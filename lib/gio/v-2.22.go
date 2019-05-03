// Code generated - DO NOT EDIT.
// +build gio_2.22 gio_2.24 gio_2.26 gio_2.28 gio_2.30 gio_2.32 gio_2.34 gio_2.36 gio_2.38 gio_2.40 gio_2.42 gio_2.44 gio_2.46 gio_2.48 gio_2.50 gio_2.52 gio_2.54 gio_2.56

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	"runtime"
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
// #include <stdlib.h>
/*

	gboolean socketservice_incomingHandler(GObject *, GSocketConnection *, GObject *, gpointer);

	static gulong SocketService_signal_connect_incoming(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "incoming", G_CALLBACK(socketservice_incomingHandler), data);
	}

*/
/*

	gboolean threadedsocketservice_runHandler(GObject *, GSocketConnection *, GObject *, gpointer);

	static gulong ThreadedSocketService_signal_connect_run(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "run", G_CALLBACK(threadedsocketservice_runHandler), data);
	}

*/
/*

	void volumemonitor_driveStopButtonHandler(GObject *, GDrive *, gpointer);

	static gulong VolumeMonitor_signal_connect_drive_stop_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "drive-stop-button", G_CALLBACK(volumemonitor_driveStopButtonHandler), data);
	}

*/
/*

	void drive_stopButtonHandler(GObject *, gpointer);

	static gulong Drive_signal_connect_stop_button(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "stop-button", G_CALLBACK(drive_stopButtonHandler), data);
	}

*/
/*

	void mount_preUnmountHandler(GObject *, gpointer);

	static gulong Mount_signal_connect_pre_unmount(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-unmount", G_CALLBACK(mount_preUnmountHandler), data);
	}

*/
import "C"

type DriveStartFlags C.GDriveStartFlags

const (
	DRIVE_START_NONE DriveStartFlags = 0
)

type SocketMsgFlags C.GSocketMsgFlags

const (
	SOCKET_MSG_NONE      SocketMsgFlags = 0
	SOCKET_MSG_OOB       SocketMsgFlags = 1
	SOCKET_MSG_PEEK      SocketMsgFlags = 2
	SOCKET_MSG_DONTROUTE SocketMsgFlags = 4
)

// Unsupported : g_cancellable_connect : unsupported parameter callback : no type generator for GObject.Callback (GCallback) for param callback

// Disconnect is a wrapper around the C function g_cancellable_disconnect.
func (recv *Cancellable) Disconnect(handlerId uint64) {
	c_handler_id := (C.gulong)(handlerId)

	C.g_cancellable_disconnect((*C.GCancellable)(recv.native), c_handler_id)

	return
}

// MakePollfd is a wrapper around the C function g_cancellable_make_pollfd.
func (recv *Cancellable) MakePollfd(pollfd *glib.PollFD) bool {
	c_pollfd := (*C.GPollFD)(C.NULL)
	if pollfd != nil {
		c_pollfd = (*C.GPollFD)(pollfd.ToC())
	}

	retC := C.g_cancellable_make_pollfd((*C.GCancellable)(recv.native), c_pollfd)
	retGo := retC == C.TRUE

	return retGo
}

// ReleaseFd is a wrapper around the C function g_cancellable_release_fd.
func (recv *Cancellable) ReleaseFd() {
	C.g_cancellable_release_fd((*C.GCancellable)(recv.native))

	return
}

