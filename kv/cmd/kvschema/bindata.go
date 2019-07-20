// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// templates/kvschema.gotmpl
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesKvschemaGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xe3\xb8\x11\xfe\x2c\xfe\x8a\x39\x7f\x38\x48\xa9\x56\x4e\x17\x05\x8a\x26\x4d\x81\x6c\x92\x5e\x8d\x4b\x93\x43\x9c\xdd\xc5\x21\x08\x0a\x5a\x1a\xdb\x84\x64\x52\x4b\x52\x4e\x04\x41\xff\xbd\x18\x52\x2f\xb6\xe3\x74\x7b\x28\x0a\xdc\x7e\xd9\xc8\x24\x9f\x79\xe6\xed\xe1\xb0\x69\xa6\x27\xc0\xae\x54\x59\x6b\xb1\x5a\x5b\xf8\x78\xfa\xc7\xbf\xc0\x4f\x4a\xad\x0a\x84\xdb\xdb\x2b\xc6\x6e\x45\x8a\xd2\x60\x06\x95\xcc\x50\x83\x5d\x23\x5c\x96\x3c\x5d\x23\x74\x2b\x31\x7c\x41\x6d\x84\x92\xf0\x31\x39\x85\x90\x36\x4c\xba\xa5\x49\x74\xce\x6a\x55\xc1\x86\xd7\x20\x95\x85\xca\x20\xd8\xb5\x30\xb0\x14\x05\x02\xbe\xa6\x58\x5a\x10\x12\x52\xb5\x29\x0b\xc1\x65\x8a\xf0\x22\xec\xda\x19\xe9\x20\x12\xf6\x6b\x07\xa0\x16\x96\x0b\x09\x1c\x52\x55\xd6\xa0\x96\xbb\xbb\x80\x5b\xc6\x00\x00\xd6\xd6\x96\xe6\x6c\x3a\x7d\x79\x79\x49\xb8\xa3\x99\x28\xbd\x9a\x16\x7e\x9b\x99\xde\xce\xae\x6e\xee\xe6\x37\x1f\x3e\x26\xa7\x8c\x7d\x96\x05\x1a\x03\x1a\xbf\x55\x42\x63\x06\x8b\x1a\x78\x59\x16\x22\xe5\x8b\x02\xa1\xe0\x2f\xa0\x34\xf0\x95\x46\xcc\xc0\x2a\x22\xfa\xa2\x85\x15\x72\x15\x83\x51\x4b\xfb\xc2\x35\xb2\x4c\x18\xab\xc5\xa2\xb2\x7b\x11\xea\x69\x09\x03\xbb\x1b\x94\x04\x2e\x61\x72\x39\x87\xd9\x7c\x02\x9f\x2e\xe7\xb3\x79\xcc\xbe\xce\x1e\xff\x71\xff\xf9\x11\xbe\x5e\x3e\x3c\x5c\xde\x3d\xce\x6e\xe6\x70\xff\x00\x57\xf7\x77\xd7\xb3\xc7\xd9\xfd\xdd\x1c\xee\xff\x0e\x97\x77\xbf\xc2\xcf\xb3\xbb\xeb\x18\x50\xd8\x35\x6a\xc0\xd7\x52\x13\x77\xa5\x41\x50\xec\x30\x4b\xd8\x1c\x71\xcf\xf8\x52\x79\x32\xa6\xc4\x54\x2c\x45\x0a\x05\x97\xab\x8a\xaf\x10\x56\x6a\x8b\x5a\x0a\xb9\x82\x12\xf5\x46\x18\xca\x9e\x01\x2e\x33\x56\x88\x8d\xb0\xdc\xba\xef\x37\xee\x24\xec\x64\xda\xb6\x8c\x35\x4d\x86\x4b\x21\x11\x26\xf9\xd6\xa4\x6b\xdc\xf0\x64\xa5\x26\x6d\x3b\x9d\xc2\x95\xca\x10\x56\x28\x51\x73\xeb\x23\x3a\xec\x99\x9c\xc3\xf5\x3d\xdc\xdd\x3f\xc2\xcd\xf5\xec\x31\x61\xac\xe4\x69\x4e\x6c\x9a\x26\xf9\xc5\xff\x99\xdc\xf1\x0d\x92\x05\xb1\x29\x95\xb6\x10\xb2\x60\xb2\x12\x76\x5d\x2d\x92\x54\x6d\xa6\x2b\x57\x96\x53\xa9\x2c\x7e\xd8\xf0\xd2\x4c\xf3\xed\x84\x45\x8c\x4d\xa7\x30\x77\x26\xa0\xd4\x6a\x2b\x32\x34\x80\xd2\x0a\x2b\xd0\xc4\xae\xb6\x94\x44\x69\x4d\x4c\x1e\x82\x90\x19\xbe\xa2\x81\x05\x4f\xf3\x2e\xe7\x90\x63\xfd\x61\xcb\x8b\x0a\x09\xca\x58\xa5\x31\x61\xd3\x29\x7d\x7c\x36\x7c\x85\x67\x6c\x3a\x6d\x1a\x57\x99\xee\x34\x24\x57\x3d\xe8\x63\x5d\xa2\x81\xd3\xb6\xa5\xcd\x40\xae\xcc\xbf\x70\xdd\xb6\x31\xa0\xd6\x70\x76\xd1\x11\x6b\xe6\x04\x7a\xe6\xb1\xdb\xa4\x69\x3a\x4f\x07\x9c\xf0\x34\x4a\xe6\x29\x97\xe1\xd3\x73\xbe\x4d\x6e\x88\x7d\xdd\xfc\x39\x86\x3f\x7d\x6c\x23\x67\x1d\x65\xd6\xb6\xcc\xd6\x25\xf6\xbe\x1a\xab\xab\xd4\x42\xc3\x02\x07\x0e\xf9\x36\x71\x7f\x30\x97\x21\xcd\xe5\x0a\x0f\x89\xb6\x6d\xd3\x58\xdc\x94\x05\xb7\x08\x93\x21\x32\x13\x48\x68\xc5\x9b\xe8\xff\xdf\xc9\xb2\x73\xda\xe7\x77\x60\x3e\x73\x81\xd8\x70\xc9\x57\x68\x80\xc3\x86\x97\xb0\xd4\x6a\x33\xee\x00\x17\x51\x43\x9d\xd3\x67\x23\xf1\x0e\x1c\x80\x8c\x8e\x18\x70\xff\x06\x4f\x82\x52\xe3\x52\xbc\xd2\x0f\xbf\xb8\xbf\xd8\x51\x7e\xa3\x23\xfb\x1c\x07\xe7\x8f\xf1\xec\x39\x11\xbf\x43\xce\x87\x3c\x47\xa0\xdf\xc6\x95\x1d\x67\xa3\xd1\x56\x9a\x1a\xee\xe8\xa2\x52\xd4\x39\xdc\xba\xce\x5b\x89\x2d\x4a\x28\xb9\x46\x69\x13\xb6\xac\x64\x0a\xa1\x81\x13\x5f\x02\xd1\x91\xf3\xa1\xdf\x0b\x43\x15\x45\x70\x72\xc4\x4a\xc3\x82\x2d\xd7\xf0\x86\x73\xe0\xb9\xc1\x8f\x6f\xcf\x34\x2c\x08\xcc\x99\xf7\xda\x78\xa7\x63\x16\x74\x6e\x9f\x75\x50\xc9\x95\x92\x29\xb7\xde\xf4\x21\xa9\xd8\xb5\xba\xdb\xe7\xb1\xa3\x98\x05\x6d\x17\xa7\x39\x5a\x30\x68\x8d\x73\x7b\x4c\x08\x37\x46\xa5\xc2\x89\x89\x6b\x40\xa4\x7c\x6d\xfb\xf6\xbc\x52\x5a\xa3\x29\x95\xcc\x48\xcc\xfa\xd6\xe6\x1a\xa1\x2a\x33\x3a\xd4\xc7\x2c\x3d\x16\x85\x88\x8c\x86\x38\xc6\x2a\x86\x2d\x34\x8d\x58\x42\x72\x2d\x34\xa6\xf6\x46\xa6\x2a\x43\xed\x9a\xa3\x30\xd8\xb6\x27\x5d\xf5\x0d\x58\x11\xf5\xb9\xd2\x14\x4f\xcc\xa9\xdf\xd3\xe4\x48\x20\x42\x8c\x58\xd0\x34\x40\xc8\xb3\x8e\x64\xdb\x52\x02\x54\x91\x8d\xce\xb2\x40\x2c\x7b\xdd\x48\x13\x93\xfc\x44\xec\xf2\x98\x36\x25\xd7\x48\x54\xa2\x73\xb7\xfe\xc3\x05\x48\x51\x90\xd1\x3e\x5f\xa8\x35\x85\xf2\x00\x60\xde\x01\x90\x7f\x74\x3c\x8c\xbe\x0b\x40\xa4\x42\x16\x04\x05\xe6\x70\x01\x05\xca\x10\xf3\x88\x05\x81\xc8\x01\xde\xf1\x6e\x47\xc4\x62\x8a\xe5\xf0\xed\x53\x1d\x9e\x46\x04\x50\x88\x1e\x50\x38\x40\x34\x30\x06\x7e\x4e\x17\x34\x0b\xa2\xa6\x81\x4e\xb8\xc6\x30\x31\x16\x90\x10\xbb\x7c\xee\x14\x86\x4b\x36\x0b\x1c\xaa\xc8\x9f\xce\x0a\xcc\x9f\x93\xcb\xb2\x44\x99\x8d\x84\x0e\xeb\x8d\x05\x74\x29\xfe\x2b\x06\xb1\xa5\x18\x79\x5b\x14\xde\xa6\x49\xfe\x89\x76\xad\x32\xbf\x31\x8c\x5c\x6c\x1c\x38\x77\x98\xa1\xb3\x21\xf2\x67\x3a\x3b\x84\x33\x49\x12\x17\x9c\x37\x69\x13\x79\x0c\x68\xde\xcd\xda\x5e\xd4\x29\xec\x0e\xc2\x24\x0f\xb8\x51\x5b\x0c\xd1\x9b\x3f\x92\xcd\x0e\xf7\xfd\x74\x1e\x40\x3b\xec\xd6\x65\xf6\x88\xe7\xdb\xdf\x8b\xdf\x33\x69\x50\xdb\xff\x8f\xdf\xdd\xef\x52\x14\x4d\x03\x28\x33\xa0\xf6\x05\xea\x67\x68\xdb\x6e\xf1\x9d\x66\x19\x0e\xf4\x1a\x95\x72\x79\x4c\xbd\xdd\xa8\x85\x3c\x5d\xfb\x5b\xa5\xa6\x51\x91\x6e\x11\x2f\x52\x33\x3f\xaa\xba\x71\xaa\xa8\x49\xa8\x68\x0c\xa0\xc9\x67\x80\x92\x80\x9b\xd2\xd6\xfe\xfa\xf1\x42\x27\x55\x27\x2d\x84\x9d\x63\x4d\xc2\xc8\x2d\xe1\x65\xca\x0d\xd3\xf8\x2a\x8c\xf5\x33\xcd\x40\xa4\x0b\x7c\x18\x01\x31\xcd\xdc\xc7\x88\xbe\xa8\x2d\x82\xa1\x56\x8b\x89\x91\x04\x4e\x70\xdf\x2a\xd4\xb5\xf7\x60\xb8\x12\xd7\xdc\x3a\x15\x25\x3b\x87\x0a\xbc\xeb\xb7\x59\xab\xaa\xc8\xa0\x0f\xb1\x22\x3c\xc7\xda\x7c\x4f\x78\x69\xd6\x41\x03\x3b\xe3\x4e\x04\xe1\xd3\xf3\xb0\x37\xf6\x38\xae\x20\x34\x9a\xaa\xb0\x54\x0e\x1b\x9e\xe3\xfe\x2e\xa7\x50\x26\xea\x3a\x5b\xc4\x80\x63\x79\xa3\x71\xa5\x71\x50\xa2\xef\x29\x74\x0c\xe1\x8f\xde\xd2\x93\x78\x8e\xfa\x12\x1e\xab\xfc\x48\x19\x4b\x51\xc4\x63\x2d\x8f\xb5\xe6\x61\x62\x5a\xdf\x9b\xc6\x3a\x51\x73\x53\xca\xad\x52\x79\x55\x7e\xaa\xc7\x60\xf6\xd5\x30\xa4\xc1\x85\xbb\x69\x46\x41\xed\xae\x89\xe9\x74\x98\xad\x28\x51\x9d\x51\x9a\x6d\x6c\xba\xa6\xfa\x6a\x9a\x84\xe6\xbd\x9b\xd7\x52\x53\x71\xd2\xb8\x63\xd7\x28\x34\x1c\xb4\x3b\x6c\xdc\x47\x5f\xa8\x8f\xeb\xbe\x24\x31\x83\x1d\x61\xa6\x67\x0c\x2f\x34\xf2\xac\x06\xa3\xf4\x9b\x6b\xf5\x80\xdf\x4e\x9a\xdf\x38\x19\x6e\xf7\xc9\x45\x10\xee\x5f\x02\xbb\x79\xcf\xb1\x7e\xf7\x4a\x1d\x8c\x7c\xaa\x2d\x1a\xba\x79\x0e\xb4\x7e\xaf\x91\xfd\xbd\x86\xe6\xcd\x8d\xd3\x8b\x06\x3d\x14\xba\xfa\xc8\xb1\xde\xd5\x30\x4a\xe0\xc1\x34\x58\x69\xa3\xf4\xce\x28\xf8\xc5\x75\xad\x1f\x05\xb5\x90\x2b\x16\xdc\x2f\x97\x06\xe9\x69\x6b\x3b\xe5\xb8\x2c\x8a\xfd\x54\xf3\x6c\x27\xd1\x87\x2d\xf6\x36\xa6\x5d\xc2\x09\x4a\xe9\x0c\xbb\xd7\xea\xd8\xf7\x5d\x06\x1f\x90\xbb\x49\x68\x81\x2b\x41\xba\x62\x21\x75\x64\xbd\x4c\xa0\xcc\x0c\xbc\x50\xe3\x93\x1e\x15\x28\x57\x76\xdd\x3f\xa4\x0f\xf2\xee\x1e\x42\x7d\xee\xdd\x2b\xd9\xae\xb9\x04\x99\xc0\x57\x3a\xaf\x3b\x3b\xc2\x38\x89\x70\xef\x77\xb4\x18\x77\xe6\xe8\xf7\x6e\x00\x03\x53\xa5\x6b\x42\x73\xa5\x5a\x19\x77\xca\x3d\xfa\x39\x98\x6a\x81\xdf\x2a\x1a\x47\x53\x5e\x14\x34\xd8\xed\xc6\xe8\x65\x4f\x5a\xf0\xd5\x82\x74\xf2\x32\xbc\x25\xfe\xbb\x12\xdc\x2d\xbe\x8e\xdd\xc9\x41\x2a\x63\x90\x94\xaa\x08\x0e\x14\xc9\x3f\xe2\xc6\x7a\x14\x16\x07\x25\xf1\xa5\x36\xb3\xf4\xd0\x55\xfa\xb8\xa8\xec\x0d\x46\x87\xa3\x48\xc4\x82\x0c\x97\xa8\x81\x50\x93\x6b\x61\x52\xae\xb3\x30\xf2\x56\x92\x39\x62\xde\xd1\x4d\x5c\x79\xed\xd6\xb2\x58\xc2\x0f\x6e\xd7\x17\x5e\x88\xac\xbb\xb3\x7d\xa0\x86\x19\x6e\x51\x2d\xdf\x14\x7b\x27\x64\x17\xd0\x1f\xae\x30\x5c\x54\xcb\xef\xcc\x96\xfd\x5c\xd9\xb1\xe9\x4a\xfb\xaf\x4e\x79\x17\xd5\xd2\x5b\x47\x33\x4e\x0c\xd4\x4c\x8b\x6a\xf9\xb4\x77\xe0\xec\x79\x9c\x19\x3a\xcd\x86\xbf\x5d\x80\xf4\x7a\xba\x8f\xfd\x87\x0b\x90\xdd\x18\x30\x6c\xed\x77\x7a\x4b\x68\x9e\xce\xe4\x73\x7f\xcf\x0f\x44\x77\x26\x1d\xe7\xe2\x1d\xbe\xda\x30\x3a\x87\xdd\x60\x9d\xef\x2e\xf9\x69\xe7\x37\x86\x65\xcf\xdc\x31\xcf\xff\x93\xa7\xde\xd2\x5e\x66\xbb\x3b\xdb\x59\xff\x19\x6b\x97\xe2\xc3\x90\x5c\x8c\xe1\xfe\x00\x61\x0f\xfb\x01\x64\xf4\xbf\x04\xaa\xfb\x6c\x87\x57\x8d\x7f\x5a\xff\x3b\x00\x00\xff\xff\x14\x31\x37\x5c\x14\x14\x00\x00")

func templatesKvschemaGotmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesKvschemaGotmpl,
		"templates/kvschema.gotmpl",
	)
}

func templatesKvschemaGotmpl() (*asset, error) {
	bytes, err := templatesKvschemaGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/kvschema.gotmpl", size: 5140, mode: os.FileMode(420), modTime: time.Unix(1563647002, 0)}
	a := &asset{bytes: bytes, info: info}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"templates/kvschema.gotmpl": templatesKvschemaGotmpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"kvschema.gotmpl": &bintree{templatesKvschemaGotmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
