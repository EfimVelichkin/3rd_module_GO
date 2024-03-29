package usergrpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/EfimVelichkin/3rd_module_GO/03-02-umanager/internal/database"
)

type usersRepository interface {
	Create(ctx context.Context, req database.CreateUserReq) (database.User, error)
	FindByID(ctx context.Context, userID uuid.UUID) (database.User, error)
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	FindAll(ctx context.Context) ([]database.User, error)
}
