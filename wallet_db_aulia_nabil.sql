--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: password_resets; Type: TABLE; Schema: public; Owner: aulianabil
--

CREATE TABLE public.password_resets (
    id bigint NOT NULL,
    user_id bigint,
    token text,
    expired_at timestamp with time zone
);


ALTER TABLE public.password_resets OWNER TO aulianabil;

--
-- Name: password_resets_id_seq; Type: SEQUENCE; Schema: public; Owner: aulianabil
--

CREATE SEQUENCE public.password_resets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.password_resets_id_seq OWNER TO aulianabil;

--
-- Name: password_resets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: aulianabil
--

ALTER SEQUENCE public.password_resets_id_seq OWNED BY public.password_resets.id;


--
-- Name: source_of_funds; Type: TABLE; Schema: public; Owner: aulianabil
--

CREATE TABLE public.source_of_funds (
    id bigint NOT NULL,
    name text
);


ALTER TABLE public.source_of_funds OWNER TO aulianabil;

--
-- Name: source_of_funds_id_seq; Type: SEQUENCE; Schema: public; Owner: aulianabil
--

CREATE SEQUENCE public.source_of_funds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.source_of_funds_id_seq OWNER TO aulianabil;

--
-- Name: source_of_funds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: aulianabil
--

ALTER SEQUENCE public.source_of_funds_id_seq OWNED BY public.source_of_funds.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: aulianabil
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    source_of_fund_id bigint,
    user_id bigint,
    destination_id bigint,
    amount bigint,
    description text,
    category text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO aulianabil;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: aulianabil
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO aulianabil;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: aulianabil
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: aulianabil
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text,
    email text,
    password text
);


ALTER TABLE public.users OWNER TO aulianabil;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: aulianabil
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO aulianabil;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: aulianabil
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: wallets; Type: TABLE; Schema: public; Owner: aulianabil
--

CREATE TABLE public.wallets (
    id bigint NOT NULL,
    user_id bigint,
    number text,
    balance bigint
);


ALTER TABLE public.wallets OWNER TO aulianabil;

--
-- Name: wallets_id_seq; Type: SEQUENCE; Schema: public; Owner: aulianabil
--

CREATE SEQUENCE public.wallets_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wallets_id_seq OWNER TO aulianabil;

--
-- Name: wallets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: aulianabil
--

ALTER SEQUENCE public.wallets_id_seq OWNED BY public.wallets.id;


--
-- Name: password_resets id; Type: DEFAULT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.password_resets ALTER COLUMN id SET DEFAULT nextval('public.password_resets_id_seq'::regclass);


