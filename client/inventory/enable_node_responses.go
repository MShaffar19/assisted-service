// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnableNodeReader is a Reader for the EnableNode structure.
type EnableNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEnableNodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnableNodeNoContent creates a EnableNodeNoContent with default headers values
func NewEnableNodeNoContent() *EnableNodeNoContent {
	return &EnableNodeNoContent{}
}

/*EnableNodeNoContent handles this case with default header values.

Enabled
*/
type EnableNodeNoContent struct {
}

func (o *EnableNodeNoContent) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/actions/enable][%d] enableNodeNoContent ", 204)
}

func (o *EnableNodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableNodeNotFound creates a EnableNodeNotFound with default headers values
func NewEnableNodeNotFound() *EnableNodeNotFound {
	return &EnableNodeNotFound{}
}

/*EnableNodeNotFound handles this case with default header values.

Node not found
*/
type EnableNodeNotFound struct {
}

func (o *EnableNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/actions/enable][%d] enableNodeNotFound ", 404)
}

func (o *EnableNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableNodeInternalServerError creates a EnableNodeInternalServerError with default headers values
func NewEnableNodeInternalServerError() *EnableNodeInternalServerError {
	return &EnableNodeInternalServerError{}
}

/*EnableNodeInternalServerError handles this case with default header values.

Internal server error
*/
type EnableNodeInternalServerError struct {
}

func (o *EnableNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/actions/enable][%d] enableNodeInternalServerError ", 500)
}

func (o *EnableNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}