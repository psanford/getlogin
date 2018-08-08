// +build windows

package getlogin

import (
	"errors"
	"os"
	"os/user"
)

// Returns the name of the user from the environment.
// This function can be easily fooled and should NOT be used for any security related purposes.
func GetLogin() string {
	return LoginFromEnv()
}

// Not supported on windows.
func UserFromStdin() (*user.User, error) {
	return nil, errors.New("not supported on windows")
}

// Returns username from environment variable USERNAME.
// This function should NOT be used for any security related purposes.
func LoginFromEnv() string {
	return os.Getenv("USERNAME")
}
