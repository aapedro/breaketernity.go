package breaketernity

import (
	"math"
)

const MAX_SIGNIFICANT_DIGITS int8 = 17

const EXP_LIMIT float64 = 9e15

const LAYER_DOWN float64 = 15.9542425094393248746 //math.Log10(9e15)

const FIRST_NEG_LAYER float64 = 1 / 1e18

const NUMBER_EXP_MAX int = 308

const NUMBER_EXP_MIN int = -324 // idk

const MAX_ES_IN_A_ROW float64 = 5

const DEFAULT_FROM_STRING_CACHE_SIZE = (1 << 10) - 1

const IGNORE_COMMAS = true

const COMMAS_ARE_DECIMAL_POINTS = false

var dOne *Decimal = dFC_NN(1, 0, 1)

var dNegOne *Decimal = dFC_NN(0, 0, 1)

var dInf *Decimal = dFC(1, math.Inf(1), math.Inf(1))

var dNegInf *Decimal = dFC(-1, math.Inf(1), math.Inf(1))

var dZero *Decimal = dFC_NN(0, 0, 0)

var CRITICAL_HEADERS = []float64{2, math.E, 3, 4, 5, 6, 7, 8, 9, 10}

var CRITICAL_TETR_VALUES = [][]float64{
	{
		1, 1.0891180521811202527, 1.1789767925673958433, 1.2701455431742086633, 1.3632090180450091941,
		1.4587818160364217007, 1.5575237916251418333, 1.6601571006859253673, 1.7674858188369780435, 1.8804192098842727359,
		2,
	},
	{
		1, 1.1121114330934078681, 1.2310389249316089299, 1.3583836963111376089, 1.4960519303993531879,
		1.6463542337511945810, 1.8121385357018724464, 1.9969713246183068478, 2.2053895545527544330, 2.4432574483385252544,
		math.E,
	},
	{
		1, 1.1187738849693603, 1.2464963939368214, 1.38527004705667, 1.5376664685821402,
		1.7068895236551784, 1.897001227148399, 2.1132403089001035, 2.362480153784171,
		2.6539010333870774, 3,
	},
	{
		1, 1.1367350847096405, 1.2889510672956703, 1.4606478703324786, 1.6570295196661111,
		1.8850062585672889, 2.1539465047453485, 2.476829779693097, 2.872061932789197,
		3.3664204535587183, 4,
	},
	{
		1, 1.1494592900767588, 1.319708228183931, 1.5166291280087583, 1.748171114438024,
		2.0253263297298045, 2.3636668498288547, 2.7858359149579424, 3.3257226212448145,
		4.035730287722532, 5,
	},
	{
		1, 1.159225940787673, 1.343712473580932, 1.5611293155111927, 1.8221199554561318,
		2.14183924486326, 2.542468319282638, 3.0574682501653316, 3.7390572020926873, 4.6719550537360774,
		6,
	},
	{
		1, 1.1670905356972596, 1.3632807444991446, 1.5979222279405536, 1.8842640123816674,
		2.2416069644878687, 2.69893426559423, 3.3012632110403577, 4.121250340630164, 5.281493033448316,
		7,
	},
	{
		1, 1.1736630594087796, 1.379783782386201, 1.6292821855668218, 1.9378971836180754,
		2.3289975651071977, 2.8384347394720835, 3.5232708454565906, 4.478242031114584,
		5.868592169644505, 8,
	},
	{
		1, 1.1793017514670474, 1.394054150657457, 1.65664127441059, 1.985170999970283,
		2.4069682290577457, 2.9647310119960752, 3.7278665320924946, 4.814462547283592,
		6.436522247411611, 9,
	},
	{
		1, 1.1840100246247336579, 1.4061375836156954169, 1.6802272208863963918, 2.026757028388618927,
		2.4770056063449647580, 3.0805252717554819987, 3.9191964192627283911, 5.1351528408331864230,
		6.9899611795347148455, 10,
	},
}

var CRITICAL_SLOG_VALUES = [][]float64{
	{
		-1, -0.9194161097107025, -0.8335625019330468, -0.7425599821143978, -0.6466611521029437,
		-0.5462617907227869, -0.4419033816638769, -0.3342645487554494, -0.224140440909962,
		-0.11241087890006762, 0,
	},
	{
		-1, -0.90603157029014, -0.80786507256596, -0.7064666939634, -0.60294836853664,
		-0.49849837513117, -0.39430303318768, -0.29147201034755, -0.19097820800866,
		-0.09361896280296, 0,
	},
	{
		-1, -0.9021579584316141, -0.8005762598234203, -0.6964780623319391, -0.5911906810998454,
		-0.486050182576545, -0.3823089430815083, -0.28106046722897615, -0.1831906535795894,
		-0.08935809204418144, 0,
	},
	{
		-1, -0.8917227442365535, -0.781258746326964, -0.6705130326902455, -0.5612813129406509,
		-0.4551067709033134, -0.35319256652135966, -0.2563741554088552, -0.1651412821106526,
		-0.0796919581982668, 0,
	},
	{
		-1, -0.8843387974366064, -0.7678744063886243, -0.6529563724510552, -0.5415870994657841,
		-0.4352842206588936, -0.33504449124791424, -0.24138853420685147, -0.15445285440944467,
		-0.07409659641336663, 0,
	},
	{
		-1, -0.8786709358426346, -0.7577735191184886, -0.6399546189952064, -0.527284921869926,
		-0.4211627631006314, -0.3223479611761232, -0.23107655627789858, -0.1472057700818259,
		-0.07035171210706326, 0,
	},
	{
		-1, -0.8740862815291583, -0.7497032990976209, -0.6297119746181752, -0.5161838335958787,
		-0.41036238255751956, -0.31277212146489963, -0.2233976621705518, -0.1418697367979619,
		-0.06762117662323441, 0,
	},
	{
		-1, -0.8702632331800649, -0.7430366914122081, -0.6213373075161548, -0.5072025698095242,
		-0.40171437727184167, -0.30517930701410456, -0.21736343968190863, -0.137710238299109,
		-0.06550774483471955, 0,
	},
	{
		-1, -0.8670016295947213, -0.7373984232432306, -0.6143173985094293, -0.49973884395492807,
		-0.394584953527678, -0.2989649949848695, -0.21245647317021688, -0.13434688362382652,
		-0.0638072667348083, 0,
	},
	{
		-1, -0.8641642839543857, -0.732534623168535, -0.6083127477059322, -0.4934049257184696,
		-0.3885773075899922, -0.29376029055315767, -0.2083678561173622, -0.13155653399373268,
		-0.062401588652553186, 0,
	},
}
