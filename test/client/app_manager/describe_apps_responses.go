// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeAppsReader is a Reader for the DescribeApps structure.
type DescribeAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAppsOK creates a DescribeAppsOK with default headers values
func NewDescribeAppsOK() *DescribeAppsOK {
	return &DescribeAppsOK{}
}

/*DescribeAppsOK handles this case with default header values.

A successful response.
*/
type DescribeAppsOK struct {
	Payload *models.OpenpitrixDescribeAppsResponse
}

func (o *DescribeAppsOK) Error() string {
	return fmt.Sprintf("[GET /api/AppManager.DescribeApps][%d] describeAppsOK  %+v", 200, o.Payload)
}

func (o *DescribeAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
