package building_util

import (
	"city_game/src/game/consts"
	"city_game/src/game/utility"
	"city_game/src/pb"
)

func HasItem(building *pb.BuildingState, item pb.Item) int {
	return utility.HasItem(&building.Items, item)
}

func HasNOfItem(building *pb.BuildingState, item pb.Item, count uint32) bool {
	return utility.HasNOfItem(&building.Items, item, count)
}

func HasItemsForRecipe(building *pb.BuildingState, recipe *pb.CraftingRecipe) bool {
	return utility.HasItemsForRecipe(&building.Items, recipe)
}

func AvailableStackIndexFor(building *pb.BuildingState, item pb.Item) int {
	is_crate := building.BuildingType == pb.Building_Crate
	stack_count := GetStackCount(building.BuildingType)

	return utility.AvailableStackIndexFor(&building.Items, item, stack_count, is_crate)
}

func GetStackCount(building pb.Building) int {
	switch building {
	case pb.Building_Crate:
		return consts.BUILDING_STACK_COUNT_CRATE
	case pb.Building_Furnace:
		return consts.BUILDING_STACK_COUNT_FURNACE
	case pb.Building_Workbench:
		return consts.BUILDING_STACK_COUNT_WORKBENCH
	default:
		return 0
	}
}

func AddItem(building *pb.BuildingState, item pb.Item) bool {
	is_crate := building.BuildingType == pb.Building_Crate
	stack_count := GetStackCount(building.BuildingType)

	return utility.AddItem(&building.Items, item, stack_count, is_crate)
}

func RemoveItem(building *pb.BuildingState, item pb.Item) bool {
	is_crate := building.BuildingType == pb.Building_Crate
	stack_count := GetStackCount(building.BuildingType)

	return utility.RemoveItem(&building.Items, item, stack_count, is_crate)
}

func RemoveNOfItem(building *pb.BuildingState, item pb.Item, count uint32) bool {
	return utility.RemoveNOfItem(&building.Items, item, count)
}

func RemoveItemsFromRecipe(building *pb.BuildingState, recipe *pb.CraftingRecipe) bool {
	return utility.RemoveItemsFromRecipe(&building.Items, recipe)
}

func HasSmeltableItem(building *pb.BuildingState) int {
	if len(building.Items) == 0 {
		return -1
	}

	for index, stack := range building.Items {
		if utility.ItemCanBeSmelted(stack.ItemId) {
			return index
		}
	}

	return -1
}

func PushRequirement(building *pb.BuildingState) int {
	req := 0.0

	for _, stack := range building.Items {
		req += utility.StackPushRequirement(stack)
	}

	return 1 + int(req)
}

func BuildingCanBePush(building pb.Building) bool {
	switch building {
	case pb.Building_Crate:
		return true
	case pb.Building_Furnace:
		return true
	case pb.Building_Workbench:
		return true
	default:
		return false
	}
}
