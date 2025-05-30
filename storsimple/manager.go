// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The StorSimple Manager.
//
// Uses Azure REST API version 2017-06-01. In version 2.x of the Azure Native provider, it used API version 2017-06-01.
type Manager struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Represents the type of StorSimple Manager.
	CisIntrinsicSettings ManagerIntrinsicSettingsResponsePtrOutput `pulumi:"cisIntrinsicSettings"`
	// The etag of the manager.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The geo location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Specifies the Sku.
	Sku ManagerSkuResponsePtrOutput `pulumi:"sku"`
	// The tags attached to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManager registers a new resource with the given unique name, arguments, and options.
func NewManager(ctx *pulumi.Context,
	name string, args *ManagerArgs, opts ...pulumi.ResourceOption) (*Manager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:Manager"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20170601:Manager"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Manager
	err := ctx.RegisterResource("azure-native:storsimple:Manager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManager gets an existing Manager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerState, opts ...pulumi.ResourceOption) (*Manager, error) {
	var resource Manager
	err := ctx.ReadResource("azure-native:storsimple:Manager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Manager resources.
type managerState struct {
}

type ManagerState struct {
}

func (ManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerState)(nil)).Elem()
}

type managerArgs struct {
	// Represents the type of StorSimple Manager.
	CisIntrinsicSettings *ManagerIntrinsicSettings `pulumi:"cisIntrinsicSettings"`
	// The geo location of the resource.
	Location *string `pulumi:"location"`
	// The manager name
	ManagerName *string `pulumi:"managerName"`
	// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the Sku.
	Sku *ManagerSku `pulumi:"sku"`
	// The tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Manager resource.
type ManagerArgs struct {
	// Represents the type of StorSimple Manager.
	CisIntrinsicSettings ManagerIntrinsicSettingsPtrInput
	// The geo location of the resource.
	Location pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringPtrInput
	// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
	ProvisioningState pulumi.StringPtrInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// Specifies the Sku.
	Sku ManagerSkuPtrInput
	// The tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (ManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerArgs)(nil)).Elem()
}

type ManagerInput interface {
	pulumi.Input

	ToManagerOutput() ManagerOutput
	ToManagerOutputWithContext(ctx context.Context) ManagerOutput
}

func (*Manager) ElementType() reflect.Type {
	return reflect.TypeOf((**Manager)(nil)).Elem()
}

func (i *Manager) ToManagerOutput() ManagerOutput {
	return i.ToManagerOutputWithContext(context.Background())
}

func (i *Manager) ToManagerOutputWithContext(ctx context.Context) ManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerOutput)
}

type ManagerOutput struct{ *pulumi.OutputState }

func (ManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Manager)(nil)).Elem()
}

func (o ManagerOutput) ToManagerOutput() ManagerOutput {
	return o
}

func (o ManagerOutput) ToManagerOutputWithContext(ctx context.Context) ManagerOutput {
	return o
}

// The Azure API version of the resource.
func (o ManagerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Represents the type of StorSimple Manager.
func (o ManagerOutput) CisIntrinsicSettings() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Manager) ManagerIntrinsicSettingsResponsePtrOutput { return v.CisIntrinsicSettings }).(ManagerIntrinsicSettingsResponsePtrOutput)
}

// The etag of the manager.
func (o ManagerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The geo location of the resource.
func (o ManagerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o ManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
func (o ManagerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Specifies the Sku.
func (o ManagerOutput) Sku() ManagerSkuResponsePtrOutput {
	return o.ApplyT(func(v *Manager) ManagerSkuResponsePtrOutput { return v.Sku }).(ManagerSkuResponsePtrOutput)
}

// The tags attached to the resource.
func (o ManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o ManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Manager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagerOutput{})
}
