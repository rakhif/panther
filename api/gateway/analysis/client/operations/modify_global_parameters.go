// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// NewModifyGlobalParams creates a new ModifyGlobalParams object
// with the default values initialized.
func NewModifyGlobalParams() *ModifyGlobalParams {
	var ()
	return &ModifyGlobalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyGlobalParamsWithTimeout creates a new ModifyGlobalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyGlobalParamsWithTimeout(timeout time.Duration) *ModifyGlobalParams {
	var ()
	return &ModifyGlobalParams{

		timeout: timeout,
	}
}

// NewModifyGlobalParamsWithContext creates a new ModifyGlobalParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyGlobalParamsWithContext(ctx context.Context) *ModifyGlobalParams {
	var ()
	return &ModifyGlobalParams{

		Context: ctx,
	}
}

// NewModifyGlobalParamsWithHTTPClient creates a new ModifyGlobalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyGlobalParamsWithHTTPClient(client *http.Client) *ModifyGlobalParams {
	var ()
	return &ModifyGlobalParams{
		HTTPClient: client,
	}
}

/*ModifyGlobalParams contains all the parameters to send to the API endpoint
for the modify global operation typically these are written to a http.Request
*/
type ModifyGlobalParams struct {

	/*Body*/
	Body *models.UpdateGlobal

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify global params
func (o *ModifyGlobalParams) WithTimeout(timeout time.Duration) *ModifyGlobalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify global params
func (o *ModifyGlobalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify global params
func (o *ModifyGlobalParams) WithContext(ctx context.Context) *ModifyGlobalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify global params
func (o *ModifyGlobalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify global params
func (o *ModifyGlobalParams) WithHTTPClient(client *http.Client) *ModifyGlobalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify global params
func (o *ModifyGlobalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the modify global params
func (o *ModifyGlobalParams) WithBody(body *models.UpdateGlobal) *ModifyGlobalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify global params
func (o *ModifyGlobalParams) SetBody(body *models.UpdateGlobal) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyGlobalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
