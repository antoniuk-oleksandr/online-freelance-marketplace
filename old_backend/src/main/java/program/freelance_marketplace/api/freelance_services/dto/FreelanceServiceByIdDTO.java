package program.freelance_marketplace.api.freelance_services.dto;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.CategoryEntity;

import java.math.BigDecimal;
import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Builder
public class FreelanceServiceByIdDTO {
    private Long id;
    private String title;
    private Timestamp createdAt;
    private String description;
    private List<String> images;
    private CategoryEntity category;
    private Long reviewsCount;
    private BigDecimal rating;
    private List<FreelanceServiceReviewDTO> reviews;
    private FreelanceServiceUserDTO freelancer;
    private List<PackageDTO> packages;
}
