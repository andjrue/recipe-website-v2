package org.springwebsite.recipewebsite.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springwebsite.recipewebsite.entity.User;

import java.util.List;
import java.util.Optional;

public interface UserRepository extends JpaRepository<User, String> {
    Optional<User> findByUsername(String username);
    Optional<User> findByEmail(String email);
    boolean existsByUsername(String username);
    boolean existsByEmail(String email);
    List<User> findAllByOrderByDateJoinedDesc();
}
