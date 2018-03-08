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
package fake

import (
	monitoring "github.com/prohori/prohori/pkg/apis/monitoring"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChecks implements CheckInterface
type FakeChecks struct {
	Fake *FakeMonitoring
	ns   string
}

var checksResource = schema.GroupVersionResource{Group: "monitoring.prohori", Version: "", Resource: "checks"}

var checksKind = schema.GroupVersionKind{Group: "monitoring.prohori", Version: "", Kind: "Check"}

// Get takes name of the check, and returns the corresponding check object, and an error if there is any.
func (c *FakeChecks) Get(name string, options v1.GetOptions) (result *monitoring.Check, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(checksResource, c.ns, name), &monitoring.Check{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring.Check), err
}

// List takes label and field selectors, and returns the list of Checks that match those selectors.
func (c *FakeChecks) List(opts v1.ListOptions) (result *monitoring.CheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(checksResource, checksKind, c.ns, opts), &monitoring.CheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoring.CheckList{}
	for _, item := range obj.(*monitoring.CheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested checks.
func (c *FakeChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(checksResource, c.ns, opts))

}

// Create takes the representation of a check and creates it.  Returns the server's representation of the check, and an error, if there is any.
func (c *FakeChecks) Create(check *monitoring.Check) (result *monitoring.Check, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(checksResource, c.ns, check), &monitoring.Check{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring.Check), err
}

// Update takes the representation of a check and updates it. Returns the server's representation of the check, and an error, if there is any.
func (c *FakeChecks) Update(check *monitoring.Check) (result *monitoring.Check, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(checksResource, c.ns, check), &monitoring.Check{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring.Check), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChecks) UpdateStatus(check *monitoring.Check) (*monitoring.Check, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(checksResource, "status", c.ns, check), &monitoring.Check{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring.Check), err
}

// Delete takes name of the check and deletes it. Returns an error if one occurs.
func (c *FakeChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(checksResource, c.ns, name), &monitoring.Check{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(checksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &monitoring.CheckList{})
	return err
}

// Patch applies the patch and returns the patched check.
func (c *FakeChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring.Check, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(checksResource, c.ns, name, data, subresources...), &monitoring.Check{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring.Check), err
}
