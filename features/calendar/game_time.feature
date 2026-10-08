Feature: Game time
  Each Game-Turn covers about a week and is divided into three Operations
  Stages, which are the basic unit of time. A time is written as the
  Operations Stage, a slash, then the Game-Turn: 1/31 is Stage 1 of
  Game-Turn 31.

  Some rules still name Game-Turns by month and week, as in
  "September III, 1940". The errata says this dating was abandoned, with
  the Roman numeral giving the week of the month. Every month has four
  Game-Turns, weeks I to IV.

  Rule: Times are written Operations Stage, then Game-Turn

    @case-4.45 @case-5.1
    Scenario Outline: Reading a time
      When the time <time> is read
      Then it is Operations Stage <stage> of Game-Turn <turn>

      Examples:
        | time  | stage | turn |
        | 1/1   | 1     | 1    |
        | 3/2   | 3     | 2    |
        | 1/31  | 1     | 31   |
        | 3/111 | 3     | 111  |

    @case-5.1 @case-64.2
    Scenario Outline: Times that are not in the campaign
      When the time <time> is read
      Then it is not a time in the campaign

      Examples:
        | time  |
        | 0/5   |
        | 4/5   |
        | 1/0   |
        | 1/112 |
        | 31    |
        | I/31  |

  Rule: Operations Stages follow in order

    @case-5.1
    Scenario Outline: The next Operations Stage
      Then the Operations Stage after <time> is <next>

      Examples:
        | time | next |
        | 1/31 | 2/31 |
        | 2/31 | 3/31 |
        | 3/30 | 1/31 |

    @case-5.1
    Scenario Outline: Comparing times
      Then <earlier> comes before <later>
      And <later> comes after <earlier>

      Examples:
        | earlier | later |
        | 3/30    | 1/31  |
        | 1/31    | 2/31  |
        | 1/1     | 3/111 |

  Rule: The campaign runs from 1/1 to 3/111

    @case-64.2
    Scenario: The first and last Operations Stages of the campaign
      Then the campaign starts at 1/1
      And the campaign ends at 3/111
      And there is no Operations Stage after 3/111

  Rule: Each Game-Turn is a week of the calendar

    @case-64.2 @errata
    Scenario Outline: Dates the rules give for Game-Turns
      Then Game-Turn <turn> is <date>

      @case-6.14
      Examples: The start of the campaign
        | turn | date                |
        | 1    | September III, 1940 |

      Examples: The arrival of Rommel
        | turn | date            |
        | 26   | March IV, 1941  |

      @case-43.13
      Examples: German bombers in Crete
        | turn | date         |
        | 35   | June I, 1941 |

      @case-56.21
      Examples: Axis convoy planning
        | turn | date               |
        | 54   | October IV, 1941   |
        | 55   | November I, 1941   |

      @case-4.48 @case-4.49
      Examples: The first Operations Stage of January 1942
        | turn | date            |
        | 63   | January I, 1942 |

    @case-29.1 @errata
    Scenario Outline: Each season starts on a Game-Turn twelve after the last
      Then <date> is Game-Turn <turn>

      Examples:
        | date                | turn |
        | September III, 1940 | 1    |
        | December II, 1940   | 12   |
        | December III, 1940  | 13   |
        | March II, 1941      | 24   |
        | March III, 1941     | 25   |
        | June II, 1941       | 36   |
        | June III, 1941      | 37   |
        | September II, 1941  | 48   |
