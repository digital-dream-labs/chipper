package noop

import (
	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/digital-dream-labs/chipper/pkg/vtt"
)

func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {

	intent := pb.IntentResponse{
		IsFinal: true,
		IntentResult: &pb.IntentResult{
			Action: "intent_play_cantdo",
		},
	}

	if err := req.Stream.Send(&intent); err != nil {
		return nil, err
	}

	r := &vtt.IntentResponse{
		Intent: &intent,
	}
	return r, nil
}
