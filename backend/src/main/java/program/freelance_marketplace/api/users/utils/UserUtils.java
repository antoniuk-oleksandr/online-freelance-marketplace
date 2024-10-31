package program.freelance_marketplace.api.users.utils;

import lombok.AllArgsConstructor;
import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.orders.mapper.OrderMapper;
import program.freelance_marketplace.api.orders.repository.OrderRepository;
import program.freelance_marketplace.api.reviews.mapper.ReviewMapper;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;
import program.freelance_marketplace.api.freelance_services.mapper.ServiceMapper;
import program.freelance_marketplace.api.reviews.dto.UserByIdReviewDTO;
import program.freelance_marketplace.api.reviews.repository.ReviewRepository;
import program.freelance_marketplace.api.freelance_services.repository.ServiceRepository;

import org.springframework.data.domain.Pageable;

import java.util.List;
import java.util.Map;

@Component
@AllArgsConstructor
public class UserUtils {
    private final OrderMapper orderMapper;
    private final ReviewRepository reviewRepository;
    private final ServiceRepository serviceRepository;
    private final ServiceMapper serviceMapper;
    private final OrderRepository orderRepository;
    private final ReviewMapper reviewMapper;

    public CountRating getReviewCountRatingByFreelancerId(Long freelancerId) {
        return orderMapper.mapReviewRatingCount(orderRepository.getReviewCountRatingByFreelancerId(freelancerId));
    }

    public List<UserByIdReviewDTO> getReviewsByUserId(Long userId, Pageable pageable) {
        return reviewMapper.mapUserReviewEntitiesToDTO(
                reviewRepository.findReviewsByUserId(userId, pageable).getContent()
        );
    }

    public List<ServiceEntity> getServicesByFreelancerId(Long freelancerId) {
        return serviceRepository.getServicesByFreelancerId(freelancerId);
    }

    public Map<Long, CountRating> getServiceReviewMap(List<ServiceEntity> serviceEntities) {
        List<Long> ids = extractServiceIds(serviceEntities);
        return serviceMapper.mapServiceReviewCountEntityToMap(serviceRepository.getServiceReviewCount(ids));
    }

    public Map<Long, Double> getServiceMinPriceMap(List<ServiceEntity> serviceEntities) {
        List<Long> ids = extractServiceIds(serviceEntities);
        return serviceMapper.mapServiceMinPriceEntityToMap(serviceRepository.getServiceMinPrice(ids));
    }

    private List<Long> extractServiceIds(List<ServiceEntity> serviceEntities) {
        return serviceEntities.stream()
                .map(ServiceEntity::getId)
                .toList();
    }
}
