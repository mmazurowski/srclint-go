package crawler

import (
	"io/fs"
	"path/filepath"
	"strings"

	"srclint/src/errors"
	"srclint/src/utils-arrays"
)

func Crawl(checkedDirectory string, ignored []string) ([]string, error) {
	var dirs []string

	var crawlErr = filepath.WalkDir(checkedDirectory, func(path string, d fs.DirEntry, err error) error {
		var sliced = strings.Replace(path, checkedDirectory+"/", "", 1)

		if d.IsDir() {
			return nil
		}

		if utils_arrays.Contains(ignored, sliced) {
			return nil
		}

		dirs = append(dirs, sliced)
		return nil
	})

	if crawlErr != nil {
		return nil, errors.CreateConfigParserError("ASDASDASD")
	}

	return dirs, nil
}
