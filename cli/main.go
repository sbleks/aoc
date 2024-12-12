package main

import (
	"aoc2md/parsemd"
	"aoc2md/scrape"
	"cli/spinner"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    list.Model
	spinner spinner.Model
	err     error
	screen  int
}

func screenOne() *model {
	days, err := scrape.ScrapeDays("")
	if err != nil {
		log.Panicf("%v\n", err)
	}

	sort.Slice(days[:], func(i, j int) bool {
		for x := range days[i] {
			if days[i][x] == days[j][x] {
				continue
			}

			a, err := strconv.Atoi(days[i][x])
			if err != nil {
				continue
			}

			b, err := strconv.Atoi(days[j][x])
			if err != nil {
				continue
			}

			return a < b
		}
		return false
	})

	var items []list.Item

	for _, day := range days {
		items = append(items, item{title: day[0], desc: day[1]})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Which day would you like to bootstrap?"
	return &m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			var title string
			if i, ok := m.list.SelectedItem().(item); ok {
				title = i.Title()
			}
			// log.Panicf("%s\n", title)
			parsemd.RunConvert(title)
			return RootScreen().SwitchScreen(screenTwo())
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

type screenTwoModel struct {
	spinner spinner.Model
	err     error
}

func screenTwo() *screenTwoModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	// s.Style = inputStyle
	return &screenTwoModel{
		spinner: s,
	}
}

func (m screenTwoModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *screenTwoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			// any other key switches the screen
			return RootScreen().SwitchScreen(screenOne())
		}
	case error:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *screenTwoModel) View() string {
	str := fmt.Sprintf("\n   %s Downloading...\n\n", m.spinner.View())
	return str
}

// Root Screen

type rootScreenModel struct {
	model         tea.Model // this will hold the current screen model
	initialRender bool
}

func RootScreen() rootScreenModel {

	// sample conditional logic to start with a specific screen
	// notice that the screen methods Update and View have been modified
	// to accept a pointer *screenXModel instead of screenXModel
	// this will allow us to modify the model's state in the View method
	// if needed

	return rootScreenModel{
		model:         screenOne(),
		initialRender: false,
	}
}

func (m rootScreenModel) Init() tea.Cmd {
	return m.model.Init() // rest methods are just wrappers for the model's methods
}

func (m rootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.model.Update(msg)
}

func (m rootScreenModel) View() string {
	return m.model.View()
}

// this is the switcher which will switch between screens
func (m rootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init() // must return .Init() to initialize the screen (and here the magic happens)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	p := tea.NewProgram(RootScreen(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
