package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name                string
	MatchesPlayed       int
	Wins, Drawns, Losts int
}

func NewTeam(name string) *Team {
	return &Team{
		Name: name,
	}
}

func (t *Team) Win() {
	t.Wins++
	t.MatchesPlayed++
}

func (t *Team) Drawn() {
	t.Drawns++
	t.MatchesPlayed++
}

func (t *Team) Lost() {
	t.Losts++
	t.MatchesPlayed++
}

func (t *Team) String() string {
	return fmt.Sprintf("%s: p=%d w=%d d=%d l=%d s=%d", t.Name, t.MatchesPlayed, t.Wins, t.Drawns, t.Losts, t.Points())
}

type Competition struct {
	teams map[string]*Team
}

func NewCompetition() *Competition {
	return &Competition{
		teams: make(map[string]*Team),
	}
}

func (c *Competition) GetTeams() []*Team {
	teams := make([]*Team, 0, len(c.teams))
	for _, v := range c.teams {
		teams = append(teams, v)
	}

	return teams
}

func (c *Competition) RegisterMatch(teamA, teamB string, result string) error {
	if err := c.registerScore(teamA, result); err != nil {
		return err
	}

	oppositeResult, err := c.oppositeResult(result)
	if err != nil {
		return err
	}

	return c.registerScore(teamB, oppositeResult)
}

func (c *Competition) oppositeResult(result string) (string, error) {
	switch result {
	case "win":
		return "loss", nil
	case "draw":
		return "draw", nil
	case "loss":
		return "win", nil
	default:
		return "", fmt.Errorf("unknown result: %q", result)
	}
}

func (c *Competition) registerScore(team string, result string) error {
	t := c.getTeam(team)
	switch result {
	case "win":
		t.Win()
	case "draw":
		t.Drawn()
	case "loss":
		t.Lost()
	default:
		return fmt.Errorf("unknown result: %q", result)
	}

	return nil
}

func (c *Competition) getTeam(name string) *Team {
	t, ok := c.teams[name]
	if !ok {
		t = NewTeam(name)
		c.teams[name] = t
	}

	return t
}

// Points : A win earns a Team 3 points. A draw earns 1. A loss earns 0.
func (t *Team) Points() int {

	return t.Wins*3 + t.Drawns
}

func Tally(reader io.Reader, writer io.Writer) error {
	competition, err := readCompetitionData(reader)
	if err != nil {
		return err
	}

	if err := printCompetition(competition, writer); err != nil {
		return err
	}

	return nil
}

func readCompetitionData(reader io.Reader) (*Competition, error) {
	competition := NewCompetition()

	s := bufio.NewScanner(reader)
	for s.Scan() {
		if s.Text() == "" || strings.HasPrefix(s.Text(), "#") {
			continue // Ignore empty lines
		}

		line := strings.Split(s.Text(), ";")
		if len(line) != 3 {
			return nil, fmt.Errorf("lines should have three parts: teamA;teamB;score : %q instead", s.Text())
		}

		teamA, teamB, score := line[0], line[1], line[2]

		if err := competition.RegisterMatch(teamA, teamB, score); err != nil {
			return nil, err
		}
	}

	return competition, nil
}

func printCompetition(competition *Competition, writer io.Writer) error {
	teams := competition.GetTeams()

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points() != teams[j].Points() {
			return teams[i].Points() > teams[j].Points()
		}

		return teams[i].Name < teams[j].Name
	})

	b := bufio.NewWriter(writer)
	_, err := b.WriteString(fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team"))
	if err != nil {
		return err
	}

	template := "%-31s| %2d | %2d | %2d | %2d | %2d\n"
	for _, team := range teams {
		_, err := b.WriteString(fmt.Sprintf(template, team.Name, team.MatchesPlayed, team.Wins, team.Drawns, team.Losts, team.Points()))
		if err != nil {
			return err
		}
	}

	return b.Flush()
}
