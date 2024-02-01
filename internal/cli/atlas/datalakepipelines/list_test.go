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

// This code was autogenerated at 2023-04-12T16:00:40+01:00. Note: Manual updates are allowed, but may be overwritten.

package datalakepipelines

import (
	"bytes"
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	mocks "github.com/andreangiolillo/mongocli-test/internal/mocks/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/pointer"
	"github.com/andreangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelinesLister(ctrl)

	expected := []atlasv2.DataLakeIngestionPipeline{
		{
			Id:    pointer.Get("2a5cbd92c036a0eb288"),
			Name:  pointer.Get("pipeline 2"),
			State: pointer.Get("PENDING"),
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		Pipelines(listOpts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID                    NAME         STATE
2a5cbd92c036a0eb288   pipeline 2   PENDING

`, buf.String())
	t.Log(buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
