package org.springwebsite.recipewebsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name= "recipe")
@Data
@Entity
public class Recipe {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    private String title;
    private String description;
    private String timeToCook;

    @ManyToOne
    @JoinColumn(name = "userId")
    private User userId;

    private boolean isDeleted;
    private String datePosted;
    private String lastEditedAt;
}
