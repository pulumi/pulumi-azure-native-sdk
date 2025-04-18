// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights private workbook definition.
//
// Uses Azure REST API version 2021-03-08.
//
// Other available API versions: 2015-05-01, 2020-10-20. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native applicationinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type MyWorkbook struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringOutput `pulumi:"category"`
	// The user-defined name of the private workbook.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource etag
	Etag pulumi.StringMapOutput `pulumi:"etag"`
	// Identity used for BYOS
	Identity MyWorkbookManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Configuration of this particular private workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringOutput `pulumi:"serializedData"`
	// Optional resourceId for a source resource.
	SourceId pulumi.StringPtrOutput `pulumi:"sourceId"`
	// BYOS Storage Account URI
	StorageUri pulumi.StringPtrOutput `pulumi:"storageUri"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this private workbook definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Unique user id of the specific user that owns this private workbook.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked private workbook.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewMyWorkbook registers a new resource with the given unique name, arguments, and options.
func NewMyWorkbook(ctx *pulumi.Context,
	name string, args *MyWorkbookArgs, opts ...pulumi.ResourceOption) (*MyWorkbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SerializedData == nil {
		return nil, errors.New("invalid value for required argument 'SerializedData'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:applicationinsights/v20150501:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:applicationinsights/v20201020:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:applicationinsights/v20210308:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210308:MyWorkbook"),
		},
		{
			Type: pulumi.String("azure-native:insights:MyWorkbook"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MyWorkbook
	err := ctx.RegisterResource("azure-native:applicationinsights:MyWorkbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMyWorkbook gets an existing MyWorkbook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMyWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MyWorkbookState, opts ...pulumi.ResourceOption) (*MyWorkbook, error) {
	var resource MyWorkbook
	err := ctx.ReadResource("azure-native:applicationinsights:MyWorkbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MyWorkbook resources.
type myWorkbookState struct {
}

type MyWorkbookState struct {
}

func (MyWorkbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*myWorkbookState)(nil)).Elem()
}

type myWorkbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The user-defined name of the private workbook.
	DisplayName string `pulumi:"displayName"`
	// Azure resource Id
	Id *string `pulumi:"id"`
	// Identity used for BYOS
	Identity *MyWorkbookManagedIdentity `pulumi:"identity"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName *string `pulumi:"resourceName"`
	// Configuration of this particular private workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// Optional resourceId for a source resource.
	SourceId *string `pulumi:"sourceId"`
	// BYOS Storage Account URI
	StorageUri *string `pulumi:"storageUri"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// This instance's version of the data model. This can change as new features are added that can be marked private workbook.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a MyWorkbook resource.
type MyWorkbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringInput
	// The user-defined name of the private workbook.
	DisplayName pulumi.StringInput
	// Azure resource Id
	Id pulumi.StringPtrInput
	// Identity used for BYOS
	Identity MyWorkbookManagedIdentityPtrInput
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringPtrInput
	// Configuration of this particular private workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringInput
	// Optional resourceId for a source resource.
	SourceId pulumi.StringPtrInput
	// BYOS Storage Account URI
	StorageUri pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// This instance's version of the data model. This can change as new features are added that can be marked private workbook.
	Version pulumi.StringPtrInput
}

func (MyWorkbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*myWorkbookArgs)(nil)).Elem()
}

type MyWorkbookInput interface {
	pulumi.Input

	ToMyWorkbookOutput() MyWorkbookOutput
	ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput
}

func (*MyWorkbook) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbook)(nil)).Elem()
}

func (i *MyWorkbook) ToMyWorkbookOutput() MyWorkbookOutput {
	return i.ToMyWorkbookOutputWithContext(context.Background())
}

func (i *MyWorkbook) ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MyWorkbookOutput)
}

type MyWorkbookOutput struct{ *pulumi.OutputState }

func (MyWorkbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MyWorkbook)(nil)).Elem()
}

func (o MyWorkbookOutput) ToMyWorkbookOutput() MyWorkbookOutput {
	return o
}

func (o MyWorkbookOutput) ToMyWorkbookOutputWithContext(ctx context.Context) MyWorkbookOutput {
	return o
}

// The Azure API version of the resource.
func (o MyWorkbookOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Workbook category, as defined by the user at creation time.
func (o MyWorkbookOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// The user-defined name of the private workbook.
func (o MyWorkbookOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource etag
func (o MyWorkbookOutput) Etag() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringMapOutput { return v.Etag }).(pulumi.StringMapOutput)
}

// Identity used for BYOS
func (o MyWorkbookOutput) Identity() MyWorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *MyWorkbook) MyWorkbookManagedIdentityResponsePtrOutput { return v.Identity }).(MyWorkbookManagedIdentityResponsePtrOutput)
}

// The kind of workbook. Choices are user and shared.
func (o MyWorkbookOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location
func (o MyWorkbookOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name
func (o MyWorkbookOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Configuration of this particular private workbook. Configuration data is a string containing valid JSON
func (o MyWorkbookOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.SerializedData }).(pulumi.StringOutput)
}

// Optional resourceId for a source resource.
func (o MyWorkbookOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.SourceId }).(pulumi.StringPtrOutput)
}

// BYOS Storage Account URI
func (o MyWorkbookOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.StorageUri }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MyWorkbookOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MyWorkbook) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o MyWorkbookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this private workbook definition.
func (o MyWorkbookOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o MyWorkbookOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Unique user id of the specific user that owns this private workbook.
func (o MyWorkbookOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// This instance's version of the data model. This can change as new features are added that can be marked private workbook.
func (o MyWorkbookOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MyWorkbook) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MyWorkbookOutput{})
}
