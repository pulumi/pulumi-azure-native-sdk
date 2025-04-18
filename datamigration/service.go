// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datamigration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Database Migration Service (classic) resource
//
// Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-30.
//
// Other available API versions: 2021-06-30, 2021-10-30-preview, 2022-01-30-preview, 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Service struct {
	pulumi.CustomResourceState

	// The time delay before the service is auto-stopped when idle.
	AutoStopDelay pulumi.StringPtrOutput `pulumi:"autoStopDelay"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Whether service resources should be deleted when stopped. (Turned on by default)
	DeleteResourcesOnStop pulumi.BoolPtrOutput `pulumi:"deleteResourcesOnStop"`
	// HTTP strong entity tag value. Ignored if submitted
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind     pulumi.StringPtrOutput `pulumi:"kind"`
	Location pulumi.StringPtrOutput `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	// The resource's provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Service SKU
	Sku        ServiceSkuResponsePtrOutput `pulumi:"sku"`
	SystemData SystemDataResponseOutput    `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput      `pulumi:"tags"`
	Type       pulumi.StringOutput         `pulumi:"type"`
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId pulumi.StringPtrOutput `pulumi:"virtualNicId"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId pulumi.StringPtrOutput `pulumi:"virtualSubnetId"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20230715preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20250315preview:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:datamigration:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:datamigration:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The time delay before the service is auto-stopped when idle.
	AutoStopDelay *string `pulumi:"autoStopDelay"`
	// Whether service resources should be deleted when stopped. (Turned on by default)
	DeleteResourcesOnStop *bool `pulumi:"deleteResourcesOnStop"`
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind     *string `pulumi:"kind"`
	Location *string `pulumi:"location"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey *string `pulumi:"publicKey"`
	// Name of the service
	ServiceName *string `pulumi:"serviceName"`
	// Service SKU
	Sku  *ServiceSku       `pulumi:"sku"`
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId *string `pulumi:"virtualNicId"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId *string `pulumi:"virtualSubnetId"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The time delay before the service is auto-stopped when idle.
	AutoStopDelay pulumi.StringPtrInput
	// Whether service resources should be deleted when stopped. (Turned on by default)
	DeleteResourcesOnStop pulumi.BoolPtrInput
	// Name of the resource group
	GroupName pulumi.StringInput
	// The resource kind. Only 'vm' (the default) is supported.
	Kind     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey pulumi.StringPtrInput
	// Name of the service
	ServiceName pulumi.StringPtrInput
	// Service SKU
	Sku  ServiceSkuPtrInput
	Tags pulumi.StringMapInput
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId pulumi.StringPtrInput
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId pulumi.StringPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// The time delay before the service is auto-stopped when idle.
func (o ServiceOutput) AutoStopDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.AutoStopDelay }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o ServiceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Whether service resources should be deleted when stopped. (Turned on by default)
func (o ServiceOutput) DeleteResourcesOnStop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.BoolPtrOutput { return v.DeleteResourcesOnStop }).(pulumi.BoolPtrOutput)
}

// HTTP strong entity tag value. Ignored if submitted
func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource kind. Only 'vm' (the default) is supported.
func (o ServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource's provisioning state
func (o ServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public key of the service, used to encrypt secrets sent to the service
func (o ServiceOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Service SKU
func (o ServiceOutput) Sku() ServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Service) ServiceSkuResponsePtrOutput { return v.Sku }).(ServiceSkuResponsePtrOutput)
}

func (o ServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Service) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The ID of the Microsoft.Network/networkInterfaces resource which the service have
func (o ServiceOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
func (o ServiceOutput) VirtualSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.VirtualSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
