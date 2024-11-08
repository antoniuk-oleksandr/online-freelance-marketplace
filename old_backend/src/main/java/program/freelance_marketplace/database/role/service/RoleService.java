package program.freelance_marketplace.database.role.service;

import program.freelance_marketplace.database.role.entity.RoleEntity;

public interface RoleService {
    RoleEntity save(RoleEntity role);

    RoleEntity findByName(String name);

    Boolean existsByName(String name);
}
