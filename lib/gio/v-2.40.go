// Code generated - DO NOT EDIT.
// +build gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"fmt"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
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
/*

	void appinfomonitor_changedHandler(GObject *, gpointer);

	static gulong AppInfoMonitor_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(appinfomonitor_changedHandler), data);
	}

*/
/*

	gint application_handleLocalOptionsHandler(GObject *, GVariantDict *, gpointer);

	static gulong Application_signal_connect_handle_local_options(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "handle-local-options", G_CALLBACK(application_handleLocalOptionsHandler), data);
	}

*/
/*

	static void _g_notification_add_button_with_target(GNotification* notification, const gchar* label, const gchar* action, const gchar* target_format) {
		return g_notification_add_button_with_target(notification, label, action, target_format);
    }
*/
/*

	static void _g_notification_set_default_action_and_target(GNotification* notification, const gchar* action, const gchar* target_format) {
		return g_notification_set_default_action_and_target(notification, action, target_format);
    }
*/
/*

	static gboolean _g_output_stream_printf(GOutputStream* stream, gsize* bytes_written, GCancellable* cancellable, GError** error, const gchar* format) {
		return g_output_stream_printf(stream, bytes_written, cancellable, error, format);
    }
*/
import "C"

type SubprocessFlags C.GSubprocessFlags

const (
	SUBPROCESS_FLAGS_NONE           SubprocessFlags = 0
	SUBPROCESS_FLAGS_STDIN_PIPE     SubprocessFlags = 1
	SUBPROCESS_FLAGS_STDIN_INHERIT  SubprocessFlags = 2
	SUBPROCESS_FLAGS_STDOUT_PIPE    SubprocessFlags = 4
	SUBPROCESS_FLAGS_STDOUT_SILENCE SubprocessFlags = 8
	SUBPROCESS_FLAGS_STDERR_PIPE    SubprocessFlags = 16
	SUBPROCESS_FLAGS_STDERR_SILENCE SubprocessFlags = 32
	SUBPROCESS_FLAGS_STDERR_MERGE   SubprocessFlags = 64
	SUBPROCESS_FLAGS_INHERIT_FDS    SubprocessFlags = 128
)

// AppInfoMonitor is a wrapper around the C record GAppInfoMonitor.
type AppInfoMonitor struct {
	native *C.GAppInfoMonitor
}

func AppInfoMonitorNewFromC(u unsafe.Pointer) *AppInfoMonitor {
	c := (*C.GAppInfoMonitor)(u)
	if c == nil {
		return nil
	}

	g := &AppInfoMonitor{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *AppInfoMonitor) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *AppInfoMonitor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AppInfoMonitor with another AppInfoMonitor, and returns true if they represent the same GObject.
func (recv *AppInfoMonitor) Equals(other *AppInfoMonitor) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *AppInfoMonitor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to AppInfoMonitor.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppInfoMonitor.
func CastToAppInfoMonitor(object *gobject.Object) *AppInfoMonitor {
	return AppInfoMonitorNewFromC(object.ToC())
}

type signalAppInfoMonitorChangedDetail struct {
	callback  AppInfoMonitorSignalChangedCallback
	handlerID C.gulong
}

var signalAppInfoMonitorChangedId int
var signalAppInfoMonitorChangedMap = make(map[int]signalAppInfoMonitorChangedDetail)
var signalAppInfoMonitorChangedLock sync.RWMutex

// AppInfoMonitorSignalChangedCallback is a callback function for a 'changed' signal emitted from a AppInfoMonitor.
type AppInfoMonitorSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the AppInfoMonitor.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *AppInfoMonitor) ConnectChanged(callback AppInfoMonitorSignalChangedCallback) int {
	signalAppInfoMonitorChangedLock.Lock()
	defer signalAppInfoMonitorChangedLock.Unlock()

	signalAppInfoMonitorChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AppInfoMonitor_signal_connect_changed(instance, C.gpointer(uintptr(signalAppInfoMonitorChangedId)))

	detail := signalAppInfoMonitorChangedDetail{callback, handlerID}
	signalAppInfoMonitorChangedMap[signalAppInfoMonitorChangedId] = detail

	return signalAppInfoMonitorChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the AppInfoMonitor.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *AppInfoMonitor) DisconnectChanged(connectionID int) {
	signalAppInfoMonitorChangedLock.Lock()
	defer signalAppInfoMonitorChangedLock.Unlock()

	detail, exists := signalAppInfoMonitorChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAppInfoMonitorChangedMap, connectionID)
}

