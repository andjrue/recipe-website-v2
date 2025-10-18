package org.springwebsite.recipewebsite.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;

import java.util.List;


public interface RecipeRepository extends JpaRepository<Recipe, String> {

    List<Recipe> findByTitleContainingIgnoreCaseAndIsDeletedFalse(String title);
    List<Recipe> findByUserIdAndIsDeletedFalse(User user);
    List<Recipe> findTop10ByOrderByDatePostedDesc();
}
