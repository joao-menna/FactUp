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
	goth.UseProviders(
		discord.New(utils.GetProviderClientId(discordProvider), utils.GetProviderClientSecret(discordProvider), utils.GetProviderCallbackUrl(discordProvider)),
		github.New(utils.GetProviderClientId(githubProvider), utils.GetProviderClientSecret(githubProvider), utils.GetProviderCallbackUrl(githubProvider)),
		google.New(utils.GetProviderClientId(googleProvider), utils.GetProviderClientSecret(googleProvider), utils.GetProviderCallbackUrl(googleProvider)),
		facebook.New(utils.GetProviderClientId(facebookProvider), utils.GetProviderClientSecret(facebookProvider), utils.GetProviderCallbackUrl(facebookProvider)),
		instagram.New(utils.GetProviderClientId(instagramProvider), utils.GetProviderClientSecret(instagramProvider), utils.GetProviderCallbackUrl(instagramProvider)),
	)

	r := g.Group("api/v1/auth")

	r.GET("login/:provider", func(ctx *gin.Context) { loginHandler(ctx, dbPool) })
	r.GET("login/:provider/callback", func(ctx *gin.Context) { loginCallbackHandler(ctx, dbPool) })
	r.GET("logout", logoutHandler)
}
