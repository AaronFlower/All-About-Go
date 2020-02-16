package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Session to handle http session.
type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value.
	SessionID() string                // back current sessionID
}

// Provider 接口用来定义 session 管理器的底层存储结构。
// 因为 session 可以用内存、数据库或文件存储，所以我们只定义接口。
// 具体实现根据自己需要来实现了。
type Provider interface {
	SessionInit(sid string) (Session, error) // Init a session
	SessionRead(sid string) (Session, error) // Get a session or create a session then return.
	SessionDestroy(sid string) error         // destory a session
	SessionGC(maxlifetime int64)             // delete expired data.
}

var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice the same name or if driver is nil, it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

// Manager to manage the application sessions.
type Manager struct {
	cookieName  string     // private cookieName, 一般都中 sessionid 吧.
	lock        sync.Mutex // protects session
	provider    Provider
	maxlifetime int64
}

// NewManager 创建一个全局的 session 管理器，要 provideName 提供的类型
// provideName 可以是内存，数据库之类的 provide, 前提是要选注册，如没有注册的话
// 会 panic.
func NewManager(providerName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknow provider %q (forgotten import ?)", providerName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// SessionStart go get session
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxlifetime),
		}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy destroys the sessionid
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName) // 根据 cookieName 获取 cookie 值.
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionDestroy(cookie.Value)
	expiration := time.Now()
	cookie2 := http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiration,
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie2)
}

// GC deletes expired data
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxlifetime)
	time.AfterFunc(time.Duration(manager.maxlifetime)*time.Second, func() { manager.GC() })
}

func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}
