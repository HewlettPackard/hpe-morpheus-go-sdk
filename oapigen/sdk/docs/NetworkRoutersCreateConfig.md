# NetworkRoutersCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaMode** | Pointer to **string** |  | [optional] 
**EdgeCluster** | Pointer to **string** |  | [optional] 
**FailOver** | Pointer to **string** |  | [optional] 
**IpManagementType** | Pointer to **string** |  | [optional] 
**IpServerId** | Pointer to **string** |  | [optional] 
**TIER0STATIC** | Pointer to **string** |  | [optional] 
**TIER0NAT** | Pointer to **string** |  | [optional] 
**TIER0IPSECLOCALIP** | Pointer to **string** |  | [optional] 
**TIER0DNSFORWARDERIP** | Pointer to **string** |  | [optional] 
**TIER0SERVICEINTERFACE** | Pointer to **string** |  | [optional] 
**TIER0EXTERNALINTERFACE** | Pointer to **string** |  | [optional] 
**TIER0LOOPBACKINTERFACE** | Pointer to **string** |  | [optional] 
**TIER0SEGMENT** | Pointer to **string** |  | [optional] 
**TIER1DNSFORWARDERIP** | Pointer to **string** |  | [optional] 
**TIER1STATIC** | Pointer to **string** |  | [optional] 
**TIER1LBVIP** | Pointer to **string** |  | [optional] 
**TIER1NAT** | Pointer to **string** |  | [optional] 
**TIER1LBSNAT** | Pointer to **string** |  | [optional] 
**TIER1IPSECLOCALENDPOINT** | Pointer to **string** |  | [optional] 
**TIER1SERVICEINTERFACE** | Pointer to **string** |  | [optional] 
**TIER1SEGMENT** | Pointer to **string** |  | [optional] 
**LOCAL_AS_NUM** | Pointer to **string** |  | [optional] 
**ECMP** | Pointer to **string** |  | [optional] 
**MULTIPATH_RELAX** | Pointer to **string** |  | [optional] 
**RESTART_MODE** | Pointer to **string** |  | [optional] 
**RESTART_TIME** | Pointer to **int64** |  | [optional] 
**STALE_ROUTE_TIME** | Pointer to **int64** |  | [optional] 
**INTER_SR_IBGP** | Pointer to **string** |  | [optional] 
**Tier0Gateway** | Pointer to **string** |  | [optional] 
**TIER1CONNECTED** | Pointer to **string** |  | [optional] 
**TIER1STATICROUTES** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkRoutersCreateConfig

`func NewNetworkRoutersCreateConfig() *NetworkRoutersCreateConfig`

NewNetworkRoutersCreateConfig instantiates a new NetworkRoutersCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetHaMode

`func (o *NetworkRoutersCreateConfig) GetHaMode() string`

GetHaMode returns the HaMode field if non-nil, zero value otherwise.

### GetHaModeOk

`func (o *NetworkRoutersCreateConfig) GetHaModeOk() (*string, bool)`

GetHaModeOk returns a tuple with the HaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaMode

`func (o *NetworkRoutersCreateConfig) SetHaMode(v string)`

SetHaMode sets HaMode field to given value.

### HasHaMode

`func (o *NetworkRoutersCreateConfig) HasHaMode() bool`

HasHaMode returns a boolean if a field has been set.

### GetEdgeCluster

`func (o *NetworkRoutersCreateConfig) GetEdgeCluster() string`

GetEdgeCluster returns the EdgeCluster field if non-nil, zero value otherwise.

### GetEdgeClusterOk

`func (o *NetworkRoutersCreateConfig) GetEdgeClusterOk() (*string, bool)`

GetEdgeClusterOk returns a tuple with the EdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCluster

`func (o *NetworkRoutersCreateConfig) SetEdgeCluster(v string)`

SetEdgeCluster sets EdgeCluster field to given value.

### HasEdgeCluster

`func (o *NetworkRoutersCreateConfig) HasEdgeCluster() bool`

HasEdgeCluster returns a boolean if a field has been set.

### GetFailOver

