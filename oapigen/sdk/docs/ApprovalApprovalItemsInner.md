# ApprovalApprovalItemsInner

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
**Approval** | Pointer to [**ApprovalApprovalItemsInnerApproval**](ApprovalApprovalItemsInnerApproval.md) |  | [optional] 
**Reference** | Pointer to [**ApprovalApprovalItemsInnerReference**](ApprovalApprovalItemsInnerReference.md) |  | [optional] 

## Methods

### NewApprovalApprovalItemsInner

`func NewApprovalApprovalItemsInner() *ApprovalApprovalItemsInner`

NewApprovalApprovalItemsInner instantiates a new ApprovalApprovalItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ApprovalApprovalItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalApprovalItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalApprovalItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalApprovalItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApprovalApprovalItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalApprovalItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalApprovalItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalApprovalItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *ApprovalApprovalItemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApprovalApprovalItemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApprovalApprovalItemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApprovalApprovalItemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ApprovalApprovalItemsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ApprovalApprovalItemsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalName

`func (o *ApprovalApprovalItemsInner) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ApprovalApprovalItemsInner) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ApprovalApprovalItemsInner) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ApprovalApprovalItemsInner) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalNameNil

`func (o *ApprovalApprovalItemsInner) SetExternalNameNil(b bool)`

 SetExternalNameNil sets the value for ExternalName to be an explicit nil

### UnsetExternalName
`func (o *ApprovalApprovalItemsInner) UnsetExternalName()`

UnsetExternalName ensures that no value is present for ExternalName, not even an explicit nil
### GetInternalId

`func (o *ApprovalApprovalItemsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ApprovalApprovalItemsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ApprovalApprovalItemsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ApprovalApprovalItemsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ApprovalApprovalItemsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ApprovalApprovalItemsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetApprovedBy

`func (o *ApprovalApprovalItemsInner) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *ApprovalApprovalItemsInner) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *ApprovalApprovalItemsInner) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *ApprovalApprovalItemsInner) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetDeniedBy

`func (o *ApprovalApprovalItemsInner) GetDeniedBy() string`

GetDeniedBy returns the DeniedBy field if non-nil, zero value otherwise.

### GetDeniedByOk

`func (o *ApprovalApprovalItemsInner) GetDeniedByOk() (*string, bool)`

GetDeniedByOk returns a tuple with the DeniedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedBy

`func (o *ApprovalApprovalItemsInner) SetDeniedBy(v string)`

SetDeniedBy sets DeniedBy field to given value.

### HasDeniedBy

`func (o *ApprovalApprovalItemsInner) HasDeniedBy() bool`

HasDeniedBy returns a boolean if a field has been set.

### SetDeniedByNil

`func (o *ApprovalApprovalItemsInner) SetDeniedByNil(b bool)`

 SetDeniedByNil sets the value for DeniedBy to be an explicit nil

### UnsetDeniedBy
`func (o *ApprovalApprovalItemsInner) UnsetDeniedBy()`

UnsetDeniedBy ensures that no value is present for DeniedBy, not even an explicit nil
### GetStatus

`func (o *ApprovalApprovalItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalApprovalItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalApprovalItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApprovalApprovalItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApprovalApprovalItemsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApprovalApprovalItemsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApprovalApprovalItemsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApprovalApprovalItemsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ApprovalApprovalItemsInner) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ApprovalApprovalItemsInner) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDateCreated

`func (o *ApprovalApprovalItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApprovalApprovalItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApprovalApprovalItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApprovalApprovalItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApprovalApprovalItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalApprovalItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalApprovalItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApprovalApprovalItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateApproved

`func (o *ApprovalApprovalItemsInner) GetDateApproved() time.Time`

GetDateApproved returns the DateApproved field if non-nil, zero value otherwise.

### GetDateApprovedOk

`func (o *ApprovalApprovalItemsInner) GetDateApprovedOk() (*time.Time, bool)`

GetDateApprovedOk returns a tuple with the DateApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateApproved

`func (o *ApprovalApprovalItemsInner) SetDateApproved(v time.Time)`

SetDateApproved sets DateApproved field to given value.

### HasDateApproved

`func (o *ApprovalApprovalItemsInner) HasDateApproved() bool`

HasDateApproved returns a boolean if a field has been set.

### GetDateDenied

`func (o *ApprovalApprovalItemsInner) GetDateDenied() time.Time`

GetDateDenied returns the DateDenied field if non-nil, zero value otherwise.

### GetDateDeniedOk

`func (o *ApprovalApprovalItemsInner) GetDateDeniedOk() (*time.Time, bool)`

GetDateDeniedOk returns a tuple with the DateDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDenied

`func (o *ApprovalApprovalItemsInner) SetDateDenied(v time.Time)`

SetDateDenied sets DateDenied field to given value.

### HasDateDenied

`func (o *ApprovalApprovalItemsInner) HasDateDenied() bool`

HasDateDenied returns a boolean if a field has been set.

### SetDateDeniedNil

`func (o *ApprovalApprovalItemsInner) SetDateDeniedNil(b bool)`

 SetDateDeniedNil sets the value for DateDenied to be an explicit nil

### UnsetDateDenied
`func (o *ApprovalApprovalItemsInner) UnsetDateDenied()`

UnsetDateDenied ensures that no value is present for DateDenied, not even an explicit nil
### GetApproval

`func (o *ApprovalApprovalItemsInner) GetApproval() ApprovalApprovalItemsInnerApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *ApprovalApprovalItemsInner) GetApprovalOk() (*ApprovalApprovalItemsInnerApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *ApprovalApprovalItemsInner) SetApproval(v ApprovalApprovalItemsInnerApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *ApprovalApprovalItemsInner) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetReference

`func (o *ApprovalApprovalItemsInner) GetReference() ApprovalApprovalItemsInnerReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApprovalApprovalItemsInner) GetReferenceOk() (*ApprovalApprovalItemsInnerReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApprovalApprovalItemsInner) SetReference(v ApprovalApprovalItemsInnerReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApprovalApprovalItemsInner) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


