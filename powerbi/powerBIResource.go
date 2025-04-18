// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2020-06-01. In version 2.x of the Azure Native provider, it used API version 2020-06-01.
type PowerBIResource struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the private endpoint connections of the resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Specifies the tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the tenant id of the resource.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPowerBIResource registers a new resource with the given unique name, arguments, and options.
func NewPowerBIResource(ctx *pulumi.Context,
	name string, args *PowerBIResourceArgs, opts ...pulumi.ResourceOption) (*PowerBIResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerbi/v20200601:PowerBIResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PowerBIResource
	err := ctx.RegisterResource("azure-native:powerbi:PowerBIResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPowerBIResource gets an existing PowerBIResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPowerBIResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PowerBIResourceState, opts ...pulumi.ResourceOption) (*PowerBIResource, error) {
	var resource PowerBIResource
	err := ctx.ReadResource("azure-native:powerbi:PowerBIResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PowerBIResource resources.
type powerBIResourceState struct {
}

type PowerBIResourceState struct {
}

func (PowerBIResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*powerBIResourceState)(nil)).Elem()
}

type powerBIResourceArgs struct {
	// The name of the Azure resource.
	AzureResourceName *string `pulumi:"azureResourceName"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the private endpoint connections of the resource.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the tenant id of the resource.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a PowerBIResource resource.
type PowerBIResourceArgs struct {
	// The name of the Azure resource.
	AzureResourceName pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the private endpoint connections of the resource.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the tags of the resource.
	Tags pulumi.StringMapInput
	// Specifies the tenant id of the resource.
	TenantId pulumi.StringPtrInput
}

func (PowerBIResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*powerBIResourceArgs)(nil)).Elem()
}

type PowerBIResourceInput interface {
	pulumi.Input

	ToPowerBIResourceOutput() PowerBIResourceOutput
	ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput
}

func (*PowerBIResource) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerBIResource)(nil)).Elem()
}

func (i *PowerBIResource) ToPowerBIResourceOutput() PowerBIResourceOutput {
	return i.ToPowerBIResourceOutputWithContext(context.Background())
}

func (i *PowerBIResource) ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerBIResourceOutput)
}

type PowerBIResourceOutput struct{ *pulumi.OutputState }

func (PowerBIResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PowerBIResource)(nil)).Elem()
}

func (o PowerBIResourceOutput) ToPowerBIResourceOutput() PowerBIResourceOutput {
	return o
}

func (o PowerBIResourceOutput) ToPowerBIResourceOutputWithContext(ctx context.Context) PowerBIResourceOutput {
	return o
}

// The Azure API version of the resource.
func (o PowerBIResourceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Specifies the location of the resource.
func (o PowerBIResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o PowerBIResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the private endpoint connections of the resource.
func (o PowerBIResourceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PowerBIResource) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The system metadata relating to this resource.
func (o PowerBIResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PowerBIResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Specifies the tags of the resource.
func (o PowerBIResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the tenant id of the resource.
func (o PowerBIResourceOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// Specifies the type of the resource.
func (o PowerBIResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PowerBIResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PowerBIResourceOutput{})
}
