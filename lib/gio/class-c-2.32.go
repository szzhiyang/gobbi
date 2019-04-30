// This is a generated file - DO NOT EDIT
// +build gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <stdlib.h>
import "C"

// InetAddressMask is a wrapper around the C record GInetAddressMask.
type InetAddressMask struct {
	native *C.GInetAddressMask
	// parent_instance : record
	// Private : priv
}

func InetAddressMaskNewFromC(u unsafe.Pointer) *InetAddressMask {
	c := (*C.GInetAddressMask)(u)
	if c == nil {
		return nil
	}

	g := &InetAddressMask{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *InetAddressMask) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *InetAddressMask) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Menu is a wrapper around the C record GMenu.
type Menu struct {
	native *C.GMenu
}

func MenuNewFromC(u unsafe.Pointer) *Menu {
	c := (*C.GMenu)(u)
	if c == nil {
		return nil
	}

	g := &Menu{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Menu) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Menu) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuAttributeIter is a wrapper around the C record GMenuAttributeIter.
type MenuAttributeIter struct {
	native *C.GMenuAttributeIter
	// parent_instance : record
	// priv : record
}

func MenuAttributeIterNewFromC(u unsafe.Pointer) *MenuAttributeIter {
	c := (*C.GMenuAttributeIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuAttributeIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuAttributeIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuAttributeIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuItem is a wrapper around the C record GMenuItem.
type MenuItem struct {
	native *C.GMenuItem
}

func MenuItemNewFromC(u unsafe.Pointer) *MenuItem {
	c := (*C.GMenuItem)(u)
	if c == nil {
		return nil
	}

	g := &MenuItem{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuItem) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuItem) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuLinkIter is a wrapper around the C record GMenuLinkIter.
type MenuLinkIter struct {
	native *C.GMenuLinkIter
	// parent_instance : record
	// priv : record
}

func MenuLinkIterNewFromC(u unsafe.Pointer) *MenuLinkIter {
	c := (*C.GMenuLinkIter)(u)
	if c == nil {
		return nil
	}

	g := &MenuLinkIter{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuLinkIter) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuLinkIter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// MenuModel is a wrapper around the C record GMenuModel.
type MenuModel struct {
	native *C.GMenuModel
	// parent_instance : record
	// priv : record
}

func MenuModelNewFromC(u unsafe.Pointer) *MenuModel {
	c := (*C.GMenuModel)(u)
	if c == nil {
		return nil
	}

	g := &MenuModel{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *MenuModel) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *MenuModel) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
