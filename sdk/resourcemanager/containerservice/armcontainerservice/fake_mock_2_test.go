package armcontainerservice_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type FakeTestSuite2 struct {
	suite.Suite

	cred    azcore.TokenCredential
	options arm.ClientOptions
}

func (testsuite *FakeTestSuite2) SetupSuite() {
	testsuite.cred = &testutil.FakeCredential{}
}

func TestFakeTest2(t *testing.T) {
	suite.Run(t, new(FakeTestSuite2))
}

func (testsuite *FakeTestSuite2) TestManagedClusters_List() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"List Managed Clusters"},
	})

	fakeServer := fake.ManagedClustersServer{}
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

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

	fakeServer.NewListPager = func(options *armcontainerservice.ManagedClustersClientListOptions) (resp azfake.PagerResponder[armcontainerservice.ManagedClustersClientListResponse]) {
		resp = azfake.PagerResponder[armcontainerservice.ManagedClustersClientListResponse]{}
		resp.AddPage(http.StatusOK, armcontainerservice.ManagedClustersClientListResponse{ManagedClusterListResult: pagerExampleRes}, nil)
		return
	}

	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		testsuite.Require().NoError(err, "Failed to advance page for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersList.json")
		testsuite.Require().True(reflect.DeepEqual(pagerExampleRes, nextResult.ManagedClusterListResult))
	}
}

func (testsuite *FakeTestSuite2) TestManagedClusters_Get() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersGet.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Get Managed Cluster"},
	})

	fakerServer := fake.ManagedClustersServer{}

	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&fakerServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
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

	fakerServer.Get = func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientGetOptions) (resp azfake.Responder[armcontainerservice.ManagedClustersClientGetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(resourceGroupName, exampleResourceGroupName)
		testsuite.Require().Equal(resourceName, exampleResourceName)
		resp = azfake.Responder[armcontainerservice.ManagedClustersClientGetResponse]{}
		resp.SetResponse(http.StatusOK, armcontainerservice.ManagedClustersClientGetResponse{ManagedCluster: exampleRes}, nil)
		return
	}

	res, err := client.Get(ctx, exampleResourceGroupName, exampleResourceName, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersGet.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.ManagedCluster))
}

func (testsuite *FakeTestSuite2) TestManagedClusters_CreateOrUpdate() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Create Managed Cluster using an agent pool snapshot"},
	})

	fakeServer := fake.ManagedClustersServer{}
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
	exampleParameters := armcontainerservice.ManagedCluster{
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
	}
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

	fakeServer.BeginCreateOrUpdate = func(ctx context.Context, resourceGroupName string, resourceName string, parameters armcontainerservice.ManagedCluster, options *armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		testsuite.Require().True(reflect.DeepEqual(exampleParameters, parameters))
		resp = azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse]{}
		resp.SetTerminalResponse(http.StatusOK, armcontainerservice.ManagedClustersClientCreateOrUpdateResponse{ManagedCluster: exampleRes}, nil)
		return
	}

	poller, err := client.BeginCreateOrUpdate(ctx, exampleResourceGroupName, exampleResourceName, exampleParameters, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json")
	res, err := poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_Snapshot.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.ManagedCluster))

	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Create Managed Cluster with AKS-managed NAT gateway as outbound type"},
	})

	//client, err = armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
	//	ClientOptions: azcore.ClientOptions{
	//		Transport: fake.NewManagedClustersServerTransport(&fakeServer),
	//	},
	//})
	//testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName = "rg1"
	exampleResourceName = "clustername1"
	exampleParameters = armcontainerservice.ManagedCluster{
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count:              to.Ptr[int32](3),
					EnableNodePublicIP: to.Ptr(false),
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
			EnablePodSecurityPolicy: to.Ptr(true),
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
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
				NatGatewayProfile: &armcontainerservice.ManagedClusterNATGatewayProfile{
					ManagedOutboundIPProfile: &armcontainerservice.ManagedClusterManagedOutboundIPProfile{
						Count: to.Ptr[int32](2),
					},
				},
				OutboundType: to.Ptr(armcontainerservice.OutboundTypeManagedNATGateway),
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
	}
	exampleRes = armcontainerservice.ManagedCluster{
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
					Type:                       to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count:                      to.Ptr[int32](3),
					CurrentOrchestratorVersion: to.Ptr("1.9.6"),
					EnableNodePublicIP:         to.Ptr(false),
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
			EnablePodSecurityPolicy:  to.Ptr(true),
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
				DNSServiceIP:    to.Ptr("10.0.0.10"),
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
				NatGatewayProfile: &armcontainerservice.ManagedClusterNATGatewayProfile{
					EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
						},
						{
							ID: to.Ptr("/subscriptions/subid1/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
						}},
					IdleTimeoutInMinutes: to.Ptr[int32](4),
					ManagedOutboundIPProfile: &armcontainerservice.ManagedClusterManagedOutboundIPProfile{
						Count: to.Ptr[int32](2),
					},
				},
				NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
				OutboundType:  to.Ptr(armcontainerservice.OutboundTypeManagedNATGateway),
				PodCidr:       to.Ptr("10.244.0.0/16"),
				ServiceCidr:   to.Ptr("10.0.0.0/16"),
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
	fakeServer.BeginCreateOrUpdate = func(ctx context.Context, resourceGroupName string, resourceName string, parameters armcontainerservice.ManagedCluster, options *armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		testsuite.Require().True(reflect.DeepEqual(exampleParameters, parameters))
		resp = azfake.PollerResponder[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse]{}
		resp.SetTerminalResponse(http.StatusOK, armcontainerservice.ManagedClustersClientCreateOrUpdateResponse{ManagedCluster: exampleRes}, nil)
		return
	}
	poller, err = client.BeginCreateOrUpdate(ctx, exampleResourceGroupName, exampleResourceName, exampleParameters, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json")
	res, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersCreate_ManagedNATGateway.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.ManagedCluster))
}

