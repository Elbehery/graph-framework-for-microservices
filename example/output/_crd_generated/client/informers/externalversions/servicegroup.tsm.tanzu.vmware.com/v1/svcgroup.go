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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	servicegrouptsmtanzuvmwarecomv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/apis/servicegroup.tsm.tanzu.vmware.com/v1"
	versioned "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned"
	internalinterfaces "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/informers/externalversions/internalinterfaces"
	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/listers/servicegroup.tsm.tanzu.vmware.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SvcGroupInformer provides access to a shared informer and lister for
// SvcGroups.
type SvcGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SvcGroupLister
}

type svcGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSvcGroupInformer constructs a new informer for SvcGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSvcGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSvcGroupInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSvcGroupInformer constructs a new informer for SvcGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSvcGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicegroupTsmV1().SvcGroups().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicegroupTsmV1().SvcGroups().Watch(context.TODO(), options)
			},
		},
		&servicegrouptsmtanzuvmwarecomv1.SvcGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *svcGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSvcGroupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *svcGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicegrouptsmtanzuvmwarecomv1.SvcGroup{}, f.defaultInformer)
}

func (f *svcGroupInformer) Lister() v1.SvcGroupLister {
	return v1.NewSvcGroupLister(f.Informer().GetIndexer())
}
