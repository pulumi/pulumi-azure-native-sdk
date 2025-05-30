// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network manager security admin configuration rule collection.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2021-02-01-preview, 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-01-01-preview, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAdminRuleCollection(ctx *pulumi.Context, args *LookupAdminRuleCollectionArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAdminRuleCollectionResult
	err := ctx.Invoke("azure-native:network:getAdminRuleCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleCollectionArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
}

// Defines the admin rule collection.
type LookupAdminRuleCollectionResult struct {
	// Groups for configuration
	AppliesToGroups []NetworkManagerSecurityGroupItemResponse `pulumi:"appliesToGroups"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// A description of the admin rule collection.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAdminRuleCollectionOutput(ctx *pulumi.Context, args LookupAdminRuleCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupAdminRuleCollectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAdminRuleCollectionResultOutput, error) {
			args := v.(LookupAdminRuleCollectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getAdminRuleCollection", args, LookupAdminRuleCollectionResultOutput{}, options).(LookupAdminRuleCollectionResultOutput), nil
		}).(LookupAdminRuleCollectionResultOutput)
}

type LookupAdminRuleCollectionOutputArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
}

func (LookupAdminRuleCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionArgs)(nil)).Elem()
}

// Defines the admin rule collection.
type LookupAdminRuleCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupAdminRuleCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleCollectionResult)(nil)).Elem()
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutput() LookupAdminRuleCollectionResultOutput {
	return o
}

func (o LookupAdminRuleCollectionResultOutput) ToLookupAdminRuleCollectionResultOutputWithContext(ctx context.Context) LookupAdminRuleCollectionResultOutput {
	return o
}

// Groups for configuration
func (o LookupAdminRuleCollectionResultOutput) AppliesToGroups() NetworkManagerSecurityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) []NetworkManagerSecurityGroupItemResponse {
		return v.AppliesToGroups
	}).(NetworkManagerSecurityGroupItemResponseArrayOutput)
}

// The Azure API version of the resource.
func (o LookupAdminRuleCollectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// A description of the admin rule collection.
func (o LookupAdminRuleCollectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupAdminRuleCollectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupAdminRuleCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAdminRuleCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupAdminRuleCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupAdminRuleCollectionResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupAdminRuleCollectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupAdminRuleCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdminRuleCollectionResultOutput{})
}
