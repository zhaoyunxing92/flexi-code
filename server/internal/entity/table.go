package entity

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/datasheet"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/node"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/space"
	"github.com/zhaoyunxing92/flexi-code/server/internal/entity/sys"
)

func Tables() []interface{} {
	return []interface{}{
		&sys.Account{},
		&node.Node{},
		&space.Space{},
		&datasheet.Datasheet{},
		&datasheet.DatasheetField{},
		&datasheet.DatasheetRecord{},
	}
}
