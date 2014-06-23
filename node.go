package at

type proto string

const (
	protohttp  proto = "http"
	prototcp   proto = "tcp"
	protohttps proto = "https"
)

type Node struct {
	Id string
}
