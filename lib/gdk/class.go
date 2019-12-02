// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	gio "github.com/pekim/gobbi/lib/gio"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
	"sync"
)

var appLaunchContextStruct *gi.Struct
var appLaunchContextStruct_Once sync.Once

func appLaunchContextStruct_Set() error {
	var err error
	appLaunchContextStruct_Once.Do(func() {
		appLaunchContextStruct, err = gi.StructNew("Gdk", "AppLaunchContext")
	})
	return err
}

type AppLaunchContext struct {
	gio.AppLaunchContext
}

var appLaunchContextNewFunction *gi.Function
var appLaunchContextNewFunction_Once sync.Once

func appLaunchContextNewFunction_Set() error {
	var err error
	appLaunchContextNewFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextNewFunction, err = appLaunchContextStruct.InvokerNew("new")
	})
	return err
}

// AppLaunchContextNew is a representation of the C type gdk_app_launch_context_new.
func AppLaunchContextNew() *AppLaunchContext {

	var ret gi.Argument

	err := appLaunchContextNewFunction_Set()
	if err == nil {
		ret = appLaunchContextNewFunction.Invoke(nil, nil)
	}

	retGo := &AppLaunchContext{}
	retGo.Native = ret.Pointer()

	return retGo
}

var appLaunchContextSetDesktopFunction *gi.Function
var appLaunchContextSetDesktopFunction_Once sync.Once

func appLaunchContextSetDesktopFunction_Set() error {
	var err error
	appLaunchContextSetDesktopFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextSetDesktopFunction, err = appLaunchContextStruct.InvokerNew("set_desktop")
	})
	return err
}

// SetDesktop is a representation of the C type gdk_app_launch_context_set_desktop.
func (recv *AppLaunchContext) SetDesktop(desktop int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(desktop)

	err := appLaunchContextSetDesktopFunction_Set()
	if err == nil {
		appLaunchContextSetDesktopFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appLaunchContextSetDisplayFunction *gi.Function
var appLaunchContextSetDisplayFunction_Once sync.Once

func appLaunchContextSetDisplayFunction_Set() error {
	var err error
	appLaunchContextSetDisplayFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextSetDisplayFunction, err = appLaunchContextStruct.InvokerNew("set_display")
	})
	return err
}

