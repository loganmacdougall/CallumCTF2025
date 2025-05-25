package building_p

import (
	"city_game/src/game/consts"
	"city_game/src/game/nsm"
	"city_game/src/game/utility"
	bu "city_game/src/game/utility/building_util"
	hu "city_game/src/game/utility/helper_util"
	"city_game/src/pb"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func HandleBuildings(mn *nsm.NextStateManager) {
	for _, building := range mn.State.BuildingStates {
		if building.BuildingType == pb.Building_Furnace {
			handle_furnace(building, mn)
		}
	}
}

func InteractWithBuilding(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	switch building_state.BuildingType {
	case pb.Building_Lumber:
		interact_lumber(building_state, helper_state, mn)
	case pb.Building_Stonepile:
		interact_stonepile(building_state, helper_state, mn)
	case pb.Building_Mine:
		interact_mine(building_state, helper_state, mn)
	case pb.Building_Workbench:
		interact_workbench(building_state, helper_state, mn)
	case pb.Building_House:
		interact_house(building_state, helper_state, mn)
	default:
		mn.AddError(fmt.Sprintf("Helper %d with action interact interacted with a building that can't be interacted with",
			helper_state.HelperId), helper_state.Coordinate)
	}
}

func interact_lumber(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	success := hu.AddItem(helper_state, pb.Item(pb.Item_Plank))
	if !success {
		mn.AddError(fmt.Sprintf("helper %d with action interact couldn't add plank to items", helper_state.HelperId), helper_state.Coordinate)
	}
}

func interact_stonepile(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	success := hu.AddItem(helper_state, pb.Item(pb.Item_Stone))
	if !success {
		mn.AddError(fmt.Sprintf("helper %d with action interact couldn't add stone to items", helper_state.HelperId), helper_state.Coordinate)
	}
}

func interact_mine(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	if hu.HasItem(helper_state, pb.Item_Pickaxe) == -1 {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried mining but didn't have a pickaxe", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	success := hu.AddItem(helper_state, pb.Item(pb.Item_Ore))
	if !success {
		mn.AddError(fmt.Sprintf("helper %d with action interact couldn't add ore to items", helper_state.HelperId), helper_state.Coordinate)
	}
}

func interact_workbench(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	if helper_state.Action.ItemId == nil {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried interacting with a workbench but didn't specify item to craft", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	item_to_craft := helper_state.Action.ItemId
	recipes := consts.GetCraftingRecipes()

	item_recipe, found := recipes[*item_to_craft]
	if !found {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried crafting an item which can't be crafted", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	if !utility.HasItemsForRecipe(&building_state.Items, item_recipe) {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried crafting an item but the workbench doesn't have the necessary items to craft the item", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	if hu.AvailableStackIndexFor(helper_state, *item_to_craft) == -1 {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried crafting an item but doesn't have space to hold the crafted item", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	bu.RemoveItemsFromRecipe(building_state, item_recipe)
	hu.AddItem(helper_state, *item_to_craft)
}

func interact_house(building_state *pb.BuildingState, helper_state *pb.HelperState, mn *nsm.NextStateManager) {
	helper_stack_idx := -1

	for i, helper_stack := range helper_state.Items {
		j := utility.HasItemForLayer(helper_stack, mn.State.LayerRequirements)

		if j == -1 {
			continue
		}

		helper_stack_idx = i
	}

	if helper_stack_idx == -1 {
		mn.AddError(fmt.Sprintf("helper %d with action interact tried to add idea to layer when helper had no item to add to the layer requirements", helper_state.HelperId), helper_state.Coordinate)
		return
	}

	item := helper_state.Items[helper_stack_idx].ItemId
	mn.AddItemToRequirement(item)
	hu.RemoveItem(helper_state, item)
}

func handle_furnace(building_state *pb.BuildingState, mn *nsm.NextStateManager) {
	if len(building_state.Items) == 0 {
		building_state.State = 0
		return
	}

	to_smelt_index := bu.HasSmeltableItem(building_state)
	if to_smelt_index == -1 {
		building_state.State = 0
		return
	}

	if building_state.State <= consts.FURNACE_SMELTING_TICKS {
		building_state.State += 1
		return
	}

	to_smelt_stack := building_state.Items[to_smelt_index]
	if building_state.State > consts.FURNACE_SMELTING_TICKS {
		smelted_item, _ := utility.ItemSmeltedIs(to_smelt_stack.ItemId)
		smelted_item_index := bu.AvailableStackIndexFor(building_state, smelted_item)
		if smelted_item_index != -1 || to_smelt_stack.Count == 1 {
			bu.RemoveItem(building_state, to_smelt_stack.ItemId)
			bu.AddItem(building_state, smelted_item)

			to_smelt_index = bu.HasSmeltableItem(building_state)
			if to_smelt_index == -1 {
				building_state.State = 0
			} else {
				building_state.State = 1
			}
		}
	}
}

func HandlePush(mn *nsm.NextStateManager) {
	for index, building := range mn.State.BuildingStates {
		if !bu.BuildingCanBePush(building.BuildingType) {
			continue
		}

		push_direction, push_count := mn.GetHighestPushDirection(index)

		if push_count == 0 {
			continue
		}

		push_requirement := bu.PushRequirement(building)

		if push_count < push_requirement {
			mn.AddError(fmt.Sprintf("Block was being pushed by %d helpers, but due to the weight of the block had a push requirement of %d", push_count, push_requirement), building.Coordinate)
			mn.AddPushError(index)
			continue
		}

		push_to_location, err := utility.GetOneStepFromCord(building.Coordinate, push_direction)
		if err != nil {
			mn.AddError("Attempted to push block out of bounds which is not allowed", building.Coordinate)
			mn.AddPushError(index)
			continue
		}

		if mn.BuildingExistsAt(push_to_location) || mn.HelperExistsAt(push_to_location) {
			mn.AddError("Attempted to push block into a coordinate that's not empty", building.Coordinate)
			mn.AddPushError(index)
			continue
		}

		pushed_from_direction := utility.GetOppositeDir(push_direction)
		previous_building_cord := building.Coordinate
		pushed_from, _ := utility.GetOneStepFromCord(building.Coordinate, pushed_from_direction)

		mn.MoveBuilding(building, push_to_location)

		for _, helper_id := range mn.PushToHelpers[index] {
			helper_index := mn.HelperIds[helper_id]
			helper := mn.State.HelperStates[helper_index]

			if proto.Equal(helper.Coordinate, pushed_from) {
				helper.Coordinate = proto.CloneOf(previous_building_cord)
			} else {
				mn.AddError(fmt.Sprintf("helper %d with action push failed to push building in their specified direction", helper_id),
					helper.Coordinate)
			}
		}

	}
}
