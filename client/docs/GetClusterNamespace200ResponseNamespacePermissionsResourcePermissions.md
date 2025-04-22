# GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllGroups** | Pointer to **bool** |  | [optional] 
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**MorpheusResourceType** | Pointer to **string** |  | [optional] 
**MorpheusResourceId** | Pointer to **int64** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Sites** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner.md) |  | [optional] 
**Plans** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner.md) |  | [optional] 

## Methods

### NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissions

`func NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissions() *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions`

NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissions instantiates a new GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissionsWithDefaults

`func NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissionsWithDefaults() *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions`

NewGetClusterNamespace200ResponseNamespacePermissionsResourcePermissionsWithDefaults instantiates a new GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllGroups

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.

### HasAllGroups

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasAllGroups() bool`

HasAllGroups returns a boolean if a field has been set.

### GetDefaultStore

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetMorpheusResourceType

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetMorpheusResourceType() string`

GetMorpheusResourceType returns the MorpheusResourceType field if non-nil, zero value otherwise.

### GetMorpheusResourceTypeOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetMorpheusResourceTypeOk() (*string, bool)`

GetMorpheusResourceTypeOk returns a tuple with the MorpheusResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceType

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetMorpheusResourceType(v string)`

SetMorpheusResourceType sets MorpheusResourceType field to given value.

### HasMorpheusResourceType

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasMorpheusResourceType() bool`

HasMorpheusResourceType returns a boolean if a field has been set.

### GetMorpheusResourceId

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetMorpheusResourceId() int64`

GetMorpheusResourceId returns the MorpheusResourceId field if non-nil, zero value otherwise.

### GetMorpheusResourceIdOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetMorpheusResourceIdOk() (*int64, bool)`

GetMorpheusResourceIdOk returns a tuple with the MorpheusResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorpheusResourceId

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetMorpheusResourceId(v int64)`

SetMorpheusResourceId sets MorpheusResourceId field to given value.

### HasMorpheusResourceId

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasMorpheusResourceId() bool`

HasMorpheusResourceId returns a boolean if a field has been set.

### GetCanManage

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetSites() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetSitesOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetSites(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetPlans() []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) GetPlansOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) SetPlans(v []ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermissionSitesInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *GetClusterNamespace200ResponseNamespacePermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


