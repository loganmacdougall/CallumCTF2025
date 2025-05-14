package utility

import (
	"city_game/src/game/consts"
	"city_game/src/pb"
)

func HasItem(stacks *[]*pb.Stack, item pb.Item) int {
	fewest_items_slot := -1
	fewest_items := 299

	for i, stack := range *stacks {
		if stack.ItemId == item {
			if stack.Count < uint32(fewest_items) {
				fewest_items_slot = i
				fewest_items = int(stack.Count)
			}
		}
	}

	return fewest_items_slot
}

func HasNOfItem(stacks *[]*pb.Stack, item pb.Item, count uint32) bool {
	current_count := uint32(0)

	for _, stack := range *stacks {
		if stack.ItemId == item {
			current_count += stack.Count

			if current_count >= count {
				return true
			}
		}

	}

	return false
}

func HasItemsForRecipe(stacks *[]*pb.Stack, recipe *pb.CraftingRecipe) bool {
	for _, req := range recipe.Requirements {
		if !HasNOfItem(stacks, req.ItemId, req.Count) {
			return false
		}
	}

	return true
}

func AvailableStackIndexFor(stacks *[]*pb.Stack, item pb.Item, stack_count int, is_crate bool) int {
	max_carry_count := GetMaxCarryCount(item)

	if is_crate {
		max_carry_count *= consts.ITEM_MAX_CRATE_MULTIPLIER
	}

	for i, stack := range *stacks {
		if stack.ItemId != item {
			continue
		}

		if stack.Count < max_carry_count {
			return i
		}
	}

	if len(*stacks) < stack_count {
		return len(*stacks)
	}

	return -1
}

func AddItem(stacks *[]*pb.Stack, item pb.Item, stack_count int, is_crate bool) bool {
	stack_index := AvailableStackIndexFor(stacks, item, stack_count, is_crate)

	if stack_index == -1 {
		return false
	}

	if len(*stacks) <= stack_index {
		*stacks = append(*stacks, &pb.Stack{
			ItemId: item,
			Count:  1,
		})
	} else {
		(*stacks)[stack_index].Count += 1
	}

	return true
}

func RemoveItem(stacks *[]*pb.Stack, item pb.Item, stack_count int, is_crate bool) bool {
	stack_index := AvailableStackIndexFor(stacks, item, stack_count, is_crate)

	if stack_index == -1 {
		return false
	}

	(*stacks)[stack_index].Count -= 1

	if (*stacks)[stack_index].Count == 0 {
		*stacks = append((*stacks)[:stack_index], (*stacks)[stack_index+1:]...)
	}

	return true
}

func RemoveNOfItem(stacks *[]*pb.Stack, item pb.Item, count uint32) bool {
	if !HasNOfItem(stacks, item, count) {
		return false
	}

	count_left_to_remove := count

	for count_left_to_remove != 0 {
		stack_index := HasItem(stacks, item)
		stack_count := (*stacks)[stack_index].Count

		if stack_count <= count_left_to_remove {
			*stacks = append((*stacks)[:stack_index], (*stacks)[stack_index+1:]...)
			count_left_to_remove -= stack_count
		} else {
			(*stacks)[stack_index].Count -= count_left_to_remove
			count_left_to_remove = 0
		}
	}

	return true
}

func RemoveItemsFromRecipe(stacks *[]*pb.Stack, recipe *pb.CraftingRecipe) bool {
	for _, stack := range recipe.Requirements {
		if !HasNOfItem(stacks, stack.ItemId, stack.Count) {
			return false
		}
	}

	for _, stack := range recipe.Requirements {
		RemoveNOfItem(stacks, stack.ItemId, stack.Count)
	}

	return true
}
