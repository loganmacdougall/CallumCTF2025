package main

import (
	"city_game/src/game"
	"city_game/src/pb"
	"testing"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

func TestDoNothing(t *testing.T) {
	state := game.CreateInitialState()

	start_state := proto.CloneOf(state)

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_N_Nothing},
		HelperInput:       []*pb.HelperInput{},
	}

	for _ = range 10 {
		next_state := game.GetNextState(state, input)
		if len(next_state.ErrorMessages) != 0 {
			t.Log(next_state.ErrorMessages[0].Message)
			t.FailNow()
		}

		state = next_state
	}

	if state.Tick != 10 {
		t.Logf("expected tick 10, but was actually tick %d", state.Tick)
		t.FailNow()
	}
	state.Tick = 0 // The tick count should be the only change

	if !proto.Equal(start_state, state) {
		t.Log("start state must match end state")
		t.Logf("Start: %s\nEnd: %s", prototext.Format(start_state), prototext.Format(state))
		t.FailNow()
	}
}
