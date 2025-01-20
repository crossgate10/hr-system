package document

import "time"

type Document struct {
	ID         uint `gorm:"primaryKey"`
	EmployeeID uint
	FileName   string
	FilePath   string
	UploadedAt time.Time
	UpdatedAt  time.Time
}

type Repository interface {
	CreateDocument(document *Document) error
	GetDocumentByEmployeeID(employeeID uint) ([]Document, error)
	GetAllDocuments() ([]Document, error)
	UpdateDocument(document *Document) error
	DeleteDocument(id uint) error
}

type Service interface {
	UploadDocument(employeeID uint, fileName, filePath string) (*Document, error)
	GetDocumentsForEmployee(employeeID uint) ([]Document, error)
	UpdateDocumentDetails(document *Document) error
	RemoveDocument(id uint) error
}
