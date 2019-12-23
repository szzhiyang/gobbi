// Code generated - DO NOT EDIT.
// +build gio_2.24

package gio

import (
	gobject "github.com/pekim/gobbi/lib/internal/c/gobject"
	"unsafe"
)

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
import "C"

type ActionEntry C.GActionEntry
type AppInfoIface C.GAppInfoIface
type AppLaunchContextClass C.GAppLaunchContextClass
type AppLaunchContextPrivate C.GAppLaunchContextPrivate
type ApplicationCommandLinePrivate C.GApplicationCommandLinePrivate
type ApplicationPrivate C.GApplicationPrivate
type AsyncInitableIface C.GAsyncInitableIface
type AsyncResultIface C.GAsyncResultIface
type BufferedInputStreamClass C.GBufferedInputStreamClass
type BufferedInputStreamPrivate C.GBufferedInputStreamPrivate
type BufferedOutputStreamClass C.GBufferedOutputStreamClass
type BufferedOutputStreamPrivate C.GBufferedOutputStreamPrivate
type CancellableClass C.GCancellableClass
type CancellablePrivate C.GCancellablePrivate
type CharsetConverterClass C.GCharsetConverterClass
type ConverterIface C.GConverterIface
type ConverterInputStreamClass C.GConverterInputStreamClass
type ConverterInputStreamPrivate C.GConverterInputStreamPrivate
type ConverterOutputStreamClass C.GConverterOutputStreamClass
type ConverterOutputStreamPrivate C.GConverterOutputStreamPrivate
type DBusInterfaceSkeletonPrivate C.GDBusInterfaceSkeletonPrivate
type DBusObjectManagerClientPrivate C.GDBusObjectManagerClientPrivate
type DBusObjectManagerServerPrivate C.GDBusObjectManagerServerPrivate
type DBusObjectProxyPrivate C.GDBusObjectProxyPrivate
type DBusObjectSkeletonPrivate C.GDBusObjectSkeletonPrivate
type DBusProxyPrivate C.GDBusProxyPrivate
type DataInputStreamClass C.GDataInputStreamClass
type DataInputStreamPrivate C.GDataInputStreamPrivate
type DataOutputStreamClass C.GDataOutputStreamClass
type DataOutputStreamPrivate C.GDataOutputStreamPrivate
type DesktopAppInfoClass C.GDesktopAppInfoClass
type DesktopAppInfoLookupIface C.GDesktopAppInfoLookupIface
type DriveIface C.GDriveIface
type EmblemClass C.GEmblemClass
type EmblemedIconClass C.GEmblemedIconClass
type EmblemedIconPrivate C.GEmblemedIconPrivate
type FileAttributeInfo C.GFileAttributeInfo
type FileAttributeInfoList C.GFileAttributeInfoList
type FileAttributeMatcher C.GFileAttributeMatcher
type FileDescriptorBasedIface C.GFileDescriptorBasedIface
type FileEnumeratorClass C.GFileEnumeratorClass
type FileEnumeratorPrivate C.GFileEnumeratorPrivate
type FileIOStreamClass C.GFileIOStreamClass
type FileIOStreamPrivate C.GFileIOStreamPrivate
type FileIconClass C.GFileIconClass
type FileIface C.GFileIface
type FileInfoClass C.GFileInfoClass
type FileInputStreamClass C.GFileInputStreamClass
type FileInputStreamPrivate C.GFileInputStreamPrivate
type FileMonitorClass C.GFileMonitorClass
type FileMonitorPrivate C.GFileMonitorPrivate
type FileOutputStreamClass C.GFileOutputStreamClass
type FileOutputStreamPrivate C.GFileOutputStreamPrivate
type FilenameCompleterClass C.GFilenameCompleterClass
type FilterInputStreamClass C.GFilterInputStreamClass
type FilterOutputStreamClass C.GFilterOutputStreamClass
type IOExtension C.GIOExtension
type IOExtensionPoint C.GIOExtensionPoint
type IOModuleClass C.GIOModuleClass
type IOSchedulerJob C.GIOSchedulerJob
type IOStreamAdapter C.GIOStreamAdapter
type IOStreamClass C.GIOStreamClass
type IOStreamPrivate C.GIOStreamPrivate
type IconIface C.GIconIface
type InetAddressClass C.GInetAddressClass
type InetAddressMaskClass C.GInetAddressMaskClass
type InetAddressMaskPrivate C.GInetAddressMaskPrivate
type InetAddressPrivate C.GInetAddressPrivate
type InetSocketAddressClass C.GInetSocketAddressClass
type InetSocketAddressPrivate C.GInetSocketAddressPrivate
type InitableIface C.GInitableIface
type InputStreamClass C.GInputStreamClass
type InputStreamPrivate C.GInputStreamPrivate
type InputVector C.GInputVector
type ListStoreClass C.GListStoreClass
type LoadableIconIface C.GLoadableIconIface
type MemoryInputStreamClass C.GMemoryInputStreamClass
type MemoryInputStreamPrivate C.GMemoryInputStreamPrivate
type MemoryOutputStreamClass C.GMemoryOutputStreamClass
type MemoryOutputStreamPrivate C.GMemoryOutputStreamPrivate
type MenuAttributeIterClass C.GMenuAttributeIterClass
type MenuAttributeIterPrivate C.GMenuAttributeIterPrivate
type MenuLinkIterClass C.GMenuLinkIterClass
type MenuLinkIterPrivate C.GMenuLinkIterPrivate
type MenuModelClass C.GMenuModelClass
type MenuModelPrivate C.GMenuModelPrivate
type MountIface C.GMountIface
type MountOperationClass C.GMountOperationClass
type MountOperationPrivate C.GMountOperationPrivate

// UNSUPPORTED : NativeSocketAddressClass : blacklisted
// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted
type NativeVolumeMonitorClass C.GNativeVolumeMonitorClass
type NetworkAddressClass C.GNetworkAddressClass
type NetworkAddressPrivate C.GNetworkAddressPrivate
type NetworkServiceClass C.GNetworkServiceClass
type NetworkServicePrivate C.GNetworkServicePrivate
type OutputStreamClass C.GOutputStreamClass
type OutputStreamPrivate C.GOutputStreamPrivate
type OutputVector C.GOutputVector
type PermissionClass C.GPermissionClass
type PermissionPrivate C.GPermissionPrivate
type ProxyAddressEnumeratorClass C.GProxyAddressEnumeratorClass
type ProxyAddressEnumeratorPrivate C.GProxyAddressEnumeratorPrivate
type ProxyAddressPrivate C.GProxyAddressPrivate
type ProxyResolverInterface C.GProxyResolverInterface
type ResolverClass C.GResolverClass
type ResolverPrivate C.GResolverPrivate
type SeekableIface C.GSeekableIface

