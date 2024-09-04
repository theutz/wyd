package main

import "os"

type Program interface {
	Exit(code int)
}

type Prog struct{}

func (p *Prog) Exit(code int) {
	os.Exit(code)
}

func Run(p Program) {
	p.Exit(1)
}

func main() {
	p := &Prog{}
	Run(p)
}
