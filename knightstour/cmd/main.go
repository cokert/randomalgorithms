package main

import (
	"flag"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cokert/knightstour/board"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	szPtr := flag.Int("s", -1, "size of the board")
	xPtr := flag.Int("x", -1, "starting x position (0 <= x < size)")
	yPtr := flag.Int("y", -1, "starting y position (0 <= y < size)")

	flag.Parse()

	if *szPtr < 0 || *xPtr < 0 || *yPtr < 0 || *szPtr <= *xPtr || *szPtr <= *yPtr {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//initial location, create board, initial
	l := board.Location{X: *xPtr, Y: *yPtr}
	b := board.NewBoard(*szPtr, l)
	p := board.Path{}
	start := time.Now()
	successes, paths := search(b, l, p)
	duration := time.Since(start)

	pr := message.NewPrinter(language.English)
	pr.Printf("done size: %v\tStart: %v\tcomplete paths: %v\tsearched paths: %v\texecution time: %v\n", *szPtr, l, successes, paths, duration)
}

func search(b board.Board, l board.Location, p board.Path) (uint64, uint64) {
	p.Push(l)
	// this limits the amount of concurrency.  once we have a path over to positions long, we switch to a linear search.
	// if unbounded, it will spin up hundreds of thousands (millions?) of gofuncs which will exaust system memory on a large-ish board.
	// it's kind of fun to watch, actually.  on linux, your system will freeze for a bit before the program gets killed.
	// if you run it while an activity monitor is running, can see that the system page everything it can to swap after it
	// (idk if a mac or windows system will kill the process or hang infinitely).  to see it do this, comment the len check below and
	// set localConcurrent to true
	localConcurrent := false
	if len(p) < 3 {
		localConcurrent = true
	}
	b = b.Visit(l)
	mvs := b.PossibleMoves(p.Last())
	if len(mvs) == 0 {
		if b.AllVisited() {
			return 1, 1
		}
		return 0, 1
	}
	var successes uint64 = 0
	var totalPaths uint64 = 0
	wg := sync.WaitGroup{}
	for _, m := range mvs {
		if localConcurrent {
			wg.Add(1)
			go func(m board.Location) {
				nb := b.Clone()
				succ, paths := search(nb, m, p)
				wg.Done()
				atomic.AddUint64(&successes, succ)
				atomic.AddUint64(&totalPaths, paths)
			}(m)
		} else {
			nb := b.Clone()
			succ, p := search(nb, m, p)
			successes += succ
			totalPaths += p
		}
	}
	if localConcurrent {
		wg.Wait()
	}
	return successes, totalPaths
}
