package app

import (
	"context"
	"github.com/ravilushqa/otelgqlgen"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/reearth/reearth-backend/internal/adapter"
	"github.com/reearth/reearth-backend/internal/adapter/gql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const enableDataLoaders = true

func graphqlAPI(
	ec *echo.Echo,
	r *echo.Group,
	conf *ServerConfig,
) {
	playgroundEnabled := conf.Debug || conf.Config.Dev

	if playgroundEnabled {
		r.GET("/graphql", echo.WrapHandler(
			playground.Handler("reearth-backend", "/api/graphql"),
		))
	}

	schema := gql.NewExecutableSchema(gql.Config{
		Resolvers: gql.NewResolver(conf.Debug),
	})

	srv := handler.NewDefaultServer(schema)
	srv.Use(otelgqlgen.Middleware())
	if conf.Config.GraphQL.ComplexityLimit > 0 {
		srv.Use(extension.FixedComplexityLimit(conf.Config.GraphQL.ComplexityLimit))
	}

	if playgroundEnabled {
		srv.Use(extension.Introspection{})
	}

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(30),
	})

	srv.SetErrorPresenter(
		// show more detailed error messgage in debug mode
		func(ctx context.Context, e error) *gqlerror.Error {
			if conf.Debug {
				return gqlerror.ErrorPathf(graphql.GetFieldContext(ctx).Path(), e.Error())
			}
			return graphql.DefaultErrorPresenter(ctx, e)
		},
	)

	r.POST("/graphql", func(c echo.Context) error {
		req := c.Request()
		ctx := req.Context()

		usecases := adapter.Usecases(ctx)
		ctx = gql.AttachUsecases(ctx, usecases, enableDataLoaders)
		c.SetRequest(req.WithContext(ctx))

		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
