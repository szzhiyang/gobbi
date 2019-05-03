// Code generated - DO NOT EDIT.
// +build glib_2.6 glib_2.8 glib_2.10 glib_2.12 glib_2.14 glib_2.16 glib_2.18 glib_2.20 glib_2.22 glib_2.24 glib_2.26 glib_2.28 glib_2.30 glib_2.32 glib_2.34 glib_2.36 glib_2.38 glib_2.40 glib_2.44 glib_2.46 glib_2.48 glib_2.50 glib_2.52 glib_2.54 glib_2.56

package glib

import "unsafe"

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <glib.h>
// #include <glib/gstdio.h>
// #include <glib-unix.h>
// #include <stdlib.h>
import "C"

const GSIZE_FORMAT string = C.G_GSIZE_FORMAT
const GSIZE_MODIFIER string = C.G_GSIZE_MODIFIER
const GSSIZE_FORMAT string = C.G_GSSIZE_FORMAT
const GSSIZE_MODIFIER string = C.G_GSSIZE_MODIFIER
const OPTION_REMAINING string = C.G_OPTION_REMAINING

// CheckVersion is a wrapper around the C function glib_check_version.
func CheckVersion(requiredMajor uint32, requiredMinor uint32, requiredMicro uint32) string {
	c_required_major := (C.guint)(requiredMajor)

	c_required_minor := (C.guint)(requiredMinor)

	c_required_micro := (C.guint)(requiredMicro)

	retC := C.glib_check_version(c_required_major, c_required_minor, c_required_micro)
	retGo := C.GoString(retC)

	return retGo
}

