# RESTful APIs

Representational State Transfer, or REST, is a popular convention that many dynamic HTTP servers follow. Not all HTTP APIs are "REST APIs", or "RESTful", but it is very common.

RESTful servers follow a loose set of rules that makes it easy to build reliable and predictable web APIs. REST is a set of conventions about how HTTP APIs should be built.

## Separate and agnostic

The big idea behind REST is that resources are transferred via well-recognized, language-agnostic client/server interactions. A RESTful style means the implementation of the client and server can be created independently of one another, as long as some simple standards, like the names of the available resources, have been established.

## Stateless

A RESTful architecture is stateless, which means the server does not need to know what state the client is in, nor does the client need to know what state the server is in. Statelessness in REST is enforced by interacting with resources instead of commands. Keep in mind, this doesn't mean the applications are stateless - what would "updating a resource" even mean if the server wasn't keeping track of its state?

## Paths in REST

In a RESTful API, the last section of the path of a URL specifies the resource. For Jello, this means an issue, a user, or a project. Depending on whether the request is a GET, POST, PUT or DELETE, the resource is read, created, updated, or deleted.

The Jello API we have been working with is a RESTful API! Take a look at the URLs we've been using:

- <https://api.boot.dev/v1/courses_rest_api/learn-http/projects>
- <https://api.boot.dev/v1/courses_rest_api/learn-http/users>
- <https://api.boot.dev/v1/courses_rest_api/learn-http/issues>

1. The first part of the path specifies the version. In this case, version 1, or v1.
2. The second part of the path tells our server that this is a REST API for the "Learn HTTP" course.
3. The last part denotes which resource is being accessed, be it a project, user, or issue.
