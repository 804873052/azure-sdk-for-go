//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

import "encoding/json"

func unmarshalActionRulePropertiesClassification(rawMsg json.RawMessage) (ActionRulePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ActionRulePropertiesClassification
	switch m["type"] {
	case string(ActionRuleTypeActionGroup):
		b = &ActionGroup{}
	case string(ActionRuleTypeDiagnostics):
		b = &Diagnostics{}
	case string(ActionRuleTypeSuppression):
		b = &Suppression{}
	default:
		b = &ActionRuleProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalActionRulePropertiesClassificationArray(rawMsg json.RawMessage) ([]ActionRulePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ActionRulePropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalActionRulePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAlertsMetaDataPropertiesClassification(rawMsg json.RawMessage) (AlertsMetaDataPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AlertsMetaDataPropertiesClassification
	switch m["metadataIdentifier"] {
	case string(MetadataIdentifierMonitorServiceList):
		b = &MonitorServiceList{}
	default:
		b = &AlertsMetaDataProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalAlertsMetaDataPropertiesClassificationArray(rawMsg json.RawMessage) ([]AlertsMetaDataPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AlertsMetaDataPropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAlertsMetaDataPropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
