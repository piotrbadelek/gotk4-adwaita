// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_preferences_group_get_type()), F: marshalPreferencesGrouper},
	})
}

// PreferencesGroup: group of preference rows.
//
// An AdwPreferencesGroup represents a group or tightly related preferences,
// which in turn are represented by adw.PreferencesRow.
//
// To summarize the role of the preferences it gathers, a group can have both a
// title and a description. The title will be used by adw.PreferencesWindow to
// let the user look for a preference.
//
//
// CSS nodes
//
// AdwPreferencesGroup has a single CSS node with name preferencesgroup.
type PreferencesGroup struct {
	gtk.Widget
}

func wrapPreferencesGroup(obj *externglib.Object) *PreferencesGroup {
	return &PreferencesGroup{
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

func marshalPreferencesGrouper(p uintptr) (interface{}, error) {
	return wrapPreferencesGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesGroup creates a new AdwPreferencesGroup.
func NewPreferencesGroup() *PreferencesGroup {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_group_new()

	var _preferencesGroup *PreferencesGroup // out

	_preferencesGroup = wrapPreferencesGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesGroup
}

// Add adds a child to self.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (self *PreferencesGroup) Add(child gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_preferences_group_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// Description gets the description of self.
func (self *PreferencesGroup) Description() string {
	var _arg0 *C.AdwPreferencesGroup // out
	var _cret *C.char                // in

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_group_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self.
func (self *PreferencesGroup) Title() string {
	var _arg0 *C.AdwPreferencesGroup // out
	var _cret *C.char                // in

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_group_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Remove removes a child from self.
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (self *PreferencesGroup) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_preferences_group_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDescription sets the description for self.
//
// The function takes the following parameters:
//
//    - description: description.
//
func (self *PreferencesGroup) SetDescription(description string) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.char                // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))
	if description != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_group_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetTitle sets the title for self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *PreferencesGroup) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesGroup // out
	var _arg1 *C.char                // out

	_arg0 = (*C.AdwPreferencesGroup)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_group_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}
