/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mergelist_test

import (
	"flag"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/testgrid/pkg/merger"
)

var lists = flag.String("list", "../../mergelists/prod.yaml ../../mergelists/canary.yaml", "Space-delimited list of mergelists to test")

func TestMergelist(t *testing.T) {
	files := strings.Split(*lists, " ")

	for _, filename := range files {
		t.Run(filename, func(t *testing.T) {
			file, err := ioutil.ReadFile(filename)
			if err != nil {
				t.Errorf("Can't read file: %v", err)
			}

			_, err = merger.ParseAndCheck(file)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
