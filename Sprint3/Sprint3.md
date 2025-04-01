# Team-Courtlink (Project Group 4) Sprint 3

### GitHub Repository Link: [CourtLink](https://github.com/Ashfaq-Ahmed-Mohammed/Team-CourtLink)

## Work Completed in Sprint 3

### BACKEND
Implemented an API to update court availability when a timeslot is booked, ensuring booking details are recorded in the Bookings table.

Optimized the getCourts API to reduce the number of database queries.

Developed an API that allows users to cancel a booking, which frees up the blocked timeslot and updates the court status in the Booking table.

Created an API that returns a JSON list of bookings for a particular customer to manage their reservations.

Added new unit tests for the CancelBookingAndUpdateSlot, UpdateCourtSlotAndBooking, and listBookings APIs, and collaborated with the front-end team to integrate the back-end with the Angular application flow while updating the database schema as needed.

Implemented an API to list courts which helps in showcasing the courts available before adding or removing courts to the admin portal.

Implemented an API to create court which is useful to add courts when there is a new court available.

Implemented an API to create sport which provides capability to add new sports which are introduced.

Implemented an API to list sports which helps in showcasing the sports available before adding or removing sports to admin portal and to sports list for main portal.

Implemented an API to delete court. This helps in removing the courts that are not needed.

Implemented logic which takes in emailID of the user and fetch the customer id which is used for the create booking logic.

Implemented logic to update booking status for delete.

Added unit test cases for Create Court, Create Sport, List Courts, List Sports and Delete Court and collaborated with front-end team to integrate angular components as per the need and enhancing the database schema as required.

Swagger documentation has been added to new API’s at BackEnd.

### FRONTEND

Admin Page :

An admin page has been added to make addition and deletion of courts and sports easier. It has the following components:

admin-sports component – The sports are tabulated and new sports can be added.

admin-courts component – The courts are tabulated against their respective sports. A filter search bar allows for easy navigation and new courts can be added.

User testing has been conducted for all 3 of these components and the subsequent pull requests have detailed descriptions of them along with pictures.

Booking Flow:

Now able to update the booking details in real time to the database by sending a JSON to the backend consisting the courtID, courtName, Slot_Index, Sport_name, Sport_ID, Customer_email.
Also made sure to reload the page after booking confirmation to show updated court slots.
Added a toast notification post booking confirmation from BackEnd otherwise an alert with Booking Failed dialog.
Also we made sure that the booking fails if the user who is trying to book a slot is not registered with us.
Performed unit testing on the courts component again in sprint two after adding functionalities.

My-Bookings :

Sends a GET request to /listBookings?email=... to fetch bookings.
We are now displaying the bookings of the user with a good UI and also have the option to CANCEL booking.
For this, a request to the backend with user information will fetch all the booking details affiliated with that user.
Performed Unit Testing on this component to make sure everything is working.

## Work planned but not completed as part of Sprint 3:

- **Implement booking functionality**
  - **Reason**: The integration of the front end to the backend of current APIs and features took up most of the time. However,we have the APIS to handle the booking process ready.

## Unit Test Case Results:

## FROM SPRINT 2

### BACKEND

- GetCourt_test and UpdateCourts_test Unit TestCase Results

![Test Case Results](Pics/Court_TestCases.jpeg)

- CreateCustomer_test Unit TestCase Results

![Test Case Results](Pics/CreateCustomer_TestCase.jpeg)

- CreateBooking_test Unit TestCase Results

![Test Case Results](Pics/CreateBooking_TestCase.jpeg)

### FRONTEND

- Courts Component and Navbar Component Unit TestCase Results

![Test Case Results](Pics/CourtNavbarComponent_TestCases.jpeg)

- Sports Component Unit TestCase Results

![Test Case Results](Pics/SportComponent_TestCase.png)

## Cypress Test FRONTEND

Cypress Test Setup

![Test Case Results](Pics/CypressTest_Setup.jpeg)

Cypress Test Result

![Test Case Results](Pics/CypressTest_Result.jpeg)

## FROM SPRINT 3

- 

## BackEnd API Documentation (Swagger)

## Court Booking API
API for managing court bookings

## Version: 1.0


## /CreateBooking

#### POST
##### Summary:

Create a new booking

##### Description:

Creates a new booking after validating the existence of the customer, sport, and court.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| booking | body | Booking data | Yes | [DataBase.Bookings](#DataBase.Bookings) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Booking record added successfully | object |
| 400 | Invalid request body | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |
| 404 | Customer, sport, or court not found | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |
| 500 | Internal server error | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |

## /Customer

#### POST
##### Summary:

Create a new customer

##### Description:

Adds a new customer to the database if they do not already exist.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| customer | body | Customer data | Yes | [DataBase.Customer](#DataBase.Customer) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Customer already exists |  |
| 201 | Customer record added successfully | object |
| 400 | Invalid request body |  |
| 500 | Internal server error |  |

## /UpdateCourtSlot

#### PUT
##### Summary:

Update court slot status

##### Description:

Toggles the availability of a court time slot. If the slot is booked, it is freed; if it is free, it is booked.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| updateRequest | body | Court slot update request | Yes | [DataBase.CourtUpdate](#DataBase.CourtUpdate) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Slot updated successfully for Court_ID: {Court_ID}, Slot_Index: {Slot_Index} | string |
| 400 | Invalid request body or Slot_Index out of range | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |
| 404 | Court time slots not found | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |
| 500 | Database error or failed to update slot | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |

## /getCourts

#### GET
##### Summary:

Get court availability

##### Description:

Fetches courts based on the selected sport and provides their availability status along with time slots.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| sport | query | Sport name | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | List of available courts with time slots | [ [DataBase.CourtAvailability](#DataBase.CourtAvailability) ] |
| 400 | Missing 'sport' query parameter | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |
| 404 | Sport not found or no courts available | [DataBase.ErrorResponse](#DataBase.ErrorResponse) |

### Models


#### DataBase.Bookings

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Booking_ID | integer |  | No |
| Booking_Status | string |  | No |
| Booking_Time | string |  | No |
| Court_ID | integer |  | No |
| Customer_ID | integer |  | No |
| Sport_ID | integer |  | No |
| court | [DataBase.Court](#DataBase.Court) |  | No |
| customer | [DataBase.Customer](#DataBase.Customer) |  | No |
| sport | [DataBase.Sport](#DataBase.Sport) |  | No |

#### DataBase.Court

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Court_ID | integer |  | No |
| Court_Location | string |  | No |
| Court_Name | string |  | No |
| Court_Status | integer |  | No |
| Sport_id | integer |  | No |
| court_Capacity | integer |  | No |
| sport | [DataBase.Sport](#DataBase.Sport) |  | No |

#### DataBase.CourtAvailability

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| CourtID | integer |  | No |
| CourtName | string |  | No |
| CourtStatus | integer |  | No |
| Slots | [ integer ] |  | No |

#### DataBase.CourtUpdate

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Court_ID | integer |  | No |
| Court_Name | string |  | No |
| Customer_ID | integer |  | No |
| Slot_Index | integer |  | No |
| Sport_name | string |  | No |

#### DataBase.Customer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Contact | string |  | No |
| Customer_ID | integer |  | No |
| email | string |  | No |
| name | string |  | No |

#### DataBase.ErrorResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |

#### DataBase.Sport

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Sport_ID | integer |  | No |
| Sport_name | string |  | No |
| sport_Description | string |  | No |

### Video Recordings
- [Sprint 3 Video Recording]()