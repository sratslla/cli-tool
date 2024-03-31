/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
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
		REPO_URL := "https://raw.githubusercontent.com/sratslla/cli-tool/main"
		YAML_FILE_PATH := "manifests/kubernetes-manifests.yaml"
		YAML_FILE_PATH2 := "manifests/loadgenerator_ui.yaml"
		err := applyManifestFromGitHub(REPO_URL, YAML_FILE_PATH)
		if err != nil {
			fmt.Println("Error applying manifest:", err)
			os.Exit(1)
		}
		err2 := applyManifestFromGitHub(REPO_URL, YAML_FILE_PATH2)
		if err2 != nil {
			fmt.Println("Error applying manifest:", err)
			os.Exit(1)
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

func applyManifestFromGitHub(repoURL, yamlFilePath string) error {
	cmd := exec.Command("kubectl", "apply", "-f", fmt.Sprintf("%s/%s", repoURL, yamlFilePath))
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error applying manifest: %v\n%s", err, output)
	}
	fmt.Println("Manifest applied successfully.", output.String())
	return nil
}
