package gconf

type reloadToken struct {
	changed  bool
	callback func()
}

func NewChangeToken() ChangeToken {
	return &reloadToken{
		changed:  false,
		callback: nil,
	}
}

func (r *reloadToken) SetCallback(callback func()) {
	r.callback = callback
}

func (r *reloadToken) HasChanged() bool {
	return r.changed
}

func (r *reloadToken) SetAsChanged() {
	r.changed = true
}

func (r *reloadToken) OnChanged() {
	r.SetAsChanged()

	if r.callback == nil {
		return
	}

	go r.callback()
}