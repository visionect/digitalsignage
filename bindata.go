package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// static_bower_json reads file data from disk. It returns an error on failure.
func static_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower.json",
		"static/bower.json",
	)
}

// static_index_html reads file data from disk. It returns an error on failure.
func static_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/index.html",
		"static/index.html",
	)
}

// static_screen_html reads file data from disk. It returns an error on failure.
func static_screen_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/screen.html",
		"static/screen.html",
	)
}

// static_bower_components_core_meta_core_meta_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_meta_core_meta_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-meta/core-meta.html",
		"static/bower_components/core-meta/core-meta.html",
	)
}

// static_bower_components_core_meta_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_meta_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-meta/bower.json",
		"static/bower_components/core-meta/bower.json",
	)
}

// static_bower_components_core_meta_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_meta_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-meta/index.html",
		"static/bower_components/core-meta/index.html",
	)
}

// static_bower_components_core_meta_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_meta_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-meta/README.md",
		"static/bower_components/core-meta/README.md",
	)
}

// static_bower_components_core_meta_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_meta_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-meta/demo.html",
		"static/bower_components/core-meta/demo.html",
	)
}

// static_bower_components_core_iconset_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/bower.json",
		"static/bower_components/core-iconset/bower.json",
	)
}

// static_bower_components_core_iconset_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/index.html",
		"static/bower_components/core-iconset/index.html",
	)
}

// static_bower_components_core_iconset_core_iconset_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_core_iconset_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/core-iconset.html",
		"static/bower_components/core-iconset/core-iconset.html",
	)
}

// static_bower_components_core_iconset_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/README.md",
		"static/bower_components/core-iconset/README.md",
	)
}

// static_bower_components_core_iconset_my_icons_png reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_my_icons_png() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/my-icons.png",
		"static/bower_components/core-iconset/my-icons.png",
	)
}

// static_bower_components_core_iconset_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/demo.html",
		"static/bower_components/core-iconset/demo.html",
	)
}

// static_bower_components_core_iconset_my_icons_big_png reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_my_icons_big_png() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset/my-icons-big.png",
		"static/bower_components/core-iconset/my-icons-big.png",
	)
}

// static_bower_components_paper_shadow_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/bower.json",
		"static/bower_components/paper-shadow/bower.json",
	)
}

// static_bower_components_paper_shadow_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/index.html",
		"static/bower_components/paper-shadow/index.html",
	)
}

// static_bower_components_paper_shadow_paper_shadow_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_paper_shadow_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/paper-shadow.html",
		"static/bower_components/paper-shadow/paper-shadow.html",
	)
}

// static_bower_components_paper_shadow_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/metadata.html",
		"static/bower_components/paper-shadow/metadata.html",
	)
}

// static_bower_components_paper_shadow_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/README.md",
		"static/bower_components/paper-shadow/README.md",
	)
}

// static_bower_components_paper_shadow_paper_shadow_css reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_paper_shadow_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/paper-shadow.css",
		"static/bower_components/paper-shadow/paper-shadow.css",
	)
}

// static_bower_components_paper_shadow_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/demo.html",
		"static/bower_components/paper-shadow/demo.html",
	)
}

// static_bower_components_paper_shadow_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/test/index.html",
		"static/bower_components/paper-shadow/test/index.html",
	)
}

// static_bower_components_paper_shadow_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_shadow_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-shadow/test/basic.html",
		"static/bower_components/paper-shadow/test/basic.html",
	)
}

// static_bower_components_polymer_build_log reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_build_log() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/build.log",
		"static/bower_components/polymer/build.log",
	)
}

// static_bower_components_polymer_layout_html reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_layout_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/layout.html",
		"static/bower_components/polymer/layout.html",
	)
}

// static_bower_components_polymer_polymer_html reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_polymer_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/polymer.html",
		"static/bower_components/polymer/polymer.html",
	)
}

// static_bower_components_polymer_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/bower.json",
		"static/bower_components/polymer/bower.json",
	)
}

// static_bower_components_polymer_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/README.md",
		"static/bower_components/polymer/README.md",
	)
}

// static_bower_components_polymer_polymer_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_polymer_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/polymer.min.js",
		"static/bower_components/polymer/polymer.min.js",
	)
}

// static_bower_components_polymer_polymer_js reads file data from disk. It returns an error on failure.
func static_bower_components_polymer_polymer_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/polymer/polymer.js",
		"static/bower_components/polymer/polymer.js",
	)
}

// static_bower_components_font_roboto_roboto_html reads file data from disk. It returns an error on failure.
func static_bower_components_font_roboto_roboto_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/font-roboto/roboto.html",
		"static/bower_components/font-roboto/roboto.html",
	)
}

// static_bower_components_paper_spinner_paper_spinner_css reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_paper_spinner_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/paper-spinner.css",
		"static/bower_components/paper-spinner/paper-spinner.css",
	)
}

// static_bower_components_paper_spinner_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/bower.json",
		"static/bower_components/paper-spinner/bower.json",
	)
}

// static_bower_components_paper_spinner_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/index.html",
		"static/bower_components/paper-spinner/index.html",
	)
}

// static_bower_components_paper_spinner_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/README.md",
		"static/bower_components/paper-spinner/README.md",
	)
}

// static_bower_components_paper_spinner_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/demo.html",
		"static/bower_components/paper-spinner/demo.html",
	)
}

// static_bower_components_paper_spinner_paper_spinner_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_spinner_paper_spinner_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-spinner/paper-spinner.html",
		"static/bower_components/paper-spinner/paper-spinner.html",
	)
}

// static_bower_components_file_input_package_json reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_package_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/package.json",
		"static/bower_components/file-input/package.json",
	)
}

// static_bower_components_file_input_gruntfile_js reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_gruntfile_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/gruntfile.js",
		"static/bower_components/file-input/gruntfile.js",
	)
}

// static_bower_components_file_input_file_input_html reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_file_input_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/file-input.html",
		"static/bower_components/file-input/file-input.html",
	)
}

// static_bower_components_file_input_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/bower.json",
		"static/bower_components/file-input/bower.json",
	)
}

// static_bower_components_file_input_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/index.html",
		"static/bower_components/file-input/index.html",
	)
}

// static_bower_components_file_input_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/README.md",
		"static/bower_components/file-input/README.md",
	)
}

// static_bower_components_file_input_grunt_tasks_karma_js reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_grunt_tasks_karma_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/grunt_tasks/karma.js",
		"static/bower_components/file-input/grunt_tasks/karma.js",
	)
}

// static_bower_components_file_input_grunt_tasks_jshint_js reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_grunt_tasks_jshint_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/grunt_tasks/jshint.js",
		"static/bower_components/file-input/grunt_tasks/jshint.js",
	)
}

// static_bower_components_file_input_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/demo.html",
		"static/bower_components/file-input/demo.html",
	)
}

// static_bower_components_file_input_file_input_js reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_file_input_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/file-input.js",
		"static/bower_components/file-input/file-input.js",
	)
}

// static_bower_components_file_input_license reads file data from disk. It returns an error on failure.
func static_bower_components_file_input_license() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/file-input/LICENSE",
		"static/bower_components/file-input/LICENSE",
	)
}

// static_bower_components_paper_fab_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/bower.json",
		"static/bower_components/paper-fab/bower.json",
	)
}

// static_bower_components_paper_fab_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/index.html",
		"static/bower_components/paper-fab/index.html",
	)
}

// static_bower_components_paper_fab_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/metadata.html",
		"static/bower_components/paper-fab/metadata.html",
	)
}

// static_bower_components_paper_fab_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/README.md",
		"static/bower_components/paper-fab/README.md",
	)
}

// static_bower_components_paper_fab_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/demo.html",
		"static/bower_components/paper-fab/demo.html",
	)
}

// static_bower_components_paper_fab_paper_fab_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_paper_fab_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/paper-fab.html",
		"static/bower_components/paper-fab/paper-fab.html",
	)
}

// static_bower_components_paper_fab_test_a11y_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_test_a11y_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/test/a11y.html",
		"static/bower_components/paper-fab/test/a11y.html",
	)
}

// static_bower_components_paper_fab_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/test/index.html",
		"static/bower_components/paper-fab/test/index.html",
	)
}

// static_bower_components_paper_fab_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_fab_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-fab/test/basic.html",
		"static/bower_components/paper-fab/test/basic.html",
	)
}

// static_bower_components_core_icon_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/bower.json",
		"static/bower_components/core-icon/bower.json",
	)
}

// static_bower_components_core_icon_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/index.html",
		"static/bower_components/core-icon/index.html",
	)
}

// static_bower_components_core_icon_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/metadata.html",
		"static/bower_components/core-icon/metadata.html",
	)
}

// static_bower_components_core_icon_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/README.md",
		"static/bower_components/core-icon/README.md",
	)
}

// static_bower_components_core_icon_core_icon_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_core_icon_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/core-icon.html",
		"static/bower_components/core-icon/core-icon.html",
	)
}

// static_bower_components_core_icon_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/demo.html",
		"static/bower_components/core-icon/demo.html",
	)
}

// static_bower_components_core_icon_core_icon_css reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_core_icon_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon/core-icon.css",
		"static/bower_components/core-icon/core-icon.css",
	)
}

// static_bower_components_core_media_query_core_media_query_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_media_query_core_media_query_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-media-query/core-media-query.html",
		"static/bower_components/core-media-query/core-media-query.html",
	)
}

// static_bower_components_core_media_query_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_media_query_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-media-query/bower.json",
		"static/bower_components/core-media-query/bower.json",
	)
}

// static_bower_components_core_media_query_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_media_query_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-media-query/index.html",
		"static/bower_components/core-media-query/index.html",
	)
}

// static_bower_components_core_media_query_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_media_query_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-media-query/README.md",
		"static/bower_components/core-media-query/README.md",
	)
}

// static_bower_components_core_media_query_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_media_query_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-media-query/demo.html",
		"static/bower_components/core-media-query/demo.html",
	)
}

// static_bower_components_core_focusable_polymer_mixin_js reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_polymer_mixin_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/polymer-mixin.js",
		"static/bower_components/core-focusable/polymer-mixin.js",
	)
}

// static_bower_components_core_focusable_core_focusable_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_core_focusable_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/core-focusable.html",
		"static/bower_components/core-focusable/core-focusable.html",
	)
}

// static_bower_components_core_focusable_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/bower.json",
		"static/bower_components/core-focusable/bower.json",
	)
}

// static_bower_components_core_focusable_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/README.md",
		"static/bower_components/core-focusable/README.md",
	)
}

