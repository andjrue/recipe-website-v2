package org.springwebsite.recipewebsite.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springwebsite.recipewebsite.entity.Comment;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.entity.User;

import java.util.List;

public interface CommentRepository extends JpaRepository<Comment, String> {
    List<Comment> findByUser(User user);
    List<Comment> findByRecipe(String recipeId);
    long countByUser(User user);
    long countByRecipe(Recipe recipe);
    List<Comment> findTop10ByOrderByDatePostedDesc();
}
