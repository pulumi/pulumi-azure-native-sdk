// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the WordPress instance resource.
func LookupWordpressInstance(ctx *pulumi.Context, args *LookupWordpressInstanceArgs, opts ...pulumi.InvokeOption) (*LookupWordpressInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWordpressInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getWordpressInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWordpressInstanceArgs struct {
	// Php workload name
	PhpWorkloadName string `pulumi:"phpWorkloadName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// WordPress instance resource
type LookupWordpressInstanceResult struct {
	// Database name used by the application
	DatabaseName *string `pulumi:"databaseName"`
	// User name used by the application to connect to database
	DatabaseUser *string `pulumi:"databaseUser"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// WordPress instance provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// Site Url to access the WordPress application
	SiteUrl string `pulumi:"siteUrl"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Application version
	Version string `pulumi:"version"`
}

func LookupWordpressInstanceOutput(ctx *pulumi.Context, args LookupWordpressInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupWordpressInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWordpressInstanceResultOutput, error) {
			args := v.(LookupWordpressInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:workloads/v20211201preview:getWordpressInstance", args, LookupWordpressInstanceResultOutput{}, options).(LookupWordpressInstanceResultOutput), nil
		}).(LookupWordpressInstanceResultOutput)
}

type LookupWordpressInstanceOutputArgs struct {
	// Php workload name
	PhpWorkloadName pulumi.StringInput `pulumi:"phpWorkloadName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWordpressInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWordpressInstanceArgs)(nil)).Elem()
}

// WordPress instance resource
type LookupWordpressInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupWordpressInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWordpressInstanceResult)(nil)).Elem()
}

func (o LookupWordpressInstanceResultOutput) ToLookupWordpressInstanceResultOutput() LookupWordpressInstanceResultOutput {
	return o
}

func (o LookupWordpressInstanceResultOutput) ToLookupWordpressInstanceResultOutputWithContext(ctx context.Context) LookupWordpressInstanceResultOutput {
	return o
}

// Database name used by the application
func (o LookupWordpressInstanceResultOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// User name used by the application to connect to database
func (o LookupWordpressInstanceResultOutput) DatabaseUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) *string { return v.DatabaseUser }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWordpressInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWordpressInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// WordPress instance provisioning state
func (o LookupWordpressInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Site Url to access the WordPress application
func (o LookupWordpressInstanceResultOutput) SiteUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.SiteUrl }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWordpressInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWordpressInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Application version
func (o LookupWordpressInstanceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWordpressInstanceResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWordpressInstanceResultOutput{})
}
