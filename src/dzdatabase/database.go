package dzdatabase

import (
	"container/list"
	"labix.org/v2/mgo"
)

type DZDatabaseSession struct {
	Session *mgo.Session
}

type DZDatabaseSessionPool struct {
	usingSessions   *list.List
	unUsingSessions *list.List
}

func (d *DZDatabaseSession) PrepareForReuse() {

}

func (d *DZDatabaseSession) CollectionPictures() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionPictures)
}

func (d *DZDatabaseSession) CollectionDeletedObjects() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionDeletedObjects)
}

func (d *DZDatabaseSession) CollectionUsers() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionUsers)
}

func (d *DZDatabaseSession) CollectionTimes() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDatabaseColletionTimes)
}

func (d *DZDatabaseSession) CollectionDiveces() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionDevices)
}

func (d *DZDatabaseSession) CollectionVersions() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionVersions)
}

func (d *DZDatabaseSession) CollectionTimeTypes() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionTimeTypes)
}

func (d *DZDatabaseSession) CollectionApps() *mgo.Collection {
	return d.Session.DB(DZDataBaseName).C(DZDataBaseColletionApps)
}

func (d *DZDatabaseSessionPool) Init() {
	d.usingSessions = list.New()
	d.unUsingSessions = list.New()
}

func (d *DZDatabaseSessionPool) OneSession() *DZDatabaseSession {

	if d.unUsingSessions.Len() > 0 {
		e := d.unUsingSessions.Back()
		s := e.Value
		s.(*DZDatabaseSession).PrepareForReuse()
		d.usingSessions.PushBack(s)
		return s.(*DZDatabaseSession)
	} else {
		session, _ := mgo.Dial(DZDataBaseServerUrl)
		dzs := &DZDatabaseSession{session}
		d.usingSessions.PushBack(dzs)
		return dzs
	}
}

func (d *DZDatabaseSessionPool) EndUseSession(s *DZDatabaseSession) {
	if s != nil {
		var elemnt *list.Element
		for e := d.usingSessions.Front(); e != nil; e = e.Next() {
			if e.Value == s {
				elemnt = e
			}
		}
		if elemnt != nil {
			d.usingSessions.Remove(elemnt)
			d.unUsingSessions.PushBack(s)
		}
	}
}

var _shareInstance *DZDatabaseSessionPool = nil

func ShareDBSessionPool() *DZDatabaseSessionPool {
	if _shareInstance == nil {
		_shareInstance = new(DZDatabaseSessionPool)
		_shareInstance.Init()
	}
	return _shareInstance
}
