package program.freelance_marketplace.api.users.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import program.freelance_marketplace.api.users.entity.UserEntity;

@Repository
public interface UserRepository extends JpaRepository<UserEntity, Long> {
    @Query(value = "SELECT U.id, U.first_name, U.surname, F.name as avatar " +
            "FROM orders O " +
            "LEFT JOIN users U ON O.customer_id = U.id " +
            "LEFT JOIN files F ON U.avatar_id = F.id " +
            "WHERE O.service_id = :serviceId AND U.id IS NOT NULL",
            nativeQuery = true)
    Object[] findUsersByServiceId(@Param("serviceId") Long serviceId);
}
