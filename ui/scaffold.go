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

var AEON_POSTS []types.Item = make([]types.Item, 0)
var PSYCHE_POSTS []types.Item = make([]types.Item, 0)
var HN_POSTS []types.WebPost = make([]types.WebPost, 0)

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
	StartIndex     int
	EndIndex       int
	CurrentPointer int
	OpenLink       bool
}

func InitModel() Model {
	AEON_POSTS = database.ReadAeonPosts()
	PSYCHE_POSTS = database.ReadPsychePosts()
	HN_POSTS = database.ReadForMemoization("t") // arbitrary choice of top posts as of now
	return Model{
		AllPanes:       ALL_PANES,
		SelectedPane:   -1,
		APRowData:      make([]types.Item, 0),
		HNRowData:      make([]types.WebPost, 0),
		PerPage:        10,
		CurrentPointer: 0,
		StartIndex:     0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1", "a":
			m.StartIndex = 0 // reset page nav
			m.SelectedPane = 0
			if m.StartIndex+m.PerPage < len(AEON_POSTS) {
				m.APRowData = AEON_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
			}
			m.HNRowData = make([]types.WebPost, 0)
		case "2", "p":
			m.StartIndex = 0 // reset page nav
			m.SelectedPane = 1
			if m.StartIndex+m.PerPage < len(PSYCHE_POSTS) {
				m.APRowData = PSYCHE_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
			}
			m.HNRowData = make([]types.WebPost, 0)
		case "3", "h":
			m.StartIndex = 0 // reset page nav
			m.SelectedPane = 2
			if m.StartIndex+m.PerPage < len(HN_POSTS) {
				m.HNRowData = HN_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
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
		case "n":
			m.StartIndex = m.StartIndex + m.PerPage
			switch m.SelectedPane {
			case 0:
				if m.StartIndex+m.PerPage <= len(AEON_POSTS) {
					m.APRowData = AEON_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.HNRowData = make([]types.WebPost, 0)
			case 1:
				if m.StartIndex+m.PerPage <= len(PSYCHE_POSTS) {
					m.APRowData = PSYCHE_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.HNRowData = make([]types.WebPost, 0)
			case 2:
				if m.StartIndex+m.PerPage <= len(HN_POSTS) {
					m.HNRowData = HN_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.APRowData = make([]types.Item, 0)

			}
		case "P":
			if m.StartIndex > 0 {
				m.StartIndex = m.StartIndex - m.PerPage
			}
			switch m.SelectedPane {
			case 0:
				if m.StartIndex+m.PerPage <= len(AEON_POSTS) {
					m.APRowData = AEON_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.HNRowData = make([]types.WebPost, 0)
			case 1:
				if m.StartIndex+m.PerPage <= len(PSYCHE_POSTS) {
					m.APRowData = PSYCHE_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.HNRowData = make([]types.WebPost, 0)
			case 2:
				if m.StartIndex+m.PerPage <= len(HN_POSTS) {
					m.HNRowData = HN_POSTS[m.StartIndex : m.StartIndex+m.PerPage]
				}
				m.APRowData = make([]types.Item, 0)

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
				s += ArticleSelectStyleOp(fmt.Sprintf(">%s %s %s", rd.Title, rd.Creator, rd.PubDate)) + "\n"
			} else {
				s += fmt.Sprintf("%s %s %s\n", rd.Title, rd.Creator, rd.PubDate)
			}
		}
	} else if len(m.HNRowData) != 0 {
		for i, rd := range m.HNRowData {
			if i == m.CurrentPointer {
				s += ArticleSelectStyleOp(fmt.Sprintf(">%s %s %s", rd.Title, rd.By, time.Unix(rd.Time, 0).Format(time.UnixDate))) + "\n"
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
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#00bb00")).Background(lipgloss.Color("#ffffff"))
	style = style.SetString(msg)
	return style.String()
}

func ArticleSelectStyleOp(msg string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#f5dd42"))
	style = style.SetString(msg)
	return style.String()
}
