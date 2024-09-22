package program.freelance_marketplace.api;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.users.entity.LanguageEntity;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "services")
public class ServiceEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String title;
    private Timestamp created_at;

    @Column(columnDefinition = "TEXT")
    private String description;

    @ManyToOne
    private FileEntity image;

    @ManyToOne
    private UserEntity freelancer;

    @ManyToOne
    private CategoryEntity category;

    @ManyToMany
    @JoinTable(
            name = "services_languages",
            joinColumns = @JoinColumn(name = "service_id"),
            inverseJoinColumns = @JoinColumn(name = "language_id")
    )
    private List<LanguageEntity> languages;

    @OneToMany
    @JoinTable(
            name = "services_files",
            joinColumns = @JoinColumn(name = "service_id"),
            inverseJoinColumns = @JoinColumn(name = "file_id")
    )
    private List<FileEntity> files;

    @OneToMany
    @JoinTable(
            name = "services_packages",
            joinColumns = @JoinColumn(name = "service_id"),
            inverseJoinColumns = @JoinColumn(name = "package_id")
    )
    private List<PackageEntity> packages;
}
