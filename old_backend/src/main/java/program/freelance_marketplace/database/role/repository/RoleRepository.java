package program.freelance_marketplace.database.role.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import program.freelance_marketplace.database.role.entity.RoleEntity;

import java.util.Optional;

public interface RoleRepository extends JpaRepository<RoleEntity, Long> {
    Optional<RoleEntity> findByName(String name);

    Boolean existsByName(String name);
}
