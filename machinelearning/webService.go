// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance of an Azure ML web service resource.
//
// Uses Azure REST API version 2017-01-01. In version 1.x of the Azure Native provider, it used API version 2017-01-01.
//
// Other available API versions: 2016-05-01-preview.
type WebService struct {
	pulumi.CustomResourceState

	// Specifies the location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphResponseOutput `pulumi:"properties"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebService registers a new resource with the given unique name, arguments, and options.
func NewWebService(ctx *pulumi.Context,
	name string, args *WebServiceArgs, opts ...pulumi.ResourceOption) (*WebService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Properties = args.Properties.ToWebServicePropertiesForGraphOutput().ApplyT(func(v WebServicePropertiesForGraph) WebServicePropertiesForGraph { return *v.Defaults() }).(WebServicePropertiesForGraphOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearning/v20160501preview:WebService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20170101:WebService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebService
	err := ctx.RegisterResource("azure-native:machinelearning:WebService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebService gets an existing WebService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebServiceState, opts ...pulumi.ResourceOption) (*WebService, error) {
	var resource WebService
	err := ctx.ReadResource("azure-native:machinelearning:WebService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebService resources.
type webServiceState struct {
}

type WebServiceState struct {
}

func (WebServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceState)(nil)).Elem()
}

type webServiceArgs struct {
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraph `pulumi:"properties"`
	// Name of the resource group in which the web service is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// The name of the web service.
	WebServiceName *string `pulumi:"webServiceName"`
}

// The set of arguments for constructing a WebService resource.
type WebServiceArgs struct {
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Contains the property payload that describes the web service.
	Properties WebServicePropertiesForGraphInput
	// Name of the resource group in which the web service is located.
	ResourceGroupName pulumi.StringInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// The name of the web service.
	WebServiceName pulumi.StringPtrInput
}

func (WebServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webServiceArgs)(nil)).Elem()
}

type WebServiceInput interface {
	pulumi.Input

	ToWebServiceOutput() WebServiceOutput
	ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput
}

func (*WebService) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (i *WebService) ToWebServiceOutput() WebServiceOutput {
	return i.ToWebServiceOutputWithContext(context.Background())
}

func (i *WebService) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebServiceOutput)
}

type WebServiceOutput struct{ *pulumi.OutputState }

func (WebServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebService)(nil)).Elem()
}

func (o WebServiceOutput) ToWebServiceOutput() WebServiceOutput {
	return o
}

func (o WebServiceOutput) ToWebServiceOutputWithContext(ctx context.Context) WebServiceOutput {
	return o
}

// Specifies the location of the resource.
func (o WebServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Specifies the name of the resource.
func (o WebServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Contains the property payload that describes the web service.
func (o WebServiceOutput) Properties() WebServicePropertiesForGraphResponseOutput {
	return o.ApplyT(func(v *WebService) WebServicePropertiesForGraphResponseOutput { return v.Properties }).(WebServicePropertiesForGraphResponseOutput)
}

// Contains resource tags defined as key/value pairs.
func (o WebServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o WebServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebServiceOutput{})
}
