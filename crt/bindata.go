package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _server_crt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x93\xc9\x92\xe2\x46\x10\x86\xef\x7a\x0a\xdf\x15\x8e\x11\x3b\x1c\xe6\x90\xb5\x68\xa5\x04\xa5\x15\xe9\x86\xd4\x58\x85\xd6\xa6\x11\x94\xd0\xd3\xbb\x07\x6c\xc7\x44\xb7\x2b\xa2\x2e\xff\x1f\x19\x99\x5f\x2e\x7f\xfe\x7a\x88\x1a\x96\xfb\x07\xa6\x5e\x60\xe9\x16\x86\x80\x3e\x55\x85\x59\x16\xee\x03\x8c\x41\xcc\x31\xe6\xa4\x3d\x45\x4b\x9a\x6d\xcc\xc3\xac\x04\x17\x15\xd5\x45\x54\x67\x63\x23\x35\x04\x3c\xd4\x81\x60\xd4\x96\x74\xcb\xa0\x32\x60\x12\x52\x24\x18\x56\x42\x9d\x0f\x74\x04\x0f\x15\x6e\x84\xa0\x60\xb8\x72\x45\xd6\x7a\x22\x6b\x44\x9d\x1a\xe1\x6f\x5e\xfe\xdd\x0b\x60\xaf\xbc\xcc\x8e\x61\xaa\x5f\x73\x23\x2a\x93\xc0\xa5\xcc\xa3\x92\xc8\x84\x44\x9c\x6f\x09\x08\x94\x99\xa8\x4e\xa6\x74\xf4\x08\xcd\x18\x7a\x65\x07\x29\xfd\xdc\xb4\xef\xca\xc9\xac\x6e\x49\x3c\xc8\x34\x76\x05\x9b\xf2\x5b\x32\xdd\xf4\xcc\xb7\xa4\x05\x89\xed\x74\xa9\x25\xee\xb9\x0b\xbc\x42\xba\x70\xb5\xa3\x11\x8d\xa7\x78\x51\x1e\xe3\x09\xa4\xd3\x89\x38\xc6\xf2\xa6\x3c\x03\xd0\xfc\x40\x02\x3a\x61\x84\x3e\xdc\x92\x6a\x6c\xac\x24\xd3\xbb\x5f\xda\xf4\x8b\x26\x8b\x74\x3e\xe0\x11\xec\x57\xe5\x49\xa0\x40\x8d\x42\xe6\x31\x49\xf9\xb3\x64\x8b\xc0\x3b\x49\xe2\x85\xf6\xf9\xbb\x34\xf6\xea\xdf\x3c\xf3\xbb\x47\xa5\xf2\x0f\xab\xf3\x85\xd5\x63\xb0\x7e\xb2\x62\x29\x2d\xfe\x2f\xa2\xc6\x07\x63\x84\xf4\x95\x9d\x31\xda\xa2\x87\x92\xcd\xc4\x62\xdb\xfc\xd7\xc0\x6a\xdb\xb8\xf7\x2c\xa0\x67\x86\xc1\xc0\xfe\xc5\xf0\xad\x6c\x46\x38\xb5\xc1\x4b\x82\x37\x43\xd4\xf9\xac\xfe\xec\x53\xdd\x73\xe3\xad\x4f\xe2\xfa\xaa\xbc\x02\x3e\xe7\x3b\x7e\x9d\x3b\xd5\x01\x76\xa8\xe4\x20\x8b\xa4\xc2\x45\x42\x61\xb9\xd9\xe9\xa1\xda\x23\xbc\x0b\xfb\xca\x54\x5b\x1a\xd4\xbd\xe2\x74\x9b\x31\x6a\xc4\x14\x1d\xf1\x3a\x3e\x48\x30\x97\xad\x11\xdc\x3a\x28\x6e\xde\xe1\x6e\x38\x1f\xaa\xea\x5e\xf1\xdb\x6a\x1b\x4f\xc7\xf5\xbe\x79\x67\x33\x75\xf0\x7b\x47\xcc\x2f\xf1\xad\xc3\x51\x54\x2b\xd9\x5c\x75\x22\xed\xec\x9b\xcd\x5d\xa7\x5b\x4d\x6a\x9d\xd9\xb0\x94\x17\xa0\xd7\x3d\x98\xf6\xa8\xb5\x0b\x51\xf6\x0f\xe4\x2d\x7f\xcc\x4f\xf3\x6d\x12\x04\xba\x91\xb9\x83\xc6\x45\xb3\xd3\x1e\x37\xaa\xe4\x7f\x7d\x10\xff\x12\xad\xb3\x89\xad\x37\xda\x8f\xb5\x7b\xb9\x8e\x0f\x58\x63\x90\x14\x20\xf8\x9f\x75\xfe\xc4\x2a\x38\xa4\xf8\x1c\xcd\x9a\x28\xa8\x94\x87\x68\xbb\x38\x1c\x39\x6a\xce\xad\xdc\x47\x4e\xb9\x7a\x44\xdd\x6e\xc5\x9d\x30\xbd\x8a\x66\xa6\xaf\xe3\x2c\x3f\x3a\xa4\xdf\x78\xd5\xbb\x9e\xcd\x4f\x83\x0a\x50\x12\x9d\xd8\xd3\x37\x2b\x5f\xde\x7d\x65\x35\x19\xf9\xc1\xab\x54\x37\x00\x2d\xb7\x0f\xfb\xd5\xf4\xe0\x0d\x24\x26\xd1\xf2\xa3\xe9\xd5\xe3\x8c\x9f\x2a\x92\x79\xa3\x33\xfc\xa8\x47\x1f\x1f\x62\x88\x8f\xa7\xc5\xf5\x22\xdc\x7a\x17\x6d\x9c\x65\xa6\xc4\xea\xee\xd1\x1c\x9b\x8f\x12\xb6\xe3\x42\xbb\xd4\xf5\x7e\xa7\x19\xb6\x99\x16\x4d\xd1\x07\x3a\x45\x46\xf1\xf3\xa7\xf2\x3c\x5c\xea\x92\xef\xc7\xfc\x77\x00\x00\x00\xff\xff\x15\x45\xca\x07\xe9\x03\x00\x00")

