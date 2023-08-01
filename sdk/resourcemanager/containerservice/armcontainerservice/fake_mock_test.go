package armcontainerservice_test

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
	"net/http"
	"reflect"
	"testing"
)

type FakeTestSuite struct {
	suite.Suite

	cred    azcore.TokenCredential
	options arm.ClientOptions

	managedClustersClientServer            fake.ManagedClustersServer
	privateEndpointConnectionsClientServer fake.PrivateEndpointConnectionsServer
}

func (testsuite *FakeTestSuite) SetupSuite() {
	testsuite.cred = &testutil.FakeCredential{}

	testsuite.managedClustersClientServer = fake.ManagedClustersServer{
		BeginCreateOrUpdate: func(ctx context.Context, resourceGroupName string, resourceName string, parameters armcontainerservice.ManagedCluster, options *armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder) {
			resp = azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse]{}
			resp.SetTerminalResponse(http.StatusOK, armcontainerservice.ManagedClustersClientCreateOrUpdateResponse{
				ManagedCluster: armcontainerservice.ManagedCluster{
					Name:     to.Ptr("clustername1"),
					Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
					ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
					Location: to.Ptr("location1"),
					Tags: map[string]*string{
						"archv2": to.Ptr(""),
						"tier":   to.Ptr("production"),
					},
					Properties: &armcontainerservice.ManagedClusterProperties{
						AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
							{
								Type:  to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
								Count: to.Ptr[int32](3),
								CreationData: &armcontainerservice.CreationData{
									SourceResourceID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
								},
								CurrentOrchestratorVersion: to.Ptr("1.9.6"),
								EnableFIPS:                 to.Ptr(true),
								EnableNodePublicIP:         to.Ptr(true),
								MaxPods:                    to.Ptr[int32](110),
								Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
								NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
								OrchestratorVersion:        to.Ptr("1.9.6"),
								OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
								ProvisioningState:          to.Ptr("Succeeded"),
								VMSize:                     to.Ptr("Standard_DS2_v2"),
								Name:                       to.Ptr("nodepool1"),
							}},
						AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
							ScaleDownDelayAfterAdd: to.Ptr("15m"),
							ScanInterval:           to.Ptr("20s"),
						},
						CurrentKubernetesVersion: to.Ptr("1.9.6"),
						DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
						DNSPrefix:                to.Ptr("dnsprefix1"),
						EnablePodSecurityPolicy:  to.Ptr(false),
						EnableRBAC:               to.Ptr(true),
						Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
						KubernetesVersion:        to.Ptr("1.9.6"),
						LinuxProfile: &armcontainerservice.LinuxProfile{
							AdminUsername: to.Ptr("azureuser"),
							SSH: &armcontainerservice.SSHConfiguration{
								PublicKeys: []*armcontainerservice.SSHPublicKey{
									{
										KeyData: to.Ptr("keydata"),
									}},
							},
						},
						MaxAgentPools: to.Ptr[int32](1),
						NetworkProfile: &armcontainerservice.NetworkProfile{
							DNSServiceIP: to.Ptr("10.0.0.10"),
							IPFamilies: []*armcontainerservice.IPFamily{
								to.Ptr(armcontainerservice.IPFamilyIPv4)},
							LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
								AllocatedOutboundPorts: to.Ptr[int32](2000),
								EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
									{
										ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
									},
									{
										ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
									}},
								IdleTimeoutInMinutes: to.Ptr[int32](10),
								ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
									Count: to.Ptr[int32](2),
								},
							},
							LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
							NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
							OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
							PodCidr:         to.Ptr("10.244.0.0/16"),
							PodCidrs: []*string{
								to.Ptr("10.244.0.0/16")},
							ServiceCidr: to.Ptr("10.0.0.0/16"),
							ServiceCidrs: []*string{
								to.Ptr("10.0.0.0/16")},
						},
						NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
						ProvisioningState: to.Ptr("Succeeded"),
						ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
							ClientID: to.Ptr("clientid"),
						},
						WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
							AdminUsername: to.Ptr("azureuser"),
						},
					},
				},
			}, nil)
			return
		},

		Get: func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientGetOptions) (resp azfake.Responder[armcontainerservice.ManagedClustersClientGetResponse], errResp azfake.ErrorResponder) {
			resp = azfake.Responder[armcontainerservice.ManagedClustersClientGetResponse]{}
			resp.SetResponse(http.StatusOK, armcontainerservice.ManagedClustersClientGetResponse{
				ManagedCluster: armcontainerservice.ManagedCluster{
					Name:     to.Ptr("clustername1"),
					Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
					ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
					Location: to.Ptr("location1"),
					Tags: map[string]*string{
						"archv2": to.Ptr(""),
						"tier":   to.Ptr("production"),
					},
					Properties: &armcontainerservice.ManagedClusterProperties{
						AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
							{
								AvailabilityZones: []*string{
									to.Ptr("1"),
									to.Ptr("2"),
									to.Ptr("3")},
								Count:                      to.Ptr[int32](3),
								CurrentOrchestratorVersion: to.Ptr("1.9.6"),
								MaxPods:                    to.Ptr[int32](110),
								NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
								OrchestratorVersion:        to.Ptr("1.9.6"),
								OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
								ProvisioningState:          to.Ptr("Succeeded"),
								UpgradeSettings: &armcontainerservice.AgentPoolUpgradeSettings{
									MaxSurge: to.Ptr("33%"),
								},
								VMSize: to.Ptr("Standard_DS1_v2"),
								Name:   to.Ptr("nodepool1"),
							}},
						AzurePortalFQDN:          to.Ptr("dnsprefix1-abcd1234.portal.hcp.eastus.azmk8s.io"),
						CurrentKubernetesVersion: to.Ptr("1.9.6"),
						DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
						DNSPrefix:                to.Ptr("dnsprefix1"),
						EnableRBAC:               to.Ptr(false),
						Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
						KubernetesVersion:        to.Ptr("1.9.6"),
						LinuxProfile: &armcontainerservice.LinuxProfile{
							AdminUsername: to.Ptr("azureuser"),
							SSH: &armcontainerservice.SSHConfiguration{
								PublicKeys: []*armcontainerservice.SSHPublicKey{
									{
										KeyData: to.Ptr("keydata"),
									}},
							},
						},
						MaxAgentPools: to.Ptr[int32](1),
						NetworkProfile: &armcontainerservice.NetworkProfile{
							DNSServiceIP: to.Ptr("10.0.0.10"),
							IPFamilies: []*armcontainerservice.IPFamily{
								to.Ptr(armcontainerservice.IPFamilyIPv4)},
							LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
								AllocatedOutboundPorts: to.Ptr[int32](2000),
								EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
									{
										ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
									},
									{
										ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
									}},
								IdleTimeoutInMinutes: to.Ptr[int32](10),
								OutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileOutboundIPs{
									PublicIPs: []*armcontainerservice.ResourceReference{
										{
											ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip1"),
										},
										{
											ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip2"),
										}},
								},
							},
							LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
							NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
							OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
							PodCidr:         to.Ptr("10.244.0.0/16"),
							PodCidrs: []*string{
								to.Ptr("10.244.0.0/16")},
							ServiceCidr: to.Ptr("10.0.0.0/16"),
							ServiceCidrs: []*string{
								to.Ptr("10.0.0.0/16")},
						},
						NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
						ProvisioningState: to.Ptr("Succeeded"),
						ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
							ClientID: to.Ptr("clientid"),
						},
					},
				},
			}, nil)
			return
		},

		BeginDelete: func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservice.ManagedClustersClientDeleteResponse], errResp azfake.ErrorResponder) {
			resp = azfake.PollerResponder[armcontainerservice.ManagedClustersClientDeleteResponse]{}
			// unexpected status code 200. acceptable values are http.StatusAccepted, http.StatusNoContent
			resp.SetTerminalResponse(http.StatusAccepted, armcontainerservice.ManagedClustersClientDeleteResponse{}, nil)
			return
		},

		NewListPager: func(options *armcontainerservice.ManagedClustersClientListOptions) (resp azfake.PagerResponder[armcontainerservice.ManagedClustersClientListResponse]) {
			resp = azfake.PagerResponder[armcontainerservice.ManagedClustersClientListResponse]{}
			resp.AddPage(http.StatusOK, armcontainerservice.ManagedClustersClientListResponse{
				ManagedClusterListResult: armcontainerservice.ManagedClusterListResult{
					Value: []*armcontainerservice.ManagedCluster{
						{
							Name:     to.Ptr("clustername1"),
							Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
							ID:       to.Ptr("/subscriptions/subid1/providers/Microsoft.ContainerService/managedClusters"),
							Location: to.Ptr("location1"),
							Tags: map[string]*string{
								"archv2": to.Ptr(""),
								"tier":   to.Ptr("production"),
							},
							Properties: &armcontainerservice.ManagedClusterProperties{
								AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
									{
										Count:                      to.Ptr[int32](3),
										CurrentOrchestratorVersion: to.Ptr("1.9.6"),
										MaxPods:                    to.Ptr[int32](110),
										OrchestratorVersion:        to.Ptr("1.9.6"),
										OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
										ProvisioningState:          to.Ptr("Succeeded"),
										VMSize:                     to.Ptr("Standard_DS1_v2"),
										Name:                       to.Ptr("nodepool1"),
									}},
								CurrentKubernetesVersion: to.Ptr("1.9.6"),
								DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
								DNSPrefix:                to.Ptr("dnsprefix1"),
								EnableRBAC:               to.Ptr(false),
								Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
								KubernetesVersion:        to.Ptr("1.9.6"),
								LinuxProfile: &armcontainerservice.LinuxProfile{
									AdminUsername: to.Ptr("azureuser"),
									SSH: &armcontainerservice.SSHConfiguration{
										PublicKeys: []*armcontainerservice.SSHPublicKey{
											{
												KeyData: to.Ptr("keydata"),
											}},
									},
								},
								MaxAgentPools: to.Ptr[int32](1),
								NetworkProfile: &armcontainerservice.NetworkProfile{
									DNSServiceIP:  to.Ptr("10.0.0.10"),
									NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
									PodCidr:       to.Ptr("10.244.0.0/16"),
									ServiceCidr:   to.Ptr("10.0.0.0/16"),
								},
								NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
								ProvisioningState: to.Ptr("Succeeded"),
								ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
									ClientID: to.Ptr("clientid"),
								},
							},
						}},
				},
			}, nil)

			return
		},
	}

	testsuite.privateEndpointConnectionsClientServer = fake.PrivateEndpointConnectionsServer{
		Update: func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, parameters armcontainerservice.PrivateEndpointConnection, options *armcontainerservice.PrivateEndpointConnectionsClientUpdateOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientUpdateResponse], errResp azfake.ErrorResponder) {
			resp.SetResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientUpdateResponse{
				PrivateEndpointConnection: armcontainerservice.PrivateEndpointConnection{
					Name: to.Ptr("privateendpointconnection1"),
					Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
					ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
					Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
						PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
							ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
						},
						PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
							Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
						},
						ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
					},
				},
			}, nil)
			return
		},

		Get: func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armcontainerservice.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder) {
			resp = azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientGetResponse]{}
			resp.SetResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientGetResponse{
				PrivateEndpointConnection: armcontainerservice.PrivateEndpointConnection{
					Name: to.Ptr("privateendpointconnection1"),
					Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
					ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
					Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
						PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
							ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
						},
						PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
							Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
						},
						ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
					},
				},
			}, nil)
			return
		},

		List: func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.PrivateEndpointConnectionsClientListOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientListResponse], errResp azfake.ErrorResponder) {
			resp = azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientListResponse]{}
			resp.SetResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientListResponse{
				PrivateEndpointConnectionListResult: armcontainerservice.PrivateEndpointConnectionListResult{
					Value: []*armcontainerservice.PrivateEndpointConnection{
						{
							Name: to.Ptr("privateendpointconnection1"),
							Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
							ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
							Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
								PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
									ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
								},
								PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
									Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
								},
								ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
							},
						},
					},
				},
			}, nil)

			return
		},

		BeginDelete: func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armcontainerservice.PrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder) {
			resp = azfake.PollerResponder[armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse]{}
			//resp.AddNonTerminalResponse(http.StatusOK, nil)
			//resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.SetTerminalResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse{}, nil)
			return
		},
	}
}

