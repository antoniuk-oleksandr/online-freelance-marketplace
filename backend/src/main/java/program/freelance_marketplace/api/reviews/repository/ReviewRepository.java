package program.freelance_marketplace.api.reviews.repository;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.reviews.entity.ReviewEntity;

@Repository
public interface ReviewRepository extends JpaRepository<ReviewEntity, Long> {
    @Query(value = "SELECT R.*, " +
            "       O.created_at, " +
            "       O.ended_at, " +
            "       U.id   AS user_id, " +
            "       U.first_name, " +
            "       U.surname, " +
            "       F.name AS avatar, " +
            "       S.id   AS service_id, " +
            "       P.price, " +
            "       F1.name as service_image, " +
            "       S.title " +
            "FROM orders O " +
            "         JOIN reviews R ON O.review_id = R.id " +
            "         LEFT JOIN users U ON U.id = O.customer_id " +
            "         LEFT JOIN files F ON U.avatar_id = F.id " +
            "         LEFT JOIN services S ON O.service_id = S.id " +
            "         LEFT JOIN packages P on P.id = O.service_package_id " +
            "         LEFT JOIN files F1 on S.image_id = F1.id " +
            "WHERE O.freelancer_id = :userId " +
            "ORDER BY O.ended_at DESC",
            nativeQuery = true)
    Page<Object[]> findReviewsByUserId(@Param("userId") Long userId, Pageable pageable);
}
