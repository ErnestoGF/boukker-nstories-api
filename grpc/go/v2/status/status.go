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

	// CodeBookmarkNotFound ...
	CodeBookmarkNotFound

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
	// CodeErrorValidateCheckLanguageOnKathe ...
	CodeErrorValidateCheckLanguageOnKathe
	// CodeErrorValidateDst ...
	CodeErrorValidateDst

	// CodeErrorValidateStoryFinished ...
	CodeErrorValidateStoryFinished
	// CodeErrorValidateStoryWithoutChaptersNotFinished ...
	CodeErrorValidateStoryWithoutChaptersNotFinished

	// ================ validation chapter ================ //

	// CodeErrorValidateContent ...
	CodeErrorValidateContent
	// CodeErrorValidateStoryID ...
	CodeErrorValidateStoryID
	// CodeErrorValidateMaxChaptersPerStory ...
	CodeErrorValidateMaxChaptersPerStory
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

	// ================ validation bookmarks ================ //

	// CodeErrorValidateAbstract ...
	CodeErrorValidateAbstract
	// CodeErrorValidatePath ...
	CodeErrorValidatePath

	// ================ validation list stories ================ //

	// CodeErrorValidateListStoriesFilterRangePublished valida el rango de fecha
	CodeErrorValidateListStoriesFilterRangePublished
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

	// StatusStoryNotFound ...
	StatusBookmarkNotFound = &commonsStatus.Status{
		CodeGRPC: codes.NotFound,
		Code:     CodeBookmarkNotFound,
		Message:  "bookmark not found",
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
	// StatusErrorValidateCheckLanguageOnKathe ...
	StatusErrorValidateCheckLanguageOnKathe = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateCheckLanguageOnKathe,
		Message:  "failed check the language-text on kathe",
	}
	// StatusErrorValidateDst ...
	StatusErrorValidateDst = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateDst,
		Message:  "destination is invalid",
	}

	// StatusErrorValidateStoryFinished ...
	StatusErrorValidateStoryFinished = &commonsStatus.Status{
		CodeGRPC: codes.FailedPrecondition,
		Code:     CodeErrorValidateStoryFinished,
		Message:  "the story is finished",
	}
	// StatusErrorValidateStoryWithoutChaptersNotFinished ...
	StatusErrorValidateStoryWithoutChaptersNotFinished = &commonsStatus.Status{
		CodeGRPC: codes.FailedPrecondition,
		Code:     CodeErrorValidateStoryWithoutChaptersNotFinished,
		Message:  "story without chapters cannot be finalized",
	}

	// ================ validation chapter ================ //

	// StatusErrorValidateContent ...
	StatusErrorValidateContent = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateContent,
		Message:  "the number the characters allowed for the description is between 60 and 2000",
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
		Message:  "maximum number of chapters per story is 200",
	}

	// StatusErrorValidatePositionInvalid ...
	StatusErrorValidatePositionInvalid = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidatePositionInvalid,
		Message:  "first position is not 1 or is less than 1",
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

	// ================ validation bookmark ================ //

	// StatusErrorValidateAbstract ...
	StatusErrorValidateAbstract = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateAbstract,
		Message:  "the number the characters allowed for the abstract is between 1 and 500",
	}
	// StatusErrorValidatePath ...
	StatusErrorValidatePath = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidatePath,
		Message:  "the number the characters allowed for the path is between 1 and 150",
	}

	// ================ validation list stories ================ //

	// StatusErrorValidateListStoriesFilterRangePublished ...
	StatusErrorValidateListStoriesFilterRangePublished = &commonsStatus.Status{
		CodeGRPC: codes.InvalidArgument,
		Code:     CodeErrorValidateListStoriesFilterRangePublished,
		Message:  "filter range published at is invalid",
	}
)
var codeStatus = map[commonsStatus.Code]*commonsStatus.Status{

	CodeStoryNotFound:        StatusStoryNotFound,
	CodeChapterNotFound:      StatusChapterNotFound,
	CodeErrorUserLoggedWrong: StatusErrorUserLoggedWrong,

	CodeErrorIDRequired: StatusErrorIDRequired,

	CodeBookmarkNotFound: StatusBookmarkNotFound,

	// ================ validation story ================ //

	CodeErrorValidateAudience:             StatusErrorValidateAudience,
	CodeErrorValidateCategoryID:           StatusErrorValidateCategoryID,
	CodeErrorValidateCharacter:            StatusErrorValidateCharacter,
	CodeErrorValidateClasification:        StatusErrorValidateClasification,
	CodeErrorValidateCommentsModeration:   StatusErrorValidateCommentsModeration,
	CodeErrorValidateCopyright:            StatusErrorValidateCopyright,
	CodeErrorValidateAuthor:               StatusErrorValidateAuthor,
	CodeErrorValidateLanguageID:           StatusErrorValidateLanguageID,
	CodeErrorValidateStatus:               StatusErrorValidateStatus,
	CodeErrorValidateTags:                 StatusErrorValidateTags,
	CodeErrorValidateTag:                  StatusErrorValidateTag,
	CodeErrorValidateTitle:                StatusErrorValidateTitle,
	CodeErrorValidateTokenCover:           StatusErrorValidateTokenCover,
	CodeErrorValidateDescription:          StatusErrorValidateDescription,
	CodeErrorValidateID:                   StatusErrorValidateID,
	CodeErrorValidateCheckLanguageOnKathe: StatusErrorValidateCheckLanguageOnKathe,
	CodeErrorValidateDst:                  StatusErrorValidateDst,

	CodeErrorValidateStoryFinished:                   StatusErrorValidateStoryFinished,
	CodeErrorValidateStoryWithoutChaptersNotFinished: StatusErrorValidateStoryWithoutChaptersNotFinished,

	// ================ validation chapter ================ //

	CodeErrorValidateContent:             StatusErrorValidateContent,
	CodeErrorValidateStoryID:             StatusErrorValidateStoryID,
	CodeErrorValidateMaxChaptersPerStory: StatusErrorValidateMaxChaptersPerStory,

	CodeErrorValidatePositionInvalid:             StatusErrorValidatePositionInvalid,
	CodeErrorValidateChaptersOrderNotConsecutive: StatusErrorValidateChaptersOrderNotConsecutive,
	CodeErrorValidateAllIDRequired:               StatusErrorValidateAllIDRequired,
	CodeErrorValidateNoneIDMatchChaptersID:       StatusErrorValidateNoneIDMatchChaptersID,
	CodeErrorValidateItemsRequired:               StatusErrorValidateItemsRequired,

	// ================ validation bookmark ================ //

	CodeErrorValidatePath:     StatusErrorValidatePath,
	CodeErrorValidateAbstract: StatusErrorValidateAbstract,

	// ================ validation list stories ================ //

	CodeErrorValidateListStoriesFilterRangePublished: StatusErrorValidateListStoriesFilterRangePublished,
}

// FromCode ...
func FromCode(code commonsStatus.Code) *commonsStatus.Status {
	if v, ok := codeStatus[code]; ok {
		return v
	}
	return commonsStatus.FromCode(code)
}
