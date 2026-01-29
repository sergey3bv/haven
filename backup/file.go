package backup

import (
	"fmt"
	"os"
)

func fileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	if !info.Mode().IsRegular() {
		return -1, fmt.Errorf("path is not a file")
	}


	return info.Size(), nil
}
