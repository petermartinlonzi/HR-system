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
    applicant_id INTEGER ,
    job_id INTEGER ,
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
    request_id INTEGER,
    approved_by INTEGER ,
    status VARCHAR(20) NOT NULL,
    comment TEXT,
    approved_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_by INTEGER,
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
    user_id INTEGER ,
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
    request_id INTEGER NOT NULL ,
    salary_scale_id INTEGER NOT NULL,
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
    department_id INTEGER,
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
