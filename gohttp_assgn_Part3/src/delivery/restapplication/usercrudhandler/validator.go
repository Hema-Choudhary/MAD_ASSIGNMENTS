package usercrudhandler

import (
	"bytes"
	"fmt"
	logger "log"
	customerrors "delivery/restapplication/packages/errors"
	"github.com/xeipuuv/gojsonschema"
)

func ValidateUserCreateUpdateRequest(rStr string) (bool, error) {
	//logger := loggerutils.GetLogger()
	schemaStr := `
	{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "object",
		"properties": {
		  "name": {
    			"type": "string",
			"minLength": 1,
			"maxLength": 60
		  },
		  "address": {
            		"type": "string",
			"minLength": 1,
			"maxLength": 120
		  },
		  "addressLine2": {
            		"type": "string",
			"minimum": 1,
			"maximum": 120
		  },
		  "url": {
            		"type": "string",
			"minimum": 1,
			"maximum": 150
		  },
		  "outcode": {
            		"type": "string",
			"minimum": 1,
			"maximum": 30
		  },
		  "postcode": {
            		"type": "string",
			"minimum": 1,
			"maximum": 30
		  },
		  "rating": {
            		"type": "float32",
			"minimum": 1,
			"maximum": 5
		  },
		  "typeOfFood": {
            		"type": "string",
			"minimum": 1,
			"maximum": 120
		  }
		  
		},
		"required": [
		  "name",
		  "address",
		  "rating"
		]
	  }`

	schema := gojsonschema.NewStringLoader(schemaStr)
	content := gojsonschema.NewStringLoader(rStr)
	result, err := gojsonschema.Validate(schema, content)

	if err != nil {
		logger.Fatalf("Invalid Json Schema Error: %v", err)
		return false, customerrors.InternalError(fmt.Sprintf("Invalid Json Schema Error: %v", err))
		//panic(err)
	}
	if result.Valid() {
		return true, nil
	} else {
		var buffer bytes.Buffer
		for _, resulterr := range result.Errors() {
			logger.Printf("- %s\n", resulterr)
			errString := fmt.Sprintf("Field: %s - %s, ", resulterr.Field(), resulterr.Description())
			buffer.Write([]byte(errString))

		}
		errorDesc := buffer.String()
		return false, customerrors.BadRequest(errorDesc)
	}

}