// static_bower_components_core_focusable_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/demo.html",
		"static/bower_components/core-focusable/demo.html",
	)
}

// static_bower_components_core_focusable_core_focusable_js reads file data from disk. It returns an error on failure.
func static_bower_components_core_focusable_core_focusable_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-focusable/core-focusable.js",
		"static/bower_components/core-focusable/core-focusable.js",
	)
}

// static_bower_components_core_icon_button_core_icon_button_css reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_core_icon_button_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/core-icon-button.css",
		"static/bower_components/core-icon-button/core-icon-button.css",
	)
}

// static_bower_components_core_icon_button_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/bower.json",
		"static/bower_components/core-icon-button/bower.json",
	)
}

// static_bower_components_core_icon_button_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/index.html",
		"static/bower_components/core-icon-button/index.html",
	)
}

// static_bower_components_core_icon_button_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/metadata.html",
		"static/bower_components/core-icon-button/metadata.html",
	)
}

// static_bower_components_core_icon_button_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/README.md",
		"static/bower_components/core-icon-button/README.md",
	)
}

// static_bower_components_core_icon_button_core_icon_button_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_core_icon_button_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/core-icon-button.html",
		"static/bower_components/core-icon-button/core-icon-button.html",
	)
}

// static_bower_components_core_icon_button_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icon_button_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icon-button/demo.html",
		"static/bower_components/core-icon-button/demo.html",
	)
}

// static_bower_components_core_header_panel_core_header_panel_css reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_core_header_panel_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/core-header-panel.css",
		"static/bower_components/core-header-panel/core-header-panel.css",
	)
}

// static_bower_components_core_header_panel_core_header_panel_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_core_header_panel_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/core-header-panel.html",
		"static/bower_components/core-header-panel/core-header-panel.html",
	)
}

// static_bower_components_core_header_panel_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/bower.json",
		"static/bower_components/core-header-panel/bower.json",
	)
}

// static_bower_components_core_header_panel_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/index.html",
		"static/bower_components/core-header-panel/index.html",
	)
}

// static_bower_components_core_header_panel_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/metadata.html",
		"static/bower_components/core-header-panel/metadata.html",
	)
}

// static_bower_components_core_header_panel_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/README.md",
		"static/bower_components/core-header-panel/README.md",
	)
}

// static_bower_components_core_header_panel_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_header_panel_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-header-panel/demo.html",
		"static/bower_components/core-header-panel/demo.html",
	)
}

// static_bower_components_core_resizable_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/bower.json",
		"static/bower_components/core-resizable/bower.json",
	)
}

// static_bower_components_core_resizable_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/index.html",
		"static/bower_components/core-resizable/index.html",
	)
}

// static_bower_components_core_resizable_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/README.md",
		"static/bower_components/core-resizable/README.md",
	)
}

// static_bower_components_core_resizable_core_resizable_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_core_resizable_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/core-resizable.html",
		"static/bower_components/core-resizable/core-resizable.html",
	)
}

// static_bower_components_core_resizable_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/test/index.html",
		"static/bower_components/core-resizable/test/index.html",
	)
}

// static_bower_components_core_resizable_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/test/basic.html",
		"static/bower_components/core-resizable/test/basic.html",
	)
}

// static_bower_components_core_resizable_test_test_elements_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_resizable_test_test_elements_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-resizable/test/test-elements.html",
		"static/bower_components/core-resizable/test/test-elements.html",
	)
}

// static_bower_components_paper_toast_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/bower.json",
		"static/bower_components/paper-toast/bower.json",
	)
}

// static_bower_components_paper_toast_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/index.html",
		"static/bower_components/paper-toast/index.html",
	)
}

// static_bower_components_paper_toast_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/metadata.html",
		"static/bower_components/paper-toast/metadata.html",
	)
}

// static_bower_components_paper_toast_paper_toast_css reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_paper_toast_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/paper-toast.css",
		"static/bower_components/paper-toast/paper-toast.css",
	)
}

// static_bower_components_paper_toast_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/README.md",
		"static/bower_components/paper-toast/README.md",
	)
}

// static_bower_components_paper_toast_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/demo.html",
		"static/bower_components/paper-toast/demo.html",
	)
}

// static_bower_components_paper_toast_paper_toast_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_toast_paper_toast_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-toast/paper-toast.html",
		"static/bower_components/paper-toast/paper-toast.html",
	)
}

// static_bower_components_core_ajax_core_ajax_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_core_ajax_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/core-ajax.html",
		"static/bower_components/core-ajax/core-ajax.html",
	)
}

// static_bower_components_core_ajax_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/bower.json",
		"static/bower_components/core-ajax/bower.json",
	)
}

// static_bower_components_core_ajax_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/index.html",
		"static/bower_components/core-ajax/index.html",
	)
}

// static_bower_components_core_ajax_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/metadata.html",
		"static/bower_components/core-ajax/metadata.html",
	)
}

// static_bower_components_core_ajax_core_xhr_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_core_xhr_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/core-xhr.html",
		"static/bower_components/core-ajax/core-xhr.html",
	)
}

// static_bower_components_core_ajax_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/README.md",
		"static/bower_components/core-ajax/README.md",
	)
}

// static_bower_components_core_ajax_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/demo.html",
		"static/bower_components/core-ajax/demo.html",
	)
}

// static_bower_components_core_ajax_demo_progress_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_demo_progress_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/demo-progress.html",
		"static/bower_components/core-ajax/demo-progress.html",
	)
}

// static_bower_components_core_ajax_test_core_ajax_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_test_core_ajax_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/test/core-ajax.html",
		"static/bower_components/core-ajax/test/core-ajax.html",
	)
}

// static_bower_components_core_ajax_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/test/index.html",
		"static/bower_components/core-ajax/test/index.html",
	)
}

// static_bower_components_core_ajax_test_core_ajax_progress_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_test_core_ajax_progress_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/test/core-ajax-progress.html",
		"static/bower_components/core-ajax/test/core-ajax-progress.html",
	)
}

// static_bower_components_core_ajax_test_core_ajax_race_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_ajax_test_core_ajax_race_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-ajax/test/core-ajax-race.html",
		"static/bower_components/core-ajax/test/core-ajax-race.html",
	)
}

// static_bower_components_core_toolbar_core_toolbar_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_core_toolbar_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/core-toolbar.html",
		"static/bower_components/core-toolbar/core-toolbar.html",
	)
}

// static_bower_components_core_toolbar_core_toolbar_css reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_core_toolbar_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/core-toolbar.css",
		"static/bower_components/core-toolbar/core-toolbar.css",
	)
}

// static_bower_components_core_toolbar_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/bower.json",
		"static/bower_components/core-toolbar/bower.json",
	)
}

// static_bower_components_core_toolbar_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/index.html",
		"static/bower_components/core-toolbar/index.html",
	)
}

// static_bower_components_core_toolbar_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/metadata.html",
		"static/bower_components/core-toolbar/metadata.html",
	)
}

// static_bower_components_core_toolbar_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/README.md",
		"static/bower_components/core-toolbar/README.md",
	)
}

// static_bower_components_core_toolbar_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/demo.html",
		"static/bower_components/core-toolbar/demo.html",
	)
}

// static_bower_components_core_toolbar_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/test/index.html",
		"static/bower_components/core-toolbar/test/index.html",
	)
}

// static_bower_components_core_toolbar_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_toolbar_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-toolbar/test/basic.html",
		"static/bower_components/core-toolbar/test/basic.html",
	)
}

// static_bower_components_paper_icon_button_paper_icon_button_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_paper_icon_button_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/paper-icon-button.html",
		"static/bower_components/paper-icon-button/paper-icon-button.html",
	)
}

// static_bower_components_paper_icon_button_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/bower.json",
		"static/bower_components/paper-icon-button/bower.json",
	)
}

// static_bower_components_paper_icon_button_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/index.html",
		"static/bower_components/paper-icon-button/index.html",
	)
}

// static_bower_components_paper_icon_button_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/metadata.html",
		"static/bower_components/paper-icon-button/metadata.html",
	)
}

// static_bower_components_paper_icon_button_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/README.md",
		"static/bower_components/paper-icon-button/README.md",
	)
}

// static_bower_components_paper_icon_button_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/demo.html",
		"static/bower_components/paper-icon-button/demo.html",
	)
}

// static_bower_components_paper_icon_button_test_a11y_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_test_a11y_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/test/a11y.html",
		"static/bower_components/paper-icon-button/test/a11y.html",
	)
}

// static_bower_components_paper_icon_button_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/test/index.html",
		"static/bower_components/paper-icon-button/test/index.html",
	)
}

// static_bower_components_paper_icon_button_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_icon_button_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-icon-button/test/basic.html",
		"static/bower_components/paper-icon-button/test/basic.html",
	)
}

// static_bower_components_core_iconset_svg_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/bower.json",
		"static/bower_components/core-iconset-svg/bower.json",
	)
}

// static_bower_components_core_iconset_svg_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/index.html",
		"static/bower_components/core-iconset-svg/index.html",
	)
}

// static_bower_components_core_iconset_svg_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/README.md",
		"static/bower_components/core-iconset-svg/README.md",
	)
}

// static_bower_components_core_iconset_svg_core_iconset_svg_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_core_iconset_svg_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/core-iconset-svg.html",
		"static/bower_components/core-iconset-svg/core-iconset-svg.html",
	)
}

// static_bower_components_core_iconset_svg_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/demo.html",
		"static/bower_components/core-iconset-svg/demo.html",
	)
}

// static_bower_components_core_iconset_svg_svg_sample_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_iconset_svg_svg_sample_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-iconset-svg/svg-sample-icons.html",
		"static/bower_components/core-iconset-svg/svg-sample-icons.html",
	)
}

// static_bower_components_webcomponentsjs_package_json reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_package_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/package.json",
		"static/bower_components/webcomponentsjs/package.json",
	)
}

// static_bower_components_webcomponentsjs_build_log reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_build_log() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/build.log",
		"static/bower_components/webcomponentsjs/build.log",
	)
}

// static_bower_components_webcomponentsjs_customelements_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_customelements_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/CustomElements.js",
		"static/bower_components/webcomponentsjs/CustomElements.js",
	)
}

// static_bower_components_webcomponentsjs_webcomponents_lite_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_webcomponents_lite_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/webcomponents-lite.js",
		"static/bower_components/webcomponentsjs/webcomponents-lite.js",
	)
}

// static_bower_components_webcomponentsjs_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/bower.json",
		"static/bower_components/webcomponentsjs/bower.json",
	)
}

