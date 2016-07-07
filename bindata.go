// Code generated by go-bindata.
// sources:
// template.txt
// DO NOT EDIT!

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x5f\x6f\xdb\x36\x10\x7f\x96\x3e\xc5\x55\x4b\x02\x79\x53\x95\xf6\x35\x43\x1e\xb2\x26\xc5\x0c\xb4\xe9\x1a\x67\xdb\x43\x10\x04\x8a\x4c\x3b\xda\x64\xc9\x25\x69\xc7\x81\xa6\xef\xbe\xbb\x23\x45\xd3\x8e\xbd\xa4\x5b\xb3\xac\xc0\x02\x04\x16\x8f\xe4\xf1\x77\xff\x7e\x3c\xee\xef\xc3\xf1\x07\x38\xfd\x70\x0e\x27\xc7\xfd\x73\x38\xff\xb1\x3f\x80\xb7\xfd\x77\x27\xf0\x22\xc4\xa9\xbe\x86\x42\xc1\x58\x54\x42\x66\x5a\x0c\xe1\xfa\x0e\xa6\x42\xaa\x42\x69\xd0\x75\x5d\x26\xa0\xea\x99\xcc\x05\x8c\x64\x3d\x81\xa6\x49\x07\x3c\x6c\xdb\x34\x9c\x66\xf9\xef\xd9\x58\x90\xf0\x27\xf3\xd9\xb6\x61\x58\x4c\xa6\xb5\xd4\x10\x87\x41\x74\x7d\xa7\x85\x8a\xf0\x43\x54\x79\x3d\x2c\xaa\xf1\xfe\x6f\xaa\xae\x48\x30\x9a\x68\xfa\x19\x17\xfa\x66\x76\x9d\xe6\xf5\x64\xff\xb6\xd6\x37\xb4\xa2\xac\xc7\x34\xa3\xb4\xcc\xeb\x6a\x6e\x3f\x71\x82\xf4\x34\xcd\x4b\x90\x59\x85\x47\xa6\x7d\x3e\x45\xe1\x81\x28\x4d\xdb\x96\xa6\x44\x35\x04\x14\xf4\xc2\x50\xdf\x4d\x09\x16\xa4\xa7\xd9\x44\xa0\x0c\x50\xc7\x2c\xd7\x0d\xc2\xf3\x74\xbc\x17\xfa\xa6\x1e\x2a\xda\x13\x8e\x66\x55\x0e\xf1\xb7\xb8\x67\xa7\xdb\xd4\xf3\x35\xc4\x4d\x63\x76\xed\x14\xc9\xce\x5c\xc3\xc1\x21\xa0\xd1\x32\x9b\x28\x3a\xbb\x18\xa1\x1c\xfe\x80\x4a\xc0\xab\xb6\x4d\x70\x23\x62\xa1\x09\x5c\x9a\xfe\x92\x49\x44\x60\xbe\x07\x65\x41\xbe\xb3\x03\x4d\x33\xe6\xdb\x39\xb0\x69\x32\x34\xc3\x13\x41\x94\x46\xdd\xaa\x73\xb4\x8b\xbe\xad\xa9\x3d\xd8\x04\xeb\x4c\xe8\x99\xac\x9e\x17\x57\x13\x06\xf3\x4c\x82\x90\xfc\x5f\xcb\x30\xf8\x34\x13\xf2\x2e\x81\x4c\x8e\x15\xa1\xe4\xd4\x48\x4f\xc5\xed\x0f\xb3\xd1\x48\xc8\xf8\xe2\x92\x24\x4d\xdb\x4b\xe0\xe2\xb2\xa8\xb4\x90\xa3\x2c\xc7\x31\x86\x8c\x35\x29\x3d\xd1\x60\x32\xc1\x0f\xe1\x5b\x99\x8d\x27\xa2\xe2\x44\x60\x73\xd3\x37\xb5\xc9\x82\x00\x07\x98\x19\x34\x24\x2b\x43\x9b\x20\x34\xc3\xba\x0e\x21\xa2\x64\xc6\x4f\x9c\x8e\x3c\x9d\xe8\x49\xe8\x5c\x79\x84\x68\x4d\x8a\xbd\x04\xd4\x27\x3e\x41\x3a\xf8\xf8\xae\x5f\x01\xa6\x13\x25\x7b\x10\x8c\x6a\x09\x57\x09\xcc\x69\xb9\xd9\x4f\x5a\xf3\x7a\xea\x5c\x66\x46\xbe\xb7\xac\xe3\x71\x77\xc0\xee\x38\x84\x6c\x3a\x45\x6c\x31\x8d\x28\x48\xec\xe6\xa3\xb2\xc8\x94\xef\x77\x16\x40\x14\xa3\x9e\xf9\x3d\x69\x0f\xa5\x3d\xd4\x48\x98\xac\x7d\xb6\x6e\x30\x1f\xa6\x25\xfa\x32\x26\x71\x02\xd1\xce\xee\x30\x4a\xe8\x6c\x6f\x5e\x64\x3a\x8e\x12\x9e\x81\x52\x54\xf1\x5f\xdb\x60\x0c\xe8\xf5\x2e\x5e\x1f\x5c\x26\xf0\xf2\x35\x16\x1c\x7b\x48\x94\x4a\x90\x9b\xd2\xf7\x99\x54\x37\x59\xc9\x61\x08\xb6\xe9\xba\xf2\x74\x25\x9c\x29\xe8\x42\xe2\x87\x6e\xfb\xe3\x50\xe0\x09\x14\x19\xdc\xfe\xe2\x10\xaa\xa2\x34\x7e\x45\x16\x49\x4f\x28\xf1\x46\x71\x34\xb1\xea\x76\xbf\x99\xf7\x4c\x36\x1e\xc0\xae\x8a\x92\xc7\xa8\x67\x60\x9d\x5f\x37\x07\xeb\x61\xeb\x56\xfc\xd3\x6e\xd7\xf4\x60\xd8\x6d\x79\x3e\x22\xbb\xb6\xe5\x87\x47\x93\xde\xa7\x29\xce\xf4\x57\x59\x68\x31\xe0\xac\xe0\x64\xe9\xb9\xc4\xb7\x65\x14\x06\xad\x65\x10\xaf\x9c\x4c\x7d\x0e\xeb\xb2\xc4\x9f\x95\xda\x0d\xb9\x3a\x8a\x65\x69\xb0\xd9\x14\x1f\xbb\xda\x79\xc0\x8c\x13\x28\xbe\xc3\x64\x0a\x96\x25\x8a\xd7\x44\x3a\x98\x22\x20\x3d\x8a\x0d\x46\x0b\x0f\x69\xc2\xec\x49\xd3\x94\xdc\x4b\xf1\x3e\x16\xd7\x33\x07\x7c\x29\xa0\x43\xcd\xaa\xce\x18\x62\x2b\x24\x45\xac\xe6\x28\x1b\x0e\x23\xf0\x2a\x3c\xfd\xb9\xb2\xe9\xc2\x75\x4f\x96\x79\x0c\xeb\x78\xc1\x5f\xc5\xb4\x53\xd1\xfc\x1a\xc7\xa2\xed\x2e\x5c\x98\x10\x0b\x17\x21\x74\x66\x9e\xe9\xb5\x78\xf1\x26\x30\x34\xe8\xe2\x64\x90\x31\xb0\x41\x9e\x55\x7c\x55\x05\x43\xa1\x18\xc5\x2a\x4f\x6e\xc4\xc9\x9b\xb6\x42\xdc\xfb\x7b\x18\x11\x02\x95\xdb\x21\xe5\xfe\xf9\x02\x1b\x81\x8f\x14\x99\xb3\xfa\xd6\x32\x4c\xe7\x70\x3e\x3d\x26\xb4\xec\xfe\x0d\x65\xea\xaa\x34\xb6\x55\xb6\x14\xd8\x30\x7a\x12\x17\xc7\x20\x90\x7c\xc3\x91\x36\x4e\x16\x3b\x44\xad\xe4\x2d\xaa\xb2\x96\xa5\xca\xd1\x4a\x87\xf4\x64\x21\xf2\x35\x94\x4f\x80\xeb\x55\xb2\x0e\x0d\xb1\xa4\xe8\x20\x75\x84\x77\x5d\x8e\x4d\x56\xcc\xa5\x65\xeb\x67\x85\x3a\xbd\xdc\x9c\x60\xcb\x34\xba\x8b\xbe\x02\x5b\xb6\x9a\x20\xc5\xa4\x9e\x8b\xaf\xda\x84\xb1\xd0\x51\x47\x05\x57\x0b\x82\x7e\x26\xd4\xac\xbc\xd7\x19\xad\x89\xdd\x35\x62\xc5\xa6\x43\x72\xd5\xbc\x24\x10\x78\x66\x9e\x71\x77\xc0\xf3\xd3\x8a\xe5\x95\xe7\xa6\x15\x54\xe3\x32\x66\xe5\xcd\xb1\x7a\x37\xb0\x95\x6b\x26\x72\x22\x18\x1b\x57\xfb\x08\xc3\x96\xae\xcb\x66\x42\xef\x06\xae\xdf\x36\xcc\xd5\xdd\xeb\x28\x84\x68\x2f\x72\x0e\xf2\x5f\x59\x76\xcd\x7a\xae\xd9\x36\x9c\x3a\x0c\xc3\xd0\xdc\x4e\x39\xd8\xf1\x26\xc8\x5e\x48\xba\x96\x67\xef\xb3\x4c\x7b\xa8\x03\xeb\x82\xb0\x21\x0a\x9b\xc3\xd0\xae\x5c\x7f\x36\x28\x7b\x57\x8b\x84\x19\x7e\x6b\xa1\xe6\xf5\xac\x5a\x2b\x55\x2f\x73\xe5\x86\xa7\x11\xee\x32\xa9\xbb\x74\x5d\xe7\x6a\x7b\xf8\x67\xdc\x73\x08\xf0\xdf\x21\xaf\x07\x1d\x51\xe2\xd3\xdd\x50\x2e\x52\xdc\x3d\xce\x65\x23\x9e\x9e\x74\xfd\x22\x42\x6a\xc1\x27\x1e\x10\x9c\xf4\x4d\x59\x2b\x41\x7c\x0b\xf8\x47\x81\x5a\x28\x8f\x53\x97\x35\xd1\x09\x6c\x69\xfc\x33\xd2\xa5\x36\x94\x4f\x3f\x15\x0b\x1d\xf3\xeb\xf4\x4b\xf1\x79\xb0\xf0\x1a\xf9\x05\xba\x9b\xf3\x80\x9f\x3d\x1b\xfa\xc9\xe7\x23\x7a\x8f\xe9\xc1\xfe\xfd\x17\x08\xdf\x56\x18\x07\x67\x8d\xd7\xbf\x2c\xa9\xdc\xcf\x4b\x66\xf7\xed\xf4\xfe\x3f\xbf\x3f\xfe\x85\x3d\x73\x08\x76\xd5\xda\x13\xfb\x51\x80\xbc\x37\xb6\xc7\xfc\xad\xa3\x25\x9b\x20\x78\x5c\xdc\xfb\xfe\x29\x89\xca\xca\xa8\x8e\x97\x14\xcb\x68\xfc\x17\xef\x9f\x01\x00\x00\xff\xff\xbd\x5c\xd3\x44\x50\x15\x00\x00")

func templateTxtBytes() ([]byte, error) {
	return bindataRead(
		_templateTxt,
		"template.txt",
	)
}

func templateTxt() (*asset, error) {
	bytes, err := templateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.txt", size: 5456, mode: os.FileMode(420), modTime: time.Unix(1467879209, 0)}
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
	"template.txt": templateTxt,
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
	"template.txt": &bintree{templateTxt, map[string]*bintree{}},
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

