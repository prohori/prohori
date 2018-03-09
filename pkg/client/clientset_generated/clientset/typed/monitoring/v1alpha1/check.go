/*
Copyright 2018 The Prohori Authors.

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
	v1alpha1 "github.com/prohori/prohori/pkg/apis/monitoring/v1alpha1"
	scheme "github.com/prohori/prohori/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChecksGetter has a method to return a CheckInterface.
// A group's client should implement this interface.
type ChecksGetter interface {
	Checks(namespace string) CheckInterface
}

// CheckInterface has methods to work with Check resources.
type CheckInterface interface {
	Create(*v1alpha1.Check) (*v1alpha1.Check, error)
	Update(*v1alpha1.Check) (*v1alpha1.Check, error)
	UpdateStatus(*v1alpha1.Check) (*v1alpha1.Check, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Check, error)
	List(opts v1.ListOptions) (*v1alpha1.CheckList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Check, err error)
	CheckExpansion
}

// checks implements CheckInterface
type checks struct {
	client rest.Interface
	ns     string
}

// newChecks returns a Checks
func newChecks(c *MonitoringV1alpha1Client, namespace string) *checks {
	return &checks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the check, and returns the corresponding check object, and an error if there is any.
func (c *checks) Get(name string, options v1.GetOptions) (result *v1alpha1.Check, err error) {
	result = &v1alpha1.Check{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("checks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Checks that match those selectors.
func (c *checks) List(opts v1.ListOptions) (result *v1alpha1.CheckList, err error) {
	result = &v1alpha1.CheckList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("checks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested checks.
func (c *checks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("checks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a check and creates it.  Returns the server's representation of the check, and an error, if there is any.
func (c *checks) Create(check *v1alpha1.Check) (result *v1alpha1.Check, err error) {
	result = &v1alpha1.Check{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("checks").
		Body(check).
		Do().
		Into(result)
	return
}

// Update takes the representation of a check and updates it. Returns the server's representation of the check, and an error, if there is any.
func (c *checks) Update(check *v1alpha1.Check) (result *v1alpha1.Check, err error) {
	result = &v1alpha1.Check{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("checks").
		Name(check.Name).
		Body(check).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *checks) UpdateStatus(check *v1alpha1.Check) (result *v1alpha1.Check, err error) {
	result = &v1alpha1.Check{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("checks").
		Name(check.Name).
		SubResource("status").
		Body(check).
		Do().
		Into(result)
	return
}

// Delete takes name of the check and deletes it. Returns an error if one occurs.
func (c *checks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("checks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *checks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("checks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched check.
func (c *checks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Check, err error) {
	result = &v1alpha1.Check{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("checks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
