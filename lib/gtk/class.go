// This is a generated file - DO NOT EDIT

package gtk

import (
	atk "github.com/pekim/gobbi/lib/atk"
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
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void AccelGroup_accelActivateHandler();

	static gulong AccelGroup_signal_connect_accel_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-activate", AccelGroup_accelActivateHandler, data);
	}

*/
/*

	void AccelGroup_accelChangedHandler();

	static gulong AccelGroup_signal_connect_accel_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-changed", AccelGroup_accelChangedHandler, data);
	}

*/
/*

	void Adjustment_changedHandler();

	static gulong Adjustment_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", Adjustment_changedHandler, data);
	}

*/
/*

	void Adjustment_valueChangedHandler();

	static gulong Adjustment_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", Adjustment_valueChangedHandler, data);
	}

*/
/*

	void AppChooserButton_customItemActivatedHandler();

	static gulong AppChooserButton_signal_connect_custom_item_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "custom-item-activated", AppChooserButton_customItemActivatedHandler, data);
	}

*/
/*

	void Assistant_escapeHandler();

	static gulong Assistant_signal_connect_escape(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "escape", Assistant_escapeHandler, data);
	}

*/
/*

	void Button_activateHandler();

	static gulong Button_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", Button_activateHandler, data);
	}

*/
/*

	void Button_clickedHandler();

	static gulong Button_signal_connect_clicked(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "clicked", Button_clickedHandler, data);
	}

*/
/*

	void Button_enterHandler();

	static gulong Button_signal_connect_enter(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enter", Button_enterHandler, data);
	}

*/
/*

	void Button_leaveHandler();

	static gulong Button_signal_connect_leave(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "leave", Button_leaveHandler, data);
	}

*/
/*

	void Button_pressedHandler();

	static gulong Button_signal_connect_pressed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pressed", Button_pressedHandler, data);
	}

*/
/*

	void Button_releasedHandler();

	static gulong Button_signal_connect_released(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "released", Button_releasedHandler, data);
	}

*/
/*

	void Calendar_daySelectedHandler();

	static gulong Calendar_signal_connect_day_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "day-selected", Calendar_daySelectedHandler, data);
	}

*/
/*

	void Calendar_daySelectedDoubleClickHandler();

	static gulong Calendar_signal_connect_day_selected_double_click(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "day-selected-double-click", Calendar_daySelectedDoubleClickHandler, data);
	}

*/
/*

	void Calendar_monthChangedHandler();

	static gulong Calendar_signal_connect_month_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "month-changed", Calendar_monthChangedHandler, data);
	}

*/
/*

	void Calendar_nextMonthHandler();

	static gulong Calendar_signal_connect_next_month(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "next-month", Calendar_nextMonthHandler, data);
	}

*/
/*

	void Calendar_nextYearHandler();

	static gulong Calendar_signal_connect_next_year(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "next-year", Calendar_nextYearHandler, data);
	}

*/
/*

	void Calendar_prevMonthHandler();

	static gulong Calendar_signal_connect_prev_month(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "prev-month", Calendar_prevMonthHandler, data);
	}

*/
/*

	void Calendar_prevYearHandler();

	static gulong Calendar_signal_connect_prev_year(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "prev-year", Calendar_prevYearHandler, data);
	}

*/
/*

	void CellRendererText_editedHandler();

	static gulong CellRendererText_signal_connect_edited(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "edited", CellRendererText_editedHandler, data);
	}

*/
/*

	void CellRendererToggle_toggledHandler();

	static gulong CellRendererToggle_signal_connect_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggled", CellRendererToggle_toggledHandler, data);
	}

*/
/*

	void CheckMenuItem_toggledHandler();

	static gulong CheckMenuItem_signal_connect_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggled", CheckMenuItem_toggledHandler, data);
	}

*/
/*

	void ColorSelection_colorChangedHandler();

	static gulong ColorSelection_signal_connect_color_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-changed", ColorSelection_colorChangedHandler, data);
	}

*/
/*

	void Container_addHandler();

	static gulong Container_signal_connect_add(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add", Container_addHandler, data);
	}

*/
/*

	void Container_checkResizeHandler();

	static gulong Container_signal_connect_check_resize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "check-resize", Container_checkResizeHandler, data);
	}

*/
/*

	void Container_removeHandler();

	static gulong Container_signal_connect_remove(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove", Container_removeHandler, data);
	}

*/
/*

	void Container_setFocusChildHandler();

	static gulong Container_signal_connect_set_focus_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "set-focus-child", Container_setFocusChildHandler, data);
	}

*/
/*

	void CssProvider_parsingErrorHandler();

	static gulong CssProvider_signal_connect_parsing_error(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "parsing-error", CssProvider_parsingErrorHandler, data);
	}

*/
/*

	void Dialog_closeHandler();

	static gulong Dialog_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", Dialog_closeHandler, data);
	}

*/
/*

	void Dialog_responseHandler();

	static gulong Dialog_signal_connect_response(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "response", Dialog_responseHandler, data);
	}

*/
/*

	void Entry_activateHandler();

	static gulong Entry_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", Entry_activateHandler, data);
	}

*/
/*

	void Entry_backspaceHandler();

	static gulong Entry_signal_connect_backspace(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "backspace", Entry_backspaceHandler, data);
	}

*/
/*

	void Entry_copyClipboardHandler();

	static gulong Entry_signal_connect_copy_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "copy-clipboard", Entry_copyClipboardHandler, data);
	}

*/
/*

	void Entry_cutClipboardHandler();

	static gulong Entry_signal_connect_cut_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cut-clipboard", Entry_cutClipboardHandler, data);
	}

*/
/*

	void Entry_deleteFromCursorHandler();

	static gulong Entry_signal_connect_delete_from_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "delete-from-cursor", Entry_deleteFromCursorHandler, data);
	}

*/
/*

	void Entry_insertAtCursorHandler();

	static gulong Entry_signal_connect_insert_at_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-at-cursor", Entry_insertAtCursorHandler, data);
	}

*/
/*

	void Entry_moveCursorHandler();

	static gulong Entry_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", Entry_moveCursorHandler, data);
	}

*/
/*

	void Entry_pasteClipboardHandler();

	static gulong Entry_signal_connect_paste_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paste-clipboard", Entry_pasteClipboardHandler, data);
	}

*/
/*

	void Entry_populatePopupHandler();

	static gulong Entry_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", Entry_populatePopupHandler, data);
	}

*/
/*

	void Entry_toggleOverwriteHandler();

	static gulong Entry_signal_connect_toggle_overwrite(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-overwrite", Entry_toggleOverwriteHandler, data);
	}

*/
/*

	void Expander_activateHandler();

	static gulong Expander_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", Expander_activateHandler, data);
	}

*/
/*

	void FileChooserWidget_desktopFolderHandler();

	static gulong FileChooserWidget_signal_connect_desktop_folder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "desktop-folder", FileChooserWidget_desktopFolderHandler, data);
	}

*/
/*

	void FileChooserWidget_downFolderHandler();

	static gulong FileChooserWidget_signal_connect_down_folder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "down-folder", FileChooserWidget_downFolderHandler, data);
	}

*/
/*

	void FileChooserWidget_homeFolderHandler();

	static gulong FileChooserWidget_signal_connect_home_folder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "home-folder", FileChooserWidget_homeFolderHandler, data);
	}

*/
/*

	void FileChooserWidget_locationPopupHandler();

	static gulong FileChooserWidget_signal_connect_location_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "location-popup", FileChooserWidget_locationPopupHandler, data);
	}

*/
/*

	void FileChooserWidget_locationPopupOnPasteHandler();

	static gulong FileChooserWidget_signal_connect_location_popup_on_paste(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "location-popup-on-paste", FileChooserWidget_locationPopupOnPasteHandler, data);
	}

*/
/*

	void FileChooserWidget_locationTogglePopupHandler();

	static gulong FileChooserWidget_signal_connect_location_toggle_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "location-toggle-popup", FileChooserWidget_locationTogglePopupHandler, data);
	}

*/
/*

	void FileChooserWidget_placesShortcutHandler();

	static gulong FileChooserWidget_signal_connect_places_shortcut(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "places-shortcut", FileChooserWidget_placesShortcutHandler, data);
	}

*/
/*

	void FileChooserWidget_quickBookmarkHandler();

	static gulong FileChooserWidget_signal_connect_quick_bookmark(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "quick-bookmark", FileChooserWidget_quickBookmarkHandler, data);
	}

*/
/*

	void FileChooserWidget_recentShortcutHandler();

	static gulong FileChooserWidget_signal_connect_recent_shortcut(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "recent-shortcut", FileChooserWidget_recentShortcutHandler, data);
	}

*/
/*

	void FileChooserWidget_searchShortcutHandler();

	static gulong FileChooserWidget_signal_connect_search_shortcut(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search-shortcut", FileChooserWidget_searchShortcutHandler, data);
	}

*/
/*

	void FileChooserWidget_showHiddenHandler();

	static gulong FileChooserWidget_signal_connect_show_hidden(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-hidden", FileChooserWidget_showHiddenHandler, data);
	}

*/
/*

	void FileChooserWidget_upFolderHandler();

	static gulong FileChooserWidget_signal_connect_up_folder(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "up-folder", FileChooserWidget_upFolderHandler, data);
	}

*/
/*

	void FlowBox_activateCursorChildHandler();

	static gulong FlowBox_signal_connect_activate_cursor_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-cursor-child", FlowBox_activateCursorChildHandler, data);
	}

*/
/*

	void FlowBox_childActivatedHandler();

	static gulong FlowBox_signal_connect_child_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "child-activated", FlowBox_childActivatedHandler, data);
	}

*/
/*

	void FlowBox_moveCursorHandler();

	static gulong FlowBox_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", FlowBox_moveCursorHandler, data);
	}

*/
/*

	void FlowBox_selectAllHandler();

	static gulong FlowBox_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", FlowBox_selectAllHandler, data);
	}

*/
/*

	void FlowBox_selectedChildrenChangedHandler();

	static gulong FlowBox_signal_connect_selected_children_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selected-children-changed", FlowBox_selectedChildrenChangedHandler, data);
	}

*/
/*

	void FlowBox_toggleCursorChildHandler();

	static gulong FlowBox_signal_connect_toggle_cursor_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-cursor-child", FlowBox_toggleCursorChildHandler, data);
	}

*/
/*

	void FlowBox_unselectAllHandler();

	static gulong FlowBox_signal_connect_unselect_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unselect-all", FlowBox_unselectAllHandler, data);
	}

*/
/*

	void FlowBoxChild_activateHandler();

	static gulong FlowBoxChild_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", FlowBoxChild_activateHandler, data);
	}

*/
/*

	void HSV_changedHandler();

	static gulong HSV_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", HSV_changedHandler, data);
	}

*/
/*

	void HSV_moveHandler();

	static gulong HSV_signal_connect_move(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move", HSV_moveHandler, data);
	}

*/
/*

	void HandleBox_childAttachedHandler();

	static gulong HandleBox_signal_connect_child_attached(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "child-attached", HandleBox_childAttachedHandler, data);
	}

*/
/*

	void HandleBox_childDetachedHandler();

	static gulong HandleBox_signal_connect_child_detached(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "child-detached", HandleBox_childDetachedHandler, data);
	}

*/
/*

	void IMContext_commitHandler();

	static gulong IMContext_signal_connect_commit(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "commit", IMContext_commitHandler, data);
	}

*/
/*

	void IMContext_deleteSurroundingHandler();

	static gulong IMContext_signal_connect_delete_surrounding(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "delete-surrounding", IMContext_deleteSurroundingHandler, data);
	}

*/
/*

	void IMContext_preeditChangedHandler();

	static gulong IMContext_signal_connect_preedit_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preedit-changed", IMContext_preeditChangedHandler, data);
	}

*/
/*

	void IMContext_preeditEndHandler();

	static gulong IMContext_signal_connect_preedit_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preedit-end", IMContext_preeditEndHandler, data);
	}

*/
/*

	void IMContext_preeditStartHandler();

	static gulong IMContext_signal_connect_preedit_start(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "preedit-start", IMContext_preeditStartHandler, data);
	}

*/
/*

	void IMContext_retrieveSurroundingHandler();

	static gulong IMContext_signal_connect_retrieve_surrounding(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "retrieve-surrounding", IMContext_retrieveSurroundingHandler, data);
	}

*/
/*

	void IconTheme_changedHandler();

	static gulong IconTheme_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", IconTheme_changedHandler, data);
	}

*/
/*

	void IconView_activateCursorItemHandler();

	static gulong IconView_signal_connect_activate_cursor_item(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-cursor-item", IconView_activateCursorItemHandler, data);
	}

*/
/*

	void IconView_itemActivatedHandler();

	static gulong IconView_signal_connect_item_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "item-activated", IconView_itemActivatedHandler, data);
	}

*/
/*

	void IconView_moveCursorHandler();

	static gulong IconView_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", IconView_moveCursorHandler, data);
	}

*/
/*

	void IconView_selectAllHandler();

	static gulong IconView_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", IconView_selectAllHandler, data);
	}

*/
/*

	void IconView_selectCursorItemHandler();

	static gulong IconView_signal_connect_select_cursor_item(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-cursor-item", IconView_selectCursorItemHandler, data);
	}

*/
/*

	void IconView_selectionChangedHandler();

	static gulong IconView_signal_connect_selection_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-changed", IconView_selectionChangedHandler, data);
	}

*/
/*

	void IconView_toggleCursorItemHandler();

	static gulong IconView_signal_connect_toggle_cursor_item(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-cursor-item", IconView_toggleCursorItemHandler, data);
	}

*/
/*

	void IconView_unselectAllHandler();

	static gulong IconView_signal_connect_unselect_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unselect-all", IconView_unselectAllHandler, data);
	}

*/
/*

	void Label_copyClipboardHandler();

	static gulong Label_signal_connect_copy_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "copy-clipboard", Label_copyClipboardHandler, data);
	}

*/
/*

	void Label_moveCursorHandler();

	static gulong Label_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", Label_moveCursorHandler, data);
	}

*/
/*

	void Label_populatePopupHandler();

	static gulong Label_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", Label_populatePopupHandler, data);
	}

*/
/*

	void LinkButton_activateLinkHandler();

	static gulong LinkButton_signal_connect_activate_link(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-link", LinkButton_activateLinkHandler, data);
	}

*/
/*

	void ListBox_activateCursorRowHandler();

	static gulong ListBox_signal_connect_activate_cursor_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-cursor-row", ListBox_activateCursorRowHandler, data);
	}

*/
/*

	void ListBox_moveCursorHandler();

	static gulong ListBox_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", ListBox_moveCursorHandler, data);
	}

*/
/*

	void ListBox_toggleCursorRowHandler();

	static gulong ListBox_signal_connect_toggle_cursor_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-cursor-row", ListBox_toggleCursorRowHandler, data);
	}

*/
/*

	void Menu_moveScrollHandler();

	static gulong Menu_signal_connect_move_scroll(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-scroll", Menu_moveScrollHandler, data);
	}

*/
/*

	void MenuItem_activateHandler();

	static gulong MenuItem_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", MenuItem_activateHandler, data);
	}

*/
/*

	void MenuItem_activateItemHandler();

	static gulong MenuItem_signal_connect_activate_item(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-item", MenuItem_activateItemHandler, data);
	}

*/
/*

	void MenuItem_deselectHandler();

	static gulong MenuItem_signal_connect_deselect(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "deselect", MenuItem_deselectHandler, data);
	}

*/
/*

	void MenuItem_selectHandler();

	static gulong MenuItem_signal_connect_select(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select", MenuItem_selectHandler, data);
	}

*/
/*

	void MenuItem_toggleSizeAllocateHandler();

	static gulong MenuItem_signal_connect_toggle_size_allocate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-size-allocate", MenuItem_toggleSizeAllocateHandler, data);
	}

*/
/*

	void MenuItem_toggleSizeRequestHandler();

	static gulong MenuItem_signal_connect_toggle_size_request(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-size-request", MenuItem_toggleSizeRequestHandler, data);
	}

*/
/*

	void MenuShell_activateCurrentHandler();

	static gulong MenuShell_signal_connect_activate_current(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-current", MenuShell_activateCurrentHandler, data);
	}

*/
/*

	void MenuShell_cancelHandler();

	static gulong MenuShell_signal_connect_cancel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cancel", MenuShell_cancelHandler, data);
	}

*/
/*

	void MenuShell_cycleFocusHandler();

	static gulong MenuShell_signal_connect_cycle_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cycle-focus", MenuShell_cycleFocusHandler, data);
	}

*/
/*

	void MenuShell_deactivateHandler();

	static gulong MenuShell_signal_connect_deactivate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "deactivate", MenuShell_deactivateHandler, data);
	}

*/
/*

	void MenuShell_moveCurrentHandler();

	static gulong MenuShell_signal_connect_move_current(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-current", MenuShell_moveCurrentHandler, data);
	}

*/
/*

	void MenuShell_selectionDoneHandler();

	static gulong MenuShell_signal_connect_selection_done(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-done", MenuShell_selectionDoneHandler, data);
	}

*/
/*

	void MenuToolButton_showMenuHandler();

	static gulong MenuToolButton_signal_connect_show_menu(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-menu", MenuToolButton_showMenuHandler, data);
	}

*/
/*

	void Notebook_changeCurrentPageHandler();

	static gulong Notebook_signal_connect_change_current_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-current-page", Notebook_changeCurrentPageHandler, data);
	}

*/
/*

	void Notebook_focusTabHandler();

	static gulong Notebook_signal_connect_focus_tab(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-tab", Notebook_focusTabHandler, data);
	}

*/
/*

	void Notebook_moveFocusOutHandler();

	static gulong Notebook_signal_connect_move_focus_out(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-focus-out", Notebook_moveFocusOutHandler, data);
	}

*/
/*

	void Notebook_reorderTabHandler();

	static gulong Notebook_signal_connect_reorder_tab(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "reorder-tab", Notebook_reorderTabHandler, data);
	}

*/
/*

	void Notebook_selectPageHandler();

	static gulong Notebook_signal_connect_select_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-page", Notebook_selectPageHandler, data);
	}

*/
/*

	void Notebook_switchPageHandler();

	static gulong Notebook_signal_connect_switch_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "switch-page", Notebook_switchPageHandler, data);
	}

*/
/*

	void PlacesSidebar_showConnectToServerHandler();

	static gulong PlacesSidebar_signal_connect_show_connect_to_server(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-connect-to-server", PlacesSidebar_showConnectToServerHandler, data);
	}

*/
/*

	void Popover_closedHandler();

	static gulong Popover_signal_connect_closed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "closed", Popover_closedHandler, data);
	}

*/
/*

	void RadioMenuItem_groupChangedHandler();

	static gulong RadioMenuItem_signal_connect_group_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "group-changed", RadioMenuItem_groupChangedHandler, data);
	}

*/
/*

	void Range_adjustBoundsHandler();

	static gulong Range_signal_connect_adjust_bounds(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "adjust-bounds", Range_adjustBoundsHandler, data);
	}

*/
/*

	void Range_moveSliderHandler();

	static gulong Range_signal_connect_move_slider(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-slider", Range_moveSliderHandler, data);
	}

*/
/*

	void Range_valueChangedHandler();

	static gulong Range_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", Range_valueChangedHandler, data);
	}

*/
/*

	void Scale_formatValueHandler();

	static gulong Scale_signal_connect_format_value(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "format-value", Scale_formatValueHandler, data);
	}

*/
/*

	void ScrolledWindow_moveFocusOutHandler();

	static gulong ScrolledWindow_signal_connect_move_focus_out(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-focus-out", ScrolledWindow_moveFocusOutHandler, data);
	}

*/
/*

	void ScrolledWindow_scrollChildHandler();

	static gulong ScrolledWindow_signal_connect_scroll_child(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-child", ScrolledWindow_scrollChildHandler, data);
	}

*/
/*

	void ShortcutsSection_changeCurrentPageHandler();

	static gulong ShortcutsSection_signal_connect_change_current_page(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-current-page", ShortcutsSection_changeCurrentPageHandler, data);
	}

*/
/*

	void ShortcutsWindow_closeHandler();

	static gulong ShortcutsWindow_signal_connect_close(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "close", ShortcutsWindow_closeHandler, data);
	}

*/
/*

	void ShortcutsWindow_searchHandler();

	static gulong ShortcutsWindow_signal_connect_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "search", ShortcutsWindow_searchHandler, data);
	}

*/
/*

	void SpinButton_changeValueHandler();

	static gulong SpinButton_signal_connect_change_value(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "change-value", SpinButton_changeValueHandler, data);
	}

*/
/*

	void SpinButton_inputHandler();

	static gulong SpinButton_signal_connect_input(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "input", SpinButton_inputHandler, data);
	}

*/
/*

	void SpinButton_outputHandler();

	static gulong SpinButton_signal_connect_output(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "output", SpinButton_outputHandler, data);
	}

*/
/*

	void SpinButton_valueChangedHandler();

	static gulong SpinButton_signal_connect_value_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "value-changed", SpinButton_valueChangedHandler, data);
	}

*/
/*

	void Statusbar_textPoppedHandler();

	static gulong Statusbar_signal_connect_text_popped(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-popped", Statusbar_textPoppedHandler, data);
	}

*/
/*

	void Statusbar_textPushedHandler();

	static gulong Statusbar_signal_connect_text_pushed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "text-pushed", Statusbar_textPushedHandler, data);
	}

*/
/*

	void Switch_activateHandler();

	static gulong Switch_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", Switch_activateHandler, data);
	}

*/
/*

	void TextBuffer_applyTagHandler();

	static gulong TextBuffer_signal_connect_apply_tag(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "apply-tag", TextBuffer_applyTagHandler, data);
	}

*/
/*

	void TextBuffer_beginUserActionHandler();

	static gulong TextBuffer_signal_connect_begin_user_action(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "begin-user-action", TextBuffer_beginUserActionHandler, data);
	}

*/
/*

	void TextBuffer_changedHandler();

	static gulong TextBuffer_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", TextBuffer_changedHandler, data);
	}

*/
/*

	void TextBuffer_deleteRangeHandler();

	static gulong TextBuffer_signal_connect_delete_range(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "delete-range", TextBuffer_deleteRangeHandler, data);
	}

*/
/*

	void TextBuffer_endUserActionHandler();

	static gulong TextBuffer_signal_connect_end_user_action(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "end-user-action", TextBuffer_endUserActionHandler, data);
	}

*/
/*

	void TextBuffer_insertChildAnchorHandler();

	static gulong TextBuffer_signal_connect_insert_child_anchor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-child-anchor", TextBuffer_insertChildAnchorHandler, data);
	}

*/
/*

	void TextBuffer_insertPixbufHandler();

	static gulong TextBuffer_signal_connect_insert_pixbuf(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-pixbuf", TextBuffer_insertPixbufHandler, data);
	}

*/
/*

	void TextBuffer_insertTextHandler();

	static gulong TextBuffer_signal_connect_insert_text(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-text", TextBuffer_insertTextHandler, data);
	}

*/
/*

	void TextBuffer_markDeletedHandler();

	static gulong TextBuffer_signal_connect_mark_deleted(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mark-deleted", TextBuffer_markDeletedHandler, data);
	}

*/
/*

	void TextBuffer_markSetHandler();

	static gulong TextBuffer_signal_connect_mark_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mark-set", TextBuffer_markSetHandler, data);
	}

*/
/*

	void TextBuffer_modifiedChangedHandler();

	static gulong TextBuffer_signal_connect_modified_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "modified-changed", TextBuffer_modifiedChangedHandler, data);
	}

*/
/*

	void TextBuffer_removeTagHandler();

	static gulong TextBuffer_signal_connect_remove_tag(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "remove-tag", TextBuffer_removeTagHandler, data);
	}

*/
/*

	void TextTagTable_tagAddedHandler();

	static gulong TextTagTable_signal_connect_tag_added(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tag-added", TextTagTable_tagAddedHandler, data);
	}

*/
/*

	void TextTagTable_tagChangedHandler();

	static gulong TextTagTable_signal_connect_tag_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tag-changed", TextTagTable_tagChangedHandler, data);
	}

*/
/*

	void TextTagTable_tagRemovedHandler();

	static gulong TextTagTable_signal_connect_tag_removed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "tag-removed", TextTagTable_tagRemovedHandler, data);
	}

*/
/*

	void TextView_backspaceHandler();

	static gulong TextView_signal_connect_backspace(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "backspace", TextView_backspaceHandler, data);
	}

*/
/*

	void TextView_copyClipboardHandler();

	static gulong TextView_signal_connect_copy_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "copy-clipboard", TextView_copyClipboardHandler, data);
	}

*/
/*

	void TextView_cutClipboardHandler();

	static gulong TextView_signal_connect_cut_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cut-clipboard", TextView_cutClipboardHandler, data);
	}

*/
/*

	void TextView_deleteFromCursorHandler();

	static gulong TextView_signal_connect_delete_from_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "delete-from-cursor", TextView_deleteFromCursorHandler, data);
	}

*/
/*

	void TextView_insertAtCursorHandler();

	static gulong TextView_signal_connect_insert_at_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "insert-at-cursor", TextView_insertAtCursorHandler, data);
	}

*/
/*

	void TextView_moveCursorHandler();

	static gulong TextView_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", TextView_moveCursorHandler, data);
	}

*/
/*

	void TextView_moveViewportHandler();

	static gulong TextView_signal_connect_move_viewport(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-viewport", TextView_moveViewportHandler, data);
	}

*/
/*

	void TextView_pasteClipboardHandler();

	static gulong TextView_signal_connect_paste_clipboard(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "paste-clipboard", TextView_pasteClipboardHandler, data);
	}

*/
/*

	void TextView_populatePopupHandler();

	static gulong TextView_signal_connect_populate_popup(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "populate-popup", TextView_populatePopupHandler, data);
	}

*/
/*

	void TextView_selectAllHandler();

	static gulong TextView_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", TextView_selectAllHandler, data);
	}

*/
/*

	void TextView_setAnchorHandler();

	static gulong TextView_signal_connect_set_anchor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "set-anchor", TextView_setAnchorHandler, data);
	}

*/
/*

	void TextView_toggleCursorVisibleHandler();

	static gulong TextView_signal_connect_toggle_cursor_visible(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-cursor-visible", TextView_toggleCursorVisibleHandler, data);
	}

*/
/*

	void TextView_toggleOverwriteHandler();

	static gulong TextView_signal_connect_toggle_overwrite(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-overwrite", TextView_toggleOverwriteHandler, data);
	}

*/
/*

	void ToggleAction_toggledHandler();

	static gulong ToggleAction_signal_connect_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggled", ToggleAction_toggledHandler, data);
	}

*/
/*

	void ToggleButton_toggledHandler();

	static gulong ToggleButton_signal_connect_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggled", ToggleButton_toggledHandler, data);
	}

*/
/*

	void ToggleToolButton_toggledHandler();

	static gulong ToggleToolButton_signal_connect_toggled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggled", ToggleToolButton_toggledHandler, data);
	}

*/
/*

	void ToolButton_clickedHandler();

	static gulong ToolButton_signal_connect_clicked(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "clicked", ToolButton_clickedHandler, data);
	}

*/
/*

	void ToolItem_createMenuProxyHandler();

	static gulong ToolItem_signal_connect_create_menu_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "create-menu-proxy", ToolItem_createMenuProxyHandler, data);
	}

*/
/*

	void ToolItem_toolbarReconfiguredHandler();

	static gulong ToolItem_signal_connect_toolbar_reconfigured(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toolbar-reconfigured", ToolItem_toolbarReconfiguredHandler, data);
	}

*/
/*

	void Toolbar_focusHomeOrEndHandler();

	static gulong Toolbar_signal_connect_focus_home_or_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-home-or-end", Toolbar_focusHomeOrEndHandler, data);
	}

*/
/*

	void Toolbar_orientationChangedHandler();

	static gulong Toolbar_signal_connect_orientation_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "orientation-changed", Toolbar_orientationChangedHandler, data);
	}

*/
/*

	void Toolbar_popupContextMenuHandler();

	static gulong Toolbar_signal_connect_popup_context_menu(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup-context-menu", Toolbar_popupContextMenuHandler, data);
	}

*/
/*

	void Toolbar_styleChangedHandler();

	static gulong Toolbar_signal_connect_style_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-changed", Toolbar_styleChangedHandler, data);
	}

*/
/*

	void TreeSelection_changedHandler();

	static gulong TreeSelection_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", TreeSelection_changedHandler, data);
	}

*/
/*

	void TreeView_columnsChangedHandler();

	static gulong TreeView_signal_connect_columns_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "columns-changed", TreeView_columnsChangedHandler, data);
	}

*/
/*

	void TreeView_cursorChangedHandler();

	static gulong TreeView_signal_connect_cursor_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "cursor-changed", TreeView_cursorChangedHandler, data);
	}

*/
/*

	void TreeView_expandCollapseCursorRowHandler();

	static gulong TreeView_signal_connect_expand_collapse_cursor_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "expand-collapse-cursor-row", TreeView_expandCollapseCursorRowHandler, data);
	}

*/
/*

	void TreeView_moveCursorHandler();

	static gulong TreeView_signal_connect_move_cursor(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-cursor", TreeView_moveCursorHandler, data);
	}

*/
/*

	void TreeView_rowActivatedHandler();

	static gulong TreeView_signal_connect_row_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-activated", TreeView_rowActivatedHandler, data);
	}

*/
/*

	void TreeView_rowCollapsedHandler();

	static gulong TreeView_signal_connect_row_collapsed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-collapsed", TreeView_rowCollapsedHandler, data);
	}

*/
/*

	void TreeView_rowExpandedHandler();

	static gulong TreeView_signal_connect_row_expanded(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "row-expanded", TreeView_rowExpandedHandler, data);
	}

*/
/*

	void TreeView_selectAllHandler();

	static gulong TreeView_signal_connect_select_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-all", TreeView_selectAllHandler, data);
	}

*/
/*

	void TreeView_selectCursorParentHandler();

	static gulong TreeView_signal_connect_select_cursor_parent(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-cursor-parent", TreeView_selectCursorParentHandler, data);
	}

*/
/*

	void TreeView_selectCursorRowHandler();

	static gulong TreeView_signal_connect_select_cursor_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "select-cursor-row", TreeView_selectCursorRowHandler, data);
	}

*/
/*

	void TreeView_startInteractiveSearchHandler();

	static gulong TreeView_signal_connect_start_interactive_search(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "start-interactive-search", TreeView_startInteractiveSearchHandler, data);
	}

*/
/*

	void TreeView_testCollapseRowHandler();

	static gulong TreeView_signal_connect_test_collapse_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "test-collapse-row", TreeView_testCollapseRowHandler, data);
	}

*/
/*

	void TreeView_testExpandRowHandler();

	static gulong TreeView_signal_connect_test_expand_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "test-expand-row", TreeView_testExpandRowHandler, data);
	}

*/
/*

	void TreeView_toggleCursorRowHandler();

	static gulong TreeView_signal_connect_toggle_cursor_row(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "toggle-cursor-row", TreeView_toggleCursorRowHandler, data);
	}

*/
/*

	void TreeView_unselectAllHandler();

	static gulong TreeView_signal_connect_unselect_all(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unselect-all", TreeView_unselectAllHandler, data);
	}

*/
/*

	void TreeViewColumn_clickedHandler();

	static gulong TreeViewColumn_signal_connect_clicked(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "clicked", TreeViewColumn_clickedHandler, data);
	}

*/
/*

	void Widget_accelClosuresChangedHandler();

	static gulong Widget_signal_connect_accel_closures_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accel-closures-changed", Widget_accelClosuresChangedHandler, data);
	}

*/
/*

	void Widget_buttonPressEventHandler();

	static gulong Widget_signal_connect_button_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-press-event", Widget_buttonPressEventHandler, data);
	}

*/
/*

	void Widget_buttonReleaseEventHandler();

	static gulong Widget_signal_connect_button_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "button-release-event", Widget_buttonReleaseEventHandler, data);
	}

*/
/*

	void Widget_canActivateAccelHandler();

	static gulong Widget_signal_connect_can_activate_accel(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "can-activate-accel", Widget_canActivateAccelHandler, data);
	}

*/
/*

	void Widget_compositedChangedHandler();

	static gulong Widget_signal_connect_composited_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "composited-changed", Widget_compositedChangedHandler, data);
	}

*/
/*

	void Widget_configureEventHandler();

	static gulong Widget_signal_connect_configure_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "configure-event", Widget_configureEventHandler, data);
	}

*/
/*

	void Widget_destroyHandler();

	static gulong Widget_signal_connect_destroy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "destroy", Widget_destroyHandler, data);
	}

*/
/*

	void Widget_directionChangedHandler();

	static gulong Widget_signal_connect_direction_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "direction-changed", Widget_directionChangedHandler, data);
	}

*/
/*

	void Widget_dragBeginHandler();

	static gulong Widget_signal_connect_drag_begin(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-begin", Widget_dragBeginHandler, data);
	}

*/
/*

	void Widget_dragDataDeleteHandler();

	static gulong Widget_signal_connect_drag_data_delete(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-delete", Widget_dragDataDeleteHandler, data);
	}

*/
/*

	void Widget_dragDataGetHandler();

	static gulong Widget_signal_connect_drag_data_get(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-get", Widget_dragDataGetHandler, data);
	}

*/
/*

	void Widget_dragDataReceivedHandler();

	static gulong Widget_signal_connect_drag_data_received(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-data-received", Widget_dragDataReceivedHandler, data);
	}

*/
/*

	void Widget_dragDropHandler();

	static gulong Widget_signal_connect_drag_drop(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-drop", Widget_dragDropHandler, data);
	}

*/
/*

	void Widget_dragEndHandler();

	static gulong Widget_signal_connect_drag_end(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-end", Widget_dragEndHandler, data);
	}

*/
/*

	void Widget_dragLeaveHandler();

	static gulong Widget_signal_connect_drag_leave(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-leave", Widget_dragLeaveHandler, data);
	}

*/
/*

	void Widget_dragMotionHandler();

	static gulong Widget_signal_connect_drag_motion(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drag-motion", Widget_dragMotionHandler, data);
	}

*/
/*

	void Widget_enterNotifyEventHandler();

	static gulong Widget_signal_connect_enter_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enter-notify-event", Widget_enterNotifyEventHandler, data);
	}

*/
/*

	void Widget_focusHandler();

	static gulong Widget_signal_connect_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus", Widget_focusHandler, data);
	}

*/
/*

	void Widget_focusInEventHandler();

	static gulong Widget_signal_connect_focus_in_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-in-event", Widget_focusInEventHandler, data);
	}

*/
/*

	void Widget_focusOutEventHandler();

	static gulong Widget_signal_connect_focus_out_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "focus-out-event", Widget_focusOutEventHandler, data);
	}

*/
/*

	void Widget_grabFocusHandler();

	static gulong Widget_signal_connect_grab_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-focus", Widget_grabFocusHandler, data);
	}

*/
/*

	void Widget_grabNotifyHandler();

	static gulong Widget_signal_connect_grab_notify(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "grab-notify", Widget_grabNotifyHandler, data);
	}

*/
/*

	void Widget_hideHandler();

	static gulong Widget_signal_connect_hide(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "hide", Widget_hideHandler, data);
	}

*/
/*

	void Widget_hierarchyChangedHandler();

	static gulong Widget_signal_connect_hierarchy_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "hierarchy-changed", Widget_hierarchyChangedHandler, data);
	}

*/
/*

	void Widget_keyPressEventHandler();

	static gulong Widget_signal_connect_key_press_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "key-press-event", Widget_keyPressEventHandler, data);
	}

*/
/*

	void Widget_keyReleaseEventHandler();

	static gulong Widget_signal_connect_key_release_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "key-release-event", Widget_keyReleaseEventHandler, data);
	}

*/
/*

	void Widget_leaveNotifyEventHandler();

	static gulong Widget_signal_connect_leave_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "leave-notify-event", Widget_leaveNotifyEventHandler, data);
	}

*/
/*

	void Widget_mapHandler();

	static gulong Widget_signal_connect_map(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "map", Widget_mapHandler, data);
	}

*/
/*

	void Widget_mapEventHandler();

	static gulong Widget_signal_connect_map_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "map-event", Widget_mapEventHandler, data);
	}

*/
/*

	void Widget_mnemonicActivateHandler();

	static gulong Widget_signal_connect_mnemonic_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "mnemonic-activate", Widget_mnemonicActivateHandler, data);
	}

*/
/*

	void Widget_motionNotifyEventHandler();

	static gulong Widget_signal_connect_motion_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "motion-notify-event", Widget_motionNotifyEventHandler, data);
	}

*/
/*

	void Widget_moveFocusHandler();

	static gulong Widget_signal_connect_move_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "move-focus", Widget_moveFocusHandler, data);
	}

*/
/*

	void Widget_parentSetHandler();

	static gulong Widget_signal_connect_parent_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "parent-set", Widget_parentSetHandler, data);
	}

*/
/*

	void Widget_popupMenuHandler();

	static gulong Widget_signal_connect_popup_menu(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "popup-menu", Widget_popupMenuHandler, data);
	}

*/
/*

	void Widget_propertyNotifyEventHandler();

	static gulong Widget_signal_connect_property_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "property-notify-event", Widget_propertyNotifyEventHandler, data);
	}

*/
/*

	void Widget_proximityInEventHandler();

	static gulong Widget_signal_connect_proximity_in_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "proximity-in-event", Widget_proximityInEventHandler, data);
	}

*/
/*

	void Widget_proximityOutEventHandler();

	static gulong Widget_signal_connect_proximity_out_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "proximity-out-event", Widget_proximityOutEventHandler, data);
	}

*/
/*

	void Widget_realizeHandler();

	static gulong Widget_signal_connect_realize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "realize", Widget_realizeHandler, data);
	}

*/
/*

	void Widget_screenChangedHandler();

	static gulong Widget_signal_connect_screen_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "screen-changed", Widget_screenChangedHandler, data);
	}

*/
/*

	void Widget_scrollEventHandler();

	static gulong Widget_signal_connect_scroll_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "scroll-event", Widget_scrollEventHandler, data);
	}

*/
/*

	void Widget_selectionClearEventHandler();

	static gulong Widget_signal_connect_selection_clear_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-clear-event", Widget_selectionClearEventHandler, data);
	}

*/
/*

	void Widget_selectionGetHandler();

	static gulong Widget_signal_connect_selection_get(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-get", Widget_selectionGetHandler, data);
	}

*/
/*

	void Widget_selectionNotifyEventHandler();

	static gulong Widget_signal_connect_selection_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-notify-event", Widget_selectionNotifyEventHandler, data);
	}

*/
/*

	void Widget_selectionReceivedHandler();

	static gulong Widget_signal_connect_selection_received(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-received", Widget_selectionReceivedHandler, data);
	}

*/
/*

	void Widget_selectionRequestEventHandler();

	static gulong Widget_signal_connect_selection_request_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "selection-request-event", Widget_selectionRequestEventHandler, data);
	}

*/
/*

	void Widget_showHandler();

	static gulong Widget_signal_connect_show(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show", Widget_showHandler, data);
	}

*/
/*

	void Widget_showHelpHandler();

	static gulong Widget_signal_connect_show_help(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "show-help", Widget_showHelpHandler, data);
	}

*/
/*

	void Widget_stateChangedHandler();

	static gulong Widget_signal_connect_state_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "state-changed", Widget_stateChangedHandler, data);
	}

*/
/*

	void Widget_styleSetHandler();

	static gulong Widget_signal_connect_style_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "style-set", Widget_styleSetHandler, data);
	}

*/
/*

	void Widget_unmapHandler();

	static gulong Widget_signal_connect_unmap(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmap", Widget_unmapHandler, data);
	}

*/
/*

	void Widget_unmapEventHandler();

	static gulong Widget_signal_connect_unmap_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unmap-event", Widget_unmapEventHandler, data);
	}

*/
/*

	void Widget_unrealizeHandler();

	static gulong Widget_signal_connect_unrealize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unrealize", Widget_unrealizeHandler, data);
	}

*/
/*

	void Widget_visibilityNotifyEventHandler();

	static gulong Widget_signal_connect_visibility_notify_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "visibility-notify-event", Widget_visibilityNotifyEventHandler, data);
	}

*/
/*

	void Widget_windowStateEventHandler();

	static gulong Widget_signal_connect_window_state_event(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "window-state-event", Widget_windowStateEventHandler, data);
	}

*/
/*

	void Window_activateDefaultHandler();

	static gulong Window_signal_connect_activate_default(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-default", Window_activateDefaultHandler, data);
	}

*/
/*

	void Window_activateFocusHandler();

	static gulong Window_signal_connect_activate_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate-focus", Window_activateFocusHandler, data);
	}

*/
/*

	void Window_enableDebuggingHandler();

	static gulong Window_signal_connect_enable_debugging(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "enable-debugging", Window_enableDebuggingHandler, data);
	}

*/
/*

	void Window_keysChangedHandler();

	static gulong Window_signal_connect_keys_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "keys-changed", Window_keysChangedHandler, data);
	}

*/
/*

	void Window_setFocusHandler();

	static gulong Window_signal_connect_set_focus(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "set-focus", Window_setFocusHandler, data);
	}

*/
import "C"

// AboutDialog is a wrapper around the C record GtkAboutDialog.
type AboutDialog struct {
	native *C.GtkAboutDialog
	// parent_instance : record
	// Private : priv
}

func AboutDialogNewFromC(u unsafe.Pointer) *AboutDialog {
	c := (*C.GtkAboutDialog)(u)
	if c == nil {
		return nil
	}

	g := &AboutDialog{native: c}

	return g
}

func (recv *AboutDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *AboutDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *AboutDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *AboutDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *AboutDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *AboutDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AboutDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AboutDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to AboutDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a AboutDialog.
func CastToAboutDialog(object *gobject.Object) *AboutDialog {
	return AboutDialogNewFromC(object.ToC())
}

// AccelGroup is a wrapper around the C record GtkAccelGroup.
type AccelGroup struct {
	native *C.GtkAccelGroup
	// parent : record
	// priv : record
}

func AccelGroupNewFromC(u unsafe.Pointer) *AccelGroup {
	c := (*C.GtkAccelGroup)(u)
	if c == nil {
		return nil
	}

	g := &AccelGroup{native: c}

	return g
}

func (recv *AccelGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *AccelGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to AccelGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a AccelGroup.
func CastToAccelGroup(object *gobject.Object) *AccelGroup {
	return AccelGroupNewFromC(object.ToC())
}

var signalAccelActivateId int
var signalAccelActivateMap = make(map[int]AccelGroupSignalAccelActivateCallback)
var signalAccelActivateLock sync.Mutex

// AccelGroupSignalAccelActivateCallback is a callback function for a 'accel-activate' signal emitted from a AccelGroup.
type AccelGroupSignalAccelActivateCallback func(acceleratable *gobject.Object, keyval uint32, modifier gdk.ModifierType) bool

func AccelGroup_accelActivateHandler() {}

var signalAccelChangedId int
var signalAccelChangedMap = make(map[int]AccelGroupSignalAccelChangedCallback)
var signalAccelChangedLock sync.Mutex

// AccelGroupSignalAccelChangedCallback is a callback function for a 'accel-changed' signal emitted from a AccelGroup.
type AccelGroupSignalAccelChangedCallback func(keyval uint32, modifier gdk.ModifierType, accelClosure *gobject.Closure)

func AccelGroup_accelChangedHandler() {}

// AccelGroupNew is a wrapper around the C function gtk_accel_group_new.
func AccelGroupNew() *AccelGroup {
	retC := C.gtk_accel_group_new()
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function gtk_accel_group_activate.
func (recv *AccelGroup) Activate(accelQuark glib.Quark, acceleratable *gobject.Object, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_quark := (C.GQuark)(accelQuark)

	c_acceleratable := (*C.GObject)(acceleratable.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_group_activate((*C.GtkAccelGroup)(recv.native), c_accel_quark, c_acceleratable, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Connect is a wrapper around the C function gtk_accel_group_connect.
func (recv *AccelGroup) Connect(accelKey uint32, accelMods gdk.ModifierType, accelFlags AccelFlags, closure *gobject.Closure) {
	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	c_accel_flags := (C.GtkAccelFlags)(accelFlags)

	c_closure := (*C.GClosure)(closure.ToC())

	C.gtk_accel_group_connect((*C.GtkAccelGroup)(recv.native), c_accel_key, c_accel_mods, c_accel_flags, c_closure)

	return
}

// ConnectByPath is a wrapper around the C function gtk_accel_group_connect_by_path.
func (recv *AccelGroup) ConnectByPath(accelPath string, closure *gobject.Closure) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	c_closure := (*C.GClosure)(closure.ToC())

	C.gtk_accel_group_connect_by_path((*C.GtkAccelGroup)(recv.native), c_accel_path, c_closure)

	return
}

// Disconnect is a wrapper around the C function gtk_accel_group_disconnect.
func (recv *AccelGroup) Disconnect(closure *gobject.Closure) bool {
	c_closure := (*C.GClosure)(closure.ToC())

	retC := C.gtk_accel_group_disconnect((*C.GtkAccelGroup)(recv.native), c_closure)
	retGo := retC == C.TRUE

	return retGo
}

// DisconnectKey is a wrapper around the C function gtk_accel_group_disconnect_key.
func (recv *AccelGroup) DisconnectKey(accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_accel_group_disconnect_key((*C.GtkAccelGroup)(recv.native), c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_accel_group_find : unsupported parameter find_func : no type generator for AccelGroupFindFunc, GtkAccelGroupFindFunc

// Lock is a wrapper around the C function gtk_accel_group_lock.
func (recv *AccelGroup) Lock() {
	C.gtk_accel_group_lock((*C.GtkAccelGroup)(recv.native))

	return
}

// Unsupported : gtk_accel_group_query : no return type

// Unlock is a wrapper around the C function gtk_accel_group_unlock.
func (recv *AccelGroup) Unlock() {
	C.gtk_accel_group_unlock((*C.GtkAccelGroup)(recv.native))

	return
}

// AccelLabel is a wrapper around the C record GtkAccelLabel.
type AccelLabel struct {
	native *C.GtkAccelLabel
	// label : record
	// priv : record
}

func AccelLabelNewFromC(u unsafe.Pointer) *AccelLabel {
	c := (*C.GtkAccelLabel)(u)
	if c == nil {
		return nil
	}

	g := &AccelLabel{native: c}

	return g
}

func (recv *AccelLabel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Label upcasts to *Label
func (recv *AccelLabel) Label() *Label {
	return LabelNewFromC(unsafe.Pointer(recv.native))
}

// Misc upcasts to *Misc
func (recv *AccelLabel) Misc() *Misc {
	return recv.Label().Misc()
}

// Widget upcasts to *Widget
func (recv *AccelLabel) Widget() *Widget {
	return recv.Label().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AccelLabel) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Label().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AccelLabel) Object() *gobject.Object {
	return recv.Label().Object()
}

// CastToWidget down casts any arbitary Object to AccelLabel.
// Exercise care, as this is a potentially dangerous function if the Object is not a AccelLabel.
func CastToAccelLabel(object *gobject.Object) *AccelLabel {
	return AccelLabelNewFromC(object.ToC())
}

// AccelLabelNew is a wrapper around the C function gtk_accel_label_new.
func AccelLabelNew(string string) *AccelLabel {
	c_string := C.CString(string)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.gtk_accel_label_new(c_string)
	retGo := AccelLabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAccelWidget is a wrapper around the C function gtk_accel_label_get_accel_widget.
func (recv *AccelLabel) GetAccelWidget() *Widget {
	retC := C.gtk_accel_label_get_accel_widget((*C.GtkAccelLabel)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAccelWidth is a wrapper around the C function gtk_accel_label_get_accel_width.
func (recv *AccelLabel) GetAccelWidth() uint32 {
	retC := C.gtk_accel_label_get_accel_width((*C.GtkAccelLabel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// Refetch is a wrapper around the C function gtk_accel_label_refetch.
func (recv *AccelLabel) Refetch() bool {
	retC := C.gtk_accel_label_refetch((*C.GtkAccelLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAccelClosure is a wrapper around the C function gtk_accel_label_set_accel_closure.
func (recv *AccelLabel) SetAccelClosure(accelClosure *gobject.Closure) {
	c_accel_closure := (*C.GClosure)(accelClosure.ToC())

	C.gtk_accel_label_set_accel_closure((*C.GtkAccelLabel)(recv.native), c_accel_closure)

	return
}

// SetAccelWidget is a wrapper around the C function gtk_accel_label_set_accel_widget.
func (recv *AccelLabel) SetAccelWidget(accelWidget *Widget) {
	c_accel_widget := (*C.GtkWidget)(accelWidget.ToC())

	C.gtk_accel_label_set_accel_widget((*C.GtkAccelLabel)(recv.native), c_accel_widget)

	return
}

// AccelMap is a wrapper around the C record GtkAccelMap.
type AccelMap struct {
	native *C.GtkAccelMap
}

func AccelMapNewFromC(u unsafe.Pointer) *AccelMap {
	c := (*C.GtkAccelMap)(u)
	if c == nil {
		return nil
	}

	g := &AccelMap{native: c}

	return g
}

func (recv *AccelMap) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *AccelMap) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to AccelMap.
// Exercise care, as this is a potentially dangerous function if the Object is not a AccelMap.
func CastToAccelMap(object *gobject.Object) *AccelMap {
	return AccelMapNewFromC(object.ToC())
}

// Accessible is a wrapper around the C record GtkAccessible.
type Accessible struct {
	native *C.GtkAccessible
	// parent : record
	// Private : priv
}

func AccessibleNewFromC(u unsafe.Pointer) *Accessible {
	c := (*C.GtkAccessible)(u)
	if c == nil {
		return nil
	}

	g := &Accessible{native: c}

	return g
}

func (recv *Accessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Accessible) Object() *atk.Object {
	return atk.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Accessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a Accessible.
func CastToAccessible(object *gobject.Object) *Accessible {
	return AccessibleNewFromC(object.ToC())
}

// ConnectWidgetDestroyed is a wrapper around the C function gtk_accessible_connect_widget_destroyed.
func (recv *Accessible) ConnectWidgetDestroyed() {
	C.gtk_accessible_connect_widget_destroyed((*C.GtkAccessible)(recv.native))

	return
}

// Action is a wrapper around the C record GtkAction.
type Action struct {
	native *C.GtkAction
	// object : record
	// Private : private_data
}

func ActionNewFromC(u unsafe.Pointer) *Action {
	c := (*C.GtkAction)(u)
	if c == nil {
		return nil
	}

	g := &Action{native: c}

	return g
}

func (recv *Action) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Action) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Action.
// Exercise care, as this is a potentially dangerous function if the Object is not a Action.
func CastToAction(object *gobject.Object) *Action {
	return ActionNewFromC(object.ToC())
}

// ActionBar is a wrapper around the C record GtkActionBar.
type ActionBar struct {
	native *C.GtkActionBar
	// Private : bin
}

func ActionBarNewFromC(u unsafe.Pointer) *ActionBar {
	c := (*C.GtkActionBar)(u)
	if c == nil {
		return nil
	}

	g := &ActionBar{native: c}

	return g
}

func (recv *ActionBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *ActionBar) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ActionBar) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *ActionBar) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ActionBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ActionBar) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to ActionBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a ActionBar.
func CastToActionBar(object *gobject.Object) *ActionBar {
	return ActionBarNewFromC(object.ToC())
}

// ActionGroup is a wrapper around the C record GtkActionGroup.
type ActionGroup struct {
	native *C.GtkActionGroup
	// parent : record
	// Private : priv
}

func ActionGroupNewFromC(u unsafe.Pointer) *ActionGroup {
	c := (*C.GtkActionGroup)(u)
	if c == nil {
		return nil
	}

	g := &ActionGroup{native: c}

	return g
}

func (recv *ActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ActionGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ActionGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a ActionGroup.
func CastToActionGroup(object *gobject.Object) *ActionGroup {
	return ActionGroupNewFromC(object.ToC())
}

// Adjustment is a wrapper around the C record GtkAdjustment.
type Adjustment struct {
	native *C.GtkAdjustment
	// parent_instance : record
	// priv : record
}

func AdjustmentNewFromC(u unsafe.Pointer) *Adjustment {
	c := (*C.GtkAdjustment)(u)
	if c == nil {
		return nil
	}

	g := &Adjustment{native: c}

	return g
}

func (recv *Adjustment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Adjustment) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Adjustment) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to Adjustment.
// Exercise care, as this is a potentially dangerous function if the Object is not a Adjustment.
func CastToAdjustment(object *gobject.Object) *Adjustment {
	return AdjustmentNewFromC(object.ToC())
}

var signalChangedId int
var signalChangedMap = make(map[int]AdjustmentSignalChangedCallback)
var signalChangedLock sync.Mutex

// AdjustmentSignalChangedCallback is a callback function for a 'changed' signal emitted from a Adjustment.
type AdjustmentSignalChangedCallback func()

func Adjustment_changedHandler() {}

var signalValueChangedId int
var signalValueChangedMap = make(map[int]AdjustmentSignalValueChangedCallback)
var signalValueChangedLock sync.Mutex

// AdjustmentSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a Adjustment.
type AdjustmentSignalValueChangedCallback func()

func Adjustment_valueChangedHandler() {}

// AdjustmentNew is a wrapper around the C function gtk_adjustment_new.
func AdjustmentNew(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) *Adjustment {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_step_increment := (C.gdouble)(stepIncrement)

	c_page_increment := (C.gdouble)(pageIncrement)

	c_page_size := (C.gdouble)(pageSize)

	retC := C.gtk_adjustment_new(c_value, c_lower, c_upper, c_step_increment, c_page_increment, c_page_size)
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Changed is a wrapper around the C function gtk_adjustment_changed.
func (recv *Adjustment) Changed() {
	C.gtk_adjustment_changed((*C.GtkAdjustment)(recv.native))

	return
}

// ClampPage is a wrapper around the C function gtk_adjustment_clamp_page.
func (recv *Adjustment) ClampPage(lower float64, upper float64) {
	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	C.gtk_adjustment_clamp_page((*C.GtkAdjustment)(recv.native), c_lower, c_upper)

	return
}

// GetValue is a wrapper around the C function gtk_adjustment_get_value.
func (recv *Adjustment) GetValue() float64 {
	retC := C.gtk_adjustment_get_value((*C.GtkAdjustment)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetValue is a wrapper around the C function gtk_adjustment_set_value.
func (recv *Adjustment) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_adjustment_set_value((*C.GtkAdjustment)(recv.native), c_value)

	return
}

// ValueChanged is a wrapper around the C function gtk_adjustment_value_changed.
func (recv *Adjustment) ValueChanged() {
	C.gtk_adjustment_value_changed((*C.GtkAdjustment)(recv.native))

	return
}

// Alignment is a wrapper around the C record GtkAlignment.
type Alignment struct {
	native *C.GtkAlignment
	// bin : record
	// Private : priv
}

func AlignmentNewFromC(u unsafe.Pointer) *Alignment {
	c := (*C.GtkAlignment)(u)
	if c == nil {
		return nil
	}

	g := &Alignment{native: c}

	return g
}

func (recv *Alignment) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Alignment) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Alignment) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Alignment) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Alignment) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Alignment) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Alignment.
// Exercise care, as this is a potentially dangerous function if the Object is not a Alignment.
func CastToAlignment(object *gobject.Object) *Alignment {
	return AlignmentNewFromC(object.ToC())
}

// AlignmentNew is a wrapper around the C function gtk_alignment_new.
func AlignmentNew(xalign float32, yalign float32, xscale float32, yscale float32) *Alignment {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	retC := C.gtk_alignment_new(c_xalign, c_yalign, c_xscale, c_yscale)
	retGo := AlignmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function gtk_alignment_set.
func (recv *Alignment) Set(xalign float32, yalign float32, xscale float32, yscale float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	C.gtk_alignment_set((*C.GtkAlignment)(recv.native), c_xalign, c_yalign, c_xscale, c_yscale)

	return
}

// AppChooserButton is a wrapper around the C record GtkAppChooserButton.
type AppChooserButton struct {
	native *C.GtkAppChooserButton
	// parent : record
	// Private : priv
}

func AppChooserButtonNewFromC(u unsafe.Pointer) *AppChooserButton {
	c := (*C.GtkAppChooserButton)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserButton{native: c}

	return g
}

func (recv *AppChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBox upcasts to *ComboBox
func (recv *AppChooserButton) ComboBox() *ComboBox {
	return ComboBoxNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *AppChooserButton) Bin() *Bin {
	return recv.ComboBox().Bin()
}

// Container upcasts to *Container
func (recv *AppChooserButton) Container() *Container {
	return recv.ComboBox().Container()
}

// Widget upcasts to *Widget
func (recv *AppChooserButton) Widget() *Widget {
	return recv.ComboBox().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AppChooserButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ComboBox().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AppChooserButton) Object() *gobject.Object {
	return recv.ComboBox().Object()
}

// CastToWidget down casts any arbitary Object to AppChooserButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppChooserButton.
func CastToAppChooserButton(object *gobject.Object) *AppChooserButton {
	return AppChooserButtonNewFromC(object.ToC())
}

var signalCustomItemActivatedId int
var signalCustomItemActivatedMap = make(map[int]AppChooserButtonSignalCustomItemActivatedCallback)
var signalCustomItemActivatedLock sync.Mutex

// AppChooserButtonSignalCustomItemActivatedCallback is a callback function for a 'custom-item-activated' signal emitted from a AppChooserButton.
type AppChooserButtonSignalCustomItemActivatedCallback func(itemName string)

func AppChooserButton_customItemActivatedHandler() {}

// GetHeading is a wrapper around the C function gtk_app_chooser_button_get_heading.
func (recv *AppChooserButton) GetHeading() string {
	retC := C.gtk_app_chooser_button_get_heading((*C.GtkAppChooserButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetHeading is a wrapper around the C function gtk_app_chooser_button_set_heading.
func (recv *AppChooserButton) SetHeading(heading string) {
	c_heading := C.CString(heading)
	defer C.free(unsafe.Pointer(c_heading))

	C.gtk_app_chooser_button_set_heading((*C.GtkAppChooserButton)(recv.native), c_heading)

	return
}

// AppChooserDialog is a wrapper around the C record GtkAppChooserDialog.
type AppChooserDialog struct {
	native *C.GtkAppChooserDialog
	// parent : record
	// Private : priv
}

func AppChooserDialogNewFromC(u unsafe.Pointer) *AppChooserDialog {
	c := (*C.GtkAppChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserDialog{native: c}

	return g
}

func (recv *AppChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *AppChooserDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *AppChooserDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *AppChooserDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *AppChooserDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *AppChooserDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AppChooserDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AppChooserDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to AppChooserDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppChooserDialog.
func CastToAppChooserDialog(object *gobject.Object) *AppChooserDialog {
	return AppChooserDialogNewFromC(object.ToC())
}

// Unsupported : gtk_app_chooser_dialog_new : unsupported parameter file : no type generator for Gio.File, GFile*

// GetHeading is a wrapper around the C function gtk_app_chooser_dialog_get_heading.
func (recv *AppChooserDialog) GetHeading() string {
	retC := C.gtk_app_chooser_dialog_get_heading((*C.GtkAppChooserDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetHeading is a wrapper around the C function gtk_app_chooser_dialog_set_heading.
func (recv *AppChooserDialog) SetHeading(heading string) {
	c_heading := C.CString(heading)
	defer C.free(unsafe.Pointer(c_heading))

	C.gtk_app_chooser_dialog_set_heading((*C.GtkAppChooserDialog)(recv.native), c_heading)

	return
}

// AppChooserWidget is a wrapper around the C record GtkAppChooserWidget.
type AppChooserWidget struct {
	native *C.GtkAppChooserWidget
	// parent : record
	// Private : priv
}

func AppChooserWidgetNewFromC(u unsafe.Pointer) *AppChooserWidget {
	c := (*C.GtkAppChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &AppChooserWidget{native: c}

	return g
}

func (recv *AppChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *AppChooserWidget) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *AppChooserWidget) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *AppChooserWidget) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AppChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AppChooserWidget) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to AppChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a AppChooserWidget.
func CastToAppChooserWidget(object *gobject.Object) *AppChooserWidget {
	return AppChooserWidgetNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// Unsupported signal : unsupported parameter application : no type generator for Gio.AppInfo,

// SetDefaultText is a wrapper around the C function gtk_app_chooser_widget_set_default_text.
func (recv *AppChooserWidget) SetDefaultText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_app_chooser_widget_set_default_text((*C.GtkAppChooserWidget)(recv.native), c_text)

	return
}

// Application is a wrapper around the C record GtkApplication.
type Application struct {
	native *C.GtkApplication
	// parent : record
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	c := (*C.GtkApplication)(u)
	if c == nil {
		return nil
	}

	g := &Application{native: c}

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Application upcasts to *Application
func (recv *Application) Application() *gio.Application {
	return gio.ApplicationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Application) Object() *gobject.Object {
	return recv.Application().Object()
}

// CastToWidget down casts any arbitary Object to Application.
// Exercise care, as this is a potentially dangerous function if the Object is not a Application.
func CastToApplication(object *gobject.Object) *Application {
	return ApplicationNewFromC(object.ToC())
}

// ApplicationWindow is a wrapper around the C record GtkApplicationWindow.
type ApplicationWindow struct {
	native *C.GtkApplicationWindow
	// parent_instance : record
	// Private : priv
}

func ApplicationWindowNewFromC(u unsafe.Pointer) *ApplicationWindow {
	c := (*C.GtkApplicationWindow)(u)
	if c == nil {
		return nil
	}

	g := &ApplicationWindow{native: c}

	return g
}

func (recv *ApplicationWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *ApplicationWindow) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ApplicationWindow) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *ApplicationWindow) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *ApplicationWindow) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ApplicationWindow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ApplicationWindow) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to ApplicationWindow.
// Exercise care, as this is a potentially dangerous function if the Object is not a ApplicationWindow.
func CastToApplicationWindow(object *gobject.Object) *ApplicationWindow {
	return ApplicationWindowNewFromC(object.ToC())
}

// Arrow is a wrapper around the C record GtkArrow.
type Arrow struct {
	native *C.GtkArrow
	// misc : record
	// Private : priv
}

func ArrowNewFromC(u unsafe.Pointer) *Arrow {
	c := (*C.GtkArrow)(u)
	if c == nil {
		return nil
	}

	g := &Arrow{native: c}

	return g
}

func (recv *Arrow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Misc upcasts to *Misc
func (recv *Arrow) Misc() *Misc {
	return MiscNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Arrow) Widget() *Widget {
	return recv.Misc().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Arrow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Misc().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Arrow) Object() *gobject.Object {
	return recv.Misc().Object()
}

// CastToWidget down casts any arbitary Object to Arrow.
// Exercise care, as this is a potentially dangerous function if the Object is not a Arrow.
func CastToArrow(object *gobject.Object) *Arrow {
	return ArrowNewFromC(object.ToC())
}

// ArrowNew is a wrapper around the C function gtk_arrow_new.
func ArrowNew(arrowType ArrowType, shadowType ShadowType) *Arrow {
	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	retC := C.gtk_arrow_new(c_arrow_type, c_shadow_type)
	retGo := ArrowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function gtk_arrow_set.
func (recv *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	c_arrow_type := (C.GtkArrowType)(arrowType)

	c_shadow_type := (C.GtkShadowType)(shadowType)

	C.gtk_arrow_set((*C.GtkArrow)(recv.native), c_arrow_type, c_shadow_type)

	return
}

// ArrowAccessible is a wrapper around the C record GtkArrowAccessible.
type ArrowAccessible struct {
	native *C.GtkArrowAccessible
	// parent : record
	// priv : record
}

func ArrowAccessibleNewFromC(u unsafe.Pointer) *ArrowAccessible {
	c := (*C.GtkArrowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ArrowAccessible{native: c}

	return g
}

func (recv *ArrowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ArrowAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *ArrowAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ArrowAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ArrowAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ArrowAccessible.
func CastToArrowAccessible(object *gobject.Object) *ArrowAccessible {
	return ArrowAccessibleNewFromC(object.ToC())
}

// AspectFrame is a wrapper around the C record GtkAspectFrame.
type AspectFrame struct {
	native *C.GtkAspectFrame
	// frame : record
	// Private : priv
}

func AspectFrameNewFromC(u unsafe.Pointer) *AspectFrame {
	c := (*C.GtkAspectFrame)(u)
	if c == nil {
		return nil
	}

	g := &AspectFrame{native: c}

	return g
}

func (recv *AspectFrame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Frame upcasts to *Frame
func (recv *AspectFrame) Frame() *Frame {
	return FrameNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *AspectFrame) Bin() *Bin {
	return recv.Frame().Bin()
}

// Container upcasts to *Container
func (recv *AspectFrame) Container() *Container {
	return recv.Frame().Container()
}

// Widget upcasts to *Widget
func (recv *AspectFrame) Widget() *Widget {
	return recv.Frame().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *AspectFrame) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Frame().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *AspectFrame) Object() *gobject.Object {
	return recv.Frame().Object()
}

// CastToWidget down casts any arbitary Object to AspectFrame.
// Exercise care, as this is a potentially dangerous function if the Object is not a AspectFrame.
func CastToAspectFrame(object *gobject.Object) *AspectFrame {
	return AspectFrameNewFromC(object.ToC())
}

// AspectFrameNew is a wrapper around the C function gtk_aspect_frame_new.
func AspectFrameNew(label string, xalign float32, yalign float32, ratio float32, obeyChild bool) *AspectFrame {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obey_child :=
		boolToGboolean(obeyChild)

	retC := C.gtk_aspect_frame_new(c_label, c_xalign, c_yalign, c_ratio, c_obey_child)
	retGo := AspectFrameNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Set is a wrapper around the C function gtk_aspect_frame_set.
func (recv *AspectFrame) Set(xalign float32, yalign float32, ratio float32, obeyChild bool) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obey_child :=
		boolToGboolean(obeyChild)

	C.gtk_aspect_frame_set((*C.GtkAspectFrame)(recv.native), c_xalign, c_yalign, c_ratio, c_obey_child)

	return
}

// Assistant is a wrapper around the C record GtkAssistant.
type Assistant struct {
	native *C.GtkAssistant
	// parent : record
	// Private : priv
}

func AssistantNewFromC(u unsafe.Pointer) *Assistant {
	c := (*C.GtkAssistant)(u)
	if c == nil {
		return nil
	}

	g := &Assistant{native: c}

	return g
}

func (recv *Assistant) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *Assistant) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *Assistant) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *Assistant) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *Assistant) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Assistant) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Assistant) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to Assistant.
// Exercise care, as this is a potentially dangerous function if the Object is not a Assistant.
func CastToAssistant(object *gobject.Object) *Assistant {
	return AssistantNewFromC(object.ToC())
}

var signalEscapeId int
var signalEscapeMap = make(map[int]AssistantSignalEscapeCallback)
var signalEscapeLock sync.Mutex

// AssistantSignalEscapeCallback is a callback function for a 'escape' signal emitted from a Assistant.
type AssistantSignalEscapeCallback func()

func Assistant_escapeHandler() {}

// Bin is a wrapper around the C record GtkBin.
type Bin struct {
	native *C.GtkBin
	// container : record
	// Private : priv
}

func BinNewFromC(u unsafe.Pointer) *Bin {
	c := (*C.GtkBin)(u)
	if c == nil {
		return nil
	}

	g := &Bin{native: c}

	return g
}

func (recv *Bin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Bin) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Bin) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Bin) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Bin) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Bin.
// Exercise care, as this is a potentially dangerous function if the Object is not a Bin.
func CastToBin(object *gobject.Object) *Bin {
	return BinNewFromC(object.ToC())
}

// GetChild is a wrapper around the C function gtk_bin_get_child.
func (recv *Bin) GetChild() *Widget {
	retC := C.gtk_bin_get_child((*C.GtkBin)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BooleanCellAccessible is a wrapper around the C record GtkBooleanCellAccessible.
type BooleanCellAccessible struct {
	native *C.GtkBooleanCellAccessible
	// parent : record
	// priv : record
}

func BooleanCellAccessibleNewFromC(u unsafe.Pointer) *BooleanCellAccessible {
	c := (*C.GtkBooleanCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &BooleanCellAccessible{native: c}

	return g
}

func (recv *BooleanCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessible upcasts to *RendererCellAccessible
func (recv *BooleanCellAccessible) RendererCellAccessible() *RendererCellAccessible {
	return RendererCellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// CellAccessible upcasts to *CellAccessible
func (recv *BooleanCellAccessible) CellAccessible() *CellAccessible {
	return recv.RendererCellAccessible().CellAccessible()
}

// Accessible upcasts to *Accessible
func (recv *BooleanCellAccessible) Accessible() *Accessible {
	return recv.RendererCellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *BooleanCellAccessible) Object() *atk.Object {
	return recv.RendererCellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to BooleanCellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a BooleanCellAccessible.
func CastToBooleanCellAccessible(object *gobject.Object) *BooleanCellAccessible {
	return BooleanCellAccessibleNewFromC(object.ToC())
}

// Box is a wrapper around the C record GtkBox.
type Box struct {
	native *C.GtkBox
	// container : record
	// Private : priv
}

func BoxNewFromC(u unsafe.Pointer) *Box {
	c := (*C.GtkBox)(u)
	if c == nil {
		return nil
	}

	g := &Box{native: c}

	return g
}

func (recv *Box) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Box) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Box) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Box) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Box) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Box.
// Exercise care, as this is a potentially dangerous function if the Object is not a Box.
func CastToBox(object *gobject.Object) *Box {
	return BoxNewFromC(object.ToC())
}

// GetHomogeneous is a wrapper around the C function gtk_box_get_homogeneous.
func (recv *Box) GetHomogeneous() bool {
	retC := C.gtk_box_get_homogeneous((*C.GtkBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_box_get_spacing.
func (recv *Box) GetSpacing() int32 {
	retC := C.gtk_box_get_spacing((*C.GtkBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_box_pack_end.
func (recv *Box) PackEnd(child *Widget, expand bool, fill bool, padding uint32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_expand :=
		boolToGboolean(expand)

	c_fill :=
		boolToGboolean(fill)

	c_padding := (C.guint)(padding)

	C.gtk_box_pack_end((*C.GtkBox)(recv.native), c_child, c_expand, c_fill, c_padding)

	return
}

// PackStart is a wrapper around the C function gtk_box_pack_start.
func (recv *Box) PackStart(child *Widget, expand bool, fill bool, padding uint32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_expand :=
		boolToGboolean(expand)

	c_fill :=
		boolToGboolean(fill)

	c_padding := (C.guint)(padding)

	C.gtk_box_pack_start((*C.GtkBox)(recv.native), c_child, c_expand, c_fill, c_padding)

	return
}

// Unsupported : gtk_box_query_child_packing : unsupported parameter pack_type : GtkPackType* with indirection level of 1

// ReorderChild is a wrapper around the C function gtk_box_reorder_child.
func (recv *Box) ReorderChild(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_box_reorder_child((*C.GtkBox)(recv.native), c_child, c_position)

	return
}

// SetChildPacking is a wrapper around the C function gtk_box_set_child_packing.
func (recv *Box) SetChildPacking(child *Widget, expand bool, fill bool, padding uint32, packType PackType) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_expand :=
		boolToGboolean(expand)

	c_fill :=
		boolToGboolean(fill)

	c_padding := (C.guint)(padding)

	c_pack_type := (C.GtkPackType)(packType)

	C.gtk_box_set_child_packing((*C.GtkBox)(recv.native), c_child, c_expand, c_fill, c_padding, c_pack_type)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_box_set_homogeneous.
func (recv *Box) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_box_set_homogeneous((*C.GtkBox)(recv.native), c_homogeneous)

	return
}

// SetSpacing is a wrapper around the C function gtk_box_set_spacing.
func (recv *Box) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_box_set_spacing((*C.GtkBox)(recv.native), c_spacing)

	return
}

// Builder is a wrapper around the C record GtkBuilder.
type Builder struct {
	native *C.GtkBuilder
	// parent_instance : record
	// priv : record
}

func BuilderNewFromC(u unsafe.Pointer) *Builder {
	c := (*C.GtkBuilder)(u)
	if c == nil {
		return nil
	}

	g := &Builder{native: c}

	return g
}

func (recv *Builder) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Builder) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Builder.
// Exercise care, as this is a potentially dangerous function if the Object is not a Builder.
func CastToBuilder(object *gobject.Object) *Builder {
	return BuilderNewFromC(object.ToC())
}

// Unsupported : gtk_builder_extend_with_template : unsupported parameter template_type : no type generator for GType, GType

// Button is a wrapper around the C record GtkButton.
type Button struct {
	native *C.GtkButton
	// Private : bin
	// Private : priv
}

func ButtonNewFromC(u unsafe.Pointer) *Button {
	c := (*C.GtkButton)(u)
	if c == nil {
		return nil
	}

	g := &Button{native: c}

	return g
}

func (recv *Button) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Button) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Button) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Button) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Button) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Button) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Button.
// Exercise care, as this is a potentially dangerous function if the Object is not a Button.
func CastToButton(object *gobject.Object) *Button {
	return ButtonNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]ButtonSignalActivateCallback)
var signalActivateLock sync.Mutex

// ButtonSignalActivateCallback is a callback function for a 'activate' signal emitted from a Button.
type ButtonSignalActivateCallback func()

func Button_activateHandler() {}

var signalClickedId int
var signalClickedMap = make(map[int]ButtonSignalClickedCallback)
var signalClickedLock sync.Mutex

// ButtonSignalClickedCallback is a callback function for a 'clicked' signal emitted from a Button.
type ButtonSignalClickedCallback func()

func Button_clickedHandler() {}

var signalEnterId int
var signalEnterMap = make(map[int]ButtonSignalEnterCallback)
var signalEnterLock sync.Mutex

// ButtonSignalEnterCallback is a callback function for a 'enter' signal emitted from a Button.
type ButtonSignalEnterCallback func()

func Button_enterHandler() {}

var signalLeaveId int
var signalLeaveMap = make(map[int]ButtonSignalLeaveCallback)
var signalLeaveLock sync.Mutex

// ButtonSignalLeaveCallback is a callback function for a 'leave' signal emitted from a Button.
type ButtonSignalLeaveCallback func()

func Button_leaveHandler() {}

var signalPressedId int
var signalPressedMap = make(map[int]ButtonSignalPressedCallback)
var signalPressedLock sync.Mutex

// ButtonSignalPressedCallback is a callback function for a 'pressed' signal emitted from a Button.
type ButtonSignalPressedCallback func()

func Button_pressedHandler() {}

var signalReleasedId int
var signalReleasedMap = make(map[int]ButtonSignalReleasedCallback)
var signalReleasedLock sync.Mutex

// ButtonSignalReleasedCallback is a callback function for a 'released' signal emitted from a Button.
type ButtonSignalReleasedCallback func()

func Button_releasedHandler() {}

// ButtonNew is a wrapper around the C function gtk_button_new.
func ButtonNew() *Button {
	retC := C.gtk_button_new()
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_button_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// ButtonNewFromStock is a wrapper around the C function gtk_button_new_from_stock.
func ButtonNewFromStock(stockId string) *Button {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_button_new_from_stock(c_stock_id)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithLabel is a wrapper around the C function gtk_button_new_with_label.
func ButtonNewWithLabel(label string) *Button {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_label(c_label)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ButtonNewWithMnemonic is a wrapper around the C function gtk_button_new_with_mnemonic.
func ButtonNewWithMnemonic(label string) *Button {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_button_new_with_mnemonic(c_label)
	retGo := ButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clicked is a wrapper around the C function gtk_button_clicked.
func (recv *Button) Clicked() {
	C.gtk_button_clicked((*C.GtkButton)(recv.native))

	return
}

// Enter is a wrapper around the C function gtk_button_enter.
func (recv *Button) Enter() {
	C.gtk_button_enter((*C.GtkButton)(recv.native))

	return
}

// GetLabel is a wrapper around the C function gtk_button_get_label.
func (recv *Button) GetLabel() string {
	retC := C.gtk_button_get_label((*C.GtkButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetRelief is a wrapper around the C function gtk_button_get_relief.
func (recv *Button) GetRelief() ReliefStyle {
	retC := C.gtk_button_get_relief((*C.GtkButton)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetUseStock is a wrapper around the C function gtk_button_get_use_stock.
func (recv *Button) GetUseStock() bool {
	retC := C.gtk_button_get_use_stock((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_button_get_use_underline.
func (recv *Button) GetUseUnderline() bool {
	retC := C.gtk_button_get_use_underline((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Leave is a wrapper around the C function gtk_button_leave.
func (recv *Button) Leave() {
	C.gtk_button_leave((*C.GtkButton)(recv.native))

	return
}

// Pressed is a wrapper around the C function gtk_button_pressed.
func (recv *Button) Pressed() {
	C.gtk_button_pressed((*C.GtkButton)(recv.native))

	return
}

// Released is a wrapper around the C function gtk_button_released.
func (recv *Button) Released() {
	C.gtk_button_released((*C.GtkButton)(recv.native))

	return
}

// SetLabel is a wrapper around the C function gtk_button_set_label.
func (recv *Button) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_button_set_label((*C.GtkButton)(recv.native), c_label)

	return
}

// SetRelief is a wrapper around the C function gtk_button_set_relief.
func (recv *Button) SetRelief(relief ReliefStyle) {
	c_relief := (C.GtkReliefStyle)(relief)

	C.gtk_button_set_relief((*C.GtkButton)(recv.native), c_relief)

	return
}

// SetUseStock is a wrapper around the C function gtk_button_set_use_stock.
func (recv *Button) SetUseStock(useStock bool) {
	c_use_stock :=
		boolToGboolean(useStock)

	C.gtk_button_set_use_stock((*C.GtkButton)(recv.native), c_use_stock)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_button_set_use_underline.
func (recv *Button) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_button_set_use_underline((*C.GtkButton)(recv.native), c_use_underline)

	return
}

// ButtonAccessible is a wrapper around the C record GtkButtonAccessible.
type ButtonAccessible struct {
	native *C.GtkButtonAccessible
	// parent : record
	// priv : record
}

func ButtonAccessibleNewFromC(u unsafe.Pointer) *ButtonAccessible {
	c := (*C.GtkButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ButtonAccessible{native: c}

	return g
}

func (recv *ButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ButtonAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ButtonAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ButtonAccessible.
func CastToButtonAccessible(object *gobject.Object) *ButtonAccessible {
	return ButtonAccessibleNewFromC(object.ToC())
}

// ButtonBox is a wrapper around the C record GtkButtonBox.
type ButtonBox struct {
	native *C.GtkButtonBox
	// box : record
	// Private : priv
}

func ButtonBoxNewFromC(u unsafe.Pointer) *ButtonBox {
	c := (*C.GtkButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &ButtonBox{native: c}

	return g
}

func (recv *ButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ButtonBox) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ButtonBox) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ButtonBox) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ButtonBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ButtonBox) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ButtonBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a ButtonBox.
func CastToButtonBox(object *gobject.Object) *ButtonBox {
	return ButtonBoxNewFromC(object.ToC())
}

// GetLayout is a wrapper around the C function gtk_button_box_get_layout.
func (recv *ButtonBox) GetLayout() ButtonBoxStyle {
	retC := C.gtk_button_box_get_layout((*C.GtkButtonBox)(recv.native))
	retGo := (ButtonBoxStyle)(retC)

	return retGo
}

// SetChildSecondary is a wrapper around the C function gtk_button_box_set_child_secondary.
func (recv *ButtonBox) SetChildSecondary(child *Widget, isSecondary bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_is_secondary :=
		boolToGboolean(isSecondary)

	C.gtk_button_box_set_child_secondary((*C.GtkButtonBox)(recv.native), c_child, c_is_secondary)

	return
}

// SetLayout is a wrapper around the C function gtk_button_box_set_layout.
func (recv *ButtonBox) SetLayout(layoutStyle ButtonBoxStyle) {
	c_layout_style := (C.GtkButtonBoxStyle)(layoutStyle)

	C.gtk_button_box_set_layout((*C.GtkButtonBox)(recv.native), c_layout_style)

	return
}

// Calendar is a wrapper around the C record GtkCalendar.
type Calendar struct {
	native *C.GtkCalendar
	// widget : record
	// priv : record
}

func CalendarNewFromC(u unsafe.Pointer) *Calendar {
	c := (*C.GtkCalendar)(u)
	if c == nil {
		return nil
	}

	g := &Calendar{native: c}

	return g
}

func (recv *Calendar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Calendar) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Calendar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Calendar) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Calendar.
// Exercise care, as this is a potentially dangerous function if the Object is not a Calendar.
func CastToCalendar(object *gobject.Object) *Calendar {
	return CalendarNewFromC(object.ToC())
}

var signalDaySelectedId int
var signalDaySelectedMap = make(map[int]CalendarSignalDaySelectedCallback)
var signalDaySelectedLock sync.Mutex

// CalendarSignalDaySelectedCallback is a callback function for a 'day-selected' signal emitted from a Calendar.
type CalendarSignalDaySelectedCallback func()

func Calendar_daySelectedHandler() {}

var signalDaySelectedDoubleClickId int
var signalDaySelectedDoubleClickMap = make(map[int]CalendarSignalDaySelectedDoubleClickCallback)
var signalDaySelectedDoubleClickLock sync.Mutex

// CalendarSignalDaySelectedDoubleClickCallback is a callback function for a 'day-selected-double-click' signal emitted from a Calendar.
type CalendarSignalDaySelectedDoubleClickCallback func()

func Calendar_daySelectedDoubleClickHandler() {}

var signalMonthChangedId int
var signalMonthChangedMap = make(map[int]CalendarSignalMonthChangedCallback)
var signalMonthChangedLock sync.Mutex

// CalendarSignalMonthChangedCallback is a callback function for a 'month-changed' signal emitted from a Calendar.
type CalendarSignalMonthChangedCallback func()

func Calendar_monthChangedHandler() {}

var signalNextMonthId int
var signalNextMonthMap = make(map[int]CalendarSignalNextMonthCallback)
var signalNextMonthLock sync.Mutex

// CalendarSignalNextMonthCallback is a callback function for a 'next-month' signal emitted from a Calendar.
type CalendarSignalNextMonthCallback func()

func Calendar_nextMonthHandler() {}

var signalNextYearId int
var signalNextYearMap = make(map[int]CalendarSignalNextYearCallback)
var signalNextYearLock sync.Mutex

// CalendarSignalNextYearCallback is a callback function for a 'next-year' signal emitted from a Calendar.
type CalendarSignalNextYearCallback func()

func Calendar_nextYearHandler() {}

var signalPrevMonthId int
var signalPrevMonthMap = make(map[int]CalendarSignalPrevMonthCallback)
var signalPrevMonthLock sync.Mutex

// CalendarSignalPrevMonthCallback is a callback function for a 'prev-month' signal emitted from a Calendar.
type CalendarSignalPrevMonthCallback func()

func Calendar_prevMonthHandler() {}

var signalPrevYearId int
var signalPrevYearMap = make(map[int]CalendarSignalPrevYearCallback)
var signalPrevYearLock sync.Mutex

// CalendarSignalPrevYearCallback is a callback function for a 'prev-year' signal emitted from a Calendar.
type CalendarSignalPrevYearCallback func()

func Calendar_prevYearHandler() {}

// CalendarNew is a wrapper around the C function gtk_calendar_new.
func CalendarNew() *Calendar {
	retC := C.gtk_calendar_new()
	retGo := CalendarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ClearMarks is a wrapper around the C function gtk_calendar_clear_marks.
func (recv *Calendar) ClearMarks() {
	C.gtk_calendar_clear_marks((*C.GtkCalendar)(recv.native))

	return
}

// GetDate is a wrapper around the C function gtk_calendar_get_date.
func (recv *Calendar) GetDate() (uint32, uint32, uint32) {
	var c_year C.guint

	var c_month C.guint

	var c_day C.guint

	C.gtk_calendar_get_date((*C.GtkCalendar)(recv.native), &c_year, &c_month, &c_day)

	year := (uint32)(c_year)

	month := (uint32)(c_month)

	day := (uint32)(c_day)

	return year, month, day
}

// MarkDay is a wrapper around the C function gtk_calendar_mark_day.
func (recv *Calendar) MarkDay(day uint32) {
	c_day := (C.guint)(day)

	C.gtk_calendar_mark_day((*C.GtkCalendar)(recv.native), c_day)

	return
}

// SelectDay is a wrapper around the C function gtk_calendar_select_day.
func (recv *Calendar) SelectDay(day uint32) {
	c_day := (C.guint)(day)

	C.gtk_calendar_select_day((*C.GtkCalendar)(recv.native), c_day)

	return
}

// SelectMonth is a wrapper around the C function gtk_calendar_select_month.
func (recv *Calendar) SelectMonth(month uint32, year uint32) {
	c_month := (C.guint)(month)

	c_year := (C.guint)(year)

	C.gtk_calendar_select_month((*C.GtkCalendar)(recv.native), c_month, c_year)

	return
}

// UnmarkDay is a wrapper around the C function gtk_calendar_unmark_day.
func (recv *Calendar) UnmarkDay(day uint32) {
	c_day := (C.guint)(day)

	C.gtk_calendar_unmark_day((*C.GtkCalendar)(recv.native), c_day)

	return
}

// CellAccessible is a wrapper around the C record GtkCellAccessible.
type CellAccessible struct {
	native *C.GtkCellAccessible
	// parent : record
	// priv : record
}

func CellAccessibleNewFromC(u unsafe.Pointer) *CellAccessible {
	c := (*C.GtkCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &CellAccessible{native: c}

	return g
}

func (recv *CellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Accessible upcasts to *Accessible
func (recv *CellAccessible) Accessible() *Accessible {
	return AccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CellAccessible) Object() *atk.Object {
	return recv.Accessible().Object()
}

// CastToWidget down casts any arbitary Object to CellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellAccessible.
func CastToCellAccessible(object *gobject.Object) *CellAccessible {
	return CellAccessibleNewFromC(object.ToC())
}

// CellArea is a wrapper around the C record GtkCellArea.
type CellArea struct {
	native *C.GtkCellArea
	// Private : parent_instance
	// Private : priv
}

func CellAreaNewFromC(u unsafe.Pointer) *CellArea {
	c := (*C.GtkCellArea)(u)
	if c == nil {
		return nil
	}

	g := &CellArea{native: c}

	return g
}

func (recv *CellArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellArea) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CellArea) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to CellArea.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellArea.
func CastToCellArea(object *gobject.Object) *CellArea {
	return CellAreaNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// CellAreaBox is a wrapper around the C record GtkCellAreaBox.
type CellAreaBox struct {
	native *C.GtkCellAreaBox
	// Private : parent_instance
	// Private : priv
}

func CellAreaBoxNewFromC(u unsafe.Pointer) *CellAreaBox {
	c := (*C.GtkCellAreaBox)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaBox{native: c}

	return g
}

func (recv *CellAreaBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellArea upcasts to *CellArea
func (recv *CellAreaBox) CellArea() *CellArea {
	return CellAreaNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellAreaBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellArea().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellAreaBox) Object() *gobject.Object {
	return recv.CellArea().Object()
}

// CastToWidget down casts any arbitary Object to CellAreaBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellAreaBox.
func CastToCellAreaBox(object *gobject.Object) *CellAreaBox {
	return CellAreaBoxNewFromC(object.ToC())
}

// CellAreaContext is a wrapper around the C record GtkCellAreaContext.
type CellAreaContext struct {
	native *C.GtkCellAreaContext
	// Private : parent_instance
	// Private : priv
}

func CellAreaContextNewFromC(u unsafe.Pointer) *CellAreaContext {
	c := (*C.GtkCellAreaContext)(u)
	if c == nil {
		return nil
	}

	g := &CellAreaContext{native: c}

	return g
}

func (recv *CellAreaContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *CellAreaContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to CellAreaContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellAreaContext.
func CastToCellAreaContext(object *gobject.Object) *CellAreaContext {
	return CellAreaContextNewFromC(object.ToC())
}

// Allocate is a wrapper around the C function gtk_cell_area_context_allocate.
func (recv *CellAreaContext) Allocate(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_cell_area_context_allocate((*C.GtkCellAreaContext)(recv.native), c_width, c_height)

	return
}

// Reset is a wrapper around the C function gtk_cell_area_context_reset.
func (recv *CellAreaContext) Reset() {
	C.gtk_cell_area_context_reset((*C.GtkCellAreaContext)(recv.native))

	return
}

// CellRenderer is a wrapper around the C record GtkCellRenderer.
type CellRenderer struct {
	native *C.GtkCellRenderer
	// parent_instance : record
	// Private : priv
}

func CellRendererNewFromC(u unsafe.Pointer) *CellRenderer {
	c := (*C.GtkCellRenderer)(u)
	if c == nil {
		return nil
	}

	g := &CellRenderer{native: c}

	return g
}

func (recv *CellRenderer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRenderer) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *CellRenderer) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to CellRenderer.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRenderer.
func CastToCellRenderer(object *gobject.Object) *CellRenderer {
	return CellRendererNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter editable : no type generator for CellEditable,

// Unsupported : gtk_cell_renderer_activate : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetFixedSize is a wrapper around the C function gtk_cell_renderer_get_fixed_size.
func (recv *CellRenderer) GetFixedSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_cell_renderer_get_fixed_size((*C.GtkCellRenderer)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// Unsupported : gtk_cell_renderer_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// Unsupported : gtk_cell_renderer_render : unsupported parameter background_area : Blacklisted record : GdkRectangle

// SetFixedSize is a wrapper around the C function gtk_cell_renderer_set_fixed_size.
func (recv *CellRenderer) SetFixedSize(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_cell_renderer_set_fixed_size((*C.GtkCellRenderer)(recv.native), c_width, c_height)

	return
}

// Unsupported : gtk_cell_renderer_start_editing : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// CellRendererAccel is a wrapper around the C record GtkCellRendererAccel.
type CellRendererAccel struct {
	native *C.GtkCellRendererAccel
	// parent : record
	// Private : priv
}

func CellRendererAccelNewFromC(u unsafe.Pointer) *CellRendererAccel {
	c := (*C.GtkCellRendererAccel)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererAccel{native: c}

	return g
}

func (recv *CellRendererAccel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererText upcasts to *CellRendererText
func (recv *CellRendererAccel) CellRendererText() *CellRendererText {
	return CellRendererTextNewFromC(unsafe.Pointer(recv.native))
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererAccel) CellRenderer() *CellRenderer {
	return recv.CellRendererText().CellRenderer()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererAccel) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRendererText().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererAccel) Object() *gobject.Object {
	return recv.CellRendererText().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererAccel.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererAccel.
func CastToCellRendererAccel(object *gobject.Object) *CellRendererAccel {
	return CellRendererAccelNewFromC(object.ToC())
}

// CellRendererCombo is a wrapper around the C record GtkCellRendererCombo.
type CellRendererCombo struct {
	native *C.GtkCellRendererCombo
	// parent : record
	// Private : priv
}

func CellRendererComboNewFromC(u unsafe.Pointer) *CellRendererCombo {
	c := (*C.GtkCellRendererCombo)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererCombo{native: c}

	return g
}

func (recv *CellRendererCombo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererText upcasts to *CellRendererText
func (recv *CellRendererCombo) CellRendererText() *CellRendererText {
	return CellRendererTextNewFromC(unsafe.Pointer(recv.native))
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererCombo) CellRenderer() *CellRenderer {
	return recv.CellRendererText().CellRenderer()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererCombo) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRendererText().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererCombo) Object() *gobject.Object {
	return recv.CellRendererText().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererCombo.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererCombo.
func CastToCellRendererCombo(object *gobject.Object) *CellRendererCombo {
	return CellRendererComboNewFromC(object.ToC())
}

// CellRendererPixbuf is a wrapper around the C record GtkCellRendererPixbuf.
type CellRendererPixbuf struct {
	native *C.GtkCellRendererPixbuf
	// parent : record
	// Private : priv
}

func CellRendererPixbufNewFromC(u unsafe.Pointer) *CellRendererPixbuf {
	c := (*C.GtkCellRendererPixbuf)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererPixbuf{native: c}

	return g
}

func (recv *CellRendererPixbuf) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererPixbuf) CellRenderer() *CellRenderer {
	return CellRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererPixbuf) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererPixbuf) Object() *gobject.Object {
	return recv.CellRenderer().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererPixbuf.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererPixbuf.
func CastToCellRendererPixbuf(object *gobject.Object) *CellRendererPixbuf {
	return CellRendererPixbufNewFromC(object.ToC())
}

// CellRendererPixbufNew is a wrapper around the C function gtk_cell_renderer_pixbuf_new.
func CellRendererPixbufNew() *CellRendererPixbuf {
	retC := C.gtk_cell_renderer_pixbuf_new()
	retGo := CellRendererPixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CellRendererProgress is a wrapper around the C record GtkCellRendererProgress.
type CellRendererProgress struct {
	native *C.GtkCellRendererProgress
	// parent_instance : record
	// Private : priv
}

func CellRendererProgressNewFromC(u unsafe.Pointer) *CellRendererProgress {
	c := (*C.GtkCellRendererProgress)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererProgress{native: c}

	return g
}

func (recv *CellRendererProgress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererProgress) CellRenderer() *CellRenderer {
	return CellRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererProgress) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererProgress) Object() *gobject.Object {
	return recv.CellRenderer().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererProgress.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererProgress.
func CastToCellRendererProgress(object *gobject.Object) *CellRendererProgress {
	return CellRendererProgressNewFromC(object.ToC())
}

// CellRendererSpin is a wrapper around the C record GtkCellRendererSpin.
type CellRendererSpin struct {
	native *C.GtkCellRendererSpin
	// parent : record
	// Private : priv
}

func CellRendererSpinNewFromC(u unsafe.Pointer) *CellRendererSpin {
	c := (*C.GtkCellRendererSpin)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpin{native: c}

	return g
}

func (recv *CellRendererSpin) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRendererText upcasts to *CellRendererText
func (recv *CellRendererSpin) CellRendererText() *CellRendererText {
	return CellRendererTextNewFromC(unsafe.Pointer(recv.native))
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererSpin) CellRenderer() *CellRenderer {
	return recv.CellRendererText().CellRenderer()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererSpin) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRendererText().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererSpin) Object() *gobject.Object {
	return recv.CellRendererText().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererSpin.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererSpin.
func CastToCellRendererSpin(object *gobject.Object) *CellRendererSpin {
	return CellRendererSpinNewFromC(object.ToC())
}

// CellRendererSpinner is a wrapper around the C record GtkCellRendererSpinner.
type CellRendererSpinner struct {
	native *C.GtkCellRendererSpinner
	// parent : record
	// Private : priv
}

func CellRendererSpinnerNewFromC(u unsafe.Pointer) *CellRendererSpinner {
	c := (*C.GtkCellRendererSpinner)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererSpinner{native: c}

	return g
}

func (recv *CellRendererSpinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererSpinner) CellRenderer() *CellRenderer {
	return CellRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererSpinner) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererSpinner) Object() *gobject.Object {
	return recv.CellRenderer().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererSpinner.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererSpinner.
func CastToCellRendererSpinner(object *gobject.Object) *CellRendererSpinner {
	return CellRendererSpinnerNewFromC(object.ToC())
}

// CellRendererText is a wrapper around the C record GtkCellRendererText.
type CellRendererText struct {
	native *C.GtkCellRendererText
	// parent : record
	// Private : priv
}

func CellRendererTextNewFromC(u unsafe.Pointer) *CellRendererText {
	c := (*C.GtkCellRendererText)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererText{native: c}

	return g
}

func (recv *CellRendererText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererText) CellRenderer() *CellRenderer {
	return CellRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererText) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererText) Object() *gobject.Object {
	return recv.CellRenderer().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererText.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererText.
func CastToCellRendererText(object *gobject.Object) *CellRendererText {
	return CellRendererTextNewFromC(object.ToC())
}

var signalEditedId int
var signalEditedMap = make(map[int]CellRendererTextSignalEditedCallback)
var signalEditedLock sync.Mutex

// CellRendererTextSignalEditedCallback is a callback function for a 'edited' signal emitted from a CellRendererText.
type CellRendererTextSignalEditedCallback func(path string, newText string)

func CellRendererText_editedHandler() {}

// CellRendererTextNew is a wrapper around the C function gtk_cell_renderer_text_new.
func CellRendererTextNew() *CellRendererText {
	retC := C.gtk_cell_renderer_text_new()
	retGo := CellRendererTextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetFixedHeightFromFont is a wrapper around the C function gtk_cell_renderer_text_set_fixed_height_from_font.
func (recv *CellRendererText) SetFixedHeightFromFont(numberOfRows int32) {
	c_number_of_rows := (C.gint)(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font((*C.GtkCellRendererText)(recv.native), c_number_of_rows)

	return
}

// CellRendererToggle is a wrapper around the C record GtkCellRendererToggle.
type CellRendererToggle struct {
	native *C.GtkCellRendererToggle
	// parent : record
	// Private : priv
}

func CellRendererToggleNewFromC(u unsafe.Pointer) *CellRendererToggle {
	c := (*C.GtkCellRendererToggle)(u)
	if c == nil {
		return nil
	}

	g := &CellRendererToggle{native: c}

	return g
}

func (recv *CellRendererToggle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellRenderer upcasts to *CellRenderer
func (recv *CellRendererToggle) CellRenderer() *CellRenderer {
	return CellRendererNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellRendererToggle) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CellRenderer().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellRendererToggle) Object() *gobject.Object {
	return recv.CellRenderer().Object()
}

// CastToWidget down casts any arbitary Object to CellRendererToggle.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellRendererToggle.
func CastToCellRendererToggle(object *gobject.Object) *CellRendererToggle {
	return CellRendererToggleNewFromC(object.ToC())
}

var signalToggledId int
var signalToggledMap = make(map[int]CellRendererToggleSignalToggledCallback)
var signalToggledLock sync.Mutex

// CellRendererToggleSignalToggledCallback is a callback function for a 'toggled' signal emitted from a CellRendererToggle.
type CellRendererToggleSignalToggledCallback func(path string)

func CellRendererToggle_toggledHandler() {}

// CellRendererToggleNew is a wrapper around the C function gtk_cell_renderer_toggle_new.
func CellRendererToggleNew() *CellRendererToggle {
	retC := C.gtk_cell_renderer_toggle_new()
	retGo := CellRendererToggleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_cell_renderer_toggle_get_active.
func (recv *CellRendererToggle) GetActive() bool {
	retC := C.gtk_cell_renderer_toggle_get_active((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRadio is a wrapper around the C function gtk_cell_renderer_toggle_get_radio.
func (recv *CellRendererToggle) GetRadio() bool {
	retC := C.gtk_cell_renderer_toggle_get_radio((*C.GtkCellRendererToggle)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_cell_renderer_toggle_set_active.
func (recv *CellRendererToggle) SetActive(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_cell_renderer_toggle_set_active((*C.GtkCellRendererToggle)(recv.native), c_setting)

	return
}

// SetRadio is a wrapper around the C function gtk_cell_renderer_toggle_set_radio.
func (recv *CellRendererToggle) SetRadio(radio bool) {
	c_radio :=
		boolToGboolean(radio)

	C.gtk_cell_renderer_toggle_set_radio((*C.GtkCellRendererToggle)(recv.native), c_radio)

	return
}

// CellView is a wrapper around the C record GtkCellView.
type CellView struct {
	native *C.GtkCellView
	// parent_instance : record
	// Private : priv
}

func CellViewNewFromC(u unsafe.Pointer) *CellView {
	c := (*C.GtkCellView)(u)
	if c == nil {
		return nil
	}

	g := &CellView{native: c}

	return g
}

func (recv *CellView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *CellView) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CellView) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CellView) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to CellView.
// Exercise care, as this is a potentially dangerous function if the Object is not a CellView.
func CastToCellView(object *gobject.Object) *CellView {
	return CellViewNewFromC(object.ToC())
}

// CheckButton is a wrapper around the C record GtkCheckButton.
type CheckButton struct {
	native *C.GtkCheckButton
	// toggle_button : record
}

func CheckButtonNewFromC(u unsafe.Pointer) *CheckButton {
	c := (*C.GtkCheckButton)(u)
	if c == nil {
		return nil
	}

	g := &CheckButton{native: c}

	return g
}

func (recv *CheckButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButton upcasts to *ToggleButton
func (recv *CheckButton) ToggleButton() *ToggleButton {
	return ToggleButtonNewFromC(unsafe.Pointer(recv.native))
}

// Button upcasts to *Button
func (recv *CheckButton) Button() *Button {
	return recv.ToggleButton().Button()
}

// Bin upcasts to *Bin
func (recv *CheckButton) Bin() *Bin {
	return recv.ToggleButton().Bin()
}

// Container upcasts to *Container
func (recv *CheckButton) Container() *Container {
	return recv.ToggleButton().Container()
}

// Widget upcasts to *Widget
func (recv *CheckButton) Widget() *Widget {
	return recv.ToggleButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CheckButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToggleButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CheckButton) Object() *gobject.Object {
	return recv.ToggleButton().Object()
}

// CastToWidget down casts any arbitary Object to CheckButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a CheckButton.
func CastToCheckButton(object *gobject.Object) *CheckButton {
	return CheckButtonNewFromC(object.ToC())
}

// CheckButtonNew is a wrapper around the C function gtk_check_button_new.
func CheckButtonNew() *CheckButton {
	retC := C.gtk_check_button_new()
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithLabel is a wrapper around the C function gtk_check_button_new_with_label.
func CheckButtonNewWithLabel(label string) *CheckButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_label(c_label)
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckButtonNewWithMnemonic is a wrapper around the C function gtk_check_button_new_with_mnemonic.
func CheckButtonNewWithMnemonic(label string) *CheckButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_button_new_with_mnemonic(c_label)
	retGo := CheckButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItem is a wrapper around the C record GtkCheckMenuItem.
type CheckMenuItem struct {
	native *C.GtkCheckMenuItem
	// menu_item : record
	// Private : priv
}

func CheckMenuItemNewFromC(u unsafe.Pointer) *CheckMenuItem {
	c := (*C.GtkCheckMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItem{native: c}

	return g
}

func (recv *CheckMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem upcasts to *MenuItem
func (recv *CheckMenuItem) MenuItem() *MenuItem {
	return MenuItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *CheckMenuItem) Bin() *Bin {
	return recv.MenuItem().Bin()
}

// Container upcasts to *Container
func (recv *CheckMenuItem) Container() *Container {
	return recv.MenuItem().Container()
}

// Widget upcasts to *Widget
func (recv *CheckMenuItem) Widget() *Widget {
	return recv.MenuItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *CheckMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *CheckMenuItem) Object() *gobject.Object {
	return recv.MenuItem().Object()
}

// CastToWidget down casts any arbitary Object to CheckMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a CheckMenuItem.
func CastToCheckMenuItem(object *gobject.Object) *CheckMenuItem {
	return CheckMenuItemNewFromC(object.ToC())
}

var signalToggledId int
var signalToggledMap = make(map[int]CheckMenuItemSignalToggledCallback)
var signalToggledLock sync.Mutex

// CheckMenuItemSignalToggledCallback is a callback function for a 'toggled' signal emitted from a CheckMenuItem.
type CheckMenuItemSignalToggledCallback func()

func CheckMenuItem_toggledHandler() {}

// CheckMenuItemNew is a wrapper around the C function gtk_check_menu_item_new.
func CheckMenuItemNew() *CheckMenuItem {
	retC := C.gtk_check_menu_item_new()
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithLabel is a wrapper around the C function gtk_check_menu_item_new_with_label.
func CheckMenuItemNewWithLabel(label string) *CheckMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_label(c_label)
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CheckMenuItemNewWithMnemonic is a wrapper around the C function gtk_check_menu_item_new_with_mnemonic.
func CheckMenuItemNewWithMnemonic(label string) *CheckMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_check_menu_item_new_with_mnemonic(c_label)
	retGo := CheckMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_check_menu_item_get_active.
func (recv *CheckMenuItem) GetActive() bool {
	retC := C.gtk_check_menu_item_get_active((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInconsistent is a wrapper around the C function gtk_check_menu_item_get_inconsistent.
func (recv *CheckMenuItem) GetInconsistent() bool {
	retC := C.gtk_check_menu_item_get_inconsistent((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_check_menu_item_set_active.
func (recv *CheckMenuItem) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_check_menu_item_set_active((*C.GtkCheckMenuItem)(recv.native), c_is_active)

	return
}

// SetInconsistent is a wrapper around the C function gtk_check_menu_item_set_inconsistent.
func (recv *CheckMenuItem) SetInconsistent(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_check_menu_item_set_inconsistent((*C.GtkCheckMenuItem)(recv.native), c_setting)

	return
}

// Toggled is a wrapper around the C function gtk_check_menu_item_toggled.
func (recv *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled((*C.GtkCheckMenuItem)(recv.native))

	return
}

// CheckMenuItemAccessible is a wrapper around the C record GtkCheckMenuItemAccessible.
type CheckMenuItemAccessible struct {
	native *C.GtkCheckMenuItemAccessible
	// parent : record
	// priv : record
}

func CheckMenuItemAccessibleNewFromC(u unsafe.Pointer) *CheckMenuItemAccessible {
	c := (*C.GtkCheckMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &CheckMenuItemAccessible{native: c}

	return g
}

func (recv *CheckMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItemAccessible upcasts to *MenuItemAccessible
func (recv *CheckMenuItemAccessible) MenuItemAccessible() *MenuItemAccessible {
	return MenuItemAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *CheckMenuItemAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.MenuItemAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *CheckMenuItemAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.MenuItemAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *CheckMenuItemAccessible) Accessible() *Accessible {
	return recv.MenuItemAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *CheckMenuItemAccessible) Object() *atk.Object {
	return recv.MenuItemAccessible().Object()
}

// CastToWidget down casts any arbitary Object to CheckMenuItemAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a CheckMenuItemAccessible.
func CastToCheckMenuItemAccessible(object *gobject.Object) *CheckMenuItemAccessible {
	return CheckMenuItemAccessibleNewFromC(object.ToC())
}

// Clipboard is a wrapper around the C record GtkClipboard.
type Clipboard struct {
	native *C.GtkClipboard
}

func ClipboardNewFromC(u unsafe.Pointer) *Clipboard {
	c := (*C.GtkClipboard)(u)
	if c == nil {
		return nil
	}

	g := &Clipboard{native: c}

	return g
}

func (recv *Clipboard) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Clipboard) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Clipboard.
// Exercise care, as this is a potentially dangerous function if the Object is not a Clipboard.
func CastToClipboard(object *gobject.Object) *Clipboard {
	return ClipboardNewFromC(object.ToC())
}

// Clear is a wrapper around the C function gtk_clipboard_clear.
func (recv *Clipboard) Clear() {
	C.gtk_clipboard_clear((*C.GtkClipboard)(recv.native))

	return
}

// GetOwner is a wrapper around the C function gtk_clipboard_get_owner.
func (recv *Clipboard) GetOwner() *gobject.Object {
	retC := C.gtk_clipboard_get_owner((*C.GtkClipboard)(recv.native))
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_clipboard_request_contents : unsupported parameter target : Blacklisted record : GdkAtom

// Unsupported : gtk_clipboard_request_text : unsupported parameter callback : no type generator for ClipboardTextReceivedFunc, GtkClipboardTextReceivedFunc

// SetText is a wrapper around the C function gtk_clipboard_set_text.
func (recv *Clipboard) SetText(text string, len int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	C.gtk_clipboard_set_text((*C.GtkClipboard)(recv.native), c_text, c_len)

	return
}

// Unsupported : gtk_clipboard_set_with_data : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_set_with_owner : unsupported parameter targets : no param type

// Unsupported : gtk_clipboard_wait_for_contents : unsupported parameter target : Blacklisted record : GdkAtom

// WaitForText is a wrapper around the C function gtk_clipboard_wait_for_text.
func (recv *Clipboard) WaitForText() string {
	retC := C.gtk_clipboard_wait_for_text((*C.GtkClipboard)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// WaitIsTextAvailable is a wrapper around the C function gtk_clipboard_wait_is_text_available.
func (recv *Clipboard) WaitIsTextAvailable() bool {
	retC := C.gtk_clipboard_wait_is_text_available((*C.GtkClipboard)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ColorButton is a wrapper around the C record GtkColorButton.
type ColorButton struct {
	native *C.GtkColorButton
	// button : record
	// Private : priv
}

func ColorButtonNewFromC(u unsafe.Pointer) *ColorButton {
	c := (*C.GtkColorButton)(u)
	if c == nil {
		return nil
	}

	g := &ColorButton{native: c}

	return g
}

func (recv *ColorButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *ColorButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ColorButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *ColorButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *ColorButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ColorButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ColorButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to ColorButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorButton.
func CastToColorButton(object *gobject.Object) *ColorButton {
	return ColorButtonNewFromC(object.ToC())
}

// ColorChooserDialog is a wrapper around the C record GtkColorChooserDialog.
type ColorChooserDialog struct {
	native *C.GtkColorChooserDialog
	// parent_instance : record
	// Private : priv
}

func ColorChooserDialogNewFromC(u unsafe.Pointer) *ColorChooserDialog {
	c := (*C.GtkColorChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserDialog{native: c}

	return g
}

func (recv *ColorChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *ColorChooserDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *ColorChooserDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *ColorChooserDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *ColorChooserDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *ColorChooserDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ColorChooserDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ColorChooserDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to ColorChooserDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorChooserDialog.
func CastToColorChooserDialog(object *gobject.Object) *ColorChooserDialog {
	return ColorChooserDialogNewFromC(object.ToC())
}

// ColorChooserWidget is a wrapper around the C record GtkColorChooserWidget.
type ColorChooserWidget struct {
	native *C.GtkColorChooserWidget
	// parent_instance : record
	// Private : priv
}

func ColorChooserWidgetNewFromC(u unsafe.Pointer) *ColorChooserWidget {
	c := (*C.GtkColorChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &ColorChooserWidget{native: c}

	return g
}

func (recv *ColorChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ColorChooserWidget) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ColorChooserWidget) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ColorChooserWidget) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ColorChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ColorChooserWidget) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ColorChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorChooserWidget.
func CastToColorChooserWidget(object *gobject.Object) *ColorChooserWidget {
	return ColorChooserWidgetNewFromC(object.ToC())
}

// ColorSelection is a wrapper around the C record GtkColorSelection.
type ColorSelection struct {
	native *C.GtkColorSelection
	// parent_instance : record
	// Private : private_data
}

func ColorSelectionNewFromC(u unsafe.Pointer) *ColorSelection {
	c := (*C.GtkColorSelection)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelection{native: c}

	return g
}

func (recv *ColorSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *ColorSelection) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ColorSelection) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *ColorSelection) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ColorSelection) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ColorSelection) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to ColorSelection.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorSelection.
func CastToColorSelection(object *gobject.Object) *ColorSelection {
	return ColorSelectionNewFromC(object.ToC())
}

var signalColorChangedId int
var signalColorChangedMap = make(map[int]ColorSelectionSignalColorChangedCallback)
var signalColorChangedLock sync.Mutex

// ColorSelectionSignalColorChangedCallback is a callback function for a 'color-changed' signal emitted from a ColorSelection.
type ColorSelectionSignalColorChangedCallback func()

func ColorSelection_colorChangedHandler() {}

// ColorSelectionNew is a wrapper around the C function gtk_color_selection_new.
func ColorSelectionNew() *ColorSelection {
	retC := C.gtk_color_selection_new()
	retGo := ColorSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentAlpha is a wrapper around the C function gtk_color_selection_get_current_alpha.
func (recv *ColorSelection) GetCurrentAlpha() uint16 {
	retC := C.gtk_color_selection_get_current_alpha((*C.GtkColorSelection)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetCurrentColor is a wrapper around the C function gtk_color_selection_get_current_color.
func (recv *ColorSelection) GetCurrentColor() *gdk.Color {
	var c_color C.GdkColor

	C.gtk_color_selection_get_current_color((*C.GtkColorSelection)(recv.native), &c_color)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetHasOpacityControl is a wrapper around the C function gtk_color_selection_get_has_opacity_control.
func (recv *ColorSelection) GetHasOpacityControl() bool {
	retC := C.gtk_color_selection_get_has_opacity_control((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasPalette is a wrapper around the C function gtk_color_selection_get_has_palette.
func (recv *ColorSelection) GetHasPalette() bool {
	retC := C.gtk_color_selection_get_has_palette((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPreviousAlpha is a wrapper around the C function gtk_color_selection_get_previous_alpha.
func (recv *ColorSelection) GetPreviousAlpha() uint16 {
	retC := C.gtk_color_selection_get_previous_alpha((*C.GtkColorSelection)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetPreviousColor is a wrapper around the C function gtk_color_selection_get_previous_color.
func (recv *ColorSelection) GetPreviousColor() *gdk.Color {
	var c_color C.GdkColor

	C.gtk_color_selection_get_previous_color((*C.GtkColorSelection)(recv.native), &c_color)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return color
}

// IsAdjusting is a wrapper around the C function gtk_color_selection_is_adjusting.
func (recv *ColorSelection) IsAdjusting() bool {
	retC := C.gtk_color_selection_is_adjusting((*C.GtkColorSelection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCurrentAlpha is a wrapper around the C function gtk_color_selection_set_current_alpha.
func (recv *ColorSelection) SetCurrentAlpha(alpha uint16) {
	c_alpha := (C.guint16)(alpha)

	C.gtk_color_selection_set_current_alpha((*C.GtkColorSelection)(recv.native), c_alpha)

	return
}

// SetCurrentColor is a wrapper around the C function gtk_color_selection_set_current_color.
func (recv *ColorSelection) SetCurrentColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_color_selection_set_current_color((*C.GtkColorSelection)(recv.native), c_color)

	return
}

// SetHasOpacityControl is a wrapper around the C function gtk_color_selection_set_has_opacity_control.
func (recv *ColorSelection) SetHasOpacityControl(hasOpacity bool) {
	c_has_opacity :=
		boolToGboolean(hasOpacity)

	C.gtk_color_selection_set_has_opacity_control((*C.GtkColorSelection)(recv.native), c_has_opacity)

	return
}

// SetHasPalette is a wrapper around the C function gtk_color_selection_set_has_palette.
func (recv *ColorSelection) SetHasPalette(hasPalette bool) {
	c_has_palette :=
		boolToGboolean(hasPalette)

	C.gtk_color_selection_set_has_palette((*C.GtkColorSelection)(recv.native), c_has_palette)

	return
}

// SetPreviousAlpha is a wrapper around the C function gtk_color_selection_set_previous_alpha.
func (recv *ColorSelection) SetPreviousAlpha(alpha uint16) {
	c_alpha := (C.guint16)(alpha)

	C.gtk_color_selection_set_previous_alpha((*C.GtkColorSelection)(recv.native), c_alpha)

	return
}

// SetPreviousColor is a wrapper around the C function gtk_color_selection_set_previous_color.
func (recv *ColorSelection) SetPreviousColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_color_selection_set_previous_color((*C.GtkColorSelection)(recv.native), c_color)

	return
}

// ColorSelectionDialog is a wrapper around the C record GtkColorSelectionDialog.
type ColorSelectionDialog struct {
	native *C.GtkColorSelectionDialog
	// parent_instance : record
	// Private : priv
}

func ColorSelectionDialogNewFromC(u unsafe.Pointer) *ColorSelectionDialog {
	c := (*C.GtkColorSelectionDialog)(u)
	if c == nil {
		return nil
	}

	g := &ColorSelectionDialog{native: c}

	return g
}

func (recv *ColorSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *ColorSelectionDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *ColorSelectionDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *ColorSelectionDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *ColorSelectionDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *ColorSelectionDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ColorSelectionDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ColorSelectionDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to ColorSelectionDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a ColorSelectionDialog.
func CastToColorSelectionDialog(object *gobject.Object) *ColorSelectionDialog {
	return ColorSelectionDialogNewFromC(object.ToC())
}

// ColorSelectionDialogNew is a wrapper around the C function gtk_color_selection_dialog_new.
func ColorSelectionDialogNew(title string) *ColorSelectionDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_color_selection_dialog_new(c_title)
	retGo := ColorSelectionDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBox is a wrapper around the C record GtkComboBox.
type ComboBox struct {
	native *C.GtkComboBox
	// parent_instance : record
	// Private : priv
}

func ComboBoxNewFromC(u unsafe.Pointer) *ComboBox {
	c := (*C.GtkComboBox)(u)
	if c == nil {
		return nil
	}

	g := &ComboBox{native: c}

	return g
}

func (recv *ComboBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *ComboBox) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ComboBox) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *ComboBox) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ComboBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ComboBox) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to ComboBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a ComboBox.
func CastToComboBox(object *gobject.Object) *ComboBox {
	return ComboBoxNewFromC(object.ToC())
}

// ComboBoxNewWithArea is a wrapper around the C function gtk_combo_box_new_with_area.
func ComboBoxNewWithArea(area *CellArea) *ComboBox {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_combo_box_new_with_area(c_area)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithAreaAndEntry is a wrapper around the C function gtk_combo_box_new_with_area_and_entry.
func ComboBoxNewWithAreaAndEntry(area *CellArea) *ComboBox {
	c_area := (*C.GtkCellArea)(area.ToC())

	retC := C.gtk_combo_box_new_with_area_and_entry(c_area)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_combo_box_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// Unsupported : gtk_combo_box_new_with_model_and_entry : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// GetAddTearoffs is a wrapper around the C function gtk_combo_box_get_add_tearoffs.
func (recv *ComboBox) GetAddTearoffs() bool {
	retC := C.gtk_combo_box_get_add_tearoffs((*C.GtkComboBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ComboBoxAccessible is a wrapper around the C record GtkComboBoxAccessible.
type ComboBoxAccessible struct {
	native *C.GtkComboBoxAccessible
	// parent : record
	// priv : record
}

func ComboBoxAccessibleNewFromC(u unsafe.Pointer) *ComboBoxAccessible {
	c := (*C.GtkComboBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxAccessible{native: c}

	return g
}

func (recv *ComboBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ComboBoxAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ComboBoxAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ComboBoxAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ComboBoxAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ComboBoxAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ComboBoxAccessible.
func CastToComboBoxAccessible(object *gobject.Object) *ComboBoxAccessible {
	return ComboBoxAccessibleNewFromC(object.ToC())
}

// ComboBoxText is a wrapper around the C record GtkComboBoxText.
type ComboBoxText struct {
	native *C.GtkComboBoxText
	// Private : parent_instance
	// Private : priv
}

func ComboBoxTextNewFromC(u unsafe.Pointer) *ComboBoxText {
	c := (*C.GtkComboBoxText)(u)
	if c == nil {
		return nil
	}

	g := &ComboBoxText{native: c}

	return g
}

func (recv *ComboBoxText) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ComboBox upcasts to *ComboBox
func (recv *ComboBoxText) ComboBox() *ComboBox {
	return ComboBoxNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ComboBoxText) Bin() *Bin {
	return recv.ComboBox().Bin()
}

// Container upcasts to *Container
func (recv *ComboBoxText) Container() *Container {
	return recv.ComboBox().Container()
}

// Widget upcasts to *Widget
func (recv *ComboBoxText) Widget() *Widget {
	return recv.ComboBox().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ComboBoxText) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ComboBox().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ComboBoxText) Object() *gobject.Object {
	return recv.ComboBox().Object()
}

// CastToWidget down casts any arbitary Object to ComboBoxText.
// Exercise care, as this is a potentially dangerous function if the Object is not a ComboBoxText.
func CastToComboBoxText(object *gobject.Object) *ComboBoxText {
	return ComboBoxTextNewFromC(object.ToC())
}

// Container is a wrapper around the C record GtkContainer.
type Container struct {
	native *C.GtkContainer
	// widget : record
	// Private : priv
}

func ContainerNewFromC(u unsafe.Pointer) *Container {
	c := (*C.GtkContainer)(u)
	if c == nil {
		return nil
	}

	g := &Container{native: c}

	return g
}

func (recv *Container) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Container) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Container) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Container) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Container.
// Exercise care, as this is a potentially dangerous function if the Object is not a Container.
func CastToContainer(object *gobject.Object) *Container {
	return ContainerNewFromC(object.ToC())
}

var signalAddId int
var signalAddMap = make(map[int]ContainerSignalAddCallback)
var signalAddLock sync.Mutex

// ContainerSignalAddCallback is a callback function for a 'add' signal emitted from a Container.
type ContainerSignalAddCallback func(object *Widget)

func Container_addHandler() {}

var signalCheckResizeId int
var signalCheckResizeMap = make(map[int]ContainerSignalCheckResizeCallback)
var signalCheckResizeLock sync.Mutex

// ContainerSignalCheckResizeCallback is a callback function for a 'check-resize' signal emitted from a Container.
type ContainerSignalCheckResizeCallback func()

func Container_checkResizeHandler() {}

var signalRemoveId int
var signalRemoveMap = make(map[int]ContainerSignalRemoveCallback)
var signalRemoveLock sync.Mutex

// ContainerSignalRemoveCallback is a callback function for a 'remove' signal emitted from a Container.
type ContainerSignalRemoveCallback func(object *Widget)

func Container_removeHandler() {}

var signalSetFocusChildId int
var signalSetFocusChildMap = make(map[int]ContainerSignalSetFocusChildCallback)
var signalSetFocusChildLock sync.Mutex

// ContainerSignalSetFocusChildCallback is a callback function for a 'set-focus-child' signal emitted from a Container.
type ContainerSignalSetFocusChildCallback func(object *Widget)

func Container_setFocusChildHandler() {}

// Add is a wrapper around the C function gtk_container_add.
func (recv *Container) Add(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_container_add((*C.GtkContainer)(recv.native), c_widget)

	return
}

// Unsupported : gtk_container_add_with_properties : unsupported parameter ... : varargs

// CheckResize is a wrapper around the C function gtk_container_check_resize.
func (recv *Container) CheckResize() {
	C.gtk_container_check_resize((*C.GtkContainer)(recv.native))

	return
}

// Unsupported : gtk_container_child_get : unsupported parameter ... : varargs

// ChildGetProperty is a wrapper around the C function gtk_container_child_get_property.
func (recv *Container) ChildGetProperty(child *Widget, propertyName string, value *gobject.Value) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_container_child_get_property((*C.GtkContainer)(recv.native), c_child, c_property_name, c_value)

	return
}

// Unsupported : gtk_container_child_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_set : unsupported parameter ... : varargs

// ChildSetProperty is a wrapper around the C function gtk_container_child_set_property.
func (recv *Container) ChildSetProperty(child *Widget, propertyName string, value *gobject.Value) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_container_child_set_property((*C.GtkContainer)(recv.native), c_child, c_property_name, c_value)

	return
}

// Unsupported : gtk_container_child_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// Unsupported : gtk_container_child_type : no return generator

// Unsupported : gtk_container_forall : unsupported parameter callback : no type generator for Callback, GtkCallback

// Unsupported : gtk_container_foreach : unsupported parameter callback : no type generator for Callback, GtkCallback

// GetBorderWidth is a wrapper around the C function gtk_container_get_border_width.
func (recv *Container) GetBorderWidth() uint32 {
	retC := C.gtk_container_get_border_width((*C.GtkContainer)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetChildren is a wrapper around the C function gtk_container_get_children.
func (recv *Container) GetChildren() *glib.List {
	retC := C.gtk_container_get_children((*C.GtkContainer)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_container_get_focus_chain : unsupported parameter focusable_widgets : record with indirection level of 2

// GetFocusHadjustment is a wrapper around the C function gtk_container_get_focus_hadjustment.
func (recv *Container) GetFocusHadjustment() *Adjustment {
	retC := C.gtk_container_get_focus_hadjustment((*C.GtkContainer)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFocusVadjustment is a wrapper around the C function gtk_container_get_focus_vadjustment.
func (recv *Container) GetFocusVadjustment() *Adjustment {
	retC := C.gtk_container_get_focus_vadjustment((*C.GtkContainer)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPathForChild is a wrapper around the C function gtk_container_get_path_for_child.
func (recv *Container) GetPathForChild(child *Widget) *WidgetPath {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_container_get_path_for_child((*C.GtkContainer)(recv.native), c_child)
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetResizeMode is a wrapper around the C function gtk_container_get_resize_mode.
func (recv *Container) GetResizeMode() ResizeMode {
	retC := C.gtk_container_get_resize_mode((*C.GtkContainer)(recv.native))
	retGo := (ResizeMode)(retC)

	return retGo
}

// PropagateDraw is a wrapper around the C function gtk_container_propagate_draw.
func (recv *Container) PropagateDraw(child *Widget, cr *cairo.Context) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_cr := (*C.cairo_t)(cr.ToC())

	C.gtk_container_propagate_draw((*C.GtkContainer)(recv.native), c_child, c_cr)

	return
}

// Remove is a wrapper around the C function gtk_container_remove.
func (recv *Container) Remove(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_container_remove((*C.GtkContainer)(recv.native), c_widget)

	return
}

// ResizeChildren is a wrapper around the C function gtk_container_resize_children.
func (recv *Container) ResizeChildren() {
	C.gtk_container_resize_children((*C.GtkContainer)(recv.native))

	return
}

// SetBorderWidth is a wrapper around the C function gtk_container_set_border_width.
func (recv *Container) SetBorderWidth(borderWidth uint32) {
	c_border_width := (C.guint)(borderWidth)

	C.gtk_container_set_border_width((*C.GtkContainer)(recv.native), c_border_width)

	return
}

// SetFocusChain is a wrapper around the C function gtk_container_set_focus_chain.
func (recv *Container) SetFocusChain(focusableWidgets *glib.List) {
	c_focusable_widgets := (*C.GList)(focusableWidgets.ToC())

	C.gtk_container_set_focus_chain((*C.GtkContainer)(recv.native), c_focusable_widgets)

	return
}

// SetFocusChild is a wrapper around the C function gtk_container_set_focus_child.
func (recv *Container) SetFocusChild(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_container_set_focus_child((*C.GtkContainer)(recv.native), c_child)

	return
}

// SetFocusHadjustment is a wrapper around the C function gtk_container_set_focus_hadjustment.
func (recv *Container) SetFocusHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_container_set_focus_hadjustment((*C.GtkContainer)(recv.native), c_adjustment)

	return
}

// SetFocusVadjustment is a wrapper around the C function gtk_container_set_focus_vadjustment.
func (recv *Container) SetFocusVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_container_set_focus_vadjustment((*C.GtkContainer)(recv.native), c_adjustment)

	return
}

// SetReallocateRedraws is a wrapper around the C function gtk_container_set_reallocate_redraws.
func (recv *Container) SetReallocateRedraws(needsRedraws bool) {
	c_needs_redraws :=
		boolToGboolean(needsRedraws)

	C.gtk_container_set_reallocate_redraws((*C.GtkContainer)(recv.native), c_needs_redraws)

	return
}

// SetResizeMode is a wrapper around the C function gtk_container_set_resize_mode.
func (recv *Container) SetResizeMode(resizeMode ResizeMode) {
	c_resize_mode := (C.GtkResizeMode)(resizeMode)

	C.gtk_container_set_resize_mode((*C.GtkContainer)(recv.native), c_resize_mode)

	return
}

// UnsetFocusChain is a wrapper around the C function gtk_container_unset_focus_chain.
func (recv *Container) UnsetFocusChain() {
	C.gtk_container_unset_focus_chain((*C.GtkContainer)(recv.native))

	return
}

// ContainerAccessible is a wrapper around the C record GtkContainerAccessible.
type ContainerAccessible struct {
	native *C.GtkContainerAccessible
	// parent : record
	// priv : record
}

func ContainerAccessibleNewFromC(u unsafe.Pointer) *ContainerAccessible {
	c := (*C.GtkContainerAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ContainerAccessible{native: c}

	return g
}

func (recv *ContainerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ContainerAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *ContainerAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ContainerAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ContainerAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContainerAccessible.
func CastToContainerAccessible(object *gobject.Object) *ContainerAccessible {
	return ContainerAccessibleNewFromC(object.ToC())
}

// ContainerCellAccessible is a wrapper around the C record GtkContainerCellAccessible.
type ContainerCellAccessible struct {
	native *C.GtkContainerCellAccessible
	// parent : record
	// priv : record
}

func ContainerCellAccessibleNewFromC(u unsafe.Pointer) *ContainerCellAccessible {
	c := (*C.GtkContainerCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ContainerCellAccessible{native: c}

	return g
}

func (recv *ContainerCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessible upcasts to *CellAccessible
func (recv *ContainerCellAccessible) CellAccessible() *CellAccessible {
	return CellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *ContainerCellAccessible) Accessible() *Accessible {
	return recv.CellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ContainerCellAccessible) Object() *atk.Object {
	return recv.CellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ContainerCellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ContainerCellAccessible.
func CastToContainerCellAccessible(object *gobject.Object) *ContainerCellAccessible {
	return ContainerCellAccessibleNewFromC(object.ToC())
}

// ContainerCellAccessibleNew is a wrapper around the C function gtk_container_cell_accessible_new.
func ContainerCellAccessibleNew() *ContainerCellAccessible {
	retC := C.gtk_container_cell_accessible_new()
	retGo := ContainerCellAccessibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddChild is a wrapper around the C function gtk_container_cell_accessible_add_child.
func (recv *ContainerCellAccessible) AddChild(child *CellAccessible) {
	c_child := (*C.GtkCellAccessible)(child.ToC())

	C.gtk_container_cell_accessible_add_child((*C.GtkContainerCellAccessible)(recv.native), c_child)

	return
}

// GetChildren is a wrapper around the C function gtk_container_cell_accessible_get_children.
func (recv *ContainerCellAccessible) GetChildren() *glib.List {
	retC := C.gtk_container_cell_accessible_get_children((*C.GtkContainerCellAccessible)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveChild is a wrapper around the C function gtk_container_cell_accessible_remove_child.
func (recv *ContainerCellAccessible) RemoveChild(child *CellAccessible) {
	c_child := (*C.GtkCellAccessible)(child.ToC())

	C.gtk_container_cell_accessible_remove_child((*C.GtkContainerCellAccessible)(recv.native), c_child)

	return
}

// CssProvider is a wrapper around the C record GtkCssProvider.
type CssProvider struct {
	native *C.GtkCssProvider
	// parent_instance : record
	// priv : record
}

func CssProviderNewFromC(u unsafe.Pointer) *CssProvider {
	c := (*C.GtkCssProvider)(u)
	if c == nil {
		return nil
	}

	g := &CssProvider{native: c}

	return g
}

func (recv *CssProvider) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *CssProvider) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to CssProvider.
// Exercise care, as this is a potentially dangerous function if the Object is not a CssProvider.
func CastToCssProvider(object *gobject.Object) *CssProvider {
	return CssProviderNewFromC(object.ToC())
}

var signalParsingErrorId int
var signalParsingErrorMap = make(map[int]CssProviderSignalParsingErrorCallback)
var signalParsingErrorLock sync.Mutex

// CssProviderSignalParsingErrorCallback is a callback function for a 'parsing-error' signal emitted from a CssProvider.
type CssProviderSignalParsingErrorCallback func(section *CssSection, error *glib.Error)

func CssProvider_parsingErrorHandler() {}

// CssProviderNew is a wrapper around the C function gtk_css_provider_new.
func CssProviderNew() *CssProvider {
	retC := C.gtk_css_provider_new()
	retGo := CssProviderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_css_provider_load_from_data : unsupported parameter data : no param type

// Unsupported : gtk_css_provider_load_from_file : unsupported parameter file : no type generator for Gio.File, GFile*

// LoadFromPath is a wrapper around the C function gtk_css_provider_load_from_path.
func (recv *CssProvider) LoadFromPath(path string) (bool, error) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	var cThrowableError *C.GError

	retC := C.gtk_css_provider_load_from_path((*C.GtkCssProvider)(recv.native), c_path, &cThrowableError)
	retGo := retC == C.TRUE

	goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
	if cThrowableError != nil {
		C.g_error_free(cThrowableError)
	}

	return retGo, goThrowableError
}

// Dialog is a wrapper around the C record GtkDialog.
type Dialog struct {
	native *C.GtkDialog
	// window : record
	// Private : priv
}

func DialogNewFromC(u unsafe.Pointer) *Dialog {
	c := (*C.GtkDialog)(u)
	if c == nil {
		return nil
	}

	g := &Dialog{native: c}

	return g
}

func (recv *Dialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *Dialog) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *Dialog) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *Dialog) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *Dialog) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Dialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Dialog) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to Dialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a Dialog.
func CastToDialog(object *gobject.Object) *Dialog {
	return DialogNewFromC(object.ToC())
}

var signalCloseId int
var signalCloseMap = make(map[int]DialogSignalCloseCallback)
var signalCloseLock sync.Mutex

// DialogSignalCloseCallback is a callback function for a 'close' signal emitted from a Dialog.
type DialogSignalCloseCallback func()

func Dialog_closeHandler() {}

var signalResponseId int
var signalResponseMap = make(map[int]DialogSignalResponseCallback)
var signalResponseLock sync.Mutex

// DialogSignalResponseCallback is a callback function for a 'response' signal emitted from a Dialog.
type DialogSignalResponseCallback func(responseId int32)

func Dialog_responseHandler() {}

// DialogNew is a wrapper around the C function gtk_dialog_new.
func DialogNew() *Dialog {
	retC := C.gtk_dialog_new()
	retGo := DialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_new_with_buttons : unsupported parameter ... : varargs

// AddActionWidget is a wrapper around the C function gtk_dialog_add_action_widget.
func (recv *Dialog) AddActionWidget(child *Widget, responseId int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_response_id := (C.gint)(responseId)

	C.gtk_dialog_add_action_widget((*C.GtkDialog)(recv.native), c_child, c_response_id)

	return
}

// AddButton is a wrapper around the C function gtk_dialog_add_button.
func (recv *Dialog) AddButton(buttonText string, responseId int32) *Widget {
	c_button_text := C.CString(buttonText)
	defer C.free(unsafe.Pointer(c_button_text))

	c_response_id := (C.gint)(responseId)

	retC := C.gtk_dialog_add_button((*C.GtkDialog)(recv.native), c_button_text, c_response_id)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_dialog_add_buttons : unsupported parameter ... : varargs

// Response is a wrapper around the C function gtk_dialog_response.
func (recv *Dialog) Response(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_dialog_response((*C.GtkDialog)(recv.native), c_response_id)

	return
}

// Run is a wrapper around the C function gtk_dialog_run.
func (recv *Dialog) Run() int32 {
	retC := C.gtk_dialog_run((*C.GtkDialog)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetDefaultResponse is a wrapper around the C function gtk_dialog_set_default_response.
func (recv *Dialog) SetDefaultResponse(responseId int32) {
	c_response_id := (C.gint)(responseId)

	C.gtk_dialog_set_default_response((*C.GtkDialog)(recv.native), c_response_id)

	return
}

// SetResponseSensitive is a wrapper around the C function gtk_dialog_set_response_sensitive.
func (recv *Dialog) SetResponseSensitive(responseId int32, setting bool) {
	c_response_id := (C.gint)(responseId)

	c_setting :=
		boolToGboolean(setting)

	C.gtk_dialog_set_response_sensitive((*C.GtkDialog)(recv.native), c_response_id, c_setting)

	return
}

// DrawingArea is a wrapper around the C record GtkDrawingArea.
type DrawingArea struct {
	native *C.GtkDrawingArea
	// widget : record
	// Private : dummy
}

func DrawingAreaNewFromC(u unsafe.Pointer) *DrawingArea {
	c := (*C.GtkDrawingArea)(u)
	if c == nil {
		return nil
	}

	g := &DrawingArea{native: c}

	return g
}

func (recv *DrawingArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *DrawingArea) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *DrawingArea) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *DrawingArea) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to DrawingArea.
// Exercise care, as this is a potentially dangerous function if the Object is not a DrawingArea.
func CastToDrawingArea(object *gobject.Object) *DrawingArea {
	return DrawingAreaNewFromC(object.ToC())
}

// DrawingAreaNew is a wrapper around the C function gtk_drawing_area_new.
func DrawingAreaNew() *DrawingArea {
	retC := C.gtk_drawing_area_new()
	retGo := DrawingAreaNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Entry is a wrapper around the C record GtkEntry.
type Entry struct {
	native *C.GtkEntry
	// Private : parent_instance
	// Private : priv
}

func EntryNewFromC(u unsafe.Pointer) *Entry {
	c := (*C.GtkEntry)(u)
	if c == nil {
		return nil
	}

	g := &Entry{native: c}

	return g
}

func (recv *Entry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Entry) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Entry) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Entry) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Entry.
// Exercise care, as this is a potentially dangerous function if the Object is not a Entry.
func CastToEntry(object *gobject.Object) *Entry {
	return EntryNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]EntrySignalActivateCallback)
var signalActivateLock sync.Mutex

// EntrySignalActivateCallback is a callback function for a 'activate' signal emitted from a Entry.
type EntrySignalActivateCallback func()

func Entry_activateHandler() {}

var signalBackspaceId int
var signalBackspaceMap = make(map[int]EntrySignalBackspaceCallback)
var signalBackspaceLock sync.Mutex

// EntrySignalBackspaceCallback is a callback function for a 'backspace' signal emitted from a Entry.
type EntrySignalBackspaceCallback func()

func Entry_backspaceHandler() {}

var signalCopyClipboardId int
var signalCopyClipboardMap = make(map[int]EntrySignalCopyClipboardCallback)
var signalCopyClipboardLock sync.Mutex

// EntrySignalCopyClipboardCallback is a callback function for a 'copy-clipboard' signal emitted from a Entry.
type EntrySignalCopyClipboardCallback func()

func Entry_copyClipboardHandler() {}

var signalCutClipboardId int
var signalCutClipboardMap = make(map[int]EntrySignalCutClipboardCallback)
var signalCutClipboardLock sync.Mutex

// EntrySignalCutClipboardCallback is a callback function for a 'cut-clipboard' signal emitted from a Entry.
type EntrySignalCutClipboardCallback func()

func Entry_cutClipboardHandler() {}

var signalDeleteFromCursorId int
var signalDeleteFromCursorMap = make(map[int]EntrySignalDeleteFromCursorCallback)
var signalDeleteFromCursorLock sync.Mutex

// EntrySignalDeleteFromCursorCallback is a callback function for a 'delete-from-cursor' signal emitted from a Entry.
type EntrySignalDeleteFromCursorCallback func(type_ DeleteType, count int32)

func Entry_deleteFromCursorHandler() {}

var signalInsertAtCursorId int
var signalInsertAtCursorMap = make(map[int]EntrySignalInsertAtCursorCallback)
var signalInsertAtCursorLock sync.Mutex

// EntrySignalInsertAtCursorCallback is a callback function for a 'insert-at-cursor' signal emitted from a Entry.
type EntrySignalInsertAtCursorCallback func(string string)

func Entry_insertAtCursorHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]EntrySignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// EntrySignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a Entry.
type EntrySignalMoveCursorCallback func(step MovementStep, count int32, extendSelection bool)

func Entry_moveCursorHandler() {}

var signalPasteClipboardId int
var signalPasteClipboardMap = make(map[int]EntrySignalPasteClipboardCallback)
var signalPasteClipboardLock sync.Mutex

// EntrySignalPasteClipboardCallback is a callback function for a 'paste-clipboard' signal emitted from a Entry.
type EntrySignalPasteClipboardCallback func()

func Entry_pasteClipboardHandler() {}

var signalPopulatePopupId int
var signalPopulatePopupMap = make(map[int]EntrySignalPopulatePopupCallback)
var signalPopulatePopupLock sync.Mutex

// EntrySignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a Entry.
type EntrySignalPopulatePopupCallback func(widget *Widget)

func Entry_populatePopupHandler() {}

var signalToggleOverwriteId int
var signalToggleOverwriteMap = make(map[int]EntrySignalToggleOverwriteCallback)
var signalToggleOverwriteLock sync.Mutex

// EntrySignalToggleOverwriteCallback is a callback function for a 'toggle-overwrite' signal emitted from a Entry.
type EntrySignalToggleOverwriteCallback func()

func Entry_toggleOverwriteHandler() {}

// EntryNew is a wrapper around the C function gtk_entry_new.
func EntryNew() *Entry {
	retC := C.gtk_entry_new()
	retGo := EntryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActivatesDefault is a wrapper around the C function gtk_entry_get_activates_default.
func (recv *Entry) GetActivatesDefault() bool {
	retC := C.gtk_entry_get_activates_default((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHasFrame is a wrapper around the C function gtk_entry_get_has_frame.
func (recv *Entry) GetHasFrame() bool {
	retC := C.gtk_entry_get_has_frame((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInvisibleChar is a wrapper around the C function gtk_entry_get_invisible_char.
func (recv *Entry) GetInvisibleChar() rune {
	retC := C.gtk_entry_get_invisible_char((*C.GtkEntry)(recv.native))
	retGo := (rune)(retC)

	return retGo
}

// GetLayout is a wrapper around the C function gtk_entry_get_layout.
func (recv *Entry) GetLayout() *pango.Layout {
	retC := C.gtk_entry_get_layout((*C.GtkEntry)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLayoutOffsets is a wrapper around the C function gtk_entry_get_layout_offsets.
func (recv *Entry) GetLayoutOffsets() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_entry_get_layout_offsets((*C.GtkEntry)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetMaxLength is a wrapper around the C function gtk_entry_get_max_length.
func (recv *Entry) GetMaxLength() int32 {
	retC := C.gtk_entry_get_max_length((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_entry_get_text.
func (recv *Entry) GetText() string {
	retC := C.gtk_entry_get_text((*C.GtkEntry)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisibility is a wrapper around the C function gtk_entry_get_visibility.
func (recv *Entry) GetVisibility() bool {
	retC := C.gtk_entry_get_visibility((*C.GtkEntry)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidthChars is a wrapper around the C function gtk_entry_get_width_chars.
func (recv *Entry) GetWidthChars() int32 {
	retC := C.gtk_entry_get_width_chars((*C.GtkEntry)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// LayoutIndexToTextIndex is a wrapper around the C function gtk_entry_layout_index_to_text_index.
func (recv *Entry) LayoutIndexToTextIndex(layoutIndex int32) int32 {
	c_layout_index := (C.gint)(layoutIndex)

	retC := C.gtk_entry_layout_index_to_text_index((*C.GtkEntry)(recv.native), c_layout_index)
	retGo := (int32)(retC)

	return retGo
}

// SetActivatesDefault is a wrapper around the C function gtk_entry_set_activates_default.
func (recv *Entry) SetActivatesDefault(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_entry_set_activates_default((*C.GtkEntry)(recv.native), c_setting)

	return
}

// SetHasFrame is a wrapper around the C function gtk_entry_set_has_frame.
func (recv *Entry) SetHasFrame(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_entry_set_has_frame((*C.GtkEntry)(recv.native), c_setting)

	return
}

// SetInvisibleChar is a wrapper around the C function gtk_entry_set_invisible_char.
func (recv *Entry) SetInvisibleChar(ch rune) {
	c_ch := (C.gunichar)(ch)

	C.gtk_entry_set_invisible_char((*C.GtkEntry)(recv.native), c_ch)

	return
}

// SetMaxLength is a wrapper around the C function gtk_entry_set_max_length.
func (recv *Entry) SetMaxLength(max int32) {
	c_max := (C.gint)(max)

	C.gtk_entry_set_max_length((*C.GtkEntry)(recv.native), c_max)

	return
}

// SetText is a wrapper around the C function gtk_entry_set_text.
func (recv *Entry) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_set_text((*C.GtkEntry)(recv.native), c_text)

	return
}

// SetVisibility is a wrapper around the C function gtk_entry_set_visibility.
func (recv *Entry) SetVisibility(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_entry_set_visibility((*C.GtkEntry)(recv.native), c_visible)

	return
}

// SetWidthChars is a wrapper around the C function gtk_entry_set_width_chars.
func (recv *Entry) SetWidthChars(nChars int32) {
	c_n_chars := (C.gint)(nChars)

	C.gtk_entry_set_width_chars((*C.GtkEntry)(recv.native), c_n_chars)

	return
}

// TextIndexToLayoutIndex is a wrapper around the C function gtk_entry_text_index_to_layout_index.
func (recv *Entry) TextIndexToLayoutIndex(textIndex int32) int32 {
	c_text_index := (C.gint)(textIndex)

	retC := C.gtk_entry_text_index_to_layout_index((*C.GtkEntry)(recv.native), c_text_index)
	retGo := (int32)(retC)

	return retGo
}

// EntryAccessible is a wrapper around the C record GtkEntryAccessible.
type EntryAccessible struct {
	native *C.GtkEntryAccessible
	// parent : record
	// priv : record
}

func EntryAccessibleNewFromC(u unsafe.Pointer) *EntryAccessible {
	c := (*C.GtkEntryAccessible)(u)
	if c == nil {
		return nil
	}

	g := &EntryAccessible{native: c}

	return g
}

func (recv *EntryAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *EntryAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *EntryAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *EntryAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to EntryAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a EntryAccessible.
func CastToEntryAccessible(object *gobject.Object) *EntryAccessible {
	return EntryAccessibleNewFromC(object.ToC())
}

// EntryBuffer is a wrapper around the C record GtkEntryBuffer.
type EntryBuffer struct {
	native *C.GtkEntryBuffer
	// parent_instance : record
	// Private : priv
}

func EntryBufferNewFromC(u unsafe.Pointer) *EntryBuffer {
	c := (*C.GtkEntryBuffer)(u)
	if c == nil {
		return nil
	}

	g := &EntryBuffer{native: c}

	return g
}

func (recv *EntryBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *EntryBuffer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to EntryBuffer.
// Exercise care, as this is a potentially dangerous function if the Object is not a EntryBuffer.
func CastToEntryBuffer(object *gobject.Object) *EntryBuffer {
	return EntryBufferNewFromC(object.ToC())
}

// EntryCompletion is a wrapper around the C record GtkEntryCompletion.
type EntryCompletion struct {
	native *C.GtkEntryCompletion
	// parent_instance : record
	// Private : priv
}

func EntryCompletionNewFromC(u unsafe.Pointer) *EntryCompletion {
	c := (*C.GtkEntryCompletion)(u)
	if c == nil {
		return nil
	}

	g := &EntryCompletion{native: c}

	return g
}

func (recv *EntryCompletion) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *EntryCompletion) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to EntryCompletion.
// Exercise care, as this is a potentially dangerous function if the Object is not a EntryCompletion.
func CastToEntryCompletion(object *gobject.Object) *EntryCompletion {
	return EntryCompletionNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported signal : unsupported parameter model : no type generator for TreeModel,

// Unsupported : EntryIconAccessible : no CType

// EventBox is a wrapper around the C record GtkEventBox.
type EventBox struct {
	native *C.GtkEventBox
	// bin : record
	// Private : priv
}

func EventBoxNewFromC(u unsafe.Pointer) *EventBox {
	c := (*C.GtkEventBox)(u)
	if c == nil {
		return nil
	}

	g := &EventBox{native: c}

	return g
}

func (recv *EventBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *EventBox) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *EventBox) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *EventBox) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *EventBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *EventBox) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to EventBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a EventBox.
func CastToEventBox(object *gobject.Object) *EventBox {
	return EventBoxNewFromC(object.ToC())
}

// EventBoxNew is a wrapper around the C function gtk_event_box_new.
func EventBoxNew() *EventBox {
	retC := C.gtk_event_box_new()
	retGo := EventBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// EventController is a wrapper around the C record GtkEventController.
type EventController struct {
	native *C.GtkEventController
}

func EventControllerNewFromC(u unsafe.Pointer) *EventController {
	c := (*C.GtkEventController)(u)
	if c == nil {
		return nil
	}

	g := &EventController{native: c}

	return g
}

func (recv *EventController) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *EventController) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to EventController.
// Exercise care, as this is a potentially dangerous function if the Object is not a EventController.
func CastToEventController(object *gobject.Object) *EventController {
	return EventControllerNewFromC(object.ToC())
}

// Expander is a wrapper around the C record GtkExpander.
type Expander struct {
	native *C.GtkExpander
	// bin : record
	// priv : record
}

func ExpanderNewFromC(u unsafe.Pointer) *Expander {
	c := (*C.GtkExpander)(u)
	if c == nil {
		return nil
	}

	g := &Expander{native: c}

	return g
}

func (recv *Expander) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Expander) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Expander) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Expander) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Expander) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Expander) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Expander.
// Exercise care, as this is a potentially dangerous function if the Object is not a Expander.
func CastToExpander(object *gobject.Object) *Expander {
	return ExpanderNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]ExpanderSignalActivateCallback)
var signalActivateLock sync.Mutex

// ExpanderSignalActivateCallback is a callback function for a 'activate' signal emitted from a Expander.
type ExpanderSignalActivateCallback func()

func Expander_activateHandler() {}

// ExpanderAccessible is a wrapper around the C record GtkExpanderAccessible.
type ExpanderAccessible struct {
	native *C.GtkExpanderAccessible
	// parent : record
	// priv : record
}

func ExpanderAccessibleNewFromC(u unsafe.Pointer) *ExpanderAccessible {
	c := (*C.GtkExpanderAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ExpanderAccessible{native: c}

	return g
}

func (recv *ExpanderAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ExpanderAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ExpanderAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ExpanderAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ExpanderAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ExpanderAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ExpanderAccessible.
func CastToExpanderAccessible(object *gobject.Object) *ExpanderAccessible {
	return ExpanderAccessibleNewFromC(object.ToC())
}

// FileChooserButton is a wrapper around the C record GtkFileChooserButton.
type FileChooserButton struct {
	native *C.GtkFileChooserButton
	// parent : record
	// Private : priv
}

func FileChooserButtonNewFromC(u unsafe.Pointer) *FileChooserButton {
	c := (*C.GtkFileChooserButton)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserButton{native: c}

	return g
}

func (recv *FileChooserButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *FileChooserButton) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *FileChooserButton) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *FileChooserButton) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FileChooserButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FileChooserButton) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to FileChooserButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserButton.
func CastToFileChooserButton(object *gobject.Object) *FileChooserButton {
	return FileChooserButtonNewFromC(object.ToC())
}

// FileChooserDialog is a wrapper around the C record GtkFileChooserDialog.
type FileChooserDialog struct {
	native *C.GtkFileChooserDialog
	// parent_instance : record
	// priv : record
}

func FileChooserDialogNewFromC(u unsafe.Pointer) *FileChooserDialog {
	c := (*C.GtkFileChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserDialog{native: c}

	return g
}

func (recv *FileChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *FileChooserDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *FileChooserDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *FileChooserDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *FileChooserDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *FileChooserDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FileChooserDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FileChooserDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to FileChooserDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserDialog.
func CastToFileChooserDialog(object *gobject.Object) *FileChooserDialog {
	return FileChooserDialogNewFromC(object.ToC())
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// FileChooserWidget is a wrapper around the C record GtkFileChooserWidget.
type FileChooserWidget struct {
	native *C.GtkFileChooserWidget
	// parent_instance : record
	// priv : record
}

func FileChooserWidgetNewFromC(u unsafe.Pointer) *FileChooserWidget {
	c := (*C.GtkFileChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &FileChooserWidget{native: c}

	return g
}

func (recv *FileChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *FileChooserWidget) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *FileChooserWidget) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *FileChooserWidget) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FileChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FileChooserWidget) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to FileChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileChooserWidget.
func CastToFileChooserWidget(object *gobject.Object) *FileChooserWidget {
	return FileChooserWidgetNewFromC(object.ToC())
}

var signalDesktopFolderId int
var signalDesktopFolderMap = make(map[int]FileChooserWidgetSignalDesktopFolderCallback)
var signalDesktopFolderLock sync.Mutex

// FileChooserWidgetSignalDesktopFolderCallback is a callback function for a 'desktop-folder' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalDesktopFolderCallback func()

func FileChooserWidget_desktopFolderHandler() {}

var signalDownFolderId int
var signalDownFolderMap = make(map[int]FileChooserWidgetSignalDownFolderCallback)
var signalDownFolderLock sync.Mutex

// FileChooserWidgetSignalDownFolderCallback is a callback function for a 'down-folder' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalDownFolderCallback func()

func FileChooserWidget_downFolderHandler() {}

var signalHomeFolderId int
var signalHomeFolderMap = make(map[int]FileChooserWidgetSignalHomeFolderCallback)
var signalHomeFolderLock sync.Mutex

// FileChooserWidgetSignalHomeFolderCallback is a callback function for a 'home-folder' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalHomeFolderCallback func()

func FileChooserWidget_homeFolderHandler() {}

var signalLocationPopupId int
var signalLocationPopupMap = make(map[int]FileChooserWidgetSignalLocationPopupCallback)
var signalLocationPopupLock sync.Mutex

// FileChooserWidgetSignalLocationPopupCallback is a callback function for a 'location-popup' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalLocationPopupCallback func(path string)

func FileChooserWidget_locationPopupHandler() {}

var signalLocationPopupOnPasteId int
var signalLocationPopupOnPasteMap = make(map[int]FileChooserWidgetSignalLocationPopupOnPasteCallback)
var signalLocationPopupOnPasteLock sync.Mutex

// FileChooserWidgetSignalLocationPopupOnPasteCallback is a callback function for a 'location-popup-on-paste' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalLocationPopupOnPasteCallback func()

func FileChooserWidget_locationPopupOnPasteHandler() {}

var signalLocationTogglePopupId int
var signalLocationTogglePopupMap = make(map[int]FileChooserWidgetSignalLocationTogglePopupCallback)
var signalLocationTogglePopupLock sync.Mutex

// FileChooserWidgetSignalLocationTogglePopupCallback is a callback function for a 'location-toggle-popup' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalLocationTogglePopupCallback func()

func FileChooserWidget_locationTogglePopupHandler() {}

var signalPlacesShortcutId int
var signalPlacesShortcutMap = make(map[int]FileChooserWidgetSignalPlacesShortcutCallback)
var signalPlacesShortcutLock sync.Mutex

// FileChooserWidgetSignalPlacesShortcutCallback is a callback function for a 'places-shortcut' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalPlacesShortcutCallback func()

func FileChooserWidget_placesShortcutHandler() {}

var signalQuickBookmarkId int
var signalQuickBookmarkMap = make(map[int]FileChooserWidgetSignalQuickBookmarkCallback)
var signalQuickBookmarkLock sync.Mutex

// FileChooserWidgetSignalQuickBookmarkCallback is a callback function for a 'quick-bookmark' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalQuickBookmarkCallback func(bookmarkIndex int32)

func FileChooserWidget_quickBookmarkHandler() {}

var signalRecentShortcutId int
var signalRecentShortcutMap = make(map[int]FileChooserWidgetSignalRecentShortcutCallback)
var signalRecentShortcutLock sync.Mutex

// FileChooserWidgetSignalRecentShortcutCallback is a callback function for a 'recent-shortcut' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalRecentShortcutCallback func()

func FileChooserWidget_recentShortcutHandler() {}

var signalSearchShortcutId int
var signalSearchShortcutMap = make(map[int]FileChooserWidgetSignalSearchShortcutCallback)
var signalSearchShortcutLock sync.Mutex

// FileChooserWidgetSignalSearchShortcutCallback is a callback function for a 'search-shortcut' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalSearchShortcutCallback func()

func FileChooserWidget_searchShortcutHandler() {}

var signalShowHiddenId int
var signalShowHiddenMap = make(map[int]FileChooserWidgetSignalShowHiddenCallback)
var signalShowHiddenLock sync.Mutex

// FileChooserWidgetSignalShowHiddenCallback is a callback function for a 'show-hidden' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalShowHiddenCallback func()

func FileChooserWidget_showHiddenHandler() {}

var signalUpFolderId int
var signalUpFolderMap = make(map[int]FileChooserWidgetSignalUpFolderCallback)
var signalUpFolderLock sync.Mutex

// FileChooserWidgetSignalUpFolderCallback is a callback function for a 'up-folder' signal emitted from a FileChooserWidget.
type FileChooserWidgetSignalUpFolderCallback func()

func FileChooserWidget_upFolderHandler() {}

// FileFilter is a wrapper around the C record GtkFileFilter.
type FileFilter struct {
	native *C.GtkFileFilter
}

func FileFilterNewFromC(u unsafe.Pointer) *FileFilter {
	c := (*C.GtkFileFilter)(u)
	if c == nil {
		return nil
	}

	g := &FileFilter{native: c}

	return g
}

func (recv *FileFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FileFilter) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *FileFilter) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to FileFilter.
// Exercise care, as this is a potentially dangerous function if the Object is not a FileFilter.
func CastToFileFilter(object *gobject.Object) *FileFilter {
	return FileFilterNewFromC(object.ToC())
}

// Unsupported : gtk_file_filter_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Fixed is a wrapper around the C record GtkFixed.
type Fixed struct {
	native *C.GtkFixed
	// container : record
	// Private : priv
}

func FixedNewFromC(u unsafe.Pointer) *Fixed {
	c := (*C.GtkFixed)(u)
	if c == nil {
		return nil
	}

	g := &Fixed{native: c}

	return g
}

func (recv *Fixed) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Fixed) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Fixed) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Fixed) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Fixed) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Fixed.
// Exercise care, as this is a potentially dangerous function if the Object is not a Fixed.
func CastToFixed(object *gobject.Object) *Fixed {
	return FixedNewFromC(object.ToC())
}

// FixedNew is a wrapper around the C function gtk_fixed_new.
func FixedNew() *Fixed {
	retC := C.gtk_fixed_new()
	retGo := FixedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Move is a wrapper around the C function gtk_fixed_move.
func (recv *Fixed) Move(widget *Widget, x int32, y int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_fixed_move((*C.GtkFixed)(recv.native), c_widget, c_x, c_y)

	return
}

// Put is a wrapper around the C function gtk_fixed_put.
func (recv *Fixed) Put(widget *Widget, x int32, y int32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_fixed_put((*C.GtkFixed)(recv.native), c_widget, c_x, c_y)

	return
}

// FlowBox is a wrapper around the C record GtkFlowBox.
type FlowBox struct {
	native *C.GtkFlowBox
	// container : record
}

func FlowBoxNewFromC(u unsafe.Pointer) *FlowBox {
	c := (*C.GtkFlowBox)(u)
	if c == nil {
		return nil
	}

	g := &FlowBox{native: c}

	return g
}

func (recv *FlowBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *FlowBox) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *FlowBox) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FlowBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FlowBox) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to FlowBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a FlowBox.
func CastToFlowBox(object *gobject.Object) *FlowBox {
	return FlowBoxNewFromC(object.ToC())
}

var signalActivateCursorChildId int
var signalActivateCursorChildMap = make(map[int]FlowBoxSignalActivateCursorChildCallback)
var signalActivateCursorChildLock sync.Mutex

// FlowBoxSignalActivateCursorChildCallback is a callback function for a 'activate-cursor-child' signal emitted from a FlowBox.
type FlowBoxSignalActivateCursorChildCallback func()

func FlowBox_activateCursorChildHandler() {}

var signalChildActivatedId int
var signalChildActivatedMap = make(map[int]FlowBoxSignalChildActivatedCallback)
var signalChildActivatedLock sync.Mutex

// FlowBoxSignalChildActivatedCallback is a callback function for a 'child-activated' signal emitted from a FlowBox.
type FlowBoxSignalChildActivatedCallback func(child *FlowBoxChild)

func FlowBox_childActivatedHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]FlowBoxSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// FlowBoxSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a FlowBox.
type FlowBoxSignalMoveCursorCallback func(step MovementStep, count int32) bool

func FlowBox_moveCursorHandler() {}

var signalSelectAllId int
var signalSelectAllMap = make(map[int]FlowBoxSignalSelectAllCallback)
var signalSelectAllLock sync.Mutex

// FlowBoxSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a FlowBox.
type FlowBoxSignalSelectAllCallback func()

func FlowBox_selectAllHandler() {}

var signalSelectedChildrenChangedId int
var signalSelectedChildrenChangedMap = make(map[int]FlowBoxSignalSelectedChildrenChangedCallback)
var signalSelectedChildrenChangedLock sync.Mutex

// FlowBoxSignalSelectedChildrenChangedCallback is a callback function for a 'selected-children-changed' signal emitted from a FlowBox.
type FlowBoxSignalSelectedChildrenChangedCallback func()

func FlowBox_selectedChildrenChangedHandler() {}

var signalToggleCursorChildId int
var signalToggleCursorChildMap = make(map[int]FlowBoxSignalToggleCursorChildCallback)
var signalToggleCursorChildLock sync.Mutex

// FlowBoxSignalToggleCursorChildCallback is a callback function for a 'toggle-cursor-child' signal emitted from a FlowBox.
type FlowBoxSignalToggleCursorChildCallback func()

func FlowBox_toggleCursorChildHandler() {}

var signalUnselectAllId int
var signalUnselectAllMap = make(map[int]FlowBoxSignalUnselectAllCallback)
var signalUnselectAllLock sync.Mutex

// FlowBoxSignalUnselectAllCallback is a callback function for a 'unselect-all' signal emitted from a FlowBox.
type FlowBoxSignalUnselectAllCallback func()

func FlowBox_unselectAllHandler() {}

// FlowBoxAccessible is a wrapper around the C record GtkFlowBoxAccessible.
type FlowBoxAccessible struct {
	native *C.GtkFlowBoxAccessible
	// parent : record
	// priv : record
}

func FlowBoxAccessibleNewFromC(u unsafe.Pointer) *FlowBoxAccessible {
	c := (*C.GtkFlowBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxAccessible{native: c}

	return g
}

func (recv *FlowBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *FlowBoxAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *FlowBoxAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *FlowBoxAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *FlowBoxAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to FlowBoxAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a FlowBoxAccessible.
func CastToFlowBoxAccessible(object *gobject.Object) *FlowBoxAccessible {
	return FlowBoxAccessibleNewFromC(object.ToC())
}

// FlowBoxChild is a wrapper around the C record GtkFlowBoxChild.
type FlowBoxChild struct {
	native *C.GtkFlowBoxChild
	// parent_instance : record
}

func FlowBoxChildNewFromC(u unsafe.Pointer) *FlowBoxChild {
	c := (*C.GtkFlowBoxChild)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChild{native: c}

	return g
}

func (recv *FlowBoxChild) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *FlowBoxChild) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *FlowBoxChild) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *FlowBoxChild) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FlowBoxChild) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FlowBoxChild) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to FlowBoxChild.
// Exercise care, as this is a potentially dangerous function if the Object is not a FlowBoxChild.
func CastToFlowBoxChild(object *gobject.Object) *FlowBoxChild {
	return FlowBoxChildNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]FlowBoxChildSignalActivateCallback)
var signalActivateLock sync.Mutex

// FlowBoxChildSignalActivateCallback is a callback function for a 'activate' signal emitted from a FlowBoxChild.
type FlowBoxChildSignalActivateCallback func()

func FlowBoxChild_activateHandler() {}

// FlowBoxChildAccessible is a wrapper around the C record GtkFlowBoxChildAccessible.
type FlowBoxChildAccessible struct {
	native *C.GtkFlowBoxChildAccessible
	// parent : record
}

func FlowBoxChildAccessibleNewFromC(u unsafe.Pointer) *FlowBoxChildAccessible {
	c := (*C.GtkFlowBoxChildAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FlowBoxChildAccessible{native: c}

	return g
}

func (recv *FlowBoxChildAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *FlowBoxChildAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *FlowBoxChildAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *FlowBoxChildAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *FlowBoxChildAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to FlowBoxChildAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a FlowBoxChildAccessible.
func CastToFlowBoxChildAccessible(object *gobject.Object) *FlowBoxChildAccessible {
	return FlowBoxChildAccessibleNewFromC(object.ToC())
}

// FontButton is a wrapper around the C record GtkFontButton.
type FontButton struct {
	native *C.GtkFontButton
	// button : record
	// Private : priv
}

func FontButtonNewFromC(u unsafe.Pointer) *FontButton {
	c := (*C.GtkFontButton)(u)
	if c == nil {
		return nil
	}

	g := &FontButton{native: c}

	return g
}

func (recv *FontButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *FontButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *FontButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *FontButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *FontButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FontButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FontButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to FontButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontButton.
func CastToFontButton(object *gobject.Object) *FontButton {
	return FontButtonNewFromC(object.ToC())
}

// FontChooserDialog is a wrapper around the C record GtkFontChooserDialog.
type FontChooserDialog struct {
	native *C.GtkFontChooserDialog
	// parent_instance : record
	// Private : priv
}

func FontChooserDialogNewFromC(u unsafe.Pointer) *FontChooserDialog {
	c := (*C.GtkFontChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserDialog{native: c}

	return g
}

func (recv *FontChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *FontChooserDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *FontChooserDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *FontChooserDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *FontChooserDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *FontChooserDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FontChooserDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FontChooserDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to FontChooserDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontChooserDialog.
func CastToFontChooserDialog(object *gobject.Object) *FontChooserDialog {
	return FontChooserDialogNewFromC(object.ToC())
}

// FontChooserWidget is a wrapper around the C record GtkFontChooserWidget.
type FontChooserWidget struct {
	native *C.GtkFontChooserWidget
	// parent_instance : record
	// Private : priv
}

func FontChooserWidgetNewFromC(u unsafe.Pointer) *FontChooserWidget {
	c := (*C.GtkFontChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &FontChooserWidget{native: c}

	return g
}

func (recv *FontChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *FontChooserWidget) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *FontChooserWidget) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *FontChooserWidget) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FontChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FontChooserWidget) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to FontChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontChooserWidget.
func CastToFontChooserWidget(object *gobject.Object) *FontChooserWidget {
	return FontChooserWidgetNewFromC(object.ToC())
}

// FontSelection is a wrapper around the C record GtkFontSelection.
type FontSelection struct {
	native *C.GtkFontSelection
	// parent_instance : record
	// Private : priv
}

func FontSelectionNewFromC(u unsafe.Pointer) *FontSelection {
	c := (*C.GtkFontSelection)(u)
	if c == nil {
		return nil
	}

	g := &FontSelection{native: c}

	return g
}

func (recv *FontSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *FontSelection) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *FontSelection) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *FontSelection) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FontSelection) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FontSelection) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to FontSelection.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontSelection.
func CastToFontSelection(object *gobject.Object) *FontSelection {
	return FontSelectionNewFromC(object.ToC())
}

// FontSelectionNew is a wrapper around the C function gtk_font_selection_new.
func FontSelectionNew() *FontSelection {
	retC := C.gtk_font_selection_new()
	retGo := FontSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_selection_get_font_name.
func (recv *FontSelection) GetFontName() string {
	retC := C.gtk_font_selection_get_font_name((*C.GtkFontSelection)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewText is a wrapper around the C function gtk_font_selection_get_preview_text.
func (recv *FontSelection) GetPreviewText() string {
	retC := C.gtk_font_selection_get_preview_text((*C.GtkFontSelection)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_selection_set_font_name.
func (recv *FontSelection) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_selection_set_font_name((*C.GtkFontSelection)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// SetPreviewText is a wrapper around the C function gtk_font_selection_set_preview_text.
func (recv *FontSelection) SetPreviewText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_font_selection_set_preview_text((*C.GtkFontSelection)(recv.native), c_text)

	return
}

// FontSelectionDialog is a wrapper around the C record GtkFontSelectionDialog.
type FontSelectionDialog struct {
	native *C.GtkFontSelectionDialog
	// parent_instance : record
	// Private : priv
}

func FontSelectionDialogNewFromC(u unsafe.Pointer) *FontSelectionDialog {
	c := (*C.GtkFontSelectionDialog)(u)
	if c == nil {
		return nil
	}

	g := &FontSelectionDialog{native: c}

	return g
}

func (recv *FontSelectionDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *FontSelectionDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *FontSelectionDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *FontSelectionDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *FontSelectionDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *FontSelectionDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *FontSelectionDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *FontSelectionDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to FontSelectionDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a FontSelectionDialog.
func CastToFontSelectionDialog(object *gobject.Object) *FontSelectionDialog {
	return FontSelectionDialogNewFromC(object.ToC())
}

// FontSelectionDialogNew is a wrapper around the C function gtk_font_selection_dialog_new.
func FontSelectionDialogNew(title string) *FontSelectionDialog {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	retC := C.gtk_font_selection_dialog_new(c_title)
	retGo := FontSelectionDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_selection_dialog_get_font_name.
func (recv *FontSelectionDialog) GetFontName() string {
	retC := C.gtk_font_selection_dialog_get_font_name((*C.GtkFontSelectionDialog)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewText is a wrapper around the C function gtk_font_selection_dialog_get_preview_text.
func (recv *FontSelectionDialog) GetPreviewText() string {
	retC := C.gtk_font_selection_dialog_get_preview_text((*C.GtkFontSelectionDialog)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_selection_dialog_set_font_name.
func (recv *FontSelectionDialog) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_selection_dialog_set_font_name((*C.GtkFontSelectionDialog)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// SetPreviewText is a wrapper around the C function gtk_font_selection_dialog_set_preview_text.
func (recv *FontSelectionDialog) SetPreviewText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_font_selection_dialog_set_preview_text((*C.GtkFontSelectionDialog)(recv.native), c_text)

	return
}

// Frame is a wrapper around the C record GtkFrame.
type Frame struct {
	native *C.GtkFrame
	// bin : record
	// Private : priv
}

func FrameNewFromC(u unsafe.Pointer) *Frame {
	c := (*C.GtkFrame)(u)
	if c == nil {
		return nil
	}

	g := &Frame{native: c}

	return g
}

func (recv *Frame) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Frame) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Frame) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Frame) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Frame) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Frame) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Frame.
// Exercise care, as this is a potentially dangerous function if the Object is not a Frame.
func CastToFrame(object *gobject.Object) *Frame {
	return FrameNewFromC(object.ToC())
}

// FrameNew is a wrapper around the C function gtk_frame_new.
func FrameNew(label string) *Frame {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_frame_new(c_label)
	retGo := FrameNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLabel is a wrapper around the C function gtk_frame_get_label.
func (recv *Frame) GetLabel() string {
	retC := C.gtk_frame_get_label((*C.GtkFrame)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelAlign is a wrapper around the C function gtk_frame_get_label_align.
func (recv *Frame) GetLabelAlign() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_frame_get_label_align((*C.GtkFrame)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetLabelWidget is a wrapper around the C function gtk_frame_get_label_widget.
func (recv *Frame) GetLabelWidget() *Widget {
	retC := C.gtk_frame_get_label_widget((*C.GtkFrame)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_frame_get_shadow_type.
func (recv *Frame) GetShadowType() ShadowType {
	retC := C.gtk_frame_get_shadow_type((*C.GtkFrame)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// SetLabel is a wrapper around the C function gtk_frame_set_label.
func (recv *Frame) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_frame_set_label((*C.GtkFrame)(recv.native), c_label)

	return
}

// SetLabelAlign is a wrapper around the C function gtk_frame_set_label_align.
func (recv *Frame) SetLabelAlign(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_frame_set_label_align((*C.GtkFrame)(recv.native), c_xalign, c_yalign)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_frame_set_label_widget.
func (recv *Frame) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(labelWidget.ToC())

	C.gtk_frame_set_label_widget((*C.GtkFrame)(recv.native), c_label_widget)

	return
}

// SetShadowType is a wrapper around the C function gtk_frame_set_shadow_type.
func (recv *Frame) SetShadowType(type_ ShadowType) {
	c_type := (C.GtkShadowType)(type_)

	C.gtk_frame_set_shadow_type((*C.GtkFrame)(recv.native), c_type)

	return
}

// FrameAccessible is a wrapper around the C record GtkFrameAccessible.
type FrameAccessible struct {
	native *C.GtkFrameAccessible
	// parent : record
	// priv : record
}

func FrameAccessibleNewFromC(u unsafe.Pointer) *FrameAccessible {
	c := (*C.GtkFrameAccessible)(u)
	if c == nil {
		return nil
	}

	g := &FrameAccessible{native: c}

	return g
}

func (recv *FrameAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *FrameAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *FrameAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *FrameAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *FrameAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to FrameAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a FrameAccessible.
func CastToFrameAccessible(object *gobject.Object) *FrameAccessible {
	return FrameAccessibleNewFromC(object.ToC())
}

// Gesture is a wrapper around the C record GtkGesture.
type Gesture struct {
	native *C.GtkGesture
}

func GestureNewFromC(u unsafe.Pointer) *Gesture {
	c := (*C.GtkGesture)(u)
	if c == nil {
		return nil
	}

	g := &Gesture{native: c}

	return g
}

func (recv *Gesture) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EventController upcasts to *EventController
func (recv *Gesture) EventController() *EventController {
	return EventControllerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Gesture) Object() *gobject.Object {
	return recv.EventController().Object()
}

// CastToWidget down casts any arbitary Object to Gesture.
// Exercise care, as this is a potentially dangerous function if the Object is not a Gesture.
func CastToGesture(object *gobject.Object) *Gesture {
	return GestureNewFromC(object.ToC())
}

// Unsupported : gtk_gesture_get_last_event : no return generator

// GestureDrag is a wrapper around the C record GtkGestureDrag.
type GestureDrag struct {
	native *C.GtkGestureDrag
}

func GestureDragNewFromC(u unsafe.Pointer) *GestureDrag {
	c := (*C.GtkGestureDrag)(u)
	if c == nil {
		return nil
	}

	g := &GestureDrag{native: c}

	return g
}

func (recv *GestureDrag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle upcasts to *GestureSingle
func (recv *GestureDrag) GestureSingle() *GestureSingle {
	return GestureSingleNewFromC(unsafe.Pointer(recv.native))
}

// Gesture upcasts to *Gesture
func (recv *GestureDrag) Gesture() *Gesture {
	return recv.GestureSingle().Gesture()
}

// EventController upcasts to *EventController
func (recv *GestureDrag) EventController() *EventController {
	return recv.GestureSingle().EventController()
}

// Object upcasts to *Object
func (recv *GestureDrag) Object() *gobject.Object {
	return recv.GestureSingle().Object()
}

// CastToWidget down casts any arbitary Object to GestureDrag.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureDrag.
func CastToGestureDrag(object *gobject.Object) *GestureDrag {
	return GestureDragNewFromC(object.ToC())
}

// GestureLongPress is a wrapper around the C record GtkGestureLongPress.
type GestureLongPress struct {
	native *C.GtkGestureLongPress
}

func GestureLongPressNewFromC(u unsafe.Pointer) *GestureLongPress {
	c := (*C.GtkGestureLongPress)(u)
	if c == nil {
		return nil
	}

	g := &GestureLongPress{native: c}

	return g
}

func (recv *GestureLongPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle upcasts to *GestureSingle
func (recv *GestureLongPress) GestureSingle() *GestureSingle {
	return GestureSingleNewFromC(unsafe.Pointer(recv.native))
}

// Gesture upcasts to *Gesture
func (recv *GestureLongPress) Gesture() *Gesture {
	return recv.GestureSingle().Gesture()
}

// EventController upcasts to *EventController
func (recv *GestureLongPress) EventController() *EventController {
	return recv.GestureSingle().EventController()
}

// Object upcasts to *Object
func (recv *GestureLongPress) Object() *gobject.Object {
	return recv.GestureSingle().Object()
}

// CastToWidget down casts any arbitary Object to GestureLongPress.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureLongPress.
func CastToGestureLongPress(object *gobject.Object) *GestureLongPress {
	return GestureLongPressNewFromC(object.ToC())
}

// GestureMultiPress is a wrapper around the C record GtkGestureMultiPress.
type GestureMultiPress struct {
	native *C.GtkGestureMultiPress
}

func GestureMultiPressNewFromC(u unsafe.Pointer) *GestureMultiPress {
	c := (*C.GtkGestureMultiPress)(u)
	if c == nil {
		return nil
	}

	g := &GestureMultiPress{native: c}

	return g
}

func (recv *GestureMultiPress) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle upcasts to *GestureSingle
func (recv *GestureMultiPress) GestureSingle() *GestureSingle {
	return GestureSingleNewFromC(unsafe.Pointer(recv.native))
}

// Gesture upcasts to *Gesture
func (recv *GestureMultiPress) Gesture() *Gesture {
	return recv.GestureSingle().Gesture()
}

// EventController upcasts to *EventController
func (recv *GestureMultiPress) EventController() *EventController {
	return recv.GestureSingle().EventController()
}

// Object upcasts to *Object
func (recv *GestureMultiPress) Object() *gobject.Object {
	return recv.GestureSingle().Object()
}

// CastToWidget down casts any arbitary Object to GestureMultiPress.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureMultiPress.
func CastToGestureMultiPress(object *gobject.Object) *GestureMultiPress {
	return GestureMultiPressNewFromC(object.ToC())
}

// GesturePan is a wrapper around the C record GtkGesturePan.
type GesturePan struct {
	native *C.GtkGesturePan
}

func GesturePanNewFromC(u unsafe.Pointer) *GesturePan {
	c := (*C.GtkGesturePan)(u)
	if c == nil {
		return nil
	}

	g := &GesturePan{native: c}

	return g
}

func (recv *GesturePan) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureDrag upcasts to *GestureDrag
func (recv *GesturePan) GestureDrag() *GestureDrag {
	return GestureDragNewFromC(unsafe.Pointer(recv.native))
}

// GestureSingle upcasts to *GestureSingle
func (recv *GesturePan) GestureSingle() *GestureSingle {
	return recv.GestureDrag().GestureSingle()
}

// Gesture upcasts to *Gesture
func (recv *GesturePan) Gesture() *Gesture {
	return recv.GestureDrag().Gesture()
}

// EventController upcasts to *EventController
func (recv *GesturePan) EventController() *EventController {
	return recv.GestureDrag().EventController()
}

// Object upcasts to *Object
func (recv *GesturePan) Object() *gobject.Object {
	return recv.GestureDrag().Object()
}

// CastToWidget down casts any arbitary Object to GesturePan.
// Exercise care, as this is a potentially dangerous function if the Object is not a GesturePan.
func CastToGesturePan(object *gobject.Object) *GesturePan {
	return GesturePanNewFromC(object.ToC())
}

// GestureRotate is a wrapper around the C record GtkGestureRotate.
type GestureRotate struct {
	native *C.GtkGestureRotate
}

func GestureRotateNewFromC(u unsafe.Pointer) *GestureRotate {
	c := (*C.GtkGestureRotate)(u)
	if c == nil {
		return nil
	}

	g := &GestureRotate{native: c}

	return g
}

func (recv *GestureRotate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gesture upcasts to *Gesture
func (recv *GestureRotate) Gesture() *Gesture {
	return GestureNewFromC(unsafe.Pointer(recv.native))
}

// EventController upcasts to *EventController
func (recv *GestureRotate) EventController() *EventController {
	return recv.Gesture().EventController()
}

// Object upcasts to *Object
func (recv *GestureRotate) Object() *gobject.Object {
	return recv.Gesture().Object()
}

// CastToWidget down casts any arbitary Object to GestureRotate.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureRotate.
func CastToGestureRotate(object *gobject.Object) *GestureRotate {
	return GestureRotateNewFromC(object.ToC())
}

// GestureSingle is a wrapper around the C record GtkGestureSingle.
type GestureSingle struct {
	native *C.GtkGestureSingle
}

func GestureSingleNewFromC(u unsafe.Pointer) *GestureSingle {
	c := (*C.GtkGestureSingle)(u)
	if c == nil {
		return nil
	}

	g := &GestureSingle{native: c}

	return g
}

func (recv *GestureSingle) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gesture upcasts to *Gesture
func (recv *GestureSingle) Gesture() *Gesture {
	return GestureNewFromC(unsafe.Pointer(recv.native))
}

// EventController upcasts to *EventController
func (recv *GestureSingle) EventController() *EventController {
	return recv.Gesture().EventController()
}

// Object upcasts to *Object
func (recv *GestureSingle) Object() *gobject.Object {
	return recv.Gesture().Object()
}

// CastToWidget down casts any arbitary Object to GestureSingle.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureSingle.
func CastToGestureSingle(object *gobject.Object) *GestureSingle {
	return GestureSingleNewFromC(object.ToC())
}

// GestureSwipe is a wrapper around the C record GtkGestureSwipe.
type GestureSwipe struct {
	native *C.GtkGestureSwipe
}

func GestureSwipeNewFromC(u unsafe.Pointer) *GestureSwipe {
	c := (*C.GtkGestureSwipe)(u)
	if c == nil {
		return nil
	}

	g := &GestureSwipe{native: c}

	return g
}

func (recv *GestureSwipe) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// GestureSingle upcasts to *GestureSingle
func (recv *GestureSwipe) GestureSingle() *GestureSingle {
	return GestureSingleNewFromC(unsafe.Pointer(recv.native))
}

// Gesture upcasts to *Gesture
func (recv *GestureSwipe) Gesture() *Gesture {
	return recv.GestureSingle().Gesture()
}

// EventController upcasts to *EventController
func (recv *GestureSwipe) EventController() *EventController {
	return recv.GestureSingle().EventController()
}

// Object upcasts to *Object
func (recv *GestureSwipe) Object() *gobject.Object {
	return recv.GestureSingle().Object()
}

// CastToWidget down casts any arbitary Object to GestureSwipe.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureSwipe.
func CastToGestureSwipe(object *gobject.Object) *GestureSwipe {
	return GestureSwipeNewFromC(object.ToC())
}

// GestureZoom is a wrapper around the C record GtkGestureZoom.
type GestureZoom struct {
	native *C.GtkGestureZoom
}

func GestureZoomNewFromC(u unsafe.Pointer) *GestureZoom {
	c := (*C.GtkGestureZoom)(u)
	if c == nil {
		return nil
	}

	g := &GestureZoom{native: c}

	return g
}

func (recv *GestureZoom) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Gesture upcasts to *Gesture
func (recv *GestureZoom) Gesture() *Gesture {
	return GestureNewFromC(unsafe.Pointer(recv.native))
}

// EventController upcasts to *EventController
func (recv *GestureZoom) EventController() *EventController {
	return recv.Gesture().EventController()
}

// Object upcasts to *Object
func (recv *GestureZoom) Object() *gobject.Object {
	return recv.Gesture().Object()
}

// CastToWidget down casts any arbitary Object to GestureZoom.
// Exercise care, as this is a potentially dangerous function if the Object is not a GestureZoom.
func CastToGestureZoom(object *gobject.Object) *GestureZoom {
	return GestureZoomNewFromC(object.ToC())
}

// Grid is a wrapper around the C record GtkGrid.
type Grid struct {
	native *C.GtkGrid
	// Private : container
	// Private : priv
}

func GridNewFromC(u unsafe.Pointer) *Grid {
	c := (*C.GtkGrid)(u)
	if c == nil {
		return nil
	}

	g := &Grid{native: c}

	return g
}

func (recv *Grid) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Grid) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Grid) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Grid) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Grid) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Grid.
// Exercise care, as this is a potentially dangerous function if the Object is not a Grid.
func CastToGrid(object *gobject.Object) *Grid {
	return GridNewFromC(object.ToC())
}

// GridNew is a wrapper around the C function gtk_grid_new.
func GridNew() *Grid {
	retC := C.gtk_grid_new()
	retGo := GridNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Attach is a wrapper around the C function gtk_grid_attach.
func (recv *Grid) Attach(child *Widget, left int32, top int32, width int32, height int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_left := (C.gint)(left)

	c_top := (C.gint)(top)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_grid_attach((*C.GtkGrid)(recv.native), c_child, c_left, c_top, c_width, c_height)

	return
}

// AttachNextTo is a wrapper around the C function gtk_grid_attach_next_to.
func (recv *Grid) AttachNextTo(child *Widget, sibling *Widget, side PositionType, width int32, height int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_sibling := (*C.GtkWidget)(sibling.ToC())

	c_side := (C.GtkPositionType)(side)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_grid_attach_next_to((*C.GtkGrid)(recv.native), c_child, c_sibling, c_side, c_width, c_height)

	return
}

// GetColumnHomogeneous is a wrapper around the C function gtk_grid_get_column_homogeneous.
func (recv *Grid) GetColumnHomogeneous() bool {
	retC := C.gtk_grid_get_column_homogeneous((*C.GtkGrid)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetColumnSpacing is a wrapper around the C function gtk_grid_get_column_spacing.
func (recv *Grid) GetColumnSpacing() uint32 {
	retC := C.gtk_grid_get_column_spacing((*C.GtkGrid)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetRowHomogeneous is a wrapper around the C function gtk_grid_get_row_homogeneous.
func (recv *Grid) GetRowHomogeneous() bool {
	retC := C.gtk_grid_get_row_homogeneous((*C.GtkGrid)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_grid_get_row_spacing.
func (recv *Grid) GetRowSpacing() uint32 {
	retC := C.gtk_grid_get_row_spacing((*C.GtkGrid)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// SetColumnHomogeneous is a wrapper around the C function gtk_grid_set_column_homogeneous.
func (recv *Grid) SetColumnHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_grid_set_column_homogeneous((*C.GtkGrid)(recv.native), c_homogeneous)

	return
}

// SetColumnSpacing is a wrapper around the C function gtk_grid_set_column_spacing.
func (recv *Grid) SetColumnSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_grid_set_column_spacing((*C.GtkGrid)(recv.native), c_spacing)

	return
}

// SetRowHomogeneous is a wrapper around the C function gtk_grid_set_row_homogeneous.
func (recv *Grid) SetRowHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_grid_set_row_homogeneous((*C.GtkGrid)(recv.native), c_homogeneous)

	return
}

// SetRowSpacing is a wrapper around the C function gtk_grid_set_row_spacing.
func (recv *Grid) SetRowSpacing(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_grid_set_row_spacing((*C.GtkGrid)(recv.native), c_spacing)

	return
}

// HBox is a wrapper around the C record GtkHBox.
type HBox struct {
	native *C.GtkHBox
	// box : record
}

func HBoxNewFromC(u unsafe.Pointer) *HBox {
	c := (*C.GtkHBox)(u)
	if c == nil {
		return nil
	}

	g := &HBox{native: c}

	return g
}

func (recv *HBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *HBox) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *HBox) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *HBox) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HBox) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to HBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a HBox.
func CastToHBox(object *gobject.Object) *HBox {
	return HBoxNewFromC(object.ToC())
}

// HBoxNew is a wrapper around the C function gtk_hbox_new.
func HBoxNew(homogeneous bool, spacing int32) *HBox {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_hbox_new(c_homogeneous, c_spacing)
	retGo := HBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HButtonBox is a wrapper around the C record GtkHButtonBox.
type HButtonBox struct {
	native *C.GtkHButtonBox
	// button_box : record
}

func HButtonBoxNewFromC(u unsafe.Pointer) *HButtonBox {
	c := (*C.GtkHButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &HButtonBox{native: c}

	return g
}

func (recv *HButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonBox upcasts to *ButtonBox
func (recv *HButtonBox) ButtonBox() *ButtonBox {
	return ButtonBoxNewFromC(unsafe.Pointer(recv.native))
}

// Box upcasts to *Box
func (recv *HButtonBox) Box() *Box {
	return recv.ButtonBox().Box()
}

// Container upcasts to *Container
func (recv *HButtonBox) Container() *Container {
	return recv.ButtonBox().Container()
}

// Widget upcasts to *Widget
func (recv *HButtonBox) Widget() *Widget {
	return recv.ButtonBox().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HButtonBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ButtonBox().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HButtonBox) Object() *gobject.Object {
	return recv.ButtonBox().Object()
}

// CastToWidget down casts any arbitary Object to HButtonBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a HButtonBox.
func CastToHButtonBox(object *gobject.Object) *HButtonBox {
	return HButtonBoxNewFromC(object.ToC())
}

// HButtonBoxNew is a wrapper around the C function gtk_hbutton_box_new.
func HButtonBoxNew() *HButtonBox {
	retC := C.gtk_hbutton_box_new()
	retGo := HButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HPaned is a wrapper around the C record GtkHPaned.
type HPaned struct {
	native *C.GtkHPaned
	// paned : record
}

func HPanedNewFromC(u unsafe.Pointer) *HPaned {
	c := (*C.GtkHPaned)(u)
	if c == nil {
		return nil
	}

	g := &HPaned{native: c}

	return g
}

func (recv *HPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Paned upcasts to *Paned
func (recv *HPaned) Paned() *Paned {
	return PanedNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *HPaned) Container() *Container {
	return recv.Paned().Container()
}

// Widget upcasts to *Widget
func (recv *HPaned) Widget() *Widget {
	return recv.Paned().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HPaned) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Paned().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HPaned) Object() *gobject.Object {
	return recv.Paned().Object()
}

// CastToWidget down casts any arbitary Object to HPaned.
// Exercise care, as this is a potentially dangerous function if the Object is not a HPaned.
func CastToHPaned(object *gobject.Object) *HPaned {
	return HPanedNewFromC(object.ToC())
}

// HPanedNew is a wrapper around the C function gtk_hpaned_new.
func HPanedNew() *HPaned {
	retC := C.gtk_hpaned_new()
	retGo := HPanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSV is a wrapper around the C record GtkHSV.
type HSV struct {
	native *C.GtkHSV
	// parent_instance : record
	// Private : priv
}

func HSVNewFromC(u unsafe.Pointer) *HSV {
	c := (*C.GtkHSV)(u)
	if c == nil {
		return nil
	}

	g := &HSV{native: c}

	return g
}

func (recv *HSV) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *HSV) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HSV) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HSV) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to HSV.
// Exercise care, as this is a potentially dangerous function if the Object is not a HSV.
func CastToHSV(object *gobject.Object) *HSV {
	return HSVNewFromC(object.ToC())
}

var signalChangedId int
var signalChangedMap = make(map[int]HSVSignalChangedCallback)
var signalChangedLock sync.Mutex

// HSVSignalChangedCallback is a callback function for a 'changed' signal emitted from a HSV.
type HSVSignalChangedCallback func()

func HSV_changedHandler() {}

var signalMoveId int
var signalMoveMap = make(map[int]HSVSignalMoveCallback)
var signalMoveLock sync.Mutex

// HSVSignalMoveCallback is a callback function for a 'move' signal emitted from a HSV.
type HSVSignalMoveCallback func(object DirectionType)

func HSV_moveHandler() {}

// HScale is a wrapper around the C record GtkHScale.
type HScale struct {
	native *C.GtkHScale
	// scale : record
}

func HScaleNewFromC(u unsafe.Pointer) *HScale {
	c := (*C.GtkHScale)(u)
	if c == nil {
		return nil
	}

	g := &HScale{native: c}

	return g
}

func (recv *HScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scale upcasts to *Scale
func (recv *HScale) Scale() *Scale {
	return ScaleNewFromC(unsafe.Pointer(recv.native))
}

// Range upcasts to *Range
func (recv *HScale) Range() *Range {
	return recv.Scale().Range()
}

// Widget upcasts to *Widget
func (recv *HScale) Widget() *Widget {
	return recv.Scale().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HScale) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Scale().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HScale) Object() *gobject.Object {
	return recv.Scale().Object()
}

// CastToWidget down casts any arbitary Object to HScale.
// Exercise care, as this is a potentially dangerous function if the Object is not a HScale.
func CastToHScale(object *gobject.Object) *HScale {
	return HScaleNewFromC(object.ToC())
}

// HScaleNew is a wrapper around the C function gtk_hscale_new.
func HScaleNew(adjustment *Adjustment) *HScale {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_hscale_new(c_adjustment)
	retGo := HScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HScaleNewWithRange is a wrapper around the C function gtk_hscale_new_with_range.
func HScaleNewWithRange(min float64, max float64, step float64) *HScale {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_hscale_new_with_range(c_min, c_max, c_step)
	retGo := HScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HScrollbar is a wrapper around the C record GtkHScrollbar.
type HScrollbar struct {
	native *C.GtkHScrollbar
	// scrollbar : record
}

func HScrollbarNewFromC(u unsafe.Pointer) *HScrollbar {
	c := (*C.GtkHScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &HScrollbar{native: c}

	return g
}

func (recv *HScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scrollbar upcasts to *Scrollbar
func (recv *HScrollbar) Scrollbar() *Scrollbar {
	return ScrollbarNewFromC(unsafe.Pointer(recv.native))
}

// Range upcasts to *Range
func (recv *HScrollbar) Range() *Range {
	return recv.Scrollbar().Range()
}

// Widget upcasts to *Widget
func (recv *HScrollbar) Widget() *Widget {
	return recv.Scrollbar().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HScrollbar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Scrollbar().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HScrollbar) Object() *gobject.Object {
	return recv.Scrollbar().Object()
}

// CastToWidget down casts any arbitary Object to HScrollbar.
// Exercise care, as this is a potentially dangerous function if the Object is not a HScrollbar.
func CastToHScrollbar(object *gobject.Object) *HScrollbar {
	return HScrollbarNewFromC(object.ToC())
}

// HScrollbarNew is a wrapper around the C function gtk_hscrollbar_new.
func HScrollbarNew(adjustment *Adjustment) *HScrollbar {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_hscrollbar_new(c_adjustment)
	retGo := HScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HSeparator is a wrapper around the C record GtkHSeparator.
type HSeparator struct {
	native *C.GtkHSeparator
	// separator : record
}

func HSeparatorNewFromC(u unsafe.Pointer) *HSeparator {
	c := (*C.GtkHSeparator)(u)
	if c == nil {
		return nil
	}

	g := &HSeparator{native: c}

	return g
}

func (recv *HSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Separator upcasts to *Separator
func (recv *HSeparator) Separator() *Separator {
	return SeparatorNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *HSeparator) Widget() *Widget {
	return recv.Separator().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HSeparator) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Separator().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HSeparator) Object() *gobject.Object {
	return recv.Separator().Object()
}

// CastToWidget down casts any arbitary Object to HSeparator.
// Exercise care, as this is a potentially dangerous function if the Object is not a HSeparator.
func CastToHSeparator(object *gobject.Object) *HSeparator {
	return HSeparatorNewFromC(object.ToC())
}

// HSeparatorNew is a wrapper around the C function gtk_hseparator_new.
func HSeparatorNew() *HSeparator {
	retC := C.gtk_hseparator_new()
	retGo := HSeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HandleBox is a wrapper around the C record GtkHandleBox.
type HandleBox struct {
	native *C.GtkHandleBox
	// bin : record
	// Private : priv
}

func HandleBoxNewFromC(u unsafe.Pointer) *HandleBox {
	c := (*C.GtkHandleBox)(u)
	if c == nil {
		return nil
	}

	g := &HandleBox{native: c}

	return g
}

func (recv *HandleBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *HandleBox) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *HandleBox) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *HandleBox) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HandleBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HandleBox) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to HandleBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a HandleBox.
func CastToHandleBox(object *gobject.Object) *HandleBox {
	return HandleBoxNewFromC(object.ToC())
}

var signalChildAttachedId int
var signalChildAttachedMap = make(map[int]HandleBoxSignalChildAttachedCallback)
var signalChildAttachedLock sync.Mutex

// HandleBoxSignalChildAttachedCallback is a callback function for a 'child-attached' signal emitted from a HandleBox.
type HandleBoxSignalChildAttachedCallback func(widget *Widget)

func HandleBox_childAttachedHandler() {}

var signalChildDetachedId int
var signalChildDetachedMap = make(map[int]HandleBoxSignalChildDetachedCallback)
var signalChildDetachedLock sync.Mutex

// HandleBoxSignalChildDetachedCallback is a callback function for a 'child-detached' signal emitted from a HandleBox.
type HandleBoxSignalChildDetachedCallback func(widget *Widget)

func HandleBox_childDetachedHandler() {}

// HandleBoxNew is a wrapper around the C function gtk_handle_box_new.
func HandleBoxNew() *HandleBox {
	retC := C.gtk_handle_box_new()
	retGo := HandleBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHandlePosition is a wrapper around the C function gtk_handle_box_get_handle_position.
func (recv *HandleBox) GetHandlePosition() PositionType {
	retC := C.gtk_handle_box_get_handle_position((*C.GtkHandleBox)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_handle_box_get_shadow_type.
func (recv *HandleBox) GetShadowType() ShadowType {
	retC := C.gtk_handle_box_get_shadow_type((*C.GtkHandleBox)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetSnapEdge is a wrapper around the C function gtk_handle_box_get_snap_edge.
func (recv *HandleBox) GetSnapEdge() PositionType {
	retC := C.gtk_handle_box_get_snap_edge((*C.GtkHandleBox)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// SetHandlePosition is a wrapper around the C function gtk_handle_box_set_handle_position.
func (recv *HandleBox) SetHandlePosition(position PositionType) {
	c_position := (C.GtkPositionType)(position)

	C.gtk_handle_box_set_handle_position((*C.GtkHandleBox)(recv.native), c_position)

	return
}

// SetShadowType is a wrapper around the C function gtk_handle_box_set_shadow_type.
func (recv *HandleBox) SetShadowType(type_ ShadowType) {
	c_type := (C.GtkShadowType)(type_)

	C.gtk_handle_box_set_shadow_type((*C.GtkHandleBox)(recv.native), c_type)

	return
}

// SetSnapEdge is a wrapper around the C function gtk_handle_box_set_snap_edge.
func (recv *HandleBox) SetSnapEdge(edge PositionType) {
	c_edge := (C.GtkPositionType)(edge)

	C.gtk_handle_box_set_snap_edge((*C.GtkHandleBox)(recv.native), c_edge)

	return
}

// HeaderBar is a wrapper around the C record GtkHeaderBar.
type HeaderBar struct {
	native *C.GtkHeaderBar
	// container : record
}

func HeaderBarNewFromC(u unsafe.Pointer) *HeaderBar {
	c := (*C.GtkHeaderBar)(u)
	if c == nil {
		return nil
	}

	g := &HeaderBar{native: c}

	return g
}

func (recv *HeaderBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *HeaderBar) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *HeaderBar) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *HeaderBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *HeaderBar) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to HeaderBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a HeaderBar.
func CastToHeaderBar(object *gobject.Object) *HeaderBar {
	return HeaderBarNewFromC(object.ToC())
}

// IMContext is a wrapper around the C record GtkIMContext.
type IMContext struct {
	native *C.GtkIMContext
	// parent_instance : record
}

func IMContextNewFromC(u unsafe.Pointer) *IMContext {
	c := (*C.GtkIMContext)(u)
	if c == nil {
		return nil
	}

	g := &IMContext{native: c}

	return g
}

func (recv *IMContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *IMContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to IMContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a IMContext.
func CastToIMContext(object *gobject.Object) *IMContext {
	return IMContextNewFromC(object.ToC())
}

var signalCommitId int
var signalCommitMap = make(map[int]IMContextSignalCommitCallback)
var signalCommitLock sync.Mutex

// IMContextSignalCommitCallback is a callback function for a 'commit' signal emitted from a IMContext.
type IMContextSignalCommitCallback func(str string)

func IMContext_commitHandler() {}

var signalDeleteSurroundingId int
var signalDeleteSurroundingMap = make(map[int]IMContextSignalDeleteSurroundingCallback)
var signalDeleteSurroundingLock sync.Mutex

// IMContextSignalDeleteSurroundingCallback is a callback function for a 'delete-surrounding' signal emitted from a IMContext.
type IMContextSignalDeleteSurroundingCallback func(offset int32, nChars int32) bool

func IMContext_deleteSurroundingHandler() {}

var signalPreeditChangedId int
var signalPreeditChangedMap = make(map[int]IMContextSignalPreeditChangedCallback)
var signalPreeditChangedLock sync.Mutex

// IMContextSignalPreeditChangedCallback is a callback function for a 'preedit-changed' signal emitted from a IMContext.
type IMContextSignalPreeditChangedCallback func()

func IMContext_preeditChangedHandler() {}

var signalPreeditEndId int
var signalPreeditEndMap = make(map[int]IMContextSignalPreeditEndCallback)
var signalPreeditEndLock sync.Mutex

// IMContextSignalPreeditEndCallback is a callback function for a 'preedit-end' signal emitted from a IMContext.
type IMContextSignalPreeditEndCallback func()

func IMContext_preeditEndHandler() {}

var signalPreeditStartId int
var signalPreeditStartMap = make(map[int]IMContextSignalPreeditStartCallback)
var signalPreeditStartLock sync.Mutex

// IMContextSignalPreeditStartCallback is a callback function for a 'preedit-start' signal emitted from a IMContext.
type IMContextSignalPreeditStartCallback func()

func IMContext_preeditStartHandler() {}

var signalRetrieveSurroundingId int
var signalRetrieveSurroundingMap = make(map[int]IMContextSignalRetrieveSurroundingCallback)
var signalRetrieveSurroundingLock sync.Mutex

// IMContextSignalRetrieveSurroundingCallback is a callback function for a 'retrieve-surrounding' signal emitted from a IMContext.
type IMContextSignalRetrieveSurroundingCallback func() bool

func IMContext_retrieveSurroundingHandler() {}

// DeleteSurrounding is a wrapper around the C function gtk_im_context_delete_surrounding.
func (recv *IMContext) DeleteSurrounding(offset int32, nChars int32) bool {
	c_offset := (C.gint)(offset)

	c_n_chars := (C.gint)(nChars)

	retC := C.gtk_im_context_delete_surrounding((*C.GtkIMContext)(recv.native), c_offset, c_n_chars)
	retGo := retC == C.TRUE

	return retGo
}

// FilterKeypress is a wrapper around the C function gtk_im_context_filter_keypress.
func (recv *IMContext) FilterKeypress(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(event.ToC())

	retC := C.gtk_im_context_filter_keypress((*C.GtkIMContext)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// FocusIn is a wrapper around the C function gtk_im_context_focus_in.
func (recv *IMContext) FocusIn() {
	C.gtk_im_context_focus_in((*C.GtkIMContext)(recv.native))

	return
}

// FocusOut is a wrapper around the C function gtk_im_context_focus_out.
func (recv *IMContext) FocusOut() {
	C.gtk_im_context_focus_out((*C.GtkIMContext)(recv.native))

	return
}

// Unsupported : gtk_im_context_get_preedit_string : unsupported parameter attrs : record with indirection level of 2

// GetSurrounding is a wrapper around the C function gtk_im_context_get_surrounding.
func (recv *IMContext) GetSurrounding() (bool, string, int32) {
	var c_text *C.gchar

	var c_cursor_index C.gint

	retC := C.gtk_im_context_get_surrounding((*C.GtkIMContext)(recv.native), &c_text, &c_cursor_index)
	retGo := retC == C.TRUE

	text := C.GoString(c_text)
	defer C.free(unsafe.Pointer(c_text))

	cursorIndex := (int32)(c_cursor_index)

	return retGo, text, cursorIndex
}

// Reset is a wrapper around the C function gtk_im_context_reset.
func (recv *IMContext) Reset() {
	C.gtk_im_context_reset((*C.GtkIMContext)(recv.native))

	return
}

// SetClientWindow is a wrapper around the C function gtk_im_context_set_client_window.
func (recv *IMContext) SetClientWindow(window *gdk.Window) {
	c_window := (*C.GdkWindow)(window.ToC())

	C.gtk_im_context_set_client_window((*C.GtkIMContext)(recv.native), c_window)

	return
}

// Unsupported : gtk_im_context_set_cursor_location : unsupported parameter area : Blacklisted record : GdkRectangle

// SetSurrounding is a wrapper around the C function gtk_im_context_set_surrounding.
func (recv *IMContext) SetSurrounding(text string, len int32, cursorIndex int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	c_cursor_index := (C.gint)(cursorIndex)

	C.gtk_im_context_set_surrounding((*C.GtkIMContext)(recv.native), c_text, c_len, c_cursor_index)

	return
}

// SetUsePreedit is a wrapper around the C function gtk_im_context_set_use_preedit.
func (recv *IMContext) SetUsePreedit(usePreedit bool) {
	c_use_preedit :=
		boolToGboolean(usePreedit)

	C.gtk_im_context_set_use_preedit((*C.GtkIMContext)(recv.native), c_use_preedit)

	return
}

// IMContextSimple is a wrapper around the C record GtkIMContextSimple.
type IMContextSimple struct {
	native *C.GtkIMContextSimple
	// object : record
	// Private : priv
}

func IMContextSimpleNewFromC(u unsafe.Pointer) *IMContextSimple {
	c := (*C.GtkIMContextSimple)(u)
	if c == nil {
		return nil
	}

	g := &IMContextSimple{native: c}

	return g
}

func (recv *IMContextSimple) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContext upcasts to *IMContext
func (recv *IMContextSimple) IMContext() *IMContext {
	return IMContextNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *IMContextSimple) Object() *gobject.Object {
	return recv.IMContext().Object()
}

// CastToWidget down casts any arbitary Object to IMContextSimple.
// Exercise care, as this is a potentially dangerous function if the Object is not a IMContextSimple.
func CastToIMContextSimple(object *gobject.Object) *IMContextSimple {
	return IMContextSimpleNewFromC(object.ToC())
}

// IMContextSimpleNew is a wrapper around the C function gtk_im_context_simple_new.
func IMContextSimpleNew() *IMContextSimple {
	retC := C.gtk_im_context_simple_new()
	retGo := IMContextSimpleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Blacklisted : gtk_im_context_simple_add_compose_file

// Unsupported : gtk_im_context_simple_add_table : unsupported parameter data : no param type

// IMMulticontext is a wrapper around the C record GtkIMMulticontext.
type IMMulticontext struct {
	native *C.GtkIMMulticontext
	// object : record
	// Private : priv
}

func IMMulticontextNewFromC(u unsafe.Pointer) *IMMulticontext {
	c := (*C.GtkIMMulticontext)(u)
	if c == nil {
		return nil
	}

	g := &IMMulticontext{native: c}

	return g
}

func (recv *IMMulticontext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// IMContext upcasts to *IMContext
func (recv *IMMulticontext) IMContext() *IMContext {
	return IMContextNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *IMMulticontext) Object() *gobject.Object {
	return recv.IMContext().Object()
}

// CastToWidget down casts any arbitary Object to IMMulticontext.
// Exercise care, as this is a potentially dangerous function if the Object is not a IMMulticontext.
func CastToIMMulticontext(object *gobject.Object) *IMMulticontext {
	return IMMulticontextNewFromC(object.ToC())
}

// IMMulticontextNew is a wrapper around the C function gtk_im_multicontext_new.
func IMMulticontextNew() *IMMulticontext {
	retC := C.gtk_im_multicontext_new()
	retGo := IMMulticontextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendMenuitems is a wrapper around the C function gtk_im_multicontext_append_menuitems.
func (recv *IMMulticontext) AppendMenuitems(menushell *MenuShell) {
	c_menushell := (*C.GtkMenuShell)(menushell.ToC())

	C.gtk_im_multicontext_append_menuitems((*C.GtkIMMulticontext)(recv.native), c_menushell)

	return
}

// IconFactory is a wrapper around the C record GtkIconFactory.
type IconFactory struct {
	native *C.GtkIconFactory
	// parent_instance : record
	// Private : priv
}

func IconFactoryNewFromC(u unsafe.Pointer) *IconFactory {
	c := (*C.GtkIconFactory)(u)
	if c == nil {
		return nil
	}

	g := &IconFactory{native: c}

	return g
}

func (recv *IconFactory) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *IconFactory) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to IconFactory.
// Exercise care, as this is a potentially dangerous function if the Object is not a IconFactory.
func CastToIconFactory(object *gobject.Object) *IconFactory {
	return IconFactoryNewFromC(object.ToC())
}

// IconFactoryNew is a wrapper around the C function gtk_icon_factory_new.
func IconFactoryNew() *IconFactory {
	retC := C.gtk_icon_factory_new()
	retGo := IconFactoryNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function gtk_icon_factory_add.
func (recv *IconFactory) Add(stockId string, iconSet *IconSet) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_icon_set := (*C.GtkIconSet)(iconSet.ToC())

	C.gtk_icon_factory_add((*C.GtkIconFactory)(recv.native), c_stock_id, c_icon_set)

	return
}

// AddDefault is a wrapper around the C function gtk_icon_factory_add_default.
func (recv *IconFactory) AddDefault() {
	C.gtk_icon_factory_add_default((*C.GtkIconFactory)(recv.native))

	return
}

// Lookup is a wrapper around the C function gtk_icon_factory_lookup.
func (recv *IconFactory) Lookup(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_icon_factory_lookup((*C.GtkIconFactory)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveDefault is a wrapper around the C function gtk_icon_factory_remove_default.
func (recv *IconFactory) RemoveDefault() {
	C.gtk_icon_factory_remove_default((*C.GtkIconFactory)(recv.native))

	return
}

// IconInfo is a wrapper around the C record GtkIconInfo.
type IconInfo struct {
	native *C.GtkIconInfo
}

func IconInfoNewFromC(u unsafe.Pointer) *IconInfo {
	c := (*C.GtkIconInfo)(u)
	if c == nil {
		return nil
	}

	g := &IconInfo{native: c}

	return g
}

func (recv *IconInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *IconInfo) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to IconInfo.
// Exercise care, as this is a potentially dangerous function if the Object is not a IconInfo.
func CastToIconInfo(object *gobject.Object) *IconInfo {
	return IconInfoNewFromC(object.ToC())
}

// IconTheme is a wrapper around the C record GtkIconTheme.
type IconTheme struct {
	native *C.GtkIconTheme
	// Private : parent_instance
	// Private : priv
}

func IconThemeNewFromC(u unsafe.Pointer) *IconTheme {
	c := (*C.GtkIconTheme)(u)
	if c == nil {
		return nil
	}

	g := &IconTheme{native: c}

	return g
}

func (recv *IconTheme) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *IconTheme) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to IconTheme.
// Exercise care, as this is a potentially dangerous function if the Object is not a IconTheme.
func CastToIconTheme(object *gobject.Object) *IconTheme {
	return IconThemeNewFromC(object.ToC())
}

var signalChangedId int
var signalChangedMap = make(map[int]IconThemeSignalChangedCallback)
var signalChangedLock sync.Mutex

// IconThemeSignalChangedCallback is a callback function for a 'changed' signal emitted from a IconTheme.
type IconThemeSignalChangedCallback func()

func IconTheme_changedHandler() {}

// IconView is a wrapper around the C record GtkIconView.
type IconView struct {
	native *C.GtkIconView
	// parent : record
	// Private : priv
}

func IconViewNewFromC(u unsafe.Pointer) *IconView {
	c := (*C.GtkIconView)(u)
	if c == nil {
		return nil
	}

	g := &IconView{native: c}

	return g
}

func (recv *IconView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *IconView) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *IconView) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *IconView) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *IconView) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to IconView.
// Exercise care, as this is a potentially dangerous function if the Object is not a IconView.
func CastToIconView(object *gobject.Object) *IconView {
	return IconViewNewFromC(object.ToC())
}

var signalActivateCursorItemId int
var signalActivateCursorItemMap = make(map[int]IconViewSignalActivateCursorItemCallback)
var signalActivateCursorItemLock sync.Mutex

// IconViewSignalActivateCursorItemCallback is a callback function for a 'activate-cursor-item' signal emitted from a IconView.
type IconViewSignalActivateCursorItemCallback func() bool

func IconView_activateCursorItemHandler() {}

var signalItemActivatedId int
var signalItemActivatedMap = make(map[int]IconViewSignalItemActivatedCallback)
var signalItemActivatedLock sync.Mutex

// IconViewSignalItemActivatedCallback is a callback function for a 'item-activated' signal emitted from a IconView.
type IconViewSignalItemActivatedCallback func(path *TreePath)

func IconView_itemActivatedHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]IconViewSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// IconViewSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a IconView.
type IconViewSignalMoveCursorCallback func(step MovementStep, count int32) bool

func IconView_moveCursorHandler() {}

var signalSelectAllId int
var signalSelectAllMap = make(map[int]IconViewSignalSelectAllCallback)
var signalSelectAllLock sync.Mutex

// IconViewSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a IconView.
type IconViewSignalSelectAllCallback func()

func IconView_selectAllHandler() {}

var signalSelectCursorItemId int
var signalSelectCursorItemMap = make(map[int]IconViewSignalSelectCursorItemCallback)
var signalSelectCursorItemLock sync.Mutex

// IconViewSignalSelectCursorItemCallback is a callback function for a 'select-cursor-item' signal emitted from a IconView.
type IconViewSignalSelectCursorItemCallback func()

func IconView_selectCursorItemHandler() {}

var signalSelectionChangedId int
var signalSelectionChangedMap = make(map[int]IconViewSignalSelectionChangedCallback)
var signalSelectionChangedLock sync.Mutex

// IconViewSignalSelectionChangedCallback is a callback function for a 'selection-changed' signal emitted from a IconView.
type IconViewSignalSelectionChangedCallback func()

func IconView_selectionChangedHandler() {}

var signalToggleCursorItemId int
var signalToggleCursorItemMap = make(map[int]IconViewSignalToggleCursorItemCallback)
var signalToggleCursorItemLock sync.Mutex

// IconViewSignalToggleCursorItemCallback is a callback function for a 'toggle-cursor-item' signal emitted from a IconView.
type IconViewSignalToggleCursorItemCallback func()

func IconView_toggleCursorItemHandler() {}

var signalUnselectAllId int
var signalUnselectAllMap = make(map[int]IconViewSignalUnselectAllCallback)
var signalUnselectAllLock sync.Mutex

// IconViewSignalUnselectAllCallback is a callback function for a 'unselect-all' signal emitted from a IconView.
type IconViewSignalUnselectAllCallback func()

func IconView_unselectAllHandler() {}

// Unsupported : gtk_icon_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// IconViewAccessible is a wrapper around the C record GtkIconViewAccessible.
type IconViewAccessible struct {
	native *C.GtkIconViewAccessible
	// parent : record
	// priv : record
}

func IconViewAccessibleNewFromC(u unsafe.Pointer) *IconViewAccessible {
	c := (*C.GtkIconViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &IconViewAccessible{native: c}

	return g
}

func (recv *IconViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *IconViewAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *IconViewAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *IconViewAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *IconViewAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to IconViewAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a IconViewAccessible.
func CastToIconViewAccessible(object *gobject.Object) *IconViewAccessible {
	return IconViewAccessibleNewFromC(object.ToC())
}

// Image is a wrapper around the C record GtkImage.
type Image struct {
	native *C.GtkImage
	// misc : record
	// Private : priv
}

func ImageNewFromC(u unsafe.Pointer) *Image {
	c := (*C.GtkImage)(u)
	if c == nil {
		return nil
	}

	g := &Image{native: c}

	return g
}

func (recv *Image) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Misc upcasts to *Misc
func (recv *Image) Misc() *Misc {
	return MiscNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Image) Widget() *Widget {
	return recv.Misc().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Image) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Misc().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Image) Object() *gobject.Object {
	return recv.Misc().Object()
}

// CastToWidget down casts any arbitary Object to Image.
// Exercise care, as this is a potentially dangerous function if the Object is not a Image.
func CastToImage(object *gobject.Object) *Image {
	return ImageNewFromC(object.ToC())
}

// ImageNew is a wrapper around the C function gtk_image_new.
func ImageNew() *Image {
	retC := C.gtk_image_new()
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromAnimation is a wrapper around the C function gtk_image_new_from_animation.
func ImageNewFromAnimation(animation *gdkpixbuf.PixbufAnimation) *Image {
	c_animation := (*C.GdkPixbufAnimation)(animation.ToC())

	retC := C.gtk_image_new_from_animation(c_animation)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageNewFromFile is a wrapper around the C function gtk_image_new_from_file.
func ImageNewFromFile(filename string) *Image {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_image_new_from_file(c_filename)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Unsupported : gtk_image_new_from_icon_name : unsupported parameter size : no type generator for gint, GtkIconSize

// Unsupported : gtk_image_new_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// ImageNewFromPixbuf is a wrapper around the C function gtk_image_new_from_pixbuf.
func ImageNewFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	retC := C.gtk_image_new_from_pixbuf(c_pixbuf)
	retGo := ImageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_new_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// GetAnimation is a wrapper around the C function gtk_image_get_animation.
func (recv *Image) GetAnimation() *gdkpixbuf.PixbufAnimation {
	retC := C.gtk_image_get_animation((*C.GtkImage)(recv.native))
	retGo := gdkpixbuf.PixbufAnimationNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_icon_set : unsupported parameter icon_set : record with indirection level of 2

// GetPixbuf is a wrapper around the C function gtk_image_get_pixbuf.
func (recv *Image) GetPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_image_get_pixbuf((*C.GtkImage)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_image_get_stock : unsupported parameter size : no type generator for gint, GtkIconSize*

// GetStorageType is a wrapper around the C function gtk_image_get_storage_type.
func (recv *Image) GetStorageType() ImageType {
	retC := C.gtk_image_get_storage_type((*C.GtkImage)(recv.native))
	retGo := (ImageType)(retC)

	return retGo
}

// SetFromAnimation is a wrapper around the C function gtk_image_set_from_animation.
func (recv *Image) SetFromAnimation(animation *gdkpixbuf.PixbufAnimation) {
	c_animation := (*C.GdkPixbufAnimation)(animation.ToC())

	C.gtk_image_set_from_animation((*C.GtkImage)(recv.native), c_animation)

	return
}

// SetFromFile is a wrapper around the C function gtk_image_set_from_file.
func (recv *Image) SetFromFile(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_image_set_from_file((*C.GtkImage)(recv.native), c_filename)

	return
}

// Unsupported : gtk_image_set_from_icon_set : unsupported parameter size : no type generator for gint, GtkIconSize

// SetFromPixbuf is a wrapper around the C function gtk_image_set_from_pixbuf.
func (recv *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_image_set_from_pixbuf((*C.GtkImage)(recv.native), c_pixbuf)

	return
}

// SetFromResource is a wrapper around the C function gtk_image_set_from_resource.
func (recv *Image) SetFromResource(resourcePath string) {
	c_resource_path := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(c_resource_path))

	C.gtk_image_set_from_resource((*C.GtkImage)(recv.native), c_resource_path)

	return
}

// Unsupported : gtk_image_set_from_stock : unsupported parameter size : no type generator for gint, GtkIconSize

// ImageAccessible is a wrapper around the C record GtkImageAccessible.
type ImageAccessible struct {
	native *C.GtkImageAccessible
	// parent : record
	// priv : record
}

func ImageAccessibleNewFromC(u unsafe.Pointer) *ImageAccessible {
	c := (*C.GtkImageAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ImageAccessible{native: c}

	return g
}

func (recv *ImageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ImageAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *ImageAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ImageAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ImageAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ImageAccessible.
func CastToImageAccessible(object *gobject.Object) *ImageAccessible {
	return ImageAccessibleNewFromC(object.ToC())
}

// ImageCellAccessible is a wrapper around the C record GtkImageCellAccessible.
type ImageCellAccessible struct {
	native *C.GtkImageCellAccessible
	// parent : record
	// priv : record
}

func ImageCellAccessibleNewFromC(u unsafe.Pointer) *ImageCellAccessible {
	c := (*C.GtkImageCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ImageCellAccessible{native: c}

	return g
}

func (recv *ImageCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessible upcasts to *RendererCellAccessible
func (recv *ImageCellAccessible) RendererCellAccessible() *RendererCellAccessible {
	return RendererCellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// CellAccessible upcasts to *CellAccessible
func (recv *ImageCellAccessible) CellAccessible() *CellAccessible {
	return recv.RendererCellAccessible().CellAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ImageCellAccessible) Accessible() *Accessible {
	return recv.RendererCellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ImageCellAccessible) Object() *atk.Object {
	return recv.RendererCellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ImageCellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ImageCellAccessible.
func CastToImageCellAccessible(object *gobject.Object) *ImageCellAccessible {
	return ImageCellAccessibleNewFromC(object.ToC())
}

// ImageMenuItem is a wrapper around the C record GtkImageMenuItem.
type ImageMenuItem struct {
	native *C.GtkImageMenuItem
	// menu_item : record
	// Private : priv
}

func ImageMenuItemNewFromC(u unsafe.Pointer) *ImageMenuItem {
	c := (*C.GtkImageMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &ImageMenuItem{native: c}

	return g
}

func (recv *ImageMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem upcasts to *MenuItem
func (recv *ImageMenuItem) MenuItem() *MenuItem {
	return MenuItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ImageMenuItem) Bin() *Bin {
	return recv.MenuItem().Bin()
}

// Container upcasts to *Container
func (recv *ImageMenuItem) Container() *Container {
	return recv.MenuItem().Container()
}

// Widget upcasts to *Widget
func (recv *ImageMenuItem) Widget() *Widget {
	return recv.MenuItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ImageMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ImageMenuItem) Object() *gobject.Object {
	return recv.MenuItem().Object()
}

// CastToWidget down casts any arbitary Object to ImageMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a ImageMenuItem.
func CastToImageMenuItem(object *gobject.Object) *ImageMenuItem {
	return ImageMenuItemNewFromC(object.ToC())
}

// ImageMenuItemNew is a wrapper around the C function gtk_image_menu_item_new.
func ImageMenuItemNew() *ImageMenuItem {
	retC := C.gtk_image_menu_item_new()
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewFromStock is a wrapper around the C function gtk_image_menu_item_new_from_stock.
func ImageMenuItemNewFromStock(stockId string, accelGroup *AccelGroup) *ImageMenuItem {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	retC := C.gtk_image_menu_item_new_from_stock(c_stock_id, c_accel_group)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithLabel is a wrapper around the C function gtk_image_menu_item_new_with_label.
func ImageMenuItemNewWithLabel(label string) *ImageMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_label(c_label)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ImageMenuItemNewWithMnemonic is a wrapper around the C function gtk_image_menu_item_new_with_mnemonic.
func ImageMenuItemNewWithMnemonic(label string) *ImageMenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_image_menu_item_new_with_mnemonic(c_label)
	retGo := ImageMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetImage is a wrapper around the C function gtk_image_menu_item_get_image.
func (recv *ImageMenuItem) GetImage() *Widget {
	retC := C.gtk_image_menu_item_get_image((*C.GtkImageMenuItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetImage is a wrapper around the C function gtk_image_menu_item_set_image.
func (recv *ImageMenuItem) SetImage(image *Widget) {
	c_image := (*C.GtkWidget)(image.ToC())

	C.gtk_image_menu_item_set_image((*C.GtkImageMenuItem)(recv.native), c_image)

	return
}

// InfoBar is a wrapper around the C record GtkInfoBar.
type InfoBar struct {
	native *C.GtkInfoBar
	// parent : record
	// Private : priv
}

func InfoBarNewFromC(u unsafe.Pointer) *InfoBar {
	c := (*C.GtkInfoBar)(u)
	if c == nil {
		return nil
	}

	g := &InfoBar{native: c}

	return g
}

func (recv *InfoBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *InfoBar) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *InfoBar) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *InfoBar) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *InfoBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *InfoBar) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to InfoBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a InfoBar.
func CastToInfoBar(object *gobject.Object) *InfoBar {
	return InfoBarNewFromC(object.ToC())
}

// Unsupported : gtk_info_bar_new_with_buttons : unsupported parameter ... : varargs

// Invisible is a wrapper around the C record GtkInvisible.
type Invisible struct {
	native *C.GtkInvisible
	// widget : record
	// Private : priv
}

func InvisibleNewFromC(u unsafe.Pointer) *Invisible {
	c := (*C.GtkInvisible)(u)
	if c == nil {
		return nil
	}

	g := &Invisible{native: c}

	return g
}

func (recv *Invisible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Invisible) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Invisible) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Invisible) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Invisible.
// Exercise care, as this is a potentially dangerous function if the Object is not a Invisible.
func CastToInvisible(object *gobject.Object) *Invisible {
	return InvisibleNewFromC(object.ToC())
}

// InvisibleNew is a wrapper around the C function gtk_invisible_new.
func InvisibleNew() *Invisible {
	retC := C.gtk_invisible_new()
	retGo := InvisibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Label is a wrapper around the C record GtkLabel.
type Label struct {
	native *C.GtkLabel
	// misc : record
	// Private : priv
}

func LabelNewFromC(u unsafe.Pointer) *Label {
	c := (*C.GtkLabel)(u)
	if c == nil {
		return nil
	}

	g := &Label{native: c}

	return g
}

func (recv *Label) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Misc upcasts to *Misc
func (recv *Label) Misc() *Misc {
	return MiscNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Label) Widget() *Widget {
	return recv.Misc().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Label) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Misc().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Label) Object() *gobject.Object {
	return recv.Misc().Object()
}

// CastToWidget down casts any arbitary Object to Label.
// Exercise care, as this is a potentially dangerous function if the Object is not a Label.
func CastToLabel(object *gobject.Object) *Label {
	return LabelNewFromC(object.ToC())
}

var signalCopyClipboardId int
var signalCopyClipboardMap = make(map[int]LabelSignalCopyClipboardCallback)
var signalCopyClipboardLock sync.Mutex

// LabelSignalCopyClipboardCallback is a callback function for a 'copy-clipboard' signal emitted from a Label.
type LabelSignalCopyClipboardCallback func()

func Label_copyClipboardHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]LabelSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// LabelSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a Label.
type LabelSignalMoveCursorCallback func(step MovementStep, count int32, extendSelection bool)

func Label_moveCursorHandler() {}

var signalPopulatePopupId int
var signalPopulatePopupMap = make(map[int]LabelSignalPopulatePopupCallback)
var signalPopulatePopupLock sync.Mutex

// LabelSignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a Label.
type LabelSignalPopulatePopupCallback func(menu *Menu)

func Label_populatePopupHandler() {}

// LabelNew is a wrapper around the C function gtk_label_new.
func LabelNew(str string) *Label {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new(c_str)
	retGo := LabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LabelNewWithMnemonic is a wrapper around the C function gtk_label_new_with_mnemonic.
func LabelNewWithMnemonic(str string) *Label {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	retC := C.gtk_label_new_with_mnemonic(c_str)
	retGo := LabelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttributes is a wrapper around the C function gtk_label_get_attributes.
func (recv *Label) GetAttributes() *pango.AttrList {
	retC := C.gtk_label_get_attributes((*C.GtkLabel)(recv.native))
	retGo := pango.AttrListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetJustify is a wrapper around the C function gtk_label_get_justify.
func (recv *Label) GetJustify() Justification {
	retC := C.gtk_label_get_justify((*C.GtkLabel)(recv.native))
	retGo := (Justification)(retC)

	return retGo
}

// GetLabel is a wrapper around the C function gtk_label_get_label.
func (recv *Label) GetLabel() string {
	retC := C.gtk_label_get_label((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLayout is a wrapper around the C function gtk_label_get_layout.
func (recv *Label) GetLayout() *pango.Layout {
	retC := C.gtk_label_get_layout((*C.GtkLabel)(recv.native))
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLayoutOffsets is a wrapper around the C function gtk_label_get_layout_offsets.
func (recv *Label) GetLayoutOffsets() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_label_get_layout_offsets((*C.GtkLabel)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetLineWrap is a wrapper around the C function gtk_label_get_line_wrap.
func (recv *Label) GetLineWrap() bool {
	retC := C.gtk_label_get_line_wrap((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMnemonicKeyval is a wrapper around the C function gtk_label_get_mnemonic_keyval.
func (recv *Label) GetMnemonicKeyval() uint32 {
	retC := C.gtk_label_get_mnemonic_keyval((*C.GtkLabel)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetMnemonicWidget is a wrapper around the C function gtk_label_get_mnemonic_widget.
func (recv *Label) GetMnemonicWidget() *Widget {
	retC := C.gtk_label_get_mnemonic_widget((*C.GtkLabel)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectable is a wrapper around the C function gtk_label_get_selectable.
func (recv *Label) GetSelectable() bool {
	retC := C.gtk_label_get_selectable((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectionBounds is a wrapper around the C function gtk_label_get_selection_bounds.
func (recv *Label) GetSelectionBounds() (bool, int32, int32) {
	var c_start C.gint

	var c_end C.gint

	retC := C.gtk_label_get_selection_bounds((*C.GtkLabel)(recv.native), &c_start, &c_end)
	retGo := retC == C.TRUE

	start := (int32)(c_start)

	end := (int32)(c_end)

	return retGo, start, end
}

// GetText is a wrapper around the C function gtk_label_get_text.
func (recv *Label) GetText() string {
	retC := C.gtk_label_get_text((*C.GtkLabel)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseMarkup is a wrapper around the C function gtk_label_get_use_markup.
func (recv *Label) GetUseMarkup() bool {
	retC := C.gtk_label_get_use_markup((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_label_get_use_underline.
func (recv *Label) GetUseUnderline() bool {
	retC := C.gtk_label_get_use_underline((*C.GtkLabel)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SelectRegion is a wrapper around the C function gtk_label_select_region.
func (recv *Label) SelectRegion(startOffset int32, endOffset int32) {
	c_start_offset := (C.gint)(startOffset)

	c_end_offset := (C.gint)(endOffset)

	C.gtk_label_select_region((*C.GtkLabel)(recv.native), c_start_offset, c_end_offset)

	return
}

// SetAttributes is a wrapper around the C function gtk_label_set_attributes.
func (recv *Label) SetAttributes(attrs *pango.AttrList) {
	c_attrs := (*C.PangoAttrList)(attrs.ToC())

	C.gtk_label_set_attributes((*C.GtkLabel)(recv.native), c_attrs)

	return
}

// SetJustify is a wrapper around the C function gtk_label_set_justify.
func (recv *Label) SetJustify(jtype Justification) {
	c_jtype := (C.GtkJustification)(jtype)

	C.gtk_label_set_justify((*C.GtkLabel)(recv.native), c_jtype)

	return
}

// SetLabel is a wrapper around the C function gtk_label_set_label.
func (recv *Label) SetLabel(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_label((*C.GtkLabel)(recv.native), c_str)

	return
}

// SetLineWrap is a wrapper around the C function gtk_label_set_line_wrap.
func (recv *Label) SetLineWrap(wrap bool) {
	c_wrap :=
		boolToGboolean(wrap)

	C.gtk_label_set_line_wrap((*C.GtkLabel)(recv.native), c_wrap)

	return
}

// SetMarkup is a wrapper around the C function gtk_label_set_markup.
func (recv *Label) SetMarkup(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_markup((*C.GtkLabel)(recv.native), c_str)

	return
}

// SetMarkupWithMnemonic is a wrapper around the C function gtk_label_set_markup_with_mnemonic.
func (recv *Label) SetMarkupWithMnemonic(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_markup_with_mnemonic((*C.GtkLabel)(recv.native), c_str)

	return
}

// SetMnemonicWidget is a wrapper around the C function gtk_label_set_mnemonic_widget.
func (recv *Label) SetMnemonicWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_label_set_mnemonic_widget((*C.GtkLabel)(recv.native), c_widget)

	return
}

// SetPattern is a wrapper around the C function gtk_label_set_pattern.
func (recv *Label) SetPattern(pattern string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_label_set_pattern((*C.GtkLabel)(recv.native), c_pattern)

	return
}

// SetSelectable is a wrapper around the C function gtk_label_set_selectable.
func (recv *Label) SetSelectable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_label_set_selectable((*C.GtkLabel)(recv.native), c_setting)

	return
}

// SetText is a wrapper around the C function gtk_label_set_text.
func (recv *Label) SetText(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_text((*C.GtkLabel)(recv.native), c_str)

	return
}

// SetTextWithMnemonic is a wrapper around the C function gtk_label_set_text_with_mnemonic.
func (recv *Label) SetTextWithMnemonic(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_text_with_mnemonic((*C.GtkLabel)(recv.native), c_str)

	return
}

// SetUseMarkup is a wrapper around the C function gtk_label_set_use_markup.
func (recv *Label) SetUseMarkup(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_label_set_use_markup((*C.GtkLabel)(recv.native), c_setting)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_label_set_use_underline.
func (recv *Label) SetUseUnderline(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_label_set_use_underline((*C.GtkLabel)(recv.native), c_setting)

	return
}

// LabelAccessible is a wrapper around the C record GtkLabelAccessible.
type LabelAccessible struct {
	native *C.GtkLabelAccessible
	// parent : record
	// priv : record
}

func LabelAccessibleNewFromC(u unsafe.Pointer) *LabelAccessible {
	c := (*C.GtkLabelAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LabelAccessible{native: c}

	return g
}

func (recv *LabelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *LabelAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *LabelAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *LabelAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to LabelAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a LabelAccessible.
func CastToLabelAccessible(object *gobject.Object) *LabelAccessible {
	return LabelAccessibleNewFromC(object.ToC())
}

// Layout is a wrapper around the C record GtkLayout.
type Layout struct {
	native *C.GtkLayout
	// container : record
	// Private : priv
}

func LayoutNewFromC(u unsafe.Pointer) *Layout {
	c := (*C.GtkLayout)(u)
	if c == nil {
		return nil
	}

	g := &Layout{native: c}

	return g
}

func (recv *Layout) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Layout) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Layout) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Layout) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Layout) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Layout.
// Exercise care, as this is a potentially dangerous function if the Object is not a Layout.
func CastToLayout(object *gobject.Object) *Layout {
	return LayoutNewFromC(object.ToC())
}

// LayoutNew is a wrapper around the C function gtk_layout_new.
func LayoutNew(hadjustment *Adjustment, vadjustment *Adjustment) *Layout {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_layout_new(c_hadjustment, c_vadjustment)
	retGo := LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_layout_get_hadjustment.
func (recv *Layout) GetHadjustment() *Adjustment {
	retC := C.gtk_layout_get_hadjustment((*C.GtkLayout)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSize is a wrapper around the C function gtk_layout_get_size.
func (recv *Layout) GetSize() (uint32, uint32) {
	var c_width C.guint

	var c_height C.guint

	C.gtk_layout_get_size((*C.GtkLayout)(recv.native), &c_width, &c_height)

	width := (uint32)(c_width)

	height := (uint32)(c_height)

	return width, height
}

// GetVadjustment is a wrapper around the C function gtk_layout_get_vadjustment.
func (recv *Layout) GetVadjustment() *Adjustment {
	retC := C.gtk_layout_get_vadjustment((*C.GtkLayout)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Move is a wrapper around the C function gtk_layout_move.
func (recv *Layout) Move(childWidget *Widget, x int32, y int32) {
	c_child_widget := (*C.GtkWidget)(childWidget.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_layout_move((*C.GtkLayout)(recv.native), c_child_widget, c_x, c_y)

	return
}

// Put is a wrapper around the C function gtk_layout_put.
func (recv *Layout) Put(childWidget *Widget, x int32, y int32) {
	c_child_widget := (*C.GtkWidget)(childWidget.ToC())

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_layout_put((*C.GtkLayout)(recv.native), c_child_widget, c_x, c_y)

	return
}

// SetHadjustment is a wrapper around the C function gtk_layout_set_hadjustment.
func (recv *Layout) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_layout_set_hadjustment((*C.GtkLayout)(recv.native), c_adjustment)

	return
}

// SetSize is a wrapper around the C function gtk_layout_set_size.
func (recv *Layout) SetSize(width uint32, height uint32) {
	c_width := (C.guint)(width)

	c_height := (C.guint)(height)

	C.gtk_layout_set_size((*C.GtkLayout)(recv.native), c_width, c_height)

	return
}

// SetVadjustment is a wrapper around the C function gtk_layout_set_vadjustment.
func (recv *Layout) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_layout_set_vadjustment((*C.GtkLayout)(recv.native), c_adjustment)

	return
}

// LevelBar is a wrapper around the C record GtkLevelBar.
type LevelBar struct {
	native *C.GtkLevelBar
	// Private : parent
	// Private : priv
}

func LevelBarNewFromC(u unsafe.Pointer) *LevelBar {
	c := (*C.GtkLevelBar)(u)
	if c == nil {
		return nil
	}

	g := &LevelBar{native: c}

	return g
}

func (recv *LevelBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *LevelBar) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *LevelBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *LevelBar) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to LevelBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a LevelBar.
func CastToLevelBar(object *gobject.Object) *LevelBar {
	return LevelBarNewFromC(object.ToC())
}

// LevelBarAccessible is a wrapper around the C record GtkLevelBarAccessible.
type LevelBarAccessible struct {
	native *C.GtkLevelBarAccessible
	// parent : record
	// priv : record
}

func LevelBarAccessibleNewFromC(u unsafe.Pointer) *LevelBarAccessible {
	c := (*C.GtkLevelBarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LevelBarAccessible{native: c}

	return g
}

func (recv *LevelBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *LevelBarAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *LevelBarAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *LevelBarAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to LevelBarAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a LevelBarAccessible.
func CastToLevelBarAccessible(object *gobject.Object) *LevelBarAccessible {
	return LevelBarAccessibleNewFromC(object.ToC())
}

// LinkButton is a wrapper around the C record GtkLinkButton.
type LinkButton struct {
	native *C.GtkLinkButton
	// Private : parent_instance
	// Private : priv
}

func LinkButtonNewFromC(u unsafe.Pointer) *LinkButton {
	c := (*C.GtkLinkButton)(u)
	if c == nil {
		return nil
	}

	g := &LinkButton{native: c}

	return g
}

func (recv *LinkButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *LinkButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *LinkButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *LinkButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *LinkButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *LinkButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *LinkButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to LinkButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a LinkButton.
func CastToLinkButton(object *gobject.Object) *LinkButton {
	return LinkButtonNewFromC(object.ToC())
}

var signalActivateLinkId int
var signalActivateLinkMap = make(map[int]LinkButtonSignalActivateLinkCallback)
var signalActivateLinkLock sync.Mutex

// LinkButtonSignalActivateLinkCallback is a callback function for a 'activate-link' signal emitted from a LinkButton.
type LinkButtonSignalActivateLinkCallback func() bool

func LinkButton_activateLinkHandler() {}

// LinkButtonAccessible is a wrapper around the C record GtkLinkButtonAccessible.
type LinkButtonAccessible struct {
	native *C.GtkLinkButtonAccessible
	// parent : record
	// priv : record
}

func LinkButtonAccessibleNewFromC(u unsafe.Pointer) *LinkButtonAccessible {
	c := (*C.GtkLinkButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LinkButtonAccessible{native: c}

	return g
}

func (recv *LinkButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *LinkButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return ButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *LinkButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *LinkButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *LinkButtonAccessible) Accessible() *Accessible {
	return recv.ButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *LinkButtonAccessible) Object() *atk.Object {
	return recv.ButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to LinkButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a LinkButtonAccessible.
func CastToLinkButtonAccessible(object *gobject.Object) *LinkButtonAccessible {
	return LinkButtonAccessibleNewFromC(object.ToC())
}

// ListBox is a wrapper around the C record GtkListBox.
type ListBox struct {
	native *C.GtkListBox
	// parent_instance : record
}

func ListBoxNewFromC(u unsafe.Pointer) *ListBox {
	c := (*C.GtkListBox)(u)
	if c == nil {
		return nil
	}

	g := &ListBox{native: c}

	return g
}

func (recv *ListBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *ListBox) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *ListBox) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ListBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ListBox) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to ListBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListBox.
func CastToListBox(object *gobject.Object) *ListBox {
	return ListBoxNewFromC(object.ToC())
}

var signalActivateCursorRowId int
var signalActivateCursorRowMap = make(map[int]ListBoxSignalActivateCursorRowCallback)
var signalActivateCursorRowLock sync.Mutex

// ListBoxSignalActivateCursorRowCallback is a callback function for a 'activate-cursor-row' signal emitted from a ListBox.
type ListBoxSignalActivateCursorRowCallback func()

func ListBox_activateCursorRowHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]ListBoxSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// ListBoxSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a ListBox.
type ListBoxSignalMoveCursorCallback func(object MovementStep, p0 int32)

func ListBox_moveCursorHandler() {}

var signalToggleCursorRowId int
var signalToggleCursorRowMap = make(map[int]ListBoxSignalToggleCursorRowCallback)
var signalToggleCursorRowLock sync.Mutex

// ListBoxSignalToggleCursorRowCallback is a callback function for a 'toggle-cursor-row' signal emitted from a ListBox.
type ListBoxSignalToggleCursorRowCallback func()

func ListBox_toggleCursorRowHandler() {}

// ListBoxAccessible is a wrapper around the C record GtkListBoxAccessible.
type ListBoxAccessible struct {
	native *C.GtkListBoxAccessible
	// parent : record
	// priv : record
}

func ListBoxAccessibleNewFromC(u unsafe.Pointer) *ListBoxAccessible {
	c := (*C.GtkListBoxAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxAccessible{native: c}

	return g
}

func (recv *ListBoxAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ListBoxAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ListBoxAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ListBoxAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ListBoxAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ListBoxAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListBoxAccessible.
func CastToListBoxAccessible(object *gobject.Object) *ListBoxAccessible {
	return ListBoxAccessibleNewFromC(object.ToC())
}

// ListBoxRow is a wrapper around the C record GtkListBoxRow.
type ListBoxRow struct {
	native *C.GtkListBoxRow
	// parent_instance : record
}

func ListBoxRowNewFromC(u unsafe.Pointer) *ListBoxRow {
	c := (*C.GtkListBoxRow)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRow{native: c}

	return g
}

func (recv *ListBoxRow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *ListBoxRow) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ListBoxRow) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *ListBoxRow) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ListBoxRow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ListBoxRow) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to ListBoxRow.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListBoxRow.
func CastToListBoxRow(object *gobject.Object) *ListBoxRow {
	return ListBoxRowNewFromC(object.ToC())
}

// ListBoxRowAccessible is a wrapper around the C record GtkListBoxRowAccessible.
type ListBoxRowAccessible struct {
	native *C.GtkListBoxRowAccessible
	// parent : record
}

func ListBoxRowAccessibleNewFromC(u unsafe.Pointer) *ListBoxRowAccessible {
	c := (*C.GtkListBoxRowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ListBoxRowAccessible{native: c}

	return g
}

func (recv *ListBoxRowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ListBoxRowAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ListBoxRowAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ListBoxRowAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ListBoxRowAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ListBoxRowAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListBoxRowAccessible.
func CastToListBoxRowAccessible(object *gobject.Object) *ListBoxRowAccessible {
	return ListBoxRowAccessibleNewFromC(object.ToC())
}

// ListStore is a wrapper around the C record GtkListStore.
type ListStore struct {
	native *C.GtkListStore
	// parent : record
	// Private : priv
}

func ListStoreNewFromC(u unsafe.Pointer) *ListStore {
	c := (*C.GtkListStore)(u)
	if c == nil {
		return nil
	}

	g := &ListStore{native: c}

	return g
}

func (recv *ListStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ListStore) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ListStore.
// Exercise care, as this is a potentially dangerous function if the Object is not a ListStore.
func CastToListStore(object *gobject.Object) *ListStore {
	return ListStoreNewFromC(object.ToC())
}

// Unsupported : gtk_list_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_newv : unsupported parameter types : no param type

// Append is a wrapper around the C function gtk_list_store_append.
func (recv *ListStore) Append() *TreeIter {
	var c_iter C.GtkTreeIter

	C.gtk_list_store_append((*C.GtkListStore)(recv.native), &c_iter)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Clear is a wrapper around the C function gtk_list_store_clear.
func (recv *ListStore) Clear() {
	C.gtk_list_store_clear((*C.GtkListStore)(recv.native))

	return
}

// Insert is a wrapper around the C function gtk_list_store_insert.
func (recv *ListStore) Insert(position int32) *TreeIter {
	var c_iter C.GtkTreeIter

	c_position := (C.gint)(position)

	C.gtk_list_store_insert((*C.GtkListStore)(recv.native), &c_iter, c_position)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// InsertAfter is a wrapper around the C function gtk_list_store_insert_after.
func (recv *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_sibling := (*C.GtkTreeIter)(sibling.ToC())

	C.gtk_list_store_insert_after((*C.GtkListStore)(recv.native), &c_iter, c_sibling)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// InsertBefore is a wrapper around the C function gtk_list_store_insert_before.
func (recv *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_sibling := (*C.GtkTreeIter)(sibling.ToC())

	C.gtk_list_store_insert_before((*C.GtkListStore)(recv.native), &c_iter, c_sibling)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Prepend is a wrapper around the C function gtk_list_store_prepend.
func (recv *ListStore) Prepend() *TreeIter {
	var c_iter C.GtkTreeIter

	C.gtk_list_store_prepend((*C.GtkListStore)(recv.native), &c_iter)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Remove is a wrapper around the C function gtk_list_store_remove.
func (recv *ListStore) Remove(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_list_store_remove((*C.GtkListStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_list_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_list_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_list_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// SetValue is a wrapper around the C function gtk_list_store_set_value.
func (recv *ListStore) SetValue(iter *TreeIter, column int32, value *gobject.Value) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_column := (C.gint)(column)

	c_value := (*C.GValue)(value.ToC())

	C.gtk_list_store_set_value((*C.GtkListStore)(recv.native), c_iter, c_column, c_value)

	return
}

// LockButton is a wrapper around the C record GtkLockButton.
type LockButton struct {
	native *C.GtkLockButton
	// parent : record
	// priv : record
}

func LockButtonNewFromC(u unsafe.Pointer) *LockButton {
	c := (*C.GtkLockButton)(u)
	if c == nil {
		return nil
	}

	g := &LockButton{native: c}

	return g
}

func (recv *LockButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *LockButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *LockButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *LockButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *LockButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *LockButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *LockButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to LockButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a LockButton.
func CastToLockButton(object *gobject.Object) *LockButton {
	return LockButtonNewFromC(object.ToC())
}

// LockButtonAccessible is a wrapper around the C record GtkLockButtonAccessible.
type LockButtonAccessible struct {
	native *C.GtkLockButtonAccessible
	// parent : record
	// priv : record
}

func LockButtonAccessibleNewFromC(u unsafe.Pointer) *LockButtonAccessible {
	c := (*C.GtkLockButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &LockButtonAccessible{native: c}

	return g
}

func (recv *LockButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *LockButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return ButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *LockButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *LockButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *LockButtonAccessible) Accessible() *Accessible {
	return recv.ButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *LockButtonAccessible) Object() *atk.Object {
	return recv.ButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to LockButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a LockButtonAccessible.
func CastToLockButtonAccessible(object *gobject.Object) *LockButtonAccessible {
	return LockButtonAccessibleNewFromC(object.ToC())
}

// Menu is a wrapper around the C record GtkMenu.
type Menu struct {
	native *C.GtkMenu
	// menu_shell : record
	// Private : priv
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GtkMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShell upcasts to *MenuShell
func (recv *Menu) MenuShell() *MenuShell {
	return MenuShellNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Menu) Container() *Container {
	return recv.MenuShell().Container()
}

// Widget upcasts to *Widget
func (recv *Menu) Widget() *Widget {
	return recv.MenuShell().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Menu) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuShell().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Menu) Object() *gobject.Object {
	return recv.MenuShell().Object()
}

// CastToWidget down casts any arbitary Object to Menu.
// Exercise care, as this is a potentially dangerous function if the Object is not a Menu.
func CastToMenu(object *gobject.Object) *Menu {
	return MenuNewFromC(object.ToC())
}

var signalMoveScrollId int
var signalMoveScrollMap = make(map[int]MenuSignalMoveScrollCallback)
var signalMoveScrollLock sync.Mutex

// MenuSignalMoveScrollCallback is a callback function for a 'move-scroll' signal emitted from a Menu.
type MenuSignalMoveScrollCallback func(scrollType ScrollType)

func Menu_moveScrollHandler() {}

// MenuNew is a wrapper around the C function gtk_menu_new.
func MenuNew() *Menu {
	retC := C.gtk_menu_new()
	retGo := MenuNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_menu_attach_to_widget : unsupported parameter detacher : no type generator for MenuDetachFunc, GtkMenuDetachFunc

// Detach is a wrapper around the C function gtk_menu_detach.
func (recv *Menu) Detach() {
	C.gtk_menu_detach((*C.GtkMenu)(recv.native))

	return
}

// GetAccelGroup is a wrapper around the C function gtk_menu_get_accel_group.
func (recv *Menu) GetAccelGroup() *AccelGroup {
	retC := C.gtk_menu_get_accel_group((*C.GtkMenu)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_menu_get_active.
func (recv *Menu) GetActive() *Widget {
	retC := C.gtk_menu_get_active((*C.GtkMenu)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAttachWidget is a wrapper around the C function gtk_menu_get_attach_widget.
func (recv *Menu) GetAttachWidget() *Widget {
	retC := C.gtk_menu_get_attach_widget((*C.GtkMenu)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTearoffState is a wrapper around the C function gtk_menu_get_tearoff_state.
func (recv *Menu) GetTearoffState() bool {
	retC := C.gtk_menu_get_tearoff_state((*C.GtkMenu)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_menu_get_title.
func (recv *Menu) GetTitle() string {
	retC := C.gtk_menu_get_title((*C.GtkMenu)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Popdown is a wrapper around the C function gtk_menu_popdown.
func (recv *Menu) Popdown() {
	C.gtk_menu_popdown((*C.GtkMenu)(recv.native))

	return
}

// Unsupported : gtk_menu_popup : unsupported parameter func : no type generator for MenuPositionFunc, GtkMenuPositionFunc

// ReorderChild is a wrapper around the C function gtk_menu_reorder_child.
func (recv *Menu) ReorderChild(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_menu_reorder_child((*C.GtkMenu)(recv.native), c_child, c_position)

	return
}

// Reposition is a wrapper around the C function gtk_menu_reposition.
func (recv *Menu) Reposition() {
	C.gtk_menu_reposition((*C.GtkMenu)(recv.native))

	return
}

// SetAccelGroup is a wrapper around the C function gtk_menu_set_accel_group.
func (recv *Menu) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_menu_set_accel_group((*C.GtkMenu)(recv.native), c_accel_group)

	return
}

// SetAccelPath is a wrapper around the C function gtk_menu_set_accel_path.
func (recv *Menu) SetAccelPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_menu_set_accel_path((*C.GtkMenu)(recv.native), c_accel_path)

	return
}

// SetActive is a wrapper around the C function gtk_menu_set_active.
func (recv *Menu) SetActive(index uint32) {
	c_index := (C.guint)(index)

	C.gtk_menu_set_active((*C.GtkMenu)(recv.native), c_index)

	return
}

// SetTearoffState is a wrapper around the C function gtk_menu_set_tearoff_state.
func (recv *Menu) SetTearoffState(tornOff bool) {
	c_torn_off :=
		boolToGboolean(tornOff)

	C.gtk_menu_set_tearoff_state((*C.GtkMenu)(recv.native), c_torn_off)

	return
}

// SetTitle is a wrapper around the C function gtk_menu_set_title.
func (recv *Menu) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_menu_set_title((*C.GtkMenu)(recv.native), c_title)

	return
}

// MenuAccessible is a wrapper around the C record GtkMenuAccessible.
type MenuAccessible struct {
	native *C.GtkMenuAccessible
	// parent : record
	// priv : record
}

func MenuAccessibleNewFromC(u unsafe.Pointer) *MenuAccessible {
	c := (*C.GtkMenuAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuAccessible{native: c}

	return g
}

func (recv *MenuAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShellAccessible upcasts to *MenuShellAccessible
func (recv *MenuAccessible) MenuShellAccessible() *MenuShellAccessible {
	return MenuShellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *MenuAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.MenuShellAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *MenuAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.MenuShellAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *MenuAccessible) Accessible() *Accessible {
	return recv.MenuShellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *MenuAccessible) Object() *atk.Object {
	return recv.MenuShellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to MenuAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuAccessible.
func CastToMenuAccessible(object *gobject.Object) *MenuAccessible {
	return MenuAccessibleNewFromC(object.ToC())
}

// MenuBar is a wrapper around the C record GtkMenuBar.
type MenuBar struct {
	native *C.GtkMenuBar
	// menu_shell : record
	// Private : priv
}

func MenuBarNewFromC(u unsafe.Pointer) *MenuBar {
	c := (*C.GtkMenuBar)(u)
	if c == nil {
		return nil
	}

	g := &MenuBar{native: c}

	return g
}

func (recv *MenuBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuShell upcasts to *MenuShell
func (recv *MenuBar) MenuShell() *MenuShell {
	return MenuShellNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *MenuBar) Container() *Container {
	return recv.MenuShell().Container()
}

// Widget upcasts to *Widget
func (recv *MenuBar) Widget() *Widget {
	return recv.MenuShell().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MenuBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuShell().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MenuBar) Object() *gobject.Object {
	return recv.MenuShell().Object()
}

// CastToWidget down casts any arbitary Object to MenuBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuBar.
func CastToMenuBar(object *gobject.Object) *MenuBar {
	return MenuBarNewFromC(object.ToC())
}

// MenuBarNew is a wrapper around the C function gtk_menu_bar_new.
func MenuBarNew() *MenuBar {
	retC := C.gtk_menu_bar_new()
	retGo := MenuBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuButton is a wrapper around the C record GtkMenuButton.
type MenuButton struct {
	native *C.GtkMenuButton
	// parent : record
	// Private : priv
}

func MenuButtonNewFromC(u unsafe.Pointer) *MenuButton {
	c := (*C.GtkMenuButton)(u)
	if c == nil {
		return nil
	}

	g := &MenuButton{native: c}

	return g
}

func (recv *MenuButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButton upcasts to *ToggleButton
func (recv *MenuButton) ToggleButton() *ToggleButton {
	return ToggleButtonNewFromC(unsafe.Pointer(recv.native))
}

// Button upcasts to *Button
func (recv *MenuButton) Button() *Button {
	return recv.ToggleButton().Button()
}

// Bin upcasts to *Bin
func (recv *MenuButton) Bin() *Bin {
	return recv.ToggleButton().Bin()
}

// Container upcasts to *Container
func (recv *MenuButton) Container() *Container {
	return recv.ToggleButton().Container()
}

// Widget upcasts to *Widget
func (recv *MenuButton) Widget() *Widget {
	return recv.ToggleButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MenuButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToggleButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MenuButton) Object() *gobject.Object {
	return recv.ToggleButton().Object()
}

// CastToWidget down casts any arbitary Object to MenuButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuButton.
func CastToMenuButton(object *gobject.Object) *MenuButton {
	return MenuButtonNewFromC(object.ToC())
}

// MenuButtonAccessible is a wrapper around the C record GtkMenuButtonAccessible.
type MenuButtonAccessible struct {
	native *C.GtkMenuButtonAccessible
	// parent : record
	// priv : record
}

func MenuButtonAccessibleNewFromC(u unsafe.Pointer) *MenuButtonAccessible {
	c := (*C.GtkMenuButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuButtonAccessible{native: c}

	return g
}

func (recv *MenuButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonAccessible upcasts to *ToggleButtonAccessible
func (recv *MenuButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {
	return ToggleButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *MenuButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return recv.ToggleButtonAccessible().ButtonAccessible()
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *MenuButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ToggleButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *MenuButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ToggleButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *MenuButtonAccessible) Accessible() *Accessible {
	return recv.ToggleButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *MenuButtonAccessible) Object() *atk.Object {
	return recv.ToggleButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to MenuButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuButtonAccessible.
func CastToMenuButtonAccessible(object *gobject.Object) *MenuButtonAccessible {
	return MenuButtonAccessibleNewFromC(object.ToC())
}

// MenuItem is a wrapper around the C record GtkMenuItem.
type MenuItem struct {
	native *C.GtkMenuItem
	// bin : record
	// Private : priv
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GtkMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *MenuItem) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *MenuItem) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *MenuItem) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MenuItem) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to MenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuItem.
func CastToMenuItem(object *gobject.Object) *MenuItem {
	return MenuItemNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]MenuItemSignalActivateCallback)
var signalActivateLock sync.Mutex

// MenuItemSignalActivateCallback is a callback function for a 'activate' signal emitted from a MenuItem.
type MenuItemSignalActivateCallback func()

func MenuItem_activateHandler() {}

var signalActivateItemId int
var signalActivateItemMap = make(map[int]MenuItemSignalActivateItemCallback)
var signalActivateItemLock sync.Mutex

// MenuItemSignalActivateItemCallback is a callback function for a 'activate-item' signal emitted from a MenuItem.
type MenuItemSignalActivateItemCallback func()

func MenuItem_activateItemHandler() {}

var signalDeselectId int
var signalDeselectMap = make(map[int]MenuItemSignalDeselectCallback)
var signalDeselectLock sync.Mutex

// MenuItemSignalDeselectCallback is a callback function for a 'deselect' signal emitted from a MenuItem.
type MenuItemSignalDeselectCallback func()

func MenuItem_deselectHandler() {}

var signalSelectId int
var signalSelectMap = make(map[int]MenuItemSignalSelectCallback)
var signalSelectLock sync.Mutex

// MenuItemSignalSelectCallback is a callback function for a 'select' signal emitted from a MenuItem.
type MenuItemSignalSelectCallback func()

func MenuItem_selectHandler() {}

var signalToggleSizeAllocateId int
var signalToggleSizeAllocateMap = make(map[int]MenuItemSignalToggleSizeAllocateCallback)
var signalToggleSizeAllocateLock sync.Mutex

// MenuItemSignalToggleSizeAllocateCallback is a callback function for a 'toggle-size-allocate' signal emitted from a MenuItem.
type MenuItemSignalToggleSizeAllocateCallback func(object int32)

func MenuItem_toggleSizeAllocateHandler() {}

var signalToggleSizeRequestId int
var signalToggleSizeRequestMap = make(map[int]MenuItemSignalToggleSizeRequestCallback)
var signalToggleSizeRequestLock sync.Mutex

// MenuItemSignalToggleSizeRequestCallback is a callback function for a 'toggle-size-request' signal emitted from a MenuItem.
type MenuItemSignalToggleSizeRequestCallback func(object uintptr)

func MenuItem_toggleSizeRequestHandler() {}

// MenuItemNew is a wrapper around the C function gtk_menu_item_new.
func MenuItemNew() *MenuItem {
	retC := C.gtk_menu_item_new()
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithLabel is a wrapper around the C function gtk_menu_item_new_with_label.
func MenuItemNewWithLabel(label string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_label(c_label)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// MenuItemNewWithMnemonic is a wrapper around the C function gtk_menu_item_new_with_mnemonic.
func MenuItemNewWithMnemonic(label string) *MenuItem {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_menu_item_new_with_mnemonic(c_label)
	retGo := MenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Activate is a wrapper around the C function gtk_menu_item_activate.
func (recv *MenuItem) Activate() {
	C.gtk_menu_item_activate((*C.GtkMenuItem)(recv.native))

	return
}

// Deselect is a wrapper around the C function gtk_menu_item_deselect.
func (recv *MenuItem) Deselect() {
	C.gtk_menu_item_deselect((*C.GtkMenuItem)(recv.native))

	return
}

// GetRightJustified is a wrapper around the C function gtk_menu_item_get_right_justified.
func (recv *MenuItem) GetRightJustified() bool {
	retC := C.gtk_menu_item_get_right_justified((*C.GtkMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSubmenu is a wrapper around the C function gtk_menu_item_get_submenu.
func (recv *MenuItem) GetSubmenu() *Widget {
	retC := C.gtk_menu_item_get_submenu((*C.GtkMenuItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Select is a wrapper around the C function gtk_menu_item_select.
func (recv *MenuItem) Select() {
	C.gtk_menu_item_select((*C.GtkMenuItem)(recv.native))

	return
}

// SetAccelPath is a wrapper around the C function gtk_menu_item_set_accel_path.
func (recv *MenuItem) SetAccelPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_menu_item_set_accel_path((*C.GtkMenuItem)(recv.native), c_accel_path)

	return
}

// SetRightJustified is a wrapper around the C function gtk_menu_item_set_right_justified.
func (recv *MenuItem) SetRightJustified(rightJustified bool) {
	c_right_justified :=
		boolToGboolean(rightJustified)

	C.gtk_menu_item_set_right_justified((*C.GtkMenuItem)(recv.native), c_right_justified)

	return
}

// SetSubmenu is a wrapper around the C function gtk_menu_item_set_submenu.
func (recv *MenuItem) SetSubmenu(submenu *Menu) {
	c_submenu := (*C.GtkWidget)(submenu.ToC())

	C.gtk_menu_item_set_submenu((*C.GtkMenuItem)(recv.native), c_submenu)

	return
}

// ToggleSizeAllocate is a wrapper around the C function gtk_menu_item_toggle_size_allocate.
func (recv *MenuItem) ToggleSizeAllocate(allocation int32) {
	c_allocation := (C.gint)(allocation)

	C.gtk_menu_item_toggle_size_allocate((*C.GtkMenuItem)(recv.native), c_allocation)

	return
}

// ToggleSizeRequest is a wrapper around the C function gtk_menu_item_toggle_size_request.
func (recv *MenuItem) ToggleSizeRequest(requisition int32) {
	c_requisition := (C.gint)(requisition)

	C.gtk_menu_item_toggle_size_request((*C.GtkMenuItem)(recv.native), &c_requisition)

	return
}

// MenuItemAccessible is a wrapper around the C record GtkMenuItemAccessible.
type MenuItemAccessible struct {
	native *C.GtkMenuItemAccessible
	// parent : record
	// priv : record
}

func MenuItemAccessibleNewFromC(u unsafe.Pointer) *MenuItemAccessible {
	c := (*C.GtkMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuItemAccessible{native: c}

	return g
}

func (recv *MenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *MenuItemAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *MenuItemAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *MenuItemAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *MenuItemAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to MenuItemAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuItemAccessible.
func CastToMenuItemAccessible(object *gobject.Object) *MenuItemAccessible {
	return MenuItemAccessibleNewFromC(object.ToC())
}

// MenuShell is a wrapper around the C record GtkMenuShell.
type MenuShell struct {
	native *C.GtkMenuShell
	// container : record
	// Private : priv
}

func MenuShellNewFromC(u unsafe.Pointer) *MenuShell {
	c := (*C.GtkMenuShell)(u)
	if c == nil {
		return nil
	}

	g := &MenuShell{native: c}

	return g
}

func (recv *MenuShell) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *MenuShell) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *MenuShell) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MenuShell) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MenuShell) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to MenuShell.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuShell.
func CastToMenuShell(object *gobject.Object) *MenuShell {
	return MenuShellNewFromC(object.ToC())
}

var signalActivateCurrentId int
var signalActivateCurrentMap = make(map[int]MenuShellSignalActivateCurrentCallback)
var signalActivateCurrentLock sync.Mutex

// MenuShellSignalActivateCurrentCallback is a callback function for a 'activate-current' signal emitted from a MenuShell.
type MenuShellSignalActivateCurrentCallback func(forceHide bool)

func MenuShell_activateCurrentHandler() {}

var signalCancelId int
var signalCancelMap = make(map[int]MenuShellSignalCancelCallback)
var signalCancelLock sync.Mutex

// MenuShellSignalCancelCallback is a callback function for a 'cancel' signal emitted from a MenuShell.
type MenuShellSignalCancelCallback func()

func MenuShell_cancelHandler() {}

var signalCycleFocusId int
var signalCycleFocusMap = make(map[int]MenuShellSignalCycleFocusCallback)
var signalCycleFocusLock sync.Mutex

// MenuShellSignalCycleFocusCallback is a callback function for a 'cycle-focus' signal emitted from a MenuShell.
type MenuShellSignalCycleFocusCallback func(direction DirectionType)

func MenuShell_cycleFocusHandler() {}

var signalDeactivateId int
var signalDeactivateMap = make(map[int]MenuShellSignalDeactivateCallback)
var signalDeactivateLock sync.Mutex

// MenuShellSignalDeactivateCallback is a callback function for a 'deactivate' signal emitted from a MenuShell.
type MenuShellSignalDeactivateCallback func()

func MenuShell_deactivateHandler() {}

var signalMoveCurrentId int
var signalMoveCurrentMap = make(map[int]MenuShellSignalMoveCurrentCallback)
var signalMoveCurrentLock sync.Mutex

// MenuShellSignalMoveCurrentCallback is a callback function for a 'move-current' signal emitted from a MenuShell.
type MenuShellSignalMoveCurrentCallback func(direction MenuDirectionType)

func MenuShell_moveCurrentHandler() {}

var signalSelectionDoneId int
var signalSelectionDoneMap = make(map[int]MenuShellSignalSelectionDoneCallback)
var signalSelectionDoneLock sync.Mutex

// MenuShellSignalSelectionDoneCallback is a callback function for a 'selection-done' signal emitted from a MenuShell.
type MenuShellSignalSelectionDoneCallback func()

func MenuShell_selectionDoneHandler() {}

// ActivateItem is a wrapper around the C function gtk_menu_shell_activate_item.
func (recv *MenuShell) ActivateItem(menuItem *Widget, forceDeactivate bool) {
	c_menu_item := (*C.GtkWidget)(menuItem.ToC())

	c_force_deactivate :=
		boolToGboolean(forceDeactivate)

	C.gtk_menu_shell_activate_item((*C.GtkMenuShell)(recv.native), c_menu_item, c_force_deactivate)

	return
}

// Append is a wrapper around the C function gtk_menu_shell_append.
func (recv *MenuShell) Append(child *MenuItem) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_menu_shell_append((*C.GtkMenuShell)(recv.native), c_child)

	return
}

// Deactivate is a wrapper around the C function gtk_menu_shell_deactivate.
func (recv *MenuShell) Deactivate() {
	C.gtk_menu_shell_deactivate((*C.GtkMenuShell)(recv.native))

	return
}

// Deselect is a wrapper around the C function gtk_menu_shell_deselect.
func (recv *MenuShell) Deselect() {
	C.gtk_menu_shell_deselect((*C.GtkMenuShell)(recv.native))

	return
}

// Insert is a wrapper around the C function gtk_menu_shell_insert.
func (recv *MenuShell) Insert(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_menu_shell_insert((*C.GtkMenuShell)(recv.native), c_child, c_position)

	return
}

// Prepend is a wrapper around the C function gtk_menu_shell_prepend.
func (recv *MenuShell) Prepend(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_menu_shell_prepend((*C.GtkMenuShell)(recv.native), c_child)

	return
}

// SelectItem is a wrapper around the C function gtk_menu_shell_select_item.
func (recv *MenuShell) SelectItem(menuItem *Widget) {
	c_menu_item := (*C.GtkWidget)(menuItem.ToC())

	C.gtk_menu_shell_select_item((*C.GtkMenuShell)(recv.native), c_menu_item)

	return
}

// MenuShellAccessible is a wrapper around the C record GtkMenuShellAccessible.
type MenuShellAccessible struct {
	native *C.GtkMenuShellAccessible
	// parent : record
	// priv : record
}

func MenuShellAccessibleNewFromC(u unsafe.Pointer) *MenuShellAccessible {
	c := (*C.GtkMenuShellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &MenuShellAccessible{native: c}

	return g
}

func (recv *MenuShellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *MenuShellAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *MenuShellAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *MenuShellAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *MenuShellAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to MenuShellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuShellAccessible.
func CastToMenuShellAccessible(object *gobject.Object) *MenuShellAccessible {
	return MenuShellAccessibleNewFromC(object.ToC())
}

// MenuToolButton is a wrapper around the C record GtkMenuToolButton.
type MenuToolButton struct {
	native *C.GtkMenuToolButton
	// parent : record
	// Private : priv
}

func MenuToolButtonNewFromC(u unsafe.Pointer) *MenuToolButton {
	c := (*C.GtkMenuToolButton)(u)
	if c == nil {
		return nil
	}

	g := &MenuToolButton{native: c}

	return g
}

func (recv *MenuToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolButton upcasts to *ToolButton
func (recv *MenuToolButton) ToolButton() *ToolButton {
	return ToolButtonNewFromC(unsafe.Pointer(recv.native))
}

// ToolItem upcasts to *ToolItem
func (recv *MenuToolButton) ToolItem() *ToolItem {
	return recv.ToolButton().ToolItem()
}

// Bin upcasts to *Bin
func (recv *MenuToolButton) Bin() *Bin {
	return recv.ToolButton().Bin()
}

// Container upcasts to *Container
func (recv *MenuToolButton) Container() *Container {
	return recv.ToolButton().Container()
}

// Widget upcasts to *Widget
func (recv *MenuToolButton) Widget() *Widget {
	return recv.ToolButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MenuToolButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToolButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MenuToolButton) Object() *gobject.Object {
	return recv.ToolButton().Object()
}

// CastToWidget down casts any arbitary Object to MenuToolButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a MenuToolButton.
func CastToMenuToolButton(object *gobject.Object) *MenuToolButton {
	return MenuToolButtonNewFromC(object.ToC())
}

var signalShowMenuId int
var signalShowMenuMap = make(map[int]MenuToolButtonSignalShowMenuCallback)
var signalShowMenuLock sync.Mutex

// MenuToolButtonSignalShowMenuCallback is a callback function for a 'show-menu' signal emitted from a MenuToolButton.
type MenuToolButtonSignalShowMenuCallback func()

func MenuToolButton_showMenuHandler() {}

// MessageDialog is a wrapper around the C record GtkMessageDialog.
type MessageDialog struct {
	native *C.GtkMessageDialog
	// parent_instance : record
	// Private : priv
}

func MessageDialogNewFromC(u unsafe.Pointer) *MessageDialog {
	c := (*C.GtkMessageDialog)(u)
	if c == nil {
		return nil
	}

	g := &MessageDialog{native: c}

	return g
}

func (recv *MessageDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *MessageDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *MessageDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *MessageDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *MessageDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *MessageDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *MessageDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *MessageDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to MessageDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a MessageDialog.
func CastToMessageDialog(object *gobject.Object) *MessageDialog {
	return MessageDialogNewFromC(object.ToC())
}

// Unsupported : gtk_message_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_message_dialog_new_with_markup : unsupported parameter ... : varargs

// Misc is a wrapper around the C record GtkMisc.
type Misc struct {
	native *C.GtkMisc
	// widget : record
	// Private : priv
}

func MiscNewFromC(u unsafe.Pointer) *Misc {
	c := (*C.GtkMisc)(u)
	if c == nil {
		return nil
	}

	g := &Misc{native: c}

	return g
}

func (recv *Misc) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Misc) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Misc) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Misc) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Misc.
// Exercise care, as this is a potentially dangerous function if the Object is not a Misc.
func CastToMisc(object *gobject.Object) *Misc {
	return MiscNewFromC(object.ToC())
}

// GetAlignment is a wrapper around the C function gtk_misc_get_alignment.
func (recv *Misc) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_misc_get_alignment((*C.GtkMisc)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetPadding is a wrapper around the C function gtk_misc_get_padding.
func (recv *Misc) GetPadding() (int32, int32) {
	var c_xpad C.gint

	var c_ypad C.gint

	C.gtk_misc_get_padding((*C.GtkMisc)(recv.native), &c_xpad, &c_ypad)

	xpad := (int32)(c_xpad)

	ypad := (int32)(c_ypad)

	return xpad, ypad
}

// SetAlignment is a wrapper around the C function gtk_misc_set_alignment.
func (recv *Misc) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_misc_set_alignment((*C.GtkMisc)(recv.native), c_xalign, c_yalign)

	return
}

// SetPadding is a wrapper around the C function gtk_misc_set_padding.
func (recv *Misc) SetPadding(xpad int32, ypad int32) {
	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_misc_set_padding((*C.GtkMisc)(recv.native), c_xpad, c_ypad)

	return
}

// ModelButton is a wrapper around the C record GtkModelButton.
type ModelButton struct {
	native *C.GtkModelButton
}

func ModelButtonNewFromC(u unsafe.Pointer) *ModelButton {
	c := (*C.GtkModelButton)(u)
	if c == nil {
		return nil
	}

	g := &ModelButton{native: c}

	return g
}

func (recv *ModelButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *ModelButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ModelButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *ModelButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *ModelButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ModelButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ModelButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to ModelButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ModelButton.
func CastToModelButton(object *gobject.Object) *ModelButton {
	return ModelButtonNewFromC(object.ToC())
}

// MountOperation is a wrapper around the C record GtkMountOperation.
type MountOperation struct {
	native *C.GtkMountOperation
	// parent_instance : record
	// priv : record
}

func MountOperationNewFromC(u unsafe.Pointer) *MountOperation {
	c := (*C.GtkMountOperation)(u)
	if c == nil {
		return nil
	}

	g := &MountOperation{native: c}

	return g
}

func (recv *MountOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MountOperation upcasts to *MountOperation
func (recv *MountOperation) MountOperation() *gio.MountOperation {
	return gio.MountOperationNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *MountOperation) Object() *gobject.Object {
	return recv.MountOperation().Object()
}

// CastToWidget down casts any arbitary Object to MountOperation.
// Exercise care, as this is a potentially dangerous function if the Object is not a MountOperation.
func CastToMountOperation(object *gobject.Object) *MountOperation {
	return MountOperationNewFromC(object.ToC())
}

// Notebook is a wrapper around the C record GtkNotebook.
type Notebook struct {
	native *C.GtkNotebook
	// Private : container
	// Private : priv
}

func NotebookNewFromC(u unsafe.Pointer) *Notebook {
	c := (*C.GtkNotebook)(u)
	if c == nil {
		return nil
	}

	g := &Notebook{native: c}

	return g
}

func (recv *Notebook) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Notebook) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Notebook) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Notebook) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Notebook) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Notebook.
// Exercise care, as this is a potentially dangerous function if the Object is not a Notebook.
func CastToNotebook(object *gobject.Object) *Notebook {
	return NotebookNewFromC(object.ToC())
}

var signalChangeCurrentPageId int
var signalChangeCurrentPageMap = make(map[int]NotebookSignalChangeCurrentPageCallback)
var signalChangeCurrentPageLock sync.Mutex

// NotebookSignalChangeCurrentPageCallback is a callback function for a 'change-current-page' signal emitted from a Notebook.
type NotebookSignalChangeCurrentPageCallback func(object int32) bool

func Notebook_changeCurrentPageHandler() {}

var signalFocusTabId int
var signalFocusTabMap = make(map[int]NotebookSignalFocusTabCallback)
var signalFocusTabLock sync.Mutex

// NotebookSignalFocusTabCallback is a callback function for a 'focus-tab' signal emitted from a Notebook.
type NotebookSignalFocusTabCallback func(object NotebookTab) bool

func Notebook_focusTabHandler() {}

var signalMoveFocusOutId int
var signalMoveFocusOutMap = make(map[int]NotebookSignalMoveFocusOutCallback)
var signalMoveFocusOutLock sync.Mutex

// NotebookSignalMoveFocusOutCallback is a callback function for a 'move-focus-out' signal emitted from a Notebook.
type NotebookSignalMoveFocusOutCallback func(object DirectionType)

func Notebook_moveFocusOutHandler() {}

var signalReorderTabId int
var signalReorderTabMap = make(map[int]NotebookSignalReorderTabCallback)
var signalReorderTabLock sync.Mutex

// NotebookSignalReorderTabCallback is a callback function for a 'reorder-tab' signal emitted from a Notebook.
type NotebookSignalReorderTabCallback func(object DirectionType, p0 bool) bool

func Notebook_reorderTabHandler() {}

var signalSelectPageId int
var signalSelectPageMap = make(map[int]NotebookSignalSelectPageCallback)
var signalSelectPageLock sync.Mutex

// NotebookSignalSelectPageCallback is a callback function for a 'select-page' signal emitted from a Notebook.
type NotebookSignalSelectPageCallback func(object bool) bool

func Notebook_selectPageHandler() {}

var signalSwitchPageId int
var signalSwitchPageMap = make(map[int]NotebookSignalSwitchPageCallback)
var signalSwitchPageLock sync.Mutex

// NotebookSignalSwitchPageCallback is a callback function for a 'switch-page' signal emitted from a Notebook.
type NotebookSignalSwitchPageCallback func(page *Widget, pageNum uint32)

func Notebook_switchPageHandler() {}

// NotebookNew is a wrapper around the C function gtk_notebook_new.
func NotebookNew() *Notebook {
	retC := C.gtk_notebook_new()
	retGo := NotebookNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendPage is a wrapper around the C function gtk_notebook_append_page.
func (recv *Notebook) AppendPage(child *Widget, tabLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	retC := C.gtk_notebook_append_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label)
	retGo := (int32)(retC)

	return retGo
}

// AppendPageMenu is a wrapper around the C function gtk_notebook_append_page_menu.
func (recv *Notebook) AppendPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	retC := C.gtk_notebook_append_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label)
	retGo := (int32)(retC)

	return retGo
}

// GetCurrentPage is a wrapper around the C function gtk_notebook_get_current_page.
func (recv *Notebook) GetCurrentPage() int32 {
	retC := C.gtk_notebook_get_current_page((*C.GtkNotebook)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMenuLabel is a wrapper around the C function gtk_notebook_get_menu_label.
func (recv *Notebook) GetMenuLabel(child *Widget) *Widget {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_menu_label((*C.GtkNotebook)(recv.native), c_child)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMenuLabelText is a wrapper around the C function gtk_notebook_get_menu_label_text.
func (recv *Notebook) GetMenuLabelText(child *Widget) string {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_menu_label_text((*C.GtkNotebook)(recv.native), c_child)
	retGo := C.GoString(retC)

	return retGo
}

// GetNthPage is a wrapper around the C function gtk_notebook_get_nth_page.
func (recv *Notebook) GetNthPage(pageNum int32) *Widget {
	c_page_num := (C.gint)(pageNum)

	retC := C.gtk_notebook_get_nth_page((*C.GtkNotebook)(recv.native), c_page_num)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScrollable is a wrapper around the C function gtk_notebook_get_scrollable.
func (recv *Notebook) GetScrollable() bool {
	retC := C.gtk_notebook_get_scrollable((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowBorder is a wrapper around the C function gtk_notebook_get_show_border.
func (recv *Notebook) GetShowBorder() bool {
	retC := C.gtk_notebook_get_show_border((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowTabs is a wrapper around the C function gtk_notebook_get_show_tabs.
func (recv *Notebook) GetShowTabs() bool {
	retC := C.gtk_notebook_get_show_tabs((*C.GtkNotebook)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTabLabel is a wrapper around the C function gtk_notebook_get_tab_label.
func (recv *Notebook) GetTabLabel(child *Widget) *Widget {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_label((*C.GtkNotebook)(recv.native), c_child)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTabLabelText is a wrapper around the C function gtk_notebook_get_tab_label_text.
func (recv *Notebook) GetTabLabelText(child *Widget) string {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_get_tab_label_text((*C.GtkNotebook)(recv.native), c_child)
	retGo := C.GoString(retC)

	return retGo
}

// GetTabPos is a wrapper around the C function gtk_notebook_get_tab_pos.
func (recv *Notebook) GetTabPos() PositionType {
	retC := C.gtk_notebook_get_tab_pos((*C.GtkNotebook)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// InsertPage is a wrapper around the C function gtk_notebook_insert_page.
func (recv *Notebook) InsertPage(child *Widget, tabLabel *Widget, position int32) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_notebook_insert_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_position)
	retGo := (int32)(retC)

	return retGo
}

// InsertPageMenu is a wrapper around the C function gtk_notebook_insert_page_menu.
func (recv *Notebook) InsertPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget, position int32) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_notebook_insert_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label, c_position)
	retGo := (int32)(retC)

	return retGo
}

// NextPage is a wrapper around the C function gtk_notebook_next_page.
func (recv *Notebook) NextPage() {
	C.gtk_notebook_next_page((*C.GtkNotebook)(recv.native))

	return
}

// PageNum is a wrapper around the C function gtk_notebook_page_num.
func (recv *Notebook) PageNum(child *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_page_num((*C.GtkNotebook)(recv.native), c_child)
	retGo := (int32)(retC)

	return retGo
}

// PopupDisable is a wrapper around the C function gtk_notebook_popup_disable.
func (recv *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable((*C.GtkNotebook)(recv.native))

	return
}

// PopupEnable is a wrapper around the C function gtk_notebook_popup_enable.
func (recv *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable((*C.GtkNotebook)(recv.native))

	return
}

// PrependPage is a wrapper around the C function gtk_notebook_prepend_page.
func (recv *Notebook) PrependPage(child *Widget, tabLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	retC := C.gtk_notebook_prepend_page((*C.GtkNotebook)(recv.native), c_child, c_tab_label)
	retGo := (int32)(retC)

	return retGo
}

// PrependPageMenu is a wrapper around the C function gtk_notebook_prepend_page_menu.
func (recv *Notebook) PrependPageMenu(child *Widget, tabLabel *Widget, menuLabel *Widget) int32 {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	retC := C.gtk_notebook_prepend_page_menu((*C.GtkNotebook)(recv.native), c_child, c_tab_label, c_menu_label)
	retGo := (int32)(retC)

	return retGo
}

// PrevPage is a wrapper around the C function gtk_notebook_prev_page.
func (recv *Notebook) PrevPage() {
	C.gtk_notebook_prev_page((*C.GtkNotebook)(recv.native))

	return
}

// RemovePage is a wrapper around the C function gtk_notebook_remove_page.
func (recv *Notebook) RemovePage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_notebook_remove_page((*C.GtkNotebook)(recv.native), c_page_num)

	return
}

// ReorderChild is a wrapper around the C function gtk_notebook_reorder_child.
func (recv *Notebook) ReorderChild(child *Widget, position int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_position := (C.gint)(position)

	C.gtk_notebook_reorder_child((*C.GtkNotebook)(recv.native), c_child, c_position)

	return
}

// SetCurrentPage is a wrapper around the C function gtk_notebook_set_current_page.
func (recv *Notebook) SetCurrentPage(pageNum int32) {
	c_page_num := (C.gint)(pageNum)

	C.gtk_notebook_set_current_page((*C.GtkNotebook)(recv.native), c_page_num)

	return
}

// SetMenuLabel is a wrapper around the C function gtk_notebook_set_menu_label.
func (recv *Notebook) SetMenuLabel(child *Widget, menuLabel *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_menu_label := (*C.GtkWidget)(menuLabel.ToC())

	C.gtk_notebook_set_menu_label((*C.GtkNotebook)(recv.native), c_child, c_menu_label)

	return
}

// SetMenuLabelText is a wrapper around the C function gtk_notebook_set_menu_label_text.
func (recv *Notebook) SetMenuLabelText(child *Widget, menuText string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_menu_text := C.CString(menuText)
	defer C.free(unsafe.Pointer(c_menu_text))

	C.gtk_notebook_set_menu_label_text((*C.GtkNotebook)(recv.native), c_child, c_menu_text)

	return
}

// SetScrollable is a wrapper around the C function gtk_notebook_set_scrollable.
func (recv *Notebook) SetScrollable(scrollable bool) {
	c_scrollable :=
		boolToGboolean(scrollable)

	C.gtk_notebook_set_scrollable((*C.GtkNotebook)(recv.native), c_scrollable)

	return
}

// SetShowBorder is a wrapper around the C function gtk_notebook_set_show_border.
func (recv *Notebook) SetShowBorder(showBorder bool) {
	c_show_border :=
		boolToGboolean(showBorder)

	C.gtk_notebook_set_show_border((*C.GtkNotebook)(recv.native), c_show_border)

	return
}

// SetShowTabs is a wrapper around the C function gtk_notebook_set_show_tabs.
func (recv *Notebook) SetShowTabs(showTabs bool) {
	c_show_tabs :=
		boolToGboolean(showTabs)

	C.gtk_notebook_set_show_tabs((*C.GtkNotebook)(recv.native), c_show_tabs)

	return
}

// SetTabLabel is a wrapper around the C function gtk_notebook_set_tab_label.
func (recv *Notebook) SetTabLabel(child *Widget, tabLabel *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_label := (*C.GtkWidget)(tabLabel.ToC())

	C.gtk_notebook_set_tab_label((*C.GtkNotebook)(recv.native), c_child, c_tab_label)

	return
}

// SetTabLabelText is a wrapper around the C function gtk_notebook_set_tab_label_text.
func (recv *Notebook) SetTabLabelText(child *Widget, tabText string) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_tab_text := C.CString(tabText)
	defer C.free(unsafe.Pointer(c_tab_text))

	C.gtk_notebook_set_tab_label_text((*C.GtkNotebook)(recv.native), c_child, c_tab_text)

	return
}

// SetTabPos is a wrapper around the C function gtk_notebook_set_tab_pos.
func (recv *Notebook) SetTabPos(pos PositionType) {
	c_pos := (C.GtkPositionType)(pos)

	C.gtk_notebook_set_tab_pos((*C.GtkNotebook)(recv.native), c_pos)

	return
}

// NotebookAccessible is a wrapper around the C record GtkNotebookAccessible.
type NotebookAccessible struct {
	native *C.GtkNotebookAccessible
	// parent : record
	// priv : record
}

func NotebookAccessibleNewFromC(u unsafe.Pointer) *NotebookAccessible {
	c := (*C.GtkNotebookAccessible)(u)
	if c == nil {
		return nil
	}

	g := &NotebookAccessible{native: c}

	return g
}

func (recv *NotebookAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *NotebookAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *NotebookAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *NotebookAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *NotebookAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to NotebookAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a NotebookAccessible.
func CastToNotebookAccessible(object *gobject.Object) *NotebookAccessible {
	return NotebookAccessibleNewFromC(object.ToC())
}

// NotebookPageAccessible is a wrapper around the C record GtkNotebookPageAccessible.
type NotebookPageAccessible struct {
	native *C.GtkNotebookPageAccessible
	// parent : record
	// priv : record
}

func NotebookPageAccessibleNewFromC(u unsafe.Pointer) *NotebookPageAccessible {
	c := (*C.GtkNotebookPageAccessible)(u)
	if c == nil {
		return nil
	}

	g := &NotebookPageAccessible{native: c}

	return g
}

func (recv *NotebookPageAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *NotebookPageAccessible) Object() *atk.Object {
	return atk.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to NotebookPageAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a NotebookPageAccessible.
func CastToNotebookPageAccessible(object *gobject.Object) *NotebookPageAccessible {
	return NotebookPageAccessibleNewFromC(object.ToC())
}

// NotebookPageAccessibleNew is a wrapper around the C function gtk_notebook_page_accessible_new.
func NotebookPageAccessibleNew(notebook *NotebookAccessible, child *Widget) *NotebookPageAccessible {
	c_notebook := (*C.GtkNotebookAccessible)(notebook.ToC())

	c_child := (*C.GtkWidget)(child.ToC())

	retC := C.gtk_notebook_page_accessible_new(c_notebook, c_child)
	retGo := NotebookPageAccessibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Invalidate is a wrapper around the C function gtk_notebook_page_accessible_invalidate.
func (recv *NotebookPageAccessible) Invalidate() {
	C.gtk_notebook_page_accessible_invalidate((*C.GtkNotebookPageAccessible)(recv.native))

	return
}

// NumerableIcon is a wrapper around the C record GtkNumerableIcon.
type NumerableIcon struct {
	native *C.GtkNumerableIcon
	// parent : record
	// Private : priv
}

func NumerableIconNewFromC(u unsafe.Pointer) *NumerableIcon {
	c := (*C.GtkNumerableIcon)(u)
	if c == nil {
		return nil
	}

	g := &NumerableIcon{native: c}

	return g
}

func (recv *NumerableIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EmblemedIcon upcasts to *EmblemedIcon
func (recv *NumerableIcon) EmblemedIcon() *gio.EmblemedIcon {
	return gio.EmblemedIconNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *NumerableIcon) Object() *gobject.Object {
	return recv.EmblemedIcon().Object()
}

// CastToWidget down casts any arbitary Object to NumerableIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a NumerableIcon.
func CastToNumerableIcon(object *gobject.Object) *NumerableIcon {
	return NumerableIconNewFromC(object.ToC())
}

// OffscreenWindow is a wrapper around the C record GtkOffscreenWindow.
type OffscreenWindow struct {
	native *C.GtkOffscreenWindow
	// parent_object : record
}

func OffscreenWindowNewFromC(u unsafe.Pointer) *OffscreenWindow {
	c := (*C.GtkOffscreenWindow)(u)
	if c == nil {
		return nil
	}

	g := &OffscreenWindow{native: c}

	return g
}

func (recv *OffscreenWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Window upcasts to *Window
func (recv *OffscreenWindow) Window() *Window {
	return WindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *OffscreenWindow) Bin() *Bin {
	return recv.Window().Bin()
}

// Container upcasts to *Container
func (recv *OffscreenWindow) Container() *Container {
	return recv.Window().Container()
}

// Widget upcasts to *Widget
func (recv *OffscreenWindow) Widget() *Widget {
	return recv.Window().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *OffscreenWindow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Window().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *OffscreenWindow) Object() *gobject.Object {
	return recv.Window().Object()
}

// CastToWidget down casts any arbitary Object to OffscreenWindow.
// Exercise care, as this is a potentially dangerous function if the Object is not a OffscreenWindow.
func CastToOffscreenWindow(object *gobject.Object) *OffscreenWindow {
	return OffscreenWindowNewFromC(object.ToC())
}

// Overlay is a wrapper around the C record GtkOverlay.
type Overlay struct {
	native *C.GtkOverlay
	// parent : record
	// priv : record
}

func OverlayNewFromC(u unsafe.Pointer) *Overlay {
	c := (*C.GtkOverlay)(u)
	if c == nil {
		return nil
	}

	g := &Overlay{native: c}

	return g
}

func (recv *Overlay) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Overlay) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Overlay) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Overlay) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Overlay) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Overlay) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Overlay.
// Exercise care, as this is a potentially dangerous function if the Object is not a Overlay.
func CastToOverlay(object *gobject.Object) *Overlay {
	return OverlayNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

// Unsupported : gtk_pad_controller_new : unsupported parameter group : no type generator for Gio.ActionGroup, GActionGroup*

// PageSetup is a wrapper around the C record GtkPageSetup.
type PageSetup struct {
	native *C.GtkPageSetup
}

func PageSetupNewFromC(u unsafe.Pointer) *PageSetup {
	c := (*C.GtkPageSetup)(u)
	if c == nil {
		return nil
	}

	g := &PageSetup{native: c}

	return g
}

func (recv *PageSetup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PageSetup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PageSetup.
// Exercise care, as this is a potentially dangerous function if the Object is not a PageSetup.
func CastToPageSetup(object *gobject.Object) *PageSetup {
	return PageSetupNewFromC(object.ToC())
}

// Unsupported : gtk_page_setup_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// Paned is a wrapper around the C record GtkPaned.
type Paned struct {
	native *C.GtkPaned
	// container : record
	// Private : priv
}

func PanedNewFromC(u unsafe.Pointer) *Paned {
	c := (*C.GtkPaned)(u)
	if c == nil {
		return nil
	}

	g := &Paned{native: c}

	return g
}

func (recv *Paned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Paned) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Paned) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Paned) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Paned) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Paned.
// Exercise care, as this is a potentially dangerous function if the Object is not a Paned.
func CastToPaned(object *gobject.Object) *Paned {
	return PanedNewFromC(object.ToC())
}

// Add1 is a wrapper around the C function gtk_paned_add1.
func (recv *Paned) Add1(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_paned_add1((*C.GtkPaned)(recv.native), c_child)

	return
}

// Add2 is a wrapper around the C function gtk_paned_add2.
func (recv *Paned) Add2(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_paned_add2((*C.GtkPaned)(recv.native), c_child)

	return
}

// GetPosition is a wrapper around the C function gtk_paned_get_position.
func (recv *Paned) GetPosition() int32 {
	retC := C.gtk_paned_get_position((*C.GtkPaned)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Pack1 is a wrapper around the C function gtk_paned_pack1.
func (recv *Paned) Pack1(child *Widget, resize bool, shrink bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_resize :=
		boolToGboolean(resize)

	c_shrink :=
		boolToGboolean(shrink)

	C.gtk_paned_pack1((*C.GtkPaned)(recv.native), c_child, c_resize, c_shrink)

	return
}

// Pack2 is a wrapper around the C function gtk_paned_pack2.
func (recv *Paned) Pack2(child *Widget, resize bool, shrink bool) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_resize :=
		boolToGboolean(resize)

	c_shrink :=
		boolToGboolean(shrink)

	C.gtk_paned_pack2((*C.GtkPaned)(recv.native), c_child, c_resize, c_shrink)

	return
}

// SetPosition is a wrapper around the C function gtk_paned_set_position.
func (recv *Paned) SetPosition(position int32) {
	c_position := (C.gint)(position)

	C.gtk_paned_set_position((*C.GtkPaned)(recv.native), c_position)

	return
}

// PanedAccessible is a wrapper around the C record GtkPanedAccessible.
type PanedAccessible struct {
	native *C.GtkPanedAccessible
	// parent : record
	// priv : record
}

func PanedAccessibleNewFromC(u unsafe.Pointer) *PanedAccessible {
	c := (*C.GtkPanedAccessible)(u)
	if c == nil {
		return nil
	}

	g := &PanedAccessible{native: c}

	return g
}

func (recv *PanedAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *PanedAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *PanedAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *PanedAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *PanedAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to PanedAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a PanedAccessible.
func CastToPanedAccessible(object *gobject.Object) *PanedAccessible {
	return PanedAccessibleNewFromC(object.ToC())
}

// PlacesSidebar is a wrapper around the C record GtkPlacesSidebar.
type PlacesSidebar struct {
	native *C.GtkPlacesSidebar
}

func PlacesSidebarNewFromC(u unsafe.Pointer) *PlacesSidebar {
	c := (*C.GtkPlacesSidebar)(u)
	if c == nil {
		return nil
	}

	g := &PlacesSidebar{native: c}

	return g
}

func (recv *PlacesSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScrolledWindow upcasts to *ScrolledWindow
func (recv *PlacesSidebar) ScrolledWindow() *ScrolledWindow {
	return ScrolledWindowNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *PlacesSidebar) Bin() *Bin {
	return recv.ScrolledWindow().Bin()
}

// Container upcasts to *Container
func (recv *PlacesSidebar) Container() *Container {
	return recv.ScrolledWindow().Container()
}

// Widget upcasts to *Widget
func (recv *PlacesSidebar) Widget() *Widget {
	return recv.ScrolledWindow().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *PlacesSidebar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ScrolledWindow().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *PlacesSidebar) Object() *gobject.Object {
	return recv.ScrolledWindow().Object()
}

// CastToWidget down casts any arbitary Object to PlacesSidebar.
// Exercise care, as this is a potentially dangerous function if the Object is not a PlacesSidebar.
func CastToPlacesSidebar(object *gobject.Object) *PlacesSidebar {
	return PlacesSidebarNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter dest_file : no type generator for Gio.File,

// Unsupported signal : unsupported parameter location : no type generator for Gio.File,

// Unsupported signal : unsupported parameter selected_item : no type generator for Gio.File,

var signalShowConnectToServerId int
var signalShowConnectToServerMap = make(map[int]PlacesSidebarSignalShowConnectToServerCallback)
var signalShowConnectToServerLock sync.Mutex

// PlacesSidebarSignalShowConnectToServerCallback is a callback function for a 'show-connect-to-server' signal emitted from a PlacesSidebar.
type PlacesSidebarSignalShowConnectToServerCallback func()

func PlacesSidebar_showConnectToServerHandler() {}

// GetShowConnectToServer is a wrapper around the C function gtk_places_sidebar_get_show_connect_to_server.
func (recv *PlacesSidebar) GetShowConnectToServer() bool {
	retC := C.gtk_places_sidebar_get_show_connect_to_server((*C.GtkPlacesSidebar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Blacklisted : GtkPlug

// Popover is a wrapper around the C record GtkPopover.
type Popover struct {
	native *C.GtkPopover
	// parent_instance : record
	// Private : priv
}

func PopoverNewFromC(u unsafe.Pointer) *Popover {
	c := (*C.GtkPopover)(u)
	if c == nil {
		return nil
	}

	g := &Popover{native: c}

	return g
}

func (recv *Popover) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Popover) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Popover) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Popover) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Popover) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Popover) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Popover.
// Exercise care, as this is a potentially dangerous function if the Object is not a Popover.
func CastToPopover(object *gobject.Object) *Popover {
	return PopoverNewFromC(object.ToC())
}

var signalClosedId int
var signalClosedMap = make(map[int]PopoverSignalClosedCallback)
var signalClosedLock sync.Mutex

// PopoverSignalClosedCallback is a callback function for a 'closed' signal emitted from a Popover.
type PopoverSignalClosedCallback func()

func Popover_closedHandler() {}

// Unsupported : gtk_popover_get_pointing_to : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetPosition is a wrapper around the C function gtk_popover_get_position.
func (recv *Popover) GetPosition() PositionType {
	retC := C.gtk_popover_get_position((*C.GtkPopover)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// PopoverAccessible is a wrapper around the C record GtkPopoverAccessible.
type PopoverAccessible struct {
	native *C.GtkPopoverAccessible
	// parent : record
}

func PopoverAccessibleNewFromC(u unsafe.Pointer) *PopoverAccessible {
	c := (*C.GtkPopoverAccessible)(u)
	if c == nil {
		return nil
	}

	g := &PopoverAccessible{native: c}

	return g
}

func (recv *PopoverAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *PopoverAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *PopoverAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *PopoverAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *PopoverAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to PopoverAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a PopoverAccessible.
func CastToPopoverAccessible(object *gobject.Object) *PopoverAccessible {
	return PopoverAccessibleNewFromC(object.ToC())
}

// PopoverMenu is a wrapper around the C record GtkPopoverMenu.
type PopoverMenu struct {
	native *C.GtkPopoverMenu
}

func PopoverMenuNewFromC(u unsafe.Pointer) *PopoverMenu {
	c := (*C.GtkPopoverMenu)(u)
	if c == nil {
		return nil
	}

	g := &PopoverMenu{native: c}

	return g
}

func (recv *PopoverMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Popover upcasts to *Popover
func (recv *PopoverMenu) Popover() *Popover {
	return PopoverNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *PopoverMenu) Bin() *Bin {
	return recv.Popover().Bin()
}

// Container upcasts to *Container
func (recv *PopoverMenu) Container() *Container {
	return recv.Popover().Container()
}

// Widget upcasts to *Widget
func (recv *PopoverMenu) Widget() *Widget {
	return recv.Popover().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *PopoverMenu) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Popover().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *PopoverMenu) Object() *gobject.Object {
	return recv.Popover().Object()
}

// CastToWidget down casts any arbitary Object to PopoverMenu.
// Exercise care, as this is a potentially dangerous function if the Object is not a PopoverMenu.
func CastToPopoverMenu(object *gobject.Object) *PopoverMenu {
	return PopoverMenuNewFromC(object.ToC())
}

// PrintContext is a wrapper around the C record GtkPrintContext.
type PrintContext struct {
	native *C.GtkPrintContext
}

func PrintContextNewFromC(u unsafe.Pointer) *PrintContext {
	c := (*C.GtkPrintContext)(u)
	if c == nil {
		return nil
	}

	g := &PrintContext{native: c}

	return g
}

func (recv *PrintContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PrintContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PrintContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintContext.
func CastToPrintContext(object *gobject.Object) *PrintContext {
	return PrintContextNewFromC(object.ToC())
}

// PrintOperation is a wrapper around the C record GtkPrintOperation.
type PrintOperation struct {
	native *C.GtkPrintOperation
	// parent_instance : record
	// Private : priv
}

func PrintOperationNewFromC(u unsafe.Pointer) *PrintOperation {
	c := (*C.GtkPrintOperation)(u)
	if c == nil {
		return nil
	}

	g := &PrintOperation{native: c}

	return g
}

func (recv *PrintOperation) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PrintOperation) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PrintOperation.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintOperation.
func CastToPrintOperation(object *gobject.Object) *PrintOperation {
	return PrintOperationNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter preview : no type generator for PrintOperationPreview,

// PrintSettings is a wrapper around the C record GtkPrintSettings.
type PrintSettings struct {
	native *C.GtkPrintSettings
}

func PrintSettingsNewFromC(u unsafe.Pointer) *PrintSettings {
	c := (*C.GtkPrintSettings)(u)
	if c == nil {
		return nil
	}

	g := &PrintSettings{native: c}

	return g
}

func (recv *PrintSettings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *PrintSettings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to PrintSettings.
// Exercise care, as this is a potentially dangerous function if the Object is not a PrintSettings.
func CastToPrintSettings(object *gobject.Object) *PrintSettings {
	return PrintSettingsNewFromC(object.ToC())
}

// Unsupported : gtk_print_settings_new_from_gvariant : unsupported parameter variant : Blacklisted record : GVariant

// ProgressBar is a wrapper around the C record GtkProgressBar.
type ProgressBar struct {
	native *C.GtkProgressBar
	// parent : record
	// Private : priv
}

func ProgressBarNewFromC(u unsafe.Pointer) *ProgressBar {
	c := (*C.GtkProgressBar)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBar{native: c}

	return g
}

func (recv *ProgressBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *ProgressBar) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ProgressBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ProgressBar) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to ProgressBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProgressBar.
func CastToProgressBar(object *gobject.Object) *ProgressBar {
	return ProgressBarNewFromC(object.ToC())
}

// ProgressBarNew is a wrapper around the C function gtk_progress_bar_new.
func ProgressBarNew() *ProgressBar {
	retC := C.gtk_progress_bar_new()
	retGo := ProgressBarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFraction is a wrapper around the C function gtk_progress_bar_get_fraction.
func (recv *ProgressBar) GetFraction() float64 {
	retC := C.gtk_progress_bar_get_fraction((*C.GtkProgressBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetInverted is a wrapper around the C function gtk_progress_bar_get_inverted.
func (recv *ProgressBar) GetInverted() bool {
	retC := C.gtk_progress_bar_get_inverted((*C.GtkProgressBar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPulseStep is a wrapper around the C function gtk_progress_bar_get_pulse_step.
func (recv *ProgressBar) GetPulseStep() float64 {
	retC := C.gtk_progress_bar_get_pulse_step((*C.GtkProgressBar)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetText is a wrapper around the C function gtk_progress_bar_get_text.
func (recv *ProgressBar) GetText() string {
	retC := C.gtk_progress_bar_get_text((*C.GtkProgressBar)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// Pulse is a wrapper around the C function gtk_progress_bar_pulse.
func (recv *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse((*C.GtkProgressBar)(recv.native))

	return
}

// SetFraction is a wrapper around the C function gtk_progress_bar_set_fraction.
func (recv *ProgressBar) SetFraction(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_progress_bar_set_fraction((*C.GtkProgressBar)(recv.native), c_fraction)

	return
}

// SetInverted is a wrapper around the C function gtk_progress_bar_set_inverted.
func (recv *ProgressBar) SetInverted(inverted bool) {
	c_inverted :=
		boolToGboolean(inverted)

	C.gtk_progress_bar_set_inverted((*C.GtkProgressBar)(recv.native), c_inverted)

	return
}

// SetPulseStep is a wrapper around the C function gtk_progress_bar_set_pulse_step.
func (recv *ProgressBar) SetPulseStep(fraction float64) {
	c_fraction := (C.gdouble)(fraction)

	C.gtk_progress_bar_set_pulse_step((*C.GtkProgressBar)(recv.native), c_fraction)

	return
}

// SetText is a wrapper around the C function gtk_progress_bar_set_text.
func (recv *ProgressBar) SetText(text string) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_progress_bar_set_text((*C.GtkProgressBar)(recv.native), c_text)

	return
}

// ProgressBarAccessible is a wrapper around the C record GtkProgressBarAccessible.
type ProgressBarAccessible struct {
	native *C.GtkProgressBarAccessible
	// parent : record
	// priv : record
}

func ProgressBarAccessibleNewFromC(u unsafe.Pointer) *ProgressBarAccessible {
	c := (*C.GtkProgressBarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ProgressBarAccessible{native: c}

	return g
}

func (recv *ProgressBarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ProgressBarAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *ProgressBarAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ProgressBarAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ProgressBarAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ProgressBarAccessible.
func CastToProgressBarAccessible(object *gobject.Object) *ProgressBarAccessible {
	return ProgressBarAccessibleNewFromC(object.ToC())
}

// RadioAction is a wrapper around the C record GtkRadioAction.
type RadioAction struct {
	native *C.GtkRadioAction
	// parent : record
	// Private : private_data
}

func RadioActionNewFromC(u unsafe.Pointer) *RadioAction {
	c := (*C.GtkRadioAction)(u)
	if c == nil {
		return nil
	}

	g := &RadioAction{native: c}

	return g
}

func (recv *RadioAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleAction upcasts to *ToggleAction
func (recv *RadioAction) ToggleAction() *ToggleAction {
	return ToggleActionNewFromC(unsafe.Pointer(recv.native))
}

// Action upcasts to *Action
func (recv *RadioAction) Action() *Action {
	return recv.ToggleAction().Action()
}

// Object upcasts to *Object
func (recv *RadioAction) Object() *gobject.Object {
	return recv.ToggleAction().Object()
}

// CastToWidget down casts any arbitary Object to RadioAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioAction.
func CastToRadioAction(object *gobject.Object) *RadioAction {
	return RadioActionNewFromC(object.ToC())
}

// RadioButton is a wrapper around the C record GtkRadioButton.
type RadioButton struct {
	native *C.GtkRadioButton
	// check_button : record
	// Private : priv
}

func RadioButtonNewFromC(u unsafe.Pointer) *RadioButton {
	c := (*C.GtkRadioButton)(u)
	if c == nil {
		return nil
	}

	g := &RadioButton{native: c}

	return g
}

func (recv *RadioButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckButton upcasts to *CheckButton
func (recv *RadioButton) CheckButton() *CheckButton {
	return CheckButtonNewFromC(unsafe.Pointer(recv.native))
}

// ToggleButton upcasts to *ToggleButton
func (recv *RadioButton) ToggleButton() *ToggleButton {
	return recv.CheckButton().ToggleButton()
}

// Button upcasts to *Button
func (recv *RadioButton) Button() *Button {
	return recv.CheckButton().Button()
}

// Bin upcasts to *Bin
func (recv *RadioButton) Bin() *Bin {
	return recv.CheckButton().Bin()
}

// Container upcasts to *Container
func (recv *RadioButton) Container() *Container {
	return recv.CheckButton().Container()
}

// Widget upcasts to *Widget
func (recv *RadioButton) Widget() *Widget {
	return recv.CheckButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RadioButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CheckButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RadioButton) Object() *gobject.Object {
	return recv.CheckButton().Object()
}

// CastToWidget down casts any arbitary Object to RadioButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioButton.
func CastToRadioButton(object *gobject.Object) *RadioButton {
	return RadioButtonNewFromC(object.ToC())
}

// RadioButtonNew is a wrapper around the C function gtk_radio_button_new.
func RadioButtonNew(group *glib.SList) *RadioButton {
	c_group := (*C.GSList)(group.ToC())

	retC := C.gtk_radio_button_new(c_group)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewFromWidget is a wrapper around the C function gtk_radio_button_new_from_widget.
func RadioButtonNewFromWidget(radioGroupMember *RadioButton) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	retC := C.gtk_radio_button_new_from_widget(c_radio_group_member)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabel is a wrapper around the C function gtk_radio_button_new_with_label.
func RadioButtonNewWithLabel(group *glib.SList, label string) *RadioButton {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label(c_group, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithLabelFromWidget is a wrapper around the C function gtk_radio_button_new_with_label_from_widget.
func RadioButtonNewWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_label_from_widget(c_radio_group_member, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonic is a wrapper around the C function gtk_radio_button_new_with_mnemonic.
func RadioButtonNewWithMnemonic(group *glib.SList, label string) *RadioButton {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic(c_group, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioButtonNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_button_new_with_mnemonic_from_widget.
func RadioButtonNewWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	c_radio_group_member := (*C.GtkRadioButton)(radioGroupMember.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_button_new_with_mnemonic_from_widget(c_radio_group_member, c_label)
	retGo := RadioButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_button_get_group.
func (recv *RadioButton) GetGroup() *glib.SList {
	retC := C.gtk_radio_button_get_group((*C.GtkRadioButton)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_button_set_group.
func (recv *RadioButton) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(group.ToC())

	C.gtk_radio_button_set_group((*C.GtkRadioButton)(recv.native), c_group)

	return
}

// RadioButtonAccessible is a wrapper around the C record GtkRadioButtonAccessible.
type RadioButtonAccessible struct {
	native *C.GtkRadioButtonAccessible
	// parent : record
	// priv : record
}

func RadioButtonAccessibleNewFromC(u unsafe.Pointer) *RadioButtonAccessible {
	c := (*C.GtkRadioButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RadioButtonAccessible{native: c}

	return g
}

func (recv *RadioButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleButtonAccessible upcasts to *ToggleButtonAccessible
func (recv *RadioButtonAccessible) ToggleButtonAccessible() *ToggleButtonAccessible {
	return ToggleButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *RadioButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return recv.ToggleButtonAccessible().ButtonAccessible()
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *RadioButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ToggleButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *RadioButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ToggleButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *RadioButtonAccessible) Accessible() *Accessible {
	return recv.ToggleButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *RadioButtonAccessible) Object() *atk.Object {
	return recv.ToggleButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to RadioButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioButtonAccessible.
func CastToRadioButtonAccessible(object *gobject.Object) *RadioButtonAccessible {
	return RadioButtonAccessibleNewFromC(object.ToC())
}

// RadioMenuItem is a wrapper around the C record GtkRadioMenuItem.
type RadioMenuItem struct {
	native *C.GtkRadioMenuItem
	// check_menu_item : record
	// Private : priv
}

func RadioMenuItemNewFromC(u unsafe.Pointer) *RadioMenuItem {
	c := (*C.GtkRadioMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItem{native: c}

	return g
}

func (recv *RadioMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItem upcasts to *CheckMenuItem
func (recv *RadioMenuItem) CheckMenuItem() *CheckMenuItem {
	return CheckMenuItemNewFromC(unsafe.Pointer(recv.native))
}

// MenuItem upcasts to *MenuItem
func (recv *RadioMenuItem) MenuItem() *MenuItem {
	return recv.CheckMenuItem().MenuItem()
}

// Bin upcasts to *Bin
func (recv *RadioMenuItem) Bin() *Bin {
	return recv.CheckMenuItem().Bin()
}

// Container upcasts to *Container
func (recv *RadioMenuItem) Container() *Container {
	return recv.CheckMenuItem().Container()
}

// Widget upcasts to *Widget
func (recv *RadioMenuItem) Widget() *Widget {
	return recv.CheckMenuItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RadioMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.CheckMenuItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RadioMenuItem) Object() *gobject.Object {
	return recv.CheckMenuItem().Object()
}

// CastToWidget down casts any arbitary Object to RadioMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioMenuItem.
func CastToRadioMenuItem(object *gobject.Object) *RadioMenuItem {
	return RadioMenuItemNewFromC(object.ToC())
}

var signalGroupChangedId int
var signalGroupChangedMap = make(map[int]RadioMenuItemSignalGroupChangedCallback)
var signalGroupChangedLock sync.Mutex

// RadioMenuItemSignalGroupChangedCallback is a callback function for a 'group-changed' signal emitted from a RadioMenuItem.
type RadioMenuItemSignalGroupChangedCallback func()

func RadioMenuItem_groupChangedHandler() {}

// RadioMenuItemNew is a wrapper around the C function gtk_radio_menu_item_new.
func RadioMenuItemNew(group *glib.SList) *RadioMenuItem {
	c_group := (*C.GSList)(group.ToC())

	retC := C.gtk_radio_menu_item_new(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabel is a wrapper around the C function gtk_radio_menu_item_new_with_label.
func RadioMenuItemNewWithLabel(group *glib.SList, label string) *RadioMenuItem {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonic is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic.
func RadioMenuItemNewWithMnemonic(group *glib.SList, label string) *RadioMenuItem {
	c_group := (*C.GSList)(group.ToC())

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_menu_item_get_group.
func (recv *RadioMenuItem) GetGroup() *glib.SList {
	retC := C.gtk_radio_menu_item_get_group((*C.GtkRadioMenuItem)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_menu_item_set_group.
func (recv *RadioMenuItem) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(group.ToC())

	C.gtk_radio_menu_item_set_group((*C.GtkRadioMenuItem)(recv.native), c_group)

	return
}

// RadioMenuItemAccessible is a wrapper around the C record GtkRadioMenuItemAccessible.
type RadioMenuItemAccessible struct {
	native *C.GtkRadioMenuItemAccessible
	// parent : record
	// priv : record
}

func RadioMenuItemAccessibleNewFromC(u unsafe.Pointer) *RadioMenuItemAccessible {
	c := (*C.GtkRadioMenuItemAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RadioMenuItemAccessible{native: c}

	return g
}

func (recv *RadioMenuItemAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CheckMenuItemAccessible upcasts to *CheckMenuItemAccessible
func (recv *RadioMenuItemAccessible) CheckMenuItemAccessible() *CheckMenuItemAccessible {
	return CheckMenuItemAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// MenuItemAccessible upcasts to *MenuItemAccessible
func (recv *RadioMenuItemAccessible) MenuItemAccessible() *MenuItemAccessible {
	return recv.CheckMenuItemAccessible().MenuItemAccessible()
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *RadioMenuItemAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.CheckMenuItemAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *RadioMenuItemAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.CheckMenuItemAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *RadioMenuItemAccessible) Accessible() *Accessible {
	return recv.CheckMenuItemAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *RadioMenuItemAccessible) Object() *atk.Object {
	return recv.CheckMenuItemAccessible().Object()
}

// CastToWidget down casts any arbitary Object to RadioMenuItemAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioMenuItemAccessible.
func CastToRadioMenuItemAccessible(object *gobject.Object) *RadioMenuItemAccessible {
	return RadioMenuItemAccessibleNewFromC(object.ToC())
}

// RadioToolButton is a wrapper around the C record GtkRadioToolButton.
type RadioToolButton struct {
	native *C.GtkRadioToolButton
	// parent : record
}

func RadioToolButtonNewFromC(u unsafe.Pointer) *RadioToolButton {
	c := (*C.GtkRadioToolButton)(u)
	if c == nil {
		return nil
	}

	g := &RadioToolButton{native: c}

	return g
}

func (recv *RadioToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToggleToolButton upcasts to *ToggleToolButton
func (recv *RadioToolButton) ToggleToolButton() *ToggleToolButton {
	return ToggleToolButtonNewFromC(unsafe.Pointer(recv.native))
}

// ToolButton upcasts to *ToolButton
func (recv *RadioToolButton) ToolButton() *ToolButton {
	return recv.ToggleToolButton().ToolButton()
}

// ToolItem upcasts to *ToolItem
func (recv *RadioToolButton) ToolItem() *ToolItem {
	return recv.ToggleToolButton().ToolItem()
}

// Bin upcasts to *Bin
func (recv *RadioToolButton) Bin() *Bin {
	return recv.ToggleToolButton().Bin()
}

// Container upcasts to *Container
func (recv *RadioToolButton) Container() *Container {
	return recv.ToggleToolButton().Container()
}

// Widget upcasts to *Widget
func (recv *RadioToolButton) Widget() *Widget {
	return recv.ToggleToolButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RadioToolButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToggleToolButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RadioToolButton) Object() *gobject.Object {
	return recv.ToggleToolButton().Object()
}

// CastToWidget down casts any arbitary Object to RadioToolButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a RadioToolButton.
func CastToRadioToolButton(object *gobject.Object) *RadioToolButton {
	return RadioToolButtonNewFromC(object.ToC())
}

// Range is a wrapper around the C record GtkRange.
type Range struct {
	native *C.GtkRange
	// widget : record
	// priv : record
}

func RangeNewFromC(u unsafe.Pointer) *Range {
	c := (*C.GtkRange)(u)
	if c == nil {
		return nil
	}

	g := &Range{native: c}

	return g
}

func (recv *Range) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Range) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Range) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Range) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Range.
// Exercise care, as this is a potentially dangerous function if the Object is not a Range.
func CastToRange(object *gobject.Object) *Range {
	return RangeNewFromC(object.ToC())
}

var signalAdjustBoundsId int
var signalAdjustBoundsMap = make(map[int]RangeSignalAdjustBoundsCallback)
var signalAdjustBoundsLock sync.Mutex

// RangeSignalAdjustBoundsCallback is a callback function for a 'adjust-bounds' signal emitted from a Range.
type RangeSignalAdjustBoundsCallback func(value float64)

func Range_adjustBoundsHandler() {}

var signalMoveSliderId int
var signalMoveSliderMap = make(map[int]RangeSignalMoveSliderCallback)
var signalMoveSliderLock sync.Mutex

// RangeSignalMoveSliderCallback is a callback function for a 'move-slider' signal emitted from a Range.
type RangeSignalMoveSliderCallback func(step ScrollType)

func Range_moveSliderHandler() {}

var signalValueChangedId int
var signalValueChangedMap = make(map[int]RangeSignalValueChangedCallback)
var signalValueChangedLock sync.Mutex

// RangeSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a Range.
type RangeSignalValueChangedCallback func()

func Range_valueChangedHandler() {}

// GetAdjustment is a wrapper around the C function gtk_range_get_adjustment.
func (recv *Range) GetAdjustment() *Adjustment {
	retC := C.gtk_range_get_adjustment((*C.GtkRange)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInverted is a wrapper around the C function gtk_range_get_inverted.
func (recv *Range) GetInverted() bool {
	retC := C.gtk_range_get_inverted((*C.GtkRange)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValue is a wrapper around the C function gtk_range_get_value.
func (recv *Range) GetValue() float64 {
	retC := C.gtk_range_get_value((*C.GtkRange)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// SetAdjustment is a wrapper around the C function gtk_range_set_adjustment.
func (recv *Range) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_range_set_adjustment((*C.GtkRange)(recv.native), c_adjustment)

	return
}

// SetIncrements is a wrapper around the C function gtk_range_set_increments.
func (recv *Range) SetIncrements(step float64, page float64) {
	c_step := (C.gdouble)(step)

	c_page := (C.gdouble)(page)

	C.gtk_range_set_increments((*C.GtkRange)(recv.native), c_step, c_page)

	return
}

// SetInverted is a wrapper around the C function gtk_range_set_inverted.
func (recv *Range) SetInverted(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_range_set_inverted((*C.GtkRange)(recv.native), c_setting)

	return
}

// SetRange is a wrapper around the C function gtk_range_set_range.
func (recv *Range) SetRange(min float64, max float64) {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	C.gtk_range_set_range((*C.GtkRange)(recv.native), c_min, c_max)

	return
}

// SetValue is a wrapper around the C function gtk_range_set_value.
func (recv *Range) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_range_set_value((*C.GtkRange)(recv.native), c_value)

	return
}

// RangeAccessible is a wrapper around the C record GtkRangeAccessible.
type RangeAccessible struct {
	native *C.GtkRangeAccessible
	// parent : record
	// priv : record
}

func RangeAccessibleNewFromC(u unsafe.Pointer) *RangeAccessible {
	c := (*C.GtkRangeAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RangeAccessible{native: c}

	return g
}

func (recv *RangeAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *RangeAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *RangeAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *RangeAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to RangeAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a RangeAccessible.
func CastToRangeAccessible(object *gobject.Object) *RangeAccessible {
	return RangeAccessibleNewFromC(object.ToC())
}

// RcStyle is a wrapper around the C record GtkRcStyle.
type RcStyle struct {
	native *C.GtkRcStyle
	// parent_instance : record
	Name string
	// no type for bg_pixmap_name
	// font_desc : record
	// no type for color_flags
	// no type for fg
	// no type for bg
	// no type for text
	// no type for base
	Xthickness int32
	Ythickness int32
	// Private : rc_properties
	// Private : rc_style_lists
	// Private : icon_factories
	// Private : engine_specified
}

func RcStyleNewFromC(u unsafe.Pointer) *RcStyle {
	c := (*C.GtkRcStyle)(u)
	if c == nil {
		return nil
	}

	g := &RcStyle{
		Name:       C.GoString(c.name),
		Xthickness: (int32)(c.xthickness),
		Ythickness: (int32)(c.ythickness),
		native:     c,
	}

	return g
}

func (recv *RcStyle) ToC() unsafe.Pointer {
	recv.native.name =
		C.CString(recv.Name)
	recv.native.xthickness =
		(C.gint)(recv.Xthickness)
	recv.native.ythickness =
		(C.gint)(recv.Ythickness)

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *RcStyle) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to RcStyle.
// Exercise care, as this is a potentially dangerous function if the Object is not a RcStyle.
func CastToRcStyle(object *gobject.Object) *RcStyle {
	return RcStyleNewFromC(object.ToC())
}

// RcStyleNew is a wrapper around the C function gtk_rc_style_new.
func RcStyleNew() *RcStyle {
	retC := C.gtk_rc_style_new()
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_rc_style_copy.
func (recv *RcStyle) Copy() *RcStyle {
	retC := C.gtk_rc_style_copy((*C.GtkRcStyle)(recv.native))
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RecentAction is a wrapper around the C record GtkRecentAction.
type RecentAction struct {
	native *C.GtkRecentAction
	// parent_instance : record
	// Private : priv
}

func RecentActionNewFromC(u unsafe.Pointer) *RecentAction {
	c := (*C.GtkRecentAction)(u)
	if c == nil {
		return nil
	}

	g := &RecentAction{native: c}

	return g
}

func (recv *RecentAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action upcasts to *Action
func (recv *RecentAction) Action() *Action {
	return ActionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *RecentAction) Object() *gobject.Object {
	return recv.Action().Object()
}

// CastToWidget down casts any arbitary Object to RecentAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentAction.
func CastToRecentAction(object *gobject.Object) *RecentAction {
	return RecentActionNewFromC(object.ToC())
}

// RecentChooserDialog is a wrapper around the C record GtkRecentChooserDialog.
type RecentChooserDialog struct {
	native *C.GtkRecentChooserDialog
	// parent_instance : record
	// Private : priv
}

func RecentChooserDialogNewFromC(u unsafe.Pointer) *RecentChooserDialog {
	c := (*C.GtkRecentChooserDialog)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserDialog{native: c}

	return g
}

func (recv *RecentChooserDialog) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Dialog upcasts to *Dialog
func (recv *RecentChooserDialog) Dialog() *Dialog {
	return DialogNewFromC(unsafe.Pointer(recv.native))
}

// Window upcasts to *Window
func (recv *RecentChooserDialog) Window() *Window {
	return recv.Dialog().Window()
}

// Bin upcasts to *Bin
func (recv *RecentChooserDialog) Bin() *Bin {
	return recv.Dialog().Bin()
}

// Container upcasts to *Container
func (recv *RecentChooserDialog) Container() *Container {
	return recv.Dialog().Container()
}

// Widget upcasts to *Widget
func (recv *RecentChooserDialog) Widget() *Widget {
	return recv.Dialog().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RecentChooserDialog) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Dialog().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RecentChooserDialog) Object() *gobject.Object {
	return recv.Dialog().Object()
}

// CastToWidget down casts any arbitary Object to RecentChooserDialog.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentChooserDialog.
func CastToRecentChooserDialog(object *gobject.Object) *RecentChooserDialog {
	return RecentChooserDialogNewFromC(object.ToC())
}

// Unsupported : gtk_recent_chooser_dialog_new : unsupported parameter ... : varargs

// Unsupported : gtk_recent_chooser_dialog_new_for_manager : unsupported parameter ... : varargs

// RecentChooserMenu is a wrapper around the C record GtkRecentChooserMenu.
type RecentChooserMenu struct {
	native *C.GtkRecentChooserMenu
	// parent_instance : record
	// Private : priv
}

func RecentChooserMenuNewFromC(u unsafe.Pointer) *RecentChooserMenu {
	c := (*C.GtkRecentChooserMenu)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserMenu{native: c}

	return g
}

func (recv *RecentChooserMenu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Menu upcasts to *Menu
func (recv *RecentChooserMenu) Menu() *Menu {
	return MenuNewFromC(unsafe.Pointer(recv.native))
}

// MenuShell upcasts to *MenuShell
func (recv *RecentChooserMenu) MenuShell() *MenuShell {
	return recv.Menu().MenuShell()
}

// Container upcasts to *Container
func (recv *RecentChooserMenu) Container() *Container {
	return recv.Menu().Container()
}

// Widget upcasts to *Widget
func (recv *RecentChooserMenu) Widget() *Widget {
	return recv.Menu().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RecentChooserMenu) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Menu().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RecentChooserMenu) Object() *gobject.Object {
	return recv.Menu().Object()
}

// CastToWidget down casts any arbitary Object to RecentChooserMenu.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentChooserMenu.
func CastToRecentChooserMenu(object *gobject.Object) *RecentChooserMenu {
	return RecentChooserMenuNewFromC(object.ToC())
}

// RecentChooserWidget is a wrapper around the C record GtkRecentChooserWidget.
type RecentChooserWidget struct {
	native *C.GtkRecentChooserWidget
	// parent_instance : record
	// Private : priv
}

func RecentChooserWidgetNewFromC(u unsafe.Pointer) *RecentChooserWidget {
	c := (*C.GtkRecentChooserWidget)(u)
	if c == nil {
		return nil
	}

	g := &RecentChooserWidget{native: c}

	return g
}

func (recv *RecentChooserWidget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *RecentChooserWidget) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *RecentChooserWidget) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *RecentChooserWidget) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RecentChooserWidget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *RecentChooserWidget) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to RecentChooserWidget.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentChooserWidget.
func CastToRecentChooserWidget(object *gobject.Object) *RecentChooserWidget {
	return RecentChooserWidgetNewFromC(object.ToC())
}

// RecentFilter is a wrapper around the C record GtkRecentFilter.
type RecentFilter struct {
	native *C.GtkRecentFilter
}

func RecentFilterNewFromC(u unsafe.Pointer) *RecentFilter {
	c := (*C.GtkRecentFilter)(u)
	if c == nil {
		return nil
	}

	g := &RecentFilter{native: c}

	return g
}

func (recv *RecentFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *RecentFilter) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *RecentFilter) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to RecentFilter.
// Exercise care, as this is a potentially dangerous function if the Object is not a RecentFilter.
func CastToRecentFilter(object *gobject.Object) *RecentFilter {
	return RecentFilterNewFromC(object.ToC())
}

// RendererCellAccessible is a wrapper around the C record GtkRendererCellAccessible.
type RendererCellAccessible struct {
	native *C.GtkRendererCellAccessible
	// parent : record
	// priv : record
}

func RendererCellAccessibleNewFromC(u unsafe.Pointer) *RendererCellAccessible {
	c := (*C.GtkRendererCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &RendererCellAccessible{native: c}

	return g
}

func (recv *RendererCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// CellAccessible upcasts to *CellAccessible
func (recv *RendererCellAccessible) CellAccessible() *CellAccessible {
	return CellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *RendererCellAccessible) Accessible() *Accessible {
	return recv.CellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *RendererCellAccessible) Object() *atk.Object {
	return recv.CellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to RendererCellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a RendererCellAccessible.
func CastToRendererCellAccessible(object *gobject.Object) *RendererCellAccessible {
	return RendererCellAccessibleNewFromC(object.ToC())
}

// RendererCellAccessibleNew is a wrapper around the C function gtk_renderer_cell_accessible_new.
func RendererCellAccessibleNew(renderer *CellRenderer) *RendererCellAccessible {
	c_renderer := (*C.GtkCellRenderer)(renderer.ToC())

	retC := C.gtk_renderer_cell_accessible_new(c_renderer)
	retGo := RendererCellAccessibleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Revealer is a wrapper around the C record GtkRevealer.
type Revealer struct {
	native *C.GtkRevealer
	// parent_instance : record
}

func RevealerNewFromC(u unsafe.Pointer) *Revealer {
	c := (*C.GtkRevealer)(u)
	if c == nil {
		return nil
	}

	g := &Revealer{native: c}

	return g
}

func (recv *Revealer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Revealer) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Revealer) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Revealer) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Revealer) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Revealer) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Revealer.
// Exercise care, as this is a potentially dangerous function if the Object is not a Revealer.
func CastToRevealer(object *gobject.Object) *Revealer {
	return RevealerNewFromC(object.ToC())
}

// Scale is a wrapper around the C record GtkScale.
type Scale struct {
	native *C.GtkScale
	// range : record
	// Private : priv
}

func ScaleNewFromC(u unsafe.Pointer) *Scale {
	c := (*C.GtkScale)(u)
	if c == nil {
		return nil
	}

	g := &Scale{native: c}

	return g
}

func (recv *Scale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Range upcasts to *Range
func (recv *Scale) Range() *Range {
	return RangeNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Scale) Widget() *Widget {
	return recv.Range().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Scale) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Range().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Scale) Object() *gobject.Object {
	return recv.Range().Object()
}

// CastToWidget down casts any arbitary Object to Scale.
// Exercise care, as this is a potentially dangerous function if the Object is not a Scale.
func CastToScale(object *gobject.Object) *Scale {
	return ScaleNewFromC(object.ToC())
}

var signalFormatValueId int
var signalFormatValueMap = make(map[int]ScaleSignalFormatValueCallback)
var signalFormatValueLock sync.Mutex

// ScaleSignalFormatValueCallback is a callback function for a 'format-value' signal emitted from a Scale.
type ScaleSignalFormatValueCallback func(value float64) string

func Scale_formatValueHandler() {}

// GetDigits is a wrapper around the C function gtk_scale_get_digits.
func (recv *Scale) GetDigits() int32 {
	retC := C.gtk_scale_get_digits((*C.GtkScale)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetDrawValue is a wrapper around the C function gtk_scale_get_draw_value.
func (recv *Scale) GetDrawValue() bool {
	retC := C.gtk_scale_get_draw_value((*C.GtkScale)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetValuePos is a wrapper around the C function gtk_scale_get_value_pos.
func (recv *Scale) GetValuePos() PositionType {
	retC := C.gtk_scale_get_value_pos((*C.GtkScale)(recv.native))
	retGo := (PositionType)(retC)

	return retGo
}

// SetDigits is a wrapper around the C function gtk_scale_set_digits.
func (recv *Scale) SetDigits(digits int32) {
	c_digits := (C.gint)(digits)

	C.gtk_scale_set_digits((*C.GtkScale)(recv.native), c_digits)

	return
}

// SetDrawValue is a wrapper around the C function gtk_scale_set_draw_value.
func (recv *Scale) SetDrawValue(drawValue bool) {
	c_draw_value :=
		boolToGboolean(drawValue)

	C.gtk_scale_set_draw_value((*C.GtkScale)(recv.native), c_draw_value)

	return
}

// SetValuePos is a wrapper around the C function gtk_scale_set_value_pos.
func (recv *Scale) SetValuePos(pos PositionType) {
	c_pos := (C.GtkPositionType)(pos)

	C.gtk_scale_set_value_pos((*C.GtkScale)(recv.native), c_pos)

	return
}

// ScaleAccessible is a wrapper around the C record GtkScaleAccessible.
type ScaleAccessible struct {
	native *C.GtkScaleAccessible
	// parent : record
	// priv : record
}

func ScaleAccessibleNewFromC(u unsafe.Pointer) *ScaleAccessible {
	c := (*C.GtkScaleAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScaleAccessible{native: c}

	return g
}

func (recv *ScaleAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RangeAccessible upcasts to *RangeAccessible
func (recv *ScaleAccessible) RangeAccessible() *RangeAccessible {
	return RangeAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ScaleAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.RangeAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ScaleAccessible) Accessible() *Accessible {
	return recv.RangeAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ScaleAccessible) Object() *atk.Object {
	return recv.RangeAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ScaleAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ScaleAccessible.
func CastToScaleAccessible(object *gobject.Object) *ScaleAccessible {
	return ScaleAccessibleNewFromC(object.ToC())
}

// ScaleButton is a wrapper around the C record GtkScaleButton.
type ScaleButton struct {
	native *C.GtkScaleButton
	// parent : record
	// Private : priv
}

func ScaleButtonNewFromC(u unsafe.Pointer) *ScaleButton {
	c := (*C.GtkScaleButton)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButton{native: c}

	return g
}

func (recv *ScaleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *ScaleButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ScaleButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *ScaleButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *ScaleButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ScaleButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ScaleButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to ScaleButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ScaleButton.
func CastToScaleButton(object *gobject.Object) *ScaleButton {
	return ScaleButtonNewFromC(object.ToC())
}

// Unsupported : gtk_scale_button_new : unsupported parameter size : no type generator for gint, GtkIconSize

// ScaleButtonAccessible is a wrapper around the C record GtkScaleButtonAccessible.
type ScaleButtonAccessible struct {
	native *C.GtkScaleButtonAccessible
	// parent : record
	// priv : record
}

func ScaleButtonAccessibleNewFromC(u unsafe.Pointer) *ScaleButtonAccessible {
	c := (*C.GtkScaleButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScaleButtonAccessible{native: c}

	return g
}

func (recv *ScaleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *ScaleButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return ButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ScaleButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ScaleButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ScaleButtonAccessible) Accessible() *Accessible {
	return recv.ButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ScaleButtonAccessible) Object() *atk.Object {
	return recv.ButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ScaleButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ScaleButtonAccessible.
func CastToScaleButtonAccessible(object *gobject.Object) *ScaleButtonAccessible {
	return ScaleButtonAccessibleNewFromC(object.ToC())
}

// Scrollbar is a wrapper around the C record GtkScrollbar.
type Scrollbar struct {
	native *C.GtkScrollbar
	// range : record
}

func ScrollbarNewFromC(u unsafe.Pointer) *Scrollbar {
	c := (*C.GtkScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &Scrollbar{native: c}

	return g
}

func (recv *Scrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Range upcasts to *Range
func (recv *Scrollbar) Range() *Range {
	return RangeNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Scrollbar) Widget() *Widget {
	return recv.Range().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Scrollbar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Range().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Scrollbar) Object() *gobject.Object {
	return recv.Range().Object()
}

// CastToWidget down casts any arbitary Object to Scrollbar.
// Exercise care, as this is a potentially dangerous function if the Object is not a Scrollbar.
func CastToScrollbar(object *gobject.Object) *Scrollbar {
	return ScrollbarNewFromC(object.ToC())
}

// ScrolledWindow is a wrapper around the C record GtkScrolledWindow.
type ScrolledWindow struct {
	native *C.GtkScrolledWindow
	// container : record
	// priv : record
}

func ScrolledWindowNewFromC(u unsafe.Pointer) *ScrolledWindow {
	c := (*C.GtkScrolledWindow)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindow{native: c}

	return g
}

func (recv *ScrolledWindow) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *ScrolledWindow) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ScrolledWindow) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *ScrolledWindow) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ScrolledWindow) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ScrolledWindow) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to ScrolledWindow.
// Exercise care, as this is a potentially dangerous function if the Object is not a ScrolledWindow.
func CastToScrolledWindow(object *gobject.Object) *ScrolledWindow {
	return ScrolledWindowNewFromC(object.ToC())
}

var signalMoveFocusOutId int
var signalMoveFocusOutMap = make(map[int]ScrolledWindowSignalMoveFocusOutCallback)
var signalMoveFocusOutLock sync.Mutex

// ScrolledWindowSignalMoveFocusOutCallback is a callback function for a 'move-focus-out' signal emitted from a ScrolledWindow.
type ScrolledWindowSignalMoveFocusOutCallback func(directionType DirectionType)

func ScrolledWindow_moveFocusOutHandler() {}

var signalScrollChildId int
var signalScrollChildMap = make(map[int]ScrolledWindowSignalScrollChildCallback)
var signalScrollChildLock sync.Mutex

// ScrolledWindowSignalScrollChildCallback is a callback function for a 'scroll-child' signal emitted from a ScrolledWindow.
type ScrolledWindowSignalScrollChildCallback func(scroll ScrollType, horizontal bool) bool

func ScrolledWindow_scrollChildHandler() {}

// ScrolledWindowNew is a wrapper around the C function gtk_scrolled_window_new.
func ScrolledWindowNew(hadjustment *Adjustment, vadjustment *Adjustment) *ScrolledWindow {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_scrolled_window_new(c_hadjustment, c_vadjustment)
	retGo := ScrolledWindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWithViewport is a wrapper around the C function gtk_scrolled_window_add_with_viewport.
func (recv *ScrolledWindow) AddWithViewport(child *Widget) {
	c_child := (*C.GtkWidget)(child.ToC())

	C.gtk_scrolled_window_add_with_viewport((*C.GtkScrolledWindow)(recv.native), c_child)

	return
}

// GetHadjustment is a wrapper around the C function gtk_scrolled_window_get_hadjustment.
func (recv *ScrolledWindow) GetHadjustment() *Adjustment {
	retC := C.gtk_scrolled_window_get_hadjustment((*C.GtkScrolledWindow)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPlacement is a wrapper around the C function gtk_scrolled_window_get_placement.
func (recv *ScrolledWindow) GetPlacement() CornerType {
	retC := C.gtk_scrolled_window_get_placement((*C.GtkScrolledWindow)(recv.native))
	retGo := (CornerType)(retC)

	return retGo
}

// Unsupported : gtk_scrolled_window_get_policy : unsupported parameter hscrollbar_policy : GtkPolicyType* with indirection level of 1

// GetShadowType is a wrapper around the C function gtk_scrolled_window_get_shadow_type.
func (recv *ScrolledWindow) GetShadowType() ShadowType {
	retC := C.gtk_scrolled_window_get_shadow_type((*C.GtkScrolledWindow)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_scrolled_window_get_vadjustment.
func (recv *ScrolledWindow) GetVadjustment() *Adjustment {
	retC := C.gtk_scrolled_window_get_vadjustment((*C.GtkScrolledWindow)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetHadjustment is a wrapper around the C function gtk_scrolled_window_set_hadjustment.
func (recv *ScrolledWindow) SetHadjustment(hadjustment *Adjustment) {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	C.gtk_scrolled_window_set_hadjustment((*C.GtkScrolledWindow)(recv.native), c_hadjustment)

	return
}

// SetPlacement is a wrapper around the C function gtk_scrolled_window_set_placement.
func (recv *ScrolledWindow) SetPlacement(windowPlacement CornerType) {
	c_window_placement := (C.GtkCornerType)(windowPlacement)

	C.gtk_scrolled_window_set_placement((*C.GtkScrolledWindow)(recv.native), c_window_placement)

	return
}

// SetPolicy is a wrapper around the C function gtk_scrolled_window_set_policy.
func (recv *ScrolledWindow) SetPolicy(hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	c_hscrollbar_policy := (C.GtkPolicyType)(hscrollbarPolicy)

	c_vscrollbar_policy := (C.GtkPolicyType)(vscrollbarPolicy)

	C.gtk_scrolled_window_set_policy((*C.GtkScrolledWindow)(recv.native), c_hscrollbar_policy, c_vscrollbar_policy)

	return
}

// SetShadowType is a wrapper around the C function gtk_scrolled_window_set_shadow_type.
func (recv *ScrolledWindow) SetShadowType(type_ ShadowType) {
	c_type := (C.GtkShadowType)(type_)

	C.gtk_scrolled_window_set_shadow_type((*C.GtkScrolledWindow)(recv.native), c_type)

	return
}

// SetVadjustment is a wrapper around the C function gtk_scrolled_window_set_vadjustment.
func (recv *ScrolledWindow) SetVadjustment(vadjustment *Adjustment) {
	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	C.gtk_scrolled_window_set_vadjustment((*C.GtkScrolledWindow)(recv.native), c_vadjustment)

	return
}

// ScrolledWindowAccessible is a wrapper around the C record GtkScrolledWindowAccessible.
type ScrolledWindowAccessible struct {
	native *C.GtkScrolledWindowAccessible
	// parent : record
	// priv : record
}

func ScrolledWindowAccessibleNewFromC(u unsafe.Pointer) *ScrolledWindowAccessible {
	c := (*C.GtkScrolledWindowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ScrolledWindowAccessible{native: c}

	return g
}

func (recv *ScrolledWindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ScrolledWindowAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ScrolledWindowAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ScrolledWindowAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ScrolledWindowAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ScrolledWindowAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ScrolledWindowAccessible.
func CastToScrolledWindowAccessible(object *gobject.Object) *ScrolledWindowAccessible {
	return ScrolledWindowAccessibleNewFromC(object.ToC())
}

// SearchBar is a wrapper around the C record GtkSearchBar.
type SearchBar struct {
	native *C.GtkSearchBar
	// Private : parent
}

func SearchBarNewFromC(u unsafe.Pointer) *SearchBar {
	c := (*C.GtkSearchBar)(u)
	if c == nil {
		return nil
	}

	g := &SearchBar{native: c}

	return g
}

func (recv *SearchBar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *SearchBar) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *SearchBar) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *SearchBar) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *SearchBar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *SearchBar) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to SearchBar.
// Exercise care, as this is a potentially dangerous function if the Object is not a SearchBar.
func CastToSearchBar(object *gobject.Object) *SearchBar {
	return SearchBarNewFromC(object.ToC())
}

// SearchEntry is a wrapper around the C record GtkSearchEntry.
type SearchEntry struct {
	native *C.GtkSearchEntry
	// parent : record
}

func SearchEntryNewFromC(u unsafe.Pointer) *SearchEntry {
	c := (*C.GtkSearchEntry)(u)
	if c == nil {
		return nil
	}

	g := &SearchEntry{native: c}

	return g
}

func (recv *SearchEntry) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Entry upcasts to *Entry
func (recv *SearchEntry) Entry() *Entry {
	return EntryNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *SearchEntry) Widget() *Widget {
	return recv.Entry().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *SearchEntry) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Entry().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *SearchEntry) Object() *gobject.Object {
	return recv.Entry().Object()
}

// CastToWidget down casts any arbitary Object to SearchEntry.
// Exercise care, as this is a potentially dangerous function if the Object is not a SearchEntry.
func CastToSearchEntry(object *gobject.Object) *SearchEntry {
	return SearchEntryNewFromC(object.ToC())
}

// Separator is a wrapper around the C record GtkSeparator.
type Separator struct {
	native *C.GtkSeparator
	// widget : record
	// priv : record
}

func SeparatorNewFromC(u unsafe.Pointer) *Separator {
	c := (*C.GtkSeparator)(u)
	if c == nil {
		return nil
	}

	g := &Separator{native: c}

	return g
}

func (recv *Separator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Separator) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Separator) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Separator) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Separator.
// Exercise care, as this is a potentially dangerous function if the Object is not a Separator.
func CastToSeparator(object *gobject.Object) *Separator {
	return SeparatorNewFromC(object.ToC())
}

// SeparatorMenuItem is a wrapper around the C record GtkSeparatorMenuItem.
type SeparatorMenuItem struct {
	native *C.GtkSeparatorMenuItem
	// menu_item : record
}

func SeparatorMenuItemNewFromC(u unsafe.Pointer) *SeparatorMenuItem {
	c := (*C.GtkSeparatorMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorMenuItem{native: c}

	return g
}

func (recv *SeparatorMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem upcasts to *MenuItem
func (recv *SeparatorMenuItem) MenuItem() *MenuItem {
	return MenuItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *SeparatorMenuItem) Bin() *Bin {
	return recv.MenuItem().Bin()
}

// Container upcasts to *Container
func (recv *SeparatorMenuItem) Container() *Container {
	return recv.MenuItem().Container()
}

// Widget upcasts to *Widget
func (recv *SeparatorMenuItem) Widget() *Widget {
	return recv.MenuItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *SeparatorMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *SeparatorMenuItem) Object() *gobject.Object {
	return recv.MenuItem().Object()
}

// CastToWidget down casts any arbitary Object to SeparatorMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a SeparatorMenuItem.
func CastToSeparatorMenuItem(object *gobject.Object) *SeparatorMenuItem {
	return SeparatorMenuItemNewFromC(object.ToC())
}

// SeparatorMenuItemNew is a wrapper around the C function gtk_separator_menu_item_new.
func SeparatorMenuItemNew() *SeparatorMenuItem {
	retC := C.gtk_separator_menu_item_new()
	retGo := SeparatorMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SeparatorToolItem is a wrapper around the C record GtkSeparatorToolItem.
type SeparatorToolItem struct {
	native *C.GtkSeparatorToolItem
	// parent : record
	// Private : priv
}

func SeparatorToolItemNewFromC(u unsafe.Pointer) *SeparatorToolItem {
	c := (*C.GtkSeparatorToolItem)(u)
	if c == nil {
		return nil
	}

	g := &SeparatorToolItem{native: c}

	return g
}

func (recv *SeparatorToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItem upcasts to *ToolItem
func (recv *SeparatorToolItem) ToolItem() *ToolItem {
	return ToolItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *SeparatorToolItem) Bin() *Bin {
	return recv.ToolItem().Bin()
}

// Container upcasts to *Container
func (recv *SeparatorToolItem) Container() *Container {
	return recv.ToolItem().Container()
}

// Widget upcasts to *Widget
func (recv *SeparatorToolItem) Widget() *Widget {
	return recv.ToolItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *SeparatorToolItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToolItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *SeparatorToolItem) Object() *gobject.Object {
	return recv.ToolItem().Object()
}

// CastToWidget down casts any arbitary Object to SeparatorToolItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a SeparatorToolItem.
func CastToSeparatorToolItem(object *gobject.Object) *SeparatorToolItem {
	return SeparatorToolItemNewFromC(object.ToC())
}

// Settings is a wrapper around the C record GtkSettings.
type Settings struct {
	native *C.GtkSettings
	// parent_instance : record
	// Private : priv
}

func SettingsNewFromC(u unsafe.Pointer) *Settings {
	c := (*C.GtkSettings)(u)
	if c == nil {
		return nil
	}

	g := &Settings{native: c}

	return g
}

func (recv *Settings) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Settings) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Settings.
// Exercise care, as this is a potentially dangerous function if the Object is not a Settings.
func CastToSettings(object *gobject.Object) *Settings {
	return SettingsNewFromC(object.ToC())
}

// SetDoubleProperty is a wrapper around the C function gtk_settings_set_double_property.
func (recv *Settings) SetDoubleProperty(name string, vDouble float64, origin string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_v_double := (C.gdouble)(vDouble)

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_double_property((*C.GtkSettings)(recv.native), c_name, c_v_double, c_origin)

	return
}

// SetLongProperty is a wrapper around the C function gtk_settings_set_long_property.
func (recv *Settings) SetLongProperty(name string, vLong int64, origin string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_v_long := (C.glong)(vLong)

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_long_property((*C.GtkSettings)(recv.native), c_name, c_v_long, c_origin)

	return
}

// SetPropertyValue is a wrapper around the C function gtk_settings_set_property_value.
func (recv *Settings) SetPropertyValue(name string, svalue *SettingsValue) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_svalue := (*C.GtkSettingsValue)(svalue.ToC())

	C.gtk_settings_set_property_value((*C.GtkSettings)(recv.native), c_name, c_svalue)

	return
}

// SetStringProperty is a wrapper around the C function gtk_settings_set_string_property.
func (recv *Settings) SetStringProperty(name string, vString string, origin string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_v_string := C.CString(vString)
	defer C.free(unsafe.Pointer(c_v_string))

	c_origin := C.CString(origin)
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_string_property((*C.GtkSettings)(recv.native), c_name, c_v_string, c_origin)

	return
}

var signalChangeCurrentPageId int
var signalChangeCurrentPageMap = make(map[int]ShortcutsSectionSignalChangeCurrentPageCallback)
var signalChangeCurrentPageLock sync.Mutex

// ShortcutsSectionSignalChangeCurrentPageCallback is a callback function for a 'change-current-page' signal emitted from a ShortcutsSection.
type ShortcutsSectionSignalChangeCurrentPageCallback func(object int32) bool

func ShortcutsSection_changeCurrentPageHandler() {}

var signalCloseId int
var signalCloseMap = make(map[int]ShortcutsWindowSignalCloseCallback)
var signalCloseLock sync.Mutex

// ShortcutsWindowSignalCloseCallback is a callback function for a 'close' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalCloseCallback func()

func ShortcutsWindow_closeHandler() {}

var signalSearchId int
var signalSearchMap = make(map[int]ShortcutsWindowSignalSearchCallback)
var signalSearchLock sync.Mutex

// ShortcutsWindowSignalSearchCallback is a callback function for a 'search' signal emitted from a ShortcutsWindow.
type ShortcutsWindowSignalSearchCallback func()

func ShortcutsWindow_searchHandler() {}

// SizeGroup is a wrapper around the C record GtkSizeGroup.
type SizeGroup struct {
	native *C.GtkSizeGroup
	// parent_instance : record
	// Private : priv
}

func SizeGroupNewFromC(u unsafe.Pointer) *SizeGroup {
	c := (*C.GtkSizeGroup)(u)
	if c == nil {
		return nil
	}

	g := &SizeGroup{native: c}

	return g
}

func (recv *SizeGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *SizeGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to SizeGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a SizeGroup.
func CastToSizeGroup(object *gobject.Object) *SizeGroup {
	return SizeGroupNewFromC(object.ToC())
}

// SizeGroupNew is a wrapper around the C function gtk_size_group_new.
func SizeGroupNew(mode SizeGroupMode) *SizeGroup {
	c_mode := (C.GtkSizeGroupMode)(mode)

	retC := C.gtk_size_group_new(c_mode)
	retGo := SizeGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWidget is a wrapper around the C function gtk_size_group_add_widget.
func (recv *SizeGroup) AddWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_size_group_add_widget((*C.GtkSizeGroup)(recv.native), c_widget)

	return
}

// GetMode is a wrapper around the C function gtk_size_group_get_mode.
func (recv *SizeGroup) GetMode() SizeGroupMode {
	retC := C.gtk_size_group_get_mode((*C.GtkSizeGroup)(recv.native))
	retGo := (SizeGroupMode)(retC)

	return retGo
}

// RemoveWidget is a wrapper around the C function gtk_size_group_remove_widget.
func (recv *SizeGroup) RemoveWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_size_group_remove_widget((*C.GtkSizeGroup)(recv.native), c_widget)

	return
}

// SetMode is a wrapper around the C function gtk_size_group_set_mode.
func (recv *SizeGroup) SetMode(mode SizeGroupMode) {
	c_mode := (C.GtkSizeGroupMode)(mode)

	C.gtk_size_group_set_mode((*C.GtkSizeGroup)(recv.native), c_mode)

	return
}

// Blacklisted : GtkSocket

// SpinButton is a wrapper around the C record GtkSpinButton.
type SpinButton struct {
	native *C.GtkSpinButton
	// entry : record
	// Private : priv
}

func SpinButtonNewFromC(u unsafe.Pointer) *SpinButton {
	c := (*C.GtkSpinButton)(u)
	if c == nil {
		return nil
	}

	g := &SpinButton{native: c}

	return g
}

func (recv *SpinButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Entry upcasts to *Entry
func (recv *SpinButton) Entry() *Entry {
	return EntryNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *SpinButton) Widget() *Widget {
	return recv.Entry().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *SpinButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Entry().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *SpinButton) Object() *gobject.Object {
	return recv.Entry().Object()
}

// CastToWidget down casts any arbitary Object to SpinButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a SpinButton.
func CastToSpinButton(object *gobject.Object) *SpinButton {
	return SpinButtonNewFromC(object.ToC())
}

var signalChangeValueId int
var signalChangeValueMap = make(map[int]SpinButtonSignalChangeValueCallback)
var signalChangeValueLock sync.Mutex

// SpinButtonSignalChangeValueCallback is a callback function for a 'change-value' signal emitted from a SpinButton.
type SpinButtonSignalChangeValueCallback func(scroll ScrollType)

func SpinButton_changeValueHandler() {}

var signalInputId int
var signalInputMap = make(map[int]SpinButtonSignalInputCallback)
var signalInputLock sync.Mutex

// SpinButtonSignalInputCallback is a callback function for a 'input' signal emitted from a SpinButton.
type SpinButtonSignalInputCallback func() int32

func SpinButton_inputHandler() {}

var signalOutputId int
var signalOutputMap = make(map[int]SpinButtonSignalOutputCallback)
var signalOutputLock sync.Mutex

// SpinButtonSignalOutputCallback is a callback function for a 'output' signal emitted from a SpinButton.
type SpinButtonSignalOutputCallback func() bool

func SpinButton_outputHandler() {}

var signalValueChangedId int
var signalValueChangedMap = make(map[int]SpinButtonSignalValueChangedCallback)
var signalValueChangedLock sync.Mutex

// SpinButtonSignalValueChangedCallback is a callback function for a 'value-changed' signal emitted from a SpinButton.
type SpinButtonSignalValueChangedCallback func()

func SpinButton_valueChangedHandler() {}

// SpinButtonNew is a wrapper around the C function gtk_spin_button_new.
func SpinButtonNew(adjustment *Adjustment, climbRate float64, digits uint32) *SpinButton {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	c_climb_rate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	retC := C.gtk_spin_button_new(c_adjustment, c_climb_rate, c_digits)
	retGo := SpinButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SpinButtonNewWithRange is a wrapper around the C function gtk_spin_button_new_with_range.
func SpinButtonNewWithRange(min float64, max float64, step float64) *SpinButton {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_spin_button_new_with_range(c_min, c_max, c_step)
	retGo := SpinButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Configure is a wrapper around the C function gtk_spin_button_configure.
func (recv *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint32) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	c_climb_rate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	C.gtk_spin_button_configure((*C.GtkSpinButton)(recv.native), c_adjustment, c_climb_rate, c_digits)

	return
}

// GetAdjustment is a wrapper around the C function gtk_spin_button_get_adjustment.
func (recv *SpinButton) GetAdjustment() *Adjustment {
	retC := C.gtk_spin_button_get_adjustment((*C.GtkSpinButton)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDigits is a wrapper around the C function gtk_spin_button_get_digits.
func (recv *SpinButton) GetDigits() uint32 {
	retC := C.gtk_spin_button_get_digits((*C.GtkSpinButton)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetIncrements is a wrapper around the C function gtk_spin_button_get_increments.
func (recv *SpinButton) GetIncrements() (float64, float64) {
	var c_step C.gdouble

	var c_page C.gdouble

	C.gtk_spin_button_get_increments((*C.GtkSpinButton)(recv.native), &c_step, &c_page)

	step := (float64)(c_step)

	page := (float64)(c_page)

	return step, page
}

// GetNumeric is a wrapper around the C function gtk_spin_button_get_numeric.
func (recv *SpinButton) GetNumeric() bool {
	retC := C.gtk_spin_button_get_numeric((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRange is a wrapper around the C function gtk_spin_button_get_range.
func (recv *SpinButton) GetRange() (float64, float64) {
	var c_min C.gdouble

	var c_max C.gdouble

	C.gtk_spin_button_get_range((*C.GtkSpinButton)(recv.native), &c_min, &c_max)

	min := (float64)(c_min)

	max := (float64)(c_max)

	return min, max
}

// GetSnapToTicks is a wrapper around the C function gtk_spin_button_get_snap_to_ticks.
func (recv *SpinButton) GetSnapToTicks() bool {
	retC := C.gtk_spin_button_get_snap_to_ticks((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUpdatePolicy is a wrapper around the C function gtk_spin_button_get_update_policy.
func (recv *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	retC := C.gtk_spin_button_get_update_policy((*C.GtkSpinButton)(recv.native))
	retGo := (SpinButtonUpdatePolicy)(retC)

	return retGo
}

// GetValue is a wrapper around the C function gtk_spin_button_get_value.
func (recv *SpinButton) GetValue() float64 {
	retC := C.gtk_spin_button_get_value((*C.GtkSpinButton)(recv.native))
	retGo := (float64)(retC)

	return retGo
}

// GetValueAsInt is a wrapper around the C function gtk_spin_button_get_value_as_int.
func (recv *SpinButton) GetValueAsInt() int32 {
	retC := C.gtk_spin_button_get_value_as_int((*C.GtkSpinButton)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetWrap is a wrapper around the C function gtk_spin_button_get_wrap.
func (recv *SpinButton) GetWrap() bool {
	retC := C.gtk_spin_button_get_wrap((*C.GtkSpinButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAdjustment is a wrapper around the C function gtk_spin_button_set_adjustment.
func (recv *SpinButton) SetAdjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_spin_button_set_adjustment((*C.GtkSpinButton)(recv.native), c_adjustment)

	return
}

// SetDigits is a wrapper around the C function gtk_spin_button_set_digits.
func (recv *SpinButton) SetDigits(digits uint32) {
	c_digits := (C.guint)(digits)

	C.gtk_spin_button_set_digits((*C.GtkSpinButton)(recv.native), c_digits)

	return
}

// SetIncrements is a wrapper around the C function gtk_spin_button_set_increments.
func (recv *SpinButton) SetIncrements(step float64, page float64) {
	c_step := (C.gdouble)(step)

	c_page := (C.gdouble)(page)

	C.gtk_spin_button_set_increments((*C.GtkSpinButton)(recv.native), c_step, c_page)

	return
}

// SetNumeric is a wrapper around the C function gtk_spin_button_set_numeric.
func (recv *SpinButton) SetNumeric(numeric bool) {
	c_numeric :=
		boolToGboolean(numeric)

	C.gtk_spin_button_set_numeric((*C.GtkSpinButton)(recv.native), c_numeric)

	return
}

// SetRange is a wrapper around the C function gtk_spin_button_set_range.
func (recv *SpinButton) SetRange(min float64, max float64) {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	C.gtk_spin_button_set_range((*C.GtkSpinButton)(recv.native), c_min, c_max)

	return
}

// SetSnapToTicks is a wrapper around the C function gtk_spin_button_set_snap_to_ticks.
func (recv *SpinButton) SetSnapToTicks(snapToTicks bool) {
	c_snap_to_ticks :=
		boolToGboolean(snapToTicks)

	C.gtk_spin_button_set_snap_to_ticks((*C.GtkSpinButton)(recv.native), c_snap_to_ticks)

	return
}

// SetUpdatePolicy is a wrapper around the C function gtk_spin_button_set_update_policy.
func (recv *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	c_policy := (C.GtkSpinButtonUpdatePolicy)(policy)

	C.gtk_spin_button_set_update_policy((*C.GtkSpinButton)(recv.native), c_policy)

	return
}

// SetValue is a wrapper around the C function gtk_spin_button_set_value.
func (recv *SpinButton) SetValue(value float64) {
	c_value := (C.gdouble)(value)

	C.gtk_spin_button_set_value((*C.GtkSpinButton)(recv.native), c_value)

	return
}

// SetWrap is a wrapper around the C function gtk_spin_button_set_wrap.
func (recv *SpinButton) SetWrap(wrap bool) {
	c_wrap :=
		boolToGboolean(wrap)

	C.gtk_spin_button_set_wrap((*C.GtkSpinButton)(recv.native), c_wrap)

	return
}

// Spin is a wrapper around the C function gtk_spin_button_spin.
func (recv *SpinButton) Spin(direction SpinType, increment float64) {
	c_direction := (C.GtkSpinType)(direction)

	c_increment := (C.gdouble)(increment)

	C.gtk_spin_button_spin((*C.GtkSpinButton)(recv.native), c_direction, c_increment)

	return
}

// Update is a wrapper around the C function gtk_spin_button_update.
func (recv *SpinButton) Update() {
	C.gtk_spin_button_update((*C.GtkSpinButton)(recv.native))

	return
}

// SpinButtonAccessible is a wrapper around the C record GtkSpinButtonAccessible.
type SpinButtonAccessible struct {
	native *C.GtkSpinButtonAccessible
	// parent : record
	// priv : record
}

func SpinButtonAccessibleNewFromC(u unsafe.Pointer) *SpinButtonAccessible {
	c := (*C.GtkSpinButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SpinButtonAccessible{native: c}

	return g
}

func (recv *SpinButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EntryAccessible upcasts to *EntryAccessible
func (recv *SpinButtonAccessible) EntryAccessible() *EntryAccessible {
	return EntryAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *SpinButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.EntryAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *SpinButtonAccessible) Accessible() *Accessible {
	return recv.EntryAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *SpinButtonAccessible) Object() *atk.Object {
	return recv.EntryAccessible().Object()
}

// CastToWidget down casts any arbitary Object to SpinButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a SpinButtonAccessible.
func CastToSpinButtonAccessible(object *gobject.Object) *SpinButtonAccessible {
	return SpinButtonAccessibleNewFromC(object.ToC())
}

// Spinner is a wrapper around the C record GtkSpinner.
type Spinner struct {
	native *C.GtkSpinner
	// parent : record
	// Private : priv
}

func SpinnerNewFromC(u unsafe.Pointer) *Spinner {
	c := (*C.GtkSpinner)(u)
	if c == nil {
		return nil
	}

	g := &Spinner{native: c}

	return g
}

func (recv *Spinner) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Spinner) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Spinner) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Spinner) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Spinner.
// Exercise care, as this is a potentially dangerous function if the Object is not a Spinner.
func CastToSpinner(object *gobject.Object) *Spinner {
	return SpinnerNewFromC(object.ToC())
}

// SpinnerAccessible is a wrapper around the C record GtkSpinnerAccessible.
type SpinnerAccessible struct {
	native *C.GtkSpinnerAccessible
	// parent : record
	// priv : record
}

func SpinnerAccessibleNewFromC(u unsafe.Pointer) *SpinnerAccessible {
	c := (*C.GtkSpinnerAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SpinnerAccessible{native: c}

	return g
}

func (recv *SpinnerAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *SpinnerAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *SpinnerAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *SpinnerAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to SpinnerAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a SpinnerAccessible.
func CastToSpinnerAccessible(object *gobject.Object) *SpinnerAccessible {
	return SpinnerAccessibleNewFromC(object.ToC())
}

// Stack is a wrapper around the C record GtkStack.
type Stack struct {
	native *C.GtkStack
	// parent_instance : record
}

func StackNewFromC(u unsafe.Pointer) *Stack {
	c := (*C.GtkStack)(u)
	if c == nil {
		return nil
	}

	g := &Stack{native: c}

	return g
}

func (recv *Stack) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Stack) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Stack) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Stack) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Stack) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Stack.
// Exercise care, as this is a potentially dangerous function if the Object is not a Stack.
func CastToStack(object *gobject.Object) *Stack {
	return StackNewFromC(object.ToC())
}

// Blacklisted : GtkStackAccessible

// StackSidebar is a wrapper around the C record GtkStackSidebar.
type StackSidebar struct {
	native *C.GtkStackSidebar
	// parent : record
}

func StackSidebarNewFromC(u unsafe.Pointer) *StackSidebar {
	c := (*C.GtkStackSidebar)(u)
	if c == nil {
		return nil
	}

	g := &StackSidebar{native: c}

	return g
}

func (recv *StackSidebar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *StackSidebar) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *StackSidebar) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *StackSidebar) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StackSidebar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *StackSidebar) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to StackSidebar.
// Exercise care, as this is a potentially dangerous function if the Object is not a StackSidebar.
func CastToStackSidebar(object *gobject.Object) *StackSidebar {
	return StackSidebarNewFromC(object.ToC())
}

// StackSwitcher is a wrapper around the C record GtkStackSwitcher.
type StackSwitcher struct {
	native *C.GtkStackSwitcher
	// widget : record
}

func StackSwitcherNewFromC(u unsafe.Pointer) *StackSwitcher {
	c := (*C.GtkStackSwitcher)(u)
	if c == nil {
		return nil
	}

	g := &StackSwitcher{native: c}

	return g
}

func (recv *StackSwitcher) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *StackSwitcher) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *StackSwitcher) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *StackSwitcher) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *StackSwitcher) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *StackSwitcher) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to StackSwitcher.
// Exercise care, as this is a potentially dangerous function if the Object is not a StackSwitcher.
func CastToStackSwitcher(object *gobject.Object) *StackSwitcher {
	return StackSwitcherNewFromC(object.ToC())
}

// StatusIcon is a wrapper around the C record GtkStatusIcon.
type StatusIcon struct {
	native *C.GtkStatusIcon
	// parent_instance : record
	// priv : record
}

func StatusIconNewFromC(u unsafe.Pointer) *StatusIcon {
	c := (*C.GtkStatusIcon)(u)
	if c == nil {
		return nil
	}

	g := &StatusIcon{native: c}

	return g
}

func (recv *StatusIcon) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *StatusIcon) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to StatusIcon.
// Exercise care, as this is a potentially dangerous function if the Object is not a StatusIcon.
func CastToStatusIcon(object *gobject.Object) *StatusIcon {
	return StatusIconNewFromC(object.ToC())
}

// Unsupported : gtk_status_icon_new_from_gicon : unsupported parameter icon : no type generator for Gio.Icon, GIcon*

// Statusbar is a wrapper around the C record GtkStatusbar.
type Statusbar struct {
	native *C.GtkStatusbar
	// parent_widget : record
	// Private : priv
}

func StatusbarNewFromC(u unsafe.Pointer) *Statusbar {
	c := (*C.GtkStatusbar)(u)
	if c == nil {
		return nil
	}

	g := &Statusbar{native: c}

	return g
}

func (recv *Statusbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *Statusbar) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Statusbar) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *Statusbar) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Statusbar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Statusbar) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to Statusbar.
// Exercise care, as this is a potentially dangerous function if the Object is not a Statusbar.
func CastToStatusbar(object *gobject.Object) *Statusbar {
	return StatusbarNewFromC(object.ToC())
}

var signalTextPoppedId int
var signalTextPoppedMap = make(map[int]StatusbarSignalTextPoppedCallback)
var signalTextPoppedLock sync.Mutex

// StatusbarSignalTextPoppedCallback is a callback function for a 'text-popped' signal emitted from a Statusbar.
type StatusbarSignalTextPoppedCallback func(contextId uint32, text string)

func Statusbar_textPoppedHandler() {}

var signalTextPushedId int
var signalTextPushedMap = make(map[int]StatusbarSignalTextPushedCallback)
var signalTextPushedLock sync.Mutex

// StatusbarSignalTextPushedCallback is a callback function for a 'text-pushed' signal emitted from a Statusbar.
type StatusbarSignalTextPushedCallback func(contextId uint32, text string)

func Statusbar_textPushedHandler() {}

// StatusbarNew is a wrapper around the C function gtk_statusbar_new.
func StatusbarNew() *Statusbar {
	retC := C.gtk_statusbar_new()
	retGo := StatusbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetContextId is a wrapper around the C function gtk_statusbar_get_context_id.
func (recv *Statusbar) GetContextId(contextDescription string) uint32 {
	c_context_description := C.CString(contextDescription)
	defer C.free(unsafe.Pointer(c_context_description))

	retC := C.gtk_statusbar_get_context_id((*C.GtkStatusbar)(recv.native), c_context_description)
	retGo := (uint32)(retC)

	return retGo
}

// Pop is a wrapper around the C function gtk_statusbar_pop.
func (recv *Statusbar) Pop(contextId uint32) {
	c_context_id := (C.guint)(contextId)

	C.gtk_statusbar_pop((*C.GtkStatusbar)(recv.native), c_context_id)

	return
}

// Push is a wrapper around the C function gtk_statusbar_push.
func (recv *Statusbar) Push(contextId uint32, text string) uint32 {
	c_context_id := (C.guint)(contextId)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_statusbar_push((*C.GtkStatusbar)(recv.native), c_context_id, c_text)
	retGo := (uint32)(retC)

	return retGo
}

// Remove is a wrapper around the C function gtk_statusbar_remove.
func (recv *Statusbar) Remove(contextId uint32, messageId uint32) {
	c_context_id := (C.guint)(contextId)

	c_message_id := (C.guint)(messageId)

	C.gtk_statusbar_remove((*C.GtkStatusbar)(recv.native), c_context_id, c_message_id)

	return
}

// StatusbarAccessible is a wrapper around the C record GtkStatusbarAccessible.
type StatusbarAccessible struct {
	native *C.GtkStatusbarAccessible
	// parent : record
	// priv : record
}

func StatusbarAccessibleNewFromC(u unsafe.Pointer) *StatusbarAccessible {
	c := (*C.GtkStatusbarAccessible)(u)
	if c == nil {
		return nil
	}

	g := &StatusbarAccessible{native: c}

	return g
}

func (recv *StatusbarAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *StatusbarAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *StatusbarAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *StatusbarAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *StatusbarAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to StatusbarAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a StatusbarAccessible.
func CastToStatusbarAccessible(object *gobject.Object) *StatusbarAccessible {
	return StatusbarAccessibleNewFromC(object.ToC())
}

// Style is a wrapper around the C record GtkStyle.
type Style struct {
	native *C.GtkStyle
	// Private : parent_instance
	// no type for fg
	// no type for bg
	// no type for light
	// no type for dark
	// no type for mid
	// no type for text
	// no type for base
	// no type for text_aa
	// black : record
	// white : record
	// font_desc : record
	Xthickness int32
	Ythickness int32
	// no type for background
	// Private : attach_count
	// Private : visual
	// Private : private_font_desc
	// Private : rc_style
	// Private : styles
	// Private : property_cache
	// Private : icon_factories
}

func StyleNewFromC(u unsafe.Pointer) *Style {
	c := (*C.GtkStyle)(u)
	if c == nil {
		return nil
	}

	g := &Style{
		Xthickness: (int32)(c.xthickness),
		Ythickness: (int32)(c.ythickness),
		native:     c,
	}

	return g
}

func (recv *Style) ToC() unsafe.Pointer {
	recv.native.xthickness =
		(C.gint)(recv.Xthickness)
	recv.native.ythickness =
		(C.gint)(recv.Ythickness)

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Style) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Style.
// Exercise care, as this is a potentially dangerous function if the Object is not a Style.
func CastToStyle(object *gobject.Object) *Style {
	return StyleNewFromC(object.ToC())
}

// StyleNew is a wrapper around the C function gtk_style_new.
func StyleNew() *Style {
	retC := C.gtk_style_new()
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ApplyDefaultBackground is a wrapper around the C function gtk_style_apply_default_background.
func (recv *Style) ApplyDefaultBackground(cr *cairo.Context, window *gdk.Window, stateType StateType, x int32, y int32, width int32, height int32) {
	c_cr := (*C.cairo_t)(cr.ToC())

	c_window := (*C.GdkWindow)(window.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_style_apply_default_background((*C.GtkStyle)(recv.native), c_cr, c_window, c_state_type, c_x, c_y, c_width, c_height)

	return
}

// Attach is a wrapper around the C function gtk_style_attach.
func (recv *Style) Attach(window *gdk.Window) *Style {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_style_attach((*C.GtkStyle)(recv.native), c_window)
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function gtk_style_copy.
func (recv *Style) Copy() *Style {
	retC := C.gtk_style_copy((*C.GtkStyle)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Detach is a wrapper around the C function gtk_style_detach.
func (recv *Style) Detach() {
	C.gtk_style_detach((*C.GtkStyle)(recv.native))

	return
}

// LookupIconSet is a wrapper around the C function gtk_style_lookup_icon_set.
func (recv *Style) LookupIconSet(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_style_lookup_icon_set((*C.GtkStyle)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_style_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// SetBackground is a wrapper around the C function gtk_style_set_background.
func (recv *Style) SetBackground(window *gdk.Window, stateType StateType) {
	c_window := (*C.GdkWindow)(window.ToC())

	c_state_type := (C.GtkStateType)(stateType)

	C.gtk_style_set_background((*C.GtkStyle)(recv.native), c_window, c_state_type)

	return
}

// StyleContext is a wrapper around the C record GtkStyleContext.
type StyleContext struct {
	native *C.GtkStyleContext
	// parent_object : record
	// priv : record
}

func StyleContextNewFromC(u unsafe.Pointer) *StyleContext {
	c := (*C.GtkStyleContext)(u)
	if c == nil {
		return nil
	}

	g := &StyleContext{native: c}

	return g
}

func (recv *StyleContext) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *StyleContext) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to StyleContext.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleContext.
func CastToStyleContext(object *gobject.Object) *StyleContext {
	return StyleContextNewFromC(object.ToC())
}

// StyleContextNew is a wrapper around the C function gtk_style_context_new.
func StyleContextNew() *StyleContext {
	retC := C.gtk_style_context_new()
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetScreen is a wrapper around the C function gtk_style_context_get_screen.
func (recv *StyleContext) GetScreen() *gdk.Screen {
	retC := C.gtk_style_context_get_screen((*C.GtkStyleContext)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSection is a wrapper around the C function gtk_style_context_get_section.
func (recv *StyleContext) GetSection(property string) *CssSection {
	c_property := C.CString(property)
	defer C.free(unsafe.Pointer(c_property))

	retC := C.gtk_style_context_get_section((*C.GtkStyleContext)(recv.native), c_property)
	retGo := CssSectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStyleProperty is a wrapper around the C function gtk_style_context_get_style_property.
func (recv *StyleContext) GetStyleProperty(propertyName string, value *gobject.Value) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_style_context_get_style_property((*C.GtkStyleContext)(recv.native), c_property_name, c_value)

	return
}

// LookupColor is a wrapper around the C function gtk_style_context_lookup_color.
func (recv *StyleContext) LookupColor(colorName string) (bool, *gdk.RGBA) {
	c_color_name := C.CString(colorName)
	defer C.free(unsafe.Pointer(c_color_name))

	var c_color C.GdkRGBA

	retC := C.gtk_style_context_lookup_color((*C.GtkStyleContext)(recv.native), c_color_name, &c_color)
	retGo := retC == C.TRUE

	color := gdk.RGBANewFromC(unsafe.Pointer(&c_color))

	return retGo, color
}

// LookupIconSet is a wrapper around the C function gtk_style_context_lookup_icon_set.
func (recv *StyleContext) LookupIconSet(stockId string) *IconSet {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_style_context_lookup_icon_set((*C.GtkStyleContext)(recv.native), c_stock_id)
	retGo := IconSetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// StyleProperties is a wrapper around the C record GtkStyleProperties.
type StyleProperties struct {
	native *C.GtkStyleProperties
	// Private : parent_object
	// Private : priv
}

func StylePropertiesNewFromC(u unsafe.Pointer) *StyleProperties {
	c := (*C.GtkStyleProperties)(u)
	if c == nil {
		return nil
	}

	g := &StyleProperties{native: c}

	return g
}

func (recv *StyleProperties) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *StyleProperties) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to StyleProperties.
// Exercise care, as this is a potentially dangerous function if the Object is not a StyleProperties.
func CastToStyleProperties(object *gobject.Object) *StyleProperties {
	return StylePropertiesNewFromC(object.ToC())
}

// StylePropertiesNew is a wrapper around the C function gtk_style_properties_new.
func StylePropertiesNew() *StyleProperties {
	retC := C.gtk_style_properties_new()
	retGo := StylePropertiesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Clear is a wrapper around the C function gtk_style_properties_clear.
func (recv *StyleProperties) Clear() {
	C.gtk_style_properties_clear((*C.GtkStyleProperties)(recv.native))

	return
}

// Switch is a wrapper around the C record GtkSwitch.
type Switch struct {
	native *C.GtkSwitch
	// Private : parent_instance
	// Private : priv
}

func SwitchNewFromC(u unsafe.Pointer) *Switch {
	c := (*C.GtkSwitch)(u)
	if c == nil {
		return nil
	}

	g := &Switch{native: c}

	return g
}

func (recv *Switch) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Widget upcasts to *Widget
func (recv *Switch) Widget() *Widget {
	return WidgetNewFromC(unsafe.Pointer(recv.native))
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Switch) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Widget().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Switch) Object() *gobject.Object {
	return recv.Widget().Object()
}

// CastToWidget down casts any arbitary Object to Switch.
// Exercise care, as this is a potentially dangerous function if the Object is not a Switch.
func CastToSwitch(object *gobject.Object) *Switch {
	return SwitchNewFromC(object.ToC())
}

var signalActivateId int
var signalActivateMap = make(map[int]SwitchSignalActivateCallback)
var signalActivateLock sync.Mutex

// SwitchSignalActivateCallback is a callback function for a 'activate' signal emitted from a Switch.
type SwitchSignalActivateCallback func()

func Switch_activateHandler() {}

// SwitchAccessible is a wrapper around the C record GtkSwitchAccessible.
type SwitchAccessible struct {
	native *C.GtkSwitchAccessible
	// parent : record
	// priv : record
}

func SwitchAccessibleNewFromC(u unsafe.Pointer) *SwitchAccessible {
	c := (*C.GtkSwitchAccessible)(u)
	if c == nil {
		return nil
	}

	g := &SwitchAccessible{native: c}

	return g
}

func (recv *SwitchAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *SwitchAccessible) WidgetAccessible() *WidgetAccessible {
	return WidgetAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Accessible upcasts to *Accessible
func (recv *SwitchAccessible) Accessible() *Accessible {
	return recv.WidgetAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *SwitchAccessible) Object() *atk.Object {
	return recv.WidgetAccessible().Object()
}

// CastToWidget down casts any arbitary Object to SwitchAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a SwitchAccessible.
func CastToSwitchAccessible(object *gobject.Object) *SwitchAccessible {
	return SwitchAccessibleNewFromC(object.ToC())
}

// Table is a wrapper around the C record GtkTable.
type Table struct {
	native *C.GtkTable
	// container : record
	// Private : priv
}

func TableNewFromC(u unsafe.Pointer) *Table {
	c := (*C.GtkTable)(u)
	if c == nil {
		return nil
	}

	g := &Table{native: c}

	return g
}

func (recv *Table) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Table) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Table) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Table) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Table) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Table.
// Exercise care, as this is a potentially dangerous function if the Object is not a Table.
func CastToTable(object *gobject.Object) *Table {
	return TableNewFromC(object.ToC())
}

// TableNew is a wrapper around the C function gtk_table_new.
func TableNew(rows uint32, columns uint32, homogeneous bool) *Table {
	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	c_homogeneous :=
		boolToGboolean(homogeneous)

	retC := C.gtk_table_new(c_rows, c_columns, c_homogeneous)
	retGo := TableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Attach is a wrapper around the C function gtk_table_attach.
func (recv *Table) Attach(child *Widget, leftAttach uint32, rightAttach uint32, topAttach uint32, bottomAttach uint32, xoptions AttachOptions, yoptions AttachOptions, xpadding uint32, ypadding uint32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_left_attach := (C.guint)(leftAttach)

	c_right_attach := (C.guint)(rightAttach)

	c_top_attach := (C.guint)(topAttach)

	c_bottom_attach := (C.guint)(bottomAttach)

	c_xoptions := (C.GtkAttachOptions)(xoptions)

	c_yoptions := (C.GtkAttachOptions)(yoptions)

	c_xpadding := (C.guint)(xpadding)

	c_ypadding := (C.guint)(ypadding)

	C.gtk_table_attach((*C.GtkTable)(recv.native), c_child, c_left_attach, c_right_attach, c_top_attach, c_bottom_attach, c_xoptions, c_yoptions, c_xpadding, c_ypadding)

	return
}

// AttachDefaults is a wrapper around the C function gtk_table_attach_defaults.
func (recv *Table) AttachDefaults(widget *Widget, leftAttach uint32, rightAttach uint32, topAttach uint32, bottomAttach uint32) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	c_left_attach := (C.guint)(leftAttach)

	c_right_attach := (C.guint)(rightAttach)

	c_top_attach := (C.guint)(topAttach)

	c_bottom_attach := (C.guint)(bottomAttach)

	C.gtk_table_attach_defaults((*C.GtkTable)(recv.native), c_widget, c_left_attach, c_right_attach, c_top_attach, c_bottom_attach)

	return
}

// GetColSpacing is a wrapper around the C function gtk_table_get_col_spacing.
func (recv *Table) GetColSpacing(column uint32) uint32 {
	c_column := (C.guint)(column)

	retC := C.gtk_table_get_col_spacing((*C.GtkTable)(recv.native), c_column)
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultColSpacing is a wrapper around the C function gtk_table_get_default_col_spacing.
func (recv *Table) GetDefaultColSpacing() uint32 {
	retC := C.gtk_table_get_default_col_spacing((*C.GtkTable)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetDefaultRowSpacing is a wrapper around the C function gtk_table_get_default_row_spacing.
func (recv *Table) GetDefaultRowSpacing() uint32 {
	retC := C.gtk_table_get_default_row_spacing((*C.GtkTable)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_table_get_homogeneous.
func (recv *Table) GetHomogeneous() bool {
	retC := C.gtk_table_get_homogeneous((*C.GtkTable)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRowSpacing is a wrapper around the C function gtk_table_get_row_spacing.
func (recv *Table) GetRowSpacing(row uint32) uint32 {
	c_row := (C.guint)(row)

	retC := C.gtk_table_get_row_spacing((*C.GtkTable)(recv.native), c_row)
	retGo := (uint32)(retC)

	return retGo
}

// Resize is a wrapper around the C function gtk_table_resize.
func (recv *Table) Resize(rows uint32, columns uint32) {
	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	C.gtk_table_resize((*C.GtkTable)(recv.native), c_rows, c_columns)

	return
}

// SetColSpacing is a wrapper around the C function gtk_table_set_col_spacing.
func (recv *Table) SetColSpacing(column uint32, spacing uint32) {
	c_column := (C.guint)(column)

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_col_spacing((*C.GtkTable)(recv.native), c_column, c_spacing)

	return
}

// SetColSpacings is a wrapper around the C function gtk_table_set_col_spacings.
func (recv *Table) SetColSpacings(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_col_spacings((*C.GtkTable)(recv.native), c_spacing)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_table_set_homogeneous.
func (recv *Table) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_table_set_homogeneous((*C.GtkTable)(recv.native), c_homogeneous)

	return
}

// SetRowSpacing is a wrapper around the C function gtk_table_set_row_spacing.
func (recv *Table) SetRowSpacing(row uint32, spacing uint32) {
	c_row := (C.guint)(row)

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_row_spacing((*C.GtkTable)(recv.native), c_row, c_spacing)

	return
}

// SetRowSpacings is a wrapper around the C function gtk_table_set_row_spacings.
func (recv *Table) SetRowSpacings(spacing uint32) {
	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_row_spacings((*C.GtkTable)(recv.native), c_spacing)

	return
}

// TearoffMenuItem is a wrapper around the C record GtkTearoffMenuItem.
type TearoffMenuItem struct {
	native *C.GtkTearoffMenuItem
	// menu_item : record
	// Private : priv
}

func TearoffMenuItemNewFromC(u unsafe.Pointer) *TearoffMenuItem {
	c := (*C.GtkTearoffMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &TearoffMenuItem{native: c}

	return g
}

func (recv *TearoffMenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem upcasts to *MenuItem
func (recv *TearoffMenuItem) MenuItem() *MenuItem {
	return MenuItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *TearoffMenuItem) Bin() *Bin {
	return recv.MenuItem().Bin()
}

// Container upcasts to *Container
func (recv *TearoffMenuItem) Container() *Container {
	return recv.MenuItem().Container()
}

// Widget upcasts to *Widget
func (recv *TearoffMenuItem) Widget() *Widget {
	return recv.MenuItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *TearoffMenuItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.MenuItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *TearoffMenuItem) Object() *gobject.Object {
	return recv.MenuItem().Object()
}

// CastToWidget down casts any arbitary Object to TearoffMenuItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a TearoffMenuItem.
func CastToTearoffMenuItem(object *gobject.Object) *TearoffMenuItem {
	return TearoffMenuItemNewFromC(object.ToC())
}

// TearoffMenuItemNew is a wrapper around the C function gtk_tearoff_menu_item_new.
func TearoffMenuItemNew() *TearoffMenuItem {
	retC := C.gtk_tearoff_menu_item_new()
	retGo := TearoffMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextBuffer is a wrapper around the C record GtkTextBuffer.
type TextBuffer struct {
	native *C.GtkTextBuffer
	// parent_instance : record
	// priv : record
}

func TextBufferNewFromC(u unsafe.Pointer) *TextBuffer {
	c := (*C.GtkTextBuffer)(u)
	if c == nil {
		return nil
	}

	g := &TextBuffer{native: c}

	return g
}

func (recv *TextBuffer) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TextBuffer) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TextBuffer.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextBuffer.
func CastToTextBuffer(object *gobject.Object) *TextBuffer {
	return TextBufferNewFromC(object.ToC())
}

var signalApplyTagId int
var signalApplyTagMap = make(map[int]TextBufferSignalApplyTagCallback)
var signalApplyTagLock sync.Mutex

// TextBufferSignalApplyTagCallback is a callback function for a 'apply-tag' signal emitted from a TextBuffer.
type TextBufferSignalApplyTagCallback func(tag *TextTag, start *TextIter, end *TextIter)

func TextBuffer_applyTagHandler() {}

var signalBeginUserActionId int
var signalBeginUserActionMap = make(map[int]TextBufferSignalBeginUserActionCallback)
var signalBeginUserActionLock sync.Mutex

// TextBufferSignalBeginUserActionCallback is a callback function for a 'begin-user-action' signal emitted from a TextBuffer.
type TextBufferSignalBeginUserActionCallback func()

func TextBuffer_beginUserActionHandler() {}

var signalChangedId int
var signalChangedMap = make(map[int]TextBufferSignalChangedCallback)
var signalChangedLock sync.Mutex

// TextBufferSignalChangedCallback is a callback function for a 'changed' signal emitted from a TextBuffer.
type TextBufferSignalChangedCallback func()

func TextBuffer_changedHandler() {}

var signalDeleteRangeId int
var signalDeleteRangeMap = make(map[int]TextBufferSignalDeleteRangeCallback)
var signalDeleteRangeLock sync.Mutex

// TextBufferSignalDeleteRangeCallback is a callback function for a 'delete-range' signal emitted from a TextBuffer.
type TextBufferSignalDeleteRangeCallback func(start *TextIter, end *TextIter)

func TextBuffer_deleteRangeHandler() {}

var signalEndUserActionId int
var signalEndUserActionMap = make(map[int]TextBufferSignalEndUserActionCallback)
var signalEndUserActionLock sync.Mutex

// TextBufferSignalEndUserActionCallback is a callback function for a 'end-user-action' signal emitted from a TextBuffer.
type TextBufferSignalEndUserActionCallback func()

func TextBuffer_endUserActionHandler() {}

var signalInsertChildAnchorId int
var signalInsertChildAnchorMap = make(map[int]TextBufferSignalInsertChildAnchorCallback)
var signalInsertChildAnchorLock sync.Mutex

// TextBufferSignalInsertChildAnchorCallback is a callback function for a 'insert-child-anchor' signal emitted from a TextBuffer.
type TextBufferSignalInsertChildAnchorCallback func(location *TextIter, anchor *TextChildAnchor)

func TextBuffer_insertChildAnchorHandler() {}

var signalInsertPixbufId int
var signalInsertPixbufMap = make(map[int]TextBufferSignalInsertPixbufCallback)
var signalInsertPixbufLock sync.Mutex

// TextBufferSignalInsertPixbufCallback is a callback function for a 'insert-pixbuf' signal emitted from a TextBuffer.
type TextBufferSignalInsertPixbufCallback func(location *TextIter, pixbuf *gdkpixbuf.Pixbuf)

func TextBuffer_insertPixbufHandler() {}

var signalInsertTextId int
var signalInsertTextMap = make(map[int]TextBufferSignalInsertTextCallback)
var signalInsertTextLock sync.Mutex

// TextBufferSignalInsertTextCallback is a callback function for a 'insert-text' signal emitted from a TextBuffer.
type TextBufferSignalInsertTextCallback func(location *TextIter, text string, len int32)

func TextBuffer_insertTextHandler() {}

var signalMarkDeletedId int
var signalMarkDeletedMap = make(map[int]TextBufferSignalMarkDeletedCallback)
var signalMarkDeletedLock sync.Mutex

// TextBufferSignalMarkDeletedCallback is a callback function for a 'mark-deleted' signal emitted from a TextBuffer.
type TextBufferSignalMarkDeletedCallback func(mark *TextMark)

func TextBuffer_markDeletedHandler() {}

var signalMarkSetId int
var signalMarkSetMap = make(map[int]TextBufferSignalMarkSetCallback)
var signalMarkSetLock sync.Mutex

// TextBufferSignalMarkSetCallback is a callback function for a 'mark-set' signal emitted from a TextBuffer.
type TextBufferSignalMarkSetCallback func(location *TextIter, mark *TextMark)

func TextBuffer_markSetHandler() {}

var signalModifiedChangedId int
var signalModifiedChangedMap = make(map[int]TextBufferSignalModifiedChangedCallback)
var signalModifiedChangedLock sync.Mutex

// TextBufferSignalModifiedChangedCallback is a callback function for a 'modified-changed' signal emitted from a TextBuffer.
type TextBufferSignalModifiedChangedCallback func()

func TextBuffer_modifiedChangedHandler() {}

var signalRemoveTagId int
var signalRemoveTagMap = make(map[int]TextBufferSignalRemoveTagCallback)
var signalRemoveTagLock sync.Mutex

// TextBufferSignalRemoveTagCallback is a callback function for a 'remove-tag' signal emitted from a TextBuffer.
type TextBufferSignalRemoveTagCallback func(tag *TextTag, start *TextIter, end *TextIter)

func TextBuffer_removeTagHandler() {}

// TextBufferNew is a wrapper around the C function gtk_text_buffer_new.
func TextBufferNew(table *TextTagTable) *TextBuffer {
	c_table := (*C.GtkTextTagTable)(table.ToC())

	retC := C.gtk_text_buffer_new(c_table)
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddSelectionClipboard is a wrapper around the C function gtk_text_buffer_add_selection_clipboard.
func (recv *TextBuffer) AddSelectionClipboard(clipboard *Clipboard) {
	c_clipboard := (*C.GtkClipboard)(clipboard.ToC())

	C.gtk_text_buffer_add_selection_clipboard((*C.GtkTextBuffer)(recv.native), c_clipboard)

	return
}

// ApplyTag is a wrapper around the C function gtk_text_buffer_apply_tag.
func (recv *TextBuffer) ApplyTag(tag *TextTag, start *TextIter, end *TextIter) {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_apply_tag((*C.GtkTextBuffer)(recv.native), c_tag, c_start, c_end)

	return
}

// ApplyTagByName is a wrapper around the C function gtk_text_buffer_apply_tag_by_name.
func (recv *TextBuffer) ApplyTagByName(name string, start *TextIter, end *TextIter) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_apply_tag_by_name((*C.GtkTextBuffer)(recv.native), c_name, c_start, c_end)

	return
}

// BeginUserAction is a wrapper around the C function gtk_text_buffer_begin_user_action.
func (recv *TextBuffer) BeginUserAction() {
	C.gtk_text_buffer_begin_user_action((*C.GtkTextBuffer)(recv.native))

	return
}

// CopyClipboard is a wrapper around the C function gtk_text_buffer_copy_clipboard.
func (recv *TextBuffer) CopyClipboard(clipboard *Clipboard) {
	c_clipboard := (*C.GtkClipboard)(clipboard.ToC())

	C.gtk_text_buffer_copy_clipboard((*C.GtkTextBuffer)(recv.native), c_clipboard)

	return
}

// CreateChildAnchor is a wrapper around the C function gtk_text_buffer_create_child_anchor.
func (recv *TextBuffer) CreateChildAnchor(iter *TextIter) *TextChildAnchor {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_buffer_create_child_anchor((*C.GtkTextBuffer)(recv.native), c_iter)
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateMark is a wrapper around the C function gtk_text_buffer_create_mark.
func (recv *TextBuffer) CreateMark(markName string, where *TextIter, leftGravity bool) *TextMark {
	c_mark_name := C.CString(markName)
	defer C.free(unsafe.Pointer(c_mark_name))

	c_where := (*C.GtkTextIter)(where.ToC())

	c_left_gravity :=
		boolToGboolean(leftGravity)

	retC := C.gtk_text_buffer_create_mark((*C.GtkTextBuffer)(recv.native), c_mark_name, c_where, c_left_gravity)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_buffer_create_tag : unsupported parameter ... : varargs

// CutClipboard is a wrapper around the C function gtk_text_buffer_cut_clipboard.
func (recv *TextBuffer) CutClipboard(clipboard *Clipboard, defaultEditable bool) {
	c_clipboard := (*C.GtkClipboard)(clipboard.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	C.gtk_text_buffer_cut_clipboard((*C.GtkTextBuffer)(recv.native), c_clipboard, c_default_editable)

	return
}

// Delete is a wrapper around the C function gtk_text_buffer_delete.
func (recv *TextBuffer) Delete(start *TextIter, end *TextIter) {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_delete((*C.GtkTextBuffer)(recv.native), c_start, c_end)

	return
}

// DeleteInteractive is a wrapper around the C function gtk_text_buffer_delete_interactive.
func (recv *TextBuffer) DeleteInteractive(startIter *TextIter, endIter *TextIter, defaultEditable bool) bool {
	c_start_iter := (*C.GtkTextIter)(startIter.ToC())

	c_end_iter := (*C.GtkTextIter)(endIter.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_delete_interactive((*C.GtkTextBuffer)(recv.native), c_start_iter, c_end_iter, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// DeleteMark is a wrapper around the C function gtk_text_buffer_delete_mark.
func (recv *TextBuffer) DeleteMark(mark *TextMark) {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	C.gtk_text_buffer_delete_mark((*C.GtkTextBuffer)(recv.native), c_mark)

	return
}

// DeleteMarkByName is a wrapper around the C function gtk_text_buffer_delete_mark_by_name.
func (recv *TextBuffer) DeleteMarkByName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_text_buffer_delete_mark_by_name((*C.GtkTextBuffer)(recv.native), c_name)

	return
}

// DeleteSelection is a wrapper around the C function gtk_text_buffer_delete_selection.
func (recv *TextBuffer) DeleteSelection(interactive bool, defaultEditable bool) bool {
	c_interactive :=
		boolToGboolean(interactive)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_delete_selection((*C.GtkTextBuffer)(recv.native), c_interactive, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// EndUserAction is a wrapper around the C function gtk_text_buffer_end_user_action.
func (recv *TextBuffer) EndUserAction() {
	C.gtk_text_buffer_end_user_action((*C.GtkTextBuffer)(recv.native))

	return
}

// GetBounds is a wrapper around the C function gtk_text_buffer_get_bounds.
func (recv *TextBuffer) GetBounds() (*TextIter, *TextIter) {
	var c_start C.GtkTextIter

	var c_end C.GtkTextIter

	C.gtk_text_buffer_get_bounds((*C.GtkTextBuffer)(recv.native), &c_start, &c_end)

	start := TextIterNewFromC(unsafe.Pointer(&c_start))

	end := TextIterNewFromC(unsafe.Pointer(&c_end))

	return start, end
}

// GetCharCount is a wrapper around the C function gtk_text_buffer_get_char_count.
func (recv *TextBuffer) GetCharCount() int32 {
	retC := C.gtk_text_buffer_get_char_count((*C.GtkTextBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetEndIter is a wrapper around the C function gtk_text_buffer_get_end_iter.
func (recv *TextBuffer) GetEndIter() *TextIter {
	var c_iter C.GtkTextIter

	C.gtk_text_buffer_get_end_iter((*C.GtkTextBuffer)(recv.native), &c_iter)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetInsert is a wrapper around the C function gtk_text_buffer_get_insert.
func (recv *TextBuffer) GetInsert() *TextMark {
	retC := C.gtk_text_buffer_get_insert((*C.GtkTextBuffer)(recv.native))
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIterAtChildAnchor is a wrapper around the C function gtk_text_buffer_get_iter_at_child_anchor.
func (recv *TextBuffer) GetIterAtChildAnchor(anchor *TextChildAnchor) *TextIter {
	var c_iter C.GtkTextIter

	c_anchor := (*C.GtkTextChildAnchor)(anchor.ToC())

	C.gtk_text_buffer_get_iter_at_child_anchor((*C.GtkTextBuffer)(recv.native), &c_iter, c_anchor)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetIterAtLine is a wrapper around the C function gtk_text_buffer_get_iter_at_line.
func (recv *TextBuffer) GetIterAtLine(lineNumber int32) *TextIter {
	var c_iter C.GtkTextIter

	c_line_number := (C.gint)(lineNumber)

	C.gtk_text_buffer_get_iter_at_line((*C.GtkTextBuffer)(recv.native), &c_iter, c_line_number)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetIterAtLineIndex is a wrapper around the C function gtk_text_buffer_get_iter_at_line_index.
func (recv *TextBuffer) GetIterAtLineIndex(lineNumber int32, byteIndex int32) *TextIter {
	var c_iter C.GtkTextIter

	c_line_number := (C.gint)(lineNumber)

	c_byte_index := (C.gint)(byteIndex)

	C.gtk_text_buffer_get_iter_at_line_index((*C.GtkTextBuffer)(recv.native), &c_iter, c_line_number, c_byte_index)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetIterAtLineOffset is a wrapper around the C function gtk_text_buffer_get_iter_at_line_offset.
func (recv *TextBuffer) GetIterAtLineOffset(lineNumber int32, charOffset int32) *TextIter {
	var c_iter C.GtkTextIter

	c_line_number := (C.gint)(lineNumber)

	c_char_offset := (C.gint)(charOffset)

	C.gtk_text_buffer_get_iter_at_line_offset((*C.GtkTextBuffer)(recv.native), &c_iter, c_line_number, c_char_offset)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetIterAtMark is a wrapper around the C function gtk_text_buffer_get_iter_at_mark.
func (recv *TextBuffer) GetIterAtMark(mark *TextMark) *TextIter {
	var c_iter C.GtkTextIter

	c_mark := (*C.GtkTextMark)(mark.ToC())

	C.gtk_text_buffer_get_iter_at_mark((*C.GtkTextBuffer)(recv.native), &c_iter, c_mark)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetIterAtOffset is a wrapper around the C function gtk_text_buffer_get_iter_at_offset.
func (recv *TextBuffer) GetIterAtOffset(charOffset int32) *TextIter {
	var c_iter C.GtkTextIter

	c_char_offset := (C.gint)(charOffset)

	C.gtk_text_buffer_get_iter_at_offset((*C.GtkTextBuffer)(recv.native), &c_iter, c_char_offset)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetLineCount is a wrapper around the C function gtk_text_buffer_get_line_count.
func (recv *TextBuffer) GetLineCount() int32 {
	retC := C.gtk_text_buffer_get_line_count((*C.GtkTextBuffer)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMark is a wrapper around the C function gtk_text_buffer_get_mark.
func (recv *TextBuffer) GetMark(name string) *TextMark {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_buffer_get_mark((*C.GtkTextBuffer)(recv.native), c_name)
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetModified is a wrapper around the C function gtk_text_buffer_get_modified.
func (recv *TextBuffer) GetModified() bool {
	retC := C.gtk_text_buffer_get_modified((*C.GtkTextBuffer)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectionBound is a wrapper around the C function gtk_text_buffer_get_selection_bound.
func (recv *TextBuffer) GetSelectionBound() *TextMark {
	retC := C.gtk_text_buffer_get_selection_bound((*C.GtkTextBuffer)(recv.native))
	retGo := TextMarkNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSelectionBounds is a wrapper around the C function gtk_text_buffer_get_selection_bounds.
func (recv *TextBuffer) GetSelectionBounds() (bool, *TextIter, *TextIter) {
	var c_start C.GtkTextIter

	var c_end C.GtkTextIter

	retC := C.gtk_text_buffer_get_selection_bounds((*C.GtkTextBuffer)(recv.native), &c_start, &c_end)
	retGo := retC == C.TRUE

	start := TextIterNewFromC(unsafe.Pointer(&c_start))

	end := TextIterNewFromC(unsafe.Pointer(&c_end))

	return retGo, start, end
}

// GetSlice is a wrapper around the C function gtk_text_buffer_get_slice.
func (recv *TextBuffer) GetSlice(start *TextIter, end *TextIter, includeHiddenChars bool) string {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_include_hidden_chars :=
		boolToGboolean(includeHiddenChars)

	retC := C.gtk_text_buffer_get_slice((*C.GtkTextBuffer)(recv.native), c_start, c_end, c_include_hidden_chars)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetStartIter is a wrapper around the C function gtk_text_buffer_get_start_iter.
func (recv *TextBuffer) GetStartIter() *TextIter {
	var c_iter C.GtkTextIter

	C.gtk_text_buffer_get_start_iter((*C.GtkTextBuffer)(recv.native), &c_iter)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// GetTagTable is a wrapper around the C function gtk_text_buffer_get_tag_table.
func (recv *TextBuffer) GetTagTable() *TextTagTable {
	retC := C.gtk_text_buffer_get_tag_table((*C.GtkTextBuffer)(recv.native))
	retGo := TextTagTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetText is a wrapper around the C function gtk_text_buffer_get_text.
func (recv *TextBuffer) GetText(start *TextIter, end *TextIter, includeHiddenChars bool) string {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_include_hidden_chars :=
		boolToGboolean(includeHiddenChars)

	retC := C.gtk_text_buffer_get_text((*C.GtkTextBuffer)(recv.native), c_start, c_end, c_include_hidden_chars)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Insert is a wrapper around the C function gtk_text_buffer_insert.
func (recv *TextBuffer) Insert(iter *TextIter, text string, len int32) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_insert((*C.GtkTextBuffer)(recv.native), c_iter, c_text, c_len)

	return
}

// InsertAtCursor is a wrapper around the C function gtk_text_buffer_insert_at_cursor.
func (recv *TextBuffer) InsertAtCursor(text string, len int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_insert_at_cursor((*C.GtkTextBuffer)(recv.native), c_text, c_len)

	return
}

// InsertChildAnchor is a wrapper around the C function gtk_text_buffer_insert_child_anchor.
func (recv *TextBuffer) InsertChildAnchor(iter *TextIter, anchor *TextChildAnchor) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_anchor := (*C.GtkTextChildAnchor)(anchor.ToC())

	C.gtk_text_buffer_insert_child_anchor((*C.GtkTextBuffer)(recv.native), c_iter, c_anchor)

	return
}

// InsertInteractive is a wrapper around the C function gtk_text_buffer_insert_interactive.
func (recv *TextBuffer) InsertInteractive(iter *TextIter, text string, len int32, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_interactive((*C.GtkTextBuffer)(recv.native), c_iter, c_text, c_len, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// InsertInteractiveAtCursor is a wrapper around the C function gtk_text_buffer_insert_interactive_at_cursor.
func (recv *TextBuffer) InsertInteractiveAtCursor(text string, len int32, defaultEditable bool) bool {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_interactive_at_cursor((*C.GtkTextBuffer)(recv.native), c_text, c_len, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// InsertPixbuf is a wrapper around the C function gtk_text_buffer_insert_pixbuf.
func (recv *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *gdkpixbuf.Pixbuf) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_text_buffer_insert_pixbuf((*C.GtkTextBuffer)(recv.native), c_iter, c_pixbuf)

	return
}

// InsertRange is a wrapper around the C function gtk_text_buffer_insert_range.
func (recv *TextBuffer) InsertRange(iter *TextIter, start *TextIter, end *TextIter) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_insert_range((*C.GtkTextBuffer)(recv.native), c_iter, c_start, c_end)

	return
}

// InsertRangeInteractive is a wrapper around the C function gtk_text_buffer_insert_range_interactive.
func (recv *TextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, defaultEditable bool) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	retC := C.gtk_text_buffer_insert_range_interactive((*C.GtkTextBuffer)(recv.native), c_iter, c_start, c_end, c_default_editable)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_buffer_insert_with_tags : unsupported parameter ... : varargs

// Unsupported : gtk_text_buffer_insert_with_tags_by_name : unsupported parameter ... : varargs

// MoveMark is a wrapper around the C function gtk_text_buffer_move_mark.
func (recv *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	c_where := (*C.GtkTextIter)(where.ToC())

	C.gtk_text_buffer_move_mark((*C.GtkTextBuffer)(recv.native), c_mark, c_where)

	return
}

// MoveMarkByName is a wrapper around the C function gtk_text_buffer_move_mark_by_name.
func (recv *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_where := (*C.GtkTextIter)(where.ToC())

	C.gtk_text_buffer_move_mark_by_name((*C.GtkTextBuffer)(recv.native), c_name, c_where)

	return
}

// PasteClipboard is a wrapper around the C function gtk_text_buffer_paste_clipboard.
func (recv *TextBuffer) PasteClipboard(clipboard *Clipboard, overrideLocation *TextIter, defaultEditable bool) {
	c_clipboard := (*C.GtkClipboard)(clipboard.ToC())

	c_override_location := (*C.GtkTextIter)(overrideLocation.ToC())

	c_default_editable :=
		boolToGboolean(defaultEditable)

	C.gtk_text_buffer_paste_clipboard((*C.GtkTextBuffer)(recv.native), c_clipboard, c_override_location, c_default_editable)

	return
}

// PlaceCursor is a wrapper around the C function gtk_text_buffer_place_cursor.
func (recv *TextBuffer) PlaceCursor(where *TextIter) {
	c_where := (*C.GtkTextIter)(where.ToC())

	C.gtk_text_buffer_place_cursor((*C.GtkTextBuffer)(recv.native), c_where)

	return
}

// RemoveAllTags is a wrapper around the C function gtk_text_buffer_remove_all_tags.
func (recv *TextBuffer) RemoveAllTags(start *TextIter, end *TextIter) {
	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_remove_all_tags((*C.GtkTextBuffer)(recv.native), c_start, c_end)

	return
}

// RemoveSelectionClipboard is a wrapper around the C function gtk_text_buffer_remove_selection_clipboard.
func (recv *TextBuffer) RemoveSelectionClipboard(clipboard *Clipboard) {
	c_clipboard := (*C.GtkClipboard)(clipboard.ToC())

	C.gtk_text_buffer_remove_selection_clipboard((*C.GtkTextBuffer)(recv.native), c_clipboard)

	return
}

// RemoveTag is a wrapper around the C function gtk_text_buffer_remove_tag.
func (recv *TextBuffer) RemoveTag(tag *TextTag, start *TextIter, end *TextIter) {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_remove_tag((*C.GtkTextBuffer)(recv.native), c_tag, c_start, c_end)

	return
}

// RemoveTagByName is a wrapper around the C function gtk_text_buffer_remove_tag_by_name.
func (recv *TextBuffer) RemoveTagByName(name string, start *TextIter, end *TextIter) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_start := (*C.GtkTextIter)(start.ToC())

	c_end := (*C.GtkTextIter)(end.ToC())

	C.gtk_text_buffer_remove_tag_by_name((*C.GtkTextBuffer)(recv.native), c_name, c_start, c_end)

	return
}

// SetModified is a wrapper around the C function gtk_text_buffer_set_modified.
func (recv *TextBuffer) SetModified(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_text_buffer_set_modified((*C.GtkTextBuffer)(recv.native), c_setting)

	return
}

// SetText is a wrapper around the C function gtk_text_buffer_set_text.
func (recv *TextBuffer) SetText(text string, len int32) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	c_len := (C.gint)(len)

	C.gtk_text_buffer_set_text((*C.GtkTextBuffer)(recv.native), c_text, c_len)

	return
}

// TextCellAccessible is a wrapper around the C record GtkTextCellAccessible.
type TextCellAccessible struct {
	native *C.GtkTextCellAccessible
	// parent : record
	// priv : record
}

func TextCellAccessibleNewFromC(u unsafe.Pointer) *TextCellAccessible {
	c := (*C.GtkTextCellAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TextCellAccessible{native: c}

	return g
}

func (recv *TextCellAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RendererCellAccessible upcasts to *RendererCellAccessible
func (recv *TextCellAccessible) RendererCellAccessible() *RendererCellAccessible {
	return RendererCellAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// CellAccessible upcasts to *CellAccessible
func (recv *TextCellAccessible) CellAccessible() *CellAccessible {
	return recv.RendererCellAccessible().CellAccessible()
}

// Accessible upcasts to *Accessible
func (recv *TextCellAccessible) Accessible() *Accessible {
	return recv.RendererCellAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *TextCellAccessible) Object() *atk.Object {
	return recv.RendererCellAccessible().Object()
}

// CastToWidget down casts any arbitary Object to TextCellAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextCellAccessible.
func CastToTextCellAccessible(object *gobject.Object) *TextCellAccessible {
	return TextCellAccessibleNewFromC(object.ToC())
}

// TextChildAnchor is a wrapper around the C record GtkTextChildAnchor.
type TextChildAnchor struct {
	native *C.GtkTextChildAnchor
	// parent_instance : record
	// Private : segment
}

func TextChildAnchorNewFromC(u unsafe.Pointer) *TextChildAnchor {
	c := (*C.GtkTextChildAnchor)(u)
	if c == nil {
		return nil
	}

	g := &TextChildAnchor{native: c}

	return g
}

func (recv *TextChildAnchor) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TextChildAnchor) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TextChildAnchor.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextChildAnchor.
func CastToTextChildAnchor(object *gobject.Object) *TextChildAnchor {
	return TextChildAnchorNewFromC(object.ToC())
}

// TextChildAnchorNew is a wrapper around the C function gtk_text_child_anchor_new.
func TextChildAnchorNew() *TextChildAnchor {
	retC := C.gtk_text_child_anchor_new()
	retGo := TextChildAnchorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeleted is a wrapper around the C function gtk_text_child_anchor_get_deleted.
func (recv *TextChildAnchor) GetDeleted() bool {
	retC := C.gtk_text_child_anchor_get_deleted((*C.GtkTextChildAnchor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidgets is a wrapper around the C function gtk_text_child_anchor_get_widgets.
func (recv *TextChildAnchor) GetWidgets() *glib.List {
	retC := C.gtk_text_child_anchor_get_widgets((*C.GtkTextChildAnchor)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextMark is a wrapper around the C record GtkTextMark.
type TextMark struct {
	native *C.GtkTextMark
	// parent_instance : record
	// Private : segment
}

func TextMarkNewFromC(u unsafe.Pointer) *TextMark {
	c := (*C.GtkTextMark)(u)
	if c == nil {
		return nil
	}

	g := &TextMark{native: c}

	return g
}

func (recv *TextMark) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TextMark) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TextMark.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextMark.
func CastToTextMark(object *gobject.Object) *TextMark {
	return TextMarkNewFromC(object.ToC())
}

// GetBuffer is a wrapper around the C function gtk_text_mark_get_buffer.
func (recv *TextMark) GetBuffer() *TextBuffer {
	retC := C.gtk_text_mark_get_buffer((*C.GtkTextMark)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDeleted is a wrapper around the C function gtk_text_mark_get_deleted.
func (recv *TextMark) GetDeleted() bool {
	retC := C.gtk_text_mark_get_deleted((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLeftGravity is a wrapper around the C function gtk_text_mark_get_left_gravity.
func (recv *TextMark) GetLeftGravity() bool {
	retC := C.gtk_text_mark_get_left_gravity((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function gtk_text_mark_get_name.
func (recv *TextMark) GetName() string {
	retC := C.gtk_text_mark_get_name((*C.GtkTextMark)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_text_mark_get_visible.
func (recv *TextMark) GetVisible() bool {
	retC := C.gtk_text_mark_get_visible((*C.GtkTextMark)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetVisible is a wrapper around the C function gtk_text_mark_set_visible.
func (recv *TextMark) SetVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_text_mark_set_visible((*C.GtkTextMark)(recv.native), c_setting)

	return
}

// TextTag is a wrapper around the C record GtkTextTag.
type TextTag struct {
	native *C.GtkTextTag
	// parent_instance : record
	// priv : record
}

func TextTagNewFromC(u unsafe.Pointer) *TextTag {
	c := (*C.GtkTextTag)(u)
	if c == nil {
		return nil
	}

	g := &TextTag{native: c}

	return g
}

func (recv *TextTag) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TextTag) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TextTag.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextTag.
func CastToTextTag(object *gobject.Object) *TextTag {
	return TextTagNewFromC(object.ToC())
}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// TextTagNew is a wrapper around the C function gtk_text_tag_new.
func TextTagNew(name string) *TextTag {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_tag_new(c_name)
	retGo := TextTagNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_tag_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// GetPriority is a wrapper around the C function gtk_text_tag_get_priority.
func (recv *TextTag) GetPriority() int32 {
	retC := C.gtk_text_tag_get_priority((*C.GtkTextTag)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// SetPriority is a wrapper around the C function gtk_text_tag_set_priority.
func (recv *TextTag) SetPriority(priority int32) {
	c_priority := (C.gint)(priority)

	C.gtk_text_tag_set_priority((*C.GtkTextTag)(recv.native), c_priority)

	return
}

// TextTagTable is a wrapper around the C record GtkTextTagTable.
type TextTagTable struct {
	native *C.GtkTextTagTable
	// parent_instance : record
	// priv : record
}

func TextTagTableNewFromC(u unsafe.Pointer) *TextTagTable {
	c := (*C.GtkTextTagTable)(u)
	if c == nil {
		return nil
	}

	g := &TextTagTable{native: c}

	return g
}

func (recv *TextTagTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TextTagTable) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TextTagTable.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextTagTable.
func CastToTextTagTable(object *gobject.Object) *TextTagTable {
	return TextTagTableNewFromC(object.ToC())
}

var signalTagAddedId int
var signalTagAddedMap = make(map[int]TextTagTableSignalTagAddedCallback)
var signalTagAddedLock sync.Mutex

// TextTagTableSignalTagAddedCallback is a callback function for a 'tag-added' signal emitted from a TextTagTable.
type TextTagTableSignalTagAddedCallback func(tag *TextTag)

func TextTagTable_tagAddedHandler() {}

var signalTagChangedId int
var signalTagChangedMap = make(map[int]TextTagTableSignalTagChangedCallback)
var signalTagChangedLock sync.Mutex

// TextTagTableSignalTagChangedCallback is a callback function for a 'tag-changed' signal emitted from a TextTagTable.
type TextTagTableSignalTagChangedCallback func(tag *TextTag, sizeChanged bool)

func TextTagTable_tagChangedHandler() {}

var signalTagRemovedId int
var signalTagRemovedMap = make(map[int]TextTagTableSignalTagRemovedCallback)
var signalTagRemovedLock sync.Mutex

// TextTagTableSignalTagRemovedCallback is a callback function for a 'tag-removed' signal emitted from a TextTagTable.
type TextTagTableSignalTagRemovedCallback func(tag *TextTag)

func TextTagTable_tagRemovedHandler() {}

// TextTagTableNew is a wrapper around the C function gtk_text_tag_table_new.
func TextTagTableNew() *TextTagTable {
	retC := C.gtk_text_tag_table_new()
	retGo := TextTagTableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Add is a wrapper around the C function gtk_text_tag_table_add.
func (recv *TextTagTable) Add(tag *TextTag) bool {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	retC := C.gtk_text_tag_table_add((*C.GtkTextTagTable)(recv.native), c_tag)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_text_tag_table_foreach : unsupported parameter func : no type generator for TextTagTableForeach, GtkTextTagTableForeach

// GetSize is a wrapper around the C function gtk_text_tag_table_get_size.
func (recv *TextTagTable) GetSize() int32 {
	retC := C.gtk_text_tag_table_get_size((*C.GtkTextTagTable)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Lookup is a wrapper around the C function gtk_text_tag_table_lookup.
func (recv *TextTagTable) Lookup(name string) *TextTag {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_text_tag_table_lookup((*C.GtkTextTagTable)(recv.native), c_name)
	retGo := TextTagNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Remove is a wrapper around the C function gtk_text_tag_table_remove.
func (recv *TextTagTable) Remove(tag *TextTag) {
	c_tag := (*C.GtkTextTag)(tag.ToC())

	C.gtk_text_tag_table_remove((*C.GtkTextTagTable)(recv.native), c_tag)

	return
}

// TextView is a wrapper around the C record GtkTextView.
type TextView struct {
	native *C.GtkTextView
	// parent_instance : record
	// Private : priv
}

func TextViewNewFromC(u unsafe.Pointer) *TextView {
	c := (*C.GtkTextView)(u)
	if c == nil {
		return nil
	}

	g := &TextView{native: c}

	return g
}

func (recv *TextView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *TextView) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *TextView) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *TextView) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *TextView) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to TextView.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextView.
func CastToTextView(object *gobject.Object) *TextView {
	return TextViewNewFromC(object.ToC())
}

var signalBackspaceId int
var signalBackspaceMap = make(map[int]TextViewSignalBackspaceCallback)
var signalBackspaceLock sync.Mutex

// TextViewSignalBackspaceCallback is a callback function for a 'backspace' signal emitted from a TextView.
type TextViewSignalBackspaceCallback func()

func TextView_backspaceHandler() {}

var signalCopyClipboardId int
var signalCopyClipboardMap = make(map[int]TextViewSignalCopyClipboardCallback)
var signalCopyClipboardLock sync.Mutex

// TextViewSignalCopyClipboardCallback is a callback function for a 'copy-clipboard' signal emitted from a TextView.
type TextViewSignalCopyClipboardCallback func()

func TextView_copyClipboardHandler() {}

var signalCutClipboardId int
var signalCutClipboardMap = make(map[int]TextViewSignalCutClipboardCallback)
var signalCutClipboardLock sync.Mutex

// TextViewSignalCutClipboardCallback is a callback function for a 'cut-clipboard' signal emitted from a TextView.
type TextViewSignalCutClipboardCallback func()

func TextView_cutClipboardHandler() {}

var signalDeleteFromCursorId int
var signalDeleteFromCursorMap = make(map[int]TextViewSignalDeleteFromCursorCallback)
var signalDeleteFromCursorLock sync.Mutex

// TextViewSignalDeleteFromCursorCallback is a callback function for a 'delete-from-cursor' signal emitted from a TextView.
type TextViewSignalDeleteFromCursorCallback func(type_ DeleteType, count int32)

func TextView_deleteFromCursorHandler() {}

var signalInsertAtCursorId int
var signalInsertAtCursorMap = make(map[int]TextViewSignalInsertAtCursorCallback)
var signalInsertAtCursorLock sync.Mutex

// TextViewSignalInsertAtCursorCallback is a callback function for a 'insert-at-cursor' signal emitted from a TextView.
type TextViewSignalInsertAtCursorCallback func(string string)

func TextView_insertAtCursorHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]TextViewSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// TextViewSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a TextView.
type TextViewSignalMoveCursorCallback func(step MovementStep, count int32, extendSelection bool)

func TextView_moveCursorHandler() {}

var signalMoveViewportId int
var signalMoveViewportMap = make(map[int]TextViewSignalMoveViewportCallback)
var signalMoveViewportLock sync.Mutex

// TextViewSignalMoveViewportCallback is a callback function for a 'move-viewport' signal emitted from a TextView.
type TextViewSignalMoveViewportCallback func(step ScrollStep, count int32)

func TextView_moveViewportHandler() {}

var signalPasteClipboardId int
var signalPasteClipboardMap = make(map[int]TextViewSignalPasteClipboardCallback)
var signalPasteClipboardLock sync.Mutex

// TextViewSignalPasteClipboardCallback is a callback function for a 'paste-clipboard' signal emitted from a TextView.
type TextViewSignalPasteClipboardCallback func()

func TextView_pasteClipboardHandler() {}

var signalPopulatePopupId int
var signalPopulatePopupMap = make(map[int]TextViewSignalPopulatePopupCallback)
var signalPopulatePopupLock sync.Mutex

// TextViewSignalPopulatePopupCallback is a callback function for a 'populate-popup' signal emitted from a TextView.
type TextViewSignalPopulatePopupCallback func(popup *Widget)

func TextView_populatePopupHandler() {}

var signalSelectAllId int
var signalSelectAllMap = make(map[int]TextViewSignalSelectAllCallback)
var signalSelectAllLock sync.Mutex

// TextViewSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a TextView.
type TextViewSignalSelectAllCallback func(select_ bool)

func TextView_selectAllHandler() {}

var signalSetAnchorId int
var signalSetAnchorMap = make(map[int]TextViewSignalSetAnchorCallback)
var signalSetAnchorLock sync.Mutex

// TextViewSignalSetAnchorCallback is a callback function for a 'set-anchor' signal emitted from a TextView.
type TextViewSignalSetAnchorCallback func()

func TextView_setAnchorHandler() {}

var signalToggleCursorVisibleId int
var signalToggleCursorVisibleMap = make(map[int]TextViewSignalToggleCursorVisibleCallback)
var signalToggleCursorVisibleLock sync.Mutex

// TextViewSignalToggleCursorVisibleCallback is a callback function for a 'toggle-cursor-visible' signal emitted from a TextView.
type TextViewSignalToggleCursorVisibleCallback func()

func TextView_toggleCursorVisibleHandler() {}

var signalToggleOverwriteId int
var signalToggleOverwriteMap = make(map[int]TextViewSignalToggleOverwriteCallback)
var signalToggleOverwriteLock sync.Mutex

// TextViewSignalToggleOverwriteCallback is a callback function for a 'toggle-overwrite' signal emitted from a TextView.
type TextViewSignalToggleOverwriteCallback func()

func TextView_toggleOverwriteHandler() {}

// TextViewNew is a wrapper around the C function gtk_text_view_new.
func TextViewNew() *TextView {
	retC := C.gtk_text_view_new()
	retGo := TextViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TextViewNewWithBuffer is a wrapper around the C function gtk_text_view_new_with_buffer.
func TextViewNewWithBuffer(buffer *TextBuffer) *TextView {
	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	retC := C.gtk_text_view_new_with_buffer(c_buffer)
	retGo := TextViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddChildAtAnchor is a wrapper around the C function gtk_text_view_add_child_at_anchor.
func (recv *TextView) AddChildAtAnchor(child *Widget, anchor *TextChildAnchor) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_anchor := (*C.GtkTextChildAnchor)(anchor.ToC())

	C.gtk_text_view_add_child_at_anchor((*C.GtkTextView)(recv.native), c_child, c_anchor)

	return
}

// AddChildInWindow is a wrapper around the C function gtk_text_view_add_child_in_window.
func (recv *TextView) AddChildInWindow(child *Widget, whichWindow TextWindowType, xpos int32, ypos int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_which_window := (C.GtkTextWindowType)(whichWindow)

	c_xpos := (C.gint)(xpos)

	c_ypos := (C.gint)(ypos)

	C.gtk_text_view_add_child_in_window((*C.GtkTextView)(recv.native), c_child, c_which_window, c_xpos, c_ypos)

	return
}

// BackwardDisplayLine is a wrapper around the C function gtk_text_view_backward_display_line.
func (recv *TextView) BackwardDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_backward_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardDisplayLineStart is a wrapper around the C function gtk_text_view_backward_display_line_start.
func (recv *TextView) BackwardDisplayLineStart(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_backward_display_line_start((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// BufferToWindowCoords is a wrapper around the C function gtk_text_view_buffer_to_window_coords.
func (recv *TextView) BufferToWindowCoords(win TextWindowType, bufferX int32, bufferY int32) (int32, int32) {
	c_win := (C.GtkTextWindowType)(win)

	c_buffer_x := (C.gint)(bufferX)

	c_buffer_y := (C.gint)(bufferY)

	var c_window_x C.gint

	var c_window_y C.gint

	C.gtk_text_view_buffer_to_window_coords((*C.GtkTextView)(recv.native), c_win, c_buffer_x, c_buffer_y, &c_window_x, &c_window_y)

	windowX := (int32)(c_window_x)

	windowY := (int32)(c_window_y)

	return windowX, windowY
}

// ForwardDisplayLine is a wrapper around the C function gtk_text_view_forward_display_line.
func (recv *TextView) ForwardDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_forward_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardDisplayLineEnd is a wrapper around the C function gtk_text_view_forward_display_line_end.
func (recv *TextView) ForwardDisplayLineEnd(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_forward_display_line_end((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// GetBorderWindowSize is a wrapper around the C function gtk_text_view_get_border_window_size.
func (recv *TextView) GetBorderWindowSize(type_ TextWindowType) int32 {
	c_type := (C.GtkTextWindowType)(type_)

	retC := C.gtk_text_view_get_border_window_size((*C.GtkTextView)(recv.native), c_type)
	retGo := (int32)(retC)

	return retGo
}

// GetBuffer is a wrapper around the C function gtk_text_view_get_buffer.
func (recv *TextView) GetBuffer() *TextBuffer {
	retC := C.gtk_text_view_get_buffer((*C.GtkTextView)(recv.native))
	retGo := TextBufferNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetCursorVisible is a wrapper around the C function gtk_text_view_get_cursor_visible.
func (recv *TextView) GetCursorVisible() bool {
	retC := C.gtk_text_view_get_cursor_visible((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDefaultAttributes is a wrapper around the C function gtk_text_view_get_default_attributes.
func (recv *TextView) GetDefaultAttributes() *TextAttributes {
	retC := C.gtk_text_view_get_default_attributes((*C.GtkTextView)(recv.native))
	retGo := TextAttributesNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetEditable is a wrapper around the C function gtk_text_view_get_editable.
func (recv *TextView) GetEditable() bool {
	retC := C.gtk_text_view_get_editable((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIndent is a wrapper around the C function gtk_text_view_get_indent.
func (recv *TextView) GetIndent() int32 {
	retC := C.gtk_text_view_get_indent((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetIterAtLocation is a wrapper around the C function gtk_text_view_get_iter_at_location.
func (recv *TextView) GetIterAtLocation(x int32, y int32) *TextIter {
	var c_iter C.GtkTextIter

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_text_view_get_iter_at_location((*C.GtkTextView)(recv.native), &c_iter, c_x, c_y)

	iter := TextIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Unsupported : gtk_text_view_get_iter_location : unsupported parameter location : Blacklisted record : GdkRectangle

// GetJustification is a wrapper around the C function gtk_text_view_get_justification.
func (recv *TextView) GetJustification() Justification {
	retC := C.gtk_text_view_get_justification((*C.GtkTextView)(recv.native))
	retGo := (Justification)(retC)

	return retGo
}

// GetLeftMargin is a wrapper around the C function gtk_text_view_get_left_margin.
func (recv *TextView) GetLeftMargin() int32 {
	retC := C.gtk_text_view_get_left_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetLineAtY is a wrapper around the C function gtk_text_view_get_line_at_y.
func (recv *TextView) GetLineAtY(y int32) (*TextIter, int32) {
	var c_target_iter C.GtkTextIter

	c_y := (C.gint)(y)

	var c_line_top C.gint

	C.gtk_text_view_get_line_at_y((*C.GtkTextView)(recv.native), &c_target_iter, c_y, &c_line_top)

	targetIter := TextIterNewFromC(unsafe.Pointer(&c_target_iter))

	lineTop := (int32)(c_line_top)

	return targetIter, lineTop
}

// GetLineYrange is a wrapper around the C function gtk_text_view_get_line_yrange.
func (recv *TextView) GetLineYrange(iter *TextIter) (int32, int32) {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	var c_y C.gint

	var c_height C.gint

	C.gtk_text_view_get_line_yrange((*C.GtkTextView)(recv.native), c_iter, &c_y, &c_height)

	y := (int32)(c_y)

	height := (int32)(c_height)

	return y, height
}

// GetPixelsAboveLines is a wrapper around the C function gtk_text_view_get_pixels_above_lines.
func (recv *TextView) GetPixelsAboveLines() int32 {
	retC := C.gtk_text_view_get_pixels_above_lines((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixelsBelowLines is a wrapper around the C function gtk_text_view_get_pixels_below_lines.
func (recv *TextView) GetPixelsBelowLines() int32 {
	retC := C.gtk_text_view_get_pixels_below_lines((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetPixelsInsideWrap is a wrapper around the C function gtk_text_view_get_pixels_inside_wrap.
func (recv *TextView) GetPixelsInsideWrap() int32 {
	retC := C.gtk_text_view_get_pixels_inside_wrap((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetRightMargin is a wrapper around the C function gtk_text_view_get_right_margin.
func (recv *TextView) GetRightMargin() int32 {
	retC := C.gtk_text_view_get_right_margin((*C.GtkTextView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTabs is a wrapper around the C function gtk_text_view_get_tabs.
func (recv *TextView) GetTabs() *pango.TabArray {
	retC := C.gtk_text_view_get_tabs((*C.GtkTextView)(recv.native))
	retGo := pango.TabArrayNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_text_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// GetWindow is a wrapper around the C function gtk_text_view_get_window.
func (recv *TextView) GetWindow(win TextWindowType) *gdk.Window {
	c_win := (C.GtkTextWindowType)(win)

	retC := C.gtk_text_view_get_window((*C.GtkTextView)(recv.native), c_win)
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWindowType is a wrapper around the C function gtk_text_view_get_window_type.
func (recv *TextView) GetWindowType(window *gdk.Window) TextWindowType {
	c_window := (*C.GdkWindow)(window.ToC())

	retC := C.gtk_text_view_get_window_type((*C.GtkTextView)(recv.native), c_window)
	retGo := (TextWindowType)(retC)

	return retGo
}

// GetWrapMode is a wrapper around the C function gtk_text_view_get_wrap_mode.
func (recv *TextView) GetWrapMode() WrapMode {
	retC := C.gtk_text_view_get_wrap_mode((*C.GtkTextView)(recv.native))
	retGo := (WrapMode)(retC)

	return retGo
}

// MoveChild is a wrapper around the C function gtk_text_view_move_child.
func (recv *TextView) MoveChild(child *Widget, xpos int32, ypos int32) {
	c_child := (*C.GtkWidget)(child.ToC())

	c_xpos := (C.gint)(xpos)

	c_ypos := (C.gint)(ypos)

	C.gtk_text_view_move_child((*C.GtkTextView)(recv.native), c_child, c_xpos, c_ypos)

	return
}

// MoveMarkOnscreen is a wrapper around the C function gtk_text_view_move_mark_onscreen.
func (recv *TextView) MoveMarkOnscreen(mark *TextMark) bool {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	retC := C.gtk_text_view_move_mark_onscreen((*C.GtkTextView)(recv.native), c_mark)
	retGo := retC == C.TRUE

	return retGo
}

// MoveVisually is a wrapper around the C function gtk_text_view_move_visually.
func (recv *TextView) MoveVisually(iter *TextIter, count int32) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_count := (C.gint)(count)

	retC := C.gtk_text_view_move_visually((*C.GtkTextView)(recv.native), c_iter, c_count)
	retGo := retC == C.TRUE

	return retGo
}

// PlaceCursorOnscreen is a wrapper around the C function gtk_text_view_place_cursor_onscreen.
func (recv *TextView) PlaceCursorOnscreen() bool {
	retC := C.gtk_text_view_place_cursor_onscreen((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ScrollMarkOnscreen is a wrapper around the C function gtk_text_view_scroll_mark_onscreen.
func (recv *TextView) ScrollMarkOnscreen(mark *TextMark) {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	C.gtk_text_view_scroll_mark_onscreen((*C.GtkTextView)(recv.native), c_mark)

	return
}

// ScrollToIter is a wrapper around the C function gtk_text_view_scroll_to_iter.
func (recv *TextView) ScrollToIter(iter *TextIter, withinMargin float64, useAlign bool, xalign float64, yalign float64) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	c_within_margin := (C.gdouble)(withinMargin)

	c_use_align :=
		boolToGboolean(useAlign)

	c_xalign := (C.gdouble)(xalign)

	c_yalign := (C.gdouble)(yalign)

	retC := C.gtk_text_view_scroll_to_iter((*C.GtkTextView)(recv.native), c_iter, c_within_margin, c_use_align, c_xalign, c_yalign)
	retGo := retC == C.TRUE

	return retGo
}

// ScrollToMark is a wrapper around the C function gtk_text_view_scroll_to_mark.
func (recv *TextView) ScrollToMark(mark *TextMark, withinMargin float64, useAlign bool, xalign float64, yalign float64) {
	c_mark := (*C.GtkTextMark)(mark.ToC())

	c_within_margin := (C.gdouble)(withinMargin)

	c_use_align :=
		boolToGboolean(useAlign)

	c_xalign := (C.gdouble)(xalign)

	c_yalign := (C.gdouble)(yalign)

	C.gtk_text_view_scroll_to_mark((*C.GtkTextView)(recv.native), c_mark, c_within_margin, c_use_align, c_xalign, c_yalign)

	return
}

// SetBorderWindowSize is a wrapper around the C function gtk_text_view_set_border_window_size.
func (recv *TextView) SetBorderWindowSize(type_ TextWindowType, size int32) {
	c_type := (C.GtkTextWindowType)(type_)

	c_size := (C.gint)(size)

	C.gtk_text_view_set_border_window_size((*C.GtkTextView)(recv.native), c_type, c_size)

	return
}

// SetBuffer is a wrapper around the C function gtk_text_view_set_buffer.
func (recv *TextView) SetBuffer(buffer *TextBuffer) {
	c_buffer := (*C.GtkTextBuffer)(buffer.ToC())

	C.gtk_text_view_set_buffer((*C.GtkTextView)(recv.native), c_buffer)

	return
}

// SetCursorVisible is a wrapper around the C function gtk_text_view_set_cursor_visible.
func (recv *TextView) SetCursorVisible(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_text_view_set_cursor_visible((*C.GtkTextView)(recv.native), c_setting)

	return
}

// SetEditable is a wrapper around the C function gtk_text_view_set_editable.
func (recv *TextView) SetEditable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_text_view_set_editable((*C.GtkTextView)(recv.native), c_setting)

	return
}

// SetIndent is a wrapper around the C function gtk_text_view_set_indent.
func (recv *TextView) SetIndent(indent int32) {
	c_indent := (C.gint)(indent)

	C.gtk_text_view_set_indent((*C.GtkTextView)(recv.native), c_indent)

	return
}

// SetJustification is a wrapper around the C function gtk_text_view_set_justification.
func (recv *TextView) SetJustification(justification Justification) {
	c_justification := (C.GtkJustification)(justification)

	C.gtk_text_view_set_justification((*C.GtkTextView)(recv.native), c_justification)

	return
}

// SetLeftMargin is a wrapper around the C function gtk_text_view_set_left_margin.
func (recv *TextView) SetLeftMargin(leftMargin int32) {
	c_left_margin := (C.gint)(leftMargin)

	C.gtk_text_view_set_left_margin((*C.GtkTextView)(recv.native), c_left_margin)

	return
}

// SetPixelsAboveLines is a wrapper around the C function gtk_text_view_set_pixels_above_lines.
func (recv *TextView) SetPixelsAboveLines(pixelsAboveLines int32) {
	c_pixels_above_lines := (C.gint)(pixelsAboveLines)

	C.gtk_text_view_set_pixels_above_lines((*C.GtkTextView)(recv.native), c_pixels_above_lines)

	return
}

// SetPixelsBelowLines is a wrapper around the C function gtk_text_view_set_pixels_below_lines.
func (recv *TextView) SetPixelsBelowLines(pixelsBelowLines int32) {
	c_pixels_below_lines := (C.gint)(pixelsBelowLines)

	C.gtk_text_view_set_pixels_below_lines((*C.GtkTextView)(recv.native), c_pixels_below_lines)

	return
}

// SetPixelsInsideWrap is a wrapper around the C function gtk_text_view_set_pixels_inside_wrap.
func (recv *TextView) SetPixelsInsideWrap(pixelsInsideWrap int32) {
	c_pixels_inside_wrap := (C.gint)(pixelsInsideWrap)

	C.gtk_text_view_set_pixels_inside_wrap((*C.GtkTextView)(recv.native), c_pixels_inside_wrap)

	return
}

// SetRightMargin is a wrapper around the C function gtk_text_view_set_right_margin.
func (recv *TextView) SetRightMargin(rightMargin int32) {
	c_right_margin := (C.gint)(rightMargin)

	C.gtk_text_view_set_right_margin((*C.GtkTextView)(recv.native), c_right_margin)

	return
}

// SetTabs is a wrapper around the C function gtk_text_view_set_tabs.
func (recv *TextView) SetTabs(tabs *pango.TabArray) {
	c_tabs := (*C.PangoTabArray)(tabs.ToC())

	C.gtk_text_view_set_tabs((*C.GtkTextView)(recv.native), c_tabs)

	return
}

// SetWrapMode is a wrapper around the C function gtk_text_view_set_wrap_mode.
func (recv *TextView) SetWrapMode(wrapMode WrapMode) {
	c_wrap_mode := (C.GtkWrapMode)(wrapMode)

	C.gtk_text_view_set_wrap_mode((*C.GtkTextView)(recv.native), c_wrap_mode)

	return
}

// StartsDisplayLine is a wrapper around the C function gtk_text_view_starts_display_line.
func (recv *TextView) StartsDisplayLine(iter *TextIter) bool {
	c_iter := (*C.GtkTextIter)(iter.ToC())

	retC := C.gtk_text_view_starts_display_line((*C.GtkTextView)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// WindowToBufferCoords is a wrapper around the C function gtk_text_view_window_to_buffer_coords.
func (recv *TextView) WindowToBufferCoords(win TextWindowType, windowX int32, windowY int32) (int32, int32) {
	c_win := (C.GtkTextWindowType)(win)

	c_window_x := (C.gint)(windowX)

	c_window_y := (C.gint)(windowY)

	var c_buffer_x C.gint

	var c_buffer_y C.gint

	C.gtk_text_view_window_to_buffer_coords((*C.GtkTextView)(recv.native), c_win, c_window_x, c_window_y, &c_buffer_x, &c_buffer_y)

	bufferX := (int32)(c_buffer_x)

	bufferY := (int32)(c_buffer_y)

	return bufferX, bufferY
}

// TextViewAccessible is a wrapper around the C record GtkTextViewAccessible.
type TextViewAccessible struct {
	native *C.GtkTextViewAccessible
	// parent : record
	// priv : record
}

func TextViewAccessibleNewFromC(u unsafe.Pointer) *TextViewAccessible {
	c := (*C.GtkTextViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TextViewAccessible{native: c}

	return g
}

func (recv *TextViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *TextViewAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *TextViewAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *TextViewAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *TextViewAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to TextViewAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a TextViewAccessible.
func CastToTextViewAccessible(object *gobject.Object) *TextViewAccessible {
	return TextViewAccessibleNewFromC(object.ToC())
}

// ThemingEngine is a wrapper around the C record GtkThemingEngine.
type ThemingEngine struct {
	native *C.GtkThemingEngine
	// parent_object : record
	// priv : record
}

func ThemingEngineNewFromC(u unsafe.Pointer) *ThemingEngine {
	c := (*C.GtkThemingEngine)(u)
	if c == nil {
		return nil
	}

	g := &ThemingEngine{native: c}

	return g
}

func (recv *ThemingEngine) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ThemingEngine) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ThemingEngine.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThemingEngine.
func CastToThemingEngine(object *gobject.Object) *ThemingEngine {
	return ThemingEngineNewFromC(object.ToC())
}

// GetScreen is a wrapper around the C function gtk_theming_engine_get_screen.
func (recv *ThemingEngine) GetScreen() *gdk.Screen {
	retC := C.gtk_theming_engine_get_screen((*C.GtkThemingEngine)(recv.native))
	retGo := gdk.ScreenNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleAction is a wrapper around the C record GtkToggleAction.
type ToggleAction struct {
	native *C.GtkToggleAction
	// parent : record
	// Private : private_data
}

func ToggleActionNewFromC(u unsafe.Pointer) *ToggleAction {
	c := (*C.GtkToggleAction)(u)
	if c == nil {
		return nil
	}

	g := &ToggleAction{native: c}

	return g
}

func (recv *ToggleAction) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Action upcasts to *Action
func (recv *ToggleAction) Action() *Action {
	return ActionNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *ToggleAction) Object() *gobject.Object {
	return recv.Action().Object()
}

// CastToWidget down casts any arbitary Object to ToggleAction.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToggleAction.
func CastToToggleAction(object *gobject.Object) *ToggleAction {
	return ToggleActionNewFromC(object.ToC())
}

var signalToggledId int
var signalToggledMap = make(map[int]ToggleActionSignalToggledCallback)
var signalToggledLock sync.Mutex

// ToggleActionSignalToggledCallback is a callback function for a 'toggled' signal emitted from a ToggleAction.
type ToggleActionSignalToggledCallback func()

func ToggleAction_toggledHandler() {}

// ToggleButton is a wrapper around the C record GtkToggleButton.
type ToggleButton struct {
	native *C.GtkToggleButton
	// Private : button
	// Private : priv
}

func ToggleButtonNewFromC(u unsafe.Pointer) *ToggleButton {
	c := (*C.GtkToggleButton)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButton{native: c}

	return g
}

func (recv *ToggleButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Button upcasts to *Button
func (recv *ToggleButton) Button() *Button {
	return ButtonNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ToggleButton) Bin() *Bin {
	return recv.Button().Bin()
}

// Container upcasts to *Container
func (recv *ToggleButton) Container() *Container {
	return recv.Button().Container()
}

// Widget upcasts to *Widget
func (recv *ToggleButton) Widget() *Widget {
	return recv.Button().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToggleButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Button().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToggleButton) Object() *gobject.Object {
	return recv.Button().Object()
}

// CastToWidget down casts any arbitary Object to ToggleButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToggleButton.
func CastToToggleButton(object *gobject.Object) *ToggleButton {
	return ToggleButtonNewFromC(object.ToC())
}

var signalToggledId int
var signalToggledMap = make(map[int]ToggleButtonSignalToggledCallback)
var signalToggledLock sync.Mutex

// ToggleButtonSignalToggledCallback is a callback function for a 'toggled' signal emitted from a ToggleButton.
type ToggleButtonSignalToggledCallback func()

func ToggleButton_toggledHandler() {}

// ToggleButtonNew is a wrapper around the C function gtk_toggle_button_new.
func ToggleButtonNew() *ToggleButton {
	retC := C.gtk_toggle_button_new()
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithLabel is a wrapper around the C function gtk_toggle_button_new_with_label.
func ToggleButtonNewWithLabel(label string) *ToggleButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_label(c_label)
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleButtonNewWithMnemonic is a wrapper around the C function gtk_toggle_button_new_with_mnemonic.
func ToggleButtonNewWithMnemonic(label string) *ToggleButton {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_toggle_button_new_with_mnemonic(c_label)
	retGo := ToggleButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_button_get_active.
func (recv *ToggleButton) GetActive() bool {
	retC := C.gtk_toggle_button_get_active((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetInconsistent is a wrapper around the C function gtk_toggle_button_get_inconsistent.
func (recv *ToggleButton) GetInconsistent() bool {
	retC := C.gtk_toggle_button_get_inconsistent((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMode is a wrapper around the C function gtk_toggle_button_get_mode.
func (recv *ToggleButton) GetMode() bool {
	retC := C.gtk_toggle_button_get_mode((*C.GtkToggleButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_toggle_button_set_active.
func (recv *ToggleButton) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_button_set_active((*C.GtkToggleButton)(recv.native), c_is_active)

	return
}

// SetInconsistent is a wrapper around the C function gtk_toggle_button_set_inconsistent.
func (recv *ToggleButton) SetInconsistent(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_toggle_button_set_inconsistent((*C.GtkToggleButton)(recv.native), c_setting)

	return
}

// SetMode is a wrapper around the C function gtk_toggle_button_set_mode.
func (recv *ToggleButton) SetMode(drawIndicator bool) {
	c_draw_indicator :=
		boolToGboolean(drawIndicator)

	C.gtk_toggle_button_set_mode((*C.GtkToggleButton)(recv.native), c_draw_indicator)

	return
}

// Toggled is a wrapper around the C function gtk_toggle_button_toggled.
func (recv *ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled((*C.GtkToggleButton)(recv.native))

	return
}

// ToggleButtonAccessible is a wrapper around the C record GtkToggleButtonAccessible.
type ToggleButtonAccessible struct {
	native *C.GtkToggleButtonAccessible
	// parent : record
	// priv : record
}

func ToggleButtonAccessibleNewFromC(u unsafe.Pointer) *ToggleButtonAccessible {
	c := (*C.GtkToggleButtonAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ToggleButtonAccessible{native: c}

	return g
}

func (recv *ToggleButtonAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonAccessible upcasts to *ButtonAccessible
func (recv *ToggleButtonAccessible) ButtonAccessible() *ButtonAccessible {
	return ButtonAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *ToggleButtonAccessible) ContainerAccessible() *ContainerAccessible {
	return recv.ButtonAccessible().ContainerAccessible()
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *ToggleButtonAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ButtonAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *ToggleButtonAccessible) Accessible() *Accessible {
	return recv.ButtonAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *ToggleButtonAccessible) Object() *atk.Object {
	return recv.ButtonAccessible().Object()
}

// CastToWidget down casts any arbitary Object to ToggleButtonAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToggleButtonAccessible.
func CastToToggleButtonAccessible(object *gobject.Object) *ToggleButtonAccessible {
	return ToggleButtonAccessibleNewFromC(object.ToC())
}

// ToggleToolButton is a wrapper around the C record GtkToggleToolButton.
type ToggleToolButton struct {
	native *C.GtkToggleToolButton
	// parent : record
	// Private : priv
}

func ToggleToolButtonNewFromC(u unsafe.Pointer) *ToggleToolButton {
	c := (*C.GtkToggleToolButton)(u)
	if c == nil {
		return nil
	}

	g := &ToggleToolButton{native: c}

	return g
}

func (recv *ToggleToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolButton upcasts to *ToolButton
func (recv *ToggleToolButton) ToolButton() *ToolButton {
	return ToolButtonNewFromC(unsafe.Pointer(recv.native))
}

// ToolItem upcasts to *ToolItem
func (recv *ToggleToolButton) ToolItem() *ToolItem {
	return recv.ToolButton().ToolItem()
}

// Bin upcasts to *Bin
func (recv *ToggleToolButton) Bin() *Bin {
	return recv.ToolButton().Bin()
}

// Container upcasts to *Container
func (recv *ToggleToolButton) Container() *Container {
	return recv.ToolButton().Container()
}

// Widget upcasts to *Widget
func (recv *ToggleToolButton) Widget() *Widget {
	return recv.ToolButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToggleToolButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToolButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToggleToolButton) Object() *gobject.Object {
	return recv.ToolButton().Object()
}

// CastToWidget down casts any arbitary Object to ToggleToolButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToggleToolButton.
func CastToToggleToolButton(object *gobject.Object) *ToggleToolButton {
	return ToggleToolButtonNewFromC(object.ToC())
}

var signalToggledId int
var signalToggledMap = make(map[int]ToggleToolButtonSignalToggledCallback)
var signalToggledLock sync.Mutex

// ToggleToolButtonSignalToggledCallback is a callback function for a 'toggled' signal emitted from a ToggleToolButton.
type ToggleToolButtonSignalToggledCallback func()

func ToggleToolButton_toggledHandler() {}

// ToolButton is a wrapper around the C record GtkToolButton.
type ToolButton struct {
	native *C.GtkToolButton
	// parent : record
	// Private : priv
}

func ToolButtonNewFromC(u unsafe.Pointer) *ToolButton {
	c := (*C.GtkToolButton)(u)
	if c == nil {
		return nil
	}

	g := &ToolButton{native: c}

	return g
}

func (recv *ToolButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ToolItem upcasts to *ToolItem
func (recv *ToolButton) ToolItem() *ToolItem {
	return ToolItemNewFromC(unsafe.Pointer(recv.native))
}

// Bin upcasts to *Bin
func (recv *ToolButton) Bin() *Bin {
	return recv.ToolItem().Bin()
}

// Container upcasts to *Container
func (recv *ToolButton) Container() *Container {
	return recv.ToolItem().Container()
}

// Widget upcasts to *Widget
func (recv *ToolButton) Widget() *Widget {
	return recv.ToolItem().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToolButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ToolItem().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToolButton) Object() *gobject.Object {
	return recv.ToolItem().Object()
}

// CastToWidget down casts any arbitary Object to ToolButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToolButton.
func CastToToolButton(object *gobject.Object) *ToolButton {
	return ToolButtonNewFromC(object.ToC())
}

var signalClickedId int
var signalClickedMap = make(map[int]ToolButtonSignalClickedCallback)
var signalClickedLock sync.Mutex

// ToolButtonSignalClickedCallback is a callback function for a 'clicked' signal emitted from a ToolButton.
type ToolButtonSignalClickedCallback func()

func ToolButton_clickedHandler() {}

// ToolItem is a wrapper around the C record GtkToolItem.
type ToolItem struct {
	native *C.GtkToolItem
	// parent : record
	// Private : priv
}

func ToolItemNewFromC(u unsafe.Pointer) *ToolItem {
	c := (*C.GtkToolItem)(u)
	if c == nil {
		return nil
	}

	g := &ToolItem{native: c}

	return g
}

func (recv *ToolItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *ToolItem) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *ToolItem) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *ToolItem) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToolItem) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToolItem) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to ToolItem.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToolItem.
func CastToToolItem(object *gobject.Object) *ToolItem {
	return ToolItemNewFromC(object.ToC())
}

var signalCreateMenuProxyId int
var signalCreateMenuProxyMap = make(map[int]ToolItemSignalCreateMenuProxyCallback)
var signalCreateMenuProxyLock sync.Mutex

// ToolItemSignalCreateMenuProxyCallback is a callback function for a 'create-menu-proxy' signal emitted from a ToolItem.
type ToolItemSignalCreateMenuProxyCallback func() bool

func ToolItem_createMenuProxyHandler() {}

var signalToolbarReconfiguredId int
var signalToolbarReconfiguredMap = make(map[int]ToolItemSignalToolbarReconfiguredCallback)
var signalToolbarReconfiguredLock sync.Mutex

// ToolItemSignalToolbarReconfiguredCallback is a callback function for a 'toolbar-reconfigured' signal emitted from a ToolItem.
type ToolItemSignalToolbarReconfiguredCallback func()

func ToolItem_toolbarReconfiguredHandler() {}

// ToolItemGroup is a wrapper around the C record GtkToolItemGroup.
type ToolItemGroup struct {
	native *C.GtkToolItemGroup
	// parent_instance : record
	// priv : record
}

func ToolItemGroupNewFromC(u unsafe.Pointer) *ToolItemGroup {
	c := (*C.GtkToolItemGroup)(u)
	if c == nil {
		return nil
	}

	g := &ToolItemGroup{native: c}

	return g
}

func (recv *ToolItemGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *ToolItemGroup) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *ToolItemGroup) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToolItemGroup) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToolItemGroup) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to ToolItemGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToolItemGroup.
func CastToToolItemGroup(object *gobject.Object) *ToolItemGroup {
	return ToolItemGroupNewFromC(object.ToC())
}

// ToolPalette is a wrapper around the C record GtkToolPalette.
type ToolPalette struct {
	native *C.GtkToolPalette
	// parent_instance : record
	// priv : record
}

func ToolPaletteNewFromC(u unsafe.Pointer) *ToolPalette {
	c := (*C.GtkToolPalette)(u)
	if c == nil {
		return nil
	}

	g := &ToolPalette{native: c}

	return g
}

func (recv *ToolPalette) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *ToolPalette) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *ToolPalette) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *ToolPalette) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *ToolPalette) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to ToolPalette.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToolPalette.
func CastToToolPalette(object *gobject.Object) *ToolPalette {
	return ToolPaletteNewFromC(object.ToC())
}

// Toolbar is a wrapper around the C record GtkToolbar.
type Toolbar struct {
	native *C.GtkToolbar
	// container : record
	// priv : record
}

func ToolbarNewFromC(u unsafe.Pointer) *Toolbar {
	c := (*C.GtkToolbar)(u)
	if c == nil {
		return nil
	}

	g := &Toolbar{native: c}

	return g
}

func (recv *Toolbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *Toolbar) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *Toolbar) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Toolbar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Toolbar) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to Toolbar.
// Exercise care, as this is a potentially dangerous function if the Object is not a Toolbar.
func CastToToolbar(object *gobject.Object) *Toolbar {
	return ToolbarNewFromC(object.ToC())
}

var signalFocusHomeOrEndId int
var signalFocusHomeOrEndMap = make(map[int]ToolbarSignalFocusHomeOrEndCallback)
var signalFocusHomeOrEndLock sync.Mutex

// ToolbarSignalFocusHomeOrEndCallback is a callback function for a 'focus-home-or-end' signal emitted from a Toolbar.
type ToolbarSignalFocusHomeOrEndCallback func(focusHome bool) bool

func Toolbar_focusHomeOrEndHandler() {}

var signalOrientationChangedId int
var signalOrientationChangedMap = make(map[int]ToolbarSignalOrientationChangedCallback)
var signalOrientationChangedLock sync.Mutex

// ToolbarSignalOrientationChangedCallback is a callback function for a 'orientation-changed' signal emitted from a Toolbar.
type ToolbarSignalOrientationChangedCallback func(orientation Orientation)

func Toolbar_orientationChangedHandler() {}

var signalPopupContextMenuId int
var signalPopupContextMenuMap = make(map[int]ToolbarSignalPopupContextMenuCallback)
var signalPopupContextMenuLock sync.Mutex

// ToolbarSignalPopupContextMenuCallback is a callback function for a 'popup-context-menu' signal emitted from a Toolbar.
type ToolbarSignalPopupContextMenuCallback func(x int32, y int32, button int32) bool

func Toolbar_popupContextMenuHandler() {}

var signalStyleChangedId int
var signalStyleChangedMap = make(map[int]ToolbarSignalStyleChangedCallback)
var signalStyleChangedLock sync.Mutex

// ToolbarSignalStyleChangedCallback is a callback function for a 'style-changed' signal emitted from a Toolbar.
type ToolbarSignalStyleChangedCallback func(style ToolbarStyle)

func Toolbar_styleChangedHandler() {}

// ToolbarNew is a wrapper around the C function gtk_toolbar_new.
func ToolbarNew() *Toolbar {
	retC := C.gtk_toolbar_new()
	retGo := ToolbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconSize is a wrapper around the C function gtk_toolbar_get_icon_size.
func (recv *Toolbar) GetIconSize() IconSize {
	retC := C.gtk_toolbar_get_icon_size((*C.GtkToolbar)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// GetStyle is a wrapper around the C function gtk_toolbar_get_style.
func (recv *Toolbar) GetStyle() ToolbarStyle {
	retC := C.gtk_toolbar_get_style((*C.GtkToolbar)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// SetIconSize is a wrapper around the C function gtk_toolbar_set_icon_size.
func (recv *Toolbar) SetIconSize(iconSize IconSize) {
	c_icon_size := (C.GtkIconSize)(iconSize)

	C.gtk_toolbar_set_icon_size((*C.GtkToolbar)(recv.native), c_icon_size)

	return
}

// SetStyle is a wrapper around the C function gtk_toolbar_set_style.
func (recv *Toolbar) SetStyle(style ToolbarStyle) {
	c_style := (C.GtkToolbarStyle)(style)

	C.gtk_toolbar_set_style((*C.GtkToolbar)(recv.native), c_style)

	return
}

// UnsetIconSize is a wrapper around the C function gtk_toolbar_unset_icon_size.
func (recv *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size((*C.GtkToolbar)(recv.native))

	return
}

// UnsetStyle is a wrapper around the C function gtk_toolbar_unset_style.
func (recv *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style((*C.GtkToolbar)(recv.native))

	return
}

// Tooltip is a wrapper around the C record GtkTooltip.
type Tooltip struct {
	native *C.GtkTooltip
}

func TooltipNewFromC(u unsafe.Pointer) *Tooltip {
	c := (*C.GtkTooltip)(u)
	if c == nil {
		return nil
	}

	g := &Tooltip{native: c}

	return g
}

func (recv *Tooltip) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *Tooltip) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to Tooltip.
// Exercise care, as this is a potentially dangerous function if the Object is not a Tooltip.
func CastToTooltip(object *gobject.Object) *Tooltip {
	return TooltipNewFromC(object.ToC())
}

// ToplevelAccessible is a wrapper around the C record GtkToplevelAccessible.
type ToplevelAccessible struct {
	native *C.GtkToplevelAccessible
	// parent : record
	// priv : record
}

func ToplevelAccessibleNewFromC(u unsafe.Pointer) *ToplevelAccessible {
	c := (*C.GtkToplevelAccessible)(u)
	if c == nil {
		return nil
	}

	g := &ToplevelAccessible{native: c}

	return g
}

func (recv *ToplevelAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *ToplevelAccessible) Object() *atk.Object {
	return atk.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to ToplevelAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a ToplevelAccessible.
func CastToToplevelAccessible(object *gobject.Object) *ToplevelAccessible {
	return ToplevelAccessibleNewFromC(object.ToC())
}

// GetChildren is a wrapper around the C function gtk_toplevel_accessible_get_children.
func (recv *ToplevelAccessible) GetChildren() *glib.List {
	retC := C.gtk_toplevel_accessible_get_children((*C.GtkToplevelAccessible)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// TreeModelFilter is a wrapper around the C record GtkTreeModelFilter.
type TreeModelFilter struct {
	native *C.GtkTreeModelFilter
	// parent : record
	// Private : priv
}

func TreeModelFilterNewFromC(u unsafe.Pointer) *TreeModelFilter {
	c := (*C.GtkTreeModelFilter)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelFilter{native: c}

	return g
}

func (recv *TreeModelFilter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TreeModelFilter) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TreeModelFilter.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeModelFilter.
func CastToTreeModelFilter(object *gobject.Object) *TreeModelFilter {
	return TreeModelFilterNewFromC(object.ToC())
}

// TreeModelSort is a wrapper around the C record GtkTreeModelSort.
type TreeModelSort struct {
	native *C.GtkTreeModelSort
	// parent : record
	// Private : priv
}

func TreeModelSortNewFromC(u unsafe.Pointer) *TreeModelSort {
	c := (*C.GtkTreeModelSort)(u)
	if c == nil {
		return nil
	}

	g := &TreeModelSort{native: c}

	return g
}

func (recv *TreeModelSort) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TreeModelSort) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TreeModelSort.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeModelSort.
func CastToTreeModelSort(object *gobject.Object) *TreeModelSort {
	return TreeModelSortNewFromC(object.ToC())
}

// ClearCache is a wrapper around the C function gtk_tree_model_sort_clear_cache.
func (recv *TreeModelSort) ClearCache() {
	C.gtk_tree_model_sort_clear_cache((*C.GtkTreeModelSort)(recv.native))

	return
}

// ConvertChildIterToIter is a wrapper around the C function gtk_tree_model_sort_convert_child_iter_to_iter.
func (recv *TreeModelSort) ConvertChildIterToIter(childIter *TreeIter) (bool, *TreeIter) {
	var c_sort_iter C.GtkTreeIter

	c_child_iter := (*C.GtkTreeIter)(childIter.ToC())

	retC := C.gtk_tree_model_sort_convert_child_iter_to_iter((*C.GtkTreeModelSort)(recv.native), &c_sort_iter, c_child_iter)
	retGo := retC == C.TRUE

	sortIter := TreeIterNewFromC(unsafe.Pointer(&c_sort_iter))

	return retGo, sortIter
}

// ConvertChildPathToPath is a wrapper around the C function gtk_tree_model_sort_convert_child_path_to_path.
func (recv *TreeModelSort) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	c_child_path := (*C.GtkTreePath)(childPath.ToC())

	retC := C.gtk_tree_model_sort_convert_child_path_to_path((*C.GtkTreeModelSort)(recv.native), c_child_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ConvertIterToChildIter is a wrapper around the C function gtk_tree_model_sort_convert_iter_to_child_iter.
func (recv *TreeModelSort) ConvertIterToChildIter(sortedIter *TreeIter) *TreeIter {
	var c_child_iter C.GtkTreeIter

	c_sorted_iter := (*C.GtkTreeIter)(sortedIter.ToC())

	C.gtk_tree_model_sort_convert_iter_to_child_iter((*C.GtkTreeModelSort)(recv.native), &c_child_iter, c_sorted_iter)

	childIter := TreeIterNewFromC(unsafe.Pointer(&c_child_iter))

	return childIter
}

// ConvertPathToChildPath is a wrapper around the C function gtk_tree_model_sort_convert_path_to_child_path.
func (recv *TreeModelSort) ConvertPathToChildPath(sortedPath *TreePath) *TreePath {
	c_sorted_path := (*C.GtkTreePath)(sortedPath.ToC())

	retC := C.gtk_tree_model_sort_convert_path_to_child_path((*C.GtkTreeModelSort)(recv.native), c_sorted_path)
	retGo := TreePathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_model_sort_get_model : no return generator

// ResetDefaultSortFunc is a wrapper around the C function gtk_tree_model_sort_reset_default_sort_func.
func (recv *TreeModelSort) ResetDefaultSortFunc() {
	C.gtk_tree_model_sort_reset_default_sort_func((*C.GtkTreeModelSort)(recv.native))

	return
}

// TreeSelection is a wrapper around the C record GtkTreeSelection.
type TreeSelection struct {
	native *C.GtkTreeSelection
	// Private : parent
	// Private : priv
}

func TreeSelectionNewFromC(u unsafe.Pointer) *TreeSelection {
	c := (*C.GtkTreeSelection)(u)
	if c == nil {
		return nil
	}

	g := &TreeSelection{native: c}

	return g
}

func (recv *TreeSelection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TreeSelection) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TreeSelection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeSelection.
func CastToTreeSelection(object *gobject.Object) *TreeSelection {
	return TreeSelectionNewFromC(object.ToC())
}

var signalChangedId int
var signalChangedMap = make(map[int]TreeSelectionSignalChangedCallback)
var signalChangedLock sync.Mutex

// TreeSelectionSignalChangedCallback is a callback function for a 'changed' signal emitted from a TreeSelection.
type TreeSelectionSignalChangedCallback func()

func TreeSelection_changedHandler() {}

// GetMode is a wrapper around the C function gtk_tree_selection_get_mode.
func (recv *TreeSelection) GetMode() SelectionMode {
	retC := C.gtk_tree_selection_get_mode((*C.GtkTreeSelection)(recv.native))
	retGo := (SelectionMode)(retC)

	return retGo
}

// Unsupported : gtk_tree_selection_get_selected : unsupported parameter model : no type generator for TreeModel, GtkTreeModel**

// GetTreeView is a wrapper around the C function gtk_tree_selection_get_tree_view.
func (recv *TreeSelection) GetTreeView() *TreeView {
	retC := C.gtk_tree_selection_get_tree_view((*C.GtkTreeSelection)(recv.native))
	retGo := TreeViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUserData is a wrapper around the C function gtk_tree_selection_get_user_data.
func (recv *TreeSelection) GetUserData() uintptr {
	retC := C.gtk_tree_selection_get_user_data((*C.GtkTreeSelection)(recv.native))
	retGo := (uintptr)(unsafe.Pointer(retC))

	return retGo
}

// IterIsSelected is a wrapper around the C function gtk_tree_selection_iter_is_selected.
func (recv *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_selection_iter_is_selected((*C.GtkTreeSelection)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// PathIsSelected is a wrapper around the C function gtk_tree_selection_path_is_selected.
func (recv *TreeSelection) PathIsSelected(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_selection_path_is_selected((*C.GtkTreeSelection)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// SelectAll is a wrapper around the C function gtk_tree_selection_select_all.
func (recv *TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all((*C.GtkTreeSelection)(recv.native))

	return
}

// SelectIter is a wrapper around the C function gtk_tree_selection_select_iter.
func (recv *TreeSelection) SelectIter(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	C.gtk_tree_selection_select_iter((*C.GtkTreeSelection)(recv.native), c_iter)

	return
}

// SelectPath is a wrapper around the C function gtk_tree_selection_select_path.
func (recv *TreeSelection) SelectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_selection_select_path((*C.GtkTreeSelection)(recv.native), c_path)

	return
}

// SelectRange is a wrapper around the C function gtk_tree_selection_select_range.
func (recv *TreeSelection) SelectRange(startPath *TreePath, endPath *TreePath) {
	c_start_path := (*C.GtkTreePath)(startPath.ToC())

	c_end_path := (*C.GtkTreePath)(endPath.ToC())

	C.gtk_tree_selection_select_range((*C.GtkTreeSelection)(recv.native), c_start_path, c_end_path)

	return
}

// Unsupported : gtk_tree_selection_selected_foreach : unsupported parameter func : no type generator for TreeSelectionForeachFunc, GtkTreeSelectionForeachFunc

// SetMode is a wrapper around the C function gtk_tree_selection_set_mode.
func (recv *TreeSelection) SetMode(type_ SelectionMode) {
	c_type := (C.GtkSelectionMode)(type_)

	C.gtk_tree_selection_set_mode((*C.GtkTreeSelection)(recv.native), c_type)

	return
}

// Unsupported : gtk_tree_selection_set_select_function : unsupported parameter func : no type generator for TreeSelectionFunc, GtkTreeSelectionFunc

// UnselectAll is a wrapper around the C function gtk_tree_selection_unselect_all.
func (recv *TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all((*C.GtkTreeSelection)(recv.native))

	return
}

// UnselectIter is a wrapper around the C function gtk_tree_selection_unselect_iter.
func (recv *TreeSelection) UnselectIter(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	C.gtk_tree_selection_unselect_iter((*C.GtkTreeSelection)(recv.native), c_iter)

	return
}

// UnselectPath is a wrapper around the C function gtk_tree_selection_unselect_path.
func (recv *TreeSelection) UnselectPath(path *TreePath) {
	c_path := (*C.GtkTreePath)(path.ToC())

	C.gtk_tree_selection_unselect_path((*C.GtkTreeSelection)(recv.native), c_path)

	return
}

// TreeStore is a wrapper around the C record GtkTreeStore.
type TreeStore struct {
	native *C.GtkTreeStore
	// parent : record
	// priv : record
}

func TreeStoreNewFromC(u unsafe.Pointer) *TreeStore {
	c := (*C.GtkTreeStore)(u)
	if c == nil {
		return nil
	}

	g := &TreeStore{native: c}

	return g
}

func (recv *TreeStore) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *TreeStore) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to TreeStore.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeStore.
func CastToTreeStore(object *gobject.Object) *TreeStore {
	return TreeStoreNewFromC(object.ToC())
}

// Unsupported : gtk_tree_store_new : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_newv : unsupported parameter types : no param type

// Append is a wrapper around the C function gtk_tree_store_append.
func (recv *TreeStore) Append(parent *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(parent.ToC())

	C.gtk_tree_store_append((*C.GtkTreeStore)(recv.native), &c_iter, c_parent)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Clear is a wrapper around the C function gtk_tree_store_clear.
func (recv *TreeStore) Clear() {
	C.gtk_tree_store_clear((*C.GtkTreeStore)(recv.native))

	return
}

// Insert is a wrapper around the C function gtk_tree_store_insert.
func (recv *TreeStore) Insert(parent *TreeIter, position int32) *TreeIter {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(parent.ToC())

	c_position := (C.gint)(position)

	C.gtk_tree_store_insert((*C.GtkTreeStore)(recv.native), &c_iter, c_parent, c_position)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// InsertAfter is a wrapper around the C function gtk_tree_store_insert_after.
func (recv *TreeStore) InsertAfter(parent *TreeIter, sibling *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(parent.ToC())

	c_sibling := (*C.GtkTreeIter)(sibling.ToC())

	C.gtk_tree_store_insert_after((*C.GtkTreeStore)(recv.native), &c_iter, c_parent, c_sibling)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// InsertBefore is a wrapper around the C function gtk_tree_store_insert_before.
func (recv *TreeStore) InsertBefore(parent *TreeIter, sibling *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(parent.ToC())

	c_sibling := (*C.GtkTreeIter)(sibling.ToC())

	C.gtk_tree_store_insert_before((*C.GtkTreeStore)(recv.native), &c_iter, c_parent, c_sibling)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// IsAncestor is a wrapper around the C function gtk_tree_store_is_ancestor.
func (recv *TreeStore) IsAncestor(iter *TreeIter, descendant *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_descendant := (*C.GtkTreeIter)(descendant.ToC())

	retC := C.gtk_tree_store_is_ancestor((*C.GtkTreeStore)(recv.native), c_iter, c_descendant)
	retGo := retC == C.TRUE

	return retGo
}

// IterDepth is a wrapper around the C function gtk_tree_store_iter_depth.
func (recv *TreeStore) IterDepth(iter *TreeIter) int32 {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_iter_depth((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := (int32)(retC)

	return retGo
}

// Prepend is a wrapper around the C function gtk_tree_store_prepend.
func (recv *TreeStore) Prepend(parent *TreeIter) *TreeIter {
	var c_iter C.GtkTreeIter

	c_parent := (*C.GtkTreeIter)(parent.ToC())

	C.gtk_tree_store_prepend((*C.GtkTreeStore)(recv.native), &c_iter, c_parent)

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return iter
}

// Remove is a wrapper around the C function gtk_tree_store_remove.
func (recv *TreeStore) Remove(iter *TreeIter) bool {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	retC := C.gtk_tree_store_remove((*C.GtkTreeStore)(recv.native), c_iter)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_store_set : unsupported parameter ... : varargs

// Unsupported : gtk_tree_store_set_column_types : unsupported parameter types : no param type

// Unsupported : gtk_tree_store_set_valist : unsupported parameter var_args : no type generator for va_list, va_list

// SetValue is a wrapper around the C function gtk_tree_store_set_value.
func (recv *TreeStore) SetValue(iter *TreeIter, column int32, value *gobject.Value) {
	c_iter := (*C.GtkTreeIter)(iter.ToC())

	c_column := (C.gint)(column)

	c_value := (*C.GValue)(value.ToC())

	C.gtk_tree_store_set_value((*C.GtkTreeStore)(recv.native), c_iter, c_column, c_value)

	return
}

// TreeView is a wrapper around the C record GtkTreeView.
type TreeView struct {
	native *C.GtkTreeView
	// parent : record
	// Private : priv
}

func TreeViewNewFromC(u unsafe.Pointer) *TreeView {
	c := (*C.GtkTreeView)(u)
	if c == nil {
		return nil
	}

	g := &TreeView{native: c}

	return g
}

func (recv *TreeView) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Container upcasts to *Container
func (recv *TreeView) Container() *Container {
	return ContainerNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *TreeView) Widget() *Widget {
	return recv.Container().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *TreeView) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Container().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *TreeView) Object() *gobject.Object {
	return recv.Container().Object()
}

// CastToWidget down casts any arbitary Object to TreeView.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeView.
func CastToTreeView(object *gobject.Object) *TreeView {
	return TreeViewNewFromC(object.ToC())
}

var signalColumnsChangedId int
var signalColumnsChangedMap = make(map[int]TreeViewSignalColumnsChangedCallback)
var signalColumnsChangedLock sync.Mutex

// TreeViewSignalColumnsChangedCallback is a callback function for a 'columns-changed' signal emitted from a TreeView.
type TreeViewSignalColumnsChangedCallback func()

func TreeView_columnsChangedHandler() {}

var signalCursorChangedId int
var signalCursorChangedMap = make(map[int]TreeViewSignalCursorChangedCallback)
var signalCursorChangedLock sync.Mutex

// TreeViewSignalCursorChangedCallback is a callback function for a 'cursor-changed' signal emitted from a TreeView.
type TreeViewSignalCursorChangedCallback func()

func TreeView_cursorChangedHandler() {}

var signalExpandCollapseCursorRowId int
var signalExpandCollapseCursorRowMap = make(map[int]TreeViewSignalExpandCollapseCursorRowCallback)
var signalExpandCollapseCursorRowLock sync.Mutex

// TreeViewSignalExpandCollapseCursorRowCallback is a callback function for a 'expand-collapse-cursor-row' signal emitted from a TreeView.
type TreeViewSignalExpandCollapseCursorRowCallback func(object bool, p0 bool, p1 bool) bool

func TreeView_expandCollapseCursorRowHandler() {}

var signalMoveCursorId int
var signalMoveCursorMap = make(map[int]TreeViewSignalMoveCursorCallback)
var signalMoveCursorLock sync.Mutex

// TreeViewSignalMoveCursorCallback is a callback function for a 'move-cursor' signal emitted from a TreeView.
type TreeViewSignalMoveCursorCallback func(step MovementStep, direction int32) bool

func TreeView_moveCursorHandler() {}

var signalRowActivatedId int
var signalRowActivatedMap = make(map[int]TreeViewSignalRowActivatedCallback)
var signalRowActivatedLock sync.Mutex

// TreeViewSignalRowActivatedCallback is a callback function for a 'row-activated' signal emitted from a TreeView.
type TreeViewSignalRowActivatedCallback func(path *TreePath, column *TreeViewColumn)

func TreeView_rowActivatedHandler() {}

var signalRowCollapsedId int
var signalRowCollapsedMap = make(map[int]TreeViewSignalRowCollapsedCallback)
var signalRowCollapsedLock sync.Mutex

// TreeViewSignalRowCollapsedCallback is a callback function for a 'row-collapsed' signal emitted from a TreeView.
type TreeViewSignalRowCollapsedCallback func(iter *TreeIter, path *TreePath)

func TreeView_rowCollapsedHandler() {}

var signalRowExpandedId int
var signalRowExpandedMap = make(map[int]TreeViewSignalRowExpandedCallback)
var signalRowExpandedLock sync.Mutex

// TreeViewSignalRowExpandedCallback is a callback function for a 'row-expanded' signal emitted from a TreeView.
type TreeViewSignalRowExpandedCallback func(iter *TreeIter, path *TreePath)

func TreeView_rowExpandedHandler() {}

var signalSelectAllId int
var signalSelectAllMap = make(map[int]TreeViewSignalSelectAllCallback)
var signalSelectAllLock sync.Mutex

// TreeViewSignalSelectAllCallback is a callback function for a 'select-all' signal emitted from a TreeView.
type TreeViewSignalSelectAllCallback func() bool

func TreeView_selectAllHandler() {}

var signalSelectCursorParentId int
var signalSelectCursorParentMap = make(map[int]TreeViewSignalSelectCursorParentCallback)
var signalSelectCursorParentLock sync.Mutex

// TreeViewSignalSelectCursorParentCallback is a callback function for a 'select-cursor-parent' signal emitted from a TreeView.
type TreeViewSignalSelectCursorParentCallback func() bool

func TreeView_selectCursorParentHandler() {}

var signalSelectCursorRowId int
var signalSelectCursorRowMap = make(map[int]TreeViewSignalSelectCursorRowCallback)
var signalSelectCursorRowLock sync.Mutex

// TreeViewSignalSelectCursorRowCallback is a callback function for a 'select-cursor-row' signal emitted from a TreeView.
type TreeViewSignalSelectCursorRowCallback func(object bool) bool

func TreeView_selectCursorRowHandler() {}

var signalStartInteractiveSearchId int
var signalStartInteractiveSearchMap = make(map[int]TreeViewSignalStartInteractiveSearchCallback)
var signalStartInteractiveSearchLock sync.Mutex

// TreeViewSignalStartInteractiveSearchCallback is a callback function for a 'start-interactive-search' signal emitted from a TreeView.
type TreeViewSignalStartInteractiveSearchCallback func() bool

func TreeView_startInteractiveSearchHandler() {}

var signalTestCollapseRowId int
var signalTestCollapseRowMap = make(map[int]TreeViewSignalTestCollapseRowCallback)
var signalTestCollapseRowLock sync.Mutex

// TreeViewSignalTestCollapseRowCallback is a callback function for a 'test-collapse-row' signal emitted from a TreeView.
type TreeViewSignalTestCollapseRowCallback func(iter *TreeIter, path *TreePath) bool

func TreeView_testCollapseRowHandler() {}

var signalTestExpandRowId int
var signalTestExpandRowMap = make(map[int]TreeViewSignalTestExpandRowCallback)
var signalTestExpandRowLock sync.Mutex

// TreeViewSignalTestExpandRowCallback is a callback function for a 'test-expand-row' signal emitted from a TreeView.
type TreeViewSignalTestExpandRowCallback func(iter *TreeIter, path *TreePath) bool

func TreeView_testExpandRowHandler() {}

var signalToggleCursorRowId int
var signalToggleCursorRowMap = make(map[int]TreeViewSignalToggleCursorRowCallback)
var signalToggleCursorRowLock sync.Mutex

// TreeViewSignalToggleCursorRowCallback is a callback function for a 'toggle-cursor-row' signal emitted from a TreeView.
type TreeViewSignalToggleCursorRowCallback func() bool

func TreeView_toggleCursorRowHandler() {}

var signalUnselectAllId int
var signalUnselectAllMap = make(map[int]TreeViewSignalUnselectAllCallback)
var signalUnselectAllLock sync.Mutex

// TreeViewSignalUnselectAllCallback is a callback function for a 'unselect-all' signal emitted from a TreeView.
type TreeViewSignalUnselectAllCallback func() bool

func TreeView_unselectAllHandler() {}

// TreeViewNew is a wrapper around the C function gtk_tree_view_new.
func TreeViewNew() *TreeView {
	retC := C.gtk_tree_view_new()
	retGo := TreeViewNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_new_with_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// AppendColumn is a wrapper around the C function gtk_tree_view_append_column.
func (recv *TreeView) AppendColumn(column *TreeViewColumn) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	retC := C.gtk_tree_view_append_column((*C.GtkTreeView)(recv.native), c_column)
	retGo := (int32)(retC)

	return retGo
}

// CollapseAll is a wrapper around the C function gtk_tree_view_collapse_all.
func (recv *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all((*C.GtkTreeView)(recv.native))

	return
}

// CollapseRow is a wrapper around the C function gtk_tree_view_collapse_row.
func (recv *TreeView) CollapseRow(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_collapse_row((*C.GtkTreeView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// ColumnsAutosize is a wrapper around the C function gtk_tree_view_columns_autosize.
func (recv *TreeView) ColumnsAutosize() {
	C.gtk_tree_view_columns_autosize((*C.GtkTreeView)(recv.native))

	return
}

// CreateRowDragIcon is a wrapper around the C function gtk_tree_view_create_row_drag_icon.
func (recv *TreeView) CreateRowDragIcon(path *TreePath) *cairo.Surface {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_create_row_drag_icon((*C.GtkTreeView)(recv.native), c_path)
	retGo := cairo.SurfaceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_enable_model_drag_dest : unsupported parameter targets : no param type

// Unsupported : gtk_tree_view_enable_model_drag_source : unsupported parameter targets : no param type

// ExpandAll is a wrapper around the C function gtk_tree_view_expand_all.
func (recv *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all((*C.GtkTreeView)(recv.native))

	return
}

// ExpandRow is a wrapper around the C function gtk_tree_view_expand_row.
func (recv *TreeView) ExpandRow(path *TreePath, openAll bool) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_open_all :=
		boolToGboolean(openAll)

	retC := C.gtk_tree_view_expand_row((*C.GtkTreeView)(recv.native), c_path, c_open_all)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_background_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetBinWindow is a wrapper around the C function gtk_tree_view_get_bin_window.
func (recv *TreeView) GetBinWindow() *gdk.Window {
	retC := C.gtk_tree_view_get_bin_window((*C.GtkTreeView)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_cell_area : unsupported parameter rect : Blacklisted record : GdkRectangle

// GetColumn is a wrapper around the C function gtk_tree_view_get_column.
func (recv *TreeView) GetColumn(n int32) *TreeViewColumn {
	c_n := (C.gint)(n)

	retC := C.gtk_tree_view_get_column((*C.GtkTreeView)(recv.native), c_n)
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetColumns is a wrapper around the C function gtk_tree_view_get_columns.
func (recv *TreeView) GetColumns() *glib.List {
	retC := C.gtk_tree_view_get_columns((*C.GtkTreeView)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_cursor : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_dest_row_at_pos : unsupported parameter path : record with indirection level of 2

// Unsupported : gtk_tree_view_get_drag_dest_row : unsupported parameter path : record with indirection level of 2

// GetEnableSearch is a wrapper around the C function gtk_tree_view_get_enable_search.
func (recv *TreeView) GetEnableSearch() bool {
	retC := C.gtk_tree_view_get_enable_search((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetExpanderColumn is a wrapper around the C function gtk_tree_view_get_expander_column.
func (recv *TreeView) GetExpanderColumn() *TreeViewColumn {
	retC := C.gtk_tree_view_get_expander_column((*C.GtkTreeView)(recv.native))
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_tree_view_get_hadjustment.
func (recv *TreeView) GetHadjustment() *Adjustment {
	retC := C.gtk_tree_view_get_hadjustment((*C.GtkTreeView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHeadersVisible is a wrapper around the C function gtk_tree_view_get_headers_visible.
func (recv *TreeView) GetHeadersVisible() bool {
	retC := C.gtk_tree_view_get_headers_visible((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_get_model : no return generator

// Unsupported : gtk_tree_view_get_path_at_pos : unsupported parameter path : record with indirection level of 2

// GetReorderable is a wrapper around the C function gtk_tree_view_get_reorderable.
func (recv *TreeView) GetReorderable() bool {
	retC := C.gtk_tree_view_get_reorderable((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRulesHint is a wrapper around the C function gtk_tree_view_get_rules_hint.
func (recv *TreeView) GetRulesHint() bool {
	retC := C.gtk_tree_view_get_rules_hint((*C.GtkTreeView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSearchColumn is a wrapper around the C function gtk_tree_view_get_search_column.
func (recv *TreeView) GetSearchColumn() int32 {
	retC := C.gtk_tree_view_get_search_column((*C.GtkTreeView)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_get_search_equal_func : no return generator

// GetSelection is a wrapper around the C function gtk_tree_view_get_selection.
func (recv *TreeView) GetSelection() *TreeSelection {
	retC := C.gtk_tree_view_get_selection((*C.GtkTreeView)(recv.native))
	retGo := TreeSelectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_tree_view_get_vadjustment.
func (recv *TreeView) GetVadjustment() *Adjustment {
	retC := C.gtk_tree_view_get_vadjustment((*C.GtkTreeView)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_get_visible_rect : unsupported parameter visible_rect : Blacklisted record : GdkRectangle

// InsertColumn is a wrapper around the C function gtk_tree_view_insert_column.
func (recv *TreeView) InsertColumn(column *TreeViewColumn, position int32) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	c_position := (C.gint)(position)

	retC := C.gtk_tree_view_insert_column((*C.GtkTreeView)(recv.native), c_column, c_position)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_tree_view_insert_column_with_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_insert_column_with_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// Unsupported : gtk_tree_view_map_expanded_rows : unsupported parameter func : no type generator for TreeViewMappingFunc, GtkTreeViewMappingFunc

// MoveColumnAfter is a wrapper around the C function gtk_tree_view_move_column_after.
func (recv *TreeView) MoveColumnAfter(column *TreeViewColumn, baseColumn *TreeViewColumn) {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	c_base_column := (*C.GtkTreeViewColumn)(baseColumn.ToC())

	C.gtk_tree_view_move_column_after((*C.GtkTreeView)(recv.native), c_column, c_base_column)

	return
}

// RemoveColumn is a wrapper around the C function gtk_tree_view_remove_column.
func (recv *TreeView) RemoveColumn(column *TreeViewColumn) int32 {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	retC := C.gtk_tree_view_remove_column((*C.GtkTreeView)(recv.native), c_column)
	retGo := (int32)(retC)

	return retGo
}

// RowActivated is a wrapper around the C function gtk_tree_view_row_activated.
func (recv *TreeView) RowActivated(path *TreePath, column *TreeViewColumn) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	C.gtk_tree_view_row_activated((*C.GtkTreeView)(recv.native), c_path, c_column)

	return
}

// RowExpanded is a wrapper around the C function gtk_tree_view_row_expanded.
func (recv *TreeView) RowExpanded(path *TreePath) bool {
	c_path := (*C.GtkTreePath)(path.ToC())

	retC := C.gtk_tree_view_row_expanded((*C.GtkTreeView)(recv.native), c_path)
	retGo := retC == C.TRUE

	return retGo
}

// ScrollToCell is a wrapper around the C function gtk_tree_view_scroll_to_cell.
func (recv *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign bool, rowAlign float32, colAlign float32) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	c_use_align :=
		boolToGboolean(useAlign)

	c_row_align := (C.gfloat)(rowAlign)

	c_col_align := (C.gfloat)(colAlign)

	C.gtk_tree_view_scroll_to_cell((*C.GtkTreeView)(recv.native), c_path, c_column, c_use_align, c_row_align, c_col_align)

	return
}

// ScrollToPoint is a wrapper around the C function gtk_tree_view_scroll_to_point.
func (recv *TreeView) ScrollToPoint(treeX int32, treeY int32) {
	c_tree_x := (C.gint)(treeX)

	c_tree_y := (C.gint)(treeY)

	C.gtk_tree_view_scroll_to_point((*C.GtkTreeView)(recv.native), c_tree_x, c_tree_y)

	return
}

// Unsupported : gtk_tree_view_set_column_drag_function : unsupported parameter func : no type generator for TreeViewColumnDropFunc, GtkTreeViewColumnDropFunc

// SetCursor is a wrapper around the C function gtk_tree_view_set_cursor.
func (recv *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing bool) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_focus_column := (*C.GtkTreeViewColumn)(focusColumn.ToC())

	c_start_editing :=
		boolToGboolean(startEditing)

	C.gtk_tree_view_set_cursor((*C.GtkTreeView)(recv.native), c_path, c_focus_column, c_start_editing)

	return
}

// Unsupported : gtk_tree_view_set_destroy_count_func : unsupported parameter func : no type generator for TreeDestroyCountFunc, GtkTreeDestroyCountFunc

// SetDragDestRow is a wrapper around the C function gtk_tree_view_set_drag_dest_row.
func (recv *TreeView) SetDragDestRow(path *TreePath, pos TreeViewDropPosition) {
	c_path := (*C.GtkTreePath)(path.ToC())

	c_pos := (C.GtkTreeViewDropPosition)(pos)

	C.gtk_tree_view_set_drag_dest_row((*C.GtkTreeView)(recv.native), c_path, c_pos)

	return
}

// SetEnableSearch is a wrapper around the C function gtk_tree_view_set_enable_search.
func (recv *TreeView) SetEnableSearch(enableSearch bool) {
	c_enable_search :=
		boolToGboolean(enableSearch)

	C.gtk_tree_view_set_enable_search((*C.GtkTreeView)(recv.native), c_enable_search)

	return
}

// SetExpanderColumn is a wrapper around the C function gtk_tree_view_set_expander_column.
func (recv *TreeView) SetExpanderColumn(column *TreeViewColumn) {
	c_column := (*C.GtkTreeViewColumn)(column.ToC())

	C.gtk_tree_view_set_expander_column((*C.GtkTreeView)(recv.native), c_column)

	return
}

// SetHadjustment is a wrapper around the C function gtk_tree_view_set_hadjustment.
func (recv *TreeView) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_tree_view_set_hadjustment((*C.GtkTreeView)(recv.native), c_adjustment)

	return
}

// SetHeadersClickable is a wrapper around the C function gtk_tree_view_set_headers_clickable.
func (recv *TreeView) SetHeadersClickable(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_tree_view_set_headers_clickable((*C.GtkTreeView)(recv.native), c_setting)

	return
}

// SetHeadersVisible is a wrapper around the C function gtk_tree_view_set_headers_visible.
func (recv *TreeView) SetHeadersVisible(headersVisible bool) {
	c_headers_visible :=
		boolToGboolean(headersVisible)

	C.gtk_tree_view_set_headers_visible((*C.GtkTreeView)(recv.native), c_headers_visible)

	return
}

// Unsupported : gtk_tree_view_set_model : unsupported parameter model : no type generator for TreeModel, GtkTreeModel*

// SetReorderable is a wrapper around the C function gtk_tree_view_set_reorderable.
func (recv *TreeView) SetReorderable(reorderable bool) {
	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_tree_view_set_reorderable((*C.GtkTreeView)(recv.native), c_reorderable)

	return
}

// SetRulesHint is a wrapper around the C function gtk_tree_view_set_rules_hint.
func (recv *TreeView) SetRulesHint(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_tree_view_set_rules_hint((*C.GtkTreeView)(recv.native), c_setting)

	return
}

// SetSearchColumn is a wrapper around the C function gtk_tree_view_set_search_column.
func (recv *TreeView) SetSearchColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_view_set_search_column((*C.GtkTreeView)(recv.native), c_column)

	return
}

// Unsupported : gtk_tree_view_set_search_equal_func : unsupported parameter search_equal_func : no type generator for TreeViewSearchEqualFunc, GtkTreeViewSearchEqualFunc

// SetVadjustment is a wrapper around the C function gtk_tree_view_set_vadjustment.
func (recv *TreeView) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_tree_view_set_vadjustment((*C.GtkTreeView)(recv.native), c_adjustment)

	return
}

// UnsetRowsDragDest is a wrapper around the C function gtk_tree_view_unset_rows_drag_dest.
func (recv *TreeView) UnsetRowsDragDest() {
	C.gtk_tree_view_unset_rows_drag_dest((*C.GtkTreeView)(recv.native))

	return
}

// UnsetRowsDragSource is a wrapper around the C function gtk_tree_view_unset_rows_drag_source.
func (recv *TreeView) UnsetRowsDragSource() {
	C.gtk_tree_view_unset_rows_drag_source((*C.GtkTreeView)(recv.native))

	return
}

// TreeViewAccessible is a wrapper around the C record GtkTreeViewAccessible.
type TreeViewAccessible struct {
	native *C.GtkTreeViewAccessible
	// parent : record
	// priv : record
}

func TreeViewAccessibleNewFromC(u unsafe.Pointer) *TreeViewAccessible {
	c := (*C.GtkTreeViewAccessible)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewAccessible{native: c}

	return g
}

func (recv *TreeViewAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *TreeViewAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *TreeViewAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *TreeViewAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *TreeViewAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to TreeViewAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeViewAccessible.
func CastToTreeViewAccessible(object *gobject.Object) *TreeViewAccessible {
	return TreeViewAccessibleNewFromC(object.ToC())
}

// TreeViewColumn is a wrapper around the C record GtkTreeViewColumn.
type TreeViewColumn struct {
	native *C.GtkTreeViewColumn
	// parent_instance : record
	// priv : record
}

func TreeViewColumnNewFromC(u unsafe.Pointer) *TreeViewColumn {
	c := (*C.GtkTreeViewColumn)(u)
	if c == nil {
		return nil
	}

	g := &TreeViewColumn{native: c}

	return g
}

func (recv *TreeViewColumn) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *TreeViewColumn) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *TreeViewColumn) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to TreeViewColumn.
// Exercise care, as this is a potentially dangerous function if the Object is not a TreeViewColumn.
func CastToTreeViewColumn(object *gobject.Object) *TreeViewColumn {
	return TreeViewColumnNewFromC(object.ToC())
}

var signalClickedId int
var signalClickedMap = make(map[int]TreeViewColumnSignalClickedCallback)
var signalClickedLock sync.Mutex

// TreeViewColumnSignalClickedCallback is a callback function for a 'clicked' signal emitted from a TreeViewColumn.
type TreeViewColumnSignalClickedCallback func()

func TreeViewColumn_clickedHandler() {}

// TreeViewColumnNew is a wrapper around the C function gtk_tree_view_column_new.
func TreeViewColumnNew() *TreeViewColumn {
	retC := C.gtk_tree_view_column_new()
	retGo := TreeViewColumnNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_tree_view_column_new_with_attributes : unsupported parameter ... : varargs

// AddAttribute is a wrapper around the C function gtk_tree_view_column_add_attribute.
func (recv *TreeViewColumn) AddAttribute(cellRenderer *CellRenderer, attribute string, column int32) {
	c_cell_renderer := (*C.GtkCellRenderer)(cellRenderer.ToC())

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_tree_view_column_add_attribute((*C.GtkTreeViewColumn)(recv.native), c_cell_renderer, c_attribute, c_column)

	return
}

// CellGetPosition is a wrapper around the C function gtk_tree_view_column_cell_get_position.
func (recv *TreeViewColumn) CellGetPosition(cellRenderer *CellRenderer) (bool, int32, int32) {
	c_cell_renderer := (*C.GtkCellRenderer)(cellRenderer.ToC())

	var c_x_offset C.gint

	var c_width C.gint

	retC := C.gtk_tree_view_column_cell_get_position((*C.GtkTreeViewColumn)(recv.native), c_cell_renderer, &c_x_offset, &c_width)
	retGo := retC == C.TRUE

	xOffset := (int32)(c_x_offset)

	width := (int32)(c_width)

	return retGo, xOffset, width
}

// Unsupported : gtk_tree_view_column_cell_get_size : unsupported parameter cell_area : Blacklisted record : GdkRectangle

// CellIsVisible is a wrapper around the C function gtk_tree_view_column_cell_is_visible.
func (recv *TreeViewColumn) CellIsVisible() bool {
	retC := C.gtk_tree_view_column_cell_is_visible((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_tree_view_column_cell_set_cell_data : unsupported parameter tree_model : no type generator for TreeModel, GtkTreeModel*

// Clear is a wrapper around the C function gtk_tree_view_column_clear.
func (recv *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear((*C.GtkTreeViewColumn)(recv.native))

	return
}

// ClearAttributes is a wrapper around the C function gtk_tree_view_column_clear_attributes.
func (recv *TreeViewColumn) ClearAttributes(cellRenderer *CellRenderer) {
	c_cell_renderer := (*C.GtkCellRenderer)(cellRenderer.ToC())

	C.gtk_tree_view_column_clear_attributes((*C.GtkTreeViewColumn)(recv.native), c_cell_renderer)

	return
}

// Clicked is a wrapper around the C function gtk_tree_view_column_clicked.
func (recv *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked((*C.GtkTreeViewColumn)(recv.native))

	return
}

// GetAlignment is a wrapper around the C function gtk_tree_view_column_get_alignment.
func (recv *TreeViewColumn) GetAlignment() float32 {
	retC := C.gtk_tree_view_column_get_alignment((*C.GtkTreeViewColumn)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetClickable is a wrapper around the C function gtk_tree_view_column_get_clickable.
func (recv *TreeViewColumn) GetClickable() bool {
	retC := C.gtk_tree_view_column_get_clickable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFixedWidth is a wrapper around the C function gtk_tree_view_column_get_fixed_width.
func (recv *TreeViewColumn) GetFixedWidth() int32 {
	retC := C.gtk_tree_view_column_get_fixed_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMaxWidth is a wrapper around the C function gtk_tree_view_column_get_max_width.
func (recv *TreeViewColumn) GetMaxWidth() int32 {
	retC := C.gtk_tree_view_column_get_max_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMinWidth is a wrapper around the C function gtk_tree_view_column_get_min_width.
func (recv *TreeViewColumn) GetMinWidth() int32 {
	retC := C.gtk_tree_view_column_get_min_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetReorderable is a wrapper around the C function gtk_tree_view_column_get_reorderable.
func (recv *TreeViewColumn) GetReorderable() bool {
	retC := C.gtk_tree_view_column_get_reorderable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetResizable is a wrapper around the C function gtk_tree_view_column_get_resizable.
func (recv *TreeViewColumn) GetResizable() bool {
	retC := C.gtk_tree_view_column_get_resizable((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSizing is a wrapper around the C function gtk_tree_view_column_get_sizing.
func (recv *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	retC := C.gtk_tree_view_column_get_sizing((*C.GtkTreeViewColumn)(recv.native))
	retGo := (TreeViewColumnSizing)(retC)

	return retGo
}

// GetSortColumnId is a wrapper around the C function gtk_tree_view_column_get_sort_column_id.
func (recv *TreeViewColumn) GetSortColumnId() int32 {
	retC := C.gtk_tree_view_column_get_sort_column_id((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSortIndicator is a wrapper around the C function gtk_tree_view_column_get_sort_indicator.
func (recv *TreeViewColumn) GetSortIndicator() bool {
	retC := C.gtk_tree_view_column_get_sort_indicator((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSortOrder is a wrapper around the C function gtk_tree_view_column_get_sort_order.
func (recv *TreeViewColumn) GetSortOrder() SortType {
	retC := C.gtk_tree_view_column_get_sort_order((*C.GtkTreeViewColumn)(recv.native))
	retGo := (SortType)(retC)

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_tree_view_column_get_spacing.
func (recv *TreeViewColumn) GetSpacing() int32 {
	retC := C.gtk_tree_view_column_get_spacing((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetTitle is a wrapper around the C function gtk_tree_view_column_get_title.
func (recv *TreeViewColumn) GetTitle() string {
	retC := C.gtk_tree_view_column_get_title((*C.GtkTreeViewColumn)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetVisible is a wrapper around the C function gtk_tree_view_column_get_visible.
func (recv *TreeViewColumn) GetVisible() bool {
	retC := C.gtk_tree_view_column_get_visible((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetWidget is a wrapper around the C function gtk_tree_view_column_get_widget.
func (recv *TreeViewColumn) GetWidget() *Widget {
	retC := C.gtk_tree_view_column_get_widget((*C.GtkTreeViewColumn)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetWidth is a wrapper around the C function gtk_tree_view_column_get_width.
func (recv *TreeViewColumn) GetWidth() int32 {
	retC := C.gtk_tree_view_column_get_width((*C.GtkTreeViewColumn)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// PackEnd is a wrapper around the C function gtk_tree_view_column_pack_end.
func (recv *TreeViewColumn) PackEnd(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_column_pack_end((*C.GtkTreeViewColumn)(recv.native), c_cell, c_expand)

	return
}

// PackStart is a wrapper around the C function gtk_tree_view_column_pack_start.
func (recv *TreeViewColumn) PackStart(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(cell.ToC())

	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_column_pack_start((*C.GtkTreeViewColumn)(recv.native), c_cell, c_expand)

	return
}

// SetAlignment is a wrapper around the C function gtk_tree_view_column_set_alignment.
func (recv *TreeViewColumn) SetAlignment(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_tree_view_column_set_alignment((*C.GtkTreeViewColumn)(recv.native), c_xalign)

	return
}

// Unsupported : gtk_tree_view_column_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_tree_view_column_set_cell_data_func : unsupported parameter func : no type generator for TreeCellDataFunc, GtkTreeCellDataFunc

// SetClickable is a wrapper around the C function gtk_tree_view_column_set_clickable.
func (recv *TreeViewColumn) SetClickable(clickable bool) {
	c_clickable :=
		boolToGboolean(clickable)

	C.gtk_tree_view_column_set_clickable((*C.GtkTreeViewColumn)(recv.native), c_clickable)

	return
}

// SetFixedWidth is a wrapper around the C function gtk_tree_view_column_set_fixed_width.
func (recv *TreeViewColumn) SetFixedWidth(fixedWidth int32) {
	c_fixed_width := (C.gint)(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width((*C.GtkTreeViewColumn)(recv.native), c_fixed_width)

	return
}

// SetMaxWidth is a wrapper around the C function gtk_tree_view_column_set_max_width.
func (recv *TreeViewColumn) SetMaxWidth(maxWidth int32) {
	c_max_width := (C.gint)(maxWidth)

	C.gtk_tree_view_column_set_max_width((*C.GtkTreeViewColumn)(recv.native), c_max_width)

	return
}

// SetMinWidth is a wrapper around the C function gtk_tree_view_column_set_min_width.
func (recv *TreeViewColumn) SetMinWidth(minWidth int32) {
	c_min_width := (C.gint)(minWidth)

	C.gtk_tree_view_column_set_min_width((*C.GtkTreeViewColumn)(recv.native), c_min_width)

	return
}

// SetReorderable is a wrapper around the C function gtk_tree_view_column_set_reorderable.
func (recv *TreeViewColumn) SetReorderable(reorderable bool) {
	c_reorderable :=
		boolToGboolean(reorderable)

	C.gtk_tree_view_column_set_reorderable((*C.GtkTreeViewColumn)(recv.native), c_reorderable)

	return
}

// SetResizable is a wrapper around the C function gtk_tree_view_column_set_resizable.
func (recv *TreeViewColumn) SetResizable(resizable bool) {
	c_resizable :=
		boolToGboolean(resizable)

	C.gtk_tree_view_column_set_resizable((*C.GtkTreeViewColumn)(recv.native), c_resizable)

	return
}

// SetSizing is a wrapper around the C function gtk_tree_view_column_set_sizing.
func (recv *TreeViewColumn) SetSizing(type_ TreeViewColumnSizing) {
	c_type := (C.GtkTreeViewColumnSizing)(type_)

	C.gtk_tree_view_column_set_sizing((*C.GtkTreeViewColumn)(recv.native), c_type)

	return
}

// SetSortColumnId is a wrapper around the C function gtk_tree_view_column_set_sort_column_id.
func (recv *TreeViewColumn) SetSortColumnId(sortColumnId int32) {
	c_sort_column_id := (C.gint)(sortColumnId)

	C.gtk_tree_view_column_set_sort_column_id((*C.GtkTreeViewColumn)(recv.native), c_sort_column_id)

	return
}

// SetSortIndicator is a wrapper around the C function gtk_tree_view_column_set_sort_indicator.
func (recv *TreeViewColumn) SetSortIndicator(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_tree_view_column_set_sort_indicator((*C.GtkTreeViewColumn)(recv.native), c_setting)

	return
}

// SetSortOrder is a wrapper around the C function gtk_tree_view_column_set_sort_order.
func (recv *TreeViewColumn) SetSortOrder(order SortType) {
	c_order := (C.GtkSortType)(order)

	C.gtk_tree_view_column_set_sort_order((*C.GtkTreeViewColumn)(recv.native), c_order)

	return
}

// SetSpacing is a wrapper around the C function gtk_tree_view_column_set_spacing.
func (recv *TreeViewColumn) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_tree_view_column_set_spacing((*C.GtkTreeViewColumn)(recv.native), c_spacing)

	return
}

// SetTitle is a wrapper around the C function gtk_tree_view_column_set_title.
func (recv *TreeViewColumn) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_tree_view_column_set_title((*C.GtkTreeViewColumn)(recv.native), c_title)

	return
}

// SetVisible is a wrapper around the C function gtk_tree_view_column_set_visible.
func (recv *TreeViewColumn) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_tree_view_column_set_visible((*C.GtkTreeViewColumn)(recv.native), c_visible)

	return
}

// SetWidget is a wrapper around the C function gtk_tree_view_column_set_widget.
func (recv *TreeViewColumn) SetWidget(widget *Widget) {
	c_widget := (*C.GtkWidget)(widget.ToC())

	C.gtk_tree_view_column_set_widget((*C.GtkTreeViewColumn)(recv.native), c_widget)

	return
}

// UIManager is a wrapper around the C record GtkUIManager.
type UIManager struct {
	native *C.GtkUIManager
	// parent : record
	// Private : private_data
}

func UIManagerNewFromC(u unsafe.Pointer) *UIManager {
	c := (*C.GtkUIManager)(u)
	if c == nil {
		return nil
	}

	g := &UIManager{native: c}

	return g
}

func (recv *UIManager) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *UIManager) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to UIManager.
// Exercise care, as this is a potentially dangerous function if the Object is not a UIManager.
func CastToUIManager(object *gobject.Object) *UIManager {
	return UIManagerNewFromC(object.ToC())
}

// VBox is a wrapper around the C record GtkVBox.
type VBox struct {
	native *C.GtkVBox
	// box : record
}

func VBoxNewFromC(u unsafe.Pointer) *VBox {
	c := (*C.GtkVBox)(u)
	if c == nil {
		return nil
	}

	g := &VBox{native: c}

	return g
}

func (recv *VBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Box upcasts to *Box
func (recv *VBox) Box() *Box {
	return BoxNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *VBox) Container() *Container {
	return recv.Box().Container()
}

// Widget upcasts to *Widget
func (recv *VBox) Widget() *Widget {
	return recv.Box().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Box().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VBox) Object() *gobject.Object {
	return recv.Box().Object()
}

// CastToWidget down casts any arbitary Object to VBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a VBox.
func CastToVBox(object *gobject.Object) *VBox {
	return VBoxNewFromC(object.ToC())
}

// VBoxNew is a wrapper around the C function gtk_vbox_new.
func VBoxNew(homogeneous bool, spacing int32) *VBox {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	c_spacing := (C.gint)(spacing)

	retC := C.gtk_vbox_new(c_homogeneous, c_spacing)
	retGo := VBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VButtonBox is a wrapper around the C record GtkVButtonBox.
type VButtonBox struct {
	native *C.GtkVButtonBox
	// button_box : record
}

func VButtonBoxNewFromC(u unsafe.Pointer) *VButtonBox {
	c := (*C.GtkVButtonBox)(u)
	if c == nil {
		return nil
	}

	g := &VButtonBox{native: c}

	return g
}

func (recv *VButtonBox) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ButtonBox upcasts to *ButtonBox
func (recv *VButtonBox) ButtonBox() *ButtonBox {
	return ButtonBoxNewFromC(unsafe.Pointer(recv.native))
}

// Box upcasts to *Box
func (recv *VButtonBox) Box() *Box {
	return recv.ButtonBox().Box()
}

// Container upcasts to *Container
func (recv *VButtonBox) Container() *Container {
	return recv.ButtonBox().Container()
}

// Widget upcasts to *Widget
func (recv *VButtonBox) Widget() *Widget {
	return recv.ButtonBox().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VButtonBox) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ButtonBox().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VButtonBox) Object() *gobject.Object {
	return recv.ButtonBox().Object()
}

// CastToWidget down casts any arbitary Object to VButtonBox.
// Exercise care, as this is a potentially dangerous function if the Object is not a VButtonBox.
func CastToVButtonBox(object *gobject.Object) *VButtonBox {
	return VButtonBoxNewFromC(object.ToC())
}

// VButtonBoxNew is a wrapper around the C function gtk_vbutton_box_new.
func VButtonBoxNew() *VButtonBox {
	retC := C.gtk_vbutton_box_new()
	retGo := VButtonBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VPaned is a wrapper around the C record GtkVPaned.
type VPaned struct {
	native *C.GtkVPaned
	// paned : record
}

func VPanedNewFromC(u unsafe.Pointer) *VPaned {
	c := (*C.GtkVPaned)(u)
	if c == nil {
		return nil
	}

	g := &VPaned{native: c}

	return g
}

func (recv *VPaned) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Paned upcasts to *Paned
func (recv *VPaned) Paned() *Paned {
	return PanedNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *VPaned) Container() *Container {
	return recv.Paned().Container()
}

// Widget upcasts to *Widget
func (recv *VPaned) Widget() *Widget {
	return recv.Paned().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VPaned) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Paned().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VPaned) Object() *gobject.Object {
	return recv.Paned().Object()
}

// CastToWidget down casts any arbitary Object to VPaned.
// Exercise care, as this is a potentially dangerous function if the Object is not a VPaned.
func CastToVPaned(object *gobject.Object) *VPaned {
	return VPanedNewFromC(object.ToC())
}

// VPanedNew is a wrapper around the C function gtk_vpaned_new.
func VPanedNew() *VPaned {
	retC := C.gtk_vpaned_new()
	retGo := VPanedNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScale is a wrapper around the C record GtkVScale.
type VScale struct {
	native *C.GtkVScale
	// scale : record
}

func VScaleNewFromC(u unsafe.Pointer) *VScale {
	c := (*C.GtkVScale)(u)
	if c == nil {
		return nil
	}

	g := &VScale{native: c}

	return g
}

func (recv *VScale) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scale upcasts to *Scale
func (recv *VScale) Scale() *Scale {
	return ScaleNewFromC(unsafe.Pointer(recv.native))
}

// Range upcasts to *Range
func (recv *VScale) Range() *Range {
	return recv.Scale().Range()
}

// Widget upcasts to *Widget
func (recv *VScale) Widget() *Widget {
	return recv.Scale().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VScale) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Scale().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VScale) Object() *gobject.Object {
	return recv.Scale().Object()
}

// CastToWidget down casts any arbitary Object to VScale.
// Exercise care, as this is a potentially dangerous function if the Object is not a VScale.
func CastToVScale(object *gobject.Object) *VScale {
	return VScaleNewFromC(object.ToC())
}

// VScaleNew is a wrapper around the C function gtk_vscale_new.
func VScaleNew(adjustment *Adjustment) *VScale {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_vscale_new(c_adjustment)
	retGo := VScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScaleNewWithRange is a wrapper around the C function gtk_vscale_new_with_range.
func VScaleNewWithRange(min float64, max float64, step float64) *VScale {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	retC := C.gtk_vscale_new_with_range(c_min, c_max, c_step)
	retGo := VScaleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VScrollbar is a wrapper around the C record GtkVScrollbar.
type VScrollbar struct {
	native *C.GtkVScrollbar
	// scrollbar : record
}

func VScrollbarNewFromC(u unsafe.Pointer) *VScrollbar {
	c := (*C.GtkVScrollbar)(u)
	if c == nil {
		return nil
	}

	g := &VScrollbar{native: c}

	return g
}

func (recv *VScrollbar) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Scrollbar upcasts to *Scrollbar
func (recv *VScrollbar) Scrollbar() *Scrollbar {
	return ScrollbarNewFromC(unsafe.Pointer(recv.native))
}

// Range upcasts to *Range
func (recv *VScrollbar) Range() *Range {
	return recv.Scrollbar().Range()
}

// Widget upcasts to *Widget
func (recv *VScrollbar) Widget() *Widget {
	return recv.Scrollbar().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VScrollbar) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Scrollbar().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VScrollbar) Object() *gobject.Object {
	return recv.Scrollbar().Object()
}

// CastToWidget down casts any arbitary Object to VScrollbar.
// Exercise care, as this is a potentially dangerous function if the Object is not a VScrollbar.
func CastToVScrollbar(object *gobject.Object) *VScrollbar {
	return VScrollbarNewFromC(object.ToC())
}

// VScrollbarNew is a wrapper around the C function gtk_vscrollbar_new.
func VScrollbarNew(adjustment *Adjustment) *VScrollbar {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	retC := C.gtk_vscrollbar_new(c_adjustment)
	retGo := VScrollbarNewFromC(unsafe.Pointer(retC))

	return retGo
}

// VSeparator is a wrapper around the C record GtkVSeparator.
type VSeparator struct {
	native *C.GtkVSeparator
	// separator : record
}

func VSeparatorNewFromC(u unsafe.Pointer) *VSeparator {
	c := (*C.GtkVSeparator)(u)
	if c == nil {
		return nil
	}

	g := &VSeparator{native: c}

	return g
}

func (recv *VSeparator) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Separator upcasts to *Separator
func (recv *VSeparator) Separator() *Separator {
	return SeparatorNewFromC(unsafe.Pointer(recv.native))
}

// Widget upcasts to *Widget
func (recv *VSeparator) Widget() *Widget {
	return recv.Separator().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VSeparator) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Separator().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VSeparator) Object() *gobject.Object {
	return recv.Separator().Object()
}

// CastToWidget down casts any arbitary Object to VSeparator.
// Exercise care, as this is a potentially dangerous function if the Object is not a VSeparator.
func CastToVSeparator(object *gobject.Object) *VSeparator {
	return VSeparatorNewFromC(object.ToC())
}

// VSeparatorNew is a wrapper around the C function gtk_vseparator_new.
func VSeparatorNew() *VSeparator {
	retC := C.gtk_vseparator_new()
	retGo := VSeparatorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Viewport is a wrapper around the C record GtkViewport.
type Viewport struct {
	native *C.GtkViewport
	// bin : record
	// Private : priv
}

func ViewportNewFromC(u unsafe.Pointer) *Viewport {
	c := (*C.GtkViewport)(u)
	if c == nil {
		return nil
	}

	g := &Viewport{native: c}

	return g
}

func (recv *Viewport) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Viewport) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Viewport) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Viewport) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Viewport) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Viewport) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Viewport.
// Exercise care, as this is a potentially dangerous function if the Object is not a Viewport.
func CastToViewport(object *gobject.Object) *Viewport {
	return ViewportNewFromC(object.ToC())
}

// ViewportNew is a wrapper around the C function gtk_viewport_new.
func ViewportNew(hadjustment *Adjustment, vadjustment *Adjustment) *Viewport {
	c_hadjustment := (*C.GtkAdjustment)(hadjustment.ToC())

	c_vadjustment := (*C.GtkAdjustment)(vadjustment.ToC())

	retC := C.gtk_viewport_new(c_hadjustment, c_vadjustment)
	retGo := ViewportNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetHadjustment is a wrapper around the C function gtk_viewport_get_hadjustment.
func (recv *Viewport) GetHadjustment() *Adjustment {
	retC := C.gtk_viewport_get_hadjustment((*C.GtkViewport)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetShadowType is a wrapper around the C function gtk_viewport_get_shadow_type.
func (recv *Viewport) GetShadowType() ShadowType {
	retC := C.gtk_viewport_get_shadow_type((*C.GtkViewport)(recv.native))
	retGo := (ShadowType)(retC)

	return retGo
}

// GetVadjustment is a wrapper around the C function gtk_viewport_get_vadjustment.
func (recv *Viewport) GetVadjustment() *Adjustment {
	retC := C.gtk_viewport_get_vadjustment((*C.GtkViewport)(recv.native))
	retGo := AdjustmentNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetHadjustment is a wrapper around the C function gtk_viewport_set_hadjustment.
func (recv *Viewport) SetHadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_viewport_set_hadjustment((*C.GtkViewport)(recv.native), c_adjustment)

	return
}

// SetShadowType is a wrapper around the C function gtk_viewport_set_shadow_type.
func (recv *Viewport) SetShadowType(type_ ShadowType) {
	c_type := (C.GtkShadowType)(type_)

	C.gtk_viewport_set_shadow_type((*C.GtkViewport)(recv.native), c_type)

	return
}

// SetVadjustment is a wrapper around the C function gtk_viewport_set_vadjustment.
func (recv *Viewport) SetVadjustment(adjustment *Adjustment) {
	c_adjustment := (*C.GtkAdjustment)(adjustment.ToC())

	C.gtk_viewport_set_vadjustment((*C.GtkViewport)(recv.native), c_adjustment)

	return
}

// VolumeButton is a wrapper around the C record GtkVolumeButton.
type VolumeButton struct {
	native *C.GtkVolumeButton
	// parent : record
}

func VolumeButtonNewFromC(u unsafe.Pointer) *VolumeButton {
	c := (*C.GtkVolumeButton)(u)
	if c == nil {
		return nil
	}

	g := &VolumeButton{native: c}

	return g
}

func (recv *VolumeButton) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ScaleButton upcasts to *ScaleButton
func (recv *VolumeButton) ScaleButton() *ScaleButton {
	return ScaleButtonNewFromC(unsafe.Pointer(recv.native))
}

// Button upcasts to *Button
func (recv *VolumeButton) Button() *Button {
	return recv.ScaleButton().Button()
}

// Bin upcasts to *Bin
func (recv *VolumeButton) Bin() *Bin {
	return recv.ScaleButton().Bin()
}

// Container upcasts to *Container
func (recv *VolumeButton) Container() *Container {
	return recv.ScaleButton().Container()
}

// Widget upcasts to *Widget
func (recv *VolumeButton) Widget() *Widget {
	return recv.ScaleButton().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *VolumeButton) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.ScaleButton().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *VolumeButton) Object() *gobject.Object {
	return recv.ScaleButton().Object()
}

// CastToWidget down casts any arbitary Object to VolumeButton.
// Exercise care, as this is a potentially dangerous function if the Object is not a VolumeButton.
func CastToVolumeButton(object *gobject.Object) *VolumeButton {
	return VolumeButtonNewFromC(object.ToC())
}

// Widget is a wrapper around the C record GtkWidget.
type Widget struct {
	native *C.GtkWidget
	// parent_instance : record
	// Private : priv
}

func WidgetNewFromC(u unsafe.Pointer) *Widget {
	c := (*C.GtkWidget)(u)
	if c == nil {
		return nil
	}

	g := &Widget{native: c}

	return g
}

func (recv *Widget) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Widget) InitiallyUnowned() *gobject.InitiallyUnowned {
	return gobject.InitiallyUnownedNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *Widget) Object() *gobject.Object {
	return recv.InitiallyUnowned().Object()
}

// CastToWidget down casts any arbitary Object to Widget.
// Exercise care, as this is a potentially dangerous function if the Object is not a Widget.
func CastToWidget(object *gobject.Object) *Widget {
	return WidgetNewFromC(object.ToC())
}

var signalAccelClosuresChangedId int
var signalAccelClosuresChangedMap = make(map[int]WidgetSignalAccelClosuresChangedCallback)
var signalAccelClosuresChangedLock sync.Mutex

// WidgetSignalAccelClosuresChangedCallback is a callback function for a 'accel-closures-changed' signal emitted from a Widget.
type WidgetSignalAccelClosuresChangedCallback func()

func Widget_accelClosuresChangedHandler() {}

var signalButtonPressEventId int
var signalButtonPressEventMap = make(map[int]WidgetSignalButtonPressEventCallback)
var signalButtonPressEventLock sync.Mutex

// WidgetSignalButtonPressEventCallback is a callback function for a 'button-press-event' signal emitted from a Widget.
type WidgetSignalButtonPressEventCallback func(event *gdk.EventButton) bool

func Widget_buttonPressEventHandler() {}

var signalButtonReleaseEventId int
var signalButtonReleaseEventMap = make(map[int]WidgetSignalButtonReleaseEventCallback)
var signalButtonReleaseEventLock sync.Mutex

// WidgetSignalButtonReleaseEventCallback is a callback function for a 'button-release-event' signal emitted from a Widget.
type WidgetSignalButtonReleaseEventCallback func(event *gdk.EventButton) bool

func Widget_buttonReleaseEventHandler() {}

var signalCanActivateAccelId int
var signalCanActivateAccelMap = make(map[int]WidgetSignalCanActivateAccelCallback)
var signalCanActivateAccelLock sync.Mutex

// WidgetSignalCanActivateAccelCallback is a callback function for a 'can-activate-accel' signal emitted from a Widget.
type WidgetSignalCanActivateAccelCallback func(signalId uint32) bool

func Widget_canActivateAccelHandler() {}

// Unsupported signal : unsupported parameter child_property : Blacklisted record : GParamSpec

var signalCompositedChangedId int
var signalCompositedChangedMap = make(map[int]WidgetSignalCompositedChangedCallback)
var signalCompositedChangedLock sync.Mutex

// WidgetSignalCompositedChangedCallback is a callback function for a 'composited-changed' signal emitted from a Widget.
type WidgetSignalCompositedChangedCallback func()

func Widget_compositedChangedHandler() {}

var signalConfigureEventId int
var signalConfigureEventMap = make(map[int]WidgetSignalConfigureEventCallback)
var signalConfigureEventLock sync.Mutex

// WidgetSignalConfigureEventCallback is a callback function for a 'configure-event' signal emitted from a Widget.
type WidgetSignalConfigureEventCallback func(event *gdk.EventConfigure) bool

func Widget_configureEventHandler() {}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

var signalDestroyId int
var signalDestroyMap = make(map[int]WidgetSignalDestroyCallback)
var signalDestroyLock sync.Mutex

// WidgetSignalDestroyCallback is a callback function for a 'destroy' signal emitted from a Widget.
type WidgetSignalDestroyCallback func()

func Widget_destroyHandler() {}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

var signalDirectionChangedId int
var signalDirectionChangedMap = make(map[int]WidgetSignalDirectionChangedCallback)
var signalDirectionChangedLock sync.Mutex

// WidgetSignalDirectionChangedCallback is a callback function for a 'direction-changed' signal emitted from a Widget.
type WidgetSignalDirectionChangedCallback func(previousDirection TextDirection)

func Widget_directionChangedHandler() {}

var signalDragBeginId int
var signalDragBeginMap = make(map[int]WidgetSignalDragBeginCallback)
var signalDragBeginLock sync.Mutex

// WidgetSignalDragBeginCallback is a callback function for a 'drag-begin' signal emitted from a Widget.
type WidgetSignalDragBeginCallback func(context *gdk.DragContext)

func Widget_dragBeginHandler() {}

var signalDragDataDeleteId int
var signalDragDataDeleteMap = make(map[int]WidgetSignalDragDataDeleteCallback)
var signalDragDataDeleteLock sync.Mutex

// WidgetSignalDragDataDeleteCallback is a callback function for a 'drag-data-delete' signal emitted from a Widget.
type WidgetSignalDragDataDeleteCallback func(context *gdk.DragContext)

func Widget_dragDataDeleteHandler() {}

var signalDragDataGetId int
var signalDragDataGetMap = make(map[int]WidgetSignalDragDataGetCallback)
var signalDragDataGetLock sync.Mutex

// WidgetSignalDragDataGetCallback is a callback function for a 'drag-data-get' signal emitted from a Widget.
type WidgetSignalDragDataGetCallback func(context *gdk.DragContext, data *SelectionData, info uint32, time uint32)

func Widget_dragDataGetHandler() {}

var signalDragDataReceivedId int
var signalDragDataReceivedMap = make(map[int]WidgetSignalDragDataReceivedCallback)
var signalDragDataReceivedLock sync.Mutex

// WidgetSignalDragDataReceivedCallback is a callback function for a 'drag-data-received' signal emitted from a Widget.
type WidgetSignalDragDataReceivedCallback func(context *gdk.DragContext, x int32, y int32, data *SelectionData, info uint32, time uint32)

func Widget_dragDataReceivedHandler() {}

var signalDragDropId int
var signalDragDropMap = make(map[int]WidgetSignalDragDropCallback)
var signalDragDropLock sync.Mutex

// WidgetSignalDragDropCallback is a callback function for a 'drag-drop' signal emitted from a Widget.
type WidgetSignalDragDropCallback func(context *gdk.DragContext, x int32, y int32, time uint32) bool

func Widget_dragDropHandler() {}

var signalDragEndId int
var signalDragEndMap = make(map[int]WidgetSignalDragEndCallback)
var signalDragEndLock sync.Mutex

// WidgetSignalDragEndCallback is a callback function for a 'drag-end' signal emitted from a Widget.
type WidgetSignalDragEndCallback func(context *gdk.DragContext)

func Widget_dragEndHandler() {}

var signalDragLeaveId int
var signalDragLeaveMap = make(map[int]WidgetSignalDragLeaveCallback)
var signalDragLeaveLock sync.Mutex

// WidgetSignalDragLeaveCallback is a callback function for a 'drag-leave' signal emitted from a Widget.
type WidgetSignalDragLeaveCallback func(context *gdk.DragContext, time uint32)

func Widget_dragLeaveHandler() {}

var signalDragMotionId int
var signalDragMotionMap = make(map[int]WidgetSignalDragMotionCallback)
var signalDragMotionLock sync.Mutex

// WidgetSignalDragMotionCallback is a callback function for a 'drag-motion' signal emitted from a Widget.
type WidgetSignalDragMotionCallback func(context *gdk.DragContext, x int32, y int32, time uint32) bool

func Widget_dragMotionHandler() {}

var signalEnterNotifyEventId int
var signalEnterNotifyEventMap = make(map[int]WidgetSignalEnterNotifyEventCallback)
var signalEnterNotifyEventLock sync.Mutex

// WidgetSignalEnterNotifyEventCallback is a callback function for a 'enter-notify-event' signal emitted from a Widget.
type WidgetSignalEnterNotifyEventCallback func(event *gdk.EventCrossing) bool

func Widget_enterNotifyEventHandler() {}

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

// Unsupported signal : unsupported parameter event : no type generator for Gdk.Event,

var signalFocusId int
var signalFocusMap = make(map[int]WidgetSignalFocusCallback)
var signalFocusLock sync.Mutex

// WidgetSignalFocusCallback is a callback function for a 'focus' signal emitted from a Widget.
type WidgetSignalFocusCallback func(direction DirectionType) bool

func Widget_focusHandler() {}

var signalFocusInEventId int
var signalFocusInEventMap = make(map[int]WidgetSignalFocusInEventCallback)
var signalFocusInEventLock sync.Mutex

// WidgetSignalFocusInEventCallback is a callback function for a 'focus-in-event' signal emitted from a Widget.
type WidgetSignalFocusInEventCallback func(event *gdk.EventFocus) bool

func Widget_focusInEventHandler() {}

var signalFocusOutEventId int
var signalFocusOutEventMap = make(map[int]WidgetSignalFocusOutEventCallback)
var signalFocusOutEventLock sync.Mutex

// WidgetSignalFocusOutEventCallback is a callback function for a 'focus-out-event' signal emitted from a Widget.
type WidgetSignalFocusOutEventCallback func(event *gdk.EventFocus) bool

func Widget_focusOutEventHandler() {}

var signalGrabFocusId int
var signalGrabFocusMap = make(map[int]WidgetSignalGrabFocusCallback)
var signalGrabFocusLock sync.Mutex

// WidgetSignalGrabFocusCallback is a callback function for a 'grab-focus' signal emitted from a Widget.
type WidgetSignalGrabFocusCallback func()

func Widget_grabFocusHandler() {}

var signalGrabNotifyId int
var signalGrabNotifyMap = make(map[int]WidgetSignalGrabNotifyCallback)
var signalGrabNotifyLock sync.Mutex

// WidgetSignalGrabNotifyCallback is a callback function for a 'grab-notify' signal emitted from a Widget.
type WidgetSignalGrabNotifyCallback func(wasGrabbed bool)

func Widget_grabNotifyHandler() {}

var signalHideId int
var signalHideMap = make(map[int]WidgetSignalHideCallback)
var signalHideLock sync.Mutex

// WidgetSignalHideCallback is a callback function for a 'hide' signal emitted from a Widget.
type WidgetSignalHideCallback func()

func Widget_hideHandler() {}

var signalHierarchyChangedId int
var signalHierarchyChangedMap = make(map[int]WidgetSignalHierarchyChangedCallback)
var signalHierarchyChangedLock sync.Mutex

// WidgetSignalHierarchyChangedCallback is a callback function for a 'hierarchy-changed' signal emitted from a Widget.
type WidgetSignalHierarchyChangedCallback func(previousToplevel *Widget)

func Widget_hierarchyChangedHandler() {}

var signalKeyPressEventId int
var signalKeyPressEventMap = make(map[int]WidgetSignalKeyPressEventCallback)
var signalKeyPressEventLock sync.Mutex

// WidgetSignalKeyPressEventCallback is a callback function for a 'key-press-event' signal emitted from a Widget.
type WidgetSignalKeyPressEventCallback func(event *gdk.EventKey) bool

func Widget_keyPressEventHandler() {}

var signalKeyReleaseEventId int
var signalKeyReleaseEventMap = make(map[int]WidgetSignalKeyReleaseEventCallback)
var signalKeyReleaseEventLock sync.Mutex

// WidgetSignalKeyReleaseEventCallback is a callback function for a 'key-release-event' signal emitted from a Widget.
type WidgetSignalKeyReleaseEventCallback func(event *gdk.EventKey) bool

func Widget_keyReleaseEventHandler() {}

var signalLeaveNotifyEventId int
var signalLeaveNotifyEventMap = make(map[int]WidgetSignalLeaveNotifyEventCallback)
var signalLeaveNotifyEventLock sync.Mutex

// WidgetSignalLeaveNotifyEventCallback is a callback function for a 'leave-notify-event' signal emitted from a Widget.
type WidgetSignalLeaveNotifyEventCallback func(event *gdk.EventCrossing) bool

func Widget_leaveNotifyEventHandler() {}

var signalMapId int
var signalMapMap = make(map[int]WidgetSignalMapCallback)
var signalMapLock sync.Mutex

// WidgetSignalMapCallback is a callback function for a 'map' signal emitted from a Widget.
type WidgetSignalMapCallback func()

func Widget_mapHandler() {}

var signalMapEventId int
var signalMapEventMap = make(map[int]WidgetSignalMapEventCallback)
var signalMapEventLock sync.Mutex

// WidgetSignalMapEventCallback is a callback function for a 'map-event' signal emitted from a Widget.
type WidgetSignalMapEventCallback func(event *gdk.EventAny) bool

func Widget_mapEventHandler() {}

var signalMnemonicActivateId int
var signalMnemonicActivateMap = make(map[int]WidgetSignalMnemonicActivateCallback)
var signalMnemonicActivateLock sync.Mutex

// WidgetSignalMnemonicActivateCallback is a callback function for a 'mnemonic-activate' signal emitted from a Widget.
type WidgetSignalMnemonicActivateCallback func(groupCycling bool) bool

func Widget_mnemonicActivateHandler() {}

var signalMotionNotifyEventId int
var signalMotionNotifyEventMap = make(map[int]WidgetSignalMotionNotifyEventCallback)
var signalMotionNotifyEventLock sync.Mutex

// WidgetSignalMotionNotifyEventCallback is a callback function for a 'motion-notify-event' signal emitted from a Widget.
type WidgetSignalMotionNotifyEventCallback func(event *gdk.EventMotion) bool

func Widget_motionNotifyEventHandler() {}

var signalMoveFocusId int
var signalMoveFocusMap = make(map[int]WidgetSignalMoveFocusCallback)
var signalMoveFocusLock sync.Mutex

// WidgetSignalMoveFocusCallback is a callback function for a 'move-focus' signal emitted from a Widget.
type WidgetSignalMoveFocusCallback func(direction DirectionType)

func Widget_moveFocusHandler() {}

var signalParentSetId int
var signalParentSetMap = make(map[int]WidgetSignalParentSetCallback)
var signalParentSetLock sync.Mutex

// WidgetSignalParentSetCallback is a callback function for a 'parent-set' signal emitted from a Widget.
type WidgetSignalParentSetCallback func(oldParent *Widget)

func Widget_parentSetHandler() {}

var signalPopupMenuId int
var signalPopupMenuMap = make(map[int]WidgetSignalPopupMenuCallback)
var signalPopupMenuLock sync.Mutex

// WidgetSignalPopupMenuCallback is a callback function for a 'popup-menu' signal emitted from a Widget.
type WidgetSignalPopupMenuCallback func() bool

func Widget_popupMenuHandler() {}

var signalPropertyNotifyEventId int
var signalPropertyNotifyEventMap = make(map[int]WidgetSignalPropertyNotifyEventCallback)
var signalPropertyNotifyEventLock sync.Mutex

// WidgetSignalPropertyNotifyEventCallback is a callback function for a 'property-notify-event' signal emitted from a Widget.
type WidgetSignalPropertyNotifyEventCallback func(event *gdk.EventProperty) bool

func Widget_propertyNotifyEventHandler() {}

var signalProximityInEventId int
var signalProximityInEventMap = make(map[int]WidgetSignalProximityInEventCallback)
var signalProximityInEventLock sync.Mutex

// WidgetSignalProximityInEventCallback is a callback function for a 'proximity-in-event' signal emitted from a Widget.
type WidgetSignalProximityInEventCallback func(event *gdk.EventProximity) bool

func Widget_proximityInEventHandler() {}

var signalProximityOutEventId int
var signalProximityOutEventMap = make(map[int]WidgetSignalProximityOutEventCallback)
var signalProximityOutEventLock sync.Mutex

// WidgetSignalProximityOutEventCallback is a callback function for a 'proximity-out-event' signal emitted from a Widget.
type WidgetSignalProximityOutEventCallback func(event *gdk.EventProximity) bool

func Widget_proximityOutEventHandler() {}

var signalRealizeId int
var signalRealizeMap = make(map[int]WidgetSignalRealizeCallback)
var signalRealizeLock sync.Mutex

// WidgetSignalRealizeCallback is a callback function for a 'realize' signal emitted from a Widget.
type WidgetSignalRealizeCallback func()

func Widget_realizeHandler() {}

var signalScreenChangedId int
var signalScreenChangedMap = make(map[int]WidgetSignalScreenChangedCallback)
var signalScreenChangedLock sync.Mutex

// WidgetSignalScreenChangedCallback is a callback function for a 'screen-changed' signal emitted from a Widget.
type WidgetSignalScreenChangedCallback func(previousScreen *gdk.Screen)

func Widget_screenChangedHandler() {}

var signalScrollEventId int
var signalScrollEventMap = make(map[int]WidgetSignalScrollEventCallback)
var signalScrollEventLock sync.Mutex

// WidgetSignalScrollEventCallback is a callback function for a 'scroll-event' signal emitted from a Widget.
type WidgetSignalScrollEventCallback func(event *gdk.EventScroll) bool

func Widget_scrollEventHandler() {}

var signalSelectionClearEventId int
var signalSelectionClearEventMap = make(map[int]WidgetSignalSelectionClearEventCallback)
var signalSelectionClearEventLock sync.Mutex

// WidgetSignalSelectionClearEventCallback is a callback function for a 'selection-clear-event' signal emitted from a Widget.
type WidgetSignalSelectionClearEventCallback func(event *gdk.EventSelection) bool

func Widget_selectionClearEventHandler() {}

var signalSelectionGetId int
var signalSelectionGetMap = make(map[int]WidgetSignalSelectionGetCallback)
var signalSelectionGetLock sync.Mutex

// WidgetSignalSelectionGetCallback is a callback function for a 'selection-get' signal emitted from a Widget.
type WidgetSignalSelectionGetCallback func(data *SelectionData, info uint32, time uint32)

func Widget_selectionGetHandler() {}

var signalSelectionNotifyEventId int
var signalSelectionNotifyEventMap = make(map[int]WidgetSignalSelectionNotifyEventCallback)
var signalSelectionNotifyEventLock sync.Mutex

// WidgetSignalSelectionNotifyEventCallback is a callback function for a 'selection-notify-event' signal emitted from a Widget.
type WidgetSignalSelectionNotifyEventCallback func(event *gdk.EventSelection) bool

func Widget_selectionNotifyEventHandler() {}

var signalSelectionReceivedId int
var signalSelectionReceivedMap = make(map[int]WidgetSignalSelectionReceivedCallback)
var signalSelectionReceivedLock sync.Mutex

// WidgetSignalSelectionReceivedCallback is a callback function for a 'selection-received' signal emitted from a Widget.
type WidgetSignalSelectionReceivedCallback func(data *SelectionData, time uint32)

func Widget_selectionReceivedHandler() {}

var signalSelectionRequestEventId int
var signalSelectionRequestEventMap = make(map[int]WidgetSignalSelectionRequestEventCallback)
var signalSelectionRequestEventLock sync.Mutex

// WidgetSignalSelectionRequestEventCallback is a callback function for a 'selection-request-event' signal emitted from a Widget.
type WidgetSignalSelectionRequestEventCallback func(event *gdk.EventSelection) bool

func Widget_selectionRequestEventHandler() {}

var signalShowId int
var signalShowMap = make(map[int]WidgetSignalShowCallback)
var signalShowLock sync.Mutex

// WidgetSignalShowCallback is a callback function for a 'show' signal emitted from a Widget.
type WidgetSignalShowCallback func()

func Widget_showHandler() {}

var signalShowHelpId int
var signalShowHelpMap = make(map[int]WidgetSignalShowHelpCallback)
var signalShowHelpLock sync.Mutex

// WidgetSignalShowHelpCallback is a callback function for a 'show-help' signal emitted from a Widget.
type WidgetSignalShowHelpCallback func(helpType WidgetHelpType) bool

func Widget_showHelpHandler() {}

// Unsupported signal : unsupported parameter allocation : Blacklisted record : GdkRectangle

var signalStateChangedId int
var signalStateChangedMap = make(map[int]WidgetSignalStateChangedCallback)
var signalStateChangedLock sync.Mutex

// WidgetSignalStateChangedCallback is a callback function for a 'state-changed' signal emitted from a Widget.
type WidgetSignalStateChangedCallback func(state StateType)

func Widget_stateChangedHandler() {}

var signalStyleSetId int
var signalStyleSetMap = make(map[int]WidgetSignalStyleSetCallback)
var signalStyleSetLock sync.Mutex

// WidgetSignalStyleSetCallback is a callback function for a 'style-set' signal emitted from a Widget.
type WidgetSignalStyleSetCallback func(previousStyle *Style)

func Widget_styleSetHandler() {}

// Unsupported signal : unsupported parameter object : no type generator for Gdk.Event,

var signalUnmapId int
var signalUnmapMap = make(map[int]WidgetSignalUnmapCallback)
var signalUnmapLock sync.Mutex

// WidgetSignalUnmapCallback is a callback function for a 'unmap' signal emitted from a Widget.
type WidgetSignalUnmapCallback func()

func Widget_unmapHandler() {}

var signalUnmapEventId int
var signalUnmapEventMap = make(map[int]WidgetSignalUnmapEventCallback)
var signalUnmapEventLock sync.Mutex

// WidgetSignalUnmapEventCallback is a callback function for a 'unmap-event' signal emitted from a Widget.
type WidgetSignalUnmapEventCallback func(event *gdk.EventAny) bool

func Widget_unmapEventHandler() {}

var signalUnrealizeId int
var signalUnrealizeMap = make(map[int]WidgetSignalUnrealizeCallback)
var signalUnrealizeLock sync.Mutex

// WidgetSignalUnrealizeCallback is a callback function for a 'unrealize' signal emitted from a Widget.
type WidgetSignalUnrealizeCallback func()

func Widget_unrealizeHandler() {}

var signalVisibilityNotifyEventId int
var signalVisibilityNotifyEventMap = make(map[int]WidgetSignalVisibilityNotifyEventCallback)
var signalVisibilityNotifyEventLock sync.Mutex

// WidgetSignalVisibilityNotifyEventCallback is a callback function for a 'visibility-notify-event' signal emitted from a Widget.
type WidgetSignalVisibilityNotifyEventCallback func(event *gdk.EventVisibility) bool

func Widget_visibilityNotifyEventHandler() {}

var signalWindowStateEventId int
var signalWindowStateEventMap = make(map[int]WidgetSignalWindowStateEventCallback)
var signalWindowStateEventLock sync.Mutex

// WidgetSignalWindowStateEventCallback is a callback function for a 'window-state-event' signal emitted from a Widget.
type WidgetSignalWindowStateEventCallback func(event *gdk.EventWindowState) bool

func Widget_windowStateEventHandler() {}

// Unsupported : gtk_widget_new : unsupported parameter type : no type generator for GType, GType

// Activate is a wrapper around the C function gtk_widget_activate.
func (recv *Widget) Activate() bool {
	retC := C.gtk_widget_activate((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddAccelerator is a wrapper around the C function gtk_widget_add_accelerator.
func (recv *Widget) AddAccelerator(accelSignal string, accelGroup *AccelGroup, accelKey uint32, accelMods gdk.ModifierType, accelFlags AccelFlags) {
	c_accel_signal := C.CString(accelSignal)
	defer C.free(unsafe.Pointer(c_accel_signal))

	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	c_accel_flags := (C.GtkAccelFlags)(accelFlags)

	C.gtk_widget_add_accelerator((*C.GtkWidget)(recv.native), c_accel_signal, c_accel_group, c_accel_key, c_accel_mods, c_accel_flags)

	return
}

// AddEvents is a wrapper around the C function gtk_widget_add_events.
func (recv *Widget) AddEvents(events int32) {
	c_events := (C.gint)(events)

	C.gtk_widget_add_events((*C.GtkWidget)(recv.native), c_events)

	return
}

// ChildFocus is a wrapper around the C function gtk_widget_child_focus.
func (recv *Widget) ChildFocus(direction DirectionType) bool {
	c_direction := (C.GtkDirectionType)(direction)

	retC := C.gtk_widget_child_focus((*C.GtkWidget)(recv.native), c_direction)
	retGo := retC == C.TRUE

	return retGo
}

// ChildNotify is a wrapper around the C function gtk_widget_child_notify.
func (recv *Widget) ChildNotify(childProperty string) {
	c_child_property := C.CString(childProperty)
	defer C.free(unsafe.Pointer(c_child_property))

	C.gtk_widget_child_notify((*C.GtkWidget)(recv.native), c_child_property)

	return
}

// ClassPath is a wrapper around the C function gtk_widget_class_path.
func (recv *Widget) ClassPath() (uint32, string, string) {
	var c_path_length C.guint

	var c_path *C.gchar

	var c_path_reversed *C.gchar

	C.gtk_widget_class_path((*C.GtkWidget)(recv.native), &c_path_length, &c_path, &c_path_reversed)

	pathLength := (uint32)(c_path_length)

	path := C.GoString(c_path)
	defer C.free(unsafe.Pointer(c_path))

	pathReversed := C.GoString(c_path_reversed)
	defer C.free(unsafe.Pointer(c_path_reversed))

	return pathLength, path, pathReversed
}

// ComputeExpand is a wrapper around the C function gtk_widget_compute_expand.
func (recv *Widget) ComputeExpand(orientation Orientation) bool {
	c_orientation := (C.GtkOrientation)(orientation)

	retC := C.gtk_widget_compute_expand((*C.GtkWidget)(recv.native), c_orientation)
	retGo := retC == C.TRUE

	return retGo
}

// CreatePangoContext is a wrapper around the C function gtk_widget_create_pango_context.
func (recv *Widget) CreatePangoContext() *pango.Context {
	retC := C.gtk_widget_create_pango_context((*C.GtkWidget)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreatePangoLayout is a wrapper around the C function gtk_widget_create_pango_layout.
func (recv *Widget) CreatePangoLayout(text string) *pango.Layout {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	retC := C.gtk_widget_create_pango_layout((*C.GtkWidget)(recv.native), c_text)
	retGo := pango.LayoutNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Destroy is a wrapper around the C function gtk_widget_destroy.
func (recv *Widget) Destroy() {
	C.gtk_widget_destroy((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_destroyed : unsupported parameter widget_pointer : record with indirection level of 2

// Unsupported : gtk_drag_begin : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// DragCheckThreshold is a wrapper around the C function gtk_drag_check_threshold.
func (recv *Widget) DragCheckThreshold(startX int32, startY int32, currentX int32, currentY int32) bool {
	c_start_x := (C.gint)(startX)

	c_start_y := (C.gint)(startY)

	c_current_x := (C.gint)(currentX)

	c_current_y := (C.gint)(currentY)

	retC := C.gtk_drag_check_threshold((*C.GtkWidget)(recv.native), c_start_x, c_start_y, c_current_x, c_current_y)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_drag_dest_find_target : return type : Blacklisted record : GdkAtom

// DragDestGetTargetList is a wrapper around the C function gtk_drag_dest_get_target_list.
func (recv *Widget) DragDestGetTargetList() *TargetList {
	retC := C.gtk_drag_dest_get_target_list((*C.GtkWidget)(recv.native))
	retGo := TargetListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_drag_dest_set : unsupported parameter targets : no param type

// DragDestSetProxy is a wrapper around the C function gtk_drag_dest_set_proxy.
func (recv *Widget) DragDestSetProxy(proxyWindow *gdk.Window, protocol gdk.DragProtocol, useCoordinates bool) {
	c_proxy_window := (*C.GdkWindow)(proxyWindow.ToC())

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_use_coordinates :=
		boolToGboolean(useCoordinates)

	C.gtk_drag_dest_set_proxy((*C.GtkWidget)(recv.native), c_proxy_window, c_protocol, c_use_coordinates)

	return
}

// DragDestSetTargetList is a wrapper around the C function gtk_drag_dest_set_target_list.
func (recv *Widget) DragDestSetTargetList(targetList *TargetList) {
	c_target_list := (*C.GtkTargetList)(targetList.ToC())

	C.gtk_drag_dest_set_target_list((*C.GtkWidget)(recv.native), c_target_list)

	return
}

// DragDestUnset is a wrapper around the C function gtk_drag_dest_unset.
func (recv *Widget) DragDestUnset() {
	C.gtk_drag_dest_unset((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_drag_get_data : unsupported parameter target : Blacklisted record : GdkAtom

// DragHighlight is a wrapper around the C function gtk_drag_highlight.
func (recv *Widget) DragHighlight() {
	C.gtk_drag_highlight((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_drag_source_set : unsupported parameter targets : no param type

// DragSourceSetIconPixbuf is a wrapper around the C function gtk_drag_source_set_icon_pixbuf.
func (recv *Widget) DragSourceSetIconPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	c_pixbuf := (*C.GdkPixbuf)(pixbuf.ToC())

	C.gtk_drag_source_set_icon_pixbuf((*C.GtkWidget)(recv.native), c_pixbuf)

	return
}

// DragSourceSetIconStock is a wrapper around the C function gtk_drag_source_set_icon_stock.
func (recv *Widget) DragSourceSetIconStock(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_drag_source_set_icon_stock((*C.GtkWidget)(recv.native), c_stock_id)

	return
}

// DragSourceUnset is a wrapper around the C function gtk_drag_source_unset.
func (recv *Widget) DragSourceUnset() {
	C.gtk_drag_source_unset((*C.GtkWidget)(recv.native))

	return
}

// DragUnhighlight is a wrapper around the C function gtk_drag_unhighlight.
func (recv *Widget) DragUnhighlight() {
	C.gtk_drag_unhighlight((*C.GtkWidget)(recv.native))

	return
}

// EnsureStyle is a wrapper around the C function gtk_widget_ensure_style.
func (recv *Widget) EnsureStyle() {
	C.gtk_widget_ensure_style((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_event : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// FreezeChildNotify is a wrapper around the C function gtk_widget_freeze_child_notify.
func (recv *Widget) FreezeChildNotify() {
	C.gtk_widget_freeze_child_notify((*C.GtkWidget)(recv.native))

	return
}

// GetAccessible is a wrapper around the C function gtk_widget_get_accessible.
func (recv *Widget) GetAccessible() *atk.Object {
	retC := C.gtk_widget_get_accessible((*C.GtkWidget)(recv.native))
	retGo := atk.ObjectNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAllocatedHeight is a wrapper around the C function gtk_widget_get_allocated_height.
func (recv *Widget) GetAllocatedHeight() int32 {
	retC := C.gtk_widget_get_allocated_height((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetAllocatedWidth is a wrapper around the C function gtk_widget_get_allocated_width.
func (recv *Widget) GetAllocatedWidth() int32 {
	retC := C.gtk_widget_get_allocated_width((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : gtk_widget_get_ancestor : unsupported parameter widget_type : no type generator for GType, GType

// GetChildRequisition is a wrapper around the C function gtk_widget_get_child_requisition.
func (recv *Widget) GetChildRequisition() *Requisition {
	var c_requisition C.GtkRequisition

	C.gtk_widget_get_child_requisition((*C.GtkWidget)(recv.native), &c_requisition)

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return requisition
}

// GetChildVisible is a wrapper around the C function gtk_widget_get_child_visible.
func (recv *Widget) GetChildVisible() bool {
	retC := C.gtk_widget_get_child_visible((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetCompositeName is a wrapper around the C function gtk_widget_get_composite_name.
func (recv *Widget) GetCompositeName() string {
	retC := C.gtk_widget_get_composite_name((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDirection is a wrapper around the C function gtk_widget_get_direction.
func (recv *Widget) GetDirection() TextDirection {
	retC := C.gtk_widget_get_direction((*C.GtkWidget)(recv.native))
	retGo := (TextDirection)(retC)

	return retGo
}

// GetEvents is a wrapper around the C function gtk_widget_get_events.
func (recv *Widget) GetEvents() int32 {
	retC := C.gtk_widget_get_events((*C.GtkWidget)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetHalign is a wrapper around the C function gtk_widget_get_halign.
func (recv *Widget) GetHalign() Align {
	retC := C.gtk_widget_get_halign((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// GetHexpand is a wrapper around the C function gtk_widget_get_hexpand.
func (recv *Widget) GetHexpand() bool {
	retC := C.gtk_widget_get_hexpand((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHexpandSet is a wrapper around the C function gtk_widget_get_hexpand_set.
func (recv *Widget) GetHexpandSet() bool {
	retC := C.gtk_widget_get_hexpand_set((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetModifierStyle is a wrapper around the C function gtk_widget_get_modifier_style.
func (recv *Widget) GetModifierStyle() *RcStyle {
	retC := C.gtk_widget_get_modifier_style((*C.GtkWidget)(recv.native))
	retGo := RcStyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_widget_get_name.
func (recv *Widget) GetName() string {
	retC := C.gtk_widget_get_name((*C.GtkWidget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPangoContext is a wrapper around the C function gtk_widget_get_pango_context.
func (recv *Widget) GetPangoContext() *pango.Context {
	retC := C.gtk_widget_get_pango_context((*C.GtkWidget)(recv.native))
	retGo := pango.ContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParent is a wrapper around the C function gtk_widget_get_parent.
func (recv *Widget) GetParent() *Widget {
	retC := C.gtk_widget_get_parent((*C.GtkWidget)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetParentWindow is a wrapper around the C function gtk_widget_get_parent_window.
func (recv *Widget) GetParentWindow() *gdk.Window {
	retC := C.gtk_widget_get_parent_window((*C.GtkWidget)(recv.native))
	retGo := gdk.WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPath is a wrapper around the C function gtk_widget_get_path.
func (recv *Widget) GetPath() *WidgetPath {
	retC := C.gtk_widget_get_path((*C.GtkWidget)(recv.native))
	retGo := WidgetPathNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPointer is a wrapper around the C function gtk_widget_get_pointer.
func (recv *Widget) GetPointer() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_widget_get_pointer((*C.GtkWidget)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// GetSettings is a wrapper around the C function gtk_widget_get_settings.
func (recv *Widget) GetSettings() *Settings {
	retC := C.gtk_widget_get_settings((*C.GtkWidget)(recv.native))
	retGo := SettingsNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSizeRequest is a wrapper around the C function gtk_widget_get_size_request.
func (recv *Widget) GetSizeRequest() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_widget_get_size_request((*C.GtkWidget)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetStyle is a wrapper around the C function gtk_widget_get_style.
func (recv *Widget) GetStyle() *Style {
	retC := C.gtk_widget_get_style((*C.GtkWidget)(recv.native))
	retGo := StyleNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetStyleContext is a wrapper around the C function gtk_widget_get_style_context.
func (recv *Widget) GetStyleContext() *StyleContext {
	retC := C.gtk_widget_get_style_context((*C.GtkWidget)(recv.native))
	retGo := StyleContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSupportMultidevice is a wrapper around the C function gtk_widget_get_support_multidevice.
func (recv *Widget) GetSupportMultidevice() bool {
	retC := C.gtk_widget_get_support_multidevice((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_get_template_child : unsupported parameter widget_type : no type generator for GType, GType

// GetToplevel is a wrapper around the C function gtk_widget_get_toplevel.
func (recv *Widget) GetToplevel() *Widget {
	retC := C.gtk_widget_get_toplevel((*C.GtkWidget)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValign is a wrapper around the C function gtk_widget_get_valign.
func (recv *Widget) GetValign() Align {
	retC := C.gtk_widget_get_valign((*C.GtkWidget)(recv.native))
	retGo := (Align)(retC)

	return retGo
}

// GetVexpand is a wrapper around the C function gtk_widget_get_vexpand.
func (recv *Widget) GetVexpand() bool {
	retC := C.gtk_widget_get_vexpand((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVexpandSet is a wrapper around the C function gtk_widget_get_vexpand_set.
func (recv *Widget) GetVexpandSet() bool {
	retC := C.gtk_widget_get_vexpand_set((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisual is a wrapper around the C function gtk_widget_get_visual.
func (recv *Widget) GetVisual() *gdk.Visual {
	retC := C.gtk_widget_get_visual((*C.GtkWidget)(recv.native))
	retGo := gdk.VisualNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GrabAdd is a wrapper around the C function gtk_grab_add.
func (recv *Widget) GrabAdd() {
	C.gtk_grab_add((*C.GtkWidget)(recv.native))

	return
}

// GrabDefault is a wrapper around the C function gtk_widget_grab_default.
func (recv *Widget) GrabDefault() {
	C.gtk_widget_grab_default((*C.GtkWidget)(recv.native))

	return
}

// GrabFocus is a wrapper around the C function gtk_widget_grab_focus.
func (recv *Widget) GrabFocus() {
	C.gtk_widget_grab_focus((*C.GtkWidget)(recv.native))

	return
}

// GrabRemove is a wrapper around the C function gtk_grab_remove.
func (recv *Widget) GrabRemove() {
	C.gtk_grab_remove((*C.GtkWidget)(recv.native))

	return
}

// Hide is a wrapper around the C function gtk_widget_hide.
func (recv *Widget) Hide() {
	C.gtk_widget_hide((*C.GtkWidget)(recv.native))

	return
}

// HideOnDelete is a wrapper around the C function gtk_widget_hide_on_delete.
func (recv *Widget) HideOnDelete() bool {
	retC := C.gtk_widget_hide_on_delete((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// InDestruction is a wrapper around the C function gtk_widget_in_destruction.
func (recv *Widget) InDestruction() bool {
	retC := C.gtk_widget_in_destruction((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_intersect : unsupported parameter area : Blacklisted record : GdkRectangle

// IsAncestor is a wrapper around the C function gtk_widget_is_ancestor.
func (recv *Widget) IsAncestor(ancestor *Widget) bool {
	c_ancestor := (*C.GtkWidget)(ancestor.ToC())

	retC := C.gtk_widget_is_ancestor((*C.GtkWidget)(recv.native), c_ancestor)
	retGo := retC == C.TRUE

	return retGo
}

// IsFocus is a wrapper around the C function gtk_widget_is_focus.
func (recv *Widget) IsFocus() bool {
	retC := C.gtk_widget_is_focus((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListAccelClosures is a wrapper around the C function gtk_widget_list_accel_closures.
func (recv *Widget) ListAccelClosures() *glib.List {
	retC := C.gtk_widget_list_accel_closures((*C.GtkWidget)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Map is a wrapper around the C function gtk_widget_map.
func (recv *Widget) Map() {
	C.gtk_widget_map((*C.GtkWidget)(recv.native))

	return
}

// MnemonicActivate is a wrapper around the C function gtk_widget_mnemonic_activate.
func (recv *Widget) MnemonicActivate(groupCycling bool) bool {
	c_group_cycling :=
		boolToGboolean(groupCycling)

	retC := C.gtk_widget_mnemonic_activate((*C.GtkWidget)(recv.native), c_group_cycling)
	retGo := retC == C.TRUE

	return retGo
}

// ModifyBase is a wrapper around the C function gtk_widget_modify_base.
func (recv *Widget) ModifyBase(state StateType, color *gdk.Color) {
	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_widget_modify_base((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// ModifyBg is a wrapper around the C function gtk_widget_modify_bg.
func (recv *Widget) ModifyBg(state StateType, color *gdk.Color) {
	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_widget_modify_bg((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// ModifyFg is a wrapper around the C function gtk_widget_modify_fg.
func (recv *Widget) ModifyFg(state StateType, color *gdk.Color) {
	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_widget_modify_fg((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// ModifyFont is a wrapper around the C function gtk_widget_modify_font.
func (recv *Widget) ModifyFont(fontDesc *pango.FontDescription) {
	c_font_desc := (*C.PangoFontDescription)(fontDesc.ToC())

	C.gtk_widget_modify_font((*C.GtkWidget)(recv.native), c_font_desc)

	return
}

// ModifyStyle is a wrapper around the C function gtk_widget_modify_style.
func (recv *Widget) ModifyStyle(style *RcStyle) {
	c_style := (*C.GtkRcStyle)(style.ToC())

	C.gtk_widget_modify_style((*C.GtkWidget)(recv.native), c_style)

	return
}

// ModifyText is a wrapper around the C function gtk_widget_modify_text.
func (recv *Widget) ModifyText(state StateType, color *gdk.Color) {
	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(color.ToC())

	C.gtk_widget_modify_text((*C.GtkWidget)(recv.native), c_state, c_color)

	return
}

// Path is a wrapper around the C function gtk_widget_path.
func (recv *Widget) Path() (uint32, string, string) {
	var c_path_length C.guint

	var c_path *C.gchar

	var c_path_reversed *C.gchar

	C.gtk_widget_path((*C.GtkWidget)(recv.native), &c_path_length, &c_path, &c_path_reversed)

	pathLength := (uint32)(c_path_length)

	path := C.GoString(c_path)
	defer C.free(unsafe.Pointer(c_path))

	pathReversed := C.GoString(c_path_reversed)
	defer C.free(unsafe.Pointer(c_path_reversed))

	return pathLength, path, pathReversed
}

// QueueComputeExpand is a wrapper around the C function gtk_widget_queue_compute_expand.
func (recv *Widget) QueueComputeExpand() {
	C.gtk_widget_queue_compute_expand((*C.GtkWidget)(recv.native))

	return
}

// QueueDraw is a wrapper around the C function gtk_widget_queue_draw.
func (recv *Widget) QueueDraw() {
	C.gtk_widget_queue_draw((*C.GtkWidget)(recv.native))

	return
}

// QueueDrawArea is a wrapper around the C function gtk_widget_queue_draw_area.
func (recv *Widget) QueueDrawArea(x int32, y int32, width int32, height int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_widget_queue_draw_area((*C.GtkWidget)(recv.native), c_x, c_y, c_width, c_height)

	return
}

// QueueResize is a wrapper around the C function gtk_widget_queue_resize.
func (recv *Widget) QueueResize() {
	C.gtk_widget_queue_resize((*C.GtkWidget)(recv.native))

	return
}

// Realize is a wrapper around the C function gtk_widget_realize.
func (recv *Widget) Realize() {
	C.gtk_widget_realize((*C.GtkWidget)(recv.native))

	return
}

// RegionIntersect is a wrapper around the C function gtk_widget_region_intersect.
func (recv *Widget) RegionIntersect(region *cairo.Region) *cairo.Region {
	c_region := (*C.cairo_region_t)(region.ToC())

	retC := C.gtk_widget_region_intersect((*C.GtkWidget)(recv.native), c_region)
	retGo := cairo.RegionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveAccelerator is a wrapper around the C function gtk_widget_remove_accelerator.
func (recv *Widget) RemoveAccelerator(accelGroup *AccelGroup, accelKey uint32, accelMods gdk.ModifierType) bool {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	c_accel_key := (C.guint)(accelKey)

	c_accel_mods := (C.GdkModifierType)(accelMods)

	retC := C.gtk_widget_remove_accelerator((*C.GtkWidget)(recv.native), c_accel_group, c_accel_key, c_accel_mods)
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : gtk_widget_render_icon : unsupported parameter size : no type generator for gint, GtkIconSize

// Reparent is a wrapper around the C function gtk_widget_reparent.
func (recv *Widget) Reparent(newParent *Widget) {
	c_new_parent := (*C.GtkWidget)(newParent.ToC())

	C.gtk_widget_reparent((*C.GtkWidget)(recv.native), c_new_parent)

	return
}

// ResetRcStyles is a wrapper around the C function gtk_widget_reset_rc_styles.
func (recv *Widget) ResetRcStyles() {
	C.gtk_widget_reset_rc_styles((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_send_expose : unsupported parameter event : no type generator for Gdk.Event, GdkEvent*

// SetAccelPath is a wrapper around the C function gtk_widget_set_accel_path.
func (recv *Widget) SetAccelPath(accelPath string, accelGroup *AccelGroup) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_widget_set_accel_path((*C.GtkWidget)(recv.native), c_accel_path, c_accel_group)

	return
}

// SetAppPaintable is a wrapper around the C function gtk_widget_set_app_paintable.
func (recv *Widget) SetAppPaintable(appPaintable bool) {
	c_app_paintable :=
		boolToGboolean(appPaintable)

	C.gtk_widget_set_app_paintable((*C.GtkWidget)(recv.native), c_app_paintable)

	return
}

// SetChildVisible is a wrapper around the C function gtk_widget_set_child_visible.
func (recv *Widget) SetChildVisible(isVisible bool) {
	c_is_visible :=
		boolToGboolean(isVisible)

	C.gtk_widget_set_child_visible((*C.GtkWidget)(recv.native), c_is_visible)

	return
}

// SetCompositeName is a wrapper around the C function gtk_widget_set_composite_name.
func (recv *Widget) SetCompositeName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_set_composite_name((*C.GtkWidget)(recv.native), c_name)

	return
}

// SetDirection is a wrapper around the C function gtk_widget_set_direction.
func (recv *Widget) SetDirection(dir TextDirection) {
	c_dir := (C.GtkTextDirection)(dir)

	C.gtk_widget_set_direction((*C.GtkWidget)(recv.native), c_dir)

	return
}

// SetDoubleBuffered is a wrapper around the C function gtk_widget_set_double_buffered.
func (recv *Widget) SetDoubleBuffered(doubleBuffered bool) {
	c_double_buffered :=
		boolToGboolean(doubleBuffered)

	C.gtk_widget_set_double_buffered((*C.GtkWidget)(recv.native), c_double_buffered)

	return
}

// SetEvents is a wrapper around the C function gtk_widget_set_events.
func (recv *Widget) SetEvents(events int32) {
	c_events := (C.gint)(events)

	C.gtk_widget_set_events((*C.GtkWidget)(recv.native), c_events)

	return
}

// SetHalign is a wrapper around the C function gtk_widget_set_halign.
func (recv *Widget) SetHalign(align Align) {
	c_align := (C.GtkAlign)(align)

	C.gtk_widget_set_halign((*C.GtkWidget)(recv.native), c_align)

	return
}

// SetHexpand is a wrapper around the C function gtk_widget_set_hexpand.
func (recv *Widget) SetHexpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_widget_set_hexpand((*C.GtkWidget)(recv.native), c_expand)

	return
}

// SetHexpandSet is a wrapper around the C function gtk_widget_set_hexpand_set.
func (recv *Widget) SetHexpandSet(set bool) {
	c_set :=
		boolToGboolean(set)

	C.gtk_widget_set_hexpand_set((*C.GtkWidget)(recv.native), c_set)

	return
}

// SetName is a wrapper around the C function gtk_widget_set_name.
func (recv *Widget) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_set_name((*C.GtkWidget)(recv.native), c_name)

	return
}

// SetParent is a wrapper around the C function gtk_widget_set_parent.
func (recv *Widget) SetParent(parent *Widget) {
	c_parent := (*C.GtkWidget)(parent.ToC())

	C.gtk_widget_set_parent((*C.GtkWidget)(recv.native), c_parent)

	return
}

// SetParentWindow is a wrapper around the C function gtk_widget_set_parent_window.
func (recv *Widget) SetParentWindow(parentWindow *gdk.Window) {
	c_parent_window := (*C.GdkWindow)(parentWindow.ToC())

	C.gtk_widget_set_parent_window((*C.GtkWidget)(recv.native), c_parent_window)

	return
}

// SetRedrawOnAllocate is a wrapper around the C function gtk_widget_set_redraw_on_allocate.
func (recv *Widget) SetRedrawOnAllocate(redrawOnAllocate bool) {
	c_redraw_on_allocate :=
		boolToGboolean(redrawOnAllocate)

	C.gtk_widget_set_redraw_on_allocate((*C.GtkWidget)(recv.native), c_redraw_on_allocate)

	return
}

// SetSensitive is a wrapper around the C function gtk_widget_set_sensitive.
func (recv *Widget) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_widget_set_sensitive((*C.GtkWidget)(recv.native), c_sensitive)

	return
}

// SetSizeRequest is a wrapper around the C function gtk_widget_set_size_request.
func (recv *Widget) SetSizeRequest(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_widget_set_size_request((*C.GtkWidget)(recv.native), c_width, c_height)

	return
}

// SetState is a wrapper around the C function gtk_widget_set_state.
func (recv *Widget) SetState(state StateType) {
	c_state := (C.GtkStateType)(state)

	C.gtk_widget_set_state((*C.GtkWidget)(recv.native), c_state)

	return
}

// SetStyle is a wrapper around the C function gtk_widget_set_style.
func (recv *Widget) SetStyle(style *Style) {
	c_style := (*C.GtkStyle)(style.ToC())

	C.gtk_widget_set_style((*C.GtkWidget)(recv.native), c_style)

	return
}

// SetValign is a wrapper around the C function gtk_widget_set_valign.
func (recv *Widget) SetValign(align Align) {
	c_align := (C.GtkAlign)(align)

	C.gtk_widget_set_valign((*C.GtkWidget)(recv.native), c_align)

	return
}

// SetVexpand is a wrapper around the C function gtk_widget_set_vexpand.
func (recv *Widget) SetVexpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_widget_set_vexpand((*C.GtkWidget)(recv.native), c_expand)

	return
}

// SetVexpandSet is a wrapper around the C function gtk_widget_set_vexpand_set.
func (recv *Widget) SetVexpandSet(set bool) {
	c_set :=
		boolToGboolean(set)

	C.gtk_widget_set_vexpand_set((*C.GtkWidget)(recv.native), c_set)

	return
}

// SetVisual is a wrapper around the C function gtk_widget_set_visual.
func (recv *Widget) SetVisual(visual *gdk.Visual) {
	c_visual := (*C.GdkVisual)(visual.ToC())

	C.gtk_widget_set_visual((*C.GtkWidget)(recv.native), c_visual)

	return
}

// Show is a wrapper around the C function gtk_widget_show.
func (recv *Widget) Show() {
	C.gtk_widget_show((*C.GtkWidget)(recv.native))

	return
}

// ShowAll is a wrapper around the C function gtk_widget_show_all.
func (recv *Widget) ShowAll() {
	C.gtk_widget_show_all((*C.GtkWidget)(recv.native))

	return
}

// ShowNow is a wrapper around the C function gtk_widget_show_now.
func (recv *Widget) ShowNow() {
	C.gtk_widget_show_now((*C.GtkWidget)(recv.native))

	return
}

// Unsupported : gtk_widget_size_allocate : unsupported parameter allocation : Blacklisted record : GdkRectangle

// SizeRequest is a wrapper around the C function gtk_widget_size_request.
func (recv *Widget) SizeRequest() *Requisition {
	var c_requisition C.GtkRequisition

	C.gtk_widget_size_request((*C.GtkWidget)(recv.native), &c_requisition)

	requisition := RequisitionNewFromC(unsafe.Pointer(&c_requisition))

	return requisition
}

// Unsupported : gtk_widget_style_get : unsupported parameter ... : varargs

// StyleGetProperty is a wrapper around the C function gtk_widget_style_get_property.
func (recv *Widget) StyleGetProperty(propertyName string, value *gobject.Value) {
	c_property_name := C.CString(propertyName)
	defer C.free(unsafe.Pointer(c_property_name))

	c_value := (*C.GValue)(value.ToC())

	C.gtk_widget_style_get_property((*C.GtkWidget)(recv.native), c_property_name, c_value)

	return
}

// Unsupported : gtk_widget_style_get_valist : unsupported parameter var_args : no type generator for va_list, va_list

// ThawChildNotify is a wrapper around the C function gtk_widget_thaw_child_notify.
func (recv *Widget) ThawChildNotify() {
	C.gtk_widget_thaw_child_notify((*C.GtkWidget)(recv.native))

	return
}

// TranslateCoordinates is a wrapper around the C function gtk_widget_translate_coordinates.
func (recv *Widget) TranslateCoordinates(destWidget *Widget, srcX int32, srcY int32) (bool, int32, int32) {
	c_dest_widget := (*C.GtkWidget)(destWidget.ToC())

	c_src_x := (C.gint)(srcX)

	c_src_y := (C.gint)(srcY)

	var c_dest_x C.gint

	var c_dest_y C.gint

	retC := C.gtk_widget_translate_coordinates((*C.GtkWidget)(recv.native), c_dest_widget, c_src_x, c_src_y, &c_dest_x, &c_dest_y)
	retGo := retC == C.TRUE

	destX := (int32)(c_dest_x)

	destY := (int32)(c_dest_y)

	return retGo, destX, destY
}

// Unmap is a wrapper around the C function gtk_widget_unmap.
func (recv *Widget) Unmap() {
	C.gtk_widget_unmap((*C.GtkWidget)(recv.native))

	return
}

// Unparent is a wrapper around the C function gtk_widget_unparent.
func (recv *Widget) Unparent() {
	C.gtk_widget_unparent((*C.GtkWidget)(recv.native))

	return
}

// Unrealize is a wrapper around the C function gtk_widget_unrealize.
func (recv *Widget) Unrealize() {
	C.gtk_widget_unrealize((*C.GtkWidget)(recv.native))

	return
}

// WidgetAccessible is a wrapper around the C record GtkWidgetAccessible.
type WidgetAccessible struct {
	native *C.GtkWidgetAccessible
	// parent : record
	// priv : record
}

func WidgetAccessibleNewFromC(u unsafe.Pointer) *WidgetAccessible {
	c := (*C.GtkWidgetAccessible)(u)
	if c == nil {
		return nil
	}

	g := &WidgetAccessible{native: c}

	return g
}

func (recv *WidgetAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Accessible upcasts to *Accessible
func (recv *WidgetAccessible) Accessible() *Accessible {
	return AccessibleNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *WidgetAccessible) Object() *atk.Object {
	return recv.Accessible().Object()
}

// CastToWidget down casts any arbitary Object to WidgetAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a WidgetAccessible.
func CastToWidgetAccessible(object *gobject.Object) *WidgetAccessible {
	return WidgetAccessibleNewFromC(object.ToC())
}

// Window is a wrapper around the C record GtkWindow.
type Window struct {
	native *C.GtkWindow
	// bin : record
	// priv : record
}

func WindowNewFromC(u unsafe.Pointer) *Window {
	c := (*C.GtkWindow)(u)
	if c == nil {
		return nil
	}

	g := &Window{native: c}

	return g
}

func (recv *Window) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Bin upcasts to *Bin
func (recv *Window) Bin() *Bin {
	return BinNewFromC(unsafe.Pointer(recv.native))
}

// Container upcasts to *Container
func (recv *Window) Container() *Container {
	return recv.Bin().Container()
}

// Widget upcasts to *Widget
func (recv *Window) Widget() *Widget {
	return recv.Bin().Widget()
}

// InitiallyUnowned upcasts to *InitiallyUnowned
func (recv *Window) InitiallyUnowned() *gobject.InitiallyUnowned {
	return recv.Bin().InitiallyUnowned()
}

// Object upcasts to *Object
func (recv *Window) Object() *gobject.Object {
	return recv.Bin().Object()
}

// CastToWidget down casts any arbitary Object to Window.
// Exercise care, as this is a potentially dangerous function if the Object is not a Window.
func CastToWindow(object *gobject.Object) *Window {
	return WindowNewFromC(object.ToC())
}

var signalActivateDefaultId int
var signalActivateDefaultMap = make(map[int]WindowSignalActivateDefaultCallback)
var signalActivateDefaultLock sync.Mutex

// WindowSignalActivateDefaultCallback is a callback function for a 'activate-default' signal emitted from a Window.
type WindowSignalActivateDefaultCallback func()

func Window_activateDefaultHandler() {}

var signalActivateFocusId int
var signalActivateFocusMap = make(map[int]WindowSignalActivateFocusCallback)
var signalActivateFocusLock sync.Mutex

// WindowSignalActivateFocusCallback is a callback function for a 'activate-focus' signal emitted from a Window.
type WindowSignalActivateFocusCallback func()

func Window_activateFocusHandler() {}

var signalEnableDebuggingId int
var signalEnableDebuggingMap = make(map[int]WindowSignalEnableDebuggingCallback)
var signalEnableDebuggingLock sync.Mutex

// WindowSignalEnableDebuggingCallback is a callback function for a 'enable-debugging' signal emitted from a Window.
type WindowSignalEnableDebuggingCallback func(toggle bool) bool

func Window_enableDebuggingHandler() {}

var signalKeysChangedId int
var signalKeysChangedMap = make(map[int]WindowSignalKeysChangedCallback)
var signalKeysChangedLock sync.Mutex

// WindowSignalKeysChangedCallback is a callback function for a 'keys-changed' signal emitted from a Window.
type WindowSignalKeysChangedCallback func()

func Window_keysChangedHandler() {}

var signalSetFocusId int
var signalSetFocusMap = make(map[int]WindowSignalSetFocusCallback)
var signalSetFocusLock sync.Mutex

// WindowSignalSetFocusCallback is a callback function for a 'set-focus' signal emitted from a Window.
type WindowSignalSetFocusCallback func(object *Widget)

func Window_setFocusHandler() {}

// WindowNew is a wrapper around the C function gtk_window_new.
func WindowNew(type_ WindowType) *Window {
	c_type := (C.GtkWindowType)(type_)

	retC := C.gtk_window_new(c_type)
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ActivateDefault is a wrapper around the C function gtk_window_activate_default.
func (recv *Window) ActivateDefault() bool {
	retC := C.gtk_window_activate_default((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ActivateFocus is a wrapper around the C function gtk_window_activate_focus.
func (recv *Window) ActivateFocus() bool {
	retC := C.gtk_window_activate_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// AddAccelGroup is a wrapper around the C function gtk_window_add_accel_group.
func (recv *Window) AddAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_window_add_accel_group((*C.GtkWindow)(recv.native), c_accel_group)

	return
}

// AddMnemonic is a wrapper around the C function gtk_window_add_mnemonic.
func (recv *Window) AddMnemonic(keyval uint32, target *Widget) {
	c_keyval := (C.guint)(keyval)

	c_target := (*C.GtkWidget)(target.ToC())

	C.gtk_window_add_mnemonic((*C.GtkWindow)(recv.native), c_keyval, c_target)

	return
}

// BeginMoveDrag is a wrapper around the C function gtk_window_begin_move_drag.
func (recv *Window) BeginMoveDrag(button int32, rootX int32, rootY int32, timestamp uint32) {
	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_begin_move_drag((*C.GtkWindow)(recv.native), c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// BeginResizeDrag is a wrapper around the C function gtk_window_begin_resize_drag.
func (recv *Window) BeginResizeDrag(edge gdk.WindowEdge, button int32, rootX int32, rootY int32, timestamp uint32) {
	c_edge := (C.GdkWindowEdge)(edge)

	c_button := (C.gint)(button)

	c_root_x := (C.gint)(rootX)

	c_root_y := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_begin_resize_drag((*C.GtkWindow)(recv.native), c_edge, c_button, c_root_x, c_root_y, c_timestamp)

	return
}

// Deiconify is a wrapper around the C function gtk_window_deiconify.
func (recv *Window) Deiconify() {
	C.gtk_window_deiconify((*C.GtkWindow)(recv.native))

	return
}

// GetDecorated is a wrapper around the C function gtk_window_get_decorated.
func (recv *Window) GetDecorated() bool {
	retC := C.gtk_window_get_decorated((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDefaultSize is a wrapper around the C function gtk_window_get_default_size.
func (recv *Window) GetDefaultSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_window_get_default_size((*C.GtkWindow)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetDestroyWithParent is a wrapper around the C function gtk_window_get_destroy_with_parent.
func (recv *Window) GetDestroyWithParent() bool {
	retC := C.gtk_window_get_destroy_with_parent((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFocus is a wrapper around the C function gtk_window_get_focus.
func (recv *Window) GetFocus() *Widget {
	retC := C.gtk_window_get_focus((*C.GtkWindow)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGravity is a wrapper around the C function gtk_window_get_gravity.
func (recv *Window) GetGravity() gdk.Gravity {
	retC := C.gtk_window_get_gravity((*C.GtkWindow)(recv.native))
	retGo := (gdk.Gravity)(retC)

	return retGo
}

// GetIcon is a wrapper around the C function gtk_window_get_icon.
func (recv *Window) GetIcon() *gdkpixbuf.Pixbuf {
	retC := C.gtk_window_get_icon((*C.GtkWindow)(recv.native))
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconList is a wrapper around the C function gtk_window_get_icon_list.
func (recv *Window) GetIconList() *glib.List {
	retC := C.gtk_window_get_icon_list((*C.GtkWindow)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMnemonicModifier is a wrapper around the C function gtk_window_get_mnemonic_modifier.
func (recv *Window) GetMnemonicModifier() gdk.ModifierType {
	retC := C.gtk_window_get_mnemonic_modifier((*C.GtkWindow)(recv.native))
	retGo := (gdk.ModifierType)(retC)

	return retGo
}

// GetModal is a wrapper around the C function gtk_window_get_modal.
func (recv *Window) GetModal() bool {
	retC := C.gtk_window_get_modal((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPosition is a wrapper around the C function gtk_window_get_position.
func (recv *Window) GetPosition() (int32, int32) {
	var c_root_x C.gint

	var c_root_y C.gint

	C.gtk_window_get_position((*C.GtkWindow)(recv.native), &c_root_x, &c_root_y)

	rootX := (int32)(c_root_x)

	rootY := (int32)(c_root_y)

	return rootX, rootY
}

// GetResizable is a wrapper around the C function gtk_window_get_resizable.
func (recv *Window) GetResizable() bool {
	retC := C.gtk_window_get_resizable((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetRole is a wrapper around the C function gtk_window_get_role.
func (recv *Window) GetRole() string {
	retC := C.gtk_window_get_role((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSize is a wrapper around the C function gtk_window_get_size.
func (recv *Window) GetSize() (int32, int32) {
	var c_width C.gint

	var c_height C.gint

	C.gtk_window_get_size((*C.GtkWindow)(recv.native), &c_width, &c_height)

	width := (int32)(c_width)

	height := (int32)(c_height)

	return width, height
}

// GetTitle is a wrapper around the C function gtk_window_get_title.
func (recv *Window) GetTitle() string {
	retC := C.gtk_window_get_title((*C.GtkWindow)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetTransientFor is a wrapper around the C function gtk_window_get_transient_for.
func (recv *Window) GetTransientFor() *Window {
	retC := C.gtk_window_get_transient_for((*C.GtkWindow)(recv.native))
	retGo := WindowNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetTypeHint is a wrapper around the C function gtk_window_get_type_hint.
func (recv *Window) GetTypeHint() gdk.WindowTypeHint {
	retC := C.gtk_window_get_type_hint((*C.GtkWindow)(recv.native))
	retGo := (gdk.WindowTypeHint)(retC)

	return retGo
}

// HasGroup is a wrapper around the C function gtk_window_has_group.
func (recv *Window) HasGroup() bool {
	retC := C.gtk_window_has_group((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Iconify is a wrapper around the C function gtk_window_iconify.
func (recv *Window) Iconify() {
	C.gtk_window_iconify((*C.GtkWindow)(recv.native))

	return
}

// Maximize is a wrapper around the C function gtk_window_maximize.
func (recv *Window) Maximize() {
	C.gtk_window_maximize((*C.GtkWindow)(recv.native))

	return
}

// MnemonicActivate is a wrapper around the C function gtk_window_mnemonic_activate.
func (recv *Window) MnemonicActivate(keyval uint32, modifier gdk.ModifierType) bool {
	c_keyval := (C.guint)(keyval)

	c_modifier := (C.GdkModifierType)(modifier)

	retC := C.gtk_window_mnemonic_activate((*C.GtkWindow)(recv.native), c_keyval, c_modifier)
	retGo := retC == C.TRUE

	return retGo
}

// Move is a wrapper around the C function gtk_window_move.
func (recv *Window) Move(x int32, y int32) {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_window_move((*C.GtkWindow)(recv.native), c_x, c_y)

	return
}

// ParseGeometry is a wrapper around the C function gtk_window_parse_geometry.
func (recv *Window) ParseGeometry(geometry string) bool {
	c_geometry := C.CString(geometry)
	defer C.free(unsafe.Pointer(c_geometry))

	retC := C.gtk_window_parse_geometry((*C.GtkWindow)(recv.native), c_geometry)
	retGo := retC == C.TRUE

	return retGo
}

// Present is a wrapper around the C function gtk_window_present.
func (recv *Window) Present() {
	C.gtk_window_present((*C.GtkWindow)(recv.native))

	return
}

// RemoveAccelGroup is a wrapper around the C function gtk_window_remove_accel_group.
func (recv *Window) RemoveAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(accelGroup.ToC())

	C.gtk_window_remove_accel_group((*C.GtkWindow)(recv.native), c_accel_group)

	return
}

// RemoveMnemonic is a wrapper around the C function gtk_window_remove_mnemonic.
func (recv *Window) RemoveMnemonic(keyval uint32, target *Widget) {
	c_keyval := (C.guint)(keyval)

	c_target := (*C.GtkWidget)(target.ToC())

	C.gtk_window_remove_mnemonic((*C.GtkWindow)(recv.native), c_keyval, c_target)

	return
}

// ReshowWithInitialSize is a wrapper around the C function gtk_window_reshow_with_initial_size.
func (recv *Window) ReshowWithInitialSize() {
	C.gtk_window_reshow_with_initial_size((*C.GtkWindow)(recv.native))

	return
}

// Resize is a wrapper around the C function gtk_window_resize.
func (recv *Window) Resize(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_resize((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetDecorated is a wrapper around the C function gtk_window_set_decorated.
func (recv *Window) SetDecorated(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_decorated((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetDefault is a wrapper around the C function gtk_window_set_default.
func (recv *Window) SetDefault(defaultWidget *Widget) {
	c_default_widget := (*C.GtkWidget)(defaultWidget.ToC())

	C.gtk_window_set_default((*C.GtkWindow)(recv.native), c_default_widget)

	return
}

// SetDefaultSize is a wrapper around the C function gtk_window_set_default_size.
func (recv *Window) SetDefaultSize(width int32, height int32) {
	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_set_default_size((*C.GtkWindow)(recv.native), c_width, c_height)

	return
}

// SetDestroyWithParent is a wrapper around the C function gtk_window_set_destroy_with_parent.
func (recv *Window) SetDestroyWithParent(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_destroy_with_parent((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetFocus is a wrapper around the C function gtk_window_set_focus.
func (recv *Window) SetFocus(focus *Widget) {
	c_focus := (*C.GtkWidget)(focus.ToC())

	C.gtk_window_set_focus((*C.GtkWindow)(recv.native), c_focus)

	return
}

// SetGeometryHints is a wrapper around the C function gtk_window_set_geometry_hints.
func (recv *Window) SetGeometryHints(geometryWidget *Widget, geometry *gdk.Geometry, geomMask gdk.WindowHints) {
	c_geometry_widget := (*C.GtkWidget)(geometryWidget.ToC())

	c_geometry := (*C.GdkGeometry)(geometry.ToC())

	c_geom_mask := (C.GdkWindowHints)(geomMask)

	C.gtk_window_set_geometry_hints((*C.GtkWindow)(recv.native), c_geometry_widget, c_geometry, c_geom_mask)

	return
}

// SetGravity is a wrapper around the C function gtk_window_set_gravity.
func (recv *Window) SetGravity(gravity gdk.Gravity) {
	c_gravity := (C.GdkGravity)(gravity)

	C.gtk_window_set_gravity((*C.GtkWindow)(recv.native), c_gravity)

	return
}

// SetIcon is a wrapper around the C function gtk_window_set_icon.
func (recv *Window) SetIcon(icon *gdkpixbuf.Pixbuf) {
	c_icon := (*C.GdkPixbuf)(icon.ToC())

	C.gtk_window_set_icon((*C.GtkWindow)(recv.native), c_icon)

	return
}

// SetIconList is a wrapper around the C function gtk_window_set_icon_list.
func (recv *Window) SetIconList(list *glib.List) {
	c_list := (*C.GList)(list.ToC())

	C.gtk_window_set_icon_list((*C.GtkWindow)(recv.native), c_list)

	return
}

// SetMnemonicModifier is a wrapper around the C function gtk_window_set_mnemonic_modifier.
func (recv *Window) SetMnemonicModifier(modifier gdk.ModifierType) {
	c_modifier := (C.GdkModifierType)(modifier)

	C.gtk_window_set_mnemonic_modifier((*C.GtkWindow)(recv.native), c_modifier)

	return
}

// SetModal is a wrapper around the C function gtk_window_set_modal.
func (recv *Window) SetModal(modal bool) {
	c_modal :=
		boolToGboolean(modal)

	C.gtk_window_set_modal((*C.GtkWindow)(recv.native), c_modal)

	return
}

// SetPosition is a wrapper around the C function gtk_window_set_position.
func (recv *Window) SetPosition(position WindowPosition) {
	c_position := (C.GtkWindowPosition)(position)

	C.gtk_window_set_position((*C.GtkWindow)(recv.native), c_position)

	return
}

// SetResizable is a wrapper around the C function gtk_window_set_resizable.
func (recv *Window) SetResizable(resizable bool) {
	c_resizable :=
		boolToGboolean(resizable)

	C.gtk_window_set_resizable((*C.GtkWindow)(recv.native), c_resizable)

	return
}

// SetRole is a wrapper around the C function gtk_window_set_role.
func (recv *Window) SetRole(role string) {
	c_role := C.CString(role)
	defer C.free(unsafe.Pointer(c_role))

	C.gtk_window_set_role((*C.GtkWindow)(recv.native), c_role)

	return
}

// SetTitle is a wrapper around the C function gtk_window_set_title.
func (recv *Window) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_window_set_title((*C.GtkWindow)(recv.native), c_title)

	return
}

// SetTransientFor is a wrapper around the C function gtk_window_set_transient_for.
func (recv *Window) SetTransientFor(parent *Window) {
	c_parent := (*C.GtkWindow)(parent.ToC())

	C.gtk_window_set_transient_for((*C.GtkWindow)(recv.native), c_parent)

	return
}

// SetTypeHint is a wrapper around the C function gtk_window_set_type_hint.
func (recv *Window) SetTypeHint(hint gdk.WindowTypeHint) {
	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gtk_window_set_type_hint((*C.GtkWindow)(recv.native), c_hint)

	return
}

// SetWmclass is a wrapper around the C function gtk_window_set_wmclass.
func (recv *Window) SetWmclass(wmclassName string, wmclassClass string) {
	c_wmclass_name := C.CString(wmclassName)
	defer C.free(unsafe.Pointer(c_wmclass_name))

	c_wmclass_class := C.CString(wmclassClass)
	defer C.free(unsafe.Pointer(c_wmclass_class))

	C.gtk_window_set_wmclass((*C.GtkWindow)(recv.native), c_wmclass_name, c_wmclass_class)

	return
}

// Stick is a wrapper around the C function gtk_window_stick.
func (recv *Window) Stick() {
	C.gtk_window_stick((*C.GtkWindow)(recv.native))

	return
}

// Unmaximize is a wrapper around the C function gtk_window_unmaximize.
func (recv *Window) Unmaximize() {
	C.gtk_window_unmaximize((*C.GtkWindow)(recv.native))

	return
}

// Unstick is a wrapper around the C function gtk_window_unstick.
func (recv *Window) Unstick() {
	C.gtk_window_unstick((*C.GtkWindow)(recv.native))

	return
}

// WindowAccessible is a wrapper around the C record GtkWindowAccessible.
type WindowAccessible struct {
	native *C.GtkWindowAccessible
	// parent : record
	// priv : record
}

func WindowAccessibleNewFromC(u unsafe.Pointer) *WindowAccessible {
	c := (*C.GtkWindowAccessible)(u)
	if c == nil {
		return nil
	}

	g := &WindowAccessible{native: c}

	return g
}

func (recv *WindowAccessible) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ContainerAccessible upcasts to *ContainerAccessible
func (recv *WindowAccessible) ContainerAccessible() *ContainerAccessible {
	return ContainerAccessibleNewFromC(unsafe.Pointer(recv.native))
}

// WidgetAccessible upcasts to *WidgetAccessible
func (recv *WindowAccessible) WidgetAccessible() *WidgetAccessible {
	return recv.ContainerAccessible().WidgetAccessible()
}

// Accessible upcasts to *Accessible
func (recv *WindowAccessible) Accessible() *Accessible {
	return recv.ContainerAccessible().Accessible()
}

// Object upcasts to *Object
func (recv *WindowAccessible) Object() *atk.Object {
	return recv.ContainerAccessible().Object()
}

// CastToWidget down casts any arbitary Object to WindowAccessible.
// Exercise care, as this is a potentially dangerous function if the Object is not a WindowAccessible.
func CastToWindowAccessible(object *gobject.Object) *WindowAccessible {
	return WindowAccessibleNewFromC(object.ToC())
}

// WindowGroup is a wrapper around the C record GtkWindowGroup.
type WindowGroup struct {
	native *C.GtkWindowGroup
	// parent_instance : record
	// priv : record
}

func WindowGroupNewFromC(u unsafe.Pointer) *WindowGroup {
	c := (*C.GtkWindowGroup)(u)
	if c == nil {
		return nil
	}

	g := &WindowGroup{native: c}

	return g
}

func (recv *WindowGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Object upcasts to *Object
func (recv *WindowGroup) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitary Object to WindowGroup.
// Exercise care, as this is a potentially dangerous function if the Object is not a WindowGroup.
func CastToWindowGroup(object *gobject.Object) *WindowGroup {
	return WindowGroupNewFromC(object.ToC())
}

// WindowGroupNew is a wrapper around the C function gtk_window_group_new.
func WindowGroupNew() *WindowGroup {
	retC := C.gtk_window_group_new()
	retGo := WindowGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddWindow is a wrapper around the C function gtk_window_group_add_window.
func (recv *WindowGroup) AddWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_window_group_add_window((*C.GtkWindowGroup)(recv.native), c_window)

	return
}

// RemoveWindow is a wrapper around the C function gtk_window_group_remove_window.
func (recv *WindowGroup) RemoveWindow(window *Window) {
	c_window := (*C.GtkWindow)(window.ToC())

	C.gtk_window_group_remove_window((*C.GtkWindowGroup)(recv.native), c_window)

	return
}
