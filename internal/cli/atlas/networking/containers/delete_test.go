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

package containers

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockContainersDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "to_delete",
			Confirm: true,
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DeleteContainer(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(nil).
		Times(1)

	err := deleteOpts.Run()
	require.NoError(t, err)
}
