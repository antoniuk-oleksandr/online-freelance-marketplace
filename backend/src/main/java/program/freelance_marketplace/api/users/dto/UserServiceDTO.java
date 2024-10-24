package program.freelance_marketplace.api.users.dto;

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
public class UserServiceDTO {
    private Long id;
    private String title;
    private Timestamp createdAt;
    private String description;
    private String image;
    private CategoryEntity category;
    private List<String> files;
    private Double minPrice;
    private Long reviewsCount;
    private BigDecimal rating;
}
