// Code generated - DO NOT EDIT.

package pango

import (
	gi "github.com/pekim/gobbi/internal/gi"
	"runtime"
	"sync"
	"unsafe"
)

var analysisStruct *gi.Struct
var analysisStruct_Once sync.Once

func analysisStruct_Set() error {
	var err error
	analysisStruct_Once.Do(func() {
		analysisStruct, err = gi.StructNew("Pango", "Analysis")
	})
	return err
}

type Analysis struct {
	native unsafe.Pointer
}

func AnalysisNewFromNative(native unsafe.Pointer) *Analysis {
	return &Analysis{native: native}
}

// FieldShapeEngine returns the C field 'shape_engine'.
func (recv *Analysis) FieldShapeEngine() *EngineShape {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "shape_engine")
	value := EngineShapeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldShapeEngine sets the value of the C field 'shape_engine'.
func (recv *Analysis) SetFieldShapeEngine(value *EngineShape) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(analysisStruct, recv.native, "shape_engine", argValue)
}

// FieldLangEngine returns the C field 'lang_engine'.
func (recv *Analysis) FieldLangEngine() *EngineLang {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "lang_engine")
	value := EngineLangNewFromNative(argValue.Pointer())
	return value
}

// SetFieldLangEngine sets the value of the C field 'lang_engine'.
func (recv *Analysis) SetFieldLangEngine(value *EngineLang) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(analysisStruct, recv.native, "lang_engine", argValue)
}

// FieldFont returns the C field 'font'.
func (recv *Analysis) FieldFont() *Font {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "font")
	value := FontNewFromNative(argValue.Pointer())
	return value
}

// SetFieldFont sets the value of the C field 'font'.
func (recv *Analysis) SetFieldFont(value *Font) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(analysisStruct, recv.native, "font", argValue)
}

// FieldLevel returns the C field 'level'.
func (recv *Analysis) FieldLevel() uint8 {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "level")
	value := argValue.Uint8()
	return value
}

// SetFieldLevel sets the value of the C field 'level'.
func (recv *Analysis) SetFieldLevel(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(analysisStruct, recv.native, "level", argValue)
}

// FieldGravity returns the C field 'gravity'.
func (recv *Analysis) FieldGravity() uint8 {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "gravity")
	value := argValue.Uint8()
	return value
}

// SetFieldGravity sets the value of the C field 'gravity'.
func (recv *Analysis) SetFieldGravity(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(analysisStruct, recv.native, "gravity", argValue)
}

// FieldFlags returns the C field 'flags'.
func (recv *Analysis) FieldFlags() uint8 {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "flags")
	value := argValue.Uint8()
	return value
}

// SetFieldFlags sets the value of the C field 'flags'.
func (recv *Analysis) SetFieldFlags(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(analysisStruct, recv.native, "flags", argValue)
}

// FieldScript returns the C field 'script'.
func (recv *Analysis) FieldScript() uint8 {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "script")
	value := argValue.Uint8()
	return value
}

// SetFieldScript sets the value of the C field 'script'.
func (recv *Analysis) SetFieldScript(value uint8) {
	var argValue gi.Argument
	argValue.SetUint8(value)
	gi.StructFieldSet(analysisStruct, recv.native, "script", argValue)
}

// FieldLanguage returns the C field 'language'.
func (recv *Analysis) FieldLanguage() *Language {
	argValue := gi.StructFieldGet(analysisStruct, recv.native, "language")
	value := LanguageNewFromNative(argValue.Pointer())
	return value
}

// SetFieldLanguage sets the value of the C field 'language'.
func (recv *Analysis) SetFieldLanguage(value *Language) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(analysisStruct, recv.native, "language", argValue)
}

// UNSUPPORTED : C value 'extra_attrs' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'extra_attrs' : for field setter : no Go type for 'GLib.SList'

// AnalysisStruct creates an uninitialised Analysis.
func AnalysisStruct() *Analysis {
	err := analysisStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AnalysisNewFromNative(analysisStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAnalysis)
	return structGo
}
func finalizeAnalysis(obj *Analysis) {
	analysisStruct.Free(obj.native)
}

var attrClassStruct *gi.Struct
var attrClassStruct_Once sync.Once

func attrClassStruct_Set() error {
	var err error
	attrClassStruct_Once.Do(func() {
		attrClassStruct, err = gi.StructNew("Pango", "AttrClass")
	})
	return err
}

type AttrClass struct {
	native unsafe.Pointer
}

func AttrClassNewFromNative(native unsafe.Pointer) *AttrClass {
	return &AttrClass{native: native}
}

// FieldType returns the C field 'type'.
func (recv *AttrClass) FieldType() AttrType {
	argValue := gi.StructFieldGet(attrClassStruct, recv.native, "type")
	value := AttrType(argValue.Int32())
	return value
}

// SetFieldType sets the value of the C field 'type'.
func (recv *AttrClass) SetFieldType(value AttrType) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(attrClassStruct, recv.native, "type", argValue)
}

// UNSUPPORTED : C value 'copy' : for field getter : missing Type

// UNSUPPORTED : C value 'copy' : for field setter : missing Type

// UNSUPPORTED : C value 'destroy' : for field getter : missing Type

// UNSUPPORTED : C value 'destroy' : for field setter : missing Type

// UNSUPPORTED : C value 'equal' : for field getter : missing Type

// UNSUPPORTED : C value 'equal' : for field setter : missing Type

// AttrClassStruct creates an uninitialised AttrClass.
func AttrClassStruct() *AttrClass {
	err := attrClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrClassNewFromNative(attrClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrClass)
	return structGo
}
func finalizeAttrClass(obj *AttrClass) {
	attrClassStruct.Free(obj.native)
}

var attrColorStruct *gi.Struct
var attrColorStruct_Once sync.Once

func attrColorStruct_Set() error {
	var err error
	attrColorStruct_Once.Do(func() {
		attrColorStruct, err = gi.StructNew("Pango", "AttrColor")
	})
	return err
}

type AttrColor struct {
	native unsafe.Pointer
}

func AttrColorNewFromNative(native unsafe.Pointer) *AttrColor {
	return &AttrColor{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrColor) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrColorStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrColor) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrColorStruct, recv.native, "attr", argValue)
}

// FieldColor returns the C field 'color'.
func (recv *AttrColor) FieldColor() *Color {
	argValue := gi.StructFieldGet(attrColorStruct, recv.native, "color")
	value := ColorNewFromNative(argValue.Pointer())
	return value
}

// SetFieldColor sets the value of the C field 'color'.
func (recv *AttrColor) SetFieldColor(value *Color) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrColorStruct, recv.native, "color", argValue)
}

// AttrColorStruct creates an uninitialised AttrColor.
func AttrColorStruct() *AttrColor {
	err := attrColorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrColorNewFromNative(attrColorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrColor)
	return structGo
}
func finalizeAttrColor(obj *AttrColor) {
	attrColorStruct.Free(obj.native)
}

var attrFloatStruct *gi.Struct
var attrFloatStruct_Once sync.Once

func attrFloatStruct_Set() error {
	var err error
	attrFloatStruct_Once.Do(func() {
		attrFloatStruct, err = gi.StructNew("Pango", "AttrFloat")
	})
	return err
}

type AttrFloat struct {
	native unsafe.Pointer
}

func AttrFloatNewFromNative(native unsafe.Pointer) *AttrFloat {
	return &AttrFloat{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrFloat) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrFloatStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrFloat) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrFloatStruct, recv.native, "attr", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *AttrFloat) FieldValue() float64 {
	argValue := gi.StructFieldGet(attrFloatStruct, recv.native, "value")
	value := argValue.Float64()
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *AttrFloat) SetFieldValue(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(attrFloatStruct, recv.native, "value", argValue)
}

// AttrFloatStruct creates an uninitialised AttrFloat.
func AttrFloatStruct() *AttrFloat {
	err := attrFloatStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrFloatNewFromNative(attrFloatStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrFloat)
	return structGo
}
func finalizeAttrFloat(obj *AttrFloat) {
	attrFloatStruct.Free(obj.native)
}

var attrFontDescStruct *gi.Struct
var attrFontDescStruct_Once sync.Once

func attrFontDescStruct_Set() error {
	var err error
	attrFontDescStruct_Once.Do(func() {
		attrFontDescStruct, err = gi.StructNew("Pango", "AttrFontDesc")
	})
	return err
}

type AttrFontDesc struct {
	native unsafe.Pointer
}

func AttrFontDescNewFromNative(native unsafe.Pointer) *AttrFontDesc {
	return &AttrFontDesc{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrFontDesc) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrFontDescStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrFontDesc) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrFontDescStruct, recv.native, "attr", argValue)
}

// FieldDesc returns the C field 'desc'.
func (recv *AttrFontDesc) FieldDesc() *FontDescription {
	argValue := gi.StructFieldGet(attrFontDescStruct, recv.native, "desc")
	value := FontDescriptionNewFromNative(argValue.Pointer())
	return value
}

// SetFieldDesc sets the value of the C field 'desc'.
func (recv *AttrFontDesc) SetFieldDesc(value *FontDescription) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrFontDescStruct, recv.native, "desc", argValue)
}

// AttrFontDescStruct creates an uninitialised AttrFontDesc.
func AttrFontDescStruct() *AttrFontDesc {
	err := attrFontDescStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrFontDescNewFromNative(attrFontDescStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrFontDesc)
	return structGo
}
func finalizeAttrFontDesc(obj *AttrFontDesc) {
	attrFontDescStruct.Free(obj.native)
}

var attrFontFeaturesStruct *gi.Struct
var attrFontFeaturesStruct_Once sync.Once

func attrFontFeaturesStruct_Set() error {
	var err error
	attrFontFeaturesStruct_Once.Do(func() {
		attrFontFeaturesStruct, err = gi.StructNew("Pango", "AttrFontFeatures")
	})
	return err
}

type AttrFontFeatures struct {
	native unsafe.Pointer
}

func AttrFontFeaturesNewFromNative(native unsafe.Pointer) *AttrFontFeatures {
	return &AttrFontFeatures{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrFontFeatures) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrFontFeaturesStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrFontFeatures) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrFontFeaturesStruct, recv.native, "attr", argValue)
}

// FieldFeatures returns the C field 'features'.
func (recv *AttrFontFeatures) FieldFeatures() string {
	argValue := gi.StructFieldGet(attrFontFeaturesStruct, recv.native, "features")
	value := argValue.String(false)
	return value
}

// SetFieldFeatures sets the value of the C field 'features'.
func (recv *AttrFontFeatures) SetFieldFeatures(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(attrFontFeaturesStruct, recv.native, "features", argValue)
}

// AttrFontFeaturesStruct creates an uninitialised AttrFontFeatures.
func AttrFontFeaturesStruct() *AttrFontFeatures {
	err := attrFontFeaturesStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrFontFeaturesNewFromNative(attrFontFeaturesStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrFontFeatures)
	return structGo
}
func finalizeAttrFontFeatures(obj *AttrFontFeatures) {
	attrFontFeaturesStruct.Free(obj.native)
}

var attrIntStruct *gi.Struct
var attrIntStruct_Once sync.Once

func attrIntStruct_Set() error {
	var err error
	attrIntStruct_Once.Do(func() {
		attrIntStruct, err = gi.StructNew("Pango", "AttrInt")
	})
	return err
}

type AttrInt struct {
	native unsafe.Pointer
}

func AttrIntNewFromNative(native unsafe.Pointer) *AttrInt {
	return &AttrInt{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrInt) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrIntStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrInt) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrIntStruct, recv.native, "attr", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *AttrInt) FieldValue() int32 {
	argValue := gi.StructFieldGet(attrIntStruct, recv.native, "value")
	value := argValue.Int32()
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *AttrInt) SetFieldValue(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(attrIntStruct, recv.native, "value", argValue)
}

// AttrIntStruct creates an uninitialised AttrInt.
func AttrIntStruct() *AttrInt {
	err := attrIntStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrIntNewFromNative(attrIntStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrInt)
	return structGo
}
func finalizeAttrInt(obj *AttrInt) {
	attrIntStruct.Free(obj.native)
}

var attrIteratorStruct *gi.Struct
var attrIteratorStruct_Once sync.Once

func attrIteratorStruct_Set() error {
	var err error
	attrIteratorStruct_Once.Do(func() {
		attrIteratorStruct, err = gi.StructNew("Pango", "AttrIterator")
	})
	return err
}

type AttrIterator struct {
	native unsafe.Pointer
}

func AttrIteratorNewFromNative(native unsafe.Pointer) *AttrIterator {
	return &AttrIterator{native: native}
}

var attrIteratorCopyFunction *gi.Function
var attrIteratorCopyFunction_Once sync.Once

func attrIteratorCopyFunction_Set() error {
	var err error
	attrIteratorCopyFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorCopyFunction, err = attrIteratorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attr_iterator_copy.
func (recv *AttrIterator) Copy() *AttrIterator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrIteratorCopyFunction_Set()
	if err == nil {
		ret = attrIteratorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttrIteratorNewFromNative(ret.Pointer())

	return retGo
}

var attrIteratorDestroyFunction *gi.Function
var attrIteratorDestroyFunction_Once sync.Once

func attrIteratorDestroyFunction_Set() error {
	var err error
	attrIteratorDestroyFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorDestroyFunction, err = attrIteratorStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type pango_attr_iterator_destroy.
func (recv *AttrIterator) Destroy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrIteratorDestroyFunction_Set()
	if err == nil {
		attrIteratorDestroyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrIteratorGetFunction *gi.Function
var attrIteratorGetFunction_Once sync.Once

func attrIteratorGetFunction_Set() error {
	var err error
	attrIteratorGetFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorGetFunction, err = attrIteratorStruct.InvokerNew("get")
	})
	return err
}

// Get is a representation of the C type pango_attr_iterator_get.
func (recv *AttrIterator) Get(type_ AttrType) *Attribute {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(type_))

	var ret gi.Argument

	err := attrIteratorGetFunction_Set()
	if err == nil {
		ret = attrIteratorGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttributeNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_iterator_get_attrs' : return type 'GLib.SList' not supported

// UNSUPPORTED : C value 'pango_attr_iterator_get_font' : parameter 'extra_attrs' of type 'GLib.SList' not supported

var attrIteratorNextFunction *gi.Function
var attrIteratorNextFunction_Once sync.Once

func attrIteratorNextFunction_Set() error {
	var err error
	attrIteratorNextFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorNextFunction, err = attrIteratorStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type pango_attr_iterator_next.
func (recv *AttrIterator) Next() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrIteratorNextFunction_Set()
	if err == nil {
		ret = attrIteratorNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var attrIteratorRangeFunction *gi.Function
var attrIteratorRangeFunction_Once sync.Once

func attrIteratorRangeFunction_Set() error {
	var err error
	attrIteratorRangeFunction_Once.Do(func() {
		err = attrIteratorStruct_Set()
		if err != nil {
			return
		}
		attrIteratorRangeFunction, err = attrIteratorStruct.InvokerNew("range")
	})
	return err
}

// Range is a representation of the C type pango_attr_iterator_range.
func (recv *AttrIterator) Range() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := attrIteratorRangeFunction_Set()
	if err == nil {
		attrIteratorRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

// AttrIteratorStruct creates an uninitialised AttrIterator.
func AttrIteratorStruct() *AttrIterator {
	err := attrIteratorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrIteratorNewFromNative(attrIteratorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrIterator)
	return structGo
}
func finalizeAttrIterator(obj *AttrIterator) {
	attrIteratorStruct.Free(obj.native)
}

var attrLanguageStruct *gi.Struct
var attrLanguageStruct_Once sync.Once

func attrLanguageStruct_Set() error {
	var err error
	attrLanguageStruct_Once.Do(func() {
		attrLanguageStruct, err = gi.StructNew("Pango", "AttrLanguage")
	})
	return err
}

type AttrLanguage struct {
	native unsafe.Pointer
}

func AttrLanguageNewFromNative(native unsafe.Pointer) *AttrLanguage {
	return &AttrLanguage{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrLanguage) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrLanguageStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrLanguage) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrLanguageStruct, recv.native, "attr", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *AttrLanguage) FieldValue() *Language {
	argValue := gi.StructFieldGet(attrLanguageStruct, recv.native, "value")
	value := LanguageNewFromNative(argValue.Pointer())
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *AttrLanguage) SetFieldValue(value *Language) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrLanguageStruct, recv.native, "value", argValue)
}

// AttrLanguageStruct creates an uninitialised AttrLanguage.
func AttrLanguageStruct() *AttrLanguage {
	err := attrLanguageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrLanguageNewFromNative(attrLanguageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrLanguage)
	return structGo
}
func finalizeAttrLanguage(obj *AttrLanguage) {
	attrLanguageStruct.Free(obj.native)
}

var attrListStruct *gi.Struct
var attrListStruct_Once sync.Once

func attrListStruct_Set() error {
	var err error
	attrListStruct_Once.Do(func() {
		attrListStruct, err = gi.StructNew("Pango", "AttrList")
	})
	return err
}

type AttrList struct {
	native unsafe.Pointer
}

func AttrListNewFromNative(native unsafe.Pointer) *AttrList {
	return &AttrList{native: native}
}

var attrListNewFunction *gi.Function
var attrListNewFunction_Once sync.Once

func attrListNewFunction_Set() error {
	var err error
	attrListNewFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListNewFunction, err = attrListStruct.InvokerNew("new")
	})
	return err
}

// AttrListNew is a representation of the C type pango_attr_list_new.
func AttrListNew() *AttrList {

	var ret gi.Argument

	err := attrListNewFunction_Set()
	if err == nil {
		ret = attrListNewFunction.Invoke(nil, nil)
	}

	retGo := AttrListNewFromNative(ret.Pointer())

	return retGo
}

var attrListChangeFunction *gi.Function
var attrListChangeFunction_Once sync.Once

func attrListChangeFunction_Set() error {
	var err error
	attrListChangeFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListChangeFunction, err = attrListStruct.InvokerNew("change")
	})
	return err
}

// Change is a representation of the C type pango_attr_list_change.
func (recv *AttrList) Change(attr *Attribute) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListChangeFunction_Set()
	if err == nil {
		attrListChangeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrListCopyFunction *gi.Function
var attrListCopyFunction_Once sync.Once

func attrListCopyFunction_Set() error {
	var err error
	attrListCopyFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListCopyFunction, err = attrListStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attr_list_copy.
func (recv *AttrList) Copy() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListCopyFunction_Set()
	if err == nil {
		ret = attrListCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttrListNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_attr_list_filter' : parameter 'func' of type 'AttrFilterFunc' not supported

var attrListGetIteratorFunction *gi.Function
var attrListGetIteratorFunction_Once sync.Once

func attrListGetIteratorFunction_Set() error {
	var err error
	attrListGetIteratorFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListGetIteratorFunction, err = attrListStruct.InvokerNew("get_iterator")
	})
	return err
}

// GetIterator is a representation of the C type pango_attr_list_get_iterator.
func (recv *AttrList) GetIterator() *AttrIterator {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListGetIteratorFunction_Set()
	if err == nil {
		ret = attrListGetIteratorFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttrIteratorNewFromNative(ret.Pointer())

	return retGo
}

var attrListInsertFunction *gi.Function
var attrListInsertFunction_Once sync.Once

func attrListInsertFunction_Set() error {
	var err error
	attrListInsertFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListInsertFunction, err = attrListStruct.InvokerNew("insert")
	})
	return err
}

// Insert is a representation of the C type pango_attr_list_insert.
func (recv *AttrList) Insert(attr *Attribute) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListInsertFunction_Set()
	if err == nil {
		attrListInsertFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrListInsertBeforeFunction *gi.Function
var attrListInsertBeforeFunction_Once sync.Once

func attrListInsertBeforeFunction_Set() error {
	var err error
	attrListInsertBeforeFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListInsertBeforeFunction, err = attrListStruct.InvokerNew("insert_before")
	})
	return err
}

