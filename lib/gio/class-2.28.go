// This is a generated file - DO NOT EDIT
// +build gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import "unsafe"

// Application is a wrapper around the C record GApplication.
type Application struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func ApplicationNewFromC(u unsafe.Pointer) *Application {
	if u == nil {
		return nil
	}

	g := &Application{native: u}

	return g
}

func (recv *Application) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_application_new : return type :

// Unsupported : g_simple_action_new : return type :

// Unsupported : g_simple_action_new_stateful : return type :

// SimpleActionGroup is a wrapper around the C record GSimpleActionGroup.
type SimpleActionGroup struct {
	native unsafe.Pointer
	// Private : parent_instance
	// Private : priv
}

func SimpleActionGroupNewFromC(u unsafe.Pointer) *SimpleActionGroup {
	if u == nil {
		return nil
	}

	g := &SimpleActionGroup{native: u}

	return g
}

func (recv *SimpleActionGroup) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_simple_action_group_new : return type :

// Unsupported : g_simple_async_result_new_take_error : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_tcp_wrapper_connection_new : return type :

// TlsCertificate is a wrapper around the C record GTlsCertificate.
type TlsCertificate struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TlsCertificateNewFromC(u unsafe.Pointer) *TlsCertificate {
	if u == nil {
		return nil
	}

	g := &TlsCertificate{native: u}

	return g
}

func (recv *TlsCertificate) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_tls_certificate_new_from_file : return type :

// Unsupported : g_tls_certificate_new_from_files : return type :

// Unsupported : g_tls_certificate_new_from_pem : return type :

// TlsConnection is a wrapper around the C record GTlsConnection.
type TlsConnection struct {
	native unsafe.Pointer
	// parent_instance : record
	// priv : record
}

func TlsConnectionNewFromC(u unsafe.Pointer) *TlsConnection {
	if u == nil {
		return nil
	}

	g := &TlsConnection{native: u}

	return g
}

func (recv *TlsConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
