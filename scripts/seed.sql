COPY codes(code, phrase, title, content, meaning)
FROM '/seed/codes.csv'
WITH (FORMAT csv, HEADER true);
