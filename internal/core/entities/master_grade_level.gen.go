// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entities

import (
	"time"
)

const TableNameMasterGradeLevel = "master_grade_level"

// MasterGradeLevel mapped from table <master_grade_level>
type MasterGradeLevel struct {
	GradeLevelID     string     `gorm:"column:gradeLevelId;type:uuid;primaryKey;default:public.uuid_generate_v4()" json:"gradeLevelId"`
	NameTh           *string    `gorm:"column:nameTh;type:character varying(255)" json:"nameTh"`
	NameEn           *string    `gorm:"column:nameEn;type:character varying(255)" json:"nameEn"`
	Slug             *string    `gorm:"column:slug;type:character varying(50)" json:"slug"`
	Group            *string    `gorm:"column:group;type:character varying(50)" json:"group"`
	Position         int32      `gorm:"column:position;type:integer;not null" json:"position"`
	GradeLevelAbbrev *string    `gorm:"column:gradeLevelAbbrev;type:character varying(50)" json:"gradeLevelAbbrev"`
	IsActive         bool       `gorm:"column:isActive;type:boolean;not null;default:true" json:"isActive"`
	CreatedAt        time.Time  `gorm:"column:createdAt;type:timestamp with time zone;not null;default:now()" json:"createdAt"`
	CreatedBy        *string    `gorm:"column:createdBy;type:uuid" json:"createdBy"`
	UpdatedAt        *time.Time `gorm:"column:updatedAt;type:timestamp with time zone" json:"updatedAt"`
	UpdatedBy        *string    `gorm:"column:updatedBy;type:uuid" json:"updatedBy"`
	DeletedAt        *time.Time `gorm:"column:deletedAt;type:timestamp with time zone" json:"deletedAt"`
	DeletedBy        *string    `gorm:"column:deletedBy;type:uuid" json:"deletedBy"`
}

// TableName MasterGradeLevel's table name
func (*MasterGradeLevel) TableName() string {
	return TableNameMasterGradeLevel
}
