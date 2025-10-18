package org.springwebsite.recipewebsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name = "comments")
@Entity
@Data
public class Comment {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @ManyToOne
    @JoinColumn(name = "recipe_id")
    private Recipe recipe;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    private String content;
    private String datePosted;
    private boolean isDeleted;
}
