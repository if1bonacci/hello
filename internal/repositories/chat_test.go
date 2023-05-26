package repositories

import (
	"reflect"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/if1bonacci/lets-go-chat/internal/models"
)

func TestNewChat(t *testing.T) {
	wants := NewChat()
	tests := []struct {
		name string
		want *chatRepository
	}{
		{
			name: "new chat test",
			want: wants,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatRepository_IsActive(t *testing.T) {
	type fields struct {
		chatUsers map[string]*websocket.Conn
	}
	type args struct {
		token string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "active",
			fields: fields{
				chatUsers: map[string]*websocket.Conn{
					"test_token_1": nil,
				},
			},
			args: args{
				token: "test_token_1",
			},
			want: true,
		},
		{
			name: "inactive",
			fields: fields{
				chatUsers: map[string]*websocket.Conn{
					"test_token_1": nil,
					"test_token_2": nil,
				},
			},
			args: args{
				token: "test_token_3",
			},
			want: false,
		},
		{
			name: "empty",
			fields: fields{
				chatUsers: nil,
			},
			args: args{
				token: "test_token_3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &chatRepository{
				chatUsers: tt.fields.chatUsers,
			}
			if got := rep.IsActive(tt.args.token); got != tt.want {
				t.Errorf("chatRepository.IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatRepository_Add(t *testing.T) {
	type fields struct {
		chatUsers map[string]*websocket.Conn
	}
	type args struct {
		user models.User
		conn *websocket.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "add new connect",
			fields: fields{
				chatUsers: map[string]*websocket.Conn{
					"test_token_1": nil,
					"test_token_2": nil,
				},
			},
			args: args{
				user: models.User{Token: "test_token_3"},
				conn: nil,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &chatRepository{
				chatUsers: tt.fields.chatUsers,
			}
			rep.Add(tt.args.user, tt.args.conn)
			if got := len(rep.chatUsers); got != tt.want {
				t.Errorf("chatRepository.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatRepository_Remove(t *testing.T) {
	type fields struct {
		chatUsers map[string]*websocket.Conn
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "rm connect",
			fields: fields{
				chatUsers: map[string]*websocket.Conn{
					"test_token_1": nil,
					"test_token_2": nil,
					"test_token_3": nil,
				},
			},
			args: args{
				token: "test_token_3",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &chatRepository{
				chatUsers: tt.fields.chatUsers,
			}
			rep.Remove(tt.args.token)
			if got := len(rep.chatUsers); got != tt.want {
				t.Errorf("chatRepository.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chatRepository_List(t *testing.T) {
	type fields struct {
		chatUsers map[string]*websocket.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]*websocket.Conn
	}{
		{
			name: "list connections",
			fields: fields{
				chatUsers: map[string]*websocket.Conn{
					"test_token_1": nil,
					"test_token_2": nil,
					"test_token_3": nil,
				},
			},
			want: map[string]*websocket.Conn{
				"test_token_1": nil,
				"test_token_2": nil,
				"test_token_3": nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &chatRepository{
				chatUsers: tt.fields.chatUsers,
			}
			if got := rep.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chatRepository.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
