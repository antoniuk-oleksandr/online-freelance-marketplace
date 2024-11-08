package program.freelance_marketplace.database.role.service.impl;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import program.freelance_marketplace.database.role.entity.RoleEntity;
import program.freelance_marketplace.database.role.repository.RoleRepository;
import program.freelance_marketplace.database.role.service.RoleService;

@Service
@RequiredArgsConstructor
public class RoleServiceImpl implements RoleService {
    private final RoleRepository roleRepository;

    @Override
    public Boolean existsByName(String name) {
        return this.roleRepository.existsByName(name);
    }

    @Override
    public RoleEntity save(RoleEntity role) {
        return existsByName(role.getName()) ? role : this.roleRepository.save(role);
    }

    @Override
    public RoleEntity findByName(String name) {
        return this.roleRepository.findByName(name).orElse(null);
    }
}
