package impl

import (
	"context"
	"fmt"

	"github.com/devnica/EasyStore/commons"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/models/response"
	"github.com/devnica/EasyStore/repositories"
	"github.com/devnica/EasyStore/services"
)

type backofficeServiceImpl struct {
	repositories.AdminCommitRepository
}

func NewBackofficeServiceImpl(repo *repositories.AdminCommitRepository) services.BackofficeService {
	return &backofficeServiceImpl{AdminCommitRepository: *repo}
}

func (srv *backofficeServiceImpl) GetKYCReview(ctx context.Context, reviewStatus string) []response.KYCReviewResponseModel {

	fmt.Println(reviewStatus)
	/* mapa de estados de revision del KYC*/
	reviewStatusMap := commons.GetReviewStatusFromDictionary()
	/* se recupera el Id del status, por medio de la clave recibida en la peticion*/
	reviewStatusId := commons.GetKeyId(reviewStatus, reviewStatusMap)

	result, err := srv.AdminCommitRepository.FetchAllKYCReview(reviewStatusId)
	exceptions.PanicLogging(err)

	/* los resultados de la consulta al repositorio se mapean para ajustarlos a los datos que retornara el servicio*/
	var list []response.KYCReviewResponseModel
	for _, kyc := range result {

		list = append(list, response.KYCReviewResponseModel{
			ReviewId:      kyc.ReviewId,
			UserRef:       kyc.UserRef,
			AdminRef:      kyc.AdminRef,
			PreRevStatus:  kyc.PreRevStatus,
			PostRevStatus: kyc.PostRevStatus,
			IsAsigned:     kyc.IsAsigned,
			CreatedAt:     kyc.CreatedAt,
			UpdatedAt:     kyc.UpdatedAt,
		})
	}

	return list
}
