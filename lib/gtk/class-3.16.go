// This is a generated file - DO NOT EDIT
// +build gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// GLArea is a wrapper around the C record GtkGLArea.
type GLArea struct {
	native unsafe.Pointer
	// Private : parent_instance
}

func GLAreaNewFromC(u unsafe.Pointer) *GLArea {
	if u == nil {
		return nil
	}

	g := &GLArea{native: u}

	return g
}

func (recv *GLArea) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : gtk_gl_area_new : return type :

// Unsupported : gtk_model_button_new : return type :

// Unsupported : gtk_popover_menu_new : return type :

// Unsupported : gtk_stack_sidebar_new : return type :
