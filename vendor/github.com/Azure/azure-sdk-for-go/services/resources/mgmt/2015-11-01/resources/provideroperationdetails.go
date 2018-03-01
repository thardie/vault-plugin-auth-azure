package resources

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

// ProviderOperationDetailsClient is the client for the ProviderOperationDetails methods of the Resources service.
type ProviderOperationDetailsClient struct {
	BaseClient
}

// NewProviderOperationDetailsClient creates an instance of the ProviderOperationDetailsClient client.
func NewProviderOperationDetailsClient(subscriptionID string) ProviderOperationDetailsClient {
	return NewProviderOperationDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderOperationDetailsClientWithBaseURI creates an instance of the ProviderOperationDetailsClient client.
func NewProviderOperationDetailsClientWithBaseURI(baseURI string, subscriptionID string) ProviderOperationDetailsClient {
	return ProviderOperationDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of resource providers.
//
// resourceProviderNamespace is resource identity.
func (client ProviderOperationDetailsClient) List(ctx context.Context, resourceProviderNamespace string) (result ProviderOperationDetailListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceProviderNamespace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.podlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "List", resp, "Failure sending request")
		return
	}

	result.podlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProviderOperationDetailsClient) ListPreparer(ctx context.Context, resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
	}

	const APIVersion = "2015-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{resourceProviderNamespace}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderOperationDetailsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProviderOperationDetailsClient) ListResponder(resp *http.Response) (result ProviderOperationDetailListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProviderOperationDetailsClient) listNextResults(lastResults ProviderOperationDetailListResult) (result ProviderOperationDetailListResult, err error) {
	req, err := lastResults.providerOperationDetailListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ProviderOperationDetailsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderOperationDetailsClient) ListComplete(ctx context.Context, resourceProviderNamespace string) (result ProviderOperationDetailListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceProviderNamespace)
	return
}
