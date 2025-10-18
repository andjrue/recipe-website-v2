package org.springwebsite.recipewebsite.controller;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springwebsite.recipewebsite.entity.Comment;
import org.springwebsite.recipewebsite.entity.Recipe;
import org.springwebsite.recipewebsite.service.CommentService;

import java.util.List;

@RestController
@RequestMapping("api/v1/comments")
public class CommentController {

    private final CommentService commentService;

    public CommentController(CommentService commentService) {
        this.commentService = commentService;
    }

    @GetMapping("/recipe/{recipeId}")
    public ResponseEntity<List<Comment>> findAllByRecipeId(@PathVariable String recipeId) {
        try {
            List<Comment> comments = this.commentService.findByRecipeId(recipeId);
            return ResponseEntity.ok(comments);
        }  catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }

    @GetMapping("/user/{userId}")
    public ResponseEntity<List<Comment>> findAllByUserId(@PathVariable String userId) {
        try {
            List<Comment> comments = this.commentService.findByUser(userId);
            return ResponseEntity.ok(comments);
        } catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }

    @GetMapping("/user/count/{userId}")
    public ResponseEntity<Long> countByUserId(@PathVariable String userId) {
        try {
            Long count = this.commentService.countByUser(userId);
            return ResponseEntity.ok(count);
        }  catch (Exception e) {
            return ResponseEntity.notFound().build();
        }
    }

    @GetMapping("/recent")
    public ResponseEntity<List<Comment>> getRecentComments() {
        List<Comment> comments = this.commentService.findMostRecent();
        return ResponseEntity.ok(comments);
    }

    @PostMapping("/recipe/{recipeId}")
    public ResponseEntity<Comment> createComment(
            @PathVariable String recipeId,
            @RequestBody Comment comment
    ) {
        String userId = "12345"; // Stub for now

        try {
            Comment savedComment = this.commentService.save(comment, recipeId, userId);
            return ResponseEntity.status(HttpStatus.CREATED).body(savedComment);
        } catch (IllegalArgumentException e) {
            return ResponseEntity.badRequest().build();
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).build();
        }
    }
}
