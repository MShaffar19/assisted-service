// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DisableNodeReader is a Reader for the DisableNode structure.
type DisableNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisableNodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDisableNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDisableNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDisableNodeNoContent creates a DisableNodeNoContent with default headers values
func NewDisableNodeNoContent() *DisableNodeNoContent {
	return &DisableNodeNoContent{}
}

/*DisableNodeNoContent handles this case with default header values.

Disabled
*/
type DisableNodeNoContent struct {
}

func (o *DisableNodeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{node_id}/actions/enable][%d] disableNodeNoContent ", 204)
}

func (o *DisableNodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableNodeNotFound creates a DisableNodeNotFound with default headers values
func NewDisableNodeNotFound() *DisableNodeNotFound {
	return &DisableNodeNotFound{}
}

/*DisableNodeNotFound handles this case with default header values.

Node not found
*/
type DisableNodeNotFound struct {
}

func (o *DisableNodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{node_id}/actions/enable][%d] disableNodeNotFound ", 404)
}

func (o *DisableNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableNodeInternalServerError creates a DisableNodeInternalServerError with default headers values
func NewDisableNodeInternalServerError() *DisableNodeInternalServerError {
	return &DisableNodeInternalServerError{}
}

/*DisableNodeInternalServerError handles this case with default header values.

Internal server error
*/
type DisableNodeInternalServerError struct {
}

func (o *DisableNodeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{node_id}/actions/enable][%d] disableNodeInternalServerError ", 500)
}

func (o *DisableNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}