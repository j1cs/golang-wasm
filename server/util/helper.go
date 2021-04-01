package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/j1cs/golang-wasm/server/logger"
)

// GettingWasmJS grab file
func GettingWasmJS() {
	logger := logger.GetLogger()
	path, _ := os.Getwd()
	file := path + filepath.FromSlash("/public/static/wasm_exec.js")
	wasm := runtime.GOROOT() + filepath.FromSlash("/misc/wasm/wasm_exec.js")
	if _, err := os.Stat(file); err == nil {
		return
	}
	err := copyFile(wasm, file)
	if err != nil {
		logger.Panicln(err)
	}
}

func copyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
