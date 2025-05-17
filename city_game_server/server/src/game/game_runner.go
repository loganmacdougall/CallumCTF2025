package game

import (
	"city_game/src/game/building_p"
	"city_game/src/game/consts"
	"city_game/src/game/helper_p"
	"city_game/src/game/necromancer_p"
	"city_game/src/game/nsm"
	"city_game/src/pb"

	"google.golang.org/protobuf/proto"
)

func GetNextState(state *pb.GameState, input *pb.GameInput) *pb.GameState {
	next_state := proto.CloneOf(state)
	next_state.ErrorMessages = []*pb.ErrorMessage{}
	next_state.Tick += 1

	mn := nsm.CreateNextStateManager(next_state, input)
	mn.UseMana(-consts.MANA_REGENERATION)

	necromancer_p.HandleNecromancer(mn)
	helper_p.HandleHelpers(mn)
	building_p.HandleBuildings(mn)

	return next_state
}

func CreateInitialState() *pb.GameState {
	state := &pb.GameState{
		Tick:              0,
		Layer:             0,
		Mana:              consts.MANA_MAX,
		LayerRequirements: consts.GetLayerRequirements(0),
		GameId:            "1234-5678-90ab-cdef",
	}

	add_building := func(building pb.Building, x int, y int) {
		building_state := &pb.BuildingState{
			BuildingType: building,
			Coordinate:   &pb.Coordinate{X: uint32(x), Y: uint32(y)},
			Items:        []*pb.Stack{},
			State:        0,
		}

		state.BuildingStates = append(state.BuildingStates, building_state)
	}

	add_building(pb.Building_House, 19, 5)
	add_building(pb.Building_House, 19, 6)
	add_building(pb.Building_House, 20, 5)
	add_building(pb.Building_House, 20, 6)

	add_building(pb.Building_Mine, 7, 16)
	add_building(pb.Building_Mine, 8, 16)

	add_building(pb.Building_Workbench, 18, 16)
	add_building(pb.Building_Stonepile, 19, 16)
	add_building(pb.Building_Stonepile, 20, 16)

	add_building(pb.Building_Lumber, 31, 16)
	add_building(pb.Building_Lumber, 32, 16)

	helper := &pb.HelperState{
		HelperId:   0,
		Coordinate: &pb.Coordinate{X: 20, Y: 9},
		Items:      []*pb.Stack{},
		Action:     &pb.Action{ActionType: pb.ActionType_A_Nothing},
	}

	state.HelperStates = append(state.HelperStates, helper)

	return state
}
