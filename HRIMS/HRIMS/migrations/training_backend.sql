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
-- Table: position
-- ========================================

-- ========================================
-- Table: employees
-- ========================================
CREATE TABLE IF NOT EXISTS public.employees
(
    id SERIAL NOT NULL,
    first_name CHARACTER VARYING(100),
    last_name CHARACTER VARYING(100),
    email CHARACTER VARYING(150),
    phone_number CHARACTER VARYING(20),
    department CHARACTER VARYING(100),
    "position" CHARACTER VARYING(100),
    hire_date DATE,
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
    CONSTRAINT employees_pkey PRIMARY KEY (id),
    CONSTRAINT employees_email_key UNIQUE (email)
);
ALTER TABLE public.employees OWNER TO postgres;
-- ========================================
-- Table: employees
-- ========================================


-- ========================================
-- Table: equipment_issues
-- ========================================
CREATE TABLE IF NOT EXISTS public.equipment_issues
(
    id SERIAL NOT NULL,
    equipment_id INTEGER,
    issued_to INTEGER,
    issue_date DATE NOT NULL,
    return_date DATE,
    returned_condition CHARACTER VARYING(100),
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
    CONSTRAINT equipment_issues_pkey PRIMARY KEY (id)
);
ALTER TABLE public.equipment_issues OWNER TO postgres;
-- ========================================
-- Table: equipment_issues
-- ========================================

-- ========================================
-- Table: fuel_records
-- ========================================
CREATE TABLE IF NOT EXISTS public.fuel_records
(
    fuel_id SERIAL NOT NULL,
    vehicle_id INTEGER,
    fueling_date DATE NOT NULL,
    fuel_type CHARACTER VARYING(20),
    quantity_liters NUMERIC(10,2) NOT NULL,
    cost NUMERIC(12,2),
    odometer_reading INTEGER,
    fueling_station CHARACTER VARYING(100),
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
    CONSTRAINT fuel_records_pkey PRIMARY KEY (fuel_id)
);
ALTER TABLE public.fuel_records OWNER TO postgres;
-- ========================================
-- Table: fuel_records
-- ========================================

-- ========================================
-- Table: internship_application
-- ========================================
CREATE TABLE IF NOT EXISTS public.internship_application
(
    id SERIAL NOT NULL,
    student_id INTEGER,
    department_id INTEGER,
    resume TEXT,
    status CHARACTER VARYING(20) DEFAULT 'Pending',
    resvd5 CHARACTER VARYING(1),
    resvd4 CHARACTER VARYING(1),
    resvd3 CHARACTER VARYING(1),
    resvd2 CHARACTER VARYING(1),
    resvd1 CHARACTER VARYING(1),
    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    applied_at TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME ZONE,
    deleted_at TIMESTAMP(0) WITH TIME ZONE,
    CONSTRAINT internship_application_pkey PRIMARY KEY (id)
);
ALTER TABLE public.internship_application OWNER TO postgres;
-- ========================================
-- Table: internship_application
-- ========================================

-- ========================================
-- Table: job_advert
-- ========================================
CREATE TABLE IF NOT EXISTS public.job_advert
(
    id SERIAL NOT NULL,
    title CHARACTER VARYING(100) NOT NULL,
    description TEXT NOT NULL,
    department_id INTEGER,
    posted_by INTEGER,
    deadline DATE,
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
    CONSTRAINT job_advert_pkey PRIMARY KEY (id)
);
ALTER TABLE public.job_advert OWNER TO postgres;
-- ========================================
-- Table: job_advert
-- ========================================

-- ========================================
-- Table: leave_balances
-- ========================================
CREATE TABLE IF NOT EXISTS public.leave_balances
(
    id SERIAL NOT NULL,
    user_id INTEGER NOT NULL,
    leave_type CHARACTER VARYING(50) NOT NULL,
    year INTEGER NOT NULL,
    total_entitled INTEGER NOT NULL,
    used_days INTEGER NOT NULL DEFAULT 0,
    remaining_days INTEGER GENERATED ALWAYS AS ((total_entitled - used_days)) STORED,
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
    CONSTRAINT leave_balances_pkey PRIMARY KEY (id),
    CONSTRAINT leave_balances_user_id_leave_type_year_key UNIQUE (user_id, leave_type, year)
);
ALTER TABLE public.leave_balances OWNER TO postgres;
-- ========================================
-- Table: leave_balances
-- ========================================



