package program.freelance_marketplace.api.users.service;

import org.springframework.data.domain.Pageable;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;

import java.util.Optional;

public interface UserService {
    Optional<UserByIdDTO> getUserById(Long id, Pageable pageable);
}
