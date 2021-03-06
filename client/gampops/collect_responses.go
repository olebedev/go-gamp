package gampops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CollectReader is a Reader for the Collect structure.
type CollectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewCollectDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewCollectDefault creates a CollectDefault with default headers values
func NewCollectDefault(code int) *CollectDefault {
	return &CollectDefault{
		_statusCode: code,
	}
}

/*CollectDefault handles this case with default header values.

CollectDefault collect default
*/
type CollectDefault struct {
	_statusCode int
}

// Code gets the status code for the collect default response
func (o *CollectDefault) Code() int {
	return o._statusCode
}

func (o *CollectDefault) Error() string {
	return fmt.Sprintf("[POST /collect][%d] Collect default ", o._statusCode)
}

func (o *CollectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
