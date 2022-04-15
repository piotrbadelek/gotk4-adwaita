// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// glib.Type values for adw-view-stack.go.
var (
	GTypeViewStack     = externglib.Type(C.adw_view_stack_get_type())
	GTypeViewStackPage = externglib.Type(C.adw_view_stack_page_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeViewStack, F: marshalViewStack},
		{T: GTypeViewStackPage, F: marshalViewStackPage},
	})
}

// ViewStackOverrider contains methods that are overridable.
type ViewStackOverrider interface {
}

// ViewStack: view container for viewswitcher.
//
// AdwViewStack is a container which only shows one page at a time. It is
// typically used to hold an application's main views.
//
// It doesn't provide a way to transition between pages. Instead, a separate
// widget such as viewswitcher can be used with AdwViewStack to provide this
// functionality.
//
// AdwViewStack pages can have a title, an icon, an attention request, and a
// numbered badge that viewswitcher will use to let users identify which page is
// which. Set them using the viewstackpage:title, viewstackpage:icon-name,
// viewstackpage:needs-attention, and viewstackpage:badge-number properties.
//
// Unlike gtk.Stack, transitions between views are not animated.
//
// AdwViewStack maintains a viewstackpage object for each added child, which
// holds additional per-child properties. You obtain the viewstackpage for a
// child with viewstack.GetPage and you can obtain a gtk.SelectionModel
// containing all the pages with viewstack.GetPages.
//
//
// AdwViewStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create viewstackpage objects
// explicitly, and set the child widget as a property on it:
//
//      <object class="AdwViewStack" id="stack">
//        <child>
//          <object class="AdwViewStackPage">
//            <property name="name">overview</property>
//            <property name="title">Overview</property>
//            <property name="child">
//              <object class="AdwStatusPage">
//                <property name="title">Welcome!</property>
//              </object>
//            </property>
//          </object>
//        </child>
//
//
//
// CSS nodes
//
// AdwViewStack has a single CSS node named stack.
type ViewStack struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ViewStack)(nil)
)

