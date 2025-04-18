// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The representation of an edge module.
//
// Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-11-01-preview.
type EdgeModule struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Internal ID generated for the instance of the Video Analyzer edge module.
	EdgeModuleId pulumi.StringOutput `pulumi:"edgeModuleId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEdgeModule registers a new resource with the given unique name, arguments, and options.
func NewEdgeModule(ctx *pulumi.Context,
	name string, args *EdgeModuleArgs, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20210501preview:EdgeModule"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:EdgeModule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EdgeModule
	err := ctx.RegisterResource("azure-native:videoanalyzer:EdgeModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgeModule gets an existing EdgeModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgeModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgeModuleState, opts ...pulumi.ResourceOption) (*EdgeModule, error) {
	var resource EdgeModule
	err := ctx.ReadResource("azure-native:videoanalyzer:EdgeModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgeModule resources.
type edgeModuleState struct {
}

type EdgeModuleState struct {
}

func (EdgeModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleState)(nil)).Elem()
}

type edgeModuleArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// The Edge Module name.
	EdgeModuleName *string `pulumi:"edgeModuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EdgeModule resource.
type EdgeModuleArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput
	// The Edge Module name.
	EdgeModuleName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (EdgeModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgeModuleArgs)(nil)).Elem()
}

type EdgeModuleInput interface {
	pulumi.Input

	ToEdgeModuleOutput() EdgeModuleOutput
	ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput
}

func (*EdgeModule) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeModule)(nil)).Elem()
}

func (i *EdgeModule) ToEdgeModuleOutput() EdgeModuleOutput {
	return i.ToEdgeModuleOutputWithContext(context.Background())
}

func (i *EdgeModule) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeModuleOutput)
}

type EdgeModuleOutput struct{ *pulumi.OutputState }

func (EdgeModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgeModule)(nil)).Elem()
}

func (o EdgeModuleOutput) ToEdgeModuleOutput() EdgeModuleOutput {
	return o
}

func (o EdgeModuleOutput) ToEdgeModuleOutputWithContext(ctx context.Context) EdgeModuleOutput {
	return o
}

// The Azure API version of the resource.
func (o EdgeModuleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeModule) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Internal ID generated for the instance of the Video Analyzer edge module.
func (o EdgeModuleOutput) EdgeModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeModule) pulumi.StringOutput { return v.EdgeModuleId }).(pulumi.StringOutput)
}

// The name of the resource
func (o EdgeModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EdgeModuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EdgeModule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EdgeModuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgeModule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EdgeModuleOutput{})
}
