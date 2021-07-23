// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
// GdkPixbuf* _gotk4_adw1_AvatarImageLoadFunc(int, gpointer);
// extern void callbackDelete(gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_avatar_get_type()), F: marshalAvatarrer},
	})
}

// AvatarImageLoadFunc: returned Pixbuf is expected to be square with width and
// height set to size. The image is cropped to a circle without any scaling or
// transformation.
type AvatarImageLoadFunc func(size int) (pixbuf *gdkpixbuf.Pixbuf)

//export _gotk4_adw1_AvatarImageLoadFunc
func _gotk4_adw1_AvatarImageLoadFunc(arg0 C.int, arg1 C.gpointer) (cret *C.GdkPixbuf) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var size int // out

	size = int(arg0)

	fn := v.(AvatarImageLoadFunc)
	pixbuf := fn(size)

	if pixbuf != nil {
		cret = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	}

	return cret
}

type Avatar struct {
	gtk.Widget
}

func wrapAvatar(obj *externglib.Object) *Avatar {
	return &Avatar{
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

func marshalAvatarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAvatar(obj), nil
}

// NewAvatar creates a new Avatar.
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

	var _avatar *Avatar // out

	_avatar = wrapAvatar(externglib.Take(unsafe.Pointer(_cret)))

	return _avatar
}

// DrawToPixbuf renders self into a pixbuf at size and scale_factor. This can be
// used to export the fallback avatar.
func (self *Avatar) DrawToPixbuf(size int, scaleFactor int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.AdwAvatar // out
	var _arg1 C.int        // out
	var _arg2 C.int        // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(size)
	_arg2 = C.int(scaleFactor)

	_cret = C.adw_avatar_draw_to_pixbuf(_arg0, _arg1, _arg2)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// IconName gets the name of the icon in the icon theme to use when the icon
// should be displayed.
func (self *Avatar) IconName() string {
	var _arg0 *C.AdwAvatar // out
	var _cret *C.char      // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_avatar_get_icon_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowInitials returns whether initials are used for the fallback or the icon.
func (self *Avatar) ShowInitials() bool {
	var _arg0 *C.AdwAvatar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_avatar_get_show_initials(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size returns the size of the avatar.
func (self *Avatar) Size() int {
	var _arg0 *C.AdwAvatar // out
	var _cret C.int        // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_avatar_get_size(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Text: get the text used to generate the fallback initials and color
func (self *Avatar) Text() string {
	var _arg0 *C.AdwAvatar // out
	var _cret *C.char      // in

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.adw_avatar_get_text(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetIconName sets the name of the icon in the icon theme to use when the icon
// should be displayed. If no name is set, the avatar-default-symbolic icon will
// be used. If the name doesn't match a valid icon, it is an error and no icon
// will be displayed. If the icon theme is changed, the image will be updated
// automatically.
func (self *Avatar) SetIconName(iconName string) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 *C.char      // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_avatar_set_icon_name(_arg0, _arg1)
}

// SetImageLoadFunc: callback which is called when the custom image need to be
// reloaded for some reason (e.g. scale-factor changes).
func (self *Avatar) SetImageLoadFunc(loadImage AvatarImageLoadFunc) {
	var _arg0 *C.AdwAvatar             // out
	var _arg1 C.AdwAvatarImageLoadFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	if loadImage != nil {
		_arg1 = (*[0]byte)(C._gotk4_adw1_AvatarImageLoadFunc)
		_arg2 = C.gpointer(gbox.Assign(loadImage))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.adw_avatar_set_image_load_func(_arg0, _arg1, _arg2, _arg3)
}

// SetShowInitials sets whether the initials should be shown on the fallback
// avatar or the icon.
func (self *Avatar) SetShowInitials(showInitials bool) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	if showInitials {
		_arg1 = C.TRUE
	}

	C.adw_avatar_set_show_initials(_arg0, _arg1)
}

// SetSize sets the size of the avatar.
func (self *Avatar) SetSize(size int) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 C.int        // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(size)

	C.adw_avatar_set_size(_arg0, _arg1)
}

// SetText: set the text used to generate the fallback initials color
func (self *Avatar) SetText(text string) {
	var _arg0 *C.AdwAvatar // out
	var _arg1 *C.char      // out

	_arg0 = (*C.AdwAvatar)(unsafe.Pointer(self.Native()))
	if text != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_avatar_set_text(_arg0, _arg1)
}
