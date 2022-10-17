package model

//Book 结构体
type Book struct {
	ID      int     //ID
	Title   string  //书名
	Author  string  //作者
	Price   float64 //价格
	Sales   int     //销量
	Stock   int     //库存
	ImgPath string
}
