package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	List    list.Model
	Spinner spinner.Model
	Serial  viewport.Model
	Width   int
	Height  int
	Focused int
}

func NewModel() model {
	elementspiner := spinner.New()
	elementspiner.Spinner = spinner.Dot

	lista := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	lista.Title = "Placas"

	return model{
		Spinner: elementspiner,
		List:    lista,
		Focused: 2,
	}
}

func (m model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.Focused = (m.Focused + 1) % 4
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}

	if m.Focused == 2 {
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	m.Spinner, cmd = m.Spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
