# chipper

Chipper is the voice back-end that drives conversations with [Vector!](https://www.digitaldreamlabs.com/pages/meet-vector)

Note: This is a reference implementation that always responds with the intent defined [here](https://github.com/digital-dream-labs/chipper/blob/e0f2402b01c6106be889175b4e7d5e16bf3bc4e6/pkg/voice_processors/noop/intent.go#L13).  It does not include speech to intent capabilities.  

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

## Using with your vector

1.  Create a [self signed TLS certificate](https://www.linode.com/docs/guides/create-a-self-signed-tls-certificate/)

2.  Append the root certificate to the vector-clouds built-in root certificate store

	You'll need to modify [this](https://github.com/digital-dream-labs/vector-cloud/blob/622564bda6473b755cf99687333a20e396c96308/internal/jdocs/client.go#L30) and [this](https://github.com/digital-dream-labs/vector-cloud/blob/622564bda6473b755cf99687333a20e396c96308/internal/jdocs/escapepod_root_cert.go#L3).  Once completed, compile and upload the binaries to your robot.  [This package](https://github.com/digital-dream-labs/vector-configurator) may help you out with that.

3.  Modify your bots server_config.json

	You'll have to edit /anki/data/assets/cozmo_resources/config/server_config.json and make it look something like this:

	```json
	{
		"jdocs": "jdocs.api.anki.com:443",
		"tms": "token.api.anki.com:443",
		"chipper": "$YOUR_CHIPPER_INSTANCE_IP:$YOUR_CHIPPER_INSTANCE_PORT",
		"check": "conncheck.global.anki-services.com/ok",
		"logfiles": "s3://anki-device-logs-prod/victor",
		"appkey": "oDoa0quieSeir6goowai7f"
	}
	```

	Once this is done, reboot your bot.

4.  Create an [environment file](https://docs.docker.com/compose/environment-variables/#the-env-file) for the docker instance

	Your contents should resemble this:
	```yaml
	DDL_RPC_PORT=YOUR_SERVICE_PORT
	DDL_RPC_TLS_CERTIFICATE=very long tls certificate
	DDL_RPC_TLS_KEY=very long tls key
	DDL_RPC_CLIENT_AUTHENTICATION=None
	```

    Note:  If you use a port other than 8084, you'll want to modify the "expose" and "ports" entries in the [docker-compose](https://github.com/digital-dream-labs/chipper/blob/main/docker-compose.yaml) file if you plan on using docker.

5.  Building and running

	```sh
	$ make build-linux
	$ docker-compose up
	```

## Writing your own implementation

The [noop implementation](https://github.com/digital-dream-labs/chipper/tree/main/pkg/voice_processors/noop) provides a bare implementation, but doesn't actually read data from the incoming streams.  

Here's a quick pseudo-code example for the IntentRequest:

```go
func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {
	
    /*
     This is copying all of the data to a byte slice, which you most likely wouldn't want
    */
    data := []byte{}
	data = append(data, req.FirstReq.InputAudio...)
	for {
		chunk, err := req.Stream.Recv()
		if err != nil {
			if err == io.EOF {
				// The incoming file is done.  Break out of the loop.
				break
			}
			/*
				Return whatever error intent you'd like here, but most likely,
				some sort of transmission error has happened and the client isn't
				listening for a response
			*/

		}
		data = append(data, chunk.InputAudio...)
	}
	
	/*
		Do something with the audio data to turn the data stream into an intent.
		this can be a static response, piped through a speech to text translation implementation,
		or anything else you want.  For now, we'll just print the data to the screen just to do
		something with it.  Yes, we're printing an ogg stream :-D
	*/
	fmt.Println(data)

	intent := pb.IntentResponse{
		IsFinal: true,
		IntentResult: &pb.IntentResult{
			// This is what would be defined by the previous speech to text processing.
			Action: "intent_play_cantdo",
		},
	}

	if err := req.Stream.Send(&intent); err != nil {
		return nil, err
	}

	/*
		This is returned for logging, metrics gathering, or whatever else you'd want.  It has no
		impact on the communication with the bot itself.
	*/
	r := &vtt.IntentResponse{
		Intent: &intent,
	}
	return r, nil
}
```