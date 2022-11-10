package existence_check

var m = map[string]string{
	"Product":         "data-platform-api-product-master-exconf-queue-from",
	"BusinessPartner": "data-platform-api-business-partner-exconf-queue-from",
}

func getQueueNameByCheckContent(content string) string {
	return m[content]
}
