// +build darwin dragonfly freebsd netbsd openbsd

package getlogin

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA
