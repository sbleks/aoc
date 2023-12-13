package main

import (
	"aoc2md/parsemd"
	"aoc2md/scrape"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, tea.Quit
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

func main() {
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

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
