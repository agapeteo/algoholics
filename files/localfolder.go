package files

import (
	"os"
	"path/filepath"
)

func LocalFolder() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