func (testsuite *FakeTestSuite2) TestManagedClusters_Delete() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Delete Managed Cluster"},
	})

	fakeServer := fake.ManagedClustersServer{}
	client, err := armcontainerservice.NewManagedClustersClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewManagedClustersServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
	fakeServer.BeginDelete = func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.ManagedClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservice.ManagedClustersClientDeleteResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		resp = azfake.PollerResponder[armcontainerservice.ManagedClustersClientDeleteResponse]{}
		// unexpected status code 200. acceptable values are http.StatusAccepted, http.StatusNoContent
		resp.SetTerminalResponse(http.StatusAccepted, armcontainerservice.ManagedClustersClientDeleteResponse{}, nil)
		return
	}

	poller, err := client.BeginDelete(ctx, exampleResourceGroupName, exampleResourceName, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersDelete.json")
}

func (testsuite *FakeTestSuite2) TestPrivateEndpointConnections_List() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsList.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"List Private Endpoint Connections by Managed Cluster"},
	})

	fakeServer := fake.PrivateEndpointConnectionsServer{}
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
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

	fakeServer.List = func(ctx context.Context, resourceGroupName string, resourceName string, options *armcontainerservice.PrivateEndpointConnectionsClientListOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientListResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		resp = azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientListResponse]{}
		resp.SetResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientListResponse{PrivateEndpointConnectionListResult: exampleRes}, nil)
		return
	}
	res, err := client.List(ctx, exampleResourceGroupName, exampleResourceName, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsList.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.PrivateEndpointConnectionListResult))
}

func (testsuite *FakeTestSuite2) TestPrivateEndpointConnections_Get() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsGet.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Get Private Endpoint Connection"},
	})

	fakeServer := fake.PrivateEndpointConnectionsServer{}
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
	examplePrivateEndpointConnectionName := "privateendpointconnection1"
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

	fakeServer.Get = func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armcontainerservice.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		testsuite.Require().Equal(examplePrivateEndpointConnectionName, privateEndpointConnectionName)
		resp = azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientGetResponse]{}
		resp.SetResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientGetResponse{PrivateEndpointConnection: exampleRes}, nil)
		return
	}

	res, err := client.Get(ctx, exampleResourceGroupName, exampleResourceName, examplePrivateEndpointConnectionName, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsGet.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.PrivateEndpointConnection))
}

func (testsuite *FakeTestSuite2) TestPrivateEndpointConnections_Update() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsUpdate.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Update Private Endpoint Connection"},
	})

	fakeServer := fake.PrivateEndpointConnectionsServer{}
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
	examplePrivateEndpointConnectionName := "privateendpointconnection1"
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
	exampleParameters := armcontainerservice.PrivateEndpointConnection{
		Properties: &armcontainerservice.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armcontainerservice.PrivateLinkServiceConnectionState{
				Status: to.Ptr(armcontainerservice.ConnectionStatusApproved),
			},
		},
	}

	fakeServer.Update = func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, parameters armcontainerservice.PrivateEndpointConnection, options *armcontainerservice.PrivateEndpointConnectionsClientUpdateOptions) (resp azfake.Responder[armcontainerservice.PrivateEndpointConnectionsClientUpdateResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		testsuite.Require().Equal(examplePrivateEndpointConnectionName, privateEndpointConnectionName)
		testsuite.Require().True(reflect.DeepEqual(exampleParameters, parameters))
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
	}

	res, err := client.Update(ctx, exampleResourceGroupName, exampleResourceName, examplePrivateEndpointConnectionName, exampleParameters, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsUpdate.json")
	testsuite.Require().True(reflect.DeepEqual(exampleRes, res.PrivateEndpointConnection))
}

func (testsuite *FakeTestSuite2) TestPrivateEndpointConnections_Delete() {
	ctx := context.Background()
	// From example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json
	ctx = runtime.WithHTTPHeader(ctx, map[string][]string{
		"example-id": {"Delete Private Endpoint Connection"},
	})

	fakeServer := fake.PrivateEndpointConnectionsServer{}
	client, err := armcontainerservice.NewPrivateEndpointConnectionsClient("subid1", testsuite.cred, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewPrivateEndpointConnectionsServerTransport(&fakeServer),
		},
	})
	testsuite.Require().NoError(err, "Failed to create client")

	exampleResourceGroupName := "rg1"
	exampleResourceName := "clustername1"
	examplePrivateEndpointConnectionName := "privateendpointconnection1"

	fakeServer.BeginDelete = func(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *armcontainerservice.PrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder) {
		testsuite.Require().Equal(exampleResourceGroupName, resourceGroupName)
		testsuite.Require().Equal(exampleResourceName, resourceName)
		testsuite.Require().Equal(examplePrivateEndpointConnectionName, privateEndpointConnectionName)
		resp = azfake.PollerResponder[armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse]{}
		//resp.AddNonTerminalResponse(http.StatusOK, nil)
		//resp.AddNonTerminalResponse(http.StatusOK, nil)
		resp.SetTerminalResponse(http.StatusOK, armcontainerservice.PrivateEndpointConnectionsClientDeleteResponse{}, nil)
		return
	}

	poller, err := client.BeginDelete(ctx, exampleResourceGroupName, exampleResourceName, examplePrivateEndpointConnectionName, nil)
	testsuite.Require().NoError(err, "Failed to get result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json")
	_, err = poller.PollUntilDone(ctx, nil)
	testsuite.Require().NoError(err, "Failed to get LRO result for example specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/PrivateEndpointConnectionsDelete.json")
}
