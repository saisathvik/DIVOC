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
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Registration API",
    "title": "Registration API",
    "version": "1.0.0"
  },
  "basePath": "/divoc/api/citizen",
  "paths": {
    "/appointment": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Book a slot in facility",
        "operationId": "bookSlotOfFacility",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "facilitySlotId",
                "enrollmentCode"
              ],
              "properties": {
                "enrollmentCode": {
                  "type": "string"
                },
                "facilitySlotId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Delete the appointment",
        "operationId": "deleteAppointment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "enrollmentCode"
              ],
              "properties": {
                "enrollmentCode": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/slots": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Get slots for facilites",
        "operationId": "getSlotsForFacilities",
        "parameters": [
          {
            "type": "string",
            "name": "facilityId",
            "in": "query"
          },
          {
            "type": "number",
            "default": 0,
            "name": "pageNumber",
            "in": "query"
          },
          {
            "type": "number",
            "default": 0,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/slots/init": {
      "post": {
        "security": [],
        "summary": "Initialize facility slots",
        "operationId": "initializeFacilitySlots",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "apiKey": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/generateOTP": {
      "post": {
        "security": [],
        "summary": "Generate OTP",
        "operationId": "generateOTP",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "phone": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "204": {
            "description": "Phone number is empty"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/recipients": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Get all the recipients",
        "operationId": "getRecipients",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "../registry/Enrollment.json#/definitions/Enrollment"
              }
            }
          },
          "401": {
            "description": "Invalid token"
          },
          "500": {
            "description": "Something went wrong"
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Enroll Recipient",
        "operationId": "enrollRecipient",
        "parameters": [
          {
            "description": "Recipient Details",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "../registry/Enrollment.json#/definitions/Enrollment"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Invalid token"
          }
        }
      }
    },
    "/verifyOTP": {
      "post": {
        "security": [],
        "summary": "Verify OTP",
        "operationId": "verifyOTP",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "otp": {
                  "type": "string"
                },
                "phone": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Bad request"
          },
          "401": {
            "description": "Invalid OTP"
          },
          "429": {
            "description": "Verify otp attempts exceeded, generate new OTP"
          },
          "500": {
            "description": "Internal error"
          }
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
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Registration API",
    "title": "Registration API",
    "version": "1.0.0"
  },
  "basePath": "/divoc/api/citizen",
  "paths": {
    "/appointment": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Book a slot in facility",
        "operationId": "bookSlotOfFacility",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "facilitySlotId",
                "enrollmentCode"
              ],
              "properties": {
                "enrollmentCode": {
                  "type": "string"
                },
                "facilitySlotId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Delete the appointment",
        "operationId": "deleteAppointment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "enrollmentCode"
              ],
              "properties": {
                "enrollmentCode": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/slots": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Get slots for facilites",
        "operationId": "getSlotsForFacilities",
        "parameters": [
          {
            "type": "string",
            "name": "facilityId",
            "in": "query"
          },
          {
            "type": "number",
            "default": 0,
            "name": "pageNumber",
            "in": "query"
          },
          {
            "type": "number",
            "default": 0,
            "name": "pageSize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/facility/slots/init": {
      "post": {
        "security": [],
        "summary": "Initialize facility slots",
        "operationId": "initializeFacilitySlots",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "apiKey": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/generateOTP": {
      "post": {
        "security": [],
        "summary": "Generate OTP",
        "operationId": "generateOTP",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "phone": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "204": {
            "description": "Phone number is empty"
          },
          "400": {
            "description": "Bad request"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/recipients": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Get all the recipients",
        "operationId": "getRecipients",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "$ref": "#/definitions/enrollment"
              }
            }
          },
          "401": {
            "description": "Invalid token"
          },
          "500": {
            "description": "Something went wrong"
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "summary": "Enroll Recipient",
        "operationId": "enrollRecipient",
        "parameters": [
          {
            "description": "Recipient Details",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enrollment"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Invalid token"
          }
        }
      }
    },
    "/verifyOTP": {
      "post": {
        "security": [],
        "summary": "Verify OTP",
        "operationId": "verifyOTP",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "otp": {
                  "type": "string"
                },
                "phone": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Bad request"
          },
          "401": {
            "description": "Invalid OTP"
          },
          "429": {
            "description": "Verify otp attempts exceeded, generate new OTP"
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "address": {
      "description": "Indian address format",
      "type": "object",
      "title": "Address",
      "required": [
        "addressLine1",
        "district",
        "state",
        "pincode"
      ],
      "properties": {
        "addressLine1": {
          "description": "Address line 1",
          "type": "string",
          "title": "The address line 1",
          "default": "",
          "$id": "#/properties/address/properties/addressLine1"
        },
        "addressLine2": {
          "type": "string",
          "title": "The address2 schema",
          "$id": "#/properties/address/properties/addressLine2"
        },
        "district": {
          "type": "string",
          "title": "The district schema",
          "$id": "#/properties/address/properties/district"
        },
        "pincode": {
          "type": "string",
          "title": "The pincode schema",
          "$id": "#/properties/address/properties/pincode"
        },
        "state": {
          "description": "State of address",
          "type": "string",
          "title": "The state schema",
          "$id": "#/properties/address/properties/state",
          "examples": [
            "Karnataka"
          ]
        }
      },
      "examples": [
        {
          "addressLine1": "no. 23, some lane, some road",
          "addressLine2": "some nagar",
          "district": "bangalore south",
          "pincode": "560000",
          "state": "Karnataka"
        }
      ]
    },
    "enrollment": {
      "type": "object",
      "required": [
        "nationalId"
      ],
      "properties": {
        "address": {
          "$ref": "#/definitions/address"
        },
        "appointmentDate": {
          "type": "string",
          "format": "date"
        },
        "appointmentSlot": {
          "type": "string"
        },
        "beneficiaryPhone": {
          "type": "string"
        },
        "certified": {
          "type": "boolean",
          "default": false
        },
        "code": {
          "type": "string"
        },
        "comorbidities": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "dob": {
          "type": "string",
          "format": "date"
        },
        "email": {
          "type": "string"
        },
        "enrollmentScopeId": {
          "type": "string"
        },
        "gender": {
          "type": "string",
          "enum": [
            "Male",
            "Female",
            "Other"
          ]
        },
        "name": {
          "type": "string"
        },
        "nationalId": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "programId": {
          "type": "string"
        },
        "yob": {
          "type": "integer"
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
