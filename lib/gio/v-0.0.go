// Code generated - DO NOT EDIT.
// +build !gio_2.18,!gio_2.20,!gio_2.22,!gio_2.24,!gio_2.26,!gio_2.28,!gio_2.30,!gio_2.32,!gio_2.34,!gio_2.36,!gio_2.38,!gio_2.40,!gio_2.42,!gio_2.44,!gio_2.46,!gio_2.48,!gio_2.50,!gio_2.52,!gio_2.54,!gio_2.56,!gio_2.58,!gio_2.60,!gio_2.62

package gio

import (
	glib "github.com/pekim/gobbi/lib/glib"
	gio "github.com/pekim/gobbi/lib/internal/c/gio"
	"unsafe"
)

// DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME is a representation of the C constant G_DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME.
const DESKTOP_APP_INFO_LOOKUP_EXTENSION_POINT_NAME = "gio-desktop-app-info-lookup"

// FILE_ATTRIBUTE_ACCESS_CAN_DELETE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_DELETE.
const FILE_ATTRIBUTE_ACCESS_CAN_DELETE = "access::can-delete"

// FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE.
const FILE_ATTRIBUTE_ACCESS_CAN_EXECUTE = "access::can-execute"

// FILE_ATTRIBUTE_ACCESS_CAN_READ is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_READ.
const FILE_ATTRIBUTE_ACCESS_CAN_READ = "access::can-read"

// FILE_ATTRIBUTE_ACCESS_CAN_RENAME is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_RENAME.
const FILE_ATTRIBUTE_ACCESS_CAN_RENAME = "access::can-rename"

// FILE_ATTRIBUTE_ACCESS_CAN_TRASH is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_TRASH.
const FILE_ATTRIBUTE_ACCESS_CAN_TRASH = "access::can-trash"

// FILE_ATTRIBUTE_ACCESS_CAN_WRITE is a representation of the C constant G_FILE_ATTRIBUTE_ACCESS_CAN_WRITE.
const FILE_ATTRIBUTE_ACCESS_CAN_WRITE = "access::can-write"

// FILE_ATTRIBUTE_DOS_IS_ARCHIVE is a representation of the C constant G_FILE_ATTRIBUTE_DOS_IS_ARCHIVE.
const FILE_ATTRIBUTE_DOS_IS_ARCHIVE = "dos::is-archive"

// FILE_ATTRIBUTE_DOS_IS_SYSTEM is a representation of the C constant G_FILE_ATTRIBUTE_DOS_IS_SYSTEM.
const FILE_ATTRIBUTE_DOS_IS_SYSTEM = "dos::is-system"

// FILE_ATTRIBUTE_ETAG_VALUE is a representation of the C constant G_FILE_ATTRIBUTE_ETAG_VALUE.
const FILE_ATTRIBUTE_ETAG_VALUE = "etag::value"

// FILE_ATTRIBUTE_FILESYSTEM_FREE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_FREE.
const FILE_ATTRIBUTE_FILESYSTEM_FREE = "filesystem::free"

// FILE_ATTRIBUTE_FILESYSTEM_READONLY is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_READONLY.
const FILE_ATTRIBUTE_FILESYSTEM_READONLY = "filesystem::readonly"

// FILE_ATTRIBUTE_FILESYSTEM_REMOTE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_REMOTE.
const FILE_ATTRIBUTE_FILESYSTEM_REMOTE = "filesystem::remote"

// FILE_ATTRIBUTE_FILESYSTEM_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_SIZE.
const FILE_ATTRIBUTE_FILESYSTEM_SIZE = "filesystem::size"

// FILE_ATTRIBUTE_FILESYSTEM_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_TYPE.
const FILE_ATTRIBUTE_FILESYSTEM_TYPE = "filesystem::type"

// FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW is a representation of the C constant G_FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW.
const FILE_ATTRIBUTE_FILESYSTEM_USE_PREVIEW = "filesystem::use-preview"

// FILE_ATTRIBUTE_GVFS_BACKEND is a representation of the C constant G_FILE_ATTRIBUTE_GVFS_BACKEND.
const FILE_ATTRIBUTE_GVFS_BACKEND = "gvfs::backend"

// FILE_ATTRIBUTE_ID_FILE is a representation of the C constant G_FILE_ATTRIBUTE_ID_FILE.
const FILE_ATTRIBUTE_ID_FILE = "id::file"

// FILE_ATTRIBUTE_ID_FILESYSTEM is a representation of the C constant G_FILE_ATTRIBUTE_ID_FILESYSTEM.
const FILE_ATTRIBUTE_ID_FILESYSTEM = "id::filesystem"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_EJECT = "mountable::can-eject"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_MOUNT = "mountable::can-mount"

// FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT.
const FILE_ATTRIBUTE_MOUNTABLE_CAN_UNMOUNT = "mountable::can-unmount"

// FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI.
const FILE_ATTRIBUTE_MOUNTABLE_HAL_UDI = "mountable::hal-udi"

// FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE is a representation of the C constant G_FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE.
const FILE_ATTRIBUTE_MOUNTABLE_UNIX_DEVICE = "mountable::unix-device"

// FILE_ATTRIBUTE_OWNER_GROUP is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_GROUP.
const FILE_ATTRIBUTE_OWNER_GROUP = "owner::group"

// FILE_ATTRIBUTE_OWNER_USER is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_USER.
const FILE_ATTRIBUTE_OWNER_USER = "owner::user"

// FILE_ATTRIBUTE_OWNER_USER_REAL is a representation of the C constant G_FILE_ATTRIBUTE_OWNER_USER_REAL.
const FILE_ATTRIBUTE_OWNER_USER_REAL = "owner::user-real"

// FILE_ATTRIBUTE_SELINUX_CONTEXT is a representation of the C constant G_FILE_ATTRIBUTE_SELINUX_CONTEXT.
const FILE_ATTRIBUTE_SELINUX_CONTEXT = "selinux::context"

// FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
const FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE = "standard::content-type"

// FILE_ATTRIBUTE_STANDARD_COPY_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_COPY_NAME.
const FILE_ATTRIBUTE_STANDARD_COPY_NAME = "standard::copy-name"

// FILE_ATTRIBUTE_STANDARD_DESCRIPTION is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_DESCRIPTION.
const FILE_ATTRIBUTE_STANDARD_DESCRIPTION = "standard::description"

// FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
const FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME = "standard::display-name"

// FILE_ATTRIBUTE_STANDARD_EDIT_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
const FILE_ATTRIBUTE_STANDARD_EDIT_NAME = "standard::edit-name"

// FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE.
const FILE_ATTRIBUTE_STANDARD_FAST_CONTENT_TYPE = "standard::fast-content-type"

// FILE_ATTRIBUTE_STANDARD_ICON is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_ICON.
const FILE_ATTRIBUTE_STANDARD_ICON = "standard::icon"

// FILE_ATTRIBUTE_STANDARD_IS_BACKUP is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_BACKUP.
const FILE_ATTRIBUTE_STANDARD_IS_BACKUP = "standard::is-backup"

// FILE_ATTRIBUTE_STANDARD_IS_HIDDEN is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
const FILE_ATTRIBUTE_STANDARD_IS_HIDDEN = "standard::is-hidden"

// FILE_ATTRIBUTE_STANDARD_IS_SYMLINK is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
const FILE_ATTRIBUTE_STANDARD_IS_SYMLINK = "standard::is-symlink"

// FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL.
const FILE_ATTRIBUTE_STANDARD_IS_VIRTUAL = "standard::is-virtual"

// FILE_ATTRIBUTE_STANDARD_NAME is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_NAME.
const FILE_ATTRIBUTE_STANDARD_NAME = "standard::name"

// FILE_ATTRIBUTE_STANDARD_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SIZE.
const FILE_ATTRIBUTE_STANDARD_SIZE = "standard::size"

// FILE_ATTRIBUTE_STANDARD_SORT_ORDER is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
const FILE_ATTRIBUTE_STANDARD_SORT_ORDER = "standard::sort-order"

// FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET.
const FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET = "standard::symlink-target"

// FILE_ATTRIBUTE_STANDARD_TARGET_URI is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TARGET_URI.
const FILE_ATTRIBUTE_STANDARD_TARGET_URI = "standard::target-uri"

// FILE_ATTRIBUTE_STANDARD_TYPE is a representation of the C constant G_FILE_ATTRIBUTE_STANDARD_TYPE.
const FILE_ATTRIBUTE_STANDARD_TYPE = "standard::type"

// FILE_ATTRIBUTE_THUMBNAILING_FAILED is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAILING_FAILED.
const FILE_ATTRIBUTE_THUMBNAILING_FAILED = "thumbnail::failed"

// FILE_ATTRIBUTE_THUMBNAIL_PATH is a representation of the C constant G_FILE_ATTRIBUTE_THUMBNAIL_PATH.
const FILE_ATTRIBUTE_THUMBNAIL_PATH = "thumbnail::path"

// FILE_ATTRIBUTE_TIME_ACCESS is a representation of the C constant G_FILE_ATTRIBUTE_TIME_ACCESS.
const FILE_ATTRIBUTE_TIME_ACCESS = "time::access"

// FILE_ATTRIBUTE_TIME_ACCESS_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_ACCESS_USEC.
const FILE_ATTRIBUTE_TIME_ACCESS_USEC = "time::access-usec"

// FILE_ATTRIBUTE_TIME_CHANGED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CHANGED.
const FILE_ATTRIBUTE_TIME_CHANGED = "time::changed"

// FILE_ATTRIBUTE_TIME_CHANGED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CHANGED_USEC.
const FILE_ATTRIBUTE_TIME_CHANGED_USEC = "time::changed-usec"

// FILE_ATTRIBUTE_TIME_CREATED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CREATED.
const FILE_ATTRIBUTE_TIME_CREATED = "time::created"

// FILE_ATTRIBUTE_TIME_CREATED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_CREATED_USEC.
const FILE_ATTRIBUTE_TIME_CREATED_USEC = "time::created-usec"

// FILE_ATTRIBUTE_TIME_MODIFIED is a representation of the C constant G_FILE_ATTRIBUTE_TIME_MODIFIED.
const FILE_ATTRIBUTE_TIME_MODIFIED = "time::modified"

// FILE_ATTRIBUTE_TIME_MODIFIED_USEC is a representation of the C constant G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC.
const FILE_ATTRIBUTE_TIME_MODIFIED_USEC = "time::modified-usec"

// FILE_ATTRIBUTE_TRASH_ITEM_COUNT is a representation of the C constant G_FILE_ATTRIBUTE_TRASH_ITEM_COUNT.
const FILE_ATTRIBUTE_TRASH_ITEM_COUNT = "trash::item-count"

// FILE_ATTRIBUTE_UNIX_BLOCKS is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_BLOCKS.
const FILE_ATTRIBUTE_UNIX_BLOCKS = "unix::blocks"

// FILE_ATTRIBUTE_UNIX_BLOCK_SIZE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_BLOCK_SIZE.
const FILE_ATTRIBUTE_UNIX_BLOCK_SIZE = "unix::block-size"

// FILE_ATTRIBUTE_UNIX_DEVICE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_DEVICE.
const FILE_ATTRIBUTE_UNIX_DEVICE = "unix::device"

// FILE_ATTRIBUTE_UNIX_GID is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_GID.
const FILE_ATTRIBUTE_UNIX_GID = "unix::gid"

// FILE_ATTRIBUTE_UNIX_INODE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_INODE.
const FILE_ATTRIBUTE_UNIX_INODE = "unix::inode"

// FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT.
const FILE_ATTRIBUTE_UNIX_IS_MOUNTPOINT = "unix::is-mountpoint"

// FILE_ATTRIBUTE_UNIX_MODE is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_MODE.
const FILE_ATTRIBUTE_UNIX_MODE = "unix::mode"

// FILE_ATTRIBUTE_UNIX_NLINK is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_NLINK.
const FILE_ATTRIBUTE_UNIX_NLINK = "unix::nlink"

// FILE_ATTRIBUTE_UNIX_RDEV is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_RDEV.
const FILE_ATTRIBUTE_UNIX_RDEV = "unix::rdev"

// FILE_ATTRIBUTE_UNIX_UID is a representation of the C constant G_FILE_ATTRIBUTE_UNIX_UID.
const FILE_ATTRIBUTE_UNIX_UID = "unix::uid"

// NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME.
const NATIVE_VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-native-volume-monitor"

// PROXY_RESOLVER_EXTENSION_POINT_NAME is a representation of the C constant G_PROXY_RESOLVER_EXTENSION_POINT_NAME.
const PROXY_RESOLVER_EXTENSION_POINT_NAME = "gio-proxy-resolver"

// SETTINGS_BACKEND_EXTENSION_POINT_NAME is a representation of the C constant G_SETTINGS_BACKEND_EXTENSION_POINT_NAME.
const SETTINGS_BACKEND_EXTENSION_POINT_NAME = "gsettings-backend"

// TLS_BACKEND_EXTENSION_POINT_NAME is a representation of the C constant G_TLS_BACKEND_EXTENSION_POINT_NAME.
const TLS_BACKEND_EXTENSION_POINT_NAME = "gio-tls-backend"

// TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT is a representation of the C constant G_TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_CLIENT = "1.3.6.1.5.5.7.3.2"

// TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER is a representation of the C constant G_TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER.
const TLS_DATABASE_PURPOSE_AUTHENTICATE_SERVER = "1.3.6.1.5.5.7.3.1"

// VFS_EXTENSION_POINT_NAME is a representation of the C constant G_VFS_EXTENSION_POINT_NAME.
const VFS_EXTENSION_POINT_NAME = "gio-vfs"

// VOLUME_IDENTIFIER_KIND_CLASS is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_CLASS.
const VOLUME_IDENTIFIER_KIND_CLASS = "class"

// VOLUME_IDENTIFIER_KIND_HAL_UDI is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_HAL_UDI.
const VOLUME_IDENTIFIER_KIND_HAL_UDI = "hal-udi"

// VOLUME_IDENTIFIER_KIND_LABEL is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_LABEL.
const VOLUME_IDENTIFIER_KIND_LABEL = "label"

// VOLUME_IDENTIFIER_KIND_NFS_MOUNT is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_NFS_MOUNT.
const VOLUME_IDENTIFIER_KIND_NFS_MOUNT = "nfs-mount"

// VOLUME_IDENTIFIER_KIND_UNIX_DEVICE is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_UNIX_DEVICE.
const VOLUME_IDENTIFIER_KIND_UNIX_DEVICE = "unix-device"

// VOLUME_IDENTIFIER_KIND_UUID is a representation of the C constant G_VOLUME_IDENTIFIER_KIND_UUID.
const VOLUME_IDENTIFIER_KIND_UUID = "uuid"

// VOLUME_MONITOR_EXTENSION_POINT_NAME is a representation of the C constant G_VOLUME_MONITOR_EXTENSION_POINT_NAME.
const VOLUME_MONITOR_EXTENSION_POINT_NAME = "gio-volume-monitor"

// AppInfoCreateFlags is a representation of the C bitfield GAppInfoCreateFlags.
type AppInfoCreateFlags int

// AppInfoCreateFlags_none is a representation of the C bitfield member G_APP_INFO_CREATE_NONE.
const AppInfoCreateFlags_none = AppInfoCreateFlags(0)

// AppInfoCreateFlags_needs_terminal is a representation of the C bitfield member G_APP_INFO_CREATE_NEEDS_TERMINAL.
const AppInfoCreateFlags_needs_terminal = AppInfoCreateFlags(1)

// AppInfoCreateFlags_supports_uris is a representation of the C bitfield member G_APP_INFO_CREATE_SUPPORTS_URIS.
const AppInfoCreateFlags_supports_uris = AppInfoCreateFlags(2)

// AppInfoCreateFlags_supports_startup_notification is a representation of the C bitfield member G_APP_INFO_CREATE_SUPPORTS_STARTUP_NOTIFICATION.
const AppInfoCreateFlags_supports_startup_notification = AppInfoCreateFlags(4)

// AskPasswordFlags is a representation of the C bitfield GAskPasswordFlags.
type AskPasswordFlags int

// AskPasswordFlags_need_password is a representation of the C bitfield member G_ASK_PASSWORD_NEED_PASSWORD.
const AskPasswordFlags_need_password = AskPasswordFlags(1)

// AskPasswordFlags_need_username is a representation of the C bitfield member G_ASK_PASSWORD_NEED_USERNAME.
const AskPasswordFlags_need_username = AskPasswordFlags(2)

// AskPasswordFlags_need_domain is a representation of the C bitfield member G_ASK_PASSWORD_NEED_DOMAIN.
const AskPasswordFlags_need_domain = AskPasswordFlags(4)

// AskPasswordFlags_saving_supported is a representation of the C bitfield member G_ASK_PASSWORD_SAVING_SUPPORTED.
const AskPasswordFlags_saving_supported = AskPasswordFlags(8)

// AskPasswordFlags_anonymous_supported is a representation of the C bitfield member G_ASK_PASSWORD_ANONYMOUS_SUPPORTED.
const AskPasswordFlags_anonymous_supported = AskPasswordFlags(16)

// AskPasswordFlags_tcrypt is a representation of the C bitfield member G_ASK_PASSWORD_TCRYPT.
const AskPasswordFlags_tcrypt = AskPasswordFlags(32)

// FileAttributeInfoFlags is a representation of the C bitfield GFileAttributeInfoFlags.
type FileAttributeInfoFlags int

// FileAttributeInfoFlags_none is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_NONE.
const FileAttributeInfoFlags_none = FileAttributeInfoFlags(0)

// FileAttributeInfoFlags_copy_with_file is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_COPY_WITH_FILE.
const FileAttributeInfoFlags_copy_with_file = FileAttributeInfoFlags(1)

// FileAttributeInfoFlags_copy_when_moved is a representation of the C bitfield member G_FILE_ATTRIBUTE_INFO_COPY_WHEN_MOVED.
const FileAttributeInfoFlags_copy_when_moved = FileAttributeInfoFlags(2)

// FileCopyFlags is a representation of the C bitfield GFileCopyFlags.
type FileCopyFlags int

// FileCopyFlags_none is a representation of the C bitfield member G_FILE_COPY_NONE.
const FileCopyFlags_none = FileCopyFlags(0)

// FileCopyFlags_overwrite is a representation of the C bitfield member G_FILE_COPY_OVERWRITE.
const FileCopyFlags_overwrite = FileCopyFlags(1)

// FileCopyFlags_backup is a representation of the C bitfield member G_FILE_COPY_BACKUP.
const FileCopyFlags_backup = FileCopyFlags(2)

// FileCopyFlags_nofollow_symlinks is a representation of the C bitfield member G_FILE_COPY_NOFOLLOW_SYMLINKS.
const FileCopyFlags_nofollow_symlinks = FileCopyFlags(4)

// FileCopyFlags_all_metadata is a representation of the C bitfield member G_FILE_COPY_ALL_METADATA.
const FileCopyFlags_all_metadata = FileCopyFlags(8)

// FileCopyFlags_no_fallback_for_move is a representation of the C bitfield member G_FILE_COPY_NO_FALLBACK_FOR_MOVE.
const FileCopyFlags_no_fallback_for_move = FileCopyFlags(16)

// FileCopyFlags_target_default_perms is a representation of the C bitfield member G_FILE_COPY_TARGET_DEFAULT_PERMS.
const FileCopyFlags_target_default_perms = FileCopyFlags(32)

// FileCreateFlags is a representation of the C bitfield GFileCreateFlags.
type FileCreateFlags int

// FileCreateFlags_none is a representation of the C bitfield member G_FILE_CREATE_NONE.
const FileCreateFlags_none = FileCreateFlags(0)

// FileCreateFlags_private is a representation of the C bitfield member G_FILE_CREATE_PRIVATE.
const FileCreateFlags_private = FileCreateFlags(1)

// FileCreateFlags_replace_destination is a representation of the C bitfield member G_FILE_CREATE_REPLACE_DESTINATION.
const FileCreateFlags_replace_destination = FileCreateFlags(2)

// FileMonitorFlags is a representation of the C bitfield GFileMonitorFlags.
type FileMonitorFlags int

// FileMonitorFlags_none is a representation of the C bitfield member G_FILE_MONITOR_NONE.
const FileMonitorFlags_none = FileMonitorFlags(0)

