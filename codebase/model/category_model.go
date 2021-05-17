package model

import (
	"fmt"

	"github.com/beego/beego/client/orm"
)

type (
	Itemcategory struct {
		SrNum        int    `orm:"column(srNum);pk"`
		CategoryId   string `orm:"column(categoryId)"`
		CategoryName string `orm:"column(categoryName); unique"`
	}
	Categories struct {
		CategoryList []Itemcategory
	}
)

func init() {
	orm.RegisterModel(new(Itemcategory))
}

func (categories *Categories) ReadAllCategoryData() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("itemcategory").All(&categories.CategoryList)
	return
}

// category name is a unique value
func (categoryInst *Itemcategory) ReadCategoryByName() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	err = o.QueryTable("itemcategory").Filter("categoryName", categoryInst.CategoryName).One(categoryInst)
	return
}

// Fetch category by category id
func (categoryInst *Itemcategory) ReadCategoryByCategoryId() (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("itemcategory").Filter("categoryId", categoryInst.CategoryId).One(&categoryInst)
	return
}

// Fetch the number of items in the `itemcategory` table
func CountCategoryRows() (numOfRows int, err error) {
	o := orm.NewOrm()
	err = o.Raw("select count(*) from itemcategory").QueryRow(&numOfRows)
	return
}

// Insert single record into the item table at once
func (category *Itemcategory) InsertRecordIntoItemCategory() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(category)
	return
}

// Update Item Record
func (category *Itemcategory) UpdateRecord(newCategoryName string) (err error) {
	// orm.Debug = true
	var tempCategory Itemcategory
	tempCategory.CategoryName = category.CategoryName
	o := orm.NewOrm()
	if o.Read(&tempCategory, "categoryName") == nil {
		tempCategory.CategoryName = newCategoryName
		_, err = o.Update(&tempCategory)
	}
	return
}

func (category *Itemcategory) DeleteRecord() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	fmt.Println("Deleting the record with category name: ", category.CategoryName)
	_, err = o.QueryTable("itemcategory").Filter("categoryName", category.CategoryName).Delete()
	return
}
