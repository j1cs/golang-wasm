package handler

import (
	"net/http"
	"strings"
	"sync/atomic"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Init handle options for methods
type Init struct {
	Directory string
	Health    *int32
}

// GetIndex it will display index.html when it requested
func (i Init) GetIndex() http.Handler {
	return http.FileServer(FileSystem{http.Dir(i.Directory)})
}

// GetStatic it will display the statics when it requested
// Im using a custom http.FileSystem to avoid display the whole folder
func (i Init) GetStatic() http.Handler {
	return http.StripPrefix("/", i.GetIndex())
}

// GetHealt check the healt of the server
func (i Init) GetHealt() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(i.Health) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