// static_bower_components_webcomponentsjs_customelements_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_customelements_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/CustomElements.min.js",
		"static/bower_components/webcomponentsjs/CustomElements.min.js",
	)
}

// static_bower_components_webcomponentsjs_htmlimports_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_htmlimports_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/HTMLImports.min.js",
		"static/bower_components/webcomponentsjs/HTMLImports.min.js",
	)
}

// static_bower_components_webcomponentsjs_webcomponents_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_webcomponents_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/webcomponents.js",
		"static/bower_components/webcomponentsjs/webcomponents.js",
	)
}

// static_bower_components_webcomponentsjs_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/README.md",
		"static/bower_components/webcomponentsjs/README.md",
	)
}

// static_bower_components_webcomponentsjs_shadowdom_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_shadowdom_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/ShadowDOM.js",
		"static/bower_components/webcomponentsjs/ShadowDOM.js",
	)
}

// static_bower_components_webcomponentsjs_shadowdom_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_shadowdom_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/ShadowDOM.min.js",
		"static/bower_components/webcomponentsjs/ShadowDOM.min.js",
	)
}

// static_bower_components_webcomponentsjs_webcomponents_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_webcomponents_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/webcomponents.min.js",
		"static/bower_components/webcomponentsjs/webcomponents.min.js",
	)
}

// static_bower_components_webcomponentsjs_htmlimports_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_htmlimports_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/HTMLImports.js",
		"static/bower_components/webcomponentsjs/HTMLImports.js",
	)
}

// static_bower_components_webcomponentsjs_webcomponents_lite_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_webcomponentsjs_webcomponents_lite_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/webcomponentsjs/webcomponents-lite.min.js",
		"static/bower_components/webcomponentsjs/webcomponents-lite.min.js",
	)
}

// static_bower_components_ajax_form_package_json reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_package_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/package.json",
		"static/bower_components/ajax-form/package.json",
	)
}

// static_bower_components_ajax_form_gruntfile_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_gruntfile_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/gruntfile.js",
		"static/bower_components/ajax-form/gruntfile.js",
	)
}

// static_bower_components_ajax_form_demo_resources_alertify_min_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_demo_resources_alertify_min_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/demo_resources/alertify.min.js",
		"static/bower_components/ajax-form/demo_resources/alertify.min.js",
	)
}

// static_bower_components_ajax_form_demo_resources_sinon_server_1_10_2_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_demo_resources_sinon_server_1_10_2_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/demo_resources/sinon-server-1.10.2.js",
		"static/bower_components/ajax-form/demo_resources/sinon-server-1.10.2.js",
	)
}

// static_bower_components_ajax_form_demo_resources_alertify_default_css reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_demo_resources_alertify_default_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/demo_resources/alertify.default.css",
		"static/bower_components/ajax-form/demo_resources/alertify.default.css",
	)
}

// static_bower_components_ajax_form_demo_resources_alertify_core_css reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_demo_resources_alertify_core_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/demo_resources/alertify.core.css",
		"static/bower_components/ajax-form/demo_resources/alertify.core.css",
	)
}

// static_bower_components_ajax_form_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/bower.json",
		"static/bower_components/ajax-form/bower.json",
	)
}

// static_bower_components_ajax_form_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/index.html",
		"static/bower_components/ajax-form/index.html",
	)
}

// static_bower_components_ajax_form_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/README.md",
		"static/bower_components/ajax-form/README.md",
	)
}

// static_bower_components_ajax_form_wct_conf_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_wct_conf_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/wct.conf.js",
		"static/bower_components/ajax-form/wct.conf.js",
	)
}

// static_bower_components_ajax_form_grunt_tasks_karma_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_grunt_tasks_karma_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/grunt_tasks/karma.js",
		"static/bower_components/ajax-form/grunt_tasks/karma.js",
	)
}

// static_bower_components_ajax_form_grunt_tasks_jshint_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_grunt_tasks_jshint_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/grunt_tasks/jshint.js",
		"static/bower_components/ajax-form/grunt_tasks/jshint.js",
	)
}

// static_bower_components_ajax_form_ajax_form_js reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_ajax_form_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/ajax-form.js",
		"static/bower_components/ajax-form/ajax-form.js",
	)
}

// static_bower_components_ajax_form_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/demo.html",
		"static/bower_components/ajax-form/demo.html",
	)
}

// static_bower_components_ajax_form_ajax_form_html reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_ajax_form_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/ajax-form.html",
		"static/bower_components/ajax-form/ajax-form.html",
	)
}

// static_bower_components_ajax_form_license reads file data from disk. It returns an error on failure.
func static_bower_components_ajax_form_license() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/ajax-form/LICENSE",
		"static/bower_components/ajax-form/LICENSE",
	)
}

// static_bower_components_paper_ripple_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/bower.json",
		"static/bower_components/paper-ripple/bower.json",
	)
}

// static_bower_components_paper_ripple_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/index.html",
		"static/bower_components/paper-ripple/index.html",
	)
}

// static_bower_components_paper_ripple_paper_ripple_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_paper_ripple_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/paper-ripple.html",
		"static/bower_components/paper-ripple/paper-ripple.html",
	)
}

// static_bower_components_paper_ripple_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/metadata.html",
		"static/bower_components/paper-ripple/metadata.html",
	)
}

// static_bower_components_paper_ripple_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/README.md",
		"static/bower_components/paper-ripple/README.md",
	)
}

// static_bower_components_paper_ripple_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/demo.html",
		"static/bower_components/paper-ripple/demo.html",
	)
}

// static_bower_components_paper_ripple_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/test/index.html",
		"static/bower_components/paper-ripple/test/index.html",
	)
}

// static_bower_components_paper_ripple_test_position_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_ripple_test_position_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-ripple/test/position.html",
		"static/bower_components/paper-ripple/test/position.html",
	)
}

// static_bower_components_core_overlay_gulpfile_js reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_gulpfile_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/gulpfile.js",
		"static/bower_components/core-overlay/gulpfile.js",
	)
}

// static_bower_components_core_overlay_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/bower.json",
		"static/bower_components/core-overlay/bower.json",
	)
}

// static_bower_components_core_overlay_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/index.html",
		"static/bower_components/core-overlay/index.html",
	)
}

// static_bower_components_core_overlay_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/README.md",
		"static/bower_components/core-overlay/README.md",
	)
}

// static_bower_components_core_overlay_core_key_helper_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_core_key_helper_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/core-key-helper.html",
		"static/bower_components/core-overlay/core-key-helper.html",
	)
}

// static_bower_components_core_overlay_core_overlay_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_core_overlay_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/core-overlay.html",
		"static/bower_components/core-overlay/core-overlay.html",
	)
}

// static_bower_components_core_overlay_tests_runner_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_runner_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/runner.html",
		"static/bower_components/core-overlay/tests/runner.html",
	)
}

// static_bower_components_core_overlay_tests_js_htmltests_js reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_js_htmltests_js() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/js/htmltests.js",
		"static/bower_components/core-overlay/tests/js/htmltests.js",
	)
}

// static_bower_components_core_overlay_tests_html_core_overlay_quick_close_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_html_core_overlay_quick_close_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/html/core-overlay-quick-close.html",
		"static/bower_components/core-overlay/tests/html/core-overlay-quick-close.html",
	)
}

// static_bower_components_core_overlay_tests_html_core_overlay_positioning_margin_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_html_core_overlay_positioning_margin_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/html/core-overlay-positioning-margin.html",
		"static/bower_components/core-overlay/tests/html/core-overlay-positioning-margin.html",
	)
}

// static_bower_components_core_overlay_tests_html_core_overlay_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_html_core_overlay_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/html/core-overlay-basic.html",
		"static/bower_components/core-overlay/tests/html/core-overlay-basic.html",
	)
}

// static_bower_components_core_overlay_tests_html_core_overlay_positioning_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_html_core_overlay_positioning_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/html/core-overlay-positioning.html",
		"static/bower_components/core-overlay/tests/html/core-overlay-positioning.html",
	)
}

// static_bower_components_core_overlay_tests_html_core_overlay_scroll_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_html_core_overlay_scroll_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/html/core-overlay-scroll.html",
		"static/bower_components/core-overlay/tests/html/core-overlay-scroll.html",
	)
}

// static_bower_components_core_overlay_tests_tests_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_tests_tests_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/tests/tests.json",
		"static/bower_components/core-overlay/tests/tests.json",
	)
}

// static_bower_components_core_overlay_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/demo.html",
		"static/bower_components/core-overlay/demo.html",
	)
}

// static_bower_components_core_overlay_core_overlay_layer_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_overlay_core_overlay_layer_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-overlay/core-overlay-layer.html",
		"static/bower_components/core-overlay/core-overlay-layer.html",
	)
}

// static_bower_components_core_component_page_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/bower.json",
		"static/bower_components/core-component-page/bower.json",
	)
}

// static_bower_components_core_component_page_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/index.html",
		"static/bower_components/core-component-page/index.html",
	)
}

// static_bower_components_core_component_page_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/README.md",
		"static/bower_components/core-component-page/README.md",
	)
}

// static_bower_components_core_component_page_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/demo.html",
		"static/bower_components/core-component-page/demo.html",
	)
}

// static_bower_components_core_component_page_bowager_logo_png reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_bowager_logo_png() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/bowager-logo.png",
		"static/bower_components/core-component-page/bowager-logo.png",
	)
}

// static_bower_components_core_component_page_core_component_page_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_component_page_core_component_page_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-component-page/core-component-page.html",
		"static/bower_components/core-component-page/core-component-page.html",
	)
}

// static_bower_components_core_icons_communication_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_communication_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/communication-icons.html",
		"static/bower_components/core-icons/communication-icons.html",
	)
}

// static_bower_components_core_icons_device_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_device_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/device-icons.html",
		"static/bower_components/core-icons/device-icons.html",
	)
}

// static_bower_components_core_icons_editor_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_editor_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/editor-icons.html",
		"static/bower_components/core-icons/editor-icons.html",
	)
}

// static_bower_components_core_icons_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/bower.json",
		"static/bower_components/core-icons/bower.json",
	)
}

// static_bower_components_core_icons_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/index.html",
		"static/bower_components/core-icons/index.html",
	)
}

// static_bower_components_core_icons_maps_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_maps_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/maps-icons.html",
		"static/bower_components/core-icons/maps-icons.html",
	)
}

// static_bower_components_core_icons_hardware_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_hardware_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/hardware-icons.html",
		"static/bower_components/core-icons/hardware-icons.html",
	)
}

