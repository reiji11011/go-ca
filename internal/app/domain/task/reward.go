package task

import "errors"

type Reward struct {
	value uint64
}

const DefaultReward = 100

func NewReward(value uint64) (Reward, error) {
	if value < DefaultReward {
		return Reward{}, errors.New("100より小さい報酬は設定できません")
	}
	return Reward{value: value}, nil
}

func (r Reward) Value() uint64 {
	return r.value
}