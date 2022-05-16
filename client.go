package nicrudns

import (
	"github.com/pkg/errors"
	"net/http"
)

type IClient interface {
	AddA(zoneName string, names []string, target string, ttl string) (*Response, error)
	AddAAAA(zoneName string, names []string, target string, ttl string) (*Response, error)
	AddCnames(zoneName string, names []string, target string, ttl string) (*Response, error)
	AddMx(zoneName string, names []string, target string, preference string, ttl string) (*Response, error)
	AddTxt(zoneName string, names []string, target string, ttl string) (*Response, error)
	CommitZone(zoneName string) (*Response, error)
	DeleteRecord(zoneName string, id int) (*Response, error)
	DownloadZone(zoneName string) (string, error)
	GetRecords(zoneName string) ([]*RR, error)
	GetARecords(zoneName string, nameFilter string, targetFilter string) ([]*RR, error)
	GetAAAARecords(zoneName string, nameFilter string, targetFilter string) ([]*RR, error)
	GetCnameRecords(zoneName string, nameFilter string, targetFilter string) ([]*RR, error)
	GetMxRecords(zoneName string, nameFilter string, targetFilter string) ([]*RR, error)
	GetTxtRecords(zoneName string, nameFilter string, targetFilter string) ([]*RR, error)
	RollbackZone(zoneName string) (*Response, error)
	GetServices() ([]*Service, error)
}

type Client struct {
	provider     *Provider
	oauth2client *http.Client
}

func NewClient(provider *Provider) IClient {
	return &Client{provider: provider}
}

func (client *Client) Do(r *http.Request) (*http.Response, error) {
	oauth2client, err := client.GetOauth2Client()
	if err != nil {
		return nil, errors.Wrap(err, Oauth2ClientError.Error())
	}
	return oauth2client.Do(r)
}
