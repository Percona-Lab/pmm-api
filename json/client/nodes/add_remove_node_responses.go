// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/Percona-Lab/pmm-api/json/models"
)

// AddRemoveNodeReader is a Reader for the AddRemoveNode structure.
type AddRemoveNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRemoveNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddRemoveNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRemoveNodeOK creates a AddRemoveNodeOK with default headers values
func NewAddRemoveNodeOK() *AddRemoveNodeOK {
	return &AddRemoveNodeOK{}
}

/*AddRemoveNodeOK handles this case with default header values.

(empty)
*/
type AddRemoveNodeOK struct {
	Payload *models.InventoryAddRemoveNodeResponse
}

func (o *AddRemoveNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/AddRemoveNode][%d] addRemoveNodeOK  %+v", 200, o.Payload)
}

func (o *AddRemoveNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddRemoveNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
