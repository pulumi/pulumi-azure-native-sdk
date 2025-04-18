// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a SQL Server registration.
//
// Uses Azure REST API version 2019-07-24-preview.
func LookupSqlServerRegistration(ctx *pulumi.Context, args *LookupSqlServerRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupSqlServerRegistrationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlServerRegistrationResult
	err := ctx.Invoke("azure-native:azuredata:getSqlServerRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlServerRegistrationArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
}

// A SQL server registration.
type LookupSqlServerRegistrationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Optional Properties as JSON string
	PropertyBag *string `pulumi:"propertyBag"`
	// Resource Group Name
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Subscription Id
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupSqlServerRegistrationOutput(ctx *pulumi.Context, args LookupSqlServerRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupSqlServerRegistrationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSqlServerRegistrationResultOutput, error) {
			args := v.(LookupSqlServerRegistrationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azuredata:getSqlServerRegistration", args, LookupSqlServerRegistrationResultOutput{}, options).(LookupSqlServerRegistrationResultOutput), nil
		}).(LookupSqlServerRegistrationResultOutput)
}

type LookupSqlServerRegistrationOutputArgs struct {
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName pulumi.StringInput `pulumi:"sqlServerRegistrationName"`
}

func (LookupSqlServerRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationArgs)(nil)).Elem()
}

// A SQL server registration.
type LookupSqlServerRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupSqlServerRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlServerRegistrationResult)(nil)).Elem()
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutput() LookupSqlServerRegistrationResultOutput {
	return o
}

func (o LookupSqlServerRegistrationResultOutput) ToLookupSqlServerRegistrationResultOutputWithContext(ctx context.Context) LookupSqlServerRegistrationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSqlServerRegistrationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlServerRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSqlServerRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlServerRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional Properties as JSON string
func (o LookupSqlServerRegistrationResultOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

// Resource Group Name
func (o LookupSqlServerRegistrationResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Subscription Id
func (o LookupSqlServerRegistrationResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Read only system data
func (o LookupSqlServerRegistrationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSqlServerRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupSqlServerRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlServerRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlServerRegistrationResultOutput{})
}
