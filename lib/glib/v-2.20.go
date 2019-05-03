// Code generated - DO NOT EDIT.
// +build glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_base64_decode_inplace : unsupported parameter out_len : array length param out_len is pointer (gsize*)

// Poll is a wrapper around the C function g_poll.
func Poll(fds *PollFD, nfds uint32, timeout int32) int32 {
	c_fds := (*C.GPollFD)(C.NULL)
	if fds != nil {
		c_fds = (*C.GPollFD)(fds.ToC())
	}

	c_nfds := (C.guint)(nfds)

	c_timeout := (C.gint)(timeout)

	retC := C.g_poll(c_fds, c_nfds, c_timeout)
	retGo := (int32)(retC)

	return retGo
}