`func (o *NetworkRoutersCreateConfig) GetFailOver() string`

GetFailOver returns the FailOver field if non-nil, zero value otherwise.

### GetFailOverOk

`func (o *NetworkRoutersCreateConfig) GetFailOverOk() (*string, bool)`

GetFailOverOk returns a tuple with the FailOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOver

`func (o *NetworkRoutersCreateConfig) SetFailOver(v string)`

SetFailOver sets FailOver field to given value.

### HasFailOver

`func (o *NetworkRoutersCreateConfig) HasFailOver() bool`

HasFailOver returns a boolean if a field has been set.

### GetIpManagementType

`func (o *NetworkRoutersCreateConfig) GetIpManagementType() string`

GetIpManagementType returns the IpManagementType field if non-nil, zero value otherwise.

### GetIpManagementTypeOk

`func (o *NetworkRoutersCreateConfig) GetIpManagementTypeOk() (*string, bool)`

GetIpManagementTypeOk returns a tuple with the IpManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpManagementType

`func (o *NetworkRoutersCreateConfig) SetIpManagementType(v string)`

SetIpManagementType sets IpManagementType field to given value.

### HasIpManagementType

`func (o *NetworkRoutersCreateConfig) HasIpManagementType() bool`

HasIpManagementType returns a boolean if a field has been set.

### GetIpServerId

`func (o *NetworkRoutersCreateConfig) GetIpServerId() string`

GetIpServerId returns the IpServerId field if non-nil, zero value otherwise.

### GetIpServerIdOk

`func (o *NetworkRoutersCreateConfig) GetIpServerIdOk() (*string, bool)`

GetIpServerIdOk returns a tuple with the IpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpServerId

`func (o *NetworkRoutersCreateConfig) SetIpServerId(v string)`

SetIpServerId sets IpServerId field to given value.

### HasIpServerId

`func (o *NetworkRoutersCreateConfig) HasIpServerId() bool`

HasIpServerId returns a boolean if a field has been set.

### GetTIER0STATIC

`func (o *NetworkRoutersCreateConfig) GetTIER0STATIC() string`

GetTIER0STATIC returns the TIER0STATIC field if non-nil, zero value otherwise.

### GetTIER0STATICOk

`func (o *NetworkRoutersCreateConfig) GetTIER0STATICOk() (*string, bool)`

GetTIER0STATICOk returns a tuple with the TIER0STATIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0STATIC

`func (o *NetworkRoutersCreateConfig) SetTIER0STATIC(v string)`

SetTIER0STATIC sets TIER0STATIC field to given value.

### HasTIER0STATIC

`func (o *NetworkRoutersCreateConfig) HasTIER0STATIC() bool`

HasTIER0STATIC returns a boolean if a field has been set.

### GetTIER0NAT

`func (o *NetworkRoutersCreateConfig) GetTIER0NAT() string`

GetTIER0NAT returns the TIER0NAT field if non-nil, zero value otherwise.

### GetTIER0NATOk

`func (o *NetworkRoutersCreateConfig) GetTIER0NATOk() (*string, bool)`

GetTIER0NATOk returns a tuple with the TIER0NAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0NAT

`func (o *NetworkRoutersCreateConfig) SetTIER0NAT(v string)`

SetTIER0NAT sets TIER0NAT field to given value.

### HasTIER0NAT

`func (o *NetworkRoutersCreateConfig) HasTIER0NAT() bool`

HasTIER0NAT returns a boolean if a field has been set.

### GetTIER0IPSECLOCALIP

`func (o *NetworkRoutersCreateConfig) GetTIER0IPSECLOCALIP() string`

GetTIER0IPSECLOCALIP returns the TIER0IPSECLOCALIP field if non-nil, zero value otherwise.

### GetTIER0IPSECLOCALIPOk

`func (o *NetworkRoutersCreateConfig) GetTIER0IPSECLOCALIPOk() (*string, bool)`

GetTIER0IPSECLOCALIPOk returns a tuple with the TIER0IPSECLOCALIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0IPSECLOCALIP

