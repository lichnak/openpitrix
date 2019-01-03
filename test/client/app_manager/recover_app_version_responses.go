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

// RecoverAppVersionReader is a Reader for the RecoverAppVersion structure.
type RecoverAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecoverAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRecoverAppVersionOK creates a RecoverAppVersionOK with default headers values
func NewRecoverAppVersionOK() *RecoverAppVersionOK {
	return &RecoverAppVersionOK{}
}

/*RecoverAppVersionOK handles this case with default header values.

A successful response.
*/
type RecoverAppVersionOK struct {
	Payload *models.OpenpitrixRecoverAppVersionResponse
}

func (o *RecoverAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /api/AppManager.RecoverAppVersion][%d] recoverAppVersionOK  %+v", 200, o.Payload)
}

func (o *RecoverAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixRecoverAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
