# [3360.Stone Removal Game][title]

## Description
Alice and Bob are playing a game where they take turns removing stones from a pile, with Alice going first.

- Alice starts by removing **exactly** 10 stones on her first turn.
- For each subsequent turn, each player removes **exactly** 1 fewer stone than the previous opponent.

The player who cannot make a move loses the game.

Given a positive integer `n`, return `true` if Alice wins the game and `false` otherwise.

**Example 1:**

```
Input: n = 12

Output: true

Explanation:

Alice removes 10 stones on her first turn, leaving 2 stones for Bob.
Bob cannot remove 9 stones, so Alice wins.
```

**Example 2:**

```
Input: n = 1

Output: false

Explanation:

Alice cannot remove 10 stones, so Alice loses.
```

## 结语

如果你同我一样热爱数据结构、算法、LeetCode，可以关注我 GitHub 上的 LeetCode 题解：[awesome-golang-algorithm][me]

[title]: https://leetcode.com/problems/stone-removal-game/
[me]: https://github.com/kylesliu/awesome-golang-algorithm
