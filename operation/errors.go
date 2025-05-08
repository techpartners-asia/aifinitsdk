package aifinitsdk_operation

import "fmt"

type ProductPriceUpdateError int

const (
	ErrProductPriceUpdateVendingMachineDoesNotExistTargetGoods int = 3501
	ErrProductPriceUpdateThereAreDuplicateProducts             int = 40502
	ErrProductPriceUpdateThereAreDownloadedGoods               int = 40504
	ErrProductPriceUpdateTheSelfDealerDoesNotExist             int = 40506
	ErrProductPriceUpdateThereAreUnknownProducts               int = 40507
	ErrProductPriceUpdateNoOperatingPermissions                int = 40531
)

func (e ProductPriceUpdateError) Error() string {
	return fmt.Sprintf("ProductPriceUpdateError: %d", e)
}

func ConvertProductPriceUpdateError(code int, message string) error {
	switch code {
	case ErrProductPriceUpdateVendingMachineDoesNotExistTargetGoods:
		return fmt.Errorf("ProductPriceUpdateVendingMachineDoesNotExistTargetGoods: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreDuplicateProducts:
		return fmt.Errorf("ProductPriceUpdateThereAreDuplicateProducts: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreDownloadedGoods:
		return fmt.Errorf("ProductPriceUpdateThereAreDownloadedGoods: %d, message: %s", code, message)
	case ErrProductPriceUpdateTheSelfDealerDoesNotExist:
		return fmt.Errorf("ProductPriceUpdateTheSelfDealerDoesNotExist: %d, message: %s", code, message)
	case ErrProductPriceUpdateThereAreUnknownProducts:
		return fmt.Errorf("ProductPriceUpdateThereAreUnknownProducts: %d, message: %s", code, message)
	case ErrProductPriceUpdateNoOperatingPermissions:
		return fmt.Errorf("ProductPriceUpdateNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("ProductPriceUpdateError: %d, message: %s", code, message)
	}
}

type GetOrderVideoError int

const (
	ErrGetOrderVideoSuccess                            int = 200
	ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound int = 404
	ErrGetOrderVideoTheOpeningRequestDoesNotExist      int = 42404
)

func (e GetOrderVideoError) Error() string {
	return fmt.Sprintf("GetOrderVideoError: %d", e)
}

func ConvertGetOrderVideoError(code int, message string) error {
	switch code {
	case ErrGetOrderVideoSuccess:
		return fmt.Errorf("GetOrderVideoSuccess: %d, message: %s", code, message)
	case ErrGetOrderVideoNoOrderOrReplenishmentRecordsFound:
		return fmt.Errorf("GetOrderVideoNoOrderOrReplenishmentRecordsFound: %d, message: %s", code, message)
	case ErrGetOrderVideoTheOpeningRequestDoesNotExist:
		return fmt.Errorf("GetOrderVideoTheOpeningRequestDoesNotExist: %d, message: %s", code, message)
	default:
		return fmt.Errorf("GetOrderVideoError: %d, message: %s", code, message)
	}
}

type SearchOpenDoorError int

const (
	ErrSearchOpenDoorSuccess                                            int = 201
	ErrSearchOpenDoorClosingSuccess                                     int = 202
	ErrSearchOpenDoorLastShoppingNotOver                                int = 2031
	ErrSearchOpenDoorLastReplenishmentNotOver                           int = 2032
	ErrSearchOpenDoorEquipmentPoweredOff                                int = 2033
	ErrSearchOpenDoorEquipmentInOperationMode                           int = 2034
	ErrSearchOpenDoorDoorOpenedFailed                                   int = 204
	ErrSearchOpenDoorDeviceBackgroundProcess                            int = 503
	ErrSearchOpenDoorDeviceReceivedMessageTimeout                       int = 504
	ErrSearchOpenDoorUnknownError                                       int = 505
	ErrSearchOpenDoorDebuggingInformationIncorrect                      int = 506
	ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed         int = 5050
	ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure            int = 5051
	ErrSearchOpenDoorEquipmentHeavyFaults                               int = 5052
	ErrSearchOpenDoorDeviceCameraAllDropped                             int = 5053
	ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal int = 5054
	ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual               int = 5055
	ErrSearchOpenDoorEquipmentPowerSupplyStatusError                    int = 5056
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen                      int = 5057
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed                    int = 5058
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2                   int = 5059
	ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3                   int = 5060
	ErrSearchOpenDoorEquipmentHasNotReportedTheResults                  int = 404
	ErrSearchOpenDoorDoorRequestIdDoesNotExist                          int = 42404
	ErrSearchOpenDoorTypeParameterTypeError                             int = 40005
	ErrSearchOpenDoorTooManyOrdersInTheShop                             int = 40526
	ErrSearchOpenDoorNoSearchPermissions                                int = 42403
)

func (e SearchOpenDoorError) Error() string {
	return fmt.Sprintf("SearchOpenDoorError: %d", e)
}

func ConvertSearchOpenDoorError(code int, message string) error {
	switch code {
	case ErrSearchOpenDoorSuccess:
		return fmt.Errorf("SearchOpenDoorSuccess: %d, message: %s", code, message)
	case ErrSearchOpenDoorClosingSuccess:
		return fmt.Errorf("SearchOpenDoorClosingSuccess: %d, message: %s", code, message)
	case ErrSearchOpenDoorLastShoppingNotOver:
		return fmt.Errorf("SearchOpenDoorLastShoppingNotOver: %d, message: %s", code, message)
	case ErrSearchOpenDoorLastReplenishmentNotOver:
		return fmt.Errorf("SearchOpenDoorLastReplenishmentNotOver: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentPoweredOff:
		return fmt.Errorf("SearchOpenDoorEquipmentPoweredOff: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentInOperationMode:
		return fmt.Errorf("SearchOpenDoorEquipmentInOperationMode: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorOpenedFailed:
		return fmt.Errorf("SearchOpenDoorDoorOpenedFailed: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceBackgroundProcess:
		return fmt.Errorf("SearchOpenDoorDeviceBackgroundProcess: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceReceivedMessageTimeout:
		return fmt.Errorf("SearchOpenDoorDeviceReceivedMessageTimeout: %d, message: %s", code, message)
	case ErrSearchOpenDoorUnknownError:
		return fmt.Errorf("SearchOpenDoorUnknownError: %d, message: %s", code, message)
	case ErrSearchOpenDoorDebuggingInformationIncorrect:
		return fmt.Errorf("SearchOpenDoorDebuggingInformationIncorrect: %d, message: %s", code, message)
	case ErrSearchOpenDoorVerificationOfSaleOfPlanningProductsFailed:
		return fmt.Errorf("SearchOpenDoorVerificationOfSaleOfPlanningProductsFailed: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedDoorOpeningEquipmentSerialFailure:
		return fmt.Errorf("SearchOpenDoorFailedDoorOpeningEquipmentSerialFailure: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentHeavyFaults:
		return fmt.Errorf("SearchOpenDoorEquipmentHeavyFaults: %d, message: %s", code, message)
	case ErrSearchOpenDoorDeviceCameraAllDropped:
		return fmt.Errorf("SearchOpenDoorDeviceCameraAllDropped: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal:
		return fmt.Errorf("SearchOpenDoorFailedDoorOpeningLocalRecognitionAlgorithmAbnormal: %d, message: %s", code, message)
	case ErrSearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual:
		return fmt.Errorf("SearchOpenDoorFailedToOpenTheDoorTheLockWasUnusual: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentPowerSupplyStatusError:
		return fmt.Errorf("SearchOpenDoorEquipmentPowerSupplyStatusError: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsOpen:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsOpen: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed2:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed2: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorLockAbnormalTheDoorIsClosed3:
		return fmt.Errorf("SearchOpenDoorDoorLockAbnormalTheDoorIsClosed3: %d, message: %s", code, message)
	case ErrSearchOpenDoorEquipmentHasNotReportedTheResults:
		return fmt.Errorf("SearchOpenDoorEquipmentHasNotReportedTheResults: %d, message: %s", code, message)
	case ErrSearchOpenDoorDoorRequestIdDoesNotExist:
		return fmt.Errorf("SearchOpenDoorDoorRequestIdDoesNotExist: %d, message: %s", code, message)
	case ErrSearchOpenDoorTypeParameterTypeError:
		return fmt.Errorf("SearchOpenDoorTypeParameterTypeError: %d, message: %s", code, message)
	case ErrSearchOpenDoorTooManyOrdersInTheShop:
		return fmt.Errorf("SearchOpenDoorTooManyOrdersInTheShop: %d, message: %s", code, message)
	case ErrSearchOpenDoorNoSearchPermissions:
		return fmt.Errorf("SearchOpenDoorNoSearchPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("SearchOpenDoorError: %d, message: %s", code, message)
	}
}

type UpdateSoldGoodsError int

const (
	ErrUpdateSoldGoodsTooManyGoods           int = 10004
	ErrUpdateSoldGoodsDuplicateGoods         int = 40502
	ErrUpdateSoldGoodsMutuallyExclusiveGoods int = 40503
	ErrUpdateSoldGoodsDownloadedGoods        int = 40504
	ErrUpdateSoldGoodsSelfDealerNotExist     int = 40506
	ErrUpdateSoldGoodsUnknownGoods           int = 40507
	ErrUpdateSoldGoodsNoOperatingPermissions int = 40531
)

func (e UpdateSoldGoodsError) Error() string {
	return fmt.Sprintf("UpdateSoldGoodsError: %d", e)
}

func ConvertUpdateSoldGoodsError(code int, message string) error {
	switch code {
	case ErrUpdateSoldGoodsTooManyGoods:
		return fmt.Errorf("UpdateSoldGoodsTooManyGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsDuplicateGoods:
		return fmt.Errorf("UpdateSoldGoodsDuplicateGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsMutuallyExclusiveGoods:
		return fmt.Errorf("UpdateSoldGoodsMutuallyExclusiveGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsDownloadedGoods:
		return fmt.Errorf("UpdateSoldGoodsDownloadedGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsSelfDealerNotExist:
		return fmt.Errorf("UpdateSoldGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsUnknownGoods:
		return fmt.Errorf("UpdateSoldGoodsUnknownGoods: %d, message: %s", code, message)
	case ErrUpdateSoldGoodsNoOperatingPermissions:
		return fmt.Errorf("UpdateSoldGoodsNoOperatingPermissions: %d, message: %s", code, message)
	default:
		return fmt.Errorf("UpdateSoldGoodsError: %d, message: %s", code, message)
	}
}

type UpdateSoldGoodsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type GetSoldGoodsError int

const (
	ErrGetSoldGoodsSelfDealerNotExist            int = 40506
	ErrGetSoldGoodsSelfDealerNotBelongToMerchant int = 40531
)

func (e GetSoldGoodsError) Error() string {
	return fmt.Sprintf("GetSoldGoodsError: %d", e)
}

func ConvertGetSoldGoodsError(code int, message string) error {
	switch code {
	case ErrGetSoldGoodsSelfDealerNotExist:
		return fmt.Errorf("GetSoldGoodsSelfDealerNotExist: %d, message: %s", code, message)
	case ErrGetSoldGoodsSelfDealerNotBelongToMerchant:
		return fmt.Errorf("GetSoldGoodsSelfDealerNotBelongToMerchant: %d, message: %s", code, message)
	default:
		return fmt.Errorf("GetSoldGoodsError: %d, message: %s", code, message)
	}
}

type OpenDoorError int

const (
	ErrOpenDoorSuccess                      int = 200
	ErrOpenDoorFailed                       int = 400
	ErrOpenDoorTimeout                      int = 503
	ErrOpenDoorUnusualMachinePackage        int = 3501
	ErrOpenDoorOfflineEquipment             int = 10416
	ErrOpenDoorSelfDealerNotInOperation     int = 40525
	ErrOpenDoorTooManyOrdersNotCompleted    int = 40526
	ErrOpenDoorNonBusinessSelfSellerMachine int = 40531
)

func (e OpenDoorError) Error() string {
	return fmt.Sprintf("OpenDoorError: %d", e)
}

func ConvertOpenDoorError(code int, message string) error {
	switch code {
	case ErrOpenDoorSuccess:
		return fmt.Errorf("OpenDoorSuccess: %d, message: %s", code, message)
	case ErrOpenDoorFailed:
		return fmt.Errorf("OpenDoorFailed: %d, message: %s", code, message)
	case ErrOpenDoorTimeout:
		return fmt.Errorf("OpenDoorTimeout: %d, message: %s", code, message)
	case ErrOpenDoorUnusualMachinePackage:
		return fmt.Errorf("OpenDoorUnusualMachinePackage: %d, message: %s", code, message)
	case ErrOpenDoorOfflineEquipment:
		return fmt.Errorf("OpenDoorOfflineEquipment: %d, message: %s", code, message)
	case ErrOpenDoorSelfDealerNotInOperation:
		return fmt.Errorf("OpenDoorSelfDealerNotInOperation: %d, message: %s", code, message)
	case ErrOpenDoorTooManyOrdersNotCompleted:
		return fmt.Errorf("OpenDoorTooManyOrdersNotCompleted: %d, message: %s", code, message)
	case ErrOpenDoorNonBusinessSelfSellerMachine:
		return fmt.Errorf("OpenDoorNonBusinessSelfSellerMachine: %d, message: %s", code, message)
	default:
		return fmt.Errorf("OpenDoorError: %d, message: %s", code, message)
	}
}
