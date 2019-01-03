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

// DescribeAppVersionsReader is a Reader for the DescribeAppVersions structure.
type DescribeAppVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAppVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAppVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAppVersionsOK creates a DescribeAppVersionsOK with default headers values
func NewDescribeAppVersionsOK() *DescribeAppVersionsOK {
	return &DescribeAppVersionsOK{}
}

/*DescribeAppVersionsOK handles this case with default header values.

A successful response.
*/
type DescribeAppVersionsOK struct {
	Payload *models.OpenpitrixDescribeAppVersionsResponse
}

func (o *DescribeAppVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/AppManager.DescribeAppVersions][%d] describeAppVersionsOK  %+v", 200, o.Payload)
}

func (o *DescribeAppVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
