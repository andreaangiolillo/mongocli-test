// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ondemand

import (
	"context"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	store "github.com/andreaangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.CompliancePolicyDescriber
}

const describeTemplate = `ID	RETENTION
{{if .}}{{.Id}}	{{.RetentionValue}} {{.RetentionUnit}}{{end}}
`

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	raw, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return err
	}

	return opts.Print(raw.OnDemandPolicyItem)
}

func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Return the on-demand policy item of the backup compliance policy for your project.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
