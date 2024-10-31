package program.freelance_marketplace.api.users.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceReviewCustomer;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.dto.UserServiceDTO;
import program.freelance_marketplace.api.users.entity.LanguageEntity;
import program.freelance_marketplace.api.users.entity.SkillEntity;
import program.freelance_marketplace.api.reviews.dto.UserByIdReviewDTO;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

@Component
public class UserMapper {
    public UserByIdDTO mapUserByIdEntityToDTO(
            UserEntity userEntity,
            CountRating reviewCountRating,
            List<UserByIdReviewDTO> reviewEntities,
            List<UserServiceDTO> userServiceDTOs
    ) {
        return UserByIdDTO.builder()
                .id(userEntity.getId())
                .avatar(mapFileEntityToString(userEntity.getAvatar()))
                .reviewsCount(reviewCountRating.count())
                .rating(reviewCountRating.rating())
                .reviews(reviewEntities)
                .firstName(userEntity.getFirstName())
                .surname(userEntity.getSurname())
                .level(userEntity.getLevel())
                .createdAt(userEntity.getCreatedAt())
                .about(userEntity.getAbout())
                .services(userServiceDTOs)
                .skills(mapSkillEntityListToStringList(userEntity.getSkills()))
                .languages(mapLanguageEntityListToStringList(userEntity.getLanguages()))
                .build();
    }

    private String mapFileEntityToString(FileEntity fileEntity) {
        if (fileEntity == null) return null;
        else return fileEntity.getName();
    }

    private List<String> mapSkillEntityListToStringList(List<SkillEntity> skillEntities) {
        return skillEntities.stream()
                .map(SkillEntity::getName)
                .collect(Collectors.toList());
    }

    public List<String> mapLanguageEntityListToStringList(List<LanguageEntity> languageEntities) {
        return languageEntities.stream()
                .map(LanguageEntity::getName)
                .collect(Collectors.toList());
    }

    public List<FreelanceServiceReviewCustomer> mapCustomerObjectsToDTOList(
            Object[] customerObjects
    ) {
        return Arrays.stream(customerObjects)
                .map(item -> {
                    Object[] customerData = (Object[]) item;
                    return new FreelanceServiceReviewCustomer(
                            (Long) customerData[0],
                            (String) customerData[1],
                            (String) customerData[2],
                            customerData[3] == null ? null : (String) customerData[3]
                    );
                })
                .toList();
    }

}
