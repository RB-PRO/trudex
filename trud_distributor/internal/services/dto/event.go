package dto

type Fields struct {
	ID           string
	EntityTypeID string
}

// Event is used in CRM as a Body message that works in webhook and we get it
//
// an example:
//
//	event=ONCRMDYNAMICITEMUPDATE&event_handler_id=509&data%5BFIELDS%5D%5BID%5D=21&data%5BFIELDS%5D%5BENTITY_TYPE_ID%5D=1032&ts=1739049397&auth%5Bdomain%5D=konsultatsiyatest.bitrix24.ru&auth%5Bclient_endpoint%5D=https%3A%2F%2Fkonsultatsiyatest.bitrix24.ru%2Frest%2F&auth%5Bserver_endpoint%5D=https%3A%2F%2Foauth.bitrix.info%2Frest%2F&auth%5Bmember_id%5D=995c2cbaf3edb993c1d640820f802149&auth%5Bapplication_token%5D=0zhkhllfzn5z3dfuulqqerwr3wduyx42
type Event struct {
	Event          string
	EventHandlerID string
	Fields         Fields
	Timestamp      int64
	Auth           struct {
		Domain           string
		ClientEndpoint   string
		ServerEndpoint   string
		MemberID         string
		ApplicationToken string
	}
}
