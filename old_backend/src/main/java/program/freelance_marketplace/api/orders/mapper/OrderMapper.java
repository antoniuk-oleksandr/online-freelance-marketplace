package program.freelance_marketplace.api.orders.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.orders.dto.CountRating;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.List;

@Component
public class OrderMapper {
    public CountRating mapReviewRatingCount(List<Object[]> reviewCount) {
        return new CountRating(
                (Long) reviewCount.getFirst()[0],
                ((BigDecimal) reviewCount.getFirst()[1]).setScale(2, RoundingMode.HALF_UP)
        );
    }
}