func TestFakeTest(t *testing.T) {
	suite.Run(t, new(FakeTestSuite))
}

func (testsuite *FakeTestSuite) TestManagedClusters_List() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"List Managed Clusters"},
	})
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&testsuite.managedClustersClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json")
		// Response check
		pagerExampleRes := armcontainerservice.ManagedClusterListResult{
			Value: []*armcontainerservice.ManagedCluster{
				{
					Name:     to.Ptr("clustername1"),
					Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
					ID:       to.Ptr("/subscriptions/subid1/providers/Microsoft.ContainerService/managedClusters"),
					Location: to.Ptr("location1"),
					Tags: map[string]*string{
						"archv2": to.Ptr(""),
						"tier":   to.Ptr("production"),
					},
					Properties: &armcontainerservice.ManagedClusterProperties{
						AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
							{
								Count:                      to.Ptr[int32](3),
								CurrentOrchestratorVersion: to.Ptr("1.9.6"),
								MaxPods:                    to.Ptr[int32](110),
								OrchestratorVersion:        to.Ptr("1.9.6"),
								OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
								ProvisioningState:          to.Ptr("Succeeded"),
								VMSize:                     to.Ptr("Standard_DS1_v2"),
								Name:                       to.Ptr("nodepool1"),
							}},
						CurrentKubernetesVersion: to.Ptr("1.9.6"),
						DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
						DNSPrefix:                to.Ptr("dnsprefix1"),
						EnableRBAC:               to.Ptr(false),
						Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
						KubernetesVersion:        to.Ptr("1.9.6"),
						LinuxProfile: &armcontainerservice.LinuxProfile{
							AdminUsername: to.Ptr("azureuser"),
							SSH: &armcontainerservice.SSHConfiguration{
								PublicKeys: []*armcontainerservice.SSHPublicKey{
									{
										KeyData: to.Ptr("keydata"),
									}},
							},
						},
						MaxAgentPools: to.Ptr[int32](1),
						NetworkProfile: &armcontainerservice.NetworkProfile{
							DNSServiceIP:  to.Ptr("10.0.0.10"),
							NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
							PodCidr:       to.Ptr("10.244.0.0/16"),
							ServiceCidr:   to.Ptr("10.0.0.0/16"),
						},
						NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
						ProvisioningState: to.Ptr("Succeeded"),
						ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
							ClientID: to.Ptr("clientid"),
						},
					},
				}},
		}
		if !reflect.DeepEqual(pagerExampleRes, nextResult.ManagedClusterListResult) {
			exampleResJson, _ := json.Marshal(pagerExampleRes)
			mockResJson, _ := json.Marshal(nextResult.ManagedClusterListResult)
			testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
		}
	}
}

