// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/terms/",
        "contact": {
            "name": "Hacktiv8",
            "email": "admin@hacktiv8.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "description": "Get All orders in detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get All Orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.GetOrderSuccess"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/views.GetOrderFailedNotFound"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Order": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customerName": {
                    "type": "string",
                    "example": "MNC GOLANG LESSON"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "productsId": {
                    "type": "integer",
                    "example": 1
                },
                "userId": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "views.GetOrderFailedNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "GET_NOT_FOUND"
                },
                "status": {
                    "type": "integer",
                    "example": 404
                }
            }
        },
        "views.GetOrderSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "GET_SUCCESS"
                },
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Orders API",
	Description:      "Ini adalah sample API Spec untuk Api orders",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
