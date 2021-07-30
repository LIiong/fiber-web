package task

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"fiber-web/api"
	"fiber-web/model"
)

// 查询
// @Tags task
// @Summary 查询
// @Description 查询
// @Security jwt
// @Accept json
// @Produce json
// @Success 200 {object} api.ResponseHTTP{data=[]model.TodoList}
// @Router /api/task/list [get]
func FindAll(c *fiber.Ctx) error {
	m := new(model.TodoList)
	todoLists, err := m.TodoLists()
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, todoLists)
}

// Save 保存
// @Summary 保存
// @Description 保存
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param task body model.TodoList true "save task"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/save [post]
func Save(c *fiber.Ctx) error {
	m := new(model.TodoList)
	if err := c.BodyParser(m); err != nil {
		return api.Response(c, err, nil)
	}
	id, err := m.Save()
	if err != nil {
		log.Println("Error while save task:", err.Error())
	}
	return api.Response(c, err, id)
}

// ChangeStatus 修改
// @Summary 修改
// @Description 修改
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param id path string true "task ID"
// @Param task body model.TodoList true "task"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/{id} [put]
func ChangeStatus(c *fiber.Ctx) error {
	m := new(model.TodoList)
	taskId := c.Params("id")
	if err := c.BodyParser(m); err != nil {
		return api.Response(c, err, nil)
	}
	id, err := m.ChangeStatus(taskId)
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, id)
}

// Remove 删除
// @Summary 删除
// @Description 删除
// @Tags task
// @Security jwt
// @Accept json
// @Produce json
// @Param id path string true "task ID"
// @Success 200 {object} api.ResponseHTTP{data=string}
// @Router /api/task/{id} [delete]
func Remove(c *fiber.Ctx) error {
	m := new(model.TodoList)
	taskId := c.Params("id")
	id, err := m.Remove(taskId)
	if err != nil {
		return api.Response(c, err, nil)
	}
	return api.Response(c, err, id)
}
