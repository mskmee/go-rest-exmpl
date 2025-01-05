package handler

import (
	"go-rest-exmpl/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type allListsResponse struct {
	Data []entities.TodoList `json:"data"`
}

type successResponse struct {
	Success bool `json:"success"`
}

func (h *Handler) createList(c *gin.Context) {
	var input entities.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	id, err := h.service.CreateList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) updateList(c *gin.Context) {
	var input entities.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.UpdateList(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	listId := c.Param("id")

	list, err := h.service.GetListById(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) getUserTodoLists(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.service.TodoList.GetUserLists(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allListsResponse{
		Data: lists,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

	lists, err := h.service.GetAllLists()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allListsResponse{
		Data: lists,
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	listId := c.Param("id")

	err := h.service.DeleteList(listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}
