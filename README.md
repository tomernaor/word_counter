# Word Counter – Backend Engineer Assignment

This project implements the Backend Engineer home assignment:

> Fetch a list of essays, count the top **N** valid words from all essays combined,  
> and print the result as pretty JSON.

A **valid word** is defined as:

- At least 3 characters long.
- Contains only alphabetic characters (A–Z, a–z).
- Exists in the provided lexicon file.

---

## ✨ Features

- **Concurrent fetching** of essay URLs using a worker pool.
- **Rate limiting** to control outgoing requests per second.
- **Configurable** worker pool size, rate limit, and top-N output.
- **Thread-safe counters** to aggregate word frequencies.
- **Progress bars** to show fetching and error counts in real time.
- **Unit tests** for core components.

---

## 🚀 How to Run

You can run the application in several ways:

### Run directly with `go run`

```bash
go run main.go \
  -lexicon input_files/words.txt \
  -urls input_files/endg-urls.txt \
  -workers 10 \
  -rate 10 \
  -top 10 \
  -stats
```
### Run using the compiled binary
```bash
make build   # builds ./word_counter
./word_counter \
-urls=input_files/endg-urls.txt \
-workers=20 \
-rate=50 \
-top=10 \
-stats
```

### Run inside Docker
```bash
docker build -t word_counter .
docker run --rm -it word_counter \
-urls=input_files/endg-urls.txt \
-workers=20 \
-rate=50 \
-top=10 \
-stats
```


## ⚙️ Command-line Flags

You can see the available flags with `--help`:

| Flag          | Default                       | Description                             |
|---------------|-------------------------------|-----------------------------------------|
| `-lexicon`    | `input_files/words.txt`       | Path to the lexicon file               |
| `-urls`       | `input_files/endg-urls.txt` | Path to the URLs file                  |
| `-workers`    | `10`                          | Number of concurrent workers           |
| `-rate`       | `10`                          | Requests per second (rate limit)       |
| `-top`        | `10`                          | Number of top words to output          |
| `-stats`      | `false`                       | Show extended statistics if set        |

## 📊 Example Output
### Without -stats
```json
 [
  {
    "word": "new",
    "count": 19118388
  },
  {
    "word": "and",
    "count": 15164238
  },
  {
    "word": "the",
    "count": 10156480
  },
  {
    "word": "must",
    "count": 7949516
  },
  {
    "word": "screen",
    "count": 7387606
  },
  {
    "word": "not",
    "count": 7061722
  },
  {
    "word": "solid",
    "count": 6487854
  },
  {
    "word": "mmmm",
    "count": 5480000
  },
  {
    "word": "for",
    "count": 4196876
  },
  {
    "word": "been",
    "count": 4140400
  }
]
```

### With -stats
```json
{
  "statistics": {
    "total_essay": 40000,
    "lexicon": {
      "total_valid": 415699,
      "total_invalid": 50848
    },
    "pool_fetch": [
      {
        "id": 0,
        "count": 1994
      },
      {
        "id": 1,
        "count": 1991
      },
      {
        "id": 2,
        "count": 2005
      },
      {
        "id": 3,
        "count": 1997
      },
      {
        "id": 4,
        "count": 2001
      },
      {
        "id": 5,
        "count": 2000
      },
      {
        "id": 6,
        "count": 1996
      },
      {
        "id": 7,
        "count": 2001
      },
      {
        "id": 8,
        "count": 2000
      },
      {
        "id": 9,
        "count": 2000
      },
      {
        "id": 10,
        "count": 2002
      },
      {
        "id": 11,
        "count": 2001
      },
      {
        "id": 12,
        "count": 2000
      },
      {
        "id": 13,
        "count": 2006
      },
      {
        "id": 14,
        "count": 1999
      },
      {
        "id": 15,
        "count": 2000
      },
      {
        "id": 16,
        "count": 2007
      },
      {
        "id": 17,
        "count": 1999
      },
      {
        "id": 18,
        "count": 2001
      },
      {
        "id": 19,
        "count": 2000
      }
    ],
    "analyze_time": "00:50:11"
  },
  "top_n_words": [
    {
      "word": "new",
      "count": 19118388
    },
    {
      "word": "and",
      "count": 15164238
    },
    {
      "word": "the",
      "count": 10156480
    },
    {
      "word": "must",
      "count": 7949516
    },
    {
      "word": "screen",
      "count": 7387606
    },
    {
      "word": "not",
      "count": 7061722
    },
    {
      "word": "solid",
      "count": 6487854
    },
    {
      "word": "mmmm",
      "count": 5480000
    },
    {
      "word": "for",
      "count": 4196876
    },
    {
      "word": "been",
      "count": 4140400
    }
  ]
}
```