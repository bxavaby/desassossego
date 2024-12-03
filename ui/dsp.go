package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Fragment struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func LoadFragments(paths ...string) ([]Fragment, error) {
	var file *os.File
	var err error

	for _, path := range paths {
		file, err = os.Open(path)
		if err == nil {
			defer file.Close()
			break
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var fragments []Fragment
	if err := json.NewDecoder(file).Decode(&fragments); err != nil {
		return nil, fmt.Errorf("failed to parse fragments: %w", err)
	}

	return fragments, nil
}

func RenderFragment(fragment Fragment, size tea.WindowSizeMsg) string {
	width := size.Width

	titleStyle := lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		Width(width).
		Align(lipgloss.Center)

	textStyle := lipgloss.NewStyle().
		Foreground(Ornament).
		Width(width - 4).
		Align(lipgloss.Left)

	title := titleStyle.Render(fmt.Sprintf("Frag.: %s", fragment.Title))
	text := textStyle.Render(strings.Join(splitToWidth(fragment.Text, width-4), "\n"))

	return fmt.Sprintf("\n%s\n\n%s", title, text)
}

func RenderPicker(width int) string {
	header := lipgloss.NewStyle().
		Foreground(Accent).
		Bold(true).
		Width(width).
		Align(lipgloss.Center).
		Render("Enter the frag. # (1-433):")

	return fmt.Sprintf("\n%s", header)
}

func CenterText(input string, width int) string {
	padding := (width - len(input)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + input
}

func splitToWidth(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	words := strings.Fields(text)
	var lines []string
	line := ""
	for _, word := range words {
		if len(line)+len(word)+1 > width {
			lines = append(lines, line)
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
