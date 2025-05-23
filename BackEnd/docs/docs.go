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
        "/AdminLogin": {
            "post": {
                "description": "Validates admin credentials (plain text match)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins"
                ],
                "summary": "Admin login",
                "parameters": [
                    {
                        "description": "Admin credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Admin.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/CancelBookingandUpdateSlot": {
            "put": {
                "description": "Cancels a booking by updating its status to \"Cancelled\" and marks the corresponding time slot (based on Booking_Time) as available (sets it to 1) in the Court_TimeSlots record.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Cancel a booking and update court time slot",
                "parameters": [
                    {
                        "description": "Cancel Booking Request",
                        "name": "cancelRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.CancelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Booking cancelled and slot updated successfully for Booking_ID: 123",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or Invalid Slot_Index",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Booking not found or Court TimeSlots not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to start transaction, database error, or transaction commit failed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/CreateBooking": {
            "post": {
                "description": "Creates a new booking after validating the existence of the customer, sport, and court.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Create a new booking",
                "parameters": [
                    {
                        "description": "Booking data",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.Bookings"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Booking record added successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Customer, sport, or court not found",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/CreateCourt": {
            "post": {
                "description": "Creates a new court and assigns time slots for bookings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Create a new court with associated time slots",
                "parameters": [
                    {
                        "description": "Court data",
                        "name": "court",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.Court"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Court created successfully",
                        "schema": {
                            "$ref": "#/definitions/Court.CourtCreationResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to create court",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/CreateSport": {
            "post": {
                "description": "Adds a new sport to the database if it does not already exist. Requires Sport_name as input.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sports"
                ],
                "summary": "Create a new sport record",
                "parameters": [
                    {
                        "description": "Sport object",
                        "name": "sport",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.Sport"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Sport record added successfully\"  example({\"message\": \"Sport record added successfully!!\", \"sport\": {\"Sport_ID\": 1, \"Sport_name\": \"Tennis\"}})",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Sport_name is required or the sport already exists or invalid request body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/Customer": {
            "post": {
                "description": "Adds a new customer to the database if they do not already exist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "summary": "Create a new customer",
                "parameters": [
                    {
                        "description": "Customer data",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Customer already exists"
                    },
                    "201": {
                        "description": "Customer record added successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request body"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/DeleteCourt": {
            "delete": {
                "description": "Deletes a court record from the database based on the court name.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Delete a court record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Court Name to be deleted",
                        "name": "court_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Court deleted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid court name",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Court not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/ListCourts": {
            "get": {
                "description": "Retrieves a list of all courts along with the corresponding sport names.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "List all courts with their associated sports",
                "responses": {
                    "200": {
                        "description": "List of courts and their associated sports\"  example([{\"court_name\": \"Court A\", \"sport_name\": \"Tennis\"}, {\"court_name\": \"Court B\", \"sport_name\": \"Basketball\"}])",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Court.CourtData"
                            }
                        }
                    },
                    "500": {
                        "description": "Database error while fetching courts",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ListSports": {
            "get": {
                "description": "Fetches all sports names from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sports"
                ],
                "summary": "Get a list of sports",
                "responses": {
                    "200": {
                        "description": "List of sport names",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch sports",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/UpdateCourtSlotandBooking": {
            "put": {
                "description": "Toggles the availability of a court time slot and, based on the provided customer email and sport name,",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Update court slot and create booking",
                "parameters": [
                    {
                        "description": "Court slot update request including Customer_email and Sport_name",
                        "name": "updateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DataBase.CourtUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Slot updated and booking created successfully for Court_ID: {Court_ID}, Slot_Index: {Slot_Index}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or Slot_Index out of range",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Court time slots, Customer, or Sport not found",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Database error or failed to update slot/booking",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/getCourts": {
            "get": {
                "description": "Fetches courts based on the selected sport and provides their availability status along with time slots.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Get court availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sport name",
                        "name": "sport",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of available courts with time slots",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/DataBase.CourtAvailability"
                            }
                        }
                    },
                    "400": {
                        "description": "Missing 'sport' query parameter",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Sport not found or no courts available",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/listBookings": {
            "get": {
                "description": "Retrieves a list of bookings for a customer by email. Returns booking details including court name, sport name, slot time, and booking status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "List bookings for a customer",
                "parameters": [
                    {
                        "type": "string",
                        "default": "john@example.com",
                        "description": "Customer email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of bookings for the customer\"  example([{\"booking_id\":1,\"court_name\":\"Court A\",\"sport_name\":\"Tennis\",\"slot_time\":\"10-11 AM\",\"booking_status\":\"Confirmed\"}])",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Bookings.BookingResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Email query parameter is required",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Customer not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Database error while fetching bookings",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/resetCourtSlots": {
            "put": {
                "description": "Sets every slot (08‑18h) back to **available** (value ` + "`" + `1` + "`" + `) for every court whose ` + "`" + `court_status == 1` + "`" + `.\u003cbr\u003e",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "courts"
                ],
                "summary": "Reset all time‑slots for available courts",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"Court A\"",
                        "description": "Reset a single court by name",
                        "name": "court_name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Slots reset successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Database error while updating slots",
                        "schema": {
                            "$ref": "#/definitions/DataBase.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Admin.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Bookings.BookingResponse": {
            "type": "object",
            "properties": {
                "booking_id": {
                    "type": "integer"
                },
                "booking_status": {
                    "type": "string"
                },
                "court_name": {
                    "type": "string"
                },
                "slot_time": {
                    "type": "string"
                },
                "sport_name": {
                    "type": "string"
                }
            }
        },
        "Court.CourtCreationResponse": {
            "type": "object",
            "properties": {
                "court": {
                    "$ref": "#/definitions/DataBase.Court"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "Court.CourtData": {
            "type": "object",
            "properties": {
                "court_name": {
                    "type": "string"
                },
                "sport_name": {
                    "type": "string"
                }
            }
        },
        "DataBase.Bookings": {
            "type": "object",
            "properties": {
                "Booking_ID": {
                    "type": "integer"
                },
                "Booking_Status": {
                    "type": "string"
                },
                "Booking_Time": {
                    "type": "integer"
                },
                "Court_ID": {
                    "type": "integer"
                },
                "Customer_ID": {
                    "type": "integer"
                },
                "Sport_ID": {
                    "type": "integer"
                },
                "court": {
                    "$ref": "#/definitions/DataBase.Court"
                },
                "customer": {
                    "$ref": "#/definitions/DataBase.Customer"
                },
                "sport": {
                    "$ref": "#/definitions/DataBase.Sport"
                }
            }
        },
        "DataBase.CancelRequest": {
            "type": "object",
            "properties": {
                "Booking_ID": {
                    "type": "integer"
                }
            }
        },
        "DataBase.Court": {
            "type": "object",
            "properties": {
                "Court_ID": {
                    "type": "integer"
                },
                "Court_Location": {
                    "type": "string"
                },
                "Court_Name": {
                    "type": "string"
                },
                "Court_Status": {
                    "type": "integer"
                },
                "Sport_id": {
                    "type": "integer"
                },
                "court_Capacity": {
                    "type": "integer"
                },
                "sport": {
                    "$ref": "#/definitions/DataBase.Sport"
                }
            }
        },
        "DataBase.CourtAvailability": {
            "type": "object",
            "properties": {
                "CourtID": {
                    "type": "integer"
                },
                "CourtName": {
                    "type": "string"
                },
                "CourtStatus": {
                    "type": "integer"
                },
                "Slots": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "DataBase.CourtUpdate": {
            "type": "object",
            "properties": {
                "Court_ID": {
                    "type": "integer"
                },
                "Court_Name": {
                    "type": "string"
                },
                "Customer_email": {
                    "type": "string"
                },
                "Slot_Index": {
                    "type": "integer"
                },
                "Sport_ID": {
                    "type": "string"
                },
                "Sport_name": {
                    "type": "string"
                }
            }
        },
        "DataBase.Customer": {
            "type": "object",
            "properties": {
                "Contact": {
                    "type": "string"
                },
                "Customer_ID": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "DataBase.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "DataBase.Sport": {
            "type": "object",
            "properties": {
                "Sport_ID": {
                    "type": "integer"
                },
                "Sport_name": {
                    "type": "string"
                },
                "sport_Description": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Court Booking API",
	Description:      "API for managing court bookings",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
