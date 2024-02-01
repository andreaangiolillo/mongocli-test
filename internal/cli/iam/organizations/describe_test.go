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

//go:build unit

package organizations

import (
	"bytes"
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/mocks"
	"github.com/andreangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationDescriber(ctrl)
	stringVal := "test"
	expected := &opsmngr.Organization{
		Links: nil,
		ID:    stringVal,
		Name:  stringVal,
	}
	buf := new(bytes.Buffer)

	mockStore.
		EXPECT().
		Organization(gomock.Eq("5a0a1e7e0f2912c554080adc")).
		Return(expected, nil).
		Times(1)

	opts := &DescribeOpts{
		store: mockStore,
		id:    "5a0a1e7e0f2912c554080adc",
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, describeTemplate, expected)
	assert.Equal(t, `ID     NAME
test   test
`, buf.String())
	t.Log(buf.String())
}
