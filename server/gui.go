package server

import (
	"time"
)

import (
	"bytes"
	"github.com/HorizontDimension/go-restful/swagger"
	"github.com/emicklei/go-restful"
	"github.com/thesyncim/at/msg"
	"io"
	"log"
	"mime"
	"net/http"
	"regexp"

	"path/filepath"
	"strings"
)

type NodeService struct {
	Url string
	Req msg.ReqTunnel
}

// +gen *
type Node struct {
	Id       string
	LastPing time.Time
	Services []NodeService
}

type NodeResource struct {
	// normally one would use DAO (data access object)
	clients *ControlRegistry
}

func (u NodeResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/nodes").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/all").To(u.allNodes).
		// docs
		Doc("get a node ").
		Operation("findNode").
		Writes([]Node{})) // on the response

	ws.Route(ws.GET("/{node-id}").To(u.getNode).
		// docs
		Doc("get a node ").
		Operation("getNode").
		Param(ws.PathParameter("node-id", "identifier of the node").DataType("string")).
		Writes(Node{})) // on the response
	ws.Route(ws.GET("/search/{node-id}").To(u.searchNodes).
		// docs
		Doc("search a node ").
		Operation("searchNode").
		Param(ws.PathParameter("node-id", "query to search").DataType("string")).
		Writes(Node{})) // on the response

	container.Add(ws)
}

//
func (n *NodeResource) getNode(request *restful.Request, response *restful.Response) {

	id := request.PathParameter("node-id")
	control := n.clients.Get(id)
	if control == nil {
		response.WriteAsJson(nil)
		return

	}

	nodeservices := []NodeService{}

	for _, tunnel := range control.tunnels {

		service := NodeService{}
		service.Req = msg.ReqTunnel{}
		service.Req = *tunnel.req
		service.Url = tunnel.url

		nodeservices = append(nodeservices, service)

	}

	node := Node{
		Id:       control.id,
		LastPing: control.lastPing,
	}
	node.Services = []NodeService{}
	node.Services = nodeservices
	response.WriteAsJson(node)
}

//
func (u *NodeResource) allNodes(request *restful.Request, response *restful.Response) {
	allnodes := u.clients.All()
	log.Println("%v", allnodes)

	nodes := []Node{}

	for _, node := range allnodes {

		nodeservices := []NodeService{}

		for _, tunnel := range node.tunnels {

			service := NodeService{}
			service.Req = msg.ReqTunnel{}
			service.Req = *tunnel.req

			service.Url = tunnel.url

			nodeservices = append(nodeservices, service)
		}

		node := Node{
			Id:       node.id,
			LastPing: node.lastPing,
		}
		node.Services = []NodeService{}
		node.Services = nodeservices

		nodes = append(nodes, node)

	}

	response.WriteAsJson(nodes)
}

func (u *NodeResource) searchNodes(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("node-id")
	allnodes := u.clients.All()
	log.Println("%v", allnodes)

	nodes := []*Node{}

	for _, node := range allnodes {

		nodeservices := []NodeService{}

		for _, tunnel := range node.tunnels {

			service := NodeService{}
			service.Req = msg.ReqTunnel{}
			service.Req = *tunnel.req

			service.Url = tunnel.url

			nodeservices = append(nodeservices, service)
		}

		node := Node{
			Id:       node.id,
			LastPing: node.lastPing,
		}
		node.Services = []NodeService{}
		node.Services = nodeservices

		nodes = append(nodes, &node)

	}

	nameMatch := func(n *Node) bool {
		res, err := regexp.MatchString(id, n.Id)
		if err != nil {
			log.Println("we got an error")
		}
		return res
	}
	nod := Nodes(nodes)
	searchResults := nod.Where(nameMatch)
	if searchResults == nil {
		response.Write([]byte("[]"))
		return
	}

	response.WriteAsJson(searchResults)

}

func InitRestInferface(c *ControlRegistry) {

	wsContainer := restful.NewContainer()

	nodeR := NodeResource{c}
	nodeR.Register(wsContainer)

	//wsContainer.Filter(wsContainer.OPTIONSFilter)
	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"AnnyTunnelApi"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Authorization"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://proxy.euroneves.pt:8080",
		ApiPath:        "/apidocs.json",

		// Optionally, specifiy where the UI is located
		SwaggerPath: "/apidocs/",
		//	SwaggerFilePath: "/root/gocode/src/github.com/HorizontDimension/twiit/swagger-ui/dist",
		StaticHandler: &BinaryHandler{},
	}

	swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on localhost:8080")
	//server := &http.Server{Addr: ":80", Handler: wsContainer}
	//log.Fatal(server.ListenAndServe())

	log.Fatalln(http.ListenAndServe(":8080", wsContainer))
}

type BinaryHandler struct {
}

func (b *BinaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	file := strings.TrimPrefix(r.RequestURI, "/apidocs/")
	if file == "" {
		file = "index.html"
	}
	mimetype := mime.TypeByExtension(filepath.Ext(file))
	//file = "dist/" + file

	w.Header().Set("Content-Type", mimetype)

	data, err := Asset(file)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	if len(data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	io.Copy(w, bytes.NewBuffer(data))
}
