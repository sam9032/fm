package tui

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/knipferrc/fm/dirfs"
	"github.com/knipferrc/fm/icons"
	"github.com/knipferrc/fm/strfmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

// statusBarView returns the status bar.
func (b Bubble) statusBarView() string {
	var logo string
	var status string

	width := lipgloss.Width
	selectedFileName := "N/A"
	fileCount := "0/0"

	if len(b.treeFiles) > 0 && b.treeFiles[b.treeCursor] != nil {
		selectedFile, err := b.treeFiles[b.treeCursor].Info()
		if err != nil {
			return "error"
		}
		fileCount = fmt.Sprintf("%d/%d", b.treeCursor+1, len(b.treeFiles))
		selectedFileName = selectedFile.Name()

		currentPath, err := dirfs.GetWorkingDirectory()
		if err != nil {
			currentPath = dirfs.CurrentDirectory
		}

		if len(b.foundFilesPaths) > 0 {
			currentPath = b.foundFilesPaths[b.treeCursor]
		}

		status = fmt.Sprintf("%s %s %s",
			selectedFile.ModTime().Format("2006-01-02 15:04:05"),
			selectedFile.Mode().String(),
			currentPath,
		)
	}

	if b.showCommandInput {
		status = b.textinput.View()
	}

	if b.moveMode {
		status = fmt.Sprintf("%s %s", "Currently moving:", b.treeFiles[b.treeCursor].Name())
	}

	if b.appConfig.Settings.ShowIcons && !b.appConfig.Settings.SimpleMode {
		logo = fmt.Sprintf("%s %s", icons.IconDef["dir"].GetGlyph(), "FM")
	} else {
		logo = "FM"
	}

	// Selected file styles
	selectedFileStyle := boldTextStyle.Copy().
		Foreground(b.theme.StatusBarSelectedFileForegroundColor).
		Background(b.theme.StatusBarSelectedFileBackgroundColor)

	if b.appConfig.Settings.SimpleMode {
		selectedFileStyle = boldTextStyle.Copy().
			Foreground(b.theme.DefaultTextColor)
	}

	selectedFileColumn := selectedFileStyle.
		Padding(0, 1).
		Height(StatusBarHeight).
		Render(truncate.StringWithTail(selectedFileName, 30, "..."))

	// File count styles
	fileCountStyle := boldTextStyle.Copy().
		Foreground(b.theme.StatusBarTotalFilesForegroundColor).
		Background(b.theme.StatusBarTotalFilesBackgroundColor)

	if b.appConfig.Settings.SimpleMode {
		fileCountStyle = boldTextStyle.Copy().
			Foreground(b.theme.DefaultTextColor)
	}

	fileCountColumn := fileCountStyle.
		Align(lipgloss.Right).
		Padding(0, 1).
		Height(StatusBarHeight).
		Render(fileCount)

	// Logo styles
	logoStyle := boldTextStyle.Copy().
		Foreground(b.theme.StatusBarLogoForegroundColor).
		Background(b.theme.StatusBarLogoBackgroundColor)

	if b.appConfig.Settings.SimpleMode {
		logoStyle = boldTextStyle.Copy().
			Foreground(b.theme.DefaultTextColor)
	}

	logoColumn := logoStyle.
		Padding(0, 1).
		Height(StatusBarHeight).
		Render(logo)

	// Status styles
	statusStyle := boldTextStyle.Copy().
		Foreground(b.theme.StatusBarBarForegroundColor).
		Background(b.theme.StatusBarBarBackgroundColor)

	if b.appConfig.Settings.SimpleMode {
		statusStyle = boldTextStyle.Copy().
			Foreground(b.theme.DefaultTextColor)
	}

	statusColumn := statusStyle.
		Padding(0, 1).
		Height(StatusBarHeight).
		Width(b.width - width(selectedFileColumn) - width(fileCountColumn) - width(logoColumn)).
		Render(truncate.StringWithTail(
			status,
			uint(b.width-width(selectedFileColumn)-width(fileCountColumn)-width(logoColumn)-3),
			"..."),
		)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		selectedFileColumn,
		statusColumn,
		fileCountColumn,
		logoColumn,
	)
}

