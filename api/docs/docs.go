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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/create-customer": {
            "post": {
                "description": "this func creates customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create customer with info",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    }
                }
            }
        },
        "/v1/create-post": {
            "post": {
                "description": "this func creates post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "create post with info",
                "parameters": [
                    {
                        "description": "Post",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.PostReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    }
                }
            }
        },
        "/v1/create-ranking": {
            "post": {
                "description": "this func creates ranking",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reyting"
                ],
                "summary": "create reyting",
                "parameters": [
                    {
                        "description": "Ranking",
                        "name": "ranking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reyting.Ranking"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reyting.Empty"
                        }
                    }
                }
            }
        },
        "/v1/delete-customer/{id}": {
            "delete": {
                "description": "this func delete customer by customer id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "delete customer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/delete-post/{id}": {
            "delete": {
                "description": "this func delete post by post id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "delete post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/get-customer/{id}": {
            "get": {
                "description": "this func get customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "get customer with info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/get-post/{id}": {
            "get": {
                "description": "this func get post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "get post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/getpostlist/{id}": {
            "get": {
                "description": "this func get list of customers by post id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "get post list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/getranking-customer/{id}": {
            "get": {
                "description": "this func get rankings of posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reyting"
                ],
                "summary": "get reyting",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reyting.Rankings"
                        }
                    }
                }
            }
        },
        "/v1/getranking-post/{id}": {
            "get": {
                "description": "this func get rankings of posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reyting"
                ],
                "summary": "get reyting",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reyting.Rankings"
                        }
                    }
                }
            }
        },
        "/v1/list-customer": {
            "get": {
                "description": "this func get list of customers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "get customer list",
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/list-post": {
            "get": {
                "description": "this func get list of posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "get post list",
                "responses": {
                    "200": {
                        "description": "success"
                    }
                }
            }
        },
        "/v1/update-customer": {
            "put": {
                "description": "this func update customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "update customer",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    }
                }
            }
        },
        "/v1/update-post": {
            "put": {
                "description": "this func update post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "update post",
                "parameters": [
                    {
                        "description": "Post",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "customer.Address": {
            "type": "object",
            "properties": {
                "district": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "customer.Customer": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Address"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Post"
                    }
                },
                "rankings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Ranking"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "customer.CustomerRequest": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Address"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "customer.Post": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "customer.Ranking": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                },
                "ranking": {
                    "type": "integer"
                }
            }
        },
        "post.Media": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Media"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "post.PostReq": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Media"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "reyting.Empty": {
            "type": "object"
        },
        "reyting.Ranking": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                },
                "ranking": {
                    "type": "integer"
                }
            }
        },
        "reyting.Rankings": {
            "type": "object",
            "properties": {
                "rankings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/reyting.Ranking"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