func server_crt_bytes() ([]byte, error) {
	return bindata_read(
		_server_crt,
		"server.crt",
	)
}

func server_crt() (*asset, error) {
	bytes, err := server_crt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server.crt", size: 1001, mode: os.FileMode(438), modTime: time.Unix(1422284128, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_csr = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x92\x4d\x8f\xba\x48\x10\xc6\xef\x7c\x8a\xb9\x93\xcd\x80\xa2\xe2\xb1\xba\x69\x1a\x90\x6e\xa7\x79\x15\x6f\x83\x8e\xb4\xf2\xa2\x8e\x68\x23\x9f\x7e\x77\x9c\x3d\x6c\x36\xf9\x57\x52\xa7\x5f\x2a\x95\xe7\xc9\xef\xaf\x9f\x41\x84\xfa\xfc\x0d\x93\x28\xf1\x5d\x1f\x43\x42\xde\x22\x22\x52\x12\x27\x2f\xaa\x31\xdf\x47\xd3\x11\x63\x48\x2b\x0c\x02\x54\xb5\xb5\x06\x3c\x42\x80\x2a\x9e\x21\x28\x12\x68\x50\xca\x22\xa6\x88\x28\x9c\x4c\x08\xdf\x81\x8b\x53\xe4\x33\xe3\x9f\x3d\x6f\xf3\xa8\xd1\xfe\x03\xbd\xff\x43\x16\x11\xe5\xa8\x17\x5b\x39\x20\x51\xe9\xa1\xa6\x98\x90\x31\x72\x48\xc4\xc0\xa6\x60\xa6\x44\xc3\x4a\xf9\x22\x1f\xd4\x36\xe7\x92\x19\x62\xa0\x23\x6c\x7f\xbf\x33\x46\x3a\xf4\x2c\xa7\x72\x16\xb6\xee\x6d\x47\xb3\x53\x91\xf0\x3a\x6c\xf9\xa3\x4c\xc8\x91\x61\xa0\x38\xbe\x52\x2d\xf6\xcb\xa9\x23\x48\x00\x51\x91\xec\xa9\x6c\x76\xd3\xe6\x5e\x4c\x9a\x5e\xd0\x7d\x5f\xe4\xcd\xed\xf7\x00\xa3\x6e\x04\x8e\xaa\xfa\x2a\xeb\x23\x5d\x2a\x03\x81\x20\x2e\xc0\x1a\x69\xa7\x9f\xdc\x45\x8d\xab\x82\xc0\x7c\xb9\x76\x53\xbd\x47\x78\x9d\xf6\xb5\xa7\x77\x24\x69\xfa\xd5\x79\x39\x66\xad\x9c\xa0\x4f\x6c\xe7\x1b\x05\xde\xbc\xa3\xc9\xfd\x0c\xd5\x3d\xda\x3c\xe8\xea\x5b\xd3\x75\x7e\xc3\xfb\x45\x98\x4f\x46\xfb\xa3\xbd\xb0\xa9\x3e\xc4\xfd\x4a\x5a\xd7\xfc\x7e\xc6\x59\xd6\x94\x96\xbe\xca\x8c\x63\xec\xb5\x0f\x97\x84\x86\x32\xce\x5e\xcb\xb6\xa2\x02\xb7\xe9\xc1\x0b\x46\xcd\xe8\x66\xf2\xd4\x3f\x51\x34\x7f\xb7\xbe\xac\xb0\x48\x12\x97\x96\x7c\x30\x84\x6c\xd7\xc6\xf3\x4e\x76\x87\x6f\x27\xbe\x66\x76\x69\x06\x6e\x6b\xbc\xdb\xfc\x7a\x1b\x9f\x60\x63\x50\x04\xe0\x13\x40\x63\x60\xbc\xca\xf8\xb7\x0b\x84\x55\x0a\x60\x51\x04\x61\xbc\xa0\x81\x39\xe0\xce\xae\xc5\x66\x98\xb3\x27\x97\x4d\xff\x91\xac\x1e\x1d\x8b\x77\xbb\x7e\x17\x8b\xef\x19\xb7\x2e\xda\xbe\x25\x2b\x16\x2f\xda\x0d\xa5\xb4\x6e\xe6\xd1\xed\x5d\xe9\x67\x99\x66\xe9\x54\x97\xa2\xfd\xb8\x7c\xee\x6e\x4e\x60\x9e\x73\xeb\x30\x33\xdb\xa7\xcf\x79\xc9\xed\x8b\x6e\x2f\x87\xa4\x5b\xca\xf5\x85\x68\xe1\x98\xce\xda\x40\xb7\x9d\xc1\x14\x8c\xb4\x12\x5b\xde\xa1\x8e\xdc\x20\x79\x14\xf3\x9e\x6e\xbf\x42\xcb\xa8\xd5\x81\xcf\x24\x33\x2f\xd5\xc9\x4a\x0f\xa7\xc3\x5a\x76\xe6\x11\x3b\x43\x2d\xd4\x22\xd7\xaa\x23\x6e\xb4\x97\x94\x84\x3b\x7f\x16\xf6\xef\x00\x00\x00\xff\xff\xc2\x81\x18\x6a\xd5\x02\x00\x00")

