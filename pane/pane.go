package pane

import (
	"github.com/knipferrc/fm/config"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width    int
	Height   int
	IsActive bool
	Viewport viewport.Model
}

func (m *Model) View() string {
	cfg := config.GetConfig()
	borderColor := cfg.Colors.Pane.InactivePane
	borderType := lipgloss.NormalBorder()

	if cfg.Settings.RoundedPanes {
		borderType = lipgloss.RoundedBorder()
	} else {
		borderType = lipgloss.NormalBorder()
	}

	if m.IsActive {
		borderColor = cfg.Colors.Pane.ActivePane
	}

	return lipgloss.NewStyle().
		BorderForeground(lipgloss.Color(borderColor)).
		Border(borderType).
		Width(m.Width).
		Height(m.Height).
		Render(m.Viewport.View())
}

func (m *Model) SetContent(content string) {
	m.Viewport.SetContent(content)
}
