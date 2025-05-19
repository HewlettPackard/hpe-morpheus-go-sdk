# NetworkDomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**DisplayName** | Pointer to **string** | Overrides displayed name in domain selection components. Useful if using many OU Paths. | [optional] 
**PublicZone** | Pointer to **bool** | Public Zone | [optional] [default to false]
**TaskSetId** | Pointer to **int64** | Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.)  | [optional] 
**Active** | Pointer to **bool** | Active | [optional] 
**DomainController** | Pointer to **bool** | Join Domain Controller | [optional] [default to true]
**DomainUsername** | Pointer to **string** | Domain Username | [optional] 
**DomainPassword** | Pointer to **string** | Domain Password | [optional] 
**DcServer** | Pointer to **string** | DC Server. If specified, must be the server name and not an IP Address | [optional] 
**OuPath** | Pointer to **string** | OU Path. (i.e. &#39;OU&#x3D;staging,DC&#x3D;ad,DC&#x3D;yourdomain,DC&#x3D;com&#39;) | [optional] 
**GuestUsername** | Pointer to **string** | Guest Username. If set, will change the instances RPC Service User after joining a Domain. | [optional] 
**GuestPassword** | Pointer to **string** | Guest Password | [optional] 

## Methods

### NewNetworkDomainCreate

`func NewNetworkDomainCreate() *NetworkDomainCreate`

NewNetworkDomainCreate instantiates a new NetworkDomainCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDomainCreateWithDefaults

`func NewNetworkDomainCreateWithDefaults() *NetworkDomainCreate`

NewNetworkDomainCreateWithDefaults instantiates a new NetworkDomainCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkDomainCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDomainCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDomainCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDomainCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkDomainCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDomainCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDomainCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkDomainCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *NetworkDomainCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NetworkDomainCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NetworkDomainCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NetworkDomainCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPublicZone

`func (o *NetworkDomainCreate) GetPublicZone() bool`

GetPublicZone returns the PublicZone field if non-nil, zero value otherwise.

### GetPublicZoneOk

`func (o *NetworkDomainCreate) GetPublicZoneOk() (*bool, bool)`

GetPublicZoneOk returns a tuple with the PublicZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicZone

`func (o *NetworkDomainCreate) SetPublicZone(v bool)`

SetPublicZone sets PublicZone field to given value.

### HasPublicZone

`func (o *NetworkDomainCreate) HasPublicZone() bool`

HasPublicZone returns a boolean if a field has been set.

### GetTaskSetId

`func (o *NetworkDomainCreate) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *NetworkDomainCreate) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *NetworkDomainCreate) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *NetworkDomainCreate) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetActive

`func (o *NetworkDomainCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkDomainCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkDomainCreate) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkDomainCreate) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDomainController

`func (o *NetworkDomainCreate) GetDomainController() bool`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *NetworkDomainCreate) GetDomainControllerOk() (*bool, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *NetworkDomainCreate) SetDomainController(v bool)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *NetworkDomainCreate) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetDomainUsername

`func (o *NetworkDomainCreate) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *NetworkDomainCreate) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *NetworkDomainCreate) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *NetworkDomainCreate) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### GetDomainPassword

`func (o *NetworkDomainCreate) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *NetworkDomainCreate) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *NetworkDomainCreate) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *NetworkDomainCreate) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.

### GetDcServer

`func (o *NetworkDomainCreate) GetDcServer() string`

GetDcServer returns the DcServer field if non-nil, zero value otherwise.

### GetDcServerOk

`func (o *NetworkDomainCreate) GetDcServerOk() (*string, bool)`

GetDcServerOk returns a tuple with the DcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcServer

`func (o *NetworkDomainCreate) SetDcServer(v string)`

SetDcServer sets DcServer field to given value.

### HasDcServer

`func (o *NetworkDomainCreate) HasDcServer() bool`

HasDcServer returns a boolean if a field has been set.

### GetOuPath

`func (o *NetworkDomainCreate) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *NetworkDomainCreate) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *NetworkDomainCreate) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *NetworkDomainCreate) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### GetGuestUsername

`func (o *NetworkDomainCreate) GetGuestUsername() string`

GetGuestUsername returns the GuestUsername field if non-nil, zero value otherwise.

### GetGuestUsernameOk

`func (o *NetworkDomainCreate) GetGuestUsernameOk() (*string, bool)`

GetGuestUsernameOk returns a tuple with the GuestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUsername

`func (o *NetworkDomainCreate) SetGuestUsername(v string)`

SetGuestUsername sets GuestUsername field to given value.

### HasGuestUsername

`func (o *NetworkDomainCreate) HasGuestUsername() bool`

HasGuestUsername returns a boolean if a field has been set.

### GetGuestPassword

`func (o *NetworkDomainCreate) GetGuestPassword() string`

GetGuestPassword returns the GuestPassword field if non-nil, zero value otherwise.

### GetGuestPasswordOk

`func (o *NetworkDomainCreate) GetGuestPasswordOk() (*string, bool)`

GetGuestPasswordOk returns a tuple with the GuestPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPassword

`func (o *NetworkDomainCreate) SetGuestPassword(v string)`

SetGuestPassword sets GuestPassword field to given value.

### HasGuestPassword

`func (o *NetworkDomainCreate) HasGuestPassword() bool`

HasGuestPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