func server_csr_bytes() ([]byte, error) {
	return bindata_read(
		_server_csr,
		"server.csr",
	)
}

func server_csr() (*asset, error) {
	bytes, err := server_csr_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server.csr", size: 725, mode: os.FileMode(438), modTime: time.Unix(1422284128, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_key = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd3\xb7\xae\xab\x58\x00\x85\xe1\x9e\xa7\xb8\x3d\x1a\x19\x30\x60\x5c\x4c\xb1\x61\x83\x89\x9b\x1c\x3b\xf0\x21\x67\x83\x49\x4f\x3f\x3a\xb7\x9e\xd5\xae\xee\x97\xbe\x7f\x7e\xc7\x8b\x2f\x05\xfd\x71\x5c\xf0\xc7\x72\x94\x00\x78\xe2\x1f\x4d\x8c\xff\x3e\x98\xa1\x28\x42\xb4\x2b\x3c\x00\x1a\x5f\xda\xf0\x43\xd0\x81\xc7\x12\xa2\xc2\xb8\x94\x5d\xb0\x6f\xc7\x0c\x89\xb9\xb9\xa3\x30\x95\x63\x69\xd4\x8f\xa4\x00\xa0\x98\xdf\x89\xc9\x96\x82\xc0\x48\x39\x87\xc5\xf3\xf6\xa0\xa9\x5d\xbd\x2f\x6b\x96\x59\x3b\x9e\xb6\x57\xf1\x90\x34\x62\x96\xeb\x29\x65\x4b\x35\x08\x83\xad\xa1\xa7\xc8\x55\x95\x3c\xe5\x7c\x7b\xf3\xa0\x5b\xe6\xf1\xd1\x09\x00\x84\x21\x01\xb0\xf7\x60\xb9\x45\xff\x32\xef\x8a\x24\x7f\x6e\xd5\xa3\x5c\x2b\x64\xf8\xc9\x72\x97\x1d\xe1\x15\x3f\x3c\x8d\x76\x0e\x7c\x41\xda\x14\x1d\x9b\xdf\x86\x99\x77\xdb\xa9\xf9\xb2\x14\xb8\x2b\x10\xd8\x80\xc7\xc0\xf8\xe2\x81\x81\x22\x4f\x43\xdc\xe1\xed\x83\x50\x11\x3f\x25\x2c\x7a\x87\xf5\xb6\xbc\x3f\xea\xeb\xe0\xcd\x70\x88\x91\xbc\xfa\x59\xd1\xa7\xd7\x32\xba\xcc\xf6\x5a\x9c\xde\xe0\x87\x90\x16\x18\xac\xfe\x21\xf5\x98\xa6\x23\x9a\x92\x41\x9b\xf6\xd6\x97\x31\x66\x12\x07\x5e\x38\x65\x4f\x29\xd7\x40\xc4\x6e\xd6\xd6\x26\xaa\x61\x69\x7d\x43\x19\x10\xaa\xbc\x1e\xaa\x44\xc1\xce\x25\x5e\x13\xdf\x1d\x63\xc5\xf3\x81\xfc\x14\xb6\x28\x7b\xe9\xbd\x16\x5d\xfd\x43\xc0\x15\x0f\x3d\xcf\xcc\x34\xa4\x0b\x68\x76\xa5\x0b\x5b\xbe\x06\xf1\x4d\x46\x6b\x53\x81\x56\x04\xb7\x36\xa5\x6b\xe6\xdd\x2c\xe5\xa2\x62\x7a\xea\x80\xdd\x3b\x9f\x15\xcb\x3b\x79\x40\x72\xb1\x32\xe5\x77\x92\x9b\xf6\x3c\x1f\xe1\x87\x9f\x4f\x34\xd1\x2a\x2c\xab\x9a\x5b\xe3\x1d\xf8\x1e\x5f\x91\x56\xd1\x71\x4c\xa2\x49\xad\xf0\x2c\x72\x2c\x3e\x39\xb3\xeb\x21\x4b\xd9\x2a\x0f\x4c\x72\x39\x6e\x94\x23\xf6\x19\x4f\xe4\xa2\xb4\xec\xbc\xd8\x4f\xdc\x64\x36\xfe\x68\xd4\xc7\x1c\xfb\x42\x60\x6f\x55\xf5\x98\xce\xf1\x49\xe6\xb5\x74\x98\x1a\x06\x1d\xe4\x1f\x8c\xad\x6a\x38\xd9\xc3\x96\xbe\x13\x3f\x81\x53\xca\x64\xb3\xc7\x5f\x19\x51\x6c\x97\xbb\x6f\xc1\xb6\x61\x7e\x45\x6f\x24\x06\x17\xf7\x05\x10\x91\xdd\xa6\xa3\xe1\x11\xab\x41\x93\x63\xa3\xe4\x29\xc5\x9a\x18\x63\xea\x8b\x16\x68\xdf\xb7\x3c\x81\xe3\xa0\xfd\xd8\x45\x94\xe2\x3e\x2d\x9d\xe7\x41\x1a\x15\x5e\xf7\xf3\x17\xaf\x3b\x60\xaf\xaa\xb8\x75\x01\xe7\x10\x24\x93\xfc\xf6\xc0\xd8\xe6\xa3\x25\xe3\x5b\x1f\x28\xcd\xc9\xcf\xcb\x43\x3d\x49\xbf\x12\xd4\xec\xb2\x29\x24\x90\xa8\xd9\xc5\x08\xeb\x1f\x2d\xdb\x71\xdb\xc1\xa5\xaf\xca\x4c\xd1\x98\xf4\x57\x9a\x48\x93\x0f\xbe\xb1\x85\x19\x26\x79\x5e\x1f\xdc\xb7\xa7\x1b\x3b\x0c\xeb\x32\x44\xbe\x6f\xee\xbf\x3d\x3c\xdd\xe2\x2d\x60\x5d\xa7\x8a\xc2\x92\x29\xee\x5f\x12\xba\xfa\xbb\x63\x7d\x9f\xfe\xde\x26\x91\x4b\x32\xd7\x65\x63\xcc\xf0\xf8\xda\x2c\x9c\xb1\x99\xc4\xa8\x27\xfc\x9a\x62\xb2\xb5\xc3\xa5\x36\x4c\x9e\xe0\x32\x93\x8e\x10\x08\xc4\xcc\x44\x3f\x29\xe3\xba\x15\xe2\xbf\xd8\x5f\x42\x22\x82\xff\x4f\xeb\xbf\x00\x00\x00\xff\xff\xe5\x4b\xab\xc0\x7b\x03\x00\x00")

