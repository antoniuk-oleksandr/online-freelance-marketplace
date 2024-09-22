package program.freelance_marketplace.api.users.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.ServiceEntity;

import java.util.List;

@Repository
public interface ServiceRepository extends JpaRepository<ServiceEntity, Long> {
    @Query("SELECT s FROM ServiceEntity s WHERE s.freelancer.id = :userId")
    List<ServiceEntity> getServicesByFreelancerId(Long userId);
}
