package driving_sword_dat

import (
	"core/mysql"
)

var (
	arrDrivingSword  []*DrivingSword
	mapTeleportCount map[int16]int
)

type DrivingSword struct {
	Level         int8 // 层级
	Width         int8 // 地图宽
	Height        int8 // 地图高
	BornX         int8 // 出生地坐标x
	BornY         int8 // 出生地坐标y
	HoleX         int8 // 传送阵坐标x
	HoleY         int8 // 传送阵坐标y
	ObstacleCount int8 // 障碍总数
}

func loadDrivingSword(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("SELECT * FROM `driving_sword` ORDER BY `level` ASC"), -1)
	if err != nil {
		panic(err)
	}

	iLevel := res.Map("level")
	iWidth := res.Map("width")
	iHeight := res.Map("height")
	iBornX := res.Map("born_x")
	iBornY := res.Map("born_y")
	iHoleX := res.Map("hole_x")
	iHoleY := res.Map("hole_y")
	iObstacleCount := res.Map("obstacle_count")

	arrDrivingSword = []*DrivingSword{}
	for _, row := range res.Rows {
		arrDrivingSword = append(arrDrivingSword, &DrivingSword{
			Level:         row.Int8(iLevel),
			Width:         row.Int8(iWidth),
			Height:        row.Int8(iHeight),
			BornX:         row.Int8(iBornX),
			BornY:         row.Int8(iBornY),
			HoleX:         row.Int8(iHoleX),
			HoleY:         row.Int8(iHoleY),
			ObstacleCount: row.Int8(iObstacleCount),
		})
	}
}

func loadTeleportCount(db *mysql.Connection) {
	res, err := db.ExecuteFetch([]byte("select `cloud_id`, count(`cloud_id`) `count` from `driving_sword_teleport` group by `cloud_id`;"), -1)
	if err != nil {
		panic(err)
	}

	iCloudId := res.Map("cloud_id")
	iCount := res.Map("count")

	mapTeleportCount = make(map[int16]int)
	for _, row := range res.Rows {
		cloud_id := row.Int16(iCloudId)
		mapTeleportCount[cloud_id] = int(row.Int8(iCount))
	}
}

func GetCloudLevel(level int16) *DrivingSword {
	return arrDrivingSword[level-1]
}

func CountObstacleByCloud(cloud int16) int {
	return int(GetCloudLevel(cloud).ObstacleCount)
}

func CountTeleportByCloud(cloud int16) int {
	return mapTeleportCount[cloud]
}
