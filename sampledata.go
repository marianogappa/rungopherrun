package main

import t "github.com/marianogappa/rungopherrun/types"

var graphEdges = [][]t.Edge{
	// []t.Edge{{1, 1}, {3, 4}},
	// []t.Edge{{0, 1}, {2, 2}},
	// []t.Edge{{1, 2}, {3, 3}},
	// []t.Edge{{0, 4}, {2, 3}},
	[]t.Edge{{81, 13}, {82, 23}, {83, 10}, {84, 7}},
	[]t.Edge{{118, 1}, {119, 16}, {120, 7}, {121, 1}},
	[]t.Edge{{94, 12}, {95, 13}, {96, 15}, {97, 4}},
	[]t.Edge{{74, 12}, {75, 21}, {76, 13}, {77, 7}},
	[]t.Edge{{95, 17}, {96, 4}, {97, 9}, {98, 23}},
	[]t.Edge{{147, 13}, {148, 14}, {149, 16}, {150, 16}},
	[]t.Edge{{141, 9}, {142, 13}, {143, 7}, {144, 5}},
	[]t.Edge{{156, 13}, {157, 7}, {158, 11}, {159, 2}},
	[]t.Edge{{13, 16}, {14, 20}, {15, 14}, {16, 9}},
	[]t.Edge{{147, 4}, {148, 25}, {149, 10}, {150, 4}},
	[]t.Edge{{157, 22}, {158, 15}, {159, 25}, {160, 1}},
	[]t.Edge{{105, 14}, {106, 14}, {107, 4}, {108, 6}},
	[]t.Edge{{51, 11}, {52, 6}, {53, 7}, {54, 17}},
	[]t.Edge{{28, 12}, {29, 3}, {30, 9}, {31, 22}},
	[]t.Edge{{163, 2}, {164, 3}, {165, 19}, {166, 23}},
	[]t.Edge{{94, 3}, {95, 14}, {96, 22}, {97, 21}},
	[]t.Edge{{23, 4}, {24, 13}, {25, 9}, {26, 17}},
	[]t.Edge{{59, 9}, {60, 19}, {61, 17}, {62, 3}},
	[]t.Edge{{78, 12}, {79, 22}, {80, 8}, {81, 16}},
	[]t.Edge{{103, 3}, {104, 19}, {105, 6}, {106, 24}},
	[]t.Edge{{25, 2}, {26, 16}, {27, 8}, {28, 13}},
	[]t.Edge{{10, 11}, {11, 11}, {12, 16}, {13, 8}},
	[]t.Edge{{98, 4}, {99, 17}, {100, 8}, {101, 10}},
	[]t.Edge{{97, 18}, {98, 13}, {99, 22}, {100, 20}},
	[]t.Edge{{126, 3}, {127, 7}, {128, 5}, {129, 17}},
	[]t.Edge{{70, 19}, {71, 12}, {72, 20}, {73, 7}},
	[]t.Edge{{52, 1}, {53, 11}, {54, 11}, {55, 13}},
	[]t.Edge{{149, 4}, {150, 19}, {151, 10}, {152, 4}},
	[]t.Edge{{24, 23}, {25, 13}, {26, 8}, {27, 17}},
	[]t.Edge{{39, 16}, {40, 12}, {41, 2}, {42, 2}},
	[]t.Edge{{40, 2}, {41, 20}, {42, 15}, {43, 6}},
	[]t.Edge{{183, 2}, {184, 16}, {185, 3}, {186, 9}},
	[]t.Edge{{167, 7}, {168, 4}, {169, 5}, {170, 23}},
	[]t.Edge{{23, 18}, {24, 9}, {25, 19}, {26, 19}},
	[]t.Edge{{166, 11}, {167, 11}, {168, 16}, {169, 5}},
	[]t.Edge{{162, 8}, {163, 16}, {164, 22}, {165, 15}},
	[]t.Edge{{30, 14}, {31, 1}, {32, 10}, {33, 21}},
	[]t.Edge{{183, 21}, {184, 10}, {185, 23}, {186, 11}},
	[]t.Edge{{165, 13}, {166, 5}, {167, 21}, {168, 24}},
	[]t.Edge{{156, 21}, {157, 17}, {158, 1}, {159, 7}},
	[]t.Edge{{29, 18}, {30, 7}, {31, 3}, {32, 12}},
	[]t.Edge{{120, 25}, {121, 13}, {122, 23}, {123, 18}},
	[]t.Edge{{88, 12}, {89, 4}, {90, 14}, {91, 19}},
	[]t.Edge{{156, 20}, {157, 8}, {158, 8}, {159, 3}},
	[]t.Edge{{75, 7}, {76, 4}, {77, 21}, {78, 18}},
	[]t.Edge{{193, 21}, {194, 22}, {195, 12}, {196, 8}},
	[]t.Edge{{120, 11}, {121, 23}, {122, 5}, {123, 12}},
	[]t.Edge{{60, 21}, {61, 5}, {62, 5}, {63, 15}},
	[]t.Edge{{60, 2}, {61, 7}, {62, 8}, {63, 17}},
	[]t.Edge{{0, 15}, {1, 13}, {2, 9}, {3, 12}},
	[]t.Edge{{4, 11}, {5, 10}, {6, 16}, {7, 20}},
	[]t.Edge{{14, 16}, {15, 13}, {16, 16}, {17, 1}},
	[]t.Edge{{84, 24}, {85, 11}, {86, 19}, {87, 22}},
	[]t.Edge{{190, 25}, {191, 20}, {192, 21}, {193, 12}},
	[]t.Edge{{138, 23}, {139, 20}, {140, 21}, {141, 25}},
	[]t.Edge{{37, 25}, {38, 19}, {39, 7}, {40, 24}},
	[]t.Edge{{152, 21}, {153, 23}, {154, 22}, {155, 18}},
	[]t.Edge{{151, 16}, {152, 8}, {153, 5}, {154, 18}},
	[]t.Edge{{129, 7}, {130, 22}, {131, 6}, {132, 12}},
	[]t.Edge{{143, 4}, {144, 17}, {145, 10}, {146, 18}},
	[]t.Edge{{8, 9}, {9, 4}, {10, 4}, {11, 25}},
	[]t.Edge{{84, 1}, {85, 25}, {86, 7}, {87, 13}},
	[]t.Edge{{86, 18}, {87, 15}, {88, 9}, {89, 10}},
	[]t.Edge{{146, 24}, {147, 25}, {148, 13}, {149, 8}},
	[]t.Edge{{174, 15}, {175, 7}, {176, 22}, {177, 9}},
	[]t.Edge{{128, 17}, {129, 2}, {130, 3}, {131, 13}},
	[]t.Edge{{36, 2}, {37, 6}, {38, 13}, {39, 15}},
	[]t.Edge{{98, 13}, {99, 6}, {100, 10}, {101, 4}},
	[]t.Edge{{159, 12}, {160, 3}, {161, 19}, {162, 21}},
	[]t.Edge{{71, 4}, {72, 16}, {73, 21}, {74, 7}},
	[]t.Edge{{5, 2}, {6, 10}, {7, 17}, {8, 7}},
	[]t.Edge{{95, 6}, {96, 10}, {97, 3}, {98, 18}},
	[]t.Edge{{78, 1}, {79, 19}, {80, 6}, {81, 9}},
	[]t.Edge{{31, 16}, {32, 9}, {33, 10}, {34, 18}},
	[]t.Edge{{23, 13}, {24, 24}, {25, 23}, {26, 4}},
	[]t.Edge{{158, 3}, {159, 22}, {160, 24}, {161, 14}},
	[]t.Edge{{131, 25}, {132, 7}, {133, 16}, {134, 7}},
	[]t.Edge{{132, 3}, {133, 3}, {134, 5}, {135, 9}},
	[]t.Edge{{130, 12}, {131, 13}, {132, 13}, {133, 5}},
	[]t.Edge{{21, 19}, {22, 24}, {23, 7}, {24, 16}},
	[]t.Edge{{122, 21}, {123, 20}, {124, 7}, {125, 25}},
	[]t.Edge{{24, 21}, {25, 20}, {26, 10}, {27, 8}},
	[]t.Edge{{124, 8}, {125, 4}, {126, 25}, {127, 1}},
	[]t.Edge{{140, 15}, {141, 15}, {142, 21}, {143, 15}},
	[]t.Edge{{181, 7}, {182, 2}, {183, 11}, {184, 5}},
	[]t.Edge{{3, 21}, {4, 18}, {5, 23}, {6, 1}},
	[]t.Edge{{177, 13}, {178, 24}, {179, 18}, {180, 10}},
	[]t.Edge{{7, 6}, {8, 6}, {9, 7}, {10, 25}},
	[]t.Edge{{144, 24}, {145, 20}, {146, 23}, {147, 9}},
	[]t.Edge{{174, 11}, {175, 6}, {176, 21}, {177, 15}},
	[]t.Edge{{123, 1}, {124, 10}, {125, 3}, {126, 21}},
	[]t.Edge{{71, 9}, {72, 24}, {73, 10}, {74, 5}},
	[]t.Edge{{154, 22}, {155, 13}, {156, 2}, {157, 11}},
	[]t.Edge{{121, 18}, {122, 2}, {123, 25}, {124, 14}},
	[]t.Edge{{120, 25}, {121, 2}, {122, 16}, {123, 24}},
	[]t.Edge{{196, 17}, {197, 1}, {198, 11}, {199, 23}},
	[]t.Edge{{113, 15}, {114, 10}, {115, 1}, {116, 22}},
	[]t.Edge{{121, 20}, {122, 14}, {123, 12}, {124, 13}},
	[]t.Edge{{88, 24}, {89, 1}, {90, 2}, {91, 20}},
	[]t.Edge{{0, 15}, {1, 6}, {2, 2}, {3, 1}},
	[]t.Edge{{51, 10}, {52, 20}, {53, 17}, {54, 14}},
	[]t.Edge{{45, 17}, {46, 18}, {47, 15}, {48, 20}},
	[]t.Edge{{21, 25}, {22, 2}, {23, 1}, {24, 11}},
	[]t.Edge{{88, 15}, {89, 11}, {90, 17}, {91, 20}},
	[]t.Edge{{182, 17}, {183, 14}, {184, 17}, {185, 23}},
	[]t.Edge{{101, 6}, {102, 3}, {103, 14}, {104, 18}},
	[]t.Edge{{158, 8}, {159, 5}, {160, 1}, {161, 9}},
	[]t.Edge{{140, 3}, {141, 10}, {142, 3}, {143, 4}},
	[]t.Edge{{37, 23}, {38, 20}, {39, 13}, {40, 22}},
	[]t.Edge{{9, 5}, {10, 7}, {11, 15}, {12, 15}},
	[]t.Edge{{23, 18}, {24, 15}, {25, 15}, {26, 6}},
	[]t.Edge{{19, 4}, {20, 11}, {21, 20}, {22, 11}},
	[]t.Edge{{121, 14}, {122, 2}, {123, 4}, {124, 8}},
	[]t.Edge{{184, 13}, {185, 9}, {186, 22}, {187, 1}},
	[]t.Edge{{115, 16}, {116, 18}, {117, 18}, {118, 17}},
	[]t.Edge{{98, 2}, {99, 8}, {100, 13}, {101, 21}},
	[]t.Edge{{116, 3}, {117, 9}, {118, 6}, {119, 20}},
	[]t.Edge{{140, 20}, {141, 24}, {142, 16}, {143, 18}},
	[]t.Edge{{82, 15}, {83, 21}, {84, 8}, {85, 19}},
	[]t.Edge{{76, 2}, {77, 22}, {78, 1}, {79, 19}},
	[]t.Edge{{162, 19}, {163, 2}, {164, 22}, {165, 14}},
	[]t.Edge{{109, 19}, {110, 21}, {111, 13}, {112, 12}},
	[]t.Edge{{41, 16}, {42, 24}, {43, 11}, {44, 2}},
	[]t.Edge{{58, 19}, {59, 18}, {60, 11}, {61, 16}},
	[]t.Edge{{81, 24}, {82, 7}, {83, 18}, {84, 25}},
	[]t.Edge{{99, 24}, {100, 2}, {101, 12}, {102, 11}},
	[]t.Edge{{92, 25}, {93, 15}, {94, 19}, {95, 14}},
	[]t.Edge{{135, 24}, {136, 1}, {137, 16}, {138, 3}},
	[]t.Edge{{190, 17}, {191, 25}, {192, 11}, {193, 20}},
	[]t.Edge{{110, 2}, {111, 17}, {112, 2}, {113, 23}},
	[]t.Edge{{60, 10}, {61, 14}, {62, 13}, {63, 24}},
	[]t.Edge{{17, 19}, {18, 16}, {19, 4}, {20, 19}},
	[]t.Edge{{1, 7}, {2, 14}, {3, 22}, {4, 3}},
	[]t.Edge{{42, 18}, {43, 8}, {44, 10}, {45, 1}},
	[]t.Edge{{107, 23}, {108, 21}, {109, 11}, {110, 24}},
	[]t.Edge{{30, 5}, {31, 13}, {32, 3}, {33, 1}},
	[]t.Edge{{114, 8}, {115, 4}, {116, 2}, {117, 24}},
	[]t.Edge{{1, 21}, {2, 1}, {3, 21}, {4, 3}},
	[]t.Edge{{99, 14}, {100, 25}, {101, 16}, {102, 1}},
	[]t.Edge{{143, 16}, {144, 11}, {145, 15}, {146, 19}},
	[]t.Edge{{151, 13}, {152, 4}, {153, 7}, {154, 7}},
	[]t.Edge{{179, 16}, {180, 11}, {181, 12}, {182, 24}},
	[]t.Edge{{114, 19}, {115, 12}, {116, 6}, {117, 11}},
	[]t.Edge{{36, 7}, {37, 9}, {38, 4}, {39, 2}},
	[]t.Edge{{12, 5}, {13, 14}, {14, 19}, {15, 13}},
	[]t.Edge{{132, 24}, {133, 24}, {134, 1}, {135, 16}},
	[]t.Edge{{77, 2}, {78, 2}, {79, 23}, {80, 9}},
	[]t.Edge{{7, 10}, {8, 15}, {9, 22}, {10, 3}},
	[]t.Edge{{164, 14}, {165, 13}, {166, 1}, {167, 18}},
	[]t.Edge{{138, 21}, {139, 23}, {140, 20}, {141, 4}},
	[]t.Edge{{8, 3}, {9, 14}, {10, 5}, {11, 3}},
	[]t.Edge{{11, 2}, {12, 14}, {13, 23}, {14, 14}},
	[]t.Edge{{180, 17}, {181, 22}, {182, 17}, {183, 15}},
	[]t.Edge{{127, 12}, {128, 14}, {129, 10}, {130, 20}},
	[]t.Edge{{37, 24}, {38, 24}, {39, 20}, {40, 21}},
	[]t.Edge{{162, 11}, {163, 5}, {164, 8}, {165, 4}},
	[]t.Edge{{82, 21}, {83, 7}, {84, 4}, {85, 15}},
	[]t.Edge{{72, 17}, {73, 16}, {74, 3}, {75, 15}},
	[]t.Edge{{120, 15}, {121, 4}, {122, 5}, {123, 1}},
	[]t.Edge{{60, 5}, {61, 13}, {62, 16}, {63, 11}},
	[]t.Edge{{198, 14}, {199, 1}, {0, 4}, {1, 2}},
	[]t.Edge{{180, 8}, {181, 13}, {182, 8}, {183, 9}},
	[]t.Edge{{64, 14}, {65, 25}, {66, 14}, {67, 10}},
	[]t.Edge{{16, 7}, {17, 13}, {18, 23}, {19, 9}},
	[]t.Edge{{44, 9}, {45, 22}, {46, 19}, {47, 2}},
	[]t.Edge{{73, 20}, {74, 6}, {75, 1}, {76, 7}},
	[]t.Edge{{155, 14}, {156, 4}, {157, 17}, {158, 12}},
	[]t.Edge{{140, 13}, {141, 4}, {142, 20}, {143, 12}},
	[]t.Edge{{112, 9}, {113, 12}, {114, 11}, {115, 22}},
	[]t.Edge{{199, 9}, {0, 5}, {1, 2}, {2, 9}},
	[]t.Edge{{16, 7}, {17, 22}, {18, 7}, {19, 15}},
	[]t.Edge{{16, 18}, {17, 25}, {18, 3}, {19, 4}},
	[]t.Edge{{66, 12}, {67, 10}, {68, 24}, {69, 5}},
	[]t.Edge{{111, 9}, {112, 5}, {113, 1}, {114, 13}},
	[]t.Edge{{191, 23}, {192, 11}, {193, 13}, {194, 15}},
	[]t.Edge{{73, 8}, {74, 4}, {75, 7}, {76, 14}},
	[]t.Edge{{17, 15}, {18, 2}, {19, 17}, {20, 11}},
	[]t.Edge{{5, 11}, {6, 6}, {7, 5}, {8, 17}},
	[]t.Edge{{163, 5}, {164, 20}, {165, 23}, {166, 9}},
	[]t.Edge{{36, 11}, {37, 4}, {38, 6}, {39, 23}},
	[]t.Edge{{107, 23}, {108, 5}, {109, 24}, {110, 17}},
	[]t.Edge{{25, 13}, {26, 23}, {27, 19}, {28, 5}},
	[]t.Edge{{4, 15}, {5, 15}, {6, 19}, {7, 18}},
	[]t.Edge{{58, 1}, {59, 6}, {60, 4}, {61, 18}},
	[]t.Edge{{125, 14}, {126, 20}, {127, 11}, {128, 11}},
	[]t.Edge{{152, 22}, {153, 23}, {154, 10}, {155, 2}},
	[]t.Edge{{187, 14}, {188, 22}, {189, 19}, {190, 9}},
	[]t.Edge{{18, 9}, {19, 17}, {20, 19}, {21, 25}},
	[]t.Edge{{121, 24}, {122, 24}, {123, 22}, {124, 5}},
	[]t.Edge{{154, 14}, {155, 11}, {156, 17}, {157, 10}},
	[]t.Edge{{132, 23}, {133, 25}, {134, 20}, {135, 25}},
	[]t.Edge{{125, 9}, {126, 4}, {127, 6}, {128, 8}},
	[]t.Edge{{117, 3}, {118, 23}, {119, 7}, {120, 23}},
	[]t.Edge{{28, 8}, {29, 6}, {30, 19}, {31, 11}},
	[]t.Edge{{193, 24}, {194, 1}, {195, 17}, {196, 22}},
	[]t.Edge{{15, 2}, {16, 7}, {17, 20}, {18, 17}},
	[]t.Edge{{109, 15}, {110, 9}, {111, 18}, {112, 16}},
	[]t.Edge{{37, 22}, {38, 20}, {39, 12}, {40, 14}},
	[]t.Edge{{24, 12}, {25, 9}, {26, 18}, {27, 8}},
	[]t.Edge{{18, 2}, {19, 21}, {20, 21}, {21, 13}},
}

