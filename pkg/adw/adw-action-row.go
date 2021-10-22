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
		{T: externglib.Type(C.adw_action_row_get_type()), F: marshalActionRower},
	})
}

// ActionRowOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionRowOverrider interface {
	// Activate activates self.
	Activate()
}

// ActionRow: gtk.ListBoxRow used to present actions.
//
// The AdwActionRow widget can have a title, a subtitle and an icon. The row can
// receive additional widgets at its end, or prefix widgets at its start.
//
// It is convenient to present a preference and its related actions.
//
// AdwActionRow is unactivatable by default, giving it an activatable widget
// will automatically make it activatable, but unsetting it won't change the
// row's activatability.
//
//
// AdwActionRow as GtkBuildable
//
// The AdwActionRow implementation of the gtk.Buildable interface supports
// adding a child at its end by specifying “suffix” or omitting the “type”
// attribute of a <child> element.
//
// It also supports adding a child as a prefix widget by specifying “prefix” as
// the “type” attribute of a <child> element.
//
//
// CSS nodes
//
// AdwActionRow has a main CSS node with name row.
//
// It contains the subnode box.header for its main horizontal box, and box.title
// for the vertical box containing the title and subtitle labels.
//
// It contains subnodes label.title and label.subtitle representing respectively
// the title label and subtitle label.
type ActionRow struct {
	PreferencesRow
}

func wrapActionRow(obj *externglib.Object) *ActionRow {
	return &ActionRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
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
				Actionable: gtk.Actionable{
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
				Object: obj,
			},
		},
	}
}

func marshalActionRower(p uintptr) (interface{}, error) {
	return wrapActionRow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewActionRow creates a new AdwActionRow.
func NewActionRow() *ActionRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_action_row_new()

	var _actionRow *ActionRow // out

	_actionRow = wrapActionRow(externglib.Take(unsafe.Pointer(_cret)))

	return _actionRow
}

// Activate activates self.
func (self *ActionRow) Activate() {
	var _arg0 *C.AdwActionRow // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	C.adw_action_row_activate(_arg0)
	runtime.KeepAlive(self)
}

// AddPrefix adds a prefix widget to self.
//
// The function takes the following parameters:
//
//    - widget: widget.
//
func (self *ActionRow) AddPrefix(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_add_prefix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// AddSuffix adds a suffix widget to self.
//
// The function takes the following parameters:
//
//    - widget: widget.
//
func (self *ActionRow) AddSuffix(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_add_suffix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// ActivatableWidget gets the widget activated when self is activated.
func (self *ActionRow) ActivatableWidget() gtk.Widgetter {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_activatable_widget(_arg0)
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

// IconName gets the icon name for self.
func (self *ActionRow) IconName() string {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Subtitle gets the subtitle for self.
func (self *ActionRow) Subtitle() string {
	var _arg0 *C.AdwActionRow // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SubtitleLines gets the number of lines at the end of which the subtitle label
// will be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
func (self *ActionRow) SubtitleLines() int {
	var _arg0 *C.AdwActionRow // out
	var _cret C.int           // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_subtitle_lines(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TitleLines gets the number of lines at the end of which the title label will
// be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
func (self *ActionRow) TitleLines() int {
	var _arg0 *C.AdwActionRow // out
	var _cret C.int           // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_title_lines(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UseUnderline gets whether underlines in title or subtitle are interpreted as
// mnemonics.
func (self *ActionRow) UseUnderline() bool {
	var _arg0 *C.AdwActionRow // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_action_row_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Remove removes a child from self.
//
// The function takes the following parameters:
//
//    - widget: child to be removed.
//
func (self *ActionRow) Remove(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.adw_action_row_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetActivatableWidget sets the widget to activate when self is activated.
//
// The function takes the following parameters:
//
//    - widget: target widget.
//
func (self *ActionRow) SetActivatableWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.adw_action_row_set_activatable_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetIconName sets the icon name for self.
//
// The function takes the following parameters:
//
//    - iconName: icon name.
//
func (self *ActionRow) SetIconName(iconName string) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_action_row_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetSubtitle sets the subtitle for self.
//
// The function takes the following parameters:
//
//    - subtitle: subtitle.
//
func (self *ActionRow) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_action_row_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetSubtitleLines sets the number of lines at the end of which the subtitle
// label will be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
//
// The function takes the following parameters:
//
//    - subtitleLines: number of lines at the end of which the subtitle label
//    will be ellipsized.
//
func (self *ActionRow) SetSubtitleLines(subtitleLines int) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.int           // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(subtitleLines)

	C.adw_action_row_set_subtitle_lines(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitleLines)
}

// SetTitleLines sets the number of lines at the end of which the title label
// will be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
//
// The function takes the following parameters:
//
//    - titleLines: number of lines at the end of which the title label will be
//    ellipsized.
//
func (self *ActionRow) SetTitleLines(titleLines int) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.int           // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(titleLines)

	C.adw_action_row_set_title_lines(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(titleLines)
}

// SetUseUnderline sets whether underlines in title or subtitle are interpreted
// as mnemonics.
//
// The function takes the following parameters:
//
//    - useUnderline: whether underlines are interpreted as mnemonics.
//
func (self *ActionRow) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwActionRow // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwActionRow)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_action_row_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// ConnectActivated: this signal is emitted after the row has been activated.
func (self *ActionRow) ConnectActivated(f func()) externglib.SignalHandle {
	return self.Connect("activated", f)
}
