// Code generated - DO NOT EDIT.
// +build gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
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

	void cellrenderercombo_changedHandler(GObject *, gchar*, GtkTreeIter *, gpointer);

	static gulong CellRendererCombo_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(cellrenderercombo_changedHandler), data);
	}

*/
/*

	gboolean statusicon_buttonPressEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-press-event", G_CALLBACK(statusicon_buttonPressEventHandler), data);
	}

*/
/*

	gboolean statusicon_buttonReleaseEventHandler(GObject *, GdkEventButton *, gpointer);

	static gulong StatusIcon_signal_connect_button_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-release-event", G_CALLBACK(statusicon_buttonReleaseEventHandler), data);
	}

*/
/*

	gboolean widget_damageEventHandler(GObject *, GdkEventExpose *, gpointer);

	static gulong Widget_signal_connect_damage_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "damage-event", G_CALLBACK(widget_damageEventHandler), data);
	}

*/
import "C"

// GetIsLocked is a wrapper around the C function gtk_accel_group_get_is_locked.
func (recv *AccelGroup) GetIsLocked() bool {
	retC := C.gtk_accel_group_get_is_locked((*C.GtkAccelGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModifierMask is a wrapper around the C function gtk_accel_group_get_modifier_mask.
func (recv *AccelGroup) GetModifierMask() gdk.ModifierType {
	retC := C.gtk_accel_group_get_modifier_mask((*C.GtkAccelGroup)(recv.native))
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// Configure is a wrapper around the C function gtk_adjustment_configure.
func (recv *Adjustment) Configure(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_step_increment := (C.gdouble)(stepIncrement)

	c_page_increment := (C.gdouble)(pageIncrement)

	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_configure((*C.GtkAdjustment)(recv.native), c_value, c_lower, c_upper, c_step_increment, c_page_increment, c_page_size)

	return
}

// GetLower is a wrapper around the C function gtk_adjustment_get_lower.
func (recv *Adjustment) GetLower() float64 {
	retC := C.gtk_adjustment_get_lower((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPageIncrement is a wrapper around the C function gtk_adjustment_get_page_increment.
func (recv *Adjustment) GetPageIncrement() float64 {
	retC := C.gtk_adjustment_get_page_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetPageSize is a wrapper around the C function gtk_adjustment_get_page_size.
func (recv *Adjustment) GetPageSize() float64 {
	retC := C.gtk_adjustment_get_page_size((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetStepIncrement is a wrapper around the C function gtk_adjustment_get_step_increment.
func (recv *Adjustment) GetStepIncrement() float64 {
	retC := C.gtk_adjustment_get_step_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetUpper is a wrapper around the C function gtk_adjustment_get_upper.
func (recv *Adjustment) GetUpper() float64 {
	retC := C.gtk_adjustment_get_upper((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetLower is a wrapper around the C function gtk_adjustment_set_lower.
func (recv *Adjustment) SetLower(lower float64) {
	c_lower := (C.gdouble)(lower)

	C.gtk_adjustment_set_lower((*C.GtkAdjustment)(recv.native), c_lower)

	return
}

// SetPageIncrement is a wrapper around the C function gtk_adjustment_set_page_increment.
func (recv *Adjustment) SetPageIncrement(pageIncrement float64) {
	c_page_increment := (C.gdouble)(pageIncrement)

	C.gtk_adjustment_set_page_increment((*C.GtkAdjustment)(recv.native), c_page_increment)

	return
}

// SetPageSize is a wrapper around the C function gtk_adjustment_set_page_size.
func (recv *Adjustment) SetPageSize(pageSize float64) {
	c_page_size := (C.gdouble)(pageSize)

	C.gtk_adjustment_set_page_size((*C.GtkAdjustment)(recv.native), c_page_size)

	return
}

// SetStepIncrement is a wrapper around the C function gtk_adjustment_set_step_increment.
func (recv *Adjustment) SetStepIncrement(stepIncrement float64) {
	c_step_increment := (C.gdouble)(stepIncrement)

	C.gtk_adjustment_set_step_increment((*C.GtkAdjustment)(recv.native), c_step_increment)

	return
}

// SetUpper is a wrapper around the C function gtk_adjustment_set_upper.
func (recv *Adjustment) SetUpper(upper float64) {
	c_upper := (C.gdouble)(upper)

	C.gtk_adjustment_set_upper((*C.GtkAdjustment)(recv.native), c_upper)

	return
}

// AddObjectsFromFile is a wrapper around the C function gtk_builder_add_objects_from_file.
func (recv *Builder) AddObjectsFromFile(filename string, objectIds []string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	c_object_ids_array := make([]*C.gchar, len(objectIds)+1, len(objectIds)+1)
	for i, item := range objectIds {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_object_ids_array[i] = c
	}
	c_object_ids_array[len(objectIds)] = nil
	c_object_ids_arrayPtr := &c_object_ids_array[0]
	c_object_ids := (**C.gchar)(unsafe.Pointer(c_object_ids_arrayPtr))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_objects_from_file((*C.GtkBuilder)(recv.native), c_filename, c_object_ids, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddObjectsFromString is a wrapper around the C function gtk_builder_add_objects_from_string.
func (recv *Builder) AddObjectsFromString(buffer string, length uint64, objectIds []string) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gsize)(length)

	c_object_ids_array := make([]*C.gchar, len(objectIds)+1, len(objectIds)+1)
	for i, item := range objectIds {
		c := C.CString(item)
		defer C.free(unsafe.Pointer(c))
		c_object_ids_array[i] = c
	}
	c_object_ids_array[len(objectIds)] = nil
	c_object_ids_arrayPtr := &c_object_ids_array[0]
	c_object_ids := (**C.gchar)(unsafe.Pointer(c_object_ids_arrayPtr))

	var cThrowableError *C.GError

	retC := C.gtk_builder_add_objects_from_string((*C.GtkBuilder)(recv.native), c_buffer, c_length, c_object_ids, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetDetailHeightRows is a wrapper around the C function gtk_calendar_get_detail_height_rows.
func (recv *Calendar) GetDetailHeightRows() int32 {
	retC := C.gtk_calendar_get_detail_height_rows((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDetailWidthChars is a wrapper around the C function gtk_calendar_get_detail_width_chars.
func (recv *Calendar) GetDetailWidthChars() int32 {
	retC := C.gtk_calendar_get_detail_width_chars((*C.GtkCalendar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_calendar_set_detail_func : unsupported parameter func : no type generator for CalendarDetailFunc (GtkCalendarDetailFunc) for param func

// SetDetailHeightRows is a wrapper around the C function gtk_calendar_set_detail_height_rows.
func (recv *Calendar) SetDetailHeightRows(rows int32) {
	c_rows := (C.gint)(rows)

	C.gtk_calendar_set_detail_height_rows((*C.GtkCalendar)(recv.native), c_rows)

	return
}

// SetDetailWidthChars is a wrapper around the C function gtk_calendar_set_detail_width_chars.
func (recv *Calendar) SetDetailWidthChars(chars int32) {
	c_chars := (C.gint)(chars)

	C.gtk_calendar_set_detail_width_chars((*C.GtkCalendar)(recv.native), c_chars)

	return
}

type signalCellRendererComboChangedDetail struct {
	callback  CellRendererComboSignalChangedCallback
	handlerID C.gulong
}

var signalCellRendererComboChangedId int
var signalCellRendererComboChangedMap = make(map[int]signalCellRendererComboChangedDetail)
var signalCellRendererComboChangedLock sync.RWMutex

// CellRendererComboSignalChangedCallback is a callback function for a 'changed' signal emitted from a CellRendererCombo.
type CellRendererComboSignalChangedCallback func(pathString string, newIter *TreeIter)

/*
ConnectChanged connects the callback to the 'changed' signal for the CellRendererCombo.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *CellRendererCombo) ConnectChanged(callback CellRendererComboSignalChangedCallback) int {
	signalCellRendererComboChangedLock.Lock()
	defer signalCellRendererComboChangedLock.Unlock()

	signalCellRendererComboChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRendererCombo_signal_connect_changed(instance, C.gpointer(uintptr(signalCellRendererComboChangedId)))

	detail := signalCellRendererComboChangedDetail{callback, handlerID}
	signalCellRendererComboChangedMap[signalCellRendererComboChangedId] = detail

	return signalCellRendererComboChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the CellRendererCombo.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *CellRendererCombo) DisconnectChanged(connectionID int) {
	signalCellRendererComboChangedLock.Lock()
	defer signalCellRendererComboChangedLock.Unlock()

	detail, exists := signalCellRendererComboChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererComboChangedMap, connectionID)
}

//export cellrenderercombo_changedHandler
func cellrenderercombo_changedHandler(_ *C.GObject, c_path_string *C.gchar, c_new_iter *C.GtkTreeIter, data C.gpointer) {
	signalCellRendererComboChangedLock.RLock()
	defer signalCellRendererComboChangedLock.RUnlock()

	pathString := C.GoString(c_path_string)

	newIter := TreeIterNewFromC(unsafe.Pointer(c_new_iter))

	index := int(uintptr(data))
	callback := signalCellRendererComboChangedMap[index].callback
	callback(pathString, newIter)
}

// Unsupported : gtk_clipboard_request_uris : unsupported parameter callback : no type generator for ClipboardURIReceivedFunc (GtkClipboardURIReceivedFunc) for param callback

// WaitForUris is a wrapper around the C function gtk_clipboard_wait_for_uris.
func (recv *Clipboard) WaitForUris() []string {
	retC := C.gtk_clipboard_wait_for_uris((*C.GtkClipboard)(recv.native))
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// WaitIsUrisAvailable is a wrapper around the C function gtk_clipboard_wait_is_uris_available.
func (recv *Clipboard) WaitIsUrisAvailable() bool {
	retC := C.gtk_clipboard_wait_is_uris_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetColorSelection is a wrapper around the C function gtk_color_selection_dialog_get_color_selection.
func (recv *ColorSelectionDialog) GetColorSelection() *Widget {
	retC := C.gtk_color_selection_dialog_get_color_selection((*C.GtkColorSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetButtonSensitivity is a wrapper around the C function gtk_combo_box_get_button_sensitivity.
func (recv *ComboBox) GetButtonSensitivity() SensitivityType {
	retC := C.gtk_combo_box_get_button_sensitivity((*C.GtkComboBox)(recv.native))
	retGo := (SensitivityType)(retC)

	return retGo
}

// SetButtonSensitivity is a wrapper around the C function gtk_combo_box_set_button_sensitivity.
func (recv *ComboBox) SetButtonSensitivity(sensitivity SensitivityType) {
	c_sensitivity := (C.GtkSensitivityType)(sensitivity)

	C.gtk_combo_box_set_button_sensitivity((*C.GtkComboBox)(recv.native), c_sensitivity)

	return
}

// GetFocusChild is a wrapper around the C function gtk_container_get_focus_child.
func (recv *Container) GetFocusChild() *Widget {
	retC := C.gtk_container_get_focus_child((*C.GtkContainer)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetActionArea is a wrapper around the C function gtk_dialog_get_action_area.
func (recv *Dialog) GetActionArea() *Widget {
	retC := C.gtk_dialog_get_action_area((*C.GtkDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContentArea is a wrapper around the C function gtk_dialog_get_content_area.
func (recv *Dialog) GetContentArea() *Box {
	retC := C.gtk_dialog_get_content_area((*C.GtkDialog)(recv.native))
	retGo := BoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOverwriteMode is a wrapper around the C function gtk_entry_get_overwrite_mode.
func (recv *Entry) GetOverwriteMode() bool {
	retC := C.gtk_entry_get_overwrite_mode((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTextLength is a wrapper around the C function gtk_entry_get_text_length.
func (recv *Entry) GetTextLength() uint16 {
	retC := C.gtk_entry_get_text_length((*C.GtkEntry)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// SetOverwriteMode is a wrapper around the C function gtk_entry_set_overwrite_mode.
func (recv *Entry) SetOverwriteMode(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_entry_set_overwrite_mode((*C.GtkEntry)(recv.native), c_overwrite)

	return
}

// GetFace is a wrapper around the C function gtk_font_selection_get_face.
func (recv *FontSelection) GetFace() *pango.FontFace {
	retC := C.gtk_font_selection_get_face((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFaceList is a wrapper around the C function gtk_font_selection_get_face_list.
func (recv *FontSelection) GetFaceList() *Widget {
	retC := C.gtk_font_selection_get_face_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamily is a wrapper around the C function gtk_font_selection_get_family.
func (recv *FontSelection) GetFamily() *pango.FontFamily {
	retC := C.gtk_font_selection_get_family((*C.GtkFontSelection)(recv.native))
	retGo := pango.FontFamilyNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFamilyList is a wrapper around the C function gtk_font_selection_get_family_list.
func (recv *FontSelection) GetFamilyList() *Widget {
	retC := C.gtk_font_selection_get_family_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewEntry is a wrapper around the C function gtk_font_selection_get_preview_entry.
func (recv *FontSelection) GetPreviewEntry() *Widget {
	retC := C.gtk_font_selection_get_preview_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSize is a wrapper around the C function gtk_font_selection_get_size.
func (recv *FontSelection) GetSize() int32 {
	retC := C.gtk_font_selection_get_size((*C.GtkFontSelection)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSizeEntry is a wrapper around the C function gtk_font_selection_get_size_entry.
func (recv *FontSelection) GetSizeEntry() *Widget {
	retC := C.gtk_font_selection_get_size_entry((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSizeList is a wrapper around the C function gtk_font_selection_get_size_list.
func (recv *FontSelection) GetSizeList() *Widget {
	retC := C.gtk_font_selection_get_size_list((*C.GtkFontSelection)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCancelButton is a wrapper around the C function gtk_font_selection_dialog_get_cancel_button.
func (recv *FontSelectionDialog) GetCancelButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_cancel_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOkButton is a wrapper around the C function gtk_font_selection_dialog_get_ok_button.
func (recv *FontSelectionDialog) GetOkButton() *Widget {
	retC := C.gtk_font_selection_dialog_get_ok_button((*C.GtkFontSelectionDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSVNew is a wrapper around the C function gtk_hsv_new.
func HSVNew() *HSV {
	retC := C.gtk_hsv_new()
	retGo := HSVNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSVToRgb is a wrapper around the C function gtk_hsv_to_rgb.
func HSVToRgb(h float64, s float64, v float64) (float64, float64, float64) {
	c_h := (C.gdouble)(h)

	c_s := (C.gdouble)(s)

	c_v := (C.gdouble)(v)

	var c_r C.gdouble

	var c_g C.gdouble

	var c_b C.gdouble

	C.gtk_hsv_to_rgb(c_h, c_s, c_v, &c_r, &c_g, &c_b)

	r := (float64)(c_r)

	g := (float64)(c_g)

	b := (float64)(c_b)

	return r, g, b
}

// GetColor is a wrapper around the C function gtk_hsv_get_color.
func (recv *HSV) GetColor() (float64, float64, float64) {
	var c_h C.gdouble

	var c_s C.gdouble

	var c_v C.gdouble

	C.gtk_hsv_get_color((*C.GtkHSV)(recv.native), &c_h, &c_s, &c_v)

	h := (float64)(c_h)

	s := (float64)(c_s)

	v := (float64)(c_v)

	return h, s, v
}

// GetMetrics is a wrapper around the C function gtk_hsv_get_metrics.
func (recv *HSV) GetMetrics() (int32, int32) {
	var c_size C.gint

	var c_ring_width C.gint

	C.gtk_hsv_get_metrics((*C.GtkHSV)(recv.native), &c_size, &c_ring_width)

	size := (int32)(c_size)

	ringWidth := (int32)(c_ring_width)

	return size, ringWidth
}

// IsAdjusting is a wrapper around the C function gtk_hsv_is_adjusting.
func (recv *HSV) IsAdjusting() bool {
	retC := C.gtk_hsv_is_adjusting((*C.GtkHSV)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetColor is a wrapper around the C function gtk_hsv_set_color.
func (recv *HSV) SetColor(h float64, s float64, v float64) {
	c_h := (C.double)(h)

	c_s := (C.double)(s)

	c_v := (C.double)(v)

	C.gtk_hsv_set_color((*C.GtkHSV)(recv.native), c_h, c_s, c_v)

	return
}

// SetMetrics is a wrapper around the C function gtk_hsv_set_metrics.
func (recv *HSV) SetMetrics(size int32, ringWidth int32) {
	c_size := (C.gint)(size)

	c_ring_width := (C.gint)(ringWidth)

	C.gtk_hsv_set_metrics((*C.GtkHSV)(recv.native), c_size, c_ring_width)

	return
}

// GetChildDetached is a wrapper around the C function gtk_handle_box_get_child_detached.
func (recv *HandleBox) GetChildDetached() bool {
	retC := C.gtk_handle_box_get_child_detached((*C.GtkHandleBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IconInfoNewForPixbuf is a wrapper around the C function gtk_icon_info_new_for_pixbuf.
func IconInfoNewForPixbuf(iconTheme *IconTheme, pixbuf *gdkpixbuf.Pixbuf) *IconInfo {
	c_icon_theme := (*C.GtkIconTheme)(C.NULL)
	if iconTheme != nil {
		c_icon_theme = (*C.GtkIconTheme)(iconTheme.ToC())
	}

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	retC := C.gtk_icon_info_new_for_pixbuf(c_icon_theme, c_pixbuf)
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// LookupByGicon is a wrapper around the C function gtk_icon_theme_lookup_by_gicon.
func (recv *IconTheme) LookupByGicon(icon *gio.Icon, size int32, flags IconLookupFlags) *IconInfo {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_by_gicon((*C.GtkIconTheme)(recv.native), c_icon, c_size, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ImageNewFromGicon is a wrapper around the C function gtk_image_new_from_gicon.
func ImageNewFromGicon(icon *gio.Icon, size IconSize) *Image {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.GtkIconSize)(size)

	retC := C.gtk_image_new_from_gicon(c_icon, c_size)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGicon is a wrapper around the C function gtk_image_get_gicon.
func (recv *Image) GetGicon() (*gio.Icon, int32) {
	var c_gicon *C.GIcon

	var c_size C.GtkIconSize

	C.gtk_image_get_gicon((*C.GtkImage)(recv.native), &c_gicon, &c_size)

	gicon := gio.IconNewFromC(unsafe.Pointer(c_gicon))

	size := (int32)(c_size)

	return gicon, size
}

// SetFromGicon is a wrapper around the C function gtk_image_set_from_gicon.
func (recv *Image) SetFromGicon(icon *gio.Icon, size IconSize) {
	c_icon := (*C.GIcon)(icon.ToC())

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_gicon((*C.GtkImage)(recv.native), c_icon, c_size)

	return
}

// GetBinWindow is a wrapper around the C function gtk_layout_get_bin_window.
func (recv *Layout) GetBinWindow() *gdk.Window {
	retC := C.gtk_layout_get_bin_window((*C.GtkLayout)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVisited is a wrapper around the C function gtk_link_button_get_visited.
func (recv *LinkButton) GetVisited() bool {
	retC := C.gtk_link_button_get_visited((*C.GtkLinkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetVisited is a wrapper around the C function gtk_link_button_set_visited.
func (recv *LinkButton) SetVisited(visited bool) {
	c_visited :=
		boolToGboolean(visited)

	C.gtk_link_button_set_visited((*C.GtkLinkButton)(recv.native), c_visited)

	return
}

// GetAccelPath is a wrapper around the C function gtk_menu_get_accel_path.
func (recv *Menu) GetAccelPath() string {
	retC := C.gtk_menu_get_accel_path((*C.GtkMenu)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetMonitor is a wrapper around the C function gtk_menu_get_monitor.
func (recv *Menu) GetMonitor() int32 {
	retC := C.gtk_menu_get_monitor((*C.GtkMenu)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetAccelPath is a wrapper around the C function gtk_menu_item_get_accel_path.
func (recv *MenuItem) GetAccelPath() string {
	retC := C.gtk_menu_item_get_accel_path((*C.GtkMenuItem)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetImage is a wrapper around the C function gtk_message_dialog_get_image.
func (recv *MessageDialog) GetImage() *Widget {
	retC := C.gtk_message_dialog_get_image((*C.GtkMessageDialog)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MountOperationNew is a wrapper around the C function gtk_mount_operation_new.
func MountOperationNew(parent *Window) *MountOperation {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	retC := C.gtk_mount_operation_new(c_parent)
	retGo := MountOperationNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetParent is a wrapper around the C function gtk_mount_operation_get_parent.
func (recv *MountOperation) GetParent() *Window {
	retC := C.gtk_mount_operation_get_parent((*C.GtkMountOperation)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_mount_operation_get_screen.
func (recv *MountOperation) GetScreen() *gdk.Screen {
	retC := C.gtk_mount_operation_get_screen((*C.GtkMountOperation)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IsShowing is a wrapper around the C function gtk_mount_operation_is_showing.
func (recv *MountOperation) IsShowing() bool {
	retC := C.gtk_mount_operation_is_showing((*C.GtkMountOperation)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetParent is a wrapper around the C function gtk_mount_operation_set_parent.
func (recv *MountOperation) SetParent(parent *Window) {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	C.gtk_mount_operation_set_parent((*C.GtkMountOperation)(recv.native), c_parent)

	return
}

// SetScreen is a wrapper around the C function gtk_mount_operation_set_screen.
func (recv *MountOperation) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_mount_operation_set_screen((*C.GtkMountOperation)(recv.native), c_screen)

	return
}

// LoadFile is a wrapper around the C function gtk_page_setup_load_file.
func (recv *PageSetup) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_file((*C.GtkPageSetup)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadKeyFile is a wrapper around the C function gtk_page_setup_load_key_file.
func (recv *PageSetup) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_page_setup_load_key_file((*C.GtkPageSetup)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetNumberUpLayout is a wrapper around the C function gtk_print_settings_get_number_up_layout.
func (recv *PrintSettings) GetNumberUpLayout() NumberUpLayout {
	retC := C.gtk_print_settings_get_number_up_layout((*C.GtkPrintSettings)(recv.native))
	retGo := (NumberUpLayout)(retC)

	return retGo
}

// LoadFile is a wrapper around the C function gtk_print_settings_load_file.
func (recv *PrintSettings) LoadFile(fileName string) (bool, error) {
	c_file_name := C.CString(fileName)
	defer C.free(unsafe.Pointer(c_file_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_file((*C.GtkPrintSettings)(recv.native), c_file_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadKeyFile is a wrapper around the C function gtk_print_settings_load_key_file.
func (recv *PrintSettings) LoadKeyFile(keyFile *glib.KeyFile, groupName string) (bool, error) {
	c_key_file := (*C.GKeyFile)(C.NULL)
	if keyFile != nil {
		c_key_file = (*C.GKeyFile)(keyFile.ToC())
	}

	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.gtk_print_settings_load_key_file((*C.GtkPrintSettings)(recv.native), c_key_file, c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetNumberUpLayout is a wrapper around the C function gtk_print_settings_set_number_up_layout.
func (recv *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	c_number_up_layout := (C.GtkNumberUpLayout)(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout((*C.GtkPrintSettings)(recv.native), c_number_up_layout)

	return
}

// GetMinusButton is a wrapper around the C function gtk_scale_button_get_minus_button.
func (recv *ScaleButton) GetMinusButton() *Button {
	retC := C.gtk_scale_button_get_minus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPlusButton is a wrapper around the C function gtk_scale_button_get_plus_button.
func (recv *ScaleButton) GetPlusButton() *Button {
	retC := C.gtk_scale_button_get_plus_button((*C.GtkScaleButton)(recv.native))
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPopup is a wrapper around the C function gtk_scale_button_get_popup.
func (recv *ScaleButton) GetPopup() *Widget {
	retC := C.gtk_scale_button_get_popup((*C.GtkScaleButton)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

type signalStatusIconButtonPressEventDetail struct {
	callback  StatusIconSignalButtonPressEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonPressEventId int
var signalStatusIconButtonPressEventMap = make(map[int]signalStatusIconButtonPressEventDetail)
var signalStatusIconButtonPressEventLock sync.RWMutex

// StatusIconSignalButtonPressEventCallback is a callback function for a 'button-press-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonPressEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonPressEvent connects the callback to the 'button-press-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonPressEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonPressEvent(callback StatusIconSignalButtonPressEventCallback) int {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	signalStatusIconButtonPressEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_press_event(instance, C.gpointer(uintptr(signalStatusIconButtonPressEventId)))

	detail := signalStatusIconButtonPressEventDetail{callback, handlerID}
	signalStatusIconButtonPressEventMap[signalStatusIconButtonPressEventId] = detail

	return signalStatusIconButtonPressEventId
}

/*
DisconnectButtonPressEvent disconnects a callback from the 'button-press-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonPressEvent.
*/
func (recv *StatusIcon) DisconnectButtonPressEvent(connectionID int) {
	signalStatusIconButtonPressEventLock.Lock()
	defer signalStatusIconButtonPressEventLock.Unlock()

	detail, exists := signalStatusIconButtonPressEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonPressEventMap, connectionID)
}

//export statusicon_buttonPressEventHandler
func statusicon_buttonPressEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalStatusIconButtonPressEventLock.RLock()
	defer signalStatusIconButtonPressEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonPressEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

type signalStatusIconButtonReleaseEventDetail struct {
	callback  StatusIconSignalButtonReleaseEventCallback
	handlerID C.gulong
}

var signalStatusIconButtonReleaseEventId int
var signalStatusIconButtonReleaseEventMap = make(map[int]signalStatusIconButtonReleaseEventDetail)
var signalStatusIconButtonReleaseEventLock sync.RWMutex

// StatusIconSignalButtonReleaseEventCallback is a callback function for a 'button-release-event' signal emitted from a StatusIcon.
type StatusIconSignalButtonReleaseEventCallback func(event *gdk.EventButton) bool

/*
ConnectButtonReleaseEvent connects the callback to the 'button-release-event' signal for the StatusIcon.

The returned value represents the connection, and may be passed to DisconnectButtonReleaseEvent to remove it.
*/
func (recv *StatusIcon) ConnectButtonReleaseEvent(callback StatusIconSignalButtonReleaseEventCallback) int {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	signalStatusIconButtonReleaseEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.StatusIcon_signal_connect_button_release_event(instance, C.gpointer(uintptr(signalStatusIconButtonReleaseEventId)))

	detail := signalStatusIconButtonReleaseEventDetail{callback, handlerID}
	signalStatusIconButtonReleaseEventMap[signalStatusIconButtonReleaseEventId] = detail

	return signalStatusIconButtonReleaseEventId
}

/*
DisconnectButtonReleaseEvent disconnects a callback from the 'button-release-event' signal for the StatusIcon.

The connectionID should be a value returned from a call to ConnectButtonReleaseEvent.
*/
func (recv *StatusIcon) DisconnectButtonReleaseEvent(connectionID int) {
	signalStatusIconButtonReleaseEventLock.Lock()
	defer signalStatusIconButtonReleaseEventLock.Unlock()

	detail, exists := signalStatusIconButtonReleaseEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStatusIconButtonReleaseEventMap, connectionID)
}

//export statusicon_buttonReleaseEventHandler
func statusicon_buttonReleaseEventHandler(_ *C.GObject, c_event *C.GdkEventButton, data C.gpointer) C.gboolean {
	signalStatusIconButtonReleaseEventLock.RLock()
	defer signalStatusIconButtonReleaseEventLock.RUnlock()

	event := gdk.EventButtonNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalStatusIconButtonReleaseEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// StatusIconNewFromGicon is a wrapper around the C function gtk_status_icon_new_from_gicon.
func StatusIconNewFromGicon(icon *gio.Icon) *StatusIcon {
	c_icon := (*C.GIcon)(icon.ToC())

	retC := C.gtk_status_icon_new_from_gicon(c_icon)
	retGo := StatusIconNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetGicon is a wrapper around the C function gtk_status_icon_get_gicon.
func (recv *StatusIcon) GetGicon() *gio.Icon {
	retC := C.gtk_status_icon_get_gicon((*C.GtkStatusIcon)(recv.native))
	var retGo (*gio.Icon)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.IconNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetX11WindowId is a wrapper around the C function gtk_status_icon_get_x11_window_id.
func (recv *StatusIcon) GetX11WindowId() uint32 {
	retC := C.gtk_status_icon_get_x11_window_id((*C.GtkStatusIcon)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// SetFromGicon is a wrapper around the C function gtk_status_icon_set_from_gicon.
func (recv *StatusIcon) SetFromGicon(icon *gio.Icon) {
	c_icon := (*C.GIcon)(icon.ToC())

	C.gtk_status_icon_set_from_gicon((*C.GtkStatusIcon)(recv.native), c_icon)

	return
}

// ToolbarReconfigured is a wrapper around the C function gtk_tool_item_toolbar_reconfigured.
func (recv *ToolItem) ToolbarReconfigured() {
	C.gtk_tool_item_toolbar_reconfigured((*C.GtkToolItem)(recv.native))

	return
}

// SetIconFromIconName is a wrapper around the C function gtk_tooltip_set_icon_from_icon_name.
func (recv *Tooltip) SetIconFromIconName(iconName string, size IconSize) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.GtkIconSize)(size)

	C.gtk_tooltip_set_icon_from_icon_name((*C.GtkTooltip)(recv.native), c_icon_name, c_size)

	return
}

// Unsupported : gtk_tree_selection_get_select_function : no return generator

type signalWidgetDamageEventDetail struct {
	callback  WidgetSignalDamageEventCallback
	handlerID C.gulong
}

var signalWidgetDamageEventId int
var signalWidgetDamageEventMap = make(map[int]signalWidgetDamageEventDetail)
var signalWidgetDamageEventLock sync.RWMutex

// WidgetSignalDamageEventCallback is a callback function for a 'damage-event' signal emitted from a Widget.
type WidgetSignalDamageEventCallback func(event *gdk.EventExpose) bool

/*
ConnectDamageEvent connects the callback to the 'damage-event' signal for the Widget.

The returned value represents the connection, and may be passed to DisconnectDamageEvent to remove it.
*/
func (recv *Widget) ConnectDamageEvent(callback WidgetSignalDamageEventCallback) int {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	signalWidgetDamageEventId++
	instance := C.gpointer(recv.native)
	handlerID := C.Widget_signal_connect_damage_event(instance, C.gpointer(uintptr(signalWidgetDamageEventId)))

	detail := signalWidgetDamageEventDetail{callback, handlerID}
	signalWidgetDamageEventMap[signalWidgetDamageEventId] = detail

	return signalWidgetDamageEventId
}

/*
DisconnectDamageEvent disconnects a callback from the 'damage-event' signal for the Widget.

The connectionID should be a value returned from a call to ConnectDamageEvent.
*/
func (recv *Widget) DisconnectDamageEvent(connectionID int) {
	signalWidgetDamageEventLock.Lock()
	defer signalWidgetDamageEventLock.Unlock()

	detail, exists := signalWidgetDamageEventMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalWidgetDamageEventMap, connectionID)
}

//export widget_damageEventHandler
func widget_damageEventHandler(_ *C.GObject, c_event *C.GdkEventExpose, data C.gpointer) C.gboolean {
	signalWidgetDamageEventLock.RLock()
	defer signalWidgetDamageEventLock.RUnlock()

	event := gdk.EventExposeNewFromC(unsafe.Pointer(c_event))

	index := int(uintptr(data))
	callback := signalWidgetDamageEventMap[index].callback
	retGo := callback(event)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// GetWindow is a wrapper around the C function gtk_widget_get_window.
func (recv *Widget) GetWindow() *gdk.Window {
	retC := C.gtk_widget_get_window((*C.GtkWidget)(recv.native))
	var retGo (*gdk.Window)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdk.WindowNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetDefaultWidget is a wrapper around the C function gtk_window_get_default_widget.
func (recv *Window) GetDefaultWidget() *Widget {
	retC := C.gtk_window_get_default_widget((*C.GtkWindow)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListWindows is a wrapper around the C function gtk_window_group_list_windows.
func (recv *WindowGroup) ListWindows() *glib.List {
	retC := C.gtk_window_group_list_windows((*C.GtkWindowGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : STOCK_PAGE_SETUP

// Blacklisted : STOCK_PRINT_ERROR

// Blacklisted : STOCK_PRINT_PAUSED

// Blacklisted : STOCK_PRINT_REPORT

// Blacklisted : STOCK_PRINT_WARNING

// RgbToHsv is a wrapper around the C function gtk_rgb_to_hsv.
func RgbToHsv(r float64, g float64, b float64) (float64, float64, float64) {
	c_r := (C.gdouble)(r)

	c_g := (C.gdouble)(g)

	c_b := (C.gdouble)(b)

	var c_h C.gdouble

	var c_s C.gdouble

	var c_v C.gdouble

	C.gtk_rgb_to_hsv(c_r, c_g, c_b, &c_h, &c_s, &c_v)

	h := (float64)(c_h)

	s := (float64)(c_s)

	v := (float64)(c_v)

	return h, s, v
}

// ShowUri is a wrapper around the C function gtk_show_uri.
func ShowUri(screen *gdk.Screen, uri string, timestamp uint32) (bool, error) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	c_timestamp := (C.guint32)(timestamp)

	var cThrowableError *C.GError

	retC := C.gtk_show_uri(c_screen, c_uri, c_timestamp, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// TestCreateSimpleWindow is a wrapper around the C function gtk_test_create_simple_window.
func TestCreateSimpleWindow(windowTitle string, dialogText string) *Widget {
	c_window_title := C.CString(windowTitle)
	defer C.free(unsafe.Pointer(c_window_title))

	c_dialog_text := C.CString(dialogText)
	defer C.free(unsafe.Pointer(c_dialog_text))

	retC := C.gtk_test_create_simple_window(c_window_title, c_dialog_text)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// TestFindLabel is a wrapper around the C function gtk_test_find_label.
func TestFindLabel(widget *Widget, labelPattern string) *Widget {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_label_pattern := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(c_label_pattern))

	retC := C.gtk_test_find_label(c_widget, c_label_pattern)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TestFindSibling is a wrapper around the C function gtk_test_find_sibling.
func TestFindSibling(baseWidget *Widget, widgetType gobject.Type) *Widget {
	c_base_widget := (*C.GtkWidget)(C.NULL)
	if baseWidget != nil {
		c_base_widget = (*C.GtkWidget)(baseWidget.ToC())
	}

	c_widget_type := (C.GType)(widgetType)

	retC := C.gtk_test_find_sibling(c_base_widget, c_widget_type)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TestFindWidget is a wrapper around the C function gtk_test_find_widget.
func TestFindWidget(widget *Widget, labelPattern string, widgetType gobject.Type) *Widget {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_label_pattern := C.CString(labelPattern)
	defer C.free(unsafe.Pointer(c_label_pattern))

	c_widget_type := (C.GType)(widgetType)

	retC := C.gtk_test_find_widget(c_widget, c_label_pattern, c_widget_type)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// Unsupported : gtk_test_init : unsupported parameter argcp : array length param argcp is pointer (int*)

// Unsupported : gtk_test_list_all_types : array return type :

// TestRegisterAllTypes is a wrapper around the C function gtk_test_register_all_types.
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()

	return
}

// TestSliderGetValue is a wrapper around the C function gtk_test_slider_get_value.
func TestSliderGetValue(widget *Widget) float64 {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_test_slider_get_value(c_widget)
	retGo := (float64)(retC)

	return retGo
}

// TestSliderSetPerc is a wrapper around the C function gtk_test_slider_set_perc.
func TestSliderSetPerc(widget *Widget, percentage float64) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_percentage := (C.double)(percentage)

	C.gtk_test_slider_set_perc(c_widget, c_percentage)

	return
}

// TestSpinButtonClick is a wrapper around the C function gtk_test_spin_button_click.
func TestSpinButtonClick(spinner *SpinButton, button uint32, upwards bool) bool {
	c_spinner := (*C.GtkSpinButton)(C.NULL)
	if spinner != nil {
		c_spinner = (*C.GtkSpinButton)(spinner.ToC())
	}

	c_button := (C.guint)(button)

	c_upwards :=
		boolToGboolean(upwards)

	retC := C.gtk_test_spin_button_click(c_spinner, c_button, c_upwards)
	retGo := retC == C.TRUE

	return retGo
}

// TestTextGet is a wrapper around the C function gtk_test_text_get.
func TestTextGet(widget *Widget) string {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	retC := C.gtk_test_text_get(c_widget)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// TestTextSet is a wrapper around the C function gtk_test_text_set.
func TestTextSet(widget *Widget, string_ string) {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	C.gtk_test_text_set(c_widget, c_string)

	return
}

// TestWidgetClick is a wrapper around the C function gtk_test_widget_click.
func TestWidgetClick(widget *Widget, button uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_button := (C.guint)(button)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_click(c_widget, c_button, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// TestWidgetSendKey is a wrapper around the C function gtk_test_widget_send_key.
func TestWidgetSendKey(widget *Widget, keyval uint32, modifiers gdk.ModifierType) bool {
	c_widget := (*C.GtkWidget)(C.NULL)
	if widget != nil {
		c_widget = (*C.GtkWidget)(widget.ToC())
	}

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	retC := C.gtk_test_widget_send_key(c_widget, c_keyval, c_modifiers)
	retGo := retC == C.TRUE

	return retGo
}

// GetCurrentFolderFile is a wrapper around the C function gtk_file_chooser_get_current_folder_file.
func (recv *FileChooser) GetCurrentFolderFile() *gio.File {
	retC := C.gtk_file_chooser_get_current_folder_file((*C.GtkFileChooser)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFile is a wrapper around the C function gtk_file_chooser_get_file.
func (recv *FileChooser) GetFile() *gio.File {
	retC := C.gtk_file_chooser_get_file((*C.GtkFileChooser)(recv.native))
	retGo := gio.FileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFiles is a wrapper around the C function gtk_file_chooser_get_files.
func (recv *FileChooser) GetFiles() *glib.SList {
	retC := C.gtk_file_chooser_get_files((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewFile is a wrapper around the C function gtk_file_chooser_get_preview_file.
func (recv *FileChooser) GetPreviewFile() *gio.File {
	retC := C.gtk_file_chooser_get_preview_file((*C.GtkFileChooser)(recv.native))
	var retGo (*gio.File)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gio.FileNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// SelectFile is a wrapper around the C function gtk_file_chooser_select_file.
func (recv *FileChooser) SelectFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_select_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCurrentFolderFile is a wrapper around the C function gtk_file_chooser_set_current_folder_file.
func (recv *FileChooser) SetCurrentFolderFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_set_current_folder_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetFile is a wrapper around the C function gtk_file_chooser_set_file.
func (recv *FileChooser) SetFile(file *gio.File) (bool, error) {
	c_file := (*C.GFile)(file.ToC())

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_set_file((*C.GtkFileChooser)(recv.native), c_file, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnselectFile is a wrapper around the C function gtk_file_chooser_unselect_file.
func (recv *FileChooser) UnselectFile(file *gio.File) {
	c_file := (*C.GFile)(file.ToC())

	C.gtk_file_chooser_unselect_file((*C.GtkFileChooser)(recv.native), c_file)

	return
}

// GetIconSize is a wrapper around the C function gtk_tool_shell_get_icon_size.
func (recv *ToolShell) GetIconSize() IconSize {
	retC := C.gtk_tool_shell_get_icon_size((*C.GtkToolShell)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_tool_shell_get_orientation.
func (recv *ToolShell) GetOrientation() Orientation {
	retC := C.gtk_tool_shell_get_orientation((*C.GtkToolShell)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_tool_shell_get_relief_style.
func (recv *ToolShell) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_shell_get_relief_style((*C.GtkToolShell)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function gtk_tool_shell_get_style.
func (recv *ToolShell) GetStyle() ToolbarStyle {
	retC := C.gtk_tool_shell_get_style((*C.GtkToolShell)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// RebuildMenu is a wrapper around the C function gtk_tool_shell_rebuild_menu.
func (recv *ToolShell) RebuildMenu() {
	C.gtk_tool_shell_rebuild_menu((*C.GtkToolShell)(recv.native))

	return
}

// BorderNew is a wrapper around the C function gtk_border_new.
func BorderNew() *Border {
	retC := C.gtk_border_new()
	retGo := BorderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_selection_data_get_data : array return type :

// GetDataType is a wrapper around the C function gtk_selection_data_get_data_type.
func (recv *SelectionData) GetDataType() gdk.Atom {
	retC := C.gtk_selection_data_get_data_type((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDisplay is a wrapper around the C function gtk_selection_data_get_display.
func (recv *SelectionData) GetDisplay() *gdk.Display {
	retC := C.gtk_selection_data_get_display((*C.GtkSelectionData)(recv.native))
	retGo := gdk.DisplayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFormat is a wrapper around the C function gtk_selection_data_get_format.
func (recv *SelectionData) GetFormat() int32 {
	retC := C.gtk_selection_data_get_format((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLength is a wrapper around the C function gtk_selection_data_get_length.
func (recv *SelectionData) GetLength() int32 {
	retC := C.gtk_selection_data_get_length((*C.GtkSelectionData)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTarget is a wrapper around the C function gtk_selection_data_get_target.
func (recv *SelectionData) GetTarget() gdk.Atom {
	retC := C.gtk_selection_data_get_target((*C.GtkSelectionData)(recv.native))
	retGo := *gdk.AtomNewFromC(unsafe.Pointer(retC))

	return retGo
}