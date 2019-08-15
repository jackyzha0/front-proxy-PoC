package main

import (
	httptransport "github.com/go-openapi/runtime/client"
	"go.hootops.com/hootsuite/core-rest/client/core"
	"log"
	"net/url"
	"strconv"
)

const _CORE_SERVICE_URL = "https://skyline-bridge1.hootsuitemedia.com:5041/service/core"

type CoreClient interface {
	GetSocialProfile(params *core.GetSocialProfileParams) (*core.GetSocialProfileOK, error)
}

type SocialProfileService interface {
	GetSocialProfile(socialProfileId int64) (*SocialProfile, error)
}

type socialProfileServiceImpl struct {
	Client *core.Client
	includeAuthData bool
}

func NewSocialProfileService() SocialProfileService {
	coreUrlString := _CORE_SERVICE_URL
	coreUrl, err := url.Parse(coreUrlString)
	includeAuthDataBool := true
	if err != nil {
		log.Fatalf("Invalid URL for Core client %+v", err)
	}
	transport := httptransport.New(coreUrl.Host, coreUrl.Path, []string{coreUrl.Scheme})
	return &socialProfileServiceImpl{Client: core.New(transport, nil), includeAuthData: includeAuthDataBool}
}

func (s *socialProfileServiceImpl) GetSocialProfile(socialProfileId int64) (*SocialProfile, error) {
	params := core.NewGetSocialProfileParams()
	params.SocialProfileID = strconv.FormatInt(socialProfileId, 10)
	params.IncludeSocialNetworkAppAuthData = &s.includeAuthData
	params.IncludeSocialNetworkAuthData = &s.includeAuthData

	response, err := s.Client.GetSocialProfile(params)
	if err != nil {
		return nil, err
	}
	profile := toSocialProfile(response)
	return &profile, nil
}