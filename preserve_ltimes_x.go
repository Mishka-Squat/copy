//go:build windows || js || plan9 || wasip1

package copy

func preserveLtimes(src, dest string) error {
	return nil // Unsupported
}
