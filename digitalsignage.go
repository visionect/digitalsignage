package main

import (
    "encoding/json"
    "flag"
    "io"
    "io/ioutil"
    "image"
    "image/draw"
    "image/jpeg"
    "image/png"
    "time"
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    "code.google.com/p/graphics-go/graphics"
    "github.com/elazarl/go-bindata-assetfs"
)

type SelectableFile struct {
    Name string
    Selected bool
}

type WaveformRender struct {
    Name string
    Waveform int
    Render bool
}

var address = flag.String("address", "0.0.0.0", "server address")
var port = flag.String("port", "4000", "server port")
var cache = flag.String("cache", "cache", "folder for resized images")

var folder = flag.String("images", "images", "folder with images")
var selected string
var waveform int
var render bool
var renderTimer *time.Timer
var files []SelectableFile


func ListFolder() {
    files = []SelectableFile{}
    raw, _ := ioutil.ReadDir(*folder)
    noneSelected := true
    for _, f := range raw {
        ext := strings.ToLower(filepath.Ext(f.Name()))
        if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
            name := filepath.ToSlash(filepath.Join("/", *folder, f.Name()))
            if name[:2] == "//" {
                name = name[1:]
            }
            isSelected := false
            if name == selected {
                isSelected = true
                noneSelected = false
            }
            files = append(files, SelectableFile{name, isSelected})
        }
    }

    if noneSelected && len(files) > 0 {
        selected = files[0].Name
        files[0].Selected = true
    }
}


func removeCached(filename string) {
    ext := filepath.Ext(filename)

    matches, _ := filepath.Glob(filepath.Join(*cache, strings.TrimRight(filepath.Base(filename), ext) + "*" + ext))
    if matches != nil {
        for _, file := range matches {
            os.Remove(file)
        }
    }
}


func GetImages(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/images/" {
        ListFolder()

        jsonData, err := json.MarshalIndent(files, "", "    ")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        if len(files) == 0 {
            w.Write([]byte("[]"))
        } else {
            w.Write(jsonData)
        }
    } else {
        filename := filepath.FromSlash(r.URL.Path[1:])

        absFilename, err := filepath.Abs(filename)
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        wd, _ := os.Getwd()
        if !strings.HasPrefix(absFilename, wd) {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if r.Method == "GET" {
            width, err1 := strconv.Atoi(r.URL.Query().Get("width"))
            height, err2 := strconv.Atoi(r.URL.Query().Get("height"))

            var x,y int
            if x, err = strconv.Atoi(r.URL.Query().Get("x")); err != nil {
                x = -1
            }
            if y, err = strconv.Atoi(r.URL.Query().Get("y")); err != nil {
                y = -1
            }

            if err1 == nil && err2 == nil && width > 0 && height > 0 {
                ext := filepath.Ext(filename)
                resizedFilename := strings.TrimRight(filepath.Base(filename), ext)
                resizedFilename += "_" + strconv.Itoa(width)
                resizedFilename += "_" + strconv.Itoa(height)
                resizedFilename += "_" + strconv.Itoa(x)
                resizedFilename += "_" + strconv.Itoa(y)
                resizedFilename += ext

                resizedFilename = filepath.Join(*cache, resizedFilename)

                if _, err := os.Stat(resizedFilename); os.IsNotExist(err) {
                    fSrc, err := os.Open(filename)
                    if err != nil {
                        fmt.Printf("Error: %s\n", err.Error())
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                    }
                    defer fSrc.Close()

                    src, _, err := image.Decode(fSrc)
                    if err != nil {
                        fmt.Printf("Error: %s\n", err.Error())
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                    }

                    if !(src.Bounds().Max.X == width && src.Bounds().Max.Y == height) || x > -1 || y > -1 {

                        var dst draw.Image
                        if x>-1 || y>-1 {
                            if x<0 {
                                x=0
                            }
                            if y<0 {
                                y=0
                            }
                            dst = image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X-x, src.Bounds().Max.Y-y))
                            draw.Draw(dst, dst.Bounds(), src, image.Point{x, y}, draw.Src)
                        } else {
                            dst = image.NewRGBA(image.Rect(0, 0, width, height))
                            graphics.Thumbnail(dst, src)
                        }

                        fDst, err := os.Create(resizedFilename)
                        if err != nil {
                            fmt.Printf("Error: %s\n", err.Error())
                            http.Error(w, err.Error(), http.StatusInternalServerError)
                            return
                        }
                        defer fDst.Close()

                        if ext == ".png" {
                            png.Encode(fDst, dst)
                        } else {
                            jpeg.Encode(fDst, dst, &jpeg.Options{95})
                        }
                        filename = resizedFilename
                    }
                } else {
                    filename = resizedFilename
                }
            }

            http.ServeFile(w, r, filename)

            fmt.Printf("Serving image %s\n", filename)

        } else if r.Method == "DELETE" {
            if err := os.Remove(filename); err != nil {
                fmt.Printf("Error: %s\n", err.Error())
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            removeCached(filename)

            fmt.Printf("Removed %s\n", filename)

            w.WriteHeader(http.StatusNoContent)
        }
    }
}