// FilenameDisplayBasename is a wrapper around the C function g_filename_display_basename.
func FilenameDisplayBasename(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_basename(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// FilenameDisplayName is a wrapper around the C function g_filename_display_name.
func FilenameDisplayName(filename string) string {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_filename_display_name(c_filename)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : g_get_filename_charsets : unsupported parameter charsets : in string with indirection level of 3

// GetLanguageNames is a wrapper around the C function g_get_language_names.
func GetLanguageNames() []string {
	retC := C.g_get_language_names()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetSystemConfigDirs is a wrapper around the C function g_get_system_config_dirs.
func GetSystemConfigDirs() []string {
	retC := C.g_get_system_config_dirs()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetSystemDataDirs is a wrapper around the C function g_get_system_data_dirs.
func GetSystemDataDirs() []string {
	retC := C.g_get_system_data_dirs()
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}

	return retGo
}

// GetUserCacheDir is a wrapper around the C function g_get_user_cache_dir.
func GetUserCacheDir() string {
	retC := C.g_get_user_cache_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserConfigDir is a wrapper around the C function g_get_user_config_dir.
func GetUserConfigDir() string {
	retC := C.g_get_user_config_dir()
	retGo := C.GoString(retC)

	return retGo
}

// GetUserDataDir is a wrapper around the C function g_get_user_data_dir.
func GetUserDataDir() string {
	retC := C.g_get_user_data_dir()
	retGo := C.GoString(retC)

	return retGo
}

// Unsupported : g_log_set_default_handler : unsupported parameter log_func : no type generator for LogFunc (GLogFunc) for param log_func

// Rmdir is a wrapper around the C function g_rmdir.
func Rmdir(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_rmdir(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// Unsupported : g_strv_length : unsupported parameter str_array : in string with indirection level of 2

// Unlink is a wrapper around the C function g_unlink.
func Unlink(filename string) int32 {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.g_unlink(c_filename)
	retGo := (int32)(retC)

	return retGo
}

// UriListExtractUris is a wrapper around the C function g_uri_list_extract_uris.
func UriListExtractUris(uriList string) []string {
	c_uri_list := C.CString(uriList)
	defer C.free(unsafe.Pointer(c_uri_list))

	retC := C.g_uri_list_extract_uris(c_uri_list)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	return retGo
}

// GetIso8601WeekOfYear is a wrapper around the C function g_date_get_iso8601_week_of_year.
func (recv *Date) GetIso8601WeekOfYear() uint32 {
	retC := C.g_date_get_iso8601_week_of_year((*C.GDate)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// KeyFileNew is a wrapper around the C function g_key_file_new.
func KeyFileNew() *KeyFile {
	retC := C.g_key_file_new()
	retGo := KeyFileNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function g_key_file_free.
func (recv *KeyFile) Free() {
	C.g_key_file_free((*C.GKeyFile)(recv.native))

	return
}

// GetBoolean is a wrapper around the C function g_key_file_get_boolean.
func (recv *KeyFile) GetBoolean(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_boolean_list : array return type :

// GetComment is a wrapper around the C function g_key_file_get_comment.
func (recv *KeyFile) GetComment(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetGroups is a wrapper around the C function g_key_file_get_groups.
func (recv *KeyFile) GetGroups() ([]string, uint64) {
	var c_length C.gsize

	retC := C.g_key_file_get_groups((*C.GKeyFile)(recv.native), &c_length)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	length := (uint64)(c_length)

	return retGo, length
}

// GetInteger is a wrapper around the C function g_key_file_get_integer.
func (recv *KeyFile) GetInteger(groupName string, key string) (int32, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := (int32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// Unsupported : g_key_file_get_integer_list : array return type :

// GetKeys is a wrapper around the C function g_key_file_get_keys.
func (recv *KeyFile) GetKeys(groupName string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_keys((*C.GKeyFile)(recv.native), c_group_name, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetLocaleString is a wrapper around the C function g_key_file_get_locale_string.
func (recv *KeyFile) GetLocaleString(groupName string, key string, locale string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetLocaleStringList is a wrapper around the C function g_key_file_get_locale_string_list.
func (recv *KeyFile) GetLocaleStringList(groupName string, key string, locale string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_locale_string_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetStartGroup is a wrapper around the C function g_key_file_get_start_group.
func (recv *KeyFile) GetStartGroup() string {
	retC := C.g_key_file_get_start_group((*C.GKeyFile)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetString is a wrapper around the C function g_key_file_get_string.
func (recv *KeyFile) GetString(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_string((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetStringList is a wrapper around the C function g_key_file_get_string_list.
func (recv *KeyFile) GetStringList(groupName string, key string) ([]string, uint64, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_get_string_list((*C.GKeyFile)(recv.native), c_group_name, c_key, &c_length, &cThrowableError)
	retGo := []string{}
	for p := retC; *p != nil; p = (**C.char)(C.gpointer((uintptr(C.gpointer(p)) + uintptr(C.sizeof_gpointer)))) {
		s := C.GoString(*p)
		retGo = append(retGo, s)
	}
	defer C.g_strfreev(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// GetValue is a wrapper around the C function g_key_file_get_value.
func (recv *KeyFile) GetValue(groupName string, key string) (string, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_get_value((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// HasGroup is a wrapper around the C function g_key_file_has_group.
func (recv *KeyFile) HasGroup(groupName string) bool {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	retC := C.g_key_file_has_group((*C.GKeyFile)(recv.native), c_group_name)
	retGo := retC == C.TRUE

	return retGo
}

// HasKey is a wrapper around the C function g_key_file_has_key.
func (recv *KeyFile) HasKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_has_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromData is a wrapper around the C function g_key_file_load_from_data.
func (recv *KeyFile) LoadFromData(data string, length uint64, flags KeyFileFlags) (bool, error) {
	c_data := C.CString(data)
	defer C.free(unsafe.Pointer(c_data))

	c_length := (C.gsize)(length)

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data((*C.GKeyFile)(recv.native), c_data, c_length, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LoadFromDataDirs is a wrapper around the C function g_key_file_load_from_data_dirs.
func (recv *KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (bool, string, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	var c_full_path *C.gchar

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_data_dirs((*C.GKeyFile)(recv.native), c_file, &c_full_path, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	fullPath := C.GoString(c_full_path)
	defer C.free(unsafe.Pointer(c_full_path))

	return retGo, fullPath, goError
}

// LoadFromFile is a wrapper around the C function g_key_file_load_from_file.
func (recv *KeyFile) LoadFromFile(file string, flags KeyFileFlags) (bool, error) {
	c_file := C.CString(file)
	defer C.free(unsafe.Pointer(c_file))

	c_flags := (C.GKeyFileFlags)(flags)

	var cThrowableError *C.GError

	retC := C.g_key_file_load_from_file((*C.GKeyFile)(recv.native), c_file, c_flags, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveComment is a wrapper around the C function g_key_file_remove_comment.
func (recv *KeyFile) RemoveComment(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveGroup is a wrapper around the C function g_key_file_remove_group.
func (recv *KeyFile) RemoveGroup(groupName string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_group((*C.GKeyFile)(recv.native), c_group_name, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveKey is a wrapper around the C function g_key_file_remove_key.
func (recv *KeyFile) RemoveKey(groupName string, key string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	var cThrowableError *C.GError

	retC := C.g_key_file_remove_key((*C.GKeyFile)(recv.native), c_group_name, c_key, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetBoolean is a wrapper around the C function g_key_file_set_boolean.
func (recv *KeyFile) SetBoolean(groupName string, key string, value bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value :=
		boolToGboolean(value)

	C.g_key_file_set_boolean((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetBooleanList is a wrapper around the C function g_key_file_set_boolean_list.
func (recv *KeyFile) SetBooleanList(groupName string, key string, list []bool) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gboolean, len(list)+1, len(list)+1)
	for i, item := range list {
		c :=
			boolToGboolean(item)
		c_list_array[i] = c
	}
	c_list_array[len(list)] = 0
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gboolean)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_boolean_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

// SetComment is a wrapper around the C function g_key_file_set_comment.
func (recv *KeyFile) SetComment(groupName string, key string, comment string) (bool, error) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_comment := C.CString(comment)
	defer C.free(unsafe.Pointer(c_comment))

	var cThrowableError *C.GError

	retC := C.g_key_file_set_comment((*C.GKeyFile)(recv.native), c_group_name, c_key, c_comment, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetInteger is a wrapper around the C function g_key_file_set_integer.
func (recv *KeyFile) SetInteger(groupName string, key string, value int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := (C.gint)(value)

	C.g_key_file_set_integer((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// SetIntegerList is a wrapper around the C function g_key_file_set_integer_list.
func (recv *KeyFile) SetIntegerList(groupName string, key string, list []int32) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_list_array := make([]C.gint, len(list)+1, len(list)+1)
	for i, item := range list {
		c := (C.gint)(item)
		c_list_array[i] = c
	}
	c_list_array[len(list)] = 0
	c_list_arrayPtr := &c_list_array[0]
	c_list := (*C.gint)(unsafe.Pointer(c_list_arrayPtr))

	c_length := (C.gsize)(len(list))

	C.g_key_file_set_integer_list((*C.GKeyFile)(recv.native), c_group_name, c_key, c_list, c_length)

	return
}

// SetListSeparator is a wrapper around the C function g_key_file_set_list_separator.
func (recv *KeyFile) SetListSeparator(separator rune) {
	c_separator := (C.gchar)(separator)

	C.g_key_file_set_list_separator((*C.GKeyFile)(recv.native), c_separator)

	return
}

// SetLocaleString is a wrapper around the C function g_key_file_set_locale_string.
func (recv *KeyFile) SetLocaleString(groupName string, key string, locale string, string_ string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_locale := C.CString(locale)
	defer C.free(unsafe.Pointer(c_locale))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_locale_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_locale, c_string)

	return
}

// Blacklisted : g_key_file_set_locale_string_list

// SetString is a wrapper around the C function g_key_file_set_string.
func (recv *KeyFile) SetString(groupName string, key string, string_ string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_string := C.CString(string_)
	defer C.free(unsafe.Pointer(c_string))

	C.g_key_file_set_string((*C.GKeyFile)(recv.native), c_group_name, c_key, c_string)

	return
}

// Blacklisted : g_key_file_set_string_list

// SetValue is a wrapper around the C function g_key_file_set_value.
func (recv *KeyFile) SetValue(groupName string, key string, value string) {
	c_group_name := C.CString(groupName)
	defer C.free(unsafe.Pointer(c_group_name))

	c_key := C.CString(key)
	defer C.free(unsafe.Pointer(c_key))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	C.g_key_file_set_value((*C.GKeyFile)(recv.native), c_group_name, c_key, c_value)

	return
}

// ToData is a wrapper around the C function g_key_file_to_data.
func (recv *KeyFile) ToData() (string, uint64, error) {
	var c_length C.gsize

	var cThrowableError *C.GError

	retC := C.g_key_file_to_data((*C.GKeyFile)(recv.native), &c_length, &cThrowableError)
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	length := (uint64)(c_length)

	return retGo, length, goError
}

// OptionContextNew is a wrapper around the C function g_option_context_new.
func OptionContextNew(parameterString string) *OptionContext {
	c_parameter_string := C.CString(parameterString)
	defer C.free(unsafe.Pointer(c_parameter_string))

	retC := C.g_option_context_new(c_parameter_string)
	retGo := OptionContextNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AddGroup is a wrapper around the C function g_option_context_add_group.
func (recv *OptionContext) AddGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_add_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// AddMainEntries is a wrapper around the C function g_option_context_add_main_entries.
func (recv *OptionContext) AddMainEntries(entries *OptionEntry, translationDomain string) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	c_translation_domain := C.CString(translationDomain)
	defer C.free(unsafe.Pointer(c_translation_domain))

	C.g_option_context_add_main_entries((*C.GOptionContext)(recv.native), c_entries, c_translation_domain)

	return
}

// Free is a wrapper around the C function g_option_context_free.
func (recv *OptionContext) Free() {
	C.g_option_context_free((*C.GOptionContext)(recv.native))

	return
}

// GetHelpEnabled is a wrapper around the C function g_option_context_get_help_enabled.
func (recv *OptionContext) GetHelpEnabled() bool {
	retC := C.g_option_context_get_help_enabled((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIgnoreUnknownOptions is a wrapper around the C function g_option_context_get_ignore_unknown_options.
func (recv *OptionContext) GetIgnoreUnknownOptions() bool {
	retC := C.g_option_context_get_ignore_unknown_options((*C.GOptionContext)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetMainGroup is a wrapper around the C function g_option_context_get_main_group.
func (recv *OptionContext) GetMainGroup() *OptionGroup {
	retC := C.g_option_context_get_main_group((*C.GOptionContext)(recv.native))
	retGo := OptionGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Parse is a wrapper around the C function g_option_context_parse.
func (recv *OptionContext) Parse(args []string) (bool, []string, error) {
	cArgc, cArgv := argsIn(args)

	var cThrowableError *C.GError

	retC := C.g_option_context_parse((*C.GOptionContext)(recv.native), &cArgc, &cArgv, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	args = argsOut(cArgc, cArgv)

	return retGo, args, goError
}

// SetHelpEnabled is a wrapper around the C function g_option_context_set_help_enabled.
func (recv *OptionContext) SetHelpEnabled(helpEnabled bool) {
	c_help_enabled :=
		boolToGboolean(helpEnabled)

	C.g_option_context_set_help_enabled((*C.GOptionContext)(recv.native), c_help_enabled)

	return
}

// SetIgnoreUnknownOptions is a wrapper around the C function g_option_context_set_ignore_unknown_options.
func (recv *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	c_ignore_unknown :=
		boolToGboolean(ignoreUnknown)

	C.g_option_context_set_ignore_unknown_options((*C.GOptionContext)(recv.native), c_ignore_unknown)

	return
}

// SetMainGroup is a wrapper around the C function g_option_context_set_main_group.
func (recv *OptionContext) SetMainGroup(group *OptionGroup) {
	c_group := (*C.GOptionGroup)(C.NULL)
	if group != nil {
		c_group = (*C.GOptionGroup)(group.ToC())
	}

	C.g_option_context_set_main_group((*C.GOptionContext)(recv.native), c_group)

	return
}

// Unsupported : g_option_group_new : unsupported parameter user_data : no type generator for gpointer (gpointer) for param user_data

// AddEntries is a wrapper around the C function g_option_group_add_entries.
func (recv *OptionGroup) AddEntries(entries *OptionEntry) {
	c_entries := (*C.GOptionEntry)(C.NULL)
	if entries != nil {
		c_entries = (*C.GOptionEntry)(entries.ToC())
	}

	C.g_option_group_add_entries((*C.GOptionGroup)(recv.native), c_entries)

	return
}

// Free is a wrapper around the C function g_option_group_free.
func (recv *OptionGroup) Free() {
	C.g_option_group_free((*C.GOptionGroup)(recv.native))

	return
}

// Unsupported : g_option_group_set_error_hook : unsupported parameter error_func : no type generator for OptionErrorFunc (GOptionErrorFunc) for param error_func

// Unsupported : g_option_group_set_parse_hooks : unsupported parameter pre_parse_func : no type generator for OptionParseFunc (GOptionParseFunc) for param pre_parse_func

// Unsupported : g_option_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GTranslateFunc) for param func

// SetTranslationDomain is a wrapper around the C function g_option_group_set_translation_domain.
func (recv *OptionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.g_option_group_set_translation_domain((*C.GOptionGroup)(recv.native), c_domain)

	return
}