package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type playertype struct {
	name       string
	dice       int
	score      int
	isplaying  bool
	endroll    []int
	prevplayer *playertype
	nextplayer *playertype
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	var (
		playerstr    string
		dicestr      string
		isGameFinish bool = false
		head         *playertype
		listplayer   []*playertype
		winner       string
		lastplayer   string
		maxscore     int = 0
	)

	fmt.Print("Pemain: ")
	fmt.Scanln(&playerstr)
	player, err := strconv.Atoi(playerstr)
	if err != nil {
		fmt.Println("input pemain salah")
		return
	}

	fmt.Print("Dadu: ")
	fmt.Scanln(&dicestr)
	dice, err := strconv.Atoi(dicestr)
	if err != nil {
		fmt.Println("input dadu salah")
		return
	}

	fmt.Println("Pemain =", player, "Dadu =", dice)
	fmt.Println("=====================")

	tmp := &playertype{}

	for i := 0; i < player; i++ {
		p := playertype{
			name:       fmt.Sprintf("Pemain #%d", i+1),
			dice:       dice,
			score:      0,
			isplaying:  true,
			prevplayer: tmp,
		}

		listplayer = append(listplayer, &p)

		if i == 0 {
			head = &p
			tmp = &p
			continue
		}

		tmp.nextplayer = &p
		tmp = &p
	}

	tmp.nextplayer = head
	head.prevplayer = tmp

	currplayer := head
	turn := 1
	for !isGameFinish {

		// start turn
		fmt.Printf("Giliran %d lempar dadu:\n", turn)
		for _, val := range listplayer {
			currroll := []int{}
			rollstr := []string{}
			if val.isplaying {
				for j := 0; j < val.dice; j++ {
					roll := random.Intn(6) + 1
					currroll = append(currroll, roll)
				}
				for _, val := range currroll {
					rollstr = append(rollstr, strconv.Itoa(val))
				}

				val.endroll = currroll
			}

			fmt.Printf("	%s (%d): ", val.name, val.score)
			if len(rollstr) > 0 {
				fmt.Println(strings.Join(rollstr, ","))
			} else {
				fmt.Println("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
		}

		// evaluate
		for i := 0; i < player; i++ {
			newendroll := []int{}
			for _, val := range currplayer.endroll {
				if val == 6 {
					currplayer.score++
					currplayer.dice--
				} else if val == 1 {
					currplayer.nextplayer.dice++
					currplayer.dice--
				} else {
					newendroll = append(newendroll, val)
				}
			}
			currplayer.endroll = newendroll
			currplayer = currplayer.nextplayer
		}

		currplayer = head

		for i := 0; i < player; i++ {
			if currplayer.dice == 0 {
				currplayer.isplaying = false
				if currplayer == head {
					head = currplayer.nextplayer
					currplayer.prevplayer.nextplayer = currplayer.nextplayer
					currplayer.nextplayer.prevplayer = currplayer.prevplayer
				} else {
					currplayer.prevplayer.nextplayer = currplayer.nextplayer
					currplayer.nextplayer.prevplayer = currplayer.prevplayer
				}
				player--
			}

			currplayer = currplayer.nextplayer
		}

		// after evaluate
		fmt.Println("Setelah evaluasi:")
		for _, val := range listplayer {
			rollstr := []string{}
			if val.isplaying {
				lastplayer = val.name
				surplus := val.dice - len(val.endroll)
				if surplus > 0 {
					for j := 0; j < surplus; j++ {
						val.endroll = append(val.endroll, 1)
					}
				}

				for _, val := range val.endroll {
					rollstr = append(rollstr, strconv.Itoa(val))
				}
			}

			fmt.Printf("	%s (%d): ", val.name, val.score)
			if len(rollstr) > 0 {
				fmt.Println(strings.Join(rollstr, ","))
			} else {
				fmt.Println("_ (Berhenti bermain karena tidak memiliki dadu)")
			}

			if val.score > maxscore {
				maxscore = val.score
				winner = val.name
			}
		}

		// reset link list pointer
		currplayer = head
		fmt.Println("=====================")

		if player == 1 {
			isGameFinish = true
		}
		turn++
	}

	fmt.Printf("Game berakhir karena hanya %s yang memiliki dadu.\n", lastplayer)
	fmt.Printf("Game dimenangkan oleh %s karena memiliki poin lebih banyak dari pemain lainnya.\n", winner)

}
