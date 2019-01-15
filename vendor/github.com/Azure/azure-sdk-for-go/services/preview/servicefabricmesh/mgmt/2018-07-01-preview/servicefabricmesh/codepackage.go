package servicefabricmesh

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

// CodePackageClient is the service Fabric Mesh Management Client
type CodePackageClient struct {
	BaseClient
}

// NewCodePackageClient creates an instance of the CodePackageClient client.
func NewCodePackageClient(subscriptionID string) CodePackageClient {
	return NewCodePackageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCodePackageClientWithBaseURI creates an instance of the CodePackageClient client.
func NewCodePackageClientWithBaseURI(baseURI string, subscriptionID string) CodePackageClient {
	return CodePackageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetContainerLog get the logs for the container of a given code package of an application.
// Parameters:
// resourceGroupName - azure resource group name
// applicationName - the identity of the application.
// serviceName - the identity of the service.
// replicaName - the identity of the service replica.
// codePackageName - the name of the code package.
// tail - number of lines to show from the end of the logs. Default is 100.
func (client CodePackageClient) GetContainerLog(ctx context.Context, resourceGroupName string, applicationName string, serviceName string, replicaName string, codePackageName string, tail *int32) (result ContainerLogs, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CodePackageClient.GetContainerLog")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetContainerLogPreparer(ctx, resourceGroupName, applicationName, serviceName, replicaName, codePackageName, tail)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.CodePackageClient", "GetContainerLog", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContainerLogSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabricmesh.CodePackageClient", "GetContainerLog", resp, "Failure sending request")
		return
	}

	result, err = client.GetContainerLogResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabricmesh.CodePackageClient", "GetContainerLog", resp, "Failure responding to request")
	}

	return
}

// GetContainerLogPreparer prepares the GetContainerLog request.
func (client CodePackageClient) GetContainerLogPreparer(ctx context.Context, resourceGroupName string, applicationName string, serviceName string, replicaName string, codePackageName string, tail *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   applicationName,
		"codePackageName":   codePackageName,
		"replicaName":       replicaName,
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       serviceName,
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if tail != nil {
		queryParameters["tail"] = autorest.Encode("query", *tail)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationName}/services/{serviceName}/replicas/{replicaName}/codePackages/{codePackageName}/logs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContainerLogSender sends the GetContainerLog request. The method will close the
// http.Response Body if it receives an error.
func (client CodePackageClient) GetContainerLogSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetContainerLogResponder handles the response to the GetContainerLog request. The method always
// closes the http.Response Body.
func (client CodePackageClient) GetContainerLogResponder(resp *http.Response) (result ContainerLogs, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}