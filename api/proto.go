package api

import "github.com/tarik02/proxyhub/pb"

func ProxyToPB(proxy Proxy) *pb.Proxy {
	return &pb.Proxy{
		Id:       proxy.ID,
		Version:  proxy.Version,
		Port:     proxy.Port,
		Started:  proxy.Started,
		ClientIp: proxy.ClientIP,
		EgressWhitelist: &pb.EgressWhitelist{
			Item: proxy.EgressWhitelist,
		},
	}
}

func ProxyFromPB(proxy *pb.Proxy) Proxy {
	if proxy == nil {
		return Proxy{}
	}

	res := Proxy{
		ID:       proxy.Id,
		Version:  proxy.Version,
		Port:     proxy.Port,
		Started:  proxy.Started,
		ClientIP: proxy.ClientIp,
	}
	if proxy.EgressWhitelist != nil {
		res.EgressWhitelist = append([]string(nil), proxy.EgressWhitelist.Item...)
	}

	return res
}
