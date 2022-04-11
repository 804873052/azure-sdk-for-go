//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesBySubscription.json
func ExampleClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.List(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesByResourceGroup.json
func ExampleClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateDomainService.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				DomainName: to.Ptr("<domain-name>"),
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
					SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
					TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
				},
				FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
					Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
					PfxCertificate:         to.Ptr("<pfx-certificate>"),
					PfxCertificatePassword: to.Ptr("<pfx-certificate-password>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.Ptr("jicha@microsoft.com"),
						to.Ptr("caalmont@microsoft.com")},
					NotifyDcAdmins:     to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
					NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.Ptr("<location>"),
						SubnetID: to.Ptr("<subnet-id>"),
					}},
			},
		},
		&armdomainservices.ClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetDomainService.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/DeleteDomainService.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		&armdomainservices.ClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdomainservices.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				ConfigDiagnostics: &armdomainservices.ConfigDiagnostics{
					LastExecuted: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2021-05-05T12:00:23Z;"); return t }()),
					ValidatorResults: []*armdomainservices.ConfigDiagnosticsValidatorResult{
						{
							Issues: []*armdomainservices.ConfigDiagnosticsValidatorResultIssue{
								{
									DescriptionParams: []*string{},
									ID:                to.Ptr("<id>"),
								}},
							ReplicaSetSubnetDisplayName: to.Ptr("<replica-set-subnet-display-name>"),
							Status:                      to.Ptr(armdomainservices.StatusWarning),
							ValidatorID:                 to.Ptr("<validator-id>"),
						}},
				},
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            to.Ptr(armdomainservices.NtlmV1Enabled),
					SyncNtlmPasswords: to.Ptr(armdomainservices.SyncNtlmPasswordsEnabled),
					TLSV1:             to.Ptr(armdomainservices.TLSV1Disabled),
				},
				FilteredSync: to.Ptr(armdomainservices.FilteredSyncEnabled),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         to.Ptr(armdomainservices.ExternalAccessEnabled),
					Ldaps:                  to.Ptr(armdomainservices.LdapsEnabled),
					PfxCertificate:         to.Ptr("<pfx-certificate>"),
					PfxCertificatePassword: to.Ptr("<pfx-certificate-password>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.Ptr("jicha@microsoft.com"),
						to.Ptr("caalmont@microsoft.com")},
					NotifyDcAdmins:     to.Ptr(armdomainservices.NotifyDcAdminsEnabled),
					NotifyGlobalAdmins: to.Ptr(armdomainservices.NotifyGlobalAdminsEnabled),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.Ptr("<location>"),
						SubnetID: to.Ptr("<subnet-id>"),
					},
					{
						Location: to.Ptr("<location>"),
						SubnetID: to.Ptr("<subnet-id>"),
					}},
			},
		},
		&armdomainservices.ClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
