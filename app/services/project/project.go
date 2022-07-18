package project

import (
	"context"
	"encoding/json"
	"fmt"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/gorm/model"
	"tmpl-go-vercel/app/grpc"

	proto "tmpl-go-vercel/gen/go/api/project/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	grpc.GRPCEndpoints.Register(proto.RegisterProjectServiceServer, &Service{})
	grpc.GRPCAuthList = append(grpc.GRPCAuthList,
		fmt.Sprintf("/%s/List", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Create", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Edit", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/Delete", proto.ProjectService_ServiceDesc.ServiceName),
		fmt.Sprintf("/%s/People", proto.ProjectService_ServiceDesc.ServiceName),
	)
}

type Service struct {
	proto.UnimplementedProjectServiceServer
}

func (s *Service) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	db := global.Db
	projects := []model.Project{}
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	zap.L().Debug("DEBUG")
	zap.S().Debugf("req.name:%s,req.personId:%d", req.Name, req.PersonId)

	if req.Name != nil && *req.Name != "" {
		result = result.Where("name LIKE ?", "%"+*req.Name+"%").Find(&projects)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
	}

	if req.PersonId != nil {
		result = result.Where("person_id = ?", req.PersonId).Find(&projects)
		if result.Error != nil {
			return nil, status.Error(codes.Internal, result.Error.Error())
		}
	}

	jsonData, err := json.Marshal(&projects)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := []*proto.Project{}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.ListResponse{
		Data: data,
	}, nil
}

func (s *Service) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	db := global.Db
	project := &model.Project{
		Name:         req.Name,
		Pin:          req.Pin,
		PersonId:     uint(req.PersonId),
		Organization: req.Organization,
		Description:  req.Description,
	}

	result := db.Create(project)

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	return &proto.CreateResponse{
		Id:           int32(project.ID),
		Name:         project.Name,
		Pin:          project.Pin,
		PersonId:     int32(project.ID),
		Organization: project.Organization,
		Description:  project.Description,
		CreatedAt:    project.CreatedAt.Unix(),
		UpdatedAt:    project.UpdatedAt.Unix(),
	}, nil
}

func (s *Service) Edit(ctx context.Context, req *proto.EditRequest) (*proto.EditResponse, error) {
	return &proto.EditResponse{}, nil
}

func (s *Service) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}

func (s *Service) People(ctx context.Context, req *proto.PeopleRequest) (*proto.PeopleResponse, error) {
	db := global.Db
	people := []model.Person{}
	result := db.Find(&people)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	jsonData, err := json.Marshal(&people)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := []*proto.Person{}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.PeopleResponse{
		Data: data,
	}, nil
}
