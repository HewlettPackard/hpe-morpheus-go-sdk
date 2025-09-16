# ListBootScripts200ResponseAllOfBootScriptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewListBootScripts200ResponseAllOfBootScriptsInner

`func NewListBootScripts200ResponseAllOfBootScriptsInner() *ListBootScripts200ResponseAllOfBootScriptsInner`

NewListBootScripts200ResponseAllOfBootScriptsInner instantiates a new ListBootScripts200ResponseAllOfBootScriptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBootScripts200ResponseAllOfBootScriptsInnerWithDefaults

`func NewListBootScripts200ResponseAllOfBootScriptsInnerWithDefaults() *ListBootScripts200ResponseAllOfBootScriptsInner`

NewListBootScripts200ResponseAllOfBootScriptsInnerWithDefaults instantiates a new ListBootScripts200ResponseAllOfBootScriptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetAccount() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetAccountOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetAccount(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetFileName

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetDescription

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetCreatedBy() ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetCreatedByOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetCreatedBy(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVisibility

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListBootScripts200ResponseAllOfBootScriptsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


