// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "oev2001@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/pets": {
            "get": {
                "description": "Retrieves all pets",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pets"
                ],
                "summary": "Get all pets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sort field",
                        "name": "sortField",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Pet"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Pet"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/current-user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Получить текущего пользователя",
                "responses": {
                    "200": {
                        "description": "Текущий пользователь",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Аутентификация",
                "parameters": [
                    {
                        "description": "Аутентификация пользователя",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.signIn.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "accessToken",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.tokenResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "accessToken": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "consumes": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "Регистрация пользователя",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.signIn.signInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "format": "email"
                },
                "password": {
                    "type": "string",
                    "format": "password"
                }
            }
        },
        "http.tokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.Pet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "breedId": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "isAvailable": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "shelterName": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Eugene"
                },
                "preferences": {
                    "description": "Предпочтения",
                    "type": "string",
                    "example": "some text about my preferences"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:12001",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Pet adoption platform api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
