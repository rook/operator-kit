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

package v1alpha1

import (
	v1alpha1 "github.com/rook/operator-kit/sample-operator/pkg/apis/myproject/v1alpha1"
	scheme "github.com/rook/operator-kit/sample-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SampleListsGetter has a method to return a SampleListInterface.
// A group's client should implement this interface.
type SampleListsGetter interface {
	SampleLists(namespace string) SampleListInterface
}

// SampleListInterface has methods to work with SampleList resources.
type SampleListInterface interface {
	Create(*v1alpha1.SampleList) (*v1alpha1.SampleList, error)
	Update(*v1alpha1.SampleList) (*v1alpha1.SampleList, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SampleList, error)
	List(opts v1.ListOptions) (*v1alpha1.SampleListList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SampleList, err error)
	SampleListExpansion
}

// sampleLists implements SampleListInterface
type sampleLists struct {
	client rest.Interface
	ns     string
}

// newSampleLists returns a SampleLists
func newSampleLists(c *MyprojectV1alpha1Client, namespace string) *sampleLists {
	return &sampleLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sampleList, and returns the corresponding sampleList object, and an error if there is any.
func (c *sampleLists) Get(name string, options v1.GetOptions) (result *v1alpha1.SampleList, err error) {
	result = &v1alpha1.SampleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samplelists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SampleLists that match those selectors.
func (c *sampleLists) List(opts v1.ListOptions) (result *v1alpha1.SampleListList, err error) {
	result = &v1alpha1.SampleListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samplelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sampleLists.
func (c *sampleLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("samplelists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sampleList and creates it.  Returns the server's representation of the sampleList, and an error, if there is any.
func (c *sampleLists) Create(sampleList *v1alpha1.SampleList) (result *v1alpha1.SampleList, err error) {
	result = &v1alpha1.SampleList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("samplelists").
		Body(sampleList).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sampleList and updates it. Returns the server's representation of the sampleList, and an error, if there is any.
func (c *sampleLists) Update(sampleList *v1alpha1.SampleList) (result *v1alpha1.SampleList, err error) {
	result = &v1alpha1.SampleList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("samplelists").
		Name(sampleList.Name).
		Body(sampleList).
		Do().
		Into(result)
	return
}

// Delete takes name of the sampleList and deletes it. Returns an error if one occurs.
func (c *sampleLists) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samplelists").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sampleLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samplelists").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sampleList.
func (c *sampleLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SampleList, err error) {
	result = &v1alpha1.SampleList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("samplelists").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
