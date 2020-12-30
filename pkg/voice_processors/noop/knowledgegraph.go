package noop

import (
	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/digital-dream-labs/chipper/pkg/vtt"
)

const (
	// NoResult defines a not-found result
	NoResult = "NoResultCommand"
	// NoResultSpoken defines the text to return with NoResult
	NoResultSpoken = "Didn't Get That"
)

// ProcessKnowledgeGraph handles knowledge graph interactions
func (s *Server) ProcessKnowledgeGraph(req *vtt.KnowledgeGraphRequest) (*vtt.KnowledgeGraphResponse, error) {
	kg := pb.KnowledgeGraphResponse{
		Session:     req.Session,
		DeviceId:    req.Device,
		CommandType: NoResult,
		SpokenText:  NoResultSpoken,
	}

	if err := req.Stream.Send(&kg); err != nil {
		return nil, err
	}

	return &vtt.KnowledgeGraphResponse{
		Intent: &kg,
	}, nil

}
