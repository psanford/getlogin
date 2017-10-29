// +build linux

package getlogin

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

const ioctlReadTermios = unix.TCGETS

func getTTY(f *os.File) (os.FileInfo, error) {
	if f == nil {
		return nil, errors.New("no such file")
	}

	fstat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fstatT, ok := fstat.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("Could not get syscall.Stat_t for " + f.Name())
	}

	link, err := os.Readlink(filepath.Join("/proc/self/fd/", strconv.Itoa(int(f.Fd()))))
	if err != nil {
		return nil, err
	}

	stat, err := os.Stat(link)
	if err != nil {
		return nil, err
	}

	statT, ok := stat.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("Could not get syscall.Stat_t for " + link)
	}
	if fstatT.Dev != statT.Dev || fstatT.Ino != statT.Ino {
		return nil, errors.New("Could not determine tty")
	}

	return stat, nil
}
