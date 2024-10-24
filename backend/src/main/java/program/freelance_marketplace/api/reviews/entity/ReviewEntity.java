package program.freelance_marketplace.api.reviews.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.FileEntity;

import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "reviews")
public class ReviewEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Integer rating;

    @OneToMany(orphanRemoval = true)
    @JoinTable(
            name = "reviews_files",
            joinColumns = @JoinColumn(name = "review_id"),
            inverseJoinColumns = @JoinColumn(name = "file_id")
    )
    private List<FileEntity> files;

    @Column(columnDefinition = "TEXT")
    private String content;
}
