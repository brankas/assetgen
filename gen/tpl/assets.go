package assets

// Code generated by assetgen. DO NOT EDIT.

import (
	"context"
	"crypto/sha1"
	"embed"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"path"
	"strings"
	"time"
)

// Files is the embedded assets.
//
%s
var Files embed.FS

const (
	// DistPath is the dist path used when building the files.
	DistPath = %q
	// ManifestFile is the name of the manifest file.
	ManifestFile = %q
)

// Asset wraps an asset.
type Asset struct {
	Hash        string
	ModTime     time.Time
	ContentType string
	Content     []byte
}

// Manifest returns a map of the asset names.
func Manifest() (map[string]string, error) {
	buf, err := Files.ReadFile(path.Join(DistPath, ManifestFile))
	if err != nil {
		return nil, err
	}
	var manifest map[string]string
	if err := json.Unmarshal(buf, &manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

// Assets returns a map of the asset contents.
func Assets() (map[string]*Asset, error) {
	modTime := time.Now()
	manifest, err := Manifest()
	if err != nil {
		return nil, err
	}
	assets := make(map[string]*Asset, len(manifest)-1)
	for k, n := range manifest {
		content, err := Files.ReadFile(path.Join(DistPath, n))
		if err != nil {
			return nil, err
		}
		hash := fmt.Sprintf("%%x", sha1.Sum(content))
		contentType := http.DetectContentType(content)
		switch {
		case strings.HasPrefix(contentType, "text/") || contentType == "":
			if i := strings.LastIndex(n, "."); i != -1 {
				contentType = mime.TypeByExtension(n[i:])
			}
		}
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		assets[k] = &Asset{
			Hash:        hash,
			ModTime:     modTime,
			ContentType: contentType,
			Content:     content,
		}
	}
	return assets, nil
}

// ManifestPath returns a manifest path conversion func.
func ManifestPath(prefixes ...string) func(string) string {
	manifest, err := Manifest()
	if err != nil {
		panic(err)
	}
	rev := make(map[string]string, len(manifest))
	for n, k := range manifest {
		rev[k] = n
	}
	prefix := path.Join(prefixes...)
	return func(s string) string {
		return path.Join(prefix, rev["/"+strings.TrimPrefix(s, "/")])
	}
}

// StaticHandler returns a static asset handler.
func StaticHandler(f func(context.Context) string) http.Handler {
	if f == nil {
		panic("f cannot be nil")
	}
	assets, err := Assets()
	if err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// retrieve asset
		asset, ok := assets[strings.TrimPrefix(f(req.Context()), "/")]
		if !ok {
			http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		// check if-modified-since header, bail if present
		if t, err := time.Parse(http.TimeFormat, req.Header.Get("If-Modified-Since")); err == nil && asset.ModTime.Unix() <= t.Unix() {
			res.WriteHeader(http.StatusNotModified) // 304
			return
		}
		// check If-None-Match header, bail if present and match hash
		if req.Header.Get("If-None-Match") == asset.Hash {
			res.WriteHeader(http.StatusNotModified) // 304
			return
		}
		// set headers
		res.Header().Set("Content-Type", asset.ContentType)
		res.Header().Set("Date", time.Now().Format(http.TimeFormat))
		// cache headers
		res.Header().Set("Cache-Control", "public, no-transform, max-age=31536000")
		res.Header().Set("Expires", time.Now().AddDate(1, 0, 0).Format(http.TimeFormat))
		res.Header().Set("Last-Modified", asset.ModTime.Format(http.TimeFormat))
		res.Header().Set("ETag", asset.Hash)
		// write data to response
		_, _ = res.Write(asset.Content)
	})
}
