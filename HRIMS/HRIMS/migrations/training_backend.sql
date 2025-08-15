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
-- Table: leave_requests
-- ========================================
CREATE TABLE IF NOT EXISTS public.leave_requests (
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL ,
    leave_type VARCHAR(50) NOT NULL,
    reason TEXT,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_days INTEGER GENERATED ALWAYS AS ((end_date - start_date) + 1) STORED,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    reviewed_by INTEGER NOT NULL,
    reviewed_at TIMESTAMP WITH TIME ZONE,
    comment TEXT,
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
ALTER TABLE public.leave_requests OWNER TO postgres;
-- ========================================
-- Table: maintenance
-- ========================================
CREATE TABLE IF NOT EXISTS public.maintenance (
    maintenance_id SERIAL NOT NULL,
    vehicle_id INTEGER,
    maintenance_type VARCHAR(50),
    description TEXT,
    maintenance_date DATE NOT NULL,
    service_provider VARCHAR(100),
    cost NUMERIC(12,2),
    next_service_due DATE,
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
    PRIMARY KEY (maintenance_id)
);
ALTER TABLE public.maintenance OWNER TO postgres;
-- ========================================
-- Table: notification
-- ========================================
CREATE TABLE IF NOT EXISTS public.notification (
    id SERIAL NOT NULL,
    user_id INTEGER ,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
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
ALTER TABLE public.notification OWNER TO postgres;
-- ========================================
-- Table: officer_roles
-- ========================================
CREATE TABLE IF NOT EXISTS public.officer_roles (
    id SERIAL NOT NULL,
    role_name VARCHAR(100) NOT NULL,
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
    PRIMARY KEY(id)
);
ALTER TABLE public.officer_roles OWNER TO postgres;
-- ========================================

-- Table: position
-- ========================================
CREATE TABLE IF NOT EXISTS public.position (
    position_id SERIAL NOT NULL,
    position_title CHARACTER VARYING(255) NOT NULL,
    position_code CHARACTER VARYING(50) NOT NULL,
    position_type CHARACTER VARYING(100) NOT NULL,
    position_qualification TEXT,
    job_description TEXT,
    department_id INTEGER NOT NULL,
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
    PRIMARY KEY(position_id),
    CONSTRAINT position_position_code_key UNIQUE (position_code),
    CONSTRAINT fk_department FOREIGN KEY (department_id)
        REFERENCES public.departments (id)
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
ALTER TABLE public.position OWNER TO postgres;
-- ========================================
-- Table: officers
-- ========================================
CREATE TABLE IF NOT EXISTS public.officers (
    id SERIAL NOT NULL,
    user_id INTEGER,
    department_id INTEGER,
    position CHARACTER VARYING(100) NOT NULL,
    phone CHARACTER VARYING(20),
    designation CHARACTER VARYING(100),
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
    CONSTRAINT officers_user_id_key UNIQUE (user_id),
    CONSTRAINT officers_department_id_fkey FOREIGN KEY (department_id)
        REFERENCES public.departments (id)
        ON UPDATE NO ACTION
        ON DELETE SET NULL,
    CONSTRAINT officers_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.users (id)
        ON UPDATE NO ACTION
        ON DELETE CASCADE
);
ALTER TABLE public.officers OWNER TO postgres;