func (testsuite *FakeTestSuite) TestManagedClusters_Get() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersGet.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Get Managed Cluster"},
	})
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&testsuite.managedClustersClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	res, err := client.Get(ctx, "rg1", "clustername1", nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersGet.json")
	// Response check
	exampleRes := armcontainerservice.ManagedCluster{
		Name:     to.Ptr("clustername1"),
		Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
		ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					AvailabilityZones: []*string{
						to.Ptr("1"),
						to.Ptr("2"),
						to.Ptr("3")},
					Count:                      to.Ptr[int32](3),
					CurrentOrchestratorVersion: to.Ptr("1.9.6"),
					MaxPods:                    to.Ptr[int32](110),
					NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
					OrchestratorVersion:        to.Ptr("1.9.6"),
					OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
					ProvisioningState:          to.Ptr("Succeeded"),
					UpgradeSettings: &armcontainerservice.AgentPoolUpgradeSettings{
						MaxSurge: to.Ptr("33%"),
					},
					VMSize: to.Ptr("Standard_DS1_v2"),
					Name:   to.Ptr("nodepool1"),
				}},
			AzurePortalFQDN:          to.Ptr("dnsprefix1-abcd1234.portal.hcp.eastus.azmk8s.io"),
			CurrentKubernetesVersion: to.Ptr("1.9.6"),
			DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
			DNSPrefix:                to.Ptr("dnsprefix1"),
			EnableRBAC:               to.Ptr(false),
			Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
			KubernetesVersion:        to.Ptr("1.9.6"),
			LinuxProfile: &armcontainerservice.LinuxProfile{
				AdminUsername: to.Ptr("azureuser"),
				SSH: &armcontainerservice.SSHConfiguration{
					PublicKeys: []*armcontainerservice.SSHPublicKey{
						{
							KeyData: to.Ptr("keydata"),
						}},
				},
			},
			MaxAgentPools: to.Ptr[int32](1),
			NetworkProfile: &armcontainerservice.NetworkProfile{
				DNSServiceIP: to.Ptr("10.0.0.10"),
				IPFamilies: []*armcontainerservice.IPFamily{
					to.Ptr(armcontainerservice.IPFamilyIPv4)},
				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
					AllocatedOutboundPorts: to.Ptr[int32](2000),
					EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
						},
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
						}},
					IdleTimeoutInMinutes: to.Ptr[int32](10),
					OutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileOutboundIPs{
						PublicIPs: []*armcontainerservice.ResourceReference{
							{
								ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip1"),
							},
							{
								ID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip2"),
							}},
					},
				},
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
				NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
				OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
				PodCidr:         to.Ptr("10.244.0.0/16"),
				PodCidrs: []*string{
					to.Ptr("10.244.0.0/16")},
				ServiceCidr: to.Ptr("10.0.0.0/16"),
				ServiceCidrs: []*string{
					to.Ptr("10.0.0.0/16")},
			},
			NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
			ProvisioningState: to.Ptr("Succeeded"),
			ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.Ptr("clientid"),
			},
		},
	}
	if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
		exampleResJson, _ := json.Marshal(exampleRes)
		mockResJson, _ := json.Marshal(res.ManagedCluster)
		testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersGet.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	}
}

