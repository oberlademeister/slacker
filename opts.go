package slacker

// SlackSenderOption is an option function to parametrize SlackSender
type SlackSenderOption func(*SlackSender)

// SetOpts allows to set options
func (ss *SlackSender) SetOpts(opts ...SlackSenderOption) {
	for _, opt := range opts {
		opt(ss)
	}
}

// FlushString sets the flushstring to s
func FlushString(s string) func(ss *SlackSender) {
	return func(ss *SlackSender) {
		ss.useFlushString = true
		ss.hideFlushString = true
		ss.flushString = s
	}
}

// SetHideFlushString sets the hideflush string property
func SetHideFlushString(b bool) func(ss *SlackSender) {
	return func(ss *SlackSender) {
		ss.hideFlushString = b
	}
}

// AlwaysFlush sets the alwaysflush bool property
func AlwaysFlush(b bool) func(ss *SlackSender) {
	return func(ss *SlackSender) {
		ss.alwaysFlush = b
	}
}