// InsertBefore is a representation of the C type pango_attr_list_insert_before.
func (recv *AttrList) InsertBefore(attr *Attribute) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr.native)

	err := attrListInsertBeforeFunction_Set()
	if err == nil {
		attrListInsertBeforeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrListRefFunction *gi.Function
var attrListRefFunction_Once sync.Once

func attrListRefFunction_Set() error {
	var err error
	attrListRefFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListRefFunction, err = attrListStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_attr_list_ref.
func (recv *AttrList) Ref() *AttrList {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attrListRefFunction_Set()
	if err == nil {
		ret = attrListRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttrListNewFromNative(ret.Pointer())

	return retGo
}

var attrListSpliceFunction *gi.Function
var attrListSpliceFunction_Once sync.Once

func attrListSpliceFunction_Set() error {
	var err error
	attrListSpliceFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListSpliceFunction, err = attrListStruct.InvokerNew("splice")
	})
	return err
}

// Splice is a representation of the C type pango_attr_list_splice.
func (recv *AttrList) Splice(other *AttrList, pos int32, len int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(other.native)
	inArgs[2].SetInt32(pos)
	inArgs[3].SetInt32(len)

	err := attrListSpliceFunction_Set()
	if err == nil {
		attrListSpliceFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrListUnrefFunction *gi.Function
var attrListUnrefFunction_Once sync.Once

func attrListUnrefFunction_Set() error {
	var err error
	attrListUnrefFunction_Once.Do(func() {
		err = attrListStruct_Set()
		if err != nil {
			return
		}
		attrListUnrefFunction, err = attrListStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_attr_list_unref.
func (recv *AttrList) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attrListUnrefFunction_Set()
	if err == nil {
		attrListUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attrShapeStruct *gi.Struct
var attrShapeStruct_Once sync.Once

func attrShapeStruct_Set() error {
	var err error
	attrShapeStruct_Once.Do(func() {
		attrShapeStruct, err = gi.StructNew("Pango", "AttrShape")
	})
	return err
}

type AttrShape struct {
	native unsafe.Pointer
}

func AttrShapeNewFromNative(native unsafe.Pointer) *AttrShape {
	return &AttrShape{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrShape) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrShapeStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrShape) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrShapeStruct, recv.native, "attr", argValue)
}

// FieldInkRect returns the C field 'ink_rect'.
func (recv *AttrShape) FieldInkRect() *Rectangle {
	argValue := gi.StructFieldGet(attrShapeStruct, recv.native, "ink_rect")
	value := RectangleNewFromNative(argValue.Pointer())
	return value
}

// SetFieldInkRect sets the value of the C field 'ink_rect'.
func (recv *AttrShape) SetFieldInkRect(value *Rectangle) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrShapeStruct, recv.native, "ink_rect", argValue)
}

// FieldLogicalRect returns the C field 'logical_rect'.
func (recv *AttrShape) FieldLogicalRect() *Rectangle {
	argValue := gi.StructFieldGet(attrShapeStruct, recv.native, "logical_rect")
	value := RectangleNewFromNative(argValue.Pointer())
	return value
}

// SetFieldLogicalRect sets the value of the C field 'logical_rect'.
func (recv *AttrShape) SetFieldLogicalRect(value *Rectangle) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrShapeStruct, recv.native, "logical_rect", argValue)
}

// UNSUPPORTED : C value 'data' : for field getter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'data' : for field setter : no Go type for 'gpointer'

// UNSUPPORTED : C value 'copy_func' : for field getter : no Go type for 'AttrDataCopyFunc'

// UNSUPPORTED : C value 'copy_func' : for field setter : no Go type for 'AttrDataCopyFunc'

// UNSUPPORTED : C value 'destroy_func' : for field getter : no Go type for 'GLib.DestroyNotify'

// UNSUPPORTED : C value 'destroy_func' : for field setter : no Go type for 'GLib.DestroyNotify'

// AttrShapeStruct creates an uninitialised AttrShape.
func AttrShapeStruct() *AttrShape {
	err := attrShapeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrShapeNewFromNative(attrShapeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrShape)
	return structGo
}
func finalizeAttrShape(obj *AttrShape) {
	attrShapeStruct.Free(obj.native)
}

var attrSizeStruct *gi.Struct
var attrSizeStruct_Once sync.Once

func attrSizeStruct_Set() error {
	var err error
	attrSizeStruct_Once.Do(func() {
		attrSizeStruct, err = gi.StructNew("Pango", "AttrSize")
	})
	return err
}

type AttrSize struct {
	native unsafe.Pointer
}

func AttrSizeNewFromNative(native unsafe.Pointer) *AttrSize {
	return &AttrSize{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrSize) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrSizeStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrSize) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrSizeStruct, recv.native, "attr", argValue)
}

// FieldSize returns the C field 'size'.
func (recv *AttrSize) FieldSize() int32 {
	argValue := gi.StructFieldGet(attrSizeStruct, recv.native, "size")
	value := argValue.Int32()
	return value
}

// SetFieldSize sets the value of the C field 'size'.
func (recv *AttrSize) SetFieldSize(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(attrSizeStruct, recv.native, "size", argValue)
}

// FieldAbsolute returns the C field 'absolute'.
func (recv *AttrSize) FieldAbsolute() uint32 {
	argValue := gi.StructFieldGet(attrSizeStruct, recv.native, "absolute")
	value := argValue.Uint32()
	return value
}

// SetFieldAbsolute sets the value of the C field 'absolute'.
func (recv *AttrSize) SetFieldAbsolute(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(attrSizeStruct, recv.native, "absolute", argValue)
}

// AttrSizeStruct creates an uninitialised AttrSize.
func AttrSizeStruct() *AttrSize {
	err := attrSizeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrSizeNewFromNative(attrSizeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrSize)
	return structGo
}
func finalizeAttrSize(obj *AttrSize) {
	attrSizeStruct.Free(obj.native)
}

var attrStringStruct *gi.Struct
var attrStringStruct_Once sync.Once

func attrStringStruct_Set() error {
	var err error
	attrStringStruct_Once.Do(func() {
		attrStringStruct, err = gi.StructNew("Pango", "AttrString")
	})
	return err
}

type AttrString struct {
	native unsafe.Pointer
}

func AttrStringNewFromNative(native unsafe.Pointer) *AttrString {
	return &AttrString{native: native}
}

// FieldAttr returns the C field 'attr'.
func (recv *AttrString) FieldAttr() *Attribute {
	argValue := gi.StructFieldGet(attrStringStruct, recv.native, "attr")
	value := AttributeNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *AttrString) SetFieldAttr(value *Attribute) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attrStringStruct, recv.native, "attr", argValue)
}

// FieldValue returns the C field 'value'.
func (recv *AttrString) FieldValue() string {
	argValue := gi.StructFieldGet(attrStringStruct, recv.native, "value")
	value := argValue.String(false)
	return value
}

// SetFieldValue sets the value of the C field 'value'.
func (recv *AttrString) SetFieldValue(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(attrStringStruct, recv.native, "value", argValue)
}

// AttrStringStruct creates an uninitialised AttrString.
func AttrStringStruct() *AttrString {
	err := attrStringStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttrStringNewFromNative(attrStringStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttrString)
	return structGo
}
func finalizeAttrString(obj *AttrString) {
	attrStringStruct.Free(obj.native)
}

var attributeStruct *gi.Struct
var attributeStruct_Once sync.Once

func attributeStruct_Set() error {
	var err error
	attributeStruct_Once.Do(func() {
		attributeStruct, err = gi.StructNew("Pango", "Attribute")
	})
	return err
}

type Attribute struct {
	native unsafe.Pointer
}

func AttributeNewFromNative(native unsafe.Pointer) *Attribute {
	return &Attribute{native: native}
}

// FieldKlass returns the C field 'klass'.
func (recv *Attribute) FieldKlass() *AttrClass {
	argValue := gi.StructFieldGet(attributeStruct, recv.native, "klass")
	value := AttrClassNewFromNative(argValue.Pointer())
	return value
}

// SetFieldKlass sets the value of the C field 'klass'.
func (recv *Attribute) SetFieldKlass(value *AttrClass) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(attributeStruct, recv.native, "klass", argValue)
}

// FieldStartIndex returns the C field 'start_index'.
func (recv *Attribute) FieldStartIndex() uint32 {
	argValue := gi.StructFieldGet(attributeStruct, recv.native, "start_index")
	value := argValue.Uint32()
	return value
}

// SetFieldStartIndex sets the value of the C field 'start_index'.
func (recv *Attribute) SetFieldStartIndex(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(attributeStruct, recv.native, "start_index", argValue)
}

// FieldEndIndex returns the C field 'end_index'.
func (recv *Attribute) FieldEndIndex() uint32 {
	argValue := gi.StructFieldGet(attributeStruct, recv.native, "end_index")
	value := argValue.Uint32()
	return value
}

// SetFieldEndIndex sets the value of the C field 'end_index'.
func (recv *Attribute) SetFieldEndIndex(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(attributeStruct, recv.native, "end_index", argValue)
}

var attributeCopyFunction *gi.Function
var attributeCopyFunction_Once sync.Once

func attributeCopyFunction_Set() error {
	var err error
	attributeCopyFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeCopyFunction, err = attributeStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_attribute_copy.
func (recv *Attribute) Copy() *Attribute {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := attributeCopyFunction_Set()
	if err == nil {
		ret = attributeCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := AttributeNewFromNative(ret.Pointer())

	return retGo
}

var attributeDestroyFunction *gi.Function
var attributeDestroyFunction_Once sync.Once

func attributeDestroyFunction_Set() error {
	var err error
	attributeDestroyFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeDestroyFunction, err = attributeStruct.InvokerNew("destroy")
	})
	return err
}

// Destroy is a representation of the C type pango_attribute_destroy.
func (recv *Attribute) Destroy() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := attributeDestroyFunction_Set()
	if err == nil {
		attributeDestroyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var attributeEqualFunction *gi.Function
var attributeEqualFunction_Once sync.Once

func attributeEqualFunction_Set() error {
	var err error
	attributeEqualFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeEqualFunction, err = attributeStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type pango_attribute_equal.
func (recv *Attribute) Equal(attr2 *Attribute) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(attr2.native)

	var ret gi.Argument

	err := attributeEqualFunction_Set()
	if err == nil {
		ret = attributeEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var attributeInitFunction *gi.Function
var attributeInitFunction_Once sync.Once

func attributeInitFunction_Set() error {
	var err error
	attributeInitFunction_Once.Do(func() {
		err = attributeStruct_Set()
		if err != nil {
			return
		}
		attributeInitFunction, err = attributeStruct.InvokerNew("init")
	})
	return err
}

// Init is a representation of the C type pango_attribute_init.
func (recv *Attribute) Init(klass *AttrClass) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(klass.native)

	err := attributeInitFunction_Set()
	if err == nil {
		attributeInitFunction.Invoke(inArgs[:], nil)
	}

	return
}

// AttributeStruct creates an uninitialised Attribute.
func AttributeStruct() *Attribute {
	err := attributeStruct_Set()
	if err != nil {
		return nil
	}

	structGo := AttributeNewFromNative(attributeStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeAttribute)
	return structGo
}
func finalizeAttribute(obj *Attribute) {
	attributeStruct.Free(obj.native)
}

var colorStruct *gi.Struct
var colorStruct_Once sync.Once

func colorStruct_Set() error {
	var err error
	colorStruct_Once.Do(func() {
		colorStruct, err = gi.StructNew("Pango", "Color")
	})
	return err
}

type Color struct {
	native unsafe.Pointer
}

func ColorNewFromNative(native unsafe.Pointer) *Color {
	return &Color{native: native}
}

// FieldRed returns the C field 'red'.
func (recv *Color) FieldRed() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "red")
	value := argValue.Uint16()
	return value
}

// SetFieldRed sets the value of the C field 'red'.
func (recv *Color) SetFieldRed(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "red", argValue)
}

// FieldGreen returns the C field 'green'.
func (recv *Color) FieldGreen() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "green")
	value := argValue.Uint16()
	return value
}

// SetFieldGreen sets the value of the C field 'green'.
func (recv *Color) SetFieldGreen(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "green", argValue)
}

// FieldBlue returns the C field 'blue'.
func (recv *Color) FieldBlue() uint16 {
	argValue := gi.StructFieldGet(colorStruct, recv.native, "blue")
	value := argValue.Uint16()
	return value
}

// SetFieldBlue sets the value of the C field 'blue'.
func (recv *Color) SetFieldBlue(value uint16) {
	var argValue gi.Argument
	argValue.SetUint16(value)
	gi.StructFieldSet(colorStruct, recv.native, "blue", argValue)
}

var colorCopyFunction *gi.Function
var colorCopyFunction_Once sync.Once

