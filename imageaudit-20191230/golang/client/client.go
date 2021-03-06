// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ScanTextRequest struct {
	Tasks  []*ScanTextRequestTasks  `json:"Tasks,omitempty" xml:"Tasks,omitempty" require:"true" type:"Repeated"`
	Labels []*ScanTextRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" require:"true" type:"Repeated"`
}

func (s ScanTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequest) GoString() string {
	return s.String()
}

func (s *ScanTextRequest) SetTasks(v []*ScanTextRequestTasks) *ScanTextRequest {
	s.Tasks = v
	return s
}

func (s *ScanTextRequest) SetLabels(v []*ScanTextRequestLabels) *ScanTextRequest {
	s.Labels = v
	return s
}

type ScanTextRequestTasks struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
}

func (s ScanTextRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequestTasks) GoString() string {
	return s.String()
}

func (s *ScanTextRequestTasks) SetContent(v string) *ScanTextRequestTasks {
	s.Content = &v
	return s
}

type ScanTextRequestLabels struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
}

func (s ScanTextRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s ScanTextRequestLabels) GoString() string {
	return s.String()
}

func (s *ScanTextRequestLabels) SetLabel(v string) *ScanTextRequestLabels {
	s.Label = &v
	return s
}

type ScanTextResponse struct {
	RequestId *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ScanTextResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ScanTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponse) GoString() string {
	return s.String()
}

func (s *ScanTextResponse) SetRequestId(v string) *ScanTextResponse {
	s.RequestId = &v
	return s
}

func (s *ScanTextResponse) SetData(v *ScanTextResponseData) *ScanTextResponse {
	s.Data = v
	return s
}

type ScanTextResponseData struct {
	Elements []*ScanTextResponseDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" require:"true" type:"Repeated"`
}

func (s ScanTextResponseData) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseData) GoString() string {
	return s.String()
}

func (s *ScanTextResponseData) SetElements(v []*ScanTextResponseDataElements) *ScanTextResponseData {
	s.Elements = v
	return s
}

type ScanTextResponseDataElements struct {
	TaskId  *string                                `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	Results []*ScanTextResponseDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" require:"true" type:"Repeated"`
}

func (s ScanTextResponseDataElements) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseDataElements) GoString() string {
	return s.String()
}

func (s *ScanTextResponseDataElements) SetTaskId(v string) *ScanTextResponseDataElements {
	s.TaskId = &v
	return s
}

func (s *ScanTextResponseDataElements) SetResults(v []*ScanTextResponseDataElementsResults) *ScanTextResponseDataElements {
	s.Results = v
	return s
}

type ScanTextResponseDataElementsResults struct {
	Label      *string                                       `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
	Suggestion *string                                       `json:"Suggestion,omitempty" xml:"Suggestion,omitempty" require:"true"`
	Rate       *float32                                      `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	Details    []*ScanTextResponseDataElementsResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" require:"true" type:"Repeated"`
}

func (s ScanTextResponseDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseDataElementsResults) GoString() string {
	return s.String()
}

func (s *ScanTextResponseDataElementsResults) SetLabel(v string) *ScanTextResponseDataElementsResults {
	s.Label = &v
	return s
}

func (s *ScanTextResponseDataElementsResults) SetSuggestion(v string) *ScanTextResponseDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *ScanTextResponseDataElementsResults) SetRate(v float32) *ScanTextResponseDataElementsResults {
	s.Rate = &v
	return s
}

func (s *ScanTextResponseDataElementsResults) SetDetails(v []*ScanTextResponseDataElementsResultsDetails) *ScanTextResponseDataElementsResults {
	s.Details = v
	return s
}

type ScanTextResponseDataElementsResultsDetails struct {
	Label    *string                                               `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
	Contexts []*ScanTextResponseDataElementsResultsDetailsContexts `json:"Contexts,omitempty" xml:"Contexts,omitempty" require:"true" type:"Repeated"`
}

func (s ScanTextResponseDataElementsResultsDetails) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseDataElementsResultsDetails) GoString() string {
	return s.String()
}

func (s *ScanTextResponseDataElementsResultsDetails) SetLabel(v string) *ScanTextResponseDataElementsResultsDetails {
	s.Label = &v
	return s
}

func (s *ScanTextResponseDataElementsResultsDetails) SetContexts(v []*ScanTextResponseDataElementsResultsDetailsContexts) *ScanTextResponseDataElementsResultsDetails {
	s.Contexts = v
	return s
}

type ScanTextResponseDataElementsResultsDetailsContexts struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty" require:"true"`
}

