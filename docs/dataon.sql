PGDMP         ,            	    {            dataon    14.5 #   14.9 (Ubuntu 14.9-0ubuntu0.22.04.1) #               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                        0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            !           1262    28044    dataon    DATABASE     Z   CREATE DATABASE dataon WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';
    DROP DATABASE dataon;
                root    false            �            1259    28053 	   companies    TABLE     \   CREATE TABLE public.companies (
    id integer NOT NULL,
    name character varying(255)
);
    DROP TABLE public.companies;
       public         heap    root    false            �            1259    28052    company_id_seq    SEQUENCE     �   CREATE SEQUENCE public.company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.company_id_seq;
       public          root    false    210            "           0    0    company_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.company_id_seq OWNED BY public.companies.id;
          public          root    false    209            �            1259    28072 	   divisions    TABLE     �   CREATE TABLE public.divisions (
    id integer NOT NULL,
    name character varying(255),
    executive_committee_id integer
);
    DROP TABLE public.divisions;
       public         heap    root    false            �            1259    28071    divisions_id_seq    SEQUENCE     �   CREATE SEQUENCE public.divisions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.divisions_id_seq;
       public          root    false    214            #           0    0    divisions_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.divisions_id_seq OWNED BY public.divisions.id;
          public          root    false    213            �            1259    28060    executivecommittees    TABLE     ~   CREATE TABLE public.executivecommittees (
    id integer NOT NULL,
    name character varying(255),
    company_id integer
);
 '   DROP TABLE public.executivecommittees;
       public         heap    root    false            �            1259    28059    executivecommittees_id_seq    SEQUENCE     �   CREATE SEQUENCE public.executivecommittees_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.executivecommittees_id_seq;
       public          root    false    212            $           0    0    executivecommittees_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.executivecommittees_id_seq OWNED BY public.executivecommittees.id;
          public          root    false    211            �            1259    28084    sub_divisions    TABLE     y   CREATE TABLE public.sub_divisions (
    id integer NOT NULL,
    name character varying(255),
    division_id integer
);
 !   DROP TABLE public.sub_divisions;
       public         heap    root    false            �            1259    28083    sub_divisions_id_seq    SEQUENCE     �   CREATE SEQUENCE public.sub_divisions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.sub_divisions_id_seq;
       public          root    false    216            %           0    0    sub_divisions_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.sub_divisions_id_seq OWNED BY public.sub_divisions.id;
          public          root    false    215            z           2604    28056    companies id    DEFAULT     j   ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);
 ;   ALTER TABLE public.companies ALTER COLUMN id DROP DEFAULT;
       public          root    false    210    209    210            |           2604    28075    divisions id    DEFAULT     l   ALTER TABLE ONLY public.divisions ALTER COLUMN id SET DEFAULT nextval('public.divisions_id_seq'::regclass);
 ;   ALTER TABLE public.divisions ALTER COLUMN id DROP DEFAULT;
       public          root    false    213    214    214            {           2604    28063    executivecommittees id    DEFAULT     �   ALTER TABLE ONLY public.executivecommittees ALTER COLUMN id SET DEFAULT nextval('public.executivecommittees_id_seq'::regclass);
 E   ALTER TABLE public.executivecommittees ALTER COLUMN id DROP DEFAULT;
       public          root    false    212    211    212            }           2604    28087    sub_divisions id    DEFAULT     t   ALTER TABLE ONLY public.sub_divisions ALTER COLUMN id SET DEFAULT nextval('public.sub_divisions_id_seq'::regclass);
 ?   ALTER TABLE public.sub_divisions ALTER COLUMN id DROP DEFAULT;
       public          root    false    216    215    216                      0    28053 	   companies 
   TABLE DATA           -   COPY public.companies (id, name) FROM stdin;
    public          root    false    210   W&                 0    28072 	   divisions 
   TABLE DATA           E   COPY public.divisions (id, name, executive_committee_id) FROM stdin;
    public          root    false    214   �&                 0    28060    executivecommittees 
   TABLE DATA           C   COPY public.executivecommittees (id, name, company_id) FROM stdin;
    public          root    false    212   �&                 0    28084    sub_divisions 
   TABLE DATA           >   COPY public.sub_divisions (id, name, division_id) FROM stdin;
    public          root    false    216   '       &           0    0    company_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.company_id_seq', 3, true);
          public          root    false    209            '           0    0    divisions_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.divisions_id_seq', 3, true);
          public          root    false    213            (           0    0    executivecommittees_id_seq    SEQUENCE SET     H   SELECT pg_catalog.setval('public.executivecommittees_id_seq', 5, true);
          public          root    false    211            )           0    0    sub_divisions_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.sub_divisions_id_seq', 6, true);
          public          root    false    215                       2606    28058    companies company_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.companies
    ADD CONSTRAINT company_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.companies DROP CONSTRAINT company_pkey;
       public            root    false    210            �           2606    28077    divisions divisions_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.divisions
    ADD CONSTRAINT divisions_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.divisions DROP CONSTRAINT divisions_pkey;
       public            root    false    214            �           2606    28065 ,   executivecommittees executivecommittees_pkey 
   CONSTRAINT     j   ALTER TABLE ONLY public.executivecommittees
    ADD CONSTRAINT executivecommittees_pkey PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.executivecommittees DROP CONSTRAINT executivecommittees_pkey;
       public            root    false    212            �           2606    28089     sub_divisions sub_divisions_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.sub_divisions
    ADD CONSTRAINT sub_divisions_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.sub_divisions DROP CONSTRAINT sub_divisions_pkey;
       public            root    false    216            �           2606    28078 /   divisions divisions_executivecommittees_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.divisions
    ADD CONSTRAINT divisions_executivecommittees_id_fkey FOREIGN KEY (executive_committee_id) REFERENCES public.executivecommittees(id);
 Y   ALTER TABLE ONLY public.divisions DROP CONSTRAINT divisions_executivecommittees_id_fkey;
       public          root    false    214    3201    212            �           2606    28066 7   executivecommittees executivecommittees_company_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.executivecommittees
    ADD CONSTRAINT executivecommittees_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.companies(id);
 a   ALTER TABLE ONLY public.executivecommittees DROP CONSTRAINT executivecommittees_company_id_fkey;
       public          root    false    210    212    3199            �           2606    28090 ,   sub_divisions sub_divisions_division_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.sub_divisions
    ADD CONSTRAINT sub_divisions_division_id_fkey FOREIGN KEY (division_id) REFERENCES public.divisions(id);
 V   ALTER TABLE ONLY public.sub_divisions DROP CONSTRAINT sub_divisions_division_id_fkey;
       public          root    false    216    214    3203               #   x�3��2��K-Wp��-H̫T0�2F���qqq �H
\         $   x�3���4�2�NM.-�,��4�2F���qqq �8	u         7   x�3�t�wQ0�4�2���,c0��2��LNU(JM�LI�+
�b
��qqq D�-         B   x�3�t�4�2�NM.�,���`l������ļ��gp~ZIybQ��k^zf^jjP4F��� �5�     