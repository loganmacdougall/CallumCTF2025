package necromancer_p

import (
	"city_game/src/game/consts"
	"city_game/src/game/nsm"
	"city_game/src/pb"
	"fmt"
)

func HandleNecromancer(mn *nsm.NextStateManager) {
	if mn.Input.NecromancerAction == nil {
		return
	}

	switch mn.Input.NecromancerAction.ActionType {
	case pb.NecromancerActionType_N_Nothing:
		break
	case pb.NecromancerActionType_Summon:
		if mn.BuildingExistsAt(mn.Input.NecromancerAction.Coordinate) {
			mn.AddError("can't summon helper on top of building", mn.Input.NecromancerAction.Coordinate)
			return
		}

		if mn.UseMana(consts.MANA_COST_SUMMON) {
			mn.SummonHelper(mn.Input.NecromancerAction.Coordinate)
		}
	case pb.NecromancerActionType_Release:
		if !mn.HelperIdExists(*mn.Input.NecromancerAction.HelperId) {
			mn.AddError(fmt.Sprintf("can't find helper with id %d to be released", *mn.Input.NecromancerAction.HelperId), nil)
			return
		}

		if mn.UseMana(consts.MANA_COST_RELEASE) {
			mn.ReleaseHelper(*mn.Input.NecromancerAction.HelperId)
		}
	}
}
