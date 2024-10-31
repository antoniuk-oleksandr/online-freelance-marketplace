package program.freelance_marketplace.api.freelance_services.controller;

import lombok.AllArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceByIdDTO;
import program.freelance_marketplace.api.freelance_services.service.FreelanceService;


@RestController
@RequestMapping("api/services")
@AllArgsConstructor
public class FreelanceServiceController {
    private FreelanceService freelanceService;

    @GetMapping("/{id}")
    public ResponseEntity<FreelanceServiceByIdDTO> getServiceById(
            @PathVariable Long id,
            Pageable pageable
    ) {
        FreelanceServiceByIdDTO serviceById = freelanceService.getFreelanceServiceById(id, pageable);
        if (serviceById == null) {
            return ResponseEntity.notFound().build();
        } else {
            return ResponseEntity.ok(serviceById);
        }
    }
}
