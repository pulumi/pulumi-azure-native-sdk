// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
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
	case "azure-native:servicebus/v20240101:DisasterRecoveryConfig":
		r = &DisasterRecoveryConfig{}
	case "azure-native:servicebus/v20240101:MigrationConfig":
		r = &MigrationConfig{}
	case "azure-native:servicebus/v20240101:Namespace":
		r = &Namespace{}
	case "azure-native:servicebus/v20240101:NamespaceAuthorizationRule":
		r = &NamespaceAuthorizationRule{}
	case "azure-native:servicebus/v20240101:NamespaceNetworkRuleSet":
		r = &NamespaceNetworkRuleSet{}
	case "azure-native:servicebus/v20240101:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:servicebus/v20240101:Queue":
		r = &Queue{}
	case "azure-native:servicebus/v20240101:QueueAuthorizationRule":
		r = &QueueAuthorizationRule{}
	case "azure-native:servicebus/v20240101:Rule":
		r = &Rule{}
	case "azure-native:servicebus/v20240101:Subscription":
		r = &Subscription{}
	case "azure-native:servicebus/v20240101:Topic":
		r = &Topic{}
	case "azure-native:servicebus/v20240101:TopicAuthorizationRule":
		r = &TopicAuthorizationRule{}
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
		"servicebus/v20240101",
		&module{version},
	)
}
