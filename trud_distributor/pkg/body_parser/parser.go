// TODO:
// Get rid of this package. This package is needed to disassemble the line that comes from CRM and parse it into a certain data format.
//
// Create a new package that will deal exclusively with this task and will apply only to those services that will participate in communication with CRM

package body_parser

import (
	"fmt"
	"net/url"
	"strconv"
	"trudex/trud_distributor/internal/services/dto"
)

func ParseEventData(data []byte) (dto.Event, error) {
	parsedData, err := url.ParseQuery(string(data))
	if err != nil {
		return dto.Event{}, fmt.Errorf("ошибка при декодировании данных: %v", err)
	}

	var eventData dto.Event

	eventData.Event = parsedData.Get("event")
	eventData.EventHandlerID = parsedData.Get("event_handler_id")
	eventData.Fields.ID = parsedData.Get("data[FIELDS][ID]")
	eventData.Fields.EntityTypeID = parsedData.Get("data[FIELDS][ENTITY_TYPE_ID]")

	timestampStr := parsedData.Get("ts")
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return dto.Event{}, fmt.Errorf("ошибка при парсинге timestamp: %v", err)
	}
	eventData.Timestamp = timestamp

	// Заполняем данные аутентификации
	eventData.Auth.Domain = parsedData.Get("auth[domain]")
	eventData.Auth.ClientEndpoint = parsedData.Get("auth[client_endpoint]")
	eventData.Auth.ServerEndpoint = parsedData.Get("auth[server_endpoint]")
	eventData.Auth.MemberID = parsedData.Get("auth[member_id]")
	eventData.Auth.ApplicationToken = parsedData.Get("auth[application_token]")

	return eventData, nil
}
