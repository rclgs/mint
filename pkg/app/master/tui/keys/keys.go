package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type global struct {
	Filter key.Binding
	Quit   key.Binding
	Help   key.Binding
	Back   key.Binding
	CtrlC  key.Binding
}

var Global = global{
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp(`/`, "filter"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	CtrlC: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

type home struct {
	Images key.Binding
	Debug  key.Binding
}

type debug struct {
	LoadDebuggableContainers key.Binding
	ChangeRuntime            key.Binding
	StartSession             key.Binding
}

var Home = home{
	Images: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "Open images view"),
	),
	Debug: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "Open debug view"),
	),
}

var Debug = debug{
	LoadDebuggableContainers: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "Load debuggable containers"),
	),
	ChangeRuntime: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Change runtime"),
	),
	StartSession: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Start session"),
	),
}
