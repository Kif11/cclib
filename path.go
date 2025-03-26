package cclib

import "path/filepath"

// FileName returns the file name without extension from a full path or file name
func FileName(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)

	return base[0 : len(base)-len(ext)]
}