-- ========================================
-- Table: student_training_enrollment
-- ========================================
CREATE TABLE IF NOT EXISTS public.student_training_enrollment
(
    id SERIAL NOT NULL,
    student_id INTEGER,
    training_program_id INTEGER,
    enrolled_at TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status CHARACTER VARYING(20) DEFAULT 'Enrolled',

    resvd5 CHARACTER VARYING(1),
    resvd4 CHARACTER VARYING(1),
    resvd3 CHARACTER VARYING(1),
    resvd2 CHARACTER VARYING(1),
    resvd1 CHARACTER VARYING(1),

    created_by INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    updated_at TIMESTAMP(0) WITH TIME ZONE,
    deleted_at TIMESTAMP(0) WITH TIME ZONE,

    CONSTRAINT student_training_enrollment_pkey PRIMARY KEY (id)
);
ALTER TABLE public.student_training_enrollment OWNER TO postgres;
-- ========================================
-- Table: student_training_enrollment
-- ========================================


-- ========================================
-- Table: team
-- ========================================
CREATE TABLE IF NOT EXISTS public.team
(
    id SERIAL NOT NULL,
    team_name CHARACTER VARYING(100) NOT NULL,
    description TEXT,
    created_by INTEGER,

    resvd5 CHARACTER VARYING(1),
    resvd4 CHARACTER VARYING(1),
    resvd3 CHARACTER VARYING(1),
    resvd2 CHARACTER VARYING(1),
    resvd1 CHARACTER VARYING(1),

    created_by_user INTEGER,
    updated_by INTEGER,
    deleted_by INTEGER,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME ZONE,
    deleted_at TIMESTAMP(0) WITH TIME ZONE,

    CONSTRAINT team_pkey PRIMARY KEY (id),
    CONSTRAINT team_team_name_key UNIQUE (team_name)
);
ALTER TABLE public.team OWNER TO postgres;
-- ========================================
-- Table: team
-- ========================================


-- ========================================
-- Table: training_program
-- ========================================
CREATE TABLE IF NOT EXISTS public.training_program
(
    id SERIAL NOT NULL,
    title CHARACTER VARYING(100),
    description TEXT,
    start_date DATE,
    end_date DATE,
    department_id INTEGER,

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

    CONSTRAINT training_program_pkey PRIMARY KEY (id)
);
ALTER TABLE public.training_program OWNER TO postgres;
-- ========================================
-- Table: training_program
-- ========================================


-- ========================================
-- Table: transport_requests
-- ========================================
CREATE TABLE IF NOT EXISTS public.transport_requests
(
    request_id SERIAL NOT NULL,
    requester_id INTEGER NOT NULL,
    driver_id INTEGER,
    vehicle_id INTEGER,
    origin CHARACTER VARYING(100) NOT NULL,
    destination CHARACTER VARYING(100) NOT NULL,
    purpose TEXT,
    requested_date DATE NOT NULL,
    departure_time TIMESTAMP(0) WITH TIME ZONE,
    return_time TIMESTAMP(0) WITH TIME ZONE,
    approval_status CHARACTER VARYING(20) DEFAULT 'Pending',

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

    CONSTRAINT transport_requests_pkey PRIMARY KEY (request_id)
);
ALTER TABLE public.transport_requests OWNER TO postgres;
-- ========================================
-- Table: transport_requests
-- ========================================

-- ========================================
-- Table: users
-- ========================================
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL NOT NULL,
    first_name CHARACTER VARYING(50) NOT NULL,
    middle_name CHARACTER VARYING(100),
    surname CHARACTER VARYING(100),
    age INTEGER,
    email CHARACTER VARYING(100) NOT NULL,
    phone_number CHARACTER VARYING(20),
    password CHARACTER VARYING(255),
    password_hash TEXT NOT NULL,
    role_id INTEGER,
    is_active BOOLEAN DEFAULT TRUE,

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

    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (first_name)
);
ALTER TABLE public.users OWNER TO postgres;
-- ========================================
-- Table: users
-- ========================================
