/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the ListHosts200ResponseAllOfServersInnerInterfacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHosts200ResponseAllOfServersInnerInterfacesInner{}

// ListHosts200ResponseAllOfServersInnerInterfacesInner struct for ListHosts200ResponseAllOfServersInnerInterfacesInner
type ListHosts200ResponseAllOfServersInnerInterfacesInner struct {
	Id                   *int64                                                                  `json:"id,omitempty"`
	RefType              *string                                                                 `json:"refType,omitempty"`
	RefId                *string                                                                 `json:"refId,omitempty"`
	Name                 *string                                                                 `json:"name,omitempty"`
	InternalId           *string                                                                 `json:"internalId,omitempty"`
	ExternalId           *string                                                                 `json:"externalId,omitempty"`
	UniqueId             *string                                                                 `json:"uniqueId,omitempty"`
	PublicIpAddress      *string                                                                 `json:"publicIpAddress,omitempty"`
	PublicIpv6Address    *string                                                                 `json:"publicIpv6Address,omitempty"`
	IpAddress            *string                                                                 `json:"ipAddress,omitempty"`
	Ipv6Address          *string                                                                 `json:"ipv6Address,omitempty"`
	IpSubnet             *string                                                                 `json:"ipSubnet,omitempty"`
	Ipv6Subnet           *string                                                                 `json:"ipv6Subnet,omitempty"`
	Description          *string                                                                 `json:"description,omitempty"`
	Dhcp                 *bool                                                                   `json:"dhcp,omitempty"`
	Active               *bool                                                                   `json:"active,omitempty"`
	PoolAssigned         *bool                                                                   `json:"poolAssigned,omitempty"`
	PrimaryInterface     *bool                                                                   `json:"primaryInterface,omitempty"`
	Network              *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"network,omitempty"`
	Subnet               *string                                                                 `json:"subnet,omitempty"`
	NetworkGroup         *string                                                                 `json:"networkGroup,omitempty"`
	NetworkPosition      *string                                                                 `json:"networkPosition,omitempty"`
	NetworkPool          *GetAlerts200ResponseAllOfCheckGroupsInnerInstance                      `json:"networkPool,omitempty"`
	NetworkDomain        *string                                                                 `json:"networkDomain,omitempty"`
	Type                 *ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType       `json:"type,omitempty"`
	IpMode               *string                                                                 `json:"ipMode,omitempty"`
	MacAddress           *string                                                                 `json:"macAddress,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _ListHosts200ResponseAllOfServersInnerInterfacesInner ListHosts200ResponseAllOfServersInnerInterfacesInner

// NewListHosts200ResponseAllOfServersInnerInterfacesInner instantiates a new ListHosts200ResponseAllOfServersInnerInterfacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHosts200ResponseAllOfServersInnerInterfacesInner() *ListHosts200ResponseAllOfServersInnerInterfacesInner {
	this := ListHosts200ResponseAllOfServersInnerInterfacesInner{}
	return &this
}

// NewListHosts200ResponseAllOfServersInnerInterfacesInnerWithDefaults instantiates a new ListHosts200ResponseAllOfServersInnerInterfacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHosts200ResponseAllOfServersInnerInterfacesInnerWithDefaults() *ListHosts200ResponseAllOfServersInnerInterfacesInner {
	this := ListHosts200ResponseAllOfServersInnerInterfacesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetId(v int64) {
	o.Id = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// IsSetRefType returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// IsSetRefId returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetRefId(v string) {
	o.RefId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetName(v string) {
	o.Name = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// IsSetInternalId returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetInternalId(v string) {
	o.InternalId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// IsSetUniqueId returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpAddress() string {
	if o == nil || IsNil(o.PublicIpAddress) {
		var ret string
		return ret
	}
	return *o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIpAddress) {
		return nil, false
	}
	return o.PublicIpAddress, true
}

// IsSetPublicIpAddress returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetPublicIpAddress() bool {
	if o != nil && !IsNil(o.PublicIpAddress) {
		return true
	}

	return false
}

// SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpAddress(v string) {
	o.PublicIpAddress = &v
}

// GetPublicIpv6Address returns the PublicIpv6Address field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpv6Address() string {
	if o == nil || IsNil(o.PublicIpv6Address) {
		var ret string
		return ret
	}
	return *o.PublicIpv6Address
}

// GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPublicIpv6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.PublicIpv6Address) {
		return nil, false
	}
	return o.PublicIpv6Address, true
}

// IsSetPublicIpv6Address returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetPublicIpv6Address() bool {
	if o != nil && !IsNil(o.PublicIpv6Address) {
		return true
	}

	return false
}

