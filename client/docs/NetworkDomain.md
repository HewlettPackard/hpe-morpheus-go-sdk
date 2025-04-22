# NetworkDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**DomainController** | Pointer to **bool** |  | [optional] 
**PublicZone** | Pointer to **bool** |  | [optional] 
**DomainUsername** | Pointer to **string** |  | [optional] 
**DomainPassword** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefSource** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**OuPath** | Pointer to **string** |  | [optional] 
**DcServer** | Pointer to **string** |  | [optional] 
**ZoneType** | Pointer to **string** |  | [optional] 
**Dnssec** | Pointer to **string** |  | [optional] 
**DomainSerial** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

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

### GetAccount

`func (o *NetworkDomain) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkDomain) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkDomain) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkDomain) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *NetworkDomain) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NetworkDomain) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NetworkDomain) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NetworkDomain) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


