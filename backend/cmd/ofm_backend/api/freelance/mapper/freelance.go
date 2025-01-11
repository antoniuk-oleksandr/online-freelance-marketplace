package mapper

import (
	"ofm_backend/cmd/ofm_backend/api/freelance/dto"
	"ofm_backend/cmd/ofm_backend/api/freelance/model"
)

func MapFreelanceModelToDTO(
	freelanceByID *model.FreelanceByID,
	reviews *[]model.Review,
) dto.Freelance {
	var packages []dto.Package

	for _, p := range *freelanceByID.Packages {
		packages = append(packages, dto.Package{
			ID:           p.ID,
			DeliveryDays: p.DeliveryDays,
			Description:  p.Description,
			Price:        p.Price,
			Title:        p.Title,
		})
	}

	return dto.Freelance{
		ID:           freelanceByID.ID,
		CreatedAt:    freelanceByID.CreatedAt,
		Description:  freelanceByID.Description,
		ReviewsCount: freelanceByID.ReviewsCount,
		Rating:       freelanceByID.Rating,
		Title:        freelanceByID.Title,
		Images:       freelanceByID.Images,
		Category:     freelanceByID.Category,
		Reviews:      reviews,
		Freelancer: &dto.FreelanceServiceFreelancer{
			ID:           freelanceByID.Freelancer.ID,
			Username:     freelanceByID.Freelancer.Username,
			FirstName:    freelanceByID.Freelancer.FirstName,
			Surname:      freelanceByID.Freelancer.Surname,
			Avatar:       freelanceByID.Freelancer.Avatar,
			Rating:       freelanceByID.Freelancer.Rating,
			Level:        freelanceByID.Freelancer.Level,
			ReviewsCount: freelanceByID.Freelancer.ReviewsCount,
		},
		Packages: &packages,
	}
}

func MapRestrictedFreelanceModelToDto(
	freelanceByIdRestrictedModel model.FreelanceByIdRestricted,
) dto.FreelanceByIdRestricted {
	var packages []dto.Package

	for _, p := range *freelanceByIdRestrictedModel.Packages {
		packages = append(packages, dto.Package{
			ID:           p.ID,
			DeliveryDays: p.DeliveryDays,
			Description:  p.Description,
			Price:        p.Price,
			Title:        p.Title,
		})
	}
	
	return dto.FreelanceByIdRestricted{
		Id: freelanceByIdRestrictedModel.Id,
		ReviewsCount: freelanceByIdRestrictedModel.ReviewsCount,
		Rating: freelanceByIdRestrictedModel.Rating,
		Title: freelanceByIdRestrictedModel.Title,
		Image: freelanceByIdRestrictedModel.Image,
		Packages: &packages,
	}
}
