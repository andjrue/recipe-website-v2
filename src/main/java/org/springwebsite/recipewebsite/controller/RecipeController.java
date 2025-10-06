package org.springwebsite.recipewebsite.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;
import org.springwebsite.recipewebsite.service.RecipeService;

import java.util.List;

@RestController
@RequestMapping("api/v1/recipes")
public class RecipeController {

    private final RecipeService recipeService;

    public RecipeController(RecipeService recipeService) {
        this.recipeService = recipeService;
    }

    @GetMapping
    public ResponseEntity<List<Recipe>> getAllRecipes() {
        List<Recipe> recipes = this.recipeService.findAll();
        return ResponseEntity.ok(recipes);
    }

    @GetMapping("/recent")
    public ResponseEntity<List<Recipe>> getRecentRecipes() {
        List<Recipe> recipes = this.recipeService.findTenMostRecentRecipes();
        return ResponseEntity.ok(recipes);
    }

    @GetMapping("/{id}")
    public ResponseEntity<List<Recipe>> getRecipesByUserId(@PathVariable String id) {
        List<Recipe> recipes = this.recipeService.findByUserId(id);
        return ResponseEntity.ok(recipes);
    }

    @PostMapping
    public ResponseEntity<Recipe> createRecipe(@RequestBody Recipe recipe, String userId) {
        try {
            // Service handles validation and saving
            Recipe newRecipe = recipeService.save(recipe, userId);
            // Returns 201 Created and the location of the new resource
            return ResponseEntity.status(HttpStatus.CREATED).body(newRecipe);
        } catch (IllegalArgumentException e) {
            // Translates validation errors into a 400 Bad Request
            return ResponseEntity.badRequest().build();
        }

    }


}
