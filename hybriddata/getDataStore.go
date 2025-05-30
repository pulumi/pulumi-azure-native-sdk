// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybriddata

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This method gets the data store/repository by name.
//
// Uses Azure REST API version 2019-06-01.
func LookupDataStore(ctx *pulumi.Context, args *LookupDataStoreArgs, opts ...pulumi.InvokeOption) (*LookupDataStoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataStoreResult
	err := ctx.Invoke("azure-native:hybriddata:getDataStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataStoreArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName string `pulumi:"dataManagerName"`
	// The data store/repository name queried.
	DataStoreName string `pulumi:"dataStoreName"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data store.
type LookupDataStoreResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets []CustomerSecretResponse `pulumi:"customerSecrets"`
	// The arm id of the data store type.
	DataStoreTypeId string `pulumi:"dataStoreTypeId"`
	// A generic json used differently by each data source type.
	ExtendedProperties interface{} `pulumi:"extendedProperties"`
	// Id of the object.
	Id string `pulumi:"id"`
	// Name of the object.
	Name string `pulumi:"name"`
	// Arm Id for the manager resource to which the data source is associated. This is optional.
	RepositoryId *string `pulumi:"repositoryId"`
	// State of the data source.
	State string `pulumi:"state"`
	// Type of the object.
	Type string `pulumi:"type"`
}

func LookupDataStoreOutput(ctx *pulumi.Context, args LookupDataStoreOutputArgs, opts ...pulumi.InvokeOption) LookupDataStoreResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataStoreResultOutput, error) {
			args := v.(LookupDataStoreArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybriddata:getDataStore", args, LookupDataStoreResultOutput{}, options).(LookupDataStoreResultOutput), nil
		}).(LookupDataStoreResultOutput)
}

type LookupDataStoreOutputArgs struct {
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName pulumi.StringInput `pulumi:"dataManagerName"`
	// The data store/repository name queried.
	DataStoreName pulumi.StringInput `pulumi:"dataStoreName"`
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataStoreArgs)(nil)).Elem()
}

// Data store.
type LookupDataStoreResultOutput struct{ *pulumi.OutputState }

func (LookupDataStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataStoreResult)(nil)).Elem()
}

func (o LookupDataStoreResultOutput) ToLookupDataStoreResultOutput() LookupDataStoreResultOutput {
	return o
}

func (o LookupDataStoreResultOutput) ToLookupDataStoreResultOutputWithContext(ctx context.Context) LookupDataStoreResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDataStoreResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
func (o LookupDataStoreResultOutput) CustomerSecrets() CustomerSecretResponseArrayOutput {
	return o.ApplyT(func(v LookupDataStoreResult) []CustomerSecretResponse { return v.CustomerSecrets }).(CustomerSecretResponseArrayOutput)
}

// The arm id of the data store type.
func (o LookupDataStoreResultOutput) DataStoreTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.DataStoreTypeId }).(pulumi.StringOutput)
}

// A generic json used differently by each data source type.
func (o LookupDataStoreResultOutput) ExtendedProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataStoreResult) interface{} { return v.ExtendedProperties }).(pulumi.AnyOutput)
}

// Id of the object.
func (o LookupDataStoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the object.
func (o LookupDataStoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// Arm Id for the manager resource to which the data source is associated. This is optional.
func (o LookupDataStoreResultOutput) RepositoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataStoreResult) *string { return v.RepositoryId }).(pulumi.StringPtrOutput)
}

// State of the data source.
func (o LookupDataStoreResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.State }).(pulumi.StringOutput)
}

// Type of the object.
func (o LookupDataStoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataStoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataStoreResultOutput{})
}