// SetDisplay is a representation of the C type gdk_app_launch_context_set_display.
func (recv *AppLaunchContext) SetDisplay(display *Display) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(display.Native)

	err := appLaunchContextSetDisplayFunction_Set()
	if err == nil {
		appLaunchContextSetDisplayFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_app_launch_context_set_icon' : parameter 'icon' of type 'Gio.Icon' not supported

var appLaunchContextSetIconNameFunction *gi.Function
var appLaunchContextSetIconNameFunction_Once sync.Once

func appLaunchContextSetIconNameFunction_Set() error {
	var err error
	appLaunchContextSetIconNameFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextSetIconNameFunction, err = appLaunchContextStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gdk_app_launch_context_set_icon_name.
func (recv *AppLaunchContext) SetIconName(iconName string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(iconName)

	err := appLaunchContextSetIconNameFunction_Set()
	if err == nil {
		appLaunchContextSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appLaunchContextSetScreenFunction *gi.Function
var appLaunchContextSetScreenFunction_Once sync.Once

func appLaunchContextSetScreenFunction_Set() error {
	var err error
	appLaunchContextSetScreenFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextSetScreenFunction, err = appLaunchContextStruct.InvokerNew("set_screen")
	})
	return err
}

// SetScreen is a representation of the C type gdk_app_launch_context_set_screen.
func (recv *AppLaunchContext) SetScreen(screen *Screen) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(screen.Native)

	err := appLaunchContextSetScreenFunction_Set()
	if err == nil {
		appLaunchContextSetScreenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var appLaunchContextSetTimestampFunction *gi.Function
var appLaunchContextSetTimestampFunction_Once sync.Once

func appLaunchContextSetTimestampFunction_Set() error {
	var err error
	appLaunchContextSetTimestampFunction_Once.Do(func() {
		err = appLaunchContextStruct_Set()
		if err != nil {
			return
		}
		appLaunchContextSetTimestampFunction, err = appLaunchContextStruct.InvokerNew("set_timestamp")
	})
	return err
}

// SetTimestamp is a representation of the C type gdk_app_launch_context_set_timestamp.
func (recv *AppLaunchContext) SetTimestamp(timestamp uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(timestamp)

	err := appLaunchContextSetTimestampFunction_Set()
	if err == nil {
		appLaunchContextSetTimestampFunction.Invoke(inArgs[:], nil)
	}

	return
}

var cursorStruct *gi.Struct
var cursorStruct_Once sync.Once

func cursorStruct_Set() error {
	var err error
	cursorStruct_Once.Do(func() {
		cursorStruct, err = gi.StructNew("Gdk", "Cursor")
	})
	return err
}

type Cursor struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gdk_cursor_new' : parameter 'cursor_type' of type 'CursorType' not supported

// UNSUPPORTED : C value 'gdk_cursor_new_for_display' : parameter 'cursor_type' of type 'CursorType' not supported

var cursorNewFromNameFunction *gi.Function
var cursorNewFromNameFunction_Once sync.Once

func cursorNewFromNameFunction_Set() error {
	var err error
	cursorNewFromNameFunction_Once.Do(func() {
		err = cursorStruct_Set()
		if err != nil {
			return
		}
		cursorNewFromNameFunction, err = cursorStruct.InvokerNew("new_from_name")
	})
	return err
}

// CursorNewFromName is a representation of the C type gdk_cursor_new_from_name.
func CursorNewFromName(display *Display, name string) *Cursor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(display.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := cursorNewFromNameFunction_Set()
	if err == nil {
		ret = cursorNewFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cursor{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_cursor_new_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_cursor_new_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_cursor_get_cursor_type' : return type 'CursorType' not supported

var cursorGetDisplayFunction *gi.Function
var cursorGetDisplayFunction_Once sync.Once

func cursorGetDisplayFunction_Set() error {
	var err error
	cursorGetDisplayFunction_Once.Do(func() {
		err = cursorStruct_Set()
		if err != nil {
			return
		}
		cursorGetDisplayFunction, err = cursorStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_cursor_get_display.
func (recv *Cursor) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := cursorGetDisplayFunction_Set()
	if err == nil {
		ret = cursorGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_cursor_get_image' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_cursor_get_surface' : return type 'cairo.Surface' not supported

var cursorRefFunction *gi.Function
var cursorRefFunction_Once sync.Once

func cursorRefFunction_Set() error {
	var err error
	cursorRefFunction_Once.Do(func() {
		err = cursorStruct_Set()
		if err != nil {
			return
		}
		cursorRefFunction, err = cursorStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type gdk_cursor_ref.
func (recv *Cursor) Ref() *Cursor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := cursorRefFunction_Set()
	if err == nil {
		ret = cursorRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cursor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var cursorUnrefFunction *gi.Function
var cursorUnrefFunction_Once sync.Once

func cursorUnrefFunction_Set() error {
	var err error
	cursorUnrefFunction_Once.Do(func() {
		err = cursorStruct_Set()
		if err != nil {
			return
		}
		cursorUnrefFunction, err = cursorStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type gdk_cursor_unref.
func (recv *Cursor) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := cursorUnrefFunction_Set()
	if err == nil {
		cursorUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var deviceStruct *gi.Struct
var deviceStruct_Once sync.Once

func deviceStruct_Set() error {
	var err error
	deviceStruct_Once.Do(func() {
		deviceStruct, err = gi.StructNew("Gdk", "Device")
	})
	return err
}

type Device struct {
	gobject.Object
}

var deviceGetAssociatedDeviceFunction *gi.Function
var deviceGetAssociatedDeviceFunction_Once sync.Once

func deviceGetAssociatedDeviceFunction_Set() error {
	var err error
	deviceGetAssociatedDeviceFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetAssociatedDeviceFunction, err = deviceStruct.InvokerNew("get_associated_device")
	})
	return err
}

// GetAssociatedDevice is a representation of the C type gdk_device_get_associated_device.
func (recv *Device) GetAssociatedDevice() *Device {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetAssociatedDeviceFunction_Set()
	if err == nil {
		ret = deviceGetAssociatedDeviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Device{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_get_axes' : return type 'AxisFlags' not supported

// UNSUPPORTED : C value 'gdk_device_get_axis' : parameter 'axes' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_device_get_axis_use' : return type 'AxisUse' not supported

// UNSUPPORTED : C value 'gdk_device_get_axis_value' : parameter 'axes' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_device_get_device_type' : return type 'DeviceType' not supported

var deviceGetDisplayFunction *gi.Function
var deviceGetDisplayFunction_Once sync.Once

func deviceGetDisplayFunction_Set() error {
	var err error
	deviceGetDisplayFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetDisplayFunction, err = deviceStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_device_get_display.
func (recv *Device) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetDisplayFunction_Set()
	if err == nil {
		ret = deviceGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

var deviceGetHasCursorFunction *gi.Function
var deviceGetHasCursorFunction_Once sync.Once

func deviceGetHasCursorFunction_Set() error {
	var err error
	deviceGetHasCursorFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetHasCursorFunction, err = deviceStruct.InvokerNew("get_has_cursor")
	})
	return err
}

// GetHasCursor is a representation of the C type gdk_device_get_has_cursor.
func (recv *Device) GetHasCursor() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetHasCursorFunction_Set()
	if err == nil {
		ret = deviceGetHasCursorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_get_history' : parameter 'events' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_device_get_key' : parameter 'modifiers' of type 'ModifierType' not supported

var deviceGetLastEventWindowFunction *gi.Function
var deviceGetLastEventWindowFunction_Once sync.Once

func deviceGetLastEventWindowFunction_Set() error {
	var err error
	deviceGetLastEventWindowFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetLastEventWindowFunction, err = deviceStruct.InvokerNew("get_last_event_window")
	})
	return err
}

// GetLastEventWindow is a representation of the C type gdk_device_get_last_event_window.
func (recv *Device) GetLastEventWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetLastEventWindowFunction_Set()
	if err == nil {
		ret = deviceGetLastEventWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_get_mode' : return type 'InputMode' not supported

var deviceGetNAxesFunction *gi.Function
var deviceGetNAxesFunction_Once sync.Once

func deviceGetNAxesFunction_Set() error {
	var err error
	deviceGetNAxesFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetNAxesFunction, err = deviceStruct.InvokerNew("get_n_axes")
	})
	return err
}

// GetNAxes is a representation of the C type gdk_device_get_n_axes.
func (recv *Device) GetNAxes() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetNAxesFunction_Set()
	if err == nil {
		ret = deviceGetNAxesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var deviceGetNKeysFunction *gi.Function
var deviceGetNKeysFunction_Once sync.Once

func deviceGetNKeysFunction_Set() error {
	var err error
	deviceGetNKeysFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetNKeysFunction, err = deviceStruct.InvokerNew("get_n_keys")
	})
	return err
}

// GetNKeys is a representation of the C type gdk_device_get_n_keys.
func (recv *Device) GetNKeys() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetNKeysFunction_Set()
	if err == nil {
		ret = deviceGetNKeysFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var deviceGetNameFunction *gi.Function
var deviceGetNameFunction_Once sync.Once

func deviceGetNameFunction_Set() error {
	var err error
	deviceGetNameFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetNameFunction, err = deviceStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gdk_device_get_name.
func (recv *Device) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetNameFunction_Set()
	if err == nil {
		ret = deviceGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var deviceGetPositionFunction *gi.Function
var deviceGetPositionFunction_Once sync.Once

func deviceGetPositionFunction_Set() error {
	var err error
	deviceGetPositionFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetPositionFunction, err = deviceStruct.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type gdk_device_get_position.
func (recv *Device) GetPosition() (*Screen, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [3]gi.Argument

	err := deviceGetPositionFunction_Set()
	if err == nil {
		deviceGetPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Screen{}
	out0.Native = outArgs[0].Pointer()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

var deviceGetPositionDoubleFunction *gi.Function
var deviceGetPositionDoubleFunction_Once sync.Once

func deviceGetPositionDoubleFunction_Set() error {
	var err error
	deviceGetPositionDoubleFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetPositionDoubleFunction, err = deviceStruct.InvokerNew("get_position_double")
	})
	return err
}

// GetPositionDouble is a representation of the C type gdk_device_get_position_double.
func (recv *Device) GetPositionDouble() (*Screen, float64, float64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [3]gi.Argument

	err := deviceGetPositionDoubleFunction_Set()
	if err == nil {
		deviceGetPositionDoubleFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Screen{}
	out0.Native = outArgs[0].Pointer()
	out1 := outArgs[1].Float64()
	out2 := outArgs[2].Float64()

	return out0, out1, out2
}

var deviceGetProductIdFunction *gi.Function
var deviceGetProductIdFunction_Once sync.Once

func deviceGetProductIdFunction_Set() error {
	var err error
	deviceGetProductIdFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetProductIdFunction, err = deviceStruct.InvokerNew("get_product_id")
	})
	return err
}

// GetProductId is a representation of the C type gdk_device_get_product_id.
func (recv *Device) GetProductId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetProductIdFunction_Set()
	if err == nil {
		ret = deviceGetProductIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var deviceGetSeatFunction *gi.Function
var deviceGetSeatFunction_Once sync.Once

func deviceGetSeatFunction_Set() error {
	var err error
	deviceGetSeatFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetSeatFunction, err = deviceStruct.InvokerNew("get_seat")
	})
	return err
}

// GetSeat is a representation of the C type gdk_device_get_seat.
func (recv *Device) GetSeat() *Seat {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetSeatFunction_Set()
	if err == nil {
		ret = deviceGetSeatFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Seat{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_get_source' : return type 'InputSource' not supported

// UNSUPPORTED : C value 'gdk_device_get_state' : parameter 'axes' of type 'nil' not supported

var deviceGetVendorIdFunction *gi.Function
var deviceGetVendorIdFunction_Once sync.Once

func deviceGetVendorIdFunction_Set() error {
	var err error
	deviceGetVendorIdFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetVendorIdFunction, err = deviceStruct.InvokerNew("get_vendor_id")
	})
	return err
}

// GetVendorId is a representation of the C type gdk_device_get_vendor_id.
func (recv *Device) GetVendorId() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceGetVendorIdFunction_Set()
	if err == nil {
		ret = deviceGetVendorIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var deviceGetWindowAtPositionFunction *gi.Function
var deviceGetWindowAtPositionFunction_Once sync.Once

func deviceGetWindowAtPositionFunction_Set() error {
	var err error
	deviceGetWindowAtPositionFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetWindowAtPositionFunction, err = deviceStruct.InvokerNew("get_window_at_position")
	})
	return err
}

// GetWindowAtPosition is a representation of the C type gdk_device_get_window_at_position.
func (recv *Device) GetWindowAtPosition() (*Window, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := deviceGetWindowAtPositionFunction_Set()
	if err == nil {
		ret = deviceGetWindowAtPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var deviceGetWindowAtPositionDoubleFunction *gi.Function
var deviceGetWindowAtPositionDoubleFunction_Once sync.Once

func deviceGetWindowAtPositionDoubleFunction_Set() error {
	var err error
	deviceGetWindowAtPositionDoubleFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceGetWindowAtPositionDoubleFunction, err = deviceStruct.InvokerNew("get_window_at_position_double")
	})
	return err
}

// GetWindowAtPositionDouble is a representation of the C type gdk_device_get_window_at_position_double.
func (recv *Device) GetWindowAtPositionDouble() (*Window, float64, float64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := deviceGetWindowAtPositionDoubleFunction_Set()
	if err == nil {
		ret = deviceGetWindowAtPositionDoubleFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()
	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return retGo, out0, out1
}

// UNSUPPORTED : C value 'gdk_device_grab' : parameter 'grab_ownership' of type 'GrabOwnership' not supported

// UNSUPPORTED : C value 'gdk_device_list_axes' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_device_list_slave_devices' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_device_set_axis_use' : parameter 'use' of type 'AxisUse' not supported

// UNSUPPORTED : C value 'gdk_device_set_key' : parameter 'modifiers' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_device_set_mode' : parameter 'mode' of type 'InputMode' not supported

var deviceUngrabFunction *gi.Function
var deviceUngrabFunction_Once sync.Once

func deviceUngrabFunction_Set() error {
	var err error
	deviceUngrabFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceUngrabFunction, err = deviceStruct.InvokerNew("ungrab")
	})
	return err
}

// Ungrab is a representation of the C type gdk_device_ungrab.
func (recv *Device) Ungrab(time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(time)

	err := deviceUngrabFunction_Set()
	if err == nil {
		deviceUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var deviceWarpFunction *gi.Function
var deviceWarpFunction_Once sync.Once

func deviceWarpFunction_Set() error {
	var err error
	deviceWarpFunction_Once.Do(func() {
		err = deviceStruct_Set()
		if err != nil {
			return
		}
		deviceWarpFunction, err = deviceStruct.InvokerNew("warp")
	})
	return err
}

// Warp is a representation of the C type gdk_device_warp.
func (recv *Device) Warp(screen *Screen, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(screen.Native)
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := deviceWarpFunction_Set()
	if err == nil {
		deviceWarpFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DeviceStruct creates an uninitialised Device.
func DeviceStruct() *Device {
	err := deviceStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Device{}
	structGo.Native = deviceStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDevice)
	return structGo
}
func finalizeDevice(obj *Device) {
	deviceStruct.Free(obj.Native)
}

var deviceManagerStruct *gi.Struct
var deviceManagerStruct_Once sync.Once

func deviceManagerStruct_Set() error {
	var err error
	deviceManagerStruct_Once.Do(func() {
		deviceManagerStruct, err = gi.StructNew("Gdk", "DeviceManager")
	})
	return err
}

type DeviceManager struct {
	gobject.Object
}

var deviceManagerGetClientPointerFunction *gi.Function
var deviceManagerGetClientPointerFunction_Once sync.Once

func deviceManagerGetClientPointerFunction_Set() error {
	var err error
	deviceManagerGetClientPointerFunction_Once.Do(func() {
		err = deviceManagerStruct_Set()
		if err != nil {
			return
		}
		deviceManagerGetClientPointerFunction, err = deviceManagerStruct.InvokerNew("get_client_pointer")
	})
	return err
}

// GetClientPointer is a representation of the C type gdk_device_manager_get_client_pointer.
func (recv *DeviceManager) GetClientPointer() *Device {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceManagerGetClientPointerFunction_Set()
	if err == nil {
		ret = deviceManagerGetClientPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Device{}
	retGo.Native = ret.Pointer()

	return retGo
}

var deviceManagerGetDisplayFunction *gi.Function
var deviceManagerGetDisplayFunction_Once sync.Once

func deviceManagerGetDisplayFunction_Set() error {
	var err error
	deviceManagerGetDisplayFunction_Once.Do(func() {
		err = deviceManagerStruct_Set()
		if err != nil {
			return
		}
		deviceManagerGetDisplayFunction, err = deviceManagerStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_device_manager_get_display.
func (recv *DeviceManager) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceManagerGetDisplayFunction_Set()
	if err == nil {
		ret = deviceManagerGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_manager_list_devices' : parameter 'type' of type 'DeviceType' not supported

// DeviceManagerStruct creates an uninitialised DeviceManager.
func DeviceManagerStruct() *DeviceManager {
	err := deviceManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DeviceManager{}
	structGo.Native = deviceManagerStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDeviceManager)
	return structGo
}
func finalizeDeviceManager(obj *DeviceManager) {
	deviceManagerStruct.Free(obj.Native)
}

var deviceToolStruct *gi.Struct
var deviceToolStruct_Once sync.Once

func deviceToolStruct_Set() error {
	var err error
	deviceToolStruct_Once.Do(func() {
		deviceToolStruct, err = gi.StructNew("Gdk", "DeviceTool")
	})
	return err
}

type DeviceTool struct {
	gobject.Object
}

var deviceToolGetHardwareIdFunction *gi.Function
var deviceToolGetHardwareIdFunction_Once sync.Once

func deviceToolGetHardwareIdFunction_Set() error {
	var err error
	deviceToolGetHardwareIdFunction_Once.Do(func() {
		err = deviceToolStruct_Set()
		if err != nil {
			return
		}
		deviceToolGetHardwareIdFunction, err = deviceToolStruct.InvokerNew("get_hardware_id")
	})
	return err
}

// GetHardwareId is a representation of the C type gdk_device_tool_get_hardware_id.
func (recv *DeviceTool) GetHardwareId() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceToolGetHardwareIdFunction_Set()
	if err == nil {
		ret = deviceToolGetHardwareIdFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

var deviceToolGetSerialFunction *gi.Function
var deviceToolGetSerialFunction_Once sync.Once

func deviceToolGetSerialFunction_Set() error {
	var err error
	deviceToolGetSerialFunction_Once.Do(func() {
		err = deviceToolStruct_Set()
		if err != nil {
			return
		}
		deviceToolGetSerialFunction, err = deviceToolStruct.InvokerNew("get_serial")
	})
	return err
}

// GetSerial is a representation of the C type gdk_device_tool_get_serial.
func (recv *DeviceTool) GetSerial() uint64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := deviceToolGetSerialFunction_Set()
	if err == nil {
		ret = deviceToolGetSerialFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint64()

	return retGo
}

// UNSUPPORTED : C value 'gdk_device_tool_get_tool_type' : return type 'DeviceToolType' not supported

// DeviceToolStruct creates an uninitialised DeviceTool.
func DeviceToolStruct() *DeviceTool {
	err := deviceToolStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DeviceTool{}
	structGo.Native = deviceToolStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDeviceTool)
	return structGo
}
func finalizeDeviceTool(obj *DeviceTool) {
	deviceToolStruct.Free(obj.Native)
}

var displayStruct *gi.Struct
var displayStruct_Once sync.Once

func displayStruct_Set() error {
	var err error
	displayStruct_Once.Do(func() {
		displayStruct, err = gi.StructNew("Gdk", "Display")
	})
	return err
}

type Display struct {
	gobject.Object
}

var displayBeepFunction *gi.Function
var displayBeepFunction_Once sync.Once

func displayBeepFunction_Set() error {
	var err error
	displayBeepFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayBeepFunction, err = displayStruct.InvokerNew("beep")
	})
	return err
}

// Beep is a representation of the C type gdk_display_beep.
func (recv *Display) Beep() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := displayBeepFunction_Set()
	if err == nil {
		displayBeepFunction.Invoke(inArgs[:], nil)
	}

	return
}

var displayCloseFunction *gi.Function
var displayCloseFunction_Once sync.Once

func displayCloseFunction_Set() error {
	var err error
	displayCloseFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayCloseFunction, err = displayStruct.InvokerNew("close")
	})
	return err
}

// Close is a representation of the C type gdk_display_close.
func (recv *Display) Close() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := displayCloseFunction_Set()
	if err == nil {
		displayCloseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var displayDeviceIsGrabbedFunction *gi.Function
var displayDeviceIsGrabbedFunction_Once sync.Once

func displayDeviceIsGrabbedFunction_Set() error {
	var err error
	displayDeviceIsGrabbedFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayDeviceIsGrabbedFunction, err = displayStruct.InvokerNew("device_is_grabbed")
	})
	return err
}

// DeviceIsGrabbed is a representation of the C type gdk_display_device_is_grabbed.
func (recv *Display) DeviceIsGrabbed(device *Device) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(device.Native)

	var ret gi.Argument

	err := displayDeviceIsGrabbedFunction_Set()
	if err == nil {
		ret = displayDeviceIsGrabbedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displayFlushFunction *gi.Function
var displayFlushFunction_Once sync.Once

func displayFlushFunction_Set() error {
	var err error
	displayFlushFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayFlushFunction, err = displayStruct.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type gdk_display_flush.
func (recv *Display) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := displayFlushFunction_Set()
	if err == nil {
		displayFlushFunction.Invoke(inArgs[:], nil)
	}

	return
}

var displayGetAppLaunchContextFunction *gi.Function
var displayGetAppLaunchContextFunction_Once sync.Once

func displayGetAppLaunchContextFunction_Set() error {
	var err error
	displayGetAppLaunchContextFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetAppLaunchContextFunction, err = displayStruct.InvokerNew("get_app_launch_context")
	})
	return err
}

// GetAppLaunchContext is a representation of the C type gdk_display_get_app_launch_context.
func (recv *Display) GetAppLaunchContext() *AppLaunchContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetAppLaunchContextFunction_Set()
	if err == nil {
		ret = displayGetAppLaunchContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &AppLaunchContext{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetDefaultCursorSizeFunction *gi.Function
var displayGetDefaultCursorSizeFunction_Once sync.Once

func displayGetDefaultCursorSizeFunction_Set() error {
	var err error
	displayGetDefaultCursorSizeFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetDefaultCursorSizeFunction, err = displayStruct.InvokerNew("get_default_cursor_size")
	})
	return err
}

// GetDefaultCursorSize is a representation of the C type gdk_display_get_default_cursor_size.
func (recv *Display) GetDefaultCursorSize() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetDefaultCursorSizeFunction_Set()
	if err == nil {
		ret = displayGetDefaultCursorSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var displayGetDefaultGroupFunction *gi.Function
var displayGetDefaultGroupFunction_Once sync.Once

func displayGetDefaultGroupFunction_Set() error {
	var err error
	displayGetDefaultGroupFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetDefaultGroupFunction, err = displayStruct.InvokerNew("get_default_group")
	})
	return err
}

// GetDefaultGroup is a representation of the C type gdk_display_get_default_group.
func (recv *Display) GetDefaultGroup() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetDefaultGroupFunction_Set()
	if err == nil {
		ret = displayGetDefaultGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetDefaultScreenFunction *gi.Function
var displayGetDefaultScreenFunction_Once sync.Once

func displayGetDefaultScreenFunction_Set() error {
	var err error
	displayGetDefaultScreenFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetDefaultScreenFunction, err = displayStruct.InvokerNew("get_default_screen")
	})
	return err
}

// GetDefaultScreen is a representation of the C type gdk_display_get_default_screen.
func (recv *Display) GetDefaultScreen() *Screen {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetDefaultScreenFunction_Set()
	if err == nil {
		ret = displayGetDefaultScreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Screen{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetDefaultSeatFunction *gi.Function
var displayGetDefaultSeatFunction_Once sync.Once

func displayGetDefaultSeatFunction_Set() error {
	var err error
	displayGetDefaultSeatFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetDefaultSeatFunction, err = displayStruct.InvokerNew("get_default_seat")
	})
	return err
}

// GetDefaultSeat is a representation of the C type gdk_display_get_default_seat.
func (recv *Display) GetDefaultSeat() *Seat {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetDefaultSeatFunction_Set()
	if err == nil {
		ret = displayGetDefaultSeatFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Seat{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetDeviceManagerFunction *gi.Function
var displayGetDeviceManagerFunction_Once sync.Once

func displayGetDeviceManagerFunction_Set() error {
	var err error
	displayGetDeviceManagerFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetDeviceManagerFunction, err = displayStruct.InvokerNew("get_device_manager")
	})
	return err
}

// GetDeviceManager is a representation of the C type gdk_display_get_device_manager.
func (recv *Display) GetDeviceManager() *DeviceManager {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetDeviceManagerFunction_Set()
	if err == nil {
		ret = displayGetDeviceManagerFunction.Invoke(inArgs[:], nil)
	}

	retGo := &DeviceManager{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_display_get_event' : return type 'Event' not supported

var displayGetMaximalCursorSizeFunction *gi.Function
var displayGetMaximalCursorSizeFunction_Once sync.Once

func displayGetMaximalCursorSizeFunction_Set() error {
	var err error
	displayGetMaximalCursorSizeFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetMaximalCursorSizeFunction, err = displayStruct.InvokerNew("get_maximal_cursor_size")
	})
	return err
}

// GetMaximalCursorSize is a representation of the C type gdk_display_get_maximal_cursor_size.
func (recv *Display) GetMaximalCursorSize() (uint32, uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := displayGetMaximalCursorSizeFunction_Set()
	if err == nil {
		displayGetMaximalCursorSizeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Uint32()

	return out0, out1
}

var displayGetMonitorFunction *gi.Function
var displayGetMonitorFunction_Once sync.Once

func displayGetMonitorFunction_Set() error {
	var err error
	displayGetMonitorFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetMonitorFunction, err = displayStruct.InvokerNew("get_monitor")
	})
	return err
}

// GetMonitor is a representation of the C type gdk_display_get_monitor.
func (recv *Display) GetMonitor(monitorNum int32) *Monitor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var ret gi.Argument

	err := displayGetMonitorFunction_Set()
	if err == nil {
		ret = displayGetMonitorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Monitor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetMonitorAtPointFunction *gi.Function
var displayGetMonitorAtPointFunction_Once sync.Once

func displayGetMonitorAtPointFunction_Set() error {
	var err error
	displayGetMonitorAtPointFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetMonitorAtPointFunction, err = displayStruct.InvokerNew("get_monitor_at_point")
	})
	return err
}

// GetMonitorAtPoint is a representation of the C type gdk_display_get_monitor_at_point.
func (recv *Display) GetMonitorAtPoint(x int32, y int32) *Monitor {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var ret gi.Argument

	err := displayGetMonitorAtPointFunction_Set()
	if err == nil {
		ret = displayGetMonitorAtPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Monitor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetMonitorAtWindowFunction *gi.Function
var displayGetMonitorAtWindowFunction_Once sync.Once

func displayGetMonitorAtWindowFunction_Set() error {
	var err error
	displayGetMonitorAtWindowFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetMonitorAtWindowFunction, err = displayStruct.InvokerNew("get_monitor_at_window")
	})
	return err
}

// GetMonitorAtWindow is a representation of the C type gdk_display_get_monitor_at_window.
func (recv *Display) GetMonitorAtWindow(window *Window) *Monitor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(window.Native)

	var ret gi.Argument

	err := displayGetMonitorAtWindowFunction_Set()
	if err == nil {
		ret = displayGetMonitorAtWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Monitor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetNMonitorsFunction *gi.Function
var displayGetNMonitorsFunction_Once sync.Once

func displayGetNMonitorsFunction_Set() error {
	var err error
	displayGetNMonitorsFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetNMonitorsFunction, err = displayStruct.InvokerNew("get_n_monitors")
	})
	return err
}

// GetNMonitors is a representation of the C type gdk_display_get_n_monitors.
func (recv *Display) GetNMonitors() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetNMonitorsFunction_Set()
	if err == nil {
		ret = displayGetNMonitorsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var displayGetNScreensFunction *gi.Function
var displayGetNScreensFunction_Once sync.Once

func displayGetNScreensFunction_Set() error {
	var err error
	displayGetNScreensFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetNScreensFunction, err = displayStruct.InvokerNew("get_n_screens")
	})
	return err
}

// GetNScreens is a representation of the C type gdk_display_get_n_screens.
func (recv *Display) GetNScreens() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetNScreensFunction_Set()
	if err == nil {
		ret = displayGetNScreensFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var displayGetNameFunction *gi.Function
var displayGetNameFunction_Once sync.Once

func displayGetNameFunction_Set() error {
	var err error
	displayGetNameFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetNameFunction, err = displayStruct.InvokerNew("get_name")
	})
	return err
}

// GetName is a representation of the C type gdk_display_get_name.
func (recv *Display) GetName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetNameFunction_Set()
	if err == nil {
		ret = displayGetNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// UNSUPPORTED : C value 'gdk_display_get_pointer' : parameter 'mask' of type 'ModifierType' not supported

var displayGetPrimaryMonitorFunction *gi.Function
var displayGetPrimaryMonitorFunction_Once sync.Once

func displayGetPrimaryMonitorFunction_Set() error {
	var err error
	displayGetPrimaryMonitorFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetPrimaryMonitorFunction, err = displayStruct.InvokerNew("get_primary_monitor")
	})
	return err
}

// GetPrimaryMonitor is a representation of the C type gdk_display_get_primary_monitor.
func (recv *Display) GetPrimaryMonitor() *Monitor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayGetPrimaryMonitorFunction_Set()
	if err == nil {
		ret = displayGetPrimaryMonitorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Monitor{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetScreenFunction *gi.Function
var displayGetScreenFunction_Once sync.Once

func displayGetScreenFunction_Set() error {
	var err error
	displayGetScreenFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetScreenFunction, err = displayStruct.InvokerNew("get_screen")
	})
	return err
}

// GetScreen is a representation of the C type gdk_display_get_screen.
func (recv *Display) GetScreen(screenNum int32) *Screen {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(screenNum)

	var ret gi.Argument

	err := displayGetScreenFunction_Set()
	if err == nil {
		ret = displayGetScreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Screen{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayGetWindowAtPointerFunction *gi.Function
var displayGetWindowAtPointerFunction_Once sync.Once

func displayGetWindowAtPointerFunction_Set() error {
	var err error
	displayGetWindowAtPointerFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayGetWindowAtPointerFunction, err = displayStruct.InvokerNew("get_window_at_pointer")
	})
	return err
}

// GetWindowAtPointer is a representation of the C type gdk_display_get_window_at_pointer.
func (recv *Display) GetWindowAtPointer() (*Window, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := displayGetWindowAtPointerFunction_Set()
	if err == nil {
		ret = displayGetWindowAtPointerFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var displayHasPendingFunction *gi.Function
var displayHasPendingFunction_Once sync.Once

func displayHasPendingFunction_Set() error {
	var err error
	displayHasPendingFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayHasPendingFunction, err = displayStruct.InvokerNew("has_pending")
	})
	return err
}

// HasPending is a representation of the C type gdk_display_has_pending.
func (recv *Display) HasPending() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayHasPendingFunction_Set()
	if err == nil {
		ret = displayHasPendingFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displayIsClosedFunction *gi.Function
var displayIsClosedFunction_Once sync.Once

func displayIsClosedFunction_Set() error {
	var err error
	displayIsClosedFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayIsClosedFunction, err = displayStruct.InvokerNew("is_closed")
	})
	return err
}

// IsClosed is a representation of the C type gdk_display_is_closed.
func (recv *Display) IsClosed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayIsClosedFunction_Set()
	if err == nil {
		ret = displayIsClosedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displayKeyboardUngrabFunction *gi.Function
var displayKeyboardUngrabFunction_Once sync.Once

func displayKeyboardUngrabFunction_Set() error {
	var err error
	displayKeyboardUngrabFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayKeyboardUngrabFunction, err = displayStruct.InvokerNew("keyboard_ungrab")
	})
	return err
}

// KeyboardUngrab is a representation of the C type gdk_display_keyboard_ungrab.
func (recv *Display) KeyboardUngrab(time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(time)

	err := displayKeyboardUngrabFunction_Set()
	if err == nil {
		displayKeyboardUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_display_list_devices' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_display_list_seats' : return type 'GLib.List' not supported

var displayNotifyStartupCompleteFunction *gi.Function
var displayNotifyStartupCompleteFunction_Once sync.Once

func displayNotifyStartupCompleteFunction_Set() error {
	var err error
	displayNotifyStartupCompleteFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayNotifyStartupCompleteFunction, err = displayStruct.InvokerNew("notify_startup_complete")
	})
	return err
}

// NotifyStartupComplete is a representation of the C type gdk_display_notify_startup_complete.
func (recv *Display) NotifyStartupComplete(startupId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(startupId)

	err := displayNotifyStartupCompleteFunction_Set()
	if err == nil {
		displayNotifyStartupCompleteFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_display_peek_event' : return type 'Event' not supported

var displayPointerIsGrabbedFunction *gi.Function
var displayPointerIsGrabbedFunction_Once sync.Once

func displayPointerIsGrabbedFunction_Set() error {
	var err error
	displayPointerIsGrabbedFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayPointerIsGrabbedFunction, err = displayStruct.InvokerNew("pointer_is_grabbed")
	})
	return err
}

// PointerIsGrabbed is a representation of the C type gdk_display_pointer_is_grabbed.
func (recv *Display) PointerIsGrabbed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayPointerIsGrabbedFunction_Set()
	if err == nil {
		ret = displayPointerIsGrabbedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displayPointerUngrabFunction *gi.Function
var displayPointerUngrabFunction_Once sync.Once

func displayPointerUngrabFunction_Set() error {
	var err error
	displayPointerUngrabFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayPointerUngrabFunction, err = displayStruct.InvokerNew("pointer_ungrab")
	})
	return err
}

// PointerUngrab is a representation of the C type gdk_display_pointer_ungrab.
func (recv *Display) PointerUngrab(time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(time)

	err := displayPointerUngrabFunction_Set()
	if err == nil {
		displayPointerUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_display_put_event' : parameter 'event' of type 'Event' not supported

var displayRequestSelectionNotificationFunction *gi.Function
var displayRequestSelectionNotificationFunction_Once sync.Once

func displayRequestSelectionNotificationFunction_Set() error {
	var err error
	displayRequestSelectionNotificationFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayRequestSelectionNotificationFunction, err = displayStruct.InvokerNew("request_selection_notification")
	})
	return err
}

// RequestSelectionNotification is a representation of the C type gdk_display_request_selection_notification.
func (recv *Display) RequestSelectionNotification(selection *Atom) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(selection.Native)

	var ret gi.Argument

	err := displayRequestSelectionNotificationFunction_Set()
	if err == nil {
		ret = displayRequestSelectionNotificationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySetDoubleClickDistanceFunction *gi.Function
var displaySetDoubleClickDistanceFunction_Once sync.Once

func displaySetDoubleClickDistanceFunction_Set() error {
	var err error
	displaySetDoubleClickDistanceFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySetDoubleClickDistanceFunction, err = displayStruct.InvokerNew("set_double_click_distance")
	})
	return err
}

// SetDoubleClickDistance is a representation of the C type gdk_display_set_double_click_distance.
func (recv *Display) SetDoubleClickDistance(distance uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(distance)

	err := displaySetDoubleClickDistanceFunction_Set()
	if err == nil {
		displaySetDoubleClickDistanceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var displaySetDoubleClickTimeFunction *gi.Function
var displaySetDoubleClickTimeFunction_Once sync.Once

func displaySetDoubleClickTimeFunction_Set() error {
	var err error
	displaySetDoubleClickTimeFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySetDoubleClickTimeFunction, err = displayStruct.InvokerNew("set_double_click_time")
	})
	return err
}

// SetDoubleClickTime is a representation of the C type gdk_display_set_double_click_time.
func (recv *Display) SetDoubleClickTime(msec uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(msec)

	err := displaySetDoubleClickTimeFunction_Set()
	if err == nil {
		displaySetDoubleClickTimeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_display_store_clipboard' : parameter 'targets' of type 'nil' not supported

var displaySupportsClipboardPersistenceFunction *gi.Function
var displaySupportsClipboardPersistenceFunction_Once sync.Once

func displaySupportsClipboardPersistenceFunction_Set() error {
	var err error
	displaySupportsClipboardPersistenceFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsClipboardPersistenceFunction, err = displayStruct.InvokerNew("supports_clipboard_persistence")
	})
	return err
}

// SupportsClipboardPersistence is a representation of the C type gdk_display_supports_clipboard_persistence.
func (recv *Display) SupportsClipboardPersistence() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsClipboardPersistenceFunction_Set()
	if err == nil {
		ret = displaySupportsClipboardPersistenceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsCompositeFunction *gi.Function
var displaySupportsCompositeFunction_Once sync.Once

func displaySupportsCompositeFunction_Set() error {
	var err error
	displaySupportsCompositeFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsCompositeFunction, err = displayStruct.InvokerNew("supports_composite")
	})
	return err
}

// SupportsComposite is a representation of the C type gdk_display_supports_composite.
func (recv *Display) SupportsComposite() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsCompositeFunction_Set()
	if err == nil {
		ret = displaySupportsCompositeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsCursorAlphaFunction *gi.Function
var displaySupportsCursorAlphaFunction_Once sync.Once

func displaySupportsCursorAlphaFunction_Set() error {
	var err error
	displaySupportsCursorAlphaFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsCursorAlphaFunction, err = displayStruct.InvokerNew("supports_cursor_alpha")
	})
	return err
}

// SupportsCursorAlpha is a representation of the C type gdk_display_supports_cursor_alpha.
func (recv *Display) SupportsCursorAlpha() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsCursorAlphaFunction_Set()
	if err == nil {
		ret = displaySupportsCursorAlphaFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsCursorColorFunction *gi.Function
var displaySupportsCursorColorFunction_Once sync.Once

func displaySupportsCursorColorFunction_Set() error {
	var err error
	displaySupportsCursorColorFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsCursorColorFunction, err = displayStruct.InvokerNew("supports_cursor_color")
	})
	return err
}

// SupportsCursorColor is a representation of the C type gdk_display_supports_cursor_color.
func (recv *Display) SupportsCursorColor() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsCursorColorFunction_Set()
	if err == nil {
		ret = displaySupportsCursorColorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsInputShapesFunction *gi.Function
var displaySupportsInputShapesFunction_Once sync.Once

func displaySupportsInputShapesFunction_Set() error {
	var err error
	displaySupportsInputShapesFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsInputShapesFunction, err = displayStruct.InvokerNew("supports_input_shapes")
	})
	return err
}

// SupportsInputShapes is a representation of the C type gdk_display_supports_input_shapes.
func (recv *Display) SupportsInputShapes() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsInputShapesFunction_Set()
	if err == nil {
		ret = displaySupportsInputShapesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsSelectionNotificationFunction *gi.Function
var displaySupportsSelectionNotificationFunction_Once sync.Once

func displaySupportsSelectionNotificationFunction_Set() error {
	var err error
	displaySupportsSelectionNotificationFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsSelectionNotificationFunction, err = displayStruct.InvokerNew("supports_selection_notification")
	})
	return err
}

// SupportsSelectionNotification is a representation of the C type gdk_display_supports_selection_notification.
func (recv *Display) SupportsSelectionNotification() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsSelectionNotificationFunction_Set()
	if err == nil {
		ret = displaySupportsSelectionNotificationFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySupportsShapesFunction *gi.Function
var displaySupportsShapesFunction_Once sync.Once

func displaySupportsShapesFunction_Set() error {
	var err error
	displaySupportsShapesFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySupportsShapesFunction, err = displayStruct.InvokerNew("supports_shapes")
	})
	return err
}

// SupportsShapes is a representation of the C type gdk_display_supports_shapes.
func (recv *Display) SupportsShapes() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displaySupportsShapesFunction_Set()
	if err == nil {
		ret = displaySupportsShapesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var displaySyncFunction *gi.Function
var displaySyncFunction_Once sync.Once

func displaySyncFunction_Set() error {
	var err error
	displaySyncFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displaySyncFunction, err = displayStruct.InvokerNew("sync")
	})
	return err
}

// Sync is a representation of the C type gdk_display_sync.
func (recv *Display) Sync() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := displaySyncFunction_Set()
	if err == nil {
		displaySyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

var displayWarpPointerFunction *gi.Function
var displayWarpPointerFunction_Once sync.Once

func displayWarpPointerFunction_Set() error {
	var err error
	displayWarpPointerFunction_Once.Do(func() {
		err = displayStruct_Set()
		if err != nil {
			return
		}
		displayWarpPointerFunction, err = displayStruct.InvokerNew("warp_pointer")
	})
	return err
}

// WarpPointer is a representation of the C type gdk_display_warp_pointer.
func (recv *Display) WarpPointer(screen *Screen, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(screen.Native)
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := displayWarpPointerFunction_Set()
	if err == nil {
		displayWarpPointerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DisplayStruct creates an uninitialised Display.
func DisplayStruct() *Display {
	err := displayStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Display{}
	structGo.Native = displayStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDisplay)
	return structGo
}
func finalizeDisplay(obj *Display) {
	displayStruct.Free(obj.Native)
}

var displayManagerStruct *gi.Struct
var displayManagerStruct_Once sync.Once

func displayManagerStruct_Set() error {
	var err error
	displayManagerStruct_Once.Do(func() {
		displayManagerStruct, err = gi.StructNew("Gdk", "DisplayManager")
	})
	return err
}

type DisplayManager struct {
	gobject.Object
}

var displayManagerGetDefaultDisplayFunction *gi.Function
var displayManagerGetDefaultDisplayFunction_Once sync.Once

func displayManagerGetDefaultDisplayFunction_Set() error {
	var err error
	displayManagerGetDefaultDisplayFunction_Once.Do(func() {
		err = displayManagerStruct_Set()
		if err != nil {
			return
		}
		displayManagerGetDefaultDisplayFunction, err = displayManagerStruct.InvokerNew("get_default_display")
	})
	return err
}

// GetDefaultDisplay is a representation of the C type gdk_display_manager_get_default_display.
func (recv *DisplayManager) GetDefaultDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := displayManagerGetDefaultDisplayFunction_Set()
	if err == nil {
		ret = displayManagerGetDefaultDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_display_manager_list_displays' : return type 'GLib.SList' not supported

var displayManagerOpenDisplayFunction *gi.Function
var displayManagerOpenDisplayFunction_Once sync.Once

func displayManagerOpenDisplayFunction_Set() error {
	var err error
	displayManagerOpenDisplayFunction_Once.Do(func() {
		err = displayManagerStruct_Set()
		if err != nil {
			return
		}
		displayManagerOpenDisplayFunction, err = displayManagerStruct.InvokerNew("open_display")
	})
	return err
}

// OpenDisplay is a representation of the C type gdk_display_manager_open_display.
func (recv *DisplayManager) OpenDisplay(name string) *Display {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	var ret gi.Argument

	err := displayManagerOpenDisplayFunction_Set()
	if err == nil {
		ret = displayManagerOpenDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

var displayManagerSetDefaultDisplayFunction *gi.Function
var displayManagerSetDefaultDisplayFunction_Once sync.Once

func displayManagerSetDefaultDisplayFunction_Set() error {
	var err error
	displayManagerSetDefaultDisplayFunction_Once.Do(func() {
		err = displayManagerStruct_Set()
		if err != nil {
			return
		}
		displayManagerSetDefaultDisplayFunction, err = displayManagerStruct.InvokerNew("set_default_display")
	})
	return err
}

// SetDefaultDisplay is a representation of the C type gdk_display_manager_set_default_display.
func (recv *DisplayManager) SetDefaultDisplay(display *Display) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(display.Native)

	err := displayManagerSetDefaultDisplayFunction_Set()
	if err == nil {
		displayManagerSetDefaultDisplayFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DisplayManagerStruct creates an uninitialised DisplayManager.
func DisplayManagerStruct() *DisplayManager {
	err := displayManagerStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DisplayManager{}
	structGo.Native = displayManagerStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDisplayManager)
	return structGo
}
func finalizeDisplayManager(obj *DisplayManager) {
	displayManagerStruct.Free(obj.Native)
}

var dragContextStruct *gi.Struct
var dragContextStruct_Once sync.Once

func dragContextStruct_Set() error {
	var err error
	dragContextStruct_Once.Do(func() {
		dragContextStruct, err = gi.StructNew("Gdk", "DragContext")
	})
	return err
}

type DragContext struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gdk_drag_context_get_actions' : return type 'DragAction' not supported

var dragContextGetDestWindowFunction *gi.Function
var dragContextGetDestWindowFunction_Once sync.Once

func dragContextGetDestWindowFunction_Set() error {
	var err error
	dragContextGetDestWindowFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextGetDestWindowFunction, err = dragContextStruct.InvokerNew("get_dest_window")
	})
	return err
}

// GetDestWindow is a representation of the C type gdk_drag_context_get_dest_window.
func (recv *DragContext) GetDestWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dragContextGetDestWindowFunction_Set()
	if err == nil {
		ret = dragContextGetDestWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var dragContextGetDeviceFunction *gi.Function
var dragContextGetDeviceFunction_Once sync.Once

func dragContextGetDeviceFunction_Set() error {
	var err error
	dragContextGetDeviceFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextGetDeviceFunction, err = dragContextStruct.InvokerNew("get_device")
	})
	return err
}

// GetDevice is a representation of the C type gdk_drag_context_get_device.
func (recv *DragContext) GetDevice() *Device {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dragContextGetDeviceFunction_Set()
	if err == nil {
		ret = dragContextGetDeviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Device{}
	retGo.Native = ret.Pointer()

	return retGo
}

var dragContextGetDragWindowFunction *gi.Function
var dragContextGetDragWindowFunction_Once sync.Once

func dragContextGetDragWindowFunction_Set() error {
	var err error
	dragContextGetDragWindowFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextGetDragWindowFunction, err = dragContextStruct.InvokerNew("get_drag_window")
	})
	return err
}

// GetDragWindow is a representation of the C type gdk_drag_context_get_drag_window.
func (recv *DragContext) GetDragWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dragContextGetDragWindowFunction_Set()
	if err == nil {
		ret = dragContextGetDragWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_drag_context_get_protocol' : return type 'DragProtocol' not supported

// UNSUPPORTED : C value 'gdk_drag_context_get_selected_action' : return type 'DragAction' not supported

var dragContextGetSourceWindowFunction *gi.Function
var dragContextGetSourceWindowFunction_Once sync.Once

func dragContextGetSourceWindowFunction_Set() error {
	var err error
	dragContextGetSourceWindowFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextGetSourceWindowFunction, err = dragContextStruct.InvokerNew("get_source_window")
	})
	return err
}

// GetSourceWindow is a representation of the C type gdk_drag_context_get_source_window.
func (recv *DragContext) GetSourceWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := dragContextGetSourceWindowFunction_Set()
	if err == nil {
		ret = dragContextGetSourceWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_drag_context_get_suggested_action' : return type 'DragAction' not supported

// UNSUPPORTED : C value 'gdk_drag_context_list_targets' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_drag_context_manage_dnd' : parameter 'actions' of type 'DragAction' not supported

var dragContextSetDeviceFunction *gi.Function
var dragContextSetDeviceFunction_Once sync.Once

func dragContextSetDeviceFunction_Set() error {
	var err error
	dragContextSetDeviceFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextSetDeviceFunction, err = dragContextStruct.InvokerNew("set_device")
	})
	return err
}

