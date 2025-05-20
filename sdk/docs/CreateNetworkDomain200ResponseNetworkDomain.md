# CreateNetworkDomain200ResponseNetworkDomain

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

### NewCreateNetworkDomain200ResponseNetworkDomain

`func NewCreateNetworkDomain200ResponseNetworkDomain() *CreateNetworkDomain200ResponseNetworkDomain`

NewCreateNetworkDomain200ResponseNetworkDomain instantiates a new CreateNetworkDomain200ResponseNetworkDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDomain200ResponseNetworkDomainWithDefaults

`func NewCreateNetworkDomain200ResponseNetworkDomainWithDefaults() *CreateNetworkDomain200ResponseNetworkDomain`

NewCreateNetworkDomain200ResponseNetworkDomainWithDefaults instantiates a new CreateNetworkDomain200ResponseNetworkDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetFqdn

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDomainController

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainController() bool`

GetDomainController returns the DomainController field if non-nil, zero value otherwise.

### GetDomainControllerOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainControllerOk() (*bool, bool)`

GetDomainControllerOk returns a tuple with the DomainController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainController

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDomainController(v bool)`

SetDomainController sets DomainController field to given value.

### HasDomainController

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDomainController() bool`

HasDomainController returns a boolean if a field has been set.

### GetPublicZone

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetPublicZone() bool`

GetPublicZone returns the PublicZone field if non-nil, zero value otherwise.

### GetPublicZoneOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetPublicZoneOk() (*bool, bool)`

GetPublicZoneOk returns a tuple with the PublicZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicZone

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetPublicZone(v bool)`

SetPublicZone sets PublicZone field to given value.

### HasPublicZone

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasPublicZone() bool`

HasPublicZone returns a boolean if a field has been set.

### GetDomainUsername

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### GetDomainPassword

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.

### GetRefType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefSource

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefSource() string`

GetRefSource returns the RefSource field if non-nil, zero value otherwise.

### GetRefSourceOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetRefSourceOk() (*string, bool)`

GetRefSourceOk returns a tuple with the RefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefSource

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetRefSource(v string)`

SetRefSource sets RefSource field to given value.

### HasRefSource

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasRefSource() bool`

HasRefSource returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetOuPath

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetOuPath() string`

GetOuPath returns the OuPath field if non-nil, zero value otherwise.

### GetOuPathOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetOuPathOk() (*string, bool)`

GetOuPathOk returns a tuple with the OuPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuPath

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetOuPath(v string)`

SetOuPath sets OuPath field to given value.

### HasOuPath

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasOuPath() bool`

HasOuPath returns a boolean if a field has been set.

### GetDcServer

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDcServer() string`

GetDcServer returns the DcServer field if non-nil, zero value otherwise.

### GetDcServerOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDcServerOk() (*string, bool)`

GetDcServerOk returns a tuple with the DcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcServer

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDcServer(v string)`

SetDcServer sets DcServer field to given value.

### HasDcServer

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDcServer() bool`

HasDcServer returns a boolean if a field has been set.

### GetZoneType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetDnssec

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDnssec() string`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDnssecOk() (*string, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDnssec(v string)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### GetDomainSerial

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainSerial() string`

GetDomainSerial returns the DomainSerial field if non-nil, zero value otherwise.

### GetDomainSerialOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetDomainSerialOk() (*string, bool)`

GetDomainSerialOk returns a tuple with the DomainSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSerial

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetDomainSerial(v string)`

SetDomainSerial sets DomainSerial field to given value.

### HasDomainSerial

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasDomainSerial() bool`

HasDomainSerial returns a boolean if a field has been set.

### GetAccount

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNetworkDomain200ResponseNetworkDomain) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNetworkDomain200ResponseNetworkDomain) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateNetworkDomain200ResponseNetworkDomain) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