// UNSUPPORTED : SettingsBackendClass : blacklisted
// UNSUPPORTED : SettingsBackendPrivate : blacklisted
type SettingsClass C.GSettingsClass
type SettingsPrivate C.GSettingsPrivate
type SettingsSchemaKey C.GSettingsSchemaKey
type SimpleActionGroupClass C.GSimpleActionGroupClass
type SimpleActionGroupPrivate C.GSimpleActionGroupPrivate
type SimpleAsyncResultClass C.GSimpleAsyncResultClass
type SimpleProxyResolverClass C.GSimpleProxyResolverClass
type SimpleProxyResolverPrivate C.GSimpleProxyResolverPrivate
type SocketAddressClass C.GSocketAddressClass
type SocketAddressEnumeratorClass C.GSocketAddressEnumeratorClass
type SocketClass C.GSocketClass
type SocketClientClass C.GSocketClientClass
type SocketClientPrivate C.GSocketClientPrivate
type SocketConnectableIface C.GSocketConnectableIface
type SocketConnectionClass C.GSocketConnectionClass
type SocketConnectionPrivate C.GSocketConnectionPrivate
type SocketControlMessageClass C.GSocketControlMessageClass
type SocketControlMessagePrivate C.GSocketControlMessagePrivate
type SocketListenerClass C.GSocketListenerClass
type SocketListenerPrivate C.GSocketListenerPrivate
type SocketPrivate C.GSocketPrivate
type SocketServiceClass C.GSocketServiceClass
type SocketServicePrivate C.GSocketServicePrivate
type SrvTarget C.GSrvTarget
type StaticResource C.GStaticResource
type TaskClass C.GTaskClass
type TcpConnectionClass C.GTcpConnectionClass
type TcpConnectionPrivate C.GTcpConnectionPrivate
type TcpWrapperConnectionClass C.GTcpWrapperConnectionClass
type TcpWrapperConnectionPrivate C.GTcpWrapperConnectionPrivate
type ThemedIconClass C.GThemedIconClass
type ThreadedSocketServiceClass C.GThreadedSocketServiceClass
type ThreadedSocketServicePrivate C.GThreadedSocketServicePrivate
type TlsCertificateClass C.GTlsCertificateClass
type TlsCertificatePrivate C.GTlsCertificatePrivate
type TlsConnectionClass C.GTlsConnectionClass
type TlsConnectionPrivate C.GTlsConnectionPrivate
type TlsDatabasePrivate C.GTlsDatabasePrivate
type TlsFileDatabaseInterface C.GTlsFileDatabaseInterface
type TlsInteractionPrivate C.GTlsInteractionPrivate
type TlsPasswordClass C.GTlsPasswordClass
type TlsPasswordPrivate C.GTlsPasswordPrivate
type UnixConnectionClass C.GUnixConnectionClass
type UnixConnectionPrivate C.GUnixConnectionPrivate
type UnixCredentialsMessagePrivate C.GUnixCredentialsMessagePrivate
type UnixFDListClass C.GUnixFDListClass
type UnixFDListPrivate C.GUnixFDListPrivate
type UnixFDMessageClass C.GUnixFDMessageClass
type UnixFDMessagePrivate C.GUnixFDMessagePrivate
type UnixInputStreamClass C.GUnixInputStreamClass
type UnixInputStreamPrivate C.GUnixInputStreamPrivate
type UnixMountEntry C.GUnixMountEntry
type UnixMountMonitorClass C.GUnixMountMonitorClass
type UnixMountPoint C.GUnixMountPoint
type UnixOutputStreamClass C.GUnixOutputStreamClass
type UnixOutputStreamPrivate C.GUnixOutputStreamPrivate
type UnixSocketAddressClass C.GUnixSocketAddressClass
type UnixSocketAddressPrivate C.GUnixSocketAddressPrivate
type VfsClass C.GVfsClass
type VolumeIface C.GVolumeIface
type VolumeMonitorClass C.GVolumeMonitorClass
type ZlibCompressorClass C.GZlibCompressorClass
type ZlibDecompressorClass C.GZlibDecompressorClass

func Fn_g_app_info_create_from_commandline(param0 string, param1 string, param2 int) {}

func Fn_g_app_info_get_all() {
	C.g_app_info_get_all()
}

func Fn_g_app_info_get_all_for_type(param0 string) {}

func Fn_g_app_info_get_default_for_type(param0 string, param1 bool) {}

func Fn_g_app_info_get_default_for_uri_scheme(param0 string) {}

