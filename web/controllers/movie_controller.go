package controllers

import (
	"errors"

	"vmc-golang/models"
	"vmc-golang/services"

	"github.com/kataras/iris"
)

type MovieController struct {
	Service services.MovieService
}

// otherwise just return the datamodels.
func (c *MovieController) Get() (results []datamodels.Movie) {
	return c.Service.GetAll()
}

// curl -i http://localhost:8080/movies/1
func (c *MovieController) GetBy(id int64) (movie datamodels.Movie, found bool) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}

// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MovieController) PutBy(ctx iris.Context, id int64) (datamodels.Movie, error) {

	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return datamodels.Movie{}, errors.New("failed due form file 'poster' missing")
	}

	file.Close()

	poster := info.Filename
	genre := ctx.FormValue("genre")

	return c.Service.UpdatePosterAndGenreByID(id, poster, genre)
}

// DeleteBy deletes a movie.
func (c *MovieController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		return iris.Map{"deleted": id}
	}
	return iris.StatusBadRequest
}
