package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestValidateClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "joj.com")
	err := client.Update("John Doe Update", "j@j.com")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "joj.com")
	err := client.Update("", "j@j.com")

	assert.Error(t, err, "name is required")
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountWhenItBelongsToAnotherClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	client2, _ := NewClient("Marie Yei", "m@y.com")
	account := NewAccount(client2)

	err := client.AddAccount(account)

	assert.Error(t, err, "account does not belong to this client")

}
