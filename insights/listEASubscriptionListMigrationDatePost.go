// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Subscription migrate date information properties
// API Version: 2017-10-01.
func ListEASubscriptionListMigrationDatePost(ctx *pulumi.Context, args *ListEASubscriptionListMigrationDatePostArgs, opts ...pulumi.InvokeOption) (*ListEASubscriptionListMigrationDatePostResult, error) {
	var rv ListEASubscriptionListMigrationDatePostResult
	err := ctx.Invoke("azure-native:insights:listEASubscriptionListMigrationDatePost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEASubscriptionListMigrationDatePostArgs struct {
}

// Subscription migrate date information properties
type ListEASubscriptionListMigrationDatePostResult struct {
	IsGrandFatherableSubscription *bool   `pulumi:"isGrandFatherableSubscription"`
	OptedInDate                   *string `pulumi:"optedInDate"`
}