//export appinfomonitor_changedHandler
func appinfomonitor_changedHandler(_ *C.GObject, data C.gpointer) {
	signalAppInfoMonitorChangedLock.RLock()
	defer signalAppInfoMonitorChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalAppInfoMonitorChangedMap[index].callback
	callback()
}

// AppInfoMonitorGet is a wrapper around the C function g_app_info_monitor_get.
func AppInfoMonitorGet() *AppInfoMonitor {
	retC := C.g_app_info_monitor_get()
	retGo := AppInfoMonitorNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalApplicationHandleLocalOptionsDetail struct {
	callback  ApplicationSignalHandleLocalOptionsCallback
	handlerID C.gulong
}

var signalApplicationHandleLocalOptionsId int
var signalApplicationHandleLocalOptionsMap = make(map[int]signalApplicationHandleLocalOptionsDetail)
var signalApplicationHandleLocalOptionsLock sync.RWMutex

// ApplicationSignalHandleLocalOptionsCallback is a callback function for a 'handle-local-options' signal emitted from a Application.
type ApplicationSignalHandleLocalOptionsCallback func(options *glib.VariantDict) int32

/*
ConnectHandleLocalOptions connects the callback to the 'handle-local-options' signal for the Application.

The returned value represents the connection, and may be passed to DisconnectHandleLocalOptions to remove it.
*/
func (recv *Application) ConnectHandleLocalOptions(callback ApplicationSignalHandleLocalOptionsCallback) int {
	signalApplicationHandleLocalOptionsLock.Lock()
	defer signalApplicationHandleLocalOptionsLock.Unlock()

	signalApplicationHandleLocalOptionsId++
	instance := C.gpointer(recv.native)
	handlerID := C.Application_signal_connect_handle_local_options(instance, C.gpointer(uintptr(signalApplicationHandleLocalOptionsId)))

	detail := signalApplicationHandleLocalOptionsDetail{callback, handlerID}
	signalApplicationHandleLocalOptionsMap[signalApplicationHandleLocalOptionsId] = detail

	return signalApplicationHandleLocalOptionsId
}

/*
DisconnectHandleLocalOptions disconnects a callback from the 'handle-local-options' signal for the Application.

The connectionID should be a value returned from a call to ConnectHandleLocalOptions.
*/
func (recv *Application) DisconnectHandleLocalOptions(connectionID int) {
	signalApplicationHandleLocalOptionsLock.Lock()
	defer signalApplicationHandleLocalOptionsLock.Unlock()

	detail, exists := signalApplicationHandleLocalOptionsMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalApplicationHandleLocalOptionsMap, connectionID)
}

//export application_handleLocalOptionsHandler
func application_handleLocalOptionsHandler(_ *C.GObject, c_options *C.GVariantDict, data C.gpointer) C.gint {
	signalApplicationHandleLocalOptionsLock.RLock()
	defer signalApplicationHandleLocalOptionsLock.RUnlock()

	options := glib.VariantDictNewFromC(unsafe.Pointer(c_options))

	index := int(uintptr(data))
	callback := signalApplicationHandleLocalOptionsMap[index].callback
	retGo := callback(options)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported : g_application_add_main_option_entries : unsupported parameter entries :

// AddOptionGroup is a wrapper around the C function g_application_add_option_group.
func (recv *Application) AddOptionGroup(group *glib.OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_application_add_option_group((*C.GApplication)(recv.native), c_group)

	return
}

// SendNotification is a wrapper around the C function g_application_send_notification.
func (recv *Application) SendNotification(id string, notification *Notification) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_notification := (*C.GNotification)(C.NULL)
	if notification != nil {
		c_notification = (*C.GNotification)(notification.ToC())
	}

	C.g_application_send_notification((*C.GApplication)(recv.native), c_id, c_notification)

	return
}

