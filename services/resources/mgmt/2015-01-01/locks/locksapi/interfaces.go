package locksapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-01-01/locks"
	"github.com/Azure/go-autorest/autorest"
)

// ManagementLocksClientAPI contains the set of methods on the ManagementLocksClient type.
type ManagementLocksClientAPI interface {
	CreateOrUpdateAtResourceGroupLevel(ctx context.Context, resourceGroupName string, lockName string, parameters locks.ManagementLockObject) (result locks.ManagementLockObject, err error)
	CreateOrUpdateAtResourceLevel(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string, parameters locks.ManagementLockObject) (result locks.ManagementLockObject, err error)
	CreateOrUpdateAtSubscriptionLevel(ctx context.Context, lockName string, parameters locks.ManagementLockObject) (result locks.ManagementLockObject, err error)
	DeleteAtResourceGroupLevel(ctx context.Context, resourceGroupName string, lockName string) (result autorest.Response, err error)
	DeleteAtResourceLevel(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string) (result autorest.Response, err error)
	DeleteAtSubscriptionLevel(ctx context.Context, lockName string) (result autorest.Response, err error)
	Get(ctx context.Context, lockName string) (result locks.ManagementLockObject, err error)
	GetAtResourceGroupLevel(ctx context.Context, resourceGroupName string, lockName string) (result locks.ManagementLockObject, err error)
	ListAtResourceGroupLevel(ctx context.Context, resourceGroupName string, filter string) (result locks.ManagementLockListResultPage, err error)
	ListAtResourceLevel(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result locks.ManagementLockListResultPage, err error)
	ListAtSubscriptionLevel(ctx context.Context, filter string) (result locks.ManagementLockListResultPage, err error)
}

var _ ManagementLocksClientAPI = (*locks.ManagementLocksClient)(nil)
