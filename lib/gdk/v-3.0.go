// Code generated - DO NOT EDIT.
// +build gdk_3.0 gdk_3.4 gdk_3.8 gdk_3.10 gdk_3.12 gdk_3.14 gdk_3.16 gdk_3.18 gdk_3.20 gdk_3.22

package gdk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gdk/gdk.h>
// #include <stdlib.h>
/*

	cairo_surface_t * window_createSurfaceHandler(GObject *, gint, gint, gpointer);

	static gulong Window_signal_connect_create_surface(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-surface", G_CALLBACK(window_createSurfaceHandler), data);
	}

*/
import "C"

// GetAssociatedDevice is a wrapper around the C function gdk_device_get_associated_device.
func (recv *Device) GetAssociatedDevice() *Device {
	retC := C.gdk_device_get_associated_device((*C.GdkDevice)(recv.native))
	var retGo (*Device)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DeviceNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetAxisValue is a wrapper around the C function gdk_device_get_axis_value.
func (recv *Device) GetAxisValue(axes []float64, axisLabel *Atom) (bool, float64) {
	c_axes_array := make([]C.gdouble, len(axes)+1, len(axes)+1)
	for i, item := range axes {
		c := (C.gdouble)(item)
		c_axes_array[i] = c
	}
	c_axes_array[len(axes)] = 0
	c_axes_arrayPtr := &c_axes_array[0]
	c_axes := (*C.gdouble)(unsafe.Pointer(c_axes_arrayPtr))

	c_axis_label := (C.GdkAtom)(C.NULL)
	if axisLabel != nil {
		c_axis_label = (C.GdkAtom)(axisLabel.ToC())
	}

	var c_value C.gdouble

	retC := C.gdk_device_get_axis_value((*C.GdkDevice)(recv.native), c_axes, c_axis_label, &c_value)
	retGo := retC == C.TRUE

	value := (float64)(c_value)

	return retGo, value
}

// GetDeviceType is a wrapper around the C function gdk_device_get_device_type.
func (recv *Device) GetDeviceType() DeviceType {
	retC := C.gdk_device_get_device_type((*C.GdkDevice)(recv.native))
	retGo := (DeviceType)(retC)

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_device_get_display.
func (recv *Device) GetDisplay() *Display {
	retC := C.gdk_device_get_display((*C.GdkDevice)(recv.native))
	retGo := DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetNAxes is a wrapper around the C function gdk_device_get_n_axes.
func (recv *Device) GetNAxes() int32 {
	retC := C.gdk_device_get_n_axes((*C.GdkDevice)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPosition is a wrapper around the C function gdk_device_get_position.
func (recv *Device) GetPosition() (*Screen, int32, int32) {
	var c_screen *C.GdkScreen

	var c_x C.gint

	var c_y C.gint

	C.gdk_device_get_position((*C.GdkDevice)(recv.native), &c_screen, &c_x, &c_y)

	screen := ScreenNewFromC(unsafe.Pointer(c_screen))

	x := (int32)(c_x)

	y := (int32)(c_y)

	return screen, x, y
}

// GetWindowAtPosition is a wrapper around the C function gdk_device_get_window_at_position.
func (recv *Device) GetWindowAtPosition() (*Window, int32, int32) {
	var c_win_x C.gint

	var c_win_y C.gint

	retC := C.gdk_device_get_window_at_position((*C.GdkDevice)(recv.native), &c_win_x, &c_win_y)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	winX := (int32)(c_win_x)

	winY := (int32)(c_win_y)

	return retGo, winX, winY
}

// GetWindowAtPositionDouble is a wrapper around the C function gdk_device_get_window_at_position_double.
func (recv *Device) GetWindowAtPositionDouble() (*Window, float64, float64) {
	var c_win_x C.gdouble

	var c_win_y C.gdouble

	retC := C.gdk_device_get_window_at_position_double((*C.GdkDevice)(recv.native), &c_win_x, &c_win_y)
	var retGo (*Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WindowNewFromC(unsafe.Pointer(retC))
	}

	winX := (float64)(c_win_x)

	winY := (float64)(c_win_y)

	return retGo, winX, winY
}

// Grab is a wrapper around the C function gdk_device_grab.
func (recv *Device) Grab(window *Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor *Cursor, time uint32) GrabStatus {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_grab_ownership := (C.GdkGrabOwnership)(grabOwnership)

	c_owner_events :=
		boolToGboolean(ownerEvents)

	c_event_mask := (C.GdkEventMask)(eventMask)

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	c_time_ := (C.guint32)(time)

	retC := C.gdk_device_grab((*C.GdkDevice)(recv.native), c_window, c_grab_ownership, c_owner_events, c_event_mask, c_cursor, c_time_)
	retGo := (GrabStatus)(retC)

	return retGo
}

// ListAxes is a wrapper around the C function gdk_device_list_axes.
func (recv *Device) ListAxes() *glib.List {
	retC := C.gdk_device_list_axes((*C.GdkDevice)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ungrab is a wrapper around the C function gdk_device_ungrab.
func (recv *Device) Ungrab(time uint32) {
	c_time_ := (C.guint32)(time)

	C.gdk_device_ungrab((*C.GdkDevice)(recv.native), c_time_)

	return
}

// Warp is a wrapper around the C function gdk_device_warp.
func (recv *Device) Warp(screen *Screen, x int32, y int32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gdk_device_warp((*C.GdkDevice)(recv.native), c_screen, c_x, c_y)

	return
}

// GetClientPointer is a wrapper around the C function gdk_device_manager_get_client_pointer.
func (recv *DeviceManager) GetClientPointer() *Device {
	retC := C.gdk_device_manager_get_client_pointer((*C.GdkDeviceManager)(recv.native))
	retGo := DeviceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gdk_device_manager_get_display.
func (recv *DeviceManager) GetDisplay() *Display {
	retC := C.gdk_device_manager_get_display((*C.GdkDeviceManager)(recv.native))
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListDevices is a wrapper around the C function gdk_device_manager_list_devices.
func (recv *DeviceManager) ListDevices(type_ DeviceType) *glib.List {
	c_type := (C.GdkDeviceType)(type_)

	retC := C.gdk_device_manager_list_devices((*C.GdkDeviceManager)(recv.native), c_type)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAppLaunchContext is a wrapper around the C function gdk_display_get_app_launch_context.
func (recv *Display) GetAppLaunchContext() *AppLaunchContext {
	retC := C.gdk_display_get_app_launch_context((*C.GdkDisplay)(recv.native))
	retGo := AppLaunchContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeviceManager is a wrapper around the C function gdk_display_get_device_manager.
func (recv *Display) GetDeviceManager() *DeviceManager {
	retC := C.gdk_display_get_device_manager((*C.GdkDisplay)(recv.native))
	var retGo (*DeviceManager)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DeviceManagerNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// HasPending is a wrapper around the C function gdk_display_has_pending.
func (recv *Display) HasPending() bool {
	retC := C.gdk_display_has_pending((*C.GdkDisplay)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// NotifyStartupComplete is a wrapper around the C function gdk_display_notify_startup_complete.
func (recv *Display) NotifyStartupComplete(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gdk_display_notify_startup_complete((*C.GdkDisplay)(recv.native), c_startup_id)

	return
}

// OpenDisplay is a wrapper around the C function gdk_display_manager_open_display.
func (recv *DisplayManager) OpenDisplay(name string) *Display {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gdk_display_manager_open_display((*C.GdkDisplayManager)(recv.native), c_name)
	var retGo (*Display)
	if retC == nil {
		retGo = nil
	} else {
		retGo = DisplayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetDestWindow is a wrapper around the C function gdk_drag_context_get_dest_window.
func (recv *DragContext) GetDestWindow() *Window {
	retC := C.gdk_drag_context_get_dest_window((*C.GdkDragContext)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProtocol is a wrapper around the C function gdk_drag_context_get_protocol.
func (recv *DragContext) GetProtocol() DragProtocol {
	retC := C.gdk_drag_context_get_protocol((*C.GdkDragContext)(recv.native))
	retGo := (DragProtocol)(retC)

	return retGo
}

// GetNumLockState is a wrapper around the C function gdk_keymap_get_num_lock_state.
func (recv *Keymap) GetNumLockState() bool {
	retC := C.gdk_keymap_get_num_lock_state((*C.GdkKeymap)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type signalWindowCreateSurfaceDetail struct {
	callback  WindowSignalCreateSurfaceCallback
	handlerID C.gulong
}

var signalWindowCreateSurfaceId int
var signalWindowCreateSurfaceMap = make(map[int]signalWindowCreateSurfaceDetail)
var signalWindowCreateSurfaceLock sync.RWMutex

// WindowSignalCreateSurfaceCallback is a callback function for a 'create-surface' signal emitted from a Window.
type WindowSignalCreateSurfaceCallback func(width int32, height int32) cairo.Surface

/*
ConnectCreateSurface connects the callback to the 'create-surface' signal for the Window.

The returned value represents the connection, and may be passed to DisconnectCreateSurface to remove it.
*/
func (recv *Window) ConnectCreateSurface(callback WindowSignalCreateSurfaceCallback) int {
	signalWindowCreateSurfaceLock.Lock()
	defer signalWindowCreateSurfaceLock.Unlock()

	signalWindowCreateSurfaceId++
	instance := C.gpointer(recv.native)
	handlerID := C.Window_signal_connect_create_surface(instance, C.gpointer(uintptr(signalWindowCreateSurfaceId)))

	detail := signalWindowCreateSurfaceDetail{callback, handlerID}
	signalWindowCreateSurfaceMap[signalWindowCreateSurfaceId] = detail

	return signalWindowCreateSurfaceId
}

/*
DisconnectCreateSurface disconnects a callback from the 'create-surface' signal for the Window.

The connectionID should be a value returned from a call to ConnectCreateSurface.
*/
func (recv *Window) DisconnectCreateSurface(connectionID int) {
	signalWindowCreateSurfaceLock.Lock()
	defer signalWindowCreateSurfaceLock.Unlock()

	detail, exists := signalWindowCreateSurfaceMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWindowCreateSurfaceMap, connectionID)
}

//export window_createSurfaceHandler
func window_createSurfaceHandler(_ *C.GObject, c_width C.gint, c_height C.gint, data C.gpointer) *C.cairo_surface_t {
	signalWindowCreateSurfaceLock.RLock()
	defer signalWindowCreateSurfaceLock.RUnlock()

	width := int32(c_width)

	height := int32(c_height)

	index := int(uintptr(data))
	callback := signalWindowCreateSurfaceMap[index].callback
	retGo := callback(width, height)
	retC :=
		(*C.cairo_surface_t)(retGo.ToC())
	return retC
}

// GetDeviceCursor is a wrapper around the C function gdk_window_get_device_cursor.
func (recv *Window) GetDeviceCursor(device *Device) *Cursor {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_window_get_device_cursor((*C.GdkWindow)(recv.native), c_device)
	var retGo (*Cursor)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CursorNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetDeviceEvents is a wrapper around the C function gdk_window_get_device_events.
func (recv *Window) GetDeviceEvents(device *Device) EventMask {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gdk_window_get_device_events((*C.GdkWindow)(recv.native), c_device)
	retGo := (EventMask)(retC)

	return retGo
}

// Unsupported : gdk_window_get_device_position : unsupported parameter mask : GdkModifierType* with indirection level of 1

// GetDragProtocol is a wrapper around the C function gdk_window_get_drag_protocol.
func (recv *Window) GetDragProtocol() (DragProtocol, *Window) {
	var c_target *C.GdkWindow

	retC := C.gdk_window_get_drag_protocol((*C.GdkWindow)(recv.native), &c_target)
	retGo := (DragProtocol)(retC)

	target := WindowNewFromC(unsafe.Pointer(c_target))

	return retGo, target
}

// GetSupportMultidevice is a wrapper around the C function gdk_window_get_support_multidevice.
func (recv *Window) GetSupportMultidevice() bool {
	retC := C.gdk_window_get_support_multidevice((*C.GdkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDeviceCursor is a wrapper around the C function gdk_window_set_device_cursor.
func (recv *Window) SetDeviceCursor(device *Device, cursor *Cursor) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_cursor := (*C.GdkCursor)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkCursor)(cursor.ToC())
	}

	C.gdk_window_set_device_cursor((*C.GdkWindow)(recv.native), c_device, c_cursor)

	return
}

// SetDeviceEvents is a wrapper around the C function gdk_window_set_device_events.
func (recv *Window) SetDeviceEvents(device *Device, eventMask EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_device_events((*C.GdkWindow)(recv.native), c_device, c_event_mask)

	return
}

// SetSourceEvents is a wrapper around the C function gdk_window_set_source_events.
func (recv *Window) SetSourceEvents(source InputSource, eventMask EventMask) {
	c_source := (C.GdkInputSource)(source)

	c_event_mask := (C.GdkEventMask)(eventMask)

	C.gdk_window_set_source_events((*C.GdkWindow)(recv.native), c_source, c_event_mask)

	return
}

// SetSupportMultidevice is a wrapper around the C function gdk_window_set_support_multidevice.
func (recv *Window) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gdk_window_set_support_multidevice((*C.GdkWindow)(recv.native), c_support_multidevice)

	return
}

// CairoSetSourceRgba is a wrapper around the C function gdk_cairo_set_source_rgba.
func CairoSetSourceRgba(cr *cairo.Context, rgba *RGBA) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gdk_cairo_set_source_rgba(c_cr, c_rgba)

	return
}

// DisableMultidevice is a wrapper around the C function gdk_disable_multidevice.
func DisableMultidevice() {
	C.gdk_disable_multidevice()

	return
}

// ErrorTrapPopIgnored is a wrapper around the C function gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()

	return
}

// Unsupported : gdk_events_get_angle : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_center : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Unsupported : gdk_events_get_distance : unsupported parameter event1 : no type generator for Event (GdkEvent*) for param event1

// Copy is a wrapper around the C function gdk_rgba_copy.
func (recv *RGBA) Copy() *RGBA {
	retC := C.gdk_rgba_copy((*C.GdkRGBA)(recv.native))
	retGo := RGBANewFromC(unsafe.Pointer(retC))

	return retGo
}

// Equal is a wrapper around the C function gdk_rgba_equal.
func (recv *RGBA) Equal(p2 *RGBA) bool {
	c_p2 := (C.gconstpointer)(C.NULL)
	if p2 != nil {
		c_p2 = (C.gconstpointer)(p2.ToC())
	}

	retC := C.gdk_rgba_equal((C.gconstpointer)(recv.native), c_p2)
	retGo := retC == C.TRUE

	return retGo
}

// Free is a wrapper around the C function gdk_rgba_free.
func (recv *RGBA) Free() {
	C.gdk_rgba_free((*C.GdkRGBA)(recv.native))

	return
}

// Hash is a wrapper around the C function gdk_rgba_hash.
func (recv *RGBA) Hash() uint32 {
	retC := C.gdk_rgba_hash((C.gconstpointer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Parse is a wrapper around the C function gdk_rgba_parse.
func (recv *RGBA) Parse(spec string) bool {
	c_spec := C.CString(spec)
	defer C.free(unsafe.Pointer(c_spec))

	retC := C.gdk_rgba_parse((*C.GdkRGBA)(recv.native), c_spec)
	retGo := retC == C.TRUE

	return retGo
}

// ToString is a wrapper around the C function gdk_rgba_to_string.
func (recv *RGBA) ToString() string {
	retC := C.gdk_rgba_to_string((*C.GdkRGBA)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}