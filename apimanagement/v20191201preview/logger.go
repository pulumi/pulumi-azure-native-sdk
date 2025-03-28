// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Logger details.
type Logger struct {
	pulumi.CustomResourceState

	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials pulumi.StringMapOutput `pulumi:"credentials"`
	// Logger description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered pulumi.BoolPtrOutput `pulumi:"isBuffered"`
	// Logger type.
	LoggerType pulumi.StringOutput `pulumi:"loggerType"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLogger registers a new resource with the given unique name, arguments, and options.
func NewLogger(ctx *pulumi.Context,
	name string, args *LoggerArgs, opts ...pulumi.ResourceOption) (*Logger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Credentials == nil {
		return nil, errors.New("invalid value for required argument 'Credentials'")
	}
	if args.LoggerType == nil {
		return nil, errors.New("invalid value for required argument 'LoggerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Logger"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Logger
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201preview:Logger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogger gets an existing Logger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggerState, opts ...pulumi.ResourceOption) (*Logger, error) {
	var resource Logger
	err := ctx.ReadResource("azure-native:apimanagement/v20191201preview:Logger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Logger resources.
type loggerState struct {
}

type LoggerState struct {
}

func (LoggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerState)(nil)).Elem()
}

type loggerArgs struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials map[string]string `pulumi:"credentials"`
	// Logger description.
	Description *string `pulumi:"description"`
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered *bool `pulumi:"isBuffered"`
	// Logger identifier. Must be unique in the API Management service instance.
	LoggerId *string `pulumi:"loggerId"`
	// Logger type.
	LoggerType string `pulumi:"loggerType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Logger resource.
type LoggerArgs struct {
	// The name and SendRule connection string of the event hub for azureEventHub logger.
	// Instrumentation key for applicationInsights logger.
	Credentials pulumi.StringMapInput
	// Logger description.
	Description pulumi.StringPtrInput
	// Whether records are buffered in the logger before publishing. Default is assumed to be true.
	IsBuffered pulumi.BoolPtrInput
	// Logger identifier. Must be unique in the API Management service instance.
	LoggerId pulumi.StringPtrInput
	// Logger type.
	LoggerType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (LoggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerArgs)(nil)).Elem()
}

type LoggerInput interface {
	pulumi.Input

	ToLoggerOutput() LoggerOutput
	ToLoggerOutputWithContext(ctx context.Context) LoggerOutput
}

func (*Logger) ElementType() reflect.Type {
	return reflect.TypeOf((**Logger)(nil)).Elem()
}

func (i *Logger) ToLoggerOutput() LoggerOutput {
	return i.ToLoggerOutputWithContext(context.Background())
}

func (i *Logger) ToLoggerOutputWithContext(ctx context.Context) LoggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerOutput)
}

type LoggerOutput struct{ *pulumi.OutputState }

func (LoggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Logger)(nil)).Elem()
}

func (o LoggerOutput) ToLoggerOutput() LoggerOutput {
	return o
}

func (o LoggerOutput) ToLoggerOutputWithContext(ctx context.Context) LoggerOutput {
	return o
}

// The name and SendRule connection string of the event hub for azureEventHub logger.
// Instrumentation key for applicationInsights logger.
func (o LoggerOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringMapOutput { return v.Credentials }).(pulumi.StringMapOutput)
}

// Logger description.
func (o LoggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether records are buffered in the logger before publishing. Default is assumed to be true.
func (o LoggerOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.BoolPtrOutput { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

// Logger type.
func (o LoggerOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.LoggerType }).(pulumi.StringOutput)
}

// Resource name.
func (o LoggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Id of a log target (either Azure Event Hub resource or Azure Application Insights resource).
func (o LoggerOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Resource type for API Management resource.
func (o LoggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoggerOutput{})
}
