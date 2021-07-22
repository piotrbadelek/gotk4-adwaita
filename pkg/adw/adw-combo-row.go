// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.adw_combo_row_get_type()), F: marshalComboRower},
	})
}

type ComboRow struct {
	ActionRow
}

func wrapComboRow(obj *externglib.Object) *ComboRow {
	return &ComboRow{
		ActionRow: ActionRow{
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
		},
	}
}

func marshalComboRower(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboRow(obj), nil
}

// NewComboRow creates a new ComboRow.
func NewComboRow() *ComboRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_combo_row_new()

	var _comboRow *ComboRow // out

	_comboRow = wrapComboRow(externglib.Take(unsafe.Pointer(_cret)))

	return _comboRow
}

// Expression gets the expression set with adw_combo_row_set_expression().
func (self *ComboRow) Expression() gtk.Expressioner {
	var _arg0 *C.AdwComboRow   // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_expression(_arg0)

	var _expression gtk.Expressioner // out

	_expression = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Expressioner)

	return _expression
}

// Factory gets the factory that's currently used to populate list items.
//
// The factory returned by this function is always used for the item in the
// button. It is also used for items in the popup if ComboRow:list-factory is
// not set.
func (self *ComboRow) Factory() *gtk.ListItemFactory {
	var _arg0 *C.AdwComboRow        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_factory(_arg0)

	var _listItemFactory *gtk.ListItemFactory // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_listItemFactory = &gtk.ListItemFactory{
			Object: obj,
		}
	}

	return _listItemFactory
}

// ListFactory gets the factory that's currently used to populate list items in
// the popup.
func (self *ComboRow) ListFactory() *gtk.ListItemFactory {
	var _arg0 *C.AdwComboRow        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_list_factory(_arg0)

	var _listItemFactory *gtk.ListItemFactory // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_listItemFactory = &gtk.ListItemFactory{
			Object: obj,
		}
	}

	return _listItemFactory
}

// Model gets the model that provides the displayed items.
func (self *ComboRow) Model() gio.ListModeller {
	var _arg0 *C.AdwComboRow // out
	var _cret *C.GListModel  // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_model(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// Selected gets the position of the selected item.
func (self *ComboRow) Selected() uint {
	var _arg0 *C.AdwComboRow // out
	var _cret C.guint        // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_selected(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedItem gets the selected item. If no item is selected, NULL is
// returned.
func (self *ComboRow) SelectedItem() *externglib.Object {
	var _arg0 *C.AdwComboRow // out
	var _cret C.gpointer     // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_selected_item(_arg0)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// UseSubtitle gets whether the current value of self should be displayed as its
// subtitle.
func (self *ComboRow) UseSubtitle() bool {
	var _arg0 *C.AdwComboRow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))

	_cret = C.adw_combo_row_get_use_subtitle(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpression sets the expression that gets evaluated to obtain strings, used
// to bind strings to labels produced by the default factory.
//
// The expression must have a value type of TYPE_STRING.
func (self *ComboRow) SetExpression(expression gtk.Expressioner) {
	var _arg0 *C.AdwComboRow   // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.adw_combo_row_set_expression(_arg0, _arg1)
}

// SetFactory sets the ListItemFactory to use for populating list items.
func (self *ComboRow) SetFactory(factory *gtk.ListItemFactory) {
	var _arg0 *C.AdwComboRow        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.adw_combo_row_set_factory(_arg0, _arg1)
}

// SetListFactory sets the ListItemFactory to use for populating list items in
// the popup.
func (self *ComboRow) SetListFactory(factory *gtk.ListItemFactory) {
	var _arg0 *C.AdwComboRow        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.adw_combo_row_set_list_factory(_arg0, _arg1)
}

// SetModel sets the Model to use.
func (self *ComboRow) SetModel(model gio.ListModeller) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 *C.GListModel  // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.adw_combo_row_set_model(_arg0, _arg1)
}

// SetSelected selects the item at the given position.
func (self *ComboRow) SetSelected(position uint) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 C.guint        // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.adw_combo_row_set_selected(_arg0, _arg1)
}

// SetUseSubtitle sets whether the current value of self should be displayed as
// its subtitle.
//
// If TRUE, you should not access AdwActionRow:subtitle.
func (self *ComboRow) SetUseSubtitle(useSubtitle bool) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(self.Native()))
	if useSubtitle {
		_arg1 = C.TRUE
	}

	C.adw_combo_row_set_use_subtitle(_arg0, _arg1)
}
