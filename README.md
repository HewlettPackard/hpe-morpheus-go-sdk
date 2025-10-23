# hpe-morpheus-go-sdk

This repository contains the Go SDKs used to drive Morpheus in the HPE Morpheus Terraform Provider.

## Packages

### oapigen
The `oapigen` package contains the SDK generated from the [Morpheus OpenAPI Specification](https://github.com/HewlettPackard/morpheus-openapi).
The SDK is generated using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator), an open source SDK and server code generator which supports a multitude of different languages.

### WARNING: FOR INTERNAL USE WITH MORPHEUS TERRAFORM PROVIDER ONLY

Do NOT use as general purpose clients for interacting with Morpheus.

These clients are intended to be used in the [HPE Morpheus Terraform Provider](https://github.com/HPE/terraform-provider-hpe).

Not all endpoints have been tested as development of the Terraform provider is ongoing.
