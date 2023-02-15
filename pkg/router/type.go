package router

type KtConf struct {
	Service  string
	Ports    [][]string
	Header   string
	Versions map[string]string
}
