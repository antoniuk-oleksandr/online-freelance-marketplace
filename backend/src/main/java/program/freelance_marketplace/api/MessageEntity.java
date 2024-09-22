package program.freelance_marketplace.api;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import program.freelance_marketplace.api.users.entity.UserEntity;

import java.sql.Timestamp;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "messages")
public class MessageEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Timestamp sent_at;
    private String content;

    @OneToMany(cascade = CascadeType.ALL)
    @JoinTable(
            name = "messages_files",
            joinColumns = @JoinColumn(name = "message_id"),
            inverseJoinColumns = @JoinColumn(name = "file_id")
    )
    private List<FileEntity> files;

    @ManyToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "sender_id")
    private UserEntity sender;

    @ManyToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "receiver_id")
    private UserEntity receiver;
}
