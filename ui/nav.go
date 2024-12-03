package ui

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Fragments  []Fragment
	Current    int
	Input      string
	Choosing   bool
	WindowSize tea.WindowSizeMsg
}

func NewModel(fragments []Fragment) Model {
	return Model{
		Fragments:  fragments,
		Current:    0,
		Input:      "",
		Choosing:   true,
		WindowSize: tea.WindowSizeMsg{Width: 80, Height: 24},
	}
}

func (m Model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Choosing {
			switch msg.String() {
			case "enter":
				if idx, ok := parseFragmentInput(m.Input, len(m.Fragments)); ok {
					m.Current = idx
					m.Choosing = false
					m.Input = ""
					return m, tea.ClearScreen
				} else {
					m.Input = ""
				}
			case "q":
				return m, tea.Quit
			case "backspace":
				if len(m.Input) > 0 {
					m.Input = m.Input[:len(m.Input)-1]
				}
			default:
				if isDigit(msg.String()) {
					m.Input += msg.String()
				}
			}
		} else {
			switch msg.String() {
			case "n":
				if m.Current < len(m.Fragments)-1 {
					m.Current++
				}
			case "p":
				if m.Current > 0 {
					m.Current--
				}
			case "b":
				m.Choosing = true
				m.Input = ""
				return m, tea.ClearScreen
			case "q":
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.WindowSize = msg
	}
	return m, nil
}

func (m Model) View() string {
	if m.Choosing {
		return fmt.Sprintf("%s\n\n%s", RenderPicker(m.WindowSize.Width), CenterText(m.Input, m.WindowSize.Width))
	}
	return RenderFragment(m.Fragments[m.Current], m.WindowSize)
}

func parseFragmentInput(input string, max int) (int, bool) {
	idx, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || idx < 1 || idx > max {
		return 0, false
	}
	return idx - 1, true
}

func isDigit(input string) bool {
	return input >= "0" && input <= "9"
}
