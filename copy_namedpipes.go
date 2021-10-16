//go:build !windows
// +build !windows

package copy

import (
	"os"
	"path/filepath"
	"syscall"
)

// pcopy is for just named pipes
func pcopy(src, dest string, info os.FileInfo, opt Options) error {
	overwriteMode := opt.OnDestExists(src, dest)
	if overwriteMode == Merge || overwriteMode == OverwriteIntersection {
		// Check if destination exists
		_, err := os.Lstat(dest)
		exists := err == nil
		if overwriteMode == Merge {
			// Skip copying if dest exists
			if exists {
				return nil
			}
		} else { // OverwriteIntersection
			// Remove existing fs item if exists
			if exists {
				if err := os.Remove(dest); err != nil {
					return err
				}
			}
		}
	}
	if err := os.MkdirAll(filepath.Dir(dest), os.ModePerm); err != nil {
		return err
	}
	return syscall.Mkfifo(dest, uint32(info.Mode()))
}
