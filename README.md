digitalsignage
==============

## Usage

1. Download latest binary for your platform from [GitHub releases](https://github.com/visionect/digitalsignage/releases)
2. Run it from Terminal or Command Prompt

    Optionally you can run it with `-h` flag to see what can be configured.

3. Set Application URL for device in Visionect Server to `http://<server_address>:<server_port>/screen` (default port is 4000)


## Development

### Backend

The whole backend is written in `digitalsignage.go`. It contains RESTful(ish) API for uploading, removing and selecting images. It also contains a template view that will display the current image.

We use [bindata](https://github.com/jteeuwen/go-bindata) to package static files in the binary.
For faster development you should use `make debug` that will only link files and not embed them in the binary (so you don't need to recompile every time you change files in `static` folder).


### Frontend

Front end is done with [Polymer](https://www.polymer-project.org/).
Everything custom is in files `static/index.html` and `static/elements/visionect-images.html`.
