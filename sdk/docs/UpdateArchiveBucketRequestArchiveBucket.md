# UpdateArchiveBucketRequestArchiveBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the archive bucket. Must be globally unique. | [optional] 
**Description** | Pointer to **string** | A description for the archive bucket | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**IsPublic** | Pointer to **bool** | Public URL - Set to true to allow anonymous access | [optional] [default to false]
**Accounts** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewUpdateArchiveBucketRequestArchiveBucket

`func NewUpdateArchiveBucketRequestArchiveBucket() *UpdateArchiveBucketRequestArchiveBucket`

NewUpdateArchiveBucketRequestArchiveBucket instantiates a new UpdateArchiveBucketRequestArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateArchiveBucketRequestArchiveBucketWithDefaults

`func NewUpdateArchiveBucketRequestArchiveBucketWithDefaults() *UpdateArchiveBucketRequestArchiveBucket`

NewUpdateArchiveBucketRequestArchiveBucketWithDefaults instantiates a new UpdateArchiveBucketRequestArchiveBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateArchiveBucketRequestArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateArchiveBucketRequestArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateArchiveBucketRequestArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateArchiveBucketRequestArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateArchiveBucketRequestArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateArchiveBucketRequestArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateArchiveBucketRequestArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateArchiveBucketRequestArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetAccounts() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateArchiveBucketRequestArchiveBucket) GetAccountsOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateArchiveBucketRequestArchiveBucket) SetAccounts(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdateArchiveBucketRequestArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