func colorCopyFunction_Set() error {
	var err error
	colorCopyFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorCopyFunction, err = colorStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_color_copy.
func (recv *Color) Copy() *Color {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorCopyFunction_Set()
	if err == nil {
		ret = colorCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ColorNewFromNative(ret.Pointer())

	return retGo
}

var colorFreeFunction *gi.Function
var colorFreeFunction_Once sync.Once

func colorFreeFunction_Set() error {
	var err error
	colorFreeFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorFreeFunction, err = colorStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_color_free.
func (recv *Color) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := colorFreeFunction_Set()
	if err == nil {
		colorFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var colorParseFunction *gi.Function
var colorParseFunction_Once sync.Once

func colorParseFunction_Set() error {
	var err error
	colorParseFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorParseFunction, err = colorStruct.InvokerNew("parse")
	})
	return err
}

// Parse is a representation of the C type pango_color_parse.
func (recv *Color) Parse(spec string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(spec)

	var ret gi.Argument

	err := colorParseFunction_Set()
	if err == nil {
		ret = colorParseFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var colorToStringFunction *gi.Function
var colorToStringFunction_Once sync.Once

func colorToStringFunction_Set() error {
	var err error
	colorToStringFunction_Once.Do(func() {
		err = colorStruct_Set()
		if err != nil {
			return
		}
		colorToStringFunction, err = colorStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_color_to_string.
func (recv *Color) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := colorToStringFunction_Set()
	if err == nil {
		ret = colorToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// ColorStruct creates an uninitialised Color.
func ColorStruct() *Color {
	err := colorStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ColorNewFromNative(colorStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeColor)
	return structGo
}
func finalizeColor(obj *Color) {
	colorStruct.Free(obj.native)
}

var contextClassStruct *gi.Struct
var contextClassStruct_Once sync.Once

func contextClassStruct_Set() error {
	var err error
	contextClassStruct_Once.Do(func() {
		contextClassStruct, err = gi.StructNew("Pango", "ContextClass")
	})
	return err
}

type ContextClass struct {
	native unsafe.Pointer
}

func ContextClassNewFromNative(native unsafe.Pointer) *ContextClass {
	return &ContextClass{native: native}
}

// ContextClassStruct creates an uninitialised ContextClass.
func ContextClassStruct() *ContextClass {
	err := contextClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ContextClassNewFromNative(contextClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeContextClass)
	return structGo
}
func finalizeContextClass(obj *ContextClass) {
	contextClassStruct.Free(obj.native)
}

var coverageStruct *gi.Struct
var coverageStruct_Once sync.Once

func coverageStruct_Set() error {
	var err error
	coverageStruct_Once.Do(func() {
		coverageStruct, err = gi.StructNew("Pango", "Coverage")
	})
	return err
}

type Coverage struct {
	native unsafe.Pointer
}

func CoverageNewFromNative(native unsafe.Pointer) *Coverage {
	return &Coverage{native: native}
}

var coverageCopyFunction *gi.Function
var coverageCopyFunction_Once sync.Once

func coverageCopyFunction_Set() error {
	var err error
	coverageCopyFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageCopyFunction, err = coverageStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_coverage_copy.
func (recv *Coverage) Copy() *Coverage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageCopyFunction_Set()
	if err == nil {
		ret = coverageCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := CoverageNewFromNative(ret.Pointer())

	return retGo
}

var coverageGetFunction *gi.Function
var coverageGetFunction_Once sync.Once

func coverageGetFunction_Set() error {
	var err error
	coverageGetFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageGetFunction, err = coverageStruct.InvokerNew("get")
	})
	return err
}

// Get is a representation of the C type pango_coverage_get.
func (recv *Coverage) Get(index int32) CoverageLevel {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)

	var ret gi.Argument

	err := coverageGetFunction_Set()
	if err == nil {
		ret = coverageGetFunction.Invoke(inArgs[:], nil)
	}

	retGo := CoverageLevel(ret.Int32())

	return retGo
}

var coverageMaxFunction *gi.Function
var coverageMaxFunction_Once sync.Once

func coverageMaxFunction_Set() error {
	var err error
	coverageMaxFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageMaxFunction, err = coverageStruct.InvokerNew("max")
	})
	return err
}

// Max is a representation of the C type pango_coverage_max.
func (recv *Coverage) Max(other *Coverage) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(other.native)

	err := coverageMaxFunction_Set()
	if err == nil {
		coverageMaxFunction.Invoke(inArgs[:], nil)
	}

	return
}

var coverageRefFunction *gi.Function
var coverageRefFunction_Once sync.Once

func coverageRefFunction_Set() error {
	var err error
	coverageRefFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageRefFunction, err = coverageStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_coverage_ref.
func (recv *Coverage) Ref() *Coverage {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := coverageRefFunction_Set()
	if err == nil {
		ret = coverageRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := CoverageNewFromNative(ret.Pointer())

	return retGo
}

var coverageSetFunction *gi.Function
var coverageSetFunction_Once sync.Once

func coverageSetFunction_Set() error {
	var err error
	coverageSetFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageSetFunction, err = coverageStruct.InvokerNew("set")
	})
	return err
}

// Set is a representation of the C type pango_coverage_set.
func (recv *Coverage) Set(index int32, level CoverageLevel) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)
	inArgs[2].SetInt32(int32(level))

	err := coverageSetFunction_Set()
	if err == nil {
		coverageSetFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_coverage_to_bytes' : parameter 'bytes' of type 'nil' not supported

var coverageUnrefFunction *gi.Function
var coverageUnrefFunction_Once sync.Once

func coverageUnrefFunction_Set() error {
	var err error
	coverageUnrefFunction_Once.Do(func() {
		err = coverageStruct_Set()
		if err != nil {
			return
		}
		coverageUnrefFunction, err = coverageStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_coverage_unref.
func (recv *Coverage) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := coverageUnrefFunction_Set()
	if err == nil {
		coverageUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

// CoverageStruct creates an uninitialised Coverage.
func CoverageStruct() *Coverage {
	err := coverageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := CoverageNewFromNative(coverageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeCoverage)
	return structGo
}
func finalizeCoverage(obj *Coverage) {
	coverageStruct.Free(obj.native)
}

var engineClassStruct *gi.Struct
var engineClassStruct_Once sync.Once

func engineClassStruct_Set() error {
	var err error
	engineClassStruct_Once.Do(func() {
		engineClassStruct, err = gi.StructNew("Pango", "EngineClass")
	})
	return err
}

type EngineClass struct {
	native unsafe.Pointer
}

func EngineClassNewFromNative(native unsafe.Pointer) *EngineClass {
	return &EngineClass{native: native}
}

// EngineClassStruct creates an uninitialised EngineClass.
func EngineClassStruct() *EngineClass {
	err := engineClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EngineClassNewFromNative(engineClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEngineClass)
	return structGo
}
func finalizeEngineClass(obj *EngineClass) {
	engineClassStruct.Free(obj.native)
}

var engineInfoStruct *gi.Struct
var engineInfoStruct_Once sync.Once

func engineInfoStruct_Set() error {
	var err error
	engineInfoStruct_Once.Do(func() {
		engineInfoStruct, err = gi.StructNew("Pango", "EngineInfo")
	})
	return err
}

type EngineInfo struct {
	native unsafe.Pointer
}

func EngineInfoNewFromNative(native unsafe.Pointer) *EngineInfo {
	return &EngineInfo{native: native}
}

// FieldId returns the C field 'id'.
func (recv *EngineInfo) FieldId() string {
	argValue := gi.StructFieldGet(engineInfoStruct, recv.native, "id")
	value := argValue.String(false)
	return value
}

// SetFieldId sets the value of the C field 'id'.
func (recv *EngineInfo) SetFieldId(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(engineInfoStruct, recv.native, "id", argValue)
}

// FieldEngineType returns the C field 'engine_type'.
func (recv *EngineInfo) FieldEngineType() string {
	argValue := gi.StructFieldGet(engineInfoStruct, recv.native, "engine_type")
	value := argValue.String(false)
	return value
}

// SetFieldEngineType sets the value of the C field 'engine_type'.
func (recv *EngineInfo) SetFieldEngineType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(engineInfoStruct, recv.native, "engine_type", argValue)
}

// FieldRenderType returns the C field 'render_type'.
func (recv *EngineInfo) FieldRenderType() string {
	argValue := gi.StructFieldGet(engineInfoStruct, recv.native, "render_type")
	value := argValue.String(false)
	return value
}

// SetFieldRenderType sets the value of the C field 'render_type'.
func (recv *EngineInfo) SetFieldRenderType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(engineInfoStruct, recv.native, "render_type", argValue)
}

// FieldScripts returns the C field 'scripts'.
func (recv *EngineInfo) FieldScripts() *EngineScriptInfo {
	argValue := gi.StructFieldGet(engineInfoStruct, recv.native, "scripts")
	value := EngineScriptInfoNewFromNative(argValue.Pointer())
	return value
}

// SetFieldScripts sets the value of the C field 'scripts'.
func (recv *EngineInfo) SetFieldScripts(value *EngineScriptInfo) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(engineInfoStruct, recv.native, "scripts", argValue)
}

// FieldNScripts returns the C field 'n_scripts'.
func (recv *EngineInfo) FieldNScripts() int32 {
	argValue := gi.StructFieldGet(engineInfoStruct, recv.native, "n_scripts")
	value := argValue.Int32()
	return value
}

// SetFieldNScripts sets the value of the C field 'n_scripts'.
func (recv *EngineInfo) SetFieldNScripts(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(engineInfoStruct, recv.native, "n_scripts", argValue)
}

// EngineInfoStruct creates an uninitialised EngineInfo.
func EngineInfoStruct() *EngineInfo {
	err := engineInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EngineInfoNewFromNative(engineInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEngineInfo)
	return structGo
}
func finalizeEngineInfo(obj *EngineInfo) {
	engineInfoStruct.Free(obj.native)
}

var engineLangClassStruct *gi.Struct
var engineLangClassStruct_Once sync.Once

func engineLangClassStruct_Set() error {
	var err error
	engineLangClassStruct_Once.Do(func() {
		engineLangClassStruct, err = gi.StructNew("Pango", "EngineLangClass")
	})
	return err
}

type EngineLangClass struct {
	native unsafe.Pointer
}

func EngineLangClassNewFromNative(native unsafe.Pointer) *EngineLangClass {
	return &EngineLangClass{native: native}
}

// UNSUPPORTED : C value 'script_break' : for field getter : missing Type

// UNSUPPORTED : C value 'script_break' : for field setter : missing Type

// EngineLangClassStruct creates an uninitialised EngineLangClass.
func EngineLangClassStruct() *EngineLangClass {
	err := engineLangClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EngineLangClassNewFromNative(engineLangClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEngineLangClass)
	return structGo
}
func finalizeEngineLangClass(obj *EngineLangClass) {
	engineLangClassStruct.Free(obj.native)
}

var engineScriptInfoStruct *gi.Struct
var engineScriptInfoStruct_Once sync.Once

func engineScriptInfoStruct_Set() error {
	var err error
	engineScriptInfoStruct_Once.Do(func() {
		engineScriptInfoStruct, err = gi.StructNew("Pango", "EngineScriptInfo")
	})
	return err
}

type EngineScriptInfo struct {
	native unsafe.Pointer
}

func EngineScriptInfoNewFromNative(native unsafe.Pointer) *EngineScriptInfo {
	return &EngineScriptInfo{native: native}
}

// FieldScript returns the C field 'script'.
func (recv *EngineScriptInfo) FieldScript() Script {
	argValue := gi.StructFieldGet(engineScriptInfoStruct, recv.native, "script")
	value := Script(argValue.Int32())
	return value
}

// SetFieldScript sets the value of the C field 'script'.
func (recv *EngineScriptInfo) SetFieldScript(value Script) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(engineScriptInfoStruct, recv.native, "script", argValue)
}

// FieldLangs returns the C field 'langs'.
func (recv *EngineScriptInfo) FieldLangs() string {
	argValue := gi.StructFieldGet(engineScriptInfoStruct, recv.native, "langs")
	value := argValue.String(false)
	return value
}

// SetFieldLangs sets the value of the C field 'langs'.
func (recv *EngineScriptInfo) SetFieldLangs(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(engineScriptInfoStruct, recv.native, "langs", argValue)
}

// EngineScriptInfoStruct creates an uninitialised EngineScriptInfo.
func EngineScriptInfoStruct() *EngineScriptInfo {
	err := engineScriptInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EngineScriptInfoNewFromNative(engineScriptInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEngineScriptInfo)
	return structGo
}
func finalizeEngineScriptInfo(obj *EngineScriptInfo) {
	engineScriptInfoStruct.Free(obj.native)
}

var engineShapeClassStruct *gi.Struct
var engineShapeClassStruct_Once sync.Once

func engineShapeClassStruct_Set() error {
	var err error
	engineShapeClassStruct_Once.Do(func() {
		engineShapeClassStruct, err = gi.StructNew("Pango", "EngineShapeClass")
	})
	return err
}

type EngineShapeClass struct {
	native unsafe.Pointer
}

func EngineShapeClassNewFromNative(native unsafe.Pointer) *EngineShapeClass {
	return &EngineShapeClass{native: native}
}

// UNSUPPORTED : C value 'script_shape' : for field getter : missing Type

// UNSUPPORTED : C value 'script_shape' : for field setter : missing Type

// UNSUPPORTED : C value 'covers' : for field getter : missing Type

// UNSUPPORTED : C value 'covers' : for field setter : missing Type

// EngineShapeClassStruct creates an uninitialised EngineShapeClass.
func EngineShapeClassStruct() *EngineShapeClass {
	err := engineShapeClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := EngineShapeClassNewFromNative(engineShapeClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeEngineShapeClass)
	return structGo
}
func finalizeEngineShapeClass(obj *EngineShapeClass) {
	engineShapeClassStruct.Free(obj.native)
}

var fontClassStruct *gi.Struct
var fontClassStruct_Once sync.Once

func fontClassStruct_Set() error {
	var err error
	fontClassStruct_Once.Do(func() {
		fontClassStruct, err = gi.StructNew("Pango", "FontClass")
	})
	return err
}

type FontClass struct {
	native unsafe.Pointer
}

func FontClassNewFromNative(native unsafe.Pointer) *FontClass {
	return &FontClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'describe' : for field getter : missing Type

// UNSUPPORTED : C value 'describe' : for field setter : missing Type

// UNSUPPORTED : C value 'get_coverage' : for field getter : missing Type

// UNSUPPORTED : C value 'get_coverage' : for field setter : missing Type

// UNSUPPORTED : C value 'find_shaper' : for field getter : missing Type

// UNSUPPORTED : C value 'find_shaper' : for field setter : missing Type

// UNSUPPORTED : C value 'get_glyph_extents' : for field getter : missing Type

// UNSUPPORTED : C value 'get_glyph_extents' : for field setter : missing Type

// UNSUPPORTED : C value 'get_metrics' : for field getter : missing Type

// UNSUPPORTED : C value 'get_metrics' : for field setter : missing Type

// UNSUPPORTED : C value 'get_font_map' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font_map' : for field setter : missing Type

// UNSUPPORTED : C value 'describe_absolute' : for field getter : missing Type

// UNSUPPORTED : C value 'describe_absolute' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field setter : missing Type

// FontClassStruct creates an uninitialised FontClass.
func FontClassStruct() *FontClass {
	err := fontClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontClassNewFromNative(fontClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontClass)
	return structGo
}
func finalizeFontClass(obj *FontClass) {
	fontClassStruct.Free(obj.native)
}

var fontDescriptionStruct *gi.Struct
var fontDescriptionStruct_Once sync.Once

func fontDescriptionStruct_Set() error {
	var err error
	fontDescriptionStruct_Once.Do(func() {
		fontDescriptionStruct, err = gi.StructNew("Pango", "FontDescription")
	})
	return err
}

type FontDescription struct {
	native unsafe.Pointer
}

func FontDescriptionNewFromNative(native unsafe.Pointer) *FontDescription {
	return &FontDescription{native: native}
}

var fontDescriptionNewFunction *gi.Function
var fontDescriptionNewFunction_Once sync.Once

func fontDescriptionNewFunction_Set() error {
	var err error
	fontDescriptionNewFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionNewFunction, err = fontDescriptionStruct.InvokerNew("new")
	})
	return err
}

// FontDescriptionNew is a representation of the C type pango_font_description_new.
func FontDescriptionNew() *FontDescription {

	var ret gi.Argument

	err := fontDescriptionNewFunction_Set()
	if err == nil {
		ret = fontDescriptionNewFunction.Invoke(nil, nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontDescriptionBetterMatchFunction *gi.Function
var fontDescriptionBetterMatchFunction_Once sync.Once

func fontDescriptionBetterMatchFunction_Set() error {
	var err error
	fontDescriptionBetterMatchFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionBetterMatchFunction, err = fontDescriptionStruct.InvokerNew("better_match")
	})
	return err
}

// BetterMatch is a representation of the C type pango_font_description_better_match.
func (recv *FontDescription) BetterMatch(oldMatch *FontDescription, newMatch *FontDescription) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(oldMatch.native)
	inArgs[2].SetPointer(newMatch.native)

	var ret gi.Argument

	err := fontDescriptionBetterMatchFunction_Set()
	if err == nil {
		ret = fontDescriptionBetterMatchFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fontDescriptionCopyFunction *gi.Function
var fontDescriptionCopyFunction_Once sync.Once

func fontDescriptionCopyFunction_Set() error {
	var err error
	fontDescriptionCopyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionCopyFunction, err = fontDescriptionStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_font_description_copy.
func (recv *FontDescription) Copy() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontDescriptionCopyStaticFunction *gi.Function
var fontDescriptionCopyStaticFunction_Once sync.Once

func fontDescriptionCopyStaticFunction_Set() error {
	var err error
	fontDescriptionCopyStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionCopyStaticFunction, err = fontDescriptionStruct.InvokerNew("copy_static")
	})
	return err
}

// CopyStatic is a representation of the C type pango_font_description_copy_static.
func (recv *FontDescription) CopyStatic() *FontDescription {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionCopyStaticFunction_Set()
	if err == nil {
		ret = fontDescriptionCopyStaticFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontDescriptionNewFromNative(ret.Pointer())

	return retGo
}

var fontDescriptionEqualFunction *gi.Function
var fontDescriptionEqualFunction_Once sync.Once

func fontDescriptionEqualFunction_Set() error {
	var err error
	fontDescriptionEqualFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionEqualFunction, err = fontDescriptionStruct.InvokerNew("equal")
	})
	return err
}

// Equal is a representation of the C type pango_font_description_equal.
func (recv *FontDescription) Equal(desc2 *FontDescription) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(desc2.native)

	var ret gi.Argument

	err := fontDescriptionEqualFunction_Set()
	if err == nil {
		ret = fontDescriptionEqualFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fontDescriptionFreeFunction *gi.Function
var fontDescriptionFreeFunction_Once sync.Once

func fontDescriptionFreeFunction_Set() error {
	var err error
	fontDescriptionFreeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionFreeFunction, err = fontDescriptionStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_font_description_free.
func (recv *FontDescription) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontDescriptionFreeFunction_Set()
	if err == nil {
		fontDescriptionFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionGetFamilyFunction *gi.Function
var fontDescriptionGetFamilyFunction_Once sync.Once

func fontDescriptionGetFamilyFunction_Set() error {
	var err error
	fontDescriptionGetFamilyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetFamilyFunction, err = fontDescriptionStruct.InvokerNew("get_family")
	})
	return err
}

// GetFamily is a representation of the C type pango_font_description_get_family.
func (recv *FontDescription) GetFamily() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetFamilyFunction_Set()
	if err == nil {
		ret = fontDescriptionGetFamilyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fontDescriptionGetGravityFunction *gi.Function
var fontDescriptionGetGravityFunction_Once sync.Once

func fontDescriptionGetGravityFunction_Set() error {
	var err error
	fontDescriptionGetGravityFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetGravityFunction, err = fontDescriptionStruct.InvokerNew("get_gravity")
	})
	return err
}

// GetGravity is a representation of the C type pango_font_description_get_gravity.
func (recv *FontDescription) GetGravity() Gravity {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetGravityFunction_Set()
	if err == nil {
		ret = fontDescriptionGetGravityFunction.Invoke(inArgs[:], nil)
	}

	retGo := Gravity(ret.Int32())

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_get_set_fields' : return type 'FontMask' not supported

var fontDescriptionGetSizeFunction *gi.Function
var fontDescriptionGetSizeFunction_Once sync.Once

func fontDescriptionGetSizeFunction_Set() error {
	var err error
	fontDescriptionGetSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetSizeFunction, err = fontDescriptionStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_font_description_get_size.
func (recv *FontDescription) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetSizeFunction_Set()
	if err == nil {
		ret = fontDescriptionGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontDescriptionGetSizeIsAbsoluteFunction *gi.Function
var fontDescriptionGetSizeIsAbsoluteFunction_Once sync.Once

func fontDescriptionGetSizeIsAbsoluteFunction_Set() error {
	var err error
	fontDescriptionGetSizeIsAbsoluteFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetSizeIsAbsoluteFunction, err = fontDescriptionStruct.InvokerNew("get_size_is_absolute")
	})
	return err
}

// GetSizeIsAbsolute is a representation of the C type pango_font_description_get_size_is_absolute.
func (recv *FontDescription) GetSizeIsAbsolute() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetSizeIsAbsoluteFunction_Set()
	if err == nil {
		ret = fontDescriptionGetSizeIsAbsoluteFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var fontDescriptionGetStretchFunction *gi.Function
var fontDescriptionGetStretchFunction_Once sync.Once

func fontDescriptionGetStretchFunction_Set() error {
	var err error
	fontDescriptionGetStretchFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetStretchFunction, err = fontDescriptionStruct.InvokerNew("get_stretch")
	})
	return err
}

// GetStretch is a representation of the C type pango_font_description_get_stretch.
func (recv *FontDescription) GetStretch() Stretch {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetStretchFunction_Set()
	if err == nil {
		ret = fontDescriptionGetStretchFunction.Invoke(inArgs[:], nil)
	}

	retGo := Stretch(ret.Int32())

	return retGo
}

var fontDescriptionGetStyleFunction *gi.Function
var fontDescriptionGetStyleFunction_Once sync.Once

func fontDescriptionGetStyleFunction_Set() error {
	var err error
	fontDescriptionGetStyleFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetStyleFunction, err = fontDescriptionStruct.InvokerNew("get_style")
	})
	return err
}

// GetStyle is a representation of the C type pango_font_description_get_style.
func (recv *FontDescription) GetStyle() Style {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetStyleFunction_Set()
	if err == nil {
		ret = fontDescriptionGetStyleFunction.Invoke(inArgs[:], nil)
	}

	retGo := Style(ret.Int32())

	return retGo
}

var fontDescriptionGetVariantFunction *gi.Function
var fontDescriptionGetVariantFunction_Once sync.Once

func fontDescriptionGetVariantFunction_Set() error {
	var err error
	fontDescriptionGetVariantFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetVariantFunction, err = fontDescriptionStruct.InvokerNew("get_variant")
	})
	return err
}

// GetVariant is a representation of the C type pango_font_description_get_variant.
func (recv *FontDescription) GetVariant() Variant {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetVariantFunction_Set()
	if err == nil {
		ret = fontDescriptionGetVariantFunction.Invoke(inArgs[:], nil)
	}

	retGo := Variant(ret.Int32())

	return retGo
}

var fontDescriptionGetVariationsFunction *gi.Function
var fontDescriptionGetVariationsFunction_Once sync.Once

func fontDescriptionGetVariationsFunction_Set() error {
	var err error
	fontDescriptionGetVariationsFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetVariationsFunction, err = fontDescriptionStruct.InvokerNew("get_variations")
	})
	return err
}

// GetVariations is a representation of the C type pango_font_description_get_variations.
func (recv *FontDescription) GetVariations() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetVariationsFunction_Set()
	if err == nil {
		ret = fontDescriptionGetVariationsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var fontDescriptionGetWeightFunction *gi.Function
var fontDescriptionGetWeightFunction_Once sync.Once

func fontDescriptionGetWeightFunction_Set() error {
	var err error
	fontDescriptionGetWeightFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionGetWeightFunction, err = fontDescriptionStruct.InvokerNew("get_weight")
	})
	return err
}

// GetWeight is a representation of the C type pango_font_description_get_weight.
func (recv *FontDescription) GetWeight() Weight {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionGetWeightFunction_Set()
	if err == nil {
		ret = fontDescriptionGetWeightFunction.Invoke(inArgs[:], nil)
	}

	retGo := Weight(ret.Int32())

	return retGo
}

var fontDescriptionHashFunction *gi.Function
var fontDescriptionHashFunction_Once sync.Once

func fontDescriptionHashFunction_Set() error {
	var err error
	fontDescriptionHashFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionHashFunction, err = fontDescriptionStruct.InvokerNew("hash")
	})
	return err
}

// Hash is a representation of the C type pango_font_description_hash.
func (recv *FontDescription) Hash() uint32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionHashFunction_Set()
	if err == nil {
		ret = fontDescriptionHashFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Uint32()

	return retGo
}

var fontDescriptionMergeFunction *gi.Function
var fontDescriptionMergeFunction_Once sync.Once

func fontDescriptionMergeFunction_Set() error {
	var err error
	fontDescriptionMergeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionMergeFunction, err = fontDescriptionStruct.InvokerNew("merge")
	})
	return err
}

// Merge is a representation of the C type pango_font_description_merge.
func (recv *FontDescription) Merge(descToMerge *FontDescription, replaceExisting bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(descToMerge.native)
	inArgs[2].SetBoolean(replaceExisting)

	err := fontDescriptionMergeFunction_Set()
	if err == nil {
		fontDescriptionMergeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionMergeStaticFunction *gi.Function
var fontDescriptionMergeStaticFunction_Once sync.Once

func fontDescriptionMergeStaticFunction_Set() error {
	var err error
	fontDescriptionMergeStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionMergeStaticFunction, err = fontDescriptionStruct.InvokerNew("merge_static")
	})
	return err
}

// MergeStatic is a representation of the C type pango_font_description_merge_static.
func (recv *FontDescription) MergeStatic(descToMerge *FontDescription, replaceExisting bool) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(descToMerge.native)
	inArgs[2].SetBoolean(replaceExisting)

	err := fontDescriptionMergeStaticFunction_Set()
	if err == nil {
		fontDescriptionMergeStaticFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetAbsoluteSizeFunction *gi.Function
var fontDescriptionSetAbsoluteSizeFunction_Once sync.Once

func fontDescriptionSetAbsoluteSizeFunction_Set() error {
	var err error
	fontDescriptionSetAbsoluteSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetAbsoluteSizeFunction, err = fontDescriptionStruct.InvokerNew("set_absolute_size")
	})
	return err
}

// SetAbsoluteSize is a representation of the C type pango_font_description_set_absolute_size.
func (recv *FontDescription) SetAbsoluteSize(size float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(size)

	err := fontDescriptionSetAbsoluteSizeFunction_Set()
	if err == nil {
		fontDescriptionSetAbsoluteSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetFamilyFunction *gi.Function
var fontDescriptionSetFamilyFunction_Once sync.Once

func fontDescriptionSetFamilyFunction_Set() error {
	var err error
	fontDescriptionSetFamilyFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetFamilyFunction, err = fontDescriptionStruct.InvokerNew("set_family")
	})
	return err
}

// SetFamily is a representation of the C type pango_font_description_set_family.
func (recv *FontDescription) SetFamily(family string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetFamilyStaticFunction *gi.Function
var fontDescriptionSetFamilyStaticFunction_Once sync.Once

func fontDescriptionSetFamilyStaticFunction_Set() error {
	var err error
	fontDescriptionSetFamilyStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetFamilyStaticFunction, err = fontDescriptionStruct.InvokerNew("set_family_static")
	})
	return err
}

// SetFamilyStatic is a representation of the C type pango_font_description_set_family_static.
func (recv *FontDescription) SetFamilyStatic(family string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(family)

	err := fontDescriptionSetFamilyStaticFunction_Set()
	if err == nil {
		fontDescriptionSetFamilyStaticFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetGravityFunction *gi.Function
var fontDescriptionSetGravityFunction_Once sync.Once

func fontDescriptionSetGravityFunction_Set() error {
	var err error
	fontDescriptionSetGravityFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetGravityFunction, err = fontDescriptionStruct.InvokerNew("set_gravity")
	})
	return err
}