// WithdrawNotification is a wrapper around the C function g_application_withdraw_notification.
func (recv *Application) WithdrawNotification(id string) {
	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	C.g_application_withdraw_notification((*C.GApplication)(recv.native), c_id)

	return
}

// GetOptionsDict is a wrapper around the C function g_application_command_line_get_options_dict.
func (recv *ApplicationCommandLine) GetOptionsDict() *glib.VariantDict {
	retC := C.g_application_command_line_get_options_dict((*C.GApplicationCommandLine)(recv.native))
	retGo := glib.VariantDictNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InetSocketAddressNewFromString is a wrapper around the C function g_inet_socket_address_new_from_string.
func InetSocketAddressNewFromString(address string, port uint32) *InetSocketAddress {
	c_address := C.CString(address)
	defer C.free(unsafe.Pointer(c_address))

	c_port := (C.guint)(port)

	retC := C.g_inet_socket_address_new_from_string(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Notification is a wrapper around the C record GNotification.
type Notification struct {
	native *C.GNotification
}

func NotificationNewFromC(u unsafe.Pointer) *Notification {
	c := (*C.GNotification)(u)
	if c == nil {
		return nil
	}

	g := &Notification{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Notification) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Notification) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Notification with another Notification, and returns true if they represent the same GObject.
func (recv *Notification) Equals(other *Notification) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Notification) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Notification.
// Exercise care, as this is a potentially dangerous function if the Object is not a Notification.
func CastToNotification(object *gobject.Object) *Notification {
	return NotificationNewFromC(object.ToC())
}

// NotificationNew is a wrapper around the C function g_notification_new.
func NotificationNew(title string) *Notification {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.g_notification_new(c_title)
	retGo := NotificationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddButton is a wrapper around the C function g_notification_add_button.
func (recv *Notification) AddButton(label string, detailedAction string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_add_button((*C.GNotification)(recv.native), c_label, c_detailed_action)

	return
}

// AddButtonWithTarget is a wrapper around the C function g_notification_add_button_with_target.
func (recv *Notification) AddButtonWithTarget(label string, action string, targetFormat string, args ...interface{}) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	goFormattedString := fmt.Sprintf(targetFormat, args...)
	c_target_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_target_format))

	C._g_notification_add_button_with_target((*C.GNotification)(recv.native), c_label, c_action, c_target_format)

	return
}

// AddButtonWithTargetValue is a wrapper around the C function g_notification_add_button_with_target_value.
func (recv *Notification) AddButtonWithTargetValue(label string, action string, target *glib.Variant) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	c_target := (*C.GVariant)(C.NULL)
	if target != nil {
		c_target = (*C.GVariant)(target.ToC())
	}

	C.g_notification_add_button_with_target_value((*C.GNotification)(recv.native), c_label, c_action, c_target)

	return
}

// SetBody is a wrapper around the C function g_notification_set_body.
func (recv *Notification) SetBody(body string) {
	c_body := C.CString(body)
	defer C.free(unsafe.Pointer(c_body))

	C.g_notification_set_body((*C.GNotification)(recv.native), c_body)

	return
}

// SetDefaultAction is a wrapper around the C function g_notification_set_default_action.
func (recv *Notification) SetDefaultAction(detailedAction string) {
	c_detailed_action := C.CString(detailedAction)
	defer C.free(unsafe.Pointer(c_detailed_action))

	C.g_notification_set_default_action((*C.GNotification)(recv.native), c_detailed_action)

	return
}

// SetDefaultActionAndTarget is a wrapper around the C function g_notification_set_default_action_and_target.
func (recv *Notification) SetDefaultActionAndTarget(action string, targetFormat string, args ...interface{}) {
	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	goFormattedString := fmt.Sprintf(targetFormat, args...)
	c_target_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_target_format))

	C._g_notification_set_default_action_and_target((*C.GNotification)(recv.native), c_action, c_target_format)

	return
}

