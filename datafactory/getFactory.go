// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a factory.
//
// Uses Azure REST API version 2018-06-01.
func LookupFactory(ctx *pulumi.Context, args *LookupFactoryArgs, opts ...pulumi.InvokeOption) (*LookupFactoryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFactoryResult
	err := ctx.Invoke("azure-native:datafactory:getFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFactoryArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Factory resource type.
type LookupFactoryResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Time the factory was created in ISO8601 format.
	CreateTime string `pulumi:"createTime"`
	// Etag identifies change in the resource.
	ETag string `pulumi:"eTag"`
	// Properties to enable Customer Managed Key for the factory.
	Encryption *EncryptionConfigurationResponse `pulumi:"encryption"`
	// List of parameters for factory.
	GlobalParameters map[string]GlobalParameterSpecificationResponse `pulumi:"globalParameters"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Managed service identity of the factory.
	Identity *FactoryIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Factory provisioning state, example Succeeded.
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for the data factory.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purview information of the factory.
	PurviewConfiguration *PurviewConfigurationResponse `pulumi:"purviewConfiguration"`
	// Git repo information of the factory.
	RepoConfiguration interface{} `pulumi:"repoConfiguration"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
	// Version of the factory.
	Version string `pulumi:"version"`
}

func LookupFactoryOutput(ctx *pulumi.Context, args LookupFactoryOutputArgs, opts ...pulumi.InvokeOption) LookupFactoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFactoryResultOutput, error) {
			args := v.(LookupFactoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datafactory:getFactory", args, LookupFactoryResultOutput{}, options).(LookupFactoryResultOutput), nil
		}).(LookupFactoryResultOutput)
}

type LookupFactoryOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFactoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFactoryArgs)(nil)).Elem()
}

// Factory resource type.
type LookupFactoryResultOutput struct{ *pulumi.OutputState }

func (LookupFactoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFactoryResult)(nil)).Elem()
}

func (o LookupFactoryResultOutput) ToLookupFactoryResultOutput() LookupFactoryResultOutput {
	return o
}

func (o LookupFactoryResultOutput) ToLookupFactoryResultOutputWithContext(ctx context.Context) LookupFactoryResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupFactoryResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Time the factory was created in ISO8601 format.
func (o LookupFactoryResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Etag identifies change in the resource.
func (o LookupFactoryResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.ETag }).(pulumi.StringOutput)
}

// Properties to enable Customer Managed Key for the factory.
func (o LookupFactoryResultOutput) Encryption() EncryptionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *EncryptionConfigurationResponse { return v.Encryption }).(EncryptionConfigurationResponsePtrOutput)
}

// List of parameters for factory.
func (o LookupFactoryResultOutput) GlobalParameters() GlobalParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupFactoryResult) map[string]GlobalParameterSpecificationResponse { return v.GlobalParameters }).(GlobalParameterSpecificationResponseMapOutput)
}

// The resource identifier.
func (o LookupFactoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity of the factory.
func (o LookupFactoryResultOutput) Identity() FactoryIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *FactoryIdentityResponse { return v.Identity }).(FactoryIdentityResponsePtrOutput)
}

// The resource location.
func (o LookupFactoryResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupFactoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Factory provisioning state, example Succeeded.
func (o LookupFactoryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for the data factory.
func (o LookupFactoryResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purview information of the factory.
func (o LookupFactoryResultOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *PurviewConfigurationResponse { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

// Git repo information of the factory.
func (o LookupFactoryResultOutput) RepoConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupFactoryResult) interface{} { return v.RepoConfiguration }).(pulumi.AnyOutput)
}

// The resource tags.
func (o LookupFactoryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFactoryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupFactoryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version of the factory.
func (o LookupFactoryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFactoryResultOutput{})
}
