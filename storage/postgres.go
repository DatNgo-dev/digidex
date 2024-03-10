package storage

import "github.com/DatNgo-dev/digidex/types"

type PostgresStorage struct{}

func (s *PostgresStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
