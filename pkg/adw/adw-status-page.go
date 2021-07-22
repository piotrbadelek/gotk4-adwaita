// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_status_page_get_type()), F: marshalStatusPager},
	})
}

type StatusPage struct {
	gtk.Widget
}

func wrapStatusPage(obj *externglib.Object) *StatusPage {
	return &StatusPage{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalStatusPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStatusPage(obj), nil
}

// NewStatusPage creates a new StatusPage.
func NewStatusPage() *StatusPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_status_page_new()

	var _statusPage *StatusPage // out

	_statusPage = wrapStatusPage(externglib.Take(unsafe.Pointer(_cret)))

	return _statusPage
}

// Child gets the child widget of self.
func (self *StatusPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.GtkWidget     // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// Description gets the description for self.
func (self *StatusPage) Description() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IconName gets the icon name for self.
func (self *StatusPage) IconName() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Title gets the title for self.
func (self *StatusPage) Title() string {
	var _arg0 *C.AdwStatusPage // out
	var _cret *C.char          // in

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_status_page_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetChild sets the child widget of self.
func (self *StatusPage) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_status_page_set_child(_arg0, _arg1)
}

// SetDescription sets the description for self.
func (self *StatusPage) SetDescription(description string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_status_page_set_description(_arg0, _arg1)
}

// SetIconName sets the icon name for self.
func (self *StatusPage) SetIconName(iconName string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_status_page_set_icon_name(_arg0, _arg1)
}

// SetTitle sets the title for self.
func (self *StatusPage) SetTitle(title string) {
	var _arg0 *C.AdwStatusPage // out
	var _arg1 *C.char          // out

	_arg0 = (*C.AdwStatusPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_status_page_set_title(_arg0, _arg1)
}
