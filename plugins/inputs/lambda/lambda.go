package lambda

import (
	"encoding/json"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"time"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	//log "gopkg.in/inconshreveable/log15.v2"
	"fmt"
)

//todo: error handling.
//todo: logging that doesn't break
var jsonOut []map[string]map[string]interface{}

type Lambda struct {
	//name of the lambda function, also the metric reported
	Name string
	//region of the lambda function
	Region string
}

var sampleConfig = `
  ### name of the lambda function
  name = "lambda_metrics"
  ### region the lambda function is in: always us-east-1
  region = "us-east-1"
  `

func (l *Lambda) SampleConfig() string {
	return sampleConfig
}

func (l *Lambda) Description() string {
	return "Read metrics from JSON output of an AWS lambda function"
}

// Gathers data from lambda function
func (l *Lambda) Gather(acc telegraf.Accumulator) error {
	if err := l.gatherMetrics(acc); err != nil {
		//log.Error(err.Error())
		return err
	}
	return nil
}

//sends metrics to telegraf accumulator
func (l *Lambda) gatherMetrics(acc telegraf.Accumulator) error {

	response := l.invokeLambdaFunction()

	now := time.Now().UTC()

	l.parseJSON(response)

	fmt.Println("jsonOut:",jsonOut)
	for _,report := range jsonOut {
		var fields = report["fields"]
		var tags = report["tags"]
		stags := make(map[string]string)
		for key, value := range tags {
			converted, ok := value.(string)
			if ok {
				stags[key] = converted
				if(len(fields)>1){
					fmt.Println("Adding fields and tags to accumulator:",fields,tags,now)
					//log.Info("Adding fields and tags to accumulator:",fields,tags)
					acc.AddFields(l.Name,fields,stags,now)

				} else {
					fmt.Println("Adding field with tags to accumulator:",fields,tags,now)
					//log.Info("Adding field with tags to accumulator:",fields,tags)
					acc.Add(l.Name,fields,stags,now)
				}
			} else {
				// Handle error
			}
		}
	}

	return nil
}

func (l *Lambda) invokeLambdaFunction() (payload []byte){
	//initialize AWS session
	sess := session.New(&aws.Config{
		Region: aws.String(l.Region),
	})

	svc := lambda.New(sess);

	//invoke the lambda function via just the name
	params := &lambda.InvokeInput{
		FunctionName:   aws.String(l.Name),
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		//log.Error(err.Error())
		return
	}

	return resp.Payload
}


//parses json format unique to the BaseAmiVersion Lambda function
func (l *Lambda) parseJSON(b []byte) {

	json.Unmarshal(b, &jsonOut);
}

func init() {
	inputs.Add("lambda", func() telegraf.Input {
		return &Lambda{}})

}



