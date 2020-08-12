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

package privateendpoints

import (
	"fmt"
	"time"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type WatchOpts struct {
	cli.GlobalOpts
	id    string
	store store.PrivateEndpointDescriber
}

func (opts *WatchOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

const defaultWait = 4 * time.Second

func (opts *WatchOpts) Run() error {
	for {
		result, err := opts.store.PrivateEndpoint(opts.ConfigProjectID(), opts.id)
		if err != nil {
			return err
		}
		if result.Status == "WAITING_FOR_USER" || result.Status == "FAILED" {
			fmt.Printf("\nPrivate endpoint changes completed.\n")
			break
		}
		fmt.Print(".")
		time.Sleep(defaultWait)
	}

	return nil
}

// mongocli atlas cluster(s) watch <name> [--projectId projectId]
func WatchBuilder() *cobra.Command {
	opts := &WatchOpts{}
	cmd := &cobra.Command{
		Use:   "watch <name>",
		Short: watchPrivateEndpoint,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	return cmd
}