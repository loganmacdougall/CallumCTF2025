package helper_p

import (
	"city_game/src/game/building_p"
	"city_game/src/game/consts"
	"city_game/src/game/nsm"
	"city_game/src/game/utility"
	bu "city_game/src/game/utility/building_util"
	hu "city_game/src/game/utility/helper_util"
	"city_game/src/pb"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func HandleHelpers(mn *nsm.NextStateManager) error {
	helper_queue := &hu.HelperActionQueue{}
	seen_walk_action := false

	for _, helper_state := range mn.State.HelperStates {
		action, changed := hu.GetAction(helper_state.HelperId, mn)

		if changed && !mn.UseMana(consts.MANA_COST_SET_ACTION) {
			mn.AddError(fmt.Sprintf("insufficient mana to set action of helper %d", helper_state.HelperId), helper_state.Coordinate)
		} else {
			helper_state.Action = action
		}

		helper_queue.Push(helper_state)
	}

	for helper_queue.Len() > 0 {
		helper_state := helper_queue.Pop().(*pb.HelperState)

		if helper_state.HelperId == mn.SummonHelperId {
			if _, changed := hu.GetAction(helper_state.HelperId, mn); changed {
				mn.AddError("can't give action to helper that's just been summoned", helper_state.Coordinate)
			}
			continue
		}

		if !seen_walk_action && helper_state.Action.ActionType == pb.ActionType_Walk {
			seen_walk_action = true
		}

		handle_helper(helper_state.HelperId, mn)
	}

	return nil
}

func handle_helper(helper_id uint32, mn *nsm.NextStateManager) {
	helper_state := mn.State.HelperStates[mn.HelperIds[helper_id]]

	switch helper_state.Action.ActionType {
	case pb.ActionType_A_Nothing:
		return
	case pb.ActionType_Walk:
		handle_helper_walking(helper_state, mn)
	case pb.ActionType_Give:
		handle_helper_give(helper_state, mn)
	case pb.ActionType_Take:
		handle_helper_take(helper_state, mn)
	case pb.ActionType_Place:
		handle_helper_place(helper_state, mn)
	case pb.ActionType_Push:
		handle_helper_push(helper_state, mn)
	case pb.ActionType_Interact:
		handle_helper_interact(helper_state, mn)
	}
}

func handle_helper_walking(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	action := helper_state.Action
	if action.Coordinate != nil {
		if proto.Equal(helper_state.Coordinate, action.Coordinate) {
			mn.AddError(fmt.Sprintf("helper %d with action to walk to (%d, %d) when currently at that coordinate",
				helper_state.HelperId, action.Coordinate.X, action.Coordinate.Y,
			), action.Coordinate)
			hu.CompleteAction(helper_state)
			return
		}

		path := utility.FindShortestPath(mn.BuildingsSlice, helper_state.Coordinate, action.Coordinate)
		if path == nil {
			mn.AddError(fmt.Sprintf("helper %d with action to walk to (%d, %d) but no path was found",
				helper_state.HelperId, action.Coordinate.X, action.Coordinate.Y,
			), helper_state.Coordinate)
		}

		helper_state.Coordinate = path[1]
		if len(path) == 2 {
			hu.CompleteAction(helper_state)
		} else if len(path) == 3 && mn.BuildingExistsAt(path[2]) {
			hu.CompleteAction(helper_state)
		}
		return

	} else if action.Direction != nil {
		dest, err := utility.GetOneStepFromCord(helper_state.Coordinate, *action.Direction)

		if err != nil || mn.BuildingExistsAt(dest) {
			hu.CompleteAction(helper_state)
			return
		}

		helper_state.Coordinate = dest

		dest, err = utility.GetOneStepFromCord(helper_state.Coordinate, *action.Direction)

		if err != nil || mn.BuildingExistsAt(dest) {
			hu.CompleteAction(helper_state)
			return
		}

	} else {
		mn.AddError(
			fmt.Sprintf("helper %d with action to walk without coordinate or direction specified", helper_state.HelperId),
			helper_state.Coordinate)
		hu.CompleteAction(helper_state)
		return
	}
}

func handle_helper_give(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	defer hu.CompleteAction(helper_state)

	if helper_state.Action.ItemId == nil {
		mn.AddError(
			fmt.Sprintf("helper %d with action to give but but must specify an item in the action", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	target_cord, err := utility.SolveCordFromCordDirPair(
		helper_state.Coordinate,
		helper_state.Action.Coordinate,
		helper_state.Action.Direction,
	)

	if err != nil {
		mn.AddError(fmt.Sprintf("helper %d with action to give had error getting coordinates from action: %s",
			helper_state.HelperId, err.Error()), helper_state.Coordinate)
		return
	}

	if utility.Distance(helper_state.Coordinate, target_cord) != 1 {
		mn.AddError(fmt.Sprintf("helper %d with action give must specify a neighbor coordinate but didn't",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	item := *helper_state.Action.ItemId
	stack_index := hu.HasItem(helper_state, item)

	if stack_index == -1 {
		mn.AddError(
			fmt.Sprintf("helper %d with action to give but doesn't have specified item", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	target_building := mn.BuildingAt(target_cord)
	if target_building != nil {
		if bu.AddItem(target_building, item) {
			hu.RemoveItem(helper_state, item)
		}
	} else {
		mn.AddError(
			fmt.Sprintf("helper %d with action to give but no block to receive it", helper_state.HelperId),
			helper_state.Coordinate)
	}
}

func handle_helper_take(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	defer hu.CompleteAction(helper_state)

	if helper_state.Action.ItemId == nil {
		mn.AddError(
			fmt.Sprintf("helper %d with action take but but must specify an item in the action", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	target_cord, err := utility.SolveCordFromCordDirPair(
		helper_state.Coordinate,
		helper_state.Action.Coordinate,
		helper_state.Action.Direction,
	)

	if err != nil {
		mn.AddError(fmt.Sprintf("helper %d with action take had error getting coordinates from action: %s",
			helper_state.HelperId, err.Error()), helper_state.Coordinate)
		return
	}

	if utility.Distance(helper_state.Coordinate, target_cord) != 1 {
		mn.AddError(fmt.Sprintf("helper %d with action take must specify a neighbor coordinate but didn't",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	target_building := mn.BuildingAt(target_cord)
	if target_building == nil {
		mn.AddError(
			fmt.Sprintf("helper %d with action take but no block to take from", helper_state.HelperId),
			helper_state.Coordinate)
	}

	item := *helper_state.Action.ItemId
	stack_index := bu.HasItem(target_building, item)

	if stack_index == -1 {
		mn.AddError(
			fmt.Sprintf("helper %d with action take but the building doesn't have the item specified", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	if hu.AddItem(helper_state, item) {
		bu.RemoveItem(target_building, item)
	}
}

func handle_helper_place(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	defer hu.CompleteAction(helper_state)

	if helper_state.Action.ItemId == nil {
		mn.AddError(
			fmt.Sprintf("helper %d with action place but but must specify an item in the action", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	target_cord, err := utility.SolveCordFromCordDirPair(
		helper_state.Coordinate,
		helper_state.Action.Coordinate,
		helper_state.Action.Direction,
	)

	if err != nil {
		mn.AddError(fmt.Sprintf("helper %d with action place had error getting coordinates from action: %s",
			helper_state.HelperId, err.Error()), helper_state.Coordinate)
		return
	}

	if utility.Distance(helper_state.Coordinate, target_cord) != 1 {
		mn.AddError(fmt.Sprintf("helper %d with action place must specify a neighbor coordinate but didn't",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	item := *helper_state.Action.ItemId
	stack_index := hu.HasItem(helper_state, item)

	if stack_index == -1 {
		mn.AddError(
			fmt.Sprintf("helper %d with action place but doesn't have specified item", helper_state.HelperId),
			helper_state.Coordinate)
		return
	}

	if mn.PlaceBuilding(target_cord, item) {
		hu.RemoveItem(helper_state, item)
	} else {
		mn.AddError(
			fmt.Sprintf("helper %d with action place but item failed to be placed", helper_state.HelperId),
			helper_state.Coordinate)
	}
}

func handle_helper_interact(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	defer hu.CompleteAction(helper_state)

	target_cord, err := utility.SolveCordFromCordDirPair(
		helper_state.Coordinate,
		helper_state.Action.Coordinate,
		helper_state.Action.Direction,
	)

	if err != nil {
		mn.AddError(fmt.Sprintf("helper %d with action interact had error getting coordinates from action: %s",
			helper_state.HelperId, err.Error()), helper_state.Coordinate)
		return
	}

	if utility.Distance(helper_state.Coordinate, target_cord) != 1 {
		mn.AddError(fmt.Sprintf("helper %d with action interact must specify a neighbor coordinate but didn't",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	target_building := mn.BuildingAt(target_cord)
	if target_building == nil {
		mn.AddError(fmt.Sprintf("helper %d with action interact but no building at coordinate",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	building_p.InteractWithBuilding(target_building, helper_state, mn)
}

func handle_helper_push(helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	defer hu.CompleteAction(helper_state)

	target_cord, err := utility.SolveCordFromCordDirPair(
		helper_state.Coordinate,
		helper_state.Action.Coordinate,
		helper_state.Action.Direction,
	)

	if err != nil {
		mn.AddError(fmt.Sprintf("helper %d with action push had error getting coordinates from action: %s",
			helper_state.HelperId, err.Error()), helper_state.Coordinate)
		return
	}

	if utility.Distance(helper_state.Coordinate, target_cord) != 1 {
		mn.AddError(fmt.Sprintf("helper %d with action push must specify a neighbor coordinate but didn't",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	target_building := mn.BuildingAt(target_cord)
	if target_building == nil {
		mn.AddError(fmt.Sprintf("helper %d with action push but no building at coordinate",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	if !bu.BuildingCanBePush(target_building.BuildingType) {
		mn.AddError(fmt.Sprintf("helper %d with action push but building is not of type which can be pushed",
			helper_state.HelperId), helper_state.Coordinate)
		return
	}

	dir, _ := utility.GetDirTowardsCord(helper_state.Coordinate, target_cord)
	mn.PushBuilding(target_building, helper_state, dir)
}
