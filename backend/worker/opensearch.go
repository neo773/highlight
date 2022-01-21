package worker

import (
	"reflect"

	"github.com/highlight-run/highlight/backend/model"
	"github.com/highlight-run/highlight/backend/opensearch"
	e "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const BATCH_SIZE = 2000

func (w *Worker) indexItem(index opensearch.Index, item interface{}) {
	val := reflect.ValueOf(item).Elem()
	id := int(val.FieldByName("ID").Int())

	// Add an item to the indexer
	if err := w.Resolver.OpenSearch.Index(index, id, item); err != nil {
		log.Error(e.Wrap(err, "OPENSEARCH_ERROR error adding field to the indexer"))
	}
}

func (w *Worker) IndexSessions(isUpdate bool) {
	modelProto := &model.Session{}
	results := &[]*model.Session{}

	inner := func(tx *gorm.DB, batch int) error {
		for _, result := range *results {
			fields := []*opensearch.OpenSearchField{}
			for _, field := range result.Fields {
				f := opensearch.OpenSearchField{
					Field:    field,
					Key:      field.Type + "_" + field.Name,
					KeyValue: field.Type + "_" + field.Name + "_" + field.Value,
				}
				fields = append(fields, &f)
			}
			os := opensearch.OpenSearchSession{
				Session: result,
				Fields:  fields,
			}
			w.indexItem(opensearch.IndexSessions, &os)
		}
		return nil
	}

	whereClause := "True"
	if isUpdate {
		whereClause = "updated_at > NOW() - interval '1 day'"
	}

	if err := w.Resolver.DB.Preload("Fields").Where(whereClause).Model(modelProto).
		FindInBatches(results, BATCH_SIZE, inner).Error; err != nil {
		log.Fatalf("OPENSEARCH_ERROR error querying objects: %+v", err)
	}
}

func (w *Worker) IndexErrors(isUpdate bool) {
	modelProto := &model.ErrorGroup{}
	results := &[]*model.ErrorGroup{}

	inner := func(tx *gorm.DB, batch int) error {
		for _, result := range *results {
			fields := []*opensearch.OpenSearchErrorField{}
			for _, field := range result.Fields {
				f := opensearch.OpenSearchErrorField{
					ErrorField: field,
					Key:        field.Name,
					KeyValue:   field.Name + "_" + field.Value,
				}
				fields = append(fields, &f)
			}
			var filename *string
			if result.MappedStackTrace != nil {
				filename = model.GetFirstFilename(*result.MappedStackTrace)
			} else {
				filename = model.GetFirstFilename(result.StackTrace)
			}
			result.FieldGroup = nil
			result.Fields = nil
			result.Environments = ""
			result.MappedStackTrace = nil
			result.StackTrace = ""
			os := opensearch.OpenSearchError{
				ErrorGroup: result,
				Fields:     fields,
				Filename:   filename,
			}
			w.indexItem(opensearch.IndexErrors, &os)
		}
		return nil
	}

	whereClause := "True"
	if isUpdate {
		whereClause = "updated_at > NOW() - interval '1 day'"
	}

	// A little hacky, but some of the error groups with low ids have a very high number of fields,
	// and it causes and error when > 65536 fields are loaded at once
	if err := w.Resolver.DB.Preload("Fields").Where(whereClause).Where("id <= 10").Model(modelProto).
		FindInBatches(results, BATCH_SIZE, inner).Error; err != nil {
		log.Fatalf("OPENSEARCH_ERROR error querying objects: %+v", err)
	}

	if err := w.Resolver.DB.Preload("Fields").Where(whereClause).Where("id > 10").Model(modelProto).
		FindInBatches(results, BATCH_SIZE, inner).Error; err != nil {
		log.Fatalf("OPENSEARCH_ERROR error querying objects: %+v", err)
	}
}

func (w *Worker) IndexTable(index opensearch.Index, modelPrototype interface{}, isUpdate bool) {
	modelProto := modelPrototype

	whereClause := "True"
	if isUpdate {
		whereClause = "updated_at > NOW() - interval '1 day'"
	}

	rows, err := w.Resolver.DB.Model(modelProto).
		Where(whereClause).
		Order("created_at asc").Rows()
	if err != nil {
		log.Fatalf("OPENSEARCH_ERROR error retrieving objects: %+v", err)
	}

	for rows.Next() {
		modelObj := modelPrototype
		if err := w.Resolver.DB.ScanRows(rows, modelObj); err != nil {
			log.Fatalf("OPENSEARCH_ERROR error scanning rows: %+v", err)
		}

		w.indexItem(index, modelObj)
	}
}

const NESTED_FIELD_MAPPINGS = `
{
	"properties": {
		"fields": {
			"properties": {
				"Key": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"KeyValue": {
					"type": "keyword",
					"normalizer": "lowercase"
				}
			}
		}
	}
}`

const FIELD_MAPPINGS = ` 
{
	"properties": {
		"Value": {
			"type": "search_as_you_type"
		}
	}
}`

// Had to write the "source" field as one line, wouldn't accept triple quotes
// even though it works in the OpenSearch dashboard.
const FIELD_APPEND_SCRIPT = `
{
	"script": {
		"lang": "painless",
		"source": "def ids = new ArrayList();for (int i = 0; i < ctx._source[params.fieldName].length; i += 1) {ids.add(ctx._source[params.fieldName].get(i).id);}for (def id : ids) {int len = params.toAppend.length;for (int i = len - 1; i >= 0; i--) {def cur_item_id = params.toAppend.get(i).id;if (id.equals(cur_item_id)) {params.toAppend.remove(i);}}}if (params.toAppend.length > 0) {ctx._source[params.fieldName].addAll(params.toAppend);} else {ctx.op = \"noop\";}"
	}
}
`

func (w *Worker) InitIndexMappings() {
	if err := w.Resolver.OpenSearch.PutMapping(opensearch.IndexSessions, NESTED_FIELD_MAPPINGS); err != nil {
		log.Warnf("OPENSEARCH_ERROR error creating session mappings: %+v", err)
	}
	if err := w.Resolver.OpenSearch.PutMapping(opensearch.IndexFields, FIELD_MAPPINGS); err != nil {
		log.Warnf("OPENSEARCH_ERROR error creating field mappings: %+v", err)
	}
	if err := w.Resolver.OpenSearch.PutMapping(opensearch.IndexErrors, NESTED_FIELD_MAPPINGS); err != nil {
		log.Warnf("OPENSEARCH_ERROR error creating error mappings: %+v", err)
	}
	if err := w.Resolver.OpenSearch.PutMapping(opensearch.IndexErrorFields, FIELD_MAPPINGS); err != nil {
		log.Warnf("OPENSEARCH_ERROR error creating error field mappings: %+v", err)
	}
	if err := w.Resolver.OpenSearch.PutScript(opensearch.ScriptAppendFields, FIELD_APPEND_SCRIPT); err != nil {
		log.Warnf("OPENSEARCH_ERROR error creating field append script: %+v", err)
	}
}
