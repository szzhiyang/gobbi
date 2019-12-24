package generate

var functionBlacklist = map[string]bool{
	// gdk
	"gdk_synthesize_window_state": true,
	"gdk_window_destroy_notify":   true,

	// gio
	"g_io_module_query": true,

	// settings backend implementations
	"g_settings_backend_get_default":           true,
	"g_settings_backend_changed":               true,
	"g_settings_backend_path_changed":          true,
	"g_settings_backend_keys_changed":          true,
	"g_settings_backend_path_writable_changed": true,
	"g_settings_backend_writable_changed":      true,
	"g_settings_backend_changed_tree":          true,
	"g_settings_backend_flatten_tree":          true,
	"g_keyfile_settings_backend_new":           true,
	"g_memory_settings_backend_new":            true,
	"g_null_settings_backend_new":              true,

	// pango
	"pango_module_register":                true,
	"pango_get_lib_subdirectory":           true,
	"pango_get_sysconf_subdirectory":       true,
	"pango_fontset_simple_new":             true,
	"pango_fontset_simple_append":          true,
	"pango_fontset_simple_size":            true,
	"pango_config_key_get":                 true,
	"pango_config_key_get_system":          true,
	"pango_default_break":                  true,
	"pango_find_map":                       true,
	"pango_font_map_get_shape_engine_type": true,

	// xlib
	"XOpenDisplay": true,
}

func (f *Function) applyAddenda() {
	if f.namespace.Name == "cairo" {
		f.blacklist = true
		return
	}

	if f.namespace.Name == "GLib" && f.Name == "clear_error" {
		f.Parameters = append(f.Parameters, &Parameter{
			Name: "err",
			Argument: Argument{
				Type: &Type{
					Name:  "Error",
					CType: "GError**",
				},
			},
		})
		return
	}

	// A small number of glib will be manually implemented, but in general
	// glib functions are not needed.
	if f.namespace.Name == "GLib" {
		f.blacklist = true
	}

	if _, ok := functionBlacklist[f.CIdentifier]; ok {
		f.blacklist = true
	}
}
