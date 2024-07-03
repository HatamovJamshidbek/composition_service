CREATE TYPE composition_status AS ENUM ('draft', 'in_progress', 'completed', 'published');

CREATE TABLE compositions (
                              id uuid primary key default gen_random_uuid(),
                              user_id uuid,
                              title VARCHAR(100) NOT NULL,
                              description TEXT,
                              status composition_status DEFAULT 'draft',
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE ,
                              deleted_at TIMESTAMP WITH TIME ZONE

);