package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*LRPInstanceStats l r p instance stats

swagger:model LRPInstanceStats
*/
type LRPInstanceStats struct {

	/* cpu
	 */
	CPU float64 `json:"cpu,omitempty"`

	/* disk
	 */
	Disk int64 `json:"disk,omitempty"`

	/* mem
	 */
	Mem int64 `json:"mem,omitempty"`

	/* time
	 */
	Time string `json:"time,omitempty"`
}

// Validate validates this l r p instance stats
func (m *LRPInstanceStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
