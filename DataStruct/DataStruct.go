package datastruct

import "time"

type StaffData struct {
	StaffId      int    `json:"staff_id" gorm:"column:StaffId"`
	DepartmentId string `json:"dpm_id" gorm:"column:DepartmentId"`
	StaffName    string `json:"staff_name" gorm:"column:StaffName"`
	PhoneNumber  string `json:"phone_number" gorm:"column:PhoneNumber"`
	Address      string `json:"add" gorm:"column:Address"`
}

type TimeKeepping struct {
	StaffId  int        `json:"staff_id" gorm:"column:StaffId"`
	CheckIn  *time.Time `json:"checkin" gorm:"column:CheckIn"`
	CheckOut *time.Time `json:"checkout" gorm:"column:CheckOut"`
}

type Product struct {
	ProductId   string  `json:"productid" gorm:"column:ProductId"`
	ProductName string  `json:"productname" gorm:"column:ProductName"`
	Price       float32 `json:"price" gorm:"column:Price"`
	Quantity    int     `json:"quantity" gorm:"column:Quantity"`
}

type ImportProduct struct {
	ProductId  string     `json:"productid" gorm:"column:ProductId"`
	ImportTime *time.Time `json:"time" gorm:"column:ImportTime"`
	Quantity   int        `json:"quantity" gorm:"column:Quantity"`
}
type SoldList struct {
	BillId    string     `json:"billid" gorm:"column:BillId"`
	ProductId string     `json:"productid" gorm:"column:ProductId"`
	StaffId   int        `json:"staffid" gorm:"column:StaffId"`
	SoldTime  *time.Time `json:"soldtime" gorm:"column:SoldTime"`
	Quantity  int        `json:"quantity" gorm:"column:Quantity"`
}
type BillCalculation struct {
	ProductId   string  `json:"productid" gorm:"column:ProductId"`
	ProductName string  `json:"productname" gorm:"column:ProductName"`
	Price       float32 `json:"price" gorm:"column:Price"`
	Quantity    int     `json:"quantity" gorm:"column:Quantity"`
}
