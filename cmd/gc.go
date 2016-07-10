// Copyright © 2016 Ramit Surana <ramitsurana@gmail.com>
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

package cmd

import (
	"fmt"
	"os/exec"
	"log"	
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "wipes off all the stopped containers",
	Long: `cleans up all the stopped containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		cmd1 := exec.Command("docker", "rm", "$(docker", "ps", "-q", "-f", "status=exited)")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Failed to remove containers ...")
        	}		
		fmt.Println("All containers wiped")
	},
}

func init() {
	RootCmd.AddCommand(gcCmd)			
}
