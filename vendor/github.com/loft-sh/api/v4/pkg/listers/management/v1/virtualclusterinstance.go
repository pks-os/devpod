// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// VirtualClusterInstanceLister helps list VirtualClusterInstances.
// All objects returned here must be treated as read-only.
type VirtualClusterInstanceLister interface {
	// List lists all VirtualClusterInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error)
	// VirtualClusterInstances returns an object that can list and get VirtualClusterInstances.
	VirtualClusterInstances(namespace string) VirtualClusterInstanceNamespaceLister
	VirtualClusterInstanceListerExpansion
}

// virtualClusterInstanceLister implements the VirtualClusterInstanceLister interface.
type virtualClusterInstanceLister struct {
	listers.ResourceIndexer[*v1.VirtualClusterInstance]
}

// NewVirtualClusterInstanceLister returns a new VirtualClusterInstanceLister.
func NewVirtualClusterInstanceLister(indexer cache.Indexer) VirtualClusterInstanceLister {
	return &virtualClusterInstanceLister{listers.New[*v1.VirtualClusterInstance](indexer, v1.Resource("virtualclusterinstance"))}
}

// VirtualClusterInstances returns an object that can list and get VirtualClusterInstances.
func (s *virtualClusterInstanceLister) VirtualClusterInstances(namespace string) VirtualClusterInstanceNamespaceLister {
	return virtualClusterInstanceNamespaceLister{listers.NewNamespaced[*v1.VirtualClusterInstance](s.ResourceIndexer, namespace)}
}

// VirtualClusterInstanceNamespaceLister helps list and get VirtualClusterInstances.
// All objects returned here must be treated as read-only.
type VirtualClusterInstanceNamespaceLister interface {
	// List lists all VirtualClusterInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualClusterInstance, err error)
	// Get retrieves the VirtualClusterInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualClusterInstance, error)
	VirtualClusterInstanceNamespaceListerExpansion
}

// virtualClusterInstanceNamespaceLister implements the VirtualClusterInstanceNamespaceLister
// interface.
type virtualClusterInstanceNamespaceLister struct {
	listers.ResourceIndexer[*v1.VirtualClusterInstance]
}
