// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// glib.Type values for adw-avatar.go.
var GTypeAvatar = externglib.Type(C.adw_avatar_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeAvatar, F: marshalAvatar},
	})
}

// AvatarOverrider contains methods that are overridable.
type AvatarOverrider interface {
}

// Avatar: widget displaying an image, with a generated fallback.
//
// <picture> <source srcset="avatar-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="avatar.png" alt="avatar"> </picture>
//
// AdwAvatar is a widget that shows a round avatar.
//
// AdwAvatar generates an avatar with the initials of the avatar:text on top of
// a colored background.
//
// The color is picked based on the hash of the avatar:text.
//
// If avatar:show-initials is set to FALSE, avatar:icon-name or
// avatar-default-symbolic is shown instead of the initials.
//
// Use avatar:custom-image to set a custom image.
//
//
// CSS nodes
//
// AdwAvatar has a single CSS node with name avatar.
type Avatar struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*Avatar)(nil)
)

func classInitAvatarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapAvatar(obj *externglib.Object) *Avatar {
	return &Avatar{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalAvatar(p uintptr) (interface{}, error) {
	return wrapAvatar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAvatar creates a new AdwAvatar.
//
// The function takes the following parameters:
//
//    - size of the avatar.
//    - text (optional) used to get the initials and color.
//    - showInitials: whether to use initials instead of an icon as fallback.
//
// The function returns the following values:
//
//    - avatar: newly created AdwAvatar.
//
func NewAvatar(size int, text string, showInitials bool) *Avatar {
	var _arg1 C.int        // out
	var _arg2 *C.char      // out
	var _arg3 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.int(size)
	if text != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if showInitials {
		_arg3 = C.TRUE
	}

	_cret = C.adw_avatar_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(size)
	runtime.KeepAlive(text)
	runtime.KeepAlive(showInitials)

	var _avatar *Avatar // out

	_avatar = wrapAvatar(externglib.Take(unsafe.Pointer(_cret)))

	return _avatar
}

// DrawToTexture renders self into a gdk.Texture at scale_factor.
//
// This can be used to export the fallback avatar.
//
// The function takes the following parameters:
//
//    - scaleFactor: scale factor.
//
// The function returns the following values:
//
//    - texture: texture.
//
func (self *Avatar) DrawToTexture(scaleFactor int) gdk.Texturer {
	var _arg0 *C.AdwAvatar  // out
	var _arg1 C.int         // out
	var _cret *C.GdkTexture // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(scaleFactor)

	_cret = C.adw_avatar_draw_to_texture(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(scaleFactor)

	var _texture gdk.Texturer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Texturer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.Texturer)
			return ok
		})
		rv, ok := casted.(gdk.Texturer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Texturer")
		}
		_texture = rv
	}

	return _texture
}

// CustomImage gets the custom image paintable.
//
// The function returns the following values:
//
//    - paintable (optional): custom image.
//
func (self *Avatar) CustomImage() gdk.Paintabler {
	var _arg0 *C.AdwAvatar    // out
	var _cret *C.GdkPaintable // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_avatar_get_custom_image(_arg0)
	runtime.KeepAlive(self)

	var _paintable gdk.Paintabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gdk.Paintabler)
				return ok
			})
			rv, ok := casted.(gdk.Paintabler)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Paintabler")
			}
			_paintable = rv
		}
	}

	return _paintable
}

// IconName gets the name of an icon to use as a fallback.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name.
//
func (self *Avatar) IconName() string {
	var _arg0 *C.AdwAvatar // out
	var _cret *C.char      // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_avatar_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowInitials gets whether initials are used instead of an icon on the
// fallback avatar.
//
// The function returns the following values:
//
//    - ok: whether initials are used instead of an icon as fallback.
//
func (self *Avatar) ShowInitials() bool {
	var _arg0 *C.AdwAvatar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_avatar_get_show_initials(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size gets the size of the avatar.
//
// The function returns the following values:
//
//    - gint: size of the avatar.
//
func (self *Avatar) Size() int {
	var _arg0 *C.AdwAvatar // out
	var _cret C.int        // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_avatar_get_size(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Text gets the text used to generate the fallback initials and color.
//
// The function returns the following values:
//
//    - utf8 (optional): text used to generate the fallback initials and color.
//
func (self *Avatar) Text() string {
	var _arg0 *C.AdwAvatar // out
	var _cret *C.char      // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_avatar_get_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetCustomImage sets the custom image paintable.
//
// The function takes the following parameters:
//
//    - customImage (optional): custom image.
//
func (self *Avatar) SetCustomImage(customImage gdk.Paintabler) {
	var _arg0 *C.AdwAvatar    // out
	var _arg1 *C.GdkPaintable // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if customImage != nil {
		_arg1 = (*C.GdkPaintable)(unsafe.Pointer(externglib.InternObject(customImage).Native()))
	}

	C.adw_avatar_set_custom_image(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(customImage)
}

// SetIconName sets the name of an icon to use as a fallback.
//
// If no name is set, avatar-default-symbolic will be used.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name.
//
func (self *Avatar) SetIconName(iconName string) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 *C.char      // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_avatar_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetShowInitials sets whether to use initials instead of an icon on the
// fallback avatar.
//
// The function takes the following parameters:
//
//    - showInitials: whether to use initials instead of an icon as fallback.
//
func (self *Avatar) SetShowInitials(showInitials bool) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if showInitials {
		_arg1 = C.TRUE
	}

	C.adw_avatar_set_show_initials(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showInitials)
}

// SetSize sets the size of the avatar.
//
// The function takes the following parameters:
//
//    - size of the avatar.
//
func (self *Avatar) SetSize(size int) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 C.int        // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(size)

	C.adw_avatar_set_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(size)
}

// SetText sets the text used to generate the fallback initials and color.
//
// The function takes the following parameters:
//
//    - text (optional) used to get the initials and color.
//
func (self *Avatar) SetText(text string) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 *C.char      // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if text != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_avatar_set_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(text)
}
