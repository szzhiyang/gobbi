// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

// Unsupported : gtk_assistant_new : return type :

// Unsupported : gtk_cell_renderer_accel_new : return type :

// Unsupported : gtk_cell_renderer_spin_new : return type :

// Unsupported : gtk_link_button_new : return type :

// Unsupported : gtk_link_button_new_with_label : return type :

// Unsupported : gtk_page_setup_new : return type :

// Unsupported : gtk_print_operation_new : return type :

// Unsupported : gtk_print_settings_new : return type :

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_menu_new : return type :

// Unsupported : gtk_recent_chooser_menu_new_for_manager : return type :

// Unsupported : gtk_recent_chooser_widget_new : return type :

// Unsupported : gtk_recent_chooser_widget_new_for_manager : return type :

// Unsupported : gtk_recent_filter_new : return type :

// Unsupported : gtk_recent_manager_new : return type :

// Unsupported : gtk_status_icon_new : return type :

// Unsupported : gtk_status_icon_new_from_file : return type :

// Unsupported : gtk_status_icon_new_from_icon_name : return type :

// Unsupported : gtk_status_icon_new_from_pixbuf : return type :

// Unsupported : gtk_status_icon_new_from_stock : return type :

// Blacklisted : STOCK_ORIENTATION_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_PORTRAIT

// Blacklisted : STOCK_ORIENTATION_REVERSE_LANDSCAPE

// Blacklisted : STOCK_ORIENTATION_REVERSE_PORTRAIT

// Blacklisted : STOCK_SELECT_ALL

type RecentChooserError int

const (
	GTK_RECENT_CHOOSER_ERROR_NOT_FOUND   RecentChooserError = 0
	GTK_RECENT_CHOOSER_ERROR_INVALID_URI RecentChooserError = 1
)

type RecentManagerError int

const (
	GTK_RECENT_MANAGER_ERROR_NOT_FOUND        RecentManagerError = 0
	GTK_RECENT_MANAGER_ERROR_INVALID_URI      RecentManagerError = 1
	GTK_RECENT_MANAGER_ERROR_INVALID_ENCODING RecentManagerError = 2
	GTK_RECENT_MANAGER_ERROR_NOT_REGISTERED   RecentManagerError = 3
	GTK_RECENT_MANAGER_ERROR_READ             RecentManagerError = 4
	GTK_RECENT_MANAGER_ERROR_WRITE            RecentManagerError = 5
	GTK_RECENT_MANAGER_ERROR_UNKNOWN          RecentManagerError = 6
)

type RecentSortType int

const (
	GTK_RECENT_SORT_NONE   RecentSortType = 0
	GTK_RECENT_SORT_MRU    RecentSortType = 1
	GTK_RECENT_SORT_LRU    RecentSortType = 2
	GTK_RECENT_SORT_CUSTOM RecentSortType = 3
)

// Unsupported : gtk_paper_size_get_default : return type :

// Unsupported : gtk_print_error_quark : return type :

// Unsupported : gtk_print_run_page_setup_dialog : return type :

// Unsupported : gtk_print_run_page_setup_dialog_async : unsupported parameter done_cb : no type generator for PageSetupDoneFunc (GtkPageSetupDoneFunc) for param done_cb

// Unsupported : gtk_target_table_free : unsupported parameter targets :

// Unsupported : gtk_target_table_new_from_list : array return type :

// Unsupported : gtk_targets_include_image : unsupported parameter targets :

// Unsupported : gtk_targets_include_rich_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_text : unsupported parameter targets :

// Unsupported : gtk_targets_include_uri : unsupported parameter targets :

// Unsupported : gtk_paper_size_new : return type :

// Unsupported : gtk_paper_size_new_custom : return type :

// Unsupported : gtk_paper_size_new_from_ppd : return type :
