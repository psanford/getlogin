// Package getlogin provides functionality similar to libc's getlogin(3).

// Similar to getlogin(3), this package should NOT be used for any security related checks.
package getlogin

import (
	"errors"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// Returns the name of the user logged in on the terminal connected to the process' stdin or failing that from the environment.
// This function can be easily fooled and should NOT be used for any security related purposes.
func GetLogin() string {
	if u, err := UserFromStdin(); err == nil {
		return u.Username
	}

	return LoginFromEnv()
}

// Returns the user that owns the Stdin TTY if Stdin is a TTY.
// This function should NOT be used for any security related purposes.
func UserFromStdin() (*user.User, error) {
	if !isTTY(os.Stdin) {
		return nil, errors.New("stdin is not a tty")
	}

	tty, err := getTTY(os.Stdin)
	if err != nil {
		return nil, errors.New("could not determine tty for stdin")
	}

	statT, ok := tty.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, errors.New("could not get syscall.Stat_t for tty")
	}

	return user.LookupId(strconv.Itoa(int(statT.Uid)))
}

// Returns username from environment variable LOGNAME.
// This function should NOT be used for any security related purposes.
func LoginFromEnv() string {
	return os.Getenv("LOGNAME")
}
