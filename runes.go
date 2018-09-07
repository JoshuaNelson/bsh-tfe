package main
// https://www.charbase.com/block/miscellaneous-technical

var LineStar        rune = 0x2217
var CircleBorder    rune = 0x2218
var CircleFill      rune = 0x2219
var ThreeF          rune = 0x222D // Waterfall
var ThreeDot        rune = 0x2234 // Predator
var ThreeDotReverse rune = 0x2235
var TwoDot          rune = 0x2236
var FourDot         rune = 0x2237
var LineDot         rune = 0x2238
var FourDotLine     rune = 0x223A
var TwoLineTilde    rune = 0x2245
var ThreeTilde      rune = 0x224B
var BiPlane         rune = 0x2256
var TwoLineCurve    rune = 0x2258 // Infantry
var TwoLineArrow    rune = 0x2259 // Infantry
var TwoLineVee      rune = 0x225A // Infantry
var TwoLineStar     rune = 0x225B // Infantry
var TwoLineCone     rune = 0x225C // Infantry
var ThreeLine       rune = 0x2261
var LeftRight       rune = 0x2276
var RightLeft       rune = 0x2277
var PokeBall        rune = 0x224E
var Ellipses        rune = 0x22EF
var Diameter        rune = 0x2300
var Electric        rune = 0x2301 // Power Station
var House           rune = 0x2302 // Barracks/Housing
var UpArrow         rune = 0x2303 // Mountain
var DownArrow       rune = 0x2304 // Brush
var Projective      rune = 0x2305 // Mountain x1
var Perspective     rune = 0x2306 // Mountain x2 / RADAR?
var SquishBox       rune = 0x2311
var Arc             rune = 0x2312 // Hill
var Segment         rune = 0x2313 // Rocks?
var Sector          rune = 0x2314 // Motion Sensor / RADAR?
var Telephone       rune = 0x2315 // Pin
var Knot            rune = 0x2318

var BottomRightCrop rune = 0x230C
var BottomLeftCrop  rune = 0x230D
var TopRightCrop    rune = 0x230F
var TopLeftCrop     rune = 0x2310

var TopLeftCorner   rune = 0x231C
var TopRightCorner  rune = 0x231D
var BotLeftCorner   rune = 0x231E
var BotRightCorner  rune = 0x231F

var Option          rune = 0x2325 // Runway?
var XRightArrow     rune = 0x2326 // Sign?
var XRectangle      rune = 0x2327 // Barricade
var WideArrowLeft   rune = 0x2329
var WideArrowRight  rune = 0x232A
var XLeftArrow      rune = 0x232B
var Countersink     rune = 0x2335 // Clean arrow down
var SkinnyRectangle rune = 0x2337
var QuadBar         rune = 0x2338 // Car
var Diamond         rune = 0x233A // Explosives / Barrels
var Jot             rune = 0x233B
var Circle          rune = 0x233C // Tank
var CircleStile     rune = 0x233D // Sensor Fence?
var CircleJot       rune = 0x233E // Resource Collectors

var QuadSlash       rune = 0x2341 // Box right slash
var QuadBackslash   rune = 0x2342
var QuadLT          rune = 0x2343
var QuadGT          rune = 0x2344
var CircleBackslash rune = 0x2349 // Resource Not Collecting
var Delta           rune = 0x234B // Alpine
var QuadDelta       rune = 0x234D // Pizza Tank
var DelStile        rune = 0x2352 // Upside down alpine
var QuadUpCar       rune = 0x2353 // IFV
var DeltaUnderbar   rune = 0x2359 // LRM launchers
var DiamondUnderbar rune = 0x235A // FLoating box
var JotUnderbar     rune = 0x235B // Drone
var CircleStar      rune = 0x235F // Command?
var QuadQuote       rune = 0x235E // Tank Standard

var QuadJotQuote    rune = 0x2360 // Tank Standard Type2
var StarDiaeresis   rune = 0x2363 // Infantry?
var JotDiaeresis    rune = 0x2364
var CircleDiaeresis rune = 0x2365
var DelTilde        rune = 0x236B // Something
var StileTilde      rune = 0x236D // Fence / Powerline?

