// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Tokens WebSite Backend Service",
    "title": "Tokens Backend Service",
    "version": "development"
  },
  "paths": {
    "/auth/refresh": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "New tokens by active refresh token",
        "summary": "Update access and refresh tokens",
        "parameters": [
          {
            "description": "Refresh Token Body",
            "name": "RefreshTokenBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RefreshTokenBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Tokens Response",
            "schema": {
              "$ref": "#/definitions/AccessTokenBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Invalid IP or token mismatch",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Refresh token reuse attempt",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/auth/token": {
      "get": {
        "description": "Get access and refresh tokens by GUID",
        "summary": "Get access and refresh tokens",
        "parameters": [
          {
            "type": "string",
            "description": "GUID",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Tokens Response",
            "schema": {
              "$ref": "#/definitions/Tokens"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AccessTokenBody": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "readOnly": true
        },
        "ip": {
          "type": "string",
          "readOnly": true
        },
        "refresh_id": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "RefreshTokenBody": {
      "type": "object",
      "properties": {
        "refresh_token": {
          "type": "string"
        }
      }
    },
    "Tokens": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Tokens WebSite Backend Service",
    "title": "Tokens Backend Service",
    "version": "development"
  },
  "paths": {
    "/auth/refresh": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "New tokens by active refresh token",
        "summary": "Update access and refresh tokens",
        "parameters": [
          {
            "description": "Refresh Token Body",
            "name": "RefreshTokenBody",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RefreshTokenBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Tokens Response",
            "schema": {
              "$ref": "#/definitions/AccessTokenBody"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "Invalid IP or token mismatch",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "Refresh token reuse attempt",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/auth/token": {
      "get": {
        "description": "Get access and refresh tokens by GUID",
        "summary": "Get access and refresh tokens",
        "parameters": [
          {
            "type": "string",
            "description": "GUID",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Tokens Response",
            "schema": {
              "$ref": "#/definitions/Tokens"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AccessTokenBody": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Principal": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "readOnly": true
        },
        "ip": {
          "type": "string",
          "readOnly": true
        },
        "refresh_id": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "RefreshTokenBody": {
      "type": "object",
      "properties": {
        "refresh_token": {
          "type": "string"
        }
      }
    },
    "Tokens": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
