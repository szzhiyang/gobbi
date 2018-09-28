// This is a generated file - DO NOT EDIT
// +build gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.20 gtk_3.22

package gtk

// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
import "C"

type CssSectionType C.GtkCssSectionType

const (
	GTK_CSS_SECTION_DOCUMENT         CssSectionType = 0
	GTK_CSS_SECTION_IMPORT           CssSectionType = 1
	GTK_CSS_SECTION_COLOR_DEFINITION CssSectionType = 2
	GTK_CSS_SECTION_BINDING_SET      CssSectionType = 3
	GTK_CSS_SECTION_RULESET          CssSectionType = 4
	GTK_CSS_SECTION_SELECTOR         CssSectionType = 5
	GTK_CSS_SECTION_DECLARATION      CssSectionType = 6
	GTK_CSS_SECTION_VALUE            CssSectionType = 7
	GTK_CSS_SECTION_KEYFRAMES        CssSectionType = 8
)