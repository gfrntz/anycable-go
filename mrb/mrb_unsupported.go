// +build !darwin,!linux !cgo

package mrb

// Supported returns true iff mruby scripting is available
func Supported() bool {
	return false
}
