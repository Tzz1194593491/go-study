package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//添加路由
	route1(r)
	route2(r)
	route3(r)
	route4(r)
	route5(r)
	route6(r)
	route7(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func route1(r *gin.Engine) {
	r.GET("route1/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("route1/", func(c *gin.Context) {
		params := c.Params
		fmt.Println(params)
	})
}

func route2(r *gin.Engine) {
	//解析路径参数
	r.GET("route2/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, "hello world"+name+action)
	})
}

func route3(r *gin.Engine) {
	//解析param
	r.GET("route3", func(c *gin.Context) {
		name := c.DefaultQuery("name", "world")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
}

func route4(r *gin.Engine) {
	//解析x-www-form-urlencoded或from-data
	r.POST("route4", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
}

func route5(r *gin.Engine) {
	//上传文件
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	//解析文件
	r.POST("route5", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		//err = c.SaveUploadedFile(file, file.Filename)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		c.String(http.StatusOK, file.Filename)
	})
}

func route6(r *gin.Engine) {
	//上传多个文件
	r.MaxMultipartMemory = 1
	r.POST("route6", func(c *gin.Context) {
		file, err := c.MultipartForm()
		if err != nil {
			c.String(500, err.Error())
			return
		}
		files := file.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("upload err %s", err.Error()))
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})
}

func route7(r *gin.Engine) {
	//routes group是为了管理一些相同的URL
	v1 := r.Group("v1")
	//使用 {} 进行规范书写
	{
		//实际请求路径为localhost:8080/v1/login
		v1.GET("/login", func(c *gin.Context) {})
		//实际请求路径为localhost:8080/v1/logout
		v1.GET("/logout", func(c *gin.Context) {})
	}
}
