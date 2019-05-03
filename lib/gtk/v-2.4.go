// Code generated - DO NOT EDIT.
// +build gtk_2.4 gtk_2.6 gtk_2.8 gtk_2.10 gtk_2.12 gtk_2.14 gtk_2.16 gtk_2.18 gtk_2.20 gtk_2.22 gtk_2.24 gtk_3.0 gtk_3.2 gtk_3.4 gtk_3.6 gtk_3.8 gtk_3.10 gtk_3.12 gtk_3.14 gtk_3.16 gtk_3.18 gtk_3.20 gtk_3.22 gtk_3.22.6 gtk_3.22.26 gtk_3.22.29

package gtk

import (
	"fmt"
	gdk "github.com/pekim/gobbi/lib/gdk"
	gdkpixbuf "github.com/pekim/gobbi/lib/gdkpixbuf"
	glib "github.com/pekim/gobbi/lib/glib"
	gobject "github.com/pekim/gobbi/lib/gobject"
	pango "github.com/pekim/gobbi/lib/pango"
	"sync"
	"unsafe"
)

// #cgo CFLAGS: -Wno-deprecated-declarations
// #cgo CFLAGS: -Wno-format-security
// #cgo CFLAGS: -Wno-incompatible-pointer-types
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// #include <stdlib.h>
/*

	void accelmap_changedHandler(GObject *, gchar*, guint, guint, gpointer);

	static gulong AccelMap_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(accelmap_changedHandler), data);
	}

*/
/*

	void action_activateHandler(GObject *, gpointer);

	static gulong Action_signal_connect_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "activate", G_CALLBACK(action_activateHandler), data);
	}

*/
/*

	void actiongroup_connectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong ActionGroup_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", G_CALLBACK(actiongroup_connectProxyHandler), data);
	}

*/
/*

	void actiongroup_disconnectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong ActionGroup_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", G_CALLBACK(actiongroup_disconnectProxyHandler), data);
	}

*/
/*

	void actiongroup_postActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong ActionGroup_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", G_CALLBACK(actiongroup_postActivateHandler), data);
	}

*/
/*

	void actiongroup_preActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong ActionGroup_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", G_CALLBACK(actiongroup_preActivateHandler), data);
	}

*/
/*

	void cellrenderer_editingCanceledHandler(GObject *, gpointer);

	static gulong CellRenderer_signal_connect_editing_canceled(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "editing-canceled", G_CALLBACK(cellrenderer_editingCanceledHandler), data);
	}

*/
/*

	void colorbutton_colorSetHandler(GObject *, gpointer);

	static gulong ColorButton_signal_connect_color_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "color-set", G_CALLBACK(colorbutton_colorSetHandler), data);
	}

*/
/*

	void combobox_changedHandler(GObject *, gpointer);

	static gulong ComboBox_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(combobox_changedHandler), data);
	}

*/
/*

	void entrycompletion_actionActivatedHandler(GObject *, gint, gpointer);

	static gulong EntryCompletion_signal_connect_action_activated(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "action-activated", G_CALLBACK(entrycompletion_actionActivatedHandler), data);
	}

*/
/*

	gboolean entrycompletion_matchSelectedHandler(GObject *, GtkTreeModel *, GtkTreeIter *, gpointer);

	static gulong EntryCompletion_signal_connect_match_selected(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "match-selected", G_CALLBACK(entrycompletion_matchSelectedHandler), data);
	}

*/
/*

	void fontbutton_fontSetHandler(GObject *, gpointer);

	static gulong FontButton_signal_connect_font_set(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "font-set", G_CALLBACK(fontbutton_fontSetHandler), data);
	}

*/
/*

	static GtkMessageDialog* _gtk_message_dialog_new_with_markup(GtkWindow* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, const gchar* message_format) {
		return gtk_message_dialog_new_with_markup(parent, flags, type, buttons, message_format);
    }
*/
/*

	void radioaction_changedHandler(GObject *, GtkRadioAction *, gpointer);

	static gulong RadioAction_signal_connect_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "changed", G_CALLBACK(radioaction_changedHandler), data);
	}

*/
/*

	void radiobutton_groupChangedHandler(GObject *, gpointer);

	static gulong RadioButton_signal_connect_group_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "group-changed", G_CALLBACK(radiobutton_groupChangedHandler), data);
	}

*/
/*

	void style_realizeHandler(GObject *, gpointer);

	static gulong Style_signal_connect_realize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "realize", G_CALLBACK(style_realizeHandler), data);
	}

*/
/*

	void style_unrealizeHandler(GObject *, gpointer);

	static gulong Style_signal_connect_unrealize(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "unrealize", G_CALLBACK(style_unrealizeHandler), data);
	}

*/
/*

	void uimanager_actionsChangedHandler(GObject *, gpointer);

	static gulong UIManager_signal_connect_actions_changed(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "actions-changed", G_CALLBACK(uimanager_actionsChangedHandler), data);
	}

*/
/*

	void uimanager_addWidgetHandler(GObject *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_add_widget(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "add-widget", G_CALLBACK(uimanager_addWidgetHandler), data);
	}

*/
/*

	void uimanager_connectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_connect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "connect-proxy", G_CALLBACK(uimanager_connectProxyHandler), data);
	}

*/
/*

	void uimanager_disconnectProxyHandler(GObject *, GtkAction *, GtkWidget *, gpointer);

	static gulong UIManager_signal_connect_disconnect_proxy(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "disconnect-proxy", G_CALLBACK(uimanager_disconnectProxyHandler), data);
	}

*/
/*

	void uimanager_postActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong UIManager_signal_connect_post_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "post-activate", G_CALLBACK(uimanager_postActivateHandler), data);
	}

*/
/*

	void uimanager_preActivateHandler(GObject *, GtkAction *, gpointer);

	static gulong UIManager_signal_connect_pre_activate(gpointer instance, gpointer data) {
		return g_signal_connect(instance, "pre-activate", G_CALLBACK(uimanager_preActivateHandler), data);
	}

*/
import "C"

type signalAccelMapChangedDetail struct {
	callback  AccelMapSignalChangedCallback
	handlerID C.gulong
}

var signalAccelMapChangedId int
var signalAccelMapChangedMap = make(map[int]signalAccelMapChangedDetail)
var signalAccelMapChangedLock sync.RWMutex

// AccelMapSignalChangedCallback is a callback function for a 'changed' signal emitted from a AccelMap.
type AccelMapSignalChangedCallback func(accelPath string, accelKey uint32, accelMods gdk.ModifierType)

/*
ConnectChanged connects the callback to the 'changed' signal for the AccelMap.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *AccelMap) ConnectChanged(callback AccelMapSignalChangedCallback) int {
	signalAccelMapChangedLock.Lock()
	defer signalAccelMapChangedLock.Unlock()

	signalAccelMapChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.AccelMap_signal_connect_changed(instance, C.gpointer(uintptr(signalAccelMapChangedId)))

	detail := signalAccelMapChangedDetail{callback, handlerID}
	signalAccelMapChangedMap[signalAccelMapChangedId] = detail

	return signalAccelMapChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the AccelMap.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *AccelMap) DisconnectChanged(connectionID int) {
	signalAccelMapChangedLock.Lock()
	defer signalAccelMapChangedLock.Unlock()

	detail, exists := signalAccelMapChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalAccelMapChangedMap, connectionID)
}

//export accelmap_changedHandler
func accelmap_changedHandler(_ *C.GObject, c_accel_path *C.gchar, c_accel_key C.guint, c_accel_mods C.guint, data C.gpointer) {
	signalAccelMapChangedLock.RLock()
	defer signalAccelMapChangedLock.RUnlock()

	accelPath := C.GoString(c_accel_path)

	accelKey := uint32(c_accel_key)

	accelMods := gdk.ModifierType(c_accel_mods)

	index := int(uintptr(data))
	callback := signalAccelMapChangedMap[index].callback
	callback(accelPath, accelKey, accelMods)
}

// AccelMapGet is a wrapper around the C function gtk_accel_map_get.
func AccelMapGet() *AccelMap {
	retC := C.gtk_accel_map_get()
	retGo := AccelMapNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AccelMapLockPath is a wrapper around the C function gtk_accel_map_lock_path.
func AccelMapLockPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_accel_map_lock_path(c_accel_path)

	return
}

// AccelMapUnlockPath is a wrapper around the C function gtk_accel_map_unlock_path.
func AccelMapUnlockPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_accel_map_unlock_path(c_accel_path)

	return
}

type signalActionActivateDetail struct {
	callback  ActionSignalActivateCallback
	handlerID C.gulong
}

var signalActionActivateId int
var signalActionActivateMap = make(map[int]signalActionActivateDetail)
var signalActionActivateLock sync.RWMutex

// ActionSignalActivateCallback is a callback function for a 'activate' signal emitted from a Action.
type ActionSignalActivateCallback func()

/*
ConnectActivate connects the callback to the 'activate' signal for the Action.

The returned value represents the connection, and may be passed to DisconnectActivate to remove it.
*/
func (recv *Action) ConnectActivate(callback ActionSignalActivateCallback) int {
	signalActionActivateLock.Lock()
	defer signalActionActivateLock.Unlock()

	signalActionActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.Action_signal_connect_activate(instance, C.gpointer(uintptr(signalActionActivateId)))

	detail := signalActionActivateDetail{callback, handlerID}
	signalActionActivateMap[signalActionActivateId] = detail

	return signalActionActivateId
}

/*
DisconnectActivate disconnects a callback from the 'activate' signal for the Action.

The connectionID should be a value returned from a call to ConnectActivate.
*/
func (recv *Action) DisconnectActivate(connectionID int) {
	signalActionActivateLock.Lock()
	defer signalActionActivateLock.Unlock()

	detail, exists := signalActionActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionActivateMap, connectionID)
}

//export action_activateHandler
func action_activateHandler(_ *C.GObject, data C.gpointer) {
	signalActionActivateLock.RLock()
	defer signalActionActivateLock.RUnlock()

	index := int(uintptr(data))
	callback := signalActionActivateMap[index].callback
	callback()
}

