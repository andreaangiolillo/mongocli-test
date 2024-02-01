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

// This code was autogenerated at 2023-06-21T13:32:19+01:00. Note: Manual updates are allowed, but may be overwritten.

package datafederation

import (
	"context"
	"fmt"
	"io"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	store "github.com/andreangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type LogOpts struct {
	cli.GlobalOpts
	cli.DownloaderOpts
	store store.DataFederationLogDownloader
	id    string
	start int64
	end   int64
}

func (opts *LogOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *LogOpts) Run() error {
	f, err := opts.store.DataFederationLogs(opts.ConfigProjectID(), opts.id, opts.start, opts.end)
	if err != nil {
		return err
	}

	defer f.Close()

	out, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	return err
}

// atlas dataFederation log [--projectId projectId].
func LogBuilder() *cobra.Command {
	opts := &LogOpts{}
	cmd := &cobra.Command{
		Use:   "logs <name>",
		Short: "Returns logs of the specified data federation database for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the data federation database.",
		},
		Example: `# download logs of data federation database 'DataFederation1':
  atlas dataFederation logs DataFederation1
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().Int64Var(&opts.start, flag.Start, 0, usage.DataFederationStartDate)
	cmd.Flags().Int64Var(&opts.end, flag.End, 0, usage.DataFederationEndDate)

	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.LogOut)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)

	_ = cmd.MarkFlagRequired(flag.Out)
	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}
