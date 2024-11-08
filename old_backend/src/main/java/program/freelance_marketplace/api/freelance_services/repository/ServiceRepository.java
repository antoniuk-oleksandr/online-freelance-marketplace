package program.freelance_marketplace.api.freelance_services.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;

import java.util.List;

@Repository
public interface ServiceRepository extends JpaRepository<ServiceEntity, Long> {
    @Query("SELECT s FROM ServiceEntity s WHERE s.freelancer.id = :userId")
    List<ServiceEntity> getServicesByFreelancerId(Long userId);

    @Query(value = "SELECT S.id, COUNT(R.id), AVG(R.rating) AS rating FROM services S " +
            "LEFT JOIN orders O ON S.id = O.service_id " +
            "LEFT JOIN reviews R ON O.review_id = R.id " +
            "WHERE S.id IN (:ids) AND R.id IS NOT NULL " +
            "GROUP BY S.id", nativeQuery = true)
    List<Object[]> getServiceReviewCount(@Param("ids") List<Long> ids);

    @Query(value = "SELECT SP.service_id AS id, MIN(P.price) " +
            "FROM services_packages SP " +
            "LEFT JOIN packages P ON P.id = SP.package_id " +
            "WHERE SP.service_id IN (:ids) " +
            "GROUP BY SP.service_id; ", nativeQuery = true)
    List<Object[]> getServiceMinPrice(@Param("ids") List<Long> ids);

    @Query(value = "SELECT COUNT(R.id), AVG(R.rating) FROM orders O " +
            "LEFT JOIN reviews R ON O.review_id = R.id " +
            "WHERE o.service_id = :serviceId AND R.id IS NOT NULL",
            nativeQuery = true)
    Object[][] getReviewCountRatingByServiceId(@Param("serviceId") Long serviceId);
}
