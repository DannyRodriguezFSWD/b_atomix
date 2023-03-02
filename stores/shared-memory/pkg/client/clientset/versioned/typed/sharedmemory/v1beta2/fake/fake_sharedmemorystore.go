// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta2 "github.com/atomix/atomix/stores/shared-memory/pkg/apis/sharedmemory/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSharedMemoryStores implements SharedMemoryStoreInterface
type FakeSharedMemoryStores struct {
	Fake *FakeSharedmemoryV1beta2
	ns   string
}

var sharedmemorystoresResource = schema.GroupVersionResource{Group: "sharedmemory.atomix.io", Version: "v1beta2", Resource: "sharedmemorystores"}

var sharedmemorystoresKind = schema.GroupVersionKind{Group: "sharedmemory.atomix.io", Version: "v1beta2", Kind: "SharedMemoryStore"}

// Get takes name of the sharedMemoryStore, and returns the corresponding sharedMemoryStore object, and an error if there is any.
func (c *FakeSharedMemoryStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.SharedMemoryStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sharedmemorystoresResource, c.ns, name), &v1beta2.SharedMemoryStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SharedMemoryStore), err
}

// List takes label and field selectors, and returns the list of SharedMemoryStores that match those selectors.
func (c *FakeSharedMemoryStores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.SharedMemoryStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sharedmemorystoresResource, sharedmemorystoresKind, c.ns, opts), &v1beta2.SharedMemoryStoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.SharedMemoryStoreList{ListMeta: obj.(*v1beta2.SharedMemoryStoreList).ListMeta}
	for _, item := range obj.(*v1beta2.SharedMemoryStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sharedMemoryStores.
func (c *FakeSharedMemoryStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sharedmemorystoresResource, c.ns, opts))

}

// Create takes the representation of a sharedMemoryStore and creates it.  Returns the server's representation of the sharedMemoryStore, and an error, if there is any.
func (c *FakeSharedMemoryStores) Create(ctx context.Context, sharedMemoryStore *v1beta2.SharedMemoryStore, opts v1.CreateOptions) (result *v1beta2.SharedMemoryStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sharedmemorystoresResource, c.ns, sharedMemoryStore), &v1beta2.SharedMemoryStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SharedMemoryStore), err
}

// Update takes the representation of a sharedMemoryStore and updates it. Returns the server's representation of the sharedMemoryStore, and an error, if there is any.
func (c *FakeSharedMemoryStores) Update(ctx context.Context, sharedMemoryStore *v1beta2.SharedMemoryStore, opts v1.UpdateOptions) (result *v1beta2.SharedMemoryStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sharedmemorystoresResource, c.ns, sharedMemoryStore), &v1beta2.SharedMemoryStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SharedMemoryStore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSharedMemoryStores) UpdateStatus(ctx context.Context, sharedMemoryStore *v1beta2.SharedMemoryStore, opts v1.UpdateOptions) (*v1beta2.SharedMemoryStore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sharedmemorystoresResource, "status", c.ns, sharedMemoryStore), &v1beta2.SharedMemoryStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SharedMemoryStore), err
}

// Delete takes name of the sharedMemoryStore and deletes it. Returns an error if one occurs.
func (c *FakeSharedMemoryStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sharedmemorystoresResource, c.ns, name, opts), &v1beta2.SharedMemoryStore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSharedMemoryStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sharedmemorystoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.SharedMemoryStoreList{})
	return err
}

// Patch applies the patch and returns the patched sharedMemoryStore.
func (c *FakeSharedMemoryStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.SharedMemoryStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sharedmemorystoresResource, c.ns, name, pt, data, subresources...), &v1beta2.SharedMemoryStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SharedMemoryStore), err
}
