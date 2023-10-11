package models

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
)

var (
	ErrEmailTaken = errors.New("models: email address is already in use")
)

var (
	ErrNotFound = errors.New("models: resource could not be found")
)

type FileError struct {
	Issue string
}

func (fe FileError) Error() string {
	return fmt.Sprintf("invalid file: %v", fe.Issue)
}

func checkContentType(r io.ReadSeeker, allowedTypes []string) error {
	test := make([]byte, 512)
	_, err := r.Read(test)
	if err != nil {
		return fmt.Errorf("checking content type: %w", err)
	}
	_, err = r.Seek(io.SeekStart, io.SeekStart)
	if err != nil {
		return fmt.Errorf("checking content type: %w", err)
	}
	contentType := http.DetectContentType(test)
	for _, t := range allowedTypes {
		if contentType == t {
			return nil
		}
	}

	return FileError{
		Issue: fmt.Sprintf("invalid content type: %v", contentType),
	}
}

func checkExtension(filename string, allowedExtensions []string) error {
	if !hasExtension(filename, allowedExtensions) {
		return FileError{
			Issue: fmt.Sprintf("invalid extension: %v", filepath.Ext(filename)),
		}
	}

	return nil
}
