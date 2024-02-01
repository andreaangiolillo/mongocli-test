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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelinesDescriber(ctrl)

	expected := &atlasv2.DataLakeIngestionPipeline{
		Id:    pointer.Get("1a5cbd92c036a0eb288"),
		Name:  pointer.Get("pipeline 1"),
		State: pointer.Get("PENDING"),
	}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		id:    "id",
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		Pipeline(describeOpts.ProjectID, describeOpts.id).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
