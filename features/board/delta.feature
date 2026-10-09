Feature: Delta hexes
  The Nile Delta is a terrain type on the Terrain Effects Chart. Its hexes
  are all on map section E. Sandstorms never reach them, and they limit
  vehicles in a Rainstorm.

  Hexes are written as on the map: the section, then two digits of row and
  two of column. Rows are numbered from south to north, so the list below
  runs from the northern coast down the Nile.

  Rule: The engine's Delta hexes match the map

    # First transcribed from Michael Miller's Hex Database (2015), which
    # is known to have errors. Each row should be desk-checked against the
    # official map section E. A correction goes here and in the engine's
    # data (internal/board/delta.go); this scenario fails until both agree.
    @case-8.37 @case-29.41
    Scenario: The Delta hexes of map section E
      Then the Delta hexes of map section E are:
        | Row | Columns       |
        | 42  | 28            |
        | 41  | 29            |
        | 39  | 20, 27-34     |
        | 38  | 19-33         |
        | 37  | 15-16, 18-34  |
        | 36  | 14-33         |
        | 35  | 14-34         |
        | 34  | 15-33         |
        | 33  | 15-34         |
        | 32  | 14-33         |
        | 31  | 17-34         |
        | 30  | 17-33         |
        | 29  | 22-34         |
        | 28  | 23-33         |
        | 27  | 25-34         |
        | 26  | 24-33         |
        | 25  | 25-34         |
        | 24  | 24-33         |
        | 23  | 25-33         |
        | 22  | 24-31         |
        | 21  | 27-31         |
        | 20  | 27-30         |
        | 19  | 28-29         |
        | 18  | 28            |
        | 17  | 29            |
        | 16  | 29            |
        | 15  | 30            |
        | 13  | 30            |
        | 12  | 29-30         |
        | 11  | 30            |
        | 10  | 29-30         |
        | 9   | 26, 30        |
        | 8   | 19, 23-27, 29 |
        | 7   | 20-27, 29     |
        | 6   | 23-26, 28     |
        | 5   | 24-26, 28-29  |
        | 4   | 23-24, 26-28  |
        | 3   | 22-24, 27-28  |
        | 2   | 26-27         |
        | 1   | 26-27         |

    # E1430 is a city on the Nile with no Delta marker, which breaks the
    # strip of Delta hexes along the river between E1330 and E1530.
    @case-8.37
    Scenario Outline: Hexes checked against the official map
      Then hex <hex> <is> a Delta hex

      Examples:
        | hex   | is     |
        | E4228 | is     |
        | E4129 | is     |
        | E3718 | is     |
        | E3414 | is not |
        | E3215 | is     |
        | E3116 | is not |
        | E3117 | is     |
        | E3120 | is     |
        | E2724 | is not |
        | E1530 | is     |
        | E1429 | is not |
        | E1430 | is not |
        | E1431 | is not |
        | E1330 | is     |
        | E0820 | is not |
        | E0819 | is     |
        | E0728 | is not |
        | E0627 | is not |
        | E0527 | is not |
        | E0425 | is not |
        | E0322 | is     |

  Rule: A hex that is partly water takes the terrain of its land

    # E3717 and E3718 look alike on the map: mostly water, with a Delta
    # marker reaching into the hex. The data has E3718 as Delta but E3717
    # as Clear. See ruling R-006.
    @case-8.37 @ruling @wip
    Scenario: A partly water hex with a Delta marker is a Delta hex
      Then hex E3717 is a Delta hex
      And hex E3718 is a Delta hex
