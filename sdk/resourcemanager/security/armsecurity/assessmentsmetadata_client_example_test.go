//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_example.json
func ExampleAssessmentsMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessmentsMetadataClient().NewListPager(nil)
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
		// page.AssessmentMetadataResponseList = armsecurity.AssessmentMetadataResponseList{
		// 	Value: []*armsecurity.AssessmentMetadataResponse{
		// 		{
		// 			Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
		// 				Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
		// 				AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
		// 				Categories: []*armsecurity.Categories{
		// 					to.Ptr(armsecurity.CategoriesCompute)},
		// 					DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
		// 					ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
		// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
		// 					RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
		// 					Severity: to.Ptr(armsecurity.SeverityMedium),
		// 					Threats: []*armsecurity.Threats{
		// 						to.Ptr(armsecurity.ThreatsDataExfiltration),
		// 						to.Ptr(armsecurity.ThreatsDataSpillage),
		// 						to.Ptr(armsecurity.ThreatsMaliciousInsider)},
		// 						UserImpact: to.Ptr(armsecurity.UserImpactLow),
		// 						PlannedDeprecationDate: to.Ptr("03/2022"),
		// 						PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
		// 							GA: to.Ptr("06/01/2021"),
		// 							Public: to.Ptr("06/01/2021"),
		// 						},
		// 						Tactics: []*armsecurity.Tactics{
		// 							to.Ptr(armsecurity.TacticsCredentialAccess),
		// 							to.Ptr(armsecurity.TacticsPersistence),
		// 							to.Ptr(armsecurity.TacticsExecution),
		// 							to.Ptr(armsecurity.TacticsDefenseEvasion),
		// 							to.Ptr(armsecurity.TacticsCollection),
		// 							to.Ptr(armsecurity.TacticsDiscovery),
		// 							to.Ptr(armsecurity.TacticsPrivilegeEscalation)},
		// 							Techniques: []*armsecurity.Techniques{
		// 								to.Ptr(armsecurity.TechniquesObfuscatedFilesOrInformation),
		// 								to.Ptr(armsecurity.TechniquesIngressToolTransfer),
		// 								to.Ptr(armsecurity.TechniquesPhishing),
		// 								to.Ptr(armsecurity.TechniquesUserExecution)},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("bc303248-3d14-44c2-96a0-55f5c326b5fe"),
		// 							Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
		// 							ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/bc303248-3d14-44c2-96a0-55f5c326b5fe"),
		// 							Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
		// 								Description: to.Ptr("Open remote management ports expose your VM to a high level of risk from internet-based attacks that attempt to brute force credentials to gain admin access to the machine."),
		// 								AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomPolicy),
		// 								Categories: []*armsecurity.Categories{
		// 									to.Ptr(armsecurity.CategoriesNetworking)},
		// 									DisplayName: to.Ptr("Close management ports on your virtual machines"),
		// 									ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
		// 									PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/22730e10-96f6-4aac-ad84-9383d35b5917"),
		// 									Preview: to.Ptr(true),
		// 									RemediationDescription: to.Ptr("We recommend that you edit the inbound rules of the below virtual machines to restrict access to specific source ranges.<br>To restrict the access to your virtual machines: 1. Click on a VM from the list below 2. At the 'Networking' blade, click on each of the rules that allow management ports (e.g. RDP-3389, WINRM-5985, SSH-22) 3. Change the 'Action' property to 'Deny' 4. Click 'Save'"),
		// 									Severity: to.Ptr(armsecurity.SeverityMedium),
		// 									Threats: []*armsecurity.Threats{
		// 										to.Ptr(armsecurity.ThreatsDataExfiltration),
		// 										to.Ptr(armsecurity.ThreatsDataSpillage),
		// 										to.Ptr(armsecurity.ThreatsMaliciousInsider)},
		// 										UserImpact: to.Ptr(armsecurity.UserImpactHigh),
		// 										PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
		// 											GA: to.Ptr("06/01/2021"),
		// 											Public: to.Ptr("06/01/2021"),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("ca039e75-a276-4175-aebc-bcd41e4b14b7"),
		// 									Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
		// 									ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/ca039e75-a276-4175-aebc-bcd41e4b14b7"),
		// 									Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
		// 										Description: to.Ptr("Assessment that my organization created to view our security assessment in Azure Security Center"),
		// 										AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
		// 										Categories: []*armsecurity.Categories{
		// 											to.Ptr(armsecurity.CategoriesCompute)},
		// 											DisplayName: to.Ptr("My organization security assessment"),
		// 											ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
		// 											RemediationDescription: to.Ptr("Fix it with these remediation instructions"),
		// 											Severity: to.Ptr(armsecurity.SeverityMedium),
		// 											Threats: []*armsecurity.Threats{
		// 											},
		// 											UserImpact: to.Ptr(armsecurity.UserImpactLow),
		// 											PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
		// 												GA: to.Ptr("06/01/2021"),
		// 												Public: to.Ptr("06/01/2021"),
		// 											},
		// 										},
		// 								}},
		// 							}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_example.json
func ExampleAssessmentsMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().Get(ctx, "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentMetadataResponse = armsecurity.AssessmentMetadataResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 		Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
	// 		AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
	// 		Categories: []*armsecurity.Categories{
	// 			to.Ptr(armsecurity.CategoriesCompute)},
	// 			DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 			ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
	// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
	// 			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
	// 				UserImpact: to.Ptr(armsecurity.UserImpactLow),
	// 				PlannedDeprecationDate: to.Ptr("03/2022"),
	// 				PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
	// 					GA: to.Ptr("06/01/2021"),
	// 					Public: to.Ptr("06/01/2021"),
	// 				},
	// 				Tactics: []*armsecurity.Tactics{
	// 					to.Ptr(armsecurity.TacticsCredentialAccess),
	// 					to.Ptr(armsecurity.TacticsPersistence),
	// 					to.Ptr(armsecurity.TacticsExecution),
	// 					to.Ptr(armsecurity.TacticsDefenseEvasion),
	// 					to.Ptr(armsecurity.TacticsCollection),
	// 					to.Ptr(armsecurity.TacticsDiscovery),
	// 					to.Ptr(armsecurity.TacticsPrivilegeEscalation)},
	// 					Techniques: []*armsecurity.Techniques{
	// 						to.Ptr(armsecurity.TechniquesObfuscatedFilesOrInformation),
	// 						to.Ptr(armsecurity.TechniquesIngressToolTransfer),
	// 						to.Ptr(armsecurity.TechniquesPhishing),
	// 						to.Ptr(armsecurity.TechniquesUserExecution)},
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessmentsMetadataClient().NewListBySubscriptionPager(nil)
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
		// page.AssessmentMetadataResponseList = armsecurity.AssessmentMetadataResponseList{
		// 	Value: []*armsecurity.AssessmentMetadataResponse{
		// 		{
		// 			Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
		// 			Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
		// 				Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
		// 				AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
		// 				Categories: []*armsecurity.Categories{
		// 					to.Ptr(armsecurity.CategoriesCompute)},
		// 					DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
		// 					ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
		// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
		// 					RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
		// 					Severity: to.Ptr(armsecurity.SeverityMedium),
		// 					Threats: []*armsecurity.Threats{
		// 						to.Ptr(armsecurity.ThreatsDataExfiltration),
		// 						to.Ptr(armsecurity.ThreatsDataSpillage),
		// 						to.Ptr(armsecurity.ThreatsMaliciousInsider)},
		// 						UserImpact: to.Ptr(armsecurity.UserImpactLow),
		// 						PlannedDeprecationDate: to.Ptr("03/2022"),
		// 						PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
		// 							GA: to.Ptr("06/01/2021"),
		// 							Public: to.Ptr("06/01/2021"),
		// 						},
		// 						Tactics: []*armsecurity.Tactics{
		// 							to.Ptr(armsecurity.TacticsCredentialAccess),
		// 							to.Ptr(armsecurity.TacticsPersistence),
		// 							to.Ptr(armsecurity.TacticsExecution),
		// 							to.Ptr(armsecurity.TacticsDefenseEvasion),
		// 							to.Ptr(armsecurity.TacticsCollection),
		// 							to.Ptr(armsecurity.TacticsDiscovery),
		// 							to.Ptr(armsecurity.TacticsPrivilegeEscalation)},
		// 							Techniques: []*armsecurity.Techniques{
		// 								to.Ptr(armsecurity.TechniquesObfuscatedFilesOrInformation),
		// 								to.Ptr(armsecurity.TechniquesIngressToolTransfer),
		// 								to.Ptr(armsecurity.TechniquesPhishing),
		// 								to.Ptr(armsecurity.TechniquesUserExecution)},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("bc303248-3d14-44c2-96a0-55f5c326b5fe"),
		// 							Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
		// 							ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/bc303248-3d14-44c2-96a0-55f5c326b5fe"),
		// 							Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
		// 								Description: to.Ptr("Open remote management ports expose your VM to a high level of risk from internet-based attacks that attempt to brute force credentials to gain admin access to the machine."),
		// 								AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomPolicy),
		// 								Categories: []*armsecurity.Categories{
		// 									to.Ptr(armsecurity.CategoriesNetworking)},
		// 									DisplayName: to.Ptr("Close management ports on your virtual machines"),
		// 									ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
		// 									PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/22730e10-96f6-4aac-ad84-9383d35b5917"),
		// 									Preview: to.Ptr(true),
		// 									RemediationDescription: to.Ptr("We recommend that you edit the inbound rules of the below virtual machines to restrict access to specific source ranges.<br>To restrict the access to your virtual machines: 1. Click on a VM from the list below 2. At the 'Networking' blade, click on each of the rules that allow management ports (e.g. RDP-3389, WINRM-5985, SSH-22) 3. Change the 'Action' property to 'Deny' 4. Click 'Save'"),
		// 									Severity: to.Ptr(armsecurity.SeverityMedium),
		// 									Threats: []*armsecurity.Threats{
		// 										to.Ptr(armsecurity.ThreatsDataExfiltration),
		// 										to.Ptr(armsecurity.ThreatsDataSpillage),
		// 										to.Ptr(armsecurity.ThreatsMaliciousInsider)},
		// 										UserImpact: to.Ptr(armsecurity.UserImpactHigh),
		// 										PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
		// 											GA: to.Ptr("06/01/2021"),
		// 											Public: to.Ptr("06/01/2021"),
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_GetInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().GetInSubscription(ctx, "21300918-b2e3-0346-785f-c77ff57d243b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentMetadataResponse = armsecurity.AssessmentMetadataResponse{
	// 	Name: to.Ptr("21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/21300918-b2e3-0346-785f-c77ff57d243b"),
	// 	Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 		Description: to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
	// 		AssessmentType: to.Ptr(armsecurity.AssessmentTypeBuiltIn),
	// 		Categories: []*armsecurity.Categories{
	// 			to.Ptr(armsecurity.CategoriesCompute)},
	// 			DisplayName: to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
	// 			ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
	// 			PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/26a828e1-e88f-464e-bbb3-c134a282b9de"),
	// 			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
	// 				UserImpact: to.Ptr(armsecurity.UserImpactLow),
	// 				PlannedDeprecationDate: to.Ptr("03/2022"),
	// 				PublishDates: &armsecurity.AssessmentMetadataPropertiesResponsePublishDates{
	// 					GA: to.Ptr("06/01/2021"),
	// 					Public: to.Ptr("06/01/2021"),
	// 				},
	// 				Tactics: []*armsecurity.Tactics{
	// 					to.Ptr(armsecurity.TacticsCredentialAccess),
	// 					to.Ptr(armsecurity.TacticsPersistence),
	// 					to.Ptr(armsecurity.TacticsExecution),
	// 					to.Ptr(armsecurity.TacticsDefenseEvasion),
	// 					to.Ptr(armsecurity.TacticsCollection),
	// 					to.Ptr(armsecurity.TacticsDiscovery),
	// 					to.Ptr(armsecurity.TacticsPrivilegeEscalation)},
	// 					Techniques: []*armsecurity.Techniques{
	// 						to.Ptr(armsecurity.TechniquesObfuscatedFilesOrInformation),
	// 						to.Ptr(armsecurity.TechniquesIngressToolTransfer),
	// 						to.Ptr(armsecurity.TechniquesPhishing),
	// 						to.Ptr(armsecurity.TechniquesUserExecution)},
	// 					},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/CreateAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_CreateInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsMetadataClient().CreateInSubscription(ctx, "ca039e75-a276-4175-aebc-bcd41e4b14b7", armsecurity.AssessmentMetadataResponse{
		Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
			Description:    to.Ptr("Install an endpoint protection solution on your virtual machines scale sets, to protect them from threats and vulnerabilities."),
			AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
			Categories: []*armsecurity.Categories{
				to.Ptr(armsecurity.CategoriesCompute)},
			DisplayName:            to.Ptr("Install endpoint protection solution on virtual machine scale sets"),
			ImplementationEffort:   to.Ptr(armsecurity.ImplementationEffortLow),
			RemediationDescription: to.Ptr("To install an endpoint protection solution: 1.  <a href=\"https://docs.microsoft.com/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-faq#how-do-i-turn-on-antimalware-in-my-virtual-machine-scale-set\">Follow the instructions in How do I turn on antimalware in my virtual machine scale set</a>"),
			Severity:               to.Ptr(armsecurity.SeverityMedium),
			Threats: []*armsecurity.Threats{
				to.Ptr(armsecurity.ThreatsDataExfiltration),
				to.Ptr(armsecurity.ThreatsDataSpillage),
				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
			UserImpact: to.Ptr(armsecurity.UserImpactLow),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessmentMetadataResponse = armsecurity.AssessmentMetadataResponse{
	// 	Name: to.Ptr("ca039e75-a276-4175-aebc-bcd41e4b14b7"),
	// 	Type: to.Ptr("Microsoft.Security/assessmentMetadata"),
	// 	ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/ca039e75-a276-4175-aebc-bcd41e4b14b7"),
	// 	Properties: &armsecurity.AssessmentMetadataPropertiesResponse{
	// 		Description: to.Ptr("Assessment that my organization created to view our security assessment in Azure Security Center"),
	// 		AssessmentType: to.Ptr(armsecurity.AssessmentTypeCustomerManaged),
	// 		Categories: []*armsecurity.Categories{
	// 			to.Ptr(armsecurity.CategoriesCompute)},
	// 			DisplayName: to.Ptr("My organization security assessment"),
	// 			ImplementationEffort: to.Ptr(armsecurity.ImplementationEffortLow),
	// 			RemediationDescription: to.Ptr("Fix it with these remediation instructions"),
	// 			Severity: to.Ptr(armsecurity.SeverityMedium),
	// 			Threats: []*armsecurity.Threats{
	// 				to.Ptr(armsecurity.ThreatsDataExfiltration),
	// 				to.Ptr(armsecurity.ThreatsDataSpillage),
	// 				to.Ptr(armsecurity.ThreatsMaliciousInsider)},
	// 				UserImpact: to.Ptr(armsecurity.UserImpactLow),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
func ExampleAssessmentsMetadataClient_DeleteInSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssessmentsMetadataClient().DeleteInSubscription(ctx, "ca039e75-a276-4175-aebc-bcd41e4b14b7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