func (testsuite *FakeTestSuite) TestManagedClusters_CreateOrUpdate() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Create Managed Cluster using an agent pool snapshot"},
	})
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&testsuite.managedClustersClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type:  to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count: to.Ptr[int32](3),
					CreationData: &armcontainerservice.CreationData{
						SourceResourceID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
					},
					EnableFIPS:         to.Ptr(true),
					EnableNodePublicIP: to.Ptr(true),
					Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
					OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
					VMSize:             to.Ptr("Standard_DS2_v2"),
					Name:               to.Ptr("nodepool1"),
				}},
			AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
				ScaleDownDelayAfterAdd: to.Ptr("15m"),
				ScanInterval:           to.Ptr("20s"),
			},
			DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
			DNSPrefix:               to.Ptr("dnsprefix1"),
			EnablePodSecurityPolicy: to.Ptr(false),
			EnableRBAC:              to.Ptr(true),
			KubernetesVersion:       to.Ptr(""),
			LinuxProfile: &armcontainerservice.LinuxProfile{
				AdminUsername: to.Ptr("azureuser"),
				SSH: &armcontainerservice.SSHConfiguration{
					PublicKeys: []*armcontainerservice.SSHPublicKey{
						{
							KeyData: to.Ptr("keydata"),
						}},
				},
			},
			NetworkProfile: &armcontainerservice.NetworkProfile{
				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
					ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
						Count: to.Ptr[int32](2),
					},
				},
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
				OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
			},
			ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.Ptr("clientid"),
				Secret:   to.Ptr("secret"),
			},
			WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
				AdminPassword: to.Ptr("replacePassword1234$"),
				AdminUsername: to.Ptr("azureuser"),
			},
		},
		SKU: &armcontainerservice.ManagedClusterSKU{
			Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
			Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
		},
	}, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json")
	// Response check
	exampleRes := armcontainerservice.ManagedCluster{
		Name:     to.Ptr("clustername1"),
		Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
		ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type:  to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count: to.Ptr[int32](3),
					CreationData: &armcontainerservice.CreationData{
						SourceResourceID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
					},
					CurrentOrchestratorVersion: to.Ptr("1.9.6"),
					EnableFIPS:                 to.Ptr(true),
					EnableNodePublicIP:         to.Ptr(true),
					MaxPods:                    to.Ptr[int32](110),
					Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
					NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
					OrchestratorVersion:        to.Ptr("1.9.6"),
					OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
					ProvisioningState:          to.Ptr("Succeeded"),
					VMSize:                     to.Ptr("Standard_DS2_v2"),
					Name:                       to.Ptr("nodepool1"),
				}},
			AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
				ScaleDownDelayAfterAdd: to.Ptr("15m"),
				ScanInterval:           to.Ptr("20s"),
			},
			CurrentKubernetesVersion: to.Ptr("1.9.6"),
			DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
			DNSPrefix:                to.Ptr("dnsprefix1"),
			EnablePodSecurityPolicy:  to.Ptr(false),
			EnableRBAC:               to.Ptr(true),
			Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
			KubernetesVersion:        to.Ptr("1.9.6"),
			LinuxProfile: &armcontainerservice.LinuxProfile{
				AdminUsername: to.Ptr("azureuser"),
				SSH: &armcontainerservice.SSHConfiguration{
					PublicKeys: []*armcontainerservice.SSHPublicKey{
						{
							KeyData: to.Ptr("keydata"),
						}},
				},
			},
			MaxAgentPools: to.Ptr[int32](1),
			NetworkProfile: &armcontainerservice.NetworkProfile{
				DNSServiceIP: to.Ptr("10.0.0.10"),
				IPFamilies: []*armcontainerservice.IPFamily{
					to.Ptr(armcontainerservice.IPFamilyIPv4)},
				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
					AllocatedOutboundPorts: to.Ptr[int32](2000),
					EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
						},
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
						}},
					IdleTimeoutInMinutes: to.Ptr[int32](10),
					ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
						Count: to.Ptr[int32](2),
					},
				},
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
				NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
				OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
				PodCidr:         to.Ptr("10.244.0.0/16"),
				PodCidrs: []*string{
					to.Ptr("10.244.0.0/16")},
				ServiceCidr: to.Ptr("10.0.0.0/16"),
				ServiceCidrs: []*string{
					to.Ptr("10.0.0.0/16")},
			},
			NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
			ProvisioningState: to.Ptr("Succeeded"),
			ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.Ptr("clientid"),
			},
			WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
				AdminUsername: to.Ptr("azureuser"),
			},
		},
	}
	if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
		exampleResJson, _ := json.Marshal(exampleRes)
		mockResJson, _ := json.Marshal(res.ManagedCluster)
		testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	}

	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with AKS-managed NAT gateway as outbound type"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(false),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			NatGatewayProfile: &armcontainerservice.ManagedClusterNATGatewayProfile{
	//				ManagedOutboundIPProfile: &armcontainerservice.ManagedClusterManagedOutboundIPProfile{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			OutboundType: to.Ptr(armcontainerservice.OutboundTypeManagedNATGateway),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(false),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP:    to.Ptr("10.0.0.10"),
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NatGatewayProfile: &armcontainerservice.ManagedClusterNATGatewayProfile{
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](4),
	//				ManagedOutboundIPProfile: &armcontainerservice.ManagedClusterManagedOutboundIPProfile{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:  to.Ptr(armcontainerservice.OutboundTypeManagedNATGateway),
	//			PodCidr:       to.Ptr("10.244.0.0/16"),
	//			ServiceCidr:   to.Ptr("10.0.0.0/16"),
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_AzureKeyvaultSecretsProvider.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with Azure KeyVault Secrets Provider Addon"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{
	//			"azureKeyvaultSecretsProvider": &armcontainerservice.ManagedClusterAddonProfile{
	//				Config: map[string]*string{
	//					"enableSecretRotation": to.Ptr("true"),
	//					"rotationPollInterval": to.Ptr("2m"),
	//				},
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_AzureKeyvaultSecretsProvider.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_AzureKeyvaultSecretsProvider.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{
	//			"azureKeyvaultSecretsProvider": &armcontainerservice.ManagedClusterAddonProfile{
	//				Config: map[string]*string{
	//					"enableSecretRotation": to.Ptr("true"),
	//					"rotationPollInterval": to.Ptr("2m"),
	//				},
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				MaxPods:                to.Ptr[int32](110),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:       to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:    to.Ptr("1.9.6"),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:      to.Ptr("Succeeded"),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		Fqdn:                    to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:       to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			ServiceCidr:     to.Ptr("10.0.0.0/16"),
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_AzureKeyvaultSecretsProvider.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DedicatedHostGroup.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with Dedicated Host Group"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				HostGroupID:        to.Ptr("/subscriptions/subid1/resourcegroups/rg/providers/Microsoft.Compute/hostGroups/hostgroup1"),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(false),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DedicatedHostGroup.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DedicatedHostGroup.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:               to.Ptr[int32](3),
	//				EnableNodePublicIP:  to.Ptr(true),
	//				HostGroupID:         to.Ptr("/subscriptions/subid1/resourcegroups/rg/providers/Microsoft.Compute/hostGroups/hostgroup1"),
	//				MaxPods:             to.Ptr[int32](110),
	//				NodeImageVersion:    to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion: to.Ptr("1.9.6"),
	//				OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:   to.Ptr("Succeeded"),
	//				VMSize:              to.Ptr("Standard_DS2_v2"),
	//				Name:                to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(false),
	//		EnableRBAC:              to.Ptr(true),
	//		Fqdn:                    to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:       to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DedicatedHostGroup.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableEncryptionAtHost.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with EncryptionAtHost enabled"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableEncryptionAtHost.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableEncryptionAtHost.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableEncryptionAtHost:     to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableEncryptionAtHost.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnabledFIPS.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with FIPS enabled OS"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableFIPS:         to.Ptr(true),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(false),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnabledFIPS.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnabledFIPS.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableFIPS:                 to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(false),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnabledFIPS.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_GPUMIG.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with GPUMIG"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				GpuInstanceProfile: to.Ptr(armcontainerservice.GPUInstanceProfileMIG3G),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_ND96asr_v4"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_GPUMIG.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_GPUMIG.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				GpuInstanceProfile:         to.Ptr(armcontainerservice.GPUInstanceProfileMIG3G),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_ND96asr_v4"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_GPUMIG.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_HTTPProxy.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with HTTP proxy configured"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_HTTPProxy.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_HTTPProxy.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_HTTPProxy.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Premium.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with LongTermSupport"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			DisableRunCommand: to.Ptr(true),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		SupportPlan: to.Ptr(armcontainerservice.KubernetesSupportPlanAKSLongTermSupport),
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBase),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierPremium),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Premium.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Premium.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableEncryptionAtHost:     to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			DisableRunCommand: to.Ptr(true),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		Fqdn:                    to.Ptr("dnsprefix1-ee788a1f.hcp.location1.azmk8s.io"),
	//		KubernetesVersion:       to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		PrivateFQDN:       to.Ptr("dnsprefix1-aae7e0f0.5cef6058-b6b5-414d-8cb1-4bd14eb0b15c.privatelink.location1.azmk8s.io"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		SupportPlan: to.Ptr(armcontainerservice.KubernetesSupportPlanAKSLongTermSupport),
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBase),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierPremium),
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Premium.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_NodePublicIPPrefix.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with Node Public IP Prefix"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                 to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                to.Ptr[int32](3),
	//				EnableNodePublicIP:   to.Ptr(true),
	//				Mode:                 to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodePublicIPPrefixID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Network/publicIPPrefixes/public-ip-prefix"),
	//				OSType:               to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:               to.Ptr("Standard_DS2_v2"),
	//				Name:                 to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_NodePublicIPPrefix.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_NodePublicIPPrefix.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				NodePublicIPPrefixID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Network/publicIPPrefixes/public-ip-prefix"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_NodePublicIPPrefix.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_OSSKU.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with OSSKU"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSSKU:              to.Ptr(armcontainerservice.OSSKUAzureLinux),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_OSSKU.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_OSSKU.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSSKU:                      to.Ptr(armcontainerservice.OSSKUAzureLinux),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		HTTPProxyConfig: &armcontainerservice.ManagedClusterHTTPProxyConfig{
	//			HTTPProxy:  to.Ptr("http://myproxy.server.com:8080"),
	//			HTTPSProxy: to.Ptr("https://myproxy.server.com:8080"),
	//			NoProxy: []*string{
	//				to.Ptr("localhost"),
	//				to.Ptr("127.0.0.1")},
	//			TrustedCa: to.Ptr("Q29uZ3JhdHMhIFlvdSBoYXZlIGZvdW5kIGEgaGlkZGVuIG1lc3NhZ2U="),
	//		},
	//		KubernetesVersion: to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_OSSKU.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PPG.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with PPG"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                      to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                     to.Ptr[int32](3),
	//				EnableNodePublicIP:        to.Ptr(true),
	//				Mode:                      to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                    to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProximityPlacementGroupID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
	//				VMSize:                    to.Ptr("Standard_DS2_v2"),
	//				Name:                      to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PPG.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PPG.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				ProximityPlacementGroupID:  to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PPG.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PodIdentity.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with PodIdentity enabled"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		PodIdentityProfile: &armcontainerservice.ManagedClusterPodIdentityProfile{
	//			AllowNetworkPluginKubenet: to.Ptr(true),
	//			Enabled:                   to.Ptr(true),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PodIdentity.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PodIdentity.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		PodIdentityProfile: &armcontainerservice.ManagedClusterPodIdentityProfile{
	//			AllowNetworkPluginKubenet: to.Ptr(true),
	//			Enabled:                   to.Ptr(true),
	//		},
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PodIdentity.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DisableRunCommand.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with RunCommand disabled"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			DisableRunCommand: to.Ptr(true),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DisableRunCommand.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DisableRunCommand.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableEncryptionAtHost:     to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			DisableRunCommand: to.Ptr(true),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-ee788a1f.hcp.location1.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		PrivateFQDN:       to.Ptr("dnsprefix1-aae7e0f0.5cef6058-b6b5-414d-8cb1-4bd14eb0b15c.privatelink.location1.azmk8s.io"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		SupportPlan: to.Ptr(armcontainerservice.KubernetesSupportPlanKubernetesOfficial),
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DisableRunCommand.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_SecurityProfile.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with Security Profile configured"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		DNSPrefix:         to.Ptr("dnsprefix1"),
	//		KubernetesVersion: to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		SecurityProfile: &armcontainerservice.ManagedClusterSecurityProfile{
	//			Defender: &armcontainerservice.ManagedClusterSecurityProfileDefender{
	//				LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/SUB_ID/resourcegroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME"),
	//				SecurityMonitoring: &armcontainerservice.ManagedClusterSecurityProfileDefenderSecurityMonitoring{
	//					Enabled: to.Ptr(true),
	//				},
	//			},
	//			WorkloadIdentity: &armcontainerservice.ManagedClusterSecurityProfileWorkloadIdentity{
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_SecurityProfile.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_SecurityProfile.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		SecurityProfile: &armcontainerservice.ManagedClusterSecurityProfile{
	//			Defender: &armcontainerservice.ManagedClusterSecurityProfileDefender{
	//				LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/SUB_ID/resourcegroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME"),
	//				SecurityMonitoring: &armcontainerservice.ManagedClusterSecurityProfileDefenderSecurityMonitoring{
	//					Enabled: to.Ptr(true),
	//				},
	//			},
	//			WorkloadIdentity: &armcontainerservice.ManagedClusterSecurityProfileWorkloadIdentity{
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_SecurityProfile.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableUltraSSD.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with UltraSSD enabled"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				EnableUltraSSD:     to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableUltraSSD.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableUltraSSD.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				EnableUltraSSD:             to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_EnableUltraSSD.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UserAssignedNATGateway.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Cluster with user-assigned NAT gateway as outbound type"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(false),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS2_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeUserAssignedNATGateway),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UserAssignedNATGateway.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UserAssignedNATGateway.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(false),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP:    to.Ptr("10.0.0.10"),
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeUserAssignedNATGateway),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			ServiceCidr:     to.Ptr("10.0.0.0/16"),
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UserAssignedNATGateway.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterPublicFQDN.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Private Cluster with Public FQDN specified"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			EnablePrivateCluster:           to.Ptr(true),
	//			EnablePrivateClusterPublicFQDN: to.Ptr(true),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterPublicFQDN.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterPublicFQDN.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableEncryptionAtHost:     to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			EnablePrivateCluster:           to.Ptr(true),
	//			EnablePrivateClusterPublicFQDN: to.Ptr(true),
	//			PrivateDNSZone:                 to.Ptr("system"),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-ee788a1f.hcp.location1.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		PrivateFQDN:       to.Ptr("dnsprefix1-aae7e0f0.5cef6058-b6b5-414d-8cb1-4bd14eb0b15c.privatelink.location1.azmk8s.io"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterPublicFQDN.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterFQDNSubdomain.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create Managed Private Cluster with fqdn subdomain specified"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                  to.Ptr[int32](3),
	//				EnableEncryptionAtHost: to.Ptr(true),
	//				EnableNodePublicIP:     to.Ptr(true),
	//				Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:                 to.Ptr("Standard_DS2_v2"),
	//				Name:                   to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			EnablePrivateCluster: to.Ptr(true),
	//			PrivateDNSZone:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Network/privateDnsZones/privatelink.location1.azmk8s.io"),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		FqdnSubdomain:           to.Ptr("domain1"),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterFQDNSubdomain.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterFQDNSubdomain.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableEncryptionAtHost:     to.Ptr(true),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS2_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
	//			EnablePrivateCluster: to.Ptr(true),
	//			PrivateDNSZone:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Network/privateDnsZones/privatelink.location1.azmk8s.io"),
	//		},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		FqdnSubdomain:            to.Ptr("domain1"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		PrivateFQDN:       to.Ptr("domain1.privatelink.location1.azmk8s.io"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_PrivateClusterFQDNSubdomain.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithEnableAzureRBAC.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create/Update AAD Managed Cluster with EnableAzureRBAC"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AADProfile: &armcontainerservice.ManagedClusterAADProfile{
	//			EnableAzureRBAC: to.Ptr(true),
	//			Managed:         to.Ptr(true),
	//		},
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS1_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithEnableAzureRBAC.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithEnableAzureRBAC.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AADProfile: &armcontainerservice.ManagedClusterAADProfile{
	//			EnableAzureRBAC: to.Ptr(true),
	//			Managed:         to.Ptr(true),
	//			TenantID:        to.Ptr("tenantID"),
	//		},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS1_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithEnableAzureRBAC.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Update.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create/Update Managed Cluster"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				ScaleDownMode:      to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
	//				VMSize:             to.Ptr("Standard_DS1_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			BalanceSimilarNodeGroups: to.Ptr("true"),
	//			Expander:                 to.Ptr(armcontainerservice.ExpanderPriority),
	//			MaxNodeProvisionTime:     to.Ptr("15m"),
	//			NewPodScaleUpDelay:       to.Ptr("1m"),
	//			ScaleDownDelayAfterAdd:   to.Ptr("15m"),
	//			ScanInterval:             to.Ptr("20s"),
	//			SkipNodesWithSystemPods:  to.Ptr("false"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Update.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Update.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//				ClientID:    to.Ptr("clientId1"),
	//				PrincipalID: to.Ptr("principalId1"),
	//			},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				ScaleDownMode:              to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
	//				VMSize:                     to.Ptr("Standard_DS1_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			BalanceSimilarNodeGroups: to.Ptr("true"),
	//			Expander:                 to.Ptr(armcontainerservice.ExpanderPriority),
	//			MaxNodeProvisionTime:     to.Ptr("15m"),
	//			NewPodScaleUpDelay:       to.Ptr("1m"),
	//			ScaleDownDelayAfterAdd:   to.Ptr("15m"),
	//			ScanInterval:             to.Ptr("20s"),
	//			SkipNodesWithSystemPods:  to.Ptr("false"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Update.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithAHUB.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create/Update Managed Cluster with EnableAHUB"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS1_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//			LicenseType:   to.Ptr(armcontainerservice.LicenseTypeWindowsServer),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithAHUB.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithAHUB.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//				ClientID:    to.Ptr("clientId1"),
	//				PrincipalID: to.Ptr("principalId1"),
	//			},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS1_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			LicenseType:   to.Ptr(armcontainerservice.LicenseTypeWindowsServer),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWithAHUB.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWindowsGmsa.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create/Update Managed Cluster with Windows gMSA enabled"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				VMSize:             to.Ptr("Standard_DS1_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//			GmsaProfile: &armcontainerservice.WindowsGmsaProfile{
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWindowsGmsa.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWindowsGmsa.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//				ClientID:    to.Ptr("clientId1"),
	//				PrincipalID: to.Ptr("principalId1"),
	//			},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.9.6"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				VMSize:                     to.Ptr("Standard_DS1_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			ScaleDownDelayAfterAdd: to.Ptr("15m"),
	//			ScanInterval:           to.Ptr("20s"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.9.6"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			GmsaProfile: &armcontainerservice.WindowsGmsaProfile{
	//				Enabled: to.Ptr(true),
	//			},
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_UpdateWindowsGmsa.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
	//
	//// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DualStackNetworking.json
	//ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
	//	"example-id": {"Create/Update Managed Cluster with dual-stack networking"},
	//})
	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &testsuite.options)
	//testsuite.Require().NoError(err, "Failed to create client")
	//poller, err = client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:              to.Ptr[int32](3),
	//				EnableNodePublicIP: to.Ptr(true),
	//				Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
	//				ScaleDownMode:      to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
	//				VMSize:             to.Ptr("Standard_DS1_v2"),
	//				Name:               to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			BalanceSimilarNodeGroups: to.Ptr("true"),
	//			Expander:                 to.Ptr(armcontainerservice.ExpanderPriority),
	//			MaxNodeProvisionTime:     to.Ptr("15m"),
	//			NewPodScaleUpDelay:       to.Ptr("1m"),
	//			ScaleDownDelayAfterAdd:   to.Ptr("15m"),
	//			ScanInterval:             to.Ptr("20s"),
	//			SkipNodesWithSystemPods:  to.Ptr("false"),
	//		},
	//		DiskEncryptionSetID:     to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:               to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy: to.Ptr(true),
	//		EnableRBAC:              to.Ptr(true),
	//		KubernetesVersion:       to.Ptr(""),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4),
	//				to.Ptr(armcontainerservice.IPFamilyIPv6)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count: to.Ptr[int32](2),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//		},
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//			Secret:   to.Ptr("secret"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminPassword: to.Ptr("replacePassword1234$"),
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//	SKU: &armcontainerservice.ManagedClusterSKU{
	//		Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
	//		Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
	//	},
	//}, nil)
	//testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DualStackNetworking.json")
	//res, err = poller.PollUntilDone(ctx, nil)
	//testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DualStackNetworking.json")
	//// Response check
	//exampleRes = armcontainerservice.ManagedCluster{
	//	Name:     to.Ptr("clustername1"),
	//	Type:     to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	//	ID:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	//	Location: to.Ptr("location1"),
	//	Tags: map[string]*string{
	//		"archv2": to.Ptr(""),
	//		"tier":   to.Ptr("production"),
	//	},
	//	Identity: &armcontainerservice.ManagedClusterIdentity{
	//		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	//		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//			"/subscriptions/subid1/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	//				ClientID:    to.Ptr("clientId1"),
	//				PrincipalID: to.Ptr("principalId1"),
	//			},
	//		},
	//	},
	//	Properties: &armcontainerservice.ManagedClusterProperties{
	//		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	//			{
	//				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	//				AvailabilityZones: []*string{
	//					to.Ptr("1"),
	//					to.Ptr("2"),
	//					to.Ptr("3")},
	//				Count:                      to.Ptr[int32](3),
	//				CurrentOrchestratorVersion: to.Ptr("1.22.1"),
	//				EnableNodePublicIP:         to.Ptr(true),
	//				MaxPods:                    to.Ptr[int32](110),
	//				Mode:                       to.Ptr(armcontainerservice.AgentPoolModeSystem),
	//				NodeImageVersion:           to.Ptr("AKSUbuntu:1604:2020.03.11"),
	//				OrchestratorVersion:        to.Ptr("1.22.1"),
	//				OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
	//				ProvisioningState:          to.Ptr("Succeeded"),
	//				ScaleDownMode:              to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
	//				VMSize:                     to.Ptr("Standard_DS1_v2"),
	//				Name:                       to.Ptr("nodepool1"),
	//			}},
	//		AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	//			BalanceSimilarNodeGroups: to.Ptr("true"),
	//			Expander:                 to.Ptr(armcontainerservice.ExpanderPriority),
	//			MaxNodeProvisionTime:     to.Ptr("15m"),
	//			NewPodScaleUpDelay:       to.Ptr("1m"),
	//			ScaleDownDelayAfterAdd:   to.Ptr("15m"),
	//			ScanInterval:             to.Ptr("20s"),
	//			SkipNodesWithSystemPods:  to.Ptr("false"),
	//		},
	//		CurrentKubernetesVersion: to.Ptr("1.22.1"),
	//		DiskEncryptionSetID:      to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	//		DNSPrefix:                to.Ptr("dnsprefix1"),
	//		EnablePodSecurityPolicy:  to.Ptr(true),
	//		EnableRBAC:               to.Ptr(true),
	//		Fqdn:                     to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	//		KubernetesVersion:        to.Ptr("1.22.1"),
	//		LinuxProfile: &armcontainerservice.LinuxProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//			SSH: &armcontainerservice.SSHConfiguration{
	//				PublicKeys: []*armcontainerservice.SSHPublicKey{
	//					{
	//						KeyData: to.Ptr("keydata"),
	//					}},
	//			},
	//		},
	//		MaxAgentPools: to.Ptr[int32](1),
	//		NetworkProfile: &armcontainerservice.NetworkProfile{
	//			DNSServiceIP: to.Ptr("10.0.0.10"),
	//			IPFamilies: []*armcontainerservice.IPFamily{
	//				to.Ptr(armcontainerservice.IPFamilyIPv4),
	//				to.Ptr(armcontainerservice.IPFamilyIPv6)},
	//			LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	//				AllocatedOutboundPorts: to.Ptr[int32](2000),
	//				EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	//					},
	//					{
	//						ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip3-ipv6"),
	//					}},
	//				IdleTimeoutInMinutes: to.Ptr[int32](10),
	//				ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	//					Count:     to.Ptr[int32](2),
	//					CountIPv6: to.Ptr[int32](1),
	//				},
	//			},
	//			LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	//			NetworkPlugin:   to.Ptr(armcontainerservice.NetworkPluginKubenet),
	//			OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	//			PodCidr:         to.Ptr("10.244.0.0/16"),
	//			PodCidrs: []*string{
	//				to.Ptr("10.244.0.0/16"),
	//				to.Ptr("fd11:1234::/64")},
	//			ServiceCidr: to.Ptr("10.0.0.0/16"),
	//			ServiceCidrs: []*string{
	//				to.Ptr("10.0.0.0/16"),
	//				to.Ptr("fd00:1234::/108")},
	//		},
	//		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	//		ProvisioningState: to.Ptr("Succeeded"),
	//		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	//			ClientID: to.Ptr("clientid"),
	//		},
	//		WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	//			AdminUsername: to.Ptr("azureuser"),
	//		},
	//	},
	//}
	//if !reflect.DeepEqual(exampleRes, res.ManagedCluster) {
	//	exampleResJson, _ := json.Marshal(exampleRes)
	//	mockResJson, _ := json.Marshal(res.ManagedCluster)
	//	testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_DualStackNetworking.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	//}
}

