// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Extension object.
type Extension struct {
	pulumi.CustomResourceState

	// Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity ExtensionResponseAksAssignedIdentityPtrOutput `pulumi:"aksAssignedIdentity"`
	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings pulumi.StringMapOutput `pulumi:"configurationProtectedSettings"`
	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings pulumi.StringMapOutput `pulumi:"configurationSettings"`
	// Custom Location settings properties.
	CustomLocationSettings pulumi.StringMapOutput `pulumi:"customLocationSettings"`
	// The error detail.
	ErrorInfo ErrorDetailResponsePtrOutput `pulumi:"errorInfo"`
	// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType pulumi.StringPtrOutput `pulumi:"extensionType"`
	// Identity of the Extension resource
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Uri of the Helm package
	PackageUri pulumi.StringOutput `pulumi:"packageUri"`
	// The provisioning state of the extension resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
	ReleaseTrain pulumi.StringPtrOutput `pulumi:"releaseTrain"`
	// Scope at which the extension is installed.
	Scope ScopeResponsePtrOutput `pulumi:"scope"`
	// Status from this extension.
	Statuses ExtensionStatusResponseArrayOutput `pulumi:"statuses"`
	// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the extension for this extension, if it is 'pinned' to a specific version. autoUpgradeMinorVersion must be 'false'.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceName'")
	}
	if args.ClusterRp == nil {
		return nil, errors.New("invalid value for required argument 'ClusterRp'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoUpgradeMinorVersion) {
		args.AutoUpgradeMinorVersion = pulumi.BoolPtr(true)
	}
	if isZero(args.ReleaseTrain) {
		args.ReleaseTrain = pulumi.StringPtr("Stable")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20200701preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20210501preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20211101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220101preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220301:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220402preview:Extension"),
		},
		{
			Type: pulumi.String("azure-native:kubernetesconfiguration/v20220701:Extension"),
		},
	})
	opts = append(opts, aliases)
	var resource Extension
	err := ctx.RegisterResource("azure-native:kubernetesconfiguration/v20210901:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure-native:kubernetesconfiguration/v20210901:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
}

type ExtensionState struct {
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity *ExtensionAksAssignedIdentity `pulumi:"aksAssignedIdentity"`
	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The name of the kubernetes cluster.
	ClusterName string `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
	ClusterResourceName string `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
	ClusterRp string `pulumi:"clusterRp"`
	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings map[string]string `pulumi:"configurationProtectedSettings"`
	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]string `pulumi:"configurationSettings"`
	// Name of the Extension.
	ExtensionName *string `pulumi:"extensionName"`
	// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType *string `pulumi:"extensionType"`
	// Identity of the Extension resource
	Identity *Identity `pulumi:"identity"`
	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
	ReleaseTrain *string `pulumi:"releaseTrain"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Scope at which the extension is installed.
	Scope *Scope `pulumi:"scope"`
	// Status from this extension.
	Statuses []ExtensionStatus `pulumi:"statuses"`
	// Version of the extension for this extension, if it is 'pinned' to a specific version. autoUpgradeMinorVersion must be 'false'.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity ExtensionAksAssignedIdentityPtrInput
	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// The name of the kubernetes cluster.
	ClusterName pulumi.StringInput
	// The Kubernetes cluster resource name - either managedClusters (for AKS clusters) or connectedClusters (for OnPrem K8S clusters).
	ClusterResourceName pulumi.StringInput
	// The Kubernetes cluster RP - either Microsoft.ContainerService (for AKS clusters) or Microsoft.Kubernetes (for OnPrem K8S clusters).
	ClusterRp pulumi.StringInput
	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings pulumi.StringMapInput
	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings pulumi.StringMapInput
	// Name of the Extension.
	ExtensionName pulumi.StringPtrInput
	// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType pulumi.StringPtrInput
	// Identity of the Extension resource
	Identity IdentityPtrInput
	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
	ReleaseTrain pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Scope at which the extension is installed.
	Scope ScopePtrInput
	// Status from this extension.
	Statuses ExtensionStatusArrayInput
	// Version of the extension for this extension, if it is 'pinned' to a specific version. autoUpgradeMinorVersion must be 'false'.
	Version pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

type ExtensionInput interface {
	pulumi.Input

	ToExtensionOutput() ExtensionOutput
	ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput
}

func (*Extension) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (i *Extension) ToExtensionOutput() ExtensionOutput {
	return i.ToExtensionOutputWithContext(context.Background())
}

func (i *Extension) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensionOutput)
}

type ExtensionOutput struct{ *pulumi.OutputState }

func (ExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Extension)(nil)).Elem()
}

func (o ExtensionOutput) ToExtensionOutput() ExtensionOutput {
	return o
}

func (o ExtensionOutput) ToExtensionOutputWithContext(ctx context.Context) ExtensionOutput {
	return o
}

// Identity of the Extension resource in an AKS cluster
func (o ExtensionOutput) AksAssignedIdentity() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ApplyT(func(v *Extension) ExtensionResponseAksAssignedIdentityPtrOutput { return v.AksAssignedIdentity }).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

// Flag to note if this extension participates in auto upgrade of minor version, or not.
func (o ExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
func (o ExtensionOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

// Configuration settings, as name-value pairs for configuring this extension.
func (o ExtensionOutput) ConfigurationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.ConfigurationSettings }).(pulumi.StringMapOutput)
}

// Custom Location settings properties.
func (o ExtensionOutput) CustomLocationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringMapOutput { return v.CustomLocationSettings }).(pulumi.StringMapOutput)
}

// The error detail.
func (o ExtensionOutput) ErrorInfo() ErrorDetailResponsePtrOutput {
	return o.ApplyT(func(v *Extension) ErrorDetailResponsePtrOutput { return v.ErrorInfo }).(ErrorDetailResponsePtrOutput)
}

// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
func (o ExtensionOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

// Identity of the Extension resource
func (o ExtensionOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Extension) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The name of the resource
func (o ExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Uri of the Helm package
func (o ExtensionOutput) PackageUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.PackageUri }).(pulumi.StringOutput)
}

// The provisioning state of the extension resource.
func (o ExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
func (o ExtensionOutput) ReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.ReleaseTrain }).(pulumi.StringPtrOutput)
}

// Scope at which the extension is installed.
func (o ExtensionOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v *Extension) ScopeResponsePtrOutput { return v.Scope }).(ScopeResponsePtrOutput)
}

// Status from this extension.
func (o ExtensionOutput) Statuses() ExtensionStatusResponseArrayOutput {
	return o.ApplyT(func(v *Extension) ExtensionStatusResponseArrayOutput { return v.Statuses }).(ExtensionStatusResponseArrayOutput)
}

// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
func (o ExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Extension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the extension for this extension, if it is 'pinned' to a specific version. autoUpgradeMinorVersion must be 'false'.
func (o ExtensionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Extension) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExtensionOutput{})
}