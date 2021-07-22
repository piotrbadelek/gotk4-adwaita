// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_enum_list_model_get_type()), F: marshalEnumListModeller},
	})
}

type EnumListModel struct {
	*externglib.Object

	gio.ListModel
}

func wrapEnumListModel(obj *externglib.Object) *EnumListModel {
	return &EnumListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalEnumListModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEnumListModel(obj), nil
}

func NewEnumListModel(enumType externglib.Type) *EnumListModel {
	var _arg1 C.GType             // out
	var _cret *C.AdwEnumListModel // in

	_arg1 = C.GType(enumType)

	_cret = C.adw_enum_list_model_new(_arg1)

	var _enumListModel *EnumListModel // out

	_enumListModel = wrapEnumListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _enumListModel
}

func (self *EnumListModel) FindPosition(value int) uint {
	var _arg0 *C.AdwEnumListModel // out
	var _arg1 C.int               // out
	var _cret C.guint             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(value)

	_cret = C.adw_enum_list_model_find_position(_arg0, _arg1)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (self *EnumListModel) EnumType() externglib.Type {
	var _arg0 *C.AdwEnumListModel // out
	var _cret C.GType             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(self.Native()))

	_cret = C.adw_enum_list_model_get_enum_type(_arg0)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}
