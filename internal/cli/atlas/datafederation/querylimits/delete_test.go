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

//go:build unit

// This code was autogenerated at 2023-06-23T15:50:57+01:00. Note: Manual updates are allowed, but may be overwritten.

package querylimits

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataFederationQueryLimitDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		store:      mockStore,
		tenantName: "DataFederation1",
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "id",
			Confirm: true,
		},
	}

	mockStore.
		EXPECT().
		DeleteDataFederationQueryLimit(deleteOpts.ProjectID, deleteOpts.tenantName, deleteOpts.Entry).
		Return(nil).
		Times(1)

	if err := deleteOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DeleteBuilder(),
		0,
		[]string{flag.ProjectID},
	)
}