// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// glib.Type values for adw-application-window.go.
var GTypeApplicationWindow = externglib.Type(C.adw_application_window_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeApplicationWindow, F: marshalApplicationWindow},
	})
}

// ApplicationWindowOverrider contains methods that are overridable.
type ApplicationWindowOverrider interface {
}

// ApplicationWindow: freeform application window.
//
// <picture> <source srcset="application-window-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="application-window.png"
// alt="application-window"> </picture>
//
// AdwApplicationWindow is a gtk.ApplicationWindow subclass providing the same
// features as window.
//
// See window for details.
//
// Using gtk.Application:menubar is not supported and may result in visual
// glitches.
type ApplicationWindow struct {
	_ [0]func() // equal guard
	gtk.ApplicationWindow
}

var (
	_ externglib.Objector = (*ApplicationWindow)(nil)
	_ gtk.Widgetter       = (*ApplicationWindow)(nil)
)

func classInitApplicationWindower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapApplicationWindow(obj *externglib.Object) *ApplicationWindow {
	return &ApplicationWindow{
		ApplicationWindow: gtk.ApplicationWindow{
			Window: gtk.Window{
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
				Object: obj,
				Root: gtk.Root{
					NativeSurface: gtk.NativeSurface{
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
					},
				},
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
			},
			Object: obj,
			ActionGroup: gio.ActionGroup{
				Object: obj,
			},
			ActionMap: gio.ActionMap{
				Object: obj,
			},
		},
	}
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	return wrapApplicationWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplicationWindow creates a new AdwApplicationWindow for app.
//
// The function takes the following parameters:
//
//    - app: application instance.
//
// The function returns the following values:
//
//    - applicationWindow: newly created AdwApplicationWindow.
//
func NewApplicationWindow(app *gtk.Application) *ApplicationWindow {
	var _arg1 *C.GtkApplication // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(externglib.InternObject(app).Native()))

	_cret = C.adw_application_window_new(_arg1)
	runtime.KeepAlive(app)

	var _applicationWindow *ApplicationWindow // out

	_applicationWindow = wrapApplicationWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _applicationWindow
}

// Content gets the content widget of self.
//
// This method should always be used instead of gtk.Window.GetChild().
//
// The function returns the following values:
//
//    - widget (optional): content widget of self.
//
func (self *ApplicationWindow) Content() gtk.Widgetter {
	var _arg0 *C.AdwApplicationWindow // out
	var _cret *C.GtkWidget            // in

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_application_window_get_content(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetContent sets the content widget of self.
//
// This method should always be used instead of gtk.Window.SetChild().
//
// The function takes the following parameters:
//
//    - content (optional) widget.
//
func (self *ApplicationWindow) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwApplicationWindow // out
	var _arg1 *C.GtkWidget            // out

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(content).Native()))
	}

	C.adw_application_window_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}
