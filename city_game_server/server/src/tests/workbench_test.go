package main

import (
	"city_game/src/game"
	"city_game/src/pb"
	"testing"
)

func TestCrafting1(t *testing.T) {
	state := game.CreateInitialState()

	workbench_index := -1
	for i, state := range state.BuildingStates {
		if state.BuildingType == pb.Building_Workbench {
			workbench_index = i
			break
		}
	}

	right := pb.Direction_Right
	down := pb.Direction_Down

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{
			ActionType: pb.NecromancerActionType_Summon,
			Coordinate: &pb.Coordinate{X: 30, Y: 16},
		},
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

	state = game.GetNextState(state, input)

	input = &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_N_Nothing},
		HelperInput: []*pb.HelperInput{
			{
				HelperId: 1,
				Action: &pb.Action{
					ActionType: pb.ActionType_Interact,
					Direction:  &right,
				},
			},
		},
	}

	for range 4 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Walk,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
			},
		},
	}

	// This will be the point where helper 0 finally reachs the stonepile
	state = game.GetNextState(state, input)

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Direction:  &down,
			},
		},
	}

	for range 12 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Walk,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
			},
		},
	}

	state = game.GetNextState(state, input)

	input.HelperInput = []*pb.HelperInput{}

	state = game.GetNextState(state, input)

	stone := pb.Item_Stone
	planks := pb.Item_Plank
	workbench := pb.Item_IWorkbench
	furnace := pb.Item_IFurnace

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &stone,
			},
		},
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &planks,
			},
		},
	}

	for range 4 {
		state = game.GetNextState(state, input)
	}

	workbench_items := state.BuildingStates[workbench_index].Items

	if workbench_items[0].ItemId != pb.Item_Plank {
		t.Log("stack index 0 doesn't have planks")
		t.FailNow()
	}

	if workbench_items[1].ItemId != pb.Item_Stone {
		t.Log("stack index 1 doesn't have stones")
		t.FailNow()
	}

	if workbench_items[0].Count != 4 {
		t.Log("stack index 0 doesn't have 4 planks")
		t.FailNow()
	}

	if workbench_items[1].Count != 4 {
		t.Log("stack index 1 doesn't have 4 stones")
		t.FailNow()
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &stone,
			},
		},
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &workbench,
			},
		},
	}

	state = game.GetNextState(state, input)

	workbench_items = state.BuildingStates[workbench_index].Items

	if len(workbench_items) != 1 {
		t.Log("After craft, the workbench should only have one stack")
		t.FailNow()
	}

	if workbench_items[0].ItemId != pb.Item_Stone {
		t.Log("After craft, the workbench should have stone in it")
		t.FailNow()
	}

	if workbench_items[0].ItemId != pb.Item_Stone {
		t.Log("After craft, the workbench should have one stone in it")
		t.FailNow()
	}

	if state.HelperStates[1].Items[0].ItemId != workbench {
		t.Log("After craft, helper id 1 should have a workbench")
		t.FailNow()
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &stone,
			},
		},
	}

	for range 7 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &furnace,
			},
		},
	}

	state = game.GetNextState(state, input)

	workbench_items = state.BuildingStates[workbench_index].Items

	if len(workbench_items) != 0 {
		t.Log("After craft, the workbench should be empty")
		t.FailNow()
	}

	if state.HelperStates[0].Items[0].ItemId != furnace {
		t.Log("After craft, helper id 0 should have a furnace")
		t.FailNow()
	}
}

func TestCrafting2(t *testing.T) {
	state := game.CreateInitialState()

	workbench_index := -1
	for i, state := range state.BuildingStates {
		if state.BuildingType == pb.Building_Workbench {
			workbench_index = i
			break
		}
	}

	right := pb.Direction_Right
	down := pb.Direction_Down

	input := &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{
			ActionType: pb.NecromancerActionType_Summon,
			Coordinate: &pb.Coordinate{X: 30, Y: 16},
		},
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

	state = game.GetNextState(state, input)

	input = &pb.GameInput{
		NecromancerAction: &pb.NecromancerAction{ActionType: pb.NecromancerActionType_N_Nothing},
		HelperInput: []*pb.HelperInput{
			{
				HelperId: 1,
				Action: &pb.Action{
					ActionType: pb.ActionType_Interact,
					Direction:  &right,
				},
			},
		},
	}

	for range 4 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Walk,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
			},
		},
	}

	// This will be the point where helper 0 finally reachs the stonepile
	state = game.GetNextState(state, input)

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Direction:  &down,
			},
		},
	}

	for range 12 {
		state = game.GetNextState(state, input)
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Walk,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
			},
		},
	}

	state = game.GetNextState(state, input)

	input.HelperInput = []*pb.HelperInput{}

	state = game.GetNextState(state, input)

	stone := pb.Item_Stone
	planks := pb.Item_Plank
	pickaxe := pb.Item_Pickaxe

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &stone,
			},
		},
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Give,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &planks,
			},
		},
	}

	for range 4 {
		state = game.GetNextState(state, input)
	}

	workbench_items := state.BuildingStates[workbench_index].Items

	if workbench_items[0].ItemId != pb.Item_Plank {
		t.Log("stack index 0 doesn't have planks")
		t.FailNow()
	}

	if workbench_items[1].ItemId != pb.Item_Stone {
		t.Log("stack index 1 doesn't have stones")
		t.FailNow()
	}

	if workbench_items[0].Count != 4 {
		t.Log("stack index 0 doesn't have 4 planks")
		t.FailNow()
	}

	if workbench_items[1].Count != 4 {
		t.Log("stack index 1 doesn't have 4 stones")
		t.FailNow()
	}

	input.HelperInput = []*pb.HelperInput{
		{
			HelperId: 0,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &pickaxe,
			},
		},
		{
			HelperId: 1,
			Action: &pb.Action{
				ActionType: pb.ActionType_Interact,
				Coordinate: &pb.Coordinate{X: 18, Y: 16},
				ItemId:     &pickaxe,
			},
		},
	}

	state = game.GetNextState(state, input)

	workbench_items = state.BuildingStates[workbench_index].Items

	if len(workbench_items) != 0 {
		t.Log("After craft, the workbench should be empty")
		t.FailNow()
	}

	if state.HelperStates[0].Items[1].ItemId != pickaxe {
		t.Log("After craft, helper id 0 should have a pickaxe")
		t.FailNow()
	}

	if state.HelperStates[1].Items[0].ItemId != pickaxe {
		t.Log("After craft, helper id 1 should have a pickaxe")
		t.FailNow()
	}
}