`func (o *NetworkRoutersCreateConfig) SetTIER0IPSECLOCALIP(v string)`

SetTIER0IPSECLOCALIP sets TIER0IPSECLOCALIP field to given value.

### HasTIER0IPSECLOCALIP

`func (o *NetworkRoutersCreateConfig) HasTIER0IPSECLOCALIP() bool`

HasTIER0IPSECLOCALIP returns a boolean if a field has been set.

### GetTIER0DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) GetTIER0DNSFORWARDERIP() string`

GetTIER0DNSFORWARDERIP returns the TIER0DNSFORWARDERIP field if non-nil, zero value otherwise.

### GetTIER0DNSFORWARDERIPOk

`func (o *NetworkRoutersCreateConfig) GetTIER0DNSFORWARDERIPOk() (*string, bool)`

GetTIER0DNSFORWARDERIPOk returns a tuple with the TIER0DNSFORWARDERIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) SetTIER0DNSFORWARDERIP(v string)`

SetTIER0DNSFORWARDERIP sets TIER0DNSFORWARDERIP field to given value.

### HasTIER0DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) HasTIER0DNSFORWARDERIP() bool`

HasTIER0DNSFORWARDERIP returns a boolean if a field has been set.

### GetTIER0SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) GetTIER0SERVICEINTERFACE() string`

GetTIER0SERVICEINTERFACE returns the TIER0SERVICEINTERFACE field if non-nil, zero value otherwise.

### GetTIER0SERVICEINTERFACEOk

`func (o *NetworkRoutersCreateConfig) GetTIER0SERVICEINTERFACEOk() (*string, bool)`

GetTIER0SERVICEINTERFACEOk returns a tuple with the TIER0SERVICEINTERFACE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) SetTIER0SERVICEINTERFACE(v string)`

SetTIER0SERVICEINTERFACE sets TIER0SERVICEINTERFACE field to given value.

### HasTIER0SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) HasTIER0SERVICEINTERFACE() bool`

HasTIER0SERVICEINTERFACE returns a boolean if a field has been set.

### GetTIER0EXTERNALINTERFACE

`func (o *NetworkRoutersCreateConfig) GetTIER0EXTERNALINTERFACE() string`

GetTIER0EXTERNALINTERFACE returns the TIER0EXTERNALINTERFACE field if non-nil, zero value otherwise.

### GetTIER0EXTERNALINTERFACEOk

`func (o *NetworkRoutersCreateConfig) GetTIER0EXTERNALINTERFACEOk() (*string, bool)`

GetTIER0EXTERNALINTERFACEOk returns a tuple with the TIER0EXTERNALINTERFACE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0EXTERNALINTERFACE

`func (o *NetworkRoutersCreateConfig) SetTIER0EXTERNALINTERFACE(v string)`

SetTIER0EXTERNALINTERFACE sets TIER0EXTERNALINTERFACE field to given value.

### HasTIER0EXTERNALINTERFACE

`func (o *NetworkRoutersCreateConfig) HasTIER0EXTERNALINTERFACE() bool`

HasTIER0EXTERNALINTERFACE returns a boolean if a field has been set.

### GetTIER0LOOPBACKINTERFACE

`func (o *NetworkRoutersCreateConfig) GetTIER0LOOPBACKINTERFACE() string`

GetTIER0LOOPBACKINTERFACE returns the TIER0LOOPBACKINTERFACE field if non-nil, zero value otherwise.

### GetTIER0LOOPBACKINTERFACEOk

`func (o *NetworkRoutersCreateConfig) GetTIER0LOOPBACKINTERFACEOk() (*string, bool)`

GetTIER0LOOPBACKINTERFACEOk returns a tuple with the TIER0LOOPBACKINTERFACE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0LOOPBACKINTERFACE

`func (o *NetworkRoutersCreateConfig) SetTIER0LOOPBACKINTERFACE(v string)`

SetTIER0LOOPBACKINTERFACE sets TIER0LOOPBACKINTERFACE field to given value.

### HasTIER0LOOPBACKINTERFACE

`func (o *NetworkRoutersCreateConfig) HasTIER0LOOPBACKINTERFACE() bool`