func (s ScanTextResponseDataElementsResultsDetailsContexts) String() string {
	return tea.Prettify(s)
}

func (s ScanTextResponseDataElementsResultsDetailsContexts) GoString() string {
	return s.String()
}

func (s *ScanTextResponseDataElementsResultsDetailsContexts) SetContext(v string) *ScanTextResponseDataElementsResultsDetailsContexts {
	s.Context = &v
	return s
}

type ScanImageRequest struct {
	Task  []*ScanImageRequestTask `json:"Task,omitempty" xml:"Task,omitempty" require:"true" type:"Repeated"`
	Scene []*string               `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true" type:"Repeated"`
}

func (s ScanImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequest) GoString() string {
	return s.String()
}

func (s *ScanImageRequest) SetTask(v []*ScanImageRequestTask) *ScanImageRequest {
	s.Task = v
	return s
}

func (s *ScanImageRequest) SetScene(v []*string) *ScanImageRequest {
	s.Scene = v
	return s
}

type ScanImageRequestTask struct {
	DataId               *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ImageURL             *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	ImageTimeMillisecond *int64  `json:"ImageTimeMillisecond,omitempty" xml:"ImageTimeMillisecond,omitempty"`
	Interval             *int    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MaxFrames            *int    `json:"MaxFrames,omitempty" xml:"MaxFrames,omitempty"`
}

func (s ScanImageRequestTask) String() string {
	return tea.Prettify(s)
}

func (s ScanImageRequestTask) GoString() string {
	return s.String()
}

func (s *ScanImageRequestTask) SetDataId(v string) *ScanImageRequestTask {
	s.DataId = &v
	return s
}

func (s *ScanImageRequestTask) SetImageURL(v string) *ScanImageRequestTask {
	s.ImageURL = &v
	return s
}

func (s *ScanImageRequestTask) SetImageTimeMillisecond(v int64) *ScanImageRequestTask {
	s.ImageTimeMillisecond = &v
	return s
}

func (s *ScanImageRequestTask) SetInterval(v int) *ScanImageRequestTask {
	s.Interval = &v
	return s
}

func (s *ScanImageRequestTask) SetMaxFrames(v int) *ScanImageRequestTask {
	s.MaxFrames = &v
	return s
}

type ScanImageResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *ScanImageResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s ScanImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponse) GoString() string {
	return s.String()
}

func (s *ScanImageResponse) SetRequestId(v string) *ScanImageResponse {
	s.RequestId = &v
	return s
}

func (s *ScanImageResponse) SetData(v *ScanImageResponseData) *ScanImageResponse {
	s.Data = v
	return s
}

type ScanImageResponseData struct {
	Results []*ScanImageResponseDataResults `json:"Results,omitempty" xml:"Results,omitempty" require:"true" type:"Repeated"`
}

func (s ScanImageResponseData) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseData) GoString() string {
	return s.String()
}

func (s *ScanImageResponseData) SetResults(v []*ScanImageResponseDataResults) *ScanImageResponseData {
	s.Results = v
	return s
}

type ScanImageResponseDataResults struct {
	DataId     *string                                   `json:"DataId,omitempty" xml:"DataId,omitempty" require:"true"`
	TaskId     *string                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
	ImageURL   *string                                   `json:"ImageURL,omitempty" xml:"ImageURL,omitempty" require:"true"`
	SubResults []*ScanImageResponseDataResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" require:"true" type:"Repeated"`
}

func (s ScanImageResponseDataResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResults) SetDataId(v string) *ScanImageResponseDataResults {
	s.DataId = &v
	return s
}

func (s *ScanImageResponseDataResults) SetTaskId(v string) *ScanImageResponseDataResults {
	s.TaskId = &v
	return s
}

func (s *ScanImageResponseDataResults) SetImageURL(v string) *ScanImageResponseDataResults {
	s.ImageURL = &v
	return s
}

