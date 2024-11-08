package program.freelance_marketplace.api.freelance_services.service;

import org.springframework.data.domain.Pageable;
import program.freelance_marketplace.api.freelance_services.dto.FreelanceServiceByIdDTO;

public interface FreelanceService {
    FreelanceServiceByIdDTO getFreelanceServiceById(Long id, Pageable pageable);
}
