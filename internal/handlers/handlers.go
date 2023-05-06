package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"tgBot/internal/commander"
	"tgBot/internal/storage"

	"github.com/pkg/errors"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	deleteCmd = "delete"
	editCmd   = "edit"
)

var BadArgument = errors.New("bad argument")

func listFunc(s string) string {
	data := storage.List()
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(s string) string {
	return "/help - list command\n" +
		"/list - list data\n" +
		"/add - add data\n" +
		"/delete - delete data\n"
}

func addFunc(data string) string {
	params := strings.Split(data, " ")
	if len(params) != 2 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}
	u, err := storage.NewUser(params[0], params[1])
	if err != nil {
		return err.Error()
	}
	err = storage.Add(u)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("user %v added", u)
}

func deleteFunc(data string) string {
	id, err := strconv.Atoi(data)
	if err != nil {
		return err.Error()
	}
	storage.Delete(uint(id))
	return fmt.Sprintf("user %v deleted", uint(id)) // TODO change output to users name, change input to users name, change error messages
}

// func editFucn(s string) string {

// // }
func AddHandlers(c *commander.Commander) {
	c.RegisterHandler(helpCmd, helpFunc)
	c.RegisterHandler(listCmd, listFunc)
	c.RegisterHandler(addCmd, addFunc)
	c.RegisterHandler(deleteCmd, deleteFunc)
	// c.RegisterHandler(editCmd, editFucn)
}
