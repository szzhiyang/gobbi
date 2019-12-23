// Code generated - DO NOT EDIT.
// +build gtk_3.14

package gtk

import (
	gdk "github.com/pekim/gobbi/lib/internal/c/gdk"
	glib "github.com/pekim/gobbi/lib/internal/c/glib"
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

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
type ActivatableIface C.GtkActivatableIface
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
type CssSection C.GtkCssSection
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
type RecentInfo C.GtkRecentInfo
type RecentManagerClass C.GtkRecentManagerClass
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

func Fn_gtk_accel_groups_activate(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_accel_groups_from_object(param0 unsafe.Pointer) {}

func Fn_gtk_accelerator_get_default_mod_mask() {
	C.gtk_accelerator_get_default_mod_mask()
}

func Fn_gtk_accelerator_get_label(param0 uint, param1 int) {}

func Fn_gtk_accelerator_get_label_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) {
}

func Fn_gtk_accelerator_name(param0 uint, param1 int) {}

func Fn_gtk_accelerator_name_with_keycode(param0 unsafe.Pointer, param1 uint, param2 uint, param3 int) {
}

func Fn_gtk_accelerator_parse(param0 string, param1 *uint, param2 int) {}

func Fn_gtk_accelerator_parse_with_keycode(param0 string, param1 *uint, param2 []*uint, param3 int) {
}

func Fn_gtk_accelerator_set_default_mod_mask(param0 int) {}

func Fn_gtk_accelerator_valid(param0 uint, param1 int) {}

func Fn_gtk_alternative_dialog_button_order(param0 unsafe.Pointer) {}

func Fn_gtk_binding_entry_add_signal_from_string(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_binding_entry_add_signall(param0 unsafe.Pointer, param1 uint, param2 int, param3 string, param4 unsafe.Pointer) {
}

func Fn_gtk_binding_entry_remove(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_binding_entry_skip(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_binding_set_by_class(param0 *unsafe.Pointer) {}

func Fn_gtk_binding_set_find(param0 string) {}

func Fn_gtk_binding_set_new(param0 string) {}

func Fn_gtk_bindings_activate(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_bindings_activate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_builder_error_quark() {
	C.gtk_builder_error_quark()
}

func Fn_gtk_cairo_should_draw_window(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cairo_transform_to_window(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_check_version(param0 uint, param1 uint, param2 uint) {}

func Fn_gtk_css_provider_error_quark() {
	C.gtk_css_provider_error_quark()
}

func Fn_gtk_device_grab_add(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_device_grab_remove(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_disable_setlocale() {
	C.gtk_disable_setlocale()
}

func Fn_gtk_distribute_natural_allocation(param0 int, param1 uint, param2 unsafe.Pointer) {}

func Fn_gtk_drag_finish(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint32) {}

func Fn_gtk_drag_get_source_widget(param0 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_default(param0 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_gicon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_drag_set_icon_name(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {}

func Fn_gtk_drag_set_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_drag_set_icon_stock(param0 unsafe.Pointer, param1 string, param2 int, param3 int) {}

func Fn_gtk_drag_set_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_drag_set_icon_widget(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int) {
}

func Fn_gtk_draw_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool, param4 int, param5 bool) {
}

func Fn_gtk_events_pending() {
	C.gtk_events_pending()
}

func Fn_gtk_false() {
	C.gtk_false()
}

func Fn_gtk_file_chooser_error_quark() {
	C.gtk_file_chooser_error_quark()
}

func Fn_gtk_get_binary_age() {
	C.gtk_get_binary_age()
}

func Fn_gtk_get_current_event() {
	C.gtk_get_current_event()
}

func Fn_gtk_get_current_event_device() {
	C.gtk_get_current_event_device()
}

func Fn_gtk_get_current_event_state(param0 int) {}

func Fn_gtk_get_current_event_time() {
	C.gtk_get_current_event_time()
}

func Fn_gtk_get_debug_flags() {
	C.gtk_get_debug_flags()
}

func Fn_gtk_get_default_language() {
	C.gtk_get_default_language()
}

func Fn_gtk_get_event_widget(param0 unsafe.Pointer) {}

func Fn_gtk_get_interface_age() {
	C.gtk_get_interface_age()
}

func Fn_gtk_get_locale_direction() {
	C.gtk_get_locale_direction()
}

func Fn_gtk_get_major_version() {
	C.gtk_get_major_version()
}

func Fn_gtk_get_micro_version() {
	C.gtk_get_micro_version()
}

func Fn_gtk_get_minor_version() {
	C.gtk_get_minor_version()
}

func Fn_gtk_get_option_group(param0 bool) {}

func Fn_gtk_grab_get_current() {
	C.gtk_grab_get_current()
}

func Fn_gtk_icon_size_from_name(param0 string) {}

func Fn_gtk_icon_size_get_name(param0 int) {}

func Fn_gtk_icon_size_lookup(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_icon_size_lookup_for_settings(param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_icon_size_register(param0 string, param1 int, param2 int) {}

func Fn_gtk_icon_size_register_alias(param0 string, param1 int) {}

func Fn_gtk_icon_theme_error_quark() {
	C.gtk_icon_theme_error_quark()
}

func Fn_gtk_init(param0 *int, param1 *[]string) {}

func Fn_gtk_init_check(param0 *int, param1 *[]string) {}

func Fn_gtk_init_with_args(param0 *int, param1 *[]string, param2 string, param3 []glib.OptionEntry, param4 string) {
}

// UNSUPPORTED : key_snooper_install : has callback

func Fn_gtk_key_snooper_remove(param0 uint) {}

func Fn_gtk_main() {
	C.gtk_main()
}

func Fn_gtk_main_do_event(param0 unsafe.Pointer) {}

func Fn_gtk_main_iteration() {
	C.gtk_main_iteration()
}

func Fn_gtk_main_iteration_do(param0 bool) {}

func Fn_gtk_main_level() {
	C.gtk_main_level()
}

func Fn_gtk_main_quit() {
	C.gtk_main_quit()
}

func Fn_gtk_paint_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 bool, param8 int, param9 int, param10 int, param11 int) {
}

func Fn_gtk_paint_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_box_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
}

func Fn_gtk_paint_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_diamond(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paint_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_flat_box(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int) {
}

func Fn_gtk_paint_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_hline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paint_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 bool, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 unsafe.Pointer) {
}

func Fn_gtk_paint_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_resize_grip(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_shadow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_shadow_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int, param11 int, param12 int) {
}

func Fn_gtk_paint_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int, param10 int) {
}

func Fn_gtk_paint_spinner(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 uint, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_tab(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 unsafe.Pointer, param5 string, param6 int, param7 int, param8 int, param9 int) {
}

func Fn_gtk_paint_vline(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 unsafe.Pointer, param4 string, param5 int, param6 int, param7 int) {
}

func Fn_gtk_paper_size_get_default() {
	C.gtk_paper_size_get_default()
}

func Fn_gtk_paper_size_get_paper_sizes(param0 bool) {}

func Fn_gtk_parse_args(param0 *int, param1 *[]string) {}

func Fn_gtk_print_error_quark() {
	C.gtk_print_error_quark()
}

func Fn_gtk_print_run_page_setup_dialog(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : print_run_page_setup_dialog_async : has callback

func Fn_gtk_propagate_event(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_rc_add_default_file(param0 string) {}

func Fn_gtk_rc_find_module_in_path(param0 string) {}

func Fn_gtk_rc_find_pixmap_in_path(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string) {}

func Fn_gtk_rc_get_default_files() {
	C.gtk_rc_get_default_files()
}

func Fn_gtk_rc_get_im_module_file() {
	C.gtk_rc_get_im_module_file()
}

func Fn_gtk_rc_get_im_module_path() {
	C.gtk_rc_get_im_module_path()
}

func Fn_gtk_rc_get_module_dir() {
	C.gtk_rc_get_module_dir()
}

func Fn_gtk_rc_get_style(param0 unsafe.Pointer) {}

func Fn_gtk_rc_get_style_by_paths(param0 unsafe.Pointer, param1 string, param2 string, param3 uint64) {
}

func Fn_gtk_rc_get_theme_dir() {
	C.gtk_rc_get_theme_dir()
}

func Fn_gtk_rc_parse(param0 string) {}

func Fn_gtk_rc_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_rc_parse_color_full(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_parse_priority(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_rc_parse_state(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_rc_parse_string(param0 string) {}

func Fn_gtk_rc_property_parse_border(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_color(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_enum(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_flags(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_property_parse_requisition(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_rc_reparse_all() {
	C.gtk_rc_reparse_all()
}

func Fn_gtk_rc_reparse_all_for_settings(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_rc_reset_styles(param0 unsafe.Pointer) {}

func Fn_gtk_rc_scanner_new() {
	C.gtk_rc_scanner_new()
}

func Fn_gtk_rc_set_default_files(param0 []string) {}

func Fn_gtk_recent_chooser_error_quark() {
	C.gtk_recent_chooser_error_quark()
}

func Fn_gtk_recent_manager_error_quark() {
	C.gtk_recent_manager_error_quark()
}

func Fn_gtk_render_activity(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_arrow(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_background(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_check(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_expander(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_extension(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
}

func Fn_gtk_render_focus(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_frame(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_frame_gap(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int, param7 float64, param8 float64) {
}

func Fn_gtk_render_handle(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_icon(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
}

func Fn_gtk_render_icon_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_render_icon_surface(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 float64, param4 float64) {
}

func Fn_gtk_render_insertion_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer, param5 int, param6 int) {
}

func Fn_gtk_render_layout(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 unsafe.Pointer) {
}

func Fn_gtk_render_line(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_option(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_render_slider(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 float64, param3 float64, param4 float64, param5 float64, param6 int) {
}

func Fn_gtk_rgb_to_hsv(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
}

func Fn_gtk_selection_add_target(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint) {
}

func Fn_gtk_selection_add_targets(param0 unsafe.Pointer, param1 gdk.Atom, param2 []TargetEntry, param3 uint) {
}

func Fn_gtk_selection_clear_targets(param0 unsafe.Pointer, param1 gdk.Atom) {}

func Fn_gtk_selection_convert(param0 unsafe.Pointer, param1 gdk.Atom, param2 gdk.Atom, param3 uint32) {
}

func Fn_gtk_selection_owner_set(param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {}

func Fn_gtk_selection_owner_set_for_display(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 gdk.Atom, param3 uint32) {
}

func Fn_gtk_selection_remove_all(param0 unsafe.Pointer) {}

func Fn_gtk_set_debug_flags(param0 uint) {}

// UNSUPPORTED : show_about_dialog : has varargs

func Fn_gtk_show_uri(param0 unsafe.Pointer, param1 string, param2 uint32) {}

func Fn_gtk_stock_add(param0 []StockItem, param1 uint) {}

func Fn_gtk_stock_add_static(param0 []StockItem, param1 uint) {}

func Fn_gtk_stock_list_ids() {
	C.gtk_stock_list_ids()
}

func Fn_gtk_stock_lookup(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : stock_set_translate_func : has callback

func Fn_gtk_target_table_free(param0 []TargetEntry, param1 int) {}

func Fn_gtk_target_table_new_from_list(param0 unsafe.Pointer, param1 *int) {}

func Fn_gtk_targets_include_image(param0 []gdk.Atom, param1 int, param2 bool) {}

func Fn_gtk_targets_include_rich_text(param0 []gdk.Atom, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_targets_include_text(param0 []gdk.Atom, param1 int) {}

func Fn_gtk_targets_include_uri(param0 []gdk.Atom, param1 int) {}

func Fn_gtk_test_create_simple_window(param0 string, param1 string) {}

// UNSUPPORTED : test_create_widget : has varargs

// UNSUPPORTED : test_display_button_window : has varargs

func Fn_gtk_test_find_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_test_find_sibling(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_test_find_widget(param0 unsafe.Pointer, param1 string, param2 uint64) {}

// UNSUPPORTED : test_init : has varargs

func Fn_gtk_test_list_all_types(param0 *uint) {}

func Fn_gtk_test_register_all_types() {
	C.gtk_test_register_all_types()
}

func Fn_gtk_test_slider_get_value(param0 unsafe.Pointer) {}

func Fn_gtk_test_slider_set_perc(param0 unsafe.Pointer, param1 float64) {}

func Fn_gtk_test_spin_button_click(param0 unsafe.Pointer, param1 uint, param2 bool) {}

func Fn_gtk_test_text_get(param0 unsafe.Pointer) {}

func Fn_gtk_test_text_set(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_test_widget_click(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_test_widget_send_key(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_test_widget_wait_for_draw(param0 unsafe.Pointer) {}

func Fn_gtk_tree_get_row_drag_data(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_gtk_tree_row_reference_deleted(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_row_reference_inserted(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_row_reference_reordered(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 []int) {
}

func Fn_gtk_tree_set_row_drag_data(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_true() {
	C.gtk_true()
}

func Fn_gtk_about_dialog_new() {
	C.gtk_about_dialog_new()
}

func Fn_gtk_about_dialog_add_credit_section(param0 string, param1 []string) {}

func Fn_gtk_about_dialog_get_artists() {}

func Fn_gtk_about_dialog_get_authors() {}

func Fn_gtk_about_dialog_get_comments() {}

func Fn_gtk_about_dialog_get_copyright() {}

func Fn_gtk_about_dialog_get_documenters() {}

func Fn_gtk_about_dialog_get_license() {}

func Fn_gtk_about_dialog_get_license_type() {}

func Fn_gtk_about_dialog_get_logo() {}

func Fn_gtk_about_dialog_get_logo_icon_name() {}

func Fn_gtk_about_dialog_get_program_name() {}

func Fn_gtk_about_dialog_get_translator_credits() {}

func Fn_gtk_about_dialog_get_version() {}

func Fn_gtk_about_dialog_get_website() {}

func Fn_gtk_about_dialog_get_website_label() {}

func Fn_gtk_about_dialog_get_wrap_license() {}

func Fn_gtk_about_dialog_set_artists(param0 []string) {}

func Fn_gtk_about_dialog_set_authors(param0 []string) {}

func Fn_gtk_about_dialog_set_comments(param0 string) {}

func Fn_gtk_about_dialog_set_copyright(param0 string) {}

func Fn_gtk_about_dialog_set_documenters(param0 []string) {}

func Fn_gtk_about_dialog_set_license(param0 string) {}

func Fn_gtk_about_dialog_set_license_type(param0 int) {}

func Fn_gtk_about_dialog_set_logo(param0 unsafe.Pointer) {}

func Fn_gtk_about_dialog_set_logo_icon_name(param0 string) {}

func Fn_gtk_about_dialog_set_program_name(param0 string) {}

func Fn_gtk_about_dialog_set_translator_credits(param0 string) {}

func Fn_gtk_about_dialog_set_version(param0 string) {}

func Fn_gtk_about_dialog_set_website(param0 string) {}

func Fn_gtk_about_dialog_set_website_label(param0 string) {}

func Fn_gtk_about_dialog_set_wrap_license(param0 bool) {}

func Fn_gtk_accel_group_new() {
	C.gtk_accel_group_new()
}

func Fn_gtk_accel_group_activate(param0 uint32, param1 unsafe.Pointer, param2 uint, param3 int) {}

func Fn_gtk_accel_group_connect(param0 uint, param1 int, param2 int, param3 unsafe.Pointer) {}

func Fn_gtk_accel_group_connect_by_path(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_accel_group_disconnect(param0 unsafe.Pointer) {}

func Fn_gtk_accel_group_disconnect_key(param0 uint, param1 int) {}

// UNSUPPORTED : find : has callback

func Fn_gtk_accel_group_get_is_locked() {}

func Fn_gtk_accel_group_get_modifier_mask() {}

func Fn_gtk_accel_group_lock() {}

func Fn_gtk_accel_group_query(param0 uint, param1 int, param2 *uint) {}

func Fn_gtk_accel_group_unlock() {}

func Fn_gtk_accel_group_from_accel_closure(param0 unsafe.Pointer) {}

func Fn_gtk_accel_label_new(param0 string) {}

func Fn_gtk_accel_label_get_accel(param0 *uint, param1 int) {}

func Fn_gtk_accel_label_get_accel_widget() {}

func Fn_gtk_accel_label_get_accel_width() {}

func Fn_gtk_accel_label_refetch() {}

func Fn_gtk_accel_label_set_accel(param0 uint, param1 int) {}

func Fn_gtk_accel_label_set_accel_closure(param0 unsafe.Pointer) {}

func Fn_gtk_accel_label_set_accel_widget(param0 unsafe.Pointer) {}

func Fn_gtk_accel_map_add_entry(param0 string, param1 uint, param2 int) {}

func Fn_gtk_accel_map_add_filter(param0 string) {}

func Fn_gtk_accel_map_change_entry(param0 string, param1 uint, param2 int, param3 bool) {}

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_unfiltered : has callback

func Fn_gtk_accel_map_get() {
	C.gtk_accel_map_get()
}

func Fn_gtk_accel_map_load(param0 string) {}

func Fn_gtk_accel_map_load_fd(param0 int) {}

func Fn_gtk_accel_map_load_scanner(param0 unsafe.Pointer) {}

func Fn_gtk_accel_map_lock_path(param0 string) {}

func Fn_gtk_accel_map_lookup_entry(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_accel_map_save(param0 string) {}

func Fn_gtk_accel_map_save_fd(param0 int) {}

func Fn_gtk_accel_map_unlock_path(param0 string) {}

func Fn_gtk_accessible_connect_widget_destroyed() {}

func Fn_gtk_accessible_get_widget() {}

func Fn_gtk_accessible_set_widget(param0 unsafe.Pointer) {}

func Fn_gtk_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_action_activate() {}

func Fn_gtk_action_block_activate() {}

func Fn_gtk_action_connect_accelerator() {}

func Fn_gtk_action_create_icon(param0 int) {}

func Fn_gtk_action_create_menu() {}

func Fn_gtk_action_create_menu_item() {}

func Fn_gtk_action_create_tool_item() {}

func Fn_gtk_action_disconnect_accelerator() {}

func Fn_gtk_action_get_accel_closure() {}

func Fn_gtk_action_get_accel_path() {}

func Fn_gtk_action_get_always_show_image() {}

func Fn_gtk_action_get_gicon() {}

func Fn_gtk_action_get_icon_name() {}

func Fn_gtk_action_get_is_important() {}

func Fn_gtk_action_get_label() {}

func Fn_gtk_action_get_name() {}

func Fn_gtk_action_get_proxies() {}

func Fn_gtk_action_get_sensitive() {}

func Fn_gtk_action_get_short_label() {}

func Fn_gtk_action_get_stock_id() {}

func Fn_gtk_action_get_tooltip() {}

func Fn_gtk_action_get_visible() {}

func Fn_gtk_action_get_visible_horizontal() {}

func Fn_gtk_action_get_visible_vertical() {}

func Fn_gtk_action_is_sensitive() {}

func Fn_gtk_action_is_visible() {}

func Fn_gtk_action_set_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_action_set_accel_path(param0 string) {}

func Fn_gtk_action_set_always_show_image(param0 bool) {}

func Fn_gtk_action_set_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_action_set_icon_name(param0 string) {}

func Fn_gtk_action_set_is_important(param0 bool) {}

func Fn_gtk_action_set_label(param0 string) {}

func Fn_gtk_action_set_sensitive(param0 bool) {}

func Fn_gtk_action_set_short_label(param0 string) {}

func Fn_gtk_action_set_stock_id(param0 string) {}

func Fn_gtk_action_set_tooltip(param0 string) {}

func Fn_gtk_action_set_visible(param0 bool) {}

func Fn_gtk_action_set_visible_horizontal(param0 bool) {}

func Fn_gtk_action_set_visible_vertical(param0 bool) {}

func Fn_gtk_action_unblock_activate() {}

func Fn_gtk_action_bar_new() {
	C.gtk_action_bar_new()
}

func Fn_gtk_action_bar_get_center_widget() {}

func Fn_gtk_action_bar_pack_end(param0 unsafe.Pointer) {}

func Fn_gtk_action_bar_pack_start(param0 unsafe.Pointer) {}

func Fn_gtk_action_bar_set_center_widget(param0 unsafe.Pointer) {}

func Fn_gtk_action_group_new(param0 string) {}

func Fn_gtk_action_group_add_action(param0 unsafe.Pointer) {}

func Fn_gtk_action_group_add_action_with_accel(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_action_group_add_actions(param0 []ActionEntry, param1 uint, param2 *unsafe.Pointer) {}

// UNSUPPORTED : add_actions_full : has callback

// UNSUPPORTED : add_radio_actions : has callback

// UNSUPPORTED : add_radio_actions_full : has callback

func Fn_gtk_action_group_add_toggle_actions(param0 []ToggleActionEntry, param1 uint, param2 *unsafe.Pointer) {
}

// UNSUPPORTED : add_toggle_actions_full : has callback

func Fn_gtk_action_group_get_accel_group() {}

func Fn_gtk_action_group_get_action(param0 string) {}

func Fn_gtk_action_group_get_name() {}

func Fn_gtk_action_group_get_sensitive() {}

func Fn_gtk_action_group_get_visible() {}

func Fn_gtk_action_group_list_actions() {}

func Fn_gtk_action_group_remove_action(param0 unsafe.Pointer) {}

func Fn_gtk_action_group_set_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_action_group_set_sensitive(param0 bool) {}

// UNSUPPORTED : set_translate_func : has callback

func Fn_gtk_action_group_set_translation_domain(param0 string) {}

func Fn_gtk_action_group_set_visible(param0 bool) {}

func Fn_gtk_action_group_translate_string(param0 string) {}

func Fn_gtk_adjustment_new(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_adjustment_changed() {}

func Fn_gtk_adjustment_clamp_page(param0 float64, param1 float64) {}

func Fn_gtk_adjustment_configure(param0 float64, param1 float64, param2 float64, param3 float64, param4 float64, param5 float64) {
}

func Fn_gtk_adjustment_get_lower() {}

func Fn_gtk_adjustment_get_minimum_increment() {}

func Fn_gtk_adjustment_get_page_increment() {}

func Fn_gtk_adjustment_get_page_size() {}

func Fn_gtk_adjustment_get_step_increment() {}

func Fn_gtk_adjustment_get_upper() {}

func Fn_gtk_adjustment_get_value() {}

func Fn_gtk_adjustment_set_lower(param0 float64) {}

func Fn_gtk_adjustment_set_page_increment(param0 float64) {}

func Fn_gtk_adjustment_set_page_size(param0 float64) {}

func Fn_gtk_adjustment_set_step_increment(param0 float64) {}

func Fn_gtk_adjustment_set_upper(param0 float64) {}

func Fn_gtk_adjustment_set_value(param0 float64) {}

func Fn_gtk_adjustment_value_changed() {}

func Fn_gtk_alignment_new(param0 float32, param1 float32, param2 float32, param3 float32) {}

func Fn_gtk_alignment_get_padding(param0 *uint, param1 *uint, param2 *uint, param3 *uint) {}

func Fn_gtk_alignment_set(param0 float32, param1 float32, param2 float32, param3 float32) {}

func Fn_gtk_alignment_set_padding(param0 uint, param1 uint, param2 uint, param3 uint) {}

func Fn_gtk_app_chooser_button_new(param0 string) {}

func Fn_gtk_app_chooser_button_append_custom_item(param0 string, param1 string, param2 unsafe.Pointer) {
}

func Fn_gtk_app_chooser_button_append_separator() {}

func Fn_gtk_app_chooser_button_get_heading() {}

func Fn_gtk_app_chooser_button_get_show_default_item() {}

func Fn_gtk_app_chooser_button_get_show_dialog_item() {}

func Fn_gtk_app_chooser_button_set_active_custom_item(param0 string) {}

func Fn_gtk_app_chooser_button_set_heading(param0 string) {}

func Fn_gtk_app_chooser_button_set_show_default_item(param0 bool) {}

func Fn_gtk_app_chooser_button_set_show_dialog_item(param0 bool) {}

func Fn_gtk_app_chooser_dialog_new(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_app_chooser_dialog_new_for_content_type(param0 unsafe.Pointer, param1 int, param2 string) {
}

func Fn_gtk_app_chooser_dialog_get_heading() {}

func Fn_gtk_app_chooser_dialog_get_widget() {}

func Fn_gtk_app_chooser_dialog_set_heading(param0 string) {}

func Fn_gtk_app_chooser_widget_new(param0 string) {}

func Fn_gtk_app_chooser_widget_get_default_text() {}

func Fn_gtk_app_chooser_widget_get_show_all() {}

func Fn_gtk_app_chooser_widget_get_show_default() {}

func Fn_gtk_app_chooser_widget_get_show_fallback() {}

func Fn_gtk_app_chooser_widget_get_show_other() {}

func Fn_gtk_app_chooser_widget_get_show_recommended() {}

func Fn_gtk_app_chooser_widget_set_default_text(param0 string) {}

func Fn_gtk_app_chooser_widget_set_show_all(param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_default(param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_fallback(param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_other(param0 bool) {}

func Fn_gtk_app_chooser_widget_set_show_recommended(param0 bool) {}

func Fn_gtk_application_new(param0 string, param1 int) {}

func Fn_gtk_application_add_accelerator(param0 string, param1 string, param2 unsafe.Pointer) {}

func Fn_gtk_application_add_window(param0 unsafe.Pointer) {}

func Fn_gtk_application_get_accels_for_action(param0 string) {}

func Fn_gtk_application_get_actions_for_accel(param0 string) {}

func Fn_gtk_application_get_active_window() {}

func Fn_gtk_application_get_app_menu() {}

func Fn_gtk_application_get_menu_by_id(param0 string) {}

func Fn_gtk_application_get_menubar() {}

func Fn_gtk_application_get_window_by_id(param0 uint) {}

func Fn_gtk_application_get_windows() {}

func Fn_gtk_application_inhibit(param0 unsafe.Pointer, param1 int, param2 string) {}

func Fn_gtk_application_is_inhibited(param0 int) {}

func Fn_gtk_application_list_action_descriptions() {}

func Fn_gtk_application_prefers_app_menu() {}

func Fn_gtk_application_remove_accelerator(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_application_remove_window(param0 unsafe.Pointer) {}

func Fn_gtk_application_set_accels_for_action(param0 string, param1 []string) {}

func Fn_gtk_application_set_app_menu(param0 unsafe.Pointer) {}

func Fn_gtk_application_set_menubar(param0 unsafe.Pointer) {}

func Fn_gtk_application_uninhibit(param0 uint) {}

func Fn_gtk_application_window_new(param0 unsafe.Pointer) {}

func Fn_gtk_application_window_get_id() {}

func Fn_gtk_application_window_get_show_menubar() {}

func Fn_gtk_application_window_set_show_menubar(param0 bool) {}

func Fn_gtk_arrow_new(param0 int, param1 int) {}

func Fn_gtk_arrow_set(param0 int, param1 int) {}

func Fn_gtk_aspect_frame_new(param0 string, param1 float32, param2 float32, param3 float32, param4 bool) {
}

func Fn_gtk_aspect_frame_set(param0 float32, param1 float32, param2 float32, param3 bool) {}

func Fn_gtk_assistant_new() {
	C.gtk_assistant_new()
}

func Fn_gtk_assistant_add_action_widget(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_append_page(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_commit() {}

func Fn_gtk_assistant_get_current_page() {}

func Fn_gtk_assistant_get_n_pages() {}

func Fn_gtk_assistant_get_nth_page(param0 int) {}

func Fn_gtk_assistant_get_page_complete(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_header_image(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_side_image(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_title(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_get_page_type(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_insert_page(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_assistant_next_page() {}

func Fn_gtk_assistant_prepend_page(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_previous_page() {}

func Fn_gtk_assistant_remove_action_widget(param0 unsafe.Pointer) {}

func Fn_gtk_assistant_remove_page(param0 int) {}

func Fn_gtk_assistant_set_current_page(param0 int) {}

// UNSUPPORTED : set_forward_page_func : has callback

func Fn_gtk_assistant_set_page_complete(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_assistant_set_page_header_image(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_assistant_set_page_side_image(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_assistant_set_page_title(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_assistant_set_page_type(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_assistant_update_buttons_state() {}

func Fn_gtk_bin_get_child() {}

func Fn_gtk_box_new(param0 int, param1 int) {}

func Fn_gtk_box_get_baseline_position() {}

func Fn_gtk_box_get_center_widget() {}

func Fn_gtk_box_get_homogeneous() {}

func Fn_gtk_box_get_spacing() {}

func Fn_gtk_box_pack_end(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {}

func Fn_gtk_box_pack_start(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint) {}

func Fn_gtk_box_query_child_packing(param0 unsafe.Pointer, param1 *bool, param2 *bool, param3 *uint, param4 int) {
}

func Fn_gtk_box_reorder_child(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_box_set_baseline_position(param0 int) {}

func Fn_gtk_box_set_center_widget(param0 unsafe.Pointer) {}

func Fn_gtk_box_set_child_packing(param0 unsafe.Pointer, param1 bool, param2 bool, param3 uint, param4 int) {
}

func Fn_gtk_box_set_homogeneous(param0 bool) {}

func Fn_gtk_box_set_spacing(param0 int) {}

func Fn_gtk_builder_new() {
	C.gtk_builder_new()
}

func Fn_gtk_builder_new_from_file(param0 string) {}

func Fn_gtk_builder_new_from_resource(param0 string) {}

func Fn_gtk_builder_new_from_string(param0 string, param1 uint64) {}

// UNSUPPORTED : add_callback_symbol : has callback

// UNSUPPORTED : add_callback_symbols : has varargs

func Fn_gtk_builder_add_from_file(param0 string) {}

func Fn_gtk_builder_add_from_resource(param0 string) {}

func Fn_gtk_builder_add_from_string(param0 string, param1 uint64) {}

func Fn_gtk_builder_add_objects_from_file(param0 string, param1 []string) {}

func Fn_gtk_builder_add_objects_from_resource(param0 string, param1 []string) {}

func Fn_gtk_builder_add_objects_from_string(param0 string, param1 uint64, param2 []string) {}

func Fn_gtk_builder_connect_signals(param0 *unsafe.Pointer) {}

// UNSUPPORTED : connect_signals_full : has callback

func Fn_gtk_builder_expose_object(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_builder_extend_with_template(param0 unsafe.Pointer, param1 uint64, param2 string, param3 uint64) {
}

func Fn_gtk_builder_get_application() {}

func Fn_gtk_builder_get_object(param0 string) {}

func Fn_gtk_builder_get_objects() {}

func Fn_gtk_builder_get_translation_domain() {}

func Fn_gtk_builder_get_type_from_name(param0 string) {}

func Fn_gtk_builder_lookup_callback_symbol(param0 string) {}

func Fn_gtk_builder_set_application(param0 unsafe.Pointer) {}

func Fn_gtk_builder_set_translation_domain(param0 string) {}

func Fn_gtk_builder_value_from_string(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {}

func Fn_gtk_builder_value_from_string_type(param0 uint64, param1 string, param2 unsafe.Pointer) {}

func Fn_gtk_button_new() {
	C.gtk_button_new()
}

func Fn_gtk_button_new_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_button_new_from_stock(param0 string) {}

func Fn_gtk_button_new_with_label(param0 string) {}

func Fn_gtk_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_button_clicked() {}

func Fn_gtk_button_enter() {}

func Fn_gtk_button_get_alignment(param0 *float32, param1 *float32) {}

func Fn_gtk_button_get_always_show_image() {}

func Fn_gtk_button_get_event_window() {}

func Fn_gtk_button_get_focus_on_click() {}

func Fn_gtk_button_get_image() {}

func Fn_gtk_button_get_image_position() {}

func Fn_gtk_button_get_label() {}

func Fn_gtk_button_get_relief() {}

func Fn_gtk_button_get_use_stock() {}

func Fn_gtk_button_get_use_underline() {}

func Fn_gtk_button_leave() {}

func Fn_gtk_button_pressed() {}

func Fn_gtk_button_released() {}

func Fn_gtk_button_set_alignment(param0 float32, param1 float32) {}

func Fn_gtk_button_set_always_show_image(param0 bool) {}

func Fn_gtk_button_set_focus_on_click(param0 bool) {}

func Fn_gtk_button_set_image(param0 unsafe.Pointer) {}

func Fn_gtk_button_set_image_position(param0 int) {}

func Fn_gtk_button_set_label(param0 string) {}

func Fn_gtk_button_set_relief(param0 int) {}

func Fn_gtk_button_set_use_stock(param0 bool) {}

func Fn_gtk_button_set_use_underline(param0 bool) {}

func Fn_gtk_button_box_new(param0 int) {}

func Fn_gtk_button_box_get_child_non_homogeneous(param0 unsafe.Pointer) {}

func Fn_gtk_button_box_get_child_secondary(param0 unsafe.Pointer) {}

func Fn_gtk_button_box_get_layout() {}

func Fn_gtk_button_box_set_child_non_homogeneous(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_button_box_set_child_secondary(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_button_box_set_layout(param0 int) {}

func Fn_gtk_calendar_new() {
	C.gtk_calendar_new()
}

func Fn_gtk_calendar_clear_marks() {}

func Fn_gtk_calendar_get_date(param0 *uint, param1 *uint, param2 *uint) {}

func Fn_gtk_calendar_get_day_is_marked(param0 uint) {}

func Fn_gtk_calendar_get_detail_height_rows() {}

func Fn_gtk_calendar_get_detail_width_chars() {}

func Fn_gtk_calendar_get_display_options() {}

func Fn_gtk_calendar_mark_day(param0 uint) {}

func Fn_gtk_calendar_select_day(param0 uint) {}

func Fn_gtk_calendar_select_month(param0 uint, param1 uint) {}

// UNSUPPORTED : set_detail_func : has callback

func Fn_gtk_calendar_set_detail_height_rows(param0 int) {}

func Fn_gtk_calendar_set_detail_width_chars(param0 int) {}

func Fn_gtk_calendar_set_display_options(param0 int) {}

func Fn_gtk_calendar_unmark_day(param0 uint) {}

func Fn_gtk_cell_area_activate(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 bool) {
}

func Fn_gtk_cell_area_activate_cell(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_area_add(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_add_focus_sibling(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : add_with_properties : has varargs

func Fn_gtk_cell_area_apply_attributes(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
}

func Fn_gtk_cell_area_attribute_connect(param0 unsafe.Pointer, param1 string, param2 int) {}

func Fn_gtk_cell_area_attribute_disconnect(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_cell_area_attribute_get_column(param0 unsafe.Pointer, param1 string) {}

// UNSUPPORTED : cell_get : has varargs

func Fn_gtk_cell_area_cell_get_property(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : cell_get_valist : has va_list

// UNSUPPORTED : cell_set : has varargs

func Fn_gtk_cell_area_cell_set_property(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : cell_set_valist : has va_list

func Fn_gtk_cell_area_copy_context(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_create_context() {}

func Fn_gtk_cell_area_event(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_area_focus(param0 int) {}

// UNSUPPORTED : foreach : has callback

// UNSUPPORTED : foreach_alloc : has callback

func Fn_gtk_cell_area_get_cell_allocation(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer) {
}

func Fn_gtk_cell_area_get_cell_at_position(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int, param4 int, param5 unsafe.Pointer) {
}

func Fn_gtk_cell_area_get_current_path_string() {}

func Fn_gtk_cell_area_get_edit_widget() {}

func Fn_gtk_cell_area_get_edited_cell() {}

func Fn_gtk_cell_area_get_focus_cell() {}

func Fn_gtk_cell_area_get_focus_from_sibling(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_get_focus_siblings(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_get_preferred_height(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
}

func Fn_gtk_cell_area_get_preferred_height_for_width(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_cell_area_get_preferred_width(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int) {
}

func Fn_gtk_cell_area_get_preferred_width_for_height(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_cell_area_get_request_mode() {}

func Fn_gtk_cell_area_has_renderer(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_inner_cell_area(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_cell_area_is_activatable() {}

func Fn_gtk_cell_area_is_focus_sibling(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_area_remove(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_remove_focus_sibling(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_area_render(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int, param6 bool) {
}

func Fn_gtk_cell_area_request_renderer(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 int, param4 *int, param5 *int) {
}

func Fn_gtk_cell_area_set_focus_cell(param0 unsafe.Pointer) {}

func Fn_gtk_cell_area_stop_editing(param0 bool) {}

func Fn_gtk_cell_area_box_new() {
	C.gtk_cell_area_box_new()
}

func Fn_gtk_cell_area_box_get_spacing() {}

func Fn_gtk_cell_area_box_pack_end(param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {}

func Fn_gtk_cell_area_box_pack_start(param0 unsafe.Pointer, param1 bool, param2 bool, param3 bool) {}

func Fn_gtk_cell_area_box_set_spacing(param0 int) {}

func Fn_gtk_cell_area_context_allocate(param0 int, param1 int) {}

func Fn_gtk_cell_area_context_get_allocation(param0 *int, param1 *int) {}

func Fn_gtk_cell_area_context_get_area() {}

func Fn_gtk_cell_area_context_get_preferred_height(param0 *int, param1 *int) {}

func Fn_gtk_cell_area_context_get_preferred_height_for_width(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_cell_area_context_get_preferred_width(param0 *int, param1 *int) {}

func Fn_gtk_cell_area_context_get_preferred_width_for_height(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_cell_area_context_push_preferred_height(param0 int, param1 int) {}

func Fn_gtk_cell_area_context_push_preferred_width(param0 int, param1 int) {}

func Fn_gtk_cell_area_context_reset() {}

func Fn_gtk_cell_renderer_activate(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
}

func Fn_gtk_cell_renderer_get_aligned_area(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

func Fn_gtk_cell_renderer_get_alignment(param0 *float32, param1 *float32) {}

func Fn_gtk_cell_renderer_get_fixed_size(param0 *int, param1 *int) {}

func Fn_gtk_cell_renderer_get_padding(param0 *int, param1 *int) {}

func Fn_gtk_cell_renderer_get_preferred_height(param0 unsafe.Pointer, param1 *int, param2 *int) {}

func Fn_gtk_cell_renderer_get_preferred_height_for_width(param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_cell_renderer_get_preferred_size(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_cell_renderer_get_preferred_width(param0 unsafe.Pointer, param1 *int, param2 *int) {}

func Fn_gtk_cell_renderer_get_preferred_width_for_height(param0 unsafe.Pointer, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_cell_renderer_get_request_mode() {}

func Fn_gtk_cell_renderer_get_sensitive() {}

func Fn_gtk_cell_renderer_get_size(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *int, param3 *int, param4 *int, param5 *int) {
}

func Fn_gtk_cell_renderer_get_state(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_cell_renderer_get_visible() {}

func Fn_gtk_cell_renderer_is_activatable() {}

func Fn_gtk_cell_renderer_render(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_cell_renderer_set_alignment(param0 float32, param1 float32) {}

func Fn_gtk_cell_renderer_set_fixed_size(param0 int, param1 int) {}

func Fn_gtk_cell_renderer_set_padding(param0 int, param1 int) {}

func Fn_gtk_cell_renderer_set_sensitive(param0 bool) {}

func Fn_gtk_cell_renderer_set_visible(param0 bool) {}

func Fn_gtk_cell_renderer_start_editing(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 string, param3 unsafe.Pointer, param4 unsafe.Pointer, param5 int) {
}

func Fn_gtk_cell_renderer_stop_editing(param0 bool) {}

func Fn_gtk_cell_renderer_accel_new() {
	C.gtk_cell_renderer_accel_new()
}

func Fn_gtk_cell_renderer_combo_new() {
	C.gtk_cell_renderer_combo_new()
}

func Fn_gtk_cell_renderer_pixbuf_new() {
	C.gtk_cell_renderer_pixbuf_new()
}

func Fn_gtk_cell_renderer_progress_new() {
	C.gtk_cell_renderer_progress_new()
}

func Fn_gtk_cell_renderer_spin_new() {
	C.gtk_cell_renderer_spin_new()
}

func Fn_gtk_cell_renderer_spinner_new() {
	C.gtk_cell_renderer_spinner_new()
}

func Fn_gtk_cell_renderer_text_new() {
	C.gtk_cell_renderer_text_new()
}

func Fn_gtk_cell_renderer_text_set_fixed_height_from_font(param0 int) {}

func Fn_gtk_cell_renderer_toggle_new() {
	C.gtk_cell_renderer_toggle_new()
}

func Fn_gtk_cell_renderer_toggle_get_activatable() {}

func Fn_gtk_cell_renderer_toggle_get_active() {}

func Fn_gtk_cell_renderer_toggle_get_radio() {}

func Fn_gtk_cell_renderer_toggle_set_activatable(param0 bool) {}

func Fn_gtk_cell_renderer_toggle_set_active(param0 bool) {}

func Fn_gtk_cell_renderer_toggle_set_radio(param0 bool) {}

func Fn_gtk_cell_view_new() {
	C.gtk_cell_view_new()
}

func Fn_gtk_cell_view_new_with_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_markup(param0 string) {}

func Fn_gtk_cell_view_new_with_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_new_with_text(param0 string) {}

func Fn_gtk_cell_view_get_displayed_row() {}

func Fn_gtk_cell_view_get_draw_sensitive() {}

func Fn_gtk_cell_view_get_fit_model() {}

func Fn_gtk_cell_view_get_model() {}

func Fn_gtk_cell_view_get_size_of_row(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_background_color(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_background_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_displayed_row(param0 unsafe.Pointer) {}

func Fn_gtk_cell_view_set_draw_sensitive(param0 bool) {}

func Fn_gtk_cell_view_set_fit_model(param0 bool) {}

func Fn_gtk_cell_view_set_model(param0 unsafe.Pointer) {}

func Fn_gtk_check_button_new() {
	C.gtk_check_button_new()
}

func Fn_gtk_check_button_new_with_label(param0 string) {}

func Fn_gtk_check_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_check_menu_item_new() {
	C.gtk_check_menu_item_new()
}

func Fn_gtk_check_menu_item_new_with_label(param0 string) {}

func Fn_gtk_check_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_check_menu_item_get_active() {}

func Fn_gtk_check_menu_item_get_draw_as_radio() {}

func Fn_gtk_check_menu_item_get_inconsistent() {}

func Fn_gtk_check_menu_item_set_active(param0 bool) {}

func Fn_gtk_check_menu_item_set_draw_as_radio(param0 bool) {}

func Fn_gtk_check_menu_item_set_inconsistent(param0 bool) {}

func Fn_gtk_check_menu_item_toggled() {}

func Fn_gtk_clipboard_clear() {}

func Fn_gtk_clipboard_get_display() {}

func Fn_gtk_clipboard_get_owner() {}

// UNSUPPORTED : request_contents : has callback

// UNSUPPORTED : request_image : has callback

// UNSUPPORTED : request_rich_text : has callback

// UNSUPPORTED : request_targets : has callback

// UNSUPPORTED : request_text : has callback

// UNSUPPORTED : request_uris : has callback

func Fn_gtk_clipboard_set_can_store(param0 []TargetEntry, param1 int) {}

func Fn_gtk_clipboard_set_image(param0 unsafe.Pointer) {}

func Fn_gtk_clipboard_set_text(param0 string, param1 int) {}

// UNSUPPORTED : set_with_data : has callback

// UNSUPPORTED : set_with_owner : has callback

func Fn_gtk_clipboard_store() {}

func Fn_gtk_clipboard_wait_for_contents(param0 gdk.Atom) {}

func Fn_gtk_clipboard_wait_for_image() {}

func Fn_gtk_clipboard_wait_for_rich_text(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *uint64) {
}

func Fn_gtk_clipboard_wait_for_targets(param0 []unsafe.Pointer, param1 *int) {}

func Fn_gtk_clipboard_wait_for_text() {}

func Fn_gtk_clipboard_wait_for_uris() {}

func Fn_gtk_clipboard_wait_is_image_available() {}

func Fn_gtk_clipboard_wait_is_rich_text_available(param0 unsafe.Pointer) {}

func Fn_gtk_clipboard_wait_is_target_available(param0 gdk.Atom) {}

func Fn_gtk_clipboard_wait_is_text_available() {}

func Fn_gtk_clipboard_wait_is_uris_available() {}

func Fn_gtk_clipboard_get(param0 gdk.Atom) {}

func Fn_gtk_clipboard_get_for_display(param0 unsafe.Pointer, param1 gdk.Atom) {}

func Fn_gtk_color_button_new() {
	C.gtk_color_button_new()
}

func Fn_gtk_color_button_new_with_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_new_with_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_alpha() {}

func Fn_gtk_color_button_get_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_get_title() {}

func Fn_gtk_color_button_get_use_alpha() {}

func Fn_gtk_color_button_set_alpha(param0 uint16) {}

func Fn_gtk_color_button_set_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_set_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_button_set_title(param0 string) {}

func Fn_gtk_color_button_set_use_alpha(param0 bool) {}

func Fn_gtk_color_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_color_chooser_widget_new() {
	C.gtk_color_chooser_widget_new()
}

func Fn_gtk_color_selection_new() {
	C.gtk_color_selection_new()
}

func Fn_gtk_color_selection_get_current_alpha() {}

func Fn_gtk_color_selection_get_current_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_get_current_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_get_has_opacity_control() {}

func Fn_gtk_color_selection_get_has_palette() {}

func Fn_gtk_color_selection_get_previous_alpha() {}

func Fn_gtk_color_selection_get_previous_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_get_previous_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_is_adjusting() {}

func Fn_gtk_color_selection_set_current_alpha(param0 uint16) {}

func Fn_gtk_color_selection_set_current_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_set_current_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_set_has_opacity_control(param0 bool) {}

func Fn_gtk_color_selection_set_has_palette(param0 bool) {}

func Fn_gtk_color_selection_set_previous_alpha(param0 uint16) {}

func Fn_gtk_color_selection_set_previous_color(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_set_previous_rgba(param0 unsafe.Pointer) {}

func Fn_gtk_color_selection_palette_from_string(param0 string, param1 []unsafe.Pointer, param2 *int) {
}

func Fn_gtk_color_selection_palette_to_string(param0 []gdk.Color, param1 int) {}

// UNSUPPORTED : set_change_palette_with_screen_hook : has callback

func Fn_gtk_color_selection_dialog_new(param0 string) {}

func Fn_gtk_color_selection_dialog_get_color_selection() {}

func Fn_gtk_combo_box_new() {
	C.gtk_combo_box_new()
}

func Fn_gtk_combo_box_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_area_and_entry(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_entry() {
	C.gtk_combo_box_new_with_entry()
}

func Fn_gtk_combo_box_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_new_with_model_and_entry(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_get_active() {}

func Fn_gtk_combo_box_get_active_id() {}

func Fn_gtk_combo_box_get_active_iter(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_get_add_tearoffs() {}

func Fn_gtk_combo_box_get_button_sensitivity() {}

func Fn_gtk_combo_box_get_column_span_column() {}

func Fn_gtk_combo_box_get_entry_text_column() {}

func Fn_gtk_combo_box_get_focus_on_click() {}

func Fn_gtk_combo_box_get_has_entry() {}

func Fn_gtk_combo_box_get_id_column() {}

func Fn_gtk_combo_box_get_model() {}

func Fn_gtk_combo_box_get_popup_accessible() {}

func Fn_gtk_combo_box_get_popup_fixed_width() {}

func Fn_gtk_combo_box_get_row_separator_func() {}

func Fn_gtk_combo_box_get_row_span_column() {}

func Fn_gtk_combo_box_get_title() {}

func Fn_gtk_combo_box_get_wrap_width() {}

func Fn_gtk_combo_box_popdown() {}

func Fn_gtk_combo_box_popup() {}

func Fn_gtk_combo_box_popup_for_device(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_active(param0 int) {}

func Fn_gtk_combo_box_set_active_id(param0 string) {}

func Fn_gtk_combo_box_set_active_iter(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_add_tearoffs(param0 bool) {}

func Fn_gtk_combo_box_set_button_sensitivity(param0 int) {}

func Fn_gtk_combo_box_set_column_span_column(param0 int) {}

func Fn_gtk_combo_box_set_entry_text_column(param0 int) {}

func Fn_gtk_combo_box_set_focus_on_click(param0 bool) {}

func Fn_gtk_combo_box_set_id_column(param0 int) {}

func Fn_gtk_combo_box_set_model(param0 unsafe.Pointer) {}

func Fn_gtk_combo_box_set_popup_fixed_width(param0 bool) {}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_combo_box_set_row_span_column(param0 int) {}

func Fn_gtk_combo_box_set_title(param0 string) {}

func Fn_gtk_combo_box_set_wrap_width(param0 int) {}

func Fn_gtk_combo_box_text_new() {
	C.gtk_combo_box_text_new()
}

func Fn_gtk_combo_box_text_new_with_entry() {
	C.gtk_combo_box_text_new_with_entry()
}

func Fn_gtk_combo_box_text_append(param0 string, param1 string) {}

func Fn_gtk_combo_box_text_append_text(param0 string) {}

func Fn_gtk_combo_box_text_get_active_text() {}

func Fn_gtk_combo_box_text_insert(param0 int, param1 string, param2 string) {}

func Fn_gtk_combo_box_text_insert_text(param0 int, param1 string) {}

func Fn_gtk_combo_box_text_prepend(param0 string, param1 string) {}

func Fn_gtk_combo_box_text_prepend_text(param0 string) {}

func Fn_gtk_combo_box_text_remove(param0 int) {}

func Fn_gtk_combo_box_text_remove_all() {}

func Fn_gtk_container_add(param0 unsafe.Pointer) {}

// UNSUPPORTED : add_with_properties : has varargs

func Fn_gtk_container_check_resize() {}

// UNSUPPORTED : child_get : has varargs

func Fn_gtk_container_child_get_property(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : child_get_valist : has va_list

func Fn_gtk_container_child_notify(param0 unsafe.Pointer, param1 string) {}

// UNSUPPORTED : child_set : has varargs

func Fn_gtk_container_child_set_property(param0 unsafe.Pointer, param1 string, param2 unsafe.Pointer) {
}

// UNSUPPORTED : child_set_valist : has va_list

func Fn_gtk_container_child_type() {}

// UNSUPPORTED : forall : has callback

// UNSUPPORTED : foreach : has callback

func Fn_gtk_container_get_border_width() {}

func Fn_gtk_container_get_children() {}

func Fn_gtk_container_get_focus_chain(param0 *unsafe.Pointer) {}

func Fn_gtk_container_get_focus_child() {}

func Fn_gtk_container_get_focus_hadjustment() {}

func Fn_gtk_container_get_focus_vadjustment() {}

func Fn_gtk_container_get_path_for_child(param0 unsafe.Pointer) {}

func Fn_gtk_container_get_resize_mode() {}

func Fn_gtk_container_propagate_draw(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_container_remove(param0 unsafe.Pointer) {}

func Fn_gtk_container_resize_children() {}

func Fn_gtk_container_set_border_width(param0 uint) {}

func Fn_gtk_container_set_focus_chain(param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_child(param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_container_set_focus_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_container_set_reallocate_redraws(param0 bool) {}

func Fn_gtk_container_set_resize_mode(param0 int) {}

func Fn_gtk_container_unset_focus_chain() {}

func Fn_gtk_container_cell_accessible_new() {
	C.gtk_container_cell_accessible_new()
}

func Fn_gtk_container_cell_accessible_add_child(param0 unsafe.Pointer) {}

func Fn_gtk_container_cell_accessible_get_children() {}

func Fn_gtk_container_cell_accessible_remove_child(param0 unsafe.Pointer) {}

func Fn_gtk_css_provider_new() {
	C.gtk_css_provider_new()
}

func Fn_gtk_css_provider_load_from_data(param0 []uint8, param1 uint64) {}

func Fn_gtk_css_provider_load_from_file(param0 unsafe.Pointer) {}

func Fn_gtk_css_provider_load_from_path(param0 string) {}

func Fn_gtk_css_provider_to_string() {}

func Fn_gtk_css_provider_get_default() {
	C.gtk_css_provider_get_default()
}

func Fn_gtk_css_provider_get_named(param0 string, param1 string) {}

func Fn_gtk_dialog_new() {
	C.gtk_dialog_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_dialog_add_action_widget(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_dialog_add_button(param0 string, param1 int) {}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_dialog_get_action_area() {}

func Fn_gtk_dialog_get_content_area() {}

func Fn_gtk_dialog_get_header_bar() {}

func Fn_gtk_dialog_get_response_for_widget(param0 unsafe.Pointer) {}

func Fn_gtk_dialog_get_widget_for_response(param0 int) {}

func Fn_gtk_dialog_response(param0 int) {}

func Fn_gtk_dialog_run() {}

// UNSUPPORTED : set_alternative_button_order : has varargs

func Fn_gtk_dialog_set_alternative_button_order_from_array(param0 int, param1 []int) {}

func Fn_gtk_dialog_set_default_response(param0 int) {}

func Fn_gtk_dialog_set_response_sensitive(param0 int, param1 bool) {}

func Fn_gtk_drawing_area_new() {
	C.gtk_drawing_area_new()
}

func Fn_gtk_entry_new() {
	C.gtk_entry_new()
}

func Fn_gtk_entry_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_entry_get_activates_default() {}

func Fn_gtk_entry_get_alignment() {}

func Fn_gtk_entry_get_attributes() {}

func Fn_gtk_entry_get_buffer() {}

func Fn_gtk_entry_get_completion() {}

func Fn_gtk_entry_get_current_icon_drag_source() {}

func Fn_gtk_entry_get_cursor_hadjustment() {}

func Fn_gtk_entry_get_has_frame() {}

func Fn_gtk_entry_get_icon_activatable(param0 int) {}

func Fn_gtk_entry_get_icon_area(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_entry_get_icon_at_pos(param0 int, param1 int) {}

func Fn_gtk_entry_get_icon_gicon(param0 int) {}

func Fn_gtk_entry_get_icon_name(param0 int) {}

func Fn_gtk_entry_get_icon_pixbuf(param0 int) {}

func Fn_gtk_entry_get_icon_sensitive(param0 int) {}

func Fn_gtk_entry_get_icon_stock(param0 int) {}

func Fn_gtk_entry_get_icon_storage_type(param0 int) {}

func Fn_gtk_entry_get_icon_tooltip_markup(param0 int) {}

func Fn_gtk_entry_get_icon_tooltip_text(param0 int) {}

func Fn_gtk_entry_get_inner_border() {}

func Fn_gtk_entry_get_input_hints() {}

func Fn_gtk_entry_get_input_purpose() {}

func Fn_gtk_entry_get_invisible_char() {}

func Fn_gtk_entry_get_layout() {}

func Fn_gtk_entry_get_layout_offsets(param0 *int, param1 *int) {}

func Fn_gtk_entry_get_max_length() {}

func Fn_gtk_entry_get_max_width_chars() {}

func Fn_gtk_entry_get_overwrite_mode() {}

func Fn_gtk_entry_get_placeholder_text() {}

func Fn_gtk_entry_get_progress_fraction() {}

func Fn_gtk_entry_get_progress_pulse_step() {}

func Fn_gtk_entry_get_tabs() {}

func Fn_gtk_entry_get_text() {}

func Fn_gtk_entry_get_text_area(param0 unsafe.Pointer) {}

func Fn_gtk_entry_get_text_length() {}

func Fn_gtk_entry_get_visibility() {}

func Fn_gtk_entry_get_width_chars() {}

func Fn_gtk_entry_im_context_filter_keypress(param0 unsafe.Pointer) {}

func Fn_gtk_entry_layout_index_to_text_index(param0 int) {}

func Fn_gtk_entry_progress_pulse() {}

func Fn_gtk_entry_reset_im_context() {}

func Fn_gtk_entry_set_activates_default(param0 bool) {}

func Fn_gtk_entry_set_alignment(param0 float32) {}

func Fn_gtk_entry_set_attributes(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_completion(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_cursor_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_has_frame(param0 bool) {}

func Fn_gtk_entry_set_icon_activatable(param0 int, param1 bool) {}

func Fn_gtk_entry_set_icon_drag_source(param0 int, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_entry_set_icon_from_gicon(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_entry_set_icon_from_icon_name(param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_from_pixbuf(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_entry_set_icon_from_stock(param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_sensitive(param0 int, param1 bool) {}

func Fn_gtk_entry_set_icon_tooltip_markup(param0 int, param1 string) {}

func Fn_gtk_entry_set_icon_tooltip_text(param0 int, param1 string) {}

func Fn_gtk_entry_set_inner_border(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_input_hints(param0 int) {}

func Fn_gtk_entry_set_input_purpose(param0 int) {}

func Fn_gtk_entry_set_invisible_char(param0 rune) {}

func Fn_gtk_entry_set_max_length(param0 int) {}

func Fn_gtk_entry_set_max_width_chars(param0 int) {}

func Fn_gtk_entry_set_overwrite_mode(param0 bool) {}

func Fn_gtk_entry_set_placeholder_text(param0 string) {}

func Fn_gtk_entry_set_progress_fraction(param0 float64) {}

func Fn_gtk_entry_set_progress_pulse_step(param0 float64) {}

func Fn_gtk_entry_set_tabs(param0 unsafe.Pointer) {}

func Fn_gtk_entry_set_text(param0 string) {}

func Fn_gtk_entry_set_visibility(param0 bool) {}

func Fn_gtk_entry_set_width_chars(param0 int) {}

func Fn_gtk_entry_text_index_to_layout_index(param0 int) {}

func Fn_gtk_entry_unset_invisible_char() {}

func Fn_gtk_entry_buffer_new(param0 string, param1 int) {}

func Fn_gtk_entry_buffer_delete_text(param0 uint, param1 int) {}

func Fn_gtk_entry_buffer_emit_deleted_text(param0 uint, param1 uint) {}

func Fn_gtk_entry_buffer_emit_inserted_text(param0 uint, param1 string, param2 uint) {}

func Fn_gtk_entry_buffer_get_bytes() {}

func Fn_gtk_entry_buffer_get_length() {}

func Fn_gtk_entry_buffer_get_max_length() {}

func Fn_gtk_entry_buffer_get_text() {}

func Fn_gtk_entry_buffer_insert_text(param0 uint, param1 string, param2 int) {}

func Fn_gtk_entry_buffer_set_max_length(param0 int) {}

func Fn_gtk_entry_buffer_set_text(param0 string, param1 int) {}

func Fn_gtk_entry_completion_new() {
	C.gtk_entry_completion_new()
}

func Fn_gtk_entry_completion_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_entry_completion_complete() {}

func Fn_gtk_entry_completion_compute_prefix(param0 string) {}

func Fn_gtk_entry_completion_delete_action(param0 int) {}

func Fn_gtk_entry_completion_get_completion_prefix() {}

func Fn_gtk_entry_completion_get_entry() {}

func Fn_gtk_entry_completion_get_inline_completion() {}

func Fn_gtk_entry_completion_get_inline_selection() {}

func Fn_gtk_entry_completion_get_minimum_key_length() {}

func Fn_gtk_entry_completion_get_model() {}

func Fn_gtk_entry_completion_get_popup_completion() {}

func Fn_gtk_entry_completion_get_popup_set_width() {}

func Fn_gtk_entry_completion_get_popup_single_match() {}

func Fn_gtk_entry_completion_get_text_column() {}

func Fn_gtk_entry_completion_insert_action_markup(param0 int, param1 string) {}

func Fn_gtk_entry_completion_insert_action_text(param0 int, param1 string) {}

func Fn_gtk_entry_completion_insert_prefix() {}

func Fn_gtk_entry_completion_set_inline_completion(param0 bool) {}

func Fn_gtk_entry_completion_set_inline_selection(param0 bool) {}

// UNSUPPORTED : set_match_func : has callback

func Fn_gtk_entry_completion_set_minimum_key_length(param0 int) {}

func Fn_gtk_entry_completion_set_model(param0 unsafe.Pointer) {}

func Fn_gtk_entry_completion_set_popup_completion(param0 bool) {}

func Fn_gtk_entry_completion_set_popup_set_width(param0 bool) {}

func Fn_gtk_entry_completion_set_popup_single_match(param0 bool) {}

func Fn_gtk_entry_completion_set_text_column(param0 int) {}

func Fn_gtk_event_box_new() {
	C.gtk_event_box_new()
}

func Fn_gtk_event_box_get_above_child() {}

func Fn_gtk_event_box_get_visible_window() {}

func Fn_gtk_event_box_set_above_child(param0 bool) {}

func Fn_gtk_event_box_set_visible_window(param0 bool) {}

func Fn_gtk_event_controller_get_propagation_phase() {}

func Fn_gtk_event_controller_get_widget() {}

func Fn_gtk_event_controller_handle_event(param0 unsafe.Pointer) {}

func Fn_gtk_event_controller_reset() {}

func Fn_gtk_event_controller_set_propagation_phase(param0 int) {}

func Fn_gtk_event_controller_key_new(param0 unsafe.Pointer) {}

func Fn_gtk_event_controller_key_forward(param0 unsafe.Pointer) {}

func Fn_gtk_event_controller_key_get_group() {}

func Fn_gtk_event_controller_key_set_im_context(param0 unsafe.Pointer) {}

func Fn_gtk_expander_new(param0 string) {}

func Fn_gtk_expander_new_with_mnemonic(param0 string) {}

func Fn_gtk_expander_get_expanded() {}

func Fn_gtk_expander_get_label() {}

func Fn_gtk_expander_get_label_fill() {}

func Fn_gtk_expander_get_label_widget() {}

func Fn_gtk_expander_get_resize_toplevel() {}

func Fn_gtk_expander_get_spacing() {}

func Fn_gtk_expander_get_use_markup() {}

func Fn_gtk_expander_get_use_underline() {}

func Fn_gtk_expander_set_expanded(param0 bool) {}

func Fn_gtk_expander_set_label(param0 string) {}

func Fn_gtk_expander_set_label_fill(param0 bool) {}

func Fn_gtk_expander_set_label_widget(param0 unsafe.Pointer) {}

func Fn_gtk_expander_set_resize_toplevel(param0 bool) {}

func Fn_gtk_expander_set_spacing(param0 int) {}

func Fn_gtk_expander_set_use_markup(param0 bool) {}

func Fn_gtk_expander_set_use_underline(param0 bool) {}

func Fn_gtk_file_chooser_button_new(param0 string, param1 int) {}

func Fn_gtk_file_chooser_button_new_with_dialog(param0 unsafe.Pointer) {}

func Fn_gtk_file_chooser_button_get_focus_on_click() {}

func Fn_gtk_file_chooser_button_get_title() {}

func Fn_gtk_file_chooser_button_get_width_chars() {}

func Fn_gtk_file_chooser_button_set_focus_on_click(param0 bool) {}

func Fn_gtk_file_chooser_button_set_title(param0 string) {}

func Fn_gtk_file_chooser_button_set_width_chars(param0 int) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_file_chooser_widget_new(param0 int) {}

func Fn_gtk_file_filter_new() {
	C.gtk_file_filter_new()
}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_file_filter_add_mime_type(param0 string) {}

func Fn_gtk_file_filter_add_pattern(param0 string) {}

func Fn_gtk_file_filter_add_pixbuf_formats() {}

func Fn_gtk_file_filter_filter(param0 unsafe.Pointer) {}

func Fn_gtk_file_filter_get_name() {}

func Fn_gtk_file_filter_get_needed() {}

func Fn_gtk_file_filter_set_name(param0 string) {}

func Fn_gtk_fixed_new() {
	C.gtk_fixed_new()
}

func Fn_gtk_fixed_move(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_fixed_put(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_flow_box_new() {
	C.gtk_flow_box_new()
}

// UNSUPPORTED : bind_model : has callback

func Fn_gtk_flow_box_get_activate_on_single_click() {}

func Fn_gtk_flow_box_get_child_at_index(param0 int) {}

func Fn_gtk_flow_box_get_column_spacing() {}

func Fn_gtk_flow_box_get_homogeneous() {}

func Fn_gtk_flow_box_get_max_children_per_line() {}

func Fn_gtk_flow_box_get_min_children_per_line() {}

func Fn_gtk_flow_box_get_row_spacing() {}

func Fn_gtk_flow_box_get_selected_children() {}

func Fn_gtk_flow_box_get_selection_mode() {}

func Fn_gtk_flow_box_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_flow_box_invalidate_filter() {}

func Fn_gtk_flow_box_invalidate_sort() {}

func Fn_gtk_flow_box_select_all() {}

func Fn_gtk_flow_box_select_child(param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_flow_box_set_activate_on_single_click(param0 bool) {}

func Fn_gtk_flow_box_set_column_spacing(param0 uint) {}

// UNSUPPORTED : set_filter_func : has callback

func Fn_gtk_flow_box_set_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_set_homogeneous(param0 bool) {}

func Fn_gtk_flow_box_set_max_children_per_line(param0 uint) {}

func Fn_gtk_flow_box_set_min_children_per_line(param0 uint) {}

func Fn_gtk_flow_box_set_row_spacing(param0 uint) {}

func Fn_gtk_flow_box_set_selection_mode(param0 int) {}

// UNSUPPORTED : set_sort_func : has callback

func Fn_gtk_flow_box_set_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_unselect_all() {}

func Fn_gtk_flow_box_unselect_child(param0 unsafe.Pointer) {}

func Fn_gtk_flow_box_child_new() {
	C.gtk_flow_box_child_new()
}

func Fn_gtk_flow_box_child_changed() {}

func Fn_gtk_flow_box_child_get_index() {}

func Fn_gtk_flow_box_child_is_selected() {}

func Fn_gtk_font_button_new() {
	C.gtk_font_button_new()
}

func Fn_gtk_font_button_new_with_font(param0 string) {}

func Fn_gtk_font_button_get_font_name() {}

func Fn_gtk_font_button_get_show_size() {}

func Fn_gtk_font_button_get_show_style() {}

func Fn_gtk_font_button_get_title() {}

func Fn_gtk_font_button_get_use_font() {}

func Fn_gtk_font_button_get_use_size() {}

func Fn_gtk_font_button_set_font_name(param0 string) {}

func Fn_gtk_font_button_set_show_size(param0 bool) {}

func Fn_gtk_font_button_set_show_style(param0 bool) {}

func Fn_gtk_font_button_set_title(param0 string) {}

func Fn_gtk_font_button_set_use_font(param0 bool) {}

func Fn_gtk_font_button_set_use_size(param0 bool) {}

func Fn_gtk_font_chooser_dialog_new(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_font_chooser_widget_new() {
	C.gtk_font_chooser_widget_new()
}

func Fn_gtk_font_selection_new() {
	C.gtk_font_selection_new()
}

func Fn_gtk_font_selection_get_face() {}

func Fn_gtk_font_selection_get_face_list() {}

func Fn_gtk_font_selection_get_family() {}

func Fn_gtk_font_selection_get_family_list() {}

func Fn_gtk_font_selection_get_font_name() {}

func Fn_gtk_font_selection_get_preview_entry() {}

func Fn_gtk_font_selection_get_preview_text() {}

func Fn_gtk_font_selection_get_size() {}

func Fn_gtk_font_selection_get_size_entry() {}

func Fn_gtk_font_selection_get_size_list() {}

func Fn_gtk_font_selection_set_font_name(param0 string) {}

func Fn_gtk_font_selection_set_preview_text(param0 string) {}

func Fn_gtk_font_selection_dialog_new(param0 string) {}

func Fn_gtk_font_selection_dialog_get_cancel_button() {}

func Fn_gtk_font_selection_dialog_get_font_name() {}

func Fn_gtk_font_selection_dialog_get_font_selection() {}

func Fn_gtk_font_selection_dialog_get_ok_button() {}

func Fn_gtk_font_selection_dialog_get_preview_text() {}

func Fn_gtk_font_selection_dialog_set_font_name(param0 string) {}

func Fn_gtk_font_selection_dialog_set_preview_text(param0 string) {}

func Fn_gtk_frame_new(param0 string) {}

func Fn_gtk_frame_get_label() {}

func Fn_gtk_frame_get_label_align(param0 *float32, param1 *float32) {}

func Fn_gtk_frame_get_label_widget() {}

func Fn_gtk_frame_get_shadow_type() {}

func Fn_gtk_frame_set_label(param0 string) {}

func Fn_gtk_frame_set_label_align(param0 float32, param1 float32) {}

func Fn_gtk_frame_set_label_widget(param0 unsafe.Pointer) {}

func Fn_gtk_frame_set_shadow_type(param0 int) {}

func Fn_gtk_gesture_get_bounding_box(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_get_bounding_box_center(param0 *float64, param1 *float64) {}

func Fn_gtk_gesture_get_device() {}

func Fn_gtk_gesture_get_group() {}

func Fn_gtk_gesture_get_last_event(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_get_last_updated_sequence() {}

func Fn_gtk_gesture_get_point(param0 unsafe.Pointer, param1 *float64, param2 *float64) {}

func Fn_gtk_gesture_get_sequence_state(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_get_sequences() {}

func Fn_gtk_gesture_get_window() {}

func Fn_gtk_gesture_group(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_handles_sequence(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_is_active() {}

func Fn_gtk_gesture_is_grouped_with(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_is_recognized() {}

func Fn_gtk_gesture_set_sequence_state(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_gesture_set_state(param0 int) {}

func Fn_gtk_gesture_set_window(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_ungroup() {}

func Fn_gtk_gesture_drag_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_drag_get_offset(param0 *float64, param1 *float64) {}

func Fn_gtk_gesture_drag_get_start_point(param0 *float64, param1 *float64) {}

func Fn_gtk_gesture_long_press_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_multi_press_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_multi_press_get_area(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_multi_press_set_area(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_pan_new(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_gesture_pan_get_orientation() {}

func Fn_gtk_gesture_pan_set_orientation(param0 int) {}

func Fn_gtk_gesture_rotate_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_rotate_get_angle_delta() {}

func Fn_gtk_gesture_single_get_button() {}

func Fn_gtk_gesture_single_get_current_button() {}

func Fn_gtk_gesture_single_get_current_sequence() {}

func Fn_gtk_gesture_single_get_exclusive() {}

func Fn_gtk_gesture_single_get_touch_only() {}

func Fn_gtk_gesture_single_set_button(param0 uint) {}

func Fn_gtk_gesture_single_set_exclusive(param0 bool) {}

func Fn_gtk_gesture_single_set_touch_only(param0 bool) {}

func Fn_gtk_gesture_swipe_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_swipe_get_velocity(param0 *float64, param1 *float64) {}

func Fn_gtk_gesture_zoom_new(param0 unsafe.Pointer) {}

func Fn_gtk_gesture_zoom_get_scale_delta() {}

func Fn_gtk_grid_new() {
	C.gtk_grid_new()
}

func Fn_gtk_grid_attach(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 int) {}

func Fn_gtk_grid_attach_next_to(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int) {
}

func Fn_gtk_grid_get_baseline_row() {}

func Fn_gtk_grid_get_child_at(param0 int, param1 int) {}

func Fn_gtk_grid_get_column_homogeneous() {}

func Fn_gtk_grid_get_column_spacing() {}

func Fn_gtk_grid_get_row_baseline_position(param0 int) {}

func Fn_gtk_grid_get_row_homogeneous() {}

func Fn_gtk_grid_get_row_spacing() {}

func Fn_gtk_grid_insert_column(param0 int) {}

func Fn_gtk_grid_insert_next_to(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_grid_insert_row(param0 int) {}

func Fn_gtk_grid_remove_column(param0 int) {}

func Fn_gtk_grid_remove_row(param0 int) {}

func Fn_gtk_grid_set_baseline_row(param0 int) {}

func Fn_gtk_grid_set_column_homogeneous(param0 bool) {}

func Fn_gtk_grid_set_column_spacing(param0 uint) {}

func Fn_gtk_grid_set_row_baseline_position(param0 int, param1 int) {}

func Fn_gtk_grid_set_row_homogeneous(param0 bool) {}

func Fn_gtk_grid_set_row_spacing(param0 uint) {}

func Fn_gtk_hbox_new(param0 bool, param1 int) {}

func Fn_gtk_hbutton_box_new() {
	C.gtk_hbutton_box_new()
}

func Fn_gtk_hpaned_new() {
	C.gtk_hpaned_new()
}

func Fn_gtk_hsv_new() {
	C.gtk_hsv_new()
}

func Fn_gtk_hsv_get_color(param0 *float64, param1 *float64, param2 *float64) {}

func Fn_gtk_hsv_get_metrics(param0 *int, param1 *int) {}

func Fn_gtk_hsv_is_adjusting() {}

func Fn_gtk_hsv_set_color(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_hsv_set_metrics(param0 int, param1 int) {}

func Fn_gtk_hsv_to_rgb(param0 float64, param1 float64, param2 float64, param3 *float64, param4 *float64, param5 *float64) {
}

func Fn_gtk_hscale_new(param0 unsafe.Pointer) {}

func Fn_gtk_hscale_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_hscrollbar_new(param0 unsafe.Pointer) {}

func Fn_gtk_hseparator_new() {
	C.gtk_hseparator_new()
}

func Fn_gtk_handle_box_new() {
	C.gtk_handle_box_new()
}

func Fn_gtk_handle_box_get_child_detached() {}

func Fn_gtk_handle_box_get_handle_position() {}

func Fn_gtk_handle_box_get_shadow_type() {}

func Fn_gtk_handle_box_get_snap_edge() {}

func Fn_gtk_handle_box_set_handle_position(param0 int) {}

func Fn_gtk_handle_box_set_shadow_type(param0 int) {}

func Fn_gtk_handle_box_set_snap_edge(param0 int) {}

func Fn_gtk_header_bar_new() {
	C.gtk_header_bar_new()
}

func Fn_gtk_header_bar_get_custom_title() {}

func Fn_gtk_header_bar_get_decoration_layout() {}

func Fn_gtk_header_bar_get_has_subtitle() {}

func Fn_gtk_header_bar_get_show_close_button() {}

func Fn_gtk_header_bar_get_subtitle() {}

func Fn_gtk_header_bar_get_title() {}

func Fn_gtk_header_bar_pack_end(param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_pack_start(param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_set_custom_title(param0 unsafe.Pointer) {}

func Fn_gtk_header_bar_set_decoration_layout(param0 string) {}

func Fn_gtk_header_bar_set_has_subtitle(param0 bool) {}

func Fn_gtk_header_bar_set_show_close_button(param0 bool) {}

func Fn_gtk_header_bar_set_subtitle(param0 string) {}

func Fn_gtk_header_bar_set_title(param0 string) {}

func Fn_gtk_im_context_delete_surrounding(param0 int, param1 int) {}

func Fn_gtk_im_context_filter_keypress(param0 unsafe.Pointer) {}

func Fn_gtk_im_context_focus_in() {}

func Fn_gtk_im_context_focus_out() {}

func Fn_gtk_im_context_get_preedit_string(param0 string, param1 *unsafe.Pointer, param2 *int) {}

func Fn_gtk_im_context_get_surrounding(param0 string, param1 *int) {}

func Fn_gtk_im_context_reset() {}

func Fn_gtk_im_context_set_client_window(param0 unsafe.Pointer) {}

func Fn_gtk_im_context_set_cursor_location(param0 unsafe.Pointer) {}

func Fn_gtk_im_context_set_surrounding(param0 string, param1 int, param2 int) {}

func Fn_gtk_im_context_set_use_preedit(param0 bool) {}

func Fn_gtk_im_context_simple_new() {
	C.gtk_im_context_simple_new()
}

func Fn_gtk_im_context_simple_add_compose_file(param0 string) {}

func Fn_gtk_im_context_simple_add_table(param0 []uint16, param1 int, param2 int) {}

func Fn_gtk_im_multicontext_new() {
	C.gtk_im_multicontext_new()
}

func Fn_gtk_im_multicontext_append_menuitems(param0 unsafe.Pointer) {}

func Fn_gtk_im_multicontext_get_context_id() {}

func Fn_gtk_im_multicontext_set_context_id(param0 string) {}

func Fn_gtk_icon_factory_new() {
	C.gtk_icon_factory_new()
}

func Fn_gtk_icon_factory_add(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_icon_factory_add_default() {}

func Fn_gtk_icon_factory_lookup(param0 string) {}

func Fn_gtk_icon_factory_remove_default() {}

func Fn_gtk_icon_factory_lookup_default(param0 string) {}

func Fn_gtk_icon_info_new_for_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_icon_info_copy() {}

func Fn_gtk_icon_info_free() {}

func Fn_gtk_icon_info_get_attach_points(param0 []unsafe.Pointer, param1 *int) {}

func Fn_gtk_icon_info_get_base_scale() {}

func Fn_gtk_icon_info_get_base_size() {}

func Fn_gtk_icon_info_get_builtin_pixbuf() {}

func Fn_gtk_icon_info_get_display_name() {}

func Fn_gtk_icon_info_get_embedded_rect(param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_get_filename() {}

func Fn_gtk_icon_info_is_symbolic() {}

func Fn_gtk_icon_info_load_icon() {}

// UNSUPPORTED : load_icon_async : has callback

func Fn_gtk_icon_info_load_icon_finish(param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_load_surface(param0 unsafe.Pointer) {}

func Fn_gtk_icon_info_load_symbolic(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *bool) {
}

// UNSUPPORTED : load_symbolic_async : has callback

func Fn_gtk_icon_info_load_symbolic_finish(param0 unsafe.Pointer, param1 *bool) {}

func Fn_gtk_icon_info_load_symbolic_for_context(param0 unsafe.Pointer, param1 *bool) {}

// UNSUPPORTED : load_symbolic_for_context_async : has callback

func Fn_gtk_icon_info_load_symbolic_for_context_finish(param0 unsafe.Pointer, param1 *bool) {}

func Fn_gtk_icon_info_load_symbolic_for_style(param0 unsafe.Pointer, param1 int, param2 *bool) {}

func Fn_gtk_icon_info_set_raw_coordinates(param0 bool) {}

func Fn_gtk_icon_theme_new() {
	C.gtk_icon_theme_new()
}

func Fn_gtk_icon_theme_add_resource_path(param0 string) {}

func Fn_gtk_icon_theme_append_search_path(param0 string) {}

func Fn_gtk_icon_theme_choose_icon(param0 []string, param1 int, param2 int) {}

func Fn_gtk_icon_theme_choose_icon_for_scale(param0 []string, param1 int, param2 int, param3 int) {}

func Fn_gtk_icon_theme_get_example_icon_name() {}

func Fn_gtk_icon_theme_get_icon_sizes(param0 string) {}

func Fn_gtk_icon_theme_get_search_path(param0 *[]string, param1 *int) {}

func Fn_gtk_icon_theme_has_icon(param0 string) {}

func Fn_gtk_icon_theme_list_contexts() {}

func Fn_gtk_icon_theme_list_icons(param0 string) {}

func Fn_gtk_icon_theme_load_icon(param0 string, param1 int, param2 int) {}

func Fn_gtk_icon_theme_load_icon_for_scale(param0 string, param1 int, param2 int, param3 int) {}

func Fn_gtk_icon_theme_load_surface(param0 string, param1 int, param2 int, param3 unsafe.Pointer, param4 int) {
}

func Fn_gtk_icon_theme_lookup_by_gicon(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_icon_theme_lookup_by_gicon_for_scale(param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
}

func Fn_gtk_icon_theme_lookup_icon(param0 string, param1 int, param2 int) {}

func Fn_gtk_icon_theme_lookup_icon_for_scale(param0 string, param1 int, param2 int, param3 int) {}

func Fn_gtk_icon_theme_prepend_search_path(param0 string) {}

func Fn_gtk_icon_theme_rescan_if_needed() {}

func Fn_gtk_icon_theme_set_custom_theme(param0 string) {}

func Fn_gtk_icon_theme_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_icon_theme_set_search_path(param0 []string, param1 int) {}

func Fn_gtk_icon_theme_add_builtin_icon(param0 string, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_icon_theme_get_default() {
	C.gtk_icon_theme_get_default()
}

func Fn_gtk_icon_theme_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_new() {
	C.gtk_icon_view_new()
}

func Fn_gtk_icon_view_new_with_area(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_convert_widget_to_bin_window_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_icon_view_create_drag_icon(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_enable_model_drag_dest(param0 []TargetEntry, param1 int, param2 int) {}

func Fn_gtk_icon_view_enable_model_drag_source(param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_icon_view_get_activate_on_single_click() {}

func Fn_gtk_icon_view_get_cell_rect(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_column_spacing() {}

func Fn_gtk_icon_view_get_columns() {}

func Fn_gtk_icon_view_get_cursor(param0 *unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_gtk_icon_view_get_dest_item_at_pos(param0 int, param1 int, param2 *unsafe.Pointer, param3 int) {
}

func Fn_gtk_icon_view_get_drag_dest_item(param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_icon_view_get_item_at_pos(param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_item_column(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_orientation() {}

func Fn_gtk_icon_view_get_item_padding() {}

func Fn_gtk_icon_view_get_item_row(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_get_item_width() {}

func Fn_gtk_icon_view_get_margin() {}

func Fn_gtk_icon_view_get_markup_column() {}

func Fn_gtk_icon_view_get_model() {}

func Fn_gtk_icon_view_get_path_at_pos(param0 int, param1 int) {}

func Fn_gtk_icon_view_get_pixbuf_column() {}

func Fn_gtk_icon_view_get_reorderable() {}

func Fn_gtk_icon_view_get_row_spacing() {}

func Fn_gtk_icon_view_get_selected_items() {}

func Fn_gtk_icon_view_get_selection_mode() {}

func Fn_gtk_icon_view_get_spacing() {}

func Fn_gtk_icon_view_get_text_column() {}

func Fn_gtk_icon_view_get_tooltip_column() {}

func Fn_gtk_icon_view_get_tooltip_context(param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_gtk_icon_view_get_visible_range(param0 *unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_gtk_icon_view_item_activated(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_path_is_selected(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_scroll_to_path(param0 unsafe.Pointer, param1 bool, param2 float32, param3 float32) {
}

func Fn_gtk_icon_view_select_all() {}

func Fn_gtk_icon_view_select_path(param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_icon_view_set_activate_on_single_click(param0 bool) {}

func Fn_gtk_icon_view_set_column_spacing(param0 int) {}

func Fn_gtk_icon_view_set_columns(param0 int) {}

func Fn_gtk_icon_view_set_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_icon_view_set_drag_dest_item(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_icon_view_set_item_orientation(param0 int) {}

func Fn_gtk_icon_view_set_item_padding(param0 int) {}

func Fn_gtk_icon_view_set_item_width(param0 int) {}

func Fn_gtk_icon_view_set_margin(param0 int) {}

func Fn_gtk_icon_view_set_markup_column(param0 int) {}

func Fn_gtk_icon_view_set_model(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_set_pixbuf_column(param0 int) {}

func Fn_gtk_icon_view_set_reorderable(param0 bool) {}

func Fn_gtk_icon_view_set_row_spacing(param0 int) {}

func Fn_gtk_icon_view_set_selection_mode(param0 int) {}

func Fn_gtk_icon_view_set_spacing(param0 int) {}

func Fn_gtk_icon_view_set_text_column(param0 int) {}

func Fn_gtk_icon_view_set_tooltip_cell(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_icon_view_set_tooltip_column(param0 int) {}

func Fn_gtk_icon_view_set_tooltip_item(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_icon_view_unselect_all() {}

func Fn_gtk_icon_view_unselect_path(param0 unsafe.Pointer) {}

func Fn_gtk_icon_view_unset_model_drag_dest() {}

func Fn_gtk_icon_view_unset_model_drag_source() {}

func Fn_gtk_image_new() {
	C.gtk_image_new()
}

func Fn_gtk_image_new_from_animation(param0 unsafe.Pointer) {}

func Fn_gtk_image_new_from_file(param0 string) {}

func Fn_gtk_image_new_from_gicon(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_new_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_image_new_from_icon_set(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_new_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_image_new_from_resource(param0 string) {}

func Fn_gtk_image_new_from_stock(param0 string, param1 int) {}

func Fn_gtk_image_new_from_surface(param0 unsafe.Pointer) {}

func Fn_gtk_image_clear() {}

func Fn_gtk_image_get_animation() {}

func Fn_gtk_image_get_gicon(param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_image_get_icon_name(param0 string, param1 int) {}

func Fn_gtk_image_get_icon_set(param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_image_get_pixbuf() {}

func Fn_gtk_image_get_pixel_size() {}

func Fn_gtk_image_get_stock(param0 string, param1 int) {}

func Fn_gtk_image_get_storage_type() {}

func Fn_gtk_image_set_from_animation(param0 unsafe.Pointer) {}

func Fn_gtk_image_set_from_file(param0 string) {}

func Fn_gtk_image_set_from_gicon(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_set_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_image_set_from_icon_set(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_image_set_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_image_set_from_resource(param0 string) {}

func Fn_gtk_image_set_from_stock(param0 string, param1 int) {}

func Fn_gtk_image_set_from_surface(param0 unsafe.Pointer) {}

func Fn_gtk_image_set_pixel_size(param0 int) {}

func Fn_gtk_image_menu_item_new() {
	C.gtk_image_menu_item_new()
}

func Fn_gtk_image_menu_item_new_from_stock(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_new_with_label(param0 string) {}

func Fn_gtk_image_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_image_menu_item_get_always_show_image() {}

func Fn_gtk_image_menu_item_get_image() {}

func Fn_gtk_image_menu_item_get_use_stock() {}

func Fn_gtk_image_menu_item_set_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_set_always_show_image(param0 bool) {}

func Fn_gtk_image_menu_item_set_image(param0 unsafe.Pointer) {}

func Fn_gtk_image_menu_item_set_use_stock(param0 bool) {}

func Fn_gtk_info_bar_new() {
	C.gtk_info_bar_new()
}

// UNSUPPORTED : new_with_buttons : has varargs

func Fn_gtk_info_bar_add_action_widget(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_info_bar_add_button(param0 string, param1 int) {}

// UNSUPPORTED : add_buttons : has varargs

func Fn_gtk_info_bar_get_action_area() {}

func Fn_gtk_info_bar_get_content_area() {}

func Fn_gtk_info_bar_get_message_type() {}

func Fn_gtk_info_bar_get_show_close_button() {}

func Fn_gtk_info_bar_response(param0 int) {}

func Fn_gtk_info_bar_set_default_response(param0 int) {}

func Fn_gtk_info_bar_set_message_type(param0 int) {}

func Fn_gtk_info_bar_set_response_sensitive(param0 int, param1 bool) {}

func Fn_gtk_info_bar_set_show_close_button(param0 bool) {}

func Fn_gtk_invisible_new() {
	C.gtk_invisible_new()
}

func Fn_gtk_invisible_new_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_invisible_get_screen() {}

func Fn_gtk_invisible_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_label_new(param0 string) {}

func Fn_gtk_label_new_with_mnemonic(param0 string) {}

func Fn_gtk_label_get_angle() {}

func Fn_gtk_label_get_attributes() {}

func Fn_gtk_label_get_current_uri() {}

func Fn_gtk_label_get_ellipsize() {}

func Fn_gtk_label_get_justify() {}

func Fn_gtk_label_get_label() {}

func Fn_gtk_label_get_layout() {}

func Fn_gtk_label_get_layout_offsets(param0 *int, param1 *int) {}

func Fn_gtk_label_get_line_wrap() {}

func Fn_gtk_label_get_line_wrap_mode() {}

func Fn_gtk_label_get_lines() {}

func Fn_gtk_label_get_max_width_chars() {}

func Fn_gtk_label_get_mnemonic_keyval() {}

func Fn_gtk_label_get_mnemonic_widget() {}

func Fn_gtk_label_get_selectable() {}

func Fn_gtk_label_get_selection_bounds(param0 *int, param1 *int) {}

func Fn_gtk_label_get_single_line_mode() {}

func Fn_gtk_label_get_text() {}

func Fn_gtk_label_get_track_visited_links() {}

func Fn_gtk_label_get_use_markup() {}

func Fn_gtk_label_get_use_underline() {}

func Fn_gtk_label_get_width_chars() {}

func Fn_gtk_label_select_region(param0 int, param1 int) {}

func Fn_gtk_label_set_angle(param0 float64) {}

func Fn_gtk_label_set_attributes(param0 unsafe.Pointer) {}

func Fn_gtk_label_set_ellipsize(param0 int) {}

func Fn_gtk_label_set_justify(param0 int) {}

func Fn_gtk_label_set_label(param0 string) {}

func Fn_gtk_label_set_line_wrap(param0 bool) {}

func Fn_gtk_label_set_line_wrap_mode(param0 int) {}

func Fn_gtk_label_set_lines(param0 int) {}

func Fn_gtk_label_set_markup(param0 string) {}

func Fn_gtk_label_set_markup_with_mnemonic(param0 string) {}

func Fn_gtk_label_set_max_width_chars(param0 int) {}

func Fn_gtk_label_set_mnemonic_widget(param0 unsafe.Pointer) {}

func Fn_gtk_label_set_pattern(param0 string) {}

func Fn_gtk_label_set_selectable(param0 bool) {}

func Fn_gtk_label_set_single_line_mode(param0 bool) {}

func Fn_gtk_label_set_text(param0 string) {}

func Fn_gtk_label_set_text_with_mnemonic(param0 string) {}

func Fn_gtk_label_set_track_visited_links(param0 bool) {}

func Fn_gtk_label_set_use_markup(param0 bool) {}

func Fn_gtk_label_set_use_underline(param0 bool) {}

func Fn_gtk_label_set_width_chars(param0 int) {}

func Fn_gtk_layout_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_layout_get_bin_window() {}

func Fn_gtk_layout_get_hadjustment() {}

func Fn_gtk_layout_get_size(param0 *uint, param1 *uint) {}

func Fn_gtk_layout_get_vadjustment() {}

func Fn_gtk_layout_move(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_layout_put(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_layout_set_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_layout_set_size(param0 uint, param1 uint) {}

func Fn_gtk_layout_set_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_level_bar_new() {
	C.gtk_level_bar_new()
}

func Fn_gtk_level_bar_new_for_interval(param0 float64, param1 float64) {}

func Fn_gtk_level_bar_add_offset_value(param0 string, param1 float64) {}

func Fn_gtk_level_bar_get_inverted() {}

func Fn_gtk_level_bar_get_max_value() {}

func Fn_gtk_level_bar_get_min_value() {}

func Fn_gtk_level_bar_get_mode() {}

func Fn_gtk_level_bar_get_offset_value(param0 string, param1 *float64) {}

func Fn_gtk_level_bar_get_value() {}

func Fn_gtk_level_bar_remove_offset_value(param0 string) {}

func Fn_gtk_level_bar_set_inverted(param0 bool) {}

func Fn_gtk_level_bar_set_max_value(param0 float64) {}

func Fn_gtk_level_bar_set_min_value(param0 float64) {}

func Fn_gtk_level_bar_set_mode(param0 int) {}

func Fn_gtk_level_bar_set_value(param0 float64) {}

func Fn_gtk_link_button_new(param0 string) {}

func Fn_gtk_link_button_new_with_label(param0 string, param1 string) {}

func Fn_gtk_link_button_get_uri() {}

func Fn_gtk_link_button_get_visited() {}

func Fn_gtk_link_button_set_uri(param0 string) {}

func Fn_gtk_link_button_set_visited(param0 bool) {}

func Fn_gtk_list_box_new() {
	C.gtk_list_box_new()
}

// UNSUPPORTED : bind_model : has callback

func Fn_gtk_list_box_drag_highlight_row(param0 unsafe.Pointer) {}

func Fn_gtk_list_box_drag_unhighlight_row() {}

func Fn_gtk_list_box_get_activate_on_single_click() {}

func Fn_gtk_list_box_get_adjustment() {}

func Fn_gtk_list_box_get_row_at_index(param0 int) {}

func Fn_gtk_list_box_get_row_at_y(param0 int) {}

func Fn_gtk_list_box_get_selected_row() {}

func Fn_gtk_list_box_get_selected_rows() {}

func Fn_gtk_list_box_get_selection_mode() {}

func Fn_gtk_list_box_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_list_box_invalidate_filter() {}

func Fn_gtk_list_box_invalidate_headers() {}

func Fn_gtk_list_box_invalidate_sort() {}

func Fn_gtk_list_box_prepend(param0 unsafe.Pointer) {}

func Fn_gtk_list_box_select_all() {}

func Fn_gtk_list_box_select_row(param0 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_list_box_set_activate_on_single_click(param0 bool) {}

func Fn_gtk_list_box_set_adjustment(param0 unsafe.Pointer) {}

// UNSUPPORTED : set_filter_func : has callback

// UNSUPPORTED : set_header_func : has callback

func Fn_gtk_list_box_set_placeholder(param0 unsafe.Pointer) {}

func Fn_gtk_list_box_set_selection_mode(param0 int) {}

// UNSUPPORTED : set_sort_func : has callback

func Fn_gtk_list_box_unselect_all() {}

func Fn_gtk_list_box_unselect_row(param0 unsafe.Pointer) {}

func Fn_gtk_list_box_row_new() {
	C.gtk_list_box_row_new()
}

func Fn_gtk_list_box_row_changed() {}

func Fn_gtk_list_box_row_get_activatable() {}

func Fn_gtk_list_box_row_get_header() {}

func Fn_gtk_list_box_row_get_index() {}

func Fn_gtk_list_box_row_get_selectable() {}

func Fn_gtk_list_box_row_is_selected() {}

func Fn_gtk_list_box_row_set_activatable(param0 bool) {}

func Fn_gtk_list_box_row_set_header(param0 unsafe.Pointer) {}

func Fn_gtk_list_box_row_set_selectable(param0 bool) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_list_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_list_store_append(param0 unsafe.Pointer) {}

func Fn_gtk_list_store_clear() {}

func Fn_gtk_list_store_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_list_store_insert_after(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_list_store_insert_before(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_list_store_insert_with_valuesv(param0 unsafe.Pointer, param1 int, param2 []int, param3 []gobject.Value, param4 int) {
}

func Fn_gtk_list_store_iter_is_valid(param0 unsafe.Pointer) {}

func Fn_gtk_list_store_move_after(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_list_store_move_before(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_list_store_prepend(param0 unsafe.Pointer) {}

func Fn_gtk_list_store_remove(param0 unsafe.Pointer) {}

func Fn_gtk_list_store_reorder(param0 []int) {}

// UNSUPPORTED : set : has varargs

func Fn_gtk_list_store_set_column_types(param0 int, param1 []uint64) {}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_list_store_set_value(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_list_store_set_valuesv(param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
}

func Fn_gtk_list_store_swap(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_lock_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_lock_button_get_permission() {}

func Fn_gtk_lock_button_set_permission(param0 unsafe.Pointer) {}

func Fn_gtk_menu_new() {
	C.gtk_menu_new()
}

func Fn_gtk_menu_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_attach(param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {}

// UNSUPPORTED : attach_to_widget : has callback

func Fn_gtk_menu_detach() {}

func Fn_gtk_menu_get_accel_group() {}

func Fn_gtk_menu_get_accel_path() {}

func Fn_gtk_menu_get_active() {}

func Fn_gtk_menu_get_attach_widget() {}

func Fn_gtk_menu_get_monitor() {}

func Fn_gtk_menu_get_reserve_toggle_size() {}

func Fn_gtk_menu_get_tearoff_state() {}

func Fn_gtk_menu_get_title() {}

func Fn_gtk_menu_popdown() {}

// UNSUPPORTED : popup : has callback

// UNSUPPORTED : popup_for_device : has callback

func Fn_gtk_menu_reorder_child(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_menu_reposition() {}

func Fn_gtk_menu_set_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_menu_set_accel_path(param0 string) {}

func Fn_gtk_menu_set_active(param0 uint) {}

func Fn_gtk_menu_set_monitor(param0 int) {}

func Fn_gtk_menu_set_reserve_toggle_size(param0 bool) {}

func Fn_gtk_menu_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_menu_set_tearoff_state(param0 bool) {}

func Fn_gtk_menu_set_title(param0 string) {}

func Fn_gtk_menu_get_for_attach_widget(param0 unsafe.Pointer) {}

func Fn_gtk_menu_bar_new() {
	C.gtk_menu_bar_new()
}

func Fn_gtk_menu_bar_new_from_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_bar_get_child_pack_direction() {}

func Fn_gtk_menu_bar_get_pack_direction() {}

func Fn_gtk_menu_bar_set_child_pack_direction(param0 int) {}

func Fn_gtk_menu_bar_set_pack_direction(param0 int) {}

func Fn_gtk_menu_button_new() {
	C.gtk_menu_button_new()
}

func Fn_gtk_menu_button_get_align_widget() {}

func Fn_gtk_menu_button_get_direction() {}

func Fn_gtk_menu_button_get_menu_model() {}

func Fn_gtk_menu_button_get_popover() {}

func Fn_gtk_menu_button_get_popup() {}

func Fn_gtk_menu_button_get_use_popover() {}

func Fn_gtk_menu_button_set_align_widget(param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_direction(param0 int) {}

func Fn_gtk_menu_button_set_menu_model(param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_popover(param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_popup(param0 unsafe.Pointer) {}

func Fn_gtk_menu_button_set_use_popover(param0 bool) {}

func Fn_gtk_menu_item_new() {
	C.gtk_menu_item_new()
}

func Fn_gtk_menu_item_new_with_label(param0 string) {}

func Fn_gtk_menu_item_new_with_mnemonic(param0 string) {}

func Fn_gtk_menu_item_activate() {}

func Fn_gtk_menu_item_deselect() {}

func Fn_gtk_menu_item_get_accel_path() {}

func Fn_gtk_menu_item_get_label() {}

func Fn_gtk_menu_item_get_reserve_indicator() {}

func Fn_gtk_menu_item_get_right_justified() {}

func Fn_gtk_menu_item_get_submenu() {}

func Fn_gtk_menu_item_get_use_underline() {}

func Fn_gtk_menu_item_select() {}

func Fn_gtk_menu_item_set_accel_path(param0 string) {}

func Fn_gtk_menu_item_set_label(param0 string) {}

func Fn_gtk_menu_item_set_reserve_indicator(param0 bool) {}

func Fn_gtk_menu_item_set_right_justified(param0 bool) {}

func Fn_gtk_menu_item_set_submenu(param0 unsafe.Pointer) {}

func Fn_gtk_menu_item_set_use_underline(param0 bool) {}

func Fn_gtk_menu_item_toggle_size_allocate(param0 int) {}

func Fn_gtk_menu_item_toggle_size_request(param0 *int) {}

func Fn_gtk_menu_shell_activate_item(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_menu_shell_append(param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_bind_model(param0 unsafe.Pointer, param1 string, param2 bool) {}

func Fn_gtk_menu_shell_cancel() {}

func Fn_gtk_menu_shell_deactivate() {}

func Fn_gtk_menu_shell_deselect() {}

func Fn_gtk_menu_shell_get_parent_shell() {}

func Fn_gtk_menu_shell_get_selected_item() {}

func Fn_gtk_menu_shell_get_take_focus() {}

func Fn_gtk_menu_shell_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_menu_shell_prepend(param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_select_first(param0 bool) {}

func Fn_gtk_menu_shell_select_item(param0 unsafe.Pointer) {}

func Fn_gtk_menu_shell_set_take_focus(param0 bool) {}

func Fn_gtk_menu_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_menu_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_menu_tool_button_get_menu() {}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_markup(param0 string) {}

func Fn_gtk_menu_tool_button_set_arrow_tooltip_text(param0 string) {}

func Fn_gtk_menu_tool_button_set_menu(param0 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_with_markup : has varargs

// UNSUPPORTED : format_secondary_markup : has varargs

// UNSUPPORTED : format_secondary_text : has varargs

func Fn_gtk_message_dialog_get_image() {}

func Fn_gtk_message_dialog_get_message_area() {}

func Fn_gtk_message_dialog_set_image(param0 unsafe.Pointer) {}

func Fn_gtk_message_dialog_set_markup(param0 string) {}

func Fn_gtk_misc_get_alignment(param0 *float32, param1 *float32) {}

func Fn_gtk_misc_get_padding(param0 *int, param1 *int) {}

func Fn_gtk_misc_set_alignment(param0 float32, param1 float32) {}

func Fn_gtk_misc_set_padding(param0 int, param1 int) {}

func Fn_gtk_mount_operation_new(param0 unsafe.Pointer) {}

func Fn_gtk_mount_operation_get_parent() {}

func Fn_gtk_mount_operation_get_screen() {}

func Fn_gtk_mount_operation_is_showing() {}

func Fn_gtk_mount_operation_set_parent(param0 unsafe.Pointer) {}

func Fn_gtk_mount_operation_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_new() {
	C.gtk_notebook_new()
}

func Fn_gtk_notebook_append_page(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_append_page_menu(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_notebook_get_action_widget(param0 int) {}

func Fn_gtk_notebook_get_current_page() {}

func Fn_gtk_notebook_get_group_name() {}

func Fn_gtk_notebook_get_menu_label(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_menu_label_text(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_n_pages() {}

func Fn_gtk_notebook_get_nth_page(param0 int) {}

func Fn_gtk_notebook_get_scrollable() {}

func Fn_gtk_notebook_get_show_border() {}

func Fn_gtk_notebook_get_show_tabs() {}

func Fn_gtk_notebook_get_tab_detachable(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_hborder() {}

func Fn_gtk_notebook_get_tab_label(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_label_text(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_pos() {}

func Fn_gtk_notebook_get_tab_reorderable(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_get_tab_vborder() {}

func Fn_gtk_notebook_insert_page(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_notebook_insert_page_menu(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 int) {
}

func Fn_gtk_notebook_next_page() {}

func Fn_gtk_notebook_page_num(param0 unsafe.Pointer) {}

func Fn_gtk_notebook_popup_disable() {}

func Fn_gtk_notebook_popup_enable() {}

func Fn_gtk_notebook_prepend_page(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_prepend_page_menu(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_notebook_prev_page() {}

func Fn_gtk_notebook_remove_page(param0 int) {}

func Fn_gtk_notebook_reorder_child(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_notebook_set_action_widget(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_notebook_set_current_page(param0 int) {}

func Fn_gtk_notebook_set_group_name(param0 string) {}

func Fn_gtk_notebook_set_menu_label(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_set_menu_label_text(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_notebook_set_scrollable(param0 bool) {}

func Fn_gtk_notebook_set_show_border(param0 bool) {}

func Fn_gtk_notebook_set_show_tabs(param0 bool) {}

func Fn_gtk_notebook_set_tab_detachable(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_notebook_set_tab_label(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_set_tab_label_text(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_notebook_set_tab_pos(param0 int) {}

func Fn_gtk_notebook_set_tab_reorderable(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_notebook_page_accessible_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_notebook_page_accessible_invalidate() {}

func Fn_gtk_numerable_icon_get_background_gicon() {}

func Fn_gtk_numerable_icon_get_background_icon_name() {}

func Fn_gtk_numerable_icon_get_count() {}

func Fn_gtk_numerable_icon_get_label() {}

func Fn_gtk_numerable_icon_get_style_context() {}

func Fn_gtk_numerable_icon_set_background_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_set_background_icon_name(param0 string) {}

func Fn_gtk_numerable_icon_set_count(param0 int) {}

func Fn_gtk_numerable_icon_set_label(param0 string) {}

func Fn_gtk_numerable_icon_set_style_context(param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new(param0 unsafe.Pointer) {}

func Fn_gtk_numerable_icon_new_with_style_context(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_offscreen_window_new() {
	C.gtk_offscreen_window_new()
}

func Fn_gtk_offscreen_window_get_pixbuf() {}

func Fn_gtk_offscreen_window_get_surface() {}

func Fn_gtk_overlay_new() {
	C.gtk_overlay_new()
}

func Fn_gtk_overlay_add_overlay(param0 unsafe.Pointer) {}

func Fn_gtk_page_setup_new() {
	C.gtk_page_setup_new()
}

func Fn_gtk_page_setup_new_from_file(param0 string) {}

func Fn_gtk_page_setup_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_page_setup_copy() {}

func Fn_gtk_page_setup_get_bottom_margin(param0 int) {}

func Fn_gtk_page_setup_get_left_margin(param0 int) {}

func Fn_gtk_page_setup_get_orientation() {}

func Fn_gtk_page_setup_get_page_height(param0 int) {}

func Fn_gtk_page_setup_get_page_width(param0 int) {}

func Fn_gtk_page_setup_get_paper_height(param0 int) {}

func Fn_gtk_page_setup_get_paper_size() {}

func Fn_gtk_page_setup_get_paper_width(param0 int) {}

func Fn_gtk_page_setup_get_right_margin(param0 int) {}

func Fn_gtk_page_setup_get_top_margin(param0 int) {}

func Fn_gtk_page_setup_load_file(param0 string) {}

func Fn_gtk_page_setup_load_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_page_setup_set_bottom_margin(param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_left_margin(param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_orientation(param0 int) {}

func Fn_gtk_page_setup_set_paper_size(param0 unsafe.Pointer) {}

func Fn_gtk_page_setup_set_paper_size_and_default_margins(param0 unsafe.Pointer) {}

func Fn_gtk_page_setup_set_right_margin(param0 float64, param1 int) {}

func Fn_gtk_page_setup_set_top_margin(param0 float64, param1 int) {}

func Fn_gtk_page_setup_to_file(param0 string) {}

func Fn_gtk_page_setup_to_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_paned_new(param0 int) {}

func Fn_gtk_paned_add1(param0 unsafe.Pointer) {}

func Fn_gtk_paned_add2(param0 unsafe.Pointer) {}

func Fn_gtk_paned_get_child1() {}

func Fn_gtk_paned_get_child2() {}

func Fn_gtk_paned_get_handle_window() {}

func Fn_gtk_paned_get_position() {}

func Fn_gtk_paned_pack1(param0 unsafe.Pointer, param1 bool, param2 bool) {}

func Fn_gtk_paned_pack2(param0 unsafe.Pointer, param1 bool, param2 bool) {}

func Fn_gtk_paned_set_position(param0 int) {}

func Fn_gtk_places_sidebar_new() {
	C.gtk_places_sidebar_new()
}

func Fn_gtk_places_sidebar_add_shortcut(param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_get_local_only() {}

func Fn_gtk_places_sidebar_get_location() {}

func Fn_gtk_places_sidebar_get_nth_bookmark(param0 int) {}

func Fn_gtk_places_sidebar_get_open_flags() {}

func Fn_gtk_places_sidebar_get_show_connect_to_server() {}

func Fn_gtk_places_sidebar_get_show_desktop() {}

func Fn_gtk_places_sidebar_get_show_enter_location() {}

func Fn_gtk_places_sidebar_list_shortcuts() {}

func Fn_gtk_places_sidebar_remove_shortcut(param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_set_local_only(param0 bool) {}

func Fn_gtk_places_sidebar_set_location(param0 unsafe.Pointer) {}

func Fn_gtk_places_sidebar_set_open_flags(param0 int) {}

func Fn_gtk_places_sidebar_set_show_connect_to_server(param0 bool) {}

func Fn_gtk_places_sidebar_set_show_desktop(param0 bool) {}

func Fn_gtk_places_sidebar_set_show_enter_location(param0 bool) {}

func Fn_gtk_plug_new(param0 uint64) {}

func Fn_gtk_plug_new_for_display(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_plug_construct(param0 uint64) {}

func Fn_gtk_plug_construct_for_display(param0 unsafe.Pointer, param1 uint64) {}

func Fn_gtk_plug_get_embedded() {}

func Fn_gtk_plug_get_id() {}

func Fn_gtk_plug_get_socket_window() {}

func Fn_gtk_popover_new(param0 unsafe.Pointer) {}

func Fn_gtk_popover_new_from_model(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_popover_bind_model(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_popover_get_modal() {}

func Fn_gtk_popover_get_pointing_to(param0 unsafe.Pointer) {}

func Fn_gtk_popover_get_position() {}

func Fn_gtk_popover_get_relative_to() {}

func Fn_gtk_popover_set_modal(param0 bool) {}

func Fn_gtk_popover_set_pointing_to(param0 unsafe.Pointer) {}

func Fn_gtk_popover_set_position(param0 int) {}

func Fn_gtk_popover_set_relative_to(param0 unsafe.Pointer) {}

func Fn_gtk_print_context_create_pango_context() {}

func Fn_gtk_print_context_create_pango_layout() {}

func Fn_gtk_print_context_get_cairo_context() {}

func Fn_gtk_print_context_get_dpi_x() {}

func Fn_gtk_print_context_get_dpi_y() {}

func Fn_gtk_print_context_get_hard_margins(param0 *float64, param1 *float64, param2 *float64, param3 *float64) {
}

func Fn_gtk_print_context_get_height() {}

func Fn_gtk_print_context_get_page_setup() {}

func Fn_gtk_print_context_get_pango_fontmap() {}

func Fn_gtk_print_context_get_width() {}

func Fn_gtk_print_context_set_cairo_context(param0 unsafe.Pointer, param1 float64, param2 float64) {}

func Fn_gtk_print_operation_new() {
	C.gtk_print_operation_new()
}

func Fn_gtk_print_operation_cancel() {}

func Fn_gtk_print_operation_draw_page_finish() {}

func Fn_gtk_print_operation_get_default_page_setup() {}

func Fn_gtk_print_operation_get_embed_page_setup() {}

func Fn_gtk_print_operation_get_error() {}

func Fn_gtk_print_operation_get_has_selection() {}

func Fn_gtk_print_operation_get_n_pages_to_print() {}

func Fn_gtk_print_operation_get_print_settings() {}

func Fn_gtk_print_operation_get_status() {}

func Fn_gtk_print_operation_get_status_string() {}

func Fn_gtk_print_operation_get_support_selection() {}

func Fn_gtk_print_operation_is_finished() {}

func Fn_gtk_print_operation_run(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_print_operation_set_allow_async(param0 bool) {}

func Fn_gtk_print_operation_set_current_page(param0 int) {}

func Fn_gtk_print_operation_set_custom_tab_label(param0 string) {}

func Fn_gtk_print_operation_set_default_page_setup(param0 unsafe.Pointer) {}

func Fn_gtk_print_operation_set_defer_drawing() {}

func Fn_gtk_print_operation_set_embed_page_setup(param0 bool) {}

func Fn_gtk_print_operation_set_export_filename(param0 string) {}

func Fn_gtk_print_operation_set_has_selection(param0 bool) {}

func Fn_gtk_print_operation_set_job_name(param0 string) {}

func Fn_gtk_print_operation_set_n_pages(param0 int) {}

func Fn_gtk_print_operation_set_print_settings(param0 unsafe.Pointer) {}

func Fn_gtk_print_operation_set_show_progress(param0 bool) {}

func Fn_gtk_print_operation_set_support_selection(param0 bool) {}

func Fn_gtk_print_operation_set_track_print_status(param0 bool) {}

func Fn_gtk_print_operation_set_unit(param0 int) {}

func Fn_gtk_print_operation_set_use_full_page(param0 bool) {}

func Fn_gtk_print_settings_new() {
	C.gtk_print_settings_new()
}

func Fn_gtk_print_settings_new_from_file(param0 string) {}

func Fn_gtk_print_settings_new_from_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_print_settings_copy() {}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_print_settings_get(param0 string) {}

func Fn_gtk_print_settings_get_bool(param0 string) {}

func Fn_gtk_print_settings_get_collate() {}

func Fn_gtk_print_settings_get_default_source() {}

func Fn_gtk_print_settings_get_dither() {}

func Fn_gtk_print_settings_get_double(param0 string) {}

func Fn_gtk_print_settings_get_double_with_default(param0 string, param1 float64) {}

func Fn_gtk_print_settings_get_duplex() {}

func Fn_gtk_print_settings_get_finishings() {}

func Fn_gtk_print_settings_get_int(param0 string) {}

func Fn_gtk_print_settings_get_int_with_default(param0 string, param1 int) {}

func Fn_gtk_print_settings_get_length(param0 string, param1 int) {}

func Fn_gtk_print_settings_get_media_type() {}

func Fn_gtk_print_settings_get_n_copies() {}

func Fn_gtk_print_settings_get_number_up() {}

func Fn_gtk_print_settings_get_number_up_layout() {}

func Fn_gtk_print_settings_get_orientation() {}

func Fn_gtk_print_settings_get_output_bin() {}

func Fn_gtk_print_settings_get_page_ranges(param0 *int) {}

func Fn_gtk_print_settings_get_page_set() {}

func Fn_gtk_print_settings_get_paper_height(param0 int) {}

func Fn_gtk_print_settings_get_paper_size() {}

func Fn_gtk_print_settings_get_paper_width(param0 int) {}

func Fn_gtk_print_settings_get_print_pages() {}

func Fn_gtk_print_settings_get_printer() {}

func Fn_gtk_print_settings_get_printer_lpi() {}

func Fn_gtk_print_settings_get_quality() {}

func Fn_gtk_print_settings_get_resolution() {}

func Fn_gtk_print_settings_get_resolution_x() {}

func Fn_gtk_print_settings_get_resolution_y() {}

func Fn_gtk_print_settings_get_reverse() {}

func Fn_gtk_print_settings_get_scale() {}

func Fn_gtk_print_settings_get_use_color() {}

func Fn_gtk_print_settings_has_key(param0 string) {}

func Fn_gtk_print_settings_load_file(param0 string) {}

func Fn_gtk_print_settings_load_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_print_settings_set(param0 string, param1 string) {}

func Fn_gtk_print_settings_set_bool(param0 string, param1 bool) {}

func Fn_gtk_print_settings_set_collate(param0 bool) {}

func Fn_gtk_print_settings_set_default_source(param0 string) {}

func Fn_gtk_print_settings_set_dither(param0 string) {}

func Fn_gtk_print_settings_set_double(param0 string, param1 float64) {}

func Fn_gtk_print_settings_set_duplex(param0 int) {}

func Fn_gtk_print_settings_set_finishings(param0 string) {}

func Fn_gtk_print_settings_set_int(param0 string, param1 int) {}

func Fn_gtk_print_settings_set_length(param0 string, param1 float64, param2 int) {}

func Fn_gtk_print_settings_set_media_type(param0 string) {}

func Fn_gtk_print_settings_set_n_copies(param0 int) {}

func Fn_gtk_print_settings_set_number_up(param0 int) {}

func Fn_gtk_print_settings_set_number_up_layout(param0 int) {}

func Fn_gtk_print_settings_set_orientation(param0 int) {}

func Fn_gtk_print_settings_set_output_bin(param0 string) {}

func Fn_gtk_print_settings_set_page_ranges(param0 []PageRange, param1 int) {}

func Fn_gtk_print_settings_set_page_set(param0 int) {}

func Fn_gtk_print_settings_set_paper_height(param0 float64, param1 int) {}

func Fn_gtk_print_settings_set_paper_size(param0 unsafe.Pointer) {}

func Fn_gtk_print_settings_set_paper_width(param0 float64, param1 int) {}

func Fn_gtk_print_settings_set_print_pages(param0 int) {}

func Fn_gtk_print_settings_set_printer(param0 string) {}

func Fn_gtk_print_settings_set_printer_lpi(param0 float64) {}

func Fn_gtk_print_settings_set_quality(param0 int) {}

func Fn_gtk_print_settings_set_resolution(param0 int) {}

func Fn_gtk_print_settings_set_resolution_xy(param0 int, param1 int) {}

func Fn_gtk_print_settings_set_reverse(param0 bool) {}

func Fn_gtk_print_settings_set_scale(param0 float64) {}

func Fn_gtk_print_settings_set_use_color(param0 bool) {}

func Fn_gtk_print_settings_to_file(param0 string) {}

func Fn_gtk_print_settings_to_key_file(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_print_settings_unset(param0 string) {}

func Fn_gtk_progress_bar_new() {
	C.gtk_progress_bar_new()
}

func Fn_gtk_progress_bar_get_ellipsize() {}

func Fn_gtk_progress_bar_get_fraction() {}

func Fn_gtk_progress_bar_get_inverted() {}

func Fn_gtk_progress_bar_get_pulse_step() {}

func Fn_gtk_progress_bar_get_show_text() {}

func Fn_gtk_progress_bar_get_text() {}

func Fn_gtk_progress_bar_pulse() {}

func Fn_gtk_progress_bar_set_ellipsize(param0 int) {}

func Fn_gtk_progress_bar_set_fraction(param0 float64) {}

func Fn_gtk_progress_bar_set_inverted(param0 bool) {}

func Fn_gtk_progress_bar_set_pulse_step(param0 float64) {}

func Fn_gtk_progress_bar_set_show_text(param0 bool) {}

func Fn_gtk_progress_bar_set_text(param0 string) {}

func Fn_gtk_radio_action_new(param0 string, param1 string, param2 string, param3 string, param4 int) {}

func Fn_gtk_radio_action_get_current_value() {}

func Fn_gtk_radio_action_get_group() {}

func Fn_gtk_radio_action_join_group(param0 unsafe.Pointer) {}

func Fn_gtk_radio_action_set_current_value(param0 int) {}

func Fn_gtk_radio_action_set_group(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_button_get_group() {}

func Fn_gtk_radio_button_join_group(param0 unsafe.Pointer) {}

func Fn_gtk_radio_button_set_group(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_menu_item_new_with_label(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_label_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_new_with_mnemonic_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_menu_item_get_group() {}

func Fn_gtk_radio_menu_item_set_group(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_from_stock(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_new_from_widget(param0 unsafe.Pointer) {}

func Fn_gtk_radio_tool_button_new_with_stock_from_widget(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_radio_tool_button_get_group() {}

func Fn_gtk_radio_tool_button_set_group(param0 unsafe.Pointer) {}

func Fn_gtk_range_get_adjustment() {}

func Fn_gtk_range_get_fill_level() {}

func Fn_gtk_range_get_flippable() {}

func Fn_gtk_range_get_inverted() {}

func Fn_gtk_range_get_lower_stepper_sensitivity() {}

func Fn_gtk_range_get_min_slider_size() {}

func Fn_gtk_range_get_range_rect(param0 unsafe.Pointer) {}

func Fn_gtk_range_get_restrict_to_fill_level() {}

func Fn_gtk_range_get_round_digits() {}

func Fn_gtk_range_get_show_fill_level() {}

func Fn_gtk_range_get_slider_range(param0 *int, param1 *int) {}

func Fn_gtk_range_get_slider_size_fixed() {}

func Fn_gtk_range_get_upper_stepper_sensitivity() {}

func Fn_gtk_range_get_value() {}

func Fn_gtk_range_set_adjustment(param0 unsafe.Pointer) {}

func Fn_gtk_range_set_fill_level(param0 float64) {}

func Fn_gtk_range_set_flippable(param0 bool) {}

func Fn_gtk_range_set_increments(param0 float64, param1 float64) {}

func Fn_gtk_range_set_inverted(param0 bool) {}

func Fn_gtk_range_set_lower_stepper_sensitivity(param0 int) {}

func Fn_gtk_range_set_min_slider_size(param0 int) {}

func Fn_gtk_range_set_range(param0 float64, param1 float64) {}

func Fn_gtk_range_set_restrict_to_fill_level(param0 bool) {}

func Fn_gtk_range_set_round_digits(param0 int) {}

func Fn_gtk_range_set_show_fill_level(param0 bool) {}

func Fn_gtk_range_set_slider_size_fixed(param0 bool) {}

func Fn_gtk_range_set_upper_stepper_sensitivity(param0 int) {}

func Fn_gtk_range_set_value(param0 float64) {}

func Fn_gtk_rc_style_new() {
	C.gtk_rc_style_new()
}

func Fn_gtk_rc_style_copy() {}

func Fn_gtk_recent_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_recent_action_new_for_manager(param0 string, param1 string, param2 string, param3 string, param4 unsafe.Pointer) {
}

func Fn_gtk_recent_action_get_show_numbers() {}

func Fn_gtk_recent_action_set_show_numbers(param0 bool) {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : new_for_manager : has varargs

func Fn_gtk_recent_chooser_menu_new() {
	C.gtk_recent_chooser_menu_new()
}

func Fn_gtk_recent_chooser_menu_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_chooser_menu_get_show_numbers() {}

func Fn_gtk_recent_chooser_menu_set_show_numbers(param0 bool) {}

func Fn_gtk_recent_chooser_widget_new() {
	C.gtk_recent_chooser_widget_new()
}

func Fn_gtk_recent_chooser_widget_new_for_manager(param0 unsafe.Pointer) {}

func Fn_gtk_recent_filter_new() {
	C.gtk_recent_filter_new()
}

func Fn_gtk_recent_filter_add_age(param0 int) {}

func Fn_gtk_recent_filter_add_application(param0 string) {}

// UNSUPPORTED : add_custom : has callback

func Fn_gtk_recent_filter_add_group(param0 string) {}

func Fn_gtk_recent_filter_add_mime_type(param0 string) {}

func Fn_gtk_recent_filter_add_pattern(param0 string) {}

func Fn_gtk_recent_filter_add_pixbuf_formats() {}

func Fn_gtk_recent_filter_filter(param0 unsafe.Pointer) {}

func Fn_gtk_recent_filter_get_name() {}

func Fn_gtk_recent_filter_get_needed() {}

func Fn_gtk_recent_filter_set_name(param0 string) {}

func Fn_gtk_recent_manager_new() {
	C.gtk_recent_manager_new()
}

func Fn_gtk_recent_manager_add_full(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_recent_manager_add_item(param0 string) {}

func Fn_gtk_recent_manager_get_items() {}

func Fn_gtk_recent_manager_has_item(param0 string) {}

func Fn_gtk_recent_manager_lookup_item(param0 string) {}

func Fn_gtk_recent_manager_move_item(param0 string, param1 string) {}

func Fn_gtk_recent_manager_purge_items() {}

func Fn_gtk_recent_manager_remove_item(param0 string) {}

func Fn_gtk_recent_manager_get_default() {
	C.gtk_recent_manager_get_default()
}

func Fn_gtk_renderer_cell_accessible_new(param0 unsafe.Pointer) {}

func Fn_gtk_revealer_new() {
	C.gtk_revealer_new()
}

func Fn_gtk_revealer_get_child_revealed() {}

func Fn_gtk_revealer_get_reveal_child() {}

func Fn_gtk_revealer_get_transition_duration() {}

func Fn_gtk_revealer_get_transition_type() {}

func Fn_gtk_revealer_set_reveal_child(param0 bool) {}

func Fn_gtk_revealer_set_transition_duration(param0 uint) {}

func Fn_gtk_revealer_set_transition_type(param0 int) {}

func Fn_gtk_scale_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scale_new_with_range(param0 int, param1 float64, param2 float64, param3 float64) {}

func Fn_gtk_scale_add_mark(param0 float64, param1 int, param2 string) {}

func Fn_gtk_scale_clear_marks() {}

func Fn_gtk_scale_get_digits() {}

func Fn_gtk_scale_get_draw_value() {}

func Fn_gtk_scale_get_has_origin() {}

func Fn_gtk_scale_get_layout() {}

func Fn_gtk_scale_get_layout_offsets(param0 *int, param1 *int) {}

func Fn_gtk_scale_get_value_pos() {}

func Fn_gtk_scale_set_digits(param0 int) {}

func Fn_gtk_scale_set_draw_value(param0 bool) {}

func Fn_gtk_scale_set_has_origin(param0 bool) {}

func Fn_gtk_scale_set_value_pos(param0 int) {}

func Fn_gtk_scale_button_new(param0 int, param1 float64, param2 float64, param3 float64, param4 []string) {
}

func Fn_gtk_scale_button_get_adjustment() {}

func Fn_gtk_scale_button_get_minus_button() {}

func Fn_gtk_scale_button_get_plus_button() {}

func Fn_gtk_scale_button_get_popup() {}

func Fn_gtk_scale_button_get_value() {}

func Fn_gtk_scale_button_set_adjustment(param0 unsafe.Pointer) {}

func Fn_gtk_scale_button_set_icons(param0 []string) {}

func Fn_gtk_scale_button_set_value(param0 float64) {}

func Fn_gtk_scrollbar_new(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_add_with_viewport(param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_get_capture_button_press() {}

func Fn_gtk_scrolled_window_get_hadjustment() {}

func Fn_gtk_scrolled_window_get_hscrollbar() {}

func Fn_gtk_scrolled_window_get_kinetic_scrolling() {}

func Fn_gtk_scrolled_window_get_min_content_height() {}

func Fn_gtk_scrolled_window_get_min_content_width() {}

func Fn_gtk_scrolled_window_get_placement() {}

func Fn_gtk_scrolled_window_get_policy(param0 int, param1 int) {}

func Fn_gtk_scrolled_window_get_shadow_type() {}

func Fn_gtk_scrolled_window_get_vadjustment() {}

func Fn_gtk_scrolled_window_get_vscrollbar() {}

func Fn_gtk_scrolled_window_set_capture_button_press(param0 bool) {}

func Fn_gtk_scrolled_window_set_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_set_kinetic_scrolling(param0 bool) {}

func Fn_gtk_scrolled_window_set_min_content_height(param0 int) {}

func Fn_gtk_scrolled_window_set_min_content_width(param0 int) {}

func Fn_gtk_scrolled_window_set_placement(param0 int) {}

func Fn_gtk_scrolled_window_set_policy(param0 int, param1 int) {}

func Fn_gtk_scrolled_window_set_shadow_type(param0 int) {}

func Fn_gtk_scrolled_window_set_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_scrolled_window_unset_placement() {}

func Fn_gtk_search_bar_new() {
	C.gtk_search_bar_new()
}

func Fn_gtk_search_bar_connect_entry(param0 unsafe.Pointer) {}

func Fn_gtk_search_bar_get_search_mode() {}

func Fn_gtk_search_bar_get_show_close_button() {}

func Fn_gtk_search_bar_handle_event(param0 unsafe.Pointer) {}

func Fn_gtk_search_bar_set_search_mode(param0 bool) {}

func Fn_gtk_search_bar_set_show_close_button(param0 bool) {}

func Fn_gtk_search_entry_new() {
	C.gtk_search_entry_new()
}

func Fn_gtk_separator_new(param0 int) {}

func Fn_gtk_separator_menu_item_new() {
	C.gtk_separator_menu_item_new()
}

func Fn_gtk_separator_tool_item_new() {
	C.gtk_separator_tool_item_new()
}

func Fn_gtk_separator_tool_item_get_draw() {}

func Fn_gtk_separator_tool_item_set_draw(param0 bool) {}

func Fn_gtk_settings_set_double_property(param0 string, param1 float64, param2 string) {}

func Fn_gtk_settings_set_long_property(param0 string, param1 int64, param2 string) {}

func Fn_gtk_settings_set_property_value(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_settings_set_string_property(param0 string, param1 string, param2 string) {}

func Fn_gtk_settings_get_default() {
	C.gtk_settings_get_default()
}

func Fn_gtk_settings_get_for_screen(param0 unsafe.Pointer) {}

func Fn_gtk_settings_install_property(param0 unsafe.Pointer) {}

// UNSUPPORTED : install_property_parser : has callback

func Fn_gtk_size_group_new(param0 int) {}

func Fn_gtk_size_group_add_widget(param0 unsafe.Pointer) {}

func Fn_gtk_size_group_get_ignore_hidden() {}

func Fn_gtk_size_group_get_mode() {}

func Fn_gtk_size_group_get_widgets() {}

func Fn_gtk_size_group_remove_widget(param0 unsafe.Pointer) {}

func Fn_gtk_size_group_set_ignore_hidden(param0 bool) {}

func Fn_gtk_size_group_set_mode(param0 int) {}

func Fn_gtk_socket_new() {
	C.gtk_socket_new()
}

func Fn_gtk_socket_add_id(param0 uint64) {}

func Fn_gtk_socket_get_id() {}

func Fn_gtk_socket_get_plug_window() {}

func Fn_gtk_spin_button_new(param0 unsafe.Pointer, param1 float64, param2 uint) {}

func Fn_gtk_spin_button_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_spin_button_configure(param0 unsafe.Pointer, param1 float64, param2 uint) {}

func Fn_gtk_spin_button_get_adjustment() {}

func Fn_gtk_spin_button_get_digits() {}

func Fn_gtk_spin_button_get_increments(param0 *float64, param1 *float64) {}

func Fn_gtk_spin_button_get_numeric() {}

func Fn_gtk_spin_button_get_range(param0 *float64, param1 *float64) {}

func Fn_gtk_spin_button_get_snap_to_ticks() {}

func Fn_gtk_spin_button_get_update_policy() {}

func Fn_gtk_spin_button_get_value() {}

func Fn_gtk_spin_button_get_value_as_int() {}

func Fn_gtk_spin_button_get_wrap() {}

func Fn_gtk_spin_button_set_adjustment(param0 unsafe.Pointer) {}

func Fn_gtk_spin_button_set_digits(param0 uint) {}

func Fn_gtk_spin_button_set_increments(param0 float64, param1 float64) {}

func Fn_gtk_spin_button_set_numeric(param0 bool) {}

func Fn_gtk_spin_button_set_range(param0 float64, param1 float64) {}

func Fn_gtk_spin_button_set_snap_to_ticks(param0 bool) {}

func Fn_gtk_spin_button_set_update_policy(param0 int) {}

func Fn_gtk_spin_button_set_value(param0 float64) {}

func Fn_gtk_spin_button_set_wrap(param0 bool) {}

func Fn_gtk_spin_button_spin(param0 int, param1 float64) {}

func Fn_gtk_spin_button_update() {}

func Fn_gtk_spinner_new() {
	C.gtk_spinner_new()
}

func Fn_gtk_spinner_start() {}

func Fn_gtk_spinner_stop() {}

func Fn_gtk_stack_new() {
	C.gtk_stack_new()
}

func Fn_gtk_stack_add_named(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_stack_add_titled(param0 unsafe.Pointer, param1 string, param2 string) {}

func Fn_gtk_stack_get_child_by_name(param0 string) {}

func Fn_gtk_stack_get_homogeneous() {}

func Fn_gtk_stack_get_transition_duration() {}

func Fn_gtk_stack_get_transition_running() {}

func Fn_gtk_stack_get_transition_type() {}

func Fn_gtk_stack_get_visible_child() {}

func Fn_gtk_stack_get_visible_child_name() {}

func Fn_gtk_stack_set_homogeneous(param0 bool) {}

func Fn_gtk_stack_set_transition_duration(param0 uint) {}

func Fn_gtk_stack_set_transition_type(param0 int) {}

func Fn_gtk_stack_set_visible_child(param0 unsafe.Pointer) {}

func Fn_gtk_stack_set_visible_child_full(param0 string, param1 int) {}

func Fn_gtk_stack_set_visible_child_name(param0 string) {}

func Fn_gtk_stack_switcher_new() {
	C.gtk_stack_switcher_new()
}

func Fn_gtk_stack_switcher_get_stack() {}

func Fn_gtk_stack_switcher_set_stack(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new() {
	C.gtk_status_icon_new()
}

func Fn_gtk_status_icon_new_from_file(param0 string) {}

func Fn_gtk_status_icon_new_from_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_icon_name(param0 string) {}

func Fn_gtk_status_icon_new_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_new_from_stock(param0 string) {}

func Fn_gtk_status_icon_get_geometry(param0 *unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_status_icon_get_gicon() {}

func Fn_gtk_status_icon_get_has_tooltip() {}

func Fn_gtk_status_icon_get_icon_name() {}

func Fn_gtk_status_icon_get_pixbuf() {}

func Fn_gtk_status_icon_get_screen() {}

func Fn_gtk_status_icon_get_size() {}

func Fn_gtk_status_icon_get_stock() {}

func Fn_gtk_status_icon_get_storage_type() {}

func Fn_gtk_status_icon_get_title() {}

func Fn_gtk_status_icon_get_tooltip_markup() {}

func Fn_gtk_status_icon_get_tooltip_text() {}

func Fn_gtk_status_icon_get_visible() {}

func Fn_gtk_status_icon_get_x11_window_id() {}

func Fn_gtk_status_icon_is_embedded() {}

func Fn_gtk_status_icon_set_from_file(param0 string) {}

func Fn_gtk_status_icon_set_from_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_from_icon_name(param0 string) {}

func Fn_gtk_status_icon_set_from_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_from_stock(param0 string) {}

func Fn_gtk_status_icon_set_has_tooltip(param0 bool) {}

func Fn_gtk_status_icon_set_name(param0 string) {}

func Fn_gtk_status_icon_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_status_icon_set_title(param0 string) {}

func Fn_gtk_status_icon_set_tooltip_markup(param0 string) {}

func Fn_gtk_status_icon_set_tooltip_text(param0 string) {}

func Fn_gtk_status_icon_set_visible(param0 bool) {}

func Fn_gtk_status_icon_position_menu(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *bool, param4 *unsafe.Pointer) {
}

func Fn_gtk_statusbar_new() {
	C.gtk_statusbar_new()
}

func Fn_gtk_statusbar_get_context_id(param0 string) {}

func Fn_gtk_statusbar_get_message_area() {}

func Fn_gtk_statusbar_pop(param0 uint) {}

func Fn_gtk_statusbar_push(param0 uint, param1 string) {}

func Fn_gtk_statusbar_remove(param0 uint, param1 uint) {}

func Fn_gtk_statusbar_remove_all(param0 uint) {}

func Fn_gtk_style_new() {
	C.gtk_style_new()
}

func Fn_gtk_style_apply_default_background(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 int, param4 int, param5 int, param6 int) {
}

func Fn_gtk_style_attach(param0 unsafe.Pointer) {}

func Fn_gtk_style_copy() {}

func Fn_gtk_style_detach() {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_get_style_property(param0 uint64, param1 string, param2 unsafe.Pointer) {}

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_has_context() {}

func Fn_gtk_style_lookup_color(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_style_lookup_icon_set(param0 string) {}

func Fn_gtk_style_render_icon(param0 unsafe.Pointer, param1 int, param2 int, param3 int, param4 unsafe.Pointer, param5 string) {
}

func Fn_gtk_style_set_background(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_style_context_new() {
	C.gtk_style_context_new()
}

func Fn_gtk_style_context_add_class(param0 string) {}

func Fn_gtk_style_context_add_provider(param0 unsafe.Pointer, param1 uint) {}

func Fn_gtk_style_context_add_region(param0 string, param1 int) {}

func Fn_gtk_style_context_cancel_animations(param0 *unsafe.Pointer) {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_context_get_background_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_border(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_border_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_direction() {}

func Fn_gtk_style_context_get_font(param0 int) {}

func Fn_gtk_style_context_get_frame_clock() {}

func Fn_gtk_style_context_get_junction_sides() {}

func Fn_gtk_style_context_get_margin(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_padding(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_get_parent() {}

func Fn_gtk_style_context_get_path() {}

func Fn_gtk_style_context_get_property(param0 string, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_style_context_get_scale() {}

func Fn_gtk_style_context_get_screen() {}

func Fn_gtk_style_context_get_section(param0 string) {}

func Fn_gtk_style_context_get_state() {}

// UNSUPPORTED : get_style : has varargs

func Fn_gtk_style_context_get_style_property(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_context_has_class(param0 string) {}

func Fn_gtk_style_context_has_region(param0 string, param1 int) {}

func Fn_gtk_style_context_invalidate() {}

func Fn_gtk_style_context_list_classes() {}

func Fn_gtk_style_context_list_regions() {}

func Fn_gtk_style_context_lookup_color(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_lookup_icon_set(param0 string) {}

func Fn_gtk_style_context_notify_state_change(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 int, param3 bool) {
}

func Fn_gtk_style_context_pop_animatable_region() {}

func Fn_gtk_style_context_push_animatable_region(param0 *unsafe.Pointer) {}

func Fn_gtk_style_context_remove_class(param0 string) {}

func Fn_gtk_style_context_remove_provider(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_remove_region(param0 string) {}

func Fn_gtk_style_context_restore() {}

func Fn_gtk_style_context_save() {}

func Fn_gtk_style_context_scroll_animations(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_style_context_set_background(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_direction(param0 int) {}

func Fn_gtk_style_context_set_frame_clock(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_junction_sides(param0 int) {}

func Fn_gtk_style_context_set_parent(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_path(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_scale(param0 int) {}

func Fn_gtk_style_context_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_style_context_set_state(param0 int) {}

func Fn_gtk_style_context_state_is_running(param0 int, param1 *float64) {}

func Fn_gtk_style_context_add_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 uint) {
}

func Fn_gtk_style_context_remove_provider_for_screen(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_style_context_reset_widgets(param0 unsafe.Pointer) {}

func Fn_gtk_style_properties_new() {
	C.gtk_style_properties_new()
}

func Fn_gtk_style_properties_clear() {}

// UNSUPPORTED : get : has varargs

func Fn_gtk_style_properties_get_property(param0 string, param1 int, param2 unsafe.Pointer) {}

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_style_properties_lookup_color(param0 string) {}

func Fn_gtk_style_properties_map_color(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_style_properties_merge(param0 unsafe.Pointer, param1 bool) {}

// UNSUPPORTED : set : has varargs

func Fn_gtk_style_properties_set_property(param0 string, param1 int, param2 unsafe.Pointer) {}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_style_properties_unset_property(param0 string, param1 int) {}

// UNSUPPORTED : lookup_property : has callback

// UNSUPPORTED : register_property : has callback

func Fn_gtk_switch_new() {
	C.gtk_switch_new()
}

func Fn_gtk_switch_get_active() {}

func Fn_gtk_switch_get_state() {}

func Fn_gtk_switch_set_active(param0 bool) {}

func Fn_gtk_switch_set_state(param0 bool) {}

func Fn_gtk_table_new(param0 uint, param1 uint, param2 bool) {}

func Fn_gtk_table_attach(param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint, param5 int, param6 int, param7 uint, param8 uint) {
}

func Fn_gtk_table_attach_defaults(param0 unsafe.Pointer, param1 uint, param2 uint, param3 uint, param4 uint) {
}

func Fn_gtk_table_get_col_spacing(param0 uint) {}

func Fn_gtk_table_get_default_col_spacing() {}

func Fn_gtk_table_get_default_row_spacing() {}

func Fn_gtk_table_get_homogeneous() {}

func Fn_gtk_table_get_row_spacing(param0 uint) {}

func Fn_gtk_table_get_size(param0 *uint, param1 *uint) {}

func Fn_gtk_table_resize(param0 uint, param1 uint) {}

func Fn_gtk_table_set_col_spacing(param0 uint, param1 uint) {}

func Fn_gtk_table_set_col_spacings(param0 uint) {}

func Fn_gtk_table_set_homogeneous(param0 bool) {}

func Fn_gtk_table_set_row_spacing(param0 uint, param1 uint) {}

func Fn_gtk_table_set_row_spacings(param0 uint) {}

func Fn_gtk_tearoff_menu_item_new() {
	C.gtk_tearoff_menu_item_new()
}

func Fn_gtk_text_buffer_new(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_add_mark(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_add_selection_clipboard(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_apply_tag(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_apply_tag_by_name(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_backspace(param0 unsafe.Pointer, param1 bool, param2 bool) {}

func Fn_gtk_text_buffer_begin_user_action() {}

func Fn_gtk_text_buffer_copy_clipboard(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_create_child_anchor(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_create_mark(param0 string, param1 unsafe.Pointer, param2 bool) {}

// UNSUPPORTED : create_tag : has varargs

func Fn_gtk_text_buffer_cut_clipboard(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_text_buffer_delete(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_delete_interactive(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {
}

func Fn_gtk_text_buffer_delete_mark(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_delete_mark_by_name(param0 string) {}

func Fn_gtk_text_buffer_delete_selection(param0 bool, param1 bool) {}

func Fn_gtk_text_buffer_deserialize(param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 []uint8, param4 uint64) {
}

func Fn_gtk_text_buffer_deserialize_get_can_create_tags(param0 gdk.Atom) {}

func Fn_gtk_text_buffer_deserialize_set_can_create_tags(param0 gdk.Atom, param1 bool) {}

func Fn_gtk_text_buffer_end_user_action() {}

func Fn_gtk_text_buffer_get_bounds(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_char_count() {}

func Fn_gtk_text_buffer_get_copy_target_list() {}

func Fn_gtk_text_buffer_get_deserialize_formats(param0 *int) {}

func Fn_gtk_text_buffer_get_end_iter(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_has_selection() {}

func Fn_gtk_text_buffer_get_insert() {}

func Fn_gtk_text_buffer_get_iter_at_child_anchor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_iter_at_line(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_text_buffer_get_iter_at_line_index(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_text_buffer_get_iter_at_line_offset(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_text_buffer_get_iter_at_mark(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_iter_at_offset(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_text_buffer_get_line_count() {}

func Fn_gtk_text_buffer_get_mark(param0 string) {}

func Fn_gtk_text_buffer_get_modified() {}

func Fn_gtk_text_buffer_get_paste_target_list() {}

func Fn_gtk_text_buffer_get_selection_bound() {}

func Fn_gtk_text_buffer_get_selection_bounds(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_serialize_formats(param0 *int) {}

func Fn_gtk_text_buffer_get_slice(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_text_buffer_get_start_iter(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_get_tag_table() {}

func Fn_gtk_text_buffer_get_text(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_text_buffer_insert(param0 unsafe.Pointer, param1 string, param2 int) {}

func Fn_gtk_text_buffer_insert_at_cursor(param0 string, param1 int) {}

func Fn_gtk_text_buffer_insert_child_anchor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_insert_interactive(param0 unsafe.Pointer, param1 string, param2 int, param3 bool) {
}

func Fn_gtk_text_buffer_insert_interactive_at_cursor(param0 string, param1 int, param2 bool) {}

func Fn_gtk_text_buffer_insert_pixbuf(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_insert_range(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_insert_range_interactive(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
}

// UNSUPPORTED : insert_with_tags : has varargs

// UNSUPPORTED : insert_with_tags_by_name : has varargs

func Fn_gtk_text_buffer_move_mark(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_move_mark_by_name(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_paste_clipboard(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_text_buffer_place_cursor(param0 unsafe.Pointer) {}

// UNSUPPORTED : register_deserialize_format : has callback

func Fn_gtk_text_buffer_register_deserialize_tagset(param0 string) {}

// UNSUPPORTED : register_serialize_format : has callback

func Fn_gtk_text_buffer_register_serialize_tagset(param0 string) {}

func Fn_gtk_text_buffer_remove_all_tags(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_remove_selection_clipboard(param0 unsafe.Pointer) {}

func Fn_gtk_text_buffer_remove_tag(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_remove_tag_by_name(param0 string, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_buffer_select_range(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_buffer_serialize(param0 unsafe.Pointer, param1 gdk.Atom, param2 unsafe.Pointer, param3 unsafe.Pointer, param4 *uint64) {
}

func Fn_gtk_text_buffer_set_modified(param0 bool) {}

func Fn_gtk_text_buffer_set_text(param0 string, param1 int) {}

func Fn_gtk_text_buffer_unregister_deserialize_format(param0 gdk.Atom) {}

func Fn_gtk_text_buffer_unregister_serialize_format(param0 gdk.Atom) {}

func Fn_gtk_text_child_anchor_new() {
	C.gtk_text_child_anchor_new()
}

func Fn_gtk_text_child_anchor_get_deleted() {}

func Fn_gtk_text_child_anchor_get_widgets() {}

func Fn_gtk_text_mark_new(param0 string, param1 bool) {}

func Fn_gtk_text_mark_get_buffer() {}

func Fn_gtk_text_mark_get_deleted() {}

func Fn_gtk_text_mark_get_left_gravity() {}

func Fn_gtk_text_mark_get_name() {}

func Fn_gtk_text_mark_get_visible() {}

func Fn_gtk_text_mark_set_visible(param0 bool) {}

func Fn_gtk_text_tag_new(param0 string) {}

func Fn_gtk_text_tag_event(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {}

func Fn_gtk_text_tag_get_priority() {}

func Fn_gtk_text_tag_set_priority(param0 int) {}

func Fn_gtk_text_tag_table_new() {
	C.gtk_text_tag_table_new()
}

func Fn_gtk_text_tag_table_add(param0 unsafe.Pointer) {}

// UNSUPPORTED : foreach : has callback

func Fn_gtk_text_tag_table_get_size() {}

func Fn_gtk_text_tag_table_lookup(param0 string) {}

func Fn_gtk_text_tag_table_remove(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_new() {
	C.gtk_text_view_new()
}

func Fn_gtk_text_view_new_with_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_add_child_at_anchor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_view_add_child_in_window(param0 unsafe.Pointer, param1 int, param2 int, param3 int) {
}

func Fn_gtk_text_view_backward_display_line(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_backward_display_line_start(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_buffer_to_window_coords(param0 int, param1 int, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_text_view_forward_display_line(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_forward_display_line_end(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_get_accepts_tab() {}

func Fn_gtk_text_view_get_border_window_size(param0 int) {}

func Fn_gtk_text_view_get_buffer() {}

func Fn_gtk_text_view_get_cursor_locations(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_text_view_get_cursor_visible() {}

func Fn_gtk_text_view_get_default_attributes() {}

func Fn_gtk_text_view_get_editable() {}

func Fn_gtk_text_view_get_hadjustment() {}

func Fn_gtk_text_view_get_indent() {}

func Fn_gtk_text_view_get_input_hints() {}

func Fn_gtk_text_view_get_input_purpose() {}

func Fn_gtk_text_view_get_iter_at_location(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_text_view_get_iter_at_position(param0 unsafe.Pointer, param1 *int, param2 int, param3 int) {
}

func Fn_gtk_text_view_get_iter_location(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_text_view_get_justification() {}

func Fn_gtk_text_view_get_left_margin() {}

func Fn_gtk_text_view_get_line_at_y(param0 unsafe.Pointer, param1 int, param2 *int) {}

func Fn_gtk_text_view_get_line_yrange(param0 unsafe.Pointer, param1 *int, param2 *int) {}

func Fn_gtk_text_view_get_overwrite() {}

func Fn_gtk_text_view_get_pixels_above_lines() {}

func Fn_gtk_text_view_get_pixels_below_lines() {}

func Fn_gtk_text_view_get_pixels_inside_wrap() {}

func Fn_gtk_text_view_get_right_margin() {}

func Fn_gtk_text_view_get_tabs() {}

func Fn_gtk_text_view_get_vadjustment() {}

func Fn_gtk_text_view_get_visible_rect(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_get_window(param0 int) {}

func Fn_gtk_text_view_get_window_type(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_get_wrap_mode() {}

func Fn_gtk_text_view_im_context_filter_keypress(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_move_child(param0 unsafe.Pointer, param1 int, param2 int) {}

func Fn_gtk_text_view_move_mark_onscreen(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_move_visually(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_text_view_place_cursor_onscreen() {}

func Fn_gtk_text_view_reset_im_context() {}

func Fn_gtk_text_view_scroll_mark_onscreen(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_scroll_to_iter(param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
}

func Fn_gtk_text_view_scroll_to_mark(param0 unsafe.Pointer, param1 float64, param2 bool, param3 float64, param4 float64) {
}

func Fn_gtk_text_view_set_accepts_tab(param0 bool) {}

func Fn_gtk_text_view_set_border_window_size(param0 int, param1 int) {}

func Fn_gtk_text_view_set_buffer(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_set_cursor_visible(param0 bool) {}

func Fn_gtk_text_view_set_editable(param0 bool) {}

func Fn_gtk_text_view_set_indent(param0 int) {}

func Fn_gtk_text_view_set_input_hints(param0 int) {}

func Fn_gtk_text_view_set_input_purpose(param0 int) {}

func Fn_gtk_text_view_set_justification(param0 int) {}

func Fn_gtk_text_view_set_left_margin(param0 int) {}

func Fn_gtk_text_view_set_overwrite(param0 bool) {}

func Fn_gtk_text_view_set_pixels_above_lines(param0 int) {}

func Fn_gtk_text_view_set_pixels_below_lines(param0 int) {}

func Fn_gtk_text_view_set_pixels_inside_wrap(param0 int) {}

func Fn_gtk_text_view_set_right_margin(param0 int) {}

func Fn_gtk_text_view_set_tabs(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_set_wrap_mode(param0 int) {}

func Fn_gtk_text_view_starts_display_line(param0 unsafe.Pointer) {}

func Fn_gtk_text_view_window_to_buffer_coords(param0 int, param1 int, param2 int, param3 *int, param4 *int) {
}

// UNSUPPORTED : get : has varargs

func Fn_gtk_theming_engine_get_background_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_border(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_border_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_direction() {}

func Fn_gtk_theming_engine_get_font(param0 int) {}

func Fn_gtk_theming_engine_get_junction_sides() {}

func Fn_gtk_theming_engine_get_margin(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_padding(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_path() {}

func Fn_gtk_theming_engine_get_property(param0 string, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_theming_engine_get_screen() {}

func Fn_gtk_theming_engine_get_state() {}

// UNSUPPORTED : get_style : has varargs

func Fn_gtk_theming_engine_get_style_property(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : get_style_valist : has va_list

// UNSUPPORTED : get_valist : has va_list

func Fn_gtk_theming_engine_has_class(param0 string) {}

func Fn_gtk_theming_engine_has_region(param0 string, param1 int) {}

func Fn_gtk_theming_engine_lookup_color(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_theming_engine_state_is_running(param0 int, param1 *float64) {}

func Fn_gtk_theming_engine_load(param0 string) {}

// UNSUPPORTED : register_property : has callback

func Fn_gtk_toggle_action_new(param0 string, param1 string, param2 string, param3 string) {}

func Fn_gtk_toggle_action_get_active() {}

func Fn_gtk_toggle_action_get_draw_as_radio() {}

func Fn_gtk_toggle_action_set_active(param0 bool) {}

func Fn_gtk_toggle_action_set_draw_as_radio(param0 bool) {}

func Fn_gtk_toggle_action_toggled() {}

func Fn_gtk_toggle_button_new() {
	C.gtk_toggle_button_new()
}

func Fn_gtk_toggle_button_new_with_label(param0 string) {}

func Fn_gtk_toggle_button_new_with_mnemonic(param0 string) {}

func Fn_gtk_toggle_button_get_active() {}

func Fn_gtk_toggle_button_get_inconsistent() {}

func Fn_gtk_toggle_button_get_mode() {}

func Fn_gtk_toggle_button_set_active(param0 bool) {}

func Fn_gtk_toggle_button_set_inconsistent(param0 bool) {}

func Fn_gtk_toggle_button_set_mode(param0 bool) {}

func Fn_gtk_toggle_button_toggled() {}

func Fn_gtk_toggle_tool_button_new() {
	C.gtk_toggle_tool_button_new()
}

func Fn_gtk_toggle_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_toggle_tool_button_get_active() {}

func Fn_gtk_toggle_tool_button_set_active(param0 bool) {}

func Fn_gtk_tool_button_new(param0 unsafe.Pointer, param1 string) {}

func Fn_gtk_tool_button_new_from_stock(param0 string) {}

func Fn_gtk_tool_button_get_icon_name() {}

func Fn_gtk_tool_button_get_icon_widget() {}

func Fn_gtk_tool_button_get_label() {}

func Fn_gtk_tool_button_get_label_widget() {}

func Fn_gtk_tool_button_get_stock_id() {}

func Fn_gtk_tool_button_get_use_underline() {}

func Fn_gtk_tool_button_set_icon_name(param0 string) {}

func Fn_gtk_tool_button_set_icon_widget(param0 unsafe.Pointer) {}

func Fn_gtk_tool_button_set_label(param0 string) {}

func Fn_gtk_tool_button_set_label_widget(param0 unsafe.Pointer) {}

func Fn_gtk_tool_button_set_stock_id(param0 string) {}

func Fn_gtk_tool_button_set_use_underline(param0 bool) {}

func Fn_gtk_tool_item_new() {
	C.gtk_tool_item_new()
}

func Fn_gtk_tool_item_get_ellipsize_mode() {}

func Fn_gtk_tool_item_get_expand() {}

func Fn_gtk_tool_item_get_homogeneous() {}

func Fn_gtk_tool_item_get_icon_size() {}

func Fn_gtk_tool_item_get_is_important() {}

func Fn_gtk_tool_item_get_orientation() {}

func Fn_gtk_tool_item_get_proxy_menu_item(param0 string) {}

func Fn_gtk_tool_item_get_relief_style() {}

func Fn_gtk_tool_item_get_text_alignment() {}

func Fn_gtk_tool_item_get_text_orientation() {}

func Fn_gtk_tool_item_get_text_size_group() {}

func Fn_gtk_tool_item_get_toolbar_style() {}

func Fn_gtk_tool_item_get_use_drag_window() {}

func Fn_gtk_tool_item_get_visible_horizontal() {}

func Fn_gtk_tool_item_get_visible_vertical() {}

func Fn_gtk_tool_item_rebuild_menu() {}

func Fn_gtk_tool_item_retrieve_proxy_menu_item() {}

func Fn_gtk_tool_item_set_expand(param0 bool) {}

func Fn_gtk_tool_item_set_homogeneous(param0 bool) {}

func Fn_gtk_tool_item_set_is_important(param0 bool) {}

func Fn_gtk_tool_item_set_proxy_menu_item(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_tool_item_set_tooltip_markup(param0 string) {}

func Fn_gtk_tool_item_set_tooltip_text(param0 string) {}

func Fn_gtk_tool_item_set_use_drag_window(param0 bool) {}

func Fn_gtk_tool_item_set_visible_horizontal(param0 bool) {}

func Fn_gtk_tool_item_set_visible_vertical(param0 bool) {}

func Fn_gtk_tool_item_toolbar_reconfigured() {}

func Fn_gtk_tool_item_group_new(param0 string) {}

func Fn_gtk_tool_item_group_get_collapsed() {}

func Fn_gtk_tool_item_group_get_drop_item(param0 int, param1 int) {}

func Fn_gtk_tool_item_group_get_ellipsize() {}

func Fn_gtk_tool_item_group_get_header_relief() {}

func Fn_gtk_tool_item_group_get_item_position(param0 unsafe.Pointer) {}

func Fn_gtk_tool_item_group_get_label() {}

func Fn_gtk_tool_item_group_get_label_widget() {}

func Fn_gtk_tool_item_group_get_n_items() {}

func Fn_gtk_tool_item_group_get_nth_item(param0 uint) {}

func Fn_gtk_tool_item_group_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_tool_item_group_set_collapsed(param0 bool) {}

func Fn_gtk_tool_item_group_set_ellipsize(param0 int) {}

func Fn_gtk_tool_item_group_set_header_relief(param0 int) {}

func Fn_gtk_tool_item_group_set_item_position(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_tool_item_group_set_label(param0 string) {}

func Fn_gtk_tool_item_group_set_label_widget(param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_new() {
	C.gtk_tool_palette_new()
}

func Fn_gtk_tool_palette_add_drag_dest(param0 unsafe.Pointer, param1 int, param2 int, param3 int) {}

func Fn_gtk_tool_palette_get_drag_item(param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_drop_group(param0 int, param1 int) {}

func Fn_gtk_tool_palette_get_drop_item(param0 int, param1 int) {}

func Fn_gtk_tool_palette_get_exclusive(param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_expand(param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_group_position(param0 unsafe.Pointer) {}

func Fn_gtk_tool_palette_get_hadjustment() {}

func Fn_gtk_tool_palette_get_icon_size() {}

func Fn_gtk_tool_palette_get_style() {}

func Fn_gtk_tool_palette_get_vadjustment() {}

func Fn_gtk_tool_palette_set_drag_source(param0 int) {}

func Fn_gtk_tool_palette_set_exclusive(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tool_palette_set_expand(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tool_palette_set_group_position(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_tool_palette_set_icon_size(param0 int) {}

func Fn_gtk_tool_palette_set_style(param0 int) {}

func Fn_gtk_tool_palette_unset_icon_size() {}

func Fn_gtk_tool_palette_unset_style() {}

func Fn_gtk_tool_palette_get_drag_target_group() {
	C.gtk_tool_palette_get_drag_target_group()
}

func Fn_gtk_tool_palette_get_drag_target_item() {
	C.gtk_tool_palette_get_drag_target_item()
}

func Fn_gtk_toolbar_new() {
	C.gtk_toolbar_new()
}

func Fn_gtk_toolbar_get_drop_index(param0 int, param1 int) {}

func Fn_gtk_toolbar_get_icon_size() {}

func Fn_gtk_toolbar_get_item_index(param0 unsafe.Pointer) {}

func Fn_gtk_toolbar_get_n_items() {}

func Fn_gtk_toolbar_get_nth_item(param0 int) {}

func Fn_gtk_toolbar_get_relief_style() {}

func Fn_gtk_toolbar_get_show_arrow() {}

func Fn_gtk_toolbar_get_style() {}

func Fn_gtk_toolbar_insert(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_toolbar_set_drop_highlight_item(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_toolbar_set_icon_size(param0 int) {}

func Fn_gtk_toolbar_set_show_arrow(param0 bool) {}

func Fn_gtk_toolbar_set_style(param0 int) {}

func Fn_gtk_toolbar_unset_icon_size() {}

func Fn_gtk_toolbar_unset_style() {}

func Fn_gtk_tooltip_set_custom(param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_set_icon(param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_set_icon_from_gicon(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_tooltip_set_icon_from_icon_name(param0 string, param1 int) {}

func Fn_gtk_tooltip_set_icon_from_stock(param0 string, param1 int) {}

func Fn_gtk_tooltip_set_markup(param0 string) {}

func Fn_gtk_tooltip_set_text(param0 string) {}

func Fn_gtk_tooltip_set_tip_area(param0 unsafe.Pointer) {}

func Fn_gtk_tooltip_trigger_tooltip_query(param0 unsafe.Pointer) {}

func Fn_gtk_toplevel_accessible_get_children() {}

func Fn_gtk_tree_model_filter_clear_cache() {}

func Fn_gtk_tree_model_filter_convert_child_iter_to_iter(param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_convert_child_path_to_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_filter_convert_iter_to_child_iter(param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_filter_convert_path_to_child_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_filter_get_model() {}

func Fn_gtk_tree_model_filter_refilter() {}

// UNSUPPORTED : set_modify_func : has callback

func Fn_gtk_tree_model_filter_set_visible_column(param0 int) {}

// UNSUPPORTED : set_visible_func : has callback

func Fn_gtk_tree_model_sort_clear_cache() {}

func Fn_gtk_tree_model_sort_convert_child_iter_to_iter(param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_convert_child_path_to_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_convert_iter_to_child_iter(param0 unsafe.Pointer, param1 unsafe.Pointer) {
}

func Fn_gtk_tree_model_sort_convert_path_to_child_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_get_model() {}

func Fn_gtk_tree_model_sort_iter_is_valid(param0 unsafe.Pointer) {}

func Fn_gtk_tree_model_sort_reset_default_sort_func() {}

func Fn_gtk_tree_selection_count_selected_rows() {}

func Fn_gtk_tree_selection_get_mode() {}

func Fn_gtk_tree_selection_get_select_function() {}

func Fn_gtk_tree_selection_get_selected(param0 *unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_selected_rows(param0 *unsafe.Pointer) {}

func Fn_gtk_tree_selection_get_tree_view() {}

func Fn_gtk_tree_selection_get_user_data() {}

func Fn_gtk_tree_selection_iter_is_selected(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_path_is_selected(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_all() {}

func Fn_gtk_tree_selection_select_iter(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_select_range(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : selected_foreach : has callback

func Fn_gtk_tree_selection_set_mode(param0 int) {}

// UNSUPPORTED : set_select_function : has callback

func Fn_gtk_tree_selection_unselect_all() {}

func Fn_gtk_tree_selection_unselect_iter(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_unselect_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_selection_unselect_range(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : new : has varargs

func Fn_gtk_tree_store_newv(param0 int, param1 []uint64) {}

func Fn_gtk_tree_store_append(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_store_clear() {}

func Fn_gtk_tree_store_insert(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_tree_store_insert_after(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_store_insert_before(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

// UNSUPPORTED : insert_with_values : has varargs

func Fn_gtk_tree_store_insert_with_valuesv(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int, param3 []int, param4 []gobject.Value, param5 int) {
}

func Fn_gtk_tree_store_is_ancestor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_store_iter_depth(param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_iter_is_valid(param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_move_after(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_store_move_before(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_store_prepend(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_store_remove(param0 unsafe.Pointer) {}

func Fn_gtk_tree_store_reorder(param0 unsafe.Pointer, param1 []int) {}

// UNSUPPORTED : set : has varargs

func Fn_gtk_tree_store_set_column_types(param0 int, param1 []uint64) {}

// UNSUPPORTED : set_valist : has va_list

func Fn_gtk_tree_store_set_value(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

func Fn_gtk_tree_store_set_valuesv(param0 unsafe.Pointer, param1 []int, param2 []gobject.Value, param3 int) {
}

func Fn_gtk_tree_store_swap(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_view_new() {
	C.gtk_tree_view_new()
}

func Fn_gtk_tree_view_new_with_model(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_append_column(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_collapse_all() {}

func Fn_gtk_tree_view_collapse_row(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_columns_autosize() {}

func Fn_gtk_tree_view_convert_bin_window_to_tree_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_bin_window_to_widget_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_tree_to_bin_window_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_tree_to_widget_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_widget_to_bin_window_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_convert_widget_to_tree_coords(param0 int, param1 int, param2 *int, param3 *int) {
}

func Fn_gtk_tree_view_create_row_drag_icon(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_enable_model_drag_dest(param0 []TargetEntry, param1 int, param2 int) {}

func Fn_gtk_tree_view_enable_model_drag_source(param0 int, param1 []TargetEntry, param2 int, param3 int) {
}

func Fn_gtk_tree_view_expand_all() {}

func Fn_gtk_tree_view_expand_row(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tree_view_expand_to_path(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_get_activate_on_single_click() {}

func Fn_gtk_tree_view_get_background_area(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_bin_window() {}

func Fn_gtk_tree_view_get_cell_area(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_column(param0 int) {}

func Fn_gtk_tree_view_get_columns() {}

func Fn_gtk_tree_view_get_cursor(param0 *unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_gtk_tree_view_get_dest_row_at_pos(param0 int, param1 int, param2 *unsafe.Pointer, param3 int) {
}

func Fn_gtk_tree_view_get_drag_dest_row(param0 *unsafe.Pointer, param1 int) {}

func Fn_gtk_tree_view_get_enable_search() {}

func Fn_gtk_tree_view_get_enable_tree_lines() {}

func Fn_gtk_tree_view_get_expander_column() {}

func Fn_gtk_tree_view_get_fixed_height_mode() {}

func Fn_gtk_tree_view_get_grid_lines() {}

func Fn_gtk_tree_view_get_hadjustment() {}

func Fn_gtk_tree_view_get_headers_clickable() {}

func Fn_gtk_tree_view_get_headers_visible() {}

func Fn_gtk_tree_view_get_hover_expand() {}

func Fn_gtk_tree_view_get_hover_selection() {}

func Fn_gtk_tree_view_get_level_indentation() {}

func Fn_gtk_tree_view_get_model() {}

func Fn_gtk_tree_view_get_n_columns() {}

func Fn_gtk_tree_view_get_path_at_pos(param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) {
}

func Fn_gtk_tree_view_get_reorderable() {}

func Fn_gtk_tree_view_get_row_separator_func() {}

func Fn_gtk_tree_view_get_rubber_banding() {}

func Fn_gtk_tree_view_get_rules_hint() {}

func Fn_gtk_tree_view_get_search_column() {}

func Fn_gtk_tree_view_get_search_entry() {}

func Fn_gtk_tree_view_get_search_equal_func() {}

func Fn_gtk_tree_view_get_search_position_func() {}

func Fn_gtk_tree_view_get_selection() {}

func Fn_gtk_tree_view_get_show_expanders() {}

func Fn_gtk_tree_view_get_tooltip_column() {}

func Fn_gtk_tree_view_get_tooltip_context(param0 *int, param1 *int, param2 bool, param3 *unsafe.Pointer, param4 *unsafe.Pointer, param5 unsafe.Pointer) {
}

func Fn_gtk_tree_view_get_vadjustment() {}

func Fn_gtk_tree_view_get_visible_range(param0 *unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_gtk_tree_view_get_visible_rect(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_insert_column(param0 unsafe.Pointer, param1 int) {}

// UNSUPPORTED : insert_column_with_attributes : has varargs

// UNSUPPORTED : insert_column_with_data_func : has callback

func Fn_gtk_tree_view_is_blank_at_pos(param0 int, param1 int, param2 *unsafe.Pointer, param3 *unsafe.Pointer, param4 *int, param5 *int) {
}

func Fn_gtk_tree_view_is_rubber_banding_active() {}

// UNSUPPORTED : map_expanded_rows : has callback

func Fn_gtk_tree_view_move_column_after(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_view_remove_column(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_row_activated(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_view_row_expanded(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_scroll_to_cell(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 float32, param4 float32) {
}

func Fn_gtk_tree_view_scroll_to_point(param0 int, param1 int) {}

func Fn_gtk_tree_view_set_activate_on_single_click(param0 bool) {}

// UNSUPPORTED : set_column_drag_function : has callback

func Fn_gtk_tree_view_set_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool) {}

func Fn_gtk_tree_view_set_cursor_on_cell(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 bool) {
}

// UNSUPPORTED : set_destroy_count_func : has callback

func Fn_gtk_tree_view_set_drag_dest_row(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_tree_view_set_enable_search(param0 bool) {}

func Fn_gtk_tree_view_set_enable_tree_lines(param0 bool) {}

func Fn_gtk_tree_view_set_expander_column(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_fixed_height_mode(param0 bool) {}

func Fn_gtk_tree_view_set_grid_lines(param0 int) {}

func Fn_gtk_tree_view_set_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_headers_clickable(param0 bool) {}

func Fn_gtk_tree_view_set_headers_visible(param0 bool) {}

func Fn_gtk_tree_view_set_hover_expand(param0 bool) {}

func Fn_gtk_tree_view_set_hover_selection(param0 bool) {}

func Fn_gtk_tree_view_set_level_indentation(param0 int) {}

func Fn_gtk_tree_view_set_model(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_reorderable(param0 bool) {}

// UNSUPPORTED : set_row_separator_func : has callback

func Fn_gtk_tree_view_set_rubber_banding(param0 bool) {}

func Fn_gtk_tree_view_set_rules_hint(param0 bool) {}

func Fn_gtk_tree_view_set_search_column(param0 int) {}

func Fn_gtk_tree_view_set_search_entry(param0 unsafe.Pointer) {}

// UNSUPPORTED : set_search_equal_func : has callback

// UNSUPPORTED : set_search_position_func : has callback

func Fn_gtk_tree_view_set_show_expanders(param0 bool) {}

func Fn_gtk_tree_view_set_tooltip_cell(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 unsafe.Pointer, param3 unsafe.Pointer) {
}

func Fn_gtk_tree_view_set_tooltip_column(param0 int) {}

func Fn_gtk_tree_view_set_tooltip_row(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_tree_view_set_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_unset_rows_drag_dest() {}

func Fn_gtk_tree_view_unset_rows_drag_source() {}

func Fn_gtk_tree_view_column_new() {
	C.gtk_tree_view_column_new()
}

func Fn_gtk_tree_view_column_new_with_area(param0 unsafe.Pointer) {}

// UNSUPPORTED : new_with_attributes : has varargs

func Fn_gtk_tree_view_column_add_attribute(param0 unsafe.Pointer, param1 string, param2 int) {}

func Fn_gtk_tree_view_column_cell_get_position(param0 unsafe.Pointer, param1 *int, param2 *int) {}

func Fn_gtk_tree_view_column_cell_get_size(param0 unsafe.Pointer, param1 *int, param2 *int, param3 *int, param4 *int) {
}

func Fn_gtk_tree_view_column_cell_is_visible() {}

func Fn_gtk_tree_view_column_cell_set_cell_data(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 bool, param3 bool) {
}

func Fn_gtk_tree_view_column_clear() {}

func Fn_gtk_tree_view_column_clear_attributes(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_column_clicked() {}

func Fn_gtk_tree_view_column_focus_cell(param0 unsafe.Pointer) {}

func Fn_gtk_tree_view_column_get_alignment() {}

func Fn_gtk_tree_view_column_get_button() {}

func Fn_gtk_tree_view_column_get_clickable() {}

func Fn_gtk_tree_view_column_get_expand() {}

func Fn_gtk_tree_view_column_get_fixed_width() {}

func Fn_gtk_tree_view_column_get_max_width() {}

func Fn_gtk_tree_view_column_get_min_width() {}

func Fn_gtk_tree_view_column_get_reorderable() {}

func Fn_gtk_tree_view_column_get_resizable() {}

func Fn_gtk_tree_view_column_get_sizing() {}

func Fn_gtk_tree_view_column_get_sort_column_id() {}

func Fn_gtk_tree_view_column_get_sort_indicator() {}

func Fn_gtk_tree_view_column_get_sort_order() {}

func Fn_gtk_tree_view_column_get_spacing() {}

func Fn_gtk_tree_view_column_get_title() {}

func Fn_gtk_tree_view_column_get_tree_view() {}

func Fn_gtk_tree_view_column_get_visible() {}

func Fn_gtk_tree_view_column_get_widget() {}

func Fn_gtk_tree_view_column_get_width() {}

func Fn_gtk_tree_view_column_get_x_offset() {}

func Fn_gtk_tree_view_column_pack_end(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tree_view_column_pack_start(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_tree_view_column_queue_resize() {}

func Fn_gtk_tree_view_column_set_alignment(param0 float32) {}

// UNSUPPORTED : set_attributes : has varargs

// UNSUPPORTED : set_cell_data_func : has callback

func Fn_gtk_tree_view_column_set_clickable(param0 bool) {}

func Fn_gtk_tree_view_column_set_expand(param0 bool) {}

func Fn_gtk_tree_view_column_set_fixed_width(param0 int) {}

func Fn_gtk_tree_view_column_set_max_width(param0 int) {}

func Fn_gtk_tree_view_column_set_min_width(param0 int) {}

func Fn_gtk_tree_view_column_set_reorderable(param0 bool) {}

func Fn_gtk_tree_view_column_set_resizable(param0 bool) {}

func Fn_gtk_tree_view_column_set_sizing(param0 int) {}

func Fn_gtk_tree_view_column_set_sort_column_id(param0 int) {}

func Fn_gtk_tree_view_column_set_sort_indicator(param0 bool) {}

func Fn_gtk_tree_view_column_set_sort_order(param0 int) {}

func Fn_gtk_tree_view_column_set_spacing(param0 int) {}

func Fn_gtk_tree_view_column_set_title(param0 string) {}

func Fn_gtk_tree_view_column_set_visible(param0 bool) {}

func Fn_gtk_tree_view_column_set_widget(param0 unsafe.Pointer) {}

func Fn_gtk_ui_manager_new() {
	C.gtk_ui_manager_new()
}

func Fn_gtk_ui_manager_add_ui(param0 uint, param1 string, param2 string, param3 string, param4 int, param5 bool) {
}

func Fn_gtk_ui_manager_add_ui_from_file(param0 string) {}

func Fn_gtk_ui_manager_add_ui_from_resource(param0 string) {}

func Fn_gtk_ui_manager_add_ui_from_string(param0 string, param1 uint64) {}

func Fn_gtk_ui_manager_ensure_update() {}

func Fn_gtk_ui_manager_get_accel_group() {}

func Fn_gtk_ui_manager_get_action(param0 string) {}

func Fn_gtk_ui_manager_get_action_groups() {}

func Fn_gtk_ui_manager_get_add_tearoffs() {}

func Fn_gtk_ui_manager_get_toplevels(param0 int) {}

func Fn_gtk_ui_manager_get_ui() {}

func Fn_gtk_ui_manager_get_widget(param0 string) {}

func Fn_gtk_ui_manager_insert_action_group(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_ui_manager_new_merge_id() {}

func Fn_gtk_ui_manager_remove_action_group(param0 unsafe.Pointer) {}

func Fn_gtk_ui_manager_remove_ui(param0 uint) {}

func Fn_gtk_ui_manager_set_add_tearoffs(param0 bool) {}

func Fn_gtk_vbox_new(param0 bool, param1 int) {}

func Fn_gtk_vbutton_box_new() {
	C.gtk_vbutton_box_new()
}

func Fn_gtk_vpaned_new() {
	C.gtk_vpaned_new()
}

func Fn_gtk_vscale_new(param0 unsafe.Pointer) {}

func Fn_gtk_vscale_new_with_range(param0 float64, param1 float64, param2 float64) {}

func Fn_gtk_vscrollbar_new(param0 unsafe.Pointer) {}

func Fn_gtk_vseparator_new() {
	C.gtk_vseparator_new()
}

func Fn_gtk_viewport_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_viewport_get_bin_window() {}

func Fn_gtk_viewport_get_hadjustment() {}

func Fn_gtk_viewport_get_shadow_type() {}

func Fn_gtk_viewport_get_vadjustment() {}

func Fn_gtk_viewport_get_view_window() {}

func Fn_gtk_viewport_set_hadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_viewport_set_shadow_type(param0 int) {}

func Fn_gtk_viewport_set_vadjustment(param0 unsafe.Pointer) {}

func Fn_gtk_volume_button_new() {
	C.gtk_volume_button_new()
}

// UNSUPPORTED : new : has varargs

func Fn_gtk_widget_activate() {}

func Fn_gtk_widget_add_accelerator(param0 string, param1 unsafe.Pointer, param2 uint, param3 int, param4 int) {
}

func Fn_gtk_widget_add_device_events(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_widget_add_events(param0 int) {}

func Fn_gtk_widget_add_mnemonic_label(param0 unsafe.Pointer) {}

// UNSUPPORTED : add_tick_callback : has callback

func Fn_gtk_widget_can_activate_accel(param0 uint) {}

func Fn_gtk_widget_child_focus(param0 int) {}

func Fn_gtk_widget_child_notify(param0 string) {}

func Fn_gtk_widget_class_path(param0 *uint, param1 string, param2 string) {}

func Fn_gtk_widget_compute_expand(param0 int) {}

func Fn_gtk_widget_create_pango_context() {}

func Fn_gtk_widget_create_pango_layout(param0 string) {}

func Fn_gtk_widget_destroy() {}

func Fn_gtk_widget_destroyed(param0 *unsafe.Pointer) {}

func Fn_gtk_widget_device_is_shadowed(param0 unsafe.Pointer) {}

func Fn_gtk_drag_begin(param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer) {}

func Fn_gtk_drag_begin_with_coordinates(param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 int, param5 int) {
}

func Fn_gtk_drag_check_threshold(param0 int, param1 int, param2 int, param3 int) {}

func Fn_gtk_drag_dest_add_image_targets() {}

func Fn_gtk_drag_dest_add_text_targets() {}

func Fn_gtk_drag_dest_add_uri_targets() {}

func Fn_gtk_drag_dest_find_target(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_drag_dest_get_target_list() {}

func Fn_gtk_drag_dest_get_track_motion() {}

func Fn_gtk_drag_dest_set(param0 int, param1 []TargetEntry, param2 int, param3 int) {}

func Fn_gtk_drag_dest_set_proxy(param0 unsafe.Pointer, param1 int, param2 bool) {}

func Fn_gtk_drag_dest_set_target_list(param0 unsafe.Pointer) {}

func Fn_gtk_drag_dest_set_track_motion(param0 bool) {}

func Fn_gtk_drag_dest_unset() {}

func Fn_gtk_drag_get_data(param0 unsafe.Pointer, param1 gdk.Atom, param2 uint32) {}

func Fn_gtk_drag_highlight() {}

func Fn_gtk_drag_source_add_image_targets() {}

func Fn_gtk_drag_source_add_text_targets() {}

func Fn_gtk_drag_source_add_uri_targets() {}

func Fn_gtk_drag_source_get_target_list() {}

func Fn_gtk_drag_source_set(param0 int, param1 []TargetEntry, param2 int, param3 int) {}

func Fn_gtk_drag_source_set_icon_gicon(param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_set_icon_name(param0 string) {}

func Fn_gtk_drag_source_set_icon_pixbuf(param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_set_icon_stock(param0 string) {}

func Fn_gtk_drag_source_set_target_list(param0 unsafe.Pointer) {}

func Fn_gtk_drag_source_unset() {}

func Fn_gtk_drag_unhighlight() {}

func Fn_gtk_widget_draw(param0 unsafe.Pointer) {}

func Fn_gtk_widget_ensure_style() {}

func Fn_gtk_widget_error_bell() {}

func Fn_gtk_widget_event(param0 unsafe.Pointer) {}

func Fn_gtk_widget_freeze_child_notify() {}

func Fn_gtk_widget_get_accessible() {}

func Fn_gtk_widget_get_allocated_baseline() {}

func Fn_gtk_widget_get_allocated_height() {}

func Fn_gtk_widget_get_allocated_width() {}

func Fn_gtk_widget_get_allocation(param0 gdk.Rectangle) {}

func Fn_gtk_widget_get_ancestor(param0 uint64) {}

func Fn_gtk_widget_get_app_paintable() {}

func Fn_gtk_widget_get_can_default() {}

func Fn_gtk_widget_get_can_focus() {}

func Fn_gtk_widget_get_child_requisition(param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_child_visible() {}

func Fn_gtk_widget_get_clip(param0 gdk.Rectangle) {}

func Fn_gtk_widget_get_clipboard(param0 gdk.Atom) {}

func Fn_gtk_widget_get_composite_name() {}

func Fn_gtk_widget_get_device_enabled(param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_device_events(param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_direction() {}

func Fn_gtk_widget_get_display() {}

func Fn_gtk_widget_get_double_buffered() {}

func Fn_gtk_widget_get_events() {}

func Fn_gtk_widget_get_frame_clock() {}

func Fn_gtk_widget_get_halign() {}

func Fn_gtk_widget_get_has_tooltip() {}

func Fn_gtk_widget_get_has_window() {}

func Fn_gtk_widget_get_hexpand() {}

func Fn_gtk_widget_get_hexpand_set() {}

func Fn_gtk_widget_get_mapped() {}

func Fn_gtk_widget_get_margin_bottom() {}

func Fn_gtk_widget_get_margin_end() {}

func Fn_gtk_widget_get_margin_left() {}

func Fn_gtk_widget_get_margin_right() {}

func Fn_gtk_widget_get_margin_start() {}

func Fn_gtk_widget_get_margin_top() {}

func Fn_gtk_widget_get_modifier_mask(param0 int) {}

func Fn_gtk_widget_get_modifier_style() {}

func Fn_gtk_widget_get_name() {}

func Fn_gtk_widget_get_no_show_all() {}

func Fn_gtk_widget_get_opacity() {}

func Fn_gtk_widget_get_pango_context() {}

func Fn_gtk_widget_get_parent() {}

func Fn_gtk_widget_get_parent_window() {}

func Fn_gtk_widget_get_path() {}

func Fn_gtk_widget_get_pointer(param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_height(param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_height_and_baseline_for_width(param0 int, param1 *int, param2 *int, param3 *int, param4 *int) {
}

func Fn_gtk_widget_get_preferred_height_for_width(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_widget_get_preferred_size(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_widget_get_preferred_width(param0 *int, param1 *int) {}

func Fn_gtk_widget_get_preferred_width_for_height(param0 int, param1 *int, param2 *int) {}

func Fn_gtk_widget_get_realized() {}

func Fn_gtk_widget_get_receives_default() {}

func Fn_gtk_widget_get_request_mode() {}

func Fn_gtk_widget_get_requisition(param0 unsafe.Pointer) {}

func Fn_gtk_widget_get_root_window() {}

func Fn_gtk_widget_get_scale_factor() {}

func Fn_gtk_widget_get_screen() {}

func Fn_gtk_widget_get_sensitive() {}

func Fn_gtk_widget_get_settings() {}

func Fn_gtk_widget_get_size_request(param0 *int, param1 *int) {}

func Fn_gtk_widget_get_state() {}

func Fn_gtk_widget_get_state_flags() {}

func Fn_gtk_widget_get_style() {}

func Fn_gtk_widget_get_style_context() {}

func Fn_gtk_widget_get_support_multidevice() {}

func Fn_gtk_widget_get_template_child(param0 uint64, param1 string) {}

func Fn_gtk_widget_get_tooltip_markup() {}

func Fn_gtk_widget_get_tooltip_text() {}

func Fn_gtk_widget_get_tooltip_window() {}

func Fn_gtk_widget_get_toplevel() {}

func Fn_gtk_widget_get_valign() {}

func Fn_gtk_widget_get_valign_with_baseline() {}

func Fn_gtk_widget_get_vexpand() {}

func Fn_gtk_widget_get_vexpand_set() {}

func Fn_gtk_widget_get_visible() {}

func Fn_gtk_widget_get_visual() {}

func Fn_gtk_widget_get_window() {}

func Fn_gtk_grab_add() {}

func Fn_gtk_widget_grab_default() {}

func Fn_gtk_widget_grab_focus() {}

func Fn_gtk_grab_remove() {}

func Fn_gtk_widget_has_default() {}

func Fn_gtk_widget_has_focus() {}

func Fn_gtk_widget_has_grab() {}

func Fn_gtk_widget_has_rc_style() {}

func Fn_gtk_widget_has_screen() {}

func Fn_gtk_widget_has_visible_focus() {}

func Fn_gtk_widget_hide() {}

func Fn_gtk_widget_hide_on_delete() {}

func Fn_gtk_widget_in_destruction() {}

func Fn_gtk_widget_init_template() {}

func Fn_gtk_widget_input_shape_combine_region(param0 unsafe.Pointer) {}

func Fn_gtk_widget_insert_action_group(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_widget_intersect(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_widget_is_ancestor(param0 unsafe.Pointer) {}

func Fn_gtk_widget_is_composited() {}

func Fn_gtk_widget_is_drawable() {}

func Fn_gtk_widget_is_focus() {}

func Fn_gtk_widget_is_sensitive() {}

func Fn_gtk_widget_is_toplevel() {}

func Fn_gtk_widget_is_visible() {}

func Fn_gtk_widget_keynav_failed(param0 int) {}

func Fn_gtk_widget_list_accel_closures() {}

func Fn_gtk_widget_list_mnemonic_labels() {}

func Fn_gtk_widget_map() {}

func Fn_gtk_widget_mnemonic_activate(param0 bool) {}

func Fn_gtk_widget_modify_base(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_bg(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_fg(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_modify_font(param0 unsafe.Pointer) {}

func Fn_gtk_widget_modify_style(param0 unsafe.Pointer) {}

func Fn_gtk_widget_modify_text(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_background_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_color(param0 int, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_cursor(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_gtk_widget_override_font(param0 unsafe.Pointer) {}

func Fn_gtk_widget_override_symbolic_color(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_widget_path(param0 *uint, param1 string, param2 string) {}

func Fn_gtk_widget_queue_compute_expand() {}

func Fn_gtk_widget_queue_draw() {}

func Fn_gtk_widget_queue_draw_area(param0 int, param1 int, param2 int, param3 int) {}

func Fn_gtk_widget_queue_draw_region(param0 unsafe.Pointer) {}

func Fn_gtk_widget_queue_resize() {}

func Fn_gtk_widget_queue_resize_no_redraw() {}

func Fn_gtk_widget_realize() {}

func Fn_gtk_widget_region_intersect(param0 unsafe.Pointer) {}

func Fn_gtk_widget_register_window(param0 unsafe.Pointer) {}

func Fn_gtk_widget_remove_accelerator(param0 unsafe.Pointer, param1 uint, param2 int) {}

func Fn_gtk_widget_remove_mnemonic_label(param0 unsafe.Pointer) {}

func Fn_gtk_widget_remove_tick_callback(param0 uint) {}

func Fn_gtk_widget_render_icon(param0 string, param1 int, param2 string) {}

func Fn_gtk_widget_render_icon_pixbuf(param0 string, param1 int) {}

func Fn_gtk_widget_reparent(param0 unsafe.Pointer) {}

func Fn_gtk_widget_reset_rc_styles() {}

func Fn_gtk_widget_reset_style() {}

func Fn_gtk_widget_send_expose(param0 unsafe.Pointer) {}

func Fn_gtk_widget_send_focus_change(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_accel_path(param0 string, param1 unsafe.Pointer) {}

func Fn_gtk_widget_set_allocation(param0 gdk.Rectangle) {}

func Fn_gtk_widget_set_app_paintable(param0 bool) {}

func Fn_gtk_widget_set_can_default(param0 bool) {}

func Fn_gtk_widget_set_can_focus(param0 bool) {}

func Fn_gtk_widget_set_child_visible(param0 bool) {}

func Fn_gtk_widget_set_clip(param0 gdk.Rectangle) {}

func Fn_gtk_widget_set_composite_name(param0 string) {}

func Fn_gtk_widget_set_device_enabled(param0 unsafe.Pointer, param1 bool) {}

func Fn_gtk_widget_set_device_events(param0 unsafe.Pointer, param1 int) {}

func Fn_gtk_widget_set_direction(param0 int) {}

func Fn_gtk_widget_set_double_buffered(param0 bool) {}

func Fn_gtk_widget_set_events(param0 int) {}

func Fn_gtk_widget_set_halign(param0 int) {}

func Fn_gtk_widget_set_has_tooltip(param0 bool) {}

func Fn_gtk_widget_set_has_window(param0 bool) {}

func Fn_gtk_widget_set_hexpand(param0 bool) {}

func Fn_gtk_widget_set_hexpand_set(param0 bool) {}

func Fn_gtk_widget_set_mapped(param0 bool) {}

func Fn_gtk_widget_set_margin_bottom(param0 int) {}

func Fn_gtk_widget_set_margin_end(param0 int) {}

func Fn_gtk_widget_set_margin_left(param0 int) {}

func Fn_gtk_widget_set_margin_right(param0 int) {}

func Fn_gtk_widget_set_margin_start(param0 int) {}

func Fn_gtk_widget_set_margin_top(param0 int) {}

func Fn_gtk_widget_set_name(param0 string) {}

func Fn_gtk_widget_set_no_show_all(param0 bool) {}

func Fn_gtk_widget_set_opacity(param0 float64) {}

func Fn_gtk_widget_set_parent(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_parent_window(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_realized(param0 bool) {}

func Fn_gtk_widget_set_receives_default(param0 bool) {}

func Fn_gtk_widget_set_redraw_on_allocate(param0 bool) {}

func Fn_gtk_widget_set_sensitive(param0 bool) {}

func Fn_gtk_widget_set_size_request(param0 int, param1 int) {}

func Fn_gtk_widget_set_state(param0 int) {}

func Fn_gtk_widget_set_state_flags(param0 int, param1 bool) {}

func Fn_gtk_widget_set_style(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_support_multidevice(param0 bool) {}

func Fn_gtk_widget_set_tooltip_markup(param0 string) {}

func Fn_gtk_widget_set_tooltip_text(param0 string) {}

func Fn_gtk_widget_set_tooltip_window(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_valign(param0 int) {}

func Fn_gtk_widget_set_vexpand(param0 bool) {}

func Fn_gtk_widget_set_vexpand_set(param0 bool) {}

func Fn_gtk_widget_set_visible(param0 bool) {}

func Fn_gtk_widget_set_visual(param0 unsafe.Pointer) {}

func Fn_gtk_widget_set_window(param0 unsafe.Pointer) {}

func Fn_gtk_widget_shape_combine_region(param0 unsafe.Pointer) {}

func Fn_gtk_widget_show() {}

func Fn_gtk_widget_show_all() {}

func Fn_gtk_widget_show_now() {}

func Fn_gtk_widget_size_allocate(param0 gdk.Rectangle) {}

func Fn_gtk_widget_size_allocate_with_baseline(param0 gdk.Rectangle, param1 int) {}

func Fn_gtk_widget_size_request(param0 unsafe.Pointer) {}

func Fn_gtk_widget_style_attach() {}

// UNSUPPORTED : style_get : has varargs

func Fn_gtk_widget_style_get_property(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : style_get_valist : has va_list

func Fn_gtk_widget_thaw_child_notify() {}

func Fn_gtk_widget_translate_coordinates(param0 unsafe.Pointer, param1 int, param2 int, param3 *int, param4 *int) {
}

func Fn_gtk_widget_trigger_tooltip_query() {}

func Fn_gtk_widget_unmap() {}

func Fn_gtk_widget_unparent() {}

func Fn_gtk_widget_unrealize() {}

func Fn_gtk_widget_unregister_window(param0 unsafe.Pointer) {}

func Fn_gtk_widget_unset_state_flags(param0 int) {}

func Fn_gtk_widget_get_default_direction() {
	C.gtk_widget_get_default_direction()
}

func Fn_gtk_widget_get_default_style() {
	C.gtk_widget_get_default_style()
}

func Fn_gtk_widget_pop_composite_child() {
	C.gtk_widget_pop_composite_child()
}

func Fn_gtk_widget_push_composite_child() {
	C.gtk_widget_push_composite_child()
}

func Fn_gtk_widget_set_default_direction(param0 int) {}

func Fn_gtk_window_new(param0 int) {}

func Fn_gtk_window_activate_default() {}

func Fn_gtk_window_activate_focus() {}

func Fn_gtk_window_activate_key(param0 unsafe.Pointer) {}

func Fn_gtk_window_add_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_window_add_mnemonic(param0 uint, param1 unsafe.Pointer) {}

func Fn_gtk_window_begin_move_drag(param0 int, param1 int, param2 int, param3 uint32) {}

func Fn_gtk_window_begin_resize_drag(param0 int, param1 int, param2 int, param3 int, param4 uint32) {}

func Fn_gtk_window_close() {}

func Fn_gtk_window_deiconify() {}

func Fn_gtk_window_fullscreen() {}

func Fn_gtk_window_get_accept_focus() {}

func Fn_gtk_window_get_application() {}

func Fn_gtk_window_get_attached_to() {}

func Fn_gtk_window_get_decorated() {}

func Fn_gtk_window_get_default_size(param0 *int, param1 *int) {}

func Fn_gtk_window_get_default_widget() {}

func Fn_gtk_window_get_deletable() {}

func Fn_gtk_window_get_destroy_with_parent() {}

func Fn_gtk_window_get_focus() {}

func Fn_gtk_window_get_focus_on_map() {}

func Fn_gtk_window_get_focus_visible() {}

func Fn_gtk_window_get_gravity() {}

func Fn_gtk_window_get_group() {}

func Fn_gtk_window_get_has_resize_grip() {}

func Fn_gtk_window_get_hide_titlebar_when_maximized() {}

func Fn_gtk_window_get_icon() {}

func Fn_gtk_window_get_icon_list() {}

func Fn_gtk_window_get_icon_name() {}

func Fn_gtk_window_get_mnemonic_modifier() {}

func Fn_gtk_window_get_mnemonics_visible() {}

func Fn_gtk_window_get_modal() {}

func Fn_gtk_window_get_opacity() {}

func Fn_gtk_window_get_position(param0 *int, param1 *int) {}

func Fn_gtk_window_get_resizable() {}

func Fn_gtk_window_get_resize_grip_area(param0 unsafe.Pointer) {}

func Fn_gtk_window_get_role() {}

func Fn_gtk_window_get_screen() {}

func Fn_gtk_window_get_size(param0 *int, param1 *int) {}

func Fn_gtk_window_get_skip_pager_hint() {}

func Fn_gtk_window_get_skip_taskbar_hint() {}

func Fn_gtk_window_get_title() {}

func Fn_gtk_window_get_transient_for() {}

func Fn_gtk_window_get_type_hint() {}

func Fn_gtk_window_get_urgency_hint() {}

func Fn_gtk_window_get_window_type() {}

func Fn_gtk_window_has_group() {}

func Fn_gtk_window_has_toplevel_focus() {}

func Fn_gtk_window_iconify() {}

func Fn_gtk_window_is_active() {}

func Fn_gtk_window_is_maximized() {}

func Fn_gtk_window_maximize() {}

func Fn_gtk_window_mnemonic_activate(param0 uint, param1 int) {}

func Fn_gtk_window_move(param0 int, param1 int) {}

func Fn_gtk_window_parse_geometry(param0 string) {}

func Fn_gtk_window_present() {}

func Fn_gtk_window_present_with_time(param0 uint32) {}

func Fn_gtk_window_propagate_key_event(param0 unsafe.Pointer) {}

func Fn_gtk_window_remove_accel_group(param0 unsafe.Pointer) {}

func Fn_gtk_window_remove_mnemonic(param0 uint, param1 unsafe.Pointer) {}

func Fn_gtk_window_reshow_with_initial_size() {}

func Fn_gtk_window_resize(param0 int, param1 int) {}

func Fn_gtk_window_resize_grip_is_visible() {}

func Fn_gtk_window_resize_to_geometry(param0 int, param1 int) {}

func Fn_gtk_window_set_accept_focus(param0 bool) {}

func Fn_gtk_window_set_application(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_attached_to(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_decorated(param0 bool) {}

func Fn_gtk_window_set_default(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_geometry(param0 int, param1 int) {}

func Fn_gtk_window_set_default_size(param0 int, param1 int) {}

func Fn_gtk_window_set_deletable(param0 bool) {}

func Fn_gtk_window_set_destroy_with_parent(param0 bool) {}

func Fn_gtk_window_set_focus(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_focus_on_map(param0 bool) {}

func Fn_gtk_window_set_focus_visible(param0 bool) {}

func Fn_gtk_window_set_geometry_hints(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_gtk_window_set_gravity(param0 int) {}

func Fn_gtk_window_set_has_resize_grip(param0 bool) {}

func Fn_gtk_window_set_has_user_ref_count(param0 bool) {}

func Fn_gtk_window_set_hide_titlebar_when_maximized(param0 bool) {}

func Fn_gtk_window_set_icon(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_icon_from_file(param0 string) {}

func Fn_gtk_window_set_icon_list(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_icon_name(param0 string) {}

func Fn_gtk_window_set_keep_above(param0 bool) {}

func Fn_gtk_window_set_keep_below(param0 bool) {}

func Fn_gtk_window_set_mnemonic_modifier(param0 int) {}

func Fn_gtk_window_set_mnemonics_visible(param0 bool) {}

func Fn_gtk_window_set_modal(param0 bool) {}

func Fn_gtk_window_set_opacity(param0 float64) {}

func Fn_gtk_window_set_position(param0 int) {}

func Fn_gtk_window_set_resizable(param0 bool) {}

func Fn_gtk_window_set_role(param0 string) {}

func Fn_gtk_window_set_screen(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_skip_pager_hint(param0 bool) {}

func Fn_gtk_window_set_skip_taskbar_hint(param0 bool) {}

func Fn_gtk_window_set_startup_id(param0 string) {}

func Fn_gtk_window_set_title(param0 string) {}

func Fn_gtk_window_set_titlebar(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_transient_for(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_type_hint(param0 int) {}

func Fn_gtk_window_set_urgency_hint(param0 bool) {}

func Fn_gtk_window_set_wmclass(param0 string, param1 string) {}

func Fn_gtk_window_stick() {}

func Fn_gtk_window_unfullscreen() {}

func Fn_gtk_window_unmaximize() {}

func Fn_gtk_window_unstick() {}

func Fn_gtk_window_get_default_icon_list() {
	C.gtk_window_get_default_icon_list()
}

func Fn_gtk_window_get_default_icon_name() {
	C.gtk_window_get_default_icon_name()
}

func Fn_gtk_window_list_toplevels() {
	C.gtk_window_list_toplevels()
}

func Fn_gtk_window_set_auto_startup_notification(param0 bool) {}

func Fn_gtk_window_set_default_icon(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_icon_from_file(param0 string) {}

func Fn_gtk_window_set_default_icon_list(param0 unsafe.Pointer) {}

func Fn_gtk_window_set_default_icon_name(param0 string) {}

func Fn_gtk_window_set_interactive_debugging(param0 bool) {}

func Fn_gtk_window_group_new() {
	C.gtk_window_group_new()
}

func Fn_gtk_window_group_add_window(param0 unsafe.Pointer) {}

func Fn_gtk_window_group_get_current_device_grab(param0 unsafe.Pointer) {}

func Fn_gtk_window_group_get_current_grab() {}

func Fn_gtk_window_group_list_windows() {}

func Fn_gtk_window_group_remove_window(param0 unsafe.Pointer) {}
