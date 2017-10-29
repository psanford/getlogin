// +build solaris

package getlogin

import (
	"os"

	"golang.org/x/sys/unix"
)

func isTTY(f *os.File) bool {
	_, err := unix.IoctlGetTermio(int(f.Fd()), unix.TCGETA)
	return err == nil
}
