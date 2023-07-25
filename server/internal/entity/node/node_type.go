package node

type Type int

const (
	DATASHEET Type = iota
	DASHBOARD
)

// 定义 ChipType 类型的方法 String(), 返回字符串。
func (t Type) String() string {
	switch t {
	case DATASHEET:
		return "datasheet"
	case DASHBOARD:
		return "dashboard"
	default:
		return "NA"
	}
}
