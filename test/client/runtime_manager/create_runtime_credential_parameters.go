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

// NewCreateRuntimeCredentialParams creates a new CreateRuntimeCredentialParams object
// with the default values initialized.
func NewCreateRuntimeCredentialParams() *CreateRuntimeCredentialParams {
	var ()
	return &CreateRuntimeCredentialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuntimeCredentialParamsWithTimeout creates a new CreateRuntimeCredentialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRuntimeCredentialParamsWithTimeout(timeout time.Duration) *CreateRuntimeCredentialParams {
	var ()
	return &CreateRuntimeCredentialParams{

		timeout: timeout,
	}
}

// NewCreateRuntimeCredentialParamsWithContext creates a new CreateRuntimeCredentialParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRuntimeCredentialParamsWithContext(ctx context.Context) *CreateRuntimeCredentialParams {
	var ()
	return &CreateRuntimeCredentialParams{

		Context: ctx,
	}
}

// NewCreateRuntimeCredentialParamsWithHTTPClient creates a new CreateRuntimeCredentialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRuntimeCredentialParamsWithHTTPClient(client *http.Client) *CreateRuntimeCredentialParams {
	var ()
	return &CreateRuntimeCredentialParams{
		HTTPClient: client,
	}
}

/*CreateRuntimeCredentialParams contains all the parameters to send to the API endpoint
for the create runtime credential operation typically these are written to a http.Request
*/
type CreateRuntimeCredentialParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateRuntimeCredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create runtime credential params
func (o *CreateRuntimeCredentialParams) WithTimeout(timeout time.Duration) *CreateRuntimeCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runtime credential params
func (o *CreateRuntimeCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runtime credential params
func (o *CreateRuntimeCredentialParams) WithContext(ctx context.Context) *CreateRuntimeCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runtime credential params
func (o *CreateRuntimeCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runtime credential params
func (o *CreateRuntimeCredentialParams) WithHTTPClient(client *http.Client) *CreateRuntimeCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runtime credential params
func (o *CreateRuntimeCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create runtime credential params
func (o *CreateRuntimeCredentialParams) WithBody(body *models.OpenpitrixCreateRuntimeCredentialRequest) *CreateRuntimeCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create runtime credential params
func (o *CreateRuntimeCredentialParams) SetBody(body *models.OpenpitrixCreateRuntimeCredentialRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuntimeCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
