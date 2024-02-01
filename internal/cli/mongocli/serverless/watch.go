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

package serverless

import (
	"context"
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	name  string
	store store.ServerlessInstanceDescriber
}

func (opts *WatchOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *WatchOpts) watcher() (bool, error) {
	result, err := opts.store.GetServerlessInstance(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return false, err
	}
	const desiredState = "IDLE"
	return *result.StateName == desiredState, nil
}

func (opts *WatchOpts) Run() error {
	if err := opts.Watch(opts.watcher); err != nil {
		return err
	}

	return opts.Print(nil)
}

// mongocli atlas serverless|sl watch <instanceName> [--projectId projectId].
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <instanceName>",
		Short: "Monitor the status of serverless instance.",
		Long: `This command checks the serverless instance's state periodically until the instance reaches an IDLE state. 
Once the instance reaches the expected state, the command prints "Instance available."
If you run the command in the terminal, it blocks the terminal session until the resource becomes idle.
You can interrupt the command's polling at any time with CTRL-C.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  %s serverless watch instanceNameSample`, cli.ExampleAtlasEntryPoint()),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"instanceNameDesc": "Name of the instance to watch.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), "\nInstance available.\n"),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
