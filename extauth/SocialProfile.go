package main

import "go.hootops.com/hootsuite/core-rest/client/core"

type SocialProfile struct {
	Auth1                 string
	Auth2                 string
	Avatar                string
	IsReauthRequired      int64
	IsSecurePost          bool
	ExtendedInfo          string
	ExternalID            string
	OrganizationID        uint64
	MemberID              uint64
	SocialNetworkAppAuth1 string
	SocialNetworkAppAuth2 string
	SocialNetworkAppAuth3 string
	SocialNetworkAppID    uint64
	SocialProfileID       uint64
	SocialProfileType     string
	UserID                string
	Username              string
}

func toSocialProfile(response *core.GetSocialProfileOK) SocialProfile {
	return SocialProfile{
		Auth1:                 response.Payload.Auth1,
		Auth2:                 response.Payload.Auth2,
		Avatar:                response.Payload.Avatar,
		IsReauthRequired:      response.Payload.IsReauthRequired,
		IsSecurePost:          response.Payload.IsSecurePost,
		ExtendedInfo:          response.Payload.ExtendedInfo,
		ExternalID:            response.Payload.ExternalID,
		OrganizationID:        response.Payload.OrganizationID,
		MemberID:              response.Payload.MemberID,
		SocialNetworkAppAuth1: response.Payload.SocialNetworkAppAuth1,
		SocialNetworkAppAuth2: response.Payload.SocialNetworkAppAuth2,
		SocialNetworkAppAuth3: response.Payload.SocialNetworkAppAuth3,
		SocialNetworkAppID:    response.Payload.SocialNetworkAppID,
		SocialProfileID:       response.Payload.SocialProfileID,
		SocialProfileType:     response.Payload.SocialProfileType,
		UserID:                response.Payload.UserID,
		Username:              response.Payload.Username,
	}
}