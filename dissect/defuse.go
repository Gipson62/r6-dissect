package dissect

import (
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *Reader) getTeamByRole(role TeamRole) int {
	for i, team := range r.Header.Teams {
		if team.Role == role {
			return i
		}
	}
	return -1
}

func (r *Reader) getAlivePlayersByTeam(teamIndex int) []string {
	var alive []string
	for _, p := range r.Header.Players {
		if p.TeamIndex == teamIndex {
			died := false
			for _, fb := range r.MatchFeedback {
				if fb.Type == Kill && fb.Target == p.Username {
					died = true
					break
				}
				if fb.Type == Death && fb.Username == p.Username {
					died = true
					break
				}
			}
			if !died && p.Username != "" {
				alive = append(alive, p.Username)
			}
		}
	}
	return alive
}

func readDefuserTimer(r *Reader) error {
	timer, err := r.String()
	if err != nil {
		return err
	}

	var playerIndex int = -1

	if r.Header.CodeVersion >= Y10S4 {
		var targetRole TeamRole
		if r.planted {
			targetRole = Defense
		} else {
			targetRole = Attack
		}

		teamIndex := r.getTeamByRole(targetRole)
		if teamIndex >= 0 {
			alive := r.getAlivePlayersByTeam(teamIndex)
			if len(alive) == 1 {
				for i, p := range r.Header.Players {
					if p.Username == alive[0] {
						playerIndex = i
						break
					}
				}
			}
		}
	} else {
		if err = r.Skip(34); err != nil {
			return err
		}
		id, err := r.Bytes(4)
		if err != nil {
			return err
		}
		playerIndex = r.PlayerIndexByID(id)
	}

	if playerIndex > -1 && r.lastDefuserPlayerIndex != playerIndex {
		a := DefuserPlantStart
		if r.planted {
			a = DefuserDisableStart
		}
		u := MatchUpdate{
			Type:          a,
			Username:      r.Header.Players[playerIndex].Username,
			Time:          r.timeRaw,
			TimeInSeconds: r.time,
		}
		r.MatchFeedback = append(r.MatchFeedback, u)
		log.Debug().Interface("match_update", u).Send()
		r.lastDefuserPlayerIndex = playerIndex
	}

	// TODO: 0.00 can be present even if defuser was not disabled.
	if !strings.HasPrefix(timer, "0.00") {
		return nil
	}
	a := DefuserDisableComplete
	if !r.planted {
		a = DefuserPlantComplete
		r.planted = true
	} else {
		for i := len(r.MatchFeedback) - 1; i >= 0; i-- {
			if r.MatchFeedback[i].Type == DefuserDisableComplete {
				return nil
			}
			if r.MatchFeedback[i].Type == DefuserPlantComplete {
				break
			}
		}
	}

	username := ""
	if r.lastDefuserPlayerIndex >= 0 && r.lastDefuserPlayerIndex < len(r.Header.Players) {
		username = r.Header.Players[r.lastDefuserPlayerIndex].Username
	}

	u := MatchUpdate{
		Type:          a,
		Username:      username,
		Time:          r.timeRaw,
		TimeInSeconds: r.time,
	}
	r.MatchFeedback = append(r.MatchFeedback, u)
	log.Debug().Interface("match_update", u).Send()
	return nil
}
