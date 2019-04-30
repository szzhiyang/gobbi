// This is a generated file - DO NOT EDIT
// +build gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import "unsafe"

// Unsupported : gtk_paper_size_new : return type :

// Unsupported : gtk_paper_size_new_custom : return type :

// Unsupported : gtk_paper_size_new_from_ppd : return type :

// RecentInfo is a wrapper around the C record GtkRecentInfo.
type RecentInfo struct {
	native unsafe.Pointer
}

func RecentInfoNewFromC(u unsafe.Pointer) *RecentInfo {
	if u == nil {
		return nil
	}

	g := &RecentInfo{native: u}

	return g
}

func (recv *RecentInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// RecentManagerClass is a wrapper around the C record GtkRecentManagerClass.
type RecentManagerClass struct {
	native unsafe.Pointer
	// Private : parent_class
	// no type for changed
	// no type for _gtk_recent1
	// no type for _gtk_recent2
	// no type for _gtk_recent3
	// no type for _gtk_recent4
}

func RecentManagerClassNewFromC(u unsafe.Pointer) *RecentManagerClass {
	if u == nil {
		return nil
	}

	g := &RecentManagerClass{native: u}

	return g
}

func (recv *RecentManagerClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
