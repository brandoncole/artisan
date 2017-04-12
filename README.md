# Artisan

Have you ever...

- Had to build a service that stores and retrieves data?
- Wondered which datastore will work best? (PostgreSQL, MongoDB, S3, Cassandra, Redis, ElasticSearch, etc.)
- Wanted to switch between datastores if operational issues arose with the current one?
- Wished you could just evaluate several data stores and pick the best option?
- Wondered how to support all the ways to make it available? (REST, WebSockets, CLI, SDKs, Ajax, etc.)
- Worried how to make a secure system that defends against OWASP attacks (XSS, CSRF, Nonces, etc.)
- Had to make sure the system supported multi-tenanacy? (ACLs, RBAC, ABAC, Basic Auth, etc.)
- Needed to ensure data compatibility? (Migration, Versioning, Forward and Backwards Compatibility, etc.)
- Wanted to have all the modern goodies? (Read Replicas, Caching, Rate Limiting, Webhooks, Queuing, etc.)
- Wished you had a standard pipeline for responding to Creation, Deletion and Modifications of data?
- Wondered the best way to balance the trade-offs of the CAP Theorem for the service?
- Heard about the 12-factor app and want to embrace those principles?
- Needed to support zero-downtime deployments for future updates?
- Envied your Ruby on Rails peers who got some of this with Ruby?

... and ...

- Don't care about implementing ANY of that from scratch!
- Know exactly what your data looks like and would love to get all that!

Then you've come to the right spot!  The goal of Artisan is to free developers from continuously re-building all this scaffolding each time we need to deliver another service and let them simply define the data and let the system take care of the rest.

### Targeted Developer Workflow

What's important to this project and remains at the forefront of the development efforts are that:

- Emphasis on limiting the amount of code developers must provide as they define the model
- Most of the code for the model that developers provide are simply Protocol Buffers
- Extensions to Protocol Buffers are provided by Artisan to help developers easily describe additional constraints
- Switching between different front-ends and back-ends should be trivial
- Artisan provides a container to generate all this code goodness from the model developers provide
- When developers do need to provide code, future code generation from updated models shouldn't overwrite their efforts

### How It Works

This is very much a work in progress and is proceeding by getting the [Northwind](examples/northwind) example to work via test-driven development.

- https://northwinddatabase.codeplex.com

Northwind is a well known, nostalgic and fictitious example of a set of data that captures several real-world realities.

1. Developer starts by providing two things - Configuration and Model Definition
    - `Configuration` - A YAML file that describes where the model definition is, where the generated code should be placed and eventually what type of front-ends and back-ends should be generated.
        - See: [examples/northwind/artisan.yaml](examples/northwind/artisan.yaml)
    - `Model Definition` - A set of Protocol Buffer files that define the model that the service is providing.  The model definition can reference the Artisan Protocol Buffer extensions that make it easy to achieve complex things
        - See: [examples/northwind/protos](examples/northwind/protos)
        - See: [artisan/api](artisan/api)
2. Developer generates stubs using Artisan and the model.  This is achieved today by running a Docker container that this project builds which mounts the inputs above and generates code.  Some of the specifics about this step are:
    - `Makefile` - Launches the Artisan container for the example
        - See: [examples/northwind/Makefile](examples/northwind/Makefile)
    - `Descriptor` - The first thing that the Artisan container generates is a binary file that encompasses all the specifics defined in the protocol buffers.  This binary file is only generated when the syntax is of the *.proto files are 100% correct.  Once this file is generated there is no need to refer back to the *.proto files
        - See: [examples/northwind/artisan/schema](examples/northwind/artisan/schema)
    - `gRPC Bindings` - The definition of the `gRPC` service is generated at the same time as the Descriptor.  This is a bunch of stub code that has no concrete implementation yet for the back-end
        - See: [examples/northwind/artisan/grpc](examples/northwind/artisan/grpc)
3. Artisan Generates Front-Ends and Back-ends
    - The CLI is in it's infancy right now but it currently demonstrates the basic reflective capabilities of interrogating the `Descriptor` to enumerate all the Services, Methods, Types and Fields that are available.  That capability can ultimately be used to generate concrete implementations of back-ends and front-ends for the `gRPC` service
        - See: [artisan/cmd/generate.go](artisan/cmd/generate.go)
    - Artisan does plan to use `Go` as the implementation language and will more than likely use the new `Plugin` capability in Go to provide extensible plugins for different back-ends and front-ends
        - See: [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

# Related Reading and Inspiration

The problems, technologies or methodologies discussed or implemented here are not necessarily novel.  Some of the inspiration for the content here can be attributed to the following:

- [Google gRPC](http://www.grpc.io/)
- [Google Protocol Buffers](https://developers.google.com/protocol-buffers/)
- [Apache Avro](https://avro.apache.org/docs/1.8.1/spec.html)
- [Microsoft COM+](https://msdn.microsoft.com/en-us/library/windows/desktop/ms685978(v=vs.85).aspx)
- [Microsoft Type Libraries, IDLs](https://msdn.microsoft.com/en-us/library/windows/desktop/aa366757(v=vs.85).aspx)
- [OpenAPI](https://www.openapis.org/)
- [Swagger](http://swagger.io/)
- [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
- [12 Factor App](https://12factor.net/)
- [OWASP](https://www.owasp.org)
- [Jepsen.io](http://jepsen.io/)
- [Ruby on Rails](http://rubyonrails.org/)
- [Django](https://www.djangoproject.com/)

# Building

Building and using this project assumes you have the following prerequisites, and so far this has only been tested on Mac:

| Component | Version |
| --------- | ------- |
| `make` | 3.81+ |
| `go`  | 1.8+ |
| `docker` | 1.13+ |

From the shell the following commands should work:

```shell
# Build Everything....
make
```

# Disclaimer

This project is in its infancy and does not currently provide all the capabilities listed here.  Over time it likely will and help from the community would be a fantastic way to move this along quickly.
