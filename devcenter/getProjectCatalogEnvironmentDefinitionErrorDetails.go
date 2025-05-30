// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets Environment Definition error details
//
// Uses Azure REST API version 2024-02-01.
//
// Other available API versions: 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetProjectCatalogEnvironmentDefinitionErrorDetails(ctx *pulumi.Context, args *GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs, opts ...pulumi.InvokeOption) (*GetProjectCatalogEnvironmentDefinitionErrorDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetProjectCatalogEnvironmentDefinitionErrorDetailsResult
	err := ctx.Invoke("azure-native:devcenter:getProjectCatalogEnvironmentDefinitionErrorDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs struct {
	// The name of the Catalog.
	CatalogName string `pulumi:"catalogName"`
	// The name of the Environment Definition.
	EnvironmentDefinitionName string `pulumi:"environmentDefinitionName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetProjectCatalogEnvironmentDefinitionErrorDetailsResult struct {
	// Errors associated with resources synchronized from the catalog.
	Errors []CatalogErrorDetailsResponse `pulumi:"errors"`
}

func GetProjectCatalogEnvironmentDefinitionErrorDetailsOutput(ctx *pulumi.Context, args GetProjectCatalogEnvironmentDefinitionErrorDetailsOutputArgs, opts ...pulumi.InvokeOption) GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput, error) {
			args := v.(GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter:getProjectCatalogEnvironmentDefinitionErrorDetails", args, GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput{}, options).(GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput), nil
		}).(GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput)
}

type GetProjectCatalogEnvironmentDefinitionErrorDetailsOutputArgs struct {
	// The name of the Catalog.
	CatalogName pulumi.StringInput `pulumi:"catalogName"`
	// The name of the Environment Definition.
	EnvironmentDefinitionName pulumi.StringInput `pulumi:"environmentDefinitionName"`
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProjectCatalogEnvironmentDefinitionErrorDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectCatalogEnvironmentDefinitionErrorDetailsArgs)(nil)).Elem()
}

// List of validator error details. Populated when changes are made to the resource or its dependent resources that impact the validity of the Catalog resource.
type GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput struct{ *pulumi.OutputState }

func (GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectCatalogEnvironmentDefinitionErrorDetailsResult)(nil)).Elem()
}

func (o GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput) ToGetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput() GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput {
	return o
}

func (o GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput) ToGetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutputWithContext(ctx context.Context) GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput {
	return o
}

// Errors associated with resources synchronized from the catalog.
func (o GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput) Errors() CatalogErrorDetailsResponseArrayOutput {
	return o.ApplyT(func(v GetProjectCatalogEnvironmentDefinitionErrorDetailsResult) []CatalogErrorDetailsResponse {
		return v.Errors
	}).(CatalogErrorDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectCatalogEnvironmentDefinitionErrorDetailsResultOutput{})
}
