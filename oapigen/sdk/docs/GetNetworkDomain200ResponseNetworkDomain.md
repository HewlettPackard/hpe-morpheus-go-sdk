# GetNetworkDomain200ResponseNetworkDomain

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
**Account** | Pointer to [**GetNetworkDomain200ResponseNetworkDomainAccount**](GetNetworkDomain200ResponseNetworkDomainAccount.md) |  | [optional] 
**Owner** | Pointer to [**GetNetworkDomain200ResponseNetworkDomainOwner**](GetNetworkDomain200ResponseNetworkDomainOwner.md) |  | [optional] 

## Methods

### NewGetNetworkDomain200ResponseNetworkDomain

`func NewGetNetworkDomain200ResponseNetworkDomain() *GetNetworkDomain200ResponseNetworkDomain`

NewGetNetworkDomain200ResponseNetworkDomain instantiates a new GetNetworkDomain200ResponseNetworkDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkDomain200ResponseNetworkDomainWithDefaults

`func NewGetNetworkDomain200ResponseNetworkDomainWithDefaults() *GetNetworkDomain200ResponseNetworkDomain`

NewGetNetworkDomain200ResponseNetworkDomainWithDefaults instantiates a new GetNetworkDomain200ResponseNetworkDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFqdn

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### SetFqdnNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetFqdnNil(b bool)`

 SetFqdnNil sets the value for Fqdn to be an explicit nil

### UnsetFqdn
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetFqdn()`

UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
### GetDescription

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDomainController

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainController() bool`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainControllerOk() (*bool, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainController(v bool)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetPublicZone

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetPublicZone() bool`

GetPublicZone returns the PublicZone field if non-nil, zero value otherwise.

### GetPublicZoneOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetPublicZoneOk() (*bool, bool)`

GetPublicZoneOk returns a tuple with the PublicZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicZone

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetPublicZone(v bool)`

SetPublicZone sets PublicZone field to given value.

### HasPublicZone

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasPublicZone() bool`

HasPublicZone returns a boolean if a field has been set.

### GetDomainUsername

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### SetDomainUsernameNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainUsernameNil(b bool)`

 SetDomainUsernameNil sets the value for DomainUsername to be an explicit nil

### UnsetDomainUsername
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDomainUsername()`

UnsetDomainUsername ensures that no value is present for DomainUsername, not even an explicit nil
### GetDomainPassword

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.

### SetDomainPasswordNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainPasswordNil(b bool)`

 SetDomainPasswordNil sets the value for DomainPassword to be an explicit nil

### UnsetDomainPassword
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDomainPassword()`

UnsetDomainPassword ensures that no value is present for DomainPassword, not even an explicit nil
### GetRefType

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetRefSource

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefSource() string`

GetRefSource returns the RefSource field if non-nil, zero value otherwise.

### GetRefSourceOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetRefSourceOk() (*string, bool)`

GetRefSourceOk returns a tuple with the RefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSource

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefSource(v string)`

SetRefSource sets RefSource field to given value.

### HasRefSource

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasRefSource() bool`

HasRefSource returns a boolean if a field has been set.

### SetRefSourceNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetRefSourceNil(b bool)`

 SetRefSourceNil sets the value for RefSource to be an explicit nil

### UnsetRefSource
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetRefSource()`

UnsetRefSource ensures that no value is present for RefSource, not even an explicit nil
### GetInternalId

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetOuPath

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### SetOuPathNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetOuPathNil(b bool)`

 SetOuPathNil sets the value for OuPath to be an explicit nil

### UnsetOuPath
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetOuPath()`

UnsetOuPath ensures that no value is present for OuPath, not even an explicit nil
### GetDcServer

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDcServer() string`

GetDcServer returns the DcServer field if non-nil, zero value otherwise.

### GetDcServerOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDcServerOk() (*string, bool)`

GetDcServerOk returns a tuple with the DcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcServer

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDcServer(v string)`

SetDcServer sets DcServer field to given value.

### HasDcServer

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDcServer() bool`

HasDcServer returns a boolean if a field has been set.

### SetDcServerNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDcServerNil(b bool)`

 SetDcServerNil sets the value for DcServer to be an explicit nil

### UnsetDcServer
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDcServer()`

UnsetDcServer ensures that no value is present for DcServer, not even an explicit nil
### GetZoneType

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### SetZoneTypeNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetZoneTypeNil(b bool)`

 SetZoneTypeNil sets the value for ZoneType to be an explicit nil

### UnsetZoneType
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetZoneType()`

UnsetZoneType ensures that no value is present for ZoneType, not even an explicit nil
### GetDnssec

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDnssec() string`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDnssecOk() (*string, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDnssec(v string)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### SetDnssecNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDnssecNil(b bool)`

 SetDnssecNil sets the value for Dnssec to be an explicit nil

### UnsetDnssec
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDnssec()`

UnsetDnssec ensures that no value is present for Dnssec, not even an explicit nil
### GetDomainSerial

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainSerial() string`

GetDomainSerial returns the DomainSerial field if non-nil, zero value otherwise.

### GetDomainSerialOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetDomainSerialOk() (*string, bool)`

GetDomainSerialOk returns a tuple with the DomainSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSerial

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainSerial(v string)`

SetDomainSerial sets DomainSerial field to given value.

### HasDomainSerial

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasDomainSerial() bool`

HasDomainSerial returns a boolean if a field has been set.

### SetDomainSerialNil

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetDomainSerialNil(b bool)`

 SetDomainSerialNil sets the value for DomainSerial to be an explicit nil

### UnsetDomainSerial
`func (o *GetNetworkDomain200ResponseNetworkDomain) UnsetDomainSerial()`

UnsetDomainSerial ensures that no value is present for DomainSerial, not even an explicit nil
### GetAccount

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetAccount() GetNetworkDomain200ResponseNetworkDomainAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetAccountOk() (*GetNetworkDomain200ResponseNetworkDomainAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetAccount(v GetNetworkDomain200ResponseNetworkDomainAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetOwner() GetNetworkDomain200ResponseNetworkDomainOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkDomain200ResponseNetworkDomain) GetOwnerOk() (*GetNetworkDomain200ResponseNetworkDomainOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkDomain200ResponseNetworkDomain) SetOwner(v GetNetworkDomain200ResponseNetworkDomainOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkDomain200ResponseNetworkDomain) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


