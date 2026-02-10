# Go-Totion 📝

A beautiful, terminal-based note-taking application written in Go, inspired by Notion. Go-Totion provides a clean and intuitive interface for creating, editing, and managing markdown notes directly from your terminal.

## ✨ Features

- **Beautiful TUI Interface**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for a smooth terminal experience
- **Markdown Support**: Create and edit notes in Markdown format (.md files)
- **List View**: Browse all your notes with modification timestamps
- **Quick Navigation**: Keyboard shortcuts for efficient note management
- **File Operations**: Create, read, update, and delete notes
- **Local Storage**: Notes are stored locally in `~/.totion` directory
- **Search & Filter**: Quickly find notes using the built-in filter

## 🚀 Installation

### Prerequisites

- Go 1.24 or higher

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/Rudraksh121a/Go-Totion.git
cd Go-Totion
```

2. Build the application:
```bash
make build
```

Or manually:
```bash
go build -o totion .
```

3. Run the application:
```bash
./totion
```

Or use:
```bash
make run
```

## 📖 Usage

When you start Go-Totion, you'll see a welcome screen with available keyboard shortcuts at the bottom.

### Creating a New Note

1. Press `Ctrl+N` to create a new note
2. Type the name of your note
3. Press `Enter` to confirm
4. Start writing your note
5. Press `Esc` to go back (note is saved automatically)

### Viewing and Editing Notes

1. Press `Ctrl+L` to view all notes
2. Use arrow keys to navigate through the list
3. Press `Enter` to open and edit a note
4. Press `Ctrl+S` to save your changes
5. Press `Esc` to go back to the list

### Deleting a Note

1. Press `Ctrl+L` to view all notes
2. Navigate to the note you want to delete
3. Press `Ctrl+D` to delete the selected note

## ⌨️ Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Ctrl+N` | Create a new note |
| `Ctrl+L` | List all notes |
| `Ctrl+S` | Save current note |
| `Ctrl+D` | Delete selected note (in list view) |
| `Esc` | Go back / Close current view |
| `Ctrl+C` or `q` | Quit the application |
| `/` | Filter notes (in list view) |
| `↑/↓` | Navigate through notes |

## 🗂️ File Storage

All notes are stored in your home directory under `.totion/`:
```
~/.totion/
├── note1.md
├── note2.md
└── ...
```

## 🛠️ Technical Details

### Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components (list, textarea, textinput)
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal layouts

### Project Structure

- `main.go` - Main application logic including:
  - Model definition and state management
  - Event handling (keyboard shortcuts)
  - View rendering
  - File operations (create, read, update, delete)

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- Built with the amazing [Charm](https://charm.sh/) TUI libraries
- Inspired by [Notion](https://www.notion.so/)

## 📧 Contact

For questions or feedback, please open an issue on GitHub.

---

Made with ❤️ using Go and Bubble Tea
