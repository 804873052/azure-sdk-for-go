//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

import "encoding/json"

func unmarshalFileShareConfigurationClassification(rawMsg json.RawMessage) (FileShareConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FileShareConfigurationClassification
	switch m["configurationType"] {
	case string(ConfigurationTypeCreateAndMount):
		b = &CreateAndMountFileShareConfiguration{}
	case string(ConfigurationTypeMount):
		b = &MountFileShareConfiguration{}
	case string(ConfigurationTypeSkip):
		b = &SkipFileShareConfiguration{}
	default:
		b = &FileShareConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalInfrastructureConfigurationClassification(rawMsg json.RawMessage) (InfrastructureConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b InfrastructureConfigurationClassification
	switch m["deploymentType"] {
	case string(SAPDeploymentTypeSingleServer):
		b = &SingleServerConfiguration{}
	case string(SAPDeploymentTypeThreeTier):
		b = &ThreeTierConfiguration{}
	default:
		b = &InfrastructureConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalOSConfigurationClassification(rawMsg json.RawMessage) (OSConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OSConfigurationClassification
	switch m["osType"] {
	case string(OSTypeLinux):
		b = &LinuxConfiguration{}
	case string(OSTypeWindows):
		b = &WindowsConfiguration{}
	default:
		b = &OSConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalProviderSpecificPropertiesClassification(rawMsg json.RawMessage) (ProviderSpecificPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ProviderSpecificPropertiesClassification
	switch m["providerType"] {
	case "Db2":
		b = &DB2ProviderInstanceProperties{}
	case "MsSqlServer":
		b = &MsSQLServerProviderInstanceProperties{}
	case "PrometheusHaCluster":
		b = &PrometheusHaClusterProviderInstanceProperties{}
	case "PrometheusOS":
		b = &PrometheusOSProviderInstanceProperties{}
	case "SapHana":
		b = &HanaDbProviderInstanceProperties{}
	case "SapNetWeaver":
		b = &SapNetWeaverProviderInstanceProperties{}
	default:
		b = &ProviderSpecificProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSAPConfigurationClassification(rawMsg json.RawMessage) (SAPConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SAPConfigurationClassification
	switch m["configurationType"] {
	case string(SAPConfigurationTypeDeployment):
		b = &DeploymentConfiguration{}
	case string(SAPConfigurationTypeDeploymentWithOSConfig):
		b = &DeploymentWithOSConfiguration{}
	case string(SAPConfigurationTypeDiscovery):
		b = &DiscoveryConfiguration{}
	default:
		b = &SAPConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSAPSizingRecommendationResultClassification(rawMsg json.RawMessage) (SAPSizingRecommendationResultClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SAPSizingRecommendationResultClassification
	switch m["deploymentType"] {
	case string(SAPDeploymentTypeSingleServer):
		b = &SingleServerRecommendationResult{}
	case string(SAPDeploymentTypeThreeTier):
		b = &ThreeTierRecommendationResult{}
	default:
		b = &SAPSizingRecommendationResult{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSingleServerCustomResourceNamesClassification(rawMsg json.RawMessage) (SingleServerCustomResourceNamesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SingleServerCustomResourceNamesClassification
	switch m["namingPatternType"] {
	case string(NamingPatternTypeFullResourceName):
		b = &SingleServerFullResourceNames{}
	default:
		b = &SingleServerCustomResourceNames{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalSoftwareConfigurationClassification(rawMsg json.RawMessage) (SoftwareConfigurationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SoftwareConfigurationClassification
	switch m["softwareInstallationType"] {
	case string(SAPSoftwareInstallationTypeExternal):
		b = &ExternalInstallationSoftwareConfiguration{}
	case string(SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig):
		b = &SAPInstallWithoutOSConfigSoftwareConfiguration{}
	case string(SAPSoftwareInstallationTypeServiceInitiated):
		b = &ServiceInitiatedSoftwareConfiguration{}
	default:
		b = &SoftwareConfiguration{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalThreeTierCustomResourceNamesClassification(rawMsg json.RawMessage) (ThreeTierCustomResourceNamesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ThreeTierCustomResourceNamesClassification
	switch m["namingPatternType"] {
	case string(NamingPatternTypeFullResourceName):
		b = &ThreeTierFullResourceNames{}
	default:
		b = &ThreeTierCustomResourceNames{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
