package utils

import "unsafe"

// GetString returns a string pointer without allocation
// #nosec G103
func GetString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
