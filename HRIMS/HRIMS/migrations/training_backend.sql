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
-- Table: Application
-- ========================================
CREATE TABLE IF NOT EXISTS public.application (
    id SERIAL PRIMARY KEY,
    applicant_id INTEGER REFERENCES public.users(id),
    job_id INTEGER REFERENCES public.job_advert(id),
    resume TEXT,
    status VARCHAR(20) DEFAULT 'Pending',
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
    deleted_at TIMESTAMP WITH TIME ZONE
);
ALTER TABLE public.application OWNER TO postgres;

-- ========================================
-- Table: Approval
-- ========================================
CREATE TABLE IF NOT EXISTS public.approval (
    id SERIAL PRIMARY KEY,
    request_id INTEGER REFERENCES public.requests(id) ON DELETE CASCADE,
    approved_by INTEGER REFERENCES public.users(id),
    status VARCHAR(20) NOT NULL,
    comment TEXT,
    approved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER REFERENCES public.officers(id) ON DELETE SET NULL,
    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
ALTER TABLE public.approval OWNER TO postgres;

-- ========================================
-- Table: Audit_logs
-- ========================================
CREATE TABLE IF NOT EXISTS public.audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES public.users(id) ON DELETE SET NULL,
    action VARCHAR(100) NOT NULL,
    table_name VARCHAR(50),
    record_id INTEGER,
    details TEXT,
    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),
    created_by INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
ALTER TABLE public.audit_logs OWNER TO postgres;


-- ========================================
-- Table: Budget
-- ========================================
CREATE TABLE IF NOT EXISTS public.budget (
    budget_id SERIAL PRIMARY KEY,
    request_id INTEGER NOT NULL REFERENCES public.requests(id) ON DELETE CASCADE,
    salary_scale_id INTEGER NOT NULL REFERENCES public.salary_scale(id) ON DELETE CASCADE,
    number_of_officers INTEGER NOT NULL,
    total_budget NUMERIC(15,2) NOT NULL,
    description TEXT,
    created_by VARCHAR(100),
    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
ALTER TABLE public.budget OWNER TO postgres;
-- ========================================
-- Table: Campus
-- ========================================
CREATE TABLE IF NOT EXISTS public.campus (
    campus_id SERIAL PRIMARY KEY,
    campus_name VARCHAR(255) NOT NULL,
    description TEXT,
    department_id INTEGER REFERENCES public.departments(id) ON DELETE SET NULL,
    resvd5 CHAR(1),
    resvd4 CHAR(1),
    resvd3 CHAR(1),
    resvd2 CHAR(1),
    resvd1 CHAR(1),
    created_by INTEGER,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);
ALTER TABLE public.campus OWNER TO postgres;