var requirements = []string{"storage room", "starcraft posters", "stainless steel cutlery", "stove", "stores nearby"}

var ads = []t.Ad{
	// {Price: 100, Location: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	// {Price: 200, Location: 2, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	// {Price: 300, Location: 3, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	// {Price: 400, Location: 2, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	// {Price: 110, Location: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	// {Price: 302, Location: 2, Description: " storage room starcraft posters stainless steel cutlery stores nearby"},
	{Price: 381, Location: 74, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 147, Location: 197, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 581, Location: 84, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 125, Location: 198, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 356, Location: 153, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 594, Location: 6, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 262, Location: 104, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 428, Location: 41, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 511, Location: 18, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 537, Location: 194, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 195, Location: 165, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 228, Location: 127, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 547, Location: 77, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 587, Location: 199, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 490, Location: 173, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 641, Location: 38, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 687, Location: 131, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 529, Location: 170, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 637, Location: 25, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 385, Location: 71, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 113, Location: 102, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 494, Location: 44, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 533, Location: 97, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 578, Location: 166, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 659, Location: 132, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 457, Location: 171, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 289, Location: 69, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 500, Location: 8, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 588, Location: 110, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 203, Location: 111, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 351, Location: 81, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 505, Location: 158, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 566, Location: 190, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 261, Location: 172, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 483, Location: 85, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 263, Location: 57, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 102, Location: 93, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 547, Location: 195, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 677, Location: 21, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 296, Location: 124, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 523, Location: 146, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 437, Location: 49, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 141, Location: 188, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 333, Location: 114, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 591, Location: 41, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 378, Location: 31, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 646, Location: 135, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 640, Location: 123, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 652, Location: 82, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 305, Location: 146, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 325, Location: 186, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 415, Location: 154, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 387, Location: 14, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 110, Location: 78, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 490, Location: 114, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 598, Location: 98, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 491, Location: 128, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 484, Location: 69, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 567, Location: 151, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 571, Location: 49, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 426, Location: 15, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 481, Location: 178, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 566, Location: 6, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 393, Location: 72, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 119, Location: 19, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 352, Location: 197, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 185, Location: 183, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 487, Location: 82, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 228, Location: 128, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 484, Location: 181, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 124, Location: 146, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 512, Location: 187, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 516, Location: 191, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 240, Location: 174, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 351, Location: 25, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 140, Location: 54, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 344, Location: 5, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 405, Location: 173, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 301, Location: 10, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 302, Location: 173, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 467, Location: 38, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 278, Location: 173, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 322, Location: 23, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 442, Location: 137, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 643, Location: 64, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 466, Location: 66, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 435, Location: 61, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 204, Location: 164, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 557, Location: 90, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 271, Location: 106, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 530, Location: 17, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 200, Location: 182, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 420, Location: 147, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 370, Location: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 347, Location: 47, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 465, Location: 162, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 129, Location: 40, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 148, Location: 26, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 195, Location: 171, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 300, Location: 179, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 529, Location: 168, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 531, Location: 128, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 186, Location: 37, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 699, Location: 187, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 347, Location: 44, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 188, Location: 42, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 603, Location: 36, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 218, Location: 26, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 119, Location: 189, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 257, Location: 198, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 175, Location: 38, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 153, Location: 158, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 117, Location: 131, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 170, Location: 175, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 686, Location: 123, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 620, Location: 20, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 422, Location: 149, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 561, Location: 88, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 320, Location: 92, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 454, Location: 25, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 160, Location: 183, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 481, Location: 76, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 616, Location: 115, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 339, Location: 91, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 233, Location: 141, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 104, Location: 33, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 609, Location: 144, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 219, Location: 156, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 340, Location: 40, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 440, Location: 30, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 384, Location: 198, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 135, Location: 108, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 321, Location: 122, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 474, Location: 19, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 570, Location: 78, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 438, Location: 193, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 644, Location: 5, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 274, Location: 65, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 424, Location: 139, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 106, Location: 163, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 652, Location: 5, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 322, Location: 16, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 417, Location: 128, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 165, Location: 24, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 179, Location: 82, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 629, Location: 193, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 621, Location: 117, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 211, Location: 152, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 153, Location: 188, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 559, Location: 73, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 108, Location: 160, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 428, Location: 74, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 674, Location: 167, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 275, Location: 180, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 206, Location: 41, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 186, Location: 60, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 214, Location: 38, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 559, Location: 99, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 223, Location: 137, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 687, Location: 114, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 674, Location: 88, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 431, Location: 29, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 683, Location: 33, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 616, Location: 46, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 377, Location: 156, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 136, Location: 56, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 205, Location: 171, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 689, Location: 88, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 587, Location: 30, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 509, Location: 70, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 659, Location: 56, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 302, Location: 187, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 445, Location: 4, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 353, Location: 195, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 220, Location: 146, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 505, Location: 148, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 409, Location: 65, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 256, Location: 82, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 680, Location: 198, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 177, Location: 37, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 178, Location: 156, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 493, Location: 34, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 508, Location: 191, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 165, Location: 124, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 584, Location: 42, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 523, Location: 187, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 373, Location: 92, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 353, Location: 32, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 627, Location: 66, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 348, Location: 168, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 231, Location: 8, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 606, Location: 116, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 231, Location: 161, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 577, Location: 27, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 554, Location: 186, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 230, Location: 175, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 687, Location: 143, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 104, Location: 106, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 493, Location: 31, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 656, Location: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
}

var sortedAds = []t.Ad{
	{Price: 100, Location: 1, Distance: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 110, Location: 1, Distance: 1, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 200, Location: 2, Distance: 3, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 300, Location: 3, Distance: 4, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
	{Price: 400, Location: 2, Distance: 3, Description: " storage room starcraft posters stainless steel cutlery stove stores nearby"},
}
