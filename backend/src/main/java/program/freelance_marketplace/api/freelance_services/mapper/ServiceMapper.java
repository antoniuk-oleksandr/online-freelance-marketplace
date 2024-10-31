package program.freelance_marketplace.api.freelance_services.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.orders.dto.CountRating;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;
import program.freelance_marketplace.api.users.dto.UserServiceDTO;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.stream.Collectors;

@Component
public class ServiceMapper {
    public Map<Long, CountRating> mapServiceReviewCountEntityToMap(List<Object[]> servicesReviewEntity) {
        return servicesReviewEntity.stream()
                .collect(Collectors.toMap(
                        row -> ((Number) row[0]).longValue(),
                        row -> new CountRating(
                                ((Number) row[1]).longValue(),
                                BigDecimal.valueOf(((Number) row[2]).doubleValue())
                        )
                ));
    }

    public Map<Long, Double> mapServiceMinPriceEntityToMap(List<Object[]> servicesMinPriceEntity) {
        return servicesMinPriceEntity.stream()
                .collect(Collectors.toMap(
                        row -> ((Number) row[0]).longValue(),
                        row -> ((Number) row[1]).doubleValue()
                ));
    }

    public List<UserServiceDTO> mapServiceEntityListToUserServiceDTOList(
            List<ServiceEntity> serviceEntities,
            Map<Long, CountRating> serviceReviewMap,
            Map<Long, Double> serviceMinPriceMap
    ) {
        CountRating defaultVal = new CountRating(0L, BigDecimal.ZERO);

        return serviceEntities.stream()
                .map(serviceEntity -> UserServiceDTO.builder()
                        .id(serviceEntity.getId())
                        .minPrice(serviceMinPriceMap.getOrDefault(serviceEntity.getId(), 0.0))
                        .reviewsCount(serviceReviewMap.getOrDefault(serviceEntity.getId(), defaultVal).count())
                        .rating(serviceReviewMap.getOrDefault(serviceEntity.getId(), defaultVal).rating())
                        .title(serviceEntity.getTitle())
                        .createdAt(serviceEntity.getCreatedAt())
                        .description(serviceEntity.getDescription())
                        .image(
                                Optional.ofNullable(serviceEntity.getFiles())
                                        .filter(files -> !files.isEmpty())
                                        .map(files -> files.getFirst().getName())
                                        .orElse(null)
                        )
                        .category(serviceEntity.getCategory())
                        .build())
                .collect(Collectors.toList());
    }

    public CountRating mapFreelanceCountRating(Object[][] freelanceCountRating) {
        Object rating = freelanceCountRating[0][1];

        return new CountRating(
                (Long) freelanceCountRating[0][0],
                rating == null ? null : BigDecimal.valueOf(((Number) rating).doubleValue())
        );
    }
}
