package nicrudns

func (client *Client) AddAAAA(zoneName string, names []string, target string, ttl string) (*Response, error) {
	request := &Request{
		RrList: &RrList{
			Rr: []*RR{},
		},
	}
	tgt := Address(target)
	for _, name := range names {
		request.RrList.Rr = append(request.RrList.Rr, &RR{
			Name: name,
			Type: `AAAA`,
			Ttl:  ttl,
			AAAA: &tgt,
		})
	}
	return client.Add(zoneName, request)

}
