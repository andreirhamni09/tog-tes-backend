PGDMP      $    	            }            tog    17.4 (Debian 17.4-1.pgdg120+2)    17.0 -    R           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            S           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            T           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            U           1262    16384    tog    DATABASE     n   CREATE DATABASE tog WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
    DROP DATABASE tog;
                     admin    false            �            1259    24589    bank_accounts    TABLE     �   CREATE TABLE public.bank_accounts (
    id bigint NOT NULL,
    customer_id bigint,
    account_number text NOT NULL,
    bank_amount bigint NOT NULL
);
 !   DROP TABLE public.bank_accounts;
       public         heap r       admin    false            �            1259    24588    bank_accounts_id_seq    SEQUENCE     }   CREATE SEQUENCE public.bank_accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.bank_accounts_id_seq;
       public               admin    false    220            V           0    0    bank_accounts_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.bank_accounts_id_seq OWNED BY public.bank_accounts.id;
          public               admin    false    219            �            1259    24578 	   customers    TABLE     o   CREATE TABLE public.customers (
    id bigint NOT NULL,
    username text,
    nama text,
    password text
);
    DROP TABLE public.customers;
       public         heap r       admin    false            �            1259    24577    customers_id_seq    SEQUENCE     y   CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.customers_id_seq;
       public               admin    false    218            W           0    0    customers_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;
          public               admin    false    217            �            1259    24629    deposito_histories    TABLE     �   CREATE TABLE public.deposito_histories (
    id bigint NOT NULL,
    deposito_id bigint,
    amount bigint NOT NULL,
    tanggal_input timestamp without time zone NOT NULL
);
 &   DROP TABLE public.deposito_histories;
       public         heap r       admin    false            �            1259    24628    deposito_histories_id_seq    SEQUENCE     �   CREATE SEQUENCE public.deposito_histories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 0   DROP SEQUENCE public.deposito_histories_id_seq;
       public               admin    false    226            X           0    0    deposito_histories_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE public.deposito_histories_id_seq OWNED BY public.deposito_histories.id;
          public               admin    false    225            �            1259    24617 	   depositos    TABLE     r   CREATE TABLE public.depositos (
    id bigint NOT NULL,
    bank_account_id bigint,
    amount bigint NOT NULL
);
    DROP TABLE public.depositos;
       public         heap r       admin    false            �            1259    24616    depositos_id_seq    SEQUENCE     y   CREATE SEQUENCE public.depositos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.depositos_id_seq;
       public               admin    false    224            Y           0    0    depositos_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.depositos_id_seq OWNED BY public.depositos.id;
          public               admin    false    223            �            1259    24605    pocket_informations    TABLE     �   CREATE TABLE public.pocket_informations (
    id bigint NOT NULL,
    bank_account_id bigint,
    amount bigint NOT NULL,
    tanggal_input timestamp without time zone NOT NULL
);
 '   DROP TABLE public.pocket_informations;
       public         heap r       admin    false            �            1259    24604    pocket_informations_id_seq    SEQUENCE     �   CREATE SEQUENCE public.pocket_informations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 1   DROP SEQUENCE public.pocket_informations_id_seq;
       public               admin    false    222            Z           0    0    pocket_informations_id_seq    SEQUENCE OWNED BY     Y   ALTER SEQUENCE public.pocket_informations_id_seq OWNED BY public.pocket_informations.id;
          public               admin    false    221            �           2604    24592    bank_accounts id    DEFAULT     t   ALTER TABLE ONLY public.bank_accounts ALTER COLUMN id SET DEFAULT nextval('public.bank_accounts_id_seq'::regclass);
 ?   ALTER TABLE public.bank_accounts ALTER COLUMN id DROP DEFAULT;
       public               admin    false    220    219    220            �           2604    24581    customers id    DEFAULT     l   ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);
 ;   ALTER TABLE public.customers ALTER COLUMN id DROP DEFAULT;
       public               admin    false    218    217    218            �           2604    24632    deposito_histories id    DEFAULT     ~   ALTER TABLE ONLY public.deposito_histories ALTER COLUMN id SET DEFAULT nextval('public.deposito_histories_id_seq'::regclass);
 D   ALTER TABLE public.deposito_histories ALTER COLUMN id DROP DEFAULT;
       public               admin    false    226    225    226            �           2604    24620    depositos id    DEFAULT     l   ALTER TABLE ONLY public.depositos ALTER COLUMN id SET DEFAULT nextval('public.depositos_id_seq'::regclass);
 ;   ALTER TABLE public.depositos ALTER COLUMN id DROP DEFAULT;
       public               admin    false    224    223    224            �           2604    24608    pocket_informations id    DEFAULT     �   ALTER TABLE ONLY public.pocket_informations ALTER COLUMN id SET DEFAULT nextval('public.pocket_informations_id_seq'::regclass);
 E   ALTER TABLE public.pocket_informations ALTER COLUMN id DROP DEFAULT;
       public               admin    false    221    222    222            I          0    24589    bank_accounts 
   TABLE DATA           U   COPY public.bank_accounts (id, customer_id, account_number, bank_amount) FROM stdin;
    public               admin    false    220   �4       G          0    24578 	   customers 
   TABLE DATA           A   COPY public.customers (id, username, nama, password) FROM stdin;
    public               admin    false    218   5       O          0    24629    deposito_histories 
   TABLE DATA           T   COPY public.deposito_histories (id, deposito_id, amount, tanggal_input) FROM stdin;
    public               admin    false    226   �5       M          0    24617 	   depositos 
   TABLE DATA           @   COPY public.depositos (id, bank_account_id, amount) FROM stdin;
    public               admin    false    224   �5       K          0    24605    pocket_informations 
   TABLE DATA           Y   COPY public.pocket_informations (id, bank_account_id, amount, tanggal_input) FROM stdin;
    public               admin    false    222    6       [           0    0    bank_accounts_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.bank_accounts_id_seq', 2, true);
          public               admin    false    219            \           0    0    customers_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.customers_id_seq', 68, true);
          public               admin    false    217            ]           0    0    deposito_histories_id_seq    SEQUENCE SET     H   SELECT pg_catalog.setval('public.deposito_histories_id_seq', 1, false);
          public               admin    false    225            ^           0    0    depositos_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.depositos_id_seq', 2, true);
          public               admin    false    223            _           0    0    pocket_informations_id_seq    SEQUENCE SET     H   SELECT pg_catalog.setval('public.pocket_informations_id_seq', 5, true);
          public               admin    false    221            �           2606    24598 .   bank_accounts uni_bank_accounts_account_number 
   CONSTRAINT     s   ALTER TABLE ONLY public.bank_accounts
    ADD CONSTRAINT uni_bank_accounts_account_number UNIQUE (account_number);
 X   ALTER TABLE ONLY public.bank_accounts DROP CONSTRAINT uni_bank_accounts_account_number;
       public                 admin    false    220            �           2606    24596 "   bank_accounts uni_bank_accounts_id 
   CONSTRAINT     `   ALTER TABLE ONLY public.bank_accounts
    ADD CONSTRAINT uni_bank_accounts_id PRIMARY KEY (id);
 L   ALTER TABLE ONLY public.bank_accounts DROP CONSTRAINT uni_bank_accounts_id;
       public                 admin    false    220            �           2606    24585    customers uni_customers_id 
   CONSTRAINT     X   ALTER TABLE ONLY public.customers
    ADD CONSTRAINT uni_customers_id PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.customers DROP CONSTRAINT uni_customers_id;
       public                 admin    false    218            �           2606    24587     customers uni_customers_username 
   CONSTRAINT     _   ALTER TABLE ONLY public.customers
    ADD CONSTRAINT uni_customers_username UNIQUE (username);
 J   ALTER TABLE ONLY public.customers DROP CONSTRAINT uni_customers_username;
       public                 admin    false    218            �           2606    24634 ,   deposito_histories uni_deposito_histories_id 
   CONSTRAINT     j   ALTER TABLE ONLY public.deposito_histories
    ADD CONSTRAINT uni_deposito_histories_id PRIMARY KEY (id);
 V   ALTER TABLE ONLY public.deposito_histories DROP CONSTRAINT uni_deposito_histories_id;
       public                 admin    false    226            �           2606    24622    depositos uni_depositos_id 
   CONSTRAINT     X   ALTER TABLE ONLY public.depositos
    ADD CONSTRAINT uni_depositos_id PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.depositos DROP CONSTRAINT uni_depositos_id;
       public                 admin    false    224            �           2606    24610 .   pocket_informations uni_pocket_informations_id 
   CONSTRAINT     l   ALTER TABLE ONLY public.pocket_informations
    ADD CONSTRAINT uni_pocket_informations_id PRIMARY KEY (id);
 X   ALTER TABLE ONLY public.pocket_informations DROP CONSTRAINT uni_pocket_informations_id;
       public                 admin    false    222            �           2606    24599 '   bank_accounts fk_bank_accounts_customer    FK CONSTRAINT     �   ALTER TABLE ONLY public.bank_accounts
    ADD CONSTRAINT fk_bank_accounts_customer FOREIGN KEY (customer_id) REFERENCES public.customers(id);
 Q   ALTER TABLE ONLY public.bank_accounts DROP CONSTRAINT fk_bank_accounts_customer;
       public               admin    false    218    3236    220            �           2606    24635 1   deposito_histories fk_deposito_histories_deposito    FK CONSTRAINT     �   ALTER TABLE ONLY public.deposito_histories
    ADD CONSTRAINT fk_deposito_histories_deposito FOREIGN KEY (deposito_id) REFERENCES public.depositos(id) ON UPDATE CASCADE ON DELETE CASCADE;
 [   ALTER TABLE ONLY public.deposito_histories DROP CONSTRAINT fk_deposito_histories_deposito;
       public               admin    false    224    3246    226            �           2606    24623 #   depositos fk_depositos_bank_account    FK CONSTRAINT     �   ALTER TABLE ONLY public.depositos
    ADD CONSTRAINT fk_depositos_bank_account FOREIGN KEY (bank_account_id) REFERENCES public.bank_accounts(id) ON UPDATE CASCADE ON DELETE CASCADE;
 M   ALTER TABLE ONLY public.depositos DROP CONSTRAINT fk_depositos_bank_account;
       public               admin    false    220    3242    224            �           2606    24611 7   pocket_informations fk_pocket_informations_bank_account    FK CONSTRAINT     �   ALTER TABLE ONLY public.pocket_informations
    ADD CONSTRAINT fk_pocket_informations_bank_account FOREIGN KEY (bank_account_id) REFERENCES public.bank_accounts(id) ON UPDATE CASCADE ON DELETE CASCADE;
 a   ALTER TABLE ONLY public.pocket_informations DROP CONSTRAINT fk_pocket_informations_bank_account;
       public               admin    false    222    3242    220            I   (   x�3�4�4422�4445 .#N#NSSS � ̏���� r��      G   �   x�e���  �s�1�2`��%z5Ƴɂ���z���K�Ф�]G	�w�^f��G��Z��X.��9睆z��M�^F��$2�2��^�3����"+4*گ4)rBjDԺ���:�Ց]����S�S�94ZHUS�>B�,��gZ�9+�;��w�g�.��J9b� ���"~$w��      O      x������ � �      M      x�3�4�4�2�4�1z\\\ �      K   :   x�3�4�45 N##S]c]C+ �2�+k�!kh�`h
�5�4�#���� :�4     