// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package azure

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9d\x5d\x73\xdb\x36\x16\x86\xef\xf3\x2b\x38\xba\xb6\xb3\x16\xf5\x61\x26\x77\x8e\xed\x38\xde\x74\x9d\x6e\xe4\xa4\xb3\xbb\xd3\xf1\x30\x12\xe3\xb0\x96\x49\x95\xa4\x9c\x26\x9d\xfc\xf7\x1d\x4a\xb2\x2c\xe2\x43\x7c\x41\x8b\x02\xe5\xbc\x37\x6d\x86\x3c\x16\x08\x9c\x07\xe7\x1c\x00\x07\xc0\xdf\xcf\x1c\xa7\x15\xf9\xb7\x41\xeb\xa5\xd3\xf2\xbf\x4f\x93\xa0\xb5\x97\x3f\x0a\xa2\xbb\xb4\xf5\xd2\xf9\xdf\x33\xc7\x71\x9c\xd6\x28\xb8\x9b\x3d\x76\x9c\xd6\x9f\xfe\xfd\xbf\x26\x49\x3c\x6a\x3d\x73\x9c\xdf\x67\x7f\x90\x04\xd7\x61\x1c\x3d\xfc\xcd\xdf\xb3\xff\x3a\x4e\x6b\x1c\x0f\xfd\x2c\x8c\xa3\xbc\x80\x8b\xe0\xab\x33\x88\xa7\xd9\x17\xe7\x37\x7f\x1c\xa4\x8b\x5f\x5a\xfe\x75\x2e\x72\x34\x4d\xb3\xc4\x1f\x87\xbe\x73\xea\xa7\xd9\x83\xc4\xf7\x38\x0a\x1e\x7e\x7d\xf6\xc8\xbf\x17\x0d\x72\xc9\xc5\xf3\xdf\x67\xff\xff\xb1\xa7\xff\x88\x8f\xe1\x30\x8b\x93\xd0\x5f\x5f\xfa\xec\x33\x03\xf8\x13\xd2\xa5\x38\xfc\x1d\x03\x3f\x76\x7e\xf5\xa7\xe3\xd8\x19\x64\x7e\x16\xa8\x3e\xe7\x55\xe2\x7f\x0f\xc7\xf3\x6f\x59\xfb\x1d\x9f\x66\x82\xb3\x8f\xc0\x3f\xe0\x32\x4e\xe2\x28\x8b\x55\x05\x1f\xfb\x91\x3f\xf2\x9d\xe3\x20\xca\xab\xb7\xb6\xe8\xe1\x4c\x74\xb8\x90\x84\x0b\xff\xf7\x34\xf8\x14\x0c\x9d\xe3\x30\xfb\xb6\xe6\x03\x4a\x19\x98\x97\x6e\xd6\xf0\xca\xf2\xe6\xdf\xef\x9c\x47\xa3\x55\x32\x54\x25\xce\x25\xc3\x99\x20\x5c\xe6\x79\xfc\x55\x09\xdc\x7d\xb9\x1f\x06\x48\xa1\xd3\x14\x2f\xf1\x4d\x1c\x5d\x3b\x6f\xe3\xe8\x5a\x55\x6c\xde\xae\xce\x51\x5a\x52\xd5\xbc\x59\xfd\xd4\xa4\x9a\x1f\xc3\xe4\x3a\x8c\xd4\x7d\x6b\x56\x66\x49\x3d\xf3\x12\x4d\x2a\x09\x94\xe7\xb8\x40\x89\xae\x49\xaf\xb9\xf9\x16\xef\x39\x03\x3f\xcc\xfc\x5b\x65\xc1\xff\xf4\x27\x7e\x54\x8e\xee\x1f\xb9\x98\x19\xb9\xef\x52\xff\x66\x4d\x91\xbf\x05\x48\x91\x5f\x83\xc7\x76\x96\xb7\x71\x12\x60\xc6\xe1\x26\x97\x34\xb6\x0d\xfa\x32\xcb\x2d\xe1\xac\x44\x43\x43\x78\x3e\x1e\x87\x51\x1c\x2a\xfd\xd1\x45\x9c\x64\x5f\x1c\xb0\x9b\x46\xb9\x70\x85\xbe\x7a\x9e\x04\x63\x3f\x1a\xe9\xcb\x3f\x9d\x26\xf1\x24\x28\x2f\x3b\x98\xcb\xe1\x30\x07\x7f\xf9\xca\x6a\xcf\xbd\x34\x58\xed\x59\x73\x57\xa8\xb6\xbe\xe0\x72\x33\x3c\x2b\xd3\xd0\x08\x0f\xc2\xe8\xda\x9f\xc4\x89\xd2\xd7\x2e\x1d\x7e\xb9\x5d\x5c\x3a\x7b\x33\xe3\xf8\x4b\x1c\x8d\xe2\x48\x55\xf6\x87\xb7\x00\xd9\xd3\x1b\x43\xac\x8f\xfd\x64\x14\x7e\xfe\xac\x29\xb0\xd4\x58\x4c\x6f\xcc\x2c\x45\xfe\x83\x4a\x62\x1e\x8a\xd5\x8a\x28\x8a\xcf\x0b\xaf\xc0\xd4\x45\x90\x7d\x09\x92\xbc\x37\x29\xb9\x9e\x7d\x01\xd0\x9b\xf2\xd2\x4d\x3b\x93\xb6\xbc\x72\x9a\xf3\xe2\x0c\x61\x3e\xf6\xc7\xe1\xe7\x38\xd1\x38\xbf\x59\xb9\x40\x0b\x9b\x34\xed\xe2\x37\x57\x1d\xaa\x54\x60\x89\xb7\x9d\x17\x29\x7a\xdb\xe5\x20\x22\x8c\xd2\xcc\x8f\x86\xc1\xe5\xb7\x49\xa0\x18\x4a\xa4\x37\xd3\x59\x54\xec\xa7\xe1\xf0\xea\xe8\xe0\xa1\xa4\x51\x90\x0e\x93\x70\x72\xff\xa1\xb2\xc0\x70\x92\xff\x65\xfb\xe1\xbb\xfd\xdb\xd6\x4b\xe7\xe0\xf9\x61\xef\xe1\x37\xc2\xf4\xa6\xf5\xd2\x71\x0f\xf0\x21\xc7\x9e\xe2\x45\x2a\x8d\x1b\x84\xf0\x7c\xe5\x71\x31\x74\x96\x5e\x88\xbf\x52\x88\x3c\xe5\xe7\xd3\x74\xf5\xe1\xd2\x3c\x09\xcf\x64\xa9\x5c\x1f\x7b\xaa\xb0\x44\x7c\xf8\x35\x28\x3e\x2c\x38\x77\xf1\xb9\x54\x59\xc1\x37\x8a\x6f\x82\x62\x9f\x94\xbd\x8a\xf8\x46\x55\xc1\x15\xa7\xb0\x27\x1b\xce\x3d\xc9\xb6\xed\xe9\xcc\xcd\x9e\xd2\x12\xec\xa9\x3a\xec\x9e\xd4\xa1\xa4\x27\x25\xd1\x65\x11\xeb\x76\x19\xd6\xed\x32\xac\xdb\x0a\xac\xbb\xc4\x9a\x58\xdb\xc4\xda\x2d\xc3\xda\x15\xb1\x76\x05\xac\x3b\xcf\x25\xaa\xfb\xa4\x9a\x54\xdb\xa4\xba\x53\x46\x75\x47\xa4\xba\x2b\x50\x7d\x28\x32\xdd\x66\x04\x42\xa8\xad\x42\xdd\x2d\x83\xba\x2b\x42\xed\x89\x11\x48\x57\x0a\xab\x19\x80\x90\x6a\x3b\x54\x0f\x32\x3f\x1a\xf9\xc9\x68\xdd\x88\x51\x29\x33\xf4\xb3\xe0\x3a\x4e\xbe\xcd\x16\x86\xf6\xd3\x20\x09\x57\x97\xac\x38\xa4\x24\xfa\xbb\x83\xbe\x7e\x54\xa9\x94\xa9\x82\xbe\x6a\xd8\x79\x48\xf4\x89\xbe\x6d\xf4\x11\xb3\xdf\x2e\xb1\xfb\xce\xbe\x33\x8c\x6f\x27\xd3\x2c\xd8\x0f\xa3\x2c\x88\xd2\xf0\x2e\x70\xee\xa7\x2b\xa5\xae\x21\xc6\x43\xbd\xbe\xd8\x31\x3a\x9e\xbb\xae\x67\xc8\x6c\x29\x41\xda\x20\x05\x7a\xc5\x94\x4d\x11\xcb\x0d\x0e\x19\x9b\x12\x6b\x63\xd6\xe0\xed\xbe\x68\x8c\xda\xee\x4f\xd5\xe4\x57\x77\xfa\x09\x16\x9d\x98\xda\x8c\x4b\xed\xd6\xa6\x0d\xa7\x0d\xb7\x6c\xc3\x11\xb8\x5d\xb3\xf0\x05\x98\x5e\x6c\x77\x7a\x44\x9f\xe8\x5b\x46\x1f\x33\xed\xae\xc2\xb4\x8b\x88\xcb\xb3\x32\x34\xed\xe4\xdb\x36\xdf\xb7\x20\xe0\xb7\x00\xe1\x6d\x29\xd0\x26\xe2\x44\xdc\x36\xe2\xfa\x55\x22\xa5\x0c\x12\xbd\x94\x2e\x23\xb9\x1e\x63\x17\x82\x6f\x19\x7c\xfd\x4a\x92\x52\x06\x01\xbf\x7c\xa9\xa9\x7f\x40\xf2\x49\xbe\x6d\xf2\xb1\xa0\xa6\xab\x88\x69\x44\xd3\xee\x89\x80\x73\x29\x95\x7c\x5b\xe7\x1b\x8c\xda\xbb\xaa\xa8\x5d\x24\xbc\x23\xcd\x39\x12\x71\x22\x6e\x1b\xf1\x1e\x80\x77\xcf\x2c\x78\x91\x86\xab\x52\xf0\xc2\x29\x47\x92\x6f\x9d\xfc\x3e\x40\x7e\xdf\x8c\x7c\xd1\xe4\xbb\x52\x54\xc3\x01\x2b\xc9\xb7\x4e\xfe\x21\x40\xfe\xa1\x19\xf9\xe5\xb9\x00\x1c\xb0\x92\x7c\xeb\xe4\x7b\x00\xf9\xde\x7a\xf2\x99\x23\x63\xd0\xdc\xd8\xf0\xc9\x53\x8c\x9e\xa4\x19\x30\xa9\xe1\x3c\x8e\x9e\x68\x4f\x6c\xdb\x13\x70\x82\xc0\x53\x4d\x10\x88\x88\xf7\xa5\x71\x12\x11\x27\xe2\xb6\x11\x7f\x01\xe0\xfd\x62\x93\x2e\xf3\x27\xcf\x72\x7c\xd5\xbe\x4d\xcb\x9b\xbc\x28\x05\xe6\x38\x76\x69\x4d\x68\x4d\x50\x68\x1f\x6b\x38\x5e\xb5\x21\x8c\xcb\x37\x5c\x88\x14\xaf\xed\xf9\xa4\x98\x14\x3b\x9b\xa4\xd8\x85\xac\xb1\x2b\x5b\x63\x71\x16\x5c\x9a\x0a\x5c\xf1\x73\x04\xb9\xf8\x90\x20\xd7\x00\x32\xc4\x71\x29\xc6\xf2\x18\x85\x14\x93\xe2\xad\x51\xdc\x85\xcc\x71\x57\x36\xc7\xe2\xd2\x8c\x3c\x9f\xd4\x61\x60\x41\x92\xb7\x47\xb2\x07\x91\xec\xc9\x24\x8b\xd3\x46\x72\x5e\x49\x9f\x03\x3d\x92\xbc\x35\x92\x4f\x80\x8d\xb0\x27\x9a\x7d\xb0\x27\xe0\xae\x7b\xc5\xb6\xb5\xde\x36\xe7\x46\x2d\xd1\xb9\x0d\xde\x36\x37\x73\x75\x82\x6c\x89\x3e\xd1\x6d\x89\xd6\xa1\x00\x64\x13\x1d\x10\x85\xe6\xa1\x00\x2d\x8c\x88\x72\x05\x20\xee\xdc\x1d\x41\x82\x2e\x70\x97\x5d\x60\x5d\x2b\x27\x73\xb6\xaf\x7e\x4d\xe2\xdb\x18\xed\x09\xa2\x34\x79\x27\xef\xbb\xc3\x3b\x64\xf0\x75\xd6\xbe\x7a\x4a\x25\xdd\x7f\xf3\xdc\x3f\xb6\x9f\x5f\x94\xc3\xdc\x7f\xd3\x90\xa0\x39\xa4\x39\xd4\xf5\x01\xdc\xfd\xab\xa4\xc9\x3b\x79\xdf\x1d\xde\x81\x03\x00\x0a\x42\x88\xfb\x2f\xcf\x9f\xed\xd2\xfd\x37\xcf\xfd\x77\x40\xf7\xdf\xa9\xe0\xfe\x9b\x86\x04\xcd\x21\xcd\xa1\xae\x0f\xe0\xee\x5f\x25\x4d\xde\xc9\xfb\xee\xf0\x0e\x1c\x83\x52\x10\x82\xd6\x81\x80\x5c\x60\x8f\x01\x40\xf3\x02\x00\xec\x64\x10\x51\x0e\x0b\x00\x9a\x07\x05\x4d\x22\x4d\xa2\xae\x17\xe0\x21\x80\x4a\x9a\xc4\x3b\x24\x7e\x87\x88\xef\x81\x76\xbf\x57\x65\xd9\xf7\x40\xec\x05\x2b\x07\xe8\x2c\xcf\x57\x60\x2f\x60\x2f\xb0\xdd\x0b\xfa\x57\x77\xd0\x64\x58\x51\x4e\x63\xeb\xfb\xdd\xe7\xd2\x1e\x18\x8e\xf8\x88\xf9\x56\x53\xfc\xfa\x29\x8a\x74\x8a\x31\x2d\x12\xdd\x76\xb9\x91\x80\x44\x6f\x91\x68\x30\x50\xa9\x32\x3e\x15\x60\xb7\x9d\xbc\x4a\xd8\x77\x19\xf6\xda\xa2\x14\x04\x7f\xc3\x64\x1d\x31\x51\x4d\xbe\xf6\x95\x73\x75\x8d\x9b\xab\x03\x53\x75\xaa\x64\xea\x34\x8c\x07\x1a\x42\x1a\x42\x4d\x07\x80\x67\xe9\xf4\x69\x3a\x84\x9d\xb0\xef\x04\xec\xc8\x38\xce\x55\x8c\xe2\x4a\x4f\x34\x60\x54\x4b\xbe\xb7\x38\x84\x73\xc1\x39\x09\x57\x35\x25\xc1\xe3\x39\x08\x73\x93\x60\x46\x40\x36\xcc\x9d\x94\x8e\x3c\x50\xdc\x8c\xc7\xe1\x58\xd3\x86\x63\x1d\xd0\x43\x77\x54\x2e\xba\x23\x6d\x15\x93\x93\xc5\xb9\x50\x4c\xcb\xb6\x55\xcb\x86\xfa\x69\x51\x10\x47\xda\xed\xd1\x59\x13\xe9\x2d\x22\x8d\x4d\x99\x55\xc9\x6e\x6f\x9a\xd3\x26\xec\xbb\x0c\x7b\x6d\xd3\x08\x46\xc9\xed\xfa\xdc\x76\xd2\x4e\xda\x9d\xe6\xd3\x8e\xa4\xb6\x1b\x66\xb6\x8b\xbb\x3a\xe4\xa0\x86\x1b\xdb\x9a\x37\x38\x03\xd3\xda\xab\x64\xb5\x37\x0d\x08\x9a\x42\x9a\x42\x4d\x0f\x80\x1d\xbf\x3e\xa3\x9d\xb4\x93\x76\x67\x27\x68\x47\x66\x2e\xba\x8a\x79\x8b\xf2\x03\x67\xb9\x1e\x4c\xc2\xb7\x39\x6b\xd1\x05\xe7\xe1\xba\xaa\x69\x38\x9e\x9f\x4c\x9a\x1b\x45\x33\xb8\xd3\xa8\xca\x46\x23\x29\x5d\x5d\xde\x72\xcf\x45\x14\xd2\x6e\x3d\x3a\xe9\x99\xc4\xe2\x2a\x61\xe2\x4e\xdc\x77\x07\xf7\x3e\x18\x8d\xf7\x55\xe1\x78\x5f\x3a\x45\x4d\xa6\xbc\xdd\x27\xe6\x3f\x2f\xe6\xf5\x85\x2a\x7d\x34\xf2\x16\x05\x71\x74\x7b\x6d\x46\xdf\x24\xb7\xf0\x74\x13\xe4\x7a\x18\xb7\x9e\x82\xda\xf2\x6b\x4a\xb8\xb4\xf7\x13\x33\x6b\x61\xc4\xe8\x81\x56\xd8\x53\x19\x61\xde\xba\x43\x9a\x1b\x45\xf3\x00\xb9\x6c\x65\xa0\xbb\x6c\x65\x50\x7d\xef\xf2\x61\x0d\x98\x5b\x5d\x7e\x6e\xda\x22\xf3\x00\xba\x47\x67\xa0\xbd\x48\x47\xa7\xdb\xf2\x6b\x24\x6a\x39\x83\x81\xba\x2d\x6a\x0d\x9b\xb7\x14\x05\x8b\x1a\x7e\xc4\xd5\x38\x3c\x67\x83\x6e\xca\xf6\x4c\xce\xc0\xec\x6a\x1c\xb5\x38\x81\x27\xf0\xbb\x03\x3c\x66\xf3\xb5\x06\xbf\xf2\x55\x28\xb5\xec\x8c\xa0\x4b\x2f\x6a\x0d\x75\xe9\x6b\xce\xd0\xd0\xbb\x74\x4b\x3a\xa6\x85\xa3\x85\x73\x4c\x7b\x81\x81\x4b\x7f\xcc\x7d\x37\x04\x9e\xc0\xdb\x07\x1e\x99\x4b\x1c\x68\x2f\xbc\xd1\xb9\xf4\xf2\xeb\x1e\xda\xb5\xac\xf5\xd0\xa7\x17\xd5\xb6\x0f\x7b\x75\x51\xd4\xae\x1e\x69\xc6\x68\xc6\x1c\x33\xd2\xc1\x5d\x2d\xb2\x28\x49\x27\xe9\xbb\x44\x3a\xcc\x79\x95\x71\x1a\xbb\x01\xbb\x81\xb3\x2b\xdd\xc0\x60\xa0\xf6\x98\x9b\xc9\x48\x3c\x89\x6f\x00\xf1\xc8\xfe\xed\x81\xf6\x6e\x32\xed\x5a\x39\x70\x2b\x8f\xeb\xd6\x91\x15\xc2\xb1\x5a\x51\x71\x78\x04\x2b\x8a\xda\xd6\x24\x4d\x19\x4d\x99\x63\xc6\xba\x87\xb3\xee\x91\x75\xb2\xbe\x78\xba\x8b\xac\xc3\xa4\x57\x19\xaf\xb1\x23\xb0\x23\x2c\x1f\x35\xbf\x23\x18\x8c\xd8\x1e\x77\x91\x24\x99\x27\xf3\x0d\x60\x1e\xdc\xdf\x2d\x0a\xa2\x79\x92\xbc\x4b\x92\x1d\xe1\xfe\x51\xb3\x3b\x02\xda\x0f\x2a\xc5\x40\x42\x2f\xd8\x52\xba\x3f\xfb\x00\xfb\x80\x63\xd0\x07\xa0\x0e\x60\x9a\x3b\x59\x7e\x5f\x0f\xa7\xee\xea\x9e\xba\x43\x73\x2c\x2a\x25\x4e\xda\x51\x30\x6d\x1b\x6d\x9b\x63\xd8\x05\xf0\xb1\xdd\x63\xae\x1f\x23\xed\xa4\xdd\x36\xed\x50\xce\xa4\x69\xca\x24\x70\x8a\x38\x37\x36\xd6\xed\xca\xd1\xe4\x9a\x4a\xb9\x35\x96\x34\x4c\xeb\x46\xeb\xe6\x18\xf6\x01\xdc\x97\x3f\xea\x5a\x04\xe2\x4e\xdc\x6d\xe3\x0e\xa5\xd5\x98\x66\xd5\x94\x9f\x0c\xce\x2d\x8d\xb5\x3b\x73\x74\xe5\xb5\xd2\xc2\xab\x25\x0d\xd3\xba\xd1\xba\x39\x86\x7d\x00\x77\xe6\x8f\xb9\xea\x80\xb8\x13\x77\xeb\xb8\xa3\xeb\xad\x95\x96\x5b\x81\x23\x86\x99\x2a\xce\x5e\xd0\x8c\x5e\x80\x1b\xfd\x47\x9d\xa9\x4d\xe0\x09\xbc\x6d\xe0\x4f\xdb\x7d\xe8\x4c\x4c\x51\x4e\x97\x4d\xc6\x6b\x9c\x48\xf9\xe2\xc9\xda\xe1\xe8\x06\xb8\xc5\x0e\x73\x95\x04\x71\x72\x79\xa9\x34\xc9\xad\x81\x5c\xec\x7a\xff\x53\xd5\xed\xfe\xd2\x59\x77\xf2\x31\xf0\x34\xb6\x44\xb6\xf0\x74\x23\xc8\x82\xb6\x56\x79\x7f\x7f\x39\xb4\xbc\x38\x8c\xd0\x6e\x1e\xda\x8e\xbb\x0f\x07\x09\x0a\xd9\x39\xbc\x1d\x91\x5e\x5e\xbd\x41\x7e\xef\x9f\xd4\xce\x2f\x78\x61\x81\x42\x94\xf4\x3a\xa4\xd7\x2e\xbd\x28\xba\x15\xb9\xe5\x8d\x74\xe4\xb6\x16\x6e\x61\x9b\x4b\x8b\x4b\x72\xef\x9f\x36\x80\x5c\xec\x72\xc4\x53\xe4\xaa\x72\xf9\x6e\x23\x5e\x55\x4e\x66\xeb\x60\x16\x34\xb6\xd0\x8d\xe4\xbc\x91\x8b\xd0\xce\x9f\xd4\x0b\x6d\xbf\x8b\x4f\x2c\x28\x64\x35\x77\x7a\x76\x65\x7c\x3d\xf2\x4b\x7e\x6b\xe1\x17\x8e\x72\x15\xb2\xe4\xd7\x21\xbf\x76\xf9\x45\xd9\xad\xc8\x2d\xaf\x01\x27\xb8\xf5\x80\x0b\x5b\x5d\x9a\x5c\x92\x7b\xff\xb4\x01\xe4\x62\xd7\x80\x9f\x22\xd7\x80\xf7\xe5\xad\x6c\xb4\xb6\x64\xb6\x06\x66\x41\x63\x0b\xdd\xf6\x2d\x53\xdb\xe6\x0e\x4c\x52\xbb\x71\x6a\x5f\x03\x37\x3f\xbf\xd6\xdc\xfb\xfc\x1a\x3c\xe3\x4b\x11\xef\x92\xe4\x9f\x95\xe4\x86\xe4\xa1\xbf\x6e\xf7\x11\xf0\xfb\x86\xe4\x8b\x59\xbe\x72\xc0\xcc\x24\x5f\xc2\xdf\x00\xf8\x53\x88\xfe\x54\x83\x7f\x5a\x9d\x7f\x8e\x17\x89\x7f\x03\xf0\x87\xf6\x9e\x4a\x82\x30\xe5\x25\xc1\xba\x0c\x8a\xb6\xc9\xd7\x45\x7f\xc6\xf5\x86\x7a\xbd\x71\xa7\x2f\x0b\xf7\xd8\xe3\xd9\xe3\x2d\xf7\x78\xa4\xb3\x6b\x36\x99\xeb\x62\x3d\x31\xf1\x4a\x1a\xb0\x73\x97\x01\xb9\xb7\xce\x3d\x62\xf1\x5d\x53\x8b\x5f\x8a\x3e\x67\xaa\x48\xbe\x75\xf2\xb1\x10\xcf\x55\x45\x78\xa5\x80\xaf\x9f\xc1\xb2\x14\xde\x75\xd0\x3a\x8b\x82\x9a\x54\x62\xc5\xca\xc9\xfa\xc9\x0b\x4b\xf5\x06\x4e\x85\x7b\xad\x39\x14\x4e\xe7\xdc\xc5\x95\x4f\x69\xaf\x36\x87\xb1\x34\x71\xb6\x4d\x5c\x17\x71\xee\x5d\x53\xe7\x5e\x8a\x3e\x67\xef\x89\xbe\x7d\xf4\x31\x4f\xd7\x55\x39\xba\x52\xc2\xd7\x8f\xdc\x2c\x79\xb9\x3e\x5a\x67\x51\x50\x93\xcc\xa3\x38\x80\xa4\x64\xa3\x90\xa5\x8a\x1f\xa2\x61\x8d\x28\x38\xaf\xf8\xa1\x74\x1e\x40\x57\x8a\x6b\x7a\x87\x4d\x8c\x6b\x3c\xa0\xce\x9e\x59\x5c\x23\xa6\x19\xc8\x67\x23\x30\xcd\x80\xe6\xdd\xba\x79\xf7\x90\xc8\xc6\x33\x8d\x6c\xca\xe1\xe7\x94\x1d\xd9\xb7\xcf\x3e\xe6\xed\x3c\x95\xb3\x2b\x47\x7c\xfd\xc0\xd5\x8e\xa7\x3b\x03\x92\x90\xce\x34\x49\x48\x67\xe0\x1c\xa5\x1c\xed\x74\xbc\x6a\x63\x78\xc3\x8e\x8b\xf6\x32\x6d\x43\x2b\x08\xae\x35\x25\xec\x0c\xa0\xef\x4c\xb3\x58\xa2\xd3\x86\x18\x7c\xca\x9b\xd4\x0f\xfb\xd5\xe2\x8e\x27\xaf\x0d\x20\xa5\xf4\x4c\x73\x7f\x91\x4e\x1b\x92\x95\x90\xef\x64\x6e\xf7\x3a\xd5\x86\xf9\x4f\x5e\x1f\xc0\x6c\xe3\x99\x66\xb6\x51\xa7\x0f\x29\xa1\x60\xf5\x46\xec\x7b\x6b\x75\x70\x58\x2d\x38\x79\xf2\x0a\xe9\x01\x0a\xe9\x99\x29\x44\x9a\x0a\xef\x76\xe5\x19\xe0\xd5\x71\x24\x15\xb2\xd2\xd8\x03\xc4\x9d\x0f\x74\xfe\x1c\xbd\x3d\x75\x63\xb7\x1a\x3c\x7d\x7d\x20\x0e\x5d\x77\x9b\xad\x56\x1f\xe5\x2e\xbd\xea\x79\xe3\x4f\x5f\x21\x88\x4f\xd7\x5d\x4a\xa8\x55\x08\xe0\xd5\x57\xfd\x0a\x35\xb2\xda\xd8\x88\x57\xd7\xdd\x2c\xa5\xd5\x08\xe2\xd7\x57\x3d\x0b\x55\x52\x68\xed\x7d\x4c\x29\xab\x62\x6c\xf7\x4d\xb4\x3b\x30\xf1\x2c\x88\xb1\xdd\x1f\xdf\xee\x48\x1c\x3b\xd0\x05\xb2\x3a\x13\x84\x44\xb2\xde\x0b\x06\x4e\x1a\x95\xec\x23\xdb\x84\x44\x39\xb6\xfc\x26\x5a\x1e\x32\x42\x3d\xd9\x08\xb1\xdd\xab\xb7\xfb\x1b\x04\xf7\x37\x32\xeb\xf2\xed\x16\x8a\xd8\xf3\xa0\xe2\xb6\x7b\xdd\xca\x87\x3c\x57\xad\x6c\xdc\x6d\x2c\x3d\x58\xd0\xd3\x2d\xa4\xa8\xdb\x2a\x2e\x9a\x9a\xda\xac\xa6\x12\x4c\x55\x09\x75\x65\x5d\x57\x98\xaa\xca\x35\x45\xfb\x57\xaf\xa6\x80\xe0\xe0\x8d\x14\x19\x88\x13\x24\x8a\x19\x2b\xea\x68\x83\x3a\x42\x5c\x94\x27\x79\x28\x64\x71\x8a\x6a\xda\x98\x9a\x7e\x81\xf6\x82\x17\xa5\x0c\xae\x34\xf3\x0e\x0e\x19\x69\xab\x9a\xbd\x83\x6c\xcd\x2a\x4a\x19\x9c\xf8\xde\xef\x54\xec\x20\x4f\xbd\xd9\x91\x9c\xf9\x82\x10\x7c\xf4\xf3\x21\xe7\xb5\x94\x2d\x8e\xe4\xf2\x15\x84\xe0\x03\xb1\x3a\x1e\x9b\x5c\xd5\xe4\xff\x6a\xbb\xde\x7e\xc7\xbd\x05\x1a\x5e\x21\xba\x30\xee\xae\xa8\x80\xce\xea\x05\x28\x0f\x77\xad\xae\xb5\x33\x8a\xa6\x33\x6c\x8c\x2a\x55\xef\x77\xe1\xaa\x17\x45\x77\xbe\xea\x68\xbd\x9f\x54\xa5\xc1\x3a\x03\x55\x2e\x0c\xd1\x1a\x5b\xe5\xd9\x11\xef\x90\xa6\x25\x49\xdd\x56\x90\xc3\x9e\x54\xef\xb2\xe1\xaa\x95\x7a\x83\x46\x4d\x92\xdc\xf1\x7a\x83\x95\x7e\x42\x35\xc6\x2a\x0c\xd4\x57\xd1\xa3\x1b\x57\xdf\x8b\xe3\x36\x90\xa1\x53\x94\xba\x37\x61\x62\x7d\xe5\xc1\x72\xdf\xab\x36\x14\x90\xc7\xc4\x8a\x86\xd9\xe0\x88\xb8\xfe\x36\xc6\x12\xeb\x65\x49\x4d\x5b\x2b\x66\x63\xfb\xeb\x73\x34\x35\xd9\xf5\x86\x3b\x0f\xaa\xd6\x1d\x48\x3a\x92\x25\x37\x5d\x77\xc3\x0f\x77\x81\x14\x90\xa2\xd4\x22\x59\x50\x34\x04\x8a\x0f\x6e\x77\xbb\xec\x19\xf3\xe6\x03\xa6\xbd\x05\x31\x36\x73\x95\x66\x46\x2d\x90\x24\xaa\x69\x6e\xc5\xd2\x76\xbb\xb3\x3e\x2d\xd9\x9a\x0d\x9a\xd7\x09\x32\x42\x92\xe8\xc6\xab\x6f\xfc\xed\xb8\xe6\x9e\xa0\xe2\x70\xbd\x35\x4b\x6d\x40\x32\x43\x41\x68\x11\x44\x0a\x9f\x2b\x4f\xaf\x76\x18\x52\xcd\x9b\x0e\xed\x15\xca\x63\x34\x81\x35\xd3\xce\xfa\x4d\x48\xd6\xba\x04\x78\x1b\x9b\x24\xb8\xe1\x8a\x9b\x7d\xf5\x49\x1b\x59\x72\x10\xc4\xea\x8e\x7d\x55\xf3\xa5\x9b\x54\xd5\x49\xee\x4b\x90\x5a\x17\xe5\x6a\x37\xdb\x5b\xa8\x37\x58\xed\x27\x55\x6b\x64\x2d\xb3\x28\x55\xb3\x2d\xaa\xbb\xc6\x1f\xa1\xa9\x83\x8f\xdb\x9d\x3a\xc0\x97\xd9\xd1\xb5\x95\x06\xae\xbc\x1b\x6b\x0a\x1a\xcb\x7e\xdc\xf2\x58\x96\xaa\x52\xa9\x0a\x09\x1c\x3f\x6e\x31\x70\xfc\x29\xb5\xf4\xcc\x71\x7e\xcf\x85\x5a\xc3\x24\x18\x05\x51\x16\xfa\xe3\x87\xf6\x5a\x2a\x6f\x92\xc4\x77\xe1\x28\x48\x72\xe5\x1c\x7d\x9f\x26\x41\x6b\xb5\xd1\x27\x63\xff\xdb\xeb\x38\xb9\xf5\xb3\xfc\xfd\xe7\x30\x18\x8f\x1e\xde\xfb\x51\x14\x67\x7e\xae\xd9\xfc\x77\xff\x7e\xf8\x9e\xc9\x17\x3f\xb9\x0d\x92\xe7\xfe\x64\x92\x0e\xe3\x51\xf0\x7c\x18\xdf\xfe\x63\x38\x9e\xa6\x59\x90\xec\x3f\x7c\x4d\xfe\x93\xab\xd5\x50\xfe\xd9\x28\x4a\xc5\x3f\x59\xfc\xc5\x8f\xe5\x87\xcc\xbe\xab\xc8\xc2\xc3\xd7\xe4\x4a\x8d\xee\x86\x71\xf4\x39\xbc\x9e\x55\xf2\xbf\x1f\xde\x9f\x5e\x5d\x9e\x5e\x1c\x5d\x5c\x5e\x9d\x9f\xac\x7c\x40\xfe\x4b\x71\x92\x93\xd7\xf2\xf3\x96\xb8\xca\x82\xc8\x8f\xb2\xab\x70\x54\x14\xfa\x23\x9d\xb3\x3c\x7f\x2d\xfe\xc4\xd8\xff\x14\xcc\xbe\xf3\x72\xf6\xda\x39\x17\xfe\x3a\x8c\x26\xd3\x6c\xfe\xe7\x7f\x65\xad\xe5\x9b\x1f\x7b\xf0\xb7\x0f\x3e\xbc\x1a\x1c\xbf\x3f\xff\xf5\xf2\xfc\xdd\x45\x49\x0d\xd2\xe9\xa7\x65\xf7\xd3\xd6\x63\x55\x48\x5b\x9b\xc1\x8a\x50\x0d\x75\x3a\xfe\xe5\xfc\xb4\x54\x1f\xc3\x71\x18\xac\xd1\xc7\xfc\xb5\xb6\x06\xc7\xb3\xd7\xf5\x7d\xfb\xe0\xf4\xf8\xfd\xe9\x25\xf0\xfd\x69\x30\x4c\x82\x6c\x5d\x1d\x06\x0a\x09\xb1\x1e\x2a\x99\x65\x5d\x26\x7e\x9a\x7e\x8d\x93\xd1\x4a\x7d\xd6\x9a\x70\xc9\x0a\x0c\xb2\x38\xf1\xaf\x6b\x33\x06\xe9\xfc\xe7\xeb\xe8\xd9\x83\xcb\x77\xef\x8f\xce\x4e\xaf\x8e\x8e\x8f\xdf\x7d\xb8\x58\xab\x8f\xc5\x57\x5c\xf9\xc3\x61\x3c\x8d\x34\x1a\x51\xbe\x5c\x2a\x63\xd6\x58\xce\xa2\xb5\x9c\x23\x95\xec\x06\x3a\xfc\xa2\x4a\x6f\x4f\xff\x83\x54\xe7\x26\xf8\xa6\xae\x8a\xf4\x62\x7d\x35\x9c\xb7\xa2\x3c\xce\xd7\xd2\xf9\xdc\x4c\x3f\x05\x49\x14\x64\x81\xc2\xf7\xdc\x05\x49\xba\x88\x0b\xda\xcf\xbd\xe7\x07\x0f\x34\x05\xd1\x9d\x80\xd1\x28\xb8\x6b\xbd\x74\xb2\x64\x1a\xac\x7a\x8c\x24\x1e\xc9\x4f\xff\xf4\x17\xcf\x9e\xad\x7e\x9a\x88\x7c\xa1\xf0\x17\x48\xe1\x9f\xfd\x71\xaa\x28\x5d\x7c\x3c\x2b\x7e\xf6\xb0\x58\x7e\xde\x24\xcf\x7e\xfc\x3f\x00\x00\xff\xff\xe0\x7f\x33\x58\xda\xae\x01\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 110298, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
