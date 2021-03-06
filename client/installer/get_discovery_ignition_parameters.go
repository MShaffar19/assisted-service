// Code generated by go-swagger; DO NOT EDIT.

package installer

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
)

// NewGetDiscoveryIgnitionParams creates a new GetDiscoveryIgnitionParams object
// with the default values initialized.
func NewGetDiscoveryIgnitionParams() *GetDiscoveryIgnitionParams {
	var ()
	return &GetDiscoveryIgnitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiscoveryIgnitionParamsWithTimeout creates a new GetDiscoveryIgnitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDiscoveryIgnitionParamsWithTimeout(timeout time.Duration) *GetDiscoveryIgnitionParams {
	var ()
	return &GetDiscoveryIgnitionParams{

		timeout: timeout,
	}
}

// NewGetDiscoveryIgnitionParamsWithContext creates a new GetDiscoveryIgnitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDiscoveryIgnitionParamsWithContext(ctx context.Context) *GetDiscoveryIgnitionParams {
	var ()
	return &GetDiscoveryIgnitionParams{

		Context: ctx,
	}
}

// NewGetDiscoveryIgnitionParamsWithHTTPClient creates a new GetDiscoveryIgnitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDiscoveryIgnitionParamsWithHTTPClient(client *http.Client) *GetDiscoveryIgnitionParams {
	var ()
	return &GetDiscoveryIgnitionParams{
		HTTPClient: client,
	}
}

/*GetDiscoveryIgnitionParams contains all the parameters to send to the API endpoint
for the get discovery ignition operation typically these are written to a http.Request
*/
type GetDiscoveryIgnitionParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) WithTimeout(timeout time.Duration) *GetDiscoveryIgnitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) WithContext(ctx context.Context) *GetDiscoveryIgnitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) WithHTTPClient(client *http.Client) *GetDiscoveryIgnitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) WithClusterID(clusterID strfmt.UUID) *GetDiscoveryIgnitionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get discovery ignition params
func (o *GetDiscoveryIgnitionParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiscoveryIgnitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
