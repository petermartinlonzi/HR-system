-- ========================================
-- Table: position
-- ========================================
CREATE TABLE IF NOT EXISTS public.position (
    id SERIAL NOT NULL,
    name CHARACTER VARYING(60) NOT NULL,
    description CHARACTER VARYING(60),
    resvd5 CHARACTER VARYING(1),
    resvd4 CHARACTER VARYING(1),
    resvd3 CHARACTER VARYING(1),
    resvd2 CHARACTER VARYING(1),
    resvd1 CHARACTER VARYING(1),
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME ZONE,
    deleted_at TIMESTAMP(0) WITH TIME ZONE,
    PRIMARY KEY(id)
);
ALTER TABLE public.position OWNER TO postgres;
-- ========================================
-- Table: requests
-- ========================================
CREATE TABLE IF NOT EXISTS public.requests (
    id SERIAL NOT NULL,
    officer_id INTEGER,
    title VARCHAR(100) NOT NULL,
    content TEXT,
    status VARCHAR(50) DEFAULT 'Pending',
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    position_id INTEGER REFERENCES public.position(id) ON DELETE SET NULL,
    department_id INTEGER REFERENCES public.departments(id) ON DELETE SET NULL,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id)
);
ALTER TABLE public.requests OWNER TO postgres;
-- ========================================
-- Table: roles
-- ========================================
CREATE TABLE IF NOT EXISTS public.roles (
    id SERIAL NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id),
    UNIQUE(name)
);
ALTER TABLE public.roles OWNER TO postgres;
-- ========================================
-- Table: salary_scale
-- ========================================
CREATE TABLE IF NOT EXISTS public.salary_scale (
    id SERIAL NOT NULL,
    salary_scale_name VARCHAR(255) NOT NULL,
    position_id INTEGER NOT NULL,
    maximum_salary NUMERIC(15,2) NOT NULL,
    currency_type VARCHAR(10) NOT NULL,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id),
    CONSTRAINT chk_salary_range CHECK (maximum_salary >= minimum_salary)
);
ALTER TABLE public.salary_scale OWNER TO postgres;
-- ========================================
-- Table: sports_equipment
-- ========================================
CREATE TABLE IF NOT EXISTS public.sports_equipment (
    id SERIAL NOT NULL,
    name VARCHAR(100) NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 0,
    condition VARCHAR(50) DEFAULT 'Good',
    location VARCHAR(100),
    last_maintenance_date DATE,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id)
);
ALTER TABLE public.sports_equipment OWNER TO postgres;
-- ========================================
-- Table: sports_events
-- ========================================
CREATE TABLE IF NOT EXISTS public.sports_events (
    id SERIAL NOT NULL,
    event_name VARCHAR(100) NOT NULL,
    description TEXT,
    location VARCHAR(100),
    event_date DATE NOT NULL,
    start_time TIME,
    end_time TIME,
    organized_by INTEGER,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id)
);
ALTER TABLE public.sports_events OWNER TO postgres;
-- ========================================
-- Table: sports_participants
-- ========================================
CREATE TABLE IF NOT EXISTS public.sports_participants (
    id SERIAL NOT NULL,
    event_id INTEGER,
    employee_id INTEGER,
    team_name VARCHAR(100),
    performance_notes TEXT,

    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,

    PRIMARY KEY(id)
);
ALTER TABLE public.sports_participants OWNER TO postgres;