func server_key_bytes() ([]byte, error) {
	return bindata_read(
		_server_key,
		"server.key",
	)
}

func server_key() (*asset, error) {
	bytes, err := server_key_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server.key", size: 891, mode: os.FileMode(438), modTime: time.Unix(1422284128, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _server_key_org = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd3\xb9\xb2\xa2\x4a\x00\x80\xe1\x9c\xa7\x38\xb9\x73\x8a\x7d\x9b\xaa\x09\x80\x6e\x59\x44\x68\x10\x10\xcc\x64\x47\x04\x9a\x16\xf0\xe0\xd3\xdf\xba\x13\xcf\x9f\xfe\xf1\xf7\xfd\x7f\x3a\x34\x6d\xef\x2b\xbc\x68\x5f\x28\xb4\x13\x2d\x82\x5f\x27\x98\xfd\x3d\x14\x22\x53\xf1\x1d\xed\xb8\xfa\xfd\x25\xfc\x82\x9e\x11\x66\x28\x82\x80\x02\xf0\xf4\x6d\x8f\xf5\xf4\xfb\x0b\xc0\xcb\x37\x04\x90\xff\x36\x74\xe3\x97\xa4\x88\x3a\xc3\x18\x3a\x04\x02\x27\x4b\x82\x44\x51\xbd\x24\xd0\x0a\x84\x8c\x25\x26\x25\xc1\x04\x39\xcb\xc5\xf4\x96\xce\x74\xaa\xc5\xe7\x2b\xde\x2e\x85\x54\x7e\x4a\xe3\xca\x26\x71\xa3\xde\xfd\x6b\xeb\x76\x9a\x2a\xe4\x7d\xfa\xc2\xc4\x88\xd7\x8e\xda\x5f\xca\xb8\xac\x5d\x00\xb0\x5a\x5d\xc4\xbb\x25\x65\x9e\x84\xc6\xfa\x01\x1e\xe7\xf5\xe6\xae\xca\xbd\x9a\x23\xc1\x72\x94\x1d\x81\xf7\x78\x08\xa4\xa5\x38\xd5\x8f\xc1\x7a\xa9\x86\xcd\x6f\xcd\x65\xa6\x62\x3a\x55\x81\x06\x4f\xf9\x50\x3c\x2f\xf9\xf0\xd2\x34\xe8\x88\x09\x92\x67\x55\x1b\x89\x02\xc4\xfb\x7c\xbc\xba\xe7\x4f\x96\x9b\x3f\x42\x4a\xde\x3e\x1b\xb6\xab\xd3\xf5\x30\x16\xba\xc5\x4d\xc9\x81\x2a\x5b\x27\x41\x93\xad\xa6\x4e\xd0\x35\xea\x09\xcb\x6b\x6c\x4f\xc5\x68\x71\x42\xea\xec\xd7\x85\x5d\x9b\x5e\xe4\x1d\x3f\x38\xeb\x2a\x80\xf7\xcc\xd2\x3f\x87\x27\xcd\xf9\xd8\x31\xea\xfb\x96\xc5\x80\xba\x70\xbb\xb4\x17\xe4\x88\xe9\xf9\x3d\x9f\xbd\x2c\xeb\x54\x50\x3c\xe4\xa2\x53\xb9\x80\x00\x95\x29\xca\x07\xea\x96\x70\x9e\x51\x10\x7a\xee\x49\x72\xcd\xa8\x8f\x8a\xbc\xca\x50\xf2\x1a\x82\xb9\xa3\xd0\x4d\x61\xca\xa7\xf3\xd2\xf3\x7a\x3d\x77\x87\x51\x59\x81\xcd\xa5\x22\xd9\xcf\x90\x86\x79\xe4\xd7\xf5\x85\x5d\x13\xe4\x2a\xbc\x85\x46\x1a\x98\x7d\x60\x68\xe5\xd1\x54\xfa\x85\x3b\x94\xd3\xa0\x52\xb1\xa6\xc4\xc2\x23\xb8\x45\x89\x1b\x74\x21\xdd\x9f\x50\x7f\xd6\xe4\xc8\x34\xda\x60\x8a\x78\x64\x8c\xca\x55\xd1\x0e\x81\x58\x0d\x40\xb6\x15\xf2\xd3\xa7\xc4\xe6\xc7\x71\xde\x53\xa9\x89\x19\xd5\xa3\xda\x4f\x29\x6b\xc7\xf3\x54\xeb\x2f\x02\xe2\x28\x13\x3f\x21\x24\x5c\xa7\x3e\x93\x3e\x79\x89\xd8\xc0\x99\x37\x11\xad\x79\x94\x9b\xee\x8b\x88\x65\x9b\x87\x70\x9b\x58\xba\xdc\xe4\xf7\xed\x96\x74\x05\x85\x6d\xcd\xcb\xf0\x3d\xed\xf5\xd3\x61\x3b\xfa\xd6\x80\x9b\x71\x1a\xec\xfe\x46\x46\x63\xdb\x0a\x55\x26\x39\x2d\x75\xc7\xce\xf1\x6f\x7b\xa6\x56\xb1\x11\xc7\x99\xd6\x54\xb3\x29\x8d\x05\x16\x3e\x29\x75\xbd\x28\x0b\x6d\x5a\x2b\x82\xbd\x2b\x3c\x9a\x1a\xcc\xed\xdb\x05\x1e\x8f\x6f\xcd\xf1\x4d\x27\xcb\x22\xd1\xdb\x6d\x4f\x3e\xc3\x76\x2c\x47\x4e\xe8\x59\x41\x68\xfa\xe7\x5a\x55\x72\xb8\xbc\x41\x8a\xa9\xbc\x5e\xbc\x56\xf7\xd9\xb6\x7d\xa3\x00\x37\x99\x1a\x92\x68\x4f\xfa\xb1\x96\x8f\x55\x9f\x18\x0c\xe1\x3c\x47\xcc\x2c\xda\x34\xcc\xca\x2b\x7e\x96\x46\xb0\x17\x36\x74\xdb\x76\x49\x9f\x6d\xb2\xa8\x54\x7a\x1f\xc6\xb4\xc5\x3f\x55\x78\xc4\xe4\x7a\xe2\x36\x91\x8d\x87\x2c\x6e\x08\x5e\x0b\x67\x92\x18\x0e\x8f\x5e\xe0\x56\x3f\xc8\xf7\x20\xef\xfd\xe4\xd7\x45\xbf\xd7\xee\x21\xca\x5d\xfb\x65\xa0\x91\xa1\x34\xbe\x97\x41\xf3\xbc\x6a\x02\xfb\xac\xf1\xcb\xae\x9f\x7c\xfb\xd9\xca\x73\xa7\x30\x8c\xd7\x0e\x44\x67\xef\x07\x27\x19\xcc\xb7\x6c\xbb\xd3\x06\x20\x62\x03\x0b\x36\x7f\xfe\x50\x7f\x29\x42\x0f\xfc\x9b\xe8\x7f\x01\x00\x00\xff\xff\xba\x06\xd9\xc6\xc3\x03\x00\x00")

func server_key_org_bytes() ([]byte, error) {
	return bindata_read(
		_server_key_org,
		"server.key.org",
	)
}

func server_key_org() (*asset, error) {
	bytes, err := server_key_org_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "server.key.org", size: 963, mode: os.FileMode(438), modTime: time.Unix(1422284128, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"server.crt": server_crt,
	"server.csr": server_csr,
	"server.key": server_key,
	"server.key.org": server_key_org,
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
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
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
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"server.crt": &_bintree_t{server_crt, map[string]*_bintree_t{
	}},
	"server.csr": &_bintree_t{server_csr, map[string]*_bintree_t{
	}},
	"server.key": &_bintree_t{server_key, map[string]*_bintree_t{
	}},
	"server.key.org": &_bintree_t{server_key_org, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

