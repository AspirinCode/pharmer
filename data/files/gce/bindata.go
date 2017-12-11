// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package gce

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x98\x41\x73\xe2\x36\x14\xc7\xef\x7c\x0a\x8d\x4f\xed\x0c\xb8\x31\xb0\x84\xc9\x2d\xf1\xd2\x6d\x76\xa6\x4d\x26\x90\xf6\xd0\xd9\x61\x84\xfc\xb0\xb5\x91\x25\xaf\x24\x93\xd2\x4c\xbe\xfb\x8e\x8c\x63\x8c\x6c\x63\x91\x43\x02\xef\xfd\x25\xfd\xde\xd3\x1f\x47\xe2\x6d\x80\x90\xc7\x71\x0a\xde\x0d\xf2\x62\x02\xde\xd0\x04\x80\xef\x94\x77\x83\xfe\x1d\x20\x84\x90\x17\xc1\xae\x08\x23\xe4\xfd\xc0\x1f\xaf\x32\x29\x22\x6f\x80\xd0\xb7\x62\x80\x84\x98\x0a\x7e\x1c\xf3\x56\xfc\xae\x12\x66\xf2\x05\x56\x1a\x24\x47\xcf\xcb\x72\x0a\x84\x3c\x26\x08\xd6\x65\xfe\x0e\xe4\x0b\x30\xd8\xa3\x50\xe4\x5c\xef\x87\x68\x29\x72\x9d\xa0\x10\x4b\xc1\x28\xc7\xc7\x41\xff\x0b\x0e\xc7\x95\x8a\x50\xae\x46\x80\x95\x0e\x46\x9b\x4a\x56\x8f\x92\xd6\x68\x81\x6f\x7e\xbe\x15\x7f\xdf\x87\x5d\xe4\xff\xc0\x79\xf2\x55\x02\xe8\x33\x66\x0c\xd4\x10\x3d\x3c\xf5\x81\xbe\x82\x59\x1c\x5b\x48\x87\xe8\xc6\x15\x29\x04\xae\x25\x66\x9d\x48\xa6\x87\x84\x32\x74\xc7\xf2\xed\x56\x0d\xd1\xbd\x78\xed\xed\x20\x39\xcc\xd9\x64\xab\x12\x76\x77\xab\x84\xdd\xe0\x2a\xb1\xbd\xb4\xc7\x8b\x5c\x8a\x0c\xda\x8b\x5a\x6a\x1f\x7d\x49\xa8\x62\x98\xf2\x21\xba\x03\x16\xd3\x3c\x3d\x5b\x15\x14\xb3\x55\xcd\x1d\x76\x64\x48\x67\xc6\xd9\x23\xb7\x8a\x62\xf4\xcb\x0a\xd3\x57\xcc\x7f\xed\xd8\x93\x04\xf3\x38\xc9\x71\x65\xf0\x83\xfa\x6c\x01\x58\x51\x5c\xda\xf5\x64\x57\x6a\xf1\x4d\x47\x9c\x5c\x86\xfe\x15\x67\x9d\xe4\x2b\xf1\xb2\x17\x43\x54\x48\xfa\x71\xb9\x90\x3a\xe9\x62\xae\x25\x9b\xe0\xb5\xa4\x4d\x5f\x3d\x69\x28\x57\x1a\x73\x02\xab\x7d\x06\x2d\xcf\x1b\xf5\x92\x17\x4f\xb2\x60\xa4\x52\xcc\xd8\x91\x36\x02\x45\x24\xcd\x3e\x4a\x6a\x0a\x08\xd6\x10\x0b\xb9\x2f\x9c\x66\xa5\x32\x33\x69\x50\xbd\x97\x38\x35\xef\xfd\xeb\xd6\xce\x96\x08\x3c\x18\x19\xd2\x08\xcb\x68\x14\x74\x72\x74\xa8\x4e\x60\xca\x7c\x1f\xcf\xc4\xbf\xfe\xe4\x0a\x34\x76\x02\x1a\x5f\x06\x34\xb6\x80\xae\x7d\x67\x9e\xa9\x13\xcf\xf4\x32\x9e\xa9\xbd\x61\xce\x38\x73\x27\x9c\xf9\x65\x38\x73\x7b\xbf\xae\x9c\xed\x33\x73\xf3\xcf\xec\x42\x03\xcd\x2c\xa2\x99\x33\xd1\x64\xec\x86\x74\xa2\x73\x61\x9a\xd8\x26\x0a\xc6\x7d\x50\x09\x8d\x93\x14\xd2\xf3\x9e\x6e\x11\xd5\x71\xfe\xa0\x71\x82\xfe\x84\xd4\xbc\xed\x71\x75\x30\x71\xe4\x39\xeb\xe9\x16\x91\x23\x8f\xed\xea\xf1\xcc\x91\xe7\xac\xa9\x5b\x44\x8e\x3c\xb6\xad\x3f\x8d\x1d\x79\xce\xbb\xba\x4d\xe5\x48\xd4\xf0\x75\x70\x35\x75\x64\xea\xf1\x75\xab\xcc\x91\xaa\xe1\xec\xf1\xd5\xdc\x81\x8a\x64\x79\xbf\xb3\x2d\x51\x83\x28\x7c\x7c\xee\xb5\xb5\xef\x4a\xd3\xeb\x6b\x4b\xe4\x42\x63\x9b\x7a\xe2\xbb\xb8\xda\x2c\xd4\xeb\x6a\x4b\xe4\x42\x63\x5b\xfa\xda\x77\xf1\xb4\x59\xa8\xdf\xd3\xb6\xca\x85\xa7\x69\xe8\xa9\xef\xe2\x68\xb3\x96\x83\xa3\x1b\x32\x17\xa6\xa6\x9d\xe7\x1f\x0e\xaa\xce\x6a\x44\x42\x04\x5c\x53\xcc\x5a\x4e\x6a\x99\x14\x3b\x1a\x81\x34\x8b\x7c\x11\x22\x66\x10\x32\x91\xd7\xfe\x21\x44\x54\x65\x0c\xef\x7f\x17\x32\xc5\xda\xa8\xbe\x2b\x51\x3b\x7e\x62\xce\x85\x2e\x8e\xa8\x66\xf2\xb7\xe3\x51\x32\x4b\xb0\x4c\x41\xfa\x38\xcb\x14\x11\x11\xf8\x44\xa4\xbf\x11\x96\x9b\x0b\xc6\xe8\x88\x64\x66\xac\x9f\x40\x5b\x87\x45\x5c\x5d\x3a\x44\x69\x21\x71\x0c\xf6\xb0\x72\xd4\x7b\xc5\xbf\xa5\xc0\xa2\xd3\xf3\xf3\xb1\x88\xc3\x1d\x9c\x08\xbe\xa5\x71\xd1\xa0\x70\xb1\x7e\x7c\x7a\xf8\xba\x08\x57\xeb\xfb\xcf\x35\x06\x33\x91\x90\x69\x79\x7b\x5f\x67\x52\x7c\x07\xa2\xd7\x34\x3a\xd5\x14\xad\xbb\x29\x9a\x6e\xf2\xf6\x14\x0c\x6f\x80\x1d\x37\x02\x15\x3b\x81\x1e\x0f\x62\x64\xab\x29\xcf\xf2\x62\x43\x34\xfc\xa7\xbd\x2a\xf3\x3e\x74\x2c\x64\xb9\x78\xfa\xfb\x3e\x5c\xac\x6f\xc3\xf0\xe1\xf9\xaf\x55\x77\x35\x0a\xe4\x8e\x12\x58\x63\x42\xcc\x95\xa9\xbd\xa4\x52\x74\xdb\xa6\x69\xaf\x6b\x79\x18\x81\x5a\x87\x9c\x14\x87\x25\xe0\x5a\x81\xe5\x2b\xfb\x4a\xf2\x92\x6f\x40\x72\xd0\x6d\xf7\x91\x1d\x48\x55\x7e\xda\x02\x7f\xee\x5f\x75\x7e\x16\xad\x6c\xf9\x0d\x4c\xcd\xd6\x11\xec\xbc\x1b\xa4\x65\x0e\x35\x07\xfe\xc0\xcd\x58\xf1\xcd\xcc\x21\x3a\xa8\x83\x17\xc0\x83\xf7\x9f\x01\x00\x00\xff\xff\x48\x7a\x29\xa3\xf2\x11\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 4594, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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
