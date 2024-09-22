package program.freelance_marketplace.api.users.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.ServiceEntity;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.dto.UserServiceDTO;
import program.freelance_marketplace.api.users.entity.SkillEntity;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.util.List;
import java.util.stream.Collectors;

@Component
public class UserMapper {
    public UserByIdDTO mapUserByIdEntityToDTO(UserEntity userEntity) {
        return UserByIdDTO.builder()
                .id(userEntity.getId())
                .firstName(userEntity.getFirstName())
                .surname(userEntity.getSurname())
                .rating(userEntity.getRating())
                .level(userEntity.getLevel())
                .createdAt(userEntity.getCreatedAt())
                .about(userEntity.getAbout())
                .skills(mapSkillEntityListToStringList(userEntity.getSkills()))
                .build();
    }

    public List<String> mapSkillEntityListToStringList(List<SkillEntity> skillEntities) {
        return skillEntities.stream()
                .map(SkillEntity::getName)
                .collect(Collectors.toList());
    }

    public List<UserServiceDTO> mapServiceEntityListToUserServiceDTOList(List<ServiceEntity> serviceEntities) {
        return serviceEntities.stream()
                .map(serviceEntity -> UserServiceDTO.builder()
                        .id(serviceEntity.getId())
                        .title(serviceEntity.getTitle())
                        .createdAt(serviceEntity.getCreated_at())
                        .description(serviceEntity.getDescription())
                        .image(serviceEntity.getImage().getName())
                        .category(serviceEntity.getCategory())
                        .languages(serviceEntity.getLanguages().stream()
                                .map(FileEntity::getName)
                                .collect(Collectors.toList()))
                        .build())
                .collect(Collectors.toList());
    }
}