func (testsuite *FakeTestSuite) TestManagedClusters_Delete() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Delete Managed Cluster"},
	})
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&testsuite.managedClustersClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	poller, err := client.BeginDelete(ctx, "rg1", "clustername1", nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json")
}

func (testsuite *FakeTestSuite) TestPrivateEndpointConnections_List() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsList.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"List Private Endpoint Connections by Managed Cluster"},
	})
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&testsuite.privateEndpointConnectionsClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	res, err := client.List(ctx, "rg1", "clustername1", nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsList.json")
	// Response check
	exampleRes := armcontainerservice.PrivateEndpointConnectionListResult{
		Value: []*armcontainerservice.PrivateEndpointConnection{
			{
				Name: to.Ptr("privateendpointconnection1"),
				Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
				ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
				Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
					PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
						ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
					},
					PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
						Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
					},
					ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
				},
			},
		},
	}
	if !reflect.DeepEqual(exampleRes, res.PrivateEndpointConnectionListResult) {
		exampleResJson, _ := json.Marshal(exampleRes)
		mockResJson, _ := json.Marshal(res.PrivateEndpointConnectionListResult)
		testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsList.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	}
}

func (testsuite *FakeTestSuite) TestPrivateEndpointConnections_Get() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsGet.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Get Private Endpoint Connection"},
	})
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&testsuite.privateEndpointConnectionsClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	res, err := client.Get(ctx, "rg1", "clustername1", "privateendpointconnection1", nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsGet.json")
	// Response check
	exampleRes := armcontainerservice.PrivateEndpointConnection{
		Name: to.Ptr("privateendpointconnection1"),
		Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
		ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
			},
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
			ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
		},
	}
	if !reflect.DeepEqual(exampleRes, res.PrivateEndpointConnection) {
		exampleResJson, _ := json.Marshal(exampleRes)
		mockResJson, _ := json.Marshal(res.PrivateEndpointConnection)
		testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsGet.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	}
}

