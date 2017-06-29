package main

import (
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	opkit "github.com/rook/operator-kit"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const resourceGroup = "mycompany.io"

var logger = capnslog.NewPackageLogger("github.com/rook/operator-kit", "sample")

func main() {
	logger.Infof("Getting kubernetes context")
	context, err := createContext()
	if err != nil {
		logger.Errorf("failed to create context. %+v\n", err)
		os.Exit(1)
	}

	logger.Infof("Creating the sample resource")
	err = opkit.CreateCustomResources(*context, []opkit.CustomResource{sampleResource})
	if err != nil {
		logger.Errorf("failed to create custom resource. %+v\n", err)
		os.Exit(1)
	}

	// start watching the sample resource
	logger.Infof("Managing the sample resource")
	mgr := &sampleManager{namespace: "default", context: *context}
	mgr.Manage()
}

func createContext() (*opkit.KubeContext, error) {
	// create the k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	// initialize the sample resource
	httpCli, err := opkit.NewHTTPClient(resourceGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create http client. %+v", err)
	}

	return &opkit.KubeContext{
		MasterHost:  config.Host,
		Clientset:   clientset,
		KubeHTTPCli: httpCli.Client,
		RetryDelay:  6,
		MaxRetries:  15,
	}, nil
}
