package utils

import "github.com/mewben/arrstate/internal/enums"

// GetCollectionName by entityType
func GetCollectionName(entityType string) string {
	switch entityType {
	case enums.EntityBusiness:
		return enums.CollBusinesses
	case enums.EntityProject:
		return enums.CollProjects
	case enums.EntityProperty:
		return enums.CollProperties
	case enums.EntityPerson:
		return enums.CollPeople
	case enums.EntityInvoice:
		return enums.CollInvoices
	default:
		return ""
	}
}
