package nicrudns

func (client *Client) AddA(zoneName string, names []string, target string, ttl string) (*Response, error) {
	request := &Request{
		RrList: &RrList{
			Rr: []*RR{},
		},
	}
	tgt := Address(target)
	for _, name := range names {
		request.RrList.Rr = append(request.RrList.Rr, &RR{
			Name: name,
			Type: `A`,
			Ttl:  ttl,
			A:    &tgt,
		})
	}
	return client.Add(zoneName, request)
}