// SetDevice is a representation of the C type gdk_drag_context_set_device.
func (recv *DragContext) SetDevice(device *Device) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(device.Native)

	err := dragContextSetDeviceFunction_Set()
	if err == nil {
		dragContextSetDeviceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragContextSetHotspotFunction *gi.Function
var dragContextSetHotspotFunction_Once sync.Once

func dragContextSetHotspotFunction_Set() error {
	var err error
	dragContextSetHotspotFunction_Once.Do(func() {
		err = dragContextStruct_Set()
		if err != nil {
			return
		}
		dragContextSetHotspotFunction, err = dragContextStruct.InvokerNew("set_hotspot")
	})
	return err
}

// SetHotspot is a representation of the C type gdk_drag_context_set_hotspot.
func (recv *DragContext) SetHotspot(hotX int32, hotY int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(hotX)
	inArgs[2].SetInt32(hotY)

	err := dragContextSetHotspotFunction_Set()
	if err == nil {
		dragContextSetHotspotFunction.Invoke(inArgs[:], nil)
	}

	return
}

// DragContextStruct creates an uninitialised DragContext.
func DragContextStruct() *DragContext {
	err := dragContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DragContext{}
	structGo.Native = dragContextStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDragContext)
	return structGo
}
func finalizeDragContext(obj *DragContext) {
	dragContextStruct.Free(obj.Native)
}

var drawingContextStruct *gi.Struct
var drawingContextStruct_Once sync.Once

func drawingContextStruct_Set() error {
	var err error
	drawingContextStruct_Once.Do(func() {
		drawingContextStruct, err = gi.StructNew("Gdk", "DrawingContext")
	})
	return err
}

type DrawingContext struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gdk_drawing_context_get_cairo_context' : return type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_drawing_context_get_clip' : return type 'cairo.Region' not supported

var drawingContextGetWindowFunction *gi.Function
var drawingContextGetWindowFunction_Once sync.Once

func drawingContextGetWindowFunction_Set() error {
	var err error
	drawingContextGetWindowFunction_Once.Do(func() {
		err = drawingContextStruct_Set()
		if err != nil {
			return
		}
		drawingContextGetWindowFunction, err = drawingContextStruct.InvokerNew("get_window")
	})
	return err
}

// GetWindow is a representation of the C type gdk_drawing_context_get_window.
func (recv *DrawingContext) GetWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := drawingContextGetWindowFunction_Set()
	if err == nil {
		ret = drawingContextGetWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var drawingContextIsValidFunction *gi.Function
var drawingContextIsValidFunction_Once sync.Once

func drawingContextIsValidFunction_Set() error {
	var err error
	drawingContextIsValidFunction_Once.Do(func() {
		err = drawingContextStruct_Set()
		if err != nil {
			return
		}
		drawingContextIsValidFunction, err = drawingContextStruct.InvokerNew("is_valid")
	})
	return err
}

// IsValid is a representation of the C type gdk_drawing_context_is_valid.
func (recv *DrawingContext) IsValid() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := drawingContextIsValidFunction_Set()
	if err == nil {
		ret = drawingContextIsValidFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// DrawingContextStruct creates an uninitialised DrawingContext.
func DrawingContextStruct() *DrawingContext {
	err := drawingContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &DrawingContext{}
	structGo.Native = drawingContextStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeDrawingContext)
	return structGo
}
func finalizeDrawingContext(obj *DrawingContext) {
	drawingContextStruct.Free(obj.Native)
}

var frameClockStruct *gi.Struct
var frameClockStruct_Once sync.Once

func frameClockStruct_Set() error {
	var err error
	frameClockStruct_Once.Do(func() {
		frameClockStruct, err = gi.StructNew("Gdk", "FrameClock")
	})
	return err
}

type FrameClock struct {
	gobject.Object
}

var frameClockBeginUpdatingFunction *gi.Function
var frameClockBeginUpdatingFunction_Once sync.Once

func frameClockBeginUpdatingFunction_Set() error {
	var err error
	frameClockBeginUpdatingFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockBeginUpdatingFunction, err = frameClockStruct.InvokerNew("begin_updating")
	})
	return err
}

// BeginUpdating is a representation of the C type gdk_frame_clock_begin_updating.
func (recv *FrameClock) BeginUpdating() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := frameClockBeginUpdatingFunction_Set()
	if err == nil {
		frameClockBeginUpdatingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var frameClockEndUpdatingFunction *gi.Function
var frameClockEndUpdatingFunction_Once sync.Once

func frameClockEndUpdatingFunction_Set() error {
	var err error
	frameClockEndUpdatingFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockEndUpdatingFunction, err = frameClockStruct.InvokerNew("end_updating")
	})
	return err
}

// EndUpdating is a representation of the C type gdk_frame_clock_end_updating.
func (recv *FrameClock) EndUpdating() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := frameClockEndUpdatingFunction_Set()
	if err == nil {
		frameClockEndUpdatingFunction.Invoke(inArgs[:], nil)
	}

	return
}

var frameClockGetCurrentTimingsFunction *gi.Function
var frameClockGetCurrentTimingsFunction_Once sync.Once

func frameClockGetCurrentTimingsFunction_Set() error {
	var err error
	frameClockGetCurrentTimingsFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetCurrentTimingsFunction, err = frameClockStruct.InvokerNew("get_current_timings")
	})
	return err
}

// GetCurrentTimings is a representation of the C type gdk_frame_clock_get_current_timings.
func (recv *FrameClock) GetCurrentTimings() *FrameTimings {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := frameClockGetCurrentTimingsFunction_Set()
	if err == nil {
		ret = frameClockGetCurrentTimingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FrameTimings{}
	retGo.Native = ret.Pointer()

	return retGo
}

var frameClockGetFrameCounterFunction *gi.Function
var frameClockGetFrameCounterFunction_Once sync.Once

func frameClockGetFrameCounterFunction_Set() error {
	var err error
	frameClockGetFrameCounterFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetFrameCounterFunction, err = frameClockStruct.InvokerNew("get_frame_counter")
	})
	return err
}

// GetFrameCounter is a representation of the C type gdk_frame_clock_get_frame_counter.
func (recv *FrameClock) GetFrameCounter() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := frameClockGetFrameCounterFunction_Set()
	if err == nil {
		ret = frameClockGetFrameCounterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var frameClockGetFrameTimeFunction *gi.Function
var frameClockGetFrameTimeFunction_Once sync.Once

func frameClockGetFrameTimeFunction_Set() error {
	var err error
	frameClockGetFrameTimeFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetFrameTimeFunction, err = frameClockStruct.InvokerNew("get_frame_time")
	})
	return err
}

// GetFrameTime is a representation of the C type gdk_frame_clock_get_frame_time.
func (recv *FrameClock) GetFrameTime() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := frameClockGetFrameTimeFunction_Set()
	if err == nil {
		ret = frameClockGetFrameTimeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var frameClockGetHistoryStartFunction *gi.Function
var frameClockGetHistoryStartFunction_Once sync.Once

func frameClockGetHistoryStartFunction_Set() error {
	var err error
	frameClockGetHistoryStartFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetHistoryStartFunction, err = frameClockStruct.InvokerNew("get_history_start")
	})
	return err
}

// GetHistoryStart is a representation of the C type gdk_frame_clock_get_history_start.
func (recv *FrameClock) GetHistoryStart() int64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := frameClockGetHistoryStartFunction_Set()
	if err == nil {
		ret = frameClockGetHistoryStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int64()

	return retGo
}

var frameClockGetRefreshInfoFunction *gi.Function
var frameClockGetRefreshInfoFunction_Once sync.Once

func frameClockGetRefreshInfoFunction_Set() error {
	var err error
	frameClockGetRefreshInfoFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetRefreshInfoFunction, err = frameClockStruct.InvokerNew("get_refresh_info")
	})
	return err
}

// GetRefreshInfo is a representation of the C type gdk_frame_clock_get_refresh_info.
func (recv *FrameClock) GetRefreshInfo(baseTime int64) (int64, int64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt64(baseTime)

	var outArgs [2]gi.Argument

	err := frameClockGetRefreshInfoFunction_Set()
	if err == nil {
		frameClockGetRefreshInfoFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int64()
	out1 := outArgs[1].Int64()

	return out0, out1
}

var frameClockGetTimingsFunction *gi.Function
var frameClockGetTimingsFunction_Once sync.Once

func frameClockGetTimingsFunction_Set() error {
	var err error
	frameClockGetTimingsFunction_Once.Do(func() {
		err = frameClockStruct_Set()
		if err != nil {
			return
		}
		frameClockGetTimingsFunction, err = frameClockStruct.InvokerNew("get_timings")
	})
	return err
}

// GetTimings is a representation of the C type gdk_frame_clock_get_timings.
func (recv *FrameClock) GetTimings(frameCounter int64) *FrameTimings {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt64(frameCounter)

	var ret gi.Argument

	err := frameClockGetTimingsFunction_Set()
	if err == nil {
		ret = frameClockGetTimingsFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FrameTimings{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_frame_clock_request_phase' : parameter 'phase' of type 'FrameClockPhase' not supported

// FrameClockStruct creates an uninitialised FrameClock.
func FrameClockStruct() *FrameClock {
	err := frameClockStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &FrameClock{}
	structGo.Native = frameClockStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeFrameClock)
	return structGo
}
func finalizeFrameClock(obj *FrameClock) {
	frameClockStruct.Free(obj.Native)
}

var gLContextStruct *gi.Struct
var gLContextStruct_Once sync.Once

func gLContextStruct_Set() error {
	var err error
	gLContextStruct_Once.Do(func() {
		gLContextStruct, err = gi.StructNew("Gdk", "GLContext")
	})
	return err
}

type GLContext struct {
	gobject.Object
}

var gLContextGetDebugEnabledFunction *gi.Function
var gLContextGetDebugEnabledFunction_Once sync.Once

func gLContextGetDebugEnabledFunction_Set() error {
	var err error
	gLContextGetDebugEnabledFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetDebugEnabledFunction, err = gLContextStruct.InvokerNew("get_debug_enabled")
	})
	return err
}

// GetDebugEnabled is a representation of the C type gdk_gl_context_get_debug_enabled.
func (recv *GLContext) GetDebugEnabled() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetDebugEnabledFunction_Set()
	if err == nil {
		ret = gLContextGetDebugEnabledFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gLContextGetDisplayFunction *gi.Function
var gLContextGetDisplayFunction_Once sync.Once

func gLContextGetDisplayFunction_Set() error {
	var err error
	gLContextGetDisplayFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetDisplayFunction, err = gLContextStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_gl_context_get_display.
func (recv *GLContext) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetDisplayFunction_Set()
	if err == nil {
		ret = gLContextGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

var gLContextGetForwardCompatibleFunction *gi.Function
var gLContextGetForwardCompatibleFunction_Once sync.Once

func gLContextGetForwardCompatibleFunction_Set() error {
	var err error
	gLContextGetForwardCompatibleFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetForwardCompatibleFunction, err = gLContextStruct.InvokerNew("get_forward_compatible")
	})
	return err
}

// GetForwardCompatible is a representation of the C type gdk_gl_context_get_forward_compatible.
func (recv *GLContext) GetForwardCompatible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetForwardCompatibleFunction_Set()
	if err == nil {
		ret = gLContextGetForwardCompatibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gLContextGetRequiredVersionFunction *gi.Function
var gLContextGetRequiredVersionFunction_Once sync.Once

func gLContextGetRequiredVersionFunction_Set() error {
	var err error
	gLContextGetRequiredVersionFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetRequiredVersionFunction, err = gLContextStruct.InvokerNew("get_required_version")
	})
	return err
}

// GetRequiredVersion is a representation of the C type gdk_gl_context_get_required_version.
func (recv *GLContext) GetRequiredVersion() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := gLContextGetRequiredVersionFunction_Set()
	if err == nil {
		gLContextGetRequiredVersionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var gLContextGetSharedContextFunction *gi.Function
var gLContextGetSharedContextFunction_Once sync.Once

func gLContextGetSharedContextFunction_Set() error {
	var err error
	gLContextGetSharedContextFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetSharedContextFunction, err = gLContextStruct.InvokerNew("get_shared_context")
	})
	return err
}

