package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseMessage_parseV2(t *testing.T) {
	pro, err := newParseMessage(nil)
	assert.Nil(t, err)
	p := pro.(*parseMessage)

	type args struct {
		message string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 map[string]interface{}
	}{
		{
			"1",
			args{message: "2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b,userid=1,orderid=xyz] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start"},
			"2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start",
			map[string]interface{}{
				"request-id": "07409b69-2595-4e38-b895-5846cf1e0d8b",
				"userid":     "1",
				"orderid":    "xyz",
				"level":      "INFO",
			},
		},
		{
			"2",
			args{message: "2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start"},
			"2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start",
			map[string]interface{}{
				"request-id": "07409b69-2595-4e38-b895-5846cf1e0d8b",
				"level":      "INFO",
			},
		},
		{
			"2.5",
			args{message: "2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b,07409b69-2595-4e38-b895-5846cf1e0das] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start"},
			"2018-11-22 11:02:35.541 INFO [pmp,07409b69-2595-4e38-b895-5846cf1e0d8b,07409b69-2595-4e38-b895-5846cf1e0das] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start",
			map[string]interface{}{
				"request-id": "07409b69-2595-4e38-b895-5846cf1e0d8b",
				"level":      "INFO",
			},
		},
		{
			"3",
			args{message: "2018-11-22 11:02:35.541 INFO [07409b69-2595-4e38-b895-5846cf1e0d8b] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start"},
			"2018-11-22 11:02:35.541 INFO [07409b69-2595-4e38-b895-5846cf1e0d8b] - [main] o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start",
			map[string]interface{}{
				"level":      "INFO",
			},
		},
		{
			"4",
			args{message: "2018-11-22 11:02:35.541 INFO xxx - xxx o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start"},
			"2018-11-22 11:02:35.541 INFO xxx - xxx o.s.j.e.a.AnnotationMBeanExporter        : Registering beans for JMX exposure on start",
			map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := p.parseV2(tt.args.message)
			if got != tt.want {
				t.Errorf("parseV2() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseV2() got1 = %v, want %v", got1, tt.want1)
			}

		})
	}
}