// static_bower_components_core_icons_av_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_av_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/av-icons.html",
		"static/bower_components/core-icons/av-icons.html",
	)
}

// static_bower_components_core_icons_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/README.md",
		"static/bower_components/core-icons/README.md",
	)
}

// static_bower_components_core_icons_image_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_image_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/image-icons.html",
		"static/bower_components/core-icons/image-icons.html",
	)
}

// static_bower_components_core_icons_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/demo.html",
		"static/bower_components/core-icons/demo.html",
	)
}

// static_bower_components_core_icons_core_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_core_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/core-icons.html",
		"static/bower_components/core-icons/core-icons.html",
	)
}

// static_bower_components_core_icons_png_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_png_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/png-icons.html",
		"static/bower_components/core-icons/png-icons.html",
	)
}

// static_bower_components_core_icons_notification_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_notification_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/notification-icons.html",
		"static/bower_components/core-icons/notification-icons.html",
	)
}

// static_bower_components_core_icons_social_icons_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_icons_social_icons_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-icons/social-icons.html",
		"static/bower_components/core-icons/social-icons.html",
	)
}

// static_bower_components_core_transition_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/bower.json",
		"static/bower_components/core-transition/bower.json",
	)
}

// static_bower_components_core_transition_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/index.html",
		"static/bower_components/core-transition/index.html",
	)
}

// static_bower_components_core_transition_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/README.md",
		"static/bower_components/core-transition/README.md",
	)
}

// static_bower_components_core_transition_core_transition_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_core_transition_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/core-transition.html",
		"static/bower_components/core-transition/core-transition.html",
	)
}

// static_bower_components_core_transition_core_transition_css_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_core_transition_css_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/core-transition-css.html",
		"static/bower_components/core-transition/core-transition-css.html",
	)
}

// static_bower_components_core_transition_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/demo.html",
		"static/bower_components/core-transition/demo.html",
	)
}

// static_bower_components_core_transition_core_transition_overlay_css reads file data from disk. It returns an error on failure.
func static_bower_components_core_transition_core_transition_overlay_css() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/core-transition/core-transition-overlay.css",
		"static/bower_components/core-transition/core-transition-overlay.css",
	)
}

// static_bower_components_paper_button_paper_button_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_paper_button_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/paper-button.html",
		"static/bower_components/paper-button/paper-button.html",
	)
}

// static_bower_components_paper_button_bower_json reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_bower_json() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/bower.json",
		"static/bower_components/paper-button/bower.json",
	)
}

// static_bower_components_paper_button_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/index.html",
		"static/bower_components/paper-button/index.html",
	)
}

// static_bower_components_paper_button_paper_button_base_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_paper_button_base_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/paper-button-base.html",
		"static/bower_components/paper-button/paper-button-base.html",
	)
}

// static_bower_components_paper_button_metadata_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_metadata_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/metadata.html",
		"static/bower_components/paper-button/metadata.html",
	)
}

// static_bower_components_paper_button_readme_md reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_readme_md() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/README.md",
		"static/bower_components/paper-button/README.md",
	)
}

// static_bower_components_paper_button_demo_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_demo_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/demo.html",
		"static/bower_components/paper-button/demo.html",
	)
}

// static_bower_components_paper_button_test_a11y_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_test_a11y_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/test/a11y.html",
		"static/bower_components/paper-button/test/a11y.html",
	)
}

// static_bower_components_paper_button_test_index_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_test_index_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/test/index.html",
		"static/bower_components/paper-button/test/index.html",
	)
}

// static_bower_components_paper_button_test_basic_html reads file data from disk. It returns an error on failure.
func static_bower_components_paper_button_test_basic_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/bower_components/paper-button/test/basic.html",
		"static/bower_components/paper-button/test/basic.html",
	)
}