// FileMonitorFlags_watch_mounts is a representation of the C bitfield member G_FILE_MONITOR_WATCH_MOUNTS.
const FileMonitorFlags_watch_mounts = FileMonitorFlags(1)

// FileMonitorFlags_send_moved is a representation of the C bitfield member G_FILE_MONITOR_SEND_MOVED.
const FileMonitorFlags_send_moved = FileMonitorFlags(2)

// FileMonitorFlags_watch_hard_links is a representation of the C bitfield member G_FILE_MONITOR_WATCH_HARD_LINKS.
const FileMonitorFlags_watch_hard_links = FileMonitorFlags(4)

// FileMonitorFlags_watch_moves is a representation of the C bitfield member G_FILE_MONITOR_WATCH_MOVES.
const FileMonitorFlags_watch_moves = FileMonitorFlags(8)

// FileQueryInfoFlags is a representation of the C bitfield GFileQueryInfoFlags.
type FileQueryInfoFlags int

// FileQueryInfoFlags_none is a representation of the C bitfield member G_FILE_QUERY_INFO_NONE.
const FileQueryInfoFlags_none = FileQueryInfoFlags(0)

// FileQueryInfoFlags_nofollow_symlinks is a representation of the C bitfield member G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS.
const FileQueryInfoFlags_nofollow_symlinks = FileQueryInfoFlags(1)

// MountMountFlags is a representation of the C bitfield GMountMountFlags.
type MountMountFlags int

// MountMountFlags_none is a representation of the C bitfield member G_MOUNT_MOUNT_NONE.
const MountMountFlags_none = MountMountFlags(0)

// MountUnmountFlags is a representation of the C bitfield GMountUnmountFlags.
type MountUnmountFlags int

// MountUnmountFlags_none is a representation of the C bitfield member G_MOUNT_UNMOUNT_NONE.
const MountUnmountFlags_none = MountUnmountFlags(0)

// MountUnmountFlags_force is a representation of the C bitfield member G_MOUNT_UNMOUNT_FORCE.
const MountUnmountFlags_force = MountUnmountFlags(1)

// OutputStreamSpliceFlags is a representation of the C bitfield GOutputStreamSpliceFlags.
type OutputStreamSpliceFlags int

// OutputStreamSpliceFlags_none is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_NONE.
const OutputStreamSpliceFlags_none = OutputStreamSpliceFlags(0)

// OutputStreamSpliceFlags_close_source is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_CLOSE_SOURCE.
const OutputStreamSpliceFlags_close_source = OutputStreamSpliceFlags(1)

// OutputStreamSpliceFlags_close_target is a representation of the C bitfield member G_OUTPUT_STREAM_SPLICE_CLOSE_TARGET.
const OutputStreamSpliceFlags_close_target = OutputStreamSpliceFlags(2)

// SettingsBindFlags is a representation of the C bitfield GSettingsBindFlags.
type SettingsBindFlags int

// SettingsBindFlags_default is a representation of the C bitfield member G_SETTINGS_BIND_DEFAULT.
const SettingsBindFlags_default = SettingsBindFlags(0)

// SettingsBindFlags_get is a representation of the C bitfield member G_SETTINGS_BIND_GET.
const SettingsBindFlags_get = SettingsBindFlags(1)

// SettingsBindFlags_set is a representation of the C bitfield member G_SETTINGS_BIND_SET.
const SettingsBindFlags_set = SettingsBindFlags(2)

// SettingsBindFlags_no_sensitivity is a representation of the C bitfield member G_SETTINGS_BIND_NO_SENSITIVITY.
const SettingsBindFlags_no_sensitivity = SettingsBindFlags(4)

// SettingsBindFlags_get_no_changes is a representation of the C bitfield member G_SETTINGS_BIND_GET_NO_CHANGES.
const SettingsBindFlags_get_no_changes = SettingsBindFlags(8)

// SettingsBindFlags_invert_boolean is a representation of the C bitfield member G_SETTINGS_BIND_INVERT_BOOLEAN.
const SettingsBindFlags_invert_boolean = SettingsBindFlags(16)

// DataStreamByteOrder is a representation of the C enumeration GDataStreamByteOrder.
type DataStreamByteOrder int

// DataStreamByteOrder_big_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_BIG_ENDIAN.
const DataStreamByteOrder_big_endian = DataStreamByteOrder(0)

// DataStreamByteOrder_little_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_LITTLE_ENDIAN.
const DataStreamByteOrder_little_endian = DataStreamByteOrder(1)

// DataStreamByteOrder_host_endian is a representation of the C enumeration member G_DATA_STREAM_BYTE_ORDER_HOST_ENDIAN.
const DataStreamByteOrder_host_endian = DataStreamByteOrder(2)

// DataStreamNewlineType is a representation of the C enumeration GDataStreamNewlineType.
type DataStreamNewlineType int

// DataStreamNewlineType_lf is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_LF.
const DataStreamNewlineType_lf = DataStreamNewlineType(0)

// DataStreamNewlineType_cr is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_CR.
const DataStreamNewlineType_cr = DataStreamNewlineType(1)

// DataStreamNewlineType_cr_lf is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_CR_LF.
const DataStreamNewlineType_cr_lf = DataStreamNewlineType(2)

// DataStreamNewlineType_any is a representation of the C enumeration member G_DATA_STREAM_NEWLINE_TYPE_ANY.
const DataStreamNewlineType_any = DataStreamNewlineType(3)

// FileAttributeStatus is a representation of the C enumeration GFileAttributeStatus.
type FileAttributeStatus int

// FileAttributeStatus_unset is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_UNSET.
const FileAttributeStatus_unset = FileAttributeStatus(0)

// FileAttributeStatus_set is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_SET.
const FileAttributeStatus_set = FileAttributeStatus(1)

// FileAttributeStatus_error_setting is a representation of the C enumeration member G_FILE_ATTRIBUTE_STATUS_ERROR_SETTING.
const FileAttributeStatus_error_setting = FileAttributeStatus(2)

// FileAttributeType is a representation of the C enumeration GFileAttributeType.
type FileAttributeType int

// FileAttributeType_invalid is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INVALID.
const FileAttributeType_invalid = FileAttributeType(0)

// FileAttributeType_string is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_STRING.
const FileAttributeType_string = FileAttributeType(1)

// FileAttributeType_byte_string is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_BYTE_STRING.
const FileAttributeType_byte_string = FileAttributeType(2)

// FileAttributeType_boolean is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
const FileAttributeType_boolean = FileAttributeType(3)

// FileAttributeType_uint32 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_UINT32.
const FileAttributeType_uint32 = FileAttributeType(4)

// FileAttributeType_int32 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INT32.
const FileAttributeType_int32 = FileAttributeType(5)

// FileAttributeType_uint64 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_UINT64.
const FileAttributeType_uint64 = FileAttributeType(6)

// FileAttributeType_int64 is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_INT64.
const FileAttributeType_int64 = FileAttributeType(7)

// FileAttributeType_object is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_OBJECT.
const FileAttributeType_object = FileAttributeType(8)

// FileAttributeType_stringv is a representation of the C enumeration member G_FILE_ATTRIBUTE_TYPE_STRINGV.
const FileAttributeType_stringv = FileAttributeType(9)

// FileMonitorEvent is a representation of the C enumeration GFileMonitorEvent.
type FileMonitorEvent int

// FileMonitorEvent_changed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CHANGED.
const FileMonitorEvent_changed = FileMonitorEvent(0)

// FileMonitorEvent_changes_done_hint is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CHANGES_DONE_HINT.
const FileMonitorEvent_changes_done_hint = FileMonitorEvent(1)

// FileMonitorEvent_deleted is a representation of the C enumeration member G_FILE_MONITOR_EVENT_DELETED.
const FileMonitorEvent_deleted = FileMonitorEvent(2)

// FileMonitorEvent_created is a representation of the C enumeration member G_FILE_MONITOR_EVENT_CREATED.
const FileMonitorEvent_created = FileMonitorEvent(3)

// FileMonitorEvent_attribute_changed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_ATTRIBUTE_CHANGED.
const FileMonitorEvent_attribute_changed = FileMonitorEvent(4)

// FileMonitorEvent_pre_unmount is a representation of the C enumeration member G_FILE_MONITOR_EVENT_PRE_UNMOUNT.
const FileMonitorEvent_pre_unmount = FileMonitorEvent(5)

// FileMonitorEvent_unmounted is a representation of the C enumeration member G_FILE_MONITOR_EVENT_UNMOUNTED.
const FileMonitorEvent_unmounted = FileMonitorEvent(6)

// FileMonitorEvent_moved is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED.
const FileMonitorEvent_moved = FileMonitorEvent(7)

// FileMonitorEvent_renamed is a representation of the C enumeration member G_FILE_MONITOR_EVENT_RENAMED.
const FileMonitorEvent_renamed = FileMonitorEvent(8)

// FileMonitorEvent_moved_in is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED_IN.
const FileMonitorEvent_moved_in = FileMonitorEvent(9)

// FileMonitorEvent_moved_out is a representation of the C enumeration member G_FILE_MONITOR_EVENT_MOVED_OUT.
const FileMonitorEvent_moved_out = FileMonitorEvent(10)

// FileType is a representation of the C enumeration GFileType.
type FileType int

// FileType_unknown is a representation of the C enumeration member G_FILE_TYPE_UNKNOWN.
const FileType_unknown = FileType(0)

// FileType_regular is a representation of the C enumeration member G_FILE_TYPE_REGULAR.
const FileType_regular = FileType(1)

// FileType_directory is a representation of the C enumeration member G_FILE_TYPE_DIRECTORY.
const FileType_directory = FileType(2)

// FileType_symbolic_link is a representation of the C enumeration member G_FILE_TYPE_SYMBOLIC_LINK.
const FileType_symbolic_link = FileType(3)

// FileType_special is a representation of the C enumeration member G_FILE_TYPE_SPECIAL.
const FileType_special = FileType(4)

// FileType_shortcut is a representation of the C enumeration member G_FILE_TYPE_SHORTCUT.
const FileType_shortcut = FileType(5)

// FileType_mountable is a representation of the C enumeration member G_FILE_TYPE_MOUNTABLE.
const FileType_mountable = FileType(6)

// FilesystemPreviewType is a representation of the C enumeration GFilesystemPreviewType.
type FilesystemPreviewType int

// FilesystemPreviewType_if_always is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_IF_ALWAYS.
const FilesystemPreviewType_if_always = FilesystemPreviewType(0)

// FilesystemPreviewType_if_local is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_IF_LOCAL.
const FilesystemPreviewType_if_local = FilesystemPreviewType(1)

// FilesystemPreviewType_never is a representation of the C enumeration member G_FILESYSTEM_PREVIEW_TYPE_NEVER.
const FilesystemPreviewType_never = FilesystemPreviewType(2)

// IOErrorEnum is a representation of the C enumeration GIOErrorEnum.
type IOErrorEnum int

// IOErrorEnum_failed is a representation of the C enumeration member G_IO_ERROR_FAILED.
const IOErrorEnum_failed = IOErrorEnum(0)

// IOErrorEnum_not_found is a representation of the C enumeration member G_IO_ERROR_NOT_FOUND.
const IOErrorEnum_not_found = IOErrorEnum(1)

// IOErrorEnum_exists is a representation of the C enumeration member G_IO_ERROR_EXISTS.
const IOErrorEnum_exists = IOErrorEnum(2)

// IOErrorEnum_is_directory is a representation of the C enumeration member G_IO_ERROR_IS_DIRECTORY.
const IOErrorEnum_is_directory = IOErrorEnum(3)

// IOErrorEnum_not_directory is a representation of the C enumeration member G_IO_ERROR_NOT_DIRECTORY.
const IOErrorEnum_not_directory = IOErrorEnum(4)

// IOErrorEnum_not_empty is a representation of the C enumeration member G_IO_ERROR_NOT_EMPTY.
const IOErrorEnum_not_empty = IOErrorEnum(5)

// IOErrorEnum_not_regular_file is a representation of the C enumeration member G_IO_ERROR_NOT_REGULAR_FILE.
const IOErrorEnum_not_regular_file = IOErrorEnum(6)

// IOErrorEnum_not_symbolic_link is a representation of the C enumeration member G_IO_ERROR_NOT_SYMBOLIC_LINK.
const IOErrorEnum_not_symbolic_link = IOErrorEnum(7)

// IOErrorEnum_not_mountable_file is a representation of the C enumeration member G_IO_ERROR_NOT_MOUNTABLE_FILE.
const IOErrorEnum_not_mountable_file = IOErrorEnum(8)

// IOErrorEnum_filename_too_long is a representation of the C enumeration member G_IO_ERROR_FILENAME_TOO_LONG.
const IOErrorEnum_filename_too_long = IOErrorEnum(9)

// IOErrorEnum_invalid_filename is a representation of the C enumeration member G_IO_ERROR_INVALID_FILENAME.
const IOErrorEnum_invalid_filename = IOErrorEnum(10)

// IOErrorEnum_too_many_links is a representation of the C enumeration member G_IO_ERROR_TOO_MANY_LINKS.
const IOErrorEnum_too_many_links = IOErrorEnum(11)

// IOErrorEnum_no_space is a representation of the C enumeration member G_IO_ERROR_NO_SPACE.
const IOErrorEnum_no_space = IOErrorEnum(12)

// IOErrorEnum_invalid_argument is a representation of the C enumeration member G_IO_ERROR_INVALID_ARGUMENT.
const IOErrorEnum_invalid_argument = IOErrorEnum(13)

// IOErrorEnum_permission_denied is a representation of the C enumeration member G_IO_ERROR_PERMISSION_DENIED.
const IOErrorEnum_permission_denied = IOErrorEnum(14)

// IOErrorEnum_not_supported is a representation of the C enumeration member G_IO_ERROR_NOT_SUPPORTED.
const IOErrorEnum_not_supported = IOErrorEnum(15)

// IOErrorEnum_not_mounted is a representation of the C enumeration member G_IO_ERROR_NOT_MOUNTED.
const IOErrorEnum_not_mounted = IOErrorEnum(16)

// IOErrorEnum_already_mounted is a representation of the C enumeration member G_IO_ERROR_ALREADY_MOUNTED.
const IOErrorEnum_already_mounted = IOErrorEnum(17)

// IOErrorEnum_closed is a representation of the C enumeration member G_IO_ERROR_CLOSED.
const IOErrorEnum_closed = IOErrorEnum(18)

// IOErrorEnum_cancelled is a representation of the C enumeration member G_IO_ERROR_CANCELLED.
const IOErrorEnum_cancelled = IOErrorEnum(19)

// IOErrorEnum_pending is a representation of the C enumeration member G_IO_ERROR_PENDING.
const IOErrorEnum_pending = IOErrorEnum(20)

// IOErrorEnum_read_only is a representation of the C enumeration member G_IO_ERROR_READ_ONLY.
const IOErrorEnum_read_only = IOErrorEnum(21)

// IOErrorEnum_cant_create_backup is a representation of the C enumeration member G_IO_ERROR_CANT_CREATE_BACKUP.
const IOErrorEnum_cant_create_backup = IOErrorEnum(22)

// IOErrorEnum_wrong_etag is a representation of the C enumeration member G_IO_ERROR_WRONG_ETAG.
const IOErrorEnum_wrong_etag = IOErrorEnum(23)

// IOErrorEnum_timed_out is a representation of the C enumeration member G_IO_ERROR_TIMED_OUT.
const IOErrorEnum_timed_out = IOErrorEnum(24)

// IOErrorEnum_would_recurse is a representation of the C enumeration member G_IO_ERROR_WOULD_RECURSE.
const IOErrorEnum_would_recurse = IOErrorEnum(25)

// IOErrorEnum_busy is a representation of the C enumeration member G_IO_ERROR_BUSY.
const IOErrorEnum_busy = IOErrorEnum(26)

// IOErrorEnum_would_block is a representation of the C enumeration member G_IO_ERROR_WOULD_BLOCK.
const IOErrorEnum_would_block = IOErrorEnum(27)

// IOErrorEnum_host_not_found is a representation of the C enumeration member G_IO_ERROR_HOST_NOT_FOUND.
const IOErrorEnum_host_not_found = IOErrorEnum(28)

// IOErrorEnum_would_merge is a representation of the C enumeration member G_IO_ERROR_WOULD_MERGE.
const IOErrorEnum_would_merge = IOErrorEnum(29)

// IOErrorEnum_failed_handled is a representation of the C enumeration member G_IO_ERROR_FAILED_HANDLED.
const IOErrorEnum_failed_handled = IOErrorEnum(30)

// IOErrorEnum_too_many_open_files is a representation of the C enumeration member G_IO_ERROR_TOO_MANY_OPEN_FILES.
const IOErrorEnum_too_many_open_files = IOErrorEnum(31)

// IOErrorEnum_not_initialized is a representation of the C enumeration member G_IO_ERROR_NOT_INITIALIZED.
const IOErrorEnum_not_initialized = IOErrorEnum(32)

// IOErrorEnum_address_in_use is a representation of the C enumeration member G_IO_ERROR_ADDRESS_IN_USE.
const IOErrorEnum_address_in_use = IOErrorEnum(33)

// IOErrorEnum_partial_input is a representation of the C enumeration member G_IO_ERROR_PARTIAL_INPUT.
const IOErrorEnum_partial_input = IOErrorEnum(34)

// IOErrorEnum_invalid_data is a representation of the C enumeration member G_IO_ERROR_INVALID_DATA.
const IOErrorEnum_invalid_data = IOErrorEnum(35)

// IOErrorEnum_dbus_error is a representation of the C enumeration member G_IO_ERROR_DBUS_ERROR.
const IOErrorEnum_dbus_error = IOErrorEnum(36)

// IOErrorEnum_host_unreachable is a representation of the C enumeration member G_IO_ERROR_HOST_UNREACHABLE.
const IOErrorEnum_host_unreachable = IOErrorEnum(37)

// IOErrorEnum_network_unreachable is a representation of the C enumeration member G_IO_ERROR_NETWORK_UNREACHABLE.
const IOErrorEnum_network_unreachable = IOErrorEnum(38)

// IOErrorEnum_connection_refused is a representation of the C enumeration member G_IO_ERROR_CONNECTION_REFUSED.
const IOErrorEnum_connection_refused = IOErrorEnum(39)

// IOErrorEnum_proxy_failed is a representation of the C enumeration member G_IO_ERROR_PROXY_FAILED.
const IOErrorEnum_proxy_failed = IOErrorEnum(40)

// IOErrorEnum_proxy_auth_failed is a representation of the C enumeration member G_IO_ERROR_PROXY_AUTH_FAILED.
const IOErrorEnum_proxy_auth_failed = IOErrorEnum(41)

// IOErrorEnum_proxy_need_auth is a representation of the C enumeration member G_IO_ERROR_PROXY_NEED_AUTH.
const IOErrorEnum_proxy_need_auth = IOErrorEnum(42)

// IOErrorEnum_proxy_not_allowed is a representation of the C enumeration member G_IO_ERROR_PROXY_NOT_ALLOWED.
const IOErrorEnum_proxy_not_allowed = IOErrorEnum(43)

// IOErrorEnum_broken_pipe is a representation of the C enumeration member G_IO_ERROR_BROKEN_PIPE.
const IOErrorEnum_broken_pipe = IOErrorEnum(44)

// IOErrorEnum_connection_closed is a representation of the C enumeration member G_IO_ERROR_CONNECTION_CLOSED.
const IOErrorEnum_connection_closed = IOErrorEnum(44)

// IOErrorEnum_not_connected is a representation of the C enumeration member G_IO_ERROR_NOT_CONNECTED.
const IOErrorEnum_not_connected = IOErrorEnum(45)

// IOErrorEnum_message_too_large is a representation of the C enumeration member G_IO_ERROR_MESSAGE_TOO_LARGE.
const IOErrorEnum_message_too_large = IOErrorEnum(46)

// MountOperationResult is a representation of the C enumeration GMountOperationResult.
type MountOperationResult int

// MountOperationResult_handled is a representation of the C enumeration member G_MOUNT_OPERATION_HANDLED.
const MountOperationResult_handled = MountOperationResult(0)

// MountOperationResult_aborted is a representation of the C enumeration member G_MOUNT_OPERATION_ABORTED.
const MountOperationResult_aborted = MountOperationResult(1)

// MountOperationResult_unhandled is a representation of the C enumeration member G_MOUNT_OPERATION_UNHANDLED.
const MountOperationResult_unhandled = MountOperationResult(2)

