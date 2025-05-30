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

// A file resource
//
// Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-30.
//
// Other available API versions: 2021-06-30, 2021-10-30-preview, 2022-01-30-preview, 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type File struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Custom file properties
	Properties ProjectFilePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFile registers a new resource with the given unique name, arguments, and options.
func NewFile(ctx *pulumi.Context,
	name string, args *FileArgs, opts ...pulumi.ResourceOption) (*File, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20230715preview:File"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20250315preview:File"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource File
	err := ctx.RegisterResource("azure-native:datamigration:File", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFile gets an existing File resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileState, opts ...pulumi.ResourceOption) (*File, error) {
	var resource File
	err := ctx.ReadResource("azure-native:datamigration:File", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering File resources.
type fileState struct {
}

type FileState struct {
}

func (FileState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileState)(nil)).Elem()
}

type fileArgs struct {
	// Name of the File
	FileName *string `pulumi:"fileName"`
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the project
	ProjectName string `pulumi:"projectName"`
	// Custom file properties
	Properties *ProjectFileProperties `pulumi:"properties"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a File resource.
type FileArgs struct {
	// Name of the File
	FileName pulumi.StringPtrInput
	// Name of the resource group
	GroupName pulumi.StringInput
	// Name of the project
	ProjectName pulumi.StringInput
	// Custom file properties
	Properties ProjectFilePropertiesPtrInput
	// Name of the service
	ServiceName pulumi.StringInput
}

func (FileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileArgs)(nil)).Elem()
}

type FileInput interface {
	pulumi.Input

	ToFileOutput() FileOutput
	ToFileOutputWithContext(ctx context.Context) FileOutput
}

func (*File) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (i *File) ToFileOutput() FileOutput {
	return i.ToFileOutputWithContext(context.Background())
}

func (i *File) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileOutput)
}

type FileOutput struct{ *pulumi.OutputState }

func (FileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**File)(nil)).Elem()
}

func (o FileOutput) ToFileOutput() FileOutput {
	return o
}

func (o FileOutput) ToFileOutputWithContext(ctx context.Context) FileOutput {
	return o
}

// The Azure API version of the resource.
func (o FileOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// HTTP strong entity tag value. This is ignored if submitted.
func (o FileOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *File) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o FileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Custom file properties
func (o FileOutput) Properties() ProjectFilePropertiesResponseOutput {
	return o.ApplyT(func(v *File) ProjectFilePropertiesResponseOutput { return v.Properties }).(ProjectFilePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o FileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *File) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o FileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *File) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileOutput{})
}
