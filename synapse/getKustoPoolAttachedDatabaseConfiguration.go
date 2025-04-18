// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns an attached database configuration.
//
// Uses Azure REST API version 2021-06-01-preview.
func LookupKustoPoolAttachedDatabaseConfiguration(ctx *pulumi.Context, args *LookupKustoPoolAttachedDatabaseConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupKustoPoolAttachedDatabaseConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoPoolAttachedDatabaseConfigurationResult
	err := ctx.Invoke("azure-native:synapse:getKustoPoolAttachedDatabaseConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoPoolAttachedDatabaseConfigurationArgs struct {
	// The name of the attached database configuration.
	AttachedDatabaseConfigurationName string `pulumi:"attachedDatabaseConfigurationName"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Class representing an attached database configuration.
type LookupKustoPoolAttachedDatabaseConfigurationResult struct {
	// The list of databases from the clusterResourceId which are currently attached to the kusto pool.
	AttachedDatabaseNames []string `pulumi:"attachedDatabaseNames"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName string `pulumi:"databaseName"`
	// The default principals modification kind
	DefaultPrincipalsModificationKind string `pulumi:"defaultPrincipalsModificationKind"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The resource id of the kusto pool where the databases you would like to attach reside.
	KustoPoolResourceId string `pulumi:"kustoPoolResourceId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Table level sharing specifications
	TableLevelSharingProperties *TableLevelSharingPropertiesResponse `pulumi:"tableLevelSharingProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupKustoPoolAttachedDatabaseConfigurationOutput(ctx *pulumi.Context, args LookupKustoPoolAttachedDatabaseConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKustoPoolAttachedDatabaseConfigurationResultOutput, error) {
			args := v.(LookupKustoPoolAttachedDatabaseConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:synapse:getKustoPoolAttachedDatabaseConfiguration", args, LookupKustoPoolAttachedDatabaseConfigurationResultOutput{}, options).(LookupKustoPoolAttachedDatabaseConfigurationResultOutput), nil
		}).(LookupKustoPoolAttachedDatabaseConfigurationResultOutput)
}

type LookupKustoPoolAttachedDatabaseConfigurationOutputArgs struct {
	// The name of the attached database configuration.
	AttachedDatabaseConfigurationName pulumi.StringInput `pulumi:"attachedDatabaseConfigurationName"`
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput `pulumi:"kustoPoolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupKustoPoolAttachedDatabaseConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolAttachedDatabaseConfigurationArgs)(nil)).Elem()
}

// Class representing an attached database configuration.
type LookupKustoPoolAttachedDatabaseConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoPoolAttachedDatabaseConfigurationResult)(nil)).Elem()
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ToLookupKustoPoolAttachedDatabaseConfigurationResultOutput() LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return o
}

func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ToLookupKustoPoolAttachedDatabaseConfigurationResultOutputWithContext(ctx context.Context) LookupKustoPoolAttachedDatabaseConfigurationResultOutput {
	return o
}

// The list of databases from the clusterResourceId which are currently attached to the kusto pool.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) AttachedDatabaseNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) []string { return v.AttachedDatabaseNames }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The default principals modification kind
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) DefaultPrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string {
		return v.DefaultPrincipalsModificationKind
	}).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource id of the kusto pool where the databases you would like to attach reside.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) KustoPoolResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.KustoPoolResourceId }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Table level sharing specifications
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) *TableLevelSharingPropertiesResponse {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKustoPoolAttachedDatabaseConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoPoolAttachedDatabaseConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoPoolAttachedDatabaseConfigurationResultOutput{})
}
