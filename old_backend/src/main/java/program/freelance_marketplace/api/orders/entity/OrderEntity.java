package program.freelance_marketplace.api.orders.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.*;
import program.freelance_marketplace.api.reviews.entity.ReviewEntity;
import program.freelance_marketplace.api.freelance_services.entity.ServiceEntity;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name = "orders")
public class OrderEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Timestamp created_at;
    private Timestamp ended_at;

    @ManyToOne
    private UserEntity customer;

    @ManyToOne
    private UserEntity freelancer;

    @ManyToOne
    private StatusEntity status;

    @ManyToOne
    private ServiceEntity service;

    @ManyToOne
    private PackageEntity servicePackage;

    @OneToOne
    private ReviewEntity review;

    @OneToOne
    private ChatEntity chat;
}
