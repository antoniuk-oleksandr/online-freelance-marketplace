package program.freelance_marketplace.api.users.service.impl;

import lombok.AllArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;
import program.freelance_marketplace.api.freelance_services.mapper.ServiceMapper;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.dto.UserServiceDTO;
import program.freelance_marketplace.api.reviews.dto.UserByIdReviewDTO;
import program.freelance_marketplace.api.users.entity.UserEntity;
import program.freelance_marketplace.api.users.mapper.UserMapper;
import program.freelance_marketplace.api.users.repository.UserRepository;
import program.freelance_marketplace.api.users.service.UserService;
import program.freelance_marketplace.api.users.utils.UserUtils;

import java.util.List;
import java.util.Map;
import java.util.Optional;

@Service
@AllArgsConstructor
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final UserMapper userMapper;
    private final ServiceMapper serviceMapper;
    private final UserUtils userUtils;

    @Override
    public Optional<UserByIdDTO> getUserById(Long id, Pageable pageable) {
        return userRepository.findById(id)
                .map(userEntity -> createUserByIdDTO(userEntity, pageable));
    }

    @Override
    public Optional<UserEntity> findByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    private UserByIdDTO createUserByIdDTO(UserEntity userEntity, Pageable pageable) {
        CountRating reviewCountRating = userUtils.getReviewCountRatingByFreelancerId(userEntity.getId());
        List<UserByIdReviewDTO> reviews = userUtils.getReviewsByUserId(userEntity.getId(), pageable);
        List<ServiceEntity> serviceEntities = userUtils.getServicesByFreelancerId(userEntity.getId());

        Map<Long, CountRating> serviceReviewMap = userUtils.getServiceReviewMap(serviceEntities);
        Map<Long, Double> serviceMinPriceMap = userUtils.getServiceMinPriceMap(serviceEntities);

        List<UserServiceDTO> userServiceDTOs = serviceMapper.mapServiceEntityListToUserServiceDTOList(
                serviceEntities, serviceReviewMap, serviceMinPriceMap
        );

        return userMapper.mapUserByIdEntityToDTO(userEntity, reviewCountRating, reviews, userServiceDTOs);
    }
}
