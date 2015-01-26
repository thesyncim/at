package server

import (
	"crypto/tls"
	//"github.com/thesyncim/at/server/assets"
	//"io/ioutil"
)

func LoadTLSConfig(crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {
	//fileOrAsset := func(path string, default_path string) ([]byte, error) {
	//	loadFn := ioutil.ReadFile
	//	if path == "" {
	//		loadFn = assets.Asset
	//		path = default_path
	//	}
	//
	//return loadFn(path)
	//	}

	var (
		crt  = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x93\xc9\x92\xe2\x46\x10\x86\xef\x7a\x0a\xdf\x15\x8e\x11\x3b\x1c\xe6\x90\xb5\x68\xa5\x04\xa5\x15\xe9\x86\xd4\x58\x85\xd6\xa6\x11\x94\xd0\xd3\xbb\x07\x6c\xc7\x44\xb7\x2b\xa2\x2e\xff\x1f\x19\x99\x5f\x2e\x7f\xfe\x7a\x88\x1a\x96\xfb\x07\xa6\x5e\x60\xe9\x16\x86\x80\x3e\x55\x85\x59\x16\xee\x03\x8c\x41\xcc\x31\xe6\xa4\x3d\x45\x4b\x9a\x6d\xcc\xc3\xac\x04\x17\x15\xd5\x45\x54\x67\x63\x23\x35\x04\x3c\xd4\x81\x60\xd4\x96\x74\xcb\xa0\x32\x60\x12\x52\x24\x18\x56\x42\x9d\x0f\x74\x04\x0f\x15\x6e\x84\xa0\x60\xb8\x72\x45\xd6\x7a\x22\x6b\x44\x9d\x1a\xe1\x6f\x5e\xfe\xdd\x0b\x60\xaf\xbc\xcc\x8e\x61\xaa\x5f\x73\x23\x2a\x93\xc0\xa5\xcc\xa3\x92\xc8\x84\x44\x9c\x6f\x09\x08\x94\x99\xa8\x4e\xa6\x74\xf4\x08\xcd\x18\x7a\x65\x07\x29\xfd\xdc\xb4\xef\xca\xc9\xac\x6e\x49\x3c\xc8\x34\x76\x05\x9b\xf2\x5b\x32\xdd\xf4\xcc\xb7\xa4\x05\x89\xed\x74\xa9\x25\xee\xb9\x0b\xbc\x42\xba\x70\xb5\xa3\x11\x8d\xa7\x78\x51\x1e\xe3\x09\xa4\xd3\x89\x38\xc6\xf2\xa6\x3c\x03\xd0\xfc\x40\x02\x3a\x61\x84\x3e\xdc\x92\x6a\x6c\xac\x24\xd3\xbb\x5f\xda\xf4\x8b\x26\x8b\x74\x3e\xe0\x11\xec\x57\xe5\x49\xa0\x40\x8d\x42\xe6\x31\x49\xf9\xb3\x64\x8b\xc0\x3b\x49\xe2\x85\xf6\xf9\xbb\x34\xf6\xea\xdf\x3c\xf3\xbb\x47\xa5\xf2\x0f\xab\xf3\x85\xd5\x63\xb0\x7e\xb2\x62\x29\x2d\xfe\x2f\xa2\xc6\x07\x63\x84\xf4\x95\x9d\x31\xda\xa2\x87\x92\xcd\xc4\x62\xdb\xfc\xd7\xc0\x6a\xdb\xb8\xf7\x2c\xa0\x67\x86\xc1\xc0\xfe\xc5\xf0\xad\x6c\x46\x38\xb5\xc1\x4b\x82\x37\x43\xd4\xf9\xac\xfe\xec\x53\xdd\x73\xe3\xad\x4f\xe2\xfa\xaa\xbc\x02\x3e\xe7\x3b\x7e\x9d\x3b\xd5\x01\x76\xa8\xe4\x20\x8b\xa4\xc2\x45\x42\x61\xb9\xd9\xe9\xa1\xda\x23\xbc\x0b\xfb\xca\x54\x5b\x1a\xd4\xbd\xe2\x74\x9b\x31\x6a\xc4\x14\x1d\xf1\x3a\x3e\x48\x30\x97\xad\x11\xdc\x3a\x28\x6e\xde\xe1\x6e\x38\x1f\xaa\xea\x5e\xf1\xdb\x6a\x1b\x4f\xc7\xf5\xbe\x79\x67\x33\x75\xf0\x7b\x47\xcc\x2f\xf1\xad\xc3\x51\x54\x2b\xd9\x5c\x75\x22\xed\xec\x9b\xcd\x5d\xa7\x5b\x4d\x6a\x9d\xd9\xb0\x94\x17\xa0\xd7\x3d\x98\xf6\xa8\xb5\x0b\x51\xf6\x0f\xe4\x2d\x7f\xcc\x4f\xf3\x6d\x12\x04\xba\x91\xb9\x83\xc6\x45\xb3\xd3\x1e\x37\xaa\xe4\x7f\x7d\x10\xff\x12\xad\xb3\x89\xad\x37\xda\x8f\xb5\x7b\xb9\x8e\x0f\x58\x63\x90\x14\x20\xf8\x9f\x75\xfe\xc4\x2a\x38\xa4\xf8\x1c\xcd\x9a\x28\xa8\x94\x87\x68\xbb\x38\x1c\x39\x6a\xce\xad\xdc\x47\x4e\xb9\x7a\x44\xdd\x6e\xc5\x9d\x30\xbd\x8a\x66\xa6\xaf\xe3\x2c\x3f\x3a\xa4\xdf\x78\xd5\xbb\x9e\xcd\x4f\x83\x0a\x50\x12\x9d\xd8\xd3\x37\x2b\x5f\xde\x7d\x65\x35\x19\xf9\xc1\xab\x54\x37\x00\x2d\xb7\x0f\xfb\xd5\xf4\xe0\x0d\x24\x26\xd1\xf2\xa3\xe9\xd5\xe3\x8c\x9f\x2a\x92\x79\xa3\x33\xfc\xa8\x47\x1f\x1f\x62\x88\x8f\xa7\xc5\xf5\x22\xdc\x7a\x17\x6d\x9c\x65\xa6\xc4\xea\xee\xd1\x1c\x9b\x8f\x12\xb6\xe3\x42\xbb\xd4\xf5\x7e\xa7\x19\xb6\x99\x16\x4d\xd1\x07\x3a\x45\x46\xf1\xf3\xa7\xf2\x3c\x5c\xea\x92\xef\xc7\xfc\x77\x00\x00\x00\xff\xff\x15\x45\xca\x07\xe9\x03\x00\x00")
		key  = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd3\xb7\xae\xab\x58\x00\x85\xe1\x9e\xa7\xb8\x3d\x1a\x19\x30\x60\x5c\x4c\xb1\x61\x83\x89\x9b\x1c\x3b\xf0\x21\x67\x83\x49\x4f\x3f\x3a\xb7\x9e\xd5\xae\xee\x97\xbe\x7f\x7e\xc7\x8b\x2f\x05\xfd\x71\x5c\xf0\xc7\x72\x94\x00\x78\xe2\x1f\x4d\x8c\xff\x3e\x98\xa1\x28\x42\xb4\x2b\x3c\x00\x1a\x5f\xda\xf0\x43\xd0\x81\xc7\x12\xa2\xc2\xb8\x94\x5d\xb0\x6f\xc7\x0c\x89\xb9\xb9\xa3\x30\x95\x63\x69\xd4\x8f\xa4\x00\xa0\x98\xdf\x89\xc9\x96\x82\xc0\x48\x39\x87\xc5\xf3\xf6\xa0\xa9\x5d\xbd\x2f\x6b\x96\x59\x3b\x9e\xb6\x57\xf1\x90\x34\x62\x96\xeb\x29\x65\x4b\x35\x08\x83\xad\xa1\xa7\xc8\x55\x95\x3c\xe5\x7c\x7b\xf3\xa0\x5b\xe6\xf1\xd1\x09\x00\x84\x21\x01\xb0\xf7\x60\xb9\x45\xff\x32\xef\x8a\x24\x7f\x6e\xd5\xa3\x5c\x2b\x64\xf8\xc9\x72\x97\x1d\xe1\x15\x3f\x3c\x8d\x76\x0e\x7c\x41\xda\x14\x1d\x9b\xdf\x86\x99\x77\xdb\xa9\xf9\xb2\x14\xb8\x2b\x10\xd8\x80\xc7\xc0\xf8\xe2\x81\x81\x22\x4f\x43\xdc\xe1\xed\x83\x50\x11\x3f\x25\x2c\x7a\x87\xf5\xb6\xbc\x3f\xea\xeb\xe0\xcd\x70\x88\x91\xbc\xfa\x59\xd1\xa7\xd7\x32\xba\xcc\xf6\x5a\x9c\xde\xe0\x87\x90\x16\x18\xac\xfe\x21\xf5\x98\xa6\x23\x9a\x92\x41\x9b\xf6\xd6\x97\x31\x66\x12\x07\x5e\x38\x65\x4f\x29\xd7\x40\xc4\x6e\xd6\xd6\x26\xaa\x61\x69\x7d\x43\x19\x10\xaa\xbc\x1e\xaa\x44\xc1\xce\x25\x5e\x13\xdf\x1d\x63\xc5\xf3\x81\xfc\x14\xb6\x28\x7b\xe9\xbd\x16\x5d\xfd\x43\xc0\x15\x0f\x3d\xcf\xcc\x34\xa4\x0b\x68\x76\xa5\x0b\x5b\xbe\x06\xf1\x4d\x46\x6b\x53\x81\x56\x04\xb7\x36\xa5\x6b\xe6\xdd\x2c\xe5\xa2\x62\x7a\xea\x80\xdd\x3b\x9f\x15\xcb\x3b\x79\x40\x72\xb1\x32\xe5\x77\x92\x9b\xf6\x3c\x1f\xe1\x87\x9f\x4f\x34\xd1\x2a\x2c\xab\x9a\x5b\xe3\x1d\xf8\x1e\x5f\x91\x56\xd1\x71\x4c\xa2\x49\xad\xf0\x2c\x72\x2c\x3e\x39\xb3\xeb\x21\x4b\xd9\x2a\x0f\x4c\x72\x39\x6e\x94\x23\xf6\x19\x4f\xe4\xa2\xb4\xec\xbc\xd8\x4f\xdc\x64\x36\xfe\x68\xd4\xc7\x1c\xfb\x42\x60\x6f\x55\xf5\x98\xce\xf1\x49\xe6\xb5\x74\x98\x1a\x06\x1d\xe4\x1f\x8c\xad\x6a\x38\xd9\xc3\x96\xbe\x13\x3f\x81\x53\xca\x64\xb3\xc7\x5f\x19\x51\x6c\x97\xbb\x6f\xc1\xb6\x61\x7e\x45\x6f\x24\x06\x17\xf7\x05\x10\x91\xdd\xa6\xa3\xe1\x11\xab\x41\x93\x63\xa3\xe4\x29\xc5\x9a\x18\x63\xea\x8b\x16\x68\xdf\xb7\x3c\x81\xe3\xa0\xfd\xd8\x45\x94\xe2\x3e\x2d\x9d\xe7\x41\x1a\x15\x5e\xf7\xf3\x17\xaf\x3b\x60\xaf\xaa\xb8\x75\x01\xe7\x10\x24\x93\xfc\xf6\xc0\xd8\xe6\xa3\x25\xe3\x5b\x1f\x28\xcd\xc9\xcf\xcb\x43\x3d\x49\xbf\x12\xd4\xec\xb2\x29\x24\x90\xa8\xd9\xc5\x08\xeb\x1f\x2d\xdb\x71\xdb\xc1\xa5\xaf\xca\x4c\xd1\x98\xf4\x57\x9a\x48\x93\x0f\xbe\xb1\x85\x19\x26\x79\x5e\x1f\xdc\xb7\xa7\x1b\x3b\x0c\xeb\x32\x44\xbe\x6f\xee\xbf\x3d\x3c\xdd\xe2\x2d\x60\x5d\xa7\x8a\xc2\x92\x29\xee\x5f\x12\xba\xfa\xbb\x63\x7d\x9f\xfe\xde\x26\x91\x4b\x32\xd7\x65\x63\xcc\xf0\xf8\xda\x2c\x9c\xb1\x99\xc4\xa8\x27\xfc\x9a\x62\xb2\xb5\xc3\xa5\x36\x4c\x9e\xe0\x32\x93\x8e\x10\x08\xc4\xcc\x44\x3f\x29\xe3\xba\x15\xe2\xbf\xd8\x5f\x42\x22\x82\xff\x4f\xeb\xbf\x00\x00\x00\xff\xff\xe5\x4b\xab\xc0\x7b\x03\x00\x00")
		cert tls.Certificate
	)

	if cert, err = tls.X509KeyPair(crt, key); err != nil {
		return
	}

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return
}
