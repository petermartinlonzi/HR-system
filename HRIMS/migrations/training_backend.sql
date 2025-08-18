-- ========================================
-- Table: position
-- ========================================
CREATE TABLE IF NOT EXISTS public.position (
    id SERIAL NOT NULL,
    name CHARACTER VARYING(60) NOT NULL,
    description CHARACTER VARYING(60),
   
    PRIMARY KEY(id) resvd5 CHARACTER VARYING(1),
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
);
ALTER TABLE public.position OWNER TO postgres;

-- ========================================
-- Table: position
-- ========================================


-- Table: public.contract_status

-- DROP TABLE IF EXISTS public.contract_status;

CREATE TABLE IF NOT EXISTS public.contract_status
(
    id integer NOT NULL DEFAULT nextval('contract_status_id_seq'::regclass),
    status_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
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
    CONSTRAINT contract_status_pkey PRIMARY KEY (id),
    CONSTRAINT contract_status_status_name_key UNIQUE (status_name)
);
ALTER TABLE IF EXISTS public.contract_status
    OWNER to postgres;


-- Table: public.contract_version

-- DROP TABLE IF EXISTS public.contract_version;

CREATE TABLE IF NOT EXISTS public.contract_version
(
    id integer NOT NULL DEFAULT nextval('contract_version_id_seq'::regclass),
    contract_id integer NOT NULL,
    version_number integer NOT NULL DEFAULT 1,
    salary numeric(10,2),
    benefits text COLLATE pg_catalog."default",
    working_hours character varying(50) COLLATE pg_catalog."default",
    probation_period character varying(50) COLLATE pg_catalog."default",
    signed_by integer,
    signed_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
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
    CONSTRAINT contract_version_pkey PRIMARY KEY (id),
    CONSTRAINT fk_contract FOREIGN KEY (contract_id)
     
);
ALTER TABLE IF EXISTS public.contract_version
    OWNER to postgres;

-- Table: public.departments

-- DROP TABLE IF EXISTS public.departments;

CREATE TABLE IF NOT EXISTS public.departments
(
    id integer NOT NULL DEFAULT nextval('departments_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
    directorate_id integer,
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
    CONSTRAINT departments_pkey PRIMARY KEY (id),
    CONSTRAINT departments_directorate_id_fkey FOREIGN KEY (directorate_id)
   
);
ALTER TABLE IF EXISTS public.departments
    OWNER to postgres;


-- Table: public.directorates

-- DROP TABLE IF EXISTS public.directorates;

CREATE TABLE IF NOT EXISTS public.directorates
(
    id integer NOT NULL DEFAULT nextval('directorates_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    description text COLLATE pg_catalog."default",
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
    CONSTRAINT directorates_pkey PRIMARY KEY (id)
);
ALTER TABLE IF EXISTS public.directorates
    OWNER to postgres;


-- Table: public.disciplinary_actions

-- DROP TABLE IF EXISTS public.disciplinary_actions;

CREATE TABLE IF NOT EXISTS public.disciplinary_actions
(
    id integer NOT NULL DEFAULT nextval('disciplinary_actions_id_seq'::regclass),
    employee_id integer,
    reported_by integer,
    incident_date date NOT NULL,
    description text COLLATE pg_catalog."default" NOT NULL,
    action_taken character varying(50) COLLATE pg_catalog."default",
    action_date date,
    duration integer,
    status character varying(20) COLLATE pg_catalog."default" DEFAULT 'Resolved'::character varying,
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
    CONSTRAINT disciplinary_actions_pkey PRIMARY KEY (id),
    CONSTRAINT disciplinary_actions_employee_id_fkey FOREIGN KEY (employee_id)
    CONSTRAINT disciplinary_actions_reported_by_fkey FOREIGN KEY (reported_by)
);
ALTER TABLE IF EXISTS public.disciplinary_actions
    OWNER to postgres;


-- Table: public.drivers

-- DROP TABLE IF EXISTS public.drivers;

CREATE TABLE IF NOT EXISTS public.drivers
(
    driver_id integer NOT NULL DEFAULT nextval('drivers_driver_id_seq'::regclass),
    first_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    license_number character varying(50) COLLATE pg_catalog."default" NOT NULL,
    license_expiry date,
    phone_number character varying(20) COLLATE pg_catalog."default",
    employment_status character varying(20) COLLATE pg_catalog."default" DEFAULT 'Active'::character varying,
    assigned_vehicle_id integer,
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
    CONSTRAINT drivers_pkey PRIMARY KEY (driver_id),
    CONSTRAINT drivers_license_number_key UNIQUE (license_number),
    CONSTRAINT drivers_assigned_vehicle_id_fkey FOREIGN KEY (assigned_vehicle_id)
);
ALTER TABLE IF EXISTS public.drivers
    OWNER to postgres;