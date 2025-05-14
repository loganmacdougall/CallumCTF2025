package main

import (
	"city_game/src/game"
	"city_game/src/pb"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestHelperWalking1(t *testing.T) {
	state := game.CreateInitialState()
	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{
			ActionType: pb.NecromancerActionType_Summon,
			Coordinate: &pb.Coordinate{X: 18, Y: 6},
		},
	}

	state = game.GetNextState(state, input)
	helper_index := 1
	helper_id := uint32(1)

	input = &pb.GameInput{
		HelperInput: []*pb.HelperInput{
			{
				HelperId: helper_id,
				Action: &pb.Action{
					ActionType: pb.ActionType_Walk,
					Coordinate: &pb.Coordinate{X: 21, Y: 6},
				},
			},
		},
	}

	state = game.GetNextState(state, input)
	input = &pb.GameInput{}

	for range 2 {
		state = game.GetNextState(state, input)
		if len(state.ErrorMessages) > 0 {
			t.Logf("error: %s", state.ErrorMessages[0])
			t.FailNow()
		}
	}

	expected_pos := &pb.Coordinate{X: 20, Y: 7}
	actual_pos := state.HelperStates[helper_index].Coordinate
	if !proto.Equal(expected_pos, actual_pos) {
		t.Logf("expected to be at coordinate (%d, %d) but was actually at (%d, %d)",
			expected_pos.X, expected_pos.Y, actual_pos.X, actual_pos.Y)
		t.FailNow()
	}

	for range 2 {
		state = game.GetNextState(state, input)
		if len(state.ErrorMessages) > 0 {
			t.Logf("error: %s", state.ErrorMessages[0])
			t.FailNow()
		}
	}

	expected_pos = &pb.Coordinate{X: 21, Y: 6}
	actual_pos = state.HelperStates[helper_index].Coordinate
	if !proto.Equal(expected_pos, actual_pos) {
		t.Logf("expected to be at coordinate (%d, %d) but was actually at (%d, %d)",
			expected_pos.X, expected_pos.Y, actual_pos.X, actual_pos.Y)
		t.FailNow()
	}

	if state.HelperStates[helper_index].Action.ActionType != pb.ActionType_A_Nothing {
		t.Log("expected action to be completed and set to nothing but wasn't")
		t.FailNow()
	}
}

func TestHelperWalking2(t *testing.T) {
	state := game.CreateInitialState()
	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{
			ActionType: pb.NecromancerActionType_Summon,
			Coordinate: &pb.Coordinate{X: 18, Y: 6},
		},
	}

	state = game.GetNextState(state, input)
	helper_index := 1
	helper_id := uint32(1)
	direction := pb.Direction_Down

	input = &pb.GameInput{
		HelperInput: []*pb.HelperInput{
			{
				HelperId: helper_id,
				Action: &pb.Action{
					ActionType: pb.ActionType_Walk,
					Direction:  &direction,
				},
			},
		},
	}

	for range 4 {
		state = game.GetNextState(state, input)
		if len(state.ErrorMessages) > 0 {
			t.Logf("error: %s", state.ErrorMessages[0])
			t.FailNow()
		}
	}

	if state.HelperStates[helper_index].Coordinate.Y != 10 {
		t.Logf("expected to be at y: %d but was actually at y: %d",
			state.HelperStates[helper_index].Coordinate.Y, 10)
		t.FailNow()
	}

	for range 5 {
		state = game.GetNextState(state, input)
		if len(state.ErrorMessages) > 0 {
			t.Logf("error: %s", state.ErrorMessages[0])
			t.FailNow()
		}
	}

	if state.HelperStates[helper_index].Coordinate.Y != 15 {
		t.Logf("expected to be at y: %d but was actually at y: %d",
			state.HelperStates[helper_index].Coordinate.Y, 15)
		t.FailNow()
	}

	if state.HelperStates[helper_index].Action.ActionType != pb.ActionType_A_Nothing {
		t.Log("expected action to be completed and set to nothing but wasn't")
		t.FailNow()
	}
}

func TestHelperWalking3(t *testing.T) {
	state := game.CreateInitialState()
	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{
			ActionType: pb.NecromancerActionType_Summon,
			Coordinate: &pb.Coordinate{X: 18, Y: 6},
		},
	}

	state = game.GetNextState(state, input)
	helper_index := 1
	helper_id := uint32(1)

	input = &pb.GameInput{
		HelperInput: []*pb.HelperInput{
			{
				HelperId: helper_id,
				Action: &pb.Action{
					ActionType: pb.ActionType_Walk,
				},
			},
		},
	}

	state = game.GetNextState(state, input)

	if len(state.ErrorMessages) != 0 {
		t.Log("expected to receive error message but didn't")
	}

	if state.HelperStates[helper_index].Action.ActionType != pb.ActionType_A_Nothing {
		t.Log("expected action to be set to nothing but wasn't")
		t.FailNow()
	}
}
