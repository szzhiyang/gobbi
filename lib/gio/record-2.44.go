// This is a generated file - DO NOT EDIT
// +build gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// #define GLIB_DISABLE_DEPRECATION_WARNINGS
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

// ListModelInterface is a wrapper around the C record GListModelInterface.
type ListModelInterface struct {
	native *C.GListModelInterface
	// g_iface : record
	// no type for get_item_type
	// no type for get_n_items
	// no type for get_item
}

func ListModelInterfaceNewFromC(u unsafe.Pointer) *ListModelInterface {
	c := (*C.GListModelInterface)(u)
	if c == nil {
		return nil
	}

	g := &ListModelInterface{native: c}

	return g
}

func (recv *ListModelInterface) toC() *C.GListModelInterface {

	return recv.native
}

// OutputMessage is a wrapper around the C record GOutputMessage.
type OutputMessage struct {
	native *C.GOutputMessage
	// address : no type generator for SocketAddress, GSocketAddress*
	// vectors : record
	NumVectors uint32
	BytesSent  uint32
	// no type for control_messages
	NumControlMessages uint32
}

func OutputMessageNewFromC(u unsafe.Pointer) *OutputMessage {
	c := (*C.GOutputMessage)(u)
	if c == nil {
		return nil
	}

	g := &OutputMessage{
		BytesSent:          (uint32)(c.bytes_sent),
		NumControlMessages: (uint32)(c.num_control_messages),
		NumVectors:         (uint32)(c.num_vectors),
		native:             c,
	}

	return g
}

func (recv *OutputMessage) toC() *C.GOutputMessage {
	recv.native.num_vectors =
		(C.guint)(recv.NumVectors)
	recv.native.bytes_sent =
		(C.guint)(recv.BytesSent)
	recv.native.num_control_messages =
		(C.guint)(recv.NumControlMessages)

	return recv.native
}