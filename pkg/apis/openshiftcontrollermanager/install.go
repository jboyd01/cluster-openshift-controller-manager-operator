package openshiftcontrollermanager

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/openshift/cluster-openshift-controller-manager-operator/pkg/apis/openshiftcontrollermanager/v1"
)

const (
	GroupName = "openshiftcontrollermanager.operator.openshift.io"
)

var (
	schemeBuilder = runtime.NewSchemeBuilder(v1.Install)
	// Install is a function which adds every version of this group to a scheme
	Install = schemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return schema.GroupResource{Group: GroupName, Resource: resource}
}

func Kind(kind string) schema.GroupKind {
	return schema.GroupKind{Group: GroupName, Kind: kind}
}
