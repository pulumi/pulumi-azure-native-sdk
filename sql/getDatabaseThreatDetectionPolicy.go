// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a database's threat detection policy.
//
// Uses Azure REST API version 2014-04-01.
func LookupDatabaseThreatDetectionPolicy(ctx *pulumi.Context, args *LookupDatabaseThreatDetectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseThreatDetectionPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseThreatDetectionPolicyResult
	err := ctx.Invoke("azure-native:sql:getDatabaseThreatDetectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseThreatDetectionPolicyArgs struct {
	// The name of the database for which database Threat Detection policy is defined.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName string `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Contains information about a database Threat Detection policy.
type LookupDatabaseThreatDetectionPolicyResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
	DisabledAlerts *string `pulumi:"disabledAlerts"`
	// Specifies that the alert is sent to the account administrators.
	EmailAccountAdmins *string `pulumi:"emailAccountAdmins"`
	// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
	EmailAddresses *string `pulumi:"emailAddresses"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource kind.
	Kind string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the number of days to keep in the Threat Detection audit logs.
	RetentionDays *int `pulumi:"retentionDays"`
	// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
	State string `pulumi:"state"`
	// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
	StorageEndpoint *string `pulumi:"storageEndpoint"`
	// Resource type.
	Type string `pulumi:"type"`
	// Specifies whether to use the default server policy.
	UseServerDefault *string `pulumi:"useServerDefault"`
}

func LookupDatabaseThreatDetectionPolicyOutput(ctx *pulumi.Context, args LookupDatabaseThreatDetectionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseThreatDetectionPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDatabaseThreatDetectionPolicyResultOutput, error) {
			args := v.(LookupDatabaseThreatDetectionPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:sql:getDatabaseThreatDetectionPolicy", args, LookupDatabaseThreatDetectionPolicyResultOutput{}, options).(LookupDatabaseThreatDetectionPolicyResultOutput), nil
		}).(LookupDatabaseThreatDetectionPolicyResultOutput)
}

type LookupDatabaseThreatDetectionPolicyOutputArgs struct {
	// The name of the database for which database Threat Detection policy is defined.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the security alert policy.
	SecurityAlertPolicyName pulumi.StringInput `pulumi:"securityAlertPolicyName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupDatabaseThreatDetectionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseThreatDetectionPolicyArgs)(nil)).Elem()
}

// Contains information about a database Threat Detection policy.
type LookupDatabaseThreatDetectionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseThreatDetectionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseThreatDetectionPolicyResult)(nil)).Elem()
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) ToLookupDatabaseThreatDetectionPolicyResultOutput() LookupDatabaseThreatDetectionPolicyResultOutput {
	return o
}

func (o LookupDatabaseThreatDetectionPolicyResultOutput) ToLookupDatabaseThreatDetectionPolicyResultOutputWithContext(ctx context.Context) LookupDatabaseThreatDetectionPolicyResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) DisabledAlerts() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.DisabledAlerts }).(pulumi.StringPtrOutput)
}

// Specifies that the alert is sent to the account administrators.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) EmailAccountAdmins() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.EmailAccountAdmins }).(pulumi.StringPtrOutput)
}

// Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) EmailAddresses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.EmailAddresses }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource kind.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDatabaseThreatDetectionPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the number of days to keep in the Threat Detection audit logs.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) StorageEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.StorageEndpoint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether to use the default server policy.
func (o LookupDatabaseThreatDetectionPolicyResultOutput) UseServerDefault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseThreatDetectionPolicyResult) *string { return v.UseServerDefault }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseThreatDetectionPolicyResultOutput{})
}
