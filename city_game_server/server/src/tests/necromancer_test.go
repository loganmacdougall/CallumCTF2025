package main

import (
	"city_game/src/game"
	"city_game/src/pb"
	"testing"
)

func TestSummon(t *testing.T) {
	state := game.CreateInitialState()

	coordinates := [3]*pb.Coordinate{{X: 0, Y: 0}, {X: 21, Y: 10}, {X: 20, Y: 5}}

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_Summon},
		HelperInput:       []*pb.HelperInput{},
	}

	for i, coordinate := range coordinates {
		input.NecromancerAction.Coordinate = coordinate
		next_state := game.GetNextState(state, input)
		if len(next_state.ErrorMessages) != 0 && i != 2 {
			t.Logf("i: %d (%d, %d)", i, coordinate.X, coordinate.Y)
			t.FailNow()
		} else if len(next_state.ErrorMessages) == 0 && i == 2 {
			t.Log("Expected error when summoned on building")
			t.Logf("i: %d (%d, %d)", i, coordinate.X, coordinate.Y)
			t.FailNow()
		}
		state = next_state
	}

	if len(state.HelperStates) != 3 ||
		state.HelperStates[0].HelperId != 0 ||
		state.HelperStates[1].HelperId != 1 ||
		state.HelperStates[2].HelperId != 2 {
		t.Log("The helper ids don't match the expected")
		t.FailNow()
	}
}

func TestRelease(t *testing.T) {
	state := game.CreateInitialState()

	coordinates := [2]*pb.Coordinate{{X: 0, Y: 0}, {X: 21, Y: 10}}

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_Summon},
		HelperInput:       []*pb.HelperInput{},
	}

	for _, coordinate := range coordinates {
		input.NecromancerAction.Coordinate = coordinate
		state = game.GetNextState(state, input)
	}

	if len(state.HelperStates) != 3 ||
		state.HelperStates[0].HelperId != 0 ||
		state.HelperStates[1].HelperId != 1 ||
		state.HelperStates[2].HelperId != 2 {
		t.Log("The helper ids don't match the expected")
		t.FailNow()
	}

	release_id := uint32(1)
	input.NecromancerAction.ActionType = pb.NecromancerActionType_Release
	input.NecromancerAction.HelperId = &release_id
	state = game.GetNextState(state, input)

	if len(state.HelperStates) != 2 ||
		state.HelperStates[0].HelperId != 0 ||
		state.HelperStates[1].HelperId != 2 {
		t.Log("The helper ids don't match the expected")
		t.FailNow()
	}

	release_id = 0
	state = game.GetNextState(state, input)

	if len(state.HelperStates) != 1 ||
		state.HelperStates[0].HelperId != 2 {
		t.Log("The helper ids don't match the expected")
		t.FailNow()
	}

	release_id = 2
	state = game.GetNextState(state, input)

	if len(state.HelperStates) != 0 {
		t.Log("The helper ids don't match the expected")
		t.FailNow()
	}
}
