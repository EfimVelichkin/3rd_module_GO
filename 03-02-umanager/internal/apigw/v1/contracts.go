package v1

import (
	"github.com/EfimVelichkin/3rd_module_GO/03-02-umanager/pkg/pb"
)

type usersClient interface {
	pb.UserServiceClient
}

type linksClient interface {
	pb.LinkServiceClient
}
