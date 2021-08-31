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

	// CodeErrorUserLoggedWrong ...
	CodeErrorUserLoggedWrong

	// CodeErrorValidateAudience ...
	CodeErrorValidateAudience
	// CodeErrorValidateCategoryID ...
	CodeErrorValidateCategoryID
	// CodeErrorValidateCharacter ...
	CodeErrorValidateCharacter
	// CodeErrorValidateClasification ...
	CodeErrorValidateClasification
	// CodeErrorValidateCommentsModeration ...
	CodeErrorValidateCommentsModeration
	// CodeErrorValidateCopyright ...
	CodeErrorValidateCopyright
	// CodeErrorValidateAuthor ...
	CodeErrorValidateAuthor
	// CodeErrorValidateLanguageID ...
	CodeErrorValidateLanguageID
	// CodeErrorValidateStatus ...
	CodeErrorValidateStatus
	// CodeErrorValidateTags ...
	CodeErrorValidateTags
	// CodeErrorValidateTag ...
	CodeErrorValidateTag
	// CodeErrorValidateTitle ...
	CodeErrorValidateTitle
	// CodeErrorValidateTokenCover ...
	CodeErrorValidateTokenCover
	// CodeErrorValidateDescription ...
	CodeErrorValidateDescription
	// CodeErrorValidateCheckLanguageOnIA ...
	CodeErrorValidateCheckLanguageOnIA
)

var (
	// StatusStoryNotFound ...
	StatusStoryNotFound = &commonsStatus.Status{
		CodeGRPC: codes.NotFound,
		Code:     CodeStoryNotFound,
		Message:  "story not found",
	}
	// StatusChapterNotFound ...
	StatusChapterNotFound = &commonsStatus.Status{
		CodeGRPC: codes.NotFound,
		Code:     CodeChapterNotFound,
		Message:  "chapter not found",
	}
	// StatusErrorUserLoggedWrong ...
	StatusErrorUserLoggedWrong = &commonsStatus.Status{
		CodeGRPC: codes.FailedPrecondition,
		Code:     CodeErrorUserLoggedWrong,
		Message:  "user not login or user id incorrect",
	}

	// ================ validation story ================ //

	// StatusErrorValidateAudience ...
	StatusErrorValidateAudience = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateAudience,
		Message:  "audience is invalid",
	}
	// StatusErrorValidateCategoryID invalid user ID
	StatusErrorValidateCategoryID = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCategoryID,
		Message:  "category id is required",
	}
	// StatusErrorValidateCharacter ...
	StatusErrorValidateCharacter = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCharacter,
		Message:  "the number the characters allowed for the character name is between 1 and 50",
	}
	// StatusErrorValidateClasification ...
	StatusErrorValidateClasification = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateClasification,
		Message:  "invalid clasification",
	}
	// StatusErrorValidateCommentsModeration ...
	StatusErrorValidateCommentsModeration = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCommentsModeration,
		Message:  "comments moderation is invalid",
	}
	// StatusErrorValidateCopyright ...
	StatusErrorValidateCopyright = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCopyright,
		Message:  "copyright is invalid",
	}
	// StatusErrorValidateAuthor ...
	StatusErrorValidateAuthor = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateAuthor,
		Message:  "author is required because story no is original",
	}
	// StatusErrorValidateLanguageID ...
	StatusErrorValidateLanguageID = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateLanguageID,
		Message:  "language id is required",
	}
	// StatusErrorValidateStatus ...
	StatusErrorValidateStatus = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateStatus,
		Message:  "status is invalid",
	}
	// StatusErrorValidateTags ...
	StatusErrorValidateTags = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateTags,
		Message:  "tags length allowed is between 1 and 20",
	}
	// StatusErrorValidateTag ...
	StatusErrorValidateTag = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateTag,
		Message:  "the number the characters allowed for the tag name is between 1 and 50",
	}
	// StatusErrorValidateTitle ...
	StatusErrorValidateTitle = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateTitle,
		Message:  "the number the characters allowed for the title is between 1 and 150",
	}
	// StatusErrorValidateTokenCover ...
	StatusErrorValidateTokenCover = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateTokenCover,
		Message:  "token cover is required or invalid",
	}
	// StatusErrorValidateDescription ...
	StatusErrorValidateDescription = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateDescription,
		Message:  "the number the characters allowed for the description is between 1 and 250",
	}
	// StatusErrorValidateCheckLanguageOnIA ...
	StatusErrorValidateCheckLanguageOnIA = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateDescription,
		Message:  "the language does not match the text",
	}
)
var codeStatus = map[commonsStatus.Code]*commonsStatus.Status{

	CodeStoryNotFound:        StatusStoryNotFound,
	CodeChapterNotFound:      StatusChapterNotFound,
	CodeErrorUserLoggedWrong: StatusErrorUserLoggedWrong,

	// ==============================

	CodeErrorValidateAudience:           StatusErrorValidateAudience,
	CodeErrorValidateCategoryID:         StatusErrorValidateCategoryID,
	CodeErrorValidateCharacter:          StatusErrorValidateCharacter,
	CodeErrorValidateClasification:      StatusErrorValidateClasification,
	CodeErrorValidateCommentsModeration: StatusErrorValidateCommentsModeration,
	CodeErrorValidateCopyright:          StatusErrorValidateCopyright,
	CodeErrorValidateAuthor:             StatusErrorValidateAuthor,
	CodeErrorValidateLanguageID:         StatusErrorValidateLanguageID,
	CodeErrorValidateStatus:             StatusErrorValidateStatus,
	CodeErrorValidateTags:               StatusErrorValidateTags,
	CodeErrorValidateTag:                StatusErrorValidateTag,
	CodeErrorValidateTitle:              StatusErrorValidateTitle,
	CodeErrorValidateTokenCover:         StatusErrorValidateTokenCover,
	CodeErrorValidateDescription:        StatusErrorValidateDescription,
	CodeErrorValidateCheckLanguageOnIA:  StatusErrorValidateCheckLanguageOnIA,
}

// FromCode ...
func FromCode(code commonsStatus.Code) *commonsStatus.Status {
	if v, ok := codeStatus[code]; ok {
		return v
	}
	return commonsStatus.FromCode(code)
}
