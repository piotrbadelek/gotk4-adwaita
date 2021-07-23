// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

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
		{T: externglib.Type(C.adw_preferences_page_get_type()), F: marshalPreferencesPager},
	})
}

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
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPreferencesPage(obj), nil
}

// NewPreferencesPage creates a new PreferencesPage.
func NewPreferencesPage() *PreferencesPage {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_page_new()

	var _preferencesPage *PreferencesPage // out

	_preferencesPage = wrapPreferencesPage(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesPage
}

func (self *PreferencesPage) Add(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(group.Native()))

	C.adw_preferences_page_add(_arg0, _arg1)
}

// IconName gets the icon name for self, or NULL.
func (self *PreferencesPage) IconName() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_icon_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self, or NULL.
func (self *PreferencesPage) Title() string {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret *C.char               // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_title(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline gets whether an embedded underline in the text of the title
// label indicates a mnemonic. See adw_preferences_page_set_use_underline().
func (self *PreferencesPage) UseUnderline() bool {
	var _arg0 *C.AdwPreferencesPage // out
	var _cret C.gboolean            // in

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_preferences_page_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (self *PreferencesPage) Remove(group *PreferencesGroup) {
	var _arg0 *C.AdwPreferencesPage  // out
	var _arg1 *C.AdwPreferencesGroup // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwPreferencesGroup)(unsafe.Pointer(group.Native()))

	C.adw_preferences_page_remove(_arg0, _arg1)
}

// SetIconName sets the icon name for self.
func (self *PreferencesPage) SetIconName(iconName string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_icon_name(_arg0, _arg1)
}

// SetTitle sets the title of self.
func (self *PreferencesPage) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 *C.char               // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_preferences_page_set_title(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the title label indicates the next
// character should be used for the mnemonic accelerator key.
func (self *PreferencesPage) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwPreferencesPage // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.AdwPreferencesPage)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_preferences_page_set_use_underline(_arg0, _arg1)
}
