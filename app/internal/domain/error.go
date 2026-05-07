package domain

import "errors"

var ErrEmptyCpuInfoPath = errors.New("empty cpu info path variable")
var ErrEmptyMemInfoPath = errors.New("empty mem info path variable")
var ErrEmptyMemInfoLimitPath = errors.New("empty mem info limit variable")

var ErrCpuModelNotFound = errors.New("cpu model not found")
var ErrMemToatlKbNotFound = errors.New("mem total kb not found")
