package program.freelance_marketplace.api.reviews.mapper;

import org.springframework.stereotype.Component;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceReviewService;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceReviewCustomer;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceReviewDTO;
import program.freelance_marketplace.api.reviews.dto.ReviewService;
import program.freelance_marketplace.api.reviews.dto.ReviewUserDTO;
import program.freelance_marketplace.api.reviews.dto.UserByIdReviewDTO;

import java.math.BigDecimal;
import java.sql.Timestamp;
import java.util.List;
import java.util.stream.IntStream;

@Component
public class ReviewMapper {
    public List<UserByIdReviewDTO> mapUserReviewEntitiesToDTO(List<Object[]> userReviewEntities) {
        return userReviewEntities.stream()
                .map(item -> new UserByIdReviewDTO(
                        (Long) item[0],
                        (Integer) item[2],
                        (String) item[1],
                        (Timestamp) item[3],
                        (Timestamp) item[4],
                        new ReviewUserDTO(
                                (Long) item[5],
                                (String) item[6],
                                (String) item[7],
                                item[8] == null ? null : (String) item[8]
                        ),
                        new ReviewService(
                                (Long) item[9],
                                (Double) item[10],
                                (String) item[11],
                                (String) item[12]
                        )
                ))
                .toList();
    }

    public List<FreelanceServiceReviewDTO> mapFreelanceServiceReviewEntitiesToDTO(
            List<Object[]> freelanceServiceReviewEntities,
            List<FreelanceServiceReviewCustomer> customers
    ) {
        return IntStream.range(0, freelanceServiceReviewEntities.size())
                .mapToObj(index -> {
                    Object[] item = freelanceServiceReviewEntities.get(index);
                    return new FreelanceServiceReviewDTO(
                            (Long) item[0],
                            (String) item[1],
                            (Integer) item[2],
                            (Timestamp) item[3],
                            (Timestamp) item[4],
                            customers.get(index),
                            new FreelanceReviewService(
                                    BigDecimal.valueOf(((Number) item[5]).doubleValue())
                            )
                    );
                })
                .toList();
    }
}
