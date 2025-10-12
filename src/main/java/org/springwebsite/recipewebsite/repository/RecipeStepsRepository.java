package org.springwebsite.recipewebsite.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springwebsite.recipewebsite.entity.RecipeSteps;

import java.util.List;

public interface RecipeStepsRepository extends JpaRepository<RecipeSteps, String> {

    List<RecipeSteps> findByRecipeIdOrderByStepNumberAsc(String recipeId);

}
