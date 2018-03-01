package sql

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

// ServerDNSAliasesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ServerDNSAliasesClient struct {
	BaseClient
}

// NewServerDNSAliasesClient creates an instance of the ServerDNSAliasesClient client.
func NewServerDNSAliasesClient(subscriptionID string) ServerDNSAliasesClient {
	return NewServerDNSAliasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerDNSAliasesClientWithBaseURI creates an instance of the ServerDNSAliasesClient client.
func NewServerDNSAliasesClientWithBaseURI(baseURI string, subscriptionID string) ServerDNSAliasesClient {
	return ServerDNSAliasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Acquire acquires server DNS alias from another server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server that the alias is pointing
// to. DNSAliasName is the name of the server dns alias.
func (client ServerDNSAliasesClient) Acquire(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string, parameters ServerDNSAliasAcquisition) (result ServerDNSAliasesAcquireFuture, err error) {
	req, err := client.AcquirePreparer(ctx, resourceGroupName, serverName, DNSAliasName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Acquire", nil, "Failure preparing request")
		return
	}

	result, err = client.AcquireSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Acquire", result.Response(), "Failure sending request")
		return
	}

	return
}

// AcquirePreparer prepares the Acquire request.
func (client ServerDNSAliasesClient) AcquirePreparer(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string, parameters ServerDNSAliasAcquisition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dnsAliasName":      autorest.Encode("path", DNSAliasName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}/acquire", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AcquireSender sends the Acquire request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDNSAliasesClient) AcquireSender(req *http.Request) (future ServerDNSAliasesAcquireFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// AcquireResponder handles the response to the Acquire request. The method always
// closes the http.Response Body.
func (client ServerDNSAliasesClient) AcquireResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate creates a server dns alias.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server that the alias is pointing
// to. DNSAliasName is the name of the server DNS alias.
func (client ServerDNSAliasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (result ServerDNSAliasesCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, DNSAliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServerDNSAliasesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dnsAliasName":      autorest.Encode("path", DNSAliasName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDNSAliasesClient) CreateOrUpdateSender(req *http.Request) (future ServerDNSAliasesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServerDNSAliasesClient) CreateOrUpdateResponder(resp *http.Response) (result ServerDNSAlias, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the server DNS alias with the given name.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server that the alias is pointing
// to. DNSAliasName is the name of the server DNS alias.
func (client ServerDNSAliasesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (result ServerDNSAliasesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, DNSAliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServerDNSAliasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dnsAliasName":      autorest.Encode("path", DNSAliasName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDNSAliasesClient) DeleteSender(req *http.Request) (future ServerDNSAliasesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServerDNSAliasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a server DNS alias.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server that the alias is pointing
// to. DNSAliasName is the name of the server DNS alias.
func (client ServerDNSAliasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (result ServerDNSAlias, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, DNSAliasName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerDNSAliasesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, DNSAliasName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dnsAliasName":      autorest.Encode("path", DNSAliasName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDNSAliasesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerDNSAliasesClient) GetResponder(resp *http.Response) (result ServerDNSAlias, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer gets a list of server DNS aliases for a server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server that the alias is pointing
// to.
func (client ServerDNSAliasesClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ServerDNSAliasListResultPage, err error) {
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.sdalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.sdalr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ServerDNSAliasesClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ServerDNSAliasesClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ServerDNSAliasesClient) ListByServerResponder(resp *http.Response) (result ServerDNSAliasListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client ServerDNSAliasesClient) listByServerNextResults(lastResults ServerDNSAliasListResult) (result ServerDNSAliasListResult, err error) {
	req, err := lastResults.serverDNSAliasListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerDNSAliasesClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServerDNSAliasesClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result ServerDNSAliasListResultIterator, err error) {
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}