// SetDefaultActionAndTargetValue is a wrapper around the C function g_notification_set_default_action_and_target_value.
func (recv *Notification) SetDefaultActionAndTargetValue(action string, target *glib.Variant) {
	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	c_target := (*C.GVariant)(C.NULL)
	if target != nil {
		c_target = (*C.GVariant)(target.ToC())
	}

	C.g_notification_set_default_action_and_target_value((*C.GNotification)(recv.native), c_action, c_target)

	return
}

// SetIcon is a wrapper around the C function g_notification_set_icon.
func (recv *Notification) SetIcon(icon *Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.g_notification_set_icon((*C.GNotification)(recv.native), c_icon)

	return
}

// SetPriority is a wrapper around the C function g_notification_set_priority.
func (recv *Notification) SetPriority(priority NotificationPriority) {
	c_priority := (C.GNotificationPriority)(priority)

	C.g_notification_set_priority((*C.GNotification)(recv.native), c_priority)

	return
}

// SetTitle is a wrapper around the C function g_notification_set_title.
func (recv *Notification) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.g_notification_set_title((*C.GNotification)(recv.native), c_title)

	return
}

// SetUrgent is a wrapper around the C function g_notification_set_urgent.
func (recv *Notification) SetUrgent(urgent bool) {
	c_urgent :=
		boolToGboolean(urgent)

	C.g_notification_set_urgent((*C.GNotification)(recv.native), c_urgent)

	return
}

// Printf is a wrapper around the C function g_output_stream_printf.
func (recv *OutputStream) Printf(cancellable *Cancellable, error *glib.Error, format string, args ...interface{}) (bool, uint64) {
	var c_bytes_written C.gsize

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	c_error := (**C.GError)(C.NULL)
	if error != nil {
		c_error = (**C.GError)(error.ToC())
	}

	goFormattedString := fmt.Sprintf(format, args...)
	c_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_format))

	retC := C._g_output_stream_printf((*C.GOutputStream)(recv.native), &c_bytes_written, c_cancellable, c_error, c_format)
	retGo := retC == C.TRUE

	bytesWritten := (uint64)(c_bytes_written)

	return retGo, bytesWritten
}

// Unsupported : g_output_stream_vprintf : unsupported parameter args : no type generator for va_list (va_list) for param args

// GetDefaultValue is a wrapper around the C function g_settings_get_default_value.
func (recv *Settings) GetDefaultValue(key string) *glib.Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_default_value((*C.GSettings)(recv.native), c_key)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetUserValue is a wrapper around the C function g_settings_get_user_value.
func (recv *Settings) GetUserValue(key string) *glib.Variant {
	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	retC := C.g_settings_get_user_value((*C.GSettings)(recv.native), c_key)
	var retGo (*glib.Variant)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.VariantNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Subprocess is a wrapper around the C record GSubprocess.
type Subprocess struct {
	native *C.GSubprocess
}

func SubprocessNewFromC(u unsafe.Pointer) *Subprocess {
	c := (*C.GSubprocess)(u)
	if c == nil {
		return nil
	}

	g := &Subprocess{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Subprocess) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Subprocess) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Subprocess with another Subprocess, and returns true if they represent the same GObject.
func (recv *Subprocess) Equals(other *Subprocess) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Subprocess) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Subprocess.
// Exercise care, as this is a potentially dangerous function if the Object is not a Subprocess.
func CastToSubprocess(object *gobject.Object) *Subprocess {
	return SubprocessNewFromC(object.ToC())
}

// Unsupported : g_subprocess_new : unsupported parameter ... : varargs

