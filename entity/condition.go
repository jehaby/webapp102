package entity

type Condition string

const (
	ConditionNew           = "NEW"
	ConditionUsedLikeNew   = "USED_LIKE_NEW"
	ConditionUsed          = "USED"
	ConditionMalfunctioned = "MALFUNCTIONED"
)