func Fn_g_app_info_launch_default_for_uri(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : app_info_launch_default_for_uri_async : has callback

func Fn_g_app_info_reset_type_associations(param0 string) {}

// UNSUPPORTED : async_initable_newv_async : has callback

// UNSUPPORTED : bus_get : has callback

// UNSUPPORTED : bus_own_name : has callback

// UNSUPPORTED : bus_own_name_on_connection : has callback

// UNSUPPORTED : bus_watch_name : has callback

// UNSUPPORTED : bus_watch_name_on_connection : has callback

func Fn_g_content_type_can_be_executable(param0 string) {}

func Fn_g_content_type_equals(param0 string, param1 string) {}

func Fn_g_content_type_from_mime_type(param0 string) {}

func Fn_g_content_type_get_description(param0 string) {}

func Fn_g_content_type_get_icon(param0 string) {}

func Fn_g_content_type_get_mime_type(param0 string) {}

func Fn_g_content_type_guess(param0 string, param1 []uint8, param2 uint64, param3 *bool) {}

func Fn_g_content_type_guess_for_tree(param0 unsafe.Pointer) {}

func Fn_g_content_type_is_a(param0 string, param1 string) {}

func Fn_g_content_type_is_unknown(param0 string) {}

func Fn_g_content_types_get_registered() {
	C.g_content_types_get_registered()
}

// UNSUPPORTED : dbus_address_get_stream : has callback

func Fn_g_dbus_error_quark() {
	C.g_dbus_error_quark()
}

func Fn_g_file_new_for_commandline_arg(param0 string) {}

func Fn_g_file_new_for_path(param0 string) {}

func Fn_g_file_new_for_uri(param0 string) {}

func Fn_g_file_parse_name(param0 string) {}

func Fn_g_icon_hash(param0 unsafe.Pointer) {}

func Fn_g_icon_new_for_string(param0 string) {}

func Fn_g_initable_newv(param0 uint64, param1 uint, param2 []gobject.Parameter, param3 unsafe.Pointer) {
}

func Fn_g_io_error_from_errno(param0 int) {}

func Fn_g_io_error_quark() {
	C.g_io_error_quark()
}

func Fn_g_io_extension_point_implement(param0 string, param1 uint64, param2 string, param3 int) {}

func Fn_g_io_extension_point_lookup(param0 string) {}

func Fn_g_io_extension_point_register(param0 string) {}

func Fn_g_io_modules_load_all_in_directory(param0 string) {}

func Fn_g_io_modules_scan_all_in_directory(param0 string) {}

func Fn_g_io_scheduler_cancel_all_jobs() {
	C.g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : io_scheduler_push_job : has callback

// UNSUPPORTED : keyfile_settings_backend_new : blacklisted
// UNSUPPORTED : memory_settings_backend_new : blacklisted
// UNSUPPORTED : null_settings_backend_new : blacklisted
func Fn_g_resolver_error_quark() {
	C.g_resolver_error_quark()
}

// UNSUPPORTED : simple_async_report_error_in_idle : has varargs

// UNSUPPORTED : simple_async_report_gerror_in_idle : has callback

// UNSUPPORTED : simple_async_report_take_gerror_in_idle : has callback

func Fn_g_srv_target_list_sort(param0 unsafe.Pointer) {}

func Fn_g_unix_is_mount_path_system_internal(param0 string) {}

func Fn_g_unix_mount_at(param0 string, param1 *uint64) {}

func Fn_g_unix_mount_compare(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_unix_mount_free(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_device_path(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_fs_type(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_get_mount_path(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_can_eject(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_icon(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_name(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_guess_should_display(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_is_readonly(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_is_system_internal(param0 unsafe.Pointer) {}

func Fn_g_unix_mount_points_changed_since(param0 uint64) {}

func Fn_g_unix_mount_points_get(param0 *uint64) {}

func Fn_g_unix_mounts_changed_since(param0 uint64) {}

func Fn_g_unix_mounts_get(param0 *uint64) {}

func Fn_g_app_launch_context_new() {
	C.g_app_launch_context_new()
}

func Fn_g_app_launch_context_get_display(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_app_launch_context_get_startup_notify_id(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_app_launch_context_launch_failed(param0 string) {}

func Fn_g_application_new(param0 string, param1 int) {}

func Fn_g_application_hold() {}

func Fn_g_application_release() {}

func Fn_g_application_id_is_valid(param0 string) {}

// UNSUPPORTED : print : has varargs

// UNSUPPORTED : printerr : has varargs

func Fn_g_buffered_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_input_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_buffered_input_stream_fill(param0 uint64, param1 unsafe.Pointer) {}

// UNSUPPORTED : fill_async : has callback

func Fn_g_buffered_input_stream_fill_finish(param0 unsafe.Pointer) {}

func Fn_g_buffered_input_stream_get_available() {}

func Fn_g_buffered_input_stream_get_buffer_size() {}

func Fn_g_buffered_input_stream_peek(param0 []uint8, param1 uint64, param2 uint64) {}

func Fn_g_buffered_input_stream_peek_buffer(param0 *uint64) {}

func Fn_g_buffered_input_stream_read_byte(param0 unsafe.Pointer) {}

func Fn_g_buffered_input_stream_set_buffer_size(param0 uint64) {}

func Fn_g_buffered_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_buffered_output_stream_new_sized(param0 unsafe.Pointer, param1 uint64) {}

func Fn_g_buffered_output_stream_get_auto_grow() {}

func Fn_g_buffered_output_stream_get_buffer_size() {}

func Fn_g_buffered_output_stream_set_auto_grow(param0 bool) {}

func Fn_g_buffered_output_stream_set_buffer_size(param0 uint64) {}

func Fn_g_cancellable_new() {
	C.g_cancellable_new()
}

func Fn_g_cancellable_cancel() {}

// UNSUPPORTED : connect : has callback

func Fn_g_cancellable_disconnect(param0 uint64) {}

func Fn_g_cancellable_get_fd() {}

func Fn_g_cancellable_is_cancelled() {}

func Fn_g_cancellable_make_pollfd(param0 unsafe.Pointer) {}

func Fn_g_cancellable_pop_current() {}

func Fn_g_cancellable_push_current() {}

func Fn_g_cancellable_release_fd() {}

func Fn_g_cancellable_reset() {}

func Fn_g_cancellable_set_error_if_cancelled() {}

func Fn_g_cancellable_get_current() {
	C.g_cancellable_get_current()
}

func Fn_g_charset_converter_new(param0 string, param1 string) {}

func Fn_g_charset_converter_get_num_fallbacks() {}

func Fn_g_charset_converter_get_use_fallback() {}

func Fn_g_charset_converter_set_use_fallback(param0 bool) {}

func Fn_g_converter_input_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_converter_input_stream_get_converter() {}

func Fn_g_converter_output_stream_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_converter_output_stream_get_converter() {}

// UNSUPPORTED : add_filter : has callback

// UNSUPPORTED : call : has callback

// UNSUPPORTED : call_with_unix_fd_list : has callback

// UNSUPPORTED : close : has callback

// UNSUPPORTED : flush : has callback

// UNSUPPORTED : register_object : has callback

// UNSUPPORTED : register_subtree : has callback

// UNSUPPORTED : send_message_with_reply : has callback

// UNSUPPORTED : signal_subscribe : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_address : has callback

func Fn_g_dbus_message_get_byte_order() {}

// UNSUPPORTED : new_method_error : has varargs

// UNSUPPORTED : new_method_error_valist : has va_list

func Fn_g_dbus_message_set_byte_order(param0 int) {}

// UNSUPPORTED : return_error : has varargs

// UNSUPPORTED : return_error_valist : has va_list

// UNSUPPORTED : new_for_bus_sync : has callback

// UNSUPPORTED : new_sync : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_dbus_object_manager_server_set_connection(param0 unsafe.Pointer) {}

// UNSUPPORTED : call : has callback

// UNSUPPORTED : call_with_unix_fd_list : has callback

// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_for_bus : has callback

func Fn_g_data_input_stream_new(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_get_byte_order() {}

func Fn_g_data_input_stream_get_newline_type() {}

func Fn_g_data_input_stream_read_byte(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_int16(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_int32(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_int64(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_line(param0 *uint64, param1 unsafe.Pointer) {}

// UNSUPPORTED : read_line_async : has callback

func Fn_g_data_input_stream_read_line_finish(param0 unsafe.Pointer, param1 *uint64) {}

func Fn_g_data_input_stream_read_uint16(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_uint32(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_uint64(param0 unsafe.Pointer) {}

func Fn_g_data_input_stream_read_until(param0 string, param1 *uint64, param2 unsafe.Pointer) {}

// UNSUPPORTED : read_until_async : has callback

func Fn_g_data_input_stream_read_until_finish(param0 unsafe.Pointer, param1 *uint64) {}

// UNSUPPORTED : read_upto_async : has callback

func Fn_g_data_input_stream_read_upto_finish(param0 unsafe.Pointer, param1 *uint64) {}

func Fn_g_data_input_stream_set_byte_order(param0 int) {}

func Fn_g_data_input_stream_set_newline_type(param0 int) {}

func Fn_g_data_output_stream_new(param0 unsafe.Pointer) {}

func Fn_g_data_output_stream_get_byte_order() {}

func Fn_g_data_output_stream_put_byte(param0 uint8, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_int16(param0 int16, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_int32(param0 int32, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_int64(param0 int64, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_string(param0 string, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_uint16(param0 uint16, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_uint32(param0 uint32, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_put_uint64(param0 uint64, param1 unsafe.Pointer) {}

func Fn_g_data_output_stream_set_byte_order(param0 int) {}

func Fn_g_desktop_app_info_new(param0 string) {}

func Fn_g_desktop_app_info_new_from_filename(param0 string) {}

func Fn_g_desktop_app_info_new_from_keyfile(param0 unsafe.Pointer) {}

func Fn_g_desktop_app_info_get_categories() {}

func Fn_g_desktop_app_info_get_filename() {}

func Fn_g_desktop_app_info_get_generic_name() {}

func Fn_g_desktop_app_info_get_is_hidden() {}

// UNSUPPORTED : launch_uris_as_manager : has callback

// UNSUPPORTED : launch_uris_as_manager_with_fds : has callback

func Fn_g_desktop_app_info_search(param0 string) {}

func Fn_g_desktop_app_info_set_desktop_env(param0 string) {}

func Fn_g_emblem_new(param0 unsafe.Pointer) {}

func Fn_g_emblem_new_with_origin(param0 unsafe.Pointer, param1 int) {}

func Fn_g_emblem_get_icon() {}

func Fn_g_emblem_get_origin() {}

func Fn_g_emblemed_icon_new(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_emblemed_icon_add_emblem(param0 unsafe.Pointer) {}

func Fn_g_emblemed_icon_get_emblems() {}

func Fn_g_emblemed_icon_get_icon() {}

func Fn_g_file_enumerator_close(param0 unsafe.Pointer) {}

// UNSUPPORTED : close_async : has callback

func Fn_g_file_enumerator_close_finish(param0 unsafe.Pointer) {}

func Fn_g_file_enumerator_get_container() {}

func Fn_g_file_enumerator_has_pending() {}

func Fn_g_file_enumerator_is_closed() {}

func Fn_g_file_enumerator_next_file(param0 unsafe.Pointer) {}

// UNSUPPORTED : next_files_async : has callback

func Fn_g_file_enumerator_next_files_finish(param0 unsafe.Pointer) {}

func Fn_g_file_enumerator_set_pending(param0 bool) {}

func Fn_g_file_io_stream_get_etag() {}

func Fn_g_file_io_stream_query_info(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_io_stream_query_info_finish(param0 unsafe.Pointer) {}

func Fn_g_file_icon_new(param0 unsafe.Pointer) {}

func Fn_g_file_icon_get_file() {}

func Fn_g_file_info_new() {
	C.g_file_info_new()
}

func Fn_g_file_info_clear_status() {}

func Fn_g_file_info_copy_into(param0 unsafe.Pointer) {}

func Fn_g_file_info_dup() {}

func Fn_g_file_info_get_attribute_as_string(param0 string) {}

func Fn_g_file_info_get_attribute_boolean(param0 string) {}

func Fn_g_file_info_get_attribute_byte_string(param0 string) {}

func Fn_g_file_info_get_attribute_data(param0 string, param1 int, param2 *unsafe.Pointer, param3 int) {}

func Fn_g_file_info_get_attribute_int32(param0 string) {}

func Fn_g_file_info_get_attribute_int64(param0 string) {}

func Fn_g_file_info_get_attribute_object(param0 string) {}

func Fn_g_file_info_get_attribute_status(param0 string) {}

func Fn_g_file_info_get_attribute_string(param0 string) {}

func Fn_g_file_info_get_attribute_stringv(param0 string) {}

func Fn_g_file_info_get_attribute_type(param0 string) {}

func Fn_g_file_info_get_attribute_uint32(param0 string) {}

func Fn_g_file_info_get_attribute_uint64(param0 string) {}

func Fn_g_file_info_get_content_type() {}

func Fn_g_file_info_get_display_name() {}

func Fn_g_file_info_get_edit_name() {}

func Fn_g_file_info_get_etag() {}

func Fn_g_file_info_get_file_type() {}

func Fn_g_file_info_get_icon() {}

func Fn_g_file_info_get_is_backup() {}

func Fn_g_file_info_get_is_hidden() {}

func Fn_g_file_info_get_is_symlink() {}

func Fn_g_file_info_get_modification_time(param0 unsafe.Pointer) {}

func Fn_g_file_info_get_name() {}

func Fn_g_file_info_get_size() {}

func Fn_g_file_info_get_sort_order() {}

func Fn_g_file_info_get_symlink_target() {}

func Fn_g_file_info_has_attribute(param0 string) {}

func Fn_g_file_info_has_namespace(param0 string) {}

func Fn_g_file_info_list_attributes(param0 string) {}

func Fn_g_file_info_remove_attribute(param0 string) {}

func Fn_g_file_info_set_attribute(param0 string, param1 int, param2 *unsafe.Pointer) {}

func Fn_g_file_info_set_attribute_boolean(param0 string, param1 bool) {}

func Fn_g_file_info_set_attribute_byte_string(param0 string, param1 string) {}

func Fn_g_file_info_set_attribute_int32(param0 string, param1 int32) {}

func Fn_g_file_info_set_attribute_int64(param0 string, param1 int64) {}

func Fn_g_file_info_set_attribute_mask(param0 unsafe.Pointer) {}

func Fn_g_file_info_set_attribute_object(param0 string, param1 unsafe.Pointer) {}

func Fn_g_file_info_set_attribute_status(param0 string, param1 int) {}

func Fn_g_file_info_set_attribute_string(param0 string, param1 string) {}

func Fn_g_file_info_set_attribute_stringv(param0 string, param1 []string) {}

func Fn_g_file_info_set_attribute_uint32(param0 string, param1 uint32) {}

func Fn_g_file_info_set_attribute_uint64(param0 string, param1 uint64) {}

func Fn_g_file_info_set_content_type(param0 string) {}

func Fn_g_file_info_set_display_name(param0 string) {}

func Fn_g_file_info_set_edit_name(param0 string) {}

func Fn_g_file_info_set_file_type(param0 int) {}

func Fn_g_file_info_set_icon(param0 unsafe.Pointer) {}

func Fn_g_file_info_set_is_hidden(param0 bool) {}

func Fn_g_file_info_set_is_symlink(param0 bool) {}

func Fn_g_file_info_set_modification_time(param0 unsafe.Pointer) {}

func Fn_g_file_info_set_name(param0 string) {}

func Fn_g_file_info_set_size(param0 int64) {}

func Fn_g_file_info_set_sort_order(param0 int32) {}

func Fn_g_file_info_set_symlink_target(param0 string) {}

func Fn_g_file_info_unset_attribute_mask() {}

func Fn_g_file_input_stream_query_info(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_input_stream_query_info_finish(param0 unsafe.Pointer) {}

func Fn_g_file_monitor_cancel() {}

func Fn_g_file_monitor_emit_event(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 int) {}

func Fn_g_file_monitor_is_cancelled() {}

func Fn_g_file_monitor_set_rate_limit(param0 int) {}

func Fn_g_file_output_stream_get_etag() {}

func Fn_g_file_output_stream_query_info(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : query_info_async : has callback

func Fn_g_file_output_stream_query_info_finish(param0 unsafe.Pointer) {}

func Fn_g_filename_completer_new() {
	C.g_filename_completer_new()
}

func Fn_g_filename_completer_get_completion_suffix(param0 string) {}

func Fn_g_filename_completer_get_completions(param0 string) {}

func Fn_g_filename_completer_set_dirs_only(param0 bool) {}

func Fn_g_filter_input_stream_get_base_stream() {}

func Fn_g_filter_input_stream_get_close_base_stream() {}

func Fn_g_filter_input_stream_set_close_base_stream(param0 bool) {}

func Fn_g_filter_output_stream_get_base_stream() {}

func Fn_g_filter_output_stream_get_close_base_stream() {}

func Fn_g_filter_output_stream_set_close_base_stream(param0 bool) {}

func Fn_g_io_module_new(param0 string) {}

func Fn_g_io_module_load() {}

func Fn_g_io_module_unload() {}

// UNSUPPORTED : query : blacklisted
func Fn_g_io_stream_clear_pending() {}

func Fn_g_io_stream_close(param0 unsafe.Pointer) {}

// UNSUPPORTED : close_async : has callback

func Fn_g_io_stream_close_finish(param0 unsafe.Pointer) {}

func Fn_g_io_stream_get_input_stream() {}

func Fn_g_io_stream_get_output_stream() {}

func Fn_g_io_stream_has_pending() {}

func Fn_g_io_stream_is_closed() {}

func Fn_g_io_stream_set_pending() {}

// UNSUPPORTED : splice_async : has callback

func Fn_g_inet_address_new_any(param0 int) {}

func Fn_g_inet_address_new_from_bytes(param0 []uint8, param1 int) {}

func Fn_g_inet_address_new_from_string(param0 string) {}

func Fn_g_inet_address_new_loopback(param0 int) {}

func Fn_g_inet_address_get_family() {}

func Fn_g_inet_address_get_is_any() {}

func Fn_g_inet_address_get_is_link_local() {}

func Fn_g_inet_address_get_is_loopback() {}

func Fn_g_inet_address_get_is_mc_global() {}

func Fn_g_inet_address_get_is_mc_link_local() {}

func Fn_g_inet_address_get_is_mc_node_local() {}

func Fn_g_inet_address_get_is_mc_org_local() {}

func Fn_g_inet_address_get_is_mc_site_local() {}

func Fn_g_inet_address_get_is_multicast() {}

func Fn_g_inet_address_get_is_site_local() {}

func Fn_g_inet_address_get_native_size() {}

func Fn_g_inet_address_to_bytes() {}

func Fn_g_inet_address_to_string() {}

func Fn_g_inet_socket_address_new(param0 unsafe.Pointer, param1 uint16) {}

func Fn_g_inet_socket_address_get_address() {}

func Fn_g_inet_socket_address_get_port() {}

func Fn_g_input_stream_clear_pending() {}

func Fn_g_input_stream_close(param0 unsafe.Pointer) {}

// UNSUPPORTED : close_async : has callback

func Fn_g_input_stream_close_finish(param0 unsafe.Pointer) {}

func Fn_g_input_stream_has_pending() {}

func Fn_g_input_stream_is_closed() {}

func Fn_g_input_stream_read(param0 []uint8, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_input_stream_read_all(param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer) {
}

// UNSUPPORTED : read_all_async : has callback

// UNSUPPORTED : read_async : has callback

// UNSUPPORTED : read_bytes_async : has callback

func Fn_g_input_stream_read_finish(param0 unsafe.Pointer) {}

func Fn_g_input_stream_set_pending() {}

func Fn_g_input_stream_skip(param0 uint64, param1 unsafe.Pointer) {}

// UNSUPPORTED : skip_async : has callback

func Fn_g_input_stream_skip_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : insert_sorted : has callback

// UNSUPPORTED : sort : has callback

func Fn_g_memory_input_stream_new() {
	C.g_memory_input_stream_new()
}

// UNSUPPORTED : new_from_data : has callback

// UNSUPPORTED : add_data : has callback

// UNSUPPORTED : new : has callback

func Fn_g_memory_output_stream_get_data() {}

func Fn_g_memory_output_stream_get_data_size() {}

func Fn_g_memory_output_stream_get_size() {}

// UNSUPPORTED : get_attribute : has varargs

// UNSUPPORTED : set_action_and_target : has varargs

// UNSUPPORTED : set_attribute : has varargs

// UNSUPPORTED : get_item_attribute : has varargs

func Fn_g_mount_operation_new() {
	C.g_mount_operation_new()
}

func Fn_g_mount_operation_get_anonymous() {}

func Fn_g_mount_operation_get_choice() {}

func Fn_g_mount_operation_get_domain() {}

func Fn_g_mount_operation_get_password() {}

func Fn_g_mount_operation_get_password_save() {}

func Fn_g_mount_operation_get_username() {}

func Fn_g_mount_operation_reply(param0 int) {}

func Fn_g_mount_operation_set_anonymous(param0 bool) {}

func Fn_g_mount_operation_set_choice(param0 int) {}

func Fn_g_mount_operation_set_domain(param0 string) {}

func Fn_g_mount_operation_set_password(param0 string) {}

func Fn_g_mount_operation_set_password_save(param0 int) {}

func Fn_g_mount_operation_set_username(param0 string) {}

func Fn_g_network_address_new(param0 string, param1 uint16) {}

func Fn_g_network_address_get_hostname() {}

func Fn_g_network_address_get_port() {}

func Fn_g_network_address_parse(param0 string, param1 uint16) {}

func Fn_g_network_service_new(param0 string, param1 string, param2 string) {}

func Fn_g_network_service_get_domain() {}

func Fn_g_network_service_get_protocol() {}

func Fn_g_network_service_get_service() {}

// UNSUPPORTED : add_button_with_target : has varargs

// UNSUPPORTED : set_default_action_and_target : has varargs

func Fn_g_notification_set_priority(param0 int) {}

func Fn_g_output_stream_clear_pending() {}

func Fn_g_output_stream_close(param0 unsafe.Pointer) {}

// UNSUPPORTED : close_async : has callback

func Fn_g_output_stream_close_finish(param0 unsafe.Pointer) {}

func Fn_g_output_stream_flush(param0 unsafe.Pointer) {}

// UNSUPPORTED : flush_async : has callback

func Fn_g_output_stream_flush_finish(param0 unsafe.Pointer) {}

func Fn_g_output_stream_has_pending() {}

func Fn_g_output_stream_is_closed() {}

func Fn_g_output_stream_is_closing() {}

// UNSUPPORTED : printf : has varargs

func Fn_g_output_stream_set_pending() {}

func Fn_g_output_stream_splice(param0 unsafe.Pointer, param1 int, param2 unsafe.Pointer) {}

// UNSUPPORTED : splice_async : has callback

func Fn_g_output_stream_splice_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : vprintf : has va_list

func Fn_g_output_stream_write(param0 []uint8, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_output_stream_write_all(param0 []uint8, param1 uint64, param2 *uint64, param3 unsafe.Pointer) {
}

// UNSUPPORTED : write_all_async : has callback

// UNSUPPORTED : write_async : has callback

func Fn_g_output_stream_write_bytes(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : write_bytes_async : has callback

func Fn_g_output_stream_write_bytes_finish(param0 unsafe.Pointer) {}

func Fn_g_output_stream_write_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : writev_all_async : has callback

// UNSUPPORTED : writev_async : has callback

// UNSUPPORTED : acquire_async : has callback

// UNSUPPORTED : release_async : has callback

func Fn_g_resolver_lookup_by_address(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : lookup_by_address_async : has callback

func Fn_g_resolver_lookup_by_address_finish(param0 unsafe.Pointer) {}

func Fn_g_resolver_lookup_by_name(param0 string, param1 unsafe.Pointer) {}

// UNSUPPORTED : lookup_by_name_async : has callback

func Fn_g_resolver_lookup_by_name_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : lookup_by_name_with_flags_async : has callback

// UNSUPPORTED : lookup_records_async : has callback

func Fn_g_resolver_lookup_service(param0 string, param1 string, param2 string, param3 unsafe.Pointer) {}

// UNSUPPORTED : lookup_service_async : has callback

func Fn_g_resolver_lookup_service_finish(param0 unsafe.Pointer) {}

func Fn_g_resolver_set_default() {}

func Fn_g_resolver_free_addresses(param0 unsafe.Pointer) {}

func Fn_g_resolver_free_targets(param0 unsafe.Pointer) {}

func Fn_g_resolver_get_default() {
	C.g_resolver_get_default()
}

func Fn_g_settings_apply() {}

// UNSUPPORTED : bind_with_mapping : has callback

// UNSUPPORTED : get : has varargs

// UNSUPPORTED : get_mapped : has callback

func Fn_g_settings_list_children() {}

func Fn_g_settings_list_keys() {}

func Fn_g_settings_reset(param0 string) {}

func Fn_g_settings_revert() {}

// UNSUPPORTED : set : has varargs

func Fn_g_settings_set_enum(param0 string, param1 int) {}

func Fn_g_settings_set_flags(param0 string, param1 uint) {}

func Fn_g_settings_sync() {
	C.g_settings_sync()
}

// UNSUPPORTED : changed : blacklisted
// UNSUPPORTED : changed_tree : blacklisted
// UNSUPPORTED : keys_changed : blacklisted
// UNSUPPORTED : path_changed : blacklisted
// UNSUPPORTED : path_writable_changed : blacklisted
// UNSUPPORTED : writable_changed : blacklisted
// UNSUPPORTED : flatten_tree : blacklisted
// UNSUPPORTED : get_default : blacklisted
// UNSUPPORTED : new : has callback

// UNSUPPORTED : new_error : has varargs

// UNSUPPORTED : new_from_error : has callback

// UNSUPPORTED : new_take_error : has callback

func Fn_g_simple_async_result_complete() {}

func Fn_g_simple_async_result_complete_in_idle() {}

func Fn_g_simple_async_result_get_op_res_gboolean() {}

func Fn_g_simple_async_result_get_op_res_gpointer() {}

func Fn_g_simple_async_result_get_op_res_gssize() {}

func Fn_g_simple_async_result_get_source_tag() {}

func Fn_g_simple_async_result_propagate_error() {}

// UNSUPPORTED : run_in_thread : has callback

// UNSUPPORTED : set_error : has varargs

// UNSUPPORTED : set_error_va : has va_list

func Fn_g_simple_async_result_set_from_error(param0 unsafe.Pointer) {}

func Fn_g_simple_async_result_set_handle_cancellation(param0 bool) {}

func Fn_g_simple_async_result_set_op_res_gboolean(param0 bool) {}

// UNSUPPORTED : set_op_res_gpointer : has callback

func Fn_g_simple_async_result_set_op_res_gssize(param0 uint64) {}

func Fn_g_simple_async_result_is_valid(param0 unsafe.Pointer, param1 unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_g_socket_new(param0 int, param1 int, param2 int) {}

func Fn_g_socket_new_from_fd(param0 int) {}

func Fn_g_socket_accept(param0 unsafe.Pointer) {}

func Fn_g_socket_bind(param0 unsafe.Pointer, param1 bool) {}

func Fn_g_socket_check_connect_result() {}

func Fn_g_socket_close() {}

func Fn_g_socket_condition_check(param0 int) {}

func Fn_g_socket_condition_wait(param0 int, param1 unsafe.Pointer) {}

func Fn_g_socket_connect(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_socket_connection_factory_create_connection() {}

func Fn_g_socket_create_source(param0 int, param1 unsafe.Pointer) {}

func Fn_g_socket_get_blocking() {}

func Fn_g_socket_get_family() {}

func Fn_g_socket_get_fd() {}

func Fn_g_socket_get_keepalive() {}

func Fn_g_socket_get_listen_backlog() {}

func Fn_g_socket_get_local_address() {}

func Fn_g_socket_get_protocol() {}

func Fn_g_socket_get_remote_address() {}

func Fn_g_socket_get_socket_type() {}

func Fn_g_socket_is_closed() {}

func Fn_g_socket_is_connected() {}

func Fn_g_socket_listen() {}

func Fn_g_socket_receive(param0 []uint8, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_socket_receive_from(param0 *unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer) {
}

func Fn_g_socket_receive_message(param0 *unsafe.Pointer, param1 []InputVector, param2 int, param3 []*unsafe.Pointer, param4 *int, param5 *int, param6 unsafe.Pointer) {
}

func Fn_g_socket_send(param0 []uint8, param1 uint64, param2 unsafe.Pointer) {}

func Fn_g_socket_send_message(param0 unsafe.Pointer, param1 []OutputVector, param2 int, param3 []unsafe.Pointer, param4 int, param5 int, param6 unsafe.Pointer) {
}

func Fn_g_socket_send_to(param0 unsafe.Pointer, param1 []uint8, param2 uint64, param3 unsafe.Pointer) {
}

func Fn_g_socket_set_blocking(param0 bool) {}

func Fn_g_socket_set_keepalive(param0 bool) {}

func Fn_g_socket_set_listen_backlog(param0 int) {}

func Fn_g_socket_shutdown(param0 bool, param1 bool) {}

func Fn_g_socket_speaks_ipv4() {}

func Fn_g_socket_address_new_from_native(param0 *unsafe.Pointer, param1 uint64) {}

func Fn_g_socket_address_get_family() {}

func Fn_g_socket_address_get_native_size() {}

func Fn_g_socket_address_to_native(param0 *unsafe.Pointer, param1 uint64) {}

func Fn_g_socket_address_enumerator_next(param0 unsafe.Pointer) {}

// UNSUPPORTED : next_async : has callback

func Fn_g_socket_address_enumerator_next_finish(param0 unsafe.Pointer) {}

func Fn_g_socket_client_new() {
	C.g_socket_client_new()
}

func Fn_g_socket_client_add_application_proxy(param0 string) {}

func Fn_g_socket_client_connect(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : connect_async : has callback

func Fn_g_socket_client_connect_finish(param0 unsafe.Pointer) {}

func Fn_g_socket_client_connect_to_host(param0 string, param1 uint16, param2 unsafe.Pointer) {}

// UNSUPPORTED : connect_to_host_async : has callback

func Fn_g_socket_client_connect_to_host_finish(param0 unsafe.Pointer) {}

func Fn_g_socket_client_connect_to_service(param0 string, param1 string, param2 unsafe.Pointer) {}

// UNSUPPORTED : connect_to_service_async : has callback

func Fn_g_socket_client_connect_to_service_finish(param0 unsafe.Pointer) {}

// UNSUPPORTED : connect_to_uri_async : has callback

func Fn_g_socket_client_get_family() {}

func Fn_g_socket_client_get_local_address() {}

func Fn_g_socket_client_get_protocol() {}

func Fn_g_socket_client_get_socket_type() {}

func Fn_g_socket_client_set_family(param0 int) {}

func Fn_g_socket_client_set_local_address(param0 unsafe.Pointer) {}

func Fn_g_socket_client_set_protocol(param0 int) {}

func Fn_g_socket_client_set_socket_type(param0 int) {}

// UNSUPPORTED : connect_async : has callback

func Fn_g_socket_connection_get_local_address() {}

func Fn_g_socket_connection_get_remote_address() {}

func Fn_g_socket_connection_get_socket() {}

func Fn_g_socket_connection_factory_lookup_type(param0 int, param1 int, param2 int) {}

func Fn_g_socket_connection_factory_register_type(param0 uint64, param1 int, param2 int, param3 int) {}

func Fn_g_socket_control_message_get_level() {}

func Fn_g_socket_control_message_get_msg_type() {}

func Fn_g_socket_control_message_get_size() {}

func Fn_g_socket_control_message_serialize(param0 *unsafe.Pointer) {}

func Fn_g_socket_control_message_deserialize(param0 int, param1 int, param2 uint64, param3 []uint8) {
}

func Fn_g_socket_listener_new() {
	C.g_socket_listener_new()
}

func Fn_g_socket_listener_accept(param0 *unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : accept_async : has callback

func Fn_g_socket_listener_accept_finish(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_socket_listener_accept_socket(param0 *unsafe.Pointer, param1 unsafe.Pointer) {}

// UNSUPPORTED : accept_socket_async : has callback

func Fn_g_socket_listener_accept_socket_finish(param0 unsafe.Pointer, param1 *unsafe.Pointer) {}

func Fn_g_socket_listener_add_address(param0 unsafe.Pointer, param1 int, param2 int, param3 unsafe.Pointer, param4 *unsafe.Pointer) {
}

func Fn_g_socket_listener_add_any_inet_port(param0 unsafe.Pointer) {}

func Fn_g_socket_listener_add_inet_port(param0 uint16, param1 unsafe.Pointer) {}

func Fn_g_socket_listener_add_socket(param0 unsafe.Pointer, param1 unsafe.Pointer) {}

func Fn_g_socket_listener_close() {}

func Fn_g_socket_listener_set_backlog(param0 int) {}

func Fn_g_socket_service_new() {
	C.g_socket_service_new()
}

func Fn_g_socket_service_is_active() {}

func Fn_g_socket_service_start() {}

func Fn_g_socket_service_stop() {}

// UNSUPPORTED : new : has varargs

// UNSUPPORTED : communicate_async : has callback

func Fn_g_subprocess_communicate_finish(param0 unsafe.Pointer, param1 *unsafe.Pointer, param2 *unsafe.Pointer) {
}

func Fn_g_subprocess_communicate_utf8(param0 string, param1 unsafe.Pointer, param2 string, param3 string) {
}

// UNSUPPORTED : communicate_utf8_async : has callback

func Fn_g_subprocess_communicate_utf8_finish(param0 unsafe.Pointer, param1 string, param2 string) {}

// UNSUPPORTED : wait_async : has callback

// UNSUPPORTED : wait_check_async : has callback

// UNSUPPORTED : set_child_setup : has callback

// UNSUPPORTED : spawn : has varargs

func Fn_g_subprocess_launcher_take_fd(param0 int, param1 int) {}

// UNSUPPORTED : new : has callback

// UNSUPPORTED : attach_source : has callback

// UNSUPPORTED : return_new_error : has varargs

// UNSUPPORTED : return_pointer : has callback

// UNSUPPORTED : run_in_thread : has callback

// UNSUPPORTED : run_in_thread_sync : has callback

// UNSUPPORTED : set_task_data : has callback

// UNSUPPORTED : report_error : has callback

// UNSUPPORTED : report_new_error : has varargs

func Fn_g_tcp_connection_get_graceful_disconnect() {}

func Fn_g_tcp_connection_set_graceful_disconnect(param0 bool) {}

func Fn_g_tcp_wrapper_connection_get_base_io_stream() {}

func Fn_g_test_dbus_new(param0 int) {}

func Fn_g_test_dbus_add_service_dir(param0 string) {}

func Fn_g_test_dbus_down() {}

func Fn_g_test_dbus_get_bus_address() {}

func Fn_g_test_dbus_get_flags() {}

func Fn_g_test_dbus_stop() {}

func Fn_g_test_dbus_up() {}

func Fn_g_test_dbus_unset() {
	C.g_test_dbus_unset()
}

func Fn_g_themed_icon_new(param0 string) {}

func Fn_g_themed_icon_new_from_names(param0 []string, param1 int) {}

func Fn_g_themed_icon_new_with_default_fallbacks(param0 string) {}

func Fn_g_themed_icon_append_name(param0 string) {}

func Fn_g_themed_icon_get_names() {}

func Fn_g_themed_icon_prepend_name(param0 string) {}

func Fn_g_threaded_socket_service_new(param0 int) {}

func Fn_g_tls_connection_get_use_system_certdb() {}

// UNSUPPORTED : handshake_async : has callback

func Fn_g_tls_connection_set_use_system_certdb(param0 bool) {}

// UNSUPPORTED : lookup_certificate_for_handle_async : has callback

// UNSUPPORTED : lookup_certificate_issuer_async : has callback

// UNSUPPORTED : lookup_certificates_issued_by_async : has callback

// UNSUPPORTED : verify_chain_async : has callback

// UNSUPPORTED : ask_password_async : has callback

// UNSUPPORTED : request_certificate_async : has callback

func Fn_g_tls_password_new(param0 int, param1 string) {}

// UNSUPPORTED : set_value_full : has callback

// UNSUPPORTED : receive_credentials_async : has callback

func Fn_g_unix_connection_receive_fd(param0 unsafe.Pointer) {}

// UNSUPPORTED : send_credentials_async : has callback

func Fn_g_unix_connection_send_fd(param0 int, param1 unsafe.Pointer) {}

func Fn_g_unix_fd_list_new() {
	C.g_unix_fd_list_new()
}

func Fn_g_unix_fd_list_new_from_array(param0 []int, param1 int) {}

func Fn_g_unix_fd_list_append(param0 int) {}

func Fn_g_unix_fd_list_get(param0 int) {}

func Fn_g_unix_fd_list_get_length() {}

func Fn_g_unix_fd_list_peek_fds(param0 *int) {}

func Fn_g_unix_fd_list_steal_fds(param0 *int) {}

func Fn_g_unix_fd_message_new() {
	C.g_unix_fd_message_new()
}

func Fn_g_unix_fd_message_new_with_fd_list(param0 unsafe.Pointer) {}

func Fn_g_unix_fd_message_append_fd(param0 int) {}

func Fn_g_unix_fd_message_get_fd_list() {}

func Fn_g_unix_fd_message_steal_fds(param0 *int) {}

func Fn_g_unix_input_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_input_stream_get_close_fd() {}

func Fn_g_unix_input_stream_get_fd() {}

func Fn_g_unix_input_stream_set_close_fd(param0 bool) {}

func Fn_g_unix_mount_monitor_new() {
	C.g_unix_mount_monitor_new()
}

func Fn_g_unix_mount_monitor_set_rate_limit(param0 int) {}

func Fn_g_unix_output_stream_new(param0 int, param1 bool) {}

func Fn_g_unix_output_stream_get_close_fd() {}

func Fn_g_unix_output_stream_get_fd() {}

func Fn_g_unix_output_stream_set_close_fd(param0 bool) {}

func Fn_g_unix_socket_address_new(param0 string) {}

func Fn_g_unix_socket_address_new_abstract(param0 []int8, param1 int) {}

func Fn_g_unix_socket_address_get_is_abstract() {}

func Fn_g_unix_socket_address_get_path() {}

func Fn_g_unix_socket_address_get_path_len() {}

func Fn_g_unix_socket_address_abstract_names_supported() {
	C.g_unix_socket_address_abstract_names_supported()
}

func Fn_g_vfs_get_file_for_path(param0 string) {}

func Fn_g_vfs_get_file_for_uri(param0 string) {}

func Fn_g_vfs_get_supported_uri_schemes() {}

func Fn_g_vfs_is_active() {}

func Fn_g_vfs_parse_name(param0 string) {}

// UNSUPPORTED : register_uri_scheme : has callback

func Fn_g_vfs_get_default() {
	C.g_vfs_get_default()
}

func Fn_g_vfs_get_local() {
	C.g_vfs_get_local()
}

func Fn_g_volume_monitor_get_connected_drives() {}

func Fn_g_volume_monitor_get_mount_for_uuid(param0 string) {}

func Fn_g_volume_monitor_get_mounts() {}

func Fn_g_volume_monitor_get_volume_for_uuid(param0 string) {}

func Fn_g_volume_monitor_get_volumes() {}

func Fn_g_volume_monitor_adopt_orphan_mount(param0 unsafe.Pointer) {}

func Fn_g_volume_monitor_get() {
	C.g_volume_monitor_get()
}

func Fn_g_zlib_compressor_new(param0 int, param1 int) {}

func Fn_g_zlib_decompressor_new(param0 int) {}
