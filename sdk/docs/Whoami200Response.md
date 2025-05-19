# Whoami200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**AddUserTenant200ResponseAllOfUser**](AddUserTenant200ResponseAllOfUser.md) |  | [optional] 
**IsMasterAccount** | Pointer to **bool** |  | [optional] 
**Permissions** | Pointer to [**[]Whoami200ResponsePermissionsInner**](Whoami200ResponsePermissionsInner.md) |  | [optional] 
**Appliance** | Pointer to [**Whoami200ResponseAppliance**](Whoami200ResponseAppliance.md) |  | [optional] 

## Methods

### NewWhoami200Response

`func NewWhoami200Response() *Whoami200Response`

NewWhoami200Response instantiates a new Whoami200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoami200ResponseWithDefaults

`func NewWhoami200ResponseWithDefaults() *Whoami200Response`

NewWhoami200ResponseWithDefaults instantiates a new Whoami200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *Whoami200Response) GetUser() AddUserTenant200ResponseAllOfUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Whoami200Response) GetUserOk() (*AddUserTenant200ResponseAllOfUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Whoami200Response) SetUser(v AddUserTenant200ResponseAllOfUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Whoami200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIsMasterAccount

`func (o *Whoami200Response) GetIsMasterAccount() bool`

GetIsMasterAccount returns the IsMasterAccount field if non-nil, zero value otherwise.

### GetIsMasterAccountOk

`func (o *Whoami200Response) GetIsMasterAccountOk() (*bool, bool)`

GetIsMasterAccountOk returns a tuple with the IsMasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMasterAccount

`func (o *Whoami200Response) SetIsMasterAccount(v bool)`

SetIsMasterAccount sets IsMasterAccount field to given value.

### HasIsMasterAccount

`func (o *Whoami200Response) HasIsMasterAccount() bool`

HasIsMasterAccount returns a boolean if a field has been set.

### GetPermissions

`func (o *Whoami200Response) GetPermissions() []Whoami200ResponsePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Whoami200Response) GetPermissionsOk() (*[]Whoami200ResponsePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Whoami200Response) SetPermissions(v []Whoami200ResponsePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Whoami200Response) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAppliance

`func (o *Whoami200Response) GetAppliance() Whoami200ResponseAppliance`

GetAppliance returns the Appliance field if non-nil, zero value otherwise.

### GetApplianceOk

`func (o *Whoami200Response) GetApplianceOk() (*Whoami200ResponseAppliance, bool)`

GetApplianceOk returns a tuple with the Appliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliance

`func (o *Whoami200Response) SetAppliance(v Whoami200ResponseAppliance)`

SetAppliance sets Appliance field to given value.

### HasAppliance

`func (o *Whoami200Response) HasAppliance() bool`

HasAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


