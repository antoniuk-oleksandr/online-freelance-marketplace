package program.freelance_marketplace.api.users.service;

import org.springframework.data.domain.Pageable;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.util.Optional;

public interface UserService {
    Optional<UserByIdDTO> getUserById(Long id, Pageable pageable);

    Optional<UserEntity> findByUsername(String username);
}
