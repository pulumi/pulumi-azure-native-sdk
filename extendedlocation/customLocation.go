// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package extendedlocation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Custom Locations definition.
// API Version: 2021-03-15-preview.
type CustomLocation struct {
	pulumi.CustomResourceState

	// This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication CustomLocationPropertiesResponseAuthenticationPtrOutput `pulumi:"authentication"`
	// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIds pulumi.StringArrayOutput `pulumi:"clusterExtensionIds"`
	// Display name for the Custom Locations location.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceId pulumi.StringPtrOutput `pulumi:"hostResourceId"`
	// Type of host the Custom Locations is referencing (Kubernetes, etc...).
	HostType pulumi.StringPtrOutput `pulumi:"hostType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Kubernetes namespace that will be created on the specified cluster.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Provisioning State for the Custom Location.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomLocation registers a new resource with the given unique name, arguments, and options.
func NewCustomLocation(ctx *pulumi.Context,
	name string, args *CustomLocationArgs, opts ...pulumi.ResourceOption) (*CustomLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:extendedlocation/v20210315preview:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-native:extendedlocation/v20210815:CustomLocation"),
		},
		{
			Type: pulumi.String("azure-native:extendedlocation/v20210831preview:CustomLocation"),
		},
	})
	opts = append(opts, aliases)
	var resource CustomLocation
	err := ctx.RegisterResource("azure-native:extendedlocation:CustomLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLocation gets an existing CustomLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLocationState, opts ...pulumi.ResourceOption) (*CustomLocation, error) {
	var resource CustomLocation
	err := ctx.ReadResource("azure-native:extendedlocation:CustomLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomLocation resources.
type customLocationState struct {
}

type CustomLocationState struct {
}

func (CustomLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLocationState)(nil)).Elem()
}

type customLocationArgs struct {
	// This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication *CustomLocationPropertiesAuthentication `pulumi:"authentication"`
	// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIds []string `pulumi:"clusterExtensionIds"`
	// Display name for the Custom Locations location.
	DisplayName *string `pulumi:"displayName"`
	// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceId *string `pulumi:"hostResourceId"`
	// Type of host the Custom Locations is referencing (Kubernetes, etc...).
	HostType *string `pulumi:"hostType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Kubernetes namespace that will be created on the specified cluster.
	Namespace *string `pulumi:"namespace"`
	// Provisioning State for the Custom Location.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Custom Locations name.
	ResourceName *string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CustomLocation resource.
type CustomLocationArgs struct {
	// This is optional input that contains the authentication that should be used to generate the namespace.
	Authentication CustomLocationPropertiesAuthenticationPtrInput
	// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
	ClusterExtensionIds pulumi.StringArrayInput
	// Display name for the Custom Locations location.
	DisplayName pulumi.StringPtrInput
	// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
	HostResourceId pulumi.StringPtrInput
	// Type of host the Custom Locations is referencing (Kubernetes, etc...).
	HostType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Kubernetes namespace that will be created on the specified cluster.
	Namespace pulumi.StringPtrInput
	// Provisioning State for the Custom Location.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Custom Locations name.
	ResourceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CustomLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLocationArgs)(nil)).Elem()
}

type CustomLocationInput interface {
	pulumi.Input

	ToCustomLocationOutput() CustomLocationOutput
	ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput
}

func (*CustomLocation) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocation)(nil)).Elem()
}

func (i *CustomLocation) ToCustomLocationOutput() CustomLocationOutput {
	return i.ToCustomLocationOutputWithContext(context.Background())
}

func (i *CustomLocation) ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationOutput)
}

type CustomLocationOutput struct{ *pulumi.OutputState }

func (CustomLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocation)(nil)).Elem()
}

func (o CustomLocationOutput) ToCustomLocationOutput() CustomLocationOutput {
	return o
}

func (o CustomLocationOutput) ToCustomLocationOutputWithContext(ctx context.Context) CustomLocationOutput {
	return o
}

// This is optional input that contains the authentication that should be used to generate the namespace.
func (o CustomLocationOutput) Authentication() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o.ApplyT(func(v *CustomLocation) CustomLocationPropertiesResponseAuthenticationPtrOutput {
		return v.Authentication
	}).(CustomLocationPropertiesResponseAuthenticationPtrOutput)
}

// Contains the reference to the add-on that contains charts to deploy CRDs and operators.
func (o CustomLocationOutput) ClusterExtensionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringArrayOutput { return v.ClusterExtensionIds }).(pulumi.StringArrayOutput)
}

// Display name for the Custom Locations location.
func (o CustomLocationOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Connected Cluster or AKS Cluster. The Custom Locations RP will perform a checkAccess API for listAdminCredentials permissions.
func (o CustomLocationOutput) HostResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringPtrOutput { return v.HostResourceId }).(pulumi.StringPtrOutput)
}

// Type of host the Custom Locations is referencing (Kubernetes, etc...).
func (o CustomLocationOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringPtrOutput { return v.HostType }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o CustomLocationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CustomLocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Kubernetes namespace that will be created on the specified cluster.
func (o CustomLocationOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Provisioning State for the Custom Location.
func (o CustomLocationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o CustomLocationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomLocation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CustomLocationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CustomLocationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLocation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomLocationOutput{})
}