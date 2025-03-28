// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets Customization Task error details
func GetCustomizationTaskErrorDetails(ctx *pulumi.Context, args *GetCustomizationTaskErrorDetailsArgs, opts ...pulumi.InvokeOption) (*GetCustomizationTaskErrorDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetCustomizationTaskErrorDetailsResult
	err := ctx.Invoke("azure-native:devcenter/v20250201:getCustomizationTaskErrorDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetCustomizationTaskErrorDetailsArgs struct {
	// The name of the Catalog.
	CatalogName string `pulumi:"catalogName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Task.
	TaskName string `pulumi:"taskName"`
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetCustomizationTaskErrorDetailsResult struct {
	// Errors associated with resources synchronized from the catalog.
	Errors []CatalogErrorDetailsResponse `pulumi:"errors"`
}

func GetCustomizationTaskErrorDetailsOutput(ctx *pulumi.Context, args GetCustomizationTaskErrorDetailsOutputArgs, opts ...pulumi.InvokeOption) GetCustomizationTaskErrorDetailsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCustomizationTaskErrorDetailsResultOutput, error) {
			args := v.(GetCustomizationTaskErrorDetailsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter/v20250201:getCustomizationTaskErrorDetails", args, GetCustomizationTaskErrorDetailsResultOutput{}, options).(GetCustomizationTaskErrorDetailsResultOutput), nil
		}).(GetCustomizationTaskErrorDetailsResultOutput)
}

type GetCustomizationTaskErrorDetailsOutputArgs struct {
	// The name of the Catalog.
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Task.
	TaskName pulumi.StringInput `pulumi:"taskName"`
}

func (GetCustomizationTaskErrorDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomizationTaskErrorDetailsArgs)(nil)).Elem()
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetCustomizationTaskErrorDetailsResultOutput struct{ *pulumi.OutputState }

func (GetCustomizationTaskErrorDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomizationTaskErrorDetailsResult)(nil)).Elem()
}

func (o GetCustomizationTaskErrorDetailsResultOutput) ToGetCustomizationTaskErrorDetailsResultOutput() GetCustomizationTaskErrorDetailsResultOutput {
	return o
}

func (o GetCustomizationTaskErrorDetailsResultOutput) ToGetCustomizationTaskErrorDetailsResultOutputWithContext(ctx context.Context) GetCustomizationTaskErrorDetailsResultOutput {
	return o
}

// Errors associated with resources synchronized from the catalog.
func (o GetCustomizationTaskErrorDetailsResultOutput) Errors() CatalogErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v GetCustomizationTaskErrorDetailsResult) []CatalogErrorDetailsResponse { return v.Errors }).(CatalogErrorDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomizationTaskErrorDetailsResultOutput{})
}
