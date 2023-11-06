package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReactSring(c *gin.Context) {
	c.String(200, "你好世界！")
}
func ReactJSon(c *gin.Context) {
	type info struct {
		Name string `json:"姓名"`
		Age  int    `json:"年龄"`
		Pwd  string `json:"-"` //不渲染Pwd
	}
	I1 := info{
		Name: "甘进", //以json的方式将Name字段换成姓名字段
		Age:  18,   //以json的方式将Age字段换成年龄字段
		Pwd:  "123456",
	}
	c.JSON(200, I1)
	c.JSON(200, gin.H{"name": "cheerio", "age": 19})
	UserMap := map[string]string{
		"text":  "这是map",
		"Uname": "ganjin",
		"Uage":  "18",
	}
	c.JSON(200, UserMap)
	c.JSON(200, gin.H{"text": "这是直接响应json"})
}

// 响应xml
func _xml(x *gin.Context) {
	x.XML(200, gin.H{"name": "cheerio", "age": 19})
}

// 响应yaml
func _yaml(x *gin.Context) {
	x.YAML(200, gin.H{"name": "cheerio", "age": 19})
}

// 响应html
func _html(c *gin.Context) {
	type info struct {
		Name string `json:"姓名"`
		Age  int    `json:"年龄"`
		Pwd  string `json:"-"` //不渲染Pwd
	}
	I1 := info{
		Name: "甘进", //以json的方式将Name字段换成姓名字段
		Age:  18,   //以json的方式将Age字段换成年龄字段
		Pwd:  "123456",
	}
	//c.HTML(200, "index.html", gin.H{"username": "ganjing"}) //加载后台传入的数据到前端
	c.HTML(200, "index.html", gin.H{"username": I1}) //加载后台传入的数据到前端
}

// 重定向
func _redirect(c *gin.Context) {
	c.Redirect(302, "http://www.youtube.com")
	//302是暂时重定向状态码，301是永久
}
func main() {
	r := gin.Default()
	//加载模板目录
	r.LoadHTMLGlob("PersonalExercise/GoWeb/1.Gin2/2_试图/templates/*")

	//加载某个静态文件，路径是以本地项目为基准的
	//浏览器通过访问static/girl.png,然后后台就会把PersonalExercise/GoWeb/1.Gin2/2_试图/static/girl.png的文件展示到前端
	r.StaticFile("static/girl.png", "PersonalExercise/GoWeb/1.Gin2/2_试图/static/girl.png")

	//浏览器访问/s就是访问"PersonalExercise/GoWeb/1.Gin2/2_试图/static/static"且还可以访问更下面的文件，注意前缀不要重复
	r.StaticFS("/s", http.Dir("PersonalExercise/GoWeb/1.Gin2/2_试图/static/static"))

	r.GET("/string", ReactSring)
	r.GET("/j", ReactJSon)
	r.GET("/x", _xml)
	r.GET("/y", _yaml)
	r.GET("/h", _html)
	r.GET("/youtube", _redirect)
	r.Run("0.0.0.0:8080")
}
