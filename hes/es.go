package gom

// https://github.com/olivere/elastic/wiki/BulkIndex
// https://github.com/olivere/elastic/wiki/BulkProcessor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/olivere/elastic/v7"
)

var __es *gomEs

// 获取es连接
func Es() *elastic.Client {
	return __es.orm
}

// 获取es批量处理连接
func BulkProcessor() *elastic.BulkProcessor {
	return __es.bulkProcessor
}

// 连接es
func EsInit(conf *EsConfig) {
	__es = newGomEs(ctx, conf)
}

// 获取es排序对象
func EsSort(sortField, sortOrder string) elastic.Sorter {
	sorter := elastic.NewFieldSort(sortField)
	if sortOrder == "desc" {
		sorter = sorter.Desc()
	} else {
		sorter = sorter.Asc()
	}
	return sorter
}

// 打印查询语句
func PrintQuery(src interface{}) {
	fmt.Println("*****")
	data, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// 生成es的id
func ESId(v interface{}) string {
	idmd5, _ := gmd5.Encrypt(v)
	return idmd5[0:20]
}

type EsConfig struct {
	Url        string `yarm:"url"`
	User       string `yarm:"user"`
	Password   string `yarm:"password"`
	LogLevel   int    `yarm:"log_level"`
	BulkWorker int    `yarm:"bulk_worker"`
}
type EsSearch struct {
	MustQuery    []elastic.Query
	MustNotQuery []elastic.Query
	ShouldQuery  []elastic.Query
	Filters      []elastic.Query
	Sorters      []elastic.Sorter
	From         int //分页
	Size         int
}

// -inner
type gomEs struct {
	orm           *elastic.Client
	conf          *EsConfig
	ctx           context.Context
	bulkProcessor *elastic.BulkProcessor
}

func newGomEs(ctx context.Context, conf *EsConfig) *gomEs {
	db := &gomEs{
		ctx:  ctx,
		conf: conf,
	}

	var err error
	clientOpts := []elastic.ClientOptionFunc{
		elastic.SetURL(db.conf.Url),
		elastic.SetSniff(false),
	}

	if db.conf.User != "" {
		clientOpts = append(clientOpts, elastic.SetBasicAuth(db.conf.User, db.conf.Password))
	}

	if db.conf.LogLevel > 0 {
		// 设置错误日志输出
		clientOpts = append(clientOpts, elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
		// 设置info日志输出
		clientOpts = append(clientOpts, elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	}

	client, err := elastic.NewClient(clientOpts...)
	if err != nil {
		panic("create elastic client failed, err:" + err.Error())
	}

	db.orm = client

	db.bulkProcessor, err = client.BulkProcessor().Name("bulk_processor").
		Workers(db.conf.BulkWorker).     // number of workers
		BulkActions(1000).               // commit if # requests >= 1000
		BulkSize(2 << 20).               // commit if size of requests >= 2 MB
		FlushInterval(10 * time.Second). // commit every 30s
		Do(ctx)
	if err != nil {
		panic("create elastic bulkProcessor failed, err:" + err.Error())
	}

	// heath
	go db.ping()

	return db
}

func (db *gomEs) ping() {
	// dur := 60 * time.Second
	// ti := time.NewTimer(dur)
	// ti := time.NewTicker(dur)
	// defer ti.Stop()
	for {
		select {
		case <-db.ctx.Done():
			// 程序退出，断开bulkProcessor
			log.Println("程序退出，断开bulkProcessor")
			db.bulkProcessor.Close()
			return
			// case <-ti.C:
			// 	pingResult, num, err := db.orm.Ping(db.conf.Urls[0]).Do(db.ctx)
			// 	if err != nil {
			// 		log.Println("es ping error:", pingResult, num, err.Error())
			// 	}
			// 	log.Println("es ping:", pingResult, num, err.Error())
		}
	}
}

/*
//判断index是否存在
func (e *gomEs) IndexExists(index string) (bool, error) {
	ok, err := e.orm.IndexExists(index).Do(e.ctx)
	if err != nil {
		return false, err
	}
	return ok, nil
}

//创建index
func (e *gomEs) CreateIndex(index, mapping string) (bool, error) {
	srv := e.orm.CreateIndex(index)
	if mapping != "" {
		srv.BodyString(mapping)
	}
	result, err := srv.Do(e.ctx)
	if err != nil {
		return false, err
	}
	return result.Acknowledged, nil
}

//删除索引
func (e *gomEs) DeleteIndex(index string) (err error) {
	srv, err := e.orm.DeleteIndex(index).Do(e.ctx)
	if err != nil {
		return
	}
	if !srv.Acknowledged {
		return errors.New("delet index err:" + index)
	}
	return nil
}

//单条插入
func (e *gomEs) Insert(index string, data interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := e.orm.Bulk()
	doc := elastic.NewBulkIndexRequest().Index(index).Doc(data)
	bulkRequest = bulkRequest.Add(doc)
	return bulkRequest.Do(e.ctx)
}

//单条插入
func (e *gomEs) InsertWithID(index, id string, data map[string]interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := e.orm.Bulk()
	doc := elastic.NewBulkIndexRequest().Index(index).Id(id).Doc(data)
	bulkRequest = bulkRequest.Add(doc)
	return bulkRequest.Do(e.ctx)
}

//批量插入
func (e *gomEs) BatchInsert(index string, data []interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := e.orm.Bulk()
	for _, v := range data {
		doc := elastic.NewBulkIndexRequest().Index(index).Doc(v)
		bulkRequest = bulkRequest.Add(doc)
	}
	return bulkRequest.Do(e.ctx)
}

//批量插入
func (e *gomEs) BatchInsertWithID(index string, data map[string]interface{}) (*elastic.BulkResponse, error) {
	bulkRequest := e.orm.Bulk()
	for id, v := range data {
		doc := elastic.NewBulkIndexRequest().Index(index).Id(id).Doc(v)
		bulkRequest = bulkRequest.Add(doc)
	}
	return bulkRequest.Do(e.ctx)
}

//获取指定Id 的文档
func (e *gomEs) GetOneByID(index string, id string) ([]byte, error) {
	result, err := e.orm.Get().Index(index).Id(id).Do(e.ctx)
	if err != nil {
		return nil, err
	}
	if !result.Found {
		return nil, errors.New("not find the document")
	}
	source, err := result.Source.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return source, nil
}

//更新doc
func (e *gomEs) Update(index, id string, doc map[string]interface{}) error {
	srv := e.orm.Update().
		Index(index).Id(id).
		Doc(doc).
		DetectNoop(true)
	_, err := srv.Do(e.ctx)
	if err != nil {
		return err
	}
	return nil
}

//term查询
func (e *gomEs) TermQuery(index string, fieldName, fieldValue string, size, page int) (*elastic.SearchResult, error) {
	query := elastic.NewTermQuery(fieldName, fieldValue)
	srv := e.orm.Search().Index(index).Query(query)
	if size > 0 {
		srv.Size(size)
	}
	if page > 0 {
		srv.From((page - 1) * size)
	}
	searchResult, err := srv.Pretty(true).Do(e.ctx)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

//query搜索
func (e *gomEs) Search(index string, query elastic.Query, size, page int) (*elastic.SearchResult, error) {
	srv := e.orm.Search(index).Query(query).Pretty(true)
	if size > 0 {
		srv.Size(size)
	}
	if page > 0 {
		srv.From((page - 1) * size)
	}
	result, err := srv.Do(e.ctx)
	if err != nil {
		return result, err
	}

	return result, nil
}

//aggregation搜索
func (e *gomEs) AggsSearch(index string, aggName string, agg elastic.Aggregation, size, page int) (*elastic.SearchResult, error) {
	srv := e.orm.Search(index).Pretty(true)
	if size > 0 {
		srv.Size(size)
	}
	if page > 0 {
		srv.From((page - 1) * size)
	}
	return srv.Aggregation(aggName, agg).Do(e.ctx)
}

//模板搜索
func (e *gomEs) TemplateSearch(index string, body interface{}) (*elastic.SearchResult, error) {
	// Get HTTP response
	res, err := e.orm.PerformRequest(e.ctx, elastic.PerformRequestOptions{
		Method: "GET",
		Path:   "/_search/template?index=" + index,
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	//组织返回数据
	var ret elastic.SearchResult
	if err := json.Unmarshal(res.Body, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

//获取搜索结果数据
func (e *gomEs) GetSearchResultData(ret *elastic.SearchResult) (int64, []json.RawMessage) {
	total := ret.Hits.TotalHits.Value
	if total == 0 {
		return total, nil
	}
	var list []json.RawMessage
	for _, hit := range ret.Hits.Hits {
		list = append(list, hit.Source)
	}
	return total, list
}

//获取搜索结果数据并返回ID
func (e *gomEs) GetSearchResultDataId(ret *elastic.SearchResult) (int64, []*ListItem) {
	total := ret.Hits.TotalHits.Value
	if total == 0 {
		return total, nil
	}

	listres := make([]*ListItem, 0)
	for _, hit := range ret.Hits.Hits {
		var res ListItem
		res.Row = hit.Source
		res.Id = hit.Id
		listres = append(listres, &res)
	}
	return total, listres
}

//获取搜索结果数据
func (e *gomEs) GetAggregationResultData(ret *elastic.SearchResult) (int64, map[string]json.RawMessage) {
	total := ret.Hits.TotalHits.Value
	if total == 0 {
		return total, nil
	}
	return total, ret.Aggregations
}
*/
