/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	Config
}

func (s *Server) NewGRPCServer(useTLS bool) *grpc.Server {
	var gRPCServer *grpc.Server
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile(s.CertFile, s.KeyFile)
		if err != nil {
			zap.S().Fatal(err)
		}
		s.grpcOptions = append(s.grpcOptions, grpc.Creds(creds))
	}
	gRPCServer = grpc.NewServer(s.grpcOptions...)
	s.grpcRegistry.ApplyTo(gRPCServer)
	return gRPCServer
}