// SetPublicIpv6Address gets a reference to the given string and assigns it to the PublicIpv6Address field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPublicIpv6Address(v string) {
	o.PublicIpv6Address = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// IsSetIpAddress returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6Address() string {
	if o == nil || IsNil(o.Ipv6Address) {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Address) {
		return nil, false
	}
	return o.Ipv6Address, true
}

// IsSetIpv6Address returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetIpv6Address() bool {
	if o != nil && !IsNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetIpSubnet returns the IpSubnet field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpSubnet() string {
	if o == nil || IsNil(o.IpSubnet) {
		var ret string
		return ret
	}
	return *o.IpSubnet
}

// GetIpSubnetOk returns a tuple with the IpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.IpSubnet) {
		return nil, false
	}
	return o.IpSubnet, true
}

// IsSetIpSubnet returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetIpSubnet() bool {
	if o != nil && !IsNil(o.IpSubnet) {
		return true
	}

	return false
}

// SetIpSubnet gets a reference to the given string and assigns it to the IpSubnet field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpSubnet(v string) {
	o.IpSubnet = &v
}

// GetIpv6Subnet returns the Ipv6Subnet field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6Subnet() string {
	if o == nil || IsNil(o.Ipv6Subnet) {
		var ret string
		return ret
	}
	return *o.Ipv6Subnet
}

// GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpv6SubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Subnet) {
		return nil, false
	}
	return o.Ipv6Subnet, true
}

// IsSetIpv6Subnet returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetIpv6Subnet() bool {
	if o != nil && !IsNil(o.Ipv6Subnet) {
		return true
	}

	return false
}

// SetIpv6Subnet gets a reference to the given string and assigns it to the Ipv6Subnet field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpv6Subnet(v string) {
	o.Ipv6Subnet = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetDescription(v string) {
	o.Description = &v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDhcp() bool {
	if o == nil || IsNil(o.Dhcp) {
		var ret bool
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetDhcpOk() (*bool, bool) {
	if o == nil || IsNil(o.Dhcp) {
		return nil, false
	}
	return o.Dhcp, true
}

// IsSetDhcp returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetDhcp() bool {
	if o != nil && !IsNil(o.Dhcp) {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given bool and assigns it to the Dhcp field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetDhcp(v bool) {
	o.Dhcp = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetActive(v bool) {
	o.Active = &v
}

// GetPoolAssigned returns the PoolAssigned field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPoolAssigned() bool {
	if o == nil || IsNil(o.PoolAssigned) {
		var ret bool
		return ret
	}
	return *o.PoolAssigned
}

// GetPoolAssignedOk returns a tuple with the PoolAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPoolAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.PoolAssigned) {
		return nil, false
	}
	return o.PoolAssigned, true
}

// IsSetPoolAssigned returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetPoolAssigned() bool {
	if o != nil && !IsNil(o.PoolAssigned) {
		return true
	}

	return false
}

// SetPoolAssigned gets a reference to the given bool and assigns it to the PoolAssigned field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPoolAssigned(v bool) {
	o.PoolAssigned = &v
}

// GetPrimaryInterface returns the PrimaryInterface field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPrimaryInterface() bool {
	if o == nil || IsNil(o.PrimaryInterface) {
		var ret bool
		return ret
	}
	return *o.PrimaryInterface
}

// GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryInterface) {
		return nil, false
	}
	return o.PrimaryInterface, true
}

// IsSetPrimaryInterface returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetPrimaryInterface() bool {
	if o != nil && !IsNil(o.PrimaryInterface) {
		return true
	}

	return false
}

// SetPrimaryInterface gets a reference to the given bool and assigns it to the PrimaryInterface field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetPrimaryInterface(v bool) {
	o.PrimaryInterface = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetwork() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Network) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// IsSetNetwork returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Network field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetwork(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Network = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// IsSetSubnet returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetSubnet(v string) {
	o.Subnet = &v
}

// GetNetworkGroup returns the NetworkGroup field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkGroup() string {
	if o == nil || IsNil(o.NetworkGroup) {
		var ret string
		return ret
	}
	return *o.NetworkGroup
}

// GetNetworkGroupOk returns a tuple with the NetworkGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkGroupOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkGroup) {
		return nil, false
	}
	return o.NetworkGroup, true
}

// IsSetNetworkGroup returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetNetworkGroup() bool {
	if o != nil && !IsNil(o.NetworkGroup) {
		return true
	}

	return false
}

// SetNetworkGroup gets a reference to the given string and assigns it to the NetworkGroup field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkGroup(v string) {
	o.NetworkGroup = &v
}