// SetGravity is a representation of the C type pango_font_description_set_gravity.
func (recv *FontDescription) SetGravity(gravity Gravity) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(gravity))

	err := fontDescriptionSetGravityFunction_Set()
	if err == nil {
		fontDescriptionSetGravityFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetSizeFunction *gi.Function
var fontDescriptionSetSizeFunction_Once sync.Once

func fontDescriptionSetSizeFunction_Set() error {
	var err error
	fontDescriptionSetSizeFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetSizeFunction, err = fontDescriptionStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type pango_font_description_set_size.
func (recv *FontDescription) SetSize(size int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(size)

	err := fontDescriptionSetSizeFunction_Set()
	if err == nil {
		fontDescriptionSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetStretchFunction *gi.Function
var fontDescriptionSetStretchFunction_Once sync.Once

func fontDescriptionSetStretchFunction_Set() error {
	var err error
	fontDescriptionSetStretchFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetStretchFunction, err = fontDescriptionStruct.InvokerNew("set_stretch")
	})
	return err
}

// SetStretch is a representation of the C type pango_font_description_set_stretch.
func (recv *FontDescription) SetStretch(stretch Stretch) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(stretch))

	err := fontDescriptionSetStretchFunction_Set()
	if err == nil {
		fontDescriptionSetStretchFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetStyleFunction *gi.Function
var fontDescriptionSetStyleFunction_Once sync.Once

func fontDescriptionSetStyleFunction_Set() error {
	var err error
	fontDescriptionSetStyleFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetStyleFunction, err = fontDescriptionStruct.InvokerNew("set_style")
	})
	return err
}

// SetStyle is a representation of the C type pango_font_description_set_style.
func (recv *FontDescription) SetStyle(style Style) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(style))

	err := fontDescriptionSetStyleFunction_Set()
	if err == nil {
		fontDescriptionSetStyleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetVariantFunction *gi.Function
var fontDescriptionSetVariantFunction_Once sync.Once

func fontDescriptionSetVariantFunction_Set() error {
	var err error
	fontDescriptionSetVariantFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetVariantFunction, err = fontDescriptionStruct.InvokerNew("set_variant")
	})
	return err
}

// SetVariant is a representation of the C type pango_font_description_set_variant.
func (recv *FontDescription) SetVariant(variant Variant) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(variant))

	err := fontDescriptionSetVariantFunction_Set()
	if err == nil {
		fontDescriptionSetVariantFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetVariationsFunction *gi.Function
var fontDescriptionSetVariationsFunction_Once sync.Once

func fontDescriptionSetVariationsFunction_Set() error {
	var err error
	fontDescriptionSetVariationsFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetVariationsFunction, err = fontDescriptionStruct.InvokerNew("set_variations")
	})
	return err
}

// SetVariations is a representation of the C type pango_font_description_set_variations.
func (recv *FontDescription) SetVariations(settings string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetVariationsStaticFunction *gi.Function
var fontDescriptionSetVariationsStaticFunction_Once sync.Once

func fontDescriptionSetVariationsStaticFunction_Set() error {
	var err error
	fontDescriptionSetVariationsStaticFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetVariationsStaticFunction, err = fontDescriptionStruct.InvokerNew("set_variations_static")
	})
	return err
}

// SetVariationsStatic is a representation of the C type pango_font_description_set_variations_static.
func (recv *FontDescription) SetVariationsStatic(settings string) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(settings)

	err := fontDescriptionSetVariationsStaticFunction_Set()
	if err == nil {
		fontDescriptionSetVariationsStaticFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionSetWeightFunction *gi.Function
var fontDescriptionSetWeightFunction_Once sync.Once

func fontDescriptionSetWeightFunction_Set() error {
	var err error
	fontDescriptionSetWeightFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionSetWeightFunction, err = fontDescriptionStruct.InvokerNew("set_weight")
	})
	return err
}

// SetWeight is a representation of the C type pango_font_description_set_weight.
func (recv *FontDescription) SetWeight(weight Weight) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(weight))

	err := fontDescriptionSetWeightFunction_Set()
	if err == nil {
		fontDescriptionSetWeightFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontDescriptionToFilenameFunction *gi.Function
var fontDescriptionToFilenameFunction_Once sync.Once

func fontDescriptionToFilenameFunction_Set() error {
	var err error
	fontDescriptionToFilenameFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionToFilenameFunction, err = fontDescriptionStruct.InvokerNew("to_filename")
	})
	return err
}

// ToFilename is a representation of the C type pango_font_description_to_filename.
func (recv *FontDescription) ToFilename() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToFilenameFunction_Set()
	if err == nil {
		ret = fontDescriptionToFilenameFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

var fontDescriptionToStringFunction *gi.Function
var fontDescriptionToStringFunction_Once sync.Once

func fontDescriptionToStringFunction_Set() error {
	var err error
	fontDescriptionToStringFunction_Once.Do(func() {
		err = fontDescriptionStruct_Set()
		if err != nil {
			return
		}
		fontDescriptionToStringFunction, err = fontDescriptionStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_font_description_to_string.
func (recv *FontDescription) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontDescriptionToStringFunction_Set()
	if err == nil {
		ret = fontDescriptionToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(true)

	return retGo
}

// UNSUPPORTED : C value 'pango_font_description_unset_fields' : parameter 'to_unset' of type 'FontMask' not supported

var fontFaceClassStruct *gi.Struct
var fontFaceClassStruct_Once sync.Once

func fontFaceClassStruct_Set() error {
	var err error
	fontFaceClassStruct_Once.Do(func() {
		fontFaceClassStruct, err = gi.StructNew("Pango", "FontFaceClass")
	})
	return err
}

type FontFaceClass struct {
	native unsafe.Pointer
}

func FontFaceClassNewFromNative(native unsafe.Pointer) *FontFaceClass {
	return &FontFaceClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_face_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_face_name' : for field setter : missing Type

// UNSUPPORTED : C value 'describe' : for field getter : missing Type

// UNSUPPORTED : C value 'describe' : for field setter : missing Type

// UNSUPPORTED : C value 'list_sizes' : for field getter : missing Type

// UNSUPPORTED : C value 'list_sizes' : for field setter : missing Type

// UNSUPPORTED : C value 'is_synthesized' : for field getter : missing Type

// UNSUPPORTED : C value 'is_synthesized' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field setter : missing Type

// FontFaceClassStruct creates an uninitialised FontFaceClass.
func FontFaceClassStruct() *FontFaceClass {
	err := fontFaceClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontFaceClassNewFromNative(fontFaceClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontFaceClass)
	return structGo
}
func finalizeFontFaceClass(obj *FontFaceClass) {
	fontFaceClassStruct.Free(obj.native)
}

var fontFamilyClassStruct *gi.Struct
var fontFamilyClassStruct_Once sync.Once

func fontFamilyClassStruct_Set() error {
	var err error
	fontFamilyClassStruct_Once.Do(func() {
		fontFamilyClassStruct, err = gi.StructNew("Pango", "FontFamilyClass")
	})
	return err
}

type FontFamilyClass struct {
	native unsafe.Pointer
}

func FontFamilyClassNewFromNative(native unsafe.Pointer) *FontFamilyClass {
	return &FontFamilyClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'list_faces' : for field getter : missing Type

// UNSUPPORTED : C value 'list_faces' : for field setter : missing Type

// UNSUPPORTED : C value 'get_name' : for field getter : missing Type

// UNSUPPORTED : C value 'get_name' : for field setter : missing Type

// UNSUPPORTED : C value 'is_monospace' : for field getter : missing Type

// UNSUPPORTED : C value 'is_monospace' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field setter : missing Type

// FontFamilyClassStruct creates an uninitialised FontFamilyClass.
func FontFamilyClassStruct() *FontFamilyClass {
	err := fontFamilyClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontFamilyClassNewFromNative(fontFamilyClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontFamilyClass)
	return structGo
}
func finalizeFontFamilyClass(obj *FontFamilyClass) {
	fontFamilyClassStruct.Free(obj.native)
}

var fontMapClassStruct *gi.Struct
var fontMapClassStruct_Once sync.Once

func fontMapClassStruct_Set() error {
	var err error
	fontMapClassStruct_Once.Do(func() {
		fontMapClassStruct, err = gi.StructNew("Pango", "FontMapClass")
	})
	return err
}

type FontMapClass struct {
	native unsafe.Pointer
}

func FontMapClassNewFromNative(native unsafe.Pointer) *FontMapClass {
	return &FontMapClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'load_font' : for field getter : missing Type

// UNSUPPORTED : C value 'load_font' : for field setter : missing Type

// UNSUPPORTED : C value 'list_families' : for field getter : missing Type

// UNSUPPORTED : C value 'list_families' : for field setter : missing Type

// UNSUPPORTED : C value 'load_fontset' : for field getter : missing Type

// UNSUPPORTED : C value 'load_fontset' : for field setter : missing Type

// FieldShapeEngineType returns the C field 'shape_engine_type'.
func (recv *FontMapClass) FieldShapeEngineType() string {
	argValue := gi.StructFieldGet(fontMapClassStruct, recv.native, "shape_engine_type")
	value := argValue.String(false)
	return value
}

// SetFieldShapeEngineType sets the value of the C field 'shape_engine_type'.
func (recv *FontMapClass) SetFieldShapeEngineType(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(fontMapClassStruct, recv.native, "shape_engine_type", argValue)
}

// UNSUPPORTED : C value 'get_serial' : for field getter : missing Type

// UNSUPPORTED : C value 'get_serial' : for field setter : missing Type

// UNSUPPORTED : C value 'changed' : for field getter : missing Type

// UNSUPPORTED : C value 'changed' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field setter : missing Type

// FontMapClassStruct creates an uninitialised FontMapClass.
func FontMapClassStruct() *FontMapClass {
	err := fontMapClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontMapClassNewFromNative(fontMapClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontMapClass)
	return structGo
}
func finalizeFontMapClass(obj *FontMapClass) {
	fontMapClassStruct.Free(obj.native)
}

var fontMetricsStruct *gi.Struct
var fontMetricsStruct_Once sync.Once

func fontMetricsStruct_Set() error {
	var err error
	fontMetricsStruct_Once.Do(func() {
		fontMetricsStruct, err = gi.StructNew("Pango", "FontMetrics")
	})
	return err
}

type FontMetrics struct {
	native unsafe.Pointer
}

func FontMetricsNewFromNative(native unsafe.Pointer) *FontMetrics {
	return &FontMetrics{native: native}
}

var fontMetricsNewFunction *gi.Function
var fontMetricsNewFunction_Once sync.Once

func fontMetricsNewFunction_Set() error {
	var err error
	fontMetricsNewFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsNewFunction, err = fontMetricsStruct.InvokerNew("new")
	})
	return err
}

// FontMetricsNew is a representation of the C type pango_font_metrics_new.
func FontMetricsNew() *FontMetrics {

	var ret gi.Argument

	err := fontMetricsNewFunction_Set()
	if err == nil {
		ret = fontMetricsNewFunction.Invoke(nil, nil)
	}

	retGo := FontMetricsNewFromNative(ret.Pointer())

	return retGo
}

var fontMetricsGetApproximateCharWidthFunction *gi.Function
var fontMetricsGetApproximateCharWidthFunction_Once sync.Once

func fontMetricsGetApproximateCharWidthFunction_Set() error {
	var err error
	fontMetricsGetApproximateCharWidthFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetApproximateCharWidthFunction, err = fontMetricsStruct.InvokerNew("get_approximate_char_width")
	})
	return err
}

// GetApproximateCharWidth is a representation of the C type pango_font_metrics_get_approximate_char_width.
func (recv *FontMetrics) GetApproximateCharWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateCharWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateCharWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetApproximateDigitWidthFunction *gi.Function
var fontMetricsGetApproximateDigitWidthFunction_Once sync.Once

func fontMetricsGetApproximateDigitWidthFunction_Set() error {
	var err error
	fontMetricsGetApproximateDigitWidthFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetApproximateDigitWidthFunction, err = fontMetricsStruct.InvokerNew("get_approximate_digit_width")
	})
	return err
}

// GetApproximateDigitWidth is a representation of the C type pango_font_metrics_get_approximate_digit_width.
func (recv *FontMetrics) GetApproximateDigitWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetApproximateDigitWidthFunction_Set()
	if err == nil {
		ret = fontMetricsGetApproximateDigitWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetAscentFunction *gi.Function
var fontMetricsGetAscentFunction_Once sync.Once

func fontMetricsGetAscentFunction_Set() error {
	var err error
	fontMetricsGetAscentFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetAscentFunction, err = fontMetricsStruct.InvokerNew("get_ascent")
	})
	return err
}

// GetAscent is a representation of the C type pango_font_metrics_get_ascent.
func (recv *FontMetrics) GetAscent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetAscentFunction_Set()
	if err == nil {
		ret = fontMetricsGetAscentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetDescentFunction *gi.Function
var fontMetricsGetDescentFunction_Once sync.Once

func fontMetricsGetDescentFunction_Set() error {
	var err error
	fontMetricsGetDescentFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetDescentFunction, err = fontMetricsStruct.InvokerNew("get_descent")
	})
	return err
}

// GetDescent is a representation of the C type pango_font_metrics_get_descent.
func (recv *FontMetrics) GetDescent() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetDescentFunction_Set()
	if err == nil {
		ret = fontMetricsGetDescentFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughPositionFunction *gi.Function
var fontMetricsGetStrikethroughPositionFunction_Once sync.Once

func fontMetricsGetStrikethroughPositionFunction_Set() error {
	var err error
	fontMetricsGetStrikethroughPositionFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetStrikethroughPositionFunction, err = fontMetricsStruct.InvokerNew("get_strikethrough_position")
	})
	return err
}

// GetStrikethroughPosition is a representation of the C type pango_font_metrics_get_strikethrough_position.
func (recv *FontMetrics) GetStrikethroughPosition() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughPositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughPositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetStrikethroughThicknessFunction *gi.Function
var fontMetricsGetStrikethroughThicknessFunction_Once sync.Once

func fontMetricsGetStrikethroughThicknessFunction_Set() error {
	var err error
	fontMetricsGetStrikethroughThicknessFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetStrikethroughThicknessFunction, err = fontMetricsStruct.InvokerNew("get_strikethrough_thickness")
	})
	return err
}

// GetStrikethroughThickness is a representation of the C type pango_font_metrics_get_strikethrough_thickness.
func (recv *FontMetrics) GetStrikethroughThickness() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetStrikethroughThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetStrikethroughThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlinePositionFunction *gi.Function
var fontMetricsGetUnderlinePositionFunction_Once sync.Once

func fontMetricsGetUnderlinePositionFunction_Set() error {
	var err error
	fontMetricsGetUnderlinePositionFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetUnderlinePositionFunction, err = fontMetricsStruct.InvokerNew("get_underline_position")
	})
	return err
}

// GetUnderlinePosition is a representation of the C type pango_font_metrics_get_underline_position.
func (recv *FontMetrics) GetUnderlinePosition() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlinePositionFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlinePositionFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsGetUnderlineThicknessFunction *gi.Function
var fontMetricsGetUnderlineThicknessFunction_Once sync.Once

func fontMetricsGetUnderlineThicknessFunction_Set() error {
	var err error
	fontMetricsGetUnderlineThicknessFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsGetUnderlineThicknessFunction, err = fontMetricsStruct.InvokerNew("get_underline_thickness")
	})
	return err
}

