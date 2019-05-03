// Code generated - DO NOT EDIT.
// +build gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	cairo "github.com/pekim/gobbi/lib/cairo"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	gio "github.com/pekim/gobbi/lib/gio"
	glib "github.com/pekim/gobbi/lib/glib"
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

	void listbox_rowActivatedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-activated", G_CALLBACK(listbox_rowActivatedHandler), data);
	}

*/
/*

	void listbox_rowSelectedHandler(GObject *, GtkListBoxRow *, gpointer);

	static gulong ListBox_signal_connect_row_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-selected", G_CALLBACK(listbox_rowSelectedHandler), data);
	}

*/
/*

	void listboxrow_activateHandler(GObject *, gpointer);

	static gulong ListBoxRow_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(listboxrow_activateHandler), data);
	}

*/
/*

	gint placessidebar_dragActionAskHandler(GObject *, gint, gpointer);

	static gulong PlacesSidebar_signal_connect_drag_action_ask(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-action-ask", G_CALLBACK(placessidebar_dragActionAskHandler), data);
	}

*/
/*

	void placessidebar_openLocationHandler(GObject *, GFile *, GtkPlacesOpenFlags, gpointer);

	static gulong PlacesSidebar_signal_connect_open_location(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "open-location", G_CALLBACK(placessidebar_openLocationHandler), data);
	}

*/
/*

	void placessidebar_populatePopupHandler(GObject *, GtkWidget *, GFile *, GVolume *, gpointer);

	static gulong PlacesSidebar_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", G_CALLBACK(placessidebar_populatePopupHandler), data);
	}

*/
/*

	void placessidebar_showErrorMessageHandler(GObject *, gchar*, gchar*, gpointer);

	static gulong PlacesSidebar_signal_connect_show_error_message(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-error-message", G_CALLBACK(placessidebar_showErrorMessageHandler), data);
	}

*/
/*

	void searchentry_searchChangedHandler(GObject *, gpointer);

	static gulong SearchEntry_signal_connect_search_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search-changed", G_CALLBACK(searchentry_searchChangedHandler), data);
	}

*/
import "C"

// GetBaselinePosition is a wrapper around the C function gtk_box_get_baseline_position.
func (recv *Box) GetBaselinePosition() BaselinePosition {
	retC := C.gtk_box_get_baseline_position((*C.GtkBox)(recv.native))
	retGo := (BaselinePosition)(retC)

	return retGo
}

// SetBaselinePosition is a wrapper around the C function gtk_box_set_baseline_position.
func (recv *Box) SetBaselinePosition(position BaselinePosition) {
	c_position := (C.GtkBaselinePosition)(position)

	C.gtk_box_set_baseline_position((*C.GtkBox)(recv.native), c_position)

	return
}

