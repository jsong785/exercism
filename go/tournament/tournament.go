package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(reader io.Reader, writer io.Writer) error {
	statsList := CreateTeamStatsList()

	if err := GetTeamStats(statsList, reader); err != nil {
		return err
	}
	if err := WriteTeamStats(statsList, writer); err != nil {
		return err
	}
	return nil
}

// parsing
func GetTeamStats(statsList *TeamStatsList, reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue
		}
		if isComment := (trimmedLine[0] == '#'); isComment {
			continue
		}

		tokens := strings.Split(trimmedLine, ";")
		if len(tokens) != 3 {
			return errors.New("invalid")
		}

		team1 := tokens[0]
		team2 := tokens[1]
		matchResult := tokens[2]
		if err := HandleLineTokens(team1, team2, matchResult, statsList); err != nil {
			return err
		}
	}
	return nil
}

func HandleLineTokens(t1, t2, results string, statsList *TeamStatsList) error {
	t1s := statsList.GetOrCreateTeam(t1)
	t2s := statsList.GetOrCreateTeam(t2)
	err := t1s.PlayedMatchAgainst(t2s, results)
	return err
}

// writing
func WriteTeamStats(statsList *TeamStatsList, writer io.Writer) error {
	if len(statsList.teams) == 0 {
		return nil
	}
	sortedStats := statsList.GetSortedTeamStats()

	buffer := bufio.NewWriter(writer)
	defer buffer.Flush()
	if _, err := buffer.WriteString(GetTeamStatsHeaderString()); err != nil {
		return err
	}

	for _, s := range sortedStats {
		if _, err := buffer.WriteString(GetTeamStatsString(s)); err != nil {
			return err
		}
	}
	return nil
}

const SPRINTF_STATS_STRING = "%-30v | %2v | %2v | %2v | %2v | %2v\n"

func GetTeamStatsHeaderString() string {
	return fmt.Sprintf(SPRINTF_STATS_STRING,
		"Team", "MP", "W", "D", "L", "P")
}

func GetTeamStatsString(t *TeamStats) string {
	return fmt.Sprintf(SPRINTF_STATS_STRING,
		t.name,
		t.GetMatchesPlayed(),
		t.wins,
		t.draws,
		t.losses,
		t.CalculatePoints())
}

// TeamStatsList
type TeamStatsList struct {
	teams map[string]*TeamStats
}

func CreateTeamStatsList() *TeamStatsList {
	statsList := new(TeamStatsList)
	statsList.teams = make(map[string]*TeamStats)
	return statsList
}

func (stats *TeamStatsList) GetOrCreateTeam(name string) *TeamStats {
	found := stats.teams[name]
	if found == nil {
		newTeam := new(TeamStats)
		newTeam.name = name

		stats.teams[name] = newTeam
		found = stats.teams[name]
	}
	return found
}

func (stats *TeamStatsList) GetSortedTeamStats() []*TeamStats {
	sortedTeamsSlice := make([]*TeamStats, 0)
	for _, team := range stats.teams {
		sortedTeamsSlice = append(sortedTeamsSlice, team)
	}

	// sort by name
	sort.Slice(sortedTeamsSlice,
		func(i, j int) bool {
			return sortedTeamsSlice[i].name < sortedTeamsSlice[j].name
		})

	// reverse sort by points
	// (stable, keep alphabetical order between teams of equal points)
	sort.SliceStable(sortedTeamsSlice,
		func(i, j int) bool {
			return sortedTeamsSlice[i].CalculatePoints() >
				sortedTeamsSlice[j].CalculatePoints()
		})

	return sortedTeamsSlice
}

// TeamStats
type Points int

const (
	POINTS_WIN  Points = 3
	POINTS_DRAW        = 1
	POINTS_LOSS        = 0
)

type TeamStats struct {
	name   string
	wins   int
	losses int
	draws  int
}

func (t *TeamStats) GetMatchesPlayed() int {
	return t.wins + t.losses + t.draws
}

func (t *TeamStats) CalculatePoints() int {
	return (t.wins * int(POINTS_WIN)) +
		(t.losses * int(POINTS_LOSS)) +
		(t.draws * int(POINTS_DRAW))
}

func (t *TeamStats) PlayedMatchAgainst(opponent *TeamStats,
	results string) error {
	switch results {
	case "win":
		t.wins += 1
		opponent.losses += 1
	case "loss":
		t.losses += 1
		opponent.wins += 1
	case "draw":
		t.draws += 1
		opponent.draws += 1
	default:
		return errors.New("Invalid MatchResults")
	}
	return nil
}
