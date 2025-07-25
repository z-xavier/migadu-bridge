// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/alias/random/new": {
            "post": {
                "description": "Create a new random email alias using the SimpleLogin API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bridges"
                ],
                "summary": "Create a random alias for SimpleLogin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Alias mode",
                        "name": "mode",
                        "in": "query"
                    },
                    {
                        "description": "Alias creation request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/sl.AliasRandomNewReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/sl.Alias"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/sl.ErrorResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/sl.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/sl.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/aliases": {
            "get": {
                "description": "Get a list of aliases with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aliases"
                ],
                "summary": "List aliases",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 200,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Order by fields",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.ListAliasResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new email alias using the Addy.io API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bridges"
                ],
                "summary": "Create an alias for Addy.io",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Alias creation request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/addy.CreateAliasReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/addy.CreateAliasResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/addy.ErrorResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/addy.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/addy.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/calllogs": {
            "get": {
                "description": "Get a list of call logs with pagination and filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "call-logs"
                ],
                "summary": "List call logs",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 200,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Order by fields",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by target email",
                        "name": "targetEmail",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by mock provider",
                        "name": "mockProvider",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by description",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by request path",
                        "name": "requestPath",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by request IP",
                        "name": "requestIp",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by request time (begin)",
                        "name": "requestAtBegin",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by request time (end)",
                        "name": "requestAtEnd",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.ListCallLogResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tokens": {
            "get": {
                "description": "Get a list of tokens with pagination and filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "List tokens",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 200,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Order by fields",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by target email",
                        "name": "targetEmail",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by mock provider",
                        "name": "mockProvider",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by description",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by expiry time (begin)",
                        "name": "expiryAtBegin",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by expiry time (end)",
                        "name": "expiryAtEnd",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by last called time (begin)",
                        "name": "lastCalledAtBegin",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by last called time (end)",
                        "name": "lastCalledAtEnd",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by updated time (begin)",
                        "name": "updatedAtBegin",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by updated time (end)",
                        "name": "updatedAtEnd",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.ListTokenResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new token for API access",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Create a new token",
                "parameters": [
                    {
                        "description": "Token creation request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CreateTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.Token"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tokens/{tokenId}": {
            "get": {
                "description": "Get a token by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Get a token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "tokenId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.Token"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "404": {
                        "description": "Token not found",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a token by ID (full update)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Update a token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "tokenId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Token update request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PutTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.Token"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "404": {
                        "description": "Token not found",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a token by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Delete a token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "tokenId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "404": {
                        "description": "Token not found",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Partially update a token by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tokens"
                ],
                "summary": "Partially update a token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token ID",
                        "name": "tokenId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Token patch request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PatchTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/v1.Token"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "404": {
                        "description": "Token not found",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "addy.Alias": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "aliasable_id": {},
                "aliasable_type": {},
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "emails_blocked": {
                    "type": "integer"
                },
                "emails_forwarded": {
                    "type": "integer"
                },
                "emails_replied": {
                    "type": "integer"
                },
                "emails_sent": {
                    "type": "integer"
                },
                "extension": {},
                "from_name": {},
                "id": {
                    "type": "string"
                },
                "last_blocked": {},
                "last_forwarded": {
                    "type": "string"
                },
                "last_replied": {},
                "last_sent": {},
                "local_part": {
                    "type": "string"
                },
                "recipients": {
                    "type": "array",
                    "items": {}
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "addy.AliasFormat": {
            "type": "string",
            "enum": [
                "random_characters",
                "uuid",
                "random_words",
                "custom"
            ],
            "x-enum-varnames": [
                "AliasFormatRandomCharacters",
                "AliasFormatUUID",
                "AliasFormatRandomWords",
                "AliasFormatCustom"
            ]
        },
        "addy.CreateAliasReq": {
            "type": "object",
            "properties": {
                "authorization": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "format": {
                    "$ref": "#/definitions/addy.AliasFormat"
                },
                "local_part": {
                    "type": "string"
                },
                "recipient_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "xrequestedWith": {
                    "type": "string"
                }
            }
        },
        "addy.CreateAliasResp": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/addy.Alias"
                }
            }
        },
        "addy.ErrorResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "enum.ProviderEnum": {
            "type": "string",
            "enum": [
                "addy",
                "sl"
            ],
            "x-enum-varnames": [
                "ProviderEnumAddy",
                "ProviderEnumSimpleLogin"
            ]
        },
        "enum.TokenStatus": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "TokenStatusInactive",
                "TokenStatusActive",
                "TokenStatusPause"
            ]
        },
        "sl.Alias": {
            "type": "object",
            "properties": {
                "creation_date": {
                    "type": "string"
                },
                "creation_timestamp": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "latest_activity": {
                    "$ref": "#/definitions/sl.LatestActivity"
                },
                "mailbox": {
                    "$ref": "#/definitions/sl.MailBox"
                },
                "mailboxes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/sl.MailBox"
                    }
                },
                "name": {
                    "type": "string"
                },
                "nb_block": {
                    "type": "integer"
                },
                "nb_forward": {
                    "type": "integer"
                },
                "nb_reply": {
                    "type": "integer"
                },
                "note": {},
                "pinned": {
                    "type": "boolean"
                }
            }
        },
        "sl.AliasRandomNewReq": {
            "type": "object",
            "properties": {
                "authentication": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                },
                "word": {
                    "type": "string"
                }
            }
        },
        "sl.ErrorResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "sl.LatestActivity": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "contact": {
                    "$ref": "#/definitions/sl.LatestActivityContact"
                },
                "timestamp": {
                    "type": "integer"
                }
            }
        },
        "sl.LatestActivityContact": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {},
                "reverse_alias": {
                    "type": "string"
                }
            }
        },
        "sl.MailBox": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.Alias": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "callLogId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mockProvider": {
                    "$ref": "#/definitions/enum.ProviderEnum"
                },
                "targetEmail": {
                    "type": "string"
                },
                "tokenId": {
                    "type": "string"
                }
            }
        },
        "v1.CallLog": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "genAlias": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "mockProvider": {
                    "$ref": "#/definitions/enum.ProviderEnum"
                },
                "requestAt": {
                    "type": "integer"
                },
                "requestIp": {
                    "type": "string"
                },
                "requestPath": {
                    "type": "string"
                },
                "targetEmail": {
                    "type": "string"
                },
                "tokenId": {
                    "type": "string"
                }
            }
        },
        "v1.CreateTokenReq": {
            "type": "object",
            "required": [
                "expiryAt",
                "mockProvider",
                "targetEmail"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 1024
                },
                "expiryAt": {
                    "type": "integer",
                    "minimum": 0
                },
                "mockProvider": {
                    "$ref": "#/definitions/enum.ProviderEnum"
                },
                "targetEmail": {
                    "type": "string"
                }
            }
        },
        "v1.ListAliasResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Alias"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "v1.ListCallLogResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.CallLog"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "v1.ListTokenResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Token"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "v1.PatchTokenReq": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/enum.TokenStatus"
                }
            }
        },
        "v1.PutTokenReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "expiryAt": {
                    "type": "integer"
                }
            }
        },
        "v1.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code 指定了业务错误码.",
                    "type": "string"
                },
                "data": {
                    "description": "Data 包含了"
                },
                "message": {
                    "description": "Message 包含了可以直接对外展示的错误信息.",
                    "type": "string"
                }
            }
        },
        "v1.Token": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "expiryAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "lastCalledAt": {
                    "type": "integer"
                },
                "mockProvider": {
                    "$ref": "#/definitions/enum.ProviderEnum"
                },
                "status": {
                    "$ref": "#/definitions/enum.TokenStatus"
                },
                "targetEmail": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
