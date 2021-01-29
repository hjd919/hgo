package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gogf/gf/util/gconv"
	"github.com/hjd919/gom"
	"github.com/olivere/elastic/v7"
)

// init index test
func TestEsInitIndex(t *testing.T) {
	esInit()

	err := (&Model{}).InitIndex()
	if err != nil {
		log.Println(err)
	}
}

// aggs test
// Aggregate query
func AggQuery(keyword string) (*elastic.SearchResult, error) {
	index := TableName()
	client := gom.Es()
	agg := elastic.NewDateHistogramAggregation().
		Field("@timestamp").
		TimeZone("Asia/Shanghai").
		MinDocCount(1).
		Interval("1m")

	// 查询一分钟前是否出现关键字keyword
	boolQuery := elastic.NewBoolQuery().
		Filter(elastic.NewRangeQuery("@timestamp").
			Format("strict_date_optional_time").
			Gte(time.Now().Add(time.Minute * -1).Format(time.RFC3339)).
			Lte(time.Now().Format(time.RFC3339))).
		Filter(elastic.NewMultiMatchQuery(keyword).
			Type("best_fields").
			Lenient(true))

	result, err := client.Search().
		Index(index).
		Query(boolQuery).
		Timeout("30000ms").
		IgnoreUnavailable(true).
		Size(500).
		Aggregation("aggs", agg).
		Version(true).
		StoredFields("*").
		Do(context.Background())

	return result, err
}

// add test
func TestEsBatchAdd(t *testing.T) {
	esInit()

	rows := []*Model{}
	for i := 60; i < 80; i++ {
		row := &Model{
			User:    "我叫胡建德" + gconv.String(i),
			Message: "我叫胡建德" + gconv.String(i),
			Content: `最近这段时间，《演员请就位2》这档综艺频上热搜，无论是导师还是演员，都贡献过不少话题，毕竟导师咖位够大，加上请了不少有争议的演员，几乎每一期都能登上热搜榜。郭敬明给表现差演员发S卡引争议、赵薇点评女演员矫情、马苏获S卡感言、孟子义改装被郭敬明怒骂等等，每次上热搜，都能引发大家热烈讨论。



			而令家居君印象最深的，当属点评嘉宾李诚儒老师了，他凭借在节目中的耿直发言也是圈粉无数。郭敬明导演发S卡时，他说自己不喜欢翻手为云覆手为雨的人，被问到对陈凯歌导演电影《无极》的印象时，他更是直言自己没看，因为不喜欢形式大于内容的电影。
			
			
			
			
			
			作为娱乐圈前辈，李诚儒老师怼郭敬明导演时，作为后辈的郭敬明虽然总想据理力争，但也不敢怼回去。但当他说没看过《无极》后，陈凯歌则用“梨园世家的子弟相对比较保守”的话，暗讽他不能接受新鲜事物。
			
			
			
			两人的这次争锋相对也自然迅速爬上热搜，但吃瓜群众更多好奇的是，李诚儒老师不是演员么？怎么敢对陈凯歌导演这么直言不讳呢？
			
			
			其实如果大家了解过李诚儒老师的身家背景后，大致就能明白原因。他在91年的时候，就在北京西单开了1000平的商场，取名“特别特”，而且商场里面还会有模特走秀，后面还有了连锁店。据称，当时商场生意好的时候，一天盈利就能到达50万，而那时的李诚儒老师，就已大牌傍身，西服七八万一套，领带上万一条，袜子100美金一双，还都是进口的，放到今天，也没多少人有这么奢侈了。而且在那时，他的座驾就是奔驰560，外汇仓里有1300多万美金，按照那时的汇率，也有6千多万人民币，而且那时的人民币，可比如今的值钱太多了。` + gconv.String(i),
			Retweets: 200,
		}
		rows = append(rows, row)
	}
	BatchAdd(rows)

	TestEsGetList(t)
}

func TestEsAdd(t *testing.T) {
	esInit()

	row := &Model{
		User: "666",
	}
	Add(row)

	TestEsGetList(t)
}

