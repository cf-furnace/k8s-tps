package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBulkLRPStatusParams creates a new BulkLRPStatusParams object
// with the default values initialized.
func NewBulkLRPStatusParams() BulkLRPStatusParams {
	var ()
	return BulkLRPStatusParams{}
}

// BulkLRPStatusParams contains all the bound params for the bulk l r p status operation
// typically these are obtained from a http.Request
//
// swagger:parameters bulkLRPStatus
type BulkLRPStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*comma separated list of guids
	  Required: true
	  In: query
	*/
	Guids string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *BulkLRPStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qGuids, qhkGuids, _ := qs.GetOK("guids")
	if err := o.bindGuids(qGuids, qhkGuids, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BulkLRPStatusParams) bindGuids(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("guids", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("guids", "query", raw); err != nil {
		return err
	}

	o.Guids = raw

	return nil
}
