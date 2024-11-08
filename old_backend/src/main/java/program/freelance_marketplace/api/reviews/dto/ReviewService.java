package program.freelance_marketplace.api.reviews.dto;

public record ReviewService(
        Long id,
        Double price,
        String image,
        String title
) {
}
