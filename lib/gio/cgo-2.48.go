// This is a generated file - DO NOT EDIT
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
import "C"

// Unsupported : g_dtls_client_connection_new : return type :

// Unsupported : g_dtls_server_connection_new : return type :

// DatagramBased is a wrapper around the C record GDatagramBased.
type DatagramBased struct {
	native *C.GDatagramBased
}

func DatagramBasedNewFromC(u unsafe.Pointer) *DatagramBased {
	c := (*C.GDatagramBased)(u)
	if c == nil {
		return nil
	}

	g := &DatagramBased{native: c}

	return g
}

func (recv *DatagramBased) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsClientConnection is a wrapper around the C record GDtlsClientConnection.
type DtlsClientConnection struct {
	native *C.GDtlsClientConnection
}

func DtlsClientConnectionNewFromC(u unsafe.Pointer) *DtlsClientConnection {
	c := (*C.GDtlsClientConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsClientConnection{native: c}

	return g
}

func (recv *DtlsClientConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsConnection is a wrapper around the C record GDtlsConnection.
type DtlsConnection struct {
	native *C.GDtlsConnection
}

func DtlsConnectionNewFromC(u unsafe.Pointer) *DtlsConnection {
	c := (*C.GDtlsConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsConnection{native: c}

	return g
}

func (recv *DtlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsServerConnection is a wrapper around the C record GDtlsServerConnection.
type DtlsServerConnection struct {
	native *C.GDtlsServerConnection
}

func DtlsServerConnectionNewFromC(u unsafe.Pointer) *DtlsServerConnection {
	c := (*C.GDtlsServerConnection)(u)
	if c == nil {
		return nil
	}

	g := &DtlsServerConnection{native: c}

	return g
}

func (recv *DtlsServerConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DatagramBasedInterface is a wrapper around the C record GDatagramBasedInterface.
type DatagramBasedInterface struct {
	native *C.GDatagramBasedInterface
	// g_iface : record
	// no type for receive_messages
	// no type for send_messages
	// no type for create_source
	// no type for condition_check
	// no type for condition_wait
}

func DatagramBasedInterfaceNewFromC(u unsafe.Pointer) *DatagramBasedInterface {
	c := (*C.GDatagramBasedInterface)(u)
	if c == nil {
		return nil
	}

	g := &DatagramBasedInterface{native: c}

	return g
}

func (recv *DatagramBasedInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsClientConnectionInterface is a wrapper around the C record GDtlsClientConnectionInterface.
type DtlsClientConnectionInterface struct {
	native *C.GDtlsClientConnectionInterface
	// g_iface : record
}

func DtlsClientConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsClientConnectionInterface {
	c := (*C.GDtlsClientConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsClientConnectionInterface{native: c}

	return g
}

func (recv *DtlsClientConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsConnectionInterface is a wrapper around the C record GDtlsConnectionInterface.
type DtlsConnectionInterface struct {
	native *C.GDtlsConnectionInterface
	// g_iface : record
	// no type for accept_certificate
	// no type for handshake
	// no type for handshake_async
	// no type for handshake_finish
	// no type for shutdown
	// no type for shutdown_async
	// no type for shutdown_finish
}

func DtlsConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsConnectionInterface {
	c := (*C.GDtlsConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsConnectionInterface{native: c}

	return g
}

func (recv *DtlsConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// DtlsServerConnectionInterface is a wrapper around the C record GDtlsServerConnectionInterface.
type DtlsServerConnectionInterface struct {
	native *C.GDtlsServerConnectionInterface
	// g_iface : record
}

func DtlsServerConnectionInterfaceNewFromC(u unsafe.Pointer) *DtlsServerConnectionInterface {
	c := (*C.GDtlsServerConnectionInterface)(u)
	if c == nil {
		return nil
	}

	g := &DtlsServerConnectionInterface{native: c}

	return g
}

func (recv *DtlsServerConnectionInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InputMessage is a wrapper around the C record GInputMessage.
type InputMessage struct {
	native *C.GInputMessage
	// address : record
	// no type for vectors
	NumVectors    uint32
	BytesReceived uint64
	Flags         int32
	// no type for control_messages
	// num_control_messages : guint* with indirection level of 1
}

func InputMessageNewFromC(u unsafe.Pointer) *InputMessage {
	c := (*C.GInputMessage)(u)
	if c == nil {
		return nil
	}

	g := &InputMessage{
		BytesReceived: (uint64)(c.bytes_received),
		Flags:         (int32)(c.flags),
		NumVectors:    (uint32)(c.num_vectors),
		native:        c,
	}

	return g
}

func (recv *InputMessage) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}