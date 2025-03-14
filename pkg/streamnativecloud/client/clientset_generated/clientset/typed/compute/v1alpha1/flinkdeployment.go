// Copyright 2025 StreamNative
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/streamnative/pulsar-resources-operator/pkg/streamnativecloud/apis/compute/v1alpha1"
	scheme "github.com/streamnative/pulsar-resources-operator/pkg/streamnativecloud/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FlinkDeploymentsGetter has a method to return a FlinkDeploymentInterface.
// A group's client should implement this interface.
type FlinkDeploymentsGetter interface {
	FlinkDeployments(namespace string) FlinkDeploymentInterface
}

// FlinkDeploymentInterface has methods to work with FlinkDeployment resources.
type FlinkDeploymentInterface interface {
	Create(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.CreateOptions) (*v1alpha1.FlinkDeployment, error)
	Update(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.UpdateOptions) (*v1alpha1.FlinkDeployment, error)
	UpdateStatus(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.UpdateOptions) (*v1alpha1.FlinkDeployment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FlinkDeployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FlinkDeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FlinkDeployment, err error)
	FlinkDeploymentExpansion
}

// flinkDeployments implements FlinkDeploymentInterface
type flinkDeployments struct {
	client rest.Interface
	ns     string
}

// newFlinkDeployments returns a FlinkDeployments
func newFlinkDeployments(c *ComputeV1alpha1Client, namespace string) *flinkDeployments {
	return &flinkDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flinkDeployment, and returns the corresponding flinkDeployment object, and an error if there is any.
func (c *flinkDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FlinkDeployment, err error) {
	result = &v1alpha1.FlinkDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkdeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlinkDeployments that match those selectors.
func (c *flinkDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FlinkDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FlinkDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flinkDeployments.
func (c *flinkDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flinkdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a flinkDeployment and creates it.  Returns the server's representation of the flinkDeployment, and an error, if there is any.
func (c *flinkDeployments) Create(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.CreateOptions) (result *v1alpha1.FlinkDeployment, err error) {
	result = &v1alpha1.FlinkDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flinkdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flinkDeployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a flinkDeployment and updates it. Returns the server's representation of the flinkDeployment, and an error, if there is any.
func (c *flinkDeployments) Update(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.UpdateOptions) (result *v1alpha1.FlinkDeployment, err error) {
	result = &v1alpha1.FlinkDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flinkdeployments").
		Name(flinkDeployment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flinkDeployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *flinkDeployments) UpdateStatus(ctx context.Context, flinkDeployment *v1alpha1.FlinkDeployment, opts v1.UpdateOptions) (result *v1alpha1.FlinkDeployment, err error) {
	result = &v1alpha1.FlinkDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flinkdeployments").
		Name(flinkDeployment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flinkDeployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the flinkDeployment and deletes it. Returns an error if one occurs.
func (c *flinkDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkdeployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flinkDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkdeployments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched flinkDeployment.
func (c *flinkDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FlinkDeployment, err error) {
	result = &v1alpha1.FlinkDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flinkdeployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
