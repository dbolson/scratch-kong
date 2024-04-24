package flags

type FlagsCmd struct {
	Get  GetCmd  `kong:"cmd"`
	List ListCmd `kong:"cmd"`
}
