package ui

import (
	"github.com/rollicks-c/kgate/internal/logic/model"
	"github.com/rollicks-c/kgate/internal/logic/ui/fancy"
	"github.com/rollicks-c/kgate/internal/logic/ui/simple"
)

func NewSimple() model.Frontend {
	return simple.New()
}

func NewFancy() model.Frontend {
	return fancy.New()
}
