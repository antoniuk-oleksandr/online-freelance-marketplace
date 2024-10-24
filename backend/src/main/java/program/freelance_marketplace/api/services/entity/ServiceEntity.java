package program.freelance_marketplace.api.services.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.CategoryEntity;
import program.freelance_marketplace.api.FileEntity;
import program.freelance_marketplace.api.PackageEntity;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.sql.Timestamp;
import java.time.Instant;
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
    private Timestamp createdAt;
    private Double rating;
    
    @Column(columnDefinition = "TEXT")
    private String description;

    @ManyToOne
    private FileEntity image;

    @ManyToOne
    private UserEntity freelancer;

    @ManyToOne
    private CategoryEntity category;

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
