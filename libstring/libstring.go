// Package libstring provides string related library functions.
package libstring

import (
	"crypto/rand"
	"encoding/base64"
	"net"
	"os"
	"os/user"
	"strings"
)

// ExpandTilde is a convenience function that expands ~ to full path.
func ExpandTilde(path string) string {
	if path == "" {
		return path
	}

	if path[:2] == "~/" {
		usr, err := user.Current()
		if err != nil || usr == nil {
			return path
		}

		if usr.Name == "root" {
			path = strings.Replace(path, "~", "/root", 1)
		} else {
			path = strings.Replace(path, "~", usr.HomeDir, 1)
		}

	}
	return path
}

// ExpandTilde is a convenience function that expands both ~ and $ENV.
func ExpandTildeAndEnv(path string) string {
	path = ExpandTilde(path)
	return os.ExpandEnv(path)
}

// GeneratePassword returns password.
// size determines length of initial seed bytes.
func GeneratePassword(size int) (string, error) {
	// Force minimum size to 32
	if size < 32 {
		size = 32
	}

	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}

// StringInSlice search exact match in a slice of strings.
func StringInSlice(beingSearched string, list []string) bool {
	for _, b := range list {
		if b == beingSearched {
			return true
		}
	}
	return false
}

// Split r.RemoteAddr, return an IP object (or nil if ParseIP fails)
func GetIP(address string) net.IP {
	// Try to parse it
	splitAddress := strings.Split(address, ":")
	if len(splitAddress) == 0 {
		return nil
	}

	// Convert to IP object
	return net.ParseIP(splitAddress[0])
}

func FindHostnameChunkInMetricKey(key string) int {
	chunks := strings.Split(key, ".")

	for i, chunk := range chunks {
		if strings.Contains(chunk, "localhost") {
			return i
		}
		if strings.HasSuffix(chunk, "_com") || strings.Contains(chunk, "_net") || strings.Contains(chunk, "_org") || strings.Contains(chunk, "_edu") || strings.Contains(chunk, "_gov") || strings.Contains(chunk, "_mil") {
			return i
		}
	}

	return -1
}
