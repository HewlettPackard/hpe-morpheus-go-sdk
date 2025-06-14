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

// checks if the ListHealth200ResponseAllOfHealthElastic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHealth200ResponseAllOfHealthElastic{}

// ListHealth200ResponseAllOfHealthElastic struct for ListHealth200ResponseAllOfHealthElastic
type ListHealth200ResponseAllOfHealthElastic struct {
	Success              *bool                                               `json:"success,omitempty"`
	Status               *string                                             `json:"status,omitempty"`
	Master               *ListHealth200ResponseAllOfHealthElasticMaster      `json:"master,omitempty"`
	Nodes                []ListHealth200ResponseAllOfHealthElasticNodesInner `json:"nodes,omitempty"`
	Stats                *ListHealth200ResponseAllOfHealthElasticStats       `json:"stats,omitempty"`
	Indices              []map[string]interface{}                            `json:"indices,omitempty"`
	BadIndices           []map[string]interface{}                            `json:"badIndices,omitempty"`
	AdditionalProperties map[string]interface{}                              `json:",remain"`
}

type _ListHealth200ResponseAllOfHealthElastic ListHealth200ResponseAllOfHealthElastic

// NewListHealth200ResponseAllOfHealthElastic instantiates a new ListHealth200ResponseAllOfHealthElastic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHealth200ResponseAllOfHealthElastic() *ListHealth200ResponseAllOfHealthElastic {
	this := ListHealth200ResponseAllOfHealthElastic{}
	return &this
}

// NewListHealth200ResponseAllOfHealthElasticWithDefaults instantiates a new ListHealth200ResponseAllOfHealthElastic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHealth200ResponseAllOfHealthElasticWithDefaults() *ListHealth200ResponseAllOfHealthElastic {
	this := ListHealth200ResponseAllOfHealthElastic{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetSuccess(v bool) {
	o.Success = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetStatus(v string) {
	o.Status = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetMaster() ListHealth200ResponseAllOfHealthElasticMaster {
	if o == nil || IsNil(o.Master) {
		var ret ListHealth200ResponseAllOfHealthElasticMaster
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetMasterOk() (*ListHealth200ResponseAllOfHealthElasticMaster, bool) {
	if o == nil || IsNil(o.Master) {
		return nil, false
	}
	return o.Master, true
}

// IsSetMaster returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetMaster() bool {
	if o != nil && !IsNil(o.Master) {
		return true
	}

	return false
}

// SetMaster gets a reference to the given ListHealth200ResponseAllOfHealthElasticMaster and assigns it to the Master field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetMaster(v ListHealth200ResponseAllOfHealthElasticMaster) {
	o.Master = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetNodes() []ListHealth200ResponseAllOfHealthElasticNodesInner {
	if o == nil || IsNil(o.Nodes) {
		var ret []ListHealth200ResponseAllOfHealthElasticNodesInner
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetNodesOk() ([]ListHealth200ResponseAllOfHealthElasticNodesInner, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// IsSetNodes returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []ListHealth200ResponseAllOfHealthElasticNodesInner and assigns it to the Nodes field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetNodes(v []ListHealth200ResponseAllOfHealthElasticNodesInner) {
	o.Nodes = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetStats() ListHealth200ResponseAllOfHealthElasticStats {
	if o == nil || IsNil(o.Stats) {
		var ret ListHealth200ResponseAllOfHealthElasticStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetStatsOk() (*ListHealth200ResponseAllOfHealthElasticStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// IsSetStats returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ListHealth200ResponseAllOfHealthElasticStats and assigns it to the Stats field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetStats(v ListHealth200ResponseAllOfHealthElasticStats) {
	o.Stats = &v
}

// GetIndices returns the Indices field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetIndices() []map[string]interface{} {
	if o == nil || IsNil(o.Indices) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Indices
}

// GetIndicesOk returns a tuple with the Indices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetIndicesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Indices) {
		return nil, false
	}
	return o.Indices, true
}

// IsSetIndices returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetIndices() bool {
	if o != nil && !IsNil(o.Indices) {
		return true
	}

	return false
}

// SetIndices gets a reference to the given []map[string]interface{} and assigns it to the Indices field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetIndices(v []map[string]interface{}) {
	o.Indices = v
}

// GetBadIndices returns the BadIndices field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthElastic) GetBadIndices() []map[string]interface{} {
	if o == nil || IsNil(o.BadIndices) {
		var ret []map[string]interface{}
		return ret
	}
	return o.BadIndices
}

// GetBadIndicesOk returns a tuple with the BadIndices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) GetBadIndicesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.BadIndices) {
		return nil, false
	}
	return o.BadIndices, true
}

// IsSetBadIndices returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthElastic) IsSetBadIndices() bool {
	if o != nil && !IsNil(o.BadIndices) {
		return true
	}

	return false
}

// SetBadIndices gets a reference to the given []map[string]interface{} and assigns it to the BadIndices field.
func (o *ListHealth200ResponseAllOfHealthElastic) SetBadIndices(v []map[string]interface{}) {
	o.BadIndices = v
}

func (o ListHealth200ResponseAllOfHealthElastic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHealth200ResponseAllOfHealthElastic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Master) {
		toSerialize["master"] = o.Master
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Indices) {
		toSerialize["indices"] = o.Indices
	}
	if !IsNil(o.BadIndices) {
		toSerialize["badIndices"] = o.BadIndices
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListHealth200ResponseAllOfHealthElastic) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
