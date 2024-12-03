# Requirements

## Go Powered "Event Booking" REST API

Type | Path | Description | Limits
--- | --- | --- | ---
GET | `/events` | Get a list of available events 
GET | `/events/<id>` | Get a list of available events
POST | `/events` | Create a new bookable event | Auth required
PUT | `/events/<id>` | Update an event | Auth required, creator only
DELETE | `/events/<id>` | Delete an event | Auth required, creator only
POST | `/signup` | Create a new user
POST | `/login` | Authenticate user
POST | `/events/<id>/register` | Register user for event | Auth required
DELETE | `/events/<id>/register` | Cancel Registration | Auth required



Authentication will be handled via JWT's (JSON Web Tokens)