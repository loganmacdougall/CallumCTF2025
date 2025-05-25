package consts

import (
	"city_game/src/pb"

	"google.golang.org/protobuf/proto"
)

var layer_requirements [][]*pb.ItemRequirement

func GetLayerRequirements(layer int) []*pb.ItemRequirement {
	initialize_layers()

	if layer > len(layer_requirements) {
		layer = len(layer_requirements) - 1
	}

	reqs := []*pb.ItemRequirement{}
	ref_reqs := layer_requirements[layer]

	for _, r := range ref_reqs {
		reqs = append(reqs, proto.CloneOf(r))
	}

	return reqs
}

func initialize_layers() {
	if layer_requirements != nil {
		return
	}

	layer_requirements = [][]*pb.ItemRequirement{
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  50,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  50,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  0,
				Total:  50,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  50,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  0,
				Total:  50,
			},
			{
				ItemId: pb.Item_Door,
				Count:  0,
				Total:  5,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  0,
				Total:  5,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  100,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  0,
				Total:  100,
			},
			{
				ItemId: pb.Item_Door,
				Count:  0,
				Total:  10,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  0,
				Total:  10,
			},
			{
				ItemId: pb.Item_Window,
				Count:  0,
				Total:  4,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  150,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  0,
				Total:  100,
			},
			{
				ItemId: pb.Item_Door,
				Count:  0,
				Total:  10,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  0,
				Total:  20,
			},
			{
				ItemId: pb.Item_Window,
				Count:  0,
				Total:  4,
			},
		},
		{
			{
				ItemId: pb.Item_Stone,
				Count:  0,
				Total:  250,
			},
			{
				ItemId: pb.Item_Plank,
				Count:  0,
				Total:  150,
			},
			{
				ItemId: pb.Item_Metal,
				Count:  0,
				Total:  50,
			},
		},
	}
}
