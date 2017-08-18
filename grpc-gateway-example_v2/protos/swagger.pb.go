package protos 

const(
swagger= `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/example/hello": {
      "post": {
        "operationId": "SayHello",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/protosReply"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/protosRequest"
            }
          }
        ],
        "tags": [
          "Greeter"
        ]
      }
    }
  },
  "definitions": {
    "protosReply": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "protosRequest": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    }
  }
}
`
)
