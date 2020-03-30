package slacker

// SBOption is an option function to parametrize SlackSender
type SBOption func(*SendBuffer)

// SetOpts allows to set options
func (ss *SendBuffer) SetOpts(opts ...SBOption) {
	for _, opt := range opts {
		opt(ss)
	}
}

// FlushString sets the flushstring to s
func FlushString(s string) func(ss *SendBuffer) {
	return func(ss *SendBuffer) {
		ss.useFlushString = true
		ss.hideFlushString = true
		ss.flushString = s
	}
}

// SetHideFlushString sets the hideflush string property
func SetHideFlushString(b bool) func(ss *SendBuffer) {
	return func(ss *SendBuffer) {
		ss.hideFlushString = b
	}
}

// AlwaysFlush sets the alwaysflush bool property
func AlwaysFlush(b bool) func(ss *SendBuffer) {
	return func(ss *SendBuffer) {
		ss.alwaysFlush = b
	}
}
