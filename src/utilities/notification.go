package utilities

import (
	"unsafe"
)

type NotificationCenter struct {
	listeners map[string]*EventChain
}
type EventCallback func(*Event)

type EventChain struct {
	chs       []chan *Event
	callbacks []*EventCallback
}

type Event struct {
	eventName string
	Params    map[string]interface{}
}

func CreateEventChain() *EventChain {
	return &EventChain{chs: []chan *Event{}, callbacks: []*EventCallback{}}
}

func CreateEvent(eventName string, params map[string]interface{}) *Event {
	return &Event{eventName: eventName, Params: params}
}

var _instance *NotificationCenter

func ShareNotificationCenter() *NotificationCenter {
	if _instance == nil {
		_instance = &NotificationCenter{}
		_instance.Init()
	}
	return _instance
}

func (n *NotificationCenter) Init() {
	n.listeners = make(map[string]*EventChain)
}

func (n *NotificationCenter) AddEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := n.listeners[eventName]
	if !ok {
		eventChain = CreateEventChain()
		n.listeners[eventName] = eventChain
	}
	exist := false

	for _, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(item))
		if a == b {
			exist = true
			break
		}
	}

	if exist {
		return
	}

	ch := make(chan *Event)
	eventChain.chs = append(eventChain.chs[:], ch)
	eventChain.callbacks = append(eventChain.callbacks[:], callback)

	go n.handle(eventName, ch, callback)
}

func (n *NotificationCenter) handle(eventName string, ch chan *Event, callback *EventCallback) {
	for {
		event := <-ch
		if event == nil {
			break
		}
		go (*callback)(event)
	}
}

func (n *NotificationCenter) RemoveEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := n.listeners[eventName]
	if !ok {
		return
	}
	var ch chan *Event
	exist := false
	key := 0
	for k, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(callback))
		if a == b {
			exist = true
			ch = eventChain.chs[k]
			key = k
			break
		}
	}
	if exist {
		ch <- nil
		eventChain.chs = append(eventChain.chs[:key], eventChain.chs[key+1:]...)
		eventChain.callbacks = append(eventChain.callbacks[:key], eventChain.callbacks[key+1:]...)
	}
}

func (n *NotificationCenter) PostEvent(event *Event) {
	eventChain, ok := n.listeners[event.eventName]
	if ok {
		for _, chEvent := range eventChain.chs {
			chEvent <- event
		}
	}
}

func (n *NotificationCenter) PosetEventWithNameAndInfos(name string, params map[string]interface{}) {
	event := &Event{name, params}
	n.PostEvent(event)
}
