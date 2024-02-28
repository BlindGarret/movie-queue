package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/BlindGarret/lexorank"
	"github.com/BlindGarret/movie-queue/cmd/cfg"
	"github.com/BlindGarret/movie-queue/ent"
	"github.com/BlindGarret/movie-queue/ent/show"
	"github.com/labstack/echo/v4"
)

const RankStepSize uint64 = 1000

type errorData struct {
	Message string
}

type FormData struct {
	Values map[string]string
}

func newFormData() FormData {
	return FormData{Values: make(map[string]string)}
}

type indexPageData struct {
	Shows          ent.Shows
	AddFormData    FormData
	FilterFormData FormData
	Error          errorData
}

func handleIndex(c echo.Context) error {
	client := c.Get(cfg.DBXContextKey).(*ent.Client)
	shows, err := client.Show.Query().All(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Error querying shows: %v", err)
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error querying shows!"})
	}
	return c.Render(http.StatusOK, "index-page", indexPageData{Shows: shows})
}

func handleGetShows(c echo.Context) error {
	client := c.Get(cfg.DBXContextKey).(*ent.Client)
	filter := c.QueryParam("filter")
	var shows ent.Shows
	var err error
	if filter == "" {
		shows, err = client.Show.Query().Order(show.ByOrder()).All(c.Request().Context())
		if err != nil {
			c.Logger().Errorf("Error querying shows: %v", err)
			return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error querying shows!"})
		}
	} else {
		shows, err = client.Show.Query().Where(show.NameContainsFold(filter)).All(c.Request().Context())
	}

	return c.Render(http.StatusOK, "show-set", shows)
}

func handleAddShow(c echo.Context) error {
	client := c.Get(cfg.DBXContextKey).(*ent.Client)
	nextRank, err := getNextRank(client)
	if err != nil {
		c.Logger().Errorf("Error getting next rank: %v", err)
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error getting next rank!"})
	}

	show, err := client.Show.
		Create().
		SetName(c.FormValue("name")).
		SetOrder(nextRank).
		SetType(show.Type(c.FormValue("type"))).
		Save(c.Request().Context())

	// Failure, keep form values, and render error
	if err != nil {
		c.Logger().Errorf("Error creating show: %v", err)
		formData := newFormData()
		formData.Values["name"] = c.FormValue("name")
		err = c.Render(http.StatusUnprocessableEntity, "add-show-form", formData)
		if err != nil {
			c.Logger().Errorf("Error rendering add-show-form: %v", err)
			return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error rendering add-show-form!"})
		}
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error creating show!"})
	}

	// Success, render OOB and Form
	err = c.Render(http.StatusOK, "add-show-form", newFormData())
	if err != nil {
		c.Logger().Errorf("Error rendering add-show-form: %v", err)
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error rendering add-show-form!"})
	}
	return c.Render(http.StatusFound, "oob-show", show)
}

func handleDeleteShow(c echo.Context) error {
	client := c.Get(cfg.DBXContextKey).(*ent.Client)
	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.Logger().Errorf("Error parsing id: %v", err)
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error parsing id!"})
	}

	err = client.Show.DeleteOneID(id).Exec(c.Request().Context())
	if err != nil {
		c.Logger().Errorf("Error deleting show: %v", err)
		return c.Render(http.StatusUnprocessableEntity, "error-toast", errorData{Message: "Error deleting show!"})
	}
	return c.NoContent(http.StatusOK)
}

func getNextRank(dbx *ent.Client) (string, error) {
	ctx := context.Background()
	tx, err := dbx.Tx(ctx)
	if err != nil {
		return "", err
	}

	// lock the table so we don't get multiple people reading this!
	rankData, err := tx.Rank.Query().ForUpdate().Only(ctx)
	if err != nil {
		return "", err
	}

	nextRank, err := lexorank.Next(rankData.Next, rankData.Next, RankStepSize)
	if err != nil {
		return "", err
	}

	returnRank := rankData.Next
	rankData.Next = nextRank

	_, err = rankData.Update().SetNext(nextRank).Save(ctx)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return returnRank, nil
}
