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

package policies

import (
	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/backup/compliancepolicy/policies/ondemand"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/atlas/backup/compliancepolicy/policies/scheduled"
	"github.com/spf13/cobra"
)

func baseCommand() *cobra.Command {
	const use = "policies"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage the individual policy items of the backup compliance policy for your project.",
	}

	return cmd
}

func Builder() *cobra.Command {
	cmd := baseCommand()

	cmd.AddCommand(
		DescribeBuilder(),
		ondemand.Builder(),
		scheduled.Builder(),
	)

	return cmd
}