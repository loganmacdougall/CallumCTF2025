package nsm

import (
	"city_game/src/game/consts"
	"city_game/src/game/utility"
	"city_game/src/pb"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type PushCount struct {
	Up    int
	Right int
	Down  int
	Left  int
}

type NextStateManager struct {
	State              *pb.GameState
	Input              *pb.GameInput
	FailedManaUses     int
	Buildings          [consts.GRID_HEIGHT][consts.GRID_WIDTH]int
	Helpers            [consts.GRID_HEIGHT][consts.GRID_WIDTH]int
	BuildingsSlice     [][]int
	HelperIds          map[uint32]int
	HelperInputIds     map[uint32]int
	BuildingPushCounts []PushCount
	PushToHelpers      [][]uint32
	SummonHelperId     uint32
}

func CreateNextStateManager(state *pb.GameState, input *pb.GameInput) *NextStateManager {
	mn := &NextStateManager{}

	buildings := [consts.GRID_HEIGHT][consts.GRID_WIDTH]int{}
	helpers := [consts.GRID_HEIGHT][consts.GRID_WIDTH]int{}
	helper_ids := map[uint32]int{}
	helper_input_ids := map[uint32]int{}

	for y := range consts.GRID_HEIGHT {
		for x := range consts.GRID_WIDTH {
			buildings[y][x] = -1
			helpers[y][x] = -1
		}
	}

	building_push_counts := make([]PushCount, len(state.BuildingStates))
	push_to_helpers := make([][]uint32, len(state.BuildingStates))

	for index, building_state := range state.BuildingStates {
		x := building_state.Coordinate.X
		y := building_state.Coordinate.Y
		buildings[y][x] = index

		building_push_counts[index] = PushCount{
			Up: 0, Down: 0, Left: 0, Right: 0,
		}
	}

	buildings_slice := make([][]int, consts.GRID_HEIGHT)
	for index := range consts.GRID_HEIGHT {
		buildings_slice[index] = buildings[index][:]
	}

	next_helper_id := uint32(0)
	for index, helper_state := range state.HelperStates {
		x := helper_state.Coordinate.X
		y := helper_state.Coordinate.Y
		helpers[y][x] = index

		helper_ids[helper_state.HelperId] = index
		next_helper_id = max(next_helper_id, helper_state.HelperId)
	}
	next_helper_id += 1

	for index, helper_input := range input.HelperInput {
		helper_input_ids[helper_input.HelperId] = index
	}

	mn.State = state
	mn.FailedManaUses = 0
	mn.Input = input
	mn.Buildings = buildings
	mn.Helpers = helpers
	mn.BuildingsSlice = buildings_slice
	mn.SummonHelperId = next_helper_id
	mn.HelperIds = helper_ids
	mn.HelperInputIds = helper_input_ids
	mn.BuildingPushCounts = building_push_counts
	mn.PushToHelpers = push_to_helpers

	return mn
}

func (mn *NextStateManager) BuildingExistsAt(cord *pb.Coordinate) bool {
	x := cord.X
	y := cord.Y
	return mn.Buildings[y][x] != -1
}

func (mn *NextStateManager) BuildingAt(cord *pb.Coordinate) *pb.BuildingState {
	x := cord.X
	y := cord.Y
	index := mn.Buildings[y][x]

	if index == -1 {
		return nil
	}
	return mn.State.BuildingStates[index]
}

func (mn *NextStateManager) PlaceBuilding(cord *pb.Coordinate, item pb.Item) bool {
	if mn.BuildingExistsAt(cord) || mn.HelperExistsAt(cord) {
		return false
	}

	building_type, err := utility.ItemToBuilding(item)
	if err != nil {
		return false
	}

	building := &pb.BuildingState{
		BuildingType: building_type,
		Coordinate:   cord,
		Items:        []*pb.Stack{},
		State:        0,
	}

	index := len(mn.State.BuildingStates)
	mn.State.BuildingStates = append(mn.State.BuildingStates, building)
	mn.PushToHelpers = append(mn.PushToHelpers, []uint32{})
	mn.BuildingPushCounts = append(mn.BuildingPushCounts, PushCount{
		Up: 0, Down: 0, Left: 0, Right: 0,
	})

	mn.Buildings[cord.Y][cord.X] = index
	mn.BuildingsSlice[cord.Y][cord.X] = index

	return true
}

func (mn *NextStateManager) MoveBuilding(building *pb.BuildingState, to *pb.Coordinate) bool {
	if mn.BuildingExistsAt(to) || mn.HelperExistsAt(to) {
		return false
	}

	from := building.Coordinate
	index := mn.Buildings[from.Y][from.X]

	mn.Buildings[from.Y][from.X] = -1
	mn.BuildingsSlice[from.Y][from.X] = -1

	building.Coordinate = proto.CloneOf(to)

	mn.Buildings[to.Y][to.X] = index
	mn.BuildingsSlice[to.Y][to.X] = index

	return true
}

func (mn *NextStateManager) HelperExistsAt(cord *pb.Coordinate) bool {
	x := cord.X
	y := cord.Y
	return mn.Helpers[y][x] != -1
}

func (mn *NextStateManager) HelperAt(cord *pb.Coordinate) *pb.HelperState {
	x := cord.X
	y := cord.Y
	index := mn.Helpers[y][x]

	if index == -1 {
		return nil
	}
	return mn.State.HelperStates[index]
}

func (mn *NextStateManager) HelperIdExists(helper_id uint32) bool {
	_, ok := mn.HelperIds[helper_id]
	return ok
}

func (mn *NextStateManager) PushBuilding(building *pb.BuildingState, helper *pb.HelperState, direction pb.Direction) {
	index := mn.Buildings[building.Coordinate.Y][building.Coordinate.X]
	push_counts := mn.BuildingPushCounts[index]

	switch direction {
	case pb.Direction_Up:
		push_counts.Up += 1
	case pb.Direction_Down:
		push_counts.Down += 1
	case pb.Direction_Left:
		push_counts.Left += 1
	case pb.Direction_Right:
		push_counts.Right += 1
	}

	mn.PushToHelpers[index] = append(mn.PushToHelpers[index], helper.HelperId)
}

func (mn *NextStateManager) GetHighestPushDirection(index int) (pb.Direction, int) {
	push_count := mn.BuildingPushCounts[index]

	highest_direction := pb.Direction_Up
	highest_count := push_count.Up

	if push_count.Down > highest_count {
		highest_direction = pb.Direction_Down
		highest_count = push_count.Down
	}
	if push_count.Left > highest_count {
		highest_direction = pb.Direction_Left
		highest_count = push_count.Left
	}
	if push_count.Right > highest_count {
		highest_direction = pb.Direction_Right
		highest_count = push_count.Right
	}

	return highest_direction, highest_count
}

func (mn *NextStateManager) SummonHelper(cord *pb.Coordinate) {
	helper := &pb.HelperState{
		Coordinate: cord,
		HelperId:   uint32(mn.SummonHelperId),
		Items:      []*pb.Stack{},
		Action: &pb.Action{
			ActionType: pb.ActionType_A_Nothing,
		},
	}

	mn.State.HelperStates = append(mn.State.HelperStates, helper)

	x := cord.X
	y := cord.Y
	mn.Helpers[y][x] = int(mn.SummonHelperId)
}

func (mn *NextStateManager) ReleaseHelper(id uint32) {
	index := mn.HelperIds[id]
	mn.State.HelperStates = append(mn.State.HelperStates[:index], mn.State.HelperStates[index+1:]...)
	delete(mn.HelperIds, id)

	for index, helper_state := range mn.State.HelperStates {
		mn.HelperIds[helper_state.HelperId] = index
	}
}

func (mn *NextStateManager) UseMana(amount int) bool {
	if amount < 0 {
		mn.State.Mana += uint32(-amount)
		if mn.State.Mana > consts.MANA_MAX {
			mn.State.Mana = consts.MANA_MAX
		}
	} else if mn.State.Mana < uint32(amount) {
		mn.FailedManaUses += 1
		return false
	} else {
		mn.State.Mana -= uint32(amount)
	}

	return true
}

func (mn *NextStateManager) AddError(err string, cord *pb.Coordinate) {
	msg := &pb.ErrorMessage{Message: err, Coordinate: cord}
	mn.State.ErrorMessages = append(mn.State.ErrorMessages, msg)
}

func (mn *NextStateManager) AddPushError(building_index int) {
	helper_ids := mn.PushToHelpers[building_index]

	for _, helper_id := range helper_ids {
		cord := mn.State.HelperStates[mn.HelperIds[helper_id]].Coordinate
		mn.AddError(fmt.Sprintf("helper %d attempted to push building but failed", helper_id), cord)
	}
}
