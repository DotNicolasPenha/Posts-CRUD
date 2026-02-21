package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupTables(conn *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := setupExtensions(conn, ctx); err != nil {
		return err
	}
	if err := setupTablePost(conn, ctx); err != nil {
		return err
	}
	if err := setupTableUser(conn, ctx); err != nil {
		return err
	}
	return nil
}
func setupExtensions(conn *pgxpool.Pool, ctx context.Context) error {
	query := `
	  CREATE EXTENSION IF NOT EXISTS pgcrypto;
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error to add extension pgcrypto: %s", err)
	}
	return nil
}
func setupTableUser(conn *pgxpool.Pool, ctx context.Context) error {
	query := ` 
	CREATE TABLE IF NOT EXISTS users (
	  ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		username TEXT NOT NULL,
		bio TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error to create the user table: %s", err)
	}
	return nil
}
func setupTablePost(conn *pgxpool.Pool, ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
	   ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		 username TEXT NOT NULL,
		 body TEXT NOT NULL,
		 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
	_, err := conn.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("Error to create the post table: %s", err)
	}
	return nil
}
