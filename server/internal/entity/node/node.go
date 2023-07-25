package node

import (
	"gorm.io/gorm"
)

type Node struct {
	SpaceId string `gorm:"index;size:32;not null"`

	ParentId string `gorm:"index;size:32;not null"`

	PreNodeId string `gorm:"index;size:32;not null"`

	NodeId string `gorm:"uniqueIndex;size:32;not null"`

	Name string `gorm:"size:32;not null"`

	Icon string `gorm:"size:32;not null"`

	Type Type

	gorm.Model
}
