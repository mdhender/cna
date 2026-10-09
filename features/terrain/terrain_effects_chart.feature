Feature: Terrain Effects Chart
  Each type of terrain has a cost in Capability Points (CP) to move into
  it, one for non-motorized and one for motorized units. Hexside features,
  such as wadis and slopes, add to the cost of the hex beyond them. Vehicles
  also pick up Breakdown Points for the terrain they move through. The chart
  also gives each terrain's shifts for combat and its stacking limit, which
  later features will use.

  The copy of the chart below follows the printed layout. A footnote marker
  is written in parentheses after the cell's value, as in "4 (3)". Text that
  spans several columns is in the first of them, and the cells it covers
  are left empty.

  Unless a scenario says otherwise, the weather is not a Rainstorm.

  The errata corrects the chart: note 4 belongs with Major City, not Swamp,
  and the "1" printed in both of Track's CP columns is wrong. Note 8 says
  what a track costs.

  Rule: The engine's chart matches the printed chart, as corrected by the errata

    # Heavy Vegetation's Breakdown Value is printed "3." on the scan.
    # The dot is a speck beside the 3 (mdhender/cnadocs#13).
    @case-8.37 @errata
    Scenario: The Terrain Effects Chart
      Then the Terrain Effects Chart reads:
        | Terrain Type            | CP non-Mot                              | CP Mot  | Breakdown Value | Barrage            | Anti Armor | Close Assault | Stacking Limit (1) |
        | Clear                   | 2                                       | 2       | 4               | -                  | -          | -             | 6                  |
        | Gravel                  | 2                                       | 2       | 6               | -                  | -          | -             | 6                  |
        | Salt Marsh (2)          | 3                                       | 2       | 6               | -                  | -          | R1            | 6                  |
        | Heavy Vegetation        | 3                                       | 3       | 3               | -                  | L1         | L1            | 6                  |
        | Rough                   | 3                                       | 4       | 8               | L1                 | L1         | L2            | 6                  |
        | Mountain                | 4                                       | 6       | 12              | L2                 | L2         | L3            | 3                  |
        | Delta                   | 2                                       | 4       | 2               | -                  | -          | -             | 6                  |
        | Desert                  | 3                                       | 4 (3)   | 24              | -                  | -          | -             | 6                  |
        | Major City (4)          | 1                                       | ½       | ½               | See Fortifications |            |               | 8                  |
        | Swamp                   | May enter only on road or railroad      |         |                 | -                  | -          | -             | 6                  |
        | Village/Bir/Oasis       | Same as terrain in hex for all purposes |         |                 |                    |            |               |                    |
        | Railroad (5)            | Same as terrain in hex for all purposes |         |                 |                    |            |               |                    |
        | Road                    | 1 (6)                                   | ½ (6)   | ½ (6)           | -                  | -          | -             | 5 (7)              |
        | Track                   | (8)                                     | (8)     | - (8)           | -                  | -          | -             | 5 (7)              |
        | Ridge                   | +2                                      | +4      | 2               | -                  | L2         | L2            | -                  |
        | Up Slope                | +2                                      | +4      | 2               | -                  | L1         | L2            | -                  |
        | Down Slope              | +1                                      | +2      | 2               | -                  | L1         | R1            | -                  |
        | Up Escarpment           | +6                                      | P       | -               | -                  | P          | L3            | -                  |
        | Down Escarpment         | +4                                      | +8 (9)  | +6              | -                  | L2         | R1            | -                  |
        | Wadi                    | +1 (10)                                 | +4 (10) | +8              | -                  | -          | L1            | -                  |
        | Major River             | +8                                      | P (11)  | -               | -                  | -          | L6            | -                  |
        | Minor River             | +3                                      | +6      | +1              | -                  | -          | L2            | -                  |
        | Fortifications          |                                         |         |                 |                    |            |               |                    |
        | Level One               | Same as terrain in hex                  |         |                 | L1 (12)            | L1         | L2            | -                  |
        | Level Two               | Same as terrain in hex                  |         |                 | L2 (12)            | L2         | L3            | -                  |
        | Level Three             | Same as terrain in hex                  |         |                 | L2 (12)            | L2         | L4            | -                  |
        | Friendly Minefield (13) | +1                                      | +4      | 0               | -                  | L1         | L1            | -                  |
        | Enemy Minefield (13)    | +4                                      | +CPA    | +2              | -                  | -          | -             | -                  |

    @case-8.37 @errata
    Scenario: Note 4 belongs to Major City, not Swamp
      Then note 4 is on the "Major City" row only

    @case-8.37 @errata
    Scenario: A track has no fixed CP cost
      Then the "Track" row gives no CP cost of its own

    @engine
    Scenario: Every type of terrain has exactly one row
      Then every type of terrain has exactly one row on the Terrain Effects Chart

  Rule: Moving into a hex costs the CP and Breakdown Points of its terrain

    @case-8.31 @case-8.37 @case-21.21
    Scenario Outline: Moving into a hex
      Given the unit is <unit>
      When it moves into a <terrain> hex
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | terrain          | cp | bp |
        | non-motorized | Clear            | 2  | 0  |
        | motorized     | Clear            | 2  | 4  |
        | non-motorized | Gravel           | 2  | 0  |
        | motorized     | Gravel           | 2  | 6  |
        | non-motorized | Salt Marsh       | 3  | 0  |
        | non-motorized | Heavy Vegetation | 3  | 0  |
        | motorized     | Heavy Vegetation | 3  | 3  |
        | non-motorized | Rough            | 3  | 0  |
        | motorized     | Rough            | 4  | 8  |
        | non-motorized | Mountain         | 4  | 0  |
        | motorized     | Mountain         | 6  | 12 |
        | non-motorized | Delta            | 2  | 0  |
        | motorized     | Delta            | 4  | 2  |
        | non-motorized | Desert           | 3  | 0  |
        | motorized     | Desert           | 4  | 24 |
        | non-motorized | Major City       | 1  | 0  |
        | motorized     | Major City       | ½  | ½  |

  Rule: Hexside features add to the cost of the hex beyond them

    @case-8.31 @case-8.37 @case-8.43 @case-21.21
    Scenario Outline: Crossing a hexside
      Given the unit is <unit>
      When it moves into a <terrain> hex across <hexsides>
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | terrain | hexsides                      | cp | bp |
        | non-motorized | Clear   | a Ridge                       | 4  | 0  |
        | motorized     | Clear   | a Ridge                       | 6  | 6  |
        | motorized     | Clear   | an Up Slope                   | 6  | 6  |
        | non-motorized | Clear   | a Down Slope                  | 3  | 0  |
        | motorized     | Clear   | a Down Slope                  | 4  | 6  |
        | non-motorized | Clear   | an Up Escarpment              | 8  | 0  |
        | non-motorized | Clear   | a Down Escarpment             | 6  | 0  |
        | non-motorized | Clear   | a Wadi                        | 3  | 0  |
        | motorized     | Clear   | a Wadi                        | 6  | 12 |
        | non-motorized | Clear   | a Major River                 | 10 | 0  |
        | non-motorized | Clear   | a Minor River                 | 5  | 0  |
        | motorized     | Clear   | a Minor River                 | 8  | 5  |
        | motorized     | Rough   | an Up Slope and a Minor River | 14 | 11 |

  Rule: A move adds up the cost of every hex entered

    # The example with Case 21.23: an Italian tank battalion of M13s moves
    # four hexes east. It follows the road for two hexes, crossing a ridge
    # on the road, which costs nothing extra. The road then bends away, so
    # the battalion leaves it for a Rough hex, and crosses a second ridge
    # into a Clear hex. The road hexes' own terrain doesn't matter.
    @case-8.33 @case-8.37 @case-21.23
    Scenario: The tank battalion in the example of Case 21.23
      Given the unit is motorized
      When it moves through these hexes:
        | route          | terrain | hexsides |
        | along a road   | Clear   | a Ridge  |
        | along a road   | Clear   |          |
        | across country | Rough   |          |
        | across country | Clear   | a Ridge  |
      Then it spends 11 Capability Points
      And it picks up 15 Breakdown Points

  Rule: Some terrain is closed to some units

    @case-8.37 @case-8.42
    Scenario Outline: Vehicles cross escarpments and major rivers only on a road or track
      Given the unit is motorized
      When it moves into a Clear hex across <hexside>
      Then it may not move there

      Examples:
        | hexside           |
        | an Up Escarpment  |
        | a Down Escarpment |
        | a Major River     |

    @case-8.37
    Scenario Outline: A Swamp may be entered only on a road
      Given the unit is <unit>
      When it moves into a Swamp hex
      Then it may not move there

      Examples:
        | unit          |
        | non-motorized |
        | motorized     |

    @case-8.37 @case-8.44
    Scenario Outline: Only light vehicles leave the track in a Salt Marsh
      Given the unit is <unit>
      When it moves into a Salt Marsh hex
      Then <result>

      Examples:
        | unit                | result                                                        |
        | motorized           | it may not move there                                         |
        | Light Trucks        | it spends 2 Capability Points and picks up 6 Breakdown Points |
        | recce               | it spends 2 Capability Points and picks up 6 Breakdown Points |
        | motorcycle infantry | it spends 2 Capability Points and picks up 6 Breakdown Points |

    @case-8.37 @case-8.45
    Scenario Outline: Light Trucks and motorcycle units never enter Desert
      Given the unit is <unit>
      When it moves <route> into a Desert hex
      Then <result>

      Examples:
        | unit                | route          | result                                                         |
        | Light Trucks        | across country | it may not move there                                          |
        | Light Trucks        | along a road   | it may not move there                                          |
        | Light Trucks        | along a track  | it may not move there                                          |
        | motorcycle infantry | across country | it may not move there                                          |
        | motorcycle recce    | across country | it may not move there                                          |
        | recce               | across country | it spends 4 Capability Points and picks up 24 Breakdown Points |

  Rule: A road replaces the terrain it runs through

    @case-8.33 @case-8.37
    Scenario Outline: Moving along a road
      Given the unit is <unit>
      When it moves along a road into a <terrain> hex
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | terrain  | cp | bp |
        | non-motorized | Rough    | 1  | 0  |
        | motorized     | Rough    | ½  | ½  |
        | motorized     | Mountain | ½  | ½  |
        | non-motorized | Swamp    | 1  | 0  |
        | motorized     | Swamp    | ½  | ½  |

    @case-8.37
    Scenario Outline: A road cancels the cost of the hexside it crosses
      Given the unit is <unit>
      When it moves along a road into a <terrain> hex across <hexsides>
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | terrain | hexsides                      | cp | bp |
        | non-motorized | Clear   | a Wadi                        | 1  | 0  |
        | motorized     | Clear   | a Wadi                        | ½  | ½  |
        | motorized     | Clear   | an Up Slope and a Minor River | ½  | ½  |
        | non-motorized | Clear   | a Major River                 | 1  | 0  |
        | motorized     | Clear   | a Major River                 | ½  | ½  |

    # Ruling R-008 is Open. The chart lets a road cancel every hexside cost,
    # but Case 8.42 keeps vehicles from going up an escarpment, and from
    # going down one except on a track. The proposal lets vehicles use a
    # road across an escarpment either way.
    @case-8.37 @case-8.42 @ruling @wip
    Scenario Outline: Vehicles use a road across an escarpment
      Given the unit is motorized
      When it moves along a road into a Clear hex across <hexside>
      Then it spends ½ Capability Points
      And it picks up ½ Breakdown Points

      Examples:
        | hexside           |
        | an Up Escarpment  |
        | a Down Escarpment |

    # Ruling R-008 is Open. Case 8.44 lets any vehicle into a Salt Marsh on
    # a road or track; the chart's note 2 names only the track. The
    # proposal follows Case 8.44.
    @case-8.37 @case-8.44 @ruling @wip
    Scenario: Vehicles use a road into a Salt Marsh
      Given the unit is motorized
      When it moves along a road into a Salt Marsh hex
      Then it spends ½ Capability Points
      And it picks up ½ Breakdown Points

  Rule: A track halves the cost of the terrain it runs through

    # Ruling R-007. The chart prints a cost of 1 CP for a track, which the
    # errata strikes in favor of note 8: a track halves the cost of the
    # terrain and hexsides, except for a vehicle going down an escarpment.
    # Case 8.46 still says 1 CP per hex, and Case 8.33 says a unit on a
    # track ignores the other terrain. We follow the errata and note 8.
    @case-8.33 @case-8.37 @case-8.46 @errata @ruling
    Scenario Outline: Moving along a track
      Given the unit is <unit>
      When it moves along a track into a <terrain> hex
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | terrain          | cp | bp |
        | non-motorized | Clear            | 1  | 0  |
        | motorized     | Clear            | 1  | 2  |
        | non-motorized | Heavy Vegetation | 1½ | 0  |
        | motorized     | Rough            | 2  | 4  |
        | motorized     | Desert           | 2  | 12 |
        | motorized     | Salt Marsh       | 1  | 3  |
        | motorized     | Major City       | ¼  | ¼  |

    @case-8.37 @case-8.46 @errata @ruling
    Scenario Outline: A track halves the cost of the hexside it crosses
      Given the unit is <unit>
      When it moves along a track into a Clear hex across <hexsides>
      Then it spends <cp> Capability Points
      And it picks up <bp> Breakdown Points

      Examples:
        | unit          | hexsides         | cp | bp |
        | non-motorized | a Wadi           | 1½ | 0  |
        | motorized     | a Wadi           | 3  | 6  |
        | motorized     | a Minor River    | 4  | 2½ |
        | non-motorized | an Up Escarpment | 4  | 0  |
        | non-motorized | a Down Escarpment | 3 | 0  |
        | non-motorized | a Major River    | 5  | 0  |

    # Note 8's exception is for vehicles; a non-motorized unit going down
    # an escarpment on a track pays half, as for any other hexside.
    @case-8.37 @case-8.42 @ruling
    Scenario: Vehicles pay in full to go down an escarpment on a track
      Given the unit is motorized
      When it moves along a track into a Clear hex across a Down Escarpment
      Then it spends 9 Capability Points
      And it picks up 8 Breakdown Points

    @case-8.37 @case-8.42
    Scenario: Vehicles never go up an escarpment, even on a track
      Given the unit is motorized
      When it moves along a track into a Clear hex across an Up Escarpment
      Then it may not move there

    @case-8.37
    Scenario Outline: A Swamp may not be entered on a track
      Given the unit is <unit>
      When it moves along a track into a Swamp hex
      Then it may not move there

      Examples:
        | unit          |
        | non-motorized |
        | motorized     |

    @case-8.37
    Scenario: Vehicles cross a Major River only on a road, not a track
      Given the unit is motorized
      When it moves along a track into a Clear hex across a Major River
      Then it may not move there

  Rule: The weather changes the cost of some terrain

    @case-8.37 @case-29.0 @wip
    Scenario: A Wadi can't be crossed off a road during a Rainstorm
      Given the unit is non-motorized
      And the weather is Rainstorm
      When it moves into a Clear hex across a Wadi
      Then it may not move there

    @case-8.37 @case-29.0 @wip
    Scenario: A road across a Wadi costs two more CP during a Rainstorm
      Given the unit is motorized
      And the weather is Rainstorm
      When it moves along a road into a Clear hex across a Wadi
      Then it spends 2½ Capability Points
