package fabric

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// StoragePoolsClient is the fabric Admin Client
type StoragePoolsClient struct {
	BaseClient
}

// NewStoragePoolsClient creates an instance of the StoragePoolsClient client.
func NewStoragePoolsClient(subscriptionID string) StoragePoolsClient {
	return NewStoragePoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStoragePoolsClientWithBaseURI creates an instance of the StoragePoolsClient client.
func NewStoragePoolsClientWithBaseURI(baseURI string, subscriptionID string) StoragePoolsClient {
	return StoragePoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a storage pool.
//
// location is location of the resource. storageSubSystem is name of the storage system. storagePool is storage
// pool name.
func (client StoragePoolsClient) Get(ctx context.Context, location string, storageSubSystem string, storagePool string) (result StoragePool, err error) {
	req, err := client.GetPreparer(ctx, location, storageSubSystem, storagePool)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StoragePoolsClient) GetPreparer(ctx context.Context, location string, storageSubSystem string, storagePool string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":         autorest.Encode("path", location),
		"storagePool":      autorest.Encode("path", storagePool),
		"storageSubSystem": autorest.Encode("path", storageSubSystem),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/storageSubSystems/{storageSubSystem}/storagePools/{storagePool}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StoragePoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StoragePoolsClient) GetResponder(resp *http.Response) (result StoragePool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get a list of all storage pools for a location.
//
// location is location of the resource. storageSubSystem is name of the storage system. filter is oData filter
// parameter.
func (client StoragePoolsClient) List(ctx context.Context, location string, storageSubSystem string, filter string) (result StoragePoolListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location, storageSubSystem, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.spl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.spl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client StoragePoolsClient) ListPreparer(ctx context.Context, location string, storageSubSystem string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":         autorest.Encode("path", location),
		"storageSubSystem": autorest.Encode("path", storageSubSystem),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/System.{location}/providers/Microsoft.Fabric.Admin/fabricLocations/{location}/storageSubSystems/{storageSubSystem}/storagePools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StoragePoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StoragePoolsClient) ListResponder(resp *http.Response) (result StoragePoolList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StoragePoolsClient) listNextResults(lastResults StoragePoolList) (result StoragePoolList, err error) {
	req, err := lastResults.storagePoolListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "fabric.StoragePoolsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StoragePoolsClient) ListComplete(ctx context.Context, location string, storageSubSystem string, filter string) (result StoragePoolListIterator, err error) {
	result.page, err = client.List(ctx, location, storageSubSystem, filter)
	return
}