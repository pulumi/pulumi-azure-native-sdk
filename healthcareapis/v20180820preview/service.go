// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180820preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the service.
//
// Deprecated: Version 2018-08-20-preview will be removed in v2 of the provider.
type Service struct {
	pulumi.CustomResourceState

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ResourceResponseIdentityPtrOutput `pulumi:"identity"`
	// The kind of the service. Valid values are: fhir, fhir-Stu3 and fhir-R4.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponseOutput `pulumi:"properties"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20190916:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20200315:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20200330:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210111:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:Service"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:Service"),
		},
	})
	opts = append(opts, aliases)
	var resource Service
	err := ctx.RegisterResource("azure-native:healthcareapis/v20180820preview:Service", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:healthcareapis/v20180820preview:Service", name, id, state, &resource, opts...)
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
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The kind of the service. Valid values are: fhir, fhir-Stu3 and fhir-R4.
	Kind Kind `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The common properties of a service.
	Properties *ServicesProperties `pulumi:"properties"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName *string `pulumi:"resourceName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity ResourceIdentityPtrInput
	// The kind of the service. Valid values are: fhir, fhir-Stu3 and fhir-R4.
	Kind KindInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The common properties of a service.
	Properties ServicesPropertiesPtrInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
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

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o ServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o ServiceOutput) Identity() ResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Service) ResourceResponseIdentityPtrOutput { return v.Identity }).(ResourceResponseIdentityPtrOutput)
}

// The kind of the service. Valid values are: fhir, fhir-Stu3 and fhir-R4.
func (o ServiceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o ServiceOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *Service) ServicesPropertiesResponseOutput { return v.Properties }).(ServicesPropertiesResponseOutput)
}

// The resource tags.
func (o ServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}