func TestEsAddById(t *testing.T) {
	id := "3"
	esInit()

	row := &Model{
		User: "hjd23",
	}
	AddById(id, row)

	TestEsGetList(t)
}

// upd test
func TestEsUpd(t *testing.T) {
	id := "3"
	esInit()

	updRow := map[string]interface{}{
		"user":     "3333",
		"retweets": 3,
	}
	UpdById(id, updRow)

	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

// incr test
func TestEsIncrId(t *testing.T) {
	id := "1"
	esInit()
	IncrFieldById(id, "retweets")

	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

// del test
func TestEsDelById(t *testing.T) {
	id := "s0Uv4HUB780-B16Wb4F_"
	esInit()

	DelById(id)

	TestEsGetList(t)
}

func TestEsDelIndexById(t *testing.T) {
	esInit()

	DelIndex()

	TestEsGetList(t)
}

// get test
func TestEsGetById(t *testing.T) {
	id := "1"
	esInit()
	r, _ := GetById(id)
	log.Println(gom.JsonEncode(r))
}

func TestEsGetList(t *testing.T) {
	esInit()

	total, rows, err := GetList(&ListParam{
		PageNum:   1,
		PageSize:  10,
		SortField: "",
		SortOrder: "",
		Keyword:   "李诚儒",
	})
	log.Println(total, gom.JsonEncode(rows), err)
}

func esInit() {
	esConf := &gom.EsConfig{
		Url:        "http://es-cn-n6w1r3anu0006zb5t.public.elasticsearch.aliyuncs.com:9200",
		User:       "elastic",
		Password:   "XZ527shortvideo",
		BulkWorker: 6,
		LogLevel:   1,
	}
	gom.EsInit(esConf)
}

// model
type Model struct {
	Id       string `json:"id"`
	User     string `json:"user"`
	Message  string `json:"message"`
	Content  string `json:"content2"`
	Retweets int    `json:"retweets"`
}

func TableName() string {
	return "tpl-2012"
}

// init index
func (t *Model) InitIndex() error {
	index := TableName()
	client := gom.Es()
	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return err
	}
	if !exists {
		mapping := `
{
	"settings": {
		"number_of_shards": 2,
		"number_of_replicas": 0
	},
	"mappings": {
		"properties": {
			"user": {
				"type": "keyword"
			},
			"message": {
				"type": "text",
				"fielddata": true
			},
			"content": {
				"type": "text",
				"analyzer": "ik_smart",
				"search_analyzer": "ik_smart"
			},
			"retweets": {
				"type": "long"
			},
			"age": {
				"type": "integer"
			},
			"tags": {
				"type": "keyword"
			},
			"location": {
				"type": "geo_point"
			},
			"create_time": {
				"format": "epoch_second",
				"type": "date"
			}
			"suggest_field": {
				"type": "completion"
			}
		}
	}
}
`
		createIndex, err := client.CreateIndex(index).Body(mapping).Do(context.Background())
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			err := fmt.Errorf("IndexInit-!createIndex.Acknowledged")
			return err
		}
	}
	return err
}

// add
func BatchAdd(rows []*Model, ps ...*elastic.BulkProcessor) (err error) {
	index := TableName()

	var p *elastic.BulkProcessor
	if len(ps) == 0 {
		p = gom.BulkProcessor()
		defer p.Flush()
	} else {
		p = ps[0]
	}

	for _, row := range rows {
		r := elastic.NewBulkIndexRequest().Index(index).Doc(row)
		// Add the request r to the processor p
		p.Add(r)
	}
	return
}

func Add(row *Model) (err error) {
	client := gom.Es()
	index := TableName()
	// save
	addResult, err := client.Index().
		Index(index).
		BodyJson(row).
		// Refresh("wait_for"). // 同步添加时需要
		Do(context.Background())
	if err != nil {
		return
	}
	row.Id = addResult.Id
	return
}

