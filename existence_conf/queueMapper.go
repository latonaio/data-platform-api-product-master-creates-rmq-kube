package existence_conf

import "data-platform-api-product-master-creates-rmq-kube/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["ProductType"] = ecQNames[0%len(ecQNames)]
	m["BaseUnit"] = ecQNames[1%len(ecQNames)]
	m["WeightUnit"] = ecQNames[1%len(ecQNames)]
	m["ProductGroup"] = ecQNames[2%len(ecQNames)]
	m["CountryOfOrigin"] = ecQNames[3%len(ecQNames)]
	m["CountryOfOriginLanguage"] = ecQNames[4%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
