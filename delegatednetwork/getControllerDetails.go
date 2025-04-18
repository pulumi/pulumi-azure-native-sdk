// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package delegatednetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details about the specified dnc controller.
//
// Uses Azure REST API version 2023-06-27-preview.
//
// Other available API versions: 2021-03-15, 2023-05-18-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native delegatednetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupControllerDetails(ctx *pulumi.Context, args *LookupControllerDetailsArgs, opts ...pulumi.InvokeOption) (*LookupControllerDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupControllerDetailsResult
	err := ctx.Invoke("azure-native:delegatednetwork:getControllerDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupControllerDetailsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName string `pulumi:"resourceName"`
}

// Represents an instance of a DNC controller.
type LookupControllerDetailsResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// dnc application id should be used by customer to authenticate with dnc gateway.
	DncAppId string `pulumi:"dncAppId"`
	// dnc endpoint url that customers can use to connect to
	DncEndpoint string `pulumi:"dncEndpoint"`
	// tenant id of dnc application id
	DncTenantId string `pulumi:"dncTenantId"`
	// An identifier that represents the resource.
	Id string `pulumi:"id"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The current state of dnc controller resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The purpose of the dnc controller resource.
	Purpose *string `pulumi:"purpose"`
	// Resource guid.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupControllerDetailsResult
func (val *LookupControllerDetailsResult) Defaults() *LookupControllerDetailsResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Purpose == nil {
		purpose_ := "prod"
		tmp.Purpose = &purpose_
	}
	return &tmp
}
func LookupControllerDetailsOutput(ctx *pulumi.Context, args LookupControllerDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupControllerDetailsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupControllerDetailsResultOutput, error) {
			args := v.(LookupControllerDetailsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:delegatednetwork:getControllerDetails", args, LookupControllerDetailsResultOutput{}, options).(LookupControllerDetailsResultOutput), nil
		}).(LookupControllerDetailsResultOutput)
}

type LookupControllerDetailsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupControllerDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerDetailsArgs)(nil)).Elem()
}

// Represents an instance of a DNC controller.
type LookupControllerDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupControllerDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerDetailsResult)(nil)).Elem()
}

func (o LookupControllerDetailsResultOutput) ToLookupControllerDetailsResultOutput() LookupControllerDetailsResultOutput {
	return o
}

func (o LookupControllerDetailsResultOutput) ToLookupControllerDetailsResultOutputWithContext(ctx context.Context) LookupControllerDetailsResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupControllerDetailsResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// dnc application id should be used by customer to authenticate with dnc gateway.
func (o LookupControllerDetailsResultOutput) DncAppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncAppId }).(pulumi.StringOutput)
}

// dnc endpoint url that customers can use to connect to
func (o LookupControllerDetailsResultOutput) DncEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncEndpoint }).(pulumi.StringOutput)
}

// tenant id of dnc application id
func (o LookupControllerDetailsResultOutput) DncTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.DncTenantId }).(pulumi.StringOutput)
}

// An identifier that represents the resource.
func (o LookupControllerDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource.
func (o LookupControllerDetailsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupControllerDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of dnc controller resource.
func (o LookupControllerDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The purpose of the dnc controller resource.
func (o LookupControllerDetailsResultOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

// Resource guid.
func (o LookupControllerDetailsResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The resource tags.
func (o LookupControllerDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of resource.
func (o LookupControllerDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControllerDetailsResultOutput{})
}
