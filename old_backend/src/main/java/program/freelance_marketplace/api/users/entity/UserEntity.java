package program.freelance_marketplace.api.users.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.database.role.entity.RoleEntity;

import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "users")
public class UserEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String firstName;
    private String surname;
    private String email;
    private Double level;
    private Timestamp createdAt;

    @OneToMany(orphanRemoval = true)
    @JoinTable(
            name = "users_skills",
            joinColumns = @JoinColumn(name = "user_id"),
            inverseJoinColumns = @JoinColumn(name = "skill_id")
    )
    private List<SkillEntity> skills;

    @OneToOne
    private FileEntity avatar;

    @Column(length = 1000)
    private String about;

    @Column(length = 2048)
    private String password;

    private String username;

    @Column(name = "private_key", length = 2048)
    private String privateKey;

    @Column(name = "public_key", length = 2048)
    private String publicKey;

    @ManyToOne
    @JoinColumn(name = "role_id", nullable = false)
    private RoleEntity role;

    @ManyToMany
    @JoinTable(
            name = "users_languages",
            joinColumns = @JoinColumn(name = "user_id"),
            inverseJoinColumns = @JoinColumn(name = "language_id")
    )
    private List<LanguageEntity> languages;
}
