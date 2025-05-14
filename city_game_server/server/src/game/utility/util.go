package utility

import (
	"city_game/src/game/consts"
	"city_game/src/pb"
	"errors"
)

func GetOneStepFromCord(cord *pb.Coordinate, dir pb.Direction) (*pb.Coordinate, error) {
	const error_message = "coordinate out of bounds"

	switch dir {
	case pb.Direction_Up:
		if cord.Y == 0 {
			return nil, errors.New(error_message)
		}
		return &pb.Coordinate{X: cord.X, Y: cord.Y - 1}, nil
	case pb.Direction_Down:
		if cord.Y == consts.GRID_HEIGHT-1 {
			return nil, errors.New(error_message)
		}
		return &pb.Coordinate{X: cord.X, Y: cord.Y + 1}, nil
	case pb.Direction_Left:
		if cord.X == 0 {
			return nil, errors.New(error_message)
		}
		return &pb.Coordinate{X: cord.X - 1, Y: cord.Y}, nil
	case pb.Direction_Right:
		if cord.X == consts.GRID_WIDTH-1 {
			return nil, errors.New(error_message)
		}
		return &pb.Coordinate{X: cord.X + 1, Y: cord.Y}, nil
	}

	return nil, nil
}

func Distance(a *pb.Coordinate, b *pb.Coordinate) int {
	return AbsI(int(a.X)-int(b.X)) + AbsI(int(a.Y)-int(b.Y))
}

func CordInbounds(a *pb.Coordinate) bool {
	return a.X < consts.GRID_WIDTH && a.Y < consts.GRID_HEIGHT
}

func GetMaxCarryCount(item pb.Item) uint32 {
	switch item {
	case pb.Item_Pickaxe:
		return consts.ITEM_MAX_CARRY_PICKAXE
	case pb.Item_Bucket:
		return consts.ITEM_MAX_CARRY_BUCKET
	case pb.Item_Plank:
		return consts.ITEM_MAX_CARRY_PLANK
	case pb.Item_Stone:
		return consts.ITEM_MAX_CARRY_STONE
	case pb.Item_Ore:
		return consts.ITEM_MAX_CARRY_ORE
	case pb.Item_SandBucket:
		return consts.ITEM_MAX_CARRY_SANDBUCKET
	case pb.Item_Glass:
		return consts.ITEM_MAX_CARRY_GLASS
	case pb.Item_Metal:
		return consts.ITEM_MAX_CARRY_METAL
	case pb.Item_Window:
		return consts.ITEM_MAX_CARRY_WINDOW
	case pb.Item_Door:
		return consts.ITEM_MAX_CARRY_DOOR
	case pb.Item_IWorkbench:
		return consts.ITEM_MAX_CARRY_WORKBENCH
	case pb.Item_IFurnace:
		return consts.ITEM_MAX_CARRY_FURNACE
	case pb.Item_ICrate:
		return consts.ITEM_MAX_CARRY_CRATE
	default:
		return 0
	}
}

func SolveCordFromCordDirPair(origin *pb.Coordinate, target *pb.Coordinate, dir *pb.Direction) (*pb.Coordinate, error) {
	if target != nil {
		if CordInbounds(target) {
			return target, nil
		} else {
			return nil, errors.New("action has coordinate that's out of bounds")
		}
	}

	if dir == nil {
		return nil, errors.New("action doesn't specify coordinate or direction")
	}

	dir_cord, err := GetOneStepFromCord(origin, *dir)
	if err != nil {
		return nil, errors.New("action points to coordinate which is out of bounds")
	}

	return dir_cord, nil
}

func GetDirTowardsCord(origin *pb.Coordinate, target *pb.Coordinate) (pb.Direction, error) {
	if Distance(origin, target) != 1 {
		return 0, errors.New("origin and target are further than 1 away")
	}

	if target.X < origin.X {
		return pb.Direction_Left, nil
	} else if target.X > origin.X {
		return pb.Direction_Right, nil
	} else if target.Y < origin.Y {
		return pb.Direction_Up, nil
	} else {
		return pb.Direction_Down, nil
	}
}

func ItemToBuilding(item pb.Item) (pb.Building, error) {
	switch item {
	case pb.Item_IWorkbench:
		return pb.Building_Workbench, nil
	case pb.Item_IFurnace:
		return pb.Building_Furnace, nil
	case pb.Item_ICrate:
		return pb.Building_Crate, nil
	default:
		return pb.Building_Crate, errors.New("no building equivalent for that item")
	}
}

func ItemCanBeSmelted(item pb.Item) bool {
	switch item {
	case pb.Item_Ore:
		return true
	case pb.Item_SandBucket:
		return true
	default:
		return false
	}
}

func ItemSmeltedIs(item pb.Item) (pb.Item, error) {
	switch item {
	case pb.Item_Ore:
		return pb.Item_Metal, nil
	case pb.Item_SandBucket:
		return pb.Item_Glass, nil
	default:
		return 0, errors.New("item has no smelted variant")
	}
}

func StackPushRequirement(stack *pb.Stack) float64 {
	item := stack.ItemId
	count := stack.Count
	item_carry_size := GetMaxCarryCount(item)

	return float64(count) / 2 * float64(item_carry_size)
}

func GetOppositeDir(dir pb.Direction) pb.Direction {
	switch dir {
	case pb.Direction_Up:
		return pb.Direction_Down
	case pb.Direction_Down:
		return pb.Direction_Up
	case pb.Direction_Right:
		return pb.Direction_Left
	case pb.Direction_Left:
		return pb.Direction_Right
	default:
		return pb.Direction_Up
	}
}