var Helm            rune = 0x2388 // Ship
var CircleBarNotch  rune = 0x2389 // Something
var CircleTriangle  rune = 0x238A // Something
var CircleArrow     rune = 0x238B

var Quad            rune = 0x2395
var TopParenRight   rune = 0x239B
var TopParenLeft    rune = 0x239E

// Curly Bracket
var RightUpperHook  rune = 0x23A7
var RightBrackMid   rune = 0x23A8
var RightLowerHook  rune = 0x23A9
var LineUp          rune = 0x23AB
var LeftUpperHook   rune = 0x23AB
var LeftBrackMid    rune = 0x23AC
var LeftLowerHook   rune = 0x23AD
var Eject           rune = 0x23CF // Armored House
var Bucket          rune = 0x2423

// 0x2500 All the borders, bold, dotted, double

var TwoLowBarPerp   rune = 0x2567
var LowBarTwoPerp   rune = 0x2568
var VerticalLowBold rune = 0x257d // Communication Tower

// 0x2581 Loading Bars Low to High

var Gravel          rune = 0x2591 // Gravel
var HeavyGravel     rune = 0x2592
var ThreeQuadrantsL rune = 0x2599 // Factory
var ThreeQuadrantsR rune = 0x259F
var SquareOutline   rune = 0x25A1
var RoundedSquare   rune = 0x25A2
var NestedSquare    rune = 0x25A3
var GrateHorizontal rune = 0x25A4
var GrateVertical   rune = 0x25A5
var GrateDiagLeft   rune = 0x25A7
var GrateDiagRight  rune = 0x25A8

// 0x25AA Small shapes

var Radioactive     rune = 0x2622
var Biohazard       rune = 0x2623
var RadioTower      rune = 0x2645 // Radio Tower
var Recycling       rune = 0x267B
var BlackFlag       rune = 0x2691
var HammerPick      rune = 0x2692 // Manufacturing // Oil
var NavalAnchor     rune = 0x2693
var CrossSwords     rune = 0x2694
var Aesculapius     rune = 0x2695 // Staff with snake
var GearTeeth       rune = 0x2699
var Atom            rune = 0x269B
var WhiteCircle     rune = 0x26AA
var BlackCircle     rune = 0x26AB
var Plane           rune = 0x2708 // Airliner facing right
var TripHashVert    rune = 0x29FB // Heat Sink

// 0x2B12 Half shaded shapes

//var RadarSplotch    rune = 0x2E0E // Splotch on radar?
var DoubleDiagSmall rune = 0x2E17
var OverpassThree   rune = 0x2FAE

// 0x2FF0 Dotted Quadrant Shapes

var HorizontalWave  rune = 0x3030
var ThreeSided      rune = 0x3107


var Generation      rune = 0x346A
var ToRaid          rune = 0x3474
var Guild           rune = 0x3479 // GuildHall
var Harvest         rune = 0x3486
var Slave           rune = 0x3492
var Watchful        rune = 0x349C
var Immortal        rune = 0x34A5
var Bandits         rune = 0x34C2

var CloudyObscure   rune = 0x53C6
var ArrowSign       rune = 0x5DE0

var Machine         rune = 0x6A5F

var CrocusIron      rune = 0x1F71E // Communications Tower
var SublimateCopper rune = 0x1F722 // Laser drill
var Crucible        rune = 0x1F765 // windy flag
var BushTwo         rune = 0x10176
var DashVee         rune = 0x10197
var RightFlag       rune = 0x10280
var Triangle        rune = 0x10285
var LeftTarget      rune = 0x102b9
var CandyCane       rune = 0x10413
var SmallCircleX    rune = 0x1044F
var PartW           rune = 0x1047F
var TentJot         rune = 0x104B2
var SmallCircleDot  rune = 0x104EB

var SlightDblDiag   rune = 0x1091A