HasTIER0LOOPBACKINTERFACE returns a boolean if a field has been set.

### GetTIER0SEGMENT

`func (o *NetworkRoutersCreateConfig) GetTIER0SEGMENT() string`

GetTIER0SEGMENT returns the TIER0SEGMENT field if non-nil, zero value otherwise.

### GetTIER0SEGMENTOk

`func (o *NetworkRoutersCreateConfig) GetTIER0SEGMENTOk() (*string, bool)`

GetTIER0SEGMENTOk returns a tuple with the TIER0SEGMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER0SEGMENT

`func (o *NetworkRoutersCreateConfig) SetTIER0SEGMENT(v string)`

SetTIER0SEGMENT sets TIER0SEGMENT field to given value.

### HasTIER0SEGMENT

`func (o *NetworkRoutersCreateConfig) HasTIER0SEGMENT() bool`

HasTIER0SEGMENT returns a boolean if a field has been set.

### GetTIER1DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) GetTIER1DNSFORWARDERIP() string`

GetTIER1DNSFORWARDERIP returns the TIER1DNSFORWARDERIP field if non-nil, zero value otherwise.

### GetTIER1DNSFORWARDERIPOk

`func (o *NetworkRoutersCreateConfig) GetTIER1DNSFORWARDERIPOk() (*string, bool)`

GetTIER1DNSFORWARDERIPOk returns a tuple with the TIER1DNSFORWARDERIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) SetTIER1DNSFORWARDERIP(v string)`

SetTIER1DNSFORWARDERIP sets TIER1DNSFORWARDERIP field to given value.

### HasTIER1DNSFORWARDERIP

`func (o *NetworkRoutersCreateConfig) HasTIER1DNSFORWARDERIP() bool`

HasTIER1DNSFORWARDERIP returns a boolean if a field has been set.

### GetTIER1STATIC

`func (o *NetworkRoutersCreateConfig) GetTIER1STATIC() string`

GetTIER1STATIC returns the TIER1STATIC field if non-nil, zero value otherwise.

### GetTIER1STATICOk

`func (o *NetworkRoutersCreateConfig) GetTIER1STATICOk() (*string, bool)`

GetTIER1STATICOk returns a tuple with the TIER1STATIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1STATIC

`func (o *NetworkRoutersCreateConfig) SetTIER1STATIC(v string)`

SetTIER1STATIC sets TIER1STATIC field to given value.

### HasTIER1STATIC

`func (o *NetworkRoutersCreateConfig) HasTIER1STATIC() bool`

HasTIER1STATIC returns a boolean if a field has been set.

### GetTIER1LBVIP

`func (o *NetworkRoutersCreateConfig) GetTIER1LBVIP() string`

GetTIER1LBVIP returns the TIER1LBVIP field if non-nil, zero value otherwise.

### GetTIER1LBVIPOk

`func (o *NetworkRoutersCreateConfig) GetTIER1LBVIPOk() (*string, bool)`

GetTIER1LBVIPOk returns a tuple with the TIER1LBVIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1LBVIP

`func (o *NetworkRoutersCreateConfig) SetTIER1LBVIP(v string)`

SetTIER1LBVIP sets TIER1LBVIP field to given value.

### HasTIER1LBVIP

`func (o *NetworkRoutersCreateConfig) HasTIER1LBVIP() bool`

HasTIER1LBVIP returns a boolean if a field has been set.

### GetTIER1NAT

`func (o *NetworkRoutersCreateConfig) GetTIER1NAT() string`

GetTIER1NAT returns the TIER1NAT field if non-nil, zero value otherwise.

### GetTIER1NATOk

`func (o *NetworkRoutersCreateConfig) GetTIER1NATOk() (*string, bool)`

GetTIER1NATOk returns a tuple with the TIER1NAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1NAT

`func (o *NetworkRoutersCreateConfig) SetTIER1NAT(v string)`

SetTIER1NAT sets TIER1NAT field to given value.

### HasTIER1NAT

`func (o *NetworkRoutersCreateConfig) HasTIER1NAT() bool`

HasTIER1NAT returns a boolean if a field has been set.

### GetTIER1LBSNAT

`func (o *NetworkRoutersCreateConfig) GetTIER1LBSNAT() string`

GetTIER1LBSNAT returns the TIER1LBSNAT field if non-nil, zero value otherwise.

### GetTIER1LBSNATOk

`func (o *NetworkRoutersCreateConfig) GetTIER1LBSNATOk() (*string, bool)`

GetTIER1LBSNATOk returns a tuple with the TIER1LBSNAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1LBSNAT

`func (o *NetworkRoutersCreateConfig) SetTIER1LBSNAT(v string)`

SetTIER1LBSNAT sets TIER1LBSNAT field to given value.

### HasTIER1LBSNAT

`func (o *NetworkRoutersCreateConfig) HasTIER1LBSNAT() bool`

HasTIER1LBSNAT returns a boolean if a field has been set.

### GetTIER1IPSECLOCALENDPOINT

`func (o *NetworkRoutersCreateConfig) GetTIER1IPSECLOCALENDPOINT() string`

GetTIER1IPSECLOCALENDPOINT returns the TIER1IPSECLOCALENDPOINT field if non-nil, zero value otherwise.

### GetTIER1IPSECLOCALENDPOINTOk

`func (o *NetworkRoutersCreateConfig) GetTIER1IPSECLOCALENDPOINTOk() (*string, bool)`

GetTIER1IPSECLOCALENDPOINTOk returns a tuple with the TIER1IPSECLOCALENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1IPSECLOCALENDPOINT

`func (o *NetworkRoutersCreateConfig) SetTIER1IPSECLOCALENDPOINT(v string)`

SetTIER1IPSECLOCALENDPOINT sets TIER1IPSECLOCALENDPOINT field to given value.

### HasTIER1IPSECLOCALENDPOINT

`func (o *NetworkRoutersCreateConfig) HasTIER1IPSECLOCALENDPOINT() bool`

HasTIER1IPSECLOCALENDPOINT returns a boolean if a field has been set.

### GetTIER1SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) GetTIER1SERVICEINTERFACE() string`

