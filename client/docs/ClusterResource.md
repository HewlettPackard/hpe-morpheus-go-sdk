# ClusterResource

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

### NewClusterResource

`func NewClusterResource() *ClusterResource`

NewClusterResource instantiates a new ClusterResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResourceWithDefaults

`func NewClusterResourceWithDefaults() *ClusterResource`

NewClusterResourceWithDefaults instantiates a new ClusterResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterResource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterResource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterResource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetName

`func (o *ClusterResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *ClusterResource) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterResource) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterResource) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterResource) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterResource) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterResource) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterResource) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterResource) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterResource) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterResource) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterResource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRawSec

`func (o *ClusterResource) GetRawSec() map[string]interface{}`

GetRawSec returns the RawSec field if non-nil, zero value otherwise.

### GetRawSecOk

`func (o *ClusterResource) GetRawSecOk() (*map[string]interface{}, bool)`

GetRawSecOk returns a tuple with the RawSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSec

`func (o *ClusterResource) SetRawSec(v map[string]interface{})`

SetRawSec sets RawSec field to given value.

### HasRawSec

`func (o *ClusterResource) HasRawSec() bool`

HasRawSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


