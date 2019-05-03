// Code generated - DO NOT EDIT.
// +build gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
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

	void cellarea_addEditableHandler(GObject *, GtkCellRenderer *, GtkCellEditable *, GdkRectangle *, gchar*, gpointer);

	static gulong CellArea_signal_connect_add_editable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add-editable", G_CALLBACK(cellarea_addEditableHandler), data);
	}

*/
/*

	void cellarea_applyAttributesHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gboolean, gboolean, gpointer);

	static gulong CellArea_signal_connect_apply_attributes(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "apply-attributes", G_CALLBACK(cellarea_applyAttributesHandler), data);
	}

*/
/*

	void cellarea_focusChangedHandler(GObject *, GtkCellRenderer *, gchar*, gpointer);

	static gulong CellArea_signal_connect_focus_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-changed", G_CALLBACK(cellarea_focusChangedHandler), data);
	}

*/
/*

	void cellarea_removeEditableHandler(GObject *, GtkCellRenderer *, GtkCellEditable *, gpointer);

	static gulong CellArea_signal_connect_remove_editable(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-editable", G_CALLBACK(cellarea_removeEditableHandler), data);
	}

*/
/*

	void stylecontext_changedHandler(GObject *, gpointer);

	static gulong StyleContext_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(stylecontext_changedHandler), data);
	}

*/
/*

	gboolean widget_drawHandler(GObject *, cairo_t *, gpointer);

	static gulong Widget_signal_connect_draw(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "draw", G_CALLBACK(widget_drawHandler), data);
	}

*/
/*

	void widget_stateFlagsChangedHandler(GObject *, GtkStateFlags, gpointer);

	static gulong Widget_signal_connect_state_flags_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-flags-changed", G_CALLBACK(widget_stateFlagsChangedHandler), data);
	}

*/
/*

	void widget_styleUpdatedHandler(GObject *, gpointer);

	static gulong Widget_signal_connect_style_updated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-updated", G_CALLBACK(widget_styleUpdatedHandler), data);
	}

*/
import "C"

// GetLicenseType is a wrapper around the C function gtk_about_dialog_get_license_type.
func (recv *AboutDialog) GetLicenseType() License {
	retC := C.gtk_about_dialog_get_license_type((*C.GtkAboutDialog)(recv.native))
	retGo := (License)(retC)

	return retGo
}

// SetLicenseType is a wrapper around the C function gtk_about_dialog_set_license_type.
func (recv *AboutDialog) SetLicenseType(licenseType License) {
	c_license_type := (C.GtkLicense)(licenseType)

	C.gtk_about_dialog_set_license_type((*C.GtkAboutDialog)(recv.native), c_license_type)

	return
}

// AppChooserButtonNew is a wrapper around the C function gtk_app_chooser_button_new.
func AppChooserButtonNew(contentType string) *AppChooserButton {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_button_new(c_content_type)
	retGo := AppChooserButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendCustomItem is a wrapper around the C function gtk_app_chooser_button_append_custom_item.
func (recv *AppChooserButton) AppendCustomItem(name string, label string, icon *gio.Icon) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_app_chooser_button_append_custom_item((*C.GtkAppChooserButton)(recv.native), c_name, c_label, c_icon)

	return
}

// AppendSeparator is a wrapper around the C function gtk_app_chooser_button_append_separator.
func (recv *AppChooserButton) AppendSeparator() {
	C.gtk_app_chooser_button_append_separator((*C.GtkAppChooserButton)(recv.native))

	return
}

