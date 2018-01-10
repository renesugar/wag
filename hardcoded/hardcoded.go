// Code generated by go-bindata.
// sources:
// _hardcoded/doer.go
// _hardcoded/middleware.go
// DO NOT EDIT!

package hardcoded

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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x3a\xff\x6f\xdb\xb6\xf2\x3f\xdb\x7f\x05\x67\x60\x9f\x4a\xa9\x23\x27\xdd\xd2\x0d\xde\x27\x7b\x58\xd3\xf4\xb5\xd8\xd6\x16\x75\xf6\x56\xa0\x28\x56\x5a\xa2\x6d\x2d\x12\xa9\x89\x54\x1c\x23\xf3\xff\xfe\xee\x8e\x94\x4c\xd9\xb2\x93\x7d\xc1\xc3\x7b\x2e\x50\xcb\xe2\xdd\xf1\xee\x78\xdf\x99\x82\xc7\xd7\x7c\x2e\x58\x9c\xa5\x42\x9a\x7e\x3f\xcd\x0b\x55\x1a\x16\xf4\x7b\x83\xe9\xca\x08\x3d\x80\x87\x58\x49\x23\x6e\x0d\x3e\x0a\x19\xab\x24\x95\xf3\xd1\xaf\x5a\x49\x7a\x51\x96\xaa\x24\xa8\x59\x4e\x10\xa9\x1a\xa5\xaa\x32\x69\x86\x3f\x32\x35\xc7\xaf\x9c\x9b\xc5\xa8\xe4\x32\xc1\x1f\x52\x18\xf7\x35\x5a\x18\x53\xe0\xb3\x5e\xc9\x18\xbf\x4d\x9a\x8b\x41\x1f\x1e\xe6\xa9\x59\x54\xd3\x28\x56\xf9\x88\xcf\xc4\xed\x68\xb1\xd2\xa6\x4c\x6f\x8f\xe7\xaa\x7e\x1c\xb4\xa1\x12\x25\xd5\x0d\x97\x8b\x34\x11\x23\x71\x03\x92\x68\x55\x95\x31\x10\xeb\xa9\x02\x7e\x95\x3c\x06\xa6\x99\x8f\xe1\xbd\xf7\x9f\x61\x0b\x40\x32\x7c\xae\x1f\x08\x3d\x72\x9a\x99\xab\x8c\xcb\x79\xa4\xca\xf9\xe8\x76\x84\xc2\x39\xad\x8d\x62\x73\xeb\xe4\x04\x6d\xcc\x45\x09\x74\x55\x71\x3d\x8f\x52\x39\xba\xc8\x80\xd7\x72\x74\xcd\x57\x37\x42\x00\xad\xe8\xe6\xe9\xc8\x02\x0d\xfa\x61\xbf\x3f\x1a\xb1\x44\x01\x42\xaa\x19\x97\x2c\x05\x72\xe5\x8c\xc7\x82\xcd\x14\x10\x49\x14\xec\x3f\x60\x48\x9a\x95\xe2\xb7\x4a\x68\xa3\x59\xa1\xb4\x4e\xa7\xd9\x8a\x2d\x81\x75\xb6\x2c\x79\x51\x00\x54\xdf\xac\x0a\xe1\x48\x35\x44\xee\xfa\xbd\xe7\x2a\x88\xd9\x11\x52\x88\x2e\xe8\xf8\x87\xac\x74\xbf\xdf\x59\x8a\x21\x0b\xea\xdf\xba\x50\x52\x8b\x21\xa3\x03\x0f\xfb\xeb\xbe\xa5\xaa\x8a\xd7\x3c\x17\x17\xe6\x96\xc1\xb1\x54\xb1\xb9\x5b\x13\xdf\x53\xae\x85\xa3\xf1\x12\xce\x3d\x83\xad\x0b\xd8\x58\x95\xb9\x66\x66\x21\x68\xbd\xc5\xbb\xa5\x86\xaf\x9f\x23\x9f\x1b\x62\xb3\x4a\xc6\x2c\x48\x9a\xa5\x90\xfd\x05\xb6\x51\xea\x52\x98\xaa\x94\xcc\x1d\x4b\x04\xd4\xca\xe8\xc2\x9e\x55\x10\x0e\x59\x0c\xd4\x48\x3c\x90\xc2\x9d\x32\x71\xc4\x93\x44\xd7\x2f\x98\x51\x6d\xc5\x5b\xee\x7d\x70\x2b\x00\xee\x97\x90\xe6\xfb\x9e\x28\x1e\xdc\x5f\x95\xa6\xdf\x03\x39\xd8\xf8\x9c\x79\x32\xa0\xcd\xe3\xa1\xe0\x6b\x58\x8d\xfe\xc5\xb3\x4a\x04\xcd\x41\xdd\xad\xc3\x28\x40\x1f\x92\x73\x00\xbd\xe1\xc0\x6b\xc1\x3c\x9b\x8e\x26\x05\x97\xfd\x1e\x88\x7f\xf5\xe6\xf9\x9b\x31\x0a\xce\xc8\x1f\x4a\x91\x71\xe3\xa4\x4f\x65\x51\x19\x96\x70\xc3\xff\xd1\xef\xa5\x33\x56\xf0\x12\x08\x20\x26\xee\xba\x4d\xed\x45\xa9\xf2\x9a\x3d\xe0\x28\xfc\xc6\x87\xff\xec\x9c\xc9\x34\x43\x4d\xf5\x80\x91\x2d\x64\xc3\x4b\x82\x72\xec\x0f\x5b\xab\x17\x8b\x34\x4b\xde\xcc\x82\x0d\xb1\x8d\x12\x42\x90\x6d\xcd\x44\xa6\xc5\xc3\x28\x23\xb8\xf5\x7b\xe2\xf8\xfb\x54\x26\xd1\x44\x98\x40\x17\x43\xd6\x7a\xfb\xee\xed\x85\x3d\xa7\x4b\x59\xe5\x21\x09\x0f\xa7\x81\x52\xeb\x22\xba\x02\xea\xa2\x0c\xc2\xe8\x95\xfc\x55\xc4\x88\xed\x99\x16\xb0\xe1\x73\xf0\xf2\xea\xea\xed\x4b\xc1\x13\x51\xea\x03\x4b\x17\xbc\x2c\x53\x20\x59\x46\xf6\x45\x08\xca\xc3\xfd\x3c\xad\x39\x83\x86\x9f\x43\x06\x41\x38\xba\x44\xe3\x98\x05\x10\xb6\xab\x2c\x91\x8f\x0c\x1c\x16\x32\xd3\x18\xef\xc2\x52\x66\xc1\xe7\x37\xe1\x80\x6c\xc9\x0a\xef\xe8\x24\x51\x82\x4e\xe1\xfb\x01\xac\x94\xab\xda\x8f\xf1\x47\x2a\x34\x3b\x3b\x79\xdf\xe5\x04\x04\xbb\xe3\x02\x9b\x0f\x39\x43\x8f\xa0\xde\xaa\x2c\x8d\x57\xec\xdd\xe6\xd9\xed\xe7\xbd\x61\x89\x98\xa5\x12\xb6\xe3\x96\x32\x84\x38\x7c\x1d\xd9\xcd\x7c\xc0\x56\x68\x03\x22\xcf\x20\xb3\xa9\xd9\x0c\xed\x16\xe5\xb2\x71\x07\xce\x6c\x8a\xce\x2c\xc1\xa6\xd3\x1c\xb5\xa1\x66\x8e\x30\x37\x46\xe4\x85\xd1\x51\xbf\x57\xa3\x06\x21\xfb\xf0\x11\x93\x52\xf4\xbc\x2a\xc1\xf6\x95\xf5\x0b\xda\x15\xb0\x62\x91\xde\x08\x4b\xd7\x57\xc4\x90\x71\xcd\x96\x22\xcb\xf0\x1b\x17\x4b\xa1\xab\xcc\xc0\x4e\x84\x5d\x27\x3e\xe7\xee\x8f\x34\xfb\xf4\x5c\x7d\x62\xb9\x30\x0b\x95\xc0\xe6\x44\x3d\x68\x45\x80\x21\xdb\x13\x00\xa6\x4a\x65\x4e\x67\x13\x10\x26\x13\x87\x34\x67\x16\xdc\x34\xa7\xc7\x6b\x6e\x99\x92\xb1\xb0\xda\xdc\x25\xd1\x8a\xe9\x1d\x0a\x05\x82\x2b\x55\x31\xbd\x40\x53\x6b\xb6\x11\x0d\xed\x53\x2d\x20\x0d\x26\x8c\xcf\x0c\x26\x1f\xc3\x66\x3c\xcd\x40\xc3\x36\x12\xee\xec\x17\xb2\xfd\x9a\xf7\x42\xf7\xd6\xca\xdd\x29\x3b\x62\xf4\x66\x42\xbb\xad\x7d\x2b\x82\x4c\x08\x07\x61\x39\x93\x4a\x1e\xbf\x7d\x33\xb9\x1a\xda\xa7\xef\xae\x2e\x5e\x6e\x52\x27\x09\x73\xf6\xfe\x7d\xd4\x6f\x22\xdf\x2b\x83\xd6\xaa\x01\xda\xb0\xb8\x2a\x31\xc8\x64\xab\xda\x5a\xe4\xca\x1e\x42\xad\x0e\x01\x29\x6a\xb5\xe7\x70\x0f\x08\x6c\x4f\x1b\xb8\x60\x5b\x27\x0e\x46\x53\x74\x1d\xbb\x7f\xf4\xa8\x13\x17\x7f\x5c\x3c\xf8\xfd\x77\x94\x28\xfa\x91\x8c\x89\x9d\x9f\xb3\x01\x0a\x3c\xe8\x7a\x8f\xe2\xe3\x02\x45\x10\x08\x54\x10\x11\x4d\xa5\x2f\x54\x22\xd8\xff\x83\x73\x9f\xf8\xb1\x65\xc6\x21\x90\xfa\x41\x02\xec\x42\x38\x35\x5f\xde\x02\x73\x20\x6b\xca\xb3\x4e\xeb\x93\x4c\x6c\x20\x5a\x4e\x6c\xad\x6e\x0f\xfe\x61\xd3\x9b\x81\xe7\x41\x49\xe0\xde\x52\xb5\xe3\xed\x92\x61\x34\x88\x4b\xc1\x35\x7a\xf8\x92\x83\xe1\xa1\x7d\x68\xaa\x4c\x84\x59\x0a\x21\x9b\x83\x1f\xb3\xd3\x93\x93\x21\x7b\x82\xff\x7d\x89\xff\x7d\x8d\xff\x61\x80\x38\x7d\x0a\x4a\xc8\xc1\x7c\x52\x6b\xc5\x9a\x3d\x1e\x1d\xb3\xaa\xc0\xfc\x77\xf6\x39\xfb\x35\x85\x80\x51\xd6\x47\xdb\x2d\xc5\x03\x0c\x1a\x33\x47\xce\xaf\x45\xb0\xb5\x3c\x64\x67\x10\x96\x25\xa4\x0e\x84\x00\x1e\x6b\x23\xff\x71\xc3\x11\x10\x00\x3e\x31\xfb\x03\xbf\xd1\x6b\xb1\x0c\xea\x87\x09\x95\xbf\x01\x21\xbc\x56\x4b\xc8\x48\x3f\xc9\xf4\xf6\x35\x97\xca\xa6\x46\x2a\x0e\x4e\xa2\x93\x33\x06\x2a\x41\xb1\xce\xb0\x3e\x8b\x41\x00\x27\x57\xbf\x87\x55\x66\xea\x88\xcf\x29\xae\xd7\x16\xf1\x21\xfd\x08\xd9\x94\x58\x7b\xcc\x5a\x4c\x07\x41\x00\x1c\x45\x2f\x32\xc5\xcd\xd3\x2f\x83\xf0\xe8\x49\x78\x7c\x1a\x1e\x89\xa3\x99\x7b\x83\x48\xb8\xbf\x15\xec\xe8\x9c\x3d\xf1\xad\x0a\xbe\xfe\x6b\x7d\x77\xdf\x01\xff\xef\x3b\xf0\x6b\xd5\x9d\x35\xac\x93\x62\x82\x94\xd8\xa6\x58\xe5\xa1\x43\x35\xb9\xc3\x65\xe1\x36\x81\xc3\x8e\x8b\xe1\x00\xb2\x2c\x80\x01\xb0\xa8\x95\xdb\xa2\xf0\x27\xb3\x40\x3b\xec\xf3\x6c\xc9\x57\x5e\xbc\x40\xf9\xf7\xec\xf6\x47\x13\xae\xc7\x83\x55\x6b\xdd\x0b\x91\x82\x5c\xbd\xd7\x56\xc3\xcf\x10\xa0\x7c\x1d\x35\xda\x00\xd5\x2e\x99\xeb\x14\xad\x19\x2b\x50\x75\x09\x6d\xac\x2d\x1d\x6c\x5b\xce\xd4\x14\x2b\xb8\x47\xba\x29\xc6\x9a\x32\x88\x24\xda\x22\x8f\x35\x76\x4d\xb4\xae\x3f\x87\x6c\x4f\xc9\x15\x6e\x43\xfa\x3d\x92\x5b\x41\xfa\xb6\x8b\x00\xca\xc3\x96\x9c\x77\xeb\x16\xe5\xd0\xef\x73\x8e\x9a\x5a\xf0\xef\x68\xda\xea\x2d\xa0\x07\xb8\xde\xea\x78\x5c\x8b\xd3\xe6\x0b\xba\x1c\x5f\x4c\x72\xb4\xcf\x00\xd5\x79\x45\xa3\x8b\x73\xa8\x78\xbd\xdf\xe4\x24\x4d\x66\xc1\x7d\x36\x6b\xd1\xc6\x30\x6d\xe7\xd4\xe1\xe3\x76\xa1\xf1\xf3\x3e\xd5\x7c\x13\x0e\xe9\xca\x2f\x8c\xa6\x2a\xc1\x1c\xc5\x62\xec\x83\x97\x50\x3f\x12\x80\xb2\x9b\x45\xec\x0d\xc0\x96\xcb\xd4\xae\x51\x14\x24\x00\x9e\x41\x4e\x4b\xd0\x7c\x78\x42\x74\xa9\x95\xae\x66\x33\xf0\x4e\x25\x9b\xe8\x96\xb4\xb6\x22\x74\xac\xbb\x22\x36\x11\x82\xd0\x90\xe1\xf1\x68\xa4\x0d\x8a\x03\xf6\x06\xd1\x79\x49\x63\x0e\xc2\x00\x67\xd2\xa3\x27\x5f\x9c\x7c\x75\xf2\xf5\x57\x4f\x47\xb8\x17\x4e\x3a\x90\xe3\x63\x35\x3b\x46\xdc\x63\x47\xfb\x18\x13\xaf\xaa\xcc\x71\xae\x92\x74\x86\xb1\xa1\x59\x01\xda\xc6\xe9\x02\x18\x04\x67\xc5\x69\x12\x9d\x41\x19\x3d\x43\xe1\xbd\x06\xa6\xad\xb0\x5e\x0f\x10\x6c\xa4\x3c\x67\x76\x94\x04\xca\xe5\xc9\x77\x59\x16\x58\x5c\xcc\x1f\xed\xa8\x89\x54\x5a\x7d\x10\xac\xc1\xab\x35\x1e\xa6\xcd\x64\x75\xdd\x8b\x69\xef\x1b\x8a\x7e\xdf\xd4\xef\x1e\x3f\x26\xfc\x4e\xd6\x7a\x65\x42\xbd\x9d\xe3\xe3\xb5\x2a\x2e\x32\xa5\xa1\x1f\xa3\xe1\x18\x66\xda\x67\xa4\xfe\x00\x78\xa6\xbc\xd6\x73\x34\xc0\x6e\x12\xc7\x02\x45\xe7\x5a\xa0\x56\x77\x65\x37\x75\x9c\x41\x4c\xcf\x84\x0c\x6a\xd3\x0b\x31\xe4\x7f\xe6\x1b\x9f\xcb\x34\x36\xb3\xd8\xbe\xcd\xf2\x38\x85\x33\xba\x76\x7b\xc1\xe9\x12\x87\x75\xe3\x41\x36\x69\xcd\x8d\x0c\x03\x0d\x64\xce\x53\x59\x27\x0d\x64\x36\x22\x0c\xb4\xe9\x9e\x2d\xa3\x33\x21\x8a\x86\x91\x0f\x8e\xc1\x8f\x61\x3b\x55\x3b\x1e\x5c\xd4\x8d\xd3\x32\xae\x52\xf3\x0c\x59\x11\x25\xf5\x7f\x69\x5e\x64\x22\xc7\x59\x9c\x0d\x65\x16\x82\x4d\x2d\x08\x2b\xb0\xe1\x2a\x25\xe5\x6c\xe8\xb4\x21\xf1\x4c\x2b\x9a\x2d\xe0\xe1\xe0\x37\x74\xc4\x10\xda\x45\x0b\x19\x5c\x86\xb3\x1c\x53\xdc\x12\x3c\x44\x60\x4b\x41\xf6\xad\x24\x64\xf7\x4c\xcd\x3d\x56\x18\x99\xa0\x2d\xe5\x08\x06\x4b\x81\x69\xa6\xe2\x6b\x2a\x00\x9a\xe2\x61\x56\xaa\x9c\xcd\x15\x0d\x36\x16\xa5\xaa\xe6\x0b\x97\xd6\x3a\x44\x3a\xd4\xd2\x5a\xfe\xed\x87\xfa\xb2\x9e\x23\x40\xb3\x18\x3b\x70\x69\x66\x80\xf4\xb1\xcf\xd1\xf7\x34\xfe\xfb\x81\x7e\xa0\x36\xd1\x1d\x1c\xea\x64\x72\xf9\x06\xda\x33\x86\x03\xd2\x08\x9f\x48\xd7\x2f\xed\x08\x14\x16\x2f\x71\xd4\x89\xf3\x41\x91\x63\xd5\x46\x65\xcc\x66\x56\xca\x6e\x52\xce\xc0\x58\xc1\xc3\x8f\x35\x02\xda\xc9\x68\x44\xc5\x91\xd0\x71\x99\x4e\x6d\x4d\x8c\x1a\x26\x6d\x61\xb2\xe7\xf5\xe6\x4e\x0d\xdb\xbb\x6d\x74\x70\x85\xcb\x07\x3e\x56\x66\xf6\x09\x67\xc5\xe3\x01\x12\x1b\x7c\xea\xf7\x48\x1d\x0f\xc6\x92\x00\x8d\x58\x2e\x49\x5c\xa8\x0a\x58\xe8\xf8\xa4\xf6\xb5\xc3\x2a\x3d\x68\xc4\xa6\x91\xc8\x5e\xdc\x6d\x6c\xd1\x40\x37\xb8\x6f\x6d\x71\x8c\x93\xf2\x07\xe0\x6e\xa0\x91\xc0\x2b\x7d\xd1\xb2\xa4\x37\x05\xb4\x1f\xde\x87\x6a\x8a\x86\x40\xda\x01\x4d\x0a\x50\x50\xf6\xcb\x39\xb1\xf5\x02\xa2\x79\x55\x8a\x43\x0a\xd8\x85\xde\x25\x92\x65\xe8\xe3\x3e\xb1\x43\x44\x5a\xd0\xfb\x88\x4d\xaa\x38\x16\x5a\x3f\x90\x98\x83\xde\x26\x36\x59\xa8\xd2\x38\x2d\x88\xe4\x3e\xf1\xda\xd0\x3b\xb4\x3c\x86\xee\x57\xd5\x1e\x86\xae\x16\x98\x05\xdf\xc2\x39\xbd\x13\x58\x88\x21\x53\xfb\x89\xec\x42\xef\xd0\x83\x28\x0b\x09\xe5\x81\x4c\x39\x68\x24\x72\x61\x3b\x19\xa8\x71\x5c\x4f\x73\x79\x2b\xe2\x0a\x33\xb6\x35\xee\x36\x91\xf8\x30\x34\x12\xfc\x01\x9c\x5e\xc6\xab\x2b\x65\x78\xf6\xa3\xe0\xf2\xb0\x71\x67\x1e\xf4\x2f\x39\x80\x03\x89\xba\xde\x83\x60\x46\xf1\x21\xc8\xba\xe2\x1a\xa4\x8a\xed\x40\x42\xc9\x2b\x8b\x5e\xc9\x99\x7a\x1e\x40\x97\x4a\xe3\xdd\x9c\x17\x1f\x6c\x00\xf8\xd8\xcc\xf3\xee\xd6\x98\xe5\xda\x3e\x3d\xee\xf2\x62\x11\xf9\x51\x02\x87\xaa\xbe\x2b\x77\xa2\x20\xd2\x26\x38\x6c\x50\x3c\x0f\xde\xc5\x73\x28\x1b\x18\xc2\xeb\x74\xdc\x71\x1b\xaf\x2b\x14\x10\x72\x97\xbf\x8e\xb7\x85\xdb\x85\xe9\xc0\x6d\xbb\xe9\xb8\x13\xb7\x05\xb3\x97\x46\xed\x0c\x87\x68\x38\x98\x1d\x1a\x5b\x4e\x39\xee\x92\xa1\x0d\xb3\x4b\xa2\xb5\x7d\xb7\x1a\xf6\x6d\xdf\xe1\x82\xe3\x6d\xdc\x5d\x98\x5d\x32\xce\xf3\x0e\xb1\xe0\x60\x08\xf7\x3e\x87\x1b\x03\xee\x3d\x2e\x4c\x74\xb2\x2d\xaf\xdc\xb5\x41\x11\x6d\x7b\x2e\x20\xae\xc3\x7e\xab\xfd\xda\x2d\x60\x42\x70\xe8\x14\x1a\x26\x37\x22\x07\x0b\x4e\xa1\x72\x8f\x69\x62\x06\x5e\x6b\x67\xe8\x92\x67\xed\x1a\x0a\x8b\x31\xae\x75\xaa\xa9\xfe\xd2\xc2\x18\x2a\x65\xbc\x52\x0b\x0a\x27\xa1\x17\x2a\x4b\x34\x95\x5b\x95\xc4\x0b\x06\x03\x8f\x34\x2e\x28\x0a\xa8\x5f\x6d\xfb\x3e\x15\xd0\xc8\xa4\xaa\x8c\x08\xfd\x27\xf0\xfd\xd2\x54\x12\xf6\xc8\xa0\xaf\xf3\x2a\x17\x9c\xdd\xe0\xd5\x85\x6d\x7b\xd8\x1c\x44\x80\x4a\x8f\xaf\x90\x15\x1c\xf4\x51\x81\x9b\x6a\xa2\x82\xb7\x50\x43\xa6\x15\xf6\x4a\xc0\x06\x5e\x5f\x63\x15\x09\x55\x2a\xcd\xf7\x25\x14\x89\x9a\xe9\xaa\xa0\x8b\x6d\xec\x05\xa6\xa5\xe2\x09\x74\x5e\x3b\x62\xe4\x58\xe9\xc6\x9a\xea\x26\x9a\xe8\xef\x16\x4f\x2c\x80\xf8\x15\x5a\xf6\xa1\x9d\x22\x30\x0d\xbd\xd4\x81\x1b\xeb\x54\x6b\x08\x4a\xa3\xb3\xa7\x51\x53\x12\xba\xba\x0e\x3b\x01\x3c\x2d\x7b\x20\x3d\x87\x33\x31\x20\x6d\x5e\x5f\xb7\x40\xf3\xe1\xde\xd3\x5c\xcf\x5f\xa3\x8a\xbd\x0b\xc9\xde\x6e\xd1\x72\x06\xa7\x26\x24\x05\x60\x7b\x4b\x25\x85\x89\x7e\xa0\x97\xc1\xc0\xc4\xc5\x60\xc8\x06\xa7\x4f\xbe\x8a\x4e\xe0\xdf\xe9\xf8\x64\xb0\xa7\xb7\x2a\xb8\x4c\xe3\xc0\x5e\x15\x51\x4f\xd5\x9b\xdb\x2b\xd0\x68\x82\x2a\x0a\x36\xdb\x74\xf1\x13\x5a\x78\x4f\x54\x54\xdf\xcf\x82\x6c\x0e\x67\x4b\x2b\x76\x7a\xc2\xea\x29\xeb\x54\xc4\xbc\x72\x1d\x0c\x52\x03\x23\xbd\x6d\x8e\x86\x43\xa5\xa2\xb0\x09\x38\x3d\x39\xae\x87\xa0\x44\xcd\xf9\x2d\x94\xfa\x32\x81\x6e\x96\x5d\x81\x79\x40\x43\x85\x63\xbc\x0a\x8f\x3f\xa7\xc9\x6d\x17\x49\x30\x96\xb9\xa0\xab\x89\x32\x5b\x39\x6a\x3c\x8e\x2b\xdb\x7c\xa0\x7f\xb9\x6b\x23\x38\x7c\xe8\x21\x84\x9d\xd3\xa0\xf9\x41\xe5\x8d\x46\x8a\x37\xdb\x78\xc1\x56\x33\xae\xaa\xd2\xdd\x2e\x39\x6a\xb5\x35\x3b\x40\x70\xb6\x9c\x9b\x78\xd1\xb0\xf2\x48\xbb\x36\x05\x5b\x7c\x30\x64\x54\x0b\xf4\x37\xb2\xab\x6f\x72\x24\x35\x0d\xfb\x58\xbc\xc0\xe1\xac\xa6\xce\x13\xb0\x5e\x50\x92\x84\xb0\x60\x67\xc6\xed\x7b\x11\x82\x01\xab\xa7\xec\x3b\xc1\x19\x38\x8d\x9e\x9b\x84\xbb\x95\x9f\xef\xd6\x2d\x04\xc8\xe3\x18\xea\xb6\x70\x88\x3c\xbe\xbf\x23\xa3\xa0\x56\xfb\x17\x56\xcf\x8c\xdd\x6a\x7c\x1d\x6c\x5d\xd1\x80\x11\xb0\x66\x94\x05\x07\x86\x33\xc7\x14\x83\x00\xd2\xe8\x55\x65\x86\xdb\x0c\xdc\x98\x62\xc0\x1e\xb3\xda\xbe\xa2\xef\x92\x04\xef\x57\x27\xb4\x3d\x59\x38\x36\xb2\xbf\x35\xe6\x4d\x36\x09\x8e\xe2\x2a\x81\x60\xf0\xcf\xcb\x2b\xb0\x71\x20\x39\x44\x6b\xb6\x08\x1d\x06\x4e\xea\x8b\xde\x02\x55\x33\x0b\x6c\x0d\x80\xe3\x2f\x09\xc7\xe5\xae\xbc\xeb\x73\x00\xfd\x60\x8f\x02\xd6\x3d\x66\x9f\xeb\xe6\x02\x15\x3f\x38\x2f\x4b\x65\x25\xe8\x17\xa9\xaf\x67\x21\x1b\xf6\xbc\xbf\x4d\x89\x26\xd5\xd4\xf6\x60\x38\x5e\x0b\x06\x40\xe9\xff\xbc\xf9\x98\x1d\xac\xfd\xf6\x1f\x67\x18\x4f\x50\xdc\x6c\xe6\xfe\x16\x33\xba\xb4\xd1\xcf\x6d\x4d\x13\x9b\xed\x82\xce\x2e\x6d\xee\xc3\xb1\x66\x8c\x7e\x92\x39\x2f\xf5\x82\x67\x81\x1d\x01\x05\xe2\x26\x7a\x0e\x01\x3b\x08\x43\x90\x57\xec\xde\x66\xef\xb0\xe5\xf8\x22\xc2\x11\xb5\x9c\x00\x3d\x70\x7b\x5f\xa8\x3c\xc7\xbf\x2b\x3a\x8c\x8a\x46\x8c\x06\x5f\x8f\x0e\x5b\x5e\xf0\xc1\xd6\x9e\x1f\x37\xa0\x3b\x4b\x60\xd0\xa2\x6f\xd7\xc1\x6a\xd1\x3b\x7d\xc7\xb4\x39\x12\xd8\x4b\x31\xd8\x88\x14\x67\x77\x0d\x70\x70\x1a\x12\xf0\x2c\x85\x84\x68\x8f\x1f\xdc\xfc\x11\x64\x35\x8d\x1e\x88\xda\xf6\x69\x29\x0f\xf3\x49\xd8\x5e\xc2\xb6\xd1\xee\xb5\xe0\xb5\xe7\x27\x2d\x8c\x2f\x42\x3f\x8e\x52\xee\x2c\x20\x71\x03\x94\x4e\x71\x8c\xb0\x14\xb6\x36\x87\xdf\x0a\xef\x9b\xd1\x50\xb6\x38\x68\x74\x5d\x8f\x4a\x7b\x1d\x41\xc0\xd3\xcb\xe6\x86\xa9\x86\xad\xfb\x80\x24\xca\xea\xd2\x3f\xbc\xe7\x5c\xeb\xf3\xe9\xac\x8f\xf1\xb8\xbb\x2b\xe7\xbf\x9b\x3f\xf7\xf7\x22\xc0\x90\x77\x6f\x06\x5e\x1a\xec\xa3\x1f\xb2\x6f\x59\x2b\xea\xfe\xdd\x0c\x6d\xfc\x12\xff\x5b\x23\xce\x3a\xbc\xb7\xc6\xfb\x8b\xb3\x76\x90\x3f\x89\xec\x34\xeb\xae\x35\xbd\x1c\xef\x8c\x2f\xeb\xca\x04\x5e\x25\x91\x37\xee\x1a\xd6\xb9\xde\xc6\x25\x52\x4b\x57\x0c\x73\xc3\x44\x3b\xad\xb5\x82\xd2\x44\xb4\x7d\x9b\xf5\xed\x79\x73\x9d\xd5\xb3\x7f\x41\x01\x16\x6c\x87\xe5\x88\xcd\xa5\xdb\x86\x66\x87\x29\xcd\xd7\x7c\xa7\x71\x79\xd8\xdb\x4c\x95\x34\xba\x0d\x06\x67\xef\xdf\x53\xc5\x63\xb7\x6e\x71\xb3\x6e\xf2\xc5\xce\xc4\x93\x0a\xa0\x07\x8e\xff\x7b\x76\xec\xfb\x07\x14\x75\x78\x5a\xbc\x67\xe8\xdd\xcc\xbb\xef\xd3\xdf\x9f\x56\xdf\x01\xed\xad\xfb\x6d\x3e\xac\xea\xd8\xe8\x88\x49\x85\x57\x64\xd4\x29\x92\xa0\xf6\x62\x7b\x05\xd5\xd6\xd1\x28\xec\x9c\x25\xff\x3b\x00\x00\xff\xff\x96\x55\x88\x81\x90\x2a\x00\x00")

func _hardcodedDoerGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedDoerGo,
		"_hardcoded/doer.go",
	)
}

func _hardcodedDoerGo() (*asset, error) {
	bytes, err := _hardcodedDoerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_hardcoded/doer.go", size: 10896, mode: os.FileMode(420), modTime: time.Unix(1515615665, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __hardcodedMiddlewareGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\x5f\x6f\xdb\x38\x12\x7f\xb6\x3e\x05\x4f\x87\xa6\x52\xe0\xca\xb9\x03\xae\x0f\x29\x7c\x40\x91\x26\x9b\x62\x9b\x36\x68\xd2\xed\x43\x51\x2c\x68\x89\x92\xd8\xc8\xa4\x96\xa4\xac\x18\x41\xbe\xfb\xce\x0c\x29\x59\x4e\x5c\xa0\x6f\x5b\xa0\x09\x35\xe4\xfc\xe6\xff\x9f\xb4\x3c\xbf\xe3\x95\x60\x56\x98\x8d\x30\x51\x24\xd7\xad\x36\x8e\x25\xd1\x2c\xce\xb5\x72\xe2\xde\xc5\x70\x2c\xd7\xf4\x4b\x09\xb7\xa8\x9d\x6b\xf1\x6c\x3a\xe5\xe4\x5a\x2c\x0a\xb1\xea\xaa\x38\x8a\x66\xba\x15\xca\x19\x9e\x4b\x55\xb1\xb8\x92\xae\xee\x56\x59\xae\xd7\x8b\x09\x7d\x7a\x7e\x55\x69\x80\x71\xbc\xb2\xbf\xf8\x7a\x11\x74\xa9\x74\x7b\x57\x65\x52\x2d\xce\x1a\x01\x2a\x2f\xee\xf8\x76\x23\x04\x3c\xc8\x36\xaf\x17\x8d\xae\x2a\x61\xe2\x28\x8d\xa2\xc5\x82\x5d\x73\x25\xf3\x2b\x59\x14\x8d\xe8\xb9\x11\x0c\x6e\x2d\xe3\x6a\xcb\x5a\xbc\xb0\x19\xbb\xd0\x86\x29\xdd\xcf\x59\x2f\x5e\xc2\x3d\x5a\x2c\x55\x27\x98\xab\x8d\xee\xd1\x10\x57\x0b\xff\x98\x75\x2d\x22\xe2\xb7\x75\xe0\x33\x66\x35\x7c\x48\xcb\xd6\x7c\xcb\x72\xc3\x6d\xed\xdf\x1a\x9d\x0b\x6b\xb3\xa8\xec\x54\xfe\x54\x7e\x52\x33\xf4\x5e\x76\xc9\x15\x50\x4c\xba\xf7\xc5\x1e\xa2\x99\x11\xae\x33\x6a\x8f\x7c\x01\x38\x09\x82\x25\xbd\xa7\x7f\x16\xb6\xd5\xca\x8a\xaf\x46\x3a\x61\xe6\xcc\xb0\xe3\x40\xff\xab\x13\xd6\xa5\x88\x33\x2b\x44\x09\x88\xc4\xe6\x09\x33\x32\xe2\xdc\x18\x76\xba\x64\x46\xe4\x1a\x3c\x97\xa4\x78\x21\x4b\x36\xde\x2d\x97\x4c\xc9\xc6\x33\x04\x65\xf0\xf8\x88\x3f\x36\xdc\x30\x61\xe8\xbf\x86\x3c\x01\x8a\xed\xa5\xcb\x6b\x36\x45\x1e\xce\x59\xe2\xb6\xad\x08\xa2\x73\x6e\xd1\x69\x06\xfc\x79\x4a\xc8\x08\xb3\x64\x90\x53\xd9\x39\x82\x95\xc9\xc0\x96\x8e\xcf\x49\xca\xf4\xf5\xf0\x04\x49\x60\x1d\xef\x1a\xf7\x13\xb0\xb8\x53\x77\x10\x53\x15\xe2\xf6\xe2\xdf\x1b\xa6\x4b\x86\xfa\xb0\x17\xb7\xf1\x7c\x44\xda\x9d\x52\x6f\x24\xfe\xf4\xf9\x93\x5d\x18\xbd\x3e\xf3\xe9\x9f\x98\x6c\x38\xa5\xa9\x97\xf1\x2e\x89\x89\x35\x9e\x93\x06\x81\xe7\xea\x21\x06\x5d\xe2\x53\xd4\x7d\xce\x62\xca\x12\xcc\x5e\x01\x24\x6f\x7d\x42\xb5\x92\xdd\xe0\x0d\x80\x3d\xa6\x63\x60\xf6\x3c\xf0\x48\x81\xa9\xb3\x1b\x2c\xc9\xcb\xdb\xdb\xeb\x04\x12\x14\x6f\x80\xe1\x91\xf2\x1a\xb0\x5d\x67\xf7\x33\x81\xf5\x86\xb7\x90\xdd\x10\x5d\x4f\x06\x02\xd2\x23\xb2\xfc\x20\x07\x68\xd5\xe5\x0e\x83\x74\x20\xb3\xa2\x99\xe7\x61\x52\x39\x14\x4b\x09\x9d\x58\x76\x7c\x08\x2a\x65\xf4\xfb\x52\xf0\x02\xf2\x2a\xd7\x85\x40\x36\x8a\xbf\xcd\x02\xce\x92\x21\x1d\x09\xfb\xbc\xd9\x53\x56\xb2\x92\xb4\x0e\xc5\xff\xa9\xfd\xc8\xd7\x22\xa8\xfb\xe0\x5d\xf0\x15\x5a\xc6\xed\xde\x35\x2f\x0a\x4b\x35\xa8\x5b\xa6\x90\xe0\x34\x78\x23\x34\x31\x56\x42\xa5\x77\xe0\x94\xd5\x96\xde\x0c\x8d\xaa\x91\x2b\xc3\xcd\x36\x63\xef\x1d\x5e\x5b\x84\xe6\xac\xd5\xa0\x3d\x38\x68\x25\x72\x8e\x4c\xd2\xbd\xb4\x2c\xe7\x4d\x23\x0a\xa0\x35\xba\x07\xeb\x26\xad\x00\x2a\xd5\x0b\x56\xcd\x96\xf5\xd0\x0f\x40\x72\xcb\xad\xd7\x46\xaa\x52\x87\xe6\x01\xdd\x02\x6e\x6a\xbe\x41\x44\x10\x63\x85\xdb\x09\xcb\xd8\xf9\x46\x20\xaa\xee\xaa\x1a\xef\x0b\x2d\xac\x7a\xe9\x58\x5e\x73\x55\x09\xc2\x1a\x8c\xe9\x51\xb2\x6c\x1a\xc2\xf2\x6d\x09\xb0\x43\xf7\xd8\xd9\x0c\xc2\xd6\x1c\xc0\x79\xd0\x16\x38\x3a\x2b\xb5\x0a\xdd\xe9\x99\x0b\x93\xdc\xdd\x0f\xcc\x43\xd2\xcf\xc1\x9d\x83\xf7\xe1\x61\xfa\xf4\x9e\x42\xec\xcc\xb5\xa3\xfa\x07\x80\xec\x0f\xde\x74\x22\xd9\x0b\xdd\xc3\x63\x9a\x25\xc7\x01\x21\xc2\x96\x13\x58\xfe\xb5\x6b\x38\xc7\x81\xb4\x0c\x02\x23\x6c\x3b\xc1\x24\x80\x0d\x99\x1f\xf4\x85\x8e\x2b\xb8\x13\x98\xee\x4a\xf4\xcc\x42\xf9\x50\xcc\x0b\xc6\x4b\x0c\x1c\x5a\xfb\xe5\xf3\x07\x88\x82\xab\xa9\xfa\xe1\xdb\xf8\x06\x99\x21\x0c\x44\xbb\x6d\xa0\x34\xad\xf7\x1c\xf1\x07\x27\x85\x67\x83\x9d\xf3\x69\xe6\x68\x78\x60\xc0\xe9\xd4\x99\x2d\xdb\x48\xce\x26\x33\x2a\xbb\x01\x98\x69\xe7\x48\x49\x54\x89\x91\x46\x01\xe2\x5e\x5a\x87\xf5\x34\xca\xa8\x29\xe9\xed\xdc\xe7\x12\xbe\xf1\x76\x15\x3e\x4d\x71\xba\xec\x86\x57\x8f\x01\x5f\x09\x8c\x6f\x2d\x9b\xc2\x9b\xc5\x1d\xf1\x85\x88\x06\xef\xfc\x13\x13\x07\x2c\x7d\xeb\x9c\x58\xb7\x94\x76\x3f\x20\xa7\x07\xb3\xc1\x94\x4a\x38\x47\xd3\x14\xdb\xa1\xaf\x88\x12\xfc\x44\x66\x07\x1f\x64\x1e\xe3\x56\x63\x4d\xc1\x02\x02\xb3\xa5\x26\xb7\x0f\xa1\xe4\x43\x71\xfb\x6c\x94\x0a\x90\xa0\x0e\xb0\x02\xb0\x4a\x2a\x0d\xd2\x15\xca\xa6\x20\xea\x0e\xd3\x60\x8b\x05\xa6\x0a\x0f\x3d\xa1\xe3\x7c\x80\xdc\x59\xc1\x57\x00\xb4\xf0\x26\x40\xe3\x80\xcc\x40\x62\x76\x0d\xc9\x13\xf9\xc1\x67\xdb\x67\x81\x8e\x68\x74\xda\x7c\x4e\x53\xf1\x74\xb9\xf7\xe0\xb7\x46\xaf\x78\x83\xe1\xc0\x31\x8b\xb6\xcd\xce\xef\xf1\xd2\x25\xd3\x67\xd8\xdc\x2f\x43\x0e\xd0\x28\xf9\xc9\xe5\x19\x37\x46\x02\x92\xc9\x3c\x21\x4d\xdf\x90\xd4\x49\xf9\xcc\x40\xc5\x7d\x1d\x6e\xd0\x8f\xa8\x69\xe2\x0d\xa3\xd1\xc2\x44\x03\x3e\xfd\x35\x86\xf9\xde\xed\x19\x26\xdd\xa7\x32\xb1\x79\x4a\x48\xe3\x96\x61\xdb\xec\x42\x2a\x69\x6b\xc8\x77\xef\x69\xa9\x7e\x88\xdc\x27\x26\x7b\xff\x0e\x27\x81\xf6\xab\x17\x76\x64\x59\x4c\x2b\x80\x26\x62\x05\xf8\xc0\xe9\xd0\x8b\x6b\x7e\x27\x92\x35\x6f\xbf\xf9\x6e\xf1\x7d\x6c\x1a\xe8\xed\xe0\x69\x90\x38\xba\xf6\x3d\xc9\x4a\x80\x34\x56\xdd\xbe\xde\xb7\x40\xbb\xe2\x2d\xf9\xf7\x00\x7d\x70\xad\x1b\x7c\x3a\xdd\x81\x30\xc0\x38\x99\x0b\x80\xbc\x43\xc9\xee\x5b\xac\xdd\x2b\x4a\x63\xf3\xca\x5f\xc5\xdf\xdf\xe0\xe5\xc3\x74\x15\xf8\xe9\xfa\xf0\xb6\x28\x86\xaf\xe7\x40\xf3\x20\x2c\x1d\x76\x2e\x5a\x49\xc0\xb2\x0f\xba\xc2\xd9\x00\x3c\xc1\x6d\x7f\xc2\x06\x27\xe4\x46\x14\x71\x7a\x70\xd9\x3b\xc8\x53\x52\x90\x02\x8f\xdf\x31\xa0\x77\x9e\x41\xcb\x7f\x92\xbc\x41\x43\x1c\x0f\x94\x0d\x66\xea\x5a\xdb\xa6\x3e\xc8\x5f\x2c\xb6\x22\x1f\x9e\x71\x60\xc2\x52\xdc\xf0\x2d\xf6\x46\x3f\x26\x73\xc8\x80\xb5\x2e\x64\xb9\x85\x89\x86\xc6\x8c\x93\x22\x8e\x77\xf2\x97\xe3\x50\x41\x99\x7e\x7e\xf8\xab\x39\x7b\x32\x47\xe6\xec\xc8\x63\x50\xae\x59\xd3\x23\xd6\xd1\xa1\x95\xc4\x3b\x82\x2e\x4e\xd9\xf0\xef\xbf\x27\x27\x94\x08\xfb\x4f\x4f\x59\x3f\x1f\xdc\x8d\x7f\x90\x50\xe9\x5d\x09\x18\xc5\x05\x2c\x61\x98\x5c\xd0\xf4\x32\x4f\x48\x87\x37\xe8\x9a\xdf\xa5\xda\xbd\xd8\xa3\x7e\xbe\x3e\xa3\xf5\xcd\x9c\xab\x6e\x9d\x4e\x71\xbf\x98\x66\x02\x3a\xf4\x19\x32\xe7\x79\x20\x47\xae\x1b\x32\xe4\x0c\xb6\xa3\x91\xb9\x03\x97\xff\xe7\x75\x02\x3e\x08\x1b\x56\x3a\xec\xf3\x3b\x12\xfb\xff\x92\xfd\xef\xe4\x24\x64\x27\xa1\xd1\x06\xbb\x53\xda\x74\x62\xcc\x37\x0c\xeb\x47\x88\x1a\x8d\x95\xd0\x5b\xa7\x4d\xb7\x11\xb8\x07\x39\xb3\xc5\xcd\x85\x1a\x3a\x05\x75\x36\xb6\x0b\xaa\x11\x1f\xb9\x5f\x58\x03\x48\x57\xe0\x39\x3a\x1a\x04\x4c\xfb\x19\x26\x31\x68\xf9\xa9\x15\x86\x3b\x58\x5a\x68\x41\x39\xde\xb5\x32\x5f\x21\xbe\xe7\x4c\x77\x65\x30\x1e\x1d\x8b\x99\x34\xa4\xad\xd7\x28\x1d\x36\xe8\xbf\x03\x00\x00\xff\xff\xa3\xdc\x93\xdc\xf3\x0e\x00\x00")

func _hardcodedMiddlewareGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedMiddlewareGo,
		"_hardcoded/middleware.go",
	)
}

func _hardcodedMiddlewareGo() (*asset, error) {
	bytes, err := _hardcodedMiddlewareGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_hardcoded/middleware.go", size: 3827, mode: os.FileMode(420), modTime: time.Unix(1493241193, 0)}
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
	"_hardcoded/doer.go":       _hardcodedDoerGo,
	"_hardcoded/middleware.go": _hardcodedMiddlewareGo,
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
	"_hardcoded": &bintree{nil, map[string]*bintree{
		"doer.go":       &bintree{_hardcodedDoerGo, map[string]*bintree{}},
		"middleware.go": &bintree{_hardcodedMiddlewareGo, map[string]*bintree{}},
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