// SubprocessNewv is a wrapper around the C function g_subprocess_newv.
func SubprocessNewv(argv []string, flags SubprocessFlags) (*Subprocess, error) {
	c_argv_array := make([]*C.gchar, len(argv)+1, len(argv)+1)
	for i, item := range argv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_argv_array[i] = c
	}
	c_argv_array[len(argv)] = nil
	c_argv_arrayPtr := &c_argv_array[0]
	c_argv := (**C.gchar)(unsafe.Pointer(c_argv_arrayPtr))

	c_flags := (C.GSubprocessFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_subprocess_newv(c_argv, c_flags, &cThrowableError)
	retGo := SubprocessNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Communicate is a wrapper around the C function g_subprocess_communicate.
func (recv *Subprocess) Communicate(stdinBuf *glib.Bytes, cancellable *Cancellable) (bool, *glib.Bytes, *glib.Bytes, error) {
	c_stdin_buf := (*C.GBytes)(C.NULL)
	if stdinBuf != nil {
		c_stdin_buf = (*C.GBytes)(stdinBuf.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_stdout_buf *C.GBytes

	var c_stderr_buf *C.GBytes

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	stdoutBuf := glib.BytesNewFromC(unsafe.Pointer(c_stdout_buf))

	stderrBuf := glib.BytesNewFromC(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goError
}

// Unsupported : g_subprocess_communicate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CommunicateFinish is a wrapper around the C function g_subprocess_communicate_finish.
func (recv *Subprocess) CommunicateFinish(result *AsyncResult) (bool, *glib.Bytes, *glib.Bytes, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_stdout_buf *C.GBytes

	var c_stderr_buf *C.GBytes

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_finish((*C.GSubprocess)(recv.native), c_result, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	stdoutBuf := glib.BytesNewFromC(unsafe.Pointer(c_stdout_buf))

	stderrBuf := glib.BytesNewFromC(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goError
}

// CommunicateUtf8 is a wrapper around the C function g_subprocess_communicate_utf8.
func (recv *Subprocess) CommunicateUtf8(stdinBuf string, cancellable *Cancellable) (bool, string, string, error) {
	c_stdin_buf := C.CString(stdinBuf)
	defer C.free(unsafe.Pointer(c_stdin_buf))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var c_stdout_buf *C.char

	var c_stderr_buf *C.char

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_utf8((*C.GSubprocess)(recv.native), c_stdin_buf, c_cancellable, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	stdoutBuf := C.GoString(c_stdout_buf)
	defer C.free(unsafe.Pointer(c_stdout_buf))

	stderrBuf := C.GoString(c_stderr_buf)
	defer C.free(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goError
}

// Unsupported : g_subprocess_communicate_utf8_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CommunicateUtf8Finish is a wrapper around the C function g_subprocess_communicate_utf8_finish.
func (recv *Subprocess) CommunicateUtf8Finish(result *AsyncResult) (bool, string, string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_stdout_buf *C.char

	var c_stderr_buf *C.char

	var cThrowableError *C.GError

	retC := C.g_subprocess_communicate_utf8_finish((*C.GSubprocess)(recv.native), c_result, &c_stdout_buf, &c_stderr_buf, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	stdoutBuf := C.GoString(c_stdout_buf)
	defer C.free(unsafe.Pointer(c_stdout_buf))

	stderrBuf := C.GoString(c_stderr_buf)
	defer C.free(unsafe.Pointer(c_stderr_buf))

	return retGo, stdoutBuf, stderrBuf, goError
}

// ForceExit is a wrapper around the C function g_subprocess_force_exit.
func (recv *Subprocess) ForceExit() {
	C.g_subprocess_force_exit((*C.GSubprocess)(recv.native))

	return
}

// GetExitStatus is a wrapper around the C function g_subprocess_get_exit_status.
func (recv *Subprocess) GetExitStatus() int32 {
	retC := C.g_subprocess_get_exit_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIdentifier is a wrapper around the C function g_subprocess_get_identifier.
func (recv *Subprocess) GetIdentifier() string {
	retC := C.g_subprocess_get_identifier((*C.GSubprocess)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIfExited is a wrapper around the C function g_subprocess_get_if_exited.
func (recv *Subprocess) GetIfExited() bool {
	retC := C.g_subprocess_get_if_exited((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIfSignaled is a wrapper around the C function g_subprocess_get_if_signaled.
func (recv *Subprocess) GetIfSignaled() bool {
	retC := C.g_subprocess_get_if_signaled((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetStatus is a wrapper around the C function g_subprocess_get_status.
func (recv *Subprocess) GetStatus() int32 {
	retC := C.g_subprocess_get_status((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetStderrPipe is a wrapper around the C function g_subprocess_get_stderr_pipe.
func (recv *Subprocess) GetStderrPipe() *InputStream {
	retC := C.g_subprocess_get_stderr_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStdinPipe is a wrapper around the C function g_subprocess_get_stdin_pipe.
func (recv *Subprocess) GetStdinPipe() *OutputStream {
	retC := C.g_subprocess_get_stdin_pipe((*C.GSubprocess)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStdoutPipe is a wrapper around the C function g_subprocess_get_stdout_pipe.
func (recv *Subprocess) GetStdoutPipe() *InputStream {
	retC := C.g_subprocess_get_stdout_pipe((*C.GSubprocess)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSuccessful is a wrapper around the C function g_subprocess_get_successful.
func (recv *Subprocess) GetSuccessful() bool {
	retC := C.g_subprocess_get_successful((*C.GSubprocess)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTermSig is a wrapper around the C function g_subprocess_get_term_sig.
func (recv *Subprocess) GetTermSig() int32 {
	retC := C.g_subprocess_get_term_sig((*C.GSubprocess)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SendSignal is a wrapper around the C function g_subprocess_send_signal.
func (recv *Subprocess) SendSignal(signalNum int32) {
	c_signal_num := (C.gint)(signalNum)

	C.g_subprocess_send_signal((*C.GSubprocess)(recv.native), c_signal_num)

	return
}

// Wait is a wrapper around the C function g_subprocess_wait.
func (recv *Subprocess) Wait(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_subprocess_wait_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WaitCheck is a wrapper around the C function g_subprocess_wait_check.
func (recv *Subprocess) WaitCheck(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_check((*C.GSubprocess)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_subprocess_wait_check_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// WaitCheckFinish is a wrapper around the C function g_subprocess_wait_check_finish.
func (recv *Subprocess) WaitCheckFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_check_finish((*C.GSubprocess)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// WaitFinish is a wrapper around the C function g_subprocess_wait_finish.
func (recv *Subprocess) WaitFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_subprocess_wait_finish((*C.GSubprocess)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SubprocessLauncher is a wrapper around the C record GSubprocessLauncher.
type SubprocessLauncher struct {
	native *C.GSubprocessLauncher
}

func SubprocessLauncherNewFromC(u unsafe.Pointer) *SubprocessLauncher {
	c := (*C.GSubprocessLauncher)(u)
	if c == nil {
		return nil
	}

	g := &SubprocessLauncher{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SubprocessLauncher) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SubprocessLauncher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SubprocessLauncher with another SubprocessLauncher, and returns true if they represent the same GObject.
func (recv *SubprocessLauncher) Equals(other *SubprocessLauncher) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SubprocessLauncher) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SubprocessLauncher.
// Exercise care, as this is a potentially dangerous function if the Object is not a SubprocessLauncher.
func CastToSubprocessLauncher(object *gobject.Object) *SubprocessLauncher {
	return SubprocessLauncherNewFromC(object.ToC())
}

// SubprocessLauncherNew is a wrapper around the C function g_subprocess_launcher_new.
func SubprocessLauncherNew(flags SubprocessFlags) *SubprocessLauncher {
	c_flags := (C.GSubprocessFlags)(flags)

	retC := C.g_subprocess_launcher_new(c_flags)
	retGo := SubprocessLauncherNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Getenv is a wrapper around the C function g_subprocess_launcher_getenv.
func (recv *SubprocessLauncher) Getenv(variable string) string {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	retC := C.g_subprocess_launcher_getenv((*C.GSubprocessLauncher)(recv.native), c_variable)
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_subprocess_launcher_set_child_setup : unsupported parameter child_setup : no type generator for GLib.SpawnChildSetupFunc (GSpawnChildSetupFunc) for param child_setup

// SetCwd is a wrapper around the C function g_subprocess_launcher_set_cwd.
func (recv *SubprocessLauncher) SetCwd(cwd string) {
	c_cwd := C.CString(cwd)
	defer C.free(unsafe.Pointer(c_cwd))

	C.g_subprocess_launcher_set_cwd((*C.GSubprocessLauncher)(recv.native), c_cwd)

	return
}

// SetEnviron is a wrapper around the C function g_subprocess_launcher_set_environ.
func (recv *SubprocessLauncher) SetEnviron(env []string) {
	c_env_array := make([]*C.gchar, len(env)+1, len(env)+1)
	for i, item := range env {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_env_array[i] = c
	}
	c_env_array[len(env)] = nil
	c_env_arrayPtr := &c_env_array[0]
	c_env := (**C.gchar)(unsafe.Pointer(c_env_arrayPtr))

	C.g_subprocess_launcher_set_environ((*C.GSubprocessLauncher)(recv.native), c_env)

	return
}

// SetFlags is a wrapper around the C function g_subprocess_launcher_set_flags.
func (recv *SubprocessLauncher) SetFlags(flags SubprocessFlags) {
	c_flags := (C.GSubprocessFlags)(flags)

	C.g_subprocess_launcher_set_flags((*C.GSubprocessLauncher)(recv.native), c_flags)

	return
}

// SetStderrFilePath is a wrapper around the C function g_subprocess_launcher_set_stderr_file_path.
func (recv *SubprocessLauncher) SetStderrFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stderr_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// SetStdinFilePath is a wrapper around the C function g_subprocess_launcher_set_stdin_file_path.
func (recv *SubprocessLauncher) SetStdinFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdin_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// SetStdoutFilePath is a wrapper around the C function g_subprocess_launcher_set_stdout_file_path.
func (recv *SubprocessLauncher) SetStdoutFilePath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.g_subprocess_launcher_set_stdout_file_path((*C.GSubprocessLauncher)(recv.native), c_path)

	return
}

// Setenv is a wrapper around the C function g_subprocess_launcher_setenv.
func (recv *SubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_overwrite :=
		boolToGboolean(overwrite)

	C.g_subprocess_launcher_setenv((*C.GSubprocessLauncher)(recv.native), c_variable, c_value, c_overwrite)

	return
}

// Unsupported : g_subprocess_launcher_spawn : unsupported parameter ... : varargs

// Spawnv is a wrapper around the C function g_subprocess_launcher_spawnv.
func (recv *SubprocessLauncher) Spawnv(argv []string) (*Subprocess, error) {
	c_argv_array := make([]*C.gchar, len(argv)+1, len(argv)+1)
	for i, item := range argv {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_argv_array[i] = c
	}
	c_argv_array[len(argv)] = nil
	c_argv_arrayPtr := &c_argv_array[0]
	c_argv := (**C.gchar)(unsafe.Pointer(c_argv_arrayPtr))

	var cThrowableError *C.GError

	retC := C.g_subprocess_launcher_spawnv((*C.GSubprocessLauncher)(recv.native), c_argv, &cThrowableError)
	retGo := SubprocessNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TakeFd is a wrapper around the C function g_subprocess_launcher_take_fd.
func (recv *SubprocessLauncher) TakeFd(sourceFd int32, targetFd int32) {
	c_source_fd := (C.gint)(sourceFd)

	c_target_fd := (C.gint)(targetFd)

	C.g_subprocess_launcher_take_fd((*C.GSubprocessLauncher)(recv.native), c_source_fd, c_target_fd)

	return
}

// TakeStderrFd is a wrapper around the C function g_subprocess_launcher_take_stderr_fd.
func (recv *SubprocessLauncher) TakeStderrFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stderr_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// TakeStdinFd is a wrapper around the C function g_subprocess_launcher_take_stdin_fd.
func (recv *SubprocessLauncher) TakeStdinFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdin_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// TakeStdoutFd is a wrapper around the C function g_subprocess_launcher_take_stdout_fd.
func (recv *SubprocessLauncher) TakeStdoutFd(fd int32) {
	c_fd := (C.gint)(fd)

	C.g_subprocess_launcher_take_stdout_fd((*C.GSubprocessLauncher)(recv.native), c_fd)

	return
}

// Unsetenv is a wrapper around the C function g_subprocess_launcher_unsetenv.
func (recv *SubprocessLauncher) Unsetenv(variable string) {
	c_variable := C.CString(variable)
	defer C.free(unsafe.Pointer(c_variable))

	C.g_subprocess_launcher_unsetenv((*C.GSubprocessLauncher)(recv.native), c_variable)

	return
}

// InvokeRequestCertificate is a wrapper around the C function g_tls_interaction_invoke_request_certificate.
func (recv *TlsInteraction) InvokeRequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GTlsConnection)(connection.ToC())
	}

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_invoke_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RequestCertificate is a wrapper around the C function g_tls_interaction_request_certificate.
func (recv *TlsInteraction) RequestCertificate(connection *TlsConnection, flags TlsCertificateRequestFlags, cancellable *Cancellable) (TlsInteractionResult, error) {
	c_connection := (*C.GTlsConnection)(C.NULL)
	if connection != nil {
		c_connection = (*C.GTlsConnection)(connection.ToC())
	}

	c_flags := (C.GTlsCertificateRequestFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_request_certificate((*C.GTlsInteraction)(recv.native), c_connection, c_flags, c_cancellable, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_tls_interaction_request_certificate_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// RequestCertificateFinish is a wrapper around the C function g_tls_interaction_request_certificate_finish.
func (recv *TlsInteraction) RequestCertificateFinish(result *AsyncResult) (TlsInteractionResult, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_tls_interaction_request_certificate_finish((*C.GTlsInteraction)(recv.native), c_result, &cThrowableError)
	retGo := (TlsInteractionResult)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

const FILE_ATTRIBUTE_THUMBNAIL_IS_VALID string = C.G_FILE_ATTRIBUTE_THUMBNAIL_IS_VALID

type TlsCertificateRequestFlags C.GTlsCertificateRequestFlags

const (
	TLS_CERTIFICATE_REQUEST_NONE TlsCertificateRequestFlags = 0
)

// Unsupported : g_file_replace_contents_bytes_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// GetKey is a wrapper around the C function g_settings_schema_get_key.
func (recv *SettingsSchema) GetKey(name string) *SettingsSchemaKey {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_get_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasKey is a wrapper around the C function g_settings_schema_has_key.
func (recv *SettingsSchema) HasKey(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.g_settings_schema_has_key((*C.GSettingsSchema)(recv.native), c_name)
	retGo := retC == C.TRUE

	return retGo
}

// GetDefaultValue is a wrapper around the C function g_settings_schema_key_get_default_value.
func (recv *SettingsSchemaKey) GetDefaultValue() *glib.Variant {
	retC := C.g_settings_schema_key_get_default_value((*C.GSettingsSchemaKey)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRange is a wrapper around the C function g_settings_schema_key_get_range.
func (recv *SettingsSchemaKey) GetRange() *glib.Variant {
	retC := C.g_settings_schema_key_get_range((*C.GSettingsSchemaKey)(recv.native))
	retGo := glib.VariantNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValueType is a wrapper around the C function g_settings_schema_key_get_value_type.
func (recv *SettingsSchemaKey) GetValueType() *glib.VariantType {
	retC := C.g_settings_schema_key_get_value_type((*C.GSettingsSchemaKey)(recv.native))
	retGo := glib.VariantTypeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RangeCheck is a wrapper around the C function g_settings_schema_key_range_check.
func (recv *SettingsSchemaKey) RangeCheck(value *glib.Variant) bool {
	c_value := (*C.GVariant)(C.NULL)
	if value != nil {
		c_value = (*C.GVariant)(value.ToC())
	}

	retC := C.g_settings_schema_key_range_check((*C.GSettingsSchemaKey)(recv.native), c_value)
	retGo := retC == C.TRUE

	return retGo
}

// Ref is a wrapper around the C function g_settings_schema_key_ref.
func (recv *SettingsSchemaKey) Ref() *SettingsSchemaKey {
	retC := C.g_settings_schema_key_ref((*C.GSettingsSchemaKey)(recv.native))
	retGo := SettingsSchemaKeyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unref is a wrapper around the C function g_settings_schema_key_unref.
func (recv *SettingsSchemaKey) Unref() {
	C.g_settings_schema_key_unref((*C.GSettingsSchemaKey)(recv.native))

	return
}

// Unsupported : g_settings_schema_source_list_schemas : unsupported parameter non_relocatable : output array param non_relocatable