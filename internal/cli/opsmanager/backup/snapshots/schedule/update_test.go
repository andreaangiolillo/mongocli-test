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

package schedule

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSnapshotScheduleUpdater(ctrl)

	expected := &opsmngr.SnapshotSchedule{}

	opts := &UpdateOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().UpdateSnapshotSchedule(opts.ConfigProjectID(), opts.clusterID, opts.newSnapshotSchedule()).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.ClusterID, flag.SnapshotRetentionDays, flag.DailySnapshotRetentionDays, flag.WeeklySnapshotRetentionWeeks,
			flag.PointInTimeWindowHours, flag.ReferenceHourOfDay, flag.ReferenceMinuteOfHour, flag.ReferenceTimeZoneOffset, flag.SnapshotIntervalHours,
			flag.ClusterCheckpointIntervalMin,
		},
	)
}
