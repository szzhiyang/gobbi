// This is a generated file - DO NOT EDIT
// +build gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Unsupported : g_memory_input_stream_new_from_bytes : return type :

// Unsupported : g_menu_item_new_from_model : return type :

// TestDBus is a wrapper around the C record GTestDBus.
type TestDBus struct {
	native unsafe.Pointer
}

func TestDBusNewFromC(u unsafe.Pointer) *TestDBus {
	if u == nil {
		return nil
	}

	g := &TestDBus{native: u}

	return g
}

func (recv *TestDBus) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_test_dbus_new : return type :
