package ds

import (
	"math/rand"
	"time"
)

type ReservoirSamplingAlgorithm struct {
	Limit int   //抽样数
	Index int   //当前输入的元素数量
	Items []int //结果集
}

func (rs *ReservoirSamplingAlgorithm) ReservoirSampling(item int) {
	rs.Index++
	if len(rs.Items) < rs.Limit {
		rs.Items = append(rs.Items, item)
		return
	}
	rate := float64(rs.Limit) / float64(rs.Index)
	for i := 0; i < len(rs.Items); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() <= rate {
			rs.Items[i] = item
			break
		}
		time.Sleep(time.Microsecond)
	}
}
