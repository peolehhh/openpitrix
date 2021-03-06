// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewModifyRuntimeCredentialParams creates a new ModifyRuntimeCredentialParams object
// with the default values initialized.
func NewModifyRuntimeCredentialParams() *ModifyRuntimeCredentialParams {
	var ()
	return &ModifyRuntimeCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyRuntimeCredentialParamsWithTimeout creates a new ModifyRuntimeCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyRuntimeCredentialParamsWithTimeout(timeout time.Duration) *ModifyRuntimeCredentialParams {
	var ()
	return &ModifyRuntimeCredentialParams{

		timeout: timeout,
	}
}

// NewModifyRuntimeCredentialParamsWithContext creates a new ModifyRuntimeCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyRuntimeCredentialParamsWithContext(ctx context.Context) *ModifyRuntimeCredentialParams {
	var ()
	return &ModifyRuntimeCredentialParams{

		Context: ctx,
	}
}

// NewModifyRuntimeCredentialParamsWithHTTPClient creates a new ModifyRuntimeCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyRuntimeCredentialParamsWithHTTPClient(client *http.Client) *ModifyRuntimeCredentialParams {
	var ()
	return &ModifyRuntimeCredentialParams{
		HTTPClient: client,
	}
}

/*ModifyRuntimeCredentialParams contains all the parameters to send to the API endpoint
for the modify runtime credential operation typically these are written to a http.Request
*/
type ModifyRuntimeCredentialParams struct {

	/*Body*/
	Body *models.OpenpitrixModifyRuntimeCredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) WithTimeout(timeout time.Duration) *ModifyRuntimeCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) WithContext(ctx context.Context) *ModifyRuntimeCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) WithHTTPClient(client *http.Client) *ModifyRuntimeCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) WithBody(body *models.OpenpitrixModifyRuntimeCredentialRequest) *ModifyRuntimeCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify runtime credential params
func (o *ModifyRuntimeCredentialParams) SetBody(body *models.OpenpitrixModifyRuntimeCredentialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyRuntimeCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
