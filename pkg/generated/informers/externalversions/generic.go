/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta1 "github.com/rabbitmq/messaging-topology-operator/api/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=rabbitmq.com, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("bindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Bindings().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("exchanges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Exchanges().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("federations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Federations().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("permissions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Permissions().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("policies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Policies().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("queues"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Queues().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("schemareplications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().SchemaReplications().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("shovels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Shovels().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("superstreams"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().SuperStreams().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("superstreamconsumers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().SuperStreamConsumers().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Users().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("vhosts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rabbitmq().V1beta1().Vhosts().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
