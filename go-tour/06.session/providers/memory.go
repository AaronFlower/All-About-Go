package memory

import (
	"container/list"
	"sync"
	"time"

	"github.com/aaronflower/ago/6.session"
)

// Provider uses memory to implements session store.
type Provider struct {
	lock     sync.Mutex               // 用来做锁
	sessions map[string]*list.Element // 用来做存储
	list     *list.List               // 用来做 gc
}

// SessionInit initializes a new session.
func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAcessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// SessionRead starts a new session
func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	}
	sess, err := pder.SessionInit(sid)
	return sess, err
}

// SessionDestroy destroys a session.
func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

// SessionGC deletes expired data.
func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAcessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// sessionUpdate updates the session access time.
func (pder *Provider) sessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAcessed = time.Now()
		pder.list.MoveToFront(element)
	}
	return nil
}

var pder = &Provider{list: list.New()}

// SessionStore sessions store.
type SessionStore struct {
	sid         string                      // session id
	timeAcessed time.Time                   // 最后访问时间
	value       map[interface{}]interface{} // session 里面存储的值
}

// Set sets a pair of (key, value)
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.sessionUpdate(st.sid)
	return nil
}

// Get returns the value fo the session by key.
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.sessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete deletes the value of session.
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.sessionUpdate(st.sid)
	return nil
}

// SessionID returns the session id.
func (st *SessionStore) SessionID() string {
	return st.sid
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}
