//go:build windows || plan9 || wasip1

package copy

import "io/fs"

func preserveOwner(src, dest string, info fs.FileInfo) (err error) {
	return nil
}
