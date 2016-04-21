#Lambda
Plugin created for telegraf that goes to the specified Lambda location, grabs the json outputted,
and converts them for telegraf's use to send as metered statistics.

#Json Format
```
[{
	"fields": {
		"count": 1
	},
	"tags": {
		"environment": "dev-tools-jenkins",
		"version": "unknown"
	}
}, {
	"fields": {
		"count": 1
	},
	"tags": {
		"environment": "unknown",
		"version": "1.0.35"
	}
}, {
	"fields": {
		"count": 1
	},
	"tags": {
		"environment": "unknown",
		"version": "1.0.36"
	}
}, {
	"fields": {
		"count": 9
	},
	"tags": {
		"environment": "unknown",
		"version": "unknown"
	}
}]
```