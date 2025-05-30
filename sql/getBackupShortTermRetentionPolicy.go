// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a database's short term retention policy.
//
// Uses Azure REST API version 2023-08-01.
//
// Other available API versions: 2017-10-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupBackupShortTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupShortTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupShortTermRetentionPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupShortTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql:getBackupShortTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupShortTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The policy name. Should always be "default".
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A short term retention policy.
type LookupBackupShortTermRetentionPolicyResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *int `pulumi:"diffBackupIntervalInHours"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `pulumi:"retentionDays"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupBackupShortTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupBackupShortTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupShortTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBackupShortTermRetentionPolicyResultOutput, error) {
			args := v.(LookupBackupShortTermRetentionPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:sql:getBackupShortTermRetentionPolicy", args, LookupBackupShortTermRetentionPolicyResultOutput{}, options).(LookupBackupShortTermRetentionPolicyResultOutput), nil
		}).(LookupBackupShortTermRetentionPolicyResultOutput)
}

type LookupBackupShortTermRetentionPolicyOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The policy name. Should always be "default".
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupBackupShortTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyArgs)(nil)).Elem()
}

// A short term retention policy.
type LookupBackupShortTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupShortTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutput() LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupBackupShortTermRetentionPolicyResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
func (o LookupBackupShortTermRetentionPolicyResultOutput) DiffBackupIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.DiffBackupIntervalInHours }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
func (o LookupBackupShortTermRetentionPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupShortTermRetentionPolicyResultOutput{})
}
