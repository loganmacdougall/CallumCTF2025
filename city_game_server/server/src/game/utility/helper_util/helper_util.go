package helper_util

import (
	"city_game/src/game/consts"
	"city_game/src/game/nsm"
	"city_game/src/game/utility"
	"city_game/src/pb"
)

func HasItem(helper *pb.HelperState, item pb.Item) int {
	return utility.HasItem(&helper.Items, item)
}

func HasNOfItem(helper *pb.HelperState, item pb.Item, count uint32) bool {
	return utility.HasNOfItem(&helper.Items, item, count)
}

func HasItemsForRecipe(helper *pb.HelperState, recipe *pb.CraftingRecipe) bool {
	return utility.HasItemsForRecipe(&helper.Items, recipe)
}

func AvailableStackIndexFor(helper *pb.HelperState, item pb.Item) int {
	return utility.AvailableStackIndexFor(&helper.Items, item, consts.HELPER_STACK_COUNT, false)
}

func GetAction(helper_id uint32, mn *nsm.NextStateManager) (*pb.Action, bool) {
	index, found := mn.HelperInputIds[helper_id]
	if found {
		return mn.Input.HelperInput[index].Action, true
	}
	return mn.State.HelperStates[mn.HelperIds[helper_id]].Action, false
}

func CompleteAction(helper_state *pb.HelperState) {
	nothing_action := &pb.Action{
		ActionType: pb.ActionType_A_Nothing,
	}
	helper_state.Action = nothing_action
}

func AddItem(helper *pb.HelperState, item pb.Item) bool {
	return utility.AddItem(&helper.Items, item, consts.HELPER_STACK_COUNT, false)
}

func RemoveItem(helper *pb.HelperState, item pb.Item) bool {
	return utility.RemoveItem(&helper.Items, item, consts.HELPER_STACK_COUNT, false)
}

func RemoveNOfItem(helper *pb.HelperState, item pb.Item, count uint32) bool {
	return utility.RemoveNOfItem(&helper.Items, item, count)
}
