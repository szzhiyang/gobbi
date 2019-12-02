// Code generated - DO NOT EDIT.

package gdk

import (
	gi "github.com/pekim/gobbi/internal/gi"
	glib "github.com/pekim/gobbi/lib/glib"
	"sync"
)

// UNSUPPORTED : C value 'gdk_add_option_entries_libgtk_only' : parameter 'group' of type 'GLib.OptionGroup' not supported

var atomInternFunction *gi.Function
var atomInternFunction_Once sync.Once

func atomInternFunction_Set() error {
	var err error
	atomInternFunction_Once.Do(func() {
		atomInternFunction, err = gi.FunctionInvokerNew("Gdk", "atom_intern")
	})
	return err
}

// AtomIntern is a representation of the C type gdk_atom_intern.
func AtomIntern(atomName string, onlyIfExists bool) *Atom {
	var inArgs [2]gi.Argument
	inArgs[0].SetString(atomName)
	inArgs[1].SetBoolean(onlyIfExists)

	var ret gi.Argument

	err := atomInternFunction_Set()
	if err == nil {
		ret = atomInternFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Atom{}
	retGo.Native = ret.Pointer()

	return retGo
}

var atomInternStaticStringFunction *gi.Function
var atomInternStaticStringFunction_Once sync.Once

func atomInternStaticStringFunction_Set() error {
	var err error
	atomInternStaticStringFunction_Once.Do(func() {
		atomInternStaticStringFunction, err = gi.FunctionInvokerNew("Gdk", "atom_intern_static_string")
	})
	return err
}

// AtomInternStaticString is a representation of the C type gdk_atom_intern_static_string.
func AtomInternStaticString(atomName string) *Atom {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(atomName)

	var ret gi.Argument

	err := atomInternStaticStringFunction_Set()
	if err == nil {
		ret = atomInternStaticStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Atom{}
	retGo.Native = ret.Pointer()

	return retGo
}

var beepFunction *gi.Function
var beepFunction_Once sync.Once

func beepFunction_Set() error {
	var err error
	beepFunction_Once.Do(func() {
		beepFunction, err = gi.FunctionInvokerNew("Gdk", "beep")
	})
	return err
}

// Beep is a representation of the C type gdk_beep.
func Beep() {

	err := beepFunction_Set()
	if err == nil {
		beepFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_cairo_create' : return type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_draw_from_gl' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_clip_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_get_drawing_context' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_rectangle' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_region_create_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_color' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_pixbuf' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_rgba' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_set_source_window' : parameter 'cr' of type 'cairo.Context' not supported

// UNSUPPORTED : C value 'gdk_cairo_surface_create_from_pixbuf' : parameter 'pixbuf' of type 'GdkPixbuf.Pixbuf' not supported

var colorParseFunction *gi.Function
var colorParseFunction_Once sync.Once

func colorParseFunction_Set() error {
	var err error
	colorParseFunction_Once.Do(func() {
		colorParseFunction, err = gi.FunctionInvokerNew("Gdk", "color_parse")
	})
	return err
}

// ColorParse is a representation of the C type gdk_color_parse.
func ColorParse(spec string) (bool, *Color) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(spec)

	var outArgs [1]gi.Argument
	var ret gi.Argument

	err := colorParseFunction_Set()
	if err == nil {
		ret = colorParseFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := &Color{}
	out0.Native = outArgs[0].Pointer()

	return retGo, out0
}

var disableMultideviceFunction *gi.Function
var disableMultideviceFunction_Once sync.Once

func disableMultideviceFunction_Set() error {
	var err error
	disableMultideviceFunction_Once.Do(func() {
		disableMultideviceFunction, err = gi.FunctionInvokerNew("Gdk", "disable_multidevice")
	})
	return err
}

// DisableMultidevice is a representation of the C type gdk_disable_multidevice.
func DisableMultidevice() {

	err := disableMultideviceFunction_Set()
	if err == nil {
		disableMultideviceFunction.Invoke(nil, nil)
	}

	return
}

var dragAbortFunction *gi.Function
var dragAbortFunction_Once sync.Once

func dragAbortFunction_Set() error {
	var err error
	dragAbortFunction_Once.Do(func() {
		dragAbortFunction, err = gi.FunctionInvokerNew("Gdk", "drag_abort")
	})
	return err
}

// DragAbort is a representation of the C type gdk_drag_abort.
func DragAbort(context *DragContext, time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native)
	inArgs[1].SetUint32(time)

	err := dragAbortFunction_Set()
	if err == nil {
		dragAbortFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_drag_begin' : parameter 'targets' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_for_device' : parameter 'targets' of type 'GLib.List' not supported

// UNSUPPORTED : C value 'gdk_drag_begin_from_point' : parameter 'targets' of type 'GLib.List' not supported

var dragDropFunction *gi.Function
var dragDropFunction_Once sync.Once

func dragDropFunction_Set() error {
	var err error
	dragDropFunction_Once.Do(func() {
		dragDropFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop")
	})
	return err
}

// DragDrop is a representation of the C type gdk_drag_drop.
func DragDrop(context *DragContext, time uint32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native)
	inArgs[1].SetUint32(time)

	err := dragDropFunction_Set()
	if err == nil {
		dragDropFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragDropDoneFunction *gi.Function
var dragDropDoneFunction_Once sync.Once

func dragDropDoneFunction_Set() error {
	var err error
	dragDropDoneFunction_Once.Do(func() {
		dragDropDoneFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop_done")
	})
	return err
}

// DragDropDone is a representation of the C type gdk_drag_drop_done.
func DragDropDone(context *DragContext, success bool) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(context.Native)
	inArgs[1].SetBoolean(success)

	err := dragDropDoneFunction_Set()
	if err == nil {
		dragDropDoneFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dragDropSucceededFunction *gi.Function
var dragDropSucceededFunction_Once sync.Once

func dragDropSucceededFunction_Set() error {
	var err error
	dragDropSucceededFunction_Once.Do(func() {
		dragDropSucceededFunction, err = gi.FunctionInvokerNew("Gdk", "drag_drop_succeeded")
	})
	return err
}

// DragDropSucceeded is a representation of the C type gdk_drag_drop_succeeded.
func DragDropSucceeded(context *DragContext) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native)

	var ret gi.Argument

	err := dragDropSucceededFunction_Set()
	if err == nil {
		ret = dragDropSucceededFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// UNSUPPORTED : C value 'gdk_drag_find_window_for_screen' : parameter 'protocol' of type 'DragProtocol' not supported

var dragGetSelectionFunction *gi.Function
var dragGetSelectionFunction_Once sync.Once

func dragGetSelectionFunction_Set() error {
	var err error
	dragGetSelectionFunction_Once.Do(func() {
		dragGetSelectionFunction, err = gi.FunctionInvokerNew("Gdk", "drag_get_selection")
	})
	return err
}

// DragGetSelection is a representation of the C type gdk_drag_get_selection.
func DragGetSelection(context *DragContext) *Atom {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(context.Native)

	var ret gi.Argument

	err := dragGetSelectionFunction_Set()
	if err == nil {
		ret = dragGetSelectionFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Atom{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_drag_motion' : parameter 'protocol' of type 'DragProtocol' not supported

// UNSUPPORTED : C value 'gdk_drag_status' : parameter 'action' of type 'DragAction' not supported

var dropFinishFunction *gi.Function
var dropFinishFunction_Once sync.Once

func dropFinishFunction_Set() error {
	var err error
	dropFinishFunction_Once.Do(func() {
		dropFinishFunction, err = gi.FunctionInvokerNew("Gdk", "drop_finish")
	})
	return err
}

// DropFinish is a representation of the C type gdk_drop_finish.
func DropFinish(context *DragContext, success bool, time uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native)
	inArgs[1].SetBoolean(success)
	inArgs[2].SetUint32(time)

	err := dropFinishFunction_Set()
	if err == nil {
		dropFinishFunction.Invoke(inArgs[:], nil)
	}

	return
}

var dropReplyFunction *gi.Function
var dropReplyFunction_Once sync.Once

func dropReplyFunction_Set() error {
	var err error
	dropReplyFunction_Once.Do(func() {
		dropReplyFunction, err = gi.FunctionInvokerNew("Gdk", "drop_reply")
	})
	return err
}

// DropReply is a representation of the C type gdk_drop_reply.
func DropReply(context *DragContext, accepted bool, time uint32) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(context.Native)
	inArgs[1].SetBoolean(accepted)
	inArgs[2].SetUint32(time)

	err := dropReplyFunction_Set()
	if err == nil {
		dropReplyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var errorTrapPopFunction *gi.Function
var errorTrapPopFunction_Once sync.Once

func errorTrapPopFunction_Set() error {
	var err error
	errorTrapPopFunction_Once.Do(func() {
		errorTrapPopFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop")
	})
	return err
}

// ErrorTrapPop is a representation of the C type gdk_error_trap_pop.
func ErrorTrapPop() int32 {

	var ret gi.Argument

	err := errorTrapPopFunction_Set()
	if err == nil {
		ret = errorTrapPopFunction.Invoke(nil, nil)
	}

	retGo := ret.Int32()

	return retGo
}

var errorTrapPopIgnoredFunction *gi.Function
var errorTrapPopIgnoredFunction_Once sync.Once

func errorTrapPopIgnoredFunction_Set() error {
	var err error
	errorTrapPopIgnoredFunction_Once.Do(func() {
		errorTrapPopIgnoredFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_pop_ignored")
	})
	return err
}

// ErrorTrapPopIgnored is a representation of the C type gdk_error_trap_pop_ignored.
func ErrorTrapPopIgnored() {

	err := errorTrapPopIgnoredFunction_Set()
	if err == nil {
		errorTrapPopIgnoredFunction.Invoke(nil, nil)
	}

	return
}

var errorTrapPushFunction *gi.Function
var errorTrapPushFunction_Once sync.Once

func errorTrapPushFunction_Set() error {
	var err error
	errorTrapPushFunction_Once.Do(func() {
		errorTrapPushFunction, err = gi.FunctionInvokerNew("Gdk", "error_trap_push")
	})
	return err
}

// ErrorTrapPush is a representation of the C type gdk_error_trap_push.
func ErrorTrapPush() {

	err := errorTrapPushFunction_Set()
	if err == nil {
		errorTrapPushFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_event_get' : return type 'Event' not supported

// UNSUPPORTED : C value 'gdk_event_handler_set' : parameter 'func' of type 'EventFunc' not supported

// UNSUPPORTED : C value 'gdk_event_peek' : return type 'Event' not supported

var eventRequestMotionsFunction *gi.Function
var eventRequestMotionsFunction_Once sync.Once

func eventRequestMotionsFunction_Set() error {
	var err error
	eventRequestMotionsFunction_Once.Do(func() {
		eventRequestMotionsFunction, err = gi.FunctionInvokerNew("Gdk", "event_request_motions")
	})
	return err
}

// EventRequestMotions is a representation of the C type gdk_event_request_motions.
func EventRequestMotions(event *EventMotion) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(event.Native)

	err := eventRequestMotionsFunction_Set()
	if err == nil {
		eventRequestMotionsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_events_get_angle' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_center' : parameter 'event1' of type 'Event' not supported

// UNSUPPORTED : C value 'gdk_events_get_distance' : parameter 'event1' of type 'Event' not supported

var eventsPendingFunction *gi.Function
var eventsPendingFunction_Once sync.Once

func eventsPendingFunction_Set() error {
	var err error
	eventsPendingFunction_Once.Do(func() {
		eventsPendingFunction, err = gi.FunctionInvokerNew("Gdk", "events_pending")
	})
	return err
}

// EventsPending is a representation of the C type gdk_events_pending.
func EventsPending() bool {

	var ret gi.Argument

	err := eventsPendingFunction_Set()
	if err == nil {
		ret = eventsPendingFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var flushFunction *gi.Function
var flushFunction_Once sync.Once

func flushFunction_Set() error {
	var err error
	flushFunction_Once.Do(func() {
		flushFunction, err = gi.FunctionInvokerNew("Gdk", "flush")
	})
	return err
}

// Flush is a representation of the C type gdk_flush.
func Flush() {

	err := flushFunction_Set()
	if err == nil {
		flushFunction.Invoke(nil, nil)
	}

	return
}

var getDefaultRootWindowFunction *gi.Function
var getDefaultRootWindowFunction_Once sync.Once

func getDefaultRootWindowFunction_Set() error {
	var err error
	getDefaultRootWindowFunction_Once.Do(func() {
		getDefaultRootWindowFunction, err = gi.FunctionInvokerNew("Gdk", "get_default_root_window")
	})
	return err
}

// GetDefaultRootWindow is a representation of the C type gdk_get_default_root_window.
func GetDefaultRootWindow() *Window {

	var ret gi.Argument

	err := getDefaultRootWindowFunction_Set()
	if err == nil {
		ret = getDefaultRootWindowFunction.Invoke(nil, nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var getDisplayFunction *gi.Function
var getDisplayFunction_Once sync.Once

func getDisplayFunction_Set() error {
	var err error
	getDisplayFunction_Once.Do(func() {
		getDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "get_display")
	})
	return err
}

// GetDisplay is a representation of the C type gdk_get_display.
func GetDisplay() string {

	var ret gi.Argument

	err := getDisplayFunction_Set()
	if err == nil {
		ret = getDisplayFunction.Invoke(nil, nil)
	}

	retGo := ret.String(true)

	return retGo
}

var getDisplayArgNameFunction *gi.Function
var getDisplayArgNameFunction_Once sync.Once

func getDisplayArgNameFunction_Set() error {
	var err error
	getDisplayArgNameFunction_Once.Do(func() {
		getDisplayArgNameFunction, err = gi.FunctionInvokerNew("Gdk", "get_display_arg_name")
	})
	return err
}

// GetDisplayArgName is a representation of the C type gdk_get_display_arg_name.
func GetDisplayArgName() string {

	var ret gi.Argument

	err := getDisplayArgNameFunction_Set()
	if err == nil {
		ret = getDisplayArgNameFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getProgramClassFunction *gi.Function
var getProgramClassFunction_Once sync.Once

func getProgramClassFunction_Set() error {
	var err error
	getProgramClassFunction_Once.Do(func() {
		getProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "get_program_class")
	})
	return err
}

// GetProgramClass is a representation of the C type gdk_get_program_class.
func GetProgramClass() string {

	var ret gi.Argument

	err := getProgramClassFunction_Set()
	if err == nil {
		ret = getProgramClassFunction.Invoke(nil, nil)
	}

	retGo := ret.String(false)

	return retGo
}

var getShowEventsFunction *gi.Function
var getShowEventsFunction_Once sync.Once

func getShowEventsFunction_Set() error {
	var err error
	getShowEventsFunction_Once.Do(func() {
		getShowEventsFunction, err = gi.FunctionInvokerNew("Gdk", "get_show_events")
	})
	return err
}

// GetShowEvents is a representation of the C type gdk_get_show_events.
func GetShowEvents() bool {

	var ret gi.Argument

	err := getShowEventsFunction_Set()
	if err == nil {
		ret = getShowEventsFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var glErrorQuarkFunction *gi.Function
var glErrorQuarkFunction_Once sync.Once

func glErrorQuarkFunction_Set() error {
	var err error
	glErrorQuarkFunction_Once.Do(func() {
		glErrorQuarkFunction, err = gi.FunctionInvokerNew("Gdk", "gl_error_quark")
	})
	return err
}

// GlErrorQuark is a representation of the C type gdk_gl_error_quark.
func GlErrorQuark() glib.Quark {

	var ret gi.Argument

	err := glErrorQuarkFunction_Set()
	if err == nil {
		ret = glErrorQuarkFunction.Invoke(nil, nil)
	}

	retGo := glib.Quark(ret.Uint32())

	return retGo
}

// UNSUPPORTED : C value 'gdk_init' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_init_check' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_keyboard_grab' : return type 'GrabStatus' not supported

var keyboardUngrabFunction *gi.Function
var keyboardUngrabFunction_Once sync.Once

func keyboardUngrabFunction_Set() error {
	var err error
	keyboardUngrabFunction_Once.Do(func() {
		keyboardUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "keyboard_ungrab")
	})
	return err
}

// KeyboardUngrab is a representation of the C type gdk_keyboard_ungrab.
func KeyboardUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := keyboardUngrabFunction_Set()
	if err == nil {
		keyboardUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var keyvalConvertCaseFunction *gi.Function
var keyvalConvertCaseFunction_Once sync.Once

func keyvalConvertCaseFunction_Set() error {
	var err error
	keyvalConvertCaseFunction_Once.Do(func() {
		keyvalConvertCaseFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_convert_case")
	})
	return err
}

// KeyvalConvertCase is a representation of the C type gdk_keyval_convert_case.
func KeyvalConvertCase(symbol uint32) (uint32, uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(symbol)

	var outArgs [2]gi.Argument

	err := keyvalConvertCaseFunction_Set()
	if err == nil {
		keyvalConvertCaseFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Uint32()
	out1 := outArgs[1].Uint32()

	return out0, out1
}

var keyvalFromNameFunction *gi.Function
var keyvalFromNameFunction_Once sync.Once

func keyvalFromNameFunction_Set() error {
	var err error
	keyvalFromNameFunction_Once.Do(func() {
		keyvalFromNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_from_name")
	})
	return err
}

// KeyvalFromName is a representation of the C type gdk_keyval_from_name.
func KeyvalFromName(keyvalName string) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(keyvalName)

	var ret gi.Argument

	err := keyvalFromNameFunction_Set()
	if err == nil {
		ret = keyvalFromNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalIsLowerFunction *gi.Function
var keyvalIsLowerFunction_Once sync.Once

func keyvalIsLowerFunction_Set() error {
	var err error
	keyvalIsLowerFunction_Once.Do(func() {
		keyvalIsLowerFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_is_lower")
	})
	return err
}

// KeyvalIsLower is a representation of the C type gdk_keyval_is_lower.
func KeyvalIsLower(keyval uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalIsLowerFunction_Set()
	if err == nil {
		ret = keyvalIsLowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keyvalIsUpperFunction *gi.Function
var keyvalIsUpperFunction_Once sync.Once

func keyvalIsUpperFunction_Set() error {
	var err error
	keyvalIsUpperFunction_Once.Do(func() {
		keyvalIsUpperFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_is_upper")
	})
	return err
}

// KeyvalIsUpper is a representation of the C type gdk_keyval_is_upper.
func KeyvalIsUpper(keyval uint32) bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalIsUpperFunction_Set()
	if err == nil {
		ret = keyvalIsUpperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var keyvalNameFunction *gi.Function
var keyvalNameFunction_Once sync.Once

func keyvalNameFunction_Set() error {
	var err error
	keyvalNameFunction_Once.Do(func() {
		keyvalNameFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_name")
	})
	return err
}

// KeyvalName is a representation of the C type gdk_keyval_name.
func KeyvalName(keyval uint32) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalNameFunction_Set()
	if err == nil {
		ret = keyvalNameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var keyvalToLowerFunction *gi.Function
var keyvalToLowerFunction_Once sync.Once

func keyvalToLowerFunction_Set() error {
	var err error
	keyvalToLowerFunction_Once.Do(func() {
		keyvalToLowerFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_lower")
	})
	return err
}

// KeyvalToLower is a representation of the C type gdk_keyval_to_lower.
func KeyvalToLower(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToLowerFunction_Set()
	if err == nil {
		ret = keyvalToLowerFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUnicodeFunction *gi.Function
var keyvalToUnicodeFunction_Once sync.Once

func keyvalToUnicodeFunction_Set() error {
	var err error
	keyvalToUnicodeFunction_Once.Do(func() {
		keyvalToUnicodeFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_unicode")
	})
	return err
}

// KeyvalToUnicode is a representation of the C type gdk_keyval_to_unicode.
func KeyvalToUnicode(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUnicodeFunction_Set()
	if err == nil {
		ret = keyvalToUnicodeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var keyvalToUpperFunction *gi.Function
var keyvalToUpperFunction_Once sync.Once

func keyvalToUpperFunction_Set() error {
	var err error
	keyvalToUpperFunction_Once.Do(func() {
		keyvalToUpperFunction, err = gi.FunctionInvokerNew("Gdk", "keyval_to_upper")
	})
	return err
}

// KeyvalToUpper is a representation of the C type gdk_keyval_to_upper.
func KeyvalToUpper(keyval uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(keyval)

	var ret gi.Argument

	err := keyvalToUpperFunction_Set()
	if err == nil {
		ret = keyvalToUpperFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

// UNSUPPORTED : C value 'gdk_list_visuals' : return type 'GLib.List' not supported

var notifyStartupCompleteFunction *gi.Function
var notifyStartupCompleteFunction_Once sync.Once

func notifyStartupCompleteFunction_Set() error {
	var err error
	notifyStartupCompleteFunction_Once.Do(func() {
		notifyStartupCompleteFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete")
	})
	return err
}

// NotifyStartupComplete is a representation of the C type gdk_notify_startup_complete.
func NotifyStartupComplete() {

	err := notifyStartupCompleteFunction_Set()
	if err == nil {
		notifyStartupCompleteFunction.Invoke(nil, nil)
	}

	return
}

var notifyStartupCompleteWithIdFunction *gi.Function
var notifyStartupCompleteWithIdFunction_Once sync.Once

func notifyStartupCompleteWithIdFunction_Set() error {
	var err error
	notifyStartupCompleteWithIdFunction_Once.Do(func() {
		notifyStartupCompleteWithIdFunction, err = gi.FunctionInvokerNew("Gdk", "notify_startup_complete_with_id")
	})
	return err
}

// NotifyStartupCompleteWithId is a representation of the C type gdk_notify_startup_complete_with_id.
func NotifyStartupCompleteWithId(startupId string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(startupId)

	err := notifyStartupCompleteWithIdFunction_Set()
	if err == nil {
		notifyStartupCompleteWithIdFunction.Invoke(inArgs[:], nil)
	}

	return
}

var offscreenWindowGetEmbedderFunction *gi.Function
var offscreenWindowGetEmbedderFunction_Once sync.Once

func offscreenWindowGetEmbedderFunction_Set() error {
	var err error
	offscreenWindowGetEmbedderFunction_Once.Do(func() {
		offscreenWindowGetEmbedderFunction, err = gi.FunctionInvokerNew("Gdk", "offscreen_window_get_embedder")
	})
	return err
}

// OffscreenWindowGetEmbedder is a representation of the C type gdk_offscreen_window_get_embedder.
func OffscreenWindowGetEmbedder(window *Window) *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native)

	var ret gi.Argument

	err := offscreenWindowGetEmbedderFunction_Set()
	if err == nil {
		ret = offscreenWindowGetEmbedderFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

// UNSUPPORTED : C value 'gdk_offscreen_window_get_surface' : return type 'cairo.Surface' not supported

var offscreenWindowSetEmbedderFunction *gi.Function
var offscreenWindowSetEmbedderFunction_Once sync.Once

func offscreenWindowSetEmbedderFunction_Set() error {
	var err error
	offscreenWindowSetEmbedderFunction_Once.Do(func() {
		offscreenWindowSetEmbedderFunction, err = gi.FunctionInvokerNew("Gdk", "offscreen_window_set_embedder")
	})
	return err
}

// OffscreenWindowSetEmbedder is a representation of the C type gdk_offscreen_window_set_embedder.
func OffscreenWindowSetEmbedder(window *Window, embedder *Window) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(window.Native)
	inArgs[1].SetPointer(embedder.Native)

	err := offscreenWindowSetEmbedderFunction_Set()
	if err == nil {
		offscreenWindowSetEmbedderFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_pango_context_get' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_display' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_context_get_for_screen' : return type 'Pango.Context' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_get_clip_region' : parameter 'layout' of type 'Pango.Layout' not supported

// UNSUPPORTED : C value 'gdk_pango_layout_line_get_clip_region' : parameter 'line' of type 'Pango.LayoutLine' not supported

// UNSUPPORTED : C value 'gdk_parse_args' : parameter 'argv' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_surface' : parameter 'surface' of type 'cairo.Surface' not supported

// UNSUPPORTED : C value 'gdk_pixbuf_get_from_window' : return type 'GdkPixbuf.Pixbuf' not supported

// UNSUPPORTED : C value 'gdk_pointer_grab' : parameter 'event_mask' of type 'EventMask' not supported

var pointerIsGrabbedFunction *gi.Function
var pointerIsGrabbedFunction_Once sync.Once

func pointerIsGrabbedFunction_Set() error {
	var err error
	pointerIsGrabbedFunction_Once.Do(func() {
		pointerIsGrabbedFunction, err = gi.FunctionInvokerNew("Gdk", "pointer_is_grabbed")
	})
	return err
}

// PointerIsGrabbed is a representation of the C type gdk_pointer_is_grabbed.
func PointerIsGrabbed() bool {

	var ret gi.Argument

	err := pointerIsGrabbedFunction_Set()
	if err == nil {
		ret = pointerIsGrabbedFunction.Invoke(nil, nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var pointerUngrabFunction *gi.Function
var pointerUngrabFunction_Once sync.Once

func pointerUngrabFunction_Set() error {
	var err error
	pointerUngrabFunction_Once.Do(func() {
		pointerUngrabFunction, err = gi.FunctionInvokerNew("Gdk", "pointer_ungrab")
	})
	return err
}

// PointerUngrab is a representation of the C type gdk_pointer_ungrab.
func PointerUngrab(time uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(time)

	err := pointerUngrabFunction_Set()
	if err == nil {
		pointerUngrabFunction.Invoke(inArgs[:], nil)
	}

	return
}

var preParseLibgtkOnlyFunction *gi.Function
var preParseLibgtkOnlyFunction_Once sync.Once

func preParseLibgtkOnlyFunction_Set() error {
	var err error
	preParseLibgtkOnlyFunction_Once.Do(func() {
		preParseLibgtkOnlyFunction, err = gi.FunctionInvokerNew("Gdk", "pre_parse_libgtk_only")
	})
	return err
}

// PreParseLibgtkOnly is a representation of the C type gdk_pre_parse_libgtk_only.
func PreParseLibgtkOnly() {

	err := preParseLibgtkOnlyFunction_Set()
	if err == nil {
		preParseLibgtkOnlyFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_property_change' : parameter 'mode' of type 'PropMode' not supported

var propertyDeleteFunction *gi.Function
var propertyDeleteFunction_Once sync.Once

func propertyDeleteFunction_Set() error {
	var err error
	propertyDeleteFunction_Once.Do(func() {
		propertyDeleteFunction, err = gi.FunctionInvokerNew("Gdk", "property_delete")
	})
	return err
}

// PropertyDelete is a representation of the C type gdk_property_delete.
func PropertyDelete(window *Window, property *Atom) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(window.Native)
	inArgs[1].SetPointer(property.Native)

	err := propertyDeleteFunction_Set()
	if err == nil {
		propertyDeleteFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_property_get' : parameter 'data' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_query_depths' : parameter 'depths' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_query_visual_types' : parameter 'visual_types' of type 'nil' not supported

var selectionConvertFunction *gi.Function
var selectionConvertFunction_Once sync.Once

func selectionConvertFunction_Set() error {
	var err error
	selectionConvertFunction_Once.Do(func() {
		selectionConvertFunction, err = gi.FunctionInvokerNew("Gdk", "selection_convert")
	})
	return err
}

// SelectionConvert is a representation of the C type gdk_selection_convert.
func SelectionConvert(requestor *Window, selection *Atom, target *Atom, time uint32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(requestor.Native)
	inArgs[1].SetPointer(selection.Native)
	inArgs[2].SetPointer(target.Native)
	inArgs[3].SetUint32(time)

	err := selectionConvertFunction_Set()
	if err == nil {
		selectionConvertFunction.Invoke(inArgs[:], nil)
	}

	return
}

var selectionOwnerGetFunction *gi.Function
var selectionOwnerGetFunction_Once sync.Once

func selectionOwnerGetFunction_Set() error {
	var err error
	selectionOwnerGetFunction_Once.Do(func() {
		selectionOwnerGetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_get")
	})
	return err
}

// SelectionOwnerGet is a representation of the C type gdk_selection_owner_get.
func SelectionOwnerGet(selection *Atom) *Window {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(selection.Native)

	var ret gi.Argument

	err := selectionOwnerGetFunction_Set()
	if err == nil {
		ret = selectionOwnerGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var selectionOwnerGetForDisplayFunction *gi.Function
var selectionOwnerGetForDisplayFunction_Once sync.Once

func selectionOwnerGetForDisplayFunction_Set() error {
	var err error
	selectionOwnerGetForDisplayFunction_Once.Do(func() {
		selectionOwnerGetForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_get_for_display")
	})
	return err
}

// SelectionOwnerGetForDisplay is a representation of the C type gdk_selection_owner_get_for_display.
func SelectionOwnerGetForDisplay(display *Display, selection *Atom) *Window {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(display.Native)
	inArgs[1].SetPointer(selection.Native)

	var ret gi.Argument

	err := selectionOwnerGetForDisplayFunction_Set()
	if err == nil {
		ret = selectionOwnerGetForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := &Window{}
	retGo.Native = ret.Pointer()

	return retGo
}

var selectionOwnerSetFunction *gi.Function
var selectionOwnerSetFunction_Once sync.Once

func selectionOwnerSetFunction_Set() error {
	var err error
	selectionOwnerSetFunction_Once.Do(func() {
		selectionOwnerSetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_set")
	})
	return err
}

// SelectionOwnerSet is a representation of the C type gdk_selection_owner_set.
func SelectionOwnerSet(owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(owner.Native)
	inArgs[1].SetPointer(selection.Native)
	inArgs[2].SetUint32(time)
	inArgs[3].SetBoolean(sendEvent)

	var ret gi.Argument

	err := selectionOwnerSetFunction_Set()
	if err == nil {
		ret = selectionOwnerSetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionOwnerSetForDisplayFunction *gi.Function
var selectionOwnerSetForDisplayFunction_Once sync.Once

func selectionOwnerSetForDisplayFunction_Set() error {
	var err error
	selectionOwnerSetForDisplayFunction_Once.Do(func() {
		selectionOwnerSetForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_owner_set_for_display")
	})
	return err
}

// SelectionOwnerSetForDisplay is a representation of the C type gdk_selection_owner_set_for_display.
func SelectionOwnerSetForDisplay(display *Display, owner *Window, selection *Atom, time uint32, sendEvent bool) bool {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(display.Native)
	inArgs[1].SetPointer(owner.Native)
	inArgs[2].SetPointer(selection.Native)
	inArgs[3].SetUint32(time)
	inArgs[4].SetBoolean(sendEvent)

	var ret gi.Argument

	err := selectionOwnerSetForDisplayFunction_Set()
	if err == nil {
		ret = selectionOwnerSetForDisplayFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var selectionPropertyGetFunction *gi.Function
var selectionPropertyGetFunction_Once sync.Once

func selectionPropertyGetFunction_Set() error {
	var err error
	selectionPropertyGetFunction_Once.Do(func() {
		selectionPropertyGetFunction, err = gi.FunctionInvokerNew("Gdk", "selection_property_get")
	})
	return err
}

// SelectionPropertyGet is a representation of the C type gdk_selection_property_get.
func SelectionPropertyGet(requestor *Window, data uint8, propType *Atom, propFormat int32) int32 {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(requestor.Native)
	inArgs[1].SetUint8(data)
	inArgs[2].SetPointer(propType.Native)
	inArgs[3].SetInt32(propFormat)

	var ret gi.Argument

	err := selectionPropertyGetFunction_Set()
	if err == nil {
		ret = selectionPropertyGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var selectionSendNotifyFunction *gi.Function
var selectionSendNotifyFunction_Once sync.Once

func selectionSendNotifyFunction_Set() error {
	var err error
	selectionSendNotifyFunction_Once.Do(func() {
		selectionSendNotifyFunction, err = gi.FunctionInvokerNew("Gdk", "selection_send_notify")
	})
	return err
}

// SelectionSendNotify is a representation of the C type gdk_selection_send_notify.
func SelectionSendNotify(requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(requestor.Native)
	inArgs[1].SetPointer(selection.Native)
	inArgs[2].SetPointer(target.Native)
	inArgs[3].SetPointer(property.Native)
	inArgs[4].SetUint32(time)

	err := selectionSendNotifyFunction_Set()
	if err == nil {
		selectionSendNotifyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var selectionSendNotifyForDisplayFunction *gi.Function
var selectionSendNotifyForDisplayFunction_Once sync.Once

func selectionSendNotifyForDisplayFunction_Set() error {
	var err error
	selectionSendNotifyForDisplayFunction_Once.Do(func() {
		selectionSendNotifyForDisplayFunction, err = gi.FunctionInvokerNew("Gdk", "selection_send_notify_for_display")
	})
	return err
}

// SelectionSendNotifyForDisplay is a representation of the C type gdk_selection_send_notify_for_display.
func SelectionSendNotifyForDisplay(display *Display, requestor *Window, selection *Atom, target *Atom, property *Atom, time uint32) {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(display.Native)
	inArgs[1].SetPointer(requestor.Native)
	inArgs[2].SetPointer(selection.Native)
	inArgs[3].SetPointer(target.Native)
	inArgs[4].SetPointer(property.Native)
	inArgs[5].SetUint32(time)

	err := selectionSendNotifyForDisplayFunction_Set()
	if err == nil {
		selectionSendNotifyForDisplayFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setAllowedBackendsFunction *gi.Function
var setAllowedBackendsFunction_Once sync.Once

func setAllowedBackendsFunction_Set() error {
	var err error
	setAllowedBackendsFunction_Once.Do(func() {
		setAllowedBackendsFunction, err = gi.FunctionInvokerNew("Gdk", "set_allowed_backends")
	})
	return err
}

// SetAllowedBackends is a representation of the C type gdk_set_allowed_backends.
func SetAllowedBackends(backends string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(backends)

	err := setAllowedBackendsFunction_Set()
	if err == nil {
		setAllowedBackendsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setDoubleClickTimeFunction *gi.Function
var setDoubleClickTimeFunction_Once sync.Once

func setDoubleClickTimeFunction_Set() error {
	var err error
	setDoubleClickTimeFunction_Once.Do(func() {
		setDoubleClickTimeFunction, err = gi.FunctionInvokerNew("Gdk", "set_double_click_time")
	})
	return err
}

// SetDoubleClickTime is a representation of the C type gdk_set_double_click_time.
func SetDoubleClickTime(msec uint32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(msec)

	err := setDoubleClickTimeFunction_Set()
	if err == nil {
		setDoubleClickTimeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setProgramClassFunction *gi.Function
var setProgramClassFunction_Once sync.Once

func setProgramClassFunction_Set() error {
	var err error
	setProgramClassFunction_Once.Do(func() {
		setProgramClassFunction, err = gi.FunctionInvokerNew("Gdk", "set_program_class")
	})
	return err
}

// SetProgramClass is a representation of the C type gdk_set_program_class.
func SetProgramClass(programClass string) {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(programClass)

	err := setProgramClassFunction_Set()
	if err == nil {
		setProgramClassFunction.Invoke(inArgs[:], nil)
	}

	return
}

var setShowEventsFunction *gi.Function
var setShowEventsFunction_Once sync.Once

func setShowEventsFunction_Set() error {
	var err error
	setShowEventsFunction_Once.Do(func() {
		setShowEventsFunction, err = gi.FunctionInvokerNew("Gdk", "set_show_events")
	})
	return err
}

// SetShowEvents is a representation of the C type gdk_set_show_events.
func SetShowEvents(showEvents bool) {
	var inArgs [1]gi.Argument
	inArgs[0].SetBoolean(showEvents)

	err := setShowEventsFunction_Set()
	if err == nil {
		setShowEventsFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_setting_get' : parameter 'value' of type 'GObject.Value' not supported

// UNSUPPORTED : C value 'gdk_synthesize_window_state' : parameter 'unset_flags' of type 'WindowState' not supported

var testRenderSyncFunction *gi.Function
var testRenderSyncFunction_Once sync.Once

func testRenderSyncFunction_Set() error {
	var err error
	testRenderSyncFunction_Once.Do(func() {
		testRenderSyncFunction, err = gi.FunctionInvokerNew("Gdk", "test_render_sync")
	})
	return err
}

// TestRenderSync is a representation of the C type gdk_test_render_sync.
func TestRenderSync(window *Window) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(window.Native)

	err := testRenderSyncFunction_Set()
	if err == nil {
		testRenderSyncFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_test_simulate_button' : parameter 'modifiers' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_test_simulate_key' : parameter 'modifiers' of type 'ModifierType' not supported

// UNSUPPORTED : C value 'gdk_text_property_to_utf8_list_for_display' : parameter 'text' of type 'nil' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_idle_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds' : parameter 'function' of type 'GLib.SourceFunc' not supported

// UNSUPPORTED : C value 'gdk_threads_add_timeout_seconds_full' : parameter 'function' of type 'GLib.SourceFunc' not supported

var threadsEnterFunction *gi.Function
var threadsEnterFunction_Once sync.Once

func threadsEnterFunction_Set() error {
	var err error
	threadsEnterFunction_Once.Do(func() {
		threadsEnterFunction, err = gi.FunctionInvokerNew("Gdk", "threads_enter")
	})
	return err
}

// ThreadsEnter is a representation of the C type gdk_threads_enter.
func ThreadsEnter() {

	err := threadsEnterFunction_Set()
	if err == nil {
		threadsEnterFunction.Invoke(nil, nil)
	}

	return
}

var threadsInitFunction *gi.Function
var threadsInitFunction_Once sync.Once

func threadsInitFunction_Set() error {
	var err error
	threadsInitFunction_Once.Do(func() {
		threadsInitFunction, err = gi.FunctionInvokerNew("Gdk", "threads_init")
	})
	return err
}

// ThreadsInit is a representation of the C type gdk_threads_init.
func ThreadsInit() {

	err := threadsInitFunction_Set()
	if err == nil {
		threadsInitFunction.Invoke(nil, nil)
	}

	return
}

var threadsLeaveFunction *gi.Function
var threadsLeaveFunction_Once sync.Once

func threadsLeaveFunction_Set() error {
	var err error
	threadsLeaveFunction_Once.Do(func() {
		threadsLeaveFunction, err = gi.FunctionInvokerNew("Gdk", "threads_leave")
	})
	return err
}

// ThreadsLeave is a representation of the C type gdk_threads_leave.
func ThreadsLeave() {

	err := threadsLeaveFunction_Set()
	if err == nil {
		threadsLeaveFunction.Invoke(nil, nil)
	}

	return
}

// UNSUPPORTED : C value 'gdk_threads_set_lock_functions' : parameter 'enter_fn' of type 'GObject.Callback' not supported

var unicodeToKeyvalFunction *gi.Function
var unicodeToKeyvalFunction_Once sync.Once

func unicodeToKeyvalFunction_Set() error {
	var err error
	unicodeToKeyvalFunction_Once.Do(func() {
		unicodeToKeyvalFunction, err = gi.FunctionInvokerNew("Gdk", "unicode_to_keyval")
	})
	return err
}

// UnicodeToKeyval is a representation of the C type gdk_unicode_to_keyval.
func UnicodeToKeyval(wc uint32) uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetUint32(wc)

	var ret gi.Argument

	err := unicodeToKeyvalFunction_Set()
	if err == nil {
		ret = unicodeToKeyvalFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var utf8ToStringTargetFunction *gi.Function
var utf8ToStringTargetFunction_Once sync.Once

func utf8ToStringTargetFunction_Set() error {
	var err error
	utf8ToStringTargetFunction_Once.Do(func() {
		utf8ToStringTargetFunction, err = gi.FunctionInvokerNew("Gdk", "utf8_to_string_target")
	})
	return err
}

// Utf8ToStringTarget is a representation of the C type gdk_utf8_to_string_target.
func Utf8ToStringTarget(str string) string {
	var inArgs [1]gi.Argument
	inArgs[0].SetString(str)

	var ret gi.Argument

	err := utf8ToStringTargetFunction_Set()
	if err == nil {
		ret = utf8ToStringTargetFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}
