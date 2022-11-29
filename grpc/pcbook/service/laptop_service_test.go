package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go-learning/grpc/pcbook/pb"
	"go-learning/grpc/pcbook/sample"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func (s *svcTestSuite) TestLaptopService_CreateLaptop() {
	t := s.T()

	t.Parallel()

	tcs := []struct {
		name     string
		buildArg func() *pb.Laptop
		wantCode codes.Code
	}{
		{
			name: "Normal With Uid Case",
			buildArg: func() *pb.Laptop {
				laptop := sample.NewLaptop()
				uid, err := uuid.NewRandom()
				require.NoError(s.T(), err)

				laptop.Uid = uid.String()

				return laptop
			},
			wantCode: codes.OK,
		},
		{
			name: "Normal Without Uid Case",
		},
		{
			name: "Invalid Uid Case",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			req := &pb.CreateLaptopReq{
				Laptop: tc.buildArg(),
			}

			svc := NewLaptopServer()
			rsp, err := svc.CreateLaptop(context.Background(), req)
			if err != nil {
				require.Empty(t, rsp)

				st, ok := status.FromError(err)
				require.True(t, ok)

				require.Equal(t, tc.wantCode, st.Code())
				return
			}

			require.NotEmpty(t, rsp)
			require.NotEmpty(t, rsp.Uid)

		})
	}
}
