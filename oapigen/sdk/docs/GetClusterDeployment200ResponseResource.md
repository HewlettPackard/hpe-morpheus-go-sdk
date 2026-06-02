# GetClusterDeployment200ResponseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RawSec** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetClusterDeployment200ResponseResource

`func NewGetClusterDeployment200ResponseResource() *GetClusterDeployment200ResponseResource`

NewGetClusterDeployment200ResponseResource instantiates a new GetClusterDeployment200ResponseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetClusterDeployment200ResponseResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterDeployment200ResponseResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterDeployment200ResponseResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterDeployment200ResponseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetClusterDeployment200ResponseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetClusterDeployment200ResponseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetClusterDeployment200ResponseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetClusterDeployment200ResponseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *GetClusterDeployment200ResponseResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterDeployment200ResponseResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterDeployment200ResponseResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterDeployment200ResponseResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterDeployment200ResponseResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterDeployment200ResponseResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterDeployment200ResponseResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterDeployment200ResponseResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *GetClusterDeployment200ResponseResource) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetClusterDeployment200ResponseResource) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetClusterDeployment200ResponseResource) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetClusterDeployment200ResponseResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *GetClusterDeployment200ResponseResource) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *GetClusterDeployment200ResponseResource) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *GetClusterDeployment200ResponseResource) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *GetClusterDeployment200ResponseResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *GetClusterDeployment200ResponseResource) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *GetClusterDeployment200ResponseResource) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetConfig

`func (o *GetClusterDeployment200ResponseResource) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetClusterDeployment200ResponseResource) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetClusterDeployment200ResponseResource) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetClusterDeployment200ResponseResource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetClusterDeployment200ResponseResource) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetClusterDeployment200ResponseResource) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetRawSec

`func (o *GetClusterDeployment200ResponseResource) GetRawSec() map[string]interface{}`

GetRawSec returns the RawSec field if non-nil, zero value otherwise.

### GetRawSecOk

`func (o *GetClusterDeployment200ResponseResource) GetRawSecOk() (*map[string]interface{}, bool)`

GetRawSecOk returns a tuple with the RawSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSec

`func (o *GetClusterDeployment200ResponseResource) SetRawSec(v map[string]interface{})`

SetRawSec sets RawSec field to given value.

### HasRawSec

`func (o *GetClusterDeployment200ResponseResource) HasRawSec() bool`

HasRawSec returns a boolean if a field has been set.

### SetRawSecNil

`func (o *GetClusterDeployment200ResponseResource) SetRawSecNil(b bool)`

 SetRawSecNil sets the value for RawSec to be an explicit nil

### UnsetRawSec
`func (o *GetClusterDeployment200ResponseResource) UnsetRawSec()`

UnsetRawSec ensures that no value is present for RawSec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