func AddById(id string, row *Model) (err error) {
	client := gom.Es()
	index := TableName()
	// save
	_, err = client.Index().
		Index(index).
		Id(id).
		BodyJson(row).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

// update
func UpdById(id string, updRow map[string]interface{}) (err error) {
	client := gom.Es()
	index := TableName()

	// save
	_, err = client.Update().
		Index(index).
		Id(id).
		Doc(updRow).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

func IncrFieldById(id string, field string) (err error) {
	client := gom.Es()
	index := TableName()

	script := elastic.NewScript(fmt.Sprintf("ctx._source.%s += params.num", field)).Param("num", 1)
	_, err = client.Update().Index(index).Id(id).
		Script(script).
		Upsert(map[string]interface{}{field: 0}). // init field
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

// del
func DelById(id string) (err error) {
	client := gom.Es()
	index := TableName()

	_, err = client.Delete().Index(index).Id(id).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DelByQuery(user string) (err error) {
	client := gom.Es()
	index := TableName()

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewTermQuery("user", user))

	_, err = client.DeleteByQuery().Index(index).Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DelIndex() (err error) {
	client := gom.Es()
	index := TableName()

	_, err = client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

// get
func GetById(id string) (row *Model, err error) {
	client := gom.Es()
	index := TableName()

	getResult, err := client.Get().
		Index(index).
		Id(id).
		Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	var r Model
	err = json.Unmarshal(getResult.Source, &r)
	if err != nil {
		return
	}
	// append id
	r.Id = getResult.Id
	row = &r
	return
}

//bool query 条件

type ListParam struct {
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	SortOrder string `json:"sort_order"`
	SortField string `json:"sort_field"`

	Keyword string `json:"keyword"`
}

func (p *ListParam) ToFilter() *gom.EsSearch {
	var search gom.EsSearch

	// match one field
	if len(p.Keyword) != 0 {
		search.ShouldQuery = append(search.ShouldQuery, elastic.NewMatchQuery("content", p.Keyword))
	}

	// match many field
	// if len(p.Keyword) != 0 {
	// 	search.ShouldQuery = append(search.ShouldQuery, elastic.NewMultiMatchQuery(p.Keyword, "user", "tag"))
	// }

	// range
	// if p.LikeCountMax > 0 {
	// 	rangeQuery := elastic.NewRangeQuery("like_count").
	// 		Gte(param.LikeCountMin).Lt(param.LikeCountMax)
	// 	boolQuery.Filter(rangeQuery)
	// }

	if len(p.SortField) != 0 {
		search.Sorters = append(search.Sorters, gom.EsSort(p.SortField, p.SortOrder))
	}

	pageNum, pageSize := gom.Page(p.PageNum, p.PageSize)
	search.From = (pageNum - 1) * pageSize
	search.Size = pageSize
	return &search
}

func GetList(param *ListParam) (total int64, rows []*Model, err error) {
	client := gom.Es()

	// filter
	filter := param.ToFilter()
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(filter.MustQuery...)
	boolQuery.MustNot(filter.MustNotQuery...)
	boolQuery.Should(filter.ShouldQuery...)
	boolQuery.Filter(filter.Filters...)

	// 当should不为空时，保证至少匹配should中的一项
	if len(filter.MustQuery) == 0 && len(filter.MustNotQuery) == 0 && len(filter.ShouldQuery) > 0 {
		boolQuery.MinimumShouldMatch("1")
	}

	index := TableName()
	searchResult, err := client.Search().
		Index(index).
		Query(boolQuery).
		SortBy(filter.Sorters...).
		From(filter.From).Size(filter.Size).
		Do(context.Background()) // execute
	if err != nil {
		log.Println(err)
		return
	}

	total = searchResult.TotalHits()
	if total == 0 {
		fmt.Print("Found no tweets\n")
		return
	}

	rows = make([]*Model, 0)
	for _, hit := range searchResult.Hits.Hits {
		var r Model
		err = json.Unmarshal(hit.Source, &r)
		if err != nil {
			return
		}
		// append id
		r.Id = hit.Id
		rows = append(rows, &r)
	}
	return
}
