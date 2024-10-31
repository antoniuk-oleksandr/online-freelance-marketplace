package program.freelance_marketplace.api.freelance_services.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceUserDTO;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.users.entity.UserEntity;

@Component
public class FreelanceMapper {
    public FreelanceServiceUserDTO mapFreelancerEntityToFreelancerDTO(
            UserEntity userEntity,
            CountRating countRating
    ) {
        String avatar = userEntity.getAvatar() == null
                ? null
                : userEntity.getAvatar().getName();

        return new FreelanceServiceUserDTO(
                userEntity.getId(),
                userEntity.getFirstName(),
                userEntity.getSurname(),
                avatar,
                countRating.rating(),
                userEntity.getLevel(),
                countRating.count()
        );
    }
}