GetTIER1SERVICEINTERFACE returns the TIER1SERVICEINTERFACE field if non-nil, zero value otherwise.

### GetTIER1SERVICEINTERFACEOk

`func (o *NetworkRoutersCreateConfig) GetTIER1SERVICEINTERFACEOk() (*string, bool)`

GetTIER1SERVICEINTERFACEOk returns a tuple with the TIER1SERVICEINTERFACE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) SetTIER1SERVICEINTERFACE(v string)`

SetTIER1SERVICEINTERFACE sets TIER1SERVICEINTERFACE field to given value.

### HasTIER1SERVICEINTERFACE

`func (o *NetworkRoutersCreateConfig) HasTIER1SERVICEINTERFACE() bool`

HasTIER1SERVICEINTERFACE returns a boolean if a field has been set.

### GetTIER1SEGMENT

`func (o *NetworkRoutersCreateConfig) GetTIER1SEGMENT() string`

GetTIER1SEGMENT returns the TIER1SEGMENT field if non-nil, zero value otherwise.

### GetTIER1SEGMENTOk

`func (o *NetworkRoutersCreateConfig) GetTIER1SEGMENTOk() (*string, bool)`

GetTIER1SEGMENTOk returns a tuple with the TIER1SEGMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1SEGMENT

`func (o *NetworkRoutersCreateConfig) SetTIER1SEGMENT(v string)`

SetTIER1SEGMENT sets TIER1SEGMENT field to given value.

### HasTIER1SEGMENT

`func (o *NetworkRoutersCreateConfig) HasTIER1SEGMENT() bool`

HasTIER1SEGMENT returns a boolean if a field has been set.

### GetLOCAL_AS_NUM

`func (o *NetworkRoutersCreateConfig) GetLOCAL_AS_NUM() string`

GetLOCAL_AS_NUM returns the LOCAL_AS_NUM field if non-nil, zero value otherwise.

### GetLOCAL_AS_NUMOk

`func (o *NetworkRoutersCreateConfig) GetLOCAL_AS_NUMOk() (*string, bool)`

GetLOCAL_AS_NUMOk returns a tuple with the LOCAL_AS_NUM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOCAL_AS_NUM

`func (o *NetworkRoutersCreateConfig) SetLOCAL_AS_NUM(v string)`

SetLOCAL_AS_NUM sets LOCAL_AS_NUM field to given value.

### HasLOCAL_AS_NUM

`func (o *NetworkRoutersCreateConfig) HasLOCAL_AS_NUM() bool`

HasLOCAL_AS_NUM returns a boolean if a field has been set.

### GetECMP

`func (o *NetworkRoutersCreateConfig) GetECMP() string`

GetECMP returns the ECMP field if non-nil, zero value otherwise.

### GetECMPOk

`func (o *NetworkRoutersCreateConfig) GetECMPOk() (*string, bool)`

GetECMPOk returns a tuple with the ECMP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECMP

`func (o *NetworkRoutersCreateConfig) SetECMP(v string)`

SetECMP sets ECMP field to given value.

### HasECMP

`func (o *NetworkRoutersCreateConfig) HasECMP() bool`

HasECMP returns a boolean if a field has been set.

### GetMULTIPATH_RELAX

`func (o *NetworkRoutersCreateConfig) GetMULTIPATH_RELAX() string`

GetMULTIPATH_RELAX returns the MULTIPATH_RELAX field if non-nil, zero value otherwise.

### GetMULTIPATH_RELAXOk

`func (o *NetworkRoutersCreateConfig) GetMULTIPATH_RELAXOk() (*string, bool)`

GetMULTIPATH_RELAXOk returns a tuple with the MULTIPATH_RELAX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMULTIPATH_RELAX

`func (o *NetworkRoutersCreateConfig) SetMULTIPATH_RELAX(v string)`

SetMULTIPATH_RELAX sets MULTIPATH_RELAX field to given value.

### HasMULTIPATH_RELAX

`func (o *NetworkRoutersCreateConfig) HasMULTIPATH_RELAX() bool`

HasMULTIPATH_RELAX returns a boolean if a field has been set.

### GetRESTART_MODE

`func (o *NetworkRoutersCreateConfig) GetRESTART_MODE() string`

GetRESTART_MODE returns the RESTART_MODE field if non-nil, zero value otherwise.

### GetRESTART_MODEOk

`func (o *NetworkRoutersCreateConfig) GetRESTART_MODEOk() (*string, bool)`

GetRESTART_MODEOk returns a tuple with the RESTART_MODE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESTART_MODE

`func (o *NetworkRoutersCreateConfig) SetRESTART_MODE(v string)`

SetRESTART_MODE sets RESTART_MODE field to given value.

### HasRESTART_MODE

`func (o *NetworkRoutersCreateConfig) HasRESTART_MODE() bool`

HasRESTART_MODE returns a boolean if a field has been set.

### GetRESTART_TIME

`func (o *NetworkRoutersCreateConfig) GetRESTART_TIME() int64`

GetRESTART_TIME returns the RESTART_TIME field if non-nil, zero value otherwise.

### GetRESTART_TIMEOk

`func (o *NetworkRoutersCreateConfig) GetRESTART_TIMEOk() (*int64, bool)`

GetRESTART_TIMEOk returns a tuple with the RESTART_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESTART_TIME

`func (o *NetworkRoutersCreateConfig) SetRESTART_TIME(v int64)`

SetRESTART_TIME sets RESTART_TIME field to given value.

### HasRESTART_TIME

`func (o *NetworkRoutersCreateConfig) HasRESTART_TIME() bool`

HasRESTART_TIME returns a boolean if a field has been set.

### GetSTALE_ROUTE_TIME

`func (o *NetworkRoutersCreateConfig) GetSTALE_ROUTE_TIME() int64`

GetSTALE_ROUTE_TIME returns the STALE_ROUTE_TIME field if non-nil, zero value otherwise.

### GetSTALE_ROUTE_TIMEOk

`func (o *NetworkRoutersCreateConfig) GetSTALE_ROUTE_TIMEOk() (*int64, bool)`

GetSTALE_ROUTE_TIMEOk returns a tuple with the STALE_ROUTE_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTALE_ROUTE_TIME

`func (o *NetworkRoutersCreateConfig) SetSTALE_ROUTE_TIME(v int64)`

SetSTALE_ROUTE_TIME sets STALE_ROUTE_TIME field to given value.

### HasSTALE_ROUTE_TIME

`func (o *NetworkRoutersCreateConfig) HasSTALE_ROUTE_TIME() bool`

HasSTALE_ROUTE_TIME returns a boolean if a field has been set.

### GetINTER_SR_IBGP

`func (o *NetworkRoutersCreateConfig) GetINTER_SR_IBGP() string`

GetINTER_SR_IBGP returns the INTER_SR_IBGP field if non-nil, zero value otherwise.

### GetINTER_SR_IBGPOk

`func (o *NetworkRoutersCreateConfig) GetINTER_SR_IBGPOk() (*string, bool)`

GetINTER_SR_IBGPOk returns a tuple with the INTER_SR_IBGP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINTER_SR_IBGP

`func (o *NetworkRoutersCreateConfig) SetINTER_SR_IBGP(v string)`

SetINTER_SR_IBGP sets INTER_SR_IBGP field to given value.

### HasINTER_SR_IBGP

`func (o *NetworkRoutersCreateConfig) HasINTER_SR_IBGP() bool`

HasINTER_SR_IBGP returns a boolean if a field has been set.

### GetTier0Gateway

`func (o *NetworkRoutersCreateConfig) GetTier0Gateway() string`

GetTier0Gateway returns the Tier0Gateway field if non-nil, zero value otherwise.

### GetTier0GatewayOk

`func (o *NetworkRoutersCreateConfig) GetTier0GatewayOk() (*string, bool)`

GetTier0GatewayOk returns a tuple with the Tier0Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier0Gateway

`func (o *NetworkRoutersCreateConfig) SetTier0Gateway(v string)`

SetTier0Gateway sets Tier0Gateway field to given value.

### HasTier0Gateway

`func (o *NetworkRoutersCreateConfig) HasTier0Gateway() bool`

HasTier0Gateway returns a boolean if a field has been set.

### GetTIER1CONNECTED

`func (o *NetworkRoutersCreateConfig) GetTIER1CONNECTED() string`

GetTIER1CONNECTED returns the TIER1CONNECTED field if non-nil, zero value otherwise.

### GetTIER1CONNECTEDOk

`func (o *NetworkRoutersCreateConfig) GetTIER1CONNECTEDOk() (*string, bool)`

GetTIER1CONNECTEDOk returns a tuple with the TIER1CONNECTED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1CONNECTED

`func (o *NetworkRoutersCreateConfig) SetTIER1CONNECTED(v string)`

SetTIER1CONNECTED sets TIER1CONNECTED field to given value.

### HasTIER1CONNECTED

`func (o *NetworkRoutersCreateConfig) HasTIER1CONNECTED() bool`

HasTIER1CONNECTED returns a boolean if a field has been set.

### GetTIER1STATICROUTES

`func (o *NetworkRoutersCreateConfig) GetTIER1STATICROUTES() string`

GetTIER1STATICROUTES returns the TIER1STATICROUTES field if non-nil, zero value otherwise.

### GetTIER1STATICROUTESOk

`func (o *NetworkRoutersCreateConfig) GetTIER1STATICROUTESOk() (*string, bool)`

GetTIER1STATICROUTESOk returns a tuple with the TIER1STATICROUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTIER1STATICROUTES

`func (o *NetworkRoutersCreateConfig) SetTIER1STATICROUTES(v string)`

SetTIER1STATICROUTES sets TIER1STATICROUTES field to given value.

### HasTIER1STATICROUTES

`func (o *NetworkRoutersCreateConfig) HasTIER1STATICROUTES() bool`

HasTIER1STATICROUTES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


