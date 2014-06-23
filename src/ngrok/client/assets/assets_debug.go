// +build debug

package assets

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

// assets_client_static_img_glyphicons_halflings_png reads file data from disk. It returns an error on failure.
func assets_client_static_img_glyphicons_halflings_png() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/img/glyphicons-halflings.png",
		"assets/client/static/img/glyphicons-halflings.png",
	)
}

// assets_client_static_js_highlight_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_highlight_min_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/highlight.min.js",
		"assets/client/static/js/highlight.min.js",
	)
}

// assets_client_static_js_angular_sanitize_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_angular_sanitize_min_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/angular-sanitize.min.js",
		"assets/client/static/js/angular-sanitize.min.js",
	)
}

// assets_client_static_js_jquery_1_9_1_min_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_jquery_1_9_1_min_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/jquery-1.9.1.min.js",
		"assets/client/static/js/jquery-1.9.1.min.js",
	)
}

// assets_client_static_js_ngrok_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_ngrok_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/ngrok.js",
		"assets/client/static/js/ngrok.js",
	)
}

// assets_client_static_js_jquery_timeago_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_jquery_timeago_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/jquery.timeago.js",
		"assets/client/static/js/jquery.timeago.js",
	)
}

// assets_client_static_js_base64_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_base64_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/base64.js",
		"assets/client/static/js/base64.js",
	)
}

// assets_client_static_js_vkbeautify_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_vkbeautify_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/vkbeautify.js",
		"assets/client/static/js/vkbeautify.js",
	)
}

// assets_client_static_js_angular_js reads file data from disk. It returns an error on failure.
func assets_client_static_js_angular_js() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/js/angular.js",
		"assets/client/static/js/angular.js",
	)
}

// assets_client_static_css_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_client_static_css_bootstrap_min_css() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/css/bootstrap.min.css",
		"assets/client/static/css/bootstrap.min.css",
	)
}

// assets_client_static_css_highlight_min_css reads file data from disk. It returns an error on failure.
func assets_client_static_css_highlight_min_css() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/static/css/highlight.min.css",
		"assets/client/static/css/highlight.min.css",
	)
}

// assets_client_page_html reads file data from disk. It returns an error on failure.
func assets_client_page_html() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/page.html",
		"assets/client/page.html",
	)
}

// assets_client_tls_snakeoilca_crt reads file data from disk. It returns an error on failure.
func assets_client_tls_snakeoilca_crt() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/tls/snakeoilca.crt",
		"assets/client/tls/snakeoilca.crt",
	)
}

// assets_client_tls_ngrokroot_crt reads file data from disk. It returns an error on failure.
func assets_client_tls_ngrokroot_crt() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/client/tls/ngrokroot.crt",
		"assets/client/tls/ngrokroot.crt",
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
	"assets/client/static/img/glyphicons-halflings.png": assets_client_static_img_glyphicons_halflings_png,
	"assets/client/static/js/highlight.min.js": assets_client_static_js_highlight_min_js,
	"assets/client/static/js/angular-sanitize.min.js": assets_client_static_js_angular_sanitize_min_js,
	"assets/client/static/js/jquery-1.9.1.min.js": assets_client_static_js_jquery_1_9_1_min_js,
	"assets/client/static/js/ngrok.js": assets_client_static_js_ngrok_js,
	"assets/client/static/js/jquery.timeago.js": assets_client_static_js_jquery_timeago_js,
	"assets/client/static/js/base64.js": assets_client_static_js_base64_js,
	"assets/client/static/js/vkbeautify.js": assets_client_static_js_vkbeautify_js,
	"assets/client/static/js/angular.js": assets_client_static_js_angular_js,
	"assets/client/static/css/bootstrap.min.css": assets_client_static_css_bootstrap_min_css,
	"assets/client/static/css/highlight.min.css": assets_client_static_css_highlight_min_css,
	"assets/client/page.html": assets_client_page_html,
	"assets/client/tls/snakeoilca.crt": assets_client_tls_snakeoilca_crt,
	"assets/client/tls/ngrokroot.crt": assets_client_tls_ngrokroot_crt,
}
