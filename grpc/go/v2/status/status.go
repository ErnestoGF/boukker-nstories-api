package status

import (
	commonsStatus "github.com/ErnestoGF/boukker-commons/go/v2/grpcutil/status"
	"google.golang.org/grpc/codes"
)

const (
	// CodeStoryNotFound ...
	CodeStoryNotFound commonsStatus.Code = iota + 10000

	// CodeChapterNotFound ...
	CodeChapterNotFound
)

var (
	// StatusStoryNotFound ...
	StatusStoryNotFound = &commonsStatus.Status{
		CodeGRPC: codes.NotFound,
		Code:     CodeStoryNotFound,
		Message:  "Story not found",
	}
	// StatusChapterNotFound ...
	StatusChapterNotFound = &commonsStatus.Status{
		CodeGRPC: codes.NotFound,
		Code:     CodeChapterNotFound,
		Message:  "Chapter not found",
	}
)
var codeStatus = map[commonsStatus.Code]*commonsStatus.Status{

	CodeStoryNotFound:   StatusStoryNotFound,
	CodeChapterNotFound: StatusChapterNotFound,
}

// FromCode ...
func FromCode(code commonsStatus.Code) *commonsStatus.Status {
	if v, ok := codeStatus[code]; ok {
		return v
	}
	return commonsStatus.FromCode(code)
}
