// This is a generated file - DO NOT EDIT

package gobject

import (
	glib "github.com/pekim/gobbi/lib/glib"
	"unsafe"
)

// CClosure is a wrapper around the C record GCClosure.
type CClosure struct {
	native unsafe.Pointer
	// closure : record
	// callback : no type generator for gpointer, gpointer
}

func CClosureNewFromC(u unsafe.Pointer) *CClosure {
	if u == nil {
		return nil
	}

	g := &CClosure{native: u}

	return g
}

func (recv *CClosure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Closure is a wrapper around the C record GClosure.
type Closure struct {
	native unsafe.Pointer
	// Private : ref_count
	// Private : meta_marshal_nouse
	// Private : n_guards
	// Private : n_fnotifiers
	// Private : n_inotifiers
	// Private : in_inotify
	// Private : floating
	// Private : derivative_flag
	// Bitfield not supported :  1 in_marshal
	// Bitfield not supported :  1 is_invalid
	// no type for marshal
	// Private : data
	// Private : notifiers
}

func ClosureNewFromC(u unsafe.Pointer) *Closure {
	if u == nil {
		return nil
	}

	g := &Closure{native: u}

	return g
}

func (recv *Closure) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_closure_new_object : return type :

// Unsupported : g_closure_new_simple : unsupported parameter data : no type generator for gpointer (gpointer) for param data

// ClosureNotifyData is a wrapper around the C record GClosureNotifyData.
type ClosureNotifyData struct {
	native unsafe.Pointer
	// data : no type generator for gpointer, gpointer
	// notify : no type generator for ClosureNotify, GClosureNotify
}

func ClosureNotifyDataNewFromC(u unsafe.Pointer) *ClosureNotifyData {
	if u == nil {
		return nil
	}

	g := &ClosureNotifyData{native: u}

	return g
}

func (recv *ClosureNotifyData) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EnumClass is a wrapper around the C record GEnumClass.
type EnumClass struct {
	native unsafe.Pointer
	// g_type_class : record
	Minimum int32
	Maximum int32
	NValues uint32
	// values : record
}

func EnumClassNewFromC(u unsafe.Pointer) *EnumClass {
	if u == nil {
		return nil
	}

	g := &EnumClass{native: u}

	return g
}

func (recv *EnumClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// EnumValue is a wrapper around the C record GEnumValue.
type EnumValue struct {
	native    unsafe.Pointer
	Value     int32
	ValueName string
	ValueNick string
}

func EnumValueNewFromC(u unsafe.Pointer) *EnumValue {
	if u == nil {
		return nil
	}

	g := &EnumValue{native: u}

	return g
}

func (recv *EnumValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlagsClass is a wrapper around the C record GFlagsClass.
type FlagsClass struct {
	native unsafe.Pointer
	// g_type_class : record
	Mask    uint32
	NValues uint32
	// values : record
}

func FlagsClassNewFromC(u unsafe.Pointer) *FlagsClass {
	if u == nil {
		return nil
	}

	g := &FlagsClass{native: u}

	return g
}

func (recv *FlagsClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// FlagsValue is a wrapper around the C record GFlagsValue.
type FlagsValue struct {
	native    unsafe.Pointer
	Value     uint32
	ValueName string
	ValueNick string
}

func FlagsValueNewFromC(u unsafe.Pointer) *FlagsValue {
	if u == nil {
		return nil
	}

	g := &FlagsValue{native: u}

	return g
}

func (recv *FlagsValue) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InitiallyUnownedClass is a wrapper around the C record GInitiallyUnownedClass.
type InitiallyUnownedClass struct {
	native unsafe.Pointer
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func InitiallyUnownedClassNewFromC(u unsafe.Pointer) *InitiallyUnownedClass {
	if u == nil {
		return nil
	}

	g := &InitiallyUnownedClass{native: u}

	return g
}

func (recv *InitiallyUnownedClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// InterfaceInfo is a wrapper around the C record GInterfaceInfo.
type InterfaceInfo struct {
	native unsafe.Pointer
	// interface_init : no type generator for InterfaceInitFunc, GInterfaceInitFunc
	// interface_finalize : no type generator for InterfaceFinalizeFunc, GInterfaceFinalizeFunc
	// interface_data : no type generator for gpointer, gpointer
}

func InterfaceInfoNewFromC(u unsafe.Pointer) *InterfaceInfo {
	if u == nil {
		return nil
	}

	g := &InterfaceInfo{native: u}

	return g
}

func (recv *InterfaceInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectClass is a wrapper around the C record GObjectClass.
type ObjectClass struct {
	native unsafe.Pointer
	// g_type_class : record
	// Private : construct_properties
	// no type for constructor
	// no type for set_property
	// no type for get_property
	// no type for dispose
	// no type for finalize
	// no type for dispatch_properties_changed
	// no type for notify
	// no type for constructed
	// Private : flags
	// Private : pdummy
}

func ObjectClassNewFromC(u unsafe.Pointer) *ObjectClass {
	if u == nil {
		return nil
	}

	g := &ObjectClass{native: u}

	return g
}

func (recv *ObjectClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ObjectConstructParam is a wrapper around the C record GObjectConstructParam.
type ObjectConstructParam struct {
	native unsafe.Pointer
	// pspec : record
	// value : record
}

func ObjectConstructParamNewFromC(u unsafe.Pointer) *ObjectConstructParam {
	if u == nil {
		return nil
	}

	g := &ObjectConstructParam{native: u}

	return g
}

func (recv *ObjectConstructParam) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecClass is a wrapper around the C record GParamSpecClass.
type ParamSpecClass struct {
	native unsafe.Pointer
	// g_type_class : record
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
	// Private : dummy
}

func ParamSpecClassNewFromC(u unsafe.Pointer) *ParamSpecClass {
	if u == nil {
		return nil
	}

	g := &ParamSpecClass{native: u}

	return g
}

func (recv *ParamSpecClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecPool is a wrapper around the C record GParamSpecPool.
type ParamSpecPool struct {
	native unsafe.Pointer
}

func ParamSpecPoolNewFromC(u unsafe.Pointer) *ParamSpecPool {
	if u == nil {
		return nil
	}

	g := &ParamSpecPool{native: u}

	return g
}

func (recv *ParamSpecPool) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ParamSpecTypeInfo is a wrapper around the C record GParamSpecTypeInfo.
type ParamSpecTypeInfo struct {
	native       unsafe.Pointer
	InstanceSize uint16
	NPreallocs   uint16
	// no type for instance_init
	ValueType Type
	// no type for finalize
	// no type for value_set_default
	// no type for value_validate
	// no type for values_cmp
}

func ParamSpecTypeInfoNewFromC(u unsafe.Pointer) *ParamSpecTypeInfo {
	if u == nil {
		return nil
	}

	g := &ParamSpecTypeInfo{native: u}

	return g
}

func (recv *ParamSpecTypeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Parameter is a wrapper around the C record GParameter.
type Parameter struct {
	native unsafe.Pointer
	Name   string
	// value : record
}

func ParameterNewFromC(u unsafe.Pointer) *Parameter {
	if u == nil {
		return nil
	}

	g := &Parameter{native: u}

	return g
}

func (recv *Parameter) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SignalInvocationHint is a wrapper around the C record GSignalInvocationHint.
type SignalInvocationHint struct {
	native   unsafe.Pointer
	SignalId uint32
	Detail   glib.Quark
	RunType  SignalFlags
}

func SignalInvocationHintNewFromC(u unsafe.Pointer) *SignalInvocationHint {
	if u == nil {
		return nil
	}

	g := &SignalInvocationHint{native: u}

	return g
}

func (recv *SignalInvocationHint) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// SignalQuery is a wrapper around the C record GSignalQuery.
type SignalQuery struct {
	native      unsafe.Pointer
	SignalId    uint32
	SignalName  string
	Itype       Type
	SignalFlags SignalFlags
	ReturnType  Type
	NParams     uint32
	// no type for param_types
}

func SignalQueryNewFromC(u unsafe.Pointer) *SignalQuery {
	if u == nil {
		return nil
	}

	g := &SignalQuery{native: u}

	return g
}

func (recv *SignalQuery) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeClass is a wrapper around the C record GTypeClass.
type TypeClass struct {
	native unsafe.Pointer
	// Private : g_type
}

func TypeClassNewFromC(u unsafe.Pointer) *TypeClass {
	if u == nil {
		return nil
	}

	g := &TypeClass{native: u}

	return g
}

func (recv *TypeClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeFundamentalInfo is a wrapper around the C record GTypeFundamentalInfo.
type TypeFundamentalInfo struct {
	native    unsafe.Pointer
	TypeFlags TypeFundamentalFlags
}

func TypeFundamentalInfoNewFromC(u unsafe.Pointer) *TypeFundamentalInfo {
	if u == nil {
		return nil
	}

	g := &TypeFundamentalInfo{native: u}

	return g
}

func (recv *TypeFundamentalInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInfo is a wrapper around the C record GTypeInfo.
type TypeInfo struct {
	native    unsafe.Pointer
	ClassSize uint16
	// base_init : no type generator for BaseInitFunc, GBaseInitFunc
	// base_finalize : no type generator for BaseFinalizeFunc, GBaseFinalizeFunc
	// class_init : no type generator for ClassInitFunc, GClassInitFunc
	// class_finalize : no type generator for ClassFinalizeFunc, GClassFinalizeFunc
	// class_data : no type generator for gpointer, gconstpointer
	InstanceSize uint16
	NPreallocs   uint16
	// instance_init : no type generator for InstanceInitFunc, GInstanceInitFunc
	// value_table : record
}

func TypeInfoNewFromC(u unsafe.Pointer) *TypeInfo {
	if u == nil {
		return nil
	}

	g := &TypeInfo{native: u}

	return g
}

func (recv *TypeInfo) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInstance is a wrapper around the C record GTypeInstance.
type TypeInstance struct {
	native unsafe.Pointer
	// Private : g_class
}

func TypeInstanceNewFromC(u unsafe.Pointer) *TypeInstance {
	if u == nil {
		return nil
	}

	g := &TypeInstance{native: u}

	return g
}

func (recv *TypeInstance) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeInterface is a wrapper around the C record GTypeInterface.
type TypeInterface struct {
	native unsafe.Pointer
	// Private : g_type
	// Private : g_instance_type
}

func TypeInterfaceNewFromC(u unsafe.Pointer) *TypeInterface {
	if u == nil {
		return nil
	}

	g := &TypeInterface{native: u}

	return g
}

func (recv *TypeInterface) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeModuleClass is a wrapper around the C record GTypeModuleClass.
type TypeModuleClass struct {
	native unsafe.Pointer
	// parent_class : record
	// no type for load
	// no type for unload
	// no type for reserved1
	// no type for reserved2
	// no type for reserved3
	// no type for reserved4
}

func TypeModuleClassNewFromC(u unsafe.Pointer) *TypeModuleClass {
	if u == nil {
		return nil
	}

	g := &TypeModuleClass{native: u}

	return g
}

func (recv *TypeModuleClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypePluginClass is a wrapper around the C record GTypePluginClass.
type TypePluginClass struct {
	native unsafe.Pointer
	// Private : base_iface
	// use_plugin : no type generator for TypePluginUse, GTypePluginUse
	// unuse_plugin : no type generator for TypePluginUnuse, GTypePluginUnuse
	// complete_type_info : no type generator for TypePluginCompleteTypeInfo, GTypePluginCompleteTypeInfo
	// complete_interface_info : no type generator for TypePluginCompleteInterfaceInfo, GTypePluginCompleteInterfaceInfo
}

func TypePluginClassNewFromC(u unsafe.Pointer) *TypePluginClass {
	if u == nil {
		return nil
	}

	g := &TypePluginClass{native: u}

	return g
}

func (recv *TypePluginClass) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeQuery is a wrapper around the C record GTypeQuery.
type TypeQuery struct {
	native       unsafe.Pointer
	Type         Type
	TypeName     string
	ClassSize    uint32
	InstanceSize uint32
}

func TypeQueryNewFromC(u unsafe.Pointer) *TypeQuery {
	if u == nil {
		return nil
	}

	g := &TypeQuery{native: u}

	return g
}

func (recv *TypeQuery) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// TypeValueTable is a wrapper around the C record GTypeValueTable.
type TypeValueTable struct {
	native unsafe.Pointer
	// no type for value_init
	// no type for value_free
	// no type for value_copy
	// no type for value_peek_pointer
	CollectFormat string
	// no type for collect_value
	LcopyFormat string
	// no type for lcopy_value
}

func TypeValueTableNewFromC(u unsafe.Pointer) *TypeValueTable {
	if u == nil {
		return nil
	}

	g := &TypeValueTable{native: u}

	return g
}

func (recv *TypeValueTable) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Value is a wrapper around the C record GValue.
type Value struct {
	native unsafe.Pointer
	// Private : g_type
	// no type for data
}

func ValueNewFromC(u unsafe.Pointer) *Value {
	if u == nil {
		return nil
	}

	g := &Value{native: u}

	return g
}

func (recv *Value) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// ValueArray is a wrapper around the C record GValueArray.
type ValueArray struct {
	native  unsafe.Pointer
	NValues uint32
	// values : record
	// Private : n_prealloced
}

func ValueArrayNewFromC(u unsafe.Pointer) *ValueArray {
	if u == nil {
		return nil
	}

	g := &ValueArray{native: u}

	return g
}

func (recv *ValueArray) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}

// Unsupported : g_value_array_new : return type :

// WeakRef is a wrapper around the C record GWeakRef.
type WeakRef struct {
	native unsafe.Pointer
}

func WeakRefNewFromC(u unsafe.Pointer) *WeakRef {
	if u == nil {
		return nil
	}

	g := &WeakRef{native: u}

	return g
}

func (recv *WeakRef) ToC() unsafe.Pointer {

	return (unsafe.Pointer)(recv.native)
}
