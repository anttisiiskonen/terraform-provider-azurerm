package synapse

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// IntegrationRuntimeObjectMetadataClient is the azure Synapse Analytics Management Client
type IntegrationRuntimeObjectMetadataClient struct {
	BaseClient
}

// NewIntegrationRuntimeObjectMetadataClient creates an instance of the IntegrationRuntimeObjectMetadataClient client.
func NewIntegrationRuntimeObjectMetadataClient(subscriptionID string) IntegrationRuntimeObjectMetadataClient {
	return NewIntegrationRuntimeObjectMetadataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeObjectMetadataClientWithBaseURI creates an instance of the
// IntegrationRuntimeObjectMetadataClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIntegrationRuntimeObjectMetadataClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeObjectMetadataClient {
	return IntegrationRuntimeObjectMetadataClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get object metadata from an integration runtime
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// integrationRuntimeName - integration runtime name
// getMetadataRequest - the parameters for getting a SSIS object metadata.
func (client IntegrationRuntimeObjectMetadataClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, getMetadataRequest *GetSsisObjectMetadataRequest) (result SsisObjectMetadataListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeObjectMetadataClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.IntegrationRuntimeObjectMetadataClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName, getMetadataRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationRuntimeObjectMetadataClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, getMetadataRequest *GetSsisObjectMetadataRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/getObjectMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if getMetadataRequest != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(getMetadataRequest))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeObjectMetadataClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeObjectMetadataClient) GetResponder(resp *http.Response) (result SsisObjectMetadataListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Refresh refresh the object metadata in an integration runtime
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// integrationRuntimeName - integration runtime name
func (client IntegrationRuntimeObjectMetadataClient) Refresh(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result SsisObjectMetadataStatusResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeObjectMetadataClient.Refresh")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.IntegrationRuntimeObjectMetadataClient", "Refresh", err.Error())
	}

	req, err := client.RefreshPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Refresh", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Refresh", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeObjectMetadataClient", "Refresh", resp, "Failure responding to request")
	}

	return
}

// RefreshPreparer prepares the Refresh request.
func (client IntegrationRuntimeObjectMetadataClient) RefreshPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/refreshObjectMetadata", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshSender sends the Refresh request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeObjectMetadataClient) RefreshSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RefreshResponder handles the response to the Refresh request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeObjectMetadataClient) RefreshResponder(resp *http.Response) (result SsisObjectMetadataStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
