package battle

import (
//"math/rand"
)

type JobInfo struct {
	Pos        int // 位置的索引
	SkillIndex int
	Type       int
	JobLevel   int
}

func (this *JobInfo) Apply(side *SideInfo) {
	if this == nil {
		return
	}

	f := side.Fighters[this.Pos]
	f.useSkillIndex = this.SkillIndex
	f.useSkillJob = f.SkillInfos[this.SkillIndex].SkillJob
}
