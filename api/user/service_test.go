package user

import (
	"context"
	"testing"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/api/store/mocks"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateDataUser(t *testing.T){
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	user := &models.User{Name: "name", Email: "email@email.com", Username: "username", Password: "password", TenantID: "tenant"}
	userChangedData := &models.User{Name: "rename", Email: "new@email.com", Username: "username", Password: "password", TenantID: "tenant"}
	userChangedPassword := &models.User{Name: "name", Email: "email@email.com", Username: "username", Password: "newpassword", TenantID: "tenant"}

	//Changed name and Email
	mock.On("UpdateDataUser", ctx, userChangedData.Username, userChangedData.Email, models.Password(userChangedData.Password), models.Tenant(userChangedData.Tenant)).Return(nil).Once()
	err := s.UpdataDataUser(ctx, userChangedData.Username, userChangedData.Email, user.Password, user.Tenant)

	assert.NoError(t, err)
	mock.AssertExpectations(t)
}
