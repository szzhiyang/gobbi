// Code generated - DO NOT EDIT.
// +build gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// SetOptionContextDescription is a wrapper around the C function g_application_set_option_context_description.
func (recv *Application) SetOptionContextDescription(description string) {
	c_description := C.CString(description)
	defer C.free(unsafe.Pointer(c_description))

	C.g_application_set_option_context_description((*C.GApplication)(recv.native), c_description)

	return
}

// SetOptionContextParameterString is a wrapper around the C function g_application_set_option_context_parameter_string.
func (recv *Application) SetOptionContextParameterString(parameterString string) {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	C.g_application_set_option_context_parameter_string((*C.GApplication)(recv.native), c_parameter_string)

	return
}

// SetOptionContextSummary is a wrapper around the C function g_application_set_option_context_summary.
func (recv *Application) SetOptionContextSummary(summary string) {
	c_summary := C.CString(summary)
	defer C.free(unsafe.Pointer(c_summary))

	C.g_application_set_option_context_summary((*C.GApplication)(recv.native), c_summary)

	return
}

// GetLocaleString is a wrapper around the C function g_desktop_app_info_get_locale_string.
func (recv *DesktopAppInfo) GetLocaleString(key string) string {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_desktop_app_info_get_locale_string((*C.GDesktopAppInfo)(recv.native), c_key)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// JoinMulticastGroupSsm is a wrapper around the C function g_socket_join_multicast_group_ssm.
func (recv *Socket) JoinMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific := (*C.GInetAddress)(C.NULL)
	if sourceSpecific != nil {
		c_source_specific = (*C.GInetAddress)(sourceSpecific.ToC())
	}

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_join_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LeaveMulticastGroupSsm is a wrapper around the C function g_socket_leave_multicast_group_ssm.
func (recv *Socket) LeaveMulticastGroupSsm(group *InetAddress, sourceSpecific *InetAddress, iface string) (bool, error) {
	c_group := (*C.GInetAddress)(C.NULL)
	if group != nil {
		c_group = (*C.GInetAddress)(group.ToC())
	}

	c_source_specific := (*C.GInetAddress)(C.NULL)
	if sourceSpecific != nil {
		c_source_specific = (*C.GInetAddress)(sourceSpecific.ToC())
	}

	c_iface := C.CString(iface)
	defer C.free(unsafe.Pointer(c_iface))

	var cThrowableError *C.GError

	retC := C.g_socket_leave_multicast_group_ssm((*C.GSocket)(recv.native), c_group, c_source_specific, c_iface, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixIsSystemDevicePath is a wrapper around the C function g_unix_is_system_device_path.
func UnixIsSystemDevicePath(devicePath string) bool {
	c_device_path := C.CString(devicePath)
	defer C.free(unsafe.Pointer(c_device_path))

	retC := C.g_unix_is_system_device_path(c_device_path)
	retGo := retC == C.TRUE

	return retGo
}

// UnixIsSystemFsType is a wrapper around the C function g_unix_is_system_fs_type.
func UnixIsSystemFsType(fsType string) bool {
	c_fs_type := C.CString(fsType)
	defer C.free(unsafe.Pointer(c_fs_type))

	retC := C.g_unix_is_system_fs_type(c_fs_type)
	retGo := retC == C.TRUE

	return retGo
}

// g_file_new_build_filename : unsupported parameter ... : varargs
// LoadBytes is a wrapper around the C function g_file_load_bytes.
func (recv *File) LoadBytes(cancellable *Cancellable) (*glib.Bytes, string, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes((*C.GFile)(recv.native), c_cancellable, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goError
}

// Unsupported : g_file_load_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LoadBytesFinish is a wrapper around the C function g_file_load_bytes_finish.
func (recv *File) LoadBytesFinish(result *AsyncResult) (*glib.Bytes, string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_etag_out *C.gchar

	var cThrowableError *C.GError

	retC := C.g_file_load_bytes_finish((*C.GFile)(recv.native), c_result, &c_etag_out, &cThrowableError)
	retGo := glib.BytesNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	etagOut := C.GoString(c_etag_out)
	defer C.free(unsafe.Pointer(c_etag_out))

	return retGo, etagOut, goError
}

// PeekPath is a wrapper around the C function g_file_peek_path.
func (recv *File) PeekPath() string {
	retC := C.g_file_peek_path((*C.GFile)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}