// This is a generated file - DO NOT EDIT
// +build pango_1.16 pango_1.18 pango_1.20 pango_1.22 pango_1.26 pango_1.30 pango_1.31.0 pango_1.32 pango_1.32.4 pango_1.34 pango_1.38

package pango

import call "github.com/pekim/gobbi/lib/internal/call"

// Unsupported : pango_attr_gravity_hint_new : return type :

// Unsupported : pango_attr_gravity_new : return type :

// Unsupported : pango_gravity_get_for_matrix : return type :

// Unsupported : pango_gravity_get_for_script : return type :

// Unsupported : pango_gravity_to_rotation : return type :

// Unsupported : pango_language_get_default : return type :

// Unsupported : pango_parse_enum : return type :

// Unsupported : pango_units_to_double : return type :

// Version is a wrapper around the C function pango_version.
func Version() int32 {
	data := call.Data{}
	call.Function(9286, data)
	retGo := int32(3)

	return retGo
}

// Unsupported : pango_version_check : return type :

// Unsupported : pango_version_string : return type :
