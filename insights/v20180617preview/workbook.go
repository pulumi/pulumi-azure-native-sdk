// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180617preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Application Insights workbook definition.
type Workbook struct {
	pulumi.CustomResourceState

	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringOutput `pulumi:"category"`
	// The user-defined name (display name) of the workbook.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name. This is GUID value. The display name should be assigned within properties field.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringOutput `pulumi:"serializedData"`
	// ResourceId for a source resource.
	SourceId pulumi.StringPtrOutput `pulumi:"sourceId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique user id of the specific user that owns this workbook.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Workbook version
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewWorkbook registers a new resource with the given unique name, arguments, and options.
func NewWorkbook(ctx *pulumi.Context,
	name string, args *WorkbookArgs, opts ...pulumi.ResourceOption) (*Workbook, error) {
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
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20150501:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20201020:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210308:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210801:Workbook"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220401:Workbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Workbook
	err := ctx.RegisterResource("azure-native:insights/v20180617preview:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkbook gets an existing Workbook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azure-native:insights/v20180617preview:Workbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workbook resources.
type workbookState struct {
}

type WorkbookState struct {
}

func (WorkbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookState)(nil)).Elem()
}

type workbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The user-defined name (display name) of the workbook.
	DisplayName string `pulumi:"displayName"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName *string `pulumi:"resourceName"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// ResourceId for a source resource.
	SourceId string `pulumi:"sourceId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Workbook version
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Workbook resource.
type WorkbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringInput
	// The user-defined name (display name) of the workbook.
	DisplayName pulumi.StringInput
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringPtrInput
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringInput
	// ResourceId for a source resource.
	SourceId pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Workbook version
	Version pulumi.StringPtrInput
}

func (WorkbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookArgs)(nil)).Elem()
}

type WorkbookInput interface {
	pulumi.Input

	ToWorkbookOutput() WorkbookOutput
	ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput
}

func (*Workbook) ElementType() reflect.Type {
	return reflect.TypeOf((**Workbook)(nil)).Elem()
}

func (i *Workbook) ToWorkbookOutput() WorkbookOutput {
	return i.ToWorkbookOutputWithContext(context.Background())
}

func (i *Workbook) ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookOutput)
}

type WorkbookOutput struct{ *pulumi.OutputState }

func (WorkbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workbook)(nil)).Elem()
}

func (o WorkbookOutput) ToWorkbookOutput() WorkbookOutput {
	return o
}

func (o WorkbookOutput) ToWorkbookOutputWithContext(ctx context.Context) WorkbookOutput {
	return o
}

// Workbook category, as defined by the user at creation time.
func (o WorkbookOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Category }).(pulumi.StringOutput)
}

// The user-defined name (display name) of the workbook.
func (o WorkbookOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The kind of workbook. Choices are user and shared.
func (o WorkbookOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location
func (o WorkbookOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name. This is GUID value. The display name should be assigned within properties field.
func (o WorkbookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration of this particular workbook. Configuration data is a string containing valid JSON
func (o WorkbookOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.SerializedData }).(pulumi.StringOutput)
}

// ResourceId for a source resource.
func (o WorkbookOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.SourceId }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o WorkbookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this workbook definition.
func (o WorkbookOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o WorkbookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique user id of the specific user that owns this workbook.
func (o WorkbookOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// Workbook version
func (o WorkbookOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workbook) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkbookOutput{})
}