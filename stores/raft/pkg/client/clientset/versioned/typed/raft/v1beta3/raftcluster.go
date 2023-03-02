// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package v1beta3

import (
	"context"
	"time"

	v1beta3 "github.com/atomix/atomix/stores/raft/pkg/apis/raft/v1beta3"
	scheme "github.com/atomix/atomix/stores/raft/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RaftClustersGetter has a method to return a RaftClusterInterface.
// A group's client should implement this interface.
type RaftClustersGetter interface {
	RaftClusters(namespace string) RaftClusterInterface
}

// RaftClusterInterface has methods to work with RaftCluster resources.
type RaftClusterInterface interface {
	Create(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.CreateOptions) (*v1beta3.RaftCluster, error)
	Update(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.UpdateOptions) (*v1beta3.RaftCluster, error)
	UpdateStatus(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.UpdateOptions) (*v1beta3.RaftCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.RaftCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.RaftClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.RaftCluster, err error)
	RaftClusterExpansion
}

// raftClusters implements RaftClusterInterface
type raftClusters struct {
	client rest.Interface
	ns     string
}

// newRaftClusters returns a RaftClusters
func newRaftClusters(c *RaftV1beta3Client, namespace string) *raftClusters {
	return &raftClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the raftCluster, and returns the corresponding raftCluster object, and an error if there is any.
func (c *raftClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta3.RaftCluster, err error) {
	result = &v1beta3.RaftCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("raftclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RaftClusters that match those selectors.
func (c *raftClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta3.RaftClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta3.RaftClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("raftclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested raftClusters.
func (c *raftClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("raftclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a raftCluster and creates it.  Returns the server's representation of the raftCluster, and an error, if there is any.
func (c *raftClusters) Create(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.CreateOptions) (result *v1beta3.RaftCluster, err error) {
	result = &v1beta3.RaftCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("raftclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a raftCluster and updates it. Returns the server's representation of the raftCluster, and an error, if there is any.
func (c *raftClusters) Update(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.UpdateOptions) (result *v1beta3.RaftCluster, err error) {
	result = &v1beta3.RaftCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("raftclusters").
		Name(raftCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *raftClusters) UpdateStatus(ctx context.Context, raftCluster *v1beta3.RaftCluster, opts v1.UpdateOptions) (result *v1beta3.RaftCluster, err error) {
	result = &v1beta3.RaftCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("raftclusters").
		Name(raftCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(raftCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the raftCluster and deletes it. Returns an error if one occurs.
func (c *raftClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("raftclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *raftClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("raftclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched raftCluster.
func (c *raftClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.RaftCluster, err error) {
	result = &v1beta3.RaftCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("raftclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
