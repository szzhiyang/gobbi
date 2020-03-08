// Code generated - DO NOT EDIT.
// +build gtk_2.2

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

static void c_gtk_binding_entry_add_signal(GtkBindingSet* binding_set, guint keyval, GdkModifierType modifiers, const gchar* signal_name, guint n_args) {
    return gtk_binding_entry_add_signal(binding_set, keyval, modifiers, signal_name, n_args, NULL);
}
*/
/*

static void c_gtk_container_add_with_properties(GtkContainer* container, GtkWidget* widget, const gchar* first_prop_name) {
    return gtk_container_add_with_properties(container, widget, first_prop_name, NULL);
}
*/
/*

static void c_gtk_container_child_get(GtkContainer* container, GtkWidget* child, const gchar* first_prop_name) {
    return gtk_container_child_get(container, child, first_prop_name, NULL);
}
*/
/*

static void c_gtk_container_child_get_valist(GtkContainer* container, GtkWidget* child, const gchar* first_property_name) {
    return gtk_container_child_get_valist(container, child, first_property_name, NULL);
}
*/
/*

static void c_gtk_container_child_set(GtkContainer* container, GtkWidget* child, const gchar* first_prop_name) {
    return gtk_container_child_set(container, child, first_prop_name, NULL);
}
*/
/*

static void c_gtk_container_child_set_valist(GtkContainer* container, GtkWidget* child, const gchar* first_property_name) {
    return gtk_container_child_set_valist(container, child, first_property_name, NULL);
}
*/
/*

static GtkWidget* c_gtk_dialog_new_with_buttons(const gchar* title, GtkWindow* parent, GtkDialogFlags flags, const gchar* first_button_text) {
    return gtk_dialog_new_with_buttons(title, parent, flags, first_button_text, NULL);
}
*/
/*

static void c_gtk_dialog_add_buttons(GtkDialog* dialog, const gchar* first_button_text) {
    return gtk_dialog_add_buttons(dialog, first_button_text, NULL);
}
*/
/*

static GtkWidget* c_gtk_info_bar_new_with_buttons(const gchar* first_button_text) {
    return gtk_info_bar_new_with_buttons(first_button_text, NULL);
}
*/
/*

static void c_gtk_list_store_set_valist(GtkListStore* list_store, GtkTreeIter* iter) {
    return gtk_list_store_set_valist(list_store, iter, NULL);
}
*/
/*

static GtkWidget* c_gtk_message_dialog_new(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, const gchar* message_format) {
    return gtk_message_dialog_new(parent, flags, type, buttons, message_format, NULL);
}
*/
/*

static GtkTextTag* c_gtk_text_buffer_create_tag(GtkTextBuffer* buffer, const gchar* tag_name, const gchar* first_property_name) {
    return gtk_text_buffer_create_tag(buffer, tag_name, first_property_name, NULL);
}
*/
/*

static void c_gtk_text_buffer_insert_with_tags(GtkTextBuffer* buffer, GtkTextIter* iter, const gchar* text, gint len, GtkTextTag* first_tag) {
    return gtk_text_buffer_insert_with_tags(buffer, iter, text, len, first_tag, NULL);
}
*/
/*

static void c_gtk_text_buffer_insert_with_tags_by_name(GtkTextBuffer* buffer, GtkTextIter* iter, const gchar* text, gint len, const gchar* first_tag_name) {
    return gtk_text_buffer_insert_with_tags_by_name(buffer, iter, text, len, first_tag_name, NULL);
}
*/
/*

static void c_gtk_tree_store_set_valist(GtkTreeStore* tree_store, GtkTreeIter* iter) {
    return gtk_tree_store_set_valist(tree_store, iter, NULL);
}
*/
/*

static gint c_gtk_tree_view_insert_column_with_attributes(GtkTreeView* tree_view, gint position, const gchar* title, GtkCellRenderer* cell) {
    return gtk_tree_view_insert_column_with_attributes(tree_view, position, title, cell, NULL);
}
*/
/*

static GtkTreeViewColumn* c_gtk_tree_view_column_new_with_attributes(const gchar* title, GtkCellRenderer* cell) {
    return gtk_tree_view_column_new_with_attributes(title, cell, NULL);
}
*/
/*

static void c_gtk_tree_view_column_set_attributes(GtkTreeViewColumn* tree_column, GtkCellRenderer* cell_renderer) {
    return gtk_tree_view_column_set_attributes(tree_column, cell_renderer, NULL);
}
*/
/*

static GtkWidget* c_gtk_widget_new(GType type, const gchar* first_property_name) {
    return gtk_widget_new(type, first_property_name, NULL);
}
*/
/*

static void c_gtk_widget_style_get(GtkWidget* widget, const gchar* first_property_name) {
    return gtk_widget_style_get(widget, first_property_name, NULL);
}
*/
/*

static void c_gtk_widget_style_get_valist(GtkWidget* widget, const gchar* first_property_name) {
    return gtk_widget_style_get_valist(widget, first_property_name, NULL);
}
*/
/*

static void c_gtk_tree_model_get(GtkTreeModel* tree_model, GtkTreeIter* iter) {
    return gtk_tree_model_get(tree_model, iter, NULL);
}
*/
/*

static void c_gtk_tree_model_get_valist(GtkTreeModel* tree_model, GtkTreeIter* iter) {
    return gtk_tree_model_get_valist(tree_model, iter, NULL);
}
*/
import "C"

func toCBool(b bool) C.gboolean {
	if b {
		return C.TRUE
	}
	return C.FALSE
}
func toGoBool(b C.gboolean) bool {
	return b == C.TRUE
}

type AboutDialogClass C.GtkAboutDialogClass
type AboutDialogPrivate C.GtkAboutDialogPrivate
type AccelGroupClass C.GtkAccelGroupClass
type AccelGroupEntry C.GtkAccelGroupEntry
type AccelGroupPrivate C.GtkAccelGroupPrivate
type AccelKey C.GtkAccelKey
type AccelLabelClass C.GtkAccelLabelClass
type AccelLabelPrivate C.GtkAccelLabelPrivate
type AccelMapClass C.GtkAccelMapClass
type AccessibleClass C.GtkAccessibleClass
type AccessiblePrivate C.GtkAccessiblePrivate
type ActionBarClass C.GtkActionBarClass
type ActionBarPrivate C.GtkActionBarPrivate
type ActionClass C.GtkActionClass
type ActionEntry C.GtkActionEntry
type ActionGroupClass C.GtkActionGroupClass
type ActionGroupPrivate C.GtkActionGroupPrivate
type ActionPrivate C.GtkActionPrivate
type ActionableInterface C.GtkActionableInterface
type AdjustmentClass C.GtkAdjustmentClass
type AdjustmentPrivate C.GtkAdjustmentPrivate
type AlignmentClass C.GtkAlignmentClass
type AlignmentPrivate C.GtkAlignmentPrivate
type AppChooserButtonClass C.GtkAppChooserButtonClass
type AppChooserButtonPrivate C.GtkAppChooserButtonPrivate
type AppChooserDialogClass C.GtkAppChooserDialogClass
type AppChooserDialogPrivate C.GtkAppChooserDialogPrivate
type AppChooserWidgetClass C.GtkAppChooserWidgetClass
type AppChooserWidgetPrivate C.GtkAppChooserWidgetPrivate
type ApplicationClass C.GtkApplicationClass
type ApplicationPrivate C.GtkApplicationPrivate
type ApplicationWindowClass C.GtkApplicationWindowClass
type ApplicationWindowPrivate C.GtkApplicationWindowPrivate
type ArrowAccessibleClass C.GtkArrowAccessibleClass
type ArrowAccessiblePrivate C.GtkArrowAccessiblePrivate
type ArrowClass C.GtkArrowClass
type ArrowPrivate C.GtkArrowPrivate
type AspectFrameClass C.GtkAspectFrameClass
type AspectFramePrivate C.GtkAspectFramePrivate
type AssistantClass C.GtkAssistantClass
type AssistantPrivate C.GtkAssistantPrivate
type BinClass C.GtkBinClass
type BinPrivate C.GtkBinPrivate
type BindingArg C.GtkBindingArg
type BindingEntry C.GtkBindingEntry
type BindingSet C.GtkBindingSet
type BindingSignal C.GtkBindingSignal
type BooleanCellAccessibleClass C.GtkBooleanCellAccessibleClass
type BooleanCellAccessiblePrivate C.GtkBooleanCellAccessiblePrivate
type Border C.GtkBorder
type BoxClass C.GtkBoxClass
type BoxPrivate C.GtkBoxPrivate
type BuildableIface C.GtkBuildableIface
type BuilderClass C.GtkBuilderClass
type BuilderPrivate C.GtkBuilderPrivate
type ButtonAccessibleClass C.GtkButtonAccessibleClass
type ButtonAccessiblePrivate C.GtkButtonAccessiblePrivate
type ButtonBoxClass C.GtkButtonBoxClass
type ButtonBoxPrivate C.GtkButtonBoxPrivate
type ButtonClass C.GtkButtonClass
type ButtonPrivate C.GtkButtonPrivate
type CalendarClass C.GtkCalendarClass
type CalendarPrivate C.GtkCalendarPrivate
type CellAccessibleClass C.GtkCellAccessibleClass
type CellAccessibleParentIface C.GtkCellAccessibleParentIface
type CellAccessiblePrivate C.GtkCellAccessiblePrivate
type CellAreaBoxClass C.GtkCellAreaBoxClass
type CellAreaBoxPrivate C.GtkCellAreaBoxPrivate
type CellAreaClass C.GtkCellAreaClass
type CellAreaContextClass C.GtkCellAreaContextClass
type CellAreaContextPrivate C.GtkCellAreaContextPrivate
type CellAreaPrivate C.GtkCellAreaPrivate
type CellEditableIface C.GtkCellEditableIface
type CellLayoutIface C.GtkCellLayoutIface
type CellRendererAccelClass C.GtkCellRendererAccelClass
type CellRendererAccelPrivate C.GtkCellRendererAccelPrivate
type CellRendererClass C.GtkCellRendererClass
type CellRendererClassPrivate C.GtkCellRendererClassPrivate
type CellRendererComboClass C.GtkCellRendererComboClass
type CellRendererComboPrivate C.GtkCellRendererComboPrivate
type CellRendererPixbufClass C.GtkCellRendererPixbufClass
type CellRendererPixbufPrivate C.GtkCellRendererPixbufPrivate
type CellRendererPrivate C.GtkCellRendererPrivate
type CellRendererProgressClass C.GtkCellRendererProgressClass
type CellRendererProgressPrivate C.GtkCellRendererProgressPrivate
type CellRendererSpinClass C.GtkCellRendererSpinClass
type CellRendererSpinPrivate C.GtkCellRendererSpinPrivate
type CellRendererSpinnerClass C.GtkCellRendererSpinnerClass
type CellRendererSpinnerPrivate C.GtkCellRendererSpinnerPrivate
type CellRendererTextClass C.GtkCellRendererTextClass
type CellRendererTextPrivate C.GtkCellRendererTextPrivate
type CellRendererToggleClass C.GtkCellRendererToggleClass
type CellRendererTogglePrivate C.GtkCellRendererTogglePrivate
type CellViewClass C.GtkCellViewClass
type CellViewPrivate C.GtkCellViewPrivate
type CheckButtonClass C.GtkCheckButtonClass
type CheckMenuItemAccessibleClass C.GtkCheckMenuItemAccessibleClass
type CheckMenuItemAccessiblePrivate C.GtkCheckMenuItemAccessiblePrivate
type CheckMenuItemClass C.GtkCheckMenuItemClass
type CheckMenuItemPrivate C.GtkCheckMenuItemPrivate
type ColorButtonClass C.GtkColorButtonClass
type ColorButtonPrivate C.GtkColorButtonPrivate
type ColorChooserDialogClass C.GtkColorChooserDialogClass
type ColorChooserDialogPrivate C.GtkColorChooserDialogPrivate
type ColorChooserInterface C.GtkColorChooserInterface
type ColorChooserWidgetClass C.GtkColorChooserWidgetClass
type ColorChooserWidgetPrivate C.GtkColorChooserWidgetPrivate
type ColorSelectionClass C.GtkColorSelectionClass
type ColorSelectionDialogClass C.GtkColorSelectionDialogClass
type ColorSelectionDialogPrivate C.GtkColorSelectionDialogPrivate
type ColorSelectionPrivate C.GtkColorSelectionPrivate
type ComboBoxAccessibleClass C.GtkComboBoxAccessibleClass
type ComboBoxAccessiblePrivate C.GtkComboBoxAccessiblePrivate
type ComboBoxClass C.GtkComboBoxClass
type ComboBoxPrivate C.GtkComboBoxPrivate
type ComboBoxTextClass C.GtkComboBoxTextClass
type ComboBoxTextPrivate C.GtkComboBoxTextPrivate
type ContainerAccessibleClass C.GtkContainerAccessibleClass
type ContainerAccessiblePrivate C.GtkContainerAccessiblePrivate
type ContainerCellAccessibleClass C.GtkContainerCellAccessibleClass
type ContainerCellAccessiblePrivate C.GtkContainerCellAccessiblePrivate
type ContainerClass C.GtkContainerClass
type ContainerPrivate C.GtkContainerPrivate
type CssProviderClass C.GtkCssProviderClass
type CssProviderPrivate C.GtkCssProviderPrivate
type DialogClass C.GtkDialogClass
type DialogPrivate C.GtkDialogPrivate
type DrawingAreaClass C.GtkDrawingAreaClass
type EditableInterface C.GtkEditableInterface
type EntryAccessibleClass C.GtkEntryAccessibleClass
type EntryAccessiblePrivate C.GtkEntryAccessiblePrivate
type EntryBufferClass C.GtkEntryBufferClass
type EntryBufferPrivate C.GtkEntryBufferPrivate
type EntryClass C.GtkEntryClass
type EntryCompletionClass C.GtkEntryCompletionClass
type EntryCompletionPrivate C.GtkEntryCompletionPrivate
type EntryPrivate C.GtkEntryPrivate
type EventBoxClass C.GtkEventBoxClass
type EventBoxPrivate C.GtkEventBoxPrivate
type EventControllerClass C.GtkEventControllerClass

// UNSUPPORTED : EventControllerMotionClass : blacklisted
// UNSUPPORTED : EventControllerScrollClass : blacklisted
type ExpanderAccessibleClass C.GtkExpanderAccessibleClass
type ExpanderAccessiblePrivate C.GtkExpanderAccessiblePrivate
type ExpanderClass C.GtkExpanderClass
type ExpanderPrivate C.GtkExpanderPrivate
type FileChooserButtonClass C.GtkFileChooserButtonClass
type FileChooserButtonPrivate C.GtkFileChooserButtonPrivate
type FileChooserDialogClass C.GtkFileChooserDialogClass
type FileChooserDialogPrivate C.GtkFileChooserDialogPrivate
type FileChooserNativeClass C.GtkFileChooserNativeClass
type FileChooserWidgetClass C.GtkFileChooserWidgetClass
type FileChooserWidgetPrivate C.GtkFileChooserWidgetPrivate
type FileFilterInfo C.GtkFileFilterInfo
type FixedChild C.GtkFixedChild
type FixedClass C.GtkFixedClass
type FixedPrivate C.GtkFixedPrivate
type FlowBoxAccessibleClass C.GtkFlowBoxAccessibleClass
type FlowBoxAccessiblePrivate C.GtkFlowBoxAccessiblePrivate
type FlowBoxChildAccessibleClass C.GtkFlowBoxChildAccessibleClass
type FlowBoxChildClass C.GtkFlowBoxChildClass
type FlowBoxClass C.GtkFlowBoxClass
type FontButtonClass C.GtkFontButtonClass
type FontButtonPrivate C.GtkFontButtonPrivate
type FontChooserDialogClass C.GtkFontChooserDialogClass
type FontChooserDialogPrivate C.GtkFontChooserDialogPrivate
type FontChooserIface C.GtkFontChooserIface
type FontChooserWidgetClass C.GtkFontChooserWidgetClass
type FontChooserWidgetPrivate C.GtkFontChooserWidgetPrivate
type FontSelectionClass C.GtkFontSelectionClass
type FontSelectionDialogClass C.GtkFontSelectionDialogClass
type FontSelectionDialogPrivate C.GtkFontSelectionDialogPrivate
type FontSelectionPrivate C.GtkFontSelectionPrivate
type FrameAccessibleClass C.GtkFrameAccessibleClass
type FrameAccessiblePrivate C.GtkFrameAccessiblePrivate
type FrameClass C.GtkFrameClass
type FramePrivate C.GtkFramePrivate
type GestureClass C.GtkGestureClass
type GestureDragClass C.GtkGestureDragClass
type GestureLongPressClass C.GtkGestureLongPressClass
type GestureMultiPressClass C.GtkGestureMultiPressClass
type GesturePanClass C.GtkGesturePanClass
type GestureRotateClass C.GtkGestureRotateClass
type GestureSingleClass C.GtkGestureSingleClass

// UNSUPPORTED : GestureStylusClass : blacklisted
type GestureSwipeClass C.GtkGestureSwipeClass
type GestureZoomClass C.GtkGestureZoomClass
type Gradient C.GtkGradient
type GridClass C.GtkGridClass
type GridPrivate C.GtkGridPrivate
type HBoxClass C.GtkHBoxClass
type HButtonBoxClass C.GtkHButtonBoxClass
type HPanedClass C.GtkHPanedClass
type HSVClass C.GtkHSVClass
type HSVPrivate C.GtkHSVPrivate
type HScaleClass C.GtkHScaleClass
type HScrollbarClass C.GtkHScrollbarClass
type HSeparatorClass C.GtkHSeparatorClass
type HandleBoxClass C.GtkHandleBoxClass
type HandleBoxPrivate C.GtkHandleBoxPrivate

// UNSUPPORTED : HeaderBarAccessibleClass : blacklisted
// UNSUPPORTED : HeaderBarAccessiblePrivate : blacklisted
type HeaderBarClass C.GtkHeaderBarClass
type HeaderBarPrivate C.GtkHeaderBarPrivate
type IMContextClass C.GtkIMContextClass
type IMContextInfo C.GtkIMContextInfo
type IMContextSimpleClass C.GtkIMContextSimpleClass
type IMContextSimplePrivate C.GtkIMContextSimplePrivate
type IMMulticontextClass C.GtkIMMulticontextClass
type IMMulticontextPrivate C.GtkIMMulticontextPrivate
type IconFactoryClass C.GtkIconFactoryClass
type IconFactoryPrivate C.GtkIconFactoryPrivate
type IconInfoClass C.GtkIconInfoClass
type IconSet C.GtkIconSet
type IconSource C.GtkIconSource
type IconThemeClass C.GtkIconThemeClass
type IconThemePrivate C.GtkIconThemePrivate
type IconViewAccessibleClass C.GtkIconViewAccessibleClass
type IconViewAccessiblePrivate C.GtkIconViewAccessiblePrivate
type IconViewClass C.GtkIconViewClass
type IconViewPrivate C.GtkIconViewPrivate
type ImageAccessibleClass C.GtkImageAccessibleClass
type ImageAccessiblePrivate C.GtkImageAccessiblePrivate
type ImageCellAccessibleClass C.GtkImageCellAccessibleClass
type ImageCellAccessiblePrivate C.GtkImageCellAccessiblePrivate
type ImageClass C.GtkImageClass
type ImageMenuItemClass C.GtkImageMenuItemClass
type ImageMenuItemPrivate C.GtkImageMenuItemPrivate
type ImagePrivate C.GtkImagePrivate
type InfoBarClass C.GtkInfoBarClass
type InfoBarPrivate C.GtkInfoBarPrivate
type InvisibleClass C.GtkInvisibleClass
type InvisiblePrivate C.GtkInvisiblePrivate
type LabelAccessibleClass C.GtkLabelAccessibleClass
type LabelAccessiblePrivate C.GtkLabelAccessiblePrivate
type LabelClass C.GtkLabelClass
type LabelPrivate C.GtkLabelPrivate
type LabelSelectionInfo C.GtkLabelSelectionInfo
type LayoutClass C.GtkLayoutClass
type LayoutPrivate C.GtkLayoutPrivate
type LevelBarAccessibleClass C.GtkLevelBarAccessibleClass
type LevelBarAccessiblePrivate C.GtkLevelBarAccessiblePrivate
type LevelBarClass C.GtkLevelBarClass
type LevelBarPrivate C.GtkLevelBarPrivate
type LinkButtonAccessibleClass C.GtkLinkButtonAccessibleClass
type LinkButtonAccessiblePrivate C.GtkLinkButtonAccessiblePrivate
type LinkButtonClass C.GtkLinkButtonClass
type LinkButtonPrivate C.GtkLinkButtonPrivate
type ListBoxAccessibleClass C.GtkListBoxAccessibleClass
type ListBoxAccessiblePrivate C.GtkListBoxAccessiblePrivate
type ListBoxClass C.GtkListBoxClass
type ListBoxRowAccessibleClass C.GtkListBoxRowAccessibleClass
type ListBoxRowClass C.GtkListBoxRowClass
type ListStoreClass C.GtkListStoreClass
type ListStorePrivate C.GtkListStorePrivate
type LockButtonAccessibleClass C.GtkLockButtonAccessibleClass
type LockButtonAccessiblePrivate C.GtkLockButtonAccessiblePrivate
type LockButtonClass C.GtkLockButtonClass
type LockButtonPrivate C.GtkLockButtonPrivate
type MenuAccessibleClass C.GtkMenuAccessibleClass
type MenuAccessiblePrivate C.GtkMenuAccessiblePrivate
type MenuBarClass C.GtkMenuBarClass
type MenuBarPrivate C.GtkMenuBarPrivate
type MenuButtonAccessibleClass C.GtkMenuButtonAccessibleClass
type MenuButtonAccessiblePrivate C.GtkMenuButtonAccessiblePrivate
type MenuButtonClass C.GtkMenuButtonClass
type MenuButtonPrivate C.GtkMenuButtonPrivate
type MenuClass C.GtkMenuClass
type MenuItemAccessibleClass C.GtkMenuItemAccessibleClass
type MenuItemAccessiblePrivate C.GtkMenuItemAccessiblePrivate
type MenuItemClass C.GtkMenuItemClass
type MenuItemPrivate C.GtkMenuItemPrivate
type MenuPrivate C.GtkMenuPrivate
type MenuShellAccessibleClass C.GtkMenuShellAccessibleClass
type MenuShellAccessiblePrivate C.GtkMenuShellAccessiblePrivate
type MenuShellClass C.GtkMenuShellClass
type MenuShellPrivate C.GtkMenuShellPrivate
type MenuToolButtonClass C.GtkMenuToolButtonClass
type MenuToolButtonPrivate C.GtkMenuToolButtonPrivate
type MessageDialogClass C.GtkMessageDialogClass
type MessageDialogPrivate C.GtkMessageDialogPrivate
type MiscClass C.GtkMiscClass
type MiscPrivate C.GtkMiscPrivate
type MountOperationClass C.GtkMountOperationClass
type MountOperationPrivate C.GtkMountOperationPrivate
type NativeDialogClass C.GtkNativeDialogClass
type NotebookAccessibleClass C.GtkNotebookAccessibleClass
type NotebookAccessiblePrivate C.GtkNotebookAccessiblePrivate
type NotebookClass C.GtkNotebookClass
type NotebookPageAccessibleClass C.GtkNotebookPageAccessibleClass
type NotebookPageAccessiblePrivate C.GtkNotebookPageAccessiblePrivate
type NotebookPrivate C.GtkNotebookPrivate
type NumerableIconClass C.GtkNumerableIconClass
type NumerableIconPrivate C.GtkNumerableIconPrivate
type OffscreenWindowClass C.GtkOffscreenWindowClass
type OrientableIface C.GtkOrientableIface
type OverlayClass C.GtkOverlayClass
type OverlayPrivate C.GtkOverlayPrivate
type PadActionEntry C.GtkPadActionEntry
type PadControllerClass C.GtkPadControllerClass
type PageRange C.GtkPageRange
type PanedAccessibleClass C.GtkPanedAccessibleClass
type PanedAccessiblePrivate C.GtkPanedAccessiblePrivate
type PanedClass C.GtkPanedClass
type PanedPrivate C.GtkPanedPrivate
type PaperSize C.GtkPaperSize
type PlacesSidebarClass C.GtkPlacesSidebarClass
type PlugClass C.GtkPlugClass
type PlugPrivate C.GtkPlugPrivate
type PopoverAccessibleClass C.GtkPopoverAccessibleClass
type PopoverClass C.GtkPopoverClass
type PopoverMenuClass C.GtkPopoverMenuClass
type PopoverPrivate C.GtkPopoverPrivate
type PrintOperationClass C.GtkPrintOperationClass
type PrintOperationPreviewIface C.GtkPrintOperationPreviewIface
type PrintOperationPrivate C.GtkPrintOperationPrivate
type ProgressBarAccessibleClass C.GtkProgressBarAccessibleClass
type ProgressBarAccessiblePrivate C.GtkProgressBarAccessiblePrivate
type ProgressBarClass C.GtkProgressBarClass
type ProgressBarPrivate C.GtkProgressBarPrivate
type RadioActionClass C.GtkRadioActionClass
type RadioActionEntry C.GtkRadioActionEntry
type RadioActionPrivate C.GtkRadioActionPrivate
type RadioButtonAccessibleClass C.GtkRadioButtonAccessibleClass
type RadioButtonAccessiblePrivate C.GtkRadioButtonAccessiblePrivate
type RadioButtonClass C.GtkRadioButtonClass
type RadioButtonPrivate C.GtkRadioButtonPrivate
type RadioMenuItemAccessibleClass C.GtkRadioMenuItemAccessibleClass
type RadioMenuItemAccessiblePrivate C.GtkRadioMenuItemAccessiblePrivate
type RadioMenuItemClass C.GtkRadioMenuItemClass
type RadioMenuItemPrivate C.GtkRadioMenuItemPrivate
type RadioToolButtonClass C.GtkRadioToolButtonClass
type RangeAccessibleClass C.GtkRangeAccessibleClass
type RangeAccessiblePrivate C.GtkRangeAccessiblePrivate
type RangeClass C.GtkRangeClass
type RangePrivate C.GtkRangePrivate
type RcContext C.GtkRcContext
type RcProperty C.GtkRcProperty
type RcStyleClass C.GtkRcStyleClass
type RecentActionClass C.GtkRecentActionClass
type RecentActionPrivate C.GtkRecentActionPrivate
type RecentChooserDialogClass C.GtkRecentChooserDialogClass
type RecentChooserDialogPrivate C.GtkRecentChooserDialogPrivate
type RecentChooserIface C.GtkRecentChooserIface
type RecentChooserMenuClass C.GtkRecentChooserMenuClass
type RecentChooserMenuPrivate C.GtkRecentChooserMenuPrivate
type RecentChooserWidgetClass C.GtkRecentChooserWidgetClass
type RecentChooserWidgetPrivate C.GtkRecentChooserWidgetPrivate
type RecentData C.GtkRecentData
type RecentFilterInfo C.GtkRecentFilterInfo
type RecentManagerPrivate C.GtkRecentManagerPrivate
type RendererCellAccessibleClass C.GtkRendererCellAccessibleClass
type RendererCellAccessiblePrivate C.GtkRendererCellAccessiblePrivate
type RequestedSize C.GtkRequestedSize
type Requisition C.GtkRequisition
type RevealerClass C.GtkRevealerClass
type ScaleAccessibleClass C.GtkScaleAccessibleClass
type ScaleAccessiblePrivate C.GtkScaleAccessiblePrivate
type ScaleButtonAccessibleClass C.GtkScaleButtonAccessibleClass
type ScaleButtonAccessiblePrivate C.GtkScaleButtonAccessiblePrivate
type ScaleButtonClass C.GtkScaleButtonClass
type ScaleButtonPrivate C.GtkScaleButtonPrivate
type ScaleClass C.GtkScaleClass
type ScalePrivate C.GtkScalePrivate
type ScrollableInterface C.GtkScrollableInterface
type ScrollbarClass C.GtkScrollbarClass
type ScrolledWindowAccessibleClass C.GtkScrolledWindowAccessibleClass
type ScrolledWindowAccessiblePrivate C.GtkScrolledWindowAccessiblePrivate
type ScrolledWindowClass C.GtkScrolledWindowClass
type ScrolledWindowPrivate C.GtkScrolledWindowPrivate
type SearchBarClass C.GtkSearchBarClass
type SearchEntryClass C.GtkSearchEntryClass
type SelectionData C.GtkSelectionData
type SeparatorClass C.GtkSeparatorClass
type SeparatorMenuItemClass C.GtkSeparatorMenuItemClass
type SeparatorPrivate C.GtkSeparatorPrivate
type SeparatorToolItemClass C.GtkSeparatorToolItemClass
type SeparatorToolItemPrivate C.GtkSeparatorToolItemPrivate
type SettingsClass C.GtkSettingsClass
type SettingsPrivate C.GtkSettingsPrivate
type SettingsValue C.GtkSettingsValue
type ShortcutLabelClass C.GtkShortcutLabelClass
type ShortcutsGroupClass C.GtkShortcutsGroupClass
type ShortcutsSectionClass C.GtkShortcutsSectionClass
type ShortcutsShortcutClass C.GtkShortcutsShortcutClass
type ShortcutsWindowClass C.GtkShortcutsWindowClass
type SizeGroupClass C.GtkSizeGroupClass
type SizeGroupPrivate C.GtkSizeGroupPrivate
type SocketClass C.GtkSocketClass
type SocketPrivate C.GtkSocketPrivate
type SpinButtonAccessibleClass C.GtkSpinButtonAccessibleClass
type SpinButtonAccessiblePrivate C.GtkSpinButtonAccessiblePrivate
type SpinButtonClass C.GtkSpinButtonClass
type SpinButtonPrivate C.GtkSpinButtonPrivate
type SpinnerAccessibleClass C.GtkSpinnerAccessibleClass
type SpinnerAccessiblePrivate C.GtkSpinnerAccessiblePrivate
type SpinnerClass C.GtkSpinnerClass
type SpinnerPrivate C.GtkSpinnerPrivate

// UNSUPPORTED : StackAccessibleClass : blacklisted
type StackClass C.GtkStackClass
type StackSidebarClass C.GtkStackSidebarClass
type StackSidebarPrivate C.GtkStackSidebarPrivate
type StackSwitcherClass C.GtkStackSwitcherClass
type StatusIconClass C.GtkStatusIconClass
type StatusIconPrivate C.GtkStatusIconPrivate
type StatusbarAccessibleClass C.GtkStatusbarAccessibleClass
type StatusbarAccessiblePrivate C.GtkStatusbarAccessiblePrivate
type StatusbarClass C.GtkStatusbarClass
type StatusbarPrivate C.GtkStatusbarPrivate
type StockItem C.GtkStockItem
type StyleClass C.GtkStyleClass
type StyleContextClass C.GtkStyleContextClass
type StyleContextPrivate C.GtkStyleContextPrivate
type StylePropertiesClass C.GtkStylePropertiesClass
type StylePropertiesPrivate C.GtkStylePropertiesPrivate
type StyleProviderIface C.GtkStyleProviderIface
type SwitchAccessibleClass C.GtkSwitchAccessibleClass
type SwitchAccessiblePrivate C.GtkSwitchAccessiblePrivate
type SwitchClass C.GtkSwitchClass
type SwitchPrivate C.GtkSwitchPrivate
type SymbolicColor C.GtkSymbolicColor
type TableChild C.GtkTableChild
type TableClass C.GtkTableClass
type TablePrivate C.GtkTablePrivate
type TableRowCol C.GtkTableRowCol
type TargetEntry C.GtkTargetEntry
type TargetList C.GtkTargetList
type TargetPair C.GtkTargetPair
type TearoffMenuItemClass C.GtkTearoffMenuItemClass
type TearoffMenuItemPrivate C.GtkTearoffMenuItemPrivate
type TextAppearance C.GtkTextAppearance
type TextAttributes C.GtkTextAttributes
type TextBTree C.GtkTextBTree
type TextBufferClass C.GtkTextBufferClass
type TextBufferPrivate C.GtkTextBufferPrivate
type TextCellAccessibleClass C.GtkTextCellAccessibleClass
type TextCellAccessiblePrivate C.GtkTextCellAccessiblePrivate
type TextChildAnchorClass C.GtkTextChildAnchorClass
type TextIter C.GtkTextIter
type TextMarkClass C.GtkTextMarkClass
type TextTagClass C.GtkTextTagClass
type TextTagPrivate C.GtkTextTagPrivate
type TextTagTableClass C.GtkTextTagTableClass
type TextTagTablePrivate C.GtkTextTagTablePrivate
type TextViewAccessibleClass C.GtkTextViewAccessibleClass
type TextViewAccessiblePrivate C.GtkTextViewAccessiblePrivate
type TextViewClass C.GtkTextViewClass
type TextViewPrivate C.GtkTextViewPrivate
type ThemeEngine C.GtkThemeEngine
type ThemingEngineClass C.GtkThemingEngineClass
type ThemingEnginePrivate C.GtkThemingEnginePrivate
type ToggleActionClass C.GtkToggleActionClass
type ToggleActionEntry C.GtkToggleActionEntry
type ToggleActionPrivate C.GtkToggleActionPrivate
type ToggleButtonAccessibleClass C.GtkToggleButtonAccessibleClass
type ToggleButtonAccessiblePrivate C.GtkToggleButtonAccessiblePrivate
type ToggleButtonClass C.GtkToggleButtonClass
type ToggleButtonPrivate C.GtkToggleButtonPrivate
type ToggleToolButtonClass C.GtkToggleToolButtonClass
type ToggleToolButtonPrivate C.GtkToggleToolButtonPrivate
type ToolButtonClass C.GtkToolButtonClass
type ToolButtonPrivate C.GtkToolButtonPrivate
type ToolItemClass C.GtkToolItemClass
type ToolItemGroupClass C.GtkToolItemGroupClass
type ToolItemGroupPrivate C.GtkToolItemGroupPrivate
type ToolItemPrivate C.GtkToolItemPrivate
type ToolPaletteClass C.GtkToolPaletteClass
type ToolPalettePrivate C.GtkToolPalettePrivate
type ToolShellIface C.GtkToolShellIface
type ToolbarClass C.GtkToolbarClass
type ToolbarPrivate C.GtkToolbarPrivate
type ToplevelAccessibleClass C.GtkToplevelAccessibleClass
type ToplevelAccessiblePrivate C.GtkToplevelAccessiblePrivate
type TreeDragDestIface C.GtkTreeDragDestIface
type TreeDragSourceIface C.GtkTreeDragSourceIface
type TreeIter C.GtkTreeIter
type TreeModelFilterClass C.GtkTreeModelFilterClass
type TreeModelFilterPrivate C.GtkTreeModelFilterPrivate
type TreeModelIface C.GtkTreeModelIface
type TreeModelSortClass C.GtkTreeModelSortClass
type TreeModelSortPrivate C.GtkTreeModelSortPrivate
type TreePath C.GtkTreePath
type TreeRowReference C.GtkTreeRowReference
type TreeSelectionClass C.GtkTreeSelectionClass
type TreeSelectionPrivate C.GtkTreeSelectionPrivate
type TreeSortableIface C.GtkTreeSortableIface
type TreeStoreClass C.GtkTreeStoreClass
type TreeStorePrivate C.GtkTreeStorePrivate
type TreeViewAccessibleClass C.GtkTreeViewAccessibleClass
type TreeViewAccessiblePrivate C.GtkTreeViewAccessiblePrivate
type TreeViewClass C.GtkTreeViewClass
type TreeViewColumnClass C.GtkTreeViewColumnClass
type TreeViewColumnPrivate C.GtkTreeViewColumnPrivate
type TreeViewPrivate C.GtkTreeViewPrivate
type UIManagerClass C.GtkUIManagerClass
type UIManagerPrivate C.GtkUIManagerPrivate
type VBoxClass C.GtkVBoxClass
type VButtonBoxClass C.GtkVButtonBoxClass
type VPanedClass C.GtkVPanedClass
type VScaleClass C.GtkVScaleClass
type VScrollbarClass C.GtkVScrollbarClass
type VSeparatorClass C.GtkVSeparatorClass
type ViewportClass C.GtkViewportClass
type ViewportPrivate C.GtkViewportPrivate
type VolumeButtonClass C.GtkVolumeButtonClass
type WidgetAccessibleClass C.GtkWidgetAccessibleClass
type WidgetAccessiblePrivate C.GtkWidgetAccessiblePrivate
type WidgetClass C.GtkWidgetClass
type WidgetClassPrivate C.GtkWidgetClassPrivate
type WidgetPath C.GtkWidgetPath
type WidgetPrivate C.GtkWidgetPrivate
type WindowAccessibleClass C.GtkWindowAccessibleClass
type WindowAccessiblePrivate C.GtkWindowAccessiblePrivate
type WindowClass C.GtkWindowClass
type WindowGeometryInfo C.GtkWindowGeometryInfo
type WindowGroupClass C.GtkWindowGroupClass
type WindowGroupPrivate C.GtkWindowGroupPrivate
type WindowPrivate C.GtkWindowPrivate

func Fn_gtk_binding_entry_add_signal(bindingSet unsafe.Pointer, keyval uint, modifiers int, signalName string, nArgs uint) {
	c_bindingSet := (*C.GtkBindingSet)(unsafe.Pointer(bindingSet))

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_signalName := (*C.gchar)(C.CString(signalName))
	defer C.free(unsafe.Pointer(c_signalName))

	c_nArgs := (C.guint)(nArgs)

	C.c_gtk_binding_entry_add_signal(c_bindingSet, c_keyval, c_modifiers, c_signalName, c_nArgs)
}

func Fn_gtk_binding_entry_add_signall(bindingSet unsafe.Pointer, keyval uint, modifiers int, signalName string, bindingArgs unsafe.Pointer) {
	c_bindingSet := (*C.GtkBindingSet)(unsafe.Pointer(bindingSet))

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_signalName := (*C.gchar)(C.CString(signalName))
	defer C.free(unsafe.Pointer(c_signalName))

	c_bindingArgs := (*C.GSList)(unsafe.Pointer(bindingArgs))

	C.gtk_binding_entry_add_signall(c_bindingSet, c_keyval, c_modifiers, c_signalName, c_bindingArgs)
}

func Fn_gtk_binding_entry_remove(bindingSet unsafe.Pointer, keyval uint, modifiers int) {
	c_bindingSet := (*C.GtkBindingSet)(unsafe.Pointer(bindingSet))

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	C.gtk_binding_entry_remove(c_bindingSet, c_keyval, c_modifiers)
}

func Fn_gtk_binding_set_activate(bindingSet unsafe.Pointer, keyval uint, modifiers int, object unsafe.Pointer) bool {
	c_bindingSet := (*C.GtkBindingSet)(unsafe.Pointer(bindingSet))

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	c_object := (*C.GObject)(unsafe.Pointer(object))

	ret := C.gtk_binding_set_activate(c_bindingSet, c_keyval, c_modifiers, c_object)

	return toGoBool(ret)
}

func Fn_gtk_binding_set_add_path(bindingSet unsafe.Pointer, pathType int, pathPattern string, priority int) {
	c_bindingSet := (*C.GtkBindingSet)(unsafe.Pointer(bindingSet))

	c_pathType := (C.GtkPathType)(pathType)

	c_pathPattern := (*C.gchar)(C.CString(pathPattern))
	defer C.free(unsafe.Pointer(c_pathPattern))

	c_priority := (C.GtkPathPriorityType)(priority)

	C.gtk_binding_set_add_path(c_bindingSet, c_pathType, c_pathPattern, c_priority)
}

func Fn_gtk_binding_set_by_class(objectClass unsafe.Pointer) unsafe.Pointer {
	c_objectClass := (C.gpointer)(objectClass)

	ret := C.gtk_binding_set_by_class(c_objectClass)

	return unsafe.Pointer(ret)
}

func Fn_gtk_binding_set_find(setName string) unsafe.Pointer {
	c_setName := (*C.gchar)(C.CString(setName))
	defer C.free(unsafe.Pointer(c_setName))

	ret := C.gtk_binding_set_find(c_setName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_binding_set_new(setName string) unsafe.Pointer {
	c_setName := (*C.gchar)(C.CString(setName))
	defer C.free(unsafe.Pointer(c_setName))

	ret := C.gtk_binding_set_new(c_setName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_border_copy(border unsafe.Pointer) unsafe.Pointer {
	c_border := (*C.GtkBorder)(unsafe.Pointer(border))

	ret := C.gtk_border_copy(c_border)

	return unsafe.Pointer(ret)
}

func Fn_gtk_border_free(border unsafe.Pointer) {
	c_border := (*C.GtkBorder)(unsafe.Pointer(border))

	C.gtk_border_free(c_border)
}

func Fn_gtk_cell_renderer_class_set_accessible_type(rendererClass unsafe.Pointer, type_ uint64) {
	c_rendererClass := (*C.GtkCellRendererClass)(unsafe.Pointer(rendererClass))

	c_type_ := (C.GType)(type_)

	C.gtk_cell_renderer_class_set_accessible_type(c_rendererClass, c_type_)
}

func Fn_gtk_container_class_find_child_property(cclass unsafe.Pointer, propertyName string) unsafe.Pointer {
	c_cclass := (*C.GObjectClass)(unsafe.Pointer(cclass))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	ret := C.gtk_container_class_find_child_property(c_cclass, c_propertyName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_class_handle_border_width(klass unsafe.Pointer) {
	c_klass := (*C.GtkContainerClass)(unsafe.Pointer(klass))

	C.gtk_container_class_handle_border_width(c_klass)
}

func Fn_gtk_container_class_install_child_property(cclass unsafe.Pointer, propertyId uint, pspec unsafe.Pointer) {
	c_cclass := (*C.GtkContainerClass)(unsafe.Pointer(cclass))

	c_propertyId := (C.guint)(propertyId)

	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	C.gtk_container_class_install_child_property(c_cclass, c_propertyId, c_pspec)
}

func Fn_gtk_container_class_list_child_properties(cclass unsafe.Pointer, nProperties *uint) []unsafe.Pointer {
	c_cclass := (*C.GObjectClass)(unsafe.Pointer(cclass))

	c_nProperties := (*C.guint)(unsafe.Pointer(nProperties))

	ret := C.gtk_container_class_list_child_properties(c_cclass, c_nProperties)

	retLen := int(*c_nProperties)
	retGo := make([]unsafe.Pointer, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_gradient_resolve : parameter 'resolved_gradient' is non array with indirect count > 1

func Fn_gtk_gradient_resolve_for_context(gradient unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	c_gradient := (*C.GtkGradient)(unsafe.Pointer(gradient))

	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	ret := C.gtk_gradient_resolve_for_context(c_gradient, c_context)

	return unsafe.Pointer(ret)
}

func Fn_gtk_gradient_to_string(gradient unsafe.Pointer) string {
	c_gradient := (*C.GtkGradient)(unsafe.Pointer(gradient))

	ret := C.gtk_gradient_to_string(c_gradient)

	return C.GoString(ret)
}

func Fn_gtk_icon_set_new() unsafe.Pointer {
	ret := C.gtk_icon_set_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_new_from_pixbuf(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gtk_icon_set_new_from_pixbuf(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_add_source(iconSet unsafe.Pointer, source unsafe.Pointer) {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	C.gtk_icon_set_add_source(c_iconSet, c_source)
}

func Fn_gtk_icon_set_copy(iconSet unsafe.Pointer) unsafe.Pointer {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	ret := C.gtk_icon_set_copy(c_iconSet)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_get_sizes(iconSet unsafe.Pointer, sizes *[]int, nSizes *int) {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	var c_sizesArrayPointer (*C.GtkIconSize)
	c_sizes := &c_sizesArrayPointer

	c_nSizes := (*C.gint)(unsafe.Pointer(nSizes))

	C.gtk_icon_set_get_sizes(c_iconSet, c_sizes, c_nSizes)

	sizesOutLen := int(*c_nSizes)
	sizesOut := make([]int, sizesOutLen, sizesOutLen)
	if sizesOutLen > 0 {
		sizesOut = (*[1 << 30](int))(unsafe.Pointer(c_sizesArrayPointer))[:sizesOutLen:sizesOutLen]
	}
	*sizes = sizesOut
}

func Fn_gtk_icon_set_ref(iconSet unsafe.Pointer) unsafe.Pointer {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	ret := C.gtk_icon_set_ref(c_iconSet)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_render_icon(iconSet unsafe.Pointer, style unsafe.Pointer, direction int, state int, size int, widget unsafe.Pointer, detail *string) unsafe.Pointer {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_direction := (C.GtkTextDirection)(direction)

	c_state := (C.GtkStateType)(state)

	c_size := (C.GtkIconSize)(size)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	ret := C.gtk_icon_set_render_icon(c_iconSet, c_style, c_direction, c_state, c_size, c_widget, c_detail)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_set_unref(iconSet unsafe.Pointer) {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	C.gtk_icon_set_unref(c_iconSet)
}

func Fn_gtk_icon_source_new() unsafe.Pointer {
	ret := C.gtk_icon_source_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_copy(source unsafe.Pointer) unsafe.Pointer {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_copy(c_source)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_free(source unsafe.Pointer) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	C.gtk_icon_source_free(c_source)
}

func Fn_gtk_icon_source_get_direction(source unsafe.Pointer) int {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_direction(c_source)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_direction_wildcarded(source unsafe.Pointer) bool {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_direction_wildcarded(c_source)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_get_filename(source unsafe.Pointer) string {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_filename(c_source)

	return C.GoString(ret)
}

func Fn_gtk_icon_source_get_icon_name(source unsafe.Pointer) string {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_icon_name(c_source)

	return C.GoString(ret)
}

func Fn_gtk_icon_source_get_pixbuf(source unsafe.Pointer) unsafe.Pointer {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_pixbuf(c_source)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_source_get_size(source unsafe.Pointer) int {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_size(c_source)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_size_wildcarded(source unsafe.Pointer) bool {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_size_wildcarded(c_source)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_get_state(source unsafe.Pointer) int {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_state(c_source)

	return (int)(ret)
}

func Fn_gtk_icon_source_get_state_wildcarded(source unsafe.Pointer) bool {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	ret := C.gtk_icon_source_get_state_wildcarded(c_source)

	return toGoBool(ret)
}

func Fn_gtk_icon_source_set_direction(source unsafe.Pointer, direction int) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_direction := (C.GtkTextDirection)(direction)

	C.gtk_icon_source_set_direction(c_source, c_direction)
}

func Fn_gtk_icon_source_set_direction_wildcarded(source unsafe.Pointer, setting bool) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_setting := toCBool(setting)

	C.gtk_icon_source_set_direction_wildcarded(c_source, c_setting)
}

func Fn_gtk_icon_source_set_filename(source unsafe.Pointer, filename string) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_icon_source_set_filename(c_source, c_filename)
}

func Fn_gtk_icon_source_set_icon_name(source unsafe.Pointer, iconName *string) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	var c_iconNameValue (*C.gchar)
	if iconName != nil {
		c_iconNameValue = (*C.gchar)(C.CString(*iconName))
		defer C.free(unsafe.Pointer(c_iconNameValue))
	}
	c_iconName := c_iconNameValue

	C.gtk_icon_source_set_icon_name(c_source, c_iconName)
}

func Fn_gtk_icon_source_set_pixbuf(source unsafe.Pointer, pixbuf unsafe.Pointer) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gtk_icon_source_set_pixbuf(c_source, c_pixbuf)
}

func Fn_gtk_icon_source_set_size(source unsafe.Pointer, size int) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_size := (C.GtkIconSize)(size)

	C.gtk_icon_source_set_size(c_source, c_size)
}

func Fn_gtk_icon_source_set_size_wildcarded(source unsafe.Pointer, setting bool) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_setting := toCBool(setting)

	C.gtk_icon_source_set_size_wildcarded(c_source, c_setting)
}

func Fn_gtk_icon_source_set_state(source unsafe.Pointer, state int) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_state := (C.GtkStateType)(state)

	C.gtk_icon_source_set_state(c_source, c_state)
}

func Fn_gtk_icon_source_set_state_wildcarded(source unsafe.Pointer, setting bool) {
	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_setting := toCBool(setting)

	C.gtk_icon_source_set_state_wildcarded(c_source, c_setting)
}

func Fn_gtk_paper_size_is_custom(size unsafe.Pointer) bool {
	c_size := (*C.GtkPaperSize)(unsafe.Pointer(size))

	ret := C.gtk_paper_size_is_custom(c_size)

	return toGoBool(ret)
}

func Fn_gtk_paper_size_is_ipp(size unsafe.Pointer) bool {
	c_size := (*C.GtkPaperSize)(unsafe.Pointer(size))

	ret := C.gtk_paper_size_is_ipp(c_size)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_border(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) bool {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	c_gstring := (*C.GString)(unsafe.Pointer(gstring))

	c_propertyValue := (*C.GValue)(unsafe.Pointer(propertyValue))

	ret := C.gtk_rc_property_parse_border(c_pspec, c_gstring, c_propertyValue)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_color(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) bool {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	c_gstring := (*C.GString)(unsafe.Pointer(gstring))

	c_propertyValue := (*C.GValue)(unsafe.Pointer(propertyValue))

	ret := C.gtk_rc_property_parse_color(c_pspec, c_gstring, c_propertyValue)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_enum(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) bool {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	c_gstring := (*C.GString)(unsafe.Pointer(gstring))

	c_propertyValue := (*C.GValue)(unsafe.Pointer(propertyValue))

	ret := C.gtk_rc_property_parse_enum(c_pspec, c_gstring, c_propertyValue)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_flags(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) bool {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	c_gstring := (*C.GString)(unsafe.Pointer(gstring))

	c_propertyValue := (*C.GValue)(unsafe.Pointer(propertyValue))

	ret := C.gtk_rc_property_parse_flags(c_pspec, c_gstring, c_propertyValue)

	return toGoBool(ret)
}

func Fn_gtk_rc_property_parse_requisition(pspec unsafe.Pointer, gstring unsafe.Pointer, propertyValue unsafe.Pointer) bool {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	c_gstring := (*C.GString)(unsafe.Pointer(gstring))

	c_propertyValue := (*C.GValue)(unsafe.Pointer(propertyValue))

	ret := C.gtk_rc_property_parse_requisition(c_pspec, c_gstring, c_propertyValue)

	return toGoBool(ret)
}

func Fn_gtk_recent_info_create_app_info(info unsafe.Pointer, appName *string, error unsafe.Pointer) unsafe.Pointer {
	c_info := (*C.GtkRecentInfo)(unsafe.Pointer(info))

	var c_appNameValue (*C.gchar)
	if appName != nil {
		c_appNameValue = (*C.gchar)(C.CString(*appName))
		defer C.free(unsafe.Pointer(c_appNameValue))
	}
	c_appName := c_appNameValue

	cError := (**C.GError)(error)

	ret := C.gtk_recent_info_create_app_info(c_info, c_appName, cError)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_recent_info_get_application_info : parameter 'app_exec' is non array with indirect count > 1

func Fn_gtk_requisition_copy(requisition unsafe.Pointer) unsafe.Pointer {
	c_requisition := (*C.GtkRequisition)(unsafe.Pointer(requisition))

	ret := C.gtk_requisition_copy(c_requisition)

	return unsafe.Pointer(ret)
}

func Fn_gtk_requisition_free(requisition unsafe.Pointer) {
	c_requisition := (*C.GtkRequisition)(unsafe.Pointer(requisition))

	C.gtk_requisition_free(c_requisition)
}

func Fn_gtk_selection_data_copy(data unsafe.Pointer) unsafe.Pointer {
	c_data := (*C.GtkSelectionData)(unsafe.Pointer(data))

	ret := C.gtk_selection_data_copy(c_data)

	return unsafe.Pointer(ret)
}

func Fn_gtk_selection_data_free(data unsafe.Pointer) {
	c_data := (*C.GtkSelectionData)(unsafe.Pointer(data))

	C.gtk_selection_data_free(c_data)
}

// UNSUPPORTED : gtk_selection_data_get_data : no array length

func Fn_gtk_selection_data_get_targets(selectionData unsafe.Pointer, targets *[]unsafe.Pointer, nAtoms *int) bool {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	var c_targetsArrayPointer (*C.GdkAtom)
	c_targets := &c_targetsArrayPointer

	c_nAtoms := (*C.gint)(unsafe.Pointer(nAtoms))

	ret := C.gtk_selection_data_get_targets(c_selectionData, c_targets, c_nAtoms)

	targetsOutLen := int(*c_nAtoms)
	targetsOut := make([]unsafe.Pointer, targetsOutLen, targetsOutLen)
	if targetsOutLen > 0 {
		targetsOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_targetsArrayPointer))[:targetsOutLen:targetsOutLen]
	}
	*targets = targetsOut

	return toGoBool(ret)
}

func Fn_gtk_selection_data_get_text(selectionData unsafe.Pointer) string {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	ret := C.gtk_selection_data_get_text(c_selectionData)

	return C.GoString((*C.char)(unsafe.Pointer(ret)))
}

// UNSUPPORTED : gtk_selection_data_get_uris : no array length

func Fn_gtk_selection_data_set(selectionData unsafe.Pointer, type_ unsafe.Pointer, format int, data []uint8, length int) {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	c_type_ := (C.GdkAtom)(type_)

	c_format := (C.gint)(format)

	c_data := (*C.guchar)(unsafe.Pointer(&data[0]))

	c_length := (C.gint)(length)

	C.gtk_selection_data_set(c_selectionData, c_type_, c_format, c_data, c_length)
}

func Fn_gtk_selection_data_set_text(selectionData unsafe.Pointer, str string, len_ int) bool {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_len_ := (C.gint)(len_)

	ret := C.gtk_selection_data_set_text(c_selectionData, c_str, c_len_)

	return toGoBool(ret)
}

func Fn_gtk_selection_data_targets_include_text(selectionData unsafe.Pointer) bool {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	ret := C.gtk_selection_data_targets_include_text(c_selectionData)

	return toGoBool(ret)
}

func Fn_gtk_stock_item_copy(item unsafe.Pointer) unsafe.Pointer {
	c_item := (*C.GtkStockItem)(unsafe.Pointer(item))

	ret := C.gtk_stock_item_copy(c_item)

	return unsafe.Pointer(ret)
}

func Fn_gtk_stock_item_free(item unsafe.Pointer) {
	c_item := (*C.GtkStockItem)(unsafe.Pointer(item))

	C.gtk_stock_item_free(c_item)
}

func Fn_gtk_symbolic_color_to_string(color unsafe.Pointer) string {
	c_color := (*C.GtkSymbolicColor)(unsafe.Pointer(color))

	ret := C.gtk_symbolic_color_to_string(c_color)

	return C.GoString(ret)
}

func Fn_gtk_target_entry_new(target string, flags uint, info uint) unsafe.Pointer {
	c_target := (*C.gchar)(C.CString(target))
	defer C.free(unsafe.Pointer(c_target))

	c_flags := (C.guint)(flags)

	c_info := (C.guint)(info)

	ret := C.gtk_target_entry_new(c_target, c_flags, c_info)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_entry_copy(data unsafe.Pointer) unsafe.Pointer {
	c_data := (*C.GtkTargetEntry)(unsafe.Pointer(data))

	ret := C.gtk_target_entry_copy(c_data)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_entry_free(data unsafe.Pointer) {
	c_data := (*C.GtkTargetEntry)(unsafe.Pointer(data))

	C.gtk_target_entry_free(c_data)
}

func Fn_gtk_target_list_new(targets []TargetEntry, ntargets uint) unsafe.Pointer {
	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_ntargets := (C.guint)(ntargets)

	ret := C.gtk_target_list_new(c_targets, c_ntargets)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_list_add(list unsafe.Pointer, target unsafe.Pointer, flags uint, info uint) {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	c_target := (C.GdkAtom)(target)

	c_flags := (C.guint)(flags)

	c_info := (C.guint)(info)

	C.gtk_target_list_add(c_list, c_target, c_flags, c_info)
}

func Fn_gtk_target_list_add_table(list unsafe.Pointer, targets []TargetEntry, ntargets uint) {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_ntargets := (C.guint)(ntargets)

	C.gtk_target_list_add_table(c_list, c_targets, c_ntargets)
}

func Fn_gtk_target_list_find(list unsafe.Pointer, target unsafe.Pointer, info *uint) bool {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	c_target := (C.GdkAtom)(target)

	c_info := (*C.guint)(unsafe.Pointer(info))

	ret := C.gtk_target_list_find(c_list, c_target, c_info)

	return toGoBool(ret)
}

func Fn_gtk_target_list_ref(list unsafe.Pointer) unsafe.Pointer {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	ret := C.gtk_target_list_ref(c_list)

	return unsafe.Pointer(ret)
}

func Fn_gtk_target_list_remove(list unsafe.Pointer, target unsafe.Pointer) {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	c_target := (C.GdkAtom)(target)

	C.gtk_target_list_remove(c_list, c_target)
}

func Fn_gtk_target_list_unref(list unsafe.Pointer) {
	c_list := (*C.GtkTargetList)(unsafe.Pointer(list))

	C.gtk_target_list_unref(c_list)
}

func Fn_gtk_text_attributes_new() unsafe.Pointer {
	ret := C.gtk_text_attributes_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_copy(src unsafe.Pointer) unsafe.Pointer {
	c_src := (*C.GtkTextAttributes)(unsafe.Pointer(src))

	ret := C.gtk_text_attributes_copy(c_src)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_copy_values(src unsafe.Pointer, dest unsafe.Pointer) {
	c_src := (*C.GtkTextAttributes)(unsafe.Pointer(src))

	c_dest := (*C.GtkTextAttributes)(unsafe.Pointer(dest))

	C.gtk_text_attributes_copy_values(c_src, c_dest)
}

func Fn_gtk_text_attributes_ref(values unsafe.Pointer) unsafe.Pointer {
	c_values := (*C.GtkTextAttributes)(unsafe.Pointer(values))

	ret := C.gtk_text_attributes_ref(c_values)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_attributes_unref(values unsafe.Pointer) {
	c_values := (*C.GtkTextAttributes)(unsafe.Pointer(values))

	C.gtk_text_attributes_unref(c_values)
}

func Fn_gtk_text_iter_backward_char(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_backward_char(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_chars(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_backward_chars(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_cursor_position(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_backward_cursor_position(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_cursor_positions(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_backward_cursor_positions(c_iter, c_count)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_backward_find_char : parameter 'pred' is callback

func Fn_gtk_text_iter_backward_line(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_backward_line(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_lines(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_backward_lines(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_search(iter unsafe.Pointer, str string, flags int, matchStart unsafe.Pointer, matchEnd unsafe.Pointer, limit unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_flags := (C.GtkTextSearchFlags)(flags)

	c_matchStart := (*C.GtkTextIter)(unsafe.Pointer(matchStart))

	c_matchEnd := (*C.GtkTextIter)(unsafe.Pointer(matchEnd))

	c_limit := (*C.GtkTextIter)(unsafe.Pointer(limit))

	ret := C.gtk_text_iter_backward_search(c_iter, c_str, c_flags, c_matchStart, c_matchEnd, c_limit)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_sentence_start(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_backward_sentence_start(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_sentence_starts(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_backward_sentence_starts(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_to_tag_toggle(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_backward_to_tag_toggle(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_word_start(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_backward_word_start(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_backward_word_starts(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_backward_word_starts(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_begins_tag(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_begins_tag(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_can_insert(iter unsafe.Pointer, defaultEditability bool) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_defaultEditability := toCBool(defaultEditability)

	ret := C.gtk_text_iter_can_insert(c_iter, c_defaultEditability)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_compare(lhs unsafe.Pointer, rhs unsafe.Pointer) int {
	c_lhs := (*C.GtkTextIter)(unsafe.Pointer(lhs))

	c_rhs := (*C.GtkTextIter)(unsafe.Pointer(rhs))

	ret := C.gtk_text_iter_compare(c_lhs, c_rhs)

	return (int)(ret)
}

func Fn_gtk_text_iter_copy(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_copy(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_editable(iter unsafe.Pointer, defaultSetting bool) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_defaultSetting := toCBool(defaultSetting)

	ret := C.gtk_text_iter_editable(c_iter, c_defaultSetting)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_line(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_ends_line(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_sentence(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_ends_sentence(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_tag(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_ends_tag(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_ends_word(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_ends_word(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_equal(lhs unsafe.Pointer, rhs unsafe.Pointer) bool {
	c_lhs := (*C.GtkTextIter)(unsafe.Pointer(lhs))

	c_rhs := (*C.GtkTextIter)(unsafe.Pointer(rhs))

	ret := C.gtk_text_iter_equal(c_lhs, c_rhs)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_char(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_char(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_chars(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_forward_chars(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_cursor_position(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_cursor_position(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_cursor_positions(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_forward_cursor_positions(c_iter, c_count)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_iter_forward_find_char : parameter 'pred' is callback

func Fn_gtk_text_iter_forward_line(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_line(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_lines(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_forward_lines(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_search(iter unsafe.Pointer, str string, flags int, matchStart unsafe.Pointer, matchEnd unsafe.Pointer, limit unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	c_flags := (C.GtkTextSearchFlags)(flags)

	c_matchStart := (*C.GtkTextIter)(unsafe.Pointer(matchStart))

	c_matchEnd := (*C.GtkTextIter)(unsafe.Pointer(matchEnd))

	c_limit := (*C.GtkTextIter)(unsafe.Pointer(limit))

	ret := C.gtk_text_iter_forward_search(c_iter, c_str, c_flags, c_matchStart, c_matchEnd, c_limit)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_sentence_end(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_sentence_end(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_sentence_ends(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_forward_sentence_ends(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_to_end(iter unsafe.Pointer) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	C.gtk_text_iter_forward_to_end(c_iter)
}

func Fn_gtk_text_iter_forward_to_line_end(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_to_line_end(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_to_tag_toggle(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_forward_to_tag_toggle(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_word_end(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_forward_word_end(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_forward_word_ends(iter unsafe.Pointer, count int) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_iter_forward_word_ends(c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_free(iter unsafe.Pointer) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	C.gtk_text_iter_free(c_iter)
}

func Fn_gtk_text_iter_get_attributes(iter unsafe.Pointer, values unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_values := (*C.GtkTextAttributes)(unsafe.Pointer(values))

	ret := C.gtk_text_iter_get_attributes(c_iter, c_values)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_get_buffer(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_buffer(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_bytes_in_line(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_bytes_in_line(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_char(iter unsafe.Pointer) rune {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_char(c_iter)

	return (rune)(ret)
}

func Fn_gtk_text_iter_get_chars_in_line(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_chars_in_line(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_child_anchor(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_child_anchor(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_language(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_language(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_line(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_line(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_line_index(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_line_index(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_line_offset(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_line_offset(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_marks(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_marks(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_offset(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_offset(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_pixbuf(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_pixbuf(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_slice(start unsafe.Pointer, end unsafe.Pointer) string {
	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_iter_get_slice(c_start, c_end)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_tags(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_tags(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_text(start unsafe.Pointer, end unsafe.Pointer) string {
	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_iter_get_text(c_start, c_end)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_toggled_tags(iter unsafe.Pointer, toggledOn bool) unsafe.Pointer {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_toggledOn := toCBool(toggledOn)

	ret := C.gtk_text_iter_get_toggled_tags(c_iter, c_toggledOn)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_iter_get_visible_line_index(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_visible_line_index(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_visible_line_offset(iter unsafe.Pointer) int {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_get_visible_line_offset(c_iter)

	return (int)(ret)
}

func Fn_gtk_text_iter_get_visible_slice(start unsafe.Pointer, end unsafe.Pointer) string {
	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_iter_get_visible_slice(c_start, c_end)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_get_visible_text(start unsafe.Pointer, end unsafe.Pointer) string {
	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_iter_get_visible_text(c_start, c_end)

	return C.GoString(ret)
}

func Fn_gtk_text_iter_has_tag(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_has_tag(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_in_range(iter unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_iter_in_range(c_iter, c_start, c_end)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_inside_sentence(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_inside_sentence(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_inside_word(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_inside_word(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_cursor_position(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_is_cursor_position(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_end(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_is_end(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_is_start(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_is_start(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_order(first unsafe.Pointer, second unsafe.Pointer) {
	c_first := (*C.GtkTextIter)(unsafe.Pointer(first))

	c_second := (*C.GtkTextIter)(unsafe.Pointer(second))

	C.gtk_text_iter_order(c_first, c_second)
}

func Fn_gtk_text_iter_set_line(iter unsafe.Pointer, lineNumber int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_lineNumber := (C.gint)(lineNumber)

	C.gtk_text_iter_set_line(c_iter, c_lineNumber)
}

func Fn_gtk_text_iter_set_line_index(iter unsafe.Pointer, byteOnLine int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_byteOnLine := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_line_index(c_iter, c_byteOnLine)
}

func Fn_gtk_text_iter_set_line_offset(iter unsafe.Pointer, charOnLine int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_charOnLine := (C.gint)(charOnLine)

	C.gtk_text_iter_set_line_offset(c_iter, c_charOnLine)
}

func Fn_gtk_text_iter_set_offset(iter unsafe.Pointer, charOffset int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_charOffset := (C.gint)(charOffset)

	C.gtk_text_iter_set_offset(c_iter, c_charOffset)
}

func Fn_gtk_text_iter_set_visible_line_index(iter unsafe.Pointer, byteOnLine int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_byteOnLine := (C.gint)(byteOnLine)

	C.gtk_text_iter_set_visible_line_index(c_iter, c_byteOnLine)
}

func Fn_gtk_text_iter_set_visible_line_offset(iter unsafe.Pointer, charOnLine int) {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_charOnLine := (C.gint)(charOnLine)

	C.gtk_text_iter_set_visible_line_offset(c_iter, c_charOnLine)
}

func Fn_gtk_text_iter_starts_line(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_starts_line(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_starts_sentence(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_starts_sentence(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_starts_word(iter unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_iter_starts_word(c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_iter_toggles_tag(iter unsafe.Pointer, tag unsafe.Pointer) bool {
	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_iter_toggles_tag(c_iter, c_tag)

	return toGoBool(ret)
}

func Fn_gtk_tree_iter_copy(iter unsafe.Pointer) unsafe.Pointer {
	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_iter_copy(c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_iter_free(iter unsafe.Pointer) {
	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_iter_free(c_iter)
}

func Fn_gtk_tree_path_new() unsafe.Pointer {
	ret := C.gtk_tree_path_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_new_first() unsafe.Pointer {
	ret := C.gtk_tree_path_new_first()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_new_from_string(path string) unsafe.Pointer {
	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	ret := C.gtk_tree_path_new_from_string(c_path)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_append_index(path unsafe.Pointer, index int) {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_index := (C.gint)(index)

	C.gtk_tree_path_append_index(c_path, c_index)
}

func Fn_gtk_tree_path_compare(a unsafe.Pointer, b unsafe.Pointer) int {
	c_a := (*C.GtkTreePath)(unsafe.Pointer(a))

	c_b := (*C.GtkTreePath)(unsafe.Pointer(b))

	ret := C.gtk_tree_path_compare(c_a, c_b)

	return (int)(ret)
}

func Fn_gtk_tree_path_copy(path unsafe.Pointer) unsafe.Pointer {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_path_copy(c_path)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_path_down(path unsafe.Pointer) {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_path_down(c_path)
}

func Fn_gtk_tree_path_free(path unsafe.Pointer) {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_path_free(c_path)
}

func Fn_gtk_tree_path_get_depth(path unsafe.Pointer) int {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_path_get_depth(c_path)

	return (int)(ret)
}

func Fn_gtk_tree_path_is_ancestor(path unsafe.Pointer, descendant unsafe.Pointer) bool {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_descendant := (*C.GtkTreePath)(unsafe.Pointer(descendant))

	ret := C.gtk_tree_path_is_ancestor(c_path, c_descendant)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_is_descendant(path unsafe.Pointer, ancestor unsafe.Pointer) bool {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_ancestor := (*C.GtkTreePath)(unsafe.Pointer(ancestor))

	ret := C.gtk_tree_path_is_descendant(c_path, c_ancestor)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_next(path unsafe.Pointer) {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_path_next(c_path)
}

func Fn_gtk_tree_path_prepend_index(path unsafe.Pointer, index int) {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_index := (C.gint)(index)

	C.gtk_tree_path_prepend_index(c_path, c_index)
}

func Fn_gtk_tree_path_prev(path unsafe.Pointer) bool {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_path_prev(c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_path_to_string(path unsafe.Pointer) string {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_path_to_string(c_path)

	return C.GoString(ret)
}

func Fn_gtk_tree_path_up(path unsafe.Pointer) bool {
	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_path_up(c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_row_reference_new(model unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	c_model := (*C.GtkTreeModel)(unsafe.Pointer(model))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_row_reference_new(c_model, c_path)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_new_proxy(proxy unsafe.Pointer, model unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	c_proxy := (*C.GObject)(unsafe.Pointer(proxy))

	c_model := (*C.GtkTreeModel)(unsafe.Pointer(model))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_row_reference_new_proxy(c_proxy, c_model, c_path)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_copy(reference unsafe.Pointer) unsafe.Pointer {
	c_reference := (*C.GtkTreeRowReference)(unsafe.Pointer(reference))

	ret := C.gtk_tree_row_reference_copy(c_reference)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_free(reference unsafe.Pointer) {
	c_reference := (*C.GtkTreeRowReference)(unsafe.Pointer(reference))

	C.gtk_tree_row_reference_free(c_reference)
}

func Fn_gtk_tree_row_reference_get_path(reference unsafe.Pointer) unsafe.Pointer {
	c_reference := (*C.GtkTreeRowReference)(unsafe.Pointer(reference))

	ret := C.gtk_tree_row_reference_get_path(c_reference)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_row_reference_valid(reference unsafe.Pointer) bool {
	c_reference := (*C.GtkTreeRowReference)(unsafe.Pointer(reference))

	ret := C.gtk_tree_row_reference_valid(c_reference)

	return toGoBool(ret)
}

func Fn_gtk_tree_row_reference_deleted(proxy unsafe.Pointer, path unsafe.Pointer) {
	c_proxy := (*C.GObject)(unsafe.Pointer(proxy))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_row_reference_deleted(c_proxy, c_path)
}

func Fn_gtk_tree_row_reference_inserted(proxy unsafe.Pointer, path unsafe.Pointer) {
	c_proxy := (*C.GObject)(unsafe.Pointer(proxy))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_row_reference_inserted(c_proxy, c_path)
}

func Fn_gtk_tree_row_reference_reordered(proxy unsafe.Pointer, path unsafe.Pointer, iter unsafe.Pointer, newOrder []int) {
	c_proxy := (*C.GObject)(unsafe.Pointer(proxy))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_newOrder := (*C.gint)(unsafe.Pointer(&newOrder[0]))

	C.gtk_tree_row_reference_reordered(c_proxy, c_path, c_iter, c_newOrder)
}

// UNSUPPORTED : gtk_widget_class_bind_template_callback_full : parameter 'callback_symbol' is callback

func Fn_gtk_widget_class_find_style_property(klass unsafe.Pointer, propertyName string) unsafe.Pointer {
	c_klass := (*C.GtkWidgetClass)(unsafe.Pointer(klass))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	ret := C.gtk_widget_class_find_style_property(c_klass, c_propertyName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_class_install_style_property(klass unsafe.Pointer, pspec unsafe.Pointer) {
	c_klass := (*C.GtkWidgetClass)(unsafe.Pointer(klass))

	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	C.gtk_widget_class_install_style_property(c_klass, c_pspec)
}

// UNSUPPORTED : gtk_widget_class_install_style_property_parser : parameter 'parser' is callback

func Fn_gtk_widget_class_list_style_properties(klass unsafe.Pointer, nProperties *uint) []unsafe.Pointer {
	c_klass := (*C.GtkWidgetClass)(unsafe.Pointer(klass))

	c_nProperties := (*C.guint)(unsafe.Pointer(nProperties))

	ret := C.gtk_widget_class_list_style_properties(c_klass, c_nProperties)

	retLen := int(*c_nProperties)
	retGo := make([]unsafe.Pointer, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

// UNSUPPORTED : gtk_widget_class_set_connect_func : parameter 'connect_func' is callback

func Fn_gtk_widget_path_iter_get_name(path unsafe.Pointer, pos int) string {
	c_path := (*C.GtkWidgetPath)(unsafe.Pointer(path))

	c_pos := (C.gint)(pos)

	ret := C.gtk_widget_path_iter_get_name(c_path, c_pos)

	return C.GoString(ret)
}

func Fn_gtk_widget_path_iter_get_sibling_index(path unsafe.Pointer, pos int) uint {
	c_path := (*C.GtkWidgetPath)(unsafe.Pointer(path))

	c_pos := (C.gint)(pos)

	ret := C.gtk_widget_path_iter_get_sibling_index(c_path, c_pos)

	return (uint)(ret)
}

func Fn_gtk_widget_path_iter_get_siblings(path unsafe.Pointer, pos int) unsafe.Pointer {
	c_path := (*C.GtkWidgetPath)(unsafe.Pointer(path))

	c_pos := (C.gint)(pos)

	ret := C.gtk_widget_path_iter_get_siblings(c_path, c_pos)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_groups_activate(object unsafe.Pointer, accelKey uint, accelMods int) bool {
	c_object := (*C.GObject)(unsafe.Pointer(object))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	ret := C.gtk_accel_groups_activate(c_object, c_accelKey, c_accelMods)

	return toGoBool(ret)
}

func Fn_gtk_accel_groups_from_object(object unsafe.Pointer) unsafe.Pointer {
	c_object := (*C.GObject)(unsafe.Pointer(object))

	ret := C.gtk_accel_groups_from_object(c_object)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accelerator_get_default_mod_mask() int {
	ret := C.gtk_accelerator_get_default_mod_mask()

	return (int)(ret)
}

func Fn_gtk_accelerator_name(acceleratorKey uint, acceleratorMods int) string {
	c_acceleratorKey := (C.guint)(acceleratorKey)

	c_acceleratorMods := (C.GdkModifierType)(acceleratorMods)

	ret := C.gtk_accelerator_name(c_acceleratorKey, c_acceleratorMods)

	return C.GoString(ret)
}

func Fn_gtk_accelerator_parse(accelerator string, acceleratorKey *uint, acceleratorMods *int) {
	c_accelerator := (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(c_accelerator))

	c_acceleratorKey := (*C.guint)(unsafe.Pointer(acceleratorKey))

	c_acceleratorMods := (*C.GdkModifierType)(unsafe.Pointer(acceleratorMods))

	C.gtk_accelerator_parse(c_accelerator, c_acceleratorKey, c_acceleratorMods)
}

func Fn_gtk_accelerator_set_default_mod_mask(defaultModMask int) {
	c_defaultModMask := (C.GdkModifierType)(defaultModMask)

	C.gtk_accelerator_set_default_mod_mask(c_defaultModMask)
}

func Fn_gtk_accelerator_valid(keyval uint, modifiers int) bool {
	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	ret := C.gtk_accelerator_valid(c_keyval, c_modifiers)

	return toGoBool(ret)
}

func Fn_gtk_bindings_activate(object unsafe.Pointer, keyval uint, modifiers int) bool {
	c_object := (*C.GObject)(unsafe.Pointer(object))

	c_keyval := (C.guint)(keyval)

	c_modifiers := (C.GdkModifierType)(modifiers)

	ret := C.gtk_bindings_activate(c_object, c_keyval, c_modifiers)

	return toGoBool(ret)
}

func Fn_gtk_check_version(requiredMajor uint, requiredMinor uint, requiredMicro uint) string {
	c_requiredMajor := (C.guint)(requiredMajor)

	c_requiredMinor := (C.guint)(requiredMinor)

	c_requiredMicro := (C.guint)(requiredMicro)

	ret := C.gtk_check_version(c_requiredMajor, c_requiredMinor, c_requiredMicro)

	return C.GoString(ret)
}

func Fn_gtk_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_gtk_distribute_natural_allocation(extraSpace int, nRequestedSizes uint, sizes unsafe.Pointer) int {
	c_extraSpace := (C.gint)(extraSpace)

	c_nRequestedSizes := (C.guint)(nRequestedSizes)

	c_sizes := (*C.GtkRequestedSize)(unsafe.Pointer(sizes))

	ret := C.gtk_distribute_natural_allocation(c_extraSpace, c_nRequestedSizes, c_sizes)

	return (int)(ret)
}

func Fn_gtk_drag_finish(context unsafe.Pointer, success bool, del bool, time uint32) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_success := toCBool(success)

	c_del := toCBool(del)

	c_time := (C.guint32)(time)

	C.gtk_drag_finish(c_context, c_success, c_del, c_time)
}

func Fn_gtk_drag_get_source_widget(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	ret := C.gtk_drag_get_source_widget(c_context)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_set_icon_default(context unsafe.Pointer) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	C.gtk_drag_set_icon_default(c_context)
}

func Fn_gtk_drag_set_icon_pixbuf(context unsafe.Pointer, pixbuf unsafe.Pointer, hotX int, hotY int) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	c_hotX := (C.gint)(hotX)

	c_hotY := (C.gint)(hotY)

	C.gtk_drag_set_icon_pixbuf(c_context, c_pixbuf, c_hotX, c_hotY)
}

func Fn_gtk_drag_set_icon_stock(context unsafe.Pointer, stockId string, hotX int, hotY int) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_hotX := (C.gint)(hotX)

	c_hotY := (C.gint)(hotY)

	C.gtk_drag_set_icon_stock(c_context, c_stockId, c_hotX, c_hotY)
}

func Fn_gtk_drag_set_icon_surface(context unsafe.Pointer, surface unsafe.Pointer) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_surface := (*C.cairo_surface_t)(unsafe.Pointer(surface))

	C.gtk_drag_set_icon_surface(c_context, c_surface)
}

func Fn_gtk_drag_set_icon_widget(context unsafe.Pointer, widget unsafe.Pointer, hotX int, hotY int) {
	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_hotX := (C.gint)(hotX)

	c_hotY := (C.gint)(hotY)

	C.gtk_drag_set_icon_widget(c_context, c_widget, c_hotX, c_hotY)
}

func Fn_gtk_events_pending() bool {
	ret := C.gtk_events_pending()

	return toGoBool(ret)
}

func Fn_gtk_false() bool {
	ret := C.gtk_false()

	return toGoBool(ret)
}

func Fn_gtk_get_current_event() unsafe.Pointer {
	ret := C.gtk_get_current_event()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_current_event_device() unsafe.Pointer {
	ret := C.gtk_get_current_event_device()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_current_event_state(state *int) bool {
	c_state := (*C.GdkModifierType)(unsafe.Pointer(state))

	ret := C.gtk_get_current_event_state(c_state)

	return toGoBool(ret)
}

func Fn_gtk_get_current_event_time() uint32 {
	ret := C.gtk_get_current_event_time()

	return (uint32)(ret)
}

func Fn_gtk_get_debug_flags() uint {
	ret := C.gtk_get_debug_flags()

	return (uint)(ret)
}

func Fn_gtk_get_default_language() unsafe.Pointer {
	ret := C.gtk_get_default_language()

	return unsafe.Pointer(ret)
}

func Fn_gtk_get_event_widget(event unsafe.Pointer) unsafe.Pointer {
	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	ret := C.gtk_get_event_widget(c_event)

	return unsafe.Pointer(ret)
}

func Fn_gtk_grab_get_current() unsafe.Pointer {
	ret := C.gtk_grab_get_current()

	return unsafe.Pointer(ret)
}

func Fn_gtk_init(argc *int, argv *[]string) {
	c_argc := (*C.int)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	C.gtk_init(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut
}

func Fn_gtk_init_check(argc *int, argv *[]string) bool {
	c_argc := (*C.int)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	ret := C.gtk_init_check(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_key_snooper_install : parameter 'snooper' is callback

func Fn_gtk_key_snooper_remove(snooperHandlerId uint) {
	c_snooperHandlerId := (C.guint)(snooperHandlerId)

	C.gtk_key_snooper_remove(c_snooperHandlerId)
}

func Fn_gtk_main() {
	C.gtk_main()
}

func Fn_gtk_main_do_event(event unsafe.Pointer) {
	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	C.gtk_main_do_event(c_event)
}

func Fn_gtk_main_iteration() bool {
	ret := C.gtk_main_iteration()

	return toGoBool(ret)
}

func Fn_gtk_main_iteration_do(blocking bool) bool {
	c_blocking := toCBool(blocking)

	ret := C.gtk_main_iteration_do(c_blocking)

	return toGoBool(ret)
}

func Fn_gtk_main_level() uint {
	ret := C.gtk_main_level()

	return (uint)(ret)
}

func Fn_gtk_main_quit() {
	C.gtk_main_quit()
}

func Fn_gtk_paint_arrow(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, arrowType int, fill bool, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_arrowType := (C.GtkArrowType)(arrowType)

	c_fill := toCBool(fill)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_arrow(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_arrowType, c_fill, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_box(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_box(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_box_gap(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gapSide := (C.GtkPositionType)(gapSide)

	c_gapX := (C.gint)(gapX)

	c_gapWidth := (C.gint)(gapWidth)

	C.gtk_paint_box_gap(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gapSide, c_gapX, c_gapWidth)
}

func Fn_gtk_paint_check(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_check(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_diamond(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_diamond(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_expander(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, x int, y int, expanderStyle int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_expanderStyle := (C.GtkExpanderStyle)(expanderStyle)

	C.gtk_paint_expander(c_style, c_cr, c_stateType, c_widget, c_detail, c_x, c_y, c_expanderStyle)
}

func Fn_gtk_paint_extension(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int, gapSide int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gapSide := (C.GtkPositionType)(gapSide)

	C.gtk_paint_extension(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gapSide)
}

func Fn_gtk_paint_flat_box(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_flat_box(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_focus(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_focus(c_style, c_cr, c_stateType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_handle(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int, orientation int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_handle(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)
}

func Fn_gtk_paint_hline(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, x1 int, x2 int, y int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x1 := (C.gint)(x1)

	c_x2 := (C.gint)(x2)

	c_y := (C.gint)(y)

	C.gtk_paint_hline(c_style, c_cr, c_stateType, c_widget, c_detail, c_x1, c_x2, c_y)
}

func Fn_gtk_paint_layout(style unsafe.Pointer, cr unsafe.Pointer, stateType int, useText bool, widget unsafe.Pointer, detail *string, x int, y int, layout unsafe.Pointer) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_useText := toCBool(useText)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_layout := (*C.PangoLayout)(unsafe.Pointer(layout))

	C.gtk_paint_layout(c_style, c_cr, c_stateType, c_useText, c_widget, c_detail, c_x, c_y, c_layout)
}

func Fn_gtk_paint_option(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_option(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_resize_grip(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, edge int, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_edge := (C.GdkWindowEdge)(edge)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_resize_grip(c_style, c_cr, c_stateType, c_widget, c_detail, c_edge, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_shadow(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_shadow(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_shadow_gap(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int, gapSide int, gapX int, gapWidth int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_gapSide := (C.GtkPositionType)(gapSide)

	c_gapX := (C.gint)(gapX)

	c_gapWidth := (C.gint)(gapWidth)

	C.gtk_paint_shadow_gap(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height, c_gapSide, c_gapX, c_gapWidth)
}

func Fn_gtk_paint_slider(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int, orientation int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	c_orientation := (C.GtkOrientation)(orientation)

	C.gtk_paint_slider(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height, c_orientation)
}

func Fn_gtk_paint_spinner(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, step uint, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_step := (C.guint)(step)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_spinner(c_style, c_cr, c_stateType, c_widget, c_detail, c_step, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_tab(style unsafe.Pointer, cr unsafe.Pointer, stateType int, shadowType int, widget unsafe.Pointer, detail *string, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_paint_tab(c_style, c_cr, c_stateType, c_shadowType, c_widget, c_detail, c_x, c_y, c_width, c_height)
}

func Fn_gtk_paint_vline(style unsafe.Pointer, cr unsafe.Pointer, stateType int, widget unsafe.Pointer, detail *string, y1 int, y2 int, x int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_stateType := (C.GtkStateType)(stateType)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	c_y1 := (C.gint)(y1)

	c_y2 := (C.gint)(y2)

	c_x := (C.gint)(x)

	C.gtk_paint_vline(c_style, c_cr, c_stateType, c_widget, c_detail, c_y1, c_y2, c_x)
}

func Fn_gtk_parse_args(argc *int, argv *[]string) bool {
	c_argc := (*C.int)(unsafe.Pointer(argc))

	var c_argvArrayPointer **C.gchar
	c_argv := &c_argvArrayPointer
	argvIndirected := *argv
	argvIndirectedLen := len(argvIndirected)
	c_argvArray := C.malloc((C.ulong)(argvIndirectedLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_argvArray))
	argvIndirectedSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArray))[:argvIndirectedLen:argvIndirectedLen]
	for argvIndirectedi, argvIndirectedString := range argvIndirected {
		argvIndirectedSlice[argvIndirectedi] = (*C.gchar)(C.CString(argvIndirectedString))
		defer C.free(unsafe.Pointer(argvIndirectedSlice[argvIndirectedi]))
	}
	if len(argvIndirectedSlice) > 0 {
		c_argvArrayPointer = &argvIndirectedSlice[0]
	}

	ret := C.gtk_parse_args(c_argc, c_argv)

	argvOutLen := int(*c_argc)
	argvOut := make([]string, argvOutLen, argvOutLen)
	if argvOutLen > 0 {
		argvOutCSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_argvArrayPointer))[:argvOutLen:argvOutLen]
		for argvOuti, argvOutCString := range argvOutCSlice {
			argvOut[argvOuti] = C.GoString(argvOutCString)
		}
	}
	*argv = argvOut

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_print_run_page_setup_dialog_async : parameter 'done_cb' is callback

func Fn_gtk_propagate_event(widget unsafe.Pointer, event unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	C.gtk_propagate_event(c_widget, c_event)
}

func Fn_gtk_rc_add_default_file(filename string) {
	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_add_default_file(c_filename)
}

func Fn_gtk_rc_find_module_in_path(moduleFile string) string {
	c_moduleFile := (*C.gchar)(C.CString(moduleFile))
	defer C.free(unsafe.Pointer(c_moduleFile))

	ret := C.gtk_rc_find_module_in_path(c_moduleFile)

	return C.GoString(ret)
}

func Fn_gtk_rc_find_pixmap_in_path(settings unsafe.Pointer, scanner unsafe.Pointer, pixmapFile string) string {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_scanner := (*C.GScanner)(unsafe.Pointer(scanner))

	c_pixmapFile := (*C.gchar)(C.CString(pixmapFile))
	defer C.free(unsafe.Pointer(c_pixmapFile))

	ret := C.gtk_rc_find_pixmap_in_path(c_settings, c_scanner, c_pixmapFile)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_rc_get_default_files : no array length

func Fn_gtk_rc_get_im_module_file() string {
	ret := C.gtk_rc_get_im_module_file()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_im_module_path() string {
	ret := C.gtk_rc_get_im_module_path()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_module_dir() string {
	ret := C.gtk_rc_get_module_dir()

	return C.GoString(ret)
}

func Fn_gtk_rc_get_style(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_rc_get_style(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_get_style_by_paths(settings unsafe.Pointer, widgetPath *string, classPath *string, type_ uint64) unsafe.Pointer {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	var c_widgetPathValue (*C.char)
	if widgetPath != nil {
		c_widgetPathValue = (*C.char)(C.CString(*widgetPath))
		defer C.free(unsafe.Pointer(c_widgetPathValue))
	}
	c_widgetPath := c_widgetPathValue

	var c_classPathValue (*C.char)
	if classPath != nil {
		c_classPathValue = (*C.char)(C.CString(*classPath))
		defer C.free(unsafe.Pointer(c_classPathValue))
	}
	c_classPath := c_classPathValue

	c_type_ := (C.GType)(type_)

	ret := C.gtk_rc_get_style_by_paths(c_settings, c_widgetPath, c_classPath, c_type_)

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_get_theme_dir() string {
	ret := C.gtk_rc_get_theme_dir()

	return C.GoString(ret)
}

func Fn_gtk_rc_parse(filename string) {
	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_rc_parse(c_filename)
}

func Fn_gtk_rc_parse_color(scanner unsafe.Pointer, color unsafe.Pointer) uint {
	c_scanner := (*C.GScanner)(unsafe.Pointer(scanner))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	ret := C.gtk_rc_parse_color(c_scanner, c_color)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_priority(scanner unsafe.Pointer, priority *int) uint {
	c_scanner := (*C.GScanner)(unsafe.Pointer(scanner))

	c_priority := (*C.GtkPathPriorityType)(unsafe.Pointer(priority))

	ret := C.gtk_rc_parse_priority(c_scanner, c_priority)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_state(scanner unsafe.Pointer, state *int) uint {
	c_scanner := (*C.GScanner)(unsafe.Pointer(scanner))

	c_state := (*C.GtkStateType)(unsafe.Pointer(state))

	ret := C.gtk_rc_parse_state(c_scanner, c_state)

	return (uint)(ret)
}

func Fn_gtk_rc_parse_string(rcString string) {
	c_rcString := (*C.gchar)(C.CString(rcString))
	defer C.free(unsafe.Pointer(c_rcString))

	C.gtk_rc_parse_string(c_rcString)
}

func Fn_gtk_rc_reparse_all() bool {
	ret := C.gtk_rc_reparse_all()

	return toGoBool(ret)
}

func Fn_gtk_rc_reparse_all_for_settings(settings unsafe.Pointer, forceLoad bool) bool {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_forceLoad := toCBool(forceLoad)

	ret := C.gtk_rc_reparse_all_for_settings(c_settings, c_forceLoad)

	return toGoBool(ret)
}

func Fn_gtk_rc_scanner_new() unsafe.Pointer {
	ret := C.gtk_rc_scanner_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_set_default_files(filenames []string) {
	filenamesLen := len(filenames)
	c_filenamesArray := C.malloc((C.ulong)(filenamesLen) * C.sizeof_gpointer)
	defer C.free(unsafe.Pointer(c_filenamesArray))
	filenamesSlice := (*[1 << 30](*C.gchar))(unsafe.Pointer(c_filenamesArray))[:filenamesLen:filenamesLen]
	for filenamesi, filenamesString := range filenames {
		filenamesSlice[filenamesi] = (*C.gchar)(C.CString(filenamesString))
		defer C.free(unsafe.Pointer(filenamesSlice[filenamesi]))
	}
	c_filenames := &filenamesSlice[0]

	C.gtk_rc_set_default_files(c_filenames)
}

func Fn_gtk_selection_add_target(widget unsafe.Pointer, selection unsafe.Pointer, target unsafe.Pointer, info uint) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	c_target := (C.GdkAtom)(target)

	c_info := (C.guint)(info)

	C.gtk_selection_add_target(c_widget, c_selection, c_target, c_info)
}

func Fn_gtk_selection_add_targets(widget unsafe.Pointer, selection unsafe.Pointer, targets []TargetEntry, ntargets uint) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_ntargets := (C.guint)(ntargets)

	C.gtk_selection_add_targets(c_widget, c_selection, c_targets, c_ntargets)
}

func Fn_gtk_selection_clear_targets(widget unsafe.Pointer, selection unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	C.gtk_selection_clear_targets(c_widget, c_selection)
}

func Fn_gtk_selection_convert(widget unsafe.Pointer, selection unsafe.Pointer, target unsafe.Pointer, time uint32) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	c_target := (C.GdkAtom)(target)

	c_time := (C.guint32)(time)

	ret := C.gtk_selection_convert(c_widget, c_selection, c_target, c_time)

	return toGoBool(ret)
}

func Fn_gtk_selection_owner_set(widget unsafe.Pointer, selection unsafe.Pointer, time uint32) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	c_time := (C.guint32)(time)

	ret := C.gtk_selection_owner_set(c_widget, c_selection, c_time)

	return toGoBool(ret)
}

func Fn_gtk_selection_owner_set_for_display(display unsafe.Pointer, widget unsafe.Pointer, selection unsafe.Pointer, time uint32) bool {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	c_time := (C.guint32)(time)

	ret := C.gtk_selection_owner_set_for_display(c_display, c_widget, c_selection, c_time)

	return toGoBool(ret)
}

func Fn_gtk_selection_remove_all(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_selection_remove_all(c_widget)
}

func Fn_gtk_set_debug_flags(flags uint) {
	c_flags := (C.guint)(flags)

	C.gtk_set_debug_flags(c_flags)
}

func Fn_gtk_stock_add(items []StockItem, nItems uint) {
	c_items := (*C.GtkStockItem)(unsafe.Pointer(&items[0]))

	c_nItems := (C.guint)(nItems)

	C.gtk_stock_add(c_items, c_nItems)
}

func Fn_gtk_stock_add_static(items []StockItem, nItems uint) {
	c_items := (*C.GtkStockItem)(unsafe.Pointer(&items[0]))

	c_nItems := (C.guint)(nItems)

	C.gtk_stock_add_static(c_items, c_nItems)
}

func Fn_gtk_stock_list_ids() unsafe.Pointer {
	ret := C.gtk_stock_list_ids()

	return unsafe.Pointer(ret)
}

func Fn_gtk_stock_lookup(stockId string, item unsafe.Pointer) bool {
	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_item := (*C.GtkStockItem)(unsafe.Pointer(item))

	ret := C.gtk_stock_lookup(c_stockId, c_item)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_stock_set_translate_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_get_row_drag_data : parameter 'tree_model' is non array with indirect count > 1

func Fn_gtk_tree_set_row_drag_data(selectionData unsafe.Pointer, treeModel unsafe.Pointer, path unsafe.Pointer) bool {
	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_set_row_drag_data(c_selectionData, c_treeModel, c_path)

	return toGoBool(ret)
}

func Fn_gtk_true() bool {
	ret := C.gtk_true()

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_about_dialog_get_artists : no array length

// UNSUPPORTED : gtk_about_dialog_get_authors : no array length

// UNSUPPORTED : gtk_about_dialog_get_documenters : no array length

func Fn_gtk_accel_group_new() unsafe.Pointer {
	ret := C.gtk_accel_group_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_group_activate(accelGroup unsafe.Pointer, accelQuark uint32, acceleratable unsafe.Pointer, accelKey uint, accelMods int) bool {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelQuark := (C.GQuark)(accelQuark)

	c_acceleratable := (*C.GObject)(unsafe.Pointer(acceleratable))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	ret := C.gtk_accel_group_activate(c_accelGroup, c_accelQuark, c_acceleratable, c_accelKey, c_accelMods)

	return toGoBool(ret)
}

func Fn_gtk_accel_group_connect(accelGroup unsafe.Pointer, accelKey uint, accelMods int, accelFlags int, closure unsafe.Pointer) {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	c_accelFlags := (C.GtkAccelFlags)(accelFlags)

	c_closure := (*C.GClosure)(unsafe.Pointer(closure))

	C.gtk_accel_group_connect(c_accelGroup, c_accelKey, c_accelMods, c_accelFlags, c_closure)
}

func Fn_gtk_accel_group_connect_by_path(accelGroup unsafe.Pointer, accelPath string, closure unsafe.Pointer) {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelPath := (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(c_accelPath))

	c_closure := (*C.GClosure)(unsafe.Pointer(closure))

	C.gtk_accel_group_connect_by_path(c_accelGroup, c_accelPath, c_closure)
}

func Fn_gtk_accel_group_disconnect(accelGroup unsafe.Pointer, closure unsafe.Pointer) bool {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_closure := (*C.GClosure)(unsafe.Pointer(closure))

	ret := C.gtk_accel_group_disconnect(c_accelGroup, c_closure)

	return toGoBool(ret)
}

func Fn_gtk_accel_group_disconnect_key(accelGroup unsafe.Pointer, accelKey uint, accelMods int) bool {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	ret := C.gtk_accel_group_disconnect_key(c_accelGroup, c_accelKey, c_accelMods)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_accel_group_find : parameter 'find_func' is callback

func Fn_gtk_accel_group_lock(accelGroup unsafe.Pointer) {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_accel_group_lock(c_accelGroup)
}

func Fn_gtk_accel_group_query(accelGroup unsafe.Pointer, accelKey uint, accelMods int, nEntries *uint) []AccelGroupEntry {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	c_nEntries := (*C.guint)(unsafe.Pointer(nEntries))

	ret := C.gtk_accel_group_query(c_accelGroup, c_accelKey, c_accelMods, c_nEntries)

	retLen := int(*c_nEntries)
	retGo := make([]AccelGroupEntry, retLen, retLen)
	if retLen > 0 {
		retGo = (*[1 << 30](AccelGroupEntry))(unsafe.Pointer(ret))[:retLen:retLen]
	}
	return retGo
}

func Fn_gtk_accel_group_unlock(accelGroup unsafe.Pointer) {
	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_accel_group_unlock(c_accelGroup)
}

func Fn_gtk_accel_group_from_accel_closure(closure unsafe.Pointer) unsafe.Pointer {
	c_closure := (*C.GClosure)(unsafe.Pointer(closure))

	ret := C.gtk_accel_group_from_accel_closure(c_closure)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_new(string_ string) unsafe.Pointer {
	c_string_ := (*C.gchar)(C.CString(string_))
	defer C.free(unsafe.Pointer(c_string_))

	ret := C.gtk_accel_label_new(c_string_)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_get_accel_widget(accelLabel unsafe.Pointer) unsafe.Pointer {
	c_accelLabel := (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel))

	ret := C.gtk_accel_label_get_accel_widget(c_accelLabel)

	return unsafe.Pointer(ret)
}

func Fn_gtk_accel_label_get_accel_width(accelLabel unsafe.Pointer) uint {
	c_accelLabel := (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel))

	ret := C.gtk_accel_label_get_accel_width(c_accelLabel)

	return (uint)(ret)
}

func Fn_gtk_accel_label_refetch(accelLabel unsafe.Pointer) bool {
	c_accelLabel := (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel))

	ret := C.gtk_accel_label_refetch(c_accelLabel)

	return toGoBool(ret)
}

func Fn_gtk_accel_label_set_accel_closure(accelLabel unsafe.Pointer, accelClosure unsafe.Pointer) {
	c_accelLabel := (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel))

	c_accelClosure := (*C.GClosure)(unsafe.Pointer(accelClosure))

	C.gtk_accel_label_set_accel_closure(c_accelLabel, c_accelClosure)
}

func Fn_gtk_accel_label_set_accel_widget(accelLabel unsafe.Pointer, accelWidget unsafe.Pointer) {
	c_accelLabel := (*C.GtkAccelLabel)(unsafe.Pointer(accelLabel))

	c_accelWidget := (*C.GtkWidget)(unsafe.Pointer(accelWidget))

	C.gtk_accel_label_set_accel_widget(c_accelLabel, c_accelWidget)
}

func Fn_gtk_accel_map_add_entry(accelPath string, accelKey uint, accelMods int) {
	c_accelPath := (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(c_accelPath))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	C.gtk_accel_map_add_entry(c_accelPath, c_accelKey, c_accelMods)
}

func Fn_gtk_accel_map_add_filter(filterPattern string) {
	c_filterPattern := (*C.gchar)(C.CString(filterPattern))
	defer C.free(unsafe.Pointer(c_filterPattern))

	C.gtk_accel_map_add_filter(c_filterPattern)
}

func Fn_gtk_accel_map_change_entry(accelPath string, accelKey uint, accelMods int, replace bool) bool {
	c_accelPath := (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(c_accelPath))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	c_replace := toCBool(replace)

	ret := C.gtk_accel_map_change_entry(c_accelPath, c_accelKey, c_accelMods, c_replace)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_accel_map_foreach : parameter 'foreach_func' is callback

// UNSUPPORTED : gtk_accel_map_foreach_unfiltered : parameter 'foreach_func' is callback

func Fn_gtk_accel_map_load(fileName string) {
	c_fileName := (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(c_fileName))

	C.gtk_accel_map_load(c_fileName)
}

func Fn_gtk_accel_map_load_fd(fd int) {
	c_fd := (C.gint)(fd)

	C.gtk_accel_map_load_fd(c_fd)
}

func Fn_gtk_accel_map_load_scanner(scanner unsafe.Pointer) {
	c_scanner := (*C.GScanner)(unsafe.Pointer(scanner))

	C.gtk_accel_map_load_scanner(c_scanner)
}

func Fn_gtk_accel_map_lookup_entry(accelPath string, key unsafe.Pointer) bool {
	c_accelPath := (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(c_accelPath))

	c_key := (*C.GtkAccelKey)(unsafe.Pointer(key))

	ret := C.gtk_accel_map_lookup_entry(c_accelPath, c_key)

	return toGoBool(ret)
}

func Fn_gtk_accel_map_save(fileName string) {
	c_fileName := (*C.gchar)(C.CString(fileName))
	defer C.free(unsafe.Pointer(c_fileName))

	C.gtk_accel_map_save(c_fileName)
}

func Fn_gtk_accel_map_save_fd(fd int) {
	c_fd := (C.gint)(fd)

	C.gtk_accel_map_save_fd(c_fd)
}

func Fn_gtk_accessible_connect_widget_destroyed(accessible unsafe.Pointer) {
	c_accessible := (*C.GtkAccessible)(unsafe.Pointer(accessible))

	C.gtk_accessible_connect_widget_destroyed(c_accessible)
}

// UNSUPPORTED : gtk_action_group_add_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_radio_actions_full : parameter 'on_change' is callback

// UNSUPPORTED : gtk_action_group_add_toggle_actions_full : parameter 'destroy' is callback

// UNSUPPORTED : gtk_action_group_set_translate_func : parameter 'func' is callback

func Fn_gtk_adjustment_new(value float64, lower float64, upper float64, stepIncrement float64, pageIncrement float64, pageSize float64) unsafe.Pointer {
	c_value := (C.gdouble)(value)

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	c_stepIncrement := (C.gdouble)(stepIncrement)

	c_pageIncrement := (C.gdouble)(pageIncrement)

	c_pageSize := (C.gdouble)(pageSize)

	ret := C.gtk_adjustment_new(c_value, c_lower, c_upper, c_stepIncrement, c_pageIncrement, c_pageSize)

	return unsafe.Pointer(ret)
}

func Fn_gtk_adjustment_changed(adjustment unsafe.Pointer) {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_adjustment_changed(c_adjustment)
}

func Fn_gtk_adjustment_clamp_page(adjustment unsafe.Pointer, lower float64, upper float64) {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	c_lower := (C.gdouble)(lower)

	c_upper := (C.gdouble)(upper)

	C.gtk_adjustment_clamp_page(c_adjustment, c_lower, c_upper)
}

func Fn_gtk_adjustment_get_value(adjustment unsafe.Pointer) float64 {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	ret := C.gtk_adjustment_get_value(c_adjustment)

	return (float64)(ret)
}

func Fn_gtk_adjustment_set_value(adjustment unsafe.Pointer, value float64) {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	c_value := (C.gdouble)(value)

	C.gtk_adjustment_set_value(c_adjustment, c_value)
}

func Fn_gtk_adjustment_value_changed(adjustment unsafe.Pointer) {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_adjustment_value_changed(c_adjustment)
}

func Fn_gtk_alignment_new(xalign float32, yalign float32, xscale float32, yscale float32) unsafe.Pointer {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	ret := C.gtk_alignment_new(c_xalign, c_yalign, c_xscale, c_yscale)

	return unsafe.Pointer(ret)
}

func Fn_gtk_alignment_set(alignment unsafe.Pointer, xalign float32, yalign float32, xscale float32, yscale float32) {
	c_alignment := (*C.GtkAlignment)(unsafe.Pointer(alignment))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_xscale := (C.gfloat)(xscale)

	c_yscale := (C.gfloat)(yscale)

	C.gtk_alignment_set(c_alignment, c_xalign, c_yalign, c_xscale, c_yscale)
}

func Fn_gtk_app_chooser_button_get_heading(self unsafe.Pointer) string {
	c_self := (*C.GtkAppChooserButton)(unsafe.Pointer(self))

	ret := C.gtk_app_chooser_button_get_heading(c_self)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_button_set_heading(self unsafe.Pointer, heading string) {
	c_self := (*C.GtkAppChooserButton)(unsafe.Pointer(self))

	c_heading := (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(c_heading))

	C.gtk_app_chooser_button_set_heading(c_self, c_heading)
}

func Fn_gtk_app_chooser_dialog_get_heading(self unsafe.Pointer) string {
	c_self := (*C.GtkAppChooserDialog)(unsafe.Pointer(self))

	ret := C.gtk_app_chooser_dialog_get_heading(c_self)

	return C.GoString(ret)
}

func Fn_gtk_app_chooser_dialog_set_heading(self unsafe.Pointer, heading string) {
	c_self := (*C.GtkAppChooserDialog)(unsafe.Pointer(self))

	c_heading := (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(c_heading))

	C.gtk_app_chooser_dialog_set_heading(c_self, c_heading)
}

func Fn_gtk_app_chooser_widget_set_default_text(self unsafe.Pointer, text string) {
	c_self := (*C.GtkAppChooserWidget)(unsafe.Pointer(self))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_app_chooser_widget_set_default_text(c_self, c_text)
}

// UNSUPPORTED : gtk_application_get_accels_for_action : no array length

// UNSUPPORTED : gtk_application_get_actions_for_accel : no array length

// UNSUPPORTED : gtk_application_list_action_descriptions : no array length

func Fn_gtk_arrow_new(arrowType int, shadowType int) unsafe.Pointer {
	c_arrowType := (C.GtkArrowType)(arrowType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	ret := C.gtk_arrow_new(c_arrowType, c_shadowType)

	return unsafe.Pointer(ret)
}

func Fn_gtk_arrow_set(arrow unsafe.Pointer, arrowType int, shadowType int) {
	c_arrow := (*C.GtkArrow)(unsafe.Pointer(arrow))

	c_arrowType := (C.GtkArrowType)(arrowType)

	c_shadowType := (C.GtkShadowType)(shadowType)

	C.gtk_arrow_set(c_arrow, c_arrowType, c_shadowType)
}

func Fn_gtk_aspect_frame_new(label *string, xalign float32, yalign float32, ratio float32, obeyChild bool) unsafe.Pointer {
	var c_labelValue (*C.gchar)
	if label != nil {
		c_labelValue = (*C.gchar)(C.CString(*label))
		defer C.free(unsafe.Pointer(c_labelValue))
	}
	c_label := c_labelValue

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obeyChild := toCBool(obeyChild)

	ret := C.gtk_aspect_frame_new(c_label, c_xalign, c_yalign, c_ratio, c_obeyChild)

	return unsafe.Pointer(ret)
}

func Fn_gtk_aspect_frame_set(aspectFrame unsafe.Pointer, xalign float32, yalign float32, ratio float32, obeyChild bool) {
	c_aspectFrame := (*C.GtkAspectFrame)(unsafe.Pointer(aspectFrame))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	c_ratio := (C.gfloat)(ratio)

	c_obeyChild := toCBool(obeyChild)

	C.gtk_aspect_frame_set(c_aspectFrame, c_xalign, c_yalign, c_ratio, c_obeyChild)
}

// UNSUPPORTED : gtk_assistant_set_forward_page_func : parameter 'page_func' is callback

func Fn_gtk_bin_get_child(bin unsafe.Pointer) unsafe.Pointer {
	c_bin := (*C.GtkBin)(unsafe.Pointer(bin))

	ret := C.gtk_bin_get_child(c_bin)

	return unsafe.Pointer(ret)
}

func Fn_gtk_box_get_homogeneous(box unsafe.Pointer) bool {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	ret := C.gtk_box_get_homogeneous(c_box)

	return toGoBool(ret)
}

func Fn_gtk_box_get_spacing(box unsafe.Pointer) int {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	ret := C.gtk_box_get_spacing(c_box)

	return (int)(ret)
}

func Fn_gtk_box_pack_end(box unsafe.Pointer, child unsafe.Pointer, expand bool, fill bool, padding uint) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_expand := toCBool(expand)

	c_fill := toCBool(fill)

	c_padding := (C.guint)(padding)

	C.gtk_box_pack_end(c_box, c_child, c_expand, c_fill, c_padding)
}

func Fn_gtk_box_pack_start(box unsafe.Pointer, child unsafe.Pointer, expand bool, fill bool, padding uint) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_expand := toCBool(expand)

	c_fill := toCBool(fill)

	c_padding := (C.guint)(padding)

	C.gtk_box_pack_start(c_box, c_child, c_expand, c_fill, c_padding)
}

func Fn_gtk_box_query_child_packing(box unsafe.Pointer, child unsafe.Pointer, expand *bool, fill *bool, padding *uint, packType *int) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_expand := (*C.gboolean)(unsafe.Pointer(expand))

	c_fill := (*C.gboolean)(unsafe.Pointer(fill))

	c_padding := (*C.guint)(unsafe.Pointer(padding))

	c_packType := (*C.GtkPackType)(unsafe.Pointer(packType))

	C.gtk_box_query_child_packing(c_box, c_child, c_expand, c_fill, c_padding, c_packType)
}

func Fn_gtk_box_reorder_child(box unsafe.Pointer, child unsafe.Pointer, position int) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_position := (C.gint)(position)

	C.gtk_box_reorder_child(c_box, c_child, c_position)
}

func Fn_gtk_box_set_child_packing(box unsafe.Pointer, child unsafe.Pointer, expand bool, fill bool, padding uint, packType int) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_expand := toCBool(expand)

	c_fill := toCBool(fill)

	c_padding := (C.guint)(padding)

	c_packType := (C.GtkPackType)(packType)

	C.gtk_box_set_child_packing(c_box, c_child, c_expand, c_fill, c_padding, c_packType)
}

func Fn_gtk_box_set_homogeneous(box unsafe.Pointer, homogeneous bool) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_homogeneous := toCBool(homogeneous)

	C.gtk_box_set_homogeneous(c_box, c_homogeneous)
}

func Fn_gtk_box_set_spacing(box unsafe.Pointer, spacing int) {
	c_box := (*C.GtkBox)(unsafe.Pointer(box))

	c_spacing := (C.gint)(spacing)

	C.gtk_box_set_spacing(c_box, c_spacing)
}

// UNSUPPORTED : gtk_builder_add_callback_symbol : parameter 'callback_symbol' is callback

// UNSUPPORTED : gtk_builder_add_callback_symbols : parameter 'first_callback_symbol' is callback

// UNSUPPORTED : gtk_builder_connect_signals_full : parameter 'func' is callback

func Fn_gtk_builder_extend_with_template(builder unsafe.Pointer, widget unsafe.Pointer, templateType uint64, buffer string, length uint64, error unsafe.Pointer) uint {
	c_builder := (*C.GtkBuilder)(unsafe.Pointer(builder))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_templateType := (C.GType)(templateType)

	c_buffer := (*C.gchar)(C.CString(buffer))
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gsize)(length)

	cError := (**C.GError)(error)

	ret := C.gtk_builder_extend_with_template(c_builder, c_widget, c_templateType, c_buffer, c_length, cError)

	return (uint)(ret)
}

// UNSUPPORTED : gtk_builder_lookup_callback_symbol : has callback return

func Fn_gtk_button_new() unsafe.Pointer {
	ret := C.gtk_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_from_stock(stockId string) unsafe.Pointer {
	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	ret := C.gtk_button_new_from_stock(c_stockId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_button_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_button_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_button_clicked(button unsafe.Pointer) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	C.gtk_button_clicked(c_button)
}

func Fn_gtk_button_enter(button unsafe.Pointer) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	C.gtk_button_enter(c_button)
}

func Fn_gtk_button_get_label(button unsafe.Pointer) string {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	ret := C.gtk_button_get_label(c_button)

	return C.GoString(ret)
}

func Fn_gtk_button_get_relief(button unsafe.Pointer) int {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	ret := C.gtk_button_get_relief(c_button)

	return (int)(ret)
}

func Fn_gtk_button_get_use_stock(button unsafe.Pointer) bool {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	ret := C.gtk_button_get_use_stock(c_button)

	return toGoBool(ret)
}

func Fn_gtk_button_get_use_underline(button unsafe.Pointer) bool {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	ret := C.gtk_button_get_use_underline(c_button)

	return toGoBool(ret)
}

func Fn_gtk_button_leave(button unsafe.Pointer) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	C.gtk_button_leave(c_button)
}

func Fn_gtk_button_pressed(button unsafe.Pointer) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	C.gtk_button_pressed(c_button)
}

func Fn_gtk_button_released(button unsafe.Pointer) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	C.gtk_button_released(c_button)
}

func Fn_gtk_button_set_label(button unsafe.Pointer, label string) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_button_set_label(c_button, c_label)
}

func Fn_gtk_button_set_relief(button unsafe.Pointer, relief int) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	c_relief := (C.GtkReliefStyle)(relief)

	C.gtk_button_set_relief(c_button, c_relief)
}

func Fn_gtk_button_set_use_stock(button unsafe.Pointer, useStock bool) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	c_useStock := toCBool(useStock)

	C.gtk_button_set_use_stock(c_button, c_useStock)
}

func Fn_gtk_button_set_use_underline(button unsafe.Pointer, useUnderline bool) {
	c_button := (*C.GtkButton)(unsafe.Pointer(button))

	c_useUnderline := toCBool(useUnderline)

	C.gtk_button_set_use_underline(c_button, c_useUnderline)
}

func Fn_gtk_button_box_get_layout(widget unsafe.Pointer) int {
	c_widget := (*C.GtkButtonBox)(unsafe.Pointer(widget))

	ret := C.gtk_button_box_get_layout(c_widget)

	return (int)(ret)
}

func Fn_gtk_button_box_set_child_secondary(widget unsafe.Pointer, child unsafe.Pointer, isSecondary bool) {
	c_widget := (*C.GtkButtonBox)(unsafe.Pointer(widget))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_isSecondary := toCBool(isSecondary)

	C.gtk_button_box_set_child_secondary(c_widget, c_child, c_isSecondary)
}

func Fn_gtk_button_box_set_layout(widget unsafe.Pointer, layoutStyle int) {
	c_widget := (*C.GtkButtonBox)(unsafe.Pointer(widget))

	c_layoutStyle := (C.GtkButtonBoxStyle)(layoutStyle)

	C.gtk_button_box_set_layout(c_widget, c_layoutStyle)
}

func Fn_gtk_calendar_new() unsafe.Pointer {
	ret := C.gtk_calendar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_calendar_clear_marks(calendar unsafe.Pointer) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	C.gtk_calendar_clear_marks(c_calendar)
}

func Fn_gtk_calendar_get_date(calendar unsafe.Pointer, year *uint, month *uint, day *uint) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	c_year := (*C.guint)(unsafe.Pointer(year))

	c_month := (*C.guint)(unsafe.Pointer(month))

	c_day := (*C.guint)(unsafe.Pointer(day))

	C.gtk_calendar_get_date(c_calendar, c_year, c_month, c_day)
}

func Fn_gtk_calendar_mark_day(calendar unsafe.Pointer, day uint) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	c_day := (C.guint)(day)

	C.gtk_calendar_mark_day(c_calendar, c_day)
}

func Fn_gtk_calendar_select_day(calendar unsafe.Pointer, day uint) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	c_day := (C.guint)(day)

	C.gtk_calendar_select_day(c_calendar, c_day)
}

func Fn_gtk_calendar_select_month(calendar unsafe.Pointer, month uint, year uint) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	c_month := (C.guint)(month)

	c_year := (C.guint)(year)

	C.gtk_calendar_select_month(c_calendar, c_month, c_year)
}

// UNSUPPORTED : gtk_calendar_set_detail_func : parameter 'func' is callback

func Fn_gtk_calendar_unmark_day(calendar unsafe.Pointer, day uint) {
	c_calendar := (*C.GtkCalendar)(unsafe.Pointer(calendar))

	c_day := (C.guint)(day)

	C.gtk_calendar_unmark_day(c_calendar, c_day)
}

// UNSUPPORTED : gtk_cell_area_foreach : parameter 'callback' is callback

// UNSUPPORTED : gtk_cell_area_foreach_alloc : parameter 'callback' is callback

func Fn_gtk_cell_area_context_allocate(context unsafe.Pointer, width int, height int) {
	c_context := (*C.GtkCellAreaContext)(unsafe.Pointer(context))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_cell_area_context_allocate(c_context, c_width, c_height)
}

func Fn_gtk_cell_area_context_reset(context unsafe.Pointer) {
	c_context := (*C.GtkCellAreaContext)(unsafe.Pointer(context))

	C.gtk_cell_area_context_reset(c_context)
}

func Fn_gtk_cell_renderer_activate(cell unsafe.Pointer, event unsafe.Pointer, widget unsafe.Pointer, path string, backgroundArea unsafe.Pointer, cellArea unsafe.Pointer, flags int) bool {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	c_backgroundArea := (*C.GdkRectangle)(unsafe.Pointer(backgroundArea))

	c_cellArea := (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	c_flags := (C.GtkCellRendererState)(flags)

	ret := C.gtk_cell_renderer_activate(c_cell, c_event, c_widget, c_path, c_backgroundArea, c_cellArea, c_flags)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_get_fixed_size(cell unsafe.Pointer, width *int, height *int) {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_cell_renderer_get_fixed_size(c_cell, c_width, c_height)
}

func Fn_gtk_cell_renderer_get_size(cell unsafe.Pointer, widget unsafe.Pointer, cellArea unsafe.Pointer, xOffset *int, yOffset *int, width *int, height *int) {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_cellArea := (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	c_xOffset := (*C.gint)(unsafe.Pointer(xOffset))

	c_yOffset := (*C.gint)(unsafe.Pointer(yOffset))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_cell_renderer_get_size(c_cell, c_widget, c_cellArea, c_xOffset, c_yOffset, c_width, c_height)
}

func Fn_gtk_cell_renderer_render(cell unsafe.Pointer, cr unsafe.Pointer, widget unsafe.Pointer, backgroundArea unsafe.Pointer, cellArea unsafe.Pointer, flags int) {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_backgroundArea := (*C.GdkRectangle)(unsafe.Pointer(backgroundArea))

	c_cellArea := (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	c_flags := (C.GtkCellRendererState)(flags)

	C.gtk_cell_renderer_render(c_cell, c_cr, c_widget, c_backgroundArea, c_cellArea, c_flags)
}

func Fn_gtk_cell_renderer_set_fixed_size(cell unsafe.Pointer, width int, height int) {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_cell_renderer_set_fixed_size(c_cell, c_width, c_height)
}

func Fn_gtk_cell_renderer_start_editing(cell unsafe.Pointer, event unsafe.Pointer, widget unsafe.Pointer, path string, backgroundArea unsafe.Pointer, cellArea unsafe.Pointer, flags int) unsafe.Pointer {
	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	c_backgroundArea := (*C.GdkRectangle)(unsafe.Pointer(backgroundArea))

	c_cellArea := (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	c_flags := (C.GtkCellRendererState)(flags)

	ret := C.gtk_cell_renderer_start_editing(c_cell, c_event, c_widget, c_path, c_backgroundArea, c_cellArea, c_flags)

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_pixbuf_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_pixbuf_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_text_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_text_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_text_set_fixed_height_from_font(renderer unsafe.Pointer, numberOfRows int) {
	c_renderer := (*C.GtkCellRendererText)(unsafe.Pointer(renderer))

	c_numberOfRows := (C.gint)(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(c_renderer, c_numberOfRows)
}

func Fn_gtk_cell_renderer_toggle_new() unsafe.Pointer {
	ret := C.gtk_cell_renderer_toggle_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_cell_renderer_toggle_get_active(toggle unsafe.Pointer) bool {
	c_toggle := (*C.GtkCellRendererToggle)(unsafe.Pointer(toggle))

	ret := C.gtk_cell_renderer_toggle_get_active(c_toggle)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_toggle_get_radio(toggle unsafe.Pointer) bool {
	c_toggle := (*C.GtkCellRendererToggle)(unsafe.Pointer(toggle))

	ret := C.gtk_cell_renderer_toggle_get_radio(c_toggle)

	return toGoBool(ret)
}

func Fn_gtk_cell_renderer_toggle_set_active(toggle unsafe.Pointer, setting bool) {
	c_toggle := (*C.GtkCellRendererToggle)(unsafe.Pointer(toggle))

	c_setting := toCBool(setting)

	C.gtk_cell_renderer_toggle_set_active(c_toggle, c_setting)
}

func Fn_gtk_cell_renderer_toggle_set_radio(toggle unsafe.Pointer, radio bool) {
	c_toggle := (*C.GtkCellRendererToggle)(unsafe.Pointer(toggle))

	c_radio := toCBool(radio)

	C.gtk_cell_renderer_toggle_set_radio(c_toggle, c_radio)
}

func Fn_gtk_check_button_new() unsafe.Pointer {
	ret := C.gtk_check_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_button_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_check_button_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_button_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_check_button_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new() unsafe.Pointer {
	ret := C.gtk_check_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_check_menu_item_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_check_menu_item_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_check_menu_item_get_active(checkMenuItem unsafe.Pointer) bool {
	c_checkMenuItem := (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem))

	ret := C.gtk_check_menu_item_get_active(c_checkMenuItem)

	return toGoBool(ret)
}

func Fn_gtk_check_menu_item_get_inconsistent(checkMenuItem unsafe.Pointer) bool {
	c_checkMenuItem := (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem))

	ret := C.gtk_check_menu_item_get_inconsistent(c_checkMenuItem)

	return toGoBool(ret)
}

func Fn_gtk_check_menu_item_set_active(checkMenuItem unsafe.Pointer, isActive bool) {
	c_checkMenuItem := (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem))

	c_isActive := toCBool(isActive)

	C.gtk_check_menu_item_set_active(c_checkMenuItem, c_isActive)
}

func Fn_gtk_check_menu_item_set_inconsistent(checkMenuItem unsafe.Pointer, setting bool) {
	c_checkMenuItem := (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem))

	c_setting := toCBool(setting)

	C.gtk_check_menu_item_set_inconsistent(c_checkMenuItem, c_setting)
}

func Fn_gtk_check_menu_item_toggled(checkMenuItem unsafe.Pointer) {
	c_checkMenuItem := (*C.GtkCheckMenuItem)(unsafe.Pointer(checkMenuItem))

	C.gtk_check_menu_item_toggled(c_checkMenuItem)
}

func Fn_gtk_clipboard_clear(clipboard unsafe.Pointer) {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	C.gtk_clipboard_clear(c_clipboard)
}

func Fn_gtk_clipboard_get_display(clipboard unsafe.Pointer) unsafe.Pointer {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	ret := C.gtk_clipboard_get_display(c_clipboard)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_get_owner(clipboard unsafe.Pointer) unsafe.Pointer {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	ret := C.gtk_clipboard_get_owner(c_clipboard)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_clipboard_request_contents : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_image : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_rich_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_targets : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_text : parameter 'callback' is callback

// UNSUPPORTED : gtk_clipboard_request_uris : parameter 'callback' is callback

func Fn_gtk_clipboard_set_text(clipboard unsafe.Pointer, text string, len_ int) {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	C.gtk_clipboard_set_text(c_clipboard, c_text, c_len_)
}

// UNSUPPORTED : gtk_clipboard_set_with_data : parameter 'get_func' is callback

// UNSUPPORTED : gtk_clipboard_set_with_owner : parameter 'get_func' is callback

func Fn_gtk_clipboard_wait_for_contents(clipboard unsafe.Pointer, target unsafe.Pointer) unsafe.Pointer {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	c_target := (C.GdkAtom)(target)

	ret := C.gtk_clipboard_wait_for_contents(c_clipboard, c_target)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_wait_for_text(clipboard unsafe.Pointer) string {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	ret := C.gtk_clipboard_wait_for_text(c_clipboard)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_clipboard_wait_for_uris : no array length

func Fn_gtk_clipboard_wait_is_text_available(clipboard unsafe.Pointer) bool {
	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	ret := C.gtk_clipboard_wait_is_text_available(c_clipboard)

	return toGoBool(ret)
}

func Fn_gtk_clipboard_get(selection unsafe.Pointer) unsafe.Pointer {
	c_selection := (C.GdkAtom)(selection)

	ret := C.gtk_clipboard_get(c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gtk_clipboard_get_for_display(display unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_selection := (C.GdkAtom)(selection)

	ret := C.gtk_clipboard_get_for_display(c_display, c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_selection_new() unsafe.Pointer {
	ret := C.gtk_color_selection_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_color_selection_get_current_alpha(colorsel unsafe.Pointer) uint16 {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	ret := C.gtk_color_selection_get_current_alpha(c_colorsel)

	return (uint16)(ret)
}

func Fn_gtk_color_selection_get_current_color(colorsel unsafe.Pointer, color unsafe.Pointer) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_color_selection_get_current_color(c_colorsel, c_color)
}

func Fn_gtk_color_selection_get_has_opacity_control(colorsel unsafe.Pointer) bool {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	ret := C.gtk_color_selection_get_has_opacity_control(c_colorsel)

	return toGoBool(ret)
}

func Fn_gtk_color_selection_get_has_palette(colorsel unsafe.Pointer) bool {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	ret := C.gtk_color_selection_get_has_palette(c_colorsel)

	return toGoBool(ret)
}

func Fn_gtk_color_selection_get_previous_alpha(colorsel unsafe.Pointer) uint16 {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	ret := C.gtk_color_selection_get_previous_alpha(c_colorsel)

	return (uint16)(ret)
}

func Fn_gtk_color_selection_get_previous_color(colorsel unsafe.Pointer, color unsafe.Pointer) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_color_selection_get_previous_color(c_colorsel, c_color)
}

func Fn_gtk_color_selection_is_adjusting(colorsel unsafe.Pointer) bool {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	ret := C.gtk_color_selection_is_adjusting(c_colorsel)

	return toGoBool(ret)
}

func Fn_gtk_color_selection_set_current_alpha(colorsel unsafe.Pointer, alpha uint16) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_alpha := (C.guint16)(alpha)

	C.gtk_color_selection_set_current_alpha(c_colorsel, c_alpha)
}

func Fn_gtk_color_selection_set_current_color(colorsel unsafe.Pointer, color unsafe.Pointer) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_color_selection_set_current_color(c_colorsel, c_color)
}

func Fn_gtk_color_selection_set_has_opacity_control(colorsel unsafe.Pointer, hasOpacity bool) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_hasOpacity := toCBool(hasOpacity)

	C.gtk_color_selection_set_has_opacity_control(c_colorsel, c_hasOpacity)
}

func Fn_gtk_color_selection_set_has_palette(colorsel unsafe.Pointer, hasPalette bool) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_hasPalette := toCBool(hasPalette)

	C.gtk_color_selection_set_has_palette(c_colorsel, c_hasPalette)
}

func Fn_gtk_color_selection_set_previous_alpha(colorsel unsafe.Pointer, alpha uint16) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_alpha := (C.guint16)(alpha)

	C.gtk_color_selection_set_previous_alpha(c_colorsel, c_alpha)
}

func Fn_gtk_color_selection_set_previous_color(colorsel unsafe.Pointer, color unsafe.Pointer) {
	c_colorsel := (*C.GtkColorSelection)(unsafe.Pointer(colorsel))

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_color_selection_set_previous_color(c_colorsel, c_color)
}

func Fn_gtk_color_selection_palette_from_string(str string, colors *[]unsafe.Pointer, nColors *int) bool {
	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	var c_colorsArrayPointer (*C.GdkColor)
	c_colors := &c_colorsArrayPointer

	c_nColors := (*C.gint)(unsafe.Pointer(nColors))

	ret := C.gtk_color_selection_palette_from_string(c_str, c_colors, c_nColors)

	colorsOutLen := int(*c_nColors)
	colorsOut := make([]unsafe.Pointer, colorsOutLen, colorsOutLen)
	if colorsOutLen > 0 {
		colorsOut = (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(c_colorsArrayPointer))[:colorsOutLen:colorsOutLen]
	}
	*colors = colorsOut

	return toGoBool(ret)
}

func Fn_gtk_color_selection_palette_to_string(colors []gdk.Color, nColors int) string {
	c_colors := (*C.GdkColor)(unsafe.Pointer(&colors[0]))

	c_nColors := (C.gint)(nColors)

	ret := C.gtk_color_selection_palette_to_string(c_colors, c_nColors)

	return C.GoString(ret)
}

// UNSUPPORTED : gtk_color_selection_set_change_palette_with_screen_hook : parameter 'func' is callback

func Fn_gtk_color_selection_dialog_new(title string) unsafe.Pointer {
	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	ret := C.gtk_color_selection_dialog_new(c_title)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_area(area unsafe.Pointer) unsafe.Pointer {
	c_area := (*C.GtkCellArea)(unsafe.Pointer(area))

	ret := C.gtk_combo_box_new_with_area(c_area)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_new_with_area_and_entry(area unsafe.Pointer) unsafe.Pointer {
	c_area := (*C.GtkCellArea)(unsafe.Pointer(area))

	ret := C.gtk_combo_box_new_with_area_and_entry(c_area)

	return unsafe.Pointer(ret)
}

func Fn_gtk_combo_box_get_add_tearoffs(comboBox unsafe.Pointer) bool {
	c_comboBox := (*C.GtkComboBox)(unsafe.Pointer(comboBox))

	ret := C.gtk_combo_box_get_add_tearoffs(c_comboBox)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_combo_box_get_row_separator_func : has callback return

// UNSUPPORTED : gtk_combo_box_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_container_add(container unsafe.Pointer, widget unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_container_add(c_container, c_widget)
}

func Fn_gtk_container_add_with_properties(container unsafe.Pointer, widget unsafe.Pointer, firstPropName string) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_firstPropName := (*C.gchar)(C.CString(firstPropName))
	defer C.free(unsafe.Pointer(c_firstPropName))

	C.c_gtk_container_add_with_properties(c_container, c_widget, c_firstPropName)
}

func Fn_gtk_container_check_resize(container unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	C.gtk_container_check_resize(c_container)
}

func Fn_gtk_container_child_get(container unsafe.Pointer, child unsafe.Pointer, firstPropName string) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_firstPropName := (*C.gchar)(C.CString(firstPropName))
	defer C.free(unsafe.Pointer(c_firstPropName))

	C.c_gtk_container_child_get(c_container, c_child, c_firstPropName)
}

func Fn_gtk_container_child_get_property(container unsafe.Pointer, child unsafe.Pointer, propertyName string, value unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_container_child_get_property(c_container, c_child, c_propertyName, c_value)
}

func Fn_gtk_container_child_get_valist(container unsafe.Pointer, child unsafe.Pointer, firstPropertyName string) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_firstPropertyName := (*C.gchar)(C.CString(firstPropertyName))
	defer C.free(unsafe.Pointer(c_firstPropertyName))

	C.c_gtk_container_child_get_valist(c_container, c_child, c_firstPropertyName)
}

func Fn_gtk_container_child_set(container unsafe.Pointer, child unsafe.Pointer, firstPropName string) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_firstPropName := (*C.gchar)(C.CString(firstPropName))
	defer C.free(unsafe.Pointer(c_firstPropName))

	C.c_gtk_container_child_set(c_container, c_child, c_firstPropName)
}

func Fn_gtk_container_child_set_property(container unsafe.Pointer, child unsafe.Pointer, propertyName string, value unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_container_child_set_property(c_container, c_child, c_propertyName, c_value)
}

func Fn_gtk_container_child_set_valist(container unsafe.Pointer, child unsafe.Pointer, firstPropertyName string) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_firstPropertyName := (*C.gchar)(C.CString(firstPropertyName))
	defer C.free(unsafe.Pointer(c_firstPropertyName))

	C.c_gtk_container_child_set_valist(c_container, c_child, c_firstPropertyName)
}

func Fn_gtk_container_child_type(container unsafe.Pointer) uint64 {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_child_type(c_container)

	return (uint64)(ret)
}

// UNSUPPORTED : gtk_container_forall : parameter 'callback' is callback

// UNSUPPORTED : gtk_container_foreach : parameter 'callback' is callback

func Fn_gtk_container_get_border_width(container unsafe.Pointer) uint {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_get_border_width(c_container)

	return (uint)(ret)
}

func Fn_gtk_container_get_children(container unsafe.Pointer) unsafe.Pointer {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_get_children(c_container)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_container_get_focus_chain : parameter 'focusable_widgets' is non array with indirect count > 1

func Fn_gtk_container_get_focus_hadjustment(container unsafe.Pointer) unsafe.Pointer {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_get_focus_hadjustment(c_container)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_focus_vadjustment(container unsafe.Pointer) unsafe.Pointer {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_get_focus_vadjustment(c_container)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_path_for_child(container unsafe.Pointer, child unsafe.Pointer) unsafe.Pointer {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_container_get_path_for_child(c_container, c_child)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_get_resize_mode(container unsafe.Pointer) int {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	ret := C.gtk_container_get_resize_mode(c_container)

	return (int)(ret)
}

func Fn_gtk_container_propagate_draw(container unsafe.Pointer, child unsafe.Pointer, cr unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	C.gtk_container_propagate_draw(c_container, c_child, c_cr)
}

func Fn_gtk_container_remove(container unsafe.Pointer, widget unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_container_remove(c_container, c_widget)
}

func Fn_gtk_container_resize_children(container unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	C.gtk_container_resize_children(c_container)
}

func Fn_gtk_container_set_border_width(container unsafe.Pointer, borderWidth uint) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_borderWidth := (C.guint)(borderWidth)

	C.gtk_container_set_border_width(c_container, c_borderWidth)
}

func Fn_gtk_container_set_focus_chain(container unsafe.Pointer, focusableWidgets unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_focusableWidgets := (*C.GList)(unsafe.Pointer(focusableWidgets))

	C.gtk_container_set_focus_chain(c_container, c_focusableWidgets)
}

func Fn_gtk_container_set_focus_child(container unsafe.Pointer, child unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_container_set_focus_child(c_container, c_child)
}

func Fn_gtk_container_set_focus_hadjustment(container unsafe.Pointer, adjustment unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_container_set_focus_hadjustment(c_container, c_adjustment)
}

func Fn_gtk_container_set_focus_vadjustment(container unsafe.Pointer, adjustment unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_container_set_focus_vadjustment(c_container, c_adjustment)
}

func Fn_gtk_container_set_reallocate_redraws(container unsafe.Pointer, needsRedraws bool) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_needsRedraws := toCBool(needsRedraws)

	C.gtk_container_set_reallocate_redraws(c_container, c_needsRedraws)
}

func Fn_gtk_container_set_resize_mode(container unsafe.Pointer, resizeMode int) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	c_resizeMode := (C.GtkResizeMode)(resizeMode)

	C.gtk_container_set_resize_mode(c_container, c_resizeMode)
}

func Fn_gtk_container_unset_focus_chain(container unsafe.Pointer) {
	c_container := (*C.GtkContainer)(unsafe.Pointer(container))

	C.gtk_container_unset_focus_chain(c_container)
}

func Fn_gtk_container_cell_accessible_new() unsafe.Pointer {
	ret := C.gtk_container_cell_accessible_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_cell_accessible_add_child(container unsafe.Pointer, child unsafe.Pointer) {
	c_container := (*C.GtkContainerCellAccessible)(unsafe.Pointer(container))

	c_child := (*C.GtkCellAccessible)(unsafe.Pointer(child))

	C.gtk_container_cell_accessible_add_child(c_container, c_child)
}

func Fn_gtk_container_cell_accessible_get_children(container unsafe.Pointer) unsafe.Pointer {
	c_container := (*C.GtkContainerCellAccessible)(unsafe.Pointer(container))

	ret := C.gtk_container_cell_accessible_get_children(c_container)

	return unsafe.Pointer(ret)
}

func Fn_gtk_container_cell_accessible_remove_child(container unsafe.Pointer, child unsafe.Pointer) {
	c_container := (*C.GtkContainerCellAccessible)(unsafe.Pointer(container))

	c_child := (*C.GtkCellAccessible)(unsafe.Pointer(child))

	C.gtk_container_cell_accessible_remove_child(c_container, c_child)
}

func Fn_gtk_css_provider_new() unsafe.Pointer {
	ret := C.gtk_css_provider_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_provider_load_from_data(cssProvider unsafe.Pointer, data []uint8, length uint64, error unsafe.Pointer) bool {
	c_cssProvider := (*C.GtkCssProvider)(unsafe.Pointer(cssProvider))

	c_data := (*C.gchar)(unsafe.Pointer(&data[0]))

	c_length := (C.gssize)(length)

	cError := (**C.GError)(error)

	ret := C.gtk_css_provider_load_from_data(c_cssProvider, c_data, c_length, cError)

	return toGoBool(ret)
}

func Fn_gtk_css_provider_load_from_file(cssProvider unsafe.Pointer, file unsafe.Pointer, error unsafe.Pointer) bool {
	c_cssProvider := (*C.GtkCssProvider)(unsafe.Pointer(cssProvider))

	c_file := (*C.GFile)(unsafe.Pointer(file))

	cError := (**C.GError)(error)

	ret := C.gtk_css_provider_load_from_file(c_cssProvider, c_file, cError)

	return toGoBool(ret)
}

func Fn_gtk_css_provider_load_from_path(cssProvider unsafe.Pointer, path string, error unsafe.Pointer) bool {
	c_cssProvider := (*C.GtkCssProvider)(unsafe.Pointer(cssProvider))

	c_path := (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(c_path))

	cError := (**C.GError)(error)

	ret := C.gtk_css_provider_load_from_path(c_cssProvider, c_path, cError)

	return toGoBool(ret)
}

func Fn_gtk_css_provider_get_default() unsafe.Pointer {
	ret := C.gtk_css_provider_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_css_provider_get_named(name string, variant *string) unsafe.Pointer {
	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	var c_variantValue (*C.gchar)
	if variant != nil {
		c_variantValue = (*C.gchar)(C.CString(*variant))
		defer C.free(unsafe.Pointer(c_variantValue))
	}
	c_variant := c_variantValue

	ret := C.gtk_css_provider_get_named(c_name, c_variant)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_new() unsafe.Pointer {
	ret := C.gtk_dialog_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_new_with_buttons(title *string, parent unsafe.Pointer, flags int, firstButtonText *string) unsafe.Pointer {
	var c_titleValue (*C.gchar)
	if title != nil {
		c_titleValue = (*C.gchar)(C.CString(*title))
		defer C.free(unsafe.Pointer(c_titleValue))
	}
	c_title := c_titleValue

	c_parent := (*C.GtkWindow)(unsafe.Pointer(parent))

	c_flags := (C.GtkDialogFlags)(flags)

	var c_firstButtonTextValue (*C.gchar)
	if firstButtonText != nil {
		c_firstButtonTextValue = (*C.gchar)(C.CString(*firstButtonText))
		defer C.free(unsafe.Pointer(c_firstButtonTextValue))
	}
	c_firstButtonText := c_firstButtonTextValue

	ret := C.c_gtk_dialog_new_with_buttons(c_title, c_parent, c_flags, c_firstButtonText)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_add_action_widget(dialog unsafe.Pointer, child unsafe.Pointer, responseId int) {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_responseId := (C.gint)(responseId)

	C.gtk_dialog_add_action_widget(c_dialog, c_child, c_responseId)
}

func Fn_gtk_dialog_add_button(dialog unsafe.Pointer, buttonText string, responseId int) unsafe.Pointer {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_buttonText := (*C.gchar)(C.CString(buttonText))
	defer C.free(unsafe.Pointer(c_buttonText))

	c_responseId := (C.gint)(responseId)

	ret := C.gtk_dialog_add_button(c_dialog, c_buttonText, c_responseId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_dialog_add_buttons(dialog unsafe.Pointer, firstButtonText string) {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_firstButtonText := (*C.gchar)(C.CString(firstButtonText))
	defer C.free(unsafe.Pointer(c_firstButtonText))

	C.c_gtk_dialog_add_buttons(c_dialog, c_firstButtonText)
}

func Fn_gtk_dialog_response(dialog unsafe.Pointer, responseId int) {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_responseId := (C.gint)(responseId)

	C.gtk_dialog_response(c_dialog, c_responseId)
}

func Fn_gtk_dialog_run(dialog unsafe.Pointer) int {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	ret := C.gtk_dialog_run(c_dialog)

	return (int)(ret)
}

func Fn_gtk_dialog_set_default_response(dialog unsafe.Pointer, responseId int) {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_responseId := (C.gint)(responseId)

	C.gtk_dialog_set_default_response(c_dialog, c_responseId)
}

func Fn_gtk_dialog_set_response_sensitive(dialog unsafe.Pointer, responseId int, setting bool) {
	c_dialog := (*C.GtkDialog)(unsafe.Pointer(dialog))

	c_responseId := (C.gint)(responseId)

	c_setting := toCBool(setting)

	C.gtk_dialog_set_response_sensitive(c_dialog, c_responseId, c_setting)
}

func Fn_gtk_drawing_area_new() unsafe.Pointer {
	ret := C.gtk_drawing_area_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_new() unsafe.Pointer {
	ret := C.gtk_entry_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_activates_default(entry unsafe.Pointer) bool {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_activates_default(c_entry)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_has_frame(entry unsafe.Pointer) bool {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_has_frame(c_entry)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_invisible_char(entry unsafe.Pointer) rune {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_invisible_char(c_entry)

	return (rune)(ret)
}

func Fn_gtk_entry_get_layout(entry unsafe.Pointer) unsafe.Pointer {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_layout(c_entry)

	return unsafe.Pointer(ret)
}

func Fn_gtk_entry_get_layout_offsets(entry unsafe.Pointer, x *int, y *int) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	C.gtk_entry_get_layout_offsets(c_entry, c_x, c_y)
}

func Fn_gtk_entry_get_max_length(entry unsafe.Pointer) int {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_max_length(c_entry)

	return (int)(ret)
}

func Fn_gtk_entry_get_text(entry unsafe.Pointer) string {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_text(c_entry)

	return C.GoString(ret)
}

func Fn_gtk_entry_get_visibility(entry unsafe.Pointer) bool {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_visibility(c_entry)

	return toGoBool(ret)
}

func Fn_gtk_entry_get_width_chars(entry unsafe.Pointer) int {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	ret := C.gtk_entry_get_width_chars(c_entry)

	return (int)(ret)
}

func Fn_gtk_entry_layout_index_to_text_index(entry unsafe.Pointer, layoutIndex int) int {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_layoutIndex := (C.gint)(layoutIndex)

	ret := C.gtk_entry_layout_index_to_text_index(c_entry, c_layoutIndex)

	return (int)(ret)
}

func Fn_gtk_entry_set_activates_default(entry unsafe.Pointer, setting bool) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_setting := toCBool(setting)

	C.gtk_entry_set_activates_default(c_entry, c_setting)
}

func Fn_gtk_entry_set_has_frame(entry unsafe.Pointer, setting bool) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_setting := toCBool(setting)

	C.gtk_entry_set_has_frame(c_entry, c_setting)
}

func Fn_gtk_entry_set_invisible_char(entry unsafe.Pointer, ch rune) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_ch := (C.gunichar)(ch)

	C.gtk_entry_set_invisible_char(c_entry, c_ch)
}

func Fn_gtk_entry_set_max_length(entry unsafe.Pointer, max int) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_max := (C.gint)(max)

	C.gtk_entry_set_max_length(c_entry, c_max)
}

func Fn_gtk_entry_set_text(entry unsafe.Pointer, text string) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_set_text(c_entry, c_text)
}

func Fn_gtk_entry_set_visibility(entry unsafe.Pointer, visible bool) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_visible := toCBool(visible)

	C.gtk_entry_set_visibility(c_entry, c_visible)
}

func Fn_gtk_entry_set_width_chars(entry unsafe.Pointer, nChars int) {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_nChars := (C.gint)(nChars)

	C.gtk_entry_set_width_chars(c_entry, c_nChars)
}

func Fn_gtk_entry_text_index_to_layout_index(entry unsafe.Pointer, textIndex int) int {
	c_entry := (*C.GtkEntry)(unsafe.Pointer(entry))

	c_textIndex := (C.gint)(textIndex)

	ret := C.gtk_entry_text_index_to_layout_index(c_entry, c_textIndex)

	return (int)(ret)
}

// UNSUPPORTED : gtk_entry_completion_set_match_func : parameter 'func' is callback

func Fn_gtk_event_box_new() unsafe.Pointer {
	ret := C.gtk_event_box_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_event_controller_key_new : blacklisted

// UNSUPPORTED : gtk_event_controller_key_forward : blacklisted

// UNSUPPORTED : gtk_event_controller_key_get_group : blacklisted

// UNSUPPORTED : gtk_event_controller_key_set_im_context : blacklisted

// UNSUPPORTED : gtk_file_filter_add_custom : parameter 'func' is callback

func Fn_gtk_fixed_new() unsafe.Pointer {
	ret := C.gtk_fixed_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_fixed_move(fixed unsafe.Pointer, widget unsafe.Pointer, x int, y int) {
	c_fixed := (*C.GtkFixed)(unsafe.Pointer(fixed))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_fixed_move(c_fixed, c_widget, c_x, c_y)
}

func Fn_gtk_fixed_put(fixed unsafe.Pointer, widget unsafe.Pointer, x int, y int) {
	c_fixed := (*C.GtkFixed)(unsafe.Pointer(fixed))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_fixed_put(c_fixed, c_widget, c_x, c_y)
}

// UNSUPPORTED : gtk_flow_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_flow_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_flow_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_flow_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_font_selection_new() unsafe.Pointer {
	ret := C.gtk_font_selection_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_get_font_name(fontsel unsafe.Pointer) string {
	c_fontsel := (*C.GtkFontSelection)(unsafe.Pointer(fontsel))

	ret := C.gtk_font_selection_get_font_name(c_fontsel)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_get_preview_text(fontsel unsafe.Pointer) string {
	c_fontsel := (*C.GtkFontSelection)(unsafe.Pointer(fontsel))

	ret := C.gtk_font_selection_get_preview_text(c_fontsel)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_set_font_name(fontsel unsafe.Pointer, fontname string) bool {
	c_fontsel := (*C.GtkFontSelection)(unsafe.Pointer(fontsel))

	c_fontname := (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(c_fontname))

	ret := C.gtk_font_selection_set_font_name(c_fontsel, c_fontname)

	return toGoBool(ret)
}

func Fn_gtk_font_selection_set_preview_text(fontsel unsafe.Pointer, text string) {
	c_fontsel := (*C.GtkFontSelection)(unsafe.Pointer(fontsel))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_font_selection_set_preview_text(c_fontsel, c_text)
}

func Fn_gtk_font_selection_dialog_new(title string) unsafe.Pointer {
	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	ret := C.gtk_font_selection_dialog_new(c_title)

	return unsafe.Pointer(ret)
}

func Fn_gtk_font_selection_dialog_get_font_name(fsd unsafe.Pointer) string {
	c_fsd := (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd))

	ret := C.gtk_font_selection_dialog_get_font_name(c_fsd)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_dialog_get_preview_text(fsd unsafe.Pointer) string {
	c_fsd := (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd))

	ret := C.gtk_font_selection_dialog_get_preview_text(c_fsd)

	return C.GoString(ret)
}

func Fn_gtk_font_selection_dialog_set_font_name(fsd unsafe.Pointer, fontname string) bool {
	c_fsd := (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd))

	c_fontname := (*C.gchar)(C.CString(fontname))
	defer C.free(unsafe.Pointer(c_fontname))

	ret := C.gtk_font_selection_dialog_set_font_name(c_fsd, c_fontname)

	return toGoBool(ret)
}

func Fn_gtk_font_selection_dialog_set_preview_text(fsd unsafe.Pointer, text string) {
	c_fsd := (*C.GtkFontSelectionDialog)(unsafe.Pointer(fsd))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_font_selection_dialog_set_preview_text(c_fsd, c_text)
}

func Fn_gtk_frame_new(label *string) unsafe.Pointer {
	var c_labelValue (*C.gchar)
	if label != nil {
		c_labelValue = (*C.gchar)(C.CString(*label))
		defer C.free(unsafe.Pointer(c_labelValue))
	}
	c_label := c_labelValue

	ret := C.gtk_frame_new(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_frame_get_label(frame unsafe.Pointer) string {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	ret := C.gtk_frame_get_label(c_frame)

	return C.GoString(ret)
}

func Fn_gtk_frame_get_label_align(frame unsafe.Pointer, xalign *float32, yalign *float32) {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	c_xalign := (*C.gfloat)(unsafe.Pointer(xalign))

	c_yalign := (*C.gfloat)(unsafe.Pointer(yalign))

	C.gtk_frame_get_label_align(c_frame, c_xalign, c_yalign)
}

func Fn_gtk_frame_get_label_widget(frame unsafe.Pointer) unsafe.Pointer {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	ret := C.gtk_frame_get_label_widget(c_frame)

	return unsafe.Pointer(ret)
}

func Fn_gtk_frame_get_shadow_type(frame unsafe.Pointer) int {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	ret := C.gtk_frame_get_shadow_type(c_frame)

	return (int)(ret)
}

func Fn_gtk_frame_set_label(frame unsafe.Pointer, label *string) {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	var c_labelValue (*C.gchar)
	if label != nil {
		c_labelValue = (*C.gchar)(C.CString(*label))
		defer C.free(unsafe.Pointer(c_labelValue))
	}
	c_label := c_labelValue

	C.gtk_frame_set_label(c_frame, c_label)
}

func Fn_gtk_frame_set_label_align(frame unsafe.Pointer, xalign float32, yalign float32) {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_frame_set_label_align(c_frame, c_xalign, c_yalign)
}

func Fn_gtk_frame_set_label_widget(frame unsafe.Pointer, labelWidget unsafe.Pointer) {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	c_labelWidget := (*C.GtkWidget)(unsafe.Pointer(labelWidget))

	C.gtk_frame_set_label_widget(c_frame, c_labelWidget)
}

func Fn_gtk_frame_set_shadow_type(frame unsafe.Pointer, type_ int) {
	c_frame := (*C.GtkFrame)(unsafe.Pointer(frame))

	c_type_ := (C.GtkShadowType)(type_)

	C.gtk_frame_set_shadow_type(c_frame, c_type_)
}

func Fn_gtk_gesture_get_last_event(gesture unsafe.Pointer, sequence unsafe.Pointer) unsafe.Pointer {
	c_gesture := (*C.GtkGesture)(unsafe.Pointer(gesture))

	c_sequence := (*C.GdkEventSequence)(unsafe.Pointer(sequence))

	ret := C.gtk_gesture_get_last_event(c_gesture, c_sequence)

	return unsafe.Pointer(ret)
}

func Fn_gtk_grid_new() unsafe.Pointer {
	ret := C.gtk_grid_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_grid_attach(grid unsafe.Pointer, child unsafe.Pointer, left int, top int, width int, height int) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_left := (C.gint)(left)

	c_top := (C.gint)(top)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_grid_attach(c_grid, c_child, c_left, c_top, c_width, c_height)
}

func Fn_gtk_grid_attach_next_to(grid unsafe.Pointer, child unsafe.Pointer, sibling unsafe.Pointer, side int, width int, height int) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_sibling := (*C.GtkWidget)(unsafe.Pointer(sibling))

	c_side := (C.GtkPositionType)(side)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_grid_attach_next_to(c_grid, c_child, c_sibling, c_side, c_width, c_height)
}

func Fn_gtk_grid_get_column_homogeneous(grid unsafe.Pointer) bool {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	ret := C.gtk_grid_get_column_homogeneous(c_grid)

	return toGoBool(ret)
}

func Fn_gtk_grid_get_column_spacing(grid unsafe.Pointer) uint {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	ret := C.gtk_grid_get_column_spacing(c_grid)

	return (uint)(ret)
}

func Fn_gtk_grid_get_row_homogeneous(grid unsafe.Pointer) bool {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	ret := C.gtk_grid_get_row_homogeneous(c_grid)

	return toGoBool(ret)
}

func Fn_gtk_grid_get_row_spacing(grid unsafe.Pointer) uint {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	ret := C.gtk_grid_get_row_spacing(c_grid)

	return (uint)(ret)
}

func Fn_gtk_grid_set_column_homogeneous(grid unsafe.Pointer, homogeneous bool) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_homogeneous := toCBool(homogeneous)

	C.gtk_grid_set_column_homogeneous(c_grid, c_homogeneous)
}

func Fn_gtk_grid_set_column_spacing(grid unsafe.Pointer, spacing uint) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_spacing := (C.guint)(spacing)

	C.gtk_grid_set_column_spacing(c_grid, c_spacing)
}

func Fn_gtk_grid_set_row_homogeneous(grid unsafe.Pointer, homogeneous bool) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_homogeneous := toCBool(homogeneous)

	C.gtk_grid_set_row_homogeneous(c_grid, c_homogeneous)
}

func Fn_gtk_grid_set_row_spacing(grid unsafe.Pointer, spacing uint) {
	c_grid := (*C.GtkGrid)(unsafe.Pointer(grid))

	c_spacing := (C.guint)(spacing)

	C.gtk_grid_set_row_spacing(c_grid, c_spacing)
}

func Fn_gtk_hbox_new(homogeneous bool, spacing int) unsafe.Pointer {
	c_homogeneous := toCBool(homogeneous)

	c_spacing := (C.gint)(spacing)

	ret := C.gtk_hbox_new(c_homogeneous, c_spacing)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hbutton_box_new() unsafe.Pointer {
	ret := C.gtk_hbutton_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_hpaned_new() unsafe.Pointer {
	ret := C.gtk_hpaned_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_hscale_new(adjustment unsafe.Pointer) unsafe.Pointer {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	ret := C.gtk_hscale_new(c_adjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hscale_new_with_range(min float64, max float64, step float64) unsafe.Pointer {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	ret := C.gtk_hscale_new_with_range(c_min, c_max, c_step)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hscrollbar_new(adjustment unsafe.Pointer) unsafe.Pointer {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	ret := C.gtk_hscrollbar_new(c_adjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_hseparator_new() unsafe.Pointer {
	ret := C.gtk_hseparator_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_handle_box_new() unsafe.Pointer {
	ret := C.gtk_handle_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_handle_box_get_handle_position(handleBox unsafe.Pointer) int {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	ret := C.gtk_handle_box_get_handle_position(c_handleBox)

	return (int)(ret)
}

func Fn_gtk_handle_box_get_shadow_type(handleBox unsafe.Pointer) int {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	ret := C.gtk_handle_box_get_shadow_type(c_handleBox)

	return (int)(ret)
}

func Fn_gtk_handle_box_get_snap_edge(handleBox unsafe.Pointer) int {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	ret := C.gtk_handle_box_get_snap_edge(c_handleBox)

	return (int)(ret)
}

func Fn_gtk_handle_box_set_handle_position(handleBox unsafe.Pointer, position int) {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	c_position := (C.GtkPositionType)(position)

	C.gtk_handle_box_set_handle_position(c_handleBox, c_position)
}

func Fn_gtk_handle_box_set_shadow_type(handleBox unsafe.Pointer, type_ int) {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	c_type_ := (C.GtkShadowType)(type_)

	C.gtk_handle_box_set_shadow_type(c_handleBox, c_type_)
}

func Fn_gtk_handle_box_set_snap_edge(handleBox unsafe.Pointer, edge int) {
	c_handleBox := (*C.GtkHandleBox)(unsafe.Pointer(handleBox))

	c_edge := (C.GtkPositionType)(edge)

	C.gtk_handle_box_set_snap_edge(c_handleBox, c_edge)
}

func Fn_gtk_im_context_delete_surrounding(context unsafe.Pointer, offset int, nChars int) bool {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_offset := (C.gint)(offset)

	c_nChars := (C.gint)(nChars)

	ret := C.gtk_im_context_delete_surrounding(c_context, c_offset, c_nChars)

	return toGoBool(ret)
}

func Fn_gtk_im_context_filter_keypress(context unsafe.Pointer, event unsafe.Pointer) bool {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_event := (*C.GdkEventKey)(unsafe.Pointer(event))

	ret := C.gtk_im_context_filter_keypress(c_context, c_event)

	return toGoBool(ret)
}

func Fn_gtk_im_context_focus_in(context unsafe.Pointer) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	C.gtk_im_context_focus_in(c_context)
}

func Fn_gtk_im_context_focus_out(context unsafe.Pointer) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	C.gtk_im_context_focus_out(c_context)
}

// UNSUPPORTED : gtk_im_context_get_preedit_string : parameter 'str' is non array with indirect count > 1

// UNSUPPORTED : gtk_im_context_get_surrounding : parameter 'text' is non array with indirect count > 1

func Fn_gtk_im_context_reset(context unsafe.Pointer) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	C.gtk_im_context_reset(c_context)
}

func Fn_gtk_im_context_set_client_window(context unsafe.Pointer, window unsafe.Pointer) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	C.gtk_im_context_set_client_window(c_context, c_window)
}

func Fn_gtk_im_context_set_cursor_location(context unsafe.Pointer, area unsafe.Pointer) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_area := (*C.GdkRectangle)(unsafe.Pointer(area))

	C.gtk_im_context_set_cursor_location(c_context, c_area)
}

func Fn_gtk_im_context_set_surrounding(context unsafe.Pointer, text string, len_ int, cursorIndex int) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	c_cursorIndex := (C.gint)(cursorIndex)

	C.gtk_im_context_set_surrounding(c_context, c_text, c_len_, c_cursorIndex)
}

func Fn_gtk_im_context_set_use_preedit(context unsafe.Pointer, usePreedit bool) {
	c_context := (*C.GtkIMContext)(unsafe.Pointer(context))

	c_usePreedit := toCBool(usePreedit)

	C.gtk_im_context_set_use_preedit(c_context, c_usePreedit)
}

func Fn_gtk_im_context_simple_new() unsafe.Pointer {
	ret := C.gtk_im_context_simple_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_im_context_simple_add_compose_file(contextSimple unsafe.Pointer, composeFile string) {
	c_contextSimple := (*C.GtkIMContextSimple)(unsafe.Pointer(contextSimple))

	c_composeFile := (*C.gchar)(C.CString(composeFile))
	defer C.free(unsafe.Pointer(c_composeFile))

	C.gtk_im_context_simple_add_compose_file(c_contextSimple, c_composeFile)
}

func Fn_gtk_im_context_simple_add_table(contextSimple unsafe.Pointer, data []uint16, maxSeqLen int, nSeqs int) {
	c_contextSimple := (*C.GtkIMContextSimple)(unsafe.Pointer(contextSimple))

	c_data := (*C.guint16)(unsafe.Pointer(&data[0]))

	c_maxSeqLen := (C.gint)(maxSeqLen)

	c_nSeqs := (C.gint)(nSeqs)

	C.gtk_im_context_simple_add_table(c_contextSimple, c_data, c_maxSeqLen, c_nSeqs)
}

func Fn_gtk_im_multicontext_new() unsafe.Pointer {
	ret := C.gtk_im_multicontext_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_im_multicontext_append_menuitems(context unsafe.Pointer, menushell unsafe.Pointer) {
	c_context := (*C.GtkIMMulticontext)(unsafe.Pointer(context))

	c_menushell := (*C.GtkMenuShell)(unsafe.Pointer(menushell))

	C.gtk_im_multicontext_append_menuitems(c_context, c_menushell)
}

func Fn_gtk_icon_factory_new() unsafe.Pointer {
	ret := C.gtk_icon_factory_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_factory_add(factory unsafe.Pointer, stockId string, iconSet unsafe.Pointer) {
	c_factory := (*C.GtkIconFactory)(unsafe.Pointer(factory))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	C.gtk_icon_factory_add(c_factory, c_stockId, c_iconSet)
}

func Fn_gtk_icon_factory_add_default(factory unsafe.Pointer) {
	c_factory := (*C.GtkIconFactory)(unsafe.Pointer(factory))

	C.gtk_icon_factory_add_default(c_factory)
}

func Fn_gtk_icon_factory_lookup(factory unsafe.Pointer, stockId string) unsafe.Pointer {
	c_factory := (*C.GtkIconFactory)(unsafe.Pointer(factory))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	ret := C.gtk_icon_factory_lookup(c_factory, c_stockId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_icon_factory_remove_default(factory unsafe.Pointer) {
	c_factory := (*C.GtkIconFactory)(unsafe.Pointer(factory))

	C.gtk_icon_factory_remove_default(c_factory)
}

func Fn_gtk_icon_factory_lookup_default(stockId string) unsafe.Pointer {
	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	ret := C.gtk_icon_factory_lookup_default(c_stockId)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_icon_info_load_icon_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_info_load_symbolic_for_context_async : parameter 'callback' is callback

// UNSUPPORTED : gtk_icon_theme_get_icon_sizes : no array length

// UNSUPPORTED : gtk_icon_view_get_cursor : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_get_dest_item_at_pos : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_get_drag_dest_item : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_get_item_at_pos : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_get_tooltip_context : parameter 'model' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_get_visible_range : parameter 'start_path' is non array with indirect count > 1

// UNSUPPORTED : gtk_icon_view_selected_foreach : parameter 'func' is callback

func Fn_gtk_image_new() unsafe.Pointer {
	ret := C.gtk_image_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_animation(animation unsafe.Pointer) unsafe.Pointer {
	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	ret := C.gtk_image_new_from_animation(c_animation)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_file(filename string) unsafe.Pointer {
	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	ret := C.gtk_image_new_from_file(c_filename)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_icon_set(iconSet unsafe.Pointer, size int) unsafe.Pointer {
	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	c_size := (C.GtkIconSize)(size)

	ret := C.gtk_image_new_from_icon_set(c_iconSet, c_size)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_pixbuf(pixbuf unsafe.Pointer) unsafe.Pointer {
	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	ret := C.gtk_image_new_from_pixbuf(c_pixbuf)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_new_from_stock(stockId string, size int) unsafe.Pointer {
	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_size := (C.GtkIconSize)(size)

	ret := C.gtk_image_new_from_stock(c_stockId, c_size)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_get_animation(image unsafe.Pointer) unsafe.Pointer {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	ret := C.gtk_image_get_animation(c_image)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_image_get_gicon : parameter 'gicon' is non array with indirect count > 1

// UNSUPPORTED : gtk_image_get_icon_name : parameter 'icon_name' is non array with indirect count > 1

// UNSUPPORTED : gtk_image_get_icon_set : parameter 'icon_set' is non array with indirect count > 1

func Fn_gtk_image_get_pixbuf(image unsafe.Pointer) unsafe.Pointer {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	ret := C.gtk_image_get_pixbuf(c_image)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_image_get_stock : parameter 'stock_id' is non array with indirect count > 1

func Fn_gtk_image_get_storage_type(image unsafe.Pointer) int {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	ret := C.gtk_image_get_storage_type(c_image)

	return (int)(ret)
}

func Fn_gtk_image_set_from_animation(image unsafe.Pointer, animation unsafe.Pointer) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	c_animation := (*C.GdkPixbufAnimation)(unsafe.Pointer(animation))

	C.gtk_image_set_from_animation(c_image, c_animation)
}

func Fn_gtk_image_set_from_file(image unsafe.Pointer, filename *string) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	var c_filenameValue (*C.gchar)
	if filename != nil {
		c_filenameValue = (*C.gchar)(C.CString(*filename))
		defer C.free(unsafe.Pointer(c_filenameValue))
	}
	c_filename := c_filenameValue

	C.gtk_image_set_from_file(c_image, c_filename)
}

func Fn_gtk_image_set_from_icon_set(image unsafe.Pointer, iconSet unsafe.Pointer, size int) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	c_iconSet := (*C.GtkIconSet)(unsafe.Pointer(iconSet))

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_icon_set(c_image, c_iconSet, c_size)
}

func Fn_gtk_image_set_from_pixbuf(image unsafe.Pointer, pixbuf unsafe.Pointer) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gtk_image_set_from_pixbuf(c_image, c_pixbuf)
}

func Fn_gtk_image_set_from_resource(image unsafe.Pointer, resourcePath *string) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	var c_resourcePathValue (*C.gchar)
	if resourcePath != nil {
		c_resourcePathValue = (*C.gchar)(C.CString(*resourcePath))
		defer C.free(unsafe.Pointer(c_resourcePathValue))
	}
	c_resourcePath := c_resourcePathValue

	C.gtk_image_set_from_resource(c_image, c_resourcePath)
}

func Fn_gtk_image_set_from_stock(image unsafe.Pointer, stockId string, size int) {
	c_image := (*C.GtkImage)(unsafe.Pointer(image))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_size := (C.GtkIconSize)(size)

	C.gtk_image_set_from_stock(c_image, c_stockId, c_size)
}

func Fn_gtk_image_menu_item_new() unsafe.Pointer {
	ret := C.gtk_image_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_from_stock(stockId string, accelGroup unsafe.Pointer) unsafe.Pointer {
	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	ret := C.gtk_image_menu_item_new_from_stock(c_stockId, c_accelGroup)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_image_menu_item_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_image_menu_item_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_get_image(imageMenuItem unsafe.Pointer) unsafe.Pointer {
	c_imageMenuItem := (*C.GtkImageMenuItem)(unsafe.Pointer(imageMenuItem))

	ret := C.gtk_image_menu_item_get_image(c_imageMenuItem)

	return unsafe.Pointer(ret)
}

func Fn_gtk_image_menu_item_set_image(imageMenuItem unsafe.Pointer, image unsafe.Pointer) {
	c_imageMenuItem := (*C.GtkImageMenuItem)(unsafe.Pointer(imageMenuItem))

	c_image := (*C.GtkWidget)(unsafe.Pointer(image))

	C.gtk_image_menu_item_set_image(c_imageMenuItem, c_image)
}

func Fn_gtk_info_bar_new_with_buttons(firstButtonText *string) unsafe.Pointer {
	var c_firstButtonTextValue (*C.gchar)
	if firstButtonText != nil {
		c_firstButtonTextValue = (*C.gchar)(C.CString(*firstButtonText))
		defer C.free(unsafe.Pointer(c_firstButtonTextValue))
	}
	c_firstButtonText := c_firstButtonTextValue

	ret := C.c_gtk_info_bar_new_with_buttons(c_firstButtonText)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_new() unsafe.Pointer {
	ret := C.gtk_invisible_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_new_for_screen(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gtk_invisible_new_for_screen(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_get_screen(invisible unsafe.Pointer) unsafe.Pointer {
	c_invisible := (*C.GtkInvisible)(unsafe.Pointer(invisible))

	ret := C.gtk_invisible_get_screen(c_invisible)

	return unsafe.Pointer(ret)
}

func Fn_gtk_invisible_set_screen(invisible unsafe.Pointer, screen unsafe.Pointer) {
	c_invisible := (*C.GtkInvisible)(unsafe.Pointer(invisible))

	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	C.gtk_invisible_set_screen(c_invisible, c_screen)
}

func Fn_gtk_label_new(str *string) unsafe.Pointer {
	var c_strValue (*C.gchar)
	if str != nil {
		c_strValue = (*C.gchar)(C.CString(*str))
		defer C.free(unsafe.Pointer(c_strValue))
	}
	c_str := c_strValue

	ret := C.gtk_label_new(c_str)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_new_with_mnemonic(str *string) unsafe.Pointer {
	var c_strValue (*C.gchar)
	if str != nil {
		c_strValue = (*C.gchar)(C.CString(*str))
		defer C.free(unsafe.Pointer(c_strValue))
	}
	c_str := c_strValue

	ret := C.gtk_label_new_with_mnemonic(c_str)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_attributes(label unsafe.Pointer) unsafe.Pointer {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_attributes(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_justify(label unsafe.Pointer) int {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_justify(c_label)

	return (int)(ret)
}

func Fn_gtk_label_get_label(label unsafe.Pointer) string {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_label(c_label)

	return C.GoString(ret)
}

func Fn_gtk_label_get_layout(label unsafe.Pointer) unsafe.Pointer {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_layout(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_layout_offsets(label unsafe.Pointer, x *int, y *int) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	C.gtk_label_get_layout_offsets(c_label, c_x, c_y)
}

func Fn_gtk_label_get_line_wrap(label unsafe.Pointer) bool {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_line_wrap(c_label)

	return toGoBool(ret)
}

func Fn_gtk_label_get_mnemonic_keyval(label unsafe.Pointer) uint {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_mnemonic_keyval(c_label)

	return (uint)(ret)
}

func Fn_gtk_label_get_mnemonic_widget(label unsafe.Pointer) unsafe.Pointer {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_mnemonic_widget(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_label_get_selectable(label unsafe.Pointer) bool {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_selectable(c_label)

	return toGoBool(ret)
}

func Fn_gtk_label_get_selection_bounds(label unsafe.Pointer, start *int, end *int) bool {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_start := (*C.gint)(unsafe.Pointer(start))

	c_end := (*C.gint)(unsafe.Pointer(end))

	ret := C.gtk_label_get_selection_bounds(c_label, c_start, c_end)

	return toGoBool(ret)
}

func Fn_gtk_label_get_text(label unsafe.Pointer) string {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_text(c_label)

	return C.GoString(ret)
}

func Fn_gtk_label_get_use_markup(label unsafe.Pointer) bool {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_use_markup(c_label)

	return toGoBool(ret)
}

func Fn_gtk_label_get_use_underline(label unsafe.Pointer) bool {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	ret := C.gtk_label_get_use_underline(c_label)

	return toGoBool(ret)
}

func Fn_gtk_label_select_region(label unsafe.Pointer, startOffset int, endOffset int) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_startOffset := (C.gint)(startOffset)

	c_endOffset := (C.gint)(endOffset)

	C.gtk_label_select_region(c_label, c_startOffset, c_endOffset)
}

func Fn_gtk_label_set_attributes(label unsafe.Pointer, attrs unsafe.Pointer) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_attrs := (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_label_set_attributes(c_label, c_attrs)
}

func Fn_gtk_label_set_justify(label unsafe.Pointer, jtype int) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_jtype := (C.GtkJustification)(jtype)

	C.gtk_label_set_justify(c_label, c_jtype)
}

func Fn_gtk_label_set_label(label unsafe.Pointer, str string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_label(c_label, c_str)
}

func Fn_gtk_label_set_line_wrap(label unsafe.Pointer, wrap bool) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_wrap := toCBool(wrap)

	C.gtk_label_set_line_wrap(c_label, c_wrap)
}

func Fn_gtk_label_set_markup(label unsafe.Pointer, str string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_markup(c_label, c_str)
}

func Fn_gtk_label_set_markup_with_mnemonic(label unsafe.Pointer, str string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_markup_with_mnemonic(c_label, c_str)
}

func Fn_gtk_label_set_mnemonic_widget(label unsafe.Pointer, widget unsafe.Pointer) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_label_set_mnemonic_widget(c_label, c_widget)
}

func Fn_gtk_label_set_pattern(label unsafe.Pointer, pattern string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_pattern := (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_label_set_pattern(c_label, c_pattern)
}

func Fn_gtk_label_set_selectable(label unsafe.Pointer, setting bool) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_setting := toCBool(setting)

	C.gtk_label_set_selectable(c_label, c_setting)
}

func Fn_gtk_label_set_text(label unsafe.Pointer, str string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_text(c_label, c_str)
}

func Fn_gtk_label_set_text_with_mnemonic(label unsafe.Pointer, str string) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_str := (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_label_set_text_with_mnemonic(c_label, c_str)
}

func Fn_gtk_label_set_use_markup(label unsafe.Pointer, setting bool) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_setting := toCBool(setting)

	C.gtk_label_set_use_markup(c_label, c_setting)
}

func Fn_gtk_label_set_use_underline(label unsafe.Pointer, setting bool) {
	c_label := (*C.GtkLabel)(unsafe.Pointer(label))

	c_setting := toCBool(setting)

	C.gtk_label_set_use_underline(c_label, c_setting)
}

func Fn_gtk_layout_new(hadjustment unsafe.Pointer, vadjustment unsafe.Pointer) unsafe.Pointer {
	c_hadjustment := (*C.GtkAdjustment)(unsafe.Pointer(hadjustment))

	c_vadjustment := (*C.GtkAdjustment)(unsafe.Pointer(vadjustment))

	ret := C.gtk_layout_new(c_hadjustment, c_vadjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_get_hadjustment(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	ret := C.gtk_layout_get_hadjustment(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_get_size(layout unsafe.Pointer, width *uint, height *uint) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_width := (*C.guint)(unsafe.Pointer(width))

	c_height := (*C.guint)(unsafe.Pointer(height))

	C.gtk_layout_get_size(c_layout, c_width, c_height)
}

func Fn_gtk_layout_get_vadjustment(layout unsafe.Pointer) unsafe.Pointer {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	ret := C.gtk_layout_get_vadjustment(c_layout)

	return unsafe.Pointer(ret)
}

func Fn_gtk_layout_move(layout unsafe.Pointer, childWidget unsafe.Pointer, x int, y int) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_childWidget := (*C.GtkWidget)(unsafe.Pointer(childWidget))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_layout_move(c_layout, c_childWidget, c_x, c_y)
}

func Fn_gtk_layout_put(layout unsafe.Pointer, childWidget unsafe.Pointer, x int, y int) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_childWidget := (*C.GtkWidget)(unsafe.Pointer(childWidget))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_layout_put(c_layout, c_childWidget, c_x, c_y)
}

func Fn_gtk_layout_set_hadjustment(layout unsafe.Pointer, adjustment unsafe.Pointer) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_layout_set_hadjustment(c_layout, c_adjustment)
}

func Fn_gtk_layout_set_size(layout unsafe.Pointer, width uint, height uint) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_width := (C.guint)(width)

	c_height := (C.guint)(height)

	C.gtk_layout_set_size(c_layout, c_width, c_height)
}

func Fn_gtk_layout_set_vadjustment(layout unsafe.Pointer, adjustment unsafe.Pointer) {
	c_layout := (*C.GtkLayout)(unsafe.Pointer(layout))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_layout_set_vadjustment(c_layout, c_adjustment)
}

// UNSUPPORTED : gtk_list_box_bind_model : parameter 'create_widget_func' is callback

// UNSUPPORTED : gtk_list_box_selected_foreach : parameter 'func' is callback

// UNSUPPORTED : gtk_list_box_set_filter_func : parameter 'filter_func' is callback

// UNSUPPORTED : gtk_list_box_set_header_func : parameter 'update_header' is callback

// UNSUPPORTED : gtk_list_box_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_list_store_newv(nColumns int, types []uint64) unsafe.Pointer {
	c_nColumns := (C.gint)(nColumns)

	c_types := (*C.GType)(unsafe.Pointer(&types[0]))

	ret := C.gtk_list_store_newv(c_nColumns, c_types)

	return unsafe.Pointer(ret)
}

func Fn_gtk_list_store_append(listStore unsafe.Pointer, iter unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_list_store_append(c_listStore, c_iter)
}

func Fn_gtk_list_store_clear(listStore unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	C.gtk_list_store_clear(c_listStore)
}

func Fn_gtk_list_store_insert(listStore unsafe.Pointer, iter unsafe.Pointer, position int) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_position := (C.gint)(position)

	C.gtk_list_store_insert(c_listStore, c_iter, c_position)
}

func Fn_gtk_list_store_insert_after(listStore unsafe.Pointer, iter unsafe.Pointer, sibling unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_sibling := (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_list_store_insert_after(c_listStore, c_iter, c_sibling)
}

func Fn_gtk_list_store_insert_before(listStore unsafe.Pointer, iter unsafe.Pointer, sibling unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_sibling := (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_list_store_insert_before(c_listStore, c_iter, c_sibling)
}

func Fn_gtk_list_store_iter_is_valid(listStore unsafe.Pointer, iter unsafe.Pointer) bool {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_list_store_iter_is_valid(c_listStore, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_list_store_move_after(store unsafe.Pointer, iter unsafe.Pointer, position unsafe.Pointer) {
	c_store := (*C.GtkListStore)(unsafe.Pointer(store))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_position := (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_list_store_move_after(c_store, c_iter, c_position)
}

func Fn_gtk_list_store_move_before(store unsafe.Pointer, iter unsafe.Pointer, position unsafe.Pointer) {
	c_store := (*C.GtkListStore)(unsafe.Pointer(store))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_position := (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_list_store_move_before(c_store, c_iter, c_position)
}

func Fn_gtk_list_store_prepend(listStore unsafe.Pointer, iter unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_list_store_prepend(c_listStore, c_iter)
}

func Fn_gtk_list_store_remove(listStore unsafe.Pointer, iter unsafe.Pointer) bool {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_list_store_remove(c_listStore, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_list_store_reorder(store unsafe.Pointer, newOrder []int) {
	c_store := (*C.GtkListStore)(unsafe.Pointer(store))

	c_newOrder := (*C.gint)(unsafe.Pointer(&newOrder[0]))

	C.gtk_list_store_reorder(c_store, c_newOrder)
}

func Fn_gtk_list_store_set_column_types(listStore unsafe.Pointer, nColumns int, types []uint64) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_nColumns := (C.gint)(nColumns)

	c_types := (*C.GType)(unsafe.Pointer(&types[0]))

	C.gtk_list_store_set_column_types(c_listStore, c_nColumns, c_types)
}

func Fn_gtk_list_store_set_valist(listStore unsafe.Pointer, iter unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.c_gtk_list_store_set_valist(c_listStore, c_iter)
}

func Fn_gtk_list_store_set_value(listStore unsafe.Pointer, iter unsafe.Pointer, column int, value unsafe.Pointer) {
	c_listStore := (*C.GtkListStore)(unsafe.Pointer(listStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_column := (C.gint)(column)

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_list_store_set_value(c_listStore, c_iter, c_column, c_value)
}

func Fn_gtk_list_store_swap(store unsafe.Pointer, a unsafe.Pointer, b unsafe.Pointer) {
	c_store := (*C.GtkListStore)(unsafe.Pointer(store))

	c_a := (*C.GtkTreeIter)(unsafe.Pointer(a))

	c_b := (*C.GtkTreeIter)(unsafe.Pointer(b))

	C.gtk_list_store_swap(c_store, c_a, c_b)
}

func Fn_gtk_menu_new() unsafe.Pointer {
	ret := C.gtk_menu_new()

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_menu_attach_to_widget : parameter 'detacher' is callback

func Fn_gtk_menu_detach(menu unsafe.Pointer) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	C.gtk_menu_detach(c_menu)
}

func Fn_gtk_menu_get_accel_group(menu unsafe.Pointer) unsafe.Pointer {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	ret := C.gtk_menu_get_accel_group(c_menu)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_active(menu unsafe.Pointer) unsafe.Pointer {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	ret := C.gtk_menu_get_active(c_menu)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_attach_widget(menu unsafe.Pointer) unsafe.Pointer {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	ret := C.gtk_menu_get_attach_widget(c_menu)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_get_tearoff_state(menu unsafe.Pointer) bool {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	ret := C.gtk_menu_get_tearoff_state(c_menu)

	return toGoBool(ret)
}

func Fn_gtk_menu_get_title(menu unsafe.Pointer) string {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	ret := C.gtk_menu_get_title(c_menu)

	return C.GoString(ret)
}

func Fn_gtk_menu_popdown(menu unsafe.Pointer) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	C.gtk_menu_popdown(c_menu)
}

// UNSUPPORTED : gtk_menu_popup : parameter 'func' is callback

// UNSUPPORTED : gtk_menu_popup_for_device : parameter 'func' is callback

func Fn_gtk_menu_reorder_child(menu unsafe.Pointer, child unsafe.Pointer, position int) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_position := (C.gint)(position)

	C.gtk_menu_reorder_child(c_menu, c_child, c_position)
}

func Fn_gtk_menu_reposition(menu unsafe.Pointer) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	C.gtk_menu_reposition(c_menu)
}

func Fn_gtk_menu_set_accel_group(menu unsafe.Pointer, accelGroup unsafe.Pointer) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_menu_set_accel_group(c_menu, c_accelGroup)
}

func Fn_gtk_menu_set_accel_path(menu unsafe.Pointer, accelPath *string) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	var c_accelPathValue (*C.gchar)
	if accelPath != nil {
		c_accelPathValue = (*C.gchar)(C.CString(*accelPath))
		defer C.free(unsafe.Pointer(c_accelPathValue))
	}
	c_accelPath := c_accelPathValue

	C.gtk_menu_set_accel_path(c_menu, c_accelPath)
}

func Fn_gtk_menu_set_active(menu unsafe.Pointer, index uint) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	c_index := (C.guint)(index)

	C.gtk_menu_set_active(c_menu, c_index)
}

func Fn_gtk_menu_set_screen(menu unsafe.Pointer, screen unsafe.Pointer) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	C.gtk_menu_set_screen(c_menu, c_screen)
}

func Fn_gtk_menu_set_tearoff_state(menu unsafe.Pointer, tornOff bool) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	c_tornOff := toCBool(tornOff)

	C.gtk_menu_set_tearoff_state(c_menu, c_tornOff)
}

func Fn_gtk_menu_set_title(menu unsafe.Pointer, title *string) {
	c_menu := (*C.GtkMenu)(unsafe.Pointer(menu))

	var c_titleValue (*C.gchar)
	if title != nil {
		c_titleValue = (*C.gchar)(C.CString(*title))
		defer C.free(unsafe.Pointer(c_titleValue))
	}
	c_title := c_titleValue

	C.gtk_menu_set_title(c_menu, c_title)
}

func Fn_gtk_menu_bar_new() unsafe.Pointer {
	ret := C.gtk_menu_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_new() unsafe.Pointer {
	ret := C.gtk_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_menu_item_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_menu_item_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_activate(menuItem unsafe.Pointer) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	C.gtk_menu_item_activate(c_menuItem)
}

func Fn_gtk_menu_item_deselect(menuItem unsafe.Pointer) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	C.gtk_menu_item_deselect(c_menuItem)
}

func Fn_gtk_menu_item_get_right_justified(menuItem unsafe.Pointer) bool {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	ret := C.gtk_menu_item_get_right_justified(c_menuItem)

	return toGoBool(ret)
}

func Fn_gtk_menu_item_get_submenu(menuItem unsafe.Pointer) unsafe.Pointer {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	ret := C.gtk_menu_item_get_submenu(c_menuItem)

	return unsafe.Pointer(ret)
}

func Fn_gtk_menu_item_select(menuItem unsafe.Pointer) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	C.gtk_menu_item_select(c_menuItem)
}

func Fn_gtk_menu_item_set_accel_path(menuItem unsafe.Pointer, accelPath *string) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	var c_accelPathValue (*C.gchar)
	if accelPath != nil {
		c_accelPathValue = (*C.gchar)(C.CString(*accelPath))
		defer C.free(unsafe.Pointer(c_accelPathValue))
	}
	c_accelPath := c_accelPathValue

	C.gtk_menu_item_set_accel_path(c_menuItem, c_accelPath)
}

func Fn_gtk_menu_item_set_right_justified(menuItem unsafe.Pointer, rightJustified bool) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	c_rightJustified := toCBool(rightJustified)

	C.gtk_menu_item_set_right_justified(c_menuItem, c_rightJustified)
}

func Fn_gtk_menu_item_set_submenu(menuItem unsafe.Pointer, submenu unsafe.Pointer) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	c_submenu := (*C.GtkWidget)(unsafe.Pointer(submenu))

	C.gtk_menu_item_set_submenu(c_menuItem, c_submenu)
}

func Fn_gtk_menu_item_toggle_size_allocate(menuItem unsafe.Pointer, allocation int) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	c_allocation := (C.gint)(allocation)

	C.gtk_menu_item_toggle_size_allocate(c_menuItem, c_allocation)
}

func Fn_gtk_menu_item_toggle_size_request(menuItem unsafe.Pointer, requisition *int) {
	c_menuItem := (*C.GtkMenuItem)(unsafe.Pointer(menuItem))

	c_requisition := (*C.gint)(unsafe.Pointer(requisition))

	C.gtk_menu_item_toggle_size_request(c_menuItem, c_requisition)
}

func Fn_gtk_menu_shell_activate_item(menuShell unsafe.Pointer, menuItem unsafe.Pointer, forceDeactivate bool) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_menuItem := (*C.GtkWidget)(unsafe.Pointer(menuItem))

	c_forceDeactivate := toCBool(forceDeactivate)

	C.gtk_menu_shell_activate_item(c_menuShell, c_menuItem, c_forceDeactivate)
}

func Fn_gtk_menu_shell_append(menuShell unsafe.Pointer, child unsafe.Pointer) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_menu_shell_append(c_menuShell, c_child)
}

func Fn_gtk_menu_shell_deactivate(menuShell unsafe.Pointer) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	C.gtk_menu_shell_deactivate(c_menuShell)
}

func Fn_gtk_menu_shell_deselect(menuShell unsafe.Pointer) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	C.gtk_menu_shell_deselect(c_menuShell)
}

func Fn_gtk_menu_shell_insert(menuShell unsafe.Pointer, child unsafe.Pointer, position int) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_position := (C.gint)(position)

	C.gtk_menu_shell_insert(c_menuShell, c_child, c_position)
}

func Fn_gtk_menu_shell_prepend(menuShell unsafe.Pointer, child unsafe.Pointer) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_menu_shell_prepend(c_menuShell, c_child)
}

func Fn_gtk_menu_shell_select_first(menuShell unsafe.Pointer, searchSensitive bool) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_searchSensitive := toCBool(searchSensitive)

	C.gtk_menu_shell_select_first(c_menuShell, c_searchSensitive)
}

func Fn_gtk_menu_shell_select_item(menuShell unsafe.Pointer, menuItem unsafe.Pointer) {
	c_menuShell := (*C.GtkMenuShell)(unsafe.Pointer(menuShell))

	c_menuItem := (*C.GtkWidget)(unsafe.Pointer(menuItem))

	C.gtk_menu_shell_select_item(c_menuShell, c_menuItem)
}

func Fn_gtk_message_dialog_new(parent unsafe.Pointer, flags int, type_ int, buttons int, messageFormat *string) unsafe.Pointer {
	c_parent := (*C.GtkWindow)(unsafe.Pointer(parent))

	c_flags := (C.GtkDialogFlags)(flags)

	c_type_ := (C.GtkMessageType)(type_)

	c_buttons := (C.GtkButtonsType)(buttons)

	var c_messageFormatValue (*C.gchar)
	if messageFormat != nil {
		c_messageFormatValue = (*C.gchar)(C.CString(*messageFormat))
		defer C.free(unsafe.Pointer(c_messageFormatValue))
	}
	c_messageFormat := c_messageFormatValue

	ret := C.c_gtk_message_dialog_new(c_parent, c_flags, c_type_, c_buttons, c_messageFormat)

	return unsafe.Pointer(ret)
}

func Fn_gtk_misc_get_alignment(misc unsafe.Pointer, xalign *float32, yalign *float32) {
	c_misc := (*C.GtkMisc)(unsafe.Pointer(misc))

	c_xalign := (*C.gfloat)(unsafe.Pointer(xalign))

	c_yalign := (*C.gfloat)(unsafe.Pointer(yalign))

	C.gtk_misc_get_alignment(c_misc, c_xalign, c_yalign)
}

func Fn_gtk_misc_get_padding(misc unsafe.Pointer, xpad *int, ypad *int) {
	c_misc := (*C.GtkMisc)(unsafe.Pointer(misc))

	c_xpad := (*C.gint)(unsafe.Pointer(xpad))

	c_ypad := (*C.gint)(unsafe.Pointer(ypad))

	C.gtk_misc_get_padding(c_misc, c_xpad, c_ypad)
}

func Fn_gtk_misc_set_alignment(misc unsafe.Pointer, xalign float32, yalign float32) {
	c_misc := (*C.GtkMisc)(unsafe.Pointer(misc))

	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_misc_set_alignment(c_misc, c_xalign, c_yalign)
}

func Fn_gtk_misc_set_padding(misc unsafe.Pointer, xpad int, ypad int) {
	c_misc := (*C.GtkMisc)(unsafe.Pointer(misc))

	c_xpad := (C.gint)(xpad)

	c_ypad := (C.gint)(ypad)

	C.gtk_misc_set_padding(c_misc, c_xpad, c_ypad)
}

func Fn_gtk_notebook_new() unsafe.Pointer {
	ret := C.gtk_notebook_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_append_page(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	ret := C.gtk_notebook_append_page(c_notebook, c_child, c_tabLabel)

	return (int)(ret)
}

func Fn_gtk_notebook_append_page_menu(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer, menuLabel unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	c_menuLabel := (*C.GtkWidget)(unsafe.Pointer(menuLabel))

	ret := C.gtk_notebook_append_page_menu(c_notebook, c_child, c_tabLabel, c_menuLabel)

	return (int)(ret)
}

func Fn_gtk_notebook_get_current_page(notebook unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_current_page(c_notebook)

	return (int)(ret)
}

func Fn_gtk_notebook_get_menu_label(notebook unsafe.Pointer, child unsafe.Pointer) unsafe.Pointer {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_get_menu_label(c_notebook, c_child)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_menu_label_text(notebook unsafe.Pointer, child unsafe.Pointer) string {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_get_menu_label_text(c_notebook, c_child)

	return C.GoString(ret)
}

func Fn_gtk_notebook_get_n_pages(notebook unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_n_pages(c_notebook)

	return (int)(ret)
}

func Fn_gtk_notebook_get_nth_page(notebook unsafe.Pointer, pageNum int) unsafe.Pointer {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_pageNum := (C.gint)(pageNum)

	ret := C.gtk_notebook_get_nth_page(c_notebook, c_pageNum)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_scrollable(notebook unsafe.Pointer) bool {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_scrollable(c_notebook)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_show_border(notebook unsafe.Pointer) bool {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_show_border(c_notebook)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_show_tabs(notebook unsafe.Pointer) bool {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_show_tabs(c_notebook)

	return toGoBool(ret)
}

func Fn_gtk_notebook_get_tab_label(notebook unsafe.Pointer, child unsafe.Pointer) unsafe.Pointer {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_get_tab_label(c_notebook, c_child)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_get_tab_label_text(notebook unsafe.Pointer, child unsafe.Pointer) string {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_get_tab_label_text(c_notebook, c_child)

	return C.GoString(ret)
}

func Fn_gtk_notebook_get_tab_pos(notebook unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	ret := C.gtk_notebook_get_tab_pos(c_notebook)

	return (int)(ret)
}

func Fn_gtk_notebook_insert_page(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer, position int) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	c_position := (C.gint)(position)

	ret := C.gtk_notebook_insert_page(c_notebook, c_child, c_tabLabel, c_position)

	return (int)(ret)
}

func Fn_gtk_notebook_insert_page_menu(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer, menuLabel unsafe.Pointer, position int) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	c_menuLabel := (*C.GtkWidget)(unsafe.Pointer(menuLabel))

	c_position := (C.gint)(position)

	ret := C.gtk_notebook_insert_page_menu(c_notebook, c_child, c_tabLabel, c_menuLabel, c_position)

	return (int)(ret)
}

func Fn_gtk_notebook_next_page(notebook unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	C.gtk_notebook_next_page(c_notebook)
}

func Fn_gtk_notebook_page_num(notebook unsafe.Pointer, child unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_page_num(c_notebook, c_child)

	return (int)(ret)
}

func Fn_gtk_notebook_popup_disable(notebook unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	C.gtk_notebook_popup_disable(c_notebook)
}

func Fn_gtk_notebook_popup_enable(notebook unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	C.gtk_notebook_popup_enable(c_notebook)
}

func Fn_gtk_notebook_prepend_page(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	ret := C.gtk_notebook_prepend_page(c_notebook, c_child, c_tabLabel)

	return (int)(ret)
}

func Fn_gtk_notebook_prepend_page_menu(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer, menuLabel unsafe.Pointer) int {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	c_menuLabel := (*C.GtkWidget)(unsafe.Pointer(menuLabel))

	ret := C.gtk_notebook_prepend_page_menu(c_notebook, c_child, c_tabLabel, c_menuLabel)

	return (int)(ret)
}

func Fn_gtk_notebook_prev_page(notebook unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	C.gtk_notebook_prev_page(c_notebook)
}

func Fn_gtk_notebook_remove_page(notebook unsafe.Pointer, pageNum int) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_pageNum := (C.gint)(pageNum)

	C.gtk_notebook_remove_page(c_notebook, c_pageNum)
}

func Fn_gtk_notebook_reorder_child(notebook unsafe.Pointer, child unsafe.Pointer, position int) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_position := (C.gint)(position)

	C.gtk_notebook_reorder_child(c_notebook, c_child, c_position)
}

func Fn_gtk_notebook_set_current_page(notebook unsafe.Pointer, pageNum int) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_pageNum := (C.gint)(pageNum)

	C.gtk_notebook_set_current_page(c_notebook, c_pageNum)
}

func Fn_gtk_notebook_set_menu_label(notebook unsafe.Pointer, child unsafe.Pointer, menuLabel unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_menuLabel := (*C.GtkWidget)(unsafe.Pointer(menuLabel))

	C.gtk_notebook_set_menu_label(c_notebook, c_child, c_menuLabel)
}

func Fn_gtk_notebook_set_menu_label_text(notebook unsafe.Pointer, child unsafe.Pointer, menuText string) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_menuText := (*C.gchar)(C.CString(menuText))
	defer C.free(unsafe.Pointer(c_menuText))

	C.gtk_notebook_set_menu_label_text(c_notebook, c_child, c_menuText)
}

func Fn_gtk_notebook_set_scrollable(notebook unsafe.Pointer, scrollable bool) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_scrollable := toCBool(scrollable)

	C.gtk_notebook_set_scrollable(c_notebook, c_scrollable)
}

func Fn_gtk_notebook_set_show_border(notebook unsafe.Pointer, showBorder bool) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_showBorder := toCBool(showBorder)

	C.gtk_notebook_set_show_border(c_notebook, c_showBorder)
}

func Fn_gtk_notebook_set_show_tabs(notebook unsafe.Pointer, showTabs bool) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_showTabs := toCBool(showTabs)

	C.gtk_notebook_set_show_tabs(c_notebook, c_showTabs)
}

func Fn_gtk_notebook_set_tab_label(notebook unsafe.Pointer, child unsafe.Pointer, tabLabel unsafe.Pointer) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabLabel := (*C.GtkWidget)(unsafe.Pointer(tabLabel))

	C.gtk_notebook_set_tab_label(c_notebook, c_child, c_tabLabel)
}

func Fn_gtk_notebook_set_tab_label_text(notebook unsafe.Pointer, child unsafe.Pointer, tabText string) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_tabText := (*C.gchar)(C.CString(tabText))
	defer C.free(unsafe.Pointer(c_tabText))

	C.gtk_notebook_set_tab_label_text(c_notebook, c_child, c_tabText)
}

func Fn_gtk_notebook_set_tab_pos(notebook unsafe.Pointer, pos int) {
	c_notebook := (*C.GtkNotebook)(unsafe.Pointer(notebook))

	c_pos := (C.GtkPositionType)(pos)

	C.gtk_notebook_set_tab_pos(c_notebook, c_pos)
}

func Fn_gtk_notebook_page_accessible_new(notebook unsafe.Pointer, child unsafe.Pointer) unsafe.Pointer {
	c_notebook := (*C.GtkNotebookAccessible)(unsafe.Pointer(notebook))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	ret := C.gtk_notebook_page_accessible_new(c_notebook, c_child)

	return unsafe.Pointer(ret)
}

func Fn_gtk_notebook_page_accessible_invalidate(page unsafe.Pointer) {
	c_page := (*C.GtkNotebookPageAccessible)(unsafe.Pointer(page))

	C.gtk_notebook_page_accessible_invalidate(c_page)
}

func Fn_gtk_paned_add1(paned unsafe.Pointer, child unsafe.Pointer) {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_paned_add1(c_paned, c_child)
}

func Fn_gtk_paned_add2(paned unsafe.Pointer, child unsafe.Pointer) {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_paned_add2(c_paned, c_child)
}

func Fn_gtk_paned_get_position(paned unsafe.Pointer) int {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	ret := C.gtk_paned_get_position(c_paned)

	return (int)(ret)
}

func Fn_gtk_paned_pack1(paned unsafe.Pointer, child unsafe.Pointer, resize bool, shrink bool) {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_resize := toCBool(resize)

	c_shrink := toCBool(shrink)

	C.gtk_paned_pack1(c_paned, c_child, c_resize, c_shrink)
}

func Fn_gtk_paned_pack2(paned unsafe.Pointer, child unsafe.Pointer, resize bool, shrink bool) {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_resize := toCBool(resize)

	c_shrink := toCBool(shrink)

	C.gtk_paned_pack2(c_paned, c_child, c_resize, c_shrink)
}

func Fn_gtk_paned_set_position(paned unsafe.Pointer, position int) {
	c_paned := (*C.GtkPaned)(unsafe.Pointer(paned))

	c_position := (C.gint)(position)

	C.gtk_paned_set_position(c_paned, c_position)
}

func Fn_gtk_places_sidebar_get_show_connect_to_server(sidebar unsafe.Pointer) bool {
	c_sidebar := (*C.GtkPlacesSidebar)(unsafe.Pointer(sidebar))

	ret := C.gtk_places_sidebar_get_show_connect_to_server(c_sidebar)

	return toGoBool(ret)
}

func Fn_gtk_plug_new(socketId uint64) unsafe.Pointer {
	c_socketId := (C.Window)(socketId)

	ret := C.gtk_plug_new(c_socketId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_plug_new_for_display(display unsafe.Pointer, socketId uint64) unsafe.Pointer {
	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_socketId := (C.Window)(socketId)

	ret := C.gtk_plug_new_for_display(c_display, c_socketId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_plug_construct(plug unsafe.Pointer, socketId uint64) {
	c_plug := (*C.GtkPlug)(unsafe.Pointer(plug))

	c_socketId := (C.Window)(socketId)

	C.gtk_plug_construct(c_plug, c_socketId)
}

func Fn_gtk_plug_construct_for_display(plug unsafe.Pointer, display unsafe.Pointer, socketId uint64) {
	c_plug := (*C.GtkPlug)(unsafe.Pointer(plug))

	c_display := (*C.GdkDisplay)(unsafe.Pointer(display))

	c_socketId := (C.Window)(socketId)

	C.gtk_plug_construct_for_display(c_plug, c_display, c_socketId)
}

func Fn_gtk_plug_get_id(plug unsafe.Pointer) uint64 {
	c_plug := (*C.GtkPlug)(unsafe.Pointer(plug))

	ret := C.gtk_plug_get_id(c_plug)

	return (uint64)(ret)
}

func Fn_gtk_popover_get_pointing_to(popover unsafe.Pointer, rect unsafe.Pointer) bool {
	c_popover := (*C.GtkPopover)(unsafe.Pointer(popover))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	ret := C.gtk_popover_get_pointing_to(c_popover, c_rect)

	return toGoBool(ret)
}

func Fn_gtk_popover_get_position(popover unsafe.Pointer) int {
	c_popover := (*C.GtkPopover)(unsafe.Pointer(popover))

	ret := C.gtk_popover_get_position(c_popover)

	return (int)(ret)
}

// UNSUPPORTED : gtk_print_settings_foreach : parameter 'func' is callback

func Fn_gtk_progress_bar_new() unsafe.Pointer {
	ret := C.gtk_progress_bar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_progress_bar_get_fraction(pbar unsafe.Pointer) float64 {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	ret := C.gtk_progress_bar_get_fraction(c_pbar)

	return (float64)(ret)
}

func Fn_gtk_progress_bar_get_inverted(pbar unsafe.Pointer) bool {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	ret := C.gtk_progress_bar_get_inverted(c_pbar)

	return toGoBool(ret)
}

func Fn_gtk_progress_bar_get_pulse_step(pbar unsafe.Pointer) float64 {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	ret := C.gtk_progress_bar_get_pulse_step(c_pbar)

	return (float64)(ret)
}

func Fn_gtk_progress_bar_get_text(pbar unsafe.Pointer) string {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	ret := C.gtk_progress_bar_get_text(c_pbar)

	return C.GoString(ret)
}

func Fn_gtk_progress_bar_pulse(pbar unsafe.Pointer) {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	C.gtk_progress_bar_pulse(c_pbar)
}

func Fn_gtk_progress_bar_set_fraction(pbar unsafe.Pointer, fraction float64) {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	c_fraction := (C.gdouble)(fraction)

	C.gtk_progress_bar_set_fraction(c_pbar, c_fraction)
}

func Fn_gtk_progress_bar_set_inverted(pbar unsafe.Pointer, inverted bool) {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	c_inverted := toCBool(inverted)

	C.gtk_progress_bar_set_inverted(c_pbar, c_inverted)
}

func Fn_gtk_progress_bar_set_pulse_step(pbar unsafe.Pointer, fraction float64) {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	c_fraction := (C.gdouble)(fraction)

	C.gtk_progress_bar_set_pulse_step(c_pbar, c_fraction)
}

func Fn_gtk_progress_bar_set_text(pbar unsafe.Pointer, text *string) {
	c_pbar := (*C.GtkProgressBar)(unsafe.Pointer(pbar))

	var c_textValue (*C.gchar)
	if text != nil {
		c_textValue = (*C.gchar)(C.CString(*text))
		defer C.free(unsafe.Pointer(c_textValue))
	}
	c_text := c_textValue

	C.gtk_progress_bar_set_text(c_pbar, c_text)
}

func Fn_gtk_radio_button_new(group unsafe.Pointer) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	ret := C.gtk_radio_button_new(c_group)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_from_widget(radioGroupMember unsafe.Pointer) unsafe.Pointer {
	c_radioGroupMember := (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember))

	ret := C.gtk_radio_button_new_from_widget(c_radioGroupMember)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_label(group unsafe.Pointer, label string) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_button_new_with_label(c_group, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_label_from_widget(radioGroupMember unsafe.Pointer, label string) unsafe.Pointer {
	c_radioGroupMember := (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_button_new_with_label_from_widget(c_radioGroupMember, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_mnemonic(group unsafe.Pointer, label string) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_button_new_with_mnemonic(c_group, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(radioGroupMember unsafe.Pointer, label string) unsafe.Pointer {
	c_radioGroupMember := (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_button_new_with_mnemonic_from_widget(c_radioGroupMember, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_get_group(radioButton unsafe.Pointer) unsafe.Pointer {
	c_radioButton := (*C.GtkRadioButton)(unsafe.Pointer(radioButton))

	ret := C.gtk_radio_button_get_group(c_radioButton)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_button_set_group(radioButton unsafe.Pointer, group unsafe.Pointer) {
	c_radioButton := (*C.GtkRadioButton)(unsafe.Pointer(radioButton))

	c_group := (*C.GSList)(unsafe.Pointer(group))

	C.gtk_radio_button_set_group(c_radioButton, c_group)
}

func Fn_gtk_radio_menu_item_new(group unsafe.Pointer) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	ret := C.gtk_radio_menu_item_new(c_group)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_label(group unsafe.Pointer, label string) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_menu_item_new_with_label(c_group, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_new_with_mnemonic(group unsafe.Pointer, label string) unsafe.Pointer {
	c_group := (*C.GSList)(unsafe.Pointer(group))

	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_radio_menu_item_new_with_mnemonic(c_group, c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_get_group(radioMenuItem unsafe.Pointer) unsafe.Pointer {
	c_radioMenuItem := (*C.GtkRadioMenuItem)(unsafe.Pointer(radioMenuItem))

	ret := C.gtk_radio_menu_item_get_group(c_radioMenuItem)

	return unsafe.Pointer(ret)
}

func Fn_gtk_radio_menu_item_set_group(radioMenuItem unsafe.Pointer, group unsafe.Pointer) {
	c_radioMenuItem := (*C.GtkRadioMenuItem)(unsafe.Pointer(radioMenuItem))

	c_group := (*C.GSList)(unsafe.Pointer(group))

	C.gtk_radio_menu_item_set_group(c_radioMenuItem, c_group)
}

func Fn_gtk_range_get_adjustment(range_ unsafe.Pointer) unsafe.Pointer {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	ret := C.gtk_range_get_adjustment(c_range_)

	return unsafe.Pointer(ret)
}

func Fn_gtk_range_get_inverted(range_ unsafe.Pointer) bool {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	ret := C.gtk_range_get_inverted(c_range_)

	return toGoBool(ret)
}

func Fn_gtk_range_get_value(range_ unsafe.Pointer) float64 {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	ret := C.gtk_range_get_value(c_range_)

	return (float64)(ret)
}

func Fn_gtk_range_set_adjustment(range_ unsafe.Pointer, adjustment unsafe.Pointer) {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_range_set_adjustment(c_range_, c_adjustment)
}

func Fn_gtk_range_set_increments(range_ unsafe.Pointer, step float64, page float64) {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	c_step := (C.gdouble)(step)

	c_page := (C.gdouble)(page)

	C.gtk_range_set_increments(c_range_, c_step, c_page)
}

func Fn_gtk_range_set_inverted(range_ unsafe.Pointer, setting bool) {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	c_setting := toCBool(setting)

	C.gtk_range_set_inverted(c_range_, c_setting)
}

func Fn_gtk_range_set_range(range_ unsafe.Pointer, min float64, max float64) {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	C.gtk_range_set_range(c_range_, c_min, c_max)
}

func Fn_gtk_range_set_value(range_ unsafe.Pointer, value float64) {
	c_range_ := (*C.GtkRange)(unsafe.Pointer(range_))

	c_value := (C.gdouble)(value)

	C.gtk_range_set_value(c_range_, c_value)
}

func Fn_gtk_rc_style_new() unsafe.Pointer {
	ret := C.gtk_rc_style_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_rc_style_copy(orig unsafe.Pointer) unsafe.Pointer {
	c_orig := (*C.GtkRcStyle)(unsafe.Pointer(orig))

	ret := C.gtk_rc_style_copy(c_orig)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_recent_filter_add_custom : parameter 'func' is callback

func Fn_gtk_renderer_cell_accessible_new(renderer unsafe.Pointer) unsafe.Pointer {
	c_renderer := (*C.GtkCellRenderer)(unsafe.Pointer(renderer))

	ret := C.gtk_renderer_cell_accessible_new(c_renderer)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scale_get_digits(scale unsafe.Pointer) int {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	ret := C.gtk_scale_get_digits(c_scale)

	return (int)(ret)
}

func Fn_gtk_scale_get_draw_value(scale unsafe.Pointer) bool {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	ret := C.gtk_scale_get_draw_value(c_scale)

	return toGoBool(ret)
}

func Fn_gtk_scale_get_value_pos(scale unsafe.Pointer) int {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	ret := C.gtk_scale_get_value_pos(c_scale)

	return (int)(ret)
}

func Fn_gtk_scale_set_digits(scale unsafe.Pointer, digits int) {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	c_digits := (C.gint)(digits)

	C.gtk_scale_set_digits(c_scale, c_digits)
}

func Fn_gtk_scale_set_draw_value(scale unsafe.Pointer, drawValue bool) {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	c_drawValue := toCBool(drawValue)

	C.gtk_scale_set_draw_value(c_scale, c_drawValue)
}

func Fn_gtk_scale_set_value_pos(scale unsafe.Pointer, pos int) {
	c_scale := (*C.GtkScale)(unsafe.Pointer(scale))

	c_pos := (C.GtkPositionType)(pos)

	C.gtk_scale_set_value_pos(c_scale, c_pos)
}

func Fn_gtk_scrolled_window_new(hadjustment unsafe.Pointer, vadjustment unsafe.Pointer) unsafe.Pointer {
	c_hadjustment := (*C.GtkAdjustment)(unsafe.Pointer(hadjustment))

	c_vadjustment := (*C.GtkAdjustment)(unsafe.Pointer(vadjustment))

	ret := C.gtk_scrolled_window_new(c_hadjustment, c_vadjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_add_with_viewport(scrolledWindow unsafe.Pointer, child unsafe.Pointer) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	C.gtk_scrolled_window_add_with_viewport(c_scrolledWindow, c_child)
}

func Fn_gtk_scrolled_window_get_hadjustment(scrolledWindow unsafe.Pointer) unsafe.Pointer {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	ret := C.gtk_scrolled_window_get_hadjustment(c_scrolledWindow)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_get_placement(scrolledWindow unsafe.Pointer) int {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	ret := C.gtk_scrolled_window_get_placement(c_scrolledWindow)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_policy(scrolledWindow unsafe.Pointer, hscrollbarPolicy *int, vscrollbarPolicy *int) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_hscrollbarPolicy := (*C.GtkPolicyType)(unsafe.Pointer(hscrollbarPolicy))

	c_vscrollbarPolicy := (*C.GtkPolicyType)(unsafe.Pointer(vscrollbarPolicy))

	C.gtk_scrolled_window_get_policy(c_scrolledWindow, c_hscrollbarPolicy, c_vscrollbarPolicy)
}

func Fn_gtk_scrolled_window_get_shadow_type(scrolledWindow unsafe.Pointer) int {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	ret := C.gtk_scrolled_window_get_shadow_type(c_scrolledWindow)

	return (int)(ret)
}

func Fn_gtk_scrolled_window_get_vadjustment(scrolledWindow unsafe.Pointer) unsafe.Pointer {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	ret := C.gtk_scrolled_window_get_vadjustment(c_scrolledWindow)

	return unsafe.Pointer(ret)
}

func Fn_gtk_scrolled_window_set_hadjustment(scrolledWindow unsafe.Pointer, hadjustment unsafe.Pointer) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_hadjustment := (*C.GtkAdjustment)(unsafe.Pointer(hadjustment))

	C.gtk_scrolled_window_set_hadjustment(c_scrolledWindow, c_hadjustment)
}

func Fn_gtk_scrolled_window_set_placement(scrolledWindow unsafe.Pointer, windowPlacement int) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_windowPlacement := (C.GtkCornerType)(windowPlacement)

	C.gtk_scrolled_window_set_placement(c_scrolledWindow, c_windowPlacement)
}

func Fn_gtk_scrolled_window_set_policy(scrolledWindow unsafe.Pointer, hscrollbarPolicy int, vscrollbarPolicy int) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_hscrollbarPolicy := (C.GtkPolicyType)(hscrollbarPolicy)

	c_vscrollbarPolicy := (C.GtkPolicyType)(vscrollbarPolicy)

	C.gtk_scrolled_window_set_policy(c_scrolledWindow, c_hscrollbarPolicy, c_vscrollbarPolicy)
}

func Fn_gtk_scrolled_window_set_shadow_type(scrolledWindow unsafe.Pointer, type_ int) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_type_ := (C.GtkShadowType)(type_)

	C.gtk_scrolled_window_set_shadow_type(c_scrolledWindow, c_type_)
}

func Fn_gtk_scrolled_window_set_vadjustment(scrolledWindow unsafe.Pointer, vadjustment unsafe.Pointer) {
	c_scrolledWindow := (*C.GtkScrolledWindow)(unsafe.Pointer(scrolledWindow))

	c_vadjustment := (*C.GtkAdjustment)(unsafe.Pointer(vadjustment))

	C.gtk_scrolled_window_set_vadjustment(c_scrolledWindow, c_vadjustment)
}

func Fn_gtk_separator_menu_item_new() unsafe.Pointer {
	ret := C.gtk_separator_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_settings_set_double_property(settings unsafe.Pointer, name string, vDouble float64, origin string) {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_vDouble := (C.gdouble)(vDouble)

	c_origin := (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_double_property(c_settings, c_name, c_vDouble, c_origin)
}

func Fn_gtk_settings_set_long_property(settings unsafe.Pointer, name string, vLong int64, origin string) {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_vLong := (C.glong)(vLong)

	c_origin := (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_long_property(c_settings, c_name, c_vLong, c_origin)
}

func Fn_gtk_settings_set_property_value(settings unsafe.Pointer, name string, svalue unsafe.Pointer) {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_svalue := (*C.GtkSettingsValue)(unsafe.Pointer(svalue))

	C.gtk_settings_set_property_value(c_settings, c_name, c_svalue)
}

func Fn_gtk_settings_set_string_property(settings unsafe.Pointer, name string, vString string, origin string) {
	c_settings := (*C.GtkSettings)(unsafe.Pointer(settings))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_vString := (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(c_vString))

	c_origin := (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(c_origin))

	C.gtk_settings_set_string_property(c_settings, c_name, c_vString, c_origin)
}

func Fn_gtk_settings_get_default() unsafe.Pointer {
	ret := C.gtk_settings_get_default()

	return unsafe.Pointer(ret)
}

func Fn_gtk_settings_get_for_screen(screen unsafe.Pointer) unsafe.Pointer {
	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	ret := C.gtk_settings_get_for_screen(c_screen)

	return unsafe.Pointer(ret)
}

func Fn_gtk_settings_install_property(pspec unsafe.Pointer) {
	c_pspec := (*C.GParamSpec)(unsafe.Pointer(pspec))

	C.gtk_settings_install_property(c_pspec)
}

// UNSUPPORTED : gtk_settings_install_property_parser : parameter 'parser' is callback

func Fn_gtk_size_group_new(mode int) unsafe.Pointer {
	c_mode := (C.GtkSizeGroupMode)(mode)

	ret := C.gtk_size_group_new(c_mode)

	return unsafe.Pointer(ret)
}

func Fn_gtk_size_group_add_widget(sizeGroup unsafe.Pointer, widget unsafe.Pointer) {
	c_sizeGroup := (*C.GtkSizeGroup)(unsafe.Pointer(sizeGroup))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_size_group_add_widget(c_sizeGroup, c_widget)
}

func Fn_gtk_size_group_get_mode(sizeGroup unsafe.Pointer) int {
	c_sizeGroup := (*C.GtkSizeGroup)(unsafe.Pointer(sizeGroup))

	ret := C.gtk_size_group_get_mode(c_sizeGroup)

	return (int)(ret)
}

func Fn_gtk_size_group_remove_widget(sizeGroup unsafe.Pointer, widget unsafe.Pointer) {
	c_sizeGroup := (*C.GtkSizeGroup)(unsafe.Pointer(sizeGroup))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_size_group_remove_widget(c_sizeGroup, c_widget)
}

func Fn_gtk_size_group_set_mode(sizeGroup unsafe.Pointer, mode int) {
	c_sizeGroup := (*C.GtkSizeGroup)(unsafe.Pointer(sizeGroup))

	c_mode := (C.GtkSizeGroupMode)(mode)

	C.gtk_size_group_set_mode(c_sizeGroup, c_mode)
}

func Fn_gtk_socket_new() unsafe.Pointer {
	ret := C.gtk_socket_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_socket_add_id(socket unsafe.Pointer, window uint64) {
	c_socket := (*C.GtkSocket)(unsafe.Pointer(socket))

	c_window := (C.Window)(window)

	C.gtk_socket_add_id(c_socket, c_window)
}

func Fn_gtk_socket_get_id(socket unsafe.Pointer) uint64 {
	c_socket := (*C.GtkSocket)(unsafe.Pointer(socket))

	ret := C.gtk_socket_get_id(c_socket)

	return (uint64)(ret)
}

func Fn_gtk_spin_button_new(adjustment unsafe.Pointer, climbRate float64, digits uint) unsafe.Pointer {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	c_climbRate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	ret := C.gtk_spin_button_new(c_adjustment, c_climbRate, c_digits)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_new_with_range(min float64, max float64, step float64) unsafe.Pointer {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	ret := C.gtk_spin_button_new_with_range(c_min, c_max, c_step)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_configure(spinButton unsafe.Pointer, adjustment unsafe.Pointer, climbRate float64, digits uint) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	c_climbRate := (C.gdouble)(climbRate)

	c_digits := (C.guint)(digits)

	C.gtk_spin_button_configure(c_spinButton, c_adjustment, c_climbRate, c_digits)
}

func Fn_gtk_spin_button_get_adjustment(spinButton unsafe.Pointer) unsafe.Pointer {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_adjustment(c_spinButton)

	return unsafe.Pointer(ret)
}

func Fn_gtk_spin_button_get_digits(spinButton unsafe.Pointer) uint {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_digits(c_spinButton)

	return (uint)(ret)
}

func Fn_gtk_spin_button_get_increments(spinButton unsafe.Pointer, step *float64, page *float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_step := (*C.gdouble)(unsafe.Pointer(step))

	c_page := (*C.gdouble)(unsafe.Pointer(page))

	C.gtk_spin_button_get_increments(c_spinButton, c_step, c_page)
}

func Fn_gtk_spin_button_get_numeric(spinButton unsafe.Pointer) bool {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_numeric(c_spinButton)

	return toGoBool(ret)
}

func Fn_gtk_spin_button_get_range(spinButton unsafe.Pointer, min *float64, max *float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_min := (*C.gdouble)(unsafe.Pointer(min))

	c_max := (*C.gdouble)(unsafe.Pointer(max))

	C.gtk_spin_button_get_range(c_spinButton, c_min, c_max)
}

func Fn_gtk_spin_button_get_snap_to_ticks(spinButton unsafe.Pointer) bool {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_snap_to_ticks(c_spinButton)

	return toGoBool(ret)
}

func Fn_gtk_spin_button_get_update_policy(spinButton unsafe.Pointer) int {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_update_policy(c_spinButton)

	return (int)(ret)
}

func Fn_gtk_spin_button_get_value(spinButton unsafe.Pointer) float64 {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_value(c_spinButton)

	return (float64)(ret)
}

func Fn_gtk_spin_button_get_value_as_int(spinButton unsafe.Pointer) int {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_value_as_int(c_spinButton)

	return (int)(ret)
}

func Fn_gtk_spin_button_get_wrap(spinButton unsafe.Pointer) bool {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	ret := C.gtk_spin_button_get_wrap(c_spinButton)

	return toGoBool(ret)
}

func Fn_gtk_spin_button_set_adjustment(spinButton unsafe.Pointer, adjustment unsafe.Pointer) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_spin_button_set_adjustment(c_spinButton, c_adjustment)
}

func Fn_gtk_spin_button_set_digits(spinButton unsafe.Pointer, digits uint) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_digits := (C.guint)(digits)

	C.gtk_spin_button_set_digits(c_spinButton, c_digits)
}

func Fn_gtk_spin_button_set_increments(spinButton unsafe.Pointer, step float64, page float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_step := (C.gdouble)(step)

	c_page := (C.gdouble)(page)

	C.gtk_spin_button_set_increments(c_spinButton, c_step, c_page)
}

func Fn_gtk_spin_button_set_numeric(spinButton unsafe.Pointer, numeric bool) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_numeric := toCBool(numeric)

	C.gtk_spin_button_set_numeric(c_spinButton, c_numeric)
}

func Fn_gtk_spin_button_set_range(spinButton unsafe.Pointer, min float64, max float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	C.gtk_spin_button_set_range(c_spinButton, c_min, c_max)
}

func Fn_gtk_spin_button_set_snap_to_ticks(spinButton unsafe.Pointer, snapToTicks bool) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_snapToTicks := toCBool(snapToTicks)

	C.gtk_spin_button_set_snap_to_ticks(c_spinButton, c_snapToTicks)
}

func Fn_gtk_spin_button_set_update_policy(spinButton unsafe.Pointer, policy int) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_policy := (C.GtkSpinButtonUpdatePolicy)(policy)

	C.gtk_spin_button_set_update_policy(c_spinButton, c_policy)
}

func Fn_gtk_spin_button_set_value(spinButton unsafe.Pointer, value float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_value := (C.gdouble)(value)

	C.gtk_spin_button_set_value(c_spinButton, c_value)
}

func Fn_gtk_spin_button_set_wrap(spinButton unsafe.Pointer, wrap bool) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_wrap := toCBool(wrap)

	C.gtk_spin_button_set_wrap(c_spinButton, c_wrap)
}

func Fn_gtk_spin_button_spin(spinButton unsafe.Pointer, direction int, increment float64) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	c_direction := (C.GtkSpinType)(direction)

	c_increment := (C.gdouble)(increment)

	C.gtk_spin_button_spin(c_spinButton, c_direction, c_increment)
}

func Fn_gtk_spin_button_update(spinButton unsafe.Pointer) {
	c_spinButton := (*C.GtkSpinButton)(unsafe.Pointer(spinButton))

	C.gtk_spin_button_update(c_spinButton)
}

// UNSUPPORTED : gtk_status_icon_get_geometry : parameter 'screen' is non array with indirect count > 1

func Fn_gtk_statusbar_new() unsafe.Pointer {
	ret := C.gtk_statusbar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_statusbar_get_context_id(statusbar unsafe.Pointer, contextDescription string) uint {
	c_statusbar := (*C.GtkStatusbar)(unsafe.Pointer(statusbar))

	c_contextDescription := (*C.gchar)(C.CString(contextDescription))
	defer C.free(unsafe.Pointer(c_contextDescription))

	ret := C.gtk_statusbar_get_context_id(c_statusbar, c_contextDescription)

	return (uint)(ret)
}

func Fn_gtk_statusbar_pop(statusbar unsafe.Pointer, contextId uint) {
	c_statusbar := (*C.GtkStatusbar)(unsafe.Pointer(statusbar))

	c_contextId := (C.guint)(contextId)

	C.gtk_statusbar_pop(c_statusbar, c_contextId)
}

func Fn_gtk_statusbar_push(statusbar unsafe.Pointer, contextId uint, text string) uint {
	c_statusbar := (*C.GtkStatusbar)(unsafe.Pointer(statusbar))

	c_contextId := (C.guint)(contextId)

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	ret := C.gtk_statusbar_push(c_statusbar, c_contextId, c_text)

	return (uint)(ret)
}

func Fn_gtk_statusbar_remove(statusbar unsafe.Pointer, contextId uint, messageId uint) {
	c_statusbar := (*C.GtkStatusbar)(unsafe.Pointer(statusbar))

	c_contextId := (C.guint)(contextId)

	c_messageId := (C.guint)(messageId)

	C.gtk_statusbar_remove(c_statusbar, c_contextId, c_messageId)
}

func Fn_gtk_style_new() unsafe.Pointer {
	ret := C.gtk_style_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_apply_default_background(style unsafe.Pointer, cr unsafe.Pointer, window unsafe.Pointer, stateType int, x int, y int, width int, height int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_cr := (*C.cairo_t)(unsafe.Pointer(cr))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_stateType := (C.GtkStateType)(stateType)

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_style_apply_default_background(c_style, c_cr, c_window, c_stateType, c_x, c_y, c_width, c_height)
}

func Fn_gtk_style_attach(style unsafe.Pointer, window unsafe.Pointer) unsafe.Pointer {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gtk_style_attach(c_style, c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_copy(style unsafe.Pointer) unsafe.Pointer {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	ret := C.gtk_style_copy(c_style)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_detach(style unsafe.Pointer) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	C.gtk_style_detach(c_style)
}

func Fn_gtk_style_lookup_icon_set(style unsafe.Pointer, stockId string) unsafe.Pointer {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	ret := C.gtk_style_lookup_icon_set(c_style, c_stockId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_render_icon(style unsafe.Pointer, source unsafe.Pointer, direction int, state int, size int, widget unsafe.Pointer, detail *string) unsafe.Pointer {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_source := (*C.GtkIconSource)(unsafe.Pointer(source))

	c_direction := (C.GtkTextDirection)(direction)

	c_state := (C.GtkStateType)(state)

	c_size := (C.GtkIconSize)(size)

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	ret := C.gtk_style_render_icon(c_style, c_source, c_direction, c_state, c_size, c_widget, c_detail)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_set_background(style unsafe.Pointer, window unsafe.Pointer, stateType int) {
	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	c_stateType := (C.GtkStateType)(stateType)

	C.gtk_style_set_background(c_style, c_window, c_stateType)
}

func Fn_gtk_style_context_new() unsafe.Pointer {
	ret := C.gtk_style_context_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_screen(context unsafe.Pointer) unsafe.Pointer {
	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	ret := C.gtk_style_context_get_screen(c_context)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_section(context unsafe.Pointer, property string) unsafe.Pointer {
	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	c_property := (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(c_property))

	ret := C.gtk_style_context_get_section(c_context, c_property)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_context_get_style_property(context unsafe.Pointer, propertyName string, value unsafe.Pointer) {
	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_style_context_get_style_property(c_context, c_propertyName, c_value)
}

func Fn_gtk_style_context_lookup_color(context unsafe.Pointer, colorName string, color unsafe.Pointer) bool {
	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	c_colorName := (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(c_colorName))

	c_color := (*C.GdkRGBA)(unsafe.Pointer(color))

	ret := C.gtk_style_context_lookup_color(c_context, c_colorName, c_color)

	return toGoBool(ret)
}

func Fn_gtk_style_context_lookup_icon_set(context unsafe.Pointer, stockId string) unsafe.Pointer {
	c_context := (*C.GtkStyleContext)(unsafe.Pointer(context))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	ret := C.gtk_style_context_lookup_icon_set(c_context, c_stockId)

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_properties_new() unsafe.Pointer {
	ret := C.gtk_style_properties_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_style_properties_clear(props unsafe.Pointer) {
	c_props := (*C.GtkStyleProperties)(unsafe.Pointer(props))

	C.gtk_style_properties_clear(c_props)
}

// UNSUPPORTED : gtk_style_properties_lookup_property : parameter 'parse_func' is callback

// UNSUPPORTED : gtk_style_properties_register_property : parameter 'parse_func' is callback

func Fn_gtk_table_new(rows uint, columns uint, homogeneous bool) unsafe.Pointer {
	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	c_homogeneous := toCBool(homogeneous)

	ret := C.gtk_table_new(c_rows, c_columns, c_homogeneous)

	return unsafe.Pointer(ret)
}

func Fn_gtk_table_attach(table unsafe.Pointer, child unsafe.Pointer, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint, xoptions int, yoptions int, xpadding uint, ypadding uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_leftAttach := (C.guint)(leftAttach)

	c_rightAttach := (C.guint)(rightAttach)

	c_topAttach := (C.guint)(topAttach)

	c_bottomAttach := (C.guint)(bottomAttach)

	c_xoptions := (C.GtkAttachOptions)(xoptions)

	c_yoptions := (C.GtkAttachOptions)(yoptions)

	c_xpadding := (C.guint)(xpadding)

	c_ypadding := (C.guint)(ypadding)

	C.gtk_table_attach(c_table, c_child, c_leftAttach, c_rightAttach, c_topAttach, c_bottomAttach, c_xoptions, c_yoptions, c_xpadding, c_ypadding)
}

func Fn_gtk_table_attach_defaults(table unsafe.Pointer, widget unsafe.Pointer, leftAttach uint, rightAttach uint, topAttach uint, bottomAttach uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_leftAttach := (C.guint)(leftAttach)

	c_rightAttach := (C.guint)(rightAttach)

	c_topAttach := (C.guint)(topAttach)

	c_bottomAttach := (C.guint)(bottomAttach)

	C.gtk_table_attach_defaults(c_table, c_widget, c_leftAttach, c_rightAttach, c_topAttach, c_bottomAttach)
}

func Fn_gtk_table_get_col_spacing(table unsafe.Pointer, column uint) uint {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_column := (C.guint)(column)

	ret := C.gtk_table_get_col_spacing(c_table, c_column)

	return (uint)(ret)
}

func Fn_gtk_table_get_default_col_spacing(table unsafe.Pointer) uint {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	ret := C.gtk_table_get_default_col_spacing(c_table)

	return (uint)(ret)
}

func Fn_gtk_table_get_default_row_spacing(table unsafe.Pointer) uint {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	ret := C.gtk_table_get_default_row_spacing(c_table)

	return (uint)(ret)
}

func Fn_gtk_table_get_homogeneous(table unsafe.Pointer) bool {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	ret := C.gtk_table_get_homogeneous(c_table)

	return toGoBool(ret)
}

func Fn_gtk_table_get_row_spacing(table unsafe.Pointer, row uint) uint {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_row := (C.guint)(row)

	ret := C.gtk_table_get_row_spacing(c_table, c_row)

	return (uint)(ret)
}

func Fn_gtk_table_resize(table unsafe.Pointer, rows uint, columns uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_rows := (C.guint)(rows)

	c_columns := (C.guint)(columns)

	C.gtk_table_resize(c_table, c_rows, c_columns)
}

func Fn_gtk_table_set_col_spacing(table unsafe.Pointer, column uint, spacing uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_column := (C.guint)(column)

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_col_spacing(c_table, c_column, c_spacing)
}

func Fn_gtk_table_set_col_spacings(table unsafe.Pointer, spacing uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_col_spacings(c_table, c_spacing)
}

func Fn_gtk_table_set_homogeneous(table unsafe.Pointer, homogeneous bool) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_homogeneous := toCBool(homogeneous)

	C.gtk_table_set_homogeneous(c_table, c_homogeneous)
}

func Fn_gtk_table_set_row_spacing(table unsafe.Pointer, row uint, spacing uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_row := (C.guint)(row)

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_row_spacing(c_table, c_row, c_spacing)
}

func Fn_gtk_table_set_row_spacings(table unsafe.Pointer, spacing uint) {
	c_table := (*C.GtkTable)(unsafe.Pointer(table))

	c_spacing := (C.guint)(spacing)

	C.gtk_table_set_row_spacings(c_table, c_spacing)
}

func Fn_gtk_tearoff_menu_item_new() unsafe.Pointer {
	ret := C.gtk_tearoff_menu_item_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_new(table unsafe.Pointer) unsafe.Pointer {
	c_table := (*C.GtkTextTagTable)(unsafe.Pointer(table))

	ret := C.gtk_text_buffer_new(c_table)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_add_selection_clipboard(buffer unsafe.Pointer, clipboard unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	C.gtk_text_buffer_add_selection_clipboard(c_buffer, c_clipboard)
}

func Fn_gtk_text_buffer_apply_tag(buffer unsafe.Pointer, tag unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_apply_tag(c_buffer, c_tag, c_start, c_end)
}

func Fn_gtk_text_buffer_apply_tag_by_name(buffer unsafe.Pointer, name string, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_apply_tag_by_name(c_buffer, c_name, c_start, c_end)
}

func Fn_gtk_text_buffer_begin_user_action(buffer unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	C.gtk_text_buffer_begin_user_action(c_buffer)
}

func Fn_gtk_text_buffer_copy_clipboard(buffer unsafe.Pointer, clipboard unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	C.gtk_text_buffer_copy_clipboard(c_buffer, c_clipboard)
}

func Fn_gtk_text_buffer_create_child_anchor(buffer unsafe.Pointer, iter unsafe.Pointer) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_buffer_create_child_anchor(c_buffer, c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_create_mark(buffer unsafe.Pointer, markName *string, where unsafe.Pointer, leftGravity bool) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	var c_markNameValue (*C.gchar)
	if markName != nil {
		c_markNameValue = (*C.gchar)(C.CString(*markName))
		defer C.free(unsafe.Pointer(c_markNameValue))
	}
	c_markName := c_markNameValue

	c_where := (*C.GtkTextIter)(unsafe.Pointer(where))

	c_leftGravity := toCBool(leftGravity)

	ret := C.gtk_text_buffer_create_mark(c_buffer, c_markName, c_where, c_leftGravity)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_create_tag(buffer unsafe.Pointer, tagName *string, firstPropertyName *string) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	var c_tagNameValue (*C.gchar)
	if tagName != nil {
		c_tagNameValue = (*C.gchar)(C.CString(*tagName))
		defer C.free(unsafe.Pointer(c_tagNameValue))
	}
	c_tagName := c_tagNameValue

	var c_firstPropertyNameValue (*C.gchar)
	if firstPropertyName != nil {
		c_firstPropertyNameValue = (*C.gchar)(C.CString(*firstPropertyName))
		defer C.free(unsafe.Pointer(c_firstPropertyNameValue))
	}
	c_firstPropertyName := c_firstPropertyNameValue

	ret := C.c_gtk_text_buffer_create_tag(c_buffer, c_tagName, c_firstPropertyName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_cut_clipboard(buffer unsafe.Pointer, clipboard unsafe.Pointer, defaultEditable bool) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	c_defaultEditable := toCBool(defaultEditable)

	C.gtk_text_buffer_cut_clipboard(c_buffer, c_clipboard, c_defaultEditable)
}

func Fn_gtk_text_buffer_delete(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_delete(c_buffer, c_start, c_end)
}

func Fn_gtk_text_buffer_delete_interactive(buffer unsafe.Pointer, startIter unsafe.Pointer, endIter unsafe.Pointer, defaultEditable bool) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_startIter := (*C.GtkTextIter)(unsafe.Pointer(startIter))

	c_endIter := (*C.GtkTextIter)(unsafe.Pointer(endIter))

	c_defaultEditable := toCBool(defaultEditable)

	ret := C.gtk_text_buffer_delete_interactive(c_buffer, c_startIter, c_endIter, c_defaultEditable)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_delete_mark(buffer unsafe.Pointer, mark unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	C.gtk_text_buffer_delete_mark(c_buffer, c_mark)
}

func Fn_gtk_text_buffer_delete_mark_by_name(buffer unsafe.Pointer, name string) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_text_buffer_delete_mark_by_name(c_buffer, c_name)
}

func Fn_gtk_text_buffer_delete_selection(buffer unsafe.Pointer, interactive bool, defaultEditable bool) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_interactive := toCBool(interactive)

	c_defaultEditable := toCBool(defaultEditable)

	ret := C.gtk_text_buffer_delete_selection(c_buffer, c_interactive, c_defaultEditable)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_end_user_action(buffer unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	C.gtk_text_buffer_end_user_action(c_buffer)
}

func Fn_gtk_text_buffer_get_bounds(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_get_bounds(c_buffer, c_start, c_end)
}

func Fn_gtk_text_buffer_get_char_count(buffer unsafe.Pointer) int {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_char_count(c_buffer)

	return (int)(ret)
}

func Fn_gtk_text_buffer_get_end_iter(buffer unsafe.Pointer, iter unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	C.gtk_text_buffer_get_end_iter(c_buffer, c_iter)
}

func Fn_gtk_text_buffer_get_insert(buffer unsafe.Pointer) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_insert(c_buffer)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_iter_at_child_anchor(buffer unsafe.Pointer, iter unsafe.Pointer, anchor unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_anchor := (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor))

	C.gtk_text_buffer_get_iter_at_child_anchor(c_buffer, c_iter, c_anchor)
}

func Fn_gtk_text_buffer_get_iter_at_line(buffer unsafe.Pointer, iter unsafe.Pointer, lineNumber int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_lineNumber := (C.gint)(lineNumber)

	C.gtk_text_buffer_get_iter_at_line(c_buffer, c_iter, c_lineNumber)
}

func Fn_gtk_text_buffer_get_iter_at_line_index(buffer unsafe.Pointer, iter unsafe.Pointer, lineNumber int, byteIndex int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_lineNumber := (C.gint)(lineNumber)

	c_byteIndex := (C.gint)(byteIndex)

	C.gtk_text_buffer_get_iter_at_line_index(c_buffer, c_iter, c_lineNumber, c_byteIndex)
}

func Fn_gtk_text_buffer_get_iter_at_line_offset(buffer unsafe.Pointer, iter unsafe.Pointer, lineNumber int, charOffset int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_lineNumber := (C.gint)(lineNumber)

	c_charOffset := (C.gint)(charOffset)

	C.gtk_text_buffer_get_iter_at_line_offset(c_buffer, c_iter, c_lineNumber, c_charOffset)
}

func Fn_gtk_text_buffer_get_iter_at_mark(buffer unsafe.Pointer, iter unsafe.Pointer, mark unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	C.gtk_text_buffer_get_iter_at_mark(c_buffer, c_iter, c_mark)
}

func Fn_gtk_text_buffer_get_iter_at_offset(buffer unsafe.Pointer, iter unsafe.Pointer, charOffset int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_charOffset := (C.gint)(charOffset)

	C.gtk_text_buffer_get_iter_at_offset(c_buffer, c_iter, c_charOffset)
}

func Fn_gtk_text_buffer_get_line_count(buffer unsafe.Pointer) int {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_line_count(c_buffer)

	return (int)(ret)
}

func Fn_gtk_text_buffer_get_mark(buffer unsafe.Pointer, name string) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.gtk_text_buffer_get_mark(c_buffer, c_name)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_modified(buffer unsafe.Pointer) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_modified(c_buffer)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_get_selection_bound(buffer unsafe.Pointer) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_selection_bound(c_buffer)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_selection_bounds(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	ret := C.gtk_text_buffer_get_selection_bounds(c_buffer, c_start, c_end)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_get_slice(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer, includeHiddenChars bool) string {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	c_includeHiddenChars := toCBool(includeHiddenChars)

	ret := C.gtk_text_buffer_get_slice(c_buffer, c_start, c_end, c_includeHiddenChars)

	return C.GoString(ret)
}

func Fn_gtk_text_buffer_get_start_iter(buffer unsafe.Pointer, iter unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	C.gtk_text_buffer_get_start_iter(c_buffer, c_iter)
}

func Fn_gtk_text_buffer_get_tag_table(buffer unsafe.Pointer) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_buffer_get_tag_table(c_buffer)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_buffer_get_text(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer, includeHiddenChars bool) string {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	c_includeHiddenChars := toCBool(includeHiddenChars)

	ret := C.gtk_text_buffer_get_text(c_buffer, c_start, c_end, c_includeHiddenChars)

	return C.GoString(ret)
}

func Fn_gtk_text_buffer_insert(buffer unsafe.Pointer, iter unsafe.Pointer, text string, len_ int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	C.gtk_text_buffer_insert(c_buffer, c_iter, c_text, c_len_)
}

func Fn_gtk_text_buffer_insert_at_cursor(buffer unsafe.Pointer, text string, len_ int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	C.gtk_text_buffer_insert_at_cursor(c_buffer, c_text, c_len_)
}

func Fn_gtk_text_buffer_insert_child_anchor(buffer unsafe.Pointer, iter unsafe.Pointer, anchor unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_anchor := (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor))

	C.gtk_text_buffer_insert_child_anchor(c_buffer, c_iter, c_anchor)
}

func Fn_gtk_text_buffer_insert_interactive(buffer unsafe.Pointer, iter unsafe.Pointer, text string, len_ int, defaultEditable bool) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	c_defaultEditable := toCBool(defaultEditable)

	ret := C.gtk_text_buffer_insert_interactive(c_buffer, c_iter, c_text, c_len_, c_defaultEditable)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_insert_interactive_at_cursor(buffer unsafe.Pointer, text string, len_ int, defaultEditable bool) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	c_defaultEditable := toCBool(defaultEditable)

	ret := C.gtk_text_buffer_insert_interactive_at_cursor(c_buffer, c_text, c_len_, c_defaultEditable)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_insert_pixbuf(buffer unsafe.Pointer, iter unsafe.Pointer, pixbuf unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gtk_text_buffer_insert_pixbuf(c_buffer, c_iter, c_pixbuf)
}

func Fn_gtk_text_buffer_insert_range(buffer unsafe.Pointer, iter unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_insert_range(c_buffer, c_iter, c_start, c_end)
}

func Fn_gtk_text_buffer_insert_range_interactive(buffer unsafe.Pointer, iter unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer, defaultEditable bool) bool {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	c_defaultEditable := toCBool(defaultEditable)

	ret := C.gtk_text_buffer_insert_range_interactive(c_buffer, c_iter, c_start, c_end, c_defaultEditable)

	return toGoBool(ret)
}

func Fn_gtk_text_buffer_insert_with_tags(buffer unsafe.Pointer, iter unsafe.Pointer, text string, len_ int, firstTag unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	c_firstTag := (*C.GtkTextTag)(unsafe.Pointer(firstTag))

	C.c_gtk_text_buffer_insert_with_tags(c_buffer, c_iter, c_text, c_len_, c_firstTag)
}

func Fn_gtk_text_buffer_insert_with_tags_by_name(buffer unsafe.Pointer, iter unsafe.Pointer, text string, len_ int, firstTagName string) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	c_firstTagName := (*C.gchar)(C.CString(firstTagName))
	defer C.free(unsafe.Pointer(c_firstTagName))

	C.c_gtk_text_buffer_insert_with_tags_by_name(c_buffer, c_iter, c_text, c_len_, c_firstTagName)
}

func Fn_gtk_text_buffer_move_mark(buffer unsafe.Pointer, mark unsafe.Pointer, where unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	c_where := (*C.GtkTextIter)(unsafe.Pointer(where))

	C.gtk_text_buffer_move_mark(c_buffer, c_mark, c_where)
}

func Fn_gtk_text_buffer_move_mark_by_name(buffer unsafe.Pointer, name string, where unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_where := (*C.GtkTextIter)(unsafe.Pointer(where))

	C.gtk_text_buffer_move_mark_by_name(c_buffer, c_name, c_where)
}

func Fn_gtk_text_buffer_paste_clipboard(buffer unsafe.Pointer, clipboard unsafe.Pointer, overrideLocation unsafe.Pointer, defaultEditable bool) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	c_overrideLocation := (*C.GtkTextIter)(unsafe.Pointer(overrideLocation))

	c_defaultEditable := toCBool(defaultEditable)

	C.gtk_text_buffer_paste_clipboard(c_buffer, c_clipboard, c_overrideLocation, c_defaultEditable)
}

func Fn_gtk_text_buffer_place_cursor(buffer unsafe.Pointer, where unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_where := (*C.GtkTextIter)(unsafe.Pointer(where))

	C.gtk_text_buffer_place_cursor(c_buffer, c_where)
}

// UNSUPPORTED : gtk_text_buffer_register_deserialize_format : parameter 'function' is callback

// UNSUPPORTED : gtk_text_buffer_register_serialize_format : parameter 'function' is callback

func Fn_gtk_text_buffer_remove_all_tags(buffer unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_remove_all_tags(c_buffer, c_start, c_end)
}

func Fn_gtk_text_buffer_remove_selection_clipboard(buffer unsafe.Pointer, clipboard unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_clipboard := (*C.GtkClipboard)(unsafe.Pointer(clipboard))

	C.gtk_text_buffer_remove_selection_clipboard(c_buffer, c_clipboard)
}

func Fn_gtk_text_buffer_remove_tag(buffer unsafe.Pointer, tag unsafe.Pointer, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_remove_tag(c_buffer, c_tag, c_start, c_end)
}

func Fn_gtk_text_buffer_remove_tag_by_name(buffer unsafe.Pointer, name string, start unsafe.Pointer, end unsafe.Pointer) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	c_start := (*C.GtkTextIter)(unsafe.Pointer(start))

	c_end := (*C.GtkTextIter)(unsafe.Pointer(end))

	C.gtk_text_buffer_remove_tag_by_name(c_buffer, c_name, c_start, c_end)
}

func Fn_gtk_text_buffer_set_modified(buffer unsafe.Pointer, setting bool) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_setting := toCBool(setting)

	C.gtk_text_buffer_set_modified(c_buffer, c_setting)
}

func Fn_gtk_text_buffer_set_text(buffer unsafe.Pointer, text string, len_ int) {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	c_text := (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(c_text))

	c_len_ := (C.gint)(len_)

	C.gtk_text_buffer_set_text(c_buffer, c_text, c_len_)
}

func Fn_gtk_text_child_anchor_new() unsafe.Pointer {
	ret := C.gtk_text_child_anchor_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_child_anchor_get_deleted(anchor unsafe.Pointer) bool {
	c_anchor := (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor))

	ret := C.gtk_text_child_anchor_get_deleted(c_anchor)

	return toGoBool(ret)
}

func Fn_gtk_text_child_anchor_get_widgets(anchor unsafe.Pointer) unsafe.Pointer {
	c_anchor := (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor))

	ret := C.gtk_text_child_anchor_get_widgets(c_anchor)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_mark_get_buffer(mark unsafe.Pointer) unsafe.Pointer {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_mark_get_buffer(c_mark)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_mark_get_deleted(mark unsafe.Pointer) bool {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_mark_get_deleted(c_mark)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_get_left_gravity(mark unsafe.Pointer) bool {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_mark_get_left_gravity(c_mark)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_get_name(mark unsafe.Pointer) string {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_mark_get_name(c_mark)

	return C.GoString(ret)
}

func Fn_gtk_text_mark_get_visible(mark unsafe.Pointer) bool {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_mark_get_visible(c_mark)

	return toGoBool(ret)
}

func Fn_gtk_text_mark_set_visible(mark unsafe.Pointer, setting bool) {
	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	c_setting := toCBool(setting)

	C.gtk_text_mark_set_visible(c_mark, c_setting)
}

func Fn_gtk_text_tag_new(name *string) unsafe.Pointer {
	var c_nameValue (*C.gchar)
	if name != nil {
		c_nameValue = (*C.gchar)(C.CString(*name))
		defer C.free(unsafe.Pointer(c_nameValue))
	}
	c_name := c_nameValue

	ret := C.gtk_text_tag_new(c_name)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_event(tag unsafe.Pointer, eventObject unsafe.Pointer, event unsafe.Pointer, iter unsafe.Pointer) bool {
	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	c_eventObject := (*C.GObject)(unsafe.Pointer(eventObject))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_tag_event(c_tag, c_eventObject, c_event, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_tag_get_priority(tag unsafe.Pointer) int {
	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_tag_get_priority(c_tag)

	return (int)(ret)
}

func Fn_gtk_text_tag_set_priority(tag unsafe.Pointer, priority int) {
	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	c_priority := (C.gint)(priority)

	C.gtk_text_tag_set_priority(c_tag, c_priority)
}

func Fn_gtk_text_tag_table_new() unsafe.Pointer {
	ret := C.gtk_text_tag_table_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_table_add(table unsafe.Pointer, tag unsafe.Pointer) bool {
	c_table := (*C.GtkTextTagTable)(unsafe.Pointer(table))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	ret := C.gtk_text_tag_table_add(c_table, c_tag)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_text_tag_table_foreach : parameter 'func' is callback

func Fn_gtk_text_tag_table_get_size(table unsafe.Pointer) int {
	c_table := (*C.GtkTextTagTable)(unsafe.Pointer(table))

	ret := C.gtk_text_tag_table_get_size(c_table)

	return (int)(ret)
}

func Fn_gtk_text_tag_table_lookup(table unsafe.Pointer, name string) unsafe.Pointer {
	c_table := (*C.GtkTextTagTable)(unsafe.Pointer(table))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.gtk_text_tag_table_lookup(c_table, c_name)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_tag_table_remove(table unsafe.Pointer, tag unsafe.Pointer) {
	c_table := (*C.GtkTextTagTable)(unsafe.Pointer(table))

	c_tag := (*C.GtkTextTag)(unsafe.Pointer(tag))

	C.gtk_text_tag_table_remove(c_table, c_tag)
}

func Fn_gtk_text_view_new() unsafe.Pointer {
	ret := C.gtk_text_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_new_with_buffer(buffer unsafe.Pointer) unsafe.Pointer {
	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	ret := C.gtk_text_view_new_with_buffer(c_buffer)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_add_child_at_anchor(textView unsafe.Pointer, child unsafe.Pointer, anchor unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_anchor := (*C.GtkTextChildAnchor)(unsafe.Pointer(anchor))

	C.gtk_text_view_add_child_at_anchor(c_textView, c_child, c_anchor)
}

func Fn_gtk_text_view_add_child_in_window(textView unsafe.Pointer, child unsafe.Pointer, whichWindow int, xpos int, ypos int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_whichWindow := (C.GtkTextWindowType)(whichWindow)

	c_xpos := (C.gint)(xpos)

	c_ypos := (C.gint)(ypos)

	C.gtk_text_view_add_child_in_window(c_textView, c_child, c_whichWindow, c_xpos, c_ypos)
}

func Fn_gtk_text_view_backward_display_line(textView unsafe.Pointer, iter unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_view_backward_display_line(c_textView, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_view_backward_display_line_start(textView unsafe.Pointer, iter unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_view_backward_display_line_start(c_textView, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_view_buffer_to_window_coords(textView unsafe.Pointer, win int, bufferX int, bufferY int, windowX *int, windowY *int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_win := (C.GtkTextWindowType)(win)

	c_bufferX := (C.gint)(bufferX)

	c_bufferY := (C.gint)(bufferY)

	c_windowX := (*C.gint)(unsafe.Pointer(windowX))

	c_windowY := (*C.gint)(unsafe.Pointer(windowY))

	C.gtk_text_view_buffer_to_window_coords(c_textView, c_win, c_bufferX, c_bufferY, c_windowX, c_windowY)
}

func Fn_gtk_text_view_forward_display_line(textView unsafe.Pointer, iter unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_view_forward_display_line(c_textView, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_view_forward_display_line_end(textView unsafe.Pointer, iter unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_view_forward_display_line_end(c_textView, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_border_window_size(textView unsafe.Pointer, type_ int) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_type_ := (C.GtkTextWindowType)(type_)

	ret := C.gtk_text_view_get_border_window_size(c_textView, c_type_)

	return (int)(ret)
}

func Fn_gtk_text_view_get_buffer(textView unsafe.Pointer) unsafe.Pointer {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_buffer(c_textView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_cursor_visible(textView unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_cursor_visible(c_textView)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_default_attributes(textView unsafe.Pointer) unsafe.Pointer {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_default_attributes(c_textView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_editable(textView unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_editable(c_textView)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_indent(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_indent(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_iter_at_location(textView unsafe.Pointer, iter unsafe.Pointer, x int, y int) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	ret := C.gtk_text_view_get_iter_at_location(c_textView, c_iter, c_x, c_y)

	return toGoBool(ret)
}

func Fn_gtk_text_view_get_iter_location(textView unsafe.Pointer, iter unsafe.Pointer, location unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_location := (*C.GdkRectangle)(unsafe.Pointer(location))

	C.gtk_text_view_get_iter_location(c_textView, c_iter, c_location)
}

func Fn_gtk_text_view_get_justification(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_justification(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_left_margin(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_left_margin(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_line_at_y(textView unsafe.Pointer, targetIter unsafe.Pointer, y int, lineTop *int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_targetIter := (*C.GtkTextIter)(unsafe.Pointer(targetIter))

	c_y := (C.gint)(y)

	c_lineTop := (*C.gint)(unsafe.Pointer(lineTop))

	C.gtk_text_view_get_line_at_y(c_textView, c_targetIter, c_y, c_lineTop)
}

func Fn_gtk_text_view_get_line_yrange(textView unsafe.Pointer, iter unsafe.Pointer, y *int, height *int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_text_view_get_line_yrange(c_textView, c_iter, c_y, c_height)
}

func Fn_gtk_text_view_get_pixels_above_lines(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_pixels_above_lines(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_pixels_below_lines(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_pixels_below_lines(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_pixels_inside_wrap(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_pixels_inside_wrap(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_right_margin(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_right_margin(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_get_tabs(textView unsafe.Pointer) unsafe.Pointer {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_tabs(c_textView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_visible_rect(textView unsafe.Pointer, visibleRect unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_visibleRect := (*C.GdkRectangle)(unsafe.Pointer(visibleRect))

	C.gtk_text_view_get_visible_rect(c_textView, c_visibleRect)
}

func Fn_gtk_text_view_get_window(textView unsafe.Pointer, win int) unsafe.Pointer {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_win := (C.GtkTextWindowType)(win)

	ret := C.gtk_text_view_get_window(c_textView, c_win)

	return unsafe.Pointer(ret)
}

func Fn_gtk_text_view_get_window_type(textView unsafe.Pointer, window unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_window := (*C.GdkWindow)(unsafe.Pointer(window))

	ret := C.gtk_text_view_get_window_type(c_textView, c_window)

	return (int)(ret)
}

func Fn_gtk_text_view_get_wrap_mode(textView unsafe.Pointer) int {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_get_wrap_mode(c_textView)

	return (int)(ret)
}

func Fn_gtk_text_view_move_child(textView unsafe.Pointer, child unsafe.Pointer, xpos int, ypos int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_child := (*C.GtkWidget)(unsafe.Pointer(child))

	c_xpos := (C.gint)(xpos)

	c_ypos := (C.gint)(ypos)

	C.gtk_text_view_move_child(c_textView, c_child, c_xpos, c_ypos)
}

func Fn_gtk_text_view_move_mark_onscreen(textView unsafe.Pointer, mark unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	ret := C.gtk_text_view_move_mark_onscreen(c_textView, c_mark)

	return toGoBool(ret)
}

func Fn_gtk_text_view_move_visually(textView unsafe.Pointer, iter unsafe.Pointer, count int) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_count := (C.gint)(count)

	ret := C.gtk_text_view_move_visually(c_textView, c_iter, c_count)

	return toGoBool(ret)
}

func Fn_gtk_text_view_place_cursor_onscreen(textView unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	ret := C.gtk_text_view_place_cursor_onscreen(c_textView)

	return toGoBool(ret)
}

func Fn_gtk_text_view_scroll_mark_onscreen(textView unsafe.Pointer, mark unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	C.gtk_text_view_scroll_mark_onscreen(c_textView, c_mark)
}

func Fn_gtk_text_view_scroll_to_iter(textView unsafe.Pointer, iter unsafe.Pointer, withinMargin float64, useAlign bool, xalign float64, yalign float64) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	c_withinMargin := (C.gdouble)(withinMargin)

	c_useAlign := toCBool(useAlign)

	c_xalign := (C.gdouble)(xalign)

	c_yalign := (C.gdouble)(yalign)

	ret := C.gtk_text_view_scroll_to_iter(c_textView, c_iter, c_withinMargin, c_useAlign, c_xalign, c_yalign)

	return toGoBool(ret)
}

func Fn_gtk_text_view_scroll_to_mark(textView unsafe.Pointer, mark unsafe.Pointer, withinMargin float64, useAlign bool, xalign float64, yalign float64) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_mark := (*C.GtkTextMark)(unsafe.Pointer(mark))

	c_withinMargin := (C.gdouble)(withinMargin)

	c_useAlign := toCBool(useAlign)

	c_xalign := (C.gdouble)(xalign)

	c_yalign := (C.gdouble)(yalign)

	C.gtk_text_view_scroll_to_mark(c_textView, c_mark, c_withinMargin, c_useAlign, c_xalign, c_yalign)
}

func Fn_gtk_text_view_set_border_window_size(textView unsafe.Pointer, type_ int, size int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_type_ := (C.GtkTextWindowType)(type_)

	c_size := (C.gint)(size)

	C.gtk_text_view_set_border_window_size(c_textView, c_type_, c_size)
}

func Fn_gtk_text_view_set_buffer(textView unsafe.Pointer, buffer unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_buffer := (*C.GtkTextBuffer)(unsafe.Pointer(buffer))

	C.gtk_text_view_set_buffer(c_textView, c_buffer)
}

func Fn_gtk_text_view_set_cursor_visible(textView unsafe.Pointer, setting bool) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_setting := toCBool(setting)

	C.gtk_text_view_set_cursor_visible(c_textView, c_setting)
}

func Fn_gtk_text_view_set_editable(textView unsafe.Pointer, setting bool) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_setting := toCBool(setting)

	C.gtk_text_view_set_editable(c_textView, c_setting)
}

func Fn_gtk_text_view_set_indent(textView unsafe.Pointer, indent int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_indent := (C.gint)(indent)

	C.gtk_text_view_set_indent(c_textView, c_indent)
}

func Fn_gtk_text_view_set_justification(textView unsafe.Pointer, justification int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_justification := (C.GtkJustification)(justification)

	C.gtk_text_view_set_justification(c_textView, c_justification)
}

func Fn_gtk_text_view_set_left_margin(textView unsafe.Pointer, leftMargin int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_leftMargin := (C.gint)(leftMargin)

	C.gtk_text_view_set_left_margin(c_textView, c_leftMargin)
}

func Fn_gtk_text_view_set_pixels_above_lines(textView unsafe.Pointer, pixelsAboveLines int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_pixelsAboveLines := (C.gint)(pixelsAboveLines)

	C.gtk_text_view_set_pixels_above_lines(c_textView, c_pixelsAboveLines)
}

func Fn_gtk_text_view_set_pixels_below_lines(textView unsafe.Pointer, pixelsBelowLines int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_pixelsBelowLines := (C.gint)(pixelsBelowLines)

	C.gtk_text_view_set_pixels_below_lines(c_textView, c_pixelsBelowLines)
}

func Fn_gtk_text_view_set_pixels_inside_wrap(textView unsafe.Pointer, pixelsInsideWrap int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_pixelsInsideWrap := (C.gint)(pixelsInsideWrap)

	C.gtk_text_view_set_pixels_inside_wrap(c_textView, c_pixelsInsideWrap)
}

func Fn_gtk_text_view_set_right_margin(textView unsafe.Pointer, rightMargin int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_rightMargin := (C.gint)(rightMargin)

	C.gtk_text_view_set_right_margin(c_textView, c_rightMargin)
}

func Fn_gtk_text_view_set_tabs(textView unsafe.Pointer, tabs unsafe.Pointer) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_tabs := (*C.PangoTabArray)(unsafe.Pointer(tabs))

	C.gtk_text_view_set_tabs(c_textView, c_tabs)
}

func Fn_gtk_text_view_set_wrap_mode(textView unsafe.Pointer, wrapMode int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_wrapMode := (C.GtkWrapMode)(wrapMode)

	C.gtk_text_view_set_wrap_mode(c_textView, c_wrapMode)
}

func Fn_gtk_text_view_starts_display_line(textView unsafe.Pointer, iter unsafe.Pointer) bool {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_iter := (*C.GtkTextIter)(unsafe.Pointer(iter))

	ret := C.gtk_text_view_starts_display_line(c_textView, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_text_view_window_to_buffer_coords(textView unsafe.Pointer, win int, windowX int, windowY int, bufferX *int, bufferY *int) {
	c_textView := (*C.GtkTextView)(unsafe.Pointer(textView))

	c_win := (C.GtkTextWindowType)(win)

	c_windowX := (C.gint)(windowX)

	c_windowY := (C.gint)(windowY)

	c_bufferX := (*C.gint)(unsafe.Pointer(bufferX))

	c_bufferY := (*C.gint)(unsafe.Pointer(bufferY))

	C.gtk_text_view_window_to_buffer_coords(c_textView, c_win, c_windowX, c_windowY, c_bufferX, c_bufferY)
}

func Fn_gtk_theming_engine_get_screen(engine unsafe.Pointer) unsafe.Pointer {
	c_engine := (*C.GtkThemingEngine)(unsafe.Pointer(engine))

	ret := C.gtk_theming_engine_get_screen(c_engine)

	return unsafe.Pointer(ret)
}

func Fn_gtk_theming_engine_load(name string) unsafe.Pointer {
	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.gtk_theming_engine_load(c_name)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_theming_engine_register_property : parameter 'parse_func' is callback

func Fn_gtk_toggle_button_new() unsafe.Pointer {
	ret := C.gtk_toggle_button_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_new_with_label(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_toggle_button_new_with_label(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_new_with_mnemonic(label string) unsafe.Pointer {
	c_label := (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(c_label))

	ret := C.gtk_toggle_button_new_with_mnemonic(c_label)

	return unsafe.Pointer(ret)
}

func Fn_gtk_toggle_button_get_active(toggleButton unsafe.Pointer) bool {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	ret := C.gtk_toggle_button_get_active(c_toggleButton)

	return toGoBool(ret)
}

func Fn_gtk_toggle_button_get_inconsistent(toggleButton unsafe.Pointer) bool {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	ret := C.gtk_toggle_button_get_inconsistent(c_toggleButton)

	return toGoBool(ret)
}

func Fn_gtk_toggle_button_get_mode(toggleButton unsafe.Pointer) bool {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	ret := C.gtk_toggle_button_get_mode(c_toggleButton)

	return toGoBool(ret)
}

func Fn_gtk_toggle_button_set_active(toggleButton unsafe.Pointer, isActive bool) {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	c_isActive := toCBool(isActive)

	C.gtk_toggle_button_set_active(c_toggleButton, c_isActive)
}

func Fn_gtk_toggle_button_set_inconsistent(toggleButton unsafe.Pointer, setting bool) {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	c_setting := toCBool(setting)

	C.gtk_toggle_button_set_inconsistent(c_toggleButton, c_setting)
}

func Fn_gtk_toggle_button_set_mode(toggleButton unsafe.Pointer, drawIndicator bool) {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	c_drawIndicator := toCBool(drawIndicator)

	C.gtk_toggle_button_set_mode(c_toggleButton, c_drawIndicator)
}

func Fn_gtk_toggle_button_toggled(toggleButton unsafe.Pointer) {
	c_toggleButton := (*C.GtkToggleButton)(unsafe.Pointer(toggleButton))

	C.gtk_toggle_button_toggled(c_toggleButton)
}

func Fn_gtk_toolbar_new() unsafe.Pointer {
	ret := C.gtk_toolbar_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_toolbar_get_icon_size(toolbar unsafe.Pointer) int {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	ret := C.gtk_toolbar_get_icon_size(c_toolbar)

	return (int)(ret)
}

func Fn_gtk_toolbar_get_style(toolbar unsafe.Pointer) int {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	ret := C.gtk_toolbar_get_style(c_toolbar)

	return (int)(ret)
}

func Fn_gtk_toolbar_set_icon_size(toolbar unsafe.Pointer, iconSize int) {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	c_iconSize := (C.GtkIconSize)(iconSize)

	C.gtk_toolbar_set_icon_size(c_toolbar, c_iconSize)
}

func Fn_gtk_toolbar_set_style(toolbar unsafe.Pointer, style int) {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	c_style := (C.GtkToolbarStyle)(style)

	C.gtk_toolbar_set_style(c_toolbar, c_style)
}

func Fn_gtk_toolbar_unset_icon_size(toolbar unsafe.Pointer) {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	C.gtk_toolbar_unset_icon_size(c_toolbar)
}

func Fn_gtk_toolbar_unset_style(toolbar unsafe.Pointer) {
	c_toolbar := (*C.GtkToolbar)(unsafe.Pointer(toolbar))

	C.gtk_toolbar_unset_style(c_toolbar)
}

func Fn_gtk_toplevel_accessible_get_children(accessible unsafe.Pointer) unsafe.Pointer {
	c_accessible := (*C.GtkToplevelAccessible)(unsafe.Pointer(accessible))

	ret := C.gtk_toplevel_accessible_get_children(c_accessible)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_model_filter_set_modify_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_model_filter_set_visible_func : parameter 'func' is callback

func Fn_gtk_tree_model_sort_clear_cache(treeModelSort unsafe.Pointer) {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	C.gtk_tree_model_sort_clear_cache(c_treeModelSort)
}

func Fn_gtk_tree_model_sort_convert_child_iter_to_iter(treeModelSort unsafe.Pointer, sortIter unsafe.Pointer, childIter unsafe.Pointer) bool {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	c_sortIter := (*C.GtkTreeIter)(unsafe.Pointer(sortIter))

	c_childIter := (*C.GtkTreeIter)(unsafe.Pointer(childIter))

	ret := C.gtk_tree_model_sort_convert_child_iter_to_iter(c_treeModelSort, c_sortIter, c_childIter)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_sort_convert_child_path_to_path(treeModelSort unsafe.Pointer, childPath unsafe.Pointer) unsafe.Pointer {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	c_childPath := (*C.GtkTreePath)(unsafe.Pointer(childPath))

	ret := C.gtk_tree_model_sort_convert_child_path_to_path(c_treeModelSort, c_childPath)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_convert_iter_to_child_iter(treeModelSort unsafe.Pointer, childIter unsafe.Pointer, sortedIter unsafe.Pointer) {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	c_childIter := (*C.GtkTreeIter)(unsafe.Pointer(childIter))

	c_sortedIter := (*C.GtkTreeIter)(unsafe.Pointer(sortedIter))

	C.gtk_tree_model_sort_convert_iter_to_child_iter(c_treeModelSort, c_childIter, c_sortedIter)
}

func Fn_gtk_tree_model_sort_convert_path_to_child_path(treeModelSort unsafe.Pointer, sortedPath unsafe.Pointer) unsafe.Pointer {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	c_sortedPath := (*C.GtkTreePath)(unsafe.Pointer(sortedPath))

	ret := C.gtk_tree_model_sort_convert_path_to_child_path(c_treeModelSort, c_sortedPath)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_get_model(treeModel unsafe.Pointer) unsafe.Pointer {
	c_treeModel := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModel))

	ret := C.gtk_tree_model_sort_get_model(c_treeModel)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_sort_iter_is_valid(treeModelSort unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_sort_iter_is_valid(c_treeModelSort, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_sort_reset_default_sort_func(treeModelSort unsafe.Pointer) {
	c_treeModelSort := (*C.GtkTreeModelSort)(unsafe.Pointer(treeModelSort))

	C.gtk_tree_model_sort_reset_default_sort_func(c_treeModelSort)
}

func Fn_gtk_tree_selection_count_selected_rows(selection unsafe.Pointer) int {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	ret := C.gtk_tree_selection_count_selected_rows(c_selection)

	return (int)(ret)
}

func Fn_gtk_tree_selection_get_mode(selection unsafe.Pointer) int {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	ret := C.gtk_tree_selection_get_mode(c_selection)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_selection_get_select_function : has callback return

// UNSUPPORTED : gtk_tree_selection_get_selected : parameter 'model' is non array with indirect count > 1

// UNSUPPORTED : gtk_tree_selection_get_selected_rows : parameter 'model' is non array with indirect count > 1

func Fn_gtk_tree_selection_get_tree_view(selection unsafe.Pointer) unsafe.Pointer {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	ret := C.gtk_tree_selection_get_tree_view(c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_selection_get_user_data(selection unsafe.Pointer) unsafe.Pointer {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	ret := C.gtk_tree_selection_get_user_data(c_selection)

	return (unsafe.Pointer)(ret)
}

func Fn_gtk_tree_selection_iter_is_selected(selection unsafe.Pointer, iter unsafe.Pointer) bool {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_selection_iter_is_selected(c_selection, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_selection_path_is_selected(selection unsafe.Pointer, path unsafe.Pointer) bool {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_selection_path_is_selected(c_selection, c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_selection_select_all(selection unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	C.gtk_tree_selection_select_all(c_selection)
}

func Fn_gtk_tree_selection_select_iter(selection unsafe.Pointer, iter unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_selection_select_iter(c_selection, c_iter)
}

func Fn_gtk_tree_selection_select_path(selection unsafe.Pointer, path unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_selection_select_path(c_selection, c_path)
}

func Fn_gtk_tree_selection_select_range(selection unsafe.Pointer, startPath unsafe.Pointer, endPath unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_startPath := (*C.GtkTreePath)(unsafe.Pointer(startPath))

	c_endPath := (*C.GtkTreePath)(unsafe.Pointer(endPath))

	C.gtk_tree_selection_select_range(c_selection, c_startPath, c_endPath)
}

// UNSUPPORTED : gtk_tree_selection_selected_foreach : parameter 'func' is callback

func Fn_gtk_tree_selection_set_mode(selection unsafe.Pointer, type_ int) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_type_ := (C.GtkSelectionMode)(type_)

	C.gtk_tree_selection_set_mode(c_selection, c_type_)
}

// UNSUPPORTED : gtk_tree_selection_set_select_function : parameter 'func' is callback

func Fn_gtk_tree_selection_unselect_all(selection unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	C.gtk_tree_selection_unselect_all(c_selection)
}

func Fn_gtk_tree_selection_unselect_iter(selection unsafe.Pointer, iter unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_selection_unselect_iter(c_selection, c_iter)
}

func Fn_gtk_tree_selection_unselect_path(selection unsafe.Pointer, path unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_selection_unselect_path(c_selection, c_path)
}

func Fn_gtk_tree_selection_unselect_range(selection unsafe.Pointer, startPath unsafe.Pointer, endPath unsafe.Pointer) {
	c_selection := (*C.GtkTreeSelection)(unsafe.Pointer(selection))

	c_startPath := (*C.GtkTreePath)(unsafe.Pointer(startPath))

	c_endPath := (*C.GtkTreePath)(unsafe.Pointer(endPath))

	C.gtk_tree_selection_unselect_range(c_selection, c_startPath, c_endPath)
}

func Fn_gtk_tree_store_newv(nColumns int, types []uint64) unsafe.Pointer {
	c_nColumns := (C.gint)(nColumns)

	c_types := (*C.GType)(unsafe.Pointer(&types[0]))

	ret := C.gtk_tree_store_newv(c_nColumns, c_types)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_store_append(treeStore unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_append(c_treeStore, c_iter, c_parent)
}

func Fn_gtk_tree_store_clear(treeStore unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	C.gtk_tree_store_clear(c_treeStore)
}

func Fn_gtk_tree_store_insert(treeStore unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer, position int) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	c_position := (C.gint)(position)

	C.gtk_tree_store_insert(c_treeStore, c_iter, c_parent, c_position)
}

func Fn_gtk_tree_store_insert_after(treeStore unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer, sibling unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	c_sibling := (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_after(c_treeStore, c_iter, c_parent, c_sibling)
}

func Fn_gtk_tree_store_insert_before(treeStore unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer, sibling unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	c_sibling := (*C.GtkTreeIter)(unsafe.Pointer(sibling))

	C.gtk_tree_store_insert_before(c_treeStore, c_iter, c_parent, c_sibling)
}

func Fn_gtk_tree_store_is_ancestor(treeStore unsafe.Pointer, iter unsafe.Pointer, descendant unsafe.Pointer) bool {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_descendant := (*C.GtkTreeIter)(unsafe.Pointer(descendant))

	ret := C.gtk_tree_store_is_ancestor(c_treeStore, c_iter, c_descendant)

	return toGoBool(ret)
}

func Fn_gtk_tree_store_iter_depth(treeStore unsafe.Pointer, iter unsafe.Pointer) int {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_store_iter_depth(c_treeStore, c_iter)

	return (int)(ret)
}

func Fn_gtk_tree_store_iter_is_valid(treeStore unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_store_iter_is_valid(c_treeStore, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_store_move_after(treeStore unsafe.Pointer, iter unsafe.Pointer, position unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_position := (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_after(c_treeStore, c_iter, c_position)
}

func Fn_gtk_tree_store_move_before(treeStore unsafe.Pointer, iter unsafe.Pointer, position unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_position := (*C.GtkTreeIter)(unsafe.Pointer(position))

	C.gtk_tree_store_move_before(c_treeStore, c_iter, c_position)
}

func Fn_gtk_tree_store_prepend(treeStore unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	C.gtk_tree_store_prepend(c_treeStore, c_iter, c_parent)
}

func Fn_gtk_tree_store_remove(treeStore unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_store_remove(c_treeStore, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_store_reorder(treeStore unsafe.Pointer, parent unsafe.Pointer, newOrder []int) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	c_newOrder := (*C.gint)(unsafe.Pointer(&newOrder[0]))

	C.gtk_tree_store_reorder(c_treeStore, c_parent, c_newOrder)
}

func Fn_gtk_tree_store_set_column_types(treeStore unsafe.Pointer, nColumns int, types []uint64) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_nColumns := (C.gint)(nColumns)

	c_types := (*C.GType)(unsafe.Pointer(&types[0]))

	C.gtk_tree_store_set_column_types(c_treeStore, c_nColumns, c_types)
}

func Fn_gtk_tree_store_set_valist(treeStore unsafe.Pointer, iter unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.c_gtk_tree_store_set_valist(c_treeStore, c_iter)
}

func Fn_gtk_tree_store_set_value(treeStore unsafe.Pointer, iter unsafe.Pointer, column int, value unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_column := (C.gint)(column)

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_tree_store_set_value(c_treeStore, c_iter, c_column, c_value)
}

func Fn_gtk_tree_store_swap(treeStore unsafe.Pointer, a unsafe.Pointer, b unsafe.Pointer) {
	c_treeStore := (*C.GtkTreeStore)(unsafe.Pointer(treeStore))

	c_a := (*C.GtkTreeIter)(unsafe.Pointer(a))

	c_b := (*C.GtkTreeIter)(unsafe.Pointer(b))

	C.gtk_tree_store_swap(c_treeStore, c_a, c_b)
}

func Fn_gtk_tree_view_new() unsafe.Pointer {
	ret := C.gtk_tree_view_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_new_with_model(model unsafe.Pointer) unsafe.Pointer {
	c_model := (*C.GtkTreeModel)(unsafe.Pointer(model))

	ret := C.gtk_tree_view_new_with_model(c_model)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_append_column(treeView unsafe.Pointer, column unsafe.Pointer) int {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	ret := C.gtk_tree_view_append_column(c_treeView, c_column)

	return (int)(ret)
}

func Fn_gtk_tree_view_collapse_all(treeView unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	C.gtk_tree_view_collapse_all(c_treeView)
}

func Fn_gtk_tree_view_collapse_row(treeView unsafe.Pointer, path unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_view_collapse_row(c_treeView, c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_columns_autosize(treeView unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	C.gtk_tree_view_columns_autosize(c_treeView)
}

func Fn_gtk_tree_view_create_row_drag_icon(treeView unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_view_create_row_drag_icon(c_treeView, c_path)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_enable_model_drag_dest(treeView unsafe.Pointer, targets []TargetEntry, nTargets int, actions int) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_tree_view_enable_model_drag_dest(c_treeView, c_targets, c_nTargets, c_actions)
}

func Fn_gtk_tree_view_enable_model_drag_source(treeView unsafe.Pointer, startButtonMask int, targets []TargetEntry, nTargets int, actions int) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_startButtonMask := (C.GdkModifierType)(startButtonMask)

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_tree_view_enable_model_drag_source(c_treeView, c_startButtonMask, c_targets, c_nTargets, c_actions)
}

func Fn_gtk_tree_view_expand_all(treeView unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	C.gtk_tree_view_expand_all(c_treeView)
}

func Fn_gtk_tree_view_expand_row(treeView unsafe.Pointer, path unsafe.Pointer, openAll bool) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_openAll := toCBool(openAll)

	ret := C.gtk_tree_view_expand_row(c_treeView, c_path, c_openAll)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_expand_to_path(treeView unsafe.Pointer, path unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_view_expand_to_path(c_treeView, c_path)
}

func Fn_gtk_tree_view_get_background_area(treeView unsafe.Pointer, path unsafe.Pointer, column unsafe.Pointer, rect unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_tree_view_get_background_area(c_treeView, c_path, c_column, c_rect)
}

func Fn_gtk_tree_view_get_bin_window(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_bin_window(c_treeView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_cell_area(treeView unsafe.Pointer, path unsafe.Pointer, column unsafe.Pointer, rect unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	c_rect := (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_tree_view_get_cell_area(c_treeView, c_path, c_column, c_rect)
}

func Fn_gtk_tree_view_get_column(treeView unsafe.Pointer, n int) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_n := (C.gint)(n)

	ret := C.gtk_tree_view_get_column(c_treeView, c_n)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_columns(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_columns(c_treeView)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_cursor : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_tree_view_get_dest_row_at_pos : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_tree_view_get_drag_dest_row : parameter 'path' is non array with indirect count > 1

func Fn_gtk_tree_view_get_enable_search(treeView unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_enable_search(c_treeView)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_expander_column(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_expander_column(c_treeView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_hadjustment(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_hadjustment(c_treeView)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_get_headers_visible(treeView unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_headers_visible(c_treeView)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_model(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_model(c_treeView)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_path_at_pos : parameter 'path' is non array with indirect count > 1

func Fn_gtk_tree_view_get_reorderable(treeView unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_reorderable(c_treeView)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_view_get_row_separator_func : has callback return

func Fn_gtk_tree_view_get_rules_hint(treeView unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_rules_hint(c_treeView)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_get_search_column(treeView unsafe.Pointer) int {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_search_column(c_treeView)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_view_get_search_equal_func : has callback return

// UNSUPPORTED : gtk_tree_view_get_search_position_func : has callback return

func Fn_gtk_tree_view_get_selection(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_selection(c_treeView)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_tooltip_context : parameter 'model' is non array with indirect count > 1

func Fn_gtk_tree_view_get_vadjustment(treeView unsafe.Pointer) unsafe.Pointer {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	ret := C.gtk_tree_view_get_vadjustment(c_treeView)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_tree_view_get_visible_range : parameter 'start_path' is non array with indirect count > 1

func Fn_gtk_tree_view_get_visible_rect(treeView unsafe.Pointer, visibleRect unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_visibleRect := (*C.GdkRectangle)(unsafe.Pointer(visibleRect))

	C.gtk_tree_view_get_visible_rect(c_treeView, c_visibleRect)
}

func Fn_gtk_tree_view_insert_column(treeView unsafe.Pointer, column unsafe.Pointer, position int) int {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	c_position := (C.gint)(position)

	ret := C.gtk_tree_view_insert_column(c_treeView, c_column, c_position)

	return (int)(ret)
}

func Fn_gtk_tree_view_insert_column_with_attributes(treeView unsafe.Pointer, position int, title string, cell unsafe.Pointer) int {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_position := (C.gint)(position)

	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	ret := C.c_gtk_tree_view_insert_column_with_attributes(c_treeView, c_position, c_title, c_cell)

	return (int)(ret)
}

// UNSUPPORTED : gtk_tree_view_insert_column_with_data_func : parameter 'func' is callback

// UNSUPPORTED : gtk_tree_view_is_blank_at_pos : parameter 'path' is non array with indirect count > 1

// UNSUPPORTED : gtk_tree_view_map_expanded_rows : parameter 'func' is callback

func Fn_gtk_tree_view_move_column_after(treeView unsafe.Pointer, column unsafe.Pointer, baseColumn unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	c_baseColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(baseColumn))

	C.gtk_tree_view_move_column_after(c_treeView, c_column, c_baseColumn)
}

func Fn_gtk_tree_view_remove_column(treeView unsafe.Pointer, column unsafe.Pointer) int {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	ret := C.gtk_tree_view_remove_column(c_treeView, c_column)

	return (int)(ret)
}

func Fn_gtk_tree_view_row_activated(treeView unsafe.Pointer, path unsafe.Pointer, column unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	C.gtk_tree_view_row_activated(c_treeView, c_path, c_column)
}

func Fn_gtk_tree_view_row_expanded(treeView unsafe.Pointer, path unsafe.Pointer) bool {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_view_row_expanded(c_treeView, c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_scroll_to_cell(treeView unsafe.Pointer, path unsafe.Pointer, column unsafe.Pointer, useAlign bool, rowAlign float32, colAlign float32) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	c_useAlign := toCBool(useAlign)

	c_rowAlign := (C.gfloat)(rowAlign)

	c_colAlign := (C.gfloat)(colAlign)

	C.gtk_tree_view_scroll_to_cell(c_treeView, c_path, c_column, c_useAlign, c_rowAlign, c_colAlign)
}

func Fn_gtk_tree_view_scroll_to_point(treeView unsafe.Pointer, treeX int, treeY int) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_treeX := (C.gint)(treeX)

	c_treeY := (C.gint)(treeY)

	C.gtk_tree_view_scroll_to_point(c_treeView, c_treeX, c_treeY)
}

// UNSUPPORTED : gtk_tree_view_set_column_drag_function : parameter 'func' is callback

func Fn_gtk_tree_view_set_cursor(treeView unsafe.Pointer, path unsafe.Pointer, focusColumn unsafe.Pointer, startEditing bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_focusColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(focusColumn))

	c_startEditing := toCBool(startEditing)

	C.gtk_tree_view_set_cursor(c_treeView, c_path, c_focusColumn, c_startEditing)
}

func Fn_gtk_tree_view_set_cursor_on_cell(treeView unsafe.Pointer, path unsafe.Pointer, focusColumn unsafe.Pointer, focusCell unsafe.Pointer, startEditing bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_focusColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(focusColumn))

	c_focusCell := (*C.GtkCellRenderer)(unsafe.Pointer(focusCell))

	c_startEditing := toCBool(startEditing)

	C.gtk_tree_view_set_cursor_on_cell(c_treeView, c_path, c_focusColumn, c_focusCell, c_startEditing)
}

// UNSUPPORTED : gtk_tree_view_set_destroy_count_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_drag_dest_row(treeView unsafe.Pointer, path unsafe.Pointer, pos int) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_pos := (C.GtkTreeViewDropPosition)(pos)

	C.gtk_tree_view_set_drag_dest_row(c_treeView, c_path, c_pos)
}

func Fn_gtk_tree_view_set_enable_search(treeView unsafe.Pointer, enableSearch bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_enableSearch := toCBool(enableSearch)

	C.gtk_tree_view_set_enable_search(c_treeView, c_enableSearch)
}

func Fn_gtk_tree_view_set_expander_column(treeView unsafe.Pointer, column unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (*C.GtkTreeViewColumn)(unsafe.Pointer(column))

	C.gtk_tree_view_set_expander_column(c_treeView, c_column)
}

func Fn_gtk_tree_view_set_hadjustment(treeView unsafe.Pointer, adjustment unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_tree_view_set_hadjustment(c_treeView, c_adjustment)
}

func Fn_gtk_tree_view_set_headers_clickable(treeView unsafe.Pointer, setting bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_setting := toCBool(setting)

	C.gtk_tree_view_set_headers_clickable(c_treeView, c_setting)
}

func Fn_gtk_tree_view_set_headers_visible(treeView unsafe.Pointer, headersVisible bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_headersVisible := toCBool(headersVisible)

	C.gtk_tree_view_set_headers_visible(c_treeView, c_headersVisible)
}

func Fn_gtk_tree_view_set_model(treeView unsafe.Pointer, model unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_model := (*C.GtkTreeModel)(unsafe.Pointer(model))

	C.gtk_tree_view_set_model(c_treeView, c_model)
}

func Fn_gtk_tree_view_set_reorderable(treeView unsafe.Pointer, reorderable bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_reorderable := toCBool(reorderable)

	C.gtk_tree_view_set_reorderable(c_treeView, c_reorderable)
}

// UNSUPPORTED : gtk_tree_view_set_row_separator_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_rules_hint(treeView unsafe.Pointer, setting bool) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_setting := toCBool(setting)

	C.gtk_tree_view_set_rules_hint(c_treeView, c_setting)
}

func Fn_gtk_tree_view_set_search_column(treeView unsafe.Pointer, column int) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_column := (C.gint)(column)

	C.gtk_tree_view_set_search_column(c_treeView, c_column)
}

// UNSUPPORTED : gtk_tree_view_set_search_equal_func : parameter 'search_equal_func' is callback

// UNSUPPORTED : gtk_tree_view_set_search_position_func : parameter 'func' is callback

func Fn_gtk_tree_view_set_vadjustment(treeView unsafe.Pointer, adjustment unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_tree_view_set_vadjustment(c_treeView, c_adjustment)
}

func Fn_gtk_tree_view_unset_rows_drag_dest(treeView unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	C.gtk_tree_view_unset_rows_drag_dest(c_treeView)
}

func Fn_gtk_tree_view_unset_rows_drag_source(treeView unsafe.Pointer) {
	c_treeView := (*C.GtkTreeView)(unsafe.Pointer(treeView))

	C.gtk_tree_view_unset_rows_drag_source(c_treeView)
}

func Fn_gtk_tree_view_column_new() unsafe.Pointer {
	ret := C.gtk_tree_view_column_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_new_with_attributes(title string, cell unsafe.Pointer) unsafe.Pointer {
	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	ret := C.c_gtk_tree_view_column_new_with_attributes(c_title, c_cell)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_add_attribute(treeColumn unsafe.Pointer, cellRenderer unsafe.Pointer, attribute string, column int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cellRenderer := (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer))

	c_attribute := (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_tree_view_column_add_attribute(c_treeColumn, c_cellRenderer, c_attribute, c_column)
}

func Fn_gtk_tree_view_column_cell_get_position(treeColumn unsafe.Pointer, cellRenderer unsafe.Pointer, xOffset *int, width *int) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cellRenderer := (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer))

	c_xOffset := (*C.gint)(unsafe.Pointer(xOffset))

	c_width := (*C.gint)(unsafe.Pointer(width))

	ret := C.gtk_tree_view_column_cell_get_position(c_treeColumn, c_cellRenderer, c_xOffset, c_width)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_cell_get_size(treeColumn unsafe.Pointer, cellArea unsafe.Pointer, xOffset *int, yOffset *int, width *int, height *int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cellArea := (*C.GdkRectangle)(unsafe.Pointer(cellArea))

	c_xOffset := (*C.gint)(unsafe.Pointer(xOffset))

	c_yOffset := (*C.gint)(unsafe.Pointer(yOffset))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_tree_view_column_cell_get_size(c_treeColumn, c_cellArea, c_xOffset, c_yOffset, c_width, c_height)
}

func Fn_gtk_tree_view_column_cell_is_visible(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_cell_is_visible(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_cell_set_cell_data(treeColumn unsafe.Pointer, treeModel unsafe.Pointer, iter unsafe.Pointer, isExpander bool, isExpanded bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_isExpander := toCBool(isExpander)

	c_isExpanded := toCBool(isExpanded)

	C.gtk_tree_view_column_cell_set_cell_data(c_treeColumn, c_treeModel, c_iter, c_isExpander, c_isExpanded)
}

func Fn_gtk_tree_view_column_clear(treeColumn unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	C.gtk_tree_view_column_clear(c_treeColumn)
}

func Fn_gtk_tree_view_column_clear_attributes(treeColumn unsafe.Pointer, cellRenderer unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cellRenderer := (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer))

	C.gtk_tree_view_column_clear_attributes(c_treeColumn, c_cellRenderer)
}

func Fn_gtk_tree_view_column_clicked(treeColumn unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	C.gtk_tree_view_column_clicked(c_treeColumn)
}

func Fn_gtk_tree_view_column_focus_cell(treeColumn unsafe.Pointer, cell unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	C.gtk_tree_view_column_focus_cell(c_treeColumn, c_cell)
}

func Fn_gtk_tree_view_column_get_alignment(treeColumn unsafe.Pointer) float32 {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_alignment(c_treeColumn)

	return (float32)(ret)
}

func Fn_gtk_tree_view_column_get_clickable(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_clickable(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_fixed_width(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_fixed_width(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_max_width(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_max_width(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_min_width(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_min_width(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_reorderable(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_reorderable(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_resizable(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_resizable(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_sizing(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_sizing(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_sort_column_id(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_sort_column_id(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_sort_indicator(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_sort_indicator(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_sort_order(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_sort_order(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_spacing(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_spacing(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_get_title(treeColumn unsafe.Pointer) string {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_title(c_treeColumn)

	return C.GoString(ret)
}

func Fn_gtk_tree_view_column_get_visible(treeColumn unsafe.Pointer) bool {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_visible(c_treeColumn)

	return toGoBool(ret)
}

func Fn_gtk_tree_view_column_get_widget(treeColumn unsafe.Pointer) unsafe.Pointer {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_widget(c_treeColumn)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_view_column_get_width(treeColumn unsafe.Pointer) int {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	ret := C.gtk_tree_view_column_get_width(c_treeColumn)

	return (int)(ret)
}

func Fn_gtk_tree_view_column_pack_end(treeColumn unsafe.Pointer, cell unsafe.Pointer, expand bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_expand := toCBool(expand)

	C.gtk_tree_view_column_pack_end(c_treeColumn, c_cell, c_expand)
}

func Fn_gtk_tree_view_column_pack_start(treeColumn unsafe.Pointer, cell unsafe.Pointer, expand bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cell := (*C.GtkCellRenderer)(unsafe.Pointer(cell))

	c_expand := toCBool(expand)

	C.gtk_tree_view_column_pack_start(c_treeColumn, c_cell, c_expand)
}

func Fn_gtk_tree_view_column_set_alignment(treeColumn unsafe.Pointer, xalign float32) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_xalign := (C.gfloat)(xalign)

	C.gtk_tree_view_column_set_alignment(c_treeColumn, c_xalign)
}

func Fn_gtk_tree_view_column_set_attributes(treeColumn unsafe.Pointer, cellRenderer unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_cellRenderer := (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer))

	C.c_gtk_tree_view_column_set_attributes(c_treeColumn, c_cellRenderer)
}

// UNSUPPORTED : gtk_tree_view_column_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_tree_view_column_set_clickable(treeColumn unsafe.Pointer, clickable bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_clickable := toCBool(clickable)

	C.gtk_tree_view_column_set_clickable(c_treeColumn, c_clickable)
}

func Fn_gtk_tree_view_column_set_fixed_width(treeColumn unsafe.Pointer, fixedWidth int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_fixedWidth := (C.gint)(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width(c_treeColumn, c_fixedWidth)
}

func Fn_gtk_tree_view_column_set_max_width(treeColumn unsafe.Pointer, maxWidth int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_maxWidth := (C.gint)(maxWidth)

	C.gtk_tree_view_column_set_max_width(c_treeColumn, c_maxWidth)
}

func Fn_gtk_tree_view_column_set_min_width(treeColumn unsafe.Pointer, minWidth int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_minWidth := (C.gint)(minWidth)

	C.gtk_tree_view_column_set_min_width(c_treeColumn, c_minWidth)
}

func Fn_gtk_tree_view_column_set_reorderable(treeColumn unsafe.Pointer, reorderable bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_reorderable := toCBool(reorderable)

	C.gtk_tree_view_column_set_reorderable(c_treeColumn, c_reorderable)
}

func Fn_gtk_tree_view_column_set_resizable(treeColumn unsafe.Pointer, resizable bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_resizable := toCBool(resizable)

	C.gtk_tree_view_column_set_resizable(c_treeColumn, c_resizable)
}

func Fn_gtk_tree_view_column_set_sizing(treeColumn unsafe.Pointer, type_ int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_type_ := (C.GtkTreeViewColumnSizing)(type_)

	C.gtk_tree_view_column_set_sizing(c_treeColumn, c_type_)
}

func Fn_gtk_tree_view_column_set_sort_column_id(treeColumn unsafe.Pointer, sortColumnId int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_sortColumnId := (C.gint)(sortColumnId)

	C.gtk_tree_view_column_set_sort_column_id(c_treeColumn, c_sortColumnId)
}

func Fn_gtk_tree_view_column_set_sort_indicator(treeColumn unsafe.Pointer, setting bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_setting := toCBool(setting)

	C.gtk_tree_view_column_set_sort_indicator(c_treeColumn, c_setting)
}

func Fn_gtk_tree_view_column_set_sort_order(treeColumn unsafe.Pointer, order int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_order := (C.GtkSortType)(order)

	C.gtk_tree_view_column_set_sort_order(c_treeColumn, c_order)
}

func Fn_gtk_tree_view_column_set_spacing(treeColumn unsafe.Pointer, spacing int) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_spacing := (C.gint)(spacing)

	C.gtk_tree_view_column_set_spacing(c_treeColumn, c_spacing)
}

func Fn_gtk_tree_view_column_set_title(treeColumn unsafe.Pointer, title string) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_tree_view_column_set_title(c_treeColumn, c_title)
}

func Fn_gtk_tree_view_column_set_visible(treeColumn unsafe.Pointer, visible bool) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_visible := toCBool(visible)

	C.gtk_tree_view_column_set_visible(c_treeColumn, c_visible)
}

func Fn_gtk_tree_view_column_set_widget(treeColumn unsafe.Pointer, widget unsafe.Pointer) {
	c_treeColumn := (*C.GtkTreeViewColumn)(unsafe.Pointer(treeColumn))

	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_tree_view_column_set_widget(c_treeColumn, c_widget)
}

func Fn_gtk_vbox_new(homogeneous bool, spacing int) unsafe.Pointer {
	c_homogeneous := toCBool(homogeneous)

	c_spacing := (C.gint)(spacing)

	ret := C.gtk_vbox_new(c_homogeneous, c_spacing)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vbutton_box_new() unsafe.Pointer {
	ret := C.gtk_vbutton_box_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_vpaned_new() unsafe.Pointer {
	ret := C.gtk_vpaned_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscale_new(adjustment unsafe.Pointer) unsafe.Pointer {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	ret := C.gtk_vscale_new(c_adjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscale_new_with_range(min float64, max float64, step float64) unsafe.Pointer {
	c_min := (C.gdouble)(min)

	c_max := (C.gdouble)(max)

	c_step := (C.gdouble)(step)

	ret := C.gtk_vscale_new_with_range(c_min, c_max, c_step)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vscrollbar_new(adjustment unsafe.Pointer) unsafe.Pointer {
	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	ret := C.gtk_vscrollbar_new(c_adjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_vseparator_new() unsafe.Pointer {
	ret := C.gtk_vseparator_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_new(hadjustment unsafe.Pointer, vadjustment unsafe.Pointer) unsafe.Pointer {
	c_hadjustment := (*C.GtkAdjustment)(unsafe.Pointer(hadjustment))

	c_vadjustment := (*C.GtkAdjustment)(unsafe.Pointer(vadjustment))

	ret := C.gtk_viewport_new(c_hadjustment, c_vadjustment)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_hadjustment(viewport unsafe.Pointer) unsafe.Pointer {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	ret := C.gtk_viewport_get_hadjustment(c_viewport)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_get_shadow_type(viewport unsafe.Pointer) int {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	ret := C.gtk_viewport_get_shadow_type(c_viewport)

	return (int)(ret)
}

func Fn_gtk_viewport_get_vadjustment(viewport unsafe.Pointer) unsafe.Pointer {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	ret := C.gtk_viewport_get_vadjustment(c_viewport)

	return unsafe.Pointer(ret)
}

func Fn_gtk_viewport_set_hadjustment(viewport unsafe.Pointer, adjustment unsafe.Pointer) {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_viewport_set_hadjustment(c_viewport, c_adjustment)
}

func Fn_gtk_viewport_set_shadow_type(viewport unsafe.Pointer, type_ int) {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	c_type_ := (C.GtkShadowType)(type_)

	C.gtk_viewport_set_shadow_type(c_viewport, c_type_)
}

func Fn_gtk_viewport_set_vadjustment(viewport unsafe.Pointer, adjustment unsafe.Pointer) {
	c_viewport := (*C.GtkViewport)(unsafe.Pointer(viewport))

	c_adjustment := (*C.GtkAdjustment)(unsafe.Pointer(adjustment))

	C.gtk_viewport_set_vadjustment(c_viewport, c_adjustment)
}

func Fn_gtk_widget_new(type_ uint64, firstPropertyName string) unsafe.Pointer {
	c_type_ := (C.GType)(type_)

	c_firstPropertyName := (*C.gchar)(C.CString(firstPropertyName))
	defer C.free(unsafe.Pointer(c_firstPropertyName))

	ret := C.c_gtk_widget_new(c_type_, c_firstPropertyName)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_activate(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_activate(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_add_accelerator(widget unsafe.Pointer, accelSignal string, accelGroup unsafe.Pointer, accelKey uint, accelMods int, accelFlags int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_accelSignal := (*C.gchar)(C.CString(accelSignal))
	defer C.free(unsafe.Pointer(c_accelSignal))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	c_accelFlags := (C.GtkAccelFlags)(accelFlags)

	C.gtk_widget_add_accelerator(c_widget, c_accelSignal, c_accelGroup, c_accelKey, c_accelMods, c_accelFlags)
}

func Fn_gtk_widget_add_events(widget unsafe.Pointer, events int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_events := (C.gint)(events)

	C.gtk_widget_add_events(c_widget, c_events)
}

// UNSUPPORTED : gtk_widget_add_tick_callback : parameter 'callback' is callback

func Fn_gtk_widget_child_focus(widget unsafe.Pointer, direction int) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_direction := (C.GtkDirectionType)(direction)

	ret := C.gtk_widget_child_focus(c_widget, c_direction)

	return toGoBool(ret)
}

func Fn_gtk_widget_child_notify(widget unsafe.Pointer, childProperty string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_childProperty := (*C.gchar)(C.CString(childProperty))
	defer C.free(unsafe.Pointer(c_childProperty))

	C.gtk_widget_child_notify(c_widget, c_childProperty)
}

// UNSUPPORTED : gtk_widget_class_path : parameter 'path' is non array with indirect count > 1

func Fn_gtk_widget_compute_expand(widget unsafe.Pointer, orientation int) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_orientation := (C.GtkOrientation)(orientation)

	ret := C.gtk_widget_compute_expand(c_widget, c_orientation)

	return toGoBool(ret)
}

func Fn_gtk_widget_create_pango_context(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_create_pango_context(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_create_pango_layout(widget unsafe.Pointer, text *string) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_textValue (*C.gchar)
	if text != nil {
		c_textValue = (*C.gchar)(C.CString(*text))
		defer C.free(unsafe.Pointer(c_textValue))
	}
	c_text := c_textValue

	ret := C.gtk_widget_create_pango_layout(c_widget, c_text)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_destroy(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_destroy(c_widget)
}

// UNSUPPORTED : gtk_widget_destroyed : parameter 'widget_pointer' is non array with indirect count > 1

func Fn_gtk_drag_begin(widget unsafe.Pointer, targets unsafe.Pointer, actions int, button int, event unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_targets := (*C.GtkTargetList)(unsafe.Pointer(targets))

	c_actions := (C.GdkDragAction)(actions)

	c_button := (C.gint)(button)

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	ret := C.gtk_drag_begin(c_widget, c_targets, c_actions, c_button, c_event)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_check_threshold(widget unsafe.Pointer, startX int, startY int, currentX int, currentY int) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_startX := (C.gint)(startX)

	c_startY := (C.gint)(startY)

	c_currentX := (C.gint)(currentX)

	c_currentY := (C.gint)(currentY)

	ret := C.gtk_drag_check_threshold(c_widget, c_startX, c_startY, c_currentX, c_currentY)

	return toGoBool(ret)
}

func Fn_gtk_drag_dest_find_target(widget unsafe.Pointer, context unsafe.Pointer, targetList unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_targetList := (*C.GtkTargetList)(unsafe.Pointer(targetList))

	ret := C.gtk_drag_dest_find_target(c_widget, c_context, c_targetList)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_dest_get_target_list(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_drag_dest_get_target_list(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_drag_dest_set(widget unsafe.Pointer, flags int, targets []TargetEntry, nTargets int, actions int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_flags := (C.GtkDestDefaults)(flags)

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_drag_dest_set(c_widget, c_flags, c_targets, c_nTargets, c_actions)
}

func Fn_gtk_drag_dest_set_proxy(widget unsafe.Pointer, proxyWindow unsafe.Pointer, protocol int, useCoordinates bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_proxyWindow := (*C.GdkWindow)(unsafe.Pointer(proxyWindow))

	c_protocol := (C.GdkDragProtocol)(protocol)

	c_useCoordinates := toCBool(useCoordinates)

	C.gtk_drag_dest_set_proxy(c_widget, c_proxyWindow, c_protocol, c_useCoordinates)
}

func Fn_gtk_drag_dest_set_target_list(widget unsafe.Pointer, targetList unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_targetList := (*C.GtkTargetList)(unsafe.Pointer(targetList))

	C.gtk_drag_dest_set_target_list(c_widget, c_targetList)
}

func Fn_gtk_drag_dest_unset(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_drag_dest_unset(c_widget)
}

func Fn_gtk_drag_get_data(widget unsafe.Pointer, context unsafe.Pointer, target unsafe.Pointer, time uint32) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_context := (*C.GdkDragContext)(unsafe.Pointer(context))

	c_target := (C.GdkAtom)(target)

	c_time := (C.guint32)(time)

	C.gtk_drag_get_data(c_widget, c_context, c_target, c_time)
}

func Fn_gtk_drag_highlight(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_drag_highlight(c_widget)
}

func Fn_gtk_drag_source_set(widget unsafe.Pointer, startButtonMask int, targets []TargetEntry, nTargets int, actions int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_startButtonMask := (C.GdkModifierType)(startButtonMask)

	c_targets := (*C.GtkTargetEntry)(unsafe.Pointer(&targets[0]))

	c_nTargets := (C.gint)(nTargets)

	c_actions := (C.GdkDragAction)(actions)

	C.gtk_drag_source_set(c_widget, c_startButtonMask, c_targets, c_nTargets, c_actions)
}

func Fn_gtk_drag_source_set_icon_pixbuf(widget unsafe.Pointer, pixbuf unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_pixbuf := (*C.GdkPixbuf)(unsafe.Pointer(pixbuf))

	C.gtk_drag_source_set_icon_pixbuf(c_widget, c_pixbuf)
}

func Fn_gtk_drag_source_set_icon_stock(widget unsafe.Pointer, stockId string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	C.gtk_drag_source_set_icon_stock(c_widget, c_stockId)
}

func Fn_gtk_drag_source_unset(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_drag_source_unset(c_widget)
}

func Fn_gtk_drag_unhighlight(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_drag_unhighlight(c_widget)
}

func Fn_gtk_widget_ensure_style(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_ensure_style(c_widget)
}

func Fn_gtk_widget_event(widget unsafe.Pointer, event unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	ret := C.gtk_widget_event(c_widget, c_event)

	return toGoBool(ret)
}

func Fn_gtk_widget_freeze_child_notify(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_freeze_child_notify(c_widget)
}

func Fn_gtk_widget_get_accessible(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_accessible(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_allocated_height(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_allocated_height(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_allocated_width(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_allocated_width(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_ancestor(widget unsafe.Pointer, widgetType uint64) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_widgetType := (C.GType)(widgetType)

	ret := C.gtk_widget_get_ancestor(c_widget, c_widgetType)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_child_requisition(widget unsafe.Pointer, requisition unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_requisition := (*C.GtkRequisition)(unsafe.Pointer(requisition))

	C.gtk_widget_get_child_requisition(c_widget, c_requisition)
}

func Fn_gtk_widget_get_child_visible(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_child_visible(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_clipboard(widget unsafe.Pointer, selection unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_selection := (C.GdkAtom)(selection)

	ret := C.gtk_widget_get_clipboard(c_widget, c_selection)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_composite_name(widget unsafe.Pointer) string {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_composite_name(c_widget)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_direction(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_direction(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_display(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_display(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_events(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_events(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_halign(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_halign(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_hexpand(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_hexpand(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_hexpand_set(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_hexpand_set(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_modifier_style(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_modifier_style(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_name(widget unsafe.Pointer) string {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_name(c_widget)

	return C.GoString(ret)
}

func Fn_gtk_widget_get_pango_context(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_pango_context(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_parent(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_parent(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_parent_window(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_parent_window(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_path(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_path(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_pointer(widget unsafe.Pointer, x *int, y *int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	C.gtk_widget_get_pointer(c_widget, c_x, c_y)
}

func Fn_gtk_widget_get_root_window(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_root_window(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_screen(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_screen(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_settings(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_settings(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_size_request(widget unsafe.Pointer, width *int, height *int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_widget_get_size_request(c_widget, c_width, c_height)
}

func Fn_gtk_widget_get_style(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_style(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_style_context(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_style_context(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_support_multidevice(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_support_multidevice(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_template_child(widget unsafe.Pointer, widgetType uint64, name string) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_widgetType := (C.GType)(widgetType)

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	ret := C.gtk_widget_get_template_child(c_widget, c_widgetType, c_name)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_toplevel(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_toplevel(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_get_valign(widget unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_valign(c_widget)

	return (int)(ret)
}

func Fn_gtk_widget_get_vexpand(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_vexpand(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_vexpand_set(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_vexpand_set(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_get_visual(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_get_visual(c_widget)

	return unsafe.Pointer(ret)
}

func Fn_gtk_grab_add(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_grab_add(c_widget)
}

func Fn_gtk_widget_grab_default(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_grab_default(c_widget)
}

func Fn_gtk_widget_grab_focus(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_grab_focus(c_widget)
}

func Fn_gtk_grab_remove(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_grab_remove(c_widget)
}

func Fn_gtk_widget_has_screen(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_has_screen(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_hide(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_hide(c_widget)
}

func Fn_gtk_widget_hide_on_delete(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_hide_on_delete(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_in_destruction(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_in_destruction(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_intersect(widget unsafe.Pointer, area unsafe.Pointer, intersection unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_area := (*C.GdkRectangle)(unsafe.Pointer(area))

	c_intersection := (*C.GdkRectangle)(unsafe.Pointer(intersection))

	ret := C.gtk_widget_intersect(c_widget, c_area, c_intersection)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_ancestor(widget unsafe.Pointer, ancestor unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_ancestor := (*C.GtkWidget)(unsafe.Pointer(ancestor))

	ret := C.gtk_widget_is_ancestor(c_widget, c_ancestor)

	return toGoBool(ret)
}

func Fn_gtk_widget_is_focus(widget unsafe.Pointer) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_is_focus(c_widget)

	return toGoBool(ret)
}

func Fn_gtk_widget_list_accel_closures(widget unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	ret := C.gtk_widget_list_accel_closures(c_widget)

	return unsafe.Pointer(ret)
}

// UNSUPPORTED : gtk_widget_list_action_prefixes : no array length

func Fn_gtk_widget_map(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_map(c_widget)
}

func Fn_gtk_widget_mnemonic_activate(widget unsafe.Pointer, groupCycling bool) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_groupCycling := toCBool(groupCycling)

	ret := C.gtk_widget_mnemonic_activate(c_widget, c_groupCycling)

	return toGoBool(ret)
}

func Fn_gtk_widget_modify_base(widget unsafe.Pointer, state int, color unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_widget_modify_base(c_widget, c_state, c_color)
}

func Fn_gtk_widget_modify_bg(widget unsafe.Pointer, state int, color unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_widget_modify_bg(c_widget, c_state, c_color)
}

func Fn_gtk_widget_modify_fg(widget unsafe.Pointer, state int, color unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_widget_modify_fg(c_widget, c_state, c_color)
}

func Fn_gtk_widget_modify_font(widget unsafe.Pointer, fontDesc unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_fontDesc := (*C.PangoFontDescription)(unsafe.Pointer(fontDesc))

	C.gtk_widget_modify_font(c_widget, c_fontDesc)
}

func Fn_gtk_widget_modify_style(widget unsafe.Pointer, style unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_style := (*C.GtkRcStyle)(unsafe.Pointer(style))

	C.gtk_widget_modify_style(c_widget, c_style)
}

func Fn_gtk_widget_modify_text(widget unsafe.Pointer, state int, color unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_state := (C.GtkStateType)(state)

	c_color := (*C.GdkColor)(unsafe.Pointer(color))

	C.gtk_widget_modify_text(c_widget, c_state, c_color)
}

// UNSUPPORTED : gtk_widget_path : parameter 'path' is non array with indirect count > 1

func Fn_gtk_widget_queue_compute_expand(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_queue_compute_expand(c_widget)
}

func Fn_gtk_widget_queue_draw(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_queue_draw(c_widget)
}

func Fn_gtk_widget_queue_draw_area(widget unsafe.Pointer, x int, y int, width int, height int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_widget_queue_draw_area(c_widget, c_x, c_y, c_width, c_height)
}

func Fn_gtk_widget_queue_resize(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_queue_resize(c_widget)
}

func Fn_gtk_widget_realize(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_realize(c_widget)
}

func Fn_gtk_widget_region_intersect(widget unsafe.Pointer, region unsafe.Pointer) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_region := (*C.cairo_region_t)(unsafe.Pointer(region))

	ret := C.gtk_widget_region_intersect(c_widget, c_region)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_remove_accelerator(widget unsafe.Pointer, accelGroup unsafe.Pointer, accelKey uint, accelMods int) bool {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	c_accelKey := (C.guint)(accelKey)

	c_accelMods := (C.GdkModifierType)(accelMods)

	ret := C.gtk_widget_remove_accelerator(c_widget, c_accelGroup, c_accelKey, c_accelMods)

	return toGoBool(ret)
}

func Fn_gtk_widget_render_icon(widget unsafe.Pointer, stockId string, size int, detail *string) unsafe.Pointer {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_stockId := (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(c_stockId))

	c_size := (C.GtkIconSize)(size)

	var c_detailValue (*C.gchar)
	if detail != nil {
		c_detailValue = (*C.gchar)(C.CString(*detail))
		defer C.free(unsafe.Pointer(c_detailValue))
	}
	c_detail := c_detailValue

	ret := C.gtk_widget_render_icon(c_widget, c_stockId, c_size, c_detail)

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_reparent(widget unsafe.Pointer, newParent unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_newParent := (*C.GtkWidget)(unsafe.Pointer(newParent))

	C.gtk_widget_reparent(c_widget, c_newParent)
}

func Fn_gtk_widget_reset_rc_styles(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_reset_rc_styles(c_widget)
}

func Fn_gtk_widget_send_expose(widget unsafe.Pointer, event unsafe.Pointer) int {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	ret := C.gtk_widget_send_expose(c_widget, c_event)

	return (int)(ret)
}

func Fn_gtk_widget_set_accel_path(widget unsafe.Pointer, accelPath *string, accelGroup unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	var c_accelPathValue (*C.gchar)
	if accelPath != nil {
		c_accelPathValue = (*C.gchar)(C.CString(*accelPath))
		defer C.free(unsafe.Pointer(c_accelPathValue))
	}
	c_accelPath := c_accelPathValue

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_widget_set_accel_path(c_widget, c_accelPath, c_accelGroup)
}

func Fn_gtk_widget_set_app_paintable(widget unsafe.Pointer, appPaintable bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_appPaintable := toCBool(appPaintable)

	C.gtk_widget_set_app_paintable(c_widget, c_appPaintable)
}

func Fn_gtk_widget_set_child_visible(widget unsafe.Pointer, isVisible bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_isVisible := toCBool(isVisible)

	C.gtk_widget_set_child_visible(c_widget, c_isVisible)
}

func Fn_gtk_widget_set_composite_name(widget unsafe.Pointer, name string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_set_composite_name(c_widget, c_name)
}

func Fn_gtk_widget_set_direction(widget unsafe.Pointer, dir int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_dir := (C.GtkTextDirection)(dir)

	C.gtk_widget_set_direction(c_widget, c_dir)
}

func Fn_gtk_widget_set_double_buffered(widget unsafe.Pointer, doubleBuffered bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_doubleBuffered := toCBool(doubleBuffered)

	C.gtk_widget_set_double_buffered(c_widget, c_doubleBuffered)
}

func Fn_gtk_widget_set_events(widget unsafe.Pointer, events int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_events := (C.gint)(events)

	C.gtk_widget_set_events(c_widget, c_events)
}

func Fn_gtk_widget_set_halign(widget unsafe.Pointer, align int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_align := (C.GtkAlign)(align)

	C.gtk_widget_set_halign(c_widget, c_align)
}

func Fn_gtk_widget_set_hexpand(widget unsafe.Pointer, expand bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_expand := toCBool(expand)

	C.gtk_widget_set_hexpand(c_widget, c_expand)
}

func Fn_gtk_widget_set_hexpand_set(widget unsafe.Pointer, set bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_set := toCBool(set)

	C.gtk_widget_set_hexpand_set(c_widget, c_set)
}

func Fn_gtk_widget_set_name(widget unsafe.Pointer, name string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_name := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_widget_set_name(c_widget, c_name)
}

func Fn_gtk_widget_set_parent(widget unsafe.Pointer, parent unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_parent := (*C.GtkWidget)(unsafe.Pointer(parent))

	C.gtk_widget_set_parent(c_widget, c_parent)
}

func Fn_gtk_widget_set_parent_window(widget unsafe.Pointer, parentWindow unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_parentWindow := (*C.GdkWindow)(unsafe.Pointer(parentWindow))

	C.gtk_widget_set_parent_window(c_widget, c_parentWindow)
}

func Fn_gtk_widget_set_redraw_on_allocate(widget unsafe.Pointer, redrawOnAllocate bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_redrawOnAllocate := toCBool(redrawOnAllocate)

	C.gtk_widget_set_redraw_on_allocate(c_widget, c_redrawOnAllocate)
}

func Fn_gtk_widget_set_sensitive(widget unsafe.Pointer, sensitive bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_sensitive := toCBool(sensitive)

	C.gtk_widget_set_sensitive(c_widget, c_sensitive)
}

func Fn_gtk_widget_set_size_request(widget unsafe.Pointer, width int, height int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_widget_set_size_request(c_widget, c_width, c_height)
}

func Fn_gtk_widget_set_state(widget unsafe.Pointer, state int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_state := (C.GtkStateType)(state)

	C.gtk_widget_set_state(c_widget, c_state)
}

func Fn_gtk_widget_set_style(widget unsafe.Pointer, style unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_style := (*C.GtkStyle)(unsafe.Pointer(style))

	C.gtk_widget_set_style(c_widget, c_style)
}

func Fn_gtk_widget_set_valign(widget unsafe.Pointer, align int) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_align := (C.GtkAlign)(align)

	C.gtk_widget_set_valign(c_widget, c_align)
}

func Fn_gtk_widget_set_vexpand(widget unsafe.Pointer, expand bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_expand := toCBool(expand)

	C.gtk_widget_set_vexpand(c_widget, c_expand)
}

func Fn_gtk_widget_set_vexpand_set(widget unsafe.Pointer, set bool) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_set := toCBool(set)

	C.gtk_widget_set_vexpand_set(c_widget, c_set)
}

func Fn_gtk_widget_set_visual(widget unsafe.Pointer, visual unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_visual := (*C.GdkVisual)(unsafe.Pointer(visual))

	C.gtk_widget_set_visual(c_widget, c_visual)
}

func Fn_gtk_widget_show(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_show(c_widget)
}

func Fn_gtk_widget_show_all(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_show_all(c_widget)
}

func Fn_gtk_widget_show_now(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_show_now(c_widget)
}

func Fn_gtk_widget_size_allocate(widget unsafe.Pointer, allocation *gdk.Rectangle) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_allocation := (*C.GtkAllocation)(unsafe.Pointer(allocation))

	C.gtk_widget_size_allocate(c_widget, c_allocation)
}

func Fn_gtk_widget_size_request(widget unsafe.Pointer, requisition unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_requisition := (*C.GtkRequisition)(unsafe.Pointer(requisition))

	C.gtk_widget_size_request(c_widget, c_requisition)
}

func Fn_gtk_widget_style_get(widget unsafe.Pointer, firstPropertyName string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_firstPropertyName := (*C.gchar)(C.CString(firstPropertyName))
	defer C.free(unsafe.Pointer(c_firstPropertyName))

	C.c_gtk_widget_style_get(c_widget, c_firstPropertyName)
}

func Fn_gtk_widget_style_get_property(widget unsafe.Pointer, propertyName string, value unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_propertyName := (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(c_propertyName))

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_widget_style_get_property(c_widget, c_propertyName, c_value)
}

func Fn_gtk_widget_style_get_valist(widget unsafe.Pointer, firstPropertyName string) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	c_firstPropertyName := (*C.gchar)(C.CString(firstPropertyName))
	defer C.free(unsafe.Pointer(c_firstPropertyName))

	C.c_gtk_widget_style_get_valist(c_widget, c_firstPropertyName)
}

func Fn_gtk_widget_thaw_child_notify(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_thaw_child_notify(c_widget)
}

func Fn_gtk_widget_translate_coordinates(srcWidget unsafe.Pointer, destWidget unsafe.Pointer, srcX int, srcY int, destX *int, destY *int) bool {
	c_srcWidget := (*C.GtkWidget)(unsafe.Pointer(srcWidget))

	c_destWidget := (*C.GtkWidget)(unsafe.Pointer(destWidget))

	c_srcX := (C.gint)(srcX)

	c_srcY := (C.gint)(srcY)

	c_destX := (*C.gint)(unsafe.Pointer(destX))

	c_destY := (*C.gint)(unsafe.Pointer(destY))

	ret := C.gtk_widget_translate_coordinates(c_srcWidget, c_destWidget, c_srcX, c_srcY, c_destX, c_destY)

	return toGoBool(ret)
}

func Fn_gtk_widget_unmap(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_unmap(c_widget)
}

func Fn_gtk_widget_unparent(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_unparent(c_widget)
}

func Fn_gtk_widget_unrealize(widget unsafe.Pointer) {
	c_widget := (*C.GtkWidget)(unsafe.Pointer(widget))

	C.gtk_widget_unrealize(c_widget)
}

func Fn_gtk_widget_get_default_direction() int {
	ret := C.gtk_widget_get_default_direction()

	return (int)(ret)
}

func Fn_gtk_widget_get_default_style() unsafe.Pointer {
	ret := C.gtk_widget_get_default_style()

	return unsafe.Pointer(ret)
}

func Fn_gtk_widget_pop_composite_child() {
	C.gtk_widget_pop_composite_child()
}

func Fn_gtk_widget_push_composite_child() {
	C.gtk_widget_push_composite_child()
}

func Fn_gtk_widget_set_default_direction(dir int) {
	c_dir := (C.GtkTextDirection)(dir)

	C.gtk_widget_set_default_direction(c_dir)
}

func Fn_gtk_window_new(type_ int) unsafe.Pointer {
	c_type_ := (C.GtkWindowType)(type_)

	ret := C.gtk_window_new(c_type_)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_activate_default(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_activate_default(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_activate_focus(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_activate_focus(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_add_accel_group(window unsafe.Pointer, accelGroup unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_window_add_accel_group(c_window, c_accelGroup)
}

func Fn_gtk_window_add_mnemonic(window unsafe.Pointer, keyval uint, target unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_keyval := (C.guint)(keyval)

	c_target := (*C.GtkWidget)(unsafe.Pointer(target))

	C.gtk_window_add_mnemonic(c_window, c_keyval, c_target)
}

func Fn_gtk_window_begin_move_drag(window unsafe.Pointer, button int, rootX int, rootY int, timestamp uint32) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_button := (C.gint)(button)

	c_rootX := (C.gint)(rootX)

	c_rootY := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_begin_move_drag(c_window, c_button, c_rootX, c_rootY, c_timestamp)
}

func Fn_gtk_window_begin_resize_drag(window unsafe.Pointer, edge int, button int, rootX int, rootY int, timestamp uint32) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_edge := (C.GdkWindowEdge)(edge)

	c_button := (C.gint)(button)

	c_rootX := (C.gint)(rootX)

	c_rootY := (C.gint)(rootY)

	c_timestamp := (C.guint32)(timestamp)

	C.gtk_window_begin_resize_drag(c_window, c_edge, c_button, c_rootX, c_rootY, c_timestamp)
}

func Fn_gtk_window_deiconify(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_deiconify(c_window)
}

func Fn_gtk_window_fullscreen(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_fullscreen(c_window)
}

func Fn_gtk_window_get_decorated(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_decorated(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_default_size(window unsafe.Pointer, width *int, height *int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_window_get_default_size(c_window, c_width, c_height)
}

func Fn_gtk_window_get_destroy_with_parent(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_destroy_with_parent(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_focus(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_focus(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_gravity(window unsafe.Pointer) int {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_gravity(c_window)

	return (int)(ret)
}

func Fn_gtk_window_get_icon(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_icon(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_icon_list(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_icon_list(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_mnemonic_modifier(window unsafe.Pointer) int {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_mnemonic_modifier(c_window)

	return (int)(ret)
}

func Fn_gtk_window_get_modal(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_modal(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_position(window unsafe.Pointer, rootX *int, rootY *int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_rootX := (*C.gint)(unsafe.Pointer(rootX))

	c_rootY := (*C.gint)(unsafe.Pointer(rootY))

	C.gtk_window_get_position(c_window, c_rootX, c_rootY)
}

func Fn_gtk_window_get_resizable(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_resizable(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_role(window unsafe.Pointer) string {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_role(c_window)

	return C.GoString(ret)
}

func Fn_gtk_window_get_screen(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_screen(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_size(window unsafe.Pointer, width *int, height *int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	C.gtk_window_get_size(c_window, c_width, c_height)
}

func Fn_gtk_window_get_skip_pager_hint(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_skip_pager_hint(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_skip_taskbar_hint(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_skip_taskbar_hint(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_get_title(window unsafe.Pointer) string {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_title(c_window)

	return C.GoString(ret)
}

func Fn_gtk_window_get_transient_for(window unsafe.Pointer) unsafe.Pointer {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_transient_for(c_window)

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_get_type_hint(window unsafe.Pointer) int {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_get_type_hint(c_window)

	return (int)(ret)
}

func Fn_gtk_window_has_group(window unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	ret := C.gtk_window_has_group(c_window)

	return toGoBool(ret)
}

func Fn_gtk_window_iconify(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_iconify(c_window)
}

func Fn_gtk_window_maximize(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_maximize(c_window)
}

func Fn_gtk_window_mnemonic_activate(window unsafe.Pointer, keyval uint, modifier int) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_keyval := (C.guint)(keyval)

	c_modifier := (C.GdkModifierType)(modifier)

	ret := C.gtk_window_mnemonic_activate(c_window, c_keyval, c_modifier)

	return toGoBool(ret)
}

func Fn_gtk_window_move(window unsafe.Pointer, x int, y int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	C.gtk_window_move(c_window, c_x, c_y)
}

func Fn_gtk_window_parse_geometry(window unsafe.Pointer, geometry string) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_geometry := (*C.gchar)(C.CString(geometry))
	defer C.free(unsafe.Pointer(c_geometry))

	ret := C.gtk_window_parse_geometry(c_window, c_geometry)

	return toGoBool(ret)
}

func Fn_gtk_window_present(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_present(c_window)
}

func Fn_gtk_window_remove_accel_group(window unsafe.Pointer, accelGroup unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_accelGroup := (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup))

	C.gtk_window_remove_accel_group(c_window, c_accelGroup)
}

func Fn_gtk_window_remove_mnemonic(window unsafe.Pointer, keyval uint, target unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_keyval := (C.guint)(keyval)

	c_target := (*C.GtkWidget)(unsafe.Pointer(target))

	C.gtk_window_remove_mnemonic(c_window, c_keyval, c_target)
}

func Fn_gtk_window_reshow_with_initial_size(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_reshow_with_initial_size(c_window)
}

func Fn_gtk_window_resize(window unsafe.Pointer, width int, height int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_resize(c_window, c_width, c_height)
}

func Fn_gtk_window_set_decorated(window unsafe.Pointer, setting bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gtk_window_set_decorated(c_window, c_setting)
}

func Fn_gtk_window_set_default(window unsafe.Pointer, defaultWidget unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_defaultWidget := (*C.GtkWidget)(unsafe.Pointer(defaultWidget))

	C.gtk_window_set_default(c_window, c_defaultWidget)
}

func Fn_gtk_window_set_default_size(window unsafe.Pointer, width int, height int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_width := (C.gint)(width)

	c_height := (C.gint)(height)

	C.gtk_window_set_default_size(c_window, c_width, c_height)
}

func Fn_gtk_window_set_destroy_with_parent(window unsafe.Pointer, setting bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gtk_window_set_destroy_with_parent(c_window, c_setting)
}

func Fn_gtk_window_set_focus(window unsafe.Pointer, focus unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_focus := (*C.GtkWidget)(unsafe.Pointer(focus))

	C.gtk_window_set_focus(c_window, c_focus)
}

func Fn_gtk_window_set_geometry_hints(window unsafe.Pointer, geometryWidget unsafe.Pointer, geometry unsafe.Pointer, geomMask int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_geometryWidget := (*C.GtkWidget)(unsafe.Pointer(geometryWidget))

	c_geometry := (*C.GdkGeometry)(unsafe.Pointer(geometry))

	c_geomMask := (C.GdkWindowHints)(geomMask)

	C.gtk_window_set_geometry_hints(c_window, c_geometryWidget, c_geometry, c_geomMask)
}

func Fn_gtk_window_set_gravity(window unsafe.Pointer, gravity int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_gravity := (C.GdkGravity)(gravity)

	C.gtk_window_set_gravity(c_window, c_gravity)
}

func Fn_gtk_window_set_icon(window unsafe.Pointer, icon unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_icon := (*C.GdkPixbuf)(unsafe.Pointer(icon))

	C.gtk_window_set_icon(c_window, c_icon)
}

func Fn_gtk_window_set_icon_from_file(window unsafe.Pointer, filename string, error unsafe.Pointer) bool {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_icon_from_file(c_window, c_filename, cError)

	return toGoBool(ret)
}

func Fn_gtk_window_set_icon_list(window unsafe.Pointer, list unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_list := (*C.GList)(unsafe.Pointer(list))

	C.gtk_window_set_icon_list(c_window, c_list)
}

func Fn_gtk_window_set_mnemonic_modifier(window unsafe.Pointer, modifier int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_modifier := (C.GdkModifierType)(modifier)

	C.gtk_window_set_mnemonic_modifier(c_window, c_modifier)
}

func Fn_gtk_window_set_modal(window unsafe.Pointer, modal bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_modal := toCBool(modal)

	C.gtk_window_set_modal(c_window, c_modal)
}

func Fn_gtk_window_set_position(window unsafe.Pointer, position int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_position := (C.GtkWindowPosition)(position)

	C.gtk_window_set_position(c_window, c_position)
}

func Fn_gtk_window_set_resizable(window unsafe.Pointer, resizable bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_resizable := toCBool(resizable)

	C.gtk_window_set_resizable(c_window, c_resizable)
}

func Fn_gtk_window_set_role(window unsafe.Pointer, role string) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_role := (*C.gchar)(C.CString(role))
	defer C.free(unsafe.Pointer(c_role))

	C.gtk_window_set_role(c_window, c_role)
}

func Fn_gtk_window_set_screen(window unsafe.Pointer, screen unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_screen := (*C.GdkScreen)(unsafe.Pointer(screen))

	C.gtk_window_set_screen(c_window, c_screen)
}

func Fn_gtk_window_set_skip_pager_hint(window unsafe.Pointer, setting bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gtk_window_set_skip_pager_hint(c_window, c_setting)
}

func Fn_gtk_window_set_skip_taskbar_hint(window unsafe.Pointer, setting bool) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_setting := toCBool(setting)

	C.gtk_window_set_skip_taskbar_hint(c_window, c_setting)
}

func Fn_gtk_window_set_title(window unsafe.Pointer, title string) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_title := (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_window_set_title(c_window, c_title)
}

func Fn_gtk_window_set_transient_for(window unsafe.Pointer, parent unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_parent := (*C.GtkWindow)(unsafe.Pointer(parent))

	C.gtk_window_set_transient_for(c_window, c_parent)
}

func Fn_gtk_window_set_type_hint(window unsafe.Pointer, hint int) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_hint := (C.GdkWindowTypeHint)(hint)

	C.gtk_window_set_type_hint(c_window, c_hint)
}

func Fn_gtk_window_set_wmclass(window unsafe.Pointer, wmclassName string, wmclassClass string) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	c_wmclassName := (*C.gchar)(C.CString(wmclassName))
	defer C.free(unsafe.Pointer(c_wmclassName))

	c_wmclassClass := (*C.gchar)(C.CString(wmclassClass))
	defer C.free(unsafe.Pointer(c_wmclassClass))

	C.gtk_window_set_wmclass(c_window, c_wmclassName, c_wmclassClass)
}

func Fn_gtk_window_stick(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_stick(c_window)
}

func Fn_gtk_window_unfullscreen(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_unfullscreen(c_window)
}

func Fn_gtk_window_unmaximize(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_unmaximize(c_window)
}

func Fn_gtk_window_unstick(window unsafe.Pointer) {
	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_unstick(c_window)
}

func Fn_gtk_window_get_default_icon_list() unsafe.Pointer {
	ret := C.gtk_window_get_default_icon_list()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_list_toplevels() unsafe.Pointer {
	ret := C.gtk_window_list_toplevels()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_set_auto_startup_notification(setting bool) {
	c_setting := toCBool(setting)

	C.gtk_window_set_auto_startup_notification(c_setting)
}

func Fn_gtk_window_set_default_icon_from_file(filename string, error unsafe.Pointer) bool {
	c_filename := (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(c_filename))

	cError := (**C.GError)(error)

	ret := C.gtk_window_set_default_icon_from_file(c_filename, cError)

	return toGoBool(ret)
}

func Fn_gtk_window_set_default_icon_list(list unsafe.Pointer) {
	c_list := (*C.GList)(unsafe.Pointer(list))

	C.gtk_window_set_default_icon_list(c_list)
}

func Fn_gtk_window_group_new() unsafe.Pointer {
	ret := C.gtk_window_group_new()

	return unsafe.Pointer(ret)
}

func Fn_gtk_window_group_add_window(windowGroup unsafe.Pointer, window unsafe.Pointer) {
	c_windowGroup := (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup))

	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_group_add_window(c_windowGroup, c_window)
}

func Fn_gtk_window_group_remove_window(windowGroup unsafe.Pointer, window unsafe.Pointer) {
	c_windowGroup := (*C.GtkWindowGroup)(unsafe.Pointer(windowGroup))

	c_window := (*C.GtkWindow)(unsafe.Pointer(window))

	C.gtk_window_group_remove_window(c_windowGroup, c_window)
}

func Fn_gtk_cell_accessible_parent_activate(parent unsafe.Pointer, cell unsafe.Pointer) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	C.gtk_cell_accessible_parent_activate(c_parent, c_cell)
}

func Fn_gtk_cell_accessible_parent_edit(parent unsafe.Pointer, cell unsafe.Pointer) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	C.gtk_cell_accessible_parent_edit(c_parent, c_cell)
}

func Fn_gtk_cell_accessible_parent_expand_collapse(parent unsafe.Pointer, cell unsafe.Pointer) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	C.gtk_cell_accessible_parent_expand_collapse(c_parent, c_cell)
}

func Fn_gtk_cell_accessible_parent_get_cell_area(parent unsafe.Pointer, cell unsafe.Pointer, cellRect unsafe.Pointer) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	c_cellRect := (*C.GdkRectangle)(unsafe.Pointer(cellRect))

	C.gtk_cell_accessible_parent_get_cell_area(c_parent, c_cell, c_cellRect)
}

func Fn_gtk_cell_accessible_parent_get_cell_extents(parent unsafe.Pointer, cell unsafe.Pointer, x *int, y *int, width *int, height *int, coordType int) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	c_x := (*C.gint)(unsafe.Pointer(x))

	c_y := (*C.gint)(unsafe.Pointer(y))

	c_width := (*C.gint)(unsafe.Pointer(width))

	c_height := (*C.gint)(unsafe.Pointer(height))

	c_coordType := (C.AtkCoordType)(coordType)

	C.gtk_cell_accessible_parent_get_cell_extents(c_parent, c_cell, c_x, c_y, c_width, c_height, c_coordType)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_cell_position : blacklisted

func Fn_gtk_cell_accessible_parent_get_child_index(parent unsafe.Pointer, cell unsafe.Pointer) int {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	ret := C.gtk_cell_accessible_parent_get_child_index(c_parent, c_cell)

	return (int)(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_column_header_cells : no array length

func Fn_gtk_cell_accessible_parent_get_renderer_state(parent unsafe.Pointer, cell unsafe.Pointer) int {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	ret := C.gtk_cell_accessible_parent_get_renderer_state(c_parent, c_cell)

	return (int)(ret)
}

// UNSUPPORTED : gtk_cell_accessible_parent_get_row_header_cells : no array length

func Fn_gtk_cell_accessible_parent_grab_focus(parent unsafe.Pointer, cell unsafe.Pointer) bool {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	ret := C.gtk_cell_accessible_parent_grab_focus(c_parent, c_cell)

	return toGoBool(ret)
}

func Fn_gtk_cell_accessible_parent_update_relationset(parent unsafe.Pointer, cell unsafe.Pointer, relationset unsafe.Pointer) {
	c_parent := (*C.GtkCellAccessibleParent)(unsafe.Pointer(parent))

	c_cell := (*C.GtkCellAccessible)(unsafe.Pointer(cell))

	c_relationset := (*C.AtkRelationSet)(unsafe.Pointer(relationset))

	C.gtk_cell_accessible_parent_update_relationset(c_parent, c_cell, c_relationset)
}

func Fn_gtk_cell_editable_editing_done(cellEditable unsafe.Pointer) {
	c_cellEditable := (*C.GtkCellEditable)(unsafe.Pointer(cellEditable))

	C.gtk_cell_editable_editing_done(c_cellEditable)
}

func Fn_gtk_cell_editable_remove_widget(cellEditable unsafe.Pointer) {
	c_cellEditable := (*C.GtkCellEditable)(unsafe.Pointer(cellEditable))

	C.gtk_cell_editable_remove_widget(c_cellEditable)
}

func Fn_gtk_cell_editable_start_editing(cellEditable unsafe.Pointer, event unsafe.Pointer) {
	c_cellEditable := (*C.GtkCellEditable)(unsafe.Pointer(cellEditable))

	c_event := (*C.GdkEvent)(unsafe.Pointer(event))

	C.gtk_cell_editable_start_editing(c_cellEditable, c_event)
}

// UNSUPPORTED : gtk_cell_layout_set_cell_data_func : parameter 'func' is callback

func Fn_gtk_editable_copy_clipboard(editable unsafe.Pointer) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	C.gtk_editable_copy_clipboard(c_editable)
}

func Fn_gtk_editable_cut_clipboard(editable unsafe.Pointer) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	C.gtk_editable_cut_clipboard(c_editable)
}

func Fn_gtk_editable_delete_selection(editable unsafe.Pointer) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	C.gtk_editable_delete_selection(c_editable)
}

func Fn_gtk_editable_delete_text(editable unsafe.Pointer, startPos int, endPos int) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	C.gtk_editable_delete_text(c_editable, c_startPos, c_endPos)
}

func Fn_gtk_editable_get_chars(editable unsafe.Pointer, startPos int, endPos int) string {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	ret := C.gtk_editable_get_chars(c_editable, c_startPos, c_endPos)

	return C.GoString(ret)
}

func Fn_gtk_editable_get_editable(editable unsafe.Pointer) bool {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	ret := C.gtk_editable_get_editable(c_editable)

	return toGoBool(ret)
}

func Fn_gtk_editable_get_position(editable unsafe.Pointer) int {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	ret := C.gtk_editable_get_position(c_editable)

	return (int)(ret)
}

func Fn_gtk_editable_get_selection_bounds(editable unsafe.Pointer, startPos *int, endPos *int) bool {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_startPos := (*C.gint)(unsafe.Pointer(startPos))

	c_endPos := (*C.gint)(unsafe.Pointer(endPos))

	ret := C.gtk_editable_get_selection_bounds(c_editable, c_startPos, c_endPos)

	return toGoBool(ret)
}

func Fn_gtk_editable_insert_text(editable unsafe.Pointer, newText string, newTextLength int, position *int) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_newText := (*C.gchar)(C.CString(newText))
	defer C.free(unsafe.Pointer(c_newText))

	c_newTextLength := (C.gint)(newTextLength)

	c_position := (*C.gint)(unsafe.Pointer(position))

	C.gtk_editable_insert_text(c_editable, c_newText, c_newTextLength, c_position)
}

func Fn_gtk_editable_paste_clipboard(editable unsafe.Pointer) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	C.gtk_editable_paste_clipboard(c_editable)
}

func Fn_gtk_editable_select_region(editable unsafe.Pointer, startPos int, endPos int) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_startPos := (C.gint)(startPos)

	c_endPos := (C.gint)(endPos)

	C.gtk_editable_select_region(c_editable, c_startPos, c_endPos)
}

func Fn_gtk_editable_set_editable(editable unsafe.Pointer, isEditable bool) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_isEditable := toCBool(isEditable)

	C.gtk_editable_set_editable(c_editable, c_isEditable)
}

func Fn_gtk_editable_set_position(editable unsafe.Pointer, position int) {
	c_editable := (*C.GtkEditable)(unsafe.Pointer(editable))

	c_position := (C.gint)(position)

	C.gtk_editable_set_position(c_editable, c_position)
}

func Fn_gtk_file_chooser_get_use_preview_label(chooser unsafe.Pointer) bool {
	c_chooser := (*C.GtkFileChooser)(unsafe.Pointer(chooser))

	ret := C.gtk_file_chooser_get_use_preview_label(c_chooser)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_font_chooser_set_filter_func : parameter 'filter' is callback

// UNSUPPORTED : gtk_recent_chooser_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tree_drag_dest_drag_data_received(dragDest unsafe.Pointer, dest unsafe.Pointer, selectionData unsafe.Pointer) bool {
	c_dragDest := (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest))

	c_dest := (*C.GtkTreePath)(unsafe.Pointer(dest))

	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	ret := C.gtk_tree_drag_dest_drag_data_received(c_dragDest, c_dest, c_selectionData)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_dest_row_drop_possible(dragDest unsafe.Pointer, destPath unsafe.Pointer, selectionData unsafe.Pointer) bool {
	c_dragDest := (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest))

	c_destPath := (*C.GtkTreePath)(unsafe.Pointer(destPath))

	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	ret := C.gtk_tree_drag_dest_row_drop_possible(c_dragDest, c_destPath, c_selectionData)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_drag_data_delete(dragSource unsafe.Pointer, path unsafe.Pointer) bool {
	c_dragSource := (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_drag_source_drag_data_delete(c_dragSource, c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_drag_data_get(dragSource unsafe.Pointer, path unsafe.Pointer, selectionData unsafe.Pointer) bool {
	c_dragSource := (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_selectionData := (*C.GtkSelectionData)(unsafe.Pointer(selectionData))

	ret := C.gtk_tree_drag_source_drag_data_get(c_dragSource, c_path, c_selectionData)

	return toGoBool(ret)
}

func Fn_gtk_tree_drag_source_row_draggable(dragSource unsafe.Pointer, path unsafe.Pointer) bool {
	c_dragSource := (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_drag_source_row_draggable(c_dragSource, c_path)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_model_foreach : parameter 'func' is callback

func Fn_gtk_tree_model_get(treeModel unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.c_gtk_tree_model_get(c_treeModel, c_iter)
}

func Fn_gtk_tree_model_get_column_type(treeModel unsafe.Pointer, index int) uint64 {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_index := (C.gint)(index)

	ret := C.gtk_tree_model_get_column_type(c_treeModel, c_index)

	return (uint64)(ret)
}

func Fn_gtk_tree_model_get_flags(treeModel unsafe.Pointer) int {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	ret := C.gtk_tree_model_get_flags(c_treeModel)

	return (int)(ret)
}

func Fn_gtk_tree_model_get_iter(treeModel unsafe.Pointer, iter unsafe.Pointer, path unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	ret := C.gtk_tree_model_get_iter(c_treeModel, c_iter, c_path)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_iter_first(treeModel unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_get_iter_first(c_treeModel, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_iter_from_string(treeModel unsafe.Pointer, iter unsafe.Pointer, pathString string) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_pathString := (*C.gchar)(C.CString(pathString))
	defer C.free(unsafe.Pointer(c_pathString))

	ret := C.gtk_tree_model_get_iter_from_string(c_treeModel, c_iter, c_pathString)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_get_n_columns(treeModel unsafe.Pointer) int {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	ret := C.gtk_tree_model_get_n_columns(c_treeModel)

	return (int)(ret)
}

func Fn_gtk_tree_model_get_path(treeModel unsafe.Pointer, iter unsafe.Pointer) unsafe.Pointer {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_get_path(c_treeModel, c_iter)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_get_string_from_iter(treeModel unsafe.Pointer, iter unsafe.Pointer) string {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_get_string_from_iter(c_treeModel, c_iter)

	return C.GoString(ret)
}

func Fn_gtk_tree_model_get_valist(treeModel unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.c_gtk_tree_model_get_valist(c_treeModel, c_iter)
}

func Fn_gtk_tree_model_get_value(treeModel unsafe.Pointer, iter unsafe.Pointer, column int, value unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_column := (C.gint)(column)

	c_value := (*C.GValue)(unsafe.Pointer(value))

	C.gtk_tree_model_get_value(c_treeModel, c_iter, c_column, c_value)
}

func Fn_gtk_tree_model_iter_children(treeModel unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	ret := C.gtk_tree_model_iter_children(c_treeModel, c_iter, c_parent)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_has_child(treeModel unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_iter_has_child(c_treeModel, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_n_children(treeModel unsafe.Pointer, iter unsafe.Pointer) int {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_iter_n_children(c_treeModel, c_iter)

	return (int)(ret)
}

func Fn_gtk_tree_model_iter_next(treeModel unsafe.Pointer, iter unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	ret := C.gtk_tree_model_iter_next(c_treeModel, c_iter)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_nth_child(treeModel unsafe.Pointer, iter unsafe.Pointer, parent unsafe.Pointer, n int) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_parent := (*C.GtkTreeIter)(unsafe.Pointer(parent))

	c_n := (C.gint)(n)

	ret := C.gtk_tree_model_iter_nth_child(c_treeModel, c_iter, c_parent, c_n)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_iter_parent(treeModel unsafe.Pointer, iter unsafe.Pointer, child unsafe.Pointer) bool {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	c_child := (*C.GtkTreeIter)(unsafe.Pointer(child))

	ret := C.gtk_tree_model_iter_parent(c_treeModel, c_iter, c_child)

	return toGoBool(ret)
}

func Fn_gtk_tree_model_ref_node(treeModel unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_model_ref_node(c_treeModel, c_iter)
}

func Fn_gtk_tree_model_row_changed(treeModel unsafe.Pointer, path unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_model_row_changed(c_treeModel, c_path, c_iter)
}

func Fn_gtk_tree_model_row_deleted(treeModel unsafe.Pointer, path unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	C.gtk_tree_model_row_deleted(c_treeModel, c_path)
}

func Fn_gtk_tree_model_row_has_child_toggled(treeModel unsafe.Pointer, path unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_model_row_has_child_toggled(c_treeModel, c_path, c_iter)
}

func Fn_gtk_tree_model_row_inserted(treeModel unsafe.Pointer, path unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_path := (*C.GtkTreePath)(unsafe.Pointer(path))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_model_row_inserted(c_treeModel, c_path, c_iter)
}

func Fn_gtk_tree_model_sort_new_with_model(childModel unsafe.Pointer) unsafe.Pointer {
	c_childModel := (*C.GtkTreeModel)(unsafe.Pointer(childModel))

	ret := C.gtk_tree_model_sort_new_with_model(c_childModel)

	return unsafe.Pointer(ret)
}

func Fn_gtk_tree_model_unref_node(treeModel unsafe.Pointer, iter unsafe.Pointer) {
	c_treeModel := (*C.GtkTreeModel)(unsafe.Pointer(treeModel))

	c_iter := (*C.GtkTreeIter)(unsafe.Pointer(iter))

	C.gtk_tree_model_unref_node(c_treeModel, c_iter)
}

func Fn_gtk_tree_sortable_get_sort_column_id(sortable unsafe.Pointer, sortColumnId *int, order *int) bool {
	c_sortable := (*C.GtkTreeSortable)(unsafe.Pointer(sortable))

	c_sortColumnId := (*C.gint)(unsafe.Pointer(sortColumnId))

	c_order := (*C.GtkSortType)(unsafe.Pointer(order))

	ret := C.gtk_tree_sortable_get_sort_column_id(c_sortable, c_sortColumnId, c_order)

	return toGoBool(ret)
}

func Fn_gtk_tree_sortable_has_default_sort_func(sortable unsafe.Pointer) bool {
	c_sortable := (*C.GtkTreeSortable)(unsafe.Pointer(sortable))

	ret := C.gtk_tree_sortable_has_default_sort_func(c_sortable)

	return toGoBool(ret)
}

// UNSUPPORTED : gtk_tree_sortable_set_default_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tree_sortable_set_sort_column_id(sortable unsafe.Pointer, sortColumnId int, order int) {
	c_sortable := (*C.GtkTreeSortable)(unsafe.Pointer(sortable))

	c_sortColumnId := (C.gint)(sortColumnId)

	c_order := (C.GtkSortType)(order)

	C.gtk_tree_sortable_set_sort_column_id(c_sortable, c_sortColumnId, c_order)
}

// UNSUPPORTED : gtk_tree_sortable_set_sort_func : parameter 'sort_func' is callback

func Fn_gtk_tree_sortable_sort_column_changed(sortable unsafe.Pointer) {
	c_sortable := (*C.GtkTreeSortable)(unsafe.Pointer(sortable))

	C.gtk_tree_sortable_sort_column_changed(c_sortable)
}
