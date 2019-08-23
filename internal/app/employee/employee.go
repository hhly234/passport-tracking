package employee

import (
	"context"

	valid "github.com/go-ozzo/ozzo-validation"
	"github.com/hhly234/passport-tracking/internal/pkg/types"
)

//Repository for Employee
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Employee, error)
	FindAll(ctx context.Context) ([]types.Employee, error)
	FindByPassportID(ctx context.Context, PassportID string) (*types.Employee, error)
	FindByDateOfExpiry(ctx context.Context, DateOfExpiry string) ([]types.Employee, error)
	Create(ctx context.Context, emp types.Employee) (string, error)
	Update(ctx context.Context, emp types.Employee) error
	Delete(ctx context.Context, id string) error
}

//PassportEmpAssocService interface
type PassportEmpAssocService interface {
	DeleleEmployeeByID(ctx context.Context, EmployeeID string) error
}

//Service is an employee
type Service struct {
	repo         Repository
	passEmpAssoc PassportEmpAssocService
}

//NewService return a new service Service(repo, passEmpAssoc)
func NewService(repo Repository, passEmpAssoc PassportEmpAssocService) *Service {
	return &Service{
		repo:         repo,
		passEmpAssoc: passEmpAssoc,
	}
}
func (s *Service) GetByID(ctx context.Context, id string) (*types.Employee, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]types.Employee, error) {
	return s.repo.FindAll(ctx)
}
func (s *Service) GetByPassportID(ctx context.Context, passportID string) (*types.Employee, error) {
	return s.repo.FindByPassportID(ctx, passportID)
}

func (s *Service) GetByDateOfExpiry(ctx context.Context, dateOfExpiry string) ([]types.Employee, error) {
	return s.repo.FindByDateOfExpiry(ctx, dateOfExpiry)
}

func (s *Service) Create(ctx context.Context, emp types.Employee) (string, error) {
	if err := valid.ValidateStruct(&emp,
		valid.Field(emp.ID, valid.Required),
		valid.Field(emp.Name, valid.Required),
		valid.Field(emp.Email, valid.Required),
		valid.Field(emp.Phone, valid.Required),
		valid.Field(emp.PassportID, valid.Required),
		valid.Field(emp.Role, valid.Required),
		valid.Field(emp.Project, valid.Required),
		valid.Field(emp.DeliveryCenter, valid.Required),
		valid.Field(emp.Manager, valid.Required),
		valid.Field(emp.EmailManager, valid.Required),
	); err != nil {
		return "", nil
	}
	return s.repo.Create(ctx, emp)
}

func (s *Service) Update(ctx context.Context, emp types.Employee) error {
	if err := valid.ValidateStruct(&emp,
		valid.Field(emp.ID, valid.Required),
		valid.Field(emp.Name, valid.Required),
		valid.Field(emp.Email, valid.Required),
		valid.Field(emp.Phone, valid.Required),
		valid.Field(emp.PassportID, valid.Required),
		valid.Field(emp.Role, valid.Required),
		valid.Field(emp.Project, valid.Required),
		valid.Field(emp.DeliveryCenter, valid.Required),
		valid.Field(emp.Manager, valid.Required),
		valid.Field(emp.EmailManager, valid.Required),
	); err != nil {
		return nil
	}
	return s.repo.Update(ctx, emp)
}
