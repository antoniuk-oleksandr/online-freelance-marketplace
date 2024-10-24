package program.freelance_marketplace.api.reviews.dto;

public record ReviewUserDTO(
        Long id,
        String firstName,
        String surname,
        String avatar
) {
}
