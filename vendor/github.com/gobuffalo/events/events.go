package events

import (
	"fmt"
	"runtime"

	"github.com/gobuffalo/mapi"
	"github.com/pkg/errors"
)

type Payload = mapi.Mapi

const (
	// ErrGeneral is emitted for general errors
	ErrGeneral = "general:err"
	// ErrPanic is emitted when a panic is recovered
	ErrPanic = "panic:err"
)

// Emit an event to all listeners
func Emit(e Event) error {
	return boss.Emit(e)
}

func EmitPayload(kind string, payload interface{}) error {
	return EmitError(kind, nil, payload)
}

func EmitError(kind string, err error, payload interface{}) error {
	var pl Payload
	pl, ok := payload.(Payload)
	if !ok {
		pl = Payload{
			"data": payload,
		}
	}
	e := Event{
		Kind:    kind,
		Payload: pl,
		Error:   err,
	}
	return Emit(e)
}

// NamedListen for events. Name is the name of the
// listener NOT the events you want to listen for,
// so something like "my-listener", "kafka-listener", etc...
func NamedListen(name string, l Listener) (DeleteFn, error) {
	return boss.Listen(name, l)
}

// Listen for events.
func Listen(l Listener) (DeleteFn, error) {
	_, file, line, _ := runtime.Caller(1)
	return NamedListen(fmt.Sprintf("%s:%d", file, line), l)
}

type listable interface {
	List() ([]string, error)
}

// List all listeners
func List() ([]string, error) {
	if l, ok := boss.(listable); ok {
		return l.List()
	}
	return []string{}, errors.Errorf("manager %T does not implemented listable", boss)
}
