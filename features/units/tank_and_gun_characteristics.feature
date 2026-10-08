Feature: Tank and Gun Characteristics
  Every TOE Strength Point of a tank, gun or anti-aircraft weapon takes
  its ratings from its nationality's Tank and Gun Characteristics Chart.
  The engine holds a copy of each chart, corrected by the errata, and a
  unit's abilities come from the TOE Strength Points in it.

  A dash on the charts means the rating doesn't apply or is zero.

  Rule: The engine's charts match the printed charts, as corrected by the errata

    # The A9 Cruiser's Armor Protection comes from the Commonwealth
    # printing; the Axis printing omits it.
    @case-4.47 @errata
    Scenario: The Commonwealth Tank and Gun Characteristics Chart
      Then the Commonwealth tank weapon systems read:
        | Type            | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | A9 Cruiser      | 25  | 1  | -       | 3          | -   | 1            | 3/3           | 2         | 1R  |
        | A10 Cruiser     | 15  | 1  | -       | 3          | -   | 2            | 3/3           | 2         | 1R  |
        | A13 Cruiser     | 30  | 0  | -       | 3          | -   | 2            | 2/3           | 2         | 1R  |
        | Churchill II    | 15  | 1  | -       | 6          | -   | 7            | 5/6           | 7         | 0   |
        | Crusader Mk.I   | 25  | 1  | -       | 3          | -   | 3            | 3/4           | 3         | 1R  |
        | Crusader Mk.II  | 25  | 1  | -       | 3          | -   | 4            | 4/4           | 3         | 1R  |
        | Crusader Mk.III | 25  | 1  | -       | 6          | -   | 4            | 4/5           | 3         | 1R  |
        | Grant M3        | 20  | 1  | -       | 7          | -   | 4            | 6/5           | 7         | 0   |
        | Mark VI Light   | 35  | 1  | -       | 0          | -   | 1            | 2/2           | 1         | 0   |
        | Matilda Mk.II   | 15  | 0  | -       | 3          | -   | 6            | 3/4           | 3         | 1R  |
        | Sherman         | 20  | 1  | -       | 6          | -   | 6            | 5/5           | 6         | 0   |
        | Stuart M3       | 35  | 1  | -       | 4          | -   | 3            | 4/4           | 4         | 0   |
        | Valentine Mk.II | 15  | 0  | -       | 3          | -   | 5            | 3/4           | 3         | 0   |
        | Scorpion        | 25  | 0  | -       | 0          | -   | 7            | 0/(2)         | 3         | 1L  |
      And the Commonwealth artillery weapon systems read:
        | Type              | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 18-pounder Gun    | 15  | -  | 7       | (2)        | 4   | -            | 0/1           | 1         | -   |
        | 18/25-pounder Gun | 15  | -  | 8       | 2          | 5   | -            | 1/1           | 1         | -   |
        | 25-pounder Gun    | 15  | -  | 8       | 5          | 6   | -            | 1/1           | 1         | -   |
        | 4.5" Gun          | 15  | -  | 11      | 1          | 9   | -            | 1/0           | 1         | -   |
        | 5.5" Gun/Howitzer | 15  | -  | 15      | 1          | 7   | -            | 1/0           | 1         | -   |
        | 60-pounder Gun    | 15  | -  | 12      | 1          | 7   | -            | 1/0           | 1         | -   |
        | SP 25-pounder Gun | 15  | -  | 8       | 4          | 3   | 5            | 2/1           | 3         | 2R  |
        | 3.7" Howitzer     | 15  | -  | 7       | 0          | 3   | -            | 0/0           | 1         | -   |
        | 4.5" Howitzer     | 15  | -  | 9       | 0          | 3   | -            | 1/0           | 1         | -   |
        | 6" Howitzer       | 15  | -  | 15      | 1          | 5   | -            | 1/0           | 1         | -   |
        | 155mm Howitzer    | 15  | -  | 15      | 1          | 6   | -            | 1/0           | 1         | -   |
        | 105mm SP Howitzer | 20  | 1  | 9       | 2          | 4   | 4            | 2/3           | 3         | 0   |
      And the Commonwealth anti-tank weapon systems read:
        | Type         | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 2-pounder    | 15  | -  | -       | 4          | 2   | -            | 1/1           | 1         | -   |
        | 6-pounder    | 15  | -  | -       | 7          | 2   | -            | 1/1           | 1         | -   |
        | 17-pounder   | 15  | -  | -       | 13         | 2   | -            | 1/1           | 1         | -   |
        | SP 6-pounder | 20  | -  | -       | 7          | 2   | 1            | 2/2           | 1         | 1L  |
      And the Commonwealth anti-air weapon systems read:
        | Type                   | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | Light AA (Bofors 40mm) | 15  | 1  | -       | (1)        | 2   | -            | (1)/(1)       | 1         | -   |
        | Heavy AA (3.7")        | 15  | 4  | -       | (7)        | 2   | -            | (1)/(1)       | 1         | -   |

    @case-4.48
    Scenario: The Italian Tank and Gun Characteristics Chart
      Then the Italian tank weapon systems read:
        | Type         | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | CV 33(L3/35) | 25  | 0  | -       | 0          | -   | 1            | 1/2           | 1         | 2R  |
        | L 6/40       | 25  | 0  | -       | 1          | -   | 2            | 2/2           | 2         | 0   |
        | M 11/39      | 20  | 1  | -       | 2          | -   | 2            | 3/3           | 2         | 1R  |
        | M 13/40      | 20  | 1  | -       | 3          | -   | 3            | 3/3           | 2         | 1R  |
        | M14/41       | 20  | 1  | -       | 3          | -   | 3            | 3/3           | 2         | 0   |
      And the Italian artillery weapon systems read:
        | Type                | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 65/17 Gun           | 15  | -  | 5       | 0          | 3   | -            | 1/1           | 1         | -   |
        | 75/18 Gun-Howitzer  | 15  | -  | 6       | 0          | 5   | -            | 1/0           | 1         | -   |
        | 75/18 Gun           | 20  | -  | 6       | 6          | 4   | 3            | 3/3           | 2         | 1R  |
        | 75/27 Gun           | 15  | -  | 6       | 2          | 4   | -            | 1/1           | 1         | -   |
        | 100/17 Howitzer     | 15  | -  | 8       | 0          | 5   | -            | 1/0           | 1         | -   |
        | 105/28 Gun          | 15  | -  | 9       | 1          | 7   | -            | 1/1           | 1         | -   |
        | 149/13 Howitzer     | 15  | -  | 15      | 0          | 5   | -            | 1/0           | 1         | -   |
        | ParaArt             | 15  | -  | 2       | 2          | 1   | -            | 1/1           | 1         | -   |
        | 149mm Vichy French  | 15  | -  | 10      | 1          | 5   | -            | 0/1           | 1         | -   |
        | 155mm Rimhailo(Fr.) | 15  | -  | 13      | 0          | 3   | -            | 0/0           | 1         | -   |
      And the Italian anti-tank weapon systems read:
        | Type          | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 47/32 Mod. 37 | 15  | -  | -       | 4          | 2   | -            | 1/1           | 1         | -   |
      And the Italian anti-air weapon systems read:
        | Type                    | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | Light (20mm M/35 Breda) | 15  | 1  | -       | 2          | 2   | -            | (1)/(1)       | 1         | -   |
        | 'I' Light               | 0+  | 1  | -       | 2          | 2   | -            | (1)/(1)       | 1         | -   |
        | Heavy (75/46 Mod. 34)   | 15  | 2  | -       | (5)        | 2   | -            | (1)/(1)       | 1         | -   |
        | 'I' Heavy               | 0+  | 2  | -       | (5)        | 2   | -            | (1)/(1)       | 0         | -   |
        | Heavy (90/53)           | 15  | 3  | -       | 9          | 2   | -            | (1)/(1)       | 1         | -   |

    # The German chart prints its guns under the tank heading with no
    # heading of their own; they are artillery here.
    # The errata replaces the Pz III E row and restores the CPAs of the
    # 7.62cm Pak(R) and the Marder III.
    @case-4.49 @errata
    Scenario: The German Tank and Gun Characteristics Chart
      Then the German tank weapon systems read:
        | Type               | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | PZ I               | 24  | 1  | -       | 0          | -   | 1            | 2/2           | 2         | 0   |
        | Pz II              | 35  | 0  | -       | 1          | -   | 2            | 2/3           | 2         | 0   |
        | Pz III E           | 25  | 1  | -       | 4          | -   | 3            | 4/4           | 3         | 0   |
        | Pz III H           | 25  | 1  | -       | 4          | -   | 4            | 4/4           | 3         | 0   |
        | Pz III J(Special)  | 26  | 1  | -       | 6          | -   | 4            | 7/5           | 3         | 0   |
        | Pz IV D            | 25  | 1  | -       | 6          | -   | 3            | 5/4           | 4         | 0   |
        | Pz IV E            | 25  | 1  | -       | 6          | -   | 3            | 5/5           | 4         | 0   |
        | Pz IV F2 (Special) | 25  | 1  | -       | 8          | -   | 4            | 6/5           | 4         | 0   |
      And the German artillery weapon systems read:
        | Type                                | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 7.5 cm(IG18) Light Infantry Gun     | 15  | -  | 6       | 3          | 2   | -            | 1/1           | 1         | -   |
        | 10.5cm(K18) Medium Gun              | 15  | -  | 9       | 1          | 9   | -            | 1/1           | 1         | -   |
        | 10.5cm(leFH18) Light Field Howitzer | 15  | -  | 9       | 2          | 5   | -            | 1/1           | 1         | -   |
        | 10.5cm SP K18 Gun                   | 15  | -  | 9       | 2          | 4   | 2            | 2/2           | 2         | 0   |
        | 149mm Vichy French                  | 15  | -  | 10      | 1          | 5   | -            | 0/1           | 1         | -   |
        | 15cm (sIG33) Medium Infantry Gun    | 15  | -  | 14      | 1          | 2   | -            | 1/1           | 1         | -   |
        | 15cm(sFH18) Medium Field Howitzer   | 15  | -  | 15      | 1          | 7   | -            | 0/1           | 1         | -   |
        | 15cm(K18) Gun                       | 15  | -  | 15      | 1          | 11  | -            | 1/1           | 1         | -   |
        | 15cm SP Gun                         | 25  | -  | 14      | 1          | 2   | 2            | 2/2           | 2         | 0   |
        | 155mm (french 1915 short)           | 15  | -  | 13      | 0          | 5   | -            | 1/1           | 1         | -   |
        | 17cm (K18) Gun                      | 15  | -  | 15      | 0          | 13  | -            | 1/0           | 1         | -   |
        | 21cm (mrs18) Howitzer               | 15  | -  | 18      | 0          | 8   | -            | 1/0           | 1         | -   |
      And the German anti-tank weapon systems read:
        | Type                         | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 2.8cm s.Pz.B.41 or 28/20 Pak | 15  | -  | -       | 4          | 2   | -            | 1/1           | 1         | -   |
        | 3.7cm Pak 35/36              | 15  | -  | -       | 2          | 1   | -            | 1/1           | 1         | -   |
        | 5cm Pak 38                   | 15  | -  | -       | 5          | 2   | -            | 1½/1          | 1         | -   |
        | 7.62cm Pak(R)                | 15  | -  | -       | 9          | 2   | -            | 1/1           | 1         | -   |
        | Pzjg 1 (SP)                  | 30  | -  | -       | 4          | 2   | 1            | 2/2           | 1         | 0   |
        | Marder III (SP)              | 25  | -  | -       | 9          | 2   | 2            | 3/2           | 2         | 0   |
      And the German anti-air weapon systems read:
        | Type                     | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | Light (20mm & 37mm Flak) | 15  | 1  | -       | (1)        | (1) | -            | (1)/(1)       | 1         | -   |
        | Heavy (88mm Flak)        | 15  | 3  | -       | 13         | 2   | -            | 2/1           | 1         | -   |
  Rule: The errata corrects the printed charts

    @case-4.49 @errata
    Scenario: The German Pz III E row is replaced
      Then the German "Pz III E" reads:
        | CPA | AA | Barrage | Anti-Armor | Vul | Armor Prtctn | Close Assault | Fuel Rate | BAR |
        | 25  | 1  | -       | 4          | -   | 3            | 4/4           | 3         | 0   |

    @case-4.49 @errata
    Scenario Outline: Missing German CPAs are restored
      Then the German "<system>" has a CPA of <cpa>

      Examples:
        | system          | cpa |
        | 7.62cm Pak(R)   | 15  |
        | Marder III (SP) | 25  |

    @case-4.47 @errata
    Scenario: The A9 Cruiser has an Armor Protection Rating
      Then the Commonwealth "A9 Cruiser" has an Armor Protection Rating of 1

  Rule: Only tanks and self-propelled guns break down

    @case-21.11 @case-21.12 @case-3.4
    Scenario Outline: Tanks and self-propelled guns break down with their chart rating
      Then the <nationality> "<system>" breaks down with a rating of <rating>

      Examples:
        | nationality  | system            | rating |
        | Commonwealth | Sherman           | 0      |
        | Commonwealth | Crusader Mk.I     | 1R     |
        | Commonwealth | Scorpion          | 1L     |
        | Commonwealth | SP 25-pounder Gun | 2R     |
        | Commonwealth | SP 6-pounder      | 1L     |
        | Italian      | CV 33(L3/35)      | 2R     |
        | Italian      | 75/18 Gun         | 1R     |
        | German       | Pz III H          | 0      |
        | German       | Marder III (SP)   | 0      |

    @case-21.11 @case-3.4 @ruling
    Scenario Outline: Towed guns never break down
      Then the <nationality> "<system>" never breaks down

      Examples:
        | nationality  | system                 |
        | Commonwealth | 25-pounder Gun         |
        | Commonwealth | 2-pounder              |
        | Commonwealth | Heavy AA (3.7")        |
        | Italian      | 75/18 Gun-Howitzer     |
        | Italian      | 'I' Heavy              |
        | German       | 15cm(K18) Gun          |
        | German       | Heavy (88mm Flak)      |

    @case-21.11 @case-3.4 @ruling
    Scenario: Every chart's ratings agree with which vehicles break down
      Then every tank and self-propelled gun has a Breakdown Adjustment Rating
      And no other weapon system has one

  Rule: Some ratings change during the campaign

    # These need the game calendar, which isn't built yet.

    @case-4.49 @wip
    Scenario: German tanks have a rating of 1R early in the campaign
      # The chart's note names the "1/31" Game-Turn; see ruling R-004.
      Given the Game-Turn is before the one named by the note to the German chart
      Then the German "Pz III H" breaks down with a rating of 1R

    @case-4.48 @case-4.49 @wip
    Scenario Outline: Some guns have no anti-armor rating until January 1942
      Given the Game-Turn is <turn>
      Then the <nationality> "<system>" has an Anti-Armor Rating of <rating>

      Examples:
        | turn | nationality | system                          | rating |
        | 62   | Italian     | 75/27 Gun                       | 0      |
        | 63   | Italian     | 75/27 Gun                       | 2      |
        | 62   | German      | 7.5 cm(IG18) Light Infantry Gun | 0      |
        | 63   | German      | 7.5 cm(IG18) Light Infantry Gun | 3      |
