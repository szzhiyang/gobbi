// Code generated - DO NOT EDIT.
// +build gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"sync"
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
// #include <gio/gnetworking.h>
// #include <stdlib.h>
/*

	gboolean dtlsconnection_acceptCertificateHandler(GObject *, GTlsCertificate *, GTlsCertificateFlags, gpointer);

	static gulong DtlsConnection_signal_connect_accept_certificate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "accept-certificate", G_CALLBACK(dtlsconnection_acceptCertificateHandler), data);
	}

*/
import "C"

// Unsupported : g_socket_receive_messages : unsupported parameter messages :

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

// Equals compares this DatagramBased with another DatagramBased, and returns true if they represent the same GObject.
func (recv *DatagramBased) Equals(other *DatagramBased) bool {
	return other.ToC() == recv.ToC()
}

// ConditionCheck is a wrapper around the C function g_datagram_based_condition_check.
func (recv *DatagramBased) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_datagram_based_condition_check((*C.GDatagramBased)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// ConditionWait is a wrapper around the C function g_datagram_based_condition_wait.
func (recv *DatagramBased) ConditionWait(condition glib.IOCondition, timeout int64, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_timeout := (C.gint64)(timeout)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_datagram_based_condition_wait((*C.GDatagramBased)(recv.native), c_condition, c_timeout, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CreateSource is a wrapper around the C function g_datagram_based_create_source.
func (recv *DatagramBased) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_datagram_based_create_source((*C.GDatagramBased)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_datagram_based_receive_messages : unsupported parameter messages :

// Unsupported : g_datagram_based_send_messages : unsupported parameter messages :

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

// Equals compares this DtlsClientConnection with another DtlsClientConnection, and returns true if they represent the same GObject.
func (recv *DtlsClientConnection) Equals(other *DtlsClientConnection) bool {
	return other.ToC() == recv.ToC()
}

// DtlsClientConnectionNew is a wrapper around the C function g_dtls_client_connection_new.
func DtlsClientConnectionNew(baseSocket *DatagramBased, serverIdentity *SocketConnectable) (*DtlsClientConnection, error) {
	c_base_socket := (*C.GDatagramBased)(baseSocket.ToC())

	c_server_identity := (*C.GSocketConnectable)(serverIdentity.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_client_connection_new(c_base_socket, c_server_identity, &cThrowableError)
	retGo := DtlsClientConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAcceptedCas is a wrapper around the C function g_dtls_client_connection_get_accepted_cas.
func (recv *DtlsClientConnection) GetAcceptedCas() *glib.List {
	retC := C.g_dtls_client_connection_get_accepted_cas((*C.GDtlsClientConnection)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetServerIdentity is a wrapper around the C function g_dtls_client_connection_get_server_identity.
func (recv *DtlsClientConnection) GetServerIdentity() *SocketConnectable {
	retC := C.g_dtls_client_connection_get_server_identity((*C.GDtlsClientConnection)(recv.native))
	retGo := SocketConnectableNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetValidationFlags is a wrapper around the C function g_dtls_client_connection_get_validation_flags.
func (recv *DtlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	retC := C.g_dtls_client_connection_get_validation_flags((*C.GDtlsClientConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// SetServerIdentity is a wrapper around the C function g_dtls_client_connection_set_server_identity.
func (recv *DtlsClientConnection) SetServerIdentity(identity *SocketConnectable) {
	c_identity := (*C.GSocketConnectable)(identity.ToC())

	C.g_dtls_client_connection_set_server_identity((*C.GDtlsClientConnection)(recv.native), c_identity)

	return
}

// SetValidationFlags is a wrapper around the C function g_dtls_client_connection_set_validation_flags.
func (recv *DtlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	c_flags := (C.GTlsCertificateFlags)(flags)

	C.g_dtls_client_connection_set_validation_flags((*C.GDtlsClientConnection)(recv.native), c_flags)

	return
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

// Equals compares this DtlsConnection with another DtlsConnection, and returns true if they represent the same GObject.
func (recv *DtlsConnection) Equals(other *DtlsConnection) bool {
	return other.ToC() == recv.ToC()
}

type signalDtlsConnectionAcceptCertificateDetail struct {
	callback  DtlsConnectionSignalAcceptCertificateCallback
	handlerID C.gulong
}

var signalDtlsConnectionAcceptCertificateId int
var signalDtlsConnectionAcceptCertificateMap = make(map[int]signalDtlsConnectionAcceptCertificateDetail)
var signalDtlsConnectionAcceptCertificateLock sync.RWMutex

// DtlsConnectionSignalAcceptCertificateCallback is a callback function for a 'accept-certificate' signal emitted from a DtlsConnection.
type DtlsConnectionSignalAcceptCertificateCallback func(peerCert *TlsCertificate, errors TlsCertificateFlags) bool

/*
ConnectAcceptCertificate connects the callback to the 'accept-certificate' signal for the DtlsConnection.

The returned value represents the connection, and may be passed to DisconnectAcceptCertificate to remove it.
*/
func (recv *DtlsConnection) ConnectAcceptCertificate(callback DtlsConnectionSignalAcceptCertificateCallback) int {
	signalDtlsConnectionAcceptCertificateLock.Lock()
	defer signalDtlsConnectionAcceptCertificateLock.Unlock()

	signalDtlsConnectionAcceptCertificateId++
	instance := C.gpointer(recv.native)
	handlerID := C.DtlsConnection_signal_connect_accept_certificate(instance, C.gpointer(uintptr(signalDtlsConnectionAcceptCertificateId)))

	detail := signalDtlsConnectionAcceptCertificateDetail{callback, handlerID}
	signalDtlsConnectionAcceptCertificateMap[signalDtlsConnectionAcceptCertificateId] = detail

	return signalDtlsConnectionAcceptCertificateId
}

/*
DisconnectAcceptCertificate disconnects a callback from the 'accept-certificate' signal for the DtlsConnection.

The connectionID should be a value returned from a call to ConnectAcceptCertificate.
*/
func (recv *DtlsConnection) DisconnectAcceptCertificate(connectionID int) {
	signalDtlsConnectionAcceptCertificateLock.Lock()
	defer signalDtlsConnectionAcceptCertificateLock.Unlock()

	detail, exists := signalDtlsConnectionAcceptCertificateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDtlsConnectionAcceptCertificateMap, connectionID)
}

//export dtlsconnection_acceptCertificateHandler
func dtlsconnection_acceptCertificateHandler(_ *C.GObject, c_peer_cert *C.GTlsCertificate, c_errors C.GTlsCertificateFlags, data C.gpointer) C.gboolean {
	signalDtlsConnectionAcceptCertificateLock.RLock()
	defer signalDtlsConnectionAcceptCertificateLock.RUnlock()

	peerCert := TlsCertificateNewFromC(unsafe.Pointer(c_peer_cert))

	errors := TlsCertificateFlags(c_errors)

	index := int(uintptr(data))
	callback := signalDtlsConnectionAcceptCertificateMap[index].callback
	retGo := callback(peerCert, errors)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// Close is a wrapper around the C function g_dtls_connection_close.
func (recv *DtlsConnection) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_close((*C.GDtlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dtls_connection_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_dtls_connection_close_finish.
func (recv *DtlsConnection) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_close_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EmitAcceptCertificate is a wrapper around the C function g_dtls_connection_emit_accept_certificate.
func (recv *DtlsConnection) EmitAcceptCertificate(peerCert *TlsCertificate, errors TlsCertificateFlags) bool {
	c_peer_cert := (*C.GTlsCertificate)(C.NULL)
	if peerCert != nil {
		c_peer_cert = (*C.GTlsCertificate)(peerCert.ToC())
	}

	c_errors := (C.GTlsCertificateFlags)(errors)

	retC := C.g_dtls_connection_emit_accept_certificate((*C.GDtlsConnection)(recv.native), c_peer_cert, c_errors)
	retGo := retC == C.TRUE

	return retGo
}

// GetCertificate is a wrapper around the C function g_dtls_connection_get_certificate.
func (recv *DtlsConnection) GetCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDatabase is a wrapper around the C function g_dtls_connection_get_database.
func (recv *DtlsConnection) GetDatabase() *TlsDatabase {
	retC := C.g_dtls_connection_get_database((*C.GDtlsConnection)(recv.native))
	retGo := TlsDatabaseNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetInteraction is a wrapper around the C function g_dtls_connection_get_interaction.
func (recv *DtlsConnection) GetInteraction() *TlsInteraction {
	retC := C.g_dtls_connection_get_interaction((*C.GDtlsConnection)(recv.native))
	retGo := TlsInteractionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificate is a wrapper around the C function g_dtls_connection_get_peer_certificate.
func (recv *DtlsConnection) GetPeerCertificate() *TlsCertificate {
	retC := C.g_dtls_connection_get_peer_certificate((*C.GDtlsConnection)(recv.native))
	retGo := TlsCertificateNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPeerCertificateErrors is a wrapper around the C function g_dtls_connection_get_peer_certificate_errors.
func (recv *DtlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	retC := C.g_dtls_connection_get_peer_certificate_errors((*C.GDtlsConnection)(recv.native))
	retGo := (TlsCertificateFlags)(retC)

	return retGo
}

// GetRehandshakeMode is a wrapper around the C function g_dtls_connection_get_rehandshake_mode.
func (recv *DtlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	retC := C.g_dtls_connection_get_rehandshake_mode((*C.GDtlsConnection)(recv.native))
	retGo := (TlsRehandshakeMode)(retC)

	return retGo
}

// GetRequireCloseNotify is a wrapper around the C function g_dtls_connection_get_require_close_notify.
func (recv *DtlsConnection) GetRequireCloseNotify() bool {
	retC := C.g_dtls_connection_get_require_close_notify((*C.GDtlsConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Handshake is a wrapper around the C function g_dtls_connection_handshake.
func (recv *DtlsConnection) Handshake(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_handshake((*C.GDtlsConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dtls_connection_handshake_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// HandshakeFinish is a wrapper around the C function g_dtls_connection_handshake_finish.
func (recv *DtlsConnection) HandshakeFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_handshake_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetCertificate is a wrapper around the C function g_dtls_connection_set_certificate.
func (recv *DtlsConnection) SetCertificate(certificate *TlsCertificate) {
	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	C.g_dtls_connection_set_certificate((*C.GDtlsConnection)(recv.native), c_certificate)

	return
}

// SetDatabase is a wrapper around the C function g_dtls_connection_set_database.
func (recv *DtlsConnection) SetDatabase(database *TlsDatabase) {
	c_database := (*C.GTlsDatabase)(C.NULL)
	if database != nil {
		c_database = (*C.GTlsDatabase)(database.ToC())
	}

	C.g_dtls_connection_set_database((*C.GDtlsConnection)(recv.native), c_database)

	return
}

// SetInteraction is a wrapper around the C function g_dtls_connection_set_interaction.
func (recv *DtlsConnection) SetInteraction(interaction *TlsInteraction) {
	c_interaction := (*C.GTlsInteraction)(C.NULL)
	if interaction != nil {
		c_interaction = (*C.GTlsInteraction)(interaction.ToC())
	}

	C.g_dtls_connection_set_interaction((*C.GDtlsConnection)(recv.native), c_interaction)

	return
}

// SetRehandshakeMode is a wrapper around the C function g_dtls_connection_set_rehandshake_mode.
func (recv *DtlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	c_mode := (C.GTlsRehandshakeMode)(mode)

	C.g_dtls_connection_set_rehandshake_mode((*C.GDtlsConnection)(recv.native), c_mode)

	return
}

// SetRequireCloseNotify is a wrapper around the C function g_dtls_connection_set_require_close_notify.
func (recv *DtlsConnection) SetRequireCloseNotify(requireCloseNotify bool) {
	c_require_close_notify :=
		boolToGboolean(requireCloseNotify)

	C.g_dtls_connection_set_require_close_notify((*C.GDtlsConnection)(recv.native), c_require_close_notify)

	return
}

// Shutdown is a wrapper around the C function g_dtls_connection_shutdown.
func (recv *DtlsConnection) Shutdown(shutdownRead bool, shutdownWrite bool, cancellable *Cancellable) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_shutdown((*C.GDtlsConnection)(recv.native), c_shutdown_read, c_shutdown_write, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_dtls_connection_shutdown_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ShutdownFinish is a wrapper around the C function g_dtls_connection_shutdown_finish.
func (recv *DtlsConnection) ShutdownFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_dtls_connection_shutdown_finish((*C.GDtlsConnection)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
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

// Equals compares this DtlsServerConnection with another DtlsServerConnection, and returns true if they represent the same GObject.
func (recv *DtlsServerConnection) Equals(other *DtlsServerConnection) bool {
	return other.ToC() == recv.ToC()
}

// DtlsServerConnectionNew is a wrapper around the C function g_dtls_server_connection_new.
func DtlsServerConnectionNew(baseSocket *DatagramBased, certificate *TlsCertificate) (*DtlsServerConnection, error) {
	c_base_socket := (*C.GDatagramBased)(baseSocket.ToC())

	c_certificate := (*C.GTlsCertificate)(C.NULL)
	if certificate != nil {
		c_certificate = (*C.GTlsCertificate)(certificate.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_dtls_server_connection_new(c_base_socket, c_certificate, &cThrowableError)
	retGo := DtlsServerConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ToString is a wrapper around the C function g_socket_connectable_to_string.
func (recv *SocketConnectable) ToString() string {
	retC := C.g_socket_connectable_to_string((*C.GSocketConnectable)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetDtlsClientConnectionType is a wrapper around the C function g_tls_backend_get_dtls_client_connection_type.
func (recv *TlsBackend) GetDtlsClientConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_client_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// GetDtlsServerConnectionType is a wrapper around the C function g_tls_backend_get_dtls_server_connection_type.
func (recv *TlsBackend) GetDtlsServerConnectionType() gobject.Type {
	retC := C.g_tls_backend_get_dtls_server_connection_type((*C.GTlsBackend)(recv.native))
	retGo := (gobject.Type)(retC)

	return retGo
}

// SupportsDtls is a wrapper around the C function g_tls_backend_supports_dtls.
func (recv *TlsBackend) SupportsDtls() bool {
	retC := C.g_tls_backend_supports_dtls((*C.GTlsBackend)(recv.native))
	retGo := retC == C.TRUE

	return retGo
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

// Equals compares this DatagramBasedInterface with another DatagramBasedInterface, and returns true if they represent the same GObject.
func (recv *DatagramBasedInterface) Equals(other *DatagramBasedInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this DtlsClientConnectionInterface with another DtlsClientConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsClientConnectionInterface) Equals(other *DtlsClientConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this DtlsConnectionInterface with another DtlsConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsConnectionInterface) Equals(other *DtlsConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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

// Equals compares this DtlsServerConnectionInterface with another DtlsServerConnectionInterface, and returns true if they represent the same GObject.
func (recv *DtlsServerConnectionInterface) Equals(other *DtlsServerConnectionInterface) bool {
	return other.ToC() == recv.ToC()
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
	recv.native.num_vectors =
		(C.guint)(recv.NumVectors)
	recv.native.bytes_received =
		(C.gsize)(recv.BytesReceived)
	recv.native.flags =
		(C.gint)(recv.Flags)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InputMessage with another InputMessage, and returns true if they represent the same GObject.
func (recv *InputMessage) Equals(other *InputMessage) bool {
	return other.ToC() == recv.ToC()
}