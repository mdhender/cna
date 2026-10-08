Feature: Rolling dice
  The game uses ordinary six-sided dice, read in three ways.
  One die gives 1 to 6. Two dice added together give 2 to 12.
  Two dice read sequentially give 11 to 66: the large die is the tens
  digit and the small die is the units digit, so the two dice must be
  told apart by size, not by which shows the higher number.

  @case-16.33
  Scenario: Rolling one die
    Given the die will roll 5
    When the player rolls one die
    Then the result is 5

  @case-3.1 @case-12.42 @case-15.73
  Scenario Outline: Reading two dice sequentially
    Given the large die will roll <large> and the small die will roll <small>
    When the player throws two dice
    Then the sequential reading is <reading>

    Examples:
      | large | small | reading |
      | 2     | 5     | 25      |
      | 5     | 2     | 52      |
      | 3     | 4     | 34      |
      | 6     | 3     | 63      |
      | 1     | 1     | 11      |
      | 6     | 6     | 66      |

  @case-15.73
  Scenario Outline: Adding two dice
    Given the large die will roll <large> and the small die will roll <small>
    When the player throws two dice
    Then the sum is <sum>

    Examples:
      | large | small | sum |
      | 1     | 1     | 2   |
      | 3     | 4     | 7   |
      | 4     | 3     | 7   |
      | 6     | 6     | 12  |

  @case-15.73
  Scenario: Reading one throw both ways
    Given the large die will roll 3 and the small die will roll 4
    When the player throws two dice
    Then the sequential reading is 34
    And the sum is 7

  @engine
  Scenario: One die rolls every face and nothing else
    Given a dice roller seeded with 1 and 2
    When the player rolls one die 1000 times
    Then every face from 1 to 6 is rolled
    And no other result is rolled

  @engine
  Scenario: A sequential read can give all 36 combinations
    Given a dice roller seeded with 3 and 4
    When the player throws two dice 5000 times
    Then every sequential reading with both digits from 1 to 6 occurs
    And no other sequential reading occurs

  @engine
  Scenario: The same seed rolls the same dice
    Given a dice roller seeded with 7 and 11
    And a second dice roller seeded with 7 and 11
    When each roller throws two dice 100 times
    Then both rollers throw the same dice in the same order

  @engine
  Scenario: Different seeds roll different dice
    Given a dice roller seeded with 7 and 11
    And a second dice roller seeded with 7 and 12
    When each roller throws two dice 100 times
    Then the rollers throw different dice
