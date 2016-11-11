package awaken_dat

import (
	"core/mysql"
)

func Load(db *mysql.Connection) {
	loadAwakenAttr(db)
	loadAwakenGraphic(db)
}

var (
	g_AwakenAttrs    map[int16]*AwakenAttr
	g_AwakenGraphics map[int8]map[int16]*AwakenGraphic
)

// 觉醒属性
type AwakenAttr struct {
	Id      int16
	IsSkill bool  // 是否为技能
	SkillId int16 // 技能id
	Type    int8  // 属性类型
	Attr    int32 // 属性值
	Lights  int8  // 希望之光需求量
}

func loadAwakenAttr(db *mysql.Connection) {
	sql := "select * from `awaken_attr`;"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	g_AwakenAttrs = make(map[int16]*AwakenAttr)

	iId := res.Map("id")
	iIsSkill := res.Map("is_skill")
	iSkillId := res.Map("skill_id")
	iType := res.Map("type")
	iAttr := res.Map("attr")
	iLights := res.Map("lights")

	for _, row := range res.Rows {
		id := row.Int16(iId)
		g_AwakenAttrs[id] = &AwakenAttr{
			Id:      id,
			IsSkill: row.Int32(iIsSkill) > 0,
			SkillId: row.Int16(iSkillId),
			Type:    row.Int8(iType),
			Attr:    row.Int32(iAttr),
			Lights:  row.Int8(iLights),
		}
	}
}

// 觉醒关联
type AwakenGraphic struct {
	Id      int32
	RoleId  int8  // 角色Id
	ImplId  int16 // 关联实例Id
	AttrId  int16 // 属性Id
	DepImpl int16 // 依赖实例
}

func loadAwakenGraphic(db *mysql.Connection) {
	sql := "select * from `awaken_graphic`;"
	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	g_AwakenGraphics = make(map[int8]map[int16]*AwakenGraphic)

	iId := res.Map("id")
	iRoleId := res.Map("role_id")
	iImplId := res.Map("impl_id")
	iAttrId := res.Map("attr_id")
	iDepImpl := res.Map("dep_impl")

	for _, row := range res.Rows {
		role_id := row.Int8(iRoleId)
		impl_id := row.Int16(iImplId)
		if g_AwakenGraphics[role_id] == nil {
			g_AwakenGraphics[role_id] = make(map[int16]*AwakenGraphic)
		}
		g_AwakenGraphics[role_id][impl_id] = &AwakenGraphic{
			Id:      row.Int32(iId),
			RoleId:  row.Int8(iRoleId),
			ImplId:  row.Int16(iImplId),
			AttrId:  row.Int16(iAttrId),
			DepImpl: row.Int16(iDepImpl),
		}
	}
}

// 查询觉醒属性
func GetAwakenAttr(id int16) *AwakenAttr {
	return g_AwakenAttrs[id]
}

// 判断觉醒属性是否为技能
func IsSkillAttr(role_id int8, impl_id int16) bool {
	return GetRoleAttr(role_id, impl_id).IsSkill
}

// 查询觉醒关联
func GetAwakenGraphic(role_id int8, impl_id int16) *AwakenGraphic {
	return g_AwakenGraphics[role_id][impl_id]
}

// 依赖实例ID
func GetDepImplId(role_id int8, impl_id int16) int16 {
	return GetAwakenGraphic(role_id, impl_id).DepImpl
}

// 角色觉醒属性
func GetRoleAttr(role_id int8, impl_id int16) *AwakenAttr {
	// 获取觉醒属性ID
	attr_id := GetAwakenGraphic(role_id, impl_id).AttrId
	return GetAwakenAttr(attr_id)
}