// ActionNew is a wrapper around the C function gtk_action_new.
func ActionNew(name string, label string, tooltip string, stockId string) *Action {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Activate is a wrapper around the C function gtk_action_activate.
func (recv *Action) Activate() {
	C.gtk_action_activate((*C.GtkAction)(recv.native))

	return
}

// ConnectAccelerator is a wrapper around the C function gtk_action_connect_accelerator.
func (recv *Action) ConnectAccelerator() {
	C.gtk_action_connect_accelerator((*C.GtkAction)(recv.native))

	return
}

// CreateIcon is a wrapper around the C function gtk_action_create_icon.
func (recv *Action) CreateIcon(iconSize IconSize) *Widget {
	c_icon_size := (C.GtkIconSize)(iconSize)

	retC := C.gtk_action_create_icon((*C.GtkAction)(recv.native), c_icon_size)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateMenuItem is a wrapper around the C function gtk_action_create_menu_item.
func (recv *Action) CreateMenuItem() *Widget {
	retC := C.gtk_action_create_menu_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// CreateToolItem is a wrapper around the C function gtk_action_create_tool_item.
func (recv *Action) CreateToolItem() *Widget {
	retC := C.gtk_action_create_tool_item((*C.GtkAction)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// DisconnectAccelerator is a wrapper around the C function gtk_action_disconnect_accelerator.
func (recv *Action) DisconnectAccelerator() {
	C.gtk_action_disconnect_accelerator((*C.GtkAction)(recv.native))

	return
}

// GetName is a wrapper around the C function gtk_action_get_name.
func (recv *Action) GetName() string {
	retC := C.gtk_action_get_name((*C.GtkAction)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetProxies is a wrapper around the C function gtk_action_get_proxies.
func (recv *Action) GetProxies() *glib.SList {
	retC := C.gtk_action_get_proxies((*C.GtkAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_action_get_sensitive.
func (recv *Action) GetSensitive() bool {
	retC := C.gtk_action_get_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_action_get_visible.
func (recv *Action) GetVisible() bool {
	retC := C.gtk_action_get_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsSensitive is a wrapper around the C function gtk_action_is_sensitive.
func (recv *Action) IsSensitive() bool {
	retC := C.gtk_action_is_sensitive((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsVisible is a wrapper around the C function gtk_action_is_visible.
func (recv *Action) IsVisible() bool {
	retC := C.gtk_action_is_visible((*C.GtkAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAccelGroup is a wrapper around the C function gtk_action_set_accel_group.
func (recv *Action) SetAccelGroup(accelGroup *AccelGroup) {
	c_accel_group := (*C.GtkAccelGroup)(C.NULL)
	if accelGroup != nil {
		c_accel_group = (*C.GtkAccelGroup)(accelGroup.ToC())
	}

	C.gtk_action_set_accel_group((*C.GtkAction)(recv.native), c_accel_group)

	return
}

// SetAccelPath is a wrapper around the C function gtk_action_set_accel_path.
func (recv *Action) SetAccelPath(accelPath string) {
	c_accel_path := C.CString(accelPath)
	defer C.free(unsafe.Pointer(c_accel_path))

	C.gtk_action_set_accel_path((*C.GtkAction)(recv.native), c_accel_path)

	return
}

type signalActionGroupConnectProxyDetail struct {
	callback  ActionGroupSignalConnectProxyCallback
	handlerID C.gulong
}

var signalActionGroupConnectProxyId int
var signalActionGroupConnectProxyMap = make(map[int]signalActionGroupConnectProxyDetail)
var signalActionGroupConnectProxyLock sync.RWMutex

// ActionGroupSignalConnectProxyCallback is a callback function for a 'connect-proxy' signal emitted from a ActionGroup.
type ActionGroupSignalConnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectConnectProxy connects the callback to the 'connect-proxy' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectConnectProxy to remove it.
*/
func (recv *ActionGroup) ConnectConnectProxy(callback ActionGroupSignalConnectProxyCallback) int {
	signalActionGroupConnectProxyLock.Lock()
	defer signalActionGroupConnectProxyLock.Unlock()

	signalActionGroupConnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalActionGroupConnectProxyId)))

	detail := signalActionGroupConnectProxyDetail{callback, handlerID}
	signalActionGroupConnectProxyMap[signalActionGroupConnectProxyId] = detail

	return signalActionGroupConnectProxyId
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *ActionGroup) DisconnectConnectProxy(connectionID int) {
	signalActionGroupConnectProxyLock.Lock()
	defer signalActionGroupConnectProxyLock.Unlock()

	detail, exists := signalActionGroupConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupConnectProxyMap, connectionID)
}

//export actiongroup_connectProxyHandler
func actiongroup_connectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalActionGroupConnectProxyLock.RLock()
	defer signalActionGroupConnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalActionGroupConnectProxyMap[index].callback
	callback(action, proxy)
}

type signalActionGroupDisconnectProxyDetail struct {
	callback  ActionGroupSignalDisconnectProxyCallback
	handlerID C.gulong
}

var signalActionGroupDisconnectProxyId int
var signalActionGroupDisconnectProxyMap = make(map[int]signalActionGroupDisconnectProxyDetail)
var signalActionGroupDisconnectProxyLock sync.RWMutex

// ActionGroupSignalDisconnectProxyCallback is a callback function for a 'disconnect-proxy' signal emitted from a ActionGroup.
type ActionGroupSignalDisconnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectDisconnectProxy connects the callback to the 'disconnect-proxy' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectDisconnectProxy to remove it.
*/
func (recv *ActionGroup) ConnectDisconnectProxy(callback ActionGroupSignalDisconnectProxyCallback) int {
	signalActionGroupDisconnectProxyLock.Lock()
	defer signalActionGroupDisconnectProxyLock.Unlock()

	signalActionGroupDisconnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalActionGroupDisconnectProxyId)))

	detail := signalActionGroupDisconnectProxyDetail{callback, handlerID}
	signalActionGroupDisconnectProxyMap[signalActionGroupDisconnectProxyId] = detail

	return signalActionGroupDisconnectProxyId
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *ActionGroup) DisconnectDisconnectProxy(connectionID int) {
	signalActionGroupDisconnectProxyLock.Lock()
	defer signalActionGroupDisconnectProxyLock.Unlock()

	detail, exists := signalActionGroupDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupDisconnectProxyMap, connectionID)
}

//export actiongroup_disconnectProxyHandler
func actiongroup_disconnectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalActionGroupDisconnectProxyLock.RLock()
	defer signalActionGroupDisconnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalActionGroupDisconnectProxyMap[index].callback
	callback(action, proxy)
}

type signalActionGroupPostActivateDetail struct {
	callback  ActionGroupSignalPostActivateCallback
	handlerID C.gulong
}

var signalActionGroupPostActivateId int
var signalActionGroupPostActivateMap = make(map[int]signalActionGroupPostActivateDetail)
var signalActionGroupPostActivateLock sync.RWMutex

// ActionGroupSignalPostActivateCallback is a callback function for a 'post-activate' signal emitted from a ActionGroup.
type ActionGroupSignalPostActivateCallback func(action *Action)

/*
ConnectPostActivate connects the callback to the 'post-activate' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectPostActivate to remove it.
*/
func (recv *ActionGroup) ConnectPostActivate(callback ActionGroupSignalPostActivateCallback) int {
	signalActionGroupPostActivateLock.Lock()
	defer signalActionGroupPostActivateLock.Unlock()

	signalActionGroupPostActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_post_activate(instance, C.gpointer(uintptr(signalActionGroupPostActivateId)))

	detail := signalActionGroupPostActivateDetail{callback, handlerID}
	signalActionGroupPostActivateMap[signalActionGroupPostActivateId] = detail

	return signalActionGroupPostActivateId
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *ActionGroup) DisconnectPostActivate(connectionID int) {
	signalActionGroupPostActivateLock.Lock()
	defer signalActionGroupPostActivateLock.Unlock()

	detail, exists := signalActionGroupPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupPostActivateMap, connectionID)
}

//export actiongroup_postActivateHandler
func actiongroup_postActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalActionGroupPostActivateLock.RLock()
	defer signalActionGroupPostActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalActionGroupPostActivateMap[index].callback
	callback(action)
}

type signalActionGroupPreActivateDetail struct {
	callback  ActionGroupSignalPreActivateCallback
	handlerID C.gulong
}

var signalActionGroupPreActivateId int
var signalActionGroupPreActivateMap = make(map[int]signalActionGroupPreActivateDetail)
var signalActionGroupPreActivateLock sync.RWMutex

// ActionGroupSignalPreActivateCallback is a callback function for a 'pre-activate' signal emitted from a ActionGroup.
type ActionGroupSignalPreActivateCallback func(action *Action)

/*
ConnectPreActivate connects the callback to the 'pre-activate' signal for the ActionGroup.

The returned value represents the connection, and may be passed to DisconnectPreActivate to remove it.
*/
func (recv *ActionGroup) ConnectPreActivate(callback ActionGroupSignalPreActivateCallback) int {
	signalActionGroupPreActivateLock.Lock()
	defer signalActionGroupPreActivateLock.Unlock()

	signalActionGroupPreActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.ActionGroup_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalActionGroupPreActivateId)))

	detail := signalActionGroupPreActivateDetail{callback, handlerID}
	signalActionGroupPreActivateMap[signalActionGroupPreActivateId] = detail

	return signalActionGroupPreActivateId
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the ActionGroup.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *ActionGroup) DisconnectPreActivate(connectionID int) {
	signalActionGroupPreActivateLock.Lock()
	defer signalActionGroupPreActivateLock.Unlock()

	detail, exists := signalActionGroupPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalActionGroupPreActivateMap, connectionID)
}

//export actiongroup_preActivateHandler
func actiongroup_preActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalActionGroupPreActivateLock.RLock()
	defer signalActionGroupPreActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalActionGroupPreActivateMap[index].callback
	callback(action)
}

// ActionGroupNew is a wrapper around the C function gtk_action_group_new.
func ActionGroupNew(name string) *ActionGroup {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	retC := C.gtk_action_group_new(c_name)
	retGo := ActionGroupNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddAction is a wrapper around the C function gtk_action_group_add_action.
func (recv *ActionGroup) AddAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_action_group_add_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// AddActionWithAccel is a wrapper around the C function gtk_action_group_add_action_with_accel.
func (recv *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	c_accelerator := C.CString(accelerator)
	defer C.free(unsafe.Pointer(c_accelerator))

	C.gtk_action_group_add_action_with_accel((*C.GtkActionGroup)(recv.native), c_action, c_accelerator)

	return
}

// Unsupported : gtk_action_group_add_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_radio_actions_full : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions : unsupported parameter entries :

// Unsupported : gtk_action_group_add_toggle_actions_full : unsupported parameter entries :

// GetAction is a wrapper around the C function gtk_action_group_get_action.
func (recv *ActionGroup) GetAction(actionName string) *Action {
	c_action_name := C.CString(actionName)
	defer C.free(unsafe.Pointer(c_action_name))

	retC := C.gtk_action_group_get_action((*C.GtkActionGroup)(recv.native), c_action_name)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetName is a wrapper around the C function gtk_action_group_get_name.
func (recv *ActionGroup) GetName() string {
	retC := C.gtk_action_group_get_name((*C.GtkActionGroup)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetSensitive is a wrapper around the C function gtk_action_group_get_sensitive.
func (recv *ActionGroup) GetSensitive() bool {
	retC := C.gtk_action_group_get_sensitive((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisible is a wrapper around the C function gtk_action_group_get_visible.
func (recv *ActionGroup) GetVisible() bool {
	retC := C.gtk_action_group_get_visible((*C.GtkActionGroup)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListActions is a wrapper around the C function gtk_action_group_list_actions.
func (recv *ActionGroup) ListActions() *glib.List {
	retC := C.gtk_action_group_list_actions((*C.GtkActionGroup)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RemoveAction is a wrapper around the C function gtk_action_group_remove_action.
func (recv *ActionGroup) RemoveAction(action *Action) {
	c_action := (*C.GtkAction)(C.NULL)
	if action != nil {
		c_action = (*C.GtkAction)(action.ToC())
	}

	C.gtk_action_group_remove_action((*C.GtkActionGroup)(recv.native), c_action)

	return
}

// SetSensitive is a wrapper around the C function gtk_action_group_set_sensitive.
func (recv *ActionGroup) SetSensitive(sensitive bool) {
	c_sensitive :=
		boolToGboolean(sensitive)

	C.gtk_action_group_set_sensitive((*C.GtkActionGroup)(recv.native), c_sensitive)

	return
}

// Unsupported : gtk_action_group_set_translate_func : unsupported parameter func : no type generator for TranslateFunc (GtkTranslateFunc) for param func

// SetTranslationDomain is a wrapper around the C function gtk_action_group_set_translation_domain.
func (recv *ActionGroup) SetTranslationDomain(domain string) {
	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	C.gtk_action_group_set_translation_domain((*C.GtkActionGroup)(recv.native), c_domain)

	return
}

// SetVisible is a wrapper around the C function gtk_action_group_set_visible.
func (recv *ActionGroup) SetVisible(visible bool) {
	c_visible :=
		boolToGboolean(visible)

	C.gtk_action_group_set_visible((*C.GtkActionGroup)(recv.native), c_visible)

	return
}

// GetPadding is a wrapper around the C function gtk_alignment_get_padding.
func (recv *Alignment) GetPadding() (uint32, uint32, uint32, uint32) {
	var c_padding_top C.guint

	var c_padding_bottom C.guint

	var c_padding_left C.guint

	var c_padding_right C.guint

	C.gtk_alignment_get_padding((*C.GtkAlignment)(recv.native), &c_padding_top, &c_padding_bottom, &c_padding_left, &c_padding_right)

	paddingTop := (uint32)(c_padding_top)

	paddingBottom := (uint32)(c_padding_bottom)

	paddingLeft := (uint32)(c_padding_left)

	paddingRight := (uint32)(c_padding_right)

	return paddingTop, paddingBottom, paddingLeft, paddingRight
}

// SetPadding is a wrapper around the C function gtk_alignment_set_padding.
func (recv *Alignment) SetPadding(paddingTop uint32, paddingBottom uint32, paddingLeft uint32, paddingRight uint32) {
	c_padding_top := (C.guint)(paddingTop)

	c_padding_bottom := (C.guint)(paddingBottom)

	c_padding_left := (C.guint)(paddingLeft)

	c_padding_right := (C.guint)(paddingRight)

	C.gtk_alignment_set_padding((*C.GtkAlignment)(recv.native), c_padding_top, c_padding_bottom, c_padding_left, c_padding_right)

	return
}

// GetAlignment is a wrapper around the C function gtk_button_get_alignment.
func (recv *Button) GetAlignment() (float32, float32) {
	var c_xalign C.gfloat

	var c_yalign C.gfloat

	C.gtk_button_get_alignment((*C.GtkButton)(recv.native), &c_xalign, &c_yalign)

	xalign := (float32)(c_xalign)

	yalign := (float32)(c_yalign)

	return xalign, yalign
}

// GetFocusOnClick is a wrapper around the C function gtk_button_get_focus_on_click.
func (recv *Button) GetFocusOnClick() bool {
	retC := C.gtk_button_get_focus_on_click((*C.GtkButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_button_set_alignment.
func (recv *Button) SetAlignment(xalign float32, yalign float32) {
	c_xalign := (C.gfloat)(xalign)

	c_yalign := (C.gfloat)(yalign)

	C.gtk_button_set_alignment((*C.GtkButton)(recv.native), c_xalign, c_yalign)

	return
}

// SetFocusOnClick is a wrapper around the C function gtk_button_set_focus_on_click.
func (recv *Button) SetFocusOnClick(focusOnClick bool) {
	c_focus_on_click :=
		boolToGboolean(focusOnClick)

	C.gtk_button_set_focus_on_click((*C.GtkButton)(recv.native), c_focus_on_click)

	return
}

// GetChildSecondary is a wrapper around the C function gtk_button_box_get_child_secondary.
func (recv *ButtonBox) GetChildSecondary(child *Widget) bool {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	retC := C.gtk_button_box_get_child_secondary((*C.GtkButtonBox)(recv.native), c_child)
	retGo := retC == C.TRUE

	return retGo
}

// GetDisplayOptions is a wrapper around the C function gtk_calendar_get_display_options.
func (recv *Calendar) GetDisplayOptions() CalendarDisplayOptions {
	retC := C.gtk_calendar_get_display_options((*C.GtkCalendar)(recv.native))
	retGo := (CalendarDisplayOptions)(retC)

	return retGo
}

// SetDisplayOptions is a wrapper around the C function gtk_calendar_set_display_options.
func (recv *Calendar) SetDisplayOptions(flags CalendarDisplayOptions) {
	c_flags := (C.GtkCalendarDisplayOptions)(flags)

	C.gtk_calendar_set_display_options((*C.GtkCalendar)(recv.native), c_flags)

	return
}

type signalCellRendererEditingCanceledDetail struct {
	callback  CellRendererSignalEditingCanceledCallback
	handlerID C.gulong
}

var signalCellRendererEditingCanceledId int
var signalCellRendererEditingCanceledMap = make(map[int]signalCellRendererEditingCanceledDetail)
var signalCellRendererEditingCanceledLock sync.RWMutex

// CellRendererSignalEditingCanceledCallback is a callback function for a 'editing-canceled' signal emitted from a CellRenderer.
type CellRendererSignalEditingCanceledCallback func()

/*
ConnectEditingCanceled connects the callback to the 'editing-canceled' signal for the CellRenderer.

The returned value represents the connection, and may be passed to DisconnectEditingCanceled to remove it.
*/
func (recv *CellRenderer) ConnectEditingCanceled(callback CellRendererSignalEditingCanceledCallback) int {
	signalCellRendererEditingCanceledLock.Lock()
	defer signalCellRendererEditingCanceledLock.Unlock()

	signalCellRendererEditingCanceledId++
	instance := C.gpointer(recv.native)
	handlerID := C.CellRenderer_signal_connect_editing_canceled(instance, C.gpointer(uintptr(signalCellRendererEditingCanceledId)))

	detail := signalCellRendererEditingCanceledDetail{callback, handlerID}
	signalCellRendererEditingCanceledMap[signalCellRendererEditingCanceledId] = detail

	return signalCellRendererEditingCanceledId
}

/*
DisconnectEditingCanceled disconnects a callback from the 'editing-canceled' signal for the CellRenderer.

The connectionID should be a value returned from a call to ConnectEditingCanceled.
*/
func (recv *CellRenderer) DisconnectEditingCanceled(connectionID int) {
	signalCellRendererEditingCanceledLock.Lock()
	defer signalCellRendererEditingCanceledLock.Unlock()

	detail, exists := signalCellRendererEditingCanceledMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalCellRendererEditingCanceledMap, connectionID)
}

//export cellrenderer_editingCanceledHandler
func cellrenderer_editingCanceledHandler(_ *C.GObject, data C.gpointer) {
	signalCellRendererEditingCanceledLock.RLock()
	defer signalCellRendererEditingCanceledLock.RUnlock()

	index := int(uintptr(data))
	callback := signalCellRendererEditingCanceledMap[index].callback
	callback()
}

// GetDrawAsRadio is a wrapper around the C function gtk_check_menu_item_get_draw_as_radio.
func (recv *CheckMenuItem) GetDrawAsRadio() bool {
	retC := C.gtk_check_menu_item_get_draw_as_radio((*C.GtkCheckMenuItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDrawAsRadio is a wrapper around the C function gtk_check_menu_item_set_draw_as_radio.
func (recv *CheckMenuItem) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_check_menu_item_set_draw_as_radio((*C.GtkCheckMenuItem)(recv.native), c_draw_as_radio)

	return
}

// Unsupported : gtk_clipboard_request_targets : unsupported parameter callback : no type generator for ClipboardTargetsReceivedFunc (GtkClipboardTargetsReceivedFunc) for param callback

// Unsupported : gtk_clipboard_wait_for_targets : unsupported parameter targets : output array param targets

type signalColorButtonColorSetDetail struct {
	callback  ColorButtonSignalColorSetCallback
	handlerID C.gulong
}

var signalColorButtonColorSetId int
var signalColorButtonColorSetMap = make(map[int]signalColorButtonColorSetDetail)
var signalColorButtonColorSetLock sync.RWMutex

// ColorButtonSignalColorSetCallback is a callback function for a 'color-set' signal emitted from a ColorButton.
type ColorButtonSignalColorSetCallback func()

/*
ConnectColorSet connects the callback to the 'color-set' signal for the ColorButton.

The returned value represents the connection, and may be passed to DisconnectColorSet to remove it.
*/
func (recv *ColorButton) ConnectColorSet(callback ColorButtonSignalColorSetCallback) int {
	signalColorButtonColorSetLock.Lock()
	defer signalColorButtonColorSetLock.Unlock()

	signalColorButtonColorSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.ColorButton_signal_connect_color_set(instance, C.gpointer(uintptr(signalColorButtonColorSetId)))

	detail := signalColorButtonColorSetDetail{callback, handlerID}
	signalColorButtonColorSetMap[signalColorButtonColorSetId] = detail

	return signalColorButtonColorSetId
}

/*
DisconnectColorSet disconnects a callback from the 'color-set' signal for the ColorButton.

The connectionID should be a value returned from a call to ConnectColorSet.
*/
func (recv *ColorButton) DisconnectColorSet(connectionID int) {
	signalColorButtonColorSetLock.Lock()
	defer signalColorButtonColorSetLock.Unlock()

	detail, exists := signalColorButtonColorSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalColorButtonColorSetMap, connectionID)
}

//export colorbutton_colorSetHandler
func colorbutton_colorSetHandler(_ *C.GObject, data C.gpointer) {
	signalColorButtonColorSetLock.RLock()
	defer signalColorButtonColorSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalColorButtonColorSetMap[index].callback
	callback()
}

// ColorButtonNew is a wrapper around the C function gtk_color_button_new.
func ColorButtonNew() *ColorButton {
	retC := C.gtk_color_button_new()
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ColorButtonNewWithColor is a wrapper around the C function gtk_color_button_new_with_color.
func ColorButtonNewWithColor(color *gdk.Color) *ColorButton {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	retC := C.gtk_color_button_new_with_color(c_color)
	retGo := ColorButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAlpha is a wrapper around the C function gtk_color_button_get_alpha.
func (recv *ColorButton) GetAlpha() uint16 {
	retC := C.gtk_color_button_get_alpha((*C.GtkColorButton)(recv.native))
	retGo := (uint16)(retC)

	return retGo
}

// GetColor is a wrapper around the C function gtk_color_button_get_color.
func (recv *ColorButton) GetColor() *gdk.Color {
	var c_color C.GdkColor

	C.gtk_color_button_get_color((*C.GtkColorButton)(recv.native), &c_color)

	color := gdk.ColorNewFromC(unsafe.Pointer(&c_color))

	return color
}

// GetTitle is a wrapper around the C function gtk_color_button_get_title.
func (recv *ColorButton) GetTitle() string {
	retC := C.gtk_color_button_get_title((*C.GtkColorButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseAlpha is a wrapper around the C function gtk_color_button_get_use_alpha.
func (recv *ColorButton) GetUseAlpha() bool {
	retC := C.gtk_color_button_get_use_alpha((*C.GtkColorButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAlpha is a wrapper around the C function gtk_color_button_set_alpha.
func (recv *ColorButton) SetAlpha(alpha uint16) {
	c_alpha := (C.guint16)(alpha)

	C.gtk_color_button_set_alpha((*C.GtkColorButton)(recv.native), c_alpha)

	return
}

// SetColor is a wrapper around the C function gtk_color_button_set_color.
func (recv *ColorButton) SetColor(color *gdk.Color) {
	c_color := (*C.GdkColor)(C.NULL)
	if color != nil {
		c_color = (*C.GdkColor)(color.ToC())
	}

	C.gtk_color_button_set_color((*C.GtkColorButton)(recv.native), c_color)

	return
}

// SetTitle is a wrapper around the C function gtk_color_button_set_title.
func (recv *ColorButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_color_button_set_title((*C.GtkColorButton)(recv.native), c_title)

	return
}

// SetUseAlpha is a wrapper around the C function gtk_color_button_set_use_alpha.
func (recv *ColorButton) SetUseAlpha(useAlpha bool) {
	c_use_alpha :=
		boolToGboolean(useAlpha)

	C.gtk_color_button_set_use_alpha((*C.GtkColorButton)(recv.native), c_use_alpha)

	return
}

type signalComboBoxChangedDetail struct {
	callback  ComboBoxSignalChangedCallback
	handlerID C.gulong
}

var signalComboBoxChangedId int
var signalComboBoxChangedMap = make(map[int]signalComboBoxChangedDetail)
var signalComboBoxChangedLock sync.RWMutex

// ComboBoxSignalChangedCallback is a callback function for a 'changed' signal emitted from a ComboBox.
type ComboBoxSignalChangedCallback func()

/*
ConnectChanged connects the callback to the 'changed' signal for the ComboBox.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *ComboBox) ConnectChanged(callback ComboBoxSignalChangedCallback) int {
	signalComboBoxChangedLock.Lock()
	defer signalComboBoxChangedLock.Unlock()

	signalComboBoxChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.ComboBox_signal_connect_changed(instance, C.gpointer(uintptr(signalComboBoxChangedId)))

	detail := signalComboBoxChangedDetail{callback, handlerID}
	signalComboBoxChangedMap[signalComboBoxChangedId] = detail

	return signalComboBoxChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the ComboBox.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *ComboBox) DisconnectChanged(connectionID int) {
	signalComboBoxChangedLock.Lock()
	defer signalComboBoxChangedLock.Unlock()

	detail, exists := signalComboBoxChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalComboBoxChangedMap, connectionID)
}

//export combobox_changedHandler
func combobox_changedHandler(_ *C.GObject, data C.gpointer) {
	signalComboBoxChangedLock.RLock()
	defer signalComboBoxChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalComboBoxChangedMap[index].callback
	callback()
}

// ComboBoxNew is a wrapper around the C function gtk_combo_box_new.
func ComboBoxNew() *ComboBox {
	retC := C.gtk_combo_box_new()
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ComboBoxNewWithModel is a wrapper around the C function gtk_combo_box_new_with_model.
func ComboBoxNewWithModel(model *TreeModel) *ComboBox {
	c_model := (*C.GtkTreeModel)(model.ToC())

	retC := C.gtk_combo_box_new_with_model(c_model)
	retGo := ComboBoxNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_combo_box_get_active.
func (recv *ComboBox) GetActive() int32 {
	retC := C.gtk_combo_box_get_active((*C.GtkComboBox)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetActiveIter is a wrapper around the C function gtk_combo_box_get_active_iter.
func (recv *ComboBox) GetActiveIter() (bool, *TreeIter) {
	var c_iter C.GtkTreeIter

	retC := C.gtk_combo_box_get_active_iter((*C.GtkComboBox)(recv.native), &c_iter)
	retGo := retC == C.TRUE

	iter := TreeIterNewFromC(unsafe.Pointer(&c_iter))

	return retGo, iter
}

// GetModel is a wrapper around the C function gtk_combo_box_get_model.
func (recv *ComboBox) GetModel() *TreeModel {
	retC := C.gtk_combo_box_get_model((*C.GtkComboBox)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Popdown is a wrapper around the C function gtk_combo_box_popdown.
func (recv *ComboBox) Popdown() {
	C.gtk_combo_box_popdown((*C.GtkComboBox)(recv.native))

	return
}

// Popup is a wrapper around the C function gtk_combo_box_popup.
func (recv *ComboBox) Popup() {
	C.gtk_combo_box_popup((*C.GtkComboBox)(recv.native))

	return
}

// SetActive is a wrapper around the C function gtk_combo_box_set_active.
func (recv *ComboBox) SetActive(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_combo_box_set_active((*C.GtkComboBox)(recv.native), c_index_)

	return
}

// SetActiveIter is a wrapper around the C function gtk_combo_box_set_active_iter.
func (recv *ComboBox) SetActiveIter(iter *TreeIter) {
	c_iter := (*C.GtkTreeIter)(C.NULL)
	if iter != nil {
		c_iter = (*C.GtkTreeIter)(iter.ToC())
	}

	C.gtk_combo_box_set_active_iter((*C.GtkComboBox)(recv.native), c_iter)

	return
}

// SetColumnSpanColumn is a wrapper around the C function gtk_combo_box_set_column_span_column.
func (recv *ComboBox) SetColumnSpanColumn(columnSpan int32) {
	c_column_span := (C.gint)(columnSpan)

	C.gtk_combo_box_set_column_span_column((*C.GtkComboBox)(recv.native), c_column_span)

	return
}

// SetModel is a wrapper around the C function gtk_combo_box_set_model.
func (recv *ComboBox) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_combo_box_set_model((*C.GtkComboBox)(recv.native), c_model)

	return
}

// SetRowSpanColumn is a wrapper around the C function gtk_combo_box_set_row_span_column.
func (recv *ComboBox) SetRowSpanColumn(rowSpan int32) {
	c_row_span := (C.gint)(rowSpan)

	C.gtk_combo_box_set_row_span_column((*C.GtkComboBox)(recv.native), c_row_span)

	return
}

// SetWrapWidth is a wrapper around the C function gtk_combo_box_set_wrap_width.
func (recv *ComboBox) SetWrapWidth(width int32) {
	c_width := (C.gint)(width)

	C.gtk_combo_box_set_wrap_width((*C.GtkComboBox)(recv.native), c_width)

	return
}

// GetAlignment is a wrapper around the C function gtk_entry_get_alignment.
func (recv *Entry) GetAlignment() float32 {
	retC := C.gtk_entry_get_alignment((*C.GtkEntry)(recv.native))
	retGo := (float32)(retC)

	return retGo
}

// GetCompletion is a wrapper around the C function gtk_entry_get_completion.
func (recv *Entry) GetCompletion() *EntryCompletion {
	retC := C.gtk_entry_get_completion((*C.GtkEntry)(recv.native))
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetAlignment is a wrapper around the C function gtk_entry_set_alignment.
func (recv *Entry) SetAlignment(xalign float32) {
	c_xalign := (C.gfloat)(xalign)

	C.gtk_entry_set_alignment((*C.GtkEntry)(recv.native), c_xalign)

	return
}

// SetCompletion is a wrapper around the C function gtk_entry_set_completion.
func (recv *Entry) SetCompletion(completion *EntryCompletion) {
	c_completion := (*C.GtkEntryCompletion)(C.NULL)
	if completion != nil {
		c_completion = (*C.GtkEntryCompletion)(completion.ToC())
	}

	C.gtk_entry_set_completion((*C.GtkEntry)(recv.native), c_completion)

	return
}

type signalEntryCompletionActionActivatedDetail struct {
	callback  EntryCompletionSignalActionActivatedCallback
	handlerID C.gulong
}

var signalEntryCompletionActionActivatedId int
var signalEntryCompletionActionActivatedMap = make(map[int]signalEntryCompletionActionActivatedDetail)
var signalEntryCompletionActionActivatedLock sync.RWMutex

// EntryCompletionSignalActionActivatedCallback is a callback function for a 'action-activated' signal emitted from a EntryCompletion.
type EntryCompletionSignalActionActivatedCallback func(Index int32)

/*
ConnectActionActivated connects the callback to the 'action-activated' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectActionActivated to remove it.
*/
func (recv *EntryCompletion) ConnectActionActivated(callback EntryCompletionSignalActionActivatedCallback) int {
	signalEntryCompletionActionActivatedLock.Lock()
	defer signalEntryCompletionActionActivatedLock.Unlock()

	signalEntryCompletionActionActivatedId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_action_activated(instance, C.gpointer(uintptr(signalEntryCompletionActionActivatedId)))

	detail := signalEntryCompletionActionActivatedDetail{callback, handlerID}
	signalEntryCompletionActionActivatedMap[signalEntryCompletionActionActivatedId] = detail

	return signalEntryCompletionActionActivatedId
}

/*
DisconnectActionActivated disconnects a callback from the 'action-activated' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectActionActivated.
*/
func (recv *EntryCompletion) DisconnectActionActivated(connectionID int) {
	signalEntryCompletionActionActivatedLock.Lock()
	defer signalEntryCompletionActionActivatedLock.Unlock()

	detail, exists := signalEntryCompletionActionActivatedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionActionActivatedMap, connectionID)
}

//export entrycompletion_actionActivatedHandler
func entrycompletion_actionActivatedHandler(_ *C.GObject, c__index C.gint, data C.gpointer) {
	signalEntryCompletionActionActivatedLock.RLock()
	defer signalEntryCompletionActionActivatedLock.RUnlock()

	Index := int32(c__index)

	index := int(uintptr(data))
	callback := signalEntryCompletionActionActivatedMap[index].callback
	callback(Index)
}

type signalEntryCompletionMatchSelectedDetail struct {
	callback  EntryCompletionSignalMatchSelectedCallback
	handlerID C.gulong
}

var signalEntryCompletionMatchSelectedId int
var signalEntryCompletionMatchSelectedMap = make(map[int]signalEntryCompletionMatchSelectedDetail)
var signalEntryCompletionMatchSelectedLock sync.RWMutex

// EntryCompletionSignalMatchSelectedCallback is a callback function for a 'match-selected' signal emitted from a EntryCompletion.
type EntryCompletionSignalMatchSelectedCallback func(model *TreeModel, iter *TreeIter) bool

/*
ConnectMatchSelected connects the callback to the 'match-selected' signal for the EntryCompletion.

The returned value represents the connection, and may be passed to DisconnectMatchSelected to remove it.
*/
func (recv *EntryCompletion) ConnectMatchSelected(callback EntryCompletionSignalMatchSelectedCallback) int {
	signalEntryCompletionMatchSelectedLock.Lock()
	defer signalEntryCompletionMatchSelectedLock.Unlock()

	signalEntryCompletionMatchSelectedId++
	instance := C.gpointer(recv.native)
	handlerID := C.EntryCompletion_signal_connect_match_selected(instance, C.gpointer(uintptr(signalEntryCompletionMatchSelectedId)))

	detail := signalEntryCompletionMatchSelectedDetail{callback, handlerID}
	signalEntryCompletionMatchSelectedMap[signalEntryCompletionMatchSelectedId] = detail

	return signalEntryCompletionMatchSelectedId
}

/*
DisconnectMatchSelected disconnects a callback from the 'match-selected' signal for the EntryCompletion.

The connectionID should be a value returned from a call to ConnectMatchSelected.
*/
func (recv *EntryCompletion) DisconnectMatchSelected(connectionID int) {
	signalEntryCompletionMatchSelectedLock.Lock()
	defer signalEntryCompletionMatchSelectedLock.Unlock()

	detail, exists := signalEntryCompletionMatchSelectedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalEntryCompletionMatchSelectedMap, connectionID)
}

//export entrycompletion_matchSelectedHandler
func entrycompletion_matchSelectedHandler(_ *C.GObject, c_model *C.GtkTreeModel, c_iter *C.GtkTreeIter, data C.gpointer) C.gboolean {
	signalEntryCompletionMatchSelectedLock.RLock()
	defer signalEntryCompletionMatchSelectedLock.RUnlock()

	model := TreeModelNewFromC(unsafe.Pointer(c_model))

	iter := TreeIterNewFromC(unsafe.Pointer(c_iter))

	index := int(uintptr(data))
	callback := signalEntryCompletionMatchSelectedMap[index].callback
	retGo := callback(model, iter)
	retC :=
		boolToGboolean(retGo)
	return retC
}

// EntryCompletionNew is a wrapper around the C function gtk_entry_completion_new.
func EntryCompletionNew() *EntryCompletion {
	retC := C.gtk_entry_completion_new()
	retGo := EntryCompletionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// Complete is a wrapper around the C function gtk_entry_completion_complete.
func (recv *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete((*C.GtkEntryCompletion)(recv.native))

	return
}

// DeleteAction is a wrapper around the C function gtk_entry_completion_delete_action.
func (recv *EntryCompletion) DeleteAction(index int32) {
	c_index_ := (C.gint)(index)

	C.gtk_entry_completion_delete_action((*C.GtkEntryCompletion)(recv.native), c_index_)

	return
}

// GetEntry is a wrapper around the C function gtk_entry_completion_get_entry.
func (recv *EntryCompletion) GetEntry() *Widget {
	retC := C.gtk_entry_completion_get_entry((*C.GtkEntryCompletion)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetMinimumKeyLength is a wrapper around the C function gtk_entry_completion_get_minimum_key_length.
func (recv *EntryCompletion) GetMinimumKeyLength() int32 {
	retC := C.gtk_entry_completion_get_minimum_key_length((*C.GtkEntryCompletion)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetModel is a wrapper around the C function gtk_entry_completion_get_model.
func (recv *EntryCompletion) GetModel() *TreeModel {
	retC := C.gtk_entry_completion_get_model((*C.GtkEntryCompletion)(recv.native))
	var retGo (*TreeModel)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreeModelNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// InsertActionMarkup is a wrapper around the C function gtk_entry_completion_insert_action_markup.
func (recv *EntryCompletion) InsertActionMarkup(index int32, markup string) {
	c_index_ := (C.gint)(index)

	c_markup := C.CString(markup)
	defer C.free(unsafe.Pointer(c_markup))

	C.gtk_entry_completion_insert_action_markup((*C.GtkEntryCompletion)(recv.native), c_index_, c_markup)

	return
}

// InsertActionText is a wrapper around the C function gtk_entry_completion_insert_action_text.
func (recv *EntryCompletion) InsertActionText(index int32, text string) {
	c_index_ := (C.gint)(index)

	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))

	C.gtk_entry_completion_insert_action_text((*C.GtkEntryCompletion)(recv.native), c_index_, c_text)

	return
}

// Unsupported : gtk_entry_completion_set_match_func : unsupported parameter func : no type generator for EntryCompletionMatchFunc (GtkEntryCompletionMatchFunc) for param func

// SetMinimumKeyLength is a wrapper around the C function gtk_entry_completion_set_minimum_key_length.
func (recv *EntryCompletion) SetMinimumKeyLength(length int32) {
	c_length := (C.gint)(length)

	C.gtk_entry_completion_set_minimum_key_length((*C.GtkEntryCompletion)(recv.native), c_length)

	return
}

// SetModel is a wrapper around the C function gtk_entry_completion_set_model.
func (recv *EntryCompletion) SetModel(model *TreeModel) {
	c_model := (*C.GtkTreeModel)(model.ToC())

	C.gtk_entry_completion_set_model((*C.GtkEntryCompletion)(recv.native), c_model)

	return
}

// SetTextColumn is a wrapper around the C function gtk_entry_completion_set_text_column.
func (recv *EntryCompletion) SetTextColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_entry_completion_set_text_column((*C.GtkEntryCompletion)(recv.native), c_column)

	return
}

// GetAboveChild is a wrapper around the C function gtk_event_box_get_above_child.
func (recv *EventBox) GetAboveChild() bool {
	retC := C.gtk_event_box_get_above_child((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleWindow is a wrapper around the C function gtk_event_box_get_visible_window.
func (recv *EventBox) GetVisibleWindow() bool {
	retC := C.gtk_event_box_get_visible_window((*C.GtkEventBox)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAboveChild is a wrapper around the C function gtk_event_box_set_above_child.
func (recv *EventBox) SetAboveChild(aboveChild bool) {
	c_above_child :=
		boolToGboolean(aboveChild)

	C.gtk_event_box_set_above_child((*C.GtkEventBox)(recv.native), c_above_child)

	return
}

// SetVisibleWindow is a wrapper around the C function gtk_event_box_set_visible_window.
func (recv *EventBox) SetVisibleWindow(visibleWindow bool) {
	c_visible_window :=
		boolToGboolean(visibleWindow)

	C.gtk_event_box_set_visible_window((*C.GtkEventBox)(recv.native), c_visible_window)

	return
}

// ExpanderNew is a wrapper around the C function gtk_expander_new.
func ExpanderNew(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ExpanderNewWithMnemonic is a wrapper around the C function gtk_expander_new_with_mnemonic.
func ExpanderNewWithMnemonic(label string) *Expander {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_expander_new_with_mnemonic(c_label)
	retGo := ExpanderNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExpanded is a wrapper around the C function gtk_expander_get_expanded.
func (recv *Expander) GetExpanded() bool {
	retC := C.gtk_expander_get_expanded((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetLabel is a wrapper around the C function gtk_expander_get_label.
func (recv *Expander) GetLabel() string {
	retC := C.gtk_expander_get_label((*C.GtkExpander)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelWidget is a wrapper around the C function gtk_expander_get_label_widget.
func (recv *Expander) GetLabelWidget() *Widget {
	retC := C.gtk_expander_get_label_widget((*C.GtkExpander)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetSpacing is a wrapper around the C function gtk_expander_get_spacing.
func (recv *Expander) GetSpacing() int32 {
	retC := C.gtk_expander_get_spacing((*C.GtkExpander)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetUseMarkup is a wrapper around the C function gtk_expander_get_use_markup.
func (recv *Expander) GetUseMarkup() bool {
	retC := C.gtk_expander_get_use_markup((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_expander_get_use_underline.
func (recv *Expander) GetUseUnderline() bool {
	retC := C.gtk_expander_get_use_underline((*C.GtkExpander)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetExpanded is a wrapper around the C function gtk_expander_set_expanded.
func (recv *Expander) SetExpanded(expanded bool) {
	c_expanded :=
		boolToGboolean(expanded)

	C.gtk_expander_set_expanded((*C.GtkExpander)(recv.native), c_expanded)

	return
}

// SetLabel is a wrapper around the C function gtk_expander_set_label.
func (recv *Expander) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_expander_set_label((*C.GtkExpander)(recv.native), c_label)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_expander_set_label_widget.
func (recv *Expander) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(C.NULL)
	if labelWidget != nil {
		c_label_widget = (*C.GtkWidget)(labelWidget.ToC())
	}

	C.gtk_expander_set_label_widget((*C.GtkExpander)(recv.native), c_label_widget)

	return
}

// SetSpacing is a wrapper around the C function gtk_expander_set_spacing.
func (recv *Expander) SetSpacing(spacing int32) {
	c_spacing := (C.gint)(spacing)

	C.gtk_expander_set_spacing((*C.GtkExpander)(recv.native), c_spacing)

	return
}

// SetUseMarkup is a wrapper around the C function gtk_expander_set_use_markup.
func (recv *Expander) SetUseMarkup(useMarkup bool) {
	c_use_markup :=
		boolToGboolean(useMarkup)

	C.gtk_expander_set_use_markup((*C.GtkExpander)(recv.native), c_use_markup)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_expander_set_use_underline.
func (recv *Expander) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_expander_set_use_underline((*C.GtkExpander)(recv.native), c_use_underline)

	return
}

// Unsupported : gtk_file_chooser_dialog_new : unsupported parameter ... : varargs

// FileChooserWidgetNew is a wrapper around the C function gtk_file_chooser_widget_new.
func FileChooserWidgetNew(action FileChooserAction) *FileChooserWidget {
	c_action := (C.GtkFileChooserAction)(action)

	retC := C.gtk_file_chooser_widget_new(c_action)
	retGo := FileChooserWidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FileFilterNew is a wrapper around the C function gtk_file_filter_new.
func FileFilterNew() *FileFilter {
	retC := C.gtk_file_filter_new()
	retGo := FileFilterNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_file_filter_add_custom : unsupported parameter func : no type generator for FileFilterFunc (GtkFileFilterFunc) for param func

// AddMimeType is a wrapper around the C function gtk_file_filter_add_mime_type.
func (recv *FileFilter) AddMimeType(mimeType string) {
	c_mime_type := C.CString(mimeType)
	defer C.free(unsafe.Pointer(c_mime_type))

	C.gtk_file_filter_add_mime_type((*C.GtkFileFilter)(recv.native), c_mime_type)

	return
}

// AddPattern is a wrapper around the C function gtk_file_filter_add_pattern.
func (recv *FileFilter) AddPattern(pattern string) {
	c_pattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(c_pattern))

	C.gtk_file_filter_add_pattern((*C.GtkFileFilter)(recv.native), c_pattern)

	return
}

// Filter is a wrapper around the C function gtk_file_filter_filter.
func (recv *FileFilter) Filter(filterInfo *FileFilterInfo) bool {
	c_filter_info := (*C.GtkFileFilterInfo)(C.NULL)
	if filterInfo != nil {
		c_filter_info = (*C.GtkFileFilterInfo)(filterInfo.ToC())
	}

	retC := C.gtk_file_filter_filter((*C.GtkFileFilter)(recv.native), c_filter_info)
	retGo := retC == C.TRUE

	return retGo
}

// GetName is a wrapper around the C function gtk_file_filter_get_name.
func (recv *FileFilter) GetName() string {
	retC := C.gtk_file_filter_get_name((*C.GtkFileFilter)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetNeeded is a wrapper around the C function gtk_file_filter_get_needed.
func (recv *FileFilter) GetNeeded() FileFilterFlags {
	retC := C.gtk_file_filter_get_needed((*C.GtkFileFilter)(recv.native))
	retGo := (FileFilterFlags)(retC)

	return retGo
}

// SetName is a wrapper around the C function gtk_file_filter_set_name.
func (recv *FileFilter) SetName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_filter_set_name((*C.GtkFileFilter)(recv.native), c_name)

	return
}

type signalFontButtonFontSetDetail struct {
	callback  FontButtonSignalFontSetCallback
	handlerID C.gulong
}

var signalFontButtonFontSetId int
var signalFontButtonFontSetMap = make(map[int]signalFontButtonFontSetDetail)
var signalFontButtonFontSetLock sync.RWMutex

// FontButtonSignalFontSetCallback is a callback function for a 'font-set' signal emitted from a FontButton.
type FontButtonSignalFontSetCallback func()

/*
ConnectFontSet connects the callback to the 'font-set' signal for the FontButton.

The returned value represents the connection, and may be passed to DisconnectFontSet to remove it.
*/
func (recv *FontButton) ConnectFontSet(callback FontButtonSignalFontSetCallback) int {
	signalFontButtonFontSetLock.Lock()
	defer signalFontButtonFontSetLock.Unlock()

	signalFontButtonFontSetId++
	instance := C.gpointer(recv.native)
	handlerID := C.FontButton_signal_connect_font_set(instance, C.gpointer(uintptr(signalFontButtonFontSetId)))

	detail := signalFontButtonFontSetDetail{callback, handlerID}
	signalFontButtonFontSetMap[signalFontButtonFontSetId] = detail

	return signalFontButtonFontSetId
}

/*
DisconnectFontSet disconnects a callback from the 'font-set' signal for the FontButton.

The connectionID should be a value returned from a call to ConnectFontSet.
*/
func (recv *FontButton) DisconnectFontSet(connectionID int) {
	signalFontButtonFontSetLock.Lock()
	defer signalFontButtonFontSetLock.Unlock()

	detail, exists := signalFontButtonFontSetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalFontButtonFontSetMap, connectionID)
}

//export fontbutton_fontSetHandler
func fontbutton_fontSetHandler(_ *C.GObject, data C.gpointer) {
	signalFontButtonFontSetLock.RLock()
	defer signalFontButtonFontSetLock.RUnlock()

	index := int(uintptr(data))
	callback := signalFontButtonFontSetMap[index].callback
	callback()
}

// FontButtonNew is a wrapper around the C function gtk_font_button_new.
func FontButtonNew() *FontButton {
	retC := C.gtk_font_button_new()
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// FontButtonNewWithFont is a wrapper around the C function gtk_font_button_new_with_font.
func FontButtonNewWithFont(fontname string) *FontButton {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_new_with_font(c_fontname)
	retGo := FontButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFontName is a wrapper around the C function gtk_font_button_get_font_name.
func (recv *FontButton) GetFontName() string {
	retC := C.gtk_font_button_get_font_name((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetShowSize is a wrapper around the C function gtk_font_button_get_show_size.
func (recv *FontButton) GetShowSize() bool {
	retC := C.gtk_font_button_get_show_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetShowStyle is a wrapper around the C function gtk_font_button_get_show_style.
func (recv *FontButton) GetShowStyle() bool {
	retC := C.gtk_font_button_get_show_style((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetTitle is a wrapper around the C function gtk_font_button_get_title.
func (recv *FontButton) GetTitle() string {
	retC := C.gtk_font_button_get_title((*C.GtkFontButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseFont is a wrapper around the C function gtk_font_button_get_use_font.
func (recv *FontButton) GetUseFont() bool {
	retC := C.gtk_font_button_get_use_font((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUseSize is a wrapper around the C function gtk_font_button_get_use_size.
func (recv *FontButton) GetUseSize() bool {
	retC := C.gtk_font_button_get_use_size((*C.GtkFontButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetFontName is a wrapper around the C function gtk_font_button_set_font_name.
func (recv *FontButton) SetFontName(fontname string) bool {
	c_fontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(c_fontname))

	retC := C.gtk_font_button_set_font_name((*C.GtkFontButton)(recv.native), c_fontname)
	retGo := retC == C.TRUE

	return retGo
}

// SetShowSize is a wrapper around the C function gtk_font_button_set_show_size.
func (recv *FontButton) SetShowSize(showSize bool) {
	c_show_size :=
		boolToGboolean(showSize)

	C.gtk_font_button_set_show_size((*C.GtkFontButton)(recv.native), c_show_size)

	return
}

// SetShowStyle is a wrapper around the C function gtk_font_button_set_show_style.
func (recv *FontButton) SetShowStyle(showStyle bool) {
	c_show_style :=
		boolToGboolean(showStyle)

	C.gtk_font_button_set_show_style((*C.GtkFontButton)(recv.native), c_show_style)

	return
}

// SetTitle is a wrapper around the C function gtk_font_button_set_title.
func (recv *FontButton) SetTitle(title string) {
	c_title := C.CString(title)
	defer C.free(unsafe.Pointer(c_title))

	C.gtk_font_button_set_title((*C.GtkFontButton)(recv.native), c_title)

	return
}

// SetUseFont is a wrapper around the C function gtk_font_button_set_use_font.
func (recv *FontButton) SetUseFont(useFont bool) {
	c_use_font :=
		boolToGboolean(useFont)

	C.gtk_font_button_set_use_font((*C.GtkFontButton)(recv.native), c_use_font)

	return
}

// SetUseSize is a wrapper around the C function gtk_font_button_set_use_size.
func (recv *FontButton) SetUseSize(useSize bool) {
	c_use_size :=
		boolToGboolean(useSize)

	C.gtk_font_button_set_use_size((*C.GtkFontButton)(recv.native), c_use_size)

	return
}

// Copy is a wrapper around the C function gtk_icon_info_copy.
func (recv *IconInfo) Copy() *IconInfo {
	retC := C.gtk_icon_info_copy((*C.GtkIconInfo)(recv.native))
	retGo := IconInfoNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Free is a wrapper around the C function gtk_icon_info_free.
func (recv *IconInfo) Free() {
	C.gtk_icon_info_free((*C.GtkIconInfo)(recv.native))

	return
}

// Unsupported : gtk_icon_info_get_attach_points : unsupported parameter points : output array param points

// GetBaseSize is a wrapper around the C function gtk_icon_info_get_base_size.
func (recv *IconInfo) GetBaseSize() int32 {
	retC := C.gtk_icon_info_get_base_size((*C.GtkIconInfo)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetBuiltinPixbuf is a wrapper around the C function gtk_icon_info_get_builtin_pixbuf.
func (recv *IconInfo) GetBuiltinPixbuf() *gdkpixbuf.Pixbuf {
	retC := C.gtk_icon_info_get_builtin_pixbuf((*C.GtkIconInfo)(recv.native))
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetDisplayName is a wrapper around the C function gtk_icon_info_get_display_name.
func (recv *IconInfo) GetDisplayName() string {
	retC := C.gtk_icon_info_get_display_name((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetEmbeddedRect is a wrapper around the C function gtk_icon_info_get_embedded_rect.
func (recv *IconInfo) GetEmbeddedRect() (bool, *gdk.Rectangle) {
	var c_rectangle C.GdkRectangle

	retC := C.gtk_icon_info_get_embedded_rect((*C.GtkIconInfo)(recv.native), &c_rectangle)
	retGo := retC == C.TRUE

	rectangle := gdk.RectangleNewFromC(unsafe.Pointer(&c_rectangle))

	return retGo, rectangle
}

// GetFilename is a wrapper around the C function gtk_icon_info_get_filename.
func (recv *IconInfo) GetFilename() string {
	retC := C.gtk_icon_info_get_filename((*C.GtkIconInfo)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// LoadIcon is a wrapper around the C function gtk_icon_info_load_icon.
func (recv *IconInfo) LoadIcon() (*gdkpixbuf.Pixbuf, error) {
	var cThrowableError *C.GError

	retC := C.gtk_icon_info_load_icon((*C.GtkIconInfo)(recv.native), &cThrowableError)
	retGo := gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SetRawCoordinates is a wrapper around the C function gtk_icon_info_set_raw_coordinates.
func (recv *IconInfo) SetRawCoordinates(rawCoordinates bool) {
	c_raw_coordinates :=
		boolToGboolean(rawCoordinates)

	C.gtk_icon_info_set_raw_coordinates((*C.GtkIconInfo)(recv.native), c_raw_coordinates)

	return
}

// IconThemeNew is a wrapper around the C function gtk_icon_theme_new.
func IconThemeNew() *IconTheme {
	retC := C.gtk_icon_theme_new()
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// IconThemeAddBuiltinIcon is a wrapper around the C function gtk_icon_theme_add_builtin_icon.
func IconThemeAddBuiltinIcon(iconName string, size int32, pixbuf *gdkpixbuf.Pixbuf) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_pixbuf := (*C.GdkPixbuf)(C.NULL)
	if pixbuf != nil {
		c_pixbuf = (*C.GdkPixbuf)(pixbuf.ToC())
	}

	C.gtk_icon_theme_add_builtin_icon(c_icon_name, c_size, c_pixbuf)

	return
}

// IconThemeGetDefault is a wrapper around the C function gtk_icon_theme_get_default.
func IconThemeGetDefault() *IconTheme {
	retC := C.gtk_icon_theme_get_default()
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// IconThemeGetForScreen is a wrapper around the C function gtk_icon_theme_get_for_screen.
func IconThemeGetForScreen(screen *gdk.Screen) *IconTheme {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	retC := C.gtk_icon_theme_get_for_screen(c_screen)
	retGo := IconThemeNewFromC(unsafe.Pointer(retC))

	return retGo
}

// AppendSearchPath is a wrapper around the C function gtk_icon_theme_append_search_path.
func (recv *IconTheme) AppendSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_append_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// GetExampleIconName is a wrapper around the C function gtk_icon_theme_get_example_icon_name.
func (recv *IconTheme) GetExampleIconName() string {
	retC := C.gtk_icon_theme_get_example_icon_name((*C.GtkIconTheme)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// Unsupported : gtk_icon_theme_get_search_path : unsupported parameter path : output array param path

// HasIcon is a wrapper around the C function gtk_icon_theme_has_icon.
func (recv *IconTheme) HasIcon(iconName string) bool {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	retC := C.gtk_icon_theme_has_icon((*C.GtkIconTheme)(recv.native), c_icon_name)
	retGo := retC == C.TRUE

	return retGo
}

// ListIcons is a wrapper around the C function gtk_icon_theme_list_icons.
func (recv *IconTheme) ListIcons(context string) *glib.List {
	c_context := C.CString(context)
	defer C.free(unsafe.Pointer(c_context))

	retC := C.gtk_icon_theme_list_icons((*C.GtkIconTheme)(recv.native), c_context)
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// LoadIcon is a wrapper around the C function gtk_icon_theme_load_icon.
func (recv *IconTheme) LoadIcon(iconName string, size int32, flags IconLookupFlags) (*gdkpixbuf.Pixbuf, error) {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	var cThrowableError *C.GError

	retC := C.gtk_icon_theme_load_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags, &cThrowableError)
	var retGo (*gdkpixbuf.Pixbuf)
	if retC == nil {
		retGo = nil
	} else {
		retGo = gdkpixbuf.PixbufNewFromC(unsafe.Pointer(retC))
	}

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// LookupIcon is a wrapper around the C function gtk_icon_theme_lookup_icon.
func (recv *IconTheme) LookupIcon(iconName string, size int32, flags IconLookupFlags) *IconInfo {
	c_icon_name := C.CString(iconName)
	defer C.free(unsafe.Pointer(c_icon_name))

	c_size := (C.gint)(size)

	c_flags := (C.GtkIconLookupFlags)(flags)

	retC := C.gtk_icon_theme_lookup_icon((*C.GtkIconTheme)(recv.native), c_icon_name, c_size, c_flags)
	var retGo (*IconInfo)
	if retC == nil {
		retGo = nil
	} else {
		retGo = IconInfoNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// PrependSearchPath is a wrapper around the C function gtk_icon_theme_prepend_search_path.
func (recv *IconTheme) PrependSearchPath(path string) {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	C.gtk_icon_theme_prepend_search_path((*C.GtkIconTheme)(recv.native), c_path)

	return
}

// RescanIfNeeded is a wrapper around the C function gtk_icon_theme_rescan_if_needed.
func (recv *IconTheme) RescanIfNeeded() bool {
	retC := C.gtk_icon_theme_rescan_if_needed((*C.GtkIconTheme)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetCustomTheme is a wrapper around the C function gtk_icon_theme_set_custom_theme.
func (recv *IconTheme) SetCustomTheme(themeName string) {
	c_theme_name := C.CString(themeName)
	defer C.free(unsafe.Pointer(c_theme_name))

	C.gtk_icon_theme_set_custom_theme((*C.GtkIconTheme)(recv.native), c_theme_name)

	return
}

// SetScreen is a wrapper around the C function gtk_icon_theme_set_screen.
func (recv *IconTheme) SetScreen(screen *gdk.Screen) {
	c_screen := (*C.GdkScreen)(C.NULL)
	if screen != nil {
		c_screen = (*C.GdkScreen)(screen.ToC())
	}

	C.gtk_icon_theme_set_screen((*C.GtkIconTheme)(recv.native), c_screen)

	return
}

// Blacklisted : gtk_icon_theme_set_search_path

// Attach is a wrapper around the C function gtk_menu_attach.
func (recv *Menu) Attach(child *Widget, leftAttach uint32, rightAttach uint32, topAttach uint32, bottomAttach uint32) {
	c_child := (*C.GtkWidget)(C.NULL)
	if child != nil {
		c_child = (*C.GtkWidget)(child.ToC())
	}

	c_left_attach := (C.guint)(leftAttach)

	c_right_attach := (C.guint)(rightAttach)

	c_top_attach := (C.guint)(topAttach)

	c_bottom_attach := (C.guint)(bottomAttach)

	C.gtk_menu_attach((*C.GtkMenu)(recv.native), c_child, c_left_attach, c_right_attach, c_top_attach, c_bottom_attach)

	return
}

// SetMonitor is a wrapper around the C function gtk_menu_set_monitor.
func (recv *Menu) SetMonitor(monitorNum int32) {
	c_monitor_num := (C.gint)(monitorNum)

	C.gtk_menu_set_monitor((*C.GtkMenu)(recv.native), c_monitor_num)

	return
}

// Cancel is a wrapper around the C function gtk_menu_shell_cancel.
func (recv *MenuShell) Cancel() {
	C.gtk_menu_shell_cancel((*C.GtkMenuShell)(recv.native))

	return
}

// MessageDialogNewWithMarkup is a wrapper around the C function gtk_message_dialog_new_with_markup.
func MessageDialogNewWithMarkup(parent *Window, flags DialogFlags, type_ MessageType, buttons ButtonsType, messageFormat string, args ...interface{}) *MessageDialog {
	c_parent := (*C.GtkWindow)(C.NULL)
	if parent != nil {
		c_parent = (*C.GtkWindow)(parent.ToC())
	}

	c_flags := (C.GtkDialogFlags)(flags)

	c_type := (C.GtkMessageType)(type_)

	c_buttons := (C.GtkButtonsType)(buttons)

	goFormattedString := fmt.Sprintf(messageFormat, args...)
	c_message_format := C.CString(goFormattedString)
	defer C.free(unsafe.Pointer(c_message_format))

	retC := C._gtk_message_dialog_new_with_markup(c_parent, c_flags, c_type, c_buttons, c_message_format)
	retGo := MessageDialogNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetMarkup is a wrapper around the C function gtk_message_dialog_set_markup.
func (recv *MessageDialog) SetMarkup(str string) {
	c_str := C.CString(str)
	defer C.free(unsafe.Pointer(c_str))

	C.gtk_message_dialog_set_markup((*C.GtkMessageDialog)(recv.native), c_str)

	return
}

// GetChild1 is a wrapper around the C function gtk_paned_get_child1.
func (recv *Paned) GetChild1() *Widget {
	retC := C.gtk_paned_get_child1((*C.GtkPaned)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetChild2 is a wrapper around the C function gtk_paned_get_child2.
func (recv *Paned) GetChild2() *Widget {
	retC := C.gtk_paned_get_child2((*C.GtkPaned)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

type signalRadioActionChangedDetail struct {
	callback  RadioActionSignalChangedCallback
	handlerID C.gulong
}

var signalRadioActionChangedId int
var signalRadioActionChangedMap = make(map[int]signalRadioActionChangedDetail)
var signalRadioActionChangedLock sync.RWMutex

// RadioActionSignalChangedCallback is a callback function for a 'changed' signal emitted from a RadioAction.
type RadioActionSignalChangedCallback func(current *RadioAction)

/*
ConnectChanged connects the callback to the 'changed' signal for the RadioAction.

The returned value represents the connection, and may be passed to DisconnectChanged to remove it.
*/
func (recv *RadioAction) ConnectChanged(callback RadioActionSignalChangedCallback) int {
	signalRadioActionChangedLock.Lock()
	defer signalRadioActionChangedLock.Unlock()

	signalRadioActionChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RadioAction_signal_connect_changed(instance, C.gpointer(uintptr(signalRadioActionChangedId)))

	detail := signalRadioActionChangedDetail{callback, handlerID}
	signalRadioActionChangedMap[signalRadioActionChangedId] = detail

	return signalRadioActionChangedId
}

/*
DisconnectChanged disconnects a callback from the 'changed' signal for the RadioAction.

The connectionID should be a value returned from a call to ConnectChanged.
*/
func (recv *RadioAction) DisconnectChanged(connectionID int) {
	signalRadioActionChangedLock.Lock()
	defer signalRadioActionChangedLock.Unlock()

	detail, exists := signalRadioActionChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRadioActionChangedMap, connectionID)
}

//export radioaction_changedHandler
func radioaction_changedHandler(_ *C.GObject, c_current *C.GtkRadioAction, data C.gpointer) {
	signalRadioActionChangedLock.RLock()
	defer signalRadioActionChangedLock.RUnlock()

	current := RadioActionNewFromC(unsafe.Pointer(c_current))

	index := int(uintptr(data))
	callback := signalRadioActionChangedMap[index].callback
	callback(current)
}

// RadioActionNew is a wrapper around the C function gtk_radio_action_new.
func RadioActionNew(name string, label string, tooltip string, stockId string, value int32) *RadioAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	c_value := (C.gint)(value)

	retC := C.gtk_radio_action_new(c_name, c_label, c_tooltip, c_stock_id, c_value)
	retGo := RadioActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetCurrentValue is a wrapper around the C function gtk_radio_action_get_current_value.
func (recv *RadioAction) GetCurrentValue() int32 {
	retC := C.gtk_radio_action_get_current_value((*C.GtkRadioAction)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_action_get_group.
func (recv *RadioAction) GetGroup() *glib.SList {
	retC := C.gtk_radio_action_get_group((*C.GtkRadioAction)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_action_set_group.
func (recv *RadioAction) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	C.gtk_radio_action_set_group((*C.GtkRadioAction)(recv.native), c_group)

	return
}

type signalRadioButtonGroupChangedDetail struct {
	callback  RadioButtonSignalGroupChangedCallback
	handlerID C.gulong
}

var signalRadioButtonGroupChangedId int
var signalRadioButtonGroupChangedMap = make(map[int]signalRadioButtonGroupChangedDetail)
var signalRadioButtonGroupChangedLock sync.RWMutex

// RadioButtonSignalGroupChangedCallback is a callback function for a 'group-changed' signal emitted from a RadioButton.
type RadioButtonSignalGroupChangedCallback func()

/*
ConnectGroupChanged connects the callback to the 'group-changed' signal for the RadioButton.

The returned value represents the connection, and may be passed to DisconnectGroupChanged to remove it.
*/
func (recv *RadioButton) ConnectGroupChanged(callback RadioButtonSignalGroupChangedCallback) int {
	signalRadioButtonGroupChangedLock.Lock()
	defer signalRadioButtonGroupChangedLock.Unlock()

	signalRadioButtonGroupChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.RadioButton_signal_connect_group_changed(instance, C.gpointer(uintptr(signalRadioButtonGroupChangedId)))

	detail := signalRadioButtonGroupChangedDetail{callback, handlerID}
	signalRadioButtonGroupChangedMap[signalRadioButtonGroupChangedId] = detail

	return signalRadioButtonGroupChangedId
}

/*
DisconnectGroupChanged disconnects a callback from the 'group-changed' signal for the RadioButton.

The connectionID should be a value returned from a call to ConnectGroupChanged.
*/
func (recv *RadioButton) DisconnectGroupChanged(connectionID int) {
	signalRadioButtonGroupChangedLock.Lock()
	defer signalRadioButtonGroupChangedLock.Unlock()

	detail, exists := signalRadioButtonGroupChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalRadioButtonGroupChangedMap, connectionID)
}

//export radiobutton_groupChangedHandler
func radiobutton_groupChangedHandler(_ *C.GObject, data C.gpointer) {
	signalRadioButtonGroupChangedLock.RLock()
	defer signalRadioButtonGroupChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalRadioButtonGroupChangedMap[index].callback
	callback()
}

// RadioMenuItemNewFromWidget is a wrapper around the C function gtk_radio_menu_item_new_from_widget.
func RadioMenuItemNewFromWidget(group *RadioMenuItem) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	retC := C.gtk_radio_menu_item_new_from_widget(c_group)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithLabelFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_label_from_widget.
func RadioMenuItemNewWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_label_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioMenuItemNewWithMnemonicFromWidget is a wrapper around the C function gtk_radio_menu_item_new_with_mnemonic_from_widget.
func RadioMenuItemNewWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	c_group := (*C.GtkRadioMenuItem)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioMenuItem)(group.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_radio_menu_item_new_with_mnemonic_from_widget(c_group, c_label)
	retGo := RadioMenuItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNew is a wrapper around the C function gtk_radio_tool_button_new.
func RadioToolButtonNew(group *glib.SList) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromStock is a wrapper around the C function gtk_radio_tool_button_new_from_stock.
func RadioToolButtonNewFromStock(group *glib.SList, stockId string) *RadioToolButton {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_from_stock(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewFromWidget is a wrapper around the C function gtk_radio_tool_button_new_from_widget.
func RadioToolButtonNewFromWidget(group *RadioToolButton) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	retC := C.gtk_radio_tool_button_new_from_widget(c_group)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// RadioToolButtonNewWithStockFromWidget is a wrapper around the C function gtk_radio_tool_button_new_with_stock_from_widget.
func RadioToolButtonNewWithStockFromWidget(group *RadioToolButton, stockId string) *RadioToolButton {
	c_group := (*C.GtkRadioToolButton)(C.NULL)
	if group != nil {
		c_group = (*C.GtkRadioToolButton)(group.ToC())
	}

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_radio_tool_button_new_with_stock_from_widget(c_group, c_stock_id)
	retGo := RadioToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetGroup is a wrapper around the C function gtk_radio_tool_button_get_group.
func (recv *RadioToolButton) GetGroup() *glib.SList {
	retC := C.gtk_radio_tool_button_get_group((*C.GtkRadioToolButton)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetGroup is a wrapper around the C function gtk_radio_tool_button_set_group.
func (recv *RadioToolButton) SetGroup(group *glib.SList) {
	c_group := (*C.GSList)(C.NULL)
	if group != nil {
		c_group = (*C.GSList)(group.ToC())
	}

	C.gtk_radio_tool_button_set_group((*C.GtkRadioToolButton)(recv.native), c_group)

	return
}

// GetLayout is a wrapper around the C function gtk_scale_get_layout.
func (recv *Scale) GetLayout() *pango.Layout {
	retC := C.gtk_scale_get_layout((*C.GtkScale)(recv.native))
	var retGo (*pango.Layout)
	if retC == nil {
		retGo = nil
	} else {
		retGo = pango.LayoutNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLayoutOffsets is a wrapper around the C function gtk_scale_get_layout_offsets.
func (recv *Scale) GetLayoutOffsets() (int32, int32) {
	var c_x C.gint

	var c_y C.gint

	C.gtk_scale_get_layout_offsets((*C.GtkScale)(recv.native), &c_x, &c_y)

	x := (int32)(c_x)

	y := (int32)(c_y)

	return x, y
}

// SeparatorToolItemNew is a wrapper around the C function gtk_separator_tool_item_new.
func SeparatorToolItemNew() *SeparatorToolItem {
	retC := C.gtk_separator_tool_item_new()
	retGo := SeparatorToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetDraw is a wrapper around the C function gtk_separator_tool_item_get_draw.
func (recv *SeparatorToolItem) GetDraw() bool {
	retC := C.gtk_separator_tool_item_get_draw((*C.GtkSeparatorToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetDraw is a wrapper around the C function gtk_separator_tool_item_set_draw.
func (recv *SeparatorToolItem) SetDraw(draw bool) {
	c_draw :=
		boolToGboolean(draw)

	C.gtk_separator_tool_item_set_draw((*C.GtkSeparatorToolItem)(recv.native), c_draw)

	return
}

type signalStyleRealizeDetail struct {
	callback  StyleSignalRealizeCallback
	handlerID C.gulong
}

var signalStyleRealizeId int
var signalStyleRealizeMap = make(map[int]signalStyleRealizeDetail)
var signalStyleRealizeLock sync.RWMutex

// StyleSignalRealizeCallback is a callback function for a 'realize' signal emitted from a Style.
type StyleSignalRealizeCallback func()

/*
ConnectRealize connects the callback to the 'realize' signal for the Style.

The returned value represents the connection, and may be passed to DisconnectRealize to remove it.
*/
func (recv *Style) ConnectRealize(callback StyleSignalRealizeCallback) int {
	signalStyleRealizeLock.Lock()
	defer signalStyleRealizeLock.Unlock()

	signalStyleRealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Style_signal_connect_realize(instance, C.gpointer(uintptr(signalStyleRealizeId)))

	detail := signalStyleRealizeDetail{callback, handlerID}
	signalStyleRealizeMap[signalStyleRealizeId] = detail

	return signalStyleRealizeId
}

/*
DisconnectRealize disconnects a callback from the 'realize' signal for the Style.

The connectionID should be a value returned from a call to ConnectRealize.
*/
func (recv *Style) DisconnectRealize(connectionID int) {
	signalStyleRealizeLock.Lock()
	defer signalStyleRealizeLock.Unlock()

	detail, exists := signalStyleRealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleRealizeMap, connectionID)
}

//export style_realizeHandler
func style_realizeHandler(_ *C.GObject, data C.gpointer) {
	signalStyleRealizeLock.RLock()
	defer signalStyleRealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleRealizeMap[index].callback
	callback()
}

type signalStyleUnrealizeDetail struct {
	callback  StyleSignalUnrealizeCallback
	handlerID C.gulong
}

var signalStyleUnrealizeId int
var signalStyleUnrealizeMap = make(map[int]signalStyleUnrealizeDetail)
var signalStyleUnrealizeLock sync.RWMutex

// StyleSignalUnrealizeCallback is a callback function for a 'unrealize' signal emitted from a Style.
type StyleSignalUnrealizeCallback func()

/*
ConnectUnrealize connects the callback to the 'unrealize' signal for the Style.

The returned value represents the connection, and may be passed to DisconnectUnrealize to remove it.
*/
func (recv *Style) ConnectUnrealize(callback StyleSignalUnrealizeCallback) int {
	signalStyleUnrealizeLock.Lock()
	defer signalStyleUnrealizeLock.Unlock()

	signalStyleUnrealizeId++
	instance := C.gpointer(recv.native)
	handlerID := C.Style_signal_connect_unrealize(instance, C.gpointer(uintptr(signalStyleUnrealizeId)))

	detail := signalStyleUnrealizeDetail{callback, handlerID}
	signalStyleUnrealizeMap[signalStyleUnrealizeId] = detail

	return signalStyleUnrealizeId
}

/*
DisconnectUnrealize disconnects a callback from the 'unrealize' signal for the Style.

The connectionID should be a value returned from a call to ConnectUnrealize.
*/
func (recv *Style) DisconnectUnrealize(connectionID int) {
	signalStyleUnrealizeLock.Lock()
	defer signalStyleUnrealizeLock.Unlock()

	detail, exists := signalStyleUnrealizeMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalStyleUnrealizeMap, connectionID)
}

//export style_unrealizeHandler
func style_unrealizeHandler(_ *C.GObject, data C.gpointer) {
	signalStyleUnrealizeLock.RLock()
	defer signalStyleUnrealizeLock.RUnlock()

	index := int(uintptr(data))
	callback := signalStyleUnrealizeMap[index].callback
	callback()
}

// SelectRange is a wrapper around the C function gtk_text_buffer_select_range.
func (recv *TextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	c_ins := (*C.GtkTextIter)(C.NULL)
	if ins != nil {
		c_ins = (*C.GtkTextIter)(ins.ToC())
	}

	c_bound := (*C.GtkTextIter)(C.NULL)
	if bound != nil {
		c_bound = (*C.GtkTextIter)(bound.ToC())
	}

	C.gtk_text_buffer_select_range((*C.GtkTextBuffer)(recv.native), c_ins, c_bound)

	return
}

// GetAcceptsTab is a wrapper around the C function gtk_text_view_get_accepts_tab.
func (recv *TextView) GetAcceptsTab() bool {
	retC := C.gtk_text_view_get_accepts_tab((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetOverwrite is a wrapper around the C function gtk_text_view_get_overwrite.
func (recv *TextView) GetOverwrite() bool {
	retC := C.gtk_text_view_get_overwrite((*C.GtkTextView)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetAcceptsTab is a wrapper around the C function gtk_text_view_set_accepts_tab.
func (recv *TextView) SetAcceptsTab(acceptsTab bool) {
	c_accepts_tab :=
		boolToGboolean(acceptsTab)

	C.gtk_text_view_set_accepts_tab((*C.GtkTextView)(recv.native), c_accepts_tab)

	return
}

// SetOverwrite is a wrapper around the C function gtk_text_view_set_overwrite.
func (recv *TextView) SetOverwrite(overwrite bool) {
	c_overwrite :=
		boolToGboolean(overwrite)

	C.gtk_text_view_set_overwrite((*C.GtkTextView)(recv.native), c_overwrite)

	return
}

// ToggleActionNew is a wrapper around the C function gtk_toggle_action_new.
func ToggleActionNew(name string, label string, tooltip string, stockId string) *ToggleAction {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	c_tooltip := C.CString(tooltip)
	defer C.free(unsafe.Pointer(c_tooltip))

	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_action_new(c_name, c_label, c_tooltip, c_stock_id)
	retGo := ToggleActionNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_action_get_active.
func (recv *ToggleAction) GetActive() bool {
	retC := C.gtk_toggle_action_get_active((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetDrawAsRadio is a wrapper around the C function gtk_toggle_action_get_draw_as_radio.
func (recv *ToggleAction) GetDrawAsRadio() bool {
	retC := C.gtk_toggle_action_get_draw_as_radio((*C.GtkToggleAction)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_toggle_action_set_active.
func (recv *ToggleAction) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_action_set_active((*C.GtkToggleAction)(recv.native), c_is_active)

	return
}

// SetDrawAsRadio is a wrapper around the C function gtk_toggle_action_set_draw_as_radio.
func (recv *ToggleAction) SetDrawAsRadio(drawAsRadio bool) {
	c_draw_as_radio :=
		boolToGboolean(drawAsRadio)

	C.gtk_toggle_action_set_draw_as_radio((*C.GtkToggleAction)(recv.native), c_draw_as_radio)

	return
}

// Toggled is a wrapper around the C function gtk_toggle_action_toggled.
func (recv *ToggleAction) Toggled() {
	C.gtk_toggle_action_toggled((*C.GtkToggleAction)(recv.native))

	return
}

// ToggleToolButtonNew is a wrapper around the C function gtk_toggle_tool_button_new.
func ToggleToolButtonNew() *ToggleToolButton {
	retC := C.gtk_toggle_tool_button_new()
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToggleToolButtonNewFromStock is a wrapper around the C function gtk_toggle_tool_button_new_from_stock.
func ToggleToolButtonNewFromStock(stockId string) *ToggleToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_toggle_tool_button_new_from_stock(c_stock_id)
	retGo := ToggleToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActive is a wrapper around the C function gtk_toggle_tool_button_get_active.
func (recv *ToggleToolButton) GetActive() bool {
	retC := C.gtk_toggle_tool_button_get_active((*C.GtkToggleToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetActive is a wrapper around the C function gtk_toggle_tool_button_set_active.
func (recv *ToggleToolButton) SetActive(isActive bool) {
	c_is_active :=
		boolToGboolean(isActive)

	C.gtk_toggle_tool_button_set_active((*C.GtkToggleToolButton)(recv.native), c_is_active)

	return
}

// ToolButtonNew is a wrapper around the C function gtk_tool_button_new.
func ToolButtonNew(iconWidget *Widget, label string) *ToolButton {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	retC := C.gtk_tool_button_new(c_icon_widget, c_label)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ToolButtonNewFromStock is a wrapper around the C function gtk_tool_button_new_from_stock.
func ToolButtonNewFromStock(stockId string) *ToolButton {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	retC := C.gtk_tool_button_new_from_stock(c_stock_id)
	retGo := ToolButtonNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetIconWidget is a wrapper around the C function gtk_tool_button_get_icon_widget.
func (recv *ToolButton) GetIconWidget() *Widget {
	retC := C.gtk_tool_button_get_icon_widget((*C.GtkToolButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLabel is a wrapper around the C function gtk_tool_button_get_label.
func (recv *ToolButton) GetLabel() string {
	retC := C.gtk_tool_button_get_label((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetLabelWidget is a wrapper around the C function gtk_tool_button_get_label_widget.
func (recv *ToolButton) GetLabelWidget() *Widget {
	retC := C.gtk_tool_button_get_label_widget((*C.GtkToolButton)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetStockId is a wrapper around the C function gtk_tool_button_get_stock_id.
func (recv *ToolButton) GetStockId() string {
	retC := C.gtk_tool_button_get_stock_id((*C.GtkToolButton)(recv.native))
	retGo := C.GoString(retC)

	return retGo
}

// GetUseUnderline is a wrapper around the C function gtk_tool_button_get_use_underline.
func (recv *ToolButton) GetUseUnderline() bool {
	retC := C.gtk_tool_button_get_use_underline((*C.GtkToolButton)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetIconWidget is a wrapper around the C function gtk_tool_button_set_icon_widget.
func (recv *ToolButton) SetIconWidget(iconWidget *Widget) {
	c_icon_widget := (*C.GtkWidget)(C.NULL)
	if iconWidget != nil {
		c_icon_widget = (*C.GtkWidget)(iconWidget.ToC())
	}

	C.gtk_tool_button_set_icon_widget((*C.GtkToolButton)(recv.native), c_icon_widget)

	return
}

// SetLabel is a wrapper around the C function gtk_tool_button_set_label.
func (recv *ToolButton) SetLabel(label string) {
	c_label := C.CString(label)
	defer C.free(unsafe.Pointer(c_label))

	C.gtk_tool_button_set_label((*C.GtkToolButton)(recv.native), c_label)

	return
}

// SetLabelWidget is a wrapper around the C function gtk_tool_button_set_label_widget.
func (recv *ToolButton) SetLabelWidget(labelWidget *Widget) {
	c_label_widget := (*C.GtkWidget)(C.NULL)
	if labelWidget != nil {
		c_label_widget = (*C.GtkWidget)(labelWidget.ToC())
	}

	C.gtk_tool_button_set_label_widget((*C.GtkToolButton)(recv.native), c_label_widget)

	return
}

// SetStockId is a wrapper around the C function gtk_tool_button_set_stock_id.
func (recv *ToolButton) SetStockId(stockId string) {
	c_stock_id := C.CString(stockId)
	defer C.free(unsafe.Pointer(c_stock_id))

	C.gtk_tool_button_set_stock_id((*C.GtkToolButton)(recv.native), c_stock_id)

	return
}

// SetUseUnderline is a wrapper around the C function gtk_tool_button_set_use_underline.
func (recv *ToolButton) SetUseUnderline(useUnderline bool) {
	c_use_underline :=
		boolToGboolean(useUnderline)

	C.gtk_tool_button_set_use_underline((*C.GtkToolButton)(recv.native), c_use_underline)

	return
}

// ToolItemNew is a wrapper around the C function gtk_tool_item_new.
func ToolItemNew() *ToolItem {
	retC := C.gtk_tool_item_new()
	retGo := ToolItemNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetExpand is a wrapper around the C function gtk_tool_item_get_expand.
func (recv *ToolItem) GetExpand() bool {
	retC := C.gtk_tool_item_get_expand((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetHomogeneous is a wrapper around the C function gtk_tool_item_get_homogeneous.
func (recv *ToolItem) GetHomogeneous() bool {
	retC := C.gtk_tool_item_get_homogeneous((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetIconSize is a wrapper around the C function gtk_tool_item_get_icon_size.
func (recv *ToolItem) GetIconSize() IconSize {
	retC := C.gtk_tool_item_get_icon_size((*C.GtkToolItem)(recv.native))
	retGo := (IconSize)(retC)

	return retGo
}

// GetIsImportant is a wrapper around the C function gtk_tool_item_get_is_important.
func (recv *ToolItem) GetIsImportant() bool {
	retC := C.gtk_tool_item_get_is_important((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetOrientation is a wrapper around the C function gtk_tool_item_get_orientation.
func (recv *ToolItem) GetOrientation() Orientation {
	retC := C.gtk_tool_item_get_orientation((*C.GtkToolItem)(recv.native))
	retGo := (Orientation)(retC)

	return retGo
}

// GetProxyMenuItem is a wrapper around the C function gtk_tool_item_get_proxy_menu_item.
func (recv *ToolItem) GetProxyMenuItem(menuItemId string) *Widget {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	retC := C.gtk_tool_item_get_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id)
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_tool_item_get_relief_style.
func (recv *ToolItem) GetReliefStyle() ReliefStyle {
	retC := C.gtk_tool_item_get_relief_style((*C.GtkToolItem)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetToolbarStyle is a wrapper around the C function gtk_tool_item_get_toolbar_style.
func (recv *ToolItem) GetToolbarStyle() ToolbarStyle {
	retC := C.gtk_tool_item_get_toolbar_style((*C.GtkToolItem)(recv.native))
	retGo := (ToolbarStyle)(retC)

	return retGo
}

// GetUseDragWindow is a wrapper around the C function gtk_tool_item_get_use_drag_window.
func (recv *ToolItem) GetUseDragWindow() bool {
	retC := C.gtk_tool_item_get_use_drag_window((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleHorizontal is a wrapper around the C function gtk_tool_item_get_visible_horizontal.
func (recv *ToolItem) GetVisibleHorizontal() bool {
	retC := C.gtk_tool_item_get_visible_horizontal((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetVisibleVertical is a wrapper around the C function gtk_tool_item_get_visible_vertical.
func (recv *ToolItem) GetVisibleVertical() bool {
	retC := C.gtk_tool_item_get_visible_vertical((*C.GtkToolItem)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// RetrieveProxyMenuItem is a wrapper around the C function gtk_tool_item_retrieve_proxy_menu_item.
func (recv *ToolItem) RetrieveProxyMenuItem() *Widget {
	retC := C.gtk_tool_item_retrieve_proxy_menu_item((*C.GtkToolItem)(recv.native))
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// SetExpand is a wrapper around the C function gtk_tool_item_set_expand.
func (recv *ToolItem) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tool_item_set_expand((*C.GtkToolItem)(recv.native), c_expand)

	return
}

// SetHomogeneous is a wrapper around the C function gtk_tool_item_set_homogeneous.
func (recv *ToolItem) SetHomogeneous(homogeneous bool) {
	c_homogeneous :=
		boolToGboolean(homogeneous)

	C.gtk_tool_item_set_homogeneous((*C.GtkToolItem)(recv.native), c_homogeneous)

	return
}

// SetIsImportant is a wrapper around the C function gtk_tool_item_set_is_important.
func (recv *ToolItem) SetIsImportant(isImportant bool) {
	c_is_important :=
		boolToGboolean(isImportant)

	C.gtk_tool_item_set_is_important((*C.GtkToolItem)(recv.native), c_is_important)

	return
}

// SetProxyMenuItem is a wrapper around the C function gtk_tool_item_set_proxy_menu_item.
func (recv *ToolItem) SetProxyMenuItem(menuItemId string, menuItem *Widget) {
	c_menu_item_id := C.CString(menuItemId)
	defer C.free(unsafe.Pointer(c_menu_item_id))

	c_menu_item := (*C.GtkWidget)(C.NULL)
	if menuItem != nil {
		c_menu_item = (*C.GtkWidget)(menuItem.ToC())
	}

	C.gtk_tool_item_set_proxy_menu_item((*C.GtkToolItem)(recv.native), c_menu_item_id, c_menu_item)

	return
}

// SetUseDragWindow is a wrapper around the C function gtk_tool_item_set_use_drag_window.
func (recv *ToolItem) SetUseDragWindow(useDragWindow bool) {
	c_use_drag_window :=
		boolToGboolean(useDragWindow)

	C.gtk_tool_item_set_use_drag_window((*C.GtkToolItem)(recv.native), c_use_drag_window)

	return
}

// SetVisibleHorizontal is a wrapper around the C function gtk_tool_item_set_visible_horizontal.
func (recv *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	c_visible_horizontal :=
		boolToGboolean(visibleHorizontal)

	C.gtk_tool_item_set_visible_horizontal((*C.GtkToolItem)(recv.native), c_visible_horizontal)

	return
}

// SetVisibleVertical is a wrapper around the C function gtk_tool_item_set_visible_vertical.
func (recv *ToolItem) SetVisibleVertical(visibleVertical bool) {
	c_visible_vertical :=
		boolToGboolean(visibleVertical)

	C.gtk_tool_item_set_visible_vertical((*C.GtkToolItem)(recv.native), c_visible_vertical)

	return
}

// GetDropIndex is a wrapper around the C function gtk_toolbar_get_drop_index.
func (recv *Toolbar) GetDropIndex(x int32, y int32) int32 {
	c_x := (C.gint)(x)

	c_y := (C.gint)(y)

	retC := C.gtk_toolbar_get_drop_index((*C.GtkToolbar)(recv.native), c_x, c_y)
	retGo := (int32)(retC)

	return retGo
}

// GetItemIndex is a wrapper around the C function gtk_toolbar_get_item_index.
func (recv *Toolbar) GetItemIndex(item *ToolItem) int32 {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	retC := C.gtk_toolbar_get_item_index((*C.GtkToolbar)(recv.native), c_item)
	retGo := (int32)(retC)

	return retGo
}

// GetNItems is a wrapper around the C function gtk_toolbar_get_n_items.
func (recv *Toolbar) GetNItems() int32 {
	retC := C.gtk_toolbar_get_n_items((*C.GtkToolbar)(recv.native))
	retGo := (int32)(retC)

	return retGo
}

// GetNthItem is a wrapper around the C function gtk_toolbar_get_nth_item.
func (recv *Toolbar) GetNthItem(n int32) *ToolItem {
	c_n := (C.gint)(n)

	retC := C.gtk_toolbar_get_nth_item((*C.GtkToolbar)(recv.native), c_n)
	var retGo (*ToolItem)
	if retC == nil {
		retGo = nil
	} else {
		retGo = ToolItemNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetReliefStyle is a wrapper around the C function gtk_toolbar_get_relief_style.
func (recv *Toolbar) GetReliefStyle() ReliefStyle {
	retC := C.gtk_toolbar_get_relief_style((*C.GtkToolbar)(recv.native))
	retGo := (ReliefStyle)(retC)

	return retGo
}

// GetShowArrow is a wrapper around the C function gtk_toolbar_get_show_arrow.
func (recv *Toolbar) GetShowArrow() bool {
	retC := C.gtk_toolbar_get_show_arrow((*C.GtkToolbar)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// Insert is a wrapper around the C function gtk_toolbar_insert.
func (recv *Toolbar) Insert(item *ToolItem, pos int32) {
	c_item := (*C.GtkToolItem)(C.NULL)
	if item != nil {
		c_item = (*C.GtkToolItem)(item.ToC())
	}

	c_pos := (C.gint)(pos)

	C.gtk_toolbar_insert((*C.GtkToolbar)(recv.native), c_item, c_pos)

	return
}

// SetDropHighlightItem is a wrapper around the C function gtk_toolbar_set_drop_highlight_item.
func (recv *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index int32) {
	c_tool_item := (*C.GtkToolItem)(C.NULL)
	if toolItem != nil {
		c_tool_item = (*C.GtkToolItem)(toolItem.ToC())
	}

	c_index_ := (C.gint)(index)

	C.gtk_toolbar_set_drop_highlight_item((*C.GtkToolbar)(recv.native), c_tool_item, c_index_)

	return
}

// SetShowArrow is a wrapper around the C function gtk_toolbar_set_show_arrow.
func (recv *Toolbar) SetShowArrow(showArrow bool) {
	c_show_arrow :=
		boolToGboolean(showArrow)

	C.gtk_toolbar_set_show_arrow((*C.GtkToolbar)(recv.native), c_show_arrow)

	return
}

// ClearCache is a wrapper around the C function gtk_tree_model_filter_clear_cache.
func (recv *TreeModelFilter) ClearCache() {
	C.gtk_tree_model_filter_clear_cache((*C.GtkTreeModelFilter)(recv.native))

	return
}

// ConvertChildIterToIter is a wrapper around the C function gtk_tree_model_filter_convert_child_iter_to_iter.
func (recv *TreeModelFilter) ConvertChildIterToIter(childIter *TreeIter) (bool, *TreeIter) {
	var c_filter_iter C.GtkTreeIter

	c_child_iter := (*C.GtkTreeIter)(C.NULL)
	if childIter != nil {
		c_child_iter = (*C.GtkTreeIter)(childIter.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_child_iter_to_iter((*C.GtkTreeModelFilter)(recv.native), &c_filter_iter, c_child_iter)
	retGo := retC == C.TRUE

	filterIter := TreeIterNewFromC(unsafe.Pointer(&c_filter_iter))

	return retGo, filterIter
}

// ConvertChildPathToPath is a wrapper around the C function gtk_tree_model_filter_convert_child_path_to_path.
func (recv *TreeModelFilter) ConvertChildPathToPath(childPath *TreePath) *TreePath {
	c_child_path := (*C.GtkTreePath)(C.NULL)
	if childPath != nil {
		c_child_path = (*C.GtkTreePath)(childPath.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_child_path_to_path((*C.GtkTreeModelFilter)(recv.native), c_child_path)
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ConvertIterToChildIter is a wrapper around the C function gtk_tree_model_filter_convert_iter_to_child_iter.
func (recv *TreeModelFilter) ConvertIterToChildIter(filterIter *TreeIter) *TreeIter {
	var c_child_iter C.GtkTreeIter

	c_filter_iter := (*C.GtkTreeIter)(C.NULL)
	if filterIter != nil {
		c_filter_iter = (*C.GtkTreeIter)(filterIter.ToC())
	}

	C.gtk_tree_model_filter_convert_iter_to_child_iter((*C.GtkTreeModelFilter)(recv.native), &c_child_iter, c_filter_iter)

	childIter := TreeIterNewFromC(unsafe.Pointer(&c_child_iter))

	return childIter
}

// ConvertPathToChildPath is a wrapper around the C function gtk_tree_model_filter_convert_path_to_child_path.
func (recv *TreeModelFilter) ConvertPathToChildPath(filterPath *TreePath) *TreePath {
	c_filter_path := (*C.GtkTreePath)(C.NULL)
	if filterPath != nil {
		c_filter_path = (*C.GtkTreePath)(filterPath.ToC())
	}

	retC := C.gtk_tree_model_filter_convert_path_to_child_path((*C.GtkTreeModelFilter)(recv.native), c_filter_path)
	var retGo (*TreePath)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TreePathNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetModel is a wrapper around the C function gtk_tree_model_filter_get_model.
func (recv *TreeModelFilter) GetModel() *TreeModel {
	retC := C.gtk_tree_model_filter_get_model((*C.GtkTreeModelFilter)(recv.native))
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// Refilter is a wrapper around the C function gtk_tree_model_filter_refilter.
func (recv *TreeModelFilter) Refilter() {
	C.gtk_tree_model_filter_refilter((*C.GtkTreeModelFilter)(recv.native))

	return
}

// Unsupported : gtk_tree_model_filter_set_modify_func : unsupported parameter func : no type generator for TreeModelFilterModifyFunc (GtkTreeModelFilterModifyFunc) for param func

// SetVisibleColumn is a wrapper around the C function gtk_tree_model_filter_set_visible_column.
func (recv *TreeModelFilter) SetVisibleColumn(column int32) {
	c_column := (C.gint)(column)

	C.gtk_tree_model_filter_set_visible_column((*C.GtkTreeModelFilter)(recv.native), c_column)

	return
}

// Unsupported : gtk_tree_model_filter_set_visible_func : unsupported parameter func : no type generator for TreeModelFilterVisibleFunc (GtkTreeModelFilterVisibleFunc) for param func

// GetExpand is a wrapper around the C function gtk_tree_view_column_get_expand.
func (recv *TreeViewColumn) GetExpand() bool {
	retC := C.gtk_tree_view_column_get_expand((*C.GtkTreeViewColumn)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// SetExpand is a wrapper around the C function gtk_tree_view_column_set_expand.
func (recv *TreeViewColumn) SetExpand(expand bool) {
	c_expand :=
		boolToGboolean(expand)

	C.gtk_tree_view_column_set_expand((*C.GtkTreeViewColumn)(recv.native), c_expand)

	return
}

type signalUIManagerActionsChangedDetail struct {
	callback  UIManagerSignalActionsChangedCallback
	handlerID C.gulong
}

var signalUIManagerActionsChangedId int
var signalUIManagerActionsChangedMap = make(map[int]signalUIManagerActionsChangedDetail)
var signalUIManagerActionsChangedLock sync.RWMutex

// UIManagerSignalActionsChangedCallback is a callback function for a 'actions-changed' signal emitted from a UIManager.
type UIManagerSignalActionsChangedCallback func()

/*
ConnectActionsChanged connects the callback to the 'actions-changed' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectActionsChanged to remove it.
*/
func (recv *UIManager) ConnectActionsChanged(callback UIManagerSignalActionsChangedCallback) int {
	signalUIManagerActionsChangedLock.Lock()
	defer signalUIManagerActionsChangedLock.Unlock()

	signalUIManagerActionsChangedId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_actions_changed(instance, C.gpointer(uintptr(signalUIManagerActionsChangedId)))

	detail := signalUIManagerActionsChangedDetail{callback, handlerID}
	signalUIManagerActionsChangedMap[signalUIManagerActionsChangedId] = detail

	return signalUIManagerActionsChangedId
}

/*
DisconnectActionsChanged disconnects a callback from the 'actions-changed' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectActionsChanged.
*/
func (recv *UIManager) DisconnectActionsChanged(connectionID int) {
	signalUIManagerActionsChangedLock.Lock()
	defer signalUIManagerActionsChangedLock.Unlock()

	detail, exists := signalUIManagerActionsChangedMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerActionsChangedMap, connectionID)
}

//export uimanager_actionsChangedHandler
func uimanager_actionsChangedHandler(_ *C.GObject, data C.gpointer) {
	signalUIManagerActionsChangedLock.RLock()
	defer signalUIManagerActionsChangedLock.RUnlock()

	index := int(uintptr(data))
	callback := signalUIManagerActionsChangedMap[index].callback
	callback()
}

type signalUIManagerAddWidgetDetail struct {
	callback  UIManagerSignalAddWidgetCallback
	handlerID C.gulong
}

var signalUIManagerAddWidgetId int
var signalUIManagerAddWidgetMap = make(map[int]signalUIManagerAddWidgetDetail)
var signalUIManagerAddWidgetLock sync.RWMutex

// UIManagerSignalAddWidgetCallback is a callback function for a 'add-widget' signal emitted from a UIManager.
type UIManagerSignalAddWidgetCallback func(widget *Widget)

/*
ConnectAddWidget connects the callback to the 'add-widget' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectAddWidget to remove it.
*/
func (recv *UIManager) ConnectAddWidget(callback UIManagerSignalAddWidgetCallback) int {
	signalUIManagerAddWidgetLock.Lock()
	defer signalUIManagerAddWidgetLock.Unlock()

	signalUIManagerAddWidgetId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_add_widget(instance, C.gpointer(uintptr(signalUIManagerAddWidgetId)))

	detail := signalUIManagerAddWidgetDetail{callback, handlerID}
	signalUIManagerAddWidgetMap[signalUIManagerAddWidgetId] = detail

	return signalUIManagerAddWidgetId
}

/*
DisconnectAddWidget disconnects a callback from the 'add-widget' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectAddWidget.
*/
func (recv *UIManager) DisconnectAddWidget(connectionID int) {
	signalUIManagerAddWidgetLock.Lock()
	defer signalUIManagerAddWidgetLock.Unlock()

	detail, exists := signalUIManagerAddWidgetMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerAddWidgetMap, connectionID)
}

//export uimanager_addWidgetHandler
func uimanager_addWidgetHandler(_ *C.GObject, c_widget *C.GtkWidget, data C.gpointer) {
	signalUIManagerAddWidgetLock.RLock()
	defer signalUIManagerAddWidgetLock.RUnlock()

	widget := WidgetNewFromC(unsafe.Pointer(c_widget))

	index := int(uintptr(data))
	callback := signalUIManagerAddWidgetMap[index].callback
	callback(widget)
}

type signalUIManagerConnectProxyDetail struct {
	callback  UIManagerSignalConnectProxyCallback
	handlerID C.gulong
}

var signalUIManagerConnectProxyId int
var signalUIManagerConnectProxyMap = make(map[int]signalUIManagerConnectProxyDetail)
var signalUIManagerConnectProxyLock sync.RWMutex

// UIManagerSignalConnectProxyCallback is a callback function for a 'connect-proxy' signal emitted from a UIManager.
type UIManagerSignalConnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectConnectProxy connects the callback to the 'connect-proxy' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectConnectProxy to remove it.
*/
func (recv *UIManager) ConnectConnectProxy(callback UIManagerSignalConnectProxyCallback) int {
	signalUIManagerConnectProxyLock.Lock()
	defer signalUIManagerConnectProxyLock.Unlock()

	signalUIManagerConnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_connect_proxy(instance, C.gpointer(uintptr(signalUIManagerConnectProxyId)))

	detail := signalUIManagerConnectProxyDetail{callback, handlerID}
	signalUIManagerConnectProxyMap[signalUIManagerConnectProxyId] = detail

	return signalUIManagerConnectProxyId
}

/*
DisconnectConnectProxy disconnects a callback from the 'connect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectConnectProxy.
*/
func (recv *UIManager) DisconnectConnectProxy(connectionID int) {
	signalUIManagerConnectProxyLock.Lock()
	defer signalUIManagerConnectProxyLock.Unlock()

	detail, exists := signalUIManagerConnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerConnectProxyMap, connectionID)
}

//export uimanager_connectProxyHandler
func uimanager_connectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalUIManagerConnectProxyLock.RLock()
	defer signalUIManagerConnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalUIManagerConnectProxyMap[index].callback
	callback(action, proxy)
}

type signalUIManagerDisconnectProxyDetail struct {
	callback  UIManagerSignalDisconnectProxyCallback
	handlerID C.gulong
}

var signalUIManagerDisconnectProxyId int
var signalUIManagerDisconnectProxyMap = make(map[int]signalUIManagerDisconnectProxyDetail)
var signalUIManagerDisconnectProxyLock sync.RWMutex

// UIManagerSignalDisconnectProxyCallback is a callback function for a 'disconnect-proxy' signal emitted from a UIManager.
type UIManagerSignalDisconnectProxyCallback func(action *Action, proxy *Widget)

/*
ConnectDisconnectProxy connects the callback to the 'disconnect-proxy' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectDisconnectProxy to remove it.
*/
func (recv *UIManager) ConnectDisconnectProxy(callback UIManagerSignalDisconnectProxyCallback) int {
	signalUIManagerDisconnectProxyLock.Lock()
	defer signalUIManagerDisconnectProxyLock.Unlock()

	signalUIManagerDisconnectProxyId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_disconnect_proxy(instance, C.gpointer(uintptr(signalUIManagerDisconnectProxyId)))

	detail := signalUIManagerDisconnectProxyDetail{callback, handlerID}
	signalUIManagerDisconnectProxyMap[signalUIManagerDisconnectProxyId] = detail

	return signalUIManagerDisconnectProxyId
}

/*
DisconnectDisconnectProxy disconnects a callback from the 'disconnect-proxy' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectDisconnectProxy.
*/
func (recv *UIManager) DisconnectDisconnectProxy(connectionID int) {
	signalUIManagerDisconnectProxyLock.Lock()
	defer signalUIManagerDisconnectProxyLock.Unlock()

	detail, exists := signalUIManagerDisconnectProxyMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerDisconnectProxyMap, connectionID)
}

//export uimanager_disconnectProxyHandler
func uimanager_disconnectProxyHandler(_ *C.GObject, c_action *C.GtkAction, c_proxy *C.GtkWidget, data C.gpointer) {
	signalUIManagerDisconnectProxyLock.RLock()
	defer signalUIManagerDisconnectProxyLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	proxy := WidgetNewFromC(unsafe.Pointer(c_proxy))

	index := int(uintptr(data))
	callback := signalUIManagerDisconnectProxyMap[index].callback
	callback(action, proxy)
}

type signalUIManagerPostActivateDetail struct {
	callback  UIManagerSignalPostActivateCallback
	handlerID C.gulong
}

var signalUIManagerPostActivateId int
var signalUIManagerPostActivateMap = make(map[int]signalUIManagerPostActivateDetail)
var signalUIManagerPostActivateLock sync.RWMutex

// UIManagerSignalPostActivateCallback is a callback function for a 'post-activate' signal emitted from a UIManager.
type UIManagerSignalPostActivateCallback func(action *Action)

/*
ConnectPostActivate connects the callback to the 'post-activate' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectPostActivate to remove it.
*/
func (recv *UIManager) ConnectPostActivate(callback UIManagerSignalPostActivateCallback) int {
	signalUIManagerPostActivateLock.Lock()
	defer signalUIManagerPostActivateLock.Unlock()

	signalUIManagerPostActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_post_activate(instance, C.gpointer(uintptr(signalUIManagerPostActivateId)))

	detail := signalUIManagerPostActivateDetail{callback, handlerID}
	signalUIManagerPostActivateMap[signalUIManagerPostActivateId] = detail

	return signalUIManagerPostActivateId
}

/*
DisconnectPostActivate disconnects a callback from the 'post-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPostActivate.
*/
func (recv *UIManager) DisconnectPostActivate(connectionID int) {
	signalUIManagerPostActivateLock.Lock()
	defer signalUIManagerPostActivateLock.Unlock()

	detail, exists := signalUIManagerPostActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerPostActivateMap, connectionID)
}

//export uimanager_postActivateHandler
func uimanager_postActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalUIManagerPostActivateLock.RLock()
	defer signalUIManagerPostActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalUIManagerPostActivateMap[index].callback
	callback(action)
}

type signalUIManagerPreActivateDetail struct {
	callback  UIManagerSignalPreActivateCallback
	handlerID C.gulong
}

var signalUIManagerPreActivateId int
var signalUIManagerPreActivateMap = make(map[int]signalUIManagerPreActivateDetail)
var signalUIManagerPreActivateLock sync.RWMutex

// UIManagerSignalPreActivateCallback is a callback function for a 'pre-activate' signal emitted from a UIManager.
type UIManagerSignalPreActivateCallback func(action *Action)

/*
ConnectPreActivate connects the callback to the 'pre-activate' signal for the UIManager.

The returned value represents the connection, and may be passed to DisconnectPreActivate to remove it.
*/
func (recv *UIManager) ConnectPreActivate(callback UIManagerSignalPreActivateCallback) int {
	signalUIManagerPreActivateLock.Lock()
	defer signalUIManagerPreActivateLock.Unlock()

	signalUIManagerPreActivateId++
	instance := C.gpointer(recv.native)
	handlerID := C.UIManager_signal_connect_pre_activate(instance, C.gpointer(uintptr(signalUIManagerPreActivateId)))

	detail := signalUIManagerPreActivateDetail{callback, handlerID}
	signalUIManagerPreActivateMap[signalUIManagerPreActivateId] = detail

	return signalUIManagerPreActivateId
}

/*
DisconnectPreActivate disconnects a callback from the 'pre-activate' signal for the UIManager.

The connectionID should be a value returned from a call to ConnectPreActivate.
*/
func (recv *UIManager) DisconnectPreActivate(connectionID int) {
	signalUIManagerPreActivateLock.Lock()
	defer signalUIManagerPreActivateLock.Unlock()

	detail, exists := signalUIManagerPreActivateMap[connectionID]
	if !exists {
		return
	}

	instance := C.gpointer(recv.native)
	C.g_signal_handler_disconnect(instance, detail.handlerID)
	delete(signalUIManagerPreActivateMap, connectionID)
}

//export uimanager_preActivateHandler
func uimanager_preActivateHandler(_ *C.GObject, c_action *C.GtkAction, data C.gpointer) {
	signalUIManagerPreActivateLock.RLock()
	defer signalUIManagerPreActivateLock.RUnlock()

	action := ActionNewFromC(unsafe.Pointer(c_action))

	index := int(uintptr(data))
	callback := signalUIManagerPreActivateMap[index].callback
	callback(action)
}

// UIManagerNew is a wrapper around the C function gtk_ui_manager_new.
func UIManagerNew() *UIManager {
	retC := C.gtk_ui_manager_new()
	retGo := UIManagerNewFromC(unsafe.Pointer(retC))

	if retC != nil {
		C.g_object_unref((C.gpointer)(retC))
	}

	return retGo
}

// AddUi is a wrapper around the C function gtk_ui_manager_add_ui.
func (recv *UIManager) AddUi(mergeId uint32, path string, name string, action string, type_ UIManagerItemType, top bool) {
	c_merge_id := (C.guint)(mergeId)

	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_action := C.CString(action)
	defer C.free(unsafe.Pointer(c_action))

	c_type := (C.GtkUIManagerItemType)(type_)

	c_top :=
		boolToGboolean(top)

	C.gtk_ui_manager_add_ui((*C.GtkUIManager)(recv.native), c_merge_id, c_path, c_name, c_action, c_type, c_top)

	return
}

// AddUiFromFile is a wrapper around the C function gtk_ui_manager_add_ui_from_file.
func (recv *UIManager) AddUiFromFile(filename string) (uint32, error) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_file((*C.GtkUIManager)(recv.native), c_filename, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddUiFromString is a wrapper around the C function gtk_ui_manager_add_ui_from_string.
func (recv *UIManager) AddUiFromString(buffer string) (uint32, error) {
	c_buffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(c_buffer))

	c_length := (C.gssize)(len(buffer))

	var cThrowableError *C.GError

	retC := C.gtk_ui_manager_add_ui_from_string((*C.GtkUIManager)(recv.native), c_buffer, c_length, &cThrowableError)
	retGo := (uint32)(retC)

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// EnsureUpdate is a wrapper around the C function gtk_ui_manager_ensure_update.
func (recv *UIManager) EnsureUpdate() {
	C.gtk_ui_manager_ensure_update((*C.GtkUIManager)(recv.native))

	return
}

// GetAccelGroup is a wrapper around the C function gtk_ui_manager_get_accel_group.
func (recv *UIManager) GetAccelGroup() *AccelGroup {
	retC := C.gtk_ui_manager_get_accel_group((*C.GtkUIManager)(recv.native))
	retGo := AccelGroupNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAction is a wrapper around the C function gtk_ui_manager_get_action.
func (recv *UIManager) GetAction(path string) *Action {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_action((*C.GtkUIManager)(recv.native), c_path)
	retGo := ActionNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetActionGroups is a wrapper around the C function gtk_ui_manager_get_action_groups.
func (recv *UIManager) GetActionGroups() *glib.List {
	retC := C.gtk_ui_manager_get_action_groups((*C.GtkUIManager)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetAddTearoffs is a wrapper around the C function gtk_ui_manager_get_add_tearoffs.
func (recv *UIManager) GetAddTearoffs() bool {
	retC := C.gtk_ui_manager_get_add_tearoffs((*C.GtkUIManager)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetToplevels is a wrapper around the C function gtk_ui_manager_get_toplevels.
func (recv *UIManager) GetToplevels(types UIManagerItemType) *glib.SList {
	c_types := (C.GtkUIManagerItemType)(types)

	retC := C.gtk_ui_manager_get_toplevels((*C.GtkUIManager)(recv.native), c_types)
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetUi is a wrapper around the C function gtk_ui_manager_get_ui.
func (recv *UIManager) GetUi() string {
	retC := C.gtk_ui_manager_get_ui((*C.GtkUIManager)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetWidget is a wrapper around the C function gtk_ui_manager_get_widget.
func (recv *UIManager) GetWidget(path string) *Widget {
	c_path := C.CString(path)
	defer C.free(unsafe.Pointer(c_path))

	retC := C.gtk_ui_manager_get_widget((*C.GtkUIManager)(recv.native), c_path)
	retGo := WidgetNewFromC(unsafe.Pointer(retC))

	return retGo
}

// InsertActionGroup is a wrapper around the C function gtk_ui_manager_insert_action_group.
func (recv *UIManager) InsertActionGroup(actionGroup *ActionGroup, pos int32) {
	c_action_group := (*C.GtkActionGroup)(C.NULL)
	if actionGroup != nil {
		c_action_group = (*C.GtkActionGroup)(actionGroup.ToC())
	}

	c_pos := (C.gint)(pos)

	C.gtk_ui_manager_insert_action_group((*C.GtkUIManager)(recv.native), c_action_group, c_pos)

	return
}

// NewMergeId is a wrapper around the C function gtk_ui_manager_new_merge_id.
func (recv *UIManager) NewMergeId() uint32 {
	retC := C.gtk_ui_manager_new_merge_id((*C.GtkUIManager)(recv.native))
	retGo := (uint32)(retC)

	return retGo
}

// RemoveActionGroup is a wrapper around the C function gtk_ui_manager_remove_action_group.
func (recv *UIManager) RemoveActionGroup(actionGroup *ActionGroup) {
	c_action_group := (*C.GtkActionGroup)(C.NULL)
	if actionGroup != nil {
		c_action_group = (*C.GtkActionGroup)(actionGroup.ToC())
	}

	C.gtk_ui_manager_remove_action_group((*C.GtkUIManager)(recv.native), c_action_group)

	return
}

// RemoveUi is a wrapper around the C function gtk_ui_manager_remove_ui.
func (recv *UIManager) RemoveUi(mergeId uint32) {
	c_merge_id := (C.guint)(mergeId)

	C.gtk_ui_manager_remove_ui((*C.GtkUIManager)(recv.native), c_merge_id)

	return
}

// SetAddTearoffs is a wrapper around the C function gtk_ui_manager_set_add_tearoffs.
func (recv *UIManager) SetAddTearoffs(addTearoffs bool) {
	c_add_tearoffs :=
		boolToGboolean(addTearoffs)

	C.gtk_ui_manager_set_add_tearoffs((*C.GtkUIManager)(recv.native), c_add_tearoffs)

	return
}

// AddMnemonicLabel is a wrapper around the C function gtk_widget_add_mnemonic_label.
func (recv *Widget) AddMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(C.NULL)
	if label != nil {
		c_label = (*C.GtkWidget)(label.ToC())
	}

	C.gtk_widget_add_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// CanActivateAccel is a wrapper around the C function gtk_widget_can_activate_accel.
func (recv *Widget) CanActivateAccel(signalId uint32) bool {
	c_signal_id := (C.guint)(signalId)

	retC := C.gtk_widget_can_activate_accel((*C.GtkWidget)(recv.native), c_signal_id)
	retGo := retC == C.TRUE

	return retGo
}

// DragSourceGetTargetList is a wrapper around the C function gtk_drag_source_get_target_list.
func (recv *Widget) DragSourceGetTargetList() *TargetList {
	retC := C.gtk_drag_source_get_target_list((*C.GtkWidget)(recv.native))
	var retGo (*TargetList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = TargetListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// DragSourceSetTargetList is a wrapper around the C function gtk_drag_source_set_target_list.
func (recv *Widget) DragSourceSetTargetList(targetList *TargetList) {
	c_target_list := (*C.GtkTargetList)(C.NULL)
	if targetList != nil {
		c_target_list = (*C.GtkTargetList)(targetList.ToC())
	}

	C.gtk_drag_source_set_target_list((*C.GtkWidget)(recv.native), c_target_list)

	return
}

// GetNoShowAll is a wrapper around the C function gtk_widget_get_no_show_all.
func (recv *Widget) GetNoShowAll() bool {
	retC := C.gtk_widget_get_no_show_all((*C.GtkWidget)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ListMnemonicLabels is a wrapper around the C function gtk_widget_list_mnemonic_labels.
func (recv *Widget) ListMnemonicLabels() *glib.List {
	retC := C.gtk_widget_list_mnemonic_labels((*C.GtkWidget)(recv.native))
	retGo := glib.ListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// QueueResizeNoRedraw is a wrapper around the C function gtk_widget_queue_resize_no_redraw.
func (recv *Widget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw((*C.GtkWidget)(recv.native))

	return
}

// RemoveMnemonicLabel is a wrapper around the C function gtk_widget_remove_mnemonic_label.
func (recv *Widget) RemoveMnemonicLabel(label *Widget) {
	c_label := (*C.GtkWidget)(C.NULL)
	if label != nil {
		c_label = (*C.GtkWidget)(label.ToC())
	}

	C.gtk_widget_remove_mnemonic_label((*C.GtkWidget)(recv.native), c_label)

	return
}

// SetNoShowAll is a wrapper around the C function gtk_widget_set_no_show_all.
func (recv *Widget) SetNoShowAll(noShowAll bool) {
	c_no_show_all :=
		boolToGboolean(noShowAll)

	C.gtk_widget_set_no_show_all((*C.GtkWidget)(recv.native), c_no_show_all)

	return
}

// WindowSetDefaultIcon is a wrapper around the C function gtk_window_set_default_icon.
func WindowSetDefaultIcon(icon *gdkpixbuf.Pixbuf) {
	c_icon := (*C.GdkPixbuf)(C.NULL)
	if icon != nil {
		c_icon = (*C.GdkPixbuf)(icon.ToC())
	}

	C.gtk_window_set_default_icon(c_icon)

	return
}

// ActivateKey is a wrapper around the C function gtk_window_activate_key.
func (recv *Window) ActivateKey(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_window_activate_key((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// GetAcceptFocus is a wrapper around the C function gtk_window_get_accept_focus.
func (recv *Window) GetAcceptFocus() bool {
	retC := C.gtk_window_get_accept_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// HasToplevelFocus is a wrapper around the C function gtk_window_has_toplevel_focus.
func (recv *Window) HasToplevelFocus() bool {
	retC := C.gtk_window_has_toplevel_focus((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// IsActive is a wrapper around the C function gtk_window_is_active.
func (recv *Window) IsActive() bool {
	retC := C.gtk_window_is_active((*C.GtkWindow)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// PropagateKeyEvent is a wrapper around the C function gtk_window_propagate_key_event.
func (recv *Window) PropagateKeyEvent(event *gdk.EventKey) bool {
	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_window_propagate_key_event((*C.GtkWindow)(recv.native), c_event)
	retGo := retC == C.TRUE

	return retGo
}

// SetAcceptFocus is a wrapper around the C function gtk_window_set_accept_focus.
func (recv *Window) SetAcceptFocus(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_accept_focus((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetKeepAbove is a wrapper around the C function gtk_window_set_keep_above.
func (recv *Window) SetKeepAbove(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_above((*C.GtkWindow)(recv.native), c_setting)

	return
}

// SetKeepBelow is a wrapper around the C function gtk_window_set_keep_below.
func (recv *Window) SetKeepBelow(setting bool) {
	c_setting :=
		boolToGboolean(setting)

	C.gtk_window_set_keep_below((*C.GtkWindow)(recv.native), c_setting)

	return
}

// Blacklisted : STOCK_DIALOG_AUTHENTICATION

// Blacklisted : STOCK_HARDDISK

// Blacklisted : STOCK_INDENT

// Blacklisted : STOCK_NETWORK

// Blacklisted : STOCK_UNINDENT

// BindingsActivateEvent is a wrapper around the C function gtk_bindings_activate_event.
func BindingsActivateEvent(object *gobject.Object, event *gdk.EventKey) bool {
	c_object := (*C.GObject)(C.NULL)
	if object != nil {
		c_object = (*C.GObject)(object.ToC())
	}

	c_event := (*C.GdkEventKey)(C.NULL)
	if event != nil {
		c_event = (*C.GdkEventKey)(event.ToC())
	}

	retC := C.gtk_bindings_activate_event(c_object, c_event)
	retGo := retC == C.TRUE

	return retGo
}

// RcResetStyles is a wrapper around the C function gtk_rc_reset_styles.
func RcResetStyles(settings *Settings) {
	c_settings := (*C.GtkSettings)(C.NULL)
	if settings != nil {
		c_settings = (*C.GtkSettings)(settings.ToC())
	}

	C.gtk_rc_reset_styles(c_settings)

	return
}

// AddAttribute is a wrapper around the C function gtk_cell_layout_add_attribute.
func (recv *CellLayout) AddAttribute(cell *CellRenderer, attribute string, column int32) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_attribute := C.CString(attribute)
	defer C.free(unsafe.Pointer(c_attribute))

	c_column := (C.gint)(column)

	C.gtk_cell_layout_add_attribute((*C.GtkCellLayout)(recv.native), c_cell, c_attribute, c_column)

	return
}

// Clear is a wrapper around the C function gtk_cell_layout_clear.
func (recv *CellLayout) Clear() {
	C.gtk_cell_layout_clear((*C.GtkCellLayout)(recv.native))

	return
}

// ClearAttributes is a wrapper around the C function gtk_cell_layout_clear_attributes.
func (recv *CellLayout) ClearAttributes(cell *CellRenderer) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	C.gtk_cell_layout_clear_attributes((*C.GtkCellLayout)(recv.native), c_cell)

	return
}

// PackEnd is a wrapper around the C function gtk_cell_layout_pack_end.
func (recv *CellLayout) PackEnd(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	C.gtk_cell_layout_pack_end((*C.GtkCellLayout)(recv.native), c_cell, c_expand)

	return
}

// PackStart is a wrapper around the C function gtk_cell_layout_pack_start.
func (recv *CellLayout) PackStart(cell *CellRenderer, expand bool) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_expand :=
		boolToGboolean(expand)

	C.gtk_cell_layout_pack_start((*C.GtkCellLayout)(recv.native), c_cell, c_expand)

	return
}

// Reorder is a wrapper around the C function gtk_cell_layout_reorder.
func (recv *CellLayout) Reorder(cell *CellRenderer, position int32) {
	c_cell := (*C.GtkCellRenderer)(C.NULL)
	if cell != nil {
		c_cell = (*C.GtkCellRenderer)(cell.ToC())
	}

	c_position := (C.gint)(position)

	C.gtk_cell_layout_reorder((*C.GtkCellLayout)(recv.native), c_cell, c_position)

	return
}

// Unsupported : gtk_cell_layout_set_attributes : unsupported parameter ... : varargs

// Unsupported : gtk_cell_layout_set_cell_data_func : unsupported parameter func : no type generator for CellLayoutDataFunc (GtkCellLayoutDataFunc) for param func

// AddFilter is a wrapper around the C function gtk_file_chooser_add_filter.
func (recv *FileChooser) AddFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_add_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// AddShortcutFolder is a wrapper around the C function gtk_file_chooser_add_shortcut_folder.
func (recv *FileChooser) AddShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// AddShortcutFolderUri is a wrapper around the C function gtk_file_chooser_add_shortcut_folder_uri.
func (recv *FileChooser) AddShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_add_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// GetAction is a wrapper around the C function gtk_file_chooser_get_action.
func (recv *FileChooser) GetAction() FileChooserAction {
	retC := C.gtk_file_chooser_get_action((*C.GtkFileChooser)(recv.native))
	retGo := (FileChooserAction)(retC)

	return retGo
}

// GetCurrentFolder is a wrapper around the C function gtk_file_chooser_get_current_folder.
func (recv *FileChooser) GetCurrentFolder() string {
	retC := C.gtk_file_chooser_get_current_folder((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetCurrentFolderUri is a wrapper around the C function gtk_file_chooser_get_current_folder_uri.
func (recv *FileChooser) GetCurrentFolderUri() string {
	retC := C.gtk_file_chooser_get_current_folder_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetExtraWidget is a wrapper around the C function gtk_file_chooser_get_extra_widget.
func (recv *FileChooser) GetExtraWidget() *Widget {
	retC := C.gtk_file_chooser_get_extra_widget((*C.GtkFileChooser)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetFilename is a wrapper around the C function gtk_file_chooser_get_filename.
func (recv *FileChooser) GetFilename() string {
	retC := C.gtk_file_chooser_get_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetFilenames is a wrapper around the C function gtk_file_chooser_get_filenames.
func (recv *FileChooser) GetFilenames() *glib.SList {
	retC := C.gtk_file_chooser_get_filenames((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// GetFilter is a wrapper around the C function gtk_file_chooser_get_filter.
func (recv *FileChooser) GetFilter() *FileFilter {
	retC := C.gtk_file_chooser_get_filter((*C.GtkFileChooser)(recv.native))
	var retGo (*FileFilter)
	if retC == nil {
		retGo = nil
	} else {
		retGo = FileFilterNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetLocalOnly is a wrapper around the C function gtk_file_chooser_get_local_only.
func (recv *FileChooser) GetLocalOnly() bool {
	retC := C.gtk_file_chooser_get_local_only((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetPreviewFilename is a wrapper around the C function gtk_file_chooser_get_preview_filename.
func (recv *FileChooser) GetPreviewFilename() string {
	retC := C.gtk_file_chooser_get_preview_filename((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewUri is a wrapper around the C function gtk_file_chooser_get_preview_uri.
func (recv *FileChooser) GetPreviewUri() string {
	retC := C.gtk_file_chooser_get_preview_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetPreviewWidget is a wrapper around the C function gtk_file_chooser_get_preview_widget.
func (recv *FileChooser) GetPreviewWidget() *Widget {
	retC := C.gtk_file_chooser_get_preview_widget((*C.GtkFileChooser)(recv.native))
	var retGo (*Widget)
	if retC == nil {
		retGo = nil
	} else {
		retGo = WidgetNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// GetPreviewWidgetActive is a wrapper around the C function gtk_file_chooser_get_preview_widget_active.
func (recv *FileChooser) GetPreviewWidgetActive() bool {
	retC := C.gtk_file_chooser_get_preview_widget_active((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetSelectMultiple is a wrapper around the C function gtk_file_chooser_get_select_multiple.
func (recv *FileChooser) GetSelectMultiple() bool {
	retC := C.gtk_file_chooser_get_select_multiple((*C.GtkFileChooser)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// GetUri is a wrapper around the C function gtk_file_chooser_get_uri.
func (recv *FileChooser) GetUri() string {
	retC := C.gtk_file_chooser_get_uri((*C.GtkFileChooser)(recv.native))
	retGo := C.GoString(retC)
	defer C.free(unsafe.Pointer(retC))

	return retGo
}

// GetUris is a wrapper around the C function gtk_file_chooser_get_uris.
func (recv *FileChooser) GetUris() *glib.SList {
	retC := C.gtk_file_chooser_get_uris((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListFilters is a wrapper around the C function gtk_file_chooser_list_filters.
func (recv *FileChooser) ListFilters() *glib.SList {
	retC := C.gtk_file_chooser_list_filters((*C.GtkFileChooser)(recv.native))
	retGo := glib.SListNewFromC(unsafe.Pointer(retC))

	return retGo
}

// ListShortcutFolderUris is a wrapper around the C function gtk_file_chooser_list_shortcut_folder_uris.
func (recv *FileChooser) ListShortcutFolderUris() *glib.SList {
	retC := C.gtk_file_chooser_list_shortcut_folder_uris((*C.GtkFileChooser)(recv.native))
	var retGo (*glib.SList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.SListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// ListShortcutFolders is a wrapper around the C function gtk_file_chooser_list_shortcut_folders.
func (recv *FileChooser) ListShortcutFolders() *glib.SList {
	retC := C.gtk_file_chooser_list_shortcut_folders((*C.GtkFileChooser)(recv.native))
	var retGo (*glib.SList)
	if retC == nil {
		retGo = nil
	} else {
		retGo = glib.SListNewFromC(unsafe.Pointer(retC))
	}

	return retGo
}

// RemoveFilter is a wrapper around the C function gtk_file_chooser_remove_filter.
func (recv *FileChooser) RemoveFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_remove_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// RemoveShortcutFolder is a wrapper around the C function gtk_file_chooser_remove_shortcut_folder.
func (recv *FileChooser) RemoveShortcutFolder(folder string) (bool, error) {
	c_folder := C.CString(folder)
	defer C.free(unsafe.Pointer(c_folder))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder((*C.GtkFileChooser)(recv.native), c_folder, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// RemoveShortcutFolderUri is a wrapper around the C function gtk_file_chooser_remove_shortcut_folder_uri.
func (recv *FileChooser) RemoveShortcutFolderUri(uri string) (bool, error) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	var cThrowableError *C.GError

	retC := C.gtk_file_chooser_remove_shortcut_folder_uri((*C.GtkFileChooser)(recv.native), c_uri, &cThrowableError)
	retGo := retC == C.TRUE

	var goError error = nil
	if cThrowableError != nil {
		goThrowableError := glib.ErrorNewFromC(unsafe.Pointer(cThrowableError))
		goError = goThrowableError

		C.g_error_free(cThrowableError)
	}

	return retGo, goError
}

// SelectAll is a wrapper around the C function gtk_file_chooser_select_all.
func (recv *FileChooser) SelectAll() {
	C.gtk_file_chooser_select_all((*C.GtkFileChooser)(recv.native))

	return
}

// SelectFilename is a wrapper around the C function gtk_file_chooser_select_filename.
func (recv *FileChooser) SelectFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_select_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SelectUri is a wrapper around the C function gtk_file_chooser_select_uri.
func (recv *FileChooser) SelectUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_select_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetAction is a wrapper around the C function gtk_file_chooser_set_action.
func (recv *FileChooser) SetAction(action FileChooserAction) {
	c_action := (C.GtkFileChooserAction)(action)

	C.gtk_file_chooser_set_action((*C.GtkFileChooser)(recv.native), c_action)

	return
}

// SetCurrentFolder is a wrapper around the C function gtk_file_chooser_set_current_folder.
func (recv *FileChooser) SetCurrentFolder(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_current_folder((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SetCurrentFolderUri is a wrapper around the C function gtk_file_chooser_set_current_folder_uri.
func (recv *FileChooser) SetCurrentFolderUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_current_folder_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetCurrentName is a wrapper around the C function gtk_file_chooser_set_current_name.
func (recv *FileChooser) SetCurrentName(name string) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	C.gtk_file_chooser_set_current_name((*C.GtkFileChooser)(recv.native), c_name)

	return
}

// SetExtraWidget is a wrapper around the C function gtk_file_chooser_set_extra_widget.
func (recv *FileChooser) SetExtraWidget(extraWidget *Widget) {
	c_extra_widget := (*C.GtkWidget)(C.NULL)
	if extraWidget != nil {
		c_extra_widget = (*C.GtkWidget)(extraWidget.ToC())
	}

	C.gtk_file_chooser_set_extra_widget((*C.GtkFileChooser)(recv.native), c_extra_widget)

	return
}

// SetFilename is a wrapper around the C function gtk_file_chooser_set_filename.
func (recv *FileChooser) SetFilename(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	retC := C.gtk_file_chooser_set_filename((*C.GtkFileChooser)(recv.native), c_filename)
	retGo := retC == C.TRUE

	return retGo
}

// SetFilter is a wrapper around the C function gtk_file_chooser_set_filter.
func (recv *FileChooser) SetFilter(filter *FileFilter) {
	c_filter := (*C.GtkFileFilter)(C.NULL)
	if filter != nil {
		c_filter = (*C.GtkFileFilter)(filter.ToC())
	}

	C.gtk_file_chooser_set_filter((*C.GtkFileChooser)(recv.native), c_filter)

	return
}

// SetLocalOnly is a wrapper around the C function gtk_file_chooser_set_local_only.
func (recv *FileChooser) SetLocalOnly(localOnly bool) {
	c_local_only :=
		boolToGboolean(localOnly)

	C.gtk_file_chooser_set_local_only((*C.GtkFileChooser)(recv.native), c_local_only)

	return
}

// SetPreviewWidget is a wrapper around the C function gtk_file_chooser_set_preview_widget.
func (recv *FileChooser) SetPreviewWidget(previewWidget *Widget) {
	c_preview_widget := (*C.GtkWidget)(C.NULL)
	if previewWidget != nil {
		c_preview_widget = (*C.GtkWidget)(previewWidget.ToC())
	}

	C.gtk_file_chooser_set_preview_widget((*C.GtkFileChooser)(recv.native), c_preview_widget)

	return
}

// SetPreviewWidgetActive is a wrapper around the C function gtk_file_chooser_set_preview_widget_active.
func (recv *FileChooser) SetPreviewWidgetActive(active bool) {
	c_active :=
		boolToGboolean(active)

	C.gtk_file_chooser_set_preview_widget_active((*C.GtkFileChooser)(recv.native), c_active)

	return
}

// SetSelectMultiple is a wrapper around the C function gtk_file_chooser_set_select_multiple.
func (recv *FileChooser) SetSelectMultiple(selectMultiple bool) {
	c_select_multiple :=
		boolToGboolean(selectMultiple)

	C.gtk_file_chooser_set_select_multiple((*C.GtkFileChooser)(recv.native), c_select_multiple)

	return
}

// SetUri is a wrapper around the C function gtk_file_chooser_set_uri.
func (recv *FileChooser) SetUri(uri string) bool {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	retC := C.gtk_file_chooser_set_uri((*C.GtkFileChooser)(recv.native), c_uri)
	retGo := retC == C.TRUE

	return retGo
}

// SetUsePreviewLabel is a wrapper around the C function gtk_file_chooser_set_use_preview_label.
func (recv *FileChooser) SetUsePreviewLabel(useLabel bool) {
	c_use_label :=
		boolToGboolean(useLabel)

	C.gtk_file_chooser_set_use_preview_label((*C.GtkFileChooser)(recv.native), c_use_label)

	return
}

// UnselectAll is a wrapper around the C function gtk_file_chooser_unselect_all.
func (recv *FileChooser) UnselectAll() {
	C.gtk_file_chooser_unselect_all((*C.GtkFileChooser)(recv.native))

	return
}

// UnselectFilename is a wrapper around the C function gtk_file_chooser_unselect_filename.
func (recv *FileChooser) UnselectFilename(filename string) {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))

	C.gtk_file_chooser_unselect_filename((*C.GtkFileChooser)(recv.native), c_filename)

	return
}

// UnselectUri is a wrapper around the C function gtk_file_chooser_unselect_uri.
func (recv *FileChooser) UnselectUri(uri string) {
	c_uri := C.CString(uri)
	defer C.free(unsafe.Pointer(c_uri))

	C.gtk_file_chooser_unselect_uri((*C.GtkFileChooser)(recv.native), c_uri)

	return
}

// FilterNew is a wrapper around the C function gtk_tree_model_filter_new.
func (recv *TreeModel) FilterNew(root *TreePath) *TreeModel {
	c_root := (*C.GtkTreePath)(C.NULL)
	if root != nil {
		c_root = (*C.GtkTreePath)(root.ToC())
	}

	retC := C.gtk_tree_model_filter_new((*C.GtkTreeModel)(recv.native), c_root)
	retGo := TreeModelNewFromC(unsafe.Pointer(retC))

	return retGo
}

// BackwardVisibleCursorPosition is a wrapper around the C function gtk_text_iter_backward_visible_cursor_position.
func (recv *TextIter) BackwardVisibleCursorPosition() bool {
	retC := C.gtk_text_iter_backward_visible_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleCursorPositions is a wrapper around the C function gtk_text_iter_backward_visible_cursor_positions.
func (recv *TextIter) BackwardVisibleCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleWordStart is a wrapper around the C function gtk_text_iter_backward_visible_word_start.
func (recv *TextIter) BackwardVisibleWordStart() bool {
	retC := C.gtk_text_iter_backward_visible_word_start((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// BackwardVisibleWordStarts is a wrapper around the C function gtk_text_iter_backward_visible_word_starts.
func (recv *TextIter) BackwardVisibleWordStarts(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_backward_visible_word_starts((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleCursorPosition is a wrapper around the C function gtk_text_iter_forward_visible_cursor_position.
func (recv *TextIter) ForwardVisibleCursorPosition() bool {
	retC := C.gtk_text_iter_forward_visible_cursor_position((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleCursorPositions is a wrapper around the C function gtk_text_iter_forward_visible_cursor_positions.
func (recv *TextIter) ForwardVisibleCursorPositions(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_cursor_positions((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleWordEnd is a wrapper around the C function gtk_text_iter_forward_visible_word_end.
func (recv *TextIter) ForwardVisibleWordEnd() bool {
	retC := C.gtk_text_iter_forward_visible_word_end((*C.GtkTextIter)(recv.native))
	retGo := retC == C.TRUE

	return retGo
}

// ForwardVisibleWordEnds is a wrapper around the C function gtk_text_iter_forward_visible_word_ends.
func (recv *TextIter) ForwardVisibleWordEnds(count int32) bool {
	c_count := (C.gint)(count)

	retC := C.gtk_text_iter_forward_visible_word_ends((*C.GtkTextIter)(recv.native), c_count)
	retGo := retC == C.TRUE

	return retGo
}