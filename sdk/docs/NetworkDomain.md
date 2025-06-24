# NetworkDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Fqdn** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**DomainController** | Pointer to **bool** |  | [optional] 
**PublicZone** | Pointer to **bool** |  | [optional] 
**DomainUsername** | Pointer to **NullableString** |  | [optional] 
**DomainPassword** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**RefSource** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**OuPath** | Pointer to **NullableString** |  | [optional] 
**DcServer** | Pointer to **NullableString** |  | [optional] 
**ZoneType** | Pointer to **NullableString** |  | [optional] 
**Dnssec** | Pointer to **NullableString** |  | [optional] 
**DomainSerial** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewNetworkDomain

`func NewNetworkDomain() *NetworkDomain`

NewNetworkDomain instantiates a new NetworkDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDomainWithDefaults

`func NewNetworkDomainWithDefaults() *NetworkDomain`

NewNetworkDomainWithDefaults instantiates a new NetworkDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkDomain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkDomain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkDomain) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *NetworkDomain) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkDomain) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkDomain) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkDomain) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFqdn

`func (o *NetworkDomain) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NetworkDomain) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NetworkDomain) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NetworkDomain) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *NetworkDomain) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *NetworkDomain) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetDescription

`func (o *NetworkDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkDomain) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkDomain) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *NetworkDomain) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkDomain) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkDomain) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkDomain) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDomainController

`func (o *NetworkDomain) GetDomainController() bool`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *NetworkDomain) GetDomainControllerOk() (*bool, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *NetworkDomain) SetDomainController(v bool)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *NetworkDomain) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetPublicZone

`func (o *NetworkDomain) GetPublicZone() bool`

GetPublicZone returns the PublicZone field if non-nil, zero value otherwise.

### GetPublicZoneOk

`func (o *NetworkDomain) GetPublicZoneOk() (*bool, bool)`

GetPublicZoneOk returns a tuple with the PublicZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicZone

`func (o *NetworkDomain) SetPublicZone(v bool)`

SetPublicZone sets PublicZone field to given value.

### HasPublicZone

`func (o *NetworkDomain) HasPublicZone() bool`

HasPublicZone returns a boolean if a field has been set.

### GetDomainUsername

`func (o *NetworkDomain) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *NetworkDomain) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *NetworkDomain) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *NetworkDomain) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### SetDomainUsernameNil

`func (o *NetworkDomain) SetDomainUsernameNil(b bool)`

 SetDomainUsernameNil sets the value for DomainUsername to be an explicit nil

### UnsetDomainUsername
`func (o *NetworkDomain) UnsetDomainUsername()`

UnsetDomainUsername ensures that no value is present for DomainUsername, not even an explicit nil
### GetDomainPassword

`func (o *NetworkDomain) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *NetworkDomain) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *NetworkDomain) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *NetworkDomain) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.

### SetDomainPasswordNil

`func (o *NetworkDomain) SetDomainPasswordNil(b bool)`

 SetDomainPasswordNil sets the value for DomainPassword to be an explicit nil

### UnsetDomainPassword
`func (o *NetworkDomain) UnsetDomainPassword()`

UnsetDomainPassword ensures that no value is present for DomainPassword, not even an explicit nil
### GetRefType

`func (o *NetworkDomain) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *NetworkDomain) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *NetworkDomain) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *NetworkDomain) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *NetworkDomain) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *NetworkDomain) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *NetworkDomain) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *NetworkDomain) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *NetworkDomain) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *NetworkDomain) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *NetworkDomain) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *NetworkDomain) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetRefSource

`func (o *NetworkDomain) GetRefSource() string`

GetRefSource returns the RefSource field if non-nil, zero value otherwise.

### GetRefSourceOk

`func (o *NetworkDomain) GetRefSourceOk() (*string, bool)`

GetRefSourceOk returns a tuple with the RefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSource

`func (o *NetworkDomain) SetRefSource(v string)`

SetRefSource sets RefSource field to given value.

### HasRefSource

`func (o *NetworkDomain) HasRefSource() bool`

HasRefSource returns a boolean if a field has been set.

### SetRefSourceNil

`func (o *NetworkDomain) SetRefSourceNil(b bool)`

 SetRefSourceNil sets the value for RefSource to be an explicit nil

### UnsetRefSource
`func (o *NetworkDomain) UnsetRefSource()`

UnsetRefSource ensures that no value is present for RefSource, not even an explicit nil
### GetInternalId

`func (o *NetworkDomain) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkDomain) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkDomain) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkDomain) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *NetworkDomain) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *NetworkDomain) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetOuPath

`func (o *NetworkDomain) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *NetworkDomain) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *NetworkDomain) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *NetworkDomain) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### SetOuPathNil

`func (o *NetworkDomain) SetOuPathNil(b bool)`

 SetOuPathNil sets the value for OuPath to be an explicit nil

### UnsetOuPath
`func (o *NetworkDomain) UnsetOuPath()`

UnsetOuPath ensures that no value is present for OuPath, not even an explicit nil
### GetDcServer

`func (o *NetworkDomain) GetDcServer() string`

GetDcServer returns the DcServer field if non-nil, zero value otherwise.

### GetDcServerOk

`func (o *NetworkDomain) GetDcServerOk() (*string, bool)`

GetDcServerOk returns a tuple with the DcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcServer

`func (o *NetworkDomain) SetDcServer(v string)`

SetDcServer sets DcServer field to given value.

### HasDcServer

`func (o *NetworkDomain) HasDcServer() bool`

HasDcServer returns a boolean if a field has been set.

### SetDcServerNil

`func (o *NetworkDomain) SetDcServerNil(b bool)`

 SetDcServerNil sets the value for DcServer to be an explicit nil

### UnsetDcServer
`func (o *NetworkDomain) UnsetDcServer()`

UnsetDcServer ensures that no value is present for DcServer, not even an explicit nil
### GetZoneType

`func (o *NetworkDomain) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *NetworkDomain) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *NetworkDomain) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *NetworkDomain) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### SetZoneTypeNil

`func (o *NetworkDomain) SetZoneTypeNil(b bool)`

 SetZoneTypeNil sets the value for ZoneType to be an explicit nil

### UnsetZoneType
`func (o *NetworkDomain) UnsetZoneType()`

UnsetZoneType ensures that no value is present for ZoneType, not even an explicit nil
### GetDnssec

`func (o *NetworkDomain) GetDnssec() string`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *NetworkDomain) GetDnssecOk() (*string, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *NetworkDomain) SetDnssec(v string)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *NetworkDomain) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### SetDnssecNil

`func (o *NetworkDomain) SetDnssecNil(b bool)`

 SetDnssecNil sets the value for Dnssec to be an explicit nil

### UnsetDnssec
`func (o *NetworkDomain) UnsetDnssec()`

UnsetDnssec ensures that no value is present for Dnssec, not even an explicit nil
### GetDomainSerial

`func (o *NetworkDomain) GetDomainSerial() string`

GetDomainSerial returns the DomainSerial field if non-nil, zero value otherwise.

### GetDomainSerialOk

`func (o *NetworkDomain) GetDomainSerialOk() (*string, bool)`

GetDomainSerialOk returns a tuple with the DomainSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSerial

`func (o *NetworkDomain) SetDomainSerial(v string)`

SetDomainSerial sets DomainSerial field to given value.

### HasDomainSerial

`func (o *NetworkDomain) HasDomainSerial() bool`

HasDomainSerial returns a boolean if a field has been set.

### SetDomainSerialNil

`func (o *NetworkDomain) SetDomainSerialNil(b bool)`

 SetDomainSerialNil sets the value for DomainSerial to be an explicit nil

### UnsetDomainSerial
`func (o *NetworkDomain) UnsetDomainSerial()`

UnsetDomainSerial ensures that no value is present for DomainSerial, not even an explicit nil
### GetAccount

`func (o *NetworkDomain) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkDomain) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkDomain) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkDomain) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *NetworkDomain) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NetworkDomain) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NetworkDomain) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NetworkDomain) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


