package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"trudex/common/logger"
	bitrixHandle "trudex/trud_distributor/internal/routers/api/v1/handle/bitrix"
	"trudex/trud_distributor/internal/routers/api/v1/middleware"
	"trudex/trud_distributor/internal/services"
)

func InitRouter(ctx context.Context, log *logger.Logger, services services.Services) *http.Server {
	r := gin.Default()
	r.Use(
		gin.Recovery(),
		middleware.Logger(log),
		middleware.CorrelationID(),
		middleware.Config(services.ConfigService),
	)

	// add: https://apidocs.bitrix24.ru/api-reference/crm/universal/events/on-crm-dynamic-item-add.html
	// update: https://apidocs.bitrix24.ru/api-reference/crm/universal/events/on-crm-dynamic-item-update.html
	r.POST("/", bitrixHandle.HandleBitrixConsumer(services.RabbitmqService))

	//event=ONCRMDYNAMICITEMUPDATE&event_handler_id=509&data[FIELDS][ID]=23&data[FIELDS][ENTITY_TYPE_ID]=1032&ts=1740948890&auth[domain]=konsultatsiyatest.bitrix24.ru&auth[client_endpoint]=https://konsultatsiyatest.bitrix24.ru/rest/&auth
	//[server_endpoint]=https://oauth.bitrix.info/rest/&auth[member_id]=995c2cbaf3edb993c1d640820f802149&auth[application_token]=0zhkhllfzn5z3dfuulqqerwr3wduyx42

	// Create a server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return srv
}
