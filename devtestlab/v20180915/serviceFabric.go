// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Service Fabric.
type ServiceFabric struct {
	pulumi.CustomResourceState

	// The applicable schedule for the virtual machine.
	ApplicableSchedule ApplicableScheduleResponseOutput `pulumi:"applicableSchedule"`
	// The resource id of the environment under which the service fabric resource is present
	EnvironmentId pulumi.StringPtrOutput `pulumi:"environmentId"`
	// The backing service fabric resource's id
	ExternalServiceFabricId pulumi.StringPtrOutput `pulumi:"externalServiceFabricId"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
}

// NewServiceFabric registers a new resource with the given unique name, arguments, and options.
func NewServiceFabric(ctx *pulumi.Context,
	name string, args *ServiceFabricArgs, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ServiceFabric"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceFabric
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:ServiceFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceFabric gets an existing ServiceFabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceFabricState, opts ...pulumi.ResourceOption) (*ServiceFabric, error) {
	var resource ServiceFabric
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:ServiceFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceFabric resources.
type serviceFabricState struct {
}

type ServiceFabricState struct {
}

func (ServiceFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricState)(nil)).Elem()
}

type serviceFabricArgs struct {
	// The resource id of the environment under which the service fabric resource is present
	EnvironmentId *string `pulumi:"environmentId"`
	// The backing service fabric resource's id
	ExternalServiceFabricId *string `pulumi:"externalServiceFabricId"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the service fabric.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the user profile.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a ServiceFabric resource.
type ServiceFabricArgs struct {
	// The resource id of the environment under which the service fabric resource is present
	EnvironmentId pulumi.StringPtrInput
	// The backing service fabric resource's id
	ExternalServiceFabricId pulumi.StringPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the service fabric.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The name of the user profile.
	UserName pulumi.StringInput
}

func (ServiceFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricArgs)(nil)).Elem()
}

type ServiceFabricInput interface {
	pulumi.Input

	ToServiceFabricOutput() ServiceFabricOutput
	ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput
}

func (*ServiceFabric) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceFabric)(nil)).Elem()
}

func (i *ServiceFabric) ToServiceFabricOutput() ServiceFabricOutput {
	return i.ToServiceFabricOutputWithContext(context.Background())
}

func (i *ServiceFabric) ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFabricOutput)
}

type ServiceFabricOutput struct{ *pulumi.OutputState }

func (ServiceFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceFabric)(nil)).Elem()
}

func (o ServiceFabricOutput) ToServiceFabricOutput() ServiceFabricOutput {
	return o
}

func (o ServiceFabricOutput) ToServiceFabricOutputWithContext(ctx context.Context) ServiceFabricOutput {
	return o
}

// The applicable schedule for the virtual machine.
func (o ServiceFabricOutput) ApplicableSchedule() ApplicableScheduleResponseOutput {
	return o.ApplyT(func(v *ServiceFabric) ApplicableScheduleResponseOutput { return v.ApplicableSchedule }).(ApplicableScheduleResponseOutput)
}

// The resource id of the environment under which the service fabric resource is present
func (o ServiceFabricOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringPtrOutput { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

// The backing service fabric resource's id
func (o ServiceFabricOutput) ExternalServiceFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringPtrOutput { return v.ExternalServiceFabricId }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o ServiceFabricOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ServiceFabricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o ServiceFabricOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o ServiceFabricOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ServiceFabricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o ServiceFabricOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabric) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceFabricOutput{})
}