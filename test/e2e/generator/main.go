//nolint: gosec
package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	randomSeed int64 = 4827085738
)

var logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))

func main() {
	NewCLI().Run()
}

// CLI is the Cobra-based command-line interface.
type CLI struct {
	root *cobra.Command
}

// NewCLI sets up the CLI.
func NewCLI() *CLI {
	cli := &CLI{}
	cli.root = &cobra.Command{
		Use:           "generator",
		Short:         "End-to-end testnet generator",
		SilenceUsage:  true,
		SilenceErrors: true, // we'll output them ourselves in Run()
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			dir := filepath.Join("networks", "generated")
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}

			r := rand.New(rand.NewSource(randomSeed))
			manifests := Generate(r)
			for i, manifest := range manifests {
				err = manifest.Save(filepath.Join(dir, fmt.Sprintf("%v.toml", i)))
				if err != nil {
					return err
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cli
}

// Run runs the CLI.
func (cli *CLI) Run() {
	if err := cli.root.Execute(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
