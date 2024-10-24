package program.freelance_marketplace.api.reviews.dto;

import program.freelance_marketplace.api.users.dto.UserByIdDTO;

import java.sql.Timestamp;

public record UserByIdReviewDTO(
        Long id,
        Integer rating,
        String content,
        Timestamp createdAt,
        Timestamp endedAt,
        ReviewUserDTO customer,
        ReviewService service
){}