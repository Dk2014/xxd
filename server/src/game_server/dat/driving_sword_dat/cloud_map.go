package driving_sword_dat

import (
	"math"
	"math/rand"
	//"time"
	"fmt"
)

var (
	driving_sword_cloud_map [] /*cloud_id*/ [] /*map_id*/ [] /*x*/ [] /*y*/ int8 /*event*/
	//driving_sword_cloud_map [] /*cloud_id*/ [] /*map_id*/ map[uint16]int8
)

const (
	MAP_EVENT_EMPTY       = 0 //无事件
	MAP_EVENT_BIRTH_POINT = 1 //出生点
	MAP_EVENT_HOLE        = 2 //跨层传送
	MAP_EVENT_TELEPORT    = 3 //同层传送
	MAP_EVENT_EVENT       = 4 //其他事件
)

const (
	//注意这几个产量可以减小不能增大，否则可能死循环
	NORMAL_EVENT_MIN_DISTANCE   = 3
	TELEPORT_EVENT_MIN_DISTANCE = 10
)

func genDrivingSwordCloudMap() {
	for _, cloudMap := range arrDrivingSword {
		cloudId := int16(cloudMap.Level)
		num := 1
		mapArray := [][][]int8{}
		for num < 100 {
			//for num < 50000 {
			num++
			mapDat, ok := doGenMap(uint8(cloudMap.BornX), uint8(cloudMap.BornY),
				uint8(cloudMap.HoleX), uint8(cloudMap.HoleY),
				uint8(cloudMap.Width), uint8(cloudMap.Height),
				//传送点数
				mapTeleportCount[int16(cloudMap.Level)],
				// 普通事件数量  障碍 探索 宝藏 拜访
				int(cloudMap.ObstacleCount)+int(mapExploringCount[cloudId])+int(mapTreasureCount[cloudId])+int(mapVisitingCount[cloudId]))
			if !ok {
				fmt.Println("gen map fail", cloudId)
				num--
				continue
			}
			mapArray = append(mapArray, mapDat)

		}
		driving_sword_cloud_map = append(driving_sword_cloud_map, mapArray)
	}
}

// def gen_map(point_birth, point_hole, width, heigh, num_teleport, num_event):I
func doGenMap(birth_x, birth_y, hole_x, hole_y uint8, width, height uint8, num_teleport, num_event int) (cloud_map [][]int8, success bool) {
	eventMap := make(map[uint16]int8, 30)
	eventMap[tuple2point(birth_x, birth_y)] = MAP_EVENT_BIRTH_POINT
	eventMap[tuple2point(hole_x, hole_y)] = MAP_EVENT_HOLE
	max_try := 500
	for num_teleport > 0 {
		find := false
		for !find {
			if max_try <= 0 {
				return nil, false
			}
			max_try--
			rand_x := uint8(rand.Int31n(int32(width)))
			rand_y := uint8(rand.Int31n(int32(height)))
			rand_point := tuple2point(rand_x, rand_y)
			valid := true
			for point, event := range eventMap {
				if event == MAP_EVENT_TELEPORT {
					if valid = isValidPoint(point, rand_point, TELEPORT_EVENT_MIN_DISTANCE); !valid {
						break
					}
				} else {
					if valid = isValidPoint(point, rand_point, NORMAL_EVENT_MIN_DISTANCE); !valid {
						break
					}
				}
			}
			if valid {
				num_teleport--
				eventMap[rand_point] = MAP_EVENT_TELEPORT
				find = true
			}
		}
	}
	max_try = 500
	for num_event > 0 {
		find := false
		for !find {
			if max_try <= 0 {
				return nil, false
			}
			max_try--
			rand_x := uint8(rand.Int31n(int32(width)))
			rand_y := uint8(rand.Int31n(int32(height)))
			rand_point := tuple2point(rand_x, rand_y)
			valid := true
			for point, _ := range eventMap {
				if valid = isValidPoint(point, rand_point, NORMAL_EVENT_MIN_DISTANCE); !valid {
					break
				}
			}
			if valid {
				num_event--
				eventMap[rand_point] = MAP_EVENT_EVENT
				find = true
			}
		}
	}
	cloud_map = make([][]int8, width, width)
	for x := uint8(0); x < width; x++ {
		for y := uint8(0); y < height; y++ {
			if cloud_map[x] == nil {
				cloud_map[x] = make([]int8, height, height)
			}
			event := eventMap[tuple2point(x, y)]
			cloud_map[x][y] = event
		}
	}
	return cloud_map, true
}

func tuple2point(x, y uint8) uint16 {
	return uint16(x)<<8 | uint16(y)
}

func point2tuple(point uint16) (x, y uint8) {
	return uint8(point >> 8), uint8(point & 0x0FF)
}

func calDistance(point_a_x, point_a_y, point_b_x, point_b_y uint8) int32 {
	return int32(math.Abs(float64(point_a_x)-float64(point_b_x)) + math.Abs(float64(point_a_y)-float64(point_b_y)))
}

func isValidPoint(point_a, point_b uint16, min_distance int32) bool {
	a_x, a_y := point2tuple(point_a)
	b_x, b_y := point2tuple(point_b)
	return calDistance(a_x, a_y, b_x, b_y) >= min_distance
}
