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
	"time"

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
	list       list.Model
	selected   string
	loadingMsg string
	spinner    spinner.Model
	err        error
	screen     int
}

func screenOne() *model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	m := model{screen: 0, spinner: s}
	return &m
}

type switchMessage int

func SwitchScreens(screen int) tea.Cmd {
	return func() tea.Msg {
		return switchMessage(screen)
	}
}

func (m *model) getData(title string) tea.Cmd {
	m.loadingMsg = fmt.Sprintf("Setting up Day %v", title)
	return func() tea.Msg {
		parsemd.RunConvert(title)
		time.Sleep(time.Second)
		return switchMessage(1)
	}
}

func (m *model) Init() tea.Cmd {
	switch m.screen {
	case 0:
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
		m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)
		m.list.Title = "Which day would you like to bootstrap?"
		return SwitchScreens(1)
	case 2:
		return tea.Batch(m.spinner.Tick, m.getData(m.selected))
	default:
		return nil
	}
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.screen == 1 {
				var title string
				if i, ok := m.list.SelectedItem().(item); ok {
					title = i.Title()
				}
				m.selected = title
				m.screen = 2
			} else {
				m.screen = 1
			}
			return m, SwitchScreens(m.screen)
		}

	case switchMessage:
		m.screen = int(msg)
		return RootScreen().SwitchScreen(m)

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.screen == 2 {
		var cmd tea.Cmd
		cmd = m.getData(m.selected)
		return m, cmd
	} else {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.screen == 1 {
		return docStyle.Render(m.list.View())
	} else {
		str := fmt.Sprintf("\n   %s %s...\n\n", m.spinner.View(), m.loadingMsg)
		return str
	}
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
