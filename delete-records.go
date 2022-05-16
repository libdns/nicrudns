package nicrudns

import (
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func (client *Client) DeleteRecord(zoneName string, id int) (*Response, error) {
	url := fmt.Sprintf(DeleteRecordsUrlPattern, client.provider.DnsServiceName, zoneName, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, RequestError.Error())
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, ResponseError.Error())
	}
	apiResponse := Response{}
	if err := xml.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, errors.Wrap(err, XmlDecodeError.Error())
	}
	if apiResponse.Status != SuccessStatus {
		return nil, errors.Wrap(ApiNonSuccessError, describeError(apiResponse.Errors.Error))
	} else {
		return &apiResponse, nil
	}
}
