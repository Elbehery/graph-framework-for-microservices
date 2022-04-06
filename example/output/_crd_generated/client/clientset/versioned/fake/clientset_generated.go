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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned"
	configtsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/config.tsm.tanzu.vmware.com/v1"
	fakeconfigtsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/config.tsm.tanzu.vmware.com/v1/fake"
	gnstsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/gns.tsm.tanzu.vmware.com/v1"
	fakegnstsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/gns.tsm.tanzu.vmware.com/v1/fake"
	nexustsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/nexus.tsm.tanzu.vmware.com/v1"
	fakenexustsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/nexus.tsm.tanzu.vmware.com/v1/fake"
	policytsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/policy.tsm.tanzu.vmware.com/v1"
	fakepolicytsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/policy.tsm.tanzu.vmware.com/v1/fake"
	roottsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/root.tsm.tanzu.vmware.com/v1"
	fakeroottsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/root.tsm.tanzu.vmware.com/v1/fake"
	servicegrouptsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/servicegroup.tsm.tanzu.vmware.com/v1"
	fakeservicegrouptsmv1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/client/clientset/versioned/typed/servicegroup.tsm.tanzu.vmware.com/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ConfigTsmV1 retrieves the ConfigTsmV1Client
func (c *Clientset) ConfigTsmV1() configtsmv1.ConfigTsmV1Interface {
	return &fakeconfigtsmv1.FakeConfigTsmV1{Fake: &c.Fake}
}

// GnsTsmV1 retrieves the GnsTsmV1Client
func (c *Clientset) GnsTsmV1() gnstsmv1.GnsTsmV1Interface {
	return &fakegnstsmv1.FakeGnsTsmV1{Fake: &c.Fake}
}

// NexusTsmV1 retrieves the NexusTsmV1Client
func (c *Clientset) NexusTsmV1() nexustsmv1.NexusTsmV1Interface {
	return &fakenexustsmv1.FakeNexusTsmV1{Fake: &c.Fake}
}

// PolicyTsmV1 retrieves the PolicyTsmV1Client
func (c *Clientset) PolicyTsmV1() policytsmv1.PolicyTsmV1Interface {
	return &fakepolicytsmv1.FakePolicyTsmV1{Fake: &c.Fake}
}

// RootTsmV1 retrieves the RootTsmV1Client
func (c *Clientset) RootTsmV1() roottsmv1.RootTsmV1Interface {
	return &fakeroottsmv1.FakeRootTsmV1{Fake: &c.Fake}
}

// ServicegroupTsmV1 retrieves the ServicegroupTsmV1Client
func (c *Clientset) ServicegroupTsmV1() servicegrouptsmv1.ServicegroupTsmV1Interface {
	return &fakeservicegrouptsmv1.FakeServicegroupTsmV1{Fake: &c.Fake}
}