// BuilderNewFromFile is a wrapper around the C function gtk_builder_new_from_file.
func BuilderNewFromFile(filename string) *Builder {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_builder_new_from_file(c_filename)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BuilderNewFromResource is a wrapper around the C function gtk_builder_new_from_resource.
func BuilderNewFromResource(resourcePath string) *Builder {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	retC := C.gtk_builder_new_from_resource(c_resource_path)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// BuilderNewFromString is a wrapper around the C function gtk_builder_new_from_string.
func BuilderNewFromString(string_ string) *Builder {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	c_length := (C.gssize)(len(string_))

	retC := C.gtk_builder_new_from_string(c_string, c_length)
	retGo := BuilderNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Unsupported : gtk_builder_add_callback_symbol : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// Unsupported : gtk_builder_add_callback_symbols : unsupported parameter first_callback_symbol : no type generator for GObject.Callback (GCallback) for param first_callback_symbol

// GetApplication is a wrapper around the C function gtk_builder_get_application.
func (recv *Builder) GetApplication() *Application {
	retC := C.gtk_builder_get_application((*C.GtkBuilder)(recv.native))
	var retGo (*Application)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ApplicationNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_builder_lookup_callback_symbol : no return generator

// SetApplication is a wrapper around the C function gtk_builder_set_application.
func (recv *Builder) SetApplication(application *Application) {
	c_application := (*C.GtkApplication)(C.NULL)
	if application != nil {
		c_application = (*C.GtkApplication)(application.ToC())
	}

	C.gtk_builder_set_application((*C.GtkBuilder)(recv.native), c_application)

	return
}

// ButtonNewFromIconName is a wrapper around the C function gtk_button_new_from_icon_name.
func ButtonNewFromIconName(iconName string, size IconSize) *Button {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_button_new_from_icon_name(c_icon_name, c_size)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTabs is a wrapper around the C function gtk_entry_get_tabs.
func (recv *Entry) GetTabs() *pango.TabArray {
	retC := C.gtk_entry_get_tabs((*C.GtkEntry)(recv.native))
	var retGo (*pango.TabArray)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.TabArrayNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetTabs is a wrapper around the C function gtk_entry_set_tabs.
func (recv *Entry) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(C.NULL)
	if tabs != nil {
		c_tabs = (*C.PangoTabArray)(tabs.ToC())
	}

	C.gtk_entry_set_tabs((*C.GtkEntry)(recv.native), c_tabs)

	return
}

// GetBaselineRow is a wrapper around the C function gtk_grid_get_baseline_row.
func (recv *Grid) GetBaselineRow() int32 {
	retC := C.gtk_grid_get_baseline_row((*C.GtkGrid)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRowBaselinePosition is a wrapper around the C function gtk_grid_get_row_baseline_position.
func (recv *Grid) GetRowBaselinePosition(row int32) BaselinePosition {
	c_row := (C.gint)(row)

	retC := C.gtk_grid_get_row_baseline_position((*C.GtkGrid)(recv.native), c_row)
	retGo := (BaselinePosition)(retC)

	return retGo
}

// RemoveColumn is a wrapper around the C function gtk_grid_remove_column.
func (recv *Grid) RemoveColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// RemoveRow is a wrapper around the C function gtk_grid_remove_row.
func (recv *Grid) RemoveRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_remove_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// SetBaselineRow is a wrapper around the C function gtk_grid_set_baseline_row.
func (recv *Grid) SetBaselineRow(row int32) {
	c_row := (C.gint)(row)

	C.gtk_grid_set_baseline_row((*C.GtkGrid)(recv.native), c_row)

	return
}

// SetRowBaselinePosition is a wrapper around the C function gtk_grid_set_row_baseline_position.
func (recv *Grid) SetRowBaselinePosition(row int32, pos BaselinePosition) {
	c_row := (C.gint)(row)

	c_pos := (C.GtkBaselinePosition)(pos)

	C.gtk_grid_set_row_baseline_position((*C.GtkGrid)(recv.native), c_row, c_pos)

	return
}

// HeaderBarNew is a wrapper around the C function gtk_header_bar_new.
func HeaderBarNew() *HeaderBar {
	retC := C.gtk_header_bar_new()
	retGo := HeaderBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCustomTitle is a wrapper around the C function gtk_header_bar_get_custom_title.
func (recv *HeaderBar) GetCustomTitle() *Widget {
	retC := C.gtk_header_bar_get_custom_title((*C.GtkHeaderBar)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_header_bar_get_show_close_button.
func (recv *HeaderBar) GetShowCloseButton() bool {
	retC := C.gtk_header_bar_get_show_close_button((*C.GtkHeaderBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSubtitle is a wrapper around the C function gtk_header_bar_get_subtitle.
func (recv *HeaderBar) GetSubtitle() string {
	retC := C.gtk_header_bar_get_subtitle((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTitle is a wrapper around the C function gtk_header_bar_get_title.
func (recv *HeaderBar) GetTitle() string {
	retC := C.gtk_header_bar_get_title((*C.GtkHeaderBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_header_bar_pack_end.
func (recv *HeaderBar) PackEnd(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_header_bar_pack_end((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// PackStart is a wrapper around the C function gtk_header_bar_pack_start.
func (recv *HeaderBar) PackStart(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_header_bar_pack_start((*C.GtkHeaderBar)(recv.native), c_child)

	return
}

// SetCustomTitle is a wrapper around the C function gtk_header_bar_set_custom_title.
func (recv *HeaderBar) SetCustomTitle(titleWidget *Widget) {
	c_title_widget := (*C.GtkWidget)(C.NULL)
	if titleWidget != nil {
		c_title_widget = (*C.GtkWidget)(titleWidget.ToC())
	}

	C.gtk_header_bar_set_custom_title((*C.GtkHeaderBar)(recv.native), c_title_widget)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_header_bar_set_show_close_button.
func (recv *HeaderBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_header_bar_set_show_close_button((*C.GtkHeaderBar)(recv.native), c_setting)

	return
}

// SetSubtitle is a wrapper around the C function gtk_header_bar_set_subtitle.
func (recv *HeaderBar) SetSubtitle(subtitle string) {
	c_subtitle := C.CString(subtitle)
	defer C.free(unsafe.Pointer(c_subtitle))

	C.gtk_header_bar_set_subtitle((*C.GtkHeaderBar)(recv.native), c_subtitle)

	return
}

// SetTitle is a wrapper around the C function gtk_header_bar_set_title.
func (recv *HeaderBar) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_header_bar_set_title((*C.GtkHeaderBar)(recv.native), c_title)

	return
}

// GetBaseScale is a wrapper around the C function gtk_icon_info_get_base_scale.
func (recv *IconInfo) GetBaseScale() int32 {
	retC := C.gtk_icon_info_get_base_scale((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// LoadSurface is a wrapper around the C function gtk_icon_info_load_surface.
func (recv *IconInfo) LoadSurface(forWindow *gdk.Window) (*cairo.Surface, error) {
	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_surface((*C.GtkIconInfo)(recv.native), c_for_window, &cThrowableError)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Blacklisted : gtk_icon_theme_choose_icon_for_scale

// LoadIconForScale is a wrapper around the C function gtk_icon_theme_load_icon_for_scale.
func (recv *IconTheme) LoadIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags, &cThrowableError)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadSurface is a wrapper around the C function gtk_icon_theme_load_surface.
func (recv *IconTheme) LoadSurface(iconName string, size int32, scale int32, forWindow *gdk.Window, flags IconLookupFlags) (*cairo.Surface, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_surface((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_for_window, c_flags, &cThrowableError)
	var retGo (*cairo.Surface)
	if retC == nil {
		retGo = nil
	} else {
		retGo = cairo.SurfaceNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LookupByGiconForScale is a wrapper around the C function gtk_icon_theme_lookup_by_gicon_for_scale.
func (recv *IconTheme) LookupByGiconForScale(icon *gio.Icon, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_by_gicon_for_scale((*C.GtkIconTheme)(recv.native), c_icon, c_size, c_scale, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// LookupIconForScale is a wrapper around the C function gtk_icon_theme_lookup_icon_for_scale.
func (recv *IconTheme) LookupIconForScale(iconName string, size int32, scale int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_scale := (C.gint)(scale)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon_for_scale((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_scale, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ImageNewFromSurface is a wrapper around the C function gtk_image_new_from_surface.
func ImageNewFromSurface(surface *cairo.Surface) *Image {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	retC := C.gtk_image_new_from_surface(c_surface)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFromSurface is a wrapper around the C function gtk_image_set_from_surface.
func (recv *Image) SetFromSurface(surface *cairo.Surface) {
	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	C.gtk_image_set_from_surface((*C.GtkImage)(recv.native), c_surface)

	return
}

// GetShowCloseButton is a wrapper around the C function gtk_info_bar_get_show_close_button.
func (recv *InfoBar) GetShowCloseButton() bool {
	retC := C.gtk_info_bar_get_show_close_button((*C.GtkInfoBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowCloseButton is a wrapper around the C function gtk_info_bar_set_show_close_button.
func (recv *InfoBar) SetShowCloseButton(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_info_bar_set_show_close_button((*C.GtkInfoBar)(recv.native), c_setting)

	return
}

// GetLines is a wrapper around the C function gtk_label_get_lines.
func (recv *Label) GetLines() int32 {
	retC := C.gtk_label_get_lines((*C.GtkLabel)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetLines is a wrapper around the C function gtk_label_set_lines.
func (recv *Label) SetLines(lines int32) {
	c_lines := (C.gint)(lines)

	C.gtk_label_set_lines((*C.GtkLabel)(recv.native), c_lines)

	return
}

type signalListBoxRowActivatedDetail struct {
	callback  ListBoxSignalRowActivatedCallback
	handlerID C.gulong
}

var signalListBoxRowActivatedId int
var signalListBoxRowActivatedMap = make(map[int]signalListBoxRowActivatedDetail)
var signalListBoxRowActivatedLock sync.RWMutex

// ListBoxSignalRowActivatedCallback is a callback function for a 'row-activated' signal emitted from a ListBox.
type ListBoxSignalRowActivatedCallback func(row *ListBoxRow)

/*
ConnectRowActivated connects the callback to the 'row-activated' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowActivated to remove it.
*/
func (recv *ListBox) ConnectRowActivated(callback ListBoxSignalRowActivatedCallback) int {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	signalListBoxRowActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_activated(instance, C.gpointer(uintptr(signalListBoxRowActivatedId)))

	detail := signalListBoxRowActivatedDetail{callback, handlerID}
	signalListBoxRowActivatedMap[signalListBoxRowActivatedId] = detail

	return signalListBoxRowActivatedId
}

/*
DisconnectRowActivated disconnects a callback from the 'row-activated' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowActivated.
*/
func (recv *ListBox) DisconnectRowActivated(connectionID int) {
	signalListBoxRowActivatedLock.Lock()
	defer signalListBoxRowActivatedLock.Unlock()

	detail, exists := signalListBoxRowActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivatedMap, connectionID)
}

//export listbox_rowActivatedHandler
func listbox_rowActivatedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	signalListBoxRowActivatedLock.RLock()
	defer signalListBoxRowActivatedLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowActivatedMap[index].callback
	callback(row)
}

type signalListBoxRowSelectedDetail struct {
	callback  ListBoxSignalRowSelectedCallback
	handlerID C.gulong
}

var signalListBoxRowSelectedId int
var signalListBoxRowSelectedMap = make(map[int]signalListBoxRowSelectedDetail)
var signalListBoxRowSelectedLock sync.RWMutex

// ListBoxSignalRowSelectedCallback is a callback function for a 'row-selected' signal emitted from a ListBox.
type ListBoxSignalRowSelectedCallback func(row *ListBoxRow)

/*
ConnectRowSelected connects the callback to the 'row-selected' signal for the ListBox.

The returned value represents the connection, and may be passed to DisconnectRowSelected to remove it.
*/
func (recv *ListBox) ConnectRowSelected(callback ListBoxSignalRowSelectedCallback) int {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	signalListBoxRowSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBox_signal_connect_row_selected(instance, C.gpointer(uintptr(signalListBoxRowSelectedId)))

	detail := signalListBoxRowSelectedDetail{callback, handlerID}
	signalListBoxRowSelectedMap[signalListBoxRowSelectedId] = detail

	return signalListBoxRowSelectedId
}

/*
DisconnectRowSelected disconnects a callback from the 'row-selected' signal for the ListBox.

The connectionID should be a value returned from a call to ConnectRowSelected.
*/
func (recv *ListBox) DisconnectRowSelected(connectionID int) {
	signalListBoxRowSelectedLock.Lock()
	defer signalListBoxRowSelectedLock.Unlock()

	detail, exists := signalListBoxRowSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowSelectedMap, connectionID)
}

//export listbox_rowSelectedHandler
func listbox_rowSelectedHandler(_ *C.GObject, c_row *C.GtkListBoxRow, data C.gpointer) {
	signalListBoxRowSelectedLock.RLock()
	defer signalListBoxRowSelectedLock.RUnlock()

	row := ListBoxRowNewFromC(unsafe.Pointer(c_row))

	index := int(uintptr(data))
	callback := signalListBoxRowSelectedMap[index].callback
	callback(row)
}

// ListBoxNew is a wrapper around the C function gtk_list_box_new.
func ListBoxNew() *ListBox {
	retC := C.gtk_list_box_new()
	retGo := ListBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DragHighlightRow is a wrapper around the C function gtk_list_box_drag_highlight_row.
func (recv *ListBox) DragHighlightRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(C.NULL)
	if row != nil {
		c_row = (*C.GtkListBoxRow)(row.ToC())
	}

	C.gtk_list_box_drag_highlight_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// DragUnhighlightRow is a wrapper around the C function gtk_list_box_drag_unhighlight_row.
func (recv *ListBox) DragUnhighlightRow() {
	C.gtk_list_box_drag_unhighlight_row((*C.GtkListBox)(recv.native))

	return
}

// GetActivateOnSingleClick is a wrapper around the C function gtk_list_box_get_activate_on_single_click.
func (recv *ListBox) GetActivateOnSingleClick() bool {
	retC := C.gtk_list_box_get_activate_on_single_click((*C.GtkListBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetAdjustment is a wrapper around the C function gtk_list_box_get_adjustment.
func (recv *ListBox) GetAdjustment() *Adjustment {
	retC := C.gtk_list_box_get_adjustment((*C.GtkListBox)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetRowAtIndex is a wrapper around the C function gtk_list_box_get_row_at_index.
func (recv *ListBox) GetRowAtIndex(index int32) *ListBoxRow {
	c_index_ := (C.gint)(index)

	retC := C.gtk_list_box_get_row_at_index((*C.GtkListBox)(recv.native), c_index_)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetRowAtY is a wrapper around the C function gtk_list_box_get_row_at_y.
func (recv *ListBox) GetRowAtY(y int32) *ListBoxRow {
	c_y := (C.gint)(y)

	retC := C.gtk_list_box_get_row_at_y((*C.GtkListBox)(recv.native), c_y)
	var retGo (*ListBoxRow)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ListBoxRowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSelectedRow is a wrapper around the C function gtk_list_box_get_selected_row.
func (recv *ListBox) GetSelectedRow() *ListBoxRow {
	retC := C.gtk_list_box_get_selected_row((*C.GtkListBox)(recv.native))
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionMode is a wrapper around the C function gtk_list_box_get_selection_mode.
func (recv *ListBox) GetSelectionMode() SelectionMode {
	retC := C.gtk_list_box_get_selection_mode((*C.GtkListBox)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Insert is a wrapper around the C function gtk_list_box_insert.
func (recv *ListBox) Insert(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_list_box_insert((*C.GtkListBox)(recv.native), c_child, c_position)

	return
}

// InvalidateFilter is a wrapper around the C function gtk_list_box_invalidate_filter.
func (recv *ListBox) InvalidateFilter() {
	C.gtk_list_box_invalidate_filter((*C.GtkListBox)(recv.native))

	return
}

// InvalidateHeaders is a wrapper around the C function gtk_list_box_invalidate_headers.
func (recv *ListBox) InvalidateHeaders() {
	C.gtk_list_box_invalidate_headers((*C.GtkListBox)(recv.native))

	return
}

// InvalidateSort is a wrapper around the C function gtk_list_box_invalidate_sort.
func (recv *ListBox) InvalidateSort() {
	C.gtk_list_box_invalidate_sort((*C.GtkListBox)(recv.native))

	return
}

// Prepend is a wrapper around the C function gtk_list_box_prepend.
func (recv *ListBox) Prepend(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_list_box_prepend((*C.GtkListBox)(recv.native), c_child)

	return
}

// SelectRow is a wrapper around the C function gtk_list_box_select_row.
func (recv *ListBox) SelectRow(row *ListBoxRow) {
	c_row := (*C.GtkListBoxRow)(C.NULL)
	if row != nil {
		c_row = (*C.GtkListBoxRow)(row.ToC())
	}

	C.gtk_list_box_select_row((*C.GtkListBox)(recv.native), c_row)

	return
}

// SetActivateOnSingleClick is a wrapper around the C function gtk_list_box_set_activate_on_single_click.
func (recv *ListBox) SetActivateOnSingleClick(single bool) {
	c_single :=
		boolToGboolean(single)

	C.gtk_list_box_set_activate_on_single_click((*C.GtkListBox)(recv.native), c_single)

	return
}

// SetAdjustment is a wrapper around the C function gtk_list_box_set_adjustment.
func (recv *ListBox) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(C.NULL)
	if adjustment != nil {
		c_adjustment = (*C.GtkAdjustment)(adjustment.ToC())
	}

	C.gtk_list_box_set_adjustment((*C.GtkListBox)(recv.native), c_adjustment)

	return
}

// Unsupported : gtk_list_box_set_filter_func : unsupported parameter filter_func : no type generator for ListBoxFilterFunc (GtkListBoxFilterFunc) for param filter_func

// Unsupported : gtk_list_box_set_header_func : unsupported parameter update_header : no type generator for ListBoxUpdateHeaderFunc (GtkListBoxUpdateHeaderFunc) for param update_header

// SetPlaceholder is a wrapper around the C function gtk_list_box_set_placeholder.
func (recv *ListBox) SetPlaceholder(placeholder *Widget) {
	c_placeholder := (*C.GtkWidget)(C.NULL)
	if placeholder != nil {
		c_placeholder = (*C.GtkWidget)(placeholder.ToC())
	}

	C.gtk_list_box_set_placeholder((*C.GtkListBox)(recv.native), c_placeholder)

	return
}

// SetSelectionMode is a wrapper around the C function gtk_list_box_set_selection_mode.
func (recv *ListBox) SetSelectionMode(mode SelectionMode) {
	c_mode := (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode((*C.GtkListBox)(recv.native), c_mode)

	return
}

// Unsupported : gtk_list_box_set_sort_func : unsupported parameter sort_func : no type generator for ListBoxSortFunc (GtkListBoxSortFunc) for param sort_func

type signalListBoxRowActivateDetail struct {
	callback  ListBoxRowSignalActivateCallback
	handlerID C.gulong
}

var signalListBoxRowActivateId int
var signalListBoxRowActivateMap = make(map[int]signalListBoxRowActivateDetail)
var signalListBoxRowActivateLock sync.RWMutex

// ListBoxRowSignalActivateCallback is a callback function for a 'activate' signal emitted from a ListBoxRow.
type ListBoxRowSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the ListBoxRow.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *ListBoxRow) ConnectActivate(callback ListBoxRowSignalActivateCallback) int {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	signalListBoxRowActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ListBoxRow_signal_connect_activate(instance, C.gpointer(uintptr(signalListBoxRowActivateId)))

	detail := signalListBoxRowActivateDetail{callback, handlerID}
	signalListBoxRowActivateMap[signalListBoxRowActivateId] = detail

	return signalListBoxRowActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the ListBoxRow.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *ListBoxRow) DisconnectActivate(connectionID int) {
	signalListBoxRowActivateLock.Lock()
	defer signalListBoxRowActivateLock.Unlock()

	detail, exists := signalListBoxRowActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalListBoxRowActivateMap, connectionID)
}

//export listboxrow_activateHandler
func listboxrow_activateHandler(_ *C.GObject, data C.gpointer) {
	signalListBoxRowActivateLock.RLock()
	defer signalListBoxRowActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalListBoxRowActivateMap[index].callback
	callback()
}

// ListBoxRowNew is a wrapper around the C function gtk_list_box_row_new.
func ListBoxRowNew() *ListBoxRow {
	retC := C.gtk_list_box_row_new()
	retGo := ListBoxRowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_list_box_row_changed.
func (recv *ListBoxRow) Changed() {
	C.gtk_list_box_row_changed((*C.GtkListBoxRow)(recv.native))

	return
}

// GetHeader is a wrapper around the C function gtk_list_box_row_get_header.
func (recv *ListBoxRow) GetHeader() *Widget {
	retC := C.gtk_list_box_row_get_header((*C.GtkListBoxRow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetIndex is a wrapper around the C function gtk_list_box_row_get_index.
func (recv *ListBoxRow) GetIndex() int32 {
	retC := C.gtk_list_box_row_get_index((*C.GtkListBoxRow)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetHeader is a wrapper around the C function gtk_list_box_row_set_header.
func (recv *ListBoxRow) SetHeader(header *Widget) {
	c_header := (*C.GtkWidget)(C.NULL)
	if header != nil {
		c_header = (*C.GtkWidget)(header.ToC())
	}

	C.gtk_list_box_row_set_header((*C.GtkListBoxRow)(recv.native), c_header)

	return
}

type signalPlacesSidebarDragActionAskDetail struct {
	callback  PlacesSidebarSignalDragActionAskCallback
	handlerID C.gulong
}

var signalPlacesSidebarDragActionAskId int
var signalPlacesSidebarDragActionAskMap = make(map[int]signalPlacesSidebarDragActionAskDetail)
var signalPlacesSidebarDragActionAskLock sync.RWMutex

// PlacesSidebarSignalDragActionAskCallback is a callback function for a 'drag-action-ask' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalDragActionAskCallback func(actions int32) int32

/*
ConnectDragActionAsk connects the callback to the 'drag-action-ask' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectDragActionAsk to remove it.
*/
func (recv *PlacesSidebar) ConnectDragActionAsk(callback PlacesSidebarSignalDragActionAskCallback) int {
	signalPlacesSidebarDragActionAskLock.Lock()
	defer signalPlacesSidebarDragActionAskLock.Unlock()

	signalPlacesSidebarDragActionAskId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_drag_action_ask(instance, C.gpointer(uintptr(signalPlacesSidebarDragActionAskId)))

	detail := signalPlacesSidebarDragActionAskDetail{callback, handlerID}
	signalPlacesSidebarDragActionAskMap[signalPlacesSidebarDragActionAskId] = detail

	return signalPlacesSidebarDragActionAskId
}

/*
DisconnectDragActionAsk disconnects a callback from the 'drag-action-ask' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectDragActionAsk.
*/
func (recv *PlacesSidebar) DisconnectDragActionAsk(connectionID int) {
	signalPlacesSidebarDragActionAskLock.Lock()
	defer signalPlacesSidebarDragActionAskLock.Unlock()

	detail, exists := signalPlacesSidebarDragActionAskMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarDragActionAskMap, connectionID)
}

//export placessidebar_dragActionAskHandler
func placessidebar_dragActionAskHandler(_ *C.GObject, c_actions C.gint, data C.gpointer) C.gint {
	signalPlacesSidebarDragActionAskLock.RLock()
	defer signalPlacesSidebarDragActionAskLock.RUnlock()

	actions := int32(c_actions)

	index := int(uintptr(data))
	callback := signalPlacesSidebarDragActionAskMap[index].callback
	retGo := callback(actions)
	retC :=
		(C.gint)(retGo)
	return retC
}

// Unsupported signal 'drag-action-requested' for PlacesSidebar : param source_file_list : gpointer

// Unsupported signal 'drag-perform-drop' for PlacesSidebar : param source_file_list : gpointer

type signalPlacesSidebarOpenLocationDetail struct {
	callback  PlacesSidebarSignalOpenLocationCallback
	handlerID C.gulong
}

var signalPlacesSidebarOpenLocationId int
var signalPlacesSidebarOpenLocationMap = make(map[int]signalPlacesSidebarOpenLocationDetail)
var signalPlacesSidebarOpenLocationLock sync.RWMutex

// PlacesSidebarSignalOpenLocationCallback is a callback function for a 'open-location' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalOpenLocationCallback func(location *gio.File, openFlags PlacesOpenFlags)

/*
ConnectOpenLocation connects the callback to the 'open-location' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectOpenLocation to remove it.
*/
func (recv *PlacesSidebar) ConnectOpenLocation(callback PlacesSidebarSignalOpenLocationCallback) int {
	signalPlacesSidebarOpenLocationLock.Lock()
	defer signalPlacesSidebarOpenLocationLock.Unlock()

	signalPlacesSidebarOpenLocationId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_open_location(instance, C.gpointer(uintptr(signalPlacesSidebarOpenLocationId)))

	detail := signalPlacesSidebarOpenLocationDetail{callback, handlerID}
	signalPlacesSidebarOpenLocationMap[signalPlacesSidebarOpenLocationId] = detail

	return signalPlacesSidebarOpenLocationId
}

/*
DisconnectOpenLocation disconnects a callback from the 'open-location' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectOpenLocation.
*/
func (recv *PlacesSidebar) DisconnectOpenLocation(connectionID int) {
	signalPlacesSidebarOpenLocationLock.Lock()
	defer signalPlacesSidebarOpenLocationLock.Unlock()

	detail, exists := signalPlacesSidebarOpenLocationMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarOpenLocationMap, connectionID)
}

//export placessidebar_openLocationHandler
func placessidebar_openLocationHandler(_ *C.GObject, c_location *C.GFile, c_open_flags C.GtkPlacesOpenFlags, data C.gpointer) {
	signalPlacesSidebarOpenLocationLock.RLock()
	defer signalPlacesSidebarOpenLocationLock.RUnlock()

	location := gio.FileNewFromC(unsafe.Pointer(c_location))

	openFlags := PlacesOpenFlags(c_open_flags)

	index := int(uintptr(data))
	callback := signalPlacesSidebarOpenLocationMap[index].callback
	callback(location, openFlags)
}

type signalPlacesSidebarPopulatePopupDetail struct {
	callback  PlacesSidebarSignalPopulatePopupCallback
	handlerID C.gulong
}

var signalPlacesSidebarPopulatePopupId int
var signalPlacesSidebarPopulatePopupMap = make(map[int]signalPlacesSidebarPopulatePopupDetail)
var signalPlacesSidebarPopulatePopupLock sync.RWMutex

// PlacesSidebarSignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalPopulatePopupCallback func(container *Widget, selectedItem *gio.File, selectedVolume *gio.Volume)

/*
ConnectPopulatePopup connects the callback to the 'populate-popup' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectPopulatePopup to remove it.
*/
func (recv *PlacesSidebar) ConnectPopulatePopup(callback PlacesSidebarSignalPopulatePopupCallback) int {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	signalPlacesSidebarPopulatePopupId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_populate_popup(instance, C.gpointer(uintptr(signalPlacesSidebarPopulatePopupId)))

	detail := signalPlacesSidebarPopulatePopupDetail{callback, handlerID}
	signalPlacesSidebarPopulatePopupMap[signalPlacesSidebarPopulatePopupId] = detail

	return signalPlacesSidebarPopulatePopupId
}

/*
DisconnectPopulatePopup disconnects a callback from the 'populate-popup' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectPopulatePopup.
*/
func (recv *PlacesSidebar) DisconnectPopulatePopup(connectionID int) {
	signalPlacesSidebarPopulatePopupLock.Lock()
	defer signalPlacesSidebarPopulatePopupLock.Unlock()

	detail, exists := signalPlacesSidebarPopulatePopupMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarPopulatePopupMap, connectionID)
}

//export placessidebar_populatePopupHandler
func placessidebar_populatePopupHandler(_ *C.GObject, c_container *C.GtkWidget, c_selected_item *C.GFile, c_selected_volume *C.GVolume, data C.gpointer) {
	signalPlacesSidebarPopulatePopupLock.RLock()
	defer signalPlacesSidebarPopulatePopupLock.RUnlock()

	container := WidgetNewFromC(unsafe.Pointer(c_container))

	selectedItem := gio.FileNewFromC(unsafe.Pointer(c_selected_item))

	selectedVolume := gio.VolumeNewFromC(unsafe.Pointer(c_selected_volume))

	index := int(uintptr(data))
	callback := signalPlacesSidebarPopulatePopupMap[index].callback
	callback(container, selectedItem, selectedVolume)
}

type signalPlacesSidebarShowErrorMessageDetail struct {
	callback  PlacesSidebarSignalShowErrorMessageCallback
	handlerID C.gulong
}

var signalPlacesSidebarShowErrorMessageId int
var signalPlacesSidebarShowErrorMessageMap = make(map[int]signalPlacesSidebarShowErrorMessageDetail)
var signalPlacesSidebarShowErrorMessageLock sync.RWMutex

// PlacesSidebarSignalShowErrorMessageCallback is a callback function for a 'show-error-message' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowErrorMessageCallback func(primary string, secondary string)

/*
ConnectShowErrorMessage connects the callback to the 'show-error-message' signal for the PlacesSidebar.

The returned value represents the connection, and may be passed to DisconnectShowErrorMessage to remove it.
*/
func (recv *PlacesSidebar) ConnectShowErrorMessage(callback PlacesSidebarSignalShowErrorMessageCallback) int {
	signalPlacesSidebarShowErrorMessageLock.Lock()
	defer signalPlacesSidebarShowErrorMessageLock.Unlock()

	signalPlacesSidebarShowErrorMessageId++
	instance := C.gpointer(recv.native)
	handlerID := C.PlacesSidebar_signal_connect_show_error_message(instance, C.gpointer(uintptr(signalPlacesSidebarShowErrorMessageId)))

	detail := signalPlacesSidebarShowErrorMessageDetail{callback, handlerID}
	signalPlacesSidebarShowErrorMessageMap[signalPlacesSidebarShowErrorMessageId] = detail

	return signalPlacesSidebarShowErrorMessageId
}

/*
DisconnectShowErrorMessage disconnects a callback from the 'show-error-message' signal for the PlacesSidebar.

The connectionID should be a value returned from a call to ConnectShowErrorMessage.
*/
func (recv *PlacesSidebar) DisconnectShowErrorMessage(connectionID int) {
	signalPlacesSidebarShowErrorMessageLock.Lock()
	defer signalPlacesSidebarShowErrorMessageLock.Unlock()

	detail, exists := signalPlacesSidebarShowErrorMessageMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalPlacesSidebarShowErrorMessageMap, connectionID)
}

//export placessidebar_showErrorMessageHandler
func placessidebar_showErrorMessageHandler(_ *C.GObject, c_primary *C.gchar, c_secondary *C.gchar, data C.gpointer) {
	signalPlacesSidebarShowErrorMessageLock.RLock()
	defer signalPlacesSidebarShowErrorMessageLock.RUnlock()

	primary := C.GoString(c_primary)

	secondary := C.GoString(c_secondary)

	index := int(uintptr(data))
	callback := signalPlacesSidebarShowErrorMessageMap[index].callback
	callback(primary, secondary)
}

// PlacesSidebarNew is a wrapper around the C function gtk_places_sidebar_new.
func PlacesSidebarNew() *PlacesSidebar {
	retC := C.gtk_places_sidebar_new()
	retGo := PlacesSidebarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddShortcut is a wrapper around the C function gtk_places_sidebar_add_shortcut.
func (recv *PlacesSidebar) AddShortcut(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_add_shortcut((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// GetLocation is a wrapper around the C function gtk_places_sidebar_get_location.
func (recv *PlacesSidebar) GetLocation() *gio.File {
	retC := C.gtk_places_sidebar_get_location((*C.GtkPlacesSidebar)(recv.native))
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetNthBookmark is a wrapper around the C function gtk_places_sidebar_get_nth_bookmark.
func (recv *PlacesSidebar) GetNthBookmark(n int32) *gio.File {
	c_n := (C.gint)(n)

	retC := C.gtk_places_sidebar_get_nth_bookmark((*C.GtkPlacesSidebar)(recv.native), c_n)
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetOpenFlags is a wrapper around the C function gtk_places_sidebar_get_open_flags.
func (recv *PlacesSidebar) GetOpenFlags() PlacesOpenFlags {
	retC := C.gtk_places_sidebar_get_open_flags((*C.GtkPlacesSidebar)(recv.native))
	retGo := (PlacesOpenFlags)(retC)

	return retGo
}

// GetShowDesktop is a wrapper around the C function gtk_places_sidebar_get_show_desktop.
func (recv *PlacesSidebar) GetShowDesktop() bool {
	retC := C.gtk_places_sidebar_get_show_desktop((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListShortcuts is a wrapper around the C function gtk_places_sidebar_list_shortcuts.
func (recv *PlacesSidebar) ListShortcuts() *glib.SList {
	retC := C.gtk_places_sidebar_list_shortcuts((*C.GtkPlacesSidebar)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveShortcut is a wrapper around the C function gtk_places_sidebar_remove_shortcut.
func (recv *PlacesSidebar) RemoveShortcut(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_remove_shortcut((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// SetLocation is a wrapper around the C function gtk_places_sidebar_set_location.
func (recv *PlacesSidebar) SetLocation(location *gio.File) {
	c_location := (*C.GFile)(location.ToC())

	C.gtk_places_sidebar_set_location((*C.GtkPlacesSidebar)(recv.native), c_location)

	return
}

// SetOpenFlags is a wrapper around the C function gtk_places_sidebar_set_open_flags.
func (recv *PlacesSidebar) SetOpenFlags(flags PlacesOpenFlags) {
	c_flags := (C.GtkPlacesOpenFlags)(flags)

	C.gtk_places_sidebar_set_open_flags((*C.GtkPlacesSidebar)(recv.native), c_flags)

	return
}

// SetShowConnectToServer is a wrapper around the C function gtk_places_sidebar_set_show_connect_to_server.
func (recv *PlacesSidebar) SetShowConnectToServer(showConnectToServer bool) {
	c_show_connect_to_server :=
		boolToGboolean(showConnectToServer)

	C.gtk_places_sidebar_set_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native), c_show_connect_to_server)

	return
}

// SetShowDesktop is a wrapper around the C function gtk_places_sidebar_set_show_desktop.
func (recv *PlacesSidebar) SetShowDesktop(showDesktop bool) {
	c_show_desktop :=
		boolToGboolean(showDesktop)

	C.gtk_places_sidebar_set_show_desktop((*C.GtkPlacesSidebar)(recv.native), c_show_desktop)

	return
}

// RevealerNew is a wrapper around the C function gtk_revealer_new.
func RevealerNew() *Revealer {
	retC := C.gtk_revealer_new()
	retGo := RevealerNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildRevealed is a wrapper around the C function gtk_revealer_get_child_revealed.
func (recv *Revealer) GetChildRevealed() bool {
	retC := C.gtk_revealer_get_child_revealed((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRevealChild is a wrapper around the C function gtk_revealer_get_reveal_child.
func (recv *Revealer) GetRevealChild() bool {
	retC := C.gtk_revealer_get_reveal_child((*C.GtkRevealer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_revealer_get_transition_duration.
func (recv *Revealer) GetTransitionDuration() uint32 {
	retC := C.gtk_revealer_get_transition_duration((*C.GtkRevealer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_revealer_get_transition_type.
func (recv *Revealer) GetTransitionType() RevealerTransitionType {
	retC := C.gtk_revealer_get_transition_type((*C.GtkRevealer)(recv.native))
	retGo := (RevealerTransitionType)(retC)

	return retGo
}

// SetRevealChild is a wrapper around the C function gtk_revealer_set_reveal_child.
func (recv *Revealer) SetRevealChild(revealChild bool) {
	c_reveal_child :=
		boolToGboolean(revealChild)

	C.gtk_revealer_set_reveal_child((*C.GtkRevealer)(recv.native), c_reveal_child)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_revealer_set_transition_duration.
func (recv *Revealer) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_revealer_set_transition_duration((*C.GtkRevealer)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_revealer_set_transition_type.
func (recv *Revealer) SetTransitionType(transition RevealerTransitionType) {
	c_transition := (C.GtkRevealerTransitionType)(transition)

	C.gtk_revealer_set_transition_type((*C.GtkRevealer)(recv.native), c_transition)

	return
}

// SearchBarNew is a wrapper around the C function gtk_search_bar_new.
func SearchBarNew() *SearchBar {
	retC := C.gtk_search_bar_new()
	retGo := SearchBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConnectEntry is a wrapper around the C function gtk_search_bar_connect_entry.
func (recv *SearchBar) ConnectEntry(entry *Entry) {
	c_entry := (*C.GtkEntry)(C.NULL)
	if entry != nil {
		c_entry = (*C.GtkEntry)(entry.ToC())
	}

	C.gtk_search_bar_connect_entry((*C.GtkSearchBar)(recv.native), c_entry)

	return
}

// GetSearchMode is a wrapper around the C function gtk_search_bar_get_search_mode.
func (recv *SearchBar) GetSearchMode() bool {
	retC := C.gtk_search_bar_get_search_mode((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowCloseButton is a wrapper around the C function gtk_search_bar_get_show_close_button.
func (recv *SearchBar) GetShowCloseButton() bool {
	retC := C.gtk_search_bar_get_show_close_button((*C.GtkSearchBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_search_bar_handle_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// SetSearchMode is a wrapper around the C function gtk_search_bar_set_search_mode.
func (recv *SearchBar) SetSearchMode(searchMode bool) {
	c_search_mode :=
		boolToGboolean(searchMode)

	C.gtk_search_bar_set_search_mode((*C.GtkSearchBar)(recv.native), c_search_mode)

	return
}

// SetShowCloseButton is a wrapper around the C function gtk_search_bar_set_show_close_button.
func (recv *SearchBar) SetShowCloseButton(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_search_bar_set_show_close_button((*C.GtkSearchBar)(recv.native), c_visible)

	return
}

type signalSearchEntrySearchChangedDetail struct {
	callback  SearchEntrySignalSearchChangedCallback
	handlerID C.gulong
}

var signalSearchEntrySearchChangedId int
var signalSearchEntrySearchChangedMap = make(map[int]signalSearchEntrySearchChangedDetail)
var signalSearchEntrySearchChangedLock sync.RWMutex

// SearchEntrySignalSearchChangedCallback is a callback function for a 'search-changed' signal emitted from a SearchEntry.
type SearchEntrySignalSearchChangedCallback func()

/*
ConnectSearchChanged connects the callback to the 'search-changed' signal for the SearchEntry.

The returned value represents the connection, and may be passed to DisconnectSearchChanged to remove it.
*/
func (recv *SearchEntry) ConnectSearchChanged(callback SearchEntrySignalSearchChangedCallback) int {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	signalSearchEntrySearchChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.SearchEntry_signal_connect_search_changed(instance, C.gpointer(uintptr(signalSearchEntrySearchChangedId)))

	detail := signalSearchEntrySearchChangedDetail{callback, handlerID}
	signalSearchEntrySearchChangedMap[signalSearchEntrySearchChangedId] = detail

	return signalSearchEntrySearchChangedId
}

/*
DisconnectSearchChanged disconnects a callback from the 'search-changed' signal for the SearchEntry.

The connectionID should be a value returned from a call to ConnectSearchChanged.
*/
func (recv *SearchEntry) DisconnectSearchChanged(connectionID int) {
	signalSearchEntrySearchChangedLock.Lock()
	defer signalSearchEntrySearchChangedLock.Unlock()

	detail, exists := signalSearchEntrySearchChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSearchEntrySearchChangedMap, connectionID)
}

//export searchentry_searchChangedHandler
func searchentry_searchChangedHandler(_ *C.GObject, data C.gpointer) {
	signalSearchEntrySearchChangedLock.RLock()
	defer signalSearchEntrySearchChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalSearchEntrySearchChangedMap[index].callback
	callback()
}

// StackNew is a wrapper around the C function gtk_stack_new.
func StackNew() *Stack {
	retC := C.gtk_stack_new()
	retGo := StackNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddNamed is a wrapper around the C function gtk_stack_add_named.
func (recv *Stack) AddNamed(child *Widget, name string) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_add_named((*C.GtkStack)(recv.native), c_child, c_name)

	return
}

// AddTitled is a wrapper around the C function gtk_stack_add_titled.
func (recv *Stack) AddTitled(child *Widget, name string, title string) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_stack_add_titled((*C.GtkStack)(recv.native), c_child, c_name, c_title)

	return
}

// GetHomogeneous is a wrapper around the C function gtk_stack_get_homogeneous.
func (recv *Stack) GetHomogeneous() bool {
	retC := C.gtk_stack_get_homogeneous((*C.GtkStack)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTransitionDuration is a wrapper around the C function gtk_stack_get_transition_duration.
func (recv *Stack) GetTransitionDuration() uint32 {
	retC := C.gtk_stack_get_transition_duration((*C.GtkStack)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetTransitionType is a wrapper around the C function gtk_stack_get_transition_type.
func (recv *Stack) GetTransitionType() StackTransitionType {
	retC := C.gtk_stack_get_transition_type((*C.GtkStack)(recv.native))
	retGo := (StackTransitionType)(retC)

	return retGo
}

// GetVisibleChild is a wrapper around the C function gtk_stack_get_visible_child.
func (recv *Stack) GetVisibleChild() *Widget {
	retC := C.gtk_stack_get_visible_child((*C.GtkStack)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetVisibleChildName is a wrapper around the C function gtk_stack_get_visible_child_name.
func (recv *Stack) GetVisibleChildName() string {
	retC := C.gtk_stack_get_visible_child_name((*C.GtkStack)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetHomogeneous is a wrapper around the C function gtk_stack_set_homogeneous.
func (recv *Stack) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_stack_set_homogeneous((*C.GtkStack)(recv.native), c_homogeneous)

	return
}

// SetTransitionDuration is a wrapper around the C function gtk_stack_set_transition_duration.
func (recv *Stack) SetTransitionDuration(duration uint32) {
	c_duration := (C.guint)(duration)

	C.gtk_stack_set_transition_duration((*C.GtkStack)(recv.native), c_duration)

	return
}

// SetTransitionType is a wrapper around the C function gtk_stack_set_transition_type.
func (recv *Stack) SetTransitionType(transition StackTransitionType) {
	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_transition_type((*C.GtkStack)(recv.native), c_transition)

	return
}

// SetVisibleChild is a wrapper around the C function gtk_stack_set_visible_child.
func (recv *Stack) SetVisibleChild(child *Widget) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	C.gtk_stack_set_visible_child((*C.GtkStack)(recv.native), c_child)

	return
}

// SetVisibleChildFull is a wrapper around the C function gtk_stack_set_visible_child_full.
func (recv *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_transition := (C.GtkStackTransitionType)(transition)

	C.gtk_stack_set_visible_child_full((*C.GtkStack)(recv.native), c_name, c_transition)

	return
}

// SetVisibleChildName is a wrapper around the C function gtk_stack_set_visible_child_name.
func (recv *Stack) SetVisibleChildName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_stack_set_visible_child_name((*C.GtkStack)(recv.native), c_name)

	return
}

// StackSwitcherNew is a wrapper around the C function gtk_stack_switcher_new.
func StackSwitcherNew() *StackSwitcher {
	retC := C.gtk_stack_switcher_new()
	retGo := StackSwitcherNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStack is a wrapper around the C function gtk_stack_switcher_get_stack.
func (recv *StackSwitcher) GetStack() *Stack {
	retC := C.gtk_stack_switcher_get_stack((*C.GtkStackSwitcher)(recv.native))
	var retGo (*Stack)
	if retC == nil {
		retGo = nil
	} else {
		retGo = StackNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SetStack is a wrapper around the C function gtk_stack_switcher_set_stack.
func (recv *StackSwitcher) SetStack(stack *Stack) {
	c_stack := (*C.GtkStack)(C.NULL)
	if stack != nil {
		c_stack = (*C.GtkStack)(stack.ToC())
	}

	C.gtk_stack_switcher_set_stack((*C.GtkStackSwitcher)(recv.native), c_stack)

	return
}

// GetScale is a wrapper around the C function gtk_style_context_get_scale.
func (recv *StyleContext) GetScale() int32 {
	retC := C.gtk_style_context_get_scale((*C.GtkStyleContext)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetScale is a wrapper around the C function gtk_style_context_set_scale.
func (recv *StyleContext) SetScale(scale int32) {
	c_scale := (C.gint)(scale)

	C.gtk_style_context_set_scale((*C.GtkStyleContext)(recv.native), c_scale)

	return
}

// Unsupported : gtk_drag_begin_with_coordinates : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// GetAllocatedBaseline is a wrapper around the C function gtk_widget_get_allocated_baseline.
func (recv *Widget) GetAllocatedBaseline() int32 {
	retC := C.gtk_widget_get_allocated_baseline((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPreferredHeightAndBaselineForWidth is a wrapper around the C function gtk_widget_get_preferred_height_and_baseline_for_width.
func (recv *Widget) GetPreferredHeightAndBaselineForWidth(width int32) (int32, int32, int32, int32) {
	c_width := (C.gint)(width)

	var c_minimum_height C.gint

	var c_natural_height C.gint

	var c_minimum_baseline C.gint

	var c_natural_baseline C.gint

	C.gtk_widget_get_preferred_height_and_baseline_for_width((*C.GtkWidget)(recv.native), c_width, &c_minimum_height, &c_natural_height, &c_minimum_baseline, &c_natural_baseline)

	minimumHeight := (int32)(c_minimum_height)

	naturalHeight := (int32)(c_natural_height)

	minimumBaseline := (int32)(c_minimum_baseline)

	naturalBaseline := (int32)(c_natural_baseline)

	return minimumHeight, naturalHeight, minimumBaseline, naturalBaseline
}

// GetScaleFactor is a wrapper around the C function gtk_widget_get_scale_factor.
func (recv *Widget) GetScaleFactor() int32 {
	retC := C.gtk_widget_get_scale_factor((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetValignWithBaseline is a wrapper around the C function gtk_widget_get_valign_with_baseline.
func (recv *Widget) GetValignWithBaseline() Align {
	retC := C.gtk_widget_get_valign_with_baseline((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// InitTemplate is a wrapper around the C function gtk_widget_init_template.
func (recv *Widget) InitTemplate() {
	C.gtk_widget_init_template((*C.GtkWidget)(recv.native))

	return
}

// SizeAllocateWithBaseline is a wrapper around the C function gtk_widget_size_allocate_with_baseline.
func (recv *Widget) SizeAllocateWithBaseline(allocation *gdk.Rectangle, baseline int32) {
	c_allocation := (*C.GdkRectangle)(C.NULL)
	if allocation != nil {
		c_allocation = (*C.GdkRectangle)(allocation.ToC())
	}

	c_baseline := (C.gint)(baseline)

	C.gtk_widget_size_allocate_with_baseline((*C.GtkWidget)(recv.native), c_allocation, c_baseline)

	return
}

// Close is a wrapper around the C function gtk_window_close.
func (recv *Window) Close() {
	C.gtk_window_close((*C.GtkWindow)(recv.native))

	return
}

// SetTitlebar is a wrapper around the C function gtk_window_set_titlebar.
func (recv *Window) SetTitlebar(titlebar *Widget) {
	c_titlebar := (*C.GtkWidget)(C.NULL)
	if titlebar != nil {
		c_titlebar = (*C.GtkWidget)(titlebar.ToC())
	}

	C.gtk_window_set_titlebar((*C.GtkWindow)(recv.native), c_titlebar)

	return
}

type BaselinePosition C.GtkBaselinePosition

const (
	GTK_BASELINE_POSITION_TOP    BaselinePosition = 0
	GTK_BASELINE_POSITION_CENTER BaselinePosition = 1
	GTK_BASELINE_POSITION_BOTTOM BaselinePosition = 2
)

// RenderIconSurface is a wrapper around the C function gtk_render_icon_surface.
func RenderIconSurface(context *StyleContext, cr *cairo.Context, surface *cairo.Surface, x float64, y float64) {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_cr := (*C.cairo_t)(C.NULL)
	if cr != nil {
		c_cr = (*C.cairo_t)(cr.ToC())
	}

	c_surface := (*C.cairo_surface_t)(C.NULL)
	if surface != nil {
		c_surface = (*C.cairo_surface_t)(surface.ToC())
	}

	c_x := (C.gdouble)(x)

	c_y := (C.gdouble)(y)

	C.gtk_render_icon_surface(c_context, c_cr, c_surface, c_x, c_y)

	return
}

// TestWidgetWaitForDraw is a wrapper around the C function gtk_test_widget_wait_for_draw.
func TestWidgetWaitForDraw(widget *Widget) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	C.gtk_test_widget_wait_for_draw(c_widget)

	return
}

// GetCurrentName is a wrapper around the C function gtk_file_chooser_get_current_name.
func (recv *FileChooser) GetCurrentName() string {
	retC := C.gtk_file_chooser_get_current_name((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// RowsReorderedWithLength is a wrapper around the C function gtk_tree_model_rows_reordered_with_length.
func (recv *TreeModel) RowsReorderedWithLength(path *TreePath, iter *TreeIter, newOrder []int32) {
	c_path := (*C.GtkTreePath)(C.NULL)
	if path != nil {
		c_path = (*C.GtkTreePath)(path.ToC())
	}

	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	c_new_order_array := make([]C.gint, len(newOrder)+1, len(newOrder)+1)
	for i, item := range newOrder {
		c := (C.gint)(item)
		c_new_order_array[i] = c
	}
	c_new_order_array[len(newOrder)] = 0
	c_new_order_arrayPtr := &c_new_order_array[0]
	c_new_order := (*C.gint)(unsafe.Pointer(c_new_order_arrayPtr))

	c_length := (C.gint)(len(newOrder))

	C.gtk_tree_model_rows_reordered_with_length((*C.GtkTreeModel)(recv.native), c_path, c_iter, c_new_order, c_length)

	return
}

// RenderIconSurface is a wrapper around the C function gtk_icon_set_render_icon_surface.
func (recv *IconSet) RenderIconSurface(context *StyleContext, size IconSize, scale int32, forWindow *gdk.Window) *cairo.Surface {
	c_context := (*C.GtkStyleContext)(C.NULL)
	if context != nil {
		c_context = (*C.GtkStyleContext)(context.ToC())
	}

	c_size := (C.GtkIconSize)(size)

	c_scale := (C.int)(scale)

	c_for_window := (*C.GdkWindow)(C.NULL)
	if forWindow != nil {
		c_for_window = (*C.GdkWindow)(forWindow.ToC())
	}

	retC := C.gtk_icon_set_render_icon_surface((*C.GtkIconSet)(recv.native), c_context, c_size, c_scale, c_for_window)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_widget_class_bind_template_callback_full : unsupported parameter callback_symbol : no type generator for GObject.Callback (GCallback) for param callback_symbol

// BindTemplateChildFull is a wrapper around the C function gtk_widget_class_bind_template_child_full.
func (recv *WidgetClass) BindTemplateChildFull(name string, internalChild bool, structOffset int64) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_internal_child :=
		boolToGboolean(internalChild)

	c_struct_offset := (C.gssize)(structOffset)

	C.gtk_widget_class_bind_template_child_full((*C.GtkWidgetClass)(recv.native), c_name, c_internal_child, c_struct_offset)

	return
}

// Unsupported : gtk_widget_class_set_connect_func : unsupported parameter connect_func : no type generator for BuilderConnectFunc (GtkBuilderConnectFunc) for param connect_func

// SetTemplate is a wrapper around the C function gtk_widget_class_set_template.
func (recv *WidgetClass) SetTemplate(templateBytes *glib.Bytes) {
	c_template_bytes := (*C.GBytes)(C.NULL)
	if templateBytes != nil {
		c_template_bytes = (*C.GBytes)(templateBytes.ToC())
	}

	C.gtk_widget_class_set_template((*C.GtkWidgetClass)(recv.native), c_template_bytes)

	return
}

// SetTemplateFromResource is a wrapper around the C function gtk_widget_class_set_template_from_resource.
func (recv *WidgetClass) SetTemplateFromResource(resourceName string) {
	c_resource_name := C.CString(resourceName)
	defer C.free(unsafe.Pointer(c_resource_name))

	C.gtk_widget_class_set_template_from_resource((*C.GtkWidgetClass)(recv.native), c_resource_name)

	return
}