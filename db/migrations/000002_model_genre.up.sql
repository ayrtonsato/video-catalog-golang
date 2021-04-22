CREATE TABLE IF NOT EXISTS "genres" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "categories_genres" (
    category_id UUID NOT NULL,
    genre_id UUID NOT NULL,
    PRIMARY KEY (category_id, genre_id),
    CONSTRAINT fk_category_genre
        FOREIGN KEY (category_id)
            REFERENCES categories(id),
    CONSTRAINT fk_genre_category
        FOREIGN KEY (genre_id)
            REFERENCES genres(id)
);