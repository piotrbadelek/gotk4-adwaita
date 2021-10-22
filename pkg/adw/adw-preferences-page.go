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
		{T: externglib.Type(C.adw_preferences_page_get_type()), F: marshalPreferencesPager},
	})
}

// PreferencesPage: page from adw.PreferencesWindow.
//
// The AdwPreferencesPage widget gathers preferences groups into a single page
// of a preferences window.
//
//
// CSS nodes
//
// AdwPreferencesPage has a single CSS node with name preferencespage.
type PreferencesPage struct {
	gtk.Widget
}

func wrapPreferencesPage(obj *externglib.Object) *PreferencesPage {
	return &PreferencesPage{
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

func marshalPreferencesPager(p uintptr) (interface{}, error) {
	return wrapPreferencesPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesPage creates a new AdwPreferencesPage.
func NewPreferencesPage() *PreferencesPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_page_new()

	var _preferencesPage *PreferencesPage // out

	_preferencesPage = wrapPreferencesPage(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesPage
}

// Add adds a preferences group to self.
//
// The function takes the following parameters:
//
//    - group to add.
//
func (self *PreferencesPage) Add(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(group.Native()))

	C.adw_preferences_page_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(group)
}

// IconName gets the icon name for self.
func (self *PreferencesPage) IconName() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of self.
func (self *PreferencesPage) Name() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self.
func (self *PreferencesPage) Title() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// UseUnderline gets whether an embedded underline in the title indicates a
// mnemonic.
func (self *PreferencesPage) UseUnderline() bool {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret C.gboolean            // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes a group from self.
//
// The function takes the following parameters:
//
//    - group to remove.
//
func (self *PreferencesPage) Remove(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(group.Native()))

	C.adw_preferences_page_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(group)
}

// SetIconName sets the icon name for self.
//
// The function takes the following parameters:
//
//    - iconName: icon name.
//
func (self *PreferencesPage) SetIconName(iconName string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetName sets the name of self.
//
// The function takes the following parameters:
//
//    - name: name.
//
func (self *PreferencesPage) SetName(name string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetTitle sets the title of self.
//
// The function takes the following parameters:
//
//    - title: title.
//
func (self *PreferencesPage) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetUseUnderline sets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (self *PreferencesPage) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_preferences_page_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}
