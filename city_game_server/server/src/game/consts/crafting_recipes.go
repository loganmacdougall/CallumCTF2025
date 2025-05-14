package consts

import (
	"city_game/src/pb"
	"sync"
)

type ItemCraftingRecipes map[pb.Item]*pb.CraftingRecipe

var (
	crafting_recipes           map[pb.Item]*pb.CraftingRecipe
	init_once_crafting_recipes sync.Once
)

func GetCraftingRecipes() map[pb.Item]*pb.CraftingRecipe {
	init_once_crafting_recipes.Do(initialize_recipes)
	return crafting_recipes
}

func initialize_recipes() {
	crafting_recipes = map[pb.Item]*pb.CraftingRecipe{
		pb.Item_IWorkbench: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Plank, Count: 4},
				{ItemId: pb.Item_Stone, Count: 4},
			},
			Result: pb.Item_IWorkbench,
		},
		pb.Item_ICrate: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Plank, Count: 8},
			},
			Result: pb.Item_ICrate,
		},
		pb.Item_IFurnace: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Stone, Count: 8},
			},
			Result: pb.Item_IFurnace,
		},
		pb.Item_Window: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Glass, Count: 2},
				{ItemId: pb.Item_Metal, Count: 2},
			},
			Result: pb.Item_Window,
		},
		pb.Item_Door: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Plank, Count: 4},
				{ItemId: pb.Item_Metal, Count: 1},
			},
			Result: pb.Item_Door,
		},
		pb.Item_Pickaxe: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Plank, Count: 2},
				{ItemId: pb.Item_Stone, Count: 2},
			},
			Result: pb.Item_Pickaxe,
		},
		pb.Item_Bucket: {
			Requirements: []*pb.Stack{
				{ItemId: pb.Item_Metal, Count: 4},
			},
			Result: pb.Item_Bucket,
		},
	}
}
