package main 

import "slices"

type Player struct {
	Name string 
	Goals int 
	Misses int 
	Assists int 
	Rating float64
}

func (p *Player) calculateRating() {
	if p.Misses != 0 {
		p.Rating = (float64(p.Goals) + float64(p.Assists) / 2) / float64(p.Misses)
	} else {
		p.Rating = float64(p.Goals) + float64(p.Assists) / 2
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	var newPlayer = Player {Name: name, Goals: goals, Misses: misses, Assists: assists}
	newPlayer.calculateRating()
	return newPlayer
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, func (a, b Player) int {
		switch{
		case a.Goals < b.Goals: 
			return 1
		case a.Goals > b. Goals:
			return -1
		default:
			if a.Name < b.Name {
				return -1 
			} else if a.Name > b.Name {
				return 1
			} else {
				return 0
			}
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Rating > b.Rating:
			return -1
		case a.Rating < b.Rating:
			return 1
		default:
			if a.Name < b.Name {
				return -1 
			} else if a.Name > b.Name {
				return 1
			} else {
				return 0
			}
		}
	})
	return players
}

func playerGm(player Player) float64 {
	var pGm float64 
	if player.Misses != 0 {
		pGm = float64(player.Goals) / float64(player.Misses)
	} else {
		pGm = float64(player.Goals)
	}
	return pGm
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		aGm := playerGm(a)
		bGm := playerGm(b)
		switch {
		case aGm > bGm:
			return -1
		case aGm > bGm:
			return 1
		default:
			if a.Name < b.Name {
				return -1 
			} else if a.Name > b.Name {
				return 1
			} else {
				return 0
			}
		}
	})
	return players
}