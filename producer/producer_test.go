package producer


import (
    "testing"
    "github.com/chenxiaoli/crawler/producer"
    "github.com/chenxiaoli/crawler/models"
)

func Test_URLProducer(t *testing.T) {

var note models.PageSaveNote
producer.URLProducer(note)
}