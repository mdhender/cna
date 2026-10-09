Feature: Weather Table
  The weather is rolled for each Operations Stage. Two dice, read
  sequentially, are rolled on the season's row of the Weather Table. A
  Sandstorm or Rainstorm strikes only some map sections: one die on the
  Foul Weather Location Table picks them, and the others have Normal
  weather. Hot weather covers the whole map.

  The printed Weather Table puts each block of Game-Turns on the row of
  the opposite season, which the errata acknowledges. Following ruling
  R-005, the weather results stay on the rows they are printed on, and
  each season's Game-Turns come from its dates in Case 29.1.

  Rule: The engine's Weather Table matches the corrected chart

    @case-29.61 @errata @ruling
    Scenario: The Weather Table
      Then the Weather Table reads:
        | Season | Game-Turns                  | Normal  | Hot     | Sandstorm | Rainstorm |
        | Fall   | 1...12, 49...60, 97...108   | 11...35 | 36...54 | 55...61   | 62...66   |
        | Winter | 13...24, 61...72, 109...111 | 11...52 | -       | -         | 53...66   |
        | Spring | 25...36, 73...84            | 11...42 | 43...55 | 56...64   | 65...66   |
        | Summer | 37...48, 85...96            | 11...23 | 24...55 | 56...66   | -         |

    @engine
    Scenario: Every throw gives exactly one weather in every season
      Then every season of the Weather Table gives exactly one weather for each sequential reading

  Rule: The Game-Turn picks the season

    @case-29.1 @ruling
    Scenario Outline: The season of a Game-Turn
      Given the time is <time>
      Then the season is <season>

      Examples:
        | time  | season |
        | 1/1   | Fall   |
        | 3/12  | Fall   |
        | 1/13  | Winter |
        | 3/24  | Winter |
        | 1/25  | Spring |
        | 3/36  | Spring |
        | 1/37  | Summer |
        | 3/48  | Summer |
        | 1/49  | Fall   |
        | 3/108 | Fall   |
        | 1/109 | Winter |
        | 3/111 | Winter |

  Rule: The dice give the weather for the season

    @case-29.1
    Scenario: A roll of 53 in Summer is Hot
      Given the time is 2/40
      And the tens die will roll 5 and the ones die will roll 3
      When the weather is determined
      Then the season is Summer
      And the weather is Hot

    @case-29.1 @case-29.61
    Scenario Outline: Reading the Weather Table
      Given the season is <season>
      Then a roll of <roll> gives <weather> weather

      Examples: Fall
        | season | roll | weather   |
        | Fall   | 11   | Normal    |
        | Fall   | 35   | Normal    |
        | Fall   | 36   | Hot       |
        | Fall   | 54   | Hot       |
        | Fall   | 55   | Sandstorm |
        | Fall   | 61   | Sandstorm |
        | Fall   | 62   | Rainstorm |
        | Fall   | 66   | Rainstorm |

      Examples: Winter
        | season | roll | weather   |
        | Winter | 11   | Normal    |
        | Winter | 52   | Normal    |
        | Winter | 53   | Rainstorm |
        | Winter | 66   | Rainstorm |

      Examples: Spring
        | season | roll | weather   |
        | Spring | 11   | Normal    |
        | Spring | 42   | Normal    |
        | Spring | 43   | Hot       |
        | Spring | 55   | Hot       |
        | Spring | 56   | Sandstorm |
        | Spring | 64   | Sandstorm |
        | Spring | 65   | Rainstorm |
        | Spring | 66   | Rainstorm |

      Examples: Summer
        | season | roll | weather   |
        | Summer | 11   | Normal    |
        | Summer | 23   | Normal    |
        | Summer | 24   | Hot       |
        | Summer | 55   | Hot       |
        | Summer | 56   | Sandstorm |
        | Summer | 66   | Sandstorm |

  Rule: Foul weather strikes only the map sections the Foul Weather Location Table gives

    @case-29.7
    Scenario: The Foul Weather Location Table
      Then the Foul Weather Location Table reads:
        | Die          | 1   | 2   | 3   | 4   | 5   | 6     |
        | Map Sections | A,B | C,D | D,E | B,C | B,D | B,C,D |

    @case-29.1 @case-29.7
    Scenario Outline: A Sandstorm in Fall
      Given the time is 1/5
      And the tens die will roll 5 and the ones die will roll 5
      And the die will roll <die>
      When the weather is determined
      Then the weather is Sandstorm on map sections <struck>
      And the weather is Normal on map sections <spared>

      Examples:
        | die | struck  | spared  |
        | 1   | A,B     | C,D,E   |
        | 2   | C,D     | A,B,E   |
        | 3   | D,E     | A,B,C   |
        | 4   | B,C     | A,D,E   |
        | 5   | B,D     | A,C,E   |
        | 6   | B,C,D   | A,E     |

    @case-29.1 @case-29.7
    Scenario: A Rainstorm in Winter
      Given the time is 3/20
      And the tens die will roll 6 and the ones die will roll 1
      And the die will roll 2
      When the weather is determined
      Then the weather is Rainstorm on map sections C,D
      And the weather is Normal on map sections A,B,E

    @case-29.31
    Scenario: Hot weather covers every map section
      Given the time is 1/40
      And the tens die will roll 3 and the ones die will roll 1
      When the weather is determined
      Then the weather is Hot on map sections A,B,C,D,E
      And no more dice are thrown

    @case-29.1
    Scenario: Normal weather needs no second roll
      Given the time is 1/15
      And the tens die will roll 2 and the ones die will roll 2
      When the weather is determined
      Then the weather is Normal on map sections A,B,C,D,E
      And no more dice are thrown

  Rule: Sandstorms never strike Delta hexes

    @case-29.41 @case-29.7
    Scenario: A Sandstorm on map section E spares the Delta
      Given the time is 1/5
      And the tens die will roll 5 and the ones die will roll 5
      And the die will roll 3
      When the weather is determined
      Then the weather is Sandstorm on map sections D,E
      And the weather is Sandstorm in hex E3116
      And the weather is Normal in hex E3117
      And the weather is Normal in hex E3215
      And the weather is Sandstorm in hex D3116

    @case-29.5 @case-29.58
    Scenario: A Rainstorm on map section E strikes the Delta too
      Given the time is 3/20
      And the tens die will roll 6 and the ones die will roll 1
      And the die will roll 3
      When the weather is determined
      Then the weather is Rainstorm in hex E3116
      And the weather is Rainstorm in hex E3117
