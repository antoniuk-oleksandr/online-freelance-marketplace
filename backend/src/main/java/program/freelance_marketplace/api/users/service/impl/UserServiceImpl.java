package program.freelance_marketplace.api.users.service.impl;

import lombok.AllArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import program.freelance_marketplace.api.ReviewEntity;
import program.freelance_marketplace.api.ServiceEntity;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.entity.UserEntity;
import program.freelance_marketplace.api.users.mapper.UserMapper;
import program.freelance_marketplace.api.users.repository.ReviewRepository;
import program.freelance_marketplace.api.users.repository.ServiceRepository;
import program.freelance_marketplace.api.users.repository.UserRepository;
import program.freelance_marketplace.api.users.service.UserService;

import java.util.List;
import java.util.Optional;

@Service
@AllArgsConstructor
public class UserServiceImpl implements UserService {
    private final UserRepository userRepository;
    private final ReviewRepository reviewRepository;
    private final ServiceRepository serviceRepository;
    private final UserMapper userMapper;

    @Override
    public Optional<UserByIdDTO> getUserById(Long id, Pageable pageable) {
        Optional<UserEntity> entity = userRepository.findById(id);
        Page<ReviewEntity> reviewEntityPage = reviewRepository.findReviewsByUserId(id, pageable);
        List<ServiceEntity> serviceEntities = serviceRepository.getServicesByFreelancerId(id);
        return entity.map(userMapper::mapUserByIdEntityToDTO)
                .map(userByIdDTO -> {
                    userByIdDTO.setReviews(reviewEntityPage.getContent());
                    userByIdDTO.setServices(userMapper
                            .mapServiceEntityListToUserServiceDTOList(serviceEntities));
                    return userByIdDTO;
                });
    }
}
