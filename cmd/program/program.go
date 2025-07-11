package program

import (
	"github.com/FrancoMusolino/film-cli/cmd/flags"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/nicksnyder/go-i18n/i18n"
)

type Program struct {
	Lang          flags.Lang
	Translate     i18n.TranslateFunc
	MoviesService movies.MoviesService
	StepChan      chan int
	DoneChan      chan bool
}

func (p *Program) Terminate() {
	close(p.StepChan)
	close(p.DoneChan)
}
