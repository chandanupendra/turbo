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

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "cleans up all your docker images",
	Long: `clean up your docker images `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Detecting images ...")
		cmd1 := exec.Command("docker", "rmi", "`docker", "images", "-a", "-q`")
		err1 := cmd1.Start()
		if err1 != nil {
                	log.Fatal(err1)
			log.Printf("Images removal failed ...")
        	}		
		fmt.Println("All images cleaned")
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)		
}
