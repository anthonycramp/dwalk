package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
)

type Point struct {
	x, y int64
}

func dwalk(N int, pos Point, directions []*Point, ch chan int64) {
	start := pos
	for i := 0; i < N; i++ {
		dir := directions[rand.Intn(len(directions))]
		pos.x += dir.x
		pos.y += dir.y
	}
	deltaX := pos.x - start.x
	deltaY := pos.y - start.y
	ch<- deltaX*deltaX + deltaY*deltaY
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s N T\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tN is the number of steps to take\n")
		fmt.Fprintf(os.Stderr, "\tT is the number of Monte Carlo iterations.\n")
		os.Exit(-1)
	}
	
	N,_ := strconv.Atoi(os.Args[1])
	T,_ := strconv.Atoi(os.Args[2])
	ch := make(chan int64)
	
	directions := []*Point{&Point{0,1},&Point{1,0},&Point{0,-1},&Point{-1,0}}
	for i := 0; i < T; i++ {
		go dwalk(N, Point{0,0}, directions, ch)
	}
	
	sum := int64(0)
	for i := 0; i < T; i++ {
		sum += <-ch
	}
	fmt.Println("Mean squared distance: ", float64(sum)/float64(T))
}
