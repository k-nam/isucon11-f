# Setup

# MEMO
- COMMIT query occurs even when I removed all tx in main.go. So I check general log.
- Check CPU consumption with `docker stats` to see the big picture

# Idea (easy)
- Remove transactions (tx)
- Use mysql connection pool (max idle conn)

# Using memory


# Can't use

# Chores

# TODO

# Bug note

# Must do before end
- Disable db log
- Disable golang log
- Disable nginx log
- Check https://blog.recruit.co.jp/rtc/2021/04/26/isucon-2021-winter/


# Score
- Initial 9300
- Stage 1. Remove some tx 9300
- Stage 2. (from 1) Add index on unread_announcements(user_id) 10500
- Stage 3. (from 2) Use db conn pool 12800
- Stage 4. (from 3) 