// GetUnderlineThickness is a representation of the C type pango_font_metrics_get_underline_thickness.
func (recv *FontMetrics) GetUnderlineThickness() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsGetUnderlineThicknessFunction_Set()
	if err == nil {
		ret = fontMetricsGetUnderlineThicknessFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var fontMetricsRefFunction *gi.Function
var fontMetricsRefFunction_Once sync.Once

func fontMetricsRefFunction_Set() error {
	var err error
	fontMetricsRefFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsRefFunction, err = fontMetricsStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_font_metrics_ref.
func (recv *FontMetrics) Ref() *FontMetrics {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := fontMetricsRefFunction_Set()
	if err == nil {
		ret = fontMetricsRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := FontMetricsNewFromNative(ret.Pointer())

	return retGo
}

var fontMetricsUnrefFunction *gi.Function
var fontMetricsUnrefFunction_Once sync.Once

func fontMetricsUnrefFunction_Set() error {
	var err error
	fontMetricsUnrefFunction_Once.Do(func() {
		err = fontMetricsStruct_Set()
		if err != nil {
			return
		}
		fontMetricsUnrefFunction, err = fontMetricsStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_font_metrics_unref.
func (recv *FontMetrics) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := fontMetricsUnrefFunction_Set()
	if err == nil {
		fontMetricsUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var fontsetClassStruct *gi.Struct
var fontsetClassStruct_Once sync.Once

func fontsetClassStruct_Set() error {
	var err error
	fontsetClassStruct_Once.Do(func() {
		fontsetClassStruct, err = gi.StructNew("Pango", "FontsetClass")
	})
	return err
}

type FontsetClass struct {
	native unsafe.Pointer
}

func FontsetClassNewFromNative(native unsafe.Pointer) *FontsetClass {
	return &FontsetClass{native: native}
}

// UNSUPPORTED : C value 'parent_class' : for field getter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'parent_class' : for field setter : no Go type for 'GObject.ObjectClass'

// UNSUPPORTED : C value 'get_font' : for field getter : missing Type

// UNSUPPORTED : C value 'get_font' : for field setter : missing Type

// UNSUPPORTED : C value 'get_metrics' : for field getter : missing Type

// UNSUPPORTED : C value 'get_metrics' : for field setter : missing Type

// UNSUPPORTED : C value 'get_language' : for field getter : missing Type

// UNSUPPORTED : C value 'get_language' : for field setter : missing Type

// UNSUPPORTED : C value 'foreach' : for field getter : missing Type

// UNSUPPORTED : C value 'foreach' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved1' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field setter : missing Type

// FontsetClassStruct creates an uninitialised FontsetClass.
func FontsetClassStruct() *FontsetClass {
	err := fontsetClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontsetClassNewFromNative(fontsetClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontsetClass)
	return structGo
}
func finalizeFontsetClass(obj *FontsetClass) {
	fontsetClassStruct.Free(obj.native)
}

var fontsetSimpleClassStruct *gi.Struct
var fontsetSimpleClassStruct_Once sync.Once

func fontsetSimpleClassStruct_Set() error {
	var err error
	fontsetSimpleClassStruct_Once.Do(func() {
		fontsetSimpleClassStruct, err = gi.StructNew("Pango", "FontsetSimpleClass")
	})
	return err
}

type FontsetSimpleClass struct {
	native unsafe.Pointer
}

func FontsetSimpleClassNewFromNative(native unsafe.Pointer) *FontsetSimpleClass {
	return &FontsetSimpleClass{native: native}
}

// FontsetSimpleClassStruct creates an uninitialised FontsetSimpleClass.
func FontsetSimpleClassStruct() *FontsetSimpleClass {
	err := fontsetSimpleClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := FontsetSimpleClassNewFromNative(fontsetSimpleClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeFontsetSimpleClass)
	return structGo
}
func finalizeFontsetSimpleClass(obj *FontsetSimpleClass) {
	fontsetSimpleClassStruct.Free(obj.native)
}

var glyphGeometryStruct *gi.Struct
var glyphGeometryStruct_Once sync.Once

func glyphGeometryStruct_Set() error {
	var err error
	glyphGeometryStruct_Once.Do(func() {
		glyphGeometryStruct, err = gi.StructNew("Pango", "GlyphGeometry")
	})
	return err
}

type GlyphGeometry struct {
	native unsafe.Pointer
}

func GlyphGeometryNewFromNative(native unsafe.Pointer) *GlyphGeometry {
	return &GlyphGeometry{native: native}
}

// FieldWidth returns the C field 'width'.
func (recv *GlyphGeometry) FieldWidth() GlyphUnit {
	argValue := gi.StructFieldGet(glyphGeometryStruct, recv.native, "width")
	value := GlyphUnit(argValue.Int32())
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *GlyphGeometry) SetFieldWidth(value GlyphUnit) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(glyphGeometryStruct, recv.native, "width", argValue)
}

// FieldXOffset returns the C field 'x_offset'.
func (recv *GlyphGeometry) FieldXOffset() GlyphUnit {
	argValue := gi.StructFieldGet(glyphGeometryStruct, recv.native, "x_offset")
	value := GlyphUnit(argValue.Int32())
	return value
}

// SetFieldXOffset sets the value of the C field 'x_offset'.
func (recv *GlyphGeometry) SetFieldXOffset(value GlyphUnit) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(glyphGeometryStruct, recv.native, "x_offset", argValue)
}

// FieldYOffset returns the C field 'y_offset'.
func (recv *GlyphGeometry) FieldYOffset() GlyphUnit {
	argValue := gi.StructFieldGet(glyphGeometryStruct, recv.native, "y_offset")
	value := GlyphUnit(argValue.Int32())
	return value
}

// SetFieldYOffset sets the value of the C field 'y_offset'.
func (recv *GlyphGeometry) SetFieldYOffset(value GlyphUnit) {
	var argValue gi.Argument
	argValue.SetInt32(int32(value))
	gi.StructFieldSet(glyphGeometryStruct, recv.native, "y_offset", argValue)
}

// GlyphGeometryStruct creates an uninitialised GlyphGeometry.
func GlyphGeometryStruct() *GlyphGeometry {
	err := glyphGeometryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlyphGeometryNewFromNative(glyphGeometryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlyphGeometry)
	return structGo
}
func finalizeGlyphGeometry(obj *GlyphGeometry) {
	glyphGeometryStruct.Free(obj.native)
}

var glyphInfoStruct *gi.Struct
var glyphInfoStruct_Once sync.Once

func glyphInfoStruct_Set() error {
	var err error
	glyphInfoStruct_Once.Do(func() {
		glyphInfoStruct, err = gi.StructNew("Pango", "GlyphInfo")
	})
	return err
}

type GlyphInfo struct {
	native unsafe.Pointer
}

func GlyphInfoNewFromNative(native unsafe.Pointer) *GlyphInfo {
	return &GlyphInfo{native: native}
}

// FieldGlyph returns the C field 'glyph'.
func (recv *GlyphInfo) FieldGlyph() Glyph {
	argValue := gi.StructFieldGet(glyphInfoStruct, recv.native, "glyph")
	value := Glyph(argValue.Uint32())
	return value
}

// SetFieldGlyph sets the value of the C field 'glyph'.
func (recv *GlyphInfo) SetFieldGlyph(value Glyph) {
	var argValue gi.Argument
	argValue.SetUint32(uint32(value))
	gi.StructFieldSet(glyphInfoStruct, recv.native, "glyph", argValue)
}

// FieldGeometry returns the C field 'geometry'.
func (recv *GlyphInfo) FieldGeometry() *GlyphGeometry {
	argValue := gi.StructFieldGet(glyphInfoStruct, recv.native, "geometry")
	value := GlyphGeometryNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGeometry sets the value of the C field 'geometry'.
func (recv *GlyphInfo) SetFieldGeometry(value *GlyphGeometry) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(glyphInfoStruct, recv.native, "geometry", argValue)
}

// FieldAttr returns the C field 'attr'.
func (recv *GlyphInfo) FieldAttr() *GlyphVisAttr {
	argValue := gi.StructFieldGet(glyphInfoStruct, recv.native, "attr")
	value := GlyphVisAttrNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAttr sets the value of the C field 'attr'.
func (recv *GlyphInfo) SetFieldAttr(value *GlyphVisAttr) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(glyphInfoStruct, recv.native, "attr", argValue)
}

// GlyphInfoStruct creates an uninitialised GlyphInfo.
func GlyphInfoStruct() *GlyphInfo {
	err := glyphInfoStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlyphInfoNewFromNative(glyphInfoStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlyphInfo)
	return structGo
}
func finalizeGlyphInfo(obj *GlyphInfo) {
	glyphInfoStruct.Free(obj.native)
}

var glyphItemStruct *gi.Struct
var glyphItemStruct_Once sync.Once

func glyphItemStruct_Set() error {
	var err error
	glyphItemStruct_Once.Do(func() {
		glyphItemStruct, err = gi.StructNew("Pango", "GlyphItem")
	})
	return err
}

type GlyphItem struct {
	native unsafe.Pointer
}

func GlyphItemNewFromNative(native unsafe.Pointer) *GlyphItem {
	return &GlyphItem{native: native}
}

// FieldItem returns the C field 'item'.
func (recv *GlyphItem) FieldItem() *Item {
	argValue := gi.StructFieldGet(glyphItemStruct, recv.native, "item")
	value := ItemNewFromNative(argValue.Pointer())
	return value
}

// SetFieldItem sets the value of the C field 'item'.
func (recv *GlyphItem) SetFieldItem(value *Item) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(glyphItemStruct, recv.native, "item", argValue)
}

// FieldGlyphs returns the C field 'glyphs'.
func (recv *GlyphItem) FieldGlyphs() *GlyphString {
	argValue := gi.StructFieldGet(glyphItemStruct, recv.native, "glyphs")
	value := GlyphStringNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGlyphs sets the value of the C field 'glyphs'.
func (recv *GlyphItem) SetFieldGlyphs(value *GlyphString) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(glyphItemStruct, recv.native, "glyphs", argValue)
}

// UNSUPPORTED : C value 'pango_glyph_item_apply_attrs' : return type 'GLib.SList' not supported

var glyphItemCopyFunction *gi.Function
var glyphItemCopyFunction_Once sync.Once

func glyphItemCopyFunction_Set() error {
	var err error
	glyphItemCopyFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemCopyFunction, err = glyphItemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_item_copy.
func (recv *GlyphItem) Copy() *GlyphItem {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemCopyFunction_Set()
	if err == nil {
		ret = glyphItemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := GlyphItemNewFromNative(ret.Pointer())

	return retGo
}

var glyphItemFreeFunction *gi.Function
var glyphItemFreeFunction_Once sync.Once

func glyphItemFreeFunction_Set() error {
	var err error
	glyphItemFreeFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemFreeFunction, err = glyphItemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_item_free.
func (recv *GlyphItem) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemFreeFunction_Set()
	if err == nil {
		glyphItemFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_glyph_item_get_logical_widths' : parameter 'logical_widths' of type 'nil' not supported

// UNSUPPORTED : C value 'pango_glyph_item_letter_space' : parameter 'log_attrs' of type 'nil' not supported

var glyphItemSplitFunction *gi.Function
var glyphItemSplitFunction_Once sync.Once

func glyphItemSplitFunction_Set() error {
	var err error
	glyphItemSplitFunction_Once.Do(func() {
		err = glyphItemStruct_Set()
		if err != nil {
			return
		}
		glyphItemSplitFunction, err = glyphItemStruct.InvokerNew("split")
	})
	return err
}

// Split is a representation of the C type pango_glyph_item_split.
func (recv *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(splitIndex)

	var ret gi.Argument

	err := glyphItemSplitFunction_Set()
	if err == nil {
		ret = glyphItemSplitFunction.Invoke(inArgs[:], nil)
	}

	retGo := GlyphItemNewFromNative(ret.Pointer())

	return retGo
}

// GlyphItemStruct creates an uninitialised GlyphItem.
func GlyphItemStruct() *GlyphItem {
	err := glyphItemStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlyphItemNewFromNative(glyphItemStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlyphItem)
	return structGo
}
func finalizeGlyphItem(obj *GlyphItem) {
	glyphItemStruct.Free(obj.native)
}

var glyphItemIterStruct *gi.Struct
var glyphItemIterStruct_Once sync.Once

func glyphItemIterStruct_Set() error {
	var err error
	glyphItemIterStruct_Once.Do(func() {
		glyphItemIterStruct, err = gi.StructNew("Pango", "GlyphItemIter")
	})
	return err
}

type GlyphItemIter struct {
	native unsafe.Pointer
}

func GlyphItemIterNewFromNative(native unsafe.Pointer) *GlyphItemIter {
	return &GlyphItemIter{native: native}
}

// FieldGlyphItem returns the C field 'glyph_item'.
func (recv *GlyphItemIter) FieldGlyphItem() *GlyphItem {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "glyph_item")
	value := GlyphItemNewFromNative(argValue.Pointer())
	return value
}

// SetFieldGlyphItem sets the value of the C field 'glyph_item'.
func (recv *GlyphItemIter) SetFieldGlyphItem(value *GlyphItem) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "glyph_item", argValue)
}

// FieldText returns the C field 'text'.
func (recv *GlyphItemIter) FieldText() string {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "text")
	value := argValue.String(false)
	return value
}

// SetFieldText sets the value of the C field 'text'.
func (recv *GlyphItemIter) SetFieldText(value string) {
	var argValue gi.Argument
	argValue.SetString(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "text", argValue)
}

// FieldStartGlyph returns the C field 'start_glyph'.
func (recv *GlyphItemIter) FieldStartGlyph() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "start_glyph")
	value := argValue.Int32()
	return value
}

// SetFieldStartGlyph sets the value of the C field 'start_glyph'.
func (recv *GlyphItemIter) SetFieldStartGlyph(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "start_glyph", argValue)
}

// FieldStartIndex returns the C field 'start_index'.
func (recv *GlyphItemIter) FieldStartIndex() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "start_index")
	value := argValue.Int32()
	return value
}

// SetFieldStartIndex sets the value of the C field 'start_index'.
func (recv *GlyphItemIter) SetFieldStartIndex(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "start_index", argValue)
}

// FieldStartChar returns the C field 'start_char'.
func (recv *GlyphItemIter) FieldStartChar() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "start_char")
	value := argValue.Int32()
	return value
}

// SetFieldStartChar sets the value of the C field 'start_char'.
func (recv *GlyphItemIter) SetFieldStartChar(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "start_char", argValue)
}

// FieldEndGlyph returns the C field 'end_glyph'.
func (recv *GlyphItemIter) FieldEndGlyph() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "end_glyph")
	value := argValue.Int32()
	return value
}

// SetFieldEndGlyph sets the value of the C field 'end_glyph'.
func (recv *GlyphItemIter) SetFieldEndGlyph(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "end_glyph", argValue)
}

// FieldEndIndex returns the C field 'end_index'.
func (recv *GlyphItemIter) FieldEndIndex() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "end_index")
	value := argValue.Int32()
	return value
}

// SetFieldEndIndex sets the value of the C field 'end_index'.
func (recv *GlyphItemIter) SetFieldEndIndex(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "end_index", argValue)
}

// FieldEndChar returns the C field 'end_char'.
func (recv *GlyphItemIter) FieldEndChar() int32 {
	argValue := gi.StructFieldGet(glyphItemIterStruct, recv.native, "end_char")
	value := argValue.Int32()
	return value
}

// SetFieldEndChar sets the value of the C field 'end_char'.
func (recv *GlyphItemIter) SetFieldEndChar(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphItemIterStruct, recv.native, "end_char", argValue)
}

var glyphItemIterCopyFunction *gi.Function
var glyphItemIterCopyFunction_Once sync.Once

func glyphItemIterCopyFunction_Set() error {
	var err error
	glyphItemIterCopyFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterCopyFunction, err = glyphItemIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_item_iter_copy.
func (recv *GlyphItemIter) Copy() *GlyphItemIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterCopyFunction_Set()
	if err == nil {
		ret = glyphItemIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := GlyphItemIterNewFromNative(ret.Pointer())

	return retGo
}

var glyphItemIterFreeFunction *gi.Function
var glyphItemIterFreeFunction_Once sync.Once

func glyphItemIterFreeFunction_Set() error {
	var err error
	glyphItemIterFreeFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterFreeFunction, err = glyphItemIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_item_iter_free.
func (recv *GlyphItemIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphItemIterFreeFunction_Set()
	if err == nil {
		glyphItemIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var glyphItemIterInitEndFunction *gi.Function
var glyphItemIterInitEndFunction_Once sync.Once

func glyphItemIterInitEndFunction_Set() error {
	var err error
	glyphItemIterInitEndFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterInitEndFunction, err = glyphItemIterStruct.InvokerNew("init_end")
	})
	return err
}

// InitEnd is a representation of the C type pango_glyph_item_iter_init_end.
func (recv *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(glyphItem.native)
	inArgs[2].SetString(text)

	var ret gi.Argument

	err := glyphItemIterInitEndFunction_Set()
	if err == nil {
		ret = glyphItemIterInitEndFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var glyphItemIterInitStartFunction *gi.Function
var glyphItemIterInitStartFunction_Once sync.Once

func glyphItemIterInitStartFunction_Set() error {
	var err error
	glyphItemIterInitStartFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterInitStartFunction, err = glyphItemIterStruct.InvokerNew("init_start")
	})
	return err
}

// InitStart is a representation of the C type pango_glyph_item_iter_init_start.
func (recv *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(glyphItem.native)
	inArgs[2].SetString(text)

	var ret gi.Argument

	err := glyphItemIterInitStartFunction_Set()
	if err == nil {
		ret = glyphItemIterInitStartFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var glyphItemIterNextClusterFunction *gi.Function
var glyphItemIterNextClusterFunction_Once sync.Once

func glyphItemIterNextClusterFunction_Set() error {
	var err error
	glyphItemIterNextClusterFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterNextClusterFunction, err = glyphItemIterStruct.InvokerNew("next_cluster")
	})
	return err
}

// NextCluster is a representation of the C type pango_glyph_item_iter_next_cluster.
func (recv *GlyphItemIter) NextCluster() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterNextClusterFunction_Set()
	if err == nil {
		ret = glyphItemIterNextClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var glyphItemIterPrevClusterFunction *gi.Function
var glyphItemIterPrevClusterFunction_Once sync.Once

func glyphItemIterPrevClusterFunction_Set() error {
	var err error
	glyphItemIterPrevClusterFunction_Once.Do(func() {
		err = glyphItemIterStruct_Set()
		if err != nil {
			return
		}
		glyphItemIterPrevClusterFunction, err = glyphItemIterStruct.InvokerNew("prev_cluster")
	})
	return err
}

// PrevCluster is a representation of the C type pango_glyph_item_iter_prev_cluster.
func (recv *GlyphItemIter) PrevCluster() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphItemIterPrevClusterFunction_Set()
	if err == nil {
		ret = glyphItemIterPrevClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// GlyphItemIterStruct creates an uninitialised GlyphItemIter.
func GlyphItemIterStruct() *GlyphItemIter {
	err := glyphItemIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlyphItemIterNewFromNative(glyphItemIterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlyphItemIter)
	return structGo
}
func finalizeGlyphItemIter(obj *GlyphItemIter) {
	glyphItemIterStruct.Free(obj.native)
}

var glyphStringStruct *gi.Struct
var glyphStringStruct_Once sync.Once

func glyphStringStruct_Set() error {
	var err error
	glyphStringStruct_Once.Do(func() {
		glyphStringStruct, err = gi.StructNew("Pango", "GlyphString")
	})
	return err
}

type GlyphString struct {
	native unsafe.Pointer
}

func GlyphStringNewFromNative(native unsafe.Pointer) *GlyphString {
	return &GlyphString{native: native}
}

// FieldNumGlyphs returns the C field 'num_glyphs'.
func (recv *GlyphString) FieldNumGlyphs() int32 {
	argValue := gi.StructFieldGet(glyphStringStruct, recv.native, "num_glyphs")
	value := argValue.Int32()
	return value
}

// SetFieldNumGlyphs sets the value of the C field 'num_glyphs'.
func (recv *GlyphString) SetFieldNumGlyphs(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphStringStruct, recv.native, "num_glyphs", argValue)
}

// UNSUPPORTED : C value 'glyphs' : for field getter : missing Type

// UNSUPPORTED : C value 'glyphs' : for field setter : missing Type

// FieldLogClusters returns the C field 'log_clusters'.
func (recv *GlyphString) FieldLogClusters() int32 {
	argValue := gi.StructFieldGet(glyphStringStruct, recv.native, "log_clusters")
	value := argValue.Int32()
	return value
}

// SetFieldLogClusters sets the value of the C field 'log_clusters'.
func (recv *GlyphString) SetFieldLogClusters(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(glyphStringStruct, recv.native, "log_clusters", argValue)
}

var glyphStringNewFunction *gi.Function
var glyphStringNewFunction_Once sync.Once

func glyphStringNewFunction_Set() error {
	var err error
	glyphStringNewFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringNewFunction, err = glyphStringStruct.InvokerNew("new")
	})
	return err
}

// GlyphStringNew is a representation of the C type pango_glyph_string_new.
func GlyphStringNew() *GlyphString {

	var ret gi.Argument

	err := glyphStringNewFunction_Set()
	if err == nil {
		ret = glyphStringNewFunction.Invoke(nil, nil)
	}

	retGo := GlyphStringNewFromNative(ret.Pointer())

	return retGo
}

var glyphStringCopyFunction *gi.Function
var glyphStringCopyFunction_Once sync.Once

func glyphStringCopyFunction_Set() error {
	var err error
	glyphStringCopyFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringCopyFunction, err = glyphStringStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_glyph_string_copy.
func (recv *GlyphString) Copy() *GlyphString {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringCopyFunction_Set()
	if err == nil {
		ret = glyphStringCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := GlyphStringNewFromNative(ret.Pointer())

	return retGo
}

var glyphStringExtentsFunction *gi.Function
var glyphStringExtentsFunction_Once sync.Once

func glyphStringExtentsFunction_Set() error {
	var err error
	glyphStringExtentsFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringExtentsFunction, err = glyphStringStruct.InvokerNew("extents")
	})
	return err
}

// Extents is a representation of the C type pango_glyph_string_extents.
func (recv *GlyphString) Extents(font *Font) (*Rectangle, *Rectangle) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(font.native)

	var outArgs [2]gi.Argument

	err := glyphStringExtentsFunction_Set()
	if err == nil {
		glyphStringExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var glyphStringExtentsRangeFunction *gi.Function
var glyphStringExtentsRangeFunction_Once sync.Once

func glyphStringExtentsRangeFunction_Set() error {
	var err error
	glyphStringExtentsRangeFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringExtentsRangeFunction, err = glyphStringStruct.InvokerNew("extents_range")
	})
	return err
}

// ExtentsRange is a representation of the C type pango_glyph_string_extents_range.
func (recv *GlyphString) ExtentsRange(start int32, end int32, font *Font) (*Rectangle, *Rectangle) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(start)
	inArgs[2].SetInt32(end)
	inArgs[3].SetPointer(font.native)

	var outArgs [2]gi.Argument

	err := glyphStringExtentsRangeFunction_Set()
	if err == nil {
		glyphStringExtentsRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var glyphStringFreeFunction *gi.Function
var glyphStringFreeFunction_Once sync.Once

func glyphStringFreeFunction_Set() error {
	var err error
	glyphStringFreeFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringFreeFunction, err = glyphStringStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_glyph_string_free.
func (recv *GlyphString) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := glyphStringFreeFunction_Set()
	if err == nil {
		glyphStringFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

// UNSUPPORTED : C value 'pango_glyph_string_get_logical_widths' : parameter 'logical_widths' of type 'nil' not supported

var glyphStringGetWidthFunction *gi.Function
var glyphStringGetWidthFunction_Once sync.Once

func glyphStringGetWidthFunction_Set() error {
	var err error
	glyphStringGetWidthFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringGetWidthFunction, err = glyphStringStruct.InvokerNew("get_width")
	})
	return err
}

// GetWidth is a representation of the C type pango_glyph_string_get_width.
func (recv *GlyphString) GetWidth() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := glyphStringGetWidthFunction_Set()
	if err == nil {
		ret = glyphStringGetWidthFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var glyphStringIndexToXFunction *gi.Function
var glyphStringIndexToXFunction_Once sync.Once

func glyphStringIndexToXFunction_Set() error {
	var err error
	glyphStringIndexToXFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringIndexToXFunction, err = glyphStringStruct.InvokerNew("index_to_x")
	})
	return err
}

// IndexToX is a representation of the C type pango_glyph_string_index_to_x.
func (recv *GlyphString) IndexToX(text string, length int32, analysis *Analysis, index int32, trailing bool) int32 {
	var inArgs [6]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)
	inArgs[3].SetPointer(analysis.native)
	inArgs[4].SetInt32(index)
	inArgs[5].SetBoolean(trailing)

	var outArgs [1]gi.Argument

	err := glyphStringIndexToXFunction_Set()
	if err == nil {
		glyphStringIndexToXFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var glyphStringSetSizeFunction *gi.Function
var glyphStringSetSizeFunction_Once sync.Once

func glyphStringSetSizeFunction_Set() error {
	var err error
	glyphStringSetSizeFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringSetSizeFunction, err = glyphStringStruct.InvokerNew("set_size")
	})
	return err
}

