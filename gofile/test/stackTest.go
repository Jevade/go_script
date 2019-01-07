package main

import (
	"fmt"
	"math/rand"

	"./stackArr"
)

//Pos is a struct to desolve palace puzzle
type Pos struct {
	X    int
	Y    int
	Dir  int
	Len  int
	maps *[]int
}

func initMap(N int) *[]int {
	maps := make([]int, N*N)
	for ix := 0; ix < N*N; ix++ {
		if bool(ix%N == 0) || bool(ix%N == N-1) || bool(ix/N == 0) || bool(ix/N == N-1) {
			maps[ix] = 1
		} else {
			seed := rand.Intn(999)
			maps[ix] = rand.Intn(seed) % 2
		}
	}
	maps[1*N+1] = 0
	maps[6*N+1] = 0
	maps[5*N+1] = 0
	maps[2*N+1] = 0
	return &maps
}

func (pos *Pos) showMap() {
	for ix := 0; ix < pos.Len*pos.Len; ix++ {
		if ix%pos.Len == 0 {
			fmt.Printf("\n|")
		}
		if (*pos.maps)[ix] == 0 {
			fmt.Printf(" %s|", "x")
		}
		if (*pos.maps)[ix] == 1 {
			fmt.Printf(" %s|", ".")
		}
		if (*pos.maps)[ix] == 10 {
			fmt.Printf(" %s|", "p")
		}

	}
}

//IsWall determine outof index or is wall
func (pos *Pos) IsWall(x, y int) bool {
	return x*(x-pos.Len+1) > 0 || y*(y-pos.Len+1) > 0 || (*pos.maps)[x*pos.Len+y] == 1 //out of index
}

//IsPort will return if is port
func (pos *Pos) IsPort() bool {
	return pos.X*(pos.X-pos.Len+1)*pos.Y*(pos.Y-pos.Len+1) == 0 && (*pos.maps)[pos.X*pos.Len+pos.Y] == 0 //wall
}

//IsPassed return if has passed
func (pos *Pos) IsPassed(x, y int) bool {
	return (*pos.maps)[x*pos.Len+y] == 10 //has passd
}

func (pos *Pos) setPassed() {
	(*pos.maps)[pos.X*pos.Len+pos.Y] = 10
}

func (pos *Pos) getNextPos() (re *Pos, flag bool) {
	switch pos.Dir {
	case 0:
		re, flag = pos.getNext(0, 1)
	case 1:
		re, flag = pos.getNext(1, 0)
	case 2:
		re, flag = pos.getNext(0, -1)
	case 3:
		re, flag = pos.getNext(-1, 0)
	}
	pos.Dir++
	return
}

func (pos *Pos) getNext(DirX, DirY int) (re *Pos, flag bool) {
	if !pos.IsWall(pos.X+DirX, pos.Y+DirY) &&
		!pos.IsPassed(pos.X+DirX, pos.Y+DirY) {
		re = &Pos{
			pos.X + DirX,
			pos.Y + DirY,
			0,
			pos.Len,
			pos.maps}
		flag = true
	}
	return
}

func getRouter(maps *[]int, Len int) (re *Pos, flag bool) {
	curpos := &Pos{1, 1, 0, Len, maps}
	curpos.setPassed()
	stackArr := stackArr.NewStack(Len * Len)
	stackArr.Push(curpos)
	for {
		curpos.showMap()
		if stackArr.IsEmpty() {
			break
		}
		curpos = stackArr.Pop().(*Pos)

		for curpos.Dir < 4 {
			if next, err := curpos.getNextPos(); err {
				if next.IsPort() {
					flag = true
					re = next
					return
				}
				stackArr.Push(curpos)
				stackArr.Push(next)

				curpos = next
				next.setPassed()
				break
			}
		}
		fmt.Println()
	}
	return
}

func main() {
	N := 7
	maps := initMap(N)

	if re, ok := getRouter(maps, N); ok {
		fmt.Printf("\ngate is %d-%d\n", re.X, re.Y)
	} else {
		fmt.Println("no gate")
	}
}
