package org.springwebsite.recipewebsite.entity;

import jakarta.persistence.*;
import lombok.*;

@Table(name = "recipe_steps")
@Entity
@Data
public class RecipeSteps {

    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private String id;

    @ManyToOne
    @JoinColumn(name = "recipe_id")
    private Recipe recipe;

    private int stepNumber;
    private String instruction;
}
