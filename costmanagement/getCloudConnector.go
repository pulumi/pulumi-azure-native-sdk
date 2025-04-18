// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a cloud connector definition
//
// Uses Azure REST API version 2019-03-01-preview.
func LookupCloudConnector(ctx *pulumi.Context, args *LookupCloudConnectorArgs, opts ...pulumi.InvokeOption) (*LookupCloudConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudConnectorResult
	err := ctx.Invoke("azure-native:costmanagement:getCloudConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudConnectorArgs struct {
	// Connector Name.
	ConnectorName string `pulumi:"connectorName"`
	// May be used to expand the collectionInfo property. By default, collectionInfo is not included.
	Expand *string `pulumi:"expand"`
}

// The Connector model definition
type LookupCloudConnectorResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Connector billing model
	BillingModel *string `pulumi:"billingModel"`
	// Collection information
	CollectionInfo ConnectorCollectionInfoResponse `pulumi:"collectionInfo"`
	// Connector definition creation datetime
	CreatedOn string `pulumi:"createdOn"`
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey *string `pulumi:"credentialsKey"`
	// Number of days remaining of trial
	DaysTrialRemaining int `pulumi:"daysTrialRemaining"`
	// Default ManagementGroupId
	DefaultManagementGroupId *string `pulumi:"defaultManagementGroupId"`
	// Connector DisplayName
	DisplayName *string `pulumi:"displayName"`
	// Associated ExternalBillingAccountId
	ExternalBillingAccountId string `pulumi:"externalBillingAccountId"`
	// Connector id
	Id string `pulumi:"id"`
	// Connector kind (eg aws)
	Kind *string `pulumi:"kind"`
	// Connector last modified datetime
	ModifiedOn string `pulumi:"modifiedOn"`
	// Connector name
	Name string `pulumi:"name"`
	// The display name of the providerBillingAccountId as defined on the external provider
	ProviderBillingAccountDisplayName string `pulumi:"providerBillingAccountDisplayName"`
	// Connector providerBillingAccountId, determined from credentials (eg AWS Consolidated account number)
	ProviderBillingAccountId string `pulumi:"providerBillingAccountId"`
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId *string `pulumi:"reportId"`
	// Connector status
	Status string `pulumi:"status"`
	// Billing SubscriptionId
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Connector type
	Type string `pulumi:"type"`
}

func LookupCloudConnectorOutput(ctx *pulumi.Context, args LookupCloudConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupCloudConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudConnectorResultOutput, error) {
			args := v.(LookupCloudConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:costmanagement:getCloudConnector", args, LookupCloudConnectorResultOutput{}, options).(LookupCloudConnectorResultOutput), nil
		}).(LookupCloudConnectorResultOutput)
}

type LookupCloudConnectorOutputArgs struct {
	// Connector Name.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// May be used to expand the collectionInfo property. By default, collectionInfo is not included.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
}

func (LookupCloudConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudConnectorArgs)(nil)).Elem()
}

// The Connector model definition
type LookupCloudConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupCloudConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudConnectorResult)(nil)).Elem()
}

func (o LookupCloudConnectorResultOutput) ToLookupCloudConnectorResultOutput() LookupCloudConnectorResultOutput {
	return o
}

func (o LookupCloudConnectorResultOutput) ToLookupCloudConnectorResultOutputWithContext(ctx context.Context) LookupCloudConnectorResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCloudConnectorResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Connector billing model
func (o LookupCloudConnectorResultOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.BillingModel }).(pulumi.StringPtrOutput)
}

// Collection information
func (o LookupCloudConnectorResultOutput) CollectionInfo() ConnectorCollectionInfoResponseOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) ConnectorCollectionInfoResponse { return v.CollectionInfo }).(ConnectorCollectionInfoResponseOutput)
}

// Connector definition creation datetime
func (o LookupCloudConnectorResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Credentials authentication key (eg AWS ARN)
func (o LookupCloudConnectorResultOutput) CredentialsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.CredentialsKey }).(pulumi.StringPtrOutput)
}

// Number of days remaining of trial
func (o LookupCloudConnectorResultOutput) DaysTrialRemaining() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) int { return v.DaysTrialRemaining }).(pulumi.IntOutput)
}

// Default ManagementGroupId
func (o LookupCloudConnectorResultOutput) DefaultManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.DefaultManagementGroupId }).(pulumi.StringPtrOutput)
}

// Connector DisplayName
func (o LookupCloudConnectorResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Associated ExternalBillingAccountId
func (o LookupCloudConnectorResultOutput) ExternalBillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ExternalBillingAccountId }).(pulumi.StringOutput)
}

// Connector id
func (o LookupCloudConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Connector kind (eg aws)
func (o LookupCloudConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Connector last modified datetime
func (o LookupCloudConnectorResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

// Connector name
func (o LookupCloudConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The display name of the providerBillingAccountId as defined on the external provider
func (o LookupCloudConnectorResultOutput) ProviderBillingAccountDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ProviderBillingAccountDisplayName }).(pulumi.StringOutput)
}

// Connector providerBillingAccountId, determined from credentials (eg AWS Consolidated account number)
func (o LookupCloudConnectorResultOutput) ProviderBillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.ProviderBillingAccountId }).(pulumi.StringOutput)
}

// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
func (o LookupCloudConnectorResultOutput) ReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.ReportId }).(pulumi.StringPtrOutput)
}

// Connector status
func (o LookupCloudConnectorResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Status }).(pulumi.StringOutput)
}

// Billing SubscriptionId
func (o LookupCloudConnectorResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Connector type
func (o LookupCloudConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudConnectorResultOutput{})
}
