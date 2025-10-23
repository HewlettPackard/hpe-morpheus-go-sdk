# GetClusterUpgradeVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | Pointer to **[]string** |  | [optional] 
**CurrentVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewGetClusterUpgradeVersions200Response

`func NewGetClusterUpgradeVersions200Response() *GetClusterUpgradeVersions200Response`

NewGetClusterUpgradeVersions200Response instantiates a new GetClusterUpgradeVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterUpgradeVersions200ResponseWithDefaults

`func NewGetClusterUpgradeVersions200ResponseWithDefaults() *GetClusterUpgradeVersions200Response`

NewGetClusterUpgradeVersions200ResponseWithDefaults instantiates a new GetClusterUpgradeVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *GetClusterUpgradeVersions200Response) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetClusterUpgradeVersions200Response) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetClusterUpgradeVersions200Response) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetClusterUpgradeVersions200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *GetClusterUpgradeVersions200Response) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *GetClusterUpgradeVersions200Response) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *GetClusterUpgradeVersions200Response) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *GetClusterUpgradeVersions200Response) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


