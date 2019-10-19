# proto-gen-swiftapiclient

proto-gen-swiftapiclient is a command line tool to create Swift API client from `.proto`.

## Install

```console
go get -u github.com/sonatard/proto-gen-swiftapiclient
```

## Usage

```console
proto-gen-swiftapiclient -i proto/ api/v1/*.proto
```

## Not supported

-	Streaming API
	-	Not create a convert method.
-	[HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule)
-	`map` type query string

