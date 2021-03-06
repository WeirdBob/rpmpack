// Copyright 2019 Google LLC
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

// rpmsample creates an rpm file with some known files, which
// can be used to test rpmpack's output against other rpm implementations.
// It is also an instructive example for using rpmpack.
package main

import (
	"log"
	"os"

	"github.com/google/rpmpack"
)

func main() {

	r, err := rpmpack.NewRPM(rpmpack.RPMMetaData{
		Name:    "rpmsample",
		Version: "0.1",
		Release: "A",
		Arch:    "noarch",
	})
	if err != nil {
		log.Fatal(err)
	}
	r.AddFile(
		rpmpack.RPMFile{
			Name:  "/var/lib/rpmpack",
			Mode:  040755,
			Owner: "root",
			Group: "root",
		})
	r.AddFile(
		rpmpack.RPMFile{
			Name:  "/var/lib/rpmpack/sample.txt",
			Body:  []byte("testsample\n"),
			Mode:  0600,
			Owner: "root",
			Group: "root",
		})
	r.AddFile(
		rpmpack.RPMFile{
			Name:  "/var/lib/rpmpack/sample2.txt",
			Body:  []byte("testsample2\n"),
			Mode:  0644,
			Owner: "root",
			Group: "root",
		})
	r.AddFile(
		rpmpack.RPMFile{
			Name:  "/var/lib/rpmpack/sample3_link.txt",
			Body:  []byte("/var/lib/rpmpack/sample.txt"),
			Mode:  0120777,
			Owner: "root",
			Group: "root",
		})
	if err := r.Write(os.Stdout); err != nil {
		log.Fatalf("write failed: %v", err)
	}

}
