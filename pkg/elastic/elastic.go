package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/nikitalier/tenderMonitoring/config"
)

//Writer ...
type Writer struct {
	client  *elasticsearch.Client
	index   string
	fields  LogField
	appName string
	msg     []byte
}

//LogField ...
type LogField struct {
	AppName string
}

//LogRecord ...
type LogRecord struct {
	Field    LogField
	Data     json.RawMessage
	DateTime time.Time
}

func readConfig(cfg *config.Logging) (c elasticsearch.Config) {
	c = elasticsearch.Config{
		Addresses: cfg.ElasticsearchHost,
	}

	return c
}

//New ...
func New(index, appName string, cfg *config.Logging) (w *Writer, err error) {
	es, err := elasticsearch.NewClient(readConfig(cfg))
	if err != nil {
		err = fmt.Errorf("Error init elasticsearcg: %w", err)
		return
	}

	w = &Writer{
		client: es,
		index:  index,
		fields: LogField{
			AppName: appName,
		},
	}

	return w, err
}

func (w *Writer) sendMessage(index string, msg []byte) error {
	dt := time.Now()

	r := LogRecord{
		Field:    w.fields,
		Data:     msg,
		DateTime: dt,
	}

	data, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}

	resp, err := w.client.Index(index, bytes.NewReader(data))
	if err != nil {
		err = fmt.Errorf("Error sending elastic log: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		log.Println(resp)
	}

	return err
}

func (w *Writer) Write(b []byte) (int, error) {
	d := make([]byte, len(b))
	n := copy(d, b)

	w.sendMessage(w.index, b)
	return n, nil
}
