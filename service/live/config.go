package live

import (
	"math"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live/cdn"
)

const (
	UPDATE_INTERVAL_SECOND = 300
)

type Live struct {
	*base.Client
	allAppInfosCache
	cdnMap map[string]cdn.CDNInterface
}

const (
	playTypeRtmp = "rtmp"
	playTypeFlv  = "flv"
	playTypeHls  = "hls"
	playTypeCmaf = "cmaf"
	playTypeDash = "dash"
)

const (
	tsBits        = 32
	minCountBits  = 8
	reservedBits  = 8
	minPushIDBits = 10

	countBits  = 8
	pushIDBits = minCountBits + reservedBits + minPushIDBits - countBits
)

var pushIDMask = int64(math.Pow(2, float64(pushIDBits))) - 1
var tsMask = int64(math.Pow(2, float64(tsBits))) - 1

type HTTPEnum string

const (
	HTTP  HTTPEnum = "http"
	HTTPS HTTPEnum = "https"
)

func NewInstance() (*Live, error) {
	instance := &Live{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		allAppInfosCache: allAppInfosCache{
			data: &sync.Map{},
		},
		cdnMap: cdn.Init(),
	}
	instance.lastCredentials = instance.Client.ServiceInfo.Credentials

	instance.mustUpdateAllAppInfosCache()
	go instance.autoFlush()
	return instance, nil
}

func NewInstanceWithRegion(region string) (*Live, error) {
	var serviceInfo *base.ServiceInfo
	var ok bool
	if serviceInfo, ok = ServiceInfoMap[region]; !ok {
		panic("Cant find the region, please check it carefully")
	}

	instance := &Live{
		Client: base.NewClient(serviceInfo, ApiInfoList),
		allAppInfosCache: allAppInfosCache{
			data: &sync.Map{},
		},
		cdnMap: cdn.Init(),
	}
	instance.lastCredentials = instance.Client.ServiceInfo.Credentials

	instance.mustUpdateAllAppInfosCache()
	go instance.autoFlush()
	return instance, nil
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "live.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "live"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "live.ap-singapore-1.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "live"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "live.us-east-1.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "live"},
		},
	}

	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "live.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "live"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"CreateStream": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CreateStream"},
				"Version": []string{"2019-10-01"},
			},
		},
		"MGetStreamsPushInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MGetStreamsPushInfo"},
				"Version": []string{"2019-10-01"},
			},
		},
		"MGetStreamsPlayInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"MGetStreamsPlayInfo"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetVODs": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetVODs"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetRecords": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRecords"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetSnapshots": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetSnapshots"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetStreamTimeShiftInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetStreamTimeShiftInfo"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetOnlineUserNum": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetOnlineUserNum"},
				"Version": []string{"2019-10-01"},
			},
		},
		"GetDesensitizedAllAppInfos": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetDesensitizedAllAppInfos"},
				"Version": []string{"2019-10-01"},
			},
		},
	}
)
