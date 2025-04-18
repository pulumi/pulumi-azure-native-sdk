// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysqldiscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the MySQLServers resource.
//
// Uses Azure REST API version 2024-09-30-preview.
func LookupMySQLServer(ctx *pulumi.Context, args *LookupMySQLServerArgs, opts ...pulumi.InvokeOption) (*LookupMySQLServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMySQLServerResult
	err := ctx.Invoke("azure-native:mysqldiscovery:getMySQLServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMySQLServerArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Server
	ServerName string `pulumi:"serverName"`
	// The name of Site
	SiteName string `pulumi:"siteName"`
}

// The MySQLServer resource definition.
type LookupMySQLServerResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// mysql server edition.
	Edition *string `pulumi:"edition"`
	// The list of errors.
	Errors []ErrorResponse `pulumi:"errors"`
	// The Server IP/host name.
	HostIp []string `pulumi:"hostIp"`
	// The Server IP/host name.
	HostName string `pulumi:"hostName"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Resource labels.
	Labels map[string]string `pulumi:"labels"`
	// discovery Machine Id
	MachineId *string `pulumi:"machineId"`
	// The mysql server version.
	MysqlVersion *string `pulumi:"mysqlVersion"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The number of database.
	NumberOfDatabase *float64 `pulumi:"numberOfDatabase"`
	// MySQL Server port number
	PortNumber string `pulumi:"portNumber"`
	// Gets or sets the provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Time when mysql version support end.
	SupportEndIn *string `pulumi:"supportEndIn"`
	// mysql version support status.
	SupportStatus *string `pulumi:"supportStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMySQLServerOutput(ctx *pulumi.Context, args LookupMySQLServerOutputArgs, opts ...pulumi.InvokeOption) LookupMySQLServerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMySQLServerResultOutput, error) {
			args := v.(LookupMySQLServerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:mysqldiscovery:getMySQLServer", args, LookupMySQLServerResultOutput{}, options).(LookupMySQLServerResultOutput), nil
		}).(LookupMySQLServerResultOutput)
}

type LookupMySQLServerOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Server
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of Site
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupMySQLServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMySQLServerArgs)(nil)).Elem()
}

// The MySQLServer resource definition.
type LookupMySQLServerResultOutput struct{ *pulumi.OutputState }

func (LookupMySQLServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMySQLServerResult)(nil)).Elem()
}

func (o LookupMySQLServerResultOutput) ToLookupMySQLServerResultOutput() LookupMySQLServerResultOutput {
	return o
}

func (o LookupMySQLServerResultOutput) ToLookupMySQLServerResultOutputWithContext(ctx context.Context) LookupMySQLServerResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMySQLServerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// mysql server edition.
func (o LookupMySQLServerResultOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.Edition }).(pulumi.StringPtrOutput)
}

// The list of errors.
func (o LookupMySQLServerResultOutput) Errors() ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) []ErrorResponse { return v.Errors }).(ErrorResponseArrayOutput)
}

// The Server IP/host name.
func (o LookupMySQLServerResultOutput) HostIp() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) []string { return v.HostIp }).(pulumi.StringArrayOutput)
}

// The Server IP/host name.
func (o LookupMySQLServerResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMySQLServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource labels.
func (o LookupMySQLServerResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// discovery Machine Id
func (o LookupMySQLServerResultOutput) MachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.MachineId }).(pulumi.StringPtrOutput)
}

// The mysql server version.
func (o LookupMySQLServerResultOutput) MysqlVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.MysqlVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupMySQLServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of database.
func (o LookupMySQLServerResultOutput) NumberOfDatabase() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *float64 { return v.NumberOfDatabase }).(pulumi.Float64PtrOutput)
}

// MySQL Server port number
func (o LookupMySQLServerResultOutput) PortNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.PortNumber }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o LookupMySQLServerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Time when mysql version support end.
func (o LookupMySQLServerResultOutput) SupportEndIn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.SupportEndIn }).(pulumi.StringPtrOutput)
}

// mysql version support status.
func (o LookupMySQLServerResultOutput) SupportStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) *string { return v.SupportStatus }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMySQLServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupMySQLServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMySQLServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMySQLServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMySQLServerResultOutput{})
}
