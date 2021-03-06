/*
MIT License

Copyright (c) 2018 Danny van Kooten

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package echojet

import (
	"bytes"
	"io"
	"path/filepath"
)

type BinLoader struct {
	// Path to directory that contains your Jet templates
	Root string

	// Asset func generated by go-bindata
	AssetFunc func(name string) ([]byte, error)
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

// Open opens the underlying reader with template content.
func (l *BinLoader) Open(name string) (io.ReadCloser, error) {
	b, err := l.AssetFunc(filepath.Join(l.Root, name))
	return nopCloser{bytes.NewBuffer(b)}, err
}

// Exists checks for template existence
func (l *BinLoader) Exists(name string) (string, bool) {
	_, err := l.AssetFunc(filepath.Join(l.Root, name))
	if err != nil {
		return "", false
	}

	return name, true
}
