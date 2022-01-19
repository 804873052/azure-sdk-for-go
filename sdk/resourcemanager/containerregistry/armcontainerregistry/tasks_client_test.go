//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcontainerregistry_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTasksClient_BeginCreate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// create task
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	createTask(t, ctx, tasksClient, rgName, registryName, location)
}

func TestTasksClient_BeginDelete(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// create task
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	task := createTask(t, ctx, tasksClient, rgName, registryName, location)
	taskName := *task.Name

	// delete task
	delPollerResp, err := tasksClient.BeginDelete(ctx, rgName, registryName, taskName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestTasksClient_BeginUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// create task
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	task := createTask(t, ctx, tasksClient, rgName, registryName, location)
	taskName := *task.Name

	// update task
	updatePollerResp, err := tasksClient.BeginUpdate(
		ctx,
		rgName,
		registryName,
		taskName,
		armcontainerregistry.TaskUpdateParameters{
			Tags: map[string]*string{
				"test": to.StringPtr("recording"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	updateResp, err := updatePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.Equal(t, "recording", updateResp.Tags["test"])
}

func TestTasksClient_Get(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// create task
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	task := createTask(t, ctx, tasksClient, rgName, registryName, location)
	taskName := *task.Name

	// get task
	getResp, err := tasksClient.Get(ctx, rgName, registryName, taskName, nil)
	require.NoError(t, err)
	require.Equal(t, taskName, *getResp.Name)
}

func TestTasksClient_GetDetails(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// create task
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	task := createTask(t, ctx, tasksClient, rgName, registryName, location)
	taskName := *task.Name

	// get task detail
	getResp, err := tasksClient.GetDetails(ctx, rgName, registryName, taskName, nil)
	require.NoError(t, err)
	require.Equal(t, taskName, *getResp.Name)
}

func TestTasksClient_List(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// registry container
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, nil)
	registry := createContainerRegistry(t, ctx, registriesClient, rgName, location)
	registryName := *registry.Name

	// list tasks
	tasksClient := armcontainerregistry.NewTasksClient(subscriptionID, cred, opt)
	pager := tasksClient.List(rgName, registryName, nil)
	require.NoError(t, pager.Err())
}

func createTask(t *testing.T, ctx context.Context, tasksClient *armcontainerregistry.TasksClient, rgName, registryName, location string) *armcontainerregistry.Task {
	taskName, _ := createRandomName(t, "task")
	taskPollerResp, err := tasksClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		taskName,
		armcontainerregistry.Task{
			Location: to.StringPtr(location),
			Properties: &armcontainerregistry.TaskProperties{
				Status: armcontainerregistry.TaskStatusEnabled.ToPtr(),
				Platform: &armcontainerregistry.PlatformProperties{
					OS:           armcontainerregistry.OSLinux.ToPtr(),
					Architecture: armcontainerregistry.ArchitectureAmd64.ToPtr(),
				},
				AgentConfiguration: &armcontainerregistry.AgentProperties{
					CPU: to.Int32Ptr(2),
				},
				Step: &armcontainerregistry.DockerBuildStep{
					Type:        armcontainerregistry.StepTypeDocker.ToPtr(),
					ContextPath: to.StringPtr("https://github.com/SteveLasker/node-helloworld"),
					ImageNames: []*string{
						to.StringPtr("testtask:v1"),
					},
					DockerFilePath: to.StringPtr("Dockerfile"),
					IsPushEnabled:  to.BoolPtr(true),
					NoCache:        to.BoolPtr(false),
				},
				Trigger: &armcontainerregistry.TriggerProperties{
					BaseImageTrigger: &armcontainerregistry.BaseImageTrigger{
						Name:                     to.StringPtr("myBaseImageTrigger"),
						BaseImageTriggerType:     armcontainerregistry.BaseImageTriggerTypeRuntime.ToPtr(),
						UpdateTriggerPayloadType: armcontainerregistry.UpdateTriggerPayloadTypeDefault.ToPtr(),
						Status:                   armcontainerregistry.TriggerStatusEnabled.ToPtr(),
					},
				},
			},
		},
		nil,
	)
	require.NoError(t, err)
	taskResp, err := taskPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, taskName, *taskResp.Name)
	return &taskResp.Task
}
