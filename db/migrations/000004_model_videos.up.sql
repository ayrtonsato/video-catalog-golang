CREATE TYPE rating_video AS ENUM ('L', '10', '12', '14', '16', '18');

CREATE TABLE IF NOT EXISTS "videos" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    year_launched INT NOT NULL,
    opened BOOLEAN NOT NULL DEFAULT TRUE,
    rating rating_video NOT NULL,
    duration INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "video_category" (
    video_id UUID NOT NULL,
    category_id UUID NOT NULL,
    PRIMARY KEY (video_id, category_id),
    CONSTRAINT fk_video_category
        FOREIGN KEY (video_id)
            REFERENCES videos(id),
    CONSTRAINT fk_category_video
        FOREIGN KEY (category_id)
            REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS "video_genre" (
    video_id UUID NOT NULL,
    genre_id UUID NOT NULL,
    PRIMARY KEY (video_id, genre_id),
    CONSTRAINT fk_video_genre
        FOREIGN KEY (video_id)
            REFERENCES videos(id),
    CONSTRAINT fk_genre_video
        FOREIGN KEY (genre_id)
            REFERENCES genres(id)
);

CREATE TABLE IF NOT EXISTS "video_castmember" (
    video_id UUID NOT NULL,
    castmember_id UUID NOT NULL,
    PRIMARY KEY (video_id, castmember_id),
    CONSTRAINT fk_video_castmember
        FOREIGN KEY (video_id)
            REFERENCES videos(id),
    CONSTRAINT fk_castmember_video
        FOREIGN KEY (castmember_id)
            REFERENCES castmembers(id)
)