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

var _templateTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x5f\x6f\xdb\x36\x10\x7f\x96\x3e\xc5\x55\x4b\x02\x79\x53\x94\xf6\x35\x43\x1e\xb2\x26\xc5\x0c\xb4\xe9\x1a\x67\xdb\x43\x10\x04\x8a\x4c\x3b\xda\x64\xc9\x25\x69\xd7\x81\xa6\xef\xbe\xbb\x23\x45\xd3\x8e\xbd\xa4\x5b\xb3\xac\xc0\x02\x04\x26\x8f\xe2\xf1\x77\xff\x7e\x3c\x1e\x1c\xc0\xc9\x7b\x38\x7b\x7f\x01\xa7\x27\xfd\x0b\xb8\xf8\xb1\x3f\x80\x37\xfd\xb7\xa7\xf0\x22\xc4\xa5\xbe\x86\x42\xc1\x58\x54\x42\x66\x5a\x0c\xe1\xe6\x0e\xa6\x42\xaa\x42\x69\xd0\x75\x5d\x26\xa0\xea\x99\xcc\x05\x8c\x64\x3d\x81\xa6\x49\x07\x3c\x6d\xdb\x34\x9c\x66\xf9\xef\xd9\x58\x90\xf0\x27\x33\x6c\xdb\x30\x2c\x26\xd3\x5a\x6a\x88\xc3\x20\xba\xb9\xd3\x42\x45\x38\x10\x55\x5e\x0f\x8b\x6a\x7c\xf0\x9b\xaa\x2b\x12\x8c\x26\x9a\x7e\xc6\x85\xbe\x9d\xdd\xa4\x79\x3d\x39\x18\xd7\xba\x98\xaa\x83\xb2\x1e\xd3\x82\xd2\x32\xaf\xab\xb9\x1d\xe2\x4e\x52\xd3\x34\xfb\x20\xb3\x0a\x4f\x4c\xfb\x7c\x88\xc2\xf3\x50\x9a\xb6\x2d\x2d\x89\x6a\x08\x28\xe8\x85\xa1\xbe\x9b\x12\x2a\x48\xcf\xb2\x89\x40\x19\xa0\x8e\x59\xae\x1b\x44\xe7\xe9\x78\x27\xf4\x6d\x3d\x54\xb4\x27\x1c\xcd\xaa\x1c\xe2\x6f\x71\xcf\x4e\xb7\xa9\xe7\x6b\x88\x9b\xc6\xec\xda\x29\x92\x9d\xb9\x86\xc3\x23\x40\x9b\x65\x36\x51\x74\x76\x31\x42\x39\xfc\x01\x95\x80\x97\x6d\x9b\xe0\x46\xc4\x42\x0b\xf8\x69\xfa\x4b\x26\x11\x81\x19\x0f\xca\x82\x5c\x67\x27\x9a\x56\xcc\xd8\xf9\xaf\x69\x32\x34\xc3\x13\x41\x94\x46\xdd\x57\x17\x68\x17\x8d\xad\xa9\x3d\xd8\x04\xeb\x5c\xe8\x99\xac\x9e\x17\x57\x13\x06\xf3\x4c\x82\x90\xfc\x5f\xcb\x30\xf8\x38\x13\xf2\x2e\x81\x4c\x8e\x15\xa1\xe4\xcc\x48\xcf\xc4\xa7\x1f\x66\xa3\x91\x90\xf1\xe5\x15\x49\x9a\xb6\x97\xc0\xe5\x55\x51\x69\x21\x47\x59\x8e\x73\x0c\x19\x6b\x52\x7a\xa2\xc1\x64\x82\x1f\xc2\x37\x32\x1b\x4f\x44\xc5\x89\xc0\xe6\xa6\xaf\x6b\x93\x05\x01\x4e\x30\x33\x68\x4a\x56\x86\x36\x41\x68\x85\x75\x1d\x41\x44\xb9\x8c\x43\x5c\x8e\x3c\x9d\xe8\x49\xe8\x5c\x79\x8c\x68\x4d\x8a\xed\x03\xea\x13\x1f\x21\x1d\x7c\x78\xdb\xaf\x00\xd3\x89\x72\x3d\x08\x46\xb5\x84\xeb\x04\xe6\xf4\xb9\xd9\x4f\x5a\xf3\x7a\xea\x5c\x66\x66\xbe\xb7\xac\xe3\x71\x77\xc0\xee\x38\x82\x6c\x3a\x45\x6c\x31\xcd\x28\x48\xec\xe6\xe3\xb2\xc8\x94\xef\x77\x16\x40\x14\xa3\x9e\xf9\x3d\x69\x0f\xa5\x3d\xd4\x48\x98\xac\x7d\xb6\x6e\x30\x1f\xa6\x25\xfa\x32\x26\x71\x02\xd1\xce\xee\x30\x4a\xe8\x6c\x6f\x5d\x64\x3a\x8e\x12\x5e\x81\x52\x54\xf1\x5f\xdb\x60\x0c\xe8\xf5\x2e\x5f\x1d\x5e\x25\xb0\xff\x0a\x0b\x8e\x3d\x24\x4a\x25\xc8\x4d\xe9\xbb\x4c\xaa\xdb\xac\xe4\x30\x04\xdb\x74\x5d\x7b\xba\x12\xce\x14\x74\x21\xd1\x43\xb7\xfd\x71\x28\xf0\x04\x8a\x0c\x6e\x7f\x71\x04\x55\x51\x1a\xbf\x22\x8b\xa4\xa7\x94\x78\xa3\x38\x9a\x58\x75\xbb\xdf\xcc\x7b\x26\x1b\x0f\x61\x57\x45\xc9\x63\xd4\x33\xb0\xce\xaf\x9b\x83\xf5\xb0\x75\x2b\xfe\x69\xb7\x6b\x7a\x30\xec\xb6\x3c\x1f\x91\x5d\xdb\xf2\xc3\xa3\x49\x6f\x68\x8a\x33\xfd\x55\x16\x5a\x0c\x38\x2b\x38\x59\x7a\x2e\xf1\x6d\x19\x85\x41\x6b\x19\xc4\x2b\x27\x53\x9f\xc3\xba\x2c\xf1\x67\xa5\x76\x43\xae\x8e\x62\x59\x1a\x6c\x36\xc5\xc7\x7e\xed\x3c\x60\xe6\x09\x14\xdf\x61\x32\x05\xcb\x12\xc5\x5b\x22\x1d\x4c\x11\x90\x1e\xc5\x06\xa3\x85\x87\x34\x61\xf6\xa4\x69\x4a\xee\xa5\x78\x9f\x88\x9b\x99\x03\xbe\x14\xd0\xa1\xe6\xab\xce\x18\x62\x2b\x24\x45\xac\xe6\x28\x1b\x0e\x23\xf0\x2a\x3c\xfd\xb9\xb2\xe9\xc2\x75\x4f\x96\x79\x0c\xeb\x78\xc1\xff\x8a\x69\xa7\xa2\xf5\x35\x8e\x45\xdb\x5d\xb8\x30\x21\x16\x2e\x42\xe8\xcc\x3c\xd3\x6b\xf1\xe2\x4d\x60\x68\xd0\xc5\xc9\x20\x63\x60\x83\x3c\xab\xf8\xaa\x0a\x86\x42\x31\x8a\x55\x9e\xdc\x88\x93\x37\x6d\x85\xb8\xf7\xf7\x30\x22\x04\x2a\xb7\x23\xca\xfd\x8b\x05\xf6\x01\x1f\x28\x32\xe7\xf5\x27\xcb\x30\x9d\xc3\xf9\xf4\x98\xd0\xb2\xfb\x37\x94\xa9\xab\xd2\xd8\x56\xd9\x52\x60\xc3\xe8\x49\x5c\x1c\x83\x40\xf2\x0d\x47\xda\x38\x59\xec\x14\xb5\x92\xb7\xa8\xca\x5a\x96\x2a\x47\x2b\x1d\xd2\xd3\x85\xc8\xd7\x50\x3e\x01\xae\x97\xc9\x3a\x34\xc4\x92\xa2\x83\xd4\x31\xde\x75\x39\xf6\x58\x31\x97\x96\xad\x9f\x15\xea\xf4\x72\x73\x82\x1d\xd3\xe8\x2e\xfa\x0a\x6c\xd9\x6a\x82\x14\x93\x7a\x2e\xbe\x6a\x13\xc6\x42\x47\x1d\x15\x5c\x2f\x08\xfa\xb9\x50\xb3\xf2\x5e\x67\xb4\x26\x76\xd7\x88\x15\x9b\x0e\xc9\x55\xf3\x92\x40\xe0\x99\x79\xc6\xdd\x01\xcf\x4f\x2b\x96\x57\x9e\x9b\x56\x50\x8d\xcb\x98\x95\x37\xc7\xea\xdd\xc0\x56\xae\x99\xc8\x89\x60\x6c\x5c\xed\x23\x0c\x5b\xba\x2e\x9b\x09\xbd\x9b\xb8\x7e\xdb\x30\x57\x77\xaf\xa3\x10\xa2\xbd\xc8\x39\xc8\x7f\x64\xd9\x6f\xd6\x73\xcd\xb6\xe1\xd4\x61\x18\x86\xe6\x76\xca\xc1\x8e\x37\x41\xf6\x42\xd2\xb5\x3c\x7b\x9f\x65\xda\x43\x1d\x58\x17\x84\x0d\x51\xd8\x1c\x86\x76\xe5\xfa\xb3\x41\xd9\xbb\x5e\x24\xcc\xf0\x5b\x0b\x35\xaf\x67\xd5\x5a\xa9\x7a\x99\x2b\x37\x3c\x8d\x70\x97\x49\xdd\xa5\xeb\x3a\x57\xdb\xc3\x3f\xe3\x9e\x43\x80\xff\x0e\x79\x3d\xe8\x88\x12\x5f\xee\x86\x72\x91\xe2\xee\x71\x2e\x1b\xf1\xf4\xa4\xeb\x17\x11\x52\x0b\x3e\xf1\x80\xe0\xa4\xaf\xcb\x5a\x09\xe2\x5b\xc0\x3f\x0a\xd4\x42\x79\x9c\xba\xac\x89\x4e\x60\x4b\xe3\x9f\x91\x2e\xb5\xa1\x7c\xfa\x99\x58\xe8\x98\x5f\xa7\x5f\x8a\xcf\x83\x85\xd7\xc8\x2f\xd0\xdd\x9c\x07\xfc\xec\xd9\xd0\x4f\x3e\x1f\xd1\x7b\x4c\x0f\xf6\xef\xbf\x40\xf8\xb6\xc2\x38\x38\x6b\xbc\xfe\x65\x49\xe5\x7e\x5e\x32\xbb\x6f\xa7\xf7\xff\xf9\xfd\xf1\x2f\xec\x99\x43\xb0\xab\xd6\x9e\xd8\x8f\x02\xe4\xbd\xb1\x3d\xe6\x6f\x1d\x2d\xd9\x04\xc1\xe3\xe2\xde\xf7\x4f\x49\x54\x56\x46\x75\xbc\xa4\x58\x46\xe3\xbf\x78\xff\x0c\x00\x00\xff\xff\xfd\x53\x8b\xbb\x4f\x15\x00\x00")

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

	info := bindataFileInfo{name: "template.txt", size: 5455, mode: os.FileMode(420), modTime: time.Unix(1466996236, 0)}
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