// SetSize is a representation of the C type pango_glyph_string_set_size.
func (recv *GlyphString) SetSize(newLen int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newLen)

	err := glyphStringSetSizeFunction_Set()
	if err == nil {
		glyphStringSetSizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var glyphStringXToIndexFunction *gi.Function
var glyphStringXToIndexFunction_Once sync.Once

func glyphStringXToIndexFunction_Set() error {
	var err error
	glyphStringXToIndexFunction_Once.Do(func() {
		err = glyphStringStruct_Set()
		if err != nil {
			return
		}
		glyphStringXToIndexFunction, err = glyphStringStruct.InvokerNew("x_to_index")
	})
	return err
}

// XToIndex is a representation of the C type pango_glyph_string_x_to_index.
func (recv *GlyphString) XToIndex(text string, length int32, analysis *Analysis, xPos int32) (int32, int32) {
	var inArgs [5]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(text)
	inArgs[2].SetInt32(length)
	inArgs[3].SetPointer(analysis.native)
	inArgs[4].SetInt32(xPos)

	var outArgs [2]gi.Argument

	err := glyphStringXToIndexFunction_Set()
	if err == nil {
		glyphStringXToIndexFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

var glyphVisAttrStruct *gi.Struct
var glyphVisAttrStruct_Once sync.Once

func glyphVisAttrStruct_Set() error {
	var err error
	glyphVisAttrStruct_Once.Do(func() {
		glyphVisAttrStruct, err = gi.StructNew("Pango", "GlyphVisAttr")
	})
	return err
}

type GlyphVisAttr struct {
	native unsafe.Pointer
}

func GlyphVisAttrNewFromNative(native unsafe.Pointer) *GlyphVisAttr {
	return &GlyphVisAttr{native: native}
}

// FieldIsClusterStart returns the C field 'is_cluster_start'.
func (recv *GlyphVisAttr) FieldIsClusterStart() uint32 {
	argValue := gi.StructFieldGet(glyphVisAttrStruct, recv.native, "is_cluster_start")
	value := argValue.Uint32()
	return value
}

// SetFieldIsClusterStart sets the value of the C field 'is_cluster_start'.
func (recv *GlyphVisAttr) SetFieldIsClusterStart(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(glyphVisAttrStruct, recv.native, "is_cluster_start", argValue)
}

// GlyphVisAttrStruct creates an uninitialised GlyphVisAttr.
func GlyphVisAttrStruct() *GlyphVisAttr {
	err := glyphVisAttrStruct_Set()
	if err != nil {
		return nil
	}

	structGo := GlyphVisAttrNewFromNative(glyphVisAttrStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeGlyphVisAttr)
	return structGo
}
func finalizeGlyphVisAttr(obj *GlyphVisAttr) {
	glyphVisAttrStruct.Free(obj.native)
}

var includedModuleStruct *gi.Struct
var includedModuleStruct_Once sync.Once

func includedModuleStruct_Set() error {
	var err error
	includedModuleStruct_Once.Do(func() {
		includedModuleStruct, err = gi.StructNew("Pango", "IncludedModule")
	})
	return err
}

type IncludedModule struct {
	native unsafe.Pointer
}

func IncludedModuleNewFromNative(native unsafe.Pointer) *IncludedModule {
	return &IncludedModule{native: native}
}

// UNSUPPORTED : C value 'list' : for field getter : missing Type

// UNSUPPORTED : C value 'list' : for field setter : missing Type

// UNSUPPORTED : C value 'init' : for field getter : missing Type

// UNSUPPORTED : C value 'init' : for field setter : missing Type

// UNSUPPORTED : C value 'exit' : for field getter : missing Type

// UNSUPPORTED : C value 'exit' : for field setter : missing Type

// UNSUPPORTED : C value 'create' : for field getter : missing Type

// UNSUPPORTED : C value 'create' : for field setter : missing Type

// IncludedModuleStruct creates an uninitialised IncludedModule.
func IncludedModuleStruct() *IncludedModule {
	err := includedModuleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := IncludedModuleNewFromNative(includedModuleStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeIncludedModule)
	return structGo
}
func finalizeIncludedModule(obj *IncludedModule) {
	includedModuleStruct.Free(obj.native)
}

var itemStruct *gi.Struct
var itemStruct_Once sync.Once

func itemStruct_Set() error {
	var err error
	itemStruct_Once.Do(func() {
		itemStruct, err = gi.StructNew("Pango", "Item")
	})
	return err
}

type Item struct {
	native unsafe.Pointer
}

func ItemNewFromNative(native unsafe.Pointer) *Item {
	return &Item{native: native}
}

// FieldOffset returns the C field 'offset'.
func (recv *Item) FieldOffset() int32 {
	argValue := gi.StructFieldGet(itemStruct, recv.native, "offset")
	value := argValue.Int32()
	return value
}

// SetFieldOffset sets the value of the C field 'offset'.
func (recv *Item) SetFieldOffset(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(itemStruct, recv.native, "offset", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *Item) FieldLength() int32 {
	argValue := gi.StructFieldGet(itemStruct, recv.native, "length")
	value := argValue.Int32()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *Item) SetFieldLength(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(itemStruct, recv.native, "length", argValue)
}

// FieldNumChars returns the C field 'num_chars'.
func (recv *Item) FieldNumChars() int32 {
	argValue := gi.StructFieldGet(itemStruct, recv.native, "num_chars")
	value := argValue.Int32()
	return value
}

// SetFieldNumChars sets the value of the C field 'num_chars'.
func (recv *Item) SetFieldNumChars(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(itemStruct, recv.native, "num_chars", argValue)
}

// FieldAnalysis returns the C field 'analysis'.
func (recv *Item) FieldAnalysis() *Analysis {
	argValue := gi.StructFieldGet(itemStruct, recv.native, "analysis")
	value := AnalysisNewFromNative(argValue.Pointer())
	return value
}

// SetFieldAnalysis sets the value of the C field 'analysis'.
func (recv *Item) SetFieldAnalysis(value *Analysis) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(itemStruct, recv.native, "analysis", argValue)
}

var itemNewFunction *gi.Function
var itemNewFunction_Once sync.Once

func itemNewFunction_Set() error {
	var err error
	itemNewFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemNewFunction, err = itemStruct.InvokerNew("new")
	})
	return err
}

// ItemNew is a representation of the C type pango_item_new.
func ItemNew() *Item {

	var ret gi.Argument

	err := itemNewFunction_Set()
	if err == nil {
		ret = itemNewFunction.Invoke(nil, nil)
	}

	retGo := ItemNewFromNative(ret.Pointer())

	return retGo
}

var itemCopyFunction *gi.Function
var itemCopyFunction_Once sync.Once

func itemCopyFunction_Set() error {
	var err error
	itemCopyFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemCopyFunction, err = itemStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_item_copy.
func (recv *Item) Copy() *Item {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := itemCopyFunction_Set()
	if err == nil {
		ret = itemCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := ItemNewFromNative(ret.Pointer())

	return retGo
}

var itemFreeFunction *gi.Function
var itemFreeFunction_Once sync.Once

func itemFreeFunction_Set() error {
	var err error
	itemFreeFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemFreeFunction, err = itemStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_item_free.
func (recv *Item) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := itemFreeFunction_Set()
	if err == nil {
		itemFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var itemSplitFunction *gi.Function
var itemSplitFunction_Once sync.Once

func itemSplitFunction_Set() error {
	var err error
	itemSplitFunction_Once.Do(func() {
		err = itemStruct_Set()
		if err != nil {
			return
		}
		itemSplitFunction, err = itemStruct.InvokerNew("split")
	})
	return err
}

// Split is a representation of the C type pango_item_split.
func (recv *Item) Split(splitIndex int32, splitOffset int32) *Item {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(splitIndex)
	inArgs[2].SetInt32(splitOffset)

	var ret gi.Argument

	err := itemSplitFunction_Set()
	if err == nil {
		ret = itemSplitFunction.Invoke(inArgs[:], nil)
	}

	retGo := ItemNewFromNative(ret.Pointer())

	return retGo
}

var languageStruct *gi.Struct
var languageStruct_Once sync.Once

func languageStruct_Set() error {
	var err error
	languageStruct_Once.Do(func() {
		languageStruct, err = gi.StructNew("Pango", "Language")
	})
	return err
}

type Language struct {
	native unsafe.Pointer
}

func LanguageNewFromNative(native unsafe.Pointer) *Language {
	return &Language{native: native}
}

var languageGetSampleStringFunction *gi.Function
var languageGetSampleStringFunction_Once sync.Once

func languageGetSampleStringFunction_Set() error {
	var err error
	languageGetSampleStringFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetSampleStringFunction, err = languageStruct.InvokerNew("get_sample_string")
	})
	return err
}

// GetSampleString is a representation of the C type pango_language_get_sample_string.
func (recv *Language) GetSampleString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageGetSampleStringFunction_Set()
	if err == nil {
		ret = languageGetSampleStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

var languageGetScriptsFunction *gi.Function
var languageGetScriptsFunction_Once sync.Once

func languageGetScriptsFunction_Set() error {
	var err error
	languageGetScriptsFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageGetScriptsFunction, err = languageStruct.InvokerNew("get_scripts")
	})
	return err
}

// GetScripts is a representation of the C type pango_language_get_scripts.
func (recv *Language) GetScripts() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := languageGetScriptsFunction_Set()
	if err == nil {
		languageGetScriptsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var languageIncludesScriptFunction *gi.Function
var languageIncludesScriptFunction_Once sync.Once

func languageIncludesScriptFunction_Set() error {
	var err error
	languageIncludesScriptFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageIncludesScriptFunction, err = languageStruct.InvokerNew("includes_script")
	})
	return err
}

// IncludesScript is a representation of the C type pango_language_includes_script.
func (recv *Language) IncludesScript(script Script) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(script))

	var ret gi.Argument

	err := languageIncludesScriptFunction_Set()
	if err == nil {
		ret = languageIncludesScriptFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var languageMatchesFunction *gi.Function
var languageMatchesFunction_Once sync.Once

func languageMatchesFunction_Set() error {
	var err error
	languageMatchesFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageMatchesFunction, err = languageStruct.InvokerNew("matches")
	})
	return err
}

// Matches is a representation of the C type pango_language_matches.
func (recv *Language) Matches(rangeList string) bool {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetString(rangeList)

	var ret gi.Argument

	err := languageMatchesFunction_Set()
	if err == nil {
		ret = languageMatchesFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var languageToStringFunction *gi.Function
var languageToStringFunction_Once sync.Once

func languageToStringFunction_Set() error {
	var err error
	languageToStringFunction_Once.Do(func() {
		err = languageStruct_Set()
		if err != nil {
			return
		}
		languageToStringFunction, err = languageStruct.InvokerNew("to_string")
	})
	return err
}

// ToString is a representation of the C type pango_language_to_string.
func (recv *Language) ToString() string {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := languageToStringFunction_Set()
	if err == nil {
		ret = languageToStringFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.String(false)

	return retGo
}

// LanguageStruct creates an uninitialised Language.
func LanguageStruct() *Language {
	err := languageStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LanguageNewFromNative(languageStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLanguage)
	return structGo
}
func finalizeLanguage(obj *Language) {
	languageStruct.Free(obj.native)
}

var layoutClassStruct *gi.Struct
var layoutClassStruct_Once sync.Once

func layoutClassStruct_Set() error {
	var err error
	layoutClassStruct_Once.Do(func() {
		layoutClassStruct, err = gi.StructNew("Pango", "LayoutClass")
	})
	return err
}

type LayoutClass struct {
	native unsafe.Pointer
}

func LayoutClassNewFromNative(native unsafe.Pointer) *LayoutClass {
	return &LayoutClass{native: native}
}

// LayoutClassStruct creates an uninitialised LayoutClass.
func LayoutClassStruct() *LayoutClass {
	err := layoutClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LayoutClassNewFromNative(layoutClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLayoutClass)
	return structGo
}
func finalizeLayoutClass(obj *LayoutClass) {
	layoutClassStruct.Free(obj.native)
}

var layoutIterStruct *gi.Struct
var layoutIterStruct_Once sync.Once

func layoutIterStruct_Set() error {
	var err error
	layoutIterStruct_Once.Do(func() {
		layoutIterStruct, err = gi.StructNew("Pango", "LayoutIter")
	})
	return err
}

type LayoutIter struct {
	native unsafe.Pointer
}

func LayoutIterNewFromNative(native unsafe.Pointer) *LayoutIter {
	return &LayoutIter{native: native}
}

var layoutIterAtLastLineFunction *gi.Function
var layoutIterAtLastLineFunction_Once sync.Once

func layoutIterAtLastLineFunction_Set() error {
	var err error
	layoutIterAtLastLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterAtLastLineFunction, err = layoutIterStruct.InvokerNew("at_last_line")
	})
	return err
}

// AtLastLine is a representation of the C type pango_layout_iter_at_last_line.
func (recv *LayoutIter) AtLastLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterAtLastLineFunction_Set()
	if err == nil {
		ret = layoutIterAtLastLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutIterCopyFunction *gi.Function
var layoutIterCopyFunction_Once sync.Once

func layoutIterCopyFunction_Set() error {
	var err error
	layoutIterCopyFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterCopyFunction, err = layoutIterStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_layout_iter_copy.
func (recv *LayoutIter) Copy() *LayoutIter {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterCopyFunction_Set()
	if err == nil {
		ret = layoutIterCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutIterNewFromNative(ret.Pointer())

	return retGo
}

var layoutIterFreeFunction *gi.Function
var layoutIterFreeFunction_Once sync.Once

func layoutIterFreeFunction_Set() error {
	var err error
	layoutIterFreeFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterFreeFunction, err = layoutIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_layout_iter_free.
func (recv *LayoutIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutIterFreeFunction_Set()
	if err == nil {
		layoutIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutIterGetBaselineFunction *gi.Function
var layoutIterGetBaselineFunction_Once sync.Once

func layoutIterGetBaselineFunction_Set() error {
	var err error
	layoutIterGetBaselineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetBaselineFunction, err = layoutIterStruct.InvokerNew("get_baseline")
	})
	return err
}

// GetBaseline is a representation of the C type pango_layout_iter_get_baseline.
func (recv *LayoutIter) GetBaseline() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetBaselineFunction_Set()
	if err == nil {
		ret = layoutIterGetBaselineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutIterGetCharExtentsFunction *gi.Function
var layoutIterGetCharExtentsFunction_Once sync.Once

func layoutIterGetCharExtentsFunction_Set() error {
	var err error
	layoutIterGetCharExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetCharExtentsFunction, err = layoutIterStruct.InvokerNew("get_char_extents")
	})
	return err
}

// GetCharExtents is a representation of the C type pango_layout_iter_get_char_extents.
func (recv *LayoutIter) GetCharExtents() *Rectangle {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [1]gi.Argument

	err := layoutIterGetCharExtentsFunction_Set()
	if err == nil {
		layoutIterGetCharExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

var layoutIterGetClusterExtentsFunction *gi.Function
var layoutIterGetClusterExtentsFunction_Once sync.Once

func layoutIterGetClusterExtentsFunction_Set() error {
	var err error
	layoutIterGetClusterExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetClusterExtentsFunction, err = layoutIterStruct.InvokerNew("get_cluster_extents")
	})
	return err
}

// GetClusterExtents is a representation of the C type pango_layout_iter_get_cluster_extents.
func (recv *LayoutIter) GetClusterExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetClusterExtentsFunction_Set()
	if err == nil {
		layoutIterGetClusterExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutIterGetIndexFunction *gi.Function
var layoutIterGetIndexFunction_Once sync.Once

func layoutIterGetIndexFunction_Set() error {
	var err error
	layoutIterGetIndexFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetIndexFunction, err = layoutIterStruct.InvokerNew("get_index")
	})
	return err
}

// GetIndex is a representation of the C type pango_layout_iter_get_index.
func (recv *LayoutIter) GetIndex() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetIndexFunction_Set()
	if err == nil {
		ret = layoutIterGetIndexFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var layoutIterGetLayoutFunction *gi.Function
var layoutIterGetLayoutFunction_Once sync.Once

func layoutIterGetLayoutFunction_Set() error {
	var err error
	layoutIterGetLayoutFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLayoutFunction, err = layoutIterStruct.InvokerNew("get_layout")
	})
	return err
}

// GetLayout is a representation of the C type pango_layout_iter_get_layout.
func (recv *LayoutIter) GetLayout() *Layout {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLayoutFunction_Set()
	if err == nil {
		ret = layoutIterGetLayoutFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutNewFromNative(ret.Pointer())

	return retGo
}

var layoutIterGetLayoutExtentsFunction *gi.Function
var layoutIterGetLayoutExtentsFunction_Once sync.Once

func layoutIterGetLayoutExtentsFunction_Set() error {
	var err error
	layoutIterGetLayoutExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLayoutExtentsFunction, err = layoutIterStruct.InvokerNew("get_layout_extents")
	})
	return err
}

// GetLayoutExtents is a representation of the C type pango_layout_iter_get_layout_extents.
func (recv *LayoutIter) GetLayoutExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLayoutExtentsFunction_Set()
	if err == nil {
		layoutIterGetLayoutExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutIterGetLineFunction *gi.Function
var layoutIterGetLineFunction_Once sync.Once

func layoutIterGetLineFunction_Set() error {
	var err error
	layoutIterGetLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineFunction, err = layoutIterStruct.InvokerNew("get_line")
	})
	return err
}

// GetLine is a representation of the C type pango_layout_iter_get_line.
func (recv *LayoutIter) GetLine() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineFunction_Set()
	if err == nil {
		ret = layoutIterGetLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

var layoutIterGetLineExtentsFunction *gi.Function
var layoutIterGetLineExtentsFunction_Once sync.Once

func layoutIterGetLineExtentsFunction_Set() error {
	var err error
	layoutIterGetLineExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineExtentsFunction, err = layoutIterStruct.InvokerNew("get_line_extents")
	})
	return err
}

// GetLineExtents is a representation of the C type pango_layout_iter_get_line_extents.
func (recv *LayoutIter) GetLineExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLineExtentsFunction_Set()
	if err == nil {
		layoutIterGetLineExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutIterGetLineReadonlyFunction *gi.Function
var layoutIterGetLineReadonlyFunction_Once sync.Once

func layoutIterGetLineReadonlyFunction_Set() error {
	var err error
	layoutIterGetLineReadonlyFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineReadonlyFunction, err = layoutIterStruct.InvokerNew("get_line_readonly")
	})
	return err
}

// GetLineReadonly is a representation of the C type pango_layout_iter_get_line_readonly.
func (recv *LayoutIter) GetLineReadonly() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterGetLineReadonlyFunction_Set()
	if err == nil {
		ret = layoutIterGetLineReadonlyFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

var layoutIterGetLineYrangeFunction *gi.Function
var layoutIterGetLineYrangeFunction_Once sync.Once

func layoutIterGetLineYrangeFunction_Set() error {
	var err error
	layoutIterGetLineYrangeFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetLineYrangeFunction, err = layoutIterStruct.InvokerNew("get_line_yrange")
	})
	return err
}

// GetLineYrange is a representation of the C type pango_layout_iter_get_line_yrange.
func (recv *LayoutIter) GetLineYrange() (int32, int32) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetLineYrangeFunction_Set()
	if err == nil {
		layoutIterGetLineYrangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run' : return type 'LayoutRun' not supported

var layoutIterGetRunExtentsFunction *gi.Function
var layoutIterGetRunExtentsFunction_Once sync.Once

func layoutIterGetRunExtentsFunction_Set() error {
	var err error
	layoutIterGetRunExtentsFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterGetRunExtentsFunction, err = layoutIterStruct.InvokerNew("get_run_extents")
	})
	return err
}

// GetRunExtents is a representation of the C type pango_layout_iter_get_run_extents.
func (recv *LayoutIter) GetRunExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutIterGetRunExtentsFunction_Set()
	if err == nil {
		layoutIterGetRunExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_iter_get_run_readonly' : return type 'LayoutRun' not supported

var layoutIterNextCharFunction *gi.Function
var layoutIterNextCharFunction_Once sync.Once

func layoutIterNextCharFunction_Set() error {
	var err error
	layoutIterNextCharFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextCharFunction, err = layoutIterStruct.InvokerNew("next_char")
	})
	return err
}

// NextChar is a representation of the C type pango_layout_iter_next_char.
func (recv *LayoutIter) NextChar() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextCharFunction_Set()
	if err == nil {
		ret = layoutIterNextCharFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutIterNextClusterFunction *gi.Function
var layoutIterNextClusterFunction_Once sync.Once

func layoutIterNextClusterFunction_Set() error {
	var err error
	layoutIterNextClusterFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextClusterFunction, err = layoutIterStruct.InvokerNew("next_cluster")
	})
	return err
}

// NextCluster is a representation of the C type pango_layout_iter_next_cluster.
func (recv *LayoutIter) NextCluster() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextClusterFunction_Set()
	if err == nil {
		ret = layoutIterNextClusterFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutIterNextLineFunction *gi.Function
var layoutIterNextLineFunction_Once sync.Once

func layoutIterNextLineFunction_Set() error {
	var err error
	layoutIterNextLineFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextLineFunction, err = layoutIterStruct.InvokerNew("next_line")
	})
	return err
}

