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

// This code was autogenerated at 2023-07-05T01:21:22Z. Note: Manual updates are allowed, but may be overwritten.

package connection

import (
	"context"
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/store"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	streamsInstance string
	store           store.ConnectionDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteConnection, opts.ConfigProjectID(), opts.streamsInstance)
}

// atlas streams connection delete <connectionName> [--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Atlas Stream Processing connection '%s' deleted\n", "Atlas Stream Processing connection not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <connectionName>",
		Short: "Remove the specified Atlas Stream Processing connection from your project.",
		Long: `The command prompts you to confirm the operation when you run the command without the --force option.

Before deleting an Atlas Streams Processing connection, you must first stop all processes associated with it. ` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"connectionNameDesc": "Name of the connection",
			"output":             opts.SuccessMessage(),
		},
		Example: fmt.Sprintf(`# deletes connection 'ExampleConnection' from instance 'ExampleInstance':
  %[1]s streams connection delete ExampleConnection --instance ExampleInstance

# deletes connection 'ExampleConnection' from instance 'ExampleInstance' without requiring confirmation:
  %[1]s streams connection delete ExampleConnection --instance ExampleInstance --force
`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				cmd.ValidateRequiredFlags,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}