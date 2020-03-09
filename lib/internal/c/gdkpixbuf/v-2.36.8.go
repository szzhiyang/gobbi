// Code generated - DO NOT EDIT.
// +build gdkpixbuf_2.36.8

package gdkpixbuf

import "unsafe"

// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <stdlib.h>
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type PixbufFormat C.GdkPixbufFormat
type PixbufLoaderClass C.GdkPixbufLoaderClass
type PixbufSimpleAnimClass C.GdkPixbufSimpleAnimClass

func Fn_gdk_pixbuf_format_copy(format unsafe.Pointer) unsafe.Pointer {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_copy(c_format)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_format_free(format unsafe.Pointer) {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	C.gdk_pixbuf_format_free(c_format)
}

func Fn_gdk_pixbuf_format_get_description(format unsafe.Pointer) string {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_get_description(c_format)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_extensions : no array length

func Fn_gdk_pixbuf_format_get_license(format unsafe.Pointer) string {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_get_license(c_format)

	return C.GoString(ret)
}

// UNSUPPORTED : gdk_pixbuf_format_get_mime_types : no array length

func Fn_gdk_pixbuf_format_get_name(format unsafe.Pointer) string {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_get_name(c_format)

	return C.GoString(ret)
}

func Fn_gdk_pixbuf_format_is_disabled(format unsafe.Pointer) bool {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_is_disabled(c_format)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_is_save_option_supported(format unsafe.Pointer, optionKey string) bool {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	c_optionKey := (*C.gchar)(C.CString(optionKey))
	defer C.free(unsafe.Pointer(c_optionKey))

	ret := C.gdk_pixbuf_format_is_save_option_supported(c_format, c_optionKey)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_is_scalable(format unsafe.Pointer) bool {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_is_scalable(c_format)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_is_writable(format unsafe.Pointer) bool {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	ret := C.gdk_pixbuf_format_is_writable(c_format)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_format_set_disabled(format unsafe.Pointer, disabled bool) {
	c_format := (*C.GdkPixbufFormat)(unsafe.Pointer(format))

	c_disabled := toCBool(disabled)

	C.gdk_pixbuf_format_set_disabled(c_format, c_disabled)
}

func Fn_gdk_pixbuf_new(colorspace int, hasAlpha bool, bitsPerSample int, width int, height int) unsafe.Pointer {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_hasAlpha := toCBool(hasAlpha)

	c_bitsPerSample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	ret := C.gdk_pixbuf_new(c_colorspace, c_hasAlpha, c_bitsPerSample, c_width, c_height)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_bytes(data unsafe.Pointer, colorspace int, hasAlpha bool, bitsPerSample int, width int, height int, rowstride int) unsafe.Pointer {
	c_data := (*C.GBytes)(unsafe.Pointer(data))

	c_colorspace := (C.GdkColorspace)(colorspace)

	c_hasAlpha := toCBool(hasAlpha)

	c_bitsPerSample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_rowstride := (C.int)(rowstride)

	ret := C.gdk_pixbuf_new_from_bytes(c_data, c_colorspace, c_hasAlpha, c_bitsPerSample, c_width, c_height, c_rowstride)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_data : parameter 'destroy_fn' is callback

func Fn_gdk_pixbuf_new_from_file(filename string, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file(c_filename, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_file_at_scale(filename string, width int, height int, preserveAspectRatio bool, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserveAspectRatio := toCBool(preserveAspectRatio)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file_at_scale(c_filename, c_width, c_height, c_preserveAspectRatio, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_file_at_size(filename string, width int, height int, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_file_at_size(c_filename, c_width, c_height, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_inline(dataLength int, data []uint8, copyPixels bool, error unsafe.Pointer) unsafe.Pointer {
	c_dataLength := (C.gint)(dataLength)

	c_data := (*C.guint8)(unsafe.Pointer(&data[0]))

	c_copyPixels := toCBool(copyPixels)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_inline(c_dataLength, c_data, c_copyPixels, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_resource(resourcePath string, error unsafe.Pointer) unsafe.Pointer {
	c_resourcePath := (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(c_resourcePath))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_resource(c_resourcePath, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_resource_at_scale(resourcePath string, width int, height int, preserveAspectRatio bool, error unsafe.Pointer) unsafe.Pointer {
	c_resourcePath := (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(c_resourcePath))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_preserveAspectRatio := toCBool(preserveAspectRatio)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_resource_at_scale(c_resourcePath, c_width, c_height, c_preserveAspectRatio, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_stream(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_stream(c_stream, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_stream_at_scale(stream unsafe.Pointer, width int, height int, preserveAspectRatio bool, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_preserveAspectRatio := toCBool(preserveAspectRatio)

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_stream_at_scale(c_stream, c_width, c_height, c_preserveAspectRatio, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_stream_finish(asyncResult unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_asyncResult := (*C.GAsyncResult)(unsafe.Pointer(asyncResult))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_new_from_stream_finish(c_asyncResult, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_new_from_xpm_data(data []string) unsafe.Pointer {
	dataLen := len(data)
	c_dataArray := C.malloc((C.ulong)(dataLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_dataArray))
	dataSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_dataArray))[:dataLen:dataLen]
	for datai, dataString := range data {
		dataSlice[datai] = (*C.gchar)(C.CString(dataString))
		defer C.free(unsafe.Pointer(dataSlice[datai]))
	}
	c_data := &dataSlice[0]

	ret := C.gdk_pixbuf_new_from_xpm_data(c_data)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_add_alpha(pixbuf unsafe.Pointer, substituteColor bool, r uint8, g uint8, b uint8) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_substituteColor := toCBool(substituteColor)

	c_r := (C.guchar)(r)

	c_g := (C.guchar)(g)

	c_b := (C.guchar)(b)

	ret := C.gdk_pixbuf_add_alpha(c_pixbuf, c_substituteColor, c_r, c_g, c_b)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_apply_embedded_orientation(src unsafe.Pointer) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	ret := C.gdk_pixbuf_apply_embedded_orientation(c_src)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_composite(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int, overallAlpha int) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	C.gdk_pixbuf_composite(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType, c_overallAlpha)
}

func Fn_gdk_pixbuf_composite_color(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int, overallAlpha int, checkX int, checkY int, checkSize int, color1 uint32, color2 uint32) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	c_checkX := (C.int)(checkX)

	c_checkY := (C.int)(checkY)

	c_checkSize := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	C.gdk_pixbuf_composite_color(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType, c_overallAlpha, c_checkX, c_checkY, c_checkSize, c_color1, c_color2)
}

func Fn_gdk_pixbuf_composite_color_simple(src unsafe.Pointer, destWidth int, destHeight int, interpType int, overallAlpha int, checkSize int, color1 uint32, color2 uint32) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_interpType := (C.GdkInterpType)(interpType)

	c_overallAlpha := (C.int)(overallAlpha)

	c_checkSize := (C.int)(checkSize)

	c_color1 := (C.guint32)(color1)

	c_color2 := (C.guint32)(color2)

	ret := C.gdk_pixbuf_composite_color_simple(c_src, c_destWidth, c_destHeight, c_interpType, c_overallAlpha, c_checkSize, c_color1, c_color2)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_copy(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_copy_area(srcPixbuf unsafe.Pointer, srcX int, srcY int, width int, height int, destPixbuf unsafe.Pointer, destX int, destY int) {
	c_srcPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(srcPixbuf))

	c_srcX := (C.int)(srcX)

	c_srcY := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	c_destPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(destPixbuf))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	C.gdk_pixbuf_copy_area(c_srcPixbuf, c_srcX, c_srcY, c_width, c_height, c_destPixbuf, c_destX, c_destY)
}

func Fn_gdk_pixbuf_copy_options(srcPixbuf unsafe.Pointer, destPixbuf unsafe.Pointer) bool {
	c_srcPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(srcPixbuf))

	c_destPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(destPixbuf))

	ret := C.gdk_pixbuf_copy_options(c_srcPixbuf, c_destPixbuf)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_fill(pixbuf unsafe.Pointer, pixel uint32) {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_pixel := (C.guint32)(pixel)

	C.gdk_pixbuf_fill(c_pixbuf, c_pixel)
}

func Fn_gdk_pixbuf_flip(src unsafe.Pointer, horizontal bool) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_horizontal := toCBool(horizontal)

	ret := C.gdk_pixbuf_flip(c_src, c_horizontal)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_get_bits_per_sample(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_bits_per_sample(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_byte_length(pixbuf unsafe.Pointer) uint64 {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_byte_length(c_pixbuf)

	return (uint64)(ret)
}

func Fn_gdk_pixbuf_get_colorspace(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_colorspace(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_has_alpha(pixbuf unsafe.Pointer) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_has_alpha(c_pixbuf)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_get_height(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_height(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_n_channels(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_n_channels(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_option(pixbuf unsafe.Pointer, key string) string {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.gdk_pixbuf_get_option(c_pixbuf, c_key)

	return C.GoString(ret)
}

func Fn_gdk_pixbuf_get_options(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_options(c_pixbuf)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_pixels : no array length

func Fn_gdk_pixbuf_get_pixels_with_length(pixbuf unsafe.Pointer, length *uint) []uint8 {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_length := (*C.guint)(unsafe.Pointer(length))

	ret := C.gdk_pixbuf_get_pixels_with_length(c_pixbuf, c_length)

	retLen := int(*c_length)
	retGo := make([]uint8, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](uint8))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gdk_pixbuf_get_rowstride(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_rowstride(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_width(pixbuf unsafe.Pointer) int {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_get_width(c_pixbuf)

	return (int)(ret)
}

func Fn_gdk_pixbuf_new_subpixbuf(srcPixbuf unsafe.Pointer, srcX int, srcY int, width int, height int) unsafe.Pointer {
	c_srcPixbuf := (*C.GdkPixbuf)(unsafe.Pointer(srcPixbuf))

	c_srcX := (C.int)(srcX)

	c_srcY := (C.int)(srcY)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	ret := C.gdk_pixbuf_new_subpixbuf(c_srcPixbuf, c_srcX, c_srcY, c_width, c_height)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_read_pixel_bytes(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_read_pixel_bytes(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_read_pixels(pixbuf unsafe.Pointer) *uint8 {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_read_pixels(c_pixbuf)

	return (*uint8)(ret)
}

func Fn_gdk_pixbuf_ref(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gdk_pixbuf_ref(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_remove_option(pixbuf unsafe.Pointer, key string) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	ret := C.gdk_pixbuf_remove_option(c_pixbuf, c_key)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_rotate_simple(src unsafe.Pointer, angle int) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_angle := (C.GdkPixbufRotation)(angle)

	ret := C.gdk_pixbuf_rotate_simple(c_src, c_angle)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_saturate_and_pixelate(src unsafe.Pointer, dest unsafe.Pointer, saturation float32, pixelate bool) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_saturation := (C.gfloat)(saturation)

	c_pixelate := toCBool(pixelate)

	C.gdk_pixbuf_saturate_and_pixelate(c_src, c_dest, c_saturation, c_pixelate)
}

// UNSUPPORTED : gdk_pixbuf_save : parameter 'error' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_buffer : blacklisted

// UNSUPPORTED : gdk_pixbuf_save_to_bufferv : parameter 'buffer' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_callback : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_callbackv : parameter 'save_func' is callback

// UNSUPPORTED : gdk_pixbuf_save_to_stream : parameter 'error' is non array with indirect count > 1

// UNSUPPORTED : gdk_pixbuf_save_to_stream_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_save_to_streamv(pixbuf unsafe.Pointer, stream unsafe.Pointer, type_ string, optionKeys []string, optionValues []string, cancellable unsafe.Pointer, error unsafe.Pointer) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_stream := (*C.GOutputStream)(unsafe.Pointer(stream))

	c_type_ := (*C.char)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	optionKeysLen := len(optionKeys)
	c_optionKeysArray := C.malloc((C.ulong)(optionKeysLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionKeysArray))
	optionKeysSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionKeysArray))[:optionKeysLen:optionKeysLen]
	for optionKeysi, optionKeysString := range optionKeys {
		optionKeysSlice[optionKeysi] = (*C.gchar)(C.CString(optionKeysString))
		defer C.free(unsafe.Pointer(optionKeysSlice[optionKeysi]))
	}
	c_optionKeys := &optionKeysSlice[0]

	optionValuesLen := len(optionValues)
	c_optionValuesArray := C.malloc((C.ulong)(optionValuesLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionValuesArray))
	optionValuesSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionValuesArray))[:optionValuesLen:optionValuesLen]
	for optionValuesi, optionValuesString := range optionValues {
		optionValuesSlice[optionValuesi] = (*C.gchar)(C.CString(optionValuesString))
		defer C.free(unsafe.Pointer(optionValuesSlice[optionValuesi]))
	}
	c_optionValues := &optionValuesSlice[0]

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_save_to_streamv(c_pixbuf, c_stream, c_type_, c_optionKeys, c_optionValues, c_cancellable, cError)

	return toGoBool(ret)
}

// UNSUPPORTED : gdk_pixbuf_save_to_streamv_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_savev(pixbuf unsafe.Pointer, filename string, type_ string, optionKeys []string, optionValues []string, error unsafe.Pointer) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	c_type_ := (*C.char)(C.CString(type_))
	defer C.free(unsafe.Pointer(c_type_))

	optionKeysLen := len(optionKeys)
	c_optionKeysArray := C.malloc((C.ulong)(optionKeysLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionKeysArray))
	optionKeysSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionKeysArray))[:optionKeysLen:optionKeysLen]
	for optionKeysi, optionKeysString := range optionKeys {
		optionKeysSlice[optionKeysi] = (*C.gchar)(C.CString(optionKeysString))
		defer C.free(unsafe.Pointer(optionKeysSlice[optionKeysi]))
	}
	c_optionKeys := &optionKeysSlice[0]

	optionValuesLen := len(optionValues)
	c_optionValuesArray := C.malloc((C.ulong)(optionValuesLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_optionValuesArray))
	optionValuesSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_optionValuesArray))[:optionValuesLen:optionValuesLen]
	for optionValuesi, optionValuesString := range optionValues {
		optionValuesSlice[optionValuesi] = (*C.gchar)(C.CString(optionValuesString))
		defer C.free(unsafe.Pointer(optionValuesSlice[optionValuesi]))
	}
	c_optionValues := &optionValuesSlice[0]

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_savev(c_pixbuf, c_filename, c_type_, c_optionKeys, c_optionValues, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_scale(src unsafe.Pointer, dest unsafe.Pointer, destX int, destY int, destWidth int, destHeight int, offsetX float64, offsetY float64, scaleX float64, scaleY float64, interpType int) {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_dest := (*C.GdkPixbuf)(unsafe.Pointer(dest))

	c_destX := (C.int)(destX)

	c_destY := (C.int)(destY)

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_offsetX := (C.double)(offsetX)

	c_offsetY := (C.double)(offsetY)

	c_scaleX := (C.double)(scaleX)

	c_scaleY := (C.double)(scaleY)

	c_interpType := (C.GdkInterpType)(interpType)

	C.gdk_pixbuf_scale(c_src, c_dest, c_destX, c_destY, c_destWidth, c_destHeight, c_offsetX, c_offsetY, c_scaleX, c_scaleY, c_interpType)
}

func Fn_gdk_pixbuf_scale_simple(src unsafe.Pointer, destWidth int, destHeight int, interpType int) unsafe.Pointer {
	c_src := (*C.GdkPixbuf)(unsafe.Pointer(src))

	c_destWidth := (C.int)(destWidth)

	c_destHeight := (C.int)(destHeight)

	c_interpType := (C.GdkInterpType)(interpType)

	ret := C.gdk_pixbuf_scale_simple(c_src, c_destWidth, c_destHeight, c_interpType)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_set_option(pixbuf unsafe.Pointer, key string, value string) bool {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_key := (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(c_key))

	c_value := (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(c_value))

	ret := C.gdk_pixbuf_set_option(c_pixbuf, c_key, c_value)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_unref(pixbuf unsafe.Pointer) {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gdk_pixbuf_unref(c_pixbuf)
}

func Fn_gdk_pixbuf_calculate_rowstride(colorspace int, hasAlpha bool, bitsPerSample int, width int, height int) int {
	c_colorspace := (C.GdkColorspace)(colorspace)

	c_hasAlpha := toCBool(hasAlpha)

	c_bitsPerSample := (C.int)(bitsPerSample)

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	ret := C.gdk_pixbuf_calculate_rowstride(c_colorspace, c_hasAlpha, c_bitsPerSample, c_width, c_height)

	return (int)(ret)
}

func Fn_gdk_pixbuf_get_file_info(filename string, width *int, height *int) unsafe.Pointer {
	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	ret := C.gdk_pixbuf_get_file_info(c_filename, c_width, c_height)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_get_file_info_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_get_file_info_finish(asyncResult unsafe.Pointer, width *int, height *int, error unsafe.Pointer) unsafe.Pointer {
	c_asyncResult := (*C.GAsyncResult)(unsafe.Pointer(asyncResult))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_get_file_info_finish(c_asyncResult, c_width, c_height, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_get_formats() unsafe.Pointer {
	ret := C.gdk_pixbuf_get_formats()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gdk_pixbuf_new_from_stream_async : parameter 'callback' is callback

// UNSUPPORTED : gdk_pixbuf_new_from_stream_at_scale_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_save_to_stream_finish(asyncResult unsafe.Pointer, error unsafe.Pointer) bool {
	c_asyncResult := (*C.GAsyncResult)(unsafe.Pointer(asyncResult))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_save_to_stream_finish(c_asyncResult, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_new_from_file(filename string, error unsafe.Pointer) unsafe.Pointer {
	c_filename := (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_file(c_filename, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_new_from_resource(resourcePath string, error unsafe.Pointer) unsafe.Pointer {
	c_resourcePath := (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(c_resourcePath))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_resource(c_resourcePath, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_new_from_stream(stream unsafe.Pointer, cancellable unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_stream := (*C.GInputStream)(unsafe.Pointer(stream))

	c_cancellable := (*C.GCancellable)(unsafe.Pointer(cancellable))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_stream(c_stream, c_cancellable, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_new_from_stream_finish(asyncResult unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	c_asyncResult := (*C.GAsyncResult)(unsafe.Pointer(asyncResult))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_animation_new_from_stream_finish(c_asyncResult, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_height(animation unsafe.Pointer) int {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_height(c_animation)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_get_iter(animation unsafe.Pointer, startTime unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	c_startTime := (*C.GTimeVal)(unsafe.Pointer(startTime))

	ret := C.gdk_pixbuf_animation_get_iter(c_animation, c_startTime)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_static_image(animation unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_static_image(c_animation)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_get_width(animation unsafe.Pointer) int {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_get_width(c_animation)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_is_static_image(animation unsafe.Pointer) bool {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_is_static_image(c_animation)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_ref(animation unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_animation_ref(c_animation)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_unref(animation unsafe.Pointer) {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	C.gdk_pixbuf_animation_unref(c_animation)
}

// UNSUPPORTED : gdk_pixbuf_animation_new_from_stream_async : parameter 'callback' is callback

func Fn_gdk_pixbuf_animation_iter_advance(iter unsafe.Pointer, currentTime unsafe.Pointer) bool {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	c_currentTime := (*C.GTimeVal)(unsafe.Pointer(currentTime))

	ret := C.gdk_pixbuf_animation_iter_advance(c_iter, c_currentTime)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_delay_time(iter unsafe.Pointer) int {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_get_delay_time(c_iter)

	return (int)(ret)
}

func Fn_gdk_pixbuf_animation_iter_get_pixbuf(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_get_pixbuf(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_animation_iter_on_currently_loading_frame(iter unsafe.Pointer) bool {
	c_iter := (*C.GdkPixbufAnimationIter)(unsafe.Pointer(iter))

	ret := C.gdk_pixbuf_animation_iter_on_currently_loading_frame(c_iter)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_new() unsafe.Pointer {
	ret := C.gdk_pixbuf_loader_new()

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_new_with_mime_type(mimeType string, error unsafe.Pointer) unsafe.Pointer {
	c_mimeType := (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(c_mimeType))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_new_with_mime_type(c_mimeType, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_new_with_type(imageType string, error unsafe.Pointer) unsafe.Pointer {
	c_imageType := (*C.char)(C.CString(imageType))
	defer C.free(unsafe.Pointer(c_imageType))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_new_with_type(c_imageType, cError)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_close(loader unsafe.Pointer, error unsafe.Pointer) bool {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_close(c_loader, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_get_animation(loader unsafe.Pointer) unsafe.Pointer {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	ret := C.gdk_pixbuf_loader_get_animation(c_loader)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_get_format(loader unsafe.Pointer) unsafe.Pointer {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	ret := C.gdk_pixbuf_loader_get_format(c_loader)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_get_pixbuf(loader unsafe.Pointer) unsafe.Pointer {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	ret := C.gdk_pixbuf_loader_get_pixbuf(c_loader)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_loader_set_size(loader unsafe.Pointer, width int, height int) {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	c_width := (C.int)(width)

	c_height := (C.int)(height)

	C.gdk_pixbuf_loader_set_size(c_loader, c_width, c_height)
}

func Fn_gdk_pixbuf_loader_write(loader unsafe.Pointer, buf []uint8, count uint64, error unsafe.Pointer) bool {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	c_buf := (*C.guchar)(unsafe.Pointer(&buf[0]))

	c_count := (C.gsize)(count)

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_write(c_loader, c_buf, c_count, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_loader_write_bytes(loader unsafe.Pointer, buffer unsafe.Pointer, error unsafe.Pointer) bool {
	c_loader := (*C.GdkPixbufLoader)(unsafe.Pointer(loader))

	c_buffer := (*C.GBytes)(unsafe.Pointer(buffer))

	cError := (**C.GError)(error)

	ret := C.gdk_pixbuf_loader_write_bytes(c_loader, c_buffer, cError)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_simple_anim_new(width int, height int, rate float32) unsafe.Pointer {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_rate := (C.gfloat)(rate)

	ret := C.gdk_pixbuf_simple_anim_new(c_width, c_height, c_rate)

	return unsafe.Pointer(ret)
}

func Fn_gdk_pixbuf_simple_anim_add_frame(animation unsafe.Pointer, pixbuf unsafe.Pointer) {
	c_animation := (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gdk_pixbuf_simple_anim_add_frame(c_animation, c_pixbuf)
}

func Fn_gdk_pixbuf_simple_anim_get_loop(animation unsafe.Pointer) bool {
	c_animation := (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation))

	ret := C.gdk_pixbuf_simple_anim_get_loop(c_animation)

	return toGoBool(ret)
}

func Fn_gdk_pixbuf_simple_anim_set_loop(animation unsafe.Pointer, loop bool) {
	c_animation := (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(animation))

	c_loop := toCBool(loop)

	C.gdk_pixbuf_simple_anim_set_loop(c_animation, c_loop)
}