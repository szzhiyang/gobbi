// This is a generated file - DO NOT EDIT
// +build glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

/*
This is the platform dependent conversion specifier for scanning
and printing values of type #gintptr.
*/
const GINTPTR_FORMAT string = C.G_GINTPTR_FORMAT

/*
The platform dependent length modifier for conversion specifiers
for scanning and printing values of type #gintptr or #guintptr.
It is a string literal.
*/
const GINTPTR_MODIFIER string = C.G_GINTPTR_MODIFIER

/*
This is the platform dependent conversion specifier
for scanning and printing values of type #guintptr.
*/
const GUINTPTR_FORMAT string = C.G_GUINTPTR_FORMAT
