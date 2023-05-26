package repositories

import (
	"reflect"
	"testing"

	"github.com/if1bonacci/lets-go-chat/internal/models"
)

func TestStoreUser(t *testing.T) {
	type args struct {
		user *models.User
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "store new user",
			args: args{
				user: &models.User{
					Token:    "token",
					Id:       "some-id",
					Password: "pass",
					UserName: "usernbamme",
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if bef := len(users); bef != 0 {
				t.Errorf("userRepository.users = %v, want %v", bef, 0)
			}

			StoreUser(tt.args.user)
			if got := len(users); got != tt.want {
				t.Errorf("userRepository.users = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserByName(t *testing.T) {
	type args struct {
		userName string
	}
	myuser := models.User{
		Token:    "token",
		Id:       "some-id",
		Password: "pass",
		UserName: "username1",
	}
	StoreUser(&myuser)

	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "exist user",
			args: args{
				userName: "username1",
			},
			want: &myuser,
		},
		{
			name: "not exist user",
			args: args{
				userName: "username2",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserByName(tt.args.userName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserByToken(t *testing.T) {
	type args struct {
		token string
	}
	myuser := models.User{
		Token:    "username1",
		Id:       "some-id",
		Password: "pass",
		UserName: "username1",
	}
	StoreUser(&myuser)

	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "exist user",
			args: args{
				token: "username1",
			},
			want: &myuser,
		},
		{
			name: "not exist user",
			args: args{
				token: "username2",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserByToken(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveToken(t *testing.T) {
	type args struct {
		user *models.User
	}

	myuser := models.User{
		Token:    "username1",
		Id:       "some-id",
		Password: "pass",
		UserName: "username1",
	}
	StoreUser(&myuser)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "remove token",
			args: args{
				user: &myuser,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.user.Token != "username1" {
				t.Errorf("User has empty token.")
			}

			RemoveToken(tt.args.user)

			if tt.args.user.Token != "" {
				t.Errorf("Remove tocken doesn't work.")
			}
		})
	}
}

func TestListOfUsers(t *testing.T) {
	myuser := models.User{
		Token:    "username1",
		Id:       "some-id",
		Password: "pass",
		UserName: "username1",
	}
	myuser2 := models.User{
		Token:    "username2",
		Id:       "some-id",
		Password: "pass",
		UserName: "username2",
	}
	StoreUser(&myuser)
	StoreUser(&myuser2)

	tests := []struct {
		name string
		want map[string]*models.User
	}{
		{
			name: "list of users",
			want: users,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListOfUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOfUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		userName string
		password string
	}

	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "create user",
			args: args{
				userName: "username",
				password: "pass",
			},
			want: args{
				userName: "username",
				password: "$2a$04$gBNNTUTQKZallgT7N03ObeNfkKeyVXnxO4rwPbYqTfRE0jCDHLWRS",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.userName, tt.args.password); got.Password == "" || got.Id == "" || !reflect.DeepEqual(got.UserName, tt.want.userName) {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
