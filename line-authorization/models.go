package line

import "github.com/uptrace/bun"

type LineMembers struct {
	bun.BaseModel `bun:"table:line_members"`

	ID    int64 `bun:"id,pk,autoincrement"`
	Token string
}
