# ApprovalItemApprovalItem

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
**Approval** | Pointer to [**ApprovalItemApprovalItemApproval**](ApprovalItemApprovalItemApproval.md) |  | [optional] 
**Reference** | Pointer to [**ApprovalItemApprovalItemReference**](ApprovalItemApprovalItemReference.md) |  | [optional] 
**Details** | Pointer to [**[]ApprovalItemApprovalItemDetailsInner**](ApprovalItemApprovalItemDetailsInner.md) | List of request change details associated with the approval item. Only present when the approval is for a reconfigure request.  | [optional] 

## Methods

### NewApprovalItemApprovalItem

`func NewApprovalItemApprovalItem() *ApprovalItemApprovalItem`

NewApprovalItemApprovalItem instantiates a new ApprovalItemApprovalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ApprovalItemApprovalItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalItemApprovalItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalItemApprovalItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalItemApprovalItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApprovalItemApprovalItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalItemApprovalItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalItemApprovalItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalItemApprovalItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *ApprovalItemApprovalItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApprovalItemApprovalItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApprovalItemApprovalItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApprovalItemApprovalItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ApprovalItemApprovalItem) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ApprovalItemApprovalItem) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalName

`func (o *ApprovalItemApprovalItem) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ApprovalItemApprovalItem) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ApprovalItemApprovalItem) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ApprovalItemApprovalItem) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalNameNil

`func (o *ApprovalItemApprovalItem) SetExternalNameNil(b bool)`

 SetExternalNameNil sets the value for ExternalName to be an explicit nil

### UnsetExternalName
`func (o *ApprovalItemApprovalItem) UnsetExternalName()`

UnsetExternalName ensures that no value is present for ExternalName, not even an explicit nil
### GetInternalId

`func (o *ApprovalItemApprovalItem) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ApprovalItemApprovalItem) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ApprovalItemApprovalItem) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ApprovalItemApprovalItem) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ApprovalItemApprovalItem) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ApprovalItemApprovalItem) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetApprovedBy

`func (o *ApprovalItemApprovalItem) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *ApprovalItemApprovalItem) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *ApprovalItemApprovalItem) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *ApprovalItemApprovalItem) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetDeniedBy

`func (o *ApprovalItemApprovalItem) GetDeniedBy() string`

GetDeniedBy returns the DeniedBy field if non-nil, zero value otherwise.

### GetDeniedByOk

`func (o *ApprovalItemApprovalItem) GetDeniedByOk() (*string, bool)`

GetDeniedByOk returns a tuple with the DeniedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedBy

`func (o *ApprovalItemApprovalItem) SetDeniedBy(v string)`

SetDeniedBy sets DeniedBy field to given value.

### HasDeniedBy

`func (o *ApprovalItemApprovalItem) HasDeniedBy() bool`

HasDeniedBy returns a boolean if a field has been set.

### SetDeniedByNil

`func (o *ApprovalItemApprovalItem) SetDeniedByNil(b bool)`

 SetDeniedByNil sets the value for DeniedBy to be an explicit nil

### UnsetDeniedBy
`func (o *ApprovalItemApprovalItem) UnsetDeniedBy()`

UnsetDeniedBy ensures that no value is present for DeniedBy, not even an explicit nil
### GetStatus

`func (o *ApprovalItemApprovalItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalItemApprovalItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalItemApprovalItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApprovalItemApprovalItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApprovalItemApprovalItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApprovalItemApprovalItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApprovalItemApprovalItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApprovalItemApprovalItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ApprovalItemApprovalItem) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ApprovalItemApprovalItem) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDateCreated

`func (o *ApprovalItemApprovalItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApprovalItemApprovalItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApprovalItemApprovalItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApprovalItemApprovalItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ApprovalItemApprovalItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalItemApprovalItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalItemApprovalItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ApprovalItemApprovalItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateApproved

`func (o *ApprovalItemApprovalItem) GetDateApproved() time.Time`

GetDateApproved returns the DateApproved field if non-nil, zero value otherwise.

### GetDateApprovedOk

`func (o *ApprovalItemApprovalItem) GetDateApprovedOk() (*time.Time, bool)`

GetDateApprovedOk returns a tuple with the DateApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateApproved

`func (o *ApprovalItemApprovalItem) SetDateApproved(v time.Time)`

SetDateApproved sets DateApproved field to given value.

### HasDateApproved

`func (o *ApprovalItemApprovalItem) HasDateApproved() bool`

HasDateApproved returns a boolean if a field has been set.

### GetDateDenied

`func (o *ApprovalItemApprovalItem) GetDateDenied() time.Time`

GetDateDenied returns the DateDenied field if non-nil, zero value otherwise.

### GetDateDeniedOk

`func (o *ApprovalItemApprovalItem) GetDateDeniedOk() (*time.Time, bool)`

GetDateDeniedOk returns a tuple with the DateDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDenied

`func (o *ApprovalItemApprovalItem) SetDateDenied(v time.Time)`

SetDateDenied sets DateDenied field to given value.

### HasDateDenied

`func (o *ApprovalItemApprovalItem) HasDateDenied() bool`

HasDateDenied returns a boolean if a field has been set.

### SetDateDeniedNil

`func (o *ApprovalItemApprovalItem) SetDateDeniedNil(b bool)`

 SetDateDeniedNil sets the value for DateDenied to be an explicit nil

### UnsetDateDenied
`func (o *ApprovalItemApprovalItem) UnsetDateDenied()`

UnsetDateDenied ensures that no value is present for DateDenied, not even an explicit nil
### GetApproval

`func (o *ApprovalItemApprovalItem) GetApproval() ApprovalItemApprovalItemApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *ApprovalItemApprovalItem) GetApprovalOk() (*ApprovalItemApprovalItemApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *ApprovalItemApprovalItem) SetApproval(v ApprovalItemApprovalItemApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *ApprovalItemApprovalItem) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetReference

`func (o *ApprovalItemApprovalItem) GetReference() ApprovalItemApprovalItemReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApprovalItemApprovalItem) GetReferenceOk() (*ApprovalItemApprovalItemReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApprovalItemApprovalItem) SetReference(v ApprovalItemApprovalItemReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApprovalItemApprovalItem) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDetails

`func (o *ApprovalItemApprovalItem) GetDetails() []ApprovalItemApprovalItemDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApprovalItemApprovalItem) GetDetailsOk() (*[]ApprovalItemApprovalItemDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApprovalItemApprovalItem) SetDetails(v []ApprovalItemApprovalItemDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApprovalItemApprovalItem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ApprovalItemApprovalItem) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ApprovalItemApprovalItem) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


