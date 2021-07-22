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
		{T: externglib.Type(C.adw_swipe_tracker_get_type()), F: marshalSwipeTrackerer},
	})
}

type SwipeTracker struct {
	*externglib.Object

	gtk.Orientable
}

func wrapSwipeTracker(obj *externglib.Object) *SwipeTracker {
	return &SwipeTracker{
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalSwipeTrackerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSwipeTracker(obj), nil
}

// NewSwipeTracker: create a new SwipeTracker object on widget.
func NewSwipeTracker(swipeable Swipeabler) *SwipeTracker {
	var _arg1 *C.AdwSwipeable    // out
	var _cret *C.AdwSwipeTracker // in

	_arg1 = (*C.AdwSwipeable)(unsafe.Pointer(swipeable.Native()))

	_cret = C.adw_swipe_tracker_new(_arg1)

	var _swipeTracker *SwipeTracker // out

	_swipeTracker = wrapSwipeTracker(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _swipeTracker
}

// AllowLongSwipes: whether to allow swiping for more than one snap point at a
// time. If the value is FALSE, each swipe can only move to the adjacent snap
// points.
func (self *SwipeTracker) AllowLongSwipes() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipe_tracker_get_allow_long_swipes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AllowMouseDrag: get whether self can be dragged with mouse pointer.
func (self *SwipeTracker) AllowMouseDrag() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipe_tracker_get_allow_mouse_drag(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Enabled: get whether self is enabled. When it's not enabled, no events will
// be processed. Generally widgets will want to expose this via a property.
func (self *SwipeTracker) Enabled() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipe_tracker_get_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reversed: get whether self is reversing the swipe direction.
func (self *SwipeTracker) Reversed() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipe_tracker_get_reversed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Swipeable: get self's swipeable widget.
func (self *SwipeTracker) Swipeable() Swipeabler {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret *C.AdwSwipeable    // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))

	_cret = C.adw_swipe_tracker_get_swipeable(_arg0)

	var _swipeable Swipeabler // out

	_swipeable = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Swipeabler)

	return _swipeable
}

// SetAllowLongSwipes sets whether to allow swiping for more than one snap point
// at a time. If the value is FALSE, each swipe can only move to the adjacent
// snap points.
func (self *SwipeTracker) SetAllowLongSwipes(allowLongSwipes bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))
	if allowLongSwipes {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_long_swipes(_arg0, _arg1)
}

// SetAllowMouseDrag: set whether self can be dragged with mouse pointer. This
// should usually be FALSE.
func (self *SwipeTracker) SetAllowMouseDrag(allowMouseDrag bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))
	if allowMouseDrag {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_mouse_drag(_arg0, _arg1)
}

// SetEnabled: set whether self is enabled. When it's not enabled, no events
// will be processed. Usually widgets will want to expose this via a property.
func (self *SwipeTracker) SetEnabled(enabled bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_enabled(_arg0, _arg1)
}

// SetReversed: set whether to reverse the swipe direction. If self is
// horizontal, can be used for supporting RTL text direction.
func (self *SwipeTracker) SetReversed(reversed bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))
	if reversed {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_reversed(_arg0, _arg1)
}

// ShiftPosition: move the current progress value by delta. This can be used to
// adjust the current position if snap points move during the gesture.
func (self *SwipeTracker) ShiftPosition(delta float64) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.double           // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(delta)

	C.adw_swipe_tracker_shift_position(_arg0, _arg1)
}
