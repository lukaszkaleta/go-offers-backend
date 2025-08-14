package postgres

type TableEntity struct {
	Name string
	Id   int
}

func (te *TableEntity) RelationEntity(relationName string) RelationEntity {
	return te.RelationEntityWithColumnName(relationName, te.Name+"_id")
}

func (te *TableEntity) RelationEntityWithColumnName(relationName string, columnName string) RelationEntity {
	return RelationEntity{
		TableName:  relationName,
		RelationId: te.Id,
		ColumnName: columnName,
	}
}

type RelationEntity struct {
	TableName  string
	RelationId int
	ColumnName string
}
