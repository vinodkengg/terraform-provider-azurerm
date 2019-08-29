package dtl

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ArtifactsClient is the the DevTest Labs Client.
type ArtifactsClient struct {
	BaseClient
}

// NewArtifactsClient creates an instance of the ArtifactsClient client.
func NewArtifactsClient(subscriptionID string) ArtifactsClient {
	return NewArtifactsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewArtifactsClientWithBaseURI creates an instance of the ArtifactsClient client.
func NewArtifactsClientWithBaseURI(baseURI string, subscriptionID string) ArtifactsClient {
	return ArtifactsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GenerateArmTemplate generates an ARM template for the given artifact, uploads the required files to a storage
// account, and validates the generated artifact.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// name - the name of the artifact.
// generateArmTemplateRequest - parameters for generating an ARM template for deploying artifacts.
func (client ArtifactsClient) GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest) (result ArmTemplateInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactsClient.GenerateArmTemplate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GenerateArmTemplatePreparer(ctx, resourceGroupName, labName, artifactSourceName, name, generateArmTemplateRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "GenerateArmTemplate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateArmTemplateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "GenerateArmTemplate", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateArmTemplateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "GenerateArmTemplate", resp, "Failure responding to request")
	}

	return
}

// GenerateArmTemplatePreparer prepares the GenerateArmTemplate request.
func (client ArtifactsClient) GenerateArmTemplatePreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest GenerateArmTemplateRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}/generateArmTemplate", pathParameters),
		autorest.WithJSON(generateArmTemplateRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateArmTemplateSender sends the GenerateArmTemplate request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactsClient) GenerateArmTemplateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GenerateArmTemplateResponder handles the response to the GenerateArmTemplate request. The method always
// closes the http.Response Body.
func (client ArtifactsClient) GenerateArmTemplateResponder(resp *http.Response) (result ArmTemplateInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get artifact.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// name - the name of the artifact.
// expand - specify the $expand query. Example: 'properties($select=title)'
func (client ArtifactsClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (result Artifact, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labName, artifactSourceName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ArtifactsClient) GetPreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"name":               autorest.Encode("path", name),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ArtifactsClient) GetResponder(resp *http.Response) (result Artifact, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list artifacts in a given artifact source.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// artifactSourceName - the name of the artifact source.
// expand - specify the $expand query. Example: 'properties($select=title)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client ArtifactsClient) List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationArtifactPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactsClient.List")
		defer func() {
			sc := -1
			if result.rwca.Response.Response != nil {
				sc = result.rwca.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwca.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "List", resp, "Failure sending request")
		return
	}

	result.rwca, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ArtifactsClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"artifactSourceName": autorest.Encode("path", artifactSourceName),
		"labName":            autorest.Encode("path", labName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ArtifactsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ArtifactsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationArtifact, err error) {
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
func (client ArtifactsClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationArtifact) (result ResponseWithContinuationArtifact, err error) {
	req, err := lastResults.responseWithContinuationArtifactPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.ArtifactsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ArtifactsClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationArtifactIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ArtifactsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labName, artifactSourceName, expand, filter, top, orderby)
	return
}
