// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

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

// NewResizeClusterParams creates a new ResizeClusterParams object
// with the default values initialized.
func NewResizeClusterParams() *ResizeClusterParams {
	var ()
	return &ResizeClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResizeClusterParamsWithTimeout creates a new ResizeClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResizeClusterParamsWithTimeout(timeout time.Duration) *ResizeClusterParams {
	var ()
	return &ResizeClusterParams{

		timeout: timeout,
	}
}

// NewResizeClusterParamsWithContext creates a new ResizeClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewResizeClusterParamsWithContext(ctx context.Context) *ResizeClusterParams {
	var ()
	return &ResizeClusterParams{

		Context: ctx,
	}
}

// NewResizeClusterParamsWithHTTPClient creates a new ResizeClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResizeClusterParamsWithHTTPClient(client *http.Client) *ResizeClusterParams {
	var ()
	return &ResizeClusterParams{
		HTTPClient: client,
	}
}

/*ResizeClusterParams contains all the parameters to send to the API endpoint
for the resize cluster operation typically these are written to a http.Request
*/
type ResizeClusterParams struct {

	/*Body*/
	Body *models.OpenpitrixResizeClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resize cluster params
func (o *ResizeClusterParams) WithTimeout(timeout time.Duration) *ResizeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resize cluster params
func (o *ResizeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resize cluster params
func (o *ResizeClusterParams) WithContext(ctx context.Context) *ResizeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resize cluster params
func (o *ResizeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resize cluster params
func (o *ResizeClusterParams) WithHTTPClient(client *http.Client) *ResizeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resize cluster params
func (o *ResizeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resize cluster params
func (o *ResizeClusterParams) WithBody(body *models.OpenpitrixResizeClusterRequest) *ResizeClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resize cluster params
func (o *ResizeClusterParams) SetBody(body *models.OpenpitrixResizeClusterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResizeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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