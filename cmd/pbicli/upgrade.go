package main

import (
	"github.com/spf13/cobra"
	"github.com/squillace/porter-pbicli/pkg/pbicli"
)

func buildUpgradeCommand(m *pbicli.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
