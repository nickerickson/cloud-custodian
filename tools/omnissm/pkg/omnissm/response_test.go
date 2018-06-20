// Copyright 2018 Capital One Services, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package omnissm_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/capitalone/cloud-custodian/tools/omnissm/pkg/omnissm"
	"github.com/capitalone/cloud-custodian/tools/omnissm/pkg/store"
)

func TestRegistrationResponseMarshal(t *testing.T) {
	testCases := []struct {
		input    *omnissm.RegistrationResponse
		expected []byte
	}{
		{
			input: &omnissm.RegistrationResponse{
				RegistrationEntry: store.RegistrationEntry{
					Id:             "id",
					ActivationId:   "aid",
					ActivationCode: "code",
					ManagedId:      "managedId",
				},
				Region: "us-east-2",
			},
			expected: []byte(`{"id":"id","ActivationId":"aid","ActivationCode":"code","ManagedId":"managedId","region":"us-east-2"}`),
		},
	}

	for i, tc := range testCases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(string(data), string(tc.expected)); diff != "" {
			t.Errorf("TestCase %d: after json.Marshal: (-got +want)\n%s", i, diff)
		}
	}
}