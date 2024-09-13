package wyd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/cli"
)

type Program interface {
	Exit(code int)
	SetArgs(args []string)
	Args() []string
	GetLogger() *log.Logger
}

type Prog struct {
	args []string
	log  *log.Logger
}

func NewProg() *Prog {
	l := log.New(os.Stderr)
	l.SetPrefix("wyd")

	p := &Prog{
		log: l,
	}

	return p
}

func (p *Prog) Exit(code int) {
	os.Exit(code)
}

func (p *Prog) SetArgs(args []string) {
	p.args = args
}

func (p *Prog) Args() []string {
	return p.args
}

func (p *Prog) GetLogger() *log.Logger {
	return p.log
}

func Run(p Program, c cli.CliRunner) {
	l := p.GetLogger()
	if err := c.Run(p.Args()...); err != nil {
		l.Error(err)
		p.Exit(1)
		return
	}

	p.Exit(0)
}

func main() {
	p := NewProg()
	p.SetArgs(os.Args[1:])
	c := cli.New(p)
	Run(p, c)
}
