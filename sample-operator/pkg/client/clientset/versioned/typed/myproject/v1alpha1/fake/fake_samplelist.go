/*
Copyright 2017 The Kubernetes Authors All rights reserved.

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

package fake

import (
	v1alpha1 "github.com/rook/operator-kit/sample-operator/pkg/apis/myproject/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSampleLists implements SampleListInterface
type FakeSampleLists struct {
	Fake *FakeMyprojectV1alpha1
	ns   string
}

var samplelistsResource = schema.GroupVersionResource{Group: "myproject", Version: "v1alpha1", Resource: "samplelists"}

var samplelistsKind = schema.GroupVersionKind{Group: "myproject", Version: "v1alpha1", Kind: "SampleList"}

// Get takes name of the sampleList, and returns the corresponding sampleList object, and an error if there is any.
func (c *FakeSampleLists) Get(name string, options v1.GetOptions) (result *v1alpha1.SampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(samplelistsResource, c.ns, name), &v1alpha1.SampleList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleList), err
}

// List takes label and field selectors, and returns the list of SampleLists that match those selectors.
func (c *FakeSampleLists) List(opts v1.ListOptions) (result *v1alpha1.SampleListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(samplelistsResource, samplelistsKind, c.ns, opts), &v1alpha1.SampleListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleListList), err
}

// Watch returns a watch.Interface that watches the requested sampleLists.
func (c *FakeSampleLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(samplelistsResource, c.ns, opts))

}

// Create takes the representation of a sampleList and creates it.  Returns the server's representation of the sampleList, and an error, if there is any.
func (c *FakeSampleLists) Create(sampleList *v1alpha1.SampleList) (result *v1alpha1.SampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(samplelistsResource, c.ns, sampleList), &v1alpha1.SampleList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleList), err
}

// Update takes the representation of a sampleList and updates it. Returns the server's representation of the sampleList, and an error, if there is any.
func (c *FakeSampleLists) Update(sampleList *v1alpha1.SampleList) (result *v1alpha1.SampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(samplelistsResource, c.ns, sampleList), &v1alpha1.SampleList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleList), err
}

// Delete takes name of the sampleList and deletes it. Returns an error if one occurs.
func (c *FakeSampleLists) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(samplelistsResource, c.ns, name), &v1alpha1.SampleList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSampleLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(samplelistsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SampleListList{})
	return err
}

// Patch applies the patch and returns the patched sampleList.
func (c *FakeSampleLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(samplelistsResource, c.ns, name, data, subresources...), &v1alpha1.SampleList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SampleList), err
}