// fileView returns the filetree view.
func (b Bubble) fileTreeView(files []fs.DirEntry) string {
	var directoryItem string
	curFiles := ""
	fileSize := ""
	selectedItemColor := b.theme.SelectedTreeItemColor
	unselectedItemColor := b.theme.UnselectedTreeItemColor

	for i, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return "Error loading directory tree"
		}

		icon, color := icons.GetIcon(fileInfo.Name(), filepath.Ext(fileInfo.Name()), icons.GetIndicator(fileInfo.Mode()))
		fileIcon := boldTextStyle.Copy().Width(2).Render(fmt.Sprintf("%s%s ", color, icon))

		if !b.appConfig.Settings.ShowIcons || b.appConfig.Settings.SimpleMode {
			fileIcon = boldTextStyle.Copy().Render("")
		}

		if b.appConfig.Settings.SimpleMode {
			selectedItemColor = b.theme.DefaultTextColor
		}

		if b.treeCursor == i {
			if len(b.fileSizes) > 0 {
				if b.fileSizes[i] != "" {
					fileSize = boldTextStyle.Copy().
						Foreground(colors["black"]).
						Background(selectedItemColor).
						Render(b.fileSizes[i])
				} else {
					fileSize = boldTextStyle.Copy().
						Foreground(colors["black"]).
						Background(selectedItemColor).
						Render(fileSizeLoadingStyle)
				}
			}

			directoryItem = boldTextStyle.Copy().
				Background(selectedItemColor).
				Width(b.primaryViewport.Width - lipgloss.Width(fileSize) - lipgloss.Width(fileIcon)).
				Foreground(colors["black"]).
				Render(
					truncate.StringWithTail(
						fileInfo.Name(), uint(b.primaryViewport.Width-lipgloss.Width(fileSize)), ellipsisStyle,
					),
				)
		} else {
			if len(b.fileSizes) > 0 {
				if b.fileSizes[i] != "" {
					fileSize = boldTextStyle.Copy().
						Foreground(unselectedItemColor).
						Render(b.fileSizes[i])
				} else {
					fileSize = boldTextStyle.Copy().
						Foreground(unselectedItemColor).
						Render(fileSizeLoadingStyle)
				}
			}

			directoryItem = boldTextStyle.Copy().
				Width(b.primaryViewport.Width - lipgloss.Width(fileSize) - lipgloss.Width(fileIcon)).
				Foreground(unselectedItemColor).
				Render(
					truncate.StringWithTail(
						fileInfo.Name(), uint(b.primaryViewport.Width-lipgloss.Width(fileSize)), ellipsisStyle,
					),
				)
		}

		row := lipgloss.JoinHorizontal(lipgloss.Top, fileIcon, directoryItem, fileSize)

		curFiles += fmt.Sprintf("%s\n", row)
	}

	if len(files) == 0 {
		curFiles = "Directory is empty"
	}

	return curFiles
}

// fileTreePreviewView returns a preview of a filetree.
func (b Bubble) fileTreePreviewView(files []fs.DirEntry) string {
	var directoryItem string
	curFiles := ""

	for _, file := range files {
		fileColor := b.theme.UnselectedTreeItemColor

		fileInfo, _ := file.Info()

		fileSize := lipgloss.NewStyle().
			Foreground(fileColor).
			Render(strfmt.ConvertBytesToSizeString(fileInfo.Size()))

		icon, color := icons.GetIcon(fileInfo.Name(), filepath.Ext(fileInfo.Name()), icons.GetIndicator(fileInfo.Mode()))
		fileIcon := fmt.Sprintf("%s%s", color, icon)

		switch {
		case b.appConfig.Settings.ShowIcons:
			directoryItem = fmt.Sprintf("%s\033[0m %s", fileIcon, boldTextStyle.Copy().
				Foreground(fileColor).
				Render(fileInfo.Name()))
		default:
			directoryItem = boldTextStyle.Copy().
				Foreground(fileColor).
				Render(fileInfo.Name())
		}

		dirItem := lipgloss.NewStyle().Width(
			b.secondaryViewport.Width - lipgloss.Width(fileSize),
		).Render(
			truncate.StringWithTail(
				directoryItem, uint(b.secondaryViewport.Width-lipgloss.Width(fileSize)), ellipsisStyle,
			),
		)

		row := lipgloss.JoinHorizontal(lipgloss.Top, dirItem, fileSize)

		curFiles += fmt.Sprintf("%s\n", row)
	}

	if len(files) == 0 {
		curFiles = "Directory is empty"
	}

	return curFiles
}

// textContentView returns some text content.
func (b Bubble) textContentView(content string) string {
	return lipgloss.NewStyle().
		Width(b.secondaryViewport.Width).
		Height(b.secondaryViewport.Height).
		Render(content)
}

// errorView returns an error message.
func (b Bubble) errorView(msg string) string {
	return lipgloss.NewStyle().
		Foreground(b.theme.ErrorColor).
		Width(b.secondaryViewport.Width).
		Height(b.secondaryViewport.Height).
		Render(msg)
}

// logView shows a list of logs.
func (b Bubble) logView() string {
	logList := ""
	title := boldTextStyle.Copy().
		Border(lipgloss.NormalBorder()).
		Italic(true).
		BorderBottom(true).
		BorderTop(false).
		BorderRight(false).
		BorderLeft(false).
		Foreground(b.theme.DefaultTextColor).
		Render("Application logs")

	for _, log := range b.logs {
		logList += fmt.Sprintf("%s\n", log)
	}

	return lipgloss.NewStyle().
		Width(b.secondaryViewport.Width).
		Height(b.secondaryViewport.Height).
		Render(lipgloss.JoinVertical(lipgloss.Top, title, logList))
}

