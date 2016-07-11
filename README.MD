Maintenance CLI Overview
---

This application is meant to run as a service on the same server as another application / service. It uses
AWS DynamoDB to create, read, and update records which indicate if the maintenance is enabled. The reason for
this application is to allow a decoupled service the ability to put a different service in/out of maintenance.

This allows for real time control of application state without the need for deployments, code changes, or manual
intervention (we don't allow SSH access into our production servers).

This application, when in polling mode, will check for a record in DynamoDB, consume it, and if maintenance is
enabled it will generate a file on the local filesystem which can contain helpful meta data (json). This file could
be used by the other application to determine state. The meta data within the file might include who enabled the
maintenance, ip addresses which can still access the system while in maintenance, etc.

Compiling
---

Make sure your `GOPATH` is set, and a reference to `GOPATH/bin` is in your `PATH`.

- Install dependencies:
	- `make tools`
- Build:
	- `make bin` (all platforms)
	- `make` (only your current platform)

Usage
---

(coming soon)

For now use `maintenance-cli -h`




