package command

type Commander interface {
	Run(name string, args ...string) (string, error)
}
