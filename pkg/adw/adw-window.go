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
		{T: externglib.Type(C.adw_window_get_type()), F: marshalWindower},
	})
}

// Window: freeform window.
//
// The AdwWindow widget is a subclass of gtk.Window which has no titlebar area.
// It means gtk.HeaderBar can be used as follows:
//
//    <object class="AdwWindow">
//      <property name="content">
//        <object class="GtkBox">
//          <property name="orientation">vertical</property>
//          <child>
//            <object class="GtkHeaderBar"/>
//          </child>
//          <child>
//            ...
//          </child>
//        </object>
//      </property>
//    </object>
//
//
// Using gtk.Window.GetTitlebar() and gtk.Window.SetTitlebar() is not supported
// and will result in a crash.
type Window struct {
	gtk.Window
}

func wrapWindow(obj *externglib.Object) *Window {
	return &Window{
		Window: gtk.Window{
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
			Root: gtk.Root{
				NativeSurface: gtk.NativeSurface{
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
				},
			},
			ShortcutManager: gtk.ShortcutManager{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalWindower(p uintptr) (interface{}, error) {
	return wrapWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindow creates a new AdwWindow.
func NewWindow() *Window {
	var _cret *C.GtkWidget // in

	_cret = C.adw_window_new()

	var _window *Window // out

	_window = wrapWindow(externglib.Take(unsafe.Pointer(_cret)))

	return _window
}

// Content gets the content widget of self.
//
// This method should always be used instead of gtk.Window.GetChild().
func (self *Window) Content() gtk.Widgetter {
	var _arg0 *C.AdwWindow // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwWindow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_window_get_content(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
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
//    - content widget.
//
func (self *Window) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwWindow // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwWindow)(unsafe.Pointer(self.Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(content.Native()))
	}

	C.adw_window_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}
