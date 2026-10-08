Feature: Breakdown Table
  Vehicles break down as they move. A unit collects Breakdown Points for
  the terrain it moves through during an Operations Stage. When it stops,
  its total picks a column of the Breakdown Table. The column shifts for
  the vehicles' Breakdown Adjustment Rating and the weather, then two dice
  read sequentially give the percentage of the unit's TOE Strength Points
  that break down.

  Unless a scenario says otherwise, the unit's Breakdown Adjustment Rating
  is 0 and the weather is Normal.

  Rule: The engine's table matches the printed chart

    @case-21.38
    Scenario: The Breakdown Table
      Then the Breakdown Table reads:
        | % Breakdown | 0...3   | 4...10  | 11...20 | 21...30 | 31...40 | 41...50 | 51...60 | 61...70 | 71+     |
        | 0           | 11...66 | 11...42 | 11...32 | 11...26 | 11...23 | 11...16 | 11...14 | -       | -       |
        | 10          | -       | 43...64 | 33...62 | 31...55 | 24...53 | 21...46 | 15...42 | 11...33 | 11...25 |
        | 25          | -       | 65      | 63...64 | 56...62 | 54...61 | 51...56 | 43...54 | 34...52 | 26...43 |
        | 33          | -       | 66      | 65      | 63...65 | 62...64 | 61...63 | 55...63 | 53...62 | 44...55 |
        | 50          | -       | -       | 66      | 66      | 65...66 | 64...65 | 64...65 | 63...64 | 56...63 |
        | 75          | -       | -       | -       | -       | -       | 66      | 66      | 65...66 | 64...66 |

    @engine
    Scenario: Every throw gives exactly one result in every column
      Then every column of the Breakdown Table gives exactly one result for each sequential reading

  Rule: The points accumulated pick the column

    @case-21.31
    Scenario Outline: Picking the column
      Given the unit has accumulated <points> Breakdown Points
      When the unit stops moving
      Then it checks for Breakdown on the <column> column

      Examples:
        | points | column  |
        | 4      | 4...10  |
        | 10     | 4...10  |
        | 11     | 11...20 |
        | 20     | 11...20 |
        | 21     | 21...30 |
        | 30     | 21...30 |
        | 31     | 31...40 |
        | 40     | 31...40 |
        | 41     | 41...50 |
        | 50     | 41...50 |
        | 51     | 51...60 |
        | 60     | 51...60 |
        | 61     | 61...70 |
        | 70     | 61...70 |
        | 71     | 71+     |
        | 150    | 71+     |

    @case-21.31
    Scenario Outline: A fraction of a point rounds up
      Given the unit has accumulated <points> Breakdown Points
      When the unit stops moving
      Then it checks for Breakdown on the <column> column

      Examples:
        | points | column  |
        | 3.5    | 4...10  |
        | 10.5   | 11...20 |
        | 20.5   | 21...30 |

    @case-21.27
    Scenario Outline: A unit with three points or fewer does not check
      Given the unit has accumulated <points> Breakdown Points
      And the unit's Breakdown Adjustment Rating is 2R
      And the weather is Hot
      When the unit stops moving
      Then it does not check for Breakdown

      Examples:
        | points |
        | 0      |
        | 0.5    |
        | 3      |

  Rule: The vehicles and the weather shift the column

    @case-21.13 @case-21.32 @case-21.37
    Scenario Outline: Shifting the column
      Given the unit has accumulated <points> Breakdown Points
      And the unit's Breakdown Adjustment Rating is <rating>
      And the weather is <weather>
      When the unit stops moving
      Then it checks for Breakdown on the <column> column

      Examples:
        | points | rating | weather | column  |
        | 35     | 0      | Normal  | 31...40 |
        | 35     | 2L     | Normal  | 11...20 |
        | 35     | 1L     | Normal  | 21...30 |
        | 35     | 1R     | Normal  | 41...50 |
        | 35     | 2R     | Normal  | 51...60 |
        | 35     | 0      | Hot     | 41...50 |
        | 35     | 2L     | Hot     | 21...30 |
        | 35     | 1R     | Hot     | 51...60 |
        | 35     | 2R     | Hot     | 61...70 |

    @case-21.37
    Scenario Outline: A Sandstorm shifts the column one to the right
      Given the unit has accumulated 35 Breakdown Points
      And the weather is <weather>
      And a Sandstorm shifts the column
      When the unit stops moving
      Then it checks for Breakdown on the <column> column

      Examples:
        | weather | column  |
        | Normal  | 41...50 |
        | Hot     | 51...60 |

    @case-21.33
    Scenario Outline: The column cannot shift past either end of the table
      Given the unit has accumulated <points> Breakdown Points
      And the unit's Breakdown Adjustment Rating is <rating>
      And the weather is <weather>
      When the unit stops moving
      Then it checks for Breakdown on the <column> column

      Examples:
        | points | rating | weather | column |
        | 71     | 1R     | Normal  | 71+    |
        | 65     | 2R     | Hot     | 71+    |
        | 8      | 2L     | Normal  | 0...3  |
        | 15     | 2L     | Normal  | 0...3  |

    @case-21.33
    Scenario: A unit shifted left of the 4...10 column suffers no Breakdown
      Given the unit has accumulated 15 Breakdown Points
      And the unit's Breakdown Adjustment Rating is 2L
      And the tens die will roll 6 and the ones die will roll 6
      When the unit stops moving
      Then the result is 0% Breakdown

  Rule: The dice, read sequentially, give the percentage

    @case-21.34
    Scenario Outline: Reading the dice
      Given the unit has accumulated <points> Breakdown Points
      And the tens die will roll <tens> and the ones die will roll <ones>
      When the unit stops moving
      Then the result is <percent>% Breakdown

      Examples:
        | points | tens | ones | percent |
        | 35     | 2    | 3    | 0       |
        | 35     | 2    | 4    | 10      |
        | 35     | 5    | 3    | 10      |
        | 35     | 5    | 4    | 25      |
        | 35     | 6    | 1    | 25      |
        | 35     | 6    | 2    | 33      |
        | 35     | 6    | 4    | 33      |
        | 35     | 6    | 5    | 50      |
        | 35     | 6    | 6    | 50      |
        | 80     | 1    | 1    | 10      |
        | 80     | 6    | 3    | 50      |
        | 80     | 6    | 4    | 75      |

    @case-21.0 @case-21.31 @case-21.34 @ruling
    Scenario: The Breakdown dice are read sequentially, not added
      Given the unit has accumulated 80 Breakdown Points
      And the tens die will roll 2 and the ones die will roll 6
      When the unit stops moving
      Then the dice read 26
      And the result is 25% Breakdown

  Rule: The percentage applies to TOE Strength Points, rounding up

    @case-21.34 @case-21.35
    Scenario Outline: Rounding up the points broken down
      Given the unit has <toe> TOE Strength Points
      When the unit suffers <percent>% Breakdown
      Then <broken> TOE Strength Points break down

      Examples:
        | toe | percent | broken |
        | 30  | 10      | 3      |
        | 20  | 33      | 7      |
        | 10  | 25      | 3      |
        | 3   | 33      | 1      |
        | 4   | 75      | 3      |
        | 7   | 0       | 0      |
        | 2   | 10      | 1      |

    @case-21.35
    Scenario Outline: A unit of one TOE Strength Point ignores a 10% result
      Given the unit has 1 TOE Strength Point
      When the unit suffers <percent>% Breakdown
      Then <broken> TOE Strength Points break down

      Examples:
        | percent | broken |
        | 10      | 0      |
        | 25      | 1      |

  Rule: A full check, from points to vehicles broken down

    @case-21.32 @case-21.34 @case-21.35 @case-21.37
    Scenario: Trucks in Hot weather
      Given the unit has 30 TOE Strength Points
      And the unit has accumulated 35 Breakdown Points
      And the unit's Breakdown Adjustment Rating is 2L
      And the weather is Hot
      And the tens die will roll 3 and the ones die will roll 3
      When the unit stops moving
      Then it checks for Breakdown on the 21...30 column
      And the result is 10% Breakdown
      And 3 TOE Strength Points break down

    @case-21.32 @case-21.34 @case-21.35 @case-21.37
    Scenario: Tanks in Hot weather
      Given the unit has 20 TOE Strength Points
      And the unit has accumulated 35 Breakdown Points
      And the unit's Breakdown Adjustment Rating is 1R
      And the weather is Hot
      And the tens die will roll 6 and the ones die will roll 1
      When the unit stops moving
      Then it checks for Breakdown on the 51...60 column
      And the result is 33% Breakdown
      And 7 TOE Strength Points break down
