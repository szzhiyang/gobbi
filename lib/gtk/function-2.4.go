// This is a generated file - DO NOT EDIT
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/gdk"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_accelerator_parse : unsupported parameter accelerator_mods : GdkModifierType* with indirection level of 1

// Unsupported : gtk_accelerator_parse_with_keycode : unsupported parameter accelerator_codes : output array param accelerator_codes

// BindingsActivateEvent is a wrapper around the C function gtk_bindings_activate_event.
func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) bool {
	c_object := (*C.GObject)(object.ToC())

	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_bindings_activate_event(c_object, c_event)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_drag_set_icon_gicon : unsupported parameter icon : no type generator for Gio.Icon (GIcon*) for param icon

// Unsupported : gtk_draw_insertion_cursor : unsupported parameter location : Blacklisted record : GdkRectangle

// FileChooserErrorQuark is a wrapper around the C function gtk_file_chooser_error_quark.
func FileChooserErrorQuark() glib.Quark {
	retC := C.gtk_file_chooser_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

// Unsupported : gtk_get_current_event : no return generator

// Unsupported : gtk_get_current_event_state : unsupported parameter state : GdkModifierType* with indirection level of 1

// Unsupported : gtk_get_event_widget : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_icon_size_from_name : no return generator

// Unsupported : gtk_icon_size_get_name : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_lookup_for_settings : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_icon_size_register : no return generator

// Unsupported : gtk_icon_size_register_alias : unsupported parameter target : no type generator for gint (GtkIconSize) for param target

// Unsupported : gtk_init_with_args : unsupported parameter entries :

// Unsupported : gtk_key_snooper_install : unsupported parameter snooper : no type generator for KeySnoopFunc (GtkKeySnoopFunc) for param snooper

// Unsupported : gtk_main_do_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_propagate_event : unsupported parameter event : no type generator for Gdk.Event (GdkEvent*) for param event

// Unsupported : gtk_rc_get_default_files : no return type

// Unsupported : gtk_rc_parse_priority : unsupported parameter priority : GtkPathPriorityType* with indirection level of 1

// Unsupported : gtk_rc_parse_state : unsupported parameter state : GtkStateType* with indirection level of 1

// Unsupported : gtk_rc_property_parse_border : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_color : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_enum : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_flags : unsupported parameter pspec : Blacklisted record : GParamSpec

// Unsupported : gtk_rc_property_parse_requisition : unsupported parameter pspec : Blacklisted record : GParamSpec

// RcResetStyles is a wrapper around the C function gtk_rc_reset_styles.
func RcResetStyles(settings *Settings) {
	c_settings := (*C.GtkSettings)(settings.ToC())

	C.gtk_rc_reset_styles(c_settings)

	return
}

// Unsupported : gtk_rc_set_default_files : unsupported parameter filenames :

// Unsupported : gtk_render_background_get_clip : unsupported parameter out_clip : Blacklisted record : GdkRectangle

// Unsupported : gtk_render_icon_pixbuf : unsupported parameter size : no type generator for gint (GtkIconSize) for param size

// Unsupported : gtk_selection_add_target : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_add_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_clear_targets : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_convert : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_selection_owner_set_for_display : unsupported parameter selection : Blacklisted record : GdkAtom

// Unsupported : gtk_show_about_dialog : unsupported parameter ... : varargs

// Unsupported : gtk_stock_add : unsupported parameter items :

// Unsupported : gtk_stock_add_static : unsupported parameter items :

// Unsupported : gtk_stock_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// Unsupported : gtk_target_table_free : unsupported parameter targets :

// Unsupported : gtk_target_table_new_from_list : no return type

// Unsupported : gtk_targets_include_image : unsupported parameter targets :

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_uri : unsupported parameter targets :

// Unsupported : gtk_test_create_widget : unsupported parameter ... : varargs

// Unsupported : gtk_test_display_button_window : unsupported parameter ... : varargs

// Unsupported : gtk_test_init : unsupported parameter argvp :

// Unsupported : gtk_test_list_all_types : no return type

// Unsupported : gtk_tree_get_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel**) for param tree_model

// Unsupported : gtk_tree_set_row_drag_data : unsupported parameter tree_model : no type generator for TreeModel (GtkTreeModel*) for param tree_model
