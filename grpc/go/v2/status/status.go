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

	// CodeErrorIDRequired ...
	CodeErrorIDRequired

	// ================ validation story ================ //

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
	// CodeErrorValidateID ...
	CodeErrorValidateID
	// CodeErrorValidateCheckLanguageOnIA ...
	CodeErrorValidateCheckLanguageOnIA
	// CodeErrorValidateDst ...
	CodeErrorValidateDst

	// ================ validation chapter ================ //

	// CodeErrorValidateContent ...
	CodeErrorValidateContent
	// CodeErrorValidateStoryID ...
	CodeErrorValidateStoryID
	// CodeErrorValidateStoryID ...
	CodeErrorValidateMaxChaptersPerStory
	// CodeErrorValidateStoryHasEnd ...
	CodeErrorValidateStoryHasEnd
	// CodeErrorValidatePositionInvalid ...
	CodeErrorValidatePositionInvalid
	// CodeErrorValidateChaptersOrderNotConsecutive
	CodeErrorValidateChaptersOrderNotConsecutive
	// CodeErrorValidateAllIDRequired ...
	CodeErrorValidateAllIDRequired
	// CodeErrorValidateNoneIDMatchChaptersID ninguno de los id de capitulos en la request conciden con los capitulos de la historia
	CodeErrorValidateNoneIDMatchChaptersID
	// CodeErrorValidateItemsRequired ...
	CodeErrorValidateItemsRequired
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

	// StatusErrorIDRequired ...
	StatusErrorIDRequired = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorIDRequired,
		Message:  "id is required",
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
		Message:  "the number the characters allowed for the character name is between 1 and 150",
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
		Message:  "author is required because story no is original(max number of characters is 150)",
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
		Message:  "tags length allowed is between 1 and 50",
	}
	// StatusErrorValidateTag ...
	StatusErrorValidateTag = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateTag,
		Message:  "the number the characters allowed for the tag name is between 1 and 80",
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
		Message:  "the number the characters allowed for the description is between 60 and 2000",
	}
	// StatusErrorValidateDescription ...
	StatusErrorValidateID = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateID,
		Message:  "id is required",
	}
	// StatusErrorValidateCheckLanguageOnIA ...
	StatusErrorValidateCheckLanguageOnIA = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCheckLanguageOnIA,
		Message:  "failed check the language-text on ia",
	}
	// StatusErrorValidateDst ...
	StatusErrorValidateDst = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateDst,
		Message:  "destination is invalid",
	}

	// ================ validation chapter ================ //

	// StatusErrorValidateContent ...
	StatusErrorValidateContent = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateContent,
		Message:  "the number the characters allowed for the description is between 60 and 100000",
	}
	// StatusErrorValidateStoryID ...
	StatusErrorValidateStoryID = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateStoryID,
		Message:  "story id is required",
	}
	// StatusErrorValidateMaxChaptersPerStory ...
	StatusErrorValidateMaxChaptersPerStory = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateMaxChaptersPerStory,
		Message:  "maximum number of chapters per story is 2500",
	}
	// StatusErrorValidateStoryHasEnd ...
	StatusErrorValidateStoryHasEnd = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateMaxChaptersPerStory,
		Message:  "the story already has an end",
	}

	// StatusErrorValidatePositionInvalid ...
	StatusErrorValidatePositionInvalid = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidatePositionInvalid,
		Message:  "first position is not 1 or is ivalid",
	}
	// StatusErrorValidateChaptersOrderNotConsecutive ...
	StatusErrorValidateChaptersOrderNotConsecutive = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateChaptersOrderNotConsecutive,
		Message:  "the order of the chapters is not consecutive",
	}
	// StatusErrorValidateAllIDRequired ...
	StatusErrorValidateAllIDRequired = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateAllIDRequired,
		Message:  "all id's are required",
	}
	// StatusErrorValidateNoneIDMatchChaptersID ...
	StatusErrorValidateNoneIDMatchChaptersID = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateNoneIDMatchChaptersID,
		Message:  "none of the chapter id match any chapter id of the story",
	}
	// StatusErrorValidateItemsRequired ...
	StatusErrorValidateItemsRequired = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateItemsRequired,
		Message:  "items are required",
	}
)
var codeStatus = map[commonsStatus.Code]*commonsStatus.Status{

	CodeStoryNotFound:        StatusStoryNotFound,
	CodeChapterNotFound:      StatusChapterNotFound,
	CodeErrorUserLoggedWrong: StatusErrorUserLoggedWrong,

	CodeErrorIDRequired: StatusErrorIDRequired,

	// ================ validation story ================ //

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
	CodeErrorValidateID:                 StatusErrorValidateID,
	CodeErrorValidateCheckLanguageOnIA:  StatusErrorValidateCheckLanguageOnIA,
	CodeErrorValidateDst:                StatusErrorValidateDst,

	// ================ validation chapter ================ //

	CodeErrorValidateContent:             StatusErrorValidateContent,
	CodeErrorValidateStoryID:             StatusErrorValidateStoryID,
	CodeErrorValidateMaxChaptersPerStory: StatusErrorValidateMaxChaptersPerStory,
	CodeErrorValidateStoryHasEnd:         StatusErrorValidateStoryHasEnd,

	CodeErrorValidatePositionInvalid:             StatusErrorValidatePositionInvalid,
	CodeErrorValidateChaptersOrderNotConsecutive: StatusErrorValidateChaptersOrderNotConsecutive,
	CodeErrorValidateAllIDRequired:               StatusErrorValidateAllIDRequired,
	CodeErrorValidateNoneIDMatchChaptersID:       StatusErrorValidateNoneIDMatchChaptersID,
	CodeErrorValidateItemsRequired:               StatusErrorValidateItemsRequired,
}

// FromCode ...
func FromCode(code commonsStatus.Code) *commonsStatus.Status {
	if v, ok := codeStatus[code]; ok {
		return v
	}
	return commonsStatus.FromCode(code)
}
