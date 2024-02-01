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

package filesystem

import (
	"context"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/store"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	store store.FileSystemsDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteFileSystem)
}

// // mongocli ops-manager admin backup fileSystem(s) delete <ID> [--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("File system configuration '%s' deleted\n", "File system configuration not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Aliases: []string{"rm"},
		Short:   "Delete a file system configuration.",
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Configuration identifier.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.initStore(cmd.Context())(); err != nil {
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

	return cmd
}
