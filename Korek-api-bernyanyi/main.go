package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Korek struct {
	hits       int
	lastPlayer string
}

// break point
const BreakPoint = 11

func main() {
	korek := make(chan *Korek)

	// channel untuk menangkap data jika sudah selesai
	done := make(chan *Korek)

	players := []string{"player 1", "player 2", "player 3", "player 4"}
	for _, p := range players {
		// tambahin kedalam parameter
		go play(p, korek, done)
	}
	korek <- new(Korek)

	// tambahin function finish
	finish(done)
}

func play(name string, korek, done chan *Korek) {
	// membuat angka random
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	for {
		select {
		case k := <-korek:
			// mengambil angka random
			v := rand.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)
			k.hits++
			k.lastPlayer = name
			fmt.Println("korek ada di", k.lastPlayer, "pada hit ke", k.hits, "dan mempunyai nilai", v)

			// tambahan
			if v%BreakPoint == 0 {
				// jika oke, maka akan mengirim value ke channel `done`
				done <- k

				// return akan memberhentikan perulangan
				return
			}

			korek <- k
		}
	}
}

func finish(done chan *Korek) {
	for {
		select {
		// jika ada data yang masuk pada channel done,
		// maka game selesai
		case d := <-done:
			fmt.Println(d.lastPlayer, "kalah pada hit ke", d.hits)
			return
		}
	}
}
