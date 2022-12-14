// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deployment script object.
//
// Deprecated: Please use one of the variants: AzureCliScript, AzurePowerShellScript.
func LookupDeploymentScript(ctx *pulumi.Context, args *LookupDeploymentScriptArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentScriptResult, error) {
	var rv LookupDeploymentScriptResult
	err := ctx.Invoke("azure-native:resources/v20201001:getDeploymentScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentScriptArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName string `pulumi:"scriptName"`
}

// Deployment script object.
type LookupDeploymentScriptResult struct {
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Type of the script.
	Kind string `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location string `pulumi:"location"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupDeploymentScriptOutput(ctx *pulumi.Context, args LookupDeploymentScriptOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentScriptResult, error) {
			args := v.(LookupDeploymentScriptArgs)
			r, err := LookupDeploymentScript(ctx, &args, opts...)
			var s LookupDeploymentScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentScriptResultOutput)
}

type LookupDeploymentScriptOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupDeploymentScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentScriptArgs)(nil)).Elem()
}

// Deployment script object.
type LookupDeploymentScriptResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentScriptResult)(nil)).Elem()
}

func (o LookupDeploymentScriptResultOutput) ToLookupDeploymentScriptResultOutput() LookupDeploymentScriptResultOutput {
	return o
}

func (o LookupDeploymentScriptResultOutput) ToLookupDeploymentScriptResultOutputWithContext(ctx context.Context) LookupDeploymentScriptResultOutput {
	return o
}

// String Id used to locate any resource on Azure.
func (o LookupDeploymentScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
func (o LookupDeploymentScriptResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Type of the script.
func (o LookupDeploymentScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The location of the ACI and the storage account for the deployment script.
func (o LookupDeploymentScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupDeploymentScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupDeploymentScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDeploymentScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o LookupDeploymentScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeploymentScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentScriptResultOutput{})
}
