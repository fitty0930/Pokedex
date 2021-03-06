package pokedex

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

// Register ... es como si fuera mi "router"
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}

func makeEndpoints(s Service) []*endpoint { // es mi pokedexservice
	list := []*endpoint{} // creo lo que se necesita en register
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/pokedex",
		function: getAll(s),
	})
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/pokedex/:ID",
		function: getOne(s),
	})
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/pokedex/:name",
		function: postOne(s),
	})
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/pokedex/:ID",
		function: deleteOne(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/pokedex/:ID/:name",
		function: changeOne(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"pokedex": s.FindAll(),
		})
	}
}

func getOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("ID"), 0, 64) // hago esto porque lo trae como string
		if err == nil {
			p := s.FindByID(id)
			c.JSON(http.StatusOK, gin.H{
				"pokemon": p,
			})
		}

	}
}

func deleteOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("ID"), 0, 64) // hago esto porque lo trae como string
		if err == nil {
			s.DeleteByID(id)
			c.JSON(http.StatusOK, gin.H{})
		}

	}
}

func postOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		s.AddPokemon(name)
		c.JSON(http.StatusOK, gin.H{})

	}
}

func changeOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		id, err := strconv.ParseInt(c.Param("ID"), 0, 64)
		if err == nil {
			s.ChangePokemon(id, name)
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}
