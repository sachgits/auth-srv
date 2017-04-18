package handler

import (
	"github.com/abaeve/auth-srv/proto"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityQueryHandler struct {
	Client client.Client
}

func (eqh *EntityQueryHandler) GetAlliances(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.AlliancesResponse) error {
	return nil
}

func (eqh *EntityQueryHandler) GetCorporations(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.CorporationsResponse) error {
	return nil
}

func (eqh *EntityQueryHandler) GetCharacters(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.CharactersResponse) error {
	return nil
}