// static_elements_visionect_images_html reads file data from disk. It returns an error on failure.
func static_elements_visionect_images_html() ([]byte, error) {
	return bindata_read(
		"/home/matevz/projects/digitalsignage/static/elements/visionect-images.html",
		"static/elements/visionect-images.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"static/bower.json": static_bower_json,
	"static/index.html": static_index_html,
	"static/screen.html": static_screen_html,
	"static/bower_components/core-meta/core-meta.html": static_bower_components_core_meta_core_meta_html,
	"static/bower_components/core-meta/bower.json": static_bower_components_core_meta_bower_json,
	"static/bower_components/core-meta/index.html": static_bower_components_core_meta_index_html,
	"static/bower_components/core-meta/README.md": static_bower_components_core_meta_readme_md,
	"static/bower_components/core-meta/demo.html": static_bower_components_core_meta_demo_html,
	"static/bower_components/core-iconset/bower.json": static_bower_components_core_iconset_bower_json,
	"static/bower_components/core-iconset/index.html": static_bower_components_core_iconset_index_html,
	"static/bower_components/core-iconset/core-iconset.html": static_bower_components_core_iconset_core_iconset_html,
	"static/bower_components/core-iconset/README.md": static_bower_components_core_iconset_readme_md,
	"static/bower_components/core-iconset/my-icons.png": static_bower_components_core_iconset_my_icons_png,
	"static/bower_components/core-iconset/demo.html": static_bower_components_core_iconset_demo_html,
	"static/bower_components/core-iconset/my-icons-big.png": static_bower_components_core_iconset_my_icons_big_png,
	"static/bower_components/paper-shadow/bower.json": static_bower_components_paper_shadow_bower_json,
	"static/bower_components/paper-shadow/index.html": static_bower_components_paper_shadow_index_html,
	"static/bower_components/paper-shadow/paper-shadow.html": static_bower_components_paper_shadow_paper_shadow_html,
	"static/bower_components/paper-shadow/metadata.html": static_bower_components_paper_shadow_metadata_html,
	"static/bower_components/paper-shadow/README.md": static_bower_components_paper_shadow_readme_md,
	"static/bower_components/paper-shadow/paper-shadow.css": static_bower_components_paper_shadow_paper_shadow_css,
	"static/bower_components/paper-shadow/demo.html": static_bower_components_paper_shadow_demo_html,
	"static/bower_components/paper-shadow/test/index.html": static_bower_components_paper_shadow_test_index_html,
	"static/bower_components/paper-shadow/test/basic.html": static_bower_components_paper_shadow_test_basic_html,
	"static/bower_components/polymer/build.log": static_bower_components_polymer_build_log,
	"static/bower_components/polymer/layout.html": static_bower_components_polymer_layout_html,
	"static/bower_components/polymer/polymer.html": static_bower_components_polymer_polymer_html,
	"static/bower_components/polymer/bower.json": static_bower_components_polymer_bower_json,
	"static/bower_components/polymer/README.md": static_bower_components_polymer_readme_md,
	"static/bower_components/polymer/polymer.min.js": static_bower_components_polymer_polymer_min_js,
	"static/bower_components/polymer/polymer.js": static_bower_components_polymer_polymer_js,
	"static/bower_components/font-roboto/roboto.html": static_bower_components_font_roboto_roboto_html,
	"static/bower_components/paper-spinner/paper-spinner.css": static_bower_components_paper_spinner_paper_spinner_css,
	"static/bower_components/paper-spinner/bower.json": static_bower_components_paper_spinner_bower_json,
	"static/bower_components/paper-spinner/index.html": static_bower_components_paper_spinner_index_html,
	"static/bower_components/paper-spinner/README.md": static_bower_components_paper_spinner_readme_md,
	"static/bower_components/paper-spinner/demo.html": static_bower_components_paper_spinner_demo_html,
	"static/bower_components/paper-spinner/paper-spinner.html": static_bower_components_paper_spinner_paper_spinner_html,
	"static/bower_components/file-input/package.json": static_bower_components_file_input_package_json,
	"static/bower_components/file-input/gruntfile.js": static_bower_components_file_input_gruntfile_js,
	"static/bower_components/file-input/file-input.html": static_bower_components_file_input_file_input_html,
	"static/bower_components/file-input/bower.json": static_bower_components_file_input_bower_json,
	"static/bower_components/file-input/index.html": static_bower_components_file_input_index_html,
	"static/bower_components/file-input/README.md": static_bower_components_file_input_readme_md,
	"static/bower_components/file-input/grunt_tasks/karma.js": static_bower_components_file_input_grunt_tasks_karma_js,
	"static/bower_components/file-input/grunt_tasks/jshint.js": static_bower_components_file_input_grunt_tasks_jshint_js,
	"static/bower_components/file-input/demo.html": static_bower_components_file_input_demo_html,
	"static/bower_components/file-input/file-input.js": static_bower_components_file_input_file_input_js,
	"static/bower_components/file-input/LICENSE": static_bower_components_file_input_license,
	"static/bower_components/paper-fab/bower.json": static_bower_components_paper_fab_bower_json,
	"static/bower_components/paper-fab/index.html": static_bower_components_paper_fab_index_html,
	"static/bower_components/paper-fab/metadata.html": static_bower_components_paper_fab_metadata_html,
	"static/bower_components/paper-fab/README.md": static_bower_components_paper_fab_readme_md,
	"static/bower_components/paper-fab/demo.html": static_bower_components_paper_fab_demo_html,
	"static/bower_components/paper-fab/paper-fab.html": static_bower_components_paper_fab_paper_fab_html,
	"static/bower_components/paper-fab/test/a11y.html": static_bower_components_paper_fab_test_a11y_html,
	"static/bower_components/paper-fab/test/index.html": static_bower_components_paper_fab_test_index_html,
	"static/bower_components/paper-fab/test/basic.html": static_bower_components_paper_fab_test_basic_html,
	"static/bower_components/core-icon/bower.json": static_bower_components_core_icon_bower_json,
	"static/bower_components/core-icon/index.html": static_bower_components_core_icon_index_html,
	"static/bower_components/core-icon/metadata.html": static_bower_components_core_icon_metadata_html,
	"static/bower_components/core-icon/README.md": static_bower_components_core_icon_readme_md,
	"static/bower_components/core-icon/core-icon.html": static_bower_components_core_icon_core_icon_html,
	"static/bower_components/core-icon/demo.html": static_bower_components_core_icon_demo_html,
	"static/bower_components/core-icon/core-icon.css": static_bower_components_core_icon_core_icon_css,
	"static/bower_components/core-media-query/core-media-query.html": static_bower_components_core_media_query_core_media_query_html,
	"static/bower_components/core-media-query/bower.json": static_bower_components_core_media_query_bower_json,
	"static/bower_components/core-media-query/index.html": static_bower_components_core_media_query_index_html,
	"static/bower_components/core-media-query/README.md": static_bower_components_core_media_query_readme_md,
	"static/bower_components/core-media-query/demo.html": static_bower_components_core_media_query_demo_html,
	"static/bower_components/core-focusable/polymer-mixin.js": static_bower_components_core_focusable_polymer_mixin_js,
	"static/bower_components/core-focusable/core-focusable.html": static_bower_components_core_focusable_core_focusable_html,
	"static/bower_components/core-focusable/bower.json": static_bower_components_core_focusable_bower_json,
	"static/bower_components/core-focusable/README.md": static_bower_components_core_focusable_readme_md,
	"static/bower_components/core-focusable/demo.html": static_bower_components_core_focusable_demo_html,
	"static/bower_components/core-focusable/core-focusable.js": static_bower_components_core_focusable_core_focusable_js,
	"static/bower_components/core-icon-button/core-icon-button.css": static_bower_components_core_icon_button_core_icon_button_css,
	"static/bower_components/core-icon-button/bower.json": static_bower_components_core_icon_button_bower_json,
	"static/bower_components/core-icon-button/index.html": static_bower_components_core_icon_button_index_html,
	"static/bower_components/core-icon-button/metadata.html": static_bower_components_core_icon_button_metadata_html,
	"static/bower_components/core-icon-button/README.md": static_bower_components_core_icon_button_readme_md,
	"static/bower_components/core-icon-button/core-icon-button.html": static_bower_components_core_icon_button_core_icon_button_html,
	"static/bower_components/core-icon-button/demo.html": static_bower_components_core_icon_button_demo_html,
	"static/bower_components/core-header-panel/core-header-panel.css": static_bower_components_core_header_panel_core_header_panel_css,
	"static/bower_components/core-header-panel/core-header-panel.html": static_bower_components_core_header_panel_core_header_panel_html,
	"static/bower_components/core-header-panel/bower.json": static_bower_components_core_header_panel_bower_json,
	"static/bower_components/core-header-panel/index.html": static_bower_components_core_header_panel_index_html,
	"static/bower_components/core-header-panel/metadata.html": static_bower_components_core_header_panel_metadata_html,
	"static/bower_components/core-header-panel/README.md": static_bower_components_core_header_panel_readme_md,
	"static/bower_components/core-header-panel/demo.html": static_bower_components_core_header_panel_demo_html,
	"static/bower_components/core-resizable/bower.json": static_bower_components_core_resizable_bower_json,
	"static/bower_components/core-resizable/index.html": static_bower_components_core_resizable_index_html,
	"static/bower_components/core-resizable/README.md": static_bower_components_core_resizable_readme_md,
	"static/bower_components/core-resizable/core-resizable.html": static_bower_components_core_resizable_core_resizable_html,
	"static/bower_components/core-resizable/test/index.html": static_bower_components_core_resizable_test_index_html,
	"static/bower_components/core-resizable/test/basic.html": static_bower_components_core_resizable_test_basic_html,
	"static/bower_components/core-resizable/test/test-elements.html": static_bower_components_core_resizable_test_test_elements_html,
	"static/bower_components/paper-toast/bower.json": static_bower_components_paper_toast_bower_json,
	"static/bower_components/paper-toast/index.html": static_bower_components_paper_toast_index_html,
	"static/bower_components/paper-toast/metadata.html": static_bower_components_paper_toast_metadata_html,
	"static/bower_components/paper-toast/paper-toast.css": static_bower_components_paper_toast_paper_toast_css,
	"static/bower_components/paper-toast/README.md": static_bower_components_paper_toast_readme_md,
	"static/bower_components/paper-toast/demo.html": static_bower_components_paper_toast_demo_html,
	"static/bower_components/paper-toast/paper-toast.html": static_bower_components_paper_toast_paper_toast_html,
	"static/bower_components/core-ajax/core-ajax.html": static_bower_components_core_ajax_core_ajax_html,
	"static/bower_components/core-ajax/bower.json": static_bower_components_core_ajax_bower_json,
	"static/bower_components/core-ajax/index.html": static_bower_components_core_ajax_index_html,
	"static/bower_components/core-ajax/metadata.html": static_bower_components_core_ajax_metadata_html,
	"static/bower_components/core-ajax/core-xhr.html": static_bower_components_core_ajax_core_xhr_html,
	"static/bower_components/core-ajax/README.md": static_bower_components_core_ajax_readme_md,
	"static/bower_components/core-ajax/demo.html": static_bower_components_core_ajax_demo_html,
	"static/bower_components/core-ajax/demo-progress.html": static_bower_components_core_ajax_demo_progress_html,
	"static/bower_components/core-ajax/test/core-ajax.html": static_bower_components_core_ajax_test_core_ajax_html,
	"static/bower_components/core-ajax/test/index.html": static_bower_components_core_ajax_test_index_html,
	"static/bower_components/core-ajax/test/core-ajax-progress.html": static_bower_components_core_ajax_test_core_ajax_progress_html,
	"static/bower_components/core-ajax/test/core-ajax-race.html": static_bower_components_core_ajax_test_core_ajax_race_html,
	"static/bower_components/core-toolbar/core-toolbar.html": static_bower_components_core_toolbar_core_toolbar_html,
	"static/bower_components/core-toolbar/core-toolbar.css": static_bower_components_core_toolbar_core_toolbar_css,
	"static/bower_components/core-toolbar/bower.json": static_bower_components_core_toolbar_bower_json,
	"static/bower_components/core-toolbar/index.html": static_bower_components_core_toolbar_index_html,
	"static/bower_components/core-toolbar/metadata.html": static_bower_components_core_toolbar_metadata_html,
	"static/bower_components/core-toolbar/README.md": static_bower_components_core_toolbar_readme_md,
	"static/bower_components/core-toolbar/demo.html": static_bower_components_core_toolbar_demo_html,
	"static/bower_components/core-toolbar/test/index.html": static_bower_components_core_toolbar_test_index_html,
	"static/bower_components/core-toolbar/test/basic.html": static_bower_components_core_toolbar_test_basic_html,
	"static/bower_components/paper-icon-button/paper-icon-button.html": static_bower_components_paper_icon_button_paper_icon_button_html,
	"static/bower_components/paper-icon-button/bower.json": static_bower_components_paper_icon_button_bower_json,
	"static/bower_components/paper-icon-button/index.html": static_bower_components_paper_icon_button_index_html,
	"static/bower_components/paper-icon-button/metadata.html": static_bower_components_paper_icon_button_metadata_html,
	"static/bower_components/paper-icon-button/README.md": static_bower_components_paper_icon_button_readme_md,
	"static/bower_components/paper-icon-button/demo.html": static_bower_components_paper_icon_button_demo_html,
	"static/bower_components/paper-icon-button/test/a11y.html": static_bower_components_paper_icon_button_test_a11y_html,
	"static/bower_components/paper-icon-button/test/index.html": static_bower_components_paper_icon_button_test_index_html,
	"static/bower_components/paper-icon-button/test/basic.html": static_bower_components_paper_icon_button_test_basic_html,
	"static/bower_components/core-iconset-svg/bower.json": static_bower_components_core_iconset_svg_bower_json,
	"static/bower_components/core-iconset-svg/index.html": static_bower_components_core_iconset_svg_index_html,
	"static/bower_components/core-iconset-svg/README.md": static_bower_components_core_iconset_svg_readme_md,
	"static/bower_components/core-iconset-svg/core-iconset-svg.html": static_bower_components_core_iconset_svg_core_iconset_svg_html,
	"static/bower_components/core-iconset-svg/demo.html": static_bower_components_core_iconset_svg_demo_html,
	"static/bower_components/core-iconset-svg/svg-sample-icons.html": static_bower_components_core_iconset_svg_svg_sample_icons_html,
	"static/bower_components/webcomponentsjs/package.json": static_bower_components_webcomponentsjs_package_json,
	"static/bower_components/webcomponentsjs/build.log": static_bower_components_webcomponentsjs_build_log,
	"static/bower_components/webcomponentsjs/CustomElements.js": static_bower_components_webcomponentsjs_customelements_js,
	"static/bower_components/webcomponentsjs/webcomponents-lite.js": static_bower_components_webcomponentsjs_webcomponents_lite_js,
	"static/bower_components/webcomponentsjs/bower.json": static_bower_components_webcomponentsjs_bower_json,
	"static/bower_components/webcomponentsjs/CustomElements.min.js": static_bower_components_webcomponentsjs_customelements_min_js,
	"static/bower_components/webcomponentsjs/HTMLImports.min.js": static_bower_components_webcomponentsjs_htmlimports_min_js,
	"static/bower_components/webcomponentsjs/webcomponents.js": static_bower_components_webcomponentsjs_webcomponents_js,
	"static/bower_components/webcomponentsjs/README.md": static_bower_components_webcomponentsjs_readme_md,
	"static/bower_components/webcomponentsjs/ShadowDOM.js": static_bower_components_webcomponentsjs_shadowdom_js,
	"static/bower_components/webcomponentsjs/ShadowDOM.min.js": static_bower_components_webcomponentsjs_shadowdom_min_js,
	"static/bower_components/webcomponentsjs/webcomponents.min.js": static_bower_components_webcomponentsjs_webcomponents_min_js,
	"static/bower_components/webcomponentsjs/HTMLImports.js": static_bower_components_webcomponentsjs_htmlimports_js,
	"static/bower_components/webcomponentsjs/webcomponents-lite.min.js": static_bower_components_webcomponentsjs_webcomponents_lite_min_js,
	"static/bower_components/ajax-form/package.json": static_bower_components_ajax_form_package_json,
	"static/bower_components/ajax-form/gruntfile.js": static_bower_components_ajax_form_gruntfile_js,
	"static/bower_components/ajax-form/demo_resources/alertify.min.js": static_bower_components_ajax_form_demo_resources_alertify_min_js,
	"static/bower_components/ajax-form/demo_resources/sinon-server-1.10.2.js": static_bower_components_ajax_form_demo_resources_sinon_server_1_10_2_js,
	"static/bower_components/ajax-form/demo_resources/alertify.default.css": static_bower_components_ajax_form_demo_resources_alertify_default_css,
	"static/bower_components/ajax-form/demo_resources/alertify.core.css": static_bower_components_ajax_form_demo_resources_alertify_core_css,
	"static/bower_components/ajax-form/bower.json": static_bower_components_ajax_form_bower_json,
	"static/bower_components/ajax-form/index.html": static_bower_components_ajax_form_index_html,
	"static/bower_components/ajax-form/README.md": static_bower_components_ajax_form_readme_md,
	"static/bower_components/ajax-form/wct.conf.js": static_bower_components_ajax_form_wct_conf_js,
	"static/bower_components/ajax-form/grunt_tasks/karma.js": static_bower_components_ajax_form_grunt_tasks_karma_js,
	"static/bower_components/ajax-form/grunt_tasks/jshint.js": static_bower_components_ajax_form_grunt_tasks_jshint_js,
	"static/bower_components/ajax-form/ajax-form.js": static_bower_components_ajax_form_ajax_form_js,
	"static/bower_components/ajax-form/demo.html": static_bower_components_ajax_form_demo_html,
	"static/bower_components/ajax-form/ajax-form.html": static_bower_components_ajax_form_ajax_form_html,
	"static/bower_components/ajax-form/LICENSE": static_bower_components_ajax_form_license,
	"static/bower_components/paper-ripple/bower.json": static_bower_components_paper_ripple_bower_json,
	"static/bower_components/paper-ripple/index.html": static_bower_components_paper_ripple_index_html,
	"static/bower_components/paper-ripple/paper-ripple.html": static_bower_components_paper_ripple_paper_ripple_html,
	"static/bower_components/paper-ripple/metadata.html": static_bower_components_paper_ripple_metadata_html,
	"static/bower_components/paper-ripple/README.md": static_bower_components_paper_ripple_readme_md,
	"static/bower_components/paper-ripple/demo.html": static_bower_components_paper_ripple_demo_html,
	"static/bower_components/paper-ripple/test/index.html": static_bower_components_paper_ripple_test_index_html,
	"static/bower_components/paper-ripple/test/position.html": static_bower_components_paper_ripple_test_position_html,
	"static/bower_components/core-overlay/gulpfile.js": static_bower_components_core_overlay_gulpfile_js,
	"static/bower_components/core-overlay/bower.json": static_bower_components_core_overlay_bower_json,
	"static/bower_components/core-overlay/index.html": static_bower_components_core_overlay_index_html,
	"static/bower_components/core-overlay/README.md": static_bower_components_core_overlay_readme_md,
	"static/bower_components/core-overlay/core-key-helper.html": static_bower_components_core_overlay_core_key_helper_html,
	"static/bower_components/core-overlay/core-overlay.html": static_bower_components_core_overlay_core_overlay_html,
	"static/bower_components/core-overlay/tests/runner.html": static_bower_components_core_overlay_tests_runner_html,
	"static/bower_components/core-overlay/tests/js/htmltests.js": static_bower_components_core_overlay_tests_js_htmltests_js,
	"static/bower_components/core-overlay/tests/html/core-overlay-quick-close.html": static_bower_components_core_overlay_tests_html_core_overlay_quick_close_html,
	"static/bower_components/core-overlay/tests/html/core-overlay-positioning-margin.html": static_bower_components_core_overlay_tests_html_core_overlay_positioning_margin_html,
	"static/bower_components/core-overlay/tests/html/core-overlay-basic.html": static_bower_components_core_overlay_tests_html_core_overlay_basic_html,
	"static/bower_components/core-overlay/tests/html/core-overlay-positioning.html": static_bower_components_core_overlay_tests_html_core_overlay_positioning_html,
	"static/bower_components/core-overlay/tests/html/core-overlay-scroll.html": static_bower_components_core_overlay_tests_html_core_overlay_scroll_html,
	"static/bower_components/core-overlay/tests/tests.json": static_bower_components_core_overlay_tests_tests_json,
	"static/bower_components/core-overlay/demo.html": static_bower_components_core_overlay_demo_html,
	"static/bower_components/core-overlay/core-overlay-layer.html": static_bower_components_core_overlay_core_overlay_layer_html,
	"static/bower_components/core-component-page/bower.json": static_bower_components_core_component_page_bower_json,
	"static/bower_components/core-component-page/index.html": static_bower_components_core_component_page_index_html,
	"static/bower_components/core-component-page/README.md": static_bower_components_core_component_page_readme_md,
	"static/bower_components/core-component-page/demo.html": static_bower_components_core_component_page_demo_html,
	"static/bower_components/core-component-page/bowager-logo.png": static_bower_components_core_component_page_bowager_logo_png,
	"static/bower_components/core-component-page/core-component-page.html": static_bower_components_core_component_page_core_component_page_html,
	"static/bower_components/core-icons/communication-icons.html": static_bower_components_core_icons_communication_icons_html,
	"static/bower_components/core-icons/device-icons.html": static_bower_components_core_icons_device_icons_html,
	"static/bower_components/core-icons/editor-icons.html": static_bower_components_core_icons_editor_icons_html,
	"static/bower_components/core-icons/bower.json": static_bower_components_core_icons_bower_json,
	"static/bower_components/core-icons/index.html": static_bower_components_core_icons_index_html,
	"static/bower_components/core-icons/maps-icons.html": static_bower_components_core_icons_maps_icons_html,
	"static/bower_components/core-icons/hardware-icons.html": static_bower_components_core_icons_hardware_icons_html,
	"static/bower_components/core-icons/av-icons.html": static_bower_components_core_icons_av_icons_html,
	"static/bower_components/core-icons/README.md": static_bower_components_core_icons_readme_md,
	"static/bower_components/core-icons/image-icons.html": static_bower_components_core_icons_image_icons_html,
	"static/bower_components/core-icons/demo.html": static_bower_components_core_icons_demo_html,
	"static/bower_components/core-icons/core-icons.html": static_bower_components_core_icons_core_icons_html,
	"static/bower_components/core-icons/png-icons.html": static_bower_components_core_icons_png_icons_html,
	"static/bower_components/core-icons/notification-icons.html": static_bower_components_core_icons_notification_icons_html,
	"static/bower_components/core-icons/social-icons.html": static_bower_components_core_icons_social_icons_html,
	"static/bower_components/core-transition/bower.json": static_bower_components_core_transition_bower_json,
	"static/bower_components/core-transition/index.html": static_bower_components_core_transition_index_html,
	"static/bower_components/core-transition/README.md": static_bower_components_core_transition_readme_md,
	"static/bower_components/core-transition/core-transition.html": static_bower_components_core_transition_core_transition_html,
	"static/bower_components/core-transition/core-transition-css.html": static_bower_components_core_transition_core_transition_css_html,
	"static/bower_components/core-transition/demo.html": static_bower_components_core_transition_demo_html,
	"static/bower_components/core-transition/core-transition-overlay.css": static_bower_components_core_transition_core_transition_overlay_css,
	"static/bower_components/paper-button/paper-button.html": static_bower_components_paper_button_paper_button_html,
	"static/bower_components/paper-button/bower.json": static_bower_components_paper_button_bower_json,
	"static/bower_components/paper-button/index.html": static_bower_components_paper_button_index_html,
	"static/bower_components/paper-button/paper-button-base.html": static_bower_components_paper_button_paper_button_base_html,
	"static/bower_components/paper-button/metadata.html": static_bower_components_paper_button_metadata_html,
	"static/bower_components/paper-button/README.md": static_bower_components_paper_button_readme_md,
	"static/bower_components/paper-button/demo.html": static_bower_components_paper_button_demo_html,
	"static/bower_components/paper-button/test/a11y.html": static_bower_components_paper_button_test_a11y_html,
	"static/bower_components/paper-button/test/index.html": static_bower_components_paper_button_test_index_html,
	"static/bower_components/paper-button/test/basic.html": static_bower_components_paper_button_test_basic_html,
	"static/elements/visionect-images.html": static_elements_visionect_images_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"bower.json": &_bintree_t{static_bower_json, map[string]*_bintree_t{
		}},
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"screen.html": &_bintree_t{static_screen_html, map[string]*_bintree_t{
		}},
		"bower_components": &_bintree_t{nil, map[string]*_bintree_t{
			"paper-fab": &_bintree_t{nil, map[string]*_bintree_t{
				"paper-fab.html": &_bintree_t{static_bower_components_paper_fab_paper_fab_html, map[string]*_bintree_t{
				}},
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"basic.html": &_bintree_t{static_bower_components_paper_fab_test_basic_html, map[string]*_bintree_t{
					}},
					"a11y.html": &_bintree_t{static_bower_components_paper_fab_test_a11y_html, map[string]*_bintree_t{
					}},
					"index.html": &_bintree_t{static_bower_components_paper_fab_test_index_html, map[string]*_bintree_t{
					}},
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_fab_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_fab_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_fab_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_fab_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_fab_demo_html, map[string]*_bintree_t{
				}},
			}},
			"core-resizable": &_bintree_t{nil, map[string]*_bintree_t{
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"index.html": &_bintree_t{static_bower_components_core_resizable_test_index_html, map[string]*_bintree_t{
					}},
					"basic.html": &_bintree_t{static_bower_components_core_resizable_test_basic_html, map[string]*_bintree_t{
					}},
					"test-elements.html": &_bintree_t{static_bower_components_core_resizable_test_test_elements_html, map[string]*_bintree_t{
					}},
				}},
				"bower.json": &_bintree_t{static_bower_components_core_resizable_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_resizable_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_resizable_readme_md, map[string]*_bintree_t{
				}},
				"core-resizable.html": &_bintree_t{static_bower_components_core_resizable_core_resizable_html, map[string]*_bintree_t{
				}},
			}},
			"paper-icon-button": &_bintree_t{nil, map[string]*_bintree_t{
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"a11y.html": &_bintree_t{static_bower_components_paper_icon_button_test_a11y_html, map[string]*_bintree_t{
					}},
					"index.html": &_bintree_t{static_bower_components_paper_icon_button_test_index_html, map[string]*_bintree_t{
					}},
					"basic.html": &_bintree_t{static_bower_components_paper_icon_button_test_basic_html, map[string]*_bintree_t{
					}},
				}},
				"paper-icon-button.html": &_bintree_t{static_bower_components_paper_icon_button_paper_icon_button_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_icon_button_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_icon_button_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_icon_button_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_icon_button_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_icon_button_demo_html, map[string]*_bintree_t{
				}},
			}},
			"paper-shadow": &_bintree_t{nil, map[string]*_bintree_t{
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"basic.html": &_bintree_t{static_bower_components_paper_shadow_test_basic_html, map[string]*_bintree_t{
					}},
					"index.html": &_bintree_t{static_bower_components_paper_shadow_test_index_html, map[string]*_bintree_t{
					}},
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_shadow_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_shadow_index_html, map[string]*_bintree_t{
				}},
				"paper-shadow.html": &_bintree_t{static_bower_components_paper_shadow_paper_shadow_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_shadow_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_shadow_readme_md, map[string]*_bintree_t{
				}},
				"paper-shadow.css": &_bintree_t{static_bower_components_paper_shadow_paper_shadow_css, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_shadow_demo_html, map[string]*_bintree_t{
				}},
			}},
			"core-focusable": &_bintree_t{nil, map[string]*_bintree_t{
				"bower.json": &_bintree_t{static_bower_components_core_focusable_bower_json, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_focusable_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_focusable_demo_html, map[string]*_bintree_t{
				}},
				"core-focusable.js": &_bintree_t{static_bower_components_core_focusable_core_focusable_js, map[string]*_bintree_t{
				}},
				"polymer-mixin.js": &_bintree_t{static_bower_components_core_focusable_polymer_mixin_js, map[string]*_bintree_t{
				}},
				"core-focusable.html": &_bintree_t{static_bower_components_core_focusable_core_focusable_html, map[string]*_bintree_t{
				}},
			}},
			"core-icon-button": &_bintree_t{nil, map[string]*_bintree_t{
				"metadata.html": &_bintree_t{static_bower_components_core_icon_button_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_icon_button_readme_md, map[string]*_bintree_t{
				}},
				"core-icon-button.html": &_bintree_t{static_bower_components_core_icon_button_core_icon_button_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_icon_button_demo_html, map[string]*_bintree_t{
				}},
				"core-icon-button.css": &_bintree_t{static_bower_components_core_icon_button_core_icon_button_css, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_icon_button_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_icon_button_index_html, map[string]*_bintree_t{
				}},
			}},
			"webcomponentsjs": &_bintree_t{nil, map[string]*_bintree_t{
				"HTMLImports.min.js": &_bintree_t{static_bower_components_webcomponentsjs_htmlimports_min_js, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_webcomponentsjs_readme_md, map[string]*_bintree_t{
				}},
				"webcomponents-lite.min.js": &_bintree_t{static_bower_components_webcomponentsjs_webcomponents_lite_min_js, map[string]*_bintree_t{
				}},
				"webcomponents-lite.js": &_bintree_t{static_bower_components_webcomponentsjs_webcomponents_lite_js, map[string]*_bintree_t{
				}},
				"CustomElements.js": &_bintree_t{static_bower_components_webcomponentsjs_customelements_js, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_webcomponentsjs_bower_json, map[string]*_bintree_t{
				}},
				"CustomElements.min.js": &_bintree_t{static_bower_components_webcomponentsjs_customelements_min_js, map[string]*_bintree_t{
				}},
				"webcomponents.js": &_bintree_t{static_bower_components_webcomponentsjs_webcomponents_js, map[string]*_bintree_t{
				}},
				"ShadowDOM.js": &_bintree_t{static_bower_components_webcomponentsjs_shadowdom_js, map[string]*_bintree_t{
				}},
				"package.json": &_bintree_t{static_bower_components_webcomponentsjs_package_json, map[string]*_bintree_t{
				}},
				"HTMLImports.js": &_bintree_t{static_bower_components_webcomponentsjs_htmlimports_js, map[string]*_bintree_t{
				}},
				"build.log": &_bintree_t{static_bower_components_webcomponentsjs_build_log, map[string]*_bintree_t{
				}},
				"webcomponents.min.js": &_bintree_t{static_bower_components_webcomponentsjs_webcomponents_min_js, map[string]*_bintree_t{
				}},
				"ShadowDOM.min.js": &_bintree_t{static_bower_components_webcomponentsjs_shadowdom_min_js, map[string]*_bintree_t{
				}},
			}},
			"core-component-page": &_bintree_t{nil, map[string]*_bintree_t{
				"core-component-page.html": &_bintree_t{static_bower_components_core_component_page_core_component_page_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_component_page_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_component_page_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_component_page_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_component_page_demo_html, map[string]*_bintree_t{
				}},
				"bowager-logo.png": &_bintree_t{static_bower_components_core_component_page_bowager_logo_png, map[string]*_bintree_t{
				}},
			}},
			"core-meta": &_bintree_t{nil, map[string]*_bintree_t{
				"index.html": &_bintree_t{static_bower_components_core_meta_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_meta_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_meta_demo_html, map[string]*_bintree_t{
				}},
				"core-meta.html": &_bintree_t{static_bower_components_core_meta_core_meta_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_meta_bower_json, map[string]*_bintree_t{
				}},
			}},
			"core-icons": &_bintree_t{nil, map[string]*_bintree_t{
				"README.md": &_bintree_t{static_bower_components_core_icons_readme_md, map[string]*_bintree_t{
				}},
				"core-icons.html": &_bintree_t{static_bower_components_core_icons_core_icons_html, map[string]*_bintree_t{
				}},
				"social-icons.html": &_bintree_t{static_bower_components_core_icons_social_icons_html, map[string]*_bintree_t{
				}},
				"editor-icons.html": &_bintree_t{static_bower_components_core_icons_editor_icons_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_icons_bower_json, map[string]*_bintree_t{
				}},
				"maps-icons.html": &_bintree_t{static_bower_components_core_icons_maps_icons_html, map[string]*_bintree_t{
				}},
				"av-icons.html": &_bintree_t{static_bower_components_core_icons_av_icons_html, map[string]*_bintree_t{
				}},
				"image-icons.html": &_bintree_t{static_bower_components_core_icons_image_icons_html, map[string]*_bintree_t{
				}},
				"device-icons.html": &_bintree_t{static_bower_components_core_icons_device_icons_html, map[string]*_bintree_t{
				}},
				"hardware-icons.html": &_bintree_t{static_bower_components_core_icons_hardware_icons_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_icons_demo_html, map[string]*_bintree_t{
				}},
				"png-icons.html": &_bintree_t{static_bower_components_core_icons_png_icons_html, map[string]*_bintree_t{
				}},
				"communication-icons.html": &_bintree_t{static_bower_components_core_icons_communication_icons_html, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_icons_index_html, map[string]*_bintree_t{
				}},
				"notification-icons.html": &_bintree_t{static_bower_components_core_icons_notification_icons_html, map[string]*_bintree_t{
				}},
			}},
			"paper-ripple": &_bintree_t{nil, map[string]*_bintree_t{
				"bower.json": &_bintree_t{static_bower_components_paper_ripple_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_ripple_index_html, map[string]*_bintree_t{
				}},
				"paper-ripple.html": &_bintree_t{static_bower_components_paper_ripple_paper_ripple_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_ripple_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_ripple_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_ripple_demo_html, map[string]*_bintree_t{
				}},
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"index.html": &_bintree_t{static_bower_components_paper_ripple_test_index_html, map[string]*_bintree_t{
					}},
					"position.html": &_bintree_t{static_bower_components_paper_ripple_test_position_html, map[string]*_bintree_t{
					}},
				}},
			}},
			"core-media-query": &_bintree_t{nil, map[string]*_bintree_t{
				"core-media-query.html": &_bintree_t{static_bower_components_core_media_query_core_media_query_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_media_query_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_media_query_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_media_query_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_media_query_demo_html, map[string]*_bintree_t{
				}},
			}},
			"core-header-panel": &_bintree_t{nil, map[string]*_bintree_t{
				"core-header-panel.html": &_bintree_t{static_bower_components_core_header_panel_core_header_panel_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_header_panel_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_header_panel_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_core_header_panel_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_header_panel_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_header_panel_demo_html, map[string]*_bintree_t{
				}},
				"core-header-panel.css": &_bintree_t{static_bower_components_core_header_panel_core_header_panel_css, map[string]*_bintree_t{
				}},
			}},
			"core-ajax": &_bintree_t{nil, map[string]*_bintree_t{
				"core-ajax.html": &_bintree_t{static_bower_components_core_ajax_core_ajax_html, map[string]*_bintree_t{
				}},
				"core-xhr.html": &_bintree_t{static_bower_components_core_ajax_core_xhr_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_ajax_demo_html, map[string]*_bintree_t{
				}},
				"demo-progress.html": &_bintree_t{static_bower_components_core_ajax_demo_progress_html, map[string]*_bintree_t{
				}},
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"core-ajax.html": &_bintree_t{static_bower_components_core_ajax_test_core_ajax_html, map[string]*_bintree_t{
					}},
					"index.html": &_bintree_t{static_bower_components_core_ajax_test_index_html, map[string]*_bintree_t{
					}},
					"core-ajax-progress.html": &_bintree_t{static_bower_components_core_ajax_test_core_ajax_progress_html, map[string]*_bintree_t{
					}},
					"core-ajax-race.html": &_bintree_t{static_bower_components_core_ajax_test_core_ajax_race_html, map[string]*_bintree_t{
					}},
				}},
				"bower.json": &_bintree_t{static_bower_components_core_ajax_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_ajax_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_core_ajax_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_ajax_readme_md, map[string]*_bintree_t{
				}},
			}},
			"core-overlay": &_bintree_t{nil, map[string]*_bintree_t{
				"core-overlay.html": &_bintree_t{static_bower_components_core_overlay_core_overlay_html, map[string]*_bintree_t{
				}},
				"tests": &_bintree_t{nil, map[string]*_bintree_t{
					"runner.html": &_bintree_t{static_bower_components_core_overlay_tests_runner_html, map[string]*_bintree_t{
					}},
					"js": &_bintree_t{nil, map[string]*_bintree_t{
						"htmltests.js": &_bintree_t{static_bower_components_core_overlay_tests_js_htmltests_js, map[string]*_bintree_t{
						}},
					}},
					"html": &_bintree_t{nil, map[string]*_bintree_t{
						"core-overlay-positioning.html": &_bintree_t{static_bower_components_core_overlay_tests_html_core_overlay_positioning_html, map[string]*_bintree_t{
						}},
						"core-overlay-scroll.html": &_bintree_t{static_bower_components_core_overlay_tests_html_core_overlay_scroll_html, map[string]*_bintree_t{
						}},
						"core-overlay-quick-close.html": &_bintree_t{static_bower_components_core_overlay_tests_html_core_overlay_quick_close_html, map[string]*_bintree_t{
						}},
						"core-overlay-positioning-margin.html": &_bintree_t{static_bower_components_core_overlay_tests_html_core_overlay_positioning_margin_html, map[string]*_bintree_t{
						}},
						"core-overlay-basic.html": &_bintree_t{static_bower_components_core_overlay_tests_html_core_overlay_basic_html, map[string]*_bintree_t{
						}},
					}},
					"tests.json": &_bintree_t{static_bower_components_core_overlay_tests_tests_json, map[string]*_bintree_t{
					}},
				}},
				"gulpfile.js": &_bintree_t{static_bower_components_core_overlay_gulpfile_js, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_overlay_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_overlay_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_overlay_readme_md, map[string]*_bintree_t{
				}},
				"core-key-helper.html": &_bintree_t{static_bower_components_core_overlay_core_key_helper_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_overlay_demo_html, map[string]*_bintree_t{
				}},
				"core-overlay-layer.html": &_bintree_t{static_bower_components_core_overlay_core_overlay_layer_html, map[string]*_bintree_t{
				}},
			}},
			"paper-spinner": &_bintree_t{nil, map[string]*_bintree_t{
				"paper-spinner.css": &_bintree_t{static_bower_components_paper_spinner_paper_spinner_css, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_spinner_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_spinner_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_spinner_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_spinner_demo_html, map[string]*_bintree_t{
				}},
				"paper-spinner.html": &_bintree_t{static_bower_components_paper_spinner_paper_spinner_html, map[string]*_bintree_t{
				}},
			}},
			"font-roboto": &_bintree_t{nil, map[string]*_bintree_t{
				"roboto.html": &_bintree_t{static_bower_components_font_roboto_roboto_html, map[string]*_bintree_t{
				}},
			}},
			"core-icon": &_bintree_t{nil, map[string]*_bintree_t{
				"core-icon.css": &_bintree_t{static_bower_components_core_icon_core_icon_css, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_icon_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_icon_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_core_icon_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_icon_readme_md, map[string]*_bintree_t{
				}},
				"core-icon.html": &_bintree_t{static_bower_components_core_icon_core_icon_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_icon_demo_html, map[string]*_bintree_t{
				}},
			}},
			"core-transition": &_bintree_t{nil, map[string]*_bintree_t{
				"core-transition-css.html": &_bintree_t{static_bower_components_core_transition_core_transition_css_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_transition_demo_html, map[string]*_bintree_t{
				}},
				"core-transition-overlay.css": &_bintree_t{static_bower_components_core_transition_core_transition_overlay_css, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_transition_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_transition_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_transition_readme_md, map[string]*_bintree_t{
				}},
				"core-transition.html": &_bintree_t{static_bower_components_core_transition_core_transition_html, map[string]*_bintree_t{
				}},
			}},
			"core-iconset": &_bintree_t{nil, map[string]*_bintree_t{
				"bower.json": &_bintree_t{static_bower_components_core_iconset_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_iconset_index_html, map[string]*_bintree_t{
				}},
				"core-iconset.html": &_bintree_t{static_bower_components_core_iconset_core_iconset_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_iconset_readme_md, map[string]*_bintree_t{
				}},
				"my-icons.png": &_bintree_t{static_bower_components_core_iconset_my_icons_png, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_iconset_demo_html, map[string]*_bintree_t{
				}},
				"my-icons-big.png": &_bintree_t{static_bower_components_core_iconset_my_icons_big_png, map[string]*_bintree_t{
				}},
			}},
			"polymer": &_bintree_t{nil, map[string]*_bintree_t{
				"build.log": &_bintree_t{static_bower_components_polymer_build_log, map[string]*_bintree_t{
				}},
				"layout.html": &_bintree_t{static_bower_components_polymer_layout_html, map[string]*_bintree_t{
				}},
				"polymer.html": &_bintree_t{static_bower_components_polymer_polymer_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_polymer_bower_json, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_polymer_readme_md, map[string]*_bintree_t{
				}},
				"polymer.min.js": &_bintree_t{static_bower_components_polymer_polymer_min_js, map[string]*_bintree_t{
				}},
				"polymer.js": &_bintree_t{static_bower_components_polymer_polymer_js, map[string]*_bintree_t{
				}},
			}},
			"paper-toast": &_bintree_t{nil, map[string]*_bintree_t{
				"paper-toast.html": &_bintree_t{static_bower_components_paper_toast_paper_toast_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_toast_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_toast_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_toast_metadata_html, map[string]*_bintree_t{
				}},
				"paper-toast.css": &_bintree_t{static_bower_components_paper_toast_paper_toast_css, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_toast_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_toast_demo_html, map[string]*_bintree_t{
				}},
			}},
			"ajax-form": &_bintree_t{nil, map[string]*_bintree_t{
				"ajax-form.html": &_bintree_t{static_bower_components_ajax_form_ajax_form_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_ajax_form_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_ajax_form_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_ajax_form_readme_md, map[string]*_bintree_t{
				}},
				"wct.conf.js": &_bintree_t{static_bower_components_ajax_form_wct_conf_js, map[string]*_bintree_t{
				}},
				"grunt_tasks": &_bintree_t{nil, map[string]*_bintree_t{
					"jshint.js": &_bintree_t{static_bower_components_ajax_form_grunt_tasks_jshint_js, map[string]*_bintree_t{
					}},
					"karma.js": &_bintree_t{static_bower_components_ajax_form_grunt_tasks_karma_js, map[string]*_bintree_t{
					}},
				}},
				"ajax-form.js": &_bintree_t{static_bower_components_ajax_form_ajax_form_js, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_ajax_form_demo_html, map[string]*_bintree_t{
				}},
				"LICENSE": &_bintree_t{static_bower_components_ajax_form_license, map[string]*_bintree_t{
				}},
				"package.json": &_bintree_t{static_bower_components_ajax_form_package_json, map[string]*_bintree_t{
				}},
				"gruntfile.js": &_bintree_t{static_bower_components_ajax_form_gruntfile_js, map[string]*_bintree_t{
				}},
				"demo_resources": &_bintree_t{nil, map[string]*_bintree_t{
					"alertify.default.css": &_bintree_t{static_bower_components_ajax_form_demo_resources_alertify_default_css, map[string]*_bintree_t{
					}},
					"alertify.core.css": &_bintree_t{static_bower_components_ajax_form_demo_resources_alertify_core_css, map[string]*_bintree_t{
					}},
					"alertify.min.js": &_bintree_t{static_bower_components_ajax_form_demo_resources_alertify_min_js, map[string]*_bintree_t{
					}},
					"sinon-server-1.10.2.js": &_bintree_t{static_bower_components_ajax_form_demo_resources_sinon_server_1_10_2_js, map[string]*_bintree_t{
					}},
				}},
			}},
			"paper-button": &_bintree_t{nil, map[string]*_bintree_t{
				"paper-button-base.html": &_bintree_t{static_bower_components_paper_button_paper_button_base_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_paper_button_metadata_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_paper_button_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_paper_button_demo_html, map[string]*_bintree_t{
				}},
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"a11y.html": &_bintree_t{static_bower_components_paper_button_test_a11y_html, map[string]*_bintree_t{
					}},
					"index.html": &_bintree_t{static_bower_components_paper_button_test_index_html, map[string]*_bintree_t{
					}},
					"basic.html": &_bintree_t{static_bower_components_paper_button_test_basic_html, map[string]*_bintree_t{
					}},
				}},
				"paper-button.html": &_bintree_t{static_bower_components_paper_button_paper_button_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_paper_button_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_paper_button_index_html, map[string]*_bintree_t{
				}},
			}},
			"file-input": &_bintree_t{nil, map[string]*_bintree_t{
				"bower.json": &_bintree_t{static_bower_components_file_input_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_file_input_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_file_input_readme_md, map[string]*_bintree_t{
				}},
				"grunt_tasks": &_bintree_t{nil, map[string]*_bintree_t{
					"jshint.js": &_bintree_t{static_bower_components_file_input_grunt_tasks_jshint_js, map[string]*_bintree_t{
					}},
					"karma.js": &_bintree_t{static_bower_components_file_input_grunt_tasks_karma_js, map[string]*_bintree_t{
					}},
				}},
				"demo.html": &_bintree_t{static_bower_components_file_input_demo_html, map[string]*_bintree_t{
				}},
				"file-input.js": &_bintree_t{static_bower_components_file_input_file_input_js, map[string]*_bintree_t{
				}},
				"LICENSE": &_bintree_t{static_bower_components_file_input_license, map[string]*_bintree_t{
				}},
				"package.json": &_bintree_t{static_bower_components_file_input_package_json, map[string]*_bintree_t{
				}},
				"gruntfile.js": &_bintree_t{static_bower_components_file_input_gruntfile_js, map[string]*_bintree_t{
				}},
				"file-input.html": &_bintree_t{static_bower_components_file_input_file_input_html, map[string]*_bintree_t{
				}},
			}},
			"core-toolbar": &_bintree_t{nil, map[string]*_bintree_t{
				"README.md": &_bintree_t{static_bower_components_core_toolbar_readme_md, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_toolbar_demo_html, map[string]*_bintree_t{
				}},
				"test": &_bintree_t{nil, map[string]*_bintree_t{
					"index.html": &_bintree_t{static_bower_components_core_toolbar_test_index_html, map[string]*_bintree_t{
					}},
					"basic.html": &_bintree_t{static_bower_components_core_toolbar_test_basic_html, map[string]*_bintree_t{
					}},
				}},
				"core-toolbar.html": &_bintree_t{static_bower_components_core_toolbar_core_toolbar_html, map[string]*_bintree_t{
				}},
				"core-toolbar.css": &_bintree_t{static_bower_components_core_toolbar_core_toolbar_css, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_toolbar_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_toolbar_index_html, map[string]*_bintree_t{
				}},
				"metadata.html": &_bintree_t{static_bower_components_core_toolbar_metadata_html, map[string]*_bintree_t{
				}},
			}},
			"core-iconset-svg": &_bintree_t{nil, map[string]*_bintree_t{
				"core-iconset-svg.html": &_bintree_t{static_bower_components_core_iconset_svg_core_iconset_svg_html, map[string]*_bintree_t{
				}},
				"demo.html": &_bintree_t{static_bower_components_core_iconset_svg_demo_html, map[string]*_bintree_t{
				}},
				"svg-sample-icons.html": &_bintree_t{static_bower_components_core_iconset_svg_svg_sample_icons_html, map[string]*_bintree_t{
				}},
				"bower.json": &_bintree_t{static_bower_components_core_iconset_svg_bower_json, map[string]*_bintree_t{
				}},
				"index.html": &_bintree_t{static_bower_components_core_iconset_svg_index_html, map[string]*_bintree_t{
				}},
				"README.md": &_bintree_t{static_bower_components_core_iconset_svg_readme_md, map[string]*_bintree_t{
				}},
			}},
		}},
		"elements": &_bintree_t{nil, map[string]*_bintree_t{
			"visionect-images.html": &_bintree_t{static_elements_visionect_images_html, map[string]*_bintree_t{
			}},
		}},
	}},
}}