// GetSharedContext is a representation of the C type gdk_gl_context_get_shared_context.
func (recv *GLContext) GetSharedContext() *GLContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetSharedContextFunction_Set()
	if err == nil {
		ret = gLContextGetSharedContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GLContext{}
	retGo.Native = ret.Pointer()

	return retGo
}

var gLContextGetUseEsFunction *gi.Function
var gLContextGetUseEsFunction_Once sync.Once

func gLContextGetUseEsFunction_Set() error {
	var err error
	gLContextGetUseEsFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetUseEsFunction, err = gLContextStruct.InvokerNew("get_use_es")
	})
	return err
}

// GetUseEs is a representation of the C type gdk_gl_context_get_use_es.
func (recv *GLContext) GetUseEs() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetUseEsFunction_Set()
	if err == nil {
		ret = gLContextGetUseEsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gLContextGetVersionFunction *gi.Function
var gLContextGetVersionFunction_Once sync.Once

func gLContextGetVersionFunction_Set() error {
	var err error
	gLContextGetVersionFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetVersionFunction, err = gLContextStruct.InvokerNew("get_version")
	})
	return err
}

// GetVersion is a representation of the C type gdk_gl_context_get_version.
func (recv *GLContext) GetVersion() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := gLContextGetVersionFunction_Set()
	if err == nil {
		gLContextGetVersionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var gLContextGetWindowFunction *gi.Function
var gLContextGetWindowFunction_Once sync.Once

func gLContextGetWindowFunction_Set() error {
	var err error
	gLContextGetWindowFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextGetWindowFunction, err = gLContextStruct.InvokerNew("get_window")
	})
	return err
}

// GetWindow is a representation of the C type gdk_gl_context_get_window.
func (recv *GLContext) GetWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextGetWindowFunction_Set()
	if err == nil {
		ret = gLContextGetWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var gLContextIsLegacyFunction *gi.Function
var gLContextIsLegacyFunction_Once sync.Once

func gLContextIsLegacyFunction_Set() error {
	var err error
	gLContextIsLegacyFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextIsLegacyFunction, err = gLContextStruct.InvokerNew("is_legacy")
	})
	return err
}

// IsLegacy is a representation of the C type gdk_gl_context_is_legacy.
func (recv *GLContext) IsLegacy() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextIsLegacyFunction_Set()
	if err == nil {
		ret = gLContextIsLegacyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gLContextMakeCurrentFunction *gi.Function
var gLContextMakeCurrentFunction_Once sync.Once

func gLContextMakeCurrentFunction_Set() error {
	var err error
	gLContextMakeCurrentFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextMakeCurrentFunction, err = gLContextStruct.InvokerNew("make_current")
	})
	return err
}

// MakeCurrent is a representation of the C type gdk_gl_context_make_current.
func (recv *GLContext) MakeCurrent() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := gLContextMakeCurrentFunction_Set()
	if err == nil {
		gLContextMakeCurrentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gLContextRealizeFunction *gi.Function
var gLContextRealizeFunction_Once sync.Once

func gLContextRealizeFunction_Set() error {
	var err error
	gLContextRealizeFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextRealizeFunction, err = gLContextStruct.InvokerNew("realize")
	})
	return err
}

// Realize is a representation of the C type gdk_gl_context_realize.
func (recv *GLContext) Realize() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := gLContextRealizeFunction_Set()
	if err == nil {
		ret = gLContextRealizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var gLContextSetDebugEnabledFunction *gi.Function
var gLContextSetDebugEnabledFunction_Once sync.Once

func gLContextSetDebugEnabledFunction_Set() error {
	var err error
	gLContextSetDebugEnabledFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextSetDebugEnabledFunction, err = gLContextStruct.InvokerNew("set_debug_enabled")
	})
	return err
}

// SetDebugEnabled is a representation of the C type gdk_gl_context_set_debug_enabled.
func (recv *GLContext) SetDebugEnabled(enabled bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(enabled)

	err := gLContextSetDebugEnabledFunction_Set()
	if err == nil {
		gLContextSetDebugEnabledFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gLContextSetForwardCompatibleFunction *gi.Function
var gLContextSetForwardCompatibleFunction_Once sync.Once

func gLContextSetForwardCompatibleFunction_Set() error {
	var err error
	gLContextSetForwardCompatibleFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextSetForwardCompatibleFunction, err = gLContextStruct.InvokerNew("set_forward_compatible")
	})
	return err
}

// SetForwardCompatible is a representation of the C type gdk_gl_context_set_forward_compatible.
func (recv *GLContext) SetForwardCompatible(compatible bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(compatible)

	err := gLContextSetForwardCompatibleFunction_Set()
	if err == nil {
		gLContextSetForwardCompatibleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gLContextSetRequiredVersionFunction *gi.Function
var gLContextSetRequiredVersionFunction_Once sync.Once

func gLContextSetRequiredVersionFunction_Set() error {
	var err error
	gLContextSetRequiredVersionFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextSetRequiredVersionFunction, err = gLContextStruct.InvokerNew("set_required_version")
	})
	return err
}

// SetRequiredVersion is a representation of the C type gdk_gl_context_set_required_version.
func (recv *GLContext) SetRequiredVersion(major int32, minor int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(major)
	inArgs[2].SetInt32(minor)

	err := gLContextSetRequiredVersionFunction_Set()
	if err == nil {
		gLContextSetRequiredVersionFunction.Invoke(inArgs[:], nil)
	}

	return
}

var gLContextSetUseEsFunction *gi.Function
var gLContextSetUseEsFunction_Once sync.Once

func gLContextSetUseEsFunction_Set() error {
	var err error
	gLContextSetUseEsFunction_Once.Do(func() {
		err = gLContextStruct_Set()
		if err != nil {
			return
		}
		gLContextSetUseEsFunction, err = gLContextStruct.InvokerNew("set_use_es")
	})
	return err
}

// SetUseEs is a representation of the C type gdk_gl_context_set_use_es.
func (recv *GLContext) SetUseEs(useEs int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(useEs)

	err := gLContextSetUseEsFunction_Set()
	if err == nil {
		gLContextSetUseEsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// GLContextStruct creates an uninitialised GLContext.
func GLContextStruct() *GLContext {
	err := gLContextStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &GLContext{}
	structGo.Native = gLContextStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeGLContext)
	return structGo
}
func finalizeGLContext(obj *GLContext) {
	gLContextStruct.Free(obj.Native)
}

var keymapStruct *gi.Struct
var keymapStruct_Once sync.Once

func keymapStruct_Set() error {
	var err error
	keymapStruct_Once.Do(func() {
		keymapStruct, err = gi.StructNew("Gdk", "Keymap")
	})
	return err
}

type Keymap struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gdk_keymap_add_virtual_modifiers' : parameter 'state' of type 'ModifierType' not supported

var keymapGetCapsLockStateFunction *gi.Function
var keymapGetCapsLockStateFunction_Once sync.Once

func keymapGetCapsLockStateFunction_Set() error {
	var err error
	keymapGetCapsLockStateFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapGetCapsLockStateFunction, err = keymapStruct.InvokerNew("get_caps_lock_state")
	})
	return err
}

// GetCapsLockState is a representation of the C type gdk_keymap_get_caps_lock_state.
func (recv *Keymap) GetCapsLockState() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := keymapGetCapsLockStateFunction_Set()
	if err == nil {
		ret = keymapGetCapsLockStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_keymap_get_direction' : return type 'Pango.Direction' not supported

// UNSUPPORTED : C value 'gdk_keymap_get_entries_for_keycode' : parameter 'keys' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_keymap_get_entries_for_keyval' : parameter 'keys' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_keymap_get_modifier_mask' : parameter 'intent' of type 'ModifierIntent' not supported

var keymapGetModifierStateFunction *gi.Function
var keymapGetModifierStateFunction_Once sync.Once

func keymapGetModifierStateFunction_Set() error {
	var err error
	keymapGetModifierStateFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapGetModifierStateFunction, err = keymapStruct.InvokerNew("get_modifier_state")
	})
	return err
}

// GetModifierState is a representation of the C type gdk_keymap_get_modifier_state.
func (recv *Keymap) GetModifierState() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := keymapGetModifierStateFunction_Set()
	if err == nil {
		ret = keymapGetModifierStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keymapGetNumLockStateFunction *gi.Function
var keymapGetNumLockStateFunction_Once sync.Once

func keymapGetNumLockStateFunction_Set() error {
	var err error
	keymapGetNumLockStateFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapGetNumLockStateFunction, err = keymapStruct.InvokerNew("get_num_lock_state")
	})
	return err
}

// GetNumLockState is a representation of the C type gdk_keymap_get_num_lock_state.
func (recv *Keymap) GetNumLockState() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := keymapGetNumLockStateFunction_Set()
	if err == nil {
		ret = keymapGetNumLockStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keymapGetScrollLockStateFunction *gi.Function
var keymapGetScrollLockStateFunction_Once sync.Once

func keymapGetScrollLockStateFunction_Set() error {
	var err error
	keymapGetScrollLockStateFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapGetScrollLockStateFunction, err = keymapStruct.InvokerNew("get_scroll_lock_state")
	})
	return err
}

// GetScrollLockState is a representation of the C type gdk_keymap_get_scroll_lock_state.
func (recv *Keymap) GetScrollLockState() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := keymapGetScrollLockStateFunction_Set()
	if err == nil {
		ret = keymapGetScrollLockStateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keymapHaveBidiLayoutsFunction *gi.Function
var keymapHaveBidiLayoutsFunction_Once sync.Once

func keymapHaveBidiLayoutsFunction_Set() error {
	var err error
	keymapHaveBidiLayoutsFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapHaveBidiLayoutsFunction, err = keymapStruct.InvokerNew("have_bidi_layouts")
	})
	return err
}

// HaveBidiLayouts is a representation of the C type gdk_keymap_have_bidi_layouts.
func (recv *Keymap) HaveBidiLayouts() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := keymapHaveBidiLayoutsFunction_Set()
	if err == nil {
		ret = keymapHaveBidiLayoutsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keymapLookupKeyFunction *gi.Function
var keymapLookupKeyFunction_Once sync.Once

func keymapLookupKeyFunction_Set() error {
	var err error
	keymapLookupKeyFunction_Once.Do(func() {
		err = keymapStruct_Set()
		if err != nil {
			return
		}
		keymapLookupKeyFunction, err = keymapStruct.InvokerNew("lookup_key")
	})
	return err
}

// LookupKey is a representation of the C type gdk_keymap_lookup_key.
func (recv *Keymap) LookupKey(key *KeymapKey) uint32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(key.Native)

	var ret gi.Argument

	err := keymapLookupKeyFunction_Set()
	if err == nil {
		ret = keymapLookupKeyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_keymap_map_virtual_modifiers' : parameter 'state' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_keymap_translate_keyboard_state' : parameter 'state' of type 'ModifierType' not supported

// KeymapStruct creates an uninitialised Keymap.
func KeymapStruct() *Keymap {
	err := keymapStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Keymap{}
	structGo.Native = keymapStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeKeymap)
	return structGo
}
func finalizeKeymap(obj *Keymap) {
	keymapStruct.Free(obj.Native)
}

var monitorStruct *gi.Struct
var monitorStruct_Once sync.Once

func monitorStruct_Set() error {
	var err error
	monitorStruct_Once.Do(func() {
		monitorStruct, err = gi.StructNew("Gdk", "Monitor")
	})
	return err
}

type Monitor struct {
	gobject.Object
}

var monitorGetDisplayFunction *gi.Function
var monitorGetDisplayFunction_Once sync.Once

func monitorGetDisplayFunction_Set() error {
	var err error
	monitorGetDisplayFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetDisplayFunction, err = monitorStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_monitor_get_display.
func (recv *Monitor) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetDisplayFunction_Set()
	if err == nil {
		ret = monitorGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

var monitorGetGeometryFunction *gi.Function
var monitorGetGeometryFunction_Once sync.Once

func monitorGetGeometryFunction_Set() error {
	var err error
	monitorGetGeometryFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetGeometryFunction, err = monitorStruct.InvokerNew("get_geometry")
	})
	return err
}

// GetGeometry is a representation of the C type gdk_monitor_get_geometry.
func (recv *Monitor) GetGeometry() *Rectangle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [1]gi.Argument

	err := monitorGetGeometryFunction_Set()
	if err == nil {
		monitorGetGeometryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var monitorGetHeightMmFunction *gi.Function
var monitorGetHeightMmFunction_Once sync.Once

func monitorGetHeightMmFunction_Set() error {
	var err error
	monitorGetHeightMmFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetHeightMmFunction, err = monitorStruct.InvokerNew("get_height_mm")
	})
	return err
}

// GetHeightMm is a representation of the C type gdk_monitor_get_height_mm.
func (recv *Monitor) GetHeightMm() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetHeightMmFunction_Set()
	if err == nil {
		ret = monitorGetHeightMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var monitorGetManufacturerFunction *gi.Function
var monitorGetManufacturerFunction_Once sync.Once

func monitorGetManufacturerFunction_Set() error {
	var err error
	monitorGetManufacturerFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetManufacturerFunction, err = monitorStruct.InvokerNew("get_manufacturer")
	})
	return err
}

// GetManufacturer is a representation of the C type gdk_monitor_get_manufacturer.
func (recv *Monitor) GetManufacturer() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetManufacturerFunction_Set()
	if err == nil {
		ret = monitorGetManufacturerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var monitorGetModelFunction *gi.Function
var monitorGetModelFunction_Once sync.Once

func monitorGetModelFunction_Set() error {
	var err error
	monitorGetModelFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetModelFunction, err = monitorStruct.InvokerNew("get_model")
	})
	return err
}

// GetModel is a representation of the C type gdk_monitor_get_model.
func (recv *Monitor) GetModel() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetModelFunction_Set()
	if err == nil {
		ret = monitorGetModelFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var monitorGetRefreshRateFunction *gi.Function
var monitorGetRefreshRateFunction_Once sync.Once

func monitorGetRefreshRateFunction_Set() error {
	var err error
	monitorGetRefreshRateFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetRefreshRateFunction, err = monitorStruct.InvokerNew("get_refresh_rate")
	})
	return err
}

// GetRefreshRate is a representation of the C type gdk_monitor_get_refresh_rate.
func (recv *Monitor) GetRefreshRate() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetRefreshRateFunction_Set()
	if err == nil {
		ret = monitorGetRefreshRateFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var monitorGetScaleFactorFunction *gi.Function
var monitorGetScaleFactorFunction_Once sync.Once

func monitorGetScaleFactorFunction_Set() error {
	var err error
	monitorGetScaleFactorFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetScaleFactorFunction, err = monitorStruct.InvokerNew("get_scale_factor")
	})
	return err
}

// GetScaleFactor is a representation of the C type gdk_monitor_get_scale_factor.
func (recv *Monitor) GetScaleFactor() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetScaleFactorFunction_Set()
	if err == nil {
		ret = monitorGetScaleFactorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_monitor_get_subpixel_layout' : return type 'SubpixelLayout' not supported

var monitorGetWidthMmFunction *gi.Function
var monitorGetWidthMmFunction_Once sync.Once

func monitorGetWidthMmFunction_Set() error {
	var err error
	monitorGetWidthMmFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetWidthMmFunction, err = monitorStruct.InvokerNew("get_width_mm")
	})
	return err
}

// GetWidthMm is a representation of the C type gdk_monitor_get_width_mm.
func (recv *Monitor) GetWidthMm() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorGetWidthMmFunction_Set()
	if err == nil {
		ret = monitorGetWidthMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var monitorGetWorkareaFunction *gi.Function
var monitorGetWorkareaFunction_Once sync.Once

func monitorGetWorkareaFunction_Set() error {
	var err error
	monitorGetWorkareaFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorGetWorkareaFunction, err = monitorStruct.InvokerNew("get_workarea")
	})
	return err
}

// GetWorkarea is a representation of the C type gdk_monitor_get_workarea.
func (recv *Monitor) GetWorkarea() *Rectangle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [1]gi.Argument

	err := monitorGetWorkareaFunction_Set()
	if err == nil {
		monitorGetWorkareaFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var monitorIsPrimaryFunction *gi.Function
var monitorIsPrimaryFunction_Once sync.Once

func monitorIsPrimaryFunction_Set() error {
	var err error
	monitorIsPrimaryFunction_Once.Do(func() {
		err = monitorStruct_Set()
		if err != nil {
			return
		}
		monitorIsPrimaryFunction, err = monitorStruct.InvokerNew("is_primary")
	})
	return err
}

// IsPrimary is a representation of the C type gdk_monitor_is_primary.
func (recv *Monitor) IsPrimary() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := monitorIsPrimaryFunction_Set()
	if err == nil {
		ret = monitorIsPrimaryFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// MonitorStruct creates an uninitialised Monitor.
func MonitorStruct() *Monitor {
	err := monitorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Monitor{}
	structGo.Native = monitorStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeMonitor)
	return structGo
}
func finalizeMonitor(obj *Monitor) {
	monitorStruct.Free(obj.Native)
}

var screenStruct *gi.Struct
var screenStruct_Once sync.Once

func screenStruct_Set() error {
	var err error
	screenStruct_Once.Do(func() {
		screenStruct, err = gi.StructNew("Gdk", "Screen")
	})
	return err
}

type Screen struct {
	gobject.Object
}

var screenGetActiveWindowFunction *gi.Function
var screenGetActiveWindowFunction_Once sync.Once

func screenGetActiveWindowFunction_Set() error {
	var err error
	screenGetActiveWindowFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetActiveWindowFunction, err = screenStruct.InvokerNew("get_active_window")
	})
	return err
}

// GetActiveWindow is a representation of the C type gdk_screen_get_active_window.
func (recv *Screen) GetActiveWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetActiveWindowFunction_Set()
	if err == nil {
		ret = screenGetActiveWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var screenGetDisplayFunction *gi.Function
var screenGetDisplayFunction_Once sync.Once

func screenGetDisplayFunction_Set() error {
	var err error
	screenGetDisplayFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetDisplayFunction, err = screenStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_screen_get_display.
func (recv *Screen) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetDisplayFunction_Set()
	if err == nil {
		ret = screenGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_get_font_options' : return type 'cairo.FontOptions' not supported

var screenGetHeightFunction *gi.Function
var screenGetHeightFunction_Once sync.Once

func screenGetHeightFunction_Set() error {
	var err error
	screenGetHeightFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetHeightFunction, err = screenStruct.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type gdk_screen_get_height.
func (recv *Screen) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetHeightFunction_Set()
	if err == nil {
		ret = screenGetHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetHeightMmFunction *gi.Function
var screenGetHeightMmFunction_Once sync.Once

func screenGetHeightMmFunction_Set() error {
	var err error
	screenGetHeightMmFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetHeightMmFunction, err = screenStruct.InvokerNew("get_height_mm")
	})
	return err
}

// GetHeightMm is a representation of the C type gdk_screen_get_height_mm.
func (recv *Screen) GetHeightMm() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetHeightMmFunction_Set()
	if err == nil {
		ret = screenGetHeightMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorAtPointFunction *gi.Function
var screenGetMonitorAtPointFunction_Once sync.Once

func screenGetMonitorAtPointFunction_Set() error {
	var err error
	screenGetMonitorAtPointFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorAtPointFunction, err = screenStruct.InvokerNew("get_monitor_at_point")
	})
	return err
}

// GetMonitorAtPoint is a representation of the C type gdk_screen_get_monitor_at_point.
func (recv *Screen) GetMonitorAtPoint(x int32, y int32) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var ret gi.Argument

	err := screenGetMonitorAtPointFunction_Set()
	if err == nil {
		ret = screenGetMonitorAtPointFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorAtWindowFunction *gi.Function
var screenGetMonitorAtWindowFunction_Once sync.Once

func screenGetMonitorAtWindowFunction_Set() error {
	var err error
	screenGetMonitorAtWindowFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorAtWindowFunction, err = screenStruct.InvokerNew("get_monitor_at_window")
	})
	return err
}

// GetMonitorAtWindow is a representation of the C type gdk_screen_get_monitor_at_window.
func (recv *Screen) GetMonitorAtWindow(window *Window) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(window.Native)

	var ret gi.Argument

	err := screenGetMonitorAtWindowFunction_Set()
	if err == nil {
		ret = screenGetMonitorAtWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorGeometryFunction *gi.Function
var screenGetMonitorGeometryFunction_Once sync.Once

func screenGetMonitorGeometryFunction_Set() error {
	var err error
	screenGetMonitorGeometryFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorGeometryFunction, err = screenStruct.InvokerNew("get_monitor_geometry")
	})
	return err
}

