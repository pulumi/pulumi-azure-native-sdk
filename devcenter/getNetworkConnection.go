// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network connection resource
//
// Uses Azure REST API version 2024-02-01.
//
// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupNetworkConnection(ctx *pulumi.Context, args *LookupNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkConnectionResult
	err := ctx.Invoke("azure-native:devcenter:getNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkConnectionArgs struct {
	// Name of the Network Connection that can be applied to a Pool.
	NetworkConnectionName string `pulumi:"networkConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network related settings
type LookupNetworkConnectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// AAD Join type.
	DomainJoinType string `pulumi:"domainJoinType"`
	// Active Directory domain name
	DomainName *string `pulumi:"domainName"`
	// The password for the account used to join domain
	DomainPassword *string `pulumi:"domainPassword"`
	// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
	DomainUsername *string `pulumi:"domainUsername"`
	// Overall health status of the network connection. Health checks are run on creation, update, and periodically to validate the network connection.
	HealthCheckStatus string `pulumi:"healthCheckStatus"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The name for resource group where NICs will be placed.
	NetworkingResourceGroupName *string `pulumi:"networkingResourceGroupName"`
	// Active Directory domain Organization Unit (OU)
	OrganizationUnit *string `pulumi:"organizationUnit"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The subnet to attach Virtual Machines to
	SubnetId string `pulumi:"subnetId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNetworkConnectionOutput(ctx *pulumi.Context, args LookupNetworkConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkConnectionResultOutput, error) {
			args := v.(LookupNetworkConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter:getNetworkConnection", args, LookupNetworkConnectionResultOutput{}, options).(LookupNetworkConnectionResultOutput), nil
		}).(LookupNetworkConnectionResultOutput)
}

type LookupNetworkConnectionOutputArgs struct {
	// Name of the Network Connection that can be applied to a Pool.
	NetworkConnectionName pulumi.StringInput `pulumi:"networkConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkConnectionArgs)(nil)).Elem()
}

// Network related settings
type LookupNetworkConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkConnectionResult)(nil)).Elem()
}

func (o LookupNetworkConnectionResultOutput) ToLookupNetworkConnectionResultOutput() LookupNetworkConnectionResultOutput {
	return o
}

func (o LookupNetworkConnectionResultOutput) ToLookupNetworkConnectionResultOutputWithContext(ctx context.Context) LookupNetworkConnectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupNetworkConnectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// AAD Join type.
func (o LookupNetworkConnectionResultOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.DomainJoinType }).(pulumi.StringOutput)
}

// Active Directory domain name
func (o LookupNetworkConnectionResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The password for the account used to join domain
func (o LookupNetworkConnectionResultOutput) DomainPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainPassword }).(pulumi.StringPtrOutput)
}

// The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
func (o LookupNetworkConnectionResultOutput) DomainUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.DomainUsername }).(pulumi.StringPtrOutput)
}

// Overall health status of the network connection. Health checks are run on creation, update, and periodically to validate the network connection.
func (o LookupNetworkConnectionResultOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNetworkConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name for resource group where NICs will be placed.
func (o LookupNetworkConnectionResultOutput) NetworkingResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.NetworkingResourceGroupName }).(pulumi.StringPtrOutput)
}

// Active Directory domain Organization Unit (OU)
func (o LookupNetworkConnectionResultOutput) OrganizationUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) *string { return v.OrganizationUnit }).(pulumi.StringPtrOutput)
}

// The provisioning state of the resource.
func (o LookupNetworkConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The subnet to attach Virtual Machines to
func (o LookupNetworkConnectionResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkConnectionResultOutput{})
}
