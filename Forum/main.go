package main

import (
	"discussionforum/hello/initializers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string `json:"email"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Postid  int    `json:"postid"`
	Email   string `json:"email"`
}

func init() {
	initializers.InitDB()
}

func main() {
	r := gin.Default()

	// Enable CORS for all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/createUser", createUser)
	r.POST("/createPost", createPost)
	r.GET("/getPost", getPost)
	r.POST("/createComment", createComment)
	r.GET("/getComment", getComment)
	r.POST("/deletePost", deletePost)
	r.POST("/editPost", editPost)

	// Start the server
	r.Run(":4090")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO user (email, username)
		VALUES (?, ?)
	`

	err := initializers.DB.Exec(query, user.Email, user.Username).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func createPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO post (title, content, email)
		VALUES (?, ?, ?)
	`

	err := initializers.DB.Exec(query, post.Title, post.Content, post.Email).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func getPost(c *gin.Context) {
	var posts []struct {
		Post
		Username string `json:"username"`
	}
	err := initializers.DB.
		Table("post").
		Select("post.*, user.username as username").
		Joins("JOIN user ON post.email = user.email").
		Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func createComment(c *gin.Context) {
	var comment Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO comment (postid, content, email)
		VALUES (?, ?, ?)
	`

	err := initializers.DB.Exec(query, comment.Postid, comment.Content, comment.Email).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func getComment(c *gin.Context) {
	var comments []struct {
		Comment
		Username string `json:"username"`
	}
	err := initializers.DB.
		Table("comment").
		Select("comment.*, user.username as username").
		Joins("JOIN user ON comment.email = user.email").
		Find(&comments).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, comments)
	}
}

func deletePost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := initializers.DB.Exec("DELETE FROM post WHERE id = ?", post.Id).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func editPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := initializers.DB.Exec("UPDATE post SET title = ?, content = ? WHERE id = ?", post.Title, post.Content, post.Id).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
