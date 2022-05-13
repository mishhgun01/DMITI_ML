package pkg

import (
	"math/rand"
	"time"
)

type Player struct {
	Score      int
	Picked     int
	TableScore []int
}

func (pl *Player) RandomizePicker(p float64) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	if p >= r {
		pl.Picked = 0
	} else {
		pl.Picked = 1
	}
}

func Winner(p1, p2 *Player) (bool, []int) {
	var flag bool
	var score []int
	switch p1.Picked {
	case 0:
		switch p2.Picked {
		case 0:
			p1.Score += 2
			p2.Score -= 2
			p1.TableScore = append(p1.TableScore, 1)
			p2.TableScore = append(p2.TableScore, 0)
			flag = true
			score = append(score, 2, 0)
		case 1:
			p1.Score -= 3
			p2.Score += 3
			p1.TableScore = append(p1.TableScore, 0)
			p2.TableScore = append(p2.TableScore, 1)
			flag = false
			score = append(score, -3, 1)
		}
	case 1:
		switch p2.Picked {
		case 0:
			p1.Score -= 1
			p2.Score += 1
			p1.TableScore = append(p1.TableScore, 0)
			p2.TableScore = append(p2.TableScore, 1)
			flag = false
			score = append(score, -1, 1)
		case 1:
			p1.Score += 2
			p2.Score -= 2
			p1.TableScore = append(p1.TableScore, 1)
			p2.TableScore = append(p2.TableScore, 0)
			flag = true
			score = append(score, 2, 1)
		}
	}
	return flag, score
}
