package program.freelance_marketplace.api.users.controller;

import lombok.AllArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import program.freelance_marketplace.api.users.dto.UserByIdDTO;
import program.freelance_marketplace.api.users.service.impl.UserServiceImpl;

import java.util.Optional;

@RestController
@RequestMapping("api/users")
@AllArgsConstructor
public class UserController {
    private UserServiceImpl userService;

    @GetMapping("/{id}")
    public ResponseEntity<UserByIdDTO> getUserByID(
            @PathVariable Long id,
            Pageable pageable
    ) {
         Optional<UserByIdDTO> optionalUserDTO =  userService.getUserById(id, pageable);
            return optionalUserDTO.map(ResponseEntity::ok)
                    .orElseGet(() -> ResponseEntity.notFound().build());
    }
}
