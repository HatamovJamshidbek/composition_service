CREATE TABLE tracks (
                        id uuid primary key default gen_random_uuid(),
                        composition_id uuid,
                        user_id uuid,
                        title VARCHAR(100) NOT NULL,
                        file_url VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE ,
                        deleted_at TIMESTAMP WITH TIME ZONE
);