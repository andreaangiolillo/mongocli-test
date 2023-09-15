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

package deployments

import (
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/search"
	"github.com/spf13/cobra"
)

var (
	errCompassNotInstalled = errors.New("MongoDB Compass not found in your system, to install MongoDB Compass follow these instructions: https://www.mongodb.com/docs/compass/current/install/")
	errMongoshNotInstalled = errors.New("mongosh not found in your system, to install MongoDB Shell follow these instructions: https://www.mongodb.com/docs/mongodb-shell/install")
)

func Builder() *cobra.Command {
	const use = "deployments"
	cmd := &cobra.Command{
		Hidden:  true,
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage Atlas and local deployments.",
	}

	cmd.AddCommand(
		SetupBuilder(),
		DeleteBuilder(),
		ListBuilder(),
		ConnectBuilder(),
		DiagnosticsBuilder(),
		search.Builder(),
	)

	return cmd
}