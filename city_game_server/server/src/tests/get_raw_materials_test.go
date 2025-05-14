package main

import (
	"city_game/src/game"
	"city_game/src/pb"
	"testing"
)

func TestGetPlanks(t *testing.T) {
	state := game.CreateInitialState()

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_N_Nothing},
		HelperInput: []*pb.HelperInput{
			{
				HelperId: 0,
				Action: &pb.Action{
					ActionType: pb.ActionType_Walk,
					Coordinate: &pb.Coordinate{X: 31, Y: 16},
				},
			},
		},
	}

	for range 17 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 31, Y: 16},
			},
		},
	}

	state = game.GetNextState(state, input)

	if len(state.HelperStates[0].Items) == 0 {
		t.Log("Failed to collect planks, as inventory is empty")
		t.FailNow()
	}

	if state.HelperStates[0].Items[0].ItemId != pb.Item_Plank {
		t.Log("Failed to collect planks, as wrong item appeared in inventory")
		t.FailNow()
	}

	right := pb.Direction_Right
	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Direction:  &right,
			},
		},
	}

	state = game.GetNextState(state, input)

	if state.HelperStates[0].Items[0].Count != 2 {
		t.Log("Expected to collect two planks but didn't")
		t.FailNow()
	}

	for range 14 {
		state = game.GetNextState(state, input)
	}

	if len(state.HelperStates[0].Items) != 1 {
		t.Log("Expected 16 planks to only require one item slot")
		t.FailNow()
	}

	state = game.GetNextState(state, input)

	if len(state.HelperStates[0].Items) != 2 {
		t.Log("Expected the 17th plank to create a second slot of planks")
		t.FailNow()
	}

	if state.HelperStates[0].Items[0].Count != 16 ||
		state.HelperStates[0].Items[1].Count != 1 {
		t.Log("The counts in the two stacks are wrong")
		t.FailNow()
	}
}

func TestGetStone(t *testing.T) {
	state := game.CreateInitialState()

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_N_Nothing},
		HelperInput: []*pb.HelperInput{
			{
				HelperId: 0,
				Action: &pb.Action{
					ActionType: pb.ActionType_Walk,
					Coordinate: &pb.Coordinate{X: 20, Y: 16},
				},
			},
		},
	}

	for range 6 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 20, Y: 16},
			},
		},
	}

	state = game.GetNextState(state, input)

	if len(state.HelperStates[0].Items) == 0 {
		t.Log("Failed to collect stones, as inventory is empty")
		t.FailNow()
	}

	if state.HelperStates[0].Items[0].ItemId != pb.Item_Stone {
		t.Log("Failed to collect stones, as wrong item appeared in inventory")
		t.FailNow()
	}

	down := pb.Direction_Down
	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Direction:  &down,
			},
		},
	}

	state = game.GetNextState(state, input)

	if state.HelperStates[0].Items[0].Count != 2 {
		t.Log("Expected to collect two stones but didn't")
		t.FailNow()
	}

	for range 14 {
		state = game.GetNextState(state, input)
	}

	if len(state.HelperStates[0].Items) != 1 {
		t.Log("Expected 16 stones to only require one item slot")
		t.FailNow()
	}

	state = game.GetNextState(state, input)

	if len(state.HelperStates[0].Items) != 2 {
		t.Log("Expected the 17th stone to create a second slot of planks")
		t.FailNow()
	}

	if state.HelperStates[0].Items[0].Count != 16 ||
		state.HelperStates[0].Items[1].Count != 1 {
		t.Log("The counts in the two stacks are wrong")
		t.FailNow()
	}
}
