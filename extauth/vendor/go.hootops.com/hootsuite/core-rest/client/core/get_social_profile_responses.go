// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "go.hootops.com/hootsuite/core-rest/models"
)

// GetSocialProfileReader is a Reader for the GetSocialProfile structure.
type GetSocialProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSocialProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSocialProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSocialProfileOK creates a GetSocialProfileOK with default headers values
func NewGetSocialProfileOK() *GetSocialProfileOK {
	return &GetSocialProfileOK{}
}

/*GetSocialProfileOK handles this case with default header values.

GetSocialProfileOK get social profile o k
*/
type GetSocialProfileOK struct {
	Payload *models.CoreSocialProfileResponse
}

func (o *GetSocialProfileOK) Error() string {
	return fmt.Sprintf("[GET /socialProfiles/{socialProfileId}][%d] getSocialProfileOK  %+v", 200, o.Payload)
}

func (o *GetSocialProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreSocialProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}