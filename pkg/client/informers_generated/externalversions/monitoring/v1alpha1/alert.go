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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	monitoring_v1alpha1 "github.com/prohori/prohori/pkg/apis/monitoring/v1alpha1"
	clientset "github.com/prohori/prohori/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/prohori/prohori/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/prohori/prohori/pkg/client/listers_generated/monitoring/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// AlertInformer provides access to a shared informer and lister for
// Alerts.
type AlertInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AlertLister
}

type alertInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAlertInformer constructs a new informer for Alert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAlertInformer constructs a new informer for Alert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().Alerts(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().Alerts(namespace).Watch(options)
			},
		},
		&monitoring_v1alpha1.Alert{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoring_v1alpha1.Alert{}, f.defaultInformer)
}

func (f *alertInformer) Lister() v1alpha1.AlertLister {
	return v1alpha1.NewAlertLister(f.Informer().GetIndexer())
}
