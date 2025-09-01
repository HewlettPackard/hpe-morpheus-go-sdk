# AddServicePlansRequestServicePlanPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllSites** | Pointer to **bool** | Gives all groups (sites) permission to Provision and Reconfigure with servicePlan. | [optional] [default to false]
**Sites** | Pointer to [**[]AddServicePlansRequestServicePlanPermissionsResourcePermissionsSitesInner**](AddServicePlansRequestServicePlanPermissionsResourcePermissionsSitesInner.md) | Group (site) permissions that determines what Groups the Service Plan will be available in for Provisioning and Reconfigure. | [optional] 

## Methods

### NewAddServicePlansRequestServicePlanPermissionsResourcePermissions

`func NewAddServicePlansRequestServicePlanPermissionsResourcePermissions() *AddServicePlansRequestServicePlanPermissionsResourcePermissions`

NewAddServicePlansRequestServicePlanPermissionsResourcePermissions instantiates a new AddServicePlansRequestServicePlanPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlansRequestServicePlanPermissionsResourcePermissionsWithDefaults

`func NewAddServicePlansRequestServicePlanPermissionsResourcePermissionsWithDefaults() *AddServicePlansRequestServicePlanPermissionsResourcePermissions`

NewAddServicePlansRequestServicePlanPermissionsResourcePermissionsWithDefaults instantiates a new AddServicePlansRequestServicePlanPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) GetAllSites() bool`

GetAllSites returns the AllSites field if non-nil, zero value otherwise.

### GetAllSitesOk

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) GetAllSitesOk() (*bool, bool)`

GetAllSitesOk returns a tuple with the AllSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) SetAllSites(v bool)`

SetAllSites sets AllSites field to given value.

### HasAllSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) HasAllSites() bool`

HasAllSites returns a boolean if a field has been set.

### GetSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) GetSites() []AddServicePlansRequestServicePlanPermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) GetSitesOk() (*[]AddServicePlansRequestServicePlanPermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) SetSites(v []AddServicePlansRequestServicePlanPermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddServicePlansRequestServicePlanPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


