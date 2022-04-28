package convert

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func CastToAny(input interface{}) (*anypb.Any, error) {
	b, _ := json.Marshal(input)
	s := &structpb.Struct{}
	err := protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	data, err := anypb.New(s)
	if err != nil {
		return nil, err
	}
	return data, nil
}