// helpView returns help text.
func (b Bubble) helpView() string {
	type helpEntry struct {
		key         string
		description string
	}

	helpScreen := ""
	helpContent := []helpEntry{
		{"ctrl+c", "Exit FM"},
		{"j/up", "Move up"},
		{"k/down", "Move down"},
		{"h/left", "Go back a directory"},
		{"l/right", "Read file or enter directory"},
		{"p", "Preview directory"},
		{"gg", "Go to top of filetree or box"},
		{"G", "Go to bottom of filetree or box"},
		{"~", "Go to home directory"},
		{"/", "Go to root directory"},
		{".", "Toggle hidden files"},
		{"S", "Only show directories"},
		{"s", "Only show files"},
		{"y", "Copy file path to clipboard"},
		{"Z", "Zip currently selected tree item"},
		{"U", "Unzip currently selected tree item"},
		{"n", "Create new file"},
		{"N", "Create new directory"},
		{"ctrl+d", "Delete currently selected tree item"},
		{"M", "Move currently selected tree item"},
		{"enter", "Process command"},
		{"E", "Edit currently selected tree item"},
		{"C", "Copy currently selected tree item"},
		{"esc", "Reset FM to initial state"},
		{"O", "Show logs if debugging enabled"},
		{"tab", "Toggle between boxes"},
	}

	for _, content := range helpContent {
		keyText := boldTextStyle.Copy().Foreground(b.theme.DefaultTextColor).Width(12).Render(content.key)
		descriptionText := lipgloss.NewStyle().Foreground(b.theme.DefaultTextColor).Render(content.description)
		row := lipgloss.JoinHorizontal(lipgloss.Top, keyText, descriptionText)
		helpScreen += fmt.Sprintf("%s\n", row)
	}

	welcomeText := boldTextStyle.Copy().
		Border(lipgloss.NormalBorder()).
		Italic(true).
		BorderBottom(true).
		BorderTop(false).
		BorderRight(false).
		BorderLeft(false).
		Foreground(b.theme.DefaultTextColor).
		Render("Welcome to FM!")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		welcomeText,
		lipgloss.NewStyle().
			Width(b.secondaryViewport.Width).
			Height(b.secondaryViewport.Height).
			Render(helpScreen),
	)
}

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	if !b.ready {
		return fmt.Sprintf("%s %s", b.spinner.View(), "loading...")
	}

	var primaryBox string
	var secondaryBox string
	primaryBoxBorder := lipgloss.NormalBorder()
	secondaryBoxBorder := lipgloss.NormalBorder()
	primaryBoxBorderColor := b.theme.InactiveBoxBorderColor
	secondaryBoxBorderColor := b.theme.InactiveBoxBorderColor

	if b.activeBox == PrimaryBoxActive {
		primaryBoxBorderColor = b.theme.ActiveBoxBorderColor
	}

	if b.activeBox == SecondaryBoxActive {
		secondaryBoxBorderColor = b.theme.ActiveBoxBorderColor
	}

	if b.appConfig.Settings.Borderless {
		primaryBoxBorder = lipgloss.HiddenBorder()
		secondaryBoxBorder = lipgloss.HiddenBorder()
	}

	if b.appConfig.Settings.SimpleMode {
		primaryBoxBorder = lipgloss.HiddenBorder()
		secondaryBoxBorder = lipgloss.HiddenBorder()
	}

	if b.moveMode && !b.appConfig.Settings.SimpleMode && !b.appConfig.Settings.Borderless {
		primaryBoxBorder = starredBorder
	}

	b.primaryViewport.Style = lipgloss.NewStyle().
		PaddingLeft(BoxPadding).
		PaddingRight(BoxPadding).
		Border(primaryBoxBorder).
		BorderForeground(primaryBoxBorderColor)

	primaryBox = b.primaryViewport.View()

	if b.showBoxSpinner {
		b.primaryViewport.Style = lipgloss.NewStyle().
			PaddingLeft(BoxPadding).
			PaddingRight(BoxPadding).
			Border(primaryBoxBorder).
			BorderForeground(primaryBoxBorderColor)

		primaryBox = b.primaryViewport.Style.Render(fmt.Sprintf("%s loading...", b.spinner.View()))
	}

	if !b.appConfig.Settings.SimpleMode {
		b.secondaryViewport.Style = lipgloss.NewStyle().
			PaddingLeft(BoxPadding).
			PaddingRight(BoxPadding).
			Border(secondaryBoxBorder).
			BorderForeground(secondaryBoxBorderColor)

		secondaryBox = b.secondaryViewport.View()
	}

	view := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			primaryBox,
			secondaryBox,
		),
		b.statusBarView(),
	)

	return view
}
