package main

import (
	"fmt"
	"math/rand"

	e2e "github.com/tendermint/tendermint/test/e2e/pkg"
)

func Generate(r *rand.Rand) []e2e.Manifest {
	manifest := e2e.Manifest{
		Nodes: map[string]e2e.ManifestNode{},
	}
	for i := 1; i <= 4; i++ {
		manifest.Nodes[fmt.Sprintf("validator%02d", i)] = generateNode(r, "validator")
	}

	return []e2e.Manifest{manifest}
}

func generateNode(r *rand.Rand, mode string) e2e.ManifestNode {
	databases := []string{"goleveldb", "cleveldb", "rocksdb", "boltdb", "badgerdb"}

	node := e2e.ManifestNode{
		Mode:     mode,
		Database: databases[r.Intn(len(databases))],
	}

	return node
}
