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

// assets_server_tls_snakeoil_crt reads file data from disk. It returns an error on failure.
func assets_server_tls_snakeoil_crt() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/server/tls/snakeoil.crt",
		"assets/server/tls/snakeoil.crt",
	)
}

// assets_server_tls_snakeoil_key reads file data from disk. It returns an error on failure.
func assets_server_tls_snakeoil_key() ([]byte, error) {
	return bindata_read(
		"/home/thesyncim/gocode/src/github.com/HorizontDimension/ngrok/assets/server/tls/snakeoil.key",
		"assets/server/tls/snakeoil.key",
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
	"assets/server/tls/snakeoil.crt": assets_server_tls_snakeoil_crt,
	"assets/server/tls/snakeoil.key": assets_server_tls_snakeoil_key,
}
