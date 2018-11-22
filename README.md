# Backlog

# Assumptions

- No transport, or business level authentication is required.
- The input archive file need not be protected with a mutex

# Scope Not Achieved

- Anything docker
- Deliberate handling of signals
- Normalising phone numbers

# Bugs

- There are no tests
- The ingestor is treating the csv header line as a user record.
- The database is volatile (in memory map) - pending bringing in a Redis 
  service.

# Robustness and Maintability

- Currenly assumes fixed name for csv file inside archive.
- Currently assumes CSV file is well-formed
- There is no top level readme overview.
- Needs explicit dependency managagement

# Flexibility

- Replace hard coded port binding with env vars.
- Replace timeouts with env vars.

# Investigate

- Log is going to stderr?
- The ingestor is reporting 104 User's found - check that 3 are duplicates
- Can we find csv file inside archive with random access rather than iterating?