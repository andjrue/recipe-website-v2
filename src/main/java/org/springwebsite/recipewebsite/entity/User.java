package org.springwebsite.recipewebsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name = "users")
@Data
@Entity
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String username;
    private String password; // will need to hash
    private String email;
    private String dateJoined;
}