func (s *ScanImageResponseDataResults) SetSubResults(v []*ScanImageResponseDataResultsSubResults) *ScanImageResponseDataResults {
	s.SubResults = v
	return s
}

type ScanImageResponseDataResultsSubResults struct {
	Label               *string                                                      `json:"Label,omitempty" xml:"Label,omitempty" require:"true"`
	Suggestion          *string                                                      `json:"Suggestion,omitempty" xml:"Suggestion,omitempty" require:"true"`
	Rate                *float32                                                     `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	Scene               *string                                                      `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	Frames              []*ScanImageResponseDataResultsSubResultsFrames              `json:"Frames,omitempty" xml:"Frames,omitempty" require:"true" type:"Repeated"`
	HintWordsInfoList   []*ScanImageResponseDataResultsSubResultsHintWordsInfoList   `json:"HintWordsInfoList,omitempty" xml:"HintWordsInfoList,omitempty" require:"true" type:"Repeated"`
	ProgramCodeDataList []*ScanImageResponseDataResultsSubResultsProgramCodeDataList `json:"ProgramCodeDataList,omitempty" xml:"ProgramCodeDataList,omitempty" require:"true" type:"Repeated"`
	LogoDataList        []*ScanImageResponseDataResultsSubResultsLogoDataList        `json:"LogoDataList,omitempty" xml:"LogoDataList,omitempty" require:"true" type:"Repeated"`
	SfaceDataList       []*ScanImageResponseDataResultsSubResultsSfaceDataList       `json:"SfaceDataList,omitempty" xml:"SfaceDataList,omitempty" require:"true" type:"Repeated"`
	OCRDataList         []*string                                                    `json:"OCRDataList,omitempty" xml:"OCRDataList,omitempty" require:"true" type:"Repeated"`
}

