// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gio "github.com/pekim/gobbi/lib/gio"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// GetMinimumIncrement is a wrapper around the C function gtk_adjustment_get_minimum_increment.
func (recv *Adjustment) GetMinimumIncrement() float64 {
	retC := C.gtk_adjustment_get_minimum_increment((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetShowDefaultItem is a wrapper around the C function gtk_app_chooser_button_get_show_default_item.
func (recv *AppChooserButton) GetShowDefaultItem() bool {
	retC := C.gtk_app_chooser_button_get_show_default_item((*C.GtkAppChooserButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetShowDefaultItem is a wrapper around the C function gtk_app_chooser_button_set_show_default_item.
func (recv *AppChooserButton) SetShowDefaultItem(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_app_chooser_button_set_show_default_item((*C.GtkAppChooserButton)(recv.native), c_setting)

	return
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// RemovePage is a wrapper around the C function gtk_assistant_remove_page.
func (recv *Assistant) RemovePage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_assistant_remove_page((*C.GtkAssistant)(recv.native), c_page_num)

	return
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// GetChildNonHomogeneous is a wrapper around the C function gtk_button_box_get_child_non_homogeneous.
func (recv *ButtonBox) GetChildNonHomogeneous(child *Widget) bool {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_button_box_get_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// SetChildNonHomogeneous is a wrapper around the C function gtk_button_box_set_child_non_homogeneous.
func (recv *ButtonBox) SetChildNonHomogeneous(child *Widget, nonHomogeneous bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_non_homogeneous :=
		boolToGboolean(nonHomogeneous)

	C.gtk_button_box_set_child_non_homogeneous((*C.GtkButtonBox)(recv.native), c_child, c_non_homogeneous)

	return
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// ChildNotify is a wrapper around the C function gtk_container_child_notify.
func (recv *Container) ChildNotify(child *Widget, childProperty string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_child_property := C.CString(childProperty)
	defer C.free(unsafe.Pointer(c_child_property))

	C.gtk_container_child_notify((*C.GtkContainer)(recv.native), c_child, c_child_property)

	return
}

// ToString is a wrapper around the C function gtk_css_provider_to_string.
func (recv *CssProvider) ToString() string {
	retC := C.gtk_css_provider_to_string((*C.GtkCssProvider)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// GetPlaceholderText is a wrapper around the C function gtk_entry_get_placeholder_text.
func (recv *Entry) GetPlaceholderText() string {
	retC := C.gtk_entry_get_placeholder_text((*C.GtkEntry)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetPlaceholderText is a wrapper around the C function gtk_entry_set_placeholder_text.
func (recv *Entry) SetPlaceholderText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_set_placeholder_text((*C.GtkEntry)(recv.native), c_text)

	return
}

// Unsupported : EntryIconAccessible : no CType

// GetResizeToplevel is a wrapper around the C function gtk_expander_get_resize_toplevel.
func (recv *Expander) GetResizeToplevel() bool {
	retC := C.gtk_expander_get_resize_toplevel((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetResizeToplevel is a wrapper around the C function gtk_expander_set_resize_toplevel.
func (recv *Expander) SetResizeToplevel(resizeToplevel bool) {
	c_resize_toplevel :=
		boolToGboolean(resizeToplevel)

	C.gtk_expander_set_resize_toplevel((*C.GtkExpander)(recv.native), c_resize_toplevel)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// FontChooserDialogNew is a wrapper around the C function gtk_font_chooser_dialog_new.
func FontChooserDialogNew(title string, parent *Window) *FontChooserDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	c_parent := (*C.GtkWindow)(parent.ToC())

	retC := C.gtk_font_chooser_dialog_new(c_title, c_parent)
	retGo := FontChooserDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontChooserWidgetNew is a wrapper around the C function gtk_font_chooser_widget_new.
func FontChooserWidgetNew() *FontChooserWidget {
	retC := C.gtk_font_chooser_widget_new()
	retGo := FontChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetChildAt is a wrapper around the C function gtk_grid_get_child_at.
func (recv *Grid) GetChildAt(left int32, top int32) *Widget {
	c_left := (C.gint)(left)

	c_top := (C.gint)(top)

	retC := C.gtk_grid_get_child_at((*C.GtkGrid)(recv.native), c_left, c_top)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertColumn is a wrapper around the C function gtk_grid_insert_column.
func (recv *Grid) InsertColumn(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_column((*C.GtkGrid)(recv.native), c_position)

	return
}

// InsertNextTo is a wrapper around the C function gtk_grid_insert_next_to.
func (recv *Grid) InsertNextTo(sibling *Widget, side PositionType) {
	c_sibling := (*C.GtkWidget)(sibling.ToC())

	c_side := (C.GtkPositionType)(side)

	C.gtk_grid_insert_next_to((*C.GtkGrid)(recv.native), c_sibling, c_side)

	return
}

// InsertRow is a wrapper around the C function gtk_grid_insert_row.
func (recv *Grid) InsertRow(position int32) {
	c_position := (C.gint)(position)

	C.gtk_grid_insert_row((*C.GtkGrid)(recv.native), c_position)

	return
}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// LockButtonNew is a wrapper around the C function gtk_lock_button_new.
func LockButtonNew(permission *gio.Permission) *LockButton {
	c_permission := (*C.GPermission)(permission.ToC())

	retC := C.gtk_lock_button_new(c_permission)
	retGo := LockButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPermission is a wrapper around the C function gtk_lock_button_get_permission.
func (recv *LockButton) GetPermission() *gio.Permission {
	retC := C.gtk_lock_button_get_permission((*C.GtkLockButton)(recv.native))
	retGo := gio.PermissionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetPermission is a wrapper around the C function gtk_lock_button_set_permission.
func (recv *LockButton) SetPermission(permission *gio.Permission) {
	c_permission := (*C.GPermission)(permission.ToC())

	C.gtk_lock_button_set_permission((*C.GtkLockButton)(recv.native), c_permission)

	return
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// OverlayNew is a wrapper around the C function gtk_overlay_new.
func OverlayNew() *Overlay {
	retC := C.gtk_overlay_new()
	retGo := OverlayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddOverlay is a wrapper around the C function gtk_overlay_add_overlay.
func (recv *Overlay) AddOverlay(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_overlay_add_overlay((*C.GtkOverlay)(recv.native), c_widget)

	return
}

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// GetXOffset is a wrapper around the C function gtk_tree_view_column_get_x_offset.
func (recv *TreeViewColumn) GetXOffset() int32 {
	retC := C.gtk_tree_view_column_get_x_offset((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Unsupported : gtk_drag_source_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// HasVisibleFocus is a wrapper around the C function gtk_widget_has_visible_focus.
func (recv *Widget) HasVisibleFocus() bool {
	retC := C.gtk_widget_has_visible_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFocusVisible is a wrapper around the C function gtk_window_get_focus_visible.
func (recv *Window) GetFocusVisible() bool {
	retC := C.gtk_window_get_focus_visible((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFocusVisible is a wrapper around the C function gtk_window_set_focus_visible.
func (recv *Window) SetFocusVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_focus_visible((*C.GtkWindow)(recv.native), c_setting)

	return
}