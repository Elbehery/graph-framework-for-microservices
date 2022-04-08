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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "gitlab.eng.vmware.com/nsx-allspark_users/nexus-sdk/compiler.git/example/output/_crd_generated/apis/policy.tsm.tanzu.vmware.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AccessControlPolicyLister helps list AccessControlPolicies.
// All objects returned here must be treated as read-only.
type AccessControlPolicyLister interface {
	// List lists all AccessControlPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AccessControlPolicy, err error)
	// Get retrieves the AccessControlPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AccessControlPolicy, error)
	AccessControlPolicyListerExpansion
}

// accessControlPolicyLister implements the AccessControlPolicyLister interface.
type accessControlPolicyLister struct {
	indexer cache.Indexer
}

// NewAccessControlPolicyLister returns a new AccessControlPolicyLister.
func NewAccessControlPolicyLister(indexer cache.Indexer) AccessControlPolicyLister {
	return &accessControlPolicyLister{indexer: indexer}
}

// List lists all AccessControlPolicies in the indexer.
func (s *accessControlPolicyLister) List(selector labels.Selector) (ret []*v1.AccessControlPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AccessControlPolicy))
	})
	return ret, err
}

// Get retrieves the AccessControlPolicy from the index for a given name.
func (s *accessControlPolicyLister) Get(name string) (*v1.AccessControlPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("accesscontrolpolicy"), name)
	}
	return obj.(*v1.AccessControlPolicy), nil
}
