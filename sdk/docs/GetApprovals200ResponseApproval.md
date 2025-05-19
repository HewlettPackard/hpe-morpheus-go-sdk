# GetApprovals200ResponseApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalName** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Approver** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**AccountIntegration** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RequestBy** | Pointer to **string** |  | [optional] 
**ApprovalItems** | Pointer to [**[]GetApprovalsItem200ResponseApprovalItem**](GetApprovalsItem200ResponseApprovalItem.md) |  | [optional] 

## Methods

### NewGetApprovals200ResponseApproval

`func NewGetApprovals200ResponseApproval() *GetApprovals200ResponseApproval`

NewGetApprovals200ResponseApproval instantiates a new GetApprovals200ResponseApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApprovals200ResponseApprovalWithDefaults

`func NewGetApprovals200ResponseApprovalWithDefaults() *GetApprovals200ResponseApproval`

NewGetApprovals200ResponseApprovalWithDefaults instantiates a new GetApprovals200ResponseApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApprovals200ResponseApproval) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApprovals200ResponseApproval) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApprovals200ResponseApproval) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetApprovals200ResponseApproval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetApprovals200ResponseApproval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApprovals200ResponseApproval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApprovals200ResponseApproval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApprovals200ResponseApproval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *GetApprovals200ResponseApproval) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetApprovals200ResponseApproval) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetApprovals200ResponseApproval) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetApprovals200ResponseApproval) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetApprovals200ResponseApproval) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetApprovals200ResponseApproval) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetApprovals200ResponseApproval) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetApprovals200ResponseApproval) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalName

`func (o *GetApprovals200ResponseApproval) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *GetApprovals200ResponseApproval) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *GetApprovals200ResponseApproval) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *GetApprovals200ResponseApproval) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetRequestType

`func (o *GetApprovals200ResponseApproval) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *GetApprovals200ResponseApproval) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *GetApprovals200ResponseApproval) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *GetApprovals200ResponseApproval) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetAccount

`func (o *GetApprovals200ResponseApproval) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetApprovals200ResponseApproval) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetApprovals200ResponseApproval) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetApprovals200ResponseApproval) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApprover

`func (o *GetApprovals200ResponseApproval) GetApprover() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *GetApprovals200ResponseApproval) GetApproverOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *GetApprovals200ResponseApproval) SetApprover(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *GetApprovals200ResponseApproval) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetAccountIntegration

`func (o *GetApprovals200ResponseApproval) GetAccountIntegration() string`

GetAccountIntegration returns the AccountIntegration field if non-nil, zero value otherwise.

### GetAccountIntegrationOk

`func (o *GetApprovals200ResponseApproval) GetAccountIntegrationOk() (*string, bool)`

GetAccountIntegrationOk returns a tuple with the AccountIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegration

`func (o *GetApprovals200ResponseApproval) SetAccountIntegration(v string)`

SetAccountIntegration sets AccountIntegration field to given value.

### HasAccountIntegration

`func (o *GetApprovals200ResponseApproval) HasAccountIntegration() bool`

HasAccountIntegration returns a boolean if a field has been set.

### GetStatus

`func (o *GetApprovals200ResponseApproval) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetApprovals200ResponseApproval) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetApprovals200ResponseApproval) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetApprovals200ResponseApproval) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetApprovals200ResponseApproval) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetApprovals200ResponseApproval) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetApprovals200ResponseApproval) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetApprovals200ResponseApproval) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetApprovals200ResponseApproval) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetApprovals200ResponseApproval) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetApprovals200ResponseApproval) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetApprovals200ResponseApproval) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetApprovals200ResponseApproval) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetApprovals200ResponseApproval) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetApprovals200ResponseApproval) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetApprovals200ResponseApproval) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRequestBy

`func (o *GetApprovals200ResponseApproval) GetRequestBy() string`

GetRequestBy returns the RequestBy field if non-nil, zero value otherwise.

### GetRequestByOk

`func (o *GetApprovals200ResponseApproval) GetRequestByOk() (*string, bool)`

GetRequestByOk returns a tuple with the RequestBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBy

`func (o *GetApprovals200ResponseApproval) SetRequestBy(v string)`

SetRequestBy sets RequestBy field to given value.

### HasRequestBy

`func (o *GetApprovals200ResponseApproval) HasRequestBy() bool`

HasRequestBy returns a boolean if a field has been set.

### GetApprovalItems

`func (o *GetApprovals200ResponseApproval) GetApprovalItems() []GetApprovalsItem200ResponseApprovalItem`

GetApprovalItems returns the ApprovalItems field if non-nil, zero value otherwise.

### GetApprovalItemsOk

`func (o *GetApprovals200ResponseApproval) GetApprovalItemsOk() (*[]GetApprovalsItem200ResponseApprovalItem, bool)`

GetApprovalItemsOk returns a tuple with the ApprovalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalItems

`func (o *GetApprovals200ResponseApproval) SetApprovalItems(v []GetApprovalsItem200ResponseApprovalItem)`

SetApprovalItems sets ApprovalItems field to given value.

### HasApprovalItems

`func (o *GetApprovals200ResponseApproval) HasApprovalItems() bool`

HasApprovalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


