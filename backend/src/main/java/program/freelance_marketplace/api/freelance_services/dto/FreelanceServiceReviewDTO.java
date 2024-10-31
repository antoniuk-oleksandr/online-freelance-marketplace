package program.freelance_marketplace.api.freelance_services.dto;

import java.sql.Timestamp;

public record FreelanceServiceReviewDTO(
        Long id,
        String content,
        Integer rating,
        Timestamp createdAt,
        Timestamp endedAt,
        FreelanceServiceReviewCustomer customer,
        FreelanceReviewService service
) {
}
