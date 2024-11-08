package program.freelance_marketplace.api.users.dto;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.reviews.dto.UserByIdReviewDTO;

import java.math.BigDecimal;
import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Builder
public class UserByIdDTO {
    private Long id;
    private String firstName;
    private String surname;
    private String avatar;
    private BigDecimal rating;
    private Double level;
    private Long reviewsCount;
    private Timestamp createdAt;
    private String about;
    private List<String> skills;
    private List<UserByIdReviewDTO> reviews;
    private List<UserServiceDTO> services;
    private List<String> languages;
}
