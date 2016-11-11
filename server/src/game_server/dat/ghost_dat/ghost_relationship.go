package ghost_dat

import (
	"core/mysql"
	"fmt"
)

var (
	mapGhostRelationship map[string]*GhostRelationship
)

type GhostRelationship struct {
	GhostA     int16 // 魂侍A
	GhostAStar int8  // 魂侍A所需星级
	GhostB     int16 // 魂侍B
	GhostBStar int8  // 魂侍B所需星级
}

func loadGhostRelationship(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM ghost_relationship ORDER BY `id` ASC"), -1)
	if err != nil {
		panic(err)
	}
	iGhostA := res.Map("ghost_a")
	iGhostAStar := res.Map("ghost_a_star")
	iGhostB := res.Map("ghost_b")
	iGhostBStar := res.Map("ghost_b_star")

	var pri_id string
	mapGhostRelationship = map[string]*GhostRelationship{}
	for _, row := range res.Rows {
		pri_id = fmt.Sprintf("%d_%d", row.Int16(iGhostA), row.Int16(iGhostB))
		mapGhostRelationship[pri_id] = &GhostRelationship{
			GhostA:     row.Int16(iGhostA),
			GhostAStar: row.Int8(iGhostAStar),
			GhostB:     row.Int16(iGhostB),
			GhostBStar: row.Int8(iGhostBStar),
		}
	}
}

// ############### 对外接口实现 coding here ###############
func GetGhostRelationship(aId, bId int16) *GhostRelationship {
	// 魂侍a,b两个，存得数据的对应键值可能是a_b 或者 b_a
	aBeforeB := fmt.Sprintf("%d_%d", aId, bId)
	if result, ok := mapGhostRelationship[aBeforeB]; ok {
		return result
	}

	bBeforeA := fmt.Sprintf("%d_%d", bId, aId)
	if result, ok := mapGhostRelationship[bBeforeA]; ok {
		return result
	}
	return nil
}

func CheckGhostsRelation(relation *GhostRelationship, aId int16, aStar int8, bId int16, bStar int8) bool {
	result := false
	if relation.GhostA == aId && relation.GhostB == bId && relation.GhostAStar <= aStar && relation.GhostBStar <= bStar {
		result = true
	}

	if relation.GhostA == bId && relation.GhostB == aId && relation.GhostAStar <= bStar && relation.GhostBStar <= aStar {
		result = true
	}
	return result
}
