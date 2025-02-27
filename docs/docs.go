// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/db-health": {
            "get": {
                "description": "Test if database ping is working",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Test database Health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates a user and retrieve a JWT for API calls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Provides a JSON Web Token",
                "operationId": "Authentication",
                "parameters": [
                    {
                        "description": "Credentials",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UserLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Test if no authenticated ping is working",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Ping No-Authenticated",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    }
                }
            }
        },
        "/v1/breeds": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Search on API and DB for cat breeds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search for cat breeds",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CatBreed"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    }
                }
            }
        },
        "/v1/ping": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Test if authenticated ping is working",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Ping with Authentication",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.HttpDefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helpers.HttpDefaultResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.CatBreed": {
            "type": "object",
            "properties": {
                "adaptability": {
                    "type": "integer"
                },
                "affection_level": {
                    "type": "integer"
                },
                "alt_names": {
                    "type": "string"
                },
                "cfa_url": {
                    "type": "string"
                },
                "child_friendly": {
                    "type": "integer"
                },
                "country_code": {
                    "type": "string"
                },
                "country_codes": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "dog_friendly": {
                    "type": "integer"
                },
                "energy_level": {
                    "type": "integer"
                },
                "experimental": {
                    "type": "integer"
                },
                "grooming": {
                    "type": "integer"
                },
                "hairless": {
                    "type": "integer"
                },
                "health_issues": {
                    "type": "integer"
                },
                "hypoallergenic": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "indoor": {
                    "type": "integer"
                },
                "intelligence": {
                    "type": "integer"
                },
                "lap": {
                    "type": "integer"
                },
                "life_span": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "natural": {
                    "type": "integer"
                },
                "origin": {
                    "type": "string"
                },
                "rare": {
                    "type": "integer"
                },
                "reference_image_id": {
                    "type": "string"
                },
                "rex": {
                    "type": "integer"
                },
                "shedding_level": {
                    "type": "integer"
                },
                "short_legs": {
                    "type": "integer"
                },
                "social_needs": {
                    "type": "integer"
                },
                "stranger_friendly": {
                    "type": "integer"
                },
                "suppressed_tail": {
                    "type": "integer"
                },
                "temperament": {
                    "type": "string"
                },
                "vcahospitals_url": {
                    "type": "string"
                },
                "vetstreet_url": {
                    "type": "string"
                },
                "vocalisation": {
                    "type": "integer"
                },
                "wikipedia_url": {
                    "type": "string"
                }
            }
        },
        "requests.UserLoginDto": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
