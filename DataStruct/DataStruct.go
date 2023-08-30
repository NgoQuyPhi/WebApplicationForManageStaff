package datastruct

type NewStaffData struct {
	StaffId      int    `json:"staff_id" gorm:"column:StaffId"`
	DepartmentId string `json:"dpm_id" gorm:"column:DepartmentId"`
	StaffName    string `json:"staff_name" gorm:"column:StaffName"`
	PhoneNumber  string `json:"phone_number" gorm:"column:PhoneNumber"`
	Address      string `json:"add" gorm:"column: Address"`
}
