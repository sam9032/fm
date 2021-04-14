package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/fm/src/components"
)

func fileManager(files []fs.FileInfo, cursor int) string {
	doc := strings.Builder{}
	curFiles := ""

	for i, file := range files {
		curFiles += fmt.Sprintf("%s\n", components.FileListing(file.Name(), cursor == i, file.IsDir(), filepath.Ext(file.Name())))
	}

	doc.WriteString(lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Width(50).Align(lipgloss.Left).Render(curFiles),
	))

	return doc.String()
}

func (m model) View() string {
	m.Viewport.SetContent(fileManager(m.Files, m.Cursor))
	var file fs.FileInfo = nil

	if len(m.Files) > 0 {
		file = m.Files[m.Cursor]
	}

	view := fmt.Sprintf("%s%s", m.Viewport.View(), components.StatusBar(m.ScreenWidth, file, m.Move, m.Rename, m.Delete, &m.TextInput))

	return view
}
