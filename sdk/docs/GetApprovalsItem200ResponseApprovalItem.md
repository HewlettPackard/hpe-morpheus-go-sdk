# GetApprovalsItem200ResponseApprovalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ApprovedBy** | Pointer to **string** |  | [optional] 
**DeniedBy** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**DateApproved** | Pointer to **time.Time** |  | [optional] 
**DateDenied** | Pointer to **time.Time** |  | [optional] 
**Approval** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Reference** | Pointer to [**GetApprovalsItem200ResponseApprovalItemReference**](GetApprovalsItem200ResponseApprovalItemReference.md) |  | [optional] 

## Methods

### NewGetApprovalsItem200ResponseApprovalItem

`func NewGetApprovalsItem200ResponseApprovalItem() *GetApprovalsItem200ResponseApprovalItem`

NewGetApprovalsItem200ResponseApprovalItem instantiates a new GetApprovalsItem200ResponseApprovalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApprovalsItem200ResponseApprovalItemWithDefaults

`func NewGetApprovalsItem200ResponseApprovalItemWithDefaults() *GetApprovalsItem200ResponseApprovalItem`

NewGetApprovalsItem200ResponseApprovalItemWithDefaults instantiates a new GetApprovalsItem200ResponseApprovalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApprovalsItem200ResponseApprovalItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApprovalsItem200ResponseApprovalItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetApprovalsItem200ResponseApprovalItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetApprovalsItem200ResponseApprovalItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApprovalsItem200ResponseApprovalItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApprovalsItem200ResponseApprovalItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalName

`func (o *GetApprovalsItem200ResponseApprovalItem) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetApprovalsItem200ResponseApprovalItem) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetApprovalsItem200ResponseApprovalItem) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetInternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetApprovalsItem200ResponseApprovalItem) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetApprovedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) GetApprovedBy() string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetApprovedByOk() (*string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) SetApprovedBy(v string)`

SetApprovedBy sets ApprovedBy field to given value.

### HasApprovedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) HasApprovedBy() bool`

HasApprovedBy returns a boolean if a field has been set.

### GetDeniedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDeniedBy() string`

GetDeniedBy returns the DeniedBy field if non-nil, zero value otherwise.

### GetDeniedByOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDeniedByOk() (*string, bool)`

GetDeniedByOk returns a tuple with the DeniedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeniedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) SetDeniedBy(v string)`

SetDeniedBy sets DeniedBy field to given value.

### HasDeniedBy

`func (o *GetApprovalsItem200ResponseApprovalItem) HasDeniedBy() bool`

HasDeniedBy returns a boolean if a field has been set.

### GetStatus

`func (o *GetApprovalsItem200ResponseApprovalItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApprovalsItem200ResponseApprovalItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApprovalsItem200ResponseApprovalItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetApprovalsItem200ResponseApprovalItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetApprovalsItem200ResponseApprovalItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetApprovalsItem200ResponseApprovalItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetApprovalsItem200ResponseApprovalItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetApprovalsItem200ResponseApprovalItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetApprovalsItem200ResponseApprovalItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetApprovalsItem200ResponseApprovalItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetApprovalsItem200ResponseApprovalItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetDateApproved

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateApproved() time.Time`

GetDateApproved returns the DateApproved field if non-nil, zero value otherwise.

### GetDateApprovedOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateApprovedOk() (*time.Time, bool)`

GetDateApprovedOk returns a tuple with the DateApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateApproved

`func (o *GetApprovalsItem200ResponseApprovalItem) SetDateApproved(v time.Time)`

SetDateApproved sets DateApproved field to given value.

### HasDateApproved

`func (o *GetApprovalsItem200ResponseApprovalItem) HasDateApproved() bool`

HasDateApproved returns a boolean if a field has been set.

### GetDateDenied

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateDenied() time.Time`

GetDateDenied returns the DateDenied field if non-nil, zero value otherwise.

### GetDateDeniedOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetDateDeniedOk() (*time.Time, bool)`

GetDateDeniedOk returns a tuple with the DateDenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDenied

`func (o *GetApprovalsItem200ResponseApprovalItem) SetDateDenied(v time.Time)`

SetDateDenied sets DateDenied field to given value.

### HasDateDenied

`func (o *GetApprovalsItem200ResponseApprovalItem) HasDateDenied() bool`

HasDateDenied returns a boolean if a field has been set.

### GetApproval

`func (o *GetApprovalsItem200ResponseApprovalItem) GetApproval() GetAlerts200ResponseAllOfChecksInnerAccount`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetApprovalOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *GetApprovalsItem200ResponseApprovalItem) SetApproval(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *GetApprovalsItem200ResponseApprovalItem) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### GetReference

`func (o *GetApprovalsItem200ResponseApprovalItem) GetReference() GetApprovalsItem200ResponseApprovalItemReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetApprovalsItem200ResponseApprovalItem) GetReferenceOk() (*GetApprovalsItem200ResponseApprovalItemReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetApprovalsItem200ResponseApprovalItem) SetReference(v GetApprovalsItem200ResponseApprovalItemReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetApprovalsItem200ResponseApprovalItem) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