--
-- Name: source_of_funds id; Type: DEFAULT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.source_of_funds ALTER COLUMN id SET DEFAULT nextval('public.source_of_funds_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: wallets id; Type: DEFAULT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.wallets ALTER COLUMN id SET DEFAULT nextval('public.wallets_id_seq'::regclass);

COPY public.source_of_funds (id, name) FROM stdin;
1	Bank Transfer
2	Credit Card
3	Cash
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: aulianabil
--

COPY public.transactions (id, source_of_fund_id, user_id, destination_id, amount, description, category, created_at, updated_at, deleted_at) FROM stdin;
9	1	2	2	1000000	Top Up from Bank Transfer	Top Up	2022-09-09 17:21:42.274626+07	2022-09-09 17:21:42.274626+07	\N
10	\N	1	2	5000	bayarin somay	Receive Money	2022-09-09 17:22:08.248752+07	2022-09-09 17:22:08.248752+07	\N
13	\N	2	3	5000	beli pulsa	Send Money	2022-09-09 17:22:35.566679+07	2022-09-09 17:22:35.566679+07	\N
20	3	4	4	889000	Top Up from Cash	Top Up	2022-09-09 17:24:03.627458+07	2022-09-09 17:24:03.627458+07	\N
21	\N	2	4	77000	beli baju	Receive Money	2022-09-09 17:24:18.534219+07	2022-09-09 17:24:18.534219+07	\N
22	\N	4	2	77000	beli baju	Send Money	2022-09-09 17:24:18.535259+07	2022-09-09 17:24:18.535259+07	\N
25	2	5	5	778111	Top Up from Credit Card	Top Up	2022-09-09 17:24:49.956445+07	2022-09-09 17:24:49.956445+07	\N
27	\N	5	4	99900	bayarin mekdi	Send Money	2022-09-09 17:25:10.580477+07	2022-09-09 17:25:10.580477+07	\N
1	1	1	1	1000000	Top Up from Bank Transfer	Top Up	2022-09-09 17:20:24.17764+07	2022-09-10 17:20:24.17764+07	\N
2	3	1	1	50000	Top Up from Cash	Top Up	2022-08-31 00:00:00+07	2022-09-09 17:20:29.898282+07	\N
4	\N	2	1	36000	Bayar netflix	Receive Money	2022-09-09 17:21:03.899876+07	2022-08-24 17:26:47+07	\N
5	\N	1	2	36000	Bayar netflix	Send Money	2022-09-09 17:21:03.901019+07	2022-05-09 17:21:03.901019+07	\N
11	\N	2	1	5000	bayarin somay	Send Money	2022-09-09 17:22:08.249826+07	2021-09-09 17:22:08.249826+07	\N
16	\N	5	3	5000	beli makan	Receive Money	2022-09-09 17:23:22.831038+07	2021-09-09 17:23:22.831038+07	\N
6	\N	4	1	80000	Bayar Hutang	Receive Money	2022-09-09 17:21:18.474015+07	2021-08-09 17:21:18.474015+07	\N
19	\N	3	5	55000	hbd	Send Money	2022-09-09 17:23:40.475604+07	2021-09-09 17:23:40.475604+07	\N
23	\N	3	4	77000	beli sendal	Receive Money	2022-09-09 17:24:26.488124+07	2022-01-09 17:24:26.488124+07	\N
7	\N	1	4	80000	Bayar Hutang	Send Money	2022-09-09 17:21:18.47527+07	2022-02-09 17:21:18.47527+07	\N
8	2	2	2	400000	Top Up from Credit Card	Top Up	2022-09-09 17:21:32.946112+07	2022-03-09 17:21:32.946112+07	\N
3	2	1	1	400000	Top Up from Credit Card	Top Up	2022-09-09 17:20:42.822551+07	2022-05-09 17:20:42.822551+07	\N
12	\N	3	2	5000	beli pulsa	Receive Money	2022-09-09 17:22:35.565533+07	2022-03-09 17:22:35.565533+07	\N
14	1	3	3	1000000	Top Up from Bank Transfer	Top Up	2022-09-09 17:23:06.810063+07	2022-06-09 17:23:06.810063+07	\N
15	3	3	3	50000	Top Up from Cash	Top Up	2022-09-09 17:23:13.11464+07	2022-10-09 17:23:13.11464+07	\N
18	\N	5	3	55000	hbd	Receive Money	2022-09-09 17:23:40.47433+07	2022-12-09 17:23:40.47433+07	\N
24	\N	4	3	77000	beli sendal	Send Money	2022-09-09 17:24:26.489297+07	2022-11-09 17:24:26.489297+07	\N
26	\N	4	5	99900	bayarin mekdi	Receive Money	2022-09-09 17:25:10.579054+07	2022-07-09 17:25:10.579054+07	\N
17	\N	3	5	5000	beli makan	Send Money	2022-09-09 17:23:22.83226+07	2022-06-09 17:23:22.83226+07	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: aulianabil
--

COPY public.users (id, name, email, password) FROM stdin;
1	nabil	nabil@user.com	$2a$04$6cp6IBREiRW27La2hkM.uekYT71m/OzGKAW7vzfVRetPZwFPgREC.
2	mario	mario@user.com	$2a$04$93AZUXoqhOu6TNb481MYke3iDbM8UAzizOHmKSEf36bQtzV3kffwm
3	gerald	gerald@user.com	$2a$04$dwl4i0hAV/x/OUEbbPB8gufGHLyXhxK7ZSfIREjQ7xcX08OTrtKFW
4	azmi	azmi@user.com	$2a$04$I0WOe4FMjq/.k9PowmJ7x.1cmBT1vBLNj/L6FBm3IuznHbdwVSPFa
5	arkin	arkin@user.com	$2a$04$ReD6.DB5iRo3Gry9AoZiaeCtLIewk6g6XTDcpZpCLbdfXG33q25ea
\.


--
-- Data for Name: wallets; Type: TABLE DATA; Schema: public; Owner: aulianabil
--

COPY public.wallets (id, user_id, number, balance) FROM stdin;
1	1	100001	1339000
2	2	100002	1503000
3	3	100003	1072000
5	5	100005	738211
4	4	100004	914900
\.


--
-- Name: password_resets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: aulianabil
--

SELECT pg_catalog.setval('public.password_resets_id_seq', 1, false);


--
-- Name: source_of_funds_id_seq; Type: SEQUENCE SET; Schema: public; Owner: aulianabil
--

SELECT pg_catalog.setval('public.source_of_funds_id_seq', 1, false);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: aulianabil
--

SELECT pg_catalog.setval('public.transactions_id_seq', 27, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: aulianabil
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- Name: wallets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: aulianabil
--

SELECT pg_catalog.setval('public.wallets_id_seq', 5, true);


--
-- Name: password_resets password_resets_pkey; Type: CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.password_resets
    ADD CONSTRAINT password_resets_pkey PRIMARY KEY (id);


--
-- Name: source_of_funds source_of_funds_pkey; Type: CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.source_of_funds
    ADD CONSTRAINT source_of_funds_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: wallets wallets_pkey; Type: CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT wallets_pkey PRIMARY KEY (id);


--
-- Name: idx_transactions_deleted_at; Type: INDEX; Schema: public; Owner: aulianabil
--

CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


--
-- Name: password_resets fk_password_resets_user; Type: FK CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.password_resets
    ADD CONSTRAINT fk_password_resets_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: transactions fk_transactions_source_of_fund; Type: FK CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_source_of_fund FOREIGN KEY (source_of_fund_id) REFERENCES public.source_of_funds(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: transactions fk_transactions_user; Type: FK CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: transactions fk_transactions_wallet; Type: FK CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT fk_transactions_wallet FOREIGN KEY (destination_id) REFERENCES public.wallets(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: wallets fk_wallets_user; Type: FK CONSTRAINT; Schema: public; Owner: aulianabil
--

ALTER TABLE ONLY public.wallets
    ADD CONSTRAINT fk_wallets_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

