-- Active: 1708385048894@@127.0.0.1@5432@api
INSERT INTO tasks 
(title, is_completed, created_at, updated_at)
VALUES
('running running running', false, now(), now()),
('reading', true, now(), now());