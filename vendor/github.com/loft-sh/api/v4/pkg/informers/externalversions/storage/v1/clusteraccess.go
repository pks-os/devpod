// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	storagev1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	versioned "github.com/loft-sh/api/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/api/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/api/v4/pkg/listers/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterAccessInformer provides access to a shared informer and lister for
// ClusterAccesses.
type ClusterAccessInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterAccessLister
}

type clusterAccessInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterAccessInformer constructs a new informer for ClusterAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterAccessInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterAccessInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterAccessInformer constructs a new informer for ClusterAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterAccessInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().ClusterAccesses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().ClusterAccesses().Watch(context.TODO(), options)
			},
		},
		&storagev1.ClusterAccess{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterAccessInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterAccessInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterAccessInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1.ClusterAccess{}, f.defaultInformer)
}

func (f *clusterAccessInformer) Lister() v1.ClusterAccessLister {
	return v1.NewClusterAccessLister(f.Informer().GetIndexer())
}
