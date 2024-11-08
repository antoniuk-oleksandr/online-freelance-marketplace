package program.freelance_marketplace.api.freelance_services.dto;

import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Builder
public class PackageDTO {
    Long id;
    String title;
    String description;
    Double price;
    Integer deliveryDays;
}
