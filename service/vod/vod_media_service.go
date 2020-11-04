// Code generated by protoc-gen-volcengine-sdk
// source: VodMediaService
// DO NOT EDIT!

package vod

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/request"
	"github.com/TTvcloud/vcloud-sdk-golang/models/vod/response"
	"google.golang.org/protobuf/encoding/protojson"
)

/**
 * UpdateVideoInfo.
 *
 * @param input *models.request.UpdateVideoInfoRequest
 * @return *models.response.UpdateVideoInfoResponse, int, error
 * @throws Exception the exception
 */
func (p *Vod) UpdateVideoInfo(req *request.UpdateVideoInfoRequest) (*response.UpdateVideoInfoResponse, int, error) {
	query := url.Values{}
	form := url.Values{}
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	jsonData := marshaler.Format(req)
	reqMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &reqMap)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range reqMap {
		var sv string
		switch ov := v.(type) {
		case string:
			sv = ov
		case int:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int8:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint8:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int16:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint16:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int32:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint32:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int64:
			sv = strconv.FormatInt(ov, 10)
		case uint64:
			sv = strconv.FormatUint(ov, 10)
		case bool:
			sv = strconv.FormatBool(ov)
		case float32:
			sv = strconv.FormatFloat(float64(ov), 'f', -1, 32)
		case float64:
			sv = strconv.FormatFloat(ov, 'f', -1, 64)
		case []byte:
			sv = string(ov)
		default:
			v2, e2 := json.Marshal(ov)
			if e2 != nil {
				return nil, 0, e2
			}
			sv = string(v2)
		}
		query.Set(k, sv)
		form.Add(k, sv)
	}

	respBody, status, err := p.Query("UpdateVideoInfo", query)

	if err != nil {
		return nil, status, err
	}
	output := &response.UpdateVideoInfoResponse{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = unmarshaler.Unmarshal(respBody, output)
	if err != nil {
		return nil, status, err
	} else {
		return output, status, nil
	}
}
