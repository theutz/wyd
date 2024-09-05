package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/cli"
)

type Program interface {
	Exit(code int)
	SetArgs(args []string)
	GetArgs() []string
}

type Prog struct {
	args []string
}

func NewProg() *Prog {
	p := &Prog{}
	return p
}

func (p *Prog) Exit(code int) {
	os.Exit(code)
}

func (p *Prog) SetArgs(args []string) {
	p.args = args
}

func (p *Prog) GetArgs() []string {
	return p.args
}

func Run(p Program, c cli.CLIer, log *log.Logger) {
	if err := c.Run(p.GetArgs()...); err != nil {
		log.Error(err)
		p.Exit(1)
		return
	}

	p.Exit(0)
}

func main() {
	p := NewProg()
	p.SetArgs(os.Args[1:])
	c := cli.New(p)
	log := log.New(os.Stderr)
	log.SetPrefix("wyd")
	Run(p, c, log)
}
