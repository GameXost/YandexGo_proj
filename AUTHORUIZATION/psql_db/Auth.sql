CREATE DATABASE Auth;
\connect Auth;



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
-- Name: cars; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cars (
    model character varying(50) NOT NULL,
    marks character varying(50) NOT NULL,
    data_a character varying(4) NOT NULL,
    data_b character varying(4) NOT NULL
);



--
-- Name: registered_drivers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.registered_drivers (
    id character varying(50) NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    phone_number character varying(10) NOT NULL,
    driver_license character varying(10) NOT NULL,
    driver_license_date character varying(10) NOT NULL,
    car_number character varying(10) NOT NULL,
    car_model character varying(20) NOT NULL,
    car_marks character varying(20) NOT NULL,
    car_color character varying(20) NOT NULL,
    password character varying(100),
    level_access integer
);





--
-- Name: registered_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.registered_users (
    id text NOT NULL,
    first_name character varying(50) NOT NULL,
    last_name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    phone_number character varying(10) NOT NULL,
    password character varying(100),
    level_access integer
);



--
-- Data for Name: cars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cars (model, marks, data_a, data_b) FROM stdin;
Acura	CL	1998	2003
Acura	EL	1997	2005
Acura	ILX	2012	2022
Acura	Integra	1995	-
Acura	MDX	2001	-
Acura	NSX	2001	-
Acura	RDX	2006	-
Acura	RL	1995	2012
Acura	RLX	2013	2021
Acura	RSX	2002	2006
Acura	TL	1999	2014
Acura	TLX	2014	-
Acura	TLX-L	2018	2021
Acura	TSX	2003	2014
Acura	ZDX	2010	2013
Alfa Romeo	146	1994	2001
Alfa Romeo	147	1998	2010
Alfa Romeo	147 GTA	2002	2010
Alfa Romeo	156	1997	2007
Alfa Romeo	156 GTA	2002	2007
Alfa Romeo	159	2005	2011
Alfa Romeo	166	1998	2007
Alfa Romeo	4C	2013	2020
Alfa Romeo	8C Competizione	2008	2010
Alfa Romeo	Brera	2005	2010
Alfa Romeo	Giulia	2016	-
Alfa Romeo	Giulietta	2010	2021
Alfa Romeo	GT	2003	2020
Alfa Romeo	GTV	1995	2006
Alfa Romeo	MiTo	2008	2018
Alfa Romeo	Spider	1998	2010
Alfa Romeo	Stelvio	2017	-
Alfa Romeo	Tonale	2022	-
Aston Martin	Cygnet	2011	2013
Aston Martin	DB11	2016	-
Aston Martin	DB12	2023	-
Aston Martin	DB9	2004	2016
Aston Martin	DBS	2003	2012
Aston Martin	DBS 770 Ultimate	2023	-
Aston Martin	DBS Superleggera	2018	-
Aston Martin	DBS Violante	2007	2012
Aston Martin	DBX	2020	-
Aston Martin	Rapide	2010	2019
Aston Martin	V12 Vanquish	2001	2008
Aston Martin	V12 Vantage	2010	-
Aston Martin	V8 Vantage	2005	2007
Aston Martin	Valkyrie	2019	2020
Aston Martin	Valour	2023	-
Aston Martin	Vanquish	2011	2018
Aston Martin	Vantage	2018	-
Aston Martin	Virage	2011	2012
Aston Martin	Zagato Coupe	2016	2016
Audi	A1	2010	-
Audi	A1 Allstreet	2022	-
Audi	A1 Citycarver	2019	2022
Audi	A2	2000	2006
Audi	A3	1996	-
Audi	A3 Allstreet	2024	-
Audi	A4	1995	-
Audi	A4 Allroad Quattro	2009	-
Audi	A5	2007	-
Audi	A6	1998	-
Audi	A7	2010	-
Audi	A8	1994	-
Audi	Allroad	2000	2011
Audi	e-tron	2019	-
Audi	e-tron GT	2020	-
Audi	e-tron S	2020	-
Audi	e-tron S Sportback	2020	-
Audi	e-tron Sportback	2020	-
Audi	Q2	2016	-
Audi	Q2L	2019	-
Audi	Q3	2011	-
Audi	Q3 Sportback	2019	-
Audi	Q4	2021	-
Audi	Q4 Sportback	2022	-
Audi	Q5	2008	-
Audi	Q5 Sportback	2021	-
Audi	Q6	2022	-
Audi	Q6 e-tron	2024	-
Audi	Q7	2006	-
Audi	Q8	2018	-
Audi	Q8 e-tron	2023	-
Audi	Q8 Sportback e-tron	2023	-
Audi	R8	2006	-
Audi	RS e-tron GT	2021	-
Audi	RS Q3	2012	-
Audi	RS Q3 Sportback	2019	-
Audi	RS Q7	0	-
Audi	RS Q8	2019	-
Audi	RS3	2011	-
Audi	RS4	2005	-
Audi	RS5	2010	-
Audi	RS6	2002	-
Audi	RS7	2013	-
Audi	S e-tron GT	2024	-
Audi	S1	2014	2018
Audi	S3	2006	-
Audi	S4	2003	-
Audi	S5	2007	-
Audi	S6	1999	-
Audi	S7	2012	-
Audi	S8	1996	-
Audi	SQ2	2019	-
Audi	SQ5	2012	-
Audi	SQ5 Sportback	2021	-
Audi	SQ6 e-tron	2024	-
Audi	SQ7	2016	-
Audi	SQ8	2019	-
Audi	SQ8 e-tron	2023	-
Audi	SQ8 Sportback e-tron	2023	-
Audi	TT	1998	-
Audi	TT RS	2008	-
Audi	TTS	2008	-
Aurus	Senat	2021	-
Avatr	11	2022	-
Avatr	12	2023	-
BAIC	BJ40	2021	-
BAIC	BJ40 Plus	2018	-
BAIC	U5 Plus	2021	-
BAIC	X35	2020	-
BAIC	X55	2020	-
Bentley	Arnage	1999	2007
Bentley	Azure	1999	2009
Bentley	Bentayga	2015	-
Bentley	Brooklands	2008	2011
Bentley	Continental	2003	2020
Bentley	Continental Flying Spur	2005	2013
Bentley	Continental GT	2011	-
Bentley	Continental GTC	2011	-
Bentley	Flying Spur	2013	-
Bentley	Mulsanne	2010	2020
BMW	1 series	2004	-
BMW	2 series	2013	-
BMW	3 series	1982	-
BMW	4 series	2013	-
BMW	5 series	1988	-
BMW	6 series	2004	-
BMW	7 series	1995	-
BMW	8 series	1990	-
BMW	i3	2013	-
BMW	i4	2021	-
BMW	i5	2023	-
BMW	i7	2022	-
BMW	i8	2014	2020
BMW	iX	2021	-
BMW	iX1	2022	-
BMW	iX2	2024	-
BMW	iX3	2020	-
BMW	M2	2015	-
BMW	M3	1992	-
BMW	M4	2012	-
BMW	M5	1998	-
BMW	M6	2005	-
BMW	M8	2019	-
BMW	X1	2009	-
BMW	X2	2017	-
BMW	X3	2004	-
BMW	X3 M	2019	-
BMW	X4	2014	-
BMW	X4 M	2019	-
BMW	X5	2000	-
BMW	X5 M	2009	-
BMW	X6	2008	-
BMW	X6 M	2009	-
BMW	X7	2018	-
BMW	XM	2023	-
BMW	Z3	1996	2003
BMW	Z4	2002	-
BMW	Z8	2000	2006
Brilliance	H230	2012	2019
Brilliance	V3	2014	2020
Brilliance	V5	2011	2020
Bugatti	Chiron	2015	2022
Bugatti	Divo	2019	2021
Bugatti	Veyron	2005	2015
Bugatti	W16 Mistral	2022	-
Buick	Century	1997	2005
Buick	Enclave	2007	-
Buick	Encore	2020	-
Buick	Envision	2014	-
Buick	GL8 ES	2021	-
Buick	La Crosse	2004	2016
Buick	LaCrosse	2016	-
Buick	Le Sabre	1999	2005
Buick	Lucerne	2005	2011
Buick	Park Avenue	1996	2012
Buick	Rainier	2003	2007
Buick	Regal	2009	-
Buick	Rendezvouz	2001	2007
Buick	Terraza	2005	2007
Buick	Verano	2012	2016
BYD	Dolphin	2021	-
BYD	Dolphin Plus	2023	-
BYD	Han	2022	-
BYD	Qin	2014	2020
BYD	Seal	2022	-
Cadillac	ATS	2012	2019
Cadillac	ATS-V	2016	2019
Cadillac	BLS	2006	2010
Cadillac	CT4	2020	-
Cadillac	CT4-V	2020	-
Cadillac	CT5	2019	-
Cadillac	CT5-V	2020	-
Cadillac	CT6	2016	-
Cadillac	CTS	2002	2019
Cadillac	De Ville	2000	2005
Cadillac	DTS	2006	2011
Cadillac	Eldorado	1992	2002
Cadillac	ELR	2013	2016
Cadillac	Escalade	2002	-
Cadillac	Lyriq	2023	-
Cadillac	Seville	1998	2004
Cadillac	SRX	2004	2016
Cadillac	STS	2004	2012
Cadillac	XLR	2003	2009
Cadillac	XT4	2018	-
Cadillac	XT5	2016	-
Cadillac	XT6	2019	-
Cadillac	XTS	2012	2019
Changan	Alsvin	2018	-
Changan	CS35	2012	2021
Changan	CS35 Plus	2018	-
Changan	CS55	2017	-
Changan	CS55 Plus	2019	-
Changan	CS75	2013	-
Changan	CS75 Plus	2019	-
Changan	CS85 Coupe	2019	-
Changan	CS95	2017	-
Changan	Eado	2011	-
Changan	Eado Plus	2020	-
Changan	Hunter Plus	2023	-
Changan	Lamore	2023	-
Changan	Raeton	2013	2017
Changan	Raeton CC	2017	-
Changan	Raeton Plus	2021	-
Changan	Uni-K	2021	-
Changan	Uni-T	2020	-
Changan	Uni-V	2021	-
Chery	Amulet	2003	2010
Chery	Arrizo 5 Plus	2021	-
Chery	Arrizo 6	2021	-
Chery	Arrizo 6 Pro	2021	-
Chery	Arrizo 7	2013	2020
Chery	Arrizo 8	2022	-
Chery	Bonus	2011	2016
Chery	Bonus 3	2013	2016
Chery	CrossEastar	2006	2015
Chery	Eastar	2003	2015
Chery	eQ	2015	2020
Chery	eQ1	2017	2021
Chery	eQ5	2020	-
Chery	eQ7	2023	-
Chery	Fora	2006	2010
Chery	iCar	2022	-
Chery	IndiS	2011	2016
Chery	Kimo	2008	-
Chery	M11	2008	2017
Chery	Omoda 5	2022	-
Chery	QQ	2003	-
Chery	QQ3	2003	2018
Chery	QQ6	2006	2010
Chery	Tiggo	2006	2016
Chery	Tiggo 2	2017	-
Chery	Tiggo 2 Pro	2020	-
Chery	Tiggo 3	2014	-
Chery	Tiggo 4	2017	-
Chery	Tiggo 4 Pro	2021	-
Chery	Tiggo 5	2014	-
Chery	Tiggo 7	2016	-
Chery	Tiggo 7 Pro	2020	-
Chery	Tiggo 7 Pro Max	2023	-
Chery	Tiggo 8	2018	-
Chery	Tiggo 8 Plus	2020	-
Chery	Tiggo 8 Pro	2021	-
Chery	Tiggo 8 Pro Max	2021	-
Chery	Tiggo 9	2023	-
Chery	Tiggo e	2019	2021
Chery	Very	2011	2016
Chevrolet	Astro	1995	2005
Chevrolet	Avalanche	2002	2013
Chevrolet	Aveo	2004	2020
Chevrolet	Beat	2010	2022
Chevrolet	Blazer	1994	2005
Chevrolet	Bolt	2017	-
Chevrolet	Bolt EUV	2021	-
Chevrolet	Camaro	1998	-
Chevrolet	Captiva	2006	2018
Chevrolet	Cavalier	1995	2005
Chevrolet	Cobalt	2005	-
Chevrolet	Colorado	2004	-
Chevrolet	Corvette	2005	-
Chevrolet	Cruze	2009	-
Chevrolet	Epica	2006	2012
Chevrolet	Equinox	2005	-
Chevrolet	Express	2003	-
Chevrolet	HHR	2006	2011
Chevrolet	Impala	2006	2020
Chevrolet	Lacetti	2002	2012
Chevrolet	Lanos	1997	2012
Chevrolet	Malibu	2004	-
Chevrolet	Monte Carlo	2005	2007
Chevrolet	Niva	2002	2020
Chevrolet	Orlando	2011	2018
Chevrolet	Rezzo	2000	2012
Chevrolet	Silverado	1999	-
Chevrolet	Silverado 1500	1999	-
Chevrolet	Silverado 2500 HD	2001	-
Chevrolet	Spark	2005	-
Chevrolet	SSR	2003	2006
Chevrolet	Suburban	2001	-
Chevrolet	Tahoe	2000	-
Chevrolet	TrailBlazer	2002	-
Chevrolet	Traverse	2009	-
Chevrolet	Trax	2013	-
Chevrolet	Uplander	2005	2009
Chevrolet	Venture	1996	2005
Chrysler	200	2010	2017
Chrysler	300	2004	-
Chrysler	300M	1998	2004
Chrysler	Aspen	2006	2009
Chrysler	Concorde	1998	2004
Chrysler	Crossfire	2003	2007
Chrysler	Grand Caravan	2021	-
Chrysler	Grand Voyager	2008	2015
Chrysler	Pacifica	2004	-
Chrysler	PT Cruiser	2000	2010
Chrysler	Sebring	2000	2010
Chrysler	Town & Country	2008	2016
Chrysler	Voyager	1995	2022
Citroen	Berlingo	1996	-
Citroen	C-Crosser	2007	2012
Citroen	C-Elysee	2012	2020
Citroen	C1	2005	2022
Citroen	C2	2003	2009
Citroen	C3	2002	-
Citroen	C3 Aircross	2017	-
Citroen	C3 Picasso	2009	2017
Citroen	C3 Pluriel	2003	2010
Citroen	C4	2004	-
Citroen	C4 Aircross	2012	2017
Citroen	C4 Cactus	2014	-
Citroen	C4 Picasso	2006	2021
Citroen	C4 SpaceTourer	2018	2021
Citroen	C4 X	2022	-
Citroen	C5	2001	-
Citroen	C5 Aircross	2017	-
Citroen	C5 X	2021	-
Citroen	C6	2004	2012
Citroen	C8	2002	2014
Citroen	DS 7 Crossback	2018	-
Citroen	DS 9	2020	-
Citroen	DS3	2009	2019
Citroen	DS4	2010	2018
Citroen	DS5	2011	2018
Citroen	Grand C4 Picasso	2006	2021
Citroen	Grand C4 SpaceTourer	2018	2022
Citroen	Jumper	2014	-
Citroen	Jumpy	2007	2016
Citroen	Nemo	2008	2018
Citroen	Saxo	1996	2004
Citroen	Spacetourer	2016	-
Citroen	Xsara	1997	2006
Citroen	Xsara Picasso	1999	2012
Dacia	Dokker	2012	2021
Dacia	Jogger	2021	-
Dacia	Lodgy	2012	2022
Dacia	Solenza	2003	2005
Dacia	Spring	2021	-
Dacia	SupeRNova	2000	2003
Daewoo	Evanda	2002	2007
Daewoo	Kalos	2002	2007
Daewoo	Leganza	1997	2002
Daewoo	Magnus	2000	2006
Daewoo	Matiz	2000	2015
Daewoo	Nexia	1995	2016
Daewoo	Nubira	1997	2008
Daihatsu	Applause	1997	2000
Daihatsu	Cast	2015	2023
Daihatsu	Copen	2002	-
Daihatsu	Cuore	2003	2011
Daihatsu	Gran Move	1996	2002
Daihatsu	Luxio	2009	-
Daihatsu	Materia	2006	2012
Daihatsu	Mebius	2013	2021
Daihatsu	Move	2014	-
Daihatsu	Rocky	2019	-
Daihatsu	Sirion	1998	2004
Daihatsu	Terios	1997	2017
Daihatsu	Trevis	2006	2011
Daihatsu	YRV	2000	2005
Datsun	mi-DO	2014	2021
Datsun	on-DO	2014	2021
Dodge	Avenger	1994	-
Dodge	Caliber	2006	2012
Dodge	Caliber SRT4	2007	2010
Dodge	Caravan	1995	2016
Dodge	Challenger	2008	-
Dodge	Charger	2006	-
Dodge	Dakota	2005	2011
Dodge	Dart	2012	2017
Dodge	Durango	1998	-
Dodge	Hornet	2023	-
Dodge	Intrepid	1998	2004
Dodge	Journey	2008	2020
Dodge	Magnum	2004	2008
Dodge	Neon	2000	2020
Dodge	Nitro	2007	2012
Dodge	Ram 1500	2001	-
Dodge	Ram 2500	2002	-
Dodge	Ram 3500	2002	-
Dodge	Ram SRT10	2004	2007
Dodge	Stratus	2001	2006
Dodge	Viper	1996	2017
Dongfeng	580	2016	-
Dongfeng	A30	2014	2017
Dongfeng	Aeolus AX7	2022	-
Dongfeng	Aeolus Yixuan MAX	2021	-
Dongfeng	AX7	2014	2020
Dongfeng	Fengon 5	2020	-
Dongfeng	Fengon 7	2020	-
Dongfeng	H30 Cross	2013	2018
Dongfeng	Rich 6	2018	-
Exeed	TXL	2019	-
Exeed	VX	2020	-
FAW	Bestune B70	2020	-
FAW	Bestune T55	2021	-
FAW	Bestune T77	2019	-
FAW	Bestune T99	2020	-
FAW	Besturn B30	2016	2020
FAW	Besturn B50	2009	2016
FAW	Besturn X40	2017	-
FAW	Besturn X80	2014	-
FAW	Oley	2012	2018
FAW	Vita	2006	2012
Ferrari	296	2022	-
Ferrari	348	1993	1995
Ferrari	360	1999	2005
Ferrari	456	1992	2003
Ferrari	458	2009	2016
Ferrari	488	2015	2020
Ferrari	512	1991	2004
Ferrari	550	1996	2002
Ferrari	575 M	2002	2006
Ferrari	599 GTB Fiorano	2006	2012
Ferrari	599 GTO	2010	2012
Ferrari	612	2004	2011
Ferrari	812	2017	-
Ferrari	California	2008	2014
Ferrari	California T	2014	2017
Ferrari	Challenge Stradale	2003	2003
Ferrari	Enzo	2002	2004
Ferrari	F12	2012	2017
Ferrari	F355	1994	1999
Ferrari	F430	2004	2010
Ferrari	F50	1995	1997
Ferrari	F512 M	1994	1996
Ferrari	F8 Spider	2019	2023
Ferrari	F8 Tributo	2019	2023
Ferrari	FF	2011	2017
Ferrari	GTC4 Lusso	2016	2020
Ferrari	LaFerrari	2013	2015
Ferrari	Portofino	2017	2021
Ferrari	Portofino M	2020	-
Ferrari	Roma	2020	-
Ferrari	SF90 Spider	2020	-
Ferrari	SF90 Stradale	2019	-
Fiat	124 Spider	2015	2020
Fiat	500	2007	-
Fiat	500L	2012	2022
Fiat	500X	2014	-
Fiat	600e	2023	-
Fiat	695	2023	-
Fiat	Albea	2002	2012
Fiat	Argo	2017	-
Fiat	Brava	1995	2003
Fiat	Bravo	1995	2016
Fiat	Coupe	1993	2009
Fiat	Croma	2005	2011
Fiat	Doblo	2001	-
Fiat	Ducato	2006	-
Fiat	E-Doblo	2022	-
Fiat	E-Ulysse	2022	-
Fiat	Egea	2015	-
Fiat	Freemont	2011	2016
Fiat	Grande Punto	2005	2009
Fiat	Idea	2004	2016
Fiat	Linea	2007	2015
Fiat	Marea	1996	2006
Fiat	Multipla	1998	2010
Fiat	Palio	1996	2017
Fiat	Panda	2003	-
Fiat	Panda 4x4	2003	2012
Fiat	Panda Cross	2014	-
Fiat	Pulse	2021	-
Fiat	Punto	1999	2018
Fiat	Qubo	2008	2021
Fiat	Sedici	2005	2014
Fiat	Siena	2014	2021
Fiat	Stilo	2001	2010
Fiat	Strada	2014	-
Fiat	Tipo	2015	-
Fiat	Ulysse	2002	2010
Fisker	Karma	2011	2012
Ford	B-Max	2012	2017
Ford	Bronco	2021	-
Ford	Bronco Sport	2021	-
Ford	C-Max	2003	2019
Ford	Cougar	1998	2004
Ford	Crown Victoria	1999	2011
Ford	EcoSport	2003	-
Ford	Edge	2007	-
Ford	Endura	2018	2021
Ford	Equator	2021	-
Ford	Escape	2001	-
Ford	Excursion	1999	2005
Ford	Expedition	1998	-
Ford	Explorer	1995	-
Ford	Explorer Sport Trac	2007	2010
Ford	F-150	1992	-
Ford	F-150 Lightning	2022	-
Ford	F-250	2011	-
Ford	F-350	2017	-
Ford	Falcon	2005	-
Ford	Fiesta	2002	-
Ford	Fiesta Active	2018	-
Ford	Five Hundred	2005	2007
Ford	Flex	2008	2019
Ford	Focus	1998	-
Ford	Focus Active	2018	-
Ford	Focus Electric	2012	2018
Ford	Freestar	2003	2009
Ford	Freestyle	2004	2007
Ford	Fusion	2002	2012
Ford	Galaxy	1995	-
Ford	Ka	2003	2021
Ford	Kuga	2008	-
Ford	Maverick	1992	2009
Ford	Mondeo	2000	-
Ford	Mustang	1995	-
Ford	Mustang Mach-E	2020	-
Ford	Mustang Shelby GT350	2016	-
Ford	Mustang Shelby GT500	2020	-
Ford	Puma	1997	-
Ford	Ranger	2006	-
Ford	S-Max	2006	-
Ford	Taurus	2000	-
Ford	Taurus X	2007	2009
Ford	Thunderbird	2000	2006
Ford	Tourneo Connect	2002	-
Ford	Transit	2001	-
Ford	Transit Connect	2002	-
Foton	Sauvana	2015	2020
Foton	Tunland G7	2019	-
GAC	GM8	2018	2023
GAC	GN8	2020	-
GAC	GS5	2011	-
GAC	GS8	2017	-
GAZ	3102	1997	2009
GAZ	31105	1997	2009
GAZ	Gazelle	1994	2010
GAZ	Gazelle Business	2010	-
GAZ	Gazelle Next	2013	-
GAZ	Gazelle NN	2019	-
GAZ	Gazelle Sity	2020	-
GAZ	Siber	2008	2010
GAZ	Sobol	2002	-
GAZ	Sobol NN	2023	-
Geely	Atlas	2017	-
Geely	Atlas Pro	2021	-
Geely	Azkarra	2019	-
Geely	Belgee X50	2023	-
Geely	Coolray	2018	-
Geely	Emgrand	2021	-
Geely	Emgrand 7	2016	2022
Geely	Emgrand EC7	2009	2017
Geely	Emgrand GS	2016	2022
Geely	Emgrand GSe	2018	2022
Geely	Emgrand X7	2011	2021
Geely	GC9	2015	2019
Geely	Geometry C	2019	-
Geely	GС6	2014	2016
Geely	MK	2008	2019
Geely	Monjaro	2021	-
Geely	Okavango	2020	-
Geely	Otaka	2006	2009
Geely	Preface	2020	-
Geely	Tugella	2019	-
Geely	Vision	2006	2013
Genesis	G70	2018	-
Genesis	G80	2017	-
Genesis	G90	2016	-
Genesis	GV60	2021	-
Genesis	GV70	2020	-
Genesis	GV80	2020	-
GMC	Acadia	2007	-
GMC	Canyon	2004	-
GMC	Envoy	2001	2009
GMC	Hummer EV	2022	-
GMC	Sierra 1500	1999	-
GMC	Sierra 2500	1999	-
GMC	Sierra 3500	2020	-
GMC	Terrain	2010	-
GMC	Yukon	2000	-
Great Wall	Cowry	2007	2014
Great Wall	Deer	2003	2013
Great Wall	Hover	2005	2020
Great Wall	Hover M2	2010	2014
Great Wall	Pegasus	2005	2007
Great Wall	Peri	2007	2015
Great Wall	Poer	2020	-
Great Wall	Poer King Kong	2022	-
Great Wall	Safe	2002	2010
Great Wall	Sailor	2004	2012
Great Wall	Sing	2004	2012
Great Wall	Socool	2004	2012
Great Wall	Wingle	2008	2021
Great Wall	Wingle 7	2018	2021
Haval	Cool Dog	2022	-
Haval	Dargo	2020	-
Haval	Dargo X	2023	-
Haval	F7	2018	-
Haval	F7x	2019	-
Haval	H4	2018	2021
Haval	H6	2017	-
Haval	H9	2015	-
Haval	Jolion	2021	-
Haval	M6 Plus	2021	-
Holden	Commodore	2004	2017
Holden	Corvette C8	2020	2021
Honda	Accord	1998	-
Honda	Amaze	2011	-
Honda	City	2014	-
Honda	Civic	2001	-
Honda	Civic Type R	2015	-
Honda	CR-V	1995	-
Honda	CR-Z	2010	2016
Honda	Crosstour	2010	2015
Honda	e	2020	-
Honda	Edix	2004	2009
Honda	Element	2003	2011
Honda	Fit	2001	2022
Honda	FR-V	2004	2009
Honda	HR-V	1998	2022
Honda	Insight	1999	-
Honda	Jade	2013	2020
Honda	Jazz	2001	-
Honda	Jazz Crosstar	2020	-
Honda	Legend	2006	2021
Honda	Odyssey	1999	2017
Honda	Passport	2019	-
Honda	Pilot	2002	-
Honda	Prelude	1996	2001
Honda	Ridgeline	2005	2014
Honda	S2000	1999	2009
Honda	Shuttle	1994	2002
Honda	Stepwgn	2022	-
Honda	Stream	2001	2015
Honda	Vezel	2013	-
Hongqi	E-HS9	2020	-
Hongqi	H5	2022	-
Hummer	H1	1992	2006
Hummer	H2	2003	2010
Hummer	H3	2005	-
Hyundai	Accent	2000	-
Hyundai	Atos Prime	1999	2014
Hyundai	Azera	2006	2011
Hyundai	Bayon	2021	-
Hyundai	Centennial	1999	2008
Hyundai	Creta	2015	-
Hyundai	Creta Grand	2021	-
Hyundai	Elantra	2000	-
Hyundai	Entourage	2007	2007
Hyundai	Eon	2011	2019
Hyundai	Equus	2009	2016
Hyundai	Galloper	1998	2003
Hyundai	Genesis	2008	2016
Hyundai	Genesis Coupe	2008	2016
Hyundai	Getz	2002	2011
Hyundai	Grandeur	2005	2017
Hyundai	H-1	1998	-
Hyundai	i10	2008	-
Hyundai	i20	2008	-
Hyundai	i20 N	2021	-
Hyundai	i30	2007	-
Hyundai	i30 N	2017	-
Hyundai	i40	2011	2019
Hyundai	Ioniq	2016	-
Hyundai	Ioniq 5	2021	-
Hyundai	Ioniq 6	2022	-
Hyundai	ix20	2010	2020
Hyundai	ix35	2010	2018
Hyundai	Kona	2017	-
Hyundai	Kona N	2021	-
Hyundai	Kusto	2021	-
Hyundai	Matrix	2001	2014
Hyundai	Mistra	2021	-
Hyundai	Nexo	2018	-
Hyundai	Palisade	2018	-
Hyundai	Porter	1996	-
Hyundai	Santa Cruz	2022	-
Hyundai	Santa Fe	2001	-
Hyundai	Solaris	2010	-
Hyundai	Sonata	1998	-
Hyundai	Staria	2021	-
Hyundai	Terracan	2001	2007
Hyundai	Trajet	2000	2008
Hyundai	Tucson	2004	-
Hyundai	Veloster	2011	-
Hyundai	Venue	2019	-
Hyundai	Veracruz	2007	2015
Hyundai	Verna	2005	2011
Hyundai	Xcent	2014	2021
Hyundai	XG	1998	2005
Infiniti	EX	2008	2014
Infiniti	FX	2003	2014
Infiniti	G	2003	2014
Infiniti	I35	2002	2004
Infiniti	JX	2012	2013
Infiniti	M	2001	2013
Infiniti	Q30	2016	2021
Infiniti	Q40	2014	2015
Infiniti	Q45	2002	2007
Infiniti	Q50	2013	-
Infiniti	Q60	2013	-
Infiniti	Q70	2013	2021
Infiniti	QX30	2016	2020
Infiniti	QX4	1997	2003
Infiniti	QX50	2013	-
Infiniti	QX55	2021	-
Infiniti	QX56	2004	2013
Infiniti	QX60	2013	-
Infiniti	QX70	2013	2017
Infiniti	QX80	2013	-
Isuzu	Ascender	2002	2008
Isuzu	Axiom	2002	2004
Isuzu	D-Max	2011	-
Isuzu	D-Max Rodeo	2007	2007
Isuzu	I280	2005	2006
Isuzu	I290	2007	2007
Isuzu	I350	2005	2006
Isuzu	I370	2007	2007
Isuzu	mu-X	2020	-
Isuzu	Rodeo	1998	2004
Isuzu	Trooper	1992	2002
Isuzu	VehiCross	1997	2001
Iveco	Daily	2006	-
Jac	iEV7S	2017	-
Jac	J7	2020	-
Jac	JS3	2021	-
Jac	JS4	2021	-
Jac	JS6	2022	-
Jac	S3 Pro	2020	-
Jac	T6	2015	-
Jac	T8	2018	-
Jac	T8 Pro	2020	-
Jac	T9	2023	-
Jaecoo	J8	2024	-
Jaguar	E-Pace	2017	-
Jaguar	F-Pace	2015	-
Jaguar	F-Type	2013	-
Jaguar	I-Pace	2018	-
Jaguar	S-Type	1998	2008
Jaguar	X-Type	2001	2009
Jaguar	XE	2015	-
Jaguar	XF	2007	-
Jaguar	XJ	1997	-
Jaguar	XK/XKR	2002	2015
Jeep	Cherokee	2002	-
Jeep	Commander	2006	-
Jeep	Compass	2007	-
Jeep	Gladiator	2019	-
Jeep	Grand Cherokee	1999	-
Jeep	Grand Wagoneer	2022	-
Jeep	Liberty	2006	2013
Jeep	Meridian	2021	-
Jeep	Patriot	2006	2016
Jeep	Renegade	2014	-
Jeep	Wagoneer	2021	-
Jeep	Wrangler	1997	-
Jetour	Dashing	2022	-
Jetour	Traveler	2023	-
Jetour	X70 Plus	2022	-
Jetour	X90 Plus	2021	-
Kaiyi	E5	2021	-
Kia	Carens	2002	-
Kia	Carnival	1999	-
Kia	Ceed	2006	-
Kia	Cerato	2004	-
Kia	Clarus	1998	2001
Kia	EV6	2021	-
Kia	EV9	2023	-
Kia	Forte	2008	-
Kia	K5	2020	-
Kia	K8	2021	-
Kia	K9	2021	-
Kia	K900	2014	-
Kia	Magentis	2001	2010
Kia	Mohave	2008	-
Kia	Niro	2016	-
Kia	Opirus	2003	2010
Kia	Optima	2010	2020
Kia	Picanto	2004	-
Kia	ProCeed	2019	-
Kia	Quoris	2012	2018
Kia	Ray	2011	-
Kia	Rio	2000	-
Kia	Rio X	2020	-
Kia	Rio X-Line	2017	2021
Kia	Seltos	2019	-
Kia	Shuma	1997	2005
Kia	Sonet	2020	-
Kia	Sorento	2002	-
Kia	Sorento Prime	2015	2020
Kia	Soul	2009	-
Kia	Spectra	2005	2009
Kia	Sportage	1993	-
Kia	Stinger	2017	-
Kia	Stonic	2017	-
Kia	Telluride	2019	-
Kia	Venga	2009	2019
Kia	XCeed	2019	-
Lamborghini	Aventador	2011	2022
Lamborghini	Centenario	2016	2017
Lamborghini	Diablo	1991	2001
Lamborghini	Gallardo	2003	2014
Lamborghini	Huracan	2014	-
Lamborghini	Murcielago	2003	2010
Lamborghini	Reventon	2008	2008
Lamborghini	Sian	2021	-
Lamborghini	Urus	2018	-
Lancia	Delta	2008	2014
Lancia	Lybra	1999	2007
Lancia	Musa	2004	2012
Lancia	Phedra	2002	2010
Lancia	Thema	2011	2014
Lancia	Thesis	2001	2009
Lancia	Ypsilon	2003	-
Land Rover	Defender	2007	-
Land Rover	Discovery	1998	-
Land Rover	Discovery Sport	2014	-
Land Rover	Evoque	2011	-
Land Rover	Freelander	1998	2014
Land Rover	Range Rover	1994	-
Land Rover	Range Rover Sport	2004	-
Land Rover	Range Rover Velar	2017	-
Lexus	CT	2010	2022
Lexus	ES	2001	-
Lexus	GS	1997	2022
Lexus	GX	2002	-
Lexus	HS	2009	-
Lexus	IS	1998	-
Lexus	LBX	2023	-
Lexus	LC	2017	-
Lexus	LFA	2010	2012
Lexus	LM	2019	-
Lexus	LS	1995	-
Lexus	LX	1998	-
Lexus	NX	2014	-
Lexus	RC	2014	-
Lexus	RC F	2014	-
Lexus	RX	1997	-
Lexus	RZ	2023	-
Lexus	SC	1999	2010
Lexus	TX	2023	-
Lexus	UX	2018	-
Lifan	Breez	2007	2012
Lifan	Cebrium	2013	2017
Lifan	Celliya	2013	2016
Lifan	Murman	2017	-
Lifan	Myway	2017	-
Lifan	Smily	2008	2018
Lifan	Solano	2008	-
Lifan	X50	2014	-
Lifan	X60	2011	-
Lifan	X70	2018	-
Lifan	X80	2017	2020
Lincoln	Aviator	2003	-
Lincoln	Corsair	2019	-
Lincoln	Mark LT	2006	2007
Lincoln	MKC	2014	2019
Lincoln	MKS	2008	2016
Lincoln	MKT	2009	2019
Lincoln	MKX	2006	2018
Lincoln	MKZ	2006	2020
Lincoln	Nautilus	2019	-
Lincoln	Navigator	1997	-
Lincoln	Town Car	1998	2011
Lincoln	Zephyr	2006	2006
Livan	S6 Pro	2023	-
Livan	X3 Pro	2022	-
Livan	X6 Pro	2023	-
LiXiang	L6	2024	-
LiXiang	L7	2023	-
LiXiang	L8	2022	-
LiXiang	L9	2022	-
LiXiang	Mega	2024	-
Lotus	Elise	2001	-
Lotus	Emira	2022	-
Lotus	Europa S	2005	2010
Lotus	Evora	2009	-
Lotus	Exige	2001	2021
Marussia	B1	2008	2014
Marussia	B2	2010	2014
Maserati	3200 GT	1998	2002
Maserati	Ghibli	2013	-
Maserati	Gran Cabrio	2010	2019
Maserati	Gran Turismo 	2007	2019
Maserati	Gran Turismo S	2008	2012
Maserati	Grecale	2022	-
Maserati	Levante	2016	-
Maserati	MC20	2020	-
Maserati	Quattroporte	2003	-
Maserati	Quattroporte S	2007	2012
Maybach	57	2002	2012
Maybach	57 S	2007	2012
Maybach	62	2002	2012
Maybach	62 S	2007	2012
Maybach	Landaulet	2003	2012
Mazda	2	2003	-
Mazda	2 Hybrid	2021	-
Mazda	3	2003	-
Mazda	323	1998	2003
Mazda	5	2005	-
Mazda	6	2002	-
Mazda	626	1997	2001
Mazda	B-Series	1999	2006
Mazda	BT-50	2011	-
Mazda	CX-3	2015	-
Mazda	CX-30	2019	-
Mazda	CX-30 EV	2021	-
Mazda	CX-4	2017	-
Mazda	CX-5	2012	-
Mazda	CX-50	2022	-
Mazda	CX-60	2022	-
Mazda	CX-7	2006	2012
Mazda	CX-8	2017	-
Mazda	CX-80	2024	-
Mazda	CX-9	2007	-
Mazda	MPV	1999	2016
Mazda	MX-30	2020	-
Mazda	MX-5	1998	-
Mazda	Premacy	1999	2006
Mazda	RX-7	1992	2002
Mazda	RX-8	2003	2012
Mazda	Tribute	2000	2007
McLaren	540C	2015	2021
McLaren	570S	2015	2021
McLaren	600LT	2018	2021
McLaren	650S	2014	2017
McLaren	675LT	2015	2017
McLaren	720S	2017	-
McLaren	720S Spider	2019	2022
McLaren	765LT	2020	-
McLaren	Artura	2021	-
McLaren	MP4-12C	2011	2014
McLaren	P1	2013	2016
Mercedes	A-class	1997	-
Mercedes	AMG GT	2014	-
Mercedes	AMG GT 4-Door	2018	-
Mercedes	B-class	2005	-
Mercedes	C-class	1997	-
Mercedes	C-class Sport Coupe	2001	2007
Mercedes	Citan	2012	-
Mercedes	CL-class	1992	2014
Mercedes	CLA-class	2013	-
Mercedes	CLC-class 	2008	2011
Mercedes	CLK-class	1997	2010
Mercedes	CLS-class	2004	-
Mercedes	E-class	1995	-
Mercedes	E-class Coupe	2017	-
Mercedes	EQA	2021	-
Mercedes	EQB	2021	-
Mercedes	EQC	2019	-
Mercedes	EQE	2022	-
Mercedes	EQE AMG	2022	-
Mercedes	EQS	2021	-
Mercedes	EQS AMG	2022	-
Mercedes	EQT	2023	-
Mercedes	EQV	2021	-
Mercedes	G-class	1996	-
Mercedes	GL-class	2006	2016
Mercedes	GLA-class	2013	-
Mercedes	GLA-class AMG	2013	-
Mercedes	GLB-class	2019	-
Mercedes	GLC-class	2015	-
Mercedes	GLC-class AMG	2015	-
Mercedes	GLC-class Coupe	2016	-
Mercedes	GLE-class	2015	-
Mercedes	GLE-class AMG	2015	-
Mercedes	GLE-class Coupe	2015	-
Mercedes	GLK-class	2008	2015
Mercedes	GLS-class	2015	-
Mercedes	GLS-class AMG	2015	-
Mercedes	M-class	1997	-
Mercedes	R-class	2005	2012
Mercedes	S-class	1990	-
Mercedes	S-class Cabrio	2015	2021
Mercedes	S-class Coupe	2015	2021
Mercedes	SL-class	2001	2021
Mercedes	SL-Class AMG	2022	-
Mercedes	SLC-class	2018	-
Mercedes	SLK-class	1996	-
Mercedes	SLR-class	2003	2010
Mercedes	SLS AMG	2010	2014
Mercedes	Sprinter	2000	-
Mercedes	Vaneo	2001	2006
Mercedes	Viano	2003	-
Mercedes	Vito	1999	2003
Mercedes	X-class	2018	-
Mercury	Grand Marquis	2003	2011
Mercury	Mariner	2005	2007
Mercury	Milan	2006	2011
Mercury	Montego	2004	2007
Mercury	Monterey	1991	2007
Mercury	Mountaineer	2003	2010
Mercury	Sable	1995	2005
MG	4	2022	-
MG	Hector	2019	-
MG	TF	2002	2011
MG	VS	2022	-
MG	XPower SV	2003	2008
MG	ZR	2001	2005
MG	ZS	2001	2005
MG	ZS EV	2021	-
MG	ZT	2001	2005
MG	ZT-T	2001	2005
Mini	Aceman	2024	-
Mini	Clubman	2007	-
Mini	Clubman S	2007	2014
Mini	Clubvan	2012	2014
Mini	Cooper	2001	-
Mini	Cooper Cabrio	2001	-
Mini	Cooper S	2001	-
Mini	Cooper S Cabrio	2001	-
Mini	Cooper S Countryman All4	2010	-
Mini	Countryman	2010	-
Mini	One	2001	-
Mitsubishi	3000 GT	1992	2001
Mitsubishi	ASX	2010	-
Mitsubishi	Carisma	1995	2004
Mitsubishi	Colt	1995	2012
Mitsubishi	Dignity	2012	2017
Mitsubishi	Eclipse	1995	2007
Mitsubishi	Eclipse Cross	2017	-
Mitsubishi	Endeavor	2004	2011
Mitsubishi	Galant	1996	2012
Mitsubishi	Grandis	2003	2011
Mitsubishi	i-MiEV	2009	2020
Mitsubishi	L200	1996	-
Mitsubishi	Lancer	1996	2017
Mitsubishi	Lancer Evo	2001	2016
Mitsubishi	Mirage	2012	-
Mitsubishi	Outlander	2003	-
Mitsubishi	Outlander Sport	2010	-
Mitsubishi	Outlander XL	2006	2012
Mitsubishi	Pajero	1990	2021
Mitsubishi	Pajero Pinin	1999	2007
Mitsubishi	Pajero Sport	1998	-
Mitsubishi	Raider	2006	2007
Mitsubishi	Space Gear	1995	2007
Mitsubishi	Space Runner	1999	2004
Mitsubishi	Space Star	1998	2004
Mitsubishi	Xpander	2017	-
Moskvich	3e	2022	-
Nissan	350Z	2002	2009
Nissan	370Z	2009	-
Nissan	Almera	2000	-
Nissan	Almera Classic	2005	2012
Nissan	Almera Tino	2000	2005
Nissan	Altima	2002	-
Nissan	Ariya	2020	-
Nissan	Armada	2003	-
Nissan	Bluebird Sylphy	2000	2012
Nissan	e-NV200	2014	-
Nissan	Frontier	1998	-
Nissan	GT-R	2007	-
Nissan	Juke	2011	-
Nissan	Leaf	2010	-
Nissan	Maxima	2000	-
Nissan	Micra	2003	-
Nissan	Murano	2002	-
Nissan	Navara	2005	-
Nissan	Note	2005	-
Nissan	NP300	2008	2013
Nissan	Pathfinder	1997	-
Nissan	Patrol	1997	-
Nissan	Primera	1996	2008
Nissan	Qashqai	2007	-
Nissan	Qashqai+2	2008	2013
Nissan	Quest	2003	2017
Nissan	Rogue	2008	-
Nissan	Sentra	1999	-
Nissan	Skyline	2001	2005
Nissan	Sylphy	2012	2020
Nissan	Teana	2005	-
Nissan	Terrano	1993	-
Nissan	Tiida	2007	-
Nissan	Titan	2003	-
Nissan	Titan XD	2016	-
Nissan	Townstar	2022	-
Nissan	X-Trail	2001	-
Nissan	XTerra	2001	2015
Nissan	Z	2003	-
Omoda	S5 GT	2023	-
Opel	Adam	2013	2020
Opel	Agila	2000	2014
Opel	Ampera-e	2017	2021
Opel	Antara	2006	2015
Opel	Astra	1998	-
Opel	Astra GTC	2011	2018
Opel	Astra OPC	2011	2018
Opel	Cascada	2013	2019
Opel	Combo	2001	-
Opel	Combo Life	2021	-
Opel	Corsa	2000	-
Opel	Corsa OPC	2007	2018
Opel	Crossland	2020	-
Opel	Crossland X	2017	2020
Opel	Frontera	1998	2004
Opel	Grandland	2021	-
Opel	Grandland X	2017	-
Opel	Insignia	2008	-
Opel	Insignia OPC	2009	2017
Opel	Karl	2014	2019
Opel	Meriva	2003	2017
Opel	Mokka	2012	-
Opel	Omega	1994	2004
Opel	Signum	2003	2008
Opel	Speedster	2000	2007
Opel	Tigra	1994	2009
Opel	Vectra	1995	2008
Opel	Vivaro	2014	-
Opel	Zafira	1999	2014
Opel	Zafira Life	2019	-
Opel	Zafira Tourer	2012	2020
Peugeot	1007	2004	2009
Peugeot	107	2005	2014
Peugeot	108	2014	-
Peugeot	2008	2013	-
Peugeot	206	1998	2012
Peugeot	207	2006	2014
Peugeot	208	2012	-
Peugeot	3008	2009	-
Peugeot	301	2012	-
Peugeot	307	2000	2011
Peugeot	308	2007	-
Peugeot	4007	2007	2012
Peugeot	4008	2012	-
Peugeot	406	1995	2004
Peugeot	407	2004	2010
Peugeot	408	2010	-
Peugeot	5008	2009	-
Peugeot	508	2011	-
Peugeot	607	2000	2010
Peugeot	807	2002	2014
Peugeot	Boxer	2008	2014
Peugeot	Expert	2016	-
Peugeot	Landtrek	2020	-
Peugeot	Manager	2017	-
Peugeot	Partner	1996	-
Peugeot	Partner Crossway	2021	-
Peugeot	Partner Rapid	2021	-
Peugeot	RCZ Sport	2010	2015
Peugeot	Rifter	2018	-
Peugeot	Traveller	2016	-
Plymouth	Road Runner	1968	1970
Pontiac	Aztec	2001	2005
Pontiac	Bonneville	1999	2005
Pontiac	Firebird	1993	2002
Pontiac	G5 Pursuit	2004	2010
Pontiac	G6	2004	2010
Pontiac	G8	2008	2009
Pontiac	Grand AM	1998	2005
Pontiac	Grand Prix	1996	2008
Pontiac	GTO	2004	2006
Pontiac	Montana	1997	2009
Pontiac	Solstice	2005	2010
Pontiac	Sunfire	1995	2005
Pontiac	Torrent	2005	2009
Pontiac	Vibe	2002	2009
Porsche	718 Boxster	2016	-
Porsche	718 Cayman	2016	-
Porsche	911	1997	-
Porsche	Boxster	1996	2016
Porsche	Cayenne	2002	-
Porsche	Cayman	2005	2016
Porsche	Macan	2014	-
Porsche	Panamera	2009	-
Porsche	Taycan	2020	-
Ravon	Gentra	2015	2019
Renault	Alaskan	2016	2020
Renault	Arkana	2019	-
Renault	Austral	2022	-
Renault	Avantime	2001	2004
Renault	Captur	2013	-
Renault	Clio	1998	-
Renault	Duster	2010	-
Renault	Duster Oroch	2015	2022
Renault	Espace	1996	-
Renault	Fluence	2010	2019
Renault	Grand Scenic	2009	-
Renault	Kadjar	2015	-
Renault	Kangoo	1998	-
Renault	Kaptur	2016	2022
Renault	Kiger	2021	-
Renault	Koleos	2008	-
Renault	Laguna	1993	2015
Renault	Latitude	2010	2018
Renault	Logan	2004	-
Renault	Logan Stepway	2018	2022
Renault	Master	2011	-
Renault	Megane	1996	-
Renault	Megane E-Tech	2022	-
Renault	Modus	2004	2012
Renault	Rafale	2024	-
Renault	Sandero	2007	-
Renault	Sandero Stepway	2008	2022
Renault	Scenic	1996	-
Renault	Symbol	2002	-
Renault	Taliant	2021	-
Renault	Talisman	2015	2022
Renault	Trafic	2001	-
Renault	Triber	2019	-
Renault	Twingo	1993	-
Renault	Twizy	2012	-
Renault	Vel Satis	2002	2009
Renault	Wind	2010	2013
Renault	Zoe	2012	-
Rolls-Royce	Cullinan	2018	-
Rolls-Royce	Dawn	2015	2022
Rolls-Royce	Ghost	2009	-
Rolls-Royce	Phantom	2003	-
Rolls-Royce	Spectre	2023	-
Rolls-Royce	Wraith	2013	2022
Rover	25	1999	2008
Rover	400	1996	2000
Rover	45	2000	2008
Rover	600	1999	2004
Rover	75	1999	2005
Rover	Streetwise	2003	2005
Saab	9-2x	2004	2007
Saab	9-3	2003	2014
Saab	9-4x	2011	2012
Saab	9-5	1997	2012
Saab	9-7x	2005	2009
Saturn	Aura	2007	-
Saturn	Ion	2003	2007
Saturn	LW	1998	2004
Saturn	Outlook	2006	2010
Saturn	Sky	2007	2010
Saturn	Vue	2002	2007
Scion	FR-S	2012	2016
Scion	tC	2004	2016
Scion	xA	2003	2007
Scion	xB	2003	2015
Scion	xD	2007	2014
Seat	Alhambra	1998	-
Seat	Altea	2004	2015
Seat	Altea Freetrack	2007	2015
Seat	Altea XL	2007	2015
Seat	Arona	2017	-
Seat	Arosa	1997	2006
Seat	Ateca	2016	-
Seat	Cordoba	1999	2009
Seat	Exeo	2008	2013
Seat	Ibiza	2002	-
Seat	Leon	1999	-
Seat	Mii	2012	2021
Seat	Tarraco	2018	-
Seat	Toledo	2004	2018
Seres	Seres 3	2022	-
Skoda	Citigo	2011	2020
Skoda	Enyaq iV	2020	-
Skoda	Fabia	1999	-
Skoda	Felicia	1995	2000
Skoda	Kamiq	2019	-
Skoda	Karoq	2017	-
Skoda	Kodiaq	2016	-
Skoda	Octavia	1995	-
Skoda	Octavia Scout	2007	2013
Skoda	Octavia Tour	2006	2013
Skoda	Praktik	2007	2015
Skoda	Rapid	2012	-
Skoda	Rapid Spaceback (NH1)	2013	2020
Skoda	Roomster	2006	2015
Skoda	Scala	2019	-
Skoda	Superb	2002	-
Skoda	Yeti	2009	-
Smart	#1	2022	-
Smart	Forfour	2004	-
Smart	Fortwo	1997	-
Smart	Roadster	2003	-
Ssang Yong	Actyon	2005	-
Ssang Yong	Actyon Sports	2006	-
Ssang Yong	Chairman	1997	-
Ssang Yong	Korando	2002	-
Ssang Yong	Kyron	2005	2016
Ssang Yong	LUVi	2015	-
Ssang Yong	Musso	1993	-
Ssang Yong	Musso Grand	2020	-
Ssang Yong	Musso Sport	2004	2007
Ssang Yong	Rexton	2001	-
Ssang Yong	Rexton Sports	2018	-
Ssang Yong	Rodius	2004	2019
Ssang Yong	Stavic	2013	-
Ssang Yong	Tivoli	2015	-
Ssang Yong	Tivoli Grand	2021	-
Ssang Yong	XLV	2016	-
Subaru	Ascent	2018	-
Subaru	Baja	2002	2007
Subaru	BRZ	2012	-
Subaru	Crosstrek	2018	-
Subaru	Exiga	2008	2015
Subaru	Forester	1997	-
Subaru	Impreza	1992	-
Subaru	Impreza WRX	2001	2013
Subaru	Impreza WRX STI	2000	2013
Subaru	Justy	1995	2006
Subaru	Legacy	1994	-
Subaru	Levorg	2014	-
Subaru	Outback	1997	-
Subaru	Solterra	2022	-
Subaru	Traviq	2001	2006
Subaru	Tribeca	2005	2014
Subaru	WRX	2014	-
Subaru	XV	2012	-
Suzuki	Alto	2002	-
Suzuki	Baleno	1995	-
Suzuki	Celerio	2014	2021
Suzuki	Ciaz	2014	-
Suzuki	Ertiga	2012	-
Suzuki	Grand Vitara	1998	2014
Suzuki	Grand Vitara XL7	1998	2014
Suzuki	Ignis	2000	-
Suzuki	Jimny	1998	-
Suzuki	Kizashi	2009	2015
Suzuki	Liana	2002	2007
Suzuki	S-Presso	2019	-
Suzuki	Splash	2008	2016
Suzuki	Swift	2005	-
Suzuki	SX4	2006	2021
Suzuki	Vitara	2015	-
Suzuki	Wagon R	2012	2017
Suzuki	Wagon R+	2000	2006
Suzuki	XL6	2019	-
Suzuki	XL7	2020	-
Tank	700	2024	-
Tesla	Model 3	2017	-
Tesla	Model S	2012	-
Tesla	Model X	2015	-
Tesla	Model Y	2019	-
Toyota	4Runner	1995	-
Toyota	Alphard	2002	-
Toyota	Auris	2007	-
Toyota	Avalon	2000	-
Toyota	Avensis	1997	2018
Toyota	Avensis Verso	2001	2010
Toyota	Aygo	2005	-
Toyota	Aygo X	2021	-
Toyota	BZ4X	2022	-
Toyota	C+pod	2020	-
Toyota	C-HR	2016	-
Toyota	Caldina	2002	2007
Toyota	Camry	1996	-
Toyota	Celica	1999	2006
Toyota	Corolla	1995	-
Toyota	Corolla Cross	2020	-
Toyota	Corolla Verso	2002	2009
Toyota	Crown	2018	-
Toyota	FJ Cruiser	2007	2016
Toyota	Fortuner	2004	-
Toyota	GT 86	2012	-
Toyota	Harrier	2013	-
Toyota	Hiace	2006	-
Toyota	Highlander	2000	-
Toyota	Hilux	2005	-
Toyota	iQ	2008	2016
Toyota	ist	2002	2016
Toyota	Land Cruiser	1990	-
Toyota	Land Cruiser Prado	2002	-
Toyota	Mark II	2000	2004
Toyota	Mirai	2014	-
Toyota	MR2	1999	2007
Toyota	Picnic	1995	2001
Toyota	Previa	2000	-
Toyota	Prius	2003	-
Toyota	Prius Prime	2017	-
Toyota	RAV4	1994	-
Toyota	Sequoia	2001	-
Toyota	Sienna	2002	-
Toyota	Supra	2019	-
Toyota	Tacoma	2005	-
Toyota	Tundra	1999	-
Toyota	Venza	2008	2016
Toyota	Verso	2009	-
Toyota	Vitz	2005	-
Toyota	Wish	2003	2017
Toyota	Yaris	1999	-
Toyota	Yaris Verso	1999	2005
UAZ	3153 Gusar	1996	2013
UAZ	3159 Bars	1999	2005
UAZ	3160	1997	2005
UAZ	3162 Simbir	1999	2005
UAZ	Buhanka	1985	-
UAZ	Patriot Sport	2010	2013
UAZ	Pickup	2008	-
UAZ	Profi	2017	-
UAZ	Патриот	2005	-
UAZ	Хантер	2003	-
VAZ	2101-2107	1976	2012
VAZ	2108, 2109, 21099	1984	2004
VAZ	2110, 2111, 2112	1996	2014
VAZ	2113, 2114, 2115	1997	2015
VAZ	4x4 Urban	2014	2021
VAZ	Granta	2011	-
VAZ	Granta Cross	2019	-
VAZ	Largus	2012	-
VAZ	Largus Cross	2015	-
VAZ	Niva Bronto	2021	-
VAZ	Niva Legend	2021	-
VAZ	Niva Travel	2021	-
VAZ	Vesta Cross	2018	-
VAZ	Vesta Sport	2018	-
VAZ	Vesta SW	2017	-
VAZ	XRay	2016	-
VAZ	XRay Cross	2018	-
VAZ	Веста	2015	-
VAZ	Калина	2004	2013
VAZ	Нива 4X4	1977	-
VAZ	Ока	1988	2008
VAZ	Приора	2007	2018
Volkswagen	Amarok	2010	-
Volkswagen	Arteon	2017	-
Volkswagen	Beetle	2011	-
Volkswagen	Bora	1998	2005
Volkswagen	Caddy	2004	-
Volkswagen	CC	2012	-
Volkswagen	Crafter	2008	-
Volkswagen	CrossGolf	2007	2009
Volkswagen	CrossPolo	2006	2009
Volkswagen	CrossTouran	2007	2015
Volkswagen	Eos	2006	2015
Volkswagen	Fox	2005	2011
Volkswagen	Golf	1991	-
Volkswagen	ID.3	2020	-
Volkswagen	ID.4	2020	-
Volkswagen	ID.4 X	2021	-
Volkswagen	ID.5	2022	-
Volkswagen	ID.6 Crozz	2021	-
Volkswagen	ID.6 X	2021	-
Volkswagen	ID.7	2023	-
Volkswagen	ID.Buzz	2022	-
Volkswagen	Jetta	2005	-
Volkswagen	Lupo	1998	2005
Volkswagen	Multivan	2003	-
Volkswagen	New Beetle	1998	2007
Volkswagen	Passat	2000	-
Volkswagen	Passat CC	2008	2012
Volkswagen	Phaeton	2002	2016
Volkswagen	Pointer	1993	1997
Volkswagen	Polo	2001	-
Volkswagen	Routan	2008	2014
Volkswagen	Scirocco	2008	2017
Volkswagen	Sharan	1995	2022
Volkswagen	T-Cross	2018	-
Volkswagen	T-Roc	2017	-
Volkswagen	Taigo	2021	-
Volkswagen	Taos	2020	-
Volkswagen	Tayron	2019	-
Volkswagen	Teramont	2017	-
Volkswagen	Teramont X	2019	-
Volkswagen	Tiguan	2007	-
Volkswagen	Tiguan X	2020	-
Volkswagen	Touareg	2002	-
Volkswagen	Touran	2003	-
Volkswagen	Transporter	1990	-
Volkswagen	Up	2011	-
Volvo	C30	2007	2013
Volvo	C40	2021	-
Volvo	C70	2006	2013
Volvo	C70 Convertible	1997	2006
Volvo	C70 Coupe	1996	2007
Volvo	EC40	2024	-
Volvo	EX30	2023	-
Volvo	EX90	2023	-
Volvo	S40	1995	2012
Volvo	S60	2000	-
Volvo	S70	1996	2000
Volvo	S80	1998	2016
Volvo	S90	1997	-
Volvo	V40	1995	2021
Volvo	V50	2004	2012
Volvo	V60	2011	-
Volvo	V70	1996	2016
Volvo	V90	1996	-
Volvo	XC40	2017	-
Volvo	XC60	2008	-
Volvo	XC70	2001	2016
Volvo	XC90	2002	-
Voyah	Dream	2021	-
Voyah	Free	2021	-
Voyah	Passion	2023	-
Wiesmann	Thunderball	2022	-
Zeekr	001	2021	-
Zeekr	009	2022	-
Zeekr	X	2023	-
\.


--
-- Data for Name: registered_drivers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.registered_drivers (id, first_name, last_name, email, phone_number, driver_license, driver_license_date, car_number, car_model, car_marks, car_color, password, level_access) FROM stdin;
\.


--
-- Data for Name: registered_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.registered_users (id, first_name, last_name, email, phone_number, password, level_access) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

