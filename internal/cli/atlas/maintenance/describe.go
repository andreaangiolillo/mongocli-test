// Copyright 2020 MongoDB Inc
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

package maintenance

import (
	"context"
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

var describeTemplate = `DAY OF THE WEEK	HOUR OF DAY	START ASAP
{{.DayOfWeek}}	{{.HourOfDay}}	{{.StartASAP}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store store.MaintenanceWindowDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.MaintenanceWindow(opts.ConfigProjectID())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// mongocli atlas maintenanceWindow(s) describe|get [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Short:       "Return the maintenance window details for your project.",
		Long:        longDesc + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{"output": describeTemplate},
		Example: fmt.Sprintf(`  # Return the maintenance window for the project with the ID 5e2211c17a3e5a48f5497de3:
  %s maintenanceWindows describe --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleAtlasEntryPoint()),
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
