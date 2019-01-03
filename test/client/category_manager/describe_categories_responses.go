// Code generated by go-swagger; DO NOT EDIT.

package category_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeCategoriesReader is a Reader for the DescribeCategories structure.
type DescribeCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeCategoriesOK creates a DescribeCategoriesOK with default headers values
func NewDescribeCategoriesOK() *DescribeCategoriesOK {
	return &DescribeCategoriesOK{}
}

/*DescribeCategoriesOK handles this case with default header values.

A successful response.
*/
type DescribeCategoriesOK struct {
	Payload *models.OpenpitrixDescribeCategoriesResponse
}

func (o *DescribeCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/CategoryManager.DescribeCategories][%d] describeCategoriesOK  %+v", 200, o.Payload)
}

func (o *DescribeCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeCategoriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