// PasswordSave is a representation of the C enumeration GPasswordSave.
type PasswordSave int

// PasswordSave_never is a representation of the C enumeration member G_PASSWORD_SAVE_NEVER.
const PasswordSave_never = PasswordSave(0)

// PasswordSave_for_session is a representation of the C enumeration member G_PASSWORD_SAVE_FOR_SESSION.
const PasswordSave_for_session = PasswordSave(1)

// PasswordSave_permanently is a representation of the C enumeration member G_PASSWORD_SAVE_PERMANENTLY.
const PasswordSave_permanently = PasswordSave(2)

// UNSUPPORTED : g_action_parse_detailed_name : parameter 'action_name' is non array with indirect count > 1

// UNSUPPORTED : g_app_info_create_from_commandline : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri : throws

// UNSUPPORTED : g_app_info_launch_default_for_uri_async : parameter 'callback' is callback

// UNSUPPORTED : g_app_info_launch_default_for_uri_finish : throws

// UNSUPPORTED : g_async_initable_newv_async : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get : parameter 'callback' is callback

// UNSUPPORTED : g_bus_get_finish : throws

// UNSUPPORTED : g_bus_get_sync : throws

// UNSUPPORTED : g_bus_own_name : parameter 'bus_acquired_handler' is callback

// UNSUPPORTED : g_bus_own_name_on_connection : parameter 'name_acquired_handler' is callback

// UNSUPPORTED : g_bus_watch_name : parameter 'name_appeared_handler' is callback

// UNSUPPORTED : g_bus_watch_name_on_connection : parameter 'name_appeared_handler' is callback

