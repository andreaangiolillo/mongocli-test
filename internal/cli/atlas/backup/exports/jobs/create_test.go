// Copyright 2022 MongoDB Inc
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

package jobs

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockExportJobsCreator(ctrl)

	expected := &atlasv2.DiskBackupExportJob{}

	createOpts := &CreateOpts{
		exportBucketID: "12345678f901a234dbdb00ca",
		clusterName:    "test-cluster",
		snapshotID:     "dbdb00ca12345678f901a234",
		customData: map[string]string{
			"name": "test",
		},
		store: mockStore,
	}

	job := createOpts.newExportJob()
	mockStore.
		EXPECT().
		CreateExportJob("", createOpts.clusterName, job).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, createTemplate, &atlasv2.DiskBackupExportJob{})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.SnapshotID, flag.BucketID, flag.ClusterName, flag.CustomData, flag.ProjectID, flag.Output},
	)
}