// GetNetworkPosition returns the NetworkPosition field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPosition() string {
	if o == nil || IsNil(o.NetworkPosition) {
		var ret string
		return ret
	}
	return *o.NetworkPosition
}

// GetNetworkPositionOk returns a tuple with the NetworkPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPositionOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkPosition) {
		return nil, false
	}
	return o.NetworkPosition, true
}

// IsSetNetworkPosition returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetNetworkPosition() bool {
	if o != nil && !IsNil(o.NetworkPosition) {
		return true
	}

	return false
}

// SetNetworkPosition gets a reference to the given string and assigns it to the NetworkPosition field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkPosition(v string) {
	o.NetworkPosition = &v
}

// GetNetworkPool returns the NetworkPool field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance {
	if o == nil || IsNil(o.NetworkPool) {
		var ret GetAlerts200ResponseAllOfCheckGroupsInnerInstance
		return ret
	}
	return *o.NetworkPool
}

// GetNetworkPoolOk returns a tuple with the NetworkPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool) {
	if o == nil || IsNil(o.NetworkPool) {
		return nil, false
	}
	return o.NetworkPool, true
}

// IsSetNetworkPool returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetNetworkPool() bool {
	if o != nil && !IsNil(o.NetworkPool) {
		return true
	}

	return false
}

// SetNetworkPool gets a reference to the given GetAlerts200ResponseAllOfCheckGroupsInnerInstance and assigns it to the NetworkPool field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance) {
	o.NetworkPool = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomain() string {
	if o == nil || IsNil(o.NetworkDomain) {
		var ret string
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetNetworkDomainOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkDomain) {
		return nil, false
	}
	return o.NetworkDomain, true
}

// IsSetNetworkDomain returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetNetworkDomain() bool {
	if o != nil && !IsNil(o.NetworkDomain) {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given string and assigns it to the NetworkDomain field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetNetworkDomain(v string) {
	o.NetworkDomain = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetType() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType {
	if o == nil || IsNil(o.Type) {
		var ret ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetTypeOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType and assigns it to the Type field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetType(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType) {
	o.Type = &v
}

// GetIpMode returns the IpMode field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpMode() string {
	if o == nil || IsNil(o.IpMode) {
		var ret string
		return ret
	}
	return *o.IpMode
}

// GetIpModeOk returns a tuple with the IpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetIpModeOk() (*string, bool) {
	if o == nil || IsNil(o.IpMode) {
		return nil, false
	}
	return o.IpMode, true
}

// IsSetIpMode returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetIpMode() bool {
	if o != nil && !IsNil(o.IpMode) {
		return true
	}

	return false
}

// SetIpMode gets a reference to the given string and assigns it to the IpMode field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetIpMode(v string) {
	o.IpMode = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// IsSetMacAddress returns a boolean if a field has been set.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) IsSetMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o ListHosts200ResponseAllOfServersInnerInterfacesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHosts200ResponseAllOfServersInnerInterfacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RefType) {
		toSerialize["refType"] = o.RefType
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.InternalId) {
		toSerialize["internalId"] = o.InternalId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.UniqueId) {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if !IsNil(o.PublicIpAddress) {
		toSerialize["publicIpAddress"] = o.PublicIpAddress
	}
	if !IsNil(o.PublicIpv6Address) {
		toSerialize["publicIpv6Address"] = o.PublicIpv6Address
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.Ipv6Address) {
		toSerialize["ipv6Address"] = o.Ipv6Address
	}
	if !IsNil(o.IpSubnet) {
		toSerialize["ipSubnet"] = o.IpSubnet
	}
	if !IsNil(o.Ipv6Subnet) {
		toSerialize["ipv6Subnet"] = o.Ipv6Subnet
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Dhcp) {
		toSerialize["dhcp"] = o.Dhcp
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.PoolAssigned) {
		toSerialize["poolAssigned"] = o.PoolAssigned
	}
	if !IsNil(o.PrimaryInterface) {
		toSerialize["primaryInterface"] = o.PrimaryInterface
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.NetworkGroup) {
		toSerialize["networkGroup"] = o.NetworkGroup
	}
	if !IsNil(o.NetworkPosition) {
		toSerialize["networkPosition"] = o.NetworkPosition
	}
	if !IsNil(o.NetworkPool) {
		toSerialize["networkPool"] = o.NetworkPool
	}
	if !IsNil(o.NetworkDomain) {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IpMode) {
		toSerialize["ipMode"] = o.IpMode
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListHosts200ResponseAllOfServersInnerInterfacesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
