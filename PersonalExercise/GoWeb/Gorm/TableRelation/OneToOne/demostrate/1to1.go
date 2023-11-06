package demostrate

type Student struct {
	StuID int `gorm:"primary_key;AUTO_INCREMENT"`
	Name  string
	Age   int

	//指定外键：
	IID int
}

type StuInfo struct {
	InfoID  int `gorm:"primary_key;AUTO_INCREMENT"`
	Address string
	//关联
	Stu Student `gorm:"ForeignKey:IID;AssociationForeignKey:InfoID"`
}