// GetMonitorGeometry is a representation of the C type gdk_screen_get_monitor_geometry.
func (recv *Screen) GetMonitorGeometry(monitorNum int32) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var outArgs [1]gi.Argument

	err := screenGetMonitorGeometryFunction_Set()
	if err == nil {
		screenGetMonitorGeometryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var screenGetMonitorHeightMmFunction *gi.Function
var screenGetMonitorHeightMmFunction_Once sync.Once

func screenGetMonitorHeightMmFunction_Set() error {
	var err error
	screenGetMonitorHeightMmFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorHeightMmFunction, err = screenStruct.InvokerNew("get_monitor_height_mm")
	})
	return err
}

// GetMonitorHeightMm is a representation of the C type gdk_screen_get_monitor_height_mm.
func (recv *Screen) GetMonitorHeightMm(monitorNum int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var ret gi.Argument

	err := screenGetMonitorHeightMmFunction_Set()
	if err == nil {
		ret = screenGetMonitorHeightMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorPlugNameFunction *gi.Function
var screenGetMonitorPlugNameFunction_Once sync.Once

func screenGetMonitorPlugNameFunction_Set() error {
	var err error
	screenGetMonitorPlugNameFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorPlugNameFunction, err = screenStruct.InvokerNew("get_monitor_plug_name")
	})
	return err
}

// GetMonitorPlugName is a representation of the C type gdk_screen_get_monitor_plug_name.
func (recv *Screen) GetMonitorPlugName(monitorNum int32) string {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var ret gi.Argument

	err := screenGetMonitorPlugNameFunction_Set()
	if err == nil {
		ret = screenGetMonitorPlugNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var screenGetMonitorScaleFactorFunction *gi.Function
var screenGetMonitorScaleFactorFunction_Once sync.Once

func screenGetMonitorScaleFactorFunction_Set() error {
	var err error
	screenGetMonitorScaleFactorFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorScaleFactorFunction, err = screenStruct.InvokerNew("get_monitor_scale_factor")
	})
	return err
}

// GetMonitorScaleFactor is a representation of the C type gdk_screen_get_monitor_scale_factor.
func (recv *Screen) GetMonitorScaleFactor(monitorNum int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var ret gi.Argument

	err := screenGetMonitorScaleFactorFunction_Set()
	if err == nil {
		ret = screenGetMonitorScaleFactorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorWidthMmFunction *gi.Function
var screenGetMonitorWidthMmFunction_Once sync.Once

func screenGetMonitorWidthMmFunction_Set() error {
	var err error
	screenGetMonitorWidthMmFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorWidthMmFunction, err = screenStruct.InvokerNew("get_monitor_width_mm")
	})
	return err
}

// GetMonitorWidthMm is a representation of the C type gdk_screen_get_monitor_width_mm.
func (recv *Screen) GetMonitorWidthMm(monitorNum int32) int32 {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var ret gi.Argument

	err := screenGetMonitorWidthMmFunction_Set()
	if err == nil {
		ret = screenGetMonitorWidthMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetMonitorWorkareaFunction *gi.Function
var screenGetMonitorWorkareaFunction_Once sync.Once

func screenGetMonitorWorkareaFunction_Set() error {
	var err error
	screenGetMonitorWorkareaFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetMonitorWorkareaFunction, err = screenStruct.InvokerNew("get_monitor_workarea")
	})
	return err
}

// GetMonitorWorkarea is a representation of the C type gdk_screen_get_monitor_workarea.
func (recv *Screen) GetMonitorWorkarea(monitorNum int32) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitorNum)

	var outArgs [1]gi.Argument

	err := screenGetMonitorWorkareaFunction_Set()
	if err == nil {
		screenGetMonitorWorkareaFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

var screenGetNMonitorsFunction *gi.Function
var screenGetNMonitorsFunction_Once sync.Once

func screenGetNMonitorsFunction_Set() error {
	var err error
	screenGetNMonitorsFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetNMonitorsFunction, err = screenStruct.InvokerNew("get_n_monitors")
	})
	return err
}

// GetNMonitors is a representation of the C type gdk_screen_get_n_monitors.
func (recv *Screen) GetNMonitors() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetNMonitorsFunction_Set()
	if err == nil {
		ret = screenGetNMonitorsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetNumberFunction *gi.Function
var screenGetNumberFunction_Once sync.Once

func screenGetNumberFunction_Set() error {
	var err error
	screenGetNumberFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetNumberFunction, err = screenStruct.InvokerNew("get_number")
	})
	return err
}

// GetNumber is a representation of the C type gdk_screen_get_number.
func (recv *Screen) GetNumber() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetNumberFunction_Set()
	if err == nil {
		ret = screenGetNumberFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetPrimaryMonitorFunction *gi.Function
var screenGetPrimaryMonitorFunction_Once sync.Once

func screenGetPrimaryMonitorFunction_Set() error {
	var err error
	screenGetPrimaryMonitorFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetPrimaryMonitorFunction, err = screenStruct.InvokerNew("get_primary_monitor")
	})
	return err
}

// GetPrimaryMonitor is a representation of the C type gdk_screen_get_primary_monitor.
func (recv *Screen) GetPrimaryMonitor() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetPrimaryMonitorFunction_Set()
	if err == nil {
		ret = screenGetPrimaryMonitorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetResolutionFunction *gi.Function
var screenGetResolutionFunction_Once sync.Once

func screenGetResolutionFunction_Set() error {
	var err error
	screenGetResolutionFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetResolutionFunction, err = screenStruct.InvokerNew("get_resolution")
	})
	return err
}

// GetResolution is a representation of the C type gdk_screen_get_resolution.
func (recv *Screen) GetResolution() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetResolutionFunction_Set()
	if err == nil {
		ret = screenGetResolutionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var screenGetRgbaVisualFunction *gi.Function
var screenGetRgbaVisualFunction_Once sync.Once

func screenGetRgbaVisualFunction_Set() error {
	var err error
	screenGetRgbaVisualFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetRgbaVisualFunction, err = screenStruct.InvokerNew("get_rgba_visual")
	})
	return err
}

// GetRgbaVisual is a representation of the C type gdk_screen_get_rgba_visual.
func (recv *Screen) GetRgbaVisual() *Visual {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetRgbaVisualFunction_Set()
	if err == nil {
		ret = screenGetRgbaVisualFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Visual{}
	retGo.Native = ret.Pointer()

	return retGo
}

var screenGetRootWindowFunction *gi.Function
var screenGetRootWindowFunction_Once sync.Once

func screenGetRootWindowFunction_Set() error {
	var err error
	screenGetRootWindowFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetRootWindowFunction, err = screenStruct.InvokerNew("get_root_window")
	})
	return err
}

// GetRootWindow is a representation of the C type gdk_screen_get_root_window.
func (recv *Screen) GetRootWindow() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetRootWindowFunction_Set()
	if err == nil {
		ret = screenGetRootWindowFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_get_setting' : parameter 'value' of type 'GObject.Value' not supported

var screenGetSystemVisualFunction *gi.Function
var screenGetSystemVisualFunction_Once sync.Once

func screenGetSystemVisualFunction_Set() error {
	var err error
	screenGetSystemVisualFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetSystemVisualFunction, err = screenStruct.InvokerNew("get_system_visual")
	})
	return err
}

// GetSystemVisual is a representation of the C type gdk_screen_get_system_visual.
func (recv *Screen) GetSystemVisual() *Visual {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetSystemVisualFunction_Set()
	if err == nil {
		ret = screenGetSystemVisualFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Visual{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_get_toplevel_windows' : return type 'GLib.List' not supported

var screenGetWidthFunction *gi.Function
var screenGetWidthFunction_Once sync.Once

func screenGetWidthFunction_Set() error {
	var err error
	screenGetWidthFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetWidthFunction, err = screenStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type gdk_screen_get_width.
func (recv *Screen) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetWidthFunction_Set()
	if err == nil {
		ret = screenGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var screenGetWidthMmFunction *gi.Function
var screenGetWidthMmFunction_Once sync.Once

func screenGetWidthMmFunction_Set() error {
	var err error
	screenGetWidthMmFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenGetWidthMmFunction, err = screenStruct.InvokerNew("get_width_mm")
	})
	return err
}

// GetWidthMm is a representation of the C type gdk_screen_get_width_mm.
func (recv *Screen) GetWidthMm() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenGetWidthMmFunction_Set()
	if err == nil {
		ret = screenGetWidthMmFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_get_window_stack' : return type 'GLib.List' not supported

var screenIsCompositedFunction *gi.Function
var screenIsCompositedFunction_Once sync.Once

func screenIsCompositedFunction_Set() error {
	var err error
	screenIsCompositedFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenIsCompositedFunction, err = screenStruct.InvokerNew("is_composited")
	})
	return err
}

// IsComposited is a representation of the C type gdk_screen_is_composited.
func (recv *Screen) IsComposited() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenIsCompositedFunction_Set()
	if err == nil {
		ret = screenIsCompositedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_list_visuals' : return type 'GLib.List' not supported

var screenMakeDisplayNameFunction *gi.Function
var screenMakeDisplayNameFunction_Once sync.Once

func screenMakeDisplayNameFunction_Set() error {
	var err error
	screenMakeDisplayNameFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenMakeDisplayNameFunction, err = screenStruct.InvokerNew("make_display_name")
	})
	return err
}

// MakeDisplayName is a representation of the C type gdk_screen_make_display_name.
func (recv *Screen) MakeDisplayName() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := screenMakeDisplayNameFunction_Set()
	if err == nil {
		ret = screenMakeDisplayNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'gdk_screen_set_font_options' : parameter 'options' of type 'cairo.FontOptions' not supported

var screenSetResolutionFunction *gi.Function
var screenSetResolutionFunction_Once sync.Once

func screenSetResolutionFunction_Set() error {
	var err error
	screenSetResolutionFunction_Once.Do(func() {
		err = screenStruct_Set()
		if err != nil {
			return
		}
		screenSetResolutionFunction, err = screenStruct.InvokerNew("set_resolution")
	})
	return err
}

// SetResolution is a representation of the C type gdk_screen_set_resolution.
func (recv *Screen) SetResolution(dpi float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetFloat64(dpi)

	err := screenSetResolutionFunction_Set()
	if err == nil {
		screenSetResolutionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// ScreenStruct creates an uninitialised Screen.
func ScreenStruct() *Screen {
	err := screenStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Screen{}
	structGo.Native = screenStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeScreen)
	return structGo
}
func finalizeScreen(obj *Screen) {
	screenStruct.Free(obj.Native)
}

var seatStruct *gi.Struct
var seatStruct_Once sync.Once

func seatStruct_Set() error {
	var err error
	seatStruct_Once.Do(func() {
		seatStruct, err = gi.StructNew("Gdk", "Seat")
	})
	return err
}

type Seat struct {
	gobject.Object
}

// UNSUPPORTED : C value 'parent_instance' : for field getter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'parent_instance' : for field setter : no Go type for 'GObject.Object'

// UNSUPPORTED : C value 'gdk_seat_get_capabilities' : return type 'SeatCapabilities' not supported

var seatGetDisplayFunction *gi.Function
var seatGetDisplayFunction_Once sync.Once

func seatGetDisplayFunction_Set() error {
	var err error
	seatGetDisplayFunction_Once.Do(func() {
		err = seatStruct_Set()
		if err != nil {
			return
		}
		seatGetDisplayFunction, err = seatStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_seat_get_display.
func (recv *Seat) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := seatGetDisplayFunction_Set()
	if err == nil {
		ret = seatGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

var seatGetKeyboardFunction *gi.Function
var seatGetKeyboardFunction_Once sync.Once

func seatGetKeyboardFunction_Set() error {
	var err error
	seatGetKeyboardFunction_Once.Do(func() {
		err = seatStruct_Set()
		if err != nil {
			return
		}
		seatGetKeyboardFunction, err = seatStruct.InvokerNew("get_keyboard")
	})
	return err
}

// GetKeyboard is a representation of the C type gdk_seat_get_keyboard.
func (recv *Seat) GetKeyboard() *Device {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := seatGetKeyboardFunction_Set()
	if err == nil {
		ret = seatGetKeyboardFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Device{}
	retGo.Native = ret.Pointer()

	return retGo
}

var seatGetPointerFunction *gi.Function
var seatGetPointerFunction_Once sync.Once

func seatGetPointerFunction_Set() error {
	var err error
	seatGetPointerFunction_Once.Do(func() {
		err = seatStruct_Set()
		if err != nil {
			return
		}
		seatGetPointerFunction, err = seatStruct.InvokerNew("get_pointer")
	})
	return err
}

// GetPointer is a representation of the C type gdk_seat_get_pointer.
func (recv *Seat) GetPointer() *Device {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := seatGetPointerFunction_Set()
	if err == nil {
		ret = seatGetPointerFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Device{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_seat_get_slaves' : parameter 'capabilities' of type 'SeatCapabilities' not supported

// UNSUPPORTED : C value 'gdk_seat_grab' : parameter 'capabilities' of type 'SeatCapabilities' not supported

var seatUngrabFunction *gi.Function
var seatUngrabFunction_Once sync.Once

func seatUngrabFunction_Set() error {
	var err error
	seatUngrabFunction_Once.Do(func() {
		err = seatStruct_Set()
		if err != nil {
			return
		}
		seatUngrabFunction, err = seatStruct.InvokerNew("ungrab")
	})
	return err
}

// Ungrab is a representation of the C type gdk_seat_ungrab.
func (recv *Seat) Ungrab() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := seatUngrabFunction_Set()
	if err == nil {
		seatUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

// SeatStruct creates an uninitialised Seat.
func SeatStruct() *Seat {
	err := seatStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Seat{}
	structGo.Native = seatStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeSeat)
	return structGo
}
func finalizeSeat(obj *Seat) {
	seatStruct.Free(obj.Native)
}

var visualStruct *gi.Struct
var visualStruct_Once sync.Once

func visualStruct_Set() error {
	var err error
	visualStruct_Once.Do(func() {
		visualStruct, err = gi.StructNew("Gdk", "Visual")
	})
	return err
}

type Visual struct {
	gobject.Object
}

var visualGetBitsPerRgbFunction *gi.Function
var visualGetBitsPerRgbFunction_Once sync.Once

func visualGetBitsPerRgbFunction_Set() error {
	var err error
	visualGetBitsPerRgbFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetBitsPerRgbFunction, err = visualStruct.InvokerNew("get_bits_per_rgb")
	})
	return err
}

// GetBitsPerRgb is a representation of the C type gdk_visual_get_bits_per_rgb.
func (recv *Visual) GetBitsPerRgb() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := visualGetBitsPerRgbFunction_Set()
	if err == nil {
		ret = visualGetBitsPerRgbFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var visualGetBluePixelDetailsFunction *gi.Function
var visualGetBluePixelDetailsFunction_Once sync.Once

func visualGetBluePixelDetailsFunction_Set() error {
	var err error
	visualGetBluePixelDetailsFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetBluePixelDetailsFunction, err = visualStruct.InvokerNew("get_blue_pixel_details")
	})
	return err
}

// GetBluePixelDetails is a representation of the C type gdk_visual_get_blue_pixel_details.
func (recv *Visual) GetBluePixelDetails() (uint32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [3]gi.Argument

	err := visualGetBluePixelDetailsFunction_Set()
	if err == nil {
		visualGetBluePixelDetailsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

// UNSUPPORTED : C value 'gdk_visual_get_byte_order' : return type 'ByteOrder' not supported

var visualGetColormapSizeFunction *gi.Function
var visualGetColormapSizeFunction_Once sync.Once

func visualGetColormapSizeFunction_Set() error {
	var err error
	visualGetColormapSizeFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetColormapSizeFunction, err = visualStruct.InvokerNew("get_colormap_size")
	})
	return err
}

// GetColormapSize is a representation of the C type gdk_visual_get_colormap_size.
func (recv *Visual) GetColormapSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := visualGetColormapSizeFunction_Set()
	if err == nil {
		ret = visualGetColormapSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var visualGetDepthFunction *gi.Function
var visualGetDepthFunction_Once sync.Once

func visualGetDepthFunction_Set() error {
	var err error
	visualGetDepthFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetDepthFunction, err = visualStruct.InvokerNew("get_depth")
	})
	return err
}

// GetDepth is a representation of the C type gdk_visual_get_depth.
func (recv *Visual) GetDepth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := visualGetDepthFunction_Set()
	if err == nil {
		ret = visualGetDepthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var visualGetGreenPixelDetailsFunction *gi.Function
var visualGetGreenPixelDetailsFunction_Once sync.Once

func visualGetGreenPixelDetailsFunction_Set() error {
	var err error
	visualGetGreenPixelDetailsFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetGreenPixelDetailsFunction, err = visualStruct.InvokerNew("get_green_pixel_details")
	})
	return err
}

// GetGreenPixelDetails is a representation of the C type gdk_visual_get_green_pixel_details.
func (recv *Visual) GetGreenPixelDetails() (uint32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [3]gi.Argument

	err := visualGetGreenPixelDetailsFunction_Set()
	if err == nil {
		visualGetGreenPixelDetailsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

var visualGetRedPixelDetailsFunction *gi.Function
var visualGetRedPixelDetailsFunction_Once sync.Once

func visualGetRedPixelDetailsFunction_Set() error {
	var err error
	visualGetRedPixelDetailsFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetRedPixelDetailsFunction, err = visualStruct.InvokerNew("get_red_pixel_details")
	})
	return err
}

// GetRedPixelDetails is a representation of the C type gdk_visual_get_red_pixel_details.
func (recv *Visual) GetRedPixelDetails() (uint32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [3]gi.Argument

	err := visualGetRedPixelDetailsFunction_Set()
	if err == nil {
		visualGetRedPixelDetailsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()

	return out0, out1, out2
}

var visualGetScreenFunction *gi.Function
var visualGetScreenFunction_Once sync.Once

func visualGetScreenFunction_Set() error {
	var err error
	visualGetScreenFunction_Once.Do(func() {
		err = visualStruct_Set()
		if err != nil {
			return
		}
		visualGetScreenFunction, err = visualStruct.InvokerNew("get_screen")
	})
	return err
}

// GetScreen is a representation of the C type gdk_visual_get_screen.
func (recv *Visual) GetScreen() *Screen {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := visualGetScreenFunction_Set()
	if err == nil {
		ret = visualGetScreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Screen{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_visual_get_visual_type' : return type 'VisualType' not supported

// VisualStruct creates an uninitialised Visual.
func VisualStruct() *Visual {
	err := visualStruct_Set()
	if err != nil {
		return nil
	}

	structGo := &Visual{}
	structGo.Native = visualStruct.Alloc()
	runtime.SetFinalizer(structGo, finalizeVisual)
	return structGo
}
func finalizeVisual(obj *Visual) {
	visualStruct.Free(obj.Native)
}

var windowStruct *gi.Struct
var windowStruct_Once sync.Once

func windowStruct_Set() error {
	var err error
	windowStruct_Once.Do(func() {
		windowStruct, err = gi.StructNew("Gdk", "Window")
	})
	return err
}

type Window struct {
	gobject.Object
}

// UNSUPPORTED : C value 'gdk_window_new' : parameter 'attributes_mask' of type 'WindowAttributesType' not supported

// UNSUPPORTED : C value 'gdk_window_add_filter' : parameter 'function' of type 'FilterFunc' not supported

var windowBeepFunction *gi.Function
var windowBeepFunction_Once sync.Once

func windowBeepFunction_Set() error {
	var err error
	windowBeepFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowBeepFunction, err = windowStruct.InvokerNew("beep")
	})
	return err
}

// Beep is a representation of the C type gdk_window_beep.
func (recv *Window) Beep() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowBeepFunction_Set()
	if err == nil {
		windowBeepFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_begin_draw_frame' : parameter 'region' of type 'cairo.Region' not supported

var windowBeginMoveDragFunction *gi.Function
var windowBeginMoveDragFunction_Once sync.Once

func windowBeginMoveDragFunction_Set() error {
	var err error
	windowBeginMoveDragFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowBeginMoveDragFunction, err = windowStruct.InvokerNew("begin_move_drag")
	})
	return err
}

// BeginMoveDrag is a representation of the C type gdk_window_begin_move_drag.
func (recv *Window) BeginMoveDrag(button int32, rootX int32, rootY int32, timestamp uint32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(button)
	inArgs[2].SetInt32(rootX)
	inArgs[3].SetInt32(rootY)
	inArgs[4].SetUint32(timestamp)

	err := windowBeginMoveDragFunction_Set()
	if err == nil {
		windowBeginMoveDragFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowBeginMoveDragForDeviceFunction *gi.Function
var windowBeginMoveDragForDeviceFunction_Once sync.Once

func windowBeginMoveDragForDeviceFunction_Set() error {
	var err error
	windowBeginMoveDragForDeviceFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowBeginMoveDragForDeviceFunction, err = windowStruct.InvokerNew("begin_move_drag_for_device")
	})
	return err
}

// BeginMoveDragForDevice is a representation of the C type gdk_window_begin_move_drag_for_device.
func (recv *Window) BeginMoveDragForDevice(device *Device, button int32, rootX int32, rootY int32, timestamp uint32) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(device.Native)
	inArgs[2].SetInt32(button)
	inArgs[3].SetInt32(rootX)
	inArgs[4].SetInt32(rootY)
	inArgs[5].SetUint32(timestamp)

	err := windowBeginMoveDragForDeviceFunction_Set()
	if err == nil {
		windowBeginMoveDragForDeviceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowBeginPaintRectFunction *gi.Function
var windowBeginPaintRectFunction_Once sync.Once

func windowBeginPaintRectFunction_Set() error {
	var err error
	windowBeginPaintRectFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowBeginPaintRectFunction, err = windowStruct.InvokerNew("begin_paint_rect")
	})
	return err
}

// BeginPaintRect is a representation of the C type gdk_window_begin_paint_rect.
func (recv *Window) BeginPaintRect(rectangle *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(rectangle.Native)

	err := windowBeginPaintRectFunction_Set()
	if err == nil {
		windowBeginPaintRectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_begin_paint_region' : parameter 'region' of type 'cairo.Region' not supported

// UNSUPPORTED : C value 'gdk_window_begin_resize_drag' : parameter 'edge' of type 'WindowEdge' not supported

// UNSUPPORTED : C value 'gdk_window_begin_resize_drag_for_device' : parameter 'edge' of type 'WindowEdge' not supported

var windowConfigureFinishedFunction *gi.Function
var windowConfigureFinishedFunction_Once sync.Once

func windowConfigureFinishedFunction_Set() error {
	var err error
	windowConfigureFinishedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowConfigureFinishedFunction, err = windowStruct.InvokerNew("configure_finished")
	})
	return err
}

// ConfigureFinished is a representation of the C type gdk_window_configure_finished.
func (recv *Window) ConfigureFinished() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowConfigureFinishedFunction_Set()
	if err == nil {
		windowConfigureFinishedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowCoordsFromParentFunction *gi.Function
var windowCoordsFromParentFunction_Once sync.Once

func windowCoordsFromParentFunction_Set() error {
	var err error
	windowCoordsFromParentFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowCoordsFromParentFunction, err = windowStruct.InvokerNew("coords_from_parent")
	})
	return err
}

// CoordsFromParent is a representation of the C type gdk_window_coords_from_parent.
func (recv *Window) CoordsFromParent(parentX float64, parentY float64) (float64, float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetFloat64(parentX)
	inArgs[2].SetFloat64(parentY)

	var outArgs [2]gi.Argument

	err := windowCoordsFromParentFunction_Set()
	if err == nil {
		windowCoordsFromParentFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1
}

var windowCoordsToParentFunction *gi.Function
var windowCoordsToParentFunction_Once sync.Once

func windowCoordsToParentFunction_Set() error {
	var err error
	windowCoordsToParentFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowCoordsToParentFunction, err = windowStruct.InvokerNew("coords_to_parent")
	})
	return err
}

// CoordsToParent is a representation of the C type gdk_window_coords_to_parent.
func (recv *Window) CoordsToParent(x float64, y float64) (float64, float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetFloat64(x)
	inArgs[2].SetFloat64(y)

	var outArgs [2]gi.Argument

	err := windowCoordsToParentFunction_Set()
	if err == nil {
		windowCoordsToParentFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1
}

var windowCreateGlContextFunction *gi.Function
var windowCreateGlContextFunction_Once sync.Once

func windowCreateGlContextFunction_Set() error {
	var err error
	windowCreateGlContextFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowCreateGlContextFunction, err = windowStruct.InvokerNew("create_gl_context")
	})
	return err
}

// CreateGlContext is a representation of the C type gdk_window_create_gl_context.
func (recv *Window) CreateGlContext() *GLContext {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowCreateGlContextFunction_Set()
	if err == nil {
		ret = windowCreateGlContextFunction.Invoke(inArgs[:], nil)
	}

	retGo := &GLContext{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_create_similar_image_surface' : return type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_window_create_similar_surface' : parameter 'content' of type 'cairo.Content' not supported

var windowDeiconifyFunction *gi.Function
var windowDeiconifyFunction_Once sync.Once

func windowDeiconifyFunction_Set() error {
	var err error
	windowDeiconifyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowDeiconifyFunction, err = windowStruct.InvokerNew("deiconify")
	})
	return err
}

// Deiconify is a representation of the C type gdk_window_deiconify.
func (recv *Window) Deiconify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowDeiconifyFunction_Set()
	if err == nil {
		windowDeiconifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowDestroyFunction *gi.Function
var windowDestroyFunction_Once sync.Once

func windowDestroyFunction_Set() error {
	var err error
	windowDestroyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowDestroyFunction, err = windowStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type gdk_window_destroy.
func (recv *Window) Destroy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowDestroyFunction_Set()
	if err == nil {
		windowDestroyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowDestroyNotifyFunction *gi.Function
var windowDestroyNotifyFunction_Once sync.Once

func windowDestroyNotifyFunction_Set() error {
	var err error
	windowDestroyNotifyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowDestroyNotifyFunction, err = windowStruct.InvokerNew("destroy_notify")
	})
	return err
}

// DestroyNotify is a representation of the C type gdk_window_destroy_notify.
func (recv *Window) DestroyNotify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowDestroyNotifyFunction_Set()
	if err == nil {
		windowDestroyNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowEnableSynchronizedConfigureFunction *gi.Function
var windowEnableSynchronizedConfigureFunction_Once sync.Once

func windowEnableSynchronizedConfigureFunction_Set() error {
	var err error
	windowEnableSynchronizedConfigureFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowEnableSynchronizedConfigureFunction, err = windowStruct.InvokerNew("enable_synchronized_configure")
	})
	return err
}

// EnableSynchronizedConfigure is a representation of the C type gdk_window_enable_synchronized_configure.
func (recv *Window) EnableSynchronizedConfigure() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowEnableSynchronizedConfigureFunction_Set()
	if err == nil {
		windowEnableSynchronizedConfigureFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowEndDrawFrameFunction *gi.Function
var windowEndDrawFrameFunction_Once sync.Once

func windowEndDrawFrameFunction_Set() error {
	var err error
	windowEndDrawFrameFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowEndDrawFrameFunction, err = windowStruct.InvokerNew("end_draw_frame")
	})
	return err
}

// EndDrawFrame is a representation of the C type gdk_window_end_draw_frame.
func (recv *Window) EndDrawFrame(context *DrawingContext) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(context.Native)

	err := windowEndDrawFrameFunction_Set()
	if err == nil {
		windowEndDrawFrameFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowEndPaintFunction *gi.Function
var windowEndPaintFunction_Once sync.Once

func windowEndPaintFunction_Set() error {
	var err error
	windowEndPaintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowEndPaintFunction, err = windowStruct.InvokerNew("end_paint")
	})
	return err
}

// EndPaint is a representation of the C type gdk_window_end_paint.
func (recv *Window) EndPaint() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowEndPaintFunction_Set()
	if err == nil {
		windowEndPaintFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowEnsureNativeFunction *gi.Function
var windowEnsureNativeFunction_Once sync.Once

func windowEnsureNativeFunction_Set() error {
	var err error
	windowEnsureNativeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowEnsureNativeFunction, err = windowStruct.InvokerNew("ensure_native")
	})
	return err
}

// EnsureNative is a representation of the C type gdk_window_ensure_native.
func (recv *Window) EnsureNative() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowEnsureNativeFunction_Set()
	if err == nil {
		ret = windowEnsureNativeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowFlushFunction *gi.Function
var windowFlushFunction_Once sync.Once

func windowFlushFunction_Set() error {
	var err error
	windowFlushFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFlushFunction, err = windowStruct.InvokerNew("flush")
	})
	return err
}

// Flush is a representation of the C type gdk_window_flush.
func (recv *Window) Flush() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowFlushFunction_Set()
	if err == nil {
		windowFlushFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowFocusFunction *gi.Function
var windowFocusFunction_Once sync.Once

func windowFocusFunction_Set() error {
	var err error
	windowFocusFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFocusFunction, err = windowStruct.InvokerNew("focus")
	})
	return err
}

// Focus is a representation of the C type gdk_window_focus.
func (recv *Window) Focus(timestamp uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetUint32(timestamp)

	err := windowFocusFunction_Set()
	if err == nil {
		windowFocusFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowFreezeToplevelUpdatesLibgtkOnlyFunction *gi.Function
var windowFreezeToplevelUpdatesLibgtkOnlyFunction_Once sync.Once

func windowFreezeToplevelUpdatesLibgtkOnlyFunction_Set() error {
	var err error
	windowFreezeToplevelUpdatesLibgtkOnlyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFreezeToplevelUpdatesLibgtkOnlyFunction, err = windowStruct.InvokerNew("freeze_toplevel_updates_libgtk_only")
	})
	return err
}

// FreezeToplevelUpdatesLibgtkOnly is a representation of the C type gdk_window_freeze_toplevel_updates_libgtk_only.
func (recv *Window) FreezeToplevelUpdatesLibgtkOnly() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowFreezeToplevelUpdatesLibgtkOnlyFunction_Set()
	if err == nil {
		windowFreezeToplevelUpdatesLibgtkOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowFreezeUpdatesFunction *gi.Function
var windowFreezeUpdatesFunction_Once sync.Once

func windowFreezeUpdatesFunction_Set() error {
	var err error
	windowFreezeUpdatesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFreezeUpdatesFunction, err = windowStruct.InvokerNew("freeze_updates")
	})
	return err
}

// FreezeUpdates is a representation of the C type gdk_window_freeze_updates.
func (recv *Window) FreezeUpdates() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowFreezeUpdatesFunction_Set()
	if err == nil {
		windowFreezeUpdatesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowFullscreenFunction *gi.Function
var windowFullscreenFunction_Once sync.Once

func windowFullscreenFunction_Set() error {
	var err error
	windowFullscreenFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFullscreenFunction, err = windowStruct.InvokerNew("fullscreen")
	})
	return err
}

// Fullscreen is a representation of the C type gdk_window_fullscreen.
func (recv *Window) Fullscreen() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowFullscreenFunction_Set()
	if err == nil {
		windowFullscreenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowFullscreenOnMonitorFunction *gi.Function
var windowFullscreenOnMonitorFunction_Once sync.Once

func windowFullscreenOnMonitorFunction_Set() error {
	var err error
	windowFullscreenOnMonitorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowFullscreenOnMonitorFunction, err = windowStruct.InvokerNew("fullscreen_on_monitor")
	})
	return err
}

// FullscreenOnMonitor is a representation of the C type gdk_window_fullscreen_on_monitor.
func (recv *Window) FullscreenOnMonitor(monitor int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(monitor)

	err := windowFullscreenOnMonitorFunction_Set()
	if err == nil {
		windowFullscreenOnMonitorFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowGeometryChangedFunction *gi.Function
var windowGeometryChangedFunction_Once sync.Once

func windowGeometryChangedFunction_Set() error {
	var err error
	windowGeometryChangedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGeometryChangedFunction, err = windowStruct.InvokerNew("geometry_changed")
	})
	return err
}

// GeometryChanged is a representation of the C type gdk_window_geometry_changed.
func (recv *Window) GeometryChanged() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowGeometryChangedFunction_Set()
	if err == nil {
		windowGeometryChangedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowGetAcceptFocusFunction *gi.Function
var windowGetAcceptFocusFunction_Once sync.Once

func windowGetAcceptFocusFunction_Set() error {
	var err error
	windowGetAcceptFocusFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetAcceptFocusFunction, err = windowStruct.InvokerNew("get_accept_focus")
	})
	return err
}

// GetAcceptFocus is a representation of the C type gdk_window_get_accept_focus.
func (recv *Window) GetAcceptFocus() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetAcceptFocusFunction_Set()
	if err == nil {
		ret = windowGetAcceptFocusFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_background_pattern' : return type 'cairo.Pattern' not supported

// UNSUPPORTED : C value 'gdk_window_get_children' : return type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_window_get_children_with_user_data' : parameter 'user_data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'gdk_window_get_clip_region' : return type 'cairo.Region' not supported

var windowGetCompositedFunction *gi.Function
var windowGetCompositedFunction_Once sync.Once

func windowGetCompositedFunction_Set() error {
	var err error
	windowGetCompositedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetCompositedFunction, err = windowStruct.InvokerNew("get_composited")
	})
	return err
}

// GetComposited is a representation of the C type gdk_window_get_composited.
func (recv *Window) GetComposited() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetCompositedFunction_Set()
	if err == nil {
		ret = windowGetCompositedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowGetCursorFunction *gi.Function
var windowGetCursorFunction_Once sync.Once

func windowGetCursorFunction_Set() error {
	var err error
	windowGetCursorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetCursorFunction, err = windowStruct.InvokerNew("get_cursor")
	})
	return err
}

// GetCursor is a representation of the C type gdk_window_get_cursor.
func (recv *Window) GetCursor() *Cursor {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetCursorFunction_Set()
	if err == nil {
		ret = windowGetCursorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cursor{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_decorations' : parameter 'decorations' of type 'WMDecoration' not supported

var windowGetDeviceCursorFunction *gi.Function
var windowGetDeviceCursorFunction_Once sync.Once

func windowGetDeviceCursorFunction_Set() error {
	var err error
	windowGetDeviceCursorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetDeviceCursorFunction, err = windowStruct.InvokerNew("get_device_cursor")
	})
	return err
}

// GetDeviceCursor is a representation of the C type gdk_window_get_device_cursor.
func (recv *Window) GetDeviceCursor(device *Device) *Cursor {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(device.Native)

	var ret gi.Argument

	err := windowGetDeviceCursorFunction_Set()
	if err == nil {
		ret = windowGetDeviceCursorFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Cursor{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_device_events' : return type 'EventMask' not supported

// UNSUPPORTED : C value 'gdk_window_get_device_position' : parameter 'mask' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_window_get_device_position_double' : parameter 'mask' of type 'ModifierType' not supported

var windowGetDisplayFunction *gi.Function
var windowGetDisplayFunction_Once sync.Once

func windowGetDisplayFunction_Set() error {
	var err error
	windowGetDisplayFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetDisplayFunction, err = windowStruct.InvokerNew("get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_window_get_display.
func (recv *Window) GetDisplay() *Display {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetDisplayFunction_Set()
	if err == nil {
		ret = windowGetDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Display{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_drag_protocol' : return type 'DragProtocol' not supported

var windowGetEffectiveParentFunction *gi.Function
var windowGetEffectiveParentFunction_Once sync.Once

func windowGetEffectiveParentFunction_Set() error {
	var err error
	windowGetEffectiveParentFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetEffectiveParentFunction, err = windowStruct.InvokerNew("get_effective_parent")
	})
	return err
}

// GetEffectiveParent is a representation of the C type gdk_window_get_effective_parent.
func (recv *Window) GetEffectiveParent() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetEffectiveParentFunction_Set()
	if err == nil {
		ret = windowGetEffectiveParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetEffectiveToplevelFunction *gi.Function
var windowGetEffectiveToplevelFunction_Once sync.Once

func windowGetEffectiveToplevelFunction_Set() error {
	var err error
	windowGetEffectiveToplevelFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetEffectiveToplevelFunction, err = windowStruct.InvokerNew("get_effective_toplevel")
	})
	return err
}

// GetEffectiveToplevel is a representation of the C type gdk_window_get_effective_toplevel.
func (recv *Window) GetEffectiveToplevel() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetEffectiveToplevelFunction_Set()
	if err == nil {
		ret = windowGetEffectiveToplevelFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetEventCompressionFunction *gi.Function
var windowGetEventCompressionFunction_Once sync.Once

func windowGetEventCompressionFunction_Set() error {
	var err error
	windowGetEventCompressionFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetEventCompressionFunction, err = windowStruct.InvokerNew("get_event_compression")
	})
	return err
}

// GetEventCompression is a representation of the C type gdk_window_get_event_compression.
func (recv *Window) GetEventCompression() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetEventCompressionFunction_Set()
	if err == nil {
		ret = windowGetEventCompressionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_events' : return type 'EventMask' not supported

var windowGetFocusOnMapFunction *gi.Function
var windowGetFocusOnMapFunction_Once sync.Once

func windowGetFocusOnMapFunction_Set() error {
	var err error
	windowGetFocusOnMapFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetFocusOnMapFunction, err = windowStruct.InvokerNew("get_focus_on_map")
	})
	return err
}

// GetFocusOnMap is a representation of the C type gdk_window_get_focus_on_map.
func (recv *Window) GetFocusOnMap() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetFocusOnMapFunction_Set()
	if err == nil {
		ret = windowGetFocusOnMapFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowGetFrameClockFunction *gi.Function
var windowGetFrameClockFunction_Once sync.Once

func windowGetFrameClockFunction_Set() error {
	var err error
	windowGetFrameClockFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetFrameClockFunction, err = windowStruct.InvokerNew("get_frame_clock")
	})
	return err
}

// GetFrameClock is a representation of the C type gdk_window_get_frame_clock.
func (recv *Window) GetFrameClock() *FrameClock {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetFrameClockFunction_Set()
	if err == nil {
		ret = windowGetFrameClockFunction.Invoke(inArgs[:], nil)
	}

	retGo := &FrameClock{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetFrameExtentsFunction *gi.Function
var windowGetFrameExtentsFunction_Once sync.Once

func windowGetFrameExtentsFunction_Set() error {
	var err error
	windowGetFrameExtentsFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetFrameExtentsFunction, err = windowStruct.InvokerNew("get_frame_extents")
	})
	return err
}

// GetFrameExtents is a representation of the C type gdk_window_get_frame_extents.
func (recv *Window) GetFrameExtents() *Rectangle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [1]gi.Argument

	err := windowGetFrameExtentsFunction_Set()
	if err == nil {
		windowGetFrameExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := &Rectangle{}
	out0.Native = outArgs[0].Pointer()

	return out0
}

// UNSUPPORTED : C value 'gdk_window_get_fullscreen_mode' : return type 'FullscreenMode' not supported

var windowGetGeometryFunction *gi.Function
var windowGetGeometryFunction_Once sync.Once

func windowGetGeometryFunction_Set() error {
	var err error
	windowGetGeometryFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetGeometryFunction, err = windowStruct.InvokerNew("get_geometry")
	})
	return err
}

// GetGeometry is a representation of the C type gdk_window_get_geometry.
func (recv *Window) GetGeometry() (int32, int32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [4]gi.Argument

	err := windowGetGeometryFunction_Set()
	if err == nil {
		windowGetGeometryFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()
	out2 := outArgs[2].Int32()
	out3 := outArgs[3].Int32()

	return out0, out1, out2, out3
}

var windowGetGroupFunction *gi.Function
var windowGetGroupFunction_Once sync.Once

func windowGetGroupFunction_Set() error {
	var err error
	windowGetGroupFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetGroupFunction, err = windowStruct.InvokerNew("get_group")
	})
	return err
}

