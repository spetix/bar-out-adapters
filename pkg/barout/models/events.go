package models

import (
	"os"
)

// EventHdlr is a function type that represents an event handler. It takes no arguments and returns an error.
type EventHdlr func() error

// EventType is a string type that represents the type of an event.
type EventType string

type Event interface {
	Type() EventType
	Name() string
	Signal() os.Signal
	IsSignal() bool
}

// EventHandlers is a map that associates EventType keys with EventHdlr values. It is used to store event handlers for different event types.
type EventHandlers map[Event]EventHdlr

// EventStrategy is an interface that defines a strategy for executing events. It has a single method, Execute, which is responsible for executing the event handling logic.
type EventStrategy interface {
	Register(k Event, h EventHdlr)
	Execute() error
}

const (
	LeftButton   EventType = "1"
	MiddleButton EventType = "2"
	RightButton  EventType = "3"
	ScrollUp     EventType = "4"
	ScrollDown   EventType = "5"
	Signal       EventType = "999"
)

func (e EventType) String() string {
	switch e {
	case LeftButton:
		return "Left Button"
	case MiddleButton:
		return "Middle Button"
	case RightButton:
		return "Right Button"
	case ScrollUp:
		return "Scroll Up"
	case ScrollDown:
		return "Scroll Down"
	default:
		return "Unknown event"
	}
}

// EventManager is an interface that defines methods for managing events. It allows registering event handlers and running the event handling logic.
type EventManager interface {
	Register(tk Event, hdlr EventHdlr)
	Run() error
}
