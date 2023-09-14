# Intro to REST

Server & Client

[Postman](https://www.postman.com/)

REST - Representational State Transfer

REST is an architectural style

"RESTFul" - REST compliant

Key notes:
- Stateless
- Helps define how server & client interact
Stateless - we don't need to remember what clients were doing. A request has
enough info for us to respond.

Client & Server interactions - REST principles can help us design a server that
is more intuitive for clients to use.


Requests consist of:
- HTTP Method (sometimes called HTTP verb)
- A Path

eg `GET /galleries`

Other info is okay as part of request - request body. headers. etc

Endpoint => HTTP Method + Path

| **HTTP Method** | **Path**       | **What happens**         |
|-----------------|----------------|--------------------------|
| `GET`           | /galleries     | Read a list of galleries |
| `GET`           | /galleries/:id | Read a single gallery    |
| `POST`          | /galleries     | Create a gallery         |
| `PUT`           | /galleries/:id | Update a gallery         |
| `DELETE`        | /galleries/:id | Delete a gallery         |


Even our pages that just render forms for editing or creating resources can 
follow this pattern:

| **HTTP Method** | **Path**            | **What happens**                   |
|-----------------|---------------------|------------------------------------|
| `GET`           | /galleries/new      | Read a form for creating galleries |
| `GET`           | /galleries/:id/edit | Read a form for editing a gallery  |


Publishing a Gallery - 3 options

1 - update

| **HTTP Method** | **Path**       | **What happens** |
|-----------------|----------------|------------------|
| `PUT`           | /galleries/:id | Update a gallery |


2 - publication resource

| **HTTP Method** | **Path**      | **What happens**                   |
|-----------------|---------------|------------------------------------|
| `POST`          | /publications | Create a publication for a gallery |

3 - semi-custom endpoint

| **HTTP Method** | **Path**               | **What happens**    |
|-----------------|------------------------|---------------------|
| `POST`          | /galleries/:id/publish | Publish a gallery   |
| `DELETE`        | /galleries/:id/publish | Unpublish a gallery |

Orgnaizing our code w/ REST:

- We will likely have a controller for each resource, with methods on that
controller for endpoints related to it.
- Views will likely map to REST endpoints

```bash
templates/
    galleries/
        show.gohtml # GET /galleries/:id
        list.gohtml # GET /galleries
        new.gohtml  # GET /galleries/new
        edit.gohtml # GET /galleries/:id/edit
```

There are exceptions; use REST as a guideline.

*Want additional info? See <https://www.codecademy.com/articles/what-is-rest>
or Google for some articles. Be wary of anything dictating a set of definitive
rules - no golden standard for that is/isn't REST.*
