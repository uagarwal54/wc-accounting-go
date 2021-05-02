package model

import (
	"fmt"

	"github.com/beego/beego/client/orm"
)

type (
	Item struct {
		SrNum            int       `orm:"column(srNum);pk" json: "srnum,omitempty"`
		ItemId           string    `orm:"column(itemId)" json: "itemid,omitempty"`
		ItemName         string    `orm:"column(itemName)" json: "itemname,omitempty"`
		ItemCategory 	 string    `orm:"column(itemCategory)" json: "itemcategory,omitempty"`
	}
)

func init() {
	orm.RegisterModel(new(Item))
}

func (items *Item) ReadAllItemData() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Item").All(&items)
	return
}

// There may be multiple items with the same name
func (item *Item) ReadItemByName() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	var i []*Item
	_, err = o.QueryTable("item").Filter("itemName", item.ItemName).All(&i)
	return
}

// Fetch all the items in a given category
func (items *Item) ReadAllItemsInACategory(categoryName string) (err error) {
	var i []*Item
	o := orm.NewOrm()
	_, err = o.QueryTable("item").Filter("itemCategory", categoryName).All(&i)
	return
}

// Fetch Item by item id
func (item *Item) ReadItemByItemId() (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("item").Filter("itemId", item.ItemId).One(&item)
	return
}

// Insert record(s) into the item table
func (items *Item) InsertIntoItem() (err error){
	o := orm.NewOrm()
	// The 1st param is the number of records to insert in one bulk statement. The 2nd param is models slice.
	_, err = o.InsertMulti(10,*items)
	return
}

// Update Item Record
func (item *Item) UpdateRecord() (err error){
	o := orm.NewOrm()
	if o.Read(item) == nil {
		_, err = o.Update(item)
	}
	return
}

func (item *Item) DeleteRecords() (err error){
	o := orm.NewOrm()
	var errSlice map[string]string
	fmt.Println("Deleting the record with item id: ", item.ItemId)
	if _, err = o.Delete(item); err != nil{
		errSlice[item.ItemId] = err.Error()
	}
	
	return  
}