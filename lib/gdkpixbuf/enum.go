// This is a generated file - DO NOT EDIT

package gdkpixbuf

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

type Colorspace C.GdkColorspace

const (
	GDK_COLORSPACE_RGB Colorspace = 0
)

type InterpType C.GdkInterpType

const (
	GDK_INTERP_NEAREST  InterpType = 0
	GDK_INTERP_TILES    InterpType = 1
	GDK_INTERP_BILINEAR InterpType = 2
	GDK_INTERP_HYPER    InterpType = 3
)

type PixbufAlphaMode C.GdkPixbufAlphaMode

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = 0
	GDK_PIXBUF_ALPHA_FULL    PixbufAlphaMode = 1
)

type PixbufError C.GdkPixbufError

const (
	GDK_PIXBUF_ERROR_CORRUPT_IMAGE         PixbufError = 0
	GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY   PixbufError = 1
	GDK_PIXBUF_ERROR_BAD_OPTION            PixbufError = 2
	GDK_PIXBUF_ERROR_UNKNOWN_TYPE          PixbufError = 3
	GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION PixbufError = 4
	GDK_PIXBUF_ERROR_FAILED                PixbufError = 5
	GDK_PIXBUF_ERROR_INCOMPLETE_ANIMATION  PixbufError = 6
)

type PixbufRotation C.GdkPixbufRotation

const (
	GDK_PIXBUF_ROTATE_NONE             PixbufRotation = 0
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	GDK_PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	GDK_PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)