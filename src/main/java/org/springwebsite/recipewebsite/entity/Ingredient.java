package org.springwebsite.recipewebsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name = "ingredients")
@Entity
@Data
public class Ingredient {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @ManyToOne
    @JoinColumn(name = "recipeId")
    private Recipe recipe;

    private String ingredient;
    private String quantity;
}
