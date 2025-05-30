// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperations

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
	case "azure-native:iotoperations:Broker":
		r = &Broker{}
	case "azure-native:iotoperations:BrokerAuthentication":
		r = &BrokerAuthentication{}
	case "azure-native:iotoperations:BrokerAuthorization":
		r = &BrokerAuthorization{}
	case "azure-native:iotoperations:BrokerListener":
		r = &BrokerListener{}
	case "azure-native:iotoperations:Dataflow":
		r = &Dataflow{}
	case "azure-native:iotoperations:DataflowEndpoint":
		r = &DataflowEndpoint{}
	case "azure-native:iotoperations:DataflowProfile":
		r = &DataflowProfile{}
	case "azure-native:iotoperations:Instance":
		r = &Instance{}
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
		"iotoperations",
		&module{version},
	)
}