var testSampleLog string = "2021-08-16 15:33:03.506 DEBUG [app-fulfillment-runtime,,] - [ConsumeMessageThread_5] API                                     : executeByDB, cost=11ms, gqlOrSql=DataStoreSqlDto{sqls=[SqlParam(sql=select `*` from `oms_FulfilmentOrderBO` where `id` = ?, params=[7007000], isSearchSQL=false, language=null)]}, reqId=8e04dadb-2e71-40f8-b7b8-f72bed365fbb, res={\"requestId\":\"6nib06b47r\",\"res\":{\"selectResult\":[{\"id\":7007000,\"inboundFreezeMark\":false,\"fulfilmentOrderCode\":\"FFO210816000018508\",\"outOrderId\":3079279,\"outOrderCode\":\"TRO20210816000030510\",\"platformOrderCode\":\"3091250\",\"fulfilmentOrderTypeDict\":\"SALES_ORDER\",\"businessTypeDict\":\"NORMAL\",\"object\":{\"address\":{\"districtList\":[{\"districtCode\":\"CODE330000\",\"districtName\":\"浙江省\",\"districtAlias\":\"zhe jiang sheng\",\"districtLevel\":1,\"isLeaf\":false,\"districtStatus\":\"ENABLED\",\"createdAt\":1609430400000,\"deletedAt\":0,\"isDeleted\":false,\"postCode\":\"330000\",\"@id\":\"@id:47bf8ff6-5494-4d9e-9948-ab2ebb9b1992\",\"id\":330000,\"region\":{\"@id\":\"@id:3dc7b216-7d42-4fc4-ad75-ace36cb69bdf\",\"id\":100000},\"_version\":1,\"updatedAt\":1609430400000},{\"districtCode\":\"CODE330100\",\"districtName\":\"杭州市\",\"districtAlias\":\"hang zhou shi\",\"districtLevel\":2,\"isLeaf\":false,\"districtStatus\":\"ENABLED\",\"parentDistrict\":{\"@id\":\"@id:24fea132-8987-43f7-b147-c51dfaf74e61\",\"id\":330000},\"createdAt\":1609430400000,\"deletedAt\":0,\"isDeleted\":false,\"postCode\":\"330100\",\"@id\":\"@id:3917e0b4-90a9-41d4-bc32-e3a5f126e3ae\",\"id\":330100,\"region\":{\"@id\":\"@id:979e0376-1144-411a-a46d-94dd08a91cfe\",\"id\":100000},\"_version\":1,\"updatedAt\":1609430400000},{\"districtCode\":\"CODE330106\",\"districtName\":\"西湖区\",\"districtAlias\":\"xi hu qu\",\"districtLevel\":3,\"isLeaf\":false,\"districtStatus\":\"ENABLED\",\"parentDistrict\":{\"@id\":\"@id:fae9781e-b772-417c-83b3-7458dae97e99\",\"id\":330100},\"createdAt\":1609430400000,\"deletedAt\":0,\"isDeleted\":false,\"postCode\":\"330106\",\"@id\":\"@id:941019f7-0891-4a78-9575-6be645de0c87\",\"id\":330106,\"region\":{\"@id\":\"@id:663ed30f-6caa-408a-bf25-343ef231ca42\",\"id\":100000},\"_version\":1,\"updatedAt\":1609430400000}],\"personName\":\"收货人名称模糊处理\",\"personPhone\":\"13514358902\",\"isDefault\":false,\"person\":{\"personName\":\"收货人名称模糊处理\",\"personPhone\":\"13514358902\",\"personStatus\":\"ENABLED\",\"createdAt\":1629099084000,\"deletedAt\":0,\"isDeleted\":false,\"@id\":\"@id:c01119b8-6c8e-490c-ab5b-06a49a97390c\",\"id\":2263272,\"_version\":1,\"entity\":\"@id:5cbfc456-b204-4b8d-83a2-5b959978cd4a\",\"updatedAt\":1629099084000},\"addressType\":\"PERSON\",\"detailAddress\":\"***模糊化详细地址***\",\"@id\":\"@id:ae17433b-a4a2-409c-a11b-0571f451f360\",\"region\":{\"regionName\":\"中国\",\"@id\":\"@id:1e530730-f52a-42ab-9dc9-35c430c45f1e\"},\"addressStatus\":\"ENABLED\"},\"subject\":{\"@id\":\"@id:5cbfc456-b204-4b8d-83a2-5b959978cd4a\",\"id\":20211707280},\"@id\":\"@id:09e8a63c-3f12-4382-8e6b-4717b2222d66\"},\"owner\":{\"subject\":{\"fulfillmentWay\":\"DOCKING_FULFILLMENT\",\"entityCode\":\"ent202106150004001\",\"entityType\":\"COMPANY\",\"createdAt\":1623766757000,\"deletedAt\":0,\"isDeleted\":false,\"createdBy\":{\"@id\":\"@id:9b4d4fd9-6ed7-45c2-89ca-0f5863095432\",\"id\":6001},\"entityStatus\":\"ENABLED\",\"entityName\":\"彼悦（北京）科技有限公司\",\"entityCategory\":\"RETAIL\",\"entityShortName\":\"ubras\",\"currency\":{\"@id\":\"@id:6cae7ca1-e4e0-42e0-b75f-bffdaf419781\",\"id\":100001},\"company\":{\"@id\":\"@id:27e7b2d0-5034-4a4b-ac9b-0653a2a500e2\",\"id\":820009},\"@id\":\"@id:c6091154-c223-4e72-a12b-8a1d2eb2cd47\",\"id\":20210434002,\"region\":{\"@id\":\"@id:e1454548-925a-4022-ac07-9825dff1cd78\",\"id\":100000},\"_version\":1,\"user\":{\"@id\":\"@id:72ab5b02-5d4e-44e8-9bfd-48c157f03ec3\",\"id\":538044},\"updatedAt\":1624848588000},\"@id\":\"@id:1d6293d6-5271-468b-aed1-929917313c9c\"},\"ownerId\":20210434002,\"channel\":{\"bizType\":[\"SALES\"],\"channelType\":\"ONLINE\",\"channelStatus\":\"ENABLED\",\"createdAt\":1626940003000,\"deletedAt\":0,\"isDeleted\":false,\"createdBy\":{\"@id\":\"@id:f541b5ff-268a-4b27-b763-19e83e6f725a\",\"id\":8888},\"channelName\":\"苏宁易购\",\"@id\":\"@id:4d499da5-e68b-4f15-a04e-5b43723068cb\",\"id\":18,\"_version\":1,\"entity\":{\"@id\":\"@id:d8f59408-7a59-4096-85c1-2424a62beed3\",\"id\":20210434002},\"updatedAt\":1626940003000,\"channelCode\":\"SUNING\",\"channelShortName\":\"SUNING\",\"channelSource\":\"EXTERNAL\"},\"channelCode\":\"SUNING\",\"shop\":{\"siteStatus\":\"ENABLED\",\"channel\":{\"bizType\":[\"SALES\"],\"channelType\":\"ONLINE\",\"channelStatus\":\"ENABLED\",\"createdAt\":1626940003000,\"deletedAt\":0,\"isDeleted\":false,\"createdBy\":{\"@id\":\"@id:bb25a541-9a53-466a-9c99-363c7ac69b45\",\"id\":8888},\"channelName\":\"苏宁易购\",\"@id\":\"@id:43ebdc60-baa6-4e33-8d4f-d032ebcb0101\",\"id\":18,\"_version\":1,\"entity\":{\"@id\":\"@id:5e17aca8-a8ba-446d-8151-ec995c0e9b7f\",\"id\":20210434002},\"updatedAt\":1626940003000,\"channelCode\":\"SUNING\",\"channelShortName\":\"SUNING\",\"channelSource\":\"EXTERNAL\"},\"siteName\":\"苏宁易购\",\"channelType\":\"ONLINE\",\"isLeaf\":true,\"districtList\":[{\"districtCode\":\"CODE110000\",\"districtName\":\"北京\",\"@id\":\"@id:de6e79de-989e-4b42-af7c-37d422304041\",\"id\":110000,\"districtLevel\":1},{\"parentDistrict\":{\"districtName\":\"北京\",\"@id\":\"@id:cc9b8b9b-0c14-4e7a-b0b5-2ea837a03d4f\",\"id\":110000},\"districtCode\":\"CODE110100\",\"districtName\":\"北京市\",\"@id\":\"@id:e3f58b19-dc0d-45bd-85d0-463924982b85\",\"id\":110100,\"districtLevel\":2},{\"parentDistrict\":{\"districtName\":\"北京市\",\"@id\":\"@id:fc61e7df-cb9d-45f4-a792-4c5cc034f53d\",\"id\":110100},\"districtCode\":\"CODE110101\",\"districtName\":\"东城区\",\"@id\":\"@id:eac7ce70-2472-418f-9074-fd1ca4cac7d5\",\"id\":110101,\"districtLevel\":3},{\"parentDistrict\":{\"districtName\":\"东城区\",\"@id\":\"@id:c9f20ce3-99cb-45f9-9d54-bd3a4457e9ed\",\"id\":110101},\"districtCode\":\"CODE11010100100\",\"districtName\":\"东华门街道\",\"@id\":\"@id:e95069c3-efbe-46fb-8b71-eca189ed2847\",\"id\":11010100100,\"districtLevel\":4}],\"createdAt\":1568941833000,\"synchronizeGoodsDict\":\"TURN_OFF\",\"isDeleted\":false,\"siteLevel\":1,\"appKey\":\"e325a7e469aadfddfd160e55b712d0dd\",\"@id\":\"@id:23d28d8c-075f-4725-8eed-9e3f3d073d19\",\"id\":96,\"_version\":2,\"gateWay\":\"https://open.suning.com/api/http/sopRequest\",\"updatedAt\":1628088800000,\"siteType\":\"SHOP\",\"updatedBy\":{\"@id\":\"@id:dae36f36-5a3f-47be-954b-3c18cbbcb0f2\",\"id\":528045},\"siteCode\":\"PT052_1\",\"siteAddress\":\"待完善\",\"isVirtualSite\":false,\"deletedAt\":0,\"createdBy\":{\"@id\":\"@id:12e79b9b-9fc7-4e10-bc12-d2557ccae48a\",\"id\":8888},\"person\":{\"@id\":\"@id:956579ca-0bde-4b68-8676-d67d7f37ea70\",\"id\":1024001},\"companyEntity\":{\"@id\":\"@id:02107ce1-70e2-4526-8e83-f2cedda6c754\",\"id\":20210434002},\"appSecret\":\"3aa847c928617a09208ade6ad629c3d1\",\"siteSort\":\"SHOP_SITE\",\"entity\":{\"@id\":\"@id:fe079550-3ce7-47a4-906e-aed9756cb7cc\",\"id\":20210434002}},\"timeLine\":{\"tradeOrderTime\":1629099149627,\"@id\":\"@id:b15a9f2f-ce26-4278-b0fb-6238d25491e8\"},\"orderStatusDict\":\"TRANSFORMED\",\"fulfilmentStatusDict\":\"WAITING\",\"allotDealTypeDict\":\"NORMAL_DEAL\",\"instructionWayDict\":\"ONE_PASSAGE\",\"exceptionTypeDict\":\"INVENTORY_OCCUPY_FAILED\",\"exceptionReason\":\"Inventory lacking,please check the inventory\",\"receiveParamTO\":{\"allotInstructionWayDict\":\"EXPRESS\",\"@id\":\"@id:018933ed-3e99-44ea-bdb0-714a4c3f0296\"},\"idempotent\":\"038c7b5c5dd363325b4503ff86c6f159\",\"outOrderVersion\":0,\"fulfilmentExt\":{\"noticeOutSystemDict\":\"CREATE_FULFILMENT_PROGRESS\",\"@id\":\"@id:584d589d-e15b-47f2-bd9a-caf32cd93fcd\"},\"tradeOrderGroupNo\":\"single:3083293:f878a135d6783d5513bf00de6e004856\",\"ifEncrypted\":false,\"_version\":2,\"createdAt\":1629099151000,\"updatedAt\":1629099152000,\"isDeleted\":false,\"deletedAt\":0,\"updatedBy\":{\"id\":528045},\"createdBy\":{\"id\":528045},\"paidAmt\":{\"type\":\"global\",\"value\":0}}]},\"success\":true}"

func Benchmark_parseMessage_parseV2(b *testing.B) {
	pro, err := newParseMessage(nil)
	if err != nil {
		b.Fatal(err)
	}
	p := pro.(*parseMessage)

	for i := 0; i < b.N; i++ {
		p.parseV2(testSampleLog)
	}
}

