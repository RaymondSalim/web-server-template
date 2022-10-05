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
        "contact": {
            "name": "Raymond Salim",
            "url": "https://raymonds.dev/#contact",
            "email": "raymond+novometrix@raymonds.dev"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/counter/add": {
            "post": {
                "description": "Add Counter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "Counter"
                ],
                "summary": "Add counter",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FooResponse"
                        }
                    }
                }
            }
        },
        "/counter/get": {
            "get": {
                "description": "Get Last Counter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "Counter"
                ],
                "summary": "Get Last counter",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FooResponse"
                        }
                    }
                }
            }
        },
        "/foo/create": {
            "post": {
                "description": "Adds a new foo to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Foo"
                ],
                "summary": "Adds a new foo",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "FooRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddFoo"
                        }
                    },
                    {
                        "type": "string",
                        "default": "174b9d6a-dafe-4f68-8e4b-6dcfbe7a804e",
                        "description": "Request ID",
                        "name": "X-Request-ID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FooResponse"
                        }
                    }
                }
            }
        },
        "/foo/delete": {
            "post": {
                "description": "Delete foo from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Foo"
                ],
                "summary": "Delete foo",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "FooRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DeleteFoo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "865782e5-ccbf-4c5f-b967-f3df1fcd1f75",
                        "name": "X-Request-ID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FooResponse"
                        }
                    }
                }
            }
        },
        "/foo/get": {
            "post": {
                "description": "Get foo from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Foo"
                ],
                "summary": "Get foo",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "FooRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetFoo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "865782e5-ccbf-4c5f-b967-f3df1fcd1f75",
                        "name": "X-Request-ID",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.FooResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Foo": {
            "type": "object",
            "properties": {
                "fooName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "request.AddFoo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "request.DeleteFoo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "request.GetFoo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "response.FooResponse": {
            "type": "object",
            "properties": {
                "foo": {
                    "$ref": "#/definitions/models.Foo"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Web Server Template",
	Description:      "This is a template for Novometrix's web server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
