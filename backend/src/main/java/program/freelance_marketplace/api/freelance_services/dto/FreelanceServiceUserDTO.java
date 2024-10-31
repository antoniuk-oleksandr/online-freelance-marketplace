package program.freelance_marketplace.api.freelance_services.dto;

import java.math.BigDecimal;

public record FreelanceServiceUserDTO (
    Long id,
    String firstName,
    String surname,
    String avatar,
    BigDecimal rating,
    Double level,
    Long reviewsCount
){}
