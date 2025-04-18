// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the Package type.
//
// Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-15-preview.
//
// Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Package struct {
	pulumi.CustomResourceState

	// Metadata pertaining to creation and last modification of the resource.
	AllOf SystemDataResponseOutput `pulumi:"allOf"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets or sets the contentLink of the Package.
	ContentLink ContentLinkResponsePtrOutput `pulumi:"contentLink"`
	// Gets or sets the isGlobal flag of the package.
	Default pulumi.BoolPtrOutput `pulumi:"default"`
	// Gets or sets the error info of the Package.
	Error PackageErrorInfoResponsePtrOutput `pulumi:"error"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the provisioning state of the Package.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the Package.
	SizeInBytes pulumi.Float64PtrOutput `pulumi:"sizeInBytes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the version of the Package.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewPackage registers a new resource with the given unique name, arguments, and options.
func NewPackage(ctx *pulumi.Context,
	name string, args *PackageArgs, opts ...pulumi.ResourceOption) (*Package, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ContentLink == nil {
		return nil, errors.New("invalid value for required argument 'ContentLink'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuntimeEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeEnvironmentName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Package"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20241023:Package"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Package
	err := ctx.RegisterResource("azure-native:automation:Package", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPackage gets an existing Package resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PackageState, opts ...pulumi.ResourceOption) (*Package, error) {
	var resource Package
	err := ctx.ReadResource("azure-native:automation:Package", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Package resources.
type packageState struct {
}

type PackageState struct {
}

func (PackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*packageState)(nil)).Elem()
}

type packageArgs struct {
	// The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
	AllOf *TrackedResource `pulumi:"allOf"`
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the package content link.
	ContentLink ContentLink `pulumi:"contentLink"`
	// The name of Package.
	PackageName *string `pulumi:"packageName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Runtime Environment.
	RuntimeEnvironmentName string `pulumi:"runtimeEnvironmentName"`
}

// The set of arguments for constructing a Package resource.
type PackageArgs struct {
	// The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags' and a 'location'
	AllOf TrackedResourcePtrInput
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the package content link.
	ContentLink ContentLinkInput
	// The name of Package.
	PackageName pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Runtime Environment.
	RuntimeEnvironmentName pulumi.StringInput
}

func (PackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packageArgs)(nil)).Elem()
}

type PackageInput interface {
	pulumi.Input

	ToPackageOutput() PackageOutput
	ToPackageOutputWithContext(ctx context.Context) PackageOutput
}

func (*Package) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (i *Package) ToPackageOutput() PackageOutput {
	return i.ToPackageOutputWithContext(context.Background())
}

func (i *Package) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageOutput)
}

type PackageOutput struct{ *pulumi.OutputState }

func (PackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Package)(nil)).Elem()
}

func (o PackageOutput) ToPackageOutput() PackageOutput {
	return o
}

func (o PackageOutput) ToPackageOutputWithContext(ctx context.Context) PackageOutput {
	return o
}

// Metadata pertaining to creation and last modification of the resource.
func (o PackageOutput) AllOf() SystemDataResponseOutput {
	return o.ApplyT(func(v *Package) SystemDataResponseOutput { return v.AllOf }).(SystemDataResponseOutput)
}

// The Azure API version of the resource.
func (o PackageOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the contentLink of the Package.
func (o PackageOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Package) ContentLinkResponsePtrOutput { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the isGlobal flag of the package.
func (o PackageOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Package) pulumi.BoolPtrOutput { return v.Default }).(pulumi.BoolPtrOutput)
}

// Gets or sets the error info of the Package.
func (o PackageOutput) Error() PackageErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v *Package) PackageErrorInfoResponsePtrOutput { return v.Error }).(PackageErrorInfoResponsePtrOutput)
}

// The geo-location where the resource lives
func (o PackageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PackageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state of the Package.
func (o PackageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the size in bytes of the Package.
func (o PackageOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Package) pulumi.Float64PtrOutput { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PackageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Package) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PackageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Package) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PackageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Package) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the version of the Package.
func (o PackageOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Package) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PackageOutput{})
}
