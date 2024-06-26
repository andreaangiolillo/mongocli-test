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

package projects

import (
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/projects/apikeys"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/projects/invitations"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/projects/settings"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/projects/teams"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/projects/users"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "projects"
	cmd := &cobra.Command{
		Use:     use,
		Short:   fmt.Sprintf("Manage your %s projects.", cli.DescriptionServiceName()),
		Long:    "Create, list and manage your MongoDB projects.",
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		ListBuilder(),
		CreateBuilder(),
		DeleteBuilder(),
		DescribeBuilder(),
		apikeys.Builder(),
		users.Builder(),
		teams.Builder(),
		invitations.Builder(),
		settings.Builder(),
	)

	return cmd
}
