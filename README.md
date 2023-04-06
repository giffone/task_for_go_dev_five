# task_for_go_dev_five
get rates from national bank

### Run local
```console
go run ./cmd/nbrates/main.go
```

1. API


**Save Request**
```
curl --request GET \
  --url http://localhost:8080/currency/save/23.12.2021
```
**Save Response**
```json
{
	"success": true
}
```
2. API


**Get Request:**
```
curl --request GET \
  --url http://localhost:8080/currency/get/21.12.2021
```
**Get Response:**
```json
[
	{
		"title": "АВСТРАЛИЙСКИЙ ДОЛЛАР",
		"code": "AUD",
		"value": 309.42,
		"date": "2021-12-21T00:00:00Z"
	},
	{
		"title": "АЗЕРБАЙДЖАНСКИЙ МАНАТ",
		"code": "AZN",
		"value": 257.74,
		"date": "2021-12-21T00:00:00Z"
	},
	{
		"title": "АРМЯНСКИЙ ДРАМ",
		"code": "AMD",
		"value": 8.88,
		"date": "2021-12-21T00:00:00Z"
	},
...
]
```