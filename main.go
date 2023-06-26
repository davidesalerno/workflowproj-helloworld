package main

import (
	"os"

	"github.com/kiegroup/kogito-serverless-operator/workflowproj"
)

func main() {
	// we are ignoring errors just for demo purposes, but don't do this!
	workflowFile, _ := os.Open("greet.sw.json")
	defer workflowFile.Close()

	// create the handler
	handler := workflowproj.New("mynamespace").
		WithWorkflow(workflowFile)

	// ... or you can save the files locally to use them later or to integrate in a GitOps process
	_ = handler.SaveAsKubernetesManifests("./output")
}
