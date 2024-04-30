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

package metrics

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestDatabasesDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockHostDatabaseMeasurementsLister(ctrl)

	listOpts := &DatabasesDescribeOpts{
		hostID: "1",
		name:   "test",
		store:  mockStore,
	}

	opts := listOpts.NewProcessMetricsListOptions()
	expected := &opsmngr.ProcessDatabaseMeasurements{
		ProcessMeasurements: &opsmngr.ProcessMeasurements{
			Measurements: []*opsmngr.Measurements{},
		},
	}
	mockStore.
		EXPECT().HostDatabaseMeasurements(listOpts.ProjectID, listOpts.hostID, listOpts.name, opts).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, databasesMetricTemplate, expected)
}
