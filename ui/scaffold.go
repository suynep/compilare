package ui

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
}

func InitModel() Model {
	return Model{
		AllPanes:       ALL_PANES,
		SelectedPane:   AeonPane,
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1", "a":
			m.SelectedPane = 0
			m.APRowData = append(m.APRowData, database.ReadAeonPosts()...)
			m.HNRowData = make([]types.WebPost, 0)
		case "2", "p":
			m.SelectedPane = 1
			m.APRowData = append(m.APRowData, database.ReadPsychePosts()...)
			m.HNRowData = make([]types.WebPost, 0)
		case "3", "h":
			m.SelectedPane = 2
			m.HNRowData = append(m.HNRowData, database.ReadForMemoization("t")...) // for testing purposes
			m.APRowData = make([]types.Item, 0)
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := ""
	for i, pane := range ALL_PANES {
		var paneDetails string
		if i == m.SelectedPane {
			paneDetails = fmt.Sprintf(">%s\n", pane)
		} else {
			paneDetails = pane + "\n"
		}
		s += paneDetails
	}
	s += "\n"

	if len(m.APRowData) != 0 {
		for _, rd := range m.APRowData {
			s += fmt.Sprintf("%s %s %s\n", rd.Title, rd.Creator, rd.PubDate)
		}
	} else if len(m.HNRowData) != 0 {
		for _, rd := range m.HNRowData {
			s += fmt.Sprintf("%s %s %s\n", rd.Title, rd.By, time.Unix(rd.Time, 0).Format(time.UnixDate))
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
