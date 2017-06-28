package main

import (
	"encoding/json"
	"fmt"
	"time"

	opkit "github.com/rook/operator-kit"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/pkg/api/v1"
)

type sampleManager struct {
	namespace string
	context   opkit.KubeContext
}

type sampleEvent struct {
	Type   watch.EventType
	Object *Sample
}

type Sample struct {
	v1.ObjectMeta `json:"metadata,omitempty"`
	Spec          SampleSpec `json:"spec"`
}

type SampleSpec struct {
	Hello string `json:"hello"`
}

var sampleResource = opkit.CustomResource{
	Name:        "sample",
	Group:       resourceGroup,
	Version:     opkit.V1Alpha1,
	Description: "A very simple example custom resource",
}

// Enter a control loop to watch for changes to the custom resource
func (s *sampleManager) Manage() {
	for {
		// load the existing sample instances
		resourceVersion, err := s.Load()
		if err != nil {
			logger.Errorf("cannot load samples. %+v. retrying...", err)
		} else {
			// watch for added/updated/deleted samples
			watcher := opkit.NewWatcher(s.context, sampleResource, s.namespace, resourceVersion, s.handleSampleEvent, nil)
			if err := watcher.Watch(); err != nil {
				logger.Errorf("failed to watch sample tpr. %+v. retrying...", err)
			}
		}

		<-time.After(time.Second * time.Duration(s.context.RetryDelay))
	}
}

func (s *sampleManager) handleSampleEvent(event *opkit.RawEvent) error {
	sampleEv := &sampleEvent{
		Type:   event.Type,
		Object: &Sample{},
	}
	err := json.Unmarshal(event.Object, sampleEv.Object)
	if err != nil {
		return fmt.Errorf("failed to unmarshal sample from data (%s): %v", sampleEv.Object, err)
	}

	sample := sampleEv.Object
	switch event.Type {
	case watch.Added:
		logger.Infof("Added Sample '%s' with Hello=%s!", sample.Name, sample.Spec.Hello)

	case watch.Modified:
		logger.Infof("Modified Sample '%s' with Hello=%s!", sample.Name, sample.Spec.Hello)

	case watch.Deleted:
		logger.Infof("Deleted Sample '%s'!", sample.Spec.Hello)
	}
	return nil
}

func (p *sampleManager) Load() (string, error) {
	logger.Info("finding existing samples...")
	sampleList, err := p.getSampleList()
	if err != nil {
		return "", err
	}

	logger.Infof("found %d samples.", len(sampleList.Items))
	return sampleList.Metadata.ResourceVersion, nil
}

func (s *sampleManager) getSampleList() (*SampleList, error) {
	b, err := opkit.GetRawListNamespaced(s.context.Clientset, sampleResource, s.namespace)
	if err != nil {
		return nil, err
	}

	samples := &SampleList{}
	if err := json.Unmarshal(b, samples); err != nil {
		return nil, err
	}
	return samples, nil
}
