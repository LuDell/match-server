package server

import "match-server/model"

type Persist int

var (
	bidOrder *model.Order
	askOrder *model.Order
	bidPosition *model.Position
	askPosition *model.Position
)

func (p *Persist)DataPersistence() error {
	
}