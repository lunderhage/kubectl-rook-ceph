/*
Copyright 2023 The Rook Authors. All rights reserved.

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

package exec

import (
	"os/exec"

	"github.com/rook/kubectl-rook-ceph/pkg/logging"
)

func ExecuteBashCommand(command string) string {
	cmd := exec.Command("/bin/bash",
		"-x", // Print commands and their arguments as they are executed
		"-e", // Exit immediately if a command exits with a non-zero status.
		"-m", // Terminal job control, allows job to be terminated by SIGTERM
		"-c", // Command to run
		command,
	)
	stdout, err := cmd.Output()
	if err != nil {
		logging.Fatal(err)
	}
	return string(stdout)
}
