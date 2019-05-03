// Code generated - DO NOT EDIT.
// +build gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void combobox_moveActiveHandler(GObject *, GtkScrollType, gpointer);

	static gulong ComboBox_signal_connect_move_active(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-active", G_CALLBACK(combobox_moveActiveHandler), data);
	}

*/
/*

	gboolean combobox_popdownHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(combobox_popdownHandler), data);
	}

*/
/*

	void combobox_popupHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(combobox_popupHandler), data);
	}

*/
/*

	gboolean entrycompletion_cursorOnMatchHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gpointer);

	static gulong EntryCompletion_signal_connect_cursor_on_match(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cursor-on-match", G_CALLBACK(entrycompletion_cursorOnMatchHandler), data);
	}

*/
/*

	void filechooserbutton_fileSetHandler(GObject *, gpointer);

	static gulong FileChooserButton_signal_connect_file_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "file-set", G_CALLBACK(filechooserbutton_fileSetHandler), data);
	}

*/
/*

	gboolean menushell_moveSelectedHandler(GObject *, gint, gpointer);

	static gulong MenuShell_signal_connect_move_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-selected", G_CALLBACK(menushell_moveSelectedHandler), data);
	}

*/
/*

	GtkNotebook * notebook_createWindowHandler(GObject *, GtkWidget *, gint, gint, gpointer);

	static gulong Notebook_signal_connect_create_window(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-window", G_CALLBACK(notebook_createWindowHandler), data);
	}

*/
/*

	void scalebutton_popdownHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popdown(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popdown", G_CALLBACK(scalebutton_popdownHandler), data);
	}

*/
/*

	void scalebutton_popupHandler(GObject *, gpointer);

	static gulong ScaleButton_signal_connect_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup", G_CALLBACK(scalebutton_popupHandler), data);
	}

*/
/*

	void scalebutton_valueChangedHandler(GObject *, gdouble, gpointer);

	static gulong ScaleButton_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", G_CALLBACK(scalebutton_valueChangedHandler), data);
	}

*/
/*

	gboolean widget_dragFailedHandler(GObject *, GdkDragContext *, GtkDragResult, gpointer);

	static gulong Widget_signal_connect_drag_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-failed", G_CALLBACK(widget_dragFailedHandler), data);
	}

*/
/*

	gboolean widget_keynavFailedHandler(GObject *, GtkDirectionType, gpointer);

	static gulong Widget_signal_connect_keynav_failed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keynav-failed", G_CALLBACK(widget_keynavFailedHandler), data);
	}

*/
/*

	gboolean widget_queryTooltipHandler(GObject *, gint, gint, gboolean, GtkTooltip *, gpointer);

	static gulong Widget_signal_connect_query_tooltip(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "query-tooltip", G_CALLBACK(widget_queryTooltipHandler), data);
	}

*/
import "C"

