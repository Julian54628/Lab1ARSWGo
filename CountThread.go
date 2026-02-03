package main

import "time"

type CountThread struct {
	start int
	end   int
}

func (ct *CountThread) Run() {
	for i := ct.start; i <= ct.end && i < ct.start+10; i++ {
		print(i, " ")
		time.Sleep(100 * time.Millisecond)
	}
	println()
}
