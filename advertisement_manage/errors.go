package advertisementmanage

import "fmt"

const (
	ErrCodeSuccess   = 200
	ErrCodeNotFound  = 404
	ErrCodeForbidden = 403
)

const (
	ErrCodeSourceMaterialNotFound     = 4440
	ErrCodeSourceMaterialNotAllowed   = 4441
	ErrCodeSourceMaterialDoesNotExist = 4448
)

const (
	ErrCodeAdvertisementNotFound     = 4450
	ErrCodeAdvertisementNotAllowed   = 4451
	ErrCodeAdvertisementInvalidInput = 4452
)

const (
	ErrAdDetailNotFound   = 4440
	ErrAdDetailNotAllowed = 4441
)

type ErrAdDetail int

func (e ErrAdDetail) Error() string {
	return fmt.Sprintf("ErrAdDetail: %d", e)
}

func ConvertAdDetailError(code int, message string) error {
	switch code {
	case ErrAdDetailNotFound:
		return fmt.Errorf("AdDetailNotFound: %d, message: %s", code, message)
	case ErrAdDetailNotAllowed:
		return fmt.Errorf("AdDetailNotAllowed: %d, message: %s", code, message)
	default:
		return fmt.Errorf("AdDetailError: %d, message: %s", code, message)
	}
}

type SourceMaterialError int

func (e SourceMaterialError) Error() string {
	return fmt.Sprintf("SourceMaterialError: %d", e)
}

func ConvertSourceMaterialError(code int, message string) error {
	switch code {
	case ErrCodeSourceMaterialNotFound:
		return fmt.Errorf("SourceMaterialNotFound: %d, message: %s", code, message)
	case ErrCodeSourceMaterialNotAllowed:
		return fmt.Errorf("SourceMaterialNotAllowed: %d, message: %s", code, message)
	case ErrCodeSourceMaterialDoesNotExist:
		return fmt.Errorf("SourceMaterialDoesNotExist: %d, message: %s", code, message)
	default:
		return fmt.Errorf("SourceMaterialError: %d, message: %s", code, message)
	}
}

type AdvertisementError int

func (e AdvertisementError) Error() string {
	return fmt.Sprintf("AdvertisementError: %d", e)
}

func ConvertAdvertisementError(code int, message string) error {
	switch code {
	case ErrCodeAdvertisementNotFound:
		return fmt.Errorf("AdvertisementNotFound: %d, message: %s", code, message)
	case ErrCodeAdvertisementNotAllowed:
		return fmt.Errorf("AdvertisementNotAllowed: %d, message: %s", code, message)
	case ErrCodeAdvertisementInvalidInput:
		return fmt.Errorf("AdvertisementInvalidInput: %d, message: %s", code, message)
	default:
		return fmt.Errorf("AdvertisementError: %d, message: %s", code, message)
	}
}