// GetEtag is a wrapper around the C function g_file_io_stream_get_etag.
func (recv *FileIOStream) GetEtag() string {
	retC := C.g_file_io_stream_get_etag((*C.GFileIOStream)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// QueryInfo is a wrapper around the C function g_file_io_stream_query_info.
func (recv *FileIOStream) QueryInfo(attributes string, cancellable *Cancellable) (*FileInfo, error) {
	c_attributes := C.CString(attributes)
	defer C.free(unsafe.Pointer(c_attributes))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_io_stream_query_info((*C.GFileIOStream)(recv.native), c_attributes, c_cancellable, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_io_stream_query_info_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// QueryInfoFinish is a wrapper around the C function g_file_io_stream_query_info_finish.
func (recv *FileIOStream) QueryInfoFinish(result *AsyncResult) (*FileInfo, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_io_stream_query_info_finish((*C.GFileIOStream)(recv.native), c_result, &cThrowableError)
	retGo := FileInfoNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAttributeStringv is a wrapper around the C function g_file_info_get_attribute_stringv.
func (recv *FileInfo) GetAttributeStringv(attribute string) []string {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	retC := C.g_file_info_get_attribute_stringv((*C.GFileInfo)(recv.native), c_attribute)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// HasNamespace is a wrapper around the C function g_file_info_has_namespace.
func (recv *FileInfo) HasNamespace(nameSpace string) bool {
	c_name_space := C.CString(nameSpace)
	defer C.free(unsafe.Pointer(c_name_space))

	retC := C.g_file_info_has_namespace((*C.GFileInfo)(recv.native), c_name_space)
	retGo := retC == C.TRUE

	return retGo
}

// SetAttributeStatus is a wrapper around the C function g_file_info_set_attribute_status.
func (recv *FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_status := (C.GFileAttributeStatus)(status)

	retC := C.g_file_info_set_attribute_status((*C.GFileInfo)(recv.native), c_attribute, c_status)
	retGo := retC == C.TRUE

	return retGo
}

// ClearPending is a wrapper around the C function g_io_stream_clear_pending.
func (recv *IOStream) ClearPending() {
	C.g_io_stream_clear_pending((*C.GIOStream)(recv.native))

	return
}

// Close is a wrapper around the C function g_io_stream_close.
func (recv *IOStream) Close(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_io_stream_close((*C.GIOStream)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_io_stream_close_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CloseFinish is a wrapper around the C function g_io_stream_close_finish.
func (recv *IOStream) CloseFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_io_stream_close_finish((*C.GIOStream)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetInputStream is a wrapper around the C function g_io_stream_get_input_stream.
func (recv *IOStream) GetInputStream() *InputStream {
	retC := C.g_io_stream_get_input_stream((*C.GIOStream)(recv.native))
	retGo := InputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetOutputStream is a wrapper around the C function g_io_stream_get_output_stream.
func (recv *IOStream) GetOutputStream() *OutputStream {
	retC := C.g_io_stream_get_output_stream((*C.GIOStream)(recv.native))
	retGo := OutputStreamNewFromC(unsafe.Pointer(retC))

	return retGo
}

// HasPending is a wrapper around the C function g_io_stream_has_pending.
func (recv *IOStream) HasPending() bool {
	retC := C.g_io_stream_has_pending((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsClosed is a wrapper around the C function g_io_stream_is_closed.
func (recv *IOStream) IsClosed() bool {
	retC := C.g_io_stream_is_closed((*C.GIOStream)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetPending is a wrapper around the C function g_io_stream_set_pending.
func (recv *IOStream) SetPending() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_io_stream_set_pending((*C.GIOStream)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// InetAddressNewAny is a wrapper around the C function g_inet_address_new_any.
func InetAddressNewAny(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_any(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewFromBytes is a wrapper around the C function g_inet_address_new_from_bytes.
func InetAddressNewFromBytes(bytes []uint8, family SocketFamily) *InetAddress {
	c_bytes_array := make([]C.guint8, len(bytes)+1, len(bytes)+1)
	for i, item := range bytes {
		c := (C.guint8)(item)
		c_bytes_array[i] = c
	}
	c_bytes_array[len(bytes)] = 0
	c_bytes_arrayPtr := &c_bytes_array[0]
	c_bytes := (*C.guint8)(unsafe.Pointer(c_bytes_arrayPtr))

	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_from_bytes(c_bytes, c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewFromString is a wrapper around the C function g_inet_address_new_from_string.
func InetAddressNewFromString(string_ string) *InetAddress {
	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	retC := C.g_inet_address_new_from_string(c_string)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// InetAddressNewLoopback is a wrapper around the C function g_inet_address_new_loopback.
func InetAddressNewLoopback(family SocketFamily) *InetAddress {
	c_family := (C.GSocketFamily)(family)

	retC := C.g_inet_address_new_loopback(c_family)
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetFamily is a wrapper around the C function g_inet_address_get_family.
func (recv *InetAddress) GetFamily() SocketFamily {
	retC := C.g_inet_address_get_family((*C.GInetAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetIsAny is a wrapper around the C function g_inet_address_get_is_any.
func (recv *InetAddress) GetIsAny() bool {
	retC := C.g_inet_address_get_is_any((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsLinkLocal is a wrapper around the C function g_inet_address_get_is_link_local.
func (recv *InetAddress) GetIsLinkLocal() bool {
	retC := C.g_inet_address_get_is_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsLoopback is a wrapper around the C function g_inet_address_get_is_loopback.
func (recv *InetAddress) GetIsLoopback() bool {
	retC := C.g_inet_address_get_is_loopback((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcGlobal is a wrapper around the C function g_inet_address_get_is_mc_global.
func (recv *InetAddress) GetIsMcGlobal() bool {
	retC := C.g_inet_address_get_is_mc_global((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcLinkLocal is a wrapper around the C function g_inet_address_get_is_mc_link_local.
func (recv *InetAddress) GetIsMcLinkLocal() bool {
	retC := C.g_inet_address_get_is_mc_link_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcNodeLocal is a wrapper around the C function g_inet_address_get_is_mc_node_local.
func (recv *InetAddress) GetIsMcNodeLocal() bool {
	retC := C.g_inet_address_get_is_mc_node_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcOrgLocal is a wrapper around the C function g_inet_address_get_is_mc_org_local.
func (recv *InetAddress) GetIsMcOrgLocal() bool {
	retC := C.g_inet_address_get_is_mc_org_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMcSiteLocal is a wrapper around the C function g_inet_address_get_is_mc_site_local.
func (recv *InetAddress) GetIsMcSiteLocal() bool {
	retC := C.g_inet_address_get_is_mc_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsMulticast is a wrapper around the C function g_inet_address_get_is_multicast.
func (recv *InetAddress) GetIsMulticast() bool {
	retC := C.g_inet_address_get_is_multicast((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIsSiteLocal is a wrapper around the C function g_inet_address_get_is_site_local.
func (recv *InetAddress) GetIsSiteLocal() bool {
	retC := C.g_inet_address_get_is_site_local((*C.GInetAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetNativeSize is a wrapper around the C function g_inet_address_get_native_size.
func (recv *InetAddress) GetNativeSize() uint64 {
	retC := C.g_inet_address_get_native_size((*C.GInetAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Blacklisted : g_inet_address_to_bytes

// ToString is a wrapper around the C function g_inet_address_to_string.
func (recv *InetAddress) ToString() string {
	retC := C.g_inet_address_to_string((*C.GInetAddress)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// InetSocketAddressNew is a wrapper around the C function g_inet_socket_address_new.
func InetSocketAddressNew(address *InetAddress, port uint16) *InetSocketAddress {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	c_port := (C.guint16)(port)

	retC := C.g_inet_socket_address_new(c_address, c_port)
	retGo := InetSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetAddress is a wrapper around the C function g_inet_socket_address_get_address.
func (recv *InetSocketAddress) GetAddress() *InetAddress {
	retC := C.g_inet_socket_address_get_address((*C.GInetSocketAddress)(recv.native))
	retGo := InetAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetPort is a wrapper around the C function g_inet_socket_address_get_port.
func (recv *InetSocketAddress) GetPort() uint16 {
	retC := C.g_inet_socket_address_get_port((*C.GInetSocketAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// Unsupported signal 'show-processes' for MountOperation : unsupported parameter processes :

// NetworkAddressNew is a wrapper around the C function g_network_address_new.
func NetworkAddressNew(hostname string, port uint16) *NetworkAddress {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	retC := C.g_network_address_new(c_hostname, c_port)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// NetworkAddressParse is a wrapper around the C function g_network_address_parse.
func NetworkAddressParse(hostAndPort string, defaultPort uint16) (*NetworkAddress, error) {
	c_host_and_port := C.CString(hostAndPort)
	defer C.free(unsafe.Pointer(c_host_and_port))

	c_default_port := (C.guint16)(defaultPort)

	var cThrowableError *C.GError

	retC := C.g_network_address_parse(c_host_and_port, c_default_port, &cThrowableError)
	retGo := NetworkAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetHostname is a wrapper around the C function g_network_address_get_hostname.
func (recv *NetworkAddress) GetHostname() string {
	retC := C.g_network_address_get_hostname((*C.GNetworkAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function g_network_address_get_port.
func (recv *NetworkAddress) GetPort() uint16 {
	retC := C.g_network_address_get_port((*C.GNetworkAddress)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// NetworkServiceNew is a wrapper around the C function g_network_service_new.
func NetworkServiceNew(service string, protocol string, domain string) *NetworkService {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	retC := C.g_network_service_new(c_service, c_protocol, c_domain)
	retGo := NetworkServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetDomain is a wrapper around the C function g_network_service_get_domain.
func (recv *NetworkService) GetDomain() string {
	retC := C.g_network_service_get_domain((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProtocol is a wrapper around the C function g_network_service_get_protocol.
func (recv *NetworkService) GetProtocol() string {
	retC := C.g_network_service_get_protocol((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetService is a wrapper around the C function g_network_service_get_service.
func (recv *NetworkService) GetService() string {
	retC := C.g_network_service_get_service((*C.GNetworkService)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// ResolverFreeAddresses is a wrapper around the C function g_resolver_free_addresses.
func ResolverFreeAddresses(addresses *glib.List) {
	c_addresses := (*C.GList)(C.NULL)
	if addresses != nil {
		c_addresses = (*C.GList)(addresses.ToC())
	}

	C.g_resolver_free_addresses(c_addresses)

	return
}

// ResolverFreeTargets is a wrapper around the C function g_resolver_free_targets.
func ResolverFreeTargets(targets *glib.List) {
	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	C.g_resolver_free_targets(c_targets)

	return
}

// ResolverGetDefault is a wrapper around the C function g_resolver_get_default.
func ResolverGetDefault() *Resolver {
	retC := C.g_resolver_get_default()
	retGo := ResolverNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LookupByAddress is a wrapper around the C function g_resolver_lookup_by_address.
func (recv *Resolver) LookupByAddress(address *InetAddress, cancellable *Cancellable) (string, error) {
	c_address := (*C.GInetAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GInetAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_address((*C.GResolver)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_resolver_lookup_by_address_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupByAddressFinish is a wrapper around the C function g_resolver_lookup_by_address_finish.
func (recv *Resolver) LookupByAddressFinish(result *AsyncResult) (string, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_address_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LookupByName is a wrapper around the C function g_resolver_lookup_by_name.
func (recv *Resolver) LookupByName(hostname string, cancellable *Cancellable) (*glib.List, error) {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_name((*C.GResolver)(recv.native), c_hostname, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_resolver_lookup_by_name_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupByNameFinish is a wrapper around the C function g_resolver_lookup_by_name_finish.
func (recv *Resolver) LookupByNameFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_by_name_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LookupService is a wrapper around the C function g_resolver_lookup_service.
func (recv *Resolver) LookupService(service string, protocol string, domain string, cancellable *Cancellable) (*glib.List, error) {
	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_service((*C.GResolver)(recv.native), c_service, c_protocol, c_domain, c_cancellable, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_resolver_lookup_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// LookupServiceFinish is a wrapper around the C function g_resolver_lookup_service_finish.
func (recv *Resolver) LookupServiceFinish(result *AsyncResult) (*glib.List, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_resolver_lookup_service_finish((*C.GResolver)(recv.native), c_result, &cThrowableError)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetDefault is a wrapper around the C function g_resolver_set_default.
func (recv *Resolver) SetDefault() {
	C.g_resolver_set_default((*C.GResolver)(recv.native))

	return
}

// Socket is a wrapper around the C record GSocket.
type Socket struct {
	native *C.GSocket
	// parent_instance : record
	// priv : record
}

func SocketNewFromC(u unsafe.Pointer) *Socket {
	c := (*C.GSocket)(u)
	if c == nil {
		return nil
	}

	g := &Socket{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *Socket) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *Socket) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Socket with another Socket, and returns true if they represent the same GObject.
func (recv *Socket) Equals(other *Socket) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *Socket) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to Socket.
// Exercise care, as this is a potentially dangerous function if the Object is not a Socket.
func CastToSocket(object *gobject.Object) *Socket {
	return SocketNewFromC(object.ToC())
}

// SocketNew is a wrapper around the C function g_socket_new.
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) (*Socket, error) {
	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	var cThrowableError *C.GError

	retC := C.g_socket_new(c_family, c_type, c_protocol, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SocketNewFromFd is a wrapper around the C function g_socket_new_from_fd.
func SocketNewFromFd(fd int32) (*Socket, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_socket_new_from_fd(c_fd, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Accept is a wrapper around the C function g_socket_accept.
func (recv *Socket) Accept(cancellable *Cancellable) (*Socket, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_accept((*C.GSocket)(recv.native), c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Bind is a wrapper around the C function g_socket_bind.
func (recv *Socket) Bind(address *SocketAddress, allowReuse bool) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_allow_reuse :=
		boolToGboolean(allowReuse)

	var cThrowableError *C.GError

	retC := C.g_socket_bind((*C.GSocket)(recv.native), c_address, c_allow_reuse, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CheckConnectResult is a wrapper around the C function g_socket_check_connect_result.
func (recv *Socket) CheckConnectResult() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_check_connect_result((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Close is a wrapper around the C function g_socket_close.
func (recv *Socket) Close() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_close((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ConditionCheck is a wrapper around the C function g_socket_condition_check.
func (recv *Socket) ConditionCheck(condition glib.IOCondition) glib.IOCondition {
	c_condition := (C.GIOCondition)(condition)

	retC := C.g_socket_condition_check((*C.GSocket)(recv.native), c_condition)
	retGo := (glib.IOCondition)(retC)

	return retGo
}

// ConditionWait is a wrapper around the C function g_socket_condition_wait.
func (recv *Socket) ConditionWait(condition glib.IOCondition, cancellable *Cancellable) (bool, error) {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_condition_wait((*C.GSocket)(recv.native), c_condition, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Connect is a wrapper around the C function g_socket_connect.
func (recv *Socket) Connect(address *SocketAddress, cancellable *Cancellable) (bool, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_connect((*C.GSocket)(recv.native), c_address, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ConnectionFactoryCreateConnection is a wrapper around the C function g_socket_connection_factory_create_connection.
func (recv *Socket) ConnectionFactoryCreateConnection() *SocketConnection {
	retC := C.g_socket_connection_factory_create_connection((*C.GSocket)(recv.native))
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateSource is a wrapper around the C function g_socket_create_source.
func (recv *Socket) CreateSource(condition glib.IOCondition, cancellable *Cancellable) *glib.Source {
	c_condition := (C.GIOCondition)(condition)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	retC := C.g_socket_create_source((*C.GSocket)(recv.native), c_condition, c_cancellable)
	retGo := glib.SourceNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetBlocking is a wrapper around the C function g_socket_get_blocking.
func (recv *Socket) GetBlocking() bool {
	retC := C.g_socket_get_blocking((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetFamily is a wrapper around the C function g_socket_get_family.
func (recv *Socket) GetFamily() SocketFamily {
	retC := C.g_socket_get_family((*C.GSocket)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetFd is a wrapper around the C function g_socket_get_fd.
func (recv *Socket) GetFd() int32 {
	retC := C.g_socket_get_fd((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetKeepalive is a wrapper around the C function g_socket_get_keepalive.
func (recv *Socket) GetKeepalive() bool {
	retC := C.g_socket_get_keepalive((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetListenBacklog is a wrapper around the C function g_socket_get_listen_backlog.
func (recv *Socket) GetListenBacklog() int32 {
	retC := C.g_socket_get_listen_backlog((*C.GSocket)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// Blacklisted : g_socket_get_local_address

// GetProtocol is a wrapper around the C function g_socket_get_protocol.
func (recv *Socket) GetProtocol() SocketProtocol {
	retC := C.g_socket_get_protocol((*C.GSocket)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// GetRemoteAddress is a wrapper around the C function g_socket_get_remote_address.
func (recv *Socket) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_get_remote_address((*C.GSocket)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSocketType is a wrapper around the C function g_socket_get_socket_type.
func (recv *Socket) GetSocketType() SocketType {
	retC := C.g_socket_get_socket_type((*C.GSocket)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// IsClosed is a wrapper around the C function g_socket_is_closed.
func (recv *Socket) IsClosed() bool {
	retC := C.g_socket_is_closed((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsConnected is a wrapper around the C function g_socket_is_connected.
func (recv *Socket) IsConnected() bool {
	retC := C.g_socket_is_connected((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Listen is a wrapper around the C function g_socket_listen.
func (recv *Socket) Listen() (bool, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_listen((*C.GSocket)(recv.native), &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Receive is a wrapper around the C function g_socket_receive.
func (recv *Socket) Receive(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive((*C.GSocket)(recv.native), c_buffer, c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReceiveFrom is a wrapper around the C function g_socket_receive_from.
func (recv *Socket) ReceiveFrom(buffer []uint8, cancellable *Cancellable) (int64, *SocketAddress, error) {
	var c_address *C.GSocketAddress

	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_receive_from((*C.GSocket)(recv.native), &c_address, c_buffer, c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	address := SocketAddressNewFromC(unsafe.Pointer(c_address))

	return retGo, address, goError
}

// Unsupported : g_socket_receive_message : unsupported parameter vectors :

// Send is a wrapper around the C function g_socket_send.
func (recv *Socket) Send(buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send((*C.GSocket)(recv.native), c_buffer, c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_send_message : unsupported parameter vectors :

// SendTo is a wrapper around the C function g_socket_send_to.
func (recv *Socket) SendTo(address *SocketAddress, buffer []uint8, cancellable *Cancellable) (int64, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_buffer_array := make([]C.guint8, len(buffer)+1, len(buffer)+1)
	for i, item := range buffer {
		c := (C.guint8)(item)
		c_buffer_array[i] = c
	}
	c_buffer_array[len(buffer)] = 0
	c_buffer_arrayPtr := &c_buffer_array[0]
	c_buffer := (*C.gchar)(unsafe.Pointer(c_buffer_arrayPtr))

	c_size := (C.gsize)(len(buffer))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_send_to((*C.GSocket)(recv.native), c_address, c_buffer, c_size, c_cancellable, &cThrowableError)
	retGo := (int64)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBlocking is a wrapper around the C function g_socket_set_blocking.
func (recv *Socket) SetBlocking(blocking bool) {
	c_blocking :=
		boolToGboolean(blocking)

	C.g_socket_set_blocking((*C.GSocket)(recv.native), c_blocking)

	return
}

// SetKeepalive is a wrapper around the C function g_socket_set_keepalive.
func (recv *Socket) SetKeepalive(keepalive bool) {
	c_keepalive :=
		boolToGboolean(keepalive)

	C.g_socket_set_keepalive((*C.GSocket)(recv.native), c_keepalive)

	return
}

// SetListenBacklog is a wrapper around the C function g_socket_set_listen_backlog.
func (recv *Socket) SetListenBacklog(backlog int32) {
	c_backlog := (C.gint)(backlog)

	C.g_socket_set_listen_backlog((*C.GSocket)(recv.native), c_backlog)

	return
}

// Shutdown is a wrapper around the C function g_socket_shutdown.
func (recv *Socket) Shutdown(shutdownRead bool, shutdownWrite bool) (bool, error) {
	c_shutdown_read :=
		boolToGboolean(shutdownRead)

	c_shutdown_write :=
		boolToGboolean(shutdownWrite)

	var cThrowableError *C.GError

	retC := C.g_socket_shutdown((*C.GSocket)(recv.native), c_shutdown_read, c_shutdown_write, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SpeaksIpv4 is a wrapper around the C function g_socket_speaks_ipv4.
func (recv *Socket) SpeaksIpv4() bool {
	retC := C.g_socket_speaks_ipv4((*C.GSocket)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_socket_address_new_from_native : unsupported parameter native : no type generator for gpointer (gpointer) for param native

// GetFamily is a wrapper around the C function g_socket_address_get_family.
func (recv *SocketAddress) GetFamily() SocketFamily {
	retC := C.g_socket_address_get_family((*C.GSocketAddress)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetNativeSize is a wrapper around the C function g_socket_address_get_native_size.
func (recv *SocketAddress) GetNativeSize() int64 {
	retC := C.g_socket_address_get_native_size((*C.GSocketAddress)(recv.native))
	retGo := (int64)(retC)

	return retGo
}

// Unsupported : g_socket_address_to_native : unsupported parameter dest : no type generator for gpointer (gpointer) for param dest

// SocketClient is a wrapper around the C record GSocketClient.
type SocketClient struct {
	native *C.GSocketClient
	// parent_instance : record
	// priv : record
}

func SocketClientNewFromC(u unsafe.Pointer) *SocketClient {
	c := (*C.GSocketClient)(u)
	if c == nil {
		return nil
	}

	g := &SocketClient{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketClient) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketClient) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketClient with another SocketClient, and returns true if they represent the same GObject.
func (recv *SocketClient) Equals(other *SocketClient) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketClient) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketClient.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketClient.
func CastToSocketClient(object *gobject.Object) *SocketClient {
	return SocketClientNewFromC(object.ToC())
}

// SocketClientNew is a wrapper around the C function g_socket_client_new.
func SocketClientNew() *SocketClient {
	retC := C.g_socket_client_new()
	retGo := SocketClientNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddApplicationProxy is a wrapper around the C function g_socket_client_add_application_proxy.
func (recv *SocketClient) AddApplicationProxy(protocol string) {
	c_protocol := C.CString(protocol)
	defer C.free(unsafe.Pointer(c_protocol))

	C.g_socket_client_add_application_proxy((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// Connect is a wrapper around the C function g_socket_client_connect.
func (recv *SocketClient) Connect(connectable *SocketConnectable, cancellable *Cancellable) (*SocketConnection, error) {
	c_connectable := (*C.GSocketConnectable)(connectable.ToC())

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect((*C.GSocketClient)(recv.native), c_connectable, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_client_connect_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectFinish is a wrapper around the C function g_socket_client_connect_finish.
func (recv *SocketClient) ConnectFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ConnectToHost is a wrapper around the C function g_socket_client_connect_to_host.
func (recv *SocketClient) ConnectToHost(hostAndPort string, defaultPort uint16, cancellable *Cancellable) (*SocketConnection, error) {
	c_host_and_port := C.CString(hostAndPort)
	defer C.free(unsafe.Pointer(c_host_and_port))

	c_default_port := (C.guint16)(defaultPort)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host((*C.GSocketClient)(recv.native), c_host_and_port, c_default_port, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_client_connect_to_host_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectToHostFinish is a wrapper around the C function g_socket_client_connect_to_host_finish.
func (recv *SocketClient) ConnectToHostFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_host_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ConnectToService is a wrapper around the C function g_socket_client_connect_to_service.
func (recv *SocketClient) ConnectToService(domain string, service string, cancellable *Cancellable) (*SocketConnection, error) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	c_service := C.CString(service)
	defer C.free(unsafe.Pointer(c_service))

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_service((*C.GSocketClient)(recv.native), c_domain, c_service, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_socket_client_connect_to_service_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ConnectToServiceFinish is a wrapper around the C function g_socket_client_connect_to_service_finish.
func (recv *SocketClient) ConnectToServiceFinish(result *AsyncResult) (*SocketConnection, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_socket_client_connect_to_service_finish((*C.GSocketClient)(recv.native), c_result, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetFamily is a wrapper around the C function g_socket_client_get_family.
func (recv *SocketClient) GetFamily() SocketFamily {
	retC := C.g_socket_client_get_family((*C.GSocketClient)(recv.native))
	retGo := (SocketFamily)(retC)

	return retGo
}

// GetLocalAddress is a wrapper around the C function g_socket_client_get_local_address.
func (recv *SocketClient) GetLocalAddress() *SocketAddress {
	retC := C.g_socket_client_get_local_address((*C.GSocketClient)(recv.native))
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetProtocol is a wrapper around the C function g_socket_client_get_protocol.
func (recv *SocketClient) GetProtocol() SocketProtocol {
	retC := C.g_socket_client_get_protocol((*C.GSocketClient)(recv.native))
	retGo := (SocketProtocol)(retC)

	return retGo
}

// GetSocketType is a wrapper around the C function g_socket_client_get_socket_type.
func (recv *SocketClient) GetSocketType() SocketType {
	retC := C.g_socket_client_get_socket_type((*C.GSocketClient)(recv.native))
	retGo := (SocketType)(retC)

	return retGo
}

// SetFamily is a wrapper around the C function g_socket_client_set_family.
func (recv *SocketClient) SetFamily(family SocketFamily) {
	c_family := (C.GSocketFamily)(family)

	C.g_socket_client_set_family((*C.GSocketClient)(recv.native), c_family)

	return
}

// SetLocalAddress is a wrapper around the C function g_socket_client_set_local_address.
func (recv *SocketClient) SetLocalAddress(address *SocketAddress) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	C.g_socket_client_set_local_address((*C.GSocketClient)(recv.native), c_address)

	return
}

// SetProtocol is a wrapper around the C function g_socket_client_set_protocol.
func (recv *SocketClient) SetProtocol(protocol SocketProtocol) {
	c_protocol := (C.GSocketProtocol)(protocol)

	C.g_socket_client_set_protocol((*C.GSocketClient)(recv.native), c_protocol)

	return
}

// SetSocketType is a wrapper around the C function g_socket_client_set_socket_type.
func (recv *SocketClient) SetSocketType(type_ SocketType) {
	c_type := (C.GSocketType)(type_)

	C.g_socket_client_set_socket_type((*C.GSocketClient)(recv.native), c_type)

	return
}

// SocketConnection is a wrapper around the C record GSocketConnection.
type SocketConnection struct {
	native *C.GSocketConnection
	// parent_instance : record
	// priv : record
}

func SocketConnectionNewFromC(u unsafe.Pointer) *SocketConnection {
	c := (*C.GSocketConnection)(u)
	if c == nil {
		return nil
	}

	g := &SocketConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketConnection with another SocketConnection, and returns true if they represent the same GObject.
func (recv *SocketConnection) Equals(other *SocketConnection) bool {
	return other.ToC() == recv.ToC()
}

// IOStream upcasts to *IOStream
func (recv *SocketConnection) IOStream() *IOStream {
	return IOStreamNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketConnection) Object() *gobject.Object {
	return recv.IOStream().Object()
}

// CastToWidget down casts any arbitrary Object to SocketConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketConnection.
func CastToSocketConnection(object *gobject.Object) *SocketConnection {
	return SocketConnectionNewFromC(object.ToC())
}

// SocketConnectionFactoryLookupType is a wrapper around the C function g_socket_connection_factory_lookup_type.
func SocketConnectionFactoryLookupType(family SocketFamily, type_ SocketType, protocolId int32) gobject.Type {
	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol_id := (C.gint)(protocolId)

	retC := C.g_socket_connection_factory_lookup_type(c_family, c_type, c_protocol_id)
	retGo := (gobject.Type)(retC)

	return retGo
}

// SocketConnectionFactoryRegisterType is a wrapper around the C function g_socket_connection_factory_register_type.
func SocketConnectionFactoryRegisterType(gType gobject.Type, family SocketFamily, type_ SocketType, protocol int32) {
	c_g_type := (C.GType)(gType)

	c_family := (C.GSocketFamily)(family)

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.gint)(protocol)

	C.g_socket_connection_factory_register_type(c_g_type, c_family, c_type, c_protocol)

	return
}

// GetLocalAddress is a wrapper around the C function g_socket_connection_get_local_address.
func (recv *SocketConnection) GetLocalAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_local_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetRemoteAddress is a wrapper around the C function g_socket_connection_get_remote_address.
func (recv *SocketConnection) GetRemoteAddress() (*SocketAddress, error) {
	var cThrowableError *C.GError

	retC := C.g_socket_connection_get_remote_address((*C.GSocketConnection)(recv.native), &cThrowableError)
	retGo := SocketAddressNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetSocket is a wrapper around the C function g_socket_connection_get_socket.
func (recv *SocketConnection) GetSocket() *Socket {
	retC := C.g_socket_connection_get_socket((*C.GSocketConnection)(recv.native))
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SocketControlMessageDeserialize is a wrapper around the C function g_socket_control_message_deserialize.
func SocketControlMessageDeserialize(level int32, type_ int32, data []uint8) *SocketControlMessage {
	c_level := (C.int)(level)

	c_type := (C.int)(type_)

	c_size := (C.gsize)(len(data))

	c_data_array := make([]C.guint8, len(data)+1, len(data)+1)
	for i, item := range data {
		c := (C.guint8)(item)
		c_data_array[i] = c
	}
	c_data_array[len(data)] = 0
	c_data_arrayPtr := &c_data_array[0]
	c_data := (C.gpointer)(unsafe.Pointer(c_data_arrayPtr))

	retC := C.g_socket_control_message_deserialize(c_level, c_type, c_size, c_data)
	retGo := SocketControlMessageNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetLevel is a wrapper around the C function g_socket_control_message_get_level.
func (recv *SocketControlMessage) GetLevel() int32 {
	retC := C.g_socket_control_message_get_level((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetMsgType is a wrapper around the C function g_socket_control_message_get_msg_type.
func (recv *SocketControlMessage) GetMsgType() int32 {
	retC := C.g_socket_control_message_get_msg_type((*C.GSocketControlMessage)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetSize is a wrapper around the C function g_socket_control_message_get_size.
func (recv *SocketControlMessage) GetSize() uint64 {
	retC := C.g_socket_control_message_get_size((*C.GSocketControlMessage)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

// Unsupported : g_socket_control_message_serialize : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// SocketListener is a wrapper around the C record GSocketListener.
type SocketListener struct {
	native *C.GSocketListener
	// parent_instance : record
	// priv : record
}

func SocketListenerNewFromC(u unsafe.Pointer) *SocketListener {
	c := (*C.GSocketListener)(u)
	if c == nil {
		return nil
	}

	g := &SocketListener{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketListener) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketListener) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketListener with another SocketListener, and returns true if they represent the same GObject.
func (recv *SocketListener) Equals(other *SocketListener) bool {
	return other.ToC() == recv.ToC()
}

// Object upcasts to *Object
func (recv *SocketListener) Object() *gobject.Object {
	return gobject.ObjectNewFromC(unsafe.Pointer(recv.native))
}

// CastToWidget down casts any arbitrary Object to SocketListener.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketListener.
func CastToSocketListener(object *gobject.Object) *SocketListener {
	return SocketListenerNewFromC(object.ToC())
}

// SocketListenerNew is a wrapper around the C function g_socket_listener_new.
func SocketListenerNew() *SocketListener {
	retC := C.g_socket_listener_new()
	retGo := SocketListenerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Accept is a wrapper around the C function g_socket_listener_accept.
func (recv *SocketListener) Accept(cancellable *Cancellable) (*SocketConnection, *gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goError
}

// Unsupported : g_socket_listener_accept_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AcceptFinish is a wrapper around the C function g_socket_listener_accept_finish.
func (recv *SocketListener) AcceptFinish(result *AsyncResult) (*SocketConnection, *gobject.Object, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_source_object *C.GObject

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_finish((*C.GSocketListener)(recv.native), c_result, &c_source_object, &cThrowableError)
	retGo := SocketConnectionNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goError
}

// AcceptSocket is a wrapper around the C function g_socket_listener_accept_socket.
func (recv *SocketListener) AcceptSocket(cancellable *Cancellable) (*Socket, *gobject.Object, error) {
	var c_source_object *C.GObject

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_socket((*C.GSocketListener)(recv.native), &c_source_object, c_cancellable, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goError
}

// Unsupported : g_socket_listener_accept_socket_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// AcceptSocketFinish is a wrapper around the C function g_socket_listener_accept_socket_finish.
func (recv *SocketListener) AcceptSocketFinish(result *AsyncResult) (*Socket, *gobject.Object, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var c_source_object *C.GObject

	var cThrowableError *C.GError

	retC := C.g_socket_listener_accept_socket_finish((*C.GSocketListener)(recv.native), c_result, &c_source_object, &cThrowableError)
	retGo := SocketNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	return retGo, sourceObject, goError
}

// AddAddress is a wrapper around the C function g_socket_listener_add_address.
func (recv *SocketListener) AddAddress(address *SocketAddress, type_ SocketType, protocol SocketProtocol, sourceObject *gobject.Object) (bool, *SocketAddress, error) {
	c_address := (*C.GSocketAddress)(C.NULL)
	if address != nil {
		c_address = (*C.GSocketAddress)(address.ToC())
	}

	c_type := (C.GSocketType)(type_)

	c_protocol := (C.GSocketProtocol)(protocol)

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var c_effective_address *C.GSocketAddress

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_address((*C.GSocketListener)(recv.native), c_address, c_type, c_protocol, c_source_object, &c_effective_address, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	effectiveAddress := SocketAddressNewFromC(unsafe.Pointer(c_effective_address))

	return retGo, effectiveAddress, goError
}

// AddInetPort is a wrapper around the C function g_socket_listener_add_inet_port.
func (recv *SocketListener) AddInetPort(port uint16, sourceObject *gobject.Object) (bool, error) {
	c_port := (C.guint16)(port)

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_inet_port((*C.GSocketListener)(recv.native), c_port, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddSocket is a wrapper around the C function g_socket_listener_add_socket.
func (recv *SocketListener) AddSocket(socket *Socket, sourceObject *gobject.Object) (bool, error) {
	c_socket := (*C.GSocket)(C.NULL)
	if socket != nil {
		c_socket = (*C.GSocket)(socket.ToC())
	}

	c_source_object := (*C.GObject)(C.NULL)
	if sourceObject != nil {
		c_source_object = (*C.GObject)(sourceObject.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_socket_listener_add_socket((*C.GSocketListener)(recv.native), c_socket, c_source_object, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Close is a wrapper around the C function g_socket_listener_close.
func (recv *SocketListener) Close() {
	C.g_socket_listener_close((*C.GSocketListener)(recv.native))

	return
}

// SetBacklog is a wrapper around the C function g_socket_listener_set_backlog.
func (recv *SocketListener) SetBacklog(listenBacklog int32) {
	c_listen_backlog := (C.int)(listenBacklog)

	C.g_socket_listener_set_backlog((*C.GSocketListener)(recv.native), c_listen_backlog)

	return
}

// SocketService is a wrapper around the C record GSocketService.
type SocketService struct {
	native *C.GSocketService
	// parent_instance : record
	// priv : record
}

func SocketServiceNewFromC(u unsafe.Pointer) *SocketService {
	c := (*C.GSocketService)(u)
	if c == nil {
		return nil
	}

	g := &SocketService{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *SocketService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *SocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this SocketService with another SocketService, and returns true if they represent the same GObject.
func (recv *SocketService) Equals(other *SocketService) bool {
	return other.ToC() == recv.ToC()
}

// SocketListener upcasts to *SocketListener
func (recv *SocketService) SocketListener() *SocketListener {
	return SocketListenerNewFromC(unsafe.Pointer(recv.native))
}

// Object upcasts to *Object
func (recv *SocketService) Object() *gobject.Object {
	return recv.SocketListener().Object()
}

// CastToWidget down casts any arbitrary Object to SocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a SocketService.
func CastToSocketService(object *gobject.Object) *SocketService {
	return SocketServiceNewFromC(object.ToC())
}

type signalSocketServiceIncomingDetail struct {
	callback  SocketServiceSignalIncomingCallback
	handlerID C.gulong
}

var signalSocketServiceIncomingId int
var signalSocketServiceIncomingMap = make(map[int]signalSocketServiceIncomingDetail)
var signalSocketServiceIncomingLock sync.RWMutex

// SocketServiceSignalIncomingCallback is a callback function for a 'incoming' signal emitted from a SocketService.
type SocketServiceSignalIncomingCallback func(connection *SocketConnection, sourceObject *gobject.Object) bool

/*
ConnectIncoming connects the callback to the 'incoming' signal for the SocketService.

The returned value represents the connection, and may be passed to DisconnectIncoming to remove it.
*/
func (recv *SocketService) ConnectIncoming(callback SocketServiceSignalIncomingCallback) int {
	signalSocketServiceIncomingLock.Lock()
	defer signalSocketServiceIncomingLock.Unlock()

	signalSocketServiceIncomingId++
	instance := C.gpointer(recv.native)
	handlerID := C.SocketService_signal_connect_incoming(instance, C.gpointer(uintptr(signalSocketServiceIncomingId)))

	detail := signalSocketServiceIncomingDetail{callback, handlerID}
	signalSocketServiceIncomingMap[signalSocketServiceIncomingId] = detail

	return signalSocketServiceIncomingId
}

/*
DisconnectIncoming disconnects a callback from the 'incoming' signal for the SocketService.

The connectionID should be a value returned from a call to ConnectIncoming.
*/
func (recv *SocketService) DisconnectIncoming(connectionID int) {
	signalSocketServiceIncomingLock.Lock()
	defer signalSocketServiceIncomingLock.Unlock()

	detail, exists := signalSocketServiceIncomingMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalSocketServiceIncomingMap, connectionID)
}

//export socketservice_incomingHandler
func socketservice_incomingHandler(_ *C.GObject, c_connection *C.GSocketConnection, c_source_object *C.GObject, data C.gpointer) C.gboolean {
	signalSocketServiceIncomingLock.RLock()
	defer signalSocketServiceIncomingLock.RUnlock()

	connection := SocketConnectionNewFromC(unsafe.Pointer(c_connection))

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	index := int(uintptr(data))
	callback := signalSocketServiceIncomingMap[index].callback
	retGo := callback(connection, sourceObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// SocketServiceNew is a wrapper around the C function g_socket_service_new.
func SocketServiceNew() *SocketService {
	retC := C.g_socket_service_new()
	retGo := SocketServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IsActive is a wrapper around the C function g_socket_service_is_active.
func (recv *SocketService) IsActive() bool {
	retC := C.g_socket_service_is_active((*C.GSocketService)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Start is a wrapper around the C function g_socket_service_start.
func (recv *SocketService) Start() {
	C.g_socket_service_start((*C.GSocketService)(recv.native))

	return
}

// Stop is a wrapper around the C function g_socket_service_stop.
func (recv *SocketService) Stop() {
	C.g_socket_service_stop((*C.GSocketService)(recv.native))

	return
}

// TcpConnection is a wrapper around the C record GTcpConnection.
type TcpConnection struct {
	native *C.GTcpConnection
	// parent_instance : record
	// priv : record
}

func TcpConnectionNewFromC(u unsafe.Pointer) *TcpConnection {
	c := (*C.GTcpConnection)(u)
	if c == nil {
		return nil
	}

	g := &TcpConnection{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *TcpConnection) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *TcpConnection) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this TcpConnection with another TcpConnection, and returns true if they represent the same GObject.
func (recv *TcpConnection) Equals(other *TcpConnection) bool {
	return other.ToC() == recv.ToC()
}

// SocketConnection upcasts to *SocketConnection
func (recv *TcpConnection) SocketConnection() *SocketConnection {
	return SocketConnectionNewFromC(unsafe.Pointer(recv.native))
}

// IOStream upcasts to *IOStream
func (recv *TcpConnection) IOStream() *IOStream {
	return recv.SocketConnection().IOStream()
}

// Object upcasts to *Object
func (recv *TcpConnection) Object() *gobject.Object {
	return recv.SocketConnection().Object()
}

// CastToWidget down casts any arbitrary Object to TcpConnection.
// Exercise care, as this is a potentially dangerous function if the Object is not a TcpConnection.
func CastToTcpConnection(object *gobject.Object) *TcpConnection {
	return TcpConnectionNewFromC(object.ToC())
}

// GetGracefulDisconnect is a wrapper around the C function g_tcp_connection_get_graceful_disconnect.
func (recv *TcpConnection) GetGracefulDisconnect() bool {
	retC := C.g_tcp_connection_get_graceful_disconnect((*C.GTcpConnection)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetGracefulDisconnect is a wrapper around the C function g_tcp_connection_set_graceful_disconnect.
func (recv *TcpConnection) SetGracefulDisconnect(gracefulDisconnect bool) {
	c_graceful_disconnect :=
		boolToGboolean(gracefulDisconnect)

	C.g_tcp_connection_set_graceful_disconnect((*C.GTcpConnection)(recv.native), c_graceful_disconnect)

	return
}

// ThreadedSocketService is a wrapper around the C record GThreadedSocketService.
type ThreadedSocketService struct {
	native *C.GThreadedSocketService
	// parent_instance : record
	// priv : record
}

func ThreadedSocketServiceNewFromC(u unsafe.Pointer) *ThreadedSocketService {
	c := (*C.GThreadedSocketService)(u)
	if c == nil {
		return nil
	}

	g := &ThreadedSocketService{native: c}

	ug := (C.gpointer)(u)
	if C.g_object_is_floating(ug) == C.TRUE {
		C.g_object_ref_sink(ug)
	} else {
		C.g_object_ref(ug)
	}
	runtime.SetFinalizer(g, func(o *ThreadedSocketService) {
		C.g_object_unref((C.gpointer)(o.native))
	})

	return g
}

func (recv *ThreadedSocketService) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this ThreadedSocketService with another ThreadedSocketService, and returns true if they represent the same GObject.
func (recv *ThreadedSocketService) Equals(other *ThreadedSocketService) bool {
	return other.ToC() == recv.ToC()
}

// SocketService upcasts to *SocketService
func (recv *ThreadedSocketService) SocketService() *SocketService {
	return SocketServiceNewFromC(unsafe.Pointer(recv.native))
}

// SocketListener upcasts to *SocketListener
func (recv *ThreadedSocketService) SocketListener() *SocketListener {
	return recv.SocketService().SocketListener()
}

// Object upcasts to *Object
func (recv *ThreadedSocketService) Object() *gobject.Object {
	return recv.SocketService().Object()
}

// CastToWidget down casts any arbitrary Object to ThreadedSocketService.
// Exercise care, as this is a potentially dangerous function if the Object is not a ThreadedSocketService.
func CastToThreadedSocketService(object *gobject.Object) *ThreadedSocketService {
	return ThreadedSocketServiceNewFromC(object.ToC())
}

type signalThreadedSocketServiceRunDetail struct {
	callback  ThreadedSocketServiceSignalRunCallback
	handlerID C.gulong
}

var signalThreadedSocketServiceRunId int
var signalThreadedSocketServiceRunMap = make(map[int]signalThreadedSocketServiceRunDetail)
var signalThreadedSocketServiceRunLock sync.RWMutex

// ThreadedSocketServiceSignalRunCallback is a callback function for a 'run' signal emitted from a ThreadedSocketService.
type ThreadedSocketServiceSignalRunCallback func(connection *SocketConnection, sourceObject *gobject.Object) bool

/*
ConnectRun connects the callback to the 'run' signal for the ThreadedSocketService.

The returned value represents the connection, and may be passed to DisconnectRun to remove it.
*/
func (recv *ThreadedSocketService) ConnectRun(callback ThreadedSocketServiceSignalRunCallback) int {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	signalThreadedSocketServiceRunId++
	instance := C.gpointer(recv.native)
	handlerID := C.ThreadedSocketService_signal_connect_run(instance, C.gpointer(uintptr(signalThreadedSocketServiceRunId)))

	detail := signalThreadedSocketServiceRunDetail{callback, handlerID}
	signalThreadedSocketServiceRunMap[signalThreadedSocketServiceRunId] = detail

	return signalThreadedSocketServiceRunId
}

/*
DisconnectRun disconnects a callback from the 'run' signal for the ThreadedSocketService.

The connectionID should be a value returned from a call to ConnectRun.
*/
func (recv *ThreadedSocketService) DisconnectRun(connectionID int) {
	signalThreadedSocketServiceRunLock.Lock()
	defer signalThreadedSocketServiceRunLock.Unlock()

	detail, exists := signalThreadedSocketServiceRunMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalThreadedSocketServiceRunMap, connectionID)
}

//export threadedsocketservice_runHandler
func threadedsocketservice_runHandler(_ *C.GObject, c_connection *C.GSocketConnection, c_source_object *C.GObject, data C.gpointer) C.gboolean {
	signalThreadedSocketServiceRunLock.RLock()
	defer signalThreadedSocketServiceRunLock.RUnlock()

	connection := SocketConnectionNewFromC(unsafe.Pointer(c_connection))

	sourceObject := gobject.ObjectNewFromC(unsafe.Pointer(c_source_object))

	index := int(uintptr(data))
	callback := signalThreadedSocketServiceRunMap[index].callback
	retGo := callback(connection, sourceObject)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// ThreadedSocketServiceNew is a wrapper around the C function g_threaded_socket_service_new.
func ThreadedSocketServiceNew(maxThreads int32) *ThreadedSocketService {
	c_max_threads := (C.int)(maxThreads)

	retC := C.g_threaded_socket_service_new(c_max_threads)
	retGo := ThreadedSocketServiceNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// ReceiveFd is a wrapper around the C function g_unix_connection_receive_fd.
func (recv *UnixConnection) ReceiveFd(cancellable *Cancellable) (int32, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_receive_fd((*C.GUnixConnection)(recv.native), c_cancellable, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SendFd is a wrapper around the C function g_unix_connection_send_fd.
func (recv *UnixConnection) SendFd(fd int32, cancellable *Cancellable) (bool, error) {
	c_fd := (C.gint)(fd)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_unix_connection_send_fd((*C.GUnixConnection)(recv.native), c_fd, c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// UnixFDMessageNew is a wrapper around the C function g_unix_fd_message_new.
func UnixFDMessageNew() *UnixFDMessage {
	retC := C.g_unix_fd_message_new()
	retGo := UnixFDMessageNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AppendFd is a wrapper around the C function g_unix_fd_message_append_fd.
func (recv *UnixFDMessage) AppendFd(fd int32) (bool, error) {
	c_fd := (C.gint)(fd)

	var cThrowableError *C.GError

	retC := C.g_unix_fd_message_append_fd((*C.GUnixFDMessage)(recv.native), c_fd, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_unix_fd_message_steal_fds : array return type :

// UnixSocketAddressNew is a wrapper around the C function g_unix_socket_address_new.
func UnixSocketAddressNew(path string) *UnixSocketAddress {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.g_unix_socket_address_new(c_path)
	retGo := UnixSocketAddressNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// UnixSocketAddressAbstractNamesSupported is a wrapper around the C function g_unix_socket_address_abstract_names_supported.
func UnixSocketAddressAbstractNamesSupported() bool {
	retC := C.g_unix_socket_address_abstract_names_supported()
	retGo := retC == C.TRUE

	return retGo
}

// GetIsAbstract is a wrapper around the C function g_unix_socket_address_get_is_abstract.
func (recv *UnixSocketAddress) GetIsAbstract() bool {
	retC := C.g_unix_socket_address_get_is_abstract((*C.GUnixSocketAddress)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPath is a wrapper around the C function g_unix_socket_address_get_path.
func (recv *UnixSocketAddress) GetPath() string {
	retC := C.g_unix_socket_address_get_path((*C.GUnixSocketAddress)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPathLen is a wrapper around the C function g_unix_socket_address_get_path_len.
func (recv *UnixSocketAddress) GetPathLen() uint64 {
	retC := C.g_unix_socket_address_get_path_len((*C.GUnixSocketAddress)(recv.native))
	retGo := (uint64)(retC)

	return retGo
}

type signalVolumeMonitorDriveStopButtonDetail struct {
	callback  VolumeMonitorSignalDriveStopButtonCallback
	handlerID C.gulong
}

var signalVolumeMonitorDriveStopButtonId int
var signalVolumeMonitorDriveStopButtonMap = make(map[int]signalVolumeMonitorDriveStopButtonDetail)
var signalVolumeMonitorDriveStopButtonLock sync.RWMutex

// VolumeMonitorSignalDriveStopButtonCallback is a callback function for a 'drive-stop-button' signal emitted from a VolumeMonitor.
type VolumeMonitorSignalDriveStopButtonCallback func(drive *Drive)

/*
ConnectDriveStopButton connects the callback to the 'drive-stop-button' signal for the VolumeMonitor.

The returned value represents the connection, and may be passed to DisconnectDriveStopButton to remove it.
*/
func (recv *VolumeMonitor) ConnectDriveStopButton(callback VolumeMonitorSignalDriveStopButtonCallback) int {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	signalVolumeMonitorDriveStopButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.VolumeMonitor_signal_connect_drive_stop_button(instance, C.gpointer(uintptr(signalVolumeMonitorDriveStopButtonId)))

	detail := signalVolumeMonitorDriveStopButtonDetail{callback, handlerID}
	signalVolumeMonitorDriveStopButtonMap[signalVolumeMonitorDriveStopButtonId] = detail

	return signalVolumeMonitorDriveStopButtonId
}

/*
DisconnectDriveStopButton disconnects a callback from the 'drive-stop-button' signal for the VolumeMonitor.

The connectionID should be a value returned from a call to ConnectDriveStopButton.
*/
func (recv *VolumeMonitor) DisconnectDriveStopButton(connectionID int) {
	signalVolumeMonitorDriveStopButtonLock.Lock()
	defer signalVolumeMonitorDriveStopButtonLock.Unlock()

	detail, exists := signalVolumeMonitorDriveStopButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalVolumeMonitorDriveStopButtonMap, connectionID)
}

//export volumemonitor_driveStopButtonHandler
func volumemonitor_driveStopButtonHandler(_ *C.GObject, c_drive *C.GDrive, data C.gpointer) {
	signalVolumeMonitorDriveStopButtonLock.RLock()
	defer signalVolumeMonitorDriveStopButtonLock.RUnlock()

	drive := DriveNewFromC(unsafe.Pointer(c_drive))

	index := int(uintptr(data))
	callback := signalVolumeMonitorDriveStopButtonMap[index].callback
	callback(drive)
}

const FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_POLL
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START
const FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_START_DEGRADED
const FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP string = C.G_FILE_ATTRIBUTE_MOUNTABLE_CAN_STOP
const FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC string = C.G_FILE_ATTRIBUTE_MOUNTABLE_IS_MEDIA_CHECK_AUTOMATIC
const FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_START_STOP_TYPE
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE string = C.G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE_FILE

type DriveStartStopType C.GDriveStartStopType

const (
	DRIVE_START_STOP_TYPE_UNKNOWN   DriveStartStopType = 0
	DRIVE_START_STOP_TYPE_SHUTDOWN  DriveStartStopType = 1
	DRIVE_START_STOP_TYPE_NETWORK   DriveStartStopType = 2
	DRIVE_START_STOP_TYPE_MULTIDISK DriveStartStopType = 3
	DRIVE_START_STOP_TYPE_PASSWORD  DriveStartStopType = 4
)

type ResolverError C.GResolverError

const (
	RESOLVER_ERROR_NOT_FOUND         ResolverError = 0
	RESOLVER_ERROR_TEMPORARY_FAILURE ResolverError = 1
	RESOLVER_ERROR_INTERNAL          ResolverError = 2
)

// ResolverErrorQuark is a wrapper around the C function g_resolver_error_quark.
func ResolverErrorQuark() glib.Quark {
	retC := C.g_resolver_error_quark()
	retGo := (glib.Quark)(retC)

	return retGo
}

type SocketFamily C.GSocketFamily

const (
	SOCKET_FAMILY_INVALID SocketFamily = 0
	SOCKET_FAMILY_UNIX    SocketFamily = 1
	SOCKET_FAMILY_IPV4    SocketFamily = 2
	SOCKET_FAMILY_IPV6    SocketFamily = 10
)

type SocketProtocol C.GSocketProtocol

const (
	SOCKET_PROTOCOL_UNKNOWN SocketProtocol = -1
	SOCKET_PROTOCOL_DEFAULT SocketProtocol = 0
	SOCKET_PROTOCOL_TCP     SocketProtocol = 6
	SOCKET_PROTOCOL_UDP     SocketProtocol = 17
	SOCKET_PROTOCOL_SCTP    SocketProtocol = 132
)

type SocketType C.GSocketType

const (
	SOCKET_TYPE_INVALID   SocketType = 0
	SOCKET_TYPE_STREAM    SocketType = 1
	SOCKET_TYPE_DATAGRAM  SocketType = 2
	SOCKET_TYPE_SEQPACKET SocketType = 3
)

// Unsupported : g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// Unsupported : g_initable_newv : unsupported parameter parameters :

// AsyncInitable is a wrapper around the C record GAsyncInitable.
type AsyncInitable struct {
	native *C.GAsyncInitable
}

func AsyncInitableNewFromC(u unsafe.Pointer) *AsyncInitable {
	c := (*C.GAsyncInitable)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitable{native: c}

	return g
}

func (recv *AsyncInitable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncInitable with another AsyncInitable, and returns true if they represent the same GObject.
func (recv *AsyncInitable) Equals(other *AsyncInitable) bool {
	return other.ToC() == recv.ToC()
}

// g_async_initable_new_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// g_async_initable_new_valist_async : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// g_async_initable_newv_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback
// Unsupported : g_async_initable_init_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// InitFinish is a wrapper around the C function g_async_initable_init_finish.
func (recv *AsyncInitable) InitFinish(res *AsyncResult) (bool, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_init_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// NewFinish is a wrapper around the C function g_async_initable_new_finish.
func (recv *AsyncInitable) NewFinish(res *AsyncResult) (*gobject.Object, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_async_initable_new_finish((*C.GAsyncInitable)(recv.native), c_res, &cThrowableError)
	retGo := gobject.ObjectNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

type signalDriveStopButtonDetail struct {
	callback  DriveSignalStopButtonCallback
	handlerID C.gulong
}

var signalDriveStopButtonId int
var signalDriveStopButtonMap = make(map[int]signalDriveStopButtonDetail)
var signalDriveStopButtonLock sync.RWMutex

// DriveSignalStopButtonCallback is a callback function for a 'stop-button' signal emitted from a Drive.
type DriveSignalStopButtonCallback func()

/*
ConnectStopButton connects the callback to the 'stop-button' signal for the Drive.

The returned value represents the connection, and may be passed to DisconnectStopButton to remove it.
*/
func (recv *Drive) ConnectStopButton(callback DriveSignalStopButtonCallback) int {
	signalDriveStopButtonLock.Lock()
	defer signalDriveStopButtonLock.Unlock()

	signalDriveStopButtonId++
	instance := C.gpointer(recv.native)
	handlerID := C.Drive_signal_connect_stop_button(instance, C.gpointer(uintptr(signalDriveStopButtonId)))

	detail := signalDriveStopButtonDetail{callback, handlerID}
	signalDriveStopButtonMap[signalDriveStopButtonId] = detail

	return signalDriveStopButtonId
}

/*
DisconnectStopButton disconnects a callback from the 'stop-button' signal for the Drive.

The connectionID should be a value returned from a call to ConnectStopButton.
*/
func (recv *Drive) DisconnectStopButton(connectionID int) {
	signalDriveStopButtonLock.Lock()
	defer signalDriveStopButtonLock.Unlock()

	detail, exists := signalDriveStopButtonMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalDriveStopButtonMap, connectionID)
}

//export drive_stopButtonHandler
func drive_stopButtonHandler(_ *C.GObject, data C.gpointer) {
	signalDriveStopButtonLock.RLock()
	defer signalDriveStopButtonLock.RUnlock()

	index := int(uintptr(data))
	callback := signalDriveStopButtonMap[index].callback
	callback()
}

// CanStart is a wrapper around the C function g_drive_can_start.
func (recv *Drive) CanStart() bool {
	retC := C.g_drive_can_start((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanStartDegraded is a wrapper around the C function g_drive_can_start_degraded.
func (recv *Drive) CanStartDegraded() bool {
	retC := C.g_drive_can_start_degraded((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// CanStop is a wrapper around the C function g_drive_can_stop.
func (recv *Drive) CanStop() bool {
	retC := C.g_drive_can_stop((*C.GDrive)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_drive_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_drive_eject_with_operation_finish.
func (recv *Drive) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_eject_with_operation_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetStartStopType is a wrapper around the C function g_drive_get_start_stop_type.
func (recv *Drive) GetStartStopType() DriveStartStopType {
	retC := C.g_drive_get_start_stop_type((*C.GDrive)(recv.native))
	retGo := (DriveStartStopType)(retC)

	return retGo
}

// Unsupported : g_drive_start : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StartFinish is a wrapper around the C function g_drive_start_finish.
func (recv *Drive) StartFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_start_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_drive_stop : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StopFinish is a wrapper around the C function g_drive_stop_finish.
func (recv *Drive) StopFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_drive_stop_finish((*C.GDrive)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// CreateReadwrite is a wrapper around the C function g_file_create_readwrite.
func (recv *File) CreateReadwrite(flags FileCreateFlags, cancellable *Cancellable) (*FileIOStream, error) {
	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite((*C.GFile)(recv.native), c_flags, c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_create_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// CreateReadwriteFinish is a wrapper around the C function g_file_create_readwrite_finish.
func (recv *File) CreateReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_create_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_eject_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectMountableWithOperationFinish is a wrapper around the C function g_file_eject_mountable_with_operation_finish.
func (recv *File) EjectMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_eject_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// OpenReadwrite is a wrapper around the C function g_file_open_readwrite.
func (recv *File) OpenReadwrite(cancellable *Cancellable) (*FileIOStream, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite((*C.GFile)(recv.native), c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_open_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// OpenReadwriteFinish is a wrapper around the C function g_file_open_readwrite_finish.
func (recv *File) OpenReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_open_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_poll_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// PollMountableFinish is a wrapper around the C function g_file_poll_mountable_finish.
func (recv *File) PollMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_poll_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// ReplaceReadwrite is a wrapper around the C function g_file_replace_readwrite.
func (recv *File) ReplaceReadwrite(etag string, makeBackup bool, flags FileCreateFlags, cancellable *Cancellable) (*FileIOStream, error) {
	c_etag := C.CString(etag)
	defer C.free(unsafe.Pointer(c_etag))

	c_make_backup :=
		boolToGboolean(makeBackup)

	c_flags := (C.GFileCreateFlags)(flags)

	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_file_replace_readwrite((*C.GFile)(recv.native), c_etag, c_make_backup, c_flags, c_cancellable, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_replace_readwrite_async : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// ReplaceReadwriteFinish is a wrapper around the C function g_file_replace_readwrite_finish.
func (recv *File) ReplaceReadwriteFinish(res *AsyncResult) (*FileIOStream, error) {
	c_res := (*C.GAsyncResult)(res.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_replace_readwrite_finish((*C.GFile)(recv.native), c_res, &cThrowableError)
	retGo := FileIOStreamNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_start_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StartMountableFinish is a wrapper around the C function g_file_start_mountable_finish.
func (recv *File) StartMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_start_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_file_stop_mountable : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// StopMountableFinish is a wrapper around the C function g_file_stop_mountable_finish.
func (recv *File) StopMountableFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_stop_mountable_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SupportsThreadContexts is a wrapper around the C function g_file_supports_thread_contexts.
func (recv *File) SupportsThreadContexts() bool {
	retC := C.g_file_supports_thread_contexts((*C.GFile)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Unsupported : g_file_unmount_mountable_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountMountableWithOperationFinish is a wrapper around the C function g_file_unmount_mountable_with_operation_finish.
func (recv *File) UnmountMountableWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_file_unmount_mountable_with_operation_finish((*C.GFile)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Initable is a wrapper around the C record GInitable.
type Initable struct {
	native *C.GInitable
}

func InitableNewFromC(u unsafe.Pointer) *Initable {
	c := (*C.GInitable)(u)
	if c == nil {
		return nil
	}

	g := &Initable{native: c}

	return g
}

func (recv *Initable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this Initable with another Initable, and returns true if they represent the same GObject.
func (recv *Initable) Equals(other *Initable) bool {
	return other.ToC() == recv.ToC()
}

// g_initable_new : unsupported parameter ... : varargs
// g_initable_new_valist : unsupported parameter var_args : no type generator for va_list (va_list) for param var_args
// g_initable_newv : unsupported parameter parameters :
// Init is a wrapper around the C function g_initable_init.
func (recv *Initable) Init(cancellable *Cancellable) (bool, error) {
	c_cancellable := (*C.GCancellable)(C.NULL)
	if cancellable != nil {
		c_cancellable = (*C.GCancellable)(cancellable.ToC())
	}

	var cThrowableError *C.GError

	retC := C.g_initable_init((*C.GInitable)(recv.native), c_cancellable, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

type signalMountPreUnmountDetail struct {
	callback  MountSignalPreUnmountCallback
	handlerID C.gulong
}

var signalMountPreUnmountId int
var signalMountPreUnmountMap = make(map[int]signalMountPreUnmountDetail)
var signalMountPreUnmountLock sync.RWMutex

// MountSignalPreUnmountCallback is a callback function for a 'pre-unmount' signal emitted from a Mount.
type MountSignalPreUnmountCallback func()

/*
ConnectPreUnmount connects the callback to the 'pre-unmount' signal for the Mount.

The returned value represents the connection, and may be passed to DisconnectPreUnmount to remove it.
*/
func (recv *Mount) ConnectPreUnmount(callback MountSignalPreUnmountCallback) int {
	signalMountPreUnmountLock.Lock()
	defer signalMountPreUnmountLock.Unlock()

	signalMountPreUnmountId++
	instance := C.gpointer(recv.native)
	handlerID := C.Mount_signal_connect_pre_unmount(instance, C.gpointer(uintptr(signalMountPreUnmountId)))

	detail := signalMountPreUnmountDetail{callback, handlerID}
	signalMountPreUnmountMap[signalMountPreUnmountId] = detail

	return signalMountPreUnmountId
}

/*
DisconnectPreUnmount disconnects a callback from the 'pre-unmount' signal for the Mount.

The connectionID should be a value returned from a call to ConnectPreUnmount.
*/
func (recv *Mount) DisconnectPreUnmount(connectionID int) {
	signalMountPreUnmountLock.Lock()
	defer signalMountPreUnmountLock.Unlock()

	detail, exists := signalMountPreUnmountMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalMountPreUnmountMap, connectionID)
}

//export mount_preUnmountHandler
func mount_preUnmountHandler(_ *C.GObject, data C.gpointer) {
	signalMountPreUnmountLock.RLock()
	defer signalMountPreUnmountLock.RUnlock()

	index := int(uintptr(data))
	callback := signalMountPreUnmountMap[index].callback
	callback()
}

// Unsupported : g_mount_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_mount_eject_with_operation_finish.
func (recv *Mount) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_eject_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_mount_unmount_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// UnmountWithOperationFinish is a wrapper around the C function g_mount_unmount_with_operation_finish.
func (recv *Mount) UnmountWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_mount_unmount_with_operation_finish((*C.GMount)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Enumerate is a wrapper around the C function g_socket_connectable_enumerate.
func (recv *SocketConnectable) Enumerate() *SocketAddressEnumerator {
	retC := C.g_socket_connectable_enumerate((*C.GSocketConnectable)(recv.native))
	retGo := SocketAddressEnumeratorNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_volume_eject_with_operation : unsupported parameter callback : no type generator for AsyncReadyCallback (GAsyncReadyCallback) for param callback

// EjectWithOperationFinish is a wrapper around the C function g_volume_eject_with_operation_finish.
func (recv *Volume) EjectWithOperationFinish(result *AsyncResult) (bool, error) {
	c_result := (*C.GAsyncResult)(result.ToC())

	var cThrowableError *C.GError

	retC := C.g_volume_eject_with_operation_finish((*C.GVolume)(recv.native), c_result, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AsyncInitableIface is a wrapper around the C record GAsyncInitableIface.
type AsyncInitableIface struct {
	native *C.GAsyncInitableIface
	// g_iface : record
	// no type for init_async
	// no type for init_finish
}

func AsyncInitableIfaceNewFromC(u unsafe.Pointer) *AsyncInitableIface {
	c := (*C.GAsyncInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &AsyncInitableIface{native: c}

	return g
}

func (recv *AsyncInitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this AsyncInitableIface with another AsyncInitableIface, and returns true if they represent the same GObject.
func (recv *AsyncInitableIface) Equals(other *AsyncInitableIface) bool {
	return other.ToC() == recv.ToC()
}

// InitableIface is a wrapper around the C record GInitableIface.
type InitableIface struct {
	native *C.GInitableIface
	// g_iface : record
	// no type for init
}

func InitableIfaceNewFromC(u unsafe.Pointer) *InitableIface {
	c := (*C.GInitableIface)(u)
	if c == nil {
		return nil
	}

	g := &InitableIface{native: c}

	return g
}

func (recv *InitableIface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InitableIface with another InitableIface, and returns true if they represent the same GObject.
func (recv *InitableIface) Equals(other *InitableIface) bool {
	return other.ToC() == recv.ToC()
}

// InputVector is a wrapper around the C record GInputVector.
type InputVector struct {
	native *C.GInputVector
	// buffer : no type generator for gpointer, gpointer
	Size uint64
}

func InputVectorNewFromC(u unsafe.Pointer) *InputVector {
	c := (*C.GInputVector)(u)
	if c == nil {
		return nil
	}

	g := &InputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *InputVector) ToC() unsafe.Pointer {
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this InputVector with another InputVector, and returns true if they represent the same GObject.
func (recv *InputVector) Equals(other *InputVector) bool {
	return other.ToC() == recv.ToC()
}

// OutputVector is a wrapper around the C record GOutputVector.
type OutputVector struct {
	native *C.GOutputVector
	// buffer : no type generator for gpointer, gconstpointer
	Size uint64
}

func OutputVectorNewFromC(u unsafe.Pointer) *OutputVector {
	c := (*C.GOutputVector)(u)
	if c == nil {
		return nil
	}

	g := &OutputVector{
		Size:   (uint64)(c.size),
		native: c,
	}

	return g
}

func (recv *OutputVector) ToC() unsafe.Pointer {
	recv.native.size =
		(C.gsize)(recv.Size)

	return (unsafe.Pointer)(recv.native)
}

// Equals compares this OutputVector with another OutputVector, and returns true if they represent the same GObject.
func (recv *OutputVector) Equals(other *OutputVector) bool {
	return other.ToC() == recv.ToC()
}

// SrvTargetNew is a wrapper around the C function g_srv_target_new.
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	c_hostname := C.CString(hostname)
	defer C.free(unsafe.Pointer(c_hostname))

	c_port := (C.guint16)(port)

	c_priority := (C.guint16)(priority)

	c_weight := (C.guint16)(weight)

	retC := C.g_srv_target_new(c_hostname, c_port, c_priority, c_weight)
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SrvTargetListSort is a wrapper around the C function g_srv_target_list_sort.
func SrvTargetListSort(targets *glib.List) *glib.List {
	c_targets := (*C.GList)(C.NULL)
	if targets != nil {
		c_targets = (*C.GList)(targets.ToC())
	}

	retC := C.g_srv_target_list_sort(c_targets)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Copy is a wrapper around the C function g_srv_target_copy.
func (recv *SrvTarget) Copy() *SrvTarget {
	retC := C.g_srv_target_copy((*C.GSrvTarget)(recv.native))
	retGo := SrvTargetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_srv_target_free.
func (recv *SrvTarget) Free() {
	C.g_srv_target_free((*C.GSrvTarget)(recv.native))

	return
}

// GetHostname is a wrapper around the C function g_srv_target_get_hostname.
func (recv *SrvTarget) GetHostname() string {
	retC := C.g_srv_target_get_hostname((*C.GSrvTarget)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetPort is a wrapper around the C function g_srv_target_get_port.
func (recv *SrvTarget) GetPort() uint16 {
	retC := C.g_srv_target_get_port((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetPriority is a wrapper around the C function g_srv_target_get_priority.
func (recv *SrvTarget) GetPriority() uint16 {
	retC := C.g_srv_target_get_priority((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetWeight is a wrapper around the C function g_srv_target_get_weight.
func (recv *SrvTarget) GetWeight() uint16 {
	retC := C.g_srv_target_get_weight((*C.GSrvTarget)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}