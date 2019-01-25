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

// DevelopAdminPassAppVersionReader is a Reader for the DevelopAdminPassAppVersion structure.
type DevelopAdminPassAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevelopAdminPassAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDevelopAdminPassAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDevelopAdminPassAppVersionOK creates a DevelopAdminPassAppVersionOK with default headers values
func NewDevelopAdminPassAppVersionOK() *DevelopAdminPassAppVersionOK {
	return &DevelopAdminPassAppVersionOK{}
}

/*DevelopAdminPassAppVersionOK handles this case with default header values.

A successful response.
*/
type DevelopAdminPassAppVersionOK struct {
	Payload *models.OpenpitrixPassAppVersionResponse
}

func (o *DevelopAdminPassAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_version/action/develop_admin/pass][%d] developAdminPassAppVersionOK  %+v", 200, o.Payload)
}

func (o *DevelopAdminPassAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixPassAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
