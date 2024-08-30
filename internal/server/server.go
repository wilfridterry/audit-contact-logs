package server

import (
	"fmt"
	"net"

	"github.com/wilfridterry/audit-log/pkg/domain"

	"google.golang.org/grpc"
)

type Server struct {
	grpcSrv     *grpc.Server
	auditServer *AuditServer
}

func New(auditServer *AuditServer) *Server {
	return &Server{
		grpc.NewServer(),
		auditServer,
	}
}

func (srv *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServiceServer(srv.grpcSrv, srv.auditServer)

	if err := srv.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