func (s ScanImageResponseDataResultsSubResults) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResults) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResults) SetLabel(v string) *ScanImageResponseDataResultsSubResults {
	s.Label = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetSuggestion(v string) *ScanImageResponseDataResultsSubResults {
	s.Suggestion = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetRate(v float32) *ScanImageResponseDataResultsSubResults {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetScene(v string) *ScanImageResponseDataResultsSubResults {
	s.Scene = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetFrames(v []*ScanImageResponseDataResultsSubResultsFrames) *ScanImageResponseDataResultsSubResults {
	s.Frames = v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetHintWordsInfoList(v []*ScanImageResponseDataResultsSubResultsHintWordsInfoList) *ScanImageResponseDataResultsSubResults {
	s.HintWordsInfoList = v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetProgramCodeDataList(v []*ScanImageResponseDataResultsSubResultsProgramCodeDataList) *ScanImageResponseDataResultsSubResults {
	s.ProgramCodeDataList = v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetLogoDataList(v []*ScanImageResponseDataResultsSubResultsLogoDataList) *ScanImageResponseDataResultsSubResults {
	s.LogoDataList = v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetSfaceDataList(v []*ScanImageResponseDataResultsSubResultsSfaceDataList) *ScanImageResponseDataResultsSubResults {
	s.SfaceDataList = v
	return s
}

func (s *ScanImageResponseDataResultsSubResults) SetOCRDataList(v []*string) *ScanImageResponseDataResultsSubResults {
	s.OCRDataList = v
	return s
}

type ScanImageResponseDataResultsSubResultsFrames struct {
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	URL  *string  `json:"URL,omitempty" xml:"URL,omitempty" require:"true"`
}

func (s ScanImageResponseDataResultsSubResultsFrames) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsFrames) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsFrames) SetRate(v float32) *ScanImageResponseDataResultsSubResultsFrames {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsFrames) SetURL(v string) *ScanImageResponseDataResultsSubResultsFrames {
	s.URL = &v
	return s
}

type ScanImageResponseDataResultsSubResultsHintWordsInfoList struct {
	Context *string `json:"Context,omitempty" xml:"Context,omitempty" require:"true"`
}

func (s ScanImageResponseDataResultsSubResultsHintWordsInfoList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsHintWordsInfoList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsHintWordsInfoList) SetContext(v string) *ScanImageResponseDataResultsSubResultsHintWordsInfoList {
	s.Context = &v
	return s
}

type ScanImageResponseDataResultsSubResultsProgramCodeDataList struct {
	X      *float32 `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
}

func (s ScanImageResponseDataResultsSubResultsProgramCodeDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsProgramCodeDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsProgramCodeDataList) SetX(v float32) *ScanImageResponseDataResultsSubResultsProgramCodeDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsProgramCodeDataList) SetY(v float32) *ScanImageResponseDataResultsSubResultsProgramCodeDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsProgramCodeDataList) SetWidth(v float32) *ScanImageResponseDataResultsSubResultsProgramCodeDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsProgramCodeDataList) SetHeight(v float32) *ScanImageResponseDataResultsSubResultsProgramCodeDataList {
	s.Height = &v
	return s
}

type ScanImageResponseDataResultsSubResultsLogoDataList struct {
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	X      *float32 `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y      *float32 `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
}

func (s ScanImageResponseDataResultsSubResultsLogoDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsLogoDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetType(v string) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.Type = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetName(v string) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.Name = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetX(v float32) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetY(v float32) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetWidth(v float32) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsLogoDataList) SetHeight(v float32) *ScanImageResponseDataResultsSubResultsLogoDataList {
	s.Height = &v
	return s
}

type ScanImageResponseDataResultsSubResultsSfaceDataList struct {
	X      *float32                                                    `json:"X,omitempty" xml:"X,omitempty" require:"true"`
	Y      *float32                                                    `json:"Y,omitempty" xml:"Y,omitempty" require:"true"`
	Width  *float32                                                    `json:"Width,omitempty" xml:"Width,omitempty" require:"true"`
	Height *float32                                                    `json:"Height,omitempty" xml:"Height,omitempty" require:"true"`
	Faces  []*ScanImageResponseDataResultsSubResultsSfaceDataListFaces `json:"Faces,omitempty" xml:"Faces,omitempty" require:"true" type:"Repeated"`
}

func (s ScanImageResponseDataResultsSubResultsSfaceDataList) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsSfaceDataList) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataList) SetX(v float32) *ScanImageResponseDataResultsSubResultsSfaceDataList {
	s.X = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataList) SetY(v float32) *ScanImageResponseDataResultsSubResultsSfaceDataList {
	s.Y = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataList) SetWidth(v float32) *ScanImageResponseDataResultsSubResultsSfaceDataList {
	s.Width = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataList) SetHeight(v float32) *ScanImageResponseDataResultsSubResultsSfaceDataList {
	s.Height = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataList) SetFaces(v []*ScanImageResponseDataResultsSubResultsSfaceDataListFaces) *ScanImageResponseDataResultsSubResultsSfaceDataList {
	s.Faces = v
	return s
}

type ScanImageResponseDataResultsSubResultsSfaceDataListFaces struct {
	Name *string  `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty" require:"true"`
	Id   *string  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
}

func (s ScanImageResponseDataResultsSubResultsSfaceDataListFaces) String() string {
	return tea.Prettify(s)
}

func (s ScanImageResponseDataResultsSubResultsSfaceDataListFaces) GoString() string {
	return s.String()
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataListFaces) SetName(v string) *ScanImageResponseDataResultsSubResultsSfaceDataListFaces {
	s.Name = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataListFaces) SetRate(v float32) *ScanImageResponseDataResultsSubResultsSfaceDataListFaces {
	s.Rate = &v
	return s
}

func (s *ScanImageResponseDataResultsSubResultsSfaceDataListFaces) SetId(v string) *ScanImageResponseDataResultsSubResultsSfaceDataListFaces {
	s.Id = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imageaudit"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ScanTextWithOptions(request *ScanTextRequest, runtime *util.RuntimeOptions) (_result *ScanTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScanTextResponse{}
	_body, _err := client.DoRequest(tea.String("ScanText"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScanText(request *ScanTextRequest) (_result *ScanTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScanTextResponse{}
	_body, _err := client.ScanTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScanImageWithOptions(request *ScanImageRequest, runtime *util.RuntimeOptions) (_result *ScanImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScanImageResponse{}
	_body, _err := client.DoRequest(tea.String("ScanImage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-12-30"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScanImage(request *ScanImageRequest) (_result *ScanImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScanImageResponse{}
	_body, _err := client.ScanImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
