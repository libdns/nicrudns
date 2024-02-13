package nicrudns

func (client *Client) AddTxt(zoneName string, names []string, target string, ttl string) (*Response, error) {
	request := &Request{
		RrList: &RrList{
			Rr: []*RR{},
		},
	}
	for _, name := range names {
		request.RrList.Rr = append(request.RrList.Rr, &RR{
			Name: name,
			Type: `TXT`,
			Ttl:  ttl,
			Txt: &Txt{
				String: target,
			},
		})
	}
	return client.Add(zoneName, request)

}
