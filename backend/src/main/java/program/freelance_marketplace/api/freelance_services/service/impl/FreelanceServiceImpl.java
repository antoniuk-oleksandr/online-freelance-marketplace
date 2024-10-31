package program.freelance_marketplace.api.freelance_services.service.impl;

import lombok.AllArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.freelance_services.dto.*;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;
import program.freelance_marketplace.api.freelance_services.mapper.FreelanceMapper;
import program.freelance_marketplace.api.freelance_services.mapper.ServiceMapper;
import program.freelance_marketplace.api.freelance_services.repository.ServiceRepository;
import program.freelance_marketplace.api.freelance_services.service.FreelanceService;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.orders.mapper.OrderMapper;
import program.freelance_marketplace.api.orders.repository.OrderRepository;
import program.freelance_marketplace.api.reviews.mapper.ReviewMapper;
import program.freelance_marketplace.api.reviews.repository.ReviewRepository;
import program.freelance_marketplace.api.users.mapper.UserMapper;
import program.freelance_marketplace.api.users.repository.UserRepository;

import java.util.List;
import java.util.Optional;

@Service
@AllArgsConstructor
public class FreelanceServiceImpl implements FreelanceService {
    private final ReviewRepository reviewRepository;
    private final ServiceRepository serviceRepository;
    private final UserRepository userRepository;
    private final OrderRepository orderRepository;
    private final ServiceMapper serviceMapper;
    private final FreelanceMapper freelanceMapper;
    private final ReviewMapper reviewMapper;
    private final UserMapper userMapper;
    private final OrderMapper orderMapper;

    @Override
    public FreelanceServiceByIdDTO getFreelanceServiceById(Long id, Pageable pageable) {
        Optional<ServiceEntity> serviceEntity = serviceRepository.findById(id);
        List<Object[]> freelanceServiceReviewObjects =
                reviewRepository.findReviewsByServiceId(id, pageable).getContent();
        Object[] customerObjects = userRepository.findUsersByServiceId(id);
        List<FreelanceServiceReviewCustomer> customers =
                userMapper.mapCustomerObjectsToDTOList(customerObjects);
        List<FreelanceServiceReviewDTO> freelanceServiceReviewDTOList =
                reviewMapper.mapFreelanceServiceReviewEntitiesToDTO(
                        freelanceServiceReviewObjects,
                        customers
                );
        CountRating countRating = serviceMapper.mapFreelanceCountRating(
                serviceRepository.getReviewCountRatingByServiceId(id)
        );

        return serviceEntity.map(entity -> FreelanceServiceByIdDTO.builder()
                .id(entity.getId())
                .images(entity.getFiles().stream().map(FileEntity::getName).toList())
                .title(entity.getTitle())
                .description(entity.getDescription())
                .createdAt(entity.getCreatedAt())
                .category(entity.getCategory())
                .freelancer(freelanceMapper.mapFreelancerEntityToFreelancerDTO(
                        entity.getFreelancer(),
                        orderMapper.mapReviewRatingCount(
                                orderRepository.getReviewCountRatingByFreelancerId(entity.getFreelancer().getId())
                        )
                ))
                .reviewsCount(countRating.count())
                .rating(countRating.rating())
                .packages(entity.getPackages().stream()
                        .map(item -> PackageDTO.builder()
                                .id(item.getId())
                                .title(item.getTitle())
                                .description(item.getDescription())
                                .price(item.getPrice())
                                .deliveryDays(item.getDeliveryDays())
                                .build())
                        .toList()
                )
                .reviews(freelanceServiceReviewDTOList)
                .build()).orElse(null);

    }
}
