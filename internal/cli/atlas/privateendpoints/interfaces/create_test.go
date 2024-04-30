// Copyright 2021 MongoDB Inc
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

package interfaces

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockInterfaceEndpointCreatorDeprecated(ctrl)

	createOpts := &CreateOpts{
		store:               mockStore,
		privateEndpointID:   "privateEndpointID",
		interfaceEndpointID: "interfaceEndpointID",
	}

	expected := &mongodbatlas.InterfaceEndpointConnectionDeprecated{}
	mockStore.
		EXPECT().
		CreateInterfaceEndpointDeprecated(createOpts.ProjectID, createOpts.privateEndpointID, createOpts.interfaceEndpointID).
		Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	require.NoError(t, err)
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.PrivateEndpointID},
	)
}
