package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir    string
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
)

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home dir", err)
	}
	vaultDir = fmt.Sprintf("%s/.totion", homedir)
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	currentFile            *os.File
	notetextarea           textarea.Model
	list                   list.Model
	showingList            bool
}

func (m model) Init() tea.Cmd {
	return nil
}
func initialModel() model {

	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {
		log.Fatal(err)
	}
	ti := textinput.New()
	ti.Placeholder = "what did you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	ta := textarea.New()
	ta.Placeholder = "Write Your Through"
	ta.Focus()
	ta.ShowLineNumbers = false
	ta.Cursor.Style = cursorStyle
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Base = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(0, 1)

	notelist := listfiles()
	finallist := list.New(notelist, list.NewDefaultDelegate(), 0, 0)
	finallist.Title = "All Notes"
	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		currentFile:            nil,
		notetextarea:           ta,
		list:                   finallist,
	}
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	// Is it a key press?
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+l":
			notelist := listfiles()
			m.list.SetItems(notelist)
			m.showingList = true
			return m, nil
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}
			if m.currentFile != nil {
				m.currentFile = nil
			}
			if m.showingList {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.showingList = false
			}
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil

		case "ctrl+d":
			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					err := deletefile(item.title)
					if err != nil {
						log.Fatalf("Error deleting file: %v", err)
						return m, nil
					}
					notelist := listfiles()
					m.list.SetItems(notelist)
				}
			}
			return m, nil

		case "ctrl+s":

			if m.currentFile == nil {
				break

			}
			err := m.currentFile.Truncate(0)
			if err != nil {
				fmt.Println("cannot save this file")
				return m, nil

			}
			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("cannot save this file")
				return m, nil
			}
			if _, err := m.currentFile.WriteString(m.notetextarea.Value()); err != nil {
				fmt.Println("cannot save this file")
				return m, nil
			}
			if err := m.currentFile.Close(); err != nil {
				fmt.Println("cannot save this file")

			}
			m.currentFile = nil
			m.notetextarea.SetValue("")

			return m, nil
		case "enter":

			if m.showingList {
				item, ok := m.list.SelectedItem().(item)
				if ok {

					filepath := fmt.Sprintf("%s/%s", vaultDir, item.title)
					content, err := os.ReadFile(filepath)
					if err != nil {
						log.Fatalf("Error reading file: %v", err)
						return m, nil
					}
					m.notetextarea.SetValue(string(content))
					f, err := os.OpenFile(filepath, os.O_RDWR, 0644)
					if err != nil {
						log.Fatalf("Error opening file: %v", err)
						return m, nil
					}
					m.currentFile = f
					m.showingList = false
				}
				return m, nil
			}
			if m.currentFile != nil {
				break
			}
			filename := m.newFileInput.Value()
			if filename != "" {
				filepath := fmt.Sprintf("%s/%s.md", vaultDir, filename)

				if _, err := os.Stat(filepath); err == nil {
					return m, nil
				}

				f, err := os.Create(filepath)
				if err != nil {
					log.Fatalf("%v", err)
				}
				m.currentFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")

			}
			return m, nil

		}
	}
	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}
	if m.currentFile != nil {
		m.notetextarea, cmd = m.notetextarea.Update(msg)

	}
	if m.showingList {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd

}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("Welcome to Totion ")
	help := "Ctrl+N: new file | Ctrl+L: list | Esc: back/save | Ctrl+Q: quit"
	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.notetextarea.View()
	}

	if m.showingList {
		view = m.list.View()
	}
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Italic(true)

	styledHelp := helpStyle.Render(help)
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, styledHelp)
}
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func listfiles() []list.Item {
	items := make([]list.Item, 0)
	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("Error Reading Notes")
	}
	for _, entries := range entries {
		if !entries.IsDir() {
			info, err := entries.Info()
			if err != nil {
				continue
			}
			modTime := info.ModTime().Format("2006-01-02 15:04")
			items = append(items, item{
				title: entries.Name(),
				desc:  fmt.Sprintf("Modified: %s", modTime),
			})
		}
	}
	return items
}
func deletefile(filename string) error {
	filepath := fmt.Sprintf("%s/%s", vaultDir, filename)
	err := os.Remove(filepath)
	return err
}
