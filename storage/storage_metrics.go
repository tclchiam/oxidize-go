package storage

import "github.com/tclchiam/oxidize-go/blockchain/entity"

type chainMetrics struct {
	entity.ChainRepository
}

func WrapWithMetrics(repository entity.ChainRepository) entity.ChainRepository {
	// TODO metrics stuff
	return &chainMetrics{
		ChainRepository: repository,
	}
}
