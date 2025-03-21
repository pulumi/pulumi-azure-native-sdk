// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a workspace instance.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The resource group name of the workspace.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Log Analytics Workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The top level Workspace resource container.
type LookupWorkspaceResult struct {
	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId string `pulumi:"customerId"`
	// The ETag of the workspace.
	ETag *string `pulumi:"eTag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// This is a legacy property and is not used anymore. Kept here for backward compatibility.
	PortalUrl string `pulumi:"portalUrl"`
	// The provisioning state of the workspace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// This is a read-only legacy property. It is always set to 'Azure' by the service. Kept here for backward compatibility.
	Source string `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResultOutput, error) {
			args := v.(LookupWorkspaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:operationalinsights/v20151101preview:getWorkspace", args, LookupWorkspaceResultOutput{}, options).(LookupWorkspaceResultOutput), nil
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	// The resource group name of the workspace.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Log Analytics Workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// The top level Workspace resource container.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

// This is a read-only property. Represents the ID associated with the workspace.
func (o LookupWorkspaceResultOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CustomerId }).(pulumi.StringOutput)
}

// The ETag of the workspace.
func (o LookupWorkspaceResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// This is a legacy property and is not used anymore. Kept here for backward compatibility.
func (o LookupWorkspaceResultOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.PortalUrl }).(pulumi.StringOutput)
}

// The provisioning state of the workspace.
func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
func (o LookupWorkspaceResultOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

// The SKU of the workspace.
func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// This is a read-only legacy property. It is always set to 'Azure' by the service. Kept here for backward compatibility.
func (o LookupWorkspaceResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Source }).(pulumi.StringOutput)
}

// Resource tags
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
