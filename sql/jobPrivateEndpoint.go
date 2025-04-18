// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A job agent private endpoint.
//
// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.
//
// Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type JobPrivateEndpoint struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint id of the private endpoint.
	PrivateEndpointId pulumi.StringOutput `pulumi:"privateEndpointId"`
	// ARM resource id of the server the private endpoint will target.
	TargetServerAzureResourceId pulumi.StringOutput `pulumi:"targetServerAzureResourceId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewJobPrivateEndpoint(ctx *pulumi.Context,
	name string, args *JobPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*JobPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TargetServerAzureResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetServerAzureResourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:JobPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801:JobPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:JobPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20240501preview:JobPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobPrivateEndpoint
	err := ctx.RegisterResource("azure-native:sql:JobPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobPrivateEndpoint gets an existing JobPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobPrivateEndpointState, opts ...pulumi.ResourceOption) (*JobPrivateEndpoint, error) {
	var resource JobPrivateEndpoint
	err := ctx.ReadResource("azure-native:sql:JobPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobPrivateEndpoint resources.
type jobPrivateEndpointState struct {
}

type JobPrivateEndpointState struct {
}

func (JobPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobPrivateEndpointState)(nil)).Elem()
}

type jobPrivateEndpointArgs struct {
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the private endpoint.
	PrivateEndpointName *string `pulumi:"privateEndpointName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// ARM resource id of the server the private endpoint will target.
	TargetServerAzureResourceId string `pulumi:"targetServerAzureResourceId"`
}

// The set of arguments for constructing a JobPrivateEndpoint resource.
type JobPrivateEndpointArgs struct {
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// ARM resource id of the server the private endpoint will target.
	TargetServerAzureResourceId pulumi.StringInput
}

func (JobPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobPrivateEndpointArgs)(nil)).Elem()
}

type JobPrivateEndpointInput interface {
	pulumi.Input

	ToJobPrivateEndpointOutput() JobPrivateEndpointOutput
	ToJobPrivateEndpointOutputWithContext(ctx context.Context) JobPrivateEndpointOutput
}

func (*JobPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPrivateEndpoint)(nil)).Elem()
}

func (i *JobPrivateEndpoint) ToJobPrivateEndpointOutput() JobPrivateEndpointOutput {
	return i.ToJobPrivateEndpointOutputWithContext(context.Background())
}

func (i *JobPrivateEndpoint) ToJobPrivateEndpointOutputWithContext(ctx context.Context) JobPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPrivateEndpointOutput)
}

type JobPrivateEndpointOutput struct{ *pulumi.OutputState }

func (JobPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobPrivateEndpoint)(nil)).Elem()
}

func (o JobPrivateEndpointOutput) ToJobPrivateEndpointOutput() JobPrivateEndpointOutput {
	return o
}

func (o JobPrivateEndpointOutput) ToJobPrivateEndpointOutputWithContext(ctx context.Context) JobPrivateEndpointOutput {
	return o
}

// The Azure API version of the resource.
func (o JobPrivateEndpointOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *JobPrivateEndpoint) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource name.
func (o JobPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint id of the private endpoint.
func (o JobPrivateEndpointOutput) PrivateEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobPrivateEndpoint) pulumi.StringOutput { return v.PrivateEndpointId }).(pulumi.StringOutput)
}

// ARM resource id of the server the private endpoint will target.
func (o JobPrivateEndpointOutput) TargetServerAzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobPrivateEndpoint) pulumi.StringOutput { return v.TargetServerAzureResourceId }).(pulumi.StringOutput)
}

// Resource type.
func (o JobPrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobPrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobPrivateEndpointOutput{})
}