func classInitViewStacker(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapViewStack(obj *externglib.Object) *ViewStack {
	return &ViewStack{
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

func marshalViewStack(p uintptr) (interface{}, error) {
	return wrapViewStack(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewStack creates a new AdwViewStack.
//
// The function returns the following values:
//
//    - viewStack: newly created AdwViewStack.
//
func NewViewStack() *ViewStack {
	var _cret *C.GtkWidget // in

	_cret = C.adw_view_stack_new()

	var _viewStack *ViewStack // out

	_viewStack = wrapViewStack(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStack
}

// Add adds a child to self.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
// The function returns the following values:
//
//    - viewStackPage: viewstackpage for child.
//
func (self *ViewStack) Add(child gtk.Widgetter) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	_cret = C.adw_view_stack_add(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// AddNamed adds a child to self.
//
// The child is identified by the name.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name (optional) for child.
//
// The function returns the following values:
//
//    - viewStackPage: AdwViewStackPage for child.
//
func (self *ViewStack) AddNamed(child gtk.Widgetter, name string) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 *C.char             // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	if name != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.adw_view_stack_add_named(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// AddTitled adds a child to self.
//
// The child is identified by the name. The title will be used by viewswitcher
// to represent child, so it should be short.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - name (optional) for child.
//    - title: human-readable title for child.
//
// The function returns the following values:
//
//    - viewStackPage: AdwViewStackPage for child.
//
func (self *ViewStack) AddTitled(child gtk.Widgetter, name, title string) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 *C.char             // out
	var _arg3 *C.char             // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	if name != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.adw_view_stack_add_titled(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
	runtime.KeepAlive(title)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// ChildByName finds the child with name in self.
//
// The function takes the following parameters:
//
//    - name of the child to find.
//
// The function returns the following values:
//
//    - widget (optional): requested child.
//
func (self *ViewStack) ChildByName(name string) gtk.Widgetter {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.char         // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_view_stack_get_child_by_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)

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

// Hhomogeneous gets whether self is horizontally homogeneous.
//
// The function returns the following values:
//
//    - ok: whether self is horizontally homogeneous.
//
func (self *ViewStack) Hhomogeneous() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_get_hhomogeneous(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page gets the viewstackpage object for child.
//
// The function takes the following parameters:
//
//    - child of self.
//
// The function returns the following values:
//
//    - viewStackPage: page object for child.
//
func (self *ViewStack) Page(child gtk.Widgetter) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	_cret = C.adw_view_stack_get_page(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// Pages returns a gio.ListModel that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and change the visible page.
//
// The function returns the following values:
//
//    - selectionModel: GtkSelectionModel for the stack's children.
//
func (self *ViewStack) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwViewStack      // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_get_pages(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel gtk.SelectionModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.SelectionModeller is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gtk.SelectionModeller)
			return ok
		})
		rv, ok := casted.(gtk.SelectionModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.SelectionModeller")
		}
		_selectionModel = rv
	}

	return _selectionModel
}

// Vhomogeneous gets whether self is vertically homogeneous.
//
// The function returns the following values:
//
//    - ok: whether self is vertically homogeneous.
//
func (self *ViewStack) Vhomogeneous() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_get_vhomogeneous(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of self, .
//
// The function returns the following values:
//
//    - widget (optional): visible child.
//
func (self *ViewStack) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwViewStack // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_get_visible_child(_arg0)
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

// VisibleChildName returns the name of the currently visible child of self.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the visible child.
//
func (self *ViewStack) VisibleChildName() string {
	var _arg0 *C.AdwViewStack // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_get_visible_child_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Remove removes a child widget from self.
//
// The function takes the following parameters:
//
//    - child to remove.
//
func (self *ViewStack) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	C.adw_view_stack_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetHhomogeneous sets self to be horizontally homogeneous or not.
//
// The function takes the following parameters:
//
//    - hhomogeneous: whether to make self horizontally homogeneous.
//
func (self *ViewStack) SetHhomogeneous(hhomogeneous bool) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if hhomogeneous {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_set_hhomogeneous(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(hhomogeneous)
}

// SetVhomogeneous sets self to be vertically homogeneous or not.
//
// The function takes the following parameters:
//
//    - vhomogeneous: whether to make self vertically homogeneous.
//
func (self *ViewStack) SetVhomogeneous(vhomogeneous bool) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if vhomogeneous {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_set_vhomogeneous(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(vhomogeneous)
}

// SetVisibleChild makes child the visible child of self.
//
// The function takes the following parameters:
//
//    - child of self.
//
func (self *ViewStack) SetVisibleChild(child gtk.Widgetter) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	C.adw_view_stack_set_visible_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetVisibleChildName makes the child with name visible.
//
// The function takes the following parameters:
//
//    - name of the child.
//
func (self *ViewStack) SetVisibleChildName(name string) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_view_stack_set_visible_child_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// ViewStackPageOverrider contains methods that are overridable.
type ViewStackPageOverrider interface {
}

// ViewStackPage: auxiliary class used by viewstack.
type ViewStackPage struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*ViewStackPage)(nil)
)

func classInitViewStackPager(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapViewStackPage(obj *externglib.Object) *ViewStackPage {
	return &ViewStackPage{
		Object: obj,
	}
}

func marshalViewStackPage(p uintptr) (interface{}, error) {
	return wrapViewStackPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BadgeNumber gets the badge number for this page.
//
// The function returns the following values:
//
//    - guint: badge number for this page.
//
func (self *ViewStackPage) BadgeNumber() uint {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.guint             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_badge_number(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Child gets the stack child to which self belongs.
//
// The function returns the following values:
//
//    - widget: child to which self belongs.
//
func (self *ViewStackPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// IconName gets the icon name of the page.
//
// The function returns the following values:
//
//    - utf8 (optional): icon name of the page.
//
func (self *ViewStackPage) IconName() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of the page.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the page.
//
func (self *ViewStackPage) Name() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NeedsAttention gets whether the page is marked as “needs attention”.
//
// The function returns the following values:
//
//    - ok: whether the page needs attention.
//
func (self *ViewStackPage) NeedsAttention() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_needs_attention(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the page title.
//
// The function returns the following values:
//
//    - utf8 (optional): page title.
//
func (self *ViewStackPage) Title() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline gets whether underlines in the page title indicate mnemonics.
//
// The function returns the following values:
//
//    - ok: whether underlines in the page title indicate mnemonics.
//
func (self *ViewStackPage) UseUnderline() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible gets whether self is visible in its AdwViewStack.
//
// This is independent from the gtk.Widget:visible property of its widget.
//
// The function returns the following values:
//
//    - ok: whether self is visible.
//
func (self *ViewStackPage) Visible() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.adw_view_stack_page_get_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetBadgeNumber sets the badge number for this page.
//
// The function takes the following parameters:
//
//    - badgeNumber: new value to set.
//
func (self *ViewStackPage) SetBadgeNumber(badgeNumber uint) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.guint             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(badgeNumber)

	C.adw_view_stack_page_set_badge_number(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(badgeNumber)
}

// SetIconName sets the icon name of the page.
//
// The function takes the following parameters:
//
//    - iconName (optional): icon name.
//
func (self *ViewStackPage) SetIconName(iconName string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetName sets the name of the page.
//
// The function takes the following parameters:
//
//    - name (optional): page name.
//
func (self *ViewStackPage) SetName(name string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetNeedsAttention sets whether the page is marked as “needs attention”.
//
// The function takes the following parameters:
//
//    - needsAttention: new value to set.
//
func (self *ViewStackPage) SetNeedsAttention(needsAttention bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if needsAttention {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_needs_attention(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(needsAttention)
}

// SetTitle sets the page title.
//
// The function takes the following parameters:
//
//    - title (optional): page title.
//
func (self *ViewStackPage) SetTitle(title string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetUseUnderline sets whether underlines in the page title indicate mnemonics.
//
// The function takes the following parameters:
//
//    - useUnderline: new value to set.
//
func (self *ViewStackPage) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// SetVisible sets whether page is visible in its AdwViewStack.
//
// The function takes the following parameters:
//
//    - visible: whether self is visible.
//
func (self *ViewStackPage) SetVisible(visible bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_visible(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(visible)
}