func (testsuite *FakeTestSuite) TestPrivateEndpointConnections_Update() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsUpdate.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Update Private Endpoint Connection"},
	})
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&testsuite.privateEndpointConnectionsClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	res, err := client.Update(ctx, "rg1", "clustername1", "privateendpointconnection1", armcontainerservice.PrivateEndpointConnection{
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
		},
	}, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsUpdate.json")
	// Response check
	exampleRes := armcontainerservice.PrivateEndpointConnection{
		Name: to.Ptr("privateendpointconnection1"),
		Type: to.Ptr("Microsoft.Network/privateLinkServices/privateEndpointConnections"),
		ID:   to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/managedCluster/clustername1/privateEndpointConnections/privateendpointconnection1"),
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armcontainerservice.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/pe2"),
			},
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
			ProvisioningState: to.Ptr(armcontainerservice.PrivateEndpointConnectionProvisioningStateSucceeded),
		},
	}
	if !reflect.DeepEqual(exampleRes, res.PrivateEndpointConnection) {
		exampleResJson, _ := json.Marshal(exampleRes)
		mockResJson, _ := json.Marshal(res.PrivateEndpointConnection)
		testsuite.Failf("Failed to validate response", "Mock response is not equal to example response for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsUpdate.json:\nmock response: %s\nexample response: %s", mockResJson, exampleResJson)
	}
}

func (testsuite *FakeTestSuite) TestPrivateEndpointConnections_Delete() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Delete Private Endpoint Connection"},
	})
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&testsuite.privateEndpointConnectionsClientServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")
	poller, err := client.BeginDelete(ctx, "rg1", "clustername1", "privateendpointconnection1", nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json")
}
