# GetApprovals200ResponseApprovalApprovalItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalName** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ApprovedBy** | Pointer to **string** |  | [optional] 
**DeniedBy** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DateApproved** | Pointer to **time.Time** |  | [optional] 
**DateDenied** | Pointer to **NullableTime** |  | [optional] 
**Approval** | Pointer to [**GetApprovals200ResponseApprovalApprovalItemsInnerApproval**](GetApprovals200ResponseApprovalApprovalItemsInnerApproval.md) |  | [optional] 
**Reference** | Pointer to [**GetApprovals200ResponseApprovalApprovalItemsInnerReference**](GetApprovals200ResponseApprovalApprovalItemsInnerReference.md) |  | [optional] 

## Methods

### NewGetApprovals200ResponseApprovalApprovalItemsInner

`func NewGetApprovals200ResponseApprovalApprovalItemsInner() *GetApprovals200ResponseApprovalApprovalItemsInner`

NewGetApprovals200ResponseApprovalApprovalItemsInner instantiates a new GetApprovals200ResponseApprovalApprovalItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApprovals200ResponseApprovalApprovalItemsInnerWithDefaults

`func NewGetApprovals200ResponseApprovalApprovalItemsInnerWithDefaults() *GetApprovals200ResponseApprovalApprovalItemsInner`

NewGetApprovals200ResponseApprovalApprovalItemsInnerWithDefaults instantiates a new GetApprovals200ResponseApprovalApprovalItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalNameNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetExternalNameNil(b bool)`

 SetExternalNameNil sets the value for ExternalName to be an explicit nil

### UnsetExternalName
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetExternalName()`

UnsetExternalName ensures that no value is present for ExternalName, not even an explicit nil
### GetInternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetApprovedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetDeniedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDeniedBy() string`

GetDeniedBy returns the DeniedBy field if non-nil, zero value otherwise.

### GetDeniedByOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDeniedByOk() (*string, bool)`

GetDeniedByOk returns a tuple with the DeniedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDeniedBy(v string)`

SetDeniedBy sets DeniedBy field to given value.

### HasDeniedBy

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasDeniedBy() bool`

HasDeniedBy returns a boolean if a field has been set.

### SetDeniedByNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDeniedByNil(b bool)`

 SetDeniedByNil sets the value for DeniedBy to be an explicit nil

### UnsetDeniedBy
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetDeniedBy()`

UnsetDeniedBy ensures that no value is present for DeniedBy, not even an explicit nil
### GetStatus

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDateCreated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateApproved

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateApproved() time.Time`

GetDateApproved returns the DateApproved field if non-nil, zero value otherwise.

### GetDateApprovedOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateApprovedOk() (*time.Time, bool)`

GetDateApprovedOk returns a tuple with the DateApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateApproved

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDateApproved(v time.Time)`

SetDateApproved sets DateApproved field to given value.

### HasDateApproved

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasDateApproved() bool`

HasDateApproved returns a boolean if a field has been set.

### GetDateDenied

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateDenied() time.Time`

GetDateDenied returns the DateDenied field if non-nil, zero value otherwise.

### GetDateDeniedOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetDateDeniedOk() (*time.Time, bool)`

GetDateDeniedOk returns a tuple with the DateDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDenied

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDateDenied(v time.Time)`

SetDateDenied sets DateDenied field to given value.

### HasDateDenied

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasDateDenied() bool`

HasDateDenied returns a boolean if a field has been set.

### SetDateDeniedNil

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetDateDeniedNil(b bool)`

 SetDateDeniedNil sets the value for DateDenied to be an explicit nil

### UnsetDateDenied
`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) UnsetDateDenied()`

UnsetDateDenied ensures that no value is present for DateDenied, not even an explicit nil
### GetApproval

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetApproval() GetApprovals200ResponseApprovalApprovalItemsInnerApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetApprovalOk() (*GetApprovals200ResponseApprovalApprovalItemsInnerApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetApproval(v GetApprovals200ResponseApprovalApprovalItemsInnerApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetReference

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetReference() GetApprovals200ResponseApprovalApprovalItemsInnerReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) GetReferenceOk() (*GetApprovals200ResponseApprovalApprovalItemsInnerReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) SetReference(v GetApprovals200ResponseApprovalApprovalItemsInnerReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetApprovals200ResponseApprovalApprovalItemsInner) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


