package auth

import (
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/instagram"
)

const (
	discordProvider   = "discord"
	githubProvider    = "github"
	googleProvider    = "google"
	facebookProvider  = "facebook"
	instagramProvider = "instagram"
)

func Routes(g *gin.Engine, dbPool *pgxpool.Pool) {
	ep := utils.NewDefaultEnvironmentProvider()

	goth.UseProviders(
		discord.New(ep.GetProviderClientId(discordProvider), ep.GetProviderClientSecret(discordProvider), ep.GetProviderCallbackUrl(discordProvider)),
		github.New(ep.GetProviderClientId(githubProvider), ep.GetProviderClientSecret(githubProvider), ep.GetProviderCallbackUrl(githubProvider)),
		google.New(ep.GetProviderClientId(googleProvider), ep.GetProviderClientSecret(googleProvider), ep.GetProviderCallbackUrl(googleProvider)),
		facebook.New(ep.GetProviderClientId(facebookProvider), ep.GetProviderClientSecret(facebookProvider), ep.GetProviderCallbackUrl(facebookProvider)),
		instagram.New(ep.GetProviderClientId(instagramProvider), ep.GetProviderClientSecret(instagramProvider), ep.GetProviderCallbackUrl(instagramProvider)),
	)

	r := g.Group("/api/v1/auth")

	ah := NewDefaultAuthHandler(dbPool)

	r.GET("login/:provider", ah.LogInUser)
	r.GET("login/:provider/callback", ah.LogInUserCallback)
	r.GET("logout", ah.LogOutUser)
}
