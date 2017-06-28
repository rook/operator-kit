package main

import (
	"fmt"
	"os"

	"github.com/coreos/pkg/capnslog"
	opkit "github.com/rook/operator-kit"
)

const resourceGroup = "mycompany.io"

var logger = capnslog.NewPackageLogger("github.com/rook/operator-kit", "sample")

func main() {
	context := opkit.KubeContext{}

	// initialize the sample resource
	httpCli, err := opkit.NewHTTPClient(resourceGroup)
	if err != nil {
		fmt.Printf("failed to create http client. %+v\n", err)
		os.Exit(1)
	}
	context.KubeHTTPCli = httpCli.Client

	err = opkit.CreateCustomResources(context, []opkit.CustomResource{sampleResource})
	if err != nil {
		fmt.Printf("failed to create custom resource. %+v\n", err)
		os.Exit(1)
	}

	// start watching the sample resource
	logger.Infof("Managing the sample custom resources")
	mgr := &sampleManager{namespace: "default", context: context}
	mgr.Manage()
}
