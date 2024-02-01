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

// This code was autogenerated at 2023-04-12T16:00:41+01:00. Note: Manual updates are allowed, but may be overwritten.

package datalakepipelines

import (
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/flag"
	mocks "github.com/andreangiolillo/mongocli-test/internal/mocks/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelinesUpdater(ctrl)

	var expected *atlasv2.DataLakeIngestionPipeline

	updateOpts := &UpdateOpts{
		pipelineName: "id",
		store:        mockStore,
	}

	request, err := updateOpts.newUpdateRequest()
	require.NoError(t, err)

	mockStore.
		EXPECT().
		UpdatePipeline(updateOpts.ConfigProjectID(), updateOpts.pipelineName, *request).Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
