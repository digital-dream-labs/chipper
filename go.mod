module github.com/digital-dream-labs/chipper

go 1.15

require (
	github.com/digital-dream-labs/api v0.0.0-20201230194801-ee4f8f160098
	github.com/digital-dream-labs/hugh v0.0.0-20210210154335-f4159b9fcd5f
	google.golang.org/grpc v1.33.1
	gopkg.in/yaml.v3 v3.0.0-20200605160147-a5ece683394c // indirect
)

replace github.com/digital-dream-labs/api => ../api
