package program.freelance_marketplace.api.orders.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.orders.entity.OrderEntity;

import java.util.List;

@Repository
public interface OrderRepository extends JpaRepository<OrderEntity, Long> {
    @Query(value = "SELECT COUNT(R.id), AVG(R.rating) AS rating " +
            "FROM orders O " +
            "LEFT JOIN reviews R ON O.review_id = R.id " +
            "WHERE freelancer_id = :id AND review_id IS NOT NULL", nativeQuery = true)
    List<Object[]> getReviewCountRatingByFreelancerId(@Param("id" ) Long freelancerId);
}
