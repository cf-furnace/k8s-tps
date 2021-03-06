package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*LRPInstance l r p instance

swagger:model LRPInstance
*/
type LRPInstance struct {

	/* details
	 */
	Details string `json:"details,omitempty"`

	/* host
	 */
	Host string `json:"host,omitempty"`

	/* index
	 */
	Index int64 `json:"index,omitempty"`

	/* instance guid
	 */
	InstanceGUID string `json:"instance_guid,omitempty"`

	/* net info
	 */
	NetInfo *ActualLRPNetInfo `json:"net_info,omitempty"`

	/* port
	 */
	Port int32 `json:"port,omitempty"`

	/* process guid
	 */
	ProcessGUID string `json:"process_guid,omitempty"`

	/* since
	 */
	Since int64 `json:"since,omitempty"`

	/* state
	 */
	State string `json:"state,omitempty"`

	/* stats
	 */
	Stats *LRPInstanceStats `json:"stats,omitempty"`

	/* uptime
	 */
	Uptime int64 `json:"uptime,omitempty"`
}

// Validate validates this l r p instance
func (m *LRPInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LRPInstance) validateNetInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.NetInfo) { // not required
		return nil
	}

	if m.NetInfo != nil {

		if err := m.NetInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *LRPInstance) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {

		if err := m.Stats.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
