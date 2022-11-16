package balls

type Ball float64

const (
	P0 Ball = 1.0
	P1      = 1.03
	P2      = 1.05
	P3      = 1.07
	P4      = 1.10
	P5      = 1.13
)

func (b Ball) String() string {
	switch b {
	case P0:
		return "P0"
	case P1:
		return "P1"
	case P2:
		return "P2"
	case P3:
		return "P3"
	case P4:
		return "P4"
	case P5:
		return "P5"
	}

	return "NA"
}
