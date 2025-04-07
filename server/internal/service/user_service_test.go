package service

import (
	"context"
	"testing"
	"time"

	"github.com/brotigen23/goph-keeper/server/internal/model"
	"github.com/brotigen23/goph-keeper/server/internal/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	timeNow := time.Now()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type argsRepo struct {
		login    string
		password string
	}
	type retRepo struct {
		user *model.User
		err  error
	}
	type prepareRepo struct {
		argsCreate argsRepo
		retCreate  retRepo
	}
	repo := mock.NewMockUsers(ctrl)

	service := NewUserService(repo)

	type userCredentials struct {
		login    string
		password string
	}

	type args struct {
		userCredentials userCredentials
		prepareRepo     *prepareRepo
	}
	type want struct {
		user *model.User
		err  error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test OK",
			args: args{
				prepareRepo: &prepareRepo{
					argsCreate: argsRepo{
						login:    "user1",
						password: gomock.Any().String(),
					},
					retCreate: retRepo{
						user: &model.User{
							ID:        1,
							Login:     "user1",
							Password:  gomock.Any().String(),
							CreatedAt: timeNow,
							UpdatedAt: timeNow,
						},
						err: nil,
					},
				},
				userCredentials: userCredentials{
					login:    "user1",
					password: "pass1",
				},
			},
			want: want{
				user: &model.User{
					ID:        1,
					Login:     "user1",
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo.EXPECT().Create(
				context.Background(),
				test.args.prepareRepo.argsCreate.login,
				gomock.Any()).
				Return(test.args.prepareRepo.retCreate.user, test.args.prepareRepo.retCreate.err)

			user, err := service.Create(
				context.Background(),
				test.args.userCredentials.login,
				test.args.userCredentials.password)

			assert.ErrorIs(t, err, test.want.err)
			assert.Equal(t, test.want.user.ID, user.ID)
			assert.Equal(t, test.want.user.Login, user.Login)
			assert.Equal(t, test.want.user.CreatedAt, user.CreatedAt)
			assert.Equal(t, test.want.user.UpdatedAt, user.UpdatedAt)
		})
	}

}
