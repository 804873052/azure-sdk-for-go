//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomersListByBillingProfile.json
func ExampleCustomersClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomersClient().NewListByBillingProfilePager("{billingAccountName}", "{billingProfileName}", &armbilling.CustomersClientListByBillingProfileOptions{Search: nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CustomerListResult = armbilling.CustomerListResult{
		// 	Value: []*armbilling.Customer{
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/22000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.CustomerProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				DisplayName: to.Ptr("customer1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000011"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/22000000-0000-0000-0000-000000000011"),
		// 			Properties: &armbilling.CustomerProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				DisplayName: to.Ptr("customer2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomersListByBillingAccount.json
func ExampleCustomersClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomersClient().NewListByBillingAccountPager("{billingAccountName}", &armbilling.CustomersClientListByBillingAccountOptions{Search: nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.CustomerListResult = armbilling.CustomerListResult{
		// 	Value: []*armbilling.Customer{
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/22000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.CustomerProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				DisplayName: to.Ptr("customer1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000011"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/22000000-0000-0000-0000-000000000011"),
		// 			Properties: &armbilling.CustomerProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				DisplayName: to.Ptr("customer2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Customer.json
func ExampleCustomersClient_Get_customer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomersClient().Get(ctx, "{billingAccountName}", "{customerName}", &armbilling.CustomersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Customer = armbilling.Customer{
	// 	Name: to.Ptr("{customerName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
	// 	Properties: &armbilling.CustomerProperties{
	// 		BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 		DisplayName: to.Ptr("customer1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerWithExpand.json
func ExampleCustomersClient_Get_customerWithExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomersClient().Get(ctx, "{billingAccountName}", "{customerName}", &armbilling.CustomersClientGetOptions{Expand: to.Ptr("enabledAzurePlans,resellers")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Customer = armbilling.Customer{
	// 	Name: to.Ptr("{customerName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/customers"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
	// 	Properties: &armbilling.CustomerProperties{
	// 		BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
	// 		DisplayName: to.Ptr("customerName1"),
	// 		EnabledAzurePlans: []*armbilling.AzurePlan{
	// 			{
	// 				SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
	// 				SKUID: to.Ptr("0002"),
	// 		}},
	// 		Resellers: []*armbilling.Reseller{
	// 			{
	// 				Description: to.Ptr("Reseller1"),
	// 				ResellerID: to.Ptr("89e87bdf-a2a2-4687-925f-4c18b27bccfd"),
	// 			},
	// 			{
	// 				Description: to.Ptr("Reseller2"),
	// 				ResellerID: to.Ptr("3b65b5a8-bd4f-4084-90e9-e1bd667a2b19"),
	// 		}},
	// 	},
	// }
}
