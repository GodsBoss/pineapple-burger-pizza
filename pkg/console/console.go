// +build: js,wasm

package console

import (
	"syscall/js"
)

// TODO: Move this to gggg.

type Console struct {
	value js.Value
}

func Global() *Console {
	return &Console{
		value: js.Global().Get("console"),
	}
}

func (console *Console) LogObjects(objects ...interface{}) {
	console.value.Call("log", objects...)
}

func (console *Console) LogMessage(msg string, substitutions ...interface{}) {
	console.LogObjects(append([]interface{}{msg}, substitutions...)...)
}
