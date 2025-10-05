package org.springwebsite.recipewebsite.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;

import java.util.List;


public interface RecipeRepository extends JpaRepository<Recipe, String> {
    // All methods will filter out soft deleted recipes, I don't feel like adding that to the name
    List<Recipe> findByTitle(String title);
    List<Recipe> findByUserId(User user);
    List<Recipe> findTop10ByOrderByDatePostedDesc();
}
