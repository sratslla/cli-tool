/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the benchmark process",
	Long:  `Start the benchmark process`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		if isKubernetesClusterRunning() {
			fmt.Println("Kubernetes cluster is running")
			// Proceed with your logic for starting services/pods
		} else {
			fmt.Println("Kubernetes cluster is not running or accessible")
			// Handle the case where Kubernetes cluster is not available
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
func isKubernetesClusterRunning() bool {
	// Execute kubectl command to get cluster info
	cmd := exec.Command("kubectl", "cluster-info")

	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		// Handle error, indicating Kubernetes cluster is not running or accessible
		return false
	}
	// If the command executed successfully, assume Kubernetes cluster is running
	fmt.Println(cmd, output.String())
	return true
}
