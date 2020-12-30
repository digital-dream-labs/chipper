# chipper

Chipper is the voice back-end that drives conversations with [Vector!](https://www.digitaldreamlabs.com/pages/meet-vector)


## Repository layout
|Directory| Description |
|--|--|
| cmd | An example main file using the noop back-end services |
| pkg/server | The main boilerplate for the chipper service |
| pkg/voiceprocessors | voice processing back-ends (and we're looking for contributions!) |
| pkg/vtt | the data types being passed around inside the services, done to avoid GetWhatever() function calls|

## Building

To build the binary, run:
```sh
# make build
```

## Required environment variables

|Variable| Description |
|--|--|
| DDL_RPC_PORT | The RPC port you'd like the service to run on |

You'll also need to either make TLS certificates or set the insecure flag in your environment.  For more information, please see the [hugh grpc library](https://github.com/digital-dream-labs/hugh/tree/main/grpc/server)