// NextLine is a representation of the C type pango_layout_iter_next_line.
func (recv *LayoutIter) NextLine() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextLineFunction_Set()
	if err == nil {
		ret = layoutIterNextLineFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var layoutIterNextRunFunction *gi.Function
var layoutIterNextRunFunction_Once sync.Once

func layoutIterNextRunFunction_Set() error {
	var err error
	layoutIterNextRunFunction_Once.Do(func() {
		err = layoutIterStruct_Set()
		if err != nil {
			return
		}
		layoutIterNextRunFunction, err = layoutIterStruct.InvokerNew("next_run")
	})
	return err
}

// NextRun is a representation of the C type pango_layout_iter_next_run.
func (recv *LayoutIter) NextRun() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutIterNextRunFunction_Set()
	if err == nil {
		ret = layoutIterNextRunFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// LayoutIterStruct creates an uninitialised LayoutIter.
func LayoutIterStruct() *LayoutIter {
	err := layoutIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LayoutIterNewFromNative(layoutIterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLayoutIter)
	return structGo
}
func finalizeLayoutIter(obj *LayoutIter) {
	layoutIterStruct.Free(obj.native)
}

var layoutLineStruct *gi.Struct
var layoutLineStruct_Once sync.Once

func layoutLineStruct_Set() error {
	var err error
	layoutLineStruct_Once.Do(func() {
		layoutLineStruct, err = gi.StructNew("Pango", "LayoutLine")
	})
	return err
}

type LayoutLine struct {
	native unsafe.Pointer
}

func LayoutLineNewFromNative(native unsafe.Pointer) *LayoutLine {
	return &LayoutLine{native: native}
}

// FieldLayout returns the C field 'layout'.
func (recv *LayoutLine) FieldLayout() *Layout {
	argValue := gi.StructFieldGet(layoutLineStruct, recv.native, "layout")
	value := LayoutNewFromNative(argValue.Pointer())
	return value
}

// SetFieldLayout sets the value of the C field 'layout'.
func (recv *LayoutLine) SetFieldLayout(value *Layout) {
	var argValue gi.Argument
	argValue.SetPointer(value.native)
	gi.StructFieldSet(layoutLineStruct, recv.native, "layout", argValue)
}

// FieldStartIndex returns the C field 'start_index'.
func (recv *LayoutLine) FieldStartIndex() int32 {
	argValue := gi.StructFieldGet(layoutLineStruct, recv.native, "start_index")
	value := argValue.Int32()
	return value
}

// SetFieldStartIndex sets the value of the C field 'start_index'.
func (recv *LayoutLine) SetFieldStartIndex(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(layoutLineStruct, recv.native, "start_index", argValue)
}

// FieldLength returns the C field 'length'.
func (recv *LayoutLine) FieldLength() int32 {
	argValue := gi.StructFieldGet(layoutLineStruct, recv.native, "length")
	value := argValue.Int32()
	return value
}

// SetFieldLength sets the value of the C field 'length'.
func (recv *LayoutLine) SetFieldLength(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(layoutLineStruct, recv.native, "length", argValue)
}

// UNSUPPORTED : C value 'runs' : for field getter : no Go type for 'GLib.SList'

// UNSUPPORTED : C value 'runs' : for field setter : no Go type for 'GLib.SList'

// FieldIsParagraphStart returns the C field 'is_paragraph_start'.
func (recv *LayoutLine) FieldIsParagraphStart() uint32 {
	argValue := gi.StructFieldGet(layoutLineStruct, recv.native, "is_paragraph_start")
	value := argValue.Uint32()
	return value
}

// SetFieldIsParagraphStart sets the value of the C field 'is_paragraph_start'.
func (recv *LayoutLine) SetFieldIsParagraphStart(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(layoutLineStruct, recv.native, "is_paragraph_start", argValue)
}

// FieldResolvedDir returns the C field 'resolved_dir'.
func (recv *LayoutLine) FieldResolvedDir() uint32 {
	argValue := gi.StructFieldGet(layoutLineStruct, recv.native, "resolved_dir")
	value := argValue.Uint32()
	return value
}

// SetFieldResolvedDir sets the value of the C field 'resolved_dir'.
func (recv *LayoutLine) SetFieldResolvedDir(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(layoutLineStruct, recv.native, "resolved_dir", argValue)
}

var layoutLineGetExtentsFunction *gi.Function
var layoutLineGetExtentsFunction_Once sync.Once

func layoutLineGetExtentsFunction_Set() error {
	var err error
	layoutLineGetExtentsFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineGetExtentsFunction, err = layoutLineStruct.InvokerNew("get_extents")
	})
	return err
}

// GetExtents is a representation of the C type pango_layout_line_get_extents.
func (recv *LayoutLine) GetExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutLineGetExtentsFunction_Set()
	if err == nil {
		layoutLineGetExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

var layoutLineGetPixelExtentsFunction *gi.Function
var layoutLineGetPixelExtentsFunction_Once sync.Once

func layoutLineGetPixelExtentsFunction_Set() error {
	var err error
	layoutLineGetPixelExtentsFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineGetPixelExtentsFunction, err = layoutLineStruct.InvokerNew("get_pixel_extents")
	})
	return err
}

// GetPixelExtents is a representation of the C type pango_layout_line_get_pixel_extents.
func (recv *LayoutLine) GetPixelExtents() (*Rectangle, *Rectangle) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := layoutLineGetPixelExtentsFunction_Set()
	if err == nil {
		layoutLineGetPixelExtentsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())
	out1 := RectangleNewFromNative(outArgs[1].Pointer())

	return out0, out1
}

// UNSUPPORTED : C value 'pango_layout_line_get_x_ranges' : parameter 'ranges' of type 'nil' not supported

var layoutLineIndexToXFunction *gi.Function
var layoutLineIndexToXFunction_Once sync.Once

func layoutLineIndexToXFunction_Set() error {
	var err error
	layoutLineIndexToXFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineIndexToXFunction, err = layoutLineStruct.InvokerNew("index_to_x")
	})
	return err
}

// IndexToX is a representation of the C type pango_layout_line_index_to_x.
func (recv *LayoutLine) IndexToX(index int32, trailing bool) int32 {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(index)
	inArgs[2].SetBoolean(trailing)

	var outArgs [1]gi.Argument

	err := layoutLineIndexToXFunction_Set()
	if err == nil {
		layoutLineIndexToXFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Int32()

	return out0
}

var layoutLineRefFunction *gi.Function
var layoutLineRefFunction_Once sync.Once

func layoutLineRefFunction_Set() error {
	var err error
	layoutLineRefFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineRefFunction, err = layoutLineStruct.InvokerNew("ref")
	})
	return err
}

// Ref is a representation of the C type pango_layout_line_ref.
func (recv *LayoutLine) Ref() *LayoutLine {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := layoutLineRefFunction_Set()
	if err == nil {
		ret = layoutLineRefFunction.Invoke(inArgs[:], nil)
	}

	retGo := LayoutLineNewFromNative(ret.Pointer())

	return retGo
}

var layoutLineUnrefFunction *gi.Function
var layoutLineUnrefFunction_Once sync.Once

func layoutLineUnrefFunction_Set() error {
	var err error
	layoutLineUnrefFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineUnrefFunction, err = layoutLineStruct.InvokerNew("unref")
	})
	return err
}

// Unref is a representation of the C type pango_layout_line_unref.
func (recv *LayoutLine) Unref() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := layoutLineUnrefFunction_Set()
	if err == nil {
		layoutLineUnrefFunction.Invoke(inArgs[:], nil)
	}

	return
}

var layoutLineXToIndexFunction *gi.Function
var layoutLineXToIndexFunction_Once sync.Once

func layoutLineXToIndexFunction_Set() error {
	var err error
	layoutLineXToIndexFunction_Once.Do(func() {
		err = layoutLineStruct_Set()
		if err != nil {
			return
		}
		layoutLineXToIndexFunction, err = layoutLineStruct.InvokerNew("x_to_index")
	})
	return err
}

// XToIndex is a representation of the C type pango_layout_line_x_to_index.
func (recv *LayoutLine) XToIndex(xPos int32) (bool, int32, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(xPos)

	var outArgs [2]gi.Argument
	var ret gi.Argument

	err := layoutLineXToIndexFunction_Set()
	if err == nil {
		ret = layoutLineXToIndexFunction.Invoke(inArgs[:], outArgs[:])
	}

	retGo := ret.Boolean()
	out0 := outArgs[0].Int32()
	out1 := outArgs[1].Int32()

	return retGo, out0, out1
}

// LayoutLineStruct creates an uninitialised LayoutLine.
func LayoutLineStruct() *LayoutLine {
	err := layoutLineStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LayoutLineNewFromNative(layoutLineStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLayoutLine)
	return structGo
}
func finalizeLayoutLine(obj *LayoutLine) {
	layoutLineStruct.Free(obj.native)
}

var logAttrStruct *gi.Struct
var logAttrStruct_Once sync.Once

func logAttrStruct_Set() error {
	var err error
	logAttrStruct_Once.Do(func() {
		logAttrStruct, err = gi.StructNew("Pango", "LogAttr")
	})
	return err
}

type LogAttr struct {
	native unsafe.Pointer
}

func LogAttrNewFromNative(native unsafe.Pointer) *LogAttr {
	return &LogAttr{native: native}
}

// FieldIsLineBreak returns the C field 'is_line_break'.
func (recv *LogAttr) FieldIsLineBreak() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_line_break")
	value := argValue.Uint32()
	return value
}

// SetFieldIsLineBreak sets the value of the C field 'is_line_break'.
func (recv *LogAttr) SetFieldIsLineBreak(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_line_break", argValue)
}

// FieldIsMandatoryBreak returns the C field 'is_mandatory_break'.
func (recv *LogAttr) FieldIsMandatoryBreak() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_mandatory_break")
	value := argValue.Uint32()
	return value
}

// SetFieldIsMandatoryBreak sets the value of the C field 'is_mandatory_break'.
func (recv *LogAttr) SetFieldIsMandatoryBreak(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_mandatory_break", argValue)
}

// FieldIsCharBreak returns the C field 'is_char_break'.
func (recv *LogAttr) FieldIsCharBreak() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_char_break")
	value := argValue.Uint32()
	return value
}

// SetFieldIsCharBreak sets the value of the C field 'is_char_break'.
func (recv *LogAttr) SetFieldIsCharBreak(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_char_break", argValue)
}

// FieldIsWhite returns the C field 'is_white'.
func (recv *LogAttr) FieldIsWhite() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_white")
	value := argValue.Uint32()
	return value
}

// SetFieldIsWhite sets the value of the C field 'is_white'.
func (recv *LogAttr) SetFieldIsWhite(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_white", argValue)
}

// FieldIsCursorPosition returns the C field 'is_cursor_position'.
func (recv *LogAttr) FieldIsCursorPosition() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_cursor_position")
	value := argValue.Uint32()
	return value
}

// SetFieldIsCursorPosition sets the value of the C field 'is_cursor_position'.
func (recv *LogAttr) SetFieldIsCursorPosition(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_cursor_position", argValue)
}

// FieldIsWordStart returns the C field 'is_word_start'.
func (recv *LogAttr) FieldIsWordStart() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_word_start")
	value := argValue.Uint32()
	return value
}

// SetFieldIsWordStart sets the value of the C field 'is_word_start'.
func (recv *LogAttr) SetFieldIsWordStart(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_word_start", argValue)
}

// FieldIsWordEnd returns the C field 'is_word_end'.
func (recv *LogAttr) FieldIsWordEnd() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_word_end")
	value := argValue.Uint32()
	return value
}

// SetFieldIsWordEnd sets the value of the C field 'is_word_end'.
func (recv *LogAttr) SetFieldIsWordEnd(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_word_end", argValue)
}

// FieldIsSentenceBoundary returns the C field 'is_sentence_boundary'.
func (recv *LogAttr) FieldIsSentenceBoundary() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_sentence_boundary")
	value := argValue.Uint32()
	return value
}

// SetFieldIsSentenceBoundary sets the value of the C field 'is_sentence_boundary'.
func (recv *LogAttr) SetFieldIsSentenceBoundary(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_sentence_boundary", argValue)
}

// FieldIsSentenceStart returns the C field 'is_sentence_start'.
func (recv *LogAttr) FieldIsSentenceStart() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_sentence_start")
	value := argValue.Uint32()
	return value
}

// SetFieldIsSentenceStart sets the value of the C field 'is_sentence_start'.
func (recv *LogAttr) SetFieldIsSentenceStart(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_sentence_start", argValue)
}

// FieldIsSentenceEnd returns the C field 'is_sentence_end'.
func (recv *LogAttr) FieldIsSentenceEnd() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_sentence_end")
	value := argValue.Uint32()
	return value
}

// SetFieldIsSentenceEnd sets the value of the C field 'is_sentence_end'.
func (recv *LogAttr) SetFieldIsSentenceEnd(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_sentence_end", argValue)
}

// FieldBackspaceDeletesCharacter returns the C field 'backspace_deletes_character'.
func (recv *LogAttr) FieldBackspaceDeletesCharacter() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "backspace_deletes_character")
	value := argValue.Uint32()
	return value
}

// SetFieldBackspaceDeletesCharacter sets the value of the C field 'backspace_deletes_character'.
func (recv *LogAttr) SetFieldBackspaceDeletesCharacter(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "backspace_deletes_character", argValue)
}

// FieldIsExpandableSpace returns the C field 'is_expandable_space'.
func (recv *LogAttr) FieldIsExpandableSpace() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_expandable_space")
	value := argValue.Uint32()
	return value
}

// SetFieldIsExpandableSpace sets the value of the C field 'is_expandable_space'.
func (recv *LogAttr) SetFieldIsExpandableSpace(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_expandable_space", argValue)
}

// FieldIsWordBoundary returns the C field 'is_word_boundary'.
func (recv *LogAttr) FieldIsWordBoundary() uint32 {
	argValue := gi.StructFieldGet(logAttrStruct, recv.native, "is_word_boundary")
	value := argValue.Uint32()
	return value
}

// SetFieldIsWordBoundary sets the value of the C field 'is_word_boundary'.
func (recv *LogAttr) SetFieldIsWordBoundary(value uint32) {
	var argValue gi.Argument
	argValue.SetUint32(value)
	gi.StructFieldSet(logAttrStruct, recv.native, "is_word_boundary", argValue)
}

// LogAttrStruct creates an uninitialised LogAttr.
func LogAttrStruct() *LogAttr {
	err := logAttrStruct_Set()
	if err != nil {
		return nil
	}

	structGo := LogAttrNewFromNative(logAttrStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeLogAttr)
	return structGo
}
func finalizeLogAttr(obj *LogAttr) {
	logAttrStruct.Free(obj.native)
}

var mapStruct *gi.Struct
var mapStruct_Once sync.Once

func mapStruct_Set() error {
	var err error
	mapStruct_Once.Do(func() {
		mapStruct, err = gi.StructNew("Pango", "Map")
	})
	return err
}

type Map struct {
	native unsafe.Pointer
}

func MapNewFromNative(native unsafe.Pointer) *Map {
	return &Map{native: native}
}

var mapGetEngineFunction *gi.Function
var mapGetEngineFunction_Once sync.Once

func mapGetEngineFunction_Set() error {
	var err error
	mapGetEngineFunction_Once.Do(func() {
		err = mapStruct_Set()
		if err != nil {
			return
		}
		mapGetEngineFunction, err = mapStruct.InvokerNew("get_engine")
	})
	return err
}

// GetEngine is a representation of the C type pango_map_get_engine.
func (recv *Map) GetEngine(script Script) *Engine {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(int32(script))

	var ret gi.Argument

	err := mapGetEngineFunction_Set()
	if err == nil {
		ret = mapGetEngineFunction.Invoke(inArgs[:], nil)
	}

	retGo := EngineNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_map_get_engines' : parameter 'exact_engines' of type 'GLib.SList' not supported

// MapStruct creates an uninitialised Map.
func MapStruct() *Map {
	err := mapStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MapNewFromNative(mapStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMap)
	return structGo
}
func finalizeMap(obj *Map) {
	mapStruct.Free(obj.native)
}

var mapEntryStruct *gi.Struct
var mapEntryStruct_Once sync.Once

func mapEntryStruct_Set() error {
	var err error
	mapEntryStruct_Once.Do(func() {
		mapEntryStruct, err = gi.StructNew("Pango", "MapEntry")
	})
	return err
}

type MapEntry struct {
	native unsafe.Pointer
}

func MapEntryNewFromNative(native unsafe.Pointer) *MapEntry {
	return &MapEntry{native: native}
}

// MapEntryStruct creates an uninitialised MapEntry.
func MapEntryStruct() *MapEntry {
	err := mapEntryStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MapEntryNewFromNative(mapEntryStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMapEntry)
	return structGo
}
func finalizeMapEntry(obj *MapEntry) {
	mapEntryStruct.Free(obj.native)
}

var matrixStruct *gi.Struct
var matrixStruct_Once sync.Once

func matrixStruct_Set() error {
	var err error
	matrixStruct_Once.Do(func() {
		matrixStruct, err = gi.StructNew("Pango", "Matrix")
	})
	return err
}

type Matrix struct {
	native unsafe.Pointer
}

func MatrixNewFromNative(native unsafe.Pointer) *Matrix {
	return &Matrix{native: native}
}

// FieldXx returns the C field 'xx'.
func (recv *Matrix) FieldXx() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "xx")
	value := argValue.Float64()
	return value
}

// SetFieldXx sets the value of the C field 'xx'.
func (recv *Matrix) SetFieldXx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "xx", argValue)
}

// FieldXy returns the C field 'xy'.
func (recv *Matrix) FieldXy() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "xy")
	value := argValue.Float64()
	return value
}

// SetFieldXy sets the value of the C field 'xy'.
func (recv *Matrix) SetFieldXy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "xy", argValue)
}

// FieldYx returns the C field 'yx'.
func (recv *Matrix) FieldYx() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "yx")
	value := argValue.Float64()
	return value
}

// SetFieldYx sets the value of the C field 'yx'.
func (recv *Matrix) SetFieldYx(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "yx", argValue)
}

// FieldYy returns the C field 'yy'.
func (recv *Matrix) FieldYy() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "yy")
	value := argValue.Float64()
	return value
}

// SetFieldYy sets the value of the C field 'yy'.
func (recv *Matrix) SetFieldYy(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "yy", argValue)
}

// FieldX0 returns the C field 'x0'.
func (recv *Matrix) FieldX0() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "x0")
	value := argValue.Float64()
	return value
}

// SetFieldX0 sets the value of the C field 'x0'.
func (recv *Matrix) SetFieldX0(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "x0", argValue)
}

// FieldY0 returns the C field 'y0'.
func (recv *Matrix) FieldY0() float64 {
	argValue := gi.StructFieldGet(matrixStruct, recv.native, "y0")
	value := argValue.Float64()
	return value
}

// SetFieldY0 sets the value of the C field 'y0'.
func (recv *Matrix) SetFieldY0(value float64) {
	var argValue gi.Argument
	argValue.SetFloat64(value)
	gi.StructFieldSet(matrixStruct, recv.native, "y0", argValue)
}

var matrixConcatFunction *gi.Function
var matrixConcatFunction_Once sync.Once

func matrixConcatFunction_Set() error {
	var err error
	matrixConcatFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixConcatFunction, err = matrixStruct.InvokerNew("concat")
	})
	return err
}

// Concat is a representation of the C type pango_matrix_concat.
func (recv *Matrix) Concat(newMatrix *Matrix) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(newMatrix.native)

	err := matrixConcatFunction_Set()
	if err == nil {
		matrixConcatFunction.Invoke(inArgs[:], nil)
	}

	return
}

var matrixCopyFunction *gi.Function
var matrixCopyFunction_Once sync.Once

func matrixCopyFunction_Set() error {
	var err error
	matrixCopyFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixCopyFunction, err = matrixStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_matrix_copy.
func (recv *Matrix) Copy() *Matrix {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matrixCopyFunction_Set()
	if err == nil {
		ret = matrixCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := MatrixNewFromNative(ret.Pointer())

	return retGo
}

var matrixFreeFunction *gi.Function
var matrixFreeFunction_Once sync.Once

func matrixFreeFunction_Set() error {
	var err error
	matrixFreeFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixFreeFunction, err = matrixStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_matrix_free.
func (recv *Matrix) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := matrixFreeFunction_Set()
	if err == nil {
		matrixFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var matrixGetFontScaleFactorFunction *gi.Function
var matrixGetFontScaleFactorFunction_Once sync.Once

func matrixGetFontScaleFactorFunction_Set() error {
	var err error
	matrixGetFontScaleFactorFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixGetFontScaleFactorFunction, err = matrixStruct.InvokerNew("get_font_scale_factor")
	})
	return err
}

// GetFontScaleFactor is a representation of the C type pango_matrix_get_font_scale_factor.
func (recv *Matrix) GetFontScaleFactor() float64 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := matrixGetFontScaleFactorFunction_Set()
	if err == nil {
		ret = matrixGetFontScaleFactorFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Float64()

	return retGo
}

var matrixGetFontScaleFactorsFunction *gi.Function
var matrixGetFontScaleFactorsFunction_Once sync.Once

func matrixGetFontScaleFactorsFunction_Set() error {
	var err error
	matrixGetFontScaleFactorsFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixGetFontScaleFactorsFunction, err = matrixStruct.InvokerNew("get_font_scale_factors")
	})
	return err
}

// GetFontScaleFactors is a representation of the C type pango_matrix_get_font_scale_factors.
func (recv *Matrix) GetFontScaleFactors() (float64, float64) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [2]gi.Argument

	err := matrixGetFontScaleFactorsFunction_Set()
	if err == nil {
		matrixGetFontScaleFactorsFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1
}

var matrixRotateFunction *gi.Function
var matrixRotateFunction_Once sync.Once

func matrixRotateFunction_Set() error {
	var err error
	matrixRotateFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixRotateFunction, err = matrixStruct.InvokerNew("rotate")
	})
	return err
}

// Rotate is a representation of the C type pango_matrix_rotate.
func (recv *Matrix) Rotate(degrees float64) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(degrees)

	err := matrixRotateFunction_Set()
	if err == nil {
		matrixRotateFunction.Invoke(inArgs[:], nil)
	}

	return
}

var matrixScaleFunction *gi.Function
var matrixScaleFunction_Once sync.Once

func matrixScaleFunction_Set() error {
	var err error
	matrixScaleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixScaleFunction, err = matrixStruct.InvokerNew("scale")
	})
	return err
}

// Scale is a representation of the C type pango_matrix_scale.
func (recv *Matrix) Scale(scaleX float64, scaleY float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(scaleX)
	inArgs[2].SetFloat64(scaleY)

	err := matrixScaleFunction_Set()
	if err == nil {
		matrixScaleFunction.Invoke(inArgs[:], nil)
	}

	return
}

var matrixTransformDistanceFunction *gi.Function
var matrixTransformDistanceFunction_Once sync.Once

func matrixTransformDistanceFunction_Set() error {
	var err error
	matrixTransformDistanceFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformDistanceFunction, err = matrixStruct.InvokerNew("transform_distance")
	})
	return err
}

// TransformDistance is a representation of the C type pango_matrix_transform_distance.
func (recv *Matrix) TransformDistance(dx float64, dy float64) (float64, float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(dx)
	inArgs[2].SetFloat64(dy)

	var outArgs [2]gi.Argument

	err := matrixTransformDistanceFunction_Set()
	if err == nil {
		matrixTransformDistanceFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1
}

var matrixTransformPixelRectangleFunction *gi.Function
var matrixTransformPixelRectangleFunction_Once sync.Once

func matrixTransformPixelRectangleFunction_Set() error {
	var err error
	matrixTransformPixelRectangleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformPixelRectangleFunction, err = matrixStruct.InvokerNew("transform_pixel_rectangle")
	})
	return err
}

// TransformPixelRectangle is a representation of the C type pango_matrix_transform_pixel_rectangle.
func (recv *Matrix) TransformPixelRectangle(rect *Rectangle) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rect.native)

	var outArgs [1]gi.Argument

	err := matrixTransformPixelRectangleFunction_Set()
	if err == nil {
		matrixTransformPixelRectangleFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

var matrixTransformPointFunction *gi.Function
var matrixTransformPointFunction_Once sync.Once

func matrixTransformPointFunction_Set() error {
	var err error
	matrixTransformPointFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformPointFunction, err = matrixStruct.InvokerNew("transform_point")
	})
	return err
}

// TransformPoint is a representation of the C type pango_matrix_transform_point.
func (recv *Matrix) TransformPoint(x float64, y float64) (float64, float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(x)
	inArgs[2].SetFloat64(y)

	var outArgs [2]gi.Argument

	err := matrixTransformPointFunction_Set()
	if err == nil {
		matrixTransformPointFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].Float64()
	out1 := outArgs[1].Float64()

	return out0, out1
}

var matrixTransformRectangleFunction *gi.Function
var matrixTransformRectangleFunction_Once sync.Once

func matrixTransformRectangleFunction_Set() error {
	var err error
	matrixTransformRectangleFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTransformRectangleFunction, err = matrixStruct.InvokerNew("transform_rectangle")
	})
	return err
}

// TransformRectangle is a representation of the C type pango_matrix_transform_rectangle.
func (recv *Matrix) TransformRectangle(rect *Rectangle) *Rectangle {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetPointer(rect.native)

	var outArgs [1]gi.Argument

	err := matrixTransformRectangleFunction_Set()
	if err == nil {
		matrixTransformRectangleFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := RectangleNewFromNative(outArgs[0].Pointer())

	return out0
}

var matrixTranslateFunction *gi.Function
var matrixTranslateFunction_Once sync.Once

func matrixTranslateFunction_Set() error {
	var err error
	matrixTranslateFunction_Once.Do(func() {
		err = matrixStruct_Set()
		if err != nil {
			return
		}
		matrixTranslateFunction, err = matrixStruct.InvokerNew("translate")
	})
	return err
}

// Translate is a representation of the C type pango_matrix_translate.
func (recv *Matrix) Translate(tx float64, ty float64) {
	var inArgs [3]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetFloat64(tx)
	inArgs[2].SetFloat64(ty)

	err := matrixTranslateFunction_Set()
	if err == nil {
		matrixTranslateFunction.Invoke(inArgs[:], nil)
	}

	return
}

// MatrixStruct creates an uninitialised Matrix.
func MatrixStruct() *Matrix {
	err := matrixStruct_Set()
	if err != nil {
		return nil
	}

	structGo := MatrixNewFromNative(matrixStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeMatrix)
	return structGo
}
func finalizeMatrix(obj *Matrix) {
	matrixStruct.Free(obj.native)
}

var rectangleStruct *gi.Struct
var rectangleStruct_Once sync.Once

func rectangleStruct_Set() error {
	var err error
	rectangleStruct_Once.Do(func() {
		rectangleStruct, err = gi.StructNew("Pango", "Rectangle")
	})
	return err
}

type Rectangle struct {
	native unsafe.Pointer
}

func RectangleNewFromNative(native unsafe.Pointer) *Rectangle {
	return &Rectangle{native: native}
}

// FieldX returns the C field 'x'.
func (recv *Rectangle) FieldX() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "x")
	value := argValue.Int32()
	return value
}

// SetFieldX sets the value of the C field 'x'.
func (recv *Rectangle) SetFieldX(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "x", argValue)
}

// FieldY returns the C field 'y'.
func (recv *Rectangle) FieldY() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "y")
	value := argValue.Int32()
	return value
}

// SetFieldY sets the value of the C field 'y'.
func (recv *Rectangle) SetFieldY(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "y", argValue)
}

// FieldWidth returns the C field 'width'.
func (recv *Rectangle) FieldWidth() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "width")
	value := argValue.Int32()
	return value
}

// SetFieldWidth sets the value of the C field 'width'.
func (recv *Rectangle) SetFieldWidth(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "width", argValue)
}

// FieldHeight returns the C field 'height'.
func (recv *Rectangle) FieldHeight() int32 {
	argValue := gi.StructFieldGet(rectangleStruct, recv.native, "height")
	value := argValue.Int32()
	return value
}

// SetFieldHeight sets the value of the C field 'height'.
func (recv *Rectangle) SetFieldHeight(value int32) {
	var argValue gi.Argument
	argValue.SetInt32(value)
	gi.StructFieldSet(rectangleStruct, recv.native, "height", argValue)
}

// RectangleStruct creates an uninitialised Rectangle.
func RectangleStruct() *Rectangle {
	err := rectangleStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RectangleNewFromNative(rectangleStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRectangle)
	return structGo
}
func finalizeRectangle(obj *Rectangle) {
	rectangleStruct.Free(obj.native)
}

var rendererClassStruct *gi.Struct
var rendererClassStruct_Once sync.Once

func rendererClassStruct_Set() error {
	var err error
	rendererClassStruct_Once.Do(func() {
		rendererClassStruct, err = gi.StructNew("Pango", "RendererClass")
	})
	return err
}

type RendererClass struct {
	native unsafe.Pointer
}

func RendererClassNewFromNative(native unsafe.Pointer) *RendererClass {
	return &RendererClass{native: native}
}

// UNSUPPORTED : C value 'draw_glyphs' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_glyphs' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_rectangle' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_rectangle' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_error_underline' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_error_underline' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_shape' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_shape' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_trapezoid' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_trapezoid' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_glyph' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_glyph' : for field setter : missing Type

// UNSUPPORTED : C value 'part_changed' : for field getter : missing Type

// UNSUPPORTED : C value 'part_changed' : for field setter : missing Type

// UNSUPPORTED : C value 'begin' : for field getter : missing Type

// UNSUPPORTED : C value 'begin' : for field setter : missing Type

// UNSUPPORTED : C value 'end' : for field getter : missing Type

// UNSUPPORTED : C value 'end' : for field setter : missing Type

// UNSUPPORTED : C value 'prepare_run' : for field getter : missing Type

// UNSUPPORTED : C value 'prepare_run' : for field setter : missing Type

// UNSUPPORTED : C value 'draw_glyph_item' : for field getter : missing Type

// UNSUPPORTED : C value 'draw_glyph_item' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved2' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved3' : for field setter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field getter : missing Type

// UNSUPPORTED : C value '_pango_reserved4' : for field setter : missing Type

// RendererClassStruct creates an uninitialised RendererClass.
func RendererClassStruct() *RendererClass {
	err := rendererClassStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RendererClassNewFromNative(rendererClassStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRendererClass)
	return structGo
}
func finalizeRendererClass(obj *RendererClass) {
	rendererClassStruct.Free(obj.native)
}

var rendererPrivateStruct *gi.Struct
var rendererPrivateStruct_Once sync.Once

func rendererPrivateStruct_Set() error {
	var err error
	rendererPrivateStruct_Once.Do(func() {
		rendererPrivateStruct, err = gi.StructNew("Pango", "RendererPrivate")
	})
	return err
}

type RendererPrivate struct {
	native unsafe.Pointer
}

func RendererPrivateNewFromNative(native unsafe.Pointer) *RendererPrivate {
	return &RendererPrivate{native: native}
}

// RendererPrivateStruct creates an uninitialised RendererPrivate.
func RendererPrivateStruct() *RendererPrivate {
	err := rendererPrivateStruct_Set()
	if err != nil {
		return nil
	}

	structGo := RendererPrivateNewFromNative(rendererPrivateStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeRendererPrivate)
	return structGo
}
func finalizeRendererPrivate(obj *RendererPrivate) {
	rendererPrivateStruct.Free(obj.native)
}

var scriptIterStruct *gi.Struct
var scriptIterStruct_Once sync.Once

func scriptIterStruct_Set() error {
	var err error
	scriptIterStruct_Once.Do(func() {
		scriptIterStruct, err = gi.StructNew("Pango", "ScriptIter")
	})
	return err
}

type ScriptIter struct {
	native unsafe.Pointer
}

func ScriptIterNewFromNative(native unsafe.Pointer) *ScriptIter {
	return &ScriptIter{native: native}
}

var scriptIterFreeFunction *gi.Function
var scriptIterFreeFunction_Once sync.Once

func scriptIterFreeFunction_Set() error {
	var err error
	scriptIterFreeFunction_Once.Do(func() {
		err = scriptIterStruct_Set()
		if err != nil {
			return
		}
		scriptIterFreeFunction, err = scriptIterStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_script_iter_free.
func (recv *ScriptIter) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := scriptIterFreeFunction_Set()
	if err == nil {
		scriptIterFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var scriptIterGetRangeFunction *gi.Function
var scriptIterGetRangeFunction_Once sync.Once

func scriptIterGetRangeFunction_Set() error {
	var err error
	scriptIterGetRangeFunction_Once.Do(func() {
		err = scriptIterStruct_Set()
		if err != nil {
			return
		}
		scriptIterGetRangeFunction, err = scriptIterStruct.InvokerNew("get_range")
	})
	return err
}

// GetRange is a representation of the C type pango_script_iter_get_range.
func (recv *ScriptIter) GetRange() (string, string, Script) {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var outArgs [3]gi.Argument

	err := scriptIterGetRangeFunction_Set()
	if err == nil {
		scriptIterGetRangeFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := outArgs[0].String(true)
	out1 := outArgs[1].String(true)
	out2 := Script(outArgs[2].Int32())

	return out0, out1, out2
}

var scriptIterNextFunction *gi.Function
var scriptIterNextFunction_Once sync.Once

func scriptIterNextFunction_Set() error {
	var err error
	scriptIterNextFunction_Once.Do(func() {
		err = scriptIterStruct_Set()
		if err != nil {
			return
		}
		scriptIterNextFunction, err = scriptIterStruct.InvokerNew("next")
	})
	return err
}

// Next is a representation of the C type pango_script_iter_next.
func (recv *ScriptIter) Next() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := scriptIterNextFunction_Set()
	if err == nil {
		ret = scriptIterNextFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

// ScriptIterStruct creates an uninitialised ScriptIter.
func ScriptIterStruct() *ScriptIter {
	err := scriptIterStruct_Set()
	if err != nil {
		return nil
	}

	structGo := ScriptIterNewFromNative(scriptIterStruct.Alloc())
	runtime.SetFinalizer(structGo, finalizeScriptIter)
	return structGo
}
func finalizeScriptIter(obj *ScriptIter) {
	scriptIterStruct.Free(obj.native)
}

var tabArrayStruct *gi.Struct
var tabArrayStruct_Once sync.Once

func tabArrayStruct_Set() error {
	var err error
	tabArrayStruct_Once.Do(func() {
		tabArrayStruct, err = gi.StructNew("Pango", "TabArray")
	})
	return err
}

type TabArray struct {
	native unsafe.Pointer
}

func TabArrayNewFromNative(native unsafe.Pointer) *TabArray {
	return &TabArray{native: native}
}

var tabArrayNewFunction *gi.Function
var tabArrayNewFunction_Once sync.Once

func tabArrayNewFunction_Set() error {
	var err error
	tabArrayNewFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayNewFunction, err = tabArrayStruct.InvokerNew("new")
	})
	return err
}

// TabArrayNew is a representation of the C type pango_tab_array_new.
func TabArrayNew(initialSize int32, positionsInPixels bool) *TabArray {
	var inArgs [2]gi.Argument
	inArgs[0].SetInt32(initialSize)
	inArgs[1].SetBoolean(positionsInPixels)

	var ret gi.Argument

	err := tabArrayNewFunction_Set()
	if err == nil {
		ret = tabArrayNewFunction.Invoke(inArgs[:], nil)
	}

	retGo := TabArrayNewFromNative(ret.Pointer())

	return retGo
}

// UNSUPPORTED : C value 'pango_tab_array_new_with_positions' : parameter '...' of type 'nil' not supported

var tabArrayCopyFunction *gi.Function
var tabArrayCopyFunction_Once sync.Once

func tabArrayCopyFunction_Set() error {
	var err error
	tabArrayCopyFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayCopyFunction, err = tabArrayStruct.InvokerNew("copy")
	})
	return err
}

// Copy is a representation of the C type pango_tab_array_copy.
func (recv *TabArray) Copy() *TabArray {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayCopyFunction_Set()
	if err == nil {
		ret = tabArrayCopyFunction.Invoke(inArgs[:], nil)
	}

	retGo := TabArrayNewFromNative(ret.Pointer())

	return retGo
}

var tabArrayFreeFunction *gi.Function
var tabArrayFreeFunction_Once sync.Once

func tabArrayFreeFunction_Set() error {
	var err error
	tabArrayFreeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayFreeFunction, err = tabArrayStruct.InvokerNew("free")
	})
	return err
}

// Free is a representation of the C type pango_tab_array_free.
func (recv *TabArray) Free() {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	err := tabArrayFreeFunction_Set()
	if err == nil {
		tabArrayFreeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tabArrayGetPositionsInPixelsFunction *gi.Function
var tabArrayGetPositionsInPixelsFunction_Once sync.Once

func tabArrayGetPositionsInPixelsFunction_Set() error {
	var err error
	tabArrayGetPositionsInPixelsFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayGetPositionsInPixelsFunction, err = tabArrayStruct.InvokerNew("get_positions_in_pixels")
	})
	return err
}

// GetPositionsInPixels is a representation of the C type pango_tab_array_get_positions_in_pixels.
func (recv *TabArray) GetPositionsInPixels() bool {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayGetPositionsInPixelsFunction_Set()
	if err == nil {
		ret = tabArrayGetPositionsInPixelsFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Boolean()

	return retGo
}

var tabArrayGetSizeFunction *gi.Function
var tabArrayGetSizeFunction_Once sync.Once

func tabArrayGetSizeFunction_Set() error {
	var err error
	tabArrayGetSizeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayGetSizeFunction, err = tabArrayStruct.InvokerNew("get_size")
	})
	return err
}

// GetSize is a representation of the C type pango_tab_array_get_size.
func (recv *TabArray) GetSize() int32 {
	var inArgs [1]gi.Argument
	inArgs[0].SetPointer(recv.native)

	var ret gi.Argument

	err := tabArrayGetSizeFunction_Set()
	if err == nil {
		ret = tabArrayGetSizeFunction.Invoke(inArgs[:], nil)
	}

	retGo := ret.Int32()

	return retGo
}

var tabArrayGetTabFunction *gi.Function
var tabArrayGetTabFunction_Once sync.Once

func tabArrayGetTabFunction_Set() error {
	var err error
	tabArrayGetTabFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayGetTabFunction, err = tabArrayStruct.InvokerNew("get_tab")
	})
	return err
}

// GetTab is a representation of the C type pango_tab_array_get_tab.
func (recv *TabArray) GetTab(tabIndex int32) (TabAlign, int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(tabIndex)

	var outArgs [2]gi.Argument

	err := tabArrayGetTabFunction_Set()
	if err == nil {
		tabArrayGetTabFunction.Invoke(inArgs[:], outArgs[:])
	}

	out0 := TabAlign(outArgs[0].Int32())
	out1 := outArgs[1].Int32()

	return out0, out1
}

// UNSUPPORTED : C value 'pango_tab_array_get_tabs' : parameter 'locations' of type 'nil' not supported

var tabArrayResizeFunction *gi.Function
var tabArrayResizeFunction_Once sync.Once

func tabArrayResizeFunction_Set() error {
	var err error
	tabArrayResizeFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArrayResizeFunction, err = tabArrayStruct.InvokerNew("resize")
	})
	return err
}

// Resize is a representation of the C type pango_tab_array_resize.
func (recv *TabArray) Resize(newSize int32) {
	var inArgs [2]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(newSize)

	err := tabArrayResizeFunction_Set()
	if err == nil {
		tabArrayResizeFunction.Invoke(inArgs[:], nil)
	}

	return
}

var tabArraySetTabFunction *gi.Function
var tabArraySetTabFunction_Once sync.Once

func tabArraySetTabFunction_Set() error {
	var err error
	tabArraySetTabFunction_Once.Do(func() {
		err = tabArrayStruct_Set()
		if err != nil {
			return
		}
		tabArraySetTabFunction, err = tabArrayStruct.InvokerNew("set_tab")
	})
	return err
}

// SetTab is a representation of the C type pango_tab_array_set_tab.
func (recv *TabArray) SetTab(tabIndex int32, alignment TabAlign, location int32) {
	var inArgs [4]gi.Argument
	inArgs[0].SetPointer(recv.native)
	inArgs[1].SetInt32(tabIndex)
	inArgs[2].SetInt32(int32(alignment))
	inArgs[3].SetInt32(location)

	err := tabArraySetTabFunction_Set()
	if err == nil {
		tabArraySetTabFunction.Invoke(inArgs[:], nil)
	}

	return
}
