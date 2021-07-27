/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/globalservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalServices implements GlobalServiceInterface
type FakeGlobalServices struct {
	Fake *FakeAzureglobalservicesV1alpha1
	ns   string
}

var globalservicesResource = schema.GroupVersionResource{Group: "azureglobalservices.networking.aks.io", Version: "v1alpha1", Resource: "globalservices"}

var globalservicesKind = schema.GroupVersionKind{Group: "azureglobalservices.networking.aks.io", Version: "v1alpha1", Kind: "GlobalService"}

// Get takes name of the globalService, and returns the corresponding globalService object, and an error if there is any.
func (c *FakeGlobalServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GlobalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globalservicesResource, c.ns, name), &v1alpha1.GlobalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalService), err
}

// List takes label and field selectors, and returns the list of GlobalServices that match those selectors.
func (c *FakeGlobalServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GlobalServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globalservicesResource, globalservicesKind, c.ns, opts), &v1alpha1.GlobalServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlobalServiceList{ListMeta: obj.(*v1alpha1.GlobalServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlobalServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalServices.
func (c *FakeGlobalServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globalservicesResource, c.ns, opts))

}

// Create takes the representation of a globalService and creates it.  Returns the server's representation of the globalService, and an error, if there is any.
func (c *FakeGlobalServices) Create(ctx context.Context, globalService *v1alpha1.GlobalService, opts v1.CreateOptions) (result *v1alpha1.GlobalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globalservicesResource, c.ns, globalService), &v1alpha1.GlobalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalService), err
}

// Update takes the representation of a globalService and updates it. Returns the server's representation of the globalService, and an error, if there is any.
func (c *FakeGlobalServices) Update(ctx context.Context, globalService *v1alpha1.GlobalService, opts v1.UpdateOptions) (result *v1alpha1.GlobalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globalservicesResource, c.ns, globalService), &v1alpha1.GlobalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalServices) UpdateStatus(ctx context.Context, globalService *v1alpha1.GlobalService, opts v1.UpdateOptions) (*v1alpha1.GlobalService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(globalservicesResource, "status", c.ns, globalService), &v1alpha1.GlobalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalService), err
}

// Delete takes name of the globalService and deletes it. Returns an error if one occurs.
func (c *FakeGlobalServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(globalservicesResource, c.ns, name), &v1alpha1.GlobalService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globalservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlobalServiceList{})
	return err
}

// Patch applies the patch and returns the patched globalService.
func (c *FakeGlobalServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GlobalService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlobalService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlobalService), err
}
