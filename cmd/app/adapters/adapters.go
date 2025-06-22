package adapters

import (
	"testWorkmate/internal/adapter/im_db"

	"github.com/sirupsen/logrus"
)

type Adapters struct {
	imDb *im_db.ImitationDb
}

func NewAdapters() *Adapters {
	return &Adapters{
		imDb: nil,
	}
}

func (a *Adapters) GetImDb() *im_db.ImitationDb {
	return a.imDb
}

func (a *Adapters) MustInit() {
	logrus.Info("adapters start initializing")
	{
		a.imDb = im_db.NewImitationDb()
		logrus.Info("imDb initialized")
	}
	logrus.Info("done")
}
