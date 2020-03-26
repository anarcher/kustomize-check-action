package finder

type PathFinder interface {
	PathFind() ([]string, error)
}
