package storage

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpcopentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
)

// NewStorageServer returns a grpc.Server serving Storage
func NewStorageServer(storage Storage) *grpc.Server {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpcopentracing.UnaryServerInterceptor()))

	s := grpc.NewServer(opts...)

	RegisterRepositoryServer(s, &repositoryServer{storage: storage})
	RegisterBranchServer(s, &branchesServer{storage: storage})
	RegisterCommitServer(s, &commitServer{storage: storage})

	return s
}

type repositoryServer struct {
	storage Storage
}

func (s *repositoryServer) Create(ctx context.Context, req *CreateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.storage.Create(ctx, req.GetOwner(), req.GetName())
}

func (s *repositoryServer) SetDescriptions(ctx context.Context, req *SetDescriptionRequest) (*empty.Empty, error) {
	return &empty.Empty{}, s.storage.SetDescription(ctx, req.GetOwner(), req.GetName(), req.GetDescription())
}

type branchesServer struct {
	storage Storage
}

func (s *branchesServer) List(ctx context.Context, req *BranchesRequest) (*BranchesResponse, error) {
	branches, err := s.storage.Branches(ctx, req.GetOwner(), req.GetName())
	if err != nil {
		return nil, err
	}

	res := &BranchesResponse{}
	for _, b := range branches {
		res.Branch = append(res.Branch, &BranchResponse{
			Name: b.Name,
			Sha1: b.Sha1,
			Type: b.Type,
		})
	}

	return res, nil
}

type commitServer struct {
	storage Storage
}

func (s *commitServer) Get(ctx context.Context, req *CommitRequest) (*CommitResponse, error) {
	c, err := s.storage.Commit(ctx, req.GetOwner(), req.GetName(), req.GetRev())
	if err != nil {
		return nil, err
	}

	return &CommitResponse{
		Hash:           c.Hash,
		Tree:           c.Tree,
		Parent:         c.Parent,
		Message:        c.Message,
		Author:         c.Author.Name,
		AuthorEmail:    c.Author.Email,
		AuthorDate:     c.Author.Date.Unix(),
		Committer:      c.Committer.Name,
		CommitterEmail: c.Committer.Email,
		CommitterDate:  c.Committer.Date.Unix(),
	}, nil
}
