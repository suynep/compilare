package ui

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/suynep/compilare/database"
	"github.com/suynep/compilare/types"
)

const (
	AeonPane = iota
	PsychePane
	HackernewsPane
)

var ALL_PANES = []string{"Aeon", "Psyche", "Hackernews"}

type Data struct {
	Title       string
	PublishedAt string
	Author      string
}

type Model struct {
	AllPanes       []string
	SelectedPane   int
	APRowData      []types.Item // Aeon/Psyche Row Data
	HNRowData      []types.WebPost
	PerPage        int
	CurrentPointer int
	OpenLink       bool
}

func InitModel() Model {
	return Model{
		AllPanes:       ALL_PANES,
		SelectedPane:   -1,
		APRowData:      make([]types.Item, 0),
		HNRowData:      make([]types.WebPost, 0),
		PerPage:        10,
		CurrentPointer: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	aeonPosts := database.ReadAeonPosts()
	psychePosts := database.ReadPsychePosts()
	hnPosts := database.ReadForMemoization("t")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1", "a":
			m.SelectedPane = 0
			if m.CurrentPointer+m.PerPage < len(aeonPosts) {
				m.APRowData = aeonPosts[m.CurrentPointer : m.CurrentPointer+m.PerPage]
			}
			m.HNRowData = make([]types.WebPost, 0)
		case "2", "p":
			m.SelectedPane = 1
			if m.CurrentPointer+m.PerPage < len(psychePosts) {
				m.APRowData = psychePosts[m.CurrentPointer : m.CurrentPointer+m.PerPage]
			}
			m.HNRowData = make([]types.WebPost, 0)
		case "3", "h":
			m.SelectedPane = 2
			if m.CurrentPointer+m.PerPage < len(hnPosts) {
				m.HNRowData = hnPosts[m.CurrentPointer : m.CurrentPointer+m.PerPage]
			}
			m.APRowData = make([]types.Item, 0)
		case "j":
			m.CurrentPointer = (m.CurrentPointer + 1) % m.PerPage
		case "k":
			if m.CurrentPointer <= 0 {
				m.CurrentPointer = m.PerPage - 1
			} else {
				m.CurrentPointer = m.CurrentPointer - 1
			}
		case "enter", "l":
			if len(m.APRowData) != 0 {
				err := exec.Command("xdg-open", m.APRowData[m.CurrentPointer].Link).Run()
				if err != nil {
					panic(err)
				}
			} else if len(m.HNRowData) != 0 {
				err := exec.Command("xdg-open", m.HNRowData[m.CurrentPointer].Url).Run()
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := ""
	for i, pane := range ALL_PANES {
		var paneDetails string
		if i == m.SelectedPane {
			paneDetails = ChosenStyleOp(fmt.Sprintf(">%s", pane)) + "\n"
		} else {
			paneDetails = pane + "\n"
		}
		s += paneDetails
	}
	s += "\n"

	if len(m.APRowData) != 0 {
		for i, rd := range m.APRowData {
			if i == m.CurrentPointer {
				s += ChosenStyleOp(fmt.Sprintf(">%s %s %s", rd.Title, rd.Creator, rd.PubDate)) + "\n"
			} else {
				s += fmt.Sprintf("%s %s %s\n", rd.Title, rd.Creator, rd.PubDate)
			}
		}
	} else if len(m.HNRowData) != 0 {
		for i, rd := range m.HNRowData {
			if i == m.CurrentPointer {
				s += ChosenStyleOp(fmt.Sprintf(">%s %s %s", rd.Title, rd.By, time.Unix(rd.Time, 0).Format(time.UnixDate))) + "\n"
			} else {
				s += fmt.Sprintf("%s %s %s\n", rd.Title, rd.By, time.Unix(rd.Time, 0).Format(time.UnixDate))
			}
		}

	}
	return s
}

func BuildUi() {
	p := tea.NewProgram(InitModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while running!\n")
		os.Exit(1)
	}
}

// styles

func ChosenStyleOp(msg string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#00bb00"))
	style = style.SetString(msg)
	return style.String()
}
