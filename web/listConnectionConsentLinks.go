// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the consent links of a connection
//
// Uses Azure REST API version 2016-06-01.
//
// Other available API versions: 2015-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListConnectionConsentLinks(ctx *pulumi.Context, args *ListConnectionConsentLinksArgs, opts ...pulumi.InvokeOption) (*ListConnectionConsentLinksResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListConnectionConsentLinksResult
	err := ctx.Invoke("azure-native:web:listConnectionConsentLinks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListConnectionConsentLinksArgs struct {
	// Connection name
	ConnectionName string `pulumi:"connectionName"`
	// Collection of resources
	Parameters []ConsentLinkParameterDefinition `pulumi:"parameters"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// Collection of consent links
type ListConnectionConsentLinksResult struct {
	// Collection of resources
	Value []ConsentLinkDefinitionResponse `pulumi:"value"`
}

func ListConnectionConsentLinksOutput(ctx *pulumi.Context, args ListConnectionConsentLinksOutputArgs, opts ...pulumi.InvokeOption) ListConnectionConsentLinksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListConnectionConsentLinksResultOutput, error) {
			args := v.(ListConnectionConsentLinksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:listConnectionConsentLinks", args, ListConnectionConsentLinksResultOutput{}, options).(ListConnectionConsentLinksResultOutput), nil
		}).(ListConnectionConsentLinksResultOutput)
}

type ListConnectionConsentLinksOutputArgs struct {
	// Connection name
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// Collection of resources
	Parameters ConsentLinkParameterDefinitionArrayInput `pulumi:"parameters"`
	// The resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Subscription Id
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (ListConnectionConsentLinksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionConsentLinksArgs)(nil)).Elem()
}

// Collection of consent links
type ListConnectionConsentLinksResultOutput struct{ *pulumi.OutputState }

func (ListConnectionConsentLinksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListConnectionConsentLinksResult)(nil)).Elem()
}

func (o ListConnectionConsentLinksResultOutput) ToListConnectionConsentLinksResultOutput() ListConnectionConsentLinksResultOutput {
	return o
}

func (o ListConnectionConsentLinksResultOutput) ToListConnectionConsentLinksResultOutputWithContext(ctx context.Context) ListConnectionConsentLinksResultOutput {
	return o
}

// Collection of resources
func (o ListConnectionConsentLinksResultOutput) Value() ConsentLinkDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ListConnectionConsentLinksResult) []ConsentLinkDefinitionResponse { return v.Value }).(ConsentLinkDefinitionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListConnectionConsentLinksResultOutput{})
}
