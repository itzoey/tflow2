// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package annotator

import (
	"testing"
	"tflow2/netflow"
)

func TestTimestampAggr(t *testing.T) {
	ca := make(chan *netflow.Flow)
	cb := make(chan *netflow.Flow)
	var aggr int64 = 60
	go Init(ca, cb, aggr, false, 1)

	testData := []struct {
		ts   int64
		want int64
	}{
		{
			ts:   1000,
			want: 960,
		},
		{
			ts:   1234,
			want: 1200,
		},
	}

	for _, test := range testData {
		fl := &netflow.Flow{
			Timestamp: test.ts,
		}

		ca <- fl
		fl = <-cb
		if fl.Timestamp != test.want {
			t.Errorf("Input: %d, Got: %d, Expected: %d, ", test.ts, fl.Timestamp, test.want)
		}
	}
}
