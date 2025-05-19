# ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner.md) |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions

`func NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions() *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions`

NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions instantiates a new ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissionsWithDefaults

`func NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissionsWithDefaults() *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions`

NewListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissionsWithDefaults instantiates a new ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetSites() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetSitesOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetSites(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ListServicePlans200ResponseAllOfServicePlansInnerPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


