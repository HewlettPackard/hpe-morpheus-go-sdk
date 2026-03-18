# CreateNetworkDomainRequestNetworkDomain

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

### NewCreateNetworkDomainRequestNetworkDomain

`func NewCreateNetworkDomainRequestNetworkDomain() *CreateNetworkDomainRequestNetworkDomain`

NewCreateNetworkDomainRequestNetworkDomain instantiates a new CreateNetworkDomainRequestNetworkDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDomainRequestNetworkDomainWithDefaults

`func NewCreateNetworkDomainRequestNetworkDomainWithDefaults() *CreateNetworkDomainRequestNetworkDomain`

NewCreateNetworkDomainRequestNetworkDomainWithDefaults instantiates a new CreateNetworkDomainRequestNetworkDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkDomainRequestNetworkDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkDomainRequestNetworkDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkDomainRequestNetworkDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPublicZone

`func (o *CreateNetworkDomainRequestNetworkDomain) GetPublicZone() bool`

GetPublicZone returns the PublicZone field if non-nil, zero value otherwise.

### GetPublicZoneOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetPublicZoneOk() (*bool, bool)`

GetPublicZoneOk returns a tuple with the PublicZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicZone

`func (o *CreateNetworkDomainRequestNetworkDomain) SetPublicZone(v bool)`

SetPublicZone sets PublicZone field to given value.

### HasPublicZone

`func (o *CreateNetworkDomainRequestNetworkDomain) HasPublicZone() bool`

HasPublicZone returns a boolean if a field has been set.

### GetTaskSetId

`func (o *CreateNetworkDomainRequestNetworkDomain) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *CreateNetworkDomainRequestNetworkDomain) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *CreateNetworkDomainRequestNetworkDomain) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetActive

`func (o *CreateNetworkDomainRequestNetworkDomain) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNetworkDomainRequestNetworkDomain) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNetworkDomainRequestNetworkDomain) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDomainController

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainController() bool`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainControllerOk() (*bool, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDomainController(v bool)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetDomainUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### GetDomainPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.

### GetDcServer

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDcServer() string`

GetDcServer returns the DcServer field if non-nil, zero value otherwise.

### GetDcServerOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetDcServerOk() (*string, bool)`

GetDcServerOk returns a tuple with the DcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcServer

`func (o *CreateNetworkDomainRequestNetworkDomain) SetDcServer(v string)`

SetDcServer sets DcServer field to given value.

### HasDcServer

`func (o *CreateNetworkDomainRequestNetworkDomain) HasDcServer() bool`

HasDcServer returns a boolean if a field has been set.

### GetOuPath

`func (o *CreateNetworkDomainRequestNetworkDomain) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *CreateNetworkDomainRequestNetworkDomain) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *CreateNetworkDomainRequestNetworkDomain) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### GetGuestUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) GetGuestUsername() string`

GetGuestUsername returns the GuestUsername field if non-nil, zero value otherwise.

### GetGuestUsernameOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetGuestUsernameOk() (*string, bool)`

GetGuestUsernameOk returns a tuple with the GuestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) SetGuestUsername(v string)`

SetGuestUsername sets GuestUsername field to given value.

### HasGuestUsername

`func (o *CreateNetworkDomainRequestNetworkDomain) HasGuestUsername() bool`

HasGuestUsername returns a boolean if a field has been set.

### GetGuestPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) GetGuestPassword() string`

GetGuestPassword returns the GuestPassword field if non-nil, zero value otherwise.

### GetGuestPasswordOk

`func (o *CreateNetworkDomainRequestNetworkDomain) GetGuestPasswordOk() (*string, bool)`

GetGuestPasswordOk returns a tuple with the GuestPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) SetGuestPassword(v string)`

SetGuestPassword sets GuestPassword field to given value.

### HasGuestPassword

`func (o *CreateNetworkDomainRequestNetworkDomain) HasGuestPassword() bool`

HasGuestPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