// GetGroup is a representation of the C type gdk_window_get_group.
func (recv *Window) GetGroup() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetGroupFunction_Set()
	if err == nil {
		ret = windowGetGroupFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetHeightFunction *gi.Function
var windowGetHeightFunction_Once sync.Once

func windowGetHeightFunction_Set() error {
	var err error
	windowGetHeightFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetHeightFunction, err = windowStruct.InvokerNew("get_height")
	})
	return err
}

// GetHeight is a representation of the C type gdk_window_get_height.
func (recv *Window) GetHeight() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetHeightFunction_Set()
	if err == nil {
		ret = windowGetHeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var windowGetModalHintFunction *gi.Function
var windowGetModalHintFunction_Once sync.Once

func windowGetModalHintFunction_Set() error {
	var err error
	windowGetModalHintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetModalHintFunction, err = windowStruct.InvokerNew("get_modal_hint")
	})
	return err
}

// GetModalHint is a representation of the C type gdk_window_get_modal_hint.
func (recv *Window) GetModalHint() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetModalHintFunction_Set()
	if err == nil {
		ret = windowGetModalHintFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowGetOriginFunction *gi.Function
var windowGetOriginFunction_Once sync.Once

func windowGetOriginFunction_Set() error {
	var err error
	windowGetOriginFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetOriginFunction, err = windowStruct.InvokerNew("get_origin")
	})
	return err
}

// GetOrigin is a representation of the C type gdk_window_get_origin.
func (recv *Window) GetOrigin() (int32, int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := windowGetOriginFunction_Set()
	if err == nil {
		ret = windowGetOriginFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Int32()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

var windowGetParentFunction *gi.Function
var windowGetParentFunction_Once sync.Once

func windowGetParentFunction_Set() error {
	var err error
	windowGetParentFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetParentFunction, err = windowStruct.InvokerNew("get_parent")
	})
	return err
}

// GetParent is a representation of the C type gdk_window_get_parent.
func (recv *Window) GetParent() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetParentFunction_Set()
	if err == nil {
		ret = windowGetParentFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetPassThroughFunction *gi.Function
var windowGetPassThroughFunction_Once sync.Once

func windowGetPassThroughFunction_Set() error {
	var err error
	windowGetPassThroughFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetPassThroughFunction, err = windowStruct.InvokerNew("get_pass_through")
	})
	return err
}

// GetPassThrough is a representation of the C type gdk_window_get_pass_through.
func (recv *Window) GetPassThrough() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetPassThroughFunction_Set()
	if err == nil {
		ret = windowGetPassThroughFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_pointer' : parameter 'mask' of type 'ModifierType' not supported

var windowGetPositionFunction *gi.Function
var windowGetPositionFunction_Once sync.Once

func windowGetPositionFunction_Set() error {
	var err error
	windowGetPositionFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetPositionFunction, err = windowStruct.InvokerNew("get_position")
	})
	return err
}

// GetPosition is a representation of the C type gdk_window_get_position.
func (recv *Window) GetPosition() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := windowGetPositionFunction_Set()
	if err == nil {
		windowGetPositionFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var windowGetRootCoordsFunction *gi.Function
var windowGetRootCoordsFunction_Once sync.Once

func windowGetRootCoordsFunction_Set() error {
	var err error
	windowGetRootCoordsFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetRootCoordsFunction, err = windowStruct.InvokerNew("get_root_coords")
	})
	return err
}

// GetRootCoords is a representation of the C type gdk_window_get_root_coords.
func (recv *Window) GetRootCoords(x int32, y int32) (int32, int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	var outArgs [2]gi.Argument

	err := windowGetRootCoordsFunction_Set()
	if err == nil {
		windowGetRootCoordsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var windowGetRootOriginFunction *gi.Function
var windowGetRootOriginFunction_Once sync.Once

func windowGetRootOriginFunction_Set() error {
	var err error
	windowGetRootOriginFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetRootOriginFunction, err = windowStruct.InvokerNew("get_root_origin")
	})
	return err
}

// GetRootOrigin is a representation of the C type gdk_window_get_root_origin.
func (recv *Window) GetRootOrigin() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var outArgs [2]gi.Argument

	err := windowGetRootOriginFunction_Set()
	if err == nil {
		windowGetRootOriginFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var windowGetScaleFactorFunction *gi.Function
var windowGetScaleFactorFunction_Once sync.Once

func windowGetScaleFactorFunction_Set() error {
	var err error
	windowGetScaleFactorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetScaleFactorFunction, err = windowStruct.InvokerNew("get_scale_factor")
	})
	return err
}

// GetScaleFactor is a representation of the C type gdk_window_get_scale_factor.
func (recv *Window) GetScaleFactor() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetScaleFactorFunction_Set()
	if err == nil {
		ret = windowGetScaleFactorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var windowGetScreenFunction *gi.Function
var windowGetScreenFunction_Once sync.Once

func windowGetScreenFunction_Set() error {
	var err error
	windowGetScreenFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetScreenFunction, err = windowStruct.InvokerNew("get_screen")
	})
	return err
}

// GetScreen is a representation of the C type gdk_window_get_screen.
func (recv *Window) GetScreen() *Screen {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetScreenFunction_Set()
	if err == nil {
		ret = windowGetScreenFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Screen{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_source_events' : parameter 'source' of type 'InputSource' not supported

// UNSUPPORTED : C value 'gdk_window_get_state' : return type 'WindowState' not supported

var windowGetSupportMultideviceFunction *gi.Function
var windowGetSupportMultideviceFunction_Once sync.Once

func windowGetSupportMultideviceFunction_Set() error {
	var err error
	windowGetSupportMultideviceFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetSupportMultideviceFunction, err = windowStruct.InvokerNew("get_support_multidevice")
	})
	return err
}

// GetSupportMultidevice is a representation of the C type gdk_window_get_support_multidevice.
func (recv *Window) GetSupportMultidevice() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetSupportMultideviceFunction_Set()
	if err == nil {
		ret = windowGetSupportMultideviceFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowGetToplevelFunction *gi.Function
var windowGetToplevelFunction_Once sync.Once

func windowGetToplevelFunction_Set() error {
	var err error
	windowGetToplevelFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetToplevelFunction, err = windowStruct.InvokerNew("get_toplevel")
	})
	return err
}

// GetToplevel is a representation of the C type gdk_window_get_toplevel.
func (recv *Window) GetToplevel() *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetToplevelFunction_Set()
	if err == nil {
		ret = windowGetToplevelFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_type_hint' : return type 'WindowTypeHint' not supported

// UNSUPPORTED : C value 'gdk_window_get_update_area' : return type 'cairo.Region' not supported

// UNSUPPORTED : C value 'gdk_window_get_user_data' : parameter 'data' of type 'gpointer' not supported

// UNSUPPORTED : C value 'gdk_window_get_visible_region' : return type 'cairo.Region' not supported

var windowGetVisualFunction *gi.Function
var windowGetVisualFunction_Once sync.Once

func windowGetVisualFunction_Set() error {
	var err error
	windowGetVisualFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetVisualFunction, err = windowStruct.InvokerNew("get_visual")
	})
	return err
}

// GetVisual is a representation of the C type gdk_window_get_visual.
func (recv *Window) GetVisual() *Visual {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetVisualFunction_Set()
	if err == nil {
		ret = windowGetVisualFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Visual{}
	retGo.Native = ret.Pointer()

	return retGo
}

var windowGetWidthFunction *gi.Function
var windowGetWidthFunction_Once sync.Once

func windowGetWidthFunction_Set() error {
	var err error
	windowGetWidthFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowGetWidthFunction, err = windowStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type gdk_window_get_width.
func (recv *Window) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowGetWidthFunction_Set()
	if err == nil {
		ret = windowGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_window_get_window_type' : return type 'WindowType' not supported

var windowHasNativeFunction *gi.Function
var windowHasNativeFunction_Once sync.Once

func windowHasNativeFunction_Set() error {
	var err error
	windowHasNativeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowHasNativeFunction, err = windowStruct.InvokerNew("has_native")
	})
	return err
}

// HasNative is a representation of the C type gdk_window_has_native.
func (recv *Window) HasNative() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowHasNativeFunction_Set()
	if err == nil {
		ret = windowHasNativeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowHideFunction *gi.Function
var windowHideFunction_Once sync.Once

func windowHideFunction_Set() error {
	var err error
	windowHideFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowHideFunction, err = windowStruct.InvokerNew("hide")
	})
	return err
}

// Hide is a representation of the C type gdk_window_hide.
func (recv *Window) Hide() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowHideFunction_Set()
	if err == nil {
		windowHideFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowIconifyFunction *gi.Function
var windowIconifyFunction_Once sync.Once

func windowIconifyFunction_Set() error {
	var err error
	windowIconifyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIconifyFunction, err = windowStruct.InvokerNew("iconify")
	})
	return err
}

// Iconify is a representation of the C type gdk_window_iconify.
func (recv *Window) Iconify() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowIconifyFunction_Set()
	if err == nil {
		windowIconifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_input_shape_combine_region' : parameter 'shape_region' of type 'cairo.Region' not supported

// UNSUPPORTED : C value 'gdk_window_invalidate_maybe_recurse' : parameter 'region' of type 'cairo.Region' not supported

var windowInvalidateRectFunction *gi.Function
var windowInvalidateRectFunction_Once sync.Once

func windowInvalidateRectFunction_Set() error {
	var err error
	windowInvalidateRectFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowInvalidateRectFunction, err = windowStruct.InvokerNew("invalidate_rect")
	})
	return err
}

// InvalidateRect is a representation of the C type gdk_window_invalidate_rect.
func (recv *Window) InvalidateRect(rect *Rectangle, invalidateChildren bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(rect.Native)
	inArgs[2].SetBoolean(invalidateChildren)

	err := windowInvalidateRectFunction_Set()
	if err == nil {
		windowInvalidateRectFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_invalidate_region' : parameter 'region' of type 'cairo.Region' not supported

var windowIsDestroyedFunction *gi.Function
var windowIsDestroyedFunction_Once sync.Once

func windowIsDestroyedFunction_Set() error {
	var err error
	windowIsDestroyedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIsDestroyedFunction, err = windowStruct.InvokerNew("is_destroyed")
	})
	return err
}

// IsDestroyed is a representation of the C type gdk_window_is_destroyed.
func (recv *Window) IsDestroyed() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowIsDestroyedFunction_Set()
	if err == nil {
		ret = windowIsDestroyedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowIsInputOnlyFunction *gi.Function
var windowIsInputOnlyFunction_Once sync.Once

func windowIsInputOnlyFunction_Set() error {
	var err error
	windowIsInputOnlyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIsInputOnlyFunction, err = windowStruct.InvokerNew("is_input_only")
	})
	return err
}

// IsInputOnly is a representation of the C type gdk_window_is_input_only.
func (recv *Window) IsInputOnly() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowIsInputOnlyFunction_Set()
	if err == nil {
		ret = windowIsInputOnlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowIsShapedFunction *gi.Function
var windowIsShapedFunction_Once sync.Once

func windowIsShapedFunction_Set() error {
	var err error
	windowIsShapedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIsShapedFunction, err = windowStruct.InvokerNew("is_shaped")
	})
	return err
}

// IsShaped is a representation of the C type gdk_window_is_shaped.
func (recv *Window) IsShaped() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowIsShapedFunction_Set()
	if err == nil {
		ret = windowIsShapedFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowIsViewableFunction *gi.Function
var windowIsViewableFunction_Once sync.Once

func windowIsViewableFunction_Set() error {
	var err error
	windowIsViewableFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIsViewableFunction, err = windowStruct.InvokerNew("is_viewable")
	})
	return err
}

// IsViewable is a representation of the C type gdk_window_is_viewable.
func (recv *Window) IsViewable() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowIsViewableFunction_Set()
	if err == nil {
		ret = windowIsViewableFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowIsVisibleFunction *gi.Function
var windowIsVisibleFunction_Once sync.Once

func windowIsVisibleFunction_Set() error {
	var err error
	windowIsVisibleFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowIsVisibleFunction, err = windowStruct.InvokerNew("is_visible")
	})
	return err
}

// IsVisible is a representation of the C type gdk_window_is_visible.
func (recv *Window) IsVisible() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	var ret gi.Argument

	err := windowIsVisibleFunction_Set()
	if err == nil {
		ret = windowIsVisibleFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowLowerFunction *gi.Function
var windowLowerFunction_Once sync.Once

func windowLowerFunction_Set() error {
	var err error
	windowLowerFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowLowerFunction, err = windowStruct.InvokerNew("lower")
	})
	return err
}

// Lower is a representation of the C type gdk_window_lower.
func (recv *Window) Lower() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowLowerFunction_Set()
	if err == nil {
		windowLowerFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_mark_paint_from_clip' : parameter 'cr' of type 'cairo.Context' not supported

var windowMaximizeFunction *gi.Function
var windowMaximizeFunction_Once sync.Once

func windowMaximizeFunction_Set() error {
	var err error
	windowMaximizeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowMaximizeFunction, err = windowStruct.InvokerNew("maximize")
	})
	return err
}

// Maximize is a representation of the C type gdk_window_maximize.
func (recv *Window) Maximize() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowMaximizeFunction_Set()
	if err == nil {
		windowMaximizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowMergeChildInputShapesFunction *gi.Function
var windowMergeChildInputShapesFunction_Once sync.Once

func windowMergeChildInputShapesFunction_Set() error {
	var err error
	windowMergeChildInputShapesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowMergeChildInputShapesFunction, err = windowStruct.InvokerNew("merge_child_input_shapes")
	})
	return err
}

// MergeChildInputShapes is a representation of the C type gdk_window_merge_child_input_shapes.
func (recv *Window) MergeChildInputShapes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowMergeChildInputShapesFunction_Set()
	if err == nil {
		windowMergeChildInputShapesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowMergeChildShapesFunction *gi.Function
var windowMergeChildShapesFunction_Once sync.Once

func windowMergeChildShapesFunction_Set() error {
	var err error
	windowMergeChildShapesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowMergeChildShapesFunction, err = windowStruct.InvokerNew("merge_child_shapes")
	})
	return err
}

// MergeChildShapes is a representation of the C type gdk_window_merge_child_shapes.
func (recv *Window) MergeChildShapes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowMergeChildShapesFunction_Set()
	if err == nil {
		windowMergeChildShapesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowMoveFunction *gi.Function
var windowMoveFunction_Once sync.Once

func windowMoveFunction_Set() error {
	var err error
	windowMoveFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowMoveFunction, err = windowStruct.InvokerNew("move")
	})
	return err
}

// Move is a representation of the C type gdk_window_move.
func (recv *Window) Move(x int32, y int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)

	err := windowMoveFunction_Set()
	if err == nil {
		windowMoveFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_move_region' : parameter 'region' of type 'cairo.Region' not supported

var windowMoveResizeFunction *gi.Function
var windowMoveResizeFunction_Once sync.Once

func windowMoveResizeFunction_Set() error {
	var err error
	windowMoveResizeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowMoveResizeFunction, err = windowStruct.InvokerNew("move_resize")
	})
	return err
}

// MoveResize is a representation of the C type gdk_window_move_resize.
func (recv *Window) MoveResize(x int32, y int32, width int32, height int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(x)
	inArgs[2].SetInt32(y)
	inArgs[3].SetInt32(width)
	inArgs[4].SetInt32(height)

	err := windowMoveResizeFunction_Set()
	if err == nil {
		windowMoveResizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_move_to_rect' : parameter 'rect_anchor' of type 'Gravity' not supported

// UNSUPPORTED : C value 'gdk_window_peek_children' : return type 'GLib.List' not supported

var windowProcessUpdatesFunction *gi.Function
var windowProcessUpdatesFunction_Once sync.Once

func windowProcessUpdatesFunction_Set() error {
	var err error
	windowProcessUpdatesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowProcessUpdatesFunction, err = windowStruct.InvokerNew("process_updates")
	})
	return err
}

// ProcessUpdates is a representation of the C type gdk_window_process_updates.
func (recv *Window) ProcessUpdates(updateChildren bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(updateChildren)

	err := windowProcessUpdatesFunction_Set()
	if err == nil {
		windowProcessUpdatesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowRaiseFunction *gi.Function
var windowRaiseFunction_Once sync.Once

func windowRaiseFunction_Set() error {
	var err error
	windowRaiseFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowRaiseFunction, err = windowStruct.InvokerNew("raise")
	})
	return err
}

// Raise is a representation of the C type gdk_window_raise.
func (recv *Window) Raise() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowRaiseFunction_Set()
	if err == nil {
		windowRaiseFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowRegisterDndFunction *gi.Function
var windowRegisterDndFunction_Once sync.Once

func windowRegisterDndFunction_Set() error {
	var err error
	windowRegisterDndFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowRegisterDndFunction, err = windowStruct.InvokerNew("register_dnd")
	})
	return err
}

// RegisterDnd is a representation of the C type gdk_window_register_dnd.
func (recv *Window) RegisterDnd() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowRegisterDndFunction_Set()
	if err == nil {
		windowRegisterDndFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_remove_filter' : parameter 'function' of type 'FilterFunc' not supported

var windowReparentFunction *gi.Function
var windowReparentFunction_Once sync.Once

func windowReparentFunction_Set() error {
	var err error
	windowReparentFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowReparentFunction, err = windowStruct.InvokerNew("reparent")
	})
	return err
}

// Reparent is a representation of the C type gdk_window_reparent.
func (recv *Window) Reparent(newParent *Window, x int32, y int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(newParent.Native)
	inArgs[2].SetInt32(x)
	inArgs[3].SetInt32(y)

	err := windowReparentFunction_Set()
	if err == nil {
		windowReparentFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowResizeFunction *gi.Function
var windowResizeFunction_Once sync.Once

func windowResizeFunction_Set() error {
	var err error
	windowResizeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowResizeFunction, err = windowStruct.InvokerNew("resize")
	})
	return err
}

// Resize is a representation of the C type gdk_window_resize.
func (recv *Window) Resize(width int32, height int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(width)
	inArgs[2].SetInt32(height)

	err := windowResizeFunction_Set()
	if err == nil {
		windowResizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowRestackFunction *gi.Function
var windowRestackFunction_Once sync.Once

func windowRestackFunction_Set() error {
	var err error
	windowRestackFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowRestackFunction, err = windowStruct.InvokerNew("restack")
	})
	return err
}

// Restack is a representation of the C type gdk_window_restack.
func (recv *Window) Restack(sibling *Window, above bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(sibling.Native)
	inArgs[2].SetBoolean(above)

	err := windowRestackFunction_Set()
	if err == nil {
		windowRestackFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowScrollFunction *gi.Function
var windowScrollFunction_Once sync.Once

func windowScrollFunction_Set() error {
	var err error
	windowScrollFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowScrollFunction, err = windowStruct.InvokerNew("scroll")
	})
	return err
}

// Scroll is a representation of the C type gdk_window_scroll.
func (recv *Window) Scroll(dx int32, dy int32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(dx)
	inArgs[2].SetInt32(dy)

	err := windowScrollFunction_Set()
	if err == nil {
		windowScrollFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetAcceptFocusFunction *gi.Function
var windowSetAcceptFocusFunction_Once sync.Once

func windowSetAcceptFocusFunction_Set() error {
	var err error
	windowSetAcceptFocusFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetAcceptFocusFunction, err = windowStruct.InvokerNew("set_accept_focus")
	})
	return err
}

// SetAcceptFocus is a representation of the C type gdk_window_set_accept_focus.
func (recv *Window) SetAcceptFocus(acceptFocus bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(acceptFocus)

	err := windowSetAcceptFocusFunction_Set()
	if err == nil {
		windowSetAcceptFocusFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetBackgroundFunction *gi.Function
var windowSetBackgroundFunction_Once sync.Once

func windowSetBackgroundFunction_Set() error {
	var err error
	windowSetBackgroundFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetBackgroundFunction, err = windowStruct.InvokerNew("set_background")
	})
	return err
}

