// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a Traffic Manager Real User Metrics key response.
//
// Deprecated: Version 2017-09-01-preview will be removed in v2 of the provider.
func LookupTrafficManagerUserMetricsKey(ctx *pulumi.Context, args *LookupTrafficManagerUserMetricsKeyArgs, opts ...pulumi.InvokeOption) (*LookupTrafficManagerUserMetricsKeyResult, error) {
	var rv LookupTrafficManagerUserMetricsKeyResult
	err := ctx.Invoke("azure-native:network/v20170901preview:getTrafficManagerUserMetricsKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficManagerUserMetricsKeyArgs struct {
}

// Class representing a Traffic Manager Real User Metrics key response.
type LookupTrafficManagerUserMetricsKeyResult struct {
	Id   string  `pulumi:"id"`
	Key  *string `pulumi:"key"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}