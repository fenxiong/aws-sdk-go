package cloudwatch

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteAlarmsRequest generates a request for the DeleteAlarms operation.
func (c *CloudWatch) DeleteAlarmsRequest(input *DeleteAlarmsInput) (req *aws.Request) {
	if opDeleteAlarms == nil {
		opDeleteAlarms = &aws.Operation{
			Name:       "DeleteAlarms",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteAlarms, input, nil)

	return
}

func (c *CloudWatch) DeleteAlarms(input *DeleteAlarmsInput) (err error) {
	req := c.DeleteAlarmsRequest(input)
	err = req.Send()
	return
}

var opDeleteAlarms *aws.Operation

// DescribeAlarmHistoryRequest generates a request for the DescribeAlarmHistory operation.
func (c *CloudWatch) DescribeAlarmHistoryRequest(input *DescribeAlarmHistoryInput) (req *aws.Request, output *DescribeAlarmHistoryOutput) {
	if opDescribeAlarmHistory == nil {
		opDescribeAlarmHistory = &aws.Operation{
			Name:       "DescribeAlarmHistory",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarmHistory, input, output)
	output = &DescribeAlarmHistoryOutput{}
	req.Data = output
	return
}

func (c *CloudWatch) DescribeAlarmHistory(input *DescribeAlarmHistoryInput) (output *DescribeAlarmHistoryOutput, err error) {
	req, out := c.DescribeAlarmHistoryRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarmHistory *aws.Operation

// DescribeAlarmsRequest generates a request for the DescribeAlarms operation.
func (c *CloudWatch) DescribeAlarmsRequest(input *DescribeAlarmsInput) (req *aws.Request, output *DescribeAlarmsOutput) {
	if opDescribeAlarms == nil {
		opDescribeAlarms = &aws.Operation{
			Name:       "DescribeAlarms",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarms, input, output)
	output = &DescribeAlarmsOutput{}
	req.Data = output
	return
}

func (c *CloudWatch) DescribeAlarms(input *DescribeAlarmsInput) (output *DescribeAlarmsOutput, err error) {
	req, out := c.DescribeAlarmsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarms *aws.Operation

// DescribeAlarmsForMetricRequest generates a request for the DescribeAlarmsForMetric operation.
func (c *CloudWatch) DescribeAlarmsForMetricRequest(input *DescribeAlarmsForMetricInput) (req *aws.Request, output *DescribeAlarmsForMetricOutput) {
	if opDescribeAlarmsForMetric == nil {
		opDescribeAlarmsForMetric = &aws.Operation{
			Name:       "DescribeAlarmsForMetric",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarmsForMetric, input, output)
	output = &DescribeAlarmsForMetricOutput{}
	req.Data = output
	return
}

func (c *CloudWatch) DescribeAlarmsForMetric(input *DescribeAlarmsForMetricInput) (output *DescribeAlarmsForMetricOutput, err error) {
	req, out := c.DescribeAlarmsForMetricRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarmsForMetric *aws.Operation

// DisableAlarmActionsRequest generates a request for the DisableAlarmActions operation.
func (c *CloudWatch) DisableAlarmActionsRequest(input *DisableAlarmActionsInput) (req *aws.Request) {
	if opDisableAlarmActions == nil {
		opDisableAlarmActions = &aws.Operation{
			Name:       "DisableAlarmActions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableAlarmActions, input, nil)

	return
}

func (c *CloudWatch) DisableAlarmActions(input *DisableAlarmActionsInput) (err error) {
	req := c.DisableAlarmActionsRequest(input)
	err = req.Send()
	return
}

var opDisableAlarmActions *aws.Operation

// EnableAlarmActionsRequest generates a request for the EnableAlarmActions operation.
func (c *CloudWatch) EnableAlarmActionsRequest(input *EnableAlarmActionsInput) (req *aws.Request) {
	if opEnableAlarmActions == nil {
		opEnableAlarmActions = &aws.Operation{
			Name:       "EnableAlarmActions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableAlarmActions, input, nil)

	return
}

func (c *CloudWatch) EnableAlarmActions(input *EnableAlarmActionsInput) (err error) {
	req := c.EnableAlarmActionsRequest(input)
	err = req.Send()
	return
}

var opEnableAlarmActions *aws.Operation

// GetMetricStatisticsRequest generates a request for the GetMetricStatistics operation.
func (c *CloudWatch) GetMetricStatisticsRequest(input *GetMetricStatisticsInput) (req *aws.Request, output *GetMetricStatisticsOutput) {
	if opGetMetricStatistics == nil {
		opGetMetricStatistics = &aws.Operation{
			Name:       "GetMetricStatistics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetMetricStatistics, input, output)
	output = &GetMetricStatisticsOutput{}
	req.Data = output
	return
}

func (c *CloudWatch) GetMetricStatistics(input *GetMetricStatisticsInput) (output *GetMetricStatisticsOutput, err error) {
	req, out := c.GetMetricStatisticsRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetMetricStatistics *aws.Operation

// ListMetricsRequest generates a request for the ListMetrics operation.
func (c *CloudWatch) ListMetricsRequest(input *ListMetricsInput) (req *aws.Request, output *ListMetricsOutput) {
	if opListMetrics == nil {
		opListMetrics = &aws.Operation{
			Name:       "ListMetrics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListMetrics, input, output)
	output = &ListMetricsOutput{}
	req.Data = output
	return
}

func (c *CloudWatch) ListMetrics(input *ListMetricsInput) (output *ListMetricsOutput, err error) {
	req, out := c.ListMetricsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListMetrics *aws.Operation

// PutMetricAlarmRequest generates a request for the PutMetricAlarm operation.
func (c *CloudWatch) PutMetricAlarmRequest(input *PutMetricAlarmInput) (req *aws.Request) {
	if opPutMetricAlarm == nil {
		opPutMetricAlarm = &aws.Operation{
			Name:       "PutMetricAlarm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutMetricAlarm, input, nil)

	return
}

func (c *CloudWatch) PutMetricAlarm(input *PutMetricAlarmInput) (err error) {
	req := c.PutMetricAlarmRequest(input)
	err = req.Send()
	return
}

var opPutMetricAlarm *aws.Operation

// PutMetricDataRequest generates a request for the PutMetricData operation.
func (c *CloudWatch) PutMetricDataRequest(input *PutMetricDataInput) (req *aws.Request) {
	if opPutMetricData == nil {
		opPutMetricData = &aws.Operation{
			Name:       "PutMetricData",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutMetricData, input, nil)

	return
}

func (c *CloudWatch) PutMetricData(input *PutMetricDataInput) (err error) {
	req := c.PutMetricDataRequest(input)
	err = req.Send()
	return
}

var opPutMetricData *aws.Operation

// SetAlarmStateRequest generates a request for the SetAlarmState operation.
func (c *CloudWatch) SetAlarmStateRequest(input *SetAlarmStateInput) (req *aws.Request) {
	if opSetAlarmState == nil {
		opSetAlarmState = &aws.Operation{
			Name:       "SetAlarmState",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetAlarmState, input, nil)

	return
}

func (c *CloudWatch) SetAlarmState(input *SetAlarmStateInput) (err error) {
	req := c.SetAlarmStateRequest(input)
	err = req.Send()
	return
}

var opSetAlarmState *aws.Operation

type AlarmHistoryItem struct {
	AlarmName       *string    `type:"string"`
	HistoryData     *string    `type:"string"`
	HistoryItemType *string    `type:"string"`
	HistorySummary  *string    `type:"string"`
	Timestamp       *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataAlarmHistoryItem `json:"-", xml:"-"`
}

type metadataAlarmHistoryItem struct {
	SDKShapeTraits bool `type:"structure"`
}

type Datapoint struct {
	Average     *float64   `type:"double"`
	Maximum     *float64   `type:"double"`
	Minimum     *float64   `type:"double"`
	SampleCount *float64   `type:"double"`
	Sum         *float64   `type:"double"`
	Timestamp   *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	Unit        *string    `type:"string"`

	metadataDatapoint `json:"-", xml:"-"`
}

type metadataDatapoint struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteAlarmsInput struct {
	AlarmNames []*string `type:"list"`

	metadataDeleteAlarmsInput `json:"-", xml:"-"`
}

type metadataDeleteAlarmsInput struct {
	SDKShapeTraits bool `type:"structure" required:"AlarmNames"`
}

type DescribeAlarmHistoryInput struct {
	AlarmName       *string    `type:"string"`
	EndDate         *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	HistoryItemType *string    `type:"string"`
	MaxRecords      *int       `type:"integer"`
	NextToken       *string    `type:"string"`
	StartDate       *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataDescribeAlarmHistoryInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmHistoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAlarmHistoryOutput struct {
	AlarmHistoryItems []*AlarmHistoryItem `type:"list"`
	NextToken         *string             `type:"string"`

	metadataDescribeAlarmHistoryOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmHistoryOutput struct {
	SDKShapeTraits bool `type:"structure" resultWrapper:"DescribeAlarmHistoryResult"`
}

type DescribeAlarmsForMetricInput struct {
	Dimensions []*Dimension `type:"list"`
	MetricName *string      `type:"string"`
	Namespace  *string      `type:"string"`
	Period     *int         `type:"integer"`
	Statistic  *string      `type:"string"`
	Unit       *string      `type:"string"`

	metadataDescribeAlarmsForMetricInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsForMetricInput struct {
	SDKShapeTraits bool `type:"structure" required:"MetricName,Namespace"`
}

type DescribeAlarmsForMetricOutput struct {
	MetricAlarms []*MetricAlarm `type:"list"`

	metadataDescribeAlarmsForMetricOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsForMetricOutput struct {
	SDKShapeTraits bool `type:"structure" resultWrapper:"DescribeAlarmsForMetricResult"`
}

type DescribeAlarmsInput struct {
	ActionPrefix    *string   `type:"string"`
	AlarmNamePrefix *string   `type:"string"`
	AlarmNames      []*string `type:"list"`
	MaxRecords      *int      `type:"integer"`
	NextToken       *string   `type:"string"`
	StateValue      *string   `type:"string"`

	metadataDescribeAlarmsInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAlarmsOutput struct {
	MetricAlarms []*MetricAlarm `type:"list"`
	NextToken    *string        `type:"string"`

	metadataDescribeAlarmsOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsOutput struct {
	SDKShapeTraits bool `type:"structure" resultWrapper:"DescribeAlarmsResult"`
}

type Dimension struct {
	Name  *string `type:"string"`
	Value *string `type:"string"`

	metadataDimension `json:"-", xml:"-"`
}

type metadataDimension struct {
	SDKShapeTraits bool `type:"structure" required:"Name,Value"`
}

type DimensionFilter struct {
	Name  *string `type:"string"`
	Value *string `type:"string"`

	metadataDimensionFilter `json:"-", xml:"-"`
}

type metadataDimensionFilter struct {
	SDKShapeTraits bool `type:"structure" required:"Name"`
}

type DisableAlarmActionsInput struct {
	AlarmNames []*string `type:"list"`

	metadataDisableAlarmActionsInput `json:"-", xml:"-"`
}

type metadataDisableAlarmActionsInput struct {
	SDKShapeTraits bool `type:"structure" required:"AlarmNames"`
}

type EnableAlarmActionsInput struct {
	AlarmNames []*string `type:"list"`

	metadataEnableAlarmActionsInput `json:"-", xml:"-"`
}

type metadataEnableAlarmActionsInput struct {
	SDKShapeTraits bool `type:"structure" required:"AlarmNames"`
}

type GetMetricStatisticsInput struct {
	Dimensions []*Dimension `type:"list"`
	EndTime    *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	MetricName *string      `type:"string"`
	Namespace  *string      `type:"string"`
	Period     *int         `type:"integer"`
	StartTime  *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	Statistics []*string    `type:"list"`
	Unit       *string      `type:"string"`

	metadataGetMetricStatisticsInput `json:"-", xml:"-"`
}

type metadataGetMetricStatisticsInput struct {
	SDKShapeTraits bool `type:"structure" required:"Namespace,MetricName,StartTime,EndTime,Period,Statistics"`
}

type GetMetricStatisticsOutput struct {
	Datapoints []*Datapoint `type:"list"`
	Label      *string      `type:"string"`

	metadataGetMetricStatisticsOutput `json:"-", xml:"-"`
}

type metadataGetMetricStatisticsOutput struct {
	SDKShapeTraits bool `type:"structure" resultWrapper:"GetMetricStatisticsResult"`
}

type InternalServiceFault struct {
	Message *string `type:"string"`

	metadataInternalServiceFault `json:"-", xml:"-"`
}

type metadataInternalServiceFault struct {
	SDKShapeTraits bool `type:"structure"`
}

type InvalidFormatFault struct {
	Message *string `locationName:"message" type:"string"`

	metadataInvalidFormatFault `json:"-", xml:"-"`
}

type metadataInvalidFormatFault struct {
	SDKShapeTraits bool `type:"structure"`
}

type InvalidNextToken struct {
	Message *string `locationName:"message" type:"string"`

	metadataInvalidNextToken `json:"-", xml:"-"`
}

type metadataInvalidNextToken struct {
	SDKShapeTraits bool `type:"structure"`
}

type InvalidParameterCombinationException struct {
	Message *string `locationName:"message" type:"string"`

	metadataInvalidParameterCombinationException `json:"-", xml:"-"`
}

type metadataInvalidParameterCombinationException struct {
	SDKShapeTraits bool `type:"structure"`
}

type InvalidParameterValueException struct {
	Message *string `locationName:"message" type:"string"`

	metadataInvalidParameterValueException `json:"-", xml:"-"`
}

type metadataInvalidParameterValueException struct {
	SDKShapeTraits bool `type:"structure"`
}

type LimitExceededFault struct {
	Message *string `locationName:"message" type:"string"`

	metadataLimitExceededFault `json:"-", xml:"-"`
}

type metadataLimitExceededFault struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListMetricsInput struct {
	Dimensions []*DimensionFilter `type:"list"`
	MetricName *string            `type:"string"`
	Namespace  *string            `type:"string"`
	NextToken  *string            `type:"string"`

	metadataListMetricsInput `json:"-", xml:"-"`
}

type metadataListMetricsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListMetricsOutput struct {
	Metrics   []*Metric `type:"list"`
	NextToken *string   `type:"string"`

	metadataListMetricsOutput `json:"-", xml:"-"`
}

type metadataListMetricsOutput struct {
	SDKShapeTraits bool `type:"structure" resultWrapper:"ListMetricsResult"`
}

type Metric struct {
	Dimensions []*Dimension `type:"list"`
	MetricName *string      `type:"string"`
	Namespace  *string      `type:"string"`

	metadataMetric `json:"-", xml:"-"`
}

type metadataMetric struct {
	SDKShapeTraits bool `type:"structure"`
}

type MetricAlarm struct {
	ActionsEnabled                     *bool        `type:"boolean"`
	AlarmARN                           *string      `locationName:"AlarmArn" type:"string"`
	AlarmActions                       []*string    `type:"list"`
	AlarmConfigurationUpdatedTimestamp *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	AlarmDescription                   *string      `type:"string"`
	AlarmName                          *string      `type:"string"`
	ComparisonOperator                 *string      `type:"string"`
	Dimensions                         []*Dimension `type:"list"`
	EvaluationPeriods                  *int         `type:"integer"`
	InsufficientDataActions            []*string    `type:"list"`
	MetricName                         *string      `type:"string"`
	Namespace                          *string      `type:"string"`
	OKActions                          []*string    `type:"list"`
	Period                             *int         `type:"integer"`
	StateReason                        *string      `type:"string"`
	StateReasonData                    *string      `type:"string"`
	StateUpdatedTimestamp              *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	StateValue                         *string      `type:"string"`
	Statistic                          *string      `type:"string"`
	Threshold                          *float64     `type:"double"`
	Unit                               *string      `type:"string"`

	metadataMetricAlarm `json:"-", xml:"-"`
}

type metadataMetricAlarm struct {
	SDKShapeTraits bool `type:"structure"`
}

type MetricDatum struct {
	Dimensions      []*Dimension  `type:"list"`
	MetricName      *string       `type:"string"`
	StatisticValues *StatisticSet `type:"structure"`
	Timestamp       *time.Time    `type:"timestamp" timestampFormat:"iso8601"`
	Unit            *string       `type:"string"`
	Value           *float64      `type:"double"`

	metadataMetricDatum `json:"-", xml:"-"`
}

type metadataMetricDatum struct {
	SDKShapeTraits bool `type:"structure" required:"MetricName"`
}

type MissingRequiredParameterException struct {
	Message *string `locationName:"message" type:"string"`

	metadataMissingRequiredParameterException `json:"-", xml:"-"`
}

type metadataMissingRequiredParameterException struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricAlarmInput struct {
	ActionsEnabled          *bool        `type:"boolean"`
	AlarmActions            []*string    `type:"list"`
	AlarmDescription        *string      `type:"string"`
	AlarmName               *string      `type:"string"`
	ComparisonOperator      *string      `type:"string"`
	Dimensions              []*Dimension `type:"list"`
	EvaluationPeriods       *int         `type:"integer"`
	InsufficientDataActions []*string    `type:"list"`
	MetricName              *string      `type:"string"`
	Namespace               *string      `type:"string"`
	OKActions               []*string    `type:"list"`
	Period                  *int         `type:"integer"`
	Statistic               *string      `type:"string"`
	Threshold               *float64     `type:"double"`
	Unit                    *string      `type:"string"`

	metadataPutMetricAlarmInput `json:"-", xml:"-"`
}

type metadataPutMetricAlarmInput struct {
	SDKShapeTraits bool `type:"structure" required:"AlarmName,MetricName,Namespace,Statistic,Period,EvaluationPeriods,Threshold,ComparisonOperator"`
}

type PutMetricDataInput struct {
	MetricData []*MetricDatum `type:"list"`
	Namespace  *string        `type:"string"`

	metadataPutMetricDataInput `json:"-", xml:"-"`
}

type metadataPutMetricDataInput struct {
	SDKShapeTraits bool `type:"structure" required:"Namespace,MetricData"`
}

type ResourceNotFound struct {
	Message *string `locationName:"message" type:"string"`

	metadataResourceNotFound `json:"-", xml:"-"`
}

type metadataResourceNotFound struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetAlarmStateInput struct {
	AlarmName       *string `type:"string"`
	StateReason     *string `type:"string"`
	StateReasonData *string `type:"string"`
	StateValue      *string `type:"string"`

	metadataSetAlarmStateInput `json:"-", xml:"-"`
}

type metadataSetAlarmStateInput struct {
	SDKShapeTraits bool `type:"structure" required:"AlarmName,StateValue,StateReason"`
}

type StatisticSet struct {
	Maximum     *float64 `type:"double"`
	Minimum     *float64 `type:"double"`
	SampleCount *float64 `type:"double"`
	Sum         *float64 `type:"double"`

	metadataStatisticSet `json:"-", xml:"-"`
}

type metadataStatisticSet struct {
	SDKShapeTraits bool `type:"structure" required:"SampleCount,Sum,Minimum,Maximum"`
}