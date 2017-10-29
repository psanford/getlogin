// +build darwin dragonfly freebsd netbsd openbsd solaris

package getlogin

import (
	"errors"
	"io/ioutil"
	"os"
	"syscall"
)

var ttyPaths = []string{
	"/dev/vc/",
	"/dev/tts/",
	"/dev/pty/",
	"/dev/pts/",
	"/dev/",
}

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

	var match os.FileInfo
OUTER:
	for _, dir := range ttyPaths {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, f := range files {
			dstatT, ok := f.Sys().(*syscall.Stat_t)
			if !ok {
				continue
			}

			if dstatT.Dev == fstatT.Dev && dstatT.Ino == fstatT.Ino {
				match = f
				break OUTER
			}
		}
	}

	if match == nil {
		return nil, errors.New("Could not determine tty for " + f.Name())
	}

	return match, nil
}
