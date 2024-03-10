package storage

import "github.com/DatNgo-dev/digidex/types"

type Storage interface {
	Get(int) *types.User
}