func SelectImage(w http.ResponseWriter, r *http.Request) {
    jsonDoc, err := ioutil.ReadAll(r.Body);
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    selectedFile := SelectableFile{}
    if err := json.Unmarshal(jsonDoc, &selectedFile); err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Printf("Selected %s\n", selectedFile.Name)

    for _, file := range files {
        if file.Name == selectedFile.Name {
            selected = selectedFile.Name
            ListFolder()

            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    w.WriteHeader(http.StatusBadRequest)
}


func UploadImage(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    defer file.Close()

    out, err := os.Create(filepath.Join(*folder, header.Filename))
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer out.Close()

    _, err = io.Copy(out, file)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    removeCached(header.Filename)

    fmt.Printf("Uploaded %s\n", header.Filename)

    w.WriteHeader(http.StatusNoContent)
}


func Screen(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Content-Type") == "application/json" {
        jsonData, err := json.MarshalIndent(WaveformRender{selected, waveform, render}, "", "    ")
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Write(jsonData)
    } else {
        data, err := Asset("static/screen.html")
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        t := template.New("screen")
        t, _ = t.Parse(string(data[:]))

        templateData := struct {
            Selected string
            X string
            Y string
            Waveform int
        }{
            selected,
            r.URL.Query().Get("x"),
            r.URL.Query().Get("y"),
            waveform,
        }

        t.Execute(w, templateData)
    }
}


func Waveform(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        w.Header().Set("Content-Type", "application/json")
        jsonData, _ := json.MarshalIndent(WaveformRender{selected, waveform, render}, "", "    ")
        w.Write(jsonData)
    } else if r.Method == "POST" {
        jsonDoc, err := ioutil.ReadAll(r.Body);
        if err != nil {
            fmt.Printf("Error: %s\n", err.Error())
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        data := WaveformRender{}
        err = json.Unmarshal(jsonDoc, &data)
        if err != nil {
            fmt.Println(err)
        }

        waveform = data.Waveform
        render = true

        if renderTimer != nil {
            renderTimer.Stop()
        }

        renderTimer = time.AfterFunc(5*time.Second, func() {
            render = false
            renderTimer = nil
        })
    }
}


func main() {
    flag.Parse()

    os.Mkdir(*cache, 0755)
    os.Mkdir(*folder, 0755)
    ListFolder()

    http.HandleFunc("/images/", GetImages)
    http.HandleFunc("/select", SelectImage)
    http.HandleFunc("/upload", UploadImage)
    http.HandleFunc("/screen", Screen)
    http.HandleFunc("/waveform", Waveform)

    http.Handle("/", http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"}))

    fmt.Println("Digital signage demo by Visionect")
    fmt.Println("http://www.visionect.com")
    fmt.Printf("\nStarting server at: http://%s:%s/ \n", *address, *port)
    fmt.Printf("Serving images from: %s \n", *folder)
    fmt.Println("\nHelp available with -h, exit with Ctrl-C")

    waveform = 0
    render = false
    renderTimer = nil

    err := http.ListenAndServe(*address + ":" + *port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
