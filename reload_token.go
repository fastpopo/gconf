package gconf

type _ReloadToken struct {
	changed  bool
	callback func()
}

func NewReloadToken() ReloadToken {
	return &_ReloadToken{
		changed:  false,
		callback: nil,
	}
}

func (r *_ReloadToken) SetCallback(callback func()) {
	r.callback = callback
}

func (r *_ReloadToken) HasChanged() bool {
	return r.changed
}

func (r *_ReloadToken) SetAsChanged() {
	r.changed = true
}

func (r *_ReloadToken) OnReload() {
	r.SetAsChanged()

	if r.callback == nil {
		return
	}

	r.callback()
}
