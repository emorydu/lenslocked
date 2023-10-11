package fileutil

import (
	"archive/zip"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// FileType uses the filetype package to determine the given file path's type.
func FileType(filepath string) (types.Type, error) {
	f, _ := os.Open(filepath)

	// We only have to pass the file header = first 262 bytes
	head := make([]byte, 261)
	_, _ = f.Read(head)

	return filetype.Match(head)
}

// FileExists returns true if the given path exists.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	return false, err
}

// DirExists returns true if the given path exists and is a directory.
func DirExists(path string) (bool, error) {
	exists, _ := FileExists(path)
	info, _ := os.Stat(path)
	if !exists || !info.IsDir() {
		return false, fmt.Errorf("path eigher doesn't exist, or is not a directory <%s>", path)
	}

	return true, nil
}

// Touch creates an empty file at the given path if it doesn't already exist.
func Touch(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return nil
		}
		defer file.Close()
	}

	return nil
}

// EnsureDir will create a directory at the given path if it doesn't already exist.
func EnsureDir(path string) error {
	exists, err := FileExists(path)
	if !exists {
		return os.Mkdir(path, 0755)
	}

	return err
}

// EnsureDirAll will create a directory at the given path along with any necessary parents if they don't already exist.
func EnsureDirAll(path string) error {
	return os.MkdirAll(path, 0755)
}

// RemoveDir removes the given dir (if it exists) along with all of its contents.
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// EmptyDir will recursively remove the contents of a directory at the given path.
func EmptyDir(path string) error {
	d, err := os.Open(path)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}

	return nil
}

// ListDir will return the contents of a given directory path as a string slice.
func ListDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		path = filepath.Dir(path)
		files, _ = os.ReadDir(path)
	}

	var dirPaths []string
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		dirPaths = append(dirPaths, filepath.Join(path, file.Name()))
	}

	return dirPaths
}

// SafeMove move src to dst in safe mode.
func SafeMove(src, dst string) error {
	err := os.Rename(src, dst)

	//nolint:nestif
	if err != nil {
		log.Printf("[fileutil] unable to rename: \"%s\" due to %s. Falling back to copying.", src, err.Error())

		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}

		err = os.Remove(src)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsZipFileUncompressed returns true if zip file in path is using 0 compression level.
func IsZipFileUncompressed(path string) (bool, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Printf("Error reading zip file %s: %s\n", path, err.Error())
		return false, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.FileInfo().IsDir() { // skip dirs, they always get store level compression.
			continue
		}
		return f.Method == 0, nil // check compression level of first actual file.
	}

	return false, nil
}

// WriteFile writes file to path creating parent directories if needed.
func WriteFile(path string, file []byte) error {
	pathErr := EnsureDirAll(filepath.Dir(path))
	if pathErr != nil {
		return fmt.Errorf("cannot ensure path %s", pathErr.Error())
	}

	err := os.WriteFile(path, file, 0600)
	if err != nil {
		return fmt.Errorf("write error for thumbnail %s: %s", path, err.Error())
	}

	return nil
}

// GetIntraDir returns a string that can be added to filepath.Join to implement directory depth, "" on error
// eg given a pattern of 0af63ce3c99162e9df23a997f62621c5 and a depth of 2 length of 3
// returns 0af/63c or 0af\63c ( dependin on os)  that can be later used like this  filepath.Join(directory, intradir,
// basename).
func GetIntraDir(pattern string, depth, length int) string {
	if depth < 1 || length < 1 || (depth*length > len(pattern)) {
		return ""
	}
	intraDir := pattern[0:length] // depth 1, get length number of characters from pattern.
	for i := 1; i < depth; i++ {
		intraDir = filepath.Join(intraDir, pattern[length*i:length*(i+1)])
	} // adding each time to intraDir the extra characters with a filepath join

	return intraDir
}

// GetParent returns the parent directory of the given path.
func GetParent(path string) *string {
	isRoot := path[len(path)-1:] == "/"
	if isRoot {
		return nil
	}
	parentPath := filepath.Clean(path + "/..")

	return &parentPath
}

// ServeFileNoCache serves the provided file, ensuring that the response
// contains headers to prevent caching.
func ServeFileNoCache(w http.ResponseWriter, r *http.Request, filepath string) {
	w.Header().Add("Cache-Control", "no-cache")

	http.ServeFile(w, r, filepath)
}

// MatchEntries returns a string slice of the entries in directory dir which
// match the regexp pattern. On error an empty slice is returned
// MatchEntries isn't recursive, only the specific 'dir' is searched
// without being expanded.
func MatchEntries(dir, pattern string) ([]string, error) {
	var (
		res []string
		err error
	)

	regx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if regx.Match([]byte(file)) {
			res = append(res, filepath.Join(dir, file))
		}
	}

	return res, err
}
