package program.freelance_marketplace.api.users.dto;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.ReviewEntity;
import program.freelance_marketplace.api.ServiceEntity;

import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Builder
public class UserByIdDTO {
    private Long id;
    private String firstName;
    private String surname;
    private Double rating;
    private Double level;
    private int reviewsCount;
    private Timestamp createdAt;
    private String about;
    private List<String> skills;
    private List<ReviewEntity> reviews;
    private List<UserServiceDTO> services;
}