// GetShowDialogItem is a wrapper around the C function gtk_app_chooser_button_get_show_dialog_item.
func (recv *AppChooserButton) GetShowDialogItem() bool {
	retC := C.gtk_app_chooser_button_get_show_dialog_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActiveCustomItem is a wrapper around the C function gtk_app_chooser_button_set_active_custom_item.
func (recv *AppChooserButton) SetActiveCustomItem(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_app_chooser_button_set_active_custom_item((*C.GtkAppChooserButton)(recv.native), c_name)

	return
}

// SetShowDialogItem is a wrapper around the C function gtk_app_chooser_button_set_show_dialog_item.
func (recv *AppChooserButton) SetShowDialogItem(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_button_set_show_dialog_item((*C.GtkAppChooserButton)(recv.native), c_setting)

	return
}

// AppChooserDialogNew is a wrapper around the C function gtk_app_chooser_dialog_new.
func AppChooserDialogNew(parent *Window, flags DialogFlags, file *gio.File) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_file := (*C.GFile)(file.ToC())

	retC := C.gtk_app_chooser_dialog_new(c_parent, c_flags, c_file)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserDialogNewForContentType is a wrapper around the C function gtk_app_chooser_dialog_new_for_content_type.
func AppChooserDialogNewForContentType(parent *Window, flags DialogFlags, contentType string) *AppChooserDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_dialog_new_for_content_type(c_parent, c_flags, c_content_type)
	retGo := AppChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_app_chooser_dialog_get_widget.
func (recv *AppChooserDialog) GetWidget() *Widget {
	retC := C.gtk_app_chooser_dialog_get_widget((*C.GtkAppChooserDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppChooserWidgetNew is a wrapper around the C function gtk_app_chooser_widget_new.
func AppChooserWidgetNew(contentType string) *AppChooserWidget {
	c_content_type := C.CString(contentType)
	defer C.free(unsafe.Pointer(c_content_type))

	retC := C.gtk_app_chooser_widget_new(c_content_type)
	retGo := AppChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDefaultText is a wrapper around the C function gtk_app_chooser_widget_get_default_text.
func (recv *AppChooserWidget) GetDefaultText() string {
	retC := C.gtk_app_chooser_widget_get_default_text((*C.GtkAppChooserWidget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetShowAll is a wrapper around the C function gtk_app_chooser_widget_get_show_all.
func (recv *AppChooserWidget) GetShowAll() bool {
	retC := C.gtk_app_chooser_widget_get_show_all((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowDefault is a wrapper around the C function gtk_app_chooser_widget_get_show_default.
func (recv *AppChooserWidget) GetShowDefault() bool {
	retC := C.gtk_app_chooser_widget_get_show_default((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowFallback is a wrapper around the C function gtk_app_chooser_widget_get_show_fallback.
func (recv *AppChooserWidget) GetShowFallback() bool {
	retC := C.gtk_app_chooser_widget_get_show_fallback((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowOther is a wrapper around the C function gtk_app_chooser_widget_get_show_other.
func (recv *AppChooserWidget) GetShowOther() bool {
	retC := C.gtk_app_chooser_widget_get_show_other((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowRecommended is a wrapper around the C function gtk_app_chooser_widget_get_show_recommended.
func (recv *AppChooserWidget) GetShowRecommended() bool {
	retC := C.gtk_app_chooser_widget_get_show_recommended((*C.GtkAppChooserWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowAll is a wrapper around the C function gtk_app_chooser_widget_set_show_all.
func (recv *AppChooserWidget) SetShowAll(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_all((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowDefault is a wrapper around the C function gtk_app_chooser_widget_set_show_default.
func (recv *AppChooserWidget) SetShowDefault(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_default((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowFallback is a wrapper around the C function gtk_app_chooser_widget_set_show_fallback.
func (recv *AppChooserWidget) SetShowFallback(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_fallback((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowOther is a wrapper around the C function gtk_app_chooser_widget_set_show_other.
func (recv *AppChooserWidget) SetShowOther(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_other((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// SetShowRecommended is a wrapper around the C function gtk_app_chooser_widget_set_show_recommended.
func (recv *AppChooserWidget) SetShowRecommended(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_widget_set_show_recommended((*C.GtkAppChooserWidget)(recv.native), c_setting)

	return
}

// ApplicationNew is a wrapper around the C function gtk_application_new.
func ApplicationNew(applicationId string, flags gio.ApplicationFlags) *Application {
	c_application_id := C.CString(applicationId)
	defer C.free(unsafe.Pointer(c_application_id))

	c_flags := (C.GApplicationFlags)(flags)

	retC := C.gtk_application_new(c_application_id, c_flags)
	retGo := ApplicationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddWindow is a wrapper around the C function gtk_application_add_window.
func (recv *Application) AddWindow(window *Window) {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	C.gtk_application_add_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// GetWindows is a wrapper around the C function gtk_application_get_windows.
func (recv *Application) GetWindows() *glib.List {
	retC := C.gtk_application_get_windows((*C.GtkApplication)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveWindow is a wrapper around the C function gtk_application_remove_window.
func (recv *Application) RemoveWindow(window *Window) {
	c_window := (*C.GtkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GtkWindow)(window.ToC())
	}

	C.gtk_application_remove_window((*C.GtkApplication)(recv.native), c_window)

	return
}

// NextPage is a wrapper around the C function gtk_assistant_next_page.
func (recv *Assistant) NextPage() {
	C.gtk_assistant_next_page((*C.GtkAssistant)(recv.native))

	return
}

// PreviousPage is a wrapper around the C function gtk_assistant_previous_page.
func (recv *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page((*C.GtkAssistant)(recv.native))

	return
}

// BoxNew is a wrapper around the C function gtk_box_new.
func BoxNew(orientation Orientation, spacing int32) *Box {
	c_orientation := (C.GtkOrientation)(orientation)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_box_new(c_orientation, c_spacing)
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonBoxNew is a wrapper around the C function gtk_button_box_new.
func ButtonBoxNew(orientation Orientation) *ButtonBox {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_button_box_new(c_orientation)
	retGo := ButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDayIsMarked is a wrapper around the C function gtk_calendar_get_day_is_marked.
func (recv *Calendar) GetDayIsMarked(day uint32) bool {
	c_day := (C.guint)(day)

	retC := C.gtk_calendar_get_day_is_marked((*C.GtkCalendar)(recv.native), c_day)
	retGo := retC == C.TRUE

	return retGo
}

type signalCellAreaAddEditableDetail struct {
	callback  CellAreaSignalAddEditableCallback
	handlerID C.gulong
}

var signalCellAreaAddEditableId int
var signalCellAreaAddEditableMap = make(map[int]signalCellAreaAddEditableDetail)
var signalCellAreaAddEditableLock sync.RWMutex

// CellAreaSignalAddEditableCallback is a callback function for a 'add-editable' signal emitted from a CellArea.
type CellAreaSignalAddEditableCallback func(renderer *CellRenderer, editable *CellEditable, cellArea *gdk.Rectangle, path string)

/*
ConnectAddEditable connects the callback to the 'add-editable' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectAddEditable to remove it.
*/
func (recv *CellArea) ConnectAddEditable(callback CellAreaSignalAddEditableCallback) int {
	signalCellAreaAddEditableLock.Lock()
	defer signalCellAreaAddEditableLock.Unlock()

	signalCellAreaAddEditableId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_add_editable(instance, C.gpointer(uintptr(signalCellAreaAddEditableId)))

	detail := signalCellAreaAddEditableDetail{callback, handlerID}
	signalCellAreaAddEditableMap[signalCellAreaAddEditableId] = detail

	return signalCellAreaAddEditableId
}

/*
DisconnectAddEditable disconnects a callback from the 'add-editable' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectAddEditable.
*/
func (recv *CellArea) DisconnectAddEditable(connectionID int) {
	signalCellAreaAddEditableLock.Lock()
	defer signalCellAreaAddEditableLock.Unlock()

	detail, exists := signalCellAreaAddEditableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaAddEditableMap, connectionID)
}

//export cellarea_addEditableHandler
func cellarea_addEditableHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_editable *C.GtkCellEditable, c_cell_area *C.GdkRectangle, c_path *C.gchar, data C.gpointer) {
	signalCellAreaAddEditableLock.RLock()
	defer signalCellAreaAddEditableLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	cellArea := gdk.RectangleNewFromC(unsafe.Pointer(c_cell_area))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellAreaAddEditableMap[index].callback
	callback(renderer, editable, cellArea, path)
}

type signalCellAreaApplyAttributesDetail struct {
	callback  CellAreaSignalApplyAttributesCallback
	handlerID C.gulong
}

var signalCellAreaApplyAttributesId int
var signalCellAreaApplyAttributesMap = make(map[int]signalCellAreaApplyAttributesDetail)
var signalCellAreaApplyAttributesLock sync.RWMutex

// CellAreaSignalApplyAttributesCallback is a callback function for a 'apply-attributes' signal emitted from a CellArea.
type CellAreaSignalApplyAttributesCallback func(model *TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)

/*
ConnectApplyAttributes connects the callback to the 'apply-attributes' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectApplyAttributes to remove it.
*/
func (recv *CellArea) ConnectApplyAttributes(callback CellAreaSignalApplyAttributesCallback) int {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	signalCellAreaApplyAttributesId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_apply_attributes(instance, C.gpointer(uintptr(signalCellAreaApplyAttributesId)))

	detail := signalCellAreaApplyAttributesDetail{callback, handlerID}
	signalCellAreaApplyAttributesMap[signalCellAreaApplyAttributesId] = detail

	return signalCellAreaApplyAttributesId
}

/*
DisconnectApplyAttributes disconnects a callback from the 'apply-attributes' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectApplyAttributes.
*/
func (recv *CellArea) DisconnectApplyAttributes(connectionID int) {
	signalCellAreaApplyAttributesLock.Lock()
	defer signalCellAreaApplyAttributesLock.Unlock()

	detail, exists := signalCellAreaApplyAttributesMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaApplyAttributesMap, connectionID)
}

//export cellarea_applyAttributesHandler
func cellarea_applyAttributesHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, c_is_expander C.gboolean, c_is_expanded C.gboolean, data C.gpointer) {
	signalCellAreaApplyAttributesLock.RLock()
	defer signalCellAreaApplyAttributesLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	isExpander := c_is_expander == C.TRUE

	isExpanded := c_is_expanded == C.TRUE

	index := int(uintptr(data))
	callback := signalCellAreaApplyAttributesMap[index].callback
	callback(model, iter, isExpander, isExpanded)
}

type signalCellAreaFocusChangedDetail struct {
	callback  CellAreaSignalFocusChangedCallback
	handlerID C.gulong
}

var signalCellAreaFocusChangedId int
var signalCellAreaFocusChangedMap = make(map[int]signalCellAreaFocusChangedDetail)
var signalCellAreaFocusChangedLock sync.RWMutex

// CellAreaSignalFocusChangedCallback is a callback function for a 'focus-changed' signal emitted from a CellArea.
type CellAreaSignalFocusChangedCallback func(renderer *CellRenderer, path string)

/*
ConnectFocusChanged connects the callback to the 'focus-changed' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectFocusChanged to remove it.
*/
func (recv *CellArea) ConnectFocusChanged(callback CellAreaSignalFocusChangedCallback) int {
	signalCellAreaFocusChangedLock.Lock()
	defer signalCellAreaFocusChangedLock.Unlock()

	signalCellAreaFocusChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_focus_changed(instance, C.gpointer(uintptr(signalCellAreaFocusChangedId)))

	detail := signalCellAreaFocusChangedDetail{callback, handlerID}
	signalCellAreaFocusChangedMap[signalCellAreaFocusChangedId] = detail

	return signalCellAreaFocusChangedId
}

/*
DisconnectFocusChanged disconnects a callback from the 'focus-changed' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectFocusChanged.
*/
func (recv *CellArea) DisconnectFocusChanged(connectionID int) {
	signalCellAreaFocusChangedLock.Lock()
	defer signalCellAreaFocusChangedLock.Unlock()

	detail, exists := signalCellAreaFocusChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaFocusChangedMap, connectionID)
}

//export cellarea_focusChangedHandler
func cellarea_focusChangedHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_path *C.gchar, data C.gpointer) {
	signalCellAreaFocusChangedLock.RLock()
	defer signalCellAreaFocusChangedLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	path := C.GoString(c_path)

	index := int(uintptr(data))
	callback := signalCellAreaFocusChangedMap[index].callback
	callback(renderer, path)
}

type signalCellAreaRemoveEditableDetail struct {
	callback  CellAreaSignalRemoveEditableCallback
	handlerID C.gulong
}

var signalCellAreaRemoveEditableId int
var signalCellAreaRemoveEditableMap = make(map[int]signalCellAreaRemoveEditableDetail)
var signalCellAreaRemoveEditableLock sync.RWMutex

// CellAreaSignalRemoveEditableCallback is a callback function for a 'remove-editable' signal emitted from a CellArea.
type CellAreaSignalRemoveEditableCallback func(renderer *CellRenderer, editable *CellEditable)

/*
ConnectRemoveEditable connects the callback to the 'remove-editable' signal for the CellArea.

The returned value represents the connection, and may be passed to DisconnectRemoveEditable to remove it.
*/
func (recv *CellArea) ConnectRemoveEditable(callback CellAreaSignalRemoveEditableCallback) int {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	signalCellAreaRemoveEditableId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellArea_signal_connect_remove_editable(instance, C.gpointer(uintptr(signalCellAreaRemoveEditableId)))

	detail := signalCellAreaRemoveEditableDetail{callback, handlerID}
	signalCellAreaRemoveEditableMap[signalCellAreaRemoveEditableId] = detail

	return signalCellAreaRemoveEditableId
}

/*
DisconnectRemoveEditable disconnects a callback from the 'remove-editable' signal for the CellArea.

The connectionID should be a value returned from a call to ConnectRemoveEditable.
*/
func (recv *CellArea) DisconnectRemoveEditable(connectionID int) {
	signalCellAreaRemoveEditableLock.Lock()
	defer signalCellAreaRemoveEditableLock.Unlock()

	detail, exists := signalCellAreaRemoveEditableMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellAreaRemoveEditableMap, connectionID)
}

//export cellarea_removeEditableHandler
func cellarea_removeEditableHandler(_ *C.GObject, c_renderer *C.GtkCellRenderer, c_editable *C.GtkCellEditable, data C.gpointer) {
	signalCellAreaRemoveEditableLock.RLock()
	defer signalCellAreaRemoveEditableLock.RUnlock()

	renderer := CellRendererNewFromC(unsafe.Pointer(c_renderer))

	editable := CellEditableNewFromC(unsafe.Pointer(c_editable))

	index := int(uintptr(data))
	callback := signalCellAreaRemoveEditableMap[index].callback
	callback(renderer, editable)
}

// Activate is a wrapper around the C function gtk_cell_area_activate.
func (recv *CellArea) Activate(context *CellAreaContext, widget *Widget, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	c_flags := (C.GtkCellRendererState)(flags)

	c_edit_only :=
		boolToGboolean(editOnly)

	retC := C.gtk_cell_area_activate((*C.GtkCellArea)(recv.native), c_context, c_widget, c_cell_area, c_flags, c_edit_only)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_activate_cell : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Add is a wrapper around the C function gtk_cell_area_add.
func (recv *CellArea) Add(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_add((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// AddFocusSibling is a wrapper around the C function gtk_cell_area_add_focus_sibling.
func (recv *CellArea) AddFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	C.gtk_cell_area_add_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Unsupported : gtk_cell_area_add_with_properties : unsupported parameter ... : varargs

// ApplyAttributes is a wrapper around the C function gtk_cell_area_apply_attributes.
func (recv *CellArea) ApplyAttributes(treeModel *TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	c_tree_model := (*C.GtkTreeModel)(treeModel.ToC())

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_is_expander :=
		boolToGboolean(isExpander)

	c_is_expanded :=
		boolToGboolean(isExpanded)

	C.gtk_cell_area_apply_attributes((*C.GtkCellArea)(recv.native), c_tree_model, c_iter, c_is_expander, c_is_expanded)

	return
}

// AttributeConnect is a wrapper around the C function gtk_cell_area_attribute_connect.
func (recv *CellArea) AttributeConnect(renderer *CellRenderer, attribute string, column int32) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_cell_area_attribute_connect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute, c_column)

	return
}

// AttributeDisconnect is a wrapper around the C function gtk_cell_area_attribute_disconnect.
func (recv *CellArea) AttributeDisconnect(renderer *CellRenderer, attribute string) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	C.gtk_cell_area_attribute_disconnect((*C.GtkCellArea)(recv.native), c_renderer, c_attribute)

	return
}

// Unsupported : gtk_cell_area_cell_get : unsupported parameter ... : varargs

// CellGetProperty is a wrapper around the C function gtk_cell_area_cell_get_property.
func (recv *CellArea) CellGetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_cell_area_cell_get_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_get_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// Unsupported : gtk_cell_area_cell_set : unsupported parameter ... : varargs

// CellSetProperty is a wrapper around the C function gtk_cell_area_cell_set_property.
func (recv *CellArea) CellSetProperty(renderer *CellRenderer, propertyName string, value *gobject.Value) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_cell_area_cell_set_property((*C.GtkCellArea)(recv.native), c_renderer, c_property_name, c_value)

	return
}

// Unsupported : gtk_cell_area_cell_set_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args

// CopyContext is a wrapper around the C function gtk_cell_area_copy_context.
func (recv *CellArea) CopyContext(context *CellAreaContext) *CellAreaContext {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	retC := C.gtk_cell_area_copy_context((*C.GtkCellArea)(recv.native), c_context)
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateContext is a wrapper around the C function gtk_cell_area_create_context.
func (recv *CellArea) CreateContext() *CellAreaContext {
	retC := C.gtk_cell_area_create_context((*C.GtkCellArea)(recv.native))
	retGo := CellAreaContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_cell_area_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Focus is a wrapper around the C function gtk_cell_area_focus.
func (recv *CellArea) Focus(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_cell_area_focus((*C.GtkCellArea)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_cell_area_foreach : unsupported parameter callback : no type generator for CellCallback (GtkCellCallback) for param callback

// Unsupported : gtk_cell_area_foreach_alloc : unsupported parameter callback : no type generator for CellAllocCallback (GtkCellAllocCallback) for param callback

// GetCellAllocation is a wrapper around the C function gtk_cell_area_get_cell_allocation.
func (recv *CellArea) GetCellAllocation(context *CellAreaContext, widget *Widget, renderer *CellRenderer, cellArea *gdk.Rectangle) *gdk.Rectangle {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	var c_allocation C.GdkRectangle

	C.gtk_cell_area_get_cell_allocation((*C.GtkCellArea)(recv.native), c_context, c_widget, c_renderer, c_cell_area, &c_allocation)

	allocation := gdk.RectangleNewFromC(unsafe.Pointer(&c_allocation))

	return allocation
}

// GetCellAtPosition is a wrapper around the C function gtk_cell_area_get_cell_at_position.
func (recv *CellArea) GetCellAtPosition(context *CellAreaContext, widget *Widget, cellArea *gdk.Rectangle, x int32, y int32) (*CellRenderer, *gdk.Rectangle) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_alloc_area C.GdkRectangle

	retC := C.gtk_cell_area_get_cell_at_position((*C.GtkCellArea)(recv.native), c_context, c_widget, c_cell_area, c_x, c_y, &c_alloc_area)
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	allocArea := gdk.RectangleNewFromC(unsafe.Pointer(&c_alloc_area))

	return retGo, allocArea
}

// GetCurrentPathString is a wrapper around the C function gtk_cell_area_get_current_path_string.
func (recv *CellArea) GetCurrentPathString() string {
	retC := C.gtk_cell_area_get_current_path_string((*C.GtkCellArea)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEditWidget is a wrapper around the C function gtk_cell_area_get_edit_widget.
func (recv *CellArea) GetEditWidget() *CellEditable {
	retC := C.gtk_cell_area_get_edit_widget((*C.GtkCellArea)(recv.native))
	retGo := CellEditableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEditedCell is a wrapper around the C function gtk_cell_area_get_edited_cell.
func (recv *CellArea) GetEditedCell() *CellRenderer {
	retC := C.gtk_cell_area_get_edited_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusCell is a wrapper around the C function gtk_cell_area_get_focus_cell.
func (recv *CellArea) GetFocusCell() *CellRenderer {
	retC := C.gtk_cell_area_get_focus_cell((*C.GtkCellArea)(recv.native))
	retGo := CellRendererNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusFromSibling is a wrapper around the C function gtk_cell_area_get_focus_from_sibling.
func (recv *CellArea) GetFocusFromSibling(renderer *CellRenderer) *CellRenderer {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_get_focus_from_sibling((*C.GtkCellArea)(recv.native), c_renderer)
	var retGo (*CellRenderer)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CellRendererNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFocusSiblings is a wrapper around the C function gtk_cell_area_get_focus_siblings.
func (recv *CellArea) GetFocusSiblings(renderer *CellRenderer) *glib.List {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_get_focus_siblings((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreferredHeight is a wrapper around the C function gtk_cell_area_get_preferred_height.
func (recv *CellArea) GetPreferredHeight(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_area_get_preferred_height_for_width.
func (recv *CellArea) GetPreferredHeightForWidth(context *CellAreaContext, widget *Widget, width int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_get_preferred_height_for_width((*C.GtkCellArea)(recv.native), c_context, c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredWidth is a wrapper around the C function gtk_cell_area_get_preferred_width.
func (recv *CellArea) GetPreferredWidth(context *CellAreaContext, widget *Widget) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width((*C.GtkCellArea)(recv.native), c_context, c_widget, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_area_get_preferred_width_for_height.
func (recv *CellArea) GetPreferredWidthForHeight(context *CellAreaContext, widget *Widget, height int32) (int32, int32) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_get_preferred_width_for_height((*C.GtkCellArea)(recv.native), c_context, c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetRequestMode is a wrapper around the C function gtk_cell_area_get_request_mode.
func (recv *CellArea) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_area_get_request_mode((*C.GtkCellArea)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// HasRenderer is a wrapper around the C function gtk_cell_area_has_renderer.
func (recv *CellArea) HasRenderer(renderer *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	retC := C.gtk_cell_area_has_renderer((*C.GtkCellArea)(recv.native), c_renderer)
	retGo := retC == C.TRUE

	return retGo
}

// InnerCellArea is a wrapper around the C function gtk_cell_area_inner_cell_area.
func (recv *CellArea) InnerCellArea(widget *Widget, cellArea *gdk.Rectangle) *gdk.Rectangle {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	var c_inner_area C.GdkRectangle

	C.gtk_cell_area_inner_cell_area((*C.GtkCellArea)(recv.native), c_widget, c_cell_area, &c_inner_area)

	innerArea := gdk.RectangleNewFromC(unsafe.Pointer(&c_inner_area))

	return innerArea
}

// IsActivatable is a wrapper around the C function gtk_cell_area_is_activatable.
func (recv *CellArea) IsActivatable() bool {
	retC := C.gtk_cell_area_is_activatable((*C.GtkCellArea)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsFocusSibling is a wrapper around the C function gtk_cell_area_is_focus_sibling.
func (recv *CellArea) IsFocusSibling(renderer *CellRenderer, sibling *CellRenderer) bool {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	retC := C.gtk_cell_area_is_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)
	retGo := retC == C.TRUE

	return retGo
}

// Remove is a wrapper around the C function gtk_cell_area_remove.
func (recv *CellArea) Remove(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_remove((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// RemoveFocusSibling is a wrapper around the C function gtk_cell_area_remove_focus_sibling.
func (recv *CellArea) RemoveFocusSibling(renderer *CellRenderer, sibling *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_sibling := (*C.GtkCellRenderer)(C.NULL)
	if sibling != nil {
		c_sibling = (*C.GtkCellRenderer)(sibling.ToC())
	}

	C.gtk_cell_area_remove_focus_sibling((*C.GtkCellArea)(recv.native), c_renderer, c_sibling)

	return
}

// Render is a wrapper around the C function gtk_cell_area_render.
func (recv *CellArea) Render(context *CellAreaContext, widget *Widget, cr *cairo.Context, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool) {
	c_context := (*C.GtkCellAreaContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkCellAreaContext)(context.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_background_area := (*C.GdkRectangle)(C.NULL)
	if backgroundArea != nil {
		c_background_area = (*C.GdkRectangle)(backgroundArea.ToC())
	}

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	c_flags := (C.GtkCellRendererState)(flags)

	c_paint_focus :=
		boolToGboolean(paintFocus)

	C.gtk_cell_area_render((*C.GtkCellArea)(recv.native), c_context, c_widget, c_cr, c_background_area, c_cell_area, c_flags, c_paint_focus)

	return
}

// RequestRenderer is a wrapper around the C function gtk_cell_area_request_renderer.
func (recv *CellArea) RequestRenderer(renderer *CellRenderer, orientation Orientation, widget *Widget, forSize int32) (int32, int32) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_orientation := (C.GtkOrientation)(orientation)

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_for_size := (C.gint)(forSize)

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_area_request_renderer((*C.GtkCellArea)(recv.native), c_renderer, c_orientation, c_widget, c_for_size, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// SetFocusCell is a wrapper around the C function gtk_cell_area_set_focus_cell.
func (recv *CellArea) SetFocusCell(renderer *CellRenderer) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	C.gtk_cell_area_set_focus_cell((*C.GtkCellArea)(recv.native), c_renderer)

	return
}

// StopEditing is a wrapper around the C function gtk_cell_area_stop_editing.
func (recv *CellArea) StopEditing(canceled bool) {
	c_canceled :=
		boolToGboolean(canceled)

	C.gtk_cell_area_stop_editing((*C.GtkCellArea)(recv.native), c_canceled)

	return
}

// CellAreaBoxNew is a wrapper around the C function gtk_cell_area_box_new.
func CellAreaBoxNew() *CellAreaBox {
	retC := C.gtk_cell_area_box_new()
	retGo := CellAreaBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_cell_area_box_get_spacing.
func (recv *CellAreaBox) GetSpacing() int32 {
	retC := C.gtk_cell_area_box_get_spacing((*C.GtkCellAreaBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_cell_area_box_pack_end.
func (recv *CellAreaBox) PackEnd(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_end((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// PackStart is a wrapper around the C function gtk_cell_area_box_pack_start.
func (recv *CellAreaBox) PackStart(renderer *CellRenderer, expand bool, align bool, fixed bool) {
	c_renderer := (*C.GtkCellRenderer)(C.NULL)
	if renderer != nil {
		c_renderer = (*C.GtkCellRenderer)(renderer.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	c_align :=
		boolToGboolean(align)

	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_cell_area_box_pack_start((*C.GtkCellAreaBox)(recv.native), c_renderer, c_expand, c_align, c_fixed)

	return
}

// SetSpacing is a wrapper around the C function gtk_cell_area_box_set_spacing.
func (recv *CellAreaBox) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_cell_area_box_set_spacing((*C.GtkCellAreaBox)(recv.native), c_spacing)

	return
}

// GetAllocation is a wrapper around the C function gtk_cell_area_context_get_allocation.
func (recv *CellAreaContext) GetAllocation() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_cell_area_context_get_allocation((*C.GtkCellAreaContext)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetArea is a wrapper around the C function gtk_cell_area_context_get_area.
func (recv *CellAreaContext) GetArea() *CellArea {
	retC := C.gtk_cell_area_context_get_area((*C.GtkCellAreaContext)(recv.native))
	retGo := CellAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreferredHeight is a wrapper around the C function gtk_cell_area_context_get_preferred_height.
func (recv *CellAreaContext) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height((*C.GtkCellAreaContext)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_area_context_get_preferred_height_for_width.
func (recv *CellAreaContext) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_area_context_get_preferred_height_for_width((*C.GtkCellAreaContext)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredWidth is a wrapper around the C function gtk_cell_area_context_get_preferred_width.
func (recv *CellAreaContext) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width((*C.GtkCellAreaContext)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_area_context_get_preferred_width_for_height.
func (recv *CellAreaContext) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_area_context_get_preferred_width_for_height((*C.GtkCellAreaContext)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// PushPreferredHeight is a wrapper around the C function gtk_cell_area_context_push_preferred_height.
func (recv *CellAreaContext) PushPreferredHeight(minimumHeight int32, naturalHeight int32) {
	c_minimum_height := (C.gint)(minimumHeight)

	c_natural_height := (C.gint)(naturalHeight)

	C.gtk_cell_area_context_push_preferred_height((*C.GtkCellAreaContext)(recv.native), c_minimum_height, c_natural_height)

	return
}

// PushPreferredWidth is a wrapper around the C function gtk_cell_area_context_push_preferred_width.
func (recv *CellAreaContext) PushPreferredWidth(minimumWidth int32, naturalWidth int32) {
	c_minimum_width := (C.gint)(minimumWidth)

	c_natural_width := (C.gint)(naturalWidth)

	C.gtk_cell_area_context_push_preferred_width((*C.GtkCellAreaContext)(recv.native), c_minimum_width, c_natural_width)

	return
}

// GetAlignedArea is a wrapper around the C function gtk_cell_renderer_get_aligned_area.
func (recv *CellRenderer) GetAlignedArea(widget *Widget, flags CellRendererState, cellArea *gdk.Rectangle) *gdk.Rectangle {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_flags := (C.GtkCellRendererState)(flags)

	c_cell_area := (*C.GdkRectangle)(C.NULL)
	if cellArea != nil {
		c_cell_area = (*C.GdkRectangle)(cellArea.ToC())
	}

	var c_aligned_area C.GdkRectangle

	C.gtk_cell_renderer_get_aligned_area((*C.GtkCellRenderer)(recv.native), c_widget, c_flags, c_cell_area, &c_aligned_area)

	alignedArea := gdk.RectangleNewFromC(unsafe.Pointer(&c_aligned_area))

	return alignedArea
}

// GetPreferredHeight is a wrapper around the C function gtk_cell_renderer_get_preferred_height.
func (recv *CellRenderer) GetPreferredHeight(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_height((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_cell_renderer_get_preferred_height_for_width.
func (recv *CellRenderer) GetPreferredHeightForWidth(widget *Widget, width int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_cell_renderer_get_preferred_height_for_width((*C.GtkCellRenderer)(recv.native), c_widget, c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredSize is a wrapper around the C function gtk_cell_renderer_get_preferred_size.
func (recv *CellRenderer) GetPreferredSize(widget *Widget) (*Requisition, *Requisition) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_cell_renderer_get_preferred_size((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// GetPreferredWidth is a wrapper around the C function gtk_cell_renderer_get_preferred_width.
func (recv *CellRenderer) GetPreferredWidth(widget *Widget) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	var c_minimum_size C.gint

	var c_natural_size C.gint

	C.gtk_cell_renderer_get_preferred_width((*C.GtkCellRenderer)(recv.native), c_widget, &c_minimum_size, &c_natural_size)

	minimumSize := (int32)(c_minimum_size)

	naturalSize := (int32)(c_natural_size)

	return minimumSize, naturalSize
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_cell_renderer_get_preferred_width_for_height.
func (recv *CellRenderer) GetPreferredWidthForHeight(widget *Widget, height int32) (int32, int32) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_cell_renderer_get_preferred_width_for_height((*C.GtkCellRenderer)(recv.native), c_widget, c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetRequestMode is a wrapper around the C function gtk_cell_renderer_get_request_mode.
func (recv *CellRenderer) GetRequestMode() SizeRequestMode {
	retC := C.gtk_cell_renderer_get_request_mode((*C.GtkCellRenderer)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// GetState is a wrapper around the C function gtk_cell_renderer_get_state.
func (recv *CellRenderer) GetState(widget *Widget, cellState CellRendererState) StateFlags {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cell_state := (C.GtkCellRendererState)(cellState)

	retC := C.gtk_cell_renderer_get_state((*C.GtkCellRenderer)(recv.native), c_widget, c_cell_state)
	retGo := (StateFlags)(retC)

	return retGo
}

// IsActivatable is a wrapper around the C function gtk_cell_renderer_is_activatable.
func (recv *CellRenderer) IsActivatable() bool {
	retC := C.gtk_cell_renderer_is_activatable((*C.GtkCellRenderer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDrawSensitive is a wrapper around the C function gtk_cell_view_get_draw_sensitive.
func (recv *CellView) GetDrawSensitive() bool {
	retC := C.gtk_cell_view_get_draw_sensitive((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFitModel is a wrapper around the C function gtk_cell_view_get_fit_model.
func (recv *CellView) GetFitModel() bool {
	retC := C.gtk_cell_view_get_fit_model((*C.GtkCellView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetBackgroundRgba is a wrapper around the C function gtk_cell_view_set_background_rgba.
func (recv *CellView) SetBackgroundRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_cell_view_set_background_rgba((*C.GtkCellView)(recv.native), c_rgba)

	return
}

// SetDrawSensitive is a wrapper around the C function gtk_cell_view_set_draw_sensitive.
func (recv *CellView) SetDrawSensitive(drawSensitive bool) {
	c_draw_sensitive :=
		boolToGboolean(drawSensitive)

	C.gtk_cell_view_set_draw_sensitive((*C.GtkCellView)(recv.native), c_draw_sensitive)

	return
}

// SetFitModel is a wrapper around the C function gtk_cell_view_set_fit_model.
func (recv *CellView) SetFitModel(fitModel bool) {
	c_fit_model :=
		boolToGboolean(fitModel)

	C.gtk_cell_view_set_fit_model((*C.GtkCellView)(recv.native), c_fit_model)

	return
}

// ColorButtonNewWithRgba is a wrapper around the C function gtk_color_button_new_with_rgba.
func ColorButtonNewWithRgba(rgba *gdk.RGBA) *ColorButton {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	retC := C.gtk_color_button_new_with_rgba(c_rgba)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRgba is a wrapper around the C function gtk_color_button_get_rgba.
func (recv *ColorButton) GetRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_button_get_rgba((*C.GtkColorButton)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// SetRgba is a wrapper around the C function gtk_color_button_set_rgba.
func (recv *ColorButton) SetRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_button_set_rgba((*C.GtkColorButton)(recv.native), c_rgba)

	return
}

// GetCurrentRgba is a wrapper around the C function gtk_color_selection_get_current_rgba.
func (recv *ColorSelection) GetCurrentRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_current_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// GetPreviousRgba is a wrapper around the C function gtk_color_selection_get_previous_rgba.
func (recv *ColorSelection) GetPreviousRgba() *gdk.RGBA {
	var c_rgba C.GdkRGBA

	C.gtk_color_selection_get_previous_rgba((*C.GtkColorSelection)(recv.native), &c_rgba)

	rgba := gdk.RGBANewFromC(unsafe.Pointer(&c_rgba))

	return rgba
}

// SetCurrentRgba is a wrapper around the C function gtk_color_selection_set_current_rgba.
func (recv *ColorSelection) SetCurrentRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_selection_set_current_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// SetPreviousRgba is a wrapper around the C function gtk_color_selection_set_previous_rgba.
func (recv *ColorSelection) SetPreviousRgba(rgba *gdk.RGBA) {
	c_rgba := (*C.GdkRGBA)(C.NULL)
	if rgba != nil {
		c_rgba = (*C.GdkRGBA)(rgba.ToC())
	}

	C.gtk_color_selection_set_previous_rgba((*C.GtkColorSelection)(recv.native), c_rgba)

	return
}

// GetActiveId is a wrapper around the C function gtk_combo_box_get_active_id.
func (recv *ComboBox) GetActiveId() string {
	retC := C.gtk_combo_box_get_active_id((*C.GtkComboBox)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetIdColumn is a wrapper around the C function gtk_combo_box_get_id_column.
func (recv *ComboBox) GetIdColumn() int32 {
	retC := C.gtk_combo_box_get_id_column((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPopupFixedWidth is a wrapper around the C function gtk_combo_box_get_popup_fixed_width.
func (recv *ComboBox) GetPopupFixedWidth() bool {
	retC := C.gtk_combo_box_get_popup_fixed_width((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PopupForDevice is a wrapper around the C function gtk_combo_box_popup_for_device.
func (recv *ComboBox) PopupForDevice(device *gdk.Device) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gtk_combo_box_popup_for_device((*C.GtkComboBox)(recv.native), c_device)

	return
}

// SetActiveId is a wrapper around the C function gtk_combo_box_set_active_id.
func (recv *ComboBox) SetActiveId(activeId string) bool {
	c_active_id := C.CString(activeId)
	defer C.free(unsafe.Pointer(c_active_id))

	retC := C.gtk_combo_box_set_active_id((*C.GtkComboBox)(recv.native), c_active_id)
	retGo := retC == C.TRUE

	return retGo
}

// SetIdColumn is a wrapper around the C function gtk_combo_box_set_id_column.
func (recv *ComboBox) SetIdColumn(idColumn int32) {
	c_id_column := (C.gint)(idColumn)

	C.gtk_combo_box_set_id_column((*C.GtkComboBox)(recv.native), c_id_column)

	return
}

// SetPopupFixedWidth is a wrapper around the C function gtk_combo_box_set_popup_fixed_width.
func (recv *ComboBox) SetPopupFixedWidth(fixed bool) {
	c_fixed :=
		boolToGboolean(fixed)

	C.gtk_combo_box_set_popup_fixed_width((*C.GtkComboBox)(recv.native), c_fixed)

	return
}

// Insert is a wrapper around the C function gtk_combo_box_text_insert.
func (recv *ComboBoxText) Insert(position int32, id string, text string) {
	c_position := (C.gint)(position)

	c_id := C.CString(id)
	defer C.free(unsafe.Pointer(c_id))

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_combo_box_text_insert((*C.GtkComboBoxText)(recv.native), c_position, c_id, c_text)

	return
}

// RemoveAll is a wrapper around the C function gtk_combo_box_text_remove_all.
func (recv *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all((*C.GtkComboBoxText)(recv.native))

	return
}

// GetIconArea is a wrapper around the C function gtk_entry_get_icon_area.
func (recv *Entry) GetIconArea(iconPos EntryIconPosition) *gdk.Rectangle {
	c_icon_pos := (C.GtkEntryIconPosition)(iconPos)

	var c_icon_area C.GdkRectangle

	C.gtk_entry_get_icon_area((*C.GtkEntry)(recv.native), c_icon_pos, &c_icon_area)

	iconArea := gdk.RectangleNewFromC(unsafe.Pointer(&c_icon_area))

	return iconArea
}

// GetTextArea is a wrapper around the C function gtk_entry_get_text_area.
func (recv *Entry) GetTextArea() *gdk.Rectangle {
	var c_text_area C.GdkRectangle

	C.gtk_entry_get_text_area((*C.GtkEntry)(recv.native), &c_text_area)

	textArea := gdk.RectangleNewFromC(unsafe.Pointer(&c_text_area))

	return textArea
}

// EntryCompletionNewWithArea is a wrapper around the C function gtk_entry_completion_new_with_area.
func EntryCompletionNewWithArea(area *CellArea) *EntryCompletion {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_entry_completion_new_with_area(c_area)
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// LoadSymbolic is a wrapper around the C function gtk_icon_info_load_symbolic.
func (recv *IconInfo) LoadSymbolic(fg *gdk.RGBA, successColor *gdk.RGBA, warningColor *gdk.RGBA, errorColor *gdk.RGBA) (*gdkpixbuf.Pixbuf, bool, error) {
	c_fg := (*C.GdkRGBA)(C.NULL)
	if fg != nil {
		c_fg = (*C.GdkRGBA)(fg.ToC())
	}

	c_success_color := (*C.GdkRGBA)(C.NULL)
	if successColor != nil {
		c_success_color = (*C.GdkRGBA)(successColor.ToC())
	}

	c_warning_color := (*C.GdkRGBA)(C.NULL)
	if warningColor != nil {
		c_warning_color = (*C.GdkRGBA)(warningColor.ToC())
	}

	c_error_color := (*C.GdkRGBA)(C.NULL)
	if errorColor != nil {
		c_error_color = (*C.GdkRGBA)(errorColor.ToC())
	}

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic((*C.GtkIconInfo)(recv.native), c_fg, c_success_color, c_warning_color, c_error_color, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goError
}

// LoadSymbolicForContext is a wrapper around the C function gtk_icon_info_load_symbolic_for_context.
func (recv *IconInfo) LoadSymbolicForContext(context *StyleContext) (*gdkpixbuf.Pixbuf, bool, error) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_context((*C.GtkIconInfo)(recv.native), c_context, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goError
}

// LoadSymbolicForStyle is a wrapper around the C function gtk_icon_info_load_symbolic_for_style.
func (recv *IconInfo) LoadSymbolicForStyle(style *Style, state StateType) (*gdkpixbuf.Pixbuf, bool, error) {
	c_style := (*C.GtkStyle)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyle)(style.ToC())
	}

	c_state := (C.GtkStateType)(state)

	var c_was_symbolic C.gboolean

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_symbolic_for_style((*C.GtkIconInfo)(recv.native), c_style, c_state, &c_was_symbolic, &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	wasSymbolic := c_was_symbolic == C.TRUE

	return retGo, wasSymbolic, goError
}

// IconViewNewWithArea is a wrapper around the C function gtk_icon_view_new_with_area.
func IconViewNewWithArea(area *CellArea) *IconView {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_icon_view_new_with_area(c_area)
	retGo := IconViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_popup_for_device : unsupported parameter func : no type generator for MenuPositionFunc (GtkMenuPositionFunc) for param func

// GetReserveIndicator is a wrapper around the C function gtk_menu_item_get_reserve_indicator.
func (recv *MenuItem) GetReserveIndicator() bool {
	retC := C.gtk_menu_item_get_reserve_indicator((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetReserveIndicator is a wrapper around the C function gtk_menu_item_set_reserve_indicator.
func (recv *MenuItem) SetReserveIndicator(reserve bool) {
	c_reserve :=
		boolToGboolean(reserve)

	C.gtk_menu_item_set_reserve_indicator((*C.GtkMenuItem)(recv.native), c_reserve)

	return
}

// GetParentShell is a wrapper around the C function gtk_menu_shell_get_parent_shell.
func (recv *MenuShell) GetParentShell() *Widget {
	retC := C.gtk_menu_shell_get_parent_shell((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectedItem is a wrapper around the C function gtk_menu_shell_get_selected_item.
func (recv *MenuShell) GetSelectedItem() *Widget {
	retC := C.gtk_menu_shell_get_selected_item((*C.GtkMenuShell)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NumerableIconNew is a wrapper around the C function gtk_numerable_icon_new.
func NumerableIconNew(baseIcon *gio.Icon) *gio.Icon {
	c_base_icon := (*C.GIcon)(baseIcon.ToC())

	retC := C.gtk_numerable_icon_new(c_base_icon)
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// NumerableIconNewWithStyleContext is a wrapper around the C function gtk_numerable_icon_new_with_style_context.
func NumerableIconNewWithStyleContext(baseIcon *gio.Icon, context *StyleContext) *gio.Icon {
	c_base_icon := (*C.GIcon)(baseIcon.ToC())

	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	retC := C.gtk_numerable_icon_new_with_style_context(c_base_icon, c_context)
	retGo := gio.IconNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBackgroundGicon is a wrapper around the C function gtk_numerable_icon_get_background_gicon.
func (recv *NumerableIcon) GetBackgroundGicon() *gio.Icon {
	retC := C.gtk_numerable_icon_get_background_gicon((*C.GtkNumerableIcon)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetBackgroundIconName is a wrapper around the C function gtk_numerable_icon_get_background_icon_name.
func (recv *NumerableIcon) GetBackgroundIconName() string {
	retC := C.gtk_numerable_icon_get_background_icon_name((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetCount is a wrapper around the C function gtk_numerable_icon_get_count.
func (recv *NumerableIcon) GetCount() int32 {
	retC := C.gtk_numerable_icon_get_count((*C.GtkNumerableIcon)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLabel is a wrapper around the C function gtk_numerable_icon_get_label.
func (recv *NumerableIcon) GetLabel() string {
	retC := C.gtk_numerable_icon_get_label((*C.GtkNumerableIcon)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetStyleContext is a wrapper around the C function gtk_numerable_icon_get_style_context.
func (recv *NumerableIcon) GetStyleContext() *StyleContext {
	retC := C.gtk_numerable_icon_get_style_context((*C.GtkNumerableIcon)(recv.native))
	var retGo (*StyleContext)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StyleContextNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetBackgroundGicon is a wrapper around the C function gtk_numerable_icon_set_background_gicon.
func (recv *NumerableIcon) SetBackgroundGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_numerable_icon_set_background_gicon((*C.GtkNumerableIcon)(recv.native), c_icon)

	return
}

// SetBackgroundIconName is a wrapper around the C function gtk_numerable_icon_set_background_icon_name.
func (recv *NumerableIcon) SetBackgroundIconName(iconName string) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	C.gtk_numerable_icon_set_background_icon_name((*C.GtkNumerableIcon)(recv.native), c_icon_name)

	return
}

// SetCount is a wrapper around the C function gtk_numerable_icon_set_count.
func (recv *NumerableIcon) SetCount(count int32) {
	c_count := (C.gint)(count)

	C.gtk_numerable_icon_set_count((*C.GtkNumerableIcon)(recv.native), c_count)

	return
}

// SetLabel is a wrapper around the C function gtk_numerable_icon_set_label.
func (recv *NumerableIcon) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_numerable_icon_set_label((*C.GtkNumerableIcon)(recv.native), c_label)

	return
}

// SetStyleContext is a wrapper around the C function gtk_numerable_icon_set_style_context.
func (recv *NumerableIcon) SetStyleContext(style *StyleContext) {
	c_style := (*C.GtkStyleContext)(C.NULL)
	if style != nil {
		c_style = (*C.GtkStyleContext)(style.ToC())
	}

	C.gtk_numerable_icon_set_style_context((*C.GtkNumerableIcon)(recv.native), c_style)

	return
}

// PanedNew is a wrapper around the C function gtk_paned_new.
func PanedNew(orientation Orientation) *Paned {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_paned_new(c_orientation)
	retGo := PanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShowText is a wrapper around the C function gtk_progress_bar_get_show_text.
func (recv *ProgressBar) GetShowText() bool {
	retC := C.gtk_progress_bar_get_show_text((*C.GtkProgressBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowText is a wrapper around the C function gtk_progress_bar_set_show_text.
func (recv *ProgressBar) SetShowText(showText bool) {
	c_show_text :=
		boolToGboolean(showText)

	C.gtk_progress_bar_set_show_text((*C.GtkProgressBar)(recv.native), c_show_text)

	return
}

// JoinGroup is a wrapper around the C function gtk_radio_action_join_group.
func (recv *RadioAction) JoinGroup(groupSource *RadioAction) {
	c_group_source := (*C.GtkRadioAction)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioAction)(groupSource.ToC())
	}

	C.gtk_radio_action_join_group((*C.GtkRadioAction)(recv.native), c_group_source)

	return
}

// JoinGroup is a wrapper around the C function gtk_radio_button_join_group.
func (recv *RadioButton) JoinGroup(groupSource *RadioButton) {
	c_group_source := (*C.GtkRadioButton)(C.NULL)
	if groupSource != nil {
		c_group_source = (*C.GtkRadioButton)(groupSource.ToC())
	}

	C.gtk_radio_button_join_group((*C.GtkRadioButton)(recv.native), c_group_source)

	return
}

// ScaleNew is a wrapper around the C function gtk_scale_new.
func ScaleNew(orientation Orientation, adjustment *Adjustment) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scale_new(c_orientation, c_adjustment)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScaleNewWithRange is a wrapper around the C function gtk_scale_new_with_range.
func ScaleNewWithRange(orientation Orientation, min float64, max float64, step float64) *Scale {
	c_orientation := (C.GtkOrientation)(orientation)

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_scale_new_with_range(c_orientation, c_min, c_max, c_step)
	retGo := ScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ScrollbarNew is a wrapper around the C function gtk_scrollbar_new.
func ScrollbarNew(orientation Orientation, adjustment *Adjustment) *Scrollbar {
	c_orientation := (C.GtkOrientation)(orientation)

	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	retC := C.gtk_scrollbar_new(c_orientation, c_adjustment)
	retGo := ScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMinContentHeight is a wrapper around the C function gtk_scrolled_window_get_min_content_height.
func (recv *ScrolledWindow) GetMinContentHeight() int32 {
	retC := C.gtk_scrolled_window_get_min_content_height((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinContentWidth is a wrapper around the C function gtk_scrolled_window_get_min_content_width.
func (recv *ScrolledWindow) GetMinContentWidth() int32 {
	retC := C.gtk_scrolled_window_get_min_content_width((*C.GtkScrolledWindow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetMinContentHeight is a wrapper around the C function gtk_scrolled_window_set_min_content_height.
func (recv *ScrolledWindow) SetMinContentHeight(height int32) {
	c_height := (C.gint)(height)

	C.gtk_scrolled_window_set_min_content_height((*C.GtkScrolledWindow)(recv.native), c_height)

	return
}

// SetMinContentWidth is a wrapper around the C function gtk_scrolled_window_set_min_content_width.
func (recv *ScrolledWindow) SetMinContentWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_scrolled_window_set_min_content_width((*C.GtkScrolledWindow)(recv.native), c_width)

	return
}

// SeparatorNew is a wrapper around the C function gtk_separator_new.
func SeparatorNew(orientation Orientation) *Separator {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_separator_new(c_orientation)
	retGo := SeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasContext is a wrapper around the C function gtk_style_has_context.
func (recv *Style) HasContext() bool {
	retC := C.gtk_style_has_context((*C.GtkStyle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

type signalStyleContextChangedDetail struct {
	callback  StyleContextSignalChangedCallback
	handlerID C.gulong
}

var signalStyleContextChangedId int
var signalStyleContextChangedMap = make(map[int]signalStyleContextChangedDetail)
var signalStyleContextChangedLock sync.RWMutex

// StyleContextSignalChangedCallback is a callback function for a 'changed' signal emitted from a StyleContext.
type StyleContextSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the StyleContext.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *StyleContext) ConnectChanged(callback StyleContextSignalChangedCallback) int {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	signalStyleContextChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.StyleContext_signal_connect_changed(instance, C.gpointer(uintptr(signalStyleContextChangedId)))

	detail := signalStyleContextChangedDetail{callback, handlerID}
	signalStyleContextChangedMap[signalStyleContextChangedId] = detail

	return signalStyleContextChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the StyleContext.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *StyleContext) DisconnectChanged(connectionID int) {
	signalStyleContextChangedLock.Lock()
	defer signalStyleContextChangedLock.Unlock()

	detail, exists := signalStyleContextChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleContextChangedMap, connectionID)
}

//export stylecontext_changedHandler
func stylecontext_changedHandler(_ *C.GObject, data C.gpointer) {
	signalStyleContextChangedLock.RLock()
	defer signalStyleContextChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleContextChangedMap[index].callback
	callback()
}

// StyleContextAddProviderForScreen is a wrapper around the C function gtk_style_context_add_provider_for_screen.
func StyleContextAddProviderForScreen(screen *gdk.Screen, provider *StyleProvider, priority uint32) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	c_priority := (C.guint)(priority)

	C.gtk_style_context_add_provider_for_screen(c_screen, c_provider, c_priority)

	return
}

// StyleContextRemoveProviderForScreen is a wrapper around the C function gtk_style_context_remove_provider_for_screen.
func StyleContextRemoveProviderForScreen(screen *gdk.Screen, provider *StyleProvider) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	C.gtk_style_context_remove_provider_for_screen(c_screen, c_provider)

	return
}

// StyleContextResetWidgets is a wrapper around the C function gtk_style_context_reset_widgets.
func StyleContextResetWidgets(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_style_context_reset_widgets(c_screen)

	return
}

// AddClass is a wrapper around the C function gtk_style_context_add_class.
func (recv *StyleContext) AddClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_add_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// AddProvider is a wrapper around the C function gtk_style_context_add_provider.
func (recv *StyleContext) AddProvider(provider *StyleProvider, priority uint32) {
	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	c_priority := (C.guint)(priority)

	C.gtk_style_context_add_provider((*C.GtkStyleContext)(recv.native), c_provider, c_priority)

	return
}

// AddRegion is a wrapper around the C function gtk_style_context_add_region.
func (recv *StyleContext) AddRegion(regionName string, flags RegionFlags) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_style_context_add_region((*C.GtkStyleContext)(recv.native), c_region_name, c_flags)

	return
}

// Unsupported : gtk_style_context_cancel_animations : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// Unsupported : gtk_style_context_get : unsupported parameter ... : varargs

// GetBackgroundColor is a wrapper around the C function gtk_style_context_get_background_color.
func (recv *StyleContext) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_background_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetBorder is a wrapper around the C function gtk_style_context_get_border.
func (recv *StyleContext) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_style_context_get_border((*C.GtkStyleContext)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// GetBorderColor is a wrapper around the C function gtk_style_context_get_border_color.
func (recv *StyleContext) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_border_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetColor is a wrapper around the C function gtk_style_context_get_color.
func (recv *StyleContext) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_style_context_get_color((*C.GtkStyleContext)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetDirection is a wrapper around the C function gtk_style_context_get_direction.
func (recv *StyleContext) GetDirection() TextDirection {
	retC := C.gtk_style_context_get_direction((*C.GtkStyleContext)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetFont is a wrapper around the C function gtk_style_context_get_font.
func (recv *StyleContext) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_style_context_get_font((*C.GtkStyleContext)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJunctionSides is a wrapper around the C function gtk_style_context_get_junction_sides.
func (recv *StyleContext) GetJunctionSides() JunctionSides {
	retC := C.gtk_style_context_get_junction_sides((*C.GtkStyleContext)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// GetMargin is a wrapper around the C function gtk_style_context_get_margin.
func (recv *StyleContext) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_style_context_get_margin((*C.GtkStyleContext)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// GetPadding is a wrapper around the C function gtk_style_context_get_padding.
func (recv *StyleContext) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_style_context_get_padding((*C.GtkStyleContext)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// GetPath is a wrapper around the C function gtk_style_context_get_path.
func (recv *StyleContext) GetPath() *WidgetPath {
	retC := C.gtk_style_context_get_path((*C.GtkStyleContext)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProperty is a wrapper around the C function gtk_style_context_get_property.
func (recv *StyleContext) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_style_context_get_property((*C.GtkStyleContext)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetState is a wrapper around the C function gtk_style_context_get_state.
func (recv *StyleContext) GetState() StateFlags {
	retC := C.gtk_style_context_get_state((*C.GtkStyleContext)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_style_context_get_style : unsupported parameter ... : varargs

// Unsupported : gtk_style_context_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_style_context_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// HasClass is a wrapper around the C function gtk_style_context_has_class.
func (recv *StyleContext) HasClass(className string) bool {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	retC := C.gtk_style_context_has_class((*C.GtkStyleContext)(recv.native), c_class_name)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_style_context_has_region : unsupported parameter flags_return : GtkRegionFlags* with indirection level of 1

// Invalidate is a wrapper around the C function gtk_style_context_invalidate.
func (recv *StyleContext) Invalidate() {
	C.gtk_style_context_invalidate((*C.GtkStyleContext)(recv.native))

	return
}

// ListClasses is a wrapper around the C function gtk_style_context_list_classes.
func (recv *StyleContext) ListClasses() *glib.List {
	retC := C.gtk_style_context_list_classes((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListRegions is a wrapper around the C function gtk_style_context_list_regions.
func (recv *StyleContext) ListRegions() *glib.List {
	retC := C.gtk_style_context_list_regions((*C.GtkStyleContext)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_context_notify_state_change : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// PopAnimatableRegion is a wrapper around the C function gtk_style_context_pop_animatable_region.
func (recv *StyleContext) PopAnimatableRegion() {
	C.gtk_style_context_pop_animatable_region((*C.GtkStyleContext)(recv.native))

	return
}

// Unsupported : gtk_style_context_push_animatable_region : unsupported parameter region_id : no type generator for gpointer (gpointer) for param region_id

// RemoveClass is a wrapper around the C function gtk_style_context_remove_class.
func (recv *StyleContext) RemoveClass(className string) {
	c_class_name := C.CString(className)
	defer C.free(unsafe.Pointer(c_class_name))

	C.gtk_style_context_remove_class((*C.GtkStyleContext)(recv.native), c_class_name)

	return
}

// RemoveProvider is a wrapper around the C function gtk_style_context_remove_provider.
func (recv *StyleContext) RemoveProvider(provider *StyleProvider) {
	c_provider := (*C.GtkStyleProvider)(provider.ToC())

	C.gtk_style_context_remove_provider((*C.GtkStyleContext)(recv.native), c_provider)

	return
}

// RemoveRegion is a wrapper around the C function gtk_style_context_remove_region.
func (recv *StyleContext) RemoveRegion(regionName string) {
	c_region_name := C.CString(regionName)
	defer C.free(unsafe.Pointer(c_region_name))

	C.gtk_style_context_remove_region((*C.GtkStyleContext)(recv.native), c_region_name)

	return
}

// Restore is a wrapper around the C function gtk_style_context_restore.
func (recv *StyleContext) Restore() {
	C.gtk_style_context_restore((*C.GtkStyleContext)(recv.native))

	return
}

// Save is a wrapper around the C function gtk_style_context_save.
func (recv *StyleContext) Save() {
	C.gtk_style_context_save((*C.GtkStyleContext)(recv.native))

	return
}

// ScrollAnimations is a wrapper around the C function gtk_style_context_scroll_animations.
func (recv *StyleContext) ScrollAnimations(window *gdk.Window, dx int32, dy int32) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	c_dx := (C.gint)(dx)

	c_dy := (C.gint)(dy)

	C.gtk_style_context_scroll_animations((*C.GtkStyleContext)(recv.native), c_window, c_dx, c_dy)

	return
}

// SetBackground is a wrapper around the C function gtk_style_context_set_background.
func (recv *StyleContext) SetBackground(window *gdk.Window) {
	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_style_context_set_background((*C.GtkStyleContext)(recv.native), c_window)

	return
}

// SetDirection is a wrapper around the C function gtk_style_context_set_direction.
func (recv *StyleContext) SetDirection(direction TextDirection) {
	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_style_context_set_direction((*C.GtkStyleContext)(recv.native), c_direction)

	return
}

// SetJunctionSides is a wrapper around the C function gtk_style_context_set_junction_sides.
func (recv *StyleContext) SetJunctionSides(sides JunctionSides) {
	c_sides := (C.GtkJunctionSides)(sides)

	C.gtk_style_context_set_junction_sides((*C.GtkStyleContext)(recv.native), c_sides)

	return
}

// SetPath is a wrapper around the C function gtk_style_context_set_path.
func (recv *StyleContext) SetPath(path *WidgetPath) {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	C.gtk_style_context_set_path((*C.GtkStyleContext)(recv.native), c_path)

	return
}

// SetScreen is a wrapper around the C function gtk_style_context_set_screen.
func (recv *StyleContext) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_style_context_set_screen((*C.GtkStyleContext)(recv.native), c_screen)

	return
}

// SetState is a wrapper around the C function gtk_style_context_set_state.
func (recv *StyleContext) SetState(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_style_context_set_state((*C.GtkStyleContext)(recv.native), c_flags)

	return
}

// StateIsRunning is a wrapper around the C function gtk_style_context_state_is_running.
func (recv *StyleContext) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_style_context_state_is_running((*C.GtkStyleContext)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

// gtk_style_properties_lookup_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser*) for param parse_func
// gtk_style_properties_register_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser) for param parse_func
// Unsupported : gtk_style_properties_get : unsupported parameter ... : varargs

// GetProperty is a wrapper around the C function gtk_style_properties_get_property.
func (recv *StyleProperties) GetProperty(property string, state StateFlags) (bool, *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	retC := C.gtk_style_properties_get_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, &c_value)
	retGo := retC == C.TRUE

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value
}

// Unsupported : gtk_style_properties_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// LookupColor is a wrapper around the C function gtk_style_properties_lookup_color.
func (recv *StyleProperties) LookupColor(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_style_properties_lookup_color((*C.GtkStyleProperties)(recv.native), c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MapColor is a wrapper around the C function gtk_style_properties_map_color.
func (recv *StyleProperties) MapColor(name string, color *SymbolicColor) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	C.gtk_style_properties_map_color((*C.GtkStyleProperties)(recv.native), c_name, c_color)

	return
}

// Merge is a wrapper around the C function gtk_style_properties_merge.
func (recv *StyleProperties) Merge(propsToMerge *StyleProperties, replace bool) {
	c_props_to_merge := (*C.GtkStyleProperties)(C.NULL)
	if propsToMerge != nil {
		c_props_to_merge = (*C.GtkStyleProperties)(propsToMerge.ToC())
	}

	c_replace :=
		boolToGboolean(replace)

	C.gtk_style_properties_merge((*C.GtkStyleProperties)(recv.native), c_props_to_merge, c_replace)

	return
}

// Unsupported : gtk_style_properties_set : unsupported parameter ... : varargs

// SetProperty is a wrapper around the C function gtk_style_properties_set_property.
func (recv *StyleProperties) SetProperty(property string, state StateFlags, value *gobject.Value) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	c_value := (*C.GValue)(C.NULL)
	if value != nil {
		c_value = (*C.GValue)(value.ToC())
	}

	C.gtk_style_properties_set_property((*C.GtkStyleProperties)(recv.native), c_property, c_state, c_value)

	return
}

// Unsupported : gtk_style_properties_set_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// UnsetProperty is a wrapper around the C function gtk_style_properties_unset_property.
func (recv *StyleProperties) UnsetProperty(property string, state StateFlags) {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	C.gtk_style_properties_unset_property((*C.GtkStyleProperties)(recv.native), c_property, c_state)

	return
}

// SwitchNew is a wrapper around the C function gtk_switch_new.
func SwitchNew() *Switch {
	retC := C.gtk_switch_new()
	retGo := SwitchNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_switch_get_active.
func (recv *Switch) GetActive() bool {
	retC := C.gtk_switch_get_active((*C.GtkSwitch)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_switch_set_active.
func (recv *Switch) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_switch_set_active((*C.GtkSwitch)(recv.native), c_is_active)

	return
}

// GetCursorLocations is a wrapper around the C function gtk_text_view_get_cursor_locations.
func (recv *TextView) GetCursorLocations(iter *TextIter) (*gdk.Rectangle, *gdk.Rectangle) {
	c_iter := (*C.GtkTextIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTextIter)(iter.ToC())
	}

	var c_strong C.GdkRectangle

	var c_weak C.GdkRectangle

	C.gtk_text_view_get_cursor_locations((*C.GtkTextView)(recv.native), c_iter, &c_strong, &c_weak)

	strong := gdk.RectangleNewFromC(unsafe.Pointer(&c_strong))

	weak := gdk.RectangleNewFromC(unsafe.Pointer(&c_weak))

	return strong, weak
}

// gtk_theming_engine_register_property : unsupported parameter parse_func : no type generator for StylePropertyParser (GtkStylePropertyParser) for param parse_func
// Unsupported : gtk_theming_engine_get : unsupported parameter ... : varargs

// GetBackgroundColor is a wrapper around the C function gtk_theming_engine_get_background_color.
func (recv *ThemingEngine) GetBackgroundColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_background_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetBorder is a wrapper around the C function gtk_theming_engine_get_border.
func (recv *ThemingEngine) GetBorder(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_border C.GtkBorder

	C.gtk_theming_engine_get_border((*C.GtkThemingEngine)(recv.native), c_state, &c_border)

	border := BorderNewFromC(unsafe.Pointer(&c_border))

	return border
}

// GetBorderColor is a wrapper around the C function gtk_theming_engine_get_border_color.
func (recv *ThemingEngine) GetBorderColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_border_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetColor is a wrapper around the C function gtk_theming_engine_get_color.
func (recv *ThemingEngine) GetColor(state StateFlags) *gdk.RGBA {
	c_state := (C.GtkStateFlags)(state)

	var c_color C.GdkRGBA

	C.gtk_theming_engine_get_color((*C.GtkThemingEngine)(recv.native), c_state, &c_color)

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetDirection is a wrapper around the C function gtk_theming_engine_get_direction.
func (recv *ThemingEngine) GetDirection() TextDirection {
	retC := C.gtk_theming_engine_get_direction((*C.GtkThemingEngine)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetFont is a wrapper around the C function gtk_theming_engine_get_font.
func (recv *ThemingEngine) GetFont(state StateFlags) *pango.FontDescription {
	c_state := (C.GtkStateFlags)(state)

	retC := C.gtk_theming_engine_get_font((*C.GtkThemingEngine)(recv.native), c_state)
	retGo := pango.FontDescriptionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJunctionSides is a wrapper around the C function gtk_theming_engine_get_junction_sides.
func (recv *ThemingEngine) GetJunctionSides() JunctionSides {
	retC := C.gtk_theming_engine_get_junction_sides((*C.GtkThemingEngine)(recv.native))
	retGo := (JunctionSides)(retC)

	return retGo
}

// GetMargin is a wrapper around the C function gtk_theming_engine_get_margin.
func (recv *ThemingEngine) GetMargin(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_margin C.GtkBorder

	C.gtk_theming_engine_get_margin((*C.GtkThemingEngine)(recv.native), c_state, &c_margin)

	margin := BorderNewFromC(unsafe.Pointer(&c_margin))

	return margin
}

// GetPadding is a wrapper around the C function gtk_theming_engine_get_padding.
func (recv *ThemingEngine) GetPadding(state StateFlags) *Border {
	c_state := (C.GtkStateFlags)(state)

	var c_padding C.GtkBorder

	C.gtk_theming_engine_get_padding((*C.GtkThemingEngine)(recv.native), c_state, &c_padding)

	padding := BorderNewFromC(unsafe.Pointer(&c_padding))

	return padding
}

// GetPath is a wrapper around the C function gtk_theming_engine_get_path.
func (recv *ThemingEngine) GetPath() *WidgetPath {
	retC := C.gtk_theming_engine_get_path((*C.GtkThemingEngine)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProperty is a wrapper around the C function gtk_theming_engine_get_property.
func (recv *ThemingEngine) GetProperty(property string, state StateFlags) *gobject.Value {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	c_state := (C.GtkStateFlags)(state)

	var c_value C.GValue

	C.gtk_theming_engine_get_property((*C.GtkThemingEngine)(recv.native), c_property, c_state, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// GetState is a wrapper around the C function gtk_theming_engine_get_state.
func (recv *ThemingEngine) GetState() StateFlags {
	retC := C.gtk_theming_engine_get_state((*C.GtkThemingEngine)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// Unsupported : gtk_theming_engine_get_style : unsupported parameter ... : varargs

// GetStyleProperty is a wrapper around the C function gtk_theming_engine_get_style_property.
func (recv *ThemingEngine) GetStyleProperty(propertyName string) *gobject.Value {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	var c_value C.GValue

	C.gtk_theming_engine_get_style_property((*C.GtkThemingEngine)(recv.native), c_property_name, &c_value)

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return value
}

// Unsupported : gtk_theming_engine_get_style_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// Unsupported : gtk_theming_engine_get_valist : unsupported parameter args : no type generator for va_list (va_list) for param args

// HasClass is a wrapper around the C function gtk_theming_engine_has_class.
func (recv *ThemingEngine) HasClass(styleClass string) bool {
	c_style_class := C.CString(styleClass)
	defer C.free(unsafe.Pointer(c_style_class))

	retC := C.gtk_theming_engine_has_class((*C.GtkThemingEngine)(recv.native), c_style_class)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_theming_engine_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// LookupColor is a wrapper around the C function gtk_theming_engine_lookup_color.
func (recv *ThemingEngine) LookupColor(colorName string) (bool, *gdk.RGBA) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkRGBA

	retC := C.gtk_theming_engine_lookup_color((*C.GtkThemingEngine)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// StateIsRunning is a wrapper around the C function gtk_theming_engine_state_is_running.
func (recv *ThemingEngine) StateIsRunning(state StateType) (bool, float64) {
	c_state := (C.GtkStateType)(state)

	var c_progress C.gdouble

	retC := C.gtk_theming_engine_state_is_running((*C.GtkThemingEngine)(recv.native), c_state, &c_progress)
	retGo := retC == C.TRUE

	progress := (float64)(c_progress)

	return retGo, progress
}

// IsBlankAtPos is a wrapper around the C function gtk_tree_view_is_blank_at_pos.
func (recv *TreeView) IsBlankAtPos(x int32, y int32) (bool, *TreePath, *TreeViewColumn, int32, int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	var c_path *C.GtkTreePath

	var c_column *C.GtkTreeViewColumn

	var c_cell_x C.gint

	var c_cell_y C.gint

	retC := C.gtk_tree_view_is_blank_at_pos((*C.GtkTreeView)(recv.native), c_x, c_y, &c_path, &c_column, &c_cell_x, &c_cell_y)
	retGo := retC == C.TRUE

	path := TreePathNewFromC(unsafe.Pointer(c_path))

	column := TreeViewColumnNewFromC(unsafe.Pointer(c_column))

	cellX := (int32)(c_cell_x)

	cellY := (int32)(c_cell_y)

	return retGo, path, column, cellX, cellY
}

// TreeViewColumnNewWithArea is a wrapper around the C function gtk_tree_view_column_new_with_area.
func TreeViewColumnNewWithArea(area *CellArea) *TreeViewColumn {
	c_area := (*C.GtkCellArea)(C.NULL)
	if area != nil {
		c_area = (*C.GtkCellArea)(area.ToC())
	}

	retC := C.gtk_tree_view_column_new_with_area(c_area)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetButton is a wrapper around the C function gtk_tree_view_column_get_button.
func (recv *TreeViewColumn) GetButton() *Widget {
	retC := C.gtk_tree_view_column_get_button((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalWidgetDrawDetail struct {
	callback  WidgetSignalDrawCallback
	handlerID C.gulong
}

var signalWidgetDrawId int
var signalWidgetDrawMap = make(map[int]signalWidgetDrawDetail)
var signalWidgetDrawLock sync.RWMutex

// WidgetSignalDrawCallback is a callback function for a 'draw' signal emitted from a Widget.
type WidgetSignalDrawCallback func(cr *cairo.Context) bool

/*
ConnectDraw connects the callback to the 'draw' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDraw to remove it.
*/
func (recv *Widget) ConnectDraw(callback WidgetSignalDrawCallback) int {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	signalWidgetDrawId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_draw(instance, C.gpointer(uintptr(signalWidgetDrawId)))

	detail := signalWidgetDrawDetail{callback, handlerID}
	signalWidgetDrawMap[signalWidgetDrawId] = detail

	return signalWidgetDrawId
}

/*
DisconnectDraw disconnects a callback from the 'draw' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDraw.
*/
func (recv *Widget) DisconnectDraw(connectionID int) {
	signalWidgetDrawLock.Lock()
	defer signalWidgetDrawLock.Unlock()

	detail, exists := signalWidgetDrawMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDrawMap, connectionID)
}

//export widget_drawHandler
func widget_drawHandler(_ *C.GObject, c_cr *C.cairo_t, data C.gpointer) C.gboolean {
	signalWidgetDrawLock.RLock()
	defer signalWidgetDrawLock.RUnlock()

	cr := cairo.ContextNewFromC(unsafe.Pointer(c_cr))

	index := int(uintptr(data))
	callback := signalWidgetDrawMap[index].callback
	retGo := callback(cr)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalWidgetStateFlagsChangedDetail struct {
	callback  WidgetSignalStateFlagsChangedCallback
	handlerID C.gulong
}

var signalWidgetStateFlagsChangedId int
var signalWidgetStateFlagsChangedMap = make(map[int]signalWidgetStateFlagsChangedDetail)
var signalWidgetStateFlagsChangedLock sync.RWMutex

// WidgetSignalStateFlagsChangedCallback is a callback function for a 'state-flags-changed' signal emitted from a Widget.
type WidgetSignalStateFlagsChangedCallback func(flags StateFlags)

/*
ConnectStateFlagsChanged connects the callback to the 'state-flags-changed' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStateFlagsChanged to remove it.
*/
func (recv *Widget) ConnectStateFlagsChanged(callback WidgetSignalStateFlagsChangedCallback) int {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	signalWidgetStateFlagsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_state_flags_changed(instance, C.gpointer(uintptr(signalWidgetStateFlagsChangedId)))

	detail := signalWidgetStateFlagsChangedDetail{callback, handlerID}
	signalWidgetStateFlagsChangedMap[signalWidgetStateFlagsChangedId] = detail

	return signalWidgetStateFlagsChangedId
}

/*
DisconnectStateFlagsChanged disconnects a callback from the 'state-flags-changed' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStateFlagsChanged.
*/
func (recv *Widget) DisconnectStateFlagsChanged(connectionID int) {
	signalWidgetStateFlagsChangedLock.Lock()
	defer signalWidgetStateFlagsChangedLock.Unlock()

	detail, exists := signalWidgetStateFlagsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStateFlagsChangedMap, connectionID)
}

//export widget_stateFlagsChangedHandler
func widget_stateFlagsChangedHandler(_ *C.GObject, c_flags C.GtkStateFlags, data C.gpointer) {
	signalWidgetStateFlagsChangedLock.RLock()
	defer signalWidgetStateFlagsChangedLock.RUnlock()

	flags := StateFlags(c_flags)

	index := int(uintptr(data))
	callback := signalWidgetStateFlagsChangedMap[index].callback
	callback(flags)
}

type signalWidgetStyleUpdatedDetail struct {
	callback  WidgetSignalStyleUpdatedCallback
	handlerID C.gulong
}

var signalWidgetStyleUpdatedId int
var signalWidgetStyleUpdatedMap = make(map[int]signalWidgetStyleUpdatedDetail)
var signalWidgetStyleUpdatedLock sync.RWMutex

// WidgetSignalStyleUpdatedCallback is a callback function for a 'style-updated' signal emitted from a Widget.
type WidgetSignalStyleUpdatedCallback func()

/*
ConnectStyleUpdated connects the callback to the 'style-updated' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectStyleUpdated to remove it.
*/
func (recv *Widget) ConnectStyleUpdated(callback WidgetSignalStyleUpdatedCallback) int {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	signalWidgetStyleUpdatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_style_updated(instance, C.gpointer(uintptr(signalWidgetStyleUpdatedId)))

	detail := signalWidgetStyleUpdatedDetail{callback, handlerID}
	signalWidgetStyleUpdatedMap[signalWidgetStyleUpdatedId] = detail

	return signalWidgetStyleUpdatedId
}

/*
DisconnectStyleUpdated disconnects a callback from the 'style-updated' signal for the Widget.

The connectionID should be a value returned from a call to ConnectStyleUpdated.
*/
func (recv *Widget) DisconnectStyleUpdated(connectionID int) {
	signalWidgetStyleUpdatedLock.Lock()
	defer signalWidgetStyleUpdatedLock.Unlock()

	detail, exists := signalWidgetStyleUpdatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetStyleUpdatedMap, connectionID)
}

//export widget_styleUpdatedHandler
func widget_styleUpdatedHandler(_ *C.GObject, data C.gpointer) {
	signalWidgetStyleUpdatedLock.RLock()
	defer signalWidgetStyleUpdatedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalWidgetStyleUpdatedMap[index].callback
	callback()
}

// AddDeviceEvents is a wrapper around the C function gtk_widget_add_device_events.
func (recv *Widget) AddDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_add_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// DeviceIsShadowed is a wrapper around the C function gtk_widget_device_is_shadowed.
func (recv *Widget) DeviceIsShadowed(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_device_is_shadowed((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// Draw is a wrapper around the C function gtk_widget_draw.
func (recv *Widget) Draw(cr *cairo.Context) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	C.gtk_widget_draw((*C.GtkWidget)(recv.native), c_cr)

	return
}

// GetDeviceEnabled is a wrapper around the C function gtk_widget_get_device_enabled.
func (recv *Widget) GetDeviceEnabled(device *gdk.Device) bool {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_get_device_enabled((*C.GtkWidget)(recv.native), c_device)
	retGo := retC == C.TRUE

	return retGo
}

// GetDeviceEvents is a wrapper around the C function gtk_widget_get_device_events.
func (recv *Widget) GetDeviceEvents(device *gdk.Device) gdk.EventMask {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_widget_get_device_events((*C.GtkWidget)(recv.native), c_device)
	retGo := (gdk.EventMask)(retC)

	return retGo
}

// GetMarginBottom is a wrapper around the C function gtk_widget_get_margin_bottom.
func (recv *Widget) GetMarginBottom() int32 {
	retC := C.gtk_widget_get_margin_bottom((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginLeft is a wrapper around the C function gtk_widget_get_margin_left.
func (recv *Widget) GetMarginLeft() int32 {
	retC := C.gtk_widget_get_margin_left((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginRight is a wrapper around the C function gtk_widget_get_margin_right.
func (recv *Widget) GetMarginRight() int32 {
	retC := C.gtk_widget_get_margin_right((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMarginTop is a wrapper around the C function gtk_widget_get_margin_top.
func (recv *Widget) GetMarginTop() int32 {
	retC := C.gtk_widget_get_margin_top((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPreferredHeight is a wrapper around the C function gtk_widget_get_preferred_height.
func (recv *Widget) GetPreferredHeight() (int32, int32) {
	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height((*C.GtkWidget)(recv.native), &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredHeightForWidth is a wrapper around the C function gtk_widget_get_preferred_height_for_width.
func (recv *Widget) GetPreferredHeightForWidth(width int32) (int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	C.gtk_widget_get_preferred_height_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	return minimumHeight, naturalHeight
}

// GetPreferredSize is a wrapper around the C function gtk_widget_get_preferred_size.
func (recv *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	var c_minimum_size C.GtkRequisition

	var c_natural_size C.GtkRequisition

	C.gtk_widget_get_preferred_size((*C.GtkWidget)(recv.native), &c_minimum_size, &c_natural_size)

	minimumSize := RequisitionNewFromC(unsafe.Pointer(&c_minimum_size))

	naturalSize := RequisitionNewFromC(unsafe.Pointer(&c_natural_size))

	return minimumSize, naturalSize
}

// GetPreferredWidth is a wrapper around the C function gtk_widget_get_preferred_width.
func (recv *Widget) GetPreferredWidth() (int32, int32) {
	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width((*C.GtkWidget)(recv.native), &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetPreferredWidthForHeight is a wrapper around the C function gtk_widget_get_preferred_width_for_height.
func (recv *Widget) GetPreferredWidthForHeight(height int32) (int32, int32) {
	c_height := (C.gint)(height)

	var c_minimum_width C.gint

	var c_natural_width C.gint

	C.gtk_widget_get_preferred_width_for_height((*C.GtkWidget)(recv.native), c_height, &c_minimum_width, &c_natural_width)

	minimumWidth := (int32)(c_minimum_width)

	naturalWidth := (int32)(c_natural_width)

	return minimumWidth, naturalWidth
}

// GetRequestMode is a wrapper around the C function gtk_widget_get_request_mode.
func (recv *Widget) GetRequestMode() SizeRequestMode {
	retC := C.gtk_widget_get_request_mode((*C.GtkWidget)(recv.native))
	retGo := (SizeRequestMode)(retC)

	return retGo
}

// GetStateFlags is a wrapper around the C function gtk_widget_get_state_flags.
func (recv *Widget) GetStateFlags() StateFlags {
	retC := C.gtk_widget_get_state_flags((*C.GtkWidget)(recv.native))
	retGo := (StateFlags)(retC)

	return retGo
}

// InputShapeCombineRegion is a wrapper around the C function gtk_widget_input_shape_combine_region.
func (recv *Widget) InputShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_input_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// OverrideBackgroundColor is a wrapper around the C function gtk_widget_override_background_color.
func (recv *Widget) OverrideBackgroundColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_background_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// OverrideColor is a wrapper around the C function gtk_widget_override_color.
func (recv *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	c_state := (C.GtkStateFlags)(state)

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_color((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// OverrideCursor is a wrapper around the C function gtk_widget_override_cursor.
func (recv *Widget) OverrideCursor(cursor *gdk.RGBA, secondaryCursor *gdk.RGBA) {
	c_cursor := (*C.GdkRGBA)(C.NULL)
	if cursor != nil {
		c_cursor = (*C.GdkRGBA)(cursor.ToC())
	}

	c_secondary_cursor := (*C.GdkRGBA)(C.NULL)
	if secondaryCursor != nil {
		c_secondary_cursor = (*C.GdkRGBA)(secondaryCursor.ToC())
	}

	C.gtk_widget_override_cursor((*C.GtkWidget)(recv.native), c_cursor, c_secondary_cursor)

	return
}

// OverrideFont is a wrapper around the C function gtk_widget_override_font.
func (recv *Widget) OverrideFont(fontDesc *pango.FontDescription) {
	c_font_desc := (*C.PangoFontDescription)(C.NULL)
	if fontDesc != nil {
		c_font_desc = (*C.PangoFontDescription)(fontDesc.ToC())
	}

	C.gtk_widget_override_font((*C.GtkWidget)(recv.native), c_font_desc)

	return
}

// OverrideSymbolicColor is a wrapper around the C function gtk_widget_override_symbolic_color.
func (recv *Widget) OverrideSymbolicColor(name string, color *gdk.RGBA) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	C.gtk_widget_override_symbolic_color((*C.GtkWidget)(recv.native), c_name, c_color)

	return
}

// QueueDrawRegion is a wrapper around the C function gtk_widget_queue_draw_region.
func (recv *Widget) QueueDrawRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_queue_draw_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// RenderIconPixbuf is a wrapper around the C function gtk_widget_render_icon_pixbuf.
func (recv *Widget) RenderIconPixbuf(stockId string, size IconSize) *gdkpixbuf.Pixbuf {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_widget_render_icon_pixbuf((*C.GtkWidget)(recv.native), c_stock_id, c_size)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ResetStyle is a wrapper around the C function gtk_widget_reset_style.
func (recv *Widget) ResetStyle() {
	C.gtk_widget_reset_style((*C.GtkWidget)(recv.native))

	return
}

// SetDeviceEnabled is a wrapper around the C function gtk_widget_set_device_enabled.
func (recv *Widget) SetDeviceEnabled(device *gdk.Device, enabled bool) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_enabled :=
		boolToGboolean(enabled)

	C.gtk_widget_set_device_enabled((*C.GtkWidget)(recv.native), c_device, c_enabled)

	return
}

// SetDeviceEvents is a wrapper around the C function gtk_widget_set_device_events.
func (recv *Widget) SetDeviceEvents(device *gdk.Device, events gdk.EventMask) {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_events := (C.GdkEventMask)(events)

	C.gtk_widget_set_device_events((*C.GtkWidget)(recv.native), c_device, c_events)

	return
}

// SetMarginBottom is a wrapper around the C function gtk_widget_set_margin_bottom.
func (recv *Widget) SetMarginBottom(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_bottom((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginLeft is a wrapper around the C function gtk_widget_set_margin_left.
func (recv *Widget) SetMarginLeft(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_left((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginRight is a wrapper around the C function gtk_widget_set_margin_right.
func (recv *Widget) SetMarginRight(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_right((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetMarginTop is a wrapper around the C function gtk_widget_set_margin_top.
func (recv *Widget) SetMarginTop(margin int32) {
	c_margin := (C.gint)(margin)

	C.gtk_widget_set_margin_top((*C.GtkWidget)(recv.native), c_margin)

	return
}

// SetStateFlags is a wrapper around the C function gtk_widget_set_state_flags.
func (recv *Widget) SetStateFlags(flags StateFlags, clear bool) {
	c_flags := (C.GtkStateFlags)(flags)

	c_clear :=
		boolToGboolean(clear)

	C.gtk_widget_set_state_flags((*C.GtkWidget)(recv.native), c_flags, c_clear)

	return
}

// SetSupportMultidevice is a wrapper around the C function gtk_widget_set_support_multidevice.
func (recv *Widget) SetSupportMultidevice(supportMultidevice bool) {
	c_support_multidevice :=
		boolToGboolean(supportMultidevice)

	C.gtk_widget_set_support_multidevice((*C.GtkWidget)(recv.native), c_support_multidevice)

	return
}

// ShapeCombineRegion is a wrapper around the C function gtk_widget_shape_combine_region.
func (recv *Widget) ShapeCombineRegion(region *cairo.Region) {
	c_region := (*C.cairo_region_t)(C.NULL)
	if region != nil {
		c_region = (*C.cairo_region_t)(region.ToC())
	}

	C.gtk_widget_shape_combine_region((*C.GtkWidget)(recv.native), c_region)

	return
}

// UnsetStateFlags is a wrapper around the C function gtk_widget_unset_state_flags.
func (recv *Widget) UnsetStateFlags(flags StateFlags) {
	c_flags := (C.GtkStateFlags)(flags)

	C.gtk_widget_unset_state_flags((*C.GtkWidget)(recv.native), c_flags)

	return
}

// GetApplication is a wrapper around the C function gtk_window_get_application.
func (recv *Window) GetApplication() *Application {
	retC := C.gtk_window_get_application((*C.GtkWindow)(recv.native))
	var retGo (*Application)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ApplicationNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetHasResizeGrip is a wrapper around the C function gtk_window_get_has_resize_grip.
func (recv *Window) GetHasResizeGrip() bool {
	retC := C.gtk_window_get_has_resize_grip((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetResizeGripArea is a wrapper around the C function gtk_window_get_resize_grip_area.
func (recv *Window) GetResizeGripArea() (bool, *gdk.Rectangle) {
	var c_rect C.GdkRectangle

	retC := C.gtk_window_get_resize_grip_area((*C.GtkWindow)(recv.native), &c_rect)
	retGo := retC == C.TRUE

	rect := gdk.RectangleNewFromC(unsafe.Pointer(&c_rect))

	return retGo, rect
}

// ResizeGripIsVisible is a wrapper around the C function gtk_window_resize_grip_is_visible.
func (recv *Window) ResizeGripIsVisible() bool {
	retC := C.gtk_window_resize_grip_is_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ResizeToGeometry is a wrapper around the C function gtk_window_resize_to_geometry.
func (recv *Window) ResizeToGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_resize_to_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetApplication is a wrapper around the C function gtk_window_set_application.
func (recv *Window) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	C.gtk_window_set_application((*C.GtkWindow)(recv.native), c_application)

	return
}

// SetDefaultGeometry is a wrapper around the C function gtk_window_set_default_geometry.
func (recv *Window) SetDefaultGeometry(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_set_default_geometry((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetHasResizeGrip is a wrapper around the C function gtk_window_set_has_resize_grip.
func (recv *Window) SetHasResizeGrip(value bool) {
	c_value :=
		boolToGboolean(value)

	C.gtk_window_set_has_resize_grip((*C.GtkWindow)(recv.native), c_value)

	return
}

// SetHasUserRefCount is a wrapper around the C function gtk_window_set_has_user_ref_count.
func (recv *Window) SetHasUserRefCount(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_has_user_ref_count((*C.GtkWindow)(recv.native), c_setting)

	return
}

// GetCurrentDeviceGrab is a wrapper around the C function gtk_window_group_get_current_device_grab.
func (recv *WindowGroup) GetCurrentDeviceGrab(device *gdk.Device) *Widget {
	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	retC := C.gtk_window_group_get_current_device_grab((*C.GtkWindowGroup)(recv.native), c_device)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type License C.GtkLicense

const (
	GTK_LICENSE_UNKNOWN       License = 0
	GTK_LICENSE_CUSTOM        License = 1
	GTK_LICENSE_GPL_2_0       License = 2
	GTK_LICENSE_GPL_3_0       License = 3
	GTK_LICENSE_LGPL_2_1      License = 4
	GTK_LICENSE_LGPL_3_0      License = 5
	GTK_LICENSE_BSD           License = 6
	GTK_LICENSE_MIT_X11       License = 7
	GTK_LICENSE_ARTISTIC      License = 8
	GTK_LICENSE_GPL_2_0_ONLY  License = 9
	GTK_LICENSE_GPL_3_0_ONLY  License = 10
	GTK_LICENSE_LGPL_2_1_ONLY License = 11
	GTK_LICENSE_LGPL_3_0_ONLY License = 12
	GTK_LICENSE_AGPL_3_0      License = 13
	GTK_LICENSE_AGPL_3_0_ONLY License = 14
)

// CairoShouldDrawWindow is a wrapper around the C function gtk_cairo_should_draw_window.
func CairoShouldDrawWindow(cr *cairo.Context, window *gdk.Window) bool {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	retC := C.gtk_cairo_should_draw_window(c_cr, c_window)
	retGo := retC == C.TRUE

	return retGo
}

// CairoTransformToWindow is a wrapper around the C function gtk_cairo_transform_to_window.
func CairoTransformToWindow(cr *cairo.Context, widget *Widget, window *gdk.Window) {
	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_window := (*C.GdkWindow)(C.NULL)
	if window != nil {
		c_window = (*C.GdkWindow)(window.ToC())
	}

	C.gtk_cairo_transform_to_window(c_cr, c_widget, c_window)

	return
}

// DeviceGrabAdd is a wrapper around the C function gtk_device_grab_add.
func DeviceGrabAdd(widget *Widget, device *gdk.Device, blockOthers bool) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	c_block_others :=
		boolToGboolean(blockOthers)

	C.gtk_device_grab_add(c_widget, c_device, c_block_others)

	return
}

// DeviceGrabRemove is a wrapper around the C function gtk_device_grab_remove.
func DeviceGrabRemove(widget *Widget, device *gdk.Device) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_device := (*C.GdkDevice)(C.NULL)
	if device != nil {
		c_device = (*C.GdkDevice)(device.ToC())
	}

	C.gtk_device_grab_remove(c_widget, c_device)

	return
}

// DrawInsertionCursor is a wrapper around the C function gtk_draw_insertion_cursor.
func DrawInsertionCursor(widget *Widget, cr *cairo.Context, location *gdk.Rectangle, isPrimary bool, direction TextDirection, drawArrow bool) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_location := (*C.GdkRectangle)(C.NULL)
	if location != nil {
		c_location = (*C.GdkRectangle)(location.ToC())
	}

	c_is_primary :=
		boolToGboolean(isPrimary)

	c_direction := (C.GtkTextDirection)(direction)

	c_draw_arrow :=
		boolToGboolean(drawArrow)

	C.gtk_draw_insertion_cursor(c_widget, c_cr, c_location, c_is_primary, c_direction, c_draw_arrow)

	return
}

// GetBinaryAge is a wrapper around the C function gtk_get_binary_age.
func GetBinaryAge() uint32 {
	retC := C.gtk_get_binary_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetInterfaceAge is a wrapper around the C function gtk_get_interface_age.
func GetInterfaceAge() uint32 {
	retC := C.gtk_get_interface_age()
	retGo := (uint32)(retC)

	return retGo
}

// GetMajorVersion is a wrapper around the C function gtk_get_major_version.
func GetMajorVersion() uint32 {
	retC := C.gtk_get_major_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMicroVersion is a wrapper around the C function gtk_get_micro_version.
func GetMicroVersion() uint32 {
	retC := C.gtk_get_micro_version()
	retGo := (uint32)(retC)

	return retGo
}

// GetMinorVersion is a wrapper around the C function gtk_get_minor_version.
func GetMinorVersion() uint32 {
	retC := C.gtk_get_minor_version()
	retGo := (uint32)(retC)

	return retGo
}

// RenderActivity is a wrapper around the C function gtk_render_activity.
func RenderActivity(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_activity(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderArrow is a wrapper around the C function gtk_render_arrow.
func RenderArrow(context *StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_angle := (C.gdouble)(angle)

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_size := (C.gdouble)(size)

	C.gtk_render_arrow(c_context, c_cr, c_angle, c_x, c_y, c_size)

	return
}

// RenderBackground is a wrapper around the C function gtk_render_background.
func RenderBackground(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_background(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderCheck is a wrapper around the C function gtk_render_check.
func RenderCheck(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_check(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderExpander is a wrapper around the C function gtk_render_expander.
func RenderExpander(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_expander(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderExtension is a wrapper around the C function gtk_render_extension.
func RenderExtension(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	C.gtk_render_extension(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side)

	return
}

// RenderFocus is a wrapper around the C function gtk_render_focus.
func RenderFocus(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_focus(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderFrame is a wrapper around the C function gtk_render_frame.
func RenderFrame(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_frame(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderFrameGap is a wrapper around the C function gtk_render_frame_gap.
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_gap_side := (C.GtkPositionType)(gapSide)

	c_xy0_gap := (C.gdouble)(xy0Gap)

	c_xy1_gap := (C.gdouble)(xy1Gap)

	C.gtk_render_frame_gap(c_context, c_cr, c_x, c_y, c_width, c_height, c_gap_side, c_xy0_gap, c_xy1_gap)

	return
}

// RenderHandle is a wrapper around the C function gtk_render_handle.
func RenderHandle(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_handle(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderIconPixbuf is a wrapper around the C function gtk_render_icon_pixbuf.
func RenderIconPixbuf(context *StyleContext, source *IconSource, size IconSize) *gdkpixbuf.Pixbuf {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_source := (*C.GtkIconSource)(C.NULL)
	if source != nil {
		c_source = (*C.GtkIconSource)(source.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_render_icon_pixbuf(c_context, c_source, c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RenderLayout is a wrapper around the C function gtk_render_layout.
func RenderLayout(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_layout := (*C.PangoLayout)(C.NULL)
	if layout != nil {
		c_layout = (*C.PangoLayout)(layout.ToC())
	}

	C.gtk_render_layout(c_context, c_cr, c_x, c_y, c_layout)

	return
}

// RenderLine is a wrapper around the C function gtk_render_line.
func RenderLine(context *StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	C.gtk_render_line(c_context, c_cr, c_x0, c_y0, c_x1, c_y1)

	return
}

// RenderOption is a wrapper around the C function gtk_render_option.
func RenderOption(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	C.gtk_render_option(c_context, c_cr, c_x, c_y, c_width, c_height)

	return
}

// RenderSlider is a wrapper around the C function gtk_render_slider.
func RenderSlider(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	c_width := (C.gdouble)(width)

	c_height := (C.gdouble)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_render_slider(c_context, c_cr, c_x, c_y, c_width, c_height, c_orientation)

	return
}

// GetAppInfo is a wrapper around the C function gtk_app_chooser_get_app_info.
func (recv *AppChooser) GetAppInfo() *gio.AppInfo {
	retC := C.gtk_app_chooser_get_app_info((*C.GtkAppChooser)(recv.native))
	var retGo (*gio.AppInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.AppInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetContentType is a wrapper around the C function gtk_app_chooser_get_content_type.
func (recv *AppChooser) GetContentType() string {
	retC := C.gtk_app_chooser_get_content_type((*C.GtkAppChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Refresh is a wrapper around the C function gtk_app_chooser_refresh.
func (recv *AppChooser) Refresh() {
	C.gtk_app_chooser_refresh((*C.GtkAppChooser)(recv.native))

	return
}

// GetArea is a wrapper around the C function gtk_cell_layout_get_area.
func (recv *CellLayout) GetArea() *CellArea {
	retC := C.gtk_cell_layout_get_area((*C.GtkCellLayout)(recv.native))
	var retGo (*CellArea)
	if retC == nil {
		retGo = nil
	} else {
		retGo = CellAreaNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_scrollable_get_hadjustment.
func (recv *Scrollable) GetHadjustment() *Adjustment {
	retC := C.gtk_scrollable_get_hadjustment((*C.GtkScrollable)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHscrollPolicy is a wrapper around the C function gtk_scrollable_get_hscroll_policy.
func (recv *Scrollable) GetHscrollPolicy() ScrollablePolicy {
	retC := C.gtk_scrollable_get_hscroll_policy((*C.GtkScrollable)(recv.native))
	retGo := (ScrollablePolicy)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_scrollable_get_vadjustment.
func (recv *Scrollable) GetVadjustment() *Adjustment {
	retC := C.gtk_scrollable_get_vadjustment((*C.GtkScrollable)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVscrollPolicy is a wrapper around the C function gtk_scrollable_get_vscroll_policy.
func (recv *Scrollable) GetVscrollPolicy() ScrollablePolicy {
	retC := C.gtk_scrollable_get_vscroll_policy((*C.GtkScrollable)(recv.native))
	retGo := (ScrollablePolicy)(retC)

	return retGo
}

// SetHadjustment is a wrapper around the C function gtk_scrollable_set_hadjustment.
func (recv *Scrollable) SetHadjustment(hadjustment *Adjustment) {
	c_hadjustment := (*C.GtkAdjustment)(C.NULL)
	if hadjustment != nil {
		c_hadjustment = (*C.GtkAdjustment)(hadjustment.ToC())
	}

	C.gtk_scrollable_set_hadjustment((*C.GtkScrollable)(recv.native), c_hadjustment)

	return
}

// SetHscrollPolicy is a wrapper around the C function gtk_scrollable_set_hscroll_policy.
func (recv *Scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	c_policy := (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_hscroll_policy((*C.GtkScrollable)(recv.native), c_policy)

	return
}

// SetVadjustment is a wrapper around the C function gtk_scrollable_set_vadjustment.
func (recv *Scrollable) SetVadjustment(vadjustment *Adjustment) {
	c_vadjustment := (*C.GtkAdjustment)(C.NULL)
	if vadjustment != nil {
		c_vadjustment = (*C.GtkAdjustment)(vadjustment.ToC())
	}

	C.gtk_scrollable_set_vadjustment((*C.GtkScrollable)(recv.native), c_vadjustment)

	return
}

// SetVscrollPolicy is a wrapper around the C function gtk_scrollable_set_vscroll_policy.
func (recv *Scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	c_policy := (C.GtkScrollablePolicy)(policy)

	C.gtk_scrollable_set_vscroll_policy((*C.GtkScrollable)(recv.native), c_policy)

	return
}

// GetIconFactory is a wrapper around the C function gtk_style_provider_get_icon_factory.
func (recv *StyleProvider) GetIconFactory(path *WidgetPath) *IconFactory {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	retC := C.gtk_style_provider_get_icon_factory((*C.GtkStyleProvider)(recv.native), c_path)
	var retGo (*IconFactory)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconFactoryNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStyle is a wrapper around the C function gtk_style_provider_get_style.
func (recv *StyleProvider) GetStyle(path *WidgetPath) *StyleProperties {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	retC := C.gtk_style_provider_get_style((*C.GtkStyleProvider)(recv.native), c_path)
	var retGo (*StyleProperties)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StylePropertiesNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStyleProperty is a wrapper around the C function gtk_style_provider_get_style_property.
func (recv *StyleProvider) GetStyleProperty(path *WidgetPath, state StateFlags, pspec *gobject.ParamSpec) (bool, *gobject.Value) {
	c_path := (*C.GtkWidgetPath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkWidgetPath)(path.ToC())
	}

	c_state := (C.GtkStateFlags)(state)

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	var c_value C.GValue

	retC := C.gtk_style_provider_get_style_property((*C.GtkStyleProvider)(recv.native), c_path, c_state, c_pspec, &c_value)
	retGo := retC == C.TRUE

	value := gobject.ValueNewFromC(unsafe.Pointer(&c_value))

	return retGo, value
}

// IterPrevious is a wrapper around the C function gtk_tree_model_iter_previous.
func (recv *TreeModel) IterPrevious(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	retC := C.gtk_tree_model_iter_previous((*C.GtkTreeModel)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// BindingEntryAddSignalFromString is a wrapper around the C function gtk_binding_entry_add_signal_from_string.
func BindingEntryAddSignalFromString(bindingSet *BindingSet, signalDesc string) glib.TokenType {
	c_binding_set := (*C.GtkBindingSet)(C.NULL)
	if bindingSet != nil {
		c_binding_set = (*C.GtkBindingSet)(bindingSet.ToC())
	}

	c_signal_desc := C.CString(signalDesc)
	defer C.free(unsafe.Pointer(c_signal_desc))

	retC := C.gtk_binding_entry_add_signal_from_string(c_binding_set, c_signal_desc)
	retGo := (glib.TokenType)(retC)

	return retGo
}

// FindCellProperty is a wrapper around the C function gtk_cell_area_class_find_cell_property.
func (recv *CellAreaClass) FindCellProperty(propertyName string) *gobject.ParamSpec {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	retC := C.gtk_cell_area_class_find_cell_property((*C.GtkCellAreaClass)(recv.native), c_property_name)
	retGo := gobject.ParamSpecNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InstallCellProperty is a wrapper around the C function gtk_cell_area_class_install_cell_property.
func (recv *CellAreaClass) InstallCellProperty(propertyId uint32, pspec *gobject.ParamSpec) {
	c_property_id := (C.guint)(propertyId)

	c_pspec := (*C.GParamSpec)(C.NULL)
	if pspec != nil {
		c_pspec = (*C.GParamSpec)(pspec.ToC())
	}

	C.gtk_cell_area_class_install_cell_property((*C.GtkCellAreaClass)(recv.native), c_property_id, c_pspec)

	return
}

// Unsupported : gtk_cell_area_class_list_cell_properties : array return type :

// GradientNewLinear is a wrapper around the C function gtk_gradient_new_linear.
func GradientNewLinear(x0 float64, y0 float64, x1 float64, y1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	retC := C.gtk_gradient_new_linear(c_x0, c_y0, c_x1, c_y1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GradientNewRadial is a wrapper around the C function gtk_gradient_new_radial.
func GradientNewRadial(x0 float64, y0 float64, radius0 float64, x1 float64, y1 float64, radius1 float64) *Gradient {
	c_x0 := (C.gdouble)(x0)

	c_y0 := (C.gdouble)(y0)

	c_radius0 := (C.gdouble)(radius0)

	c_x1 := (C.gdouble)(x1)

	c_y1 := (C.gdouble)(y1)

	c_radius1 := (C.gdouble)(radius1)

	retC := C.gtk_gradient_new_radial(c_x0, c_y0, c_radius0, c_x1, c_y1, c_radius1)
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddColorStop is a wrapper around the C function gtk_gradient_add_color_stop.
func (recv *Gradient) AddColorStop(offset float64, color *SymbolicColor) {
	c_offset := (C.gdouble)(offset)

	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	C.gtk_gradient_add_color_stop((*C.GtkGradient)(recv.native), c_offset, c_color)

	return
}

// Ref is a wrapper around the C function gtk_gradient_ref.
func (recv *Gradient) Ref() *Gradient {
	retC := C.gtk_gradient_ref((*C.GtkGradient)(recv.native))
	retGo := GradientNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Resolve is a wrapper around the C function gtk_gradient_resolve.
func (recv *Gradient) Resolve(props *StyleProperties) (bool, *cairo.Pattern) {
	c_props := (*C.GtkStyleProperties)(C.NULL)
	if props != nil {
		c_props = (*C.GtkStyleProperties)(props.ToC())
	}

	var c_resolved_gradient *C.cairo_pattern_t

	retC := C.gtk_gradient_resolve((*C.GtkGradient)(recv.native), c_props, &c_resolved_gradient)
	retGo := retC == C.TRUE

	resolvedGradient := cairo.PatternNewFromC(unsafe.Pointer(c_resolved_gradient))

	return retGo, resolvedGradient
}

// Unref is a wrapper around the C function gtk_gradient_unref.
func (recv *Gradient) Unref() {
	C.gtk_gradient_unref((*C.GtkGradient)(recv.native))

	return
}

// RenderIconPixbuf is a wrapper around the C function gtk_icon_set_render_icon_pixbuf.
func (recv *IconSet) RenderIconPixbuf(context *StyleContext, size IconSize) *gdkpixbuf.Pixbuf {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_icon_set_render_icon_pixbuf((*C.GtkIconSet)(recv.native), c_context, c_size)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RequisitionNew is a wrapper around the C function gtk_requisition_new.
func RequisitionNew() *Requisition {
	retC := C.gtk_requisition_new()
	retGo := RequisitionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_selection_data_get_data_with_length : array return type :

// SymbolicColorNewAlpha is a wrapper around the C function gtk_symbolic_color_new_alpha.
func SymbolicColorNewAlpha(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_alpha(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewLiteral is a wrapper around the C function gtk_symbolic_color_new_literal.
func SymbolicColorNewLiteral(color *gdk.RGBA) *SymbolicColor {
	c_color := (*C.GdkRGBA)(C.NULL)
	if color != nil {
		c_color = (*C.GdkRGBA)(color.ToC())
	}

	retC := C.gtk_symbolic_color_new_literal(c_color)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewMix is a wrapper around the C function gtk_symbolic_color_new_mix.
func SymbolicColorNewMix(color1 *SymbolicColor, color2 *SymbolicColor, factor float64) *SymbolicColor {
	c_color1 := (*C.GtkSymbolicColor)(C.NULL)
	if color1 != nil {
		c_color1 = (*C.GtkSymbolicColor)(color1.ToC())
	}

	c_color2 := (*C.GtkSymbolicColor)(C.NULL)
	if color2 != nil {
		c_color2 = (*C.GtkSymbolicColor)(color2.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_mix(c_color1, c_color2, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewName is a wrapper around the C function gtk_symbolic_color_new_name.
func SymbolicColorNewName(name string) *SymbolicColor {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_symbolic_color_new_name(c_name)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SymbolicColorNewShade is a wrapper around the C function gtk_symbolic_color_new_shade.
func SymbolicColorNewShade(color *SymbolicColor, factor float64) *SymbolicColor {
	c_color := (*C.GtkSymbolicColor)(C.NULL)
	if color != nil {
		c_color = (*C.GtkSymbolicColor)(color.ToC())
	}

	c_factor := (C.gdouble)(factor)

	retC := C.gtk_symbolic_color_new_shade(c_color, c_factor)
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Ref is a wrapper around the C function gtk_symbolic_color_ref.
func (recv *SymbolicColor) Ref() *SymbolicColor {
	retC := C.gtk_symbolic_color_ref((*C.GtkSymbolicColor)(recv.native))
	retGo := SymbolicColorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Resolve is a wrapper around the C function gtk_symbolic_color_resolve.
func (recv *SymbolicColor) Resolve(props *StyleProperties) (bool, *gdk.RGBA) {
	c_props := (*C.GtkStyleProperties)(C.NULL)
	if props != nil {
		c_props = (*C.GtkStyleProperties)(props.ToC())
	}

	var c_resolved_color C.GdkRGBA

	retC := C.gtk_symbolic_color_resolve((*C.GtkSymbolicColor)(recv.native), c_props, &c_resolved_color)
	retGo := retC == C.TRUE

	resolvedColor := gdk.RGBANewFromC(unsafe.Pointer(&c_resolved_color))

	return retGo, resolvedColor
}

// Unref is a wrapper around the C function gtk_symbolic_color_unref.
func (recv *SymbolicColor) Unref() {
	C.gtk_symbolic_color_unref((*C.GtkSymbolicColor)(recv.native))

	return
}

// Unsupported : gtk_tree_path_get_indices_with_depth : array return type :

// WidgetPathNew is a wrapper around the C function gtk_widget_path_new.
func WidgetPathNew() *WidgetPath {
	retC := C.gtk_widget_path_new()
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendType is a wrapper around the C function gtk_widget_path_append_type.
func (recv *WidgetPath) AppendType(type_ gobject.Type) int32 {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_append_type((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := (int32)(retC)

	return retGo
}

// Copy is a wrapper around the C function gtk_widget_path_copy.
func (recv *WidgetPath) Copy() *WidgetPath {
	retC := C.gtk_widget_path_copy((*C.GtkWidgetPath)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_widget_path_free.
func (recv *WidgetPath) Free() {
	C.gtk_widget_path_free((*C.GtkWidgetPath)(recv.native))

	return
}

// GetObjectType is a wrapper around the C function gtk_widget_path_get_object_type.
func (recv *WidgetPath) GetObjectType() gobject.Type {
	retC := C.gtk_widget_path_get_object_type((*C.GtkWidgetPath)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// HasParent is a wrapper around the C function gtk_widget_path_has_parent.
func (recv *WidgetPath) HasParent(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_has_parent((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// IsType is a wrapper around the C function gtk_widget_path_is_type.
func (recv *WidgetPath) IsType(type_ gobject.Type) bool {
	c_type := (C.GType)(type_)

	retC := C.gtk_widget_path_is_type((*C.GtkWidgetPath)(recv.native), c_type)
	retGo := retC == C.TRUE

	return retGo
}

// IterAddClass is a wrapper around the C function gtk_widget_path_iter_add_class.
func (recv *WidgetPath) IterAddClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_add_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterAddRegion is a wrapper around the C function gtk_widget_path_iter_add_region.
func (recv *WidgetPath) IterAddRegion(pos int32, name string, flags RegionFlags) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_flags := (C.GtkRegionFlags)(flags)

	C.gtk_widget_path_iter_add_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name, c_flags)

	return
}

// IterClearClasses is a wrapper around the C function gtk_widget_path_iter_clear_classes.
func (recv *WidgetPath) IterClearClasses(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_classes((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// IterClearRegions is a wrapper around the C function gtk_widget_path_iter_clear_regions.
func (recv *WidgetPath) IterClearRegions(pos int32) {
	c_pos := (C.gint)(pos)

	C.gtk_widget_path_iter_clear_regions((*C.GtkWidgetPath)(recv.native), c_pos)

	return
}

// IterGetObjectType is a wrapper around the C function gtk_widget_path_iter_get_object_type.
func (recv *WidgetPath) IterGetObjectType(pos int32) gobject.Type {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_get_object_type((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := (gobject.Type)(retC)

	return retGo
}

// IterHasClass is a wrapper around the C function gtk_widget_path_iter_has_class.
func (recv *WidgetPath) IterHasClass(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasName is a wrapper around the C function gtk_widget_path_iter_has_name.
func (recv *WidgetPath) IterHasName(pos int32, name string) bool {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_widget_path_iter_has_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasQclass is a wrapper around the C function gtk_widget_path_iter_has_qclass.
func (recv *WidgetPath) IterHasQclass(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qclass((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// IterHasQname is a wrapper around the C function gtk_widget_path_iter_has_qname.
func (recv *WidgetPath) IterHasQname(pos int32, qname glib.Quark) bool {
	c_pos := (C.gint)(pos)

	c_qname := (C.GQuark)(qname)

	retC := C.gtk_widget_path_iter_has_qname((*C.GtkWidgetPath)(recv.native), c_pos, c_qname)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_path_iter_has_qregion : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// Unsupported : gtk_widget_path_iter_has_region : unsupported parameter flags : GtkRegionFlags* with indirection level of 1

// IterListClasses is a wrapper around the C function gtk_widget_path_iter_list_classes.
func (recv *WidgetPath) IterListClasses(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_classes((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterListRegions is a wrapper around the C function gtk_widget_path_iter_list_regions.
func (recv *WidgetPath) IterListRegions(pos int32) *glib.SList {
	c_pos := (C.gint)(pos)

	retC := C.gtk_widget_path_iter_list_regions((*C.GtkWidgetPath)(recv.native), c_pos)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IterRemoveClass is a wrapper around the C function gtk_widget_path_iter_remove_class.
func (recv *WidgetPath) IterRemoveClass(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_class((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterRemoveRegion is a wrapper around the C function gtk_widget_path_iter_remove_region.
func (recv *WidgetPath) IterRemoveRegion(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_remove_region((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterSetName is a wrapper around the C function gtk_widget_path_iter_set_name.
func (recv *WidgetPath) IterSetName(pos int32, name string) {
	c_pos := (C.gint)(pos)

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_path_iter_set_name((*C.GtkWidgetPath)(recv.native), c_pos, c_name)

	return
}

// IterSetObjectType is a wrapper around the C function gtk_widget_path_iter_set_object_type.
func (recv *WidgetPath) IterSetObjectType(pos int32, type_ gobject.Type) {
	c_pos := (C.gint)(pos)

	c_type := (C.GType)(type_)

	C.gtk_widget_path_iter_set_object_type((*C.GtkWidgetPath)(recv.native), c_pos, c_type)

	return
}

// Length is a wrapper around the C function gtk_widget_path_length.
func (recv *WidgetPath) Length() int32 {
	retC := C.gtk_widget_path_length((*C.GtkWidgetPath)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// PrependType is a wrapper around the C function gtk_widget_path_prepend_type.
func (recv *WidgetPath) PrependType(type_ gobject.Type) {
	c_type := (C.GType)(type_)

	C.gtk_widget_path_prepend_type((*C.GtkWidgetPath)(recv.native), c_type)

	return
}