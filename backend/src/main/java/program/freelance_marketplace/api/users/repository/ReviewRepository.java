package program.freelance_marketplace.api.users.repository;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.ReviewEntity;

@Repository
public interface ReviewRepository extends JpaRepository<ReviewEntity, Long> {

    @Query(value = "SELECT R.* FROM orders O " +
            "JOIN reviews R ON O.review_id = R.id " +
            "WHERE O.freelancer_id = :userId",
            nativeQuery = true)
    Page<ReviewEntity> findReviewsByUserId(@Param("userId") Long userId, Pageable pageable);
}
