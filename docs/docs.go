// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-02-16 16:21:08.519063 -0700 MST m=+0.023555764

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is Rhombus Exchnage rates API",
        "title": "Swagger Exchange API",
        "contact": {
            "name": "API Support",
            "email": "enormousrage@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "1.0"
    },
    "host": "ethergram.tk:8080",
    "basePath": "/",
    "paths": {
        "/rate/{currency}": {
            "get": {
                "description": "Return exchange rate of specific currency group",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Return exchange rate of specific currency group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "currency",
                        "name": "currency",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ExchangeResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.ExchangeResponse": {
            "type": "object",
            "properties": {
                "rate": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
