package main 

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