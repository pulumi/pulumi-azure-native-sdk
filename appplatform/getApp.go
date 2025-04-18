// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an App and its properties.
//
// Uses Azure REST API version 2024-01-01-preview.
//
// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:appplatform:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAppArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
	// Indicates whether sync status
	SyncStatus *string `pulumi:"syncStatus"`
}

// App resource payload
type LookupAppResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The Managed Identity type of the app resource
	Identity *ManagedIdentityPropertiesResponse `pulumi:"identity"`
	// The GEO location of the application, always the same with its parent resource
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the App resource
	Properties AppResourcePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAppResult
func (val *LookupAppResult) Defaults() *LookupAppResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAppResultOutput, error) {
			args := v.(LookupAppArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:appplatform:getApp", args, LookupAppResultOutput{}, options).(LookupAppResultOutput), nil
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Indicates whether sync status
	SyncStatus pulumi.StringPtrInput `pulumi:"syncStatus"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

// App resource payload
type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAppResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource.
func (o LookupAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Managed Identity type of the app resource
func (o LookupAppResultOutput) Identity() ManagedIdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAppResult) *ManagedIdentityPropertiesResponse { return v.Identity }).(ManagedIdentityPropertiesResponsePtrOutput)
}

// The GEO location of the application, always the same with its parent resource
func (o LookupAppResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the App resource
func (o LookupAppResultOutput) Properties() AppResourcePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAppResult) AppResourcePropertiesResponse { return v.Properties }).(AppResourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAppResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAppResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
