// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubecontrollermanager_v1alpha1 "github.com/openshift/cluster-kube-controller-manager-operator/pkg/apis/kubecontrollermanager/v1alpha1"
	versioned "github.com/openshift/cluster-kube-controller-manager-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-kube-controller-manager-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/cluster-kube-controller-manager-operator/pkg/generated/listers/kubecontrollermanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeApiserverOperatorConfigInformer provides access to a shared informer and lister for
// KubeApiserverOperatorConfigs.
type KubeApiserverOperatorConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubeApiserverOperatorConfigLister
}

type kubeApiserverOperatorConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeApiserverOperatorConfigInformer constructs a new informer for KubeApiserverOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeApiserverOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeApiserverOperatorConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeApiserverOperatorConfigInformer constructs a new informer for KubeApiserverOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeApiserverOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubecontrollermanagerV1alpha1().KubeApiserverOperatorConfigs().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubecontrollermanagerV1alpha1().KubeApiserverOperatorConfigs().Watch(options)
			},
		},
		&kubecontrollermanager_v1alpha1.KubeApiserverOperatorConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeApiserverOperatorConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeApiserverOperatorConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeApiserverOperatorConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubecontrollermanager_v1alpha1.KubeApiserverOperatorConfig{}, f.defaultInformer)
}

func (f *kubeApiserverOperatorConfigInformer) Lister() v1alpha1.KubeApiserverOperatorConfigLister {
	return v1alpha1.NewKubeApiserverOperatorConfigLister(f.Informer().GetIndexer())
}
