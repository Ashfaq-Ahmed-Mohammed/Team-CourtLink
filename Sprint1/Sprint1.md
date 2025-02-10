# Team-Courtlink (Project Group 4) Sprint 1

### GitHub Repository Link: [CourtLink](https://github.com/Ashfaq-Ahmed-Mohammed/Team-CourtLink)

## User Stories

- **As a developer**, I want to learn and implement Angular 17 best practices, so that I can build a good application.
- **As a developer**, I want to build modular Angular components for the navbar, sports tiles, and court selection, so that the application is maintainable and scalable.
- **As a developer**, I want to use Angular Material and Tailwind CSS for styling, so that the UI looks modern and responsive.
- **As a developer**, I want to minimize unnecessary re-renders, so that the UI remains smooth and efficient.
- **As a developer**, I want to ensure smooth navigation in the single-page application, so that users do not experience page reloads while switching between sections.
- **As a user**, I want a fast and secure way to log into the application.
- **As a user**, I want a responsive navigation bar, so that I can easily navigate between different sections of the application.
- **As a user**, I want to see a list of available sports, so that I can quickly choose the sport I want to book.
- **As a user**, I want each sport in the grid to have a logo, so that I can easily identify sports visually.
- **As a user**, I want to view available courts for each sport, so that I can select a court based on availability.
- **As a developer**, I want to design and implement a database schema for storing user personal information to store customer details for future interactions and bookings securely.
- **As a developer**, I want to create a database schema for court bookings so that reservations are accurately recorded and managed.
- **As a developer**, I want to design a database schema for court information so that details about each court are systematically stored and retrievable.
- **As a developer**, I want to modify the database schema to support multiple courts for each sport so that users can book specific courts under a particular sport category.
- **As a developer**, I want to implement a user registration API so that new users can create accounts securely.
- **As a developer**, I want to create an endpoint that fetches available court slots so that the front end can display them to users.
- **As a developer**, I want to add a time slot column to the bookings table so that each booking can store specific time slot information.
- **As a developer**, I want to implement a customer registration system that prevents double registrations so that users cannot register on multiple logins.
- **As a developer**, I want to add a location filter to the court information API so that users can search for available courts based on their preferred location.
- **As a developer**, I want to create and maintain comprehensive documentation for the backend APIs so that front-end developers and other team members can easily understand and integrate with the system.

## Issues

**Issues that were planned for Sprint 1:**

- Design and implement a database to store, retrieve, and manage information related to customers, sports, bookings, and courts.
- REST API - HTTP POST to authenticate users using Google or Apple IDs, and also store information like name, email, and contact in the application database for setting up their personalized profiles.
- REST API - HTTP GET to fetch information on the availability of courts and specific time slots based on the sport selected by the user at a specific location.
- Documenting backend APIs using Swagger.
- Conceptualize the layout for the application, keeping ease of use and speed in mind.
- Create a navigation bar component and provide basic styling to the webpage.
- Create a sports component that displays a list of all available sports that users may choose from, making it visually appealing while maintaining clarity. Provide pictures associated with each sport to make the process faster.
- Integrate authentication via Auth0 through the navbar component to provide fast and secure log-in functionality to users.
- Work with the back-end team to develop APIs for fetching the stored user information from Auth0’s databases to our own.
- Start work on routing and develop a component that can receive JSON tokens containing sport details for booking and render them on one premade template, thus reducing the number of individual components required.
- Continue work on the aforementioned component and conceptualize the layout of the sport selection screen.
- Display the court and availability information provided by the backend team in this component and complete the booking flow.

**Issues that were planned but are not part of Sprint 1:**

- Documenting backend APIs using Swagger.
  - **Why:** The API code is constantly changing due to the addition of JWT authentication, validation, and additional logic for GET and POST API methods. Hence, Swagger documentation was not taken up this sprint. It will be part of the next sprint and will be added to the upcoming APIs as well.
- Display the court and availability information provided by the backend team in this component and complete the booking flow.
  - **Why:** We were unable to focus on developing the time slot booking system as planned, as our efforts were directed toward creating a single-page application. This required additional time to properly configure the routing of the components.


### Video Recordings
- [FrontEnd Recording](https://drive.google.com/file/d/1ljUqdiY-nMspxK2W6NZ-VHV2VZV7shOx/view?usp=sharing)
- [BackEnd Recording](https://drive.google.com/file/d/1iWfi_NnLlPfjPQI-3P1WekB22cshHVUZ/view?usp=drive_link)