// SetBackground is a representation of the C type gdk_window_set_background.
func (recv *Window) SetBackground(color *Color) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(color.Native)

	err := windowSetBackgroundFunction_Set()
	if err == nil {
		windowSetBackgroundFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_background_pattern' : parameter 'pattern' of type 'cairo.Pattern' not supported

var windowSetBackgroundRgbaFunction *gi.Function
var windowSetBackgroundRgbaFunction_Once sync.Once

func windowSetBackgroundRgbaFunction_Set() error {
	var err error
	windowSetBackgroundRgbaFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetBackgroundRgbaFunction, err = windowStruct.InvokerNew("set_background_rgba")
	})
	return err
}

// SetBackgroundRgba is a representation of the C type gdk_window_set_background_rgba.
func (recv *Window) SetBackgroundRgba(rgba *RGBA) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(rgba.Native)

	err := windowSetBackgroundRgbaFunction_Set()
	if err == nil {
		windowSetBackgroundRgbaFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetChildInputShapesFunction *gi.Function
var windowSetChildInputShapesFunction_Once sync.Once

func windowSetChildInputShapesFunction_Set() error {
	var err error
	windowSetChildInputShapesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetChildInputShapesFunction, err = windowStruct.InvokerNew("set_child_input_shapes")
	})
	return err
}

// SetChildInputShapes is a representation of the C type gdk_window_set_child_input_shapes.
func (recv *Window) SetChildInputShapes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowSetChildInputShapesFunction_Set()
	if err == nil {
		windowSetChildInputShapesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetChildShapesFunction *gi.Function
var windowSetChildShapesFunction_Once sync.Once

func windowSetChildShapesFunction_Set() error {
	var err error
	windowSetChildShapesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetChildShapesFunction, err = windowStruct.InvokerNew("set_child_shapes")
	})
	return err
}

// SetChildShapes is a representation of the C type gdk_window_set_child_shapes.
func (recv *Window) SetChildShapes() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowSetChildShapesFunction_Set()
	if err == nil {
		windowSetChildShapesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetCompositedFunction *gi.Function
var windowSetCompositedFunction_Once sync.Once

func windowSetCompositedFunction_Set() error {
	var err error
	windowSetCompositedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetCompositedFunction, err = windowStruct.InvokerNew("set_composited")
	})
	return err
}

// SetComposited is a representation of the C type gdk_window_set_composited.
func (recv *Window) SetComposited(composited bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(composited)

	err := windowSetCompositedFunction_Set()
	if err == nil {
		windowSetCompositedFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetCursorFunction *gi.Function
var windowSetCursorFunction_Once sync.Once

func windowSetCursorFunction_Set() error {
	var err error
	windowSetCursorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetCursorFunction, err = windowStruct.InvokerNew("set_cursor")
	})
	return err
}

// SetCursor is a representation of the C type gdk_window_set_cursor.
func (recv *Window) SetCursor(cursor *Cursor) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(cursor.Native)

	err := windowSetCursorFunction_Set()
	if err == nil {
		windowSetCursorFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_decorations' : parameter 'decorations' of type 'WMDecoration' not supported

var windowSetDeviceCursorFunction *gi.Function
var windowSetDeviceCursorFunction_Once sync.Once

func windowSetDeviceCursorFunction_Set() error {
	var err error
	windowSetDeviceCursorFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetDeviceCursorFunction, err = windowStruct.InvokerNew("set_device_cursor")
	})
	return err
}

// SetDeviceCursor is a representation of the C type gdk_window_set_device_cursor.
func (recv *Window) SetDeviceCursor(device *Device, cursor *Cursor) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(device.Native)
	inArgs[2].SetPointer(cursor.Native)

	err := windowSetDeviceCursorFunction_Set()
	if err == nil {
		windowSetDeviceCursorFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_device_events' : parameter 'event_mask' of type 'EventMask' not supported

var windowSetEventCompressionFunction *gi.Function
var windowSetEventCompressionFunction_Once sync.Once

func windowSetEventCompressionFunction_Set() error {
	var err error
	windowSetEventCompressionFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetEventCompressionFunction, err = windowStruct.InvokerNew("set_event_compression")
	})
	return err
}

// SetEventCompression is a representation of the C type gdk_window_set_event_compression.
func (recv *Window) SetEventCompression(eventCompression bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(eventCompression)

	err := windowSetEventCompressionFunction_Set()
	if err == nil {
		windowSetEventCompressionFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_events' : parameter 'event_mask' of type 'EventMask' not supported

var windowSetFocusOnMapFunction *gi.Function
var windowSetFocusOnMapFunction_Once sync.Once

func windowSetFocusOnMapFunction_Set() error {
	var err error
	windowSetFocusOnMapFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetFocusOnMapFunction, err = windowStruct.InvokerNew("set_focus_on_map")
	})
	return err
}

// SetFocusOnMap is a representation of the C type gdk_window_set_focus_on_map.
func (recv *Window) SetFocusOnMap(focusOnMap bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(focusOnMap)

	err := windowSetFocusOnMapFunction_Set()
	if err == nil {
		windowSetFocusOnMapFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_fullscreen_mode' : parameter 'mode' of type 'FullscreenMode' not supported

// UNSUPPORTED : C value 'gdk_window_set_functions' : parameter 'functions' of type 'WMFunction' not supported

// UNSUPPORTED : C value 'gdk_window_set_geometry_hints' : parameter 'geom_mask' of type 'WindowHints' not supported

var windowSetGroupFunction *gi.Function
var windowSetGroupFunction_Once sync.Once

func windowSetGroupFunction_Set() error {
	var err error
	windowSetGroupFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetGroupFunction, err = windowStruct.InvokerNew("set_group")
	})
	return err
}

// SetGroup is a representation of the C type gdk_window_set_group.
func (recv *Window) SetGroup(leader *Window) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(leader.Native)

	err := windowSetGroupFunction_Set()
	if err == nil {
		windowSetGroupFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_icon_list' : parameter 'pixbufs' of type 'GLib.List' not supported

var windowSetIconNameFunction *gi.Function
var windowSetIconNameFunction_Once sync.Once

func windowSetIconNameFunction_Set() error {
	var err error
	windowSetIconNameFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetIconNameFunction, err = windowStruct.InvokerNew("set_icon_name")
	})
	return err
}

// SetIconName is a representation of the C type gdk_window_set_icon_name.
func (recv *Window) SetIconName(name string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(name)

	err := windowSetIconNameFunction_Set()
	if err == nil {
		windowSetIconNameFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_invalidate_handler' : parameter 'handler' of type 'WindowInvalidateHandlerFunc' not supported

var windowSetKeepAboveFunction *gi.Function
var windowSetKeepAboveFunction_Once sync.Once

func windowSetKeepAboveFunction_Set() error {
	var err error
	windowSetKeepAboveFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetKeepAboveFunction, err = windowStruct.InvokerNew("set_keep_above")
	})
	return err
}

// SetKeepAbove is a representation of the C type gdk_window_set_keep_above.
func (recv *Window) SetKeepAbove(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(setting)

	err := windowSetKeepAboveFunction_Set()
	if err == nil {
		windowSetKeepAboveFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetKeepBelowFunction *gi.Function
var windowSetKeepBelowFunction_Once sync.Once

func windowSetKeepBelowFunction_Set() error {
	var err error
	windowSetKeepBelowFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetKeepBelowFunction, err = windowStruct.InvokerNew("set_keep_below")
	})
	return err
}

// SetKeepBelow is a representation of the C type gdk_window_set_keep_below.
func (recv *Window) SetKeepBelow(setting bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(setting)

	err := windowSetKeepBelowFunction_Set()
	if err == nil {
		windowSetKeepBelowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetModalHintFunction *gi.Function
var windowSetModalHintFunction_Once sync.Once

func windowSetModalHintFunction_Set() error {
	var err error
	windowSetModalHintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetModalHintFunction, err = windowStruct.InvokerNew("set_modal_hint")
	})
	return err
}

// SetModalHint is a representation of the C type gdk_window_set_modal_hint.
func (recv *Window) SetModalHint(modal bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(modal)

	err := windowSetModalHintFunction_Set()
	if err == nil {
		windowSetModalHintFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetOpacityFunction *gi.Function
var windowSetOpacityFunction_Once sync.Once

func windowSetOpacityFunction_Set() error {
	var err error
	windowSetOpacityFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetOpacityFunction, err = windowStruct.InvokerNew("set_opacity")
	})
	return err
}

// SetOpacity is a representation of the C type gdk_window_set_opacity.
func (recv *Window) SetOpacity(opacity float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetFloat64(opacity)

	err := windowSetOpacityFunction_Set()
	if err == nil {
		windowSetOpacityFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_opaque_region' : parameter 'region' of type 'cairo.Region' not supported

var windowSetOverrideRedirectFunction *gi.Function
var windowSetOverrideRedirectFunction_Once sync.Once

func windowSetOverrideRedirectFunction_Set() error {
	var err error
	windowSetOverrideRedirectFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetOverrideRedirectFunction, err = windowStruct.InvokerNew("set_override_redirect")
	})
	return err
}

// SetOverrideRedirect is a representation of the C type gdk_window_set_override_redirect.
func (recv *Window) SetOverrideRedirect(overrideRedirect bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(overrideRedirect)

	err := windowSetOverrideRedirectFunction_Set()
	if err == nil {
		windowSetOverrideRedirectFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetPassThroughFunction *gi.Function
var windowSetPassThroughFunction_Once sync.Once

func windowSetPassThroughFunction_Set() error {
	var err error
	windowSetPassThroughFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetPassThroughFunction, err = windowStruct.InvokerNew("set_pass_through")
	})
	return err
}

// SetPassThrough is a representation of the C type gdk_window_set_pass_through.
func (recv *Window) SetPassThrough(passThrough bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(passThrough)

	err := windowSetPassThroughFunction_Set()
	if err == nil {
		windowSetPassThroughFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetRoleFunction *gi.Function
var windowSetRoleFunction_Once sync.Once

func windowSetRoleFunction_Set() error {
	var err error
	windowSetRoleFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetRoleFunction, err = windowStruct.InvokerNew("set_role")
	})
	return err
}

// SetRole is a representation of the C type gdk_window_set_role.
func (recv *Window) SetRole(role string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(role)

	err := windowSetRoleFunction_Set()
	if err == nil {
		windowSetRoleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetShadowWidthFunction *gi.Function
var windowSetShadowWidthFunction_Once sync.Once

func windowSetShadowWidthFunction_Set() error {
	var err error
	windowSetShadowWidthFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetShadowWidthFunction, err = windowStruct.InvokerNew("set_shadow_width")
	})
	return err
}

// SetShadowWidth is a representation of the C type gdk_window_set_shadow_width.
func (recv *Window) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetInt32(left)
	inArgs[2].SetInt32(right)
	inArgs[3].SetInt32(top)
	inArgs[4].SetInt32(bottom)

	err := windowSetShadowWidthFunction_Set()
	if err == nil {
		windowSetShadowWidthFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetSkipPagerHintFunction *gi.Function
var windowSetSkipPagerHintFunction_Once sync.Once

func windowSetSkipPagerHintFunction_Set() error {
	var err error
	windowSetSkipPagerHintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetSkipPagerHintFunction, err = windowStruct.InvokerNew("set_skip_pager_hint")
	})
	return err
}

// SetSkipPagerHint is a representation of the C type gdk_window_set_skip_pager_hint.
func (recv *Window) SetSkipPagerHint(skipsPager bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(skipsPager)

	err := windowSetSkipPagerHintFunction_Set()
	if err == nil {
		windowSetSkipPagerHintFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetSkipTaskbarHintFunction *gi.Function
var windowSetSkipTaskbarHintFunction_Once sync.Once

func windowSetSkipTaskbarHintFunction_Set() error {
	var err error
	windowSetSkipTaskbarHintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetSkipTaskbarHintFunction, err = windowStruct.InvokerNew("set_skip_taskbar_hint")
	})
	return err
}

// SetSkipTaskbarHint is a representation of the C type gdk_window_set_skip_taskbar_hint.
func (recv *Window) SetSkipTaskbarHint(skipsTaskbar bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(skipsTaskbar)

	err := windowSetSkipTaskbarHintFunction_Set()
	if err == nil {
		windowSetSkipTaskbarHintFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_source_events' : parameter 'source' of type 'InputSource' not supported

var windowSetStartupIdFunction *gi.Function
var windowSetStartupIdFunction_Once sync.Once

func windowSetStartupIdFunction_Set() error {
	var err error
	windowSetStartupIdFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetStartupIdFunction, err = windowStruct.InvokerNew("set_startup_id")
	})
	return err
}

// SetStartupId is a representation of the C type gdk_window_set_startup_id.
func (recv *Window) SetStartupId(startupId string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(startupId)

	err := windowSetStartupIdFunction_Set()
	if err == nil {
		windowSetStartupIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetStaticGravitiesFunction *gi.Function
var windowSetStaticGravitiesFunction_Once sync.Once

func windowSetStaticGravitiesFunction_Set() error {
	var err error
	windowSetStaticGravitiesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetStaticGravitiesFunction, err = windowStruct.InvokerNew("set_static_gravities")
	})
	return err
}

// SetStaticGravities is a representation of the C type gdk_window_set_static_gravities.
func (recv *Window) SetStaticGravities(useStatic bool) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(useStatic)

	var ret gi.Argument

	err := windowSetStaticGravitiesFunction_Set()
	if err == nil {
		ret = windowSetStaticGravitiesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var windowSetSupportMultideviceFunction *gi.Function
var windowSetSupportMultideviceFunction_Once sync.Once

func windowSetSupportMultideviceFunction_Set() error {
	var err error
	windowSetSupportMultideviceFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetSupportMultideviceFunction, err = windowStruct.InvokerNew("set_support_multidevice")
	})
	return err
}

// SetSupportMultidevice is a representation of the C type gdk_window_set_support_multidevice.
func (recv *Window) SetSupportMultidevice(supportMultidevice bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(supportMultidevice)

	err := windowSetSupportMultideviceFunction_Set()
	if err == nil {
		windowSetSupportMultideviceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetTitleFunction *gi.Function
var windowSetTitleFunction_Once sync.Once

func windowSetTitleFunction_Set() error {
	var err error
	windowSetTitleFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetTitleFunction, err = windowStruct.InvokerNew("set_title")
	})
	return err
}

// SetTitle is a representation of the C type gdk_window_set_title.
func (recv *Window) SetTitle(title string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetString(title)

	err := windowSetTitleFunction_Set()
	if err == nil {
		windowSetTitleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowSetTransientForFunction *gi.Function
var windowSetTransientForFunction_Once sync.Once

func windowSetTransientForFunction_Set() error {
	var err error
	windowSetTransientForFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetTransientForFunction, err = windowStruct.InvokerNew("set_transient_for")
	})
	return err
}

// SetTransientFor is a representation of the C type gdk_window_set_transient_for.
func (recv *Window) SetTransientFor(parent *Window) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetPointer(parent.Native)

	err := windowSetTransientForFunction_Set()
	if err == nil {
		windowSetTransientForFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_type_hint' : parameter 'hint' of type 'WindowTypeHint' not supported

var windowSetUrgencyHintFunction *gi.Function
var windowSetUrgencyHintFunction_Once sync.Once

func windowSetUrgencyHintFunction_Set() error {
	var err error
	windowSetUrgencyHintFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowSetUrgencyHintFunction, err = windowStruct.InvokerNew("set_urgency_hint")
	})
	return err
}

// SetUrgencyHint is a representation of the C type gdk_window_set_urgency_hint.
func (recv *Window) SetUrgencyHint(urgent bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.Native)
	inArgs[1].SetBoolean(urgent)

	err := windowSetUrgencyHintFunction_Set()
	if err == nil {
		windowSetUrgencyHintFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_set_user_data' : parameter 'user_data' of type 'GObject.Object' not supported

// UNSUPPORTED : C value 'gdk_window_shape_combine_region' : parameter 'shape_region' of type 'cairo.Region' not supported

var windowShowFunction *gi.Function
var windowShowFunction_Once sync.Once

func windowShowFunction_Set() error {
	var err error
	windowShowFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowShowFunction, err = windowStruct.InvokerNew("show")
	})
	return err
}

// Show is a representation of the C type gdk_window_show.
func (recv *Window) Show() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowShowFunction_Set()
	if err == nil {
		windowShowFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowShowUnraisedFunction *gi.Function
var windowShowUnraisedFunction_Once sync.Once

func windowShowUnraisedFunction_Set() error {
	var err error
	windowShowUnraisedFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowShowUnraisedFunction, err = windowStruct.InvokerNew("show_unraised")
	})
	return err
}

// ShowUnraised is a representation of the C type gdk_window_show_unraised.
func (recv *Window) ShowUnraised() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowShowUnraisedFunction_Set()
	if err == nil {
		windowShowUnraisedFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_window_show_window_menu' : parameter 'event' of type 'Event' not supported

var windowStickFunction *gi.Function
var windowStickFunction_Once sync.Once

func windowStickFunction_Set() error {
	var err error
	windowStickFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowStickFunction, err = windowStruct.InvokerNew("stick")
	})
	return err
}

// Stick is a representation of the C type gdk_window_stick.
func (recv *Window) Stick() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowStickFunction_Set()
	if err == nil {
		windowStickFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowThawToplevelUpdatesLibgtkOnlyFunction *gi.Function
var windowThawToplevelUpdatesLibgtkOnlyFunction_Once sync.Once

func windowThawToplevelUpdatesLibgtkOnlyFunction_Set() error {
	var err error
	windowThawToplevelUpdatesLibgtkOnlyFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowThawToplevelUpdatesLibgtkOnlyFunction, err = windowStruct.InvokerNew("thaw_toplevel_updates_libgtk_only")
	})
	return err
}

// ThawToplevelUpdatesLibgtkOnly is a representation of the C type gdk_window_thaw_toplevel_updates_libgtk_only.
func (recv *Window) ThawToplevelUpdatesLibgtkOnly() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowThawToplevelUpdatesLibgtkOnlyFunction_Set()
	if err == nil {
		windowThawToplevelUpdatesLibgtkOnlyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowThawUpdatesFunction *gi.Function
var windowThawUpdatesFunction_Once sync.Once

func windowThawUpdatesFunction_Set() error {
	var err error
	windowThawUpdatesFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowThawUpdatesFunction, err = windowStruct.InvokerNew("thaw_updates")
	})
	return err
}

// ThawUpdates is a representation of the C type gdk_window_thaw_updates.
func (recv *Window) ThawUpdates() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowThawUpdatesFunction_Set()
	if err == nil {
		windowThawUpdatesFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowUnfullscreenFunction *gi.Function
var windowUnfullscreenFunction_Once sync.Once

func windowUnfullscreenFunction_Set() error {
	var err error
	windowUnfullscreenFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowUnfullscreenFunction, err = windowStruct.InvokerNew("unfullscreen")
	})
	return err
}

// Unfullscreen is a representation of the C type gdk_window_unfullscreen.
func (recv *Window) Unfullscreen() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowUnfullscreenFunction_Set()
	if err == nil {
		windowUnfullscreenFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowUnmaximizeFunction *gi.Function
var windowUnmaximizeFunction_Once sync.Once

func windowUnmaximizeFunction_Set() error {
	var err error
	windowUnmaximizeFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowUnmaximizeFunction, err = windowStruct.InvokerNew("unmaximize")
	})
	return err
}

// Unmaximize is a representation of the C type gdk_window_unmaximize.
func (recv *Window) Unmaximize() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowUnmaximizeFunction_Set()
	if err == nil {
		windowUnmaximizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowUnstickFunction *gi.Function
var windowUnstickFunction_Once sync.Once

func windowUnstickFunction_Set() error {
	var err error
	windowUnstickFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowUnstickFunction, err = windowStruct.InvokerNew("unstick")
	})
	return err
}

// Unstick is a representation of the C type gdk_window_unstick.
func (recv *Window) Unstick() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowUnstickFunction_Set()
	if err == nil {
		windowUnstickFunction.Invoke(inArgs[:], nil)
	}

	return
}

var windowWithdrawFunction *gi.Function
var windowWithdrawFunction_Once sync.Once

func windowWithdrawFunction_Set() error {
	var err error
	windowWithdrawFunction_Once.Do(func() {
		err = windowStruct_Set()
		if err != nil {
			return
		}
		windowWithdrawFunction, err = windowStruct.InvokerNew("withdraw")
	})
	return err
}

// Withdraw is a representation of the C type gdk_window_withdraw.
func (recv *Window) Withdraw() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.Native)

	err := windowWithdrawFunction_Set()
	if err == nil {
		windowWithdrawFunction.Invoke(inArgs[:], nil)
	}

	return
}
