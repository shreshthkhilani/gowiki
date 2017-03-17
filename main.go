package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
    "fmt"
    "gowiki/page"
)

func viewHandler(c *gin.Context) {
    title := c.Param("title")
    p, err := page.LoadPage(title)
    if err != nil {
        fmt.Println("viewHandler", err)
        c.Redirect(http.StatusFound, "/edit/" + title)
        return
    }
    c.HTML(http.StatusOK, "view.html", gin.H{
        "Title": p.Title,
        "Body": p.Body,
    })
}

func editHandler(c *gin.Context) {
    title := c.Param("title")
    p, err := page.LoadPage(title)
    if err != nil {
        fmt.Println("editHandler", err)
        c.HTML(http.StatusOK, "edit.html", gin.H{
            "Title": title,
        })
        return
    }
    c.HTML(http.StatusOK, "edit.html", gin.H{
        "Title": p.Title,
        "Body": p.Body,
    })
}

func saveHandler(c *gin.Context) {
    title := c.Param("title")
    body := c.PostForm("body")
    p := &page.Page{Title: title, Body: body}
    err := p.Save()
    if err != nil {
        fmt.Println("saveHandler", err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    c.Redirect(http.StatusFound, "/view/" + title)
}

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/view/:title", viewHandler)
    router.GET("/edit/:title", editHandler)
    router.POST("/save/:title", saveHandler)

    router.Run(":8080")
}