// GetProgramName is a wrapper around the C function gtk_about_dialog_get_program_name.
func (recv *AboutDialog) GetProgramName() string {
	retC := C.gtk_about_dialog_get_program_name((*C.GtkAboutDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetProgramName is a wrapper around the C function gtk_about_dialog_set_program_name.
func (recv *AboutDialog) SetProgramName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_about_dialog_set_program_name((*C.GtkAboutDialog)(recv.native), c_name)

	return
}

// CreateMenu is a wrapper around the C function gtk_action_create_menu.
func (recv *Action) CreateMenu() *Widget {
	retC := C.gtk_action_create_menu((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BuilderNew is a wrapper around the C function gtk_builder_new.
func BuilderNew() *Builder {
	retC := C.gtk_builder_new()
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddFromFile is a wrapper around the C function gtk_builder_add_from_file.
func (recv *Builder) AddFromFile(filename string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_file((*C.GtkBuilder)(recv.native), c_filename, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddFromString is a wrapper around the C function gtk_builder_add_from_string.
func (recv *Builder) AddFromString(buffer string, length uint64) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gsize)(length)

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_from_string((*C.GtkBuilder)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : gtk_builder_connect_signals : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// Unsupported : gtk_builder_connect_signals_full : unsupported parameter func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param func

// GetObject is a wrapper around the C function gtk_builder_get_object.
func (recv *Builder) GetObject(name string) *gobject.Object {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_builder_get_object((*C.GtkBuilder)(recv.native), c_name)
	var retGo (*gobject.Object)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gobject.ObjectNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetObjects is a wrapper around the C function gtk_builder_get_objects.
func (recv *Builder) GetObjects() *glib.SList {
	retC := C.gtk_builder_get_objects((*C.GtkBuilder)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTranslationDomain is a wrapper around the C function gtk_builder_get_translation_domain.
func (recv *Builder) GetTranslationDomain() string {
	retC := C.gtk_builder_get_translation_domain((*C.GtkBuilder)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTypeFromName is a wrapper around the C function gtk_builder_get_type_from_name.
func (recv *Builder) GetTypeFromName(typeName string) gobject.Type {
	c_type_name := C.CString(typeName)
	defer C.free(unsafe.Pointer(c_type_name))

	retC := C.gtk_builder_get_type_from_name((*C.GtkBuilder)(recv.native), c_type_name)
	retGo := (gobject.Type)(retC)

	return retGo
}

// SetTranslationDomain is a wrapper around the C function gtk_builder_set_translation_domain.
func (recv *Builder) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_builder_set_translation_domain((*C.GtkBuilder)(recv.native), c_domain)

	return
}

// ValueFromString is a wrapper around the C function gtk_builder_value_from_string.
func (recv *Builder) ValueFromString(pspec *gobject.ParamSpec, string_ string) (bool, *gobject.Value, error) {
	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	var c_value C.GValue

	var cThrowableError *C.GError

	retC := C.gtk_builder_value_from_string((*C.GtkBuilder)(recv.native), c_pspec, c_string, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value, goError
}

// ValueFromStringType is a wrapper around the C function gtk_builder_value_from_string_type.
func (recv *Builder) ValueFromStringType(type_ gobject.Type, string_ string) (bool, *gobject.Value, error) {
	c_type := (C.GType)(type_)

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	var c_value C.GValue

	var cThrowableError *C.GError

	retC := C.gtk_builder_value_from_string_type((*C.GtkBuilder)(recv.native), c_type, c_string, &c_value, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value, goError
}

type signalComboBoxMoveActiveDetail struct {
	callback  ComboBoxSignalMoveActiveCallback
	handlerID C.gulong
}

var signalComboBoxMoveActiveId int
var signalComboBoxMoveActiveMap = make(map[int]signalComboBoxMoveActiveDetail)
var signalComboBoxMoveActiveLock sync.RWMutex

// ComboBoxSignalMoveActiveCallback is a callback function for a 'move-active' signal emitted from a ComboBox.
type ComboBoxSignalMoveActiveCallback func(scrollType ScrollType)

/*
ConnectMoveActive connects the callback to the 'move-active' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectMoveActive to remove it.
*/
func (recv *ComboBox) ConnectMoveActive(callback ComboBoxSignalMoveActiveCallback) int {
	signalComboBoxMoveActiveLock.Lock()
	defer signalComboBoxMoveActiveLock.Unlock()

	signalComboBoxMoveActiveId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_move_active(instance, C.gpointer(uintptr(signalComboBoxMoveActiveId)))

	detail := signalComboBoxMoveActiveDetail{callback, handlerID}
	signalComboBoxMoveActiveMap[signalComboBoxMoveActiveId] = detail

	return signalComboBoxMoveActiveId
}

/*
DisconnectMoveActive disconnects a callback from the 'move-active' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectMoveActive.
*/
func (recv *ComboBox) DisconnectMoveActive(connectionID int) {
	signalComboBoxMoveActiveLock.Lock()
	defer signalComboBoxMoveActiveLock.Unlock()

	detail, exists := signalComboBoxMoveActiveMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxMoveActiveMap, connectionID)
}

//export combobox_moveActiveHandler
func combobox_moveActiveHandler(_ *C.GObject, c_scroll_type C.GtkScrollType, data C.gpointer) {
	signalComboBoxMoveActiveLock.RLock()
	defer signalComboBoxMoveActiveLock.RUnlock()

	scrollType := ScrollType(c_scroll_type)

	index := int(uintptr(data))
	callback := signalComboBoxMoveActiveMap[index].callback
	callback(scrollType)
}

type signalComboBoxPopdownDetail struct {
	callback  ComboBoxSignalPopdownCallback
	handlerID C.gulong
}

var signalComboBoxPopdownId int
var signalComboBoxPopdownMap = make(map[int]signalComboBoxPopdownDetail)
var signalComboBoxPopdownLock sync.RWMutex

// ComboBoxSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ComboBox.
type ComboBoxSignalPopdownCallback func() bool

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ComboBox) ConnectPopdown(callback ComboBoxSignalPopdownCallback) int {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	signalComboBoxPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popdown(instance, C.gpointer(uintptr(signalComboBoxPopdownId)))

	detail := signalComboBoxPopdownDetail{callback, handlerID}
	signalComboBoxPopdownMap[signalComboBoxPopdownId] = detail

	return signalComboBoxPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ComboBox) DisconnectPopdown(connectionID int) {
	signalComboBoxPopdownLock.Lock()
	defer signalComboBoxPopdownLock.Unlock()

	detail, exists := signalComboBoxPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopdownMap, connectionID)
}

//export combobox_popdownHandler
func combobox_popdownHandler(_ *C.GObject, data C.gpointer) C.gboolean {
	signalComboBoxPopdownLock.RLock()
	defer signalComboBoxPopdownLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxPopdownMap[index].callback
	retGo := callback()
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalComboBoxPopupDetail struct {
	callback  ComboBoxSignalPopupCallback
	handlerID C.gulong
}

var signalComboBoxPopupId int
var signalComboBoxPopupMap = make(map[int]signalComboBoxPopupDetail)
var signalComboBoxPopupLock sync.RWMutex

// ComboBoxSignalPopupCallback is a callback function for a 'popup' signal emitted from a ComboBox.
type ComboBoxSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ComboBox) ConnectPopup(callback ComboBoxSignalPopupCallback) int {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	signalComboBoxPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_popup(instance, C.gpointer(uintptr(signalComboBoxPopupId)))

	detail := signalComboBoxPopupDetail{callback, handlerID}
	signalComboBoxPopupMap[signalComboBoxPopupId] = detail

	return signalComboBoxPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ComboBox) DisconnectPopup(connectionID int) {
	signalComboBoxPopupLock.Lock()
	defer signalComboBoxPopupLock.Unlock()

	detail, exists := signalComboBoxPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxPopupMap, connectionID)
}

//export combobox_popupHandler
func combobox_popupHandler(_ *C.GObject, data C.gpointer) {
	signalComboBoxPopupLock.RLock()
	defer signalComboBoxPopupLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxPopupMap[index].callback
	callback()
}

// GetCursorHadjustment is a wrapper around the C function gtk_entry_get_cursor_hadjustment.
func (recv *Entry) GetCursorHadjustment() *Adjustment {
	retC := C.gtk_entry_get_cursor_hadjustment((*C.GtkEntry)(recv.native))
	var retGo (*Adjustment)
	if retC == nil {
		retGo = nil
	} else {
		retGo = AdjustmentNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetCursorHadjustment is a wrapper around the C function gtk_entry_set_cursor_hadjustment.
func (recv *Entry) SetCursorHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_entry_set_cursor_hadjustment((*C.GtkEntry)(recv.native), c_adjustment)

	return
}

type signalEntryCompletionCursorOnMatchDetail struct {
	callback  EntryCompletionSignalCursorOnMatchCallback
	handlerID C.gulong
}

var signalEntryCompletionCursorOnMatchId int
var signalEntryCompletionCursorOnMatchMap = make(map[int]signalEntryCompletionCursorOnMatchDetail)
var signalEntryCompletionCursorOnMatchLock sync.RWMutex

// EntryCompletionSignalCursorOnMatchCallback is a callback function for a 'cursor-on-match' signal emitted from a EntryCompletion.
type EntryCompletionSignalCursorOnMatchCallback func(model *TreeModel, iter *TreeIter) bool

/*
ConnectCursorOnMatch connects the callback to the 'cursor-on-match' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectCursorOnMatch to remove it.
*/
func (recv *EntryCompletion) ConnectCursorOnMatch(callback EntryCompletionSignalCursorOnMatchCallback) int {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	signalEntryCompletionCursorOnMatchId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_cursor_on_match(instance, C.gpointer(uintptr(signalEntryCompletionCursorOnMatchId)))

	detail := signalEntryCompletionCursorOnMatchDetail{callback, handlerID}
	signalEntryCompletionCursorOnMatchMap[signalEntryCompletionCursorOnMatchId] = detail

	return signalEntryCompletionCursorOnMatchId
}

/*
DisconnectCursorOnMatch disconnects a callback from the 'cursor-on-match' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectCursorOnMatch.
*/
func (recv *EntryCompletion) DisconnectCursorOnMatch(connectionID int) {
	signalEntryCompletionCursorOnMatchLock.Lock()
	defer signalEntryCompletionCursorOnMatchLock.Unlock()

	detail, exists := signalEntryCompletionCursorOnMatchMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionCursorOnMatchMap, connectionID)
}

//export entrycompletion_cursorOnMatchHandler
func entrycompletion_cursorOnMatchHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	signalEntryCompletionCursorOnMatchLock.RLock()
	defer signalEntryCompletionCursorOnMatchLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionCursorOnMatchMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// GetCompletionPrefix is a wrapper around the C function gtk_entry_completion_get_completion_prefix.
func (recv *EntryCompletion) GetCompletionPrefix() string {
	retC := C.gtk_entry_completion_get_completion_prefix((*C.GtkEntryCompletion)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetInlineSelection is a wrapper around the C function gtk_entry_completion_get_inline_selection.
func (recv *EntryCompletion) GetInlineSelection() bool {
	retC := C.gtk_entry_completion_get_inline_selection((*C.GtkEntryCompletion)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetInlineSelection is a wrapper around the C function gtk_entry_completion_set_inline_selection.
func (recv *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	c_inline_selection :=
		boolToGboolean(inlineSelection)

	C.gtk_entry_completion_set_inline_selection((*C.GtkEntryCompletion)(recv.native), c_inline_selection)

	return
}

type signalFileChooserButtonFileSetDetail struct {
	callback  FileChooserButtonSignalFileSetCallback
	handlerID C.gulong
}

var signalFileChooserButtonFileSetId int
var signalFileChooserButtonFileSetMap = make(map[int]signalFileChooserButtonFileSetDetail)
var signalFileChooserButtonFileSetLock sync.RWMutex

// FileChooserButtonSignalFileSetCallback is a callback function for a 'file-set' signal emitted from a FileChooserButton.
type FileChooserButtonSignalFileSetCallback func()

/*
ConnectFileSet connects the callback to the 'file-set' signal for the FileChooserButton.

The returned value represents the connection, and may be passed to DisconnectFileSet to remove it.
*/
func (recv *FileChooserButton) ConnectFileSet(callback FileChooserButtonSignalFileSetCallback) int {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	signalFileChooserButtonFileSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.FileChooserButton_signal_connect_file_set(instance, C.gpointer(uintptr(signalFileChooserButtonFileSetId)))

	detail := signalFileChooserButtonFileSetDetail{callback, handlerID}
	signalFileChooserButtonFileSetMap[signalFileChooserButtonFileSetId] = detail

	return signalFileChooserButtonFileSetId
}

/*
DisconnectFileSet disconnects a callback from the 'file-set' signal for the FileChooserButton.

The connectionID should be a value returned from a call to ConnectFileSet.
*/
func (recv *FileChooserButton) DisconnectFileSet(connectionID int) {
	signalFileChooserButtonFileSetLock.Lock()
	defer signalFileChooserButtonFileSetLock.Unlock()

	detail, exists := signalFileChooserButtonFileSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFileChooserButtonFileSetMap, connectionID)
}

//export filechooserbutton_fileSetHandler
func filechooserbutton_fileSetHandler(_ *C.GObject, data C.gpointer) {
	signalFileChooserButtonFileSetLock.RLock()
	defer signalFileChooserButtonFileSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFileChooserButtonFileSetMap[index].callback
	callback()
}

// Blacklisted : gtk_icon_theme_choose_icon

// ListContexts is a wrapper around the C function gtk_icon_theme_list_contexts.
func (recv *IconTheme) ListContexts() *glib.List {
	retC := C.gtk_icon_theme_list_contexts((*C.GtkIconTheme)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConvertWidgetToBinWindowCoords is a wrapper around the C function gtk_icon_view_convert_widget_to_bin_window_coords.
func (recv *IconView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_icon_view_convert_widget_to_bin_window_coords((*C.GtkIconView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// GetTooltipColumn is a wrapper around the C function gtk_icon_view_get_tooltip_column.
func (recv *IconView) GetTooltipColumn() int32 {
	retC := C.gtk_icon_view_get_tooltip_column((*C.GtkIconView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTooltipContext is a wrapper around the C function gtk_icon_view_get_tooltip_context.
func (recv *IconView) GetTooltipContext(x int32, y int32, keyboardTip bool) (bool, *TreeModel, *TreePath, *TreeIter) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyboard_tip :=
		boolToGboolean(keyboardTip)

	var c_model *C.GtkTreeModel

	var c_path *C.GtkTreePath

	var c_iter C.GtkTreeIter

	retC := C.gtk_icon_view_get_tooltip_context((*C.GtkIconView)(recv.native), &c_x, &c_y, c_keyboard_tip, &c_model, &c_path, &c_iter)
	retGo := retC == C.TRUE

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, model, path, iter
}

// SetTooltipCell is a wrapper around the C function gtk_icon_view_set_tooltip_cell.
func (recv *IconView) SetTooltipCell(tooltip *Tooltip, path *TreePath, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_icon_view_set_tooltip_cell((*C.GtkIconView)(recv.native), c_tooltip, c_path, c_cell)

	return
}

// SetTooltipColumn is a wrapper around the C function gtk_icon_view_set_tooltip_column.
func (recv *IconView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_icon_view_set_tooltip_column((*C.GtkIconView)(recv.native), c_column)

	return
}

// SetTooltipItem is a wrapper around the C function gtk_icon_view_set_tooltip_item.
func (recv *IconView) SetTooltipItem(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_icon_view_set_tooltip_item((*C.GtkIconView)(recv.native), c_tooltip, c_path)

	return
}

// Unsupported : gtk_list_store_set_valuesv : unsupported parameter values :

type signalMenuShellMoveSelectedDetail struct {
	callback  MenuShellSignalMoveSelectedCallback
	handlerID C.gulong
}

var signalMenuShellMoveSelectedId int
var signalMenuShellMoveSelectedMap = make(map[int]signalMenuShellMoveSelectedDetail)
var signalMenuShellMoveSelectedLock sync.RWMutex

// MenuShellSignalMoveSelectedCallback is a callback function for a 'move-selected' signal emitted from a MenuShell.
type MenuShellSignalMoveSelectedCallback func(distance int32) bool

/*
ConnectMoveSelected connects the callback to the 'move-selected' signal for the MenuShell.

The returned value represents the connection, and may be passed to DisconnectMoveSelected to remove it.
*/
func (recv *MenuShell) ConnectMoveSelected(callback MenuShellSignalMoveSelectedCallback) int {
	signalMenuShellMoveSelectedLock.Lock()
	defer signalMenuShellMoveSelectedLock.Unlock()

	signalMenuShellMoveSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.MenuShell_signal_connect_move_selected(instance, C.gpointer(uintptr(signalMenuShellMoveSelectedId)))

	detail := signalMenuShellMoveSelectedDetail{callback, handlerID}
	signalMenuShellMoveSelectedMap[signalMenuShellMoveSelectedId] = detail

	return signalMenuShellMoveSelectedId
}

/*
DisconnectMoveSelected disconnects a callback from the 'move-selected' signal for the MenuShell.

The connectionID should be a value returned from a call to ConnectMoveSelected.
*/
func (recv *MenuShell) DisconnectMoveSelected(connectionID int) {
	signalMenuShellMoveSelectedLock.Lock()
	defer signalMenuShellMoveSelectedLock.Unlock()

	detail, exists := signalMenuShellMoveSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMenuShellMoveSelectedMap, connectionID)
}

//export menushell_moveSelectedHandler
func menushell_moveSelectedHandler(_ *C.GObject, c_distance C.gint, data C.gpointer) C.gboolean {
	signalMenuShellMoveSelectedLock.RLock()
	defer signalMenuShellMoveSelectedLock.RUnlock()

	distance := int32(c_distance)

	index := int(uintptr(data))
	callback := signalMenuShellMoveSelectedMap[index].callback
	retGo := callback(distance)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// SetArrowTooltipMarkup is a wrapper around the C function gtk_menu_tool_button_set_arrow_tooltip_markup.
func (recv *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup((*C.GtkMenuToolButton)(recv.native), c_markup)

	return
}

// SetArrowTooltipText is a wrapper around the C function gtk_menu_tool_button_set_arrow_tooltip_text.
func (recv *MenuToolButton) SetArrowTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_menu_tool_button_set_arrow_tooltip_text((*C.GtkMenuToolButton)(recv.native), c_text)

	return
}

type signalNotebookCreateWindowDetail struct {
	callback  NotebookSignalCreateWindowCallback
	handlerID C.gulong
}

var signalNotebookCreateWindowId int
var signalNotebookCreateWindowMap = make(map[int]signalNotebookCreateWindowDetail)
var signalNotebookCreateWindowLock sync.RWMutex

// NotebookSignalCreateWindowCallback is a callback function for a 'create-window' signal emitted from a Notebook.
type NotebookSignalCreateWindowCallback func(page *Widget, x int32, y int32) Notebook

/*
ConnectCreateWindow connects the callback to the 'create-window' signal for the Notebook.

The returned value represents the connection, and may be passed to DisconnectCreateWindow to remove it.
*/
func (recv *Notebook) ConnectCreateWindow(callback NotebookSignalCreateWindowCallback) int {
	signalNotebookCreateWindowLock.Lock()
	defer signalNotebookCreateWindowLock.Unlock()

	signalNotebookCreateWindowId++
	instance := C.gpointer(recv.native)
	handlerID := C.Notebook_signal_connect_create_window(instance, C.gpointer(uintptr(signalNotebookCreateWindowId)))

	detail := signalNotebookCreateWindowDetail{callback, handlerID}
	signalNotebookCreateWindowMap[signalNotebookCreateWindowId] = detail

	return signalNotebookCreateWindowId
}

/*
DisconnectCreateWindow disconnects a callback from the 'create-window' signal for the Notebook.

The connectionID should be a value returned from a call to ConnectCreateWindow.
*/
func (recv *Notebook) DisconnectCreateWindow(connectionID int) {
	signalNotebookCreateWindowLock.Lock()
	defer signalNotebookCreateWindowLock.Unlock()

	detail, exists := signalNotebookCreateWindowMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalNotebookCreateWindowMap, connectionID)
}

//export notebook_createWindowHandler
func notebook_createWindowHandler(_ *C.GObject, c_page *C.GtkWidget, c_x C.gint, c_y C.gint, data C.gpointer) *C.GtkNotebook {
	signalNotebookCreateWindowLock.RLock()
	defer signalNotebookCreateWindowLock.RUnlock()

	page := WidgetNewFromC(unsafe.Pointer(c_page))

	x := int32(c_x)

	y := int32(c_y)

	index := int(uintptr(data))
	callback := signalNotebookCreateWindowMap[index].callback
	retGo := callback(page, x, y)
	retC :=
		(*C.GtkNotebook)(retGo.ToC())
	return retC
}

// PageSetupNewFromFile is a wrapper around the C function gtk_page_setup_new_from_file.
func PageSetupNewFromFile(fileName string) (*PageSetup, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_file(c_file_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

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

// PageSetupNewFromKeyFile is a wrapper around the C function gtk_page_setup_new_from_key_file.
func PageSetupNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PageSetup, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PageSetupNewFromC(unsafe.Pointer(retC))

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

// ToFile is a wrapper around the C function gtk_page_setup_to_file.
func (recv *PageSetup) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_to_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToKeyFile is a wrapper around the C function gtk_page_setup_to_key_file.
func (recv *PageSetup) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_page_setup_to_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name)

	return
}

// PrintSettingsNewFromFile is a wrapper around the C function gtk_print_settings_new_from_file.
func PrintSettingsNewFromFile(fileName string) (*PrintSettings, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_file(c_file_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

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

// PrintSettingsNewFromKeyFile is a wrapper around the C function gtk_print_settings_new_from_key_file.
func PrintSettingsNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PrintSettings, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PrintSettingsNewFromC(unsafe.Pointer(retC))

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

// ToFile is a wrapper around the C function gtk_print_settings_to_file.
func (recv *PrintSettings) ToFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_to_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToKeyFile is a wrapper around the C function gtk_print_settings_to_key_file.
func (recv *PrintSettings) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_print_settings_to_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name)

	return
}

// GetFillLevel is a wrapper around the C function gtk_range_get_fill_level.
func (recv *Range) GetFillLevel() float64 {
	retC := C.gtk_range_get_fill_level((*C.GtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetRestrictToFillLevel is a wrapper around the C function gtk_range_get_restrict_to_fill_level.
func (recv *Range) GetRestrictToFillLevel() bool {
	retC := C.gtk_range_get_restrict_to_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowFillLevel is a wrapper around the C function gtk_range_get_show_fill_level.
func (recv *Range) GetShowFillLevel() bool {
	retC := C.gtk_range_get_show_fill_level((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFillLevel is a wrapper around the C function gtk_range_set_fill_level.
func (recv *Range) SetFillLevel(fillLevel float64) {
	c_fill_level := (C.gdouble)(fillLevel)

	C.gtk_range_set_fill_level((*C.GtkRange)(recv.native), c_fill_level)

	return
}

// SetRestrictToFillLevel is a wrapper around the C function gtk_range_set_restrict_to_fill_level.
func (recv *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	c_restrict_to_fill_level :=
		boolToGboolean(restrictToFillLevel)

	C.gtk_range_set_restrict_to_fill_level((*C.GtkRange)(recv.native), c_restrict_to_fill_level)

	return
}

// SetShowFillLevel is a wrapper around the C function gtk_range_set_show_fill_level.
func (recv *Range) SetShowFillLevel(showFillLevel bool) {
	c_show_fill_level :=
		boolToGboolean(showFillLevel)

	C.gtk_range_set_show_fill_level((*C.GtkRange)(recv.native), c_show_fill_level)

	return
}

// RecentActionNew is a wrapper around the C function gtk_recent_action_new.
func RecentActionNew(name string, label string, tooltip string, stockId string) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_recent_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// RecentActionNewForManager is a wrapper around the C function gtk_recent_action_new_for_manager.
func RecentActionNewForManager(name string, label string, tooltip string, stockId string, manager *RecentManager) *RecentAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_manager := (*C.GtkRecentManager)(C.NULL)
	if manager != nil {
		c_manager = (*C.GtkRecentManager)(manager.ToC())
	}

	retC := C.gtk_recent_action_new_for_manager(c_name, c_label, c_tooltip, c_stock_id, c_manager)
	retGo := RecentActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetShowNumbers is a wrapper around the C function gtk_recent_action_get_show_numbers.
func (recv *RecentAction) GetShowNumbers() bool {
	retC := C.gtk_recent_action_get_show_numbers((*C.GtkRecentAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowNumbers is a wrapper around the C function gtk_recent_action_set_show_numbers.
func (recv *RecentAction) SetShowNumbers(showNumbers bool) {
	c_show_numbers :=
		boolToGboolean(showNumbers)

	C.gtk_recent_action_set_show_numbers((*C.GtkRecentAction)(recv.native), c_show_numbers)

	return
}

type signalScaleButtonPopdownDetail struct {
	callback  ScaleButtonSignalPopdownCallback
	handlerID C.gulong
}

var signalScaleButtonPopdownId int
var signalScaleButtonPopdownMap = make(map[int]signalScaleButtonPopdownDetail)
var signalScaleButtonPopdownLock sync.RWMutex

// ScaleButtonSignalPopdownCallback is a callback function for a 'popdown' signal emitted from a ScaleButton.
type ScaleButtonSignalPopdownCallback func()

/*
ConnectPopdown connects the callback to the 'popdown' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopdown to remove it.
*/
func (recv *ScaleButton) ConnectPopdown(callback ScaleButtonSignalPopdownCallback) int {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	signalScaleButtonPopdownId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popdown(instance, C.gpointer(uintptr(signalScaleButtonPopdownId)))

	detail := signalScaleButtonPopdownDetail{callback, handlerID}
	signalScaleButtonPopdownMap[signalScaleButtonPopdownId] = detail

	return signalScaleButtonPopdownId
}

/*
DisconnectPopdown disconnects a callback from the 'popdown' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopdown.
*/
func (recv *ScaleButton) DisconnectPopdown(connectionID int) {
	signalScaleButtonPopdownLock.Lock()
	defer signalScaleButtonPopdownLock.Unlock()

	detail, exists := signalScaleButtonPopdownMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopdownMap, connectionID)
}

//export scalebutton_popdownHandler
func scalebutton_popdownHandler(_ *C.GObject, data C.gpointer) {
	signalScaleButtonPopdownLock.RLock()
	defer signalScaleButtonPopdownLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScaleButtonPopdownMap[index].callback
	callback()
}

type signalScaleButtonPopupDetail struct {
	callback  ScaleButtonSignalPopupCallback
	handlerID C.gulong
}

var signalScaleButtonPopupId int
var signalScaleButtonPopupMap = make(map[int]signalScaleButtonPopupDetail)
var signalScaleButtonPopupLock sync.RWMutex

// ScaleButtonSignalPopupCallback is a callback function for a 'popup' signal emitted from a ScaleButton.
type ScaleButtonSignalPopupCallback func()

/*
ConnectPopup connects the callback to the 'popup' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectPopup to remove it.
*/
func (recv *ScaleButton) ConnectPopup(callback ScaleButtonSignalPopupCallback) int {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	signalScaleButtonPopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_popup(instance, C.gpointer(uintptr(signalScaleButtonPopupId)))

	detail := signalScaleButtonPopupDetail{callback, handlerID}
	signalScaleButtonPopupMap[signalScaleButtonPopupId] = detail

	return signalScaleButtonPopupId
}

/*
DisconnectPopup disconnects a callback from the 'popup' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectPopup.
*/
func (recv *ScaleButton) DisconnectPopup(connectionID int) {
	signalScaleButtonPopupLock.Lock()
	defer signalScaleButtonPopupLock.Unlock()

	detail, exists := signalScaleButtonPopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonPopupMap, connectionID)
}

//export scalebutton_popupHandler
func scalebutton_popupHandler(_ *C.GObject, data C.gpointer) {
	signalScaleButtonPopupLock.RLock()
	defer signalScaleButtonPopupLock.RUnlock()

	index := int(uintptr(data))
	callback := signalScaleButtonPopupMap[index].callback
	callback()
}

type signalScaleButtonValueChangedDetail struct {
	callback  ScaleButtonSignalValueChangedCallback
	handlerID C.gulong
}

var signalScaleButtonValueChangedId int
var signalScaleButtonValueChangedMap = make(map[int]signalScaleButtonValueChangedDetail)
var signalScaleButtonValueChangedLock sync.RWMutex

// ScaleButtonSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a ScaleButton.
type ScaleButtonSignalValueChangedCallback func(value float64)

/*
ConnectValueChanged connects the callback to the 'value-changed' signal for the ScaleButton.

The returned value represents the connection, and may be passed to DisconnectValueChanged to remove it.
*/
func (recv *ScaleButton) ConnectValueChanged(callback ScaleButtonSignalValueChangedCallback) int {
	signalScaleButtonValueChangedLock.Lock()
	defer signalScaleButtonValueChangedLock.Unlock()

	signalScaleButtonValueChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ScaleButton_signal_connect_value_changed(instance, C.gpointer(uintptr(signalScaleButtonValueChangedId)))

	detail := signalScaleButtonValueChangedDetail{callback, handlerID}
	signalScaleButtonValueChangedMap[signalScaleButtonValueChangedId] = detail

	return signalScaleButtonValueChangedId
}

/*
DisconnectValueChanged disconnects a callback from the 'value-changed' signal for the ScaleButton.

The connectionID should be a value returned from a call to ConnectValueChanged.
*/
func (recv *ScaleButton) DisconnectValueChanged(connectionID int) {
	signalScaleButtonValueChangedLock.Lock()
	defer signalScaleButtonValueChangedLock.Unlock()

	detail, exists := signalScaleButtonValueChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalScaleButtonValueChangedMap, connectionID)
}

//export scalebutton_valueChangedHandler
func scalebutton_valueChangedHandler(_ *C.GObject, c_value C.gdouble, data C.gpointer) {
	signalScaleButtonValueChangedLock.RLock()
	defer signalScaleButtonValueChangedLock.RUnlock()

	value := float64(c_value)

	index := int(uintptr(data))
	callback := signalScaleButtonValueChangedMap[index].callback
	callback(value)
}

// ScaleButtonNew is a wrapper around the C function gtk_scale_button_new.
func ScaleButtonNew(size IconSize, min float64, max float64, step float64, icons []string) *ScaleButton {
	c_size := (C.GtkIconSize)(size)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	c_icons_array := make([]*C.gchar, len(icons)+1, len(icons)+1)
	for i, item := range icons {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_icons_array[i] = c
	}
	c_icons_array[len(icons)] = nil
	c_icons_arrayPtr := &c_icons_array[0]
	c_icons := (**C.gchar)(unsafe.Pointer(c_icons_arrayPtr))

	retC := C.gtk_scale_button_new(c_size, c_min, c_max, c_step, c_icons)
	retGo := ScaleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAdjustment is a wrapper around the C function gtk_scale_button_get_adjustment.
func (recv *ScaleButton) GetAdjustment() *Adjustment {
	retC := C.gtk_scale_button_get_adjustment((*C.GtkScaleButton)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValue is a wrapper around the C function gtk_scale_button_get_value.
func (recv *ScaleButton) GetValue() float64 {
	retC := C.gtk_scale_button_get_value((*C.GtkScaleButton)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetAdjustment is a wrapper around the C function gtk_scale_button_set_adjustment.
func (recv *ScaleButton) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_scale_button_set_adjustment((*C.GtkScaleButton)(recv.native), c_adjustment)

	return
}

// SetIcons is a wrapper around the C function gtk_scale_button_set_icons.
func (recv *ScaleButton) SetIcons(icons []string) {
	c_icons_array := make([]*C.gchar, len(icons)+1, len(icons)+1)
	for i, item := range icons {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_icons_array[i] = c
	}
	c_icons_array[len(icons)] = nil
	c_icons_arrayPtr := &c_icons_array[0]
	c_icons := (**C.gchar)(unsafe.Pointer(c_icons_arrayPtr))

	C.gtk_scale_button_set_icons((*C.GtkScaleButton)(recv.native), c_icons)

	return
}

// SetValue is a wrapper around the C function gtk_scale_button_set_value.
func (recv *ScaleButton) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_scale_button_set_value((*C.GtkScaleButton)(recv.native), c_value)

	return
}

// GetScreen is a wrapper around the C function gtk_status_icon_get_screen.
func (recv *StatusIcon) GetScreen() *gdk.Screen {
	retC := C.gtk_status_icon_get_screen((*C.GtkStatusIcon)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetScreen is a wrapper around the C function gtk_status_icon_set_screen.
func (recv *StatusIcon) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_status_icon_set_screen((*C.GtkStatusIcon)(recv.native), c_screen)

	return
}

// AddMark is a wrapper around the C function gtk_text_buffer_add_mark.
func (recv *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	c_mark := (*C.GtkTextMark)(C.NULL)
	if mark != nil {
		c_mark = (*C.GtkTextMark)(mark.ToC())
	}

	c_where := (*C.GtkTextIter)(C.NULL)
	if where != nil {
		c_where = (*C.GtkTextIter)(where.ToC())
	}

	C.gtk_text_buffer_add_mark((*C.GtkTextBuffer)(recv.native), c_mark, c_where)

	return
}

// TextMarkNew is a wrapper around the C function gtk_text_mark_new.
func TextMarkNew(name string, leftGravity bool) *TextMark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_left_gravity :=
		boolToGboolean(leftGravity)

	retC := C.gtk_text_mark_new(c_name, c_left_gravity)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// SetTooltipMarkup is a wrapper around the C function gtk_tool_item_set_tooltip_markup.
func (recv *ToolItem) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tool_item_set_tooltip_markup((*C.GtkToolItem)(recv.native), c_markup)

	return
}

// SetTooltipText is a wrapper around the C function gtk_tool_item_set_tooltip_text.
func (recv *ToolItem) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tool_item_set_tooltip_text((*C.GtkToolItem)(recv.native), c_text)

	return
}

// TooltipTriggerTooltipQuery is a wrapper around the C function gtk_tooltip_trigger_tooltip_query.
func TooltipTriggerTooltipQuery(display *gdk.Display) {
	c_display := (*C.GdkDisplay)(C.NULL)
	if display != nil {
		c_display = (*C.GdkDisplay)(display.ToC())
	}

	C.gtk_tooltip_trigger_tooltip_query(c_display)

	return
}

// SetCustom is a wrapper around the C function gtk_tooltip_set_custom.
func (recv *Tooltip) SetCustom(customWidget *Widget) {
	c_custom_widget := (*C.GtkWidget)(C.NULL)
	if customWidget != nil {
		c_custom_widget = (*C.GtkWidget)(customWidget.ToC())
	}

	C.gtk_tooltip_set_custom((*C.GtkTooltip)(recv.native), c_custom_widget)

	return
}

// SetIcon is a wrapper around the C function gtk_tooltip_set_icon.
func (recv *Tooltip) SetIcon(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_tooltip_set_icon((*C.GtkTooltip)(recv.native), c_pixbuf)

	return
}

// SetIconFromStock is a wrapper around the C function gtk_tooltip_set_icon_from_stock.
func (recv *Tooltip) SetIconFromStock(stockId string, size IconSize) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_size := (C.GtkIconSize)(size)

	C.gtk_tooltip_set_icon_from_stock((*C.GtkTooltip)(recv.native), c_stock_id, c_size)

	return
}

// SetMarkup is a wrapper around the C function gtk_tooltip_set_markup.
func (recv *Tooltip) SetMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_tooltip_set_markup((*C.GtkTooltip)(recv.native), c_markup)

	return
}

// SetText is a wrapper around the C function gtk_tooltip_set_text.
func (recv *Tooltip) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_tooltip_set_text((*C.GtkTooltip)(recv.native), c_text)

	return
}

// SetTipArea is a wrapper around the C function gtk_tooltip_set_tip_area.
func (recv *Tooltip) SetTipArea(rect *gdk.Rectangle) {
	c_rect := (*C.GdkRectangle)(C.NULL)
	if rect != nil {
		c_rect = (*C.GdkRectangle)(rect.ToC())
	}

	C.gtk_tooltip_set_tip_area((*C.GtkTooltip)(recv.native), c_rect)

	return
}

// Unsupported : gtk_tree_store_set_valuesv : unsupported parameter values :

// ConvertBinWindowToTreeCoords is a wrapper around the C function gtk_tree_view_convert_bin_window_to_tree_coords.
func (recv *TreeView) ConvertBinWindowToTreeCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_bin_window_to_tree_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// ConvertBinWindowToWidgetCoords is a wrapper around the C function gtk_tree_view_convert_bin_window_to_widget_coords.
func (recv *TreeView) ConvertBinWindowToWidgetCoords(bx int32, by int32) (int32, int32) {
	c_bx := (C.gint)(bx)

	c_by := (C.gint)(by)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_bin_window_to_widget_coords((*C.GtkTreeView)(recv.native), c_bx, c_by, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// ConvertTreeToBinWindowCoords is a wrapper around the C function gtk_tree_view_convert_tree_to_bin_window_coords.
func (recv *TreeView) ConvertTreeToBinWindowCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_tree_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// ConvertTreeToWidgetCoords is a wrapper around the C function gtk_tree_view_convert_tree_to_widget_coords.
func (recv *TreeView) ConvertTreeToWidgetCoords(tx int32, ty int32) (int32, int32) {
	c_tx := (C.gint)(tx)

	c_ty := (C.gint)(ty)

	var c_wx C.gint

	var c_wy C.gint

	C.gtk_tree_view_convert_tree_to_widget_coords((*C.GtkTreeView)(recv.native), c_tx, c_ty, &c_wx, &c_wy)

	wx := (int32)(c_wx)

	wy := (int32)(c_wy)

	return wx, wy
}

// ConvertWidgetToBinWindowCoords is a wrapper around the C function gtk_tree_view_convert_widget_to_bin_window_coords.
func (recv *TreeView) ConvertWidgetToBinWindowCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_bx C.gint

	var c_by C.gint

	C.gtk_tree_view_convert_widget_to_bin_window_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_bx, &c_by)

	bx := (int32)(c_bx)

	by := (int32)(c_by)

	return bx, by
}

// ConvertWidgetToTreeCoords is a wrapper around the C function gtk_tree_view_convert_widget_to_tree_coords.
func (recv *TreeView) ConvertWidgetToTreeCoords(wx int32, wy int32) (int32, int32) {
	c_wx := (C.gint)(wx)

	c_wy := (C.gint)(wy)

	var c_tx C.gint

	var c_ty C.gint

	C.gtk_tree_view_convert_widget_to_tree_coords((*C.GtkTreeView)(recv.native), c_wx, c_wy, &c_tx, &c_ty)

	tx := (int32)(c_tx)

	ty := (int32)(c_ty)

	return tx, ty
}

// GetLevelIndentation is a wrapper around the C function gtk_tree_view_get_level_indentation.
func (recv *TreeView) GetLevelIndentation() int32 {
	retC := C.gtk_tree_view_get_level_indentation((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetShowExpanders is a wrapper around the C function gtk_tree_view_get_show_expanders.
func (recv *TreeView) GetShowExpanders() bool {
	retC := C.gtk_tree_view_get_show_expanders((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTooltipColumn is a wrapper around the C function gtk_tree_view_get_tooltip_column.
func (recv *TreeView) GetTooltipColumn() int32 {
	retC := C.gtk_tree_view_get_tooltip_column((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTooltipContext is a wrapper around the C function gtk_tree_view_get_tooltip_context.
func (recv *TreeView) GetTooltipContext(x int32, y int32, keyboardTip bool) (bool, *TreeModel, *TreePath, *TreeIter) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_keyboard_tip :=
		boolToGboolean(keyboardTip)

	var c_model *C.GtkTreeModel

	var c_path *C.GtkTreePath

	var c_iter C.GtkTreeIter

	retC := C.gtk_tree_view_get_tooltip_context((*C.GtkTreeView)(recv.native), &c_x, &c_y, c_keyboard_tip, &c_model, &c_path, &c_iter)
	retGo := retC == C.TRUE

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, model, path, iter
}

// IsRubberBandingActive is a wrapper around the C function gtk_tree_view_is_rubber_banding_active.
func (recv *TreeView) IsRubberBandingActive() bool {
	retC := C.gtk_tree_view_is_rubber_banding_active((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetLevelIndentation is a wrapper around the C function gtk_tree_view_set_level_indentation.
func (recv *TreeView) SetLevelIndentation(indentation int32) {
	c_indentation := (C.gint)(indentation)

	C.gtk_tree_view_set_level_indentation((*C.GtkTreeView)(recv.native), c_indentation)

	return
}

// SetShowExpanders is a wrapper around the C function gtk_tree_view_set_show_expanders.
func (recv *TreeView) SetShowExpanders(enabled bool) {
	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_tree_view_set_show_expanders((*C.GtkTreeView)(recv.native), c_enabled)

	return
}

// SetTooltipCell is a wrapper around the C function gtk_tree_view_set_tooltip_cell.
func (recv *TreeView) SetTooltipCell(tooltip *Tooltip, path *TreePath, column *TreeViewColumn, cell *CellRenderer) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_column := (*C.GtkTreeViewColumn)(C.NULL)
	if column != nil {
		c_column = (*C.GtkTreeViewColumn)(column.ToC())
	}

	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_tree_view_set_tooltip_cell((*C.GtkTreeView)(recv.native), c_tooltip, c_path, c_column, c_cell)

	return
}

// SetTooltipColumn is a wrapper around the C function gtk_tree_view_set_tooltip_column.
func (recv *TreeView) SetTooltipColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_view_set_tooltip_column((*C.GtkTreeView)(recv.native), c_column)

	return
}

// SetTooltipRow is a wrapper around the C function gtk_tree_view_set_tooltip_row.
func (recv *TreeView) SetTooltipRow(tooltip *Tooltip, path *TreePath) {
	c_tooltip := (*C.GtkTooltip)(C.NULL)
	if tooltip != nil {
		c_tooltip = (*C.GtkTooltip)(tooltip.ToC())
	}

	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	C.gtk_tree_view_set_tooltip_row((*C.GtkTreeView)(recv.native), c_tooltip, c_path)

	return
}

// GetTreeView is a wrapper around the C function gtk_tree_view_column_get_tree_view.
func (recv *TreeViewColumn) GetTreeView() *Widget {
	retC := C.gtk_tree_view_column_get_tree_view((*C.GtkTreeViewColumn)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// VolumeButtonNew is a wrapper around the C function gtk_volume_button_new.
func VolumeButtonNew() *VolumeButton {
	retC := C.gtk_volume_button_new()
	retGo := VolumeButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalWidgetDragFailedDetail struct {
	callback  WidgetSignalDragFailedCallback
	handlerID C.gulong
}

var signalWidgetDragFailedId int
var signalWidgetDragFailedMap = make(map[int]signalWidgetDragFailedDetail)
var signalWidgetDragFailedLock sync.RWMutex

// WidgetSignalDragFailedCallback is a callback function for a 'drag-failed' signal emitted from a Widget.
type WidgetSignalDragFailedCallback func(context *gdk.DragContext, result DragResult) bool

/*
ConnectDragFailed connects the callback to the 'drag-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDragFailed to remove it.
*/
func (recv *Widget) ConnectDragFailed(callback WidgetSignalDragFailedCallback) int {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	signalWidgetDragFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_drag_failed(instance, C.gpointer(uintptr(signalWidgetDragFailedId)))

	detail := signalWidgetDragFailedDetail{callback, handlerID}
	signalWidgetDragFailedMap[signalWidgetDragFailedId] = detail

	return signalWidgetDragFailedId
}

/*
DisconnectDragFailed disconnects a callback from the 'drag-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDragFailed.
*/
func (recv *Widget) DisconnectDragFailed(connectionID int) {
	signalWidgetDragFailedLock.Lock()
	defer signalWidgetDragFailedLock.Unlock()

	detail, exists := signalWidgetDragFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDragFailedMap, connectionID)
}

//export widget_dragFailedHandler
func widget_dragFailedHandler(_ *C.GObject, c_context *C.GdkDragContext, c_result C.GtkDragResult, data C.gpointer) C.gboolean {
	signalWidgetDragFailedLock.RLock()
	defer signalWidgetDragFailedLock.RUnlock()

	context := gdk.DragContextNewFromC(unsafe.Pointer(c_context))

	result := DragResult(c_result)

	index := int(uintptr(data))
	callback := signalWidgetDragFailedMap[index].callback
	retGo := callback(context, result)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetKeynavFailedDetail struct {
	callback  WidgetSignalKeynavFailedCallback
	handlerID C.gulong
}

var signalWidgetKeynavFailedId int
var signalWidgetKeynavFailedMap = make(map[int]signalWidgetKeynavFailedDetail)
var signalWidgetKeynavFailedLock sync.RWMutex

// WidgetSignalKeynavFailedCallback is a callback function for a 'keynav-failed' signal emitted from a Widget.
type WidgetSignalKeynavFailedCallback func(direction DirectionType) bool

/*
ConnectKeynavFailed connects the callback to the 'keynav-failed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectKeynavFailed to remove it.
*/
func (recv *Widget) ConnectKeynavFailed(callback WidgetSignalKeynavFailedCallback) int {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	signalWidgetKeynavFailedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_keynav_failed(instance, C.gpointer(uintptr(signalWidgetKeynavFailedId)))

	detail := signalWidgetKeynavFailedDetail{callback, handlerID}
	signalWidgetKeynavFailedMap[signalWidgetKeynavFailedId] = detail

	return signalWidgetKeynavFailedId
}

/*
DisconnectKeynavFailed disconnects a callback from the 'keynav-failed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectKeynavFailed.
*/
func (recv *Widget) DisconnectKeynavFailed(connectionID int) {
	signalWidgetKeynavFailedLock.Lock()
	defer signalWidgetKeynavFailedLock.Unlock()

	detail, exists := signalWidgetKeynavFailedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetKeynavFailedMap, connectionID)
}

//export widget_keynavFailedHandler
func widget_keynavFailedHandler(_ *C.GObject, c_direction C.GtkDirectionType, data C.gpointer) C.gboolean {
	signalWidgetKeynavFailedLock.RLock()
	defer signalWidgetKeynavFailedLock.RUnlock()

	direction := DirectionType(c_direction)

	index := int(uintptr(data))
	callback := signalWidgetKeynavFailedMap[index].callback
	retGo := callback(direction)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetQueryTooltipDetail struct {
	callback  WidgetSignalQueryTooltipCallback
	handlerID C.gulong
}

var signalWidgetQueryTooltipId int
var signalWidgetQueryTooltipMap = make(map[int]signalWidgetQueryTooltipDetail)
var signalWidgetQueryTooltipLock sync.RWMutex

// WidgetSignalQueryTooltipCallback is a callback function for a 'query-tooltip' signal emitted from a Widget.
type WidgetSignalQueryTooltipCallback func(x int32, y int32, keyboardMode bool, tooltip *Tooltip) bool

/*
ConnectQueryTooltip connects the callback to the 'query-tooltip' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectQueryTooltip to remove it.
*/
func (recv *Widget) ConnectQueryTooltip(callback WidgetSignalQueryTooltipCallback) int {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	signalWidgetQueryTooltipId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_query_tooltip(instance, C.gpointer(uintptr(signalWidgetQueryTooltipId)))

	detail := signalWidgetQueryTooltipDetail{callback, handlerID}
	signalWidgetQueryTooltipMap[signalWidgetQueryTooltipId] = detail

	return signalWidgetQueryTooltipId
}

/*
DisconnectQueryTooltip disconnects a callback from the 'query-tooltip' signal for the Widget.

The connectionID should be a value returned from a call to ConnectQueryTooltip.
*/
func (recv *Widget) DisconnectQueryTooltip(connectionID int) {
	signalWidgetQueryTooltipLock.Lock()
	defer signalWidgetQueryTooltipLock.Unlock()

	detail, exists := signalWidgetQueryTooltipMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetQueryTooltipMap, connectionID)
}

//export widget_queryTooltipHandler
func widget_queryTooltipHandler(_ *C.GObject, c_x C.gint, c_y C.gint, c_keyboard_mode C.gboolean, c_tooltip *C.GtkTooltip, data C.gpointer) C.gboolean {
	signalWidgetQueryTooltipLock.RLock()
	defer signalWidgetQueryTooltipLock.RUnlock()

	x := int32(c_x)

	y := int32(c_y)

	keyboardMode := c_keyboard_mode == C.TRUE

	tooltip := TooltipNewFromC(unsafe.Pointer(c_tooltip))

	index := int(uintptr(data))
	callback := signalWidgetQueryTooltipMap[index].callback
	retGo := callback(x, y, keyboardMode, tooltip)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// ErrorBell is a wrapper around the C function gtk_widget_error_bell.
func (recv *Widget) ErrorBell() {
	C.gtk_widget_error_bell((*C.GtkWidget)(recv.native))

	return
}

// GetHasTooltip is a wrapper around the C function gtk_widget_get_has_tooltip.
func (recv *Widget) GetHasTooltip() bool {
	retC := C.gtk_widget_get_has_tooltip((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTooltipMarkup is a wrapper around the C function gtk_widget_get_tooltip_markup.
func (recv *Widget) GetTooltipMarkup() string {
	retC := C.gtk_widget_get_tooltip_markup((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipText is a wrapper around the C function gtk_widget_get_tooltip_text.
func (recv *Widget) GetTooltipText() string {
	retC := C.gtk_widget_get_tooltip_text((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetTooltipWindow is a wrapper around the C function gtk_widget_get_tooltip_window.
func (recv *Widget) GetTooltipWindow() *Window {
	retC := C.gtk_widget_get_tooltip_window((*C.GtkWidget)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// KeynavFailed is a wrapper around the C function gtk_widget_keynav_failed.
func (recv *Widget) KeynavFailed(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_widget_keynav_failed((*C.GtkWidget)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// ModifyCursor is a wrapper around the C function gtk_widget_modify_cursor.
func (recv *Widget) ModifyCursor(primary *gdk.Color, secondary *gdk.Color) {
	c_primary := (*C.GdkColor)(C.NULL)
	if primary != nil {
		c_primary = (*C.GdkColor)(primary.ToC())
	}

	c_secondary := (*C.GdkColor)(C.NULL)
	if secondary != nil {
		c_secondary = (*C.GdkColor)(secondary.ToC())
	}

	C.gtk_widget_modify_cursor((*C.GtkWidget)(recv.native), c_primary, c_secondary)

	return
}

// SetHasTooltip is a wrapper around the C function gtk_widget_set_has_tooltip.
func (recv *Widget) SetHasTooltip(hasTooltip bool) {
	c_has_tooltip :=
		boolToGboolean(hasTooltip)

	C.gtk_widget_set_has_tooltip((*C.GtkWidget)(recv.native), c_has_tooltip)

	return
}

// SetTooltipMarkup is a wrapper around the C function gtk_widget_set_tooltip_markup.
func (recv *Widget) SetTooltipMarkup(markup string) {
	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_widget_set_tooltip_markup((*C.GtkWidget)(recv.native), c_markup)

	return
}

// SetTooltipText is a wrapper around the C function gtk_widget_set_tooltip_text.
func (recv *Widget) SetTooltipText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_widget_set_tooltip_text((*C.GtkWidget)(recv.native), c_text)

	return
}

// SetTooltipWindow is a wrapper around the C function gtk_widget_set_tooltip_window.
func (recv *Widget) SetTooltipWindow(customWindow *Window) {
	c_custom_window := (*C.GtkWindow)(C.NULL)
	if customWindow != nil {
		c_custom_window = (*C.GtkWindow)(customWindow.ToC())
	}

	C.gtk_widget_set_tooltip_window((*C.GtkWidget)(recv.native), c_custom_window)

	return
}

// TriggerTooltipQuery is a wrapper around the C function gtk_widget_trigger_tooltip_query.
func (recv *Widget) TriggerTooltipQuery() {
	C.gtk_widget_trigger_tooltip_query((*C.GtkWidget)(recv.native))

	return
}

// GetOpacity is a wrapper around the C function gtk_window_get_opacity.
func (recv *Window) GetOpacity() float64 {
	retC := C.gtk_window_get_opacity((*C.GtkWindow)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetOpacity is a wrapper around the C function gtk_window_set_opacity.
func (recv *Window) SetOpacity(opacity float64) {
	c_opacity := (C.gdouble)(opacity)

	C.gtk_window_set_opacity((*C.GtkWindow)(recv.native), c_opacity)

	return
}

// SetStartupId is a wrapper around the C function gtk_window_set_startup_id.
func (recv *Window) SetStartupId(startupId string) {
	c_startup_id := C.CString(startupId)
	defer C.free(unsafe.Pointer(c_startup_id))

	C.gtk_window_set_startup_id((*C.GtkWindow)(recv.native), c_startup_id)

	return
}

// Blacklisted : STOCK_DISCARD

// RcParseColorFull is a wrapper around the C function gtk_rc_parse_color_full.
func RcParseColorFull(scanner *glib.Scanner, style *RcStyle) (uint32, *gdk.Color) {
	c_scanner := (*C.GScanner)(C.NULL)
	if scanner != nil {
		c_scanner = (*C.GScanner)(scanner.ToC())
	}

	c_style := (*C.GtkRcStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkRcStyle)(style.ToC())
	}

	var c_color C.GdkColor

	retC := C.gtk_rc_parse_color_full(c_scanner, c_style, &c_color)
	retGo := (uint32)(retC)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// AddChild is a wrapper around the C function gtk_buildable_add_child.
func (recv *Buildable) AddChild(builder *Builder, child *gobject.Object, type_ string) {
	c_builder := (*C.GtkBuilder)(C.NULL)
	if builder != nil {
		c_builder = (*C.GtkBuilder)(builder.ToC())
	}

	c_child := (*C.GObject)(C.NULL)
	if child != nil {
		c_child = (*C.GObject)(child.ToC())
	}

	c_type := C.CString(type_)
	defer C.free(unsafe.Pointer(c_type))

	C.gtk_buildable_add_child((*C.GtkBuildable)(recv.native), c_builder, c_child, c_type)

	return
}

// ConstructChild is a wrapper around the C function gtk_buildable_construct_child.
func (recv *Buildable) ConstructChild(builder *Builder, name string) *gobject.Object {
	c_builder := (*C.GtkBuilder)(C.NULL)
	if builder != nil {
		c_builder = (*C.GtkBuilder)(builder.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_buildable_construct_child((*C.GtkBuildable)(recv.native), c_builder, c_name)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_buildable_custom_finished : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// Unsupported : gtk_buildable_custom_tag_end : unsupported parameter data : no type generator for gpointer (gpointer*) for param data

// Unsupported : gtk_buildable_custom_tag_start : unsupported parameter data : no type generator for gpointer (gpointer*) for param data

// GetInternalChild is a wrapper around the C function gtk_buildable_get_internal_child.
func (recv *Buildable) GetInternalChild(builder *Builder, childname string) *gobject.Object {
	c_builder := (*C.GtkBuilder)(C.NULL)
	if builder != nil {
		c_builder = (*C.GtkBuilder)(builder.ToC())
	}

	c_childname := C.CString(childname)
	defer C.free(unsafe.Pointer(c_childname))

	retC := C.gtk_buildable_get_internal_child((*C.GtkBuildable)(recv.native), c_builder, c_childname)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_buildable_get_name.
func (recv *Buildable) GetName() string {
	retC := C.gtk_buildable_get_name((*C.GtkBuildable)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// ParserFinished is a wrapper around the C function gtk_buildable_parser_finished.
func (recv *Buildable) ParserFinished(builder *Builder) {
	c_builder := (*C.GtkBuilder)(C.NULL)
	if builder != nil {
		c_builder = (*C.GtkBuilder)(builder.ToC())
	}

	C.gtk_buildable_parser_finished((*C.GtkBuildable)(recv.native), c_builder)

	return
}

// SetBuildableProperty is a wrapper around the C function gtk_buildable_set_buildable_property.
func (recv *Buildable) SetBuildableProperty(builder *Builder, name string, value *gobject.Value) {
	c_builder := (*C.GtkBuilder)(C.NULL)
	if builder != nil {
		c_builder = (*C.GtkBuilder)(builder.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_buildable_set_buildable_property((*C.GtkBuildable)(recv.native), c_builder, c_name, c_value)

	return
}

// SetName is a wrapper around the C function gtk_buildable_set_name.
func (recv *Buildable) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_buildable_set_name((*C.GtkBuildable)(recv.native), c_name)

	return
}

// GetCells is a wrapper around the C function gtk_cell_layout_get_cells.
func (recv *CellLayout) GetCells() *glib.List {
	retC := C.gtk_cell_layout_get_cells((*C.GtkCellLayout)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BindingEntrySkip is a wrapper around the C function gtk_binding_entry_skip.
func BindingEntrySkip(bindingSet *BindingSet, keyval uint32, modifiers gdk.ModifierType) {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_skip(c_binding_set, c_keyval, c_modifiers)

	return
}

// PaperSizeNewFromKeyFile is a wrapper around the C function gtk_paper_size_new_from_key_file.
func PaperSizeNewFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PaperSize, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_paper_size_new_from_key_file(c_key_file, c_group_name, &cThrowableError)
	retGo := PaperSizeNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// PaperSizeGetPaperSizes is a wrapper around the C function gtk_paper_size_get_paper_sizes.
func PaperSizeGetPaperSizes(includeCustom bool) *glib.List {
	c_include_custom :=
		boolToGboolean(includeCustom)

	retC := C.gtk_paper_size_get_paper_sizes(c_include_custom)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToKeyFile is a wrapper around the C function gtk_paper_size_to_key_file.
func (recv *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	C.gtk_paper_size_to_key_file((*C.GtkPaperSize)(recv.native), c_key_file, c_group_name)

	return
}