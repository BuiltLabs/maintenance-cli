## Maintenance CLI Overview

This application is meant to run as a service on the same server as another application / service. It uses
AWS DynamoDB to create, read, and update records which indicate if the maintenance is enabled. The reason for
this application is to allow a decoupled service the ability to put a different service in/out of maintenance.

This allows for real time control of application state without the need for deployments, code changes, or manual
intervention (we don't allow SSH access into our production servers).

This application, when in polling mode, will check for a record in DynamoDB, consume it, and if maintenance is
enabled it will generate a file on the local filesystem which can contain helpful meta data (json). This file could
be used by the other application to determine state. The meta data within the file might include who enabled the
maintenance, ip addresses which can still access the system while in maintenance, etc.

## Compiling

Make sure your `GOPATH` is set, and a reference to `GOPATH/bin` is in your `PATH`.

- Install dependencies:
	- `make tools`
- Build:
	- `make bin` (all platforms)
	- `make` (only your current platform)

## Basic Usage

#### Environment

**IMPORTANT** This application requires AWS credentials either in the environment or in `~/.aws/credentials` For more details visit
[AWS Credentials](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs).

## Help

Commands and basic usage are covered using the `--help, -h` option.

```
$ ./maintenance-cli -h

NAME:
   maintenance-cli - Used to control maintenance state and poll

USAGE:
   maintenance-cli [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     poll     polls remote state and creates/deletes maintenance file
     enable   flags maintenance mode as enabled in dynamodb
     disable  flags maintenance mode as disabled in dynamodb
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```
$ ./maintenance-cli poll -h

NAME:
   maintenance-cli poll - polls remote state and creates/deletes maintenance file

USAGE:
   maintenance-cli poll [command options] [arguments...]

OPTIONS:
   --file FILE, -f FILE       Location to create/delete FILE for maintenance events (default: "/tmp/maintenance.enable")
   --table value, -t value    Table where maintenance state is stored (default: "maintenanceFlags")
   --key value, -k value      Primary lookup key to query maintenance data (default: "ops")
   --keyName value, -n value  Primary lookup key column name (default: "environment")
```

## Examples
> Synopsis: Enable maintenance mode which can be consumed by `maintenance-cli poll` on a different machine. NOTE: if the IAM key provided has permission, it will create the table if it doesn't exist
```
$ AWS_ACCESS_KEY_ID=... ./maintenance-cli enable --table maintenance-table --key development
[2016-07-11 13:45:45.59624503 -0500 CDT] action:enable ** creating table [maintenance-table]
[2016-07-11 13:45:45.59624503 -0500 CDT] action:enable ** waiting for table creation [maintenance-table]
[2016-07-11 13:45:45.59624503 -0500 CDT] action:enable ** maintenance enabled
```
---
> Synopsis: Using the maintenance-cli tool to poll for maintenance mode (using the paremeters noted above), writing a file named /application/maintenance.json if enabled. NOTE: This program will poll every 5 seconds until terminated.
```
$ AWS_ACCESS_KEY_ID=... ./maintenance-cli poll --file /app/maintenance.json --table maintenance-table --key development
[2016-07-11 14:11:22.118842334 -0500 CDT] action:poll ** maintenance enabled
[2016-07-11 14:11:22.118842334 -0500 CDT] action:poll ** maintenance enabled
...

$ ls -al /app/maintenance.json
-rw-r--r--  1 foo  wheel  0 Jul 11 15:16 /app/maintenance.json
```