// This is a generated file - DO NOT EDIT
// +build gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

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

// GetNetworkMetered is a wrapper around the C function g_network_monitor_get_network_metered.
func (recv *NetworkMonitor) GetNetworkMetered() bool {
	retC := C.g_network_monitor_get_network_metered((*C.GNetworkMonitor)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CopySessionState is a wrapper around the C function g_tls_client_connection_copy_session_state.
func (recv *TlsClientConnection) CopySessionState(source *TlsClientConnection) {
	c_source := (*C.GTlsClientConnection)(source.ToC())

	C.g_tls_client_connection_copy_session_state((*C.GTlsClientConnection)(recv.native), c_source)

	return
}
