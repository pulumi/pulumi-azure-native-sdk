// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package purview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the kafka configuration for the account
//
// Uses Azure REST API version 2024-04-01-preview.
//
// Other available API versions: 2021-12-01, 2023-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native purview [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupKafkaConfiguration(ctx *pulumi.Context, args *LookupKafkaConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupKafkaConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKafkaConfigurationResult
	err := ctx.Invoke("azure-native:purview:getKafkaConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupKafkaConfigurationArgs struct {
	// The name of the account.
	AccountName string `pulumi:"accountName"`
	// Name of kafka configuration.
	KafkaConfigurationName string `pulumi:"kafkaConfigurationName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
type LookupKafkaConfigurationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Consumer group for hook event hub.
	ConsumerGroup *string `pulumi:"consumerGroup"`
	// Credentials to access the event streaming service attached to the purview account.
	Credentials *CredentialsResponse `pulumi:"credentials"`
	// Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
	EventHubPartitionId *string `pulumi:"eventHubPartitionId"`
	EventHubResourceId  *string `pulumi:"eventHubResourceId"`
	// The event hub type.
	EventHubType *string `pulumi:"eventHubType"`
	// The state of the event streaming service
	EventStreamingState *string `pulumi:"eventStreamingState"`
	// The event streaming service type
	EventStreamingType *string `pulumi:"eventStreamingType"`
	// Gets or sets the identifier.
	Id string `pulumi:"id"`
	// Gets or sets the name.
	Name string `pulumi:"name"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData ProxyResourceResponseSystemData `pulumi:"systemData"`
	// Gets or sets the type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupKafkaConfigurationResult
func (val *LookupKafkaConfigurationResult) Defaults() *LookupKafkaConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EventStreamingState == nil {
		eventStreamingState_ := "Enabled"
		tmp.EventStreamingState = &eventStreamingState_
	}
	if tmp.EventStreamingType == nil {
		eventStreamingType_ := "None"
		tmp.EventStreamingType = &eventStreamingType_
	}
	return &tmp
}
func LookupKafkaConfigurationOutput(ctx *pulumi.Context, args LookupKafkaConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupKafkaConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKafkaConfigurationResultOutput, error) {
			args := v.(LookupKafkaConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:purview:getKafkaConfiguration", args, LookupKafkaConfigurationResultOutput{}, options).(LookupKafkaConfigurationResultOutput), nil
		}).(LookupKafkaConfigurationResultOutput)
}

type LookupKafkaConfigurationOutputArgs struct {
	// The name of the account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Name of kafka configuration.
	KafkaConfigurationName pulumi.StringInput `pulumi:"kafkaConfigurationName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKafkaConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConfigurationArgs)(nil)).Elem()
}

// The configuration of the event streaming service resource attached to the Purview account for kafka notifications.
type LookupKafkaConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupKafkaConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKafkaConfigurationResult)(nil)).Elem()
}

func (o LookupKafkaConfigurationResultOutput) ToLookupKafkaConfigurationResultOutput() LookupKafkaConfigurationResultOutput {
	return o
}

func (o LookupKafkaConfigurationResultOutput) ToLookupKafkaConfigurationResultOutputWithContext(ctx context.Context) LookupKafkaConfigurationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupKafkaConfigurationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Consumer group for hook event hub.
func (o LookupKafkaConfigurationResultOutput) ConsumerGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.ConsumerGroup }).(pulumi.StringPtrOutput)
}

// Credentials to access the event streaming service attached to the purview account.
func (o LookupKafkaConfigurationResultOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *CredentialsResponse { return v.Credentials }).(CredentialsResponsePtrOutput)
}

// Optional partition Id for notification event hub. If not set, all partitions will be leveraged.
func (o LookupKafkaConfigurationResultOutput) EventHubPartitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubPartitionId }).(pulumi.StringPtrOutput)
}

func (o LookupKafkaConfigurationResultOutput) EventHubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubResourceId }).(pulumi.StringPtrOutput)
}

// The event hub type.
func (o LookupKafkaConfigurationResultOutput) EventHubType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventHubType }).(pulumi.StringPtrOutput)
}

// The state of the event streaming service
func (o LookupKafkaConfigurationResultOutput) EventStreamingState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventStreamingState }).(pulumi.StringPtrOutput)
}

// The event streaming service type
func (o LookupKafkaConfigurationResultOutput) EventStreamingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) *string { return v.EventStreamingType }).(pulumi.StringPtrOutput)
}

// Gets or sets the identifier.
func (o LookupKafkaConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name.
func (o LookupKafkaConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupKafkaConfigurationResultOutput) SystemData() ProxyResourceResponseSystemDataOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) ProxyResourceResponseSystemData { return v.SystemData }).(ProxyResourceResponseSystemDataOutput)
}

// Gets or sets the type.
func (o LookupKafkaConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKafkaConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKafkaConfigurationResultOutput{})
}
