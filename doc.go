// Package rwfs provides interfaces that represent writable filesystems. It
// builds on top of the standard library types such as [fs.FS] and [fs.File],
// acting as an extension to enable writing on top of them.
//
// It includes a concrete implementation for the OS filesystem.
package rwfs
