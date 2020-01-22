// Code generated - DO NOT EDIT.
// +build gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29 gtk_3.24

package gtk

import "unsafe"

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
// UNSUPPORTED : GestureStylusClass : blacklisted
// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
// UNSUPPORTED : StackAccessibleClass : blacklisted
func Fn_gtk_container_class_install_child_properties(paramInstance unsafe.Pointer, param0 uint, param1 []unsafe.Pointer) {
	cValueInstance := (*C.GtkContainerClass)(unsafe.Pointer(paramInstance))

	cValue0 := (C.guint)(param0)

	cValue1 := (**C.GParamSpec)(unsafe.Pointer(&param1[0]))

	C.gtk_container_class_install_child_properties(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

// UNSUPPORTED : gtk_selection_data_set_uris : parameter 'uris' is array parameter without length parameter

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

// UNSUPPORTED : gtk_accelerator_parse_with_keycode : parameter 'accelerator_codes' is array parameter without length parameter

// UNSUPPORTED : gtk_init_with_args : parameter 'entries' is array parameter without length parameter

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

// UNSUPPORTED : gtk_rc_get_default_files : no array length

// UNSUPPORTED : gtk_rc_set_default_files : parameter 'filenames' is array parameter without length parameter

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_row_reference_reordered : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_add_credit_section : parameter 'people' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

// UNSUPPORTED : gtk_about_dialog_set_artists : parameter 'artists' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_authors : parameter 'authors' is array parameter without length parameter

// UNSUPPORTED : gtk_about_dialog_set_documenters : parameter 'documenters' is array parameter without length parameter

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

// UNSUPPORTED : gtk_application_set_accels_for_action : parameter 'accels' is array parameter without length parameter

func Fn_gtk_assistant_get_page_has_padding(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_assistant_get_page_has_padding(cValueInstance, cValue0)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

func Fn_gtk_assistant_set_page_has_padding(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkAssistant)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_assistant_set_page_has_padding(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_objects_from_file : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_resource : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_add_objects_from_string : parameter 'object_ids' is array parameter without length parameter

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_container_child_notify_by_pspec(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkContainer)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (*C.GParamSpec)(unsafe.Pointer(param1))

	C.gtk_container_child_notify_by_pspec(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_gesture_stylus_get_axes : parameter 'axes' is array parameter without length parameter

// UNSUPPORTED : gtk_im_context_simple_add_table : parameter 'data' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_choose_icon : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_choose_icon_for_scale : parameter 'icon_names' is array parameter without length parameter

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_list_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_overlay_get_overlay_pass_through(paramInstance unsafe.Pointer, param0 unsafe.Pointer) bool {
	cValueInstance := (*C.GtkOverlay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	ret := C.gtk_overlay_get_overlay_pass_through(cValueInstance, cValue0)

	return toGoBool(ret)
}

func Fn_gtk_overlay_reorder_overlay(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkOverlay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := (C.int)(param1)

	C.gtk_overlay_reorder_overlay(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_overlay_set_overlay_pass_through(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 bool) {
	cValueInstance := (*C.GtkOverlay)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	cValue1 := toCBool(param1)

	C.gtk_overlay_set_overlay_pass_through(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_places_sidebar_get_show_other_locations(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_other_locations(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_get_show_recent(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_recent(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_get_show_trash(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	ret := C.gtk_places_sidebar_get_show_trash(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_places_sidebar_set_drop_targets_visible(paramInstance unsafe.Pointer, param0 bool, param1 unsafe.Pointer) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	cValue1 := (*C.GdkDragContext)(unsafe.Pointer(param1))

	C.gtk_places_sidebar_set_drop_targets_visible(cValueInstance, cValue0, cValue1)
}

func Fn_gtk_places_sidebar_set_show_other_locations(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_other_locations(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_show_recent(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_recent(cValueInstance, cValue0)
}

func Fn_gtk_places_sidebar_set_show_trash(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkPlacesSidebar)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_places_sidebar_set_show_trash(cValueInstance, cValue0)
}

func Fn_gtk_popover_get_default_widget(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	ret := C.gtk_popover_get_default_widget(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_popover_set_default_widget(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkPopover)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkWidget)(unsafe.Pointer(param0))

	C.gtk_popover_set_default_widget(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_radio_menu_item_join_group(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkRadioMenuItem)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GtkRadioMenuItem)(unsafe.Pointer(param0))

	C.gtk_radio_menu_item_join_group(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

// UNSUPPORTED : gtk_scale_button_new : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_scale_button_set_icons : parameter 'icons' is array parameter without length parameter

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_stack_get_interpolate_size(paramInstance unsafe.Pointer) bool {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	ret := C.gtk_stack_get_interpolate_size(cValueInstance)

	return toGoBool(ret)
}

func Fn_gtk_stack_set_interpolate_size(paramInstance unsafe.Pointer, param0 bool) {
	cValueInstance := (*C.GtkStack)(unsafe.Pointer(paramInstance))

	cValue0 := toCBool(param0)

	C.gtk_stack_set_interpolate_size(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_view_get_bottom_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_bottom_margin(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_get_top_margin(paramInstance unsafe.Pointer) int {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	ret := C.gtk_text_view_get_top_margin(cValueInstance)

	return (int)(ret)
}

func Fn_gtk_text_view_set_bottom_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_bottom_margin(cValueInstance, cValue0)
}

func Fn_gtk_text_view_set_top_margin(paramInstance unsafe.Pointer, param0 int) {
	cValueInstance := (*C.GtkTextView)(unsafe.Pointer(paramInstance))

	cValue0 := (C.gint)(param0)

	C.gtk_text_view_set_top_margin(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_store_reorder : parameter 'new_order' is array parameter without length parameter

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_get_font_map(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_font_map(cValueInstance)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_font_options(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	ret := C.gtk_widget_get_font_options(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_set_font_map(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontMap)(unsafe.Pointer(param0))

	C.gtk_widget_set_font_map(cValueInstance, cValue0)
}

func Fn_gtk_widget_set_font_options(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkWidget)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.cairo_font_options_t)(unsafe.Pointer(param0))

	C.gtk_widget_set_font_options(cValueInstance, cValue0)
}

func Fn_gtk_window_fullscreen_on_monitor(paramInstance unsafe.Pointer, param0 unsafe.Pointer, param1 int) {
	cValueInstance := (*C.GtkWindow)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.GdkScreen)(unsafe.Pointer(param0))

	cValue1 := (C.gint)(param1)

	C.gtk_window_fullscreen_on_monitor(cValueInstance, cValue0, cValue1)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_file_chooser_add_choice : parameter 'options' is array parameter without length parameter

func Fn_gtk_font_chooser_get_font_map(paramInstance unsafe.Pointer) unsafe.Pointer {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	ret := C.gtk_font_chooser_get_font_map(cValueInstance)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

func Fn_gtk_font_chooser_set_font_map(paramInstance unsafe.Pointer, param0 unsafe.Pointer) {
	cValueInstance := (*C.GtkFontChooser)(unsafe.Pointer(paramInstance))

	cValue0 := (*C.PangoFontMap)(unsafe.Pointer(param0))

	C.gtk_font_chooser_set_font_map(cValueInstance, cValue0)
}

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback