// Copyright 2017-2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	time "time"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumNetworkPolicyInformer provides access to a shared informer and lister for
// CiliumNetworkPolicies.
type CiliumNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumNetworkPolicyLister
}

type ciliumNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCiliumNetworkPolicyInformer constructs a new informer for CiliumNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumNetworkPolicyInformer constructs a new informer for CiliumNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumNetworkPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumNetworkPolicies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumNetworkPolicies(namespace).Watch(options)
			},
		},
		&ciliumiov2.CiliumNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumNetworkPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumNetworkPolicy{}, f.defaultInformer)
}

func (f *ciliumNetworkPolicyInformer) Lister() v2.CiliumNetworkPolicyLister {
	return v2.NewCiliumNetworkPolicyLister(f.Informer().GetIndexer())
}
