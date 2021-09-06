// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValuesRequest
type GetDimensionValuesInput struct {
	_ struct{} `type:"structure"`

	// The context for the call to GetDimensionValues. This can be RESERVATIONS
	// or COST_AND_USAGE. The default value is COST_AND_USAGE. If the context is
	// set to RESERVATIONS, the resulting dimension values can be used in the GetReservationUtilization
	// operation. If the context is set to COST_AND_USAGE, the resulting dimension
	// values can be used in the GetCostAndUsage operation.
	//
	// If you set the context to COST_AND_USAGE, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * DATABASE_ENGINE - The Amazon Relational Database Service database. Examples
	//    are Aurora or MySQL.
	//
	//    * INSTANCE_TYPE - The type of Amazon EC2 instance. An example is m4.xlarge.
	//
	//    * LEGAL_ENTITY_NAME - The name of the organization that sells you AWS
	//    services, such as Amazon Web Services.
	//
	//    * LINKED_ACCOUNT - The description in the attribute map that includes
	//    the full name of the member account. The value field contains the AWS
	//    ID of the member account.
	//
	//    * OPERATING_SYSTEM - The operating system. Examples are Windows or Linux.
	//
	//    * OPERATION - The action performed. Examples include RunInstance and CreateBucket.
	//
	//    * PLATFORM - The Amazon EC2 operating system. Examples are Windows or
	//    Linux.
	//
	//    * PURCHASE_TYPE - The reservation type of the purchase to which this usage
	//    is related. Examples include On-Demand Instances and Standard Reserved
	//    Instances.
	//
	//    * SERVICE - The AWS service such as Amazon DynamoDB.
	//
	//    * USAGE_TYPE - The type of usage. An example is DataTransfer-In-Bytes.
	//    The response for the GetDimensionValues operation includes a unit attribute.
	//    Examples include GB and Hrs.
	//
	//    * USAGE_TYPE_GROUP - The grouping of common usage types. An example is
	//    Amazon EC2: CloudWatch – Alarms. The response for this operation includes
	//    a unit attribute.
	//
	//    * RECORD_TYPE - The different types of charges such as RI fees, usage
	//    costs, tax refunds, and credits.
	//
	// If you set the context to RESERVATIONS, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * CACHE_ENGINE - The Amazon ElastiCache operating system. Examples are
	//    Windows or Linux.
	//
	//    * DEPLOYMENT_OPTION - The scope of Amazon Relational Database Service
	//    deployments. Valid values are SingleAZ and MultiAZ.
	//
	//    * INSTANCE_TYPE - The type of Amazon EC2 instance. An example is m4.xlarge.
	//
	//    * LINKED_ACCOUNT - The description in the attribute map that includes
	//    the full name of the member account. The value field contains the AWS
	//    ID of the member account.
	//
	//    * PLATFORM - The Amazon EC2 operating system. Examples are Windows or
	//    Linux.
	//
	//    * REGION - The AWS Region.
	//
	//    * SCOPE (Utilization only) - The scope of a Reserved Instance (RI). Values
	//    are regional or a single Availability Zone.
	//
	//    * TAG (Coverage only) - The tags that are associated with a Reserved Instance
	//    (RI).
	//
	//    * TENANCY - The tenancy of a resource. Examples are shared or dedicated.
	Context Context `type:"string" enum:"true"`

	// The name of the dimension. Each Dimension is available for a different Context.
	// For more information, see Context.
	//
	// Dimension is a required field
	Dimension Dimension `type:"string" required:"true" enum:"true"`

	// The token to retrieve the next set of results. AWS provides the token when
	// the response from a previous call has more results than the maximum page
	// size.
	NextPageToken *string `type:"string"`

	// The value that you want to search the filter values for.
	SearchString *string `type:"string"`

	// The start and end dates for retrieving the dimension values. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// TimePeriod is a required field
	TimePeriod *DateInterval `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDimensionValuesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDimensionValuesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDimensionValuesInput"}
	if len(s.Dimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Dimension"))
	}

	if s.TimePeriod == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimePeriod"))
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			invalidParams.AddNested("TimePeriod", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValuesResponse
type GetDimensionValuesOutput struct {
	_ struct{} `type:"structure"`

	// The filters that you used to filter your request. Some dimensions are available
	// only for a specific context.
	//
	// If you set the context to COST_AND_USAGE, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * DATABASE_ENGINE - The Amazon Relational Database Service database. Examples
	//    are Aurora or MySQL.
	//
	//    * INSTANCE_TYPE - The type of Amazon EC2 instance. An example is m4.xlarge.
	//
	//    * LEGAL_ENTITY_NAME - The name of the organization that sells you AWS
	//    services, such as Amazon Web Services.
	//
	//    * LINKED_ACCOUNT - The description in the attribute map that includes
	//    the full name of the member account. The value field contains the AWS
	//    ID of the member account.
	//
	//    * OPERATING_SYSTEM - The operating system. Examples are Windows or Linux.
	//
	//    * OPERATION - The action performed. Examples include RunInstance and CreateBucket.
	//
	//    * PLATFORM - The Amazon EC2 operating system. Examples are Windows or
	//    Linux.
	//
	//    * PURCHASE_TYPE - The reservation type of the purchase to which this usage
	//    is related. Examples include On-Demand Instances and Standard Reserved
	//    Instances.
	//
	//    * SERVICE - The AWS service such as Amazon DynamoDB.
	//
	//    * USAGE_TYPE - The type of usage. An example is DataTransfer-In-Bytes.
	//    The response for the GetDimensionValues operation includes a unit attribute.
	//    Examples include GB and Hrs.
	//
	//    * USAGE_TYPE_GROUP - The grouping of common usage types. An example is
	//    Amazon EC2: CloudWatch – Alarms. The response for this operation includes
	//    a unit attribute.
	//
	//    * RECORD_TYPE - The different types of charges such as RI fees, usage
	//    costs, tax refunds, and credits.
	//
	// If you set the context to RESERVATIONS, you can use the following dimensions
	// for searching:
	//
	//    * AZ - The Availability Zone. An example is us-east-1a.
	//
	//    * CACHE_ENGINE - The Amazon ElastiCache operating system. Examples are
	//    Windows or Linux.
	//
	//    * DEPLOYMENT_OPTION - The scope of Amazon Relational Database Service
	//    deployments. Valid values are SingleAZ and MultiAZ.
	//
	//    * INSTANCE_TYPE - The type of Amazon EC2 instance. An example is m4.xlarge.
	//
	//    * LINKED_ACCOUNT - The description in the attribute map that includes
	//    the full name of the member account. The value field contains the AWS
	//    ID of the member account.
	//
	//    * PLATFORM - The Amazon EC2 operating system. Examples are Windows or
	//    Linux.
	//
	//    * REGION - The AWS Region.
	//
	//    * SCOPE (Utilization only) - The scope of a Reserved Instance (RI). Values
	//    are regional or a single Availability Zone.
	//
	//    * TAG (Coverage only) - The tags that are associated with a Reserved Instance
	//    (RI).
	//
	//    * TENANCY - The tenancy of a resource. Examples are shared or dedicated.
	//
	// DimensionValues is a required field
	DimensionValues []DimensionValuesWithAttributes `type:"list" required:"true"`

	// The token for the next set of retrievable results. AWS provides the token
	// when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string `type:"string"`

	// The number of results that AWS returned at one time.
	//
	// ReturnSize is a required field
	ReturnSize *int64 `type:"integer" required:"true"`

	// The total number of search results.
	//
	// TotalSize is a required field
	TotalSize *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetDimensionValuesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDimensionValues = "GetDimensionValues"

// GetDimensionValuesRequest returns a request value for making API operation for
// AWS Cost Explorer Service.
//
// Retrieves all available filter values for a specified filter over a period
// of time. You can search the dimension values for an arbitrary string.
//
//    // Example sending a request using GetDimensionValuesRequest.
//    req := client.GetDimensionValuesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ce-2017-10-25/GetDimensionValues
func (c *Client) GetDimensionValuesRequest(input *GetDimensionValuesInput) GetDimensionValuesRequest {
	op := &aws.Operation{
		Name:       opGetDimensionValues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDimensionValuesInput{}
	}

	req := c.newRequest(op, input, &GetDimensionValuesOutput{})
	return GetDimensionValuesRequest{Request: req, Input: input, Copy: c.GetDimensionValuesRequest}
}

// GetDimensionValuesRequest is the request type for the
// GetDimensionValues API operation.
type GetDimensionValuesRequest struct {
	*aws.Request
	Input *GetDimensionValuesInput
	Copy  func(*GetDimensionValuesInput) GetDimensionValuesRequest
}

// Send marshals and sends the GetDimensionValues API request.
func (r GetDimensionValuesRequest) Send(ctx context.Context) (*GetDimensionValuesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDimensionValuesResponse{
		GetDimensionValuesOutput: r.Request.Data.(*GetDimensionValuesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDimensionValuesResponse is the response type for the
// GetDimensionValues API operation.
type GetDimensionValuesResponse struct {
	*GetDimensionValuesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDimensionValues request.
func (r *GetDimensionValuesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
