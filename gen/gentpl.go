// Code generated by vfsgen; DO NOT EDIT.

package gen

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// files statically implements the virtual filesystem provided to vfsgen.
var files = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/assets.anko": &vfsgen۰FileInfo{
			name:    "assets.anko",
			modTime: time.Date(2019, 1, 4, 2, 42, 23, 580940004, time.UTC),
			content: []byte("\x23\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x70\x6c\x61\x63\x65\x68\x6f\x6c\x64\x65\x72\x20\x73\x63\x72\x69\x70\x74\x0a\x0a\x23\x20\x6a\x73\x28\x22\x6a\x73\x2f\x61\x70\x70\x2e\x6a\x73\x22\x2c\x20\x2e\x2e\x2e\x29\x0a"),
		},
		"/gitignore": &vfsgen۰FileInfo{
			name:    "gitignore",
			modTime: time.Date(2019, 1, 4, 2, 42, 23, 580974274, time.UTC),
			content: []byte("\x2f\x61\x73\x73\x65\x74\x73\x2e\x67\x6f\x0a\x2f\x6d\x61\x6e\x69\x66\x65\x73\x74\x2e\x67\x6f\x0a\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2e\x67\x6f\x0a\x2f\x67\x65\x6f\x69\x70\x2f\x67\x65\x6f\x69\x70\x2e\x67\x6f\x0a\x2f\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x2a\x2e\x68\x74\x6d\x6c\x2e\x67\x6f\x0a\x2a\x2e\x6d\x6f\x0a"),
		},
		"/manifest.go": &vfsgen۰CompressedFileInfo{
			name:             "manifest.go",
			modTime:          time.Date(2019, 1, 4, 2, 42, 23, 581024314, time.UTC),
			uncompressedSize: 3637,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x6d\x6f\xdb\x36\x10\xfe\x2c\xfd\x8a\xab\x80\x14\x64\x27\x4b\x09\xba\xed\x43\x3a\x17\xc8\xf2\xb2\x04\x68\x82\xa0\xc9\xd0\x0f\x41\x3e\x30\x12\x69\x13\x91\x48\x97\xa4\x13\x67\x45\xfe\xd3\xfe\xcf\xfe\xcc\x70\x7c\xb1\x65\xd7\xcd\x5a\x60\x01\x02\x59\xd4\xdd\x73\xf7\xdc\x1d\x1f\x72\xc6\x9a\x7b\x36\xe1\xc0\xac\xe5\xce\xe6\xb9\xec\x67\xda\x38\x20\x79\x56\xdc\x3d\x39\x6e\x8b\x3c\x2b\x1a\xdd\xcf\x0c\xb7\xb6\x9e\xfc\x25\x67\x61\x41\x39\xbe\x70\xfe\xa7\x79\x9a\x39\x5d\xdb\x29\xdb\xc3\x57\xd1\xfb\x55\xa9\x6b\xa9\xe7\x4e\x76\xf8\xd2\xe9\x09\x3e\x7a\xd9\x73\x7c\x2a\xee\xea\xa9\x73\x1e\x48\x7b\x7c\xeb\x8c\x54\x13\xff\xd3\x79\xa3\x3c\x2b\x26\xd2\x4d\xe7\x77\x55\xa3\xfb\xda\x4e\xe7\xa6\xd1\xfa\x83\xf7\x12\xb6\x7e\x10\x36\x40\xd3\x3c\xaf\x6b\x78\x10\x76\xc2\xd5\x3f\x7f\x1f\x20\x03\x90\x16\x18\x58\xc7\x9c\x6c\x02\xa7\x2a\x77\x4f\x33\xbe\x61\x65\x9d\x99\x37\x0e\xbe\xe4\xd9\x11\x73\x0c\xe2\xdf\xcd\x2d\x52\xce\xb3\x43\xa4\xa7\xdc\x35\xfa\x85\xdc\xf2\xec\x5c\xb7\xd7\xb2\xe7\xde\x0e\x93\xac\xf0\x2d\xcf\xae\x4e\x0f\xf6\x92\x7b\x34\x7d\x5e\xcb\xea\x6e\x2e\xbb\xf6\x9c\x29\x29\xb8\x75\x3e\xb8\x05\xbf\x66\xa1\x8f\xab\xb1\xf6\x55\x2e\xe6\xaa\x79\xc9\x91\x50\x20\x3d\x9b\xdd\x84\x40\xb7\x6b\x94\x4a\xe0\xc6\x68\x43\x91\xd3\x12\x78\x7f\x0c\x09\x81\xd0\x3c\x0b\x71\x70\xb5\x67\xf7\xfc\x05\xa8\x8e\x2b\x92\x40\x28\xcd\x33\x6e\x0c\x7a\xc5\xc2\x57\x9f\x58\x77\x4f\x42\x4a\x25\x14\x75\x51\x02\x66\x4e\x54\xac\x40\x09\x42\x82\xb6\xd5\x89\xec\xf8\x99\x12\xda\xa7\x96\xd2\xf3\x0f\x4c\x32\xb3\x8f\xd2\x35\x53\xff\xb3\x61\x96\x7b\xa3\x57\x63\x50\xb2\xdb\xcf\xb3\x2c\x33\xdc\xcd\x8d\xc2\xd5\x64\x20\x64\x75\x66\x8f\xa4\x21\x74\x68\xa0\x64\x97\x67\xd9\x73\x9e\x67\x99\x28\x41\xdf\x63\xa2\x21\xb7\x8a\x24\x5a\x27\x57\xf4\x46\xdd\xe6\x59\x26\x05\xbc\xd2\xf7\x3e\x66\xf2\x17\xbd\xab\x8e\x31\x29\x41\x0a\xa5\x43\x2f\x60\x67\xc7\x16\x25\x28\xba\x44\x56\x09\x3a\x95\xe5\xbb\xf0\x1a\x3d\xef\x5a\x50\xda\x81\x90\xaa\x85\x19\x73\x53\x10\xda\xac\xc3\x23\x3a\x8c\xb1\x8e\xf0\x13\x08\x85\xe1\x1e\x98\x81\x16\x07\x33\x4d\x64\x2a\xd6\x02\x53\x10\x15\xc1\x91\xa6\xab\xd2\xbd\x49\x44\x0f\xe3\x4e\xe5\x6d\x2a\x7e\x28\x55\x68\xc1\xfe\x18\x70\x07\x57\x17\xfc\xf1\x23\x67\x2d\x37\xc4\x6f\xf1\xc1\xfb\xa2\x6a\x96\x08\x71\x17\x60\xff\x3d\xd1\x55\x7f\x02\xdf\xf5\x0e\x79\x22\x19\x26\x1d\x62\x8d\x21\x08\x40\x85\xc8\x07\x5d\x47\xcc\x0f\xe0\xac\x93\x5a\xa3\xe2\xcb\x32\x06\x4c\xd4\xa7\x97\x3a\xd4\x0c\xf6\xec\xfe\x18\x50\x2a\xaa\x23\xee\x78\xe3\x06\xbb\x99\xa0\x37\xdd\x32\x7b\x51\x81\xaa\x53\x66\x2f\x0d\x17\x72\x41\x06\x78\x25\x14\xa8\x76\x75\x11\xe6\x4e\x0a\x90\x18\x22\xf9\x7c\x60\xd6\x9d\xa9\x96\x2f\x08\x4e\x49\x51\x15\xf4\x1d\x48\x64\x38\xda\x8b\x04\x87\xb9\x8d\xa1\xf7\xea\xf1\x34\xe3\xbf\x3f\x1d\x2f\x1c\x57\x56\x6a\x45\x84\xba\x91\xfb\xb7\x34\x15\x00\xff\x3b\x3d\xa9\x2e\x8d\x54\x4e\x90\xe2\xfd\xfb\xf7\x20\x14\xce\x0d\xd4\x75\x1c\x1f\x8c\x36\x40\xa6\x58\x85\xb0\xc3\x6f\x84\xba\x85\xf1\xba\xe4\xf9\x4c\x50\xec\xf6\x93\x5c\xf9\x5e\xe1\xea\xa0\x40\xfb\x43\x44\xff\x31\xea\x5e\xf0\x12\xb2\x8a\xef\x84\xfa\xaf\x28\x7f\x4b\x40\x9c\xfc\xab\x59\xcc\x78\x67\x67\x51\x94\x80\xc7\x42\x75\x35\xef\x43\xe1\xbd\x8f\xef\xd6\x70\xef\x3e\xd3\x7c\xcb\x58\xac\x2c\xca\x30\x19\xe8\x17\x17\x59\x94\x1e\x74\x0f\x6a\x7b\xe5\x05\xff\x94\xa9\xb6\xe3\x06\x82\x99\x05\x37\xe5\x2b\x91\xf5\x9b\x6f\x78\x30\xc0\x34\x98\x47\xdd\x5d\x83\x20\x73\xd3\x85\xdd\x8a\xc2\x16\x8f\xbb\xea\x30\x3c\x69\x6c\x3d\x0d\x53\x96\xa2\x7e\xf1\x34\x92\xe3\x78\x45\x65\xc6\x94\x6c\x48\x31\x84\x84\x86\x29\x94\x85\x3b\x8e\x56\x05\x0d\xf4\x12\xaf\x95\xdc\x7e\xfb\x24\xd8\x56\xb3\x10\x88\x1b\x43\xd7\xca\x35\xcc\xf2\x04\xf9\x78\x52\x86\xdb\xf0\xe5\x23\xb7\x33\xad\x2c\xff\x64\xa4\xe3\xa6\x04\xc3\x3f\xc3\x9b\xf8\xe5\xf3\x1c\x4f\x00\x0f\x5e\xd7\x30\x31\xec\x0e\xa4\x12\x3a\x4d\x5a\x12\xc5\x38\x76\x91\x21\x31\xfc\x73\xaa\x15\xa1\x74\x53\x27\x3d\xb4\x57\x48\xcc\xa1\x0c\x49\x60\xf5\xe7\xf6\x1a\x3d\x06\xef\x17\xda\x9d\xe8\xb9\x6a\xe9\x9a\xd5\x72\x75\xa5\xba\x69\xae\xea\x1a\x9a\x29\x6f\xee\x41\x8a\x51\xaf\x5b\x29\x24\x6f\x47\x56\xaa\x86\xc3\xd4\x2b\x5c\x09\x77\x4c\x76\x20\x05\xa0\xca\x05\xfd\x90\x02\xdc\xb2\xe8\xfe\x6c\xbf\x64\xc6\xf2\x90\x08\x0e\xfb\x89\x36\x3d\x73\xbe\x32\xd5\xa9\x87\xa9\xfe\xe0\x8e\x14\x67\x62\x74\x9e\x82\x5c\x61\x90\x82\xd2\x77\x41\xfd\x42\x53\x5e\xbf\x8e\x77\x90\xb8\x6b\xaa\x3f\x95\x5c\x10\x0a\xbf\x8d\xc1\xa5\xdf\xf1\xec\xb0\x95\xaf\x7f\x80\xdf\xa8\x41\x0a\x42\x71\xf3\xbf\xdd\xfd\xf9\x05\xde\x67\x62\x74\xa1\x15\x1f\x9d\x33\xd4\xb7\x6f\x70\x06\xa6\x5a\xe8\xbd\x05\x6e\xd0\x50\x82\x2d\xe4\x56\x48\x05\x45\x4a\x81\x8b\xbf\xf2\xfc\x5f\x59\xfb\x7d\xe8\xbd\x6d\x1e\x00\x23\x16\xad\xae\x30\x8b\xa8\x4d\x23\xd4\xa3\xa2\x8c\x19\x1c\x0e\x55\x6f\x9b\xd7\x11\x73\x68\xed\x7b\x79\xa1\x1f\x09\xad\x42\x0b\x37\x5b\x4a\x69\x2a\x1e\x6b\xa6\xfc\xe5\x44\xd0\x62\x84\x91\x8d\xee\x8a\x12\x8a\xd9\xfc\xae\x93\x4d\x09\x4a\x8f\x9c\x61\xca\x0a\x6d\xfa\x12\x7a\xb6\x18\xb1\x09\x1f\xbf\xdd\xfb\xe5\xed\xaf\xbb\xbb\xbb\xc5\xf6\x04\x8f\x17\x33\x69\xb8\x5d\xcf\xf1\xa0\x6d\x31\x71\xb2\x57\xc2\x6e\x09\xbb\x2f\x24\xbd\x0d\x12\x8f\xa1\xe5\x38\x2e\x4b\x95\x06\xef\x87\xa0\x8e\xaf\xd9\x64\x89\x80\xed\x4e\x65\x7a\xc4\x66\x87\x9b\x89\xd3\x60\xa2\x70\xe4\x83\x49\x20\xc1\xe9\x28\x9c\xb3\xcf\x34\xea\x73\x52\xaf\x35\x69\x0e\x2a\x9c\x04\x3a\xca\xf0\xea\xbe\x0a\x83\xfb\x69\x78\xe0\xd4\x45\x4d\x4b\xda\x98\xbc\x63\x9c\x8f\xfc\x81\x1b\xcb\xb7\x86\x33\xe1\xdb\xf6\xb0\x1b\x8e\xdf\x19\xdd\xf0\x87\x0d\x82\x97\xa8\xf0\x5f\x93\x5c\xde\xfc\x14\xeb\xf9\x06\x53\x74\x21\xb8\xbe\x3c\x56\xfe\x93\xed\x0d\x9a\xdf\xae\x73\xfe\x2a\xf2\xfa\xc9\xb7\x11\x7c\xe0\xf4\x9d\xb1\x0d\x7f\x58\x86\xfd\x37\x00\x00\xff\xff\xfd\xea\x46\xbb\x35\x0e\x00\x00"),
		},
		"/manifest.go.extra": &vfsgen۰FileInfo{
			name:    "manifest.go.extra",
			modTime: time.Date(2019, 1, 4, 2, 42, 23, 581049750, time.UTC),
			content: []byte("\x76\x61\x72\x20\x28\x0a\x09\x76\x66\x73\x67\x65\x6e\xdb\xb0\x6d\x61\x6e\x69\x66\x65\x73\x74\x20\x3d\x20\x25\x23\x76\x0a\x09\x76\x66\x73\x67\x65\x6e\xdb\xb0\x72\x65\x76\x20\x20\x20\x20\x20\x20\x3d\x20\x25\x23\x76\x0a\x29\x0a"),
		},
		"/package.json": &vfsgen۰CompressedFileInfo{
			name:             "package.json",
			modTime:          time.Date(2019, 1, 4, 2, 42, 23, 581076428, time.UTC),
			uncompressedSize: 174,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\x31\xcb\xc2\x30\x10\xc6\xf1\x3d\x9f\xe2\x38\xc8\xf6\xae\xef\xd2\xc1\xc5\x76\x10\xa4\x8b\x38\x89\x43\xbd\x1c\x78\x50\x93\x78\x17\x75\x28\xfd\xee\x92\x08\x3a\x3e\xcf\xef\xbf\x38\x00\x8c\xd3\x8d\xb1\x03\x7f\xff\xab\x2b\xb0\x91\x4a\x2e\x92\xe2\xef\x9c\x85\x38\x5a\xad\xf0\x38\xee\x77\xdb\x61\x3c\x0c\x3d\x36\xca\x2a\xcf\xa9\x54\x2a\xfa\xe0\x76\x5d\x34\xbd\x8c\xd5\x66\xb1\x82\x1d\x9c\x1c\x00\x00\x6e\xe0\xdf\x7b\x74\x00\xe7\x16\xd1\x44\x57\xee\x45\x99\x4a\x52\x61\xab\xa1\xb7\x2f\x07\xce\x1c\x03\x47\xfa\xd0\xb2\xba\xd5\xbd\x03\x00\x00\xff\xff\x18\x68\x1c\x80\xae\x00\x00\x00"),
		},
		"/sass.js": &vfsgen۰CompressedFileInfo{
			name:             "sass.js",
			modTime:          time.Date(2019, 1, 4, 2, 42, 23, 581118051, time.UTC),
			uncompressedSize: 1679,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5d\x6f\xe2\x3a\x10\x7d\xcf\xaf\x98\xbe\x34\x46\xa5\xa6\x7d\x25\x97\x87\x0a\x55\x57\xba\x57\x6a\xa5\x65\x3f\x1e\x58\xb4\xf2\x26\x13\xb0\xd6\xd8\xe9\xd8\xe9\x16\xa1\xfc\xf7\xd5\xc4\x09\x84\xd2\xaa\xda\x17\x42\x3c\x67\x66\xce\x9c\x39\xf1\xb3\x22\xb0\x18\x60\x06\x84\x4f\xb5\x26\x14\xa9\xc5\x90\x8e\xb2\x84\x23\x5e\x79\x7f\x12\x72\x05\x5e\xf3\x61\x0f\x28\x50\xf9\x9d\xcd\x87\x98\xee\x88\x11\x89\x2e\x41\x5c\x54\xe4\x72\xf4\x5e\xa2\x7d\x96\x77\x8b\xc5\xfd\xe7\x7f\xef\x1f\x7e\x2c\x1e\xe7\xff\x8f\x60\x9f\x00\xe4\xce\x7a\x67\x50\x22\x91\x23\x91\xb6\x8f\x69\x3a\x86\xf4\x04\x0b\xda\x83\x75\x01\x7c\x24\x07\x70\xa8\xfa\xa2\x83\xb8\x1d\x65\x49\x93\x24\x93\x09\x14\xee\x13\x3e\x81\x47\x5b\x78\x50\x2d\x29\xf4\x01\x82\x83\xb0\x41\xd0\x55\x0e\xb9\x32\xe6\xa7\xca\x7f\x81\x47\x7a\x46\x92\x49\x59\xdb\x3c\x68\x67\x63\xaa\xa8\x14\xa9\xad\x1f\x1f\x70\x91\x24\xcf\x9a\x1b\x98\xb1\x56\x32\x27\x54\x01\xe7\xce\x5a\x6c\x33\xc5\xfb\x13\xf2\x7c\x46\x3a\x2b\xd2\x3c\xc2\xd3\x31\xf4\x0d\x45\x2c\x0d\x30\x99\x40\x17\x6d\x5f\x73\x23\x7f\x93\x0e\x28\xfe\x5b\x3c\x3e\x48\x1f\x48\xdb\xb5\x2e\x77\x1d\xb3\x11\x5c\x41\xfa\xdd\x46\x11\x9a\xf6\xb7\xeb\x50\xa8\xa0\x86\xe5\xf9\x7d\xd0\x82\x5f\xdb\xff\x3c\x4b\x05\x33\x68\xcb\x57\x8a\x3c\x46\x68\xd6\x46\x79\x65\x15\x5c\x5e\x42\x15\x37\xd2\x57\x78\x7f\x51\x3d\x30\xeb\x70\x67\x9b\xe1\xc3\x06\xd0\x78\x3c\x56\xbf\xa8\x24\xa1\xaf\x4d\xf8\xa0\x3e\x6c\xb5\xf7\xda\xae\x21\xa2\xd3\x0f\xba\x44\x01\xbb\xd5\x89\x43\x8f\xac\x17\xb6\x40\x1f\xc8\xed\xc4\xb9\x7a\x6d\xbb\xa1\x7c\xd8\x33\x7b\x9d\xf6\xbe\x12\xd8\xc5\xdf\xe0\xd6\x0c\x1c\x3a\x57\xc6\x40\x85\x54\x3a\xda\xb2\x4b\x99\xef\x87\xde\xe4\x24\x51\xda\x48\x8a\x30\xd4\x64\xcf\x9d\xc4\xab\x55\xb4\xe6\x6f\x76\xb9\x8a\x5c\x4a\x47\x20\xf8\x5c\xc3\x0c\x6e\x32\xd0\xf0\x0f\x43\xea\x2d\xda\xe0\xa5\x41\xbb\x0e\x1b\xb8\x86\xdb\x0c\xf4\xd5\xd5\x71\x19\x5c\x65\xa9\x57\x30\x3b\x82\x97\x7a\x25\xd7\x18\xbe\x2a\x53\xa3\x38\x08\x7e\xe8\x5b\x38\x8b\x30\x83\x52\x19\x8f\xd9\xe1\x94\xda\xbb\xc5\xd6\xc6\xc4\xb3\xf8\x91\xed\xc3\xae\x42\x98\x42\xca\x13\xb3\x85\x5a\x6b\xc3\x14\xf6\x56\x6d\x39\x50\xda\x71\x1c\x64\xda\x3e\x9a\x66\xb0\x98\x81\x25\x63\x75\xea\x3d\xd1\x51\x08\x54\x77\x0c\x9a\x8e\x66\x77\x27\x49\xe3\x5c\xf5\x6d\xa3\x0d\x8a\xa1\x74\xbd\x9c\x17\x9c\x9f\xb5\x49\x6d\x56\xb0\x5c\x6c\x57\xa1\x2b\x19\x12\x4b\x75\x58\xbe\x04\x25\xc7\xfc\x32\x58\x99\x6f\x14\xdd\x05\x71\x33\x92\xc1\x7d\xa9\x2a\xa4\xb9\xf2\x28\xf8\x53\x0d\x56\x7a\xa3\x73\x14\xb7\xa3\x95\x20\x8c\x56\x6c\x7a\x2f\xac\x31\x1c\x97\x6e\xb4\x0f\xc9\xb9\x92\xaf\x14\x63\xd4\x75\xcf\xde\xa7\x6f\x08\xb3\x75\x45\xcd\xe6\x7c\xa9\x1c\x05\xb6\xc2\xbe\xe1\xa6\x6f\x19\x81\x3a\x03\x9c\x2c\x9f\x31\x25\x8f\x4e\x4b\xdd\x99\xe8\xb4\xe4\xb2\xb4\x6c\x8c\xa3\x27\xdb\x99\x92\x57\xfa\xb3\x8c\x7f\xad\xfb\x9f\x00\x00\x00\xff\xff\xa9\x7c\x25\x5f\x8f\x06\x00\x00"),
		},
		"/yarnrc": &vfsgen۰FileInfo{
			name:    "yarnrc",
			modTime: time.Date(2019, 1, 4, 2, 42, 23, 581141257, time.UTC),
			content: []byte("\x2d\x2d\x6d\x6f\x64\x75\x6c\x65\x73\x2d\x66\x6f\x6c\x64\x65\x72\x20\x25\x71\x0a\x2d\x2d\x2a\x2e\x6e\x6f\x2d\x62\x69\x6e\x2d\x6c\x69\x6e\x6b\x73\x20\x74\x72\x75\x65\x0a"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/assets.anko"].(os.FileInfo),
		fs["/gitignore"].(os.FileInfo),
		fs["/manifest.go"].(os.FileInfo),
		fs["/manifest.go.extra"].(os.FileInfo),
		fs["/package.json"].(os.FileInfo),
		fs["/sass.js"].(os.FileInfo),
		fs["/yarnrc"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
