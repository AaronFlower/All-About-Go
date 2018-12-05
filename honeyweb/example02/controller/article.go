package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/aaronflower/ago/honeyweb/example02/model/article"
	"github.com/aaronflower/honey"
)

// ArticleController defines a Article RESTful resource.
type ArticleController struct {
	honey.Controller
}

// Get returns the articles list
func (c *ArticleController) Get() {
	v := article.GetAll()
	articles, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(c.Ct.ResponseWriter, string(articles))
}

// Post creates a article.
func (c *ArticleController) Post() {
	var a article.Article
	var content, title string
	r, w := c.Ct.Request, c.Ct.ResponseWriter
	title = r.FormValue("title")
	content = r.FormValue("content")
	if len(title) == 0 || len(content) == 0 {
		fmt.Fprintf(w, "Name and content can't be null!")
		return
	}
	a.Title = title
	a.Content = content
	a, err := article.Save(a)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the Age!")
		return
	}
	articleByte, err := json.Marshal(a)

	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the Age!")
		return
	}
	fmt.Fprintf(w, string(articleByte))
}

// Delete deletes a article.
func (c *ArticleController) Delete() {
	r, w := c.Ct.Request, c.Ct.ResponseWriter
	id := r.FormValue("id")
	fmt.Printf("r.Form = %+v\n", r.Form)
	if len(id) == 0 {
		fmt.Fprintf(w, "Please input the id 1")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the id 2")
		return
	}

	err = article.Delete(idInt)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(w, "Please input the id 3")
	}
}
