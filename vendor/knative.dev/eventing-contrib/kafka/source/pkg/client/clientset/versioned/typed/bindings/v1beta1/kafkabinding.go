/*
Copyright 2020 The Knative Authors

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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/bindings/v1beta1"
	scheme "knative.dev/eventing-contrib/kafka/source/pkg/client/clientset/versioned/scheme"
)

// KafkaBindingsGetter has a method to return a KafkaBindingInterface.
// A group's client should implement this interface.
type KafkaBindingsGetter interface {
	KafkaBindings(namespace string) KafkaBindingInterface
}

// KafkaBindingInterface has methods to work with KafkaBinding resources.
type KafkaBindingInterface interface {
	Create(*v1beta1.KafkaBinding) (*v1beta1.KafkaBinding, error)
	Update(*v1beta1.KafkaBinding) (*v1beta1.KafkaBinding, error)
	UpdateStatus(*v1beta1.KafkaBinding) (*v1beta1.KafkaBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.KafkaBinding, error)
	List(opts v1.ListOptions) (*v1beta1.KafkaBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KafkaBinding, err error)
	KafkaBindingExpansion
}

// kafkaBindings implements KafkaBindingInterface
type kafkaBindings struct {
	client rest.Interface
	ns     string
}

// newKafkaBindings returns a KafkaBindings
func newKafkaBindings(c *BindingsV1beta1Client, namespace string) *kafkaBindings {
	return &kafkaBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaBinding, and returns the corresponding kafkaBinding object, and an error if there is any.
func (c *kafkaBindings) Get(name string, options v1.GetOptions) (result *v1beta1.KafkaBinding, err error) {
	result = &v1beta1.KafkaBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkabindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaBindings that match those selectors.
func (c *kafkaBindings) List(opts v1.ListOptions) (result *v1beta1.KafkaBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KafkaBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkabindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaBindings.
func (c *kafkaBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkabindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kafkaBinding and creates it.  Returns the server's representation of the kafkaBinding, and an error, if there is any.
func (c *kafkaBindings) Create(kafkaBinding *v1beta1.KafkaBinding) (result *v1beta1.KafkaBinding, err error) {
	result = &v1beta1.KafkaBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kafkabindings").
		Body(kafkaBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kafkaBinding and updates it. Returns the server's representation of the kafkaBinding, and an error, if there is any.
func (c *kafkaBindings) Update(kafkaBinding *v1beta1.KafkaBinding) (result *v1beta1.KafkaBinding, err error) {
	result = &v1beta1.KafkaBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkabindings").
		Name(kafkaBinding.Name).
		Body(kafkaBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kafkaBindings) UpdateStatus(kafkaBinding *v1beta1.KafkaBinding) (result *v1beta1.KafkaBinding, err error) {
	result = &v1beta1.KafkaBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kafkabindings").
		Name(kafkaBinding.Name).
		SubResource("status").
		Body(kafkaBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the kafkaBinding and deletes it. Returns an error if one occurs.
func (c *kafkaBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkabindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kafkaBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kafkabindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kafkaBinding.
func (c *kafkaBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.KafkaBinding, err error) {
	result = &v1beta1.KafkaBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kafkabindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
