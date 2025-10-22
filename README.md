# Elo Rating System for 5v5 Games

## Overview

This project implements an Elo-based rating system specifically designed for 5v5 competitive games like Counter-Strike, Valorant, or Rainbow Six Siege. It calculates player MMR (Matchmaking Rating) changes based on the results of a match.

## Features

- Calculates MMR changes for all players in a 5v5 match.
- Considers team strength by comparing the average MMR of each team.
- Factors in individual performance through a performance rating.
- Adjusts MMR gains/losses based on the round difference in the final score.

## How it works

The MMR change calculation is based on several factors:

1.  **Expected Score:** Based on the Elo rating system, we calculate the probability of winning for each team based on their average MMR.
2.  **Individual Performance:** A player's performance rating (which can be derived from stats like K/D ratio, ADR, etc.) acts as a multiplier, rewarding players who perform well and penalizing those who underperform.
3.  **Round Modifier:** The magnitude of the win/loss is considered. A decisive victory (e.g., 13-2) will result in a larger MMR change than a close match (e.g., 13-11).

## Running Tests

To run the tests for this project, execute the following command in the root directory:

```sh
go test ./...
```
