Feature: Breakdown Adjustment Ratings
  Each type of vehicle has a Breakdown Adjustment Rating, printed on the
  Tank and Gun Characteristics Charts. A rating shifts the Breakdown
  Table column left (L) or right (R), or not at all (0). A dash means the
  vehicles never break down, which is not the same as a rating of 0.

  Rule: A printed rating shifts the column

    @case-21.13
    Scenario Outline: Reading a printed rating
      When the printed rating <printed> is read
      Then the vehicles break down with the column shifted <shift>

      Examples:
        | printed | shift           |
        | 0       | 0 columns       |
        | 1L      | 1 column left   |
        | 2L      | 2 columns left  |
        | 1R      | 1 column right  |
        | 2R      | 2 columns right |

  Rule: A dash means the vehicles never break down

    @case-3.4 @case-21.11 @ruling
    Scenario: Reading a dash
      When the printed rating - is read
      Then the vehicles never break down

    @case-21.13 @ruling
    Scenario: A rating of 0 is not a dash
      When the printed rating 0 is read
      Then the vehicles break down with the column shifted 0 columns
