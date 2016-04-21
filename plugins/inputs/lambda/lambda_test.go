package lambda

import (
)

// Generates a pointer to a mock Lambda object
func genMockLambda() Lambda{
	return &Lambda{
		Name:   "test-funtion",
		Region: "us-east-1",
	}
}

const invalidJSON = "I don't think this is JSON"

const empty = ""

const validJSON_1Field = `[{
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
}]`

const validJSON_2Fields = `[{
	"fields": {
		"count": 1,
		"name": "one"
	},
	"tags": {
		"environment": "dev-tools-jenkins",
		"version": "unknown"
	}
}, {
	"fields": {
		"count": 1,
		"name": "one"
	},
	"tags": {
		"environment": "unknown",
		"version": "1.0.35"
	}
}, {
	"fields": {
		"count": 1,
		"name": "one"
	},
	"tags": {
		"environment": "unknown",
		"version": "1.0.36"
	}
}, {
	"fields": {
		"count": 9,
		"name": "nine"
	},
	"tags": {
		"environment": "unknown",
		"version": "unknown"
	}
}]`

const validJSON_mixedFields = `[{
	"fields": {
		"count": 1,
		"name": "one"
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
		"count": 1,
		"name": "one"
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
}]`
