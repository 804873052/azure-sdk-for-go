//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_ListByFactory.json
func ExamplePipelinesClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByFactoryPager("exampleResourceGroup",
		"exampleFactoryName",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
func ExamplePipelinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"examplePipeline",
		armdatafactory.PipelineResource{
			Properties: &armdatafactory.Pipeline{
				Activities: []armdatafactory.ActivityClassification{
					&armdatafactory.ForEachActivity{
						Name: to.Ptr("ExampleForeachActivity"),
						Type: to.Ptr("ForEach"),
						TypeProperties: &armdatafactory.ForEachActivityTypeProperties{
							Activities: []armdatafactory.ActivityClassification{
								&armdatafactory.CopyActivity{
									Name: to.Ptr("ExampleCopyActivity"),
									Type: to.Ptr("Copy"),
									Inputs: []*armdatafactory.DatasetReference{
										{
											Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
											Parameters: map[string]interface{}{
												"MyFileName":   "examplecontainer.csv",
												"MyFolderPath": "examplecontainer",
											},
											ReferenceName: to.Ptr("exampleDataset"),
										}},
									Outputs: []*armdatafactory.DatasetReference{
										{
											Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
											Parameters: map[string]interface{}{
												"MyFileName": map[string]interface{}{
													"type":  "Expression",
													"value": "@item()",
												},
												"MyFolderPath": "examplecontainer",
											},
											ReferenceName: to.Ptr("exampleDataset"),
										}},
									TypeProperties: &armdatafactory.CopyActivityTypeProperties{
										DataIntegrationUnits: float64(32),
										Sink: &armdatafactory.BlobSink{
											Type: to.Ptr("BlobSink"),
										},
										Source: &armdatafactory.BlobSource{
											Type: to.Ptr("BlobSource"),
										},
									},
								}},
							IsSequential: to.Ptr(true),
							Items: &armdatafactory.Expression{
								Type:  to.Ptr(armdatafactory.ExpressionTypeExpression),
								Value: to.Ptr("@pipeline().parameters.OutputBlobNameList"),
							},
						},
					}},
				Parameters: map[string]*armdatafactory.ParameterSpecification{
					"JobId": {
						Type: to.Ptr(armdatafactory.ParameterTypeString),
					},
					"OutputBlobNameList": {
						Type: to.Ptr(armdatafactory.ParameterTypeArray),
					},
				},
				Policy: &armdatafactory.PipelinePolicy{
					ElapsedTimeMetric: &armdatafactory.PipelineElapsedTimeMetricPolicy{
						Duration: "0.00:10:00",
					},
				},
				RunDimensions: map[string]interface{}{
					"JobId": map[string]interface{}{
						"type":  "Expression",
						"value": "@pipeline().parameters.JobId",
					},
				},
				Variables: map[string]*armdatafactory.VariableSpecification{
					"TestVariableArray": {
						Type: to.Ptr(armdatafactory.VariableTypeArray),
					},
				},
			},
		},
		&armdatafactory.PipelinesClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Get.json
func ExamplePipelinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"examplePipeline",
		&armdatafactory.PipelinesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Delete.json
func ExamplePipelinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"examplePipeline",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_CreateRun.json
func ExamplePipelinesClient_CreateRun() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewPipelinesClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateRun(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"examplePipeline",
		&armdatafactory.PipelinesClientCreateRunOptions{ReferencePipelineRunID: nil,
			IsRecovery:        nil,
			StartActivityName: nil,
			StartFromFailure:  nil,
			Parameters: map[string]interface{}{
				"OutputBlobNameList": []interface{}{
					"exampleoutput.csv",
				},
			},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
