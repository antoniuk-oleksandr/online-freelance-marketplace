package program.freelance_marketplace.api.users.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class PackageDTO {
    private Long id;
    private String title;
    private String description;
    private double price;
    private int deliveryDays;
    private String image;
}