// ContentTypeCanBeExecutable wraps the C function g_content_type_can_be_executable.
func ContentTypeCanBeExecutable(type_ string) bool {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_can_be_executable(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeEquals wraps the C function g_content_type_equals.
func ContentTypeEquals(type1 string, type2 string) bool {
	sys_type1 := type1
	sys_type2 := type2
	retSys := gio.Fn_g_content_type_equals(sys_type1, sys_type2)
	ret := retSys

	return ret
}

// ContentTypeGetDescription wraps the C function g_content_type_get_description.
func ContentTypeGetDescription(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_description(sys_type_)
	ret := retSys

	return ret
}

// ContentTypeGetIcon wraps the C function g_content_type_get_icon.
func ContentTypeGetIcon(type_ string) *Icon {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_icon(sys_type_)
	ret := IconNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_content_type_get_mime_dirs : no array length

// ContentTypeGetMimeType wraps the C function g_content_type_get_mime_type.
func ContentTypeGetMimeType(type_ string) string {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_get_mime_type(sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_content_type_guess : has array param, data

// UNSUPPORTED : g_content_type_guess_for_tree : no array length

// ContentTypeIsA wraps the C function g_content_type_is_a.
func ContentTypeIsA(type_ string, supertype string) bool {
	sys_type_ := type_
	sys_supertype := supertype
	retSys := gio.Fn_g_content_type_is_a(sys_type_, sys_supertype)
	ret := retSys

	return ret
}

// ContentTypeIsUnknown wraps the C function g_content_type_is_unknown.
func ContentTypeIsUnknown(type_ string) bool {
	sys_type_ := type_
	retSys := gio.Fn_g_content_type_is_unknown(sys_type_)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_content_type_set_mime_dirs : has array param, dirs

// ContentTypesGetRegistered wraps the C function g_content_types_get_registered.
func ContentTypesGetRegistered() *glib.List {
	retSys := gio.Fn_g_content_types_get_registered()
	ret := glib.ListNewFromC(retSys)

	return ret
}

// UNSUPPORTED : g_dbus_address_get_for_bus_sync : throws

// UNSUPPORTED : g_dbus_address_get_stream : parameter 'callback' is callback

// UNSUPPORTED : g_dbus_address_get_stream_finish : parameter 'out_guid' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_address_get_stream_sync : parameter 'out_guid' is non array with indirect count > 1

// UNSUPPORTED : g_dbus_annotation_info_lookup : has array param, annotations

// UNSUPPORTED : g_dbus_error_register_error_domain : has array param, entries

// UNSUPPORTED : g_dbus_gvariant_to_gvalue : has [in]out param, out_gvalue

// UNSUPPORTED : g_dbus_is_supported_address : throws

// UNSUPPORTED : g_dtls_client_connection_new : throws

// UNSUPPORTED : g_dtls_server_connection_new : throws

// UNSUPPORTED : g_file_new_tmp : parameter 'iostream' is non array with indirect count > 1

// UNSUPPORTED : g_icon_new_for_string : throws

// UNSUPPORTED : g_initable_newv : throws

// IoErrorFromErrno wraps the C function g_io_error_from_errno.
func IoErrorFromErrno(errNo int) int {
	sys_errNo := errNo
	retSys := gio.Fn_g_io_error_from_errno(sys_errNo)
	ret := retSys

	return ret
}

// IoErrorQuark wraps the C function g_io_error_quark.
func IoErrorQuark() uint32 {
	retSys := gio.Fn_g_io_error_quark()
	ret := retSys

	return ret
}

// UNSUPPORTED : g_io_modules_load_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_load_all_in_directory_with_scope : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory : blacklisted

// UNSUPPORTED : g_io_modules_scan_all_in_directory_with_scope : blacklisted

// IoSchedulerCancelAllJobs wraps the C function g_io_scheduler_cancel_all_jobs.
func IoSchedulerCancelAllJobs() {
	gio.Fn_g_io_scheduler_cancel_all_jobs()
}

// UNSUPPORTED : g_io_scheduler_push_job : parameter 'job_func' is callback

// UNSUPPORTED : g_keyfile_settings_backend_new : blacklisted

// UNSUPPORTED : g_memory_settings_backend_new : blacklisted

// UNSUPPORTED : g_null_settings_backend_new : blacklisted

// UNSUPPORTED : g_pollable_stream_read : throws

// UNSUPPORTED : g_pollable_stream_write : throws

// UNSUPPORTED : g_pollable_stream_write_all : throws

// UNSUPPORTED : g_resource_load : throws

// UNSUPPORTED : g_resources_enumerate_children : no array length

// UNSUPPORTED : g_resources_get_info : throws

// UNSUPPORTED : g_resources_lookup_data : throws

// UNSUPPORTED : g_resources_open_stream : throws

// UNSUPPORTED : g_simple_async_report_error_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_simple_async_report_take_gerror_in_idle : parameter 'callback' is callback

// UNSUPPORTED : g_tls_client_connection_new : throws

// UNSUPPORTED : g_tls_file_database_new : throws

// UNSUPPORTED : g_tls_server_connection_new : throws

// UnixIsMountPathSystemInternal wraps the C function g_unix_is_mount_path_system_internal.
func UnixIsMountPathSystemInternal(mountPath string) bool {
	sys_mountPath := mountPath
	retSys := gio.Fn_g_unix_is_mount_path_system_internal(sys_mountPath)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_unix_mount_at : has [in]out param, time_read

// UnixMountCompare wraps the C function g_unix_mount_compare.
func UnixMountCompare(mount1 *UnixMountEntry, mount2 *UnixMountEntry) int {
	sys_mount1 := mount1.ToC()
	sys_mount2 := mount2.ToC()
	retSys := gio.Fn_g_unix_mount_compare(sys_mount1, sys_mount2)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_unix_mount_for : has [in]out param, time_read

// UnixMountFree wraps the C function g_unix_mount_free.
func UnixMountFree(mountEntry *UnixMountEntry) {
	sys_mountEntry := mountEntry.ToC()
	gio.Fn_g_unix_mount_free(sys_mountEntry)
}

// UnixMountGetDevicePath wraps the C function g_unix_mount_get_device_path.
func UnixMountGetDevicePath(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_device_path(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGetFsType wraps the C function g_unix_mount_get_fs_type.
func UnixMountGetFsType(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_fs_type(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGetMountPath wraps the C function g_unix_mount_get_mount_path.
func UnixMountGetMountPath(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_get_mount_path(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessCanEject wraps the C function g_unix_mount_guess_can_eject.
func UnixMountGuessCanEject(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_can_eject(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessIcon wraps the C function g_unix_mount_guess_icon.
func UnixMountGuessIcon(mountEntry *UnixMountEntry) *Icon {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_icon(sys_mountEntry)
	ret := IconNewFromC(retSys)

	return ret
}

// UnixMountGuessName wraps the C function g_unix_mount_guess_name.
func UnixMountGuessName(mountEntry *UnixMountEntry) string {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_name(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountGuessShouldDisplay wraps the C function g_unix_mount_guess_should_display.
func UnixMountGuessShouldDisplay(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_guess_should_display(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountIsReadonly wraps the C function g_unix_mount_is_readonly.
func UnixMountIsReadonly(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_is_readonly(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountIsSystemInternal wraps the C function g_unix_mount_is_system_internal.
func UnixMountIsSystemInternal(mountEntry *UnixMountEntry) bool {
	sys_mountEntry := mountEntry.ToC()
	retSys := gio.Fn_g_unix_mount_is_system_internal(sys_mountEntry)
	ret := retSys

	return ret
}

// UnixMountPointsChangedSince wraps the C function g_unix_mount_points_changed_since.
func UnixMountPointsChangedSince(time uint64) bool {
	sys_time := time
	retSys := gio.Fn_g_unix_mount_points_changed_since(sys_time)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_unix_mount_points_get : has [in]out param, time_read

// UnixMountsChangedSince wraps the C function g_unix_mounts_changed_since.
func UnixMountsChangedSince(time uint64) bool {
	sys_time := time
	retSys := gio.Fn_g_unix_mounts_changed_since(sys_time)
	ret := retSys

	return ret
}

// UNSUPPORTED : g_unix_mounts_get : has [in]out param, time_read

// ActionEntry is a representation of the C record GActionEntry.
type ActionEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionEntry that represents the ActionEntry.
func (recv *ActionEntry) ToC() unsafe.Pointer {
	return recv.native
}

// ActionEntryNewFromC creates a new ActionEntry from a pointer to the C GActionEntry that represents the ActionEntry.
func ActionEntryNewFromC(native unsafe.Pointer) *ActionEntry {
	return &ActionEntry{native: native}
}

// AppInfoIface is a representation of the C record GAppInfoIface.
type AppInfoIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfoIface that represents the AppInfoIface.
func (recv *AppInfoIface) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoIfaceNewFromC creates a new AppInfoIface from a pointer to the C GAppInfoIface that represents the AppInfoIface.
func AppInfoIfaceNewFromC(native unsafe.Pointer) *AppInfoIface {
	return &AppInfoIface{native: native}
}

// AppLaunchContextClass is a representation of the C record GAppLaunchContextClass.
type AppLaunchContextClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextClass that represents the AppLaunchContextClass.
func (recv *AppLaunchContextClass) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextClassNewFromC creates a new AppLaunchContextClass from a pointer to the C GAppLaunchContextClass that represents the AppLaunchContextClass.
func AppLaunchContextClassNewFromC(native unsafe.Pointer) *AppLaunchContextClass {
	return &AppLaunchContextClass{native: native}
}

// AppLaunchContextPrivate is a representation of the C record GAppLaunchContextPrivate.
type AppLaunchContextPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContextPrivate that represents the AppLaunchContextPrivate.
func (recv *AppLaunchContextPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextPrivateNewFromC creates a new AppLaunchContextPrivate from a pointer to the C GAppLaunchContextPrivate that represents the AppLaunchContextPrivate.
func AppLaunchContextPrivateNewFromC(native unsafe.Pointer) *AppLaunchContextPrivate {
	return &AppLaunchContextPrivate{native: native}
}

// ApplicationCommandLinePrivate is a representation of the C record GApplicationCommandLinePrivate.
type ApplicationCommandLinePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLinePrivate that represents the ApplicationCommandLinePrivate.
func (recv *ApplicationCommandLinePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLinePrivateNewFromC creates a new ApplicationCommandLinePrivate from a pointer to the C GApplicationCommandLinePrivate that represents the ApplicationCommandLinePrivate.
func ApplicationCommandLinePrivateNewFromC(native unsafe.Pointer) *ApplicationCommandLinePrivate {
	return &ApplicationCommandLinePrivate{native: native}
}

// ApplicationPrivate is a representation of the C record GApplicationPrivate.
type ApplicationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationPrivate that represents the ApplicationPrivate.
func (recv *ApplicationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationPrivateNewFromC creates a new ApplicationPrivate from a pointer to the C GApplicationPrivate that represents the ApplicationPrivate.
func ApplicationPrivateNewFromC(native unsafe.Pointer) *ApplicationPrivate {
	return &ApplicationPrivate{native: native}
}

// AsyncResultIface is a representation of the C record GAsyncResultIface.
type AsyncResultIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResultIface that represents the AsyncResultIface.
func (recv *AsyncResultIface) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResultIfaceNewFromC creates a new AsyncResultIface from a pointer to the C GAsyncResultIface that represents the AsyncResultIface.
func AsyncResultIfaceNewFromC(native unsafe.Pointer) *AsyncResultIface {
	return &AsyncResultIface{native: native}
}

// BufferedInputStreamClass is a representation of the C record GBufferedInputStreamClass.
type BufferedInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamClass that represents the BufferedInputStreamClass.
func (recv *BufferedInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamClassNewFromC creates a new BufferedInputStreamClass from a pointer to the C GBufferedInputStreamClass that represents the BufferedInputStreamClass.
func BufferedInputStreamClassNewFromC(native unsafe.Pointer) *BufferedInputStreamClass {
	return &BufferedInputStreamClass{native: native}
}

// BufferedInputStreamPrivate is a representation of the C record GBufferedInputStreamPrivate.
type BufferedInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStreamPrivate that represents the BufferedInputStreamPrivate.
func (recv *BufferedInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamPrivateNewFromC creates a new BufferedInputStreamPrivate from a pointer to the C GBufferedInputStreamPrivate that represents the BufferedInputStreamPrivate.
func BufferedInputStreamPrivateNewFromC(native unsafe.Pointer) *BufferedInputStreamPrivate {
	return &BufferedInputStreamPrivate{native: native}
}

// BufferedOutputStreamClass is a representation of the C record GBufferedOutputStreamClass.
type BufferedOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamClass that represents the BufferedOutputStreamClass.
func (recv *BufferedOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamClassNewFromC creates a new BufferedOutputStreamClass from a pointer to the C GBufferedOutputStreamClass that represents the BufferedOutputStreamClass.
func BufferedOutputStreamClassNewFromC(native unsafe.Pointer) *BufferedOutputStreamClass {
	return &BufferedOutputStreamClass{native: native}
}

// BufferedOutputStreamPrivate is a representation of the C record GBufferedOutputStreamPrivate.
type BufferedOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStreamPrivate that represents the BufferedOutputStreamPrivate.
func (recv *BufferedOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamPrivateNewFromC creates a new BufferedOutputStreamPrivate from a pointer to the C GBufferedOutputStreamPrivate that represents the BufferedOutputStreamPrivate.
func BufferedOutputStreamPrivateNewFromC(native unsafe.Pointer) *BufferedOutputStreamPrivate {
	return &BufferedOutputStreamPrivate{native: native}
}

// CancellableClass is a representation of the C record GCancellableClass.
type CancellableClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellableClass that represents the CancellableClass.
func (recv *CancellableClass) ToC() unsafe.Pointer {
	return recv.native
}

// CancellableClassNewFromC creates a new CancellableClass from a pointer to the C GCancellableClass that represents the CancellableClass.
func CancellableClassNewFromC(native unsafe.Pointer) *CancellableClass {
	return &CancellableClass{native: native}
}

// CancellablePrivate is a representation of the C record GCancellablePrivate.
type CancellablePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellablePrivate that represents the CancellablePrivate.
func (recv *CancellablePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// CancellablePrivateNewFromC creates a new CancellablePrivate from a pointer to the C GCancellablePrivate that represents the CancellablePrivate.
func CancellablePrivateNewFromC(native unsafe.Pointer) *CancellablePrivate {
	return &CancellablePrivate{native: native}
}

// CharsetConverterClass is a representation of the C record GCharsetConverterClass.
type CharsetConverterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverterClass that represents the CharsetConverterClass.
func (recv *CharsetConverterClass) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverterClassNewFromC creates a new CharsetConverterClass from a pointer to the C GCharsetConverterClass that represents the CharsetConverterClass.
func CharsetConverterClassNewFromC(native unsafe.Pointer) *CharsetConverterClass {
	return &CharsetConverterClass{native: native}
}

// ConverterInputStreamClass is a representation of the C record GConverterInputStreamClass.
type ConverterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamClass that represents the ConverterInputStreamClass.
func (recv *ConverterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamClassNewFromC creates a new ConverterInputStreamClass from a pointer to the C GConverterInputStreamClass that represents the ConverterInputStreamClass.
func ConverterInputStreamClassNewFromC(native unsafe.Pointer) *ConverterInputStreamClass {
	return &ConverterInputStreamClass{native: native}
}

// ConverterInputStreamPrivate is a representation of the C record GConverterInputStreamPrivate.
type ConverterInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStreamPrivate that represents the ConverterInputStreamPrivate.
func (recv *ConverterInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamPrivateNewFromC creates a new ConverterInputStreamPrivate from a pointer to the C GConverterInputStreamPrivate that represents the ConverterInputStreamPrivate.
func ConverterInputStreamPrivateNewFromC(native unsafe.Pointer) *ConverterInputStreamPrivate {
	return &ConverterInputStreamPrivate{native: native}
}

// ConverterOutputStreamClass is a representation of the C record GConverterOutputStreamClass.
type ConverterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamClass that represents the ConverterOutputStreamClass.
func (recv *ConverterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamClassNewFromC creates a new ConverterOutputStreamClass from a pointer to the C GConverterOutputStreamClass that represents the ConverterOutputStreamClass.
func ConverterOutputStreamClassNewFromC(native unsafe.Pointer) *ConverterOutputStreamClass {
	return &ConverterOutputStreamClass{native: native}
}

// ConverterOutputStreamPrivate is a representation of the C record GConverterOutputStreamPrivate.
type ConverterOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStreamPrivate that represents the ConverterOutputStreamPrivate.
func (recv *ConverterOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamPrivateNewFromC creates a new ConverterOutputStreamPrivate from a pointer to the C GConverterOutputStreamPrivate that represents the ConverterOutputStreamPrivate.
func ConverterOutputStreamPrivateNewFromC(native unsafe.Pointer) *ConverterOutputStreamPrivate {
	return &ConverterOutputStreamPrivate{native: native}
}

// DBusInterfaceSkeletonPrivate is a representation of the C record GDBusInterfaceSkeletonPrivate.
type DBusInterfaceSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusInterfaceSkeletonPrivate that represents the DBusInterfaceSkeletonPrivate.
func (recv *DBusInterfaceSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusInterfaceSkeletonPrivateNewFromC creates a new DBusInterfaceSkeletonPrivate from a pointer to the C GDBusInterfaceSkeletonPrivate that represents the DBusInterfaceSkeletonPrivate.
func DBusInterfaceSkeletonPrivateNewFromC(native unsafe.Pointer) *DBusInterfaceSkeletonPrivate {
	return &DBusInterfaceSkeletonPrivate{native: native}
}

// DBusObjectManagerClientPrivate is a representation of the C record GDBusObjectManagerClientPrivate.
type DBusObjectManagerClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerClientPrivate that represents the DBusObjectManagerClientPrivate.
func (recv *DBusObjectManagerClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerClientPrivateNewFromC creates a new DBusObjectManagerClientPrivate from a pointer to the C GDBusObjectManagerClientPrivate that represents the DBusObjectManagerClientPrivate.
func DBusObjectManagerClientPrivateNewFromC(native unsafe.Pointer) *DBusObjectManagerClientPrivate {
	return &DBusObjectManagerClientPrivate{native: native}
}

// DBusObjectManagerServerPrivate is a representation of the C record GDBusObjectManagerServerPrivate.
type DBusObjectManagerServerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManagerServerPrivate that represents the DBusObjectManagerServerPrivate.
func (recv *DBusObjectManagerServerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerServerPrivateNewFromC creates a new DBusObjectManagerServerPrivate from a pointer to the C GDBusObjectManagerServerPrivate that represents the DBusObjectManagerServerPrivate.
func DBusObjectManagerServerPrivateNewFromC(native unsafe.Pointer) *DBusObjectManagerServerPrivate {
	return &DBusObjectManagerServerPrivate{native: native}
}

// DBusObjectProxyPrivate is a representation of the C record GDBusObjectProxyPrivate.
type DBusObjectProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectProxyPrivate that represents the DBusObjectProxyPrivate.
func (recv *DBusObjectProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectProxyPrivateNewFromC creates a new DBusObjectProxyPrivate from a pointer to the C GDBusObjectProxyPrivate that represents the DBusObjectProxyPrivate.
func DBusObjectProxyPrivateNewFromC(native unsafe.Pointer) *DBusObjectProxyPrivate {
	return &DBusObjectProxyPrivate{native: native}
}

// DBusObjectSkeletonPrivate is a representation of the C record GDBusObjectSkeletonPrivate.
type DBusObjectSkeletonPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectSkeletonPrivate that represents the DBusObjectSkeletonPrivate.
func (recv *DBusObjectSkeletonPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectSkeletonPrivateNewFromC creates a new DBusObjectSkeletonPrivate from a pointer to the C GDBusObjectSkeletonPrivate that represents the DBusObjectSkeletonPrivate.
func DBusObjectSkeletonPrivateNewFromC(native unsafe.Pointer) *DBusObjectSkeletonPrivate {
	return &DBusObjectSkeletonPrivate{native: native}
}

// DBusProxyPrivate is a representation of the C record GDBusProxyPrivate.
type DBusProxyPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusProxyPrivate that represents the DBusProxyPrivate.
func (recv *DBusProxyPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DBusProxyPrivateNewFromC creates a new DBusProxyPrivate from a pointer to the C GDBusProxyPrivate that represents the DBusProxyPrivate.
func DBusProxyPrivateNewFromC(native unsafe.Pointer) *DBusProxyPrivate {
	return &DBusProxyPrivate{native: native}
}

// DataInputStreamClass is a representation of the C record GDataInputStreamClass.
type DataInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamClass that represents the DataInputStreamClass.
func (recv *DataInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamClassNewFromC creates a new DataInputStreamClass from a pointer to the C GDataInputStreamClass that represents the DataInputStreamClass.
func DataInputStreamClassNewFromC(native unsafe.Pointer) *DataInputStreamClass {
	return &DataInputStreamClass{native: native}
}

// DataInputStreamPrivate is a representation of the C record GDataInputStreamPrivate.
type DataInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStreamPrivate that represents the DataInputStreamPrivate.
func (recv *DataInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamPrivateNewFromC creates a new DataInputStreamPrivate from a pointer to the C GDataInputStreamPrivate that represents the DataInputStreamPrivate.
func DataInputStreamPrivateNewFromC(native unsafe.Pointer) *DataInputStreamPrivate {
	return &DataInputStreamPrivate{native: native}
}

// DataOutputStreamClass is a representation of the C record GDataOutputStreamClass.
type DataOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamClass that represents the DataOutputStreamClass.
func (recv *DataOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamClassNewFromC creates a new DataOutputStreamClass from a pointer to the C GDataOutputStreamClass that represents the DataOutputStreamClass.
func DataOutputStreamClassNewFromC(native unsafe.Pointer) *DataOutputStreamClass {
	return &DataOutputStreamClass{native: native}
}

// DataOutputStreamPrivate is a representation of the C record GDataOutputStreamPrivate.
type DataOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStreamPrivate that represents the DataOutputStreamPrivate.
func (recv *DataOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamPrivateNewFromC creates a new DataOutputStreamPrivate from a pointer to the C GDataOutputStreamPrivate that represents the DataOutputStreamPrivate.
func DataOutputStreamPrivateNewFromC(native unsafe.Pointer) *DataOutputStreamPrivate {
	return &DataOutputStreamPrivate{native: native}
}

// DesktopAppInfoClass is a representation of the C record GDesktopAppInfoClass.
type DesktopAppInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoClass that represents the DesktopAppInfoClass.
func (recv *DesktopAppInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoClassNewFromC creates a new DesktopAppInfoClass from a pointer to the C GDesktopAppInfoClass that represents the DesktopAppInfoClass.
func DesktopAppInfoClassNewFromC(native unsafe.Pointer) *DesktopAppInfoClass {
	return &DesktopAppInfoClass{native: native}
}

// DesktopAppInfoLookupIface is a representation of the C record GDesktopAppInfoLookupIface.
type DesktopAppInfoLookupIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookupIface that represents the DesktopAppInfoLookupIface.
func (recv *DesktopAppInfoLookupIface) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookupIfaceNewFromC creates a new DesktopAppInfoLookupIface from a pointer to the C GDesktopAppInfoLookupIface that represents the DesktopAppInfoLookupIface.
func DesktopAppInfoLookupIfaceNewFromC(native unsafe.Pointer) *DesktopAppInfoLookupIface {
	return &DesktopAppInfoLookupIface{native: native}
}

// DriveIface is a representation of the C record GDriveIface.
type DriveIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDriveIface that represents the DriveIface.
func (recv *DriveIface) ToC() unsafe.Pointer {
	return recv.native
}

// DriveIfaceNewFromC creates a new DriveIface from a pointer to the C GDriveIface that represents the DriveIface.
func DriveIfaceNewFromC(native unsafe.Pointer) *DriveIface {
	return &DriveIface{native: native}
}

// EmblemClass is a representation of the C record GEmblemClass.
type EmblemClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemClass that represents the EmblemClass.
func (recv *EmblemClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemClassNewFromC creates a new EmblemClass from a pointer to the C GEmblemClass that represents the EmblemClass.
func EmblemClassNewFromC(native unsafe.Pointer) *EmblemClass {
	return &EmblemClass{native: native}
}

// EmblemedIconClass is a representation of the C record GEmblemedIconClass.
type EmblemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconClass that represents the EmblemedIconClass.
func (recv *EmblemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconClassNewFromC creates a new EmblemedIconClass from a pointer to the C GEmblemedIconClass that represents the EmblemedIconClass.
func EmblemedIconClassNewFromC(native unsafe.Pointer) *EmblemedIconClass {
	return &EmblemedIconClass{native: native}
}

// EmblemedIconPrivate is a representation of the C record GEmblemedIconPrivate.
type EmblemedIconPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIconPrivate that represents the EmblemedIconPrivate.
func (recv *EmblemedIconPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconPrivateNewFromC creates a new EmblemedIconPrivate from a pointer to the C GEmblemedIconPrivate that represents the EmblemedIconPrivate.
func EmblemedIconPrivateNewFromC(native unsafe.Pointer) *EmblemedIconPrivate {
	return &EmblemedIconPrivate{native: native}
}

// FileAttributeInfo is a representation of the C record GFileAttributeInfo.
type FileAttributeInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfo that represents the FileAttributeInfo.
func (recv *FileAttributeInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfoNewFromC creates a new FileAttributeInfo from a pointer to the C GFileAttributeInfo that represents the FileAttributeInfo.
func FileAttributeInfoNewFromC(native unsafe.Pointer) *FileAttributeInfo {
	return &FileAttributeInfo{native: native}
}

// FileAttributeInfoList is a representation of the C record GFileAttributeInfoList.
type FileAttributeInfoList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeInfoList that represents the FileAttributeInfoList.
func (recv *FileAttributeInfoList) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeInfoListNewFromC creates a new FileAttributeInfoList from a pointer to the C GFileAttributeInfoList that represents the FileAttributeInfoList.
func FileAttributeInfoListNewFromC(native unsafe.Pointer) *FileAttributeInfoList {
	return &FileAttributeInfoList{native: native}
}

// FileAttributeMatcher is a representation of the C record GFileAttributeMatcher.
type FileAttributeMatcher struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileAttributeMatcher that represents the FileAttributeMatcher.
func (recv *FileAttributeMatcher) ToC() unsafe.Pointer {
	return recv.native
}

// FileAttributeMatcherNewFromC creates a new FileAttributeMatcher from a pointer to the C GFileAttributeMatcher that represents the FileAttributeMatcher.
func FileAttributeMatcherNewFromC(native unsafe.Pointer) *FileAttributeMatcher {
	return &FileAttributeMatcher{native: native}
}

// FileDescriptorBasedIface is a representation of the C record GFileDescriptorBasedIface.
type FileDescriptorBasedIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileDescriptorBasedIface that represents the FileDescriptorBasedIface.
func (recv *FileDescriptorBasedIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileDescriptorBasedIfaceNewFromC creates a new FileDescriptorBasedIface from a pointer to the C GFileDescriptorBasedIface that represents the FileDescriptorBasedIface.
func FileDescriptorBasedIfaceNewFromC(native unsafe.Pointer) *FileDescriptorBasedIface {
	return &FileDescriptorBasedIface{native: native}
}

// FileEnumeratorClass is a representation of the C record GFileEnumeratorClass.
type FileEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorClass that represents the FileEnumeratorClass.
func (recv *FileEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorClassNewFromC creates a new FileEnumeratorClass from a pointer to the C GFileEnumeratorClass that represents the FileEnumeratorClass.
func FileEnumeratorClassNewFromC(native unsafe.Pointer) *FileEnumeratorClass {
	return &FileEnumeratorClass{native: native}
}

// FileEnumeratorPrivate is a representation of the C record GFileEnumeratorPrivate.
type FileEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumeratorPrivate that represents the FileEnumeratorPrivate.
func (recv *FileEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorPrivateNewFromC creates a new FileEnumeratorPrivate from a pointer to the C GFileEnumeratorPrivate that represents the FileEnumeratorPrivate.
func FileEnumeratorPrivateNewFromC(native unsafe.Pointer) *FileEnumeratorPrivate {
	return &FileEnumeratorPrivate{native: native}
}

// FileIOStreamClass is a representation of the C record GFileIOStreamClass.
type FileIOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamClass that represents the FileIOStreamClass.
func (recv *FileIOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamClassNewFromC creates a new FileIOStreamClass from a pointer to the C GFileIOStreamClass that represents the FileIOStreamClass.
func FileIOStreamClassNewFromC(native unsafe.Pointer) *FileIOStreamClass {
	return &FileIOStreamClass{native: native}
}

// FileIOStreamPrivate is a representation of the C record GFileIOStreamPrivate.
type FileIOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIOStreamPrivate that represents the FileIOStreamPrivate.
func (recv *FileIOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileIOStreamPrivateNewFromC creates a new FileIOStreamPrivate from a pointer to the C GFileIOStreamPrivate that represents the FileIOStreamPrivate.
func FileIOStreamPrivateNewFromC(native unsafe.Pointer) *FileIOStreamPrivate {
	return &FileIOStreamPrivate{native: native}
}

// FileIconClass is a representation of the C record GFileIconClass.
type FileIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIconClass that represents the FileIconClass.
func (recv *FileIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileIconClassNewFromC creates a new FileIconClass from a pointer to the C GFileIconClass that represents the FileIconClass.
func FileIconClassNewFromC(native unsafe.Pointer) *FileIconClass {
	return &FileIconClass{native: native}
}

// FileIface is a representation of the C record GFileIface.
type FileIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIface that represents the FileIface.
func (recv *FileIface) ToC() unsafe.Pointer {
	return recv.native
}

// FileIfaceNewFromC creates a new FileIface from a pointer to the C GFileIface that represents the FileIface.
func FileIfaceNewFromC(native unsafe.Pointer) *FileIface {
	return &FileIface{native: native}
}

// FileInfoClass is a representation of the C record GFileInfoClass.
type FileInfoClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfoClass that represents the FileInfoClass.
func (recv *FileInfoClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfoClassNewFromC creates a new FileInfoClass from a pointer to the C GFileInfoClass that represents the FileInfoClass.
func FileInfoClassNewFromC(native unsafe.Pointer) *FileInfoClass {
	return &FileInfoClass{native: native}
}

// FileInputStreamClass is a representation of the C record GFileInputStreamClass.
type FileInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamClass that represents the FileInputStreamClass.
func (recv *FileInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamClassNewFromC creates a new FileInputStreamClass from a pointer to the C GFileInputStreamClass that represents the FileInputStreamClass.
func FileInputStreamClassNewFromC(native unsafe.Pointer) *FileInputStreamClass {
	return &FileInputStreamClass{native: native}
}

// FileInputStreamPrivate is a representation of the C record GFileInputStreamPrivate.
type FileInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStreamPrivate that represents the FileInputStreamPrivate.
func (recv *FileInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamPrivateNewFromC creates a new FileInputStreamPrivate from a pointer to the C GFileInputStreamPrivate that represents the FileInputStreamPrivate.
func FileInputStreamPrivateNewFromC(native unsafe.Pointer) *FileInputStreamPrivate {
	return &FileInputStreamPrivate{native: native}
}

// FileMonitorClass is a representation of the C record GFileMonitorClass.
type FileMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorClass that represents the FileMonitorClass.
func (recv *FileMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorClassNewFromC creates a new FileMonitorClass from a pointer to the C GFileMonitorClass that represents the FileMonitorClass.
func FileMonitorClassNewFromC(native unsafe.Pointer) *FileMonitorClass {
	return &FileMonitorClass{native: native}
}

// FileMonitorPrivate is a representation of the C record GFileMonitorPrivate.
type FileMonitorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitorPrivate that represents the FileMonitorPrivate.
func (recv *FileMonitorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorPrivateNewFromC creates a new FileMonitorPrivate from a pointer to the C GFileMonitorPrivate that represents the FileMonitorPrivate.
func FileMonitorPrivateNewFromC(native unsafe.Pointer) *FileMonitorPrivate {
	return &FileMonitorPrivate{native: native}
}

// FileOutputStreamClass is a representation of the C record GFileOutputStreamClass.
type FileOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamClass that represents the FileOutputStreamClass.
func (recv *FileOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamClassNewFromC creates a new FileOutputStreamClass from a pointer to the C GFileOutputStreamClass that represents the FileOutputStreamClass.
func FileOutputStreamClassNewFromC(native unsafe.Pointer) *FileOutputStreamClass {
	return &FileOutputStreamClass{native: native}
}

// FileOutputStreamPrivate is a representation of the C record GFileOutputStreamPrivate.
type FileOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStreamPrivate that represents the FileOutputStreamPrivate.
func (recv *FileOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamPrivateNewFromC creates a new FileOutputStreamPrivate from a pointer to the C GFileOutputStreamPrivate that represents the FileOutputStreamPrivate.
func FileOutputStreamPrivateNewFromC(native unsafe.Pointer) *FileOutputStreamPrivate {
	return &FileOutputStreamPrivate{native: native}
}

// FilenameCompleterClass is a representation of the C record GFilenameCompleterClass.
type FilenameCompleterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleterClass that represents the FilenameCompleterClass.
func (recv *FilenameCompleterClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleterClassNewFromC creates a new FilenameCompleterClass from a pointer to the C GFilenameCompleterClass that represents the FilenameCompleterClass.
func FilenameCompleterClassNewFromC(native unsafe.Pointer) *FilenameCompleterClass {
	return &FilenameCompleterClass{native: native}
}

// FilterInputStreamClass is a representation of the C record GFilterInputStreamClass.
type FilterInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStreamClass that represents the FilterInputStreamClass.
func (recv *FilterInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStreamClassNewFromC creates a new FilterInputStreamClass from a pointer to the C GFilterInputStreamClass that represents the FilterInputStreamClass.
func FilterInputStreamClassNewFromC(native unsafe.Pointer) *FilterInputStreamClass {
	return &FilterInputStreamClass{native: native}
}

// FilterOutputStreamClass is a representation of the C record GFilterOutputStreamClass.
type FilterOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStreamClass that represents the FilterOutputStreamClass.
func (recv *FilterOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStreamClassNewFromC creates a new FilterOutputStreamClass from a pointer to the C GFilterOutputStreamClass that represents the FilterOutputStreamClass.
func FilterOutputStreamClassNewFromC(native unsafe.Pointer) *FilterOutputStreamClass {
	return &FilterOutputStreamClass{native: native}
}

// IOExtension is a representation of the C record GIOExtension.
type IOExtension struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtension that represents the IOExtension.
func (recv *IOExtension) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtensionNewFromC creates a new IOExtension from a pointer to the C GIOExtension that represents the IOExtension.
func IOExtensionNewFromC(native unsafe.Pointer) *IOExtension {
	return &IOExtension{native: native}
}

// IOExtensionPoint is a representation of the C record GIOExtensionPoint.
type IOExtensionPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOExtensionPoint that represents the IOExtensionPoint.
func (recv *IOExtensionPoint) ToC() unsafe.Pointer {
	return recv.native
}

// IOExtensionPointNewFromC creates a new IOExtensionPoint from a pointer to the C GIOExtensionPoint that represents the IOExtensionPoint.
func IOExtensionPointNewFromC(native unsafe.Pointer) *IOExtensionPoint {
	return &IOExtensionPoint{native: native}
}

// IOModuleClass is a representation of the C record GIOModuleClass.
type IOModuleClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModuleClass that represents the IOModuleClass.
func (recv *IOModuleClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleClassNewFromC creates a new IOModuleClass from a pointer to the C GIOModuleClass that represents the IOModuleClass.
func IOModuleClassNewFromC(native unsafe.Pointer) *IOModuleClass {
	return &IOModuleClass{native: native}
}

// IOSchedulerJob is a representation of the C record GIOSchedulerJob.
type IOSchedulerJob struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOSchedulerJob that represents the IOSchedulerJob.
func (recv *IOSchedulerJob) ToC() unsafe.Pointer {
	return recv.native
}

// IOSchedulerJobNewFromC creates a new IOSchedulerJob from a pointer to the C GIOSchedulerJob that represents the IOSchedulerJob.
func IOSchedulerJobNewFromC(native unsafe.Pointer) *IOSchedulerJob {
	return &IOSchedulerJob{native: native}
}

// IOStreamAdapter is a representation of the C record GIOStreamAdapter.
type IOStreamAdapter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamAdapter that represents the IOStreamAdapter.
func (recv *IOStreamAdapter) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamAdapterNewFromC creates a new IOStreamAdapter from a pointer to the C GIOStreamAdapter that represents the IOStreamAdapter.
func IOStreamAdapterNewFromC(native unsafe.Pointer) *IOStreamAdapter {
	return &IOStreamAdapter{native: native}
}

// IOStreamClass is a representation of the C record GIOStreamClass.
type IOStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamClass that represents the IOStreamClass.
func (recv *IOStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamClassNewFromC creates a new IOStreamClass from a pointer to the C GIOStreamClass that represents the IOStreamClass.
func IOStreamClassNewFromC(native unsafe.Pointer) *IOStreamClass {
	return &IOStreamClass{native: native}
}

// IOStreamPrivate is a representation of the C record GIOStreamPrivate.
type IOStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOStreamPrivate that represents the IOStreamPrivate.
func (recv *IOStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// IOStreamPrivateNewFromC creates a new IOStreamPrivate from a pointer to the C GIOStreamPrivate that represents the IOStreamPrivate.
func IOStreamPrivateNewFromC(native unsafe.Pointer) *IOStreamPrivate {
	return &IOStreamPrivate{native: native}
}

// IconIface is a representation of the C record GIconIface.
type IconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIconIface that represents the IconIface.
func (recv *IconIface) ToC() unsafe.Pointer {
	return recv.native
}

// IconIfaceNewFromC creates a new IconIface from a pointer to the C GIconIface that represents the IconIface.
func IconIfaceNewFromC(native unsafe.Pointer) *IconIface {
	return &IconIface{native: native}
}

// InetAddressClass is a representation of the C record GInetAddressClass.
type InetAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressClass that represents the InetAddressClass.
func (recv *InetAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressClassNewFromC creates a new InetAddressClass from a pointer to the C GInetAddressClass that represents the InetAddressClass.
func InetAddressClassNewFromC(native unsafe.Pointer) *InetAddressClass {
	return &InetAddressClass{native: native}
}

// InetAddressMaskClass is a representation of the C record GInetAddressMaskClass.
type InetAddressMaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskClass that represents the InetAddressMaskClass.
func (recv *InetAddressMaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskClassNewFromC creates a new InetAddressMaskClass from a pointer to the C GInetAddressMaskClass that represents the InetAddressMaskClass.
func InetAddressMaskClassNewFromC(native unsafe.Pointer) *InetAddressMaskClass {
	return &InetAddressMaskClass{native: native}
}

// InetAddressMaskPrivate is a representation of the C record GInetAddressMaskPrivate.
type InetAddressMaskPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressMaskPrivate that represents the InetAddressMaskPrivate.
func (recv *InetAddressMaskPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressMaskPrivateNewFromC creates a new InetAddressMaskPrivate from a pointer to the C GInetAddressMaskPrivate that represents the InetAddressMaskPrivate.
func InetAddressMaskPrivateNewFromC(native unsafe.Pointer) *InetAddressMaskPrivate {
	return &InetAddressMaskPrivate{native: native}
}

// InetAddressPrivate is a representation of the C record GInetAddressPrivate.
type InetAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddressPrivate that represents the InetAddressPrivate.
func (recv *InetAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressPrivateNewFromC creates a new InetAddressPrivate from a pointer to the C GInetAddressPrivate that represents the InetAddressPrivate.
func InetAddressPrivateNewFromC(native unsafe.Pointer) *InetAddressPrivate {
	return &InetAddressPrivate{native: native}
}

// InetSocketAddressClass is a representation of the C record GInetSocketAddressClass.
type InetSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressClass that represents the InetSocketAddressClass.
func (recv *InetSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressClassNewFromC creates a new InetSocketAddressClass from a pointer to the C GInetSocketAddressClass that represents the InetSocketAddressClass.
func InetSocketAddressClassNewFromC(native unsafe.Pointer) *InetSocketAddressClass {
	return &InetSocketAddressClass{native: native}
}

// InetSocketAddressPrivate is a representation of the C record GInetSocketAddressPrivate.
type InetSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddressPrivate that represents the InetSocketAddressPrivate.
func (recv *InetSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressPrivateNewFromC creates a new InetSocketAddressPrivate from a pointer to the C GInetSocketAddressPrivate that represents the InetSocketAddressPrivate.
func InetSocketAddressPrivateNewFromC(native unsafe.Pointer) *InetSocketAddressPrivate {
	return &InetSocketAddressPrivate{native: native}
}

// InputStreamClass is a representation of the C record GInputStreamClass.
type InputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamClass that represents the InputStreamClass.
func (recv *InputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamClassNewFromC creates a new InputStreamClass from a pointer to the C GInputStreamClass that represents the InputStreamClass.
func InputStreamClassNewFromC(native unsafe.Pointer) *InputStreamClass {
	return &InputStreamClass{native: native}
}

// InputStreamPrivate is a representation of the C record GInputStreamPrivate.
type InputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStreamPrivate that represents the InputStreamPrivate.
func (recv *InputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamPrivateNewFromC creates a new InputStreamPrivate from a pointer to the C GInputStreamPrivate that represents the InputStreamPrivate.
func InputStreamPrivateNewFromC(native unsafe.Pointer) *InputStreamPrivate {
	return &InputStreamPrivate{native: native}
}

// ListStoreClass is a representation of the C record GListStoreClass.
type ListStoreClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStoreClass that represents the ListStoreClass.
func (recv *ListStoreClass) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreClassNewFromC creates a new ListStoreClass from a pointer to the C GListStoreClass that represents the ListStoreClass.
func ListStoreClassNewFromC(native unsafe.Pointer) *ListStoreClass {
	return &ListStoreClass{native: native}
}

// LoadableIconIface is a representation of the C record GLoadableIconIface.
type LoadableIconIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIconIface that represents the LoadableIconIface.
func (recv *LoadableIconIface) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIconIfaceNewFromC creates a new LoadableIconIface from a pointer to the C GLoadableIconIface that represents the LoadableIconIface.
func LoadableIconIfaceNewFromC(native unsafe.Pointer) *LoadableIconIface {
	return &LoadableIconIface{native: native}
}

// MemoryInputStreamClass is a representation of the C record GMemoryInputStreamClass.
type MemoryInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamClass that represents the MemoryInputStreamClass.
func (recv *MemoryInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamClassNewFromC creates a new MemoryInputStreamClass from a pointer to the C GMemoryInputStreamClass that represents the MemoryInputStreamClass.
func MemoryInputStreamClassNewFromC(native unsafe.Pointer) *MemoryInputStreamClass {
	return &MemoryInputStreamClass{native: native}
}

// MemoryInputStreamPrivate is a representation of the C record GMemoryInputStreamPrivate.
type MemoryInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStreamPrivate that represents the MemoryInputStreamPrivate.
func (recv *MemoryInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamPrivateNewFromC creates a new MemoryInputStreamPrivate from a pointer to the C GMemoryInputStreamPrivate that represents the MemoryInputStreamPrivate.
func MemoryInputStreamPrivateNewFromC(native unsafe.Pointer) *MemoryInputStreamPrivate {
	return &MemoryInputStreamPrivate{native: native}
}

// MemoryOutputStreamClass is a representation of the C record GMemoryOutputStreamClass.
type MemoryOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamClass that represents the MemoryOutputStreamClass.
func (recv *MemoryOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamClassNewFromC creates a new MemoryOutputStreamClass from a pointer to the C GMemoryOutputStreamClass that represents the MemoryOutputStreamClass.
func MemoryOutputStreamClassNewFromC(native unsafe.Pointer) *MemoryOutputStreamClass {
	return &MemoryOutputStreamClass{native: native}
}

// MemoryOutputStreamPrivate is a representation of the C record GMemoryOutputStreamPrivate.
type MemoryOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStreamPrivate that represents the MemoryOutputStreamPrivate.
func (recv *MemoryOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamPrivateNewFromC creates a new MemoryOutputStreamPrivate from a pointer to the C GMemoryOutputStreamPrivate that represents the MemoryOutputStreamPrivate.
func MemoryOutputStreamPrivateNewFromC(native unsafe.Pointer) *MemoryOutputStreamPrivate {
	return &MemoryOutputStreamPrivate{native: native}
}

// MenuAttributeIterClass is a representation of the C record GMenuAttributeIterClass.
type MenuAttributeIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterClass that represents the MenuAttributeIterClass.
func (recv *MenuAttributeIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterClassNewFromC creates a new MenuAttributeIterClass from a pointer to the C GMenuAttributeIterClass that represents the MenuAttributeIterClass.
func MenuAttributeIterClassNewFromC(native unsafe.Pointer) *MenuAttributeIterClass {
	return &MenuAttributeIterClass{native: native}
}

// MenuAttributeIterPrivate is a representation of the C record GMenuAttributeIterPrivate.
type MenuAttributeIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuAttributeIterPrivate that represents the MenuAttributeIterPrivate.
func (recv *MenuAttributeIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuAttributeIterPrivateNewFromC creates a new MenuAttributeIterPrivate from a pointer to the C GMenuAttributeIterPrivate that represents the MenuAttributeIterPrivate.
func MenuAttributeIterPrivateNewFromC(native unsafe.Pointer) *MenuAttributeIterPrivate {
	return &MenuAttributeIterPrivate{native: native}
}

// MenuLinkIterClass is a representation of the C record GMenuLinkIterClass.
type MenuLinkIterClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterClass that represents the MenuLinkIterClass.
func (recv *MenuLinkIterClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterClassNewFromC creates a new MenuLinkIterClass from a pointer to the C GMenuLinkIterClass that represents the MenuLinkIterClass.
func MenuLinkIterClassNewFromC(native unsafe.Pointer) *MenuLinkIterClass {
	return &MenuLinkIterClass{native: native}
}

// MenuLinkIterPrivate is a representation of the C record GMenuLinkIterPrivate.
type MenuLinkIterPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuLinkIterPrivate that represents the MenuLinkIterPrivate.
func (recv *MenuLinkIterPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuLinkIterPrivateNewFromC creates a new MenuLinkIterPrivate from a pointer to the C GMenuLinkIterPrivate that represents the MenuLinkIterPrivate.
func MenuLinkIterPrivateNewFromC(native unsafe.Pointer) *MenuLinkIterPrivate {
	return &MenuLinkIterPrivate{native: native}
}

// MenuModelClass is a representation of the C record GMenuModelClass.
type MenuModelClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelClass that represents the MenuModelClass.
func (recv *MenuModelClass) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelClassNewFromC creates a new MenuModelClass from a pointer to the C GMenuModelClass that represents the MenuModelClass.
func MenuModelClassNewFromC(native unsafe.Pointer) *MenuModelClass {
	return &MenuModelClass{native: native}
}

// MenuModelPrivate is a representation of the C record GMenuModelPrivate.
type MenuModelPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMenuModelPrivate that represents the MenuModelPrivate.
func (recv *MenuModelPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MenuModelPrivateNewFromC creates a new MenuModelPrivate from a pointer to the C GMenuModelPrivate that represents the MenuModelPrivate.
func MenuModelPrivateNewFromC(native unsafe.Pointer) *MenuModelPrivate {
	return &MenuModelPrivate{native: native}
}

// MountIface is a representation of the C record GMountIface.
type MountIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountIface that represents the MountIface.
func (recv *MountIface) ToC() unsafe.Pointer {
	return recv.native
}

// MountIfaceNewFromC creates a new MountIface from a pointer to the C GMountIface that represents the MountIface.
func MountIfaceNewFromC(native unsafe.Pointer) *MountIface {
	return &MountIface{native: native}
}

// MountOperationClass is a representation of the C record GMountOperationClass.
type MountOperationClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationClass that represents the MountOperationClass.
func (recv *MountOperationClass) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationClassNewFromC creates a new MountOperationClass from a pointer to the C GMountOperationClass that represents the MountOperationClass.
func MountOperationClassNewFromC(native unsafe.Pointer) *MountOperationClass {
	return &MountOperationClass{native: native}
}

// MountOperationPrivate is a representation of the C record GMountOperationPrivate.
type MountOperationPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperationPrivate that represents the MountOperationPrivate.
func (recv *MountOperationPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationPrivateNewFromC creates a new MountOperationPrivate from a pointer to the C GMountOperationPrivate that represents the MountOperationPrivate.
func MountOperationPrivateNewFromC(native unsafe.Pointer) *MountOperationPrivate {
	return &MountOperationPrivate{native: native}
}

// UNSUPPORTED : NativeSocketAddressClass : blacklisted

// UNSUPPORTED : NativeSocketAddressPrivate : blacklisted

// NativeVolumeMonitorClass is a representation of the C record GNativeVolumeMonitorClass.
type NativeVolumeMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeVolumeMonitorClass that represents the NativeVolumeMonitorClass.
func (recv *NativeVolumeMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// NativeVolumeMonitorClassNewFromC creates a new NativeVolumeMonitorClass from a pointer to the C GNativeVolumeMonitorClass that represents the NativeVolumeMonitorClass.
func NativeVolumeMonitorClassNewFromC(native unsafe.Pointer) *NativeVolumeMonitorClass {
	return &NativeVolumeMonitorClass{native: native}
}

// NetworkAddressClass is a representation of the C record GNetworkAddressClass.
type NetworkAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressClass that represents the NetworkAddressClass.
func (recv *NetworkAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressClassNewFromC creates a new NetworkAddressClass from a pointer to the C GNetworkAddressClass that represents the NetworkAddressClass.
func NetworkAddressClassNewFromC(native unsafe.Pointer) *NetworkAddressClass {
	return &NetworkAddressClass{native: native}
}

// NetworkAddressPrivate is a representation of the C record GNetworkAddressPrivate.
type NetworkAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddressPrivate that represents the NetworkAddressPrivate.
func (recv *NetworkAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressPrivateNewFromC creates a new NetworkAddressPrivate from a pointer to the C GNetworkAddressPrivate that represents the NetworkAddressPrivate.
func NetworkAddressPrivateNewFromC(native unsafe.Pointer) *NetworkAddressPrivate {
	return &NetworkAddressPrivate{native: native}
}

// NetworkServiceClass is a representation of the C record GNetworkServiceClass.
type NetworkServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServiceClass that represents the NetworkServiceClass.
func (recv *NetworkServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServiceClassNewFromC creates a new NetworkServiceClass from a pointer to the C GNetworkServiceClass that represents the NetworkServiceClass.
func NetworkServiceClassNewFromC(native unsafe.Pointer) *NetworkServiceClass {
	return &NetworkServiceClass{native: native}
}

// NetworkServicePrivate is a representation of the C record GNetworkServicePrivate.
type NetworkServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkServicePrivate that represents the NetworkServicePrivate.
func (recv *NetworkServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServicePrivateNewFromC creates a new NetworkServicePrivate from a pointer to the C GNetworkServicePrivate that represents the NetworkServicePrivate.
func NetworkServicePrivateNewFromC(native unsafe.Pointer) *NetworkServicePrivate {
	return &NetworkServicePrivate{native: native}
}

// OutputStreamClass is a representation of the C record GOutputStreamClass.
type OutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamClass that represents the OutputStreamClass.
func (recv *OutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamClassNewFromC creates a new OutputStreamClass from a pointer to the C GOutputStreamClass that represents the OutputStreamClass.
func OutputStreamClassNewFromC(native unsafe.Pointer) *OutputStreamClass {
	return &OutputStreamClass{native: native}
}

// OutputStreamPrivate is a representation of the C record GOutputStreamPrivate.
type OutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStreamPrivate that represents the OutputStreamPrivate.
func (recv *OutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamPrivateNewFromC creates a new OutputStreamPrivate from a pointer to the C GOutputStreamPrivate that represents the OutputStreamPrivate.
func OutputStreamPrivateNewFromC(native unsafe.Pointer) *OutputStreamPrivate {
	return &OutputStreamPrivate{native: native}
}

// PermissionClass is a representation of the C record GPermissionClass.
type PermissionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionClass that represents the PermissionClass.
func (recv *PermissionClass) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionClassNewFromC creates a new PermissionClass from a pointer to the C GPermissionClass that represents the PermissionClass.
func PermissionClassNewFromC(native unsafe.Pointer) *PermissionClass {
	return &PermissionClass{native: native}
}

// PermissionPrivate is a representation of the C record GPermissionPrivate.
type PermissionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermissionPrivate that represents the PermissionPrivate.
func (recv *PermissionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionPrivateNewFromC creates a new PermissionPrivate from a pointer to the C GPermissionPrivate that represents the PermissionPrivate.
func PermissionPrivateNewFromC(native unsafe.Pointer) *PermissionPrivate {
	return &PermissionPrivate{native: native}
}

// ProxyAddressEnumeratorClass is a representation of the C record GProxyAddressEnumeratorClass.
type ProxyAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorClass that represents the ProxyAddressEnumeratorClass.
func (recv *ProxyAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorClassNewFromC creates a new ProxyAddressEnumeratorClass from a pointer to the C GProxyAddressEnumeratorClass that represents the ProxyAddressEnumeratorClass.
func ProxyAddressEnumeratorClassNewFromC(native unsafe.Pointer) *ProxyAddressEnumeratorClass {
	return &ProxyAddressEnumeratorClass{native: native}
}

// ProxyAddressEnumeratorPrivate is a representation of the C record GProxyAddressEnumeratorPrivate.
type ProxyAddressEnumeratorPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumeratorPrivate that represents the ProxyAddressEnumeratorPrivate.
func (recv *ProxyAddressEnumeratorPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorPrivateNewFromC creates a new ProxyAddressEnumeratorPrivate from a pointer to the C GProxyAddressEnumeratorPrivate that represents the ProxyAddressEnumeratorPrivate.
func ProxyAddressEnumeratorPrivateNewFromC(native unsafe.Pointer) *ProxyAddressEnumeratorPrivate {
	return &ProxyAddressEnumeratorPrivate{native: native}
}

// ProxyAddressPrivate is a representation of the C record GProxyAddressPrivate.
type ProxyAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressPrivate that represents the ProxyAddressPrivate.
func (recv *ProxyAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressPrivateNewFromC creates a new ProxyAddressPrivate from a pointer to the C GProxyAddressPrivate that represents the ProxyAddressPrivate.
func ProxyAddressPrivateNewFromC(native unsafe.Pointer) *ProxyAddressPrivate {
	return &ProxyAddressPrivate{native: native}
}

// ProxyResolverInterface is a representation of the C record GProxyResolverInterface.
type ProxyResolverInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyResolverInterface that represents the ProxyResolverInterface.
func (recv *ProxyResolverInterface) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyResolverInterfaceNewFromC creates a new ProxyResolverInterface from a pointer to the C GProxyResolverInterface that represents the ProxyResolverInterface.
func ProxyResolverInterfaceNewFromC(native unsafe.Pointer) *ProxyResolverInterface {
	return &ProxyResolverInterface{native: native}
}

// ResolverClass is a representation of the C record GResolverClass.
type ResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverClass that represents the ResolverClass.
func (recv *ResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverClassNewFromC creates a new ResolverClass from a pointer to the C GResolverClass that represents the ResolverClass.
func ResolverClassNewFromC(native unsafe.Pointer) *ResolverClass {
	return &ResolverClass{native: native}
}

// ResolverPrivate is a representation of the C record GResolverPrivate.
type ResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolverPrivate that represents the ResolverPrivate.
func (recv *ResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverPrivateNewFromC creates a new ResolverPrivate from a pointer to the C GResolverPrivate that represents the ResolverPrivate.
func ResolverPrivateNewFromC(native unsafe.Pointer) *ResolverPrivate {
	return &ResolverPrivate{native: native}
}

// SeekableIface is a representation of the C record GSeekableIface.
type SeekableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekableIface that represents the SeekableIface.
func (recv *SeekableIface) ToC() unsafe.Pointer {
	return recv.native
}

// SeekableIfaceNewFromC creates a new SeekableIface from a pointer to the C GSeekableIface that represents the SeekableIface.
func SeekableIfaceNewFromC(native unsafe.Pointer) *SeekableIface {
	return &SeekableIface{native: native}
}

// UNSUPPORTED : SettingsBackendClass : blacklisted

// UNSUPPORTED : SettingsBackendPrivate : blacklisted

// SettingsClass is a representation of the C record GSettingsClass.
type SettingsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsClass that represents the SettingsClass.
func (recv *SettingsClass) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsClassNewFromC creates a new SettingsClass from a pointer to the C GSettingsClass that represents the SettingsClass.
func SettingsClassNewFromC(native unsafe.Pointer) *SettingsClass {
	return &SettingsClass{native: native}
}

// SettingsPrivate is a representation of the C record GSettingsPrivate.
type SettingsPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsPrivate that represents the SettingsPrivate.
func (recv *SettingsPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsPrivateNewFromC creates a new SettingsPrivate from a pointer to the C GSettingsPrivate that represents the SettingsPrivate.
func SettingsPrivateNewFromC(native unsafe.Pointer) *SettingsPrivate {
	return &SettingsPrivate{native: native}
}

// SettingsSchemaKey is a representation of the C record GSettingsSchemaKey.
type SettingsSchemaKey struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsSchemaKey that represents the SettingsSchemaKey.
func (recv *SettingsSchemaKey) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsSchemaKeyNewFromC creates a new SettingsSchemaKey from a pointer to the C GSettingsSchemaKey that represents the SettingsSchemaKey.
func SettingsSchemaKeyNewFromC(native unsafe.Pointer) *SettingsSchemaKey {
	return &SettingsSchemaKey{native: native}
}

// SimpleActionGroupClass is a representation of the C record GSimpleActionGroupClass.
type SimpleActionGroupClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupClass that represents the SimpleActionGroupClass.
func (recv *SimpleActionGroupClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupClassNewFromC creates a new SimpleActionGroupClass from a pointer to the C GSimpleActionGroupClass that represents the SimpleActionGroupClass.
func SimpleActionGroupClassNewFromC(native unsafe.Pointer) *SimpleActionGroupClass {
	return &SimpleActionGroupClass{native: native}
}

// SimpleActionGroupPrivate is a representation of the C record GSimpleActionGroupPrivate.
type SimpleActionGroupPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleActionGroupPrivate that represents the SimpleActionGroupPrivate.
func (recv *SimpleActionGroupPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionGroupPrivateNewFromC creates a new SimpleActionGroupPrivate from a pointer to the C GSimpleActionGroupPrivate that represents the SimpleActionGroupPrivate.
func SimpleActionGroupPrivateNewFromC(native unsafe.Pointer) *SimpleActionGroupPrivate {
	return &SimpleActionGroupPrivate{native: native}
}

// SimpleAsyncResultClass is a representation of the C record GSimpleAsyncResultClass.
type SimpleAsyncResultClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResultClass that represents the SimpleAsyncResultClass.
func (recv *SimpleAsyncResultClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultClassNewFromC creates a new SimpleAsyncResultClass from a pointer to the C GSimpleAsyncResultClass that represents the SimpleAsyncResultClass.
func SimpleAsyncResultClassNewFromC(native unsafe.Pointer) *SimpleAsyncResultClass {
	return &SimpleAsyncResultClass{native: native}
}

// SimpleProxyResolverClass is a representation of the C record GSimpleProxyResolverClass.
type SimpleProxyResolverClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverClass that represents the SimpleProxyResolverClass.
func (recv *SimpleProxyResolverClass) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverClassNewFromC creates a new SimpleProxyResolverClass from a pointer to the C GSimpleProxyResolverClass that represents the SimpleProxyResolverClass.
func SimpleProxyResolverClassNewFromC(native unsafe.Pointer) *SimpleProxyResolverClass {
	return &SimpleProxyResolverClass{native: native}
}

// SimpleProxyResolverPrivate is a representation of the C record GSimpleProxyResolverPrivate.
type SimpleProxyResolverPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleProxyResolverPrivate that represents the SimpleProxyResolverPrivate.
func (recv *SimpleProxyResolverPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleProxyResolverPrivateNewFromC creates a new SimpleProxyResolverPrivate from a pointer to the C GSimpleProxyResolverPrivate that represents the SimpleProxyResolverPrivate.
func SimpleProxyResolverPrivateNewFromC(native unsafe.Pointer) *SimpleProxyResolverPrivate {
	return &SimpleProxyResolverPrivate{native: native}
}

// SocketAddressClass is a representation of the C record GSocketAddressClass.
type SocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressClass that represents the SocketAddressClass.
func (recv *SocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressClassNewFromC creates a new SocketAddressClass from a pointer to the C GSocketAddressClass that represents the SocketAddressClass.
func SocketAddressClassNewFromC(native unsafe.Pointer) *SocketAddressClass {
	return &SocketAddressClass{native: native}
}

// SocketAddressEnumeratorClass is a representation of the C record GSocketAddressEnumeratorClass.
type SocketAddressEnumeratorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumeratorClass that represents the SocketAddressEnumeratorClass.
func (recv *SocketAddressEnumeratorClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumeratorClassNewFromC creates a new SocketAddressEnumeratorClass from a pointer to the C GSocketAddressEnumeratorClass that represents the SocketAddressEnumeratorClass.
func SocketAddressEnumeratorClassNewFromC(native unsafe.Pointer) *SocketAddressEnumeratorClass {
	return &SocketAddressEnumeratorClass{native: native}
}

// SocketClass is a representation of the C record GSocketClass.
type SocketClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClass that represents the SocketClass.
func (recv *SocketClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClassNewFromC creates a new SocketClass from a pointer to the C GSocketClass that represents the SocketClass.
func SocketClassNewFromC(native unsafe.Pointer) *SocketClass {
	return &SocketClass{native: native}
}

// SocketClientClass is a representation of the C record GSocketClientClass.
type SocketClientClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientClass that represents the SocketClientClass.
func (recv *SocketClientClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientClassNewFromC creates a new SocketClientClass from a pointer to the C GSocketClientClass that represents the SocketClientClass.
func SocketClientClassNewFromC(native unsafe.Pointer) *SocketClientClass {
	return &SocketClientClass{native: native}
}

// SocketClientPrivate is a representation of the C record GSocketClientPrivate.
type SocketClientPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketClientPrivate that represents the SocketClientPrivate.
func (recv *SocketClientPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketClientPrivateNewFromC creates a new SocketClientPrivate from a pointer to the C GSocketClientPrivate that represents the SocketClientPrivate.
func SocketClientPrivateNewFromC(native unsafe.Pointer) *SocketClientPrivate {
	return &SocketClientPrivate{native: native}
}

// SocketConnectableIface is a representation of the C record GSocketConnectableIface.
type SocketConnectableIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectableIface that represents the SocketConnectableIface.
func (recv *SocketConnectableIface) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectableIfaceNewFromC creates a new SocketConnectableIface from a pointer to the C GSocketConnectableIface that represents the SocketConnectableIface.
func SocketConnectableIfaceNewFromC(native unsafe.Pointer) *SocketConnectableIface {
	return &SocketConnectableIface{native: native}
}

// SocketConnectionClass is a representation of the C record GSocketConnectionClass.
type SocketConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionClass that represents the SocketConnectionClass.
func (recv *SocketConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionClassNewFromC creates a new SocketConnectionClass from a pointer to the C GSocketConnectionClass that represents the SocketConnectionClass.
func SocketConnectionClassNewFromC(native unsafe.Pointer) *SocketConnectionClass {
	return &SocketConnectionClass{native: native}
}

// SocketConnectionPrivate is a representation of the C record GSocketConnectionPrivate.
type SocketConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectionPrivate that represents the SocketConnectionPrivate.
func (recv *SocketConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectionPrivateNewFromC creates a new SocketConnectionPrivate from a pointer to the C GSocketConnectionPrivate that represents the SocketConnectionPrivate.
func SocketConnectionPrivateNewFromC(native unsafe.Pointer) *SocketConnectionPrivate {
	return &SocketConnectionPrivate{native: native}
}

// SocketControlMessageClass is a representation of the C record GSocketControlMessageClass.
type SocketControlMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessageClass that represents the SocketControlMessageClass.
func (recv *SocketControlMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessageClassNewFromC creates a new SocketControlMessageClass from a pointer to the C GSocketControlMessageClass that represents the SocketControlMessageClass.
func SocketControlMessageClassNewFromC(native unsafe.Pointer) *SocketControlMessageClass {
	return &SocketControlMessageClass{native: native}
}

// SocketControlMessagePrivate is a representation of the C record GSocketControlMessagePrivate.
type SocketControlMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketControlMessagePrivate that represents the SocketControlMessagePrivate.
func (recv *SocketControlMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketControlMessagePrivateNewFromC creates a new SocketControlMessagePrivate from a pointer to the C GSocketControlMessagePrivate that represents the SocketControlMessagePrivate.
func SocketControlMessagePrivateNewFromC(native unsafe.Pointer) *SocketControlMessagePrivate {
	return &SocketControlMessagePrivate{native: native}
}

// SocketListenerClass is a representation of the C record GSocketListenerClass.
type SocketListenerClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerClass that represents the SocketListenerClass.
func (recv *SocketListenerClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerClassNewFromC creates a new SocketListenerClass from a pointer to the C GSocketListenerClass that represents the SocketListenerClass.
func SocketListenerClassNewFromC(native unsafe.Pointer) *SocketListenerClass {
	return &SocketListenerClass{native: native}
}

// SocketListenerPrivate is a representation of the C record GSocketListenerPrivate.
type SocketListenerPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketListenerPrivate that represents the SocketListenerPrivate.
func (recv *SocketListenerPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketListenerPrivateNewFromC creates a new SocketListenerPrivate from a pointer to the C GSocketListenerPrivate that represents the SocketListenerPrivate.
func SocketListenerPrivateNewFromC(native unsafe.Pointer) *SocketListenerPrivate {
	return &SocketListenerPrivate{native: native}
}

// SocketPrivate is a representation of the C record GSocketPrivate.
type SocketPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketPrivate that represents the SocketPrivate.
func (recv *SocketPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketPrivateNewFromC creates a new SocketPrivate from a pointer to the C GSocketPrivate that represents the SocketPrivate.
func SocketPrivateNewFromC(native unsafe.Pointer) *SocketPrivate {
	return &SocketPrivate{native: native}
}

// SocketServiceClass is a representation of the C record GSocketServiceClass.
type SocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServiceClass that represents the SocketServiceClass.
func (recv *SocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServiceClassNewFromC creates a new SocketServiceClass from a pointer to the C GSocketServiceClass that represents the SocketServiceClass.
func SocketServiceClassNewFromC(native unsafe.Pointer) *SocketServiceClass {
	return &SocketServiceClass{native: native}
}

// SocketServicePrivate is a representation of the C record GSocketServicePrivate.
type SocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketServicePrivate that represents the SocketServicePrivate.
func (recv *SocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// SocketServicePrivateNewFromC creates a new SocketServicePrivate from a pointer to the C GSocketServicePrivate that represents the SocketServicePrivate.
func SocketServicePrivateNewFromC(native unsafe.Pointer) *SocketServicePrivate {
	return &SocketServicePrivate{native: native}
}

// SrvTarget is a representation of the C record GSrvTarget.
type SrvTarget struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSrvTarget that represents the SrvTarget.
func (recv *SrvTarget) ToC() unsafe.Pointer {
	return recv.native
}

// SrvTargetNewFromC creates a new SrvTarget from a pointer to the C GSrvTarget that represents the SrvTarget.
func SrvTargetNewFromC(native unsafe.Pointer) *SrvTarget {
	return &SrvTarget{native: native}
}

// StaticResource is a representation of the C record GStaticResource.
type StaticResource struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GStaticResource that represents the StaticResource.
func (recv *StaticResource) ToC() unsafe.Pointer {
	return recv.native
}

// StaticResourceNewFromC creates a new StaticResource from a pointer to the C GStaticResource that represents the StaticResource.
func StaticResourceNewFromC(native unsafe.Pointer) *StaticResource {
	return &StaticResource{native: native}
}

// TaskClass is a representation of the C record GTaskClass.
type TaskClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTaskClass that represents the TaskClass.
func (recv *TaskClass) ToC() unsafe.Pointer {
	return recv.native
}

// TaskClassNewFromC creates a new TaskClass from a pointer to the C GTaskClass that represents the TaskClass.
func TaskClassNewFromC(native unsafe.Pointer) *TaskClass {
	return &TaskClass{native: native}
}

// TcpConnectionClass is a representation of the C record GTcpConnectionClass.
type TcpConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionClass that represents the TcpConnectionClass.
func (recv *TcpConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionClassNewFromC creates a new TcpConnectionClass from a pointer to the C GTcpConnectionClass that represents the TcpConnectionClass.
func TcpConnectionClassNewFromC(native unsafe.Pointer) *TcpConnectionClass {
	return &TcpConnectionClass{native: native}
}

// TcpConnectionPrivate is a representation of the C record GTcpConnectionPrivate.
type TcpConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpConnectionPrivate that represents the TcpConnectionPrivate.
func (recv *TcpConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TcpConnectionPrivateNewFromC creates a new TcpConnectionPrivate from a pointer to the C GTcpConnectionPrivate that represents the TcpConnectionPrivate.
func TcpConnectionPrivateNewFromC(native unsafe.Pointer) *TcpConnectionPrivate {
	return &TcpConnectionPrivate{native: native}
}

// TcpWrapperConnectionClass is a representation of the C record GTcpWrapperConnectionClass.
type TcpWrapperConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionClass that represents the TcpWrapperConnectionClass.
func (recv *TcpWrapperConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionClassNewFromC creates a new TcpWrapperConnectionClass from a pointer to the C GTcpWrapperConnectionClass that represents the TcpWrapperConnectionClass.
func TcpWrapperConnectionClassNewFromC(native unsafe.Pointer) *TcpWrapperConnectionClass {
	return &TcpWrapperConnectionClass{native: native}
}

// TcpWrapperConnectionPrivate is a representation of the C record GTcpWrapperConnectionPrivate.
type TcpWrapperConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTcpWrapperConnectionPrivate that represents the TcpWrapperConnectionPrivate.
func (recv *TcpWrapperConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TcpWrapperConnectionPrivateNewFromC creates a new TcpWrapperConnectionPrivate from a pointer to the C GTcpWrapperConnectionPrivate that represents the TcpWrapperConnectionPrivate.
func TcpWrapperConnectionPrivateNewFromC(native unsafe.Pointer) *TcpWrapperConnectionPrivate {
	return &TcpWrapperConnectionPrivate{native: native}
}

// ThemedIconClass is a representation of the C record GThemedIconClass.
type ThemedIconClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIconClass that represents the ThemedIconClass.
func (recv *ThemedIconClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIconClassNewFromC creates a new ThemedIconClass from a pointer to the C GThemedIconClass that represents the ThemedIconClass.
func ThemedIconClassNewFromC(native unsafe.Pointer) *ThemedIconClass {
	return &ThemedIconClass{native: native}
}

// ThreadedSocketServiceClass is a representation of the C record GThreadedSocketServiceClass.
type ThreadedSocketServiceClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServiceClass that represents the ThreadedSocketServiceClass.
func (recv *ThreadedSocketServiceClass) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServiceClassNewFromC creates a new ThreadedSocketServiceClass from a pointer to the C GThreadedSocketServiceClass that represents the ThreadedSocketServiceClass.
func ThreadedSocketServiceClassNewFromC(native unsafe.Pointer) *ThreadedSocketServiceClass {
	return &ThreadedSocketServiceClass{native: native}
}

// ThreadedSocketServicePrivate is a representation of the C record GThreadedSocketServicePrivate.
type ThreadedSocketServicePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThreadedSocketServicePrivate that represents the ThreadedSocketServicePrivate.
func (recv *ThreadedSocketServicePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// ThreadedSocketServicePrivateNewFromC creates a new ThreadedSocketServicePrivate from a pointer to the C GThreadedSocketServicePrivate that represents the ThreadedSocketServicePrivate.
func ThreadedSocketServicePrivateNewFromC(native unsafe.Pointer) *ThreadedSocketServicePrivate {
	return &ThreadedSocketServicePrivate{native: native}
}

// TlsCertificateClass is a representation of the C record GTlsCertificateClass.
type TlsCertificateClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificateClass that represents the TlsCertificateClass.
func (recv *TlsCertificateClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificateClassNewFromC creates a new TlsCertificateClass from a pointer to the C GTlsCertificateClass that represents the TlsCertificateClass.
func TlsCertificateClassNewFromC(native unsafe.Pointer) *TlsCertificateClass {
	return &TlsCertificateClass{native: native}
}

// TlsCertificatePrivate is a representation of the C record GTlsCertificatePrivate.
type TlsCertificatePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsCertificatePrivate that represents the TlsCertificatePrivate.
func (recv *TlsCertificatePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsCertificatePrivateNewFromC creates a new TlsCertificatePrivate from a pointer to the C GTlsCertificatePrivate that represents the TlsCertificatePrivate.
func TlsCertificatePrivateNewFromC(native unsafe.Pointer) *TlsCertificatePrivate {
	return &TlsCertificatePrivate{native: native}
}

// TlsConnectionClass is a representation of the C record GTlsConnectionClass.
type TlsConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionClass that represents the TlsConnectionClass.
func (recv *TlsConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionClassNewFromC creates a new TlsConnectionClass from a pointer to the C GTlsConnectionClass that represents the TlsConnectionClass.
func TlsConnectionClassNewFromC(native unsafe.Pointer) *TlsConnectionClass {
	return &TlsConnectionClass{native: native}
}

// TlsConnectionPrivate is a representation of the C record GTlsConnectionPrivate.
type TlsConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsConnectionPrivate that represents the TlsConnectionPrivate.
func (recv *TlsConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsConnectionPrivateNewFromC creates a new TlsConnectionPrivate from a pointer to the C GTlsConnectionPrivate that represents the TlsConnectionPrivate.
func TlsConnectionPrivateNewFromC(native unsafe.Pointer) *TlsConnectionPrivate {
	return &TlsConnectionPrivate{native: native}
}

// TlsDatabasePrivate is a representation of the C record GTlsDatabasePrivate.
type TlsDatabasePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsDatabasePrivate that represents the TlsDatabasePrivate.
func (recv *TlsDatabasePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsDatabasePrivateNewFromC creates a new TlsDatabasePrivate from a pointer to the C GTlsDatabasePrivate that represents the TlsDatabasePrivate.
func TlsDatabasePrivateNewFromC(native unsafe.Pointer) *TlsDatabasePrivate {
	return &TlsDatabasePrivate{native: native}
}

// TlsFileDatabaseInterface is a representation of the C record GTlsFileDatabaseInterface.
type TlsFileDatabaseInterface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsFileDatabaseInterface that represents the TlsFileDatabaseInterface.
func (recv *TlsFileDatabaseInterface) ToC() unsafe.Pointer {
	return recv.native
}

// TlsFileDatabaseInterfaceNewFromC creates a new TlsFileDatabaseInterface from a pointer to the C GTlsFileDatabaseInterface that represents the TlsFileDatabaseInterface.
func TlsFileDatabaseInterfaceNewFromC(native unsafe.Pointer) *TlsFileDatabaseInterface {
	return &TlsFileDatabaseInterface{native: native}
}

// TlsInteractionPrivate is a representation of the C record GTlsInteractionPrivate.
type TlsInteractionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsInteractionPrivate that represents the TlsInteractionPrivate.
func (recv *TlsInteractionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsInteractionPrivateNewFromC creates a new TlsInteractionPrivate from a pointer to the C GTlsInteractionPrivate that represents the TlsInteractionPrivate.
func TlsInteractionPrivateNewFromC(native unsafe.Pointer) *TlsInteractionPrivate {
	return &TlsInteractionPrivate{native: native}
}

// TlsPasswordClass is a representation of the C record GTlsPasswordClass.
type TlsPasswordClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordClass that represents the TlsPasswordClass.
func (recv *TlsPasswordClass) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordClassNewFromC creates a new TlsPasswordClass from a pointer to the C GTlsPasswordClass that represents the TlsPasswordClass.
func TlsPasswordClassNewFromC(native unsafe.Pointer) *TlsPasswordClass {
	return &TlsPasswordClass{native: native}
}

// TlsPasswordPrivate is a representation of the C record GTlsPasswordPrivate.
type TlsPasswordPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTlsPasswordPrivate that represents the TlsPasswordPrivate.
func (recv *TlsPasswordPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// TlsPasswordPrivateNewFromC creates a new TlsPasswordPrivate from a pointer to the C GTlsPasswordPrivate that represents the TlsPasswordPrivate.
func TlsPasswordPrivateNewFromC(native unsafe.Pointer) *TlsPasswordPrivate {
	return &TlsPasswordPrivate{native: native}
}

// UnixConnectionClass is a representation of the C record GUnixConnectionClass.
type UnixConnectionClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionClass that represents the UnixConnectionClass.
func (recv *UnixConnectionClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionClassNewFromC creates a new UnixConnectionClass from a pointer to the C GUnixConnectionClass that represents the UnixConnectionClass.
func UnixConnectionClassNewFromC(native unsafe.Pointer) *UnixConnectionClass {
	return &UnixConnectionClass{native: native}
}

// UnixConnectionPrivate is a representation of the C record GUnixConnectionPrivate.
type UnixConnectionPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixConnectionPrivate that represents the UnixConnectionPrivate.
func (recv *UnixConnectionPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixConnectionPrivateNewFromC creates a new UnixConnectionPrivate from a pointer to the C GUnixConnectionPrivate that represents the UnixConnectionPrivate.
func UnixConnectionPrivateNewFromC(native unsafe.Pointer) *UnixConnectionPrivate {
	return &UnixConnectionPrivate{native: native}
}

// UnixCredentialsMessagePrivate is a representation of the C record GUnixCredentialsMessagePrivate.
type UnixCredentialsMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixCredentialsMessagePrivate that represents the UnixCredentialsMessagePrivate.
func (recv *UnixCredentialsMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixCredentialsMessagePrivateNewFromC creates a new UnixCredentialsMessagePrivate from a pointer to the C GUnixCredentialsMessagePrivate that represents the UnixCredentialsMessagePrivate.
func UnixCredentialsMessagePrivateNewFromC(native unsafe.Pointer) *UnixCredentialsMessagePrivate {
	return &UnixCredentialsMessagePrivate{native: native}
}

// UnixFDListClass is a representation of the C record GUnixFDListClass.
type UnixFDListClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListClass that represents the UnixFDListClass.
func (recv *UnixFDListClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListClassNewFromC creates a new UnixFDListClass from a pointer to the C GUnixFDListClass that represents the UnixFDListClass.
func UnixFDListClassNewFromC(native unsafe.Pointer) *UnixFDListClass {
	return &UnixFDListClass{native: native}
}

// UnixFDListPrivate is a representation of the C record GUnixFDListPrivate.
type UnixFDListPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDListPrivate that represents the UnixFDListPrivate.
func (recv *UnixFDListPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListPrivateNewFromC creates a new UnixFDListPrivate from a pointer to the C GUnixFDListPrivate that represents the UnixFDListPrivate.
func UnixFDListPrivateNewFromC(native unsafe.Pointer) *UnixFDListPrivate {
	return &UnixFDListPrivate{native: native}
}

// UnixFDMessageClass is a representation of the C record GUnixFDMessageClass.
type UnixFDMessageClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessageClass that represents the UnixFDMessageClass.
func (recv *UnixFDMessageClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessageClassNewFromC creates a new UnixFDMessageClass from a pointer to the C GUnixFDMessageClass that represents the UnixFDMessageClass.
func UnixFDMessageClassNewFromC(native unsafe.Pointer) *UnixFDMessageClass {
	return &UnixFDMessageClass{native: native}
}

// UnixFDMessagePrivate is a representation of the C record GUnixFDMessagePrivate.
type UnixFDMessagePrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessagePrivate that represents the UnixFDMessagePrivate.
func (recv *UnixFDMessagePrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessagePrivateNewFromC creates a new UnixFDMessagePrivate from a pointer to the C GUnixFDMessagePrivate that represents the UnixFDMessagePrivate.
func UnixFDMessagePrivateNewFromC(native unsafe.Pointer) *UnixFDMessagePrivate {
	return &UnixFDMessagePrivate{native: native}
}

// UnixInputStreamClass is a representation of the C record GUnixInputStreamClass.
type UnixInputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamClass that represents the UnixInputStreamClass.
func (recv *UnixInputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamClassNewFromC creates a new UnixInputStreamClass from a pointer to the C GUnixInputStreamClass that represents the UnixInputStreamClass.
func UnixInputStreamClassNewFromC(native unsafe.Pointer) *UnixInputStreamClass {
	return &UnixInputStreamClass{native: native}
}

// UnixInputStreamPrivate is a representation of the C record GUnixInputStreamPrivate.
type UnixInputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStreamPrivate that represents the UnixInputStreamPrivate.
func (recv *UnixInputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamPrivateNewFromC creates a new UnixInputStreamPrivate from a pointer to the C GUnixInputStreamPrivate that represents the UnixInputStreamPrivate.
func UnixInputStreamPrivateNewFromC(native unsafe.Pointer) *UnixInputStreamPrivate {
	return &UnixInputStreamPrivate{native: native}
}

// UnixMountEntry is a representation of the C record GUnixMountEntry.
type UnixMountEntry struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountEntry that represents the UnixMountEntry.
func (recv *UnixMountEntry) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountEntryNewFromC creates a new UnixMountEntry from a pointer to the C GUnixMountEntry that represents the UnixMountEntry.
func UnixMountEntryNewFromC(native unsafe.Pointer) *UnixMountEntry {
	return &UnixMountEntry{native: native}
}

// UnixMountMonitorClass is a representation of the C record GUnixMountMonitorClass.
type UnixMountMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitorClass that represents the UnixMountMonitorClass.
func (recv *UnixMountMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorClassNewFromC creates a new UnixMountMonitorClass from a pointer to the C GUnixMountMonitorClass that represents the UnixMountMonitorClass.
func UnixMountMonitorClassNewFromC(native unsafe.Pointer) *UnixMountMonitorClass {
	return &UnixMountMonitorClass{native: native}
}

// UnixMountPoint is a representation of the C record GUnixMountPoint.
type UnixMountPoint struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountPoint that represents the UnixMountPoint.
func (recv *UnixMountPoint) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountPointNewFromC creates a new UnixMountPoint from a pointer to the C GUnixMountPoint that represents the UnixMountPoint.
func UnixMountPointNewFromC(native unsafe.Pointer) *UnixMountPoint {
	return &UnixMountPoint{native: native}
}

// UnixOutputStreamClass is a representation of the C record GUnixOutputStreamClass.
type UnixOutputStreamClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamClass that represents the UnixOutputStreamClass.
func (recv *UnixOutputStreamClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamClassNewFromC creates a new UnixOutputStreamClass from a pointer to the C GUnixOutputStreamClass that represents the UnixOutputStreamClass.
func UnixOutputStreamClassNewFromC(native unsafe.Pointer) *UnixOutputStreamClass {
	return &UnixOutputStreamClass{native: native}
}

// UnixOutputStreamPrivate is a representation of the C record GUnixOutputStreamPrivate.
type UnixOutputStreamPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStreamPrivate that represents the UnixOutputStreamPrivate.
func (recv *UnixOutputStreamPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamPrivateNewFromC creates a new UnixOutputStreamPrivate from a pointer to the C GUnixOutputStreamPrivate that represents the UnixOutputStreamPrivate.
func UnixOutputStreamPrivateNewFromC(native unsafe.Pointer) *UnixOutputStreamPrivate {
	return &UnixOutputStreamPrivate{native: native}
}

// UnixSocketAddressClass is a representation of the C record GUnixSocketAddressClass.
type UnixSocketAddressClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressClass that represents the UnixSocketAddressClass.
func (recv *UnixSocketAddressClass) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressClassNewFromC creates a new UnixSocketAddressClass from a pointer to the C GUnixSocketAddressClass that represents the UnixSocketAddressClass.
func UnixSocketAddressClassNewFromC(native unsafe.Pointer) *UnixSocketAddressClass {
	return &UnixSocketAddressClass{native: native}
}

// UnixSocketAddressPrivate is a representation of the C record GUnixSocketAddressPrivate.
type UnixSocketAddressPrivate struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddressPrivate that represents the UnixSocketAddressPrivate.
func (recv *UnixSocketAddressPrivate) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressPrivateNewFromC creates a new UnixSocketAddressPrivate from a pointer to the C GUnixSocketAddressPrivate that represents the UnixSocketAddressPrivate.
func UnixSocketAddressPrivateNewFromC(native unsafe.Pointer) *UnixSocketAddressPrivate {
	return &UnixSocketAddressPrivate{native: native}
}

// VfsClass is a representation of the C record GVfsClass.
type VfsClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfsClass that represents the VfsClass.
func (recv *VfsClass) ToC() unsafe.Pointer {
	return recv.native
}

// VfsClassNewFromC creates a new VfsClass from a pointer to the C GVfsClass that represents the VfsClass.
func VfsClassNewFromC(native unsafe.Pointer) *VfsClass {
	return &VfsClass{native: native}
}

// VolumeIface is a representation of the C record GVolumeIface.
type VolumeIface struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeIface that represents the VolumeIface.
func (recv *VolumeIface) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeIfaceNewFromC creates a new VolumeIface from a pointer to the C GVolumeIface that represents the VolumeIface.
func VolumeIfaceNewFromC(native unsafe.Pointer) *VolumeIface {
	return &VolumeIface{native: native}
}

// VolumeMonitorClass is a representation of the C record GVolumeMonitorClass.
type VolumeMonitorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitorClass that represents the VolumeMonitorClass.
func (recv *VolumeMonitorClass) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitorClassNewFromC creates a new VolumeMonitorClass from a pointer to the C GVolumeMonitorClass that represents the VolumeMonitorClass.
func VolumeMonitorClassNewFromC(native unsafe.Pointer) *VolumeMonitorClass {
	return &VolumeMonitorClass{native: native}
}

// ZlibCompressorClass is a representation of the C record GZlibCompressorClass.
type ZlibCompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressorClass that represents the ZlibCompressorClass.
func (recv *ZlibCompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressorClassNewFromC creates a new ZlibCompressorClass from a pointer to the C GZlibCompressorClass that represents the ZlibCompressorClass.
func ZlibCompressorClassNewFromC(native unsafe.Pointer) *ZlibCompressorClass {
	return &ZlibCompressorClass{native: native}
}

// ZlibDecompressorClass is a representation of the C record GZlibDecompressorClass.
type ZlibDecompressorClass struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressorClass that represents the ZlibDecompressorClass.
func (recv *ZlibDecompressorClass) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressorClassNewFromC creates a new ZlibDecompressorClass from a pointer to the C GZlibDecompressorClass that represents the ZlibDecompressorClass.
func ZlibDecompressorClassNewFromC(native unsafe.Pointer) *ZlibDecompressorClass {
	return &ZlibDecompressorClass{native: native}
}

// AppLaunchContext is a representation of the C record GAppLaunchContext.
type AppLaunchContext struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppLaunchContext that represents the AppLaunchContext.
func (recv *AppLaunchContext) ToC() unsafe.Pointer {
	return recv.native
}

// AppLaunchContextNewFromC creates a new AppLaunchContext from a pointer to the C GAppLaunchContext that represents the AppLaunchContext.
func AppLaunchContextNewFromC(native unsafe.Pointer) *AppLaunchContext {
	return &AppLaunchContext{native: native}
}

// ApplicationCommandLine is a representation of the C record GApplicationCommandLine.
type ApplicationCommandLine struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GApplicationCommandLine that represents the ApplicationCommandLine.
func (recv *ApplicationCommandLine) ToC() unsafe.Pointer {
	return recv.native
}

// ApplicationCommandLineNewFromC creates a new ApplicationCommandLine from a pointer to the C GApplicationCommandLine that represents the ApplicationCommandLine.
func ApplicationCommandLineNewFromC(native unsafe.Pointer) *ApplicationCommandLine {
	return &ApplicationCommandLine{native: native}
}

// BufferedInputStream is a representation of the C record GBufferedInputStream.
type BufferedInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedInputStream that represents the BufferedInputStream.
func (recv *BufferedInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedInputStreamNewFromC creates a new BufferedInputStream from a pointer to the C GBufferedInputStream that represents the BufferedInputStream.
func BufferedInputStreamNewFromC(native unsafe.Pointer) *BufferedInputStream {
	return &BufferedInputStream{native: native}
}

// BufferedOutputStream is a representation of the C record GBufferedOutputStream.
type BufferedOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GBufferedOutputStream that represents the BufferedOutputStream.
func (recv *BufferedOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// BufferedOutputStreamNewFromC creates a new BufferedOutputStream from a pointer to the C GBufferedOutputStream that represents the BufferedOutputStream.
func BufferedOutputStreamNewFromC(native unsafe.Pointer) *BufferedOutputStream {
	return &BufferedOutputStream{native: native}
}

// Cancellable is a representation of the C record GCancellable.
type Cancellable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCancellable that represents the Cancellable.
func (recv *Cancellable) ToC() unsafe.Pointer {
	return recv.native
}

// CancellableNewFromC creates a new Cancellable from a pointer to the C GCancellable that represents the Cancellable.
func CancellableNewFromC(native unsafe.Pointer) *Cancellable {
	return &Cancellable{native: native}
}

// CharsetConverter is a representation of the C record GCharsetConverter.
type CharsetConverter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GCharsetConverter that represents the CharsetConverter.
func (recv *CharsetConverter) ToC() unsafe.Pointer {
	return recv.native
}

// CharsetConverterNewFromC creates a new CharsetConverter from a pointer to the C GCharsetConverter that represents the CharsetConverter.
func CharsetConverterNewFromC(native unsafe.Pointer) *CharsetConverter {
	return &CharsetConverter{native: native}
}

// ConverterInputStream is a representation of the C record GConverterInputStream.
type ConverterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterInputStream that represents the ConverterInputStream.
func (recv *ConverterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterInputStreamNewFromC creates a new ConverterInputStream from a pointer to the C GConverterInputStream that represents the ConverterInputStream.
func ConverterInputStreamNewFromC(native unsafe.Pointer) *ConverterInputStream {
	return &ConverterInputStream{native: native}
}

// ConverterOutputStream is a representation of the C record GConverterOutputStream.
type ConverterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GConverterOutputStream that represents the ConverterOutputStream.
func (recv *ConverterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// ConverterOutputStreamNewFromC creates a new ConverterOutputStream from a pointer to the C GConverterOutputStream that represents the ConverterOutputStream.
func ConverterOutputStreamNewFromC(native unsafe.Pointer) *ConverterOutputStream {
	return &ConverterOutputStream{native: native}
}

// DBusActionGroup is a representation of the C record GDBusActionGroup.
type DBusActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusActionGroup that represents the DBusActionGroup.
func (recv *DBusActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// DBusActionGroupNewFromC creates a new DBusActionGroup from a pointer to the C GDBusActionGroup that represents the DBusActionGroup.
func DBusActionGroupNewFromC(native unsafe.Pointer) *DBusActionGroup {
	return &DBusActionGroup{native: native}
}

// DBusMenuModel is a representation of the C record GDBusMenuModel.
type DBusMenuModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusMenuModel that represents the DBusMenuModel.
func (recv *DBusMenuModel) ToC() unsafe.Pointer {
	return recv.native
}

// DBusMenuModelNewFromC creates a new DBusMenuModel from a pointer to the C GDBusMenuModel that represents the DBusMenuModel.
func DBusMenuModelNewFromC(native unsafe.Pointer) *DBusMenuModel {
	return &DBusMenuModel{native: native}
}

// DataInputStream is a representation of the C record GDataInputStream.
type DataInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataInputStream that represents the DataInputStream.
func (recv *DataInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DataInputStreamNewFromC creates a new DataInputStream from a pointer to the C GDataInputStream that represents the DataInputStream.
func DataInputStreamNewFromC(native unsafe.Pointer) *DataInputStream {
	return &DataInputStream{native: native}
}

// DataOutputStream is a representation of the C record GDataOutputStream.
type DataOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDataOutputStream that represents the DataOutputStream.
func (recv *DataOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// DataOutputStreamNewFromC creates a new DataOutputStream from a pointer to the C GDataOutputStream that represents the DataOutputStream.
func DataOutputStreamNewFromC(native unsafe.Pointer) *DataOutputStream {
	return &DataOutputStream{native: native}
}

// DesktopAppInfo is a representation of the C record GDesktopAppInfo.
type DesktopAppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfo that represents the DesktopAppInfo.
func (recv *DesktopAppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoNewFromC creates a new DesktopAppInfo from a pointer to the C GDesktopAppInfo that represents the DesktopAppInfo.
func DesktopAppInfoNewFromC(native unsafe.Pointer) *DesktopAppInfo {
	return &DesktopAppInfo{native: native}
}

// Emblem is a representation of the C record GEmblem.
type Emblem struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblem that represents the Emblem.
func (recv *Emblem) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemNewFromC creates a new Emblem from a pointer to the C GEmblem that represents the Emblem.
func EmblemNewFromC(native unsafe.Pointer) *Emblem {
	return &Emblem{native: native}
}

// EmblemedIcon is a representation of the C record GEmblemedIcon.
type EmblemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GEmblemedIcon that represents the EmblemedIcon.
func (recv *EmblemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// EmblemedIconNewFromC creates a new EmblemedIcon from a pointer to the C GEmblemedIcon that represents the EmblemedIcon.
func EmblemedIconNewFromC(native unsafe.Pointer) *EmblemedIcon {
	return &EmblemedIcon{native: native}
}

// FileEnumerator is a representation of the C record GFileEnumerator.
type FileEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileEnumerator that represents the FileEnumerator.
func (recv *FileEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// FileEnumeratorNewFromC creates a new FileEnumerator from a pointer to the C GFileEnumerator that represents the FileEnumerator.
func FileEnumeratorNewFromC(native unsafe.Pointer) *FileEnumerator {
	return &FileEnumerator{native: native}
}

// FileIcon is a representation of the C record GFileIcon.
type FileIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileIcon that represents the FileIcon.
func (recv *FileIcon) ToC() unsafe.Pointer {
	return recv.native
}

// FileIconNewFromC creates a new FileIcon from a pointer to the C GFileIcon that represents the FileIcon.
func FileIconNewFromC(native unsafe.Pointer) *FileIcon {
	return &FileIcon{native: native}
}

// FileInfo is a representation of the C record GFileInfo.
type FileInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInfo that represents the FileInfo.
func (recv *FileInfo) ToC() unsafe.Pointer {
	return recv.native
}

// FileInfoNewFromC creates a new FileInfo from a pointer to the C GFileInfo that represents the FileInfo.
func FileInfoNewFromC(native unsafe.Pointer) *FileInfo {
	return &FileInfo{native: native}
}

// FileInputStream is a representation of the C record GFileInputStream.
type FileInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileInputStream that represents the FileInputStream.
func (recv *FileInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileInputStreamNewFromC creates a new FileInputStream from a pointer to the C GFileInputStream that represents the FileInputStream.
func FileInputStreamNewFromC(native unsafe.Pointer) *FileInputStream {
	return &FileInputStream{native: native}
}

// FileMonitor is a representation of the C record GFileMonitor.
type FileMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileMonitor that represents the FileMonitor.
func (recv *FileMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// FileMonitorNewFromC creates a new FileMonitor from a pointer to the C GFileMonitor that represents the FileMonitor.
func FileMonitorNewFromC(native unsafe.Pointer) *FileMonitor {
	return &FileMonitor{native: native}
}

// FileOutputStream is a representation of the C record GFileOutputStream.
type FileOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFileOutputStream that represents the FileOutputStream.
func (recv *FileOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FileOutputStreamNewFromC creates a new FileOutputStream from a pointer to the C GFileOutputStream that represents the FileOutputStream.
func FileOutputStreamNewFromC(native unsafe.Pointer) *FileOutputStream {
	return &FileOutputStream{native: native}
}

// FilenameCompleter is a representation of the C record GFilenameCompleter.
type FilenameCompleter struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilenameCompleter that represents the FilenameCompleter.
func (recv *FilenameCompleter) ToC() unsafe.Pointer {
	return recv.native
}

// FilenameCompleterNewFromC creates a new FilenameCompleter from a pointer to the C GFilenameCompleter that represents the FilenameCompleter.
func FilenameCompleterNewFromC(native unsafe.Pointer) *FilenameCompleter {
	return &FilenameCompleter{native: native}
}

// FilterInputStream is a representation of the C record GFilterInputStream.
type FilterInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterInputStream that represents the FilterInputStream.
func (recv *FilterInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilterInputStreamNewFromC creates a new FilterInputStream from a pointer to the C GFilterInputStream that represents the FilterInputStream.
func FilterInputStreamNewFromC(native unsafe.Pointer) *FilterInputStream {
	return &FilterInputStream{native: native}
}

// FilterOutputStream is a representation of the C record GFilterOutputStream.
type FilterOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFilterOutputStream that represents the FilterOutputStream.
func (recv *FilterOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// FilterOutputStreamNewFromC creates a new FilterOutputStream from a pointer to the C GFilterOutputStream that represents the FilterOutputStream.
func FilterOutputStreamNewFromC(native unsafe.Pointer) *FilterOutputStream {
	return &FilterOutputStream{native: native}
}

// IOModule is a representation of the C record GIOModule.
type IOModule struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIOModule that represents the IOModule.
func (recv *IOModule) ToC() unsafe.Pointer {
	return recv.native
}

// IOModuleNewFromC creates a new IOModule from a pointer to the C GIOModule that represents the IOModule.
func IOModuleNewFromC(native unsafe.Pointer) *IOModule {
	return &IOModule{native: native}
}

// InetAddress is a representation of the C record GInetAddress.
type InetAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetAddress that represents the InetAddress.
func (recv *InetAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InetAddressNewFromC creates a new InetAddress from a pointer to the C GInetAddress that represents the InetAddress.
func InetAddressNewFromC(native unsafe.Pointer) *InetAddress {
	return &InetAddress{native: native}
}

// InetSocketAddress is a representation of the C record GInetSocketAddress.
type InetSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInetSocketAddress that represents the InetSocketAddress.
func (recv *InetSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// InetSocketAddressNewFromC creates a new InetSocketAddress from a pointer to the C GInetSocketAddress that represents the InetSocketAddress.
func InetSocketAddressNewFromC(native unsafe.Pointer) *InetSocketAddress {
	return &InetSocketAddress{native: native}
}

// InputStream is a representation of the C record GInputStream.
type InputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GInputStream that represents the InputStream.
func (recv *InputStream) ToC() unsafe.Pointer {
	return recv.native
}

// InputStreamNewFromC creates a new InputStream from a pointer to the C GInputStream that represents the InputStream.
func InputStreamNewFromC(native unsafe.Pointer) *InputStream {
	return &InputStream{native: native}
}

// ListStore is a representation of the C record GListStore.
type ListStore struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListStore that represents the ListStore.
func (recv *ListStore) ToC() unsafe.Pointer {
	return recv.native
}

// ListStoreNewFromC creates a new ListStore from a pointer to the C GListStore that represents the ListStore.
func ListStoreNewFromC(native unsafe.Pointer) *ListStore {
	return &ListStore{native: native}
}

// MemoryInputStream is a representation of the C record GMemoryInputStream.
type MemoryInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryInputStream that represents the MemoryInputStream.
func (recv *MemoryInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryInputStreamNewFromC creates a new MemoryInputStream from a pointer to the C GMemoryInputStream that represents the MemoryInputStream.
func MemoryInputStreamNewFromC(native unsafe.Pointer) *MemoryInputStream {
	return &MemoryInputStream{native: native}
}

// MemoryOutputStream is a representation of the C record GMemoryOutputStream.
type MemoryOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMemoryOutputStream that represents the MemoryOutputStream.
func (recv *MemoryOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// MemoryOutputStreamNewFromC creates a new MemoryOutputStream from a pointer to the C GMemoryOutputStream that represents the MemoryOutputStream.
func MemoryOutputStreamNewFromC(native unsafe.Pointer) *MemoryOutputStream {
	return &MemoryOutputStream{native: native}
}

// MountOperation is a representation of the C record GMountOperation.
type MountOperation struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMountOperation that represents the MountOperation.
func (recv *MountOperation) ToC() unsafe.Pointer {
	return recv.native
}

// MountOperationNewFromC creates a new MountOperation from a pointer to the C GMountOperation that represents the MountOperation.
func MountOperationNewFromC(native unsafe.Pointer) *MountOperation {
	return &MountOperation{native: native}
}

// NativeSocketAddress is a representation of the C record GNativeSocketAddress.
type NativeSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeSocketAddress that represents the NativeSocketAddress.
func (recv *NativeSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NativeSocketAddressNewFromC creates a new NativeSocketAddress from a pointer to the C GNativeSocketAddress that represents the NativeSocketAddress.
func NativeSocketAddressNewFromC(native unsafe.Pointer) *NativeSocketAddress {
	return &NativeSocketAddress{native: native}
}

// NativeVolumeMonitor is a representation of the C record GNativeVolumeMonitor.
type NativeVolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNativeVolumeMonitor that represents the NativeVolumeMonitor.
func (recv *NativeVolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// NativeVolumeMonitorNewFromC creates a new NativeVolumeMonitor from a pointer to the C GNativeVolumeMonitor that represents the NativeVolumeMonitor.
func NativeVolumeMonitorNewFromC(native unsafe.Pointer) *NativeVolumeMonitor {
	return &NativeVolumeMonitor{native: native}
}

// NetworkAddress is a representation of the C record GNetworkAddress.
type NetworkAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkAddress that represents the NetworkAddress.
func (recv *NetworkAddress) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkAddressNewFromC creates a new NetworkAddress from a pointer to the C GNetworkAddress that represents the NetworkAddress.
func NetworkAddressNewFromC(native unsafe.Pointer) *NetworkAddress {
	return &NetworkAddress{native: native}
}

// NetworkService is a representation of the C record GNetworkService.
type NetworkService struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GNetworkService that represents the NetworkService.
func (recv *NetworkService) ToC() unsafe.Pointer {
	return recv.native
}

// NetworkServiceNewFromC creates a new NetworkService from a pointer to the C GNetworkService that represents the NetworkService.
func NetworkServiceNewFromC(native unsafe.Pointer) *NetworkService {
	return &NetworkService{native: native}
}

// OutputStream is a representation of the C record GOutputStream.
type OutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GOutputStream that represents the OutputStream.
func (recv *OutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// OutputStreamNewFromC creates a new OutputStream from a pointer to the C GOutputStream that represents the OutputStream.
func OutputStreamNewFromC(native unsafe.Pointer) *OutputStream {
	return &OutputStream{native: native}
}

// Permission is a representation of the C record GPermission.
type Permission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GPermission that represents the Permission.
func (recv *Permission) ToC() unsafe.Pointer {
	return recv.native
}

// PermissionNewFromC creates a new Permission from a pointer to the C GPermission that represents the Permission.
func PermissionNewFromC(native unsafe.Pointer) *Permission {
	return &Permission{native: native}
}

// ProxyAddressEnumerator is a representation of the C record GProxyAddressEnumerator.
type ProxyAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GProxyAddressEnumerator that represents the ProxyAddressEnumerator.
func (recv *ProxyAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// ProxyAddressEnumeratorNewFromC creates a new ProxyAddressEnumerator from a pointer to the C GProxyAddressEnumerator that represents the ProxyAddressEnumerator.
func ProxyAddressEnumeratorNewFromC(native unsafe.Pointer) *ProxyAddressEnumerator {
	return &ProxyAddressEnumerator{native: native}
}

// Resolver is a representation of the C record GResolver.
type Resolver struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GResolver that represents the Resolver.
func (recv *Resolver) ToC() unsafe.Pointer {
	return recv.native
}

// ResolverNewFromC creates a new Resolver from a pointer to the C GResolver that represents the Resolver.
func ResolverNewFromC(native unsafe.Pointer) *Resolver {
	return &Resolver{native: native}
}

// Settings is a representation of the C record GSettings.
type Settings struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettings that represents the Settings.
func (recv *Settings) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsNewFromC creates a new Settings from a pointer to the C GSettings that represents the Settings.
func SettingsNewFromC(native unsafe.Pointer) *Settings {
	return &Settings{native: native}
}

// SettingsBackend is a representation of the C record GSettingsBackend.
type SettingsBackend struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSettingsBackend that represents the SettingsBackend.
func (recv *SettingsBackend) ToC() unsafe.Pointer {
	return recv.native
}

// SettingsBackendNewFromC creates a new SettingsBackend from a pointer to the C GSettingsBackend that represents the SettingsBackend.
func SettingsBackendNewFromC(native unsafe.Pointer) *SettingsBackend {
	return &SettingsBackend{native: native}
}

// SimpleAction is a representation of the C record GSimpleAction.
type SimpleAction struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAction that represents the SimpleAction.
func (recv *SimpleAction) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleActionNewFromC creates a new SimpleAction from a pointer to the C GSimpleAction that represents the SimpleAction.
func SimpleActionNewFromC(native unsafe.Pointer) *SimpleAction {
	return &SimpleAction{native: native}
}

// SimpleAsyncResult is a representation of the C record GSimpleAsyncResult.
type SimpleAsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimpleAsyncResult that represents the SimpleAsyncResult.
func (recv *SimpleAsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// SimpleAsyncResultNewFromC creates a new SimpleAsyncResult from a pointer to the C GSimpleAsyncResult that represents the SimpleAsyncResult.
func SimpleAsyncResultNewFromC(native unsafe.Pointer) *SimpleAsyncResult {
	return &SimpleAsyncResult{native: native}
}

// SimplePermission is a representation of the C record GSimplePermission.
type SimplePermission struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSimplePermission that represents the SimplePermission.
func (recv *SimplePermission) ToC() unsafe.Pointer {
	return recv.native
}

// SimplePermissionNewFromC creates a new SimplePermission from a pointer to the C GSimplePermission that represents the SimplePermission.
func SimplePermissionNewFromC(native unsafe.Pointer) *SimplePermission {
	return &SimplePermission{native: native}
}

// SocketAddress is a representation of the C record GSocketAddress.
type SocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddress that represents the SocketAddress.
func (recv *SocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressNewFromC creates a new SocketAddress from a pointer to the C GSocketAddress that represents the SocketAddress.
func SocketAddressNewFromC(native unsafe.Pointer) *SocketAddress {
	return &SocketAddress{native: native}
}

// SocketAddressEnumerator is a representation of the C record GSocketAddressEnumerator.
type SocketAddressEnumerator struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketAddressEnumerator that represents the SocketAddressEnumerator.
func (recv *SocketAddressEnumerator) ToC() unsafe.Pointer {
	return recv.native
}

// SocketAddressEnumeratorNewFromC creates a new SocketAddressEnumerator from a pointer to the C GSocketAddressEnumerator that represents the SocketAddressEnumerator.
func SocketAddressEnumeratorNewFromC(native unsafe.Pointer) *SocketAddressEnumerator {
	return &SocketAddressEnumerator{native: native}
}

// Task is a representation of the C record GTask.
type Task struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GTask that represents the Task.
func (recv *Task) ToC() unsafe.Pointer {
	return recv.native
}

// TaskNewFromC creates a new Task from a pointer to the C GTask that represents the Task.
func TaskNewFromC(native unsafe.Pointer) *Task {
	return &Task{native: native}
}

// ThemedIcon is a representation of the C record GThemedIcon.
type ThemedIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GThemedIcon that represents the ThemedIcon.
func (recv *ThemedIcon) ToC() unsafe.Pointer {
	return recv.native
}

// ThemedIconNewFromC creates a new ThemedIcon from a pointer to the C GThemedIcon that represents the ThemedIcon.
func ThemedIconNewFromC(native unsafe.Pointer) *ThemedIcon {
	return &ThemedIcon{native: native}
}

// UnixFDList is a representation of the C record GUnixFDList.
type UnixFDList struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDList that represents the UnixFDList.
func (recv *UnixFDList) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDListNewFromC creates a new UnixFDList from a pointer to the C GUnixFDList that represents the UnixFDList.
func UnixFDListNewFromC(native unsafe.Pointer) *UnixFDList {
	return &UnixFDList{native: native}
}

// UnixFDMessage is a representation of the C record GUnixFDMessage.
type UnixFDMessage struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixFDMessage that represents the UnixFDMessage.
func (recv *UnixFDMessage) ToC() unsafe.Pointer {
	return recv.native
}

// UnixFDMessageNewFromC creates a new UnixFDMessage from a pointer to the C GUnixFDMessage that represents the UnixFDMessage.
func UnixFDMessageNewFromC(native unsafe.Pointer) *UnixFDMessage {
	return &UnixFDMessage{native: native}
}

// UnixInputStream is a representation of the C record GUnixInputStream.
type UnixInputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixInputStream that represents the UnixInputStream.
func (recv *UnixInputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixInputStreamNewFromC creates a new UnixInputStream from a pointer to the C GUnixInputStream that represents the UnixInputStream.
func UnixInputStreamNewFromC(native unsafe.Pointer) *UnixInputStream {
	return &UnixInputStream{native: native}
}

// UnixMountMonitor is a representation of the C record GUnixMountMonitor.
type UnixMountMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixMountMonitor that represents the UnixMountMonitor.
func (recv *UnixMountMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// UnixMountMonitorNewFromC creates a new UnixMountMonitor from a pointer to the C GUnixMountMonitor that represents the UnixMountMonitor.
func UnixMountMonitorNewFromC(native unsafe.Pointer) *UnixMountMonitor {
	return &UnixMountMonitor{native: native}
}

// UnixOutputStream is a representation of the C record GUnixOutputStream.
type UnixOutputStream struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixOutputStream that represents the UnixOutputStream.
func (recv *UnixOutputStream) ToC() unsafe.Pointer {
	return recv.native
}

// UnixOutputStreamNewFromC creates a new UnixOutputStream from a pointer to the C GUnixOutputStream that represents the UnixOutputStream.
func UnixOutputStreamNewFromC(native unsafe.Pointer) *UnixOutputStream {
	return &UnixOutputStream{native: native}
}

// UnixSocketAddress is a representation of the C record GUnixSocketAddress.
type UnixSocketAddress struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GUnixSocketAddress that represents the UnixSocketAddress.
func (recv *UnixSocketAddress) ToC() unsafe.Pointer {
	return recv.native
}

// UnixSocketAddressNewFromC creates a new UnixSocketAddress from a pointer to the C GUnixSocketAddress that represents the UnixSocketAddress.
func UnixSocketAddressNewFromC(native unsafe.Pointer) *UnixSocketAddress {
	return &UnixSocketAddress{native: native}
}

// Vfs is a representation of the C record GVfs.
type Vfs struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVfs that represents the Vfs.
func (recv *Vfs) ToC() unsafe.Pointer {
	return recv.native
}

// VfsNewFromC creates a new Vfs from a pointer to the C GVfs that represents the Vfs.
func VfsNewFromC(native unsafe.Pointer) *Vfs {
	return &Vfs{native: native}
}

// VolumeMonitor is a representation of the C record GVolumeMonitor.
type VolumeMonitor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolumeMonitor that represents the VolumeMonitor.
func (recv *VolumeMonitor) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeMonitorNewFromC creates a new VolumeMonitor from a pointer to the C GVolumeMonitor that represents the VolumeMonitor.
func VolumeMonitorNewFromC(native unsafe.Pointer) *VolumeMonitor {
	return &VolumeMonitor{native: native}
}

// ZlibCompressor is a representation of the C record GZlibCompressor.
type ZlibCompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibCompressor that represents the ZlibCompressor.
func (recv *ZlibCompressor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibCompressorNewFromC creates a new ZlibCompressor from a pointer to the C GZlibCompressor that represents the ZlibCompressor.
func ZlibCompressorNewFromC(native unsafe.Pointer) *ZlibCompressor {
	return &ZlibCompressor{native: native}
}

// ZlibDecompressor is a representation of the C record GZlibDecompressor.
type ZlibDecompressor struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GZlibDecompressor that represents the ZlibDecompressor.
func (recv *ZlibDecompressor) ToC() unsafe.Pointer {
	return recv.native
}

// ZlibDecompressorNewFromC creates a new ZlibDecompressor from a pointer to the C GZlibDecompressor that represents the ZlibDecompressor.
func ZlibDecompressorNewFromC(native unsafe.Pointer) *ZlibDecompressor {
	return &ZlibDecompressor{native: native}
}

// Action is a representation of the C interface GAction.
type Action struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAction that represents the Action.
func (recv *Action) ToC() unsafe.Pointer {
	return recv.native
}

// ActionNewFromC creates a new Action from a pointer to the C GAction that represents the Action.
func ActionNewFromC(native unsafe.Pointer) *Action {
	return &Action{native: native}
}

// ActionGroup is a representation of the C interface GActionGroup.
type ActionGroup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GActionGroup that represents the ActionGroup.
func (recv *ActionGroup) ToC() unsafe.Pointer {
	return recv.native
}

// ActionGroupNewFromC creates a new ActionGroup from a pointer to the C GActionGroup that represents the ActionGroup.
func ActionGroupNewFromC(native unsafe.Pointer) *ActionGroup {
	return &ActionGroup{native: native}
}

// AppInfo is a representation of the C interface GAppInfo.
type AppInfo struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAppInfo that represents the AppInfo.
func (recv *AppInfo) ToC() unsafe.Pointer {
	return recv.native
}

// AppInfoNewFromC creates a new AppInfo from a pointer to the C GAppInfo that represents the AppInfo.
func AppInfoNewFromC(native unsafe.Pointer) *AppInfo {
	return &AppInfo{native: native}
}

// AsyncResult is a representation of the C interface GAsyncResult.
type AsyncResult struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GAsyncResult that represents the AsyncResult.
func (recv *AsyncResult) ToC() unsafe.Pointer {
	return recv.native
}

// AsyncResultNewFromC creates a new AsyncResult from a pointer to the C GAsyncResult that represents the AsyncResult.
func AsyncResultNewFromC(native unsafe.Pointer) *AsyncResult {
	return &AsyncResult{native: native}
}

// DBusObject is a representation of the C interface GDBusObject.
type DBusObject struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObject that represents the DBusObject.
func (recv *DBusObject) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectNewFromC creates a new DBusObject from a pointer to the C GDBusObject that represents the DBusObject.
func DBusObjectNewFromC(native unsafe.Pointer) *DBusObject {
	return &DBusObject{native: native}
}

// DBusObjectManager is a representation of the C interface GDBusObjectManager.
type DBusObjectManager struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDBusObjectManager that represents the DBusObjectManager.
func (recv *DBusObjectManager) ToC() unsafe.Pointer {
	return recv.native
}

// DBusObjectManagerNewFromC creates a new DBusObjectManager from a pointer to the C GDBusObjectManager that represents the DBusObjectManager.
func DBusObjectManagerNewFromC(native unsafe.Pointer) *DBusObjectManager {
	return &DBusObjectManager{native: native}
}

// DesktopAppInfoLookup is a representation of the C interface GDesktopAppInfoLookup.
type DesktopAppInfoLookup struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDesktopAppInfoLookup that represents the DesktopAppInfoLookup.
func (recv *DesktopAppInfoLookup) ToC() unsafe.Pointer {
	return recv.native
}

// DesktopAppInfoLookupNewFromC creates a new DesktopAppInfoLookup from a pointer to the C GDesktopAppInfoLookup that represents the DesktopAppInfoLookup.
func DesktopAppInfoLookupNewFromC(native unsafe.Pointer) *DesktopAppInfoLookup {
	return &DesktopAppInfoLookup{native: native}
}

// Drive is a representation of the C interface GDrive.
type Drive struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GDrive that represents the Drive.
func (recv *Drive) ToC() unsafe.Pointer {
	return recv.native
}

// DriveNewFromC creates a new Drive from a pointer to the C GDrive that represents the Drive.
func DriveNewFromC(native unsafe.Pointer) *Drive {
	return &Drive{native: native}
}

// File is a representation of the C interface GFile.
type File struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GFile that represents the File.
func (recv *File) ToC() unsafe.Pointer {
	return recv.native
}

// FileNewFromC creates a new File from a pointer to the C GFile that represents the File.
func FileNewFromC(native unsafe.Pointer) *File {
	return &File{native: native}
}

// Icon is a representation of the C interface GIcon.
type Icon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GIcon that represents the Icon.
func (recv *Icon) ToC() unsafe.Pointer {
	return recv.native
}

// IconNewFromC creates a new Icon from a pointer to the C GIcon that represents the Icon.
func IconNewFromC(native unsafe.Pointer) *Icon {
	return &Icon{native: native}
}

// ListModel is a representation of the C interface GListModel.
type ListModel struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GListModel that represents the ListModel.
func (recv *ListModel) ToC() unsafe.Pointer {
	return recv.native
}

// ListModelNewFromC creates a new ListModel from a pointer to the C GListModel that represents the ListModel.
func ListModelNewFromC(native unsafe.Pointer) *ListModel {
	return &ListModel{native: native}
}

// LoadableIcon is a representation of the C interface GLoadableIcon.
type LoadableIcon struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GLoadableIcon that represents the LoadableIcon.
func (recv *LoadableIcon) ToC() unsafe.Pointer {
	return recv.native
}

// LoadableIconNewFromC creates a new LoadableIcon from a pointer to the C GLoadableIcon that represents the LoadableIcon.
func LoadableIconNewFromC(native unsafe.Pointer) *LoadableIcon {
	return &LoadableIcon{native: native}
}

// Mount is a representation of the C interface GMount.
type Mount struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GMount that represents the Mount.
func (recv *Mount) ToC() unsafe.Pointer {
	return recv.native
}

// MountNewFromC creates a new Mount from a pointer to the C GMount that represents the Mount.
func MountNewFromC(native unsafe.Pointer) *Mount {
	return &Mount{native: native}
}

// Seekable is a representation of the C interface GSeekable.
type Seekable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSeekable that represents the Seekable.
func (recv *Seekable) ToC() unsafe.Pointer {
	return recv.native
}

// SeekableNewFromC creates a new Seekable from a pointer to the C GSeekable that represents the Seekable.
func SeekableNewFromC(native unsafe.Pointer) *Seekable {
	return &Seekable{native: native}
}

// SocketConnectable is a representation of the C interface GSocketConnectable.
type SocketConnectable struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GSocketConnectable that represents the SocketConnectable.
func (recv *SocketConnectable) ToC() unsafe.Pointer {
	return recv.native
}

// SocketConnectableNewFromC creates a new SocketConnectable from a pointer to the C GSocketConnectable that represents the SocketConnectable.
func SocketConnectableNewFromC(native unsafe.Pointer) *SocketConnectable {
	return &SocketConnectable{native: native}
}

// Volume is a representation of the C interface GVolume.
type Volume struct {
	native unsafe.Pointer
}

// ToC returns a pointer to the C GVolume that represents the Volume.
func (recv *Volume) ToC() unsafe.Pointer {
	return recv.native
}

// VolumeNewFromC creates a new Volume from a pointer to the C GVolume that represents the Volume.
func VolumeNewFromC(native unsafe.Pointer) *Volume {
	return &Volume{native: native}
}