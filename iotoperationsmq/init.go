// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsmq

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:iotoperationsmq:Broker":
		r = &Broker{}
	case "azure-native:iotoperationsmq:BrokerAuthentication":
		r = &BrokerAuthentication{}
	case "azure-native:iotoperationsmq:BrokerAuthorization":
		r = &BrokerAuthorization{}
	case "azure-native:iotoperationsmq:BrokerListener":
		r = &BrokerListener{}
	case "azure-native:iotoperationsmq:DataLakeConnector":
		r = &DataLakeConnector{}
	case "azure-native:iotoperationsmq:DataLakeConnectorTopicMap":
		r = &DataLakeConnectorTopicMap{}
	case "azure-native:iotoperationsmq:DiagnosticService":
		r = &DiagnosticService{}
	case "azure-native:iotoperationsmq:KafkaConnector":
		r = &KafkaConnector{}
	case "azure-native:iotoperationsmq:KafkaConnectorTopicMap":
		r = &KafkaConnectorTopicMap{}
	case "azure-native:iotoperationsmq:Mq":
		r = &Mq{}
	case "azure-native:iotoperationsmq:MqttBridgeConnector":
		r = &MqttBridgeConnector{}
	case "azure-native:iotoperationsmq:MqttBridgeTopicMap":
		r = &MqttBridgeTopicMap{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"iotoperationsmq",
		&module{version},